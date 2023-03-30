package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpServer HttpServer
}

type HttpServer struct {
	Port string
	Host string
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
