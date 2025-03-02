package models

// StacksNames - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum StacksNames
type StacksNames string

// values for TableExtraNameEnum
const (
	SvgStackName     StacksNames = "svg"     // the gongdoc application launch a svg stack for visualisation. this is its name
	GongdocStackName StacksNames = "gongdoc" // the default gongdoc stack has this name
)
