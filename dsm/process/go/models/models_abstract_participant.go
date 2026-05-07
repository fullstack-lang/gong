package models

type Participant struct {
	Name string

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
