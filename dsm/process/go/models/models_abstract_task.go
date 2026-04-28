package models

type Task struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*Task)(nil)
