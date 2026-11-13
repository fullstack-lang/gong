package models

type Process struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	// SVG is a raw <svg>...</svg> string rendering an icon representing the process.
	//gong:text gong:width 600 gong:height 300
	SVG string

	parentProcess *Process

	DiagramProcesss                   []*DiagramProcess
	DiagramProcessWhoseNodeIsExpanded []*DiagramProcess

	IsSubProcessNodeExpanded bool
	SubProcesses             []*Process

	Participants                   []*Participant
	ParticipantWhoseNodeIsExpanded []*Participant

	DataFlows               []*DataFlow
	IsDataFlowsNodeExpanded bool

	ExternalParticipants                   []*Participant
	ExternalParticipantWhoseNodeIsExpanded []*Participant
}

var _ AbstractType = (*Process)(nil)
