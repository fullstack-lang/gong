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

	project_go "github.com/fullstack-lang/gong/dsm/project/go"
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
	Diagrams                map[*Diagram]struct{}
	Diagrams_instance       map[*Diagram]*Diagram
	Diagrams_mapString      map[string]*Diagram
	DiagramOrder            uint
	Diagram_stagedOrder     map[*Diagram]uint
	Diagram_orderStaged     map[uint]*Diagram
	Diagrams_reference      map[*Diagram]*Diagram
	Diagrams_referenceOrder map[*Diagram]uint

	// insertion point for slice of pointers maps
	Diagram_Product_Shapes_reverseMap map[*ProductShape]*Diagram

	Diagram_ProductsWhoseNodeIsExpanded_reverseMap map[*Product]*Diagram

	Diagram_ProductComposition_Shapes_reverseMap map[*ProductCompositionShape]*Diagram

	Diagram_Task_Shapes_reverseMap map[*TaskShape]*Diagram

	Diagram_TasksWhoseNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TasksWhoseInputNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TasksWhoseOutputNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TaskGroupShapes_reverseMap map[*TaskGroupShape]*Diagram

	Diagram_TaskGroupsWhoseNodeIsExpanded_reverseMap map[*TaskGroup]*Diagram

	Diagram_TaskComposition_Shapes_reverseMap map[*TaskCompositionShape]*Diagram

	Diagram_TaskInputShapes_reverseMap map[*TaskInputShape]*Diagram

	Diagram_TaskOutputShapes_reverseMap map[*TaskOutputShape]*Diagram

	Diagram_Note_Shapes_reverseMap map[*NoteShape]*Diagram

	Diagram_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*Diagram

	Diagram_NoteProductShapes_reverseMap map[*NoteProductShape]*Diagram

	Diagram_NoteTaskShapes_reverseMap map[*NoteTaskShape]*Diagram

	Diagram_NoteResourceShapes_reverseMap map[*NoteResourceShape]*Diagram

	Diagram_Resource_Shapes_reverseMap map[*ResourceShape]*Diagram

	Diagram_ResourcesWhoseNodeIsExpanded_reverseMap map[*Resource]*Diagram

	Diagram_ResourceComposition_Shapes_reverseMap map[*ResourceCompositionShape]*Diagram

	Diagram_ResourceTaskShapes_reverseMap map[*ResourceTaskShape]*Diagram

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
	Library_SubLibraries_reverseMap map[*Library]*Library

	Library_RootProducts_reverseMap map[*Product]*Library

	Library_RootTasks_reverseMap map[*Task]*Library

	Library_RootTaskGroups_reverseMap map[*TaskGroup]*Library

	Library_RootResources_reverseMap map[*Resource]*Library

	Library_Notes_reverseMap map[*Note]*Library

	Library_Diagrams_reverseMap map[*Diagram]*Library

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
	Note_Products_reverseMap map[*Product]*Note

	Note_Tasks_reverseMap map[*Task]*Note

	Note_Resources_reverseMap map[*Resource]*Note

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

	NoteResourceShapes                map[*NoteResourceShape]struct{}
	NoteResourceShapes_instance       map[*NoteResourceShape]*NoteResourceShape
	NoteResourceShapes_mapString      map[string]*NoteResourceShape
	NoteResourceShapeOrder            uint
	NoteResourceShape_stagedOrder     map[*NoteResourceShape]uint
	NoteResourceShape_orderStaged     map[uint]*NoteResourceShape
	NoteResourceShapes_reference      map[*NoteResourceShape]*NoteResourceShape
	NoteResourceShapes_referenceOrder map[*NoteResourceShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteResourceShapeCreateCallback OnAfterCreateInterface[NoteResourceShape]
	OnAfterNoteResourceShapeUpdateCallback OnAfterUpdateInterface[NoteResourceShape]
	OnAfterNoteResourceShapeDeleteCallback OnAfterDeleteInterface[NoteResourceShape]
	OnAfterNoteResourceShapeReadCallback   OnAfterReadInterface[NoteResourceShape]

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

	Products                map[*Product]struct{}
	Products_instance       map[*Product]*Product
	Products_mapString      map[string]*Product
	ProductOrder            uint
	Product_stagedOrder     map[*Product]uint
	Product_orderStaged     map[uint]*Product
	Products_reference      map[*Product]*Product
	Products_referenceOrder map[*Product]uint

	// insertion point for slice of pointers maps
	Product_SubProducts_reverseMap map[*Product]*Product

	OnAfterProductCreateCallback OnAfterCreateInterface[Product]
	OnAfterProductUpdateCallback OnAfterUpdateInterface[Product]
	OnAfterProductDeleteCallback OnAfterDeleteInterface[Product]
	OnAfterProductReadCallback   OnAfterReadInterface[Product]

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

	Resources                map[*Resource]struct{}
	Resources_instance       map[*Resource]*Resource
	Resources_mapString      map[string]*Resource
	ResourceOrder            uint
	Resource_stagedOrder     map[*Resource]uint
	Resource_orderStaged     map[uint]*Resource
	Resources_reference      map[*Resource]*Resource
	Resources_referenceOrder map[*Resource]uint

	// insertion point for slice of pointers maps
	Resource_Tasks_reverseMap map[*Task]*Resource

	Resource_SubResources_reverseMap map[*Resource]*Resource

	OnAfterResourceCreateCallback OnAfterCreateInterface[Resource]
	OnAfterResourceUpdateCallback OnAfterUpdateInterface[Resource]
	OnAfterResourceDeleteCallback OnAfterDeleteInterface[Resource]
	OnAfterResourceReadCallback   OnAfterReadInterface[Resource]

	ResourceCompositionShapes                map[*ResourceCompositionShape]struct{}
	ResourceCompositionShapes_instance       map[*ResourceCompositionShape]*ResourceCompositionShape
	ResourceCompositionShapes_mapString      map[string]*ResourceCompositionShape
	ResourceCompositionShapeOrder            uint
	ResourceCompositionShape_stagedOrder     map[*ResourceCompositionShape]uint
	ResourceCompositionShape_orderStaged     map[uint]*ResourceCompositionShape
	ResourceCompositionShapes_reference      map[*ResourceCompositionShape]*ResourceCompositionShape
	ResourceCompositionShapes_referenceOrder map[*ResourceCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterResourceCompositionShapeCreateCallback OnAfterCreateInterface[ResourceCompositionShape]
	OnAfterResourceCompositionShapeUpdateCallback OnAfterUpdateInterface[ResourceCompositionShape]
	OnAfterResourceCompositionShapeDeleteCallback OnAfterDeleteInterface[ResourceCompositionShape]
	OnAfterResourceCompositionShapeReadCallback   OnAfterReadInterface[ResourceCompositionShape]

	ResourceShapes                map[*ResourceShape]struct{}
	ResourceShapes_instance       map[*ResourceShape]*ResourceShape
	ResourceShapes_mapString      map[string]*ResourceShape
	ResourceShapeOrder            uint
	ResourceShape_stagedOrder     map[*ResourceShape]uint
	ResourceShape_orderStaged     map[uint]*ResourceShape
	ResourceShapes_reference      map[*ResourceShape]*ResourceShape
	ResourceShapes_referenceOrder map[*ResourceShape]uint

	// insertion point for slice of pointers maps
	OnAfterResourceShapeCreateCallback OnAfterCreateInterface[ResourceShape]
	OnAfterResourceShapeUpdateCallback OnAfterUpdateInterface[ResourceShape]
	OnAfterResourceShapeDeleteCallback OnAfterDeleteInterface[ResourceShape]
	OnAfterResourceShapeReadCallback   OnAfterReadInterface[ResourceShape]

	ResourceTaskShapes                map[*ResourceTaskShape]struct{}
	ResourceTaskShapes_instance       map[*ResourceTaskShape]*ResourceTaskShape
	ResourceTaskShapes_mapString      map[string]*ResourceTaskShape
	ResourceTaskShapeOrder            uint
	ResourceTaskShape_stagedOrder     map[*ResourceTaskShape]uint
	ResourceTaskShape_orderStaged     map[uint]*ResourceTaskShape
	ResourceTaskShapes_reference      map[*ResourceTaskShape]*ResourceTaskShape
	ResourceTaskShapes_referenceOrder map[*ResourceTaskShape]uint

	// insertion point for slice of pointers maps
	OnAfterResourceTaskShapeCreateCallback OnAfterCreateInterface[ResourceTaskShape]
	OnAfterResourceTaskShapeUpdateCallback OnAfterUpdateInterface[ResourceTaskShape]
	OnAfterResourceTaskShapeDeleteCallback OnAfterDeleteInterface[ResourceTaskShape]
	OnAfterResourceTaskShapeReadCallback   OnAfterReadInterface[ResourceTaskShape]

	Tasks                map[*Task]struct{}
	Tasks_instance       map[*Task]*Task
	Tasks_mapString      map[string]*Task
	TaskOrder            uint
	Task_stagedOrder     map[*Task]uint
	Task_orderStaged     map[uint]*Task
	Tasks_reference      map[*Task]*Task
	Tasks_referenceOrder map[*Task]uint

	// insertion point for slice of pointers maps
	Task_Predecessors_reverseMap map[*Task]*Task

	Task_SubTasks_reverseMap map[*Task]*Task

	Task_Inputs_reverseMap map[*Product]*Task

	Task_Outputs_reverseMap map[*Product]*Task

	Task_TaskGroupsToDisplay_reverseMap map[*TaskGroup]*Task

	OnAfterTaskCreateCallback OnAfterCreateInterface[Task]
	OnAfterTaskUpdateCallback OnAfterUpdateInterface[Task]
	OnAfterTaskDeleteCallback OnAfterDeleteInterface[Task]
	OnAfterTaskReadCallback   OnAfterReadInterface[Task]

	TaskCompositionShapes                map[*TaskCompositionShape]struct{}
	TaskCompositionShapes_instance       map[*TaskCompositionShape]*TaskCompositionShape
	TaskCompositionShapes_mapString      map[string]*TaskCompositionShape
	TaskCompositionShapeOrder            uint
	TaskCompositionShape_stagedOrder     map[*TaskCompositionShape]uint
	TaskCompositionShape_orderStaged     map[uint]*TaskCompositionShape
	TaskCompositionShapes_reference      map[*TaskCompositionShape]*TaskCompositionShape
	TaskCompositionShapes_referenceOrder map[*TaskCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskCompositionShapeCreateCallback OnAfterCreateInterface[TaskCompositionShape]
	OnAfterTaskCompositionShapeUpdateCallback OnAfterUpdateInterface[TaskCompositionShape]
	OnAfterTaskCompositionShapeDeleteCallback OnAfterDeleteInterface[TaskCompositionShape]
	OnAfterTaskCompositionShapeReadCallback   OnAfterReadInterface[TaskCompositionShape]

	TaskGroups                map[*TaskGroup]struct{}
	TaskGroups_instance       map[*TaskGroup]*TaskGroup
	TaskGroups_mapString      map[string]*TaskGroup
	TaskGroupOrder            uint
	TaskGroup_stagedOrder     map[*TaskGroup]uint
	TaskGroup_orderStaged     map[uint]*TaskGroup
	TaskGroups_reference      map[*TaskGroup]*TaskGroup
	TaskGroups_referenceOrder map[*TaskGroup]uint

	// insertion point for slice of pointers maps
	TaskGroup_Tasks_reverseMap map[*Task]*TaskGroup

	OnAfterTaskGroupCreateCallback OnAfterCreateInterface[TaskGroup]
	OnAfterTaskGroupUpdateCallback OnAfterUpdateInterface[TaskGroup]
	OnAfterTaskGroupDeleteCallback OnAfterDeleteInterface[TaskGroup]
	OnAfterTaskGroupReadCallback   OnAfterReadInterface[TaskGroup]

	TaskGroupShapes                map[*TaskGroupShape]struct{}
	TaskGroupShapes_instance       map[*TaskGroupShape]*TaskGroupShape
	TaskGroupShapes_mapString      map[string]*TaskGroupShape
	TaskGroupShapeOrder            uint
	TaskGroupShape_stagedOrder     map[*TaskGroupShape]uint
	TaskGroupShape_orderStaged     map[uint]*TaskGroupShape
	TaskGroupShapes_reference      map[*TaskGroupShape]*TaskGroupShape
	TaskGroupShapes_referenceOrder map[*TaskGroupShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskGroupShapeCreateCallback OnAfterCreateInterface[TaskGroupShape]
	OnAfterTaskGroupShapeUpdateCallback OnAfterUpdateInterface[TaskGroupShape]
	OnAfterTaskGroupShapeDeleteCallback OnAfterDeleteInterface[TaskGroupShape]
	OnAfterTaskGroupShapeReadCallback   OnAfterReadInterface[TaskGroupShape]

	TaskInputShapes                map[*TaskInputShape]struct{}
	TaskInputShapes_instance       map[*TaskInputShape]*TaskInputShape
	TaskInputShapes_mapString      map[string]*TaskInputShape
	TaskInputShapeOrder            uint
	TaskInputShape_stagedOrder     map[*TaskInputShape]uint
	TaskInputShape_orderStaged     map[uint]*TaskInputShape
	TaskInputShapes_reference      map[*TaskInputShape]*TaskInputShape
	TaskInputShapes_referenceOrder map[*TaskInputShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskInputShapeCreateCallback OnAfterCreateInterface[TaskInputShape]
	OnAfterTaskInputShapeUpdateCallback OnAfterUpdateInterface[TaskInputShape]
	OnAfterTaskInputShapeDeleteCallback OnAfterDeleteInterface[TaskInputShape]
	OnAfterTaskInputShapeReadCallback   OnAfterReadInterface[TaskInputShape]

	TaskOutputShapes                map[*TaskOutputShape]struct{}
	TaskOutputShapes_instance       map[*TaskOutputShape]*TaskOutputShape
	TaskOutputShapes_mapString      map[string]*TaskOutputShape
	TaskOutputShapeOrder            uint
	TaskOutputShape_stagedOrder     map[*TaskOutputShape]uint
	TaskOutputShape_orderStaged     map[uint]*TaskOutputShape
	TaskOutputShapes_reference      map[*TaskOutputShape]*TaskOutputShape
	TaskOutputShapes_referenceOrder map[*TaskOutputShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskOutputShapeCreateCallback OnAfterCreateInterface[TaskOutputShape]
	OnAfterTaskOutputShapeUpdateCallback OnAfterUpdateInterface[TaskOutputShape]
	OnAfterTaskOutputShapeDeleteCallback OnAfterDeleteInterface[TaskOutputShape]
	OnAfterTaskOutputShapeReadCallback   OnAfterReadInterface[TaskOutputShape]

	TaskShapes                map[*TaskShape]struct{}
	TaskShapes_instance       map[*TaskShape]*TaskShape
	TaskShapes_mapString      map[string]*TaskShape
	TaskShapeOrder            uint
	TaskShape_stagedOrder     map[*TaskShape]uint
	TaskShape_orderStaged     map[uint]*TaskShape
	TaskShapes_reference      map[*TaskShape]*TaskShape
	TaskShapes_referenceOrder map[*TaskShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskShapeCreateCallback OnAfterCreateInterface[TaskShape]
	OnAfterTaskShapeUpdateCallback OnAfterUpdateInterface[TaskShape]
	OnAfterTaskShapeDeleteCallback OnAfterDeleteInterface[TaskShape]
	OnAfterTaskShapeReadCallback   OnAfterReadInterface[TaskShape]

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

	stage.NoteResourceShapes_reference = make(map[*NoteResourceShape]*NoteResourceShape)
	stage.NoteResourceShapes_instance = make(map[*NoteResourceShape]*NoteResourceShape)
	stage.NoteResourceShapes_referenceOrder = make(map[*NoteResourceShape]uint)

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)

	stage.NoteTaskShapes_reference = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_instance = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint)

	stage.Products_reference = make(map[*Product]*Product)
	stage.Products_instance = make(map[*Product]*Product)
	stage.Products_referenceOrder = make(map[*Product]uint)

	stage.ProductCompositionShapes_reference = make(map[*ProductCompositionShape]*ProductCompositionShape)
	stage.ProductCompositionShapes_instance = make(map[*ProductCompositionShape]*ProductCompositionShape)
	stage.ProductCompositionShapes_referenceOrder = make(map[*ProductCompositionShape]uint)

	stage.ProductShapes_reference = make(map[*ProductShape]*ProductShape)
	stage.ProductShapes_instance = make(map[*ProductShape]*ProductShape)
	stage.ProductShapes_referenceOrder = make(map[*ProductShape]uint)

	stage.Resources_reference = make(map[*Resource]*Resource)
	stage.Resources_instance = make(map[*Resource]*Resource)
	stage.Resources_referenceOrder = make(map[*Resource]uint)

	stage.ResourceCompositionShapes_reference = make(map[*ResourceCompositionShape]*ResourceCompositionShape)
	stage.ResourceCompositionShapes_instance = make(map[*ResourceCompositionShape]*ResourceCompositionShape)
	stage.ResourceCompositionShapes_referenceOrder = make(map[*ResourceCompositionShape]uint)

	stage.ResourceShapes_reference = make(map[*ResourceShape]*ResourceShape)
	stage.ResourceShapes_instance = make(map[*ResourceShape]*ResourceShape)
	stage.ResourceShapes_referenceOrder = make(map[*ResourceShape]uint)

	stage.ResourceTaskShapes_reference = make(map[*ResourceTaskShape]*ResourceTaskShape)
	stage.ResourceTaskShapes_instance = make(map[*ResourceTaskShape]*ResourceTaskShape)
	stage.ResourceTaskShapes_referenceOrder = make(map[*ResourceTaskShape]uint)

	stage.Tasks_reference = make(map[*Task]*Task)
	stage.Tasks_instance = make(map[*Task]*Task)
	stage.Tasks_referenceOrder = make(map[*Task]uint)

	stage.TaskCompositionShapes_reference = make(map[*TaskCompositionShape]*TaskCompositionShape)
	stage.TaskCompositionShapes_instance = make(map[*TaskCompositionShape]*TaskCompositionShape)
	stage.TaskCompositionShapes_referenceOrder = make(map[*TaskCompositionShape]uint)

	stage.TaskGroups_reference = make(map[*TaskGroup]*TaskGroup)
	stage.TaskGroups_instance = make(map[*TaskGroup]*TaskGroup)
	stage.TaskGroups_referenceOrder = make(map[*TaskGroup]uint)

	stage.TaskGroupShapes_reference = make(map[*TaskGroupShape]*TaskGroupShape)
	stage.TaskGroupShapes_instance = make(map[*TaskGroupShape]*TaskGroupShape)
	stage.TaskGroupShapes_referenceOrder = make(map[*TaskGroupShape]uint)

	stage.TaskInputShapes_reference = make(map[*TaskInputShape]*TaskInputShape)
	stage.TaskInputShapes_instance = make(map[*TaskInputShape]*TaskInputShape)
	stage.TaskInputShapes_referenceOrder = make(map[*TaskInputShape]uint)

	stage.TaskOutputShapes_reference = make(map[*TaskOutputShape]*TaskOutputShape)
	stage.TaskOutputShapes_instance = make(map[*TaskOutputShape]*TaskOutputShape)
	stage.TaskOutputShapes_referenceOrder = make(map[*TaskOutputShape]uint)

	stage.TaskShapes_reference = make(map[*TaskShape]*TaskShape)
	stage.TaskShapes_instance = make(map[*TaskShape]*TaskShape)
	stage.TaskShapes_referenceOrder = make(map[*TaskShape]uint)

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

	var maxNoteResourceShapeOrder uint
	var foundNoteResourceShape bool
	for _, order := range stage.NoteResourceShape_stagedOrder {
		if !foundNoteResourceShape || order > maxNoteResourceShapeOrder {
			maxNoteResourceShapeOrder = order
			foundNoteResourceShape = true
		}
	}
	if foundNoteResourceShape {
		stage.NoteResourceShapeOrder = maxNoteResourceShapeOrder + 1
	} else {
		stage.NoteResourceShapeOrder = 0
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

	var maxProductOrder uint
	var foundProduct bool
	for _, order := range stage.Product_stagedOrder {
		if !foundProduct || order > maxProductOrder {
			maxProductOrder = order
			foundProduct = true
		}
	}
	if foundProduct {
		stage.ProductOrder = maxProductOrder + 1
	} else {
		stage.ProductOrder = 0
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

	var maxResourceOrder uint
	var foundResource bool
	for _, order := range stage.Resource_stagedOrder {
		if !foundResource || order > maxResourceOrder {
			maxResourceOrder = order
			foundResource = true
		}
	}
	if foundResource {
		stage.ResourceOrder = maxResourceOrder + 1
	} else {
		stage.ResourceOrder = 0
	}

	var maxResourceCompositionShapeOrder uint
	var foundResourceCompositionShape bool
	for _, order := range stage.ResourceCompositionShape_stagedOrder {
		if !foundResourceCompositionShape || order > maxResourceCompositionShapeOrder {
			maxResourceCompositionShapeOrder = order
			foundResourceCompositionShape = true
		}
	}
	if foundResourceCompositionShape {
		stage.ResourceCompositionShapeOrder = maxResourceCompositionShapeOrder + 1
	} else {
		stage.ResourceCompositionShapeOrder = 0
	}

	var maxResourceShapeOrder uint
	var foundResourceShape bool
	for _, order := range stage.ResourceShape_stagedOrder {
		if !foundResourceShape || order > maxResourceShapeOrder {
			maxResourceShapeOrder = order
			foundResourceShape = true
		}
	}
	if foundResourceShape {
		stage.ResourceShapeOrder = maxResourceShapeOrder + 1
	} else {
		stage.ResourceShapeOrder = 0
	}

	var maxResourceTaskShapeOrder uint
	var foundResourceTaskShape bool
	for _, order := range stage.ResourceTaskShape_stagedOrder {
		if !foundResourceTaskShape || order > maxResourceTaskShapeOrder {
			maxResourceTaskShapeOrder = order
			foundResourceTaskShape = true
		}
	}
	if foundResourceTaskShape {
		stage.ResourceTaskShapeOrder = maxResourceTaskShapeOrder + 1
	} else {
		stage.ResourceTaskShapeOrder = 0
	}

	var maxTaskOrder uint
	var foundTask bool
	for _, order := range stage.Task_stagedOrder {
		if !foundTask || order > maxTaskOrder {
			maxTaskOrder = order
			foundTask = true
		}
	}
	if foundTask {
		stage.TaskOrder = maxTaskOrder + 1
	} else {
		stage.TaskOrder = 0
	}

	var maxTaskCompositionShapeOrder uint
	var foundTaskCompositionShape bool
	for _, order := range stage.TaskCompositionShape_stagedOrder {
		if !foundTaskCompositionShape || order > maxTaskCompositionShapeOrder {
			maxTaskCompositionShapeOrder = order
			foundTaskCompositionShape = true
		}
	}
	if foundTaskCompositionShape {
		stage.TaskCompositionShapeOrder = maxTaskCompositionShapeOrder + 1
	} else {
		stage.TaskCompositionShapeOrder = 0
	}

	var maxTaskGroupOrder uint
	var foundTaskGroup bool
	for _, order := range stage.TaskGroup_stagedOrder {
		if !foundTaskGroup || order > maxTaskGroupOrder {
			maxTaskGroupOrder = order
			foundTaskGroup = true
		}
	}
	if foundTaskGroup {
		stage.TaskGroupOrder = maxTaskGroupOrder + 1
	} else {
		stage.TaskGroupOrder = 0
	}

	var maxTaskGroupShapeOrder uint
	var foundTaskGroupShape bool
	for _, order := range stage.TaskGroupShape_stagedOrder {
		if !foundTaskGroupShape || order > maxTaskGroupShapeOrder {
			maxTaskGroupShapeOrder = order
			foundTaskGroupShape = true
		}
	}
	if foundTaskGroupShape {
		stage.TaskGroupShapeOrder = maxTaskGroupShapeOrder + 1
	} else {
		stage.TaskGroupShapeOrder = 0
	}

	var maxTaskInputShapeOrder uint
	var foundTaskInputShape bool
	for _, order := range stage.TaskInputShape_stagedOrder {
		if !foundTaskInputShape || order > maxTaskInputShapeOrder {
			maxTaskInputShapeOrder = order
			foundTaskInputShape = true
		}
	}
	if foundTaskInputShape {
		stage.TaskInputShapeOrder = maxTaskInputShapeOrder + 1
	} else {
		stage.TaskInputShapeOrder = 0
	}

	var maxTaskOutputShapeOrder uint
	var foundTaskOutputShape bool
	for _, order := range stage.TaskOutputShape_stagedOrder {
		if !foundTaskOutputShape || order > maxTaskOutputShapeOrder {
			maxTaskOutputShapeOrder = order
			foundTaskOutputShape = true
		}
	}
	if foundTaskOutputShape {
		stage.TaskOutputShapeOrder = maxTaskOutputShapeOrder + 1
	} else {
		stage.TaskOutputShapeOrder = 0
	}

	var maxTaskShapeOrder uint
	var foundTaskShape bool
	for _, order := range stage.TaskShape_stagedOrder {
		if !foundTaskShape || order > maxTaskShapeOrder {
			maxTaskShapeOrder = order
			foundTaskShape = true
		}
	}
	if foundTaskShape {
		stage.TaskShapeOrder = maxTaskShapeOrder + 1
	} else {
		stage.TaskShapeOrder = 0
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
	case *NoteResourceShape:
		tmp := GetStructInstancesByOrder(stage.NoteResourceShapes, stage.NoteResourceShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteResourceShape implements.
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
	case *Product:
		tmp := GetStructInstancesByOrder(stage.Products, stage.Product_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Product implements.
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
	case *Resource:
		tmp := GetStructInstancesByOrder(stage.Resources, stage.Resource_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Resource implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ResourceCompositionShape:
		tmp := GetStructInstancesByOrder(stage.ResourceCompositionShapes, stage.ResourceCompositionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ResourceCompositionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ResourceShape:
		tmp := GetStructInstancesByOrder(stage.ResourceShapes, stage.ResourceShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ResourceShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ResourceTaskShape:
		tmp := GetStructInstancesByOrder(stage.ResourceTaskShapes, stage.ResourceTaskShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ResourceTaskShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Task:
		tmp := GetStructInstancesByOrder(stage.Tasks, stage.Task_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Task implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskCompositionShape:
		tmp := GetStructInstancesByOrder(stage.TaskCompositionShapes, stage.TaskCompositionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskCompositionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskGroup:
		tmp := GetStructInstancesByOrder(stage.TaskGroups, stage.TaskGroup_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskGroup implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskGroupShape:
		tmp := GetStructInstancesByOrder(stage.TaskGroupShapes, stage.TaskGroupShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskGroupShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskInputShape:
		tmp := GetStructInstancesByOrder(stage.TaskInputShapes, stage.TaskInputShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskInputShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskOutputShape:
		tmp := GetStructInstancesByOrder(stage.TaskOutputShapes, stage.TaskOutputShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskOutputShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *TaskShape:
		tmp := GetStructInstancesByOrder(stage.TaskShapes, stage.TaskShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *TaskShape implements.
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
	case "Diagram":
		res = GetNamedStructInstances(stage.Diagrams, stage.Diagram_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Note":
		res = GetNamedStructInstances(stage.Notes, stage.Note_stagedOrder)
	case "NoteProductShape":
		res = GetNamedStructInstances(stage.NoteProductShapes, stage.NoteProductShape_stagedOrder)
	case "NoteResourceShape":
		res = GetNamedStructInstances(stage.NoteResourceShapes, stage.NoteResourceShape_stagedOrder)
	case "NoteShape":
		res = GetNamedStructInstances(stage.NoteShapes, stage.NoteShape_stagedOrder)
	case "NoteTaskShape":
		res = GetNamedStructInstances(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder)
	case "Product":
		res = GetNamedStructInstances(stage.Products, stage.Product_stagedOrder)
	case "ProductCompositionShape":
		res = GetNamedStructInstances(stage.ProductCompositionShapes, stage.ProductCompositionShape_stagedOrder)
	case "ProductShape":
		res = GetNamedStructInstances(stage.ProductShapes, stage.ProductShape_stagedOrder)
	case "Resource":
		res = GetNamedStructInstances(stage.Resources, stage.Resource_stagedOrder)
	case "ResourceCompositionShape":
		res = GetNamedStructInstances(stage.ResourceCompositionShapes, stage.ResourceCompositionShape_stagedOrder)
	case "ResourceShape":
		res = GetNamedStructInstances(stage.ResourceShapes, stage.ResourceShape_stagedOrder)
	case "ResourceTaskShape":
		res = GetNamedStructInstances(stage.ResourceTaskShapes, stage.ResourceTaskShape_stagedOrder)
	case "Task":
		res = GetNamedStructInstances(stage.Tasks, stage.Task_stagedOrder)
	case "TaskCompositionShape":
		res = GetNamedStructInstances(stage.TaskCompositionShapes, stage.TaskCompositionShape_stagedOrder)
	case "TaskGroup":
		res = GetNamedStructInstances(stage.TaskGroups, stage.TaskGroup_stagedOrder)
	case "TaskGroupShape":
		res = GetNamedStructInstances(stage.TaskGroupShapes, stage.TaskGroupShape_stagedOrder)
	case "TaskInputShape":
		res = GetNamedStructInstances(stage.TaskInputShapes, stage.TaskInputShape_stagedOrder)
	case "TaskOutputShape":
		res = GetNamedStructInstances(stage.TaskOutputShapes, stage.TaskOutputShape_stagedOrder)
	case "TaskShape":
		res = GetNamedStructInstances(stage.TaskShapes, stage.TaskShape_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/project/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return project_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return project_go.GoDiagramsDir
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
	CommitDiagram(diagram *Diagram)
	CheckoutDiagram(diagram *Diagram)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitNote(note *Note)
	CheckoutNote(note *Note)
	CommitNoteProductShape(noteproductshape *NoteProductShape)
	CheckoutNoteProductShape(noteproductshape *NoteProductShape)
	CommitNoteResourceShape(noteresourceshape *NoteResourceShape)
	CheckoutNoteResourceShape(noteresourceshape *NoteResourceShape)
	CommitNoteShape(noteshape *NoteShape)
	CheckoutNoteShape(noteshape *NoteShape)
	CommitNoteTaskShape(notetaskshape *NoteTaskShape)
	CheckoutNoteTaskShape(notetaskshape *NoteTaskShape)
	CommitProduct(product *Product)
	CheckoutProduct(product *Product)
	CommitProductCompositionShape(productcompositionshape *ProductCompositionShape)
	CheckoutProductCompositionShape(productcompositionshape *ProductCompositionShape)
	CommitProductShape(productshape *ProductShape)
	CheckoutProductShape(productshape *ProductShape)
	CommitResource(resource *Resource)
	CheckoutResource(resource *Resource)
	CommitResourceCompositionShape(resourcecompositionshape *ResourceCompositionShape)
	CheckoutResourceCompositionShape(resourcecompositionshape *ResourceCompositionShape)
	CommitResourceShape(resourceshape *ResourceShape)
	CheckoutResourceShape(resourceshape *ResourceShape)
	CommitResourceTaskShape(resourcetaskshape *ResourceTaskShape)
	CheckoutResourceTaskShape(resourcetaskshape *ResourceTaskShape)
	CommitTask(task *Task)
	CheckoutTask(task *Task)
	CommitTaskCompositionShape(taskcompositionshape *TaskCompositionShape)
	CheckoutTaskCompositionShape(taskcompositionshape *TaskCompositionShape)
	CommitTaskGroup(taskgroup *TaskGroup)
	CheckoutTaskGroup(taskgroup *TaskGroup)
	CommitTaskGroupShape(taskgroupshape *TaskGroupShape)
	CheckoutTaskGroupShape(taskgroupshape *TaskGroupShape)
	CommitTaskInputShape(taskinputshape *TaskInputShape)
	CheckoutTaskInputShape(taskinputshape *TaskInputShape)
	CommitTaskOutputShape(taskoutputshape *TaskOutputShape)
	CheckoutTaskOutputShape(taskoutputshape *TaskOutputShape)
	CommitTaskShape(taskshape *TaskShape)
	CheckoutTaskShape(taskshape *TaskShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		Diagrams:           make(map[*Diagram]struct{}),
		Diagrams_mapString: make(map[string]*Diagram),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		NoteProductShapes:           make(map[*NoteProductShape]struct{}),
		NoteProductShapes_mapString: make(map[string]*NoteProductShape),

		NoteResourceShapes:           make(map[*NoteResourceShape]struct{}),
		NoteResourceShapes_mapString: make(map[string]*NoteResourceShape),

		NoteShapes:           make(map[*NoteShape]struct{}),
		NoteShapes_mapString: make(map[string]*NoteShape),

		NoteTaskShapes:           make(map[*NoteTaskShape]struct{}),
		NoteTaskShapes_mapString: make(map[string]*NoteTaskShape),

		Products:           make(map[*Product]struct{}),
		Products_mapString: make(map[string]*Product),

		ProductCompositionShapes:           make(map[*ProductCompositionShape]struct{}),
		ProductCompositionShapes_mapString: make(map[string]*ProductCompositionShape),

		ProductShapes:           make(map[*ProductShape]struct{}),
		ProductShapes_mapString: make(map[string]*ProductShape),

		Resources:           make(map[*Resource]struct{}),
		Resources_mapString: make(map[string]*Resource),

		ResourceCompositionShapes:           make(map[*ResourceCompositionShape]struct{}),
		ResourceCompositionShapes_mapString: make(map[string]*ResourceCompositionShape),

		ResourceShapes:           make(map[*ResourceShape]struct{}),
		ResourceShapes_mapString: make(map[string]*ResourceShape),

		ResourceTaskShapes:           make(map[*ResourceTaskShape]struct{}),
		ResourceTaskShapes_mapString: make(map[string]*ResourceTaskShape),

		Tasks:           make(map[*Task]struct{}),
		Tasks_mapString: make(map[string]*Task),

		TaskCompositionShapes:           make(map[*TaskCompositionShape]struct{}),
		TaskCompositionShapes_mapString: make(map[string]*TaskCompositionShape),

		TaskGroups:           make(map[*TaskGroup]struct{}),
		TaskGroups_mapString: make(map[string]*TaskGroup),

		TaskGroupShapes:           make(map[*TaskGroupShape]struct{}),
		TaskGroupShapes_mapString: make(map[string]*TaskGroupShape),

		TaskInputShapes:           make(map[*TaskInputShape]struct{}),
		TaskInputShapes_mapString: make(map[string]*TaskInputShape),

		TaskOutputShapes:           make(map[*TaskOutputShape]struct{}),
		TaskOutputShapes_mapString: make(map[string]*TaskOutputShape),

		TaskShapes:           make(map[*TaskShape]struct{}),
		TaskShapes_mapString: make(map[string]*TaskShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
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

		NoteResourceShape_stagedOrder: make(map[*NoteResourceShape]uint),
		NoteResourceShape_orderStaged: make(map[uint]*NoteResourceShape),
		NoteResourceShapes_reference:  make(map[*NoteResourceShape]*NoteResourceShape),

		NoteShape_stagedOrder: make(map[*NoteShape]uint),
		NoteShape_orderStaged: make(map[uint]*NoteShape),
		NoteShapes_reference:  make(map[*NoteShape]*NoteShape),

		NoteTaskShape_stagedOrder: make(map[*NoteTaskShape]uint),
		NoteTaskShape_orderStaged: make(map[uint]*NoteTaskShape),
		NoteTaskShapes_reference:  make(map[*NoteTaskShape]*NoteTaskShape),

		Product_stagedOrder: make(map[*Product]uint),
		Product_orderStaged: make(map[uint]*Product),
		Products_reference:  make(map[*Product]*Product),

		ProductCompositionShape_stagedOrder: make(map[*ProductCompositionShape]uint),
		ProductCompositionShape_orderStaged: make(map[uint]*ProductCompositionShape),
		ProductCompositionShapes_reference:  make(map[*ProductCompositionShape]*ProductCompositionShape),

		ProductShape_stagedOrder: make(map[*ProductShape]uint),
		ProductShape_orderStaged: make(map[uint]*ProductShape),
		ProductShapes_reference:  make(map[*ProductShape]*ProductShape),

		Resource_stagedOrder: make(map[*Resource]uint),
		Resource_orderStaged: make(map[uint]*Resource),
		Resources_reference:  make(map[*Resource]*Resource),

		ResourceCompositionShape_stagedOrder: make(map[*ResourceCompositionShape]uint),
		ResourceCompositionShape_orderStaged: make(map[uint]*ResourceCompositionShape),
		ResourceCompositionShapes_reference:  make(map[*ResourceCompositionShape]*ResourceCompositionShape),

		ResourceShape_stagedOrder: make(map[*ResourceShape]uint),
		ResourceShape_orderStaged: make(map[uint]*ResourceShape),
		ResourceShapes_reference:  make(map[*ResourceShape]*ResourceShape),

		ResourceTaskShape_stagedOrder: make(map[*ResourceTaskShape]uint),
		ResourceTaskShape_orderStaged: make(map[uint]*ResourceTaskShape),
		ResourceTaskShapes_reference:  make(map[*ResourceTaskShape]*ResourceTaskShape),

		Task_stagedOrder: make(map[*Task]uint),
		Task_orderStaged: make(map[uint]*Task),
		Tasks_reference:  make(map[*Task]*Task),

		TaskCompositionShape_stagedOrder: make(map[*TaskCompositionShape]uint),
		TaskCompositionShape_orderStaged: make(map[uint]*TaskCompositionShape),
		TaskCompositionShapes_reference:  make(map[*TaskCompositionShape]*TaskCompositionShape),

		TaskGroup_stagedOrder: make(map[*TaskGroup]uint),
		TaskGroup_orderStaged: make(map[uint]*TaskGroup),
		TaskGroups_reference:  make(map[*TaskGroup]*TaskGroup),

		TaskGroupShape_stagedOrder: make(map[*TaskGroupShape]uint),
		TaskGroupShape_orderStaged: make(map[uint]*TaskGroupShape),
		TaskGroupShapes_reference:  make(map[*TaskGroupShape]*TaskGroupShape),

		TaskInputShape_stagedOrder: make(map[*TaskInputShape]uint),
		TaskInputShape_orderStaged: make(map[uint]*TaskInputShape),
		TaskInputShapes_reference:  make(map[*TaskInputShape]*TaskInputShape),

		TaskOutputShape_stagedOrder: make(map[*TaskOutputShape]uint),
		TaskOutputShape_orderStaged: make(map[uint]*TaskOutputShape),
		TaskOutputShapes_reference:  make(map[*TaskOutputShape]*TaskOutputShape),

		TaskShape_stagedOrder: make(map[*TaskShape]uint),
		TaskShape_orderStaged: make(map[uint]*TaskShape),
		TaskShapes_reference:  make(map[*TaskShape]*TaskShape),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Diagram": &DiagramUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteProductShape": &NoteProductShapeUnmarshaller{},

			"NoteResourceShape": &NoteResourceShapeUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"NoteTaskShape": &NoteTaskShapeUnmarshaller{},

			"Product": &ProductUnmarshaller{},

			"ProductCompositionShape": &ProductCompositionShapeUnmarshaller{},

			"ProductShape": &ProductShapeUnmarshaller{},

			"Resource": &ResourceUnmarshaller{},

			"ResourceCompositionShape": &ResourceCompositionShapeUnmarshaller{},

			"ResourceShape": &ResourceShapeUnmarshaller{},

			"ResourceTaskShape": &ResourceTaskShapeUnmarshaller{},

			"Task": &TaskUnmarshaller{},

			"TaskCompositionShape": &TaskCompositionShapeUnmarshaller{},

			"TaskGroup": &TaskGroupUnmarshaller{},

			"TaskGroupShape": &TaskGroupShapeUnmarshaller{},

			"TaskInputShape": &TaskInputShapeUnmarshaller{},

			"TaskOutputShape": &TaskOutputShapeUnmarshaller{},

			"TaskShape": &TaskShapeUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Diagram"},
			{name: "Library"},
			{name: "Note"},
			{name: "NoteProductShape"},
			{name: "NoteResourceShape"},
			{name: "NoteShape"},
			{name: "NoteTaskShape"},
			{name: "Product"},
			{name: "ProductCompositionShape"},
			{name: "ProductShape"},
			{name: "Resource"},
			{name: "ResourceCompositionShape"},
			{name: "ResourceShape"},
			{name: "ResourceTaskShape"},
			{name: "Task"},
			{name: "TaskCompositionShape"},
			{name: "TaskGroup"},
			{name: "TaskGroupShape"},
			{name: "TaskInputShape"},
			{name: "TaskOutputShape"},
			{name: "TaskShape"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteProductShape:
		return stage.NoteProductShape_stagedOrder[instance]
	case *NoteResourceShape:
		return stage.NoteResourceShape_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *NoteTaskShape:
		return stage.NoteTaskShape_stagedOrder[instance]
	case *Product:
		return stage.Product_stagedOrder[instance]
	case *ProductCompositionShape:
		return stage.ProductCompositionShape_stagedOrder[instance]
	case *ProductShape:
		return stage.ProductShape_stagedOrder[instance]
	case *Resource:
		return stage.Resource_stagedOrder[instance]
	case *ResourceCompositionShape:
		return stage.ResourceCompositionShape_stagedOrder[instance]
	case *ResourceShape:
		return stage.ResourceShape_stagedOrder[instance]
	case *ResourceTaskShape:
		return stage.ResourceTaskShape_stagedOrder[instance]
	case *Task:
		return stage.Task_stagedOrder[instance]
	case *TaskCompositionShape:
		return stage.TaskCompositionShape_stagedOrder[instance]
	case *TaskGroup:
		return stage.TaskGroup_stagedOrder[instance]
	case *TaskGroupShape:
		return stage.TaskGroupShape_stagedOrder[instance]
	case *TaskInputShape:
		return stage.TaskInputShape_stagedOrder[instance]
	case *TaskOutputShape:
		return stage.TaskOutputShape_stagedOrder[instance]
	case *TaskShape:
		return stage.TaskShape_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *Diagram:
		return any(stage.Diagram_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Note:
		return any(stage.Note_orderStaged[order]).(Type)
	case *NoteProductShape:
		return any(stage.NoteProductShape_orderStaged[order]).(Type)
	case *NoteResourceShape:
		return any(stage.NoteResourceShape_orderStaged[order]).(Type)
	case *NoteShape:
		return any(stage.NoteShape_orderStaged[order]).(Type)
	case *NoteTaskShape:
		return any(stage.NoteTaskShape_orderStaged[order]).(Type)
	case *Product:
		return any(stage.Product_orderStaged[order]).(Type)
	case *ProductCompositionShape:
		return any(stage.ProductCompositionShape_orderStaged[order]).(Type)
	case *ProductShape:
		return any(stage.ProductShape_orderStaged[order]).(Type)
	case *Resource:
		return any(stage.Resource_orderStaged[order]).(Type)
	case *ResourceCompositionShape:
		return any(stage.ResourceCompositionShape_orderStaged[order]).(Type)
	case *ResourceShape:
		return any(stage.ResourceShape_orderStaged[order]).(Type)
	case *ResourceTaskShape:
		return any(stage.ResourceTaskShape_orderStaged[order]).(Type)
	case *Task:
		return any(stage.Task_orderStaged[order]).(Type)
	case *TaskCompositionShape:
		return any(stage.TaskCompositionShape_orderStaged[order]).(Type)
	case *TaskGroup:
		return any(stage.TaskGroup_orderStaged[order]).(Type)
	case *TaskGroupShape:
		return any(stage.TaskGroupShape_orderStaged[order]).(Type)
	case *TaskInputShape:
		return any(stage.TaskInputShape_orderStaged[order]).(Type)
	case *TaskOutputShape:
		return any(stage.TaskOutputShape_orderStaged[order]).(Type)
	case *TaskShape:
		return any(stage.TaskShape_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteProductShape:
		return stage.NoteProductShape_stagedOrder[instance]
	case *NoteResourceShape:
		return stage.NoteResourceShape_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *NoteTaskShape:
		return stage.NoteTaskShape_stagedOrder[instance]
	case *Product:
		return stage.Product_stagedOrder[instance]
	case *ProductCompositionShape:
		return stage.ProductCompositionShape_stagedOrder[instance]
	case *ProductShape:
		return stage.ProductShape_stagedOrder[instance]
	case *Resource:
		return stage.Resource_stagedOrder[instance]
	case *ResourceCompositionShape:
		return stage.ResourceCompositionShape_stagedOrder[instance]
	case *ResourceShape:
		return stage.ResourceShape_stagedOrder[instance]
	case *ResourceTaskShape:
		return stage.ResourceTaskShape_stagedOrder[instance]
	case *Task:
		return stage.Task_stagedOrder[instance]
	case *TaskCompositionShape:
		return stage.TaskCompositionShape_stagedOrder[instance]
	case *TaskGroup:
		return stage.TaskGroup_stagedOrder[instance]
	case *TaskGroupShape:
		return stage.TaskGroupShape_stagedOrder[instance]
	case *TaskInputShape:
		return stage.TaskInputShape_stagedOrder[instance]
	case *TaskOutputShape:
		return stage.TaskOutputShape_stagedOrder[instance]
	case *TaskShape:
		return stage.TaskShape_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteProductShape"] = len(stage.NoteProductShapes)
	stage.Map_GongStructName_InstancesNb["NoteResourceShape"] = len(stage.NoteResourceShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteTaskShape"] = len(stage.NoteTaskShapes)
	stage.Map_GongStructName_InstancesNb["Product"] = len(stage.Products)
	stage.Map_GongStructName_InstancesNb["ProductCompositionShape"] = len(stage.ProductCompositionShapes)
	stage.Map_GongStructName_InstancesNb["ProductShape"] = len(stage.ProductShapes)
	stage.Map_GongStructName_InstancesNb["Resource"] = len(stage.Resources)
	stage.Map_GongStructName_InstancesNb["ResourceCompositionShape"] = len(stage.ResourceCompositionShapes)
	stage.Map_GongStructName_InstancesNb["ResourceShape"] = len(stage.ResourceShapes)
	stage.Map_GongStructName_InstancesNb["ResourceTaskShape"] = len(stage.ResourceTaskShapes)
	stage.Map_GongStructName_InstancesNb["Task"] = len(stage.Tasks)
	stage.Map_GongStructName_InstancesNb["TaskCompositionShape"] = len(stage.TaskCompositionShapes)
	stage.Map_GongStructName_InstancesNb["TaskGroup"] = len(stage.TaskGroups)
	stage.Map_GongStructName_InstancesNb["TaskGroupShape"] = len(stage.TaskGroupShapes)
	stage.Map_GongStructName_InstancesNb["TaskInputShape"] = len(stage.TaskInputShapes)
	stage.Map_GongStructName_InstancesNb["TaskOutputShape"] = len(stage.TaskOutputShapes)
	stage.Map_GongStructName_InstancesNb["TaskShape"] = len(stage.TaskShapes)
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

// Stage puts noteresourceshape to the model stage
func (noteresourceshape *NoteResourceShape) Stage(stage *Stage) *NoteResourceShape {
	if _, ok := stage.NoteResourceShapes[noteresourceshape]; !ok {
		stage.NoteResourceShapes[noteresourceshape] = struct{}{}
		stage.NoteResourceShape_stagedOrder[noteresourceshape] = stage.NoteResourceShapeOrder
		stage.NoteResourceShape_orderStaged[stage.NoteResourceShapeOrder] = noteresourceshape
		stage.NoteResourceShapeOrder++
	}
	stage.NoteResourceShapes_mapString[noteresourceshape.Name] = noteresourceshape

	return noteresourceshape
}

// StagePreserveOrder puts noteresourceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteResourceShapeOrder
// - update stage.NoteResourceShapeOrder accordingly
func (noteresourceshape *NoteResourceShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteResourceShapes[noteresourceshape]; !ok {
		stage.NoteResourceShapes[noteresourceshape] = struct{}{}

		if order > stage.NoteResourceShapeOrder {
			stage.NoteResourceShapeOrder = order
		}
		stage.NoteResourceShape_stagedOrder[noteresourceshape] = order
		stage.NoteResourceShape_orderStaged[order] = noteresourceshape
		stage.NoteResourceShapeOrder++
	}
	stage.NoteResourceShapes_mapString[noteresourceshape.Name] = noteresourceshape
}

// Unstage removes noteresourceshape off the model stage
func (noteresourceshape *NoteResourceShape) Unstage(stage *Stage) *NoteResourceShape {
	delete(stage.NoteResourceShapes, noteresourceshape)
	// issue1150
	// delete(stage.NoteResourceShape_stagedOrder, noteresourceshape)
	delete(stage.NoteResourceShapes_mapString, noteresourceshape.Name)

	return noteresourceshape
}

// UnstageVoid removes noteresourceshape off the model stage
func (noteresourceshape *NoteResourceShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteResourceShapes, noteresourceshape)
	// issue1150
	// delete(stage.NoteResourceShape_stagedOrder, noteresourceshape)
	delete(stage.NoteResourceShapes_mapString, noteresourceshape.Name)
}

// commit noteresourceshape to the back repo (if it is already staged)
func (noteresourceshape *NoteResourceShape) Commit(stage *Stage) *NoteResourceShape {
	if _, ok := stage.NoteResourceShapes[noteresourceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteResourceShape(noteresourceshape)
		}
	}
	return noteresourceshape
}

func (noteresourceshape *NoteResourceShape) CommitVoid(stage *Stage) {
	noteresourceshape.Commit(stage)
}

func (noteresourceshape *NoteResourceShape) StageVoid(stage *Stage) {
	noteresourceshape.Stage(stage)
}

// Checkout noteresourceshape to the back repo (if it is already staged)
func (noteresourceshape *NoteResourceShape) Checkout(stage *Stage) *NoteResourceShape {
	if _, ok := stage.NoteResourceShapes[noteresourceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteResourceShape(noteresourceshape)
		}
	}
	return noteresourceshape
}

// for satisfaction of GongStruct interface
func (noteresourceshape *NoteResourceShape) GetName() (res string) {
	return noteresourceshape.Name
}

// for satisfaction of GongStruct interface
func (noteresourceshape *NoteResourceShape) SetName(name string) {
	noteresourceshape.Name = name
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

// Stage puts product to the model stage
func (product *Product) Stage(stage *Stage) *Product {
	if _, ok := stage.Products[product]; !ok {
		stage.Products[product] = struct{}{}
		stage.Product_stagedOrder[product] = stage.ProductOrder
		stage.Product_orderStaged[stage.ProductOrder] = product
		stage.ProductOrder++
	}
	stage.Products_mapString[product.Name] = product

	return product
}

// StagePreserveOrder puts product to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductOrder
// - update stage.ProductOrder accordingly
func (product *Product) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Products[product]; !ok {
		stage.Products[product] = struct{}{}

		if order > stage.ProductOrder {
			stage.ProductOrder = order
		}
		stage.Product_stagedOrder[product] = order
		stage.Product_orderStaged[order] = product
		stage.ProductOrder++
	}
	stage.Products_mapString[product.Name] = product
}

// Unstage removes product off the model stage
func (product *Product) Unstage(stage *Stage) *Product {
	delete(stage.Products, product)
	// issue1150
	// delete(stage.Product_stagedOrder, product)
	delete(stage.Products_mapString, product.Name)

	return product
}

// UnstageVoid removes product off the model stage
func (product *Product) UnstageVoid(stage *Stage) {
	delete(stage.Products, product)
	// issue1150
	// delete(stage.Product_stagedOrder, product)
	delete(stage.Products_mapString, product.Name)
}

// commit product to the back repo (if it is already staged)
func (product *Product) Commit(stage *Stage) *Product {
	if _, ok := stage.Products[product]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProduct(product)
		}
	}
	return product
}

func (product *Product) CommitVoid(stage *Stage) {
	product.Commit(stage)
}

func (product *Product) StageVoid(stage *Stage) {
	product.Stage(stage)
}

// Checkout product to the back repo (if it is already staged)
func (product *Product) Checkout(stage *Stage) *Product {
	if _, ok := stage.Products[product]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProduct(product)
		}
	}
	return product
}

// for satisfaction of GongStruct interface
func (product *Product) GetName() (res string) {
	return product.Name
}

// for satisfaction of GongStruct interface
func (product *Product) SetName(name string) {
	product.Name = name
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

// Stage puts resource to the model stage
func (resource *Resource) Stage(stage *Stage) *Resource {
	if _, ok := stage.Resources[resource]; !ok {
		stage.Resources[resource] = struct{}{}
		stage.Resource_stagedOrder[resource] = stage.ResourceOrder
		stage.Resource_orderStaged[stage.ResourceOrder] = resource
		stage.ResourceOrder++
	}
	stage.Resources_mapString[resource.Name] = resource

	return resource
}

// StagePreserveOrder puts resource to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceOrder
// - update stage.ResourceOrder accordingly
func (resource *Resource) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Resources[resource]; !ok {
		stage.Resources[resource] = struct{}{}

		if order > stage.ResourceOrder {
			stage.ResourceOrder = order
		}
		stage.Resource_stagedOrder[resource] = order
		stage.Resource_orderStaged[order] = resource
		stage.ResourceOrder++
	}
	stage.Resources_mapString[resource.Name] = resource
}

// Unstage removes resource off the model stage
func (resource *Resource) Unstage(stage *Stage) *Resource {
	delete(stage.Resources, resource)
	// issue1150
	// delete(stage.Resource_stagedOrder, resource)
	delete(stage.Resources_mapString, resource.Name)

	return resource
}

// UnstageVoid removes resource off the model stage
func (resource *Resource) UnstageVoid(stage *Stage) {
	delete(stage.Resources, resource)
	// issue1150
	// delete(stage.Resource_stagedOrder, resource)
	delete(stage.Resources_mapString, resource.Name)
}

// commit resource to the back repo (if it is already staged)
func (resource *Resource) Commit(stage *Stage) *Resource {
	if _, ok := stage.Resources[resource]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitResource(resource)
		}
	}
	return resource
}

func (resource *Resource) CommitVoid(stage *Stage) {
	resource.Commit(stage)
}

func (resource *Resource) StageVoid(stage *Stage) {
	resource.Stage(stage)
}

// Checkout resource to the back repo (if it is already staged)
func (resource *Resource) Checkout(stage *Stage) *Resource {
	if _, ok := stage.Resources[resource]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutResource(resource)
		}
	}
	return resource
}

// for satisfaction of GongStruct interface
func (resource *Resource) GetName() (res string) {
	return resource.Name
}

// for satisfaction of GongStruct interface
func (resource *Resource) SetName(name string) {
	resource.Name = name
}

// Stage puts resourcecompositionshape to the model stage
func (resourcecompositionshape *ResourceCompositionShape) Stage(stage *Stage) *ResourceCompositionShape {
	if _, ok := stage.ResourceCompositionShapes[resourcecompositionshape]; !ok {
		stage.ResourceCompositionShapes[resourcecompositionshape] = struct{}{}
		stage.ResourceCompositionShape_stagedOrder[resourcecompositionshape] = stage.ResourceCompositionShapeOrder
		stage.ResourceCompositionShape_orderStaged[stage.ResourceCompositionShapeOrder] = resourcecompositionshape
		stage.ResourceCompositionShapeOrder++
	}
	stage.ResourceCompositionShapes_mapString[resourcecompositionshape.Name] = resourcecompositionshape

	return resourcecompositionshape
}

// StagePreserveOrder puts resourcecompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceCompositionShapeOrder
// - update stage.ResourceCompositionShapeOrder accordingly
func (resourcecompositionshape *ResourceCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ResourceCompositionShapes[resourcecompositionshape]; !ok {
		stage.ResourceCompositionShapes[resourcecompositionshape] = struct{}{}

		if order > stage.ResourceCompositionShapeOrder {
			stage.ResourceCompositionShapeOrder = order
		}
		stage.ResourceCompositionShape_stagedOrder[resourcecompositionshape] = order
		stage.ResourceCompositionShape_orderStaged[order] = resourcecompositionshape
		stage.ResourceCompositionShapeOrder++
	}
	stage.ResourceCompositionShapes_mapString[resourcecompositionshape.Name] = resourcecompositionshape
}

// Unstage removes resourcecompositionshape off the model stage
func (resourcecompositionshape *ResourceCompositionShape) Unstage(stage *Stage) *ResourceCompositionShape {
	delete(stage.ResourceCompositionShapes, resourcecompositionshape)
	// issue1150
	// delete(stage.ResourceCompositionShape_stagedOrder, resourcecompositionshape)
	delete(stage.ResourceCompositionShapes_mapString, resourcecompositionshape.Name)

	return resourcecompositionshape
}

// UnstageVoid removes resourcecompositionshape off the model stage
func (resourcecompositionshape *ResourceCompositionShape) UnstageVoid(stage *Stage) {
	delete(stage.ResourceCompositionShapes, resourcecompositionshape)
	// issue1150
	// delete(stage.ResourceCompositionShape_stagedOrder, resourcecompositionshape)
	delete(stage.ResourceCompositionShapes_mapString, resourcecompositionshape.Name)
}

// commit resourcecompositionshape to the back repo (if it is already staged)
func (resourcecompositionshape *ResourceCompositionShape) Commit(stage *Stage) *ResourceCompositionShape {
	if _, ok := stage.ResourceCompositionShapes[resourcecompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitResourceCompositionShape(resourcecompositionshape)
		}
	}
	return resourcecompositionshape
}

func (resourcecompositionshape *ResourceCompositionShape) CommitVoid(stage *Stage) {
	resourcecompositionshape.Commit(stage)
}

func (resourcecompositionshape *ResourceCompositionShape) StageVoid(stage *Stage) {
	resourcecompositionshape.Stage(stage)
}

// Checkout resourcecompositionshape to the back repo (if it is already staged)
func (resourcecompositionshape *ResourceCompositionShape) Checkout(stage *Stage) *ResourceCompositionShape {
	if _, ok := stage.ResourceCompositionShapes[resourcecompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutResourceCompositionShape(resourcecompositionshape)
		}
	}
	return resourcecompositionshape
}

// for satisfaction of GongStruct interface
func (resourcecompositionshape *ResourceCompositionShape) GetName() (res string) {
	return resourcecompositionshape.Name
}

// for satisfaction of GongStruct interface
func (resourcecompositionshape *ResourceCompositionShape) SetName(name string) {
	resourcecompositionshape.Name = name
}

// Stage puts resourceshape to the model stage
func (resourceshape *ResourceShape) Stage(stage *Stage) *ResourceShape {
	if _, ok := stage.ResourceShapes[resourceshape]; !ok {
		stage.ResourceShapes[resourceshape] = struct{}{}
		stage.ResourceShape_stagedOrder[resourceshape] = stage.ResourceShapeOrder
		stage.ResourceShape_orderStaged[stage.ResourceShapeOrder] = resourceshape
		stage.ResourceShapeOrder++
	}
	stage.ResourceShapes_mapString[resourceshape.Name] = resourceshape

	return resourceshape
}

// StagePreserveOrder puts resourceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceShapeOrder
// - update stage.ResourceShapeOrder accordingly
func (resourceshape *ResourceShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ResourceShapes[resourceshape]; !ok {
		stage.ResourceShapes[resourceshape] = struct{}{}

		if order > stage.ResourceShapeOrder {
			stage.ResourceShapeOrder = order
		}
		stage.ResourceShape_stagedOrder[resourceshape] = order
		stage.ResourceShape_orderStaged[order] = resourceshape
		stage.ResourceShapeOrder++
	}
	stage.ResourceShapes_mapString[resourceshape.Name] = resourceshape
}

// Unstage removes resourceshape off the model stage
func (resourceshape *ResourceShape) Unstage(stage *Stage) *ResourceShape {
	delete(stage.ResourceShapes, resourceshape)
	// issue1150
	// delete(stage.ResourceShape_stagedOrder, resourceshape)
	delete(stage.ResourceShapes_mapString, resourceshape.Name)

	return resourceshape
}

// UnstageVoid removes resourceshape off the model stage
func (resourceshape *ResourceShape) UnstageVoid(stage *Stage) {
	delete(stage.ResourceShapes, resourceshape)
	// issue1150
	// delete(stage.ResourceShape_stagedOrder, resourceshape)
	delete(stage.ResourceShapes_mapString, resourceshape.Name)
}

// commit resourceshape to the back repo (if it is already staged)
func (resourceshape *ResourceShape) Commit(stage *Stage) *ResourceShape {
	if _, ok := stage.ResourceShapes[resourceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitResourceShape(resourceshape)
		}
	}
	return resourceshape
}

func (resourceshape *ResourceShape) CommitVoid(stage *Stage) {
	resourceshape.Commit(stage)
}

func (resourceshape *ResourceShape) StageVoid(stage *Stage) {
	resourceshape.Stage(stage)
}

// Checkout resourceshape to the back repo (if it is already staged)
func (resourceshape *ResourceShape) Checkout(stage *Stage) *ResourceShape {
	if _, ok := stage.ResourceShapes[resourceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutResourceShape(resourceshape)
		}
	}
	return resourceshape
}

// for satisfaction of GongStruct interface
func (resourceshape *ResourceShape) GetName() (res string) {
	return resourceshape.Name
}

// for satisfaction of GongStruct interface
func (resourceshape *ResourceShape) SetName(name string) {
	resourceshape.Name = name
}

// Stage puts resourcetaskshape to the model stage
func (resourcetaskshape *ResourceTaskShape) Stage(stage *Stage) *ResourceTaskShape {
	if _, ok := stage.ResourceTaskShapes[resourcetaskshape]; !ok {
		stage.ResourceTaskShapes[resourcetaskshape] = struct{}{}
		stage.ResourceTaskShape_stagedOrder[resourcetaskshape] = stage.ResourceTaskShapeOrder
		stage.ResourceTaskShape_orderStaged[stage.ResourceTaskShapeOrder] = resourcetaskshape
		stage.ResourceTaskShapeOrder++
	}
	stage.ResourceTaskShapes_mapString[resourcetaskshape.Name] = resourcetaskshape

	return resourcetaskshape
}

// StagePreserveOrder puts resourcetaskshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceTaskShapeOrder
// - update stage.ResourceTaskShapeOrder accordingly
func (resourcetaskshape *ResourceTaskShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ResourceTaskShapes[resourcetaskshape]; !ok {
		stage.ResourceTaskShapes[resourcetaskshape] = struct{}{}

		if order > stage.ResourceTaskShapeOrder {
			stage.ResourceTaskShapeOrder = order
		}
		stage.ResourceTaskShape_stagedOrder[resourcetaskshape] = order
		stage.ResourceTaskShape_orderStaged[order] = resourcetaskshape
		stage.ResourceTaskShapeOrder++
	}
	stage.ResourceTaskShapes_mapString[resourcetaskshape.Name] = resourcetaskshape
}

// Unstage removes resourcetaskshape off the model stage
func (resourcetaskshape *ResourceTaskShape) Unstage(stage *Stage) *ResourceTaskShape {
	delete(stage.ResourceTaskShapes, resourcetaskshape)
	// issue1150
	// delete(stage.ResourceTaskShape_stagedOrder, resourcetaskshape)
	delete(stage.ResourceTaskShapes_mapString, resourcetaskshape.Name)

	return resourcetaskshape
}

// UnstageVoid removes resourcetaskshape off the model stage
func (resourcetaskshape *ResourceTaskShape) UnstageVoid(stage *Stage) {
	delete(stage.ResourceTaskShapes, resourcetaskshape)
	// issue1150
	// delete(stage.ResourceTaskShape_stagedOrder, resourcetaskshape)
	delete(stage.ResourceTaskShapes_mapString, resourcetaskshape.Name)
}

// commit resourcetaskshape to the back repo (if it is already staged)
func (resourcetaskshape *ResourceTaskShape) Commit(stage *Stage) *ResourceTaskShape {
	if _, ok := stage.ResourceTaskShapes[resourcetaskshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitResourceTaskShape(resourcetaskshape)
		}
	}
	return resourcetaskshape
}

func (resourcetaskshape *ResourceTaskShape) CommitVoid(stage *Stage) {
	resourcetaskshape.Commit(stage)
}

func (resourcetaskshape *ResourceTaskShape) StageVoid(stage *Stage) {
	resourcetaskshape.Stage(stage)
}

// Checkout resourcetaskshape to the back repo (if it is already staged)
func (resourcetaskshape *ResourceTaskShape) Checkout(stage *Stage) *ResourceTaskShape {
	if _, ok := stage.ResourceTaskShapes[resourcetaskshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutResourceTaskShape(resourcetaskshape)
		}
	}
	return resourcetaskshape
}

// for satisfaction of GongStruct interface
func (resourcetaskshape *ResourceTaskShape) GetName() (res string) {
	return resourcetaskshape.Name
}

// for satisfaction of GongStruct interface
func (resourcetaskshape *ResourceTaskShape) SetName(name string) {
	resourcetaskshape.Name = name
}

// Stage puts task to the model stage
func (task *Task) Stage(stage *Stage) *Task {
	if _, ok := stage.Tasks[task]; !ok {
		stage.Tasks[task] = struct{}{}
		stage.Task_stagedOrder[task] = stage.TaskOrder
		stage.Task_orderStaged[stage.TaskOrder] = task
		stage.TaskOrder++
	}
	stage.Tasks_mapString[task.Name] = task

	return task
}

// StagePreserveOrder puts task to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskOrder
// - update stage.TaskOrder accordingly
func (task *Task) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Tasks[task]; !ok {
		stage.Tasks[task] = struct{}{}

		if order > stage.TaskOrder {
			stage.TaskOrder = order
		}
		stage.Task_stagedOrder[task] = order
		stage.Task_orderStaged[order] = task
		stage.TaskOrder++
	}
	stage.Tasks_mapString[task.Name] = task
}

// Unstage removes task off the model stage
func (task *Task) Unstage(stage *Stage) *Task {
	delete(stage.Tasks, task)
	// issue1150
	// delete(stage.Task_stagedOrder, task)
	delete(stage.Tasks_mapString, task.Name)

	return task
}

// UnstageVoid removes task off the model stage
func (task *Task) UnstageVoid(stage *Stage) {
	delete(stage.Tasks, task)
	// issue1150
	// delete(stage.Task_stagedOrder, task)
	delete(stage.Tasks_mapString, task.Name)
}

// commit task to the back repo (if it is already staged)
func (task *Task) Commit(stage *Stage) *Task {
	if _, ok := stage.Tasks[task]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTask(task)
		}
	}
	return task
}

func (task *Task) CommitVoid(stage *Stage) {
	task.Commit(stage)
}

func (task *Task) StageVoid(stage *Stage) {
	task.Stage(stage)
}

// Checkout task to the back repo (if it is already staged)
func (task *Task) Checkout(stage *Stage) *Task {
	if _, ok := stage.Tasks[task]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTask(task)
		}
	}
	return task
}

// for satisfaction of GongStruct interface
func (task *Task) GetName() (res string) {
	return task.Name
}

// for satisfaction of GongStruct interface
func (task *Task) SetName(name string) {
	task.Name = name
}

// Stage puts taskcompositionshape to the model stage
func (taskcompositionshape *TaskCompositionShape) Stage(stage *Stage) *TaskCompositionShape {
	if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; !ok {
		stage.TaskCompositionShapes[taskcompositionshape] = struct{}{}
		stage.TaskCompositionShape_stagedOrder[taskcompositionshape] = stage.TaskCompositionShapeOrder
		stage.TaskCompositionShape_orderStaged[stage.TaskCompositionShapeOrder] = taskcompositionshape
		stage.TaskCompositionShapeOrder++
	}
	stage.TaskCompositionShapes_mapString[taskcompositionshape.Name] = taskcompositionshape

	return taskcompositionshape
}

// StagePreserveOrder puts taskcompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskCompositionShapeOrder
// - update stage.TaskCompositionShapeOrder accordingly
func (taskcompositionshape *TaskCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; !ok {
		stage.TaskCompositionShapes[taskcompositionshape] = struct{}{}

		if order > stage.TaskCompositionShapeOrder {
			stage.TaskCompositionShapeOrder = order
		}
		stage.TaskCompositionShape_stagedOrder[taskcompositionshape] = order
		stage.TaskCompositionShape_orderStaged[order] = taskcompositionshape
		stage.TaskCompositionShapeOrder++
	}
	stage.TaskCompositionShapes_mapString[taskcompositionshape.Name] = taskcompositionshape
}

// Unstage removes taskcompositionshape off the model stage
func (taskcompositionshape *TaskCompositionShape) Unstage(stage *Stage) *TaskCompositionShape {
	delete(stage.TaskCompositionShapes, taskcompositionshape)
	// issue1150
	// delete(stage.TaskCompositionShape_stagedOrder, taskcompositionshape)
	delete(stage.TaskCompositionShapes_mapString, taskcompositionshape.Name)

	return taskcompositionshape
}

// UnstageVoid removes taskcompositionshape off the model stage
func (taskcompositionshape *TaskCompositionShape) UnstageVoid(stage *Stage) {
	delete(stage.TaskCompositionShapes, taskcompositionshape)
	// issue1150
	// delete(stage.TaskCompositionShape_stagedOrder, taskcompositionshape)
	delete(stage.TaskCompositionShapes_mapString, taskcompositionshape.Name)
}

// commit taskcompositionshape to the back repo (if it is already staged)
func (taskcompositionshape *TaskCompositionShape) Commit(stage *Stage) *TaskCompositionShape {
	if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskCompositionShape(taskcompositionshape)
		}
	}
	return taskcompositionshape
}

func (taskcompositionshape *TaskCompositionShape) CommitVoid(stage *Stage) {
	taskcompositionshape.Commit(stage)
}

func (taskcompositionshape *TaskCompositionShape) StageVoid(stage *Stage) {
	taskcompositionshape.Stage(stage)
}

// Checkout taskcompositionshape to the back repo (if it is already staged)
func (taskcompositionshape *TaskCompositionShape) Checkout(stage *Stage) *TaskCompositionShape {
	if _, ok := stage.TaskCompositionShapes[taskcompositionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskCompositionShape(taskcompositionshape)
		}
	}
	return taskcompositionshape
}

// for satisfaction of GongStruct interface
func (taskcompositionshape *TaskCompositionShape) GetName() (res string) {
	return taskcompositionshape.Name
}

// for satisfaction of GongStruct interface
func (taskcompositionshape *TaskCompositionShape) SetName(name string) {
	taskcompositionshape.Name = name
}

// Stage puts taskgroup to the model stage
func (taskgroup *TaskGroup) Stage(stage *Stage) *TaskGroup {
	if _, ok := stage.TaskGroups[taskgroup]; !ok {
		stage.TaskGroups[taskgroup] = struct{}{}
		stage.TaskGroup_stagedOrder[taskgroup] = stage.TaskGroupOrder
		stage.TaskGroup_orderStaged[stage.TaskGroupOrder] = taskgroup
		stage.TaskGroupOrder++
	}
	stage.TaskGroups_mapString[taskgroup.Name] = taskgroup

	return taskgroup
}

// StagePreserveOrder puts taskgroup to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskGroupOrder
// - update stage.TaskGroupOrder accordingly
func (taskgroup *TaskGroup) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TaskGroups[taskgroup]; !ok {
		stage.TaskGroups[taskgroup] = struct{}{}

		if order > stage.TaskGroupOrder {
			stage.TaskGroupOrder = order
		}
		stage.TaskGroup_stagedOrder[taskgroup] = order
		stage.TaskGroup_orderStaged[order] = taskgroup
		stage.TaskGroupOrder++
	}
	stage.TaskGroups_mapString[taskgroup.Name] = taskgroup
}

// Unstage removes taskgroup off the model stage
func (taskgroup *TaskGroup) Unstage(stage *Stage) *TaskGroup {
	delete(stage.TaskGroups, taskgroup)
	// issue1150
	// delete(stage.TaskGroup_stagedOrder, taskgroup)
	delete(stage.TaskGroups_mapString, taskgroup.Name)

	return taskgroup
}

// UnstageVoid removes taskgroup off the model stage
func (taskgroup *TaskGroup) UnstageVoid(stage *Stage) {
	delete(stage.TaskGroups, taskgroup)
	// issue1150
	// delete(stage.TaskGroup_stagedOrder, taskgroup)
	delete(stage.TaskGroups_mapString, taskgroup.Name)
}

// commit taskgroup to the back repo (if it is already staged)
func (taskgroup *TaskGroup) Commit(stage *Stage) *TaskGroup {
	if _, ok := stage.TaskGroups[taskgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskGroup(taskgroup)
		}
	}
	return taskgroup
}

func (taskgroup *TaskGroup) CommitVoid(stage *Stage) {
	taskgroup.Commit(stage)
}

func (taskgroup *TaskGroup) StageVoid(stage *Stage) {
	taskgroup.Stage(stage)
}

// Checkout taskgroup to the back repo (if it is already staged)
func (taskgroup *TaskGroup) Checkout(stage *Stage) *TaskGroup {
	if _, ok := stage.TaskGroups[taskgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskGroup(taskgroup)
		}
	}
	return taskgroup
}

// for satisfaction of GongStruct interface
func (taskgroup *TaskGroup) GetName() (res string) {
	return taskgroup.Name
}

// for satisfaction of GongStruct interface
func (taskgroup *TaskGroup) SetName(name string) {
	taskgroup.Name = name
}

// Stage puts taskgroupshape to the model stage
func (taskgroupshape *TaskGroupShape) Stage(stage *Stage) *TaskGroupShape {
	if _, ok := stage.TaskGroupShapes[taskgroupshape]; !ok {
		stage.TaskGroupShapes[taskgroupshape] = struct{}{}
		stage.TaskGroupShape_stagedOrder[taskgroupshape] = stage.TaskGroupShapeOrder
		stage.TaskGroupShape_orderStaged[stage.TaskGroupShapeOrder] = taskgroupshape
		stage.TaskGroupShapeOrder++
	}
	stage.TaskGroupShapes_mapString[taskgroupshape.Name] = taskgroupshape

	return taskgroupshape
}

// StagePreserveOrder puts taskgroupshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskGroupShapeOrder
// - update stage.TaskGroupShapeOrder accordingly
func (taskgroupshape *TaskGroupShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TaskGroupShapes[taskgroupshape]; !ok {
		stage.TaskGroupShapes[taskgroupshape] = struct{}{}

		if order > stage.TaskGroupShapeOrder {
			stage.TaskGroupShapeOrder = order
		}
		stage.TaskGroupShape_stagedOrder[taskgroupshape] = order
		stage.TaskGroupShape_orderStaged[order] = taskgroupshape
		stage.TaskGroupShapeOrder++
	}
	stage.TaskGroupShapes_mapString[taskgroupshape.Name] = taskgroupshape
}

// Unstage removes taskgroupshape off the model stage
func (taskgroupshape *TaskGroupShape) Unstage(stage *Stage) *TaskGroupShape {
	delete(stage.TaskGroupShapes, taskgroupshape)
	// issue1150
	// delete(stage.TaskGroupShape_stagedOrder, taskgroupshape)
	delete(stage.TaskGroupShapes_mapString, taskgroupshape.Name)

	return taskgroupshape
}

// UnstageVoid removes taskgroupshape off the model stage
func (taskgroupshape *TaskGroupShape) UnstageVoid(stage *Stage) {
	delete(stage.TaskGroupShapes, taskgroupshape)
	// issue1150
	// delete(stage.TaskGroupShape_stagedOrder, taskgroupshape)
	delete(stage.TaskGroupShapes_mapString, taskgroupshape.Name)
}

// commit taskgroupshape to the back repo (if it is already staged)
func (taskgroupshape *TaskGroupShape) Commit(stage *Stage) *TaskGroupShape {
	if _, ok := stage.TaskGroupShapes[taskgroupshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskGroupShape(taskgroupshape)
		}
	}
	return taskgroupshape
}

