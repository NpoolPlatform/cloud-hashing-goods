// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/goodextrainfo"
	"github.com/google/uuid"
)

// GoodExtraInfo is the model entity for the GoodExtraInfo schema.
type GoodExtraInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// Posters holds the value of the "posters" field.
	Posters []string `json:"posters,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt int64 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt int64 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodExtraInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodextrainfo.FieldPosters:
			values[i] = new([]byte)
		case goodextrainfo.FieldCreateAt, goodextrainfo.FieldUpdateAt, goodextrainfo.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case goodextrainfo.FieldID, goodextrainfo.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodExtraInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodExtraInfo fields.
func (gei *GoodExtraInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodextrainfo.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gei.ID = *value
			}
		case goodextrainfo.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				gei.GoodID = *value
			}
		case goodextrainfo.FieldPosters:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field posters", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gei.Posters); err != nil {
					return fmt.Errorf("unmarshal field posters: %w", err)
				}
			}
		case goodextrainfo.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				gei.CreateAt = value.Int64
			}
		case goodextrainfo.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				gei.UpdateAt = value.Int64
			}
		case goodextrainfo.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				gei.DeleteAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodExtraInfo.
// Note that you need to call GoodExtraInfo.Unwrap() before calling this method if this GoodExtraInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (gei *GoodExtraInfo) Update() *GoodExtraInfoUpdateOne {
	return (&GoodExtraInfoClient{config: gei.config}).UpdateOne(gei)
}

// Unwrap unwraps the GoodExtraInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gei *GoodExtraInfo) Unwrap() *GoodExtraInfo {
	tx, ok := gei.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodExtraInfo is not a transactional entity")
	}
	gei.config.driver = tx.drv
	return gei
}

// String implements the fmt.Stringer.
func (gei *GoodExtraInfo) String() string {
	var builder strings.Builder
	builder.WriteString("GoodExtraInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", gei.ID))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", gei.GoodID))
	builder.WriteString(", posters=")
	builder.WriteString(fmt.Sprintf("%v", gei.Posters))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", gei.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", gei.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", gei.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// GoodExtraInfos is a parsable slice of GoodExtraInfo.
type GoodExtraInfos []*GoodExtraInfo

func (gei GoodExtraInfos) config(cfg config) {
	for _i := range gei {
		gei[_i].config = cfg
	}
}