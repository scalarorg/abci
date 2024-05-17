package node

type Config struct {
	ScalarisAddr string `mapstructure:"scalaris_addr"`
}

// DefaultConfig returns a default configuration for a CometBFT node.
func DefaultConfig() *Config {
	return &Config{
		ScalarisAddr: "tcp://127.0.0.1:8080",
	}
}
