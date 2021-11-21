// Code generated by entc, DO NOT EDIT.

package goodextrainfo

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the goodextrainfo type in the database.
	Label = "good_extra_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldPosters holds the string denoting the posters field in the database.
	FieldPosters = "posters"
	// FieldLabels holds the string denoting the labels field in the database.
	FieldLabels = "labels"
	// FieldOutSale holds the string denoting the out_sale field in the database.
	FieldOutSale = "out_sale"
	// FieldPreSale holds the string denoting the pre_sale field in the database.
	FieldPreSale = "pre_sale"
	// FieldVoteCount holds the string denoting the vote_count field in the database.
	FieldVoteCount = "vote_count"
	// FieldRating holds the string denoting the rating field in the database.
	FieldRating = "rating"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the goodextrainfo in the database.
	Table = "good_extra_infos"
)

// Columns holds all SQL columns for goodextrainfo fields.
var Columns = []string{
	FieldID,
	FieldGoodID,
	FieldPosters,
	FieldLabels,
	FieldOutSale,
	FieldPreSale,
	FieldVoteCount,
	FieldRating,
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
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() int64
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() int64
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() int64
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
