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

	Diagram_ProductReference_Shapes_reverseMap map[*ProductReferenceShape]*Diagram

	Diagram_Task_Shapes_reverseMap map[*TaskShape]*Diagram

	Diagram_TasksWhoseNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TasksWhoseInputNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TasksWhoseOutputNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TasksWhosePredecessorNodeIsExpanded_reverseMap map[*Task]*Diagram

	Diagram_TaskGroupShapes_reverseMap map[*TaskGroupShape]*Diagram

	Diagram_TaskGroupsWhoseNodeIsExpanded_reverseMap map[*TaskGroup]*Diagram

	Diagram_TaskComposition_Shapes_reverseMap map[*TaskCompositionShape]*Diagram

	Diagram_TaskInputShapes_reverseMap map[*TaskInputShape]*Diagram

	Diagram_TaskOutputShapes_reverseMap map[*TaskOutputShape]*Diagram

	Diagram_TaskPredecessorShapes_reverseMap map[*TaskPredecessorShape]*Diagram

	Diagram_Note_Shapes_reverseMap map[*NoteShape]*Diagram

	Diagram_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*Diagram

	Diagram_NoteProductShapes_reverseMap map[*NoteProductShape]*Diagram

	Diagram_NoteTaskShapes_reverseMap map[*NoteTaskShape]*Diagram

	Diagram_NoteResourceShapes_reverseMap map[*NoteResourceShape]*Diagram

	Diagram_Resource_Shapes_reverseMap map[*ResourceShape]*Diagram

	Diagram_ResourcesWhoseNodeIsExpanded_reverseMap map[*Resource]*Diagram

	Diagram_ResourceComposition_Shapes_reverseMap map[*ResourceCompositionShape]*Diagram

	Diagram_ResourceTaskShapes_reverseMap map[*ResourceTaskShape]*Diagram

	OnAfterDiagramCreateCallback GongOnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback GongOnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback GongOnAfterDeleteInterface[Diagram]

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
	Note_Products_reverseMap map[*Product]*Note

	Note_Tasks_reverseMap map[*Task]*Note

	Note_Resources_reverseMap map[*Resource]*Note

	OnAfterNoteCreateCallback GongOnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback GongOnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback GongOnAfterDeleteInterface[Note]

	NoteProductShapes                map[*NoteProductShape]struct{}
	NoteProductShapes_instance       map[*NoteProductShape]*NoteProductShape
	NoteProductShapes_mapString      map[string]*NoteProductShape
	NoteProductShapeOrder            uint
	NoteProductShape_stagedOrder     map[*NoteProductShape]uint
	NoteProductShape_orderStaged     map[uint]*NoteProductShape
	NoteProductShapes_reference      map[*NoteProductShape]*NoteProductShape
	NoteProductShapes_referenceOrder map[*NoteProductShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteProductShapeCreateCallback GongOnAfterCreateInterface[NoteProductShape]
	OnAfterNoteProductShapeUpdateCallback GongOnAfterUpdateInterface[NoteProductShape]
	OnAfterNoteProductShapeDeleteCallback GongOnAfterDeleteInterface[NoteProductShape]

	NoteResourceShapes                map[*NoteResourceShape]struct{}
	NoteResourceShapes_instance       map[*NoteResourceShape]*NoteResourceShape
	NoteResourceShapes_mapString      map[string]*NoteResourceShape
	NoteResourceShapeOrder            uint
	NoteResourceShape_stagedOrder     map[*NoteResourceShape]uint
	NoteResourceShape_orderStaged     map[uint]*NoteResourceShape
	NoteResourceShapes_reference      map[*NoteResourceShape]*NoteResourceShape
	NoteResourceShapes_referenceOrder map[*NoteResourceShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteResourceShapeCreateCallback GongOnAfterCreateInterface[NoteResourceShape]
	OnAfterNoteResourceShapeUpdateCallback GongOnAfterUpdateInterface[NoteResourceShape]
	OnAfterNoteResourceShapeDeleteCallback GongOnAfterDeleteInterface[NoteResourceShape]

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

	NoteTaskShapes                map[*NoteTaskShape]struct{}
	NoteTaskShapes_instance       map[*NoteTaskShape]*NoteTaskShape
	NoteTaskShapes_mapString      map[string]*NoteTaskShape
	NoteTaskShapeOrder            uint
	NoteTaskShape_stagedOrder     map[*NoteTaskShape]uint
	NoteTaskShape_orderStaged     map[uint]*NoteTaskShape
	NoteTaskShapes_reference      map[*NoteTaskShape]*NoteTaskShape
	NoteTaskShapes_referenceOrder map[*NoteTaskShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteTaskShapeCreateCallback GongOnAfterCreateInterface[NoteTaskShape]
	OnAfterNoteTaskShapeUpdateCallback GongOnAfterUpdateInterface[NoteTaskShape]
	OnAfterNoteTaskShapeDeleteCallback GongOnAfterDeleteInterface[NoteTaskShape]

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

	OnAfterProductCreateCallback GongOnAfterCreateInterface[Product]
	OnAfterProductUpdateCallback GongOnAfterUpdateInterface[Product]
	OnAfterProductDeleteCallback GongOnAfterDeleteInterface[Product]

	ProductCompositionShapes                map[*ProductCompositionShape]struct{}
	ProductCompositionShapes_instance       map[*ProductCompositionShape]*ProductCompositionShape
	ProductCompositionShapes_mapString      map[string]*ProductCompositionShape
	ProductCompositionShapeOrder            uint
	ProductCompositionShape_stagedOrder     map[*ProductCompositionShape]uint
	ProductCompositionShape_orderStaged     map[uint]*ProductCompositionShape
	ProductCompositionShapes_reference      map[*ProductCompositionShape]*ProductCompositionShape
	ProductCompositionShapes_referenceOrder map[*ProductCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterProductCompositionShapeCreateCallback GongOnAfterCreateInterface[ProductCompositionShape]
	OnAfterProductCompositionShapeUpdateCallback GongOnAfterUpdateInterface[ProductCompositionShape]
	OnAfterProductCompositionShapeDeleteCallback GongOnAfterDeleteInterface[ProductCompositionShape]

	ProductReferenceShapes                map[*ProductReferenceShape]struct{}
	ProductReferenceShapes_instance       map[*ProductReferenceShape]*ProductReferenceShape
	ProductReferenceShapes_mapString      map[string]*ProductReferenceShape
	ProductReferenceShapeOrder            uint
	ProductReferenceShape_stagedOrder     map[*ProductReferenceShape]uint
	ProductReferenceShape_orderStaged     map[uint]*ProductReferenceShape
	ProductReferenceShapes_reference      map[*ProductReferenceShape]*ProductReferenceShape
	ProductReferenceShapes_referenceOrder map[*ProductReferenceShape]uint

	// insertion point for slice of pointers maps
	OnAfterProductReferenceShapeCreateCallback GongOnAfterCreateInterface[ProductReferenceShape]
	OnAfterProductReferenceShapeUpdateCallback GongOnAfterUpdateInterface[ProductReferenceShape]
	OnAfterProductReferenceShapeDeleteCallback GongOnAfterDeleteInterface[ProductReferenceShape]

	ProductShapes                map[*ProductShape]struct{}
	ProductShapes_instance       map[*ProductShape]*ProductShape
	ProductShapes_mapString      map[string]*ProductShape
	ProductShapeOrder            uint
	ProductShape_stagedOrder     map[*ProductShape]uint
	ProductShape_orderStaged     map[uint]*ProductShape
	ProductShapes_reference      map[*ProductShape]*ProductShape
	ProductShapes_referenceOrder map[*ProductShape]uint

	// insertion point for slice of pointers maps
	OnAfterProductShapeCreateCallback GongOnAfterCreateInterface[ProductShape]
	OnAfterProductShapeUpdateCallback GongOnAfterUpdateInterface[ProductShape]
	OnAfterProductShapeDeleteCallback GongOnAfterDeleteInterface[ProductShape]

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

	OnAfterResourceCreateCallback GongOnAfterCreateInterface[Resource]
	OnAfterResourceUpdateCallback GongOnAfterUpdateInterface[Resource]
	OnAfterResourceDeleteCallback GongOnAfterDeleteInterface[Resource]

	ResourceCompositionShapes                map[*ResourceCompositionShape]struct{}
	ResourceCompositionShapes_instance       map[*ResourceCompositionShape]*ResourceCompositionShape
	ResourceCompositionShapes_mapString      map[string]*ResourceCompositionShape
	ResourceCompositionShapeOrder            uint
	ResourceCompositionShape_stagedOrder     map[*ResourceCompositionShape]uint
	ResourceCompositionShape_orderStaged     map[uint]*ResourceCompositionShape
	ResourceCompositionShapes_reference      map[*ResourceCompositionShape]*ResourceCompositionShape
	ResourceCompositionShapes_referenceOrder map[*ResourceCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterResourceCompositionShapeCreateCallback GongOnAfterCreateInterface[ResourceCompositionShape]
	OnAfterResourceCompositionShapeUpdateCallback GongOnAfterUpdateInterface[ResourceCompositionShape]
	OnAfterResourceCompositionShapeDeleteCallback GongOnAfterDeleteInterface[ResourceCompositionShape]

	ResourceShapes                map[*ResourceShape]struct{}
	ResourceShapes_instance       map[*ResourceShape]*ResourceShape
	ResourceShapes_mapString      map[string]*ResourceShape
	ResourceShapeOrder            uint
	ResourceShape_stagedOrder     map[*ResourceShape]uint
	ResourceShape_orderStaged     map[uint]*ResourceShape
	ResourceShapes_reference      map[*ResourceShape]*ResourceShape
	ResourceShapes_referenceOrder map[*ResourceShape]uint

	// insertion point for slice of pointers maps
	OnAfterResourceShapeCreateCallback GongOnAfterCreateInterface[ResourceShape]
	OnAfterResourceShapeUpdateCallback GongOnAfterUpdateInterface[ResourceShape]
	OnAfterResourceShapeDeleteCallback GongOnAfterDeleteInterface[ResourceShape]

	ResourceTaskShapes                map[*ResourceTaskShape]struct{}
	ResourceTaskShapes_instance       map[*ResourceTaskShape]*ResourceTaskShape
	ResourceTaskShapes_mapString      map[string]*ResourceTaskShape
	ResourceTaskShapeOrder            uint
	ResourceTaskShape_stagedOrder     map[*ResourceTaskShape]uint
	ResourceTaskShape_orderStaged     map[uint]*ResourceTaskShape
	ResourceTaskShapes_reference      map[*ResourceTaskShape]*ResourceTaskShape
	ResourceTaskShapes_referenceOrder map[*ResourceTaskShape]uint

	// insertion point for slice of pointers maps
	OnAfterResourceTaskShapeCreateCallback GongOnAfterCreateInterface[ResourceTaskShape]
	OnAfterResourceTaskShapeUpdateCallback GongOnAfterUpdateInterface[ResourceTaskShape]
	OnAfterResourceTaskShapeDeleteCallback GongOnAfterDeleteInterface[ResourceTaskShape]

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

	Task_Inputs_reverseMap map[*Product]*Task

	Task_Outputs_reverseMap map[*Product]*Task

	Task_SubTasks_reverseMap map[*Task]*Task

	Task_TaskGroupsToDisplay_reverseMap map[*TaskGroup]*Task

	OnAfterTaskCreateCallback GongOnAfterCreateInterface[Task]
	OnAfterTaskUpdateCallback GongOnAfterUpdateInterface[Task]
	OnAfterTaskDeleteCallback GongOnAfterDeleteInterface[Task]

	TaskCompositionShapes                map[*TaskCompositionShape]struct{}
	TaskCompositionShapes_instance       map[*TaskCompositionShape]*TaskCompositionShape
	TaskCompositionShapes_mapString      map[string]*TaskCompositionShape
	TaskCompositionShapeOrder            uint
	TaskCompositionShape_stagedOrder     map[*TaskCompositionShape]uint
	TaskCompositionShape_orderStaged     map[uint]*TaskCompositionShape
	TaskCompositionShapes_reference      map[*TaskCompositionShape]*TaskCompositionShape
	TaskCompositionShapes_referenceOrder map[*TaskCompositionShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskCompositionShapeCreateCallback GongOnAfterCreateInterface[TaskCompositionShape]
	OnAfterTaskCompositionShapeUpdateCallback GongOnAfterUpdateInterface[TaskCompositionShape]
	OnAfterTaskCompositionShapeDeleteCallback GongOnAfterDeleteInterface[TaskCompositionShape]

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

	OnAfterTaskGroupCreateCallback GongOnAfterCreateInterface[TaskGroup]
	OnAfterTaskGroupUpdateCallback GongOnAfterUpdateInterface[TaskGroup]
	OnAfterTaskGroupDeleteCallback GongOnAfterDeleteInterface[TaskGroup]

	TaskGroupShapes                map[*TaskGroupShape]struct{}
	TaskGroupShapes_instance       map[*TaskGroupShape]*TaskGroupShape
	TaskGroupShapes_mapString      map[string]*TaskGroupShape
	TaskGroupShapeOrder            uint
	TaskGroupShape_stagedOrder     map[*TaskGroupShape]uint
	TaskGroupShape_orderStaged     map[uint]*TaskGroupShape
	TaskGroupShapes_reference      map[*TaskGroupShape]*TaskGroupShape
	TaskGroupShapes_referenceOrder map[*TaskGroupShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskGroupShapeCreateCallback GongOnAfterCreateInterface[TaskGroupShape]
	OnAfterTaskGroupShapeUpdateCallback GongOnAfterUpdateInterface[TaskGroupShape]
	OnAfterTaskGroupShapeDeleteCallback GongOnAfterDeleteInterface[TaskGroupShape]

	TaskInputShapes                map[*TaskInputShape]struct{}
	TaskInputShapes_instance       map[*TaskInputShape]*TaskInputShape
	TaskInputShapes_mapString      map[string]*TaskInputShape
	TaskInputShapeOrder            uint
	TaskInputShape_stagedOrder     map[*TaskInputShape]uint
	TaskInputShape_orderStaged     map[uint]*TaskInputShape
	TaskInputShapes_reference      map[*TaskInputShape]*TaskInputShape
	TaskInputShapes_referenceOrder map[*TaskInputShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskInputShapeCreateCallback GongOnAfterCreateInterface[TaskInputShape]
	OnAfterTaskInputShapeUpdateCallback GongOnAfterUpdateInterface[TaskInputShape]
	OnAfterTaskInputShapeDeleteCallback GongOnAfterDeleteInterface[TaskInputShape]

	TaskOutputShapes                map[*TaskOutputShape]struct{}
	TaskOutputShapes_instance       map[*TaskOutputShape]*TaskOutputShape
	TaskOutputShapes_mapString      map[string]*TaskOutputShape
	TaskOutputShapeOrder            uint
	TaskOutputShape_stagedOrder     map[*TaskOutputShape]uint
	TaskOutputShape_orderStaged     map[uint]*TaskOutputShape
	TaskOutputShapes_reference      map[*TaskOutputShape]*TaskOutputShape
	TaskOutputShapes_referenceOrder map[*TaskOutputShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskOutputShapeCreateCallback GongOnAfterCreateInterface[TaskOutputShape]
	OnAfterTaskOutputShapeUpdateCallback GongOnAfterUpdateInterface[TaskOutputShape]
	OnAfterTaskOutputShapeDeleteCallback GongOnAfterDeleteInterface[TaskOutputShape]

	TaskPredecessorShapes                map[*TaskPredecessorShape]struct{}
	TaskPredecessorShapes_instance       map[*TaskPredecessorShape]*TaskPredecessorShape
	TaskPredecessorShapes_mapString      map[string]*TaskPredecessorShape
	TaskPredecessorShapeOrder            uint
	TaskPredecessorShape_stagedOrder     map[*TaskPredecessorShape]uint
	TaskPredecessorShape_orderStaged     map[uint]*TaskPredecessorShape
	TaskPredecessorShapes_reference      map[*TaskPredecessorShape]*TaskPredecessorShape
	TaskPredecessorShapes_referenceOrder map[*TaskPredecessorShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskPredecessorShapeCreateCallback GongOnAfterCreateInterface[TaskPredecessorShape]
	OnAfterTaskPredecessorShapeUpdateCallback GongOnAfterUpdateInterface[TaskPredecessorShape]
	OnAfterTaskPredecessorShapeDeleteCallback GongOnAfterDeleteInterface[TaskPredecessorShape]

	TaskShapes                map[*TaskShape]struct{}
	TaskShapes_instance       map[*TaskShape]*TaskShape
	TaskShapes_mapString      map[string]*TaskShape
	TaskShapeOrder            uint
	TaskShape_stagedOrder     map[*TaskShape]uint
	TaskShape_orderStaged     map[uint]*TaskShape
	TaskShapes_reference      map[*TaskShape]*TaskShape
	TaskShapes_referenceOrder map[*TaskShape]uint

	// insertion point for slice of pointers maps
	OnAfterTaskShapeCreateCallback GongOnAfterCreateInterface[TaskShape]
	OnAfterTaskShapeUpdateCallback GongOnAfterUpdateInterface[TaskShape]
	OnAfterTaskShapeDeleteCallback GongOnAfterDeleteInterface[TaskShape]

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
	__gong__clearReferences(&stage.Diagrams_reference, &stage.Diagrams_instance, &stage.Diagrams_referenceOrder)

	__gong__clearReferences(&stage.Librarys_reference, &stage.Librarys_instance, &stage.Librarys_referenceOrder)

	__gong__clearReferences(&stage.Notes_reference, &stage.Notes_instance, &stage.Notes_referenceOrder)

	__gong__clearReferences(&stage.NoteProductShapes_reference, &stage.NoteProductShapes_instance, &stage.NoteProductShapes_referenceOrder)

	__gong__clearReferences(&stage.NoteResourceShapes_reference, &stage.NoteResourceShapes_instance, &stage.NoteResourceShapes_referenceOrder)

	__gong__clearReferences(&stage.NoteShapes_reference, &stage.NoteShapes_instance, &stage.NoteShapes_referenceOrder)

	__gong__clearReferences(&stage.NoteTaskShapes_reference, &stage.NoteTaskShapes_instance, &stage.NoteTaskShapes_referenceOrder)

	__gong__clearReferences(&stage.Products_reference, &stage.Products_instance, &stage.Products_referenceOrder)

	__gong__clearReferences(&stage.ProductCompositionShapes_reference, &stage.ProductCompositionShapes_instance, &stage.ProductCompositionShapes_referenceOrder)

	__gong__clearReferences(&stage.ProductReferenceShapes_reference, &stage.ProductReferenceShapes_instance, &stage.ProductReferenceShapes_referenceOrder)

	__gong__clearReferences(&stage.ProductShapes_reference, &stage.ProductShapes_instance, &stage.ProductShapes_referenceOrder)

	__gong__clearReferences(&stage.Resources_reference, &stage.Resources_instance, &stage.Resources_referenceOrder)

	__gong__clearReferences(&stage.ResourceCompositionShapes_reference, &stage.ResourceCompositionShapes_instance, &stage.ResourceCompositionShapes_referenceOrder)

	__gong__clearReferences(&stage.ResourceShapes_reference, &stage.ResourceShapes_instance, &stage.ResourceShapes_referenceOrder)

	__gong__clearReferences(&stage.ResourceTaskShapes_reference, &stage.ResourceTaskShapes_instance, &stage.ResourceTaskShapes_referenceOrder)

	__gong__clearReferences(&stage.Tasks_reference, &stage.Tasks_instance, &stage.Tasks_referenceOrder)

	__gong__clearReferences(&stage.TaskCompositionShapes_reference, &stage.TaskCompositionShapes_instance, &stage.TaskCompositionShapes_referenceOrder)

	__gong__clearReferences(&stage.TaskGroups_reference, &stage.TaskGroups_instance, &stage.TaskGroups_referenceOrder)

	__gong__clearReferences(&stage.TaskGroupShapes_reference, &stage.TaskGroupShapes_instance, &stage.TaskGroupShapes_referenceOrder)

	__gong__clearReferences(&stage.TaskInputShapes_reference, &stage.TaskInputShapes_instance, &stage.TaskInputShapes_referenceOrder)

	__gong__clearReferences(&stage.TaskOutputShapes_reference, &stage.TaskOutputShapes_instance, &stage.TaskOutputShapes_referenceOrder)

	__gong__clearReferences(&stage.TaskPredecessorShapes_reference, &stage.TaskPredecessorShapes_instance, &stage.TaskPredecessorShapes_referenceOrder)

	__gong__clearReferences(&stage.TaskShapes_reference, &stage.TaskShapes_instance, &stage.TaskShapes_referenceOrder)

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
	stage.DiagramOrder = __gong__recomputeOrder(stage.Diagram_stagedOrder)

	stage.LibraryOrder = __gong__recomputeOrder(stage.Library_stagedOrder)

	stage.NoteOrder = __gong__recomputeOrder(stage.Note_stagedOrder)

	stage.NoteProductShapeOrder = __gong__recomputeOrder(stage.NoteProductShape_stagedOrder)

	stage.NoteResourceShapeOrder = __gong__recomputeOrder(stage.NoteResourceShape_stagedOrder)

	stage.NoteShapeOrder = __gong__recomputeOrder(stage.NoteShape_stagedOrder)

	stage.NoteTaskShapeOrder = __gong__recomputeOrder(stage.NoteTaskShape_stagedOrder)

	stage.ProductOrder = __gong__recomputeOrder(stage.Product_stagedOrder)

	stage.ProductCompositionShapeOrder = __gong__recomputeOrder(stage.ProductCompositionShape_stagedOrder)

	stage.ProductReferenceShapeOrder = __gong__recomputeOrder(stage.ProductReferenceShape_stagedOrder)

	stage.ProductShapeOrder = __gong__recomputeOrder(stage.ProductShape_stagedOrder)

	stage.ResourceOrder = __gong__recomputeOrder(stage.Resource_stagedOrder)

	stage.ResourceCompositionShapeOrder = __gong__recomputeOrder(stage.ResourceCompositionShape_stagedOrder)

	stage.ResourceShapeOrder = __gong__recomputeOrder(stage.ResourceShape_stagedOrder)

	stage.ResourceTaskShapeOrder = __gong__recomputeOrder(stage.ResourceTaskShape_stagedOrder)

	stage.TaskOrder = __gong__recomputeOrder(stage.Task_stagedOrder)

	stage.TaskCompositionShapeOrder = __gong__recomputeOrder(stage.TaskCompositionShape_stagedOrder)

	stage.TaskGroupOrder = __gong__recomputeOrder(stage.TaskGroup_stagedOrder)

	stage.TaskGroupShapeOrder = __gong__recomputeOrder(stage.TaskGroupShape_stagedOrder)

	stage.TaskInputShapeOrder = __gong__recomputeOrder(stage.TaskInputShape_stagedOrder)

	stage.TaskOutputShapeOrder = __gong__recomputeOrder(stage.TaskOutputShape_stagedOrder)

	stage.TaskPredecessorShapeOrder = __gong__recomputeOrder(stage.TaskPredecessorShape_stagedOrder)

	stage.TaskShapeOrder = __gong__recomputeOrder(stage.TaskShape_stagedOrder)

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
	case *Diagram:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Diagrams, stage.Diagram_stagedOrder))
	case *Library:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder))
	case *Note:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Notes, stage.Note_stagedOrder))
	case *NoteProductShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteProductShapes, stage.NoteProductShape_stagedOrder))
	case *NoteResourceShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteResourceShapes, stage.NoteResourceShape_stagedOrder))
	case *NoteShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteShapes, stage.NoteShape_stagedOrder))
	case *NoteTaskShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder))
	case *Product:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Products, stage.Product_stagedOrder))
	case *ProductCompositionShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ProductCompositionShapes, stage.ProductCompositionShape_stagedOrder))
	case *ProductReferenceShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ProductReferenceShapes, stage.ProductReferenceShape_stagedOrder))
	case *ProductShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ProductShapes, stage.ProductShape_stagedOrder))
	case *Resource:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Resources, stage.Resource_stagedOrder))
	case *ResourceCompositionShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ResourceCompositionShapes, stage.ResourceCompositionShape_stagedOrder))
	case *ResourceShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ResourceShapes, stage.ResourceShape_stagedOrder))
	case *ResourceTaskShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ResourceTaskShapes, stage.ResourceTaskShape_stagedOrder))
	case *Task:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Tasks, stage.Task_stagedOrder))
	case *TaskCompositionShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TaskCompositionShapes, stage.TaskCompositionShape_stagedOrder))
	case *TaskGroup:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TaskGroups, stage.TaskGroup_stagedOrder))
	case *TaskGroupShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TaskGroupShapes, stage.TaskGroupShape_stagedOrder))
	case *TaskInputShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TaskInputShapes, stage.TaskInputShape_stagedOrder))
	case *TaskOutputShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TaskOutputShapes, stage.TaskOutputShape_stagedOrder))
	case *TaskPredecessorShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TaskPredecessorShapes, stage.TaskPredecessorShape_stagedOrder))
	case *TaskShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.TaskShapes, stage.TaskShape_stagedOrder))

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
	return "github.com/fullstack-lang/gong/dsm/project/go/models"
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

		ProductReferenceShapes:           make(map[*ProductReferenceShape]struct{}),
		ProductReferenceShapes_mapString: make(map[string]*ProductReferenceShape),

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

		TaskPredecessorShapes:           make(map[*TaskPredecessorShape]struct{}),
		TaskPredecessorShapes_mapString: make(map[string]*TaskPredecessorShape),

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

		ProductReferenceShape_stagedOrder: make(map[*ProductReferenceShape]uint),
		ProductReferenceShape_orderStaged: make(map[uint]*ProductReferenceShape),
		ProductReferenceShapes_reference:  make(map[*ProductReferenceShape]*ProductReferenceShape),

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

		TaskPredecessorShape_stagedOrder: make(map[*TaskPredecessorShape]uint),
		TaskPredecessorShape_orderStaged: make(map[uint]*TaskPredecessorShape),
		TaskPredecessorShapes_reference:  make(map[*TaskPredecessorShape]*TaskPredecessorShape),

		TaskShape_stagedOrder: make(map[*TaskShape]uint),
		TaskShape_orderStaged: make(map[uint]*TaskShape),
		TaskShapes_reference:  make(map[*TaskShape]*TaskShape),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Diagram": &DiagramUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteProductShape": &NoteProductShapeUnmarshaller{},

			"NoteResourceShape": &NoteResourceShapeUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"NoteTaskShape": &NoteTaskShapeUnmarshaller{},

			"Product": &ProductUnmarshaller{},

			"ProductCompositionShape": &ProductCompositionShapeUnmarshaller{},

			"ProductReferenceShape": &ProductReferenceShapeUnmarshaller{},

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

			"TaskPredecessorShape": &TaskPredecessorShapeUnmarshaller{},

			"TaskShape": &TaskShapeUnmarshaller{},

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
	case *ProductReferenceShape:
		return any(stage.ProductReferenceShape_orderStaged[order]).(Type)
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
	case *TaskPredecessorShape:
		return any(stage.TaskPredecessorShape_orderStaged[order]).(Type)
	case *TaskShape:
		return any(stage.TaskShape_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteProductShape"] = len(stage.NoteProductShapes)
	stage.Map_GongStructName_InstancesNb["NoteResourceShape"] = len(stage.NoteResourceShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteTaskShape"] = len(stage.NoteTaskShapes)
	stage.Map_GongStructName_InstancesNb["Product"] = len(stage.Products)
	stage.Map_GongStructName_InstancesNb["ProductCompositionShape"] = len(stage.ProductCompositionShapes)
	stage.Map_GongStructName_InstancesNb["ProductReferenceShape"] = len(stage.ProductReferenceShapes)
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
	stage.Map_GongStructName_InstancesNb["TaskPredecessorShape"] = len(stage.TaskPredecessorShapes)
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

// Stage puts noteproductshape to the model stage
func (noteproductshape *NoteProductShape) Stage(stage *Stage) *NoteProductShape {
	__gong__stage(stage.NoteProductShapes, stage.NoteProductShape_stagedOrder, stage.NoteProductShape_orderStaged, &stage.NoteProductShapeOrder, stage.NoteProductShapes_mapString, noteproductshape, noteproductshape.Name)
	return noteproductshape
}

// StagePreserveOrder puts noteproductshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteProductShapeOrder
// - update stage.NoteProductShapeOrder accordingly
func (noteproductshape *NoteProductShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteProductShapes, stage.NoteProductShape_stagedOrder, stage.NoteProductShape_orderStaged, &stage.NoteProductShapeOrder, stage.NoteProductShapes_mapString, noteproductshape, order, noteproductshape.Name)
}

// Unstage removes noteproductshape off the model stage
func (noteproductshape *NoteProductShape) Unstage(stage *Stage) *NoteProductShape {
	__gong__unstage(stage.NoteProductShapes, stage.NoteProductShapes_mapString, noteproductshape, noteproductshape.Name)
	return noteproductshape
}

// UnstageVoid removes noteproductshape off the model stage
func (noteproductshape *NoteProductShape) UnstageVoid(stage *Stage) {
	noteproductshape.Unstage(stage)
}

func (noteproductshape *NoteProductShape) StageVoid(stage *Stage) {
	noteproductshape.Stage(stage)
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
	__gong__stage(stage.NoteResourceShapes, stage.NoteResourceShape_stagedOrder, stage.NoteResourceShape_orderStaged, &stage.NoteResourceShapeOrder, stage.NoteResourceShapes_mapString, noteresourceshape, noteresourceshape.Name)
	return noteresourceshape
}

// StagePreserveOrder puts noteresourceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteResourceShapeOrder
// - update stage.NoteResourceShapeOrder accordingly
func (noteresourceshape *NoteResourceShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteResourceShapes, stage.NoteResourceShape_stagedOrder, stage.NoteResourceShape_orderStaged, &stage.NoteResourceShapeOrder, stage.NoteResourceShapes_mapString, noteresourceshape, order, noteresourceshape.Name)
}

