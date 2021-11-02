// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/vendorlocation"
)

// VendorLocationDelete is the builder for deleting a VendorLocation entity.
type VendorLocationDelete struct {
	config
	hooks    []Hook
	mutation *VendorLocationMutation
}

// Where appends a list predicates to the VendorLocationDelete builder.
func (vld *VendorLocationDelete) Where(ps ...predicate.VendorLocation) *VendorLocationDelete {
	vld.mutation.Where(ps...)
	return vld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vld *VendorLocationDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vld.hooks) == 0 {
		affected, err = vld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VendorLocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vld.mutation = mutation
			affected, err = vld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vld.hooks) - 1; i >= 0; i-- {
			if vld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vld *VendorLocationDelete) ExecX(ctx context.Context) int {
	n, err := vld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vld *VendorLocationDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vendorlocation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vendorlocation.FieldID,
			},
		},
	}
	if ps := vld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vld.driver, _spec)
}

// VendorLocationDeleteOne is the builder for deleting a single VendorLocation entity.
type VendorLocationDeleteOne struct {
	vld *VendorLocationDelete
}

// Exec executes the deletion query.
func (vldo *VendorLocationDeleteOne) Exec(ctx context.Context) error {
	n, err := vldo.vld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vendorlocation.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vldo *VendorLocationDeleteOne) ExecX(ctx context.Context) {
	vldo.vld.ExecX(ctx)
}