func (taskgroupshape *TaskGroupShape) CommitVoid(stage *Stage) {
	taskgroupshape.Commit(stage)
}

func (taskgroupshape *TaskGroupShape) StageVoid(stage *Stage) {
	taskgroupshape.Stage(stage)
}

// Checkout taskgroupshape to the back repo (if it is already staged)
func (taskgroupshape *TaskGroupShape) Checkout(stage *Stage) *TaskGroupShape {
	if _, ok := stage.TaskGroupShapes[taskgroupshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskGroupShape(taskgroupshape)
		}
	}
	return taskgroupshape
}

// for satisfaction of GongStruct interface
func (taskgroupshape *TaskGroupShape) GetName() (res string) {
	return taskgroupshape.Name
}

// for satisfaction of GongStruct interface
func (taskgroupshape *TaskGroupShape) SetName(name string) {
	taskgroupshape.Name = name
}

// Stage puts taskinputshape to the model stage
func (taskinputshape *TaskInputShape) Stage(stage *Stage) *TaskInputShape {
	if _, ok := stage.TaskInputShapes[taskinputshape]; !ok {
		stage.TaskInputShapes[taskinputshape] = struct{}{}
		stage.TaskInputShape_stagedOrder[taskinputshape] = stage.TaskInputShapeOrder
		stage.TaskInputShape_orderStaged[stage.TaskInputShapeOrder] = taskinputshape
		stage.TaskInputShapeOrder++
	}
	stage.TaskInputShapes_mapString[taskinputshape.Name] = taskinputshape

	return taskinputshape
}

