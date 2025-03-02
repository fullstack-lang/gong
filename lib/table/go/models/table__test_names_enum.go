package models

// TableTestNameEnum -
// Authorized name for form group
type TableTestNameEnum string

// values for TableTestNameEnum
const (
	ManualyEditedTableStackName TableTestNameEnum = "manualy edited table"
	ManualyEditedFormStackName  TableTestNameEnum = "manualy edited form"
	GeneratedTableStackName     TableTestNameEnum = "generated table"
)
