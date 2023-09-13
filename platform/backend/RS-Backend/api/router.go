package api

import (
	v1 "RS-Backend/api/v1"
	"RS-Backend/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the routes for the application.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1Group := api.Group("/v1")
		{
			v1Group.POST("/register", v1.Register)
			v1Group.POST("/login", v1.Login)
			v1Group.GET("/users", middleware.AuthMiddleware(), v1.GetUsers) // Protected route
		}
	}

	return r
}