// Unstage removes noteresourceshape off the model stage
func (noteresourceshape *NoteResourceShape) Unstage(stage *Stage) *NoteResourceShape {
	__gong__unstage(stage.NoteResourceShapes, stage.NoteResourceShapes_mapString, noteresourceshape, noteresourceshape.Name)
	return noteresourceshape
}

// UnstageVoid removes noteresourceshape off the model stage
func (noteresourceshape *NoteResourceShape) UnstageVoid(stage *Stage) {
	noteresourceshape.Unstage(stage)
}

func (noteresourceshape *NoteResourceShape) StageVoid(stage *Stage) {
	noteresourceshape.Stage(stage)
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

// Stage puts product to the model stage
func (product *Product) Stage(stage *Stage) *Product {
	__gong__stage(stage.Products, stage.Product_stagedOrder, stage.Product_orderStaged, &stage.ProductOrder, stage.Products_mapString, product, product.Name)
	return product
}

// StagePreserveOrder puts product to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductOrder
// - update stage.ProductOrder accordingly
func (product *Product) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Products, stage.Product_stagedOrder, stage.Product_orderStaged, &stage.ProductOrder, stage.Products_mapString, product, order, product.Name)
}

// Unstage removes product off the model stage
func (product *Product) Unstage(stage *Stage) *Product {
	__gong__unstage(stage.Products, stage.Products_mapString, product, product.Name)
	return product
}

// UnstageVoid removes product off the model stage
func (product *Product) UnstageVoid(stage *Stage) {
	product.Unstage(stage)
}

func (product *Product) StageVoid(stage *Stage) {
	product.Stage(stage)
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
	__gong__stage(stage.ProductCompositionShapes, stage.ProductCompositionShape_stagedOrder, stage.ProductCompositionShape_orderStaged, &stage.ProductCompositionShapeOrder, stage.ProductCompositionShapes_mapString, productcompositionshape, productcompositionshape.Name)
	return productcompositionshape
}

// StagePreserveOrder puts productcompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductCompositionShapeOrder
// - update stage.ProductCompositionShapeOrder accordingly
func (productcompositionshape *ProductCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ProductCompositionShapes, stage.ProductCompositionShape_stagedOrder, stage.ProductCompositionShape_orderStaged, &stage.ProductCompositionShapeOrder, stage.ProductCompositionShapes_mapString, productcompositionshape, order, productcompositionshape.Name)
}

// Unstage removes productcompositionshape off the model stage
func (productcompositionshape *ProductCompositionShape) Unstage(stage *Stage) *ProductCompositionShape {
	__gong__unstage(stage.ProductCompositionShapes, stage.ProductCompositionShapes_mapString, productcompositionshape, productcompositionshape.Name)
	return productcompositionshape
}

