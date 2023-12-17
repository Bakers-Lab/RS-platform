package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register handles user registration
func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
