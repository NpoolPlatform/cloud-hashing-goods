// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/deviceinfo"
	"github.com/google/uuid"
)

// DeviceInfo is the model entity for the DeviceInfo schema.
type DeviceInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Manufacturer holds the value of the "manufacturer" field.
	Manufacturer string `json:"manufacturer,omitempty"`
	// PowerComsuption holds the value of the "power_comsuption" field.
	PowerComsuption int32 `json:"power_comsuption,omitempty"`
	// ShipmentAt holds the value of the "shipment_at" field.
	ShipmentAt int32 `json:"shipment_at,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt int64 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt int64 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeviceInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deviceinfo.FieldPowerComsuption, deviceinfo.FieldShipmentAt, deviceinfo.FieldCreateAt, deviceinfo.FieldUpdateAt, deviceinfo.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case deviceinfo.FieldType, deviceinfo.FieldManufacturer:
			values[i] = new(sql.NullString)
		case deviceinfo.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeviceInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeviceInfo fields.
func (di *DeviceInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deviceinfo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				di.ID = *value
			}
		case deviceinfo.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				di.Type = value.String
			}
		case deviceinfo.FieldManufacturer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer", values[i])
			} else if value.Valid {
				di.Manufacturer = value.String
			}
		case deviceinfo.FieldPowerComsuption:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field power_comsuption", values[i])
			} else if value.Valid {
				di.PowerComsuption = int32(value.Int64)
			}
		case deviceinfo.FieldShipmentAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shipment_at", values[i])
			} else if value.Valid {
				di.ShipmentAt = int32(value.Int64)
			}
		case deviceinfo.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				di.CreateAt = value.Int64
			}
		case deviceinfo.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				di.UpdateAt = value.Int64
			}
		case deviceinfo.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				di.DeleteAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DeviceInfo.
// Note that you need to call DeviceInfo.Unwrap() before calling this method if this DeviceInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (di *DeviceInfo) Update() *DeviceInfoUpdateOne {
	return (&DeviceInfoClient{config: di.config}).UpdateOne(di)
}

// Unwrap unwraps the DeviceInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (di *DeviceInfo) Unwrap() *DeviceInfo {
	tx, ok := di.config.driver.(*txDriver)
	if !ok {
		panic("ent: DeviceInfo is not a transactional entity")
	}
	di.config.driver = tx.drv
	return di
}

// String implements the fmt.Stringer.
func (di *DeviceInfo) String() string {
	var builder strings.Builder
	builder.WriteString("DeviceInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", di.ID))
	builder.WriteString(", type=")
	builder.WriteString(di.Type)
	builder.WriteString(", manufacturer=")
	builder.WriteString(di.Manufacturer)
	builder.WriteString(", power_comsuption=")
	builder.WriteString(fmt.Sprintf("%v", di.PowerComsuption))
	builder.WriteString(", shipment_at=")
	builder.WriteString(fmt.Sprintf("%v", di.ShipmentAt))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", di.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", di.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", di.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// DeviceInfos is a parsable slice of DeviceInfo.
type DeviceInfos []*DeviceInfo

func (di DeviceInfos) config(cfg config) {
	for _i := range di {
		di[_i].config = cfg
	}
}
