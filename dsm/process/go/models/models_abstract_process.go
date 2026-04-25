package models

type Process struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	parentProcess *Process

	DiagramProcesss []*DiagramProcess

	IsSubProcessNodeExpanded bool
	SubProcesses             []*Process

	IsParticipantsNodeExpanded     bool
	Participants                   []*Participant
	ParticipantWhoseNodeIsExpanded []*Participant
}

var (
	_ AbstractType = (*Process)(nil)
)

type Participant struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields
}

var (
	_ AbstractType = (*Participant)(nil)
)
