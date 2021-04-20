package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("password"),
		field.String("username").Default("unknown"),
		field.Time("create_time").Default(time.Now().Local),
		field.Time("update_time").Default(time.Now().Local),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
