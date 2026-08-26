package transaction

import (
	"clothing_store_api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/transactions/:id", h.GetTransaction)
		protected.POST("/transactions/:id/cancel", h.CancelTransaction)
		protected.POST("/transactions/:id/accept", h.AcceptTransaction)
		protected.POST("/transactions/:id/reject", h.RejectTransaction)
		protected.POST("/transactions/:id/complete", h.CompleteTransaction)
		protected.POST("/transactions", h.InsertTransaction)
	}
}
