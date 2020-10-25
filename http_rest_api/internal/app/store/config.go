package store

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig returns pointer on config
// TODO: so it returns new copy of config or what?
func NewConfig() *Config {
	return &Config{}
}
