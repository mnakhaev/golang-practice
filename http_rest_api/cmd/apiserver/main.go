package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gopherschool/http-rest-api/internal/app/apiserver"
	"log"
)

// Set path to TOML config as a flag before running binary file
var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse() // parse all variables defined in init()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}