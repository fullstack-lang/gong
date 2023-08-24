package models

// TableExtraNameEnum - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum TableExtraNameEnum
type TableExtraNameEnum string

// values for TableExtraNameEnum
const (
	TableSelectExtraName TableExtraNameEnum = "tmp-picker"
	TableSortExtraName   TableExtraNameEnum = "tmp-sort"
)
