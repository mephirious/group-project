package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/mephirious/group-project/internal/model"
	usecase "github.com/mephirious/group-project/internal/usecase/category"
)

// CategoryHandler is responsible for handling category-related HTTP requests.
type CategoryHandler struct {
	categoryUsecase *usecase.Category
}

func (h *CategoryHandler) Routes(r *gin.Engine) {
	r.POST("/categories", h.Create)
	r.GET("/categories", h.GetAll)
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
