package diagrammer

type ModelElementNode interface {
	ModelNode

	GetElement() any

	// CanBeAddedToDiagram returns true if the model element can be displayed
	CanBeAddedToDiagram() bool
	AddToDiagram()
	RemoveFromDiagram()

	// GetLinkedNode returns a node to navigate to
	GetLinkedNode() ModelElementNode
}
