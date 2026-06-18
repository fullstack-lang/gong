package models

import svg "github.com/fullstack-lang/gong/lib/svg/go/models"

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

	Product_Shapes                  []*ProductShape
	map_Product_ProductShape        map[*Deliverable]*ProductShape
	ProductsWhoseNodeIsExpanded     []*Deliverable // to be made private once in production (no need to persist)
	ProductsWhoseConceptsNodeIsExpanded []*Deliverable
	IsPBSNodeExpanded               bool

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

	Concept_Shapes                          []*ConceptShape
	map_Concept_ConceptShape                map[*Concept]*ConceptShape
	ConceptsWhoseNodeIsExpanded             []*Concept
	ConceptsWhoseDeliverablesNodeIsExpanded []*Concept

	DeliverableConceptShapes                          []*DeliverableConceptShape
	map_DeliverableConceptKey_DeliverableConceptShape map[deliverableConceptKey]*DeliverableConceptShape

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

func (d *Diagram) GetIsChecked() bool                   { return d.IsChecked }
func (d *Diagram) SetIsChecked(b bool)                  { d.IsChecked = b }
func (d *Diagram) GetIsShowPrefix() bool                { return d.ShowPrefix }
func (d *Diagram) SetIsShowPrefix(b bool)               { d.ShowPrefix = b }
func (d *Diagram) GetDefaultBoxWidth() float64          { return d.DefaultBoxWidth }
func (d *Diagram) GetDefaultBoxHeigth() float64         { return d.DefaultBoxHeigth }
func (d *Diagram) GetDiagramListElement() AbstractType  { return d.elementWhoseDiagramListIsDisplayed }
func (d *Diagram) SetDiagramListElement(a AbstractType) { d.elementWhoseDiagramListIsDisplayed = a }
