package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
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
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {

	return []ent.Edge{
		edge.From("owner", Role.Type).Ref("roles").Unique(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
