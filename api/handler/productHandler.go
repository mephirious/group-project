package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/usecase"
)

type ProductHandler struct {
	useCase usecase.ProductUseCase
}

func NewProductHandler(router *gin.Engine, useCase usecase.ProductUseCase) {
	handler := &ProductHandler{useCase: useCase}

	router.GET("/products", handler.GetAllProducts)
	router.GET("/products/:id", handler.GetProductByID)
	router.POST("/products", handler.CreateProduct)
	router.PUT("/products/:id", handler.UpdateProduct)
	router.DELETE("/products/:id", handler.DeleteProduct)
}

func (p *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := p.useCase.GetAllProducts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed:  %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, products)
	slog.Info(fmt.Sprintf("Method %s finished successful", c.Request.Method))
}

func (p *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := p.useCase.GetProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed:  %s", c.Request.Method, err))
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		slog.Error(fmt.Sprintf("Method %s failed:  %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, product)
	slog.Info(fmt.Sprintf("Method %s finished successful", c.Request.Method))
}

func (p *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed:  %s", c.Request.Method, err))
		return
	}

	err = p.useCase.CreateProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed:  %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusCreated, product)
	slog.Info(fmt.Sprintf("Method %s finished successful", c.Request.Method))
}

func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product domain.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed:  %s", c.Request.Method, err))
		return
	}

	product.ID = id
	err = p.useCase.UpdateProduct(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed:  %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, product)
	slog.Info(fmt.Sprintf("Method %s finished successful", c.Request.Method))
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := h.useCase.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed:  %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	slog.Info(fmt.Sprintf("Method %s finished successful", c.Request.Method))
}
