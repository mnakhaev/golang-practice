package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config will be based on env variables using `cleanenv` ;ibrary
	Config struct {
		IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
		IsDevelopment bool `env:"IS_DEV" env-default:"true"`
		Listen        struct {
			// the app can work using port or unix socket. if socket, then env socket file is required
			Type       string `env:"LISTEN_TYPE" env-default:"port"`
			BindIP     string `env:"BIND_IP" env-default:"0.0.0.0"`
			Port       string `env:"PORT" env-default:"10000"`
			SocketFile string `env:"SOCKET_FILE" env-default:"app.sock"`
		}
		AppConfig struct {
			LogLevel  string `env:"LOG_LEVEL" env-default:"trace"`
			AdminUser struct {
				Email    string `env:"ADMIN_EMAIL" env-default:"admin"`
				Password string `env:"ADMIN_PWD" env-default:"admin"`
			}
		}
		PostgreSQL struct {
			Username string `env:"PSQL_USERNAME" env-required:"true"`
			Password string `env:"PSQL_PASSWORD" env-required:"true"`
			Host     string `env:"PSQL_HOST" env-required:"true"`
			Port     string `env:"PSQL_PORT" env-required:"true"`
			Database string `env:"PSQL_DATABASE" env-required:"true"`
		}
	}
)

// Config will be initialized as singletone
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
