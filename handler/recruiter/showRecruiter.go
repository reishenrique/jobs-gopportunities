package recruiter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reishenrique/job-gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Show recruiter
// @Description Show a recruiter by id
// @Tags Recruiters
// @Accept json
// @Produce json
// @Param id query string true "Recruiter identification"
// @Success 200 {object} ShowRecruiterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /recruiter [get]
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