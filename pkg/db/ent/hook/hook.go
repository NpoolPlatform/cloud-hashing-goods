// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent"
)

// The AppGoodFunc type is an adapter to allow the use of ordinary
// function as AppGood mutator.
type AppGoodFunc func(context.Context, *ent.AppGoodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodMutation", m)
	}
	return f(ctx, mv)
}

// The AppGoodPromotionFunc type is an adapter to allow the use of ordinary
// function as AppGoodPromotion mutator.
type AppGoodPromotionFunc func(context.Context, *ent.AppGoodPromotionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodPromotionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodPromotionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodPromotionMutation", m)
	}
	return f(ctx, mv)
}

// The AppGoodTargetAreaFunc type is an adapter to allow the use of ordinary
// function as AppGoodTargetArea mutator.
type AppGoodTargetAreaFunc func(context.Context, *ent.AppGoodTargetAreaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodTargetAreaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodTargetAreaMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodTargetAreaMutation", m)
	}
	return f(ctx, mv)
}

// The AppTargetAreaFunc type is an adapter to allow the use of ordinary
// function as AppTargetArea mutator.
type AppTargetAreaFunc func(context.Context, *ent.AppTargetAreaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppTargetAreaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppTargetAreaMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppTargetAreaMutation", m)
	}
	return f(ctx, mv)
}

// The DeviceInfoFunc type is an adapter to allow the use of ordinary
// function as DeviceInfo mutator.
type DeviceInfoFunc func(context.Context, *ent.DeviceInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeviceInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeviceInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeviceInfoMutation", m)
	}
	return f(ctx, mv)
}

// The FeeFunc type is an adapter to allow the use of ordinary
// function as Fee mutator.
type FeeFunc func(context.Context, *ent.FeeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FeeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeeMutation", m)
	}
	return f(ctx, mv)
}

// The FeeTypeFunc type is an adapter to allow the use of ordinary
// function as FeeType mutator.
type FeeTypeFunc func(context.Context, *ent.FeeTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeeTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FeeTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeeTypeMutation", m)
	}
	return f(ctx, mv)
}

// The GoodCommentFunc type is an adapter to allow the use of ordinary
// function as GoodComment mutator.
type GoodCommentFunc func(context.Context, *ent.GoodCommentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodCommentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodCommentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodCommentMutation", m)
	}
	return f(ctx, mv)
}

// The GoodExtraInfoFunc type is an adapter to allow the use of ordinary
// function as GoodExtraInfo mutator.
type GoodExtraInfoFunc func(context.Context, *ent.GoodExtraInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodExtraInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodExtraInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodExtraInfoMutation", m)
	}
	return f(ctx, mv)
}

// The GoodInfoFunc type is an adapter to allow the use of ordinary
// function as GoodInfo mutator.
type GoodInfoFunc func(context.Context, *ent.GoodInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodInfoMutation", m)
	}
	return f(ctx, mv)
}

// The GoodReviewFunc type is an adapter to allow the use of ordinary
// function as GoodReview mutator.
type GoodReviewFunc func(context.Context, *ent.GoodReviewMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodReviewFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodReviewMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodReviewMutation", m)
	}
	return f(ctx, mv)
}

// The PriceCurrencyFunc type is an adapter to allow the use of ordinary
// function as PriceCurrency mutator.
type PriceCurrencyFunc func(context.Context, *ent.PriceCurrencyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PriceCurrencyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PriceCurrencyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PriceCurrencyMutation", m)
	}
	return f(ctx, mv)
}

// The RecommendFunc type is an adapter to allow the use of ordinary
// function as Recommend mutator.
type RecommendFunc func(context.Context, *ent.RecommendMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RecommendFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RecommendMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RecommendMutation", m)
	}
	return f(ctx, mv)
}

// The TargetAreaFunc type is an adapter to allow the use of ordinary
// function as TargetArea mutator.
type TargetAreaFunc func(context.Context, *ent.TargetAreaMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TargetAreaFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TargetAreaMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TargetAreaMutation", m)
	}
	return f(ctx, mv)
}

// The VendorLocationFunc type is an adapter to allow the use of ordinary
// function as VendorLocation mutator.
type VendorLocationFunc func(context.Context, *ent.VendorLocationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VendorLocationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.VendorLocationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VendorLocationMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
