package models

type Note struct {
	//gong:text width:300 height:300
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Tasks []*Task
}

var _ AbstractType = (*Note)(nil)
