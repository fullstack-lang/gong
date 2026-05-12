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

	// A task is an instance of a process within another process.
	// This is an instance of the whole/part meta model pattern.
	// by default, the task name is the process type name, but it can be overridden by the user.
	Type                     *Process
	IsTaskNameNotProcessName bool

	owningParticipant *Participant
}

var _ AbstractType = (*Task)(nil)

func (s *Task) GetOwningParticipant() *Participant {
	return s.owningParticipant
}
