// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"
)

// TargetAreaUpdate is the builder for updating TargetArea entities.
type TargetAreaUpdate struct {
	config
	hooks    []Hook
	mutation *TargetAreaMutation
}

// Where appends a list predicates to the TargetAreaUpdate builder.
func (tau *TargetAreaUpdate) Where(ps ...predicate.TargetArea) *TargetAreaUpdate {
	tau.mutation.Where(ps...)
	return tau
}

// SetContinent sets the "continent" field.
func (tau *TargetAreaUpdate) SetContinent(s string) *TargetAreaUpdate {
	tau.mutation.SetContinent(s)
	return tau
}

// SetCountry sets the "country" field.
func (tau *TargetAreaUpdate) SetCountry(s string) *TargetAreaUpdate {
	tau.mutation.SetCountry(s)
	return tau
}

// SetCreateAt sets the "create_at" field.
func (tau *TargetAreaUpdate) SetCreateAt(t time.Time) *TargetAreaUpdate {
	tau.mutation.SetCreateAt(t)
	return tau
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (tau *TargetAreaUpdate) SetNillableCreateAt(t *time.Time) *TargetAreaUpdate {
	if t != nil {
		tau.SetCreateAt(*t)
	}
	return tau
}

// SetUpdateAt sets the "update_at" field.
func (tau *TargetAreaUpdate) SetUpdateAt(t time.Time) *TargetAreaUpdate {
	tau.mutation.SetUpdateAt(t)
	return tau
}

// SetDeleteAt sets the "delete_at" field.
func (tau *TargetAreaUpdate) SetDeleteAt(t time.Time) *TargetAreaUpdate {
	tau.mutation.SetDeleteAt(t)
	return tau
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tau *TargetAreaUpdate) SetNillableDeleteAt(t *time.Time) *TargetAreaUpdate {
	if t != nil {
		tau.SetDeleteAt(*t)
	}
	return tau
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (tau *TargetAreaUpdate) ClearDeleteAt() *TargetAreaUpdate {
	tau.mutation.ClearDeleteAt()
	return tau
}

// Mutation returns the TargetAreaMutation object of the builder.
func (tau *TargetAreaUpdate) Mutation() *TargetAreaMutation {
	return tau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tau *TargetAreaUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tau.defaults()
	if len(tau.hooks) == 0 {
		if err = tau.check(); err != nil {
			return 0, err
		}
		affected, err = tau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetAreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tau.check(); err != nil {
				return 0, err
			}
			tau.mutation = mutation
			affected, err = tau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tau.hooks) - 1; i >= 0; i-- {
			if tau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tau *TargetAreaUpdate) SaveX(ctx context.Context) int {
	affected, err := tau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tau *TargetAreaUpdate) Exec(ctx context.Context) error {
	_, err := tau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tau *TargetAreaUpdate) ExecX(ctx context.Context) {
	if err := tau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tau *TargetAreaUpdate) defaults() {
	if _, ok := tau.mutation.UpdateAt(); !ok {
		v := targetarea.UpdateDefaultUpdateAt()
		tau.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tau *TargetAreaUpdate) check() error {
	if v, ok := tau.mutation.Continent(); ok {
		if err := targetarea.ContinentValidator(v); err != nil {
			return &ValidationError{Name: "continent", err: fmt.Errorf("ent: validator failed for field \"continent\": %w", err)}
		}
	}
	if v, ok := tau.mutation.Country(); ok {
		if err := targetarea.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf("ent: validator failed for field \"country\": %w", err)}
		}
	}
	return nil
}

func (tau *TargetAreaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   targetarea.Table,
			Columns: targetarea.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: targetarea.FieldID,
			},
		},
	}
	if ps := tau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tau.mutation.Continent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: targetarea.FieldContinent,
		})
	}
	if value, ok := tau.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: targetarea.FieldCountry,
		})
	}
	if value, ok := tau.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: targetarea.FieldCreateAt,
		})
	}
	if value, ok := tau.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: targetarea.FieldUpdateAt,
		})
	}
	if value, ok := tau.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: targetarea.FieldDeleteAt,
		})
	}
	if tau.mutation.DeleteAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: targetarea.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{targetarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TargetAreaUpdateOne is the builder for updating a single TargetArea entity.
type TargetAreaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TargetAreaMutation
}

// SetContinent sets the "continent" field.
func (tauo *TargetAreaUpdateOne) SetContinent(s string) *TargetAreaUpdateOne {
	tauo.mutation.SetContinent(s)
	return tauo
}

