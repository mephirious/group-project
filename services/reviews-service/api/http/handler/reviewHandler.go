package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mephirious/group-project/services/products-service/domain"
	"github.com/mephirious/group-project/services/products-service/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewHandler struct {
	useCase usecase.ReviewUseCase
}

func NewReviewHandler(router *gin.Engine, useCase usecase.ReviewUseCase) {
	handler := &ReviewHandler{useCase: useCase}

	router.GET("/reviews", handler.GetAllReviews)
	router.GET("/reviews/:id", handler.GetReviewByID)
	router.GET("/reviews/customer/:customer_id", handler.GetReviewsByCustomerID)
	router.GET("/reviews/product/:product_id", handler.GetReviewsByProductID)
	router.PUT("/reviews/:id", handler.UpdateReview)
	router.POST("/reviews", handler.CreateReview)
	router.DELETE("/reviews/:id", handler.DeleteReview)
}

func (h *ReviewHandler) CreateReview(g *gin.Context) {
	var req struct {
		CustomerID string  `json:"customer_id" binding:"required"`
		ProductID  string  `json:"product_id" binding:"required"`
		Content    string  `json:"content" binding:"required"`
		Rating     float64 `json:"rating" binding:"required"`
	}

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	customerID, err := primitive.ObjectIDFromHex(req.CustomerID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid customer ID", g.Request.Method))
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid product ID", g.Request.Method))
		return
	}

	verified := false

	review := domain.Review{
		ID:         primitive.NewObjectID(),
		CustomerID: customerID,
		ProductID:  productID,
		Content:    req.Content,
		Rating:     req.Rating,
		Verified:   verified,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := h.useCase.CreateReview(g.Request.Context(), &review); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusCreated, review)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ReviewHandler) GetAllReviews(g *gin.Context) {
	limit, _ := strconv.Atoi(g.DefaultQuery("limit", "10"))
	skip, _ := strconv.Atoi(g.DefaultQuery("skip", "0"))
	sortField := g.DefaultQuery("sortField", "created_at")
	sortOrder := g.DefaultQuery("sortOrder", "desc")
	verified := g.DefaultQuery("verified", "")

	var verifiedPtr *bool
	if verified != "" {
		verifiedVal := verified == "true"
		verifiedPtr = &verifiedVal
	}

	reviews, err := h.useCase.GetAllReviews(g.Request.Context(), limit, skip, sortField, sortOrder, verifiedPtr)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, reviews)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ReviewHandler) GetReviewByID(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	review, err := h.useCase.GetReviewByID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if review == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Review not found", g.Request.Method))
		return
	}

	g.JSON(http.StatusOK, review)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ReviewHandler) GetReviewsByCustomerID(g *gin.Context) {
	customerID := g.Param("customer_id")
	objID, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	limit := 10
	if l := g.Query("limit"); l != "" {
		if limit, err = strconv.Atoi(l); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			slog.Error(fmt.Sprintf("Method %s failed: Invalid limit", g.Request.Method))
			return
		}
	}

	var verified *bool
	if v := g.Query("verified"); v != "" {
		verifiedVal := v == "true"
		verified = &verifiedVal
	}

	reviews, err := h.useCase.GetReviewsByCustomerID(g.Request.Context(), objID, limit, verified)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, reviews)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ReviewHandler) GetReviewsByProductID(g *gin.Context) {
	productID := g.Param("product_id")
	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	limit := 10
	if l := g.Query("limit"); l != "" {
		if limit, err = strconv.Atoi(l); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			slog.Error(fmt.Sprintf("Method %s failed: Invalid limit", g.Request.Method))
			return
		}
	}

	var verified *bool
	if v := g.Query("verified"); v != "" {
		verifiedVal := v == "true"
		verified = &verifiedVal
	}

	reviews, totalReviews, averageRating, err := h.useCase.GetReviewsByProductID(g.Request.Context(), objID, limit, verified)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	response := gin.H{
		"reviews":        reviews,
		"total":          totalReviews,
		"average_rating": averageRating,
	}

	g.JSON(http.StatusOK, response)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ReviewHandler) UpdateReview(g *gin.Context) {
	id := g.Param("id")
	reviewID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	var req struct {
		CustomerID string  `json:"customer_id" binding:"required"`
		ProductID  string  `json:"product_id" binding:"required"`
		Content    string  `json:"content" binding:"required"`
		Rating     float64 `json:"rating" binding:"required"`
		Verified   bool    `json:"verified"`
	}

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	customerID, err := primitive.ObjectIDFromHex(req.CustomerID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid customer ID", g.Request.Method))
		return
	}

	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid product ID", g.Request.Method))
		return
	}

	review := domain.Review{
		ID:         reviewID,
		CustomerID: customerID,
		ProductID:  productID,
		Content:    req.Content,
		Rating:     req.Rating,
		Verified:   req.Verified,
		UpdatedAt:  time.Now(),
	}

	if err := h.useCase.UpdateReview(g.Request.Context(), reviewID, &review); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *ReviewHandler) DeleteReview(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	if err := h.useCase.DeleteReview(g.Request.Context(), objID); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}
