package main

import (
	"context"
	"fmt"

	"github.com/mephirious/group-project/internal/adapters/http"
	"github.com/mephirious/group-project/internal/app"
	"github.com/mephirious/group-project/internal/usecase"
)

func main() {
	// Initialize usecases
	productUsecase := usecase.NewProductUsecase()
	reviewUsecase := usecase.NewReviewUsecase()
	categoryUsecase := usecase.NewCategoryUsecase()

	// Create handlers with usecases
	productHandler := &http.ProductHandler{ProductUsecase: productUsecase}
	reviewHandler := &http.ReviewHandler{ReviewUsecase: reviewUsecase}
	categoryHandler := &http.CategoryHandler{CategoryUsecase: categoryUsecase}

	// Create a new server with the handlers
	server := http.NewServer(productHandler, reviewHandler, categoryHandler)

	// Initialize the App with the server
	app := app.New(server)

	// Run the server in the context of the app
	ctx := context.Background()
	app.SimpleServer.Run(ctx)

	// Optionally, add graceful shutdown handling
	select {
	case <-ctx.Done():
		// Graceful shutdown or cleanup if necessary
		fmt.Println("Shutting down the server...")
	}
}