// StagePreserveOrder puts taskinputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskInputShapeOrder
// - update stage.TaskInputShapeOrder accordingly
func (taskinputshape *TaskInputShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TaskInputShapes[taskinputshape]; !ok {
		stage.TaskInputShapes[taskinputshape] = struct{}{}

		if order > stage.TaskInputShapeOrder {
			stage.TaskInputShapeOrder = order
		}
		stage.TaskInputShape_stagedOrder[taskinputshape] = order
		stage.TaskInputShape_orderStaged[order] = taskinputshape
		stage.TaskInputShapeOrder++
	}
	stage.TaskInputShapes_mapString[taskinputshape.Name] = taskinputshape
}

// Unstage removes taskinputshape off the model stage
func (taskinputshape *TaskInputShape) Unstage(stage *Stage) *TaskInputShape {
	delete(stage.TaskInputShapes, taskinputshape)
	// issue1150
	// delete(stage.TaskInputShape_stagedOrder, taskinputshape)
	delete(stage.TaskInputShapes_mapString, taskinputshape.Name)

	return taskinputshape
}

// UnstageVoid removes taskinputshape off the model stage
func (taskinputshape *TaskInputShape) UnstageVoid(stage *Stage) {
	delete(stage.TaskInputShapes, taskinputshape)
	// issue1150
	// delete(stage.TaskInputShape_stagedOrder, taskinputshape)
	delete(stage.TaskInputShapes_mapString, taskinputshape.Name)
}

