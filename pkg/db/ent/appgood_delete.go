// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgood"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
)

// AppGoodDelete is the builder for deleting a AppGood entity.
type AppGoodDelete struct {
	config
	hooks    []Hook
	mutation *AppGoodMutation
}

// Where appends a list predicates to the AppGoodDelete builder.
func (agd *AppGoodDelete) Where(ps ...predicate.AppGood) *AppGoodDelete {
	agd.mutation.Where(ps...)
	return agd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (agd *AppGoodDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(agd.hooks) == 0 {
		affected, err = agd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agd.mutation = mutation
			affected, err = agd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agd.hooks) - 1; i >= 0; i-- {
			if agd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (agd *AppGoodDelete) ExecX(ctx context.Context) int {
	n, err := agd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (agd *AppGoodDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appgood.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appgood.FieldID,
			},
		},
	}
	if ps := agd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, agd.driver, _spec)
}

// AppGoodDeleteOne is the builder for deleting a single AppGood entity.
type AppGoodDeleteOne struct {
	agd *AppGoodDelete
}

// Exec executes the deletion query.
func (agdo *AppGoodDeleteOne) Exec(ctx context.Context) error {
	n, err := agdo.agd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appgood.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (agdo *AppGoodDeleteOne) ExecX(ctx context.Context) {
	agdo.agd.ExecX(ctx)
}
