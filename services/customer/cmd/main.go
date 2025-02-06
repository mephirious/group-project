package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/customer/internal/cart"
	"github.com/mephirious/group-project/services/customer/internal/database"
	"github.com/mephirious/group-project/services/customer/pkg/logger"
)

func main() {
	// Инициализируем логгер
	logger.InitLogger()
	logger.Log.Info("Starting Cart Service... 🚀")

	// Используем локальную MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // ✅ Подключаемся к локальному MongoDB
	}

	// Подключаемся к MongoDB
	database.ConnectMongoDB(mongoURI)
	logger.Log.Info("Connected to MongoDB: ", mongoURI)

	// Инициализируем репозиторий и обработчик
	repo := cart.NewCartRepository()
	handler := cart.NewCartHandler(repo)

	// Создаем маршруты
	r := gin.Default()
	cart.RegisterRoutes(r, handler)

	logger.Log.Info("Cart Service is running on port 8080")
	r.Run(":8080")
}
