package cart

import (
	"clothing_store_api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/cart", h.InsertToCart)
		protected.DELETE("/cart/:id", h.DeleteFromCart)
		protected.GET("/cart", h.GetCart)
	}
}
