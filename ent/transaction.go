// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// ContractAddress holds the value of the "contract_address" field.
	ContractAddress string `json:"contract_address,omitempty"`
	// EntryPointSelector holds the value of the "entry_point_selector" field.
	EntryPointSelector string `json:"entry_point_selector,omitempty"`
	// TransactionHash holds the value of the "transaction_hash" field.
	TransactionHash string `json:"transaction_hash,omitempty"`
	// Calldata holds the value of the "calldata" field.
	Calldata []string `json:"calldata,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature []string `json:"signature,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce string `json:"nonce,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges                 TransactionEdges `json:"edges"`
	block_transactions    *string
	contract_transactions *string
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// Block holds the value of the block edge.
	Block *Block `json:"block,omitempty"`
	// Receipt holds the value of the receipt edge.
	Receipt *TransactionReceipt `json:"receipt,omitempty"`
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]*int
}

// BlockOrErr returns the Block value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) BlockOrErr() (*Block, error) {
	if e.loadedTypes[0] {
		if e.Block == nil {
			// The edge block was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: block.Label}
		}
		return e.Block, nil
	}
	return nil, &NotLoadedError{edge: "block"}
}

// ReceiptOrErr returns the Receipt value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) ReceiptOrErr() (*TransactionReceipt, error) {
	if e.loadedTypes[1] {
		if e.Receipt == nil {
			// The edge receipt was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: transactionreceipt.Label}
		}
		return e.Receipt, nil
	}
	return nil, &NotLoadedError{edge: "receipt"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[2] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldCalldata, transaction.FieldSignature:
			values[i] = new([]byte)
		case transaction.FieldID, transaction.FieldContractAddress, transaction.FieldEntryPointSelector, transaction.FieldTransactionHash, transaction.FieldNonce:
			values[i] = new(sql.NullString)
		case transaction.ForeignKeys[0]: // block_transactions
			values[i] = new(sql.NullString)
		case transaction.ForeignKeys[1]: // contract_transactions
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case transaction.FieldContractAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contract_address", values[i])
			} else if value.Valid {
				t.ContractAddress = value.String
			}
		case transaction.FieldEntryPointSelector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entry_point_selector", values[i])
			} else if value.Valid {
				t.EntryPointSelector = value.String
			}
		case transaction.FieldTransactionHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_hash", values[i])
			} else if value.Valid {
				t.TransactionHash = value.String
			}
		case transaction.FieldCalldata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field calldata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Calldata); err != nil {
					return fmt.Errorf("unmarshal field calldata: %w", err)
				}
			}
		case transaction.FieldSignature:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field signature", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Signature); err != nil {
					return fmt.Errorf("unmarshal field signature: %w", err)
				}
			}
		case transaction.FieldNonce:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				t.Nonce = value.String
			}
		case transaction.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field block_transactions", values[i])
			} else if value.Valid {
				t.block_transactions = new(string)
				*t.block_transactions = value.String
			}
		case transaction.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contract_transactions", values[i])
			} else if value.Valid {
				t.contract_transactions = new(string)
				*t.contract_transactions = value.String
			}
		}
	}
	return nil
}

// QueryBlock queries the "block" edge of the Transaction entity.
func (t *Transaction) QueryBlock() *BlockQuery {
	return (&TransactionClient{config: t.config}).QueryBlock(t)
}

// QueryReceipt queries the "receipt" edge of the Transaction entity.
func (t *Transaction) QueryReceipt() *TransactionReceiptQuery {
	return (&TransactionClient{config: t.config}).QueryReceipt(t)
}

// QueryEvents queries the "events" edge of the Transaction entity.
func (t *Transaction) QueryEvents() *EventQuery {
	return (&TransactionClient{config: t.config}).QueryEvents(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return (&TransactionClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("contract_address=")
	builder.WriteString(t.ContractAddress)
	builder.WriteString(", ")
	builder.WriteString("entry_point_selector=")
	builder.WriteString(t.EntryPointSelector)
	builder.WriteString(", ")
	builder.WriteString("transaction_hash=")
	builder.WriteString(t.TransactionHash)
	builder.WriteString(", ")
	builder.WriteString("calldata=")
	builder.WriteString(fmt.Sprintf("%v", t.Calldata))
	builder.WriteString(", ")
	builder.WriteString("signature=")
	builder.WriteString(fmt.Sprintf("%v", t.Signature))
	builder.WriteString(", ")
	builder.WriteString("nonce=")
	builder.WriteString(t.Nonce)
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction

func (t Transactions) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
