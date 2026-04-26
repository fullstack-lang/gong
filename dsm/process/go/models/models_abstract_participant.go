package models

type Participant struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*Participant)(nil)
