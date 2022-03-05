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
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/pricecurrency"
	"github.com/google/uuid"
)

// PriceCurrencyQuery is the builder for querying PriceCurrency entities.
type PriceCurrencyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PriceCurrency
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PriceCurrencyQuery builder.
func (pcq *PriceCurrencyQuery) Where(ps ...predicate.PriceCurrency) *PriceCurrencyQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit adds a limit step to the query.
func (pcq *PriceCurrencyQuery) Limit(limit int) *PriceCurrencyQuery {
	pcq.limit = &limit
	return pcq
}

// Offset adds an offset step to the query.
func (pcq *PriceCurrencyQuery) Offset(offset int) *PriceCurrencyQuery {
	pcq.offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PriceCurrencyQuery) Unique(unique bool) *PriceCurrencyQuery {
	pcq.unique = &unique
	return pcq
}

// Order adds an order step to the query.
func (pcq *PriceCurrencyQuery) Order(o ...OrderFunc) *PriceCurrencyQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// First returns the first PriceCurrency entity from the query.
// Returns a *NotFoundError when no PriceCurrency was found.
func (pcq *PriceCurrencyQuery) First(ctx context.Context) (*PriceCurrency, error) {
	nodes, err := pcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pricecurrency.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *PriceCurrencyQuery) FirstX(ctx context.Context) *PriceCurrency {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PriceCurrency ID from the query.
// Returns a *NotFoundError when no PriceCurrency ID was found.
func (pcq *PriceCurrencyQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pricecurrency.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *PriceCurrencyQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PriceCurrency entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PriceCurrency entity is found.
// Returns a *NotFoundError when no PriceCurrency entities are found.
func (pcq *PriceCurrencyQuery) Only(ctx context.Context) (*PriceCurrency, error) {
	nodes, err := pcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pricecurrency.Label}
	default:
		return nil, &NotSingularError{pricecurrency.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *PriceCurrencyQuery) OnlyX(ctx context.Context) *PriceCurrency {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PriceCurrency ID in the query.
// Returns a *NotSingularError when more than one PriceCurrency ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *PriceCurrencyQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = &NotSingularError{pricecurrency.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *PriceCurrencyQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PriceCurrencies.
func (pcq *PriceCurrencyQuery) All(ctx context.Context) ([]*PriceCurrency, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pcq *PriceCurrencyQuery) AllX(ctx context.Context) []*PriceCurrency {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PriceCurrency IDs.
func (pcq *PriceCurrencyQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := pcq.Select(pricecurrency.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *PriceCurrencyQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *PriceCurrencyQuery) Count(ctx context.Context) (int, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *PriceCurrencyQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *PriceCurrencyQuery) Exist(ctx context.Context) (bool, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *PriceCurrencyQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PriceCurrencyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *PriceCurrencyQuery) Clone() *PriceCurrencyQuery {
	if pcq == nil {
		return nil
	}
	return &PriceCurrencyQuery{
		config:     pcq.config,
		limit:      pcq.limit,
		offset:     pcq.offset,
		order:      append([]OrderFunc{}, pcq.order...),
		predicates: append([]predicate.PriceCurrency{}, pcq.predicates...),
		// clone intermediate query.
		sql:    pcq.sql.Clone(),
		path:   pcq.path,
		unique: pcq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PriceCurrency.Query().
//		GroupBy(pricecurrency.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pcq *PriceCurrencyQuery) GroupBy(field string, fields ...string) *PriceCurrencyGroupBy {
	group := &PriceCurrencyGroupBy{config: pcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.PriceCurrency.Query().
//		Select(pricecurrency.FieldName).
//		Scan(ctx, &v)
//
func (pcq *PriceCurrencyQuery) Select(fields ...string) *PriceCurrencySelect {
	pcq.fields = append(pcq.fields, fields...)
	return &PriceCurrencySelect{PriceCurrencyQuery: pcq}
}

func (pcq *PriceCurrencyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pcq.fields {
		if !pricecurrency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *PriceCurrencyQuery) sqlAll(ctx context.Context) ([]*PriceCurrency, error) {
	var (
		nodes = []*PriceCurrency{}
		_spec = pcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PriceCurrency{config: pcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pcq *PriceCurrencyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	_spec.Node.Columns = pcq.fields
	if len(pcq.fields) > 0 {
		_spec.Unique = pcq.unique != nil && *pcq.unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PriceCurrencyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pcq *PriceCurrencyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pricecurrency.Table,
			Columns: pricecurrency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: pricecurrency.FieldID,
			},
		},
		From:   pcq.sql,
		Unique: true,
	}
	if unique := pcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pricecurrency.FieldID)
		for i := range fields {
			if fields[i] != pricecurrency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *PriceCurrencyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(pricecurrency.Table)
	columns := pcq.fields
	if len(columns) == 0 {
		columns = pricecurrency.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.unique != nil && *pcq.unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PriceCurrencyGroupBy is the group-by builder for PriceCurrency entities.
type PriceCurrencyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PriceCurrencyGroupBy) Aggregate(fns ...AggregateFunc) *PriceCurrencyGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pcgb *PriceCurrencyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pcgb.path(ctx)
	if err != nil {
		return err
	}
	pcgb.sql = query
	return pcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcgb *PriceCurrencyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pcgb.fields) > 1 {
		return nil, errors.New("ent: PriceCurrencyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) StringsX(ctx context.Context) []string {
	v, err := pcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcgb *PriceCurrencyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = fmt.Errorf("ent: PriceCurrencyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) StringX(ctx context.Context) string {
	v, err := pcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcgb *PriceCurrencyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pcgb.fields) > 1 {
		return nil, errors.New("ent: PriceCurrencyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) IntsX(ctx context.Context) []int {
	v, err := pcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcgb *PriceCurrencyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = fmt.Errorf("ent: PriceCurrencyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) IntX(ctx context.Context) int {
	v, err := pcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcgb *PriceCurrencyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcgb.fields) > 1 {
		return nil, errors.New("ent: PriceCurrencyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcgb *PriceCurrencyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = fmt.Errorf("ent: PriceCurrencyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcgb *PriceCurrencyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pcgb.fields) > 1 {
		return nil, errors.New("ent: PriceCurrencyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcgb *PriceCurrencyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = fmt.Errorf("ent: PriceCurrencyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcgb *PriceCurrencyGroupBy) BoolX(ctx context.Context) bool {
	v, err := pcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcgb *PriceCurrencyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pcgb.fields {
		if !pricecurrency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pcgb *PriceCurrencyGroupBy) sqlQuery() *sql.Selector {
	selector := pcgb.sql.Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pcgb.fields)+len(pcgb.fns))
		for _, f := range pcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pcgb.fields...)...)
}

// PriceCurrencySelect is the builder for selecting fields of PriceCurrency entities.
type PriceCurrencySelect struct {
	*PriceCurrencyQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PriceCurrencySelect) Scan(ctx context.Context, v interface{}) error {
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	pcs.sql = pcs.PriceCurrencyQuery.sqlQuery(ctx)
	return pcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcs *PriceCurrencySelect) ScanX(ctx context.Context, v interface{}) {
	if err := pcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pcs *PriceCurrencySelect) Strings(ctx context.Context) ([]string, error) {
	if len(pcs.fields) > 1 {
		return nil, errors.New("ent: PriceCurrencySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcs *PriceCurrencySelect) StringsX(ctx context.Context) []string {
	v, err := pcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pcs *PriceCurrencySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = fmt.Errorf("ent: PriceCurrencySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcs *PriceCurrencySelect) StringX(ctx context.Context) string {
	v, err := pcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pcs *PriceCurrencySelect) Ints(ctx context.Context) ([]int, error) {
	if len(pcs.fields) > 1 {
		return nil, errors.New("ent: PriceCurrencySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcs *PriceCurrencySelect) IntsX(ctx context.Context) []int {
	v, err := pcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pcs *PriceCurrencySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = fmt.Errorf("ent: PriceCurrencySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcs *PriceCurrencySelect) IntX(ctx context.Context) int {
	v, err := pcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pcs *PriceCurrencySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcs.fields) > 1 {
		return nil, errors.New("ent: PriceCurrencySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcs *PriceCurrencySelect) Float64sX(ctx context.Context) []float64 {
	v, err := pcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pcs *PriceCurrencySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = fmt.Errorf("ent: PriceCurrencySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcs *PriceCurrencySelect) Float64X(ctx context.Context) float64 {
	v, err := pcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pcs *PriceCurrencySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pcs.fields) > 1 {
		return nil, errors.New("ent: PriceCurrencySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcs *PriceCurrencySelect) BoolsX(ctx context.Context) []bool {
	v, err := pcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pcs *PriceCurrencySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{pricecurrency.Label}
	default:
		err = fmt.Errorf("ent: PriceCurrencySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcs *PriceCurrencySelect) BoolX(ctx context.Context) bool {
	v, err := pcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcs *PriceCurrencySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pcs.sql.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
