package logger

import (
	"github.com/sinyavcev/authorization/config"
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger(config config.LoggerConfig) *logrus.Logger {
	log := logrus.New()

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to set logging level: %s", err.Error())
	} else {
		log.SetLevel(level)
	}
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	return log
}
