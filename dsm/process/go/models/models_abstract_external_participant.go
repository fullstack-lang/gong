package models

type ExternalParticipant struct {
	Name string

	ParticipantAbstractFields

	// Processes is the list of processes performed by the external participant
	Processes               []*Process
	IsProcessesNodeExpanded bool

	inDataFlows  []*DataFlow
	outDataFlows []*DataFlow
}

var _ AbstractType = (*ExternalParticipant)(nil)
