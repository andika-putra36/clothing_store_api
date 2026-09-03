package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	public := router.Group("/")
	{
		public.POST("/auth/login", h.LogIn)
		public.POST("/auth/refresh-token", h.RefreshToken)
		public.POST("/auth/register-customer", h.RegisterCustomer)
		public.POST("/auth/register-admin", h.RegisterAdmin)
	}

	// protected := router.Group("/")
	// protected.Use(middleware.AuthMiddleware())
	// {
	// 	protected.GET("auth/get-claims", func(c *gin.Context) {
	// 		claims := c.MustGet("claims").(*jwt.Claims)
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"user_id":     claims.UserID,
	// 			"email":       claims.Email,
	// 			"role_id":     claims.RoleID,
	// 			"customer_id": claims.CustomerID,
	// 			"admin_id":    claims.AdminID,
	// 		})
	// 	})
	// }

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"test": "test",
	// 	})
	// })
}
