// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"
	"github.com/google/uuid"
)

// TargetAreaCreate is the builder for creating a TargetArea entity.
type TargetAreaCreate struct {
	config
	mutation *TargetAreaMutation
	hooks    []Hook
}

// SetContinent sets the "continent" field.
func (tac *TargetAreaCreate) SetContinent(s string) *TargetAreaCreate {
	tac.mutation.SetContinent(s)
	return tac
}

// SetNillableContinent sets the "continent" field if the given value is not nil.
func (tac *TargetAreaCreate) SetNillableContinent(s *string) *TargetAreaCreate {
	if s != nil {
		tac.SetContinent(*s)
	}
	return tac
}

// SetCountry sets the "country" field.
func (tac *TargetAreaCreate) SetCountry(s string) *TargetAreaCreate {
	tac.mutation.SetCountry(s)
	return tac
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (tac *TargetAreaCreate) SetNillableCountry(s *string) *TargetAreaCreate {
	if s != nil {
		tac.SetCountry(*s)
	}
	return tac
}

// SetCreateAt sets the "create_at" field.
func (tac *TargetAreaCreate) SetCreateAt(i int64) *TargetAreaCreate {
	tac.mutation.SetCreateAt(i)
	return tac
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (tac *TargetAreaCreate) SetNillableCreateAt(i *int64) *TargetAreaCreate {
	if i != nil {
		tac.SetCreateAt(*i)
	}
	return tac
}

// SetUpdateAt sets the "update_at" field.
func (tac *TargetAreaCreate) SetUpdateAt(i int64) *TargetAreaCreate {
	tac.mutation.SetUpdateAt(i)
	return tac
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (tac *TargetAreaCreate) SetNillableUpdateAt(i *int64) *TargetAreaCreate {
	if i != nil {
		tac.SetUpdateAt(*i)
	}
	return tac
}

// SetDeleteAt sets the "delete_at" field.
func (tac *TargetAreaCreate) SetDeleteAt(i int64) *TargetAreaCreate {
	tac.mutation.SetDeleteAt(i)
	return tac
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (tac *TargetAreaCreate) SetNillableDeleteAt(i *int64) *TargetAreaCreate {
	if i != nil {
		tac.SetDeleteAt(*i)
	}
	return tac
}

// SetID sets the "id" field.
func (tac *TargetAreaCreate) SetID(u uuid.UUID) *TargetAreaCreate {
	tac.mutation.SetID(u)
	return tac
}

// Mutation returns the TargetAreaMutation object of the builder.
func (tac *TargetAreaCreate) Mutation() *TargetAreaMutation {
	return tac.mutation
}

// Save creates the TargetArea in the database.
func (tac *TargetAreaCreate) Save(ctx context.Context) (*TargetArea, error) {
	var (
		err  error
		node *TargetArea
	)
	tac.defaults()
	if len(tac.hooks) == 0 {
		if err = tac.check(); err != nil {
			return nil, err
		}
		node, err = tac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetAreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tac.check(); err != nil {
				return nil, err
			}
			tac.mutation = mutation
			if node, err = tac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tac.hooks) - 1; i >= 0; i-- {
			if tac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tac *TargetAreaCreate) SaveX(ctx context.Context) *TargetArea {
	v, err := tac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tac *TargetAreaCreate) Exec(ctx context.Context) error {
	_, err := tac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tac *TargetAreaCreate) ExecX(ctx context.Context) {
	if err := tac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tac *TargetAreaCreate) defaults() {
	if _, ok := tac.mutation.Continent(); !ok {
		v := targetarea.DefaultContinent
		tac.mutation.SetContinent(v)
	}
	if _, ok := tac.mutation.Country(); !ok {
		v := targetarea.DefaultCountry
		tac.mutation.SetCountry(v)
	}
	if _, ok := tac.mutation.CreateAt(); !ok {
		v := targetarea.DefaultCreateAt()
		tac.mutation.SetCreateAt(v)
	}
	if _, ok := tac.mutation.UpdateAt(); !ok {
		v := targetarea.DefaultUpdateAt()
		tac.mutation.SetUpdateAt(v)
	}
	if _, ok := tac.mutation.DeleteAt(); !ok {
		v := targetarea.DefaultDeleteAt()
		tac.mutation.SetDeleteAt(v)
	}
	if _, ok := tac.mutation.ID(); !ok {
		v := targetarea.DefaultID()
		tac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tac *TargetAreaCreate) check() error {
	if _, ok := tac.mutation.Continent(); !ok {
		return &ValidationError{Name: "continent", err: errors.New(`ent: missing required field "continent"`)}
	}
	if v, ok := tac.mutation.Continent(); ok {
		if err := targetarea.ContinentValidator(v); err != nil {
			return &ValidationError{Name: "continent", err: fmt.Errorf(`ent: validator failed for field "continent": %w`, err)}
		}
	}
	if _, ok := tac.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "country"`)}
	}
	if v, ok := tac.mutation.Country(); ok {
		if err := targetarea.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "country": %w`, err)}
		}
	}
	if _, ok := tac.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := tac.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := tac.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (tac *TargetAreaCreate) sqlSave(ctx context.Context) (*TargetArea, error) {
	_node, _spec := tac.createSpec()
	if err := sqlgraph.CreateNode(ctx, tac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (tac *TargetAreaCreate) createSpec() (*TargetArea, *sqlgraph.CreateSpec) {
	var (
		_node = &TargetArea{config: tac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: targetarea.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: targetarea.FieldID,
			},
		}
	)
	if id, ok := tac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tac.mutation.Continent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: targetarea.FieldContinent,
		})
		_node.Continent = value
	}
	if value, ok := tac.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: targetarea.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := tac.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: targetarea.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := tac.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: targetarea.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := tac.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: targetarea.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// TargetAreaCreateBulk is the builder for creating many TargetArea entities in bulk.
type TargetAreaCreateBulk struct {
	config
	builders []*TargetAreaCreate
}

// Save creates the TargetArea entities in the database.
func (tacb *TargetAreaCreateBulk) Save(ctx context.Context) ([]*TargetArea, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tacb.builders))
	nodes := make([]*TargetArea, len(tacb.builders))
	mutators := make([]Mutator, len(tacb.builders))
	for i := range tacb.builders {
		func(i int, root context.Context) {
			builder := tacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TargetAreaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tacb *TargetAreaCreateBulk) SaveX(ctx context.Context) []*TargetArea {
	v, err := tacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tacb *TargetAreaCreateBulk) Exec(ctx context.Context) error {
	_, err := tacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tacb *TargetAreaCreateBulk) ExecX(ctx context.Context) {
	if err := tacb.Exec(ctx); err != nil {
		panic(err)
	}
}
