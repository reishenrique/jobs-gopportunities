package recruiter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reishenrique/job-gopportunities/schemas"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")

	ctx.JSON(code, gin.H {
		"message": message,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")

	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data": data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	ErrorCode string `json:"data"`
}

type CreateRecruiterResponse struct {
	Message string `json:"message"`
	Data schemas.RecruitersResponse `json:"data"`
}

type ShowRecruiterResponse struct {
	Message string `json:"message"`
	Data schemas.RecruitersResponse `json:"data"`
}

type ListRecruitersResponse struct {
	Message string `json:"message"`
	Data []schemas.RecruitersResponse `json:"data"`
}

type UpdateRecruiterResponse struct {
	Message string `json:"message"`
	Data schemas.RecruitersResponse `json:"data"`
}

type DeleteRecruiterResponse struct {
	Message string `json:"message"`
	Data schemas.RecruitersResponse `json:"data"`
}