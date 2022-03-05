// Code generated by entc, DO NOT EDIT.

package appgood

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appgood type in the database.
	Label = "app_good"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldOnline holds the string denoting the online field in the database.
	FieldOnline = "online"
	// FieldInitAreaStrategy holds the string denoting the init_area_strategy field in the database.
	FieldInitAreaStrategy = "init_area_strategy"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldDisplayIndex holds the string denoting the display_index field in the database.
	FieldDisplayIndex = "display_index"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the appgood in the database.
	Table = "app_goods"
)

// Columns holds all SQL columns for appgood fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldGoodID,
	FieldOnline,
	FieldInitAreaStrategy,
	FieldPrice,
	FieldDisplayIndex,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultOnline holds the default value on creation for the "online" field.
	DefaultOnline bool
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// InitAreaStrategy defines the type for the "init_area_strategy" enum field.
type InitAreaStrategy string

// InitAreaStrategy values.
const (
	InitAreaStrategyAll  InitAreaStrategy = "all"
	InitAreaStrategyNone InitAreaStrategy = "none"
)

func (ias InitAreaStrategy) String() string {
	return string(ias)
}

// InitAreaStrategyValidator is a validator for the "init_area_strategy" field enum values. It is called by the builders before save.
func InitAreaStrategyValidator(ias InitAreaStrategy) error {
	switch ias {
	case InitAreaStrategyAll, InitAreaStrategyNone:
		return nil
	default:
		return fmt.Errorf("appgood: invalid enum value for init_area_strategy field: %q", ias)
	}
}
