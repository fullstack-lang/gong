package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

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



type Diagram struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsChecked   bool
	IsEditable_ bool

	ShowPrefix bool // display shapes with their prefix

	DefaultBoxWidth  float64
	DefaultBoxHeigth float64

	Width  float64
	Height float64

	ConcernsWhoseRequirementsNodeIsExpanded []*Concern

	IsRequirementsNodeExpanded bool
	IsConceptsNodeExpanded     bool

	Product_Shapes              []*ProductShape
	map_Product_ProductShape    map[*Deliverable]*ProductShape
	ProductsWhoseNodeIsExpanded []*Deliverable // to be made private once in production (no need to persist)
	IsPBSNodeExpanded           bool

	ProductComposition_Shapes           []*ProductCompositionShape
	map_Product_ProductCompositionShape map[*Deliverable]*ProductCompositionShape

	IsConcernsNodeExpanded bool

	Concern_Shapes                         []*ConcernShape
	map_Concern_ConcernShape               map[*Concern]*ConcernShape
	ConcernsWhoseNodeIsExpanded            []*Concern // to be made private once in production (no need to persist)ExpandableNodeObject
	ConcernsWhoseInputNodeIsExpanded       []*Concern
	ConcernsWhoseStakeholderNodeIsExpanded []*Concern
	ConcernssWhoseOutputNodeIsExpanded     []*Concern

	ConcernComposition_Shapes           []*ConcernCompositionShape
	map_Concern_ConcernCompositionShape map[*Concern]*ConcernCompositionShape

	ConcernInputShapes         []*ConcernInputShape
	map_Concern_TaskInputShape map[concernProductKey]*ConcernInputShape

	ConcernOutputShapes            []*ConcernOutputShape
	map_Concern_ConcernOutputShape map[concernProductKey]*ConcernOutputShape

	Note_Shapes              []*NoteShape
	map_Note_NoteShape       map[*Note]*NoteShape
	NotesWhoseNodeIsExpanded []*Note
	IsNotesNodeExpanded      bool

	NoteProductShapes         []*NoteProductShape
	map_Note_NoteProductShape map[noteProductKey]*NoteProductShape

	NoteTaskShapes         []*NoteTaskShape
	map_Note_NoteTaskShape map[noteTaskKey]*NoteTaskShape

	NoteResourceShapes         []*NoteStakeholderShape
	map_Note_NoteResourceShape map[noteResourceKey]*NoteStakeholderShape

	Stakeholder_Shapes               []*StakeholderShape
	map_Stakeholder_StakeholderShape map[*Stakeholder]*StakeholderShape
	ResourcesWhoseNodeIsExpanded     []*Stakeholder
	IsStakeholdersNodeExpanded       bool

	ResourceComposition_Shapes            []*StakeholderCompositionShape
	map_Resource_ResourceCompositionShape map[*Stakeholder]*StakeholderCompositionShape

	StakeholderConcernShapes                          []*StakeholderConcernShape
	map_StakeholderConcernKey_StakeholderConcernShape map[stakeholderConcernKey]*StakeholderConcernShape

	Requirement_Shapes               []*RequirementShape
	map_Requirement_RequirementShape map[*Requirement]*RequirementShape
	RequirementsWhoseNodeIsExpanded  []*Requirement

	Concept_Shapes              []*ConceptShape
	map_Concept_ConceptShape    map[*Concept]*ConceptShape
	ConceptsWhoseNodeIsExpanded []*Concept

	map_Product_Rect     map[*Deliverable]*svg.Rect
	map_Task_Rect        map[*Concern]*svg.Rect
	map_Note_Rect        map[*Note]*svg.Rect
	map_Stakeholder_Rect map[*Stakeholder]*svg.Rect
	map_Requirement_Rect map[*Requirement]*svg.Rect
	map_Concept_Rect     map[*Concept]*svg.Rect

	map_SvgRect_ProductShape     map[*svg.Rect]*ProductShape
	map_SvgRect_ConcernShape     map[*svg.Rect]*ConcernShape
	map_SvgRect_NoteShape        map[*svg.Rect]*NoteShape
	map_SvgRect_StakeholderShape map[*svg.Rect]*StakeholderShape
	map_SvgRect_RequirementShape map[*svg.Rect]*RequirementShape
	map_SvgRect_ConceptShape     map[*svg.Rect]*ConceptShape

	elementWhoseDiagramListIsDisplayed AbstractType
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}

func (d *Diagram) GetIsChecked() bool { return d.IsChecked }
func (d *Diagram) SetIsChecked(b bool) { d.IsChecked = b }
func (d *Diagram) GetIsShowPrefix() bool { return d.ShowPrefix }
func (d *Diagram) SetIsShowPrefix(b bool) { d.ShowPrefix = b }
func (d *Diagram) GetDefaultBoxWidth() float64 { return d.DefaultBoxWidth }
func (d *Diagram) GetDefaultBoxHeigth() float64 { return d.DefaultBoxHeigth }
func (d *Diagram) GetDiagramListElement() AbstractType { return d.elementWhoseDiagramListIsDisplayed }
func (d *Diagram) SetDiagramListElement(a AbstractType) { d.elementWhoseDiagramListIsDisplayed = a }

// Note brings information to a diagram
type Note struct {
	//gong:text width:300 height:300
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Products  []*Deliverable
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
// [models.Product] and [models.Task] are in Product Breakdown Structure (PBS)
// and Work Breakdown Structure (WBS)
// PBS/WBS have 2 invariants that are enforced at each UX loop:
// - They are Trees
// - A [models.Product]/[models.Task] belongs to at most one PBS/WBS.
// Those invariants allow prefix and parent to be computed at each UX loop
const NoteSemantic = ""

type Deliverable struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	//gong:text width:300 height:300
	Description string

	SubProducts []*Deliverable

	// producers are computed from [models.Task.Outputs]
	// this is a computed field, therefore, not exported
	producers               []*Concern
	IsProducersNodeExpanded bool

	// consumers are computed from [models.Task.Inputs]
	// this is a computed field, therefore, not exported
	consumers               []*Concern
	IsConsumersNodeExpanded bool

	// parentProduct is a computed field
	// since a Product belongs to at most one WBS,
	// a parentProduct is computed at each UX look. It can be null if the
	// product is a root product or an orphaned product
	parentProduct *Deliverable

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
	Name          string

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
	Name  string

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
