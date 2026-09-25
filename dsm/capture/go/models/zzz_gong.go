// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"
)

// can be used for
//
//	days := __gong__abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __gong__abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __gong__abs
	_ = strings.Clone("")
)

const (
	GongProbeTreeSidebarSuffix           = ":sidebar of the probe"
	GongProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	GongProbeTableSuffix                 = ":table of the probe"
	GongProbeNotificationTableSuffix     = ":notification table of the probe"
	GongProbeFormSuffix                  = ":form of the probe"
	GongProbeSplitSuffix                 = ":probe of the probe"
	GongProbeLoadSuffix                  = ":load of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeSplitSuffix
}

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var _ = fmt.Sprintf

// idem for math package when not need by generated code
var _ = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// MetaPackageImport represents a package import needed by a meta/diagram file
type MetaPackageImport struct {
	Alias string
	Path  string
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
	OnAfterAnalysisNeedCreateCallback GongOnAfterCreateInterface[AnalysisNeed]
	OnAfterAnalysisNeedUpdateCallback GongOnAfterUpdateInterface[AnalysisNeed]
	OnAfterAnalysisNeedDeleteCallback GongOnAfterDeleteInterface[AnalysisNeed]

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

	OnAfterConceptCreateCallback GongOnAfterCreateInterface[Concept]
	OnAfterConceptUpdateCallback GongOnAfterUpdateInterface[Concept]
	OnAfterConceptDeleteCallback GongOnAfterDeleteInterface[Concept]

	ConceptShapes                map[*ConceptShape]struct{}
	ConceptShapes_instance       map[*ConceptShape]*ConceptShape
	ConceptShapes_mapString      map[string]*ConceptShape
	ConceptShapeOrder            uint
	ConceptShape_stagedOrder     map[*ConceptShape]uint
	ConceptShape_orderStaged     map[uint]*ConceptShape
	ConceptShapes_reference      map[*ConceptShape]*ConceptShape
	ConceptShapes_referenceOrder map[*ConceptShape]uint

	// insertion point for slice of pointers maps
	OnAfterConceptShapeCreateCallback GongOnAfterCreateInterface[ConceptShape]
	OnAfterConceptShapeUpdateCallback GongOnAfterUpdateInterface[ConceptShape]
	OnAfterConceptShapeDeleteCallback GongOnAfterDeleteInterface[ConceptShape]

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

	OnAfterConcernCreateCallback GongOnAfterCreateInterface[Concern]
	OnAfterConcernUpdateCallback GongOnAfterUpdateInterface[Concern]
	OnAfterConcernDeleteCallback GongOnAfterDeleteInterface[Concern]

	ConcernCompositionShapes                map[*ConcernCompositionShape]struct{}
	ConcernCompositionShapes_instance       map[*ConcernCompositionShape]*ConcernCompositionShape
	ConcernCompositionShapes_mapString      map[string]*ConcernCompositionShape
	ConcernCompositionShapeOrder            uint
	ConcernCompositionShape_stagedOrder     map[*ConcernCompositionShape]uint
	ConcernCompositionShape_orderStaged     map[uint]*ConcernCompositionShape
	ConcernCompositionShapes_reference      map[*ConcernCompositionShape]*ConcernCompositionShape
	ConcernCompositionShapes_referenceOrder map[*ConcernCompositionShape]uint

	// insertion point for slice of pointers maps
	ConcernCompositionShape_ControlPointShapes_reverseMap map[*ControlPointShape]*ConcernCompositionShape

	OnAfterConcernCompositionShapeCreateCallback GongOnAfterCreateInterface[ConcernCompositionShape]
	OnAfterConcernCompositionShapeUpdateCallback GongOnAfterUpdateInterface[ConcernCompositionShape]
	OnAfterConcernCompositionShapeDeleteCallback GongOnAfterDeleteInterface[ConcernCompositionShape]

	ConcernInputShapes                map[*ConcernInputShape]struct{}
	ConcernInputShapes_instance       map[*ConcernInputShape]*ConcernInputShape
	ConcernInputShapes_mapString      map[string]*ConcernInputShape
	ConcernInputShapeOrder            uint
	ConcernInputShape_stagedOrder     map[*ConcernInputShape]uint
	ConcernInputShape_orderStaged     map[uint]*ConcernInputShape
	ConcernInputShapes_reference      map[*ConcernInputShape]*ConcernInputShape
	ConcernInputShapes_referenceOrder map[*ConcernInputShape]uint

	// insertion point for slice of pointers maps
	ConcernInputShape_ControlPointShapes_reverseMap map[*ControlPointShape]*ConcernInputShape

	OnAfterConcernInputShapeCreateCallback GongOnAfterCreateInterface[ConcernInputShape]
	OnAfterConcernInputShapeUpdateCallback GongOnAfterUpdateInterface[ConcernInputShape]
	OnAfterConcernInputShapeDeleteCallback GongOnAfterDeleteInterface[ConcernInputShape]

	ConcernOutputShapes                map[*ConcernOutputShape]struct{}
	ConcernOutputShapes_instance       map[*ConcernOutputShape]*ConcernOutputShape
	ConcernOutputShapes_mapString      map[string]*ConcernOutputShape
	ConcernOutputShapeOrder            uint
	ConcernOutputShape_stagedOrder     map[*ConcernOutputShape]uint
	ConcernOutputShape_orderStaged     map[uint]*ConcernOutputShape
	ConcernOutputShapes_reference      map[*ConcernOutputShape]*ConcernOutputShape
	ConcernOutputShapes_referenceOrder map[*ConcernOutputShape]uint

	// insertion point for slice of pointers maps
	ConcernOutputShape_ControlPointShapes_reverseMap map[*ControlPointShape]*ConcernOutputShape

	OnAfterConcernOutputShapeCreateCallback GongOnAfterCreateInterface[ConcernOutputShape]
	OnAfterConcernOutputShapeUpdateCallback GongOnAfterUpdateInterface[ConcernOutputShape]
	OnAfterConcernOutputShapeDeleteCallback GongOnAfterDeleteInterface[ConcernOutputShape]

	ConcernShapes                map[*ConcernShape]struct{}
	ConcernShapes_instance       map[*ConcernShape]*ConcernShape
	ConcernShapes_mapString      map[string]*ConcernShape
	ConcernShapeOrder            uint
	ConcernShape_stagedOrder     map[*ConcernShape]uint
	ConcernShape_orderStaged     map[uint]*ConcernShape
	ConcernShapes_reference      map[*ConcernShape]*ConcernShape
	ConcernShapes_referenceOrder map[*ConcernShape]uint

	// insertion point for slice of pointers maps
	OnAfterConcernShapeCreateCallback GongOnAfterCreateInterface[ConcernShape]
	OnAfterConcernShapeUpdateCallback GongOnAfterUpdateInterface[ConcernShape]
	OnAfterConcernShapeDeleteCallback GongOnAfterDeleteInterface[ConcernShape]

	ControlPointShapes                map[*ControlPointShape]struct{}
	ControlPointShapes_instance       map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_mapString      map[string]*ControlPointShape
	ControlPointShapeOrder            uint
	ControlPointShape_stagedOrder     map[*ControlPointShape]uint
	ControlPointShape_orderStaged     map[uint]*ControlPointShape
	ControlPointShapes_reference      map[*ControlPointShape]*ControlPointShape
	ControlPointShapes_referenceOrder map[*ControlPointShape]uint

	// insertion point for slice of pointers maps
	OnAfterControlPointShapeCreateCallback GongOnAfterCreateInterface[ControlPointShape]
	OnAfterControlPointShapeUpdateCallback GongOnAfterUpdateInterface[ControlPointShape]
	OnAfterControlPointShapeDeleteCallback GongOnAfterDeleteInterface[ControlPointShape]

	Deliverables                map[*Deliverable]struct{}
	Deliverables_instance       map[*Deliverable]*Deliverable
	Deliverables_mapString      map[string]*Deliverable
	DeliverableOrder            uint
	Deliverable_stagedOrder     map[*Deliverable]uint
	Deliverable_orderStaged     map[uint]*Deliverable
	Deliverables_reference      map[*Deliverable]*Deliverable
	Deliverables_referenceOrder map[*Deliverable]uint

	// insertion point for slice of pointers maps
	Deliverable_SubDeliverables_reverseMap map[*Deliverable]*Deliverable

	Deliverable_Concepts_reverseMap map[*Concept]*Deliverable

	OnAfterDeliverableCreateCallback GongOnAfterCreateInterface[Deliverable]
	OnAfterDeliverableUpdateCallback GongOnAfterUpdateInterface[Deliverable]
	OnAfterDeliverableDeleteCallback GongOnAfterDeleteInterface[Deliverable]

	DeliverableCompositionShapes                map[*DeliverableCompositionShape]struct{}
	DeliverableCompositionShapes_instance       map[*DeliverableCompositionShape]*DeliverableCompositionShape
	DeliverableCompositionShapes_mapString      map[string]*DeliverableCompositionShape
	DeliverableCompositionShapeOrder            uint
	DeliverableCompositionShape_stagedOrder     map[*DeliverableCompositionShape]uint
	DeliverableCompositionShape_orderStaged     map[uint]*DeliverableCompositionShape
	DeliverableCompositionShapes_reference      map[*DeliverableCompositionShape]*DeliverableCompositionShape
	DeliverableCompositionShapes_referenceOrder map[*DeliverableCompositionShape]uint

	// insertion point for slice of pointers maps
	DeliverableCompositionShape_ControlPointShapes_reverseMap map[*ControlPointShape]*DeliverableCompositionShape

	OnAfterDeliverableCompositionShapeCreateCallback GongOnAfterCreateInterface[DeliverableCompositionShape]
	OnAfterDeliverableCompositionShapeUpdateCallback GongOnAfterUpdateInterface[DeliverableCompositionShape]
	OnAfterDeliverableCompositionShapeDeleteCallback GongOnAfterDeleteInterface[DeliverableCompositionShape]

	DeliverableConceptShapes                map[*DeliverableConceptShape]struct{}
	DeliverableConceptShapes_instance       map[*DeliverableConceptShape]*DeliverableConceptShape
	DeliverableConceptShapes_mapString      map[string]*DeliverableConceptShape
	DeliverableConceptShapeOrder            uint
	DeliverableConceptShape_stagedOrder     map[*DeliverableConceptShape]uint
	DeliverableConceptShape_orderStaged     map[uint]*DeliverableConceptShape
	DeliverableConceptShapes_reference      map[*DeliverableConceptShape]*DeliverableConceptShape
	DeliverableConceptShapes_referenceOrder map[*DeliverableConceptShape]uint

	// insertion point for slice of pointers maps
	DeliverableConceptShape_ControlPointShapes_reverseMap map[*ControlPointShape]*DeliverableConceptShape

	OnAfterDeliverableConceptShapeCreateCallback GongOnAfterCreateInterface[DeliverableConceptShape]
	OnAfterDeliverableConceptShapeUpdateCallback GongOnAfterUpdateInterface[DeliverableConceptShape]
	OnAfterDeliverableConceptShapeDeleteCallback GongOnAfterDeleteInterface[DeliverableConceptShape]

	DeliverableShapes                map[*DeliverableShape]struct{}
	DeliverableShapes_instance       map[*DeliverableShape]*DeliverableShape
	DeliverableShapes_mapString      map[string]*DeliverableShape
	DeliverableShapeOrder            uint
	DeliverableShape_stagedOrder     map[*DeliverableShape]uint
	DeliverableShape_orderStaged     map[uint]*DeliverableShape
	DeliverableShapes_reference      map[*DeliverableShape]*DeliverableShape
	DeliverableShapes_referenceOrder map[*DeliverableShape]uint

	// insertion point for slice of pointers maps
	OnAfterDeliverableShapeCreateCallback GongOnAfterCreateInterface[DeliverableShape]
	OnAfterDeliverableShapeUpdateCallback GongOnAfterUpdateInterface[DeliverableShape]
	OnAfterDeliverableShapeDeleteCallback GongOnAfterDeleteInterface[DeliverableShape]

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

	Diagram_Deliverable_Shapes_reverseMap map[*DeliverableShape]*Diagram

	Diagram_DeliverablesWhoseNodeIsExpanded_reverseMap map[*Deliverable]*Diagram

	Diagram_DeliverablesWhoseConceptsNodeIsExpanded_reverseMap map[*Deliverable]*Diagram

	Diagram_DeliverableComposition_Shapes_reverseMap map[*DeliverableCompositionShape]*Diagram

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

	Diagram_NoteDeliverableShapes_reverseMap map[*NoteDeliverableShape]*Diagram

	Diagram_NoteTaskShapes_reverseMap map[*NoteTaskShape]*Diagram

	Diagram_NoteResourceShapes_reverseMap map[*NoteStakeholderShape]*Diagram

	Diagram_Stakeholder_Shapes_reverseMap map[*StakeholderShape]*Diagram

	Diagram_ResourcesWhoseNodeIsExpanded_reverseMap map[*Stakeholder]*Diagram

	Diagram_ResourceComposition_Shapes_reverseMap map[*StakeholderCompositionShape]*Diagram

	Diagram_StakeholderConcernShapes_reverseMap map[*StakeholderConcernShape]*Diagram

	Diagram_Requirement_Shapes_reverseMap map[*RequirementShape]*Diagram

	Diagram_RequirementsWhoseNodeIsExpanded_reverseMap map[*Requirement]*Diagram

	Diagram_Concept_Shapes_reverseMap map[*ConceptShape]*Diagram

	Diagram_ConceptsWhoseNodeIsExpanded_reverseMap map[*Concept]*Diagram

	Diagram_ConceptsWhoseDeliverablesNodeIsExpanded_reverseMap map[*Concept]*Diagram

	Diagram_DeliverableConceptShapes_reverseMap map[*DeliverableConceptShape]*Diagram

	Diagram_Diagram_Shapes_reverseMap map[*DiagramShape]*Diagram

	Diagram_DiagramsWhoseNodeIsExpanded_reverseMap map[*Diagram]*Diagram

	OnAfterDiagramCreateCallback GongOnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback GongOnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback GongOnAfterDeleteInterface[Diagram]

	DiagramShapes                map[*DiagramShape]struct{}
	DiagramShapes_instance       map[*DiagramShape]*DiagramShape
	DiagramShapes_mapString      map[string]*DiagramShape
	DiagramShapeOrder            uint
	DiagramShape_stagedOrder     map[*DiagramShape]uint
	DiagramShape_orderStaged     map[uint]*DiagramShape
	DiagramShapes_reference      map[*DiagramShape]*DiagramShape
	DiagramShapes_referenceOrder map[*DiagramShape]uint

	// insertion point for slice of pointers maps
	OnAfterDiagramShapeCreateCallback GongOnAfterCreateInterface[DiagramShape]
	OnAfterDiagramShapeUpdateCallback GongOnAfterUpdateInterface[DiagramShape]
	OnAfterDiagramShapeDeleteCallback GongOnAfterDeleteInterface[DiagramShape]

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

	OnAfterLibraryCreateCallback GongOnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback GongOnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback GongOnAfterDeleteInterface[Library]

	Notes                map[*Note]struct{}
	Notes_instance       map[*Note]*Note
	Notes_mapString      map[string]*Note
	NoteOrder            uint
	Note_stagedOrder     map[*Note]uint
	Note_orderStaged     map[uint]*Note
	Notes_reference      map[*Note]*Note
	Notes_referenceOrder map[*Note]uint

	// insertion point for slice of pointers maps
	Note_Deliverables_reverseMap map[*Deliverable]*Note

	Note_Tasks_reverseMap map[*Concern]*Note

	Note_Resources_reverseMap map[*Stakeholder]*Note

	OnAfterNoteCreateCallback GongOnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback GongOnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback GongOnAfterDeleteInterface[Note]

	NoteDeliverableShapes                map[*NoteDeliverableShape]struct{}
	NoteDeliverableShapes_instance       map[*NoteDeliverableShape]*NoteDeliverableShape
	NoteDeliverableShapes_mapString      map[string]*NoteDeliverableShape
	NoteDeliverableShapeOrder            uint
	NoteDeliverableShape_stagedOrder     map[*NoteDeliverableShape]uint
	NoteDeliverableShape_orderStaged     map[uint]*NoteDeliverableShape
	NoteDeliverableShapes_reference      map[*NoteDeliverableShape]*NoteDeliverableShape
	NoteDeliverableShapes_referenceOrder map[*NoteDeliverableShape]uint

	// insertion point for slice of pointers maps
	NoteDeliverableShape_ControlPointShapes_reverseMap map[*ControlPointShape]*NoteDeliverableShape

	OnAfterNoteDeliverableShapeCreateCallback GongOnAfterCreateInterface[NoteDeliverableShape]
	OnAfterNoteDeliverableShapeUpdateCallback GongOnAfterUpdateInterface[NoteDeliverableShape]
	OnAfterNoteDeliverableShapeDeleteCallback GongOnAfterDeleteInterface[NoteDeliverableShape]

	NoteShapes                map[*NoteShape]struct{}
	NoteShapes_instance       map[*NoteShape]*NoteShape
	NoteShapes_mapString      map[string]*NoteShape
	NoteShapeOrder            uint
	NoteShape_stagedOrder     map[*NoteShape]uint
	NoteShape_orderStaged     map[uint]*NoteShape
	NoteShapes_reference      map[*NoteShape]*NoteShape
	NoteShapes_referenceOrder map[*NoteShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteShapeCreateCallback GongOnAfterCreateInterface[NoteShape]
	OnAfterNoteShapeUpdateCallback GongOnAfterUpdateInterface[NoteShape]
	OnAfterNoteShapeDeleteCallback GongOnAfterDeleteInterface[NoteShape]

	NoteStakeholderShapes                map[*NoteStakeholderShape]struct{}
	NoteStakeholderShapes_instance       map[*NoteStakeholderShape]*NoteStakeholderShape
	NoteStakeholderShapes_mapString      map[string]*NoteStakeholderShape
	NoteStakeholderShapeOrder            uint
	NoteStakeholderShape_stagedOrder     map[*NoteStakeholderShape]uint
	NoteStakeholderShape_orderStaged     map[uint]*NoteStakeholderShape
	NoteStakeholderShapes_reference      map[*NoteStakeholderShape]*NoteStakeholderShape
	NoteStakeholderShapes_referenceOrder map[*NoteStakeholderShape]uint

	// insertion point for slice of pointers maps
	NoteStakeholderShape_ControlPointShapes_reverseMap map[*ControlPointShape]*NoteStakeholderShape

	OnAfterNoteStakeholderShapeCreateCallback GongOnAfterCreateInterface[NoteStakeholderShape]
	OnAfterNoteStakeholderShapeUpdateCallback GongOnAfterUpdateInterface[NoteStakeholderShape]
	OnAfterNoteStakeholderShapeDeleteCallback GongOnAfterDeleteInterface[NoteStakeholderShape]

	NoteTaskShapes                map[*NoteTaskShape]struct{}
	NoteTaskShapes_instance       map[*NoteTaskShape]*NoteTaskShape
	NoteTaskShapes_mapString      map[string]*NoteTaskShape
	NoteTaskShapeOrder            uint
	NoteTaskShape_stagedOrder     map[*NoteTaskShape]uint
	NoteTaskShape_orderStaged     map[uint]*NoteTaskShape
	NoteTaskShapes_reference      map[*NoteTaskShape]*NoteTaskShape
	NoteTaskShapes_referenceOrder map[*NoteTaskShape]uint

	// insertion point for slice of pointers maps
	NoteTaskShape_ControlPointShapes_reverseMap map[*ControlPointShape]*NoteTaskShape

	OnAfterNoteTaskShapeCreateCallback GongOnAfterCreateInterface[NoteTaskShape]
	OnAfterNoteTaskShapeUpdateCallback GongOnAfterUpdateInterface[NoteTaskShape]
	OnAfterNoteTaskShapeDeleteCallback GongOnAfterDeleteInterface[NoteTaskShape]

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

	OnAfterRequirementCreateCallback GongOnAfterCreateInterface[Requirement]
	OnAfterRequirementUpdateCallback GongOnAfterUpdateInterface[Requirement]
	OnAfterRequirementDeleteCallback GongOnAfterDeleteInterface[Requirement]

	RequirementShapes                map[*RequirementShape]struct{}
	RequirementShapes_instance       map[*RequirementShape]*RequirementShape
	RequirementShapes_mapString      map[string]*RequirementShape
	RequirementShapeOrder            uint
	RequirementShape_stagedOrder     map[*RequirementShape]uint
	RequirementShape_orderStaged     map[uint]*RequirementShape
	RequirementShapes_reference      map[*RequirementShape]*RequirementShape
	RequirementShapes_referenceOrder map[*RequirementShape]uint

	// insertion point for slice of pointers maps
	OnAfterRequirementShapeCreateCallback GongOnAfterCreateInterface[RequirementShape]
	OnAfterRequirementShapeUpdateCallback GongOnAfterUpdateInterface[RequirementShape]
	OnAfterRequirementShapeDeleteCallback GongOnAfterDeleteInterface[RequirementShape]

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

	OnAfterStakeholderCreateCallback GongOnAfterCreateInterface[Stakeholder]
	OnAfterStakeholderUpdateCallback GongOnAfterUpdateInterface[Stakeholder]
	OnAfterStakeholderDeleteCallback GongOnAfterDeleteInterface[Stakeholder]

	StakeholderCompositionShapes                map[*StakeholderCompositionShape]struct{}
	StakeholderCompositionShapes_instance       map[*StakeholderCompositionShape]*StakeholderCompositionShape
	StakeholderCompositionShapes_mapString      map[string]*StakeholderCompositionShape
	StakeholderCompositionShapeOrder            uint
	StakeholderCompositionShape_stagedOrder     map[*StakeholderCompositionShape]uint
	StakeholderCompositionShape_orderStaged     map[uint]*StakeholderCompositionShape
	StakeholderCompositionShapes_reference      map[*StakeholderCompositionShape]*StakeholderCompositionShape
	StakeholderCompositionShapes_referenceOrder map[*StakeholderCompositionShape]uint

	// insertion point for slice of pointers maps
	StakeholderCompositionShape_ControlPointShapes_reverseMap map[*ControlPointShape]*StakeholderCompositionShape

	OnAfterStakeholderCompositionShapeCreateCallback GongOnAfterCreateInterface[StakeholderCompositionShape]
	OnAfterStakeholderCompositionShapeUpdateCallback GongOnAfterUpdateInterface[StakeholderCompositionShape]
	OnAfterStakeholderCompositionShapeDeleteCallback GongOnAfterDeleteInterface[StakeholderCompositionShape]

	StakeholderConcernShapes                map[*StakeholderConcernShape]struct{}
	StakeholderConcernShapes_instance       map[*StakeholderConcernShape]*StakeholderConcernShape
	StakeholderConcernShapes_mapString      map[string]*StakeholderConcernShape
	StakeholderConcernShapeOrder            uint
	StakeholderConcernShape_stagedOrder     map[*StakeholderConcernShape]uint
	StakeholderConcernShape_orderStaged     map[uint]*StakeholderConcernShape
	StakeholderConcernShapes_reference      map[*StakeholderConcernShape]*StakeholderConcernShape
	StakeholderConcernShapes_referenceOrder map[*StakeholderConcernShape]uint

	// insertion point for slice of pointers maps
	StakeholderConcernShape_ControlPointShapes_reverseMap map[*ControlPointShape]*StakeholderConcernShape

	OnAfterStakeholderConcernShapeCreateCallback GongOnAfterCreateInterface[StakeholderConcernShape]
	OnAfterStakeholderConcernShapeUpdateCallback GongOnAfterUpdateInterface[StakeholderConcernShape]
	OnAfterStakeholderConcernShapeDeleteCallback GongOnAfterDeleteInterface[StakeholderConcernShape]

	StakeholderShapes                map[*StakeholderShape]struct{}
	StakeholderShapes_instance       map[*StakeholderShape]*StakeholderShape
	StakeholderShapes_mapString      map[string]*StakeholderShape
	StakeholderShapeOrder            uint
	StakeholderShape_stagedOrder     map[*StakeholderShape]uint
	StakeholderShape_orderStaged     map[uint]*StakeholderShape
	StakeholderShapes_reference      map[*StakeholderShape]*StakeholderShape
	StakeholderShapes_referenceOrder map[*StakeholderShape]uint

	// insertion point for slice of pointers maps
	OnAfterStakeholderShapeCreateCallback GongOnAfterCreateInterface[StakeholderShape]
	OnAfterStakeholderShapeUpdateCallback GongOnAfterUpdateInterface[StakeholderShape]
	OnAfterStakeholderShapeDeleteCallback GongOnAfterDeleteInterface[StakeholderShape]

	SupportLevels                map[*SupportLevel]struct{}
	SupportLevels_instance       map[*SupportLevel]*SupportLevel
	SupportLevels_mapString      map[string]*SupportLevel
	SupportLevelOrder            uint
	SupportLevel_stagedOrder     map[*SupportLevel]uint
	SupportLevel_orderStaged     map[uint]*SupportLevel
	SupportLevels_reference      map[*SupportLevel]*SupportLevel
	SupportLevels_referenceOrder map[*SupportLevel]uint

	// insertion point for slice of pointers maps
	OnAfterSupportLevelCreateCallback GongOnAfterCreateInterface[SupportLevel]
	OnAfterSupportLevelUpdateCallback GongOnAfterUpdateInterface[SupportLevel]
	OnAfterSupportLevelDeleteCallback GongOnAfterDeleteInterface[SupportLevel]

	Tools                map[*Tool]struct{}
	Tools_instance       map[*Tool]*Tool
	Tools_mapString      map[string]*Tool
	ToolOrder            uint
	Tool_stagedOrder     map[*Tool]uint
	Tool_orderStaged     map[uint]*Tool
	Tools_reference      map[*Tool]*Tool
	Tools_referenceOrder map[*Tool]uint

	// insertion point for slice of pointers maps
	OnAfterToolCreateCallback GongOnAfterCreateInterface[Tool]
	OnAfterToolUpdateCallback GongOnAfterUpdateInterface[Tool]
	OnAfterToolDeleteCallback GongOnAfterDeleteInterface[Tool]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string
	MetaPackageImports     []*MetaPackageImport

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	// end of insertion point

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]GongModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF GongProbeIF

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

type GongStage = Stage

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
	err := stage.ParseAstString(commitToApply, true)
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
	err := stage.ParseAstString(commitToApply, true)
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
	__gong__clearReferences(&stage.AnalysisNeeds_reference, &stage.AnalysisNeeds_instance, &stage.AnalysisNeeds_referenceOrder)

	__gong__clearReferences(&stage.Concepts_reference, &stage.Concepts_instance, &stage.Concepts_referenceOrder)

	__gong__clearReferences(&stage.ConceptShapes_reference, &stage.ConceptShapes_instance, &stage.ConceptShapes_referenceOrder)

	__gong__clearReferences(&stage.Concerns_reference, &stage.Concerns_instance, &stage.Concerns_referenceOrder)

	__gong__clearReferences(&stage.ConcernCompositionShapes_reference, &stage.ConcernCompositionShapes_instance, &stage.ConcernCompositionShapes_referenceOrder)

	__gong__clearReferences(&stage.ConcernInputShapes_reference, &stage.ConcernInputShapes_instance, &stage.ConcernInputShapes_referenceOrder)

	__gong__clearReferences(&stage.ConcernOutputShapes_reference, &stage.ConcernOutputShapes_instance, &stage.ConcernOutputShapes_referenceOrder)

	__gong__clearReferences(&stage.ConcernShapes_reference, &stage.ConcernShapes_instance, &stage.ConcernShapes_referenceOrder)

	__gong__clearReferences(&stage.ControlPointShapes_reference, &stage.ControlPointShapes_instance, &stage.ControlPointShapes_referenceOrder)

	__gong__clearReferences(&stage.Deliverables_reference, &stage.Deliverables_instance, &stage.Deliverables_referenceOrder)

	__gong__clearReferences(&stage.DeliverableCompositionShapes_reference, &stage.DeliverableCompositionShapes_instance, &stage.DeliverableCompositionShapes_referenceOrder)

	__gong__clearReferences(&stage.DeliverableConceptShapes_reference, &stage.DeliverableConceptShapes_instance, &stage.DeliverableConceptShapes_referenceOrder)

	__gong__clearReferences(&stage.DeliverableShapes_reference, &stage.DeliverableShapes_instance, &stage.DeliverableShapes_referenceOrder)

	__gong__clearReferences(&stage.Diagrams_reference, &stage.Diagrams_instance, &stage.Diagrams_referenceOrder)

	__gong__clearReferences(&stage.DiagramShapes_reference, &stage.DiagramShapes_instance, &stage.DiagramShapes_referenceOrder)

	__gong__clearReferences(&stage.Librarys_reference, &stage.Librarys_instance, &stage.Librarys_referenceOrder)

	__gong__clearReferences(&stage.Notes_reference, &stage.Notes_instance, &stage.Notes_referenceOrder)

	__gong__clearReferences(&stage.NoteDeliverableShapes_reference, &stage.NoteDeliverableShapes_instance, &stage.NoteDeliverableShapes_referenceOrder)

	__gong__clearReferences(&stage.NoteShapes_reference, &stage.NoteShapes_instance, &stage.NoteShapes_referenceOrder)

	__gong__clearReferences(&stage.NoteStakeholderShapes_reference, &stage.NoteStakeholderShapes_instance, &stage.NoteStakeholderShapes_referenceOrder)

	__gong__clearReferences(&stage.NoteTaskShapes_reference, &stage.NoteTaskShapes_instance, &stage.NoteTaskShapes_referenceOrder)

	__gong__clearReferences(&stage.Requirements_reference, &stage.Requirements_instance, &stage.Requirements_referenceOrder)

	__gong__clearReferences(&stage.RequirementShapes_reference, &stage.RequirementShapes_instance, &stage.RequirementShapes_referenceOrder)

	__gong__clearReferences(&stage.Stakeholders_reference, &stage.Stakeholders_instance, &stage.Stakeholders_referenceOrder)

	__gong__clearReferences(&stage.StakeholderCompositionShapes_reference, &stage.StakeholderCompositionShapes_instance, &stage.StakeholderCompositionShapes_referenceOrder)

	__gong__clearReferences(&stage.StakeholderConcernShapes_reference, &stage.StakeholderConcernShapes_instance, &stage.StakeholderConcernShapes_referenceOrder)

	__gong__clearReferences(&stage.StakeholderShapes_reference, &stage.StakeholderShapes_instance, &stage.StakeholderShapes_referenceOrder)

	__gong__clearReferences(&stage.SupportLevels_reference, &stage.SupportLevels_instance, &stage.SupportLevels_referenceOrder)

	__gong__clearReferences(&stage.Tools_reference, &stage.Tools_instance, &stage.Tools_referenceOrder)

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
	stage.AnalysisNeedOrder = __gong__recomputeOrder(stage.AnalysisNeed_stagedOrder)

	stage.ConceptOrder = __gong__recomputeOrder(stage.Concept_stagedOrder)

	stage.ConceptShapeOrder = __gong__recomputeOrder(stage.ConceptShape_stagedOrder)

	stage.ConcernOrder = __gong__recomputeOrder(stage.Concern_stagedOrder)

	stage.ConcernCompositionShapeOrder = __gong__recomputeOrder(stage.ConcernCompositionShape_stagedOrder)

	stage.ConcernInputShapeOrder = __gong__recomputeOrder(stage.ConcernInputShape_stagedOrder)

	stage.ConcernOutputShapeOrder = __gong__recomputeOrder(stage.ConcernOutputShape_stagedOrder)

	stage.ConcernShapeOrder = __gong__recomputeOrder(stage.ConcernShape_stagedOrder)

	stage.ControlPointShapeOrder = __gong__recomputeOrder(stage.ControlPointShape_stagedOrder)

	stage.DeliverableOrder = __gong__recomputeOrder(stage.Deliverable_stagedOrder)

	stage.DeliverableCompositionShapeOrder = __gong__recomputeOrder(stage.DeliverableCompositionShape_stagedOrder)

	stage.DeliverableConceptShapeOrder = __gong__recomputeOrder(stage.DeliverableConceptShape_stagedOrder)

	stage.DeliverableShapeOrder = __gong__recomputeOrder(stage.DeliverableShape_stagedOrder)

	stage.DiagramOrder = __gong__recomputeOrder(stage.Diagram_stagedOrder)

	stage.DiagramShapeOrder = __gong__recomputeOrder(stage.DiagramShape_stagedOrder)

	stage.LibraryOrder = __gong__recomputeOrder(stage.Library_stagedOrder)

	stage.NoteOrder = __gong__recomputeOrder(stage.Note_stagedOrder)

	stage.NoteDeliverableShapeOrder = __gong__recomputeOrder(stage.NoteDeliverableShape_stagedOrder)

	stage.NoteShapeOrder = __gong__recomputeOrder(stage.NoteShape_stagedOrder)

	stage.NoteStakeholderShapeOrder = __gong__recomputeOrder(stage.NoteStakeholderShape_stagedOrder)

	stage.NoteTaskShapeOrder = __gong__recomputeOrder(stage.NoteTaskShape_stagedOrder)

	stage.RequirementOrder = __gong__recomputeOrder(stage.Requirement_stagedOrder)

	stage.RequirementShapeOrder = __gong__recomputeOrder(stage.RequirementShape_stagedOrder)

	stage.StakeholderOrder = __gong__recomputeOrder(stage.Stakeholder_stagedOrder)

	stage.StakeholderCompositionShapeOrder = __gong__recomputeOrder(stage.StakeholderCompositionShape_stagedOrder)

	stage.StakeholderConcernShapeOrder = __gong__recomputeOrder(stage.StakeholderConcernShape_stagedOrder)

	stage.StakeholderShapeOrder = __gong__recomputeOrder(stage.StakeholderShape_stagedOrder)

	stage.SupportLevelOrder = __gong__recomputeOrder(stage.SupportLevel_stagedOrder)

	stage.ToolOrder = __gong__recomputeOrder(stage.Tool_stagedOrder)

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF GongProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() GongProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *AnalysisNeed:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.AnalysisNeeds, stage.AnalysisNeed_stagedOrder))
	case *Concept:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Concepts, stage.Concept_stagedOrder))
	case *ConceptShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ConceptShapes, stage.ConceptShape_stagedOrder))
	case *Concern:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Concerns, stage.Concern_stagedOrder))
	case *ConcernCompositionShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ConcernCompositionShapes, stage.ConcernCompositionShape_stagedOrder))
	case *ConcernInputShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ConcernInputShapes, stage.ConcernInputShape_stagedOrder))
	case *ConcernOutputShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ConcernOutputShapes, stage.ConcernOutputShape_stagedOrder))
	case *ConcernShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ConcernShapes, stage.ConcernShape_stagedOrder))
	case *ControlPointShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder))
	case *Deliverable:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Deliverables, stage.Deliverable_stagedOrder))
	case *DeliverableCompositionShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DeliverableCompositionShapes, stage.DeliverableCompositionShape_stagedOrder))
	case *DeliverableConceptShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DeliverableConceptShapes, stage.DeliverableConceptShape_stagedOrder))
	case *DeliverableShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DeliverableShapes, stage.DeliverableShape_stagedOrder))
	case *Diagram:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Diagrams, stage.Diagram_stagedOrder))
	case *DiagramShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DiagramShapes, stage.DiagramShape_stagedOrder))
	case *Library:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder))
	case *Note:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Notes, stage.Note_stagedOrder))
	case *NoteDeliverableShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteDeliverableShapes, stage.NoteDeliverableShape_stagedOrder))
	case *NoteShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteShapes, stage.NoteShape_stagedOrder))
	case *NoteStakeholderShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteStakeholderShapes, stage.NoteStakeholderShape_stagedOrder))
	case *NoteTaskShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder))
	case *Requirement:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Requirements, stage.Requirement_stagedOrder))
	case *RequirementShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.RequirementShapes, stage.RequirementShape_stagedOrder))
	case *Stakeholder:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Stakeholders, stage.Stakeholder_stagedOrder))
	case *StakeholderCompositionShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.StakeholderCompositionShapes, stage.StakeholderCompositionShape_stagedOrder))
	case *StakeholderConcernShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.StakeholderConcernShapes, stage.StakeholderConcernShape_stagedOrder))
	case *StakeholderShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.StakeholderShapes, stage.StakeholderShape_stagedOrder))
	case *SupportLevel:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SupportLevels, stage.SupportLevel_stagedOrder))
	case *Tool:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Tools, stage.Tool_stagedOrder))

	}
	return
}

