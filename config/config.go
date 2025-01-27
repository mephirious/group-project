package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		URI  string `mapstructure:"uri"`
		Name string `mapstructure:"name"`
	} `mapstructure:"database"`
	Logging struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logging"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	currentDir, _ := os.Getwd()

	var configPath string
	if filepath.Base(currentDir) == "cmd" {
		configPath = "../config"
	} else {
		configPath = "./config"
	}

	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
