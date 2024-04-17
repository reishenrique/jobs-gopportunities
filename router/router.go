package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router with default Gin settings
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Running the server
	router.Run("0.0.0.0:" + port)
}