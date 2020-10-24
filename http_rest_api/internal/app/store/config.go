package store

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig returns pointer on config
func NewConfig() *Config {
	return &Config{}
}
