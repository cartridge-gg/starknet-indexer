// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/ent/event"
	"github.com/cartridge-gg/starknet-indexer/ent/predicate"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt"
)

// TransactionQuery is the builder for querying Transaction entities.
type TransactionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Transaction
	// eager-loading edges.
	withBlock   *BlockQuery
	withReceipt *TransactionReceiptQuery
	withEvents  *EventQuery
	withFKs     bool
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*Transaction) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransactionQuery builder.
func (tq *TransactionQuery) Where(ps ...predicate.Transaction) *TransactionQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TransactionQuery) Limit(limit int) *TransactionQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TransactionQuery) Offset(offset int) *TransactionQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TransactionQuery) Unique(unique bool) *TransactionQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TransactionQuery) Order(o ...OrderFunc) *TransactionQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryBlock chains the current query on the "block" edge.
func (tq *TransactionQuery) QueryBlock() *BlockQuery {
	query := &BlockQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, selector),
			sqlgraph.To(block.Table, block.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transaction.BlockTable, transaction.BlockColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReceipt chains the current query on the "receipt" edge.
func (tq *TransactionQuery) QueryReceipt() *TransactionReceiptQuery {
	query := &TransactionReceiptQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, selector),
			sqlgraph.To(transactionreceipt.Table, transactionreceipt.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, transaction.ReceiptTable, transaction.ReceiptColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvents chains the current query on the "events" edge.
func (tq *TransactionQuery) QueryEvents() *EventQuery {
	query := &EventQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transaction.EventsTable, transaction.EventsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Transaction entity from the query.
// Returns a *NotFoundError when no Transaction was found.
func (tq *TransactionQuery) First(ctx context.Context) (*Transaction, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transaction.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TransactionQuery) FirstX(ctx context.Context) *Transaction {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Transaction ID from the query.
// Returns a *NotFoundError when no Transaction ID was found.
func (tq *TransactionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transaction.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TransactionQuery) FirstIDX(ctx context.Context) string {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Transaction entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Transaction entity is found.
// Returns a *NotFoundError when no Transaction entities are found.
func (tq *TransactionQuery) Only(ctx context.Context) (*Transaction, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transaction.Label}
	default:
		return nil, &NotSingularError{transaction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TransactionQuery) OnlyX(ctx context.Context) *Transaction {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Transaction ID in the query.
// Returns a *NotSingularError when more than one Transaction ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TransactionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transaction.Label}
	default:
		err = &NotSingularError{transaction.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TransactionQuery) OnlyIDX(ctx context.Context) string {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Transactions.
func (tq *TransactionQuery) All(ctx context.Context) ([]*Transaction, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TransactionQuery) AllX(ctx context.Context) []*Transaction {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Transaction IDs.
func (tq *TransactionQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := tq.Select(transaction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TransactionQuery) IDsX(ctx context.Context) []string {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TransactionQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TransactionQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TransactionQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TransactionQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransactionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TransactionQuery) Clone() *TransactionQuery {
	if tq == nil {
		return nil
	}
	return &TransactionQuery{
		config:      tq.config,
		limit:       tq.limit,
		offset:      tq.offset,
		order:       append([]OrderFunc{}, tq.order...),
		predicates:  append([]predicate.Transaction{}, tq.predicates...),
		withBlock:   tq.withBlock.Clone(),
		withReceipt: tq.withReceipt.Clone(),
		withEvents:  tq.withEvents.Clone(),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// WithBlock tells the query-builder to eager-load the nodes that are connected to
// the "block" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TransactionQuery) WithBlock(opts ...func(*BlockQuery)) *TransactionQuery {
	query := &BlockQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withBlock = query
	return tq
}

// WithReceipt tells the query-builder to eager-load the nodes that are connected to
// the "receipt" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TransactionQuery) WithReceipt(opts ...func(*TransactionReceiptQuery)) *TransactionQuery {
	query := &TransactionReceiptQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withReceipt = query
	return tq
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TransactionQuery) WithEvents(opts ...func(*EventQuery)) *TransactionQuery {
	query := &EventQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withEvents = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ContractAddress string `json:"contract_address,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Transaction.Query().
//		GroupBy(transaction.FieldContractAddress).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TransactionQuery) GroupBy(field string, fields ...string) *TransactionGroupBy {
	grbuild := &TransactionGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = transaction.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ContractAddress string `json:"contract_address,omitempty"`
//	}
//
//	client.Transaction.Query().
//		Select(transaction.FieldContractAddress).
//		Scan(ctx, &v)
//
func (tq *TransactionQuery) Select(fields ...string) *TransactionSelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &TransactionSelect{TransactionQuery: tq}
	selbuild.label = transaction.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
}

func (tq *TransactionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !transaction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TransactionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Transaction, error) {
	var (
		nodes       = []*Transaction{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [3]bool{
			tq.withBlock != nil,
			tq.withReceipt != nil,
			tq.withEvents != nil,
		}
	)
	if tq.withBlock != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Transaction).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Transaction{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withBlock; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Transaction)
		for i := range nodes {
			if nodes[i].block_transactions == nil {
				continue
			}
			fk := *nodes[i].block_transactions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(block.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "block_transactions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Block = n
			}
		}
	}

	if query := tq.withReceipt; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*Transaction)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.TransactionReceipt(func(s *sql.Selector) {
			s.Where(sql.InValues(transaction.ReceiptColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.transaction_receipt
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "transaction_receipt" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "transaction_receipt" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Receipt = n
		}
	}

	if query := tq.withEvents; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*Transaction)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Events = []*Event{}
		}
		query.withFKs = true
		query.Where(predicate.Event(func(s *sql.Selector) {
			s.Where(sql.InValues(transaction.EventsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.transaction_events
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "transaction_events" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "transaction_events" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Events = append(node.Edges.Events, n)
		}
	}

	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TransactionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TransactionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *TransactionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transaction.Table,
			Columns: transaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transaction.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for i := range fields {
			if fields[i] != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TransactionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(transaction.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = transaction.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TransactionGroupBy is the group-by builder for Transaction entities.
type TransactionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TransactionGroupBy) Aggregate(fns ...AggregateFunc) *TransactionGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TransactionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

func (tgb *TransactionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !transaction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TransactionGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TransactionSelect is the builder for selecting fields of Transaction entities.
type TransactionSelect struct {
	*TransactionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TransactionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TransactionQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

func (ts *TransactionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
