package config

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT     string
	Prefix   string
	Services []ServiceHealth
}

type ServiceHealth struct {
	Name string
	URL  string
}

func (c *Config) HealthCheckLoop() {
	for {
		for _, service := range c.Services {
			if service.URL == "" {
				log.Printf("[ERROR] %s has an empty URL (%s). Check your environment variables.", service.Name, service.URL)
				continue
			}

			// Make health check request
			resp, err := http.Get(service.URL)
			if err != nil || resp.StatusCode != http.StatusOK {
				log.Printf("[ERROR] %s is DOWN (%v)", service.Name, err)
			} else {
				log.Printf("[OK] %s is UP", service.Name)
			}
			if resp != nil {
				resp.Body.Close()
			}
		}
		time.Sleep(30 * time.Second)
	}
}
func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file")
	}

	cfg := &Config{
		PORT:   ":8080",
		Prefix: "/api/v1",
	}

	cfg.Services = append(cfg.Services, ServiceHealth{Name: "auth-service", URL: os.Getenv("AUTH_SERVICE_URL") + "/auth/api/v1/health"})
	cfg.Services = append(cfg.Services, ServiceHealth{Name: "products-service", URL: os.Getenv("PRODUCTS_SERVICE_URL") + "/brands"})
	cfg.Services = append(cfg.Services, ServiceHealth{Name: "blogs-service", URL: os.Getenv("BLOGS_SERVICE_URL") + "/blog-posts"})
	cfg.Services = append(cfg.Services, ServiceHealth{Name: "reviews-service", URL: os.Getenv("REVIEWS_SERVICE_URL") + "/reviews"})

	return cfg
}
