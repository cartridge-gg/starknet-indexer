// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlocksColumns holds the columns for the "blocks" table.
	BlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "block_hash", Type: field.TypeString, Unique: true},
		{Name: "parent_block_hash", Type: field.TypeString},
		{Name: "block_number", Type: field.TypeUint64, Unique: true},
		{Name: "state_root", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ACCEPTED_ON_L1", "ACCEPTED_ON_L2"}},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// BlocksTable holds the schema information for the "blocks" table.
	BlocksTable = &schema.Table{
		Name:       "blocks",
		Columns:    BlocksColumns,
		PrimaryKey: []*schema.Column{BlocksColumns[0]},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "contract_address", Type: field.TypeString},
		{Name: "entry_point_selector", Type: field.TypeString},
		{Name: "entry_point_type", Type: field.TypeString},
		{Name: "transaction_hash", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"INVOKE_FUNCTION", "DEPLOY"}},
		{Name: "nonce", Type: field.TypeString},
		{Name: "block_transactions", Type: field.TypeString, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_blocks_transactions",
				Columns:    []*schema.Column{TransactionsColumns[7]},
				RefColumns: []*schema.Column{BlocksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlocksTable,
		TransactionsTable,
	}
)

func init() {
	TransactionsTable.ForeignKeys[0].RefTable = BlocksTable
}
