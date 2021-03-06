// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodcomment"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
)

// GoodCommentDelete is the builder for deleting a GoodComment entity.
type GoodCommentDelete struct {
	config
	hooks    []Hook
	mutation *GoodCommentMutation
}

// Where appends a list predicates to the GoodCommentDelete builder.
func (gcd *GoodCommentDelete) Where(ps ...predicate.GoodComment) *GoodCommentDelete {
	gcd.mutation.Where(ps...)
	return gcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gcd *GoodCommentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gcd.hooks) == 0 {
		affected, err = gcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodCommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcd.mutation = mutation
			affected, err = gcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gcd.hooks) - 1; i >= 0; i-- {
			if gcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcd *GoodCommentDelete) ExecX(ctx context.Context) int {
	n, err := gcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gcd *GoodCommentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: goodcomment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodcomment.FieldID,
			},
		},
	}
	if ps := gcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gcd.driver, _spec)
}

// GoodCommentDeleteOne is the builder for deleting a single GoodComment entity.
type GoodCommentDeleteOne struct {
	gcd *GoodCommentDelete
}

// Exec executes the deletion query.
func (gcdo *GoodCommentDeleteOne) Exec(ctx context.Context) error {
	n, err := gcdo.gcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodcomment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gcdo *GoodCommentDeleteOne) ExecX(ctx context.Context) {
	gcdo.gcd.ExecX(ctx)
}
