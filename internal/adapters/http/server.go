package http

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/internal/service"
)

// Server represents the HTTP server
type Server struct {
	srv             *gin.Engine
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

	// Return a new Server instance with Gin engine
	return &Server{
		srv:             r,
		productHandler:  productHandler,
		reviewHandler:   reviewHandler,
		categoryHandler: categoryHandler,
	}
}

// Run starts the HTTP server
func (s *Server) Run(ctx context.Context) {
	// Start the server on port 9000
	if err := s.srv.Run(":8080"); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
