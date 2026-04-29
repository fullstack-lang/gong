package models

type Task struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsStartTask bool
	IsEndTask   bool
}

var _ AbstractType = (*Task)(nil)
