// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tarrencev/starknet-indexer/ent/balance"
	"github.com/tarrencev/starknet-indexer/ent/contract"
	"github.com/tarrencev/starknet-indexer/ent/predicate"
)

// BalanceQuery is the builder for querying Balance entities.
type BalanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Balance
	// eager-loading edges.
	withAccount  *ContractQuery
	withContract *ContractQuery
	withFKs      bool
	modifiers    []func(*sql.Selector)
	loadTotal    []func(context.Context, []*Balance) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BalanceQuery builder.
func (bq *BalanceQuery) Where(ps ...predicate.Balance) *BalanceQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BalanceQuery) Limit(limit int) *BalanceQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BalanceQuery) Offset(offset int) *BalanceQuery {
	bq.offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BalanceQuery) Unique(unique bool) *BalanceQuery {
	bq.unique = &unique
	return bq
}

// Order adds an order step to the query.
func (bq *BalanceQuery) Order(o ...OrderFunc) *BalanceQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryAccount chains the current query on the "account" edge.
func (bq *BalanceQuery) QueryAccount() *ContractQuery {
	query := &ContractQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(balance.Table, balance.FieldID, selector),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, balance.AccountTable, balance.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryContract chains the current query on the "contract" edge.
func (bq *BalanceQuery) QueryContract() *ContractQuery {
	query := &ContractQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(balance.Table, balance.FieldID, selector),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, balance.ContractTable, balance.ContractColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Balance entity from the query.
// Returns a *NotFoundError when no Balance was found.
func (bq *BalanceQuery) First(ctx context.Context) (*Balance, error) {
	nodes, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{balance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BalanceQuery) FirstX(ctx context.Context) *Balance {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Balance ID from the query.
// Returns a *NotFoundError when no Balance ID was found.
func (bq *BalanceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{balance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BalanceQuery) FirstIDX(ctx context.Context) string {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Balance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Balance entity is found.
// Returns a *NotFoundError when no Balance entities are found.
func (bq *BalanceQuery) Only(ctx context.Context) (*Balance, error) {
	nodes, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{balance.Label}
	default:
		return nil, &NotSingularError{balance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BalanceQuery) OnlyX(ctx context.Context) *Balance {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Balance ID in the query.
// Returns a *NotSingularError when more than one Balance ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BalanceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{balance.Label}
	default:
		err = &NotSingularError{balance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BalanceQuery) OnlyIDX(ctx context.Context) string {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Balances.
func (bq *BalanceQuery) All(ctx context.Context) ([]*Balance, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BalanceQuery) AllX(ctx context.Context) []*Balance {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Balance IDs.
func (bq *BalanceQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := bq.Select(balance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BalanceQuery) IDsX(ctx context.Context) []string {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BalanceQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BalanceQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BalanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BalanceQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BalanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BalanceQuery) Clone() *BalanceQuery {
	if bq == nil {
		return nil
	}
	return &BalanceQuery{
		config:       bq.config,
		limit:        bq.limit,
		offset:       bq.offset,
		order:        append([]OrderFunc{}, bq.order...),
		predicates:   append([]predicate.Balance{}, bq.predicates...),
		withAccount:  bq.withAccount.Clone(),
		withContract: bq.withContract.Clone(),
		// clone intermediate query.
		sql:    bq.sql.Clone(),
		path:   bq.path,
		unique: bq.unique,
	}
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BalanceQuery) WithAccount(opts ...func(*ContractQuery)) *BalanceQuery {
	query := &ContractQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withAccount = query
	return bq
}

// WithContract tells the query-builder to eager-load the nodes that are connected to
// the "contract" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BalanceQuery) WithContract(opts ...func(*ContractQuery)) *BalanceQuery {
	query := &ContractQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withContract = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Balance uint64 `json:"balance,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Balance.Query().
//		GroupBy(balance.FieldBalance).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bq *BalanceQuery) GroupBy(field string, fields ...string) *BalanceGroupBy {
	grbuild := &BalanceGroupBy{config: bq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(ctx), nil
	}
	grbuild.label = balance.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Balance uint64 `json:"balance,omitempty"`
//	}
//
//	client.Balance.Query().
//		Select(balance.FieldBalance).
//		Scan(ctx, &v)
//
func (bq *BalanceQuery) Select(fields ...string) *BalanceSelect {
	bq.fields = append(bq.fields, fields...)
	selbuild := &BalanceSelect{BalanceQuery: bq}
	selbuild.label = balance.Label
	selbuild.flds, selbuild.scan = &bq.fields, selbuild.Scan
	return selbuild
}

func (bq *BalanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bq.fields {
		if !balance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BalanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Balance, error) {
	var (
		nodes       = []*Balance{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [2]bool{
			bq.withAccount != nil,
			bq.withContract != nil,
		}
	)
	if bq.withAccount != nil || bq.withContract != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, balance.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Balance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Balance{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bq.withAccount; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Balance)
		for i := range nodes {
			if nodes[i].balance_account == nil {
				continue
			}
			fk := *nodes[i].balance_account
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(contract.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "balance_account" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Account = n
			}
		}
	}

	if query := bq.withContract; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Balance)
		for i := range nodes {
			if nodes[i].balance_contract == nil {
				continue
			}
			fk := *nodes[i].balance_contract
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(contract.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "balance_contract" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Contract = n
			}
		}
	}

	for i := range bq.loadTotal {
		if err := bq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BalanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	_spec.Node.Columns = bq.fields
	if len(bq.fields) > 0 {
		_spec.Unique = bq.unique != nil && *bq.unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BalanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bq *BalanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   balance.Table,
			Columns: balance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: balance.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if unique := bq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, balance.FieldID)
		for i := range fields {
			if fields[i] != balance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BalanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(balance.Table)
	columns := bq.fields
	if len(columns) == 0 {
		columns = balance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.unique != nil && *bq.unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BalanceGroupBy is the group-by builder for Balance entities.
type BalanceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BalanceGroupBy) Aggregate(fns ...AggregateFunc) *BalanceGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bgb *BalanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

func (bgb *BalanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bgb.fields {
		if !balance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BalanceGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql.Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
		for _, f := range bgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bgb.fields...)...)
}

// BalanceSelect is the builder for selecting fields of Balance entities.
type BalanceSelect struct {
	*BalanceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BalanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	bs.sql = bs.BalanceQuery.sqlQuery(ctx)
	return bs.sqlScan(ctx, v)
}

func (bs *BalanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bs.sql.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
