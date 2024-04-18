package handler

import (
	"github.com/reishenrique/job-gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializeRecruiterHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}