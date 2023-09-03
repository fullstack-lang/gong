package models

// TableExtraPathEnum - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum TableExtraPathEnum
type TableExtraPathEnum string

// values for TableExtraPathEnum
const (
	StackNamePostFixForTableForAssociation        TableExtraPathEnum = "-table"
	StackNamePostFixForTableForAssociationSorting TableExtraPathEnum = "-table-sort"
)
