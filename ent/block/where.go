// Code generated by entc, DO NOT EDIT.

package block

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cartridge-gg/starknet-indexer/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BlockHash applies equality check predicate on the "block_hash" field. It's identical to BlockHashEQ.
func BlockHash(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockHash), v))
	})
}

// ParentBlockHash applies equality check predicate on the "parent_block_hash" field. It's identical to ParentBlockHashEQ.
func ParentBlockHash(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentBlockHash), v))
	})
}

// BlockNumber applies equality check predicate on the "block_number" field. It's identical to BlockNumberEQ.
func BlockNumber(v uint64) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockNumber), v))
	})
}

// StateRoot applies equality check predicate on the "state_root" field. It's identical to StateRootEQ.
func StateRoot(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStateRoot), v))
	})
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// BlockHashEQ applies the EQ predicate on the "block_hash" field.
func BlockHashEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockHash), v))
	})
}

// BlockHashNEQ applies the NEQ predicate on the "block_hash" field.
func BlockHashNEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBlockHash), v))
	})
}

// BlockHashIn applies the In predicate on the "block_hash" field.
func BlockHashIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBlockHash), v...))
	})
}

// BlockHashNotIn applies the NotIn predicate on the "block_hash" field.
func BlockHashNotIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBlockHash), v...))
	})
}

// BlockHashGT applies the GT predicate on the "block_hash" field.
func BlockHashGT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBlockHash), v))
	})
}

// BlockHashGTE applies the GTE predicate on the "block_hash" field.
func BlockHashGTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBlockHash), v))
	})
}

// BlockHashLT applies the LT predicate on the "block_hash" field.
func BlockHashLT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBlockHash), v))
	})
}

// BlockHashLTE applies the LTE predicate on the "block_hash" field.
func BlockHashLTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBlockHash), v))
	})
}

// BlockHashContains applies the Contains predicate on the "block_hash" field.
func BlockHashContains(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBlockHash), v))
	})
}

// BlockHashHasPrefix applies the HasPrefix predicate on the "block_hash" field.
func BlockHashHasPrefix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBlockHash), v))
	})
}

// BlockHashHasSuffix applies the HasSuffix predicate on the "block_hash" field.
func BlockHashHasSuffix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBlockHash), v))
	})
}

// BlockHashEqualFold applies the EqualFold predicate on the "block_hash" field.
func BlockHashEqualFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBlockHash), v))
	})
}

// BlockHashContainsFold applies the ContainsFold predicate on the "block_hash" field.
func BlockHashContainsFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBlockHash), v))
	})
}

// ParentBlockHashEQ applies the EQ predicate on the "parent_block_hash" field.
func ParentBlockHashEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashNEQ applies the NEQ predicate on the "parent_block_hash" field.
func ParentBlockHashNEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashIn applies the In predicate on the "parent_block_hash" field.
func ParentBlockHashIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldParentBlockHash), v...))
	})
}

// ParentBlockHashNotIn applies the NotIn predicate on the "parent_block_hash" field.
func ParentBlockHashNotIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldParentBlockHash), v...))
	})
}

// ParentBlockHashGT applies the GT predicate on the "parent_block_hash" field.
func ParentBlockHashGT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashGTE applies the GTE predicate on the "parent_block_hash" field.
func ParentBlockHashGTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashLT applies the LT predicate on the "parent_block_hash" field.
func ParentBlockHashLT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashLTE applies the LTE predicate on the "parent_block_hash" field.
func ParentBlockHashLTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashContains applies the Contains predicate on the "parent_block_hash" field.
func ParentBlockHashContains(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashHasPrefix applies the HasPrefix predicate on the "parent_block_hash" field.
func ParentBlockHashHasPrefix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashHasSuffix applies the HasSuffix predicate on the "parent_block_hash" field.
func ParentBlockHashHasSuffix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashEqualFold applies the EqualFold predicate on the "parent_block_hash" field.
func ParentBlockHashEqualFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldParentBlockHash), v))
	})
}

