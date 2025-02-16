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

type TypeRequest struct {
	TypeName string `json:"type_name" binding:"required"`
}

type TypeHandler struct {
	useCase usecase.TypeUseCase
}

func NewTypeHandler(router *gin.Engine, useCase usecase.TypeUseCase) {
	handler := &TypeHandler{useCase: useCase}

	router.GET("/types", handler.GetAllTypes)
	router.GET("/types/:id", handler.GetTypeByID)
	router.GET("/types/name/:name", handler.GetTypeByName)
	router.POST("/types", handler.CreateType)
	router.PUT("/types/:id", handler.UpdateType)
	router.DELETE("/types/:id", handler.DeleteType)
}

func (t *TypeHandler) GetAllTypes(g *gin.Context) {
	types, err := t.useCase.GetAllTypes(g.Request.Context())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, types)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (t *TypeHandler) GetTypeByID(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	typeEntity, err := t.useCase.GetTypeByID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if typeEntity == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Type not found"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, typeEntity)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (t *TypeHandler) GetTypeByName(g *gin.Context) {
	name := g.Param("name")

	typeEntity, err := t.useCase.GetTypeByName(g.Request.Context(), name)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if typeEntity == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Type not found"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, typeEntity)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (t *TypeHandler) CreateType(g *gin.Context) {
	var req TypeRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "type_name is required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	typeEntity := domain.Type{
		ID:        primitive.NewObjectID(),
		TypeName:  req.TypeName,
		CreatedAt: time.Now(), // Ensure creation timestamp is set
		UpdatedAt: time.Now(),
	}

	if err := t.useCase.CreateType(g.Request.Context(), &typeEntity); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusCreated, typeEntity)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (t *TypeHandler) UpdateType(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	// Fetch the existing type to preserve created_at
	existingType, err := t.useCase.GetTypeByID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Type not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Type not found", g.Request.Method))
		return
	}

	var req TypeRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "type_name is required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	// Preserve created_at and update other fields
	typeEntity := domain.Type{
		ID:        objID,
		TypeName:  req.TypeName,
		CreatedAt: existingType.CreatedAt, // Preserve original created_at
		UpdatedAt: time.Now(),             // Update timestamp
	}

	if err := t.useCase.UpdateType(g.Request.Context(), &typeEntity); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, typeEntity)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (t *TypeHandler) DeleteType(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	err = t.useCase.DeleteType(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Type deleted"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}
