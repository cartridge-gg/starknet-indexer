// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/cartridge-gg/starknet-indexer/ent/migrate"

	"github.com/cartridge-gg/starknet-indexer/ent/balance"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/ent/contract"
	"github.com/cartridge-gg/starknet-indexer/ent/event"
	"github.com/cartridge-gg/starknet-indexer/ent/token"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Balance is the client for interacting with the Balance builders.
	Balance *BalanceClient
	// Block is the client for interacting with the Block builders.
	Block *BlockClient
	// Contract is the client for interacting with the Contract builders.
	Contract *ContractClient
	// Event is the client for interacting with the Event builders.
	Event *EventClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// Transaction is the client for interacting with the Transaction builders.
	Transaction *TransactionClient
	// TransactionReceipt is the client for interacting with the TransactionReceipt builders.
	TransactionReceipt *TransactionReceiptClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Balance = NewBalanceClient(c.config)
	c.Block = NewBlockClient(c.config)
	c.Contract = NewContractClient(c.config)
	c.Event = NewEventClient(c.config)
	c.Token = NewTokenClient(c.config)
	c.Transaction = NewTransactionClient(c.config)
	c.TransactionReceipt = NewTransactionReceiptClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		Balance:            NewBalanceClient(cfg),
		Block:              NewBlockClient(cfg),
		Contract:           NewContractClient(cfg),
		Event:              NewEventClient(cfg),
		Token:              NewTokenClient(cfg),
		Transaction:        NewTransactionClient(cfg),
		TransactionReceipt: NewTransactionReceiptClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		Balance:            NewBalanceClient(cfg),
		Block:              NewBlockClient(cfg),
		Contract:           NewContractClient(cfg),
		Event:              NewEventClient(cfg),
		Token:              NewTokenClient(cfg),
		Transaction:        NewTransactionClient(cfg),
		TransactionReceipt: NewTransactionReceiptClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Balance.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Balance.Use(hooks...)
	c.Block.Use(hooks...)
	c.Contract.Use(hooks...)
	c.Event.Use(hooks...)
	c.Token.Use(hooks...)
	c.Transaction.Use(hooks...)
	c.TransactionReceipt.Use(hooks...)
}

// BalanceClient is a client for the Balance schema.
type BalanceClient struct {
	config
}

// NewBalanceClient returns a client for the Balance from the given config.
func NewBalanceClient(c config) *BalanceClient {
	return &BalanceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `balance.Hooks(f(g(h())))`.
func (c *BalanceClient) Use(hooks ...Hook) {
	c.hooks.Balance = append(c.hooks.Balance, hooks...)
}

// Create returns a builder for creating a Balance entity.
func (c *BalanceClient) Create() *BalanceCreate {
	mutation := newBalanceMutation(c.config, OpCreate)
	return &BalanceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Balance entities.
func (c *BalanceClient) CreateBulk(builders ...*BalanceCreate) *BalanceCreateBulk {
	return &BalanceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Balance.
func (c *BalanceClient) Update() *BalanceUpdate {
	mutation := newBalanceMutation(c.config, OpUpdate)
	return &BalanceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BalanceClient) UpdateOne(b *Balance) *BalanceUpdateOne {
	mutation := newBalanceMutation(c.config, OpUpdateOne, withBalance(b))
	return &BalanceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BalanceClient) UpdateOneID(id string) *BalanceUpdateOne {
	mutation := newBalanceMutation(c.config, OpUpdateOne, withBalanceID(id))
	return &BalanceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Balance.
func (c *BalanceClient) Delete() *BalanceDelete {
	mutation := newBalanceMutation(c.config, OpDelete)
	return &BalanceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BalanceClient) DeleteOne(b *Balance) *BalanceDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BalanceClient) DeleteOneID(id string) *BalanceDeleteOne {
	builder := c.Delete().Where(balance.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BalanceDeleteOne{builder}
}

// Query returns a query builder for Balance.
func (c *BalanceClient) Query() *BalanceQuery {
	return &BalanceQuery{
		config: c.config,
	}
}

// Get returns a Balance entity by its id.
func (c *BalanceClient) Get(ctx context.Context, id string) (*Balance, error) {
	return c.Query().Where(balance.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BalanceClient) GetX(ctx context.Context, id string) *Balance {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a Balance.
func (c *BalanceClient) QueryAccount(b *Balance) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(balance.Table, balance.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, balance.AccountTable, balance.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryContract queries the contract edge of a Balance.
func (c *BalanceClient) QueryContract(b *Balance) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(balance.Table, balance.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, balance.ContractTable, balance.ContractColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BalanceClient) Hooks() []Hook {
	return c.hooks.Balance
}

// BlockClient is a client for the Block schema.
type BlockClient struct {
	config
}

// NewBlockClient returns a client for the Block from the given config.
func NewBlockClient(c config) *BlockClient {
	return &BlockClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `block.Hooks(f(g(h())))`.
func (c *BlockClient) Use(hooks ...Hook) {
	c.hooks.Block = append(c.hooks.Block, hooks...)
}

// Create returns a builder for creating a Block entity.
func (c *BlockClient) Create() *BlockCreate {
	mutation := newBlockMutation(c.config, OpCreate)
	return &BlockCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Block entities.
func (c *BlockClient) CreateBulk(builders ...*BlockCreate) *BlockCreateBulk {
	return &BlockCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Block.
func (c *BlockClient) Update() *BlockUpdate {
	mutation := newBlockMutation(c.config, OpUpdate)
	return &BlockUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BlockClient) UpdateOne(b *Block) *BlockUpdateOne {
	mutation := newBlockMutation(c.config, OpUpdateOne, withBlock(b))
	return &BlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BlockClient) UpdateOneID(id string) *BlockUpdateOne {
	mutation := newBlockMutation(c.config, OpUpdateOne, withBlockID(id))
	return &BlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Block.
func (c *BlockClient) Delete() *BlockDelete {
	mutation := newBlockMutation(c.config, OpDelete)
	return &BlockDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BlockClient) DeleteOne(b *Block) *BlockDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BlockClient) DeleteOneID(id string) *BlockDeleteOne {
	builder := c.Delete().Where(block.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BlockDeleteOne{builder}
}

// Query returns a query builder for Block.
func (c *BlockClient) Query() *BlockQuery {
	return &BlockQuery{
		config: c.config,
	}
}

// Get returns a Block entity by its id.
func (c *BlockClient) Get(ctx context.Context, id string) (*Block, error) {
	return c.Query().Where(block.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BlockClient) GetX(ctx context.Context, id string) *Block {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTransactions queries the transactions edge of a Block.
func (c *BlockClient) QueryTransactions(b *Block) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(block.Table, block.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, block.TransactionsTable, block.TransactionsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTransactionReceipts queries the transaction_receipts edge of a Block.
func (c *BlockClient) QueryTransactionReceipts(b *Block) *TransactionReceiptQuery {
	query := &TransactionReceiptQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(block.Table, block.FieldID, id),
			sqlgraph.To(transactionreceipt.Table, transactionreceipt.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, block.TransactionReceiptsTable, block.TransactionReceiptsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BlockClient) Hooks() []Hook {
	return c.hooks.Block
}

// ContractClient is a client for the Contract schema.
type ContractClient struct {
	config
}

// NewContractClient returns a client for the Contract from the given config.
func NewContractClient(c config) *ContractClient {
	return &ContractClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contract.Hooks(f(g(h())))`.
func (c *ContractClient) Use(hooks ...Hook) {
	c.hooks.Contract = append(c.hooks.Contract, hooks...)
}

// Create returns a builder for creating a Contract entity.
func (c *ContractClient) Create() *ContractCreate {
	mutation := newContractMutation(c.config, OpCreate)
	return &ContractCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contract entities.
func (c *ContractClient) CreateBulk(builders ...*ContractCreate) *ContractCreateBulk {
	return &ContractCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contract.
func (c *ContractClient) Update() *ContractUpdate {
	mutation := newContractMutation(c.config, OpUpdate)
	return &ContractUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContractClient) UpdateOne(co *Contract) *ContractUpdateOne {
	mutation := newContractMutation(c.config, OpUpdateOne, withContract(co))
	return &ContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContractClient) UpdateOneID(id string) *ContractUpdateOne {
	mutation := newContractMutation(c.config, OpUpdateOne, withContractID(id))
	return &ContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contract.
func (c *ContractClient) Delete() *ContractDelete {
	mutation := newContractMutation(c.config, OpDelete)
	return &ContractDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContractClient) DeleteOne(co *Contract) *ContractDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ContractClient) DeleteOneID(id string) *ContractDeleteOne {
	builder := c.Delete().Where(contract.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContractDeleteOne{builder}
}

// Query returns a query builder for Contract.
func (c *ContractClient) Query() *ContractQuery {
	return &ContractQuery{
		config: c.config,
	}
}

// Get returns a Contract entity by its id.
func (c *ContractClient) Get(ctx context.Context, id string) (*Contract, error) {
	return c.Query().Where(contract.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContractClient) GetX(ctx context.Context, id string) *Contract {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTransactions queries the transactions edge of a Contract.
func (c *ContractClient) QueryTransactions(co *Contract) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(contract.Table, contract.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, contract.TransactionsTable, contract.TransactionsColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ContractClient) Hooks() []Hook {
	return c.hooks.Contract
}

// EventClient is a client for the Event schema.
type EventClient struct {
	config
}

// NewEventClient returns a client for the Event from the given config.
func NewEventClient(c config) *EventClient {
	return &EventClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `event.Hooks(f(g(h())))`.
func (c *EventClient) Use(hooks ...Hook) {
	c.hooks.Event = append(c.hooks.Event, hooks...)
}

// Create returns a builder for creating a Event entity.
func (c *EventClient) Create() *EventCreate {
	mutation := newEventMutation(c.config, OpCreate)
	return &EventCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Event entities.
func (c *EventClient) CreateBulk(builders ...*EventCreate) *EventCreateBulk {
	return &EventCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Event.
func (c *EventClient) Update() *EventUpdate {
	mutation := newEventMutation(c.config, OpUpdate)
	return &EventUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EventClient) UpdateOne(e *Event) *EventUpdateOne {
	mutation := newEventMutation(c.config, OpUpdateOne, withEvent(e))
	return &EventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EventClient) UpdateOneID(id string) *EventUpdateOne {
	mutation := newEventMutation(c.config, OpUpdateOne, withEventID(id))
	return &EventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Event.
func (c *EventClient) Delete() *EventDelete {
	mutation := newEventMutation(c.config, OpDelete)
	return &EventDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EventClient) DeleteOne(e *Event) *EventDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EventClient) DeleteOneID(id string) *EventDeleteOne {
	builder := c.Delete().Where(event.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EventDeleteOne{builder}
}

// Query returns a query builder for Event.
func (c *EventClient) Query() *EventQuery {
	return &EventQuery{
		config: c.config,
	}
}

// Get returns a Event entity by its id.
func (c *EventClient) Get(ctx context.Context, id string) (*Event, error) {
	return c.Query().Where(event.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EventClient) GetX(ctx context.Context, id string) *Event {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTransaction queries the transaction edge of a Event.
func (c *EventClient) QueryTransaction(e *Event) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, event.TransactionTable, event.TransactionColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EventClient) Hooks() []Hook {
	return c.hooks.Event
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Create returns a builder for creating a Token entity.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id string) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TokenClient) DeleteOneID(id string) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id string) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id string) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Token.
func (c *TokenClient) QueryOwner(t *Token) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, token.OwnerTable, token.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryContract queries the contract edge of a Token.
func (c *TokenClient) QueryContract(t *Token) *ContractQuery {
	query := &ContractQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, id),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, token.ContractTable, token.ContractColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	return c.hooks.Token
}

// TransactionClient is a client for the Transaction schema.
type TransactionClient struct {
	config
}

// NewTransactionClient returns a client for the Transaction from the given config.
func NewTransactionClient(c config) *TransactionClient {
	return &TransactionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transaction.Hooks(f(g(h())))`.
func (c *TransactionClient) Use(hooks ...Hook) {
	c.hooks.Transaction = append(c.hooks.Transaction, hooks...)
}

// Create returns a builder for creating a Transaction entity.
func (c *TransactionClient) Create() *TransactionCreate {
	mutation := newTransactionMutation(c.config, OpCreate)
	return &TransactionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transaction entities.
func (c *TransactionClient) CreateBulk(builders ...*TransactionCreate) *TransactionCreateBulk {
	return &TransactionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transaction.
func (c *TransactionClient) Update() *TransactionUpdate {
	mutation := newTransactionMutation(c.config, OpUpdate)
	return &TransactionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransactionClient) UpdateOne(t *Transaction) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransaction(t))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransactionClient) UpdateOneID(id string) *TransactionUpdateOne {
	mutation := newTransactionMutation(c.config, OpUpdateOne, withTransactionID(id))
	return &TransactionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transaction.
func (c *TransactionClient) Delete() *TransactionDelete {
	mutation := newTransactionMutation(c.config, OpDelete)
	return &TransactionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransactionClient) DeleteOne(t *Transaction) *TransactionDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TransactionClient) DeleteOneID(id string) *TransactionDeleteOne {
	builder := c.Delete().Where(transaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransactionDeleteOne{builder}
}

// Query returns a query builder for Transaction.
func (c *TransactionClient) Query() *TransactionQuery {
	return &TransactionQuery{
		config: c.config,
	}
}

// Get returns a Transaction entity by its id.
func (c *TransactionClient) Get(ctx context.Context, id string) (*Transaction, error) {
	return c.Query().Where(transaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransactionClient) GetX(ctx context.Context, id string) *Transaction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBlock queries the block edge of a Transaction.
func (c *TransactionClient) QueryBlock(t *Transaction) *BlockQuery {
	query := &BlockQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(block.Table, block.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transaction.BlockTable, transaction.BlockColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReceipt queries the receipt edge of a Transaction.
func (c *TransactionClient) QueryReceipt(t *Transaction) *TransactionReceiptQuery {
	query := &TransactionReceiptQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(transactionreceipt.Table, transactionreceipt.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, transaction.ReceiptTable, transaction.ReceiptColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEvents queries the events edge of a Transaction.
func (c *TransactionClient) QueryEvents(t *Transaction) *EventQuery {
	query := &EventQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transaction.Table, transaction.FieldID, id),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, transaction.EventsTable, transaction.EventsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TransactionClient) Hooks() []Hook {
	return c.hooks.Transaction
}

// TransactionReceiptClient is a client for the TransactionReceipt schema.
type TransactionReceiptClient struct {
	config
}

// NewTransactionReceiptClient returns a client for the TransactionReceipt from the given config.
func NewTransactionReceiptClient(c config) *TransactionReceiptClient {
	return &TransactionReceiptClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transactionreceipt.Hooks(f(g(h())))`.
func (c *TransactionReceiptClient) Use(hooks ...Hook) {
	c.hooks.TransactionReceipt = append(c.hooks.TransactionReceipt, hooks...)
}

// Create returns a builder for creating a TransactionReceipt entity.
func (c *TransactionReceiptClient) Create() *TransactionReceiptCreate {
	mutation := newTransactionReceiptMutation(c.config, OpCreate)
	return &TransactionReceiptCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TransactionReceipt entities.
func (c *TransactionReceiptClient) CreateBulk(builders ...*TransactionReceiptCreate) *TransactionReceiptCreateBulk {
	return &TransactionReceiptCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TransactionReceipt.
func (c *TransactionReceiptClient) Update() *TransactionReceiptUpdate {
	mutation := newTransactionReceiptMutation(c.config, OpUpdate)
	return &TransactionReceiptUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransactionReceiptClient) UpdateOne(tr *TransactionReceipt) *TransactionReceiptUpdateOne {
	mutation := newTransactionReceiptMutation(c.config, OpUpdateOne, withTransactionReceipt(tr))
	return &TransactionReceiptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransactionReceiptClient) UpdateOneID(id string) *TransactionReceiptUpdateOne {
	mutation := newTransactionReceiptMutation(c.config, OpUpdateOne, withTransactionReceiptID(id))
	return &TransactionReceiptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TransactionReceipt.
func (c *TransactionReceiptClient) Delete() *TransactionReceiptDelete {
	mutation := newTransactionReceiptMutation(c.config, OpDelete)
	return &TransactionReceiptDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransactionReceiptClient) DeleteOne(tr *TransactionReceipt) *TransactionReceiptDeleteOne {
	return c.DeleteOneID(tr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TransactionReceiptClient) DeleteOneID(id string) *TransactionReceiptDeleteOne {
	builder := c.Delete().Where(transactionreceipt.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransactionReceiptDeleteOne{builder}
}

// Query returns a query builder for TransactionReceipt.
func (c *TransactionReceiptClient) Query() *TransactionReceiptQuery {
	return &TransactionReceiptQuery{
		config: c.config,
	}
}

// Get returns a TransactionReceipt entity by its id.
func (c *TransactionReceiptClient) Get(ctx context.Context, id string) (*TransactionReceipt, error) {
	return c.Query().Where(transactionreceipt.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransactionReceiptClient) GetX(ctx context.Context, id string) *TransactionReceipt {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBlock queries the block edge of a TransactionReceipt.
func (c *TransactionReceiptClient) QueryBlock(tr *TransactionReceipt) *BlockQuery {
	query := &BlockQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transactionreceipt.Table, transactionreceipt.FieldID, id),
			sqlgraph.To(block.Table, block.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transactionreceipt.BlockTable, transactionreceipt.BlockColumn),
		)
		fromV = sqlgraph.Neighbors(tr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTransaction queries the transaction edge of a TransactionReceipt.
func (c *TransactionReceiptClient) QueryTransaction(tr *TransactionReceipt) *TransactionQuery {
	query := &TransactionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := tr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(transactionreceipt.Table, transactionreceipt.FieldID, id),
			sqlgraph.To(transaction.Table, transaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, transactionreceipt.TransactionTable, transactionreceipt.TransactionColumn),
		)
		fromV = sqlgraph.Neighbors(tr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TransactionReceiptClient) Hooks() []Hook {
	return c.hooks.TransactionReceipt
}
