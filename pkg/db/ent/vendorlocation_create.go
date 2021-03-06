// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/vendorlocation"
	"github.com/google/uuid"
)

// VendorLocationCreate is the builder for creating a VendorLocation entity.
type VendorLocationCreate struct {
	config
	mutation *VendorLocationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCountry sets the "country" field.
func (vlc *VendorLocationCreate) SetCountry(s string) *VendorLocationCreate {
	vlc.mutation.SetCountry(s)
	return vlc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (vlc *VendorLocationCreate) SetNillableCountry(s *string) *VendorLocationCreate {
	if s != nil {
		vlc.SetCountry(*s)
	}
	return vlc
}

// SetProvince sets the "province" field.
func (vlc *VendorLocationCreate) SetProvince(s string) *VendorLocationCreate {
	vlc.mutation.SetProvince(s)
	return vlc
}

// SetNillableProvince sets the "province" field if the given value is not nil.
func (vlc *VendorLocationCreate) SetNillableProvince(s *string) *VendorLocationCreate {
	if s != nil {
		vlc.SetProvince(*s)
	}
	return vlc
}

// SetCity sets the "city" field.
func (vlc *VendorLocationCreate) SetCity(s string) *VendorLocationCreate {
	vlc.mutation.SetCity(s)
	return vlc
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (vlc *VendorLocationCreate) SetNillableCity(s *string) *VendorLocationCreate {
	if s != nil {
		vlc.SetCity(*s)
	}
	return vlc
}

// SetAddress sets the "address" field.
func (vlc *VendorLocationCreate) SetAddress(s string) *VendorLocationCreate {
	vlc.mutation.SetAddress(s)
	return vlc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (vlc *VendorLocationCreate) SetNillableAddress(s *string) *VendorLocationCreate {
	if s != nil {
		vlc.SetAddress(*s)
	}
	return vlc
}

// SetCreateAt sets the "create_at" field.
func (vlc *VendorLocationCreate) SetCreateAt(i int64) *VendorLocationCreate {
	vlc.mutation.SetCreateAt(i)
	return vlc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (vlc *VendorLocationCreate) SetNillableCreateAt(i *int64) *VendorLocationCreate {
	if i != nil {
		vlc.SetCreateAt(*i)
	}
	return vlc
}

// SetUpdateAt sets the "update_at" field.
func (vlc *VendorLocationCreate) SetUpdateAt(i int64) *VendorLocationCreate {
	vlc.mutation.SetUpdateAt(i)
	return vlc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (vlc *VendorLocationCreate) SetNillableUpdateAt(i *int64) *VendorLocationCreate {
	if i != nil {
		vlc.SetUpdateAt(*i)
	}
	return vlc
}

// SetDeleteAt sets the "delete_at" field.
func (vlc *VendorLocationCreate) SetDeleteAt(i int64) *VendorLocationCreate {
	vlc.mutation.SetDeleteAt(i)
	return vlc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (vlc *VendorLocationCreate) SetNillableDeleteAt(i *int64) *VendorLocationCreate {
	if i != nil {
		vlc.SetDeleteAt(*i)
	}
	return vlc
}

// SetID sets the "id" field.
func (vlc *VendorLocationCreate) SetID(u uuid.UUID) *VendorLocationCreate {
	vlc.mutation.SetID(u)
	return vlc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vlc *VendorLocationCreate) SetNillableID(u *uuid.UUID) *VendorLocationCreate {
	if u != nil {
		vlc.SetID(*u)
	}
	return vlc
}

// Mutation returns the VendorLocationMutation object of the builder.
func (vlc *VendorLocationCreate) Mutation() *VendorLocationMutation {
	return vlc.mutation
}

// Save creates the VendorLocation in the database.
func (vlc *VendorLocationCreate) Save(ctx context.Context) (*VendorLocation, error) {
	var (
		err  error
		node *VendorLocation
	)
	vlc.defaults()
	if len(vlc.hooks) == 0 {
		if err = vlc.check(); err != nil {
			return nil, err
		}
		node, err = vlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VendorLocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vlc.check(); err != nil {
				return nil, err
			}
			vlc.mutation = mutation
			if node, err = vlc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vlc.hooks) - 1; i >= 0; i-- {
			if vlc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vlc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vlc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vlc *VendorLocationCreate) SaveX(ctx context.Context) *VendorLocation {
	v, err := vlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vlc *VendorLocationCreate) Exec(ctx context.Context) error {
	_, err := vlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vlc *VendorLocationCreate) ExecX(ctx context.Context) {
	if err := vlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vlc *VendorLocationCreate) defaults() {
	if _, ok := vlc.mutation.Country(); !ok {
		v := vendorlocation.DefaultCountry
		vlc.mutation.SetCountry(v)
	}
	if _, ok := vlc.mutation.Province(); !ok {
		v := vendorlocation.DefaultProvince
		vlc.mutation.SetProvince(v)
	}
	if _, ok := vlc.mutation.City(); !ok {
		v := vendorlocation.DefaultCity
		vlc.mutation.SetCity(v)
	}
	if _, ok := vlc.mutation.Address(); !ok {
		v := vendorlocation.DefaultAddress
		vlc.mutation.SetAddress(v)
	}
	if _, ok := vlc.mutation.CreateAt(); !ok {
		v := vendorlocation.DefaultCreateAt()
		vlc.mutation.SetCreateAt(v)
	}
	if _, ok := vlc.mutation.UpdateAt(); !ok {
		v := vendorlocation.DefaultUpdateAt()
		vlc.mutation.SetUpdateAt(v)
	}
	if _, ok := vlc.mutation.DeleteAt(); !ok {
		v := vendorlocation.DefaultDeleteAt()
		vlc.mutation.SetDeleteAt(v)
	}
	if _, ok := vlc.mutation.ID(); !ok {
		v := vendorlocation.DefaultID()
		vlc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vlc *VendorLocationCreate) check() error {
	if _, ok := vlc.mutation.Country(); !ok {
		return &ValidationError{Name: "country", err: errors.New(`ent: missing required field "VendorLocation.country"`)}
	}
	if v, ok := vlc.mutation.Country(); ok {
		if err := vendorlocation.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`ent: validator failed for field "VendorLocation.country": %w`, err)}
		}
	}
	if _, ok := vlc.mutation.Province(); !ok {
		return &ValidationError{Name: "province", err: errors.New(`ent: missing required field "VendorLocation.province"`)}
	}
	if v, ok := vlc.mutation.Province(); ok {
		if err := vendorlocation.ProvinceValidator(v); err != nil {
			return &ValidationError{Name: "province", err: fmt.Errorf(`ent: validator failed for field "VendorLocation.province": %w`, err)}
		}
	}
	if _, ok := vlc.mutation.City(); !ok {
		return &ValidationError{Name: "city", err: errors.New(`ent: missing required field "VendorLocation.city"`)}
	}
	if v, ok := vlc.mutation.City(); ok {
		if err := vendorlocation.CityValidator(v); err != nil {
			return &ValidationError{Name: "city", err: fmt.Errorf(`ent: validator failed for field "VendorLocation.city": %w`, err)}
		}
	}
	if _, ok := vlc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "VendorLocation.address"`)}
	}
	if v, ok := vlc.mutation.Address(); ok {
		if err := vendorlocation.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "VendorLocation.address": %w`, err)}
		}
	}
	if _, ok := vlc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "VendorLocation.create_at"`)}
	}
	if _, ok := vlc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "VendorLocation.update_at"`)}
	}
	if _, ok := vlc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "VendorLocation.delete_at"`)}
	}
	return nil
}

func (vlc *VendorLocationCreate) sqlSave(ctx context.Context) (*VendorLocation, error) {
	_node, _spec := vlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (vlc *VendorLocationCreate) createSpec() (*VendorLocation, *sqlgraph.CreateSpec) {
	var (
		_node = &VendorLocation{config: vlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vendorlocation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vendorlocation.FieldID,
			},
		}
	)
	_spec.OnConflict = vlc.conflict
	if id, ok := vlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vlc.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vendorlocation.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := vlc.mutation.Province(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vendorlocation.FieldProvince,
		})
		_node.Province = value
	}
	if value, ok := vlc.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vendorlocation.FieldCity,
		})
		_node.City = value
	}
	if value, ok := vlc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vendorlocation.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := vlc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vendorlocation.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := vlc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vendorlocation.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := vlc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vendorlocation.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.VendorLocation.Create().
//		SetCountry(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VendorLocationUpsert) {
//			SetCountry(v+v).
//		}).
//		Exec(ctx)
//
func (vlc *VendorLocationCreate) OnConflict(opts ...sql.ConflictOption) *VendorLocationUpsertOne {
	vlc.conflict = opts
	return &VendorLocationUpsertOne{
		create: vlc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.VendorLocation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (vlc *VendorLocationCreate) OnConflictColumns(columns ...string) *VendorLocationUpsertOne {
	vlc.conflict = append(vlc.conflict, sql.ConflictColumns(columns...))
	return &VendorLocationUpsertOne{
		create: vlc,
	}
}

type (
	// VendorLocationUpsertOne is the builder for "upsert"-ing
	//  one VendorLocation node.
	VendorLocationUpsertOne struct {
		create *VendorLocationCreate
	}

	// VendorLocationUpsert is the "OnConflict" setter.
	VendorLocationUpsert struct {
		*sql.UpdateSet
	}
)

// SetCountry sets the "country" field.
func (u *VendorLocationUpsert) SetCountry(v string) *VendorLocationUpsert {
	u.Set(vendorlocation.FieldCountry, v)
	return u
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *VendorLocationUpsert) UpdateCountry() *VendorLocationUpsert {
	u.SetExcluded(vendorlocation.FieldCountry)
	return u
}

// SetProvince sets the "province" field.
func (u *VendorLocationUpsert) SetProvince(v string) *VendorLocationUpsert {
	u.Set(vendorlocation.FieldProvince, v)
	return u
}

// UpdateProvince sets the "province" field to the value that was provided on create.
func (u *VendorLocationUpsert) UpdateProvince() *VendorLocationUpsert {
	u.SetExcluded(vendorlocation.FieldProvince)
	return u
}

// SetCity sets the "city" field.
func (u *VendorLocationUpsert) SetCity(v string) *VendorLocationUpsert {
	u.Set(vendorlocation.FieldCity, v)
	return u
}

// UpdateCity sets the "city" field to the value that was provided on create.
func (u *VendorLocationUpsert) UpdateCity() *VendorLocationUpsert {
	u.SetExcluded(vendorlocation.FieldCity)
	return u
}

// SetAddress sets the "address" field.
func (u *VendorLocationUpsert) SetAddress(v string) *VendorLocationUpsert {
	u.Set(vendorlocation.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *VendorLocationUpsert) UpdateAddress() *VendorLocationUpsert {
	u.SetExcluded(vendorlocation.FieldAddress)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *VendorLocationUpsert) SetCreateAt(v int64) *VendorLocationUpsert {
	u.Set(vendorlocation.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *VendorLocationUpsert) UpdateCreateAt() *VendorLocationUpsert {
	u.SetExcluded(vendorlocation.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *VendorLocationUpsert) AddCreateAt(v int64) *VendorLocationUpsert {
	u.Add(vendorlocation.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *VendorLocationUpsert) SetUpdateAt(v int64) *VendorLocationUpsert {
	u.Set(vendorlocation.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *VendorLocationUpsert) UpdateUpdateAt() *VendorLocationUpsert {
	u.SetExcluded(vendorlocation.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *VendorLocationUpsert) AddUpdateAt(v int64) *VendorLocationUpsert {
	u.Add(vendorlocation.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *VendorLocationUpsert) SetDeleteAt(v int64) *VendorLocationUpsert {
	u.Set(vendorlocation.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *VendorLocationUpsert) UpdateDeleteAt() *VendorLocationUpsert {
	u.SetExcluded(vendorlocation.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *VendorLocationUpsert) AddDeleteAt(v int64) *VendorLocationUpsert {
	u.Add(vendorlocation.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.VendorLocation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(vendorlocation.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *VendorLocationUpsertOne) UpdateNewValues() *VendorLocationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(vendorlocation.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.VendorLocation.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *VendorLocationUpsertOne) Ignore() *VendorLocationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VendorLocationUpsertOne) DoNothing() *VendorLocationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VendorLocationCreate.OnConflict
// documentation for more info.
func (u *VendorLocationUpsertOne) Update(set func(*VendorLocationUpsert)) *VendorLocationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VendorLocationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCountry sets the "country" field.
func (u *VendorLocationUpsertOne) SetCountry(v string) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetCountry(v)
	})
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *VendorLocationUpsertOne) UpdateCountry() *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateCountry()
	})
}

// SetProvince sets the "province" field.
func (u *VendorLocationUpsertOne) SetProvince(v string) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetProvince(v)
	})
}

// UpdateProvince sets the "province" field to the value that was provided on create.
func (u *VendorLocationUpsertOne) UpdateProvince() *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateProvince()
	})
}

// SetCity sets the "city" field.
func (u *VendorLocationUpsertOne) SetCity(v string) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetCity(v)
	})
}

// UpdateCity sets the "city" field to the value that was provided on create.
func (u *VendorLocationUpsertOne) UpdateCity() *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateCity()
	})
}

// SetAddress sets the "address" field.
func (u *VendorLocationUpsertOne) SetAddress(v string) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *VendorLocationUpsertOne) UpdateAddress() *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateAddress()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *VendorLocationUpsertOne) SetCreateAt(v int64) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *VendorLocationUpsertOne) AddCreateAt(v int64) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *VendorLocationUpsertOne) UpdateCreateAt() *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *VendorLocationUpsertOne) SetUpdateAt(v int64) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *VendorLocationUpsertOne) AddUpdateAt(v int64) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *VendorLocationUpsertOne) UpdateUpdateAt() *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *VendorLocationUpsertOne) SetDeleteAt(v int64) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *VendorLocationUpsertOne) AddDeleteAt(v int64) *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *VendorLocationUpsertOne) UpdateDeleteAt() *VendorLocationUpsertOne {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *VendorLocationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VendorLocationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VendorLocationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *VendorLocationUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: VendorLocationUpsertOne.ID is not supported by MySQL driver. Use VendorLocationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *VendorLocationUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// VendorLocationCreateBulk is the builder for creating many VendorLocation entities in bulk.
type VendorLocationCreateBulk struct {
	config
	builders []*VendorLocationCreate
	conflict []sql.ConflictOption
}

// Save creates the VendorLocation entities in the database.
func (vlcb *VendorLocationCreateBulk) Save(ctx context.Context) ([]*VendorLocation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vlcb.builders))
	nodes := make([]*VendorLocation, len(vlcb.builders))
	mutators := make([]Mutator, len(vlcb.builders))
	for i := range vlcb.builders {
		func(i int, root context.Context) {
			builder := vlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VendorLocationMutation)
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
					_, err = mutators[i+1].Mutate(root, vlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = vlcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vlcb *VendorLocationCreateBulk) SaveX(ctx context.Context) []*VendorLocation {
	v, err := vlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vlcb *VendorLocationCreateBulk) Exec(ctx context.Context) error {
	_, err := vlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vlcb *VendorLocationCreateBulk) ExecX(ctx context.Context) {
	if err := vlcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.VendorLocation.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VendorLocationUpsert) {
//			SetCountry(v+v).
//		}).
//		Exec(ctx)
//
func (vlcb *VendorLocationCreateBulk) OnConflict(opts ...sql.ConflictOption) *VendorLocationUpsertBulk {
	vlcb.conflict = opts
	return &VendorLocationUpsertBulk{
		create: vlcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.VendorLocation.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (vlcb *VendorLocationCreateBulk) OnConflictColumns(columns ...string) *VendorLocationUpsertBulk {
	vlcb.conflict = append(vlcb.conflict, sql.ConflictColumns(columns...))
	return &VendorLocationUpsertBulk{
		create: vlcb,
	}
}

// VendorLocationUpsertBulk is the builder for "upsert"-ing
// a bulk of VendorLocation nodes.
type VendorLocationUpsertBulk struct {
	create *VendorLocationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.VendorLocation.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(vendorlocation.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *VendorLocationUpsertBulk) UpdateNewValues() *VendorLocationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(vendorlocation.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.VendorLocation.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *VendorLocationUpsertBulk) Ignore() *VendorLocationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VendorLocationUpsertBulk) DoNothing() *VendorLocationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VendorLocationCreateBulk.OnConflict
// documentation for more info.
func (u *VendorLocationUpsertBulk) Update(set func(*VendorLocationUpsert)) *VendorLocationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VendorLocationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCountry sets the "country" field.
func (u *VendorLocationUpsertBulk) SetCountry(v string) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetCountry(v)
	})
}

// UpdateCountry sets the "country" field to the value that was provided on create.
func (u *VendorLocationUpsertBulk) UpdateCountry() *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateCountry()
	})
}

// SetProvince sets the "province" field.
func (u *VendorLocationUpsertBulk) SetProvince(v string) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetProvince(v)
	})
}

// UpdateProvince sets the "province" field to the value that was provided on create.
func (u *VendorLocationUpsertBulk) UpdateProvince() *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateProvince()
	})
}

// SetCity sets the "city" field.
func (u *VendorLocationUpsertBulk) SetCity(v string) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetCity(v)
	})
}

// UpdateCity sets the "city" field to the value that was provided on create.
func (u *VendorLocationUpsertBulk) UpdateCity() *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateCity()
	})
}

// SetAddress sets the "address" field.
func (u *VendorLocationUpsertBulk) SetAddress(v string) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *VendorLocationUpsertBulk) UpdateAddress() *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateAddress()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *VendorLocationUpsertBulk) SetCreateAt(v int64) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *VendorLocationUpsertBulk) AddCreateAt(v int64) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *VendorLocationUpsertBulk) UpdateCreateAt() *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *VendorLocationUpsertBulk) SetUpdateAt(v int64) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *VendorLocationUpsertBulk) AddUpdateAt(v int64) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *VendorLocationUpsertBulk) UpdateUpdateAt() *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *VendorLocationUpsertBulk) SetDeleteAt(v int64) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *VendorLocationUpsertBulk) AddDeleteAt(v int64) *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *VendorLocationUpsertBulk) UpdateDeleteAt() *VendorLocationUpsertBulk {
	return u.Update(func(s *VendorLocationUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *VendorLocationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the VendorLocationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VendorLocationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VendorLocationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
