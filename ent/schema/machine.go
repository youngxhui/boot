package schema

import "entgo.io/ent"

// Machine holds the schema definition for the Machine entity.
type Machine struct {
	ent.Schema
}

// Fields of the Machine.
func (Machine) Fields() []ent.Field {
	return nil
}

// Edges of the Machine.
func (Machine) Edges() []ent.Edge {
	return nil
}
