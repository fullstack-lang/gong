package models

type ControlFlow struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	Start *Port

	End *Port
}

var _ AbstractType = (*ControlFlow)(nil)
