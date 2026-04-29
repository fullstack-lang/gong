package models

type Task struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsStartTask bool
	IsEndTask   bool

	outcontrolFlows []*ControlFlow
	incontrolFlows  []*ControlFlow

	owningParticipant *Participant
}

var _ AbstractType = (*Task)(nil)
