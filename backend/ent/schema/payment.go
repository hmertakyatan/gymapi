package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Enum("type").
			Values("in", "out"),
		field.String("description"),
		field.Float("amount"),
		field.Time("date_time").
			Default(time.Now()),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return nil
}
