package rpc

import (
	"github.com/cometbft/cometbft/crypto"
	sm "github.com/cometbft/cometbft/state"
)

var (
	// set by Node
	env *Environment
)

// SetEnvironment sets up the given Environment.
// It will race if multiple Node call SetEnvironment.
func SetEnvironment(e *Environment) {
	env = e
}

type Environment struct {
	BlockStore sm.BlockStore
	PubKey     crypto.PubKey
}
