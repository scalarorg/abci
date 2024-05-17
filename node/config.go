package node

import (
	cometcfg "github.com/cometbft/cometbft/config"
)

type Config struct {
	CometConfig  *cometcfg.Config
	ScalarisAddr string `mapstructure:"scalaris_addr"`
}

// DefaultConfig returns a default configuration for a CometBFT node.
func DefaultConfig() *Config {
	return &Config{
		ScalarisAddr: "tcp://127.0.0.1:8080",
		CometConfig:  cometcfg.DefaultConfig(),
	}
}
