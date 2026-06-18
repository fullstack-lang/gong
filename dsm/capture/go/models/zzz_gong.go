// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	capture_go "github.com/fullstack-lang/gong/dsm/capture/go"
)

// can be used for
//
//	days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __Gong__Abs
	_ = strings.Clone("")
)

const (
	ProbeTreeSidebarSuffix           = ":sidebar of the probe"
	ProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	ProbeTableSuffix                 = ":table of the probe"
	ProbeNotificationTableSuffix     = ":notification table of the probe"
	ProbeFormSuffix                  = ":form of the probe"
	ProbeSplitSuffix                 = ":probe of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
}

// Stage enables storage of staged instances
type Stage struct {
	name string

	// isInDeltaMode is true when the stage is used to compute difference between
	// succesive commit
	isInDeltaMode bool

	// gongMarshallingMode set the marshalling mode
	gongMarshallingMode GongMarshallingMode
	// some stages have semantic rules that forbids them to be empty
	// like for git, the commit #0 (genesis commit) cannot be rolled back
	isWithGenesisCommit bool

	// insertion point for definition of arrays registering instances
	AnalysisNeeds                map[*AnalysisNeed]struct{}
	AnalysisNeeds_instance       map[*AnalysisNeed]*AnalysisNeed
	AnalysisNeeds_mapString      map[string]*AnalysisNeed
	AnalysisNeedOrder            uint
	AnalysisNeed_stagedOrder     map[*AnalysisNeed]uint
	AnalysisNeed_orderStaged     map[uint]*AnalysisNeed
	AnalysisNeeds_reference      map[*AnalysisNeed]*AnalysisNeed
	AnalysisNeeds_referenceOrder map[*AnalysisNeed]uint

	// insertion point for slice of pointers maps
	OnAfterAnalysisNeedCreateCallback OnAfterCreateInterface[AnalysisNeed]
	OnAfterAnalysisNeedUpdateCallback OnAfterUpdateInterface[AnalysisNeed]
	OnAfterAnalysisNeedDeleteCallback OnAfterDeleteInterface[AnalysisNeed]
	OnAfterAnalysisNeedReadCallback   OnAfterReadInterface[AnalysisNeed]

	Concepts                map[*Concept]struct{}
	Concepts_instance       map[*Concept]*Concept
	Concepts_mapString      map[string]*Concept
	ConceptOrder            uint
	Concept_stagedOrder     map[*Concept]uint
	Concept_orderStaged     map[uint]*Concept
	Concepts_reference      map[*Concept]*Concept
	Concepts_referenceOrder map[*Concept]uint

	// insertion point for slice of pointers maps
	Concept_Tools_reverseMap map[*Tool]*Concept

	OnAfterConceptCreateCallback OnAfterCreateInterface[Concept]
	OnAfterConceptUpdateCallback OnAfterUpdateInterface[Concept]
	OnAfterConceptDeleteCallback OnAfterDeleteInterface[Concept]
	OnAfterConceptReadCallback   OnAfterReadInterface[Concept]

	Concerns                map[*Concern]struct{}
	Concerns_instance       map[*Concern]*Concern
	Concerns_mapString      map[string]*Concern
	ConcernOrder            uint
	Concern_stagedOrder     map[*Concern]uint
	Concern_orderStaged     map[uint]*Concern
	Concerns_reference      map[*Concern]*Concern
	Concerns_referenceOrder map[*Concern]uint

	// insertion point for slice of pointers maps
	Concern_SubConcerns_reverseMap map[*Concern]*Concern

	Concern_Inputs_reverseMap map[*Deliverable]*Concern

	Concern_Outputs_reverseMap map[*Deliverable]*Concern

	Concern_Requirements_reverseMap map[*Requirement]*Concern

	OnAfterConcernCreateCallback OnAfterCreateInterface[Concern]
	OnAfterConcernUpdateCallback OnAfterUpdateInterface[Concern]
	OnAfterConcernDeleteCallback OnAfterDeleteInterface[Concern]
	OnAfterConcernReadCallback   OnAfterReadInterface[Concern]

	ConcernCompositionShapes                map[*ConcernCompositionShape]struct{}
	ConcernCompositionShapes_instance       map[*ConcernCompositionShape]*ConcernCompositionShape
	ConcernCompositionShapes_mapString      map[string]*ConcernCompositionShape
	ConcernCompositionShapeOrder            uint
	ConcernCompositionShape_stagedOrder     map[*ConcernCompositionShape]uint
	ConcernCompositionShape_orderStaged     map[uint]*ConcernCompositionShape
	ConcernCompositionShapes_reference      map[*ConcernCompositionShape]*ConcernCompositionShape
	ConcernCompositionShapes_referenceOrder map[*ConcernCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterConcernCompositionShapeCreateCallback OnAfterCreateInterface[ConcernCompositionShape]
	OnAfterConcernCompositionShapeUpdateCallback OnAfterUpdateInterface[ConcernCompositionShape]
	OnAfterConcernCompositionShapeDeleteCallback OnAfterDeleteInterface[ConcernCompositionShape]
	OnAfterConcernCompositionShapeReadCallback   OnAfterReadInterface[ConcernCompositionShape]

	ConcernInputShapes                map[*ConcernInputShape]struct{}
	ConcernInputShapes_instance       map[*ConcernInputShape]*ConcernInputShape
	ConcernInputShapes_mapString      map[string]*ConcernInputShape
	ConcernInputShapeOrder            uint
	ConcernInputShape_stagedOrder     map[*ConcernInputShape]uint
	ConcernInputShape_orderStaged     map[uint]*ConcernInputShape
	ConcernInputShapes_reference      map[*ConcernInputShape]*ConcernInputShape
	ConcernInputShapes_referenceOrder map[*ConcernInputShape]uint

	// insertion point for slice of pointers maps
	OnAfterConcernInputShapeCreateCallback OnAfterCreateInterface[ConcernInputShape]
	OnAfterConcernInputShapeUpdateCallback OnAfterUpdateInterface[ConcernInputShape]
	OnAfterConcernInputShapeDeleteCallback OnAfterDeleteInterface[ConcernInputShape]
	OnAfterConcernInputShapeReadCallback   OnAfterReadInterface[ConcernInputShape]

	ConcernOutputShapes                map[*ConcernOutputShape]struct{}
	ConcernOutputShapes_instance       map[*ConcernOutputShape]*ConcernOutputShape
	ConcernOutputShapes_mapString      map[string]*ConcernOutputShape
	ConcernOutputShapeOrder            uint
	ConcernOutputShape_stagedOrder     map[*ConcernOutputShape]uint
	ConcernOutputShape_orderStaged     map[uint]*ConcernOutputShape
	ConcernOutputShapes_reference      map[*ConcernOutputShape]*ConcernOutputShape
	ConcernOutputShapes_referenceOrder map[*ConcernOutputShape]uint

	// insertion point for slice of pointers maps
	OnAfterConcernOutputShapeCreateCallback OnAfterCreateInterface[ConcernOutputShape]
	OnAfterConcernOutputShapeUpdateCallback OnAfterUpdateInterface[ConcernOutputShape]
	OnAfterConcernOutputShapeDeleteCallback OnAfterDeleteInterface[ConcernOutputShape]
	OnAfterConcernOutputShapeReadCallback   OnAfterReadInterface[ConcernOutputShape]

	ConcernShapes                map[*ConcernShape]struct{}
	ConcernShapes_instance       map[*ConcernShape]*ConcernShape
	ConcernShapes_mapString      map[string]*ConcernShape
	ConcernShapeOrder            uint
	ConcernShape_stagedOrder     map[*ConcernShape]uint
	ConcernShape_orderStaged     map[uint]*ConcernShape
	ConcernShapes_reference      map[*ConcernShape]*ConcernShape
	ConcernShapes_referenceOrder map[*ConcernShape]uint

	// insertion point for slice of pointers maps
	OnAfterConcernShapeCreateCallback OnAfterCreateInterface[ConcernShape]
	OnAfterConcernShapeUpdateCallback OnAfterUpdateInterface[ConcernShape]
	OnAfterConcernShapeDeleteCallback OnAfterDeleteInterface[ConcernShape]
	OnAfterConcernShapeReadCallback   OnAfterReadInterface[ConcernShape]

	Deliverables                map[*Deliverable]struct{}
	Deliverables_instance       map[*Deliverable]*Deliverable
	Deliverables_mapString      map[string]*Deliverable
	DeliverableOrder            uint
	Deliverable_stagedOrder     map[*Deliverable]uint
	Deliverable_orderStaged     map[uint]*Deliverable
	Deliverables_reference      map[*Deliverable]*Deliverable
	Deliverables_referenceOrder map[*Deliverable]uint

	// insertion point for slice of pointers maps
	Deliverable_SubProducts_reverseMap map[*Deliverable]*Deliverable

	Deliverable_Concepts_reverseMap map[*Concept]*Deliverable

	OnAfterDeliverableCreateCallback OnAfterCreateInterface[Deliverable]
	OnAfterDeliverableUpdateCallback OnAfterUpdateInterface[Deliverable]
	OnAfterDeliverableDeleteCallback OnAfterDeleteInterface[Deliverable]
	OnAfterDeliverableReadCallback   OnAfterReadInterface[Deliverable]

	Diagrams                map[*Diagram]struct{}
	Diagrams_instance       map[*Diagram]*Diagram
	Diagrams_mapString      map[string]*Diagram
	DiagramOrder            uint
	Diagram_stagedOrder     map[*Diagram]uint
	Diagram_orderStaged     map[uint]*Diagram
	Diagrams_reference      map[*Diagram]*Diagram
	Diagrams_referenceOrder map[*Diagram]uint

	// insertion point for slice of pointers maps
	Diagram_ConcernsWhoseRequirementsNodeIsExpanded_reverseMap map[*Concern]*Diagram

	Diagram_Product_Shapes_reverseMap map[*ProductShape]*Diagram

	Diagram_ProductsWhoseNodeIsExpanded_reverseMap map[*Deliverable]*Diagram

	Diagram_ProductComposition_Shapes_reverseMap map[*ProductCompositionShape]*Diagram

	Diagram_Concern_Shapes_reverseMap map[*ConcernShape]*Diagram

	Diagram_ConcernsWhoseNodeIsExpanded_reverseMap map[*Concern]*Diagram

	Diagram_ConcernsWhoseInputNodeIsExpanded_reverseMap map[*Concern]*Diagram

	Diagram_ConcernsWhoseStakeholderNodeIsExpanded_reverseMap map[*Concern]*Diagram

	Diagram_ConcernssWhoseOutputNodeIsExpanded_reverseMap map[*Concern]*Diagram

	Diagram_ConcernComposition_Shapes_reverseMap map[*ConcernCompositionShape]*Diagram

	Diagram_ConcernInputShapes_reverseMap map[*ConcernInputShape]*Diagram

	Diagram_ConcernOutputShapes_reverseMap map[*ConcernOutputShape]*Diagram

	Diagram_Note_Shapes_reverseMap map[*NoteShape]*Diagram

	Diagram_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*Diagram

	Diagram_NoteProductShapes_reverseMap map[*NoteProductShape]*Diagram

	Diagram_NoteTaskShapes_reverseMap map[*NoteTaskShape]*Diagram

	Diagram_NoteResourceShapes_reverseMap map[*NoteStakeholderShape]*Diagram

	Diagram_Stakeholder_Shapes_reverseMap map[*StakeholderShape]*Diagram

	Diagram_ResourcesWhoseNodeIsExpanded_reverseMap map[*Stakeholder]*Diagram

	Diagram_ResourceComposition_Shapes_reverseMap map[*StakeholderCompositionShape]*Diagram

	Diagram_StakeholderConcernShapes_reverseMap map[*StakeholderConcernShape]*Diagram

	OnAfterDiagramCreateCallback OnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback OnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback OnAfterDeleteInterface[Diagram]
	OnAfterDiagramReadCallback   OnAfterReadInterface[Diagram]

	Librarys                map[*Library]struct{}
	Librarys_instance       map[*Library]*Library
	Librarys_mapString      map[string]*Library
	LibraryOrder            uint
	Library_stagedOrder     map[*Library]uint
	Library_orderStaged     map[uint]*Library
	Librarys_reference      map[*Library]*Library
	Librarys_referenceOrder map[*Library]uint

	// insertion point for slice of pointers maps
	Library_RootDeliverables_reverseMap map[*Deliverable]*Library

	Library_RootConcerns_reverseMap map[*Concern]*Library

	Library_RootStakeholders_reverseMap map[*Stakeholder]*Library

	Library_RootRequirements_reverseMap map[*Requirement]*Library

	Library_RootConcepts_reverseMap map[*Concept]*Library

	Library_AnalysisNeeds_reverseMap map[*AnalysisNeed]*Library

	Library_Notes_reverseMap map[*Note]*Library

	Library_Diagrams_reverseMap map[*Diagram]*Library

	Library_SubLibraries_reverseMap map[*Library]*Library

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	Notes                map[*Note]struct{}
	Notes_instance       map[*Note]*Note
	Notes_mapString      map[string]*Note
	NoteOrder            uint
	Note_stagedOrder     map[*Note]uint
	Note_orderStaged     map[uint]*Note
	Notes_reference      map[*Note]*Note
	Notes_referenceOrder map[*Note]uint

	// insertion point for slice of pointers maps
	Note_Products_reverseMap map[*Deliverable]*Note

	Note_Tasks_reverseMap map[*Concern]*Note

	Note_Resources_reverseMap map[*Stakeholder]*Note

	OnAfterNoteCreateCallback OnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback OnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback OnAfterDeleteInterface[Note]
	OnAfterNoteReadCallback   OnAfterReadInterface[Note]

	NoteProductShapes                map[*NoteProductShape]struct{}
	NoteProductShapes_instance       map[*NoteProductShape]*NoteProductShape
	NoteProductShapes_mapString      map[string]*NoteProductShape
	NoteProductShapeOrder            uint
	NoteProductShape_stagedOrder     map[*NoteProductShape]uint
	NoteProductShape_orderStaged     map[uint]*NoteProductShape
	NoteProductShapes_reference      map[*NoteProductShape]*NoteProductShape
	NoteProductShapes_referenceOrder map[*NoteProductShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteProductShapeCreateCallback OnAfterCreateInterface[NoteProductShape]
	OnAfterNoteProductShapeUpdateCallback OnAfterUpdateInterface[NoteProductShape]
	OnAfterNoteProductShapeDeleteCallback OnAfterDeleteInterface[NoteProductShape]
	OnAfterNoteProductShapeReadCallback   OnAfterReadInterface[NoteProductShape]

	NoteShapes                map[*NoteShape]struct{}
	NoteShapes_instance       map[*NoteShape]*NoteShape
	NoteShapes_mapString      map[string]*NoteShape
	NoteShapeOrder            uint
	NoteShape_stagedOrder     map[*NoteShape]uint
	NoteShape_orderStaged     map[uint]*NoteShape
	NoteShapes_reference      map[*NoteShape]*NoteShape
	NoteShapes_referenceOrder map[*NoteShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteShapeCreateCallback OnAfterCreateInterface[NoteShape]
	OnAfterNoteShapeUpdateCallback OnAfterUpdateInterface[NoteShape]
	OnAfterNoteShapeDeleteCallback OnAfterDeleteInterface[NoteShape]
	OnAfterNoteShapeReadCallback   OnAfterReadInterface[NoteShape]

	NoteStakeholderShapes                map[*NoteStakeholderShape]struct{}
	NoteStakeholderShapes_instance       map[*NoteStakeholderShape]*NoteStakeholderShape
	NoteStakeholderShapes_mapString      map[string]*NoteStakeholderShape
	NoteStakeholderShapeOrder            uint
	NoteStakeholderShape_stagedOrder     map[*NoteStakeholderShape]uint
	NoteStakeholderShape_orderStaged     map[uint]*NoteStakeholderShape
	NoteStakeholderShapes_reference      map[*NoteStakeholderShape]*NoteStakeholderShape
	NoteStakeholderShapes_referenceOrder map[*NoteStakeholderShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteStakeholderShapeCreateCallback OnAfterCreateInterface[NoteStakeholderShape]
	OnAfterNoteStakeholderShapeUpdateCallback OnAfterUpdateInterface[NoteStakeholderShape]
	OnAfterNoteStakeholderShapeDeleteCallback OnAfterDeleteInterface[NoteStakeholderShape]
	OnAfterNoteStakeholderShapeReadCallback   OnAfterReadInterface[NoteStakeholderShape]

	NoteTaskShapes                map[*NoteTaskShape]struct{}
	NoteTaskShapes_instance       map[*NoteTaskShape]*NoteTaskShape
	NoteTaskShapes_mapString      map[string]*NoteTaskShape
	NoteTaskShapeOrder            uint
	NoteTaskShape_stagedOrder     map[*NoteTaskShape]uint
	NoteTaskShape_orderStaged     map[uint]*NoteTaskShape
	NoteTaskShapes_reference      map[*NoteTaskShape]*NoteTaskShape
	NoteTaskShapes_referenceOrder map[*NoteTaskShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteTaskShapeCreateCallback OnAfterCreateInterface[NoteTaskShape]
	OnAfterNoteTaskShapeUpdateCallback OnAfterUpdateInterface[NoteTaskShape]
	OnAfterNoteTaskShapeDeleteCallback OnAfterDeleteInterface[NoteTaskShape]
	OnAfterNoteTaskShapeReadCallback   OnAfterReadInterface[NoteTaskShape]

	ProductCompositionShapes                map[*ProductCompositionShape]struct{}
	ProductCompositionShapes_instance       map[*ProductCompositionShape]*ProductCompositionShape
	ProductCompositionShapes_mapString      map[string]*ProductCompositionShape
	ProductCompositionShapeOrder            uint
	ProductCompositionShape_stagedOrder     map[*ProductCompositionShape]uint
	ProductCompositionShape_orderStaged     map[uint]*ProductCompositionShape
	ProductCompositionShapes_reference      map[*ProductCompositionShape]*ProductCompositionShape
	ProductCompositionShapes_referenceOrder map[*ProductCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterProductCompositionShapeCreateCallback OnAfterCreateInterface[ProductCompositionShape]
	OnAfterProductCompositionShapeUpdateCallback OnAfterUpdateInterface[ProductCompositionShape]
	OnAfterProductCompositionShapeDeleteCallback OnAfterDeleteInterface[ProductCompositionShape]
	OnAfterProductCompositionShapeReadCallback   OnAfterReadInterface[ProductCompositionShape]

	ProductShapes                map[*ProductShape]struct{}
	ProductShapes_instance       map[*ProductShape]*ProductShape
	ProductShapes_mapString      map[string]*ProductShape
	ProductShapeOrder            uint
	ProductShape_stagedOrder     map[*ProductShape]uint
	ProductShape_orderStaged     map[uint]*ProductShape
	ProductShapes_reference      map[*ProductShape]*ProductShape
	ProductShapes_referenceOrder map[*ProductShape]uint

	// insertion point for slice of pointers maps
	OnAfterProductShapeCreateCallback OnAfterCreateInterface[ProductShape]
	OnAfterProductShapeUpdateCallback OnAfterUpdateInterface[ProductShape]
	OnAfterProductShapeDeleteCallback OnAfterDeleteInterface[ProductShape]
	OnAfterProductShapeReadCallback   OnAfterReadInterface[ProductShape]

	Requirements                map[*Requirement]struct{}
	Requirements_instance       map[*Requirement]*Requirement
	Requirements_mapString      map[string]*Requirement
	RequirementOrder            uint
	Requirement_stagedOrder     map[*Requirement]uint
	Requirement_orderStaged     map[uint]*Requirement
	Requirements_reference      map[*Requirement]*Requirement
	Requirements_referenceOrder map[*Requirement]uint

	// insertion point for slice of pointers maps
	Requirement_SupportLevels_reverseMap map[*SupportLevel]*Requirement

	Requirement_Concepts_reverseMap map[*Concept]*Requirement

	OnAfterRequirementCreateCallback OnAfterCreateInterface[Requirement]
	OnAfterRequirementUpdateCallback OnAfterUpdateInterface[Requirement]
	OnAfterRequirementDeleteCallback OnAfterDeleteInterface[Requirement]
	OnAfterRequirementReadCallback   OnAfterReadInterface[Requirement]

	Stakeholders                map[*Stakeholder]struct{}
	Stakeholders_instance       map[*Stakeholder]*Stakeholder
	Stakeholders_mapString      map[string]*Stakeholder
	StakeholderOrder            uint
	Stakeholder_stagedOrder     map[*Stakeholder]uint
	Stakeholder_orderStaged     map[uint]*Stakeholder
	Stakeholders_reference      map[*Stakeholder]*Stakeholder
	Stakeholders_referenceOrder map[*Stakeholder]uint

	// insertion point for slice of pointers maps
	Stakeholder_Concerns_reverseMap map[*Concern]*Stakeholder

	Stakeholder_SubStakeholders_reverseMap map[*Stakeholder]*Stakeholder

	OnAfterStakeholderCreateCallback OnAfterCreateInterface[Stakeholder]
	OnAfterStakeholderUpdateCallback OnAfterUpdateInterface[Stakeholder]
	OnAfterStakeholderDeleteCallback OnAfterDeleteInterface[Stakeholder]
	OnAfterStakeholderReadCallback   OnAfterReadInterface[Stakeholder]

	StakeholderCompositionShapes                map[*StakeholderCompositionShape]struct{}
	StakeholderCompositionShapes_instance       map[*StakeholderCompositionShape]*StakeholderCompositionShape
	StakeholderCompositionShapes_mapString      map[string]*StakeholderCompositionShape
	StakeholderCompositionShapeOrder            uint
	StakeholderCompositionShape_stagedOrder     map[*StakeholderCompositionShape]uint
	StakeholderCompositionShape_orderStaged     map[uint]*StakeholderCompositionShape
	StakeholderCompositionShapes_reference      map[*StakeholderCompositionShape]*StakeholderCompositionShape
	StakeholderCompositionShapes_referenceOrder map[*StakeholderCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterStakeholderCompositionShapeCreateCallback OnAfterCreateInterface[StakeholderCompositionShape]
	OnAfterStakeholderCompositionShapeUpdateCallback OnAfterUpdateInterface[StakeholderCompositionShape]
	OnAfterStakeholderCompositionShapeDeleteCallback OnAfterDeleteInterface[StakeholderCompositionShape]
	OnAfterStakeholderCompositionShapeReadCallback   OnAfterReadInterface[StakeholderCompositionShape]

	StakeholderConcernShapes                map[*StakeholderConcernShape]struct{}
	StakeholderConcernShapes_instance       map[*StakeholderConcernShape]*StakeholderConcernShape
	StakeholderConcernShapes_mapString      map[string]*StakeholderConcernShape
	StakeholderConcernShapeOrder            uint
	StakeholderConcernShape_stagedOrder     map[*StakeholderConcernShape]uint
	StakeholderConcernShape_orderStaged     map[uint]*StakeholderConcernShape
	StakeholderConcernShapes_reference      map[*StakeholderConcernShape]*StakeholderConcernShape
	StakeholderConcernShapes_referenceOrder map[*StakeholderConcernShape]uint

	// insertion point for slice of pointers maps
	OnAfterStakeholderConcernShapeCreateCallback OnAfterCreateInterface[StakeholderConcernShape]
	OnAfterStakeholderConcernShapeUpdateCallback OnAfterUpdateInterface[StakeholderConcernShape]
	OnAfterStakeholderConcernShapeDeleteCallback OnAfterDeleteInterface[StakeholderConcernShape]
	OnAfterStakeholderConcernShapeReadCallback   OnAfterReadInterface[StakeholderConcernShape]

	StakeholderShapes                map[*StakeholderShape]struct{}
	StakeholderShapes_instance       map[*StakeholderShape]*StakeholderShape
	StakeholderShapes_mapString      map[string]*StakeholderShape
	StakeholderShapeOrder            uint
	StakeholderShape_stagedOrder     map[*StakeholderShape]uint
	StakeholderShape_orderStaged     map[uint]*StakeholderShape
	StakeholderShapes_reference      map[*StakeholderShape]*StakeholderShape
	StakeholderShapes_referenceOrder map[*StakeholderShape]uint

	// insertion point for slice of pointers maps
	OnAfterStakeholderShapeCreateCallback OnAfterCreateInterface[StakeholderShape]
	OnAfterStakeholderShapeUpdateCallback OnAfterUpdateInterface[StakeholderShape]
	OnAfterStakeholderShapeDeleteCallback OnAfterDeleteInterface[StakeholderShape]
	OnAfterStakeholderShapeReadCallback   OnAfterReadInterface[StakeholderShape]

	SupportLevels                map[*SupportLevel]struct{}
	SupportLevels_instance       map[*SupportLevel]*SupportLevel
	SupportLevels_mapString      map[string]*SupportLevel
	SupportLevelOrder            uint
	SupportLevel_stagedOrder     map[*SupportLevel]uint
	SupportLevel_orderStaged     map[uint]*SupportLevel
	SupportLevels_reference      map[*SupportLevel]*SupportLevel
	SupportLevels_referenceOrder map[*SupportLevel]uint

	// insertion point for slice of pointers maps
	OnAfterSupportLevelCreateCallback OnAfterCreateInterface[SupportLevel]
	OnAfterSupportLevelUpdateCallback OnAfterUpdateInterface[SupportLevel]
	OnAfterSupportLevelDeleteCallback OnAfterDeleteInterface[SupportLevel]
	OnAfterSupportLevelReadCallback   OnAfterReadInterface[SupportLevel]

	Tools                map[*Tool]struct{}
	Tools_instance       map[*Tool]*Tool
	Tools_mapString      map[string]*Tool
	ToolOrder            uint
	Tool_stagedOrder     map[*Tool]uint
	Tool_orderStaged     map[uint]*Tool
	Tools_reference      map[*Tool]*Tool
	Tools_referenceOrder map[*Tool]uint

	// insertion point for slice of pointers maps
	OnAfterToolCreateCallback OnAfterCreateInterface[Tool]
	OnAfterToolUpdateCallback OnAfterUpdateInterface[Tool]
	OnAfterToolDeleteCallback OnAfterDeleteInterface[Tool]
	OnAfterToolReadCallback   OnAfterReadInterface[Tool]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	// end of insertion point

	NamedStructs []*NamedStruct

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]ModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF ProbeIF

	forwardCommits  []string
	backwardCommits []string

	// when navigating the commit history
	// navigationMode is set to Navigating
	navigationMode gongStageNavigationMode
	commitsBehind  int // the number of commits the stage is behind the front of the history

	isApplyingBackwardCommit bool
	isApplyingForwardCommit  bool
	isSquashing              bool

	modified bool

	lock sync.RWMutex
}

func (s *Stage) SetGongMarshallingMode(mode GongMarshallingMode) {
	s.gongMarshallingMode = mode
}

func (s *Stage) GetGongMarshallingMode() GongMarshallingMode {
	return s.gongMarshallingMode
}

func (s *Stage) SetIsWithGenesisCommit(isWithGenesisCommit bool) {
	s.isWithGenesisCommit = isWithGenesisCommit
}

func (s *Stage) GetIsWithGenesisCommit() bool {
	return s.isWithGenesisCommit
}

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.beforeCommitHooks = append(s.beforeCommitHooks, hook)
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.afterCommitHooks = append(s.afterCommitHooks, hook)
}

type gongStageNavigationMode string

const (
	GongNavigationModeNormal gongStageNavigationMode = "Normal"
	// when the mode is navigating, each commit backward and forward
	// it is possible to go apply the nbCommitsBackward forward commits
	GongNavigationModeNavigating gongStageNavigationMode = "Navigating"
)

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	if len(stage.backwardCommits) == 0 {
		return errors.New("no backward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	if stage.isWithGenesisCommit && stage.commitsBehind >= len(stage.backwardCommits)-1 {
		return errors.New("cannot rollback genesis commit")
	}

	if stage.commitsBehind >= len(stage.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := stage.backwardCommits[len(stage.backwardCommits)-1-stage.commitsBehind]

	// umarshall the backward commit to the stage

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind++
	stage.isApplyingBackwardCommit = true
	err := GongParseAstString(stage, commitToApply, true)
	stage.isApplyingBackwardCommit = false
	if err != nil {
		log.Println("error during ApplyBackwardCommit: ", err)
		return err
	}

	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetForwardCommits() []string {
	return stage.forwardCommits
}

func (stage *Stage) GetBackwardCommits() []string {
	return stage.backwardCommits
}

func (stage *Stage) ApplyForwardCommit() error {
	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.commitsBehind == 0 {
		return errors.New("no more forward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	commitToApply := stage.forwardCommits[len(stage.forwardCommits)-1-stage.commitsBehind+1]

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind--
	stage.isApplyingForwardCommit = true
	err := GongParseAstString(stage, commitToApply, true)
	stage.isApplyingForwardCommit = false
	if err != nil {
		log.Println("error during ApplyForwardCommit: ", err)
		return err
	}
	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetCommitsBehind() int {
	return stage.commitsBehind
}

func (stage *Stage) Lock() {
	stage.lock.Lock()
}

func (stage *Stage) Unlock() {
	stage.lock.Unlock()
}

func (stage *Stage) RLock() {
	stage.lock.RLock()
}

func (stage *Stage) RUnlock() {
	stage.lock.RUnlock()
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	newCommitsLen := len(stage.forwardCommits) - stage.GetCommitsBehind()

	stage.forwardCommits = stage.forwardCommits[:newCommitsLen]
	stage.backwardCommits = stage.backwardCommits[:newCommitsLen]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.forwardCommits = stage.forwardCommits[:0]
	stage.backwardCommits = stage.backwardCommits[:0]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.modified = true
	stage.isSquashing = true

	// insertion point for clear references
	stage.AnalysisNeeds_reference = make(map[*AnalysisNeed]*AnalysisNeed)
	stage.AnalysisNeeds_instance = make(map[*AnalysisNeed]*AnalysisNeed)
	stage.AnalysisNeeds_referenceOrder = make(map[*AnalysisNeed]uint)

	stage.Concepts_reference = make(map[*Concept]*Concept)
	stage.Concepts_instance = make(map[*Concept]*Concept)
	stage.Concepts_referenceOrder = make(map[*Concept]uint)

	stage.Concerns_reference = make(map[*Concern]*Concern)
	stage.Concerns_instance = make(map[*Concern]*Concern)
	stage.Concerns_referenceOrder = make(map[*Concern]uint)

	stage.ConcernCompositionShapes_reference = make(map[*ConcernCompositionShape]*ConcernCompositionShape)
	stage.ConcernCompositionShapes_instance = make(map[*ConcernCompositionShape]*ConcernCompositionShape)
	stage.ConcernCompositionShapes_referenceOrder = make(map[*ConcernCompositionShape]uint)

	stage.ConcernInputShapes_reference = make(map[*ConcernInputShape]*ConcernInputShape)
	stage.ConcernInputShapes_instance = make(map[*ConcernInputShape]*ConcernInputShape)
	stage.ConcernInputShapes_referenceOrder = make(map[*ConcernInputShape]uint)

	stage.ConcernOutputShapes_reference = make(map[*ConcernOutputShape]*ConcernOutputShape)
	stage.ConcernOutputShapes_instance = make(map[*ConcernOutputShape]*ConcernOutputShape)
	stage.ConcernOutputShapes_referenceOrder = make(map[*ConcernOutputShape]uint)

	stage.ConcernShapes_reference = make(map[*ConcernShape]*ConcernShape)
	stage.ConcernShapes_instance = make(map[*ConcernShape]*ConcernShape)
	stage.ConcernShapes_referenceOrder = make(map[*ConcernShape]uint)

	stage.Deliverables_reference = make(map[*Deliverable]*Deliverable)
	stage.Deliverables_instance = make(map[*Deliverable]*Deliverable)
	stage.Deliverables_referenceOrder = make(map[*Deliverable]uint)

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_instance = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint)

	stage.NoteProductShapes_reference = make(map[*NoteProductShape]*NoteProductShape)
	stage.NoteProductShapes_instance = make(map[*NoteProductShape]*NoteProductShape)
	stage.NoteProductShapes_referenceOrder = make(map[*NoteProductShape]uint)

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)

	stage.NoteStakeholderShapes_reference = make(map[*NoteStakeholderShape]*NoteStakeholderShape)
	stage.NoteStakeholderShapes_instance = make(map[*NoteStakeholderShape]*NoteStakeholderShape)
	stage.NoteStakeholderShapes_referenceOrder = make(map[*NoteStakeholderShape]uint)

	stage.NoteTaskShapes_reference = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_instance = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint)

	stage.ProductCompositionShapes_reference = make(map[*ProductCompositionShape]*ProductCompositionShape)
	stage.ProductCompositionShapes_instance = make(map[*ProductCompositionShape]*ProductCompositionShape)
	stage.ProductCompositionShapes_referenceOrder = make(map[*ProductCompositionShape]uint)

	stage.ProductShapes_reference = make(map[*ProductShape]*ProductShape)
	stage.ProductShapes_instance = make(map[*ProductShape]*ProductShape)
	stage.ProductShapes_referenceOrder = make(map[*ProductShape]uint)

	stage.Requirements_reference = make(map[*Requirement]*Requirement)
	stage.Requirements_instance = make(map[*Requirement]*Requirement)
	stage.Requirements_referenceOrder = make(map[*Requirement]uint)

	stage.Stakeholders_reference = make(map[*Stakeholder]*Stakeholder)
	stage.Stakeholders_instance = make(map[*Stakeholder]*Stakeholder)
	stage.Stakeholders_referenceOrder = make(map[*Stakeholder]uint)

	stage.StakeholderCompositionShapes_reference = make(map[*StakeholderCompositionShape]*StakeholderCompositionShape)
	stage.StakeholderCompositionShapes_instance = make(map[*StakeholderCompositionShape]*StakeholderCompositionShape)
	stage.StakeholderCompositionShapes_referenceOrder = make(map[*StakeholderCompositionShape]uint)

	stage.StakeholderConcernShapes_reference = make(map[*StakeholderConcernShape]*StakeholderConcernShape)
	stage.StakeholderConcernShapes_instance = make(map[*StakeholderConcernShape]*StakeholderConcernShape)
	stage.StakeholderConcernShapes_referenceOrder = make(map[*StakeholderConcernShape]uint)

	stage.StakeholderShapes_reference = make(map[*StakeholderShape]*StakeholderShape)
	stage.StakeholderShapes_instance = make(map[*StakeholderShape]*StakeholderShape)
	stage.StakeholderShapes_referenceOrder = make(map[*StakeholderShape]uint)

	stage.SupportLevels_reference = make(map[*SupportLevel]*SupportLevel)
	stage.SupportLevels_instance = make(map[*SupportLevel]*SupportLevel)
	stage.SupportLevels_referenceOrder = make(map[*SupportLevel]uint)

	stage.Tools_reference = make(map[*Tool]*Tool)
	stage.Tools_instance = make(map[*Tool]*Tool)
	stage.Tools_referenceOrder = make(map[*Tool]uint)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}

	stage.isSquashing = false
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	var maxAnalysisNeedOrder uint
	var foundAnalysisNeed bool
	for _, order := range stage.AnalysisNeed_stagedOrder {
		if !foundAnalysisNeed || order > maxAnalysisNeedOrder {
			maxAnalysisNeedOrder = order
			foundAnalysisNeed = true
		}
	}
	if foundAnalysisNeed {
		stage.AnalysisNeedOrder = maxAnalysisNeedOrder + 1
	} else {
		stage.AnalysisNeedOrder = 0
	}

	var maxConceptOrder uint
	var foundConcept bool
	for _, order := range stage.Concept_stagedOrder {
		if !foundConcept || order > maxConceptOrder {
			maxConceptOrder = order
			foundConcept = true
		}
	}
	if foundConcept {
		stage.ConceptOrder = maxConceptOrder + 1
	} else {
		stage.ConceptOrder = 0
	}

	var maxConcernOrder uint
	var foundConcern bool
	for _, order := range stage.Concern_stagedOrder {
		if !foundConcern || order > maxConcernOrder {
			maxConcernOrder = order
			foundConcern = true
		}
	}
	if foundConcern {
		stage.ConcernOrder = maxConcernOrder + 1
	} else {
		stage.ConcernOrder = 0
	}

	var maxConcernCompositionShapeOrder uint
	var foundConcernCompositionShape bool
	for _, order := range stage.ConcernCompositionShape_stagedOrder {
		if !foundConcernCompositionShape || order > maxConcernCompositionShapeOrder {
			maxConcernCompositionShapeOrder = order
			foundConcernCompositionShape = true
		}
	}
	if foundConcernCompositionShape {
		stage.ConcernCompositionShapeOrder = maxConcernCompositionShapeOrder + 1
	} else {
		stage.ConcernCompositionShapeOrder = 0
	}

	var maxConcernInputShapeOrder uint
	var foundConcernInputShape bool
	for _, order := range stage.ConcernInputShape_stagedOrder {
		if !foundConcernInputShape || order > maxConcernInputShapeOrder {
			maxConcernInputShapeOrder = order
			foundConcernInputShape = true
		}
	}
	if foundConcernInputShape {
		stage.ConcernInputShapeOrder = maxConcernInputShapeOrder + 1
	} else {
		stage.ConcernInputShapeOrder = 0
	}

	var maxConcernOutputShapeOrder uint
	var foundConcernOutputShape bool
	for _, order := range stage.ConcernOutputShape_stagedOrder {
		if !foundConcernOutputShape || order > maxConcernOutputShapeOrder {
			maxConcernOutputShapeOrder = order
			foundConcernOutputShape = true
		}
	}
	if foundConcernOutputShape {
		stage.ConcernOutputShapeOrder = maxConcernOutputShapeOrder + 1
	} else {
		stage.ConcernOutputShapeOrder = 0
	}

	var maxConcernShapeOrder uint
	var foundConcernShape bool
	for _, order := range stage.ConcernShape_stagedOrder {
		if !foundConcernShape || order > maxConcernShapeOrder {
			maxConcernShapeOrder = order
			foundConcernShape = true
		}
	}
	if foundConcernShape {
		stage.ConcernShapeOrder = maxConcernShapeOrder + 1
	} else {
		stage.ConcernShapeOrder = 0
	}

	var maxDeliverableOrder uint
	var foundDeliverable bool
	for _, order := range stage.Deliverable_stagedOrder {
		if !foundDeliverable || order > maxDeliverableOrder {
			maxDeliverableOrder = order
			foundDeliverable = true
		}
	}
	if foundDeliverable {
		stage.DeliverableOrder = maxDeliverableOrder + 1
	} else {
		stage.DeliverableOrder = 0
	}

	var maxDiagramOrder uint
	var foundDiagram bool
	for _, order := range stage.Diagram_stagedOrder {
		if !foundDiagram || order > maxDiagramOrder {
			maxDiagramOrder = order
			foundDiagram = true
		}
	}
	if foundDiagram {
		stage.DiagramOrder = maxDiagramOrder + 1
	} else {
		stage.DiagramOrder = 0
	}

	var maxLibraryOrder uint
	var foundLibrary bool
	for _, order := range stage.Library_stagedOrder {
		if !foundLibrary || order > maxLibraryOrder {
			maxLibraryOrder = order
			foundLibrary = true
		}
	}
	if foundLibrary {
		stage.LibraryOrder = maxLibraryOrder + 1
	} else {
		stage.LibraryOrder = 0
	}

	var maxNoteOrder uint
	var foundNote bool
	for _, order := range stage.Note_stagedOrder {
		if !foundNote || order > maxNoteOrder {
			maxNoteOrder = order
			foundNote = true
		}
	}
	if foundNote {
		stage.NoteOrder = maxNoteOrder + 1
	} else {
		stage.NoteOrder = 0
	}

	var maxNoteProductShapeOrder uint
	var foundNoteProductShape bool
	for _, order := range stage.NoteProductShape_stagedOrder {
		if !foundNoteProductShape || order > maxNoteProductShapeOrder {
			maxNoteProductShapeOrder = order
			foundNoteProductShape = true
		}
	}
	if foundNoteProductShape {
		stage.NoteProductShapeOrder = maxNoteProductShapeOrder + 1
	} else {
		stage.NoteProductShapeOrder = 0
	}

	var maxNoteShapeOrder uint
	var foundNoteShape bool
	for _, order := range stage.NoteShape_stagedOrder {
		if !foundNoteShape || order > maxNoteShapeOrder {
			maxNoteShapeOrder = order
			foundNoteShape = true
		}
	}
	if foundNoteShape {
		stage.NoteShapeOrder = maxNoteShapeOrder + 1
	} else {
		stage.NoteShapeOrder = 0
	}

	var maxNoteStakeholderShapeOrder uint
	var foundNoteStakeholderShape bool
	for _, order := range stage.NoteStakeholderShape_stagedOrder {
		if !foundNoteStakeholderShape || order > maxNoteStakeholderShapeOrder {
			maxNoteStakeholderShapeOrder = order
			foundNoteStakeholderShape = true
		}
	}
	if foundNoteStakeholderShape {
		stage.NoteStakeholderShapeOrder = maxNoteStakeholderShapeOrder + 1
	} else {
		stage.NoteStakeholderShapeOrder = 0
	}

	var maxNoteTaskShapeOrder uint
	var foundNoteTaskShape bool
	for _, order := range stage.NoteTaskShape_stagedOrder {
		if !foundNoteTaskShape || order > maxNoteTaskShapeOrder {
			maxNoteTaskShapeOrder = order
			foundNoteTaskShape = true
		}
	}
	if foundNoteTaskShape {
		stage.NoteTaskShapeOrder = maxNoteTaskShapeOrder + 1
	} else {
		stage.NoteTaskShapeOrder = 0
	}

	var maxProductCompositionShapeOrder uint
	var foundProductCompositionShape bool
	for _, order := range stage.ProductCompositionShape_stagedOrder {
		if !foundProductCompositionShape || order > maxProductCompositionShapeOrder {
			maxProductCompositionShapeOrder = order
			foundProductCompositionShape = true
		}
	}
	if foundProductCompositionShape {
		stage.ProductCompositionShapeOrder = maxProductCompositionShapeOrder + 1
	} else {
		stage.ProductCompositionShapeOrder = 0
	}

	var maxProductShapeOrder uint
	var foundProductShape bool
	for _, order := range stage.ProductShape_stagedOrder {
		if !foundProductShape || order > maxProductShapeOrder {
			maxProductShapeOrder = order
			foundProductShape = true
		}
	}
	if foundProductShape {
		stage.ProductShapeOrder = maxProductShapeOrder + 1
	} else {
		stage.ProductShapeOrder = 0
	}

	var maxRequirementOrder uint
	var foundRequirement bool
	for _, order := range stage.Requirement_stagedOrder {
		if !foundRequirement || order > maxRequirementOrder {
			maxRequirementOrder = order
			foundRequirement = true
		}
	}
	if foundRequirement {
		stage.RequirementOrder = maxRequirementOrder + 1
	} else {
		stage.RequirementOrder = 0
	}

	var maxStakeholderOrder uint
	var foundStakeholder bool
	for _, order := range stage.Stakeholder_stagedOrder {
		if !foundStakeholder || order > maxStakeholderOrder {
			maxStakeholderOrder = order
			foundStakeholder = true
		}
	}
	if foundStakeholder {
		stage.StakeholderOrder = maxStakeholderOrder + 1
	} else {
		stage.StakeholderOrder = 0
	}

	var maxStakeholderCompositionShapeOrder uint
	var foundStakeholderCompositionShape bool
	for _, order := range stage.StakeholderCompositionShape_stagedOrder {
		if !foundStakeholderCompositionShape || order > maxStakeholderCompositionShapeOrder {
			maxStakeholderCompositionShapeOrder = order
			foundStakeholderCompositionShape = true
		}
	}
	if foundStakeholderCompositionShape {
		stage.StakeholderCompositionShapeOrder = maxStakeholderCompositionShapeOrder + 1
	} else {
		stage.StakeholderCompositionShapeOrder = 0
	}

	var maxStakeholderConcernShapeOrder uint
	var foundStakeholderConcernShape bool
	for _, order := range stage.StakeholderConcernShape_stagedOrder {
		if !foundStakeholderConcernShape || order > maxStakeholderConcernShapeOrder {
			maxStakeholderConcernShapeOrder = order
			foundStakeholderConcernShape = true
		}
	}
	if foundStakeholderConcernShape {
		stage.StakeholderConcernShapeOrder = maxStakeholderConcernShapeOrder + 1
	} else {
		stage.StakeholderConcernShapeOrder = 0
	}

	var maxStakeholderShapeOrder uint
	var foundStakeholderShape bool
	for _, order := range stage.StakeholderShape_stagedOrder {
		if !foundStakeholderShape || order > maxStakeholderShapeOrder {
			maxStakeholderShapeOrder = order
			foundStakeholderShape = true
		}
	}
	if foundStakeholderShape {
		stage.StakeholderShapeOrder = maxStakeholderShapeOrder + 1
	} else {
		stage.StakeholderShapeOrder = 0
	}

	var maxSupportLevelOrder uint
	var foundSupportLevel bool
	for _, order := range stage.SupportLevel_stagedOrder {
		if !foundSupportLevel || order > maxSupportLevelOrder {
			maxSupportLevelOrder = order
			foundSupportLevel = true
		}
	}
	if foundSupportLevel {
		stage.SupportLevelOrder = maxSupportLevelOrder + 1
	} else {
		stage.SupportLevelOrder = 0
	}

	var maxToolOrder uint
	var foundTool bool
	for _, order := range stage.Tool_stagedOrder {
		if !foundTool || order > maxToolOrder {
			maxToolOrder = order
			foundTool = true
		}
	}
	if foundTool {
		stage.ToolOrder = maxToolOrder + 1
	} else {
		stage.ToolOrder = 0
	}

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF ProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() ProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {
	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []string) {
	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance.GetName())
	}

	return
}

// GetStructInstancesByOrderAuto returns a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *AnalysisNeed:
		tmp := GetStructInstancesByOrder(stage.AnalysisNeeds, stage.AnalysisNeed_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AnalysisNeed implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Concept:
		tmp := GetStructInstancesByOrder(stage.Concepts, stage.Concept_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Concept implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Concern:
		tmp := GetStructInstancesByOrder(stage.Concerns, stage.Concern_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Concern implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ConcernCompositionShape:
		tmp := GetStructInstancesByOrder(stage.ConcernCompositionShapes, stage.ConcernCompositionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ConcernCompositionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ConcernInputShape:
		tmp := GetStructInstancesByOrder(stage.ConcernInputShapes, stage.ConcernInputShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ConcernInputShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ConcernOutputShape:
		tmp := GetStructInstancesByOrder(stage.ConcernOutputShapes, stage.ConcernOutputShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ConcernOutputShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ConcernShape:
		tmp := GetStructInstancesByOrder(stage.ConcernShapes, stage.ConcernShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ConcernShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Deliverable:
		tmp := GetStructInstancesByOrder(stage.Deliverables, stage.Deliverable_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Deliverable implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Diagram:
		tmp := GetStructInstancesByOrder(stage.Diagrams, stage.Diagram_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Diagram implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Library:
		tmp := GetStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Library implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Note:
		tmp := GetStructInstancesByOrder(stage.Notes, stage.Note_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Note implements.
			res = append(res, any(v).(T))
		}
		return res
	case *NoteProductShape:
		tmp := GetStructInstancesByOrder(stage.NoteProductShapes, stage.NoteProductShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteProductShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *NoteShape:
		tmp := GetStructInstancesByOrder(stage.NoteShapes, stage.NoteShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *NoteStakeholderShape:
		tmp := GetStructInstancesByOrder(stage.NoteStakeholderShapes, stage.NoteStakeholderShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteStakeholderShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *NoteTaskShape:
		tmp := GetStructInstancesByOrder(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteTaskShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ProductCompositionShape:
		tmp := GetStructInstancesByOrder(stage.ProductCompositionShapes, stage.ProductCompositionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ProductCompositionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ProductShape:
		tmp := GetStructInstancesByOrder(stage.ProductShapes, stage.ProductShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ProductShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Requirement:
		tmp := GetStructInstancesByOrder(stage.Requirements, stage.Requirement_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Requirement implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Stakeholder:
		tmp := GetStructInstancesByOrder(stage.Stakeholders, stage.Stakeholder_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Stakeholder implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StakeholderCompositionShape:
		tmp := GetStructInstancesByOrder(stage.StakeholderCompositionShapes, stage.StakeholderCompositionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StakeholderCompositionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StakeholderConcernShape:
		tmp := GetStructInstancesByOrder(stage.StakeholderConcernShapes, stage.StakeholderConcernShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StakeholderConcernShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StakeholderShape:
		tmp := GetStructInstancesByOrder(stage.StakeholderShapes, stage.StakeholderShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StakeholderShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SupportLevel:
		tmp := GetStructInstancesByOrder(stage.SupportLevels, stage.SupportLevel_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SupportLevel implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Tool:
		tmp := GetStructInstancesByOrder(stage.Tools, stage.Tool_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Tool implements.
			res = append(res, any(v).(T))
		}
		return res

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []T) {
	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {
	switch namedStructName {
	// insertion point for case
	case "AnalysisNeed":
		res = GetNamedStructInstances(stage.AnalysisNeeds, stage.AnalysisNeed_stagedOrder)
	case "Concept":
		res = GetNamedStructInstances(stage.Concepts, stage.Concept_stagedOrder)
	case "Concern":
		res = GetNamedStructInstances(stage.Concerns, stage.Concern_stagedOrder)
	case "ConcernCompositionShape":
		res = GetNamedStructInstances(stage.ConcernCompositionShapes, stage.ConcernCompositionShape_stagedOrder)
	case "ConcernInputShape":
		res = GetNamedStructInstances(stage.ConcernInputShapes, stage.ConcernInputShape_stagedOrder)
	case "ConcernOutputShape":
		res = GetNamedStructInstances(stage.ConcernOutputShapes, stage.ConcernOutputShape_stagedOrder)
	case "ConcernShape":
		res = GetNamedStructInstances(stage.ConcernShapes, stage.ConcernShape_stagedOrder)
	case "Deliverable":
		res = GetNamedStructInstances(stage.Deliverables, stage.Deliverable_stagedOrder)
	case "Diagram":
		res = GetNamedStructInstances(stage.Diagrams, stage.Diagram_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Note":
		res = GetNamedStructInstances(stage.Notes, stage.Note_stagedOrder)
	case "NoteProductShape":
		res = GetNamedStructInstances(stage.NoteProductShapes, stage.NoteProductShape_stagedOrder)
	case "NoteShape":
		res = GetNamedStructInstances(stage.NoteShapes, stage.NoteShape_stagedOrder)
	case "NoteStakeholderShape":
		res = GetNamedStructInstances(stage.NoteStakeholderShapes, stage.NoteStakeholderShape_stagedOrder)
	case "NoteTaskShape":
		res = GetNamedStructInstances(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder)
	case "ProductCompositionShape":
		res = GetNamedStructInstances(stage.ProductCompositionShapes, stage.ProductCompositionShape_stagedOrder)
	case "ProductShape":
		res = GetNamedStructInstances(stage.ProductShapes, stage.ProductShape_stagedOrder)
	case "Requirement":
		res = GetNamedStructInstances(stage.Requirements, stage.Requirement_stagedOrder)
	case "Stakeholder":
		res = GetNamedStructInstances(stage.Stakeholders, stage.Stakeholder_stagedOrder)
	case "StakeholderCompositionShape":
		res = GetNamedStructInstances(stage.StakeholderCompositionShapes, stage.StakeholderCompositionShape_stagedOrder)
	case "StakeholderConcernShape":
		res = GetNamedStructInstances(stage.StakeholderConcernShapes, stage.StakeholderConcernShape_stagedOrder)
	case "StakeholderShape":
		res = GetNamedStructInstances(stage.StakeholderShapes, stage.StakeholderShape_stagedOrder)
	case "SupportLevel":
		res = GetNamedStructInstances(stage.SupportLevels, stage.SupportLevel_stagedOrder)
	case "Tool":
		res = GetNamedStructInstances(stage.Tools, stage.Tool_stagedOrder)
	}

	return
}

type NamedStruct struct {
	name string
}

func (namedStruct *NamedStruct) GetName() string {
	return namedStruct.name
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/dsm/capture/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return capture_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return capture_go.GoDiagramsDir
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitAnalysisNeed(analysisneed *AnalysisNeed)
	CheckoutAnalysisNeed(analysisneed *AnalysisNeed)
	CommitConcept(concept *Concept)
	CheckoutConcept(concept *Concept)
	CommitConcern(concern *Concern)
	CheckoutConcern(concern *Concern)
	CommitConcernCompositionShape(concerncompositionshape *ConcernCompositionShape)
	CheckoutConcernCompositionShape(concerncompositionshape *ConcernCompositionShape)
	CommitConcernInputShape(concerninputshape *ConcernInputShape)
	CheckoutConcernInputShape(concerninputshape *ConcernInputShape)
	CommitConcernOutputShape(concernoutputshape *ConcernOutputShape)
	CheckoutConcernOutputShape(concernoutputshape *ConcernOutputShape)
	CommitConcernShape(concernshape *ConcernShape)
	CheckoutConcernShape(concernshape *ConcernShape)
	CommitDeliverable(deliverable *Deliverable)
	CheckoutDeliverable(deliverable *Deliverable)
	CommitDiagram(diagram *Diagram)
	CheckoutDiagram(diagram *Diagram)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitNote(note *Note)
	CheckoutNote(note *Note)
	CommitNoteProductShape(noteproductshape *NoteProductShape)
	CheckoutNoteProductShape(noteproductshape *NoteProductShape)
	CommitNoteShape(noteshape *NoteShape)
	CheckoutNoteShape(noteshape *NoteShape)
	CommitNoteStakeholderShape(notestakeholdershape *NoteStakeholderShape)
	CheckoutNoteStakeholderShape(notestakeholdershape *NoteStakeholderShape)
	CommitNoteTaskShape(notetaskshape *NoteTaskShape)
	CheckoutNoteTaskShape(notetaskshape *NoteTaskShape)
	CommitProductCompositionShape(productcompositionshape *ProductCompositionShape)
	CheckoutProductCompositionShape(productcompositionshape *ProductCompositionShape)
	CommitProductShape(productshape *ProductShape)
	CheckoutProductShape(productshape *ProductShape)
	CommitRequirement(requirement *Requirement)
	CheckoutRequirement(requirement *Requirement)
	CommitStakeholder(stakeholder *Stakeholder)
	CheckoutStakeholder(stakeholder *Stakeholder)
	CommitStakeholderCompositionShape(stakeholdercompositionshape *StakeholderCompositionShape)
	CheckoutStakeholderCompositionShape(stakeholdercompositionshape *StakeholderCompositionShape)
	CommitStakeholderConcernShape(stakeholderconcernshape *StakeholderConcernShape)
	CheckoutStakeholderConcernShape(stakeholderconcernshape *StakeholderConcernShape)
	CommitStakeholderShape(stakeholdershape *StakeholderShape)
	CheckoutStakeholderShape(stakeholdershape *StakeholderShape)
	CommitSupportLevel(supportlevel *SupportLevel)
	CheckoutSupportLevel(supportlevel *SupportLevel)
	CommitTool(tool *Tool)
	CheckoutTool(tool *Tool)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		AnalysisNeeds:           make(map[*AnalysisNeed]struct{}),
		AnalysisNeeds_mapString: make(map[string]*AnalysisNeed),

		Concepts:           make(map[*Concept]struct{}),
		Concepts_mapString: make(map[string]*Concept),

		Concerns:           make(map[*Concern]struct{}),
		Concerns_mapString: make(map[string]*Concern),

		ConcernCompositionShapes:           make(map[*ConcernCompositionShape]struct{}),
		ConcernCompositionShapes_mapString: make(map[string]*ConcernCompositionShape),

		ConcernInputShapes:           make(map[*ConcernInputShape]struct{}),
		ConcernInputShapes_mapString: make(map[string]*ConcernInputShape),

		ConcernOutputShapes:           make(map[*ConcernOutputShape]struct{}),
		ConcernOutputShapes_mapString: make(map[string]*ConcernOutputShape),

		ConcernShapes:           make(map[*ConcernShape]struct{}),
		ConcernShapes_mapString: make(map[string]*ConcernShape),

		Deliverables:           make(map[*Deliverable]struct{}),
		Deliverables_mapString: make(map[string]*Deliverable),

		Diagrams:           make(map[*Diagram]struct{}),
		Diagrams_mapString: make(map[string]*Diagram),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		NoteProductShapes:           make(map[*NoteProductShape]struct{}),
		NoteProductShapes_mapString: make(map[string]*NoteProductShape),

		NoteShapes:           make(map[*NoteShape]struct{}),
		NoteShapes_mapString: make(map[string]*NoteShape),

		NoteStakeholderShapes:           make(map[*NoteStakeholderShape]struct{}),
		NoteStakeholderShapes_mapString: make(map[string]*NoteStakeholderShape),

		NoteTaskShapes:           make(map[*NoteTaskShape]struct{}),
		NoteTaskShapes_mapString: make(map[string]*NoteTaskShape),

		ProductCompositionShapes:           make(map[*ProductCompositionShape]struct{}),
		ProductCompositionShapes_mapString: make(map[string]*ProductCompositionShape),

		ProductShapes:           make(map[*ProductShape]struct{}),
		ProductShapes_mapString: make(map[string]*ProductShape),

		Requirements:           make(map[*Requirement]struct{}),
		Requirements_mapString: make(map[string]*Requirement),

		Stakeholders:           make(map[*Stakeholder]struct{}),
		Stakeholders_mapString: make(map[string]*Stakeholder),

		StakeholderCompositionShapes:           make(map[*StakeholderCompositionShape]struct{}),
		StakeholderCompositionShapes_mapString: make(map[string]*StakeholderCompositionShape),

		StakeholderConcernShapes:           make(map[*StakeholderConcernShape]struct{}),
		StakeholderConcernShapes_mapString: make(map[string]*StakeholderConcernShape),

		StakeholderShapes:           make(map[*StakeholderShape]struct{}),
		StakeholderShapes_mapString: make(map[string]*StakeholderShape),

		SupportLevels:           make(map[*SupportLevel]struct{}),
		SupportLevels_mapString: make(map[string]*SupportLevel),

		Tools:           make(map[*Tool]struct{}),
		Tools_mapString: make(map[string]*Tool),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AnalysisNeed_stagedOrder: make(map[*AnalysisNeed]uint),
		AnalysisNeed_orderStaged: make(map[uint]*AnalysisNeed),
		AnalysisNeeds_reference:  make(map[*AnalysisNeed]*AnalysisNeed),

		Concept_stagedOrder: make(map[*Concept]uint),
		Concept_orderStaged: make(map[uint]*Concept),
		Concepts_reference:  make(map[*Concept]*Concept),

		Concern_stagedOrder: make(map[*Concern]uint),
		Concern_orderStaged: make(map[uint]*Concern),
		Concerns_reference:  make(map[*Concern]*Concern),

		ConcernCompositionShape_stagedOrder: make(map[*ConcernCompositionShape]uint),
		ConcernCompositionShape_orderStaged: make(map[uint]*ConcernCompositionShape),
		ConcernCompositionShapes_reference:  make(map[*ConcernCompositionShape]*ConcernCompositionShape),

		ConcernInputShape_stagedOrder: make(map[*ConcernInputShape]uint),
		ConcernInputShape_orderStaged: make(map[uint]*ConcernInputShape),
		ConcernInputShapes_reference:  make(map[*ConcernInputShape]*ConcernInputShape),

		ConcernOutputShape_stagedOrder: make(map[*ConcernOutputShape]uint),
		ConcernOutputShape_orderStaged: make(map[uint]*ConcernOutputShape),
		ConcernOutputShapes_reference:  make(map[*ConcernOutputShape]*ConcernOutputShape),

		ConcernShape_stagedOrder: make(map[*ConcernShape]uint),
		ConcernShape_orderStaged: make(map[uint]*ConcernShape),
		ConcernShapes_reference:  make(map[*ConcernShape]*ConcernShape),

		Deliverable_stagedOrder: make(map[*Deliverable]uint),
		Deliverable_orderStaged: make(map[uint]*Deliverable),
		Deliverables_reference:  make(map[*Deliverable]*Deliverable),

		Diagram_stagedOrder: make(map[*Diagram]uint),
		Diagram_orderStaged: make(map[uint]*Diagram),
		Diagrams_reference:  make(map[*Diagram]*Diagram),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		Note_stagedOrder: make(map[*Note]uint),
		Note_orderStaged: make(map[uint]*Note),
		Notes_reference:  make(map[*Note]*Note),

		NoteProductShape_stagedOrder: make(map[*NoteProductShape]uint),
		NoteProductShape_orderStaged: make(map[uint]*NoteProductShape),
		NoteProductShapes_reference:  make(map[*NoteProductShape]*NoteProductShape),

		NoteShape_stagedOrder: make(map[*NoteShape]uint),
		NoteShape_orderStaged: make(map[uint]*NoteShape),
		NoteShapes_reference:  make(map[*NoteShape]*NoteShape),

		NoteStakeholderShape_stagedOrder: make(map[*NoteStakeholderShape]uint),
		NoteStakeholderShape_orderStaged: make(map[uint]*NoteStakeholderShape),
		NoteStakeholderShapes_reference:  make(map[*NoteStakeholderShape]*NoteStakeholderShape),

		NoteTaskShape_stagedOrder: make(map[*NoteTaskShape]uint),
		NoteTaskShape_orderStaged: make(map[uint]*NoteTaskShape),
		NoteTaskShapes_reference:  make(map[*NoteTaskShape]*NoteTaskShape),

		ProductCompositionShape_stagedOrder: make(map[*ProductCompositionShape]uint),
		ProductCompositionShape_orderStaged: make(map[uint]*ProductCompositionShape),
		ProductCompositionShapes_reference:  make(map[*ProductCompositionShape]*ProductCompositionShape),

		ProductShape_stagedOrder: make(map[*ProductShape]uint),
		ProductShape_orderStaged: make(map[uint]*ProductShape),
		ProductShapes_reference:  make(map[*ProductShape]*ProductShape),

		Requirement_stagedOrder: make(map[*Requirement]uint),
		Requirement_orderStaged: make(map[uint]*Requirement),
		Requirements_reference:  make(map[*Requirement]*Requirement),

		Stakeholder_stagedOrder: make(map[*Stakeholder]uint),
		Stakeholder_orderStaged: make(map[uint]*Stakeholder),
		Stakeholders_reference:  make(map[*Stakeholder]*Stakeholder),

		StakeholderCompositionShape_stagedOrder: make(map[*StakeholderCompositionShape]uint),
		StakeholderCompositionShape_orderStaged: make(map[uint]*StakeholderCompositionShape),
		StakeholderCompositionShapes_reference:  make(map[*StakeholderCompositionShape]*StakeholderCompositionShape),

		StakeholderConcernShape_stagedOrder: make(map[*StakeholderConcernShape]uint),
		StakeholderConcernShape_orderStaged: make(map[uint]*StakeholderConcernShape),
		StakeholderConcernShapes_reference:  make(map[*StakeholderConcernShape]*StakeholderConcernShape),

		StakeholderShape_stagedOrder: make(map[*StakeholderShape]uint),
		StakeholderShape_orderStaged: make(map[uint]*StakeholderShape),
		StakeholderShapes_reference:  make(map[*StakeholderShape]*StakeholderShape),

		SupportLevel_stagedOrder: make(map[*SupportLevel]uint),
		SupportLevel_orderStaged: make(map[uint]*SupportLevel),
		SupportLevels_reference:  make(map[*SupportLevel]*SupportLevel),

		Tool_stagedOrder: make(map[*Tool]uint),
		Tool_orderStaged: make(map[uint]*Tool),
		Tools_reference:  make(map[*Tool]*Tool),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"AnalysisNeed": &AnalysisNeedUnmarshaller{},

			"Concept": &ConceptUnmarshaller{},

			"Concern": &ConcernUnmarshaller{},

			"ConcernCompositionShape": &ConcernCompositionShapeUnmarshaller{},

			"ConcernInputShape": &ConcernInputShapeUnmarshaller{},

			"ConcernOutputShape": &ConcernOutputShapeUnmarshaller{},

			"ConcernShape": &ConcernShapeUnmarshaller{},

			"Deliverable": &DeliverableUnmarshaller{},

			"Diagram": &DiagramUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteProductShape": &NoteProductShapeUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"NoteStakeholderShape": &NoteStakeholderShapeUnmarshaller{},

			"NoteTaskShape": &NoteTaskShapeUnmarshaller{},

			"ProductCompositionShape": &ProductCompositionShapeUnmarshaller{},

			"ProductShape": &ProductShapeUnmarshaller{},

			"Requirement": &RequirementUnmarshaller{},

			"Stakeholder": &StakeholderUnmarshaller{},

			"StakeholderCompositionShape": &StakeholderCompositionShapeUnmarshaller{},

			"StakeholderConcernShape": &StakeholderConcernShapeUnmarshaller{},

			"StakeholderShape": &StakeholderShapeUnmarshaller{},

			"SupportLevel": &SupportLevelUnmarshaller{},

			"Tool": &ToolUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AnalysisNeed"},
			{name: "Concept"},
			{name: "Concern"},
			{name: "ConcernCompositionShape"},
			{name: "ConcernInputShape"},
			{name: "ConcernOutputShape"},
			{name: "ConcernShape"},
			{name: "Deliverable"},
			{name: "Diagram"},
			{name: "Library"},
			{name: "Note"},
			{name: "NoteProductShape"},
			{name: "NoteShape"},
			{name: "NoteStakeholderShape"},
			{name: "NoteTaskShape"},
			{name: "ProductCompositionShape"},
			{name: "ProductShape"},
			{name: "Requirement"},
			{name: "Stakeholder"},
			{name: "StakeholderCompositionShape"},
			{name: "StakeholderConcernShape"},
			{name: "StakeholderShape"},
			{name: "SupportLevel"},
			{name: "Tool"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AnalysisNeed:
		return stage.AnalysisNeed_stagedOrder[instance]
	case *Concept:
		return stage.Concept_stagedOrder[instance]
	case *Concern:
		return stage.Concern_stagedOrder[instance]
	case *ConcernCompositionShape:
		return stage.ConcernCompositionShape_stagedOrder[instance]
	case *ConcernInputShape:
		return stage.ConcernInputShape_stagedOrder[instance]
	case *ConcernOutputShape:
		return stage.ConcernOutputShape_stagedOrder[instance]
	case *ConcernShape:
		return stage.ConcernShape_stagedOrder[instance]
	case *Deliverable:
		return stage.Deliverable_stagedOrder[instance]
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteProductShape:
		return stage.NoteProductShape_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *NoteStakeholderShape:
		return stage.NoteStakeholderShape_stagedOrder[instance]
	case *NoteTaskShape:
		return stage.NoteTaskShape_stagedOrder[instance]
	case *ProductCompositionShape:
		return stage.ProductCompositionShape_stagedOrder[instance]
	case *ProductShape:
		return stage.ProductShape_stagedOrder[instance]
	case *Requirement:
		return stage.Requirement_stagedOrder[instance]
	case *Stakeholder:
		return stage.Stakeholder_stagedOrder[instance]
	case *StakeholderCompositionShape:
		return stage.StakeholderCompositionShape_stagedOrder[instance]
	case *StakeholderConcernShape:
		return stage.StakeholderConcernShape_stagedOrder[instance]
	case *StakeholderShape:
		return stage.StakeholderShape_stagedOrder[instance]
	case *SupportLevel:
		return stage.SupportLevel_stagedOrder[instance]
	case *Tool:
		return stage.Tool_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *AnalysisNeed:
		return any(stage.AnalysisNeed_orderStaged[order]).(Type)
	case *Concept:
		return any(stage.Concept_orderStaged[order]).(Type)
	case *Concern:
		return any(stage.Concern_orderStaged[order]).(Type)
	case *ConcernCompositionShape:
		return any(stage.ConcernCompositionShape_orderStaged[order]).(Type)
	case *ConcernInputShape:
		return any(stage.ConcernInputShape_orderStaged[order]).(Type)
	case *ConcernOutputShape:
		return any(stage.ConcernOutputShape_orderStaged[order]).(Type)
	case *ConcernShape:
		return any(stage.ConcernShape_orderStaged[order]).(Type)
	case *Deliverable:
		return any(stage.Deliverable_orderStaged[order]).(Type)
	case *Diagram:
		return any(stage.Diagram_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Note:
		return any(stage.Note_orderStaged[order]).(Type)
	case *NoteProductShape:
		return any(stage.NoteProductShape_orderStaged[order]).(Type)
	case *NoteShape:
		return any(stage.NoteShape_orderStaged[order]).(Type)
	case *NoteStakeholderShape:
		return any(stage.NoteStakeholderShape_orderStaged[order]).(Type)
	case *NoteTaskShape:
		return any(stage.NoteTaskShape_orderStaged[order]).(Type)
	case *ProductCompositionShape:
		return any(stage.ProductCompositionShape_orderStaged[order]).(Type)
	case *ProductShape:
		return any(stage.ProductShape_orderStaged[order]).(Type)
	case *Requirement:
		return any(stage.Requirement_orderStaged[order]).(Type)
	case *Stakeholder:
		return any(stage.Stakeholder_orderStaged[order]).(Type)
	case *StakeholderCompositionShape:
		return any(stage.StakeholderCompositionShape_orderStaged[order]).(Type)
	case *StakeholderConcernShape:
		return any(stage.StakeholderConcernShape_orderStaged[order]).(Type)
	case *StakeholderShape:
		return any(stage.StakeholderShape_orderStaged[order]).(Type)
	case *SupportLevel:
		return any(stage.SupportLevel_orderStaged[order]).(Type)
	case *Tool:
		return any(stage.Tool_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AnalysisNeed:
		return stage.AnalysisNeed_stagedOrder[instance]
	case *Concept:
		return stage.Concept_stagedOrder[instance]
	case *Concern:
		return stage.Concern_stagedOrder[instance]
	case *ConcernCompositionShape:
		return stage.ConcernCompositionShape_stagedOrder[instance]
	case *ConcernInputShape:
		return stage.ConcernInputShape_stagedOrder[instance]
	case *ConcernOutputShape:
		return stage.ConcernOutputShape_stagedOrder[instance]
	case *ConcernShape:
		return stage.ConcernShape_stagedOrder[instance]
	case *Deliverable:
		return stage.Deliverable_stagedOrder[instance]
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteProductShape:
		return stage.NoteProductShape_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *NoteStakeholderShape:
		return stage.NoteStakeholderShape_stagedOrder[instance]
	case *NoteTaskShape:
		return stage.NoteTaskShape_stagedOrder[instance]
	case *ProductCompositionShape:
		return stage.ProductCompositionShape_stagedOrder[instance]
	case *ProductShape:
		return stage.ProductShape_stagedOrder[instance]
	case *Requirement:
		return stage.Requirement_stagedOrder[instance]
	case *Stakeholder:
		return stage.Stakeholder_stagedOrder[instance]
	case *StakeholderCompositionShape:
		return stage.StakeholderCompositionShape_stagedOrder[instance]
	case *StakeholderConcernShape:
		return stage.StakeholderConcernShape_stagedOrder[instance]
	case *StakeholderShape:
		return stage.StakeholderShape_stagedOrder[instance]
	case *SupportLevel:
		return stage.SupportLevel_stagedOrder[instance]
	case *Tool:
		return stage.Tool_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {
	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	tmp2 := stage.beforeCommitHooks
	stage.beforeCommitHooks = nil
	tmp3 := stage.afterCommitHooks
	stage.afterCommitHooks = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.beforeCommitHooks = tmp2
	stage.afterCommitHooks = tmp3
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.isInDeltaMode && stage.navigationMode == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
		stage.ResetHard()
	}

	if stage.IsInDeltaMode() {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
		if stage.probeIF != nil {
			stage.probeIF.RefreshNavigationTree()
		}
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AnalysisNeed"] = len(stage.AnalysisNeeds)
	stage.Map_GongStructName_InstancesNb["Concept"] = len(stage.Concepts)
	stage.Map_GongStructName_InstancesNb["Concern"] = len(stage.Concerns)
	stage.Map_GongStructName_InstancesNb["ConcernCompositionShape"] = len(stage.ConcernCompositionShapes)
	stage.Map_GongStructName_InstancesNb["ConcernInputShape"] = len(stage.ConcernInputShapes)
	stage.Map_GongStructName_InstancesNb["ConcernOutputShape"] = len(stage.ConcernOutputShapes)
	stage.Map_GongStructName_InstancesNb["ConcernShape"] = len(stage.ConcernShapes)
	stage.Map_GongStructName_InstancesNb["Deliverable"] = len(stage.Deliverables)
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteProductShape"] = len(stage.NoteProductShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteStakeholderShape"] = len(stage.NoteStakeholderShapes)
	stage.Map_GongStructName_InstancesNb["NoteTaskShape"] = len(stage.NoteTaskShapes)
	stage.Map_GongStructName_InstancesNb["ProductCompositionShape"] = len(stage.ProductCompositionShapes)
	stage.Map_GongStructName_InstancesNb["ProductShape"] = len(stage.ProductShapes)
	stage.Map_GongStructName_InstancesNb["Requirement"] = len(stage.Requirements)
	stage.Map_GongStructName_InstancesNb["Stakeholder"] = len(stage.Stakeholders)
	stage.Map_GongStructName_InstancesNb["StakeholderCompositionShape"] = len(stage.StakeholderCompositionShapes)
	stage.Map_GongStructName_InstancesNb["StakeholderConcernShape"] = len(stage.StakeholderConcernShapes)
	stage.Map_GongStructName_InstancesNb["StakeholderShape"] = len(stage.StakeholderShapes)
	stage.Map_GongStructName_InstancesNb["SupportLevel"] = len(stage.SupportLevels)
	stage.Map_GongStructName_InstancesNb["Tool"] = len(stage.Tools)
}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts analysisneed to the model stage
func (analysisneed *AnalysisNeed) Stage(stage *Stage) *AnalysisNeed {
	if _, ok := stage.AnalysisNeeds[analysisneed]; !ok {
		stage.AnalysisNeeds[analysisneed] = struct{}{}
		stage.AnalysisNeed_stagedOrder[analysisneed] = stage.AnalysisNeedOrder
		stage.AnalysisNeed_orderStaged[stage.AnalysisNeedOrder] = analysisneed
		stage.AnalysisNeedOrder++
	}
	stage.AnalysisNeeds_mapString[analysisneed.Name] = analysisneed

	return analysisneed
}

// StagePreserveOrder puts analysisneed to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AnalysisNeedOrder
// - update stage.AnalysisNeedOrder accordingly
func (analysisneed *AnalysisNeed) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AnalysisNeeds[analysisneed]; !ok {
		stage.AnalysisNeeds[analysisneed] = struct{}{}

		if order > stage.AnalysisNeedOrder {
			stage.AnalysisNeedOrder = order
		}
		stage.AnalysisNeed_stagedOrder[analysisneed] = order
		stage.AnalysisNeed_orderStaged[order] = analysisneed
		stage.AnalysisNeedOrder++
	}
	stage.AnalysisNeeds_mapString[analysisneed.Name] = analysisneed
}

// Unstage removes analysisneed off the model stage
func (analysisneed *AnalysisNeed) Unstage(stage *Stage) *AnalysisNeed {
	delete(stage.AnalysisNeeds, analysisneed)
	// issue1150
	// delete(stage.AnalysisNeed_stagedOrder, analysisneed)
	delete(stage.AnalysisNeeds_mapString, analysisneed.Name)

	return analysisneed
}

// UnstageVoid removes analysisneed off the model stage
func (analysisneed *AnalysisNeed) UnstageVoid(stage *Stage) {
	delete(stage.AnalysisNeeds, analysisneed)
	// issue1150
	// delete(stage.AnalysisNeed_stagedOrder, analysisneed)
	delete(stage.AnalysisNeeds_mapString, analysisneed.Name)
}

// commit analysisneed to the back repo (if it is already staged)
func (analysisneed *AnalysisNeed) Commit(stage *Stage) *AnalysisNeed {
	if _, ok := stage.AnalysisNeeds[analysisneed]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAnalysisNeed(analysisneed)
		}
	}
	return analysisneed
}

func (analysisneed *AnalysisNeed) CommitVoid(stage *Stage) {
	analysisneed.Commit(stage)
}

func (analysisneed *AnalysisNeed) StageVoid(stage *Stage) {
	analysisneed.Stage(stage)
}

// Checkout analysisneed to the back repo (if it is already staged)
func (analysisneed *AnalysisNeed) Checkout(stage *Stage) *AnalysisNeed {
	if _, ok := stage.AnalysisNeeds[analysisneed]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAnalysisNeed(analysisneed)
		}
	}
	return analysisneed
}

// for satisfaction of GongStruct interface
func (analysisneed *AnalysisNeed) GetName() (res string) {
	return analysisneed.Name
}

// for satisfaction of GongStruct interface
func (analysisneed *AnalysisNeed) SetName(name string) {
	analysisneed.Name = name
}

// Stage puts concept to the model stage
func (concept *Concept) Stage(stage *Stage) *Concept {
	if _, ok := stage.Concepts[concept]; !ok {
		stage.Concepts[concept] = struct{}{}
		stage.Concept_stagedOrder[concept] = stage.ConceptOrder
		stage.Concept_orderStaged[stage.ConceptOrder] = concept
		stage.ConceptOrder++
	}
	stage.Concepts_mapString[concept.Name] = concept

	return concept
}

// StagePreserveOrder puts concept to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConceptOrder
// - update stage.ConceptOrder accordingly
func (concept *Concept) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Concepts[concept]; !ok {
		stage.Concepts[concept] = struct{}{}

		if order > stage.ConceptOrder {
			stage.ConceptOrder = order
		}
		stage.Concept_stagedOrder[concept] = order
		stage.Concept_orderStaged[order] = concept
		stage.ConceptOrder++
	}
	stage.Concepts_mapString[concept.Name] = concept
}

// Unstage removes concept off the model stage
func (concept *Concept) Unstage(stage *Stage) *Concept {
	delete(stage.Concepts, concept)
	// issue1150
	// delete(stage.Concept_stagedOrder, concept)
	delete(stage.Concepts_mapString, concept.Name)

	return concept
}

// UnstageVoid removes concept off the model stage
func (concept *Concept) UnstageVoid(stage *Stage) {
	delete(stage.Concepts, concept)
	// issue1150
	// delete(stage.Concept_stagedOrder, concept)
	delete(stage.Concepts_mapString, concept.Name)
}

// commit concept to the back repo (if it is already staged)
func (concept *Concept) Commit(stage *Stage) *Concept {
	if _, ok := stage.Concepts[concept]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitConcept(concept)
		}
	}
	return concept
}

func (concept *Concept) CommitVoid(stage *Stage) {
	concept.Commit(stage)
}

func (concept *Concept) StageVoid(stage *Stage) {
	concept.Stage(stage)
}

// Checkout concept to the back repo (if it is already staged)
func (concept *Concept) Checkout(stage *Stage) *Concept {
	if _, ok := stage.Concepts[concept]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutConcept(concept)
		}
	}
	return concept
}

// for satisfaction of GongStruct interface
func (concept *Concept) GetName() (res string) {
	return concept.Name
}

// for satisfaction of GongStruct interface
func (concept *Concept) SetName(name string) {
	concept.Name = name
}

// Stage puts concern to the model stage
func (concern *Concern) Stage(stage *Stage) *Concern {
	if _, ok := stage.Concerns[concern]; !ok {
		stage.Concerns[concern] = struct{}{}
		stage.Concern_stagedOrder[concern] = stage.ConcernOrder
		stage.Concern_orderStaged[stage.ConcernOrder] = concern
		stage.ConcernOrder++
	}
	stage.Concerns_mapString[concern.Name] = concern

	return concern
}

// StagePreserveOrder puts concern to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernOrder
// - update stage.ConcernOrder accordingly
func (concern *Concern) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Concerns[concern]; !ok {
		stage.Concerns[concern] = struct{}{}

		if order > stage.ConcernOrder {
			stage.ConcernOrder = order
		}
		stage.Concern_stagedOrder[concern] = order
		stage.Concern_orderStaged[order] = concern
		stage.ConcernOrder++
	}
	stage.Concerns_mapString[concern.Name] = concern
}

// Unstage removes concern off the model stage
func (concern *Concern) Unstage(stage *Stage) *Concern {
	delete(stage.Concerns, concern)
	// issue1150
	// delete(stage.Concern_stagedOrder, concern)
	delete(stage.Concerns_mapString, concern.Name)

	return concern
}

// UnstageVoid removes concern off the model stage
func (concern *Concern) UnstageVoid(stage *Stage) {
	delete(stage.Concerns, concern)
	// issue1150
	// delete(stage.Concern_stagedOrder, concern)
	delete(stage.Concerns_mapString, concern.Name)
}

// commit concern to the back repo (if it is already staged)
func (concern *Concern) Commit(stage *Stage) *Concern {
	if _, ok := stage.Concerns[concern]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitConcern(concern)
		}
	}
	return concern
}

func (concern *Concern) CommitVoid(stage *Stage) {
	concern.Commit(stage)
}

func (concern *Concern) StageVoid(stage *Stage) {
	concern.Stage(stage)
}

// Checkout concern to the back repo (if it is already staged)
func (concern *Concern) Checkout(stage *Stage) *Concern {
	if _, ok := stage.Concerns[concern]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutConcern(concern)
		}
	}
	return concern
}

// for satisfaction of GongStruct interface
func (concern *Concern) GetName() (res string) {
	return concern.Name
}

// for satisfaction of GongStruct interface
func (concern *Concern) SetName(name string) {
	concern.Name = name
}

// Stage puts concerncompositionshape to the model stage
func (concerncompositionshape *ConcernCompositionShape) Stage(stage *Stage) *ConcernCompositionShape {
	if _, ok := stage.ConcernCompositionShapes[concerncompositionshape]; !ok {
		stage.ConcernCompositionShapes[concerncompositionshape] = struct{}{}
		stage.ConcernCompositionShape_stagedOrder[concerncompositionshape] = stage.ConcernCompositionShapeOrder
		stage.ConcernCompositionShape_orderStaged[stage.ConcernCompositionShapeOrder] = concerncompositionshape
		stage.ConcernCompositionShapeOrder++
	}
	stage.ConcernCompositionShapes_mapString[concerncompositionshape.Name] = concerncompositionshape

	return concerncompositionshape
}

// StagePreserveOrder puts concerncompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernCompositionShapeOrder
// - update stage.ConcernCompositionShapeOrder accordingly
func (concerncompositionshape *ConcernCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ConcernCompositionShapes[concerncompositionshape]; !ok {
		stage.ConcernCompositionShapes[concerncompositionshape] = struct{}{}

		if order > stage.ConcernCompositionShapeOrder {
			stage.ConcernCompositionShapeOrder = order
		}
		stage.ConcernCompositionShape_stagedOrder[concerncompositionshape] = order
		stage.ConcernCompositionShape_orderStaged[order] = concerncompositionshape
		stage.ConcernCompositionShapeOrder++
	}
	stage.ConcernCompositionShapes_mapString[concerncompositionshape.Name] = concerncompositionshape
}

// Unstage removes concerncompositionshape off the model stage
func (concerncompositionshape *ConcernCompositionShape) Unstage(stage *Stage) *ConcernCompositionShape {
	delete(stage.ConcernCompositionShapes, concerncompositionshape)
	// issue1150
	// delete(stage.ConcernCompositionShape_stagedOrder, concerncompositionshape)
	delete(stage.ConcernCompositionShapes_mapString, concerncompositionshape.Name)

	return concerncompositionshape
}

// UnstageVoid removes concerncompositionshape off the model stage
func (concerncompositionshape *ConcernCompositionShape) UnstageVoid(stage *Stage) {
	delete(stage.ConcernCompositionShapes, concerncompositionshape)
	// issue1150
	// delete(stage.ConcernCompositionShape_stagedOrder, concerncompositionshape)
	delete(stage.ConcernCompositionShapes_mapString, concerncompositionshape.Name)
}

// commit concerncompositionshape to the back repo (if it is already staged)
func (concerncompositionshape *ConcernCompositionShape) Commit(stage *Stage) *ConcernCompositionShape {
	if _, ok := stage.ConcernCompositionShapes[concerncompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitConcernCompositionShape(concerncompositionshape)
		}
	}
	return concerncompositionshape
}

func (concerncompositionshape *ConcernCompositionShape) CommitVoid(stage *Stage) {
	concerncompositionshape.Commit(stage)
}

func (concerncompositionshape *ConcernCompositionShape) StageVoid(stage *Stage) {
	concerncompositionshape.Stage(stage)
}

// Checkout concerncompositionshape to the back repo (if it is already staged)
func (concerncompositionshape *ConcernCompositionShape) Checkout(stage *Stage) *ConcernCompositionShape {
	if _, ok := stage.ConcernCompositionShapes[concerncompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutConcernCompositionShape(concerncompositionshape)
		}
	}
	return concerncompositionshape
}

// for satisfaction of GongStruct interface
func (concerncompositionshape *ConcernCompositionShape) GetName() (res string) {
	return concerncompositionshape.Name
}

// for satisfaction of GongStruct interface
func (concerncompositionshape *ConcernCompositionShape) SetName(name string) {
	concerncompositionshape.Name = name
}

// Stage puts concerninputshape to the model stage
func (concerninputshape *ConcernInputShape) Stage(stage *Stage) *ConcernInputShape {
	if _, ok := stage.ConcernInputShapes[concerninputshape]; !ok {
		stage.ConcernInputShapes[concerninputshape] = struct{}{}
		stage.ConcernInputShape_stagedOrder[concerninputshape] = stage.ConcernInputShapeOrder
		stage.ConcernInputShape_orderStaged[stage.ConcernInputShapeOrder] = concerninputshape
		stage.ConcernInputShapeOrder++
	}
	stage.ConcernInputShapes_mapString[concerninputshape.Name] = concerninputshape

	return concerninputshape
}

// StagePreserveOrder puts concerninputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernInputShapeOrder
// - update stage.ConcernInputShapeOrder accordingly
func (concerninputshape *ConcernInputShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ConcernInputShapes[concerninputshape]; !ok {
		stage.ConcernInputShapes[concerninputshape] = struct{}{}

		if order > stage.ConcernInputShapeOrder {
			stage.ConcernInputShapeOrder = order
		}
		stage.ConcernInputShape_stagedOrder[concerninputshape] = order
		stage.ConcernInputShape_orderStaged[order] = concerninputshape
		stage.ConcernInputShapeOrder++
	}
	stage.ConcernInputShapes_mapString[concerninputshape.Name] = concerninputshape
}

// Unstage removes concerninputshape off the model stage
func (concerninputshape *ConcernInputShape) Unstage(stage *Stage) *ConcernInputShape {
	delete(stage.ConcernInputShapes, concerninputshape)
	// issue1150
	// delete(stage.ConcernInputShape_stagedOrder, concerninputshape)
	delete(stage.ConcernInputShapes_mapString, concerninputshape.Name)

	return concerninputshape
}

// UnstageVoid removes concerninputshape off the model stage
func (concerninputshape *ConcernInputShape) UnstageVoid(stage *Stage) {
	delete(stage.ConcernInputShapes, concerninputshape)
	// issue1150
	// delete(stage.ConcernInputShape_stagedOrder, concerninputshape)
	delete(stage.ConcernInputShapes_mapString, concerninputshape.Name)
}

// commit concerninputshape to the back repo (if it is already staged)
func (concerninputshape *ConcernInputShape) Commit(stage *Stage) *ConcernInputShape {
	if _, ok := stage.ConcernInputShapes[concerninputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitConcernInputShape(concerninputshape)
		}
	}
	return concerninputshape
}

func (concerninputshape *ConcernInputShape) CommitVoid(stage *Stage) {
	concerninputshape.Commit(stage)
}

func (concerninputshape *ConcernInputShape) StageVoid(stage *Stage) {
	concerninputshape.Stage(stage)
}

// Checkout concerninputshape to the back repo (if it is already staged)
func (concerninputshape *ConcernInputShape) Checkout(stage *Stage) *ConcernInputShape {
	if _, ok := stage.ConcernInputShapes[concerninputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutConcernInputShape(concerninputshape)
		}
	}
	return concerninputshape
}

// for satisfaction of GongStruct interface
func (concerninputshape *ConcernInputShape) GetName() (res string) {
	return concerninputshape.Name
}

// for satisfaction of GongStruct interface
func (concerninputshape *ConcernInputShape) SetName(name string) {
	concerninputshape.Name = name
}

// Stage puts concernoutputshape to the model stage
func (concernoutputshape *ConcernOutputShape) Stage(stage *Stage) *ConcernOutputShape {
	if _, ok := stage.ConcernOutputShapes[concernoutputshape]; !ok {
		stage.ConcernOutputShapes[concernoutputshape] = struct{}{}
		stage.ConcernOutputShape_stagedOrder[concernoutputshape] = stage.ConcernOutputShapeOrder
		stage.ConcernOutputShape_orderStaged[stage.ConcernOutputShapeOrder] = concernoutputshape
		stage.ConcernOutputShapeOrder++
	}
	stage.ConcernOutputShapes_mapString[concernoutputshape.Name] = concernoutputshape

	return concernoutputshape
}

// StagePreserveOrder puts concernoutputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernOutputShapeOrder
// - update stage.ConcernOutputShapeOrder accordingly
func (concernoutputshape *ConcernOutputShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ConcernOutputShapes[concernoutputshape]; !ok {
		stage.ConcernOutputShapes[concernoutputshape] = struct{}{}

		if order > stage.ConcernOutputShapeOrder {
			stage.ConcernOutputShapeOrder = order
		}
		stage.ConcernOutputShape_stagedOrder[concernoutputshape] = order
		stage.ConcernOutputShape_orderStaged[order] = concernoutputshape
		stage.ConcernOutputShapeOrder++
	}
	stage.ConcernOutputShapes_mapString[concernoutputshape.Name] = concernoutputshape
}

// Unstage removes concernoutputshape off the model stage
func (concernoutputshape *ConcernOutputShape) Unstage(stage *Stage) *ConcernOutputShape {
	delete(stage.ConcernOutputShapes, concernoutputshape)
	// issue1150
	// delete(stage.ConcernOutputShape_stagedOrder, concernoutputshape)
	delete(stage.ConcernOutputShapes_mapString, concernoutputshape.Name)

	return concernoutputshape
}

// UnstageVoid removes concernoutputshape off the model stage
func (concernoutputshape *ConcernOutputShape) UnstageVoid(stage *Stage) {
	delete(stage.ConcernOutputShapes, concernoutputshape)
	// issue1150
	// delete(stage.ConcernOutputShape_stagedOrder, concernoutputshape)
	delete(stage.ConcernOutputShapes_mapString, concernoutputshape.Name)
}

// commit concernoutputshape to the back repo (if it is already staged)
func (concernoutputshape *ConcernOutputShape) Commit(stage *Stage) *ConcernOutputShape {
	if _, ok := stage.ConcernOutputShapes[concernoutputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitConcernOutputShape(concernoutputshape)
		}
	}
	return concernoutputshape
}

func (concernoutputshape *ConcernOutputShape) CommitVoid(stage *Stage) {
	concernoutputshape.Commit(stage)
}

func (concernoutputshape *ConcernOutputShape) StageVoid(stage *Stage) {
	concernoutputshape.Stage(stage)
}

// Checkout concernoutputshape to the back repo (if it is already staged)
func (concernoutputshape *ConcernOutputShape) Checkout(stage *Stage) *ConcernOutputShape {
	if _, ok := stage.ConcernOutputShapes[concernoutputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutConcernOutputShape(concernoutputshape)
		}
	}
	return concernoutputshape
}

// for satisfaction of GongStruct interface
func (concernoutputshape *ConcernOutputShape) GetName() (res string) {
	return concernoutputshape.Name
}

// for satisfaction of GongStruct interface
func (concernoutputshape *ConcernOutputShape) SetName(name string) {
	concernoutputshape.Name = name
}

// Stage puts concernshape to the model stage
func (concernshape *ConcernShape) Stage(stage *Stage) *ConcernShape {
	if _, ok := stage.ConcernShapes[concernshape]; !ok {
		stage.ConcernShapes[concernshape] = struct{}{}
		stage.ConcernShape_stagedOrder[concernshape] = stage.ConcernShapeOrder
		stage.ConcernShape_orderStaged[stage.ConcernShapeOrder] = concernshape
		stage.ConcernShapeOrder++
	}
	stage.ConcernShapes_mapString[concernshape.Name] = concernshape

	return concernshape
}

// StagePreserveOrder puts concernshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernShapeOrder
// - update stage.ConcernShapeOrder accordingly
func (concernshape *ConcernShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ConcernShapes[concernshape]; !ok {
		stage.ConcernShapes[concernshape] = struct{}{}

		if order > stage.ConcernShapeOrder {
			stage.ConcernShapeOrder = order
		}
		stage.ConcernShape_stagedOrder[concernshape] = order
		stage.ConcernShape_orderStaged[order] = concernshape
		stage.ConcernShapeOrder++
	}
	stage.ConcernShapes_mapString[concernshape.Name] = concernshape
}

// Unstage removes concernshape off the model stage
func (concernshape *ConcernShape) Unstage(stage *Stage) *ConcernShape {
	delete(stage.ConcernShapes, concernshape)
	// issue1150
	// delete(stage.ConcernShape_stagedOrder, concernshape)
	delete(stage.ConcernShapes_mapString, concernshape.Name)

	return concernshape
}

// UnstageVoid removes concernshape off the model stage
func (concernshape *ConcernShape) UnstageVoid(stage *Stage) {
	delete(stage.ConcernShapes, concernshape)
	// issue1150
	// delete(stage.ConcernShape_stagedOrder, concernshape)
	delete(stage.ConcernShapes_mapString, concernshape.Name)
}

// commit concernshape to the back repo (if it is already staged)
func (concernshape *ConcernShape) Commit(stage *Stage) *ConcernShape {
	if _, ok := stage.ConcernShapes[concernshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitConcernShape(concernshape)
		}
	}
	return concernshape
}

func (concernshape *ConcernShape) CommitVoid(stage *Stage) {
	concernshape.Commit(stage)
}

func (concernshape *ConcernShape) StageVoid(stage *Stage) {
	concernshape.Stage(stage)
}

// Checkout concernshape to the back repo (if it is already staged)
func (concernshape *ConcernShape) Checkout(stage *Stage) *ConcernShape {
	if _, ok := stage.ConcernShapes[concernshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutConcernShape(concernshape)
		}
	}
	return concernshape
}

// for satisfaction of GongStruct interface
func (concernshape *ConcernShape) GetName() (res string) {
	return concernshape.Name
}

// for satisfaction of GongStruct interface
func (concernshape *ConcernShape) SetName(name string) {
	concernshape.Name = name
}

// Stage puts deliverable to the model stage
func (deliverable *Deliverable) Stage(stage *Stage) *Deliverable {
	if _, ok := stage.Deliverables[deliverable]; !ok {
		stage.Deliverables[deliverable] = struct{}{}
		stage.Deliverable_stagedOrder[deliverable] = stage.DeliverableOrder
		stage.Deliverable_orderStaged[stage.DeliverableOrder] = deliverable
		stage.DeliverableOrder++
	}
	stage.Deliverables_mapString[deliverable.Name] = deliverable

	return deliverable
}

// StagePreserveOrder puts deliverable to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DeliverableOrder
// - update stage.DeliverableOrder accordingly
func (deliverable *Deliverable) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Deliverables[deliverable]; !ok {
		stage.Deliverables[deliverable] = struct{}{}

		if order > stage.DeliverableOrder {
			stage.DeliverableOrder = order
		}
		stage.Deliverable_stagedOrder[deliverable] = order
		stage.Deliverable_orderStaged[order] = deliverable
		stage.DeliverableOrder++
	}
	stage.Deliverables_mapString[deliverable.Name] = deliverable
}

// Unstage removes deliverable off the model stage
func (deliverable *Deliverable) Unstage(stage *Stage) *Deliverable {
	delete(stage.Deliverables, deliverable)
	// issue1150
	// delete(stage.Deliverable_stagedOrder, deliverable)
	delete(stage.Deliverables_mapString, deliverable.Name)

	return deliverable
}

// UnstageVoid removes deliverable off the model stage
func (deliverable *Deliverable) UnstageVoid(stage *Stage) {
	delete(stage.Deliverables, deliverable)
	// issue1150
	// delete(stage.Deliverable_stagedOrder, deliverable)
	delete(stage.Deliverables_mapString, deliverable.Name)
}

// commit deliverable to the back repo (if it is already staged)
func (deliverable *Deliverable) Commit(stage *Stage) *Deliverable {
	if _, ok := stage.Deliverables[deliverable]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDeliverable(deliverable)
		}
	}
	return deliverable
}

func (deliverable *Deliverable) CommitVoid(stage *Stage) {
	deliverable.Commit(stage)
}

func (deliverable *Deliverable) StageVoid(stage *Stage) {
	deliverable.Stage(stage)
}

// Checkout deliverable to the back repo (if it is already staged)
func (deliverable *Deliverable) Checkout(stage *Stage) *Deliverable {
	if _, ok := stage.Deliverables[deliverable]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDeliverable(deliverable)
		}
	}
	return deliverable
}

// for satisfaction of GongStruct interface
func (deliverable *Deliverable) GetName() (res string) {
	return deliverable.Name
}

// for satisfaction of GongStruct interface
func (deliverable *Deliverable) SetName(name string) {
	deliverable.Name = name
}

// Stage puts diagram to the model stage
func (diagram *Diagram) Stage(stage *Stage) *Diagram {
	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}
		stage.Diagram_stagedOrder[diagram] = stage.DiagramOrder
		stage.Diagram_orderStaged[stage.DiagramOrder] = diagram
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram

	return diagram
}

// StagePreserveOrder puts diagram to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramOrder
// - update stage.DiagramOrder accordingly
func (diagram *Diagram) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}

		if order > stage.DiagramOrder {
			stage.DiagramOrder = order
		}
		stage.Diagram_stagedOrder[diagram] = order
		stage.Diagram_orderStaged[order] = diagram
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram
}

