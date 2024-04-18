package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

type CreateRecruiterRequest struct {
	FullName string `json:"full_name"`
	Recruiter *bool `json:"recruiter"`
	CompanyName string `json:"company_name"`
	CompanyEmail string	`json:"company_email"`
	CompanyWebsite string `json:"company_website"`
	Phone int64 `json:"phone"`
}

func (r *CreateRecruiterRequest) Validate() error {
	if r.FullName == "" {
		return errParamIsRequired("full_name", "string")
	}

	if r.Recruiter == nil {
		return errParamIsRequired("recruiter", "bool")
	}

	if r.CompanyName == "" {
		return errParamIsRequired("company_name", "string")
	}

	if r.CompanyEmail == "" {
		return errParamIsRequired("company_email", "string")
	}

	if r.CompanyWebsite == "" {
		return errParamIsRequired("company_website", "string")
	}

	if r.Phone <= 0 {
		return errParamIsRequired("phone", "string")
	}

	return nil
}