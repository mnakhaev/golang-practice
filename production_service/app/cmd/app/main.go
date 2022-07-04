package main

import (
	"log"

	"production_service/internal/app"
	"production_service/internal/config"
	"production_service/pkg/logging"
)

func main() {
	log.Print("initializing config")
	cfg := config.GetConfig()

	log.Print("initializing logger")
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running application")
	a.Run()
}