// UnstageVoid removes productcompositionshape off the model stage
func (productcompositionshape *ProductCompositionShape) UnstageVoid(stage *Stage) {
	productcompositionshape.Unstage(stage)
}

func (productcompositionshape *ProductCompositionShape) StageVoid(stage *Stage) {
	productcompositionshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (productcompositionshape *ProductCompositionShape) GetName() (res string) {
	return productcompositionshape.Name
}

// for satisfaction of GongStruct interface
func (productcompositionshape *ProductCompositionShape) SetName(name string) {
	productcompositionshape.Name = name
}

// Stage puts productreferenceshape to the model stage
func (productreferenceshape *ProductReferenceShape) Stage(stage *Stage) *ProductReferenceShape {
	__gong__stage(stage.ProductReferenceShapes, stage.ProductReferenceShape_stagedOrder, stage.ProductReferenceShape_orderStaged, &stage.ProductReferenceShapeOrder, stage.ProductReferenceShapes_mapString, productreferenceshape, productreferenceshape.Name)
	return productreferenceshape
}

// StagePreserveOrder puts productreferenceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductReferenceShapeOrder
// - update stage.ProductReferenceShapeOrder accordingly
func (productreferenceshape *ProductReferenceShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ProductReferenceShapes, stage.ProductReferenceShape_stagedOrder, stage.ProductReferenceShape_orderStaged, &stage.ProductReferenceShapeOrder, stage.ProductReferenceShapes_mapString, productreferenceshape, order, productreferenceshape.Name)
}

// Unstage removes productreferenceshape off the model stage
func (productreferenceshape *ProductReferenceShape) Unstage(stage *Stage) *ProductReferenceShape {
	__gong__unstage(stage.ProductReferenceShapes, stage.ProductReferenceShapes_mapString, productreferenceshape, productreferenceshape.Name)
	return productreferenceshape
}

// UnstageVoid removes productreferenceshape off the model stage
func (productreferenceshape *ProductReferenceShape) UnstageVoid(stage *Stage) {
	productreferenceshape.Unstage(stage)
}

func (productreferenceshape *ProductReferenceShape) StageVoid(stage *Stage) {
	productreferenceshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (productreferenceshape *ProductReferenceShape) GetName() (res string) {
	return productreferenceshape.Name
}

// for satisfaction of GongStruct interface
func (productreferenceshape *ProductReferenceShape) SetName(name string) {
	productreferenceshape.Name = name
}

// Stage puts productshape to the model stage
func (productshape *ProductShape) Stage(stage *Stage) *ProductShape {
	__gong__stage(stage.ProductShapes, stage.ProductShape_stagedOrder, stage.ProductShape_orderStaged, &stage.ProductShapeOrder, stage.ProductShapes_mapString, productshape, productshape.Name)
	return productshape
}

// StagePreserveOrder puts productshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductShapeOrder
// - update stage.ProductShapeOrder accordingly
func (productshape *ProductShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ProductShapes, stage.ProductShape_stagedOrder, stage.ProductShape_orderStaged, &stage.ProductShapeOrder, stage.ProductShapes_mapString, productshape, order, productshape.Name)
}

// Unstage removes productshape off the model stage
func (productshape *ProductShape) Unstage(stage *Stage) *ProductShape {
	__gong__unstage(stage.ProductShapes, stage.ProductShapes_mapString, productshape, productshape.Name)
	return productshape
}

// UnstageVoid removes productshape off the model stage
func (productshape *ProductShape) UnstageVoid(stage *Stage) {
	productshape.Unstage(stage)
}

func (productshape *ProductShape) StageVoid(stage *Stage) {
	productshape.Stage(stage)
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
	__gong__stage(stage.Resources, stage.Resource_stagedOrder, stage.Resource_orderStaged, &stage.ResourceOrder, stage.Resources_mapString, resource, resource.Name)
	return resource
}

// StagePreserveOrder puts resource to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceOrder
// - update stage.ResourceOrder accordingly
func (resource *Resource) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Resources, stage.Resource_stagedOrder, stage.Resource_orderStaged, &stage.ResourceOrder, stage.Resources_mapString, resource, order, resource.Name)
}

// Unstage removes resource off the model stage
func (resource *Resource) Unstage(stage *Stage) *Resource {
	__gong__unstage(stage.Resources, stage.Resources_mapString, resource, resource.Name)
	return resource
}

// UnstageVoid removes resource off the model stage
func (resource *Resource) UnstageVoid(stage *Stage) {
	resource.Unstage(stage)
}

func (resource *Resource) StageVoid(stage *Stage) {
	resource.Stage(stage)
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
	__gong__stage(stage.ResourceCompositionShapes, stage.ResourceCompositionShape_stagedOrder, stage.ResourceCompositionShape_orderStaged, &stage.ResourceCompositionShapeOrder, stage.ResourceCompositionShapes_mapString, resourcecompositionshape, resourcecompositionshape.Name)
	return resourcecompositionshape
}

// StagePreserveOrder puts resourcecompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceCompositionShapeOrder
// - update stage.ResourceCompositionShapeOrder accordingly
func (resourcecompositionshape *ResourceCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ResourceCompositionShapes, stage.ResourceCompositionShape_stagedOrder, stage.ResourceCompositionShape_orderStaged, &stage.ResourceCompositionShapeOrder, stage.ResourceCompositionShapes_mapString, resourcecompositionshape, order, resourcecompositionshape.Name)
}

// Unstage removes resourcecompositionshape off the model stage
func (resourcecompositionshape *ResourceCompositionShape) Unstage(stage *Stage) *ResourceCompositionShape {
	__gong__unstage(stage.ResourceCompositionShapes, stage.ResourceCompositionShapes_mapString, resourcecompositionshape, resourcecompositionshape.Name)
	return resourcecompositionshape
}

// UnstageVoid removes resourcecompositionshape off the model stage
func (resourcecompositionshape *ResourceCompositionShape) UnstageVoid(stage *Stage) {
	resourcecompositionshape.Unstage(stage)
}

func (resourcecompositionshape *ResourceCompositionShape) StageVoid(stage *Stage) {
	resourcecompositionshape.Stage(stage)
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
	__gong__stage(stage.ResourceShapes, stage.ResourceShape_stagedOrder, stage.ResourceShape_orderStaged, &stage.ResourceShapeOrder, stage.ResourceShapes_mapString, resourceshape, resourceshape.Name)
	return resourceshape
}

// StagePreserveOrder puts resourceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceShapeOrder
// - update stage.ResourceShapeOrder accordingly
func (resourceshape *ResourceShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ResourceShapes, stage.ResourceShape_stagedOrder, stage.ResourceShape_orderStaged, &stage.ResourceShapeOrder, stage.ResourceShapes_mapString, resourceshape, order, resourceshape.Name)
}

// Unstage removes resourceshape off the model stage
func (resourceshape *ResourceShape) Unstage(stage *Stage) *ResourceShape {
	__gong__unstage(stage.ResourceShapes, stage.ResourceShapes_mapString, resourceshape, resourceshape.Name)
	return resourceshape
}

// UnstageVoid removes resourceshape off the model stage
func (resourceshape *ResourceShape) UnstageVoid(stage *Stage) {
	resourceshape.Unstage(stage)
}

