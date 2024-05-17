package client

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cometbft/cometbft/abci/types"
	cmtnet "github.com/cometbft/cometbft/libs/net"
	"github.com/cometbft/cometbft/libs/service"
	consensus "github.com/scalarorg/abci/api/proto/consensus"
)

var _ Client = (*grpcClient)(nil)

// A stripped copy of the remoteClient that makes
// synchronous calls using grpc.
type grpcClient struct {
	service.BaseService
	mustConnect bool

	client   consensus.ConsensusApiClient
	conn     *grpc.ClientConn
	chReqRes chan *ReqRes // dispatches "async" responses to callbacks *in order*, needed by mempool

	mtx   sync.Mutex
	addr  string
	err   error
	resCb func(*types.Request, *types.Response) // listens to all callbacks
}

func NewGRPCClient(addr string, mustConnect bool) Client {
	cli := &grpcClient{
		addr:        addr,
		mustConnect: mustConnect,
		// Buffering the channel is needed to make calls appear asynchronous,
		// which is required when the caller makes multiple async calls before
		// processing callbacks (e.g. due to holding locks). 64 means that a
		// caller can make up to 64 async calls before a callback must be
		// processed (otherwise it deadlocks). It also means that we can make 64
		// gRPC calls while processing a slow callback at the channel head.
		chReqRes: make(chan *ReqRes, 64),
	}
	cli.BaseService = *service.NewBaseService(nil, "grpcClient", cli)
	return cli
}

func dialerFunc(_ context.Context, addr string) (net.Conn, error) {
	return cmtnet.Connect(addr)
}

func (cli *grpcClient) OnStart() error {
	if err := cli.BaseService.OnStart(); err != nil {
		return err
	}

	// This processes asynchronous request/response messages and dispatches
	// them to callbacks.
	go func() {
		// Use a separate function to use defer for mutex unlocks (this handles panics)
		callCb := func(reqres *ReqRes) {
			cli.mtx.Lock()
			defer cli.mtx.Unlock()

			reqres.Done()

			// Notify client listener if set
			if cli.resCb != nil {
				cli.resCb(reqres.Request, reqres.Response)
			}

			// Notify reqRes listener if set
			reqres.InvokeCallback()
		}
		for reqres := range cli.chReqRes {
			if reqres != nil {
				callCb(reqres)
			} else {
				cli.Logger.Error("Received nil reqres")
			}
		}
	}()

RETRY_LOOP:
	for {
		conn, err := grpc.Dial(cli.addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(dialerFunc),
		)
		if err != nil {
			if cli.mustConnect {
				return err
			}
			cli.Logger.Error(fmt.Sprintf("abci.grpcClient failed to connect to %v.  Retrying...\n", cli.addr), "err", err)
			time.Sleep(time.Second * dialRetryIntervalSeconds)
			continue RETRY_LOOP
		}

		cli.Logger.Info("Dialed server. Waiting for echo.", "addr", cli.addr)
		client := types.NewABCIClient(conn)
		cli.conn = conn

	ENSURE_CONNECTED:
		for {
			_, err := client.Echo(context.Background(), &types.EchoRequest{Message: "hello"}, grpc.WaitForReady(true))
			if err == nil {
				break ENSURE_CONNECTED
			}
			cli.Logger.Error("Echo failed", "err", err)
			time.Sleep(time.Second * echoRetryIntervalSeconds)
		}

		cli.client = client
		return nil
	}
}

func (cli *grpcClient) OnStop() {
	cli.BaseService.OnStop()

	if cli.conn != nil {
		cli.conn.Close()
	}
	close(cli.chReqRes)
}

func (cli *grpcClient) StopForError(err error) {
	if !cli.IsRunning() {
		return
	}

	cli.mtx.Lock()
	if cli.err == nil {
		cli.err = err
	}
	cli.mtx.Unlock()

	cli.Logger.Error(fmt.Sprintf("Stopping abci.grpcClient for error: %v", err.Error()))
	if err := cli.Stop(); err != nil {
		cli.Logger.Error("Error stopping abci.grpcClient", "err", err)
	}
}

func (cli *grpcClient) Error() error {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()
	return cli.err
}

// Set listener for all responses
// NOTE: callback may get internally generated flush responses.
func (cli *grpcClient) SetResponseCallback(resCb Callback) {
	cli.mtx.Lock()
	cli.resCb = resCb
	cli.mtx.Unlock()
}

// ----------------------------------------

func (cli *grpcClient) InitTransaction(ctx context.Context) (consensus.ConsensusApi_InitTransactionClient, error) {
	client, err := cli.client.InitTransaction(ctx)
	if err != nil {
		cli.StopForError(err)
		return nil, err
	}
	return client, nil
}