// commit taskinputshape to the back repo (if it is already staged)
func (taskinputshape *TaskInputShape) Commit(stage *Stage) *TaskInputShape {
	if _, ok := stage.TaskInputShapes[taskinputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskInputShape(taskinputshape)
		}
	}
	return taskinputshape
}

func (taskinputshape *TaskInputShape) CommitVoid(stage *Stage) {
	taskinputshape.Commit(stage)
}

func (taskinputshape *TaskInputShape) StageVoid(stage *Stage) {
	taskinputshape.Stage(stage)
}

// Checkout taskinputshape to the back repo (if it is already staged)
func (taskinputshape *TaskInputShape) Checkout(stage *Stage) *TaskInputShape {
	if _, ok := stage.TaskInputShapes[taskinputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskInputShape(taskinputshape)
		}
	}
	return taskinputshape
}

// for satisfaction of GongStruct interface
func (taskinputshape *TaskInputShape) GetName() (res string) {
	return taskinputshape.Name
}

// for satisfaction of GongStruct interface
func (taskinputshape *TaskInputShape) SetName(name string) {
	taskinputshape.Name = name
}

// Stage puts taskoutputshape to the model stage
func (taskoutputshape *TaskOutputShape) Stage(stage *Stage) *TaskOutputShape {
	if _, ok := stage.TaskOutputShapes[taskoutputshape]; !ok {
		stage.TaskOutputShapes[taskoutputshape] = struct{}{}
		stage.TaskOutputShape_stagedOrder[taskoutputshape] = stage.TaskOutputShapeOrder
		stage.TaskOutputShape_orderStaged[stage.TaskOutputShapeOrder] = taskoutputshape
		stage.TaskOutputShapeOrder++
	}
	stage.TaskOutputShapes_mapString[taskoutputshape.Name] = taskoutputshape

	return taskoutputshape
}

// StagePreserveOrder puts taskoutputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskOutputShapeOrder
// - update stage.TaskOutputShapeOrder accordingly
func (taskoutputshape *TaskOutputShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TaskOutputShapes[taskoutputshape]; !ok {
		stage.TaskOutputShapes[taskoutputshape] = struct{}{}

		if order > stage.TaskOutputShapeOrder {
			stage.TaskOutputShapeOrder = order
		}
		stage.TaskOutputShape_stagedOrder[taskoutputshape] = order
		stage.TaskOutputShape_orderStaged[order] = taskoutputshape
		stage.TaskOutputShapeOrder++
	}
	stage.TaskOutputShapes_mapString[taskoutputshape.Name] = taskoutputshape
}

// Unstage removes taskoutputshape off the model stage
func (taskoutputshape *TaskOutputShape) Unstage(stage *Stage) *TaskOutputShape {
	delete(stage.TaskOutputShapes, taskoutputshape)
	// issue1150
	// delete(stage.TaskOutputShape_stagedOrder, taskoutputshape)
	delete(stage.TaskOutputShapes_mapString, taskoutputshape.Name)

	return taskoutputshape
}

// UnstageVoid removes taskoutputshape off the model stage
func (taskoutputshape *TaskOutputShape) UnstageVoid(stage *Stage) {
	delete(stage.TaskOutputShapes, taskoutputshape)
	// issue1150
	// delete(stage.TaskOutputShape_stagedOrder, taskoutputshape)
	delete(stage.TaskOutputShapes_mapString, taskoutputshape.Name)
}

// commit taskoutputshape to the back repo (if it is already staged)
func (taskoutputshape *TaskOutputShape) Commit(stage *Stage) *TaskOutputShape {
	if _, ok := stage.TaskOutputShapes[taskoutputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskOutputShape(taskoutputshape)
		}
	}
	return taskoutputshape
}

func (taskoutputshape *TaskOutputShape) CommitVoid(stage *Stage) {
	taskoutputshape.Commit(stage)
}

func (taskoutputshape *TaskOutputShape) StageVoid(stage *Stage) {
	taskoutputshape.Stage(stage)
}

// Checkout taskoutputshape to the back repo (if it is already staged)
func (taskoutputshape *TaskOutputShape) Checkout(stage *Stage) *TaskOutputShape {
	if _, ok := stage.TaskOutputShapes[taskoutputshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskOutputShape(taskoutputshape)
		}
	}
	return taskoutputshape
}

// for satisfaction of GongStruct interface
func (taskoutputshape *TaskOutputShape) GetName() (res string) {
	return taskoutputshape.Name
}

// for satisfaction of GongStruct interface
func (taskoutputshape *TaskOutputShape) SetName(name string) {
	taskoutputshape.Name = name
}

// Stage puts taskshape to the model stage
func (taskshape *TaskShape) Stage(stage *Stage) *TaskShape {
	if _, ok := stage.TaskShapes[taskshape]; !ok {
		stage.TaskShapes[taskshape] = struct{}{}
		stage.TaskShape_stagedOrder[taskshape] = stage.TaskShapeOrder
		stage.TaskShape_orderStaged[stage.TaskShapeOrder] = taskshape
		stage.TaskShapeOrder++
	}
	stage.TaskShapes_mapString[taskshape.Name] = taskshape

	return taskshape
}

// StagePreserveOrder puts taskshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskShapeOrder
// - update stage.TaskShapeOrder accordingly
func (taskshape *TaskShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.TaskShapes[taskshape]; !ok {
		stage.TaskShapes[taskshape] = struct{}{}

		if order > stage.TaskShapeOrder {
			stage.TaskShapeOrder = order
		}
		stage.TaskShape_stagedOrder[taskshape] = order
		stage.TaskShape_orderStaged[order] = taskshape
		stage.TaskShapeOrder++
	}
	stage.TaskShapes_mapString[taskshape.Name] = taskshape
}

// Unstage removes taskshape off the model stage
func (taskshape *TaskShape) Unstage(stage *Stage) *TaskShape {
	delete(stage.TaskShapes, taskshape)
	// issue1150
	// delete(stage.TaskShape_stagedOrder, taskshape)
	delete(stage.TaskShapes_mapString, taskshape.Name)

	return taskshape
}

// UnstageVoid removes taskshape off the model stage
func (taskshape *TaskShape) UnstageVoid(stage *Stage) {
	delete(stage.TaskShapes, taskshape)
	// issue1150
	// delete(stage.TaskShape_stagedOrder, taskshape)
	delete(stage.TaskShapes_mapString, taskshape.Name)
}

// commit taskshape to the back repo (if it is already staged)
func (taskshape *TaskShape) Commit(stage *Stage) *TaskShape {
	if _, ok := stage.TaskShapes[taskshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTaskShape(taskshape)
		}
	}
	return taskshape
}

func (taskshape *TaskShape) CommitVoid(stage *Stage) {
	taskshape.Commit(stage)
}

func (taskshape *TaskShape) StageVoid(stage *Stage) {
	taskshape.Stage(stage)
}

// Checkout taskshape to the back repo (if it is already staged)
func (taskshape *TaskShape) Checkout(stage *Stage) *TaskShape {
	if _, ok := stage.TaskShapes[taskshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTaskShape(taskshape)
		}
	}
	return taskshape
}

// for satisfaction of GongStruct interface
func (taskshape *TaskShape) GetName() (res string) {
	return taskshape.Name
}

