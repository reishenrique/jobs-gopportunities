package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router with default Gin settings
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Running the server
	router.Run(":8080")
}