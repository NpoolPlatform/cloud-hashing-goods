// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodinfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodInfoQuery is the builder for querying GoodInfo entities.
type GoodInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodInfoQuery builder.
func (giq *GoodInfoQuery) Where(ps ...predicate.GoodInfo) *GoodInfoQuery {
	giq.predicates = append(giq.predicates, ps...)
	return giq
}

// Limit adds a limit step to the query.
func (giq *GoodInfoQuery) Limit(limit int) *GoodInfoQuery {
	giq.limit = &limit
	return giq
}

// Offset adds an offset step to the query.
func (giq *GoodInfoQuery) Offset(offset int) *GoodInfoQuery {
	giq.offset = &offset
	return giq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (giq *GoodInfoQuery) Unique(unique bool) *GoodInfoQuery {
	giq.unique = &unique
	return giq
}

// Order adds an order step to the query.
func (giq *GoodInfoQuery) Order(o ...OrderFunc) *GoodInfoQuery {
	giq.order = append(giq.order, o...)
	return giq
}

// First returns the first GoodInfo entity from the query.
// Returns a *NotFoundError when no GoodInfo was found.
func (giq *GoodInfoQuery) First(ctx context.Context) (*GoodInfo, error) {
	nodes, err := giq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (giq *GoodInfoQuery) FirstX(ctx context.Context) *GoodInfo {
	node, err := giq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodInfo ID from the query.
// Returns a *NotFoundError when no GoodInfo ID was found.
func (giq *GoodInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = giq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (giq *GoodInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := giq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GoodInfo entity is not found.
// Returns a *NotFoundError when no GoodInfo entities are found.
func (giq *GoodInfoQuery) Only(ctx context.Context) (*GoodInfo, error) {
	nodes, err := giq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodinfo.Label}
	default:
		return nil, &NotSingularError{goodinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (giq *GoodInfoQuery) OnlyX(ctx context.Context) *GoodInfo {
	node, err := giq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodInfo ID in the query.
// Returns a *NotSingularError when exactly one GoodInfo ID is not found.
// Returns a *NotFoundError when no entities are found.
func (giq *GoodInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = giq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = &NotSingularError{goodinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (giq *GoodInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := giq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodInfos.
func (giq *GoodInfoQuery) All(ctx context.Context) ([]*GoodInfo, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return giq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (giq *GoodInfoQuery) AllX(ctx context.Context) []*GoodInfo {
	nodes, err := giq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodInfo IDs.
func (giq *GoodInfoQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := giq.Select(goodinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (giq *GoodInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := giq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (giq *GoodInfoQuery) Count(ctx context.Context) (int, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return giq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (giq *GoodInfoQuery) CountX(ctx context.Context) int {
	count, err := giq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (giq *GoodInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := giq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return giq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (giq *GoodInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := giq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (giq *GoodInfoQuery) Clone() *GoodInfoQuery {
	if giq == nil {
		return nil
	}
	return &GoodInfoQuery{
		config:     giq.config,
		limit:      giq.limit,
		offset:     giq.offset,
		order:      append([]OrderFunc{}, giq.order...),
		predicates: append([]predicate.GoodInfo{}, giq.predicates...),
		// clone intermediate query.
		sql:  giq.sql.Clone(),
		path: giq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DeviceInfoID uuid.UUID `json:"device_info_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoodInfo.Query().
//		GroupBy(goodinfo.FieldDeviceInfoID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (giq *GoodInfoQuery) GroupBy(field string, fields ...string) *GoodInfoGroupBy {
	group := &GoodInfoGroupBy{config: giq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := giq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return giq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DeviceInfoID uuid.UUID `json:"device_info_id,omitempty"`
//	}
//
//	client.GoodInfo.Query().
//		Select(goodinfo.FieldDeviceInfoID).
//		Scan(ctx, &v)
//
func (giq *GoodInfoQuery) Select(fields ...string) *GoodInfoSelect {
	giq.fields = append(giq.fields, fields...)
	return &GoodInfoSelect{GoodInfoQuery: giq}
}

func (giq *GoodInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range giq.fields {
		if !goodinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if giq.path != nil {
		prev, err := giq.path(ctx)
		if err != nil {
			return err
		}
		giq.sql = prev
	}
	return nil
}

func (giq *GoodInfoQuery) sqlAll(ctx context.Context) ([]*GoodInfo, error) {
	var (
		nodes = []*GoodInfo{}
		_spec = giq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GoodInfo{config: giq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, giq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (giq *GoodInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := giq.querySpec()
	_spec.Node.Columns = giq.fields
	if len(giq.fields) > 0 {
		_spec.Unique = giq.unique != nil && *giq.unique
	}
	return sqlgraph.CountNodes(ctx, giq.driver, _spec)
}

func (giq *GoodInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := giq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (giq *GoodInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodinfo.Table,
			Columns: goodinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodinfo.FieldID,
			},
		},
		From:   giq.sql,
		Unique: true,
	}
	if unique := giq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := giq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodinfo.FieldID)
		for i := range fields {
			if fields[i] != goodinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := giq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := giq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := giq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := giq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (giq *GoodInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(giq.driver.Dialect())
	t1 := builder.Table(goodinfo.Table)
	columns := giq.fields
	if len(columns) == 0 {
		columns = goodinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if giq.sql != nil {
		selector = giq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if giq.unique != nil && *giq.unique {
		selector.Distinct()
	}
	for _, p := range giq.predicates {
		p(selector)
	}
	for _, p := range giq.order {
		p(selector)
	}
	if offset := giq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := giq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodInfoGroupBy is the group-by builder for GoodInfo entities.
type GoodInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gigb *GoodInfoGroupBy) Aggregate(fns ...AggregateFunc) *GoodInfoGroupBy {
	gigb.fns = append(gigb.fns, fns...)
	return gigb
}

// Scan applies the group-by query and scans the result into the given value.
func (gigb *GoodInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gigb.path(ctx)
	if err != nil {
		return err
	}
	gigb.sql = query
	return gigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GoodInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := gigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = fmt.Errorf("ent: GoodInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) StringX(ctx context.Context) string {
	v, err := gigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GoodInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := gigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = fmt.Errorf("ent: GoodInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) IntX(ctx context.Context) int {
	v, err := gigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GoodInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = fmt.Errorf("ent: GoodInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GoodInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gigb *GoodInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = fmt.Errorf("ent: GoodInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gigb *GoodInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := gigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gigb *GoodInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gigb.fields {
		if !goodinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gigb *GoodInfoGroupBy) sqlQuery() *sql.Selector {
	selector := gigb.sql.Select()
	aggregation := make([]string, 0, len(gigb.fns))
	for _, fn := range gigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gigb.fields)+len(gigb.fns))
		for _, f := range gigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gigb.fields...)...)
}

// GoodInfoSelect is the builder for selecting fields of GoodInfo entities.
type GoodInfoSelect struct {
	*GoodInfoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gis *GoodInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gis.prepareQuery(ctx); err != nil {
		return err
	}
	gis.sql = gis.GoodInfoQuery.sqlQuery(ctx)
	return gis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gis *GoodInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gis *GoodInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GoodInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gis *GoodInfoSelect) StringsX(ctx context.Context) []string {
	v, err := gis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gis *GoodInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = fmt.Errorf("ent: GoodInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gis *GoodInfoSelect) StringX(ctx context.Context) string {
	v, err := gis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gis *GoodInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GoodInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gis *GoodInfoSelect) IntsX(ctx context.Context) []int {
	v, err := gis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gis *GoodInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = fmt.Errorf("ent: GoodInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gis *GoodInfoSelect) IntX(ctx context.Context) int {
	v, err := gis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gis *GoodInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GoodInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gis *GoodInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gis *GoodInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = fmt.Errorf("ent: GoodInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gis *GoodInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := gis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gis *GoodInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GoodInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gis *GoodInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := gis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gis *GoodInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodinfo.Label}
	default:
		err = fmt.Errorf("ent: GoodInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gis *GoodInfoSelect) BoolX(ctx context.Context) bool {
	v, err := gis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gis *GoodInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gis.sql.Query()
	if err := gis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
