package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dontpanicdao/caigo/types"
)

// TransactionReceipt defines the TransactionReceipt type schema.
type TransactionReceipt struct {
	ent.Schema
}

// Fields returns TransactionReceipt fields.
func (TransactionReceipt) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("transaction_hash"),
		field.Enum("status").Values("UNKNOWN",
			"RECEIVED",
			"PENDING",
			"ACCEPTED_ON_L2",
			"ACCEPTED_ON_L1",
			"REJECTED"),
		field.String("status_data"),
		field.JSON("messages_sent", []types.L1Message{}).Annotations(
			entgql.Skip(),
		),
		field.JSON("l1_origin_message", types.L2Message{}),
		field.JSON("events", []types.Event{}).Annotations(
			entgql.Skip(),
		),
	}
}

// Edges returns TransactionReceipt edges.
func (TransactionReceipt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("block", Block.Type).Ref("transaction_receipts").
			Unique(),
	}
}

// Annotations returns TransactionReceipt annotations.
func (TransactionReceipt) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
