package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Membership_type holds the schema definition for the Membership_type entity.
type Membership_type struct {
	ent.Schema
}

// Fields of the Membership_type.
func (Membership_type) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.String("description"),
		field.Uint8("membership_month"),
		field.Float("price"),
	}
}

// Edges of the Membership_type.
func (Membership_type) Edges() []ent.Edge {
	return nil
}
