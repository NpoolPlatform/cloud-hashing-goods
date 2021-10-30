// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GoodInfosColumns holds the columns for the "good_infos" table.
	GoodInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "device_info_id", Type: field.TypeUUID},
		{Name: "gas_price", Type: field.TypeInt},
		{Name: "separate_gas_fee", Type: field.TypeBool},
		{Name: "unit_power", Type: field.TypeFloat64},
		{Name: "duration", Type: field.TypeInt},
		{Name: "coin_info_id", Type: field.TypeUUID},
		{Name: "actuals", Type: field.TypeBool},
		{Name: "delivery_time", Type: field.TypeTime},
		{Name: "inherit_from_good_id", Type: field.TypeUUID},
		{Name: "vendor_location_id", Type: field.TypeUUID},
		{Name: "price", Type: field.TypeInt},
		{Name: "benefit_type", Type: field.TypeString},
		{Name: "classic", Type: field.TypeBool},
		{Name: "support_coin_type_ids", Type: field.TypeJSON},
		{Name: "reviewer_id", Type: field.TypeUUID},
		{Name: "state", Type: field.TypeString},
		{Name: "total", Type: field.TypeInt},
	}
	// GoodInfosTable holds the schema information for the "good_infos" table.
	GoodInfosTable = &schema.Table{
		Name:       "good_infos",
		Columns:    GoodInfosColumns,
		PrimaryKey: []*schema.Column{GoodInfosColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GoodInfosTable,
	}
)

func init() {
}
