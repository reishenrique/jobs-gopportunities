package recruiter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reishenrique/job-gopportunities/schemas"
)

func ListRecruitersHandler(ctx *gin.Context) {
	recruiters := []schemas.Recruiters{}

	if err := db.Find(&recruiters).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing recruiters")
		return
	}

	sendSuccess(ctx, "list-recruiters", recruiters)
}	