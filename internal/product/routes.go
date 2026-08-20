package product

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	router.POST("/products", h.InsertProduct)
	router.PATCH("/products/:id", h.UpdateProduct)
	router.DELETE("/products/:id", h.DeleteProduct)
	router.GET("/products", h.GetProducts)
	router.GET("/products/:id", h.GetProduct)
}
