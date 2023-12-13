package requests

type DtoPaymentRequest struct {
	Type        string  `json:"type,omitempty"`
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
}
