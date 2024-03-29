// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"boot/ent/migrate"

	"boot/ent/machine"
	"boot/ent/notice"
	"boot/ent/role"
	"boot/ent/tool"
	"boot/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Machine is the client for interacting with the Machine builders.
	Machine *MachineClient
	// Notice is the client for interacting with the Notice builders.
	Notice *NoticeClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// Tool is the client for interacting with the Tool builders.
	Tool *ToolClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.Machine = NewMachineClient(c.config)
	c.Notice = NewNoticeClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.Tool = NewToolClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:     ctx,
		config:  cfg,
		Machine: NewMachineClient(cfg),
		Notice:  NewNoticeClient(cfg),
		Role:    NewRoleClient(cfg),
		Tool:    NewToolClient(cfg),
		User:    NewUserClient(cfg),
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
		config:  cfg,
		Machine: NewMachineClient(cfg),
		Notice:  NewNoticeClient(cfg),
		Role:    NewRoleClient(cfg),
		Tool:    NewToolClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Machine.
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
	c.Machine.Use(hooks...)
	c.Notice.Use(hooks...)
	c.Role.Use(hooks...)
	c.Tool.Use(hooks...)
	c.User.Use(hooks...)
}

// MachineClient is a client for the Machine schema.
type MachineClient struct {
	config
}

// NewMachineClient returns a client for the Machine from the given config.
func NewMachineClient(c config) *MachineClient {
	return &MachineClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `machine.Hooks(f(g(h())))`.
func (c *MachineClient) Use(hooks ...Hook) {
	c.hooks.Machine = append(c.hooks.Machine, hooks...)
}

// Create returns a create builder for Machine.
func (c *MachineClient) Create() *MachineCreate {
	mutation := newMachineMutation(c.config, OpCreate)
	return &MachineCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Machine entities.
func (c *MachineClient) CreateBulk(builders ...*MachineCreate) *MachineCreateBulk {
	return &MachineCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Machine.
func (c *MachineClient) Update() *MachineUpdate {
	mutation := newMachineMutation(c.config, OpUpdate)
	return &MachineUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MachineClient) UpdateOne(m *Machine) *MachineUpdateOne {
	mutation := newMachineMutation(c.config, OpUpdateOne, withMachine(m))
	return &MachineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MachineClient) UpdateOneID(id int) *MachineUpdateOne {
	mutation := newMachineMutation(c.config, OpUpdateOne, withMachineID(id))
	return &MachineUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Machine.
func (c *MachineClient) Delete() *MachineDelete {
	mutation := newMachineMutation(c.config, OpDelete)
	return &MachineDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MachineClient) DeleteOne(m *Machine) *MachineDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MachineClient) DeleteOneID(id int) *MachineDeleteOne {
	builder := c.Delete().Where(machine.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MachineDeleteOne{builder}
}

// Query returns a query builder for Machine.
func (c *MachineClient) Query() *MachineQuery {
	return &MachineQuery{
		config: c.config,
	}
}

// Get returns a Machine entity by its id.
func (c *MachineClient) Get(ctx context.Context, id int) (*Machine, error) {
	return c.Query().Where(machine.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MachineClient) GetX(ctx context.Context, id int) *Machine {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MachineClient) Hooks() []Hook {
	return c.hooks.Machine
}

// NoticeClient is a client for the Notice schema.
type NoticeClient struct {
	config
}

// NewNoticeClient returns a client for the Notice from the given config.
func NewNoticeClient(c config) *NoticeClient {
	return &NoticeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notice.Hooks(f(g(h())))`.
func (c *NoticeClient) Use(hooks ...Hook) {
	c.hooks.Notice = append(c.hooks.Notice, hooks...)
}

// Create returns a create builder for Notice.
func (c *NoticeClient) Create() *NoticeCreate {
	mutation := newNoticeMutation(c.config, OpCreate)
	return &NoticeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Notice entities.
func (c *NoticeClient) CreateBulk(builders ...*NoticeCreate) *NoticeCreateBulk {
	return &NoticeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Notice.
func (c *NoticeClient) Update() *NoticeUpdate {
	mutation := newNoticeMutation(c.config, OpUpdate)
	return &NoticeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NoticeClient) UpdateOne(n *Notice) *NoticeUpdateOne {
	mutation := newNoticeMutation(c.config, OpUpdateOne, withNotice(n))
	return &NoticeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NoticeClient) UpdateOneID(id int) *NoticeUpdateOne {
	mutation := newNoticeMutation(c.config, OpUpdateOne, withNoticeID(id))
	return &NoticeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Notice.
func (c *NoticeClient) Delete() *NoticeDelete {
	mutation := newNoticeMutation(c.config, OpDelete)
	return &NoticeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NoticeClient) DeleteOne(n *Notice) *NoticeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NoticeClient) DeleteOneID(id int) *NoticeDeleteOne {
	builder := c.Delete().Where(notice.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NoticeDeleteOne{builder}
}

// Query returns a query builder for Notice.
func (c *NoticeClient) Query() *NoticeQuery {
	return &NoticeQuery{
		config: c.config,
	}
}

// Get returns a Notice entity by its id.
func (c *NoticeClient) Get(ctx context.Context, id int) (*Notice, error) {
	return c.Query().Where(notice.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NoticeClient) GetX(ctx context.Context, id int) *Notice {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NoticeClient) Hooks() []Hook {
	return c.hooks.Notice
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Create returns a create builder for Role.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a Role.
func (c *RoleClient) QueryRoles(r *Role) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, role.RolesTable, role.RolesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// ToolClient is a client for the Tool schema.
type ToolClient struct {
	config
}

// NewToolClient returns a client for the Tool from the given config.
func NewToolClient(c config) *ToolClient {
	return &ToolClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tool.Hooks(f(g(h())))`.
func (c *ToolClient) Use(hooks ...Hook) {
	c.hooks.Tool = append(c.hooks.Tool, hooks...)
}

// Create returns a create builder for Tool.
func (c *ToolClient) Create() *ToolCreate {
	mutation := newToolMutation(c.config, OpCreate)
	return &ToolCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tool entities.
func (c *ToolClient) CreateBulk(builders ...*ToolCreate) *ToolCreateBulk {
	return &ToolCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tool.
func (c *ToolClient) Update() *ToolUpdate {
	mutation := newToolMutation(c.config, OpUpdate)
	return &ToolUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ToolClient) UpdateOne(t *Tool) *ToolUpdateOne {
	mutation := newToolMutation(c.config, OpUpdateOne, withTool(t))
	return &ToolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ToolClient) UpdateOneID(id int) *ToolUpdateOne {
	mutation := newToolMutation(c.config, OpUpdateOne, withToolID(id))
	return &ToolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tool.
func (c *ToolClient) Delete() *ToolDelete {
	mutation := newToolMutation(c.config, OpDelete)
	return &ToolDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ToolClient) DeleteOne(t *Tool) *ToolDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ToolClient) DeleteOneID(id int) *ToolDeleteOne {
	builder := c.Delete().Where(tool.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ToolDeleteOne{builder}
}

// Query returns a query builder for Tool.
func (c *ToolClient) Query() *ToolQuery {
	return &ToolQuery{
		config: c.config,
	}
}

// Get returns a Tool entity by its id.
func (c *ToolClient) Get(ctx context.Context, id int) (*Tool, error) {
	return c.Query().Where(tool.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ToolClient) GetX(ctx context.Context, id int) *Tool {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ToolClient) Hooks() []Hook {
	return c.hooks.Tool
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a User.
func (c *UserClient) QueryOwner(u *User) *RoleQuery {
	query := &RoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.OwnerTable, user.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
