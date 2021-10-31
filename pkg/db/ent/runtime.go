// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodinfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/schema"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	goodinfoFields := schema.GoodInfo{}.Fields()
	_ = goodinfoFields
	// goodinfoDescGasPrice is the schema descriptor for gas_price field.
	goodinfoDescGasPrice := goodinfoFields[2].Descriptor()
	// goodinfo.GasPriceValidator is a validator for the "gas_price" field. It is called by the builders before save.
	goodinfo.GasPriceValidator = goodinfoDescGasPrice.Validators[0].(func(int) error)
	// goodinfoDescUnitPower is the schema descriptor for unit_power field.
	goodinfoDescUnitPower := goodinfoFields[4].Descriptor()
	// goodinfo.UnitPowerValidator is a validator for the "unit_power" field. It is called by the builders before save.
	goodinfo.UnitPowerValidator = goodinfoDescUnitPower.Validators[0].(func(float64) error)
	// goodinfoDescDuration is the schema descriptor for duration field.
	goodinfoDescDuration := goodinfoFields[5].Descriptor()
	// goodinfo.DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	goodinfo.DurationValidator = goodinfoDescDuration.Validators[0].(func(int) error)
	// goodinfoDescPrice is the schema descriptor for price field.
	goodinfoDescPrice := goodinfoFields[11].Descriptor()
	// goodinfo.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	goodinfo.PriceValidator = goodinfoDescPrice.Validators[0].(func(int) error)
	// goodinfoDescTotal is the schema descriptor for total field.
	goodinfoDescTotal := goodinfoFields[17].Descriptor()
	// goodinfo.TotalValidator is a validator for the "total" field. It is called by the builders before save.
	goodinfo.TotalValidator = goodinfoDescTotal.Validators[0].(func(int) error)
	// goodinfoDescID is the schema descriptor for id field.
	goodinfoDescID := goodinfoFields[0].Descriptor()
	// goodinfo.DefaultID holds the default value on creation for the id field.
	goodinfo.DefaultID = goodinfoDescID.Default.(func() uuid.UUID)
	targetareaFields := schema.TargetArea{}.Fields()
	_ = targetareaFields
	// targetareaDescID is the schema descriptor for id field.
	targetareaDescID := targetareaFields[0].Descriptor()
	// targetarea.DefaultID holds the default value on creation for the id field.
	targetarea.DefaultID = targetareaDescID.Default.(func() uuid.UUID)
}
