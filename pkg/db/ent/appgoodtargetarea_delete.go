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
)

// AppGoodTargetAreaDelete is the builder for deleting a AppGoodTargetArea entity.
type AppGoodTargetAreaDelete struct {
	config
	hooks    []Hook
	mutation *AppGoodTargetAreaMutation
}

// Where appends a list predicates to the AppGoodTargetAreaDelete builder.
func (agtad *AppGoodTargetAreaDelete) Where(ps ...predicate.AppGoodTargetArea) *AppGoodTargetAreaDelete {
	agtad.mutation.Where(ps...)
	return agtad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (agtad *AppGoodTargetAreaDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(agtad.hooks) == 0 {
		affected, err = agtad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodTargetAreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agtad.mutation = mutation
			affected, err = agtad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agtad.hooks) - 1; i >= 0; i-- {
			if agtad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agtad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agtad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (agtad *AppGoodTargetAreaDelete) ExecX(ctx context.Context) int {
	n, err := agtad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (agtad *AppGoodTargetAreaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appgoodtargetarea.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appgoodtargetarea.FieldID,
			},
		},
	}
	if ps := agtad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, agtad.driver, _spec)
}

// AppGoodTargetAreaDeleteOne is the builder for deleting a single AppGoodTargetArea entity.
type AppGoodTargetAreaDeleteOne struct {
	agtad *AppGoodTargetAreaDelete
}

// Exec executes the deletion query.
func (agtado *AppGoodTargetAreaDeleteOne) Exec(ctx context.Context) error {
	n, err := agtado.agtad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appgoodtargetarea.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (agtado *AppGoodTargetAreaDeleteOne) ExecX(ctx context.Context) {
	agtado.agtad.ExecX(ctx)
}
