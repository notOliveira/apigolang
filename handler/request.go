package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type: %s) is required", name, typ)
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {

	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("request body is invalid or empty")
	}

	if r == nil {
		return fmt.Errorf("request body is required")
	}

	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}

// UpdateOpening

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {

	// If any field is provided, it's valid
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	
	// If no fields are provided, return an error
	return fmt.Errorf("at least one field must be provided for update")

}