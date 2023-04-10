package app

import (
	"github.com/sinyavcev/authorization/config"
	"github.com/sinyavcev/authorization/internal/common/logger"
	"github.com/sinyavcev/authorization/internal/controller/http"
	"log"
	"os"
)

func Run(config config.Config) {
	logger, err := logger.NewLogger(config.LoggerConfig)
	if err != nil {
		log.Printf("logger.NewLogger: %w", err)
		os.Exit(1)
	}

	httpConrtoller := http.NewController(usecases, logger)
	server := NewServer(config, httpConrtoller)
	server.Run()
}
