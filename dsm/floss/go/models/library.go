package models

type Library struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	// There is one and only one root library per stage.
	IsRootLibrary bool

	SubLibraries []*Library

	// DSM specifc
	objects []AbstractType

	IsSubLibrariesNodeExpanded      bool
	SubLibrariesWhoseNodeIsExpanded []*Library

	NbPixPerCharacter float64 // stored at the root Library only

	//gong:width 600 gong:height 300
	LogoSVGFile string // the content of the logo file, used for the static site generation

	// DSM specific fields
	RootSystemes               []*System
	IsSystemesNodeExpanded     bool
	SystemsWhoseNodeIsExpanded []*System

	RootComplexitys                []*Complexity
	IsComplexitysNodeExpanded      bool
	ComplexitysWhoseNodeIsExpanded []*Complexity

	RootPerformances                []*Performance
	IsPerformancesNodeExpanded      bool
	PerformancesWhoseNodeIsExpanded []*Performance

	RootEfforts                []*Effort
	IsEffortsNodeExpanded      bool
	EffortsWhoseNodeIsExpanded []*Effort

	// temporary persistance of the library expand status.
	IsExpandedTmp bool
}

