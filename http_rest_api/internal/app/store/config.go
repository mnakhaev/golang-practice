package store

// Config for the store
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig is helper function which returns pointer on config
// TODO: so it returns new copy of config or what?
func NewConfig() *Config {
	return &Config{}
}
