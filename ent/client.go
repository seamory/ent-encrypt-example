// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"main/ent/migrate"

	"main/ent/accountinfo"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AccountInfo is the client for interacting with the AccountInfo builders.
	AccountInfo *AccountInfoClient
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
	c.AccountInfo = NewAccountInfoClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		AccountInfo: NewAccountInfoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:         ctx,
		config:      cfg,
		AccountInfo: NewAccountInfoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AccountInfo.
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
	c.AccountInfo.Use(hooks...)
}

// AccountInfoClient is a client for the AccountInfo schema.
type AccountInfoClient struct {
	config
}

// NewAccountInfoClient returns a client for the AccountInfo from the given config.
func NewAccountInfoClient(c config) *AccountInfoClient {
	return &AccountInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `accountinfo.Hooks(f(g(h())))`.
func (c *AccountInfoClient) Use(hooks ...Hook) {
	c.hooks.AccountInfo = append(c.hooks.AccountInfo, hooks...)
}

// Create returns a builder for creating a AccountInfo entity.
func (c *AccountInfoClient) Create() *AccountInfoCreate {
	mutation := newAccountInfoMutation(c.config, OpCreate)
	return &AccountInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AccountInfo entities.
func (c *AccountInfoClient) CreateBulk(builders ...*AccountInfoCreate) *AccountInfoCreateBulk {
	return &AccountInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AccountInfo.
func (c *AccountInfoClient) Update() *AccountInfoUpdate {
	mutation := newAccountInfoMutation(c.config, OpUpdate)
	return &AccountInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountInfoClient) UpdateOne(ai *AccountInfo) *AccountInfoUpdateOne {
	mutation := newAccountInfoMutation(c.config, OpUpdateOne, withAccountInfo(ai))
	return &AccountInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountInfoClient) UpdateOneID(id int) *AccountInfoUpdateOne {
	mutation := newAccountInfoMutation(c.config, OpUpdateOne, withAccountInfoID(id))
	return &AccountInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AccountInfo.
func (c *AccountInfoClient) Delete() *AccountInfoDelete {
	mutation := newAccountInfoMutation(c.config, OpDelete)
	return &AccountInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountInfoClient) DeleteOne(ai *AccountInfo) *AccountInfoDeleteOne {
	return c.DeleteOneID(ai.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AccountInfoClient) DeleteOneID(id int) *AccountInfoDeleteOne {
	builder := c.Delete().Where(accountinfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountInfoDeleteOne{builder}
}

// Query returns a query builder for AccountInfo.
func (c *AccountInfoClient) Query() *AccountInfoQuery {
	return &AccountInfoQuery{
		config: c.config,
	}
}

// Get returns a AccountInfo entity by its id.
func (c *AccountInfoClient) Get(ctx context.Context, id int) (*AccountInfo, error) {
	return c.Query().Where(accountinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountInfoClient) GetX(ctx context.Context, id int) *AccountInfo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountInfoClient) Hooks() []Hook {
	return c.hooks.AccountInfo
}