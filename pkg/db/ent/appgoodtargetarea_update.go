// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgoodtargetarea"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppGoodTargetAreaUpdate is the builder for updating AppGoodTargetArea entities.
type AppGoodTargetAreaUpdate struct {
	config
	hooks    []Hook
	mutation *AppGoodTargetAreaMutation
}

// Where appends a list predicates to the AppGoodTargetAreaUpdate builder.
func (agtau *AppGoodTargetAreaUpdate) Where(ps ...predicate.AppGoodTargetArea) *AppGoodTargetAreaUpdate {
	agtau.mutation.Where(ps...)
	return agtau
}

// SetAppID sets the "app_id" field.
func (agtau *AppGoodTargetAreaUpdate) SetAppID(u uuid.UUID) *AppGoodTargetAreaUpdate {
	agtau.mutation.SetAppID(u)
	return agtau
}

// SetGoodID sets the "good_id" field.
func (agtau *AppGoodTargetAreaUpdate) SetGoodID(u uuid.UUID) *AppGoodTargetAreaUpdate {
	agtau.mutation.SetGoodID(u)
	return agtau
}

// SetTargetAreaID sets the "target_area_id" field.
func (agtau *AppGoodTargetAreaUpdate) SetTargetAreaID(u uuid.UUID) *AppGoodTargetAreaUpdate {
	agtau.mutation.SetTargetAreaID(u)
	return agtau
}

// SetCreateAt sets the "create_at" field.
func (agtau *AppGoodTargetAreaUpdate) SetCreateAt(i int64) *AppGoodTargetAreaUpdate {
	agtau.mutation.ResetCreateAt()
	agtau.mutation.SetCreateAt(i)
	return agtau
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (agtau *AppGoodTargetAreaUpdate) SetNillableCreateAt(i *int64) *AppGoodTargetAreaUpdate {
	if i != nil {
		agtau.SetCreateAt(*i)
	}
	return agtau
}

// AddCreateAt adds i to the "create_at" field.
func (agtau *AppGoodTargetAreaUpdate) AddCreateAt(i int64) *AppGoodTargetAreaUpdate {
	agtau.mutation.AddCreateAt(i)
	return agtau
}

// SetUpdateAt sets the "update_at" field.
func (agtau *AppGoodTargetAreaUpdate) SetUpdateAt(i int64) *AppGoodTargetAreaUpdate {
	agtau.mutation.ResetUpdateAt()
	agtau.mutation.SetUpdateAt(i)
	return agtau
}

// AddUpdateAt adds i to the "update_at" field.
func (agtau *AppGoodTargetAreaUpdate) AddUpdateAt(i int64) *AppGoodTargetAreaUpdate {
	agtau.mutation.AddUpdateAt(i)
	return agtau
}

// SetDeleteAt sets the "delete_at" field.
func (agtau *AppGoodTargetAreaUpdate) SetDeleteAt(i int64) *AppGoodTargetAreaUpdate {
	agtau.mutation.ResetDeleteAt()
	agtau.mutation.SetDeleteAt(i)
	return agtau
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (agtau *AppGoodTargetAreaUpdate) SetNillableDeleteAt(i *int64) *AppGoodTargetAreaUpdate {
	if i != nil {
		agtau.SetDeleteAt(*i)
	}
	return agtau
}

// AddDeleteAt adds i to the "delete_at" field.
func (agtau *AppGoodTargetAreaUpdate) AddDeleteAt(i int64) *AppGoodTargetAreaUpdate {
	agtau.mutation.AddDeleteAt(i)
	return agtau
}

// Mutation returns the AppGoodTargetAreaMutation object of the builder.
func (agtau *AppGoodTargetAreaUpdate) Mutation() *AppGoodTargetAreaMutation {
	return agtau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (agtau *AppGoodTargetAreaUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	agtau.defaults()
	if len(agtau.hooks) == 0 {
		affected, err = agtau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodTargetAreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agtau.mutation = mutation
			affected, err = agtau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agtau.hooks) - 1; i >= 0; i-- {
			if agtau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agtau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agtau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (agtau *AppGoodTargetAreaUpdate) SaveX(ctx context.Context) int {
	affected, err := agtau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (agtau *AppGoodTargetAreaUpdate) Exec(ctx context.Context) error {
	_, err := agtau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agtau *AppGoodTargetAreaUpdate) ExecX(ctx context.Context) {
	if err := agtau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agtau *AppGoodTargetAreaUpdate) defaults() {
	if _, ok := agtau.mutation.UpdateAt(); !ok {
		v := appgoodtargetarea.UpdateDefaultUpdateAt()
		agtau.mutation.SetUpdateAt(v)
	}
}

func (agtau *AppGoodTargetAreaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgoodtargetarea.Table,
			Columns: appgoodtargetarea.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appgoodtargetarea.FieldID,
			},
		},
	}
	if ps := agtau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := agtau.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgoodtargetarea.FieldAppID,
		})
	}
	if value, ok := agtau.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgoodtargetarea.FieldGoodID,
		})
	}
	if value, ok := agtau.mutation.TargetAreaID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgoodtargetarea.FieldTargetAreaID,
		})
	}
	if value, ok := agtau.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldCreateAt,
		})
	}
	if value, ok := agtau.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldCreateAt,
		})
	}
	if value, ok := agtau.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldUpdateAt,
		})
	}
	if value, ok := agtau.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldUpdateAt,
		})
	}
	if value, ok := agtau.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldDeleteAt,
		})
	}
	if value, ok := agtau.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, agtau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appgoodtargetarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppGoodTargetAreaUpdateOne is the builder for updating a single AppGoodTargetArea entity.
type AppGoodTargetAreaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppGoodTargetAreaMutation
}

// SetAppID sets the "app_id" field.
func (agtauo *AppGoodTargetAreaUpdateOne) SetAppID(u uuid.UUID) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.SetAppID(u)
	return agtauo
}

// SetGoodID sets the "good_id" field.
func (agtauo *AppGoodTargetAreaUpdateOne) SetGoodID(u uuid.UUID) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.SetGoodID(u)
	return agtauo
}

// SetTargetAreaID sets the "target_area_id" field.
func (agtauo *AppGoodTargetAreaUpdateOne) SetTargetAreaID(u uuid.UUID) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.SetTargetAreaID(u)
	return agtauo
}

// SetCreateAt sets the "create_at" field.
func (agtauo *AppGoodTargetAreaUpdateOne) SetCreateAt(i int64) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.ResetCreateAt()
	agtauo.mutation.SetCreateAt(i)
	return agtauo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (agtauo *AppGoodTargetAreaUpdateOne) SetNillableCreateAt(i *int64) *AppGoodTargetAreaUpdateOne {
	if i != nil {
		agtauo.SetCreateAt(*i)
	}
	return agtauo
}

// AddCreateAt adds i to the "create_at" field.
func (agtauo *AppGoodTargetAreaUpdateOne) AddCreateAt(i int64) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.AddCreateAt(i)
	return agtauo
}

// SetUpdateAt sets the "update_at" field.
func (agtauo *AppGoodTargetAreaUpdateOne) SetUpdateAt(i int64) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.ResetUpdateAt()
	agtauo.mutation.SetUpdateAt(i)
	return agtauo
}

// AddUpdateAt adds i to the "update_at" field.
func (agtauo *AppGoodTargetAreaUpdateOne) AddUpdateAt(i int64) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.AddUpdateAt(i)
	return agtauo
}

// SetDeleteAt sets the "delete_at" field.
func (agtauo *AppGoodTargetAreaUpdateOne) SetDeleteAt(i int64) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.ResetDeleteAt()
	agtauo.mutation.SetDeleteAt(i)
	return agtauo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (agtauo *AppGoodTargetAreaUpdateOne) SetNillableDeleteAt(i *int64) *AppGoodTargetAreaUpdateOne {
	if i != nil {
		agtauo.SetDeleteAt(*i)
	}
	return agtauo
}

// AddDeleteAt adds i to the "delete_at" field.
func (agtauo *AppGoodTargetAreaUpdateOne) AddDeleteAt(i int64) *AppGoodTargetAreaUpdateOne {
	agtauo.mutation.AddDeleteAt(i)
	return agtauo
}

// Mutation returns the AppGoodTargetAreaMutation object of the builder.
func (agtauo *AppGoodTargetAreaUpdateOne) Mutation() *AppGoodTargetAreaMutation {
	return agtauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (agtauo *AppGoodTargetAreaUpdateOne) Select(field string, fields ...string) *AppGoodTargetAreaUpdateOne {
	agtauo.fields = append([]string{field}, fields...)
	return agtauo
}

// Save executes the query and returns the updated AppGoodTargetArea entity.
func (agtauo *AppGoodTargetAreaUpdateOne) Save(ctx context.Context) (*AppGoodTargetArea, error) {
	var (
		err  error
		node *AppGoodTargetArea
	)
	agtauo.defaults()
	if len(agtauo.hooks) == 0 {
		node, err = agtauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodTargetAreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agtauo.mutation = mutation
			node, err = agtauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(agtauo.hooks) - 1; i >= 0; i-- {
			if agtauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agtauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agtauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (agtauo *AppGoodTargetAreaUpdateOne) SaveX(ctx context.Context) *AppGoodTargetArea {
	node, err := agtauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (agtauo *AppGoodTargetAreaUpdateOne) Exec(ctx context.Context) error {
	_, err := agtauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agtauo *AppGoodTargetAreaUpdateOne) ExecX(ctx context.Context) {
	if err := agtauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agtauo *AppGoodTargetAreaUpdateOne) defaults() {
	if _, ok := agtauo.mutation.UpdateAt(); !ok {
		v := appgoodtargetarea.UpdateDefaultUpdateAt()
		agtauo.mutation.SetUpdateAt(v)
	}
}

func (agtauo *AppGoodTargetAreaUpdateOne) sqlSave(ctx context.Context) (_node *AppGoodTargetArea, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgoodtargetarea.Table,
			Columns: appgoodtargetarea.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appgoodtargetarea.FieldID,
			},
		},
	}
	id, ok := agtauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AppGoodTargetArea.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := agtauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appgoodtargetarea.FieldID)
		for _, f := range fields {
			if !appgoodtargetarea.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appgoodtargetarea.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := agtauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := agtauo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgoodtargetarea.FieldAppID,
		})
	}
	if value, ok := agtauo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgoodtargetarea.FieldGoodID,
		})
	}
	if value, ok := agtauo.mutation.TargetAreaID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgoodtargetarea.FieldTargetAreaID,
		})
	}
	if value, ok := agtauo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldCreateAt,
		})
	}
	if value, ok := agtauo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldCreateAt,
		})
	}
	if value, ok := agtauo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldUpdateAt,
		})
	}
	if value, ok := agtauo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldUpdateAt,
		})
	}
	if value, ok := agtauo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldDeleteAt,
		})
	}
	if value, ok := agtauo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: appgoodtargetarea.FieldDeleteAt,
		})
	}
	_node = &AppGoodTargetArea{config: agtauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, agtauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appgoodtargetarea.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
