package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/customer/internal/cart"
	"github.com/mephirious/group-project/services/customer/internal/database"
	"github.com/mephirious/group-project/services/customer/pkg/logger"
)

func main() {
	// Инициализируем логгер
	logger.InitLogger()
	logger.Log.Info("Starting Cart Service... 🚀")

	// Подключаемся к MongoDB
	database.ConnectMongoDB("mongodb://localhost:27017")
	logger.Log.Info("Connected to MongoDB")

	// Инициализируем репозиторий и обработчик
	repo := cart.NewCartRepository()
	handler := cart.NewCartHandler(repo)

	// Создаем маршруты
	r := gin.Default()
	r.GET("/cart/:user_id", handler.GetCartHandler)
	r.POST("/cart", handler.AddToCartHandler)
	r.DELETE("/cart/:user_id/item/:item_id", handler.RemoveFromCartHandler)
	r.DELETE("/cart/:user_id", handler.ClearCartHandler)

	logger.Log.Info("Cart Service is running on port 8080")
	r.Run(":8080")
}
