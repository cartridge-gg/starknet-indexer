package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Block defines the Block type schema.
type Block struct {
	ent.Schema
}

// Fields returns Block fields.
func (Block) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("block_hash").Unique(),
		field.String("parent_block_hash"),
		field.Uint64("block_number").Unique().Annotations(
			entgql.Type("Long"),
			entgql.OrderField("BLOCK_NUMBER"),
		),
		field.String("state_root"),
		field.Enum("status").Values("ACCEPTED_ON_L1", "ACCEPTED_ON_L2").
			Annotations(
				entgql.Type("BlockStatus"),
			),
		field.Time("timestamp").
			Immutable().
			Annotations(
				entgql.OrderField("TIMESTAMP"),
			),
	}
}

// Edges returns Block edges.
func (Block) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type).
			Annotations(entgql.Unbind()),
	}
}

// Annotations returns Block annotations.
func (Block) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
