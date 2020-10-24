package apiserver

import "github.com/gopherschool/http-rest-api/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"` // Address used for web server start
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
