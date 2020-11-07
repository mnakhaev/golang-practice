package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"` // Address used for web server start
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
