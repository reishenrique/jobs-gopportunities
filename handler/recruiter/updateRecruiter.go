package recruiter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reishenrique/job-gopportunities/schemas"
)

func UpdateRecruiterHandler(ctx *gin.Context) {
	request := UpdateRecruiterRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id",
		"queryParameter").Error())
		return
	}

	recruiter := schemas.Recruiters{}

	if err := db.First(&recruiter, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "recruiter not found")
		return
	}

	if request.FullName != "" {
		recruiter.FullName = request.FullName
	}
	if request.Recruiter != nil {
		recruiter.Recruiter = *request.Recruiter
	}
	if request.CompanyName != "" {
		recruiter.CompanyName = request.CompanyName
	}
	if request.CompanyEmail != "" {
		recruiter.CompanyEmail = request.CompanyEmail
	}
	if request.CompanyWebsite != "" {
		recruiter.CompanyWebsite = request.CompanyWebsite
	}
	if request.Phone > 0 {
		recruiter.Phone = request.Phone
	}

	if err := db.Save(&recruiter).Error; err != nil {
		logger.Errorf("error updating recruiter: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating recruiter")
		return
	}

	sendSuccess(ctx, "update-recruiter", recruiter)
}