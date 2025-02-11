package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mephirious/group-project/services/gateway-service/config"
	"github.com/mephirious/group-project/services/gateway-service/internal/middleware"
	"github.com/mephirious/group-project/services/gateway-service/internal/proxy"
)

func main() {
	log.Printf("STARTING GATEWAY SERVICE (%v environment) ...", os.Getenv("SERVICE_ENV"))

	cfg := config.NewConfig()

	// Load environment variables
	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set in .env file")
	}
	productsServiceURL := os.Getenv("PRODUCTS_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("PRODUCTS_SERVICE_URL not set in .env file")
	}

	// Start automatic health checking
	go cfg.HealthCheckLoop()

	// Set up routes
	http.Handle("/auth/", middleware.Logging(proxy.ReverseProxyHandler(authServiceURL)))
	http.Handle("/products/", middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)))

	// Start Gateway Server
	port := ":8080"
	log.Printf("Gateway running on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
