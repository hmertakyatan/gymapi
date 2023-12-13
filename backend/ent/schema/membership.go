package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Membership holds the schema definition for the Membership entity.
type Membership struct {
	ent.Schema
}

// Fields of the Membership.
func (Membership) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("operation_date").
			Default(time.Now()),
		field.Time("start_date"),
		field.Time("end_date"),
		field.Float("amount_paid").
			Optional(),
		field.Float("amount_remaining"),
	}
}

// Edges of the Membership.
func (Membership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).
			Ref("membership").
			Unique().
			Required(),
	}
}