func (resourceshape *ResourceShape) StageVoid(stage *Stage) {
	resourceshape.Stage(stage)
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
	__gong__stage(stage.ResourceTaskShapes, stage.ResourceTaskShape_stagedOrder, stage.ResourceTaskShape_orderStaged, &stage.ResourceTaskShapeOrder, stage.ResourceTaskShapes_mapString, resourcetaskshape, resourcetaskshape.Name)
	return resourcetaskshape
}

// StagePreserveOrder puts resourcetaskshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ResourceTaskShapeOrder
// - update stage.ResourceTaskShapeOrder accordingly
func (resourcetaskshape *ResourceTaskShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ResourceTaskShapes, stage.ResourceTaskShape_stagedOrder, stage.ResourceTaskShape_orderStaged, &stage.ResourceTaskShapeOrder, stage.ResourceTaskShapes_mapString, resourcetaskshape, order, resourcetaskshape.Name)
}

// Unstage removes resourcetaskshape off the model stage
func (resourcetaskshape *ResourceTaskShape) Unstage(stage *Stage) *ResourceTaskShape {
	__gong__unstage(stage.ResourceTaskShapes, stage.ResourceTaskShapes_mapString, resourcetaskshape, resourcetaskshape.Name)
	return resourcetaskshape
}

// UnstageVoid removes resourcetaskshape off the model stage
func (resourcetaskshape *ResourceTaskShape) UnstageVoid(stage *Stage) {
	resourcetaskshape.Unstage(stage)
}

func (resourcetaskshape *ResourceTaskShape) StageVoid(stage *Stage) {
	resourcetaskshape.Stage(stage)
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
	__gong__stage(stage.Tasks, stage.Task_stagedOrder, stage.Task_orderStaged, &stage.TaskOrder, stage.Tasks_mapString, task, task.Name)
	return task
}

// StagePreserveOrder puts task to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskOrder
// - update stage.TaskOrder accordingly
func (task *Task) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Tasks, stage.Task_stagedOrder, stage.Task_orderStaged, &stage.TaskOrder, stage.Tasks_mapString, task, order, task.Name)
}

// Unstage removes task off the model stage
func (task *Task) Unstage(stage *Stage) *Task {
	__gong__unstage(stage.Tasks, stage.Tasks_mapString, task, task.Name)
	return task
}

// UnstageVoid removes task off the model stage
func (task *Task) UnstageVoid(stage *Stage) {
	task.Unstage(stage)
}

func (task *Task) StageVoid(stage *Stage) {
	task.Stage(stage)
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
	__gong__stage(stage.TaskCompositionShapes, stage.TaskCompositionShape_stagedOrder, stage.TaskCompositionShape_orderStaged, &stage.TaskCompositionShapeOrder, stage.TaskCompositionShapes_mapString, taskcompositionshape, taskcompositionshape.Name)
	return taskcompositionshape
}

// StagePreserveOrder puts taskcompositionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskCompositionShapeOrder
// - update stage.TaskCompositionShapeOrder accordingly
func (taskcompositionshape *TaskCompositionShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TaskCompositionShapes, stage.TaskCompositionShape_stagedOrder, stage.TaskCompositionShape_orderStaged, &stage.TaskCompositionShapeOrder, stage.TaskCompositionShapes_mapString, taskcompositionshape, order, taskcompositionshape.Name)
}

// Unstage removes taskcompositionshape off the model stage
func (taskcompositionshape *TaskCompositionShape) Unstage(stage *Stage) *TaskCompositionShape {
	__gong__unstage(stage.TaskCompositionShapes, stage.TaskCompositionShapes_mapString, taskcompositionshape, taskcompositionshape.Name)
	return taskcompositionshape
}

// UnstageVoid removes taskcompositionshape off the model stage
func (taskcompositionshape *TaskCompositionShape) UnstageVoid(stage *Stage) {
	taskcompositionshape.Unstage(stage)
}

func (taskcompositionshape *TaskCompositionShape) StageVoid(stage *Stage) {
	taskcompositionshape.Stage(stage)
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
	__gong__stage(stage.TaskGroups, stage.TaskGroup_stagedOrder, stage.TaskGroup_orderStaged, &stage.TaskGroupOrder, stage.TaskGroups_mapString, taskgroup, taskgroup.Name)
	return taskgroup
}

// StagePreserveOrder puts taskgroup to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskGroupOrder
// - update stage.TaskGroupOrder accordingly
func (taskgroup *TaskGroup) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TaskGroups, stage.TaskGroup_stagedOrder, stage.TaskGroup_orderStaged, &stage.TaskGroupOrder, stage.TaskGroups_mapString, taskgroup, order, taskgroup.Name)
}

// Unstage removes taskgroup off the model stage
func (taskgroup *TaskGroup) Unstage(stage *Stage) *TaskGroup {
	__gong__unstage(stage.TaskGroups, stage.TaskGroups_mapString, taskgroup, taskgroup.Name)
	return taskgroup
}

// UnstageVoid removes taskgroup off the model stage
func (taskgroup *TaskGroup) UnstageVoid(stage *Stage) {
	taskgroup.Unstage(stage)
}

func (taskgroup *TaskGroup) StageVoid(stage *Stage) {
	taskgroup.Stage(stage)
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
	__gong__stage(stage.TaskGroupShapes, stage.TaskGroupShape_stagedOrder, stage.TaskGroupShape_orderStaged, &stage.TaskGroupShapeOrder, stage.TaskGroupShapes_mapString, taskgroupshape, taskgroupshape.Name)
	return taskgroupshape
}

// StagePreserveOrder puts taskgroupshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskGroupShapeOrder
// - update stage.TaskGroupShapeOrder accordingly
func (taskgroupshape *TaskGroupShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TaskGroupShapes, stage.TaskGroupShape_stagedOrder, stage.TaskGroupShape_orderStaged, &stage.TaskGroupShapeOrder, stage.TaskGroupShapes_mapString, taskgroupshape, order, taskgroupshape.Name)
}

// Unstage removes taskgroupshape off the model stage
func (taskgroupshape *TaskGroupShape) Unstage(stage *Stage) *TaskGroupShape {
	__gong__unstage(stage.TaskGroupShapes, stage.TaskGroupShapes_mapString, taskgroupshape, taskgroupshape.Name)
	return taskgroupshape
}

// UnstageVoid removes taskgroupshape off the model stage
func (taskgroupshape *TaskGroupShape) UnstageVoid(stage *Stage) {
	taskgroupshape.Unstage(stage)
}

func (taskgroupshape *TaskGroupShape) StageVoid(stage *Stage) {
	taskgroupshape.Stage(stage)
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
	__gong__stage(stage.TaskInputShapes, stage.TaskInputShape_stagedOrder, stage.TaskInputShape_orderStaged, &stage.TaskInputShapeOrder, stage.TaskInputShapes_mapString, taskinputshape, taskinputshape.Name)
	return taskinputshape
}

// StagePreserveOrder puts taskinputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskInputShapeOrder
// - update stage.TaskInputShapeOrder accordingly
func (taskinputshape *TaskInputShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TaskInputShapes, stage.TaskInputShape_stagedOrder, stage.TaskInputShape_orderStaged, &stage.TaskInputShapeOrder, stage.TaskInputShapes_mapString, taskinputshape, order, taskinputshape.Name)
}

// Unstage removes taskinputshape off the model stage
func (taskinputshape *TaskInputShape) Unstage(stage *Stage) *TaskInputShape {
	__gong__unstage(stage.TaskInputShapes, stage.TaskInputShapes_mapString, taskinputshape, taskinputshape.Name)
	return taskinputshape
}

// UnstageVoid removes taskinputshape off the model stage
func (taskinputshape *TaskInputShape) UnstageVoid(stage *Stage) {
	taskinputshape.Unstage(stage)
}

func (taskinputshape *TaskInputShape) StageVoid(stage *Stage) {
	taskinputshape.Stage(stage)
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
	__gong__stage(stage.TaskOutputShapes, stage.TaskOutputShape_stagedOrder, stage.TaskOutputShape_orderStaged, &stage.TaskOutputShapeOrder, stage.TaskOutputShapes_mapString, taskoutputshape, taskoutputshape.Name)
	return taskoutputshape
}

// StagePreserveOrder puts taskoutputshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskOutputShapeOrder
// - update stage.TaskOutputShapeOrder accordingly
func (taskoutputshape *TaskOutputShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TaskOutputShapes, stage.TaskOutputShape_stagedOrder, stage.TaskOutputShape_orderStaged, &stage.TaskOutputShapeOrder, stage.TaskOutputShapes_mapString, taskoutputshape, order, taskoutputshape.Name)
}

// Unstage removes taskoutputshape off the model stage
func (taskoutputshape *TaskOutputShape) Unstage(stage *Stage) *TaskOutputShape {
	__gong__unstage(stage.TaskOutputShapes, stage.TaskOutputShapes_mapString, taskoutputshape, taskoutputshape.Name)
	return taskoutputshape
}

// UnstageVoid removes taskoutputshape off the model stage
func (taskoutputshape *TaskOutputShape) UnstageVoid(stage *Stage) {
	taskoutputshape.Unstage(stage)
}

func (taskoutputshape *TaskOutputShape) StageVoid(stage *Stage) {
	taskoutputshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (taskoutputshape *TaskOutputShape) GetName() (res string) {
	return taskoutputshape.Name
}

// for satisfaction of GongStruct interface
func (taskoutputshape *TaskOutputShape) SetName(name string) {
	taskoutputshape.Name = name
}

// Stage puts taskpredecessorshape to the model stage
func (taskpredecessorshape *TaskPredecessorShape) Stage(stage *Stage) *TaskPredecessorShape {
	__gong__stage(stage.TaskPredecessorShapes, stage.TaskPredecessorShape_stagedOrder, stage.TaskPredecessorShape_orderStaged, &stage.TaskPredecessorShapeOrder, stage.TaskPredecessorShapes_mapString, taskpredecessorshape, taskpredecessorshape.Name)
	return taskpredecessorshape
}

// StagePreserveOrder puts taskpredecessorshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskPredecessorShapeOrder
// - update stage.TaskPredecessorShapeOrder accordingly
func (taskpredecessorshape *TaskPredecessorShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TaskPredecessorShapes, stage.TaskPredecessorShape_stagedOrder, stage.TaskPredecessorShape_orderStaged, &stage.TaskPredecessorShapeOrder, stage.TaskPredecessorShapes_mapString, taskpredecessorshape, order, taskpredecessorshape.Name)
}

