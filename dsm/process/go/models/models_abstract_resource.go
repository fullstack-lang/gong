package models

type Resource struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields
}

var _ AbstractType = (*Resource)(nil)

func (s *Resource) GetOwningLibrary() *Library {
	return s.owningLibrary
}

func (s *Resource) SetOwningLibrary(library *Library) {
	s.owningLibrary = library
}
