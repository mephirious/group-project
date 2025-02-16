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

type BlogPostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Image   string `json:"image"`
}

type BlogPostHandler struct {
	useCase usecase.BlogPostUseCase
}

func NewBlogPostHandler(router *gin.Engine, useCase usecase.BlogPostUseCase) {
	handler := &BlogPostHandler{useCase: useCase}

	router.GET("/blog-posts", handler.GetAllBlogPosts)
	router.GET("/blog-posts/:id", handler.GetBlogPostByID)
	router.GET("/blog-posts/title/:title", handler.GetBlogPostByTitle)
	router.POST("/blog-posts", handler.CreateBlogPost)
	router.PUT("/blog-posts/:id", handler.UpdateBlogPost)
	router.DELETE("/blog-posts/:id", handler.DeleteBlogPost)
}

func (h *BlogPostHandler) GetAllBlogPosts(g *gin.Context) {
	posts, err := h.useCase.GetAllBlogPosts(g.Request.Context())
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, posts)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *BlogPostHandler) GetBlogPostByID(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog post ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	post, err := h.useCase.GetBlogPostByID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if post == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Blog post not found", g.Request.Method))
		return
	}

	g.JSON(http.StatusOK, post)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *BlogPostHandler) GetBlogPostByTitle(g *gin.Context) {
	title := g.Param("title")

	post, err := h.useCase.GetBlogPostByTitle(g.Request.Context(), title)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	if post == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Blog post not found", g.Request.Method))
		return
	}

	g.JSON(http.StatusOK, post)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *BlogPostHandler) CreateBlogPost(g *gin.Context) {
	var req BlogPostRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "title and content are required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	post := domain.BlogPost{
		ID:        primitive.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		Image:     req.Image,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.useCase.CreateBlogPost(g.Request.Context(), &post); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusCreated, post)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *BlogPostHandler) UpdateBlogPost(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog post ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	// Fetch the existing post to preserve created_at
	existingPost, err := h.useCase.GetBlogPostByID(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		slog.Error(fmt.Sprintf("Method %s failed: Blog post not found", g.Request.Method))
		return
	}

	var req BlogPostRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "title and content are required"})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	// Preserve created_at and update other fields
	post := domain.BlogPost{
		ID:        objID,
		Title:     req.Title,
		Content:   req.Content,
		Image:     req.Image,
		CreatedAt: existingPost.CreatedAt, // Preserve original created_at
		UpdatedAt: time.Now(),             // Update timestamp
	}

	if err := h.useCase.UpdateBlogPost(g.Request.Context(), &post); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, post)
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}

func (h *BlogPostHandler) DeleteBlogPost(g *gin.Context) {
	id := g.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog post ID"})
		slog.Error(fmt.Sprintf("Method %s failed: Invalid ID format", g.Request.Method))
		return
	}

	err = h.useCase.DeleteBlogPost(g.Request.Context(), objID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(fmt.Sprintf("Method %s failed: %s", g.Request.Method, err))
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Blog post deleted"})
	slog.Info(fmt.Sprintf("Method %s finished successfully", g.Request.Method))
}
