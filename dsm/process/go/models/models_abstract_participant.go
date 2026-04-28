package models

type Participant struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Tasks []*Task
}

var _ AbstractType = (*Participant)(nil)
