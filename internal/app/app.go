package app

import (
	"fmt"

	"github.com/sinyavcev/authorization/config"
	"github.com/sinyavcev/authorization/internal/common/logger"
	"github.com/sinyavcev/authorization/internal/controller/http"
)

func Run(config config.Config) {
	logger, err := logger.NewLogger(config.LoggerConfig)
	if err != nil {
		fmt.Errorf("logger.NewLogger: %v", err)
	}

	httpConrtoller := http.NewController(usecases, logger)
	server := NewServer(config, httpConrtoller)
	server.Run()
}