func __gong__getStructInstancesByOrder[T GongstructPtr](set map[T]struct{}, order map[T]uint) (res []T) {
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
			log.Fatalf("getStructInstancesByOrder: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func __gong__castSlice[T any, S any](s []S) []T {
	res := make([]T, len(s))
	for i, v := range s {
		res[i] = any(v).(T)
	}
	return res
}

func __gong__stage[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	order *uint,
	mapString map[string]T,
	instance T,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		stagedOrder[instance] = *order
		orderStaged[*order] = instance
		*order++
	}
	mapString[name] = instance
}

func __gong__stagePreserveOrder[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	currentOrder *uint,
	mapString map[string]T,
	instance T,
	order uint,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		if order > *currentOrder {
			*currentOrder = order
		}
		stagedOrder[instance] = order
		orderStaged[order] = instance
		*currentOrder++
	}
	mapString[name] = instance
}

func __gong__unstage[T comparable](
	instances map[T]struct{},
	mapString map[string]T,
	instance T,
	name string,
) {
	delete(instances, instance)
	delete(mapString, name)
}

func __gong__recomputeOrder[T comparable](stagedOrder map[T]uint) uint {
	var maxOrder uint
	var found bool
	for _, order := range stagedOrder {
		if !found || order > maxOrder {
			maxOrder = order
			found = true
		}
	}
	if found {
		return maxOrder + 1
	}
	return 0
}

