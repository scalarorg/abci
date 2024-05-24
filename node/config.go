package node

import (
	cmtcfg "github.com/cometbft/cometbft/config"
)

// Config defines the top level configuration for a CometBFT node
type Config struct {
	// Top level options use an anonymous struct
	cmtcfg.BaseConfig `mapstructure:",squash"`

	// Options for services
	RPC             *cmtcfg.RPCConfig             `mapstructure:"rpc"`
	P2P             *cmtcfg.P2PConfig             `mapstructure:"p2p"`
	Mempool         *cmtcfg.MempoolConfig         `mapstructure:"mempool"`
	StateSync       *cmtcfg.StateSyncConfig       `mapstructure:"statesync"`
	BlockSync       *cmtcfg.BlockSyncConfig       `mapstructure:"blocksync"`
	Consensus       *cmtcfg.ConsensusConfig       `mapstructure:"consensus"`
	Storage         *cmtcfg.StorageConfig         `mapstructure:"storage"`
	TxIndex         *cmtcfg.TxIndexConfig         `mapstructure:"tx_index"`
	Instrumentation *cmtcfg.InstrumentationConfig `mapstructure:"instrumentation"`
	ScalarisAddr    string                        `mapstructure:"scalaris_addr"`
}

// DefaultConfig returns a default configuration for a CometBFT node.
func DefaultConfig() *Config {
	return &Config{
		BaseConfig:      cmtcfg.DefaultBaseConfig(),
		RPC:             cmtcfg.DefaultRPCConfig(),
		P2P:             cmtcfg.DefaultP2PConfig(),
		Mempool:         cmtcfg.DefaultMempoolConfig(),
		StateSync:       cmtcfg.DefaultStateSyncConfig(),
		BlockSync:       cmtcfg.DefaultBlockSyncConfig(),
		Consensus:       cmtcfg.DefaultConsensusConfig(),
		Storage:         cmtcfg.DefaultStorageConfig(),
		TxIndex:         cmtcfg.DefaultTxIndexConfig(),
		Instrumentation: cmtcfg.DefaultInstrumentationConfig(),
		ScalarisAddr:    "127.0.0.1:8080",
	}
}
