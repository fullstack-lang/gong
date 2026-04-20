package models

type Library struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	RootProducts  []*Product
	RootTasks     []*Task
	RootResources []*Resource

	Notes []*Note

	Diagrams []*Diagram

	objects []AbstractType

	SubLibraries []*Library

	NbPixPerCharacter float64 // stored at the root Library only

	//gong:width 600 gong:height 300
	LogoSVGFile string // the content of the logo file, used for the static site generation
}

type LibraryAbstractFields struct {
	owningLibrary *Library
}

type LibraryOwnedType interface {
	GetOwningLibrary() *Library
	SetOwningLibrary(library *Library)
}

func (r *LibraryAbstractFields) GetOwningLibrary() *Library {
	return r.owningLibrary
}

func (r *LibraryAbstractFields) SetOwningLibrary(library *Library) {
	r.owningLibrary = library
}
