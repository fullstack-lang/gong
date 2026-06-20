package models

// an AllocatedResourceShape is a shape that represents an allocated resource to a participant.
// It is used to display the allocated resource in the diagram.
type AllocatedResourceShape struct {
	Name string

	Participant *Participant
	Resource    *Resource
}
