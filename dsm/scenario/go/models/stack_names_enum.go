package models

// StacksNames - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum StacksNames
type StacksNames string

// values for TableExtraNameEnum
const (
	// we have three application stacks. All have the same name
	Weberv2specsStackName StacksNames = "weberv2specs"

	GongsvgStackName   StacksNames = "gongsvg"
	GongtreeStackName  StacksNames = "gongtree"
	GongtableStackName StacksNames = "gongtable"
	GongsimStackName   StacksNames = "gongsim"
)
