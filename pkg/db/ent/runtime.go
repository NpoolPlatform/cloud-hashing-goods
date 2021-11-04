// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgood"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/appgoodtargetarea"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/apptargetarea"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/deviceinfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodcomment"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodextrainfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodinfo"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodreview"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/schema"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/vendorlocation"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appgoodFields := schema.AppGood{}.Fields()
	_ = appgoodFields
	// appgoodDescAuthorized is the schema descriptor for authorized field.
	appgoodDescAuthorized := appgoodFields[3].Descriptor()
	// appgood.DefaultAuthorized holds the default value on creation for the authorized field.
	appgood.DefaultAuthorized = appgoodDescAuthorized.Default.(bool)
	// appgoodDescOnline is the schema descriptor for online field.
	appgoodDescOnline := appgoodFields[4].Descriptor()
	// appgood.DefaultOnline holds the default value on creation for the online field.
	appgood.DefaultOnline = appgoodDescOnline.Default.(bool)
	// appgoodDescInvitationOnly is the schema descriptor for invitation_only field.
	appgoodDescInvitationOnly := appgoodFields[8].Descriptor()
	// appgood.DefaultInvitationOnly holds the default value on creation for the invitation_only field.
	appgood.DefaultInvitationOnly = appgoodDescInvitationOnly.Default.(bool)
	// appgoodDescCreateAt is the schema descriptor for create_at field.
	appgoodDescCreateAt := appgoodFields[9].Descriptor()
	// appgood.DefaultCreateAt holds the default value on creation for the create_at field.
	appgood.DefaultCreateAt = appgoodDescCreateAt.Default.(func() int64)
	// appgoodDescUpdateAt is the schema descriptor for update_at field.
	appgoodDescUpdateAt := appgoodFields[10].Descriptor()
	// appgood.DefaultUpdateAt holds the default value on creation for the update_at field.
	appgood.DefaultUpdateAt = appgoodDescUpdateAt.Default.(func() int64)
	// appgood.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appgood.UpdateDefaultUpdateAt = appgoodDescUpdateAt.UpdateDefault.(func() int64)
	// appgoodDescDeleteAt is the schema descriptor for delete_at field.
	appgoodDescDeleteAt := appgoodFields[11].Descriptor()
	// appgood.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appgood.DefaultDeleteAt = appgoodDescDeleteAt.Default.(func() int64)
	// appgoodDescID is the schema descriptor for id field.
	appgoodDescID := appgoodFields[0].Descriptor()
	// appgood.DefaultID holds the default value on creation for the id field.
	appgood.DefaultID = appgoodDescID.Default.(func() uuid.UUID)
	appgoodtargetareaFields := schema.AppGoodTargetArea{}.Fields()
	_ = appgoodtargetareaFields
	// appgoodtargetareaDescCreateAt is the schema descriptor for create_at field.
	appgoodtargetareaDescCreateAt := appgoodtargetareaFields[4].Descriptor()
	// appgoodtargetarea.DefaultCreateAt holds the default value on creation for the create_at field.
	appgoodtargetarea.DefaultCreateAt = appgoodtargetareaDescCreateAt.Default.(func() int64)
	// appgoodtargetareaDescUpdateAt is the schema descriptor for update_at field.
	appgoodtargetareaDescUpdateAt := appgoodtargetareaFields[5].Descriptor()
	// appgoodtargetarea.DefaultUpdateAt holds the default value on creation for the update_at field.
	appgoodtargetarea.DefaultUpdateAt = appgoodtargetareaDescUpdateAt.Default.(func() int64)
	// appgoodtargetarea.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appgoodtargetarea.UpdateDefaultUpdateAt = appgoodtargetareaDescUpdateAt.UpdateDefault.(func() int64)
	// appgoodtargetareaDescDeleteAt is the schema descriptor for delete_at field.
	appgoodtargetareaDescDeleteAt := appgoodtargetareaFields[6].Descriptor()
	// appgoodtargetarea.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appgoodtargetarea.DefaultDeleteAt = appgoodtargetareaDescDeleteAt.Default.(func() int64)
	// appgoodtargetareaDescID is the schema descriptor for id field.
	appgoodtargetareaDescID := appgoodtargetareaFields[0].Descriptor()
	// appgoodtargetarea.DefaultID holds the default value on creation for the id field.
	appgoodtargetarea.DefaultID = appgoodtargetareaDescID.Default.(func() uuid.UUID)
	apptargetareaFields := schema.AppTargetArea{}.Fields()
	_ = apptargetareaFields
	// apptargetareaDescCreateAt is the schema descriptor for create_at field.
	apptargetareaDescCreateAt := apptargetareaFields[3].Descriptor()
	// apptargetarea.DefaultCreateAt holds the default value on creation for the create_at field.
	apptargetarea.DefaultCreateAt = apptargetareaDescCreateAt.Default.(func() int64)
	// apptargetareaDescUpdateAt is the schema descriptor for update_at field.
	apptargetareaDescUpdateAt := apptargetareaFields[4].Descriptor()
	// apptargetarea.DefaultUpdateAt holds the default value on creation for the update_at field.
	apptargetarea.DefaultUpdateAt = apptargetareaDescUpdateAt.Default.(func() int64)
	// apptargetarea.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	apptargetarea.UpdateDefaultUpdateAt = apptargetareaDescUpdateAt.UpdateDefault.(func() int64)
	// apptargetareaDescDeleteAt is the schema descriptor for delete_at field.
	apptargetareaDescDeleteAt := apptargetareaFields[5].Descriptor()
	// apptargetarea.DefaultDeleteAt holds the default value on creation for the delete_at field.
	apptargetarea.DefaultDeleteAt = apptargetareaDescDeleteAt.Default.(func() int64)
	// apptargetareaDescID is the schema descriptor for id field.
	apptargetareaDescID := apptargetareaFields[0].Descriptor()
	// apptargetarea.DefaultID holds the default value on creation for the id field.
	apptargetarea.DefaultID = apptargetareaDescID.Default.(func() uuid.UUID)
	deviceinfoFields := schema.DeviceInfo{}.Fields()
	_ = deviceinfoFields
	// deviceinfoDescType is the schema descriptor for type field.
	deviceinfoDescType := deviceinfoFields[1].Descriptor()
	// deviceinfo.DefaultType holds the default value on creation for the type field.
	deviceinfo.DefaultType = deviceinfoDescType.Default.(string)
	// deviceinfo.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	deviceinfo.TypeValidator = deviceinfoDescType.Validators[0].(func(string) error)
	// deviceinfoDescManufacturer is the schema descriptor for manufacturer field.
	deviceinfoDescManufacturer := deviceinfoFields[2].Descriptor()
	// deviceinfo.DefaultManufacturer holds the default value on creation for the manufacturer field.
	deviceinfo.DefaultManufacturer = deviceinfoDescManufacturer.Default.(string)
	// deviceinfo.ManufacturerValidator is a validator for the "manufacturer" field. It is called by the builders before save.
	deviceinfo.ManufacturerValidator = deviceinfoDescManufacturer.Validators[0].(func(string) error)
	// deviceinfoDescCreateAt is the schema descriptor for create_at field.
	deviceinfoDescCreateAt := deviceinfoFields[5].Descriptor()
	// deviceinfo.DefaultCreateAt holds the default value on creation for the create_at field.
	deviceinfo.DefaultCreateAt = deviceinfoDescCreateAt.Default.(func() int64)
	// deviceinfoDescUpdateAt is the schema descriptor for update_at field.
	deviceinfoDescUpdateAt := deviceinfoFields[6].Descriptor()
	// deviceinfo.DefaultUpdateAt holds the default value on creation for the update_at field.
	deviceinfo.DefaultUpdateAt = deviceinfoDescUpdateAt.Default.(func() int64)
	// deviceinfo.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	deviceinfo.UpdateDefaultUpdateAt = deviceinfoDescUpdateAt.UpdateDefault.(func() int64)
	// deviceinfoDescDeleteAt is the schema descriptor for delete_at field.
	deviceinfoDescDeleteAt := deviceinfoFields[7].Descriptor()
	// deviceinfo.DefaultDeleteAt holds the default value on creation for the delete_at field.
	deviceinfo.DefaultDeleteAt = deviceinfoDescDeleteAt.Default.(func() int64)
	// deviceinfoDescID is the schema descriptor for id field.
	deviceinfoDescID := deviceinfoFields[0].Descriptor()
	// deviceinfo.DefaultID holds the default value on creation for the id field.
	deviceinfo.DefaultID = deviceinfoDescID.Default.(func() uuid.UUID)
	goodcommentFields := schema.GoodComment{}.Fields()
	_ = goodcommentFields
	// goodcommentDescContent is the schema descriptor for content field.
	goodcommentDescContent := goodcommentFields[6].Descriptor()
	// goodcomment.DefaultContent holds the default value on creation for the content field.
	goodcomment.DefaultContent = goodcommentDescContent.Default.(string)
	// goodcommentDescCreateAt is the schema descriptor for create_at field.
	goodcommentDescCreateAt := goodcommentFields[7].Descriptor()
	// goodcomment.DefaultCreateAt holds the default value on creation for the create_at field.
	goodcomment.DefaultCreateAt = goodcommentDescCreateAt.Default.(func() int64)
	// goodcommentDescUpdateAt is the schema descriptor for update_at field.
	goodcommentDescUpdateAt := goodcommentFields[8].Descriptor()
	// goodcomment.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodcomment.DefaultUpdateAt = goodcommentDescUpdateAt.Default.(func() int64)
	// goodcomment.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodcomment.UpdateDefaultUpdateAt = goodcommentDescUpdateAt.UpdateDefault.(func() int64)
	// goodcommentDescDeleteAt is the schema descriptor for delete_at field.
	goodcommentDescDeleteAt := goodcommentFields[9].Descriptor()
	// goodcomment.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodcomment.DefaultDeleteAt = goodcommentDescDeleteAt.Default.(func() int64)
	// goodcommentDescID is the schema descriptor for id field.
	goodcommentDescID := goodcommentFields[0].Descriptor()
	// goodcomment.DefaultID holds the default value on creation for the id field.
	goodcomment.DefaultID = goodcommentDescID.Default.(func() uuid.UUID)
	goodextrainfoFields := schema.GoodExtraInfo{}.Fields()
	_ = goodextrainfoFields
	// goodextrainfoDescCreateAt is the schema descriptor for create_at field.
	goodextrainfoDescCreateAt := goodextrainfoFields[3].Descriptor()
	// goodextrainfo.DefaultCreateAt holds the default value on creation for the create_at field.
	goodextrainfo.DefaultCreateAt = goodextrainfoDescCreateAt.Default.(func() int64)
	// goodextrainfoDescUpdateAt is the schema descriptor for update_at field.
	goodextrainfoDescUpdateAt := goodextrainfoFields[4].Descriptor()
	// goodextrainfo.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodextrainfo.DefaultUpdateAt = goodextrainfoDescUpdateAt.Default.(func() int64)
	// goodextrainfo.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodextrainfo.UpdateDefaultUpdateAt = goodextrainfoDescUpdateAt.UpdateDefault.(func() int64)
	// goodextrainfoDescDeleteAt is the schema descriptor for delete_at field.
	goodextrainfoDescDeleteAt := goodextrainfoFields[5].Descriptor()
	// goodextrainfo.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodextrainfo.DefaultDeleteAt = goodextrainfoDescDeleteAt.Default.(func() int64)
	// goodextrainfoDescID is the schema descriptor for id field.
	goodextrainfoDescID := goodextrainfoFields[0].Descriptor()
	// goodextrainfo.DefaultID holds the default value on creation for the id field.
	goodextrainfo.DefaultID = goodextrainfoDescID.Default.(func() uuid.UUID)
	goodinfoFields := schema.GoodInfo{}.Fields()
	_ = goodinfoFields
	// goodinfoDescGasPrice is the schema descriptor for gas_price field.
	goodinfoDescGasPrice := goodinfoFields[2].Descriptor()
	// goodinfo.GasPriceValidator is a validator for the "gas_price" field. It is called by the builders before save.
	goodinfo.GasPriceValidator = goodinfoDescGasPrice.Validators[0].(func(uint64) error)
	// goodinfoDescUnitPower is the schema descriptor for unit_power field.
	goodinfoDescUnitPower := goodinfoFields[4].Descriptor()
	// goodinfo.UnitPowerValidator is a validator for the "unit_power" field. It is called by the builders before save.
	goodinfo.UnitPowerValidator = goodinfoDescUnitPower.Validators[0].(func(int32) error)
	// goodinfoDescDurationDays is the schema descriptor for duration_days field.
	goodinfoDescDurationDays := goodinfoFields[5].Descriptor()
	// goodinfo.DurationDaysValidator is a validator for the "duration_days" field. It is called by the builders before save.
	goodinfo.DurationDaysValidator = goodinfoDescDurationDays.Validators[0].(func(int32) error)
	// goodinfoDescPrice is the schema descriptor for price field.
	goodinfoDescPrice := goodinfoFields[11].Descriptor()
	// goodinfo.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	goodinfo.PriceValidator = goodinfoDescPrice.Validators[0].(func(uint64) error)
	// goodinfoDescTotal is the schema descriptor for total field.
	goodinfoDescTotal := goodinfoFields[15].Descriptor()
	// goodinfo.TotalValidator is a validator for the "total" field. It is called by the builders before save.
	goodinfo.TotalValidator = goodinfoDescTotal.Validators[0].(func(int32) error)
	// goodinfoDescCreateAt is the schema descriptor for create_at field.
	goodinfoDescCreateAt := goodinfoFields[16].Descriptor()
	// goodinfo.DefaultCreateAt holds the default value on creation for the create_at field.
	goodinfo.DefaultCreateAt = goodinfoDescCreateAt.Default.(func() int64)
	// goodinfoDescUpdateAt is the schema descriptor for update_at field.
	goodinfoDescUpdateAt := goodinfoFields[17].Descriptor()
	// goodinfo.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodinfo.DefaultUpdateAt = goodinfoDescUpdateAt.Default.(func() int64)
	// goodinfo.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodinfo.UpdateDefaultUpdateAt = goodinfoDescUpdateAt.UpdateDefault.(func() int64)
	// goodinfoDescDeleteAt is the schema descriptor for delete_at field.
	goodinfoDescDeleteAt := goodinfoFields[18].Descriptor()
	// goodinfo.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodinfo.DefaultDeleteAt = goodinfoDescDeleteAt.Default.(func() int64)
	// goodinfoDescID is the schema descriptor for id field.
	goodinfoDescID := goodinfoFields[0].Descriptor()
	// goodinfo.DefaultID holds the default value on creation for the id field.
	goodinfo.DefaultID = goodinfoDescID.Default.(func() uuid.UUID)
	goodreviewFields := schema.GoodReview{}.Fields()
	_ = goodreviewFields
	// goodreviewDescMessage is the schema descriptor for message field.
	goodreviewDescMessage := goodreviewFields[5].Descriptor()
	// goodreview.DefaultMessage holds the default value on creation for the message field.
	goodreview.DefaultMessage = goodreviewDescMessage.Default.(string)
	// goodreviewDescCreateAt is the schema descriptor for create_at field.
	goodreviewDescCreateAt := goodreviewFields[6].Descriptor()
	// goodreview.DefaultCreateAt holds the default value on creation for the create_at field.
	goodreview.DefaultCreateAt = goodreviewDescCreateAt.Default.(func() int64)
	// goodreviewDescUpdateAt is the schema descriptor for update_at field.
	goodreviewDescUpdateAt := goodreviewFields[7].Descriptor()
	// goodreview.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodreview.DefaultUpdateAt = goodreviewDescUpdateAt.Default.(func() int64)
	// goodreview.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodreview.UpdateDefaultUpdateAt = goodreviewDescUpdateAt.UpdateDefault.(func() int64)
	// goodreviewDescDeleteAt is the schema descriptor for delete_at field.
	goodreviewDescDeleteAt := goodreviewFields[8].Descriptor()
	// goodreview.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodreview.DefaultDeleteAt = goodreviewDescDeleteAt.Default.(func() int64)
	// goodreviewDescID is the schema descriptor for id field.
	goodreviewDescID := goodreviewFields[0].Descriptor()
	// goodreview.DefaultID holds the default value on creation for the id field.
	goodreview.DefaultID = goodreviewDescID.Default.(func() uuid.UUID)
	targetareaFields := schema.TargetArea{}.Fields()
	_ = targetareaFields
	// targetareaDescContinent is the schema descriptor for continent field.
	targetareaDescContinent := targetareaFields[1].Descriptor()
	// targetarea.DefaultContinent holds the default value on creation for the continent field.
	targetarea.DefaultContinent = targetareaDescContinent.Default.(string)
	// targetarea.ContinentValidator is a validator for the "continent" field. It is called by the builders before save.
	targetarea.ContinentValidator = targetareaDescContinent.Validators[0].(func(string) error)
	// targetareaDescCountry is the schema descriptor for country field.
	targetareaDescCountry := targetareaFields[2].Descriptor()
	// targetarea.DefaultCountry holds the default value on creation for the country field.
	targetarea.DefaultCountry = targetareaDescCountry.Default.(string)
	// targetarea.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	targetarea.CountryValidator = targetareaDescCountry.Validators[0].(func(string) error)
	// targetareaDescCreateAt is the schema descriptor for create_at field.
	targetareaDescCreateAt := targetareaFields[3].Descriptor()
	// targetarea.DefaultCreateAt holds the default value on creation for the create_at field.
	targetarea.DefaultCreateAt = targetareaDescCreateAt.Default.(func() int64)
	// targetareaDescUpdateAt is the schema descriptor for update_at field.
	targetareaDescUpdateAt := targetareaFields[4].Descriptor()
	// targetarea.DefaultUpdateAt holds the default value on creation for the update_at field.
	targetarea.DefaultUpdateAt = targetareaDescUpdateAt.Default.(func() int64)
	// targetarea.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	targetarea.UpdateDefaultUpdateAt = targetareaDescUpdateAt.UpdateDefault.(func() int64)
	// targetareaDescDeleteAt is the schema descriptor for delete_at field.
	targetareaDescDeleteAt := targetareaFields[5].Descriptor()
	// targetarea.DefaultDeleteAt holds the default value on creation for the delete_at field.
	targetarea.DefaultDeleteAt = targetareaDescDeleteAt.Default.(func() int64)
	// targetareaDescID is the schema descriptor for id field.
	targetareaDescID := targetareaFields[0].Descriptor()
	// targetarea.DefaultID holds the default value on creation for the id field.
	targetarea.DefaultID = targetareaDescID.Default.(func() uuid.UUID)
	vendorlocationFields := schema.VendorLocation{}.Fields()
	_ = vendorlocationFields
	// vendorlocationDescCountry is the schema descriptor for country field.
	vendorlocationDescCountry := vendorlocationFields[1].Descriptor()
	// vendorlocation.DefaultCountry holds the default value on creation for the country field.
	vendorlocation.DefaultCountry = vendorlocationDescCountry.Default.(string)
	// vendorlocation.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	vendorlocation.CountryValidator = vendorlocationDescCountry.Validators[0].(func(string) error)
	// vendorlocationDescProvince is the schema descriptor for province field.
	vendorlocationDescProvince := vendorlocationFields[2].Descriptor()
	// vendorlocation.DefaultProvince holds the default value on creation for the province field.
	vendorlocation.DefaultProvince = vendorlocationDescProvince.Default.(string)
	// vendorlocation.ProvinceValidator is a validator for the "province" field. It is called by the builders before save.
	vendorlocation.ProvinceValidator = vendorlocationDescProvince.Validators[0].(func(string) error)
	// vendorlocationDescCity is the schema descriptor for city field.
	vendorlocationDescCity := vendorlocationFields[3].Descriptor()
	// vendorlocation.DefaultCity holds the default value on creation for the city field.
	vendorlocation.DefaultCity = vendorlocationDescCity.Default.(string)
	// vendorlocation.CityValidator is a validator for the "city" field. It is called by the builders before save.
	vendorlocation.CityValidator = vendorlocationDescCity.Validators[0].(func(string) error)
	// vendorlocationDescAddress is the schema descriptor for address field.
	vendorlocationDescAddress := vendorlocationFields[4].Descriptor()
	// vendorlocation.DefaultAddress holds the default value on creation for the address field.
	vendorlocation.DefaultAddress = vendorlocationDescAddress.Default.(string)
	// vendorlocation.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	vendorlocation.AddressValidator = vendorlocationDescAddress.Validators[0].(func(string) error)
	// vendorlocationDescCreateAt is the schema descriptor for create_at field.
	vendorlocationDescCreateAt := vendorlocationFields[5].Descriptor()
	// vendorlocation.DefaultCreateAt holds the default value on creation for the create_at field.
	vendorlocation.DefaultCreateAt = vendorlocationDescCreateAt.Default.(func() int64)
	// vendorlocationDescUpdateAt is the schema descriptor for update_at field.
	vendorlocationDescUpdateAt := vendorlocationFields[6].Descriptor()
	// vendorlocation.DefaultUpdateAt holds the default value on creation for the update_at field.
	vendorlocation.DefaultUpdateAt = vendorlocationDescUpdateAt.Default.(func() int64)
	// vendorlocation.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	vendorlocation.UpdateDefaultUpdateAt = vendorlocationDescUpdateAt.UpdateDefault.(func() int64)
	// vendorlocationDescDeleteAt is the schema descriptor for delete_at field.
	vendorlocationDescDeleteAt := vendorlocationFields[7].Descriptor()
	// vendorlocation.DefaultDeleteAt holds the default value on creation for the delete_at field.
	vendorlocation.DefaultDeleteAt = vendorlocationDescDeleteAt.Default.(func() int64)
	// vendorlocationDescID is the schema descriptor for id field.
	vendorlocationDescID := vendorlocationFields[0].Descriptor()
	// vendorlocation.DefaultID holds the default value on creation for the id field.
	vendorlocation.DefaultID = vendorlocationDescID.Default.(func() uuid.UUID)
}
