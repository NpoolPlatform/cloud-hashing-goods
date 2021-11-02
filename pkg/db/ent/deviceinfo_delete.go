// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/deviceinfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
)

// DeviceInfoDelete is the builder for deleting a DeviceInfo entity.
type DeviceInfoDelete struct {
	config
	hooks    []Hook
	mutation *DeviceInfoMutation
}

// Where appends a list predicates to the DeviceInfoDelete builder.
func (did *DeviceInfoDelete) Where(ps ...predicate.DeviceInfo) *DeviceInfoDelete {
	did.mutation.Where(ps...)
	return did
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (did *DeviceInfoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(did.hooks) == 0 {
		affected, err = did.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			did.mutation = mutation
			affected, err = did.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(did.hooks) - 1; i >= 0; i-- {
			if did.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = did.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, did.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (did *DeviceInfoDelete) ExecX(ctx context.Context) int {
	n, err := did.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (did *DeviceInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: deviceinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deviceinfo.FieldID,
			},
		},
	}
	if ps := did.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, did.driver, _spec)
}

// DeviceInfoDeleteOne is the builder for deleting a single DeviceInfo entity.
type DeviceInfoDeleteOne struct {
	did *DeviceInfoDelete
}

// Exec executes the deletion query.
func (dido *DeviceInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := dido.did.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deviceinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dido *DeviceInfoDeleteOne) ExecX(ctx context.Context) {
	dido.did.ExecX(ctx)
}