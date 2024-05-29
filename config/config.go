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

func LoadConfig() (Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath("config")

	if err := v.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := gotenv.Load(); err != nil {
		return Config{}, err
	}

	var c Config

	if err := v.Unmarshal(&c); err != nil {
		return Config{}, err
	}

	return c, nil
}
