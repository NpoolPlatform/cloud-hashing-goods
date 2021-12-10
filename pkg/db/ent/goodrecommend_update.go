// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodrecommend"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodRecommendUpdate is the builder for updating GoodRecommend entities.
type GoodRecommendUpdate struct {
	config
	hooks    []Hook
	mutation *GoodRecommendMutation
}

// Where appends a list predicates to the GoodRecommendUpdate builder.
func (gru *GoodRecommendUpdate) Where(ps ...predicate.GoodRecommend) *GoodRecommendUpdate {
	gru.mutation.Where(ps...)
	return gru
}

// SetUserID sets the "user_id" field.
func (gru *GoodRecommendUpdate) SetUserID(u uuid.UUID) *GoodRecommendUpdate {
	gru.mutation.SetUserID(u)
	return gru
}

// SetGoodID sets the "good_id" field.
func (gru *GoodRecommendUpdate) SetGoodID(u uuid.UUID) *GoodRecommendUpdate {
	gru.mutation.SetGoodID(u)
	return gru
}

// SetContent sets the "content" field.
func (gru *GoodRecommendUpdate) SetContent(s string) *GoodRecommendUpdate {
	gru.mutation.SetContent(s)
	return gru
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (gru *GoodRecommendUpdate) SetNillableContent(s *string) *GoodRecommendUpdate {
	if s != nil {
		gru.SetContent(*s)
	}
	return gru
}

// SetCreateAt sets the "create_at" field.
func (gru *GoodRecommendUpdate) SetCreateAt(i int64) *GoodRecommendUpdate {
	gru.mutation.ResetCreateAt()
	gru.mutation.SetCreateAt(i)
	return gru
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (gru *GoodRecommendUpdate) SetNillableCreateAt(i *int64) *GoodRecommendUpdate {
	if i != nil {
		gru.SetCreateAt(*i)
	}
	return gru
}

// AddCreateAt adds i to the "create_at" field.
func (gru *GoodRecommendUpdate) AddCreateAt(i int64) *GoodRecommendUpdate {
	gru.mutation.AddCreateAt(i)
	return gru
}

// SetUpdateAt sets the "update_at" field.
func (gru *GoodRecommendUpdate) SetUpdateAt(i int64) *GoodRecommendUpdate {
	gru.mutation.ResetUpdateAt()
	gru.mutation.SetUpdateAt(i)
	return gru
}

// AddUpdateAt adds i to the "update_at" field.
func (gru *GoodRecommendUpdate) AddUpdateAt(i int64) *GoodRecommendUpdate {
	gru.mutation.AddUpdateAt(i)
	return gru
}

// SetDeleteAt sets the "delete_at" field.
func (gru *GoodRecommendUpdate) SetDeleteAt(i int64) *GoodRecommendUpdate {
	gru.mutation.ResetDeleteAt()
	gru.mutation.SetDeleteAt(i)
	return gru
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (gru *GoodRecommendUpdate) SetNillableDeleteAt(i *int64) *GoodRecommendUpdate {
	if i != nil {
		gru.SetDeleteAt(*i)
	}
	return gru
}

// AddDeleteAt adds i to the "delete_at" field.
func (gru *GoodRecommendUpdate) AddDeleteAt(i int64) *GoodRecommendUpdate {
	gru.mutation.AddDeleteAt(i)
	return gru
}

// Mutation returns the GoodRecommendMutation object of the builder.
func (gru *GoodRecommendUpdate) Mutation() *GoodRecommendMutation {
	return gru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gru *GoodRecommendUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gru.defaults()
	if len(gru.hooks) == 0 {
		affected, err = gru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodRecommendMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gru.mutation = mutation
			affected, err = gru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gru.hooks) - 1; i >= 0; i-- {
			if gru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gru *GoodRecommendUpdate) SaveX(ctx context.Context) int {
	affected, err := gru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gru *GoodRecommendUpdate) Exec(ctx context.Context) error {
	_, err := gru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gru *GoodRecommendUpdate) ExecX(ctx context.Context) {
	if err := gru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gru *GoodRecommendUpdate) defaults() {
	if _, ok := gru.mutation.UpdateAt(); !ok {
		v := goodrecommend.UpdateDefaultUpdateAt()
		gru.mutation.SetUpdateAt(v)
	}
}

func (gru *GoodRecommendUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodrecommend.Table,
			Columns: goodrecommend.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodrecommend.FieldID,
			},
		},
	}
	if ps := gru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gru.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrecommend.FieldUserID,
		})
	}
	if value, ok := gru.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrecommend.FieldGoodID,
		})
	}
	if value, ok := gru.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodrecommend.FieldContent,
		})
	}
	if value, ok := gru.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldCreateAt,
		})
	}
	if value, ok := gru.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldCreateAt,
		})
	}
	if value, ok := gru.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldUpdateAt,
		})
	}
	if value, ok := gru.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldUpdateAt,
		})
	}
	if value, ok := gru.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldDeleteAt,
		})
	}
	if value, ok := gru.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodrecommend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// GoodRecommendUpdateOne is the builder for updating a single GoodRecommend entity.
type GoodRecommendUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoodRecommendMutation
}

// SetUserID sets the "user_id" field.
func (gruo *GoodRecommendUpdateOne) SetUserID(u uuid.UUID) *GoodRecommendUpdateOne {
	gruo.mutation.SetUserID(u)
	return gruo
}

