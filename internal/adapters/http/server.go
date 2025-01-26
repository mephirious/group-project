package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	service "github.com/mephirious/group-project/internal/adapters/http/service"
)

// Server represents the HTTP server
type Server struct {
	srv             *http.Server
	productHandler  *service.ProductHandler
	reviewHandler   *service.ReviewHandler
	categoryHandler *service.CategoryHandler
}

// NewServer creates a new Server with the provided handlers
func NewServer(productHandler *service.ProductHandler, reviewHandler *service.ReviewHandler, categoryHandler *service.CategoryHandler) *Server {
	// Create a new Gin engine
	r := gin.Default()

	// Register handlers with the Gin router
	productHandler.Routes(r)
	reviewHandler.Routes(r)
	categoryHandler.Routes(r)

	// Create an http.Server with the Gin router as the handler
	return &Server{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: r,
		},
		productHandler:  productHandler,
		reviewHandler:   reviewHandler,
		categoryHandler: categoryHandler,
	}
}

// Run starts the HTTP server
func (s *Server) Run(ctx context.Context) {
	// Start the server in a separate goroutine
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server failed to start:", err)
		}
	}()
	<-ctx.Done()
	s.Stop(ctx)
}

// Stop gracefully shuts down the server
func (s *Server) Stop(ctx context.Context) {
	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Attempt to shut down the server gracefully
	if err := s.srv.Shutdown(ctx); err != nil {
		fmt.Println("Server shutdown failed:", err)
	} else {
		fmt.Println("Server stopped gracefully")
	}
}
