package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/mephirious/group-project/services/products-service/adapter/mongo"
	"github.com/mephirious/group-project/services/products-service/api/http/handler"
	"github.com/mephirious/group-project/services/products-service/config"
	"github.com/mephirious/group-project/services/products-service/repository"
	"github.com/mephirious/group-project/services/products-service/usecase"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("Error occured while loading config: %s", err))
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := db.ConnectToMongoDB(ctx, cfg.Database.URI)
	if err != nil {
		slog.Error(fmt.Sprintf("Error occured while connecting mongoDB: %s", err))
		os.Exit(1)
	}
	defer db.DisconnectFromMongoDB(ctx, client)

	database := client.Database(cfg.Database.Name)
	brandRepository := repository.NewBrandRepository(database)
	brandUseCase := usecase.NewBrandUseCase(brandRepository)
	categoryRepository := repository.NewCategoryRepository(database)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	typeRepository := repository.NewTypeRepository(database)
	typeUseCase := usecase.NewTypeUseCase(typeRepository)
	inventoryRepository := repository.NewInventoryRepository(database)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository)
	productRepository := repository.NewProductRepository(database)
	productUseCase := usecase.NewProductUseCase(productRepository, categoryRepository, brandRepository, typeRepository)

	router := gin.Default()
	handler.NewBrandHandler(router, brandUseCase)
	handler.NewCategoryHandler(router, categoryUseCase)
	handler.NewTypeHandler(router, typeUseCase)
	handler.NewInventoryHandler(router, inventoryUseCase)
	handler.NewProductHandler(router, productUseCase)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	slog.Info("Starting server", slog.Int("port", cfg.Server.Port))
	err = router.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		slog.Error("Failed to start server on port %d", cfg.Server.Port)
		os.Exit(1)
	}
}
