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

type Participant struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields
}

var (
	_ AbstractType = (*Participant)(nil)
)