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

	floss_go "github.com/fullstack-lang/gong/dsm/floss/go"
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
	ProbeLoadSuffix                  = ":load of the probe"
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

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeLoadSuffix
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
	CompareAnalysiss                map[*CompareAnalysis]struct{}
	CompareAnalysiss_instance       map[*CompareAnalysis]*CompareAnalysis
	CompareAnalysiss_mapString      map[string]*CompareAnalysis
	CompareAnalysisOrder            uint
	CompareAnalysis_stagedOrder     map[*CompareAnalysis]uint
	CompareAnalysis_orderStaged     map[uint]*CompareAnalysis
	CompareAnalysiss_reference      map[*CompareAnalysis]*CompareAnalysis
	CompareAnalysiss_referenceOrder map[*CompareAnalysis]uint

	// insertion point for slice of pointers maps
	CompareAnalysis_DiagramFlossEquations_reverseMap map[*DiagramFlossEquation]*CompareAnalysis

	CompareAnalysis_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap map[*DiagramFlossEquation]*CompareAnalysis

	OnAfterCompareAnalysisCreateCallback OnAfterCreateInterface[CompareAnalysis]
	OnAfterCompareAnalysisUpdateCallback OnAfterUpdateInterface[CompareAnalysis]
	OnAfterCompareAnalysisDeleteCallback OnAfterDeleteInterface[CompareAnalysis]
	OnAfterCompareAnalysisReadCallback   OnAfterReadInterface[CompareAnalysis]

	Complexitys                map[*Complexity]struct{}
	Complexitys_instance       map[*Complexity]*Complexity
	Complexitys_mapString      map[string]*Complexity
	ComplexityOrder            uint
	Complexity_stagedOrder     map[*Complexity]uint
	Complexity_orderStaged     map[uint]*Complexity
	Complexitys_reference      map[*Complexity]*Complexity
	Complexitys_referenceOrder map[*Complexity]uint

	// insertion point for slice of pointers maps
	OnAfterComplexityCreateCallback OnAfterCreateInterface[Complexity]
	OnAfterComplexityUpdateCallback OnAfterUpdateInterface[Complexity]
	OnAfterComplexityDeleteCallback OnAfterDeleteInterface[Complexity]
	OnAfterComplexityReadCallback   OnAfterReadInterface[Complexity]

	DiagramFlossEquations                map[*DiagramFlossEquation]struct{}
	DiagramFlossEquations_instance       map[*DiagramFlossEquation]*DiagramFlossEquation
	DiagramFlossEquations_mapString      map[string]*DiagramFlossEquation
	DiagramFlossEquationOrder            uint
	DiagramFlossEquation_stagedOrder     map[*DiagramFlossEquation]uint
	DiagramFlossEquation_orderStaged     map[uint]*DiagramFlossEquation
	DiagramFlossEquations_reference      map[*DiagramFlossEquation]*DiagramFlossEquation
	DiagramFlossEquations_referenceOrder map[*DiagramFlossEquation]uint

	// insertion point for slice of pointers maps
	DiagramFlossEquation_Note_Shapes_reverseMap map[*NoteShape]*DiagramFlossEquation

	DiagramFlossEquation_NoteComplexityShapes_reverseMap map[*NoteComplexityShape]*DiagramFlossEquation

	DiagramFlossEquation_NotePerformanceShapes_reverseMap map[*NotePerformanceShape]*DiagramFlossEquation

	DiagramFlossEquation_NoteEffortShapes_reverseMap map[*NoteEffortShape]*DiagramFlossEquation

	DiagramFlossEquation_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*DiagramFlossEquation

	DiagramFlossEquation_ComplexitysWhoseNodeIsExpanded_reverseMap map[*Complexity]*DiagramFlossEquation

	DiagramFlossEquation_PerformancesWhoseNodeIsExpanded_reverseMap map[*Performance]*DiagramFlossEquation

	DiagramFlossEquation_EffortsWhoseNodeIsExpanded_reverseMap map[*Effort]*DiagramFlossEquation

	OnAfterDiagramFlossEquationCreateCallback OnAfterCreateInterface[DiagramFlossEquation]
	OnAfterDiagramFlossEquationUpdateCallback OnAfterUpdateInterface[DiagramFlossEquation]
	OnAfterDiagramFlossEquationDeleteCallback OnAfterDeleteInterface[DiagramFlossEquation]
	OnAfterDiagramFlossEquationReadCallback   OnAfterReadInterface[DiagramFlossEquation]

	Efforts                map[*Effort]struct{}
	Efforts_instance       map[*Effort]*Effort
	Efforts_mapString      map[string]*Effort
	EffortOrder            uint
	Effort_stagedOrder     map[*Effort]uint
	Effort_orderStaged     map[uint]*Effort
	Efforts_reference      map[*Effort]*Effort
	Efforts_referenceOrder map[*Effort]uint

	// insertion point for slice of pointers maps
	OnAfterEffortCreateCallback OnAfterCreateInterface[Effort]
	OnAfterEffortUpdateCallback OnAfterUpdateInterface[Effort]
	OnAfterEffortDeleteCallback OnAfterDeleteInterface[Effort]
	OnAfterEffortReadCallback   OnAfterReadInterface[Effort]

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

	Library_RootSystems_reverseMap map[*System]*Library

	Library_RootComplexitys_reverseMap map[*Complexity]*Library

	Library_RootPerformances_reverseMap map[*Performance]*Library

	Library_RootEfforts_reverseMap map[*Effort]*Library

	Library_RootCompareAnalysis_reverseMap map[*CompareAnalysis]*Library

	Library_RootNotes_reverseMap map[*Note]*Library

	Library_SubLibrariesWhoseNodeIsExpanded_reverseMap map[*Library]*Library

	Library_SystemsWhoseNodeIsExpanded_reverseMap map[*System]*Library

	Library_ComplexitysWhoseNodeIsExpanded_reverseMap map[*Complexity]*Library

	Library_PerformancesWhoseNodeIsExpanded_reverseMap map[*Performance]*Library

	Library_EffortsWhoseNodeIsExpanded_reverseMap map[*Effort]*Library

	Library_CompareAnalysisWhoseNodeIsExpanded_reverseMap map[*CompareAnalysis]*Library

	Library_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*Library

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
	Note_Complexities_reverseMap map[*Complexity]*Note

	Note_Performances_reverseMap map[*Performance]*Note

	Note_Efforts_reverseMap map[*Effort]*Note

	OnAfterNoteCreateCallback OnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback OnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback OnAfterDeleteInterface[Note]
	OnAfterNoteReadCallback   OnAfterReadInterface[Note]

	NoteComplexityShapes                map[*NoteComplexityShape]struct{}
	NoteComplexityShapes_instance       map[*NoteComplexityShape]*NoteComplexityShape
	NoteComplexityShapes_mapString      map[string]*NoteComplexityShape
	NoteComplexityShapeOrder            uint
	NoteComplexityShape_stagedOrder     map[*NoteComplexityShape]uint
	NoteComplexityShape_orderStaged     map[uint]*NoteComplexityShape
	NoteComplexityShapes_reference      map[*NoteComplexityShape]*NoteComplexityShape
	NoteComplexityShapes_referenceOrder map[*NoteComplexityShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteComplexityShapeCreateCallback OnAfterCreateInterface[NoteComplexityShape]
	OnAfterNoteComplexityShapeUpdateCallback OnAfterUpdateInterface[NoteComplexityShape]
	OnAfterNoteComplexityShapeDeleteCallback OnAfterDeleteInterface[NoteComplexityShape]
	OnAfterNoteComplexityShapeReadCallback   OnAfterReadInterface[NoteComplexityShape]

	NoteEffortShapes                map[*NoteEffortShape]struct{}
	NoteEffortShapes_instance       map[*NoteEffortShape]*NoteEffortShape
	NoteEffortShapes_mapString      map[string]*NoteEffortShape
	NoteEffortShapeOrder            uint
	NoteEffortShape_stagedOrder     map[*NoteEffortShape]uint
	NoteEffortShape_orderStaged     map[uint]*NoteEffortShape
	NoteEffortShapes_reference      map[*NoteEffortShape]*NoteEffortShape
	NoteEffortShapes_referenceOrder map[*NoteEffortShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteEffortShapeCreateCallback OnAfterCreateInterface[NoteEffortShape]
	OnAfterNoteEffortShapeUpdateCallback OnAfterUpdateInterface[NoteEffortShape]
	OnAfterNoteEffortShapeDeleteCallback OnAfterDeleteInterface[NoteEffortShape]
	OnAfterNoteEffortShapeReadCallback   OnAfterReadInterface[NoteEffortShape]

	NotePerformanceShapes                map[*NotePerformanceShape]struct{}
	NotePerformanceShapes_instance       map[*NotePerformanceShape]*NotePerformanceShape
	NotePerformanceShapes_mapString      map[string]*NotePerformanceShape
	NotePerformanceShapeOrder            uint
	NotePerformanceShape_stagedOrder     map[*NotePerformanceShape]uint
	NotePerformanceShape_orderStaged     map[uint]*NotePerformanceShape
	NotePerformanceShapes_reference      map[*NotePerformanceShape]*NotePerformanceShape
	NotePerformanceShapes_referenceOrder map[*NotePerformanceShape]uint

	// insertion point for slice of pointers maps
	OnAfterNotePerformanceShapeCreateCallback OnAfterCreateInterface[NotePerformanceShape]
	OnAfterNotePerformanceShapeUpdateCallback OnAfterUpdateInterface[NotePerformanceShape]
	OnAfterNotePerformanceShapeDeleteCallback OnAfterDeleteInterface[NotePerformanceShape]
	OnAfterNotePerformanceShapeReadCallback   OnAfterReadInterface[NotePerformanceShape]

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

	Performances                map[*Performance]struct{}
	Performances_instance       map[*Performance]*Performance
	Performances_mapString      map[string]*Performance
	PerformanceOrder            uint
	Performance_stagedOrder     map[*Performance]uint
	Performance_orderStaged     map[uint]*Performance
	Performances_reference      map[*Performance]*Performance
	Performances_referenceOrder map[*Performance]uint

	// insertion point for slice of pointers maps
	OnAfterPerformanceCreateCallback OnAfterCreateInterface[Performance]
	OnAfterPerformanceUpdateCallback OnAfterUpdateInterface[Performance]
	OnAfterPerformanceDeleteCallback OnAfterDeleteInterface[Performance]
	OnAfterPerformanceReadCallback   OnAfterReadInterface[Performance]

	Systems                map[*System]struct{}
	Systems_instance       map[*System]*System
	Systems_mapString      map[string]*System
	SystemOrder            uint
	System_stagedOrder     map[*System]uint
	System_orderStaged     map[uint]*System
	Systems_reference      map[*System]*System
	Systems_referenceOrder map[*System]uint

	// insertion point for slice of pointers maps
	System_Complexities_reverseMap map[*Complexity]*System

	System_Performances_reverseMap map[*Performance]*System

	System_Efforts_reverseMap map[*Effort]*System

	System_SubSystems_reverseMap map[*System]*System

	System_DiagramFlossEquations_reverseMap map[*DiagramFlossEquation]*System

	System_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap map[*DiagramFlossEquation]*System

	System_ComplexitysWhoseNodeIsExpanded_reverseMap map[*Complexity]*System

	System_PerformancesWhoseNodeIsExpanded_reverseMap map[*Performance]*System

	System_EffortsWhoseNodeIsExpanded_reverseMap map[*Effort]*System

	OnAfterSystemCreateCallback OnAfterCreateInterface[System]
	OnAfterSystemUpdateCallback OnAfterUpdateInterface[System]
	OnAfterSystemDeleteCallback OnAfterDeleteInterface[System]
	OnAfterSystemReadCallback   OnAfterReadInterface[System]

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
	stage.CompareAnalysiss_reference = make(map[*CompareAnalysis]*CompareAnalysis)
	stage.CompareAnalysiss_instance = make(map[*CompareAnalysis]*CompareAnalysis)
	stage.CompareAnalysiss_referenceOrder = make(map[*CompareAnalysis]uint)

	stage.Complexitys_reference = make(map[*Complexity]*Complexity)
	stage.Complexitys_instance = make(map[*Complexity]*Complexity)
	stage.Complexitys_referenceOrder = make(map[*Complexity]uint)

	stage.DiagramFlossEquations_reference = make(map[*DiagramFlossEquation]*DiagramFlossEquation)
	stage.DiagramFlossEquations_instance = make(map[*DiagramFlossEquation]*DiagramFlossEquation)
	stage.DiagramFlossEquations_referenceOrder = make(map[*DiagramFlossEquation]uint)

	stage.Efforts_reference = make(map[*Effort]*Effort)
	stage.Efforts_instance = make(map[*Effort]*Effort)
	stage.Efforts_referenceOrder = make(map[*Effort]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_instance = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint)

	stage.NoteComplexityShapes_reference = make(map[*NoteComplexityShape]*NoteComplexityShape)
	stage.NoteComplexityShapes_instance = make(map[*NoteComplexityShape]*NoteComplexityShape)
	stage.NoteComplexityShapes_referenceOrder = make(map[*NoteComplexityShape]uint)

	stage.NoteEffortShapes_reference = make(map[*NoteEffortShape]*NoteEffortShape)
	stage.NoteEffortShapes_instance = make(map[*NoteEffortShape]*NoteEffortShape)
	stage.NoteEffortShapes_referenceOrder = make(map[*NoteEffortShape]uint)

	stage.NotePerformanceShapes_reference = make(map[*NotePerformanceShape]*NotePerformanceShape)
	stage.NotePerformanceShapes_instance = make(map[*NotePerformanceShape]*NotePerformanceShape)
	stage.NotePerformanceShapes_referenceOrder = make(map[*NotePerformanceShape]uint)

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)

	stage.Performances_reference = make(map[*Performance]*Performance)
	stage.Performances_instance = make(map[*Performance]*Performance)
	stage.Performances_referenceOrder = make(map[*Performance]uint)

	stage.Systems_reference = make(map[*System]*System)
	stage.Systems_instance = make(map[*System]*System)
	stage.Systems_referenceOrder = make(map[*System]uint)

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
	var maxCompareAnalysisOrder uint
	var foundCompareAnalysis bool
	for _, order := range stage.CompareAnalysis_stagedOrder {
		if !foundCompareAnalysis || order > maxCompareAnalysisOrder {
			maxCompareAnalysisOrder = order
			foundCompareAnalysis = true
		}
	}
	if foundCompareAnalysis {
		stage.CompareAnalysisOrder = maxCompareAnalysisOrder + 1
	} else {
		stage.CompareAnalysisOrder = 0
	}

	var maxComplexityOrder uint
	var foundComplexity bool
	for _, order := range stage.Complexity_stagedOrder {
		if !foundComplexity || order > maxComplexityOrder {
			maxComplexityOrder = order
			foundComplexity = true
		}
	}
	if foundComplexity {
		stage.ComplexityOrder = maxComplexityOrder + 1
	} else {
		stage.ComplexityOrder = 0
	}

	var maxDiagramFlossEquationOrder uint
	var foundDiagramFlossEquation bool
	for _, order := range stage.DiagramFlossEquation_stagedOrder {
		if !foundDiagramFlossEquation || order > maxDiagramFlossEquationOrder {
			maxDiagramFlossEquationOrder = order
			foundDiagramFlossEquation = true
		}
	}
	if foundDiagramFlossEquation {
		stage.DiagramFlossEquationOrder = maxDiagramFlossEquationOrder + 1
	} else {
		stage.DiagramFlossEquationOrder = 0
	}

	var maxEffortOrder uint
	var foundEffort bool
	for _, order := range stage.Effort_stagedOrder {
		if !foundEffort || order > maxEffortOrder {
			maxEffortOrder = order
			foundEffort = true
		}
	}
	if foundEffort {
		stage.EffortOrder = maxEffortOrder + 1
	} else {
		stage.EffortOrder = 0
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

	var maxNoteComplexityShapeOrder uint
	var foundNoteComplexityShape bool
	for _, order := range stage.NoteComplexityShape_stagedOrder {
		if !foundNoteComplexityShape || order > maxNoteComplexityShapeOrder {
			maxNoteComplexityShapeOrder = order
			foundNoteComplexityShape = true
		}
	}
	if foundNoteComplexityShape {
		stage.NoteComplexityShapeOrder = maxNoteComplexityShapeOrder + 1
	} else {
		stage.NoteComplexityShapeOrder = 0
	}

	var maxNoteEffortShapeOrder uint
	var foundNoteEffortShape bool
	for _, order := range stage.NoteEffortShape_stagedOrder {
		if !foundNoteEffortShape || order > maxNoteEffortShapeOrder {
			maxNoteEffortShapeOrder = order
			foundNoteEffortShape = true
		}
	}
	if foundNoteEffortShape {
		stage.NoteEffortShapeOrder = maxNoteEffortShapeOrder + 1
	} else {
		stage.NoteEffortShapeOrder = 0
	}

	var maxNotePerformanceShapeOrder uint
	var foundNotePerformanceShape bool
	for _, order := range stage.NotePerformanceShape_stagedOrder {
		if !foundNotePerformanceShape || order > maxNotePerformanceShapeOrder {
			maxNotePerformanceShapeOrder = order
			foundNotePerformanceShape = true
		}
	}
	if foundNotePerformanceShape {
		stage.NotePerformanceShapeOrder = maxNotePerformanceShapeOrder + 1
	} else {
		stage.NotePerformanceShapeOrder = 0
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

	var maxPerformanceOrder uint
	var foundPerformance bool
	for _, order := range stage.Performance_stagedOrder {
		if !foundPerformance || order > maxPerformanceOrder {
			maxPerformanceOrder = order
			foundPerformance = true
		}
	}
	if foundPerformance {
		stage.PerformanceOrder = maxPerformanceOrder + 1
	} else {
		stage.PerformanceOrder = 0
	}

	var maxSystemOrder uint
	var foundSystem bool
	for _, order := range stage.System_stagedOrder {
		if !foundSystem || order > maxSystemOrder {
			maxSystemOrder = order
			foundSystem = true
		}
	}
	if foundSystem {
		stage.SystemOrder = maxSystemOrder + 1
	} else {
		stage.SystemOrder = 0
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
	case *CompareAnalysis:
		tmp := GetStructInstancesByOrder(stage.CompareAnalysiss, stage.CompareAnalysis_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CompareAnalysis implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Complexity:
		tmp := GetStructInstancesByOrder(stage.Complexitys, stage.Complexity_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Complexity implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DiagramFlossEquation:
		tmp := GetStructInstancesByOrder(stage.DiagramFlossEquations, stage.DiagramFlossEquation_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DiagramFlossEquation implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Effort:
		tmp := GetStructInstancesByOrder(stage.Efforts, stage.Effort_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Effort implements.
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
	case *NoteComplexityShape:
		tmp := GetStructInstancesByOrder(stage.NoteComplexityShapes, stage.NoteComplexityShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteComplexityShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *NoteEffortShape:
		tmp := GetStructInstancesByOrder(stage.NoteEffortShapes, stage.NoteEffortShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteEffortShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *NotePerformanceShape:
		tmp := GetStructInstancesByOrder(stage.NotePerformanceShapes, stage.NotePerformanceShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NotePerformanceShape implements.
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
	case *Performance:
		tmp := GetStructInstancesByOrder(stage.Performances, stage.Performance_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Performance implements.
			res = append(res, any(v).(T))
		}
		return res
	case *System:
		tmp := GetStructInstancesByOrder(stage.Systems, stage.System_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *System implements.
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
	case "CompareAnalysis":
		res = GetNamedStructInstances(stage.CompareAnalysiss, stage.CompareAnalysis_stagedOrder)
	case "Complexity":
		res = GetNamedStructInstances(stage.Complexitys, stage.Complexity_stagedOrder)
	case "DiagramFlossEquation":
		res = GetNamedStructInstances(stage.DiagramFlossEquations, stage.DiagramFlossEquation_stagedOrder)
	case "Effort":
		res = GetNamedStructInstances(stage.Efforts, stage.Effort_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Note":
		res = GetNamedStructInstances(stage.Notes, stage.Note_stagedOrder)
	case "NoteComplexityShape":
		res = GetNamedStructInstances(stage.NoteComplexityShapes, stage.NoteComplexityShape_stagedOrder)
	case "NoteEffortShape":
		res = GetNamedStructInstances(stage.NoteEffortShapes, stage.NoteEffortShape_stagedOrder)
	case "NotePerformanceShape":
		res = GetNamedStructInstances(stage.NotePerformanceShapes, stage.NotePerformanceShape_stagedOrder)
	case "NoteShape":
		res = GetNamedStructInstances(stage.NoteShapes, stage.NoteShape_stagedOrder)
	case "Performance":
		res = GetNamedStructInstances(stage.Performances, stage.Performance_stagedOrder)
	case "System":
		res = GetNamedStructInstances(stage.Systems, stage.System_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/floss/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return floss_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return floss_go.GoDiagramsDir
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
	CommitCompareAnalysis(compareanalysis *CompareAnalysis)
	CheckoutCompareAnalysis(compareanalysis *CompareAnalysis)
	CommitComplexity(complexity *Complexity)
	CheckoutComplexity(complexity *Complexity)
	CommitDiagramFlossEquation(diagramflossequation *DiagramFlossEquation)
	CheckoutDiagramFlossEquation(diagramflossequation *DiagramFlossEquation)
	CommitEffort(effort *Effort)
	CheckoutEffort(effort *Effort)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitNote(note *Note)
	CheckoutNote(note *Note)
	CommitNoteComplexityShape(notecomplexityshape *NoteComplexityShape)
	CheckoutNoteComplexityShape(notecomplexityshape *NoteComplexityShape)
	CommitNoteEffortShape(noteeffortshape *NoteEffortShape)
	CheckoutNoteEffortShape(noteeffortshape *NoteEffortShape)
	CommitNotePerformanceShape(noteperformanceshape *NotePerformanceShape)
	CheckoutNotePerformanceShape(noteperformanceshape *NotePerformanceShape)
	CommitNoteShape(noteshape *NoteShape)
	CheckoutNoteShape(noteshape *NoteShape)
	CommitPerformance(performance *Performance)
	CheckoutPerformance(performance *Performance)
	CommitSystem(system *System)
	CheckoutSystem(system *System)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		CompareAnalysiss:           make(map[*CompareAnalysis]struct{}),
		CompareAnalysiss_mapString: make(map[string]*CompareAnalysis),

		Complexitys:           make(map[*Complexity]struct{}),
		Complexitys_mapString: make(map[string]*Complexity),

		DiagramFlossEquations:           make(map[*DiagramFlossEquation]struct{}),
		DiagramFlossEquations_mapString: make(map[string]*DiagramFlossEquation),

		Efforts:           make(map[*Effort]struct{}),
		Efforts_mapString: make(map[string]*Effort),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		NoteComplexityShapes:           make(map[*NoteComplexityShape]struct{}),
		NoteComplexityShapes_mapString: make(map[string]*NoteComplexityShape),

		NoteEffortShapes:           make(map[*NoteEffortShape]struct{}),
		NoteEffortShapes_mapString: make(map[string]*NoteEffortShape),

		NotePerformanceShapes:           make(map[*NotePerformanceShape]struct{}),
		NotePerformanceShapes_mapString: make(map[string]*NotePerformanceShape),

		NoteShapes:           make(map[*NoteShape]struct{}),
		NoteShapes_mapString: make(map[string]*NoteShape),

		Performances:           make(map[*Performance]struct{}),
		Performances_mapString: make(map[string]*Performance),

		Systems:           make(map[*System]struct{}),
		Systems_mapString: make(map[string]*System),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CompareAnalysis_stagedOrder: make(map[*CompareAnalysis]uint),
		CompareAnalysis_orderStaged: make(map[uint]*CompareAnalysis),
		CompareAnalysiss_reference:  make(map[*CompareAnalysis]*CompareAnalysis),

		Complexity_stagedOrder: make(map[*Complexity]uint),
		Complexity_orderStaged: make(map[uint]*Complexity),
		Complexitys_reference:  make(map[*Complexity]*Complexity),

		DiagramFlossEquation_stagedOrder: make(map[*DiagramFlossEquation]uint),
		DiagramFlossEquation_orderStaged: make(map[uint]*DiagramFlossEquation),
		DiagramFlossEquations_reference:  make(map[*DiagramFlossEquation]*DiagramFlossEquation),

		Effort_stagedOrder: make(map[*Effort]uint),
		Effort_orderStaged: make(map[uint]*Effort),
		Efforts_reference:  make(map[*Effort]*Effort),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		Note_stagedOrder: make(map[*Note]uint),
		Note_orderStaged: make(map[uint]*Note),
		Notes_reference:  make(map[*Note]*Note),

		NoteComplexityShape_stagedOrder: make(map[*NoteComplexityShape]uint),
		NoteComplexityShape_orderStaged: make(map[uint]*NoteComplexityShape),
		NoteComplexityShapes_reference:  make(map[*NoteComplexityShape]*NoteComplexityShape),

		NoteEffortShape_stagedOrder: make(map[*NoteEffortShape]uint),
		NoteEffortShape_orderStaged: make(map[uint]*NoteEffortShape),
		NoteEffortShapes_reference:  make(map[*NoteEffortShape]*NoteEffortShape),

		NotePerformanceShape_stagedOrder: make(map[*NotePerformanceShape]uint),
		NotePerformanceShape_orderStaged: make(map[uint]*NotePerformanceShape),
		NotePerformanceShapes_reference:  make(map[*NotePerformanceShape]*NotePerformanceShape),

		NoteShape_stagedOrder: make(map[*NoteShape]uint),
		NoteShape_orderStaged: make(map[uint]*NoteShape),
		NoteShapes_reference:  make(map[*NoteShape]*NoteShape),

		Performance_stagedOrder: make(map[*Performance]uint),
		Performance_orderStaged: make(map[uint]*Performance),
		Performances_reference:  make(map[*Performance]*Performance),

		System_stagedOrder: make(map[*System]uint),
		System_orderStaged: make(map[uint]*System),
		Systems_reference:  make(map[*System]*System),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"CompareAnalysis": &CompareAnalysisUnmarshaller{},

			"Complexity": &ComplexityUnmarshaller{},

			"DiagramFlossEquation": &DiagramFlossEquationUnmarshaller{},

			"Effort": &EffortUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteComplexityShape": &NoteComplexityShapeUnmarshaller{},

			"NoteEffortShape": &NoteEffortShapeUnmarshaller{},

			"NotePerformanceShape": &NotePerformanceShapeUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"Performance": &PerformanceUnmarshaller{},

			"System": &SystemUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "CompareAnalysis"},
			{name: "Complexity"},
			{name: "DiagramFlossEquation"},
			{name: "Effort"},
			{name: "Library"},
			{name: "Note"},
			{name: "NoteComplexityShape"},
			{name: "NoteEffortShape"},
			{name: "NotePerformanceShape"},
			{name: "NoteShape"},
			{name: "Performance"},
			{name: "System"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *CompareAnalysis:
		return stage.CompareAnalysis_stagedOrder[instance]
	case *Complexity:
		return stage.Complexity_stagedOrder[instance]
	case *DiagramFlossEquation:
		return stage.DiagramFlossEquation_stagedOrder[instance]
	case *Effort:
		return stage.Effort_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteComplexityShape:
		return stage.NoteComplexityShape_stagedOrder[instance]
	case *NoteEffortShape:
		return stage.NoteEffortShape_stagedOrder[instance]
	case *NotePerformanceShape:
		return stage.NotePerformanceShape_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *Performance:
		return stage.Performance_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *CompareAnalysis:
		return any(stage.CompareAnalysis_orderStaged[order]).(Type)
	case *Complexity:
		return any(stage.Complexity_orderStaged[order]).(Type)
	case *DiagramFlossEquation:
		return any(stage.DiagramFlossEquation_orderStaged[order]).(Type)
	case *Effort:
		return any(stage.Effort_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Note:
		return any(stage.Note_orderStaged[order]).(Type)
	case *NoteComplexityShape:
		return any(stage.NoteComplexityShape_orderStaged[order]).(Type)
	case *NoteEffortShape:
		return any(stage.NoteEffortShape_orderStaged[order]).(Type)
	case *NotePerformanceShape:
		return any(stage.NotePerformanceShape_orderStaged[order]).(Type)
	case *NoteShape:
		return any(stage.NoteShape_orderStaged[order]).(Type)
	case *Performance:
		return any(stage.Performance_orderStaged[order]).(Type)
	case *System:
		return any(stage.System_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *CompareAnalysis:
		return stage.CompareAnalysis_stagedOrder[instance]
	case *Complexity:
		return stage.Complexity_stagedOrder[instance]
	case *DiagramFlossEquation:
		return stage.DiagramFlossEquation_stagedOrder[instance]
	case *Effort:
		return stage.Effort_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteComplexityShape:
		return stage.NoteComplexityShape_stagedOrder[instance]
	case *NoteEffortShape:
		return stage.NoteEffortShape_stagedOrder[instance]
	case *NotePerformanceShape:
		return stage.NotePerformanceShape_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *Performance:
		return stage.Performance_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["CompareAnalysis"] = len(stage.CompareAnalysiss)
	stage.Map_GongStructName_InstancesNb["Complexity"] = len(stage.Complexitys)
	stage.Map_GongStructName_InstancesNb["DiagramFlossEquation"] = len(stage.DiagramFlossEquations)
	stage.Map_GongStructName_InstancesNb["Effort"] = len(stage.Efforts)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteComplexityShape"] = len(stage.NoteComplexityShapes)
	stage.Map_GongStructName_InstancesNb["NoteEffortShape"] = len(stage.NoteEffortShapes)
	stage.Map_GongStructName_InstancesNb["NotePerformanceShape"] = len(stage.NotePerformanceShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["Performance"] = len(stage.Performances)
	stage.Map_GongStructName_InstancesNb["System"] = len(stage.Systems)
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
// Stage puts compareanalysis to the model stage
func (compareanalysis *CompareAnalysis) Stage(stage *Stage) *CompareAnalysis {
	if _, ok := stage.CompareAnalysiss[compareanalysis]; !ok {
		stage.CompareAnalysiss[compareanalysis] = struct{}{}
		stage.CompareAnalysis_stagedOrder[compareanalysis] = stage.CompareAnalysisOrder
		stage.CompareAnalysis_orderStaged[stage.CompareAnalysisOrder] = compareanalysis
		stage.CompareAnalysisOrder++
	}
	stage.CompareAnalysiss_mapString[compareanalysis.Name] = compareanalysis

	return compareanalysis
}

// StagePreserveOrder puts compareanalysis to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CompareAnalysisOrder
// - update stage.CompareAnalysisOrder accordingly
func (compareanalysis *CompareAnalysis) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CompareAnalysiss[compareanalysis]; !ok {
		stage.CompareAnalysiss[compareanalysis] = struct{}{}

		if order > stage.CompareAnalysisOrder {
			stage.CompareAnalysisOrder = order
		}
		stage.CompareAnalysis_stagedOrder[compareanalysis] = order
		stage.CompareAnalysis_orderStaged[order] = compareanalysis
		stage.CompareAnalysisOrder++
	}
	stage.CompareAnalysiss_mapString[compareanalysis.Name] = compareanalysis
}

// Unstage removes compareanalysis off the model stage
func (compareanalysis *CompareAnalysis) Unstage(stage *Stage) *CompareAnalysis {
	delete(stage.CompareAnalysiss, compareanalysis)
	// issue1150
	// delete(stage.CompareAnalysis_stagedOrder, compareanalysis)
	delete(stage.CompareAnalysiss_mapString, compareanalysis.Name)

	return compareanalysis
}

// UnstageVoid removes compareanalysis off the model stage
func (compareanalysis *CompareAnalysis) UnstageVoid(stage *Stage) {
	delete(stage.CompareAnalysiss, compareanalysis)
	// issue1150
	// delete(stage.CompareAnalysis_stagedOrder, compareanalysis)
	delete(stage.CompareAnalysiss_mapString, compareanalysis.Name)
}

// commit compareanalysis to the back repo (if it is already staged)
func (compareanalysis *CompareAnalysis) Commit(stage *Stage) *CompareAnalysis {
	if _, ok := stage.CompareAnalysiss[compareanalysis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCompareAnalysis(compareanalysis)
		}
	}
	return compareanalysis
}

func (compareanalysis *CompareAnalysis) CommitVoid(stage *Stage) {
	compareanalysis.Commit(stage)
}

func (compareanalysis *CompareAnalysis) StageVoid(stage *Stage) {
	compareanalysis.Stage(stage)
}

// Checkout compareanalysis to the back repo (if it is already staged)
func (compareanalysis *CompareAnalysis) Checkout(stage *Stage) *CompareAnalysis {
	if _, ok := stage.CompareAnalysiss[compareanalysis]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCompareAnalysis(compareanalysis)
		}
	}
	return compareanalysis
}

// for satisfaction of GongStruct interface
func (compareanalysis *CompareAnalysis) GetName() (res string) {
	return compareanalysis.Name
}

// for satisfaction of GongStruct interface
func (compareanalysis *CompareAnalysis) SetName(name string) {
	compareanalysis.Name = name
}

// Stage puts complexity to the model stage
func (complexity *Complexity) Stage(stage *Stage) *Complexity {
	if _, ok := stage.Complexitys[complexity]; !ok {
		stage.Complexitys[complexity] = struct{}{}
		stage.Complexity_stagedOrder[complexity] = stage.ComplexityOrder
		stage.Complexity_orderStaged[stage.ComplexityOrder] = complexity
		stage.ComplexityOrder++
	}
	stage.Complexitys_mapString[complexity.Name] = complexity

	return complexity
}

// StagePreserveOrder puts complexity to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ComplexityOrder
// - update stage.ComplexityOrder accordingly
func (complexity *Complexity) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Complexitys[complexity]; !ok {
		stage.Complexitys[complexity] = struct{}{}

		if order > stage.ComplexityOrder {
			stage.ComplexityOrder = order
		}
		stage.Complexity_stagedOrder[complexity] = order
		stage.Complexity_orderStaged[order] = complexity
		stage.ComplexityOrder++
	}
	stage.Complexitys_mapString[complexity.Name] = complexity
}

// Unstage removes complexity off the model stage
func (complexity *Complexity) Unstage(stage *Stage) *Complexity {
	delete(stage.Complexitys, complexity)
	// issue1150
	// delete(stage.Complexity_stagedOrder, complexity)
	delete(stage.Complexitys_mapString, complexity.Name)

	return complexity
}

// UnstageVoid removes complexity off the model stage
func (complexity *Complexity) UnstageVoid(stage *Stage) {
	delete(stage.Complexitys, complexity)
	// issue1150
	// delete(stage.Complexity_stagedOrder, complexity)
	delete(stage.Complexitys_mapString, complexity.Name)
}

// commit complexity to the back repo (if it is already staged)
func (complexity *Complexity) Commit(stage *Stage) *Complexity {
	if _, ok := stage.Complexitys[complexity]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitComplexity(complexity)
		}
	}
	return complexity
}

func (complexity *Complexity) CommitVoid(stage *Stage) {
	complexity.Commit(stage)
}

func (complexity *Complexity) StageVoid(stage *Stage) {
	complexity.Stage(stage)
}

// Checkout complexity to the back repo (if it is already staged)
func (complexity *Complexity) Checkout(stage *Stage) *Complexity {
	if _, ok := stage.Complexitys[complexity]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutComplexity(complexity)
		}
	}
	return complexity
}

// for satisfaction of GongStruct interface
func (complexity *Complexity) GetName() (res string) {
	return complexity.Name
}

// for satisfaction of GongStruct interface
func (complexity *Complexity) SetName(name string) {
	complexity.Name = name
}

// Stage puts diagramflossequation to the model stage
func (diagramflossequation *DiagramFlossEquation) Stage(stage *Stage) *DiagramFlossEquation {
	if _, ok := stage.DiagramFlossEquations[diagramflossequation]; !ok {
		stage.DiagramFlossEquations[diagramflossequation] = struct{}{}
		stage.DiagramFlossEquation_stagedOrder[diagramflossequation] = stage.DiagramFlossEquationOrder
		stage.DiagramFlossEquation_orderStaged[stage.DiagramFlossEquationOrder] = diagramflossequation
		stage.DiagramFlossEquationOrder++
	}
	stage.DiagramFlossEquations_mapString[diagramflossequation.Name] = diagramflossequation

	return diagramflossequation
}

// StagePreserveOrder puts diagramflossequation to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramFlossEquationOrder
// - update stage.DiagramFlossEquationOrder accordingly
func (diagramflossequation *DiagramFlossEquation) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DiagramFlossEquations[diagramflossequation]; !ok {
		stage.DiagramFlossEquations[diagramflossequation] = struct{}{}

		if order > stage.DiagramFlossEquationOrder {
			stage.DiagramFlossEquationOrder = order
		}
		stage.DiagramFlossEquation_stagedOrder[diagramflossequation] = order
		stage.DiagramFlossEquation_orderStaged[order] = diagramflossequation
		stage.DiagramFlossEquationOrder++
	}
	stage.DiagramFlossEquations_mapString[diagramflossequation.Name] = diagramflossequation
}

// Unstage removes diagramflossequation off the model stage
func (diagramflossequation *DiagramFlossEquation) Unstage(stage *Stage) *DiagramFlossEquation {
	delete(stage.DiagramFlossEquations, diagramflossequation)
	// issue1150
	// delete(stage.DiagramFlossEquation_stagedOrder, diagramflossequation)
	delete(stage.DiagramFlossEquations_mapString, diagramflossequation.Name)

	return diagramflossequation
}

// UnstageVoid removes diagramflossequation off the model stage
func (diagramflossequation *DiagramFlossEquation) UnstageVoid(stage *Stage) {
	delete(stage.DiagramFlossEquations, diagramflossequation)
	// issue1150
	// delete(stage.DiagramFlossEquation_stagedOrder, diagramflossequation)
	delete(stage.DiagramFlossEquations_mapString, diagramflossequation.Name)
}

// commit diagramflossequation to the back repo (if it is already staged)
func (diagramflossequation *DiagramFlossEquation) Commit(stage *Stage) *DiagramFlossEquation {
	if _, ok := stage.DiagramFlossEquations[diagramflossequation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagramFlossEquation(diagramflossequation)
		}
	}
	return diagramflossequation
}

func (diagramflossequation *DiagramFlossEquation) CommitVoid(stage *Stage) {
	diagramflossequation.Commit(stage)
}

func (diagramflossequation *DiagramFlossEquation) StageVoid(stage *Stage) {
	diagramflossequation.Stage(stage)
}

// Checkout diagramflossequation to the back repo (if it is already staged)
func (diagramflossequation *DiagramFlossEquation) Checkout(stage *Stage) *DiagramFlossEquation {
	if _, ok := stage.DiagramFlossEquations[diagramflossequation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagramFlossEquation(diagramflossequation)
		}
	}
	return diagramflossequation
}

// for satisfaction of GongStruct interface
func (diagramflossequation *DiagramFlossEquation) GetName() (res string) {
	return diagramflossequation.Name
}

// for satisfaction of GongStruct interface
func (diagramflossequation *DiagramFlossEquation) SetName(name string) {
	diagramflossequation.Name = name
}

// Stage puts effort to the model stage
func (effort *Effort) Stage(stage *Stage) *Effort {
	if _, ok := stage.Efforts[effort]; !ok {
		stage.Efforts[effort] = struct{}{}
		stage.Effort_stagedOrder[effort] = stage.EffortOrder
		stage.Effort_orderStaged[stage.EffortOrder] = effort
		stage.EffortOrder++
	}
	stage.Efforts_mapString[effort.Name] = effort

	return effort
}

// StagePreserveOrder puts effort to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EffortOrder
// - update stage.EffortOrder accordingly
func (effort *Effort) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Efforts[effort]; !ok {
		stage.Efforts[effort] = struct{}{}

		if order > stage.EffortOrder {
			stage.EffortOrder = order
		}
		stage.Effort_stagedOrder[effort] = order
		stage.Effort_orderStaged[order] = effort
		stage.EffortOrder++
	}
	stage.Efforts_mapString[effort.Name] = effort
}

// Unstage removes effort off the model stage
func (effort *Effort) Unstage(stage *Stage) *Effort {
	delete(stage.Efforts, effort)
	// issue1150
	// delete(stage.Effort_stagedOrder, effort)
	delete(stage.Efforts_mapString, effort.Name)

	return effort
}

// UnstageVoid removes effort off the model stage
func (effort *Effort) UnstageVoid(stage *Stage) {
	delete(stage.Efforts, effort)
	// issue1150
	// delete(stage.Effort_stagedOrder, effort)
	delete(stage.Efforts_mapString, effort.Name)
}

// commit effort to the back repo (if it is already staged)
func (effort *Effort) Commit(stage *Stage) *Effort {
	if _, ok := stage.Efforts[effort]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEffort(effort)
		}
	}
	return effort
}

func (effort *Effort) CommitVoid(stage *Stage) {
	effort.Commit(stage)
}

func (effort *Effort) StageVoid(stage *Stage) {
	effort.Stage(stage)
}

// Checkout effort to the back repo (if it is already staged)
func (effort *Effort) Checkout(stage *Stage) *Effort {
	if _, ok := stage.Efforts[effort]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEffort(effort)
		}
	}
	return effort
}

