// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/recommendgood"
	"github.com/google/uuid"
)

// RecommendGoodCreate is the builder for creating a RecommendGood entity.
type RecommendGoodCreate struct {
	config
	mutation *RecommendGoodMutation
	hooks    []Hook
}

// SetRecommenderID sets the "recommender_id" field.
func (rgc *RecommendGoodCreate) SetRecommenderID(u uuid.UUID) *RecommendGoodCreate {
	rgc.mutation.SetRecommenderID(u)
	return rgc
}

// SetGoodID sets the "good_id" field.
func (rgc *RecommendGoodCreate) SetGoodID(u uuid.UUID) *RecommendGoodCreate {
	rgc.mutation.SetGoodID(u)
	return rgc
}

// SetMessage sets the "message" field.
func (rgc *RecommendGoodCreate) SetMessage(s string) *RecommendGoodCreate {
	rgc.mutation.SetMessage(s)
	return rgc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (rgc *RecommendGoodCreate) SetNillableMessage(s *string) *RecommendGoodCreate {
	if s != nil {
		rgc.SetMessage(*s)
	}
	return rgc
}

// SetCreateAt sets the "create_at" field.
func (rgc *RecommendGoodCreate) SetCreateAt(i int64) *RecommendGoodCreate {
	rgc.mutation.SetCreateAt(i)
	return rgc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (rgc *RecommendGoodCreate) SetNillableCreateAt(i *int64) *RecommendGoodCreate {
	if i != nil {
		rgc.SetCreateAt(*i)
	}
	return rgc
}

// SetUpdateAt sets the "update_at" field.
func (rgc *RecommendGoodCreate) SetUpdateAt(i int64) *RecommendGoodCreate {
	rgc.mutation.SetUpdateAt(i)
	return rgc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (rgc *RecommendGoodCreate) SetNillableUpdateAt(i *int64) *RecommendGoodCreate {
	if i != nil {
		rgc.SetUpdateAt(*i)
	}
	return rgc
}

// SetDeleteAt sets the "delete_at" field.
func (rgc *RecommendGoodCreate) SetDeleteAt(i int64) *RecommendGoodCreate {
	rgc.mutation.SetDeleteAt(i)
	return rgc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (rgc *RecommendGoodCreate) SetNillableDeleteAt(i *int64) *RecommendGoodCreate {
	if i != nil {
		rgc.SetDeleteAt(*i)
	}
	return rgc
}

// SetID sets the "id" field.
func (rgc *RecommendGoodCreate) SetID(u uuid.UUID) *RecommendGoodCreate {
	rgc.mutation.SetID(u)
	return rgc
}

// Mutation returns the RecommendGoodMutation object of the builder.
func (rgc *RecommendGoodCreate) Mutation() *RecommendGoodMutation {
	return rgc.mutation
}

// Save creates the RecommendGood in the database.
func (rgc *RecommendGoodCreate) Save(ctx context.Context) (*RecommendGood, error) {
	var (
		err  error
		node *RecommendGood
	)
	rgc.defaults()
	if len(rgc.hooks) == 0 {
		if err = rgc.check(); err != nil {
			return nil, err
		}
		node, err = rgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecommendGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rgc.check(); err != nil {
				return nil, err
			}
			rgc.mutation = mutation
			if node, err = rgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rgc.hooks) - 1; i >= 0; i-- {
			if rgc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rgc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rgc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rgc *RecommendGoodCreate) SaveX(ctx context.Context) *RecommendGood {
	v, err := rgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rgc *RecommendGoodCreate) Exec(ctx context.Context) error {
	_, err := rgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rgc *RecommendGoodCreate) ExecX(ctx context.Context) {
	if err := rgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rgc *RecommendGoodCreate) defaults() {
	if _, ok := rgc.mutation.Message(); !ok {
		v := recommendgood.DefaultMessage
		rgc.mutation.SetMessage(v)
	}
	if _, ok := rgc.mutation.CreateAt(); !ok {
		v := recommendgood.DefaultCreateAt()
		rgc.mutation.SetCreateAt(v)
	}
	if _, ok := rgc.mutation.UpdateAt(); !ok {
		v := recommendgood.DefaultUpdateAt()
		rgc.mutation.SetUpdateAt(v)
	}
	if _, ok := rgc.mutation.DeleteAt(); !ok {
		v := recommendgood.DefaultDeleteAt()
		rgc.mutation.SetDeleteAt(v)
	}
	if _, ok := rgc.mutation.ID(); !ok {
		v := recommendgood.DefaultID()
		rgc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rgc *RecommendGoodCreate) check() error {
	if _, ok := rgc.mutation.RecommenderID(); !ok {
		return &ValidationError{Name: "recommender_id", err: errors.New(`ent: missing required field "recommender_id"`)}
	}
	if _, ok := rgc.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "good_id"`)}
	}
	if _, ok := rgc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if _, ok := rgc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := rgc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := rgc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (rgc *RecommendGoodCreate) sqlSave(ctx context.Context) (*RecommendGood, error) {
	_node, _spec := rgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rgc.driver, _spec); err != nil {
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

func (rgc *RecommendGoodCreate) createSpec() (*RecommendGood, *sqlgraph.CreateSpec) {
	var (
		_node = &RecommendGood{config: rgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recommendgood.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: recommendgood.FieldID,
			},
		}
	)
	if id, ok := rgc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rgc.mutation.RecommenderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommendgood.FieldRecommenderID,
		})
		_node.RecommenderID = value
	}
	if value, ok := rgc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommendgood.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := rgc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recommendgood.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := rgc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommendgood.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := rgc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommendgood.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := rgc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: recommendgood.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// RecommendGoodCreateBulk is the builder for creating many RecommendGood entities in bulk.
type RecommendGoodCreateBulk struct {
	config
	builders []*RecommendGoodCreate
}

// Save creates the RecommendGood entities in the database.
func (rgcb *RecommendGoodCreateBulk) Save(ctx context.Context) ([]*RecommendGood, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rgcb.builders))
	nodes := make([]*RecommendGood, len(rgcb.builders))
	mutators := make([]Mutator, len(rgcb.builders))
	for i := range rgcb.builders {
		func(i int, root context.Context) {
			builder := rgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecommendGoodMutation)
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
					_, err = mutators[i+1].Mutate(root, rgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rgcb *RecommendGoodCreateBulk) SaveX(ctx context.Context) []*RecommendGood {
	v, err := rgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rgcb *RecommendGoodCreateBulk) Exec(ctx context.Context) error {
	_, err := rgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rgcb *RecommendGoodCreateBulk) ExecX(ctx context.Context) {
	if err := rgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
