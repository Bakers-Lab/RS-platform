package api

import (
	v1 "RS-Backend/api/v1"
	"RS-Backend/dal/db"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configures the routes for the application.
func SetupRouter(dB db.IDB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1Group := api.Group("/v1")
		{
			v1Group.GET("/test", v1.TestHandler)
			// 集成 Swagger
			v1.RegisterDatasetRoutes(v1Group, dB)
			v1.RegisterUserRoutes(v1Group, dB)
			v1.RegisterInferRoutes(v1Group, dB)
			v1.RegisterEvalRoutes(v1Group, dB)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
