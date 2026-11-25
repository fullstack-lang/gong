package models

type ControlFlow struct {
	Name string

	Datas []*Data

	//gong:text width:300 height:300
	Description string

	Guard string

	LibraryAbstractFields
	AbstractTypeFields

	Start *Task

	End *Task

	IsDatasNodeExpanded bool
}

var _ AbstractType = (*ControlFlow)(nil)
