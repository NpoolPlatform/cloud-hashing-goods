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
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodextrainfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodExtraInfoQuery is the builder for querying GoodExtraInfo entities.
type GoodExtraInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodExtraInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodExtraInfoQuery builder.
func (geiq *GoodExtraInfoQuery) Where(ps ...predicate.GoodExtraInfo) *GoodExtraInfoQuery {
	geiq.predicates = append(geiq.predicates, ps...)
	return geiq
}

// Limit adds a limit step to the query.
func (geiq *GoodExtraInfoQuery) Limit(limit int) *GoodExtraInfoQuery {
	geiq.limit = &limit
	return geiq
}

// Offset adds an offset step to the query.
func (geiq *GoodExtraInfoQuery) Offset(offset int) *GoodExtraInfoQuery {
	geiq.offset = &offset
	return geiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (geiq *GoodExtraInfoQuery) Unique(unique bool) *GoodExtraInfoQuery {
	geiq.unique = &unique
	return geiq
}

// Order adds an order step to the query.
func (geiq *GoodExtraInfoQuery) Order(o ...OrderFunc) *GoodExtraInfoQuery {
	geiq.order = append(geiq.order, o...)
	return geiq
}

// First returns the first GoodExtraInfo entity from the query.
// Returns a *NotFoundError when no GoodExtraInfo was found.
func (geiq *GoodExtraInfoQuery) First(ctx context.Context) (*GoodExtraInfo, error) {
	nodes, err := geiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodextrainfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (geiq *GoodExtraInfoQuery) FirstX(ctx context.Context) *GoodExtraInfo {
	node, err := geiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodExtraInfo ID from the query.
// Returns a *NotFoundError when no GoodExtraInfo ID was found.
func (geiq *GoodExtraInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = geiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodextrainfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (geiq *GoodExtraInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := geiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodExtraInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GoodExtraInfo entity is found.
// Returns a *NotFoundError when no GoodExtraInfo entities are found.
func (geiq *GoodExtraInfoQuery) Only(ctx context.Context) (*GoodExtraInfo, error) {
	nodes, err := geiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodextrainfo.Label}
	default:
		return nil, &NotSingularError{goodextrainfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (geiq *GoodExtraInfoQuery) OnlyX(ctx context.Context) *GoodExtraInfo {
	node, err := geiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodExtraInfo ID in the query.
// Returns a *NotSingularError when more than one GoodExtraInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (geiq *GoodExtraInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = geiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = &NotSingularError{goodextrainfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (geiq *GoodExtraInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := geiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodExtraInfos.
func (geiq *GoodExtraInfoQuery) All(ctx context.Context) ([]*GoodExtraInfo, error) {
	if err := geiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return geiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (geiq *GoodExtraInfoQuery) AllX(ctx context.Context) []*GoodExtraInfo {
	nodes, err := geiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodExtraInfo IDs.
func (geiq *GoodExtraInfoQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := geiq.Select(goodextrainfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (geiq *GoodExtraInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := geiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (geiq *GoodExtraInfoQuery) Count(ctx context.Context) (int, error) {
	if err := geiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return geiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (geiq *GoodExtraInfoQuery) CountX(ctx context.Context) int {
	count, err := geiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (geiq *GoodExtraInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := geiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return geiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (geiq *GoodExtraInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := geiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodExtraInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (geiq *GoodExtraInfoQuery) Clone() *GoodExtraInfoQuery {
	if geiq == nil {
		return nil
	}
	return &GoodExtraInfoQuery{
		config:     geiq.config,
		limit:      geiq.limit,
		offset:     geiq.offset,
		order:      append([]OrderFunc{}, geiq.order...),
		predicates: append([]predicate.GoodExtraInfo{}, geiq.predicates...),
		// clone intermediate query.
		sql:    geiq.sql.Clone(),
		path:   geiq.path,
		unique: geiq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GoodID uuid.UUID `json:"good_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoodExtraInfo.Query().
//		GroupBy(goodextrainfo.FieldGoodID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (geiq *GoodExtraInfoQuery) GroupBy(field string, fields ...string) *GoodExtraInfoGroupBy {
	group := &GoodExtraInfoGroupBy{config: geiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := geiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return geiq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GoodID uuid.UUID `json:"good_id,omitempty"`
//	}
//
//	client.GoodExtraInfo.Query().
//		Select(goodextrainfo.FieldGoodID).
//		Scan(ctx, &v)
//
func (geiq *GoodExtraInfoQuery) Select(fields ...string) *GoodExtraInfoSelect {
	geiq.fields = append(geiq.fields, fields...)
	return &GoodExtraInfoSelect{GoodExtraInfoQuery: geiq}
}

func (geiq *GoodExtraInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range geiq.fields {
		if !goodextrainfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if geiq.path != nil {
		prev, err := geiq.path(ctx)
		if err != nil {
			return err
		}
		geiq.sql = prev
	}
	return nil
}

func (geiq *GoodExtraInfoQuery) sqlAll(ctx context.Context) ([]*GoodExtraInfo, error) {
	var (
		nodes = []*GoodExtraInfo{}
		_spec = geiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GoodExtraInfo{config: geiq.config}
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
	if err := sqlgraph.QueryNodes(ctx, geiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (geiq *GoodExtraInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := geiq.querySpec()
	_spec.Node.Columns = geiq.fields
	if len(geiq.fields) > 0 {
		_spec.Unique = geiq.unique != nil && *geiq.unique
	}
	return sqlgraph.CountNodes(ctx, geiq.driver, _spec)
}

func (geiq *GoodExtraInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := geiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (geiq *GoodExtraInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodextrainfo.Table,
			Columns: goodextrainfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodextrainfo.FieldID,
			},
		},
		From:   geiq.sql,
		Unique: true,
	}
	if unique := geiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := geiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodextrainfo.FieldID)
		for i := range fields {
			if fields[i] != goodextrainfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := geiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := geiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := geiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := geiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (geiq *GoodExtraInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(geiq.driver.Dialect())
	t1 := builder.Table(goodextrainfo.Table)
	columns := geiq.fields
	if len(columns) == 0 {
		columns = goodextrainfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if geiq.sql != nil {
		selector = geiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if geiq.unique != nil && *geiq.unique {
		selector.Distinct()
	}
	for _, p := range geiq.predicates {
		p(selector)
	}
	for _, p := range geiq.order {
		p(selector)
	}
	if offset := geiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := geiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodExtraInfoGroupBy is the group-by builder for GoodExtraInfo entities.
type GoodExtraInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (geigb *GoodExtraInfoGroupBy) Aggregate(fns ...AggregateFunc) *GoodExtraInfoGroupBy {
	geigb.fns = append(geigb.fns, fns...)
	return geigb
}

// Scan applies the group-by query and scans the result into the given value.
func (geigb *GoodExtraInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := geigb.path(ctx)
	if err != nil {
		return err
	}
	geigb.sql = query
	return geigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := geigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (geigb *GoodExtraInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(geigb.fields) > 1 {
		return nil, errors.New("ent: GoodExtraInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := geigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := geigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (geigb *GoodExtraInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = geigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = fmt.Errorf("ent: GoodExtraInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) StringX(ctx context.Context) string {
	v, err := geigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (geigb *GoodExtraInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(geigb.fields) > 1 {
		return nil, errors.New("ent: GoodExtraInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := geigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := geigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (geigb *GoodExtraInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = geigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = fmt.Errorf("ent: GoodExtraInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) IntX(ctx context.Context) int {
	v, err := geigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (geigb *GoodExtraInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(geigb.fields) > 1 {
		return nil, errors.New("ent: GoodExtraInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := geigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := geigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (geigb *GoodExtraInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = geigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = fmt.Errorf("ent: GoodExtraInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := geigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (geigb *GoodExtraInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(geigb.fields) > 1 {
		return nil, errors.New("ent: GoodExtraInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := geigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := geigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (geigb *GoodExtraInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = geigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = fmt.Errorf("ent: GoodExtraInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (geigb *GoodExtraInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := geigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (geigb *GoodExtraInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range geigb.fields {
		if !goodextrainfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := geigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := geigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (geigb *GoodExtraInfoGroupBy) sqlQuery() *sql.Selector {
	selector := geigb.sql.Select()
	aggregation := make([]string, 0, len(geigb.fns))
	for _, fn := range geigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(geigb.fields)+len(geigb.fns))
		for _, f := range geigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(geigb.fields...)...)
}

// GoodExtraInfoSelect is the builder for selecting fields of GoodExtraInfo entities.
type GoodExtraInfoSelect struct {
	*GoodExtraInfoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (geis *GoodExtraInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := geis.prepareQuery(ctx); err != nil {
		return err
	}
	geis.sql = geis.GoodExtraInfoQuery.sqlQuery(ctx)
	return geis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := geis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (geis *GoodExtraInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(geis.fields) > 1 {
		return nil, errors.New("ent: GoodExtraInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := geis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) StringsX(ctx context.Context) []string {
	v, err := geis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (geis *GoodExtraInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = geis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = fmt.Errorf("ent: GoodExtraInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) StringX(ctx context.Context) string {
	v, err := geis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (geis *GoodExtraInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(geis.fields) > 1 {
		return nil, errors.New("ent: GoodExtraInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := geis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) IntsX(ctx context.Context) []int {
	v, err := geis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (geis *GoodExtraInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = geis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = fmt.Errorf("ent: GoodExtraInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) IntX(ctx context.Context) int {
	v, err := geis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (geis *GoodExtraInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(geis.fields) > 1 {
		return nil, errors.New("ent: GoodExtraInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := geis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := geis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (geis *GoodExtraInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = geis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = fmt.Errorf("ent: GoodExtraInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := geis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (geis *GoodExtraInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(geis.fields) > 1 {
		return nil, errors.New("ent: GoodExtraInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := geis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := geis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (geis *GoodExtraInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = geis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodextrainfo.Label}
	default:
		err = fmt.Errorf("ent: GoodExtraInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (geis *GoodExtraInfoSelect) BoolX(ctx context.Context) bool {
	v, err := geis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (geis *GoodExtraInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := geis.sql.Query()
	if err := geis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
