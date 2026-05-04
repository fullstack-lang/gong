package models

type DataFlow struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Start *Task

	End *Task

	IsDatasNodeExpanded bool
	Datas               []*Data
}

var _ AbstractType = (*DataFlow)(nil)
