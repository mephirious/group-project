package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/mephirious/group-project/internal/model"
	usecase "github.com/mephirious/group-project/internal/usecase/product"
)

// ProductHandler is responsible for handling product-related HTTP requests.
type ProductHandler struct {
	productUsecase *usecase.Product
}

func (h *ProductHandler) Routes(r *gin.Engine) {
	r.POST("/products", h.Create)
	r.GET("/products", h.GetAll)
	r.GET("/products/:id", h.Get)
	r.PUT("/products/:id", h.Update)
	r.DELETE("/products/:id", h.Delete)
}

// Create creates a new product.
func (h *ProductHandler) Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.productUsecase.Create(c, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// GetAll retrieves all products.
func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.productUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Get retrieves a product by its ID.
func (h *ProductHandler) Get(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productUsecase.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Update updates an existing product.
func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.productUsecase.Update(c, id, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// Delete deletes a product.
func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := h.productUsecase.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
