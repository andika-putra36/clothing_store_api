package cart

import (
	"clothing_store_api/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) InsertToCart(c *gin.Context) {
	claims, ok := jwt.GetClaims(c)
	if !ok {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	var input InsertToCartRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request body is required",
		})
		return
	}

	err = h.service.InsertToCart(claims.CustomerID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product is added to cart successfully!",
	})
}
