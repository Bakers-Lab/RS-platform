package middleware

import (
	"RS-Backend/services"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the JWT token from the request header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization token required"})
			return
		}

		userID, err := services.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization token"})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
