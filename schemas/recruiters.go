package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Recruiters struct {
	gorm.Model
	FullName string
	Recruiter bool
	CompanyName string
	CompanyEmail string
	CompanyWebsite string
	Phone int64
}

type RecruitersResponse struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt,omitempty"`
	FullName string `json:"fullName"`
	Recruiter bool `json:"recruiter"`
	CompanyName string `json:"companyName"`
	CompanyEmail string `json:"companyEmail"`
	CompanyWebsite string `json:"companyWebsite"`
	Phone int64 `json:"phone"`
}