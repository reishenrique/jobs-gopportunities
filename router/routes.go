package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/reishenrique/job-gopportunities/handler/opening"

	docs "github.com/reishenrique/job-gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath
	
	v1 := router.Group(basePath)

	{
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}