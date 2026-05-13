package models

type AllocatedProcessShape struct {
	Name string

	Participant *Participant
	Process     *Process
}

