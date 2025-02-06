package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRequest struct {
	ModelName      string                `json:"model_name" binding:"required"`
	Specifications domain.Specifications `json:"specifications" binding:"required"`
	Content        string                `json:"content" binding:"required"`
	LaptopImage    []string              `json:"laptop_image" binding:"required"`
}

type ProductHandler struct {
	useCase usecase.ProductUseCase
}

func NewProductHandler(router *gin.Engine, useCase usecase.ProductUseCase) {
	handler := &ProductHandler{useCase: useCase}

	router.GET("/products", handler.GetAllProducts)
	router.GET("/products/:id", handler.GetProductByID)
	router.GET("/products/model/:model_name", handler.GetProductByModelName)
	router.POST("/products", handler.CreateProduct)
	router.PUT("/products/:id", handler.UpdateProduct)
	router.DELETE("/products/:id", handler.DeleteProduct)
}

func (h *ProductHandler) GetAllProducts(g *gin.Context) {
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "10"))
	skip, _ := strconv.Atoi(g.DefaultQuery("skip", "0"))
	sortField := g.DefaultQuery("sortField", "model_name")
	sortOrder := g.DefaultQuery("sortOrder", "asc")
	search := g.DefaultQuery("search", "")

	products, err := h.useCase.GetAllProducts(g.Request.Context(), limit, skip, sortField, sortOrder, search)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, products)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ProductHandler) GetProductByID(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	product, err := h.useCase.GetProductByID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if product == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Product not found", g.Request.Method))
		return
	}

	g.JSON(http.StatusOK, product)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ProductHandler) GetProductByModelName(g *gin.Context) {
	modelName := g.Param("model_name")

	product, err := h.useCase.GetProductByName(g.Request.Context(), modelName)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if product == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Product not found", g.Request.Method))
		return
	}

	g.JSON(http.StatusOK, product)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ProductHandler) CreateProduct(g *gin.Context) {
	var req ProductRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "ModelName, Specifications, Content, and LaptopImage are required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	product := domain.Product{
		ID:             primitive.NewObjectID(),
		ModelName:      req.ModelName,
		Specifications: req.Specifications,
		Content:        req.Content,
		LaptopImage:    req.LaptopImage,
	}

	err := h.useCase.CreateProduct(g.Request.Context(), &product)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusCreated, product)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ProductHandler) UpdateProduct(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	var req ProductRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "ModelName, Specifications, Content, and LaptopImage are required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	product := domain.Product{
		ID:             objID,
		ModelName:      req.ModelName,
		Specifications: req.Specifications,
		Content:        req.Content,
		LaptopImage:    req.LaptopImage,
	}

	err = h.useCase.UpdateProduct(g.Request.Context(), &product)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, product)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ProductHandler) DeleteProduct(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	err = h.useCase.DeleteProduct(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}
