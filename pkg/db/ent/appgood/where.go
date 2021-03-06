// Code generated by entc, DO NOT EDIT.

package appgood

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// Online applies equality check predicate on the "online" field. It's identical to OnlineEQ.
func Online(v bool) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnline), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v uint64) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// DisplayIndex applies equality check predicate on the "display_index" field. It's identical to DisplayIndexEQ.
func DisplayIndex(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayIndex), v))
	})
}

// Visible applies equality check predicate on the "visible" field. It's identical to VisibleEQ.
func Visible(v bool) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisible), v))
	})
}

// PurchaseLimit applies equality check predicate on the "purchase_limit" field. It's identical to PurchaseLimitEQ.
func PurchaseLimit(v int32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPurchaseLimit), v))
	})
}

// CommissionPercent applies equality check predicate on the "commission_percent" field. It's identical to CommissionPercentEQ.
func CommissionPercent(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommissionPercent), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// OnlineEQ applies the EQ predicate on the "online" field.
func OnlineEQ(v bool) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnline), v))
	})
}

// OnlineNEQ applies the NEQ predicate on the "online" field.
func OnlineNEQ(v bool) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOnline), v))
	})
}

// InitAreaStrategyEQ applies the EQ predicate on the "init_area_strategy" field.
func InitAreaStrategyEQ(v InitAreaStrategy) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInitAreaStrategy), v))
	})
}

// InitAreaStrategyNEQ applies the NEQ predicate on the "init_area_strategy" field.
func InitAreaStrategyNEQ(v InitAreaStrategy) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInitAreaStrategy), v))
	})
}

// InitAreaStrategyIn applies the In predicate on the "init_area_strategy" field.
func InitAreaStrategyIn(vs ...InitAreaStrategy) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInitAreaStrategy), v...))
	})
}

// InitAreaStrategyNotIn applies the NotIn predicate on the "init_area_strategy" field.
func InitAreaStrategyNotIn(vs ...InitAreaStrategy) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInitAreaStrategy), v...))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v uint64) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v uint64) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...uint64) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...uint64) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v uint64) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v uint64) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v uint64) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v uint64) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// DisplayIndexEQ applies the EQ predicate on the "display_index" field.
func DisplayIndexEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDisplayIndex), v))
	})
}

// DisplayIndexNEQ applies the NEQ predicate on the "display_index" field.
func DisplayIndexNEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDisplayIndex), v))
	})
}

// DisplayIndexIn applies the In predicate on the "display_index" field.
func DisplayIndexIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDisplayIndex), v...))
	})
}

// DisplayIndexNotIn applies the NotIn predicate on the "display_index" field.
func DisplayIndexNotIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDisplayIndex), v...))
	})
}

// DisplayIndexGT applies the GT predicate on the "display_index" field.
func DisplayIndexGT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDisplayIndex), v))
	})
}

// DisplayIndexGTE applies the GTE predicate on the "display_index" field.
func DisplayIndexGTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDisplayIndex), v))
	})
}

// DisplayIndexLT applies the LT predicate on the "display_index" field.
func DisplayIndexLT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDisplayIndex), v))
	})
}

// DisplayIndexLTE applies the LTE predicate on the "display_index" field.
func DisplayIndexLTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDisplayIndex), v))
	})
}

// VisibleEQ applies the EQ predicate on the "visible" field.
func VisibleEQ(v bool) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisible), v))
	})
}

// VisibleNEQ applies the NEQ predicate on the "visible" field.
func VisibleNEQ(v bool) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVisible), v))
	})
}

// PurchaseLimitEQ applies the EQ predicate on the "purchase_limit" field.
func PurchaseLimitEQ(v int32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPurchaseLimit), v))
	})
}

// PurchaseLimitNEQ applies the NEQ predicate on the "purchase_limit" field.
func PurchaseLimitNEQ(v int32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPurchaseLimit), v))
	})
}

// PurchaseLimitIn applies the In predicate on the "purchase_limit" field.
func PurchaseLimitIn(vs ...int32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPurchaseLimit), v...))
	})
}

// PurchaseLimitNotIn applies the NotIn predicate on the "purchase_limit" field.
func PurchaseLimitNotIn(vs ...int32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPurchaseLimit), v...))
	})
}

// PurchaseLimitGT applies the GT predicate on the "purchase_limit" field.
func PurchaseLimitGT(v int32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPurchaseLimit), v))
	})
}

// PurchaseLimitGTE applies the GTE predicate on the "purchase_limit" field.
func PurchaseLimitGTE(v int32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPurchaseLimit), v))
	})
}

// PurchaseLimitLT applies the LT predicate on the "purchase_limit" field.
func PurchaseLimitLT(v int32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPurchaseLimit), v))
	})
}

// PurchaseLimitLTE applies the LTE predicate on the "purchase_limit" field.
func PurchaseLimitLTE(v int32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPurchaseLimit), v))
	})
}

// CommissionPercentEQ applies the EQ predicate on the "commission_percent" field.
func CommissionPercentEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommissionPercent), v))
	})
}

// CommissionPercentNEQ applies the NEQ predicate on the "commission_percent" field.
func CommissionPercentNEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCommissionPercent), v))
	})
}

// CommissionPercentIn applies the In predicate on the "commission_percent" field.
func CommissionPercentIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCommissionPercent), v...))
	})
}

// CommissionPercentNotIn applies the NotIn predicate on the "commission_percent" field.
func CommissionPercentNotIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCommissionPercent), v...))
	})
}

// CommissionPercentGT applies the GT predicate on the "commission_percent" field.
func CommissionPercentGT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCommissionPercent), v))
	})
}

// CommissionPercentGTE applies the GTE predicate on the "commission_percent" field.
func CommissionPercentGTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCommissionPercent), v))
	})
}

// CommissionPercentLT applies the LT predicate on the "commission_percent" field.
func CommissionPercentLT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCommissionPercent), v))
	})
}

// CommissionPercentLTE applies the LTE predicate on the "commission_percent" field.
func CommissionPercentLTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCommissionPercent), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.AppGood {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppGood(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppGood) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppGood) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppGood) predicate.AppGood {
	return predicate.AppGood(func(s *sql.Selector) {
		p(s.Not())
	})
}
