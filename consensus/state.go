package consensus

import (
	cfg "github.com/cometbft/cometbft/config"
	cstypes "github.com/cometbft/cometbft/consensus/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/libs/service"
	cmtsync "github.com/cometbft/cometbft/libs/sync"
	sm "github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/types"
	sstate "github.com/scalarorg/abci/state"
)

// State handles execution of the consensus algorithm.
// It processes votes and proposals, and upon reaching agreement,
// commits blocks to the chain and executes them against the application.
// The internal state machine receives input from peers, the internal validator, and from a timer.
type State struct {
	service.BaseService
	// config details
	config        *cfg.ConsensusConfig
	privValidator types.PrivValidator // for signing votes

	// store blocks and commits
	blockStore sm.BlockStore

	// create and execute blocks
	blockExec *sstate.BlockExecutor

	// notify us if txs are available
	// txNotifier txNotifier

	// internal state
	mtx cmtsync.RWMutex
	cstypes.RoundState
	state sm.State // State until height-1.
	// privValidator pubkey, memoized for the duration of one block
	// to avoid extra requests to HSM
	privValidatorPubKey crypto.PubKey

	// we use eventBus to trigger msg broadcasts in the reactor,
	// and to notify external subscribers, eg. through a websocket
	eventBus *types.EventBus
}

// NewState returns a new State.
func NewState(
	config *cfg.ConsensusConfig,
	state sm.State,
	blockExec *sstate.BlockExecutor,
	blockStore sm.BlockStore,
	// txNotifier txNotifier,
	// evpool evidencePool,
	// options ...StateOption,
) *State {
	cs := &State{
		config:     config,
		blockExec:  blockExec,
		blockStore: blockStore,
		// txNotifier: txNotifier,
		// peerMsgQueue:     make(chan msgInfo, msgQueueSize),
		// internalMsgQueue: make(chan msgInfo, msgQueueSize),
		// timeoutTicker:    NewTimeoutTicker(),
		// statsMsgQueue:    make(chan msgInfo, msgQueueSize),
		// done:             make(chan struct{}),
		// doWALCatchup:     true,
		// wal:              nilWAL{},
		// evpool:           evpool,
		// evsw:             cmtevents.NewEventSwitch(),
		// metrics:          cmtc.NopMetrics(),
	}
	// for _, option := range options {
	// 	option(cs)
	// }
	// // set function defaults (may be overwritten before calling Start)
	// cs.decideProposal = cs.defaultDecideProposal
	// cs.doPrevote = cs.defaultDoPrevote
	// cs.setProposal = cs.defaultSetProposal

	// We have no votes, so reconstruct LastCommit from SeenCommit.
	if state.LastBlockHeight > 0 {
		// In case of out of band performed statesync, the state store
		// will have a state but no extended commit (as no block has been downloaded).
		// If the height at which the vote extensions are enabled is lower
		// than the height at which we statesync, consensus will panic because
		// it will try to reconstruct the extended commit here.
		// cs.reconstructLastCommit(state)
	}

	//cs.updateToState(state)

	// NOTE: we do not call scheduleRound0 yet, we do that upon Start()

	cs.BaseService = *service.NewBaseService(nil, "State", cs)

	return cs
}

// SetEventBus sets event bus.
func (cs *State) SetEventBus(b *types.EventBus) {
	cs.eventBus = b
	cs.blockExec.SetEventBus(b)
}

// String returns a string.
func (cs *State) String() string {
	// better not to access shared variables
	return "ConsensusState"
}

// GetRoundState returns a shallow copy of the internal consensus state.
func (cs *State) GetRoundState() *cstypes.RoundState {
	cs.mtx.RLock()
	rs := cs.RoundState // copy
	cs.mtx.RUnlock()
	return &rs
}

// GetValidators returns a copy of the current validators.
func (cs *State) GetValidators() (int64, []*types.Validator) {
	cs.mtx.RLock()
	defer cs.mtx.RUnlock()
	return cs.state.LastBlockHeight, cs.state.Validators.Copy().Validators
}

// SetPrivValidator sets the private validator account for signing votes. It
// immediately requests pubkey and caches it.
func (cs *State) SetPrivValidator(priv types.PrivValidator) {
	cs.mtx.Lock()
	defer cs.mtx.Unlock()

	cs.privValidator = priv

	if err := cs.updatePrivValidatorPubKey(); err != nil {
		cs.Logger.Error("failed to get private validator pubkey", "err", err)
	}
}

// updatePrivValidatorPubKey get's the private validator public key and
// memoizes it. This func returns an error if the private validator is not
// responding or responds with an error.
func (cs *State) updatePrivValidatorPubKey() error {
	if cs.privValidator == nil {
		return nil
	}

	pubKey, err := cs.privValidator.GetPubKey()
	if err != nil {
		return err
	}
	cs.privValidatorPubKey = pubKey
	return nil
}

// OnStart loads the latest state via the WAL, and starts the timeout and
// receive routines.
func (cs *State) OnStart() error {
	return nil
}

// OnStop implements service.Service.
func (cs *State) OnStop() {

}
