package recruiter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reishenrique/job-gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Delete recruiter
// @Description Delete a recruiter by id
// @Tags Recruiters
// @Accept json
// @Produce json
// @Param id query string true "Recruiter identification"
// @Success 200 {object} DeleteRecruiterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /recruiter [delete]
func DeleteRecruiterHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	recruiter := schemas.Recruiters{}

	if err := db.First(&recruiter, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("recruiter with id: %v not found", id))
		return
	}

	if err := db.Delete(&recruiter).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting recruiter with id: %v", id))
		return
	}

	sendSuccess(ctx, "delete-recruiter", recruiter)
}