// for satisfaction of GongStruct interface
func (effort *Effort) GetName() (res string) {
	return effort.Name
}

// for satisfaction of GongStruct interface
func (effort *Effort) SetName(name string) {
	effort.Name = name
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

// Stage puts notecomplexityshape to the model stage
func (notecomplexityshape *NoteComplexityShape) Stage(stage *Stage) *NoteComplexityShape {
	if _, ok := stage.NoteComplexityShapes[notecomplexityshape]; !ok {
		stage.NoteComplexityShapes[notecomplexityshape] = struct{}{}
		stage.NoteComplexityShape_stagedOrder[notecomplexityshape] = stage.NoteComplexityShapeOrder
		stage.NoteComplexityShape_orderStaged[stage.NoteComplexityShapeOrder] = notecomplexityshape
		stage.NoteComplexityShapeOrder++
	}
	stage.NoteComplexityShapes_mapString[notecomplexityshape.Name] = notecomplexityshape

	return notecomplexityshape
}

// StagePreserveOrder puts notecomplexityshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteComplexityShapeOrder
// - update stage.NoteComplexityShapeOrder accordingly
func (notecomplexityshape *NoteComplexityShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteComplexityShapes[notecomplexityshape]; !ok {
		stage.NoteComplexityShapes[notecomplexityshape] = struct{}{}

		if order > stage.NoteComplexityShapeOrder {
			stage.NoteComplexityShapeOrder = order
		}
		stage.NoteComplexityShape_stagedOrder[notecomplexityshape] = order
		stage.NoteComplexityShape_orderStaged[order] = notecomplexityshape
		stage.NoteComplexityShapeOrder++
	}
	stage.NoteComplexityShapes_mapString[notecomplexityshape.Name] = notecomplexityshape
}

// Unstage removes notecomplexityshape off the model stage
func (notecomplexityshape *NoteComplexityShape) Unstage(stage *Stage) *NoteComplexityShape {
	delete(stage.NoteComplexityShapes, notecomplexityshape)
	// issue1150
	// delete(stage.NoteComplexityShape_stagedOrder, notecomplexityshape)
	delete(stage.NoteComplexityShapes_mapString, notecomplexityshape.Name)

	return notecomplexityshape
}

// UnstageVoid removes notecomplexityshape off the model stage
func (notecomplexityshape *NoteComplexityShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteComplexityShapes, notecomplexityshape)
	// issue1150
	// delete(stage.NoteComplexityShape_stagedOrder, notecomplexityshape)
	delete(stage.NoteComplexityShapes_mapString, notecomplexityshape.Name)
}

