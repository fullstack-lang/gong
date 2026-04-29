package models

type Participant struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Tasks []*Task

	ControlFlows []*ControlFlow
}

var _ AbstractType = (*Participant)(nil)
