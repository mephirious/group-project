package cart

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes регистрирует все маршруты для Cart Service
func RegisterRoutes(router *gin.Engine, handler *CartHandler) {
	routes := router.Group("/cart")
	{
		routes.POST("/", handler.AddToCartHandler)
		routes.GET("/:user_id", handler.GetCartHandler)
		routes.PUT("/:user_id/item/:item_id", handler.UpdateCartItemHandler)
		routes.DELETE("/:user_id/item/:item_id", handler.RemoveFromCartHandler)
		routes.DELETE("/:user_id", handler.ClearCartHandler)
	}
}
