package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HttpConfig
	LoggerConfig
	PostgresConfig
}

type HttpConfig struct {
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type LoggerConfig struct {
	LogLevel string
}

type PostgresConfig struct {
	Port       string
	Host       string
	User       string
	DBName     string
	DBPassword string
}

func LoadConfig(path string) (Config, error) {
	var (
		config Config
	)
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("viper.ReadInConfig: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("viper.Unmarshal: %w", err)
	}

	return config, nil
}
