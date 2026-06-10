package models

type Library struct {

	// DSM mandatory
	Name string

	SubLibraries []*Library

	NbPixPerCharacter float64 // stored at the root Library only

	//gong:width 600 gong:height 300
	LogoSVGFile string // the content of the logo file, used for the static site generation

	LibraryAbstractFields
	AbstractTypeFields

	// There is one and only one root library per stage.
	IsRootLibrary bool

	// DSM specific
	RootProducts   []*Product
	RootTasks      []*Task
	RootTaskGroups []*TaskGroup
	RootResources  []*Resource

	Notes []*Note

	Diagrams []*Diagram

	objects []AbstractType
}
