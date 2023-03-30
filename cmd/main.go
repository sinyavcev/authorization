package main

import (
	"github.com/sinyavcev/authorization/config"
	"github.com/sinyavcev/authorization/internal/app"
	"log"
)

func main() {
	config, err := config.LoadConfig("config")
	if err != nil {
		log.Fatal(err)
	}
	app.Run(config)
}
