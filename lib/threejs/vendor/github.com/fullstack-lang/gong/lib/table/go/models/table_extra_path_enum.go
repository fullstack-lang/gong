package models

// TableExtraPathEnum - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum TableExtraPathEnum
type TableExtraPathEnum string

// values for TableExtraPathEnum
const (

	// those five default values are used by the "data" when used in combination for accessing
	// data
	StackNamePostFixForTableForMainTable          TableExtraPathEnum = "-table"
	StackNamePostFixForTableForMainForm           TableExtraPathEnum = "-form"
	StackNamePostFixForTableForMainTree           TableExtraPathEnum = "-sidebar"
	StackNamePostFixForTableForAssociation        TableExtraPathEnum = "-table-pick"
	StackNamePostFixForTableForAssociationSorting TableExtraPathEnum = "-table-sort"
)