// Unstage removes diagram off the model stage
func (diagram *Diagram) Unstage(stage *Stage) *Diagram {
	delete(stage.Diagrams, diagram)
	// issue1150
	// delete(stage.Diagram_stagedOrder, diagram)
	delete(stage.Diagrams_mapString, diagram.Name)

	return diagram
}

// UnstageVoid removes diagram off the model stage
func (diagram *Diagram) UnstageVoid(stage *Stage) {
	delete(stage.Diagrams, diagram)
	// issue1150
	// delete(stage.Diagram_stagedOrder, diagram)
	delete(stage.Diagrams_mapString, diagram.Name)
}

// commit diagram to the back repo (if it is already staged)
func (diagram *Diagram) Commit(stage *Stage) *Diagram {
	if _, ok := stage.Diagrams[diagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagram(diagram)
		}
	}
	return diagram
}

func (diagram *Diagram) CommitVoid(stage *Stage) {
	diagram.Commit(stage)
}

func (diagram *Diagram) StageVoid(stage *Stage) {
	diagram.Stage(stage)
}

// Checkout diagram to the back repo (if it is already staged)
func (diagram *Diagram) Checkout(stage *Stage) *Diagram {
	if _, ok := stage.Diagrams[diagram]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagram(diagram)
		}
	}
	return diagram
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) GetName() (res string) {
	return diagram.Name
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) SetName(name string) {
	diagram.Name = name
}

