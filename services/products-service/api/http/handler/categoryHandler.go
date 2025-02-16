package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"`
}

type CategoryHandler struct {
	useCase usecase.CategoryUseCase
}

func NewCategoryHandler(router *gin.Engine, useCase usecase.CategoryUseCase) {
	handler := &CategoryHandler{useCase: useCase}

	router.GET("/categories", handler.GetAllCategories)
	router.GET("/categories/:id", handler.GetCategoryByID)
	router.GET("/categories/name/:name", handler.GetCategoryByName)
	router.POST("/categories", handler.CreateCategory)
	router.PUT("/categories/:id", handler.UpdateCategory)
	router.DELETE("/categories/:id", handler.DeleteCategory)
}

func (c *CategoryHandler) GetAllCategories(g *gin.Context) {
	categories, err := c.useCase.GetAllCategories(g.Request.Context())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, categories)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (c *CategoryHandler) GetCategoryByID(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	category, err := c.useCase.GetCategoryByID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if category == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, category)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (c *CategoryHandler) GetCategoryByName(g *gin.Context) {
	name := g.Param("name")

	category, err := c.useCase.GetCategoryByName(g.Request.Context(), name)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if category == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, category)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (c *CategoryHandler) CreateCategory(g *gin.Context) {
	var req CategoryRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "category_name is required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	category := domain.Category{
		ID:           primitive.NewObjectID(),
		CategoryName: req.CategoryName,
		CreatedAt:    time.Now(), // Ensure creation timestamp is set
		UpdatedAt:    time.Now(),
	}

	if err := c.useCase.CreateCategory(g.Request.Context(), &category); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusCreated, category)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (c *CategoryHandler) UpdateCategory(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	// Fetch the existing category to preserve created_at
	existingCategory, err := c.useCase.GetCategoryByID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Category not found", g.Request.Method))
		return
	}

	var req CategoryRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "category_name is required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	// Preserve created_at and update other fields
	category := domain.Category{
		ID:           objID,
		CategoryName: req.CategoryName,
		CreatedAt:    existingCategory.CreatedAt, // Preserve original created_at
		UpdatedAt:    time.Now(),                 // Update timestamp
	}

	if err := c.useCase.UpdateCategory(g.Request.Context(), &category); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, category)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (c *CategoryHandler) DeleteCategory(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	err = c.useCase.DeleteCategory(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}
