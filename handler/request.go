package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role is required", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company is required", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location is required", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link is required", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote is required", "bool")
	}
	if r.Salary == 0 {
		return errParamIsRequired("salary is required", "ini64")
	}
	return nil
}
