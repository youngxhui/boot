package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Notice holds the schema definition for the Notice entity.
type Notice struct {
	ent.Schema
}

// Fields of the Notice.
func (Notice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("noticeType"),
		field.Int("userId"),
		field.String("content"),
	}
}

// Edges of the Notice.
func (Notice) Edges() []ent.Edge {
	return nil
}

func (Notice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
