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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not set in .env file")
	}

	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("AUTH_SERVICE_URL not set in .env file")
	}
	productsServiceURL := os.Getenv("PRODUCTS_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("PRODUCTS_SERVICE_URL not set in .env file")
	}
	blogsServiceURL := os.Getenv("BLOGS_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("BLOGS_SERVICE_URL not set in .env file")
	}
	reviewsServiceURL := os.Getenv("REVIEWS_SERVICE_URL")
	if authServiceURL == "" {
		log.Fatal("REVIEWS_SERVICE_URL not set in .env file")
	}

	go cfg.HealthCheckLoop()

	http.Handle("/auth/", middleware.CORS(middleware.Logging(proxy.ReverseProxyHandler(authServiceURL))))
	http.Handle("/products/", middleware.CORS(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL))))
	http.Handle("/blogs/", middleware.CORS(middleware.Logging(proxy.ReverseProxyHandler(blogsServiceURL))))
	http.Handle("/reviews/", middleware.CORS(middleware.Logging(proxy.ReverseProxyHandler(reviewsServiceURL))))

	brandPermissions := map[string]string{
		"GET":    "",
		"POST":   "admin",
		"PUT":    "admin",
		"DELETE": "admin",
	}

	http.Handle("/products/brands", middleware.CORS(middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions)))
	http.Handle("/products/categories", middleware.CORS(middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions)))
	http.Handle("/products/inventory", middleware.CORS(middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions)))
	http.Handle("/products/products", middleware.CORS(middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions)))
	http.Handle("/products/types", middleware.CORS(middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(productsServiceURL)), brandPermissions)))

	http.Handle("/blogs/blog-posts", middleware.CORS(middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(blogsServiceURL)), brandPermissions)))

	reviewPermissions := map[string]string{
		"GET":    "",
		"POST":   "",
		"PUT":    "admin",
		"DELETE": "admin",
	}

	http.Handle("/reviews/reviews", middleware.CORS(middleware.AuthMiddleware(middleware.Logging(proxy.ReverseProxyHandler(reviewsServiceURL)), reviewPermissions)))

	// Start Gateway Server
	log.Printf("Gateway running on %s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
