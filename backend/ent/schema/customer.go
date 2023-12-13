package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("full_name"),
		field.Time("birth_date"),
		field.Time("registration_date").
			Default(time.Now()),
		field.String("customer_picture_url").
			Optional(),
		field.Bool("status"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("membership", Membership.Type),
		edge.To("personnel_rel_customer", PersonnelRelCustomer.Type),
	}
}