func __gong__rebuildMapString[T interface {
	comparable
	GetName() string
}](staged map[T]struct{}, mapString *map[string]T) {
	*mapString = make(map[string]T, len(staged))
	for instance := range staged {
		(*mapString)[instance.GetName()] = instance
	}
}

func __gong__clearReferences[T comparable](ref *map[T]T, inst *map[T]T, refOrder *map[T]uint) {
	*ref = make(map[T]T)
	*inst = make(map[T]T)
	*refOrder = make(map[T]uint)
}

func __gong__resetStageType[T comparable](staged *map[T]struct{}, mapString *map[string]T, stagedOrder *map[T]uint, order *uint) {
	*staged = make(map[T]struct{})
	*mapString = make(map[string]T)
	*stagedOrder = make(map[T]uint)
	*order = 0
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/dsm/capture/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type GongOnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

type OnInitCommitInterface = GongOnInitCommitInterface

// GongOnAfterCreateInterface callback when an instance is updated from the front
type GongOnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

type OnAfterCreateInterface[Type Gongstruct] = GongOnAfterCreateInterface[Type]

// GongOnAfterUpdateInterface callback when an instance is updated from the front
type GongOnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

type OnAfterUpdateInterface[Type Gongstruct] = GongOnAfterUpdateInterface[Type]

// GongOnAfterDeleteInterface callback when an instance is updated from the front
type GongOnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type OnAfterDeleteInterface[Type Gongstruct] = GongOnAfterDeleteInterface[Type]

type GongBackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

type BackRepoInterface = GongBackRepoInterface

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		AnalysisNeeds:           make(map[*AnalysisNeed]struct{}),
		AnalysisNeeds_mapString: make(map[string]*AnalysisNeed),

		Concepts:           make(map[*Concept]struct{}),
		Concepts_mapString: make(map[string]*Concept),

		ConceptShapes:           make(map[*ConceptShape]struct{}),
		ConceptShapes_mapString: make(map[string]*ConceptShape),

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

		ControlPointShapes:           make(map[*ControlPointShape]struct{}),
		ControlPointShapes_mapString: make(map[string]*ControlPointShape),

		Deliverables:           make(map[*Deliverable]struct{}),
		Deliverables_mapString: make(map[string]*Deliverable),

		DeliverableCompositionShapes:           make(map[*DeliverableCompositionShape]struct{}),
		DeliverableCompositionShapes_mapString: make(map[string]*DeliverableCompositionShape),

		DeliverableConceptShapes:           make(map[*DeliverableConceptShape]struct{}),
		DeliverableConceptShapes_mapString: make(map[string]*DeliverableConceptShape),

		DeliverableShapes:           make(map[*DeliverableShape]struct{}),
		DeliverableShapes_mapString: make(map[string]*DeliverableShape),

		Diagrams:           make(map[*Diagram]struct{}),
		Diagrams_mapString: make(map[string]*Diagram),

		DiagramShapes:           make(map[*DiagramShape]struct{}),
		DiagramShapes_mapString: make(map[string]*DiagramShape),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		NoteDeliverableShapes:           make(map[*NoteDeliverableShape]struct{}),
		NoteDeliverableShapes_mapString: make(map[string]*NoteDeliverableShape),

		NoteShapes:           make(map[*NoteShape]struct{}),
		NoteShapes_mapString: make(map[string]*NoteShape),

		NoteStakeholderShapes:           make(map[*NoteStakeholderShape]struct{}),
		NoteStakeholderShapes_mapString: make(map[string]*NoteStakeholderShape),

		NoteTaskShapes:           make(map[*NoteTaskShape]struct{}),
		NoteTaskShapes_mapString: make(map[string]*NoteTaskShape),

		Requirements:           make(map[*Requirement]struct{}),
		Requirements_mapString: make(map[string]*Requirement),

		RequirementShapes:           make(map[*RequirementShape]struct{}),
		RequirementShapes_mapString: make(map[string]*RequirementShape),

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

		ConceptShape_stagedOrder: make(map[*ConceptShape]uint),
		ConceptShape_orderStaged: make(map[uint]*ConceptShape),
		ConceptShapes_reference:  make(map[*ConceptShape]*ConceptShape),

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

		ControlPointShape_stagedOrder: make(map[*ControlPointShape]uint),
		ControlPointShape_orderStaged: make(map[uint]*ControlPointShape),
		ControlPointShapes_reference:  make(map[*ControlPointShape]*ControlPointShape),

		Deliverable_stagedOrder: make(map[*Deliverable]uint),
		Deliverable_orderStaged: make(map[uint]*Deliverable),
		Deliverables_reference:  make(map[*Deliverable]*Deliverable),

		DeliverableCompositionShape_stagedOrder: make(map[*DeliverableCompositionShape]uint),
		DeliverableCompositionShape_orderStaged: make(map[uint]*DeliverableCompositionShape),
		DeliverableCompositionShapes_reference:  make(map[*DeliverableCompositionShape]*DeliverableCompositionShape),

		DeliverableConceptShape_stagedOrder: make(map[*DeliverableConceptShape]uint),
		DeliverableConceptShape_orderStaged: make(map[uint]*DeliverableConceptShape),
		DeliverableConceptShapes_reference:  make(map[*DeliverableConceptShape]*DeliverableConceptShape),

		DeliverableShape_stagedOrder: make(map[*DeliverableShape]uint),
		DeliverableShape_orderStaged: make(map[uint]*DeliverableShape),
		DeliverableShapes_reference:  make(map[*DeliverableShape]*DeliverableShape),

		Diagram_stagedOrder: make(map[*Diagram]uint),
		Diagram_orderStaged: make(map[uint]*Diagram),
		Diagrams_reference:  make(map[*Diagram]*Diagram),

		DiagramShape_stagedOrder: make(map[*DiagramShape]uint),
		DiagramShape_orderStaged: make(map[uint]*DiagramShape),
		DiagramShapes_reference:  make(map[*DiagramShape]*DiagramShape),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		Note_stagedOrder: make(map[*Note]uint),
		Note_orderStaged: make(map[uint]*Note),
		Notes_reference:  make(map[*Note]*Note),

		NoteDeliverableShape_stagedOrder: make(map[*NoteDeliverableShape]uint),
		NoteDeliverableShape_orderStaged: make(map[uint]*NoteDeliverableShape),
		NoteDeliverableShapes_reference:  make(map[*NoteDeliverableShape]*NoteDeliverableShape),

		NoteShape_stagedOrder: make(map[*NoteShape]uint),
		NoteShape_orderStaged: make(map[uint]*NoteShape),
		NoteShapes_reference:  make(map[*NoteShape]*NoteShape),

		NoteStakeholderShape_stagedOrder: make(map[*NoteStakeholderShape]uint),
		NoteStakeholderShape_orderStaged: make(map[uint]*NoteStakeholderShape),
		NoteStakeholderShapes_reference:  make(map[*NoteStakeholderShape]*NoteStakeholderShape),

		NoteTaskShape_stagedOrder: make(map[*NoteTaskShape]uint),
		NoteTaskShape_orderStaged: make(map[uint]*NoteTaskShape),
		NoteTaskShapes_reference:  make(map[*NoteTaskShape]*NoteTaskShape),

		Requirement_stagedOrder: make(map[*Requirement]uint),
		Requirement_orderStaged: make(map[uint]*Requirement),
		Requirements_reference:  make(map[*Requirement]*Requirement),

		RequirementShape_stagedOrder: make(map[*RequirementShape]uint),
		RequirementShape_orderStaged: make(map[uint]*RequirementShape),
		RequirementShapes_reference:  make(map[*RequirementShape]*RequirementShape),

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
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AnalysisNeed": &AnalysisNeedUnmarshaller{},

			"Concept": &ConceptUnmarshaller{},

			"ConceptShape": &ConceptShapeUnmarshaller{},

			"Concern": &ConcernUnmarshaller{},

			"ConcernCompositionShape": &ConcernCompositionShapeUnmarshaller{},

			"ConcernInputShape": &ConcernInputShapeUnmarshaller{},

			"ConcernOutputShape": &ConcernOutputShapeUnmarshaller{},

			"ConcernShape": &ConcernShapeUnmarshaller{},

			"ControlPointShape": &ControlPointShapeUnmarshaller{},

			"Deliverable": &DeliverableUnmarshaller{},

			"DeliverableCompositionShape": &DeliverableCompositionShapeUnmarshaller{},

			"DeliverableConceptShape": &DeliverableConceptShapeUnmarshaller{},

			"DeliverableShape": &DeliverableShapeUnmarshaller{},

			"Diagram": &DiagramUnmarshaller{},

			"DiagramShape": &DiagramShapeUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteDeliverableShape": &NoteDeliverableShapeUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"NoteStakeholderShape": &NoteStakeholderShapeUnmarshaller{},

			"NoteTaskShape": &NoteTaskShapeUnmarshaller{},

			"Requirement": &RequirementUnmarshaller{},

			"RequirementShape": &RequirementShapeUnmarshaller{},

			"Stakeholder": &StakeholderUnmarshaller{},

			"StakeholderCompositionShape": &StakeholderCompositionShapeUnmarshaller{},

			"StakeholderConcernShape": &StakeholderConcernShapeUnmarshaller{},

			"StakeholderShape": &StakeholderShapeUnmarshaller{},

			"SupportLevel": &SupportLevelUnmarshaller{},

			"Tool": &ToolUnmarshaller{},

			// end of insertion point
		},

		navigationMode: GongNavigationModeNormal,
	}

	return
}

// GetOrder is the Stage method returning the order of a gongstruct instance.
func (stage *Stage) GetOrder(instance GongstructIF) uint {
	if instance != nil {
		return instance.GongGetOrder(stage)
	}
	return 0
}

// GetInstanceFromOrder is the Stage method returning a gongstruct instance from its order.
func (stage *Stage) GetInstanceFromOrder[Type GongstructPtr](order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *AnalysisNeed:
		return any(stage.AnalysisNeed_orderStaged[order]).(Type)
	case *Concept:
		return any(stage.Concept_orderStaged[order]).(Type)
	case *ConceptShape:
		return any(stage.ConceptShape_orderStaged[order]).(Type)
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
	case *ControlPointShape:
		return any(stage.ControlPointShape_orderStaged[order]).(Type)
	case *Deliverable:
		return any(stage.Deliverable_orderStaged[order]).(Type)
	case *DeliverableCompositionShape:
		return any(stage.DeliverableCompositionShape_orderStaged[order]).(Type)
	case *DeliverableConceptShape:
		return any(stage.DeliverableConceptShape_orderStaged[order]).(Type)
	case *DeliverableShape:
		return any(stage.DeliverableShape_orderStaged[order]).(Type)
	case *Diagram:
		return any(stage.Diagram_orderStaged[order]).(Type)
	case *DiagramShape:
		return any(stage.DiagramShape_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Note:
		return any(stage.Note_orderStaged[order]).(Type)
	case *NoteDeliverableShape:
		return any(stage.NoteDeliverableShape_orderStaged[order]).(Type)
	case *NoteShape:
		return any(stage.NoteShape_orderStaged[order]).(Type)
	case *NoteStakeholderShape:
		return any(stage.NoteStakeholderShape_orderStaged[order]).(Type)
	case *NoteTaskShape:
		return any(stage.NoteTaskShape_orderStaged[order]).(Type)
	case *Requirement:
		return any(stage.Requirement_orderStaged[order]).(Type)
	case *RequirementShape:
		return any(stage.RequirementShape_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["ConceptShape"] = len(stage.ConceptShapes)
	stage.Map_GongStructName_InstancesNb["Concern"] = len(stage.Concerns)
	stage.Map_GongStructName_InstancesNb["ConcernCompositionShape"] = len(stage.ConcernCompositionShapes)
	stage.Map_GongStructName_InstancesNb["ConcernInputShape"] = len(stage.ConcernInputShapes)
	stage.Map_GongStructName_InstancesNb["ConcernOutputShape"] = len(stage.ConcernOutputShapes)
	stage.Map_GongStructName_InstancesNb["ConcernShape"] = len(stage.ConcernShapes)
	stage.Map_GongStructName_InstancesNb["ControlPointShape"] = len(stage.ControlPointShapes)
	stage.Map_GongStructName_InstancesNb["Deliverable"] = len(stage.Deliverables)
	stage.Map_GongStructName_InstancesNb["DeliverableCompositionShape"] = len(stage.DeliverableCompositionShapes)
	stage.Map_GongStructName_InstancesNb["DeliverableConceptShape"] = len(stage.DeliverableConceptShapes)
	stage.Map_GongStructName_InstancesNb["DeliverableShape"] = len(stage.DeliverableShapes)
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["DiagramShape"] = len(stage.DiagramShapes)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteDeliverableShape"] = len(stage.NoteDeliverableShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteStakeholderShape"] = len(stage.NoteStakeholderShapes)
	stage.Map_GongStructName_InstancesNb["NoteTaskShape"] = len(stage.NoteTaskShapes)
	stage.Map_GongStructName_InstancesNb["Requirement"] = len(stage.Requirements)
	stage.Map_GongStructName_InstancesNb["RequirementShape"] = len(stage.RequirementShapes)
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
	__gong__stage(stage.AnalysisNeeds, stage.AnalysisNeed_stagedOrder, stage.AnalysisNeed_orderStaged, &stage.AnalysisNeedOrder, stage.AnalysisNeeds_mapString, analysisneed, analysisneed.Name)
	return analysisneed
}

// StagePreserveOrder puts analysisneed to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AnalysisNeedOrder
// - update stage.AnalysisNeedOrder accordingly
func (analysisneed *AnalysisNeed) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AnalysisNeeds, stage.AnalysisNeed_stagedOrder, stage.AnalysisNeed_orderStaged, &stage.AnalysisNeedOrder, stage.AnalysisNeeds_mapString, analysisneed, order, analysisneed.Name)
}

// Unstage removes analysisneed off the model stage
func (analysisneed *AnalysisNeed) Unstage(stage *Stage) *AnalysisNeed {
	__gong__unstage(stage.AnalysisNeeds, stage.AnalysisNeeds_mapString, analysisneed, analysisneed.Name)
	return analysisneed
}

// UnstageVoid removes analysisneed off the model stage
func (analysisneed *AnalysisNeed) UnstageVoid(stage *Stage) {
	analysisneed.Unstage(stage)
}

func (analysisneed *AnalysisNeed) StageVoid(stage *Stage) {
	analysisneed.Stage(stage)
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
	__gong__stage(stage.Concepts, stage.Concept_stagedOrder, stage.Concept_orderStaged, &stage.ConceptOrder, stage.Concepts_mapString, concept, concept.Name)
	return concept
}

// StagePreserveOrder puts concept to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConceptOrder
// - update stage.ConceptOrder accordingly
func (concept *Concept) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Concepts, stage.Concept_stagedOrder, stage.Concept_orderStaged, &stage.ConceptOrder, stage.Concepts_mapString, concept, order, concept.Name)
}

// Unstage removes concept off the model stage
func (concept *Concept) Unstage(stage *Stage) *Concept {
	__gong__unstage(stage.Concepts, stage.Concepts_mapString, concept, concept.Name)
	return concept
}

// UnstageVoid removes concept off the model stage
func (concept *Concept) UnstageVoid(stage *Stage) {
	concept.Unstage(stage)
}

func (concept *Concept) StageVoid(stage *Stage) {
	concept.Stage(stage)
}