// Stage puts library to the model stage
func (library *Library) Stage(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; !ok {
		stage.Librarys[library] = struct{}{}
		stage.Library_stagedOrder[library] = stage.LibraryOrder
		stage.Library_orderStaged[stage.LibraryOrder] = library
		stage.LibraryOrder++
	}
	stage.Librarys_mapString[library.Name] = library

	return library
}

// StagePreserveOrder puts library to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LibraryOrder
// - update stage.LibraryOrder accordingly
func (library *Library) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Librarys[library]; !ok {
		stage.Librarys[library] = struct{}{}

		if order > stage.LibraryOrder {
			stage.LibraryOrder = order
		}
		stage.Library_stagedOrder[library] = order
		stage.Library_orderStaged[order] = library
		stage.LibraryOrder++
	}
	stage.Librarys_mapString[library.Name] = library
}

// Unstage removes library off the model stage
func (library *Library) Unstage(stage *Stage) *Library {
	delete(stage.Librarys, library)
	// issue1150
	// delete(stage.Library_stagedOrder, library)
	delete(stage.Librarys_mapString, library.Name)

	return library
}

// UnstageVoid removes library off the model stage
func (library *Library) UnstageVoid(stage *Stage) {
	delete(stage.Librarys, library)
	// issue1150
	// delete(stage.Library_stagedOrder, library)
	delete(stage.Librarys_mapString, library.Name)
}

