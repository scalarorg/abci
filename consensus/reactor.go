package consensus

import (
	cmtc "github.com/cometbft/cometbft/consensus"
	cstypes "github.com/cometbft/cometbft/consensus/types"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/types"
)

// Reactor defines a reactor for the consensus service.
type Reactor struct {
	p2p.BaseReactor // BaseService + p2p.Switch

	conS *State

	mtx      cmtsync.RWMutex
	waitSync bool
	eventBus *types.EventBus
	rs       *cstypes.RoundState

	Metrics *cmtc.Metrics
}

type ReactorOption func(*Reactor)

// ReactorMetrics sets the metrics
func ReactorMetrics(metrics *cmtc.Metrics) ReactorOption {
	return func(conR *Reactor) { conR.Metrics = metrics }
}

// NewReactor returns a new Reactor with the given
// consensusState.
func NewReactor(consensusState *State, waitSync bool, options ...ReactorOption) *Reactor {
	conR := &Reactor{
		conS:     consensusState,
		waitSync: waitSync,
		rs:       consensusState.GetRoundState(),
		Metrics:  cmtc.NopMetrics(),
	}
	conR.BaseReactor = *p2p.NewBaseReactor("Consensus", conR)

	for _, option := range options {
		option(conR)
	}

	return conR
}

// OnStart implements BaseService by subscribing to events, which later will be
// broadcasted to other peers and starting state if we're not in block sync.
func (conR *Reactor) OnStart() error {
	conR.Logger.Info("Reactor ", "waitSync", conR.WaitSync())

	// start routine that computes peer statistics for evaluating peer quality
	// go conR.peerStatsRoutine()

	//conR.subscribeToBroadcastEvents()
	//go conR.updateRoundStateRoutine()

	if !conR.WaitSync() {
		err := conR.conS.Start()
		if err != nil {
			return err
		}
	}

	return nil
}

// OnStop implements BaseService by unsubscribing from events and stopping
// state.
func (conR *Reactor) OnStop() {
	//conR.unsubscribeFromBroadcastEvents()
	if err := conR.conS.Stop(); err != nil {
		conR.Logger.Error("Error stopping consensus state", "err", err)
	}
	if !conR.WaitSync() {
		conR.conS.Wait()
	}
}

// SetEventBus sets event bus.
func (conR *Reactor) SetEventBus(b *types.EventBus) {
	conR.eventBus = b
	conR.conS.SetEventBus(b)
}

// WaitSync returns whether the consensus reactor is waiting for state/block sync.
func (conR *Reactor) WaitSync() bool {
	conR.mtx.RLock()
	defer conR.mtx.RUnlock()
	return conR.waitSync
}

//--------------------------------------

// // subscribeToBroadcastEvents subscribes for new round steps and votes
// // using internal pubsub defined on state to broadcast
// // them to peers upon receiving.
// func (conR *Reactor) subscribeToBroadcastEvents() {
// 	const subscriber = "consensus-reactor"
// 	if err := conR.conS.evsw.AddListenerForEvent(subscriber, types.EventNewRoundStep,
// 		func(data cmtevents.EventData) {
// 			conR.broadcastNewRoundStepMessage(data.(*cstypes.RoundState))
// 		}); err != nil {
// 		conR.Logger.Error("Error adding listener for events", "err", err)
// 	}

// 	if err := conR.conS.evsw.AddListenerForEvent(subscriber, types.EventValidBlock,
// 		func(data cmtevents.EventData) {
// 			conR.broadcastNewValidBlockMessage(data.(*cstypes.RoundState))
// 		}); err != nil {
// 		conR.Logger.Error("Error adding listener for events", "err", err)
// 	}

// 	if err := conR.conS.evsw.AddListenerForEvent(subscriber, types.EventVote,
// 		func(data cmtevents.EventData) {
// 			conR.broadcastHasVoteMessage(data.(*types.Vote))
// 		}); err != nil {
// 		conR.Logger.Error("Error adding listener for events", "err", err)
// 	}

// }

// func (conR *Reactor) unsubscribeFromBroadcastEvents() {
// 	const subscriber = "consensus-reactor"
// 	conR.conS.evsw.RemoveListener(subscriber)
// }
