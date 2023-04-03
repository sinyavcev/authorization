package app

import (
	"github.com/sinyavcev/authorization/config"
	"github.com/sinyavcev/authorization/internal/controller/http"
)

func Run(config config.Config) {

	httpConrtoller := http.NewController(usecases)

	server := NewServer(config, httpConrtoller)
	server.Run()
}
