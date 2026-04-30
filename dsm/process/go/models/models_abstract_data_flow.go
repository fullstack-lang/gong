package models

type DataFlow struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Start *Task

	End *Task
}
