package models

type ParticipantAbstractFields struct {
	// When the participant has one or more process/resource
	// it might be conveniant to consider it, semanticaly/visualy, as
	// the process/resource itself.
	// IsNameOverridenWithFirstAllocatedProcessResource is true in this case
	IsProcessResource bool

	//gong:text width:300 height:300
	Description string

	// SVG is a raw <svg>...</svg> string rendering an icon representing the participant.
	//gong:text gong:width 600 gong:height 300
	SVG string

	// Resources is the list of resources used by the participant
	Resources               []*Resource
	IsResourcesNodeExpanded bool

	LibraryAbstractFields
	AbstractTypeFields

	owningProcess *Process
}

func (s *ParticipantAbstractFields) GetOwningProcess() *Process {
	return s.owningProcess
}

type Participant struct {
	Name string

	ParticipantAbstractFields

	IsTasksNodeExpanded bool
	Tasks               []*Task

	IsControlFlowsNodeExpanded bool
	ControlFlows               []*ControlFlow

	TaskWhoseOutControlFlowsNodeIsExpanded []*Task
	TaskWhoseInControlFlowsNodeIsExpanded  []*Task

	IsDataFlowsNodeExpanded bool

	TaskWhoseOutDataFlowsNodeIsExpanded []*Task
	TaskWhoseInDataFlowsNodeIsExpanded  []*Task
}

var _ AbstractType = (*Participant)(nil)
