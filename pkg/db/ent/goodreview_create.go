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
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodreview"
	"github.com/google/uuid"
)

// GoodReviewCreate is the builder for creating a GoodReview entity.
type GoodReviewCreate struct {
	config
	mutation *GoodReviewMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetType sets the "type" field.
func (grc *GoodReviewCreate) SetType(_go goodreview.Type) *GoodReviewCreate {
	grc.mutation.SetType(_go)
	return grc
}

// SetReviewedID sets the "reviewed_id" field.
func (grc *GoodReviewCreate) SetReviewedID(u uuid.UUID) *GoodReviewCreate {
	grc.mutation.SetReviewedID(u)
	return grc
}

// SetReviewerID sets the "reviewer_id" field.
func (grc *GoodReviewCreate) SetReviewerID(u uuid.UUID) *GoodReviewCreate {
	grc.mutation.SetReviewerID(u)
	return grc
}

// SetState sets the "state" field.
func (grc *GoodReviewCreate) SetState(_go goodreview.State) *GoodReviewCreate {
	grc.mutation.SetState(_go)
	return grc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (grc *GoodReviewCreate) SetNillableState(_go *goodreview.State) *GoodReviewCreate {
	if _go != nil {
		grc.SetState(*_go)
	}
	return grc
}

// SetMessage sets the "message" field.
func (grc *GoodReviewCreate) SetMessage(s string) *GoodReviewCreate {
	grc.mutation.SetMessage(s)
	return grc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (grc *GoodReviewCreate) SetNillableMessage(s *string) *GoodReviewCreate {
	if s != nil {
		grc.SetMessage(*s)
	}
	return grc
}

// SetCreateAt sets the "create_at" field.
func (grc *GoodReviewCreate) SetCreateAt(i int64) *GoodReviewCreate {
	grc.mutation.SetCreateAt(i)
	return grc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (grc *GoodReviewCreate) SetNillableCreateAt(i *int64) *GoodReviewCreate {
	if i != nil {
		grc.SetCreateAt(*i)
	}
	return grc
}

// SetUpdateAt sets the "update_at" field.
func (grc *GoodReviewCreate) SetUpdateAt(i int64) *GoodReviewCreate {
	grc.mutation.SetUpdateAt(i)
	return grc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (grc *GoodReviewCreate) SetNillableUpdateAt(i *int64) *GoodReviewCreate {
	if i != nil {
		grc.SetUpdateAt(*i)
	}
	return grc
}

// SetDeleteAt sets the "delete_at" field.
func (grc *GoodReviewCreate) SetDeleteAt(i int64) *GoodReviewCreate {
	grc.mutation.SetDeleteAt(i)
	return grc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (grc *GoodReviewCreate) SetNillableDeleteAt(i *int64) *GoodReviewCreate {
	if i != nil {
		grc.SetDeleteAt(*i)
	}
	return grc
}

// SetID sets the "id" field.
func (grc *GoodReviewCreate) SetID(u uuid.UUID) *GoodReviewCreate {
	grc.mutation.SetID(u)
	return grc
}

// Mutation returns the GoodReviewMutation object of the builder.
func (grc *GoodReviewCreate) Mutation() *GoodReviewMutation {
	return grc.mutation
}

// Save creates the GoodReview in the database.
func (grc *GoodReviewCreate) Save(ctx context.Context) (*GoodReview, error) {
	var (
		err  error
		node *GoodReview
	)
	grc.defaults()
	if len(grc.hooks) == 0 {
		if err = grc.check(); err != nil {
			return nil, err
		}
		node, err = grc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = grc.check(); err != nil {
				return nil, err
			}
			grc.mutation = mutation
			if node, err = grc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(grc.hooks) - 1; i >= 0; i-- {
			if grc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = grc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, grc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (grc *GoodReviewCreate) SaveX(ctx context.Context) *GoodReview {
	v, err := grc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (grc *GoodReviewCreate) Exec(ctx context.Context) error {
	_, err := grc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grc *GoodReviewCreate) ExecX(ctx context.Context) {
	if err := grc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (grc *GoodReviewCreate) defaults() {
	if _, ok := grc.mutation.State(); !ok {
		v := goodreview.DefaultState
		grc.mutation.SetState(v)
	}
	if _, ok := grc.mutation.Message(); !ok {
		v := goodreview.DefaultMessage
		grc.mutation.SetMessage(v)
	}
	if _, ok := grc.mutation.CreateAt(); !ok {
		v := goodreview.DefaultCreateAt()
		grc.mutation.SetCreateAt(v)
	}
	if _, ok := grc.mutation.UpdateAt(); !ok {
		v := goodreview.DefaultUpdateAt()
		grc.mutation.SetUpdateAt(v)
	}
	if _, ok := grc.mutation.DeleteAt(); !ok {
		v := goodreview.DefaultDeleteAt()
		grc.mutation.SetDeleteAt(v)
	}
	if _, ok := grc.mutation.ID(); !ok {
		v := goodreview.DefaultID()
		grc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (grc *GoodReviewCreate) check() error {
	if _, ok := grc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := grc.mutation.GetType(); ok {
		if err := goodreview.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := grc.mutation.ReviewedID(); !ok {
		return &ValidationError{Name: "reviewed_id", err: errors.New(`ent: missing required field "reviewed_id"`)}
	}
	if _, ok := grc.mutation.ReviewerID(); !ok {
		return &ValidationError{Name: "reviewer_id", err: errors.New(`ent: missing required field "reviewer_id"`)}
	}
	if _, ok := grc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := grc.mutation.State(); ok {
		if err := goodreview.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := grc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if _, ok := grc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := grc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := grc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (grc *GoodReviewCreate) sqlSave(ctx context.Context) (*GoodReview, error) {
	_node, _spec := grc.createSpec()
	if err := sqlgraph.CreateNode(ctx, grc.driver, _spec); err != nil {
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

func (grc *GoodReviewCreate) createSpec() (*GoodReview, *sqlgraph.CreateSpec) {
	var (
		_node = &GoodReview{config: grc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goodreview.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodreview.FieldID,
			},
		}
	)
	_spec.OnConflict = grc.conflict
	if id, ok := grc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := grc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: goodreview.FieldType,
		})
		_node.Type = value
	}
	if value, ok := grc.mutation.ReviewedID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodreview.FieldReviewedID,
		})
		_node.ReviewedID = value
	}
	if value, ok := grc.mutation.ReviewerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodreview.FieldReviewerID,
		})
		_node.ReviewerID = value
	}
	if value, ok := grc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: goodreview.FieldState,
		})
		_node.State = value
	}
	if value, ok := grc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodreview.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := grc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodreview.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := grc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodreview.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := grc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: goodreview.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoodReview.Create().
//		SetType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoodReviewUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
//
func (grc *GoodReviewCreate) OnConflict(opts ...sql.ConflictOption) *GoodReviewUpsertOne {
	grc.conflict = opts
	return &GoodReviewUpsertOne{
		create: grc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoodReview.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (grc *GoodReviewCreate) OnConflictColumns(columns ...string) *GoodReviewUpsertOne {
	grc.conflict = append(grc.conflict, sql.ConflictColumns(columns...))
	return &GoodReviewUpsertOne{
		create: grc,
	}
}

type (
	// GoodReviewUpsertOne is the builder for "upsert"-ing
	//  one GoodReview node.
	GoodReviewUpsertOne struct {
		create *GoodReviewCreate
	}

	// GoodReviewUpsert is the "OnConflict" setter.
	GoodReviewUpsert struct {
		*sql.UpdateSet
	}
)

// SetType sets the "type" field.
func (u *GoodReviewUpsert) SetType(v goodreview.Type) *GoodReviewUpsert {
	u.Set(goodreview.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GoodReviewUpsert) UpdateType() *GoodReviewUpsert {
	u.SetExcluded(goodreview.FieldType)
	return u
}

// SetReviewedID sets the "reviewed_id" field.
func (u *GoodReviewUpsert) SetReviewedID(v uuid.UUID) *GoodReviewUpsert {
	u.Set(goodreview.FieldReviewedID, v)
	return u
}

// UpdateReviewedID sets the "reviewed_id" field to the value that was provided on create.
func (u *GoodReviewUpsert) UpdateReviewedID() *GoodReviewUpsert {
	u.SetExcluded(goodreview.FieldReviewedID)
	return u
}

// SetReviewerID sets the "reviewer_id" field.
func (u *GoodReviewUpsert) SetReviewerID(v uuid.UUID) *GoodReviewUpsert {
	u.Set(goodreview.FieldReviewerID, v)
	return u
}

// UpdateReviewerID sets the "reviewer_id" field to the value that was provided on create.
func (u *GoodReviewUpsert) UpdateReviewerID() *GoodReviewUpsert {
	u.SetExcluded(goodreview.FieldReviewerID)
	return u
}

// SetState sets the "state" field.
func (u *GoodReviewUpsert) SetState(v goodreview.State) *GoodReviewUpsert {
	u.Set(goodreview.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *GoodReviewUpsert) UpdateState() *GoodReviewUpsert {
	u.SetExcluded(goodreview.FieldState)
	return u
}

// SetMessage sets the "message" field.
func (u *GoodReviewUpsert) SetMessage(v string) *GoodReviewUpsert {
	u.Set(goodreview.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *GoodReviewUpsert) UpdateMessage() *GoodReviewUpsert {
	u.SetExcluded(goodreview.FieldMessage)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *GoodReviewUpsert) SetCreateAt(v int64) *GoodReviewUpsert {
	u.Set(goodreview.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *GoodReviewUpsert) UpdateCreateAt() *GoodReviewUpsert {
	u.SetExcluded(goodreview.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *GoodReviewUpsert) SetUpdateAt(v int64) *GoodReviewUpsert {
	u.Set(goodreview.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *GoodReviewUpsert) UpdateUpdateAt() *GoodReviewUpsert {
	u.SetExcluded(goodreview.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *GoodReviewUpsert) SetDeleteAt(v int64) *GoodReviewUpsert {
	u.Set(goodreview.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *GoodReviewUpsert) UpdateDeleteAt() *GoodReviewUpsert {
	u.SetExcluded(goodreview.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GoodReview.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(goodreview.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *GoodReviewUpsertOne) UpdateNewValues() *GoodReviewUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(goodreview.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.GoodReview.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *GoodReviewUpsertOne) Ignore() *GoodReviewUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoodReviewUpsertOne) DoNothing() *GoodReviewUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoodReviewCreate.OnConflict
// documentation for more info.
func (u *GoodReviewUpsertOne) Update(set func(*GoodReviewUpsert)) *GoodReviewUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoodReviewUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *GoodReviewUpsertOne) SetType(v goodreview.Type) *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GoodReviewUpsertOne) UpdateType() *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateType()
	})
}

// SetReviewedID sets the "reviewed_id" field.
func (u *GoodReviewUpsertOne) SetReviewedID(v uuid.UUID) *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetReviewedID(v)
	})
}

// UpdateReviewedID sets the "reviewed_id" field to the value that was provided on create.
func (u *GoodReviewUpsertOne) UpdateReviewedID() *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateReviewedID()
	})
}

// SetReviewerID sets the "reviewer_id" field.
func (u *GoodReviewUpsertOne) SetReviewerID(v uuid.UUID) *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetReviewerID(v)
	})
}

// UpdateReviewerID sets the "reviewer_id" field to the value that was provided on create.
func (u *GoodReviewUpsertOne) UpdateReviewerID() *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateReviewerID()
	})
}

