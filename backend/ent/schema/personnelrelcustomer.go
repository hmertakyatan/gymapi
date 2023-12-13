package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PersonelRelCustomer holds the schema definition for the PersonelRelCustomer entity.
type PersonnelRelCustomer struct {
	ent.Schema
}

// Fields of the PersonelRelCustomer.
func (PersonnelRelCustomer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("description"),
		field.Float("price"),
	}
}

// Edges of the PersonelRelCustomer.
func (PersonnelRelCustomer) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("personnel", Personnel.Type).
			Ref("personnel_rel_customer").
			Unique().
			Required(),
		edge.From("customer", Customer.Type).
			Ref("personnel_rel_customer").
			Unique().
			Required(),
	}
}
