package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Token struct {
	ent.Schema
}

// Fields returns Balance fields.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		// contract:tokenId
		field.String("id").Unique().Immutable(),
		field.Uint64("tokenId"),
	}
}

// Edges returns Balance edges.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", Contract.Type).Unique(),
		edge.To("contract", Contract.Type).Unique(),
	}
}

// Annotations returns Transaction annotations.
func (Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
