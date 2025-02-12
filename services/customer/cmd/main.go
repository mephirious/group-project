package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/customer/internal/cart"
	"github.com/mephirious/group-project/services/customer/internal/config"
	"github.com/mephirious/group-project/services/customer/internal/database"
	"github.com/mephirious/group-project/services/customer/pkg/logger"
)

func main() {
	// Load app config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize logger
	logger.InitLogger()
	logger.Log.Info("Starting Cart Service... 🚀")

	// Connect to MongoDB
	database.ConnectMongoDB(cfg)
	logger.Log.Info("Connected to MongoDB: ", cfg)

	// Initialize repository and handler
	repo := cart.NewCartRepository()
	handler := cart.NewCartHandler(repo)

	// Create routes
	r := gin.Default()
	cart.RegisterRoutes(r, handler)

	logger.Log.Info("Cart Service is running on port 8080")
	r.Run(":8080")
}
