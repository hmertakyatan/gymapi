package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Personnel holds the schema definition for the Personnel entity.
type Personnel struct {
	ent.Schema
}

// Fields of the Personnel.
func (Personnel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.Float("salary"),
		field.String("description"),
		field.Time("bith_date"),
		field.Time("start_date").
			Default(time.Now()),
		field.Bool("status").
			Default(true),
	}
}

// Edges of the Personnel.
func (Personnel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("personnel_rel_customer", PersonnelRelCustomer.Type),
	}
}
