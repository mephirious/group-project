package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/mephirious/group-project/internal/model"
	"github.com/mephirious/group-project/internal/usecase"
)

// ReviewHandler is responsible for handling review-related HTTP requests.
type ReviewHandler struct {
	reviewUsecase *usecase.Review
}

// NewReviewHandler creates a new ReviewHandler and registers routes.
func NewReviewHandler(r *gin.Engine, reviewUsecase *usecase.Review) {
	handler := &ReviewHandler{reviewUsecase}

	r.POST("/reviews/:productID", handler.Add)
	r.GET("/reviews/:productID", handler.GetAll)
}

// Add adds a new review for a product.
func (h *ReviewHandler) Add(c *gin.Context) {
	productID := c.Param("productID")
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.reviewUsecase.Add(c, productID, review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

// GetAll retrieves all reviews for a product.
func (h *ReviewHandler) GetAll(c *gin.Context) {
	productID := c.Param("productID")
	reviews, err := h.reviewUsecase.GetAll(c, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
