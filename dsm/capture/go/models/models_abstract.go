package models

type Library struct {
	Name string

	IsRootLibrary bool

	LibraryAbstractFields
	AbstractTypeFields

	RootDeliverables []*Deliverable
	RootConcerns     []*Concern
	RootStakeholders []*Stakeholder
	RootRequirements []*Requirement
	RootConcepts     []*Concept
	AnalysisNeeds    []*AnalysisNeed

	Notes []*Note

	Diagrams []*Diagram

	objects []AbstractType

	SubLibraries []*Library

	NbPixPerCharacter float64 // stored at the root Library only
}

// Note brings information to a diagram
type Note struct {
	//gong:text width:300 height:300
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Deliverables  []*Deliverable
	Tasks     []*Concern
	Resources []*Stakeholder
}

type Priority string

const (
	Low     Priority = "Low"
	Medium  Priority = "Medium"
	High    Priority = "High"
	Unknown Priority = "Unknown"
)

type Concern struct {
	Name string

	IDAirbus string

	Priority Priority

	LibraryAbstractFields
	AbstractTypeFields

	//gong:text width:300 height:300
	Description string

	SubConcerns []*Concern

	Inputs               []*Deliverable
	IsInputsNodeExpanded bool

	Outputs               []*Deliverable
	IsOutputsNodeExpanded bool

	// parentConcern is a computed field
	// since a Task belongs to at most one WBS,
	// a parentConcern is computed at each UX look. It can be null if the
	// task is a root task or an orphaned task
	parentConcern *Concern

	// Completion Management
	IsWithCompletion bool
	Completion       CompletionEnum

	Requirements []*Requirement
}

// GetIDAirbus implements [withIDAirbus].
func (concern *Concern) GetIDAirbus() string {
	return concern.IDAirbus
}

// CompletionEnum
type CompletionEnum string

const (
	PERCENT_100 CompletionEnum = "100 %"
	PERCENT_075 CompletionEnum = "75 %"
	PERCENT_050 CompletionEnum = "50 %"
	PERCENT_025 CompletionEnum = "25 %"
	PERCENT_000 CompletionEnum = "0 %"
)

// GONGDOC(NoteSemantic): Note on the models semantic
//
// [models.Deliverable] and [models.Task] are in Deliverable Breakdown Structure (PBS)
// and Work Breakdown Structure (WBS)
// PBS/WBS have 2 invariants that are enforced at each UX loop:
// - They are Trees
// - A [models.Deliverable]/[models.Task] belongs to at most one PBS/WBS.
// Those invariants allow prefix and parent to be computed at each UX loop
const NoteSemantic = ""

type Deliverable struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	//gong:text width:300 height:300
	Description string

	SubDeliverables []*Deliverable

	// producers are computed from [models.Task.Outputs]
	// this is a computed field, therefore, not exported
	producers               []*Concern
	IsProducersNodeExpanded bool

	// consumers are computed from [models.Task.Inputs]
	// this is a computed field, therefore, not exported
	consumers               []*Concern
	IsConsumersNodeExpanded bool

	// parentDeliverable is a computed field
	// since a Deliverable belongs to at most one WBS,
	// a parentDeliverable is computed at each UX look. It can be null if the
	// deliverable is a root deliverable or an orphaned deliverable
	parentDeliverable *Deliverable

	Concepts []*Concept
}

type Stakeholder struct {
	Name string

	IDAirbus string

	LibraryAbstractFields
	AbstractTypeFields

	//gong:text width:300 height:300
	Description string

	Concerns []*Concern

	SubStakeholders []*Stakeholder

	// parentStakeholder is a computed field
	// since a Resource belongs to at most one RBS,
	// a parentStakeholder is computed at each UX look. It can be null if the
	// resource is a root resource or an orphaned resource
	parentStakeholder *Stakeholder
}

type Requirement struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	SupportLevels []*SupportLevel
	Concepts      []*Concept
}

type SupportLevel struct {
	Name string
	Tool *Tool
}

type Concept struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Tools []*Tool
}

type Tool struct {
	Name string
}

// GetIDAirbus implements [withIDAirbus].
func (stakeholder *Stakeholder) GetIDAirbus() string {
	return stakeholder.IDAirbus
}

type AnalysisNeed struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields
}

var (
	_ AbstractType = (*Deliverable)(nil)
	_ AbstractType = (*Library)(nil)
	_ AbstractType = (*Concern)(nil)
	_ AbstractType = (*Note)(nil)
	_ AbstractType = (*Stakeholder)(nil)
	_ AbstractType = (*Requirement)(nil)
	_ AbstractType = (*Concept)(nil)
	_ AbstractType = (*AnalysisNeed)(nil)
)

type withIDAirbus interface {
	GetIDAirbus() string
}

var (
	_ withIDAirbus = (*Concern)(nil)
	_ withIDAirbus = (*Stakeholder)(nil)
)
