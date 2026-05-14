package models

type Process struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	// SVG_Path is a a SVG Path rendering an icon representing the process.
	//gong:width 600 gong:height 300
	SVG_Path string

	// $960$ is a "magic number" because it is perfectly divisible by $24, 40, 48,$ and $60$
	InverseAppliedScaling float64

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
