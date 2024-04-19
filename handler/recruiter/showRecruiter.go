package recruiter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reishenrique/job-gopportunities/schemas"
)

func ShowRecruiterHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
		"queryParameter").Error())
		return
	}

	recruiter := schemas.Recruiters{}

	if err := db.First(&recruiter, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Recruiter with id: %v not found", id))
		return
	}

	sendSuccess(ctx, "show-recruiter", recruiter)
}