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
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodcomment"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodCommentQuery is the builder for querying GoodComment entities.
type GoodCommentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GoodComment
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodCommentQuery builder.
func (gcq *GoodCommentQuery) Where(ps ...predicate.GoodComment) *GoodCommentQuery {
	gcq.predicates = append(gcq.predicates, ps...)
	return gcq
}

// Limit adds a limit step to the query.
func (gcq *GoodCommentQuery) Limit(limit int) *GoodCommentQuery {
	gcq.limit = &limit
	return gcq
}

// Offset adds an offset step to the query.
func (gcq *GoodCommentQuery) Offset(offset int) *GoodCommentQuery {
	gcq.offset = &offset
	return gcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gcq *GoodCommentQuery) Unique(unique bool) *GoodCommentQuery {
	gcq.unique = &unique
	return gcq
}

// Order adds an order step to the query.
func (gcq *GoodCommentQuery) Order(o ...OrderFunc) *GoodCommentQuery {
	gcq.order = append(gcq.order, o...)
	return gcq
}

// First returns the first GoodComment entity from the query.
// Returns a *NotFoundError when no GoodComment was found.
func (gcq *GoodCommentQuery) First(ctx context.Context) (*GoodComment, error) {
	nodes, err := gcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goodcomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gcq *GoodCommentQuery) FirstX(ctx context.Context) *GoodComment {
	node, err := gcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GoodComment ID from the query.
// Returns a *NotFoundError when no GoodComment ID was found.
func (gcq *GoodCommentQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goodcomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gcq *GoodCommentQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GoodComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GoodComment entity is not found.
// Returns a *NotFoundError when no GoodComment entities are found.
func (gcq *GoodCommentQuery) Only(ctx context.Context) (*GoodComment, error) {
	nodes, err := gcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goodcomment.Label}
	default:
		return nil, &NotSingularError{goodcomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gcq *GoodCommentQuery) OnlyX(ctx context.Context) *GoodComment {
	node, err := gcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GoodComment ID in the query.
// Returns a *NotSingularError when exactly one GoodComment ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gcq *GoodCommentQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = &NotSingularError{goodcomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gcq *GoodCommentQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodComments.
func (gcq *GoodCommentQuery) All(ctx context.Context) ([]*GoodComment, error) {
	if err := gcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gcq *GoodCommentQuery) AllX(ctx context.Context) []*GoodComment {
	nodes, err := gcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GoodComment IDs.
func (gcq *GoodCommentQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gcq.Select(goodcomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gcq *GoodCommentQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gcq *GoodCommentQuery) Count(ctx context.Context) (int, error) {
	if err := gcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gcq *GoodCommentQuery) CountX(ctx context.Context) int {
	count, err := gcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gcq *GoodCommentQuery) Exist(ctx context.Context) (bool, error) {
	if err := gcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gcq *GoodCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := gcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gcq *GoodCommentQuery) Clone() *GoodCommentQuery {
	if gcq == nil {
		return nil
	}
	return &GoodCommentQuery{
		config:     gcq.config,
		limit:      gcq.limit,
		offset:     gcq.offset,
		order:      append([]OrderFunc{}, gcq.order...),
		predicates: append([]predicate.GoodComment{}, gcq.predicates...),
		// clone intermediate query.
		sql:  gcq.sql.Clone(),
		path: gcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ReplyToID uuid.UUID `json:"reply_to_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GoodComment.Query().
//		GroupBy(goodcomment.FieldReplyToID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gcq *GoodCommentQuery) GroupBy(field string, fields ...string) *GoodCommentGroupBy {
	group := &GoodCommentGroupBy{config: gcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ReplyToID uuid.UUID `json:"reply_to_id,omitempty"`
//	}
//
//	client.GoodComment.Query().
//		Select(goodcomment.FieldReplyToID).
//		Scan(ctx, &v)
//
func (gcq *GoodCommentQuery) Select(fields ...string) *GoodCommentSelect {
	gcq.fields = append(gcq.fields, fields...)
	return &GoodCommentSelect{GoodCommentQuery: gcq}
}

func (gcq *GoodCommentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gcq.fields {
		if !goodcomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gcq.path != nil {
		prev, err := gcq.path(ctx)
		if err != nil {
			return err
		}
		gcq.sql = prev
	}
	return nil
}

func (gcq *GoodCommentQuery) sqlAll(ctx context.Context) ([]*GoodComment, error) {
	var (
		nodes = []*GoodComment{}
		_spec = gcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GoodComment{config: gcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gcq *GoodCommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gcq.querySpec()
	_spec.Node.Columns = gcq.fields
	if len(gcq.fields) > 0 {
		_spec.Unique = gcq.unique != nil && *gcq.unique
	}
	return sqlgraph.CountNodes(ctx, gcq.driver, _spec)
}

func (gcq *GoodCommentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gcq *GoodCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodcomment.Table,
			Columns: goodcomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodcomment.FieldID,
			},
		},
		From:   gcq.sql,
		Unique: true,
	}
	if unique := gcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodcomment.FieldID)
		for i := range fields {
			if fields[i] != goodcomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gcq *GoodCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gcq.driver.Dialect())
	t1 := builder.Table(goodcomment.Table)
	columns := gcq.fields
	if len(columns) == 0 {
		columns = goodcomment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gcq.sql != nil {
		selector = gcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gcq.unique != nil && *gcq.unique {
		selector.Distinct()
	}
	for _, p := range gcq.predicates {
		p(selector)
	}
	for _, p := range gcq.order {
		p(selector)
	}
	if offset := gcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GoodCommentGroupBy is the group-by builder for GoodComment entities.
type GoodCommentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gcgb *GoodCommentGroupBy) Aggregate(fns ...AggregateFunc) *GoodCommentGroupBy {
	gcgb.fns = append(gcgb.fns, fns...)
	return gcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gcgb *GoodCommentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gcgb.path(ctx)
	if err != nil {
		return err
	}
	gcgb.sql = query
	return gcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GoodCommentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("ent: GoodCommentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) StringsX(ctx context.Context) []string {
	v, err := gcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GoodCommentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = fmt.Errorf("ent: GoodCommentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) StringX(ctx context.Context) string {
	v, err := gcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GoodCommentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("ent: GoodCommentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) IntsX(ctx context.Context) []int {
	v, err := gcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GoodCommentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = fmt.Errorf("ent: GoodCommentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) IntX(ctx context.Context) int {
	v, err := gcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GoodCommentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("ent: GoodCommentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GoodCommentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = fmt.Errorf("ent: GoodCommentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GoodCommentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("ent: GoodCommentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GoodCommentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = fmt.Errorf("ent: GoodCommentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gcgb *GoodCommentGroupBy) BoolX(ctx context.Context) bool {
	v, err := gcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gcgb *GoodCommentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gcgb.fields {
		if !goodcomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gcgb *GoodCommentGroupBy) sqlQuery() *sql.Selector {
	selector := gcgb.sql.Select()
	aggregation := make([]string, 0, len(gcgb.fns))
	for _, fn := range gcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gcgb.fields)+len(gcgb.fns))
		for _, f := range gcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gcgb.fields...)...)
}

// GoodCommentSelect is the builder for selecting fields of GoodComment entities.
type GoodCommentSelect struct {
	*GoodCommentQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gcs *GoodCommentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gcs.prepareQuery(ctx); err != nil {
		return err
	}
	gcs.sql = gcs.GoodCommentQuery.sqlQuery(ctx)
	return gcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gcs *GoodCommentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gcs *GoodCommentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("ent: GoodCommentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gcs *GoodCommentSelect) StringsX(ctx context.Context) []string {
	v, err := gcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gcs *GoodCommentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = fmt.Errorf("ent: GoodCommentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gcs *GoodCommentSelect) StringX(ctx context.Context) string {
	v, err := gcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gcs *GoodCommentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("ent: GoodCommentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gcs *GoodCommentSelect) IntsX(ctx context.Context) []int {
	v, err := gcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gcs *GoodCommentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = fmt.Errorf("ent: GoodCommentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gcs *GoodCommentSelect) IntX(ctx context.Context) int {
	v, err := gcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gcs *GoodCommentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("ent: GoodCommentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gcs *GoodCommentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gcs *GoodCommentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = fmt.Errorf("ent: GoodCommentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gcs *GoodCommentSelect) Float64X(ctx context.Context) float64 {
	v, err := gcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gcs *GoodCommentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("ent: GoodCommentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gcs *GoodCommentSelect) BoolsX(ctx context.Context) []bool {
	v, err := gcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gcs *GoodCommentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goodcomment.Label}
	default:
		err = fmt.Errorf("ent: GoodCommentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gcs *GoodCommentSelect) BoolX(ctx context.Context) bool {
	v, err := gcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gcs *GoodCommentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gcs.sql.Query()
	if err := gcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