// Unstage removes taskpredecessorshape off the model stage
func (taskpredecessorshape *TaskPredecessorShape) Unstage(stage *Stage) *TaskPredecessorShape {
	__gong__unstage(stage.TaskPredecessorShapes, stage.TaskPredecessorShapes_mapString, taskpredecessorshape, taskpredecessorshape.Name)
	return taskpredecessorshape
}

// UnstageVoid removes taskpredecessorshape off the model stage
func (taskpredecessorshape *TaskPredecessorShape) UnstageVoid(stage *Stage) {
	taskpredecessorshape.Unstage(stage)
}

func (taskpredecessorshape *TaskPredecessorShape) StageVoid(stage *Stage) {
	taskpredecessorshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (taskpredecessorshape *TaskPredecessorShape) GetName() (res string) {
	return taskpredecessorshape.Name
}

// for satisfaction of GongStruct interface
func (taskpredecessorshape *TaskPredecessorShape) SetName(name string) {
	taskpredecessorshape.Name = name
}

// Stage puts taskshape to the model stage
func (taskshape *TaskShape) Stage(stage *Stage) *TaskShape {
	__gong__stage(stage.TaskShapes, stage.TaskShape_stagedOrder, stage.TaskShape_orderStaged, &stage.TaskShapeOrder, stage.TaskShapes_mapString, taskshape, taskshape.Name)
	return taskshape
}

// StagePreserveOrder puts taskshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskShapeOrder
// - update stage.TaskShapeOrder accordingly
func (taskshape *TaskShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.TaskShapes, stage.TaskShape_stagedOrder, stage.TaskShape_orderStaged, &stage.TaskShapeOrder, stage.TaskShapes_mapString, taskshape, order, taskshape.Name)
}

// Unstage removes taskshape off the model stage
func (taskshape *TaskShape) Unstage(stage *Stage) *TaskShape {
	__gong__unstage(stage.TaskShapes, stage.TaskShapes_mapString, taskshape, taskshape.Name)
	return taskshape
}

// UnstageVoid removes taskshape off the model stage
func (taskshape *TaskShape) UnstageVoid(stage *Stage) {
	taskshape.Unstage(stage)
}

