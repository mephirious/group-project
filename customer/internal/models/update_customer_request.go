package models

type UpdateCustomerRequest struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
