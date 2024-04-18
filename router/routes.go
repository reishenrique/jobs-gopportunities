package router

import (
	"github.com/gin-gonic/gin"
	openingHandler "github.com/reishenrique/job-gopportunities/handler/opening"
	recruiterHandler "github.com/reishenrique/job-gopportunities/handler/recruiter"

	docs "github.com/reishenrique/job-gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	openingHandler.InitializeHandler()
	recruiterHandler.InitializeRecruiterHandler()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath
	
	v1 := router.Group(basePath)

	{
		v1.POST("/opening", openingHandler.CreateOpeningHandler)
		v1.GET("/opening", openingHandler.ShowOpeningHandler)
		v1.PUT("/opening", openingHandler.UpdateOpeningHandler)
		v1.DELETE("/opening", openingHandler.DeleteOpeningHandler)
		v1.GET("/openings", openingHandler.ListOpeningsHandler)
	}
	
	{
		v1.POST("/recruiter", recruiterHandler.CreateRecruiterHandler)
		v1.GET("/recruiter")
		v1.PUT("/recruiter")
		v1.DELETE("/recruiter")
		v1.GET("/recruiters")
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}