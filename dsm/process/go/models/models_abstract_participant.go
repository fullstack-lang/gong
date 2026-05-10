package models

type Participant struct {
	Name string

	Resources               []*Resource
	IsResourcesNodeExpanded bool

	LibraryAbstractFields
	AbstractTypeFields

	IsTasksNodeExpanded bool
	Tasks               []*Task

	IsControlFlowsNodeExpanded bool
	ControlFlows               []*ControlFlow

	TaskWhoseOutControlFlowsNodeIsExpanded []*Task
	TaskWhoseInControlFlowsNodeIsExpanded  []*Task

	IsDataFlowsNodeExpanded bool
	inDataFlows             []*DataFlow
	outDataFlows            []*DataFlow

	TaskWhoseOutDataFlowsNodeIsExpanded []*Task
	TaskWhoseInDataFlowsNodeIsExpanded  []*Task

	owningProcess *Process
}

var _ AbstractType = (*Participant)(nil)

func (s *Participant) GetOwningProcess() *Process {
	return s.owningProcess
}
