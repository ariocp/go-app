package config

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type (
	Config struct {
		Server   ServerConfig
		Database DatabaseConfig
	}

	ServerConfig struct {
		Host string
		Port string
	}

	DatabaseConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		SSLMode  string
	}
)

func NewConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := gotenv.Load(); err != nil {
		panic(err)
	}

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		return Config{}, err
	}

	return c, nil
}
