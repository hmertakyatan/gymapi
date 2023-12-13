package requests

type DtoMembershipRequest struct {
	StartDateStr     string  `json:"start_date"`
	AmountPaid       float64 `json:"amount_paid"`
	CustomerID       string  `json:"customer_id"`
	MembershipTypeID string  `json:"membership_type_id"`
}
