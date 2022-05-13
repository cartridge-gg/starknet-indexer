package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transaction defines the Transaction type schema.
type Transaction struct {
	ent.Schema
}

// Fields returns Transaction fields.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("contract_address"),
		field.String("entry_point_selector").Optional(),
		field.String("transaction_hash"),
		field.Strings("calldata"),
		field.Strings("signature").Optional(),
		field.String("nonce").Optional().
			Annotations(
				entgql.OrderField("NONCE"),
			),
	}
}

// Edges returns Transaction edges.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("block", Block.Type).Ref("transactions").
			Unique(),
		edge.From("contract", Contract.Type).Ref("transactions"),
		edge.To("receipts", TransactionReceipt.Type).Unique(),
		edge.To("events", Event.Type)}
}

// Annotations returns Transaction annotations.
func (Transaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
