// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodinfo"
	"github.com/google/uuid"
)

// GoodInfoCreate is the builder for creating a GoodInfo entity.
type GoodInfoCreate struct {
	config
	mutation *GoodInfoMutation
	hooks    []Hook
}

// SetDeviceInfoID sets the "device_info_id" field.
func (gic *GoodInfoCreate) SetDeviceInfoID(u uuid.UUID) *GoodInfoCreate {
	gic.mutation.SetDeviceInfoID(u)
	return gic
}

// SetSeparateFee sets the "separate_fee" field.
func (gic *GoodInfoCreate) SetSeparateFee(b bool) *GoodInfoCreate {
	gic.mutation.SetSeparateFee(b)
	return gic
}

// SetUnitPower sets the "unit_power" field.
func (gic *GoodInfoCreate) SetUnitPower(i int32) *GoodInfoCreate {
	gic.mutation.SetUnitPower(i)
	return gic
}

// SetDurationDays sets the "duration_days" field.
func (gic *GoodInfoCreate) SetDurationDays(i int32) *GoodInfoCreate {
	gic.mutation.SetDurationDays(i)
	return gic
}

// SetCoinInfoID sets the "coin_info_id" field.
func (gic *GoodInfoCreate) SetCoinInfoID(u uuid.UUID) *GoodInfoCreate {
	gic.mutation.SetCoinInfoID(u)
	return gic
}

// SetActuals sets the "actuals" field.
func (gic *GoodInfoCreate) SetActuals(b bool) *GoodInfoCreate {
	gic.mutation.SetActuals(b)
	return gic
}

// SetDeliveryAt sets the "delivery_at" field.
func (gic *GoodInfoCreate) SetDeliveryAt(u uint32) *GoodInfoCreate {
	gic.mutation.SetDeliveryAt(u)
	return gic
}

// SetInheritFromGoodID sets the "inherit_from_good_id" field.
func (gic *GoodInfoCreate) SetInheritFromGoodID(u uuid.UUID) *GoodInfoCreate {
	gic.mutation.SetInheritFromGoodID(u)
	return gic
}

// SetVendorLocationID sets the "vendor_location_id" field.
func (gic *GoodInfoCreate) SetVendorLocationID(u uuid.UUID) *GoodInfoCreate {
	gic.mutation.SetVendorLocationID(u)
	return gic
}

// SetPrice sets the "price" field.
func (gic *GoodInfoCreate) SetPrice(u uint64) *GoodInfoCreate {
	gic.mutation.SetPrice(u)
	return gic
}

// SetPriceCurrency sets the "price_currency" field.
func (gic *GoodInfoCreate) SetPriceCurrency(u uuid.UUID) *GoodInfoCreate {
	gic.mutation.SetPriceCurrency(u)
	return gic
}

// SetBenefitType sets the "benefit_type" field.
func (gic *GoodInfoCreate) SetBenefitType(gt goodinfo.BenefitType) *GoodInfoCreate {
	gic.mutation.SetBenefitType(gt)
	return gic
}

// SetClassic sets the "classic" field.
func (gic *GoodInfoCreate) SetClassic(b bool) *GoodInfoCreate {
	gic.mutation.SetClassic(b)
	return gic
}

// SetTitle sets the "title" field.
func (gic *GoodInfoCreate) SetTitle(s string) *GoodInfoCreate {
	gic.mutation.SetTitle(s)
	return gic
}

// SetUnit sets the "unit" field.
func (gic *GoodInfoCreate) SetUnit(s string) *GoodInfoCreate {
	gic.mutation.SetUnit(s)
	return gic
}

// SetSupportCoinTypeIds sets the "support_coin_type_ids" field.
func (gic *GoodInfoCreate) SetSupportCoinTypeIds(u []uuid.UUID) *GoodInfoCreate {
	gic.mutation.SetSupportCoinTypeIds(u)
	return gic
}

// SetFeeIds sets the "fee_ids" field.
func (gic *GoodInfoCreate) SetFeeIds(u []uuid.UUID) *GoodInfoCreate {
	gic.mutation.SetFeeIds(u)
	return gic
}

// SetTotal sets the "total" field.
func (gic *GoodInfoCreate) SetTotal(i int32) *GoodInfoCreate {
	gic.mutation.SetTotal(i)
	return gic
}

