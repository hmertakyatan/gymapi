package requests

import "time"

type DtoCustomerRequest struct {
	FullName           string    `json:"full_name,omitempty"`
	BirthDateStr       string    `json:"birth_date,omitempty"`
	RegistrationDate   time.Time `json:"registration_date,omitempty"`
	CustomerPictureURL string    `json:"customer_picture_url"`
	Status             bool      `json:"status,omitempty"`
}
