package diagrammer

type Model interface {
	GenerateProgeny() []ModelNode
	GetChildren() []ModelNode
}
