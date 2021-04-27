package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Machine holds the schema definition for the Machine entity.
type Machine struct {
	ent.Schema
}

// Fields of the Machine.
func (Machine) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Machine.
func (Machine) Edges() []ent.Edge {
	return nil
}

func (Machine) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
