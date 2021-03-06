package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"

	constant "github.com/NpoolPlatform/cloud-hashing-goods/pkg/const"
)

// GoodInfo holds the schema definition for the GoodInfo entity.
type GoodInfo struct {
	ent.Schema
}

// Fields of the GoodInfo.
func (GoodInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("device_info_id", uuid.UUID{}),
		field.Bool("separate_fee"),
		field.Int32("unit_power").
			Positive(),
		field.Int32("duration_days").
			Positive(),
		field.UUID("coin_info_id", uuid.UUID{}),
		field.Bool("actuals"),
		field.Uint32("delivery_at"),
		field.UUID("inherit_from_good_id", uuid.UUID{}),
		field.UUID("vendor_location_id", uuid.UUID{}),
		field.Uint64("price").
			Positive(),
		field.UUID("price_currency", uuid.UUID{}),
		field.Enum("benefit_type").
			Values(constant.BenefitTypePool, constant.BenefitTypePlatform),
		field.Bool("classic"),
		field.String("title"),
		field.String("unit"),
		field.JSON("support_coin_type_ids", []uuid.UUID{}),
		field.JSON("fee_ids", []uuid.UUID{}),
		field.Uint32("start_at"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the GoodInfo.
func (GoodInfo) Edges() []ent.Edge {
	return nil
}
