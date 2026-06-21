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
	RootProcesses               []*Process
	IsProcessesNodeExpanded     bool
	ProcesssWhoseNodeIsExpanded []*Process

	RootDataFlows                []*DataFlow
	IsDataFlowsNodeExpanded      bool
	DataFlowsWhoseNodeIsExpanded []*DataFlow

	RootDatas                []*Data
	IsDatasNodeExpanded      bool
	DatasWhoseNodeIsExpanded []*Data

	RootResources                   []*Resource
	IsResourcesNodeExpanded         bool
	ResourcesWhoseNodeIsExpanded    []*Resource
	ParticipantsWhoseNodeIsExpanded []*Participant

	RootNotes                []*Note
	IsNotesNodeExpanded      bool
	NotesWhoseNodeIsExpanded []*Note

	// temporary persistance of the library expand status.
	IsExpandedTmp bool
}
