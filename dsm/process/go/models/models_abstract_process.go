package models

type Process struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	parentProcess *Process
	SubProcesses  []*Process

	DiagramProcesss []*DiagramProcess

	IsSubProcessNodeExpanded bool
}

var (
	_ AbstractType = (*Process)(nil)
)
