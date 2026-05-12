package models

type ControlFlow struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	Start *Task

	End *Task
}

var _ AbstractType = (*ControlFlow)(nil)