// commit library to the back repo (if it is already staged)
func (library *Library) Commit(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLibrary(library)
		}
	}
	return library
}

func (library *Library) CommitVoid(stage *Stage) {
	library.Commit(stage)
}

func (library *Library) StageVoid(stage *Stage) {
	library.Stage(stage)
}

// Checkout library to the back repo (if it is already staged)
func (library *Library) Checkout(stage *Stage) *Library {
	if _, ok := stage.Librarys[library]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLibrary(library)
		}
	}
	return library
}

// for satisfaction of GongStruct interface
func (library *Library) GetName() (res string) {
	return library.Name
}

// for satisfaction of GongStruct interface
func (library *Library) SetName(name string) {
	library.Name = name
}

// Stage puts note to the model stage
func (note *Note) Stage(stage *Stage) *Note {
	if _, ok := stage.Notes[note]; !ok {
		stage.Notes[note] = struct{}{}
		stage.Note_stagedOrder[note] = stage.NoteOrder
		stage.Note_orderStaged[stage.NoteOrder] = note
		stage.NoteOrder++
	}
	stage.Notes_mapString[note.Name] = note

	return note
}

// StagePreserveOrder puts note to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteOrder
// - update stage.NoteOrder accordingly
func (note *Note) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Notes[note]; !ok {
		stage.Notes[note] = struct{}{}

		if order > stage.NoteOrder {
			stage.NoteOrder = order
		}
		stage.Note_stagedOrder[note] = order
		stage.Note_orderStaged[order] = note
		stage.NoteOrder++
	}
	stage.Notes_mapString[note.Name] = note
}

// Unstage removes note off the model stage
func (note *Note) Unstage(stage *Stage) *Note {
	delete(stage.Notes, note)
	// issue1150
	// delete(stage.Note_stagedOrder, note)
	delete(stage.Notes_mapString, note.Name)

	return note
}

// UnstageVoid removes note off the model stage
func (note *Note) UnstageVoid(stage *Stage) {
	delete(stage.Notes, note)
	// issue1150
	// delete(stage.Note_stagedOrder, note)
	delete(stage.Notes_mapString, note.Name)
}

// commit note to the back repo (if it is already staged)
func (note *Note) Commit(stage *Stage) *Note {
	if _, ok := stage.Notes[note]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNote(note)
		}
	}
	return note
}

func (note *Note) CommitVoid(stage *Stage) {
	note.Commit(stage)
}

func (note *Note) StageVoid(stage *Stage) {
	note.Stage(stage)
}

// Checkout note to the back repo (if it is already staged)
func (note *Note) Checkout(stage *Stage) *Note {
	if _, ok := stage.Notes[note]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNote(note)
		}
	}
	return note
}

// for satisfaction of GongStruct interface
func (note *Note) GetName() (res string) {
	return note.Name
}

// for satisfaction of GongStruct interface
func (note *Note) SetName(name string) {
	note.Name = name
}

// Stage puts noteproductshape to the model stage
func (noteproductshape *NoteProductShape) Stage(stage *Stage) *NoteProductShape {
	if _, ok := stage.NoteProductShapes[noteproductshape]; !ok {
		stage.NoteProductShapes[noteproductshape] = struct{}{}
		stage.NoteProductShape_stagedOrder[noteproductshape] = stage.NoteProductShapeOrder
		stage.NoteProductShape_orderStaged[stage.NoteProductShapeOrder] = noteproductshape
		stage.NoteProductShapeOrder++
	}
	stage.NoteProductShapes_mapString[noteproductshape.Name] = noteproductshape

	return noteproductshape
}

// StagePreserveOrder puts noteproductshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteProductShapeOrder
// - update stage.NoteProductShapeOrder accordingly
func (noteproductshape *NoteProductShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteProductShapes[noteproductshape]; !ok {
		stage.NoteProductShapes[noteproductshape] = struct{}{}

		if order > stage.NoteProductShapeOrder {
			stage.NoteProductShapeOrder = order
		}
		stage.NoteProductShape_stagedOrder[noteproductshape] = order
		stage.NoteProductShape_orderStaged[order] = noteproductshape
		stage.NoteProductShapeOrder++
	}
	stage.NoteProductShapes_mapString[noteproductshape.Name] = noteproductshape
}

// Unstage removes noteproductshape off the model stage
func (noteproductshape *NoteProductShape) Unstage(stage *Stage) *NoteProductShape {
	delete(stage.NoteProductShapes, noteproductshape)
	// issue1150
	// delete(stage.NoteProductShape_stagedOrder, noteproductshape)
	delete(stage.NoteProductShapes_mapString, noteproductshape.Name)

	return noteproductshape
}

// UnstageVoid removes noteproductshape off the model stage
func (noteproductshape *NoteProductShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteProductShapes, noteproductshape)
	// issue1150
	// delete(stage.NoteProductShape_stagedOrder, noteproductshape)
	delete(stage.NoteProductShapes_mapString, noteproductshape.Name)
}

// commit noteproductshape to the back repo (if it is already staged)
func (noteproductshape *NoteProductShape) Commit(stage *Stage) *NoteProductShape {
	if _, ok := stage.NoteProductShapes[noteproductshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteProductShape(noteproductshape)
		}
	}
	return noteproductshape
}

func (noteproductshape *NoteProductShape) CommitVoid(stage *Stage) {
	noteproductshape.Commit(stage)
}

func (noteproductshape *NoteProductShape) StageVoid(stage *Stage) {
	noteproductshape.Stage(stage)
}

// Checkout noteproductshape to the back repo (if it is already staged)
func (noteproductshape *NoteProductShape) Checkout(stage *Stage) *NoteProductShape {
	if _, ok := stage.NoteProductShapes[noteproductshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteProductShape(noteproductshape)
		}
	}
	return noteproductshape
}

// for satisfaction of GongStruct interface
func (noteproductshape *NoteProductShape) GetName() (res string) {
	return noteproductshape.Name
}

// for satisfaction of GongStruct interface
func (noteproductshape *NoteProductShape) SetName(name string) {
	noteproductshape.Name = name
}

// Stage puts noteshape to the model stage
func (noteshape *NoteShape) Stage(stage *Stage) *NoteShape {
	if _, ok := stage.NoteShapes[noteshape]; !ok {
		stage.NoteShapes[noteshape] = struct{}{}
		stage.NoteShape_stagedOrder[noteshape] = stage.NoteShapeOrder
		stage.NoteShape_orderStaged[stage.NoteShapeOrder] = noteshape
		stage.NoteShapeOrder++
	}
	stage.NoteShapes_mapString[noteshape.Name] = noteshape

	return noteshape
}

// StagePreserveOrder puts noteshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteShapeOrder
// - update stage.NoteShapeOrder accordingly
func (noteshape *NoteShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteShapes[noteshape]; !ok {
		stage.NoteShapes[noteshape] = struct{}{}

		if order > stage.NoteShapeOrder {
			stage.NoteShapeOrder = order
		}
		stage.NoteShape_stagedOrder[noteshape] = order
		stage.NoteShape_orderStaged[order] = noteshape
		stage.NoteShapeOrder++
	}
	stage.NoteShapes_mapString[noteshape.Name] = noteshape
}

// Unstage removes noteshape off the model stage
func (noteshape *NoteShape) Unstage(stage *Stage) *NoteShape {
	delete(stage.NoteShapes, noteshape)
	// issue1150
	// delete(stage.NoteShape_stagedOrder, noteshape)
	delete(stage.NoteShapes_mapString, noteshape.Name)

	return noteshape
}

// UnstageVoid removes noteshape off the model stage
func (noteshape *NoteShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteShapes, noteshape)
	// issue1150
	// delete(stage.NoteShape_stagedOrder, noteshape)
	delete(stage.NoteShapes_mapString, noteshape.Name)
}

// commit noteshape to the back repo (if it is already staged)
func (noteshape *NoteShape) Commit(stage *Stage) *NoteShape {
	if _, ok := stage.NoteShapes[noteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteShape(noteshape)
		}
	}
	return noteshape
}

func (noteshape *NoteShape) CommitVoid(stage *Stage) {
	noteshape.Commit(stage)
}

func (noteshape *NoteShape) StageVoid(stage *Stage) {
	noteshape.Stage(stage)
}

// Checkout noteshape to the back repo (if it is already staged)
func (noteshape *NoteShape) Checkout(stage *Stage) *NoteShape {
	if _, ok := stage.NoteShapes[noteshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteShape(noteshape)
		}
	}
	return noteshape
}

// for satisfaction of GongStruct interface
func (noteshape *NoteShape) GetName() (res string) {
	return noteshape.Name
}

// for satisfaction of GongStruct interface
func (noteshape *NoteShape) SetName(name string) {
	noteshape.Name = name
}

// Stage puts notestakeholdershape to the model stage
func (notestakeholdershape *NoteStakeholderShape) Stage(stage *Stage) *NoteStakeholderShape {
	if _, ok := stage.NoteStakeholderShapes[notestakeholdershape]; !ok {
		stage.NoteStakeholderShapes[notestakeholdershape] = struct{}{}
		stage.NoteStakeholderShape_stagedOrder[notestakeholdershape] = stage.NoteStakeholderShapeOrder
		stage.NoteStakeholderShape_orderStaged[stage.NoteStakeholderShapeOrder] = notestakeholdershape
		stage.NoteStakeholderShapeOrder++
	}
	stage.NoteStakeholderShapes_mapString[notestakeholdershape.Name] = notestakeholdershape

	return notestakeholdershape
}

// StagePreserveOrder puts notestakeholdershape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteStakeholderShapeOrder
// - update stage.NoteStakeholderShapeOrder accordingly
func (notestakeholdershape *NoteStakeholderShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteStakeholderShapes[notestakeholdershape]; !ok {
		stage.NoteStakeholderShapes[notestakeholdershape] = struct{}{}

		if order > stage.NoteStakeholderShapeOrder {
			stage.NoteStakeholderShapeOrder = order
		}
		stage.NoteStakeholderShape_stagedOrder[notestakeholdershape] = order
		stage.NoteStakeholderShape_orderStaged[order] = notestakeholdershape
		stage.NoteStakeholderShapeOrder++
	}
	stage.NoteStakeholderShapes_mapString[notestakeholdershape.Name] = notestakeholdershape
}

// Unstage removes notestakeholdershape off the model stage
func (notestakeholdershape *NoteStakeholderShape) Unstage(stage *Stage) *NoteStakeholderShape {
	delete(stage.NoteStakeholderShapes, notestakeholdershape)
	// issue1150
	// delete(stage.NoteStakeholderShape_stagedOrder, notestakeholdershape)
	delete(stage.NoteStakeholderShapes_mapString, notestakeholdershape.Name)

	return notestakeholdershape
}

// UnstageVoid removes notestakeholdershape off the model stage
func (notestakeholdershape *NoteStakeholderShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteStakeholderShapes, notestakeholdershape)
	// issue1150
	// delete(stage.NoteStakeholderShape_stagedOrder, notestakeholdershape)
	delete(stage.NoteStakeholderShapes_mapString, notestakeholdershape.Name)
}

// commit notestakeholdershape to the back repo (if it is already staged)
func (notestakeholdershape *NoteStakeholderShape) Commit(stage *Stage) *NoteStakeholderShape {
	if _, ok := stage.NoteStakeholderShapes[notestakeholdershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteStakeholderShape(notestakeholdershape)
		}
	}
	return notestakeholdershape
}

func (notestakeholdershape *NoteStakeholderShape) CommitVoid(stage *Stage) {
	notestakeholdershape.Commit(stage)
}

func (notestakeholdershape *NoteStakeholderShape) StageVoid(stage *Stage) {
	notestakeholdershape.Stage(stage)
}

// Checkout notestakeholdershape to the back repo (if it is already staged)
func (notestakeholdershape *NoteStakeholderShape) Checkout(stage *Stage) *NoteStakeholderShape {
	if _, ok := stage.NoteStakeholderShapes[notestakeholdershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteStakeholderShape(notestakeholdershape)
		}
	}
	return notestakeholdershape
}

// for satisfaction of GongStruct interface
func (notestakeholdershape *NoteStakeholderShape) GetName() (res string) {
	return notestakeholdershape.Name
}

// for satisfaction of GongStruct interface
func (notestakeholdershape *NoteStakeholderShape) SetName(name string) {
	notestakeholdershape.Name = name
}

// Stage puts notetaskshape to the model stage
func (notetaskshape *NoteTaskShape) Stage(stage *Stage) *NoteTaskShape {
	if _, ok := stage.NoteTaskShapes[notetaskshape]; !ok {
		stage.NoteTaskShapes[notetaskshape] = struct{}{}
		stage.NoteTaskShape_stagedOrder[notetaskshape] = stage.NoteTaskShapeOrder
		stage.NoteTaskShape_orderStaged[stage.NoteTaskShapeOrder] = notetaskshape
		stage.NoteTaskShapeOrder++
	}
	stage.NoteTaskShapes_mapString[notetaskshape.Name] = notetaskshape

	return notetaskshape
}

// StagePreserveOrder puts notetaskshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteTaskShapeOrder
// - update stage.NoteTaskShapeOrder accordingly
func (notetaskshape *NoteTaskShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteTaskShapes[notetaskshape]; !ok {
		stage.NoteTaskShapes[notetaskshape] = struct{}{}

		if order > stage.NoteTaskShapeOrder {
			stage.NoteTaskShapeOrder = order
		}
		stage.NoteTaskShape_stagedOrder[notetaskshape] = order
		stage.NoteTaskShape_orderStaged[order] = notetaskshape
		stage.NoteTaskShapeOrder++
	}
	stage.NoteTaskShapes_mapString[notetaskshape.Name] = notetaskshape
}

// Unstage removes notetaskshape off the model stage
func (notetaskshape *NoteTaskShape) Unstage(stage *Stage) *NoteTaskShape {
	delete(stage.NoteTaskShapes, notetaskshape)
	// issue1150
	// delete(stage.NoteTaskShape_stagedOrder, notetaskshape)
	delete(stage.NoteTaskShapes_mapString, notetaskshape.Name)

	return notetaskshape
}

// UnstageVoid removes notetaskshape off the model stage
func (notetaskshape *NoteTaskShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteTaskShapes, notetaskshape)
	// issue1150
	// delete(stage.NoteTaskShape_stagedOrder, notetaskshape)
	delete(stage.NoteTaskShapes_mapString, notetaskshape.Name)
}

// commit notetaskshape to the back repo (if it is already staged)
func (notetaskshape *NoteTaskShape) Commit(stage *Stage) *NoteTaskShape {
	if _, ok := stage.NoteTaskShapes[notetaskshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteTaskShape(notetaskshape)
		}
	}
	return notetaskshape
}

func (notetaskshape *NoteTaskShape) CommitVoid(stage *Stage) {
	notetaskshape.Commit(stage)
}

func (notetaskshape *NoteTaskShape) StageVoid(stage *Stage) {
	notetaskshape.Stage(stage)
}

// Checkout notetaskshape to the back repo (if it is already staged)
func (notetaskshape *NoteTaskShape) Checkout(stage *Stage) *NoteTaskShape {
	if _, ok := stage.NoteTaskShapes[notetaskshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteTaskShape(notetaskshape)
		}
	}
	return notetaskshape
}

// for satisfaction of GongStruct interface
func (notetaskshape *NoteTaskShape) GetName() (res string) {
	return notetaskshape.Name
}

// for satisfaction of GongStruct interface
func (notetaskshape *NoteTaskShape) SetName(name string) {
	notetaskshape.Name = name
}

// Stage puts productcompositionshape to the model stage
func (productcompositionshape *ProductCompositionShape) Stage(stage *Stage) *ProductCompositionShape {
	if _, ok := stage.ProductCompositionShapes[productcompositionshape]; !ok {
		stage.ProductCompositionShapes[productcompositionshape] = struct{}{}
		stage.ProductCompositionShape_stagedOrder[productcompositionshape] = stage.ProductCompositionShapeOrder
		stage.ProductCompositionShape_orderStaged[stage.ProductCompositionShapeOrder] = productcompositionshape
		stage.ProductCompositionShapeOrder++
	}
	stage.ProductCompositionShapes_mapString[productcompositionshape.Name] = productcompositionshape

	return productcompositionshape
}

// StagePreserveOrder puts productcompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductCompositionShapeOrder
// - update stage.ProductCompositionShapeOrder accordingly
func (productcompositionshape *ProductCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ProductCompositionShapes[productcompositionshape]; !ok {
		stage.ProductCompositionShapes[productcompositionshape] = struct{}{}

		if order > stage.ProductCompositionShapeOrder {
			stage.ProductCompositionShapeOrder = order
		}
		stage.ProductCompositionShape_stagedOrder[productcompositionshape] = order
		stage.ProductCompositionShape_orderStaged[order] = productcompositionshape
		stage.ProductCompositionShapeOrder++
	}
	stage.ProductCompositionShapes_mapString[productcompositionshape.Name] = productcompositionshape
}

// Unstage removes productcompositionshape off the model stage
func (productcompositionshape *ProductCompositionShape) Unstage(stage *Stage) *ProductCompositionShape {
	delete(stage.ProductCompositionShapes, productcompositionshape)
	// issue1150
	// delete(stage.ProductCompositionShape_stagedOrder, productcompositionshape)
	delete(stage.ProductCompositionShapes_mapString, productcompositionshape.Name)

	return productcompositionshape
}

// UnstageVoid removes productcompositionshape off the model stage
func (productcompositionshape *ProductCompositionShape) UnstageVoid(stage *Stage) {
	delete(stage.ProductCompositionShapes, productcompositionshape)
	// issue1150
	// delete(stage.ProductCompositionShape_stagedOrder, productcompositionshape)
	delete(stage.ProductCompositionShapes_mapString, productcompositionshape.Name)
}

// commit productcompositionshape to the back repo (if it is already staged)
func (productcompositionshape *ProductCompositionShape) Commit(stage *Stage) *ProductCompositionShape {
	if _, ok := stage.ProductCompositionShapes[productcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProductCompositionShape(productcompositionshape)
		}
	}
	return productcompositionshape
}

func (productcompositionshape *ProductCompositionShape) CommitVoid(stage *Stage) {
	productcompositionshape.Commit(stage)
}

func (productcompositionshape *ProductCompositionShape) StageVoid(stage *Stage) {
	productcompositionshape.Stage(stage)
}

// Checkout productcompositionshape to the back repo (if it is already staged)
func (productcompositionshape *ProductCompositionShape) Checkout(stage *Stage) *ProductCompositionShape {
	if _, ok := stage.ProductCompositionShapes[productcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProductCompositionShape(productcompositionshape)
		}
	}
	return productcompositionshape
}

// for satisfaction of GongStruct interface
func (productcompositionshape *ProductCompositionShape) GetName() (res string) {
	return productcompositionshape.Name
}

// for satisfaction of GongStruct interface
func (productcompositionshape *ProductCompositionShape) SetName(name string) {
	productcompositionshape.Name = name
}

// Stage puts productshape to the model stage
func (productshape *ProductShape) Stage(stage *Stage) *ProductShape {
	if _, ok := stage.ProductShapes[productshape]; !ok {
		stage.ProductShapes[productshape] = struct{}{}
		stage.ProductShape_stagedOrder[productshape] = stage.ProductShapeOrder
		stage.ProductShape_orderStaged[stage.ProductShapeOrder] = productshape
		stage.ProductShapeOrder++
	}
	stage.ProductShapes_mapString[productshape.Name] = productshape

	return productshape
}

// StagePreserveOrder puts productshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductShapeOrder
// - update stage.ProductShapeOrder accordingly
func (productshape *ProductShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ProductShapes[productshape]; !ok {
		stage.ProductShapes[productshape] = struct{}{}

		if order > stage.ProductShapeOrder {
			stage.ProductShapeOrder = order
		}
		stage.ProductShape_stagedOrder[productshape] = order
		stage.ProductShape_orderStaged[order] = productshape
		stage.ProductShapeOrder++
	}
	stage.ProductShapes_mapString[productshape.Name] = productshape
}

// Unstage removes productshape off the model stage
func (productshape *ProductShape) Unstage(stage *Stage) *ProductShape {
	delete(stage.ProductShapes, productshape)
	// issue1150
	// delete(stage.ProductShape_stagedOrder, productshape)
	delete(stage.ProductShapes_mapString, productshape.Name)

	return productshape
}

// UnstageVoid removes productshape off the model stage
func (productshape *ProductShape) UnstageVoid(stage *Stage) {
	delete(stage.ProductShapes, productshape)
	// issue1150
	// delete(stage.ProductShape_stagedOrder, productshape)
	delete(stage.ProductShapes_mapString, productshape.Name)
}

// commit productshape to the back repo (if it is already staged)
func (productshape *ProductShape) Commit(stage *Stage) *ProductShape {
	if _, ok := stage.ProductShapes[productshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProductShape(productshape)
		}
	}
	return productshape
}

func (productshape *ProductShape) CommitVoid(stage *Stage) {
	productshape.Commit(stage)
}

func (productshape *ProductShape) StageVoid(stage *Stage) {
	productshape.Stage(stage)
}

// Checkout productshape to the back repo (if it is already staged)
func (productshape *ProductShape) Checkout(stage *Stage) *ProductShape {
	if _, ok := stage.ProductShapes[productshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProductShape(productshape)
		}
	}
	return productshape
}

// for satisfaction of GongStruct interface
func (productshape *ProductShape) GetName() (res string) {
	return productshape.Name
}

// for satisfaction of GongStruct interface
func (productshape *ProductShape) SetName(name string) {
	productshape.Name = name
}

// Stage puts requirement to the model stage
func (requirement *Requirement) Stage(stage *Stage) *Requirement {
	if _, ok := stage.Requirements[requirement]; !ok {
		stage.Requirements[requirement] = struct{}{}
		stage.Requirement_stagedOrder[requirement] = stage.RequirementOrder
		stage.Requirement_orderStaged[stage.RequirementOrder] = requirement
		stage.RequirementOrder++
	}
	stage.Requirements_mapString[requirement.Name] = requirement

	return requirement
}

// StagePreserveOrder puts requirement to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RequirementOrder
// - update stage.RequirementOrder accordingly
func (requirement *Requirement) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Requirements[requirement]; !ok {
		stage.Requirements[requirement] = struct{}{}

		if order > stage.RequirementOrder {
			stage.RequirementOrder = order
		}
		stage.Requirement_stagedOrder[requirement] = order
		stage.Requirement_orderStaged[order] = requirement
		stage.RequirementOrder++
	}
	stage.Requirements_mapString[requirement.Name] = requirement
}

// Unstage removes requirement off the model stage
func (requirement *Requirement) Unstage(stage *Stage) *Requirement {
	delete(stage.Requirements, requirement)
	// issue1150
	// delete(stage.Requirement_stagedOrder, requirement)
	delete(stage.Requirements_mapString, requirement.Name)

	return requirement
}

// UnstageVoid removes requirement off the model stage
func (requirement *Requirement) UnstageVoid(stage *Stage) {
	delete(stage.Requirements, requirement)
	// issue1150
	// delete(stage.Requirement_stagedOrder, requirement)
	delete(stage.Requirements_mapString, requirement.Name)
}

// commit requirement to the back repo (if it is already staged)
func (requirement *Requirement) Commit(stage *Stage) *Requirement {
	if _, ok := stage.Requirements[requirement]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRequirement(requirement)
		}
	}
	return requirement
}

func (requirement *Requirement) CommitVoid(stage *Stage) {
	requirement.Commit(stage)
}

func (requirement *Requirement) StageVoid(stage *Stage) {
	requirement.Stage(stage)
}

// Checkout requirement to the back repo (if it is already staged)
func (requirement *Requirement) Checkout(stage *Stage) *Requirement {
	if _, ok := stage.Requirements[requirement]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRequirement(requirement)
		}
	}
	return requirement
}

// for satisfaction of GongStruct interface
func (requirement *Requirement) GetName() (res string) {
	return requirement.Name
}

// for satisfaction of GongStruct interface
func (requirement *Requirement) SetName(name string) {
	requirement.Name = name
}

// Stage puts stakeholder to the model stage
func (stakeholder *Stakeholder) Stage(stage *Stage) *Stakeholder {
	if _, ok := stage.Stakeholders[stakeholder]; !ok {
		stage.Stakeholders[stakeholder] = struct{}{}
		stage.Stakeholder_stagedOrder[stakeholder] = stage.StakeholderOrder
		stage.Stakeholder_orderStaged[stage.StakeholderOrder] = stakeholder
		stage.StakeholderOrder++
	}
	stage.Stakeholders_mapString[stakeholder.Name] = stakeholder

	return stakeholder
}

// StagePreserveOrder puts stakeholder to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StakeholderOrder
// - update stage.StakeholderOrder accordingly
func (stakeholder *Stakeholder) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Stakeholders[stakeholder]; !ok {
		stage.Stakeholders[stakeholder] = struct{}{}

		if order > stage.StakeholderOrder {
			stage.StakeholderOrder = order
		}
		stage.Stakeholder_stagedOrder[stakeholder] = order
		stage.Stakeholder_orderStaged[order] = stakeholder
		stage.StakeholderOrder++
	}
	stage.Stakeholders_mapString[stakeholder.Name] = stakeholder
}

// Unstage removes stakeholder off the model stage
func (stakeholder *Stakeholder) Unstage(stage *Stage) *Stakeholder {
	delete(stage.Stakeholders, stakeholder)
	// issue1150
	// delete(stage.Stakeholder_stagedOrder, stakeholder)
	delete(stage.Stakeholders_mapString, stakeholder.Name)

	return stakeholder
}

// UnstageVoid removes stakeholder off the model stage
func (stakeholder *Stakeholder) UnstageVoid(stage *Stage) {
	delete(stage.Stakeholders, stakeholder)
	// issue1150
	// delete(stage.Stakeholder_stagedOrder, stakeholder)
	delete(stage.Stakeholders_mapString, stakeholder.Name)
}

