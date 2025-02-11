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

	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set in .env file")
	}
	productsServiceURL := os.Getenv("PRODUCTS_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("PRODUCTS_SERVICE_URL not set in .env file")
	}

	go cfg.HealthCheckLoop()

	http.Handle("/auth/", middleware.Logging(proxy.ReverseProxyHandler(authServiceURL)))
	http.Handle("/products/", middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)))

	brandPermissions := map[string]string{
		"GET":    "",
		"POST":   "admin",
		"PUT":    "admin",
		"DELETE": "admin",
	}
	http.Handle("/products/brands", middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions))
	http.Handle("/products/categories", middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions))
	http.Handle("/products/inventory", middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions))
	http.Handle("/products/products", middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions))
	http.Handle("/products/types", middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions))

	// Start Gateway Server
	port := ":8080"
	log.Printf("Gateway running on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