// for satisfaction of GongStruct interface
func (concept *Concept) GetName() (res string) {
	return concept.Name
}

// for satisfaction of GongStruct interface
func (concept *Concept) SetName(name string) {
	concept.Name = name
}

// Stage puts conceptshape to the model stage
func (conceptshape *ConceptShape) Stage(stage *Stage) *ConceptShape {
	__gong__stage(stage.ConceptShapes, stage.ConceptShape_stagedOrder, stage.ConceptShape_orderStaged, &stage.ConceptShapeOrder, stage.ConceptShapes_mapString, conceptshape, conceptshape.Name)
	return conceptshape
}

// StagePreserveOrder puts conceptshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConceptShapeOrder
// - update stage.ConceptShapeOrder accordingly
func (conceptshape *ConceptShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ConceptShapes, stage.ConceptShape_stagedOrder, stage.ConceptShape_orderStaged, &stage.ConceptShapeOrder, stage.ConceptShapes_mapString, conceptshape, order, conceptshape.Name)
}

// Unstage removes conceptshape off the model stage
func (conceptshape *ConceptShape) Unstage(stage *Stage) *ConceptShape {
	__gong__unstage(stage.ConceptShapes, stage.ConceptShapes_mapString, conceptshape, conceptshape.Name)
	return conceptshape
}

// UnstageVoid removes conceptshape off the model stage
func (conceptshape *ConceptShape) UnstageVoid(stage *Stage) {
	conceptshape.Unstage(stage)
}

func (conceptshape *ConceptShape) StageVoid(stage *Stage) {
	conceptshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (conceptshape *ConceptShape) GetName() (res string) {
	return conceptshape.Name
}

// for satisfaction of GongStruct interface
func (conceptshape *ConceptShape) SetName(name string) {
	conceptshape.Name = name
}

// Stage puts concern to the model stage
func (concern *Concern) Stage(stage *Stage) *Concern {
	__gong__stage(stage.Concerns, stage.Concern_stagedOrder, stage.Concern_orderStaged, &stage.ConcernOrder, stage.Concerns_mapString, concern, concern.Name)
	return concern
}

// StagePreserveOrder puts concern to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernOrder
// - update stage.ConcernOrder accordingly
func (concern *Concern) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Concerns, stage.Concern_stagedOrder, stage.Concern_orderStaged, &stage.ConcernOrder, stage.Concerns_mapString, concern, order, concern.Name)
}

// Unstage removes concern off the model stage
func (concern *Concern) Unstage(stage *Stage) *Concern {
	__gong__unstage(stage.Concerns, stage.Concerns_mapString, concern, concern.Name)
	return concern
}

// UnstageVoid removes concern off the model stage
func (concern *Concern) UnstageVoid(stage *Stage) {
	concern.Unstage(stage)
}

func (concern *Concern) StageVoid(stage *Stage) {
	concern.Stage(stage)
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
	__gong__stage(stage.ConcernCompositionShapes, stage.ConcernCompositionShape_stagedOrder, stage.ConcernCompositionShape_orderStaged, &stage.ConcernCompositionShapeOrder, stage.ConcernCompositionShapes_mapString, concerncompositionshape, concerncompositionshape.Name)
	return concerncompositionshape
}

// StagePreserveOrder puts concerncompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernCompositionShapeOrder
// - update stage.ConcernCompositionShapeOrder accordingly
func (concerncompositionshape *ConcernCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ConcernCompositionShapes, stage.ConcernCompositionShape_stagedOrder, stage.ConcernCompositionShape_orderStaged, &stage.ConcernCompositionShapeOrder, stage.ConcernCompositionShapes_mapString, concerncompositionshape, order, concerncompositionshape.Name)
}

// Unstage removes concerncompositionshape off the model stage
func (concerncompositionshape *ConcernCompositionShape) Unstage(stage *Stage) *ConcernCompositionShape {
	__gong__unstage(stage.ConcernCompositionShapes, stage.ConcernCompositionShapes_mapString, concerncompositionshape, concerncompositionshape.Name)
	return concerncompositionshape
}

// UnstageVoid removes concerncompositionshape off the model stage
func (concerncompositionshape *ConcernCompositionShape) UnstageVoid(stage *Stage) {
	concerncompositionshape.Unstage(stage)
}

func (concerncompositionshape *ConcernCompositionShape) StageVoid(stage *Stage) {
	concerncompositionshape.Stage(stage)
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
	__gong__stage(stage.ConcernInputShapes, stage.ConcernInputShape_stagedOrder, stage.ConcernInputShape_orderStaged, &stage.ConcernInputShapeOrder, stage.ConcernInputShapes_mapString, concerninputshape, concerninputshape.Name)
	return concerninputshape
}

// StagePreserveOrder puts concerninputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernInputShapeOrder
// - update stage.ConcernInputShapeOrder accordingly
func (concerninputshape *ConcernInputShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ConcernInputShapes, stage.ConcernInputShape_stagedOrder, stage.ConcernInputShape_orderStaged, &stage.ConcernInputShapeOrder, stage.ConcernInputShapes_mapString, concerninputshape, order, concerninputshape.Name)
}

// Unstage removes concerninputshape off the model stage
func (concerninputshape *ConcernInputShape) Unstage(stage *Stage) *ConcernInputShape {
	__gong__unstage(stage.ConcernInputShapes, stage.ConcernInputShapes_mapString, concerninputshape, concerninputshape.Name)
	return concerninputshape
}

// UnstageVoid removes concerninputshape off the model stage
func (concerninputshape *ConcernInputShape) UnstageVoid(stage *Stage) {
	concerninputshape.Unstage(stage)
}

func (concerninputshape *ConcernInputShape) StageVoid(stage *Stage) {
	concerninputshape.Stage(stage)
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
	__gong__stage(stage.ConcernOutputShapes, stage.ConcernOutputShape_stagedOrder, stage.ConcernOutputShape_orderStaged, &stage.ConcernOutputShapeOrder, stage.ConcernOutputShapes_mapString, concernoutputshape, concernoutputshape.Name)
	return concernoutputshape
}

// StagePreserveOrder puts concernoutputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernOutputShapeOrder
// - update stage.ConcernOutputShapeOrder accordingly
func (concernoutputshape *ConcernOutputShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ConcernOutputShapes, stage.ConcernOutputShape_stagedOrder, stage.ConcernOutputShape_orderStaged, &stage.ConcernOutputShapeOrder, stage.ConcernOutputShapes_mapString, concernoutputshape, order, concernoutputshape.Name)
}

// Unstage removes concernoutputshape off the model stage
func (concernoutputshape *ConcernOutputShape) Unstage(stage *Stage) *ConcernOutputShape {
	__gong__unstage(stage.ConcernOutputShapes, stage.ConcernOutputShapes_mapString, concernoutputshape, concernoutputshape.Name)
	return concernoutputshape
}

// UnstageVoid removes concernoutputshape off the model stage
func (concernoutputshape *ConcernOutputShape) UnstageVoid(stage *Stage) {
	concernoutputshape.Unstage(stage)
}

func (concernoutputshape *ConcernOutputShape) StageVoid(stage *Stage) {
	concernoutputshape.Stage(stage)
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
	__gong__stage(stage.ConcernShapes, stage.ConcernShape_stagedOrder, stage.ConcernShape_orderStaged, &stage.ConcernShapeOrder, stage.ConcernShapes_mapString, concernshape, concernshape.Name)
	return concernshape
}

// StagePreserveOrder puts concernshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ConcernShapeOrder
// - update stage.ConcernShapeOrder accordingly
func (concernshape *ConcernShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ConcernShapes, stage.ConcernShape_stagedOrder, stage.ConcernShape_orderStaged, &stage.ConcernShapeOrder, stage.ConcernShapes_mapString, concernshape, order, concernshape.Name)
}

// Unstage removes concernshape off the model stage
func (concernshape *ConcernShape) Unstage(stage *Stage) *ConcernShape {
	__gong__unstage(stage.ConcernShapes, stage.ConcernShapes_mapString, concernshape, concernshape.Name)
	return concernshape
}

// UnstageVoid removes concernshape off the model stage
func (concernshape *ConcernShape) UnstageVoid(stage *Stage) {
	concernshape.Unstage(stage)
}

func (concernshape *ConcernShape) StageVoid(stage *Stage) {
	concernshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (concernshape *ConcernShape) GetName() (res string) {
	return concernshape.Name
}

// for satisfaction of GongStruct interface
func (concernshape *ConcernShape) SetName(name string) {
	concernshape.Name = name
}

// Stage puts controlpointshape to the model stage
func (controlpointshape *ControlPointShape) Stage(stage *Stage) *ControlPointShape {
	__gong__stage(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder, stage.ControlPointShape_orderStaged, &stage.ControlPointShapeOrder, stage.ControlPointShapes_mapString, controlpointshape, controlpointshape.Name)
	return controlpointshape
}

// StagePreserveOrder puts controlpointshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlPointShapeOrder
// - update stage.ControlPointShapeOrder accordingly
func (controlpointshape *ControlPointShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ControlPointShapes, stage.ControlPointShape_stagedOrder, stage.ControlPointShape_orderStaged, &stage.ControlPointShapeOrder, stage.ControlPointShapes_mapString, controlpointshape, order, controlpointshape.Name)
}

// Unstage removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) Unstage(stage *Stage) *ControlPointShape {
	__gong__unstage(stage.ControlPointShapes, stage.ControlPointShapes_mapString, controlpointshape, controlpointshape.Name)
	return controlpointshape
}

// UnstageVoid removes controlpointshape off the model stage
func (controlpointshape *ControlPointShape) UnstageVoid(stage *Stage) {
	controlpointshape.Unstage(stage)
}

func (controlpointshape *ControlPointShape) StageVoid(stage *Stage) {
	controlpointshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (controlpointshape *ControlPointShape) GetName() (res string) {
	return controlpointshape.Name
}

// for satisfaction of GongStruct interface
func (controlpointshape *ControlPointShape) SetName(name string) {
	controlpointshape.Name = name
}

// Stage puts deliverable to the model stage
func (deliverable *Deliverable) Stage(stage *Stage) *Deliverable {
	__gong__stage(stage.Deliverables, stage.Deliverable_stagedOrder, stage.Deliverable_orderStaged, &stage.DeliverableOrder, stage.Deliverables_mapString, deliverable, deliverable.Name)
	return deliverable
}

// StagePreserveOrder puts deliverable to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DeliverableOrder
// - update stage.DeliverableOrder accordingly
func (deliverable *Deliverable) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Deliverables, stage.Deliverable_stagedOrder, stage.Deliverable_orderStaged, &stage.DeliverableOrder, stage.Deliverables_mapString, deliverable, order, deliverable.Name)
}

// Unstage removes deliverable off the model stage
func (deliverable *Deliverable) Unstage(stage *Stage) *Deliverable {
	__gong__unstage(stage.Deliverables, stage.Deliverables_mapString, deliverable, deliverable.Name)
	return deliverable
}

// UnstageVoid removes deliverable off the model stage
func (deliverable *Deliverable) UnstageVoid(stage *Stage) {
	deliverable.Unstage(stage)
}

func (deliverable *Deliverable) StageVoid(stage *Stage) {
	deliverable.Stage(stage)
}

// for satisfaction of GongStruct interface
func (deliverable *Deliverable) GetName() (res string) {
	return deliverable.Name
}

// for satisfaction of GongStruct interface
func (deliverable *Deliverable) SetName(name string) {
	deliverable.Name = name
}

// Stage puts deliverablecompositionshape to the model stage
func (deliverablecompositionshape *DeliverableCompositionShape) Stage(stage *Stage) *DeliverableCompositionShape {
	__gong__stage(stage.DeliverableCompositionShapes, stage.DeliverableCompositionShape_stagedOrder, stage.DeliverableCompositionShape_orderStaged, &stage.DeliverableCompositionShapeOrder, stage.DeliverableCompositionShapes_mapString, deliverablecompositionshape, deliverablecompositionshape.Name)
	return deliverablecompositionshape
}

// StagePreserveOrder puts deliverablecompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DeliverableCompositionShapeOrder
// - update stage.DeliverableCompositionShapeOrder accordingly
func (deliverablecompositionshape *DeliverableCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DeliverableCompositionShapes, stage.DeliverableCompositionShape_stagedOrder, stage.DeliverableCompositionShape_orderStaged, &stage.DeliverableCompositionShapeOrder, stage.DeliverableCompositionShapes_mapString, deliverablecompositionshape, order, deliverablecompositionshape.Name)
}

// Unstage removes deliverablecompositionshape off the model stage
func (deliverablecompositionshape *DeliverableCompositionShape) Unstage(stage *Stage) *DeliverableCompositionShape {
	__gong__unstage(stage.DeliverableCompositionShapes, stage.DeliverableCompositionShapes_mapString, deliverablecompositionshape, deliverablecompositionshape.Name)
	return deliverablecompositionshape
}

// UnstageVoid removes deliverablecompositionshape off the model stage
func (deliverablecompositionshape *DeliverableCompositionShape) UnstageVoid(stage *Stage) {
	deliverablecompositionshape.Unstage(stage)
}

func (deliverablecompositionshape *DeliverableCompositionShape) StageVoid(stage *Stage) {
	deliverablecompositionshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (deliverablecompositionshape *DeliverableCompositionShape) GetName() (res string) {
	return deliverablecompositionshape.Name
}

// for satisfaction of GongStruct interface
func (deliverablecompositionshape *DeliverableCompositionShape) SetName(name string) {
	deliverablecompositionshape.Name = name
}

// Stage puts deliverableconceptshape to the model stage
func (deliverableconceptshape *DeliverableConceptShape) Stage(stage *Stage) *DeliverableConceptShape {
	__gong__stage(stage.DeliverableConceptShapes, stage.DeliverableConceptShape_stagedOrder, stage.DeliverableConceptShape_orderStaged, &stage.DeliverableConceptShapeOrder, stage.DeliverableConceptShapes_mapString, deliverableconceptshape, deliverableconceptshape.Name)
	return deliverableconceptshape
}

// StagePreserveOrder puts deliverableconceptshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DeliverableConceptShapeOrder
// - update stage.DeliverableConceptShapeOrder accordingly
func (deliverableconceptshape *DeliverableConceptShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DeliverableConceptShapes, stage.DeliverableConceptShape_stagedOrder, stage.DeliverableConceptShape_orderStaged, &stage.DeliverableConceptShapeOrder, stage.DeliverableConceptShapes_mapString, deliverableconceptshape, order, deliverableconceptshape.Name)
}

// Unstage removes deliverableconceptshape off the model stage
func (deliverableconceptshape *DeliverableConceptShape) Unstage(stage *Stage) *DeliverableConceptShape {
	__gong__unstage(stage.DeliverableConceptShapes, stage.DeliverableConceptShapes_mapString, deliverableconceptshape, deliverableconceptshape.Name)
	return deliverableconceptshape
}

// UnstageVoid removes deliverableconceptshape off the model stage
func (deliverableconceptshape *DeliverableConceptShape) UnstageVoid(stage *Stage) {
	deliverableconceptshape.Unstage(stage)
}

func (deliverableconceptshape *DeliverableConceptShape) StageVoid(stage *Stage) {
	deliverableconceptshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (deliverableconceptshape *DeliverableConceptShape) GetName() (res string) {
	return deliverableconceptshape.Name
}

// for satisfaction of GongStruct interface
func (deliverableconceptshape *DeliverableConceptShape) SetName(name string) {
	deliverableconceptshape.Name = name
}

// Stage puts deliverableshape to the model stage
func (deliverableshape *DeliverableShape) Stage(stage *Stage) *DeliverableShape {
	__gong__stage(stage.DeliverableShapes, stage.DeliverableShape_stagedOrder, stage.DeliverableShape_orderStaged, &stage.DeliverableShapeOrder, stage.DeliverableShapes_mapString, deliverableshape, deliverableshape.Name)
	return deliverableshape
}

// StagePreserveOrder puts deliverableshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DeliverableShapeOrder
// - update stage.DeliverableShapeOrder accordingly
func (deliverableshape *DeliverableShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DeliverableShapes, stage.DeliverableShape_stagedOrder, stage.DeliverableShape_orderStaged, &stage.DeliverableShapeOrder, stage.DeliverableShapes_mapString, deliverableshape, order, deliverableshape.Name)
}

// Unstage removes deliverableshape off the model stage
func (deliverableshape *DeliverableShape) Unstage(stage *Stage) *DeliverableShape {
	__gong__unstage(stage.DeliverableShapes, stage.DeliverableShapes_mapString, deliverableshape, deliverableshape.Name)
	return deliverableshape
}

// UnstageVoid removes deliverableshape off the model stage
func (deliverableshape *DeliverableShape) UnstageVoid(stage *Stage) {
	deliverableshape.Unstage(stage)
}

func (deliverableshape *DeliverableShape) StageVoid(stage *Stage) {
	deliverableshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (deliverableshape *DeliverableShape) GetName() (res string) {
	return deliverableshape.Name
}

// for satisfaction of GongStruct interface
func (deliverableshape *DeliverableShape) SetName(name string) {
	deliverableshape.Name = name
}

// Stage puts diagram to the model stage
func (diagram *Diagram) Stage(stage *Stage) *Diagram {
	__gong__stage(stage.Diagrams, stage.Diagram_stagedOrder, stage.Diagram_orderStaged, &stage.DiagramOrder, stage.Diagrams_mapString, diagram, diagram.Name)
	return diagram
}

// StagePreserveOrder puts diagram to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramOrder
// - update stage.DiagramOrder accordingly
func (diagram *Diagram) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Diagrams, stage.Diagram_stagedOrder, stage.Diagram_orderStaged, &stage.DiagramOrder, stage.Diagrams_mapString, diagram, order, diagram.Name)
}

