package models

type Participant struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsTasksNodeExpanded bool
	Tasks               []*Task

	ControlFlows []*ControlFlow
}

var _ AbstractType = (*Participant)(nil)