// ParentBlockHashContainsFold applies the ContainsFold predicate on the "parent_block_hash" field.
func ParentBlockHashContainsFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldParentBlockHash), v))
	})
}

// BlockNumberEQ applies the EQ predicate on the "block_number" field.
func BlockNumberEQ(v uint64) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberNEQ applies the NEQ predicate on the "block_number" field.
func BlockNumberNEQ(v uint64) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberIn applies the In predicate on the "block_number" field.
func BlockNumberIn(vs ...uint64) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBlockNumber), v...))
	})
}

// BlockNumberNotIn applies the NotIn predicate on the "block_number" field.
func BlockNumberNotIn(vs ...uint64) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBlockNumber), v...))
	})
}

// BlockNumberGT applies the GT predicate on the "block_number" field.
func BlockNumberGT(v uint64) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberGTE applies the GTE predicate on the "block_number" field.
func BlockNumberGTE(v uint64) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberLT applies the LT predicate on the "block_number" field.
func BlockNumberLT(v uint64) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberLTE applies the LTE predicate on the "block_number" field.
func BlockNumberLTE(v uint64) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBlockNumber), v))
	})
}

// StateRootEQ applies the EQ predicate on the "state_root" field.
func StateRootEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStateRoot), v))
	})
}

// StateRootNEQ applies the NEQ predicate on the "state_root" field.
func StateRootNEQ(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStateRoot), v))
	})
}

// StateRootIn applies the In predicate on the "state_root" field.
func StateRootIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStateRoot), v...))
	})
}

// StateRootNotIn applies the NotIn predicate on the "state_root" field.
func StateRootNotIn(vs ...string) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStateRoot), v...))
	})
}

// StateRootGT applies the GT predicate on the "state_root" field.
func StateRootGT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStateRoot), v))
	})
}

// StateRootGTE applies the GTE predicate on the "state_root" field.
func StateRootGTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStateRoot), v))
	})
}

// StateRootLT applies the LT predicate on the "state_root" field.
func StateRootLT(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStateRoot), v))
	})
}

// StateRootLTE applies the LTE predicate on the "state_root" field.
func StateRootLTE(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStateRoot), v))
	})
}

// StateRootContains applies the Contains predicate on the "state_root" field.
func StateRootContains(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStateRoot), v))
	})
}

// StateRootHasPrefix applies the HasPrefix predicate on the "state_root" field.
func StateRootHasPrefix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStateRoot), v))
	})
}

// StateRootHasSuffix applies the HasSuffix predicate on the "state_root" field.
func StateRootHasSuffix(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStateRoot), v))
	})
}

// StateRootEqualFold applies the EqualFold predicate on the "state_root" field.
func StateRootEqualFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStateRoot), v))
	})
}

// StateRootContainsFold applies the ContainsFold predicate on the "state_root" field.
func StateRootContainsFold(v string) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStateRoot), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTimestamp), v))
	})
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTimestamp), v))
	})
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTimestamp), v...))
	})
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Block {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Block(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTimestamp), v...))
	})
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTimestamp), v))
	})
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTimestamp), v))
	})
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTimestamp), v))
	})
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTimestamp), v))
	})
}

// HasTransactions applies the HasEdge predicate on the "transactions" edge.
func HasTransactions() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionsWith applies the HasEdge predicate on the "transactions" edge with a given conditions (other predicates).
func HasTransactionsWith(preds ...predicate.Transaction) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionsTable, TransactionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTransactionReceipts applies the HasEdge predicate on the "transaction_receipts" edge.
func HasTransactionReceipts() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionReceiptsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionReceiptsTable, TransactionReceiptsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionReceiptsWith applies the HasEdge predicate on the "transaction_receipts" edge with a given conditions (other predicates).
func HasTransactionReceiptsWith(preds ...predicate.TransactionReceipt) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionReceiptsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionReceiptsTable, TransactionReceiptsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Block) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Block) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Block) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		p(s.Not())
	})
}
