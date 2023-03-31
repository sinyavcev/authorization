package app

import (
	"github.com/sinyavcev/authorization/config"
)

func Run(config config.Config) {
	server := NewServer(config)
	server.Run()
}