// SetGoodID sets the "good_id" field.
func (gruo *GoodRecommendUpdateOne) SetGoodID(u uuid.UUID) *GoodRecommendUpdateOne {
	gruo.mutation.SetGoodID(u)
	return gruo
}

// SetContent sets the "content" field.
func (gruo *GoodRecommendUpdateOne) SetContent(s string) *GoodRecommendUpdateOne {
	gruo.mutation.SetContent(s)
	return gruo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (gruo *GoodRecommendUpdateOne) SetNillableContent(s *string) *GoodRecommendUpdateOne {
	if s != nil {
		gruo.SetContent(*s)
	}
	return gruo
}

// SetCreateAt sets the "create_at" field.
func (gruo *GoodRecommendUpdateOne) SetCreateAt(i int64) *GoodRecommendUpdateOne {
	gruo.mutation.ResetCreateAt()
	gruo.mutation.SetCreateAt(i)
	return gruo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (gruo *GoodRecommendUpdateOne) SetNillableCreateAt(i *int64) *GoodRecommendUpdateOne {
	if i != nil {
		gruo.SetCreateAt(*i)
	}
	return gruo
}

// AddCreateAt adds i to the "create_at" field.
func (gruo *GoodRecommendUpdateOne) AddCreateAt(i int64) *GoodRecommendUpdateOne {
	gruo.mutation.AddCreateAt(i)
	return gruo
}

// SetUpdateAt sets the "update_at" field.
func (gruo *GoodRecommendUpdateOne) SetUpdateAt(i int64) *GoodRecommendUpdateOne {
	gruo.mutation.ResetUpdateAt()
	gruo.mutation.SetUpdateAt(i)
	return gruo
}

// AddUpdateAt adds i to the "update_at" field.
func (gruo *GoodRecommendUpdateOne) AddUpdateAt(i int64) *GoodRecommendUpdateOne {
	gruo.mutation.AddUpdateAt(i)
	return gruo
}

// SetDeleteAt sets the "delete_at" field.
func (gruo *GoodRecommendUpdateOne) SetDeleteAt(i int64) *GoodRecommendUpdateOne {
	gruo.mutation.ResetDeleteAt()
	gruo.mutation.SetDeleteAt(i)
	return gruo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (gruo *GoodRecommendUpdateOne) SetNillableDeleteAt(i *int64) *GoodRecommendUpdateOne {
	if i != nil {
		gruo.SetDeleteAt(*i)
	}
	return gruo
}

// AddDeleteAt adds i to the "delete_at" field.
func (gruo *GoodRecommendUpdateOne) AddDeleteAt(i int64) *GoodRecommendUpdateOne {
	gruo.mutation.AddDeleteAt(i)
	return gruo
}

// Mutation returns the GoodRecommendMutation object of the builder.
func (gruo *GoodRecommendUpdateOne) Mutation() *GoodRecommendMutation {
	return gruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gruo *GoodRecommendUpdateOne) Select(field string, fields ...string) *GoodRecommendUpdateOne {
	gruo.fields = append([]string{field}, fields...)
	return gruo
}

// Save executes the query and returns the updated GoodRecommend entity.
func (gruo *GoodRecommendUpdateOne) Save(ctx context.Context) (*GoodRecommend, error) {
	var (
		err  error
		node *GoodRecommend
	)
	gruo.defaults()
	if len(gruo.hooks) == 0 {
		node, err = gruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodRecommendMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gruo.mutation = mutation
			node, err = gruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gruo.hooks) - 1; i >= 0; i-- {
			if gruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gruo *GoodRecommendUpdateOne) SaveX(ctx context.Context) *GoodRecommend {
	node, err := gruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gruo *GoodRecommendUpdateOne) Exec(ctx context.Context) error {
	_, err := gruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gruo *GoodRecommendUpdateOne) ExecX(ctx context.Context) {
	if err := gruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gruo *GoodRecommendUpdateOne) defaults() {
	if _, ok := gruo.mutation.UpdateAt(); !ok {
		v := goodrecommend.UpdateDefaultUpdateAt()
		gruo.mutation.SetUpdateAt(v)
	}
}

func (gruo *GoodRecommendUpdateOne) sqlSave(ctx context.Context) (_node *GoodRecommend, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodrecommend.Table,
			Columns: goodrecommend.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodrecommend.FieldID,
			},
		},
	}
	id, ok := gruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GoodRecommend.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := gruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodrecommend.FieldID)
		for _, f := range fields {
			if !goodrecommend.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodrecommend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gruo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrecommend.FieldUserID,
		})
	}
	if value, ok := gruo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrecommend.FieldGoodID,
		})
	}
	if value, ok := gruo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodrecommend.FieldContent,
		})
	}
	if value, ok := gruo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldCreateAt,
		})
	}
	if value, ok := gruo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldCreateAt,
		})
	}
	if value, ok := gruo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldUpdateAt,
		})
	}
	if value, ok := gruo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldUpdateAt,
		})
	}
	if value, ok := gruo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldDeleteAt,
		})
	}
	if value, ok := gruo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodrecommend.FieldDeleteAt,
		})
	}
	_node = &GoodRecommend{config: gruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodrecommend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
