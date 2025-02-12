package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds all the application configurations
type Config struct {
	MongoURI string
	Port     string
}

// LoadConfig loads configurations from .env file or environment variables
func LoadConfig() (*Config, error) {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"), // Default to local MongoDB
		Port:     getEnv("PORT", "8080"),                           // Default to 8080
	}, nil
}

// getEnv is a helper function to retrieve environment variables with a fallback default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
