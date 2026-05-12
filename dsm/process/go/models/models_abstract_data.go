package models

type Data struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*Data)(nil)
