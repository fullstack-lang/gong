package diagrammer

type ModelNode interface {
	GenerateProgeny() []ModelNode
	GetChildren() []ModelNode
	// GetParent returns the node above
	GetParent() ModelNode

	GetName() string
	IsNameEditable() bool

	IsExpanded() bool
	SetIsExpanded(isExpanded bool)

	HasSecondCheckbox() bool
	HasSecondName() bool
	GetSecondName() string
}