// Unstage removes diagram off the model stage
func (diagram *Diagram) Unstage(stage *Stage) *Diagram {
	__gong__unstage(stage.Diagrams, stage.Diagrams_mapString, diagram, diagram.Name)
	return diagram
}

// UnstageVoid removes diagram off the model stage
func (diagram *Diagram) UnstageVoid(stage *Stage) {
	diagram.Unstage(stage)
}

func (diagram *Diagram) StageVoid(stage *Stage) {
	diagram.Stage(stage)
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) GetName() (res string) {
	return diagram.Name
}

// for satisfaction of GongStruct interface
func (diagram *Diagram) SetName(name string) {
	diagram.Name = name
}

// Stage puts diagramshape to the model stage
func (diagramshape *DiagramShape) Stage(stage *Stage) *DiagramShape {
	__gong__stage(stage.DiagramShapes, stage.DiagramShape_stagedOrder, stage.DiagramShape_orderStaged, &stage.DiagramShapeOrder, stage.DiagramShapes_mapString, diagramshape, diagramshape.Name)
	return diagramshape
}

// StagePreserveOrder puts diagramshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramShapeOrder
// - update stage.DiagramShapeOrder accordingly
func (diagramshape *DiagramShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DiagramShapes, stage.DiagramShape_stagedOrder, stage.DiagramShape_orderStaged, &stage.DiagramShapeOrder, stage.DiagramShapes_mapString, diagramshape, order, diagramshape.Name)
}

// Unstage removes diagramshape off the model stage
func (diagramshape *DiagramShape) Unstage(stage *Stage) *DiagramShape {
	__gong__unstage(stage.DiagramShapes, stage.DiagramShapes_mapString, diagramshape, diagramshape.Name)
	return diagramshape
}

// UnstageVoid removes diagramshape off the model stage
func (diagramshape *DiagramShape) UnstageVoid(stage *Stage) {
	diagramshape.Unstage(stage)
}

func (diagramshape *DiagramShape) StageVoid(stage *Stage) {
	diagramshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (diagramshape *DiagramShape) GetName() (res string) {
	return diagramshape.Name
}

// for satisfaction of GongStruct interface
func (diagramshape *DiagramShape) SetName(name string) {
	diagramshape.Name = name
}

// Stage puts library to the model stage
func (library *Library) Stage(stage *Stage) *Library {
	__gong__stage(stage.Librarys, stage.Library_stagedOrder, stage.Library_orderStaged, &stage.LibraryOrder, stage.Librarys_mapString, library, library.Name)
	return library
}

// StagePreserveOrder puts library to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LibraryOrder
// - update stage.LibraryOrder accordingly
func (library *Library) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Librarys, stage.Library_stagedOrder, stage.Library_orderStaged, &stage.LibraryOrder, stage.Librarys_mapString, library, order, library.Name)
}

// Unstage removes library off the model stage
func (library *Library) Unstage(stage *Stage) *Library {
	__gong__unstage(stage.Librarys, stage.Librarys_mapString, library, library.Name)
	return library
}

// UnstageVoid removes library off the model stage
func (library *Library) UnstageVoid(stage *Stage) {
	library.Unstage(stage)
}

func (library *Library) StageVoid(stage *Stage) {
	library.Stage(stage)
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
	__gong__stage(stage.Notes, stage.Note_stagedOrder, stage.Note_orderStaged, &stage.NoteOrder, stage.Notes_mapString, note, note.Name)
	return note
}

// StagePreserveOrder puts note to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteOrder
// - update stage.NoteOrder accordingly
func (note *Note) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Notes, stage.Note_stagedOrder, stage.Note_orderStaged, &stage.NoteOrder, stage.Notes_mapString, note, order, note.Name)
}

// Unstage removes note off the model stage
func (note *Note) Unstage(stage *Stage) *Note {
	__gong__unstage(stage.Notes, stage.Notes_mapString, note, note.Name)
	return note
}

// UnstageVoid removes note off the model stage
func (note *Note) UnstageVoid(stage *Stage) {
	note.Unstage(stage)
}

func (note *Note) StageVoid(stage *Stage) {
	note.Stage(stage)
}

// for satisfaction of GongStruct interface
func (note *Note) GetName() (res string) {
	return note.Name
}

// for satisfaction of GongStruct interface
func (note *Note) SetName(name string) {
	note.Name = name
}

// Stage puts notedeliverableshape to the model stage
func (notedeliverableshape *NoteDeliverableShape) Stage(stage *Stage) *NoteDeliverableShape {
	__gong__stage(stage.NoteDeliverableShapes, stage.NoteDeliverableShape_stagedOrder, stage.NoteDeliverableShape_orderStaged, &stage.NoteDeliverableShapeOrder, stage.NoteDeliverableShapes_mapString, notedeliverableshape, notedeliverableshape.Name)
	return notedeliverableshape
}

// StagePreserveOrder puts notedeliverableshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteDeliverableShapeOrder
// - update stage.NoteDeliverableShapeOrder accordingly
func (notedeliverableshape *NoteDeliverableShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteDeliverableShapes, stage.NoteDeliverableShape_stagedOrder, stage.NoteDeliverableShape_orderStaged, &stage.NoteDeliverableShapeOrder, stage.NoteDeliverableShapes_mapString, notedeliverableshape, order, notedeliverableshape.Name)
}

// Unstage removes notedeliverableshape off the model stage
func (notedeliverableshape *NoteDeliverableShape) Unstage(stage *Stage) *NoteDeliverableShape {
	__gong__unstage(stage.NoteDeliverableShapes, stage.NoteDeliverableShapes_mapString, notedeliverableshape, notedeliverableshape.Name)
	return notedeliverableshape
}

// UnstageVoid removes notedeliverableshape off the model stage
func (notedeliverableshape *NoteDeliverableShape) UnstageVoid(stage *Stage) {
	notedeliverableshape.Unstage(stage)
}

func (notedeliverableshape *NoteDeliverableShape) StageVoid(stage *Stage) {
	notedeliverableshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notedeliverableshape *NoteDeliverableShape) GetName() (res string) {
	return notedeliverableshape.Name
}

// for satisfaction of GongStruct interface
func (notedeliverableshape *NoteDeliverableShape) SetName(name string) {
	notedeliverableshape.Name = name
}

// Stage puts noteshape to the model stage
func (noteshape *NoteShape) Stage(stage *Stage) *NoteShape {
	__gong__stage(stage.NoteShapes, stage.NoteShape_stagedOrder, stage.NoteShape_orderStaged, &stage.NoteShapeOrder, stage.NoteShapes_mapString, noteshape, noteshape.Name)
	return noteshape
}

// StagePreserveOrder puts noteshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteShapeOrder
// - update stage.NoteShapeOrder accordingly
func (noteshape *NoteShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteShapes, stage.NoteShape_stagedOrder, stage.NoteShape_orderStaged, &stage.NoteShapeOrder, stage.NoteShapes_mapString, noteshape, order, noteshape.Name)
}

// Unstage removes noteshape off the model stage
func (noteshape *NoteShape) Unstage(stage *Stage) *NoteShape {
	__gong__unstage(stage.NoteShapes, stage.NoteShapes_mapString, noteshape, noteshape.Name)
	return noteshape
}

// UnstageVoid removes noteshape off the model stage
func (noteshape *NoteShape) UnstageVoid(stage *Stage) {
	noteshape.Unstage(stage)
}

func (noteshape *NoteShape) StageVoid(stage *Stage) {
	noteshape.Stage(stage)
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
	__gong__stage(stage.NoteStakeholderShapes, stage.NoteStakeholderShape_stagedOrder, stage.NoteStakeholderShape_orderStaged, &stage.NoteStakeholderShapeOrder, stage.NoteStakeholderShapes_mapString, notestakeholdershape, notestakeholdershape.Name)
	return notestakeholdershape
}

// StagePreserveOrder puts notestakeholdershape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteStakeholderShapeOrder
// - update stage.NoteStakeholderShapeOrder accordingly
func (notestakeholdershape *NoteStakeholderShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteStakeholderShapes, stage.NoteStakeholderShape_stagedOrder, stage.NoteStakeholderShape_orderStaged, &stage.NoteStakeholderShapeOrder, stage.NoteStakeholderShapes_mapString, notestakeholdershape, order, notestakeholdershape.Name)
}

// Unstage removes notestakeholdershape off the model stage
func (notestakeholdershape *NoteStakeholderShape) Unstage(stage *Stage) *NoteStakeholderShape {
	__gong__unstage(stage.NoteStakeholderShapes, stage.NoteStakeholderShapes_mapString, notestakeholdershape, notestakeholdershape.Name)
	return notestakeholdershape
}

// UnstageVoid removes notestakeholdershape off the model stage
func (notestakeholdershape *NoteStakeholderShape) UnstageVoid(stage *Stage) {
	notestakeholdershape.Unstage(stage)
}

func (notestakeholdershape *NoteStakeholderShape) StageVoid(stage *Stage) {
	notestakeholdershape.Stage(stage)
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
	__gong__stage(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder, stage.NoteTaskShape_orderStaged, &stage.NoteTaskShapeOrder, stage.NoteTaskShapes_mapString, notetaskshape, notetaskshape.Name)
	return notetaskshape
}

// StagePreserveOrder puts notetaskshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteTaskShapeOrder
// - update stage.NoteTaskShapeOrder accordingly
func (notetaskshape *NoteTaskShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder, stage.NoteTaskShape_orderStaged, &stage.NoteTaskShapeOrder, stage.NoteTaskShapes_mapString, notetaskshape, order, notetaskshape.Name)
}

// Unstage removes notetaskshape off the model stage
func (notetaskshape *NoteTaskShape) Unstage(stage *Stage) *NoteTaskShape {
	__gong__unstage(stage.NoteTaskShapes, stage.NoteTaskShapes_mapString, notetaskshape, notetaskshape.Name)
	return notetaskshape
}

// UnstageVoid removes notetaskshape off the model stage
func (notetaskshape *NoteTaskShape) UnstageVoid(stage *Stage) {
	notetaskshape.Unstage(stage)
}

func (notetaskshape *NoteTaskShape) StageVoid(stage *Stage) {
	notetaskshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notetaskshape *NoteTaskShape) GetName() (res string) {
	return notetaskshape.Name
}

// for satisfaction of GongStruct interface
func (notetaskshape *NoteTaskShape) SetName(name string) {
	notetaskshape.Name = name
}

// Stage puts requirement to the model stage
func (requirement *Requirement) Stage(stage *Stage) *Requirement {
	__gong__stage(stage.Requirements, stage.Requirement_stagedOrder, stage.Requirement_orderStaged, &stage.RequirementOrder, stage.Requirements_mapString, requirement, requirement.Name)
	return requirement
}

// StagePreserveOrder puts requirement to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RequirementOrder
// - update stage.RequirementOrder accordingly
func (requirement *Requirement) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Requirements, stage.Requirement_stagedOrder, stage.Requirement_orderStaged, &stage.RequirementOrder, stage.Requirements_mapString, requirement, order, requirement.Name)
}

// Unstage removes requirement off the model stage
func (requirement *Requirement) Unstage(stage *Stage) *Requirement {
	__gong__unstage(stage.Requirements, stage.Requirements_mapString, requirement, requirement.Name)
	return requirement
}

// UnstageVoid removes requirement off the model stage
func (requirement *Requirement) UnstageVoid(stage *Stage) {
	requirement.Unstage(stage)
}

func (requirement *Requirement) StageVoid(stage *Stage) {
	requirement.Stage(stage)
}

// for satisfaction of GongStruct interface
func (requirement *Requirement) GetName() (res string) {
	return requirement.Name
}

// for satisfaction of GongStruct interface
func (requirement *Requirement) SetName(name string) {
	requirement.Name = name
}

// Stage puts requirementshape to the model stage
func (requirementshape *RequirementShape) Stage(stage *Stage) *RequirementShape {
	__gong__stage(stage.RequirementShapes, stage.RequirementShape_stagedOrder, stage.RequirementShape_orderStaged, &stage.RequirementShapeOrder, stage.RequirementShapes_mapString, requirementshape, requirementshape.Name)
	return requirementshape
}

// StagePreserveOrder puts requirementshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RequirementShapeOrder
// - update stage.RequirementShapeOrder accordingly
func (requirementshape *RequirementShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.RequirementShapes, stage.RequirementShape_stagedOrder, stage.RequirementShape_orderStaged, &stage.RequirementShapeOrder, stage.RequirementShapes_mapString, requirementshape, order, requirementshape.Name)
}

// Unstage removes requirementshape off the model stage
func (requirementshape *RequirementShape) Unstage(stage *Stage) *RequirementShape {
	__gong__unstage(stage.RequirementShapes, stage.RequirementShapes_mapString, requirementshape, requirementshape.Name)
	return requirementshape
}

// UnstageVoid removes requirementshape off the model stage
func (requirementshape *RequirementShape) UnstageVoid(stage *Stage) {
	requirementshape.Unstage(stage)
}

func (requirementshape *RequirementShape) StageVoid(stage *Stage) {
	requirementshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (requirementshape *RequirementShape) GetName() (res string) {
	return requirementshape.Name
}

// for satisfaction of GongStruct interface
func (requirementshape *RequirementShape) SetName(name string) {
	requirementshape.Name = name
}

// Stage puts stakeholder to the model stage
func (stakeholder *Stakeholder) Stage(stage *Stage) *Stakeholder {
	__gong__stage(stage.Stakeholders, stage.Stakeholder_stagedOrder, stage.Stakeholder_orderStaged, &stage.StakeholderOrder, stage.Stakeholders_mapString, stakeholder, stakeholder.Name)
	return stakeholder
}

// StagePreserveOrder puts stakeholder to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StakeholderOrder
// - update stage.StakeholderOrder accordingly
func (stakeholder *Stakeholder) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Stakeholders, stage.Stakeholder_stagedOrder, stage.Stakeholder_orderStaged, &stage.StakeholderOrder, stage.Stakeholders_mapString, stakeholder, order, stakeholder.Name)
}

// Unstage removes stakeholder off the model stage
func (stakeholder *Stakeholder) Unstage(stage *Stage) *Stakeholder {
	__gong__unstage(stage.Stakeholders, stage.Stakeholders_mapString, stakeholder, stakeholder.Name)
	return stakeholder
}

// UnstageVoid removes stakeholder off the model stage
func (stakeholder *Stakeholder) UnstageVoid(stage *Stage) {
	stakeholder.Unstage(stage)
}

func (stakeholder *Stakeholder) StageVoid(stage *Stage) {
	stakeholder.Stage(stage)
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
	__gong__stage(stage.StakeholderCompositionShapes, stage.StakeholderCompositionShape_stagedOrder, stage.StakeholderCompositionShape_orderStaged, &stage.StakeholderCompositionShapeOrder, stage.StakeholderCompositionShapes_mapString, stakeholdercompositionshape, stakeholdercompositionshape.Name)
	return stakeholdercompositionshape
}

// StagePreserveOrder puts stakeholdercompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StakeholderCompositionShapeOrder
// - update stage.StakeholderCompositionShapeOrder accordingly
func (stakeholdercompositionshape *StakeholderCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StakeholderCompositionShapes, stage.StakeholderCompositionShape_stagedOrder, stage.StakeholderCompositionShape_orderStaged, &stage.StakeholderCompositionShapeOrder, stage.StakeholderCompositionShapes_mapString, stakeholdercompositionshape, order, stakeholdercompositionshape.Name)
}

// Unstage removes stakeholdercompositionshape off the model stage
func (stakeholdercompositionshape *StakeholderCompositionShape) Unstage(stage *Stage) *StakeholderCompositionShape {
	__gong__unstage(stage.StakeholderCompositionShapes, stage.StakeholderCompositionShapes_mapString, stakeholdercompositionshape, stakeholdercompositionshape.Name)
	return stakeholdercompositionshape
}

// UnstageVoid removes stakeholdercompositionshape off the model stage
func (stakeholdercompositionshape *StakeholderCompositionShape) UnstageVoid(stage *Stage) {
	stakeholdercompositionshape.Unstage(stage)
}

func (stakeholdercompositionshape *StakeholderCompositionShape) StageVoid(stage *Stage) {
	stakeholdercompositionshape.Stage(stage)
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
	__gong__stage(stage.StakeholderConcernShapes, stage.StakeholderConcernShape_stagedOrder, stage.StakeholderConcernShape_orderStaged, &stage.StakeholderConcernShapeOrder, stage.StakeholderConcernShapes_mapString, stakeholderconcernshape, stakeholderconcernshape.Name)
	return stakeholderconcernshape
}

// StagePreserveOrder puts stakeholderconcernshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StakeholderConcernShapeOrder
// - update stage.StakeholderConcernShapeOrder accordingly
func (stakeholderconcernshape *StakeholderConcernShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StakeholderConcernShapes, stage.StakeholderConcernShape_stagedOrder, stage.StakeholderConcernShape_orderStaged, &stage.StakeholderConcernShapeOrder, stage.StakeholderConcernShapes_mapString, stakeholderconcernshape, order, stakeholderconcernshape.Name)
}

// Unstage removes stakeholderconcernshape off the model stage
func (stakeholderconcernshape *StakeholderConcernShape) Unstage(stage *Stage) *StakeholderConcernShape {
	__gong__unstage(stage.StakeholderConcernShapes, stage.StakeholderConcernShapes_mapString, stakeholderconcernshape, stakeholderconcernshape.Name)
	return stakeholderconcernshape
}

// UnstageVoid removes stakeholderconcernshape off the model stage
func (stakeholderconcernshape *StakeholderConcernShape) UnstageVoid(stage *Stage) {
	stakeholderconcernshape.Unstage(stage)
}

func (stakeholderconcernshape *StakeholderConcernShape) StageVoid(stage *Stage) {
	stakeholderconcernshape.Stage(stage)
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
	__gong__stage(stage.StakeholderShapes, stage.StakeholderShape_stagedOrder, stage.StakeholderShape_orderStaged, &stage.StakeholderShapeOrder, stage.StakeholderShapes_mapString, stakeholdershape, stakeholdershape.Name)
	return stakeholdershape
}

