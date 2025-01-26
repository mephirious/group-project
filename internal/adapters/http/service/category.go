package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/mephirious/group-project/internal/model"
	"github.com/mephirious/group-project/internal/usecase"
)

// CategoryHandler is responsible for handling category-related HTTP requests.
type CategoryHandler struct {
	categoryUsecase *usecase.Category
}

// NewCategoryHandler creates a new CategoryHandler and registers routes.
func NewCategoryHandler(r *gin.Engine, categoryUsecase *usecase.Category) {
	handler := &CategoryHandler{categoryUsecase}

	r.POST("/categories", handler.Create)
	r.GET("/categories", handler.GetAll)
}

// Create creates a new category.
func (h *CategoryHandler) Create(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := h.categoryUsecase.Create(c, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

// GetAll retrieves all categories.
func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.categoryUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
