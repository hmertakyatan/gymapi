package requests

type DtoMembershipTypeRequest struct {
	Name            string  `json:"name,omitempty"`
	Description     string  `json:"description,omitempty"`
	MembershipMonth uint8   `json:"membership_month,omitempty"`
	Price           float64 `json:"price,omitempty"`
}
