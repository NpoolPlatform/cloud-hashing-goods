// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgood"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/apptargetarea"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/deviceinfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodinfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/vendorlocation"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AppGood is the client for interacting with the AppGood builders.
	AppGood *AppGoodClient
	// AppTargetArea is the client for interacting with the AppTargetArea builders.
	AppTargetArea *AppTargetAreaClient
	// DeviceInfo is the client for interacting with the DeviceInfo builders.
	DeviceInfo *DeviceInfoClient
	// GoodInfo is the client for interacting with the GoodInfo builders.
	GoodInfo *GoodInfoClient
	// TargetArea is the client for interacting with the TargetArea builders.
	TargetArea *TargetAreaClient
	// VendorLocation is the client for interacting with the VendorLocation builders.
	VendorLocation *VendorLocationClient
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
	c.AppGood = NewAppGoodClient(c.config)
	c.AppTargetArea = NewAppTargetAreaClient(c.config)
	c.DeviceInfo = NewDeviceInfoClient(c.config)
	c.GoodInfo = NewGoodInfoClient(c.config)
	c.TargetArea = NewTargetAreaClient(c.config)
	c.VendorLocation = NewVendorLocationClient(c.config)
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
		ctx:            ctx,
		config:         cfg,
		AppGood:        NewAppGoodClient(cfg),
		AppTargetArea:  NewAppTargetAreaClient(cfg),
		DeviceInfo:     NewDeviceInfoClient(cfg),
		GoodInfo:       NewGoodInfoClient(cfg),
		TargetArea:     NewTargetAreaClient(cfg),
		VendorLocation: NewVendorLocationClient(cfg),
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
		config:         cfg,
		AppGood:        NewAppGoodClient(cfg),
		AppTargetArea:  NewAppTargetAreaClient(cfg),
		DeviceInfo:     NewDeviceInfoClient(cfg),
		GoodInfo:       NewGoodInfoClient(cfg),
		TargetArea:     NewTargetAreaClient(cfg),
		VendorLocation: NewVendorLocationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AppGood.
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
	c.AppGood.Use(hooks...)
	c.AppTargetArea.Use(hooks...)
	c.DeviceInfo.Use(hooks...)
	c.GoodInfo.Use(hooks...)
	c.TargetArea.Use(hooks...)
	c.VendorLocation.Use(hooks...)
}

// AppGoodClient is a client for the AppGood schema.
type AppGoodClient struct {
	config
}

