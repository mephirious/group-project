package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		URI  string
		Name string
	}
	Logging struct {
		Level string
	}
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		slog.Error(fmt.Sprintf("Error occured while loading config: %s", err))
	}

	config := &Config{}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		port = 8080
	}

	config.Server.Port = port
	config.Database.URI = os.Getenv("DATABASE_URI")
	config.Database.Name = os.Getenv("DATABASE_NAME")
	config.Logging.Level = os.Getenv("LOGGING_LEVEL")

	return config, nil
}
