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

type InventoryRequest struct {
	ProductID primitive.ObjectID `json:"product_id" binding:"required"`
	Quantity  int                `json:"quantity" binding:"required"`
}

type InventoryHandler struct {
	useCase usecase.InventoryUseCase
}

func NewInventoryHandler(router *gin.Engine, useCase usecase.InventoryUseCase) {
	handler := &InventoryHandler{useCase: useCase}

	router.GET("/inventory", handler.GetAllInventory)
	router.GET("/inventory/:id", handler.GetInventoryByID)
	router.GET("/inventory/product/:product_id", handler.GetInventoryByProductID)
	router.POST("/inventory", handler.CreateInventory)
	router.PUT("/inventory/:id", handler.UpdateInventory)
	router.DELETE("/inventory/:id", handler.DeleteInventory)
}

func (h *InventoryHandler) GetAllInventory(g *gin.Context) {
	inventories, err := h.useCase.GetAllInventories(g.Request.Context())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, inventories)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *InventoryHandler) GetInventoryByID(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	inventory, err := h.useCase.GetInventoryByProductID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if inventory == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Inventory not found", g.Request.Method))
		return
	}

	g.JSON(http.StatusOK, inventory)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *InventoryHandler) GetInventoryByProductID(g *gin.Context) {
	productID := g.Param("product_id")
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid product ID format", g.Request.Method))
		return
	}

	inventory, err := h.useCase.GetInventoryByProductID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if inventory == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Inventory not found", g.Request.Method))
		return
	}

	g.JSON(http.StatusOK, inventory)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *InventoryHandler) CreateInventory(g *gin.Context) {
	var req InventoryRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "ProductID and quantity are required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if req.Quantity < 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be a non-negative integer"})
		slog.Error(fmt.Sprintf("Method %s failed: Quantity must be a non-negative integer", g.Request.Method))
		return
	}

	inventory := domain.Inventory{
		ID:        primitive.NewObjectID(),
		ProductID: req.ProductID,
		Quantity:  uint32(req.Quantity),
	}

	err := h.useCase.CreateInventory(g.Request.Context(), &inventory)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusCreated, inventory)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *InventoryHandler) UpdateInventory(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	var req InventoryRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "ProductID and quantity are required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if req.Quantity < 0 {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be a non-negative integer"})
		slog.Error(fmt.Sprintf("Method %s failed: Quantity must be a non-negative integer", g.Request.Method))
		return
	}

	inventory := domain.Inventory{
		ID:        objID,
		ProductID: req.ProductID,
		Quantity:  uint32(req.Quantity),
	}

	err = h.useCase.UpdateInventory(g.Request.Context(), &inventory)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, inventory)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *InventoryHandler) DeleteInventory(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	err = h.useCase.DeleteInventory(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Inventory deleted"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}
