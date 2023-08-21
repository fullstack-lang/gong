package models

// TableExtraPathEnum - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum TableExtraPathEnum
type TableExtraPathEnum string

// values for TableExtraPathEnum
const (
	TableSelectExtra TableExtraPathEnum = "-table"
	TableSortExtra   TableExtraPathEnum = "-table-sort"
)
