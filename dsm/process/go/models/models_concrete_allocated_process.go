package models

type AllocatedProcessShape struct {
	Name string

	ExternalParticipant *ExternalParticipant
	Process             *Process
}
