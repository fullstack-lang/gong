package models

type Data struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*Data)(nil)
