// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppGoodsColumns holds the columns for the "app_goods" table.
	AppGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "authorized", Type: field.TypeBool, Default: false},
		{Name: "online", Type: field.TypeBool, Default: false},
		{Name: "init_area_strategy", Type: field.TypeEnum, Enums: []string{"all", "none"}},
		{Name: "price", Type: field.TypeUint64},
		{Name: "invitation_only", Type: field.TypeBool, Default: false},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// AppGoodsTable holds the schema information for the "app_goods" table.
	AppGoodsTable = &schema.Table{
		Name:       "app_goods",
		Columns:    AppGoodsColumns,
		PrimaryKey: []*schema.Column{AppGoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appgood_good_id_app_id",
				Unique:  true,
				Columns: []*schema.Column{AppGoodsColumns[2], AppGoodsColumns[1]},
			},
		},
	}
	// AppGoodTargetAreasColumns holds the columns for the "app_good_target_areas" table.
	AppGoodTargetAreasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "target_area_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// AppGoodTargetAreasTable holds the schema information for the "app_good_target_areas" table.
	AppGoodTargetAreasTable = &schema.Table{
		Name:       "app_good_target_areas",
		Columns:    AppGoodTargetAreasColumns,
		PrimaryKey: []*schema.Column{AppGoodTargetAreasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appgoodtargetarea_good_id_app_id_target_area_id",
				Unique:  true,
				Columns: []*schema.Column{AppGoodTargetAreasColumns[2], AppGoodTargetAreasColumns[1], AppGoodTargetAreasColumns[3]},
			},
		},
	}
	// AppTargetAreasColumns holds the columns for the "app_target_areas" table.
	AppTargetAreasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "target_area_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// AppTargetAreasTable holds the schema information for the "app_target_areas" table.
	AppTargetAreasTable = &schema.Table{
		Name:       "app_target_areas",
		Columns:    AppTargetAreasColumns,
		PrimaryKey: []*schema.Column{AppTargetAreasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "apptargetarea_target_area_id_app_id",
				Unique:  true,
				Columns: []*schema.Column{AppTargetAreasColumns[1], AppTargetAreasColumns[2]},
			},
		},
	}
	// DeviceInfosColumns holds the columns for the "device_infos" table.
	DeviceInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "type", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "manufacturer", Type: field.TypeString, Size: 64, Default: ""},
		{Name: "power_comsuption", Type: field.TypeInt32},
		{Name: "shipment_at", Type: field.TypeInt32},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// DeviceInfosTable holds the schema information for the "device_infos" table.
	DeviceInfosTable = &schema.Table{
		Name:       "device_infos",
		Columns:    DeviceInfosColumns,
		PrimaryKey: []*schema.Column{DeviceInfosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "deviceinfo_type_manufacturer_shipment_at_power_comsuption",
				Unique:  true,
				Columns: []*schema.Column{DeviceInfosColumns[1], DeviceInfosColumns[2], DeviceInfosColumns[4], DeviceInfosColumns[3]},
			},
		},
	}
	// FeesColumns holds the columns for the "fees" table.
	FeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "fee_type_id", Type: field.TypeUUID},
		{Name: "value", Type: field.TypeUint64},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// FeesTable holds the schema information for the "fees" table.
	FeesTable = &schema.Table{
		Name:       "fees",
		Columns:    FeesColumns,
		PrimaryKey: []*schema.Column{FeesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "fee_app_id_fee_type_id_value",
				Unique:  true,
				Columns: []*schema.Column{FeesColumns[1], FeesColumns[2], FeesColumns[3]},
			},
		},
	}
	// FeeTypesColumns holds the columns for the "fee_types" table.
	FeeTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "fee_type", Type: field.TypeString, Unique: true},
		{Name: "fee_description", Type: field.TypeString, Size: 256},
		{Name: "pay_type", Type: field.TypeEnum, Enums: []string{"percent", "amount"}},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// FeeTypesTable holds the schema information for the "fee_types" table.
	FeeTypesTable = &schema.Table{
		Name:       "fee_types",
		Columns:    FeeTypesColumns,
		PrimaryKey: []*schema.Column{FeeTypesColumns[0]},
	}
	// GoodCommentsColumns holds the columns for the "good_comments" table.
	GoodCommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "reply_to_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "content", Type: field.TypeString, Default: ""},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// GoodCommentsTable holds the schema information for the "good_comments" table.
	GoodCommentsTable = &schema.Table{
		Name:       "good_comments",
		Columns:    GoodCommentsColumns,
		PrimaryKey: []*schema.Column{GoodCommentsColumns[0]},
	}
	// GoodExtraInfosColumns holds the columns for the "good_extra_infos" table.
	GoodExtraInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID, Unique: true},
		{Name: "posters", Type: field.TypeJSON},
		{Name: "labels", Type: field.TypeJSON},
		{Name: "out_sale", Type: field.TypeBool},
		{Name: "pre_sale", Type: field.TypeBool},
		{Name: "vote_count", Type: field.TypeUint32},
		{Name: "rating", Type: field.TypeFloat32},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// GoodExtraInfosTable holds the schema information for the "good_extra_infos" table.
	GoodExtraInfosTable = &schema.Table{
		Name:       "good_extra_infos",
		Columns:    GoodExtraInfosColumns,
		PrimaryKey: []*schema.Column{GoodExtraInfosColumns[0]},
	}
	// GoodInfosColumns holds the columns for the "good_infos" table.
	GoodInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "device_info_id", Type: field.TypeUUID},
		{Name: "separate_fee", Type: field.TypeBool},
		{Name: "unit_power", Type: field.TypeInt32},
		{Name: "duration_days", Type: field.TypeInt32},
		{Name: "coin_info_id", Type: field.TypeUUID},
		{Name: "actuals", Type: field.TypeBool},
		{Name: "delivery_at", Type: field.TypeUint32},
		{Name: "inherit_from_good_id", Type: field.TypeUUID},
		{Name: "vendor_location_id", Type: field.TypeUUID},
		{Name: "price", Type: field.TypeUint64},
		{Name: "price_currency", Type: field.TypeUUID},
		{Name: "benefit_type", Type: field.TypeEnum, Enums: []string{"pool", "platform"}},
		{Name: "classic", Type: field.TypeBool},
		{Name: "title", Type: field.TypeString},
		{Name: "unit", Type: field.TypeString},
		{Name: "support_coin_type_ids", Type: field.TypeJSON},
		{Name: "fee_ids", Type: field.TypeJSON},
		{Name: "total", Type: field.TypeInt32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// GoodInfosTable holds the schema information for the "good_infos" table.
	GoodInfosTable = &schema.Table{
		Name:       "good_infos",
		Columns:    GoodInfosColumns,
		PrimaryKey: []*schema.Column{GoodInfosColumns[0]},
	}
	// GoodReviewsColumns holds the columns for the "good_reviews" table.
	GoodReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "entity_type", Type: field.TypeEnum, Enums: []string{"good", "appgood", "apptargetarea", "appgoodtargetarea"}},
		{Name: "reviewed_id", Type: field.TypeUUID},
		{Name: "reviewer_id", Type: field.TypeUUID},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"approved", "rejected", "none"}, Default: "none"},
		{Name: "message", Type: field.TypeString, Default: ""},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// GoodReviewsTable holds the schema information for the "good_reviews" table.
	GoodReviewsTable = &schema.Table{
		Name:       "good_reviews",
		Columns:    GoodReviewsColumns,
		PrimaryKey: []*schema.Column{GoodReviewsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "goodreview_entity_type_reviewed_id",
				Unique:  true,
				Columns: []*schema.Column{GoodReviewsColumns[1], GoodReviewsColumns[2]},
			},
		},
	}
	// PriceCurrenciesColumns holds the columns for the "price_currencies" table.
	PriceCurrenciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "unit", Type: field.TypeString},
		{Name: "symbol", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// PriceCurrenciesTable holds the schema information for the "price_currencies" table.
	PriceCurrenciesTable = &schema.Table{
		Name:       "price_currencies",
		Columns:    PriceCurrenciesColumns,
		PrimaryKey: []*schema.Column{PriceCurrenciesColumns[0]},
	}
	// RecommendsColumns holds the columns for the "recommends" table.
	RecommendsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "recommender_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Default: ""},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// RecommendsTable holds the schema information for the "recommends" table.
	RecommendsTable = &schema.Table{
		Name:       "recommends",
		Columns:    RecommendsColumns,
		PrimaryKey: []*schema.Column{RecommendsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "recommend_app_id_good_id",
				Unique:  false,
				Columns: []*schema.Column{RecommendsColumns[1], RecommendsColumns[2]},
			},
		},
	}
	// TargetAreasColumns holds the columns for the "target_areas" table.
	TargetAreasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "continent", Type: field.TypeString, Size: 128, Default: ""},
		{Name: "country", Type: field.TypeString, Size: 128, Default: ""},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// TargetAreasTable holds the schema information for the "target_areas" table.
	TargetAreasTable = &schema.Table{
		Name:       "target_areas",
		Columns:    TargetAreasColumns,
		PrimaryKey: []*schema.Column{TargetAreasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "targetarea_continent_country",
				Unique:  true,
				Columns: []*schema.Column{TargetAreasColumns[1], TargetAreasColumns[2]},
			},
		},
	}
	// VendorLocationsColumns holds the columns for the "vendor_locations" table.
	VendorLocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "country", Type: field.TypeString, Size: 128, Default: ""},
		{Name: "province", Type: field.TypeString, Size: 128, Default: ""},
		{Name: "city", Type: field.TypeString, Size: 128, Default: ""},
		{Name: "address", Type: field.TypeString, Size: 256, Default: ""},
		{Name: "create_at", Type: field.TypeInt64},
		{Name: "update_at", Type: field.TypeInt64},
		{Name: "delete_at", Type: field.TypeInt64},
	}
	// VendorLocationsTable holds the schema information for the "vendor_locations" table.
	VendorLocationsTable = &schema.Table{
		Name:       "vendor_locations",
		Columns:    VendorLocationsColumns,
		PrimaryKey: []*schema.Column{VendorLocationsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vendorlocation_country_province_city_address",
				Unique:  true,
				Columns: []*schema.Column{VendorLocationsColumns[1], VendorLocationsColumns[2], VendorLocationsColumns[3], VendorLocationsColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppGoodsTable,
		AppGoodTargetAreasTable,
		AppTargetAreasTable,
		DeviceInfosTable,
		FeesTable,
		FeeTypesTable,
		GoodCommentsTable,
		GoodExtraInfosTable,
		GoodInfosTable,
		GoodReviewsTable,
		PriceCurrenciesTable,
		RecommendsTable,
		TargetAreasTable,
		VendorLocationsTable,
	}
)

func init() {
}