// SetCountry sets the "country" field.
func (tauo *TargetAreaUpdateOne) SetCountry(s string) *TargetAreaUpdateOne {
	tauo.mutation.SetCountry(s)
	return tauo
}

// SetCreateAt sets the "create_at" field.
func (tauo *TargetAreaUpdateOne) SetCreateAt(t time.Time) *TargetAreaUpdateOne {
	tauo.mutation.SetCreateAt(t)
	return tauo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (tauo *TargetAreaUpdateOne) SetNillableCreateAt(t *time.Time) *TargetAreaUpdateOne {
	if t != nil {
		tauo.SetCreateAt(*t)
	}
	return tauo
}

// SetUpdateAt sets the "update_at" field.
func (tauo *TargetAreaUpdateOne) SetUpdateAt(t time.Time) *TargetAreaUpdateOne {
	tauo.mutation.SetUpdateAt(t)
	return tauo
}

// SetDeleteAt sets the "delete_at" field.
func (tauo *TargetAreaUpdateOne) SetDeleteAt(t time.Time) *TargetAreaUpdateOne {
	tauo.mutation.SetDeleteAt(t)
	return tauo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tauo *TargetAreaUpdateOne) SetNillableDeleteAt(t *time.Time) *TargetAreaUpdateOne {
	if t != nil {
		tauo.SetDeleteAt(*t)
	}
	return tauo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (tauo *TargetAreaUpdateOne) ClearDeleteAt() *TargetAreaUpdateOne {
	tauo.mutation.ClearDeleteAt()
	return tauo
}

// Mutation returns the TargetAreaMutation object of the builder.
func (tauo *TargetAreaUpdateOne) Mutation() *TargetAreaMutation {
	return tauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tauo *TargetAreaUpdateOne) Select(field string, fields ...string) *TargetAreaUpdateOne {
	tauo.fields = append([]string{field}, fields...)
	return tauo
}

// Save executes the query and returns the updated TargetArea entity.
func (tauo *TargetAreaUpdateOne) Save(ctx context.Context) (*TargetArea, error) {
	var (
		err  error
		node *TargetArea
	)
	tauo.defaults()
	if len(tauo.hooks) == 0 {
		if err = tauo.check(); err != nil {
			return nil, err
		}
		node, err = tauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetAreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tauo.check(); err != nil {
				return nil, err
			}
			tauo.mutation = mutation
			node, err = tauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tauo.hooks) - 1; i >= 0; i-- {
			if tauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tauo *TargetAreaUpdateOne) SaveX(ctx context.Context) *TargetArea {
	node, err := tauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tauo *TargetAreaUpdateOne) Exec(ctx context.Context) error {
	_, err := tauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tauo *TargetAreaUpdateOne) ExecX(ctx context.Context) {
	if err := tauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tauo *TargetAreaUpdateOne) defaults() {
	if _, ok := tauo.mutation.UpdateAt(); !ok {
		v := targetarea.UpdateDefaultUpdateAt()
		tauo.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tauo *TargetAreaUpdateOne) check() error {
	if v, ok := tauo.mutation.Continent(); ok {
		if err := targetarea.ContinentValidator(v); err != nil {
			return &ValidationError{Name: "continent", err: fmt.Errorf("ent: validator failed for field \"continent\": %w", err)}
		}
	}
	if v, ok := tauo.mutation.Country(); ok {
		if err := targetarea.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf("ent: validator failed for field \"country\": %w", err)}
		}
	}
	return nil
}

func (tauo *TargetAreaUpdateOne) sqlSave(ctx context.Context) (_node *TargetArea, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   targetarea.Table,
			Columns: targetarea.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: targetarea.FieldID,
			},
		},
	}
	id, ok := tauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing TargetArea.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, targetarea.FieldID)
		for _, f := range fields {
			if !targetarea.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != targetarea.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tauo.mutation.Continent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: targetarea.FieldContinent,
		})
	}
	if value, ok := tauo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: targetarea.FieldCountry,
		})
	}
	if value, ok := tauo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: targetarea.FieldCreateAt,
		})
	}
	if value, ok := tauo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: targetarea.FieldUpdateAt,
		})
	}
	if value, ok := tauo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: targetarea.FieldDeleteAt,
		})
	}
	if tauo.mutation.DeleteAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: targetarea.FieldDeleteAt,
		})
	}
	_node = &TargetArea{config: tauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{targetarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
