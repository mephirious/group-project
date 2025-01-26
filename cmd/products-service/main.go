package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mephirious/group-project/internal/adapters/http"
	"github.com/mephirious/group-project/pkg"
)

func main() {
	// Load configuration from config.yaml
	config, err := pkg.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create a context for the application lifecycle
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up a channel to listen for OS signals (graceful shutdown)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Initialize the database
	db, err := pkg.NewDB(ctx, *config)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}
	defer func() {
		if err := db.Close(ctx); err != nil {
			log.Printf("Error closing MongoDB connection: %v", err)
		}
	}()

	// Initialize HTTP handlers with the database repository
	productHandler := &http.service.ProductHandler{}
	reviewHandler := &http.service.ReviewHandler{}
	categoryHandler := &http.service.CategoryHandler{}

	// Create the server
	server := http.NewServer(productHandler, reviewHandler, categoryHandler)

	server.Run(ctx)

	// Wait for a termination signal
	<-signalChan
	log.Println("Shutdown signal received")

	// Perform graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	server.Stop(shutdownCtx)

	log.Println("Server stopped")
}
