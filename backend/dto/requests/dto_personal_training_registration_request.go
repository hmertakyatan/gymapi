package requests

type DtoPersonalTrainingRegistartionRequest struct {
	CustomerID  string  `json:"customer_id,omitempty"`
	PersonnelID string  `json:"personnel_id,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"derscription,omitempty"`
}
