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
		field.String("entry_point_selector"),
		field.String("entry_point_type"),
		field.String("transaction_hash"),
		field.Strings("calldata"),
		field.Strings("signature"),
		field.Enum("type").Values("INVOKE_FUNCTION", "DEPLOY"),
		field.String("nonce").
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
		edge.To("receipts", TransactionReceipt.Type).
			Unique(),
	}
}

// Annotations returns Transaction annotations.
func (Transaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
