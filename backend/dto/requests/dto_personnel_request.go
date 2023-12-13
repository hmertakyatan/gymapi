package requests

type DtoPersonnelRequest struct {
	Name         string  `json:"name,omitempty"`
	Salary       float64 `json:"salary,omitempty"`
	Description  string  `json:"description,omitempty"`
	BirthDateStr string  `json:"birth_date,omitempty"`
	StartDate    string  `json:"start_date,omitempty"`
	Status       bool    `json:"status,omitempty"`
}
