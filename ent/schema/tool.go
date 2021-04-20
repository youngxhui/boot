package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tool holds the schema definition for the Tool entity.
type Tool struct {
	ent.Schema
}

// Fields of the Tool.
func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.Int("machine_id").Positive(),
		field.Int("status").Default(0),
	}
}

// Edges of the Tool.
func (Tool) Edges() []ent.Edge {
	return nil
}
