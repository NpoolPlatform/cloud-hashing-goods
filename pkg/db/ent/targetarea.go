// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-goods/pkg/db/ent/targetarea"
	"github.com/google/uuid"
)

// TargetArea is the model entity for the TargetArea schema.
type TargetArea struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Continent holds the value of the "continent" field.
	Continent string `json:"continent,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt int64 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt int64 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt int64 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TargetArea) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case targetarea.FieldCreateAt, targetarea.FieldUpdateAt, targetarea.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case targetarea.FieldContinent, targetarea.FieldCountry:
			values[i] = new(sql.NullString)
		case targetarea.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TargetArea", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TargetArea fields.
func (ta *TargetArea) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case targetarea.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ta.ID = *value
			}
		case targetarea.FieldContinent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field continent", values[i])
			} else if value.Valid {
				ta.Continent = value.String
			}
		case targetarea.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				ta.Country = value.String
			}
		case targetarea.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ta.CreateAt = value.Int64
			}
		case targetarea.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ta.UpdateAt = value.Int64
			}
		case targetarea.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ta.DeleteAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TargetArea.
// Note that you need to call TargetArea.Unwrap() before calling this method if this TargetArea
// was returned from a transaction, and the transaction was committed or rolled back.
func (ta *TargetArea) Update() *TargetAreaUpdateOne {
	return (&TargetAreaClient{config: ta.config}).UpdateOne(ta)
}

// Unwrap unwraps the TargetArea entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ta *TargetArea) Unwrap() *TargetArea {
	tx, ok := ta.config.driver.(*txDriver)
	if !ok {
		panic("ent: TargetArea is not a transactional entity")
	}
	ta.config.driver = tx.drv
	return ta
}

// String implements the fmt.Stringer.
func (ta *TargetArea) String() string {
	var builder strings.Builder
	builder.WriteString("TargetArea(")
	builder.WriteString(fmt.Sprintf("id=%v", ta.ID))
	builder.WriteString(", continent=")
	builder.WriteString(ta.Continent)
	builder.WriteString(", country=")
	builder.WriteString(ta.Country)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ta.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ta.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ta.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// TargetAreas is a parsable slice of TargetArea.
type TargetAreas []*TargetArea

func (ta TargetAreas) config(cfg config) {
	for _i := range ta {
		ta[_i].config = cfg
	}
}
