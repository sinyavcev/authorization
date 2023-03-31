package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HttpServer HttpServer
}

type HttpServer struct {
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func LoadConfig(path string) (Config, error) {
	var (
		config Config
	)
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("viper.ReadInConfig: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("viper.Unmarshal: %w", err)
	}

	return config, nil
}