// SetState sets the "state" field.
func (u *GoodReviewUpsertOne) SetState(v goodreview.State) *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *GoodReviewUpsertOne) UpdateState() *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateState()
	})
}

// SetMessage sets the "message" field.
func (u *GoodReviewUpsertOne) SetMessage(v string) *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *GoodReviewUpsertOne) UpdateMessage() *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateMessage()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *GoodReviewUpsertOne) SetCreateAt(v int64) *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *GoodReviewUpsertOne) UpdateCreateAt() *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *GoodReviewUpsertOne) SetUpdateAt(v int64) *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *GoodReviewUpsertOne) UpdateUpdateAt() *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *GoodReviewUpsertOne) SetDeleteAt(v int64) *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *GoodReviewUpsertOne) UpdateDeleteAt() *GoodReviewUpsertOne {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *GoodReviewUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoodReviewCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoodReviewUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GoodReviewUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GoodReviewUpsertOne.ID is not supported by MySQL driver. Use GoodReviewUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GoodReviewUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GoodReviewCreateBulk is the builder for creating many GoodReview entities in bulk.
type GoodReviewCreateBulk struct {
	config
	builders []*GoodReviewCreate
	conflict []sql.ConflictOption
}

// Save creates the GoodReview entities in the database.
func (grcb *GoodReviewCreateBulk) Save(ctx context.Context) ([]*GoodReview, error) {
	specs := make([]*sqlgraph.CreateSpec, len(grcb.builders))
	nodes := make([]*GoodReview, len(grcb.builders))
	mutators := make([]Mutator, len(grcb.builders))
	for i := range grcb.builders {
		func(i int, root context.Context) {
			builder := grcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodReviewMutation)
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
					_, err = mutators[i+1].Mutate(root, grcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = grcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, grcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, grcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (grcb *GoodReviewCreateBulk) SaveX(ctx context.Context) []*GoodReview {
	v, err := grcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (grcb *GoodReviewCreateBulk) Exec(ctx context.Context) error {
	_, err := grcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grcb *GoodReviewCreateBulk) ExecX(ctx context.Context) {
	if err := grcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoodReview.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoodReviewUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
//
func (grcb *GoodReviewCreateBulk) OnConflict(opts ...sql.ConflictOption) *GoodReviewUpsertBulk {
	grcb.conflict = opts
	return &GoodReviewUpsertBulk{
		create: grcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoodReview.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (grcb *GoodReviewCreateBulk) OnConflictColumns(columns ...string) *GoodReviewUpsertBulk {
	grcb.conflict = append(grcb.conflict, sql.ConflictColumns(columns...))
	return &GoodReviewUpsertBulk{
		create: grcb,
	}
}

// GoodReviewUpsertBulk is the builder for "upsert"-ing
// a bulk of GoodReview nodes.
type GoodReviewUpsertBulk struct {
	create *GoodReviewCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GoodReview.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(goodreview.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *GoodReviewUpsertBulk) UpdateNewValues() *GoodReviewUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(goodreview.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoodReview.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *GoodReviewUpsertBulk) Ignore() *GoodReviewUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoodReviewUpsertBulk) DoNothing() *GoodReviewUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoodReviewCreateBulk.OnConflict
// documentation for more info.
func (u *GoodReviewUpsertBulk) Update(set func(*GoodReviewUpsert)) *GoodReviewUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoodReviewUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *GoodReviewUpsertBulk) SetType(v goodreview.Type) *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GoodReviewUpsertBulk) UpdateType() *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateType()
	})
}

// SetReviewedID sets the "reviewed_id" field.
func (u *GoodReviewUpsertBulk) SetReviewedID(v uuid.UUID) *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetReviewedID(v)
	})
}

