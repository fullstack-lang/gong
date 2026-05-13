package models

type Participant struct {
	Name string

	// When the participant has one or more process/resource
	// it might be conveniant to consider it, semanticaly/visualy, as
	// the process/resource itself.
	// IsNameOverridenWithFirstAllocatedProcessResource is true in this case
	IsProcessResource bool

	//gong:text width:300 height:300
	Description string

	// Resources is the list of resources used by the participant
	Resources               []*Resource
	IsResourcesNodeExpanded bool

	// Processess is the list of processes performed by the participant
	Processes               []*Process
	IsProcessesNodeExpanded bool

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
