package node

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/libs/log"
	cmtos "github.com/cometbft/cometbft/libs/os"
	"github.com/spf13/cobra"
)

var (
	config = DefaultConfig()
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
)

var genesisHash []byte

// AddNodeFlags exposes some common configuration options on the command-line
// These are exposed for convenience of commands embedding a CometBFT node.
func AddNodeFlags(cmd *cobra.Command) {

	// abci flags
	cmd.Flags().String(
		"sclaris_addr",
		config.ScalarisAddr,
		"scalaris server address")
	//cmd.Flags().String("abci", cometBFTConfig.ABCI, "specify abci transport (socket | grpc)")
}

// // TrapSignal catches the SIGTERM/SIGINT and executes cb function. After that it exits
// // with code 0.
// func TrapSignal(logger log.Logger, cb func()) {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
// 	go func() {
// 		for sig := range c {
// 			logger.Info("signal trapped", "msg", log.NewLazySprintf("captured %v, exiting...", sig))
// 			if cb != nil {
// 				cb()
// 			}
// 			os.Exit(0)
// 		}
// 	}()
// }

// NewRunNodeCmd returns the command that allows the CLI to start a node.
// It can be used with a custom PrivValidator and in-process ABCI application.
func NewRunNodeCmd(nodeProvider Provider) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Aliases: []string{"node", "run"},
		Short:   "Run the scalaris Connector",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkGenesisHash(config.CometConfig); err != nil {
				return err
			}

			n, err := nodeProvider(config, logger)
			if err != nil {
				return fmt.Errorf("failed to create node: %w", err)
			}

			if err := n.Start(); err != nil {
				return fmt.Errorf("failed to start node: %w", err)
			}

			logger.Info("Started node", "nodeInfo", n.Switch().NodeInfo())

			// Stop upon receiving SIGTERM or CTRL-C.
			cmtos.TrapSignal(logger, func() {
				if n.IsRunning() {
					if err := n.Stop(); err != nil {
						logger.Error("unable to stop the node", "error", err)
					}
				}
			})

			// Run forever.
			select {}
		},
	}

	AddNodeFlags(cmd)
	return cmd
}

func checkGenesisHash(config *cfg.Config) error {
	if len(genesisHash) == 0 || config.Genesis == "" {
		return nil
	}

	// Calculate SHA-256 hash of the genesis file.
	f, err := os.Open(config.GenesisFile())
	if err != nil {
		return fmt.Errorf("can't open genesis file: %w", err)
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return fmt.Errorf("error when hashing genesis file: %w", err)
	}
	actualHash := h.Sum(nil)

	// Compare with the flag.
	if !bytes.Equal(genesisHash, actualHash) {
		return fmt.Errorf(
			"--genesis_hash=%X does not match %s hash: %X",
			genesisHash, config.GenesisFile(), actualHash)
	}

	return nil
}