// commit stakeholder to the back repo (if it is already staged)
func (stakeholder *Stakeholder) Commit(stage *Stage) *Stakeholder {
	if _, ok := stage.Stakeholders[stakeholder]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStakeholder(stakeholder)
		}
	}
	return stakeholder
}

func (stakeholder *Stakeholder) CommitVoid(stage *Stage) {
	stakeholder.Commit(stage)
}

func (stakeholder *Stakeholder) StageVoid(stage *Stage) {
	stakeholder.Stage(stage)
}

// Checkout stakeholder to the back repo (if it is already staged)
func (stakeholder *Stakeholder) Checkout(stage *Stage) *Stakeholder {
	if _, ok := stage.Stakeholders[stakeholder]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStakeholder(stakeholder)
		}
	}
	return stakeholder
}

// for satisfaction of GongStruct interface
func (stakeholder *Stakeholder) GetName() (res string) {
	return stakeholder.Name
}

// for satisfaction of GongStruct interface
func (stakeholder *Stakeholder) SetName(name string) {
	stakeholder.Name = name
}

// Stage puts stakeholdercompositionshape to the model stage
func (stakeholdercompositionshape *StakeholderCompositionShape) Stage(stage *Stage) *StakeholderCompositionShape {
	if _, ok := stage.StakeholderCompositionShapes[stakeholdercompositionshape]; !ok {
		stage.StakeholderCompositionShapes[stakeholdercompositionshape] = struct{}{}
		stage.StakeholderCompositionShape_stagedOrder[stakeholdercompositionshape] = stage.StakeholderCompositionShapeOrder
		stage.StakeholderCompositionShape_orderStaged[stage.StakeholderCompositionShapeOrder] = stakeholdercompositionshape
		stage.StakeholderCompositionShapeOrder++
	}
	stage.StakeholderCompositionShapes_mapString[stakeholdercompositionshape.Name] = stakeholdercompositionshape

	return stakeholdercompositionshape
}

// StagePreserveOrder puts stakeholdercompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StakeholderCompositionShapeOrder
// - update stage.StakeholderCompositionShapeOrder accordingly
func (stakeholdercompositionshape *StakeholderCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StakeholderCompositionShapes[stakeholdercompositionshape]; !ok {
		stage.StakeholderCompositionShapes[stakeholdercompositionshape] = struct{}{}

		if order > stage.StakeholderCompositionShapeOrder {
			stage.StakeholderCompositionShapeOrder = order
		}
		stage.StakeholderCompositionShape_stagedOrder[stakeholdercompositionshape] = order
		stage.StakeholderCompositionShape_orderStaged[order] = stakeholdercompositionshape
		stage.StakeholderCompositionShapeOrder++
	}
	stage.StakeholderCompositionShapes_mapString[stakeholdercompositionshape.Name] = stakeholdercompositionshape
}

// Unstage removes stakeholdercompositionshape off the model stage
func (stakeholdercompositionshape *StakeholderCompositionShape) Unstage(stage *Stage) *StakeholderCompositionShape {
	delete(stage.StakeholderCompositionShapes, stakeholdercompositionshape)
	// issue1150
	// delete(stage.StakeholderCompositionShape_stagedOrder, stakeholdercompositionshape)
	delete(stage.StakeholderCompositionShapes_mapString, stakeholdercompositionshape.Name)

	return stakeholdercompositionshape
}

// UnstageVoid removes stakeholdercompositionshape off the model stage
func (stakeholdercompositionshape *StakeholderCompositionShape) UnstageVoid(stage *Stage) {
	delete(stage.StakeholderCompositionShapes, stakeholdercompositionshape)
	// issue1150
	// delete(stage.StakeholderCompositionShape_stagedOrder, stakeholdercompositionshape)
	delete(stage.StakeholderCompositionShapes_mapString, stakeholdercompositionshape.Name)
}

// commit stakeholdercompositionshape to the back repo (if it is already staged)
func (stakeholdercompositionshape *StakeholderCompositionShape) Commit(stage *Stage) *StakeholderCompositionShape {
	if _, ok := stage.StakeholderCompositionShapes[stakeholdercompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStakeholderCompositionShape(stakeholdercompositionshape)
		}
	}
	return stakeholdercompositionshape
}

func (stakeholdercompositionshape *StakeholderCompositionShape) CommitVoid(stage *Stage) {
	stakeholdercompositionshape.Commit(stage)
}

func (stakeholdercompositionshape *StakeholderCompositionShape) StageVoid(stage *Stage) {
	stakeholdercompositionshape.Stage(stage)
}

// Checkout stakeholdercompositionshape to the back repo (if it is already staged)
func (stakeholdercompositionshape *StakeholderCompositionShape) Checkout(stage *Stage) *StakeholderCompositionShape {
	if _, ok := stage.StakeholderCompositionShapes[stakeholdercompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStakeholderCompositionShape(stakeholdercompositionshape)
		}
	}
	return stakeholdercompositionshape
}

// for satisfaction of GongStruct interface
func (stakeholdercompositionshape *StakeholderCompositionShape) GetName() (res string) {
	return stakeholdercompositionshape.Name
}

// for satisfaction of GongStruct interface
func (stakeholdercompositionshape *StakeholderCompositionShape) SetName(name string) {
	stakeholdercompositionshape.Name = name
}

// Stage puts stakeholderconcernshape to the model stage
func (stakeholderconcernshape *StakeholderConcernShape) Stage(stage *Stage) *StakeholderConcernShape {
	if _, ok := stage.StakeholderConcernShapes[stakeholderconcernshape]; !ok {
		stage.StakeholderConcernShapes[stakeholderconcernshape] = struct{}{}
		stage.StakeholderConcernShape_stagedOrder[stakeholderconcernshape] = stage.StakeholderConcernShapeOrder
		stage.StakeholderConcernShape_orderStaged[stage.StakeholderConcernShapeOrder] = stakeholderconcernshape
		stage.StakeholderConcernShapeOrder++
	}
	stage.StakeholderConcernShapes_mapString[stakeholderconcernshape.Name] = stakeholderconcernshape

	return stakeholderconcernshape
}

// StagePreserveOrder puts stakeholderconcernshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StakeholderConcernShapeOrder
// - update stage.StakeholderConcernShapeOrder accordingly
func (stakeholderconcernshape *StakeholderConcernShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StakeholderConcernShapes[stakeholderconcernshape]; !ok {
		stage.StakeholderConcernShapes[stakeholderconcernshape] = struct{}{}

		if order > stage.StakeholderConcernShapeOrder {
			stage.StakeholderConcernShapeOrder = order
		}
		stage.StakeholderConcernShape_stagedOrder[stakeholderconcernshape] = order
		stage.StakeholderConcernShape_orderStaged[order] = stakeholderconcernshape
		stage.StakeholderConcernShapeOrder++
	}
	stage.StakeholderConcernShapes_mapString[stakeholderconcernshape.Name] = stakeholderconcernshape
}

// Unstage removes stakeholderconcernshape off the model stage
func (stakeholderconcernshape *StakeholderConcernShape) Unstage(stage *Stage) *StakeholderConcernShape {
	delete(stage.StakeholderConcernShapes, stakeholderconcernshape)
	// issue1150
	// delete(stage.StakeholderConcernShape_stagedOrder, stakeholderconcernshape)
	delete(stage.StakeholderConcernShapes_mapString, stakeholderconcernshape.Name)

	return stakeholderconcernshape
}

// UnstageVoid removes stakeholderconcernshape off the model stage
func (stakeholderconcernshape *StakeholderConcernShape) UnstageVoid(stage *Stage) {
	delete(stage.StakeholderConcernShapes, stakeholderconcernshape)
	// issue1150
	// delete(stage.StakeholderConcernShape_stagedOrder, stakeholderconcernshape)
	delete(stage.StakeholderConcernShapes_mapString, stakeholderconcernshape.Name)
}

// commit stakeholderconcernshape to the back repo (if it is already staged)
func (stakeholderconcernshape *StakeholderConcernShape) Commit(stage *Stage) *StakeholderConcernShape {
	if _, ok := stage.StakeholderConcernShapes[stakeholderconcernshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStakeholderConcernShape(stakeholderconcernshape)
		}
	}
	return stakeholderconcernshape
}

func (stakeholderconcernshape *StakeholderConcernShape) CommitVoid(stage *Stage) {
	stakeholderconcernshape.Commit(stage)
}

func (stakeholderconcernshape *StakeholderConcernShape) StageVoid(stage *Stage) {
	stakeholderconcernshape.Stage(stage)
}

// Checkout stakeholderconcernshape to the back repo (if it is already staged)
func (stakeholderconcernshape *StakeholderConcernShape) Checkout(stage *Stage) *StakeholderConcernShape {
	if _, ok := stage.StakeholderConcernShapes[stakeholderconcernshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStakeholderConcernShape(stakeholderconcernshape)
		}
	}
	return stakeholderconcernshape
}

// for satisfaction of GongStruct interface
func (stakeholderconcernshape *StakeholderConcernShape) GetName() (res string) {
	return stakeholderconcernshape.Name
}

// for satisfaction of GongStruct interface
func (stakeholderconcernshape *StakeholderConcernShape) SetName(name string) {
	stakeholderconcernshape.Name = name
}

// Stage puts stakeholdershape to the model stage
func (stakeholdershape *StakeholderShape) Stage(stage *Stage) *StakeholderShape {
	if _, ok := stage.StakeholderShapes[stakeholdershape]; !ok {
		stage.StakeholderShapes[stakeholdershape] = struct{}{}
		stage.StakeholderShape_stagedOrder[stakeholdershape] = stage.StakeholderShapeOrder
		stage.StakeholderShape_orderStaged[stage.StakeholderShapeOrder] = stakeholdershape
		stage.StakeholderShapeOrder++
	}
	stage.StakeholderShapes_mapString[stakeholdershape.Name] = stakeholdershape

	return stakeholdershape
}

// StagePreserveOrder puts stakeholdershape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StakeholderShapeOrder
// - update stage.StakeholderShapeOrder accordingly
func (stakeholdershape *StakeholderShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StakeholderShapes[stakeholdershape]; !ok {
		stage.StakeholderShapes[stakeholdershape] = struct{}{}

		if order > stage.StakeholderShapeOrder {
			stage.StakeholderShapeOrder = order
		}
		stage.StakeholderShape_stagedOrder[stakeholdershape] = order
		stage.StakeholderShape_orderStaged[order] = stakeholdershape
		stage.StakeholderShapeOrder++
	}
	stage.StakeholderShapes_mapString[stakeholdershape.Name] = stakeholdershape
}

// Unstage removes stakeholdershape off the model stage
func (stakeholdershape *StakeholderShape) Unstage(stage *Stage) *StakeholderShape {
	delete(stage.StakeholderShapes, stakeholdershape)
	// issue1150
	// delete(stage.StakeholderShape_stagedOrder, stakeholdershape)
	delete(stage.StakeholderShapes_mapString, stakeholdershape.Name)

	return stakeholdershape
}

// UnstageVoid removes stakeholdershape off the model stage
func (stakeholdershape *StakeholderShape) UnstageVoid(stage *Stage) {
	delete(stage.StakeholderShapes, stakeholdershape)
	// issue1150
	// delete(stage.StakeholderShape_stagedOrder, stakeholdershape)
	delete(stage.StakeholderShapes_mapString, stakeholdershape.Name)
}

// commit stakeholdershape to the back repo (if it is already staged)
func (stakeholdershape *StakeholderShape) Commit(stage *Stage) *StakeholderShape {
	if _, ok := stage.StakeholderShapes[stakeholdershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStakeholderShape(stakeholdershape)
		}
	}
	return stakeholdershape
}

func (stakeholdershape *StakeholderShape) CommitVoid(stage *Stage) {
	stakeholdershape.Commit(stage)
}

func (stakeholdershape *StakeholderShape) StageVoid(stage *Stage) {
	stakeholdershape.Stage(stage)
}

// Checkout stakeholdershape to the back repo (if it is already staged)
func (stakeholdershape *StakeholderShape) Checkout(stage *Stage) *StakeholderShape {
	if _, ok := stage.StakeholderShapes[stakeholdershape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStakeholderShape(stakeholdershape)
		}
	}
	return stakeholdershape
}

// for satisfaction of GongStruct interface
func (stakeholdershape *StakeholderShape) GetName() (res string) {
	return stakeholdershape.Name
}

// for satisfaction of GongStruct interface
func (stakeholdershape *StakeholderShape) SetName(name string) {
	stakeholdershape.Name = name
}

// Stage puts supportlevel to the model stage
func (supportlevel *SupportLevel) Stage(stage *Stage) *SupportLevel {
	if _, ok := stage.SupportLevels[supportlevel]; !ok {
		stage.SupportLevels[supportlevel] = struct{}{}
		stage.SupportLevel_stagedOrder[supportlevel] = stage.SupportLevelOrder
		stage.SupportLevel_orderStaged[stage.SupportLevelOrder] = supportlevel
		stage.SupportLevelOrder++
	}
	stage.SupportLevels_mapString[supportlevel.Name] = supportlevel

	return supportlevel
}

// StagePreserveOrder puts supportlevel to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SupportLevelOrder
// - update stage.SupportLevelOrder accordingly
func (supportlevel *SupportLevel) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SupportLevels[supportlevel]; !ok {
		stage.SupportLevels[supportlevel] = struct{}{}

		if order > stage.SupportLevelOrder {
			stage.SupportLevelOrder = order
		}
		stage.SupportLevel_stagedOrder[supportlevel] = order
		stage.SupportLevel_orderStaged[order] = supportlevel
		stage.SupportLevelOrder++
	}
	stage.SupportLevels_mapString[supportlevel.Name] = supportlevel
}

// Unstage removes supportlevel off the model stage
func (supportlevel *SupportLevel) Unstage(stage *Stage) *SupportLevel {
	delete(stage.SupportLevels, supportlevel)
	// issue1150
	// delete(stage.SupportLevel_stagedOrder, supportlevel)
	delete(stage.SupportLevels_mapString, supportlevel.Name)

	return supportlevel
}

// UnstageVoid removes supportlevel off the model stage
func (supportlevel *SupportLevel) UnstageVoid(stage *Stage) {
	delete(stage.SupportLevels, supportlevel)
	// issue1150
	// delete(stage.SupportLevel_stagedOrder, supportlevel)
	delete(stage.SupportLevels_mapString, supportlevel.Name)
}

// commit supportlevel to the back repo (if it is already staged)
func (supportlevel *SupportLevel) Commit(stage *Stage) *SupportLevel {
	if _, ok := stage.SupportLevels[supportlevel]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSupportLevel(supportlevel)
		}
	}
	return supportlevel
}

func (supportlevel *SupportLevel) CommitVoid(stage *Stage) {
	supportlevel.Commit(stage)
}

func (supportlevel *SupportLevel) StageVoid(stage *Stage) {
	supportlevel.Stage(stage)
}

// Checkout supportlevel to the back repo (if it is already staged)
func (supportlevel *SupportLevel) Checkout(stage *Stage) *SupportLevel {
	if _, ok := stage.SupportLevels[supportlevel]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSupportLevel(supportlevel)
		}
	}
	return supportlevel
}

// for satisfaction of GongStruct interface
func (supportlevel *SupportLevel) GetName() (res string) {
	return supportlevel.Name
}

// for satisfaction of GongStruct interface
func (supportlevel *SupportLevel) SetName(name string) {
	supportlevel.Name = name
}

// Stage puts tool to the model stage
func (tool *Tool) Stage(stage *Stage) *Tool {
	if _, ok := stage.Tools[tool]; !ok {
		stage.Tools[tool] = struct{}{}
		stage.Tool_stagedOrder[tool] = stage.ToolOrder
		stage.Tool_orderStaged[stage.ToolOrder] = tool
		stage.ToolOrder++
	}
	stage.Tools_mapString[tool.Name] = tool

	return tool
}

// StagePreserveOrder puts tool to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ToolOrder
// - update stage.ToolOrder accordingly
func (tool *Tool) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Tools[tool]; !ok {
		stage.Tools[tool] = struct{}{}

		if order > stage.ToolOrder {
			stage.ToolOrder = order
		}
		stage.Tool_stagedOrder[tool] = order
		stage.Tool_orderStaged[order] = tool
		stage.ToolOrder++
	}
	stage.Tools_mapString[tool.Name] = tool
}

// Unstage removes tool off the model stage
func (tool *Tool) Unstage(stage *Stage) *Tool {
	delete(stage.Tools, tool)
	// issue1150
	// delete(stage.Tool_stagedOrder, tool)
	delete(stage.Tools_mapString, tool.Name)

	return tool
}

// UnstageVoid removes tool off the model stage
func (tool *Tool) UnstageVoid(stage *Stage) {
	delete(stage.Tools, tool)
	// issue1150
	// delete(stage.Tool_stagedOrder, tool)
	delete(stage.Tools_mapString, tool.Name)
}

// commit tool to the back repo (if it is already staged)
func (tool *Tool) Commit(stage *Stage) *Tool {
	if _, ok := stage.Tools[tool]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTool(tool)
		}
	}
	return tool
}

func (tool *Tool) CommitVoid(stage *Stage) {
	tool.Commit(stage)
}

func (tool *Tool) StageVoid(stage *Stage) {
	tool.Stage(stage)
}

// Checkout tool to the back repo (if it is already staged)
func (tool *Tool) Checkout(stage *Stage) *Tool {
	if _, ok := stage.Tools[tool]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTool(tool)
		}
	}
	return tool
}

// for satisfaction of GongStruct interface
func (tool *Tool) GetName() (res string) {
	return tool.Name
}

// for satisfaction of GongStruct interface
func (tool *Tool) SetName(name string) {
	tool.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAnalysisNeed(AnalysisNeed *AnalysisNeed)
	CreateORMConcept(Concept *Concept)
	CreateORMConcern(Concern *Concern)
	CreateORMConcernCompositionShape(ConcernCompositionShape *ConcernCompositionShape)
	CreateORMConcernInputShape(ConcernInputShape *ConcernInputShape)
	CreateORMConcernOutputShape(ConcernOutputShape *ConcernOutputShape)
	CreateORMConcernShape(ConcernShape *ConcernShape)
	CreateORMDeliverable(Deliverable *Deliverable)
	CreateORMDiagram(Diagram *Diagram)
	CreateORMLibrary(Library *Library)
	CreateORMNote(Note *Note)
	CreateORMNoteProductShape(NoteProductShape *NoteProductShape)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMNoteStakeholderShape(NoteStakeholderShape *NoteStakeholderShape)
	CreateORMNoteTaskShape(NoteTaskShape *NoteTaskShape)
	CreateORMProductCompositionShape(ProductCompositionShape *ProductCompositionShape)
	CreateORMProductShape(ProductShape *ProductShape)
	CreateORMRequirement(Requirement *Requirement)
	CreateORMStakeholder(Stakeholder *Stakeholder)
	CreateORMStakeholderCompositionShape(StakeholderCompositionShape *StakeholderCompositionShape)
	CreateORMStakeholderConcernShape(StakeholderConcernShape *StakeholderConcernShape)
	CreateORMStakeholderShape(StakeholderShape *StakeholderShape)
	CreateORMSupportLevel(SupportLevel *SupportLevel)
	CreateORMTool(Tool *Tool)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAnalysisNeed(AnalysisNeed *AnalysisNeed)
	DeleteORMConcept(Concept *Concept)
	DeleteORMConcern(Concern *Concern)
	DeleteORMConcernCompositionShape(ConcernCompositionShape *ConcernCompositionShape)
	DeleteORMConcernInputShape(ConcernInputShape *ConcernInputShape)
	DeleteORMConcernOutputShape(ConcernOutputShape *ConcernOutputShape)
	DeleteORMConcernShape(ConcernShape *ConcernShape)
	DeleteORMDeliverable(Deliverable *Deliverable)
	DeleteORMDiagram(Diagram *Diagram)
	DeleteORMLibrary(Library *Library)
	DeleteORMNote(Note *Note)
	DeleteORMNoteProductShape(NoteProductShape *NoteProductShape)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMNoteStakeholderShape(NoteStakeholderShape *NoteStakeholderShape)
	DeleteORMNoteTaskShape(NoteTaskShape *NoteTaskShape)
	DeleteORMProductCompositionShape(ProductCompositionShape *ProductCompositionShape)
	DeleteORMProductShape(ProductShape *ProductShape)
	DeleteORMRequirement(Requirement *Requirement)
	DeleteORMStakeholder(Stakeholder *Stakeholder)
	DeleteORMStakeholderCompositionShape(StakeholderCompositionShape *StakeholderCompositionShape)
	DeleteORMStakeholderConcernShape(StakeholderConcernShape *StakeholderConcernShape)
	DeleteORMStakeholderShape(StakeholderShape *StakeholderShape)
	DeleteORMSupportLevel(SupportLevel *SupportLevel)
	DeleteORMTool(Tool *Tool)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.AnalysisNeeds = make(map[*AnalysisNeed]struct{})
	stage.AnalysisNeeds_mapString = make(map[string]*AnalysisNeed)
	stage.AnalysisNeed_stagedOrder = make(map[*AnalysisNeed]uint)
	stage.AnalysisNeedOrder = 0

	stage.Concepts = make(map[*Concept]struct{})
	stage.Concepts_mapString = make(map[string]*Concept)
	stage.Concept_stagedOrder = make(map[*Concept]uint)
	stage.ConceptOrder = 0

	stage.Concerns = make(map[*Concern]struct{})
	stage.Concerns_mapString = make(map[string]*Concern)
	stage.Concern_stagedOrder = make(map[*Concern]uint)
	stage.ConcernOrder = 0

	stage.ConcernCompositionShapes = make(map[*ConcernCompositionShape]struct{})
	stage.ConcernCompositionShapes_mapString = make(map[string]*ConcernCompositionShape)
	stage.ConcernCompositionShape_stagedOrder = make(map[*ConcernCompositionShape]uint)
	stage.ConcernCompositionShapeOrder = 0

	stage.ConcernInputShapes = make(map[*ConcernInputShape]struct{})
	stage.ConcernInputShapes_mapString = make(map[string]*ConcernInputShape)
	stage.ConcernInputShape_stagedOrder = make(map[*ConcernInputShape]uint)
	stage.ConcernInputShapeOrder = 0

	stage.ConcernOutputShapes = make(map[*ConcernOutputShape]struct{})
	stage.ConcernOutputShapes_mapString = make(map[string]*ConcernOutputShape)
	stage.ConcernOutputShape_stagedOrder = make(map[*ConcernOutputShape]uint)
	stage.ConcernOutputShapeOrder = 0

	stage.ConcernShapes = make(map[*ConcernShape]struct{})
	stage.ConcernShapes_mapString = make(map[string]*ConcernShape)
	stage.ConcernShape_stagedOrder = make(map[*ConcernShape]uint)
	stage.ConcernShapeOrder = 0

	stage.Deliverables = make(map[*Deliverable]struct{})
	stage.Deliverables_mapString = make(map[string]*Deliverable)
	stage.Deliverable_stagedOrder = make(map[*Deliverable]uint)
	stage.DeliverableOrder = 0

	stage.Diagrams = make(map[*Diagram]struct{})
	stage.Diagrams_mapString = make(map[string]*Diagram)
	stage.Diagram_stagedOrder = make(map[*Diagram]uint)
	stage.DiagramOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Notes = make(map[*Note]struct{})
	stage.Notes_mapString = make(map[string]*Note)
	stage.Note_stagedOrder = make(map[*Note]uint)
	stage.NoteOrder = 0

	stage.NoteProductShapes = make(map[*NoteProductShape]struct{})
	stage.NoteProductShapes_mapString = make(map[string]*NoteProductShape)
	stage.NoteProductShape_stagedOrder = make(map[*NoteProductShape]uint)
	stage.NoteProductShapeOrder = 0

	stage.NoteShapes = make(map[*NoteShape]struct{})
	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	stage.NoteShape_stagedOrder = make(map[*NoteShape]uint)
	stage.NoteShapeOrder = 0

	stage.NoteStakeholderShapes = make(map[*NoteStakeholderShape]struct{})
	stage.NoteStakeholderShapes_mapString = make(map[string]*NoteStakeholderShape)
	stage.NoteStakeholderShape_stagedOrder = make(map[*NoteStakeholderShape]uint)
	stage.NoteStakeholderShapeOrder = 0

	stage.NoteTaskShapes = make(map[*NoteTaskShape]struct{})
	stage.NoteTaskShapes_mapString = make(map[string]*NoteTaskShape)
	stage.NoteTaskShape_stagedOrder = make(map[*NoteTaskShape]uint)
	stage.NoteTaskShapeOrder = 0

	stage.ProductCompositionShapes = make(map[*ProductCompositionShape]struct{})
	stage.ProductCompositionShapes_mapString = make(map[string]*ProductCompositionShape)
	stage.ProductCompositionShape_stagedOrder = make(map[*ProductCompositionShape]uint)
	stage.ProductCompositionShapeOrder = 0

	stage.ProductShapes = make(map[*ProductShape]struct{})
	stage.ProductShapes_mapString = make(map[string]*ProductShape)
	stage.ProductShape_stagedOrder = make(map[*ProductShape]uint)
	stage.ProductShapeOrder = 0

	stage.Requirements = make(map[*Requirement]struct{})
	stage.Requirements_mapString = make(map[string]*Requirement)
	stage.Requirement_stagedOrder = make(map[*Requirement]uint)
	stage.RequirementOrder = 0

	stage.Stakeholders = make(map[*Stakeholder]struct{})
	stage.Stakeholders_mapString = make(map[string]*Stakeholder)
	stage.Stakeholder_stagedOrder = make(map[*Stakeholder]uint)
	stage.StakeholderOrder = 0

	stage.StakeholderCompositionShapes = make(map[*StakeholderCompositionShape]struct{})
	stage.StakeholderCompositionShapes_mapString = make(map[string]*StakeholderCompositionShape)
	stage.StakeholderCompositionShape_stagedOrder = make(map[*StakeholderCompositionShape]uint)
	stage.StakeholderCompositionShapeOrder = 0

	stage.StakeholderConcernShapes = make(map[*StakeholderConcernShape]struct{})
	stage.StakeholderConcernShapes_mapString = make(map[string]*StakeholderConcernShape)
	stage.StakeholderConcernShape_stagedOrder = make(map[*StakeholderConcernShape]uint)
	stage.StakeholderConcernShapeOrder = 0

	stage.StakeholderShapes = make(map[*StakeholderShape]struct{})
	stage.StakeholderShapes_mapString = make(map[string]*StakeholderShape)
	stage.StakeholderShape_stagedOrder = make(map[*StakeholderShape]uint)
	stage.StakeholderShapeOrder = 0

	stage.SupportLevels = make(map[*SupportLevel]struct{})
	stage.SupportLevels_mapString = make(map[string]*SupportLevel)
	stage.SupportLevel_stagedOrder = make(map[*SupportLevel]uint)
	stage.SupportLevelOrder = 0

	stage.Tools = make(map[*Tool]struct{})
	stage.Tools_mapString = make(map[string]*Tool)
	stage.Tool_stagedOrder = make(map[*Tool]uint)
	stage.ToolOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.AnalysisNeeds = nil
	stage.AnalysisNeeds_mapString = nil

	stage.Concepts = nil
	stage.Concepts_mapString = nil

	stage.Concerns = nil
	stage.Concerns_mapString = nil

	stage.ConcernCompositionShapes = nil
	stage.ConcernCompositionShapes_mapString = nil

	stage.ConcernInputShapes = nil
	stage.ConcernInputShapes_mapString = nil

	stage.ConcernOutputShapes = nil
	stage.ConcernOutputShapes_mapString = nil

	stage.ConcernShapes = nil
	stage.ConcernShapes_mapString = nil

	stage.Deliverables = nil
	stage.Deliverables_mapString = nil

	stage.Diagrams = nil
	stage.Diagrams_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Notes = nil
	stage.Notes_mapString = nil

	stage.NoteProductShapes = nil
	stage.NoteProductShapes_mapString = nil

	stage.NoteShapes = nil
	stage.NoteShapes_mapString = nil

	stage.NoteStakeholderShapes = nil
	stage.NoteStakeholderShapes_mapString = nil

	stage.NoteTaskShapes = nil
	stage.NoteTaskShapes_mapString = nil

	stage.ProductCompositionShapes = nil
	stage.ProductCompositionShapes_mapString = nil

	stage.ProductShapes = nil
	stage.ProductShapes_mapString = nil

	stage.Requirements = nil
	stage.Requirements_mapString = nil

	stage.Stakeholders = nil
	stage.Stakeholders_mapString = nil

	stage.StakeholderCompositionShapes = nil
	stage.StakeholderCompositionShapes_mapString = nil

	stage.StakeholderConcernShapes = nil
	stage.StakeholderConcernShapes_mapString = nil

	stage.StakeholderShapes = nil
	stage.StakeholderShapes_mapString = nil

	stage.SupportLevels = nil
	stage.SupportLevels_mapString = nil

	stage.Tools = nil
	stage.Tools_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for analysisneed := range stage.AnalysisNeeds {
		analysisneed.Unstage(stage)
	}

	for concept := range stage.Concepts {
		concept.Unstage(stage)
	}

	for concern := range stage.Concerns {
		concern.Unstage(stage)
	}

	for concerncompositionshape := range stage.ConcernCompositionShapes {
		concerncompositionshape.Unstage(stage)
	}

	for concerninputshape := range stage.ConcernInputShapes {
		concerninputshape.Unstage(stage)
	}

	for concernoutputshape := range stage.ConcernOutputShapes {
		concernoutputshape.Unstage(stage)
	}

	for concernshape := range stage.ConcernShapes {
		concernshape.Unstage(stage)
	}

	for deliverable := range stage.Deliverables {
		deliverable.Unstage(stage)
	}

	for diagram := range stage.Diagrams {
		diagram.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for note := range stage.Notes {
		note.Unstage(stage)
	}

	for noteproductshape := range stage.NoteProductShapes {
		noteproductshape.Unstage(stage)
	}

	for noteshape := range stage.NoteShapes {
		noteshape.Unstage(stage)
	}

	for notestakeholdershape := range stage.NoteStakeholderShapes {
		notestakeholdershape.Unstage(stage)
	}

	for notetaskshape := range stage.NoteTaskShapes {
		notetaskshape.Unstage(stage)
	}

	for productcompositionshape := range stage.ProductCompositionShapes {
		productcompositionshape.Unstage(stage)
	}

	for productshape := range stage.ProductShapes {
		productshape.Unstage(stage)
	}

	for requirement := range stage.Requirements {
		requirement.Unstage(stage)
	}

	for stakeholder := range stage.Stakeholders {
		stakeholder.Unstage(stage)
	}

	for stakeholdercompositionshape := range stage.StakeholderCompositionShapes {
		stakeholdercompositionshape.Unstage(stage)
	}

	for stakeholderconcernshape := range stage.StakeholderConcernShapes {
		stakeholderconcernshape.Unstage(stage)
	}

	for stakeholdershape := range stage.StakeholderShapes {
		stakeholdershape.Unstage(stage)
	}

	for supportlevel := range stage.SupportLevels {
		supportlevel.Unstage(stage)
	}

	for tool := range stage.Tools {
		tool.Unstage(stage)
	}

	// end of insertion point for array nil
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface{}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage) (modified bool)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) GongstructIF
	GongGetUUID(stage *Stage) string
}
type PointerToGongstruct interface {
	GongstructIF
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]struct{}) (sortedSlice []T) {
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *Stage) (sortedSlice []T) {
	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*AnalysisNeed]any:
		return any(&stage.AnalysisNeeds).(*Type)
	case map[*Concept]any:
		return any(&stage.Concepts).(*Type)
	case map[*Concern]any:
		return any(&stage.Concerns).(*Type)
	case map[*ConcernCompositionShape]any:
		return any(&stage.ConcernCompositionShapes).(*Type)
	case map[*ConcernInputShape]any:
		return any(&stage.ConcernInputShapes).(*Type)
	case map[*ConcernOutputShape]any:
		return any(&stage.ConcernOutputShapes).(*Type)
	case map[*ConcernShape]any:
		return any(&stage.ConcernShapes).(*Type)
	case map[*Deliverable]any:
		return any(&stage.Deliverables).(*Type)
	case map[*Diagram]any:
		return any(&stage.Diagrams).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Note]any:
		return any(&stage.Notes).(*Type)
	case map[*NoteProductShape]any:
		return any(&stage.NoteProductShapes).(*Type)
	case map[*NoteShape]any:
		return any(&stage.NoteShapes).(*Type)
	case map[*NoteStakeholderShape]any:
		return any(&stage.NoteStakeholderShapes).(*Type)
	case map[*NoteTaskShape]any:
		return any(&stage.NoteTaskShapes).(*Type)
	case map[*ProductCompositionShape]any:
		return any(&stage.ProductCompositionShapes).(*Type)
	case map[*ProductShape]any:
		return any(&stage.ProductShapes).(*Type)
	case map[*Requirement]any:
		return any(&stage.Requirements).(*Type)
	case map[*Stakeholder]any:
		return any(&stage.Stakeholders).(*Type)
	case map[*StakeholderCompositionShape]any:
		return any(&stage.StakeholderCompositionShapes).(*Type)
	case map[*StakeholderConcernShape]any:
		return any(&stage.StakeholderConcernShapes).(*Type)
	case map[*StakeholderShape]any:
		return any(&stage.StakeholderShapes).(*Type)
	case map[*SupportLevel]any:
		return any(&stage.SupportLevels).(*Type)
	case map[*Tool]any:
		return any(&stage.Tools).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged Gonstruct instance by their name
