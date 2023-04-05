package app

import (
	"github.com/sinyavcev/authorization/config"
	"github.com/sinyavcev/authorization/internal/common/logger"
	"github.com/sinyavcev/authorization/internal/controller/http"
)

func Run(config config.Config) {
	logger := logger.NewLogger(config.LoggerConfig)

	httpConrtoller := http.NewController(usecases, logger)
	server := NewServer(config, httpConrtoller)
	server.Run()
}