// for satisfaction of GongStruct interface
func (taskshape *TaskShape) SetName(name string) {
	taskshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMDiagram(Diagram *Diagram)
	CreateORMLibrary(Library *Library)
	CreateORMNote(Note *Note)
	CreateORMNoteProductShape(NoteProductShape *NoteProductShape)
	CreateORMNoteResourceShape(NoteResourceShape *NoteResourceShape)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMNoteTaskShape(NoteTaskShape *NoteTaskShape)
	CreateORMProduct(Product *Product)
	CreateORMProductCompositionShape(ProductCompositionShape *ProductCompositionShape)
	CreateORMProductShape(ProductShape *ProductShape)
	CreateORMResource(Resource *Resource)
	CreateORMResourceCompositionShape(ResourceCompositionShape *ResourceCompositionShape)
	CreateORMResourceShape(ResourceShape *ResourceShape)
	CreateORMResourceTaskShape(ResourceTaskShape *ResourceTaskShape)
	CreateORMTask(Task *Task)
	CreateORMTaskCompositionShape(TaskCompositionShape *TaskCompositionShape)
	CreateORMTaskGroup(TaskGroup *TaskGroup)
	CreateORMTaskGroupShape(TaskGroupShape *TaskGroupShape)
	CreateORMTaskInputShape(TaskInputShape *TaskInputShape)
	CreateORMTaskOutputShape(TaskOutputShape *TaskOutputShape)
	CreateORMTaskShape(TaskShape *TaskShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMDiagram(Diagram *Diagram)
	DeleteORMLibrary(Library *Library)
	DeleteORMNote(Note *Note)
	DeleteORMNoteProductShape(NoteProductShape *NoteProductShape)
	DeleteORMNoteResourceShape(NoteResourceShape *NoteResourceShape)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMNoteTaskShape(NoteTaskShape *NoteTaskShape)
	DeleteORMProduct(Product *Product)
	DeleteORMProductCompositionShape(ProductCompositionShape *ProductCompositionShape)
	DeleteORMProductShape(ProductShape *ProductShape)
	DeleteORMResource(Resource *Resource)
	DeleteORMResourceCompositionShape(ResourceCompositionShape *ResourceCompositionShape)
	DeleteORMResourceShape(ResourceShape *ResourceShape)
	DeleteORMResourceTaskShape(ResourceTaskShape *ResourceTaskShape)
	DeleteORMTask(Task *Task)
	DeleteORMTaskCompositionShape(TaskCompositionShape *TaskCompositionShape)
	DeleteORMTaskGroup(TaskGroup *TaskGroup)
	DeleteORMTaskGroupShape(TaskGroupShape *TaskGroupShape)
	DeleteORMTaskInputShape(TaskInputShape *TaskInputShape)
	DeleteORMTaskOutputShape(TaskOutputShape *TaskOutputShape)
	DeleteORMTaskShape(TaskShape *TaskShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
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

	stage.NoteResourceShapes = make(map[*NoteResourceShape]struct{})
	stage.NoteResourceShapes_mapString = make(map[string]*NoteResourceShape)
	stage.NoteResourceShape_stagedOrder = make(map[*NoteResourceShape]uint)
	stage.NoteResourceShapeOrder = 0

	stage.NoteShapes = make(map[*NoteShape]struct{})
	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	stage.NoteShape_stagedOrder = make(map[*NoteShape]uint)
	stage.NoteShapeOrder = 0

	stage.NoteTaskShapes = make(map[*NoteTaskShape]struct{})
	stage.NoteTaskShapes_mapString = make(map[string]*NoteTaskShape)
	stage.NoteTaskShape_stagedOrder = make(map[*NoteTaskShape]uint)
	stage.NoteTaskShapeOrder = 0

	stage.Products = make(map[*Product]struct{})
	stage.Products_mapString = make(map[string]*Product)
	stage.Product_stagedOrder = make(map[*Product]uint)
	stage.ProductOrder = 0

	stage.ProductCompositionShapes = make(map[*ProductCompositionShape]struct{})
	stage.ProductCompositionShapes_mapString = make(map[string]*ProductCompositionShape)
	stage.ProductCompositionShape_stagedOrder = make(map[*ProductCompositionShape]uint)
	stage.ProductCompositionShapeOrder = 0

	stage.ProductShapes = make(map[*ProductShape]struct{})
	stage.ProductShapes_mapString = make(map[string]*ProductShape)
	stage.ProductShape_stagedOrder = make(map[*ProductShape]uint)
	stage.ProductShapeOrder = 0

	stage.Resources = make(map[*Resource]struct{})
	stage.Resources_mapString = make(map[string]*Resource)
	stage.Resource_stagedOrder = make(map[*Resource]uint)
	stage.ResourceOrder = 0

	stage.ResourceCompositionShapes = make(map[*ResourceCompositionShape]struct{})
	stage.ResourceCompositionShapes_mapString = make(map[string]*ResourceCompositionShape)
	stage.ResourceCompositionShape_stagedOrder = make(map[*ResourceCompositionShape]uint)
	stage.ResourceCompositionShapeOrder = 0

	stage.ResourceShapes = make(map[*ResourceShape]struct{})
	stage.ResourceShapes_mapString = make(map[string]*ResourceShape)
	stage.ResourceShape_stagedOrder = make(map[*ResourceShape]uint)
	stage.ResourceShapeOrder = 0

	stage.ResourceTaskShapes = make(map[*ResourceTaskShape]struct{})
	stage.ResourceTaskShapes_mapString = make(map[string]*ResourceTaskShape)
	stage.ResourceTaskShape_stagedOrder = make(map[*ResourceTaskShape]uint)
	stage.ResourceTaskShapeOrder = 0

	stage.Tasks = make(map[*Task]struct{})
	stage.Tasks_mapString = make(map[string]*Task)
	stage.Task_stagedOrder = make(map[*Task]uint)
	stage.TaskOrder = 0

	stage.TaskCompositionShapes = make(map[*TaskCompositionShape]struct{})
	stage.TaskCompositionShapes_mapString = make(map[string]*TaskCompositionShape)
	stage.TaskCompositionShape_stagedOrder = make(map[*TaskCompositionShape]uint)
	stage.TaskCompositionShapeOrder = 0

	stage.TaskGroups = make(map[*TaskGroup]struct{})
	stage.TaskGroups_mapString = make(map[string]*TaskGroup)
	stage.TaskGroup_stagedOrder = make(map[*TaskGroup]uint)
	stage.TaskGroupOrder = 0

	stage.TaskGroupShapes = make(map[*TaskGroupShape]struct{})
	stage.TaskGroupShapes_mapString = make(map[string]*TaskGroupShape)
	stage.TaskGroupShape_stagedOrder = make(map[*TaskGroupShape]uint)
	stage.TaskGroupShapeOrder = 0

	stage.TaskInputShapes = make(map[*TaskInputShape]struct{})
	stage.TaskInputShapes_mapString = make(map[string]*TaskInputShape)
	stage.TaskInputShape_stagedOrder = make(map[*TaskInputShape]uint)
	stage.TaskInputShapeOrder = 0

	stage.TaskOutputShapes = make(map[*TaskOutputShape]struct{})
	stage.TaskOutputShapes_mapString = make(map[string]*TaskOutputShape)
	stage.TaskOutputShape_stagedOrder = make(map[*TaskOutputShape]uint)
	stage.TaskOutputShapeOrder = 0

	stage.TaskShapes = make(map[*TaskShape]struct{})
	stage.TaskShapes_mapString = make(map[string]*TaskShape)
	stage.TaskShape_stagedOrder = make(map[*TaskShape]uint)
	stage.TaskShapeOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Diagrams = nil
	stage.Diagrams_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Notes = nil
	stage.Notes_mapString = nil

	stage.NoteProductShapes = nil
	stage.NoteProductShapes_mapString = nil

	stage.NoteResourceShapes = nil
	stage.NoteResourceShapes_mapString = nil

	stage.NoteShapes = nil
	stage.NoteShapes_mapString = nil

	stage.NoteTaskShapes = nil
	stage.NoteTaskShapes_mapString = nil

	stage.Products = nil
	stage.Products_mapString = nil

	stage.ProductCompositionShapes = nil
	stage.ProductCompositionShapes_mapString = nil

	stage.ProductShapes = nil
	stage.ProductShapes_mapString = nil

	stage.Resources = nil
	stage.Resources_mapString = nil

	stage.ResourceCompositionShapes = nil
	stage.ResourceCompositionShapes_mapString = nil

	stage.ResourceShapes = nil
	stage.ResourceShapes_mapString = nil

	stage.ResourceTaskShapes = nil
	stage.ResourceTaskShapes_mapString = nil

	stage.Tasks = nil
	stage.Tasks_mapString = nil

	stage.TaskCompositionShapes = nil
	stage.TaskCompositionShapes_mapString = nil

	stage.TaskGroups = nil
	stage.TaskGroups_mapString = nil

	stage.TaskGroupShapes = nil
	stage.TaskGroupShapes_mapString = nil

	stage.TaskInputShapes = nil
	stage.TaskInputShapes_mapString = nil

	stage.TaskOutputShapes = nil
	stage.TaskOutputShapes_mapString = nil

	stage.TaskShapes = nil
	stage.TaskShapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
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

	for noteresourceshape := range stage.NoteResourceShapes {
		noteresourceshape.Unstage(stage)
	}

	for noteshape := range stage.NoteShapes {
		noteshape.Unstage(stage)
	}

	for notetaskshape := range stage.NoteTaskShapes {
		notetaskshape.Unstage(stage)
	}

	for product := range stage.Products {
		product.Unstage(stage)
	}

	for productcompositionshape := range stage.ProductCompositionShapes {
		productcompositionshape.Unstage(stage)
	}

	for productshape := range stage.ProductShapes {
		productshape.Unstage(stage)
	}

	for resource := range stage.Resources {
		resource.Unstage(stage)
	}

	for resourcecompositionshape := range stage.ResourceCompositionShapes {
		resourcecompositionshape.Unstage(stage)
	}

	for resourceshape := range stage.ResourceShapes {
		resourceshape.Unstage(stage)
	}

	for resourcetaskshape := range stage.ResourceTaskShapes {
		resourcetaskshape.Unstage(stage)
	}

	for task := range stage.Tasks {
		task.Unstage(stage)
	}

	for taskcompositionshape := range stage.TaskCompositionShapes {
		taskcompositionshape.Unstage(stage)
	}

	for taskgroup := range stage.TaskGroups {
		taskgroup.Unstage(stage)
	}

	for taskgroupshape := range stage.TaskGroupShapes {
		taskgroupshape.Unstage(stage)
	}

	for taskinputshape := range stage.TaskInputShapes {
		taskinputshape.Unstage(stage)
	}

	for taskoutputshape := range stage.TaskOutputShapes {
		taskoutputshape.Unstage(stage)
	}

	for taskshape := range stage.TaskShapes {
		taskshape.Unstage(stage)
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
	case map[*Diagram]any:
		return any(&stage.Diagrams).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Note]any:
		return any(&stage.Notes).(*Type)
	case map[*NoteProductShape]any:
		return any(&stage.NoteProductShapes).(*Type)
	case map[*NoteResourceShape]any:
		return any(&stage.NoteResourceShapes).(*Type)
	case map[*NoteShape]any:
		return any(&stage.NoteShapes).(*Type)
	case map[*NoteTaskShape]any:
		return any(&stage.NoteTaskShapes).(*Type)
	case map[*Product]any:
		return any(&stage.Products).(*Type)
	case map[*ProductCompositionShape]any:
		return any(&stage.ProductCompositionShapes).(*Type)
	case map[*ProductShape]any:
		return any(&stage.ProductShapes).(*Type)
	case map[*Resource]any:
		return any(&stage.Resources).(*Type)
	case map[*ResourceCompositionShape]any:
		return any(&stage.ResourceCompositionShapes).(*Type)
	case map[*ResourceShape]any:
		return any(&stage.ResourceShapes).(*Type)
	case map[*ResourceTaskShape]any:
		return any(&stage.ResourceTaskShapes).(*Type)
	case map[*Task]any:
		return any(&stage.Tasks).(*Type)
	case map[*TaskCompositionShape]any:
		return any(&stage.TaskCompositionShapes).(*Type)
	case map[*TaskGroup]any:
		return any(&stage.TaskGroups).(*Type)
	case map[*TaskGroupShape]any:
		return any(&stage.TaskGroupShapes).(*Type)
	case map[*TaskInputShape]any:
		return any(&stage.TaskInputShapes).(*Type)
	case map[*TaskOutputShape]any:
		return any(&stage.TaskOutputShapes).(*Type)
	case map[*TaskShape]any:
		return any(&stage.TaskShapes).(*Type)
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
	case *Diagram:
		return any(stage.Diagrams_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Note:
		return any(stage.Notes_mapString).(map[string]Type)
	case *NoteProductShape:
		return any(stage.NoteProductShapes_mapString).(map[string]Type)
	case *NoteResourceShape:
		return any(stage.NoteResourceShapes_mapString).(map[string]Type)
	case *NoteShape:
		return any(stage.NoteShapes_mapString).(map[string]Type)
	case *NoteTaskShape:
		return any(stage.NoteTaskShapes_mapString).(map[string]Type)
	case *Product:
		return any(stage.Products_mapString).(map[string]Type)
	case *ProductCompositionShape:
		return any(stage.ProductCompositionShapes_mapString).(map[string]Type)
	case *ProductShape:
		return any(stage.ProductShapes_mapString).(map[string]Type)
	case *Resource:
		return any(stage.Resources_mapString).(map[string]Type)
	case *ResourceCompositionShape:
		return any(stage.ResourceCompositionShapes_mapString).(map[string]Type)
	case *ResourceShape:
		return any(stage.ResourceShapes_mapString).(map[string]Type)
	case *ResourceTaskShape:
		return any(stage.ResourceTaskShapes_mapString).(map[string]Type)
	case *Task:
		return any(stage.Tasks_mapString).(map[string]Type)
	case *TaskCompositionShape:
		return any(stage.TaskCompositionShapes_mapString).(map[string]Type)
	case *TaskGroup:
		return any(stage.TaskGroups_mapString).(map[string]Type)
	case *TaskGroupShape:
		return any(stage.TaskGroupShapes_mapString).(map[string]Type)
	case *TaskInputShape:
		return any(stage.TaskInputShapes_mapString).(map[string]Type)
	case *TaskOutputShape:
		return any(stage.TaskOutputShapes_mapString).(map[string]Type)
	case *TaskShape:
		return any(stage.TaskShapes_mapString).(map[string]Type)
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
	case Diagram:
		return any(&stage.Diagrams).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Note:
		return any(&stage.Notes).(*map[*Type]struct{})
	case NoteProductShape:
		return any(&stage.NoteProductShapes).(*map[*Type]struct{})
	case NoteResourceShape:
		return any(&stage.NoteResourceShapes).(*map[*Type]struct{})
	case NoteShape:
		return any(&stage.NoteShapes).(*map[*Type]struct{})
	case NoteTaskShape:
		return any(&stage.NoteTaskShapes).(*map[*Type]struct{})
	case Product:
		return any(&stage.Products).(*map[*Type]struct{})
	case ProductCompositionShape:
		return any(&stage.ProductCompositionShapes).(*map[*Type]struct{})
	case ProductShape:
		return any(&stage.ProductShapes).(*map[*Type]struct{})
	case Resource:
		return any(&stage.Resources).(*map[*Type]struct{})
	case ResourceCompositionShape:
		return any(&stage.ResourceCompositionShapes).(*map[*Type]struct{})
	case ResourceShape:
		return any(&stage.ResourceShapes).(*map[*Type]struct{})
	case ResourceTaskShape:
		return any(&stage.ResourceTaskShapes).(*map[*Type]struct{})
	case Task:
		return any(&stage.Tasks).(*map[*Type]struct{})
	case TaskCompositionShape:
		return any(&stage.TaskCompositionShapes).(*map[*Type]struct{})
	case TaskGroup:
		return any(&stage.TaskGroups).(*map[*Type]struct{})
	case TaskGroupShape:
		return any(&stage.TaskGroupShapes).(*map[*Type]struct{})
	case TaskInputShape:
		return any(&stage.TaskInputShapes).(*map[*Type]struct{})
	case TaskOutputShape:
		return any(&stage.TaskOutputShapes).(*map[*Type]struct{})
	case TaskShape:
		return any(&stage.TaskShapes).(*map[*Type]struct{})
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
	case *Diagram:
		return any(&stage.Diagrams).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Note:
		return any(&stage.Notes).(*map[Type]struct{})
	case *NoteProductShape:
		return any(&stage.NoteProductShapes).(*map[Type]struct{})
	case *NoteResourceShape:
		return any(&stage.NoteResourceShapes).(*map[Type]struct{})
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]struct{})
	case *NoteTaskShape:
		return any(&stage.NoteTaskShapes).(*map[Type]struct{})
	case *Product:
		return any(&stage.Products).(*map[Type]struct{})
	case *ProductCompositionShape:
		return any(&stage.ProductCompositionShapes).(*map[Type]struct{})
	case *ProductShape:
		return any(&stage.ProductShapes).(*map[Type]struct{})
	case *Resource:
		return any(&stage.Resources).(*map[Type]struct{})
	case *ResourceCompositionShape:
		return any(&stage.ResourceCompositionShapes).(*map[Type]struct{})
	case *ResourceShape:
		return any(&stage.ResourceShapes).(*map[Type]struct{})
	case *ResourceTaskShape:
		return any(&stage.ResourceTaskShapes).(*map[Type]struct{})
	case *Task:
		return any(&stage.Tasks).(*map[Type]struct{})
	case *TaskCompositionShape:
		return any(&stage.TaskCompositionShapes).(*map[Type]struct{})
	case *TaskGroup:
		return any(&stage.TaskGroups).(*map[Type]struct{})
	case *TaskGroupShape:
		return any(&stage.TaskGroupShapes).(*map[Type]struct{})
	case *TaskInputShape:
		return any(&stage.TaskInputShapes).(*map[Type]struct{})
	case *TaskOutputShape:
		return any(&stage.TaskOutputShapes).(*map[Type]struct{})
	case *TaskShape:
		return any(&stage.TaskShapes).(*map[Type]struct{})
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
	case Diagram:
		return any(&stage.Diagrams_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Note:
		return any(&stage.Notes_mapString).(*map[string]*Type)
	case NoteProductShape:
		return any(&stage.NoteProductShapes_mapString).(*map[string]*Type)
	case NoteResourceShape:
		return any(&stage.NoteResourceShapes_mapString).(*map[string]*Type)
	case NoteShape:
		return any(&stage.NoteShapes_mapString).(*map[string]*Type)
	case NoteTaskShape:
		return any(&stage.NoteTaskShapes_mapString).(*map[string]*Type)
	case Product:
		return any(&stage.Products_mapString).(*map[string]*Type)
	case ProductCompositionShape:
		return any(&stage.ProductCompositionShapes_mapString).(*map[string]*Type)
	case ProductShape:
		return any(&stage.ProductShapes_mapString).(*map[string]*Type)
	case Resource:
		return any(&stage.Resources_mapString).(*map[string]*Type)
	case ResourceCompositionShape:
		return any(&stage.ResourceCompositionShapes_mapString).(*map[string]*Type)
	case ResourceShape:
		return any(&stage.ResourceShapes_mapString).(*map[string]*Type)
	case ResourceTaskShape:
		return any(&stage.ResourceTaskShapes_mapString).(*map[string]*Type)
	case Task:
		return any(&stage.Tasks_mapString).(*map[string]*Type)
	case TaskCompositionShape:
		return any(&stage.TaskCompositionShapes_mapString).(*map[string]*Type)
	case TaskGroup:
		return any(&stage.TaskGroups_mapString).(*map[string]*Type)
	case TaskGroupShape:
		return any(&stage.TaskGroupShapes_mapString).(*map[string]*Type)
	case TaskInputShape:
		return any(&stage.TaskInputShapes_mapString).(*map[string]*Type)
	case TaskOutputShape:
		return any(&stage.TaskOutputShapes_mapString).(*map[string]*Type)
	case TaskShape:
		return any(&stage.TaskShapes_mapString).(*map[string]*Type)
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
	case Diagram:
		return any(&Diagram{
			// Initialisation of associations
			// field is initialized with an instance of ProductShape with the name of the field
			Product_Shapes: []*ProductShape{{Name: "Product_Shapes"}},
			// field is initialized with an instance of Product with the name of the field
			ProductsWhoseNodeIsExpanded: []*Product{{Name: "ProductsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ProductCompositionShape with the name of the field
			ProductComposition_Shapes: []*ProductCompositionShape{{Name: "ProductComposition_Shapes"}},
			// field is initialized with an instance of TaskShape with the name of the field
			Task_Shapes: []*TaskShape{{Name: "Task_Shapes"}},
			// field is initialized with an instance of Task with the name of the field
			TasksWhoseNodeIsExpanded: []*Task{{Name: "TasksWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TasksWhoseInputNodeIsExpanded: []*Task{{Name: "TasksWhoseInputNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TasksWhoseOutputNodeIsExpanded: []*Task{{Name: "TasksWhoseOutputNodeIsExpanded"}},
			// field is initialized with an instance of TaskGroupShape with the name of the field
			TaskGroupShapes: []*TaskGroupShape{{Name: "TaskGroupShapes"}},
			// field is initialized with an instance of TaskGroup with the name of the field
			TaskGroupsWhoseNodeIsExpanded: []*TaskGroup{{Name: "TaskGroupsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of TaskCompositionShape with the name of the field
			TaskComposition_Shapes: []*TaskCompositionShape{{Name: "TaskComposition_Shapes"}},
			// field is initialized with an instance of TaskInputShape with the name of the field
			TaskInputShapes: []*TaskInputShape{{Name: "TaskInputShapes"}},
			// field is initialized with an instance of TaskOutputShape with the name of the field
			TaskOutputShapes: []*TaskOutputShape{{Name: "TaskOutputShapes"}},
			// field is initialized with an instance of NoteShape with the name of the field
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of NoteProductShape with the name of the field
			NoteProductShapes: []*NoteProductShape{{Name: "NoteProductShapes"}},
			// field is initialized with an instance of NoteTaskShape with the name of the field
			NoteTaskShapes: []*NoteTaskShape{{Name: "NoteTaskShapes"}},
			// field is initialized with an instance of NoteResourceShape with the name of the field
			NoteResourceShapes: []*NoteResourceShape{{Name: "NoteResourceShapes"}},
			// field is initialized with an instance of ResourceShape with the name of the field
			Resource_Shapes: []*ResourceShape{{Name: "Resource_Shapes"}},
			// field is initialized with an instance of Resource with the name of the field
			ResourcesWhoseNodeIsExpanded: []*Resource{{Name: "ResourcesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ResourceCompositionShape with the name of the field
			ResourceComposition_Shapes: []*ResourceCompositionShape{{Name: "ResourceComposition_Shapes"}},
			// field is initialized with an instance of ResourceTaskShape with the name of the field
			ResourceTaskShapes: []*ResourceTaskShape{{Name: "ResourceTaskShapes"}},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Product with the name of the field
			RootProducts: []*Product{{Name: "RootProducts"}},
			// field is initialized with an instance of Task with the name of the field
			RootTasks: []*Task{{Name: "RootTasks"}},
			// field is initialized with an instance of TaskGroup with the name of the field
			RootTaskGroups: []*TaskGroup{{Name: "RootTaskGroups"}},
			// field is initialized with an instance of Resource with the name of the field
			RootResources: []*Resource{{Name: "RootResources"}},
			// field is initialized with an instance of Note with the name of the field
			Notes: []*Note{{Name: "Notes"}},
			// field is initialized with an instance of Diagram with the name of the field
			Diagrams: []*Diagram{{Name: "Diagrams"}},
		}).(*Type)
	case Note:
		return any(&Note{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			Products: []*Product{{Name: "Products"}},
			// field is initialized with an instance of Task with the name of the field
			Tasks: []*Task{{Name: "Tasks"}},
			// field is initialized with an instance of Resource with the name of the field
			Resources: []*Resource{{Name: "Resources"}},
		}).(*Type)
	case NoteProductShape:
		return any(&NoteProductShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Product with the name of the field
			Product: &Product{Name: "Product"},
		}).(*Type)
	case NoteResourceShape:
		return any(&NoteResourceShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Resource with the name of the field
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
		}).(*Type)
	case NoteTaskShape:
		return any(&NoteTaskShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
		}).(*Type)
	case Product:
		return any(&Product{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			ReferencedProduct: &Product{Name: "ReferencedProduct"},
			// field is initialized with an instance of Product with the name of the field
			SubProducts: []*Product{{Name: "SubProducts"}},
		}).(*Type)
	case ProductCompositionShape:
		return any(&ProductCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			Product: &Product{Name: "Product"},
		}).(*Type)
	case ProductShape:
		return any(&ProductShape{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			Product: &Product{Name: "Product"},
		}).(*Type)
	case Resource:
		return any(&Resource{
			// Initialisation of associations
			// field is initialized with an instance of Resource with the name of the field
			ReferencedResource: &Resource{Name: "ReferencedResource"},
			// field is initialized with an instance of Task with the name of the field
			Tasks: []*Task{{Name: "Tasks"}},
			// field is initialized with an instance of Resource with the name of the field
			SubResources: []*Resource{{Name: "SubResources"}},
		}).(*Type)
	case ResourceCompositionShape:
		return any(&ResourceCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Resource with the name of the field
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
	case ResourceShape:
		return any(&ResourceShape{
			// Initialisation of associations
			// field is initialized with an instance of Resource with the name of the field
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
	case ResourceTaskShape:
		return any(&ResourceTaskShape{
			// Initialisation of associations
			// field is initialized with an instance of Resource with the name of the field
			Resource: &Resource{Name: "Resource"},
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
		}).(*Type)
	case Task:
		return any(&Task{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			ReferencedTask: &Task{Name: "ReferencedTask"},
			// field is initialized with an instance of Task with the name of the field
			Predecessors: []*Task{{Name: "Predecessors"}},
			// field is initialized with an instance of Task with the name of the field
			SubTasks: []*Task{{Name: "SubTasks"}},
			// field is initialized with an instance of Product with the name of the field
			Inputs: []*Product{{Name: "Inputs"}},
			// field is initialized with an instance of Product with the name of the field
			Outputs: []*Product{{Name: "Outputs"}},
			// field is initialized with an instance of TaskGroup with the name of the field
			TaskGroupsToDisplay: []*TaskGroup{{Name: "TaskGroupsToDisplay"}},
		}).(*Type)
	case TaskCompositionShape:
		return any(&TaskCompositionShape{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
		}).(*Type)
	case TaskGroup:
		return any(&TaskGroup{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Tasks: []*Task{{Name: "Tasks"}},
		}).(*Type)
	case TaskGroupShape:
		return any(&TaskGroupShape{
			// Initialisation of associations
			// field is initialized with an instance of TaskGroup with the name of the field
			TaskGroup: &TaskGroup{Name: "TaskGroup"},
		}).(*Type)
	case TaskInputShape:
		return any(&TaskInputShape{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			Product: &Product{Name: "Product"},
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
		}).(*Type)
	case TaskOutputShape:
		return any(&TaskOutputShape{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
			// field is initialized with an instance of Product with the name of the field
			Product: &Product{Name: "Product"},
		}).(*Type)
	case TaskShape:
		return any(&TaskShape{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
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
			res := make(map[*Product][]*NoteProductShape)
			for noteproductshape := range stage.NoteProductShapes {
				if noteproductshape.Product != nil {
					product_ := noteproductshape.Product
					var noteproductshapes []*NoteProductShape
					_, ok := res[product_]
					if ok {
						noteproductshapes = res[product_]
					} else {
						noteproductshapes = make([]*NoteProductShape, 0)
					}
					noteproductshapes = append(noteproductshapes, noteproductshape)
					res[product_] = noteproductshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteResourceShape
	case NoteResourceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteResourceShape)
			for noteresourceshape := range stage.NoteResourceShapes {
				if noteresourceshape.Note != nil {
					note_ := noteresourceshape.Note
					var noteresourceshapes []*NoteResourceShape
					_, ok := res[note_]
					if ok {
						noteresourceshapes = res[note_]
					} else {
						noteresourceshapes = make([]*NoteResourceShape, 0)
					}
					noteresourceshapes = append(noteresourceshapes, noteresourceshape)
					res[note_] = noteresourceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Resource":
			res := make(map[*Resource][]*NoteResourceShape)
			for noteresourceshape := range stage.NoteResourceShapes {
				if noteresourceshape.Resource != nil {
					resource_ := noteresourceshape.Resource
					var noteresourceshapes []*NoteResourceShape
					_, ok := res[resource_]
					if ok {
						noteresourceshapes = res[resource_]
					} else {
						noteresourceshapes = make([]*NoteResourceShape, 0)
					}
					noteresourceshapes = append(noteresourceshapes, noteresourceshape)
					res[resource_] = noteresourceshapes
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
			res := make(map[*Task][]*NoteTaskShape)
			for notetaskshape := range stage.NoteTaskShapes {
				if notetaskshape.Task != nil {
					task_ := notetaskshape.Task
					var notetaskshapes []*NoteTaskShape
					_, ok := res[task_]
					if ok {
						notetaskshapes = res[task_]
					} else {
						notetaskshapes = make([]*NoteTaskShape, 0)
					}
					notetaskshapes = append(notetaskshapes, notetaskshape)
					res[task_] = notetaskshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Product
	case Product:
		switch fieldname {
		// insertion point for per direct association field
		case "ReferencedProduct":
			res := make(map[*Product][]*Product)
			for product := range stage.Products {
				if product.ReferencedProduct != nil {
					product_ := product.ReferencedProduct
					var products []*Product
					_, ok := res[product_]
					if ok {
						products = res[product_]
					} else {
						products = make([]*Product, 0)
					}
					products = append(products, product)
					res[product_] = products
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProductCompositionShape
	case ProductCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Product":
			res := make(map[*Product][]*ProductCompositionShape)
			for productcompositionshape := range stage.ProductCompositionShapes {
				if productcompositionshape.Product != nil {
					product_ := productcompositionshape.Product
					var productcompositionshapes []*ProductCompositionShape
					_, ok := res[product_]
					if ok {
						productcompositionshapes = res[product_]
					} else {
						productcompositionshapes = make([]*ProductCompositionShape, 0)
					}
					productcompositionshapes = append(productcompositionshapes, productcompositionshape)
					res[product_] = productcompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProductShape
	case ProductShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Product":
			res := make(map[*Product][]*ProductShape)
			for productshape := range stage.ProductShapes {
				if productshape.Product != nil {
					product_ := productshape.Product
					var productshapes []*ProductShape
					_, ok := res[product_]
					if ok {
						productshapes = res[product_]
					} else {
						productshapes = make([]*ProductShape, 0)
					}
					productshapes = append(productshapes, productshape)
					res[product_] = productshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Resource
	case Resource:
		switch fieldname {
		// insertion point for per direct association field
		case "ReferencedResource":
			res := make(map[*Resource][]*Resource)
			for resource := range stage.Resources {
				if resource.ReferencedResource != nil {
					resource_ := resource.ReferencedResource
					var resources []*Resource
					_, ok := res[resource_]
					if ok {
						resources = res[resource_]
					} else {
						resources = make([]*Resource, 0)
					}
					resources = append(resources, resource)
					res[resource_] = resources
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ResourceCompositionShape
	case ResourceCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Resource":
			res := make(map[*Resource][]*ResourceCompositionShape)
			for resourcecompositionshape := range stage.ResourceCompositionShapes {
				if resourcecompositionshape.Resource != nil {
					resource_ := resourcecompositionshape.Resource
					var resourcecompositionshapes []*ResourceCompositionShape
					_, ok := res[resource_]
					if ok {
						resourcecompositionshapes = res[resource_]
					} else {
						resourcecompositionshapes = make([]*ResourceCompositionShape, 0)
					}
					resourcecompositionshapes = append(resourcecompositionshapes, resourcecompositionshape)
					res[resource_] = resourcecompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ResourceShape
	case ResourceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Resource":
			res := make(map[*Resource][]*ResourceShape)
			for resourceshape := range stage.ResourceShapes {
				if resourceshape.Resource != nil {
					resource_ := resourceshape.Resource
					var resourceshapes []*ResourceShape
					_, ok := res[resource_]
					if ok {
						resourceshapes = res[resource_]
					} else {
						resourceshapes = make([]*ResourceShape, 0)
					}
					resourceshapes = append(resourceshapes, resourceshape)
					res[resource_] = resourceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ResourceTaskShape
	case ResourceTaskShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Resource":
			res := make(map[*Resource][]*ResourceTaskShape)
			for resourcetaskshape := range stage.ResourceTaskShapes {
				if resourcetaskshape.Resource != nil {
					resource_ := resourcetaskshape.Resource
					var resourcetaskshapes []*ResourceTaskShape
					_, ok := res[resource_]
					if ok {
						resourcetaskshapes = res[resource_]
					} else {
						resourcetaskshapes = make([]*ResourceTaskShape, 0)
					}
					resourcetaskshapes = append(resourcetaskshapes, resourcetaskshape)
					res[resource_] = resourcetaskshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Task":
			res := make(map[*Task][]*ResourceTaskShape)
			for resourcetaskshape := range stage.ResourceTaskShapes {
				if resourcetaskshape.Task != nil {
					task_ := resourcetaskshape.Task
					var resourcetaskshapes []*ResourceTaskShape
					_, ok := res[task_]
					if ok {
						resourcetaskshapes = res[task_]
					} else {
						resourcetaskshapes = make([]*ResourceTaskShape, 0)
					}
					resourcetaskshapes = append(resourcetaskshapes, resourcetaskshape)
					res[task_] = resourcetaskshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
		case "ReferencedTask":
			res := make(map[*Task][]*Task)
			for task := range stage.Tasks {
				if task.ReferencedTask != nil {
					task_ := task.ReferencedTask
					var tasks []*Task
					_, ok := res[task_]
					if ok {
						tasks = res[task_]
					} else {
						tasks = make([]*Task, 0)
					}
					tasks = append(tasks, task)
					res[task_] = tasks
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskCompositionShape
	case TaskCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Task":
			res := make(map[*Task][]*TaskCompositionShape)
			for taskcompositionshape := range stage.TaskCompositionShapes {
				if taskcompositionshape.Task != nil {
					task_ := taskcompositionshape.Task
					var taskcompositionshapes []*TaskCompositionShape
					_, ok := res[task_]
					if ok {
						taskcompositionshapes = res[task_]
					} else {
						taskcompositionshapes = make([]*TaskCompositionShape, 0)
					}
					taskcompositionshapes = append(taskcompositionshapes, taskcompositionshape)
					res[task_] = taskcompositionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskGroup
	case TaskGroup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TaskGroupShape
	case TaskGroupShape:
		switch fieldname {
		// insertion point for per direct association field
		case "TaskGroup":
			res := make(map[*TaskGroup][]*TaskGroupShape)
			for taskgroupshape := range stage.TaskGroupShapes {
				if taskgroupshape.TaskGroup != nil {
					taskgroup_ := taskgroupshape.TaskGroup
					var taskgroupshapes []*TaskGroupShape
					_, ok := res[taskgroup_]
					if ok {
						taskgroupshapes = res[taskgroup_]
					} else {
						taskgroupshapes = make([]*TaskGroupShape, 0)
					}
					taskgroupshapes = append(taskgroupshapes, taskgroupshape)
					res[taskgroup_] = taskgroupshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskInputShape
	case TaskInputShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Product":
			res := make(map[*Product][]*TaskInputShape)
			for taskinputshape := range stage.TaskInputShapes {
				if taskinputshape.Product != nil {
					product_ := taskinputshape.Product
					var taskinputshapes []*TaskInputShape
					_, ok := res[product_]
					if ok {
						taskinputshapes = res[product_]
					} else {
						taskinputshapes = make([]*TaskInputShape, 0)
					}
					taskinputshapes = append(taskinputshapes, taskinputshape)
					res[product_] = taskinputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Task":
			res := make(map[*Task][]*TaskInputShape)
			for taskinputshape := range stage.TaskInputShapes {
				if taskinputshape.Task != nil {
					task_ := taskinputshape.Task
					var taskinputshapes []*TaskInputShape
					_, ok := res[task_]
					if ok {
						taskinputshapes = res[task_]
					} else {
						taskinputshapes = make([]*TaskInputShape, 0)
					}
					taskinputshapes = append(taskinputshapes, taskinputshape)
					res[task_] = taskinputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskOutputShape
	case TaskOutputShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Task":
			res := make(map[*Task][]*TaskOutputShape)
			for taskoutputshape := range stage.TaskOutputShapes {
				if taskoutputshape.Task != nil {
					task_ := taskoutputshape.Task
					var taskoutputshapes []*TaskOutputShape
					_, ok := res[task_]
					if ok {
						taskoutputshapes = res[task_]
					} else {
						taskoutputshapes = make([]*TaskOutputShape, 0)
					}
					taskoutputshapes = append(taskoutputshapes, taskoutputshape)
					res[task_] = taskoutputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Product":
			res := make(map[*Product][]*TaskOutputShape)
			for taskoutputshape := range stage.TaskOutputShapes {
				if taskoutputshape.Product != nil {
					product_ := taskoutputshape.Product
					var taskoutputshapes []*TaskOutputShape
					_, ok := res[product_]
					if ok {
						taskoutputshapes = res[product_]
					} else {
						taskoutputshapes = make([]*TaskOutputShape, 0)
					}
					taskoutputshapes = append(taskoutputshapes, taskoutputshape)
					res[product_] = taskoutputshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskShape
	case TaskShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Task":
			res := make(map[*Task][]*TaskShape)
			for taskshape := range stage.TaskShapes {
				if taskshape.Task != nil {
					task_ := taskshape.Task
					var taskshapes []*TaskShape
					_, ok := res[task_]
					if ok {
						taskshapes = res[task_]
					} else {
						taskshapes = make([]*TaskShape, 0)
					}
					taskshapes = append(taskshapes, taskshape)
					res[task_] = taskshapes
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		case "Product_Shapes":
			res := make(map[*ProductShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, productshape_ := range diagram.Product_Shapes {
					res[productshape_] = append(res[productshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProductsWhoseNodeIsExpanded":
			res := make(map[*Product][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, product_ := range diagram.ProductsWhoseNodeIsExpanded {
					res[product_] = append(res[product_], diagram)
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
		case "Task_Shapes":
			res := make(map[*TaskShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskshape_ := range diagram.Task_Shapes {
					res[taskshape_] = append(res[taskshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TasksWhoseNodeIsExpanded":
			res := make(map[*Task][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, task_ := range diagram.TasksWhoseNodeIsExpanded {
					res[task_] = append(res[task_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TasksWhoseInputNodeIsExpanded":
			res := make(map[*Task][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, task_ := range diagram.TasksWhoseInputNodeIsExpanded {
					res[task_] = append(res[task_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TasksWhoseOutputNodeIsExpanded":
			res := make(map[*Task][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, task_ := range diagram.TasksWhoseOutputNodeIsExpanded {
					res[task_] = append(res[task_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskGroupShapes":
			res := make(map[*TaskGroupShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskgroupshape_ := range diagram.TaskGroupShapes {
					res[taskgroupshape_] = append(res[taskgroupshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskGroupsWhoseNodeIsExpanded":
			res := make(map[*TaskGroup][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskgroup_ := range diagram.TaskGroupsWhoseNodeIsExpanded {
					res[taskgroup_] = append(res[taskgroup_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskComposition_Shapes":
			res := make(map[*TaskCompositionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskcompositionshape_ := range diagram.TaskComposition_Shapes {
					res[taskcompositionshape_] = append(res[taskcompositionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskInputShapes":
			res := make(map[*TaskInputShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskinputshape_ := range diagram.TaskInputShapes {
					res[taskinputshape_] = append(res[taskinputshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskOutputShapes":
			res := make(map[*TaskOutputShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskoutputshape_ := range diagram.TaskOutputShapes {
					res[taskoutputshape_] = append(res[taskoutputshape_], diagram)
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
			res := make(map[*NoteResourceShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, noteresourceshape_ := range diagram.NoteResourceShapes {
					res[noteresourceshape_] = append(res[noteresourceshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Resource_Shapes":
			res := make(map[*ResourceShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, resourceshape_ := range diagram.Resource_Shapes {
					res[resourceshape_] = append(res[resourceshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ResourcesWhoseNodeIsExpanded":
			res := make(map[*Resource][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, resource_ := range diagram.ResourcesWhoseNodeIsExpanded {
					res[resource_] = append(res[resource_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ResourceComposition_Shapes":
			res := make(map[*ResourceCompositionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, resourcecompositionshape_ := range diagram.ResourceComposition_Shapes {
					res[resourcecompositionshape_] = append(res[resourcecompositionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ResourceTaskShapes":
			res := make(map[*ResourceTaskShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, resourcetaskshape_ := range diagram.ResourceTaskShapes {
					res[resourcetaskshape_] = append(res[resourcetaskshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		case "SubLibraries":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibraries {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootProducts":
			res := make(map[*Product][]*Library)
			for library := range stage.Librarys {
				for _, product_ := range library.RootProducts {
					res[product_] = append(res[product_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootTasks":
			res := make(map[*Task][]*Library)
			for library := range stage.Librarys {
				for _, task_ := range library.RootTasks {
					res[task_] = append(res[task_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootTaskGroups":
			res := make(map[*TaskGroup][]*Library)
			for library := range stage.Librarys {
				for _, taskgroup_ := range library.RootTaskGroups {
					res[taskgroup_] = append(res[taskgroup_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootResources":
			res := make(map[*Resource][]*Library)
			for library := range stage.Librarys {
				for _, resource_ := range library.RootResources {
					res[resource_] = append(res[resource_], library)
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
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "Products":
			res := make(map[*Product][]*Note)
			for note := range stage.Notes {
				for _, product_ := range note.Products {
					res[product_] = append(res[product_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tasks":
			res := make(map[*Task][]*Note)
			for note := range stage.Notes {
				for _, task_ := range note.Tasks {
					res[task_] = append(res[task_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Resources":
			res := make(map[*Resource][]*Note)
			for note := range stage.Notes {
				for _, resource_ := range note.Resources {
					res[resource_] = append(res[resource_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteProductShape
	case NoteProductShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteResourceShape
	case NoteResourceShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteTaskShape
	case NoteTaskShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Product
	case Product:
		switch fieldname {
		// insertion point for per direct association field
		case "SubProducts":
			res := make(map[*Product][]*Product)
			for product := range stage.Products {
				for _, product_ := range product.SubProducts {
					res[product_] = append(res[product_], product)
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of Resource
	case Resource:
		switch fieldname {
		// insertion point for per direct association field
		case "Tasks":
			res := make(map[*Task][]*Resource)
			for resource := range stage.Resources {
				for _, task_ := range resource.Tasks {
					res[task_] = append(res[task_], resource)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubResources":
			res := make(map[*Resource][]*Resource)
			for resource := range stage.Resources {
				for _, resource_ := range resource.SubResources {
					res[resource_] = append(res[resource_], resource)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ResourceCompositionShape
	case ResourceCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ResourceShape
	case ResourceShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ResourceTaskShape
	case ResourceTaskShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
		case "Predecessors":
			res := make(map[*Task][]*Task)
			for task := range stage.Tasks {
				for _, task_ := range task.Predecessors {
					res[task_] = append(res[task_], task)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubTasks":
			res := make(map[*Task][]*Task)
			for task := range stage.Tasks {
				for _, task_ := range task.SubTasks {
					res[task_] = append(res[task_], task)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Inputs":
			res := make(map[*Product][]*Task)
			for task := range stage.Tasks {
				for _, product_ := range task.Inputs {
					res[product_] = append(res[product_], task)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Outputs":
			res := make(map[*Product][]*Task)
			for task := range stage.Tasks {
				for _, product_ := range task.Outputs {
					res[product_] = append(res[product_], task)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskGroupsToDisplay":
			res := make(map[*TaskGroup][]*Task)
			for task := range stage.Tasks {
				for _, taskgroup_ := range task.TaskGroupsToDisplay {
					res[taskgroup_] = append(res[taskgroup_], task)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskCompositionShape
	case TaskCompositionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TaskGroup
	case TaskGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "Tasks":
			res := make(map[*Task][]*TaskGroup)
			for taskgroup := range stage.TaskGroups {
				for _, task_ := range taskgroup.Tasks {
					res[task_] = append(res[task_], taskgroup)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TaskGroupShape
	case TaskGroupShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TaskInputShape
	case TaskInputShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TaskOutputShape
	case TaskOutputShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TaskShape
	case TaskShape:
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
	case *Diagram:
		res = "Diagram"
	case *Library:
		res = "Library"
	case *Note:
		res = "Note"
	case *NoteProductShape:
		res = "NoteProductShape"
	case *NoteResourceShape:
		res = "NoteResourceShape"
	case *NoteShape:
		res = "NoteShape"
	case *NoteTaskShape:
		res = "NoteTaskShape"
	case *Product:
		res = "Product"
	case *ProductCompositionShape:
		res = "ProductCompositionShape"
	case *ProductShape:
		res = "ProductShape"
	case *Resource:
		res = "Resource"
	case *ResourceCompositionShape:
		res = "ResourceCompositionShape"
	case *ResourceShape:
		res = "ResourceShape"
	case *ResourceTaskShape:
		res = "ResourceTaskShape"
	case *Task:
		res = "Task"
	case *TaskCompositionShape:
		res = "TaskCompositionShape"
	case *TaskGroup:
		res = "TaskGroup"
	case *TaskGroupShape:
		res = "TaskGroupShape"
	case *TaskInputShape:
		res = "TaskInputShape"
	case *TaskOutputShape:
		res = "TaskOutputShape"
	case *TaskShape:
		res = "TaskShape"
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
	case *NoteResourceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NoteResourceShapes"
		res = append(res, rf)
	case *NoteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Note_Shapes"
		res = append(res, rf)
	case *NoteTaskShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NoteTaskShapes"
		res = append(res, rf)
	case *Product:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ProductsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootProducts"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Products"
		res = append(res, rf)
		rf.GongstructName = "Product"
		rf.Fieldname = "SubProducts"
		res = append(res, rf)
		rf.GongstructName = "Task"
		rf.Fieldname = "Inputs"
		res = append(res, rf)
		rf.GongstructName = "Task"
		rf.Fieldname = "Outputs"
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
	case *Resource:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ResourcesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootResources"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Resources"
		res = append(res, rf)
		rf.GongstructName = "Resource"
		rf.Fieldname = "SubResources"
		res = append(res, rf)
	case *ResourceCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ResourceComposition_Shapes"
		res = append(res, rf)
	case *ResourceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Resource_Shapes"
		res = append(res, rf)
	case *ResourceTaskShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ResourceTaskShapes"
		res = append(res, rf)
	case *Task:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TasksWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TasksWhoseInputNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TasksWhoseOutputNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootTasks"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Tasks"
		res = append(res, rf)
		rf.GongstructName = "Resource"
		rf.Fieldname = "Tasks"
		res = append(res, rf)
		rf.GongstructName = "Task"
		rf.Fieldname = "Predecessors"
		res = append(res, rf)
		rf.GongstructName = "Task"
		rf.Fieldname = "SubTasks"
		res = append(res, rf)
		rf.GongstructName = "TaskGroup"
		rf.Fieldname = "Tasks"
		res = append(res, rf)
	case *TaskCompositionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TaskComposition_Shapes"
		res = append(res, rf)
	case *TaskGroup:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TaskGroupsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootTaskGroups"
		res = append(res, rf)
		rf.GongstructName = "Task"
		rf.Fieldname = "TaskGroupsToDisplay"
		res = append(res, rf)
	case *TaskGroupShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TaskGroupShapes"
		res = append(res, rf)
	case *TaskInputShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TaskInputShapes"
		res = append(res, rf)
	case *TaskOutputShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TaskOutputShapes"
		res = append(res, rf)
	case *TaskShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Task_Shapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
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
			Name:               "IsShowPrefix",
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
			Name:                 "Product_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProductShape",
		},
		{
			Name:                 "ProductsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
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
			Name:               "IsWBSNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Task_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskShape",
		},
		{
			Name:                 "TasksWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "TasksWhoseInputNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "TasksWhoseOutputNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:               "IsTaskGroupsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "TaskGroupShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskGroupShape",
		},
		{
			Name:                 "TaskGroupsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskGroup",
		},
		{
			Name:               "DateFormat",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TaskComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskCompositionShape",
		},
		{
			Name:                 "TaskInputShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskInputShape",
		},
		{
			Name:                 "TaskOutputShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskOutputShape",
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
			TargetGongstructName: "NoteResourceShape",
		},
		{
			Name:                 "Resource_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ResourceShape",
		},
		{
			Name:                 "ResourcesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
		{
			Name:               "IsResourcesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ResourceComposition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ResourceCompositionShape",
		},
		{
			Name:                 "ResourceTaskShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ResourceTaskShape",
		},
		{
			Name:               "IsTimeDiagram",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ComputedStart",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "ComputedEnd",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "ComputedDuration",
			GongFieldValueType: GongFieldValueTypeIntDuration,
		},
		{
			Name:               "UseManualStartAndEndDates",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ManualStart",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "ManualEnd",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "TimeStep",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "TimeStepScale",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TimeStepScaleEnum",
		},
		{
			Name:               "LaneHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "RatioBarToLaneHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "YTopMargin",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "XLeftText",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TextHeight",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "XLeftLanes",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "XRightMargin",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ArrowLengthToTheRightOfStartBar",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ArrowTipLenght",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TimeLine_Color",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeLine_FillOpacity",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "TimeLine_Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "TimeLine_StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DrawVerticalTimeLines",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Group_Stroke",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Group_StrokeWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Group_StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DateYOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "AlignOnStartEndOnYearStart",
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
			Name:                 "SubLibraries",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "NbPixPerCharacter",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "LogoSVGFile",
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
			Name:               "IsRootLibrary",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "RootProducts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:                 "RootTasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "RootTaskGroups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskGroup",
		},
		{
			Name:                 "RootResources",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
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
			TargetGongstructName: "Product",
		},
		{
			Name:                 "Tasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "Resources",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
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
			TargetGongstructName: "Product",
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

func (noteresourceshape *NoteResourceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Resource",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Resource",
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
			Name:               "OverideLayoutDirection",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
			TargetGongstructName: "Task",
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

func (product *Product) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "IsImport",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ReferencedProduct",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Product",
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SubProducts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsProducersNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsConsumersNodeExpanded",
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
			TargetGongstructName: "Product",
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
			TargetGongstructName: "Product",
		},
		{
			Name:               "OverideLayoutDirection",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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

func (resource *Resource) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "IsImport",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ReferencedResource",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Resource",
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Tasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "SubResources",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
	}
	return
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Resource",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Resource",
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

func (resourceshape *ResourceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Resource",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Resource",
		},
		{
			Name:               "OverideLayoutDirection",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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

func (resourcetaskshape *ResourceTaskShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Resource",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Resource",
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
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

func (task *Task) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "IsImport",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ReferencedTask",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:               "Start",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "End",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeIntDuration,
		},
		{
			Name:               "IsEndDateComputedFromDuration",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Predecessors",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:               "IsStartDateComputedFromPredecessors",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsMilestone",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SubTasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "Inputs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsInputsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Outputs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
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
			Name:               "DisplayVerticalBar",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "TaskGroupsToDisplay",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskGroup",
		},
		{
			Name:                 "TextPosition",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "TextPositionEnum",
		},
		{
			Name:               "XOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "YOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (taskcompositionshape *TaskCompositionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
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

func (taskgroup *TaskGroup) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Tasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
	}
	return
}

func (taskgroupshape *TaskGroupShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "TaskGroup",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "TaskGroup",
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

func (taskinputshape *TaskInputShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Product",
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
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

func (taskoutputshape *TaskOutputShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "Product",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Product",
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

func (taskshape *TaskShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Task",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:               "IsShowDate",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "OverideLayoutDirection",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "LayoutDirection",
			GongFieldValueType:   GongFieldValueTypeInt,
			TargetGongstructName: "LayoutDirection",
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
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagram.IsShowPrefix)
		res.valueBool = diagram.IsShowPrefix
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
	case "IsWBSNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsWBSNodeExpanded)
		res.valueBool = diagram.IsWBSNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Task_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Task_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TasksWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TasksWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TasksWhoseInputNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TasksWhoseInputNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TasksWhoseOutputNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TasksWhoseOutputNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsTaskGroupsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsTaskGroupsNodeExpanded)
		res.valueBool = diagram.IsTaskGroupsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "TaskGroupShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TaskGroupShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TaskGroupsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TaskGroupsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DateFormat":
		res.valueString = diagram.DateFormat
	case "TaskComposition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TaskComposition_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TaskInputShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TaskInputShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TaskOutputShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TaskOutputShapes {
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
	case "Resource_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Resource_Shapes {
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
	case "IsResourcesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsResourcesNodeExpanded)
		res.valueBool = diagram.IsResourcesNodeExpanded
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
	case "ResourceTaskShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ResourceTaskShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsTimeDiagram":
		res.valueString = fmt.Sprintf("%t", diagram.IsTimeDiagram)
		res.valueBool = diagram.IsTimeDiagram
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedStart":
		res.valueString = diagram.ComputedStart.String()
	case "ComputedEnd":
		res.valueString = diagram.ComputedEnd.String()
	case "ComputedDuration":
		if math.Abs(diagram.ComputedDuration.Hours()) >= 24 {
			days := __Gong__Abs(int(int(diagram.ComputedDuration.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(diagram.ComputedDuration.Hours()) % 24
			remainingMinutes := int(diagram.ComputedDuration.Minutes()) % 60
			remainingSeconds := int(diagram.ComputedDuration.Seconds()) % 60

			if diagram.ComputedDuration.Hours() < 0 {
				res.valueString = "- "
			}

			if months > 0 {
				if months > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d months", months)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d month", months)
				}
			}
			if days > 0 {
				if months != 0 {
					res.valueString = res.valueString + ", "
				}
				if days > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d days", days)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d day", days)
				}

			}
			if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
				if days != 0 || (days == 0 && months != 0) {
					res.valueString = res.valueString + ", "
				}
				res.valueString = res.valueString + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
			}
		} else {
			res.valueString = fmt.Sprintf("%s\n", diagram.ComputedDuration.String())
		}
	case "UseManualStartAndEndDates":
		res.valueString = fmt.Sprintf("%t", diagram.UseManualStartAndEndDates)
		res.valueBool = diagram.UseManualStartAndEndDates
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ManualStart":
		res.valueString = diagram.ManualStart.String()
	case "ManualEnd":
		res.valueString = diagram.ManualEnd.String()
	case "TimeStep":
		res.valueString = fmt.Sprintf("%d", diagram.TimeStep)
		res.valueInt = diagram.TimeStep
		res.GongFieldValueType = GongFieldValueTypeInt
	case "TimeStepScale":
		enum := diagram.TimeStepScale
		res.valueString = enum.ToCodeString()
	case "LaneHeight":
		res.valueString = fmt.Sprintf("%f", diagram.LaneHeight)
		res.valueFloat = diagram.LaneHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "RatioBarToLaneHeight":
		res.valueString = fmt.Sprintf("%f", diagram.RatioBarToLaneHeight)
		res.valueFloat = diagram.RatioBarToLaneHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "YTopMargin":
		res.valueString = fmt.Sprintf("%f", diagram.YTopMargin)
		res.valueFloat = diagram.YTopMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XLeftText":
		res.valueString = fmt.Sprintf("%f", diagram.XLeftText)
		res.valueFloat = diagram.XLeftText
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TextHeight":
		res.valueString = fmt.Sprintf("%f", diagram.TextHeight)
		res.valueFloat = diagram.TextHeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XLeftLanes":
		res.valueString = fmt.Sprintf("%f", diagram.XLeftLanes)
		res.valueFloat = diagram.XLeftLanes
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "XRightMargin":
		res.valueString = fmt.Sprintf("%f", diagram.XRightMargin)
		res.valueFloat = diagram.XRightMargin
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ArrowLengthToTheRightOfStartBar":
		res.valueString = fmt.Sprintf("%f", diagram.ArrowLengthToTheRightOfStartBar)
		res.valueFloat = diagram.ArrowLengthToTheRightOfStartBar
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ArrowTipLenght":
		res.valueString = fmt.Sprintf("%f", diagram.ArrowTipLenght)
		res.valueFloat = diagram.ArrowTipLenght
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TimeLine_Color":
		res.valueString = diagram.TimeLine_Color
	case "TimeLine_FillOpacity":
		res.valueString = fmt.Sprintf("%f", diagram.TimeLine_FillOpacity)
		res.valueFloat = diagram.TimeLine_FillOpacity
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "TimeLine_Stroke":
		res.valueString = diagram.TimeLine_Stroke
	case "TimeLine_StrokeWidth":
		res.valueString = fmt.Sprintf("%f", diagram.TimeLine_StrokeWidth)
		res.valueFloat = diagram.TimeLine_StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DrawVerticalTimeLines":
		res.valueString = fmt.Sprintf("%t", diagram.DrawVerticalTimeLines)
		res.valueBool = diagram.DrawVerticalTimeLines
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Group_Stroke":
		res.valueString = diagram.Group_Stroke
	case "Group_StrokeWidth":
		res.valueString = fmt.Sprintf("%f", diagram.Group_StrokeWidth)
		res.valueFloat = diagram.Group_StrokeWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Group_StrokeDashArray":
		res.valueString = diagram.Group_StrokeDashArray
	case "DateYOffset":
		res.valueString = fmt.Sprintf("%f", diagram.DateYOffset)
		res.valueFloat = diagram.DateYOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "AlignOnStartEndOnYearStart":
		res.valueString = fmt.Sprintf("%t", diagram.AlignOnStartEndOnYearStart)
		res.valueBool = diagram.AlignOnStartEndOnYearStart
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (library *Library) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = library.Name
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
	case "LogoSVGFile":
		res.valueString = library.LogoSVGFile
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsExpanded)
		res.valueBool = library.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := library.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsRootLibrary":
		res.valueString = fmt.Sprintf("%t", library.IsRootLibrary)
		res.valueBool = library.IsRootLibrary
		res.GongFieldValueType = GongFieldValueTypeBool
	case "RootProducts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootProducts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootTasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootTasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootTaskGroups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootTaskGroups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootResources":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootResources {
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

func (noteresourceshape *NoteResourceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = noteresourceshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteresourceshape.Note != nil {
			res.valueString = noteresourceshape.Note.Name
			res.ids = noteresourceshape.Note.GongGetUUID(stage)
		}
	case "Resource":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteresourceshape.Resource != nil {
			res.valueString = noteresourceshape.Resource.Name
			res.ids = noteresourceshape.Resource.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", noteresourceshape.StartRatio)
		res.valueFloat = noteresourceshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", noteresourceshape.EndRatio)
		res.valueFloat = noteresourceshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := noteresourceshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := noteresourceshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", noteresourceshape.CornerOffsetRatio)
		res.valueFloat = noteresourceshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", noteresourceshape.IsHidden)
		res.valueBool = noteresourceshape.IsHidden
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
	case "OverideLayoutDirection":
		res.valueString = fmt.Sprintf("%t", noteshape.OverideLayoutDirection)
		res.valueBool = noteshape.OverideLayoutDirection
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := noteshape.LayoutDirection
		res.valueString = enum.ToCodeString()
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

func (product *Product) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = product.Name
	case "ComputedPrefix":
		res.valueString = product.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", product.IsExpanded)
		res.valueBool = product.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := product.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsImport":
		res.valueString = fmt.Sprintf("%t", product.IsImport)
		res.valueBool = product.IsImport
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ReferencedProduct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if product.ReferencedProduct != nil {
			res.valueString = product.ReferencedProduct.Name
			res.ids = product.ReferencedProduct.GongGetUUID(stage)
		}
	case "Description":
		res.valueString = product.Description
	case "SubProducts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range product.SubProducts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsProducersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", product.IsProducersNodeExpanded)
		res.valueBool = product.IsProducersNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsConsumersNodeExpanded":
		res.valueString = fmt.Sprintf("%t", product.IsConsumersNodeExpanded)
		res.valueBool = product.IsConsumersNodeExpanded
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
	case "OverideLayoutDirection":
		res.valueString = fmt.Sprintf("%t", productshape.OverideLayoutDirection)
		res.valueBool = productshape.OverideLayoutDirection
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := productshape.LayoutDirection
		res.valueString = enum.ToCodeString()
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

func (resource *Resource) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = resource.Name
	case "ComputedPrefix":
		res.valueString = resource.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", resource.IsExpanded)
		res.valueBool = resource.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := resource.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsImport":
		res.valueString = fmt.Sprintf("%t", resource.IsImport)
		res.valueBool = resource.IsImport
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ReferencedResource":
		res.GongFieldValueType = GongFieldValueTypePointer
		if resource.ReferencedResource != nil {
			res.valueString = resource.ReferencedResource.Name
			res.ids = resource.ReferencedResource.GongGetUUID(stage)
		}
	case "Description":
		res.valueString = resource.Description
	case "Tasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range resource.Tasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SubResources":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range resource.SubResources {
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

func (resourcecompositionshape *ResourceCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = resourcecompositionshape.Name
	case "Resource":
		res.GongFieldValueType = GongFieldValueTypePointer
		if resourcecompositionshape.Resource != nil {
			res.valueString = resourcecompositionshape.Resource.Name
			res.ids = resourcecompositionshape.Resource.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", resourcecompositionshape.StartRatio)
		res.valueFloat = resourcecompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", resourcecompositionshape.EndRatio)
		res.valueFloat = resourcecompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := resourcecompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := resourcecompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", resourcecompositionshape.CornerOffsetRatio)
		res.valueFloat = resourcecompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", resourcecompositionshape.IsHidden)
		res.valueBool = resourcecompositionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (resourceshape *ResourceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = resourceshape.Name
	case "Resource":
		res.GongFieldValueType = GongFieldValueTypePointer
		if resourceshape.Resource != nil {
			res.valueString = resourceshape.Resource.Name
			res.ids = resourceshape.Resource.GongGetUUID(stage)
		}
	case "OverideLayoutDirection":
		res.valueString = fmt.Sprintf("%t", resourceshape.OverideLayoutDirection)
		res.valueBool = resourceshape.OverideLayoutDirection
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := resourceshape.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "X":
		res.valueString = fmt.Sprintf("%f", resourceshape.X)
		res.valueFloat = resourceshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", resourceshape.Y)
		res.valueFloat = resourceshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", resourceshape.Width)
		res.valueFloat = resourceshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", resourceshape.Height)
		res.valueFloat = resourceshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", resourceshape.IsHidden)
		res.valueBool = resourceshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (resourcetaskshape *ResourceTaskShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = resourcetaskshape.Name
	case "Resource":
		res.GongFieldValueType = GongFieldValueTypePointer
		if resourcetaskshape.Resource != nil {
			res.valueString = resourcetaskshape.Resource.Name
			res.ids = resourcetaskshape.Resource.GongGetUUID(stage)
		}
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if resourcetaskshape.Task != nil {
			res.valueString = resourcetaskshape.Task.Name
			res.ids = resourcetaskshape.Task.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", resourcetaskshape.StartRatio)
		res.valueFloat = resourcetaskshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", resourcetaskshape.EndRatio)
		res.valueFloat = resourcetaskshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := resourcetaskshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := resourcetaskshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", resourcetaskshape.CornerOffsetRatio)
		res.valueFloat = resourcetaskshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", resourcetaskshape.IsHidden)
		res.valueBool = resourcetaskshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (task *Task) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = task.Name
	case "ComputedPrefix":
		res.valueString = task.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsExpanded)
		res.valueBool = task.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := task.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "IsImport":
		res.valueString = fmt.Sprintf("%t", task.IsImport)
		res.valueBool = task.IsImport
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ReferencedTask":
		res.GongFieldValueType = GongFieldValueTypePointer
		if task.ReferencedTask != nil {
			res.valueString = task.ReferencedTask.Name
			res.ids = task.ReferencedTask.GongGetUUID(stage)
		}
	case "Start":
		res.valueString = task.Start.String()
	case "End":
		res.valueString = task.End.String()
	case "Duration":
		if math.Abs(task.Duration.Hours()) >= 24 {
			days := __Gong__Abs(int(int(task.Duration.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(task.Duration.Hours()) % 24
			remainingMinutes := int(task.Duration.Minutes()) % 60
			remainingSeconds := int(task.Duration.Seconds()) % 60

			if task.Duration.Hours() < 0 {
				res.valueString = "- "
			}

			if months > 0 {
				if months > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d months", months)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d month", months)
				}
			}
			if days > 0 {
				if months != 0 {
					res.valueString = res.valueString + ", "
				}
				if days > 1 {
					res.valueString = res.valueString + fmt.Sprintf("%d days", days)
				} else {
					res.valueString = res.valueString + fmt.Sprintf("%d day", days)
				}

			}
			if remainingHours != 0 || remainingMinutes != 0 || remainingSeconds != 0 {
				if days != 0 || (days == 0 && months != 0) {
					res.valueString = res.valueString + ", "
				}
				res.valueString = res.valueString + fmt.Sprintf("%d hours, %d minutes, %d seconds\n", remainingHours, remainingMinutes, remainingSeconds)
			}
		} else {
			res.valueString = fmt.Sprintf("%s\n", task.Duration.String())
		}
	case "IsEndDateComputedFromDuration":
		res.valueString = fmt.Sprintf("%t", task.IsEndDateComputedFromDuration)
		res.valueBool = task.IsEndDateComputedFromDuration
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Predecessors":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range task.Predecessors {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsStartDateComputedFromPredecessors":
		res.valueString = fmt.Sprintf("%t", task.IsStartDateComputedFromPredecessors)
		res.valueBool = task.IsStartDateComputedFromPredecessors
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsMilestone":
		res.valueString = fmt.Sprintf("%t", task.IsMilestone)
		res.valueBool = task.IsMilestone
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Description":
		res.valueString = task.Description
	case "SubTasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range task.SubTasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Inputs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range task.Inputs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsInputsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsInputsNodeExpanded)
		res.valueBool = task.IsInputsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Outputs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range task.Outputs {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsOutputsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsOutputsNodeExpanded)
		res.valueBool = task.IsOutputsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsWithCompletion":
		res.valueString = fmt.Sprintf("%t", task.IsWithCompletion)
		res.valueBool = task.IsWithCompletion
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Completion":
		enum := task.Completion
		res.valueString = enum.ToCodeString()
	case "DisplayVerticalBar":
		res.valueString = fmt.Sprintf("%t", task.DisplayVerticalBar)
		res.valueBool = task.DisplayVerticalBar
		res.GongFieldValueType = GongFieldValueTypeBool
	case "TaskGroupsToDisplay":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range task.TaskGroupsToDisplay {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TextPosition":
		enum := task.TextPosition
		res.valueString = enum.ToCodeString()
	case "XOffset":
		res.valueString = fmt.Sprintf("%f", task.XOffset)
		res.valueFloat = task.XOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "YOffset":
		res.valueString = fmt.Sprintf("%f", task.YOffset)
		res.valueFloat = task.YOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (taskcompositionshape *TaskCompositionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskcompositionshape.Name
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskcompositionshape.Task != nil {
			res.valueString = taskcompositionshape.Task.Name
			res.ids = taskcompositionshape.Task.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", taskcompositionshape.StartRatio)
		res.valueFloat = taskcompositionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", taskcompositionshape.EndRatio)
		res.valueFloat = taskcompositionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := taskcompositionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := taskcompositionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", taskcompositionshape.CornerOffsetRatio)
		res.valueFloat = taskcompositionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", taskcompositionshape.IsHidden)
		res.valueBool = taskcompositionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (taskgroup *TaskGroup) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskgroup.Name
	case "ComputedPrefix":
		res.valueString = taskgroup.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", taskgroup.IsExpanded)
		res.valueBool = taskgroup.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := taskgroup.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "Tasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range taskgroup.Tasks {
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

func (taskgroupshape *TaskGroupShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskgroupshape.Name
	case "TaskGroup":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskgroupshape.TaskGroup != nil {
			res.valueString = taskgroupshape.TaskGroup.Name
			res.ids = taskgroupshape.TaskGroup.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", taskgroupshape.X)
		res.valueFloat = taskgroupshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", taskgroupshape.Y)
		res.valueFloat = taskgroupshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", taskgroupshape.Width)
		res.valueFloat = taskgroupshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", taskgroupshape.Height)
		res.valueFloat = taskgroupshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", taskgroupshape.IsHidden)
		res.valueBool = taskgroupshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (taskinputshape *TaskInputShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskinputshape.Name
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskinputshape.Product != nil {
			res.valueString = taskinputshape.Product.Name
			res.ids = taskinputshape.Product.GongGetUUID(stage)
		}
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskinputshape.Task != nil {
			res.valueString = taskinputshape.Task.Name
			res.ids = taskinputshape.Task.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", taskinputshape.StartRatio)
		res.valueFloat = taskinputshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", taskinputshape.EndRatio)
		res.valueFloat = taskinputshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := taskinputshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := taskinputshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", taskinputshape.CornerOffsetRatio)
		res.valueFloat = taskinputshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", taskinputshape.IsHidden)
		res.valueBool = taskinputshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (taskoutputshape *TaskOutputShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskoutputshape.Name
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskoutputshape.Task != nil {
			res.valueString = taskoutputshape.Task.Name
			res.ids = taskoutputshape.Task.GongGetUUID(stage)
		}
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskoutputshape.Product != nil {
			res.valueString = taskoutputshape.Product.Name
			res.ids = taskoutputshape.Product.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", taskoutputshape.StartRatio)
		res.valueFloat = taskoutputshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", taskoutputshape.EndRatio)
		res.valueFloat = taskoutputshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := taskoutputshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := taskoutputshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", taskoutputshape.CornerOffsetRatio)
		res.valueFloat = taskoutputshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", taskoutputshape.IsHidden)
		res.valueBool = taskoutputshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (taskshape *TaskShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskshape.Name
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskshape.Task != nil {
			res.valueString = taskshape.Task.Name
			res.ids = taskshape.Task.GongGetUUID(stage)
		}
	case "IsShowDate":
		res.valueString = fmt.Sprintf("%t", taskshape.IsShowDate)
		res.valueBool = taskshape.IsShowDate
		res.GongFieldValueType = GongFieldValueTypeBool
	case "OverideLayoutDirection":
		res.valueString = fmt.Sprintf("%t", taskshape.OverideLayoutDirection)
		res.valueBool = taskshape.OverideLayoutDirection
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := taskshape.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "X":
		res.valueString = fmt.Sprintf("%f", taskshape.X)
		res.valueFloat = taskshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", taskshape.Y)
		res.valueFloat = taskshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", taskshape.Width)
		res.valueFloat = taskshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", taskshape.Height)
		res.valueFloat = taskshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", taskshape.IsHidden)
		res.valueBool = taskshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
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
	case "IsShowPrefix":
		diagram.IsShowPrefix = value.GetValueBool()
	case "DefaultBoxWidth":
		diagram.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagram.DefaultBoxHeigth = value.GetValueFloat()
	case "Width":
		diagram.Width = value.GetValueFloat()
	case "Height":
		diagram.Height = value.GetValueFloat()
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
		diagram.ProductsWhoseNodeIsExpanded = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.Product_stagedOrder[__instance__] == uint(id) {
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
	case "IsWBSNodeExpanded":
		diagram.IsWBSNodeExpanded = value.GetValueBool()
	case "Task_Shapes":
		diagram.Task_Shapes = make([]*TaskShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskShapes {
					if stage.TaskShape_stagedOrder[__instance__] == uint(id) {
						diagram.Task_Shapes = append(diagram.Task_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "TasksWhoseNodeIsExpanded":
		diagram.TasksWhoseNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						diagram.TasksWhoseNodeIsExpanded = append(diagram.TasksWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "TasksWhoseInputNodeIsExpanded":
		diagram.TasksWhoseInputNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						diagram.TasksWhoseInputNodeIsExpanded = append(diagram.TasksWhoseInputNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "TasksWhoseOutputNodeIsExpanded":
		diagram.TasksWhoseOutputNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						diagram.TasksWhoseOutputNodeIsExpanded = append(diagram.TasksWhoseOutputNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsTaskGroupsNodeExpanded":
		diagram.IsTaskGroupsNodeExpanded = value.GetValueBool()
	case "TaskGroupShapes":
		diagram.TaskGroupShapes = make([]*TaskGroupShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskGroupShapes {
					if stage.TaskGroupShape_stagedOrder[__instance__] == uint(id) {
						diagram.TaskGroupShapes = append(diagram.TaskGroupShapes, __instance__)
						break
					}
				}
			}
		}
	case "TaskGroupsWhoseNodeIsExpanded":
		diagram.TaskGroupsWhoseNodeIsExpanded = make([]*TaskGroup, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskGroups {
					if stage.TaskGroup_stagedOrder[__instance__] == uint(id) {
						diagram.TaskGroupsWhoseNodeIsExpanded = append(diagram.TaskGroupsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "DateFormat":
		diagram.DateFormat = value.GetValueString()
	case "TaskComposition_Shapes":
		diagram.TaskComposition_Shapes = make([]*TaskCompositionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskCompositionShapes {
					if stage.TaskCompositionShape_stagedOrder[__instance__] == uint(id) {
						diagram.TaskComposition_Shapes = append(diagram.TaskComposition_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "TaskInputShapes":
		diagram.TaskInputShapes = make([]*TaskInputShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskInputShapes {
					if stage.TaskInputShape_stagedOrder[__instance__] == uint(id) {
						diagram.TaskInputShapes = append(diagram.TaskInputShapes, __instance__)
						break
					}
				}
			}
		}
	case "TaskOutputShapes":
		diagram.TaskOutputShapes = make([]*TaskOutputShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskOutputShapes {
					if stage.TaskOutputShape_stagedOrder[__instance__] == uint(id) {
						diagram.TaskOutputShapes = append(diagram.TaskOutputShapes, __instance__)
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
		diagram.NoteResourceShapes = make([]*NoteResourceShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteResourceShapes {
					if stage.NoteResourceShape_stagedOrder[__instance__] == uint(id) {
						diagram.NoteResourceShapes = append(diagram.NoteResourceShapes, __instance__)
						break
					}
				}
			}
		}
	case "Resource_Shapes":
		diagram.Resource_Shapes = make([]*ResourceShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ResourceShapes {
					if stage.ResourceShape_stagedOrder[__instance__] == uint(id) {
						diagram.Resource_Shapes = append(diagram.Resource_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ResourcesWhoseNodeIsExpanded":
		diagram.ResourcesWhoseNodeIsExpanded = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						diagram.ResourcesWhoseNodeIsExpanded = append(diagram.ResourcesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsResourcesNodeExpanded":
		diagram.IsResourcesNodeExpanded = value.GetValueBool()
	case "ResourceComposition_Shapes":
		diagram.ResourceComposition_Shapes = make([]*ResourceCompositionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ResourceCompositionShapes {
					if stage.ResourceCompositionShape_stagedOrder[__instance__] == uint(id) {
						diagram.ResourceComposition_Shapes = append(diagram.ResourceComposition_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "ResourceTaskShapes":
		diagram.ResourceTaskShapes = make([]*ResourceTaskShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ResourceTaskShapes {
					if stage.ResourceTaskShape_stagedOrder[__instance__] == uint(id) {
						diagram.ResourceTaskShapes = append(diagram.ResourceTaskShapes, __instance__)
						break
					}
				}
			}
		}
	case "IsTimeDiagram":
		diagram.IsTimeDiagram = value.GetValueBool()
	case "UseManualStartAndEndDates":
		diagram.UseManualStartAndEndDates = value.GetValueBool()
	case "TimeStep":
		diagram.TimeStep = int(value.GetValueInt())
	case "TimeStepScale":
		diagram.TimeStepScale.FromCodeString(value.GetValueString())
	case "LaneHeight":
		diagram.LaneHeight = value.GetValueFloat()
	case "RatioBarToLaneHeight":
		diagram.RatioBarToLaneHeight = value.GetValueFloat()
	case "YTopMargin":
		diagram.YTopMargin = value.GetValueFloat()
	case "XLeftText":
		diagram.XLeftText = value.GetValueFloat()
	case "TextHeight":
		diagram.TextHeight = value.GetValueFloat()
	case "XLeftLanes":
		diagram.XLeftLanes = value.GetValueFloat()
	case "XRightMargin":
		diagram.XRightMargin = value.GetValueFloat()
	case "ArrowLengthToTheRightOfStartBar":
		diagram.ArrowLengthToTheRightOfStartBar = value.GetValueFloat()
	case "ArrowTipLenght":
		diagram.ArrowTipLenght = value.GetValueFloat()
	case "TimeLine_Color":
		diagram.TimeLine_Color = value.GetValueString()
	case "TimeLine_FillOpacity":
		diagram.TimeLine_FillOpacity = value.GetValueFloat()
	case "TimeLine_Stroke":
		diagram.TimeLine_Stroke = value.GetValueString()
	case "TimeLine_StrokeWidth":
		diagram.TimeLine_StrokeWidth = value.GetValueFloat()
	case "DrawVerticalTimeLines":
		diagram.DrawVerticalTimeLines = value.GetValueBool()
	case "Group_Stroke":
		diagram.Group_Stroke = value.GetValueString()
	case "Group_StrokeWidth":
		diagram.Group_StrokeWidth = value.GetValueFloat()
	case "Group_StrokeDashArray":
		diagram.Group_StrokeDashArray = value.GetValueString()
	case "DateYOffset":
		diagram.DateYOffset = value.GetValueFloat()
	case "AlignOnStartEndOnYearStart":
		diagram.AlignOnStartEndOnYearStart = value.GetValueBool()
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
	case "LogoSVGFile":
		library.LogoSVGFile = value.GetValueString()
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		library.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		library.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
	case "RootProducts":
		library.RootProducts = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.Product_stagedOrder[__instance__] == uint(id) {
						library.RootProducts = append(library.RootProducts, __instance__)
						break
					}
				}
			}
		}
	case "RootTasks":
		library.RootTasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						library.RootTasks = append(library.RootTasks, __instance__)
						break
					}
				}
			}
		}
	case "RootTaskGroups":
		library.RootTaskGroups = make([]*TaskGroup, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskGroups {
					if stage.TaskGroup_stagedOrder[__instance__] == uint(id) {
						library.RootTaskGroups = append(library.RootTaskGroups, __instance__)
						break
					}
				}
			}
		}
	case "RootResources":
		library.RootResources = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						library.RootResources = append(library.RootResources, __instance__)
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
		note.Products = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.Product_stagedOrder[__instance__] == uint(id) {
						note.Products = append(note.Products, __instance__)
						break
					}
				}
			}
		}
	case "Tasks":
		note.Tasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						note.Tasks = append(note.Tasks, __instance__)
						break
					}
				}
			}
		}
	case "Resources":
		note.Resources = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
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
			for __instance__ := range stage.Products {
				if stage.Product_stagedOrder[__instance__] == uint(id) {
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

func (noteresourceshape *NoteResourceShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		noteresourceshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteresourceshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					noteresourceshape.Note = __instance__
					break
				}
			}
		}
	case "Resource":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteresourceshape.Resource = nil
			for __instance__ := range stage.Resources {
				if stage.Resource_stagedOrder[__instance__] == uint(id) {
					noteresourceshape.Resource = __instance__
					break
				}
			}
		}
	case "StartRatio":
		noteresourceshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		noteresourceshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		noteresourceshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		noteresourceshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		noteresourceshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		noteresourceshape.IsHidden = value.GetValueBool()
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
	case "OverideLayoutDirection":
		noteshape.OverideLayoutDirection = value.GetValueBool()
	case "LayoutDirection":
		noteshape.LayoutDirection.FromCodeString(value.GetValueString())
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
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
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

func (product *Product) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		product.Name = value.GetValueString()
	case "ComputedPrefix":
		product.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		product.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		product.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsImport":
		product.IsImport = value.GetValueBool()
	case "ReferencedProduct":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			product.ReferencedProduct = nil
			for __instance__ := range stage.Products {
				if stage.Product_stagedOrder[__instance__] == uint(id) {
					product.ReferencedProduct = __instance__
					break
				}
			}
		}
	case "Description":
		product.Description = value.GetValueString()
	case "SubProducts":
		product.SubProducts = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.Product_stagedOrder[__instance__] == uint(id) {
						product.SubProducts = append(product.SubProducts, __instance__)
						break
					}
				}
			}
		}
	case "IsProducersNodeExpanded":
		product.IsProducersNodeExpanded = value.GetValueBool()
	case "IsConsumersNodeExpanded":
		product.IsConsumersNodeExpanded = value.GetValueBool()
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
			for __instance__ := range stage.Products {
				if stage.Product_stagedOrder[__instance__] == uint(id) {
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
			for __instance__ := range stage.Products {
				if stage.Product_stagedOrder[__instance__] == uint(id) {
					productshape.Product = __instance__
					break
				}
			}
		}
	case "OverideLayoutDirection":
		productshape.OverideLayoutDirection = value.GetValueBool()
	case "LayoutDirection":
		productshape.LayoutDirection.FromCodeString(value.GetValueString())
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

func (resource *Resource) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		resource.Name = value.GetValueString()
	case "ComputedPrefix":
		resource.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		resource.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		resource.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsImport":
		resource.IsImport = value.GetValueBool()
	case "ReferencedResource":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			resource.ReferencedResource = nil
			for __instance__ := range stage.Resources {
				if stage.Resource_stagedOrder[__instance__] == uint(id) {
					resource.ReferencedResource = __instance__
					break
				}
			}
		}
	case "Description":
		resource.Description = value.GetValueString()
	case "Tasks":
		resource.Tasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						resource.Tasks = append(resource.Tasks, __instance__)
						break
					}
				}
			}
		}
	case "SubResources":
		resource.SubResources = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						resource.SubResources = append(resource.SubResources, __instance__)
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

func (resourcecompositionshape *ResourceCompositionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		resourcecompositionshape.Name = value.GetValueString()
	case "Resource":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			resourcecompositionshape.Resource = nil
			for __instance__ := range stage.Resources {
				if stage.Resource_stagedOrder[__instance__] == uint(id) {
					resourcecompositionshape.Resource = __instance__
					break
				}
			}
		}
	case "StartRatio":
		resourcecompositionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		resourcecompositionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		resourcecompositionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		resourcecompositionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		resourcecompositionshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		resourcecompositionshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (resourceshape *ResourceShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		resourceshape.Name = value.GetValueString()
	case "Resource":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			resourceshape.Resource = nil
			for __instance__ := range stage.Resources {
				if stage.Resource_stagedOrder[__instance__] == uint(id) {
					resourceshape.Resource = __instance__
					break
				}
			}
		}
	case "OverideLayoutDirection":
		resourceshape.OverideLayoutDirection = value.GetValueBool()
	case "LayoutDirection":
		resourceshape.LayoutDirection.FromCodeString(value.GetValueString())
	case "X":
		resourceshape.X = value.GetValueFloat()
	case "Y":
		resourceshape.Y = value.GetValueFloat()
	case "Width":
		resourceshape.Width = value.GetValueFloat()
	case "Height":
		resourceshape.Height = value.GetValueFloat()
	case "IsHidden":
		resourceshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (resourcetaskshape *ResourceTaskShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		resourcetaskshape.Name = value.GetValueString()
	case "Resource":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			resourcetaskshape.Resource = nil
			for __instance__ := range stage.Resources {
				if stage.Resource_stagedOrder[__instance__] == uint(id) {
					resourcetaskshape.Resource = __instance__
					break
				}
			}
		}
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			resourcetaskshape.Task = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					resourcetaskshape.Task = __instance__
					break
				}
			}
		}
	case "StartRatio":
		resourcetaskshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		resourcetaskshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		resourcetaskshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		resourcetaskshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		resourcetaskshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		resourcetaskshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (task *Task) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		task.Name = value.GetValueString()
	case "ComputedPrefix":
		task.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		task.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		task.LayoutDirection.FromCodeString(value.GetValueString())
	case "IsImport":
		task.IsImport = value.GetValueBool()
	case "ReferencedTask":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			task.ReferencedTask = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					task.ReferencedTask = __instance__
					break
				}
			}
		}
	case "IsEndDateComputedFromDuration":
		task.IsEndDateComputedFromDuration = value.GetValueBool()
	case "Predecessors":
		task.Predecessors = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						task.Predecessors = append(task.Predecessors, __instance__)
						break
					}
				}
			}
		}
	case "IsStartDateComputedFromPredecessors":
		task.IsStartDateComputedFromPredecessors = value.GetValueBool()
	case "IsMilestone":
		task.IsMilestone = value.GetValueBool()
	case "Description":
		task.Description = value.GetValueString()
	case "SubTasks":
		task.SubTasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						task.SubTasks = append(task.SubTasks, __instance__)
						break
					}
				}
			}
		}
	case "Inputs":
		task.Inputs = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.Product_stagedOrder[__instance__] == uint(id) {
						task.Inputs = append(task.Inputs, __instance__)
						break
					}
				}
			}
		}
	case "IsInputsNodeExpanded":
		task.IsInputsNodeExpanded = value.GetValueBool()
	case "Outputs":
		task.Outputs = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.Product_stagedOrder[__instance__] == uint(id) {
						task.Outputs = append(task.Outputs, __instance__)
						break
					}
				}
			}
		}
	case "IsOutputsNodeExpanded":
		task.IsOutputsNodeExpanded = value.GetValueBool()
	case "IsWithCompletion":
		task.IsWithCompletion = value.GetValueBool()
	case "Completion":
		task.Completion.FromCodeString(value.GetValueString())
	case "DisplayVerticalBar":
		task.DisplayVerticalBar = value.GetValueBool()
	case "TaskGroupsToDisplay":
		task.TaskGroupsToDisplay = make([]*TaskGroup, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskGroups {
					if stage.TaskGroup_stagedOrder[__instance__] == uint(id) {
						task.TaskGroupsToDisplay = append(task.TaskGroupsToDisplay, __instance__)
						break
					}
				}
			}
		}
	case "TextPosition":
		task.TextPosition.FromCodeString(value.GetValueString())
	case "XOffset":
		task.XOffset = value.GetValueFloat()
	case "YOffset":
		task.YOffset = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (taskcompositionshape *TaskCompositionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskcompositionshape.Name = value.GetValueString()
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskcompositionshape.Task = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					taskcompositionshape.Task = __instance__
					break
				}
			}
		}
	case "StartRatio":
		taskcompositionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		taskcompositionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		taskcompositionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		taskcompositionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		taskcompositionshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		taskcompositionshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (taskgroup *TaskGroup) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskgroup.Name = value.GetValueString()
	case "ComputedPrefix":
		taskgroup.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		taskgroup.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		taskgroup.LayoutDirection.FromCodeString(value.GetValueString())
	case "Tasks":
		taskgroup.Tasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						taskgroup.Tasks = append(taskgroup.Tasks, __instance__)
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

func (taskgroupshape *TaskGroupShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskgroupshape.Name = value.GetValueString()
	case "TaskGroup":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskgroupshape.TaskGroup = nil
			for __instance__ := range stage.TaskGroups {
				if stage.TaskGroup_stagedOrder[__instance__] == uint(id) {
					taskgroupshape.TaskGroup = __instance__
					break
				}
			}
		}
	case "X":
		taskgroupshape.X = value.GetValueFloat()
	case "Y":
		taskgroupshape.Y = value.GetValueFloat()
	case "Width":
		taskgroupshape.Width = value.GetValueFloat()
	case "Height":
		taskgroupshape.Height = value.GetValueFloat()
	case "IsHidden":
		taskgroupshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (taskinputshape *TaskInputShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskinputshape.Name = value.GetValueString()
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskinputshape.Product = nil
			for __instance__ := range stage.Products {
				if stage.Product_stagedOrder[__instance__] == uint(id) {
					taskinputshape.Product = __instance__
					break
				}
			}
		}
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskinputshape.Task = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					taskinputshape.Task = __instance__
					break
				}
			}
		}
	case "StartRatio":
		taskinputshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		taskinputshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		taskinputshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		taskinputshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		taskinputshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		taskinputshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (taskoutputshape *TaskOutputShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskoutputshape.Name = value.GetValueString()
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskoutputshape.Task = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					taskoutputshape.Task = __instance__
					break
				}
			}
		}
	case "Product":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskoutputshape.Product = nil
			for __instance__ := range stage.Products {
				if stage.Product_stagedOrder[__instance__] == uint(id) {
					taskoutputshape.Product = __instance__
					break
				}
			}
		}
	case "StartRatio":
		taskoutputshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		taskoutputshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		taskoutputshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		taskoutputshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		taskoutputshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		taskoutputshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (taskshape *TaskShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		taskshape.Name = value.GetValueString()
	case "Task":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			taskshape.Task = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					taskshape.Task = __instance__
					break
				}
			}
		}
	case "IsShowDate":
		taskshape.IsShowDate = value.GetValueBool()
	case "OverideLayoutDirection":
		taskshape.OverideLayoutDirection = value.GetValueBool()
	case "LayoutDirection":
		taskshape.LayoutDirection.FromCodeString(value.GetValueString())
	case "X":
		taskshape.X = value.GetValueFloat()
	case "Y":
		taskshape.Y = value.GetValueFloat()
	case "Width":
		taskshape.Width = value.GetValueFloat()
	case "Height":
		taskshape.Height = value.GetValueFloat()
	case "IsHidden":
		taskshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
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

func (noteresourceshape *NoteResourceShape) GongGetGongstructName() string {
	return "NoteResourceShape"
}

func (noteshape *NoteShape) GongGetGongstructName() string {
	return "NoteShape"
}

func (notetaskshape *NoteTaskShape) GongGetGongstructName() string {
	return "NoteTaskShape"
}

func (product *Product) GongGetGongstructName() string {
	return "Product"
}

func (productcompositionshape *ProductCompositionShape) GongGetGongstructName() string {
	return "ProductCompositionShape"
}

func (productshape *ProductShape) GongGetGongstructName() string {
	return "ProductShape"
}

func (resource *Resource) GongGetGongstructName() string {
	return "Resource"
}

func (resourcecompositionshape *ResourceCompositionShape) GongGetGongstructName() string {
	return "ResourceCompositionShape"
}

func (resourceshape *ResourceShape) GongGetGongstructName() string {
	return "ResourceShape"
}

func (resourcetaskshape *ResourceTaskShape) GongGetGongstructName() string {
	return "ResourceTaskShape"
}

func (task *Task) GongGetGongstructName() string {
	return "Task"
}

func (taskcompositionshape *TaskCompositionShape) GongGetGongstructName() string {
	return "TaskCompositionShape"
}

func (taskgroup *TaskGroup) GongGetGongstructName() string {
	return "TaskGroup"
}

func (taskgroupshape *TaskGroupShape) GongGetGongstructName() string {
	return "TaskGroupShape"
}

func (taskinputshape *TaskInputShape) GongGetGongstructName() string {
	return "TaskInputShape"
}

func (taskoutputshape *TaskOutputShape) GongGetGongstructName() string {
	return "TaskOutputShape"
}

func (taskshape *TaskShape) GongGetGongstructName() string {
	return "TaskShape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
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

	stage.NoteResourceShapes_mapString = make(map[string]*NoteResourceShape)
	for noteresourceshape := range stage.NoteResourceShapes {
		stage.NoteResourceShapes_mapString[noteresourceshape.Name] = noteresourceshape
	}

	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	for noteshape := range stage.NoteShapes {
		stage.NoteShapes_mapString[noteshape.Name] = noteshape
	}

	stage.NoteTaskShapes_mapString = make(map[string]*NoteTaskShape)
	for notetaskshape := range stage.NoteTaskShapes {
		stage.NoteTaskShapes_mapString[notetaskshape.Name] = notetaskshape
	}

	stage.Products_mapString = make(map[string]*Product)
	for product := range stage.Products {
		stage.Products_mapString[product.Name] = product
	}

	stage.ProductCompositionShapes_mapString = make(map[string]*ProductCompositionShape)
	for productcompositionshape := range stage.ProductCompositionShapes {
		stage.ProductCompositionShapes_mapString[productcompositionshape.Name] = productcompositionshape
	}

	stage.ProductShapes_mapString = make(map[string]*ProductShape)
	for productshape := range stage.ProductShapes {
		stage.ProductShapes_mapString[productshape.Name] = productshape
	}

	stage.Resources_mapString = make(map[string]*Resource)
	for resource := range stage.Resources {
		stage.Resources_mapString[resource.Name] = resource
	}

	stage.ResourceCompositionShapes_mapString = make(map[string]*ResourceCompositionShape)
	for resourcecompositionshape := range stage.ResourceCompositionShapes {
		stage.ResourceCompositionShapes_mapString[resourcecompositionshape.Name] = resourcecompositionshape
	}

	stage.ResourceShapes_mapString = make(map[string]*ResourceShape)
	for resourceshape := range stage.ResourceShapes {
		stage.ResourceShapes_mapString[resourceshape.Name] = resourceshape
	}

	stage.ResourceTaskShapes_mapString = make(map[string]*ResourceTaskShape)
	for resourcetaskshape := range stage.ResourceTaskShapes {
		stage.ResourceTaskShapes_mapString[resourcetaskshape.Name] = resourcetaskshape
	}

	stage.Tasks_mapString = make(map[string]*Task)
	for task := range stage.Tasks {
		stage.Tasks_mapString[task.Name] = task
	}

	stage.TaskCompositionShapes_mapString = make(map[string]*TaskCompositionShape)
	for taskcompositionshape := range stage.TaskCompositionShapes {
		stage.TaskCompositionShapes_mapString[taskcompositionshape.Name] = taskcompositionshape
	}

	stage.TaskGroups_mapString = make(map[string]*TaskGroup)
	for taskgroup := range stage.TaskGroups {
		stage.TaskGroups_mapString[taskgroup.Name] = taskgroup
	}

	stage.TaskGroupShapes_mapString = make(map[string]*TaskGroupShape)
	for taskgroupshape := range stage.TaskGroupShapes {
		stage.TaskGroupShapes_mapString[taskgroupshape.Name] = taskgroupshape
	}

	stage.TaskInputShapes_mapString = make(map[string]*TaskInputShape)
	for taskinputshape := range stage.TaskInputShapes {
		stage.TaskInputShapes_mapString[taskinputshape.Name] = taskinputshape
	}

	stage.TaskOutputShapes_mapString = make(map[string]*TaskOutputShape)
	for taskoutputshape := range stage.TaskOutputShapes {
		stage.TaskOutputShapes_mapString[taskoutputshape.Name] = taskoutputshape
	}

	stage.TaskShapes_mapString = make(map[string]*TaskShape)
	for taskshape := range stage.TaskShapes {
		stage.TaskShapes_mapString[taskshape.Name] = taskshape
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
