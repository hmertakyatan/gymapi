package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User is the ent schema for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("username").
			Unique(),
		field.String("password"),
		field.String("role"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now()),
		field.Bool("status").
			Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
