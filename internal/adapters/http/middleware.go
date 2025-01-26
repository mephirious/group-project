package http

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs the request details.
func LoggerMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()
	duration := time.Since(start)
	log.Printf("Request: %s %s took %v", c.Request.Method, c.Request.URL, duration)
}
