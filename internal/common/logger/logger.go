package logger

import (
	"fmt"
	"os"

	"github.com/sinyavcev/authorization/config"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(config config.LoggerConfig) (*Logger, error) {
	log := logrus.New()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return &Logger{log}, fmt.Errorf("logrus.ParseLevel: %w", err)
	}
	log.SetLevel(level)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return &Logger{log}, nil
}
