package models

type ControlFlow struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Start *Task

	End *Task
}
