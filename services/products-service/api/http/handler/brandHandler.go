package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandRequest struct {
	BrandName string `json:"brand_name" binding:"required"`
}

type BrandHandler struct {
	useCase usecase.BrandUseCase
}

func NewBrandHandler(router *gin.Engine, useCase usecase.BrandUseCase) {
	handler := &BrandHandler{useCase: useCase}

	router.GET("/brands", handler.GetAllBrands)
	router.GET("/brands/:id", handler.GetBrandByID)
	router.GET("/brands/name/:name", handler.GetBrandByName)
	router.POST("/brands", handler.CreateBrand)
	router.PUT("/brands/:id", handler.UpdateBrand)
	router.DELETE("/brands/:id", handler.DeleteBrand)
}

func (b *BrandHandler) GetAllBrands(c *gin.Context) {
	brands, err := b.useCase.GetAllBrands(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, brands)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (b *BrandHandler) GetBrandByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", c.Request.Method))
		return
	}

	brand, err := b.useCase.GetBrandByID(c.Request.Context(), objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	if brand == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, brand)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (b *BrandHandler) GetBrandByName(c *gin.Context) {
	name := c.Param("name")

	brand, err := b.useCase.GetBrandByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	if brand == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, brand)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (b *BrandHandler) CreateBrand(c *gin.Context) {
	var req BrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brand_name is required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	brand := domain.Brand{
		ID:        primitive.NewObjectID(),
		BrandName: req.BrandName,
	}

	err := b.useCase.CreateBrand(c.Request.Context(), &brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusCreated, brand)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (b *BrandHandler) UpdateBrand(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", c.Request.Method))
		return
	}

	var req BrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "brand_name is required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	brand := domain.Brand{
		ID:        objID,
		BrandName: req.BrandName,
	}

	err = b.useCase.UpdateBrand(c.Request.Context(), &brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, brand)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (b *BrandHandler) DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", c.Request.Method))
		return
	}

	err = b.useCase.DeleteBrand(c.Request.Context(), objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}
