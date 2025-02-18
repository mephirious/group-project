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

type InventoryRequest struct {
	ProductID    string `json:"product_id" binding:"required"`
	SerialNumber string `json:"serial_number" binding:"required"`
	Status       string `json:"status" binding:"required"`
}

type InventoryPOSTRequest struct {
	ProductID    string `json:"product_id" binding:"required"`
	SerialNumber string `json:"serial_number" binding:"required"`
}

type InventoryHandler struct {
	useCase usecase.InventoryUseCase
}

func NewInventoryHandler(router *gin.Engine, useCase usecase.InventoryUseCase) {
	handler := &InventoryHandler{useCase: useCase}

	router.GET("/inventories", handler.GetAllInventories)
	router.GET("/inventories/:id", handler.GetInventoryByID)
	router.GET("/inventories/product/:product_id", handler.GetInventoryByProductID)
	router.GET("/inventories/serial/:serial_number", handler.GetInventoryBySerialNumber)
	router.POST("/inventories", handler.CreateInventory)
	router.PUT("/inventories/:id", handler.UpdateInventory)
	router.DELETE("/inventories/:id", handler.DeleteInventory)
	router.GET("/inventories/product/:product_id/quantity", handler.GetProductQuantity)

	router.POST("/payment/start", handler.StartPayment)
	router.POST("/payment/cancel", handler.CancelPayment)
	router.POST("/payment/success", handler.PaymentSuccess)
}

func (i *InventoryHandler) GetAllInventories(c *gin.Context) {
	inventories, err := i.useCase.GetAllInventories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, inventories)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) GetInventoryByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", c.Request.Method))
		return
	}

	inventory, err := i.useCase.GetInventoryByID(c.Request.Context(), objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	if inventory == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Inventory not found", c.Request.Method))
		return
	}

	c.JSON(http.StatusOK, inventory)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) GetInventoryByProductID(c *gin.Context) {
	productID := c.Param("product_id")
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", c.Request.Method))
		return
	}

	inventories, err := i.useCase.GetInventoryByProductID(c.Request.Context(), objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, inventories)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) GetInventoryBySerialNumber(c *gin.Context) {
	serialNumber := c.Param("serial_number")

	inventory, err := i.useCase.GetInventoryBySerialNumber(c.Request.Context(), serialNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	if inventory == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Inventory not found", c.Request.Method))
		return
	}

	c.JSON(http.StatusOK, inventory)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) CreateInventory(c *gin.Context) {
	var req InventoryPOSTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid product ID", c.Request.Method))
		return
	}

	inventory := domain.Inventory{
		ID:           primitive.NewObjectID(),
		ProductID:    productID,
		SerialNumber: req.SerialNumber,
		Status:       "in_stock",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := i.useCase.CreateInventory(c.Request.Context(), &inventory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusCreated, inventory)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", c.Request.Method))
		return
	}

	var req InventoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid product ID", c.Request.Method))
		return
	}

	inventory := domain.Inventory{
		ID:           objID,
		ProductID:    productID,
		SerialNumber: req.SerialNumber,
		Status:       req.Status,
		UpdatedAt:    time.Now(),
	}

	if err := i.useCase.UpdateInventory(c.Request.Context(), &inventory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, inventory)
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", c.Request.Method))
		return
	}

	err = i.useCase.DeleteInventory(c.Request.Context(), objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory deleted"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) GetProductQuantity(c *gin.Context) {
	productID := c.Param("product_id")
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", c.Request.Method))
		return
	}

	quantity, err := i.useCase.GetProductQuantity(c.Request.Context(), objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": productID, "quantity": quantity})
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) StartPayment(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	err := i.useCase.ReserveProducts(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Products reserved"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) CancelPayment(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	err := i.useCase.CancelReservation(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation canceled"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}

func (i *InventoryHandler) PaymentSuccess(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	err := i.useCase.MarkProductsAsSold(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", c.Request.Method, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Products marked as sold"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", c.Request.Method))
}
