package models

// Repository is a first rank element
type Repository struct {
	Name         string
	ParameterUse []*ParameterShape
	GroupUse     []*GroupUse
	LibraryAbstractFields
	AbstractTypeFields
}