// StagePreserveOrder puts stakeholdershape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StakeholderShapeOrder
// - update stage.StakeholderShapeOrder accordingly
func (stakeholdershape *StakeholderShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StakeholderShapes, stage.StakeholderShape_stagedOrder, stage.StakeholderShape_orderStaged, &stage.StakeholderShapeOrder, stage.StakeholderShapes_mapString, stakeholdershape, order, stakeholdershape.Name)
}

// Unstage removes stakeholdershape off the model stage
func (stakeholdershape *StakeholderShape) Unstage(stage *Stage) *StakeholderShape {
	__gong__unstage(stage.StakeholderShapes, stage.StakeholderShapes_mapString, stakeholdershape, stakeholdershape.Name)
	return stakeholdershape
}

// UnstageVoid removes stakeholdershape off the model stage
func (stakeholdershape *StakeholderShape) UnstageVoid(stage *Stage) {
	stakeholdershape.Unstage(stage)
}

func (stakeholdershape *StakeholderShape) StageVoid(stage *Stage) {
	stakeholdershape.Stage(stage)
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
	__gong__stage(stage.SupportLevels, stage.SupportLevel_stagedOrder, stage.SupportLevel_orderStaged, &stage.SupportLevelOrder, stage.SupportLevels_mapString, supportlevel, supportlevel.Name)
	return supportlevel
}

// StagePreserveOrder puts supportlevel to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SupportLevelOrder
// - update stage.SupportLevelOrder accordingly
func (supportlevel *SupportLevel) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SupportLevels, stage.SupportLevel_stagedOrder, stage.SupportLevel_orderStaged, &stage.SupportLevelOrder, stage.SupportLevels_mapString, supportlevel, order, supportlevel.Name)
}

// Unstage removes supportlevel off the model stage
func (supportlevel *SupportLevel) Unstage(stage *Stage) *SupportLevel {
	__gong__unstage(stage.SupportLevels, stage.SupportLevels_mapString, supportlevel, supportlevel.Name)
	return supportlevel
}

// UnstageVoid removes supportlevel off the model stage
func (supportlevel *SupportLevel) UnstageVoid(stage *Stage) {
	supportlevel.Unstage(stage)
}

func (supportlevel *SupportLevel) StageVoid(stage *Stage) {
	supportlevel.Stage(stage)
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
	__gong__stage(stage.Tools, stage.Tool_stagedOrder, stage.Tool_orderStaged, &stage.ToolOrder, stage.Tools_mapString, tool, tool.Name)
	return tool
}

// StagePreserveOrder puts tool to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ToolOrder
// - update stage.ToolOrder accordingly
func (tool *Tool) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Tools, stage.Tool_stagedOrder, stage.Tool_orderStaged, &stage.ToolOrder, stage.Tools_mapString, tool, order, tool.Name)
}

// Unstage removes tool off the model stage
func (tool *Tool) Unstage(stage *Stage) *Tool {
	__gong__unstage(stage.Tools, stage.Tools_mapString, tool, tool.Name)
	return tool
}

// UnstageVoid removes tool off the model stage
func (tool *Tool) UnstageVoid(stage *Stage) {
	tool.Unstage(stage)
}

func (tool *Tool) StageVoid(stage *Stage) {
	tool.Stage(stage)
}

// for satisfaction of GongStruct interface
func (tool *Tool) GetName() (res string) {
	return tool.Name
}

// for satisfaction of GongStruct interface
func (tool *Tool) SetName(name string) {
	tool.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.AnalysisNeeds, &stage.AnalysisNeeds_mapString, &stage.AnalysisNeed_stagedOrder, &stage.AnalysisNeedOrder)

	__gong__resetStageType(&stage.Concepts, &stage.Concepts_mapString, &stage.Concept_stagedOrder, &stage.ConceptOrder)

	__gong__resetStageType(&stage.ConceptShapes, &stage.ConceptShapes_mapString, &stage.ConceptShape_stagedOrder, &stage.ConceptShapeOrder)

	__gong__resetStageType(&stage.Concerns, &stage.Concerns_mapString, &stage.Concern_stagedOrder, &stage.ConcernOrder)

	__gong__resetStageType(&stage.ConcernCompositionShapes, &stage.ConcernCompositionShapes_mapString, &stage.ConcernCompositionShape_stagedOrder, &stage.ConcernCompositionShapeOrder)

	__gong__resetStageType(&stage.ConcernInputShapes, &stage.ConcernInputShapes_mapString, &stage.ConcernInputShape_stagedOrder, &stage.ConcernInputShapeOrder)

	__gong__resetStageType(&stage.ConcernOutputShapes, &stage.ConcernOutputShapes_mapString, &stage.ConcernOutputShape_stagedOrder, &stage.ConcernOutputShapeOrder)

	__gong__resetStageType(&stage.ConcernShapes, &stage.ConcernShapes_mapString, &stage.ConcernShape_stagedOrder, &stage.ConcernShapeOrder)

	__gong__resetStageType(&stage.ControlPointShapes, &stage.ControlPointShapes_mapString, &stage.ControlPointShape_stagedOrder, &stage.ControlPointShapeOrder)

	__gong__resetStageType(&stage.Deliverables, &stage.Deliverables_mapString, &stage.Deliverable_stagedOrder, &stage.DeliverableOrder)

	__gong__resetStageType(&stage.DeliverableCompositionShapes, &stage.DeliverableCompositionShapes_mapString, &stage.DeliverableCompositionShape_stagedOrder, &stage.DeliverableCompositionShapeOrder)

	__gong__resetStageType(&stage.DeliverableConceptShapes, &stage.DeliverableConceptShapes_mapString, &stage.DeliverableConceptShape_stagedOrder, &stage.DeliverableConceptShapeOrder)

	__gong__resetStageType(&stage.DeliverableShapes, &stage.DeliverableShapes_mapString, &stage.DeliverableShape_stagedOrder, &stage.DeliverableShapeOrder)

	__gong__resetStageType(&stage.Diagrams, &stage.Diagrams_mapString, &stage.Diagram_stagedOrder, &stage.DiagramOrder)

	__gong__resetStageType(&stage.DiagramShapes, &stage.DiagramShapes_mapString, &stage.DiagramShape_stagedOrder, &stage.DiagramShapeOrder)

	__gong__resetStageType(&stage.Librarys, &stage.Librarys_mapString, &stage.Library_stagedOrder, &stage.LibraryOrder)

	__gong__resetStageType(&stage.Notes, &stage.Notes_mapString, &stage.Note_stagedOrder, &stage.NoteOrder)

	__gong__resetStageType(&stage.NoteDeliverableShapes, &stage.NoteDeliverableShapes_mapString, &stage.NoteDeliverableShape_stagedOrder, &stage.NoteDeliverableShapeOrder)

	__gong__resetStageType(&stage.NoteShapes, &stage.NoteShapes_mapString, &stage.NoteShape_stagedOrder, &stage.NoteShapeOrder)

	__gong__resetStageType(&stage.NoteStakeholderShapes, &stage.NoteStakeholderShapes_mapString, &stage.NoteStakeholderShape_stagedOrder, &stage.NoteStakeholderShapeOrder)

	__gong__resetStageType(&stage.NoteTaskShapes, &stage.NoteTaskShapes_mapString, &stage.NoteTaskShape_stagedOrder, &stage.NoteTaskShapeOrder)

	__gong__resetStageType(&stage.Requirements, &stage.Requirements_mapString, &stage.Requirement_stagedOrder, &stage.RequirementOrder)

	__gong__resetStageType(&stage.RequirementShapes, &stage.RequirementShapes_mapString, &stage.RequirementShape_stagedOrder, &stage.RequirementShapeOrder)

	__gong__resetStageType(&stage.Stakeholders, &stage.Stakeholders_mapString, &stage.Stakeholder_stagedOrder, &stage.StakeholderOrder)

	__gong__resetStageType(&stage.StakeholderCompositionShapes, &stage.StakeholderCompositionShapes_mapString, &stage.StakeholderCompositionShape_stagedOrder, &stage.StakeholderCompositionShapeOrder)

	__gong__resetStageType(&stage.StakeholderConcernShapes, &stage.StakeholderConcernShapes_mapString, &stage.StakeholderConcernShape_stagedOrder, &stage.StakeholderConcernShapeOrder)

	__gong__resetStageType(&stage.StakeholderShapes, &stage.StakeholderShapes_mapString, &stage.StakeholderShape_stagedOrder, &stage.StakeholderShapeOrder)

	__gong__resetStageType(&stage.SupportLevels, &stage.SupportLevels_mapString, &stage.SupportLevel_stagedOrder, &stage.SupportLevelOrder)

	__gong__resetStageType(&stage.Tools, &stage.Tools_mapString, &stage.Tool_stagedOrder, &stage.ToolOrder)

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct any

type GongstructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

type GongtructBasicField = GongstructBasicField

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) string
	GongGetUUID(stage *Stage) string
	GongAfterCreateFromFront(stage *Stage)
	GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF)
	GongAfterDeleteFromFront(stage *Stage, front GongstructIF)
	GongIsStaged(stage *Stage) bool
	GongStageBranch(stage *Stage)
	GongUnstageBranch(stage *Stage)
}
type GongstructPtr interface {
	GongstructIF
	comparable
}

type PointerToGongstruct = GongstructPtr

func GongCompareGongstructByName[T GongstructPtr](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func GongSortGongstructSetByName[T GongstructPtr](set map[T]struct{}) (sortedSlice []T) {
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, GongCompareGongstructByName)

	return
}

// GetInstancesSorted is the Stage method returning sorted instances of a gongstruct.
func (stage *Stage) GetInstancesSorted[T GongstructPtr]() (sortedSlice []T) {
	set := stage.GetInstancesSet[T]()
	sortedSlice = GongSortGongstructSetByName(*set)

	return
}

// GetInstancesMapByName is the Stage method returning a map of staged instances by their name.
func (stage *Stage) GetInstancesMapByName[Type GongstructIF]() map[string]Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AnalysisNeed:
		return any(stage.AnalysisNeeds_mapString).(map[string]Type)
	case *Concept:
		return any(stage.Concepts_mapString).(map[string]Type)
	case *ConceptShape:
		return any(stage.ConceptShapes_mapString).(map[string]Type)
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
	case *ControlPointShape:
		return any(stage.ControlPointShapes_mapString).(map[string]Type)
	case *Deliverable:
		return any(stage.Deliverables_mapString).(map[string]Type)
	case *DeliverableCompositionShape:
		return any(stage.DeliverableCompositionShapes_mapString).(map[string]Type)
	case *DeliverableConceptShape:
		return any(stage.DeliverableConceptShapes_mapString).(map[string]Type)
	case *DeliverableShape:
		return any(stage.DeliverableShapes_mapString).(map[string]Type)
	case *Diagram:
		return any(stage.Diagrams_mapString).(map[string]Type)
	case *DiagramShape:
		return any(stage.DiagramShapes_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Note:
		return any(stage.Notes_mapString).(map[string]Type)
	case *NoteDeliverableShape:
		return any(stage.NoteDeliverableShapes_mapString).(map[string]Type)
	case *NoteShape:
		return any(stage.NoteShapes_mapString).(map[string]Type)
	case *NoteStakeholderShape:
		return any(stage.NoteStakeholderShapes_mapString).(map[string]Type)
	case *NoteTaskShape:
		return any(stage.NoteTaskShapes_mapString).(map[string]Type)
	case *Requirement:
		return any(stage.Requirements_mapString).(map[string]Type)
	case *RequirementShape:
		return any(stage.RequirementShapes_mapString).(map[string]Type)
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

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AnalysisNeed:
		return any(&stage.AnalysisNeeds).(*map[Type]struct{})
	case *Concept:
		return any(&stage.Concepts).(*map[Type]struct{})
	case *ConceptShape:
		return any(&stage.ConceptShapes).(*map[Type]struct{})
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
	case *ControlPointShape:
		return any(&stage.ControlPointShapes).(*map[Type]struct{})
	case *Deliverable:
		return any(&stage.Deliverables).(*map[Type]struct{})
	case *DeliverableCompositionShape:
		return any(&stage.DeliverableCompositionShapes).(*map[Type]struct{})
	case *DeliverableConceptShape:
		return any(&stage.DeliverableConceptShapes).(*map[Type]struct{})
	case *DeliverableShape:
		return any(&stage.DeliverableShapes).(*map[Type]struct{})
	case *Diagram:
		return any(&stage.Diagrams).(*map[Type]struct{})
	case *DiagramShape:
		return any(&stage.DiagramShapes).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Note:
		return any(&stage.Notes).(*map[Type]struct{})
	case *NoteDeliverableShape:
		return any(&stage.NoteDeliverableShapes).(*map[Type]struct{})
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]struct{})
	case *NoteStakeholderShape:
		return any(&stage.NoteStakeholderShapes).(*map[Type]struct{})
	case *NoteTaskShape:
		return any(&stage.NoteTaskShapes).(*map[Type]struct{})
	case *Requirement:
		return any(&stage.Requirements).(*map[Type]struct{})
	case *RequirementShape:
		return any(&stage.RequirementShapes).(*map[Type]struct{})
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

// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Concept:
		return any(&Concept{
			Tools: []*Tool{{Name: "Tools"}},
		}).(*Type)
	case ConceptShape:
		return any(&ConceptShape{
			Concept: &Concept{Name: "Concept"},
		}).(*Type)
	case Concern:
		return any(&Concern{
			SubConcerns: []*Concern{{Name: "SubConcerns"}},
			Inputs: []*Deliverable{{Name: "Inputs"}},
			Outputs: []*Deliverable{{Name: "Outputs"}},
			Requirements: []*Requirement{{Name: "Requirements"}},
		}).(*Type)
	case ConcernCompositionShape:
		return any(&ConcernCompositionShape{
			Concern: &Concern{Name: "Concern"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case ConcernInputShape:
		return any(&ConcernInputShape{
			Deliverable: &Deliverable{Name: "Deliverable"},
			Concern: &Concern{Name: "Concern"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case ConcernOutputShape:
		return any(&ConcernOutputShape{
			Concern: &Concern{Name: "Concern"},
			Deliverable: &Deliverable{Name: "Deliverable"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case ConcernShape:
		return any(&ConcernShape{
			Concern: &Concern{Name: "Concern"},
		}).(*Type)
	case Deliverable:
		return any(&Deliverable{
			SubDeliverables: []*Deliverable{{Name: "SubDeliverables"}},
			Concepts: []*Concept{{Name: "Concepts"}},
		}).(*Type)
	case DeliverableCompositionShape:
		return any(&DeliverableCompositionShape{
			Deliverable: &Deliverable{Name: "Deliverable"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case DeliverableConceptShape:
		return any(&DeliverableConceptShape{
			Deliverable: &Deliverable{Name: "Deliverable"},
			Concept: &Concept{Name: "Concept"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case DeliverableShape:
		return any(&DeliverableShape{
			Deliverable: &Deliverable{Name: "Deliverable"},
		}).(*Type)
	case Diagram:
		return any(&Diagram{
			ConcernsWhoseRequirementsNodeIsExpanded: []*Concern{{Name: "ConcernsWhoseRequirementsNodeIsExpanded"}},
			Deliverable_Shapes: []*DeliverableShape{{Name: "Deliverable_Shapes"}},
			DeliverablesWhoseNodeIsExpanded: []*Deliverable{{Name: "DeliverablesWhoseNodeIsExpanded"}},
			DeliverablesWhoseConceptsNodeIsExpanded: []*Deliverable{{Name: "DeliverablesWhoseConceptsNodeIsExpanded"}},
			DeliverableComposition_Shapes: []*DeliverableCompositionShape{{Name: "DeliverableComposition_Shapes"}},
			Concern_Shapes: []*ConcernShape{{Name: "Concern_Shapes"}},
			ConcernsWhoseNodeIsExpanded: []*Concern{{Name: "ConcernsWhoseNodeIsExpanded"}},
			ConcernsWhoseInputNodeIsExpanded: []*Concern{{Name: "ConcernsWhoseInputNodeIsExpanded"}},
			ConcernsWhoseStakeholderNodeIsExpanded: []*Concern{{Name: "ConcernsWhoseStakeholderNodeIsExpanded"}},
			ConcernssWhoseOutputNodeIsExpanded: []*Concern{{Name: "ConcernssWhoseOutputNodeIsExpanded"}},
			ConcernComposition_Shapes: []*ConcernCompositionShape{{Name: "ConcernComposition_Shapes"}},
			ConcernInputShapes: []*ConcernInputShape{{Name: "ConcernInputShapes"}},
			ConcernOutputShapes: []*ConcernOutputShape{{Name: "ConcernOutputShapes"}},
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			NoteDeliverableShapes: []*NoteDeliverableShape{{Name: "NoteDeliverableShapes"}},
			NoteTaskShapes: []*NoteTaskShape{{Name: "NoteTaskShapes"}},
			NoteResourceShapes: []*NoteStakeholderShape{{Name: "NoteResourceShapes"}},
			Stakeholder_Shapes: []*StakeholderShape{{Name: "Stakeholder_Shapes"}},
			ResourcesWhoseNodeIsExpanded: []*Stakeholder{{Name: "ResourcesWhoseNodeIsExpanded"}},
			ResourceComposition_Shapes: []*StakeholderCompositionShape{{Name: "ResourceComposition_Shapes"}},
			StakeholderConcernShapes: []*StakeholderConcernShape{{Name: "StakeholderConcernShapes"}},
			Requirement_Shapes: []*RequirementShape{{Name: "Requirement_Shapes"}},
			RequirementsWhoseNodeIsExpanded: []*Requirement{{Name: "RequirementsWhoseNodeIsExpanded"}},
			Concept_Shapes: []*ConceptShape{{Name: "Concept_Shapes"}},
			ConceptsWhoseNodeIsExpanded: []*Concept{{Name: "ConceptsWhoseNodeIsExpanded"}},
			ConceptsWhoseDeliverablesNodeIsExpanded: []*Concept{{Name: "ConceptsWhoseDeliverablesNodeIsExpanded"}},
			DeliverableConceptShapes: []*DeliverableConceptShape{{Name: "DeliverableConceptShapes"}},
			Diagram_Shapes: []*DiagramShape{{Name: "Diagram_Shapes"}},
			DiagramsWhoseNodeIsExpanded: []*Diagram{{Name: "DiagramsWhoseNodeIsExpanded"}},
		}).(*Type)
	case DiagramShape:
		return any(&DiagramShape{
			Diagram: &Diagram{Name: "Diagram"},
		}).(*Type)
	case Library:
		return any(&Library{
			RootDeliverables: []*Deliverable{{Name: "RootDeliverables"}},
			RootConcerns: []*Concern{{Name: "RootConcerns"}},
			RootStakeholders: []*Stakeholder{{Name: "RootStakeholders"}},
			RootRequirements: []*Requirement{{Name: "RootRequirements"}},
			RootConcepts: []*Concept{{Name: "RootConcepts"}},
			AnalysisNeeds: []*AnalysisNeed{{Name: "AnalysisNeeds"}},
			Notes: []*Note{{Name: "Notes"}},
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			SubLibraries: []*Library{{Name: "SubLibraries"}},
		}).(*Type)
	case Note:
		return any(&Note{
			Deliverables: []*Deliverable{{Name: "Deliverables"}},
			Tasks: []*Concern{{Name: "Tasks"}},
			Resources: []*Stakeholder{{Name: "Resources"}},
		}).(*Type)
	case NoteDeliverableShape:
		return any(&NoteDeliverableShape{
			Note: &Note{Name: "Note"},
			Deliverable: &Deliverable{Name: "Deliverable"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			Note: &Note{Name: "Note"},
		}).(*Type)
	case NoteStakeholderShape:
		return any(&NoteStakeholderShape{
			Note: &Note{Name: "Note"},
			Stakeholder: &Stakeholder{Name: "Stakeholder"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case NoteTaskShape:
		return any(&NoteTaskShape{
			Note: &Note{Name: "Note"},
			Task: &Concern{Name: "Task"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case Requirement:
		return any(&Requirement{
			SupportLevels: []*SupportLevel{{Name: "SupportLevels"}},
			Concepts: []*Concept{{Name: "Concepts"}},
		}).(*Type)
	case RequirementShape:
		return any(&RequirementShape{
			Requirement: &Requirement{Name: "Requirement"},
		}).(*Type)
	case Stakeholder:
		return any(&Stakeholder{
			Concerns: []*Concern{{Name: "Concerns"}},
			SubStakeholders: []*Stakeholder{{Name: "SubStakeholders"}},
		}).(*Type)
	case StakeholderCompositionShape:
		return any(&StakeholderCompositionShape{
			Stakeholder: &Stakeholder{Name: "Stakeholder"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case StakeholderConcernShape:
		return any(&StakeholderConcernShape{
			Stakeholder: &Stakeholder{Name: "Stakeholder"},
			Concern: &Concern{Name: "Concern"},
			ControlPointShapes: []*ControlPointShape{{Name: "ControlPointShapes"}},
		}).(*Type)
	case StakeholderShape:
		return any(&StakeholderShape{
			Stakeholder: &Stakeholder{Name: "Stakeholder"},
		}).(*Type)
	case SupportLevel:
		return any(&SupportLevel{
			Tool: &Tool{Name: "Tool"},
		}).(*Type)
	default:
		return &ret
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
// GetPointerReverseMap is the Stage method for backtrack navigation of pointer associations.
func (stage *Stage) GetPointerReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
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
	// reverse maps of direct associations of ConceptShape
	case ConceptShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Concept":
			res := make(map[*Concept][]*ConceptShape)
			for conceptshape := range stage.ConceptShapes {
				if conceptshape.Concept != nil {
					concept_ := conceptshape.Concept
					var conceptshapes []*ConceptShape
					_, ok := res[concept_]
					if ok {
						conceptshapes = res[concept_]
					} else {
						conceptshapes = make([]*ConceptShape, 0)
					}
					conceptshapes = append(conceptshapes, conceptshape)
					res[concept_] = conceptshapes
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "Concern":
			res := make(map[*Concern][]*ConcernOutputShape)
			for concernoutputshape := range stage.ConcernOutputShapes {
				if concernoutputshape.Concern != nil {
					concern_ := concernoutputshape.Concern
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
		case "Deliverable":
			res := make(map[*Deliverable][]*ConcernOutputShape)
			for concernoutputshape := range stage.ConcernOutputShapes {
				if concernoutputshape.Deliverable != nil {
					deliverable_ := concernoutputshape.Deliverable
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
	// reverse maps of direct associations of ControlPointShape
	case ControlPointShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Deliverable
	case Deliverable:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DeliverableCompositionShape
	case DeliverableCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Deliverable":
			res := make(map[*Deliverable][]*DeliverableCompositionShape)
			for deliverablecompositionshape := range stage.DeliverableCompositionShapes {
				if deliverablecompositionshape.Deliverable != nil {
					deliverable_ := deliverablecompositionshape.Deliverable
					var deliverablecompositionshapes []*DeliverableCompositionShape
					_, ok := res[deliverable_]
					if ok {
						deliverablecompositionshapes = res[deliverable_]
					} else {
						deliverablecompositionshapes = make([]*DeliverableCompositionShape, 0)
					}
					deliverablecompositionshapes = append(deliverablecompositionshapes, deliverablecompositionshape)
					res[deliverable_] = deliverablecompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DeliverableConceptShape
	case DeliverableConceptShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Deliverable":
			res := make(map[*Deliverable][]*DeliverableConceptShape)
			for deliverableconceptshape := range stage.DeliverableConceptShapes {
				if deliverableconceptshape.Deliverable != nil {
					deliverable_ := deliverableconceptshape.Deliverable
					var deliverableconceptshapes []*DeliverableConceptShape
					_, ok := res[deliverable_]
					if ok {
						deliverableconceptshapes = res[deliverable_]
					} else {
						deliverableconceptshapes = make([]*DeliverableConceptShape, 0)
					}
					deliverableconceptshapes = append(deliverableconceptshapes, deliverableconceptshape)
					res[deliverable_] = deliverableconceptshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Concept":
			res := make(map[*Concept][]*DeliverableConceptShape)
			for deliverableconceptshape := range stage.DeliverableConceptShapes {
				if deliverableconceptshape.Concept != nil {
					concept_ := deliverableconceptshape.Concept
					var deliverableconceptshapes []*DeliverableConceptShape
					_, ok := res[concept_]
					if ok {
						deliverableconceptshapes = res[concept_]
					} else {
						deliverableconceptshapes = make([]*DeliverableConceptShape, 0)
					}
					deliverableconceptshapes = append(deliverableconceptshapes, deliverableconceptshape)
					res[concept_] = deliverableconceptshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DeliverableShape
	case DeliverableShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Deliverable":
			res := make(map[*Deliverable][]*DeliverableShape)
			for deliverableshape := range stage.DeliverableShapes {
				if deliverableshape.Deliverable != nil {
					deliverable_ := deliverableshape.Deliverable
					var deliverableshapes []*DeliverableShape
					_, ok := res[deliverable_]
					if ok {
						deliverableshapes = res[deliverable_]
					} else {
						deliverableshapes = make([]*DeliverableShape, 0)
					}
					deliverableshapes = append(deliverableshapes, deliverableshape)
					res[deliverable_] = deliverableshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DiagramShape
	case DiagramShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Diagram":
			res := make(map[*Diagram][]*DiagramShape)
			for diagramshape := range stage.DiagramShapes {
				if diagramshape.Diagram != nil {
					diagram_ := diagramshape.Diagram
					var diagramshapes []*DiagramShape
					_, ok := res[diagram_]
					if ok {
						diagramshapes = res[diagram_]
					} else {
						diagramshapes = make([]*DiagramShape, 0)
					}
					diagramshapes = append(diagramshapes, diagramshape)
					res[diagram_] = diagramshapes
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of NoteDeliverableShape
	case NoteDeliverableShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteDeliverableShape)
			for notedeliverableshape := range stage.NoteDeliverableShapes {
				if notedeliverableshape.Note != nil {
					note_ := notedeliverableshape.Note
					var notedeliverableshapes []*NoteDeliverableShape
					_, ok := res[note_]
					if ok {
						notedeliverableshapes = res[note_]
					} else {
						notedeliverableshapes = make([]*NoteDeliverableShape, 0)
					}
					notedeliverableshapes = append(notedeliverableshapes, notedeliverableshape)
					res[note_] = notedeliverableshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Deliverable":
			res := make(map[*Deliverable][]*NoteDeliverableShape)
			for notedeliverableshape := range stage.NoteDeliverableShapes {
				if notedeliverableshape.Deliverable != nil {
					deliverable_ := notedeliverableshape.Deliverable
					var notedeliverableshapes []*NoteDeliverableShape
					_, ok := res[deliverable_]
					if ok {
						notedeliverableshapes = res[deliverable_]
					} else {
						notedeliverableshapes = make([]*NoteDeliverableShape, 0)
					}
					notedeliverableshapes = append(notedeliverableshapes, notedeliverableshape)
					res[deliverable_] = notedeliverableshapes
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
	// reverse maps of direct associations of Requirement
	case Requirement:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RequirementShape
	case RequirementShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Requirement":
			res := make(map[*Requirement][]*RequirementShape)
			for requirementshape := range stage.RequirementShapes {
				if requirementshape.Requirement != nil {
					requirement_ := requirementshape.Requirement
					var requirementshapes []*RequirementShape
					_, ok := res[requirement_]
					if ok {
						requirementshapes = res[requirement_]
					} else {
						requirementshapes = make([]*RequirementShape, 0)
					}
					requirementshapes = append(requirementshapes, requirementshape)
					res[requirement_] = requirementshapes
				}
			}
			return any(res).(map[*End][]*Start)
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

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
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
	// reverse maps of direct associations of ConceptShape
	case ConceptShape:
		switch fieldname {
		// insertion point for per direct association field
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
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*ConcernCompositionShape)
			for concerncompositionshape := range stage.ConcernCompositionShapes {
				for _, controlpointshape_ := range concerncompositionshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], concerncompositionshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ConcernInputShape
	case ConcernInputShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*ConcernInputShape)
			for concerninputshape := range stage.ConcernInputShapes {
				for _, controlpointshape_ := range concerninputshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], concerninputshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ConcernOutputShape
	case ConcernOutputShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*ConcernOutputShape)
			for concernoutputshape := range stage.ConcernOutputShapes {
				for _, controlpointshape_ := range concernoutputshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], concernoutputshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ConcernShape
	case ConcernShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ControlPointShape
	case ControlPointShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Deliverable
	case Deliverable:
		switch fieldname {
		// insertion point for per direct association field
		case "SubDeliverables":
			res := make(map[*Deliverable][]*Deliverable)
			for deliverable := range stage.Deliverables {
				for _, deliverable_ := range deliverable.SubDeliverables {
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
	// reverse maps of direct associations of DeliverableCompositionShape
	case DeliverableCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*DeliverableCompositionShape)
			for deliverablecompositionshape := range stage.DeliverableCompositionShapes {
				for _, controlpointshape_ := range deliverablecompositionshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], deliverablecompositionshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DeliverableConceptShape
	case DeliverableConceptShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*DeliverableConceptShape)
			for deliverableconceptshape := range stage.DeliverableConceptShapes {
				for _, controlpointshape_ := range deliverableconceptshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], deliverableconceptshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DeliverableShape
	case DeliverableShape:
		switch fieldname {
		// insertion point for per direct association field
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
		case "Deliverable_Shapes":
			res := make(map[*DeliverableShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, deliverableshape_ := range diagram.Deliverable_Shapes {
					res[deliverableshape_] = append(res[deliverableshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DeliverablesWhoseNodeIsExpanded":
			res := make(map[*Deliverable][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, deliverable_ := range diagram.DeliverablesWhoseNodeIsExpanded {
					res[deliverable_] = append(res[deliverable_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DeliverablesWhoseConceptsNodeIsExpanded":
			res := make(map[*Deliverable][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, deliverable_ := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
					res[deliverable_] = append(res[deliverable_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DeliverableComposition_Shapes":
			res := make(map[*DeliverableCompositionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, deliverablecompositionshape_ := range diagram.DeliverableComposition_Shapes {
					res[deliverablecompositionshape_] = append(res[deliverablecompositionshape_], diagram)
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
		case "NoteDeliverableShapes":
			res := make(map[*NoteDeliverableShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, notedeliverableshape_ := range diagram.NoteDeliverableShapes {
					res[notedeliverableshape_] = append(res[notedeliverableshape_], diagram)
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
		case "Requirement_Shapes":
			res := make(map[*RequirementShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, requirementshape_ := range diagram.Requirement_Shapes {
					res[requirementshape_] = append(res[requirementshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RequirementsWhoseNodeIsExpanded":
			res := make(map[*Requirement][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, requirement_ := range diagram.RequirementsWhoseNodeIsExpanded {
					res[requirement_] = append(res[requirement_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Concept_Shapes":
			res := make(map[*ConceptShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, conceptshape_ := range diagram.Concept_Shapes {
					res[conceptshape_] = append(res[conceptshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConceptsWhoseNodeIsExpanded":
			res := make(map[*Concept][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concept_ := range diagram.ConceptsWhoseNodeIsExpanded {
					res[concept_] = append(res[concept_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ConceptsWhoseDeliverablesNodeIsExpanded":
			res := make(map[*Concept][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, concept_ := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
					res[concept_] = append(res[concept_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DeliverableConceptShapes":
			res := make(map[*DeliverableConceptShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, deliverableconceptshape_ := range diagram.DeliverableConceptShapes {
					res[deliverableconceptshape_] = append(res[deliverableconceptshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Diagram_Shapes":
			res := make(map[*DiagramShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, diagramshape_ := range diagram.Diagram_Shapes {
					res[diagramshape_] = append(res[diagramshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramsWhoseNodeIsExpanded":
			res := make(map[*Diagram][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, diagram_ := range diagram.DiagramsWhoseNodeIsExpanded {
					res[diagram_] = append(res[diagram_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DiagramShape
	case DiagramShape:
		switch fieldname {
		// insertion point for per direct association field
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
		case "Deliverables":
			res := make(map[*Deliverable][]*Note)
			for note := range stage.Notes {
				for _, deliverable_ := range note.Deliverables {
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
	// reverse maps of direct associations of NoteDeliverableShape
	case NoteDeliverableShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*NoteDeliverableShape)
			for notedeliverableshape := range stage.NoteDeliverableShapes {
				for _, controlpointshape_ := range notedeliverableshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], notedeliverableshape)
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*NoteStakeholderShape)
			for notestakeholdershape := range stage.NoteStakeholderShapes {
				for _, controlpointshape_ := range notestakeholdershape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], notestakeholdershape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteTaskShape
	case NoteTaskShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*NoteTaskShape)
			for notetaskshape := range stage.NoteTaskShapes {
				for _, controlpointshape_ := range notetaskshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], notetaskshape)
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of RequirementShape
	case RequirementShape:
		switch fieldname {
		// insertion point for per direct association field
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
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*StakeholderCompositionShape)
			for stakeholdercompositionshape := range stage.StakeholderCompositionShapes {
				for _, controlpointshape_ := range stakeholdercompositionshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], stakeholdercompositionshape)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StakeholderConcernShape
	case StakeholderConcernShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlPointShapes":
			res := make(map[*ControlPointShape][]*StakeholderConcernShape)
			for stakeholderconcernshape := range stage.StakeholderConcernShapes {
				for _, controlpointshape_ := range stakeholderconcernshape.ControlPointShapes {
					res[controlpointshape_] = append(res[controlpointshape_], stakeholderconcernshape)
				}
			}
			return any(res).(map[*End][]*Start)
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

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *AnalysisNeed:
		res = any(new(AnalysisNeed)).(Type)
	case *Concept:
		res = any(new(Concept)).(Type)
	case *ConceptShape:
		res = any(new(ConceptShape)).(Type)
	case *Concern:
		res = any(new(Concern)).(Type)
	case *ConcernCompositionShape:
		res = any(new(ConcernCompositionShape)).(Type)
	case *ConcernInputShape:
		res = any(new(ConcernInputShape)).(Type)
	case *ConcernOutputShape:
		res = any(new(ConcernOutputShape)).(Type)
	case *ConcernShape:
		res = any(new(ConcernShape)).(Type)
	case *ControlPointShape:
		res = any(new(ControlPointShape)).(Type)
	case *Deliverable:
		res = any(new(Deliverable)).(Type)
	case *DeliverableCompositionShape:
		res = any(new(DeliverableCompositionShape)).(Type)
	case *DeliverableConceptShape:
		res = any(new(DeliverableConceptShape)).(Type)
	case *DeliverableShape:
		res = any(new(DeliverableShape)).(Type)
	case *Diagram:
		res = any(new(Diagram)).(Type)
	case *DiagramShape:
		res = any(new(DiagramShape)).(Type)
	case *Library:
		res = any(new(Library)).(Type)
	case *Note:
		res = any(new(Note)).(Type)
	case *NoteDeliverableShape:
		res = any(new(NoteDeliverableShape)).(Type)
	case *NoteShape:
		res = any(new(NoteShape)).(Type)
	case *NoteStakeholderShape:
		res = any(new(NoteStakeholderShape)).(Type)
	case *NoteTaskShape:
		res = any(new(NoteTaskShape)).(Type)
	case *Requirement:
		res = any(new(Requirement)).(Type)
	case *RequirementShape:
		res = any(new(RequirementShape)).(Type)
	case *Stakeholder:
		res = any(new(Stakeholder)).(Type)
	case *StakeholderCompositionShape:
		res = any(new(StakeholderCompositionShape)).(Type)
	case *StakeholderConcernShape:
		res = any(new(StakeholderConcernShape)).(Type)
	case *StakeholderShape:
		res = any(new(StakeholderShape)).(Type)
	case *SupportLevel:
		res = any(new(SupportLevel)).(Type)
	case *Tool:
		res = any(new(Tool)).(Type)
	}
	return res
}

func NewInstance[Type GongstructPtr]() (res Type) {
	return GongNewInstance[Type]()
}

func (stage *Stage) GongNewInstance[Type GongstructPtr]() (res Type) {
	res = GongNewInstance[Type]()
	var zero Type
	if res != zero {
		res.StageVoid(stage)
	}
	return res
}

func (stage *Stage) NewInstance[Type GongstructPtr]() (res Type) {
	return stage.GongNewInstance[Type]()
}

// GongGetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GongGetPointerToGongstructName[Type GongstructIF]() (res string) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AnalysisNeed:
		res = "AnalysisNeed"
	case *Concept:
		res = "Concept"
	case *ConceptShape:
		res = "ConceptShape"
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
	case *ControlPointShape:
		res = "ControlPointShape"
	case *Deliverable:
		res = "Deliverable"
	case *DeliverableCompositionShape:
		res = "DeliverableCompositionShape"
	case *DeliverableConceptShape:
		res = "DeliverableConceptShape"
	case *DeliverableShape:
		res = "DeliverableShape"
	case *Diagram:
		res = "Diagram"
	case *DiagramShape:
		res = "DiagramShape"
	case *Library:
		res = "Library"
	case *Note:
		res = "Note"
	case *NoteDeliverableShape:
		res = "NoteDeliverableShape"
	case *NoteShape:
		res = "NoteShape"
	case *NoteStakeholderShape:
		res = "NoteStakeholderShape"
	case *NoteTaskShape:
		res = "NoteTaskShape"
	case *Requirement:
		res = "Requirement"
	case *RequirementShape:
		res = "RequirementShape"
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

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	res = make([]GongReverseField, 0)

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
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConceptsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ConceptsWhoseDeliverablesNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootConcepts"
		res = append(res, rf)
		rf.GongstructName = "Requirement"
		rf.Fieldname = "Concepts"
		res = append(res, rf)
	case *ConceptShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Concept_Shapes"
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
	case *ControlPointShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ConcernCompositionShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "ConcernInputShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "ConcernOutputShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "DeliverableCompositionShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "DeliverableConceptShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "NoteDeliverableShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "NoteStakeholderShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "NoteTaskShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "StakeholderCompositionShape"
		rf.Fieldname = "ControlPointShapes"
		res = append(res, rf)
		rf.GongstructName = "StakeholderConcernShape"
		rf.Fieldname = "ControlPointShapes"
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
		rf.Fieldname = "SubDeliverables"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "DeliverablesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "DeliverablesWhoseConceptsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootDeliverables"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Deliverables"
		res = append(res, rf)
	case *DeliverableCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "DeliverableComposition_Shapes"
		res = append(res, rf)
	case *DeliverableConceptShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "DeliverableConceptShapes"
		res = append(res, rf)
	case *DeliverableShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Deliverable_Shapes"
		res = append(res, rf)
	case *Diagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "DiagramsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
	case *DiagramShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Diagram_Shapes"
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
	case *NoteDeliverableShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NoteDeliverableShapes"
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
	case *Requirement:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Concern"
		rf.Fieldname = "Requirements"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "RequirementsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootRequirements"
		res = append(res, rf)
	case *RequirementShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Requirement_Shapes"
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

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
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
			Name:                 "Tools",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Tool",
		},
	}
	return
}

func (conceptshape *ConceptShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Concept",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concept",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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
			Name:                 "Concern",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concern",
		},
		{
			Name:                 "Deliverable",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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

func (controlpointshape *ControlPointShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X_Relative",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Y_Relative",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsStartShapeTheClosestShape",
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SubDeliverables",
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

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
		},
	}
	return
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Concept",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Concept",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
		},
	}
	return
}

func (deliverableshape *DeliverableShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Deliverable_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DeliverableShape",
		},
		{
			Name:                 "DeliverablesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:                 "DeliverablesWhoseConceptsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Deliverable",
		},
		{
			Name:               "IsPBSNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "DeliverableComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DeliverableCompositionShape",
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
			Name:                 "NoteDeliverableShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteDeliverableShape",
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
		{
			Name:                 "Requirement_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "RequirementShape",
		},
		{
			Name:                 "RequirementsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Requirement",
		},
		{
			Name:                 "Concept_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ConceptShape",
		},
		{
			Name:                 "ConceptsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concept",
		},
		{
			Name:                 "ConceptsWhoseDeliverablesNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Concept",
		},
		{
			Name:                 "DeliverableConceptShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DeliverableConceptShape",
		},
		{
			Name:                 "Diagram_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramShape",
		},
		{
			Name:               "IsDiagramsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "DiagramsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
	}
	return
}

func (diagramshape *DiagramShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Diagram",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Diagram",
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
			Name:                 "Deliverables",
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

func (notedeliverableshape *NoteDeliverableShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Deliverable",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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

func (requirementshape *RequirementShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Requirement",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Requirement",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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
		{
			Name:                 "ControlPointShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlPointShape",
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

// GongGetFieldsFromPointer return the array of the fields
func GongGetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
}

func GetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	return GongGetFieldsFromPointer[Type]()
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

func (conceptshape *ConceptShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = conceptshape.Name
	case "Concept":
		res.GongFieldValueType = GongFieldValueTypePointer
		if conceptshape.Concept != nil {
			res.valueString = conceptshape.Concept.Name
			res.ids = conceptshape.Concept.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", conceptshape.IsExpanded)
		res.valueBool = conceptshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", conceptshape.X)
		res.valueFloat = conceptshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", conceptshape.Y)
		res.valueFloat = conceptshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", conceptshape.Width)
		res.valueFloat = conceptshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", conceptshape.Height)
		res.valueFloat = conceptshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", conceptshape.IsHidden)
		res.valueBool = conceptshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range concerncompositionshape.ControlPointShapes {
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
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range concerninputshape.ControlPointShapes {
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

func (concernoutputshape *ConcernOutputShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = concernoutputshape.Name
	case "Concern":
		res.GongFieldValueType = GongFieldValueTypePointer
		if concernoutputshape.Concern != nil {
			res.valueString = concernoutputshape.Concern.Name
			res.ids = concernoutputshape.Concern.GongGetUUID(stage)
		}
	case "Deliverable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if concernoutputshape.Deliverable != nil {
			res.valueString = concernoutputshape.Deliverable.Name
			res.ids = concernoutputshape.Deliverable.GongGetUUID(stage)
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
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range concernoutputshape.ControlPointShapes {
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

func (controlpointshape *ControlPointShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlpointshape.Name
	case "X_Relative":
		res.valueString = fmt.Sprintf("%f", controlpointshape.X_Relative)
		res.valueFloat = controlpointshape.X_Relative
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y_Relative":
		res.valueString = fmt.Sprintf("%f", controlpointshape.Y_Relative)
		res.valueFloat = controlpointshape.Y_Relative
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsStartShapeTheClosestShape":
		res.valueString = fmt.Sprintf("%t", controlpointshape.IsStartShapeTheClosestShape)
		res.valueBool = controlpointshape.IsStartShapeTheClosestShape
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
	case "Description":
		res.valueString = deliverable.Description
	case "SubDeliverables":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range deliverable.SubDeliverables {
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

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = deliverablecompositionshape.Name
	case "Deliverable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if deliverablecompositionshape.Deliverable != nil {
			res.valueString = deliverablecompositionshape.Deliverable.Name
			res.ids = deliverablecompositionshape.Deliverable.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", deliverablecompositionshape.StartRatio)
		res.valueFloat = deliverablecompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", deliverablecompositionshape.EndRatio)
		res.valueFloat = deliverablecompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := deliverablecompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := deliverablecompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", deliverablecompositionshape.CornerOffsetRatio)
		res.valueFloat = deliverablecompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", deliverablecompositionshape.IsHidden)
		res.valueBool = deliverablecompositionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range deliverablecompositionshape.ControlPointShapes {
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

func (deliverableconceptshape *DeliverableConceptShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = deliverableconceptshape.Name
	case "Deliverable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if deliverableconceptshape.Deliverable != nil {
			res.valueString = deliverableconceptshape.Deliverable.Name
			res.ids = deliverableconceptshape.Deliverable.GongGetUUID(stage)
		}
	case "Concept":
		res.GongFieldValueType = GongFieldValueTypePointer
		if deliverableconceptshape.Concept != nil {
			res.valueString = deliverableconceptshape.Concept.Name
			res.ids = deliverableconceptshape.Concept.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", deliverableconceptshape.StartRatio)
		res.valueFloat = deliverableconceptshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", deliverableconceptshape.EndRatio)
		res.valueFloat = deliverableconceptshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := deliverableconceptshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := deliverableconceptshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", deliverableconceptshape.CornerOffsetRatio)
		res.valueFloat = deliverableconceptshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", deliverableconceptshape.IsHidden)
		res.valueBool = deliverableconceptshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range deliverableconceptshape.ControlPointShapes {
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

func (deliverableshape *DeliverableShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = deliverableshape.Name
	case "Deliverable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if deliverableshape.Deliverable != nil {
			res.valueString = deliverableshape.Deliverable.Name
			res.ids = deliverableshape.Deliverable.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", deliverableshape.IsExpanded)
		res.valueBool = deliverableshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", deliverableshape.X)
		res.valueFloat = deliverableshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", deliverableshape.Y)
		res.valueFloat = deliverableshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", deliverableshape.Width)
		res.valueFloat = deliverableshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", deliverableshape.Height)
		res.valueFloat = deliverableshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", deliverableshape.IsHidden)
		res.valueBool = deliverableshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "Deliverable_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Deliverable_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DeliverablesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.DeliverablesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DeliverablesWhoseConceptsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.DeliverablesWhoseConceptsNodeIsExpanded {
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
	case "DeliverableComposition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.DeliverableComposition_Shapes {
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
	case "NoteDeliverableShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.NoteDeliverableShapes {
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
	case "Requirement_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Requirement_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RequirementsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.RequirementsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Concept_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Concept_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConceptsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConceptsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ConceptsWhoseDeliverablesNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ConceptsWhoseDeliverablesNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DeliverableConceptShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.DeliverableConceptShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Diagram_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Diagram_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDiagramsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsDiagramsNodeExpanded)
		res.valueBool = diagram.IsDiagramsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DiagramsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.DiagramsWhoseNodeIsExpanded {
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

func (diagramshape *DiagramShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramshape.Name
	case "Diagram":
		res.GongFieldValueType = GongFieldValueTypePointer
		if diagramshape.Diagram != nil {
			res.valueString = diagramshape.Diagram.Name
			res.ids = diagramshape.Diagram.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagramshape.IsExpanded)
		res.valueBool = diagramshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", diagramshape.X)
		res.valueFloat = diagramshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", diagramshape.Y)
		res.valueFloat = diagramshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramshape.Width)
		res.valueFloat = diagramshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramshape.Height)
		res.valueFloat = diagramshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", diagramshape.IsHidden)
		res.valueBool = diagramshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "Deliverables":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Deliverables {
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

func (notedeliverableshape *NoteDeliverableShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notedeliverableshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notedeliverableshape.Note != nil {
			res.valueString = notedeliverableshape.Note.Name
			res.ids = notedeliverableshape.Note.GongGetUUID(stage)
		}
	case "Deliverable":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notedeliverableshape.Deliverable != nil {
			res.valueString = notedeliverableshape.Deliverable.Name
			res.ids = notedeliverableshape.Deliverable.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", notedeliverableshape.StartRatio)
		res.valueFloat = notedeliverableshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", notedeliverableshape.EndRatio)
		res.valueFloat = notedeliverableshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := notedeliverableshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := notedeliverableshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", notedeliverableshape.CornerOffsetRatio)
		res.valueFloat = notedeliverableshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", notedeliverableshape.IsHidden)
		res.valueBool = notedeliverableshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notedeliverableshape.ControlPointShapes {
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
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notestakeholdershape.ControlPointShapes {
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
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range notetaskshape.ControlPointShapes {
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

func (requirementshape *RequirementShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = requirementshape.Name
	case "Requirement":
		res.GongFieldValueType = GongFieldValueTypePointer
		if requirementshape.Requirement != nil {
			res.valueString = requirementshape.Requirement.Name
			res.ids = requirementshape.Requirement.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", requirementshape.IsExpanded)
		res.valueBool = requirementshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", requirementshape.X)
		res.valueFloat = requirementshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", requirementshape.Y)
		res.valueFloat = requirementshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", requirementshape.Width)
		res.valueFloat = requirementshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", requirementshape.Height)
		res.valueFloat = requirementshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", requirementshape.IsHidden)
		res.valueBool = requirementshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stakeholdercompositionshape.ControlPointShapes {
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
	case "ControlPointShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range stakeholderconcernshape.ControlPointShapes {
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (analysisneed *AnalysisNeed) GongGetGongstructName() string {
	return "AnalysisNeed"
}

func (concept *Concept) GongGetGongstructName() string {
	return "Concept"
}

func (conceptshape *ConceptShape) GongGetGongstructName() string {
	return "ConceptShape"
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

func (controlpointshape *ControlPointShape) GongGetGongstructName() string {
	return "ControlPointShape"
}

func (deliverable *Deliverable) GongGetGongstructName() string {
	return "Deliverable"
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongGetGongstructName() string {
	return "DeliverableCompositionShape"
}

func (deliverableconceptshape *DeliverableConceptShape) GongGetGongstructName() string {
	return "DeliverableConceptShape"
}

func (deliverableshape *DeliverableShape) GongGetGongstructName() string {
	return "DeliverableShape"
}

func (diagram *Diagram) GongGetGongstructName() string {
	return "Diagram"
}

func (diagramshape *DiagramShape) GongGetGongstructName() string {
	return "DiagramShape"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (note *Note) GongGetGongstructName() string {
	return "Note"
}

func (notedeliverableshape *NoteDeliverableShape) GongGetGongstructName() string {
	return "NoteDeliverableShape"
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

func (requirement *Requirement) GongGetGongstructName() string {
	return "Requirement"
}

func (requirementshape *RequirementShape) GongGetGongstructName() string {
	return "RequirementShape"
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

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	__gong__rebuildMapString(stage.AnalysisNeeds, &stage.AnalysisNeeds_mapString)

	__gong__rebuildMapString(stage.Concepts, &stage.Concepts_mapString)

	__gong__rebuildMapString(stage.ConceptShapes, &stage.ConceptShapes_mapString)

	__gong__rebuildMapString(stage.Concerns, &stage.Concerns_mapString)

	__gong__rebuildMapString(stage.ConcernCompositionShapes, &stage.ConcernCompositionShapes_mapString)

	__gong__rebuildMapString(stage.ConcernInputShapes, &stage.ConcernInputShapes_mapString)

	__gong__rebuildMapString(stage.ConcernOutputShapes, &stage.ConcernOutputShapes_mapString)

	__gong__rebuildMapString(stage.ConcernShapes, &stage.ConcernShapes_mapString)

	__gong__rebuildMapString(stage.ControlPointShapes, &stage.ControlPointShapes_mapString)

	__gong__rebuildMapString(stage.Deliverables, &stage.Deliverables_mapString)

	__gong__rebuildMapString(stage.DeliverableCompositionShapes, &stage.DeliverableCompositionShapes_mapString)

	__gong__rebuildMapString(stage.DeliverableConceptShapes, &stage.DeliverableConceptShapes_mapString)

	__gong__rebuildMapString(stage.DeliverableShapes, &stage.DeliverableShapes_mapString)

	__gong__rebuildMapString(stage.Diagrams, &stage.Diagrams_mapString)

	__gong__rebuildMapString(stage.DiagramShapes, &stage.DiagramShapes_mapString)

	__gong__rebuildMapString(stage.Librarys, &stage.Librarys_mapString)

	__gong__rebuildMapString(stage.Notes, &stage.Notes_mapString)

	__gong__rebuildMapString(stage.NoteDeliverableShapes, &stage.NoteDeliverableShapes_mapString)

	__gong__rebuildMapString(stage.NoteShapes, &stage.NoteShapes_mapString)

	__gong__rebuildMapString(stage.NoteStakeholderShapes, &stage.NoteStakeholderShapes_mapString)

	__gong__rebuildMapString(stage.NoteTaskShapes, &stage.NoteTaskShapes_mapString)

	__gong__rebuildMapString(stage.Requirements, &stage.Requirements_mapString)

	__gong__rebuildMapString(stage.RequirementShapes, &stage.RequirementShapes_mapString)

	__gong__rebuildMapString(stage.Stakeholders, &stage.Stakeholders_mapString)

	__gong__rebuildMapString(stage.StakeholderCompositionShapes, &stage.StakeholderCompositionShapes_mapString)

	__gong__rebuildMapString(stage.StakeholderConcernShapes, &stage.StakeholderConcernShapes_mapString)

	__gong__rebuildMapString(stage.StakeholderShapes, &stage.StakeholderShapes_mapString)

	__gong__rebuildMapString(stage.SupportLevels, &stage.SupportLevels_mapString)

	__gong__rebuildMapString(stage.Tools, &stage.Tools_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
