// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/otakakot/go-orm-s/gen/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/otakakot/go-orm-s/gen/ent/entuser"
	"github.com/otakakot/go-orm-s/gen/ent/entusername"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// EntUser is the client for interacting with the EntUser builders.
	EntUser *EntUserClient
	// EntUserName is the client for interacting with the EntUserName builders.
	EntUserName *EntUserNameClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.EntUser = NewEntUserClient(c.config)
	c.EntUserName = NewEntUserNameClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		EntUser:     NewEntUserClient(cfg),
		EntUserName: NewEntUserNameClient(cfg),
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
		EntUser:     NewEntUserClient(cfg),
		EntUserName: NewEntUserNameClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		EntUser.
//		Query().
//		Count(ctx)
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
	c.EntUser.Use(hooks...)
	c.EntUserName.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.EntUser.Intercept(interceptors...)
	c.EntUserName.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *EntUserMutation:
		return c.EntUser.mutate(ctx, m)
	case *EntUserNameMutation:
		return c.EntUserName.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// EntUserClient is a client for the EntUser schema.
type EntUserClient struct {
	config
}

// NewEntUserClient returns a client for the EntUser from the given config.
func NewEntUserClient(c config) *EntUserClient {
	return &EntUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `entuser.Hooks(f(g(h())))`.
func (c *EntUserClient) Use(hooks ...Hook) {
	c.hooks.EntUser = append(c.hooks.EntUser, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `entuser.Intercept(f(g(h())))`.
func (c *EntUserClient) Intercept(interceptors ...Interceptor) {
	c.inters.EntUser = append(c.inters.EntUser, interceptors...)
}

// Create returns a builder for creating a EntUser entity.
func (c *EntUserClient) Create() *EntUserCreate {
	mutation := newEntUserMutation(c.config, OpCreate)
	return &EntUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EntUser entities.
func (c *EntUserClient) CreateBulk(builders ...*EntUserCreate) *EntUserCreateBulk {
	return &EntUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EntUser.
func (c *EntUserClient) Update() *EntUserUpdate {
	mutation := newEntUserMutation(c.config, OpUpdate)
	return &EntUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EntUserClient) UpdateOne(eu *EntUser) *EntUserUpdateOne {
	mutation := newEntUserMutation(c.config, OpUpdateOne, withEntUser(eu))
	return &EntUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EntUserClient) UpdateOneID(id uuid.UUID) *EntUserUpdateOne {
	mutation := newEntUserMutation(c.config, OpUpdateOne, withEntUserID(id))
	return &EntUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EntUser.
func (c *EntUserClient) Delete() *EntUserDelete {
	mutation := newEntUserMutation(c.config, OpDelete)
	return &EntUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EntUserClient) DeleteOne(eu *EntUser) *EntUserDeleteOne {
	return c.DeleteOneID(eu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EntUserClient) DeleteOneID(id uuid.UUID) *EntUserDeleteOne {
	builder := c.Delete().Where(entuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EntUserDeleteOne{builder}
}

// Query returns a query builder for EntUser.
func (c *EntUserClient) Query() *EntUserQuery {
	return &EntUserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEntUser},
		inters: c.Interceptors(),
	}
}

// Get returns a EntUser entity by its id.
func (c *EntUserClient) Get(ctx context.Context, id uuid.UUID) (*EntUser, error) {
	return c.Query().Where(entuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EntUserClient) GetX(ctx context.Context, id uuid.UUID) *EntUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEntUserNames queries the ent_user_names edge of a EntUser.
func (c *EntUserClient) QueryEntUserNames(eu *EntUser) *EntUserNameQuery {
	query := (&EntUserNameClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := eu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(entuser.Table, entuser.FieldID, id),
			sqlgraph.To(entusername.Table, entusername.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, entuser.EntUserNamesTable, entuser.EntUserNamesColumn),
		)
		fromV = sqlgraph.Neighbors(eu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EntUserClient) Hooks() []Hook {
	return c.hooks.EntUser
}

// Interceptors returns the client interceptors.
func (c *EntUserClient) Interceptors() []Interceptor {
	return c.inters.EntUser
}

func (c *EntUserClient) mutate(ctx context.Context, m *EntUserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EntUserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EntUserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EntUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EntUserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown EntUser mutation op: %q", m.Op())
	}
}

// EntUserNameClient is a client for the EntUserName schema.
type EntUserNameClient struct {
	config
}

// NewEntUserNameClient returns a client for the EntUserName from the given config.
func NewEntUserNameClient(c config) *EntUserNameClient {
	return &EntUserNameClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `entusername.Hooks(f(g(h())))`.
func (c *EntUserNameClient) Use(hooks ...Hook) {
	c.hooks.EntUserName = append(c.hooks.EntUserName, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `entusername.Intercept(f(g(h())))`.
func (c *EntUserNameClient) Intercept(interceptors ...Interceptor) {
	c.inters.EntUserName = append(c.inters.EntUserName, interceptors...)
}

// Create returns a builder for creating a EntUserName entity.
func (c *EntUserNameClient) Create() *EntUserNameCreate {
	mutation := newEntUserNameMutation(c.config, OpCreate)
	return &EntUserNameCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EntUserName entities.
func (c *EntUserNameClient) CreateBulk(builders ...*EntUserNameCreate) *EntUserNameCreateBulk {
	return &EntUserNameCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EntUserName.
func (c *EntUserNameClient) Update() *EntUserNameUpdate {
	mutation := newEntUserNameMutation(c.config, OpUpdate)
	return &EntUserNameUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EntUserNameClient) UpdateOne(eun *EntUserName) *EntUserNameUpdateOne {
	mutation := newEntUserNameMutation(c.config, OpUpdateOne, withEntUserName(eun))
	return &EntUserNameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EntUserNameClient) UpdateOneID(id uuid.UUID) *EntUserNameUpdateOne {
	mutation := newEntUserNameMutation(c.config, OpUpdateOne, withEntUserNameID(id))
	return &EntUserNameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EntUserName.
func (c *EntUserNameClient) Delete() *EntUserNameDelete {
	mutation := newEntUserNameMutation(c.config, OpDelete)
	return &EntUserNameDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EntUserNameClient) DeleteOne(eun *EntUserName) *EntUserNameDeleteOne {
	return c.DeleteOneID(eun.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EntUserNameClient) DeleteOneID(id uuid.UUID) *EntUserNameDeleteOne {
	builder := c.Delete().Where(entusername.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EntUserNameDeleteOne{builder}
}

// Query returns a query builder for EntUserName.
func (c *EntUserNameClient) Query() *EntUserNameQuery {
	return &EntUserNameQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEntUserName},
		inters: c.Interceptors(),
	}
}

// Get returns a EntUserName entity by its id.
func (c *EntUserNameClient) Get(ctx context.Context, id uuid.UUID) (*EntUserName, error) {
	return c.Query().Where(entusername.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EntUserNameClient) GetX(ctx context.Context, id uuid.UUID) *EntUserName {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEntUser queries the ent_user edge of a EntUserName.
func (c *EntUserNameClient) QueryEntUser(eun *EntUserName) *EntUserQuery {
	query := (&EntUserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := eun.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(entusername.Table, entusername.FieldID, id),
			sqlgraph.To(entuser.Table, entuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entusername.EntUserTable, entusername.EntUserColumn),
		)
		fromV = sqlgraph.Neighbors(eun.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EntUserNameClient) Hooks() []Hook {
	return c.hooks.EntUserName
}

// Interceptors returns the client interceptors.
func (c *EntUserNameClient) Interceptors() []Interceptor {
	return c.inters.EntUserName
}

func (c *EntUserNameClient) mutate(ctx context.Context, m *EntUserNameMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EntUserNameCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EntUserNameUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EntUserNameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EntUserNameDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown EntUserName mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		EntUser, EntUserName []ent.Hook
	}
	inters struct {
		EntUser, EntUserName []ent.Interceptor
	}
)
