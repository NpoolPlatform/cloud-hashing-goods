// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgoodpromotion"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
)

// AppGoodPromotionDelete is the builder for deleting a AppGoodPromotion entity.
type AppGoodPromotionDelete struct {
	config
	hooks    []Hook
	mutation *AppGoodPromotionMutation
}

// Where appends a list predicates to the AppGoodPromotionDelete builder.
func (agpd *AppGoodPromotionDelete) Where(ps ...predicate.AppGoodPromotion) *AppGoodPromotionDelete {
	agpd.mutation.Where(ps...)
	return agpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (agpd *AppGoodPromotionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(agpd.hooks) == 0 {
		affected, err = agpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodPromotionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agpd.mutation = mutation
			affected, err = agpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agpd.hooks) - 1; i >= 0; i-- {
			if agpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (agpd *AppGoodPromotionDelete) ExecX(ctx context.Context) int {
	n, err := agpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (agpd *AppGoodPromotionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appgoodpromotion.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appgoodpromotion.FieldID,
			},
		},
	}
	if ps := agpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, agpd.driver, _spec)
}

// AppGoodPromotionDeleteOne is the builder for deleting a single AppGoodPromotion entity.
type AppGoodPromotionDeleteOne struct {
	agpd *AppGoodPromotionDelete
}

// Exec executes the deletion query.
func (agpdo *AppGoodPromotionDeleteOne) Exec(ctx context.Context) error {
	n, err := agpdo.agpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appgoodpromotion.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (agpdo *AppGoodPromotionDeleteOne) ExecX(ctx context.Context) {
	agpdo.agpd.ExecX(ctx)
}