package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CartHandler - cart handler
type CartHandler struct {
	Repo *CartRepository
}

// NewCartHandler creating the handler
func NewCartHandler(repo *CartRepository) *CartHandler {
	return &CartHandler{Repo: repo}
}

// GetCartHandler return cart to the user
func (h *CartHandler) GetCartHandler(c *gin.Context) {
	userID := c.Param("user_id")
	cart, err := h.Repo.GetCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, cart)
}

// AddToCartHandler adds the product to the cart
func (h *CartHandler) AddToCartHandler(c *gin.Context) {
	var req struct {
		UserID    string `json:"user_id"`
		ProductID string `json:"product_id"`
		Amount    int    `json:"amount"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.Repo.AddToCart(c.Request.Context(), req.UserID, req.ProductID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added"})
}

// RemoveFromCartHandler deletes product by its ID
func (h *CartHandler) RemoveFromCartHandler(c *gin.Context) {
	userID := c.Param("user_id")
	itemID := c.Param("item_id")

	err := h.Repo.RemoveFromCart(c.Request.Context(), userID, itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item from cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed"})
}

// ClearCartHandler clears cart
func (h *CartHandler) ClearCartHandler(c *gin.Context) {
	userID := c.Param("user_id")

	err := h.Repo.ClearCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}

// UpdateCartItemHandler updates quantity of product or deletes, if amount < 1
func (h *CartHandler) UpdateCartItemHandler(c *gin.Context) {
	userID := c.Param("user_id")
	itemID := c.Param("item_id")

	var req struct {
		Amount int `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	message, err := h.Repo.UpdateCartItem(c.Request.Context(), userID, itemID, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
