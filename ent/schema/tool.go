package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
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
		field.Time("create_time").Default(time.Now().Local),
		field.Time("update_time").Default(time.Now().Local),
	}
}

// Edges of the Tool.
func (Tool) Edges() []ent.Edge {
	return nil
}
