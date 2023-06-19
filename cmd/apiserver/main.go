package main

import (
	"flag"
	"github.com/diplodocx/Test-Go-API/intrernal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configs-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	config, _ := apiserver.NewConfig(configPath)
	server := apiserver.New(config)
	if err := apiserver.Start(server); err != nil {
		log.Fatal(err)
	}
}