func (taskshape *TaskShape) StageVoid(stage *Stage) {
	taskshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (taskshape *TaskShape) GetName() (res string) {
	return taskshape.Name
}

// for satisfaction of GongStruct interface
func (taskshape *TaskShape) SetName(name string) {
	taskshape.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Diagrams, &stage.Diagrams_mapString, &stage.Diagram_stagedOrder, &stage.DiagramOrder)

	__gong__resetStageType(&stage.Librarys, &stage.Librarys_mapString, &stage.Library_stagedOrder, &stage.LibraryOrder)

	__gong__resetStageType(&stage.Notes, &stage.Notes_mapString, &stage.Note_stagedOrder, &stage.NoteOrder)

	__gong__resetStageType(&stage.NoteProductShapes, &stage.NoteProductShapes_mapString, &stage.NoteProductShape_stagedOrder, &stage.NoteProductShapeOrder)

	__gong__resetStageType(&stage.NoteResourceShapes, &stage.NoteResourceShapes_mapString, &stage.NoteResourceShape_stagedOrder, &stage.NoteResourceShapeOrder)

	__gong__resetStageType(&stage.NoteShapes, &stage.NoteShapes_mapString, &stage.NoteShape_stagedOrder, &stage.NoteShapeOrder)

	__gong__resetStageType(&stage.NoteTaskShapes, &stage.NoteTaskShapes_mapString, &stage.NoteTaskShape_stagedOrder, &stage.NoteTaskShapeOrder)

	__gong__resetStageType(&stage.Products, &stage.Products_mapString, &stage.Product_stagedOrder, &stage.ProductOrder)

	__gong__resetStageType(&stage.ProductCompositionShapes, &stage.ProductCompositionShapes_mapString, &stage.ProductCompositionShape_stagedOrder, &stage.ProductCompositionShapeOrder)

	__gong__resetStageType(&stage.ProductReferenceShapes, &stage.ProductReferenceShapes_mapString, &stage.ProductReferenceShape_stagedOrder, &stage.ProductReferenceShapeOrder)

	__gong__resetStageType(&stage.ProductShapes, &stage.ProductShapes_mapString, &stage.ProductShape_stagedOrder, &stage.ProductShapeOrder)

	__gong__resetStageType(&stage.Resources, &stage.Resources_mapString, &stage.Resource_stagedOrder, &stage.ResourceOrder)

	__gong__resetStageType(&stage.ResourceCompositionShapes, &stage.ResourceCompositionShapes_mapString, &stage.ResourceCompositionShape_stagedOrder, &stage.ResourceCompositionShapeOrder)

	__gong__resetStageType(&stage.ResourceShapes, &stage.ResourceShapes_mapString, &stage.ResourceShape_stagedOrder, &stage.ResourceShapeOrder)

	__gong__resetStageType(&stage.ResourceTaskShapes, &stage.ResourceTaskShapes_mapString, &stage.ResourceTaskShape_stagedOrder, &stage.ResourceTaskShapeOrder)

	__gong__resetStageType(&stage.Tasks, &stage.Tasks_mapString, &stage.Task_stagedOrder, &stage.TaskOrder)

	__gong__resetStageType(&stage.TaskCompositionShapes, &stage.TaskCompositionShapes_mapString, &stage.TaskCompositionShape_stagedOrder, &stage.TaskCompositionShapeOrder)

	__gong__resetStageType(&stage.TaskGroups, &stage.TaskGroups_mapString, &stage.TaskGroup_stagedOrder, &stage.TaskGroupOrder)

	__gong__resetStageType(&stage.TaskGroupShapes, &stage.TaskGroupShapes_mapString, &stage.TaskGroupShape_stagedOrder, &stage.TaskGroupShapeOrder)

	__gong__resetStageType(&stage.TaskInputShapes, &stage.TaskInputShapes_mapString, &stage.TaskInputShape_stagedOrder, &stage.TaskInputShapeOrder)

	__gong__resetStageType(&stage.TaskOutputShapes, &stage.TaskOutputShapes_mapString, &stage.TaskOutputShape_stagedOrder, &stage.TaskOutputShapeOrder)

	__gong__resetStageType(&stage.TaskPredecessorShapes, &stage.TaskPredecessorShapes_mapString, &stage.TaskPredecessorShape_stagedOrder, &stage.TaskPredecessorShapeOrder)

	__gong__resetStageType(&stage.TaskShapes, &stage.TaskShapes_mapString, &stage.TaskShape_stagedOrder, &stage.TaskShapeOrder)

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
	case *ProductReferenceShape:
		return any(stage.ProductReferenceShapes_mapString).(map[string]Type)
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
	case *TaskPredecessorShape:
		return any(stage.TaskPredecessorShapes_mapString).(map[string]Type)
	case *TaskShape:
		return any(stage.TaskShapes_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
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
	case *ProductReferenceShape:
		return any(&stage.ProductReferenceShapes).(*map[Type]struct{})
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
	case *TaskPredecessorShape:
		return any(&stage.TaskPredecessorShapes).(*map[Type]struct{})
	case *TaskShape:
		return any(&stage.TaskShapes).(*map[Type]struct{})
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
	case Diagram:
		return any(&Diagram{
			Product_Shapes: []*ProductShape{{Name: "Product_Shapes"}},
			ProductsWhoseNodeIsExpanded: []*Product{{Name: "ProductsWhoseNodeIsExpanded"}},
			ProductComposition_Shapes: []*ProductCompositionShape{{Name: "ProductComposition_Shapes"}},
			ProductReference_Shapes: []*ProductReferenceShape{{Name: "ProductReference_Shapes"}},
			Task_Shapes: []*TaskShape{{Name: "Task_Shapes"}},
			TasksWhoseNodeIsExpanded: []*Task{{Name: "TasksWhoseNodeIsExpanded"}},
			TasksWhoseInputNodeIsExpanded: []*Task{{Name: "TasksWhoseInputNodeIsExpanded"}},
			TasksWhoseOutputNodeIsExpanded: []*Task{{Name: "TasksWhoseOutputNodeIsExpanded"}},
			TasksWhosePredecessorNodeIsExpanded: []*Task{{Name: "TasksWhosePredecessorNodeIsExpanded"}},
			TaskGroupShapes: []*TaskGroupShape{{Name: "TaskGroupShapes"}},
			TaskGroupsWhoseNodeIsExpanded: []*TaskGroup{{Name: "TaskGroupsWhoseNodeIsExpanded"}},
			TaskComposition_Shapes: []*TaskCompositionShape{{Name: "TaskComposition_Shapes"}},
			TaskInputShapes: []*TaskInputShape{{Name: "TaskInputShapes"}},
			TaskOutputShapes: []*TaskOutputShape{{Name: "TaskOutputShapes"}},
			TaskPredecessorShapes: []*TaskPredecessorShape{{Name: "TaskPredecessorShapes"}},
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			NoteProductShapes: []*NoteProductShape{{Name: "NoteProductShapes"}},
			NoteTaskShapes: []*NoteTaskShape{{Name: "NoteTaskShapes"}},
			NoteResourceShapes: []*NoteResourceShape{{Name: "NoteResourceShapes"}},
			Resource_Shapes: []*ResourceShape{{Name: "Resource_Shapes"}},
			ResourcesWhoseNodeIsExpanded: []*Resource{{Name: "ResourcesWhoseNodeIsExpanded"}},
			ResourceComposition_Shapes: []*ResourceCompositionShape{{Name: "ResourceComposition_Shapes"}},
			ResourceTaskShapes: []*ResourceTaskShape{{Name: "ResourceTaskShapes"}},
		}).(*Type)
	case Library:
		return any(&Library{
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			RootProducts: []*Product{{Name: "RootProducts"}},
			RootTasks: []*Task{{Name: "RootTasks"}},
			RootTaskGroups: []*TaskGroup{{Name: "RootTaskGroups"}},
			RootResources: []*Resource{{Name: "RootResources"}},
			Notes: []*Note{{Name: "Notes"}},
			Diagrams: []*Diagram{{Name: "Diagrams"}},
		}).(*Type)
	case Note:
		return any(&Note{
			Products: []*Product{{Name: "Products"}},
			Tasks: []*Task{{Name: "Tasks"}},
			Resources: []*Resource{{Name: "Resources"}},
		}).(*Type)
	case NoteProductShape:
		return any(&NoteProductShape{
			Note: &Note{Name: "Note"},
			Product: &Product{Name: "Product"},
		}).(*Type)
	case NoteResourceShape:
		return any(&NoteResourceShape{
			Note: &Note{Name: "Note"},
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			Note: &Note{Name: "Note"},
		}).(*Type)
	case NoteTaskShape:
		return any(&NoteTaskShape{
			Note: &Note{Name: "Note"},
			Task: &Task{Name: "Task"},
		}).(*Type)
	case Product:
		return any(&Product{
			SubProducts: []*Product{{Name: "SubProducts"}},
			ReferencedProduct: &Product{Name: "ReferencedProduct"},
		}).(*Type)
	case ProductCompositionShape:
		return any(&ProductCompositionShape{
			Product: &Product{Name: "Product"},
		}).(*Type)
	case ProductReferenceShape:
		return any(&ProductReferenceShape{
			Product: &Product{Name: "Product"},
			ReferencedProduct: &Product{Name: "ReferencedProduct"},
		}).(*Type)
	case ProductShape:
		return any(&ProductShape{
			Product: &Product{Name: "Product"},
		}).(*Type)
	case Resource:
		return any(&Resource{
			Tasks: []*Task{{Name: "Tasks"}},
			SubResources: []*Resource{{Name: "SubResources"}},
			ReferencedResource: &Resource{Name: "ReferencedResource"},
		}).(*Type)
	case ResourceCompositionShape:
		return any(&ResourceCompositionShape{
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
	case ResourceShape:
		return any(&ResourceShape{
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
	case ResourceTaskShape:
		return any(&ResourceTaskShape{
			Resource: &Resource{Name: "Resource"},
			Task: &Task{Name: "Task"},
		}).(*Type)
	case Task:
		return any(&Task{
			Predecessors: []*Task{{Name: "Predecessors"}},
			Inputs: []*Product{{Name: "Inputs"}},
			Outputs: []*Product{{Name: "Outputs"}},
			SubTasks: []*Task{{Name: "SubTasks"}},
			TaskGroupsToDisplay: []*TaskGroup{{Name: "TaskGroupsToDisplay"}},
			ReferencedTask: &Task{Name: "ReferencedTask"},
		}).(*Type)
	case TaskCompositionShape:
		return any(&TaskCompositionShape{
			Task: &Task{Name: "Task"},
		}).(*Type)
	case TaskGroup:
		return any(&TaskGroup{
			Tasks: []*Task{{Name: "Tasks"}},
		}).(*Type)
	case TaskGroupShape:
		return any(&TaskGroupShape{
			TaskGroup: &TaskGroup{Name: "TaskGroup"},
		}).(*Type)
	case TaskInputShape:
		return any(&TaskInputShape{
			Product: &Product{Name: "Product"},
			Task: &Task{Name: "Task"},
		}).(*Type)
	case TaskOutputShape:
		return any(&TaskOutputShape{
			Task: &Task{Name: "Task"},
			Product: &Product{Name: "Product"},
		}).(*Type)
	case TaskPredecessorShape:
		return any(&TaskPredecessorShape{
			Predecessor: &Task{Name: "Predecessor"},
			Task: &Task{Name: "Task"},
		}).(*Type)
	case TaskShape:
		return any(&TaskShape{
			Task: &Task{Name: "Task"},
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
	// reverse maps of direct associations of ProductReferenceShape
	case ProductReferenceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Product":
			res := make(map[*Product][]*ProductReferenceShape)
			for productreferenceshape := range stage.ProductReferenceShapes {
				if productreferenceshape.Product != nil {
					product_ := productreferenceshape.Product
					var productreferenceshapes []*ProductReferenceShape
					_, ok := res[product_]
					if ok {
						productreferenceshapes = res[product_]
					} else {
						productreferenceshapes = make([]*ProductReferenceShape, 0)
					}
					productreferenceshapes = append(productreferenceshapes, productreferenceshape)
					res[product_] = productreferenceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "ReferencedProduct":
			res := make(map[*Product][]*ProductReferenceShape)
			for productreferenceshape := range stage.ProductReferenceShapes {
				if productreferenceshape.ReferencedProduct != nil {
					product_ := productreferenceshape.ReferencedProduct
					var productreferenceshapes []*ProductReferenceShape
					_, ok := res[product_]
					if ok {
						productreferenceshapes = res[product_]
					} else {
						productreferenceshapes = make([]*ProductReferenceShape, 0)
					}
					productreferenceshapes = append(productreferenceshapes, productreferenceshape)
					res[product_] = productreferenceshapes
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
	// reverse maps of direct associations of TaskPredecessorShape
	case TaskPredecessorShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Predecessor":
			res := make(map[*Task][]*TaskPredecessorShape)
			for taskpredecessorshape := range stage.TaskPredecessorShapes {
				if taskpredecessorshape.Predecessor != nil {
					task_ := taskpredecessorshape.Predecessor
					var taskpredecessorshapes []*TaskPredecessorShape
					_, ok := res[task_]
					if ok {
						taskpredecessorshapes = res[task_]
					} else {
						taskpredecessorshapes = make([]*TaskPredecessorShape, 0)
					}
					taskpredecessorshapes = append(taskpredecessorshapes, taskpredecessorshape)
					res[task_] = taskpredecessorshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Task":
			res := make(map[*Task][]*TaskPredecessorShape)
			for taskpredecessorshape := range stage.TaskPredecessorShapes {
				if taskpredecessorshape.Task != nil {
					task_ := taskpredecessorshape.Task
					var taskpredecessorshapes []*TaskPredecessorShape
					_, ok := res[task_]
					if ok {
						taskpredecessorshapes = res[task_]
					} else {
						taskpredecessorshapes = make([]*TaskPredecessorShape, 0)
					}
					taskpredecessorshapes = append(taskpredecessorshapes, taskpredecessorshape)
					res[task_] = taskpredecessorshapes
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

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
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
		case "ProductReference_Shapes":
			res := make(map[*ProductReferenceShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, productreferenceshape_ := range diagram.ProductReference_Shapes {
					res[productreferenceshape_] = append(res[productreferenceshape_], diagram)
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
		case "TasksWhosePredecessorNodeIsExpanded":
			res := make(map[*Task][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, task_ := range diagram.TasksWhosePredecessorNodeIsExpanded {
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
		case "TaskPredecessorShapes":
			res := make(map[*TaskPredecessorShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, taskpredecessorshape_ := range diagram.TaskPredecessorShapes {
					res[taskpredecessorshape_] = append(res[taskpredecessorshape_], diagram)
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
	// reverse maps of direct associations of ProductReferenceShape
	case ProductReferenceShape:
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
		case "SubTasks":
			res := make(map[*Task][]*Task)
			for task := range stage.Tasks {
				for _, task_ := range task.SubTasks {
					res[task_] = append(res[task_], task)
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
	// reverse maps of direct associations of TaskPredecessorShape
	case TaskPredecessorShape:
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

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *Diagram:
		res = any(new(Diagram)).(Type)
	case *Library:
		res = any(new(Library)).(Type)
	case *Note:
		res = any(new(Note)).(Type)
	case *NoteProductShape:
		res = any(new(NoteProductShape)).(Type)
	case *NoteResourceShape:
		res = any(new(NoteResourceShape)).(Type)
	case *NoteShape:
		res = any(new(NoteShape)).(Type)
	case *NoteTaskShape:
		res = any(new(NoteTaskShape)).(Type)
	case *Product:
		res = any(new(Product)).(Type)
	case *ProductCompositionShape:
		res = any(new(ProductCompositionShape)).(Type)
	case *ProductReferenceShape:
		res = any(new(ProductReferenceShape)).(Type)
	case *ProductShape:
		res = any(new(ProductShape)).(Type)
	case *Resource:
		res = any(new(Resource)).(Type)
	case *ResourceCompositionShape:
		res = any(new(ResourceCompositionShape)).(Type)
	case *ResourceShape:
		res = any(new(ResourceShape)).(Type)
	case *ResourceTaskShape:
		res = any(new(ResourceTaskShape)).(Type)
	case *Task:
		res = any(new(Task)).(Type)
	case *TaskCompositionShape:
		res = any(new(TaskCompositionShape)).(Type)
	case *TaskGroup:
		res = any(new(TaskGroup)).(Type)
	case *TaskGroupShape:
		res = any(new(TaskGroupShape)).(Type)
	case *TaskInputShape:
		res = any(new(TaskInputShape)).(Type)
	case *TaskOutputShape:
		res = any(new(TaskOutputShape)).(Type)
	case *TaskPredecessorShape:
		res = any(new(TaskPredecessorShape)).(Type)
	case *TaskShape:
		res = any(new(TaskShape)).(Type)
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
	case *ProductReferenceShape:
		res = "ProductReferenceShape"
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
	case *TaskPredecessorShape:
		res = "TaskPredecessorShape"
	case *TaskShape:
		res = "TaskShape"
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
	case *ProductReferenceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "ProductReference_Shapes"
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
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TasksWhosePredecessorNodeIsExpanded"
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
	case *TaskPredecessorShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "TaskPredecessorShapes"
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

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
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
			Name:               "DefaultBoxWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DefaultBoxHeigth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DateFormat",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "DrawVerticalTimeLines",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HideWeekendsPeriod",
			GongFieldValueType: GongFieldValueTypeBool,
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
			Name:               "IsShowPrefix",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsInAutoLayoutMode",
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
			Name:                 "ProductReference_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProductReferenceShape",
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
			Name:                 "TasksWhosePredecessorNodeIsExpanded",
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
			Name:                 "TaskPredecessorShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskPredecessorShape",
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

func (productreferenceshape *ProductReferenceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ReferencedProduct",
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
			Name:               "IsShowType",
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

func (resource *Resource) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "IsAllDay",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsMilestone",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Predecessors",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "DependencyType",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DependencyTypeEnum",
		},
		{
			Name:               "DependencyDurationYears",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DependencyDurationMonths",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DependencyDurationWeeks",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DependencyDurationDays",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DependencyDurationHours",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DurationYears",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DurationMonths",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DurationWeeks",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DurationDays",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DurationHours",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsEndDateComputedFromDuration",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Inputs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:                 "Outputs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
		{
			Name:                 "SubTasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
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
			Name:               "IsInputsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsOutputsNodeExpanded",
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

func (taskpredecessorshape *TaskPredecessorShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Predecessor",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
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
			Name:               "VerticalOffset",
			GongFieldValueType: GongFieldValueTypeFloat,
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
func (diagram *Diagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagram.Name
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagram.DefaultBoxWidth)
		res.valueFloat = diagram.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagram.DefaultBoxHeigth)
		res.valueFloat = diagram.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DateFormat":
		res.valueString = diagram.DateFormat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagram.Width)
		res.valueFloat = diagram.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagram.Height)
		res.valueFloat = diagram.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
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
			days := __gong__abs(int(int(diagram.ComputedDuration.Hours()) / 24))
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
	case "DrawVerticalTimeLines":
		res.valueString = fmt.Sprintf("%t", diagram.DrawVerticalTimeLines)
		res.valueBool = diagram.DrawVerticalTimeLines
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HideWeekendsPeriod":
		res.valueString = fmt.Sprintf("%t", diagram.HideWeekendsPeriod)
		res.valueBool = diagram.HideWeekendsPeriod
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagram.IsShowPrefix)
		res.valueBool = diagram.IsShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInAutoLayoutMode":
		res.valueString = fmt.Sprintf("%t", diagram.IsInAutoLayoutMode)
		res.valueBool = diagram.IsInAutoLayoutMode
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
	case "ProductReference_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.ProductReference_Shapes {
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
	case "TasksWhosePredecessorNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TasksWhosePredecessorNodeIsExpanded {
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
	case "TaskPredecessorShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.TaskPredecessorShapes {
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
	case "ComputedPrefix":
		res.valueString = product.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", product.IsExpanded)
		res.valueBool = product.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := product.LayoutDirection
		res.valueString = enum.ToCodeString()
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

func (productreferenceshape *ProductReferenceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = productreferenceshape.Name
	case "Product":
		res.GongFieldValueType = GongFieldValueTypePointer
		if productreferenceshape.Product != nil {
			res.valueString = productreferenceshape.Product.Name
			res.ids = productreferenceshape.Product.GongGetUUID(stage)
		}
	case "ReferencedProduct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if productreferenceshape.ReferencedProduct != nil {
			res.valueString = productreferenceshape.ReferencedProduct.Name
			res.ids = productreferenceshape.ReferencedProduct.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", productreferenceshape.StartRatio)
		res.valueFloat = productreferenceshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", productreferenceshape.EndRatio)
		res.valueFloat = productreferenceshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := productreferenceshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := productreferenceshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", productreferenceshape.CornerOffsetRatio)
		res.valueFloat = productreferenceshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", productreferenceshape.IsHidden)
		res.valueBool = productreferenceshape.IsHidden
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
	case "IsShowType":
		res.valueString = fmt.Sprintf("%t", productshape.IsShowType)
		res.valueBool = productshape.IsShowType
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "Description":
		res.valueString = task.Description
	case "Start":
		res.valueString = task.Start.String()
	case "End":
		res.valueString = task.End.String()
	case "IsAllDay":
		res.valueString = fmt.Sprintf("%t", task.IsAllDay)
		res.valueBool = task.IsAllDay
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsMilestone":
		res.valueString = fmt.Sprintf("%t", task.IsMilestone)
		res.valueBool = task.IsMilestone
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
	case "DependencyType":
		enum := task.DependencyType
		res.valueString = enum.ToCodeString()
	case "DependencyDurationYears":
		res.valueString = fmt.Sprintf("%f", task.DependencyDurationYears)
		res.valueFloat = task.DependencyDurationYears
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DependencyDurationMonths":
		res.valueString = fmt.Sprintf("%f", task.DependencyDurationMonths)
		res.valueFloat = task.DependencyDurationMonths
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DependencyDurationWeeks":
		res.valueString = fmt.Sprintf("%f", task.DependencyDurationWeeks)
		res.valueFloat = task.DependencyDurationWeeks
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DependencyDurationDays":
		res.valueString = fmt.Sprintf("%f", task.DependencyDurationDays)
		res.valueFloat = task.DependencyDurationDays
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DependencyDurationHours":
		res.valueString = fmt.Sprintf("%f", task.DependencyDurationHours)
		res.valueFloat = task.DependencyDurationHours
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DurationYears":
		res.valueString = fmt.Sprintf("%f", task.DurationYears)
		res.valueFloat = task.DurationYears
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DurationMonths":
		res.valueString = fmt.Sprintf("%f", task.DurationMonths)
		res.valueFloat = task.DurationMonths
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DurationWeeks":
		res.valueString = fmt.Sprintf("%f", task.DurationWeeks)
		res.valueFloat = task.DurationWeeks
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DurationDays":
		res.valueString = fmt.Sprintf("%f", task.DurationDays)
		res.valueFloat = task.DurationDays
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DurationHours":
		res.valueString = fmt.Sprintf("%f", task.DurationHours)
		res.valueFloat = task.DurationHours
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsEndDateComputedFromDuration":
		res.valueString = fmt.Sprintf("%t", task.IsEndDateComputedFromDuration)
		res.valueBool = task.IsEndDateComputedFromDuration
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "IsInputsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsInputsNodeExpanded)
		res.valueBool = task.IsInputsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsOutputsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsOutputsNodeExpanded)
		res.valueBool = task.IsOutputsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = task.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsExpanded)
		res.valueBool = task.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := task.LayoutDirection
		res.valueString = enum.ToCodeString()
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

func (taskpredecessorshape *TaskPredecessorShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = taskpredecessorshape.Name
	case "Predecessor":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskpredecessorshape.Predecessor != nil {
			res.valueString = taskpredecessorshape.Predecessor.Name
			res.ids = taskpredecessorshape.Predecessor.GongGetUUID(stage)
		}
	case "Task":
		res.GongFieldValueType = GongFieldValueTypePointer
		if taskpredecessorshape.Task != nil {
			res.valueString = taskpredecessorshape.Task.Name
			res.ids = taskpredecessorshape.Task.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", taskpredecessorshape.StartRatio)
		res.valueFloat = taskpredecessorshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", taskpredecessorshape.EndRatio)
		res.valueFloat = taskpredecessorshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := taskpredecessorshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := taskpredecessorshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", taskpredecessorshape.CornerOffsetRatio)
		res.valueFloat = taskpredecessorshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", taskpredecessorshape.IsHidden)
		res.valueBool = taskpredecessorshape.IsHidden
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
	case "VerticalOffset":
		res.valueString = fmt.Sprintf("%f", taskshape.VerticalOffset)
		res.valueFloat = taskshape.VerticalOffset
		res.GongFieldValueType = GongFieldValueTypeFloat
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
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

func (productreferenceshape *ProductReferenceShape) GongGetGongstructName() string {
	return "ProductReferenceShape"
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

func (taskpredecessorshape *TaskPredecessorShape) GongGetGongstructName() string {
	return "TaskPredecessorShape"
}

func (taskshape *TaskShape) GongGetGongstructName() string {
	return "TaskShape"
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
	__gong__rebuildMapString(stage.Diagrams, &stage.Diagrams_mapString)

	__gong__rebuildMapString(stage.Librarys, &stage.Librarys_mapString)

	__gong__rebuildMapString(stage.Notes, &stage.Notes_mapString)

	__gong__rebuildMapString(stage.NoteProductShapes, &stage.NoteProductShapes_mapString)

	__gong__rebuildMapString(stage.NoteResourceShapes, &stage.NoteResourceShapes_mapString)

	__gong__rebuildMapString(stage.NoteShapes, &stage.NoteShapes_mapString)

	__gong__rebuildMapString(stage.NoteTaskShapes, &stage.NoteTaskShapes_mapString)

	__gong__rebuildMapString(stage.Products, &stage.Products_mapString)

	__gong__rebuildMapString(stage.ProductCompositionShapes, &stage.ProductCompositionShapes_mapString)

	__gong__rebuildMapString(stage.ProductReferenceShapes, &stage.ProductReferenceShapes_mapString)

	__gong__rebuildMapString(stage.ProductShapes, &stage.ProductShapes_mapString)

	__gong__rebuildMapString(stage.Resources, &stage.Resources_mapString)

	__gong__rebuildMapString(stage.ResourceCompositionShapes, &stage.ResourceCompositionShapes_mapString)

	__gong__rebuildMapString(stage.ResourceShapes, &stage.ResourceShapes_mapString)

	__gong__rebuildMapString(stage.ResourceTaskShapes, &stage.ResourceTaskShapes_mapString)

	__gong__rebuildMapString(stage.Tasks, &stage.Tasks_mapString)

	__gong__rebuildMapString(stage.TaskCompositionShapes, &stage.TaskCompositionShapes_mapString)

	__gong__rebuildMapString(stage.TaskGroups, &stage.TaskGroups_mapString)

	__gong__rebuildMapString(stage.TaskGroupShapes, &stage.TaskGroupShapes_mapString)

	__gong__rebuildMapString(stage.TaskInputShapes, &stage.TaskInputShapes_mapString)

	__gong__rebuildMapString(stage.TaskOutputShapes, &stage.TaskOutputShapes_mapString)

	__gong__rebuildMapString(stage.TaskPredecessorShapes, &stage.TaskPredecessorShapes_mapString)

	__gong__rebuildMapString(stage.TaskShapes, &stage.TaskShapes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
