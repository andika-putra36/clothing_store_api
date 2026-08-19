package admin

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	router.POST("admin/products", h.InsertProduct)
	router.PATCH("admin/products/:id", h.UpdateProduct)
	router.DELETE("admin/products/:id", h.DeleteProduct)
}