// UpdateReviewedID sets the "reviewed_id" field to the value that was provided on create.
func (u *GoodReviewUpsertBulk) UpdateReviewedID() *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateReviewedID()
	})
}

// SetReviewerID sets the "reviewer_id" field.
func (u *GoodReviewUpsertBulk) SetReviewerID(v uuid.UUID) *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetReviewerID(v)
	})
}

// UpdateReviewerID sets the "reviewer_id" field to the value that was provided on create.
func (u *GoodReviewUpsertBulk) UpdateReviewerID() *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateReviewerID()
	})
}

// SetState sets the "state" field.
func (u *GoodReviewUpsertBulk) SetState(v goodreview.State) *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *GoodReviewUpsertBulk) UpdateState() *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateState()
	})
}

// SetMessage sets the "message" field.
func (u *GoodReviewUpsertBulk) SetMessage(v string) *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *GoodReviewUpsertBulk) UpdateMessage() *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateMessage()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *GoodReviewUpsertBulk) SetCreateAt(v int64) *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *GoodReviewUpsertBulk) UpdateCreateAt() *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *GoodReviewUpsertBulk) SetUpdateAt(v int64) *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *GoodReviewUpsertBulk) UpdateUpdateAt() *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *GoodReviewUpsertBulk) SetDeleteAt(v int64) *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *GoodReviewUpsertBulk) UpdateDeleteAt() *GoodReviewUpsertBulk {
	return u.Update(func(s *GoodReviewUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *GoodReviewUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GoodReviewCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoodReviewCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoodReviewUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
