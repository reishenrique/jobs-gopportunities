package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	
	ctx.JSON(code, gin.H {
		"message": message,
		"errorCode": code,
	})
} 

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Context-type", "application/json")

	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data": data,
	})

}