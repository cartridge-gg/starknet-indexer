package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/schema/big"
)

type Balance struct {
	ent.Schema
}

// Fields returns Balance fields.
func (Balance) Fields() []ent.Field {
	return []ent.Field{
		// account:contract
		field.String("id").Unique().Immutable(),
		field.Int("tokenId").
			Optional().
			GoType(big.Int{}).
			SchemaType(big.IntSchemaType).
			DefaultFunc(func() big.Int {
				return big.NewInt(0)
			}).
			Annotations(
				entgql.Type("BigInt"),
			),
		field.Int("balance").
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

// Edges returns Balance edges.
func (Balance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Contract.Type).
			Unique(),
		edge.To("contract", Contract.Type).
			Unique(),
	}
}

// Annotations returns Transaction annotations.
func (Balance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
