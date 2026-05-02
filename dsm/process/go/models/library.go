package models

type Library struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	objects []AbstractType

	SubLibraries                    []*Library
	IsSubLibrariesNodeExpanded      bool
	SubLibrariesWhoseNodeIsExpanded []*Library

	NbPixPerCharacter float64 // stored at the root Library only

	//gong:width 600 gong:height 300
	LogoSVGFile string // the content of the logo file, used for the static site generation

	// DSM specific fields
	RootProcesses               []*Process
	IsProcessesNodeExpanded     bool
	ProcesssWhoseNodeIsExpanded []*Process

	RootDataFlows                []*DataFlow
	IsDataFlowsNodeExpanded      bool
	DataFlowsWhoseNodeIsExpanded []*DataFlow

	// temporary persistance of the library expand status.
	IsExpandedTmp bool
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
