package models

// StacksNames - enumeration of possible 'type' values for an HTML <input> element
// swagger:enum StacksNames
type StacksNames string

// values for TableExtraNameEnum
const (
	GongsvgStack    StacksNames = "GongsvgStack"
	ObjectTree      StacksNames = "Object tree"
	DiagramTree     StacksNames = "Diagram tree"
	ButtonStackName StacksNames = "Button Stack Name"

	GongtreeStackName  StacksNames = "gongtree"
	GongtableStackName StacksNames = "gongtable"
	GongsimStackName   StacksNames = "gongsim"
)
