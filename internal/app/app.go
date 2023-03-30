package app

import (
	"github.com/sinyavcev/authorization/config"
	controller "github.com/sinyavcev/authorization/internal/controller/http"
)

func Run(config config.Config) {

	repository := controller.NewRepository(db)
	service := controller.NewService(repository)
	controller := controller.NewController(service)

	server := NewServer(config, controller)
	server.Run()
}
