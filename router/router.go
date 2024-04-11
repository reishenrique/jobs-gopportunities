package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the router with default Gin settings
	router := gin.Default()
	// Defining an endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Running api on port 8080 
	router.Run(":8080")
}