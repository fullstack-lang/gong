package models

type Participant struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsTasksNodeExpanded bool
	Tasks               []*Task

	IsControlFlowsNodeExpanded bool
	ControlFlows               []*ControlFlow
}

var _ AbstractType = (*Participant)(nil)
