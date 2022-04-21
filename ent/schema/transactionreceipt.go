package schema

import (
	"encoding/json"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type L1ToL2ConsumedMessage struct {
	FromAddress string   `json:"from_address"`
	ToAddress   string   `json:"to_address"`
	Selector    string   `json:"selector"`
	Payload     []string `json:"payload"`
}

type BuiltinInstanceCounter struct {
	PedersenBuiltin   uint64 `json:"pedersen_builtin"`
	RangeCheckBuiltin uint64 `json:"range_check_builtin"`
	BitwiseBuiltin    uint64 `json:"bitwise_builtin"`
	OutputBuiltin     uint64 `json:"output_builtin"`
	EcdsaBuiltin      uint64 `json:"ecdsa_builtin"`
	EcOpBuiltin       uint64 `json:"ec_op_builtin"`
}

type ExecutionResources struct {
	NSteps                 uint64                 `json:"n_steps"`
	BuiltinInstanceCounter BuiltinInstanceCounter `json:"builtin_instance_counter"`
	NMemoryHoles           uint64                 `json:"n_memory_holes"`
}

type Event struct {
	json.RawMessage
}

// TransactionReceipt defines the TransactionReceipt type schema.
type TransactionReceipt struct {
	ent.Schema
}

// Fields returns TransactionReceipt fields.
func (TransactionReceipt) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.Int32("transaction_index").
			Annotations(
				entgql.OrderField("INDEX"),
			),
		field.String("transaction_hash"),
		field.JSON("l1_to_l2_consumed_message", L1ToL2ConsumedMessage{}).Annotations(
			entgql.Type("L1ToL2ConsumedMessage"),
		),
		field.JSON("execution_resources", ExecutionResources{}).Annotations(
			entgql.Type("ExecutionResources"),
		),
		field.JSON("events", json.RawMessage{}).Annotations(
			entgql.Type("JSON"),
		),
		field.JSON("l2_to_l1_messages", json.RawMessage{}).Annotations(
			entgql.Type("JSON"),
		),
	}
}

// Edges returns TransactionReceipt edges.
func (TransactionReceipt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("block", Block.Type).Ref("transaction_receipts").
			Annotations(entgql.Unbind()).
			Unique(),
		edge.From("transaction", Transaction.Type).Ref("receipts").
			Annotations(entgql.Unbind()).
			Unique(),
	}
}

// Annotations returns TransactionReceipt annotations.
func (TransactionReceipt) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