// Can be usefull if names are unique
func GongGetMap[Type GongstructIF](stage *Stage) map[string]Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AnalysisNeed:
		return any(stage.AnalysisNeeds_mapString).(map[string]Type)
	case *Concept:
		return any(stage.Concepts_mapString).(map[string]Type)
	case *Concern:
		return any(stage.Concerns_mapString).(map[string]Type)
	case *ConcernCompositionShape:
		return any(stage.ConcernCompositionShapes_mapString).(map[string]Type)
	case *ConcernInputShape:
		return any(stage.ConcernInputShapes_mapString).(map[string]Type)
	case *ConcernOutputShape:
		return any(stage.ConcernOutputShapes_mapString).(map[string]Type)
	case *ConcernShape:
		return any(stage.ConcernShapes_mapString).(map[string]Type)
	case *Deliverable:
		return any(stage.Deliverables_mapString).(map[string]Type)
	case *Diagram:
		return any(stage.Diagrams_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Note:
		return any(stage.Notes_mapString).(map[string]Type)
	case *NoteProductShape:
		return any(stage.NoteProductShapes_mapString).(map[string]Type)
	case *NoteShape:
		return any(stage.NoteShapes_mapString).(map[string]Type)
	case *NoteStakeholderShape:
		return any(stage.NoteStakeholderShapes_mapString).(map[string]Type)
	case *NoteTaskShape:
		return any(stage.NoteTaskShapes_mapString).(map[string]Type)
	case *ProductCompositionShape:
		return any(stage.ProductCompositionShapes_mapString).(map[string]Type)
	case *ProductShape:
		return any(stage.ProductShapes_mapString).(map[string]Type)
	case *Requirement:
		return any(stage.Requirements_mapString).(map[string]Type)
	case *Stakeholder:
		return any(stage.Stakeholders_mapString).(map[string]Type)
	case *StakeholderCompositionShape:
		return any(stage.StakeholderCompositionShapes_mapString).(map[string]Type)
	case *StakeholderConcernShape:
		return any(stage.StakeholderConcernShapes_mapString).(map[string]Type)
	case *StakeholderShape:
		return any(stage.StakeholderShapes_mapString).(map[string]Type)
	case *SupportLevel:
		return any(stage.SupportLevels_mapString).(map[string]Type)
	case *Tool:
		return any(stage.Tools_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AnalysisNeed:
		return any(&stage.AnalysisNeeds).(*map[*Type]struct{})
	case Concept:
		return any(&stage.Concepts).(*map[*Type]struct{})
	case Concern:
		return any(&stage.Concerns).(*map[*Type]struct{})
	case ConcernCompositionShape:
		return any(&stage.ConcernCompositionShapes).(*map[*Type]struct{})
	case ConcernInputShape:
		return any(&stage.ConcernInputShapes).(*map[*Type]struct{})
	case ConcernOutputShape:
		return any(&stage.ConcernOutputShapes).(*map[*Type]struct{})
	case ConcernShape:
		return any(&stage.ConcernShapes).(*map[*Type]struct{})
	case Deliverable:
		return any(&stage.Deliverables).(*map[*Type]struct{})
	case Diagram:
		return any(&stage.Diagrams).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Note:
		return any(&stage.Notes).(*map[*Type]struct{})
	case NoteProductShape:
		return any(&stage.NoteProductShapes).(*map[*Type]struct{})
	case NoteShape:
		return any(&stage.NoteShapes).(*map[*Type]struct{})
	case NoteStakeholderShape:
		return any(&stage.NoteStakeholderShapes).(*map[*Type]struct{})
	case NoteTaskShape:
		return any(&stage.NoteTaskShapes).(*map[*Type]struct{})
	case ProductCompositionShape:
		return any(&stage.ProductCompositionShapes).(*map[*Type]struct{})
	case ProductShape:
		return any(&stage.ProductShapes).(*map[*Type]struct{})
	case Requirement:
		return any(&stage.Requirements).(*map[*Type]struct{})
	case Stakeholder:
		return any(&stage.Stakeholders).(*map[*Type]struct{})
	case StakeholderCompositionShape:
		return any(&stage.StakeholderCompositionShapes).(*map[*Type]struct{})
	case StakeholderConcernShape:
		return any(&stage.StakeholderConcernShapes).(*map[*Type]struct{})
	case StakeholderShape:
		return any(&stage.StakeholderShapes).(*map[*Type]struct{})
	case SupportLevel:
		return any(&stage.SupportLevels).(*map[*Type]struct{})
	case Tool:
		return any(&stage.Tools).(*map[*Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AnalysisNeed:
		return any(&stage.AnalysisNeeds).(*map[Type]struct{})
	case *Concept:
		return any(&stage.Concepts).(*map[Type]struct{})
	case *Concern:
		return any(&stage.Concerns).(*map[Type]struct{})
	case *ConcernCompositionShape:
		return any(&stage.ConcernCompositionShapes).(*map[Type]struct{})
	case *ConcernInputShape:
		return any(&stage.ConcernInputShapes).(*map[Type]struct{})
	case *ConcernOutputShape:
		return any(&stage.ConcernOutputShapes).(*map[Type]struct{})
	case *ConcernShape:
		return any(&stage.ConcernShapes).(*map[Type]struct{})
	case *Deliverable:
		return any(&stage.Deliverables).(*map[Type]struct{})
	case *Diagram:
		return any(&stage.Diagrams).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Note:
		return any(&stage.Notes).(*map[Type]struct{})
	case *NoteProductShape:
		return any(&stage.NoteProductShapes).(*map[Type]struct{})
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]struct{})
	case *NoteStakeholderShape:
		return any(&stage.NoteStakeholderShapes).(*map[Type]struct{})
	case *NoteTaskShape:
		return any(&stage.NoteTaskShapes).(*map[Type]struct{})
	case *ProductCompositionShape:
		return any(&stage.ProductCompositionShapes).(*map[Type]struct{})
	case *ProductShape:
		return any(&stage.ProductShapes).(*map[Type]struct{})
	case *Requirement:
		return any(&stage.Requirements).(*map[Type]struct{})
	case *Stakeholder:
		return any(&stage.Stakeholders).(*map[Type]struct{})
	case *StakeholderCompositionShape:
		return any(&stage.StakeholderCompositionShapes).(*map[Type]struct{})
	case *StakeholderConcernShape:
		return any(&stage.StakeholderConcernShapes).(*map[Type]struct{})
	case *StakeholderShape:
		return any(&stage.StakeholderShapes).(*map[Type]struct{})
	case *SupportLevel:
		return any(&stage.SupportLevels).(*map[Type]struct{})
	case *Tool:
		return any(&stage.Tools).(*map[Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AnalysisNeed:
		return any(&stage.AnalysisNeeds_mapString).(*map[string]*Type)
	case Concept:
		return any(&stage.Concepts_mapString).(*map[string]*Type)
	case Concern:
		return any(&stage.Concerns_mapString).(*map[string]*Type)
	case ConcernCompositionShape:
		return any(&stage.ConcernCompositionShapes_mapString).(*map[string]*Type)
	case ConcernInputShape:
		return any(&stage.ConcernInputShapes_mapString).(*map[string]*Type)
	case ConcernOutputShape:
		return any(&stage.ConcernOutputShapes_mapString).(*map[string]*Type)
	case ConcernShape:
		return any(&stage.ConcernShapes_mapString).(*map[string]*Type)
	case Deliverable:
		return any(&stage.Deliverables_mapString).(*map[string]*Type)
	case Diagram:
		return any(&stage.Diagrams_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Note:
		return any(&stage.Notes_mapString).(*map[string]*Type)
	case NoteProductShape:
		return any(&stage.NoteProductShapes_mapString).(*map[string]*Type)
	case NoteShape:
		return any(&stage.NoteShapes_mapString).(*map[string]*Type)
	case NoteStakeholderShape:
		return any(&stage.NoteStakeholderShapes_mapString).(*map[string]*Type)
	case NoteTaskShape:
		return any(&stage.NoteTaskShapes_mapString).(*map[string]*Type)
	case ProductCompositionShape:
		return any(&stage.ProductCompositionShapes_mapString).(*map[string]*Type)
	case ProductShape:
		return any(&stage.ProductShapes_mapString).(*map[string]*Type)
	case Requirement:
		return any(&stage.Requirements_mapString).(*map[string]*Type)
	case Stakeholder:
		return any(&stage.Stakeholders_mapString).(*map[string]*Type)
	case StakeholderCompositionShape:
		return any(&stage.StakeholderCompositionShapes_mapString).(*map[string]*Type)
	case StakeholderConcernShape:
		return any(&stage.StakeholderConcernShapes_mapString).(*map[string]*Type)
	case StakeholderShape:
		return any(&stage.StakeholderShapes_mapString).(*map[string]*Type)
	case SupportLevel:
		return any(&stage.SupportLevels_mapString).(*map[string]*Type)
	case Tool:
		return any(&stage.Tools_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case AnalysisNeed:
		return any(&AnalysisNeed{
			// Initialisation of associations
		}).(*Type)
	case Concept:
		return any(&Concept{
			// Initialisation of associations
			// field is initialized with an instance of Tool with the name of the field
			Tools: []*Tool{{Name: "Tools"}},
		}).(*Type)
	case Concern:
		return any(&Concern{
			// Initialisation of associations
			// field is initialized with an instance of Concern with the name of the field
			SubConcerns: []*Concern{{Name: "SubConcerns"}},
			// field is initialized with an instance of Deliverable with the name of the field
			Inputs: []*Deliverable{{Name: "Inputs"}},
			// field is initialized with an instance of Deliverable with the name of the field
			Outputs: []*Deliverable{{Name: "Outputs"}},
			// field is initialized with an instance of Requirement with the name of the field
			Requirements: []*Requirement{{Name: "Requirements"}},
		}).(*Type)
	case ConcernCompositionShape:
		return any(&ConcernCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Concern with the name of the field
			Concern: &Concern{Name: "Concern"},
		}).(*Type)
	case ConcernInputShape:
		return any(&ConcernInputShape{
			// Initialisation of associations
			// field is initialized with an instance of Deliverable with the name of the field
			Deliverable: &Deliverable{Name: "Deliverable"},
			// field is initialized with an instance of Concern with the name of the field
			Concern: &Concern{Name: "Concern"},
		}).(*Type)
	case ConcernOutputShape:
		return any(&ConcernOutputShape{
			// Initialisation of associations
			// field is initialized with an instance of Concern with the name of the field
			Task: &Concern{Name: "Task"},
			// field is initialized with an instance of Deliverable with the name of the field
			Product: &Deliverable{Name: "Product"},
		}).(*Type)
	case ConcernShape:
		return any(&ConcernShape{
			// Initialisation of associations
			// field is initialized with an instance of Concern with the name of the field
			Concern: &Concern{Name: "Concern"},
		}).(*Type)
	case Deliverable:
		return any(&Deliverable{
			// Initialisation of associations
			// field is initialized with an instance of Deliverable with the name of the field
			SubProducts: []*Deliverable{{Name: "SubProducts"}},
			// field is initialized with an instance of Concept with the name of the field
			Concepts: []*Concept{{Name: "Concepts"}},
		}).(*Type)
	case Diagram:
		return any(&Diagram{
			// Initialisation of associations
			// field is initialized with an instance of Concern with the name of the field
			ConcernsWhoseRequirementsNodeIsExpanded: []*Concern{{Name: "ConcernsWhoseRequirementsNodeIsExpanded"}},
			// field is initialized with an instance of ProductShape with the name of the field
			Product_Shapes: []*ProductShape{{Name: "Product_Shapes"}},
			// field is initialized with an instance of Deliverable with the name of the field
			ProductsWhoseNodeIsExpanded: []*Deliverable{{Name: "ProductsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ProductCompositionShape with the name of the field
			ProductComposition_Shapes: []*ProductCompositionShape{{Name: "ProductComposition_Shapes"}},
			// field is initialized with an instance of ConcernShape with the name of the field
			Concern_Shapes: []*ConcernShape{{Name: "Concern_Shapes"}},
			// field is initialized with an instance of Concern with the name of the field
			ConcernsWhoseNodeIsExpanded: []*Concern{{Name: "ConcernsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Concern with the name of the field
			ConcernsWhoseInputNodeIsExpanded: []*Concern{{Name: "ConcernsWhoseInputNodeIsExpanded"}},
			// field is initialized with an instance of Concern with the name of the field
			ConcernsWhoseStakeholderNodeIsExpanded: []*Concern{{Name: "ConcernsWhoseStakeholderNodeIsExpanded"}},
			// field is initialized with an instance of Concern with the name of the field
			ConcernssWhoseOutputNodeIsExpanded: []*Concern{{Name: "ConcernssWhoseOutputNodeIsExpanded"}},
			// field is initialized with an instance of ConcernCompositionShape with the name of the field
			ConcernComposition_Shapes: []*ConcernCompositionShape{{Name: "ConcernComposition_Shapes"}},
			// field is initialized with an instance of ConcernInputShape with the name of the field
			ConcernInputShapes: []*ConcernInputShape{{Name: "ConcernInputShapes"}},
			// field is initialized with an instance of ConcernOutputShape with the name of the field
			ConcernOutputShapes: []*ConcernOutputShape{{Name: "ConcernOutputShapes"}},
			// field is initialized with an instance of NoteShape with the name of the field
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of NoteProductShape with the name of the field
			NoteProductShapes: []*NoteProductShape{{Name: "NoteProductShapes"}},
			// field is initialized with an instance of NoteTaskShape with the name of the field
			NoteTaskShapes: []*NoteTaskShape{{Name: "NoteTaskShapes"}},
			// field is initialized with an instance of NoteStakeholderShape with the name of the field
			NoteResourceShapes: []*NoteStakeholderShape{{Name: "NoteResourceShapes"}},
			// field is initialized with an instance of StakeholderShape with the name of the field
			Stakeholder_Shapes: []*StakeholderShape{{Name: "Stakeholder_Shapes"}},
			// field is initialized with an instance of Stakeholder with the name of the field
			ResourcesWhoseNodeIsExpanded: []*Stakeholder{{Name: "ResourcesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of StakeholderCompositionShape with the name of the field
			ResourceComposition_Shapes: []*StakeholderCompositionShape{{Name: "ResourceComposition_Shapes"}},
			// field is initialized with an instance of StakeholderConcernShape with the name of the field
			StakeholderConcernShapes: []*StakeholderConcernShape{{Name: "StakeholderConcernShapes"}},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Deliverable with the name of the field
			RootDeliverables: []*Deliverable{{Name: "RootDeliverables"}},
			// field is initialized with an instance of Concern with the name of the field
			RootConcerns: []*Concern{{Name: "RootConcerns"}},
			// field is initialized with an instance of Stakeholder with the name of the field
			RootStakeholders: []*Stakeholder{{Name: "RootStakeholders"}},
			// field is initialized with an instance of Requirement with the name of the field
			RootRequirements: []*Requirement{{Name: "RootRequirements"}},
			// field is initialized with an instance of Concept with the name of the field
			RootConcepts: []*Concept{{Name: "RootConcepts"}},
			// field is initialized with an instance of AnalysisNeed with the name of the field
			AnalysisNeeds: []*AnalysisNeed{{Name: "AnalysisNeeds"}},
			// field is initialized with an instance of Note with the name of the field
			Notes: []*Note{{Name: "Notes"}},
			// field is initialized with an instance of Diagram with the name of the field
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
		}).(*Type)
	case Note:
		return any(&Note{
			// Initialisation of associations
			// field is initialized with an instance of Deliverable with the name of the field
			Products: []*Deliverable{{Name: "Products"}},
			// field is initialized with an instance of Concern with the name of the field
			Tasks: []*Concern{{Name: "Tasks"}},
			// field is initialized with an instance of Stakeholder with the name of the field
			Resources: []*Stakeholder{{Name: "Resources"}},
		}).(*Type)
	case NoteProductShape:
		return any(&NoteProductShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Deliverable with the name of the field
			Product: &Deliverable{Name: "Product"},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
		}).(*Type)
	case NoteStakeholderShape:
		return any(&NoteStakeholderShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Stakeholder with the name of the field
			Stakeholder: &Stakeholder{Name: "Stakeholder"},
		}).(*Type)
	case NoteTaskShape:
		return any(&NoteTaskShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Concern with the name of the field
			Task: &Concern{Name: "Task"},
		}).(*Type)
	case ProductCompositionShape:
		return any(&ProductCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Deliverable with the name of the field
			Product: &Deliverable{Name: "Product"},
		}).(*Type)
	case ProductShape:
		return any(&ProductShape{
			// Initialisation of associations
			// field is initialized with an instance of Deliverable with the name of the field
			Product: &Deliverable{Name: "Product"},
		}).(*Type)
	case Requirement:
		return any(&Requirement{
			// Initialisation of associations
			// field is initialized with an instance of SupportLevel with the name of the field
			SupportLevels: []*SupportLevel{{Name: "SupportLevels"}},
			// field is initialized with an instance of Concept with the name of the field
			Concepts: []*Concept{{Name: "Concepts"}},
		}).(*Type)
	case Stakeholder:
		return any(&Stakeholder{
			// Initialisation of associations
			// field is initialized with an instance of Concern with the name of the field
			Concerns: []*Concern{{Name: "Concerns"}},
			// field is initialized with an instance of Stakeholder with the name of the field
			SubStakeholders: []*Stakeholder{{Name: "SubStakeholders"}},
		}).(*Type)
	case StakeholderCompositionShape:
		return any(&StakeholderCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Stakeholder with the name of the field
			Stakeholder: &Stakeholder{Name: "Stakeholder"},
		}).(*Type)
	case StakeholderConcernShape:
		return any(&StakeholderConcernShape{
			// Initialisation of associations
			// field is initialized with an instance of Stakeholder with the name of the field
			Stakeholder: &Stakeholder{Name: "Stakeholder"},
			// field is initialized with an instance of Concern with the name of the field
			Concern: &Concern{Name: "Concern"},
		}).(*Type)
	case StakeholderShape:
		return any(&StakeholderShape{
			// Initialisation of associations
			// field is initialized with an instance of Stakeholder with the name of the field
			Stakeholder: &Stakeholder{Name: "Stakeholder"},
		}).(*Type)
	case SupportLevel:
		return any(&SupportLevel{
			// Initialisation of associations
			// field is initialized with an instance of Tool with the name of the field
			Tool: &Tool{Name: "Tool"},
		}).(*Type)
	case Tool:
		return any(&Tool{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AnalysisNeed
	case AnalysisNeed:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Concept
	case Concept:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Concern
	case Concern:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ConcernCompositionShape
	case ConcernCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Concern":
			res := make(map[*Concern][]*ConcernCompositionShape)
			for concerncompositionshape := range stage.ConcernCompositionShapes {
				if concerncompositionshape.Concern != nil {
					concern_ := concerncompositionshape.Concern
					var concerncompositionshapes []*ConcernCompositionShape
					_, ok := res[concern_]
					if ok {
						concerncompositionshapes = res[concern_]
					} else {
						concerncompositionshapes = make([]*ConcernCompositionShape, 0)
					}
					concerncompositionshapes = append(concerncompositionshapes, concerncompositionshape)
					res[concern_] = concerncompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ConcernInputShape
	case ConcernInputShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Deliverable":
			res := make(map[*Deliverable][]*ConcernInputShape)
			for concerninputshape := range stage.ConcernInputShapes {
				if concerninputshape.Deliverable != nil {
					deliverable_ := concerninputshape.Deliverable
					var concerninputshapes []*ConcernInputShape
					_, ok := res[deliverable_]
					if ok {
						concerninputshapes = res[deliverable_]
					} else {
						concerninputshapes = make([]*ConcernInputShape, 0)
					}
					concerninputshapes = append(concerninputshapes, concerninputshape)
					res[deliverable_] = concerninputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Concern":
			res := make(map[*Concern][]*ConcernInputShape)
			for concerninputshape := range stage.ConcernInputShapes {
				if concerninputshape.Concern != nil {
					concern_ := concerninputshape.Concern
					var concerninputshapes []*ConcernInputShape
					_, ok := res[concern_]
					if ok {
						concerninputshapes = res[concern_]
					} else {
						concerninputshapes = make([]*ConcernInputShape, 0)
					}
					concerninputshapes = append(concerninputshapes, concerninputshape)
					res[concern_] = concerninputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ConcernOutputShape
	case ConcernOutputShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Task":
			res := make(map[*Concern][]*ConcernOutputShape)
			for concernoutputshape := range stage.ConcernOutputShapes {
				if concernoutputshape.Task != nil {
					concern_ := concernoutputshape.Task
					var concernoutputshapes []*ConcernOutputShape
					_, ok := res[concern_]
					if ok {
						concernoutputshapes = res[concern_]
					} else {
						concernoutputshapes = make([]*ConcernOutputShape, 0)
					}
					concernoutputshapes = append(concernoutputshapes, concernoutputshape)
					res[concern_] = concernoutputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Product":
			res := make(map[*Deliverable][]*ConcernOutputShape)
			for concernoutputshape := range stage.ConcernOutputShapes {
				if concernoutputshape.Product != nil {
					deliverable_ := concernoutputshape.Product
					var concernoutputshapes []*ConcernOutputShape
					_, ok := res[deliverable_]
					if ok {
						concernoutputshapes = res[deliverable_]
					} else {
						concernoutputshapes = make([]*ConcernOutputShape, 0)
					}
					concernoutputshapes = append(concernoutputshapes, concernoutputshape)
					res[deliverable_] = concernoutputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ConcernShape
	case ConcernShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Concern":
			res := make(map[*Concern][]*ConcernShape)
			for concernshape := range stage.ConcernShapes {
				if concernshape.Concern != nil {
					concern_ := concernshape.Concern
					var concernshapes []*ConcernShape
					_, ok := res[concern_]
					if ok {
						concernshapes = res[concern_]
					} else {
						concernshapes = make([]*ConcernShape, 0)
					}
					concernshapes = append(concernshapes, concernshape)
					res[concern_] = concernshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Deliverable
	case Deliverable:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteProductShape
	case NoteProductShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteProductShape)
			for noteproductshape := range stage.NoteProductShapes {
				if noteproductshape.Note != nil {
					note_ := noteproductshape.Note
					var noteproductshapes []*NoteProductShape
					_, ok := res[note_]
					if ok {
						noteproductshapes = res[note_]
					} else {
						noteproductshapes = make([]*NoteProductShape, 0)
					}
					noteproductshapes = append(noteproductshapes, noteproductshape)
					res[note_] = noteproductshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Product":
			res := make(map[*Deliverable][]*NoteProductShape)
			for noteproductshape := range stage.NoteProductShapes {
				if noteproductshape.Product != nil {
					deliverable_ := noteproductshape.Product
					var noteproductshapes []*NoteProductShape
					_, ok := res[deliverable_]
					if ok {
						noteproductshapes = res[deliverable_]
					} else {
						noteproductshapes = make([]*NoteProductShape, 0)
					}
					noteproductshapes = append(noteproductshapes, noteproductshape)
					res[deliverable_] = noteproductshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteShape)
			for noteshape := range stage.NoteShapes {
				if noteshape.Note != nil {
					note_ := noteshape.Note
					var noteshapes []*NoteShape
					_, ok := res[note_]
					if ok {
						noteshapes = res[note_]
					} else {
						noteshapes = make([]*NoteShape, 0)
					}
					noteshapes = append(noteshapes, noteshape)
					res[note_] = noteshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteStakeholderShape
	case NoteStakeholderShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteStakeholderShape)
			for notestakeholdershape := range stage.NoteStakeholderShapes {
				if notestakeholdershape.Note != nil {
					note_ := notestakeholdershape.Note
					var notestakeholdershapes []*NoteStakeholderShape
					_, ok := res[note_]
					if ok {
						notestakeholdershapes = res[note_]
					} else {
						notestakeholdershapes = make([]*NoteStakeholderShape, 0)
					}
					notestakeholdershapes = append(notestakeholdershapes, notestakeholdershape)
					res[note_] = notestakeholdershapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Stakeholder":
			res := make(map[*Stakeholder][]*NoteStakeholderShape)
			for notestakeholdershape := range stage.NoteStakeholderShapes {
				if notestakeholdershape.Stakeholder != nil {
					stakeholder_ := notestakeholdershape.Stakeholder
					var notestakeholdershapes []*NoteStakeholderShape
					_, ok := res[stakeholder_]
					if ok {
						notestakeholdershapes = res[stakeholder_]
					} else {
						notestakeholdershapes = make([]*NoteStakeholderShape, 0)
					}
					notestakeholdershapes = append(notestakeholdershapes, notestakeholdershape)
					res[stakeholder_] = notestakeholdershapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteTaskShape
	case NoteTaskShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteTaskShape)
			for notetaskshape := range stage.NoteTaskShapes {
				if notetaskshape.Note != nil {
					note_ := notetaskshape.Note
					var notetaskshapes []*NoteTaskShape
					_, ok := res[note_]
					if ok {
						notetaskshapes = res[note_]
					} else {
						notetaskshapes = make([]*NoteTaskShape, 0)
					}
					notetaskshapes = append(notetaskshapes, notetaskshape)
					res[note_] = notetaskshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Task":
			res := make(map[*Concern][]*NoteTaskShape)
			for notetaskshape := range stage.NoteTaskShapes {
				if notetaskshape.Task != nil {
					concern_ := notetaskshape.Task
					var notetaskshapes []*NoteTaskShape
					_, ok := res[concern_]
					if ok {
						notetaskshapes = res[concern_]
					} else {
						notetaskshapes = make([]*NoteTaskShape, 0)
					}
					notetaskshapes = append(notetaskshapes, notetaskshape)
					res[concern_] = notetaskshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProductCompositionShape
	case ProductCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Product":
			res := make(map[*Deliverable][]*ProductCompositionShape)
			for productcompositionshape := range stage.ProductCompositionShapes {
				if productcompositionshape.Product != nil {
					deliverable_ := productcompositionshape.Product
					var productcompositionshapes []*ProductCompositionShape
					_, ok := res[deliverable_]
					if ok {
						productcompositionshapes = res[deliverable_]
					} else {
						productcompositionshapes = make([]*ProductCompositionShape, 0)
					}
					productcompositionshapes = append(productcompositionshapes, productcompositionshape)
					res[deliverable_] = productcompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProductShape
	case ProductShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Product":
			res := make(map[*Deliverable][]*ProductShape)
			for productshape := range stage.ProductShapes {
				if productshape.Product != nil {
					deliverable_ := productshape.Product
					var productshapes []*ProductShape
					_, ok := res[deliverable_]
					if ok {
						productshapes = res[deliverable_]
					} else {
						productshapes = make([]*ProductShape, 0)
					}
					productshapes = append(productshapes, productshape)
					res[deliverable_] = productshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Requirement
	case Requirement:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Stakeholder
	case Stakeholder:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StakeholderCompositionShape
	case StakeholderCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Stakeholder":
			res := make(map[*Stakeholder][]*StakeholderCompositionShape)
			for stakeholdercompositionshape := range stage.StakeholderCompositionShapes {
				if stakeholdercompositionshape.Stakeholder != nil {
					stakeholder_ := stakeholdercompositionshape.Stakeholder
					var stakeholdercompositionshapes []*StakeholderCompositionShape
					_, ok := res[stakeholder_]
					if ok {
						stakeholdercompositionshapes = res[stakeholder_]
					} else {
						stakeholdercompositionshapes = make([]*StakeholderCompositionShape, 0)
					}
					stakeholdercompositionshapes = append(stakeholdercompositionshapes, stakeholdercompositionshape)
					res[stakeholder_] = stakeholdercompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StakeholderConcernShape
	case StakeholderConcernShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Stakeholder":
			res := make(map[*Stakeholder][]*StakeholderConcernShape)
			for stakeholderconcernshape := range stage.StakeholderConcernShapes {
				if stakeholderconcernshape.Stakeholder != nil {
					stakeholder_ := stakeholderconcernshape.Stakeholder
					var stakeholderconcernshapes []*StakeholderConcernShape
					_, ok := res[stakeholder_]
					if ok {
						stakeholderconcernshapes = res[stakeholder_]
					} else {
						stakeholderconcernshapes = make([]*StakeholderConcernShape, 0)
					}
					stakeholderconcernshapes = append(stakeholderconcernshapes, stakeholderconcernshape)
					res[stakeholder_] = stakeholderconcernshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Concern":
			res := make(map[*Concern][]*StakeholderConcernShape)
			for stakeholderconcernshape := range stage.StakeholderConcernShapes {
				if stakeholderconcernshape.Concern != nil {
					concern_ := stakeholderconcernshape.Concern
					var stakeholderconcernshapes []*StakeholderConcernShape
					_, ok := res[concern_]
					if ok {
						stakeholderconcernshapes = res[concern_]
					} else {
						stakeholderconcernshapes = make([]*StakeholderConcernShape, 0)
					}
					stakeholderconcernshapes = append(stakeholderconcernshapes, stakeholderconcernshape)
					res[concern_] = stakeholderconcernshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StakeholderShape
	case StakeholderShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Stakeholder":
			res := make(map[*Stakeholder][]*StakeholderShape)
			for stakeholdershape := range stage.StakeholderShapes {
				if stakeholdershape.Stakeholder != nil {
					stakeholder_ := stakeholdershape.Stakeholder
					var stakeholdershapes []*StakeholderShape
					_, ok := res[stakeholder_]
					if ok {
						stakeholdershapes = res[stakeholder_]
					} else {
						stakeholdershapes = make([]*StakeholderShape, 0)
					}
					stakeholdershapes = append(stakeholdershapes, stakeholdershape)
					res[stakeholder_] = stakeholdershapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SupportLevel
	case SupportLevel:
		switch fieldname {
		// insertion point for per direct association field
		case "Tool":
			res := make(map[*Tool][]*SupportLevel)
			for supportlevel := range stage.SupportLevels {
				if supportlevel.Tool != nil {
					tool_ := supportlevel.Tool
					var supportlevels []*SupportLevel
					_, ok := res[tool_]
					if ok {
						supportlevels = res[tool_]
					} else {
						supportlevels = make([]*SupportLevel, 0)
					}
					supportlevels = append(supportlevels, supportlevel)
					res[tool_] = supportlevels
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Tool
	case Tool:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {
	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AnalysisNeed
	case AnalysisNeed:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Concept
	case Concept:
		switch fieldname {
		// insertion point for per direct association field
		case "Tools":
			res := make(map[*Tool][]*Concept)
			for concept := range stage.Concepts {
				for _, tool_ := range concept.Tools {
					res[tool_] = append(res[tool_], concept)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Concern
	case Concern:
		switch fieldname {
		// insertion point for per direct association field
		case "SubConcerns":
			res := make(map[*Concern][]*Concern)
			for concern := range stage.Concerns {
				for _, concern_ := range concern.SubConcerns {
					res[concern_] = append(res[concern_], concern)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Inputs":
			res := make(map[*Deliverable][]*Concern)
			for concern := range stage.Concerns {
				for _, deliverable_ := range concern.Inputs {
					res[deliverable_] = append(res[deliverable_], concern)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Outputs":
			res := make(map[*Deliverable][]*Concern)
			for concern := range stage.Concerns {
				for _, deliverable_ := range concern.Outputs {
					res[deliverable_] = append(res[deliverable_], concern)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Requirements":
			res := make(map[*Requirement][]*Concern)
			for concern := range stage.Concerns {
				for _, requirement_ := range concern.Requirements {
					res[requirement_] = append(res[requirement_], concern)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ConcernCompositionShape
	case ConcernCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ConcernInputShape
	case ConcernInputShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ConcernOutputShape
	case ConcernOutputShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ConcernShape
	case ConcernShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Deliverable
	case Deliverable:
		switch fieldname {
		// insertion point for per direct association field
		case "SubProducts":
			res := make(map[*Deliverable][]*Deliverable)
			for deliverable := range stage.Deliverables {
				for _, deliverable_ := range deliverable.SubProducts {
					res[deliverable_] = append(res[deliverable_], deliverable)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Concepts":
			res := make(map[*Concept][]*Deliverable)
			for deliverable := range stage.Deliverables {
				for _, concept_ := range deliverable.Concepts {
					res[concept_] = append(res[concept_], deliverable)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		case "ConcernsWhoseRequirementsNodeIsExpanded":
			res := make(map[*Concern][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concern_ := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
					res[concern_] = append(res[concern_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Product_Shapes":
			res := make(map[*ProductShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, productshape_ := range diagram.Product_Shapes {
					res[productshape_] = append(res[productshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProductsWhoseNodeIsExpanded":
			res := make(map[*Deliverable][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, deliverable_ := range diagram.ProductsWhoseNodeIsExpanded {
					res[deliverable_] = append(res[deliverable_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProductComposition_Shapes":
			res := make(map[*ProductCompositionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, productcompositionshape_ := range diagram.ProductComposition_Shapes {
					res[productcompositionshape_] = append(res[productcompositionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Concern_Shapes":
			res := make(map[*ConcernShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concernshape_ := range diagram.Concern_Shapes {
					res[concernshape_] = append(res[concernshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConcernsWhoseNodeIsExpanded":
			res := make(map[*Concern][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concern_ := range diagram.ConcernsWhoseNodeIsExpanded {
					res[concern_] = append(res[concern_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConcernsWhoseInputNodeIsExpanded":
			res := make(map[*Concern][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concern_ := range diagram.ConcernsWhoseInputNodeIsExpanded {
					res[concern_] = append(res[concern_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConcernsWhoseStakeholderNodeIsExpanded":
			res := make(map[*Concern][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concern_ := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
					res[concern_] = append(res[concern_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConcernssWhoseOutputNodeIsExpanded":
			res := make(map[*Concern][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concern_ := range diagram.ConcernssWhoseOutputNodeIsExpanded {
					res[concern_] = append(res[concern_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConcernComposition_Shapes":
			res := make(map[*ConcernCompositionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concerncompositionshape_ := range diagram.ConcernComposition_Shapes {
					res[concerncompositionshape_] = append(res[concerncompositionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConcernInputShapes":
			res := make(map[*ConcernInputShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concerninputshape_ := range diagram.ConcernInputShapes {
					res[concerninputshape_] = append(res[concerninputshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConcernOutputShapes":
			res := make(map[*ConcernOutputShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concernoutputshape_ := range diagram.ConcernOutputShapes {
					res[concernoutputshape_] = append(res[concernoutputshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Note_Shapes":
			res := make(map[*NoteShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, noteshape_ := range diagram.Note_Shapes {
					res[noteshape_] = append(res[noteshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotesWhoseNodeIsExpanded":
			res := make(map[*Note][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, note_ := range diagram.NotesWhoseNodeIsExpanded {
					res[note_] = append(res[note_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteProductShapes":
			res := make(map[*NoteProductShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, noteproductshape_ := range diagram.NoteProductShapes {
					res[noteproductshape_] = append(res[noteproductshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteTaskShapes":
			res := make(map[*NoteTaskShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, notetaskshape_ := range diagram.NoteTaskShapes {
					res[notetaskshape_] = append(res[notetaskshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteResourceShapes":
			res := make(map[*NoteStakeholderShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, notestakeholdershape_ := range diagram.NoteResourceShapes {
					res[notestakeholdershape_] = append(res[notestakeholdershape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Stakeholder_Shapes":
			res := make(map[*StakeholderShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, stakeholdershape_ := range diagram.Stakeholder_Shapes {
					res[stakeholdershape_] = append(res[stakeholdershape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ResourcesWhoseNodeIsExpanded":
			res := make(map[*Stakeholder][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, stakeholder_ := range diagram.ResourcesWhoseNodeIsExpanded {
					res[stakeholder_] = append(res[stakeholder_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ResourceComposition_Shapes":
			res := make(map[*StakeholderCompositionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, stakeholdercompositionshape_ := range diagram.ResourceComposition_Shapes {
					res[stakeholdercompositionshape_] = append(res[stakeholdercompositionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "StakeholderConcernShapes":
			res := make(map[*StakeholderConcernShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, stakeholderconcernshape_ := range diagram.StakeholderConcernShapes {
					res[stakeholderconcernshape_] = append(res[stakeholderconcernshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		case "RootDeliverables":
			res := make(map[*Deliverable][]*Library)
			for library := range stage.Librarys {
				for _, deliverable_ := range library.RootDeliverables {
					res[deliverable_] = append(res[deliverable_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootConcerns":
			res := make(map[*Concern][]*Library)
			for library := range stage.Librarys {
				for _, concern_ := range library.RootConcerns {
					res[concern_] = append(res[concern_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootStakeholders":
			res := make(map[*Stakeholder][]*Library)
			for library := range stage.Librarys {
				for _, stakeholder_ := range library.RootStakeholders {
					res[stakeholder_] = append(res[stakeholder_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootRequirements":
			res := make(map[*Requirement][]*Library)
			for library := range stage.Librarys {
				for _, requirement_ := range library.RootRequirements {
					res[requirement_] = append(res[requirement_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootConcepts":
			res := make(map[*Concept][]*Library)
			for library := range stage.Librarys {
				for _, concept_ := range library.RootConcepts {
					res[concept_] = append(res[concept_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AnalysisNeeds":
			res := make(map[*AnalysisNeed][]*Library)
			for library := range stage.Librarys {
				for _, analysisneed_ := range library.AnalysisNeeds {
					res[analysisneed_] = append(res[analysisneed_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Notes":
			res := make(map[*Note][]*Library)
			for library := range stage.Librarys {
				for _, note_ := range library.Notes {
					res[note_] = append(res[note_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Diagrams":
			res := make(map[*Diagram][]*Library)
			for library := range stage.Librarys {
				for _, diagram_ := range library.Diagrams {
					res[diagram_] = append(res[diagram_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubLibraries":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibraries {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "Products":
			res := make(map[*Deliverable][]*Note)
			for note := range stage.Notes {
				for _, deliverable_ := range note.Products {
					res[deliverable_] = append(res[deliverable_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tasks":
			res := make(map[*Concern][]*Note)
			for note := range stage.Notes {
				for _, concern_ := range note.Tasks {
					res[concern_] = append(res[concern_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Resources":
			res := make(map[*Stakeholder][]*Note)
			for note := range stage.Notes {
				for _, stakeholder_ := range note.Resources {
					res[stakeholder_] = append(res[stakeholder_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteProductShape
	case NoteProductShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteStakeholderShape
	case NoteStakeholderShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteTaskShape
	case NoteTaskShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ProductCompositionShape
	case ProductCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ProductShape
	case ProductShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Requirement
	case Requirement:
		switch fieldname {
		// insertion point for per direct association field
		case "SupportLevels":
			res := make(map[*SupportLevel][]*Requirement)
			for requirement := range stage.Requirements {
				for _, supportlevel_ := range requirement.SupportLevels {
					res[supportlevel_] = append(res[supportlevel_], requirement)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Concepts":
			res := make(map[*Concept][]*Requirement)
			for requirement := range stage.Requirements {
				for _, concept_ := range requirement.Concepts {
					res[concept_] = append(res[concept_], requirement)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Stakeholder
	case Stakeholder:
		switch fieldname {
		// insertion point for per direct association field
		case "Concerns":
			res := make(map[*Concern][]*Stakeholder)
			for stakeholder := range stage.Stakeholders {
				for _, concern_ := range stakeholder.Concerns {
					res[concern_] = append(res[concern_], stakeholder)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubStakeholders":
			res := make(map[*Stakeholder][]*Stakeholder)
			for stakeholder := range stage.Stakeholders {
				for _, stakeholder_ := range stakeholder.SubStakeholders {
					res[stakeholder_] = append(res[stakeholder_], stakeholder)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StakeholderCompositionShape
	case StakeholderCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StakeholderConcernShape
	case StakeholderConcernShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of StakeholderShape
	case StakeholderShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SupportLevel
	case SupportLevel:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tool
	case Tool:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AnalysisNeed:
		res = "AnalysisNeed"
	case *Concept:
		res = "Concept"
	case *Concern:
		res = "Concern"
	case *ConcernCompositionShape:
		res = "ConcernCompositionShape"
	case *ConcernInputShape:
		res = "ConcernInputShape"
	case *ConcernOutputShape:
		res = "ConcernOutputShape"
	case *ConcernShape:
		res = "ConcernShape"
	case *Deliverable:
		res = "Deliverable"
	case *Diagram:
		res = "Diagram"
	case *Library:
		res = "Library"
	case *Note:
		res = "Note"
	case *NoteProductShape:
		res = "NoteProductShape"
	case *NoteShape:
		res = "NoteShape"
	case *NoteStakeholderShape:
		res = "NoteStakeholderShape"
	case *NoteTaskShape:
		res = "NoteTaskShape"
	case *ProductCompositionShape:
		res = "ProductCompositionShape"
	case *ProductShape:
		res = "ProductShape"
	case *Requirement:
		res = "Requirement"
	case *Stakeholder:
		res = "Stakeholder"
	case *StakeholderCompositionShape:
		res = "StakeholderCompositionShape"
	case *StakeholderConcernShape:
		res = "StakeholderConcernShape"
	case *StakeholderShape:
		res = "StakeholderShape"
	case *SupportLevel:
		res = "SupportLevel"
	case *Tool:
		res = "Tool"
	}
	return res
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {
	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *AnalysisNeed:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "AnalysisNeeds"
		res = append(res, rf)
	case *Concept:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Deliverable"
		rf.Fieldname = "Concepts"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootConcepts"
		res = append(res, rf)
		rf.GongstructName = "Requirement"
		rf.Fieldname = "Concepts"
		res = append(res, rf)
	case *Concern:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Concern"
		rf.Fieldname = "SubConcerns"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConcernsWhoseRequirementsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConcernsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConcernsWhoseInputNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConcernsWhoseStakeholderNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConcernssWhoseOutputNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootConcerns"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Tasks"
		res = append(res, rf)
		rf.GongstructName = "Stakeholder"
		rf.Fieldname = "Concerns"
		res = append(res, rf)
	case *ConcernCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConcernComposition_Shapes"
		res = append(res, rf)
	case *ConcernInputShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConcernInputShapes"
		res = append(res, rf)
	case *ConcernOutputShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConcernOutputShapes"
		res = append(res, rf)
	case *ConcernShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Concern_Shapes"
		res = append(res, rf)
	case *Deliverable:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Concern"
		rf.Fieldname = "Inputs"
		res = append(res, rf)
		rf.GongstructName = "Concern"
		rf.Fieldname = "Outputs"
		res = append(res, rf)
		rf.GongstructName = "Deliverable"
		rf.Fieldname = "SubProducts"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ProductsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootDeliverables"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Products"
		res = append(res, rf)
	case *Diagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
	case *Note:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NotesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "Notes"
		res = append(res, rf)
	case *NoteProductShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NoteProductShapes"
		res = append(res, rf)
	case *NoteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Note_Shapes"
		res = append(res, rf)
	case *NoteStakeholderShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NoteResourceShapes"
		res = append(res, rf)
	case *NoteTaskShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NoteTaskShapes"
		res = append(res, rf)
	case *ProductCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ProductComposition_Shapes"
		res = append(res, rf)
	case *ProductShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Product_Shapes"
		res = append(res, rf)
	case *Requirement:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Concern"
		rf.Fieldname = "Requirements"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootRequirements"
		res = append(res, rf)
	case *Stakeholder:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ResourcesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootStakeholders"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Resources"
		res = append(res, rf)
		rf.GongstructName = "Stakeholder"
		rf.Fieldname = "SubStakeholders"
		res = append(res, rf)
	case *StakeholderCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ResourceComposition_Shapes"
		res = append(res, rf)
	case *StakeholderConcernShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "StakeholderConcernShapes"
		res = append(res, rf)
	case *StakeholderShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Stakeholder_Shapes"
		res = append(res, rf)
	case *SupportLevel:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Requirement"
		rf.Fieldname = "SupportLevels"
		res = append(res, rf)
	case *Tool:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Concept"
		rf.Fieldname = "Tools"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (analysisneed *AnalysisNeed) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
	}
	return
}

func (concept *Concept) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:                 "Tools",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Tool",
		},
	}
	return
}

func (concern *Concern) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IDAirbus",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Priority",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Priority",
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SubConcerns",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "Inputs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "IsInputsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Outputs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "IsOutputsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsWithCompletion",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Completion",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "CompletionEnum",
		},
		{
			Name:                 "Requirements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Requirement",
		},
	}
	return
}

func (concerncompositionshape *ConcernCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Concern",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concern",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (concerninputshape *ConcernInputShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Deliverable",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:                 "Concern",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concern",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (concernoutputshape *ConcernOutputShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (concernshape *ConcernShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Concern",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concern",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (deliverable *Deliverable) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SubProducts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "IsProducersNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsConsumersNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Concepts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concept",
		},
	}
	return
}

func (diagram *Diagram) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEditable_",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowPrefix",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "DefaultBoxWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DefaultBoxHeigth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "ConcernsWhoseRequirementsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:               "IsRequirementsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsConceptsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Product_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProductShape",
		},
		{
			Name:                 "ProductsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "IsPBSNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ProductComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProductCompositionShape",
		},
		{
			Name:               "IsConcernsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Concern_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ConcernShape",
		},
		{
			Name:                 "ConcernsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "ConcernsWhoseInputNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "ConcernsWhoseStakeholderNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "ConcernssWhoseOutputNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "ConcernComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ConcernCompositionShape",
		},
		{
			Name:                 "ConcernInputShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ConcernInputShape",
		},
		{
			Name:                 "ConcernOutputShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ConcernOutputShape",
		},
		{
			Name:                 "Note_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteShape",
		},
		{
			Name:                 "NotesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
		},
		{
			Name:               "IsNotesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "NoteProductShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteProductShape",
		},
		{
			Name:                 "NoteTaskShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteTaskShape",
		},
		{
			Name:                 "NoteResourceShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteStakeholderShape",
		},
		{
			Name:                 "Stakeholder_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StakeholderShape",
		},
		{
			Name:                 "ResourcesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Stakeholder",
		},
		{
			Name:               "IsStakeholdersNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ResourceComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StakeholderCompositionShape",
		},
		{
			Name:                 "StakeholderConcernShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StakeholderConcernShape",
		},
	}
	return
}

func (library *Library) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsRootLibrary",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:                 "RootDeliverables",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:                 "RootConcerns",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "RootStakeholders",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Stakeholder",
		},
		{
			Name:                 "RootRequirements",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Requirement",
		},
		{
			Name:                 "RootConcepts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concept",
		},
		{
			Name:                 "AnalysisNeeds",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AnalysisNeed",
		},
		{
			Name:                 "Notes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
		},
		{
			Name:                 "Diagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
		{
			Name:                 "SubLibraries",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "NbPixPerCharacter",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (note *Note) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:                 "Products",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:                 "Tasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "Resources",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Stakeholder",
		},
	}
	return
}

func (noteproductshape *NoteProductShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Note",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Note",
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (noteshape *NoteShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Note",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Note",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (notestakeholdershape *NoteStakeholderShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Note",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Note",
		},
		{
			Name:                 "Stakeholder",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Stakeholder",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (notetaskshape *NoteTaskShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Note",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Note",
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concern",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (productcompositionshape *ProductCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (productshape *ProductShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (requirement *Requirement) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:                 "SupportLevels",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SupportLevel",
		},
		{
			Name:                 "Concepts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concept",
		},
	}
	return
}

func (stakeholder *Stakeholder) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IDAirbus",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Concerns",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "SubStakeholders",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Stakeholder",
		},
	}
	return
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Stakeholder",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Stakeholder",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Stakeholder",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Stakeholder",
		},
		{
			Name:                 "Concern",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concern",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "StartOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:                 "EndOrientation",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "OrientationType",
		},
		{
			Name:               "CornerOffsetRatio",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (stakeholdershape *StakeholderShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Stakeholder",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Stakeholder",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (supportlevel *SupportLevel) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Tool",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Tool",
		},
	}
	return
}

func (tool *Tool) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt             GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeIntDuration     GongFieldValueType = "GongFieldValueTypeIntDuration"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeDate            GongFieldValueType = "GongFieldValueTypeDate"
	GongFieldValueTypeBasicKind       GongFieldValueType = "GongFieldValueTypeBasicKind"
	GongFieldValueTypePointer         GongFieldValueType = "GongFieldValueTypePointer"
	GongFieldValueTypeSliceOfPointers GongFieldValueType = "GongFieldValueTypeSliceOfPointers"
)

type GongFieldValue struct {
	GongFieldValueType
	valueString string
	valueInt    int
	valueFloat  float64
	valueBool   bool

	// in case of a pointer, the ID of the pointed element
	// in case of a slice of pointers, the IDs, separated by semi columbs
	ids string
}

type GongFieldHeader struct {
	Name string
	GongFieldValueType
	TargetGongstructName string
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

// insertion point for generic get gongstruct field value
func (analysisneed *AnalysisNeed) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = analysisneed.Name
	case "ComputedPrefix":
		res.valueString = analysisneed.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", analysisneed.IsExpanded)
		res.valueBool = analysisneed.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := analysisneed.LayoutDirection
		res.valueString = enum.ToCodeString()
	}
	return
}

func (concept *Concept) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = concept.Name
	case "ComputedPrefix":
		res.valueString = concept.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", concept.IsExpanded)
		res.valueBool = concept.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := concept.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Tools":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range concept.Tools {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (concern *Concern) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = concern.Name
	case "IDAirbus":
		res.valueString = concern.IDAirbus
	case "Priority":
		enum := concern.Priority
		res.valueString = enum.ToCodeString()
	case "ComputedPrefix":
		res.valueString = concern.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", concern.IsExpanded)
		res.valueBool = concern.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := concern.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Description":
		res.valueString = concern.Description
	case "SubConcerns":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range concern.SubConcerns {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Inputs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range concern.Inputs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsInputsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", concern.IsInputsNodeExpanded)
		res.valueBool = concern.IsInputsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Outputs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range concern.Outputs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsOutputsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", concern.IsOutputsNodeExpanded)
		res.valueBool = concern.IsOutputsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsWithCompletion":
		res.valueString = fmt.Sprintf("%t", concern.IsWithCompletion)
		res.valueBool = concern.IsWithCompletion
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Completion":
		enum := concern.Completion
		res.valueString = enum.ToCodeString()
	case "Requirements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range concern.Requirements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (concerncompositionshape *ConcernCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = concerncompositionshape.Name
	case "Concern":
		res.GongFieldValueType = GongFieldValueTypePointer
		if concerncompositionshape.Concern != nil {
			res.valueString = concerncompositionshape.Concern.Name
			res.ids = concerncompositionshape.Concern.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", concerncompositionshape.StartRatio)
		res.valueFloat = concerncompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", concerncompositionshape.EndRatio)
		res.valueFloat = concerncompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := concerncompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := concerncompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", concerncompositionshape.CornerOffsetRatio)
		res.valueFloat = concerncompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", concerncompositionshape.IsHidden)
		res.valueBool = concerncompositionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (concerninputshape *ConcernInputShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = concerninputshape.Name
	case "Deliverable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if concerninputshape.Deliverable != nil {
			res.valueString = concerninputshape.Deliverable.Name
			res.ids = concerninputshape.Deliverable.GongGetUUID(stage)
		}
	case "Concern":
		res.GongFieldValueType = GongFieldValueTypePointer
		if concerninputshape.Concern != nil {
			res.valueString = concerninputshape.Concern.Name
			res.ids = concerninputshape.Concern.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", concerninputshape.StartRatio)
		res.valueFloat = concerninputshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", concerninputshape.EndRatio)
		res.valueFloat = concerninputshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := concerninputshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := concerninputshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", concerninputshape.CornerOffsetRatio)
		res.valueFloat = concerninputshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", concerninputshape.IsHidden)
		res.valueBool = concerninputshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (concernoutputshape *ConcernOutputShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = concernoutputshape.Name
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if concernoutputshape.Task != nil {
			res.valueString = concernoutputshape.Task.Name
			res.ids = concernoutputshape.Task.GongGetUUID(stage)
		}
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if concernoutputshape.Product != nil {
			res.valueString = concernoutputshape.Product.Name
			res.ids = concernoutputshape.Product.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", concernoutputshape.StartRatio)
		res.valueFloat = concernoutputshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", concernoutputshape.EndRatio)
		res.valueFloat = concernoutputshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := concernoutputshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := concernoutputshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", concernoutputshape.CornerOffsetRatio)
		res.valueFloat = concernoutputshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", concernoutputshape.IsHidden)
		res.valueBool = concernoutputshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (concernshape *ConcernShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = concernshape.Name
	case "Concern":
		res.GongFieldValueType = GongFieldValueTypePointer
		if concernshape.Concern != nil {
			res.valueString = concernshape.Concern.Name
			res.ids = concernshape.Concern.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", concernshape.IsExpanded)
		res.valueBool = concernshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", concernshape.X)
		res.valueFloat = concernshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", concernshape.Y)
		res.valueFloat = concernshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", concernshape.Width)
		res.valueFloat = concernshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", concernshape.Height)
		res.valueFloat = concernshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", concernshape.IsHidden)
		res.valueBool = concernshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (deliverable *Deliverable) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = deliverable.Name
	case "ComputedPrefix":
		res.valueString = deliverable.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", deliverable.IsExpanded)
		res.valueBool = deliverable.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := deliverable.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Description":
		res.valueString = deliverable.Description
	case "SubProducts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range deliverable.SubProducts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsProducersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", deliverable.IsProducersNodeExpanded)
		res.valueBool = deliverable.IsProducersNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsConsumersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", deliverable.IsConsumersNodeExpanded)
		res.valueBool = deliverable.IsConsumersNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Concepts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range deliverable.Concepts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (diagram *Diagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagram.Name
	case "ComputedPrefix":
		res.valueString = diagram.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsExpanded)
		res.valueBool = diagram.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := diagram.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagram.IsChecked)
		res.valueBool = diagram.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagram.IsEditable_)
		res.valueBool = diagram.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagram.ShowPrefix)
		res.valueBool = diagram.ShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagram.DefaultBoxWidth)
		res.valueFloat = diagram.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagram.DefaultBoxHeigth)
		res.valueFloat = diagram.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagram.Width)
		res.valueFloat = diagram.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagram.Height)
		res.valueFloat = diagram.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ConcernsWhoseRequirementsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConcernsWhoseRequirementsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsRequirementsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsRequirementsNodeExpanded)
		res.valueBool = diagram.IsRequirementsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsConceptsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsConceptsNodeExpanded)
		res.valueBool = diagram.IsConceptsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Product_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Product_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ProductsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ProductsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsPBSNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsPBSNodeExpanded)
		res.valueBool = diagram.IsPBSNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ProductComposition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ProductComposition_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsConcernsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsConcernsNodeExpanded)
		res.valueBool = diagram.IsConcernsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Concern_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Concern_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConcernsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConcernsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConcernsWhoseInputNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConcernsWhoseInputNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConcernsWhoseStakeholderNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConcernsWhoseStakeholderNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConcernssWhoseOutputNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConcernssWhoseOutputNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConcernComposition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConcernComposition_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConcernInputShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConcernInputShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConcernOutputShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConcernOutputShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Note_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Note_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NotesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.NotesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsNotesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsNotesNodeExpanded)
		res.valueBool = diagram.IsNotesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NoteProductShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.NoteProductShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NoteTaskShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.NoteTaskShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NoteResourceShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.NoteResourceShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Stakeholder_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Stakeholder_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ResourcesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ResourcesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsStakeholdersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsStakeholdersNodeExpanded)
		res.valueBool = diagram.IsStakeholdersNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ResourceComposition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ResourceComposition_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "StakeholderConcernShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.StakeholderConcernShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (library *Library) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = library.Name
	case "IsRootLibrary":
		res.valueString = fmt.Sprintf("%t", library.IsRootLibrary)
		res.valueBool = library.IsRootLibrary
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsExpanded)
		res.valueBool = library.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := library.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "RootDeliverables":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootDeliverables {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootConcerns":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootConcerns {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootStakeholders":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootStakeholders {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootRequirements":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootRequirements {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootConcepts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootConcepts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AnalysisNeeds":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.AnalysisNeeds {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Notes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.Notes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.Diagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SubLibraries":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SubLibraries {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NbPixPerCharacter":
		res.valueString = fmt.Sprintf("%f", library.NbPixPerCharacter)
		res.valueFloat = library.NbPixPerCharacter
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (note *Note) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = note.Name
	case "ComputedPrefix":
		res.valueString = note.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsExpanded)
		res.valueBool = note.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := note.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Products":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Products {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Tasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Tasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Resources":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Resources {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (noteproductshape *NoteProductShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = noteproductshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteproductshape.Note != nil {
			res.valueString = noteproductshape.Note.Name
			res.ids = noteproductshape.Note.GongGetUUID(stage)
		}
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteproductshape.Product != nil {
			res.valueString = noteproductshape.Product.Name
			res.ids = noteproductshape.Product.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", noteproductshape.StartRatio)
		res.valueFloat = noteproductshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", noteproductshape.EndRatio)
		res.valueFloat = noteproductshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := noteproductshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := noteproductshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", noteproductshape.CornerOffsetRatio)
		res.valueFloat = noteproductshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", noteproductshape.IsHidden)
		res.valueBool = noteproductshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (noteshape *NoteShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = noteshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteshape.Note != nil {
			res.valueString = noteshape.Note.Name
			res.ids = noteshape.Note.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", noteshape.IsExpanded)
		res.valueBool = noteshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", noteshape.X)
		res.valueFloat = noteshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", noteshape.Y)
		res.valueFloat = noteshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", noteshape.Width)
		res.valueFloat = noteshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", noteshape.Height)
		res.valueFloat = noteshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", noteshape.IsHidden)
		res.valueBool = noteshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (notestakeholdershape *NoteStakeholderShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notestakeholdershape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notestakeholdershape.Note != nil {
			res.valueString = notestakeholdershape.Note.Name
			res.ids = notestakeholdershape.Note.GongGetUUID(stage)
		}
	case "Stakeholder":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notestakeholdershape.Stakeholder != nil {
			res.valueString = notestakeholdershape.Stakeholder.Name
			res.ids = notestakeholdershape.Stakeholder.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", notestakeholdershape.StartRatio)
		res.valueFloat = notestakeholdershape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", notestakeholdershape.EndRatio)
		res.valueFloat = notestakeholdershape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := notestakeholdershape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := notestakeholdershape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", notestakeholdershape.CornerOffsetRatio)
		res.valueFloat = notestakeholdershape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", notestakeholdershape.IsHidden)
		res.valueBool = notestakeholdershape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (notetaskshape *NoteTaskShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notetaskshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notetaskshape.Note != nil {
			res.valueString = notetaskshape.Note.Name
			res.ids = notetaskshape.Note.GongGetUUID(stage)
		}
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notetaskshape.Task != nil {
			res.valueString = notetaskshape.Task.Name
			res.ids = notetaskshape.Task.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", notetaskshape.StartRatio)
		res.valueFloat = notetaskshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", notetaskshape.EndRatio)
		res.valueFloat = notetaskshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := notetaskshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := notetaskshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", notetaskshape.CornerOffsetRatio)
		res.valueFloat = notetaskshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", notetaskshape.IsHidden)
		res.valueBool = notetaskshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (productcompositionshape *ProductCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = productcompositionshape.Name
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if productcompositionshape.Product != nil {
			res.valueString = productcompositionshape.Product.Name
			res.ids = productcompositionshape.Product.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", productcompositionshape.StartRatio)
		res.valueFloat = productcompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", productcompositionshape.EndRatio)
		res.valueFloat = productcompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := productcompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := productcompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", productcompositionshape.CornerOffsetRatio)
		res.valueFloat = productcompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", productcompositionshape.IsHidden)
		res.valueBool = productcompositionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (productshape *ProductShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = productshape.Name
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if productshape.Product != nil {
			res.valueString = productshape.Product.Name
			res.ids = productshape.Product.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", productshape.IsExpanded)
		res.valueBool = productshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", productshape.X)
		res.valueFloat = productshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", productshape.Y)
		res.valueFloat = productshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", productshape.Width)
		res.valueFloat = productshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", productshape.Height)
		res.valueFloat = productshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", productshape.IsHidden)
		res.valueBool = productshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (requirement *Requirement) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = requirement.Name
	case "ComputedPrefix":
		res.valueString = requirement.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", requirement.IsExpanded)
		res.valueBool = requirement.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := requirement.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "SupportLevels":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range requirement.SupportLevels {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Concepts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range requirement.Concepts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (stakeholder *Stakeholder) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stakeholder.Name
	case "IDAirbus":
		res.valueString = stakeholder.IDAirbus
	case "ComputedPrefix":
		res.valueString = stakeholder.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", stakeholder.IsExpanded)
		res.valueBool = stakeholder.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := stakeholder.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Description":
		res.valueString = stakeholder.Description
	case "Concerns":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stakeholder.Concerns {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SubStakeholders":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stakeholder.SubStakeholders {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	}
	return
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stakeholdercompositionshape.Name
	case "Stakeholder":
		res.GongFieldValueType = GongFieldValueTypePointer
		if stakeholdercompositionshape.Stakeholder != nil {
			res.valueString = stakeholdercompositionshape.Stakeholder.Name
			res.ids = stakeholdercompositionshape.Stakeholder.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", stakeholdercompositionshape.StartRatio)
		res.valueFloat = stakeholdercompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", stakeholdercompositionshape.EndRatio)
		res.valueFloat = stakeholdercompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := stakeholdercompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := stakeholdercompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", stakeholdercompositionshape.CornerOffsetRatio)
		res.valueFloat = stakeholdercompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", stakeholdercompositionshape.IsHidden)
		res.valueBool = stakeholdercompositionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stakeholderconcernshape.Name
	case "Stakeholder":
		res.GongFieldValueType = GongFieldValueTypePointer
		if stakeholderconcernshape.Stakeholder != nil {
			res.valueString = stakeholderconcernshape.Stakeholder.Name
			res.ids = stakeholderconcernshape.Stakeholder.GongGetUUID(stage)
		}
	case "Concern":
		res.GongFieldValueType = GongFieldValueTypePointer
		if stakeholderconcernshape.Concern != nil {
			res.valueString = stakeholderconcernshape.Concern.Name
			res.ids = stakeholderconcernshape.Concern.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", stakeholderconcernshape.StartRatio)
		res.valueFloat = stakeholderconcernshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", stakeholderconcernshape.EndRatio)
		res.valueFloat = stakeholderconcernshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := stakeholderconcernshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := stakeholderconcernshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", stakeholderconcernshape.CornerOffsetRatio)
		res.valueFloat = stakeholderconcernshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", stakeholderconcernshape.IsHidden)
		res.valueBool = stakeholderconcernshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (stakeholdershape *StakeholderShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stakeholdershape.Name
	case "Stakeholder":
		res.GongFieldValueType = GongFieldValueTypePointer
		if stakeholdershape.Stakeholder != nil {
			res.valueString = stakeholdershape.Stakeholder.Name
			res.ids = stakeholdershape.Stakeholder.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", stakeholdershape.IsExpanded)
		res.valueBool = stakeholdershape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", stakeholdershape.X)
		res.valueFloat = stakeholdershape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", stakeholdershape.Y)
		res.valueFloat = stakeholdershape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", stakeholdershape.Width)
		res.valueFloat = stakeholdershape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", stakeholdershape.Height)
		res.valueFloat = stakeholdershape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", stakeholdershape.IsHidden)
		res.valueBool = stakeholdershape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (supportlevel *SupportLevel) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = supportlevel.Name
	case "Tool":
		res.GongFieldValueType = GongFieldValueTypePointer
		if supportlevel.Tool != nil {
			res.valueString = supportlevel.Tool.Name
			res.ids = supportlevel.Tool.GongGetUUID(stage)
		}
	}
	return
}

func (tool *Tool) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = tool.Name
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (analysisneed *AnalysisNeed) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		analysisneed.Name = value.GetValueString()
	case "ComputedPrefix":
		analysisneed.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		analysisneed.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		analysisneed.LayoutDirection.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (concept *Concept) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		concept.Name = value.GetValueString()
	case "ComputedPrefix":
		concept.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		concept.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		concept.LayoutDirection.FromCodeString(value.GetValueString())
	case "Tools":
		concept.Tools = make([]*Tool, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tools {
					if stage.Tool_stagedOrder[__instance__] == uint(id) {
						concept.Tools = append(concept.Tools, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (concern *Concern) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		concern.Name = value.GetValueString()
	case "IDAirbus":
		concern.IDAirbus = value.GetValueString()
	case "Priority":
		concern.Priority.FromCodeString(value.GetValueString())
	case "ComputedPrefix":
		concern.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		concern.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		concern.LayoutDirection.FromCodeString(value.GetValueString())
	case "Description":
		concern.Description = value.GetValueString()
	case "SubConcerns":
		concern.SubConcerns = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						concern.SubConcerns = append(concern.SubConcerns, __instance__)
						break
					}
				}
			}
		}
	case "Inputs":
		concern.Inputs = make([]*Deliverable, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Deliverables {
					if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
						concern.Inputs = append(concern.Inputs, __instance__)
						break
					}
				}
			}
		}
	case "IsInputsNodeExpanded":
		concern.IsInputsNodeExpanded = value.GetValueBool()
	case "Outputs":
		concern.Outputs = make([]*Deliverable, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Deliverables {
					if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
						concern.Outputs = append(concern.Outputs, __instance__)
						break
					}
				}
			}
		}
	case "IsOutputsNodeExpanded":
		concern.IsOutputsNodeExpanded = value.GetValueBool()
	case "IsWithCompletion":
		concern.IsWithCompletion = value.GetValueBool()
	case "Completion":
		concern.Completion.FromCodeString(value.GetValueString())
	case "Requirements":
		concern.Requirements = make([]*Requirement, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Requirements {
					if stage.Requirement_stagedOrder[__instance__] == uint(id) {
						concern.Requirements = append(concern.Requirements, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (concerncompositionshape *ConcernCompositionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		concerncompositionshape.Name = value.GetValueString()
	case "Concern":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			concerncompositionshape.Concern = nil
			for __instance__ := range stage.Concerns {
				if stage.Concern_stagedOrder[__instance__] == uint(id) {
					concerncompositionshape.Concern = __instance__
					break
				}
			}
		}
	case "StartRatio":
		concerncompositionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		concerncompositionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		concerncompositionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		concerncompositionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		concerncompositionshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		concerncompositionshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (concerninputshape *ConcernInputShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		concerninputshape.Name = value.GetValueString()
	case "Deliverable":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			concerninputshape.Deliverable = nil
			for __instance__ := range stage.Deliverables {
				if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
					concerninputshape.Deliverable = __instance__
					break
				}
			}
		}
	case "Concern":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			concerninputshape.Concern = nil
			for __instance__ := range stage.Concerns {
				if stage.Concern_stagedOrder[__instance__] == uint(id) {
					concerninputshape.Concern = __instance__
					break
				}
			}
		}
	case "StartRatio":
		concerninputshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		concerninputshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		concerninputshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		concerninputshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		concerninputshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		concerninputshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (concernoutputshape *ConcernOutputShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		concernoutputshape.Name = value.GetValueString()
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			concernoutputshape.Task = nil
			for __instance__ := range stage.Concerns {
				if stage.Concern_stagedOrder[__instance__] == uint(id) {
					concernoutputshape.Task = __instance__
					break
				}
			}
		}
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			concernoutputshape.Product = nil
			for __instance__ := range stage.Deliverables {
				if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
					concernoutputshape.Product = __instance__
					break
				}
			}
		}
	case "StartRatio":
		concernoutputshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		concernoutputshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		concernoutputshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		concernoutputshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		concernoutputshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		concernoutputshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (concernshape *ConcernShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		concernshape.Name = value.GetValueString()
	case "Concern":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			concernshape.Concern = nil
			for __instance__ := range stage.Concerns {
				if stage.Concern_stagedOrder[__instance__] == uint(id) {
					concernshape.Concern = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		concernshape.IsExpanded = value.GetValueBool()
	case "X":
		concernshape.X = value.GetValueFloat()
	case "Y":
		concernshape.Y = value.GetValueFloat()
	case "Width":
		concernshape.Width = value.GetValueFloat()
	case "Height":
		concernshape.Height = value.GetValueFloat()
	case "IsHidden":
		concernshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (deliverable *Deliverable) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		deliverable.Name = value.GetValueString()
	case "ComputedPrefix":
		deliverable.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		deliverable.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		deliverable.LayoutDirection.FromCodeString(value.GetValueString())
	case "Description":
		deliverable.Description = value.GetValueString()
	case "SubProducts":
		deliverable.SubProducts = make([]*Deliverable, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Deliverables {
					if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
						deliverable.SubProducts = append(deliverable.SubProducts, __instance__)
						break
					}
				}
			}
		}
	case "IsProducersNodeExpanded":
		deliverable.IsProducersNodeExpanded = value.GetValueBool()
	case "IsConsumersNodeExpanded":
		deliverable.IsConsumersNodeExpanded = value.GetValueBool()
	case "Concepts":
		deliverable.Concepts = make([]*Concept, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concepts {
					if stage.Concept_stagedOrder[__instance__] == uint(id) {
						deliverable.Concepts = append(deliverable.Concepts, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (diagram *Diagram) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagram.Name = value.GetValueString()
	case "ComputedPrefix":
		diagram.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagram.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		diagram.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsChecked":
		diagram.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagram.IsEditable_ = value.GetValueBool()
	case "ShowPrefix":
		diagram.ShowPrefix = value.GetValueBool()
	case "DefaultBoxWidth":
		diagram.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagram.DefaultBoxHeigth = value.GetValueFloat()
	case "Width":
		diagram.Width = value.GetValueFloat()
	case "Height":
		diagram.Height = value.GetValueFloat()
	case "ConcernsWhoseRequirementsNodeIsExpanded":
		diagram.ConcernsWhoseRequirementsNodeIsExpanded = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						diagram.ConcernsWhoseRequirementsNodeIsExpanded = append(diagram.ConcernsWhoseRequirementsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsRequirementsNodeExpanded":
		diagram.IsRequirementsNodeExpanded = value.GetValueBool()
	case "IsConceptsNodeExpanded":
		diagram.IsConceptsNodeExpanded = value.GetValueBool()
	case "Product_Shapes":
		diagram.Product_Shapes = make([]*ProductShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ProductShapes {
					if stage.ProductShape_stagedOrder[__instance__] == uint(id) {
						diagram.Product_Shapes = append(diagram.Product_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ProductsWhoseNodeIsExpanded":
		diagram.ProductsWhoseNodeIsExpanded = make([]*Deliverable, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Deliverables {
					if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
						diagram.ProductsWhoseNodeIsExpanded = append(diagram.ProductsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsPBSNodeExpanded":
		diagram.IsPBSNodeExpanded = value.GetValueBool()
	case "ProductComposition_Shapes":
		diagram.ProductComposition_Shapes = make([]*ProductCompositionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ProductCompositionShapes {
					if stage.ProductCompositionShape_stagedOrder[__instance__] == uint(id) {
						diagram.ProductComposition_Shapes = append(diagram.ProductComposition_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsConcernsNodeExpanded":
		diagram.IsConcernsNodeExpanded = value.GetValueBool()
	case "Concern_Shapes":
		diagram.Concern_Shapes = make([]*ConcernShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ConcernShapes {
					if stage.ConcernShape_stagedOrder[__instance__] == uint(id) {
						diagram.Concern_Shapes = append(diagram.Concern_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ConcernsWhoseNodeIsExpanded":
		diagram.ConcernsWhoseNodeIsExpanded = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						diagram.ConcernsWhoseNodeIsExpanded = append(diagram.ConcernsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ConcernsWhoseInputNodeIsExpanded":
		diagram.ConcernsWhoseInputNodeIsExpanded = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						diagram.ConcernsWhoseInputNodeIsExpanded = append(diagram.ConcernsWhoseInputNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ConcernsWhoseStakeholderNodeIsExpanded":
		diagram.ConcernsWhoseStakeholderNodeIsExpanded = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						diagram.ConcernsWhoseStakeholderNodeIsExpanded = append(diagram.ConcernsWhoseStakeholderNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ConcernssWhoseOutputNodeIsExpanded":
		diagram.ConcernssWhoseOutputNodeIsExpanded = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						diagram.ConcernssWhoseOutputNodeIsExpanded = append(diagram.ConcernssWhoseOutputNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ConcernComposition_Shapes":
		diagram.ConcernComposition_Shapes = make([]*ConcernCompositionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ConcernCompositionShapes {
					if stage.ConcernCompositionShape_stagedOrder[__instance__] == uint(id) {
						diagram.ConcernComposition_Shapes = append(diagram.ConcernComposition_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ConcernInputShapes":
		diagram.ConcernInputShapes = make([]*ConcernInputShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ConcernInputShapes {
					if stage.ConcernInputShape_stagedOrder[__instance__] == uint(id) {
						diagram.ConcernInputShapes = append(diagram.ConcernInputShapes, __instance__)
						break
					}
				}
			}
		}
	case "ConcernOutputShapes":
		diagram.ConcernOutputShapes = make([]*ConcernOutputShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ConcernOutputShapes {
					if stage.ConcernOutputShape_stagedOrder[__instance__] == uint(id) {
						diagram.ConcernOutputShapes = append(diagram.ConcernOutputShapes, __instance__)
						break
					}
				}
			}
		}
	case "Note_Shapes":
		diagram.Note_Shapes = make([]*NoteShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteShapes {
					if stage.NoteShape_stagedOrder[__instance__] == uint(id) {
						diagram.Note_Shapes = append(diagram.Note_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "NotesWhoseNodeIsExpanded":
		diagram.NotesWhoseNodeIsExpanded = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						diagram.NotesWhoseNodeIsExpanded = append(diagram.NotesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsNotesNodeExpanded":
		diagram.IsNotesNodeExpanded = value.GetValueBool()
	case "NoteProductShapes":
		diagram.NoteProductShapes = make([]*NoteProductShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteProductShapes {
					if stage.NoteProductShape_stagedOrder[__instance__] == uint(id) {
						diagram.NoteProductShapes = append(diagram.NoteProductShapes, __instance__)
						break
					}
				}
			}
		}
	case "NoteTaskShapes":
		diagram.NoteTaskShapes = make([]*NoteTaskShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteTaskShapes {
					if stage.NoteTaskShape_stagedOrder[__instance__] == uint(id) {
						diagram.NoteTaskShapes = append(diagram.NoteTaskShapes, __instance__)
						break
					}
				}
			}
		}
	case "NoteResourceShapes":
		diagram.NoteResourceShapes = make([]*NoteStakeholderShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteStakeholderShapes {
					if stage.NoteStakeholderShape_stagedOrder[__instance__] == uint(id) {
						diagram.NoteResourceShapes = append(diagram.NoteResourceShapes, __instance__)
						break
					}
				}
			}
		}
	case "Stakeholder_Shapes":
		diagram.Stakeholder_Shapes = make([]*StakeholderShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StakeholderShapes {
					if stage.StakeholderShape_stagedOrder[__instance__] == uint(id) {
						diagram.Stakeholder_Shapes = append(diagram.Stakeholder_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ResourcesWhoseNodeIsExpanded":
		diagram.ResourcesWhoseNodeIsExpanded = make([]*Stakeholder, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Stakeholders {
					if stage.Stakeholder_stagedOrder[__instance__] == uint(id) {
						diagram.ResourcesWhoseNodeIsExpanded = append(diagram.ResourcesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsStakeholdersNodeExpanded":
		diagram.IsStakeholdersNodeExpanded = value.GetValueBool()
	case "ResourceComposition_Shapes":
		diagram.ResourceComposition_Shapes = make([]*StakeholderCompositionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StakeholderCompositionShapes {
					if stage.StakeholderCompositionShape_stagedOrder[__instance__] == uint(id) {
						diagram.ResourceComposition_Shapes = append(diagram.ResourceComposition_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "StakeholderConcernShapes":
		diagram.StakeholderConcernShapes = make([]*StakeholderConcernShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StakeholderConcernShapes {
					if stage.StakeholderConcernShape_stagedOrder[__instance__] == uint(id) {
						diagram.StakeholderConcernShapes = append(diagram.StakeholderConcernShapes, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (library *Library) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		library.Name = value.GetValueString()
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		library.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		library.LayoutDirection.FromCodeString(value.GetValueString())
	case "RootDeliverables":
		library.RootDeliverables = make([]*Deliverable, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Deliverables {
					if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
						library.RootDeliverables = append(library.RootDeliverables, __instance__)
						break
					}
				}
			}
		}
	case "RootConcerns":
		library.RootConcerns = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						library.RootConcerns = append(library.RootConcerns, __instance__)
						break
					}
				}
			}
		}
	case "RootStakeholders":
		library.RootStakeholders = make([]*Stakeholder, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Stakeholders {
					if stage.Stakeholder_stagedOrder[__instance__] == uint(id) {
						library.RootStakeholders = append(library.RootStakeholders, __instance__)
						break
					}
				}
			}
		}
	case "RootRequirements":
		library.RootRequirements = make([]*Requirement, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Requirements {
					if stage.Requirement_stagedOrder[__instance__] == uint(id) {
						library.RootRequirements = append(library.RootRequirements, __instance__)
						break
					}
				}
			}
		}
	case "RootConcepts":
		library.RootConcepts = make([]*Concept, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concepts {
					if stage.Concept_stagedOrder[__instance__] == uint(id) {
						library.RootConcepts = append(library.RootConcepts, __instance__)
						break
					}
				}
			}
		}
	case "AnalysisNeeds":
		library.AnalysisNeeds = make([]*AnalysisNeed, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AnalysisNeeds {
					if stage.AnalysisNeed_stagedOrder[__instance__] == uint(id) {
						library.AnalysisNeeds = append(library.AnalysisNeeds, __instance__)
						break
					}
				}
			}
		}
	case "Notes":
		library.Notes = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						library.Notes = append(library.Notes, __instance__)
						break
					}
				}
			}
		}
	case "Diagrams":
		library.Diagrams = make([]*Diagram, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Diagrams {
					if stage.Diagram_stagedOrder[__instance__] == uint(id) {
						library.Diagrams = append(library.Diagrams, __instance__)
						break
					}
				}
			}
		}
	case "SubLibraries":
		library.SubLibraries = make([]*Library, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Librarys {
					if stage.Library_stagedOrder[__instance__] == uint(id) {
						library.SubLibraries = append(library.SubLibraries, __instance__)
						break
					}
				}
			}
		}
	case "NbPixPerCharacter":
		library.NbPixPerCharacter = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (note *Note) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		note.Name = value.GetValueString()
	case "ComputedPrefix":
		note.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		note.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		note.LayoutDirection.FromCodeString(value.GetValueString())
	case "Products":
		note.Products = make([]*Deliverable, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Deliverables {
					if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
						note.Products = append(note.Products, __instance__)
						break
					}
				}
			}
		}
	case "Tasks":
		note.Tasks = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						note.Tasks = append(note.Tasks, __instance__)
						break
					}
				}
			}
		}
	case "Resources":
		note.Resources = make([]*Stakeholder, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Stakeholders {
					if stage.Stakeholder_stagedOrder[__instance__] == uint(id) {
						note.Resources = append(note.Resources, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (noteproductshape *NoteProductShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		noteproductshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteproductshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					noteproductshape.Note = __instance__
					break
				}
			}
		}
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteproductshape.Product = nil
			for __instance__ := range stage.Deliverables {
				if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
					noteproductshape.Product = __instance__
					break
				}
			}
		}
	case "StartRatio":
		noteproductshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		noteproductshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		noteproductshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		noteproductshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		noteproductshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		noteproductshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (noteshape *NoteShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		noteshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					noteshape.Note = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		noteshape.IsExpanded = value.GetValueBool()
	case "X":
		noteshape.X = value.GetValueFloat()
	case "Y":
		noteshape.Y = value.GetValueFloat()
	case "Width":
		noteshape.Width = value.GetValueFloat()
	case "Height":
		noteshape.Height = value.GetValueFloat()
	case "IsHidden":
		noteshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (notestakeholdershape *NoteStakeholderShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		notestakeholdershape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notestakeholdershape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					notestakeholdershape.Note = __instance__
					break
				}
			}
		}
	case "Stakeholder":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notestakeholdershape.Stakeholder = nil
			for __instance__ := range stage.Stakeholders {
				if stage.Stakeholder_stagedOrder[__instance__] == uint(id) {
					notestakeholdershape.Stakeholder = __instance__
					break
				}
			}
		}
	case "StartRatio":
		notestakeholdershape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		notestakeholdershape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		notestakeholdershape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		notestakeholdershape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		notestakeholdershape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		notestakeholdershape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (notetaskshape *NoteTaskShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		notetaskshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notetaskshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					notetaskshape.Note = __instance__
					break
				}
			}
		}
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notetaskshape.Task = nil
			for __instance__ := range stage.Concerns {
				if stage.Concern_stagedOrder[__instance__] == uint(id) {
					notetaskshape.Task = __instance__
					break
				}
			}
		}
	case "StartRatio":
		notetaskshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		notetaskshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		notetaskshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		notetaskshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		notetaskshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		notetaskshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (productcompositionshape *ProductCompositionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		productcompositionshape.Name = value.GetValueString()
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			productcompositionshape.Product = nil
			for __instance__ := range stage.Deliverables {
				if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
					productcompositionshape.Product = __instance__
					break
				}
			}
		}
	case "StartRatio":
		productcompositionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		productcompositionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		productcompositionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		productcompositionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		productcompositionshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		productcompositionshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (productshape *ProductShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		productshape.Name = value.GetValueString()
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			productshape.Product = nil
			for __instance__ := range stage.Deliverables {
				if stage.Deliverable_stagedOrder[__instance__] == uint(id) {
					productshape.Product = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		productshape.IsExpanded = value.GetValueBool()
	case "X":
		productshape.X = value.GetValueFloat()
	case "Y":
		productshape.Y = value.GetValueFloat()
	case "Width":
		productshape.Width = value.GetValueFloat()
	case "Height":
		productshape.Height = value.GetValueFloat()
	case "IsHidden":
		productshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (requirement *Requirement) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		requirement.Name = value.GetValueString()
	case "ComputedPrefix":
		requirement.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		requirement.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		requirement.LayoutDirection.FromCodeString(value.GetValueString())
	case "SupportLevels":
		requirement.SupportLevels = make([]*SupportLevel, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SupportLevels {
					if stage.SupportLevel_stagedOrder[__instance__] == uint(id) {
						requirement.SupportLevels = append(requirement.SupportLevels, __instance__)
						break
					}
				}
			}
		}
	case "Concepts":
		requirement.Concepts = make([]*Concept, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concepts {
					if stage.Concept_stagedOrder[__instance__] == uint(id) {
						requirement.Concepts = append(requirement.Concepts, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (stakeholder *Stakeholder) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stakeholder.Name = value.GetValueString()
	case "IDAirbus":
		stakeholder.IDAirbus = value.GetValueString()
	case "ComputedPrefix":
		stakeholder.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		stakeholder.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		stakeholder.LayoutDirection.FromCodeString(value.GetValueString())
	case "Description":
		stakeholder.Description = value.GetValueString()
	case "Concerns":
		stakeholder.Concerns = make([]*Concern, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Concerns {
					if stage.Concern_stagedOrder[__instance__] == uint(id) {
						stakeholder.Concerns = append(stakeholder.Concerns, __instance__)
						break
					}
				}
			}
		}
	case "SubStakeholders":
		stakeholder.SubStakeholders = make([]*Stakeholder, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Stakeholders {
					if stage.Stakeholder_stagedOrder[__instance__] == uint(id) {
						stakeholder.SubStakeholders = append(stakeholder.SubStakeholders, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stakeholdercompositionshape.Name = value.GetValueString()
	case "Stakeholder":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			stakeholdercompositionshape.Stakeholder = nil
			for __instance__ := range stage.Stakeholders {
				if stage.Stakeholder_stagedOrder[__instance__] == uint(id) {
					stakeholdercompositionshape.Stakeholder = __instance__
					break
				}
			}
		}
	case "StartRatio":
		stakeholdercompositionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		stakeholdercompositionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		stakeholdercompositionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		stakeholdercompositionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		stakeholdercompositionshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		stakeholdercompositionshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (stakeholderconcernshape *StakeholderConcernShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stakeholderconcernshape.Name = value.GetValueString()
	case "Stakeholder":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			stakeholderconcernshape.Stakeholder = nil
			for __instance__ := range stage.Stakeholders {
				if stage.Stakeholder_stagedOrder[__instance__] == uint(id) {
					stakeholderconcernshape.Stakeholder = __instance__
					break
				}
			}
		}
	case "Concern":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			stakeholderconcernshape.Concern = nil
			for __instance__ := range stage.Concerns {
				if stage.Concern_stagedOrder[__instance__] == uint(id) {
					stakeholderconcernshape.Concern = __instance__
					break
				}
			}
		}
	case "StartRatio":
		stakeholderconcernshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		stakeholderconcernshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		stakeholderconcernshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		stakeholderconcernshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		stakeholderconcernshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		stakeholderconcernshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (stakeholdershape *StakeholderShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stakeholdershape.Name = value.GetValueString()
	case "Stakeholder":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			stakeholdershape.Stakeholder = nil
			for __instance__ := range stage.Stakeholders {
				if stage.Stakeholder_stagedOrder[__instance__] == uint(id) {
					stakeholdershape.Stakeholder = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		stakeholdershape.IsExpanded = value.GetValueBool()
	case "X":
		stakeholdershape.X = value.GetValueFloat()
	case "Y":
		stakeholdershape.Y = value.GetValueFloat()
	case "Width":
		stakeholdershape.Width = value.GetValueFloat()
	case "Height":
		stakeholdershape.Height = value.GetValueFloat()
	case "IsHidden":
		stakeholdershape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (supportlevel *SupportLevel) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		supportlevel.Name = value.GetValueString()
	case "Tool":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			supportlevel.Tool = nil
			for __instance__ := range stage.Tools {
				if stage.Tool_stagedOrder[__instance__] == uint(id) {
					supportlevel.Tool = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (tool *Tool) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		tool.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (analysisneed *AnalysisNeed) GongGetGongstructName() string {
	return "AnalysisNeed"
}

func (concept *Concept) GongGetGongstructName() string {
	return "Concept"
}

func (concern *Concern) GongGetGongstructName() string {
	return "Concern"
}

func (concerncompositionshape *ConcernCompositionShape) GongGetGongstructName() string {
	return "ConcernCompositionShape"
}

func (concerninputshape *ConcernInputShape) GongGetGongstructName() string {
	return "ConcernInputShape"
}

func (concernoutputshape *ConcernOutputShape) GongGetGongstructName() string {
	return "ConcernOutputShape"
}

func (concernshape *ConcernShape) GongGetGongstructName() string {
	return "ConcernShape"
}

func (deliverable *Deliverable) GongGetGongstructName() string {
	return "Deliverable"
}

func (diagram *Diagram) GongGetGongstructName() string {
	return "Diagram"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (note *Note) GongGetGongstructName() string {
	return "Note"
}

func (noteproductshape *NoteProductShape) GongGetGongstructName() string {
	return "NoteProductShape"
}

func (noteshape *NoteShape) GongGetGongstructName() string {
	return "NoteShape"
}

func (notestakeholdershape *NoteStakeholderShape) GongGetGongstructName() string {
	return "NoteStakeholderShape"
}

func (notetaskshape *NoteTaskShape) GongGetGongstructName() string {
	return "NoteTaskShape"
}

func (productcompositionshape *ProductCompositionShape) GongGetGongstructName() string {
	return "ProductCompositionShape"
}

func (productshape *ProductShape) GongGetGongstructName() string {
	return "ProductShape"
}

func (requirement *Requirement) GongGetGongstructName() string {
	return "Requirement"
}

func (stakeholder *Stakeholder) GongGetGongstructName() string {
	return "Stakeholder"
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongGetGongstructName() string {
	return "StakeholderCompositionShape"
}

func (stakeholderconcernshape *StakeholderConcernShape) GongGetGongstructName() string {
	return "StakeholderConcernShape"
}

func (stakeholdershape *StakeholderShape) GongGetGongstructName() string {
	return "StakeholderShape"
}

func (supportlevel *SupportLevel) GongGetGongstructName() string {
	return "SupportLevel"
}

func (tool *Tool) GongGetGongstructName() string {
	return "Tool"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.AnalysisNeeds_mapString = make(map[string]*AnalysisNeed)
	for analysisneed := range stage.AnalysisNeeds {
		stage.AnalysisNeeds_mapString[analysisneed.Name] = analysisneed
	}

	stage.Concepts_mapString = make(map[string]*Concept)
	for concept := range stage.Concepts {
		stage.Concepts_mapString[concept.Name] = concept
	}

	stage.Concerns_mapString = make(map[string]*Concern)
	for concern := range stage.Concerns {
		stage.Concerns_mapString[concern.Name] = concern
	}

	stage.ConcernCompositionShapes_mapString = make(map[string]*ConcernCompositionShape)
	for concerncompositionshape := range stage.ConcernCompositionShapes {
		stage.ConcernCompositionShapes_mapString[concerncompositionshape.Name] = concerncompositionshape
	}

	stage.ConcernInputShapes_mapString = make(map[string]*ConcernInputShape)
	for concerninputshape := range stage.ConcernInputShapes {
		stage.ConcernInputShapes_mapString[concerninputshape.Name] = concerninputshape
	}

	stage.ConcernOutputShapes_mapString = make(map[string]*ConcernOutputShape)
	for concernoutputshape := range stage.ConcernOutputShapes {
		stage.ConcernOutputShapes_mapString[concernoutputshape.Name] = concernoutputshape
	}

	stage.ConcernShapes_mapString = make(map[string]*ConcernShape)
	for concernshape := range stage.ConcernShapes {
		stage.ConcernShapes_mapString[concernshape.Name] = concernshape
	}

	stage.Deliverables_mapString = make(map[string]*Deliverable)
	for deliverable := range stage.Deliverables {
		stage.Deliverables_mapString[deliverable.Name] = deliverable
	}

	stage.Diagrams_mapString = make(map[string]*Diagram)
	for diagram := range stage.Diagrams {
		stage.Diagrams_mapString[diagram.Name] = diagram
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Notes_mapString = make(map[string]*Note)
	for note := range stage.Notes {
		stage.Notes_mapString[note.Name] = note
	}

	stage.NoteProductShapes_mapString = make(map[string]*NoteProductShape)
	for noteproductshape := range stage.NoteProductShapes {
		stage.NoteProductShapes_mapString[noteproductshape.Name] = noteproductshape
	}

	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	for noteshape := range stage.NoteShapes {
		stage.NoteShapes_mapString[noteshape.Name] = noteshape
	}

	stage.NoteStakeholderShapes_mapString = make(map[string]*NoteStakeholderShape)
	for notestakeholdershape := range stage.NoteStakeholderShapes {
		stage.NoteStakeholderShapes_mapString[notestakeholdershape.Name] = notestakeholdershape
	}

	stage.NoteTaskShapes_mapString = make(map[string]*NoteTaskShape)
	for notetaskshape := range stage.NoteTaskShapes {
		stage.NoteTaskShapes_mapString[notetaskshape.Name] = notetaskshape
	}

	stage.ProductCompositionShapes_mapString = make(map[string]*ProductCompositionShape)
	for productcompositionshape := range stage.ProductCompositionShapes {
		stage.ProductCompositionShapes_mapString[productcompositionshape.Name] = productcompositionshape
	}

	stage.ProductShapes_mapString = make(map[string]*ProductShape)
	for productshape := range stage.ProductShapes {
		stage.ProductShapes_mapString[productshape.Name] = productshape
	}

	stage.Requirements_mapString = make(map[string]*Requirement)
	for requirement := range stage.Requirements {
		stage.Requirements_mapString[requirement.Name] = requirement
	}

	stage.Stakeholders_mapString = make(map[string]*Stakeholder)
	for stakeholder := range stage.Stakeholders {
		stage.Stakeholders_mapString[stakeholder.Name] = stakeholder
	}

	stage.StakeholderCompositionShapes_mapString = make(map[string]*StakeholderCompositionShape)
	for stakeholdercompositionshape := range stage.StakeholderCompositionShapes {
		stage.StakeholderCompositionShapes_mapString[stakeholdercompositionshape.Name] = stakeholdercompositionshape
	}

	stage.StakeholderConcernShapes_mapString = make(map[string]*StakeholderConcernShape)
	for stakeholderconcernshape := range stage.StakeholderConcernShapes {
		stage.StakeholderConcernShapes_mapString[stakeholderconcernshape.Name] = stakeholderconcernshape
	}

	stage.StakeholderShapes_mapString = make(map[string]*StakeholderShape)
	for stakeholdershape := range stage.StakeholderShapes {
		stage.StakeholderShapes_mapString[stakeholdershape.Name] = stakeholdershape
	}

	stage.SupportLevels_mapString = make(map[string]*SupportLevel)
	for supportlevel := range stage.SupportLevels {
		stage.SupportLevels_mapString[supportlevel.Name] = supportlevel
	}

	stage.Tools_mapString = make(map[string]*Tool)
	for tool := range stage.Tools {
		stage.Tools_mapString[tool.Name] = tool
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
