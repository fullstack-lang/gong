package models

type Task struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsStartTask bool
	IsEndTask   bool

	outControlFlows []*ControlFlow
	inControlFlows  []*ControlFlow

	outDataFlows []*DataFlow
	inDataFlows  []*DataFlow

	owningParticipant *Participant
}

var _ AbstractType = (*Task)(nil)
