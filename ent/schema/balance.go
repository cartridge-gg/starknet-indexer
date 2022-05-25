package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Balance struct {
	ent.Schema
}

// Fields returns Balance fields.
func (Balance) Fields() []ent.Field {
	return []ent.Field{
		// account:contract
		field.String("id").Unique().Immutable(),
		field.Uint64("balance").Default(0),
	}
}

// Edges returns Balance edges.
func (Balance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Contract.Type).Unique(),
		edge.To("contract", Contract.Type).Unique(),
	}
}

// Annotations returns Transaction annotations.
func (Balance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
