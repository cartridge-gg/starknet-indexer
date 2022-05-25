package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/schema/big"
)

type Token struct {
	ent.Schema
}

// Fields returns Token fields.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		// contract:tokenId
		field.String("id").Unique().Immutable(),
		field.Int("tokenId").
			GoType(big.Int{}).
			SchemaType(big.IntSchemaType).
			DefaultFunc(func() big.Int {
				return big.NewInt(0)
			}).
			Annotations(
				entgql.Type("BigInt"),
			),
	}
}

// Edges returns Token edges.
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