// commit notecomplexityshape to the back repo (if it is already staged)
func (notecomplexityshape *NoteComplexityShape) Commit(stage *Stage) *NoteComplexityShape {
	if _, ok := stage.NoteComplexityShapes[notecomplexityshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteComplexityShape(notecomplexityshape)
		}
	}
	return notecomplexityshape
}

func (notecomplexityshape *NoteComplexityShape) CommitVoid(stage *Stage) {
	notecomplexityshape.Commit(stage)
}

func (notecomplexityshape *NoteComplexityShape) StageVoid(stage *Stage) {
	notecomplexityshape.Stage(stage)
}

// Checkout notecomplexityshape to the back repo (if it is already staged)
func (notecomplexityshape *NoteComplexityShape) Checkout(stage *Stage) *NoteComplexityShape {
	if _, ok := stage.NoteComplexityShapes[notecomplexityshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteComplexityShape(notecomplexityshape)
		}
	}
	return notecomplexityshape
}

// for satisfaction of GongStruct interface
func (notecomplexityshape *NoteComplexityShape) GetName() (res string) {
	return notecomplexityshape.Name
}

// for satisfaction of GongStruct interface
func (notecomplexityshape *NoteComplexityShape) SetName(name string) {
	notecomplexityshape.Name = name
}

// Stage puts noteeffortshape to the model stage
func (noteeffortshape *NoteEffortShape) Stage(stage *Stage) *NoteEffortShape {
	if _, ok := stage.NoteEffortShapes[noteeffortshape]; !ok {
		stage.NoteEffortShapes[noteeffortshape] = struct{}{}
		stage.NoteEffortShape_stagedOrder[noteeffortshape] = stage.NoteEffortShapeOrder
		stage.NoteEffortShape_orderStaged[stage.NoteEffortShapeOrder] = noteeffortshape
		stage.NoteEffortShapeOrder++
	}
	stage.NoteEffortShapes_mapString[noteeffortshape.Name] = noteeffortshape

	return noteeffortshape
}

// StagePreserveOrder puts noteeffortshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteEffortShapeOrder
// - update stage.NoteEffortShapeOrder accordingly
func (noteeffortshape *NoteEffortShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteEffortShapes[noteeffortshape]; !ok {
		stage.NoteEffortShapes[noteeffortshape] = struct{}{}

		if order > stage.NoteEffortShapeOrder {
			stage.NoteEffortShapeOrder = order
		}
		stage.NoteEffortShape_stagedOrder[noteeffortshape] = order
		stage.NoteEffortShape_orderStaged[order] = noteeffortshape
		stage.NoteEffortShapeOrder++
	}
	stage.NoteEffortShapes_mapString[noteeffortshape.Name] = noteeffortshape
}

