package middleware

import (
	"clothing_store_api/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization") // e.g. "Bearer <token>"
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.ValidateAccessToken(tokenString) // your token validation logic
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		// attach customer id to context so handlers can access it
		c.Set("claims", claims)
		c.Next() // proceed to the actual handler
	}
}

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "missing token",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		parts := strings.Split(authHeader, " ")
// 		if len(parts) != 2 || parts[0] != "Bearer" {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "invalid token format",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := jwt.ValidateAccessToken(parts[1])
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "invalid or expired token",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("claims", claims)
// 		c.Next()
// 	}
// }
