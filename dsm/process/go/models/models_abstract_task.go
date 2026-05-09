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

func (s *Task) GetOwningParticipant() *Participant {
	return s.owningParticipant
}
	