// Unstage removes noteeffortshape off the model stage
func (noteeffortshape *NoteEffortShape) Unstage(stage *Stage) *NoteEffortShape {
	delete(stage.NoteEffortShapes, noteeffortshape)
	// issue1150
	// delete(stage.NoteEffortShape_stagedOrder, noteeffortshape)
	delete(stage.NoteEffortShapes_mapString, noteeffortshape.Name)

	return noteeffortshape
}

// UnstageVoid removes noteeffortshape off the model stage
func (noteeffortshape *NoteEffortShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteEffortShapes, noteeffortshape)
	// issue1150
	// delete(stage.NoteEffortShape_stagedOrder, noteeffortshape)
	delete(stage.NoteEffortShapes_mapString, noteeffortshape.Name)
}

// commit noteeffortshape to the back repo (if it is already staged)
func (noteeffortshape *NoteEffortShape) Commit(stage *Stage) *NoteEffortShape {
	if _, ok := stage.NoteEffortShapes[noteeffortshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteEffortShape(noteeffortshape)
		}
	}
	return noteeffortshape
}

func (noteeffortshape *NoteEffortShape) CommitVoid(stage *Stage) {
	noteeffortshape.Commit(stage)
}

func (noteeffortshape *NoteEffortShape) StageVoid(stage *Stage) {
	noteeffortshape.Stage(stage)
}

// Checkout noteeffortshape to the back repo (if it is already staged)
func (noteeffortshape *NoteEffortShape) Checkout(stage *Stage) *NoteEffortShape {
	if _, ok := stage.NoteEffortShapes[noteeffortshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteEffortShape(noteeffortshape)
		}
	}
	return noteeffortshape
}

// for satisfaction of GongStruct interface
func (noteeffortshape *NoteEffortShape) GetName() (res string) {
	return noteeffortshape.Name
}

// for satisfaction of GongStruct interface
func (noteeffortshape *NoteEffortShape) SetName(name string) {
	noteeffortshape.Name = name
}

// Stage puts noteperformanceshape to the model stage
func (noteperformanceshape *NotePerformanceShape) Stage(stage *Stage) *NotePerformanceShape {
	if _, ok := stage.NotePerformanceShapes[noteperformanceshape]; !ok {
		stage.NotePerformanceShapes[noteperformanceshape] = struct{}{}
		stage.NotePerformanceShape_stagedOrder[noteperformanceshape] = stage.NotePerformanceShapeOrder
		stage.NotePerformanceShape_orderStaged[stage.NotePerformanceShapeOrder] = noteperformanceshape
		stage.NotePerformanceShapeOrder++
	}
	stage.NotePerformanceShapes_mapString[noteperformanceshape.Name] = noteperformanceshape

	return noteperformanceshape
}

// StagePreserveOrder puts noteperformanceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NotePerformanceShapeOrder
// - update stage.NotePerformanceShapeOrder accordingly
func (noteperformanceshape *NotePerformanceShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NotePerformanceShapes[noteperformanceshape]; !ok {
		stage.NotePerformanceShapes[noteperformanceshape] = struct{}{}

		if order > stage.NotePerformanceShapeOrder {
			stage.NotePerformanceShapeOrder = order
		}
		stage.NotePerformanceShape_stagedOrder[noteperformanceshape] = order
		stage.NotePerformanceShape_orderStaged[order] = noteperformanceshape
		stage.NotePerformanceShapeOrder++
	}
	stage.NotePerformanceShapes_mapString[noteperformanceshape.Name] = noteperformanceshape
}

// Unstage removes noteperformanceshape off the model stage
func (noteperformanceshape *NotePerformanceShape) Unstage(stage *Stage) *NotePerformanceShape {
	delete(stage.NotePerformanceShapes, noteperformanceshape)
	// issue1150
	// delete(stage.NotePerformanceShape_stagedOrder, noteperformanceshape)
	delete(stage.NotePerformanceShapes_mapString, noteperformanceshape.Name)

	return noteperformanceshape
}

// UnstageVoid removes noteperformanceshape off the model stage
func (noteperformanceshape *NotePerformanceShape) UnstageVoid(stage *Stage) {
	delete(stage.NotePerformanceShapes, noteperformanceshape)
	// issue1150
	// delete(stage.NotePerformanceShape_stagedOrder, noteperformanceshape)
	delete(stage.NotePerformanceShapes_mapString, noteperformanceshape.Name)
}

// commit noteperformanceshape to the back repo (if it is already staged)
func (noteperformanceshape *NotePerformanceShape) Commit(stage *Stage) *NotePerformanceShape {
	if _, ok := stage.NotePerformanceShapes[noteperformanceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNotePerformanceShape(noteperformanceshape)
		}
	}
	return noteperformanceshape
}

func (noteperformanceshape *NotePerformanceShape) CommitVoid(stage *Stage) {
	noteperformanceshape.Commit(stage)
}

func (noteperformanceshape *NotePerformanceShape) StageVoid(stage *Stage) {
	noteperformanceshape.Stage(stage)
}

// Checkout noteperformanceshape to the back repo (if it is already staged)
func (noteperformanceshape *NotePerformanceShape) Checkout(stage *Stage) *NotePerformanceShape {
	if _, ok := stage.NotePerformanceShapes[noteperformanceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNotePerformanceShape(noteperformanceshape)
		}
	}
	return noteperformanceshape
}

// for satisfaction of GongStruct interface
func (noteperformanceshape *NotePerformanceShape) GetName() (res string) {
	return noteperformanceshape.Name
}

// for satisfaction of GongStruct interface
func (noteperformanceshape *NotePerformanceShape) SetName(name string) {
	noteperformanceshape.Name = name
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

// Stage puts performance to the model stage
func (performance *Performance) Stage(stage *Stage) *Performance {
	if _, ok := stage.Performances[performance]; !ok {
		stage.Performances[performance] = struct{}{}
		stage.Performance_stagedOrder[performance] = stage.PerformanceOrder
		stage.Performance_orderStaged[stage.PerformanceOrder] = performance
		stage.PerformanceOrder++
	}
	stage.Performances_mapString[performance.Name] = performance

	return performance
}

// StagePreserveOrder puts performance to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PerformanceOrder
// - update stage.PerformanceOrder accordingly
func (performance *Performance) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Performances[performance]; !ok {
		stage.Performances[performance] = struct{}{}

		if order > stage.PerformanceOrder {
			stage.PerformanceOrder = order
		}
		stage.Performance_stagedOrder[performance] = order
		stage.Performance_orderStaged[order] = performance
		stage.PerformanceOrder++
	}
	stage.Performances_mapString[performance.Name] = performance
}

// Unstage removes performance off the model stage
func (performance *Performance) Unstage(stage *Stage) *Performance {
	delete(stage.Performances, performance)
	// issue1150
	// delete(stage.Performance_stagedOrder, performance)
	delete(stage.Performances_mapString, performance.Name)

	return performance
}

// UnstageVoid removes performance off the model stage
func (performance *Performance) UnstageVoid(stage *Stage) {
	delete(stage.Performances, performance)
	// issue1150
	// delete(stage.Performance_stagedOrder, performance)
	delete(stage.Performances_mapString, performance.Name)
}

// commit performance to the back repo (if it is already staged)
func (performance *Performance) Commit(stage *Stage) *Performance {
	if _, ok := stage.Performances[performance]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPerformance(performance)
		}
	}
	return performance
}

func (performance *Performance) CommitVoid(stage *Stage) {
	performance.Commit(stage)
}

func (performance *Performance) StageVoid(stage *Stage) {
	performance.Stage(stage)
}

// Checkout performance to the back repo (if it is already staged)
func (performance *Performance) Checkout(stage *Stage) *Performance {
	if _, ok := stage.Performances[performance]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPerformance(performance)
		}
	}
	return performance
}

// for satisfaction of GongStruct interface
func (performance *Performance) GetName() (res string) {
	return performance.Name
}

// for satisfaction of GongStruct interface
func (performance *Performance) SetName(name string) {
	performance.Name = name
}

// Stage puts system to the model stage
func (system *System) Stage(stage *Stage) *System {
	if _, ok := stage.Systems[system]; !ok {
		stage.Systems[system] = struct{}{}
		stage.System_stagedOrder[system] = stage.SystemOrder
		stage.System_orderStaged[stage.SystemOrder] = system
		stage.SystemOrder++
	}
	stage.Systems_mapString[system.Name] = system

	return system
}

// StagePreserveOrder puts system to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemOrder
// - update stage.SystemOrder accordingly
func (system *System) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Systems[system]; !ok {
		stage.Systems[system] = struct{}{}

		if order > stage.SystemOrder {
			stage.SystemOrder = order
		}
		stage.System_stagedOrder[system] = order
		stage.System_orderStaged[order] = system
		stage.SystemOrder++
	}
	stage.Systems_mapString[system.Name] = system
}

// Unstage removes system off the model stage
func (system *System) Unstage(stage *Stage) *System {
	delete(stage.Systems, system)
	// issue1150
	// delete(stage.System_stagedOrder, system)
	delete(stage.Systems_mapString, system.Name)

	return system
}

// UnstageVoid removes system off the model stage
func (system *System) UnstageVoid(stage *Stage) {
	delete(stage.Systems, system)
	// issue1150
	// delete(stage.System_stagedOrder, system)
	delete(stage.Systems_mapString, system.Name)
}

// commit system to the back repo (if it is already staged)
func (system *System) Commit(stage *Stage) *System {
	if _, ok := stage.Systems[system]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSystem(system)
		}
	}
	return system
}

func (system *System) CommitVoid(stage *Stage) {
	system.Commit(stage)
}

func (system *System) StageVoid(stage *Stage) {
	system.Stage(stage)
}

// Checkout system to the back repo (if it is already staged)
func (system *System) Checkout(stage *Stage) *System {
	if _, ok := stage.Systems[system]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSystem(system)
		}
	}
	return system
}

// for satisfaction of GongStruct interface
func (system *System) GetName() (res string) {
	return system.Name
}

