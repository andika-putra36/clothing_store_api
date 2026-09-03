package product

import (
	"clothing_store_api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	public := router.Group("/")
	{
		public.GET("/products", h.GetProducts)
		public.GET("/products/:id", h.GetProduct)
	}
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/products", h.InsertProduct)
		protected.PATCH("/products/:id", h.UpdateProduct)
		protected.DELETE("/products/:id", h.DeleteProduct)
	}

}
