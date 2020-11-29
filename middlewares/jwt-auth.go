package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gireeshcse/go-gin/entity"
)

// JWTAuth To validate Token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Bearer")

		token, err := entity.ValidateToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "UnAuthorized Access"})
		}

		c.Set("token", token)
		c.Next()
	}
}
