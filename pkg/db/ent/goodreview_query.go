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
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodreview"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodReviewQuery is the builder for querying GoodReview entities.
type GoodReviewQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodReview
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodReviewQuery builder.
func (grq *GoodReviewQuery) Where(ps ...predicate.GoodReview) *GoodReviewQuery {
	grq.predicates = append(grq.predicates, ps...)
	return grq
}

// Limit adds a limit step to the query.
func (grq *GoodReviewQuery) Limit(limit int) *GoodReviewQuery {
	grq.limit = &limit
	return grq
}

// Offset adds an offset step to the query.
func (grq *GoodReviewQuery) Offset(offset int) *GoodReviewQuery {
	grq.offset = &offset
	return grq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (grq *GoodReviewQuery) Unique(unique bool) *GoodReviewQuery {
	grq.unique = &unique
	return grq
}

// Order adds an order step to the query.
func (grq *GoodReviewQuery) Order(o ...OrderFunc) *GoodReviewQuery {
	grq.order = append(grq.order, o...)
	return grq
}

// First returns the first GoodReview entity from the query.
// Returns a *NotFoundError when no GoodReview was found.
func (grq *GoodReviewQuery) First(ctx context.Context) (*GoodReview, error) {
	nodes, err := grq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodreview.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (grq *GoodReviewQuery) FirstX(ctx context.Context) *GoodReview {
	node, err := grq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodReview ID from the query.
// Returns a *NotFoundError when no GoodReview ID was found.
func (grq *GoodReviewQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = grq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodreview.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (grq *GoodReviewQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := grq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodReview entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GoodReview entity is not found.
// Returns a *NotFoundError when no GoodReview entities are found.
func (grq *GoodReviewQuery) Only(ctx context.Context) (*GoodReview, error) {
	nodes, err := grq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodreview.Label}
	default:
		return nil, &NotSingularError{goodreview.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (grq *GoodReviewQuery) OnlyX(ctx context.Context) *GoodReview {
	node, err := grq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodReview ID in the query.
// Returns a *NotSingularError when exactly one GoodReview ID is not found.
// Returns a *NotFoundError when no entities are found.
func (grq *GoodReviewQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = grq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = &NotSingularError{goodreview.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (grq *GoodReviewQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := grq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodReviews.
func (grq *GoodReviewQuery) All(ctx context.Context) ([]*GoodReview, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return grq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (grq *GoodReviewQuery) AllX(ctx context.Context) []*GoodReview {
	nodes, err := grq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodReview IDs.
func (grq *GoodReviewQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := grq.Select(goodreview.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (grq *GoodReviewQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := grq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (grq *GoodReviewQuery) Count(ctx context.Context) (int, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return grq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (grq *GoodReviewQuery) CountX(ctx context.Context) int {
	count, err := grq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (grq *GoodReviewQuery) Exist(ctx context.Context) (bool, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return grq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (grq *GoodReviewQuery) ExistX(ctx context.Context) bool {
	exist, err := grq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodReviewQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (grq *GoodReviewQuery) Clone() *GoodReviewQuery {
	if grq == nil {
		return nil
	}
	return &GoodReviewQuery{
		config:     grq.config,
		limit:      grq.limit,
		offset:     grq.offset,
		order:      append([]OrderFunc{}, grq.order...),
		predicates: append([]predicate.GoodReview{}, grq.predicates...),
		// clone intermediate query.
		sql:  grq.sql.Clone(),
		path: grq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EntityType goodreview.EntityType `json:"entity_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoodReview.Query().
//		GroupBy(goodreview.FieldEntityType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (grq *GoodReviewQuery) GroupBy(field string, fields ...string) *GoodReviewGroupBy {
	group := &GoodReviewGroupBy{config: grq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return grq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EntityType goodreview.EntityType `json:"entity_type,omitempty"`
//	}
//
//	client.GoodReview.Query().
//		Select(goodreview.FieldEntityType).
//		Scan(ctx, &v)
//
func (grq *GoodReviewQuery) Select(fields ...string) *GoodReviewSelect {
	grq.fields = append(grq.fields, fields...)
	return &GoodReviewSelect{GoodReviewQuery: grq}
}

func (grq *GoodReviewQuery) prepareQuery(ctx context.Context) error {
	for _, f := range grq.fields {
		if !goodreview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if grq.path != nil {
		prev, err := grq.path(ctx)
		if err != nil {
			return err
		}
		grq.sql = prev
	}
	return nil
}

func (grq *GoodReviewQuery) sqlAll(ctx context.Context) ([]*GoodReview, error) {
	var (
		nodes = []*GoodReview{}
		_spec = grq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GoodReview{config: grq.config}
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
	if err := sqlgraph.QueryNodes(ctx, grq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (grq *GoodReviewQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := grq.querySpec()
	return sqlgraph.CountNodes(ctx, grq.driver, _spec)
}

func (grq *GoodReviewQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := grq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (grq *GoodReviewQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodreview.Table,
			Columns: goodreview.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodreview.FieldID,
			},
		},
		From:   grq.sql,
		Unique: true,
	}
	if unique := grq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := grq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodreview.FieldID)
		for i := range fields {
			if fields[i] != goodreview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := grq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := grq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := grq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := grq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (grq *GoodReviewQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(grq.driver.Dialect())
	t1 := builder.Table(goodreview.Table)
	columns := grq.fields
	if len(columns) == 0 {
		columns = goodreview.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if grq.sql != nil {
		selector = grq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range grq.predicates {
		p(selector)
	}
	for _, p := range grq.order {
		p(selector)
	}
	if offset := grq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := grq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodReviewGroupBy is the group-by builder for GoodReview entities.
type GoodReviewGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (grgb *GoodReviewGroupBy) Aggregate(fns ...AggregateFunc) *GoodReviewGroupBy {
	grgb.fns = append(grgb.fns, fns...)
	return grgb
}

// Scan applies the group-by query and scans the result into the given value.
func (grgb *GoodReviewGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := grgb.path(ctx)
	if err != nil {
		return err
	}
	grgb.sql = query
	return grgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := grgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (grgb *GoodReviewGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("ent: GoodReviewGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) StringsX(ctx context.Context) []string {
	v, err := grgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (grgb *GoodReviewGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = grgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = fmt.Errorf("ent: GoodReviewGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) StringX(ctx context.Context) string {
	v, err := grgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (grgb *GoodReviewGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("ent: GoodReviewGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) IntsX(ctx context.Context) []int {
	v, err := grgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (grgb *GoodReviewGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = grgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = fmt.Errorf("ent: GoodReviewGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) IntX(ctx context.Context) int {
	v, err := grgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (grgb *GoodReviewGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("ent: GoodReviewGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := grgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (grgb *GoodReviewGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = grgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = fmt.Errorf("ent: GoodReviewGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) Float64X(ctx context.Context) float64 {
	v, err := grgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (grgb *GoodReviewGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("ent: GoodReviewGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := grgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (grgb *GoodReviewGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = grgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = fmt.Errorf("ent: GoodReviewGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (grgb *GoodReviewGroupBy) BoolX(ctx context.Context) bool {
	v, err := grgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (grgb *GoodReviewGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range grgb.fields {
		if !goodreview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := grgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := grgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (grgb *GoodReviewGroupBy) sqlQuery() *sql.Selector {
	selector := grgb.sql.Select()
	aggregation := make([]string, 0, len(grgb.fns))
	for _, fn := range grgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(grgb.fields)+len(grgb.fns))
		for _, f := range grgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(grgb.fields...)...)
}

// GoodReviewSelect is the builder for selecting fields of GoodReview entities.
type GoodReviewSelect struct {
	*GoodReviewQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (grs *GoodReviewSelect) Scan(ctx context.Context, v interface{}) error {
	if err := grs.prepareQuery(ctx); err != nil {
		return err
	}
	grs.sql = grs.GoodReviewQuery.sqlQuery(ctx)
	return grs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (grs *GoodReviewSelect) ScanX(ctx context.Context, v interface{}) {
	if err := grs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (grs *GoodReviewSelect) Strings(ctx context.Context) ([]string, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("ent: GoodReviewSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (grs *GoodReviewSelect) StringsX(ctx context.Context) []string {
	v, err := grs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (grs *GoodReviewSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = grs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = fmt.Errorf("ent: GoodReviewSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (grs *GoodReviewSelect) StringX(ctx context.Context) string {
	v, err := grs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (grs *GoodReviewSelect) Ints(ctx context.Context) ([]int, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("ent: GoodReviewSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (grs *GoodReviewSelect) IntsX(ctx context.Context) []int {
	v, err := grs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (grs *GoodReviewSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = grs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = fmt.Errorf("ent: GoodReviewSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (grs *GoodReviewSelect) IntX(ctx context.Context) int {
	v, err := grs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (grs *GoodReviewSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("ent: GoodReviewSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (grs *GoodReviewSelect) Float64sX(ctx context.Context) []float64 {
	v, err := grs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (grs *GoodReviewSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = grs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = fmt.Errorf("ent: GoodReviewSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (grs *GoodReviewSelect) Float64X(ctx context.Context) float64 {
	v, err := grs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (grs *GoodReviewSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("ent: GoodReviewSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (grs *GoodReviewSelect) BoolsX(ctx context.Context) []bool {
	v, err := grs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (grs *GoodReviewSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = grs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodreview.Label}
	default:
		err = fmt.Errorf("ent: GoodReviewSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (grs *GoodReviewSelect) BoolX(ctx context.Context) bool {
	v, err := grs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (grs *GoodReviewSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := grs.sql.Query()
	if err := grs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