// NewAppGoodClient returns a client for the AppGood from the given config.
func NewAppGoodClient(c config) *AppGoodClient {
	return &AppGoodClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `appgood.Hooks(f(g(h())))`.
func (c *AppGoodClient) Use(hooks ...Hook) {
	c.hooks.AppGood = append(c.hooks.AppGood, hooks...)
}

// Create returns a create builder for AppGood.
func (c *AppGoodClient) Create() *AppGoodCreate {
	mutation := newAppGoodMutation(c.config, OpCreate)
	return &AppGoodCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AppGood entities.
func (c *AppGoodClient) CreateBulk(builders ...*AppGoodCreate) *AppGoodCreateBulk {
	return &AppGoodCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AppGood.
func (c *AppGoodClient) Update() *AppGoodUpdate {
	mutation := newAppGoodMutation(c.config, OpUpdate)
	return &AppGoodUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppGoodClient) UpdateOne(ag *AppGood) *AppGoodUpdateOne {
	mutation := newAppGoodMutation(c.config, OpUpdateOne, withAppGood(ag))
	return &AppGoodUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppGoodClient) UpdateOneID(id uuid.UUID) *AppGoodUpdateOne {
	mutation := newAppGoodMutation(c.config, OpUpdateOne, withAppGoodID(id))
	return &AppGoodUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AppGood.
func (c *AppGoodClient) Delete() *AppGoodDelete {
	mutation := newAppGoodMutation(c.config, OpDelete)
	return &AppGoodDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AppGoodClient) DeleteOne(ag *AppGood) *AppGoodDeleteOne {
	return c.DeleteOneID(ag.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AppGoodClient) DeleteOneID(id uuid.UUID) *AppGoodDeleteOne {
	builder := c.Delete().Where(appgood.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppGoodDeleteOne{builder}
}

// Query returns a query builder for AppGood.
func (c *AppGoodClient) Query() *AppGoodQuery {
	return &AppGoodQuery{
		config: c.config,
	}
}

// Get returns a AppGood entity by its id.
func (c *AppGoodClient) Get(ctx context.Context, id uuid.UUID) (*AppGood, error) {
	return c.Query().Where(appgood.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppGoodClient) GetX(ctx context.Context, id uuid.UUID) *AppGood {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppGoodClient) Hooks() []Hook {
	return c.hooks.AppGood
}

// AppTargetAreaClient is a client for the AppTargetArea schema.
type AppTargetAreaClient struct {
	config
}

// NewAppTargetAreaClient returns a client for the AppTargetArea from the given config.
func NewAppTargetAreaClient(c config) *AppTargetAreaClient {
	return &AppTargetAreaClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `apptargetarea.Hooks(f(g(h())))`.
func (c *AppTargetAreaClient) Use(hooks ...Hook) {
	c.hooks.AppTargetArea = append(c.hooks.AppTargetArea, hooks...)
}

// Create returns a create builder for AppTargetArea.
func (c *AppTargetAreaClient) Create() *AppTargetAreaCreate {
	mutation := newAppTargetAreaMutation(c.config, OpCreate)
	return &AppTargetAreaCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AppTargetArea entities.
func (c *AppTargetAreaClient) CreateBulk(builders ...*AppTargetAreaCreate) *AppTargetAreaCreateBulk {
	return &AppTargetAreaCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AppTargetArea.
func (c *AppTargetAreaClient) Update() *AppTargetAreaUpdate {
	mutation := newAppTargetAreaMutation(c.config, OpUpdate)
	return &AppTargetAreaUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AppTargetAreaClient) UpdateOne(ata *AppTargetArea) *AppTargetAreaUpdateOne {
	mutation := newAppTargetAreaMutation(c.config, OpUpdateOne, withAppTargetArea(ata))
	return &AppTargetAreaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AppTargetAreaClient) UpdateOneID(id uuid.UUID) *AppTargetAreaUpdateOne {
	mutation := newAppTargetAreaMutation(c.config, OpUpdateOne, withAppTargetAreaID(id))
	return &AppTargetAreaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AppTargetArea.
func (c *AppTargetAreaClient) Delete() *AppTargetAreaDelete {
	mutation := newAppTargetAreaMutation(c.config, OpDelete)
	return &AppTargetAreaDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AppTargetAreaClient) DeleteOne(ata *AppTargetArea) *AppTargetAreaDeleteOne {
	return c.DeleteOneID(ata.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AppTargetAreaClient) DeleteOneID(id uuid.UUID) *AppTargetAreaDeleteOne {
	builder := c.Delete().Where(apptargetarea.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AppTargetAreaDeleteOne{builder}
}

// Query returns a query builder for AppTargetArea.
func (c *AppTargetAreaClient) Query() *AppTargetAreaQuery {
	return &AppTargetAreaQuery{
		config: c.config,
	}
}

// Get returns a AppTargetArea entity by its id.
func (c *AppTargetAreaClient) Get(ctx context.Context, id uuid.UUID) (*AppTargetArea, error) {
	return c.Query().Where(apptargetarea.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AppTargetAreaClient) GetX(ctx context.Context, id uuid.UUID) *AppTargetArea {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AppTargetAreaClient) Hooks() []Hook {
	return c.hooks.AppTargetArea
}

// DeviceInfoClient is a client for the DeviceInfo schema.
type DeviceInfoClient struct {
	config
}

// NewDeviceInfoClient returns a client for the DeviceInfo from the given config.
func NewDeviceInfoClient(c config) *DeviceInfoClient {
	return &DeviceInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `deviceinfo.Hooks(f(g(h())))`.
func (c *DeviceInfoClient) Use(hooks ...Hook) {
	c.hooks.DeviceInfo = append(c.hooks.DeviceInfo, hooks...)
}

// Create returns a create builder for DeviceInfo.
func (c *DeviceInfoClient) Create() *DeviceInfoCreate {
	mutation := newDeviceInfoMutation(c.config, OpCreate)
	return &DeviceInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DeviceInfo entities.
func (c *DeviceInfoClient) CreateBulk(builders ...*DeviceInfoCreate) *DeviceInfoCreateBulk {
	return &DeviceInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DeviceInfo.
func (c *DeviceInfoClient) Update() *DeviceInfoUpdate {
	mutation := newDeviceInfoMutation(c.config, OpUpdate)
	return &DeviceInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DeviceInfoClient) UpdateOne(di *DeviceInfo) *DeviceInfoUpdateOne {
	mutation := newDeviceInfoMutation(c.config, OpUpdateOne, withDeviceInfo(di))
	return &DeviceInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DeviceInfoClient) UpdateOneID(id uuid.UUID) *DeviceInfoUpdateOne {
	mutation := newDeviceInfoMutation(c.config, OpUpdateOne, withDeviceInfoID(id))
	return &DeviceInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DeviceInfo.
func (c *DeviceInfoClient) Delete() *DeviceInfoDelete {
	mutation := newDeviceInfoMutation(c.config, OpDelete)
	return &DeviceInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DeviceInfoClient) DeleteOne(di *DeviceInfo) *DeviceInfoDeleteOne {
	return c.DeleteOneID(di.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DeviceInfoClient) DeleteOneID(id uuid.UUID) *DeviceInfoDeleteOne {
	builder := c.Delete().Where(deviceinfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DeviceInfoDeleteOne{builder}
}

// Query returns a query builder for DeviceInfo.
func (c *DeviceInfoClient) Query() *DeviceInfoQuery {
	return &DeviceInfoQuery{
		config: c.config,
	}
}

// Get returns a DeviceInfo entity by its id.
func (c *DeviceInfoClient) Get(ctx context.Context, id uuid.UUID) (*DeviceInfo, error) {
	return c.Query().Where(deviceinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DeviceInfoClient) GetX(ctx context.Context, id uuid.UUID) *DeviceInfo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DeviceInfoClient) Hooks() []Hook {
	return c.hooks.DeviceInfo
}

// GoodInfoClient is a client for the GoodInfo schema.
type GoodInfoClient struct {
	config
}

// NewGoodInfoClient returns a client for the GoodInfo from the given config.
func NewGoodInfoClient(c config) *GoodInfoClient {
	return &GoodInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `goodinfo.Hooks(f(g(h())))`.
func (c *GoodInfoClient) Use(hooks ...Hook) {
	c.hooks.GoodInfo = append(c.hooks.GoodInfo, hooks...)
}

// Create returns a create builder for GoodInfo.
func (c *GoodInfoClient) Create() *GoodInfoCreate {
	mutation := newGoodInfoMutation(c.config, OpCreate)
	return &GoodInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GoodInfo entities.
func (c *GoodInfoClient) CreateBulk(builders ...*GoodInfoCreate) *GoodInfoCreateBulk {
	return &GoodInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GoodInfo.
func (c *GoodInfoClient) Update() *GoodInfoUpdate {
	mutation := newGoodInfoMutation(c.config, OpUpdate)
	return &GoodInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GoodInfoClient) UpdateOne(gi *GoodInfo) *GoodInfoUpdateOne {
	mutation := newGoodInfoMutation(c.config, OpUpdateOne, withGoodInfo(gi))
	return &GoodInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GoodInfoClient) UpdateOneID(id uuid.UUID) *GoodInfoUpdateOne {
	mutation := newGoodInfoMutation(c.config, OpUpdateOne, withGoodInfoID(id))
	return &GoodInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GoodInfo.
func (c *GoodInfoClient) Delete() *GoodInfoDelete {
	mutation := newGoodInfoMutation(c.config, OpDelete)
	return &GoodInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GoodInfoClient) DeleteOne(gi *GoodInfo) *GoodInfoDeleteOne {
	return c.DeleteOneID(gi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GoodInfoClient) DeleteOneID(id uuid.UUID) *GoodInfoDeleteOne {
	builder := c.Delete().Where(goodinfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GoodInfoDeleteOne{builder}
}

// Query returns a query builder for GoodInfo.
func (c *GoodInfoClient) Query() *GoodInfoQuery {
	return &GoodInfoQuery{
		config: c.config,
	}
}

// Get returns a GoodInfo entity by its id.
func (c *GoodInfoClient) Get(ctx context.Context, id uuid.UUID) (*GoodInfo, error) {
	return c.Query().Where(goodinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GoodInfoClient) GetX(ctx context.Context, id uuid.UUID) *GoodInfo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GoodInfoClient) Hooks() []Hook {
	return c.hooks.GoodInfo
}

// TargetAreaClient is a client for the TargetArea schema.
type TargetAreaClient struct {
	config
}

// NewTargetAreaClient returns a client for the TargetArea from the given config.
func NewTargetAreaClient(c config) *TargetAreaClient {
	return &TargetAreaClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `targetarea.Hooks(f(g(h())))`.
func (c *TargetAreaClient) Use(hooks ...Hook) {
	c.hooks.TargetArea = append(c.hooks.TargetArea, hooks...)
}

// Create returns a create builder for TargetArea.
func (c *TargetAreaClient) Create() *TargetAreaCreate {
	mutation := newTargetAreaMutation(c.config, OpCreate)
	return &TargetAreaCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TargetArea entities.
func (c *TargetAreaClient) CreateBulk(builders ...*TargetAreaCreate) *TargetAreaCreateBulk {
	return &TargetAreaCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TargetArea.
func (c *TargetAreaClient) Update() *TargetAreaUpdate {
	mutation := newTargetAreaMutation(c.config, OpUpdate)
	return &TargetAreaUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TargetAreaClient) UpdateOne(ta *TargetArea) *TargetAreaUpdateOne {
	mutation := newTargetAreaMutation(c.config, OpUpdateOne, withTargetArea(ta))
	return &TargetAreaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TargetAreaClient) UpdateOneID(id uuid.UUID) *TargetAreaUpdateOne {
	mutation := newTargetAreaMutation(c.config, OpUpdateOne, withTargetAreaID(id))
	return &TargetAreaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TargetArea.
func (c *TargetAreaClient) Delete() *TargetAreaDelete {
	mutation := newTargetAreaMutation(c.config, OpDelete)
	return &TargetAreaDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TargetAreaClient) DeleteOne(ta *TargetArea) *TargetAreaDeleteOne {
	return c.DeleteOneID(ta.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TargetAreaClient) DeleteOneID(id uuid.UUID) *TargetAreaDeleteOne {
	builder := c.Delete().Where(targetarea.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TargetAreaDeleteOne{builder}
}

// Query returns a query builder for TargetArea.
func (c *TargetAreaClient) Query() *TargetAreaQuery {
	return &TargetAreaQuery{
		config: c.config,
	}
}

// Get returns a TargetArea entity by its id.
func (c *TargetAreaClient) Get(ctx context.Context, id uuid.UUID) (*TargetArea, error) {
	return c.Query().Where(targetarea.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TargetAreaClient) GetX(ctx context.Context, id uuid.UUID) *TargetArea {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TargetAreaClient) Hooks() []Hook {
	return c.hooks.TargetArea
}

// VendorLocationClient is a client for the VendorLocation schema.
type VendorLocationClient struct {
	config
}

// NewVendorLocationClient returns a client for the VendorLocation from the given config.
func NewVendorLocationClient(c config) *VendorLocationClient {
	return &VendorLocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `vendorlocation.Hooks(f(g(h())))`.
func (c *VendorLocationClient) Use(hooks ...Hook) {
	c.hooks.VendorLocation = append(c.hooks.VendorLocation, hooks...)
}

// Create returns a create builder for VendorLocation.
func (c *VendorLocationClient) Create() *VendorLocationCreate {
	mutation := newVendorLocationMutation(c.config, OpCreate)
	return &VendorLocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of VendorLocation entities.
func (c *VendorLocationClient) CreateBulk(builders ...*VendorLocationCreate) *VendorLocationCreateBulk {
	return &VendorLocationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for VendorLocation.
func (c *VendorLocationClient) Update() *VendorLocationUpdate {
	mutation := newVendorLocationMutation(c.config, OpUpdate)
	return &VendorLocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VendorLocationClient) UpdateOne(vl *VendorLocation) *VendorLocationUpdateOne {
	mutation := newVendorLocationMutation(c.config, OpUpdateOne, withVendorLocation(vl))
	return &VendorLocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VendorLocationClient) UpdateOneID(id uuid.UUID) *VendorLocationUpdateOne {
	mutation := newVendorLocationMutation(c.config, OpUpdateOne, withVendorLocationID(id))
	return &VendorLocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for VendorLocation.
func (c *VendorLocationClient) Delete() *VendorLocationDelete {
	mutation := newVendorLocationMutation(c.config, OpDelete)
	return &VendorLocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VendorLocationClient) DeleteOne(vl *VendorLocation) *VendorLocationDeleteOne {
	return c.DeleteOneID(vl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VendorLocationClient) DeleteOneID(id uuid.UUID) *VendorLocationDeleteOne {
	builder := c.Delete().Where(vendorlocation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VendorLocationDeleteOne{builder}
}

// Query returns a query builder for VendorLocation.
func (c *VendorLocationClient) Query() *VendorLocationQuery {
	return &VendorLocationQuery{
		config: c.config,
	}
}

// Get returns a VendorLocation entity by its id.
func (c *VendorLocationClient) Get(ctx context.Context, id uuid.UUID) (*VendorLocation, error) {
	return c.Query().Where(vendorlocation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VendorLocationClient) GetX(ctx context.Context, id uuid.UUID) *VendorLocation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VendorLocationClient) Hooks() []Hook {
	return c.hooks.VendorLocation
}
