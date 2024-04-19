package recruiter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reishenrique/job-gopportunities/schemas"
)

func CreateRecruiterHandler(ctx *gin.Context) {
	request := CreateRecruiterRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	recruiter := schemas.Recruiters {
		FullName: request.FullName,
		Recruiter: *request.Recruiter,
		CompanyName: request.CompanyName,
		CompanyEmail: request.CompanyEmail,
		CompanyWebsite: request.CompanyWebsite,
		Phone: request.Phone,
	}

	if err := db.Create(&recruiter).Error; err != nil {
		logger.Errorf("Error creating recruiter: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error saving recruiter to the database")
		return
	}

	sendSuccess(ctx, "create-recruiter", recruiter)
}