// for satisfaction of GongStruct interface
func (system *System) SetName(name string) {
	system.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCompareAnalysis(CompareAnalysis *CompareAnalysis)
	CreateORMComplexity(Complexity *Complexity)
	CreateORMDiagramFlossEquation(DiagramFlossEquation *DiagramFlossEquation)
	CreateORMEffort(Effort *Effort)
	CreateORMLibrary(Library *Library)
	CreateORMNote(Note *Note)
	CreateORMNoteComplexityShape(NoteComplexityShape *NoteComplexityShape)
	CreateORMNoteEffortShape(NoteEffortShape *NoteEffortShape)
	CreateORMNotePerformanceShape(NotePerformanceShape *NotePerformanceShape)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMPerformance(Performance *Performance)
	CreateORMSystem(System *System)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCompareAnalysis(CompareAnalysis *CompareAnalysis)
	DeleteORMComplexity(Complexity *Complexity)
	DeleteORMDiagramFlossEquation(DiagramFlossEquation *DiagramFlossEquation)
	DeleteORMEffort(Effort *Effort)
	DeleteORMLibrary(Library *Library)
	DeleteORMNote(Note *Note)
	DeleteORMNoteComplexityShape(NoteComplexityShape *NoteComplexityShape)
	DeleteORMNoteEffortShape(NoteEffortShape *NoteEffortShape)
	DeleteORMNotePerformanceShape(NotePerformanceShape *NotePerformanceShape)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMPerformance(Performance *Performance)
	DeleteORMSystem(System *System)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.CompareAnalysiss = make(map[*CompareAnalysis]struct{})
	stage.CompareAnalysiss_mapString = make(map[string]*CompareAnalysis)
	stage.CompareAnalysis_stagedOrder = make(map[*CompareAnalysis]uint)
	stage.CompareAnalysisOrder = 0

	stage.Complexitys = make(map[*Complexity]struct{})
	stage.Complexitys_mapString = make(map[string]*Complexity)
	stage.Complexity_stagedOrder = make(map[*Complexity]uint)
	stage.ComplexityOrder = 0

	stage.DiagramFlossEquations = make(map[*DiagramFlossEquation]struct{})
	stage.DiagramFlossEquations_mapString = make(map[string]*DiagramFlossEquation)
	stage.DiagramFlossEquation_stagedOrder = make(map[*DiagramFlossEquation]uint)
	stage.DiagramFlossEquationOrder = 0

	stage.Efforts = make(map[*Effort]struct{})
	stage.Efforts_mapString = make(map[string]*Effort)
	stage.Effort_stagedOrder = make(map[*Effort]uint)
	stage.EffortOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Notes = make(map[*Note]struct{})
	stage.Notes_mapString = make(map[string]*Note)
	stage.Note_stagedOrder = make(map[*Note]uint)
	stage.NoteOrder = 0

	stage.NoteComplexityShapes = make(map[*NoteComplexityShape]struct{})
	stage.NoteComplexityShapes_mapString = make(map[string]*NoteComplexityShape)
	stage.NoteComplexityShape_stagedOrder = make(map[*NoteComplexityShape]uint)
	stage.NoteComplexityShapeOrder = 0

	stage.NoteEffortShapes = make(map[*NoteEffortShape]struct{})
	stage.NoteEffortShapes_mapString = make(map[string]*NoteEffortShape)
	stage.NoteEffortShape_stagedOrder = make(map[*NoteEffortShape]uint)
	stage.NoteEffortShapeOrder = 0

	stage.NotePerformanceShapes = make(map[*NotePerformanceShape]struct{})
	stage.NotePerformanceShapes_mapString = make(map[string]*NotePerformanceShape)
	stage.NotePerformanceShape_stagedOrder = make(map[*NotePerformanceShape]uint)
	stage.NotePerformanceShapeOrder = 0

	stage.NoteShapes = make(map[*NoteShape]struct{})
	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	stage.NoteShape_stagedOrder = make(map[*NoteShape]uint)
	stage.NoteShapeOrder = 0

	stage.Performances = make(map[*Performance]struct{})
	stage.Performances_mapString = make(map[string]*Performance)
	stage.Performance_stagedOrder = make(map[*Performance]uint)
	stage.PerformanceOrder = 0

	stage.Systems = make(map[*System]struct{})
	stage.Systems_mapString = make(map[string]*System)
	stage.System_stagedOrder = make(map[*System]uint)
	stage.SystemOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.CompareAnalysiss = nil
	stage.CompareAnalysiss_mapString = nil

	stage.Complexitys = nil
	stage.Complexitys_mapString = nil

	stage.DiagramFlossEquations = nil
	stage.DiagramFlossEquations_mapString = nil

	stage.Efforts = nil
	stage.Efforts_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Notes = nil
	stage.Notes_mapString = nil

	stage.NoteComplexityShapes = nil
	stage.NoteComplexityShapes_mapString = nil

	stage.NoteEffortShapes = nil
	stage.NoteEffortShapes_mapString = nil

	stage.NotePerformanceShapes = nil
	stage.NotePerformanceShapes_mapString = nil

	stage.NoteShapes = nil
	stage.NoteShapes_mapString = nil

	stage.Performances = nil
	stage.Performances_mapString = nil

	stage.Systems = nil
	stage.Systems_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for compareanalysis := range stage.CompareAnalysiss {
		compareanalysis.Unstage(stage)
	}

	for complexity := range stage.Complexitys {
		complexity.Unstage(stage)
	}

	for diagramflossequation := range stage.DiagramFlossEquations {
		diagramflossequation.Unstage(stage)
	}

	for effort := range stage.Efforts {
		effort.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for note := range stage.Notes {
		note.Unstage(stage)
	}

	for notecomplexityshape := range stage.NoteComplexityShapes {
		notecomplexityshape.Unstage(stage)
	}

	for noteeffortshape := range stage.NoteEffortShapes {
		noteeffortshape.Unstage(stage)
	}

	for noteperformanceshape := range stage.NotePerformanceShapes {
		noteperformanceshape.Unstage(stage)
	}

	for noteshape := range stage.NoteShapes {
		noteshape.Unstage(stage)
	}

	for performance := range stage.Performances {
		performance.Unstage(stage)
	}

	for system := range stage.Systems {
		system.Unstage(stage)
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
	case map[*CompareAnalysis]any:
		return any(&stage.CompareAnalysiss).(*Type)
	case map[*Complexity]any:
		return any(&stage.Complexitys).(*Type)
	case map[*DiagramFlossEquation]any:
		return any(&stage.DiagramFlossEquations).(*Type)
	case map[*Effort]any:
		return any(&stage.Efforts).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Note]any:
		return any(&stage.Notes).(*Type)
	case map[*NoteComplexityShape]any:
		return any(&stage.NoteComplexityShapes).(*Type)
	case map[*NoteEffortShape]any:
		return any(&stage.NoteEffortShapes).(*Type)
	case map[*NotePerformanceShape]any:
		return any(&stage.NotePerformanceShapes).(*Type)
	case map[*NoteShape]any:
		return any(&stage.NoteShapes).(*Type)
	case map[*Performance]any:
		return any(&stage.Performances).(*Type)
	case map[*System]any:
		return any(&stage.Systems).(*Type)
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
	case *CompareAnalysis:
		return any(stage.CompareAnalysiss_mapString).(map[string]Type)
	case *Complexity:
		return any(stage.Complexitys_mapString).(map[string]Type)
	case *DiagramFlossEquation:
		return any(stage.DiagramFlossEquations_mapString).(map[string]Type)
	case *Effort:
		return any(stage.Efforts_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Note:
		return any(stage.Notes_mapString).(map[string]Type)
	case *NoteComplexityShape:
		return any(stage.NoteComplexityShapes_mapString).(map[string]Type)
	case *NoteEffortShape:
		return any(stage.NoteEffortShapes_mapString).(map[string]Type)
	case *NotePerformanceShape:
		return any(stage.NotePerformanceShapes_mapString).(map[string]Type)
	case *NoteShape:
		return any(stage.NoteShapes_mapString).(map[string]Type)
	case *Performance:
		return any(stage.Performances_mapString).(map[string]Type)
	case *System:
		return any(stage.Systems_mapString).(map[string]Type)
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
	case CompareAnalysis:
		return any(&stage.CompareAnalysiss).(*map[*Type]struct{})
	case Complexity:
		return any(&stage.Complexitys).(*map[*Type]struct{})
	case DiagramFlossEquation:
		return any(&stage.DiagramFlossEquations).(*map[*Type]struct{})
	case Effort:
		return any(&stage.Efforts).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Note:
		return any(&stage.Notes).(*map[*Type]struct{})
	case NoteComplexityShape:
		return any(&stage.NoteComplexityShapes).(*map[*Type]struct{})
	case NoteEffortShape:
		return any(&stage.NoteEffortShapes).(*map[*Type]struct{})
	case NotePerformanceShape:
		return any(&stage.NotePerformanceShapes).(*map[*Type]struct{})
	case NoteShape:
		return any(&stage.NoteShapes).(*map[*Type]struct{})
	case Performance:
		return any(&stage.Performances).(*map[*Type]struct{})
	case System:
		return any(&stage.Systems).(*map[*Type]struct{})
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
	case *CompareAnalysis:
		return any(&stage.CompareAnalysiss).(*map[Type]struct{})
	case *Complexity:
		return any(&stage.Complexitys).(*map[Type]struct{})
	case *DiagramFlossEquation:
		return any(&stage.DiagramFlossEquations).(*map[Type]struct{})
	case *Effort:
		return any(&stage.Efforts).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Note:
		return any(&stage.Notes).(*map[Type]struct{})
	case *NoteComplexityShape:
		return any(&stage.NoteComplexityShapes).(*map[Type]struct{})
	case *NoteEffortShape:
		return any(&stage.NoteEffortShapes).(*map[Type]struct{})
	case *NotePerformanceShape:
		return any(&stage.NotePerformanceShapes).(*map[Type]struct{})
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]struct{})
	case *Performance:
		return any(&stage.Performances).(*map[Type]struct{})
	case *System:
		return any(&stage.Systems).(*map[Type]struct{})
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
	case CompareAnalysis:
		return any(&stage.CompareAnalysiss_mapString).(*map[string]*Type)
	case Complexity:
		return any(&stage.Complexitys_mapString).(*map[string]*Type)
	case DiagramFlossEquation:
		return any(&stage.DiagramFlossEquations_mapString).(*map[string]*Type)
	case Effort:
		return any(&stage.Efforts_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Note:
		return any(&stage.Notes_mapString).(*map[string]*Type)
	case NoteComplexityShape:
		return any(&stage.NoteComplexityShapes_mapString).(*map[string]*Type)
	case NoteEffortShape:
		return any(&stage.NoteEffortShapes_mapString).(*map[string]*Type)
	case NotePerformanceShape:
		return any(&stage.NotePerformanceShapes_mapString).(*map[string]*Type)
	case NoteShape:
		return any(&stage.NoteShapes_mapString).(*map[string]*Type)
	case Performance:
		return any(&stage.Performances_mapString).(*map[string]*Type)
	case System:
		return any(&stage.Systems_mapString).(*map[string]*Type)
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
	case CompareAnalysis:
		return any(&CompareAnalysis{
			// Initialisation of associations
			// field is initialized with an instance of System with the name of the field
			FromSystem: &System{Name: "FromSystem"},
			// field is initialized with an instance of System with the name of the field
			ToSystem: &System{Name: "ToSystem"},
			// field is initialized with an instance of DiagramFlossEquation with the name of the field
			DiagramFlossEquations: []*DiagramFlossEquation{{Name: "DiagramFlossEquations"}},
			// field is initialized with an instance of DiagramFlossEquation with the name of the field
			DiagramFlossEquationsWhoseNodeIsExpanded: []*DiagramFlossEquation{{Name: "DiagramFlossEquationsWhoseNodeIsExpanded"}},
		}).(*Type)
	case Complexity:
		return any(&Complexity{
			// Initialisation of associations
		}).(*Type)
	case DiagramFlossEquation:
		return any(&DiagramFlossEquation{
			// Initialisation of associations
			// field is initialized with an instance of NoteShape with the name of the field
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			// field is initialized with an instance of NoteComplexityShape with the name of the field
			NoteComplexityShapes: []*NoteComplexityShape{{Name: "NoteComplexityShapes"}},
			// field is initialized with an instance of NotePerformanceShape with the name of the field
			NotePerformanceShapes: []*NotePerformanceShape{{Name: "NotePerformanceShapes"}},
			// field is initialized with an instance of NoteEffortShape with the name of the field
			NoteEffortShapes: []*NoteEffortShape{{Name: "NoteEffortShapes"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Complexity with the name of the field
			ComplexitysWhoseNodeIsExpanded: []*Complexity{{Name: "ComplexitysWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Performance with the name of the field
			PerformancesWhoseNodeIsExpanded: []*Performance{{Name: "PerformancesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Effort with the name of the field
			EffortsWhoseNodeIsExpanded: []*Effort{{Name: "EffortsWhoseNodeIsExpanded"}},
		}).(*Type)
	case Effort:
		return any(&Effort{
			// Initialisation of associations
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of System with the name of the field
			RootSystems: []*System{{Name: "RootSystems"}},
			// field is initialized with an instance of Complexity with the name of the field
			RootComplexitys: []*Complexity{{Name: "RootComplexitys"}},
			// field is initialized with an instance of Performance with the name of the field
			RootPerformances: []*Performance{{Name: "RootPerformances"}},
			// field is initialized with an instance of Effort with the name of the field
			RootEfforts: []*Effort{{Name: "RootEfforts"}},
			// field is initialized with an instance of CompareAnalysis with the name of the field
			RootCompareAnalysis: []*CompareAnalysis{{Name: "RootCompareAnalysis"}},
			// field is initialized with an instance of Note with the name of the field
			RootNotes: []*Note{{Name: "RootNotes"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of System with the name of the field
			SystemsWhoseNodeIsExpanded: []*System{{Name: "SystemsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Complexity with the name of the field
			ComplexitysWhoseNodeIsExpanded: []*Complexity{{Name: "ComplexitysWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Performance with the name of the field
			PerformancesWhoseNodeIsExpanded: []*Performance{{Name: "PerformancesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Effort with the name of the field
			EffortsWhoseNodeIsExpanded: []*Effort{{Name: "EffortsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of CompareAnalysis with the name of the field
			CompareAnalysisWhoseNodeIsExpanded: []*CompareAnalysis{{Name: "CompareAnalysisWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
		}).(*Type)
	case Note:
		return any(&Note{
			// Initialisation of associations
			// field is initialized with an instance of Complexity with the name of the field
			Complexities: []*Complexity{{Name: "Complexities"}},
			// field is initialized with an instance of Performance with the name of the field
			Performances: []*Performance{{Name: "Performances"}},
			// field is initialized with an instance of Effort with the name of the field
			Efforts: []*Effort{{Name: "Efforts"}},
		}).(*Type)
	case NoteComplexityShape:
		return any(&NoteComplexityShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Complexity with the name of the field
			Complexity: &Complexity{Name: "Complexity"},
		}).(*Type)
	case NoteEffortShape:
		return any(&NoteEffortShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Effort with the name of the field
			Effort: &Effort{Name: "Effort"},
		}).(*Type)
	case NotePerformanceShape:
		return any(&NotePerformanceShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Performance with the name of the field
			Performance: &Performance{Name: "Performance"},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
		}).(*Type)
	case Performance:
		return any(&Performance{
			// Initialisation of associations
		}).(*Type)
	case System:
		return any(&System{
			// Initialisation of associations
			// field is initialized with an instance of Complexity with the name of the field
			Complexities: []*Complexity{{Name: "Complexities"}},
			// field is initialized with an instance of Performance with the name of the field
			Performances: []*Performance{{Name: "Performances"}},
			// field is initialized with an instance of Effort with the name of the field
			Efforts: []*Effort{{Name: "Efforts"}},
			// field is initialized with an instance of System with the name of the field
			SubSystems: []*System{{Name: "SubSystems"}},
			// field is initialized with an instance of DiagramFlossEquation with the name of the field
			DiagramFlossEquations: []*DiagramFlossEquation{{Name: "DiagramFlossEquations"}},
			// field is initialized with an instance of DiagramFlossEquation with the name of the field
			DiagramFlossEquationsWhoseNodeIsExpanded: []*DiagramFlossEquation{{Name: "DiagramFlossEquationsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Complexity with the name of the field
			ComplexitysWhoseNodeIsExpanded: []*Complexity{{Name: "ComplexitysWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Performance with the name of the field
			PerformancesWhoseNodeIsExpanded: []*Performance{{Name: "PerformancesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Effort with the name of the field
			EffortsWhoseNodeIsExpanded: []*Effort{{Name: "EffortsWhoseNodeIsExpanded"}},
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
	// reverse maps of direct associations of CompareAnalysis
	case CompareAnalysis:
		switch fieldname {
		// insertion point for per direct association field
		case "FromSystem":
			res := make(map[*System][]*CompareAnalysis)
			for compareanalysis := range stage.CompareAnalysiss {
				if compareanalysis.FromSystem != nil {
					system_ := compareanalysis.FromSystem
					var compareanalysiss []*CompareAnalysis
					_, ok := res[system_]
					if ok {
						compareanalysiss = res[system_]
					} else {
						compareanalysiss = make([]*CompareAnalysis, 0)
					}
					compareanalysiss = append(compareanalysiss, compareanalysis)
					res[system_] = compareanalysiss
				}
			}
			return any(res).(map[*End][]*Start)
		case "ToSystem":
			res := make(map[*System][]*CompareAnalysis)
			for compareanalysis := range stage.CompareAnalysiss {
				if compareanalysis.ToSystem != nil {
					system_ := compareanalysis.ToSystem
					var compareanalysiss []*CompareAnalysis
					_, ok := res[system_]
					if ok {
						compareanalysiss = res[system_]
					} else {
						compareanalysiss = make([]*CompareAnalysis, 0)
					}
					compareanalysiss = append(compareanalysiss, compareanalysis)
					res[system_] = compareanalysiss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Complexity
	case Complexity:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DiagramFlossEquation
	case DiagramFlossEquation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Effort
	case Effort:
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
	// reverse maps of direct associations of NoteComplexityShape
	case NoteComplexityShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteComplexityShape)
			for notecomplexityshape := range stage.NoteComplexityShapes {
				if notecomplexityshape.Note != nil {
					note_ := notecomplexityshape.Note
					var notecomplexityshapes []*NoteComplexityShape
					_, ok := res[note_]
					if ok {
						notecomplexityshapes = res[note_]
					} else {
						notecomplexityshapes = make([]*NoteComplexityShape, 0)
					}
					notecomplexityshapes = append(notecomplexityshapes, notecomplexityshape)
					res[note_] = notecomplexityshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Complexity":
			res := make(map[*Complexity][]*NoteComplexityShape)
			for notecomplexityshape := range stage.NoteComplexityShapes {
				if notecomplexityshape.Complexity != nil {
					complexity_ := notecomplexityshape.Complexity
					var notecomplexityshapes []*NoteComplexityShape
					_, ok := res[complexity_]
					if ok {
						notecomplexityshapes = res[complexity_]
					} else {
						notecomplexityshapes = make([]*NoteComplexityShape, 0)
					}
					notecomplexityshapes = append(notecomplexityshapes, notecomplexityshape)
					res[complexity_] = notecomplexityshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteEffortShape
	case NoteEffortShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteEffortShape)
			for noteeffortshape := range stage.NoteEffortShapes {
				if noteeffortshape.Note != nil {
					note_ := noteeffortshape.Note
					var noteeffortshapes []*NoteEffortShape
					_, ok := res[note_]
					if ok {
						noteeffortshapes = res[note_]
					} else {
						noteeffortshapes = make([]*NoteEffortShape, 0)
					}
					noteeffortshapes = append(noteeffortshapes, noteeffortshape)
					res[note_] = noteeffortshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Effort":
			res := make(map[*Effort][]*NoteEffortShape)
			for noteeffortshape := range stage.NoteEffortShapes {
				if noteeffortshape.Effort != nil {
					effort_ := noteeffortshape.Effort
					var noteeffortshapes []*NoteEffortShape
					_, ok := res[effort_]
					if ok {
						noteeffortshapes = res[effort_]
					} else {
						noteeffortshapes = make([]*NoteEffortShape, 0)
					}
					noteeffortshapes = append(noteeffortshapes, noteeffortshape)
					res[effort_] = noteeffortshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NotePerformanceShape
	case NotePerformanceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NotePerformanceShape)
			for noteperformanceshape := range stage.NotePerformanceShapes {
				if noteperformanceshape.Note != nil {
					note_ := noteperformanceshape.Note
					var noteperformanceshapes []*NotePerformanceShape
					_, ok := res[note_]
					if ok {
						noteperformanceshapes = res[note_]
					} else {
						noteperformanceshapes = make([]*NotePerformanceShape, 0)
					}
					noteperformanceshapes = append(noteperformanceshapes, noteperformanceshape)
					res[note_] = noteperformanceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Performance":
			res := make(map[*Performance][]*NotePerformanceShape)
			for noteperformanceshape := range stage.NotePerformanceShapes {
				if noteperformanceshape.Performance != nil {
					performance_ := noteperformanceshape.Performance
					var noteperformanceshapes []*NotePerformanceShape
					_, ok := res[performance_]
					if ok {
						noteperformanceshapes = res[performance_]
					} else {
						noteperformanceshapes = make([]*NotePerformanceShape, 0)
					}
					noteperformanceshapes = append(noteperformanceshapes, noteperformanceshape)
					res[performance_] = noteperformanceshapes
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
	// reverse maps of direct associations of Performance
	case Performance:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of System
	case System:
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
	// reverse maps of direct associations of CompareAnalysis
	case CompareAnalysis:
		switch fieldname {
		// insertion point for per direct association field
		case "DiagramFlossEquations":
			res := make(map[*DiagramFlossEquation][]*CompareAnalysis)
			for compareanalysis := range stage.CompareAnalysiss {
				for _, diagramflossequation_ := range compareanalysis.DiagramFlossEquations {
					res[diagramflossequation_] = append(res[diagramflossequation_], compareanalysis)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			res := make(map[*DiagramFlossEquation][]*CompareAnalysis)
			for compareanalysis := range stage.CompareAnalysiss {
				for _, diagramflossequation_ := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
					res[diagramflossequation_] = append(res[diagramflossequation_], compareanalysis)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Complexity
	case Complexity:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DiagramFlossEquation
	case DiagramFlossEquation:
		switch fieldname {
		// insertion point for per direct association field
		case "Note_Shapes":
			res := make(map[*NoteShape][]*DiagramFlossEquation)
			for diagramflossequation := range stage.DiagramFlossEquations {
				for _, noteshape_ := range diagramflossequation.Note_Shapes {
					res[noteshape_] = append(res[noteshape_], diagramflossequation)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteComplexityShapes":
			res := make(map[*NoteComplexityShape][]*DiagramFlossEquation)
			for diagramflossequation := range stage.DiagramFlossEquations {
				for _, notecomplexityshape_ := range diagramflossequation.NoteComplexityShapes {
					res[notecomplexityshape_] = append(res[notecomplexityshape_], diagramflossequation)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotePerformanceShapes":
			res := make(map[*NotePerformanceShape][]*DiagramFlossEquation)
			for diagramflossequation := range stage.DiagramFlossEquations {
				for _, noteperformanceshape_ := range diagramflossequation.NotePerformanceShapes {
					res[noteperformanceshape_] = append(res[noteperformanceshape_], diagramflossequation)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteEffortShapes":
			res := make(map[*NoteEffortShape][]*DiagramFlossEquation)
			for diagramflossequation := range stage.DiagramFlossEquations {
				for _, noteeffortshape_ := range diagramflossequation.NoteEffortShapes {
					res[noteeffortshape_] = append(res[noteeffortshape_], diagramflossequation)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotesWhoseNodeIsExpanded":
			res := make(map[*Note][]*DiagramFlossEquation)
			for diagramflossequation := range stage.DiagramFlossEquations {
				for _, note_ := range diagramflossequation.NotesWhoseNodeIsExpanded {
					res[note_] = append(res[note_], diagramflossequation)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexitysWhoseNodeIsExpanded":
			res := make(map[*Complexity][]*DiagramFlossEquation)
			for diagramflossequation := range stage.DiagramFlossEquations {
				for _, complexity_ := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
					res[complexity_] = append(res[complexity_], diagramflossequation)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PerformancesWhoseNodeIsExpanded":
			res := make(map[*Performance][]*DiagramFlossEquation)
			for diagramflossequation := range stage.DiagramFlossEquations {
				for _, performance_ := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
					res[performance_] = append(res[performance_], diagramflossequation)
				}
			}
			return any(res).(map[*End][]*Start)
		case "EffortsWhoseNodeIsExpanded":
			res := make(map[*Effort][]*DiagramFlossEquation)
			for diagramflossequation := range stage.DiagramFlossEquations {
				for _, effort_ := range diagramflossequation.EffortsWhoseNodeIsExpanded {
					res[effort_] = append(res[effort_], diagramflossequation)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Effort
	case Effort:
		switch fieldname {
		// insertion point for per direct association field
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
		case "RootSystems":
			res := make(map[*System][]*Library)
			for library := range stage.Librarys {
				for _, system_ := range library.RootSystems {
					res[system_] = append(res[system_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootComplexitys":
			res := make(map[*Complexity][]*Library)
			for library := range stage.Librarys {
				for _, complexity_ := range library.RootComplexitys {
					res[complexity_] = append(res[complexity_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootPerformances":
			res := make(map[*Performance][]*Library)
			for library := range stage.Librarys {
				for _, performance_ := range library.RootPerformances {
					res[performance_] = append(res[performance_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootEfforts":
			res := make(map[*Effort][]*Library)
			for library := range stage.Librarys {
				for _, effort_ := range library.RootEfforts {
					res[effort_] = append(res[effort_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootCompareAnalysis":
			res := make(map[*CompareAnalysis][]*Library)
			for library := range stage.Librarys {
				for _, compareanalysis_ := range library.RootCompareAnalysis {
					res[compareanalysis_] = append(res[compareanalysis_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootNotes":
			res := make(map[*Note][]*Library)
			for library := range stage.Librarys {
				for _, note_ := range library.RootNotes {
					res[note_] = append(res[note_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubLibrariesWhoseNodeIsExpanded":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibrariesWhoseNodeIsExpanded {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SystemsWhoseNodeIsExpanded":
			res := make(map[*System][]*Library)
			for library := range stage.Librarys {
				for _, system_ := range library.SystemsWhoseNodeIsExpanded {
					res[system_] = append(res[system_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexitysWhoseNodeIsExpanded":
			res := make(map[*Complexity][]*Library)
			for library := range stage.Librarys {
				for _, complexity_ := range library.ComplexitysWhoseNodeIsExpanded {
					res[complexity_] = append(res[complexity_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PerformancesWhoseNodeIsExpanded":
			res := make(map[*Performance][]*Library)
			for library := range stage.Librarys {
				for _, performance_ := range library.PerformancesWhoseNodeIsExpanded {
					res[performance_] = append(res[performance_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "EffortsWhoseNodeIsExpanded":
			res := make(map[*Effort][]*Library)
			for library := range stage.Librarys {
				for _, effort_ := range library.EffortsWhoseNodeIsExpanded {
					res[effort_] = append(res[effort_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "CompareAnalysisWhoseNodeIsExpanded":
			res := make(map[*CompareAnalysis][]*Library)
			for library := range stage.Librarys {
				for _, compareanalysis_ := range library.CompareAnalysisWhoseNodeIsExpanded {
					res[compareanalysis_] = append(res[compareanalysis_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotesWhoseNodeIsExpanded":
			res := make(map[*Note][]*Library)
			for library := range stage.Librarys {
				for _, note_ := range library.NotesWhoseNodeIsExpanded {
					res[note_] = append(res[note_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "Complexities":
			res := make(map[*Complexity][]*Note)
			for note := range stage.Notes {
				for _, complexity_ := range note.Complexities {
					res[complexity_] = append(res[complexity_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Performances":
			res := make(map[*Performance][]*Note)
			for note := range stage.Notes {
				for _, performance_ := range note.Performances {
					res[performance_] = append(res[performance_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Efforts":
			res := make(map[*Effort][]*Note)
			for note := range stage.Notes {
				for _, effort_ := range note.Efforts {
					res[effort_] = append(res[effort_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteComplexityShape
	case NoteComplexityShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteEffortShape
	case NoteEffortShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NotePerformanceShape
	case NotePerformanceShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Performance
	case Performance:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of System
	case System:
		switch fieldname {
		// insertion point for per direct association field
		case "Complexities":
			res := make(map[*Complexity][]*System)
			for system := range stage.Systems {
				for _, complexity_ := range system.Complexities {
					res[complexity_] = append(res[complexity_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Performances":
			res := make(map[*Performance][]*System)
			for system := range stage.Systems {
				for _, performance_ := range system.Performances {
					res[performance_] = append(res[performance_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Efforts":
			res := make(map[*Effort][]*System)
			for system := range stage.Systems {
				for _, effort_ := range system.Efforts {
					res[effort_] = append(res[effort_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubSystems":
			res := make(map[*System][]*System)
			for system := range stage.Systems {
				for _, system_ := range system.SubSystems {
					res[system_] = append(res[system_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramFlossEquations":
			res := make(map[*DiagramFlossEquation][]*System)
			for system := range stage.Systems {
				for _, diagramflossequation_ := range system.DiagramFlossEquations {
					res[diagramflossequation_] = append(res[diagramflossequation_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			res := make(map[*DiagramFlossEquation][]*System)
			for system := range stage.Systems {
				for _, diagramflossequation_ := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
					res[diagramflossequation_] = append(res[diagramflossequation_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexitysWhoseNodeIsExpanded":
			res := make(map[*Complexity][]*System)
			for system := range stage.Systems {
				for _, complexity_ := range system.ComplexitysWhoseNodeIsExpanded {
					res[complexity_] = append(res[complexity_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PerformancesWhoseNodeIsExpanded":
			res := make(map[*Performance][]*System)
			for system := range stage.Systems {
				for _, performance_ := range system.PerformancesWhoseNodeIsExpanded {
					res[performance_] = append(res[performance_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "EffortsWhoseNodeIsExpanded":
			res := make(map[*Effort][]*System)
			for system := range stage.Systems {
				for _, effort_ := range system.EffortsWhoseNodeIsExpanded {
					res[effort_] = append(res[effort_], system)
				}
			}
			return any(res).(map[*End][]*Start)
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
	case *CompareAnalysis:
		res = "CompareAnalysis"
	case *Complexity:
		res = "Complexity"
	case *DiagramFlossEquation:
		res = "DiagramFlossEquation"
	case *Effort:
		res = "Effort"
	case *Library:
		res = "Library"
	case *Note:
		res = "Note"
	case *NoteComplexityShape:
		res = "NoteComplexityShape"
	case *NoteEffortShape:
		res = "NoteEffortShape"
	case *NotePerformanceShape:
		res = "NotePerformanceShape"
	case *NoteShape:
		res = "NoteShape"
	case *Performance:
		res = "Performance"
	case *System:
		res = "System"
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
	case *CompareAnalysis:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "RootCompareAnalysis"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "CompareAnalysisWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Complexity:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "ComplexitysWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootComplexitys"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "ComplexitysWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Complexities"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "Complexities"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "ComplexitysWhoseNodeIsExpanded"
		res = append(res, rf)
	case *DiagramFlossEquation:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "CompareAnalysis"
		rf.Fieldname = "DiagramFlossEquations"
		res = append(res, rf)
		rf.GongstructName = "CompareAnalysis"
		rf.Fieldname = "DiagramFlossEquationsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramFlossEquations"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramFlossEquationsWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Effort:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "EffortsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootEfforts"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "EffortsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Efforts"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "Efforts"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "EffortsWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibrariesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Note:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "NotesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootNotes"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "NotesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *NoteComplexityShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "NoteComplexityShapes"
		res = append(res, rf)
	case *NoteEffortShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "NoteEffortShapes"
		res = append(res, rf)
	case *NotePerformanceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "NotePerformanceShapes"
		res = append(res, rf)
	case *NoteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "Note_Shapes"
		res = append(res, rf)
	case *Performance:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "PerformancesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootPerformances"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "PerformancesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Performances"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "Performances"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "PerformancesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *System:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "RootSystems"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SystemsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "SubSystems"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (compareanalysis *CompareAnalysis) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "FromSystem",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "System",
		},
		{
			Name:                 "ToSystem",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "System",
		},
		{
			Name:               "Mu",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Epsilon",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "DiagramFlossEquations",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramFlossEquation",
		},
		{
			Name:                 "DiagramFlossEquationsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramFlossEquation",
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

func (complexity *Complexity) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Strength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Description",
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

func (diagramflossequation *DiagramFlossEquation) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "Scale",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "FontSize",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "FontSize",
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
			Name:               "IsInDelta3ColumnsMode",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AreQuantitativeElementsVisible",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AreSubsystemsVisible",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AreCommonElementsHidden",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AreCPEArrowsVisible",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "AreColumnTitlesVisible",
			GongFieldValueType: GongFieldValueTypeBool,
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
			Name:               "DefaultBoxWidth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "DefaultBoxHeigth",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "Note_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteShape",
		},
		{
			Name:                 "NoteComplexityShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteComplexityShape",
		},
		{
			Name:                 "NotePerformanceShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NotePerformanceShape",
		},
		{
			Name:                 "NoteEffortShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteEffortShape",
		},
		{
			Name:               "IsNotesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "NotesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
		},
		{
			Name:               "IsComplexitysNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ComplexitysWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Complexity",
		},
		{
			Name:               "IsPerformancesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "PerformancesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Performance",
		},
		{
			Name:               "IsEffortsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "EffortsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Effort",
		},
	}
	return
}

func (effort *Effort) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Strength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Description",
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

func (library *Library) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibraries",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:                 "RootSystems",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:                 "RootComplexitys",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Complexity",
		},
		{
			Name:                 "RootPerformances",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Performance",
		},
		{
			Name:                 "RootEfforts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Effort",
		},
		{
			Name:                 "RootCompareAnalysis",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CompareAnalysis",
		},
		{
			Name:                 "RootNotes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
		},
		{
			Name:               "IsRootLibrary",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsSubLibrariesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibrariesWhoseNodeIsExpanded",
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
			Name:               "IsSystemsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SystemsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:               "IsComplexitysNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ComplexitysWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Complexity",
		},
		{
			Name:               "IsPerformancesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "PerformancesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Performance",
		},
		{
			Name:               "IsEffortsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "EffortsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Effort",
		},
		{
			Name:               "IsCompareAnalysisNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "CompareAnalysisWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "CompareAnalysis",
		},
		{
			Name:               "IsNotesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "NotesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
		},
		{
			Name:               "IsExpandedTmp",
			GongFieldValueType: GongFieldValueTypeBool,
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Complexities",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Complexity",
		},
		{
			Name:                 "Performances",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Performance",
		},
		{
			Name:                 "Efforts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Effort",
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
			Name:               "IsComplexitysNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsPerformancesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEffortsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (notecomplexityshape *NoteComplexityShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Complexity",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Complexity",
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

func (noteeffortshape *NoteEffortShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Effort",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Effort",
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

func (noteperformanceshape *NotePerformanceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Performance",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Performance",
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

func (performance *Performance) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Strength",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Description",
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

func (system *System) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Complexities",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Complexity",
		},
		{
			Name:                 "Performances",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Performance",
		},
		{
			Name:                 "Efforts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Effort",
		},
		{
			Name:                 "SubSystems",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "System",
		},
		{
			Name:               "AreCPEsCompoundedFromSubSystems",
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
			Name:               "SVG_Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InverseAppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:                 "DiagramFlossEquations",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramFlossEquation",
		},
		{
			Name:                 "DiagramFlossEquationsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramFlossEquation",
		},
		{
			Name:               "IsSubSystemNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsComplexitysNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ComplexitysWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Complexity",
		},
		{
			Name:               "IsPerformancesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "PerformancesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Performance",
		},
		{
			Name:               "IsEffortsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "EffortsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Effort",
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
func (compareanalysis *CompareAnalysis) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = compareanalysis.Name
	case "FromSystem":
		res.GongFieldValueType = GongFieldValueTypePointer
		if compareanalysis.FromSystem != nil {
			res.valueString = compareanalysis.FromSystem.Name
			res.ids = compareanalysis.FromSystem.GongGetUUID(stage)
		}
	case "ToSystem":
		res.GongFieldValueType = GongFieldValueTypePointer
		if compareanalysis.ToSystem != nil {
			res.valueString = compareanalysis.ToSystem.Name
			res.ids = compareanalysis.ToSystem.GongGetUUID(stage)
		}
	case "Mu":
		res.valueString = fmt.Sprintf("%f", compareanalysis.Mu)
		res.valueFloat = compareanalysis.Mu
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Epsilon":
		res.valueString = fmt.Sprintf("%f", compareanalysis.Epsilon)
		res.valueFloat = compareanalysis.Epsilon
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DiagramFlossEquations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range compareanalysis.DiagramFlossEquations {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = compareanalysis.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", compareanalysis.IsExpanded)
		res.valueBool = compareanalysis.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (complexity *Complexity) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = complexity.Name
	case "Strength":
		res.valueString = fmt.Sprintf("%f", complexity.Strength)
		res.valueFloat = complexity.Strength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Description":
		res.valueString = complexity.Description
	case "ComputedPrefix":
		res.valueString = complexity.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", complexity.IsExpanded)
		res.valueBool = complexity.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (diagramflossequation *DiagramFlossEquation) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramflossequation.Name
	case "Description":
		res.valueString = diagramflossequation.Description
	case "Scale":
		res.valueString = fmt.Sprintf("%f", diagramflossequation.Scale)
		res.valueFloat = diagramflossequation.Scale
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "FontSize":
		enum := diagramflossequation.FontSize
		res.valueString = enum.ToCodeString()
	case "ComputedPrefix":
		res.valueString = diagramflossequation.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.IsExpanded)
		res.valueBool = diagramflossequation.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.IsChecked)
		res.valueBool = diagramflossequation.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.IsEditable_)
		res.valueBool = diagramflossequation.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInDelta3ColumnsMode":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.IsInDelta3ColumnsMode)
		res.valueBool = diagramflossequation.IsInDelta3ColumnsMode
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AreQuantitativeElementsVisible":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.AreQuantitativeElementsVisible)
		res.valueBool = diagramflossequation.AreQuantitativeElementsVisible
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AreSubsystemsVisible":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.AreSubsystemsVisible)
		res.valueBool = diagramflossequation.AreSubsystemsVisible
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AreCommonElementsHidden":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.AreCommonElementsHidden)
		res.valueBool = diagramflossequation.AreCommonElementsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AreCPEArrowsVisible":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.AreCPEArrowsVisible)
		res.valueBool = diagramflossequation.AreCPEArrowsVisible
		res.GongFieldValueType = GongFieldValueTypeBool
	case "AreColumnTitlesVisible":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.AreColumnTitlesVisible)
		res.valueBool = diagramflossequation.AreColumnTitlesVisible
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramflossequation.Width)
		res.valueFloat = diagramflossequation.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramflossequation.Height)
		res.valueFloat = diagramflossequation.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagramflossequation.DefaultBoxWidth)
		res.valueFloat = diagramflossequation.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagramflossequation.DefaultBoxHeigth)
		res.valueFloat = diagramflossequation.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Note_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramflossequation.Note_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NoteComplexityShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramflossequation.NoteComplexityShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NotePerformanceShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramflossequation.NotePerformanceShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NoteEffortShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramflossequation.NoteEffortShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsNotesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.IsNotesNodeExpanded)
		res.valueBool = diagramflossequation.IsNotesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NotesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramflossequation.NotesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsComplexitysNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.IsComplexitysNodeExpanded)
		res.valueBool = diagramflossequation.IsComplexitysNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComplexitysWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsPerformancesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.IsPerformancesNodeExpanded)
		res.valueBool = diagramflossequation.IsPerformancesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PerformancesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsEffortsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramflossequation.IsEffortsNodeExpanded)
		res.valueBool = diagramflossequation.IsEffortsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "EffortsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramflossequation.EffortsWhoseNodeIsExpanded {
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

func (effort *Effort) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = effort.Name
	case "Strength":
		res.valueString = fmt.Sprintf("%f", effort.Strength)
		res.valueFloat = effort.Strength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Description":
		res.valueString = effort.Description
	case "ComputedPrefix":
		res.valueString = effort.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", effort.IsExpanded)
		res.valueBool = effort.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (library *Library) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = library.Name
	case "Description":
		res.valueString = library.Description
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsExpanded)
		res.valueBool = library.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "RootSystems":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootSystems {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootComplexitys":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootComplexitys {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootPerformances":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootPerformances {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootEfforts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootEfforts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootCompareAnalysis":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootCompareAnalysis {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootNotes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootNotes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsRootLibrary":
		res.valueString = fmt.Sprintf("%t", library.IsRootLibrary)
		res.valueBool = library.IsRootLibrary
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsSubLibrariesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsSubLibrariesNodeExpanded)
		res.valueBool = library.IsSubLibrariesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubLibrariesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SubLibrariesWhoseNodeIsExpanded {
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
	case "IsSystemsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsSystemsNodeExpanded)
		res.valueBool = library.IsSystemsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SystemsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.SystemsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsComplexitysNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsComplexitysNodeExpanded)
		res.valueBool = library.IsComplexitysNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComplexitysWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.ComplexitysWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsPerformancesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsPerformancesNodeExpanded)
		res.valueBool = library.IsPerformancesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PerformancesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.PerformancesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsEffortsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsEffortsNodeExpanded)
		res.valueBool = library.IsEffortsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "EffortsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.EffortsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsCompareAnalysisNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsCompareAnalysisNodeExpanded)
		res.valueBool = library.IsCompareAnalysisNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "CompareAnalysisWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.CompareAnalysisWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsNotesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsNotesNodeExpanded)
		res.valueBool = library.IsNotesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NotesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.NotesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsExpandedTmp":
		res.valueString = fmt.Sprintf("%t", library.IsExpandedTmp)
		res.valueBool = library.IsExpandedTmp
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (note *Note) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = note.Name
	case "Description":
		res.valueString = note.Description
	case "Complexities":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Complexities {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Performances":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Performances {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Efforts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Efforts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ComputedPrefix":
		res.valueString = note.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsExpanded)
		res.valueBool = note.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsComplexitysNodeExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsComplexitysNodeExpanded)
		res.valueBool = note.IsComplexitysNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsPerformancesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsPerformancesNodeExpanded)
		res.valueBool = note.IsPerformancesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEffortsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsEffortsNodeExpanded)
		res.valueBool = note.IsEffortsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (notecomplexityshape *NoteComplexityShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notecomplexityshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notecomplexityshape.Note != nil {
			res.valueString = notecomplexityshape.Note.Name
			res.ids = notecomplexityshape.Note.GongGetUUID(stage)
		}
	case "Complexity":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notecomplexityshape.Complexity != nil {
			res.valueString = notecomplexityshape.Complexity.Name
			res.ids = notecomplexityshape.Complexity.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", notecomplexityshape.StartRatio)
		res.valueFloat = notecomplexityshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", notecomplexityshape.EndRatio)
		res.valueFloat = notecomplexityshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := notecomplexityshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := notecomplexityshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", notecomplexityshape.CornerOffsetRatio)
		res.valueFloat = notecomplexityshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", notecomplexityshape.IsHidden)
		res.valueBool = notecomplexityshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (noteeffortshape *NoteEffortShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = noteeffortshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteeffortshape.Note != nil {
			res.valueString = noteeffortshape.Note.Name
			res.ids = noteeffortshape.Note.GongGetUUID(stage)
		}
	case "Effort":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteeffortshape.Effort != nil {
			res.valueString = noteeffortshape.Effort.Name
			res.ids = noteeffortshape.Effort.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", noteeffortshape.StartRatio)
		res.valueFloat = noteeffortshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", noteeffortshape.EndRatio)
		res.valueFloat = noteeffortshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := noteeffortshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := noteeffortshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", noteeffortshape.CornerOffsetRatio)
		res.valueFloat = noteeffortshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", noteeffortshape.IsHidden)
		res.valueBool = noteeffortshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (noteperformanceshape *NotePerformanceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = noteperformanceshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteperformanceshape.Note != nil {
			res.valueString = noteperformanceshape.Note.Name
			res.ids = noteperformanceshape.Note.GongGetUUID(stage)
		}
	case "Performance":
		res.GongFieldValueType = GongFieldValueTypePointer
		if noteperformanceshape.Performance != nil {
			res.valueString = noteperformanceshape.Performance.Name
			res.ids = noteperformanceshape.Performance.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", noteperformanceshape.StartRatio)
		res.valueFloat = noteperformanceshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", noteperformanceshape.EndRatio)
		res.valueFloat = noteperformanceshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := noteperformanceshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := noteperformanceshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", noteperformanceshape.CornerOffsetRatio)
		res.valueFloat = noteperformanceshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", noteperformanceshape.IsHidden)
		res.valueBool = noteperformanceshape.IsHidden
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

func (performance *Performance) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = performance.Name
	case "Strength":
		res.valueString = fmt.Sprintf("%f", performance.Strength)
		res.valueFloat = performance.Strength
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Description":
		res.valueString = performance.Description
	case "ComputedPrefix":
		res.valueString = performance.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", performance.IsExpanded)
		res.valueBool = performance.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (system *System) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = system.Name
	case "Description":
		res.valueString = system.Description
	case "Complexities":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.Complexities {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Performances":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.Performances {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Efforts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.Efforts {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SubSystems":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.SubSystems {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AreCPEsCompoundedFromSubSystems":
		res.valueString = fmt.Sprintf("%t", system.AreCPEsCompoundedFromSubSystems)
		res.valueBool = system.AreCPEsCompoundedFromSubSystems
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = system.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsExpanded)
		res.valueBool = system.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SVG_Path":
		res.valueString = system.SVG_Path
	case "InverseAppliedScaling":
		res.valueString = fmt.Sprintf("%f", system.InverseAppliedScaling)
		res.valueFloat = system.InverseAppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DiagramFlossEquations":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramFlossEquations {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubSystemNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsSubSystemNodeExpanded)
		res.valueBool = system.IsSubSystemNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsComplexitysNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsComplexitysNodeExpanded)
		res.valueBool = system.IsComplexitysNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComplexitysWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.ComplexitysWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsPerformancesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsPerformancesNodeExpanded)
		res.valueBool = system.IsPerformancesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PerformancesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.PerformancesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsEffortsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", system.IsEffortsNodeExpanded)
		res.valueBool = system.IsEffortsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "EffortsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.EffortsWhoseNodeIsExpanded {
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

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (compareanalysis *CompareAnalysis) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		compareanalysis.Name = value.GetValueString()
	case "FromSystem":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			compareanalysis.FromSystem = nil
			for __instance__ := range stage.Systems {
				if stage.System_stagedOrder[__instance__] == uint(id) {
					compareanalysis.FromSystem = __instance__
					break
				}
			}
		}
	case "ToSystem":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			compareanalysis.ToSystem = nil
			for __instance__ := range stage.Systems {
				if stage.System_stagedOrder[__instance__] == uint(id) {
					compareanalysis.ToSystem = __instance__
					break
				}
			}
		}
	case "Mu":
		compareanalysis.Mu = value.GetValueFloat()
	case "Epsilon":
		compareanalysis.Epsilon = value.GetValueFloat()
	case "DiagramFlossEquations":
		compareanalysis.DiagramFlossEquations = make([]*DiagramFlossEquation, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramFlossEquations {
					if stage.DiagramFlossEquation_stagedOrder[__instance__] == uint(id) {
						compareanalysis.DiagramFlossEquations = append(compareanalysis.DiagramFlossEquations, __instance__)
						break
					}
				}
			}
		}
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded = make([]*DiagramFlossEquation, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramFlossEquations {
					if stage.DiagramFlossEquation_stagedOrder[__instance__] == uint(id) {
						compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded = append(compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ComputedPrefix":
		compareanalysis.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		compareanalysis.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (complexity *Complexity) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		complexity.Name = value.GetValueString()
	case "Strength":
		complexity.Strength = value.GetValueFloat()
	case "Description":
		complexity.Description = value.GetValueString()
	case "ComputedPrefix":
		complexity.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		complexity.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (diagramflossequation *DiagramFlossEquation) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagramflossequation.Name = value.GetValueString()
	case "Description":
		diagramflossequation.Description = value.GetValueString()
	case "Scale":
		diagramflossequation.Scale = value.GetValueFloat()
	case "FontSize":
		diagramflossequation.FontSize.FromCodeString(value.GetValueString())
	case "ComputedPrefix":
		diagramflossequation.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagramflossequation.IsExpanded = value.GetValueBool()
	case "IsChecked":
		diagramflossequation.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagramflossequation.IsEditable_ = value.GetValueBool()
	case "IsInDelta3ColumnsMode":
		diagramflossequation.IsInDelta3ColumnsMode = value.GetValueBool()
	case "AreQuantitativeElementsVisible":
		diagramflossequation.AreQuantitativeElementsVisible = value.GetValueBool()
	case "AreSubsystemsVisible":
		diagramflossequation.AreSubsystemsVisible = value.GetValueBool()
	case "AreCommonElementsHidden":
		diagramflossequation.AreCommonElementsHidden = value.GetValueBool()
	case "AreCPEArrowsVisible":
		diagramflossequation.AreCPEArrowsVisible = value.GetValueBool()
	case "AreColumnTitlesVisible":
		diagramflossequation.AreColumnTitlesVisible = value.GetValueBool()
	case "Width":
		diagramflossequation.Width = value.GetValueFloat()
	case "Height":
		diagramflossequation.Height = value.GetValueFloat()
	case "DefaultBoxWidth":
		diagramflossequation.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagramflossequation.DefaultBoxHeigth = value.GetValueFloat()
	case "Note_Shapes":
		diagramflossequation.Note_Shapes = make([]*NoteShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteShapes {
					if stage.NoteShape_stagedOrder[__instance__] == uint(id) {
						diagramflossequation.Note_Shapes = append(diagramflossequation.Note_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "NoteComplexityShapes":
		diagramflossequation.NoteComplexityShapes = make([]*NoteComplexityShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteComplexityShapes {
					if stage.NoteComplexityShape_stagedOrder[__instance__] == uint(id) {
						diagramflossequation.NoteComplexityShapes = append(diagramflossequation.NoteComplexityShapes, __instance__)
						break
					}
				}
			}
		}
	case "NotePerformanceShapes":
		diagramflossequation.NotePerformanceShapes = make([]*NotePerformanceShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NotePerformanceShapes {
					if stage.NotePerformanceShape_stagedOrder[__instance__] == uint(id) {
						diagramflossequation.NotePerformanceShapes = append(diagramflossequation.NotePerformanceShapes, __instance__)
						break
					}
				}
			}
		}
	case "NoteEffortShapes":
		diagramflossequation.NoteEffortShapes = make([]*NoteEffortShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteEffortShapes {
					if stage.NoteEffortShape_stagedOrder[__instance__] == uint(id) {
						diagramflossequation.NoteEffortShapes = append(diagramflossequation.NoteEffortShapes, __instance__)
						break
					}
				}
			}
		}
	case "IsNotesNodeExpanded":
		diagramflossequation.IsNotesNodeExpanded = value.GetValueBool()
	case "NotesWhoseNodeIsExpanded":
		diagramflossequation.NotesWhoseNodeIsExpanded = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						diagramflossequation.NotesWhoseNodeIsExpanded = append(diagramflossequation.NotesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsComplexitysNodeExpanded":
		diagramflossequation.IsComplexitysNodeExpanded = value.GetValueBool()
	case "ComplexitysWhoseNodeIsExpanded":
		diagramflossequation.ComplexitysWhoseNodeIsExpanded = make([]*Complexity, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Complexitys {
					if stage.Complexity_stagedOrder[__instance__] == uint(id) {
						diagramflossequation.ComplexitysWhoseNodeIsExpanded = append(diagramflossequation.ComplexitysWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsPerformancesNodeExpanded":
		diagramflossequation.IsPerformancesNodeExpanded = value.GetValueBool()
	case "PerformancesWhoseNodeIsExpanded":
		diagramflossequation.PerformancesWhoseNodeIsExpanded = make([]*Performance, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Performances {
					if stage.Performance_stagedOrder[__instance__] == uint(id) {
						diagramflossequation.PerformancesWhoseNodeIsExpanded = append(diagramflossequation.PerformancesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsEffortsNodeExpanded":
		diagramflossequation.IsEffortsNodeExpanded = value.GetValueBool()
	case "EffortsWhoseNodeIsExpanded":
		diagramflossequation.EffortsWhoseNodeIsExpanded = make([]*Effort, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Efforts {
					if stage.Effort_stagedOrder[__instance__] == uint(id) {
						diagramflossequation.EffortsWhoseNodeIsExpanded = append(diagramflossequation.EffortsWhoseNodeIsExpanded, __instance__)
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

func (effort *Effort) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		effort.Name = value.GetValueString()
	case "Strength":
		effort.Strength = value.GetValueFloat()
	case "Description":
		effort.Description = value.GetValueString()
	case "ComputedPrefix":
		effort.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		effort.IsExpanded = value.GetValueBool()
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
	case "Description":
		library.Description = value.GetValueString()
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		library.IsExpanded = value.GetValueBool()
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
	case "RootSystems":
		library.RootSystems = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						library.RootSystems = append(library.RootSystems, __instance__)
						break
					}
				}
			}
		}
	case "RootComplexitys":
		library.RootComplexitys = make([]*Complexity, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Complexitys {
					if stage.Complexity_stagedOrder[__instance__] == uint(id) {
						library.RootComplexitys = append(library.RootComplexitys, __instance__)
						break
					}
				}
			}
		}
	case "RootPerformances":
		library.RootPerformances = make([]*Performance, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Performances {
					if stage.Performance_stagedOrder[__instance__] == uint(id) {
						library.RootPerformances = append(library.RootPerformances, __instance__)
						break
					}
				}
			}
		}
	case "RootEfforts":
		library.RootEfforts = make([]*Effort, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Efforts {
					if stage.Effort_stagedOrder[__instance__] == uint(id) {
						library.RootEfforts = append(library.RootEfforts, __instance__)
						break
					}
				}
			}
		}
	case "RootCompareAnalysis":
		library.RootCompareAnalysis = make([]*CompareAnalysis, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.CompareAnalysiss {
					if stage.CompareAnalysis_stagedOrder[__instance__] == uint(id) {
						library.RootCompareAnalysis = append(library.RootCompareAnalysis, __instance__)
						break
					}
				}
			}
		}
	case "RootNotes":
		library.RootNotes = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						library.RootNotes = append(library.RootNotes, __instance__)
						break
					}
				}
			}
		}
	case "IsRootLibrary":
		library.IsRootLibrary = value.GetValueBool()
	case "IsSubLibrariesNodeExpanded":
		library.IsSubLibrariesNodeExpanded = value.GetValueBool()
	case "SubLibrariesWhoseNodeIsExpanded":
		library.SubLibrariesWhoseNodeIsExpanded = make([]*Library, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Librarys {
					if stage.Library_stagedOrder[__instance__] == uint(id) {
						library.SubLibrariesWhoseNodeIsExpanded = append(library.SubLibrariesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "NbPixPerCharacter":
		library.NbPixPerCharacter = value.GetValueFloat()
	case "LogoSVGFile":
		library.LogoSVGFile = value.GetValueString()
	case "IsSystemsNodeExpanded":
		library.IsSystemsNodeExpanded = value.GetValueBool()
	case "SystemsWhoseNodeIsExpanded":
		library.SystemsWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						library.SystemsWhoseNodeIsExpanded = append(library.SystemsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsComplexitysNodeExpanded":
		library.IsComplexitysNodeExpanded = value.GetValueBool()
	case "ComplexitysWhoseNodeIsExpanded":
		library.ComplexitysWhoseNodeIsExpanded = make([]*Complexity, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Complexitys {
					if stage.Complexity_stagedOrder[__instance__] == uint(id) {
						library.ComplexitysWhoseNodeIsExpanded = append(library.ComplexitysWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsPerformancesNodeExpanded":
		library.IsPerformancesNodeExpanded = value.GetValueBool()
	case "PerformancesWhoseNodeIsExpanded":
		library.PerformancesWhoseNodeIsExpanded = make([]*Performance, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Performances {
					if stage.Performance_stagedOrder[__instance__] == uint(id) {
						library.PerformancesWhoseNodeIsExpanded = append(library.PerformancesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsEffortsNodeExpanded":
		library.IsEffortsNodeExpanded = value.GetValueBool()
	case "EffortsWhoseNodeIsExpanded":
		library.EffortsWhoseNodeIsExpanded = make([]*Effort, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Efforts {
					if stage.Effort_stagedOrder[__instance__] == uint(id) {
						library.EffortsWhoseNodeIsExpanded = append(library.EffortsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsCompareAnalysisNodeExpanded":
		library.IsCompareAnalysisNodeExpanded = value.GetValueBool()
	case "CompareAnalysisWhoseNodeIsExpanded":
		library.CompareAnalysisWhoseNodeIsExpanded = make([]*CompareAnalysis, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.CompareAnalysiss {
					if stage.CompareAnalysis_stagedOrder[__instance__] == uint(id) {
						library.CompareAnalysisWhoseNodeIsExpanded = append(library.CompareAnalysisWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsNotesNodeExpanded":
		library.IsNotesNodeExpanded = value.GetValueBool()
	case "NotesWhoseNodeIsExpanded":
		library.NotesWhoseNodeIsExpanded = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						library.NotesWhoseNodeIsExpanded = append(library.NotesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsExpandedTmp":
		library.IsExpandedTmp = value.GetValueBool()
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
	case "Description":
		note.Description = value.GetValueString()
	case "Complexities":
		note.Complexities = make([]*Complexity, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Complexitys {
					if stage.Complexity_stagedOrder[__instance__] == uint(id) {
						note.Complexities = append(note.Complexities, __instance__)
						break
					}
				}
			}
		}
	case "Performances":
		note.Performances = make([]*Performance, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Performances {
					if stage.Performance_stagedOrder[__instance__] == uint(id) {
						note.Performances = append(note.Performances, __instance__)
						break
					}
				}
			}
		}
	case "Efforts":
		note.Efforts = make([]*Effort, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Efforts {
					if stage.Effort_stagedOrder[__instance__] == uint(id) {
						note.Efforts = append(note.Efforts, __instance__)
						break
					}
				}
			}
		}
	case "ComputedPrefix":
		note.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		note.IsExpanded = value.GetValueBool()
	case "IsComplexitysNodeExpanded":
		note.IsComplexitysNodeExpanded = value.GetValueBool()
	case "IsPerformancesNodeExpanded":
		note.IsPerformancesNodeExpanded = value.GetValueBool()
	case "IsEffortsNodeExpanded":
		note.IsEffortsNodeExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (notecomplexityshape *NoteComplexityShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		notecomplexityshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notecomplexityshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					notecomplexityshape.Note = __instance__
					break
				}
			}
		}
	case "Complexity":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notecomplexityshape.Complexity = nil
			for __instance__ := range stage.Complexitys {
				if stage.Complexity_stagedOrder[__instance__] == uint(id) {
					notecomplexityshape.Complexity = __instance__
					break
				}
			}
		}
	case "StartRatio":
		notecomplexityshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		notecomplexityshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		notecomplexityshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		notecomplexityshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		notecomplexityshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		notecomplexityshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (noteeffortshape *NoteEffortShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		noteeffortshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteeffortshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					noteeffortshape.Note = __instance__
					break
				}
			}
		}
	case "Effort":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteeffortshape.Effort = nil
			for __instance__ := range stage.Efforts {
				if stage.Effort_stagedOrder[__instance__] == uint(id) {
					noteeffortshape.Effort = __instance__
					break
				}
			}
		}
	case "StartRatio":
		noteeffortshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		noteeffortshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		noteeffortshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		noteeffortshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		noteeffortshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		noteeffortshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (noteperformanceshape *NotePerformanceShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		noteperformanceshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteperformanceshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					noteperformanceshape.Note = __instance__
					break
				}
			}
		}
	case "Performance":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			noteperformanceshape.Performance = nil
			for __instance__ := range stage.Performances {
				if stage.Performance_stagedOrder[__instance__] == uint(id) {
					noteperformanceshape.Performance = __instance__
					break
				}
			}
		}
	case "StartRatio":
		noteperformanceshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		noteperformanceshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		noteperformanceshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		noteperformanceshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		noteperformanceshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		noteperformanceshape.IsHidden = value.GetValueBool()
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

func (performance *Performance) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		performance.Name = value.GetValueString()
	case "Strength":
		performance.Strength = value.GetValueFloat()
	case "Description":
		performance.Description = value.GetValueString()
	case "ComputedPrefix":
		performance.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		performance.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (system *System) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		system.Name = value.GetValueString()
	case "Description":
		system.Description = value.GetValueString()
	case "Complexities":
		system.Complexities = make([]*Complexity, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Complexitys {
					if stage.Complexity_stagedOrder[__instance__] == uint(id) {
						system.Complexities = append(system.Complexities, __instance__)
						break
					}
				}
			}
		}
	case "Performances":
		system.Performances = make([]*Performance, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Performances {
					if stage.Performance_stagedOrder[__instance__] == uint(id) {
						system.Performances = append(system.Performances, __instance__)
						break
					}
				}
			}
		}
	case "Efforts":
		system.Efforts = make([]*Effort, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Efforts {
					if stage.Effort_stagedOrder[__instance__] == uint(id) {
						system.Efforts = append(system.Efforts, __instance__)
						break
					}
				}
			}
		}
	case "SubSystems":
		system.SubSystems = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						system.SubSystems = append(system.SubSystems, __instance__)
						break
					}
				}
			}
		}
	case "AreCPEsCompoundedFromSubSystems":
		system.AreCPEsCompoundedFromSubSystems = value.GetValueBool()
	case "ComputedPrefix":
		system.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		system.IsExpanded = value.GetValueBool()
	case "SVG_Path":
		system.SVG_Path = value.GetValueString()
	case "InverseAppliedScaling":
		system.InverseAppliedScaling = value.GetValueFloat()
	case "DiagramFlossEquations":
		system.DiagramFlossEquations = make([]*DiagramFlossEquation, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramFlossEquations {
					if stage.DiagramFlossEquation_stagedOrder[__instance__] == uint(id) {
						system.DiagramFlossEquations = append(system.DiagramFlossEquations, __instance__)
						break
					}
				}
			}
		}
	case "DiagramFlossEquationsWhoseNodeIsExpanded":
		system.DiagramFlossEquationsWhoseNodeIsExpanded = make([]*DiagramFlossEquation, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramFlossEquations {
					if stage.DiagramFlossEquation_stagedOrder[__instance__] == uint(id) {
						system.DiagramFlossEquationsWhoseNodeIsExpanded = append(system.DiagramFlossEquationsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsSubSystemNodeExpanded":
		system.IsSubSystemNodeExpanded = value.GetValueBool()
	case "IsComplexitysNodeExpanded":
		system.IsComplexitysNodeExpanded = value.GetValueBool()
	case "ComplexitysWhoseNodeIsExpanded":
		system.ComplexitysWhoseNodeIsExpanded = make([]*Complexity, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Complexitys {
					if stage.Complexity_stagedOrder[__instance__] == uint(id) {
						system.ComplexitysWhoseNodeIsExpanded = append(system.ComplexitysWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsPerformancesNodeExpanded":
		system.IsPerformancesNodeExpanded = value.GetValueBool()
	case "PerformancesWhoseNodeIsExpanded":
		system.PerformancesWhoseNodeIsExpanded = make([]*Performance, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Performances {
					if stage.Performance_stagedOrder[__instance__] == uint(id) {
						system.PerformancesWhoseNodeIsExpanded = append(system.PerformancesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsEffortsNodeExpanded":
		system.IsEffortsNodeExpanded = value.GetValueBool()
	case "EffortsWhoseNodeIsExpanded":
		system.EffortsWhoseNodeIsExpanded = make([]*Effort, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Efforts {
					if stage.Effort_stagedOrder[__instance__] == uint(id) {
						system.EffortsWhoseNodeIsExpanded = append(system.EffortsWhoseNodeIsExpanded, __instance__)
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

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (compareanalysis *CompareAnalysis) GongGetGongstructName() string {
	return "CompareAnalysis"
}

func (complexity *Complexity) GongGetGongstructName() string {
	return "Complexity"
}

func (diagramflossequation *DiagramFlossEquation) GongGetGongstructName() string {
	return "DiagramFlossEquation"
}

func (effort *Effort) GongGetGongstructName() string {
	return "Effort"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (note *Note) GongGetGongstructName() string {
	return "Note"
}

func (notecomplexityshape *NoteComplexityShape) GongGetGongstructName() string {
	return "NoteComplexityShape"
}

func (noteeffortshape *NoteEffortShape) GongGetGongstructName() string {
	return "NoteEffortShape"
}

func (noteperformanceshape *NotePerformanceShape) GongGetGongstructName() string {
	return "NotePerformanceShape"
}

func (noteshape *NoteShape) GongGetGongstructName() string {
	return "NoteShape"
}

func (performance *Performance) GongGetGongstructName() string {
	return "Performance"
}

func (system *System) GongGetGongstructName() string {
	return "System"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.CompareAnalysiss_mapString = make(map[string]*CompareAnalysis)
	for compareanalysis := range stage.CompareAnalysiss {
		stage.CompareAnalysiss_mapString[compareanalysis.Name] = compareanalysis
	}

	stage.Complexitys_mapString = make(map[string]*Complexity)
	for complexity := range stage.Complexitys {
		stage.Complexitys_mapString[complexity.Name] = complexity
	}

	stage.DiagramFlossEquations_mapString = make(map[string]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		stage.DiagramFlossEquations_mapString[diagramflossequation.Name] = diagramflossequation
	}

	stage.Efforts_mapString = make(map[string]*Effort)
	for effort := range stage.Efforts {
		stage.Efforts_mapString[effort.Name] = effort
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Notes_mapString = make(map[string]*Note)
	for note := range stage.Notes {
		stage.Notes_mapString[note.Name] = note
	}

	stage.NoteComplexityShapes_mapString = make(map[string]*NoteComplexityShape)
	for notecomplexityshape := range stage.NoteComplexityShapes {
		stage.NoteComplexityShapes_mapString[notecomplexityshape.Name] = notecomplexityshape
	}

	stage.NoteEffortShapes_mapString = make(map[string]*NoteEffortShape)
	for noteeffortshape := range stage.NoteEffortShapes {
		stage.NoteEffortShapes_mapString[noteeffortshape.Name] = noteeffortshape
	}

	stage.NotePerformanceShapes_mapString = make(map[string]*NotePerformanceShape)
	for noteperformanceshape := range stage.NotePerformanceShapes {
		stage.NotePerformanceShapes_mapString[noteperformanceshape.Name] = noteperformanceshape
	}

	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	for noteshape := range stage.NoteShapes {
		stage.NoteShapes_mapString[noteshape.Name] = noteshape
	}

	stage.Performances_mapString = make(map[string]*Performance)
	for performance := range stage.Performances {
		stage.Performances_mapString[performance.Name] = performance
	}

	stage.Systems_mapString = make(map[string]*System)
	for system := range stage.Systems {
		stage.Systems_mapString[system.Name] = system
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
