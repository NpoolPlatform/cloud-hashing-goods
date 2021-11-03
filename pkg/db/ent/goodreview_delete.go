// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodreview"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
)

// GoodReviewDelete is the builder for deleting a GoodReview entity.
type GoodReviewDelete struct {
	config
	hooks    []Hook
	mutation *GoodReviewMutation
}

// Where appends a list predicates to the GoodReviewDelete builder.
func (grd *GoodReviewDelete) Where(ps ...predicate.GoodReview) *GoodReviewDelete {
	grd.mutation.Where(ps...)
	return grd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (grd *GoodReviewDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(grd.hooks) == 0 {
		affected, err = grd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			grd.mutation = mutation
			affected, err = grd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(grd.hooks) - 1; i >= 0; i-- {
			if grd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = grd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, grd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (grd *GoodReviewDelete) ExecX(ctx context.Context) int {
	n, err := grd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (grd *GoodReviewDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: goodreview.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodreview.FieldID,
			},
		},
	}
	if ps := grd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, grd.driver, _spec)
}

// GoodReviewDeleteOne is the builder for deleting a single GoodReview entity.
type GoodReviewDeleteOne struct {
	grd *GoodReviewDelete
}

// Exec executes the deletion query.
func (grdo *GoodReviewDeleteOne) Exec(ctx context.Context) error {
	n, err := grdo.grd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodreview.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (grdo *GoodReviewDeleteOne) ExecX(ctx context.Context) {
	grdo.grd.ExecX(ctx)
}