// SetCreateAt sets the "create_at" field.
func (gic *GoodInfoCreate) SetCreateAt(u uint32) *GoodInfoCreate {
	gic.mutation.SetCreateAt(u)
	return gic
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (gic *GoodInfoCreate) SetNillableCreateAt(u *uint32) *GoodInfoCreate {
	if u != nil {
		gic.SetCreateAt(*u)
	}
	return gic
}

// SetUpdateAt sets the "update_at" field.
func (gic *GoodInfoCreate) SetUpdateAt(u uint32) *GoodInfoCreate {
	gic.mutation.SetUpdateAt(u)
	return gic
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (gic *GoodInfoCreate) SetNillableUpdateAt(u *uint32) *GoodInfoCreate {
	if u != nil {
		gic.SetUpdateAt(*u)
	}
	return gic
}

// SetDeleteAt sets the "delete_at" field.
func (gic *GoodInfoCreate) SetDeleteAt(u uint32) *GoodInfoCreate {
	gic.mutation.SetDeleteAt(u)
	return gic
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (gic *GoodInfoCreate) SetNillableDeleteAt(u *uint32) *GoodInfoCreate {
	if u != nil {
		gic.SetDeleteAt(*u)
	}
	return gic
}

// SetID sets the "id" field.
func (gic *GoodInfoCreate) SetID(u uuid.UUID) *GoodInfoCreate {
	gic.mutation.SetID(u)
	return gic
}

// Mutation returns the GoodInfoMutation object of the builder.
func (gic *GoodInfoCreate) Mutation() *GoodInfoMutation {
	return gic.mutation
}

// Save creates the GoodInfo in the database.
func (gic *GoodInfoCreate) Save(ctx context.Context) (*GoodInfo, error) {
	var (
		err  error
		node *GoodInfo
	)
	gic.defaults()
	if len(gic.hooks) == 0 {
		if err = gic.check(); err != nil {
			return nil, err
		}
		node, err = gic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gic.check(); err != nil {
				return nil, err
			}
			gic.mutation = mutation
			if node, err = gic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gic.hooks) - 1; i >= 0; i-- {
			if gic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gic *GoodInfoCreate) SaveX(ctx context.Context) *GoodInfo {
	v, err := gic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gic *GoodInfoCreate) Exec(ctx context.Context) error {
	_, err := gic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gic *GoodInfoCreate) ExecX(ctx context.Context) {
	if err := gic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gic *GoodInfoCreate) defaults() {
	if _, ok := gic.mutation.CreateAt(); !ok {
		v := goodinfo.DefaultCreateAt()
		gic.mutation.SetCreateAt(v)
	}
	if _, ok := gic.mutation.UpdateAt(); !ok {
		v := goodinfo.DefaultUpdateAt()
		gic.mutation.SetUpdateAt(v)
	}
	if _, ok := gic.mutation.DeleteAt(); !ok {
		v := goodinfo.DefaultDeleteAt()
		gic.mutation.SetDeleteAt(v)
	}
	if _, ok := gic.mutation.ID(); !ok {
		v := goodinfo.DefaultID()
		gic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gic *GoodInfoCreate) check() error {
	if _, ok := gic.mutation.DeviceInfoID(); !ok {
		return &ValidationError{Name: "device_info_id", err: errors.New(`ent: missing required field "device_info_id"`)}
	}
	if _, ok := gic.mutation.SeparateFee(); !ok {
		return &ValidationError{Name: "separate_fee", err: errors.New(`ent: missing required field "separate_fee"`)}
	}
	if _, ok := gic.mutation.UnitPower(); !ok {
		return &ValidationError{Name: "unit_power", err: errors.New(`ent: missing required field "unit_power"`)}
	}
	if v, ok := gic.mutation.UnitPower(); ok {
		if err := goodinfo.UnitPowerValidator(v); err != nil {
			return &ValidationError{Name: "unit_power", err: fmt.Errorf(`ent: validator failed for field "unit_power": %w`, err)}
		}
	}
	if _, ok := gic.mutation.DurationDays(); !ok {
		return &ValidationError{Name: "duration_days", err: errors.New(`ent: missing required field "duration_days"`)}
	}
	if v, ok := gic.mutation.DurationDays(); ok {
		if err := goodinfo.DurationDaysValidator(v); err != nil {
			return &ValidationError{Name: "duration_days", err: fmt.Errorf(`ent: validator failed for field "duration_days": %w`, err)}
		}
	}
	if _, ok := gic.mutation.CoinInfoID(); !ok {
		return &ValidationError{Name: "coin_info_id", err: errors.New(`ent: missing required field "coin_info_id"`)}
	}
	if _, ok := gic.mutation.Actuals(); !ok {
		return &ValidationError{Name: "actuals", err: errors.New(`ent: missing required field "actuals"`)}
	}
	if _, ok := gic.mutation.DeliveryAt(); !ok {
		return &ValidationError{Name: "delivery_at", err: errors.New(`ent: missing required field "delivery_at"`)}
	}
	if _, ok := gic.mutation.InheritFromGoodID(); !ok {
		return &ValidationError{Name: "inherit_from_good_id", err: errors.New(`ent: missing required field "inherit_from_good_id"`)}
	}
	if _, ok := gic.mutation.VendorLocationID(); !ok {
		return &ValidationError{Name: "vendor_location_id", err: errors.New(`ent: missing required field "vendor_location_id"`)}
	}
	if _, ok := gic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "price"`)}
	}
	if v, ok := gic.mutation.Price(); ok {
		if err := goodinfo.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "price": %w`, err)}
		}
	}
	if _, ok := gic.mutation.PriceCurrency(); !ok {
		return &ValidationError{Name: "price_currency", err: errors.New(`ent: missing required field "price_currency"`)}
	}
	if _, ok := gic.mutation.BenefitType(); !ok {
		return &ValidationError{Name: "benefit_type", err: errors.New(`ent: missing required field "benefit_type"`)}
	}
	if v, ok := gic.mutation.BenefitType(); ok {
		if err := goodinfo.BenefitTypeValidator(v); err != nil {
			return &ValidationError{Name: "benefit_type", err: fmt.Errorf(`ent: validator failed for field "benefit_type": %w`, err)}
		}
	}
	if _, ok := gic.mutation.Classic(); !ok {
		return &ValidationError{Name: "classic", err: errors.New(`ent: missing required field "classic"`)}
	}
	if _, ok := gic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if _, ok := gic.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New(`ent: missing required field "unit"`)}
	}
	if _, ok := gic.mutation.SupportCoinTypeIds(); !ok {
		return &ValidationError{Name: "support_coin_type_ids", err: errors.New(`ent: missing required field "support_coin_type_ids"`)}
	}
	if _, ok := gic.mutation.FeeIds(); !ok {
		return &ValidationError{Name: "fee_ids", err: errors.New(`ent: missing required field "fee_ids"`)}
	}
	if _, ok := gic.mutation.Total(); !ok {
		return &ValidationError{Name: "total", err: errors.New(`ent: missing required field "total"`)}
	}
	if v, ok := gic.mutation.Total(); ok {
		if err := goodinfo.TotalValidator(v); err != nil {
			return &ValidationError{Name: "total", err: fmt.Errorf(`ent: validator failed for field "total": %w`, err)}
		}
	}
	if _, ok := gic.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := gic.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := gic.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (gic *GoodInfoCreate) sqlSave(ctx context.Context) (*GoodInfo, error) {
	_node, _spec := gic.createSpec()
	if err := sqlgraph.CreateNode(ctx, gic.driver, _spec); err != nil {
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

func (gic *GoodInfoCreate) createSpec() (*GoodInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &GoodInfo{config: gic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goodinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodinfo.FieldID,
			},
		}
	)
	if id, ok := gic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gic.mutation.DeviceInfoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodinfo.FieldDeviceInfoID,
		})
		_node.DeviceInfoID = value
	}
	if value, ok := gic.mutation.SeparateFee(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodinfo.FieldSeparateFee,
		})
		_node.SeparateFee = value
	}
	if value, ok := gic.mutation.UnitPower(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: goodinfo.FieldUnitPower,
		})
		_node.UnitPower = value
	}
	if value, ok := gic.mutation.DurationDays(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: goodinfo.FieldDurationDays,
		})
		_node.DurationDays = value
	}
	if value, ok := gic.mutation.CoinInfoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodinfo.FieldCoinInfoID,
		})
		_node.CoinInfoID = value
	}
	if value, ok := gic.mutation.Actuals(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodinfo.FieldActuals,
		})
		_node.Actuals = value
	}
	if value, ok := gic.mutation.DeliveryAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodinfo.FieldDeliveryAt,
		})
		_node.DeliveryAt = value
	}
	if value, ok := gic.mutation.InheritFromGoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodinfo.FieldInheritFromGoodID,
		})
		_node.InheritFromGoodID = value
	}
	if value, ok := gic.mutation.VendorLocationID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodinfo.FieldVendorLocationID,
		})
		_node.VendorLocationID = value
	}
	if value, ok := gic.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: goodinfo.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := gic.mutation.PriceCurrency(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodinfo.FieldPriceCurrency,
		})
		_node.PriceCurrency = value
	}
	if value, ok := gic.mutation.BenefitType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: goodinfo.FieldBenefitType,
		})
		_node.BenefitType = value
	}
	if value, ok := gic.mutation.Classic(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodinfo.FieldClassic,
		})
		_node.Classic = value
	}
	if value, ok := gic.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodinfo.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := gic.mutation.Unit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodinfo.FieldUnit,
		})
		_node.Unit = value
	}
	if value, ok := gic.mutation.SupportCoinTypeIds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: goodinfo.FieldSupportCoinTypeIds,
		})
		_node.SupportCoinTypeIds = value
	}
	if value, ok := gic.mutation.FeeIds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: goodinfo.FieldFeeIds,
		})
		_node.FeeIds = value
	}
	if value, ok := gic.mutation.Total(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: goodinfo.FieldTotal,
		})
		_node.Total = value
	}
	if value, ok := gic.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodinfo.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := gic.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodinfo.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := gic.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodinfo.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// GoodInfoCreateBulk is the builder for creating many GoodInfo entities in bulk.
type GoodInfoCreateBulk struct {
	config
	builders []*GoodInfoCreate
}

// Save creates the GoodInfo entities in the database.
func (gicb *GoodInfoCreateBulk) Save(ctx context.Context) ([]*GoodInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gicb.builders))
	nodes := make([]*GoodInfo, len(gicb.builders))
	mutators := make([]Mutator, len(gicb.builders))
	for i := range gicb.builders {
		func(i int, root context.Context) {
			builder := gicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, gicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gicb *GoodInfoCreateBulk) SaveX(ctx context.Context) []*GoodInfo {
	v, err := gicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gicb *GoodInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := gicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gicb *GoodInfoCreateBulk) ExecX(ctx context.Context) {
	if err := gicb.Exec(ctx); err != nil {
		panic(err)
	}
}
