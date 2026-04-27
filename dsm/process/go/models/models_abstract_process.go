package models

type Process struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	parentProcess *Process

	DiagramProcesss                   []*DiagramProcess
	DiagramProcessWhoseNodeIsExpanded []*DiagramProcess

	IsSubProcessNodeExpanded bool
	SubProcesses             []*Process

	Participants                   []*Participant
	ParticipantWhoseNodeIsExpanded []*Participant
}

var (
	_ AbstractType = (*Process)(nil)
)

