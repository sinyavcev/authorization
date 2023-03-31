package app

import (
	"github.com/sinyavcev/authorization/config"
	"github.com/sinyavcev/authorization/internal/controller/http"
	"github.com/sinyavcev/authorization/internal/usecases"
)

func Run(config config.Config) {
	usecases := usecases.NewBackendUsecases()
	conrtoller := http.NewController(usecases)

	server := NewServer(config, conrtoller)
	server.Run()
}
