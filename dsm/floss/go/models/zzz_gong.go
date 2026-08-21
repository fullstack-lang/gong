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

	ComplexityShapes                map[*ComplexityShape]struct{}
	ComplexityShapes_instance       map[*ComplexityShape]*ComplexityShape
	ComplexityShapes_mapString      map[string]*ComplexityShape
	ComplexityShapeOrder            uint
	ComplexityShape_stagedOrder     map[*ComplexityShape]uint
	ComplexityShape_orderStaged     map[uint]*ComplexityShape
	ComplexityShapes_reference      map[*ComplexityShape]*ComplexityShape
	ComplexityShapes_referenceOrder map[*ComplexityShape]uint

	// insertion point for slice of pointers maps
	OnAfterComplexityShapeCreateCallback OnAfterCreateInterface[ComplexityShape]
	OnAfterComplexityShapeUpdateCallback OnAfterUpdateInterface[ComplexityShape]
	OnAfterComplexityShapeDeleteCallback OnAfterDeleteInterface[ComplexityShape]
	OnAfterComplexityShapeReadCallback   OnAfterReadInterface[ComplexityShape]

	DiagramFlosss                map[*DiagramFloss]struct{}
	DiagramFlosss_instance       map[*DiagramFloss]*DiagramFloss
	DiagramFlosss_mapString      map[string]*DiagramFloss
	DiagramFlossOrder            uint
	DiagramFloss_stagedOrder     map[*DiagramFloss]uint
	DiagramFloss_orderStaged     map[uint]*DiagramFloss
	DiagramFlosss_reference      map[*DiagramFloss]*DiagramFloss
	DiagramFlosss_referenceOrder map[*DiagramFloss]uint

	// insertion point for slice of pointers maps
	DiagramFloss_System_Shapes_reverseMap map[*SystemShape]*DiagramFloss

	DiagramFloss_SystemsWhoseNodeIsExpanded_reverseMap map[*System]*DiagramFloss

	DiagramFloss_Complexity_Shapes_reverseMap map[*ComplexityShape]*DiagramFloss

	DiagramFloss_ComplexitysWhoseNodeIsExpanded_reverseMap map[*Complexity]*DiagramFloss

	DiagramFloss_Performance_Shapes_reverseMap map[*PerformanceShape]*DiagramFloss

	DiagramFloss_PerformancesWhoseNodeIsExpanded_reverseMap map[*Performance]*DiagramFloss

	DiagramFloss_Effort_Shapes_reverseMap map[*EffortShape]*DiagramFloss

	DiagramFloss_EffortsWhoseNodeIsExpanded_reverseMap map[*Effort]*DiagramFloss

	DiagramFloss_Note_Shapes_reverseMap map[*NoteShape]*DiagramFloss

	DiagramFloss_NoteComplexityShapes_reverseMap map[*NoteComplexityShape]*DiagramFloss

	DiagramFloss_NotePerformanceShapes_reverseMap map[*NotePerformanceShape]*DiagramFloss

	DiagramFloss_NoteEffortShapes_reverseMap map[*NoteEffortShape]*DiagramFloss

	DiagramFloss_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*DiagramFloss

	OnAfterDiagramFlossCreateCallback OnAfterCreateInterface[DiagramFloss]
	OnAfterDiagramFlossUpdateCallback OnAfterUpdateInterface[DiagramFloss]
	OnAfterDiagramFlossDeleteCallback OnAfterDeleteInterface[DiagramFloss]
	OnAfterDiagramFlossReadCallback   OnAfterReadInterface[DiagramFloss]

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

	EffortShapes                map[*EffortShape]struct{}
	EffortShapes_instance       map[*EffortShape]*EffortShape
	EffortShapes_mapString      map[string]*EffortShape
	EffortShapeOrder            uint
	EffortShape_stagedOrder     map[*EffortShape]uint
	EffortShape_orderStaged     map[uint]*EffortShape
	EffortShapes_reference      map[*EffortShape]*EffortShape
	EffortShapes_referenceOrder map[*EffortShape]uint

	// insertion point for slice of pointers maps
	OnAfterEffortShapeCreateCallback OnAfterCreateInterface[EffortShape]
	OnAfterEffortShapeUpdateCallback OnAfterUpdateInterface[EffortShape]
	OnAfterEffortShapeDeleteCallback OnAfterDeleteInterface[EffortShape]
	OnAfterEffortShapeReadCallback   OnAfterReadInterface[EffortShape]

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

	PerformanceShapes                map[*PerformanceShape]struct{}
	PerformanceShapes_instance       map[*PerformanceShape]*PerformanceShape
	PerformanceShapes_mapString      map[string]*PerformanceShape
	PerformanceShapeOrder            uint
	PerformanceShape_stagedOrder     map[*PerformanceShape]uint
	PerformanceShape_orderStaged     map[uint]*PerformanceShape
	PerformanceShapes_reference      map[*PerformanceShape]*PerformanceShape
	PerformanceShapes_referenceOrder map[*PerformanceShape]uint

	// insertion point for slice of pointers maps
	OnAfterPerformanceShapeCreateCallback OnAfterCreateInterface[PerformanceShape]
	OnAfterPerformanceShapeUpdateCallback OnAfterUpdateInterface[PerformanceShape]
	OnAfterPerformanceShapeDeleteCallback OnAfterDeleteInterface[PerformanceShape]
	OnAfterPerformanceShapeReadCallback   OnAfterReadInterface[PerformanceShape]

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

	System_DiagramFlosses_reverseMap map[*DiagramFloss]*System

	System_DiagramFlossWhoseNodeIsExpanded_reverseMap map[*DiagramFloss]*System

	System_SubSystemes_reverseMap map[*System]*System

	System_ComplexitysWhoseNodeIsExpanded_reverseMap map[*Complexity]*System

	System_PerformancesWhoseNodeIsExpanded_reverseMap map[*Performance]*System

	System_EffortsWhoseNodeIsExpanded_reverseMap map[*Effort]*System

	OnAfterSystemCreateCallback OnAfterCreateInterface[System]
	OnAfterSystemUpdateCallback OnAfterUpdateInterface[System]
	OnAfterSystemDeleteCallback OnAfterDeleteInterface[System]
	OnAfterSystemReadCallback   OnAfterReadInterface[System]

	SystemShapes                map[*SystemShape]struct{}
	SystemShapes_instance       map[*SystemShape]*SystemShape
	SystemShapes_mapString      map[string]*SystemShape
	SystemShapeOrder            uint
	SystemShape_stagedOrder     map[*SystemShape]uint
	SystemShape_orderStaged     map[uint]*SystemShape
	SystemShapes_reference      map[*SystemShape]*SystemShape
	SystemShapes_referenceOrder map[*SystemShape]uint

	// insertion point for slice of pointers maps
	OnAfterSystemShapeCreateCallback OnAfterCreateInterface[SystemShape]
	OnAfterSystemShapeUpdateCallback OnAfterUpdateInterface[SystemShape]
	OnAfterSystemShapeDeleteCallback OnAfterDeleteInterface[SystemShape]
	OnAfterSystemShapeReadCallback   OnAfterReadInterface[SystemShape]

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

	stage.ComplexityShapes_reference = make(map[*ComplexityShape]*ComplexityShape)
	stage.ComplexityShapes_instance = make(map[*ComplexityShape]*ComplexityShape)
	stage.ComplexityShapes_referenceOrder = make(map[*ComplexityShape]uint)

	stage.DiagramFlosss_reference = make(map[*DiagramFloss]*DiagramFloss)
	stage.DiagramFlosss_instance = make(map[*DiagramFloss]*DiagramFloss)
	stage.DiagramFlosss_referenceOrder = make(map[*DiagramFloss]uint)

	stage.DiagramFlossEquations_reference = make(map[*DiagramFlossEquation]*DiagramFlossEquation)
	stage.DiagramFlossEquations_instance = make(map[*DiagramFlossEquation]*DiagramFlossEquation)
	stage.DiagramFlossEquations_referenceOrder = make(map[*DiagramFlossEquation]uint)

	stage.Efforts_reference = make(map[*Effort]*Effort)
	stage.Efforts_instance = make(map[*Effort]*Effort)
	stage.Efforts_referenceOrder = make(map[*Effort]uint)

	stage.EffortShapes_reference = make(map[*EffortShape]*EffortShape)
	stage.EffortShapes_instance = make(map[*EffortShape]*EffortShape)
	stage.EffortShapes_referenceOrder = make(map[*EffortShape]uint)

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

	stage.PerformanceShapes_reference = make(map[*PerformanceShape]*PerformanceShape)
	stage.PerformanceShapes_instance = make(map[*PerformanceShape]*PerformanceShape)
	stage.PerformanceShapes_referenceOrder = make(map[*PerformanceShape]uint)

	stage.Systems_reference = make(map[*System]*System)
	stage.Systems_instance = make(map[*System]*System)
	stage.Systems_referenceOrder = make(map[*System]uint)

	stage.SystemShapes_reference = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_instance = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint)

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

	var maxComplexityShapeOrder uint
	var foundComplexityShape bool
	for _, order := range stage.ComplexityShape_stagedOrder {
		if !foundComplexityShape || order > maxComplexityShapeOrder {
			maxComplexityShapeOrder = order
			foundComplexityShape = true
		}
	}
	if foundComplexityShape {
		stage.ComplexityShapeOrder = maxComplexityShapeOrder + 1
	} else {
		stage.ComplexityShapeOrder = 0
	}

	var maxDiagramFlossOrder uint
	var foundDiagramFloss bool
	for _, order := range stage.DiagramFloss_stagedOrder {
		if !foundDiagramFloss || order > maxDiagramFlossOrder {
			maxDiagramFlossOrder = order
			foundDiagramFloss = true
		}
	}
	if foundDiagramFloss {
		stage.DiagramFlossOrder = maxDiagramFlossOrder + 1
	} else {
		stage.DiagramFlossOrder = 0
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

	var maxEffortShapeOrder uint
	var foundEffortShape bool
	for _, order := range stage.EffortShape_stagedOrder {
		if !foundEffortShape || order > maxEffortShapeOrder {
			maxEffortShapeOrder = order
			foundEffortShape = true
		}
	}
	if foundEffortShape {
		stage.EffortShapeOrder = maxEffortShapeOrder + 1
	} else {
		stage.EffortShapeOrder = 0
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

	var maxPerformanceShapeOrder uint
	var foundPerformanceShape bool
	for _, order := range stage.PerformanceShape_stagedOrder {
		if !foundPerformanceShape || order > maxPerformanceShapeOrder {
			maxPerformanceShapeOrder = order
			foundPerformanceShape = true
		}
	}
	if foundPerformanceShape {
		stage.PerformanceShapeOrder = maxPerformanceShapeOrder + 1
	} else {
		stage.PerformanceShapeOrder = 0
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

	var maxSystemShapeOrder uint
	var foundSystemShape bool
	for _, order := range stage.SystemShape_stagedOrder {
		if !foundSystemShape || order > maxSystemShapeOrder {
			maxSystemShapeOrder = order
			foundSystemShape = true
		}
	}
	if foundSystemShape {
		stage.SystemShapeOrder = maxSystemShapeOrder + 1
	} else {
		stage.SystemShapeOrder = 0
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
	case *ComplexityShape:
		tmp := GetStructInstancesByOrder(stage.ComplexityShapes, stage.ComplexityShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ComplexityShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DiagramFloss:
		tmp := GetStructInstancesByOrder(stage.DiagramFlosss, stage.DiagramFloss_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DiagramFloss implements.
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
	case *EffortShape:
		tmp := GetStructInstancesByOrder(stage.EffortShapes, stage.EffortShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *EffortShape implements.
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
	case *PerformanceShape:
		tmp := GetStructInstancesByOrder(stage.PerformanceShapes, stage.PerformanceShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *PerformanceShape implements.
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
	case *SystemShape:
		tmp := GetStructInstancesByOrder(stage.SystemShapes, stage.SystemShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SystemShape implements.
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
	case "ComplexityShape":
		res = GetNamedStructInstances(stage.ComplexityShapes, stage.ComplexityShape_stagedOrder)
	case "DiagramFloss":
		res = GetNamedStructInstances(stage.DiagramFlosss, stage.DiagramFloss_stagedOrder)
	case "DiagramFlossEquation":
		res = GetNamedStructInstances(stage.DiagramFlossEquations, stage.DiagramFlossEquation_stagedOrder)
	case "Effort":
		res = GetNamedStructInstances(stage.Efforts, stage.Effort_stagedOrder)
	case "EffortShape":
		res = GetNamedStructInstances(stage.EffortShapes, stage.EffortShape_stagedOrder)
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
	case "PerformanceShape":
		res = GetNamedStructInstances(stage.PerformanceShapes, stage.PerformanceShape_stagedOrder)
	case "System":
		res = GetNamedStructInstances(stage.Systems, stage.System_stagedOrder)
	case "SystemShape":
		res = GetNamedStructInstances(stage.SystemShapes, stage.SystemShape_stagedOrder)
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
	CommitComplexityShape(complexityshape *ComplexityShape)
	CheckoutComplexityShape(complexityshape *ComplexityShape)
	CommitDiagramFloss(diagramfloss *DiagramFloss)
	CheckoutDiagramFloss(diagramfloss *DiagramFloss)
	CommitDiagramFlossEquation(diagramflossequation *DiagramFlossEquation)
	CheckoutDiagramFlossEquation(diagramflossequation *DiagramFlossEquation)
	CommitEffort(effort *Effort)
	CheckoutEffort(effort *Effort)
	CommitEffortShape(effortshape *EffortShape)
	CheckoutEffortShape(effortshape *EffortShape)
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
	CommitPerformanceShape(performanceshape *PerformanceShape)
	CheckoutPerformanceShape(performanceshape *PerformanceShape)
	CommitSystem(system *System)
	CheckoutSystem(system *System)
	CommitSystemShape(systemshape *SystemShape)
	CheckoutSystemShape(systemshape *SystemShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		CompareAnalysiss:           make(map[*CompareAnalysis]struct{}),
		CompareAnalysiss_mapString: make(map[string]*CompareAnalysis),

		Complexitys:           make(map[*Complexity]struct{}),
		Complexitys_mapString: make(map[string]*Complexity),

		ComplexityShapes:           make(map[*ComplexityShape]struct{}),
		ComplexityShapes_mapString: make(map[string]*ComplexityShape),

		DiagramFlosss:           make(map[*DiagramFloss]struct{}),
		DiagramFlosss_mapString: make(map[string]*DiagramFloss),

		DiagramFlossEquations:           make(map[*DiagramFlossEquation]struct{}),
		DiagramFlossEquations_mapString: make(map[string]*DiagramFlossEquation),

		Efforts:           make(map[*Effort]struct{}),
		Efforts_mapString: make(map[string]*Effort),

		EffortShapes:           make(map[*EffortShape]struct{}),
		EffortShapes_mapString: make(map[string]*EffortShape),

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

		PerformanceShapes:           make(map[*PerformanceShape]struct{}),
		PerformanceShapes_mapString: make(map[string]*PerformanceShape),

		Systems:           make(map[*System]struct{}),
		Systems_mapString: make(map[string]*System),

		SystemShapes:           make(map[*SystemShape]struct{}),
		SystemShapes_mapString: make(map[string]*SystemShape),

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

		ComplexityShape_stagedOrder: make(map[*ComplexityShape]uint),
		ComplexityShape_orderStaged: make(map[uint]*ComplexityShape),
		ComplexityShapes_reference:  make(map[*ComplexityShape]*ComplexityShape),

		DiagramFloss_stagedOrder: make(map[*DiagramFloss]uint),
		DiagramFloss_orderStaged: make(map[uint]*DiagramFloss),
		DiagramFlosss_reference:  make(map[*DiagramFloss]*DiagramFloss),

		DiagramFlossEquation_stagedOrder: make(map[*DiagramFlossEquation]uint),
		DiagramFlossEquation_orderStaged: make(map[uint]*DiagramFlossEquation),
		DiagramFlossEquations_reference:  make(map[*DiagramFlossEquation]*DiagramFlossEquation),

		Effort_stagedOrder: make(map[*Effort]uint),
		Effort_orderStaged: make(map[uint]*Effort),
		Efforts_reference:  make(map[*Effort]*Effort),

		EffortShape_stagedOrder: make(map[*EffortShape]uint),
		EffortShape_orderStaged: make(map[uint]*EffortShape),
		EffortShapes_reference:  make(map[*EffortShape]*EffortShape),

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

		PerformanceShape_stagedOrder: make(map[*PerformanceShape]uint),
		PerformanceShape_orderStaged: make(map[uint]*PerformanceShape),
		PerformanceShapes_reference:  make(map[*PerformanceShape]*PerformanceShape),

		System_stagedOrder: make(map[*System]uint),
		System_orderStaged: make(map[uint]*System),
		Systems_reference:  make(map[*System]*System),

		SystemShape_stagedOrder: make(map[*SystemShape]uint),
		SystemShape_orderStaged: make(map[uint]*SystemShape),
		SystemShapes_reference:  make(map[*SystemShape]*SystemShape),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"CompareAnalysis": &CompareAnalysisUnmarshaller{},

			"Complexity": &ComplexityUnmarshaller{},

			"ComplexityShape": &ComplexityShapeUnmarshaller{},

			"DiagramFloss": &DiagramFlossUnmarshaller{},

			"DiagramFlossEquation": &DiagramFlossEquationUnmarshaller{},

			"Effort": &EffortUnmarshaller{},

			"EffortShape": &EffortShapeUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteComplexityShape": &NoteComplexityShapeUnmarshaller{},

			"NoteEffortShape": &NoteEffortShapeUnmarshaller{},

			"NotePerformanceShape": &NotePerformanceShapeUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"Performance": &PerformanceUnmarshaller{},

			"PerformanceShape": &PerformanceShapeUnmarshaller{},

			"System": &SystemUnmarshaller{},

			"SystemShape": &SystemShapeUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "CompareAnalysis"},
			{name: "Complexity"},
			{name: "ComplexityShape"},
			{name: "DiagramFloss"},
			{name: "DiagramFlossEquation"},
			{name: "Effort"},
			{name: "EffortShape"},
			{name: "Library"},
			{name: "Note"},
			{name: "NoteComplexityShape"},
			{name: "NoteEffortShape"},
			{name: "NotePerformanceShape"},
			{name: "NoteShape"},
			{name: "Performance"},
			{name: "PerformanceShape"},
			{name: "System"},
			{name: "SystemShape"},
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
	case *ComplexityShape:
		return stage.ComplexityShape_stagedOrder[instance]
	case *DiagramFloss:
		return stage.DiagramFloss_stagedOrder[instance]
	case *DiagramFlossEquation:
		return stage.DiagramFlossEquation_stagedOrder[instance]
	case *Effort:
		return stage.Effort_stagedOrder[instance]
	case *EffortShape:
		return stage.EffortShape_stagedOrder[instance]
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
	case *PerformanceShape:
		return stage.PerformanceShape_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
	case *SystemShape:
		return stage.SystemShape_stagedOrder[instance]
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
	case *ComplexityShape:
		return any(stage.ComplexityShape_orderStaged[order]).(Type)
	case *DiagramFloss:
		return any(stage.DiagramFloss_orderStaged[order]).(Type)
	case *DiagramFlossEquation:
		return any(stage.DiagramFlossEquation_orderStaged[order]).(Type)
	case *Effort:
		return any(stage.Effort_orderStaged[order]).(Type)
	case *EffortShape:
		return any(stage.EffortShape_orderStaged[order]).(Type)
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
	case *PerformanceShape:
		return any(stage.PerformanceShape_orderStaged[order]).(Type)
	case *System:
		return any(stage.System_orderStaged[order]).(Type)
	case *SystemShape:
		return any(stage.SystemShape_orderStaged[order]).(Type)
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
	case *ComplexityShape:
		return stage.ComplexityShape_stagedOrder[instance]
	case *DiagramFloss:
		return stage.DiagramFloss_stagedOrder[instance]
	case *DiagramFlossEquation:
		return stage.DiagramFlossEquation_stagedOrder[instance]
	case *Effort:
		return stage.Effort_stagedOrder[instance]
	case *EffortShape:
		return stage.EffortShape_stagedOrder[instance]
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
	case *PerformanceShape:
		return stage.PerformanceShape_stagedOrder[instance]
	case *System:
		return stage.System_stagedOrder[instance]
	case *SystemShape:
		return stage.SystemShape_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["ComplexityShape"] = len(stage.ComplexityShapes)
	stage.Map_GongStructName_InstancesNb["DiagramFloss"] = len(stage.DiagramFlosss)
	stage.Map_GongStructName_InstancesNb["DiagramFlossEquation"] = len(stage.DiagramFlossEquations)
	stage.Map_GongStructName_InstancesNb["Effort"] = len(stage.Efforts)
	stage.Map_GongStructName_InstancesNb["EffortShape"] = len(stage.EffortShapes)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteComplexityShape"] = len(stage.NoteComplexityShapes)
	stage.Map_GongStructName_InstancesNb["NoteEffortShape"] = len(stage.NoteEffortShapes)
	stage.Map_GongStructName_InstancesNb["NotePerformanceShape"] = len(stage.NotePerformanceShapes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["Performance"] = len(stage.Performances)
	stage.Map_GongStructName_InstancesNb["PerformanceShape"] = len(stage.PerformanceShapes)
	stage.Map_GongStructName_InstancesNb["System"] = len(stage.Systems)
	stage.Map_GongStructName_InstancesNb["SystemShape"] = len(stage.SystemShapes)
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

// Stage puts complexityshape to the model stage
func (complexityshape *ComplexityShape) Stage(stage *Stage) *ComplexityShape {
	if _, ok := stage.ComplexityShapes[complexityshape]; !ok {
		stage.ComplexityShapes[complexityshape] = struct{}{}
		stage.ComplexityShape_stagedOrder[complexityshape] = stage.ComplexityShapeOrder
		stage.ComplexityShape_orderStaged[stage.ComplexityShapeOrder] = complexityshape
		stage.ComplexityShapeOrder++
	}
	stage.ComplexityShapes_mapString[complexityshape.Name] = complexityshape

	return complexityshape
}

// StagePreserveOrder puts complexityshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ComplexityShapeOrder
// - update stage.ComplexityShapeOrder accordingly
func (complexityshape *ComplexityShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ComplexityShapes[complexityshape]; !ok {
		stage.ComplexityShapes[complexityshape] = struct{}{}

		if order > stage.ComplexityShapeOrder {
			stage.ComplexityShapeOrder = order
		}
		stage.ComplexityShape_stagedOrder[complexityshape] = order
		stage.ComplexityShape_orderStaged[order] = complexityshape
		stage.ComplexityShapeOrder++
	}
	stage.ComplexityShapes_mapString[complexityshape.Name] = complexityshape
}

// Unstage removes complexityshape off the model stage
func (complexityshape *ComplexityShape) Unstage(stage *Stage) *ComplexityShape {
	delete(stage.ComplexityShapes, complexityshape)
	// issue1150
	// delete(stage.ComplexityShape_stagedOrder, complexityshape)
	delete(stage.ComplexityShapes_mapString, complexityshape.Name)

	return complexityshape
}

// UnstageVoid removes complexityshape off the model stage
func (complexityshape *ComplexityShape) UnstageVoid(stage *Stage) {
	delete(stage.ComplexityShapes, complexityshape)
	// issue1150
	// delete(stage.ComplexityShape_stagedOrder, complexityshape)
	delete(stage.ComplexityShapes_mapString, complexityshape.Name)
}

// commit complexityshape to the back repo (if it is already staged)
func (complexityshape *ComplexityShape) Commit(stage *Stage) *ComplexityShape {
	if _, ok := stage.ComplexityShapes[complexityshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitComplexityShape(complexityshape)
		}
	}
	return complexityshape
}

func (complexityshape *ComplexityShape) CommitVoid(stage *Stage) {
	complexityshape.Commit(stage)
}

func (complexityshape *ComplexityShape) StageVoid(stage *Stage) {
	complexityshape.Stage(stage)
}

// Checkout complexityshape to the back repo (if it is already staged)
func (complexityshape *ComplexityShape) Checkout(stage *Stage) *ComplexityShape {
	if _, ok := stage.ComplexityShapes[complexityshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutComplexityShape(complexityshape)
		}
	}
	return complexityshape
}

// for satisfaction of GongStruct interface
func (complexityshape *ComplexityShape) GetName() (res string) {
	return complexityshape.Name
}

// for satisfaction of GongStruct interface
func (complexityshape *ComplexityShape) SetName(name string) {
	complexityshape.Name = name
}

// Stage puts diagramfloss to the model stage
func (diagramfloss *DiagramFloss) Stage(stage *Stage) *DiagramFloss {
	if _, ok := stage.DiagramFlosss[diagramfloss]; !ok {
		stage.DiagramFlosss[diagramfloss] = struct{}{}
		stage.DiagramFloss_stagedOrder[diagramfloss] = stage.DiagramFlossOrder
		stage.DiagramFloss_orderStaged[stage.DiagramFlossOrder] = diagramfloss
		stage.DiagramFlossOrder++
	}
	stage.DiagramFlosss_mapString[diagramfloss.Name] = diagramfloss

	return diagramfloss
}

// StagePreserveOrder puts diagramfloss to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramFlossOrder
// - update stage.DiagramFlossOrder accordingly
func (diagramfloss *DiagramFloss) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DiagramFlosss[diagramfloss]; !ok {
		stage.DiagramFlosss[diagramfloss] = struct{}{}

		if order > stage.DiagramFlossOrder {
			stage.DiagramFlossOrder = order
		}
		stage.DiagramFloss_stagedOrder[diagramfloss] = order
		stage.DiagramFloss_orderStaged[order] = diagramfloss
		stage.DiagramFlossOrder++
	}
	stage.DiagramFlosss_mapString[diagramfloss.Name] = diagramfloss
}

// Unstage removes diagramfloss off the model stage
func (diagramfloss *DiagramFloss) Unstage(stage *Stage) *DiagramFloss {
	delete(stage.DiagramFlosss, diagramfloss)
	// issue1150
	// delete(stage.DiagramFloss_stagedOrder, diagramfloss)
	delete(stage.DiagramFlosss_mapString, diagramfloss.Name)

	return diagramfloss
}

// UnstageVoid removes diagramfloss off the model stage
func (diagramfloss *DiagramFloss) UnstageVoid(stage *Stage) {
	delete(stage.DiagramFlosss, diagramfloss)
	// issue1150
	// delete(stage.DiagramFloss_stagedOrder, diagramfloss)
	delete(stage.DiagramFlosss_mapString, diagramfloss.Name)
}

// commit diagramfloss to the back repo (if it is already staged)
func (diagramfloss *DiagramFloss) Commit(stage *Stage) *DiagramFloss {
	if _, ok := stage.DiagramFlosss[diagramfloss]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagramFloss(diagramfloss)
		}
	}
	return diagramfloss
}

func (diagramfloss *DiagramFloss) CommitVoid(stage *Stage) {
	diagramfloss.Commit(stage)
}

func (diagramfloss *DiagramFloss) StageVoid(stage *Stage) {
	diagramfloss.Stage(stage)
}

// Checkout diagramfloss to the back repo (if it is already staged)
func (diagramfloss *DiagramFloss) Checkout(stage *Stage) *DiagramFloss {
	if _, ok := stage.DiagramFlosss[diagramfloss]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagramFloss(diagramfloss)
		}
	}
	return diagramfloss
}

// for satisfaction of GongStruct interface
func (diagramfloss *DiagramFloss) GetName() (res string) {
	return diagramfloss.Name
}

// for satisfaction of GongStruct interface
func (diagramfloss *DiagramFloss) SetName(name string) {
	diagramfloss.Name = name
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

// Stage puts effortshape to the model stage
func (effortshape *EffortShape) Stage(stage *Stage) *EffortShape {
	if _, ok := stage.EffortShapes[effortshape]; !ok {
		stage.EffortShapes[effortshape] = struct{}{}
		stage.EffortShape_stagedOrder[effortshape] = stage.EffortShapeOrder
		stage.EffortShape_orderStaged[stage.EffortShapeOrder] = effortshape
		stage.EffortShapeOrder++
	}
	stage.EffortShapes_mapString[effortshape.Name] = effortshape

	return effortshape
}

// StagePreserveOrder puts effortshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EffortShapeOrder
// - update stage.EffortShapeOrder accordingly
func (effortshape *EffortShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.EffortShapes[effortshape]; !ok {
		stage.EffortShapes[effortshape] = struct{}{}

		if order > stage.EffortShapeOrder {
			stage.EffortShapeOrder = order
		}
		stage.EffortShape_stagedOrder[effortshape] = order
		stage.EffortShape_orderStaged[order] = effortshape
		stage.EffortShapeOrder++
	}
	stage.EffortShapes_mapString[effortshape.Name] = effortshape
}

// Unstage removes effortshape off the model stage
func (effortshape *EffortShape) Unstage(stage *Stage) *EffortShape {
	delete(stage.EffortShapes, effortshape)
	// issue1150
	// delete(stage.EffortShape_stagedOrder, effortshape)
	delete(stage.EffortShapes_mapString, effortshape.Name)

	return effortshape
}

// UnstageVoid removes effortshape off the model stage
func (effortshape *EffortShape) UnstageVoid(stage *Stage) {
	delete(stage.EffortShapes, effortshape)
	// issue1150
	// delete(stage.EffortShape_stagedOrder, effortshape)
	delete(stage.EffortShapes_mapString, effortshape.Name)
}

// commit effortshape to the back repo (if it is already staged)
func (effortshape *EffortShape) Commit(stage *Stage) *EffortShape {
	if _, ok := stage.EffortShapes[effortshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEffortShape(effortshape)
		}
	}
	return effortshape
}

func (effortshape *EffortShape) CommitVoid(stage *Stage) {
	effortshape.Commit(stage)
}

func (effortshape *EffortShape) StageVoid(stage *Stage) {
	effortshape.Stage(stage)
}

// Checkout effortshape to the back repo (if it is already staged)
func (effortshape *EffortShape) Checkout(stage *Stage) *EffortShape {
	if _, ok := stage.EffortShapes[effortshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEffortShape(effortshape)
		}
	}
	return effortshape
}

// for satisfaction of GongStruct interface
func (effortshape *EffortShape) GetName() (res string) {
	return effortshape.Name
}

// for satisfaction of GongStruct interface
func (effortshape *EffortShape) SetName(name string) {
	effortshape.Name = name
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

// Stage puts performanceshape to the model stage
func (performanceshape *PerformanceShape) Stage(stage *Stage) *PerformanceShape {
	if _, ok := stage.PerformanceShapes[performanceshape]; !ok {
		stage.PerformanceShapes[performanceshape] = struct{}{}
		stage.PerformanceShape_stagedOrder[performanceshape] = stage.PerformanceShapeOrder
		stage.PerformanceShape_orderStaged[stage.PerformanceShapeOrder] = performanceshape
		stage.PerformanceShapeOrder++
	}
	stage.PerformanceShapes_mapString[performanceshape.Name] = performanceshape

	return performanceshape
}

// StagePreserveOrder puts performanceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.PerformanceShapeOrder
// - update stage.PerformanceShapeOrder accordingly
func (performanceshape *PerformanceShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.PerformanceShapes[performanceshape]; !ok {
		stage.PerformanceShapes[performanceshape] = struct{}{}

		if order > stage.PerformanceShapeOrder {
			stage.PerformanceShapeOrder = order
		}
		stage.PerformanceShape_stagedOrder[performanceshape] = order
		stage.PerformanceShape_orderStaged[order] = performanceshape
		stage.PerformanceShapeOrder++
	}
	stage.PerformanceShapes_mapString[performanceshape.Name] = performanceshape
}

// Unstage removes performanceshape off the model stage
func (performanceshape *PerformanceShape) Unstage(stage *Stage) *PerformanceShape {
	delete(stage.PerformanceShapes, performanceshape)
	// issue1150
	// delete(stage.PerformanceShape_stagedOrder, performanceshape)
	delete(stage.PerformanceShapes_mapString, performanceshape.Name)

	return performanceshape
}

// UnstageVoid removes performanceshape off the model stage
func (performanceshape *PerformanceShape) UnstageVoid(stage *Stage) {
	delete(stage.PerformanceShapes, performanceshape)
	// issue1150
	// delete(stage.PerformanceShape_stagedOrder, performanceshape)
	delete(stage.PerformanceShapes_mapString, performanceshape.Name)
}

// commit performanceshape to the back repo (if it is already staged)
func (performanceshape *PerformanceShape) Commit(stage *Stage) *PerformanceShape {
	if _, ok := stage.PerformanceShapes[performanceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPerformanceShape(performanceshape)
		}
	}
	return performanceshape
}

func (performanceshape *PerformanceShape) CommitVoid(stage *Stage) {
	performanceshape.Commit(stage)
}

func (performanceshape *PerformanceShape) StageVoid(stage *Stage) {
	performanceshape.Stage(stage)
}

// Checkout performanceshape to the back repo (if it is already staged)
func (performanceshape *PerformanceShape) Checkout(stage *Stage) *PerformanceShape {
	if _, ok := stage.PerformanceShapes[performanceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPerformanceShape(performanceshape)
		}
	}
	return performanceshape
}

// for satisfaction of GongStruct interface
func (performanceshape *PerformanceShape) GetName() (res string) {
	return performanceshape.Name
}

// for satisfaction of GongStruct interface
func (performanceshape *PerformanceShape) SetName(name string) {
	performanceshape.Name = name
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

// Stage puts systemshape to the model stage
func (systemshape *SystemShape) Stage(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; !ok {
		stage.SystemShapes[systemshape] = struct{}{}
		stage.SystemShape_stagedOrder[systemshape] = stage.SystemShapeOrder
		stage.SystemShape_orderStaged[stage.SystemShapeOrder] = systemshape
		stage.SystemShapeOrder++
	}
	stage.SystemShapes_mapString[systemshape.Name] = systemshape

	return systemshape
}

// StagePreserveOrder puts systemshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SystemShapeOrder
// - update stage.SystemShapeOrder accordingly
func (systemshape *SystemShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SystemShapes[systemshape]; !ok {
		stage.SystemShapes[systemshape] = struct{}{}

		if order > stage.SystemShapeOrder {
			stage.SystemShapeOrder = order
		}
		stage.SystemShape_stagedOrder[systemshape] = order
		stage.SystemShape_orderStaged[order] = systemshape
		stage.SystemShapeOrder++
	}
	stage.SystemShapes_mapString[systemshape.Name] = systemshape
}

// Unstage removes systemshape off the model stage
func (systemshape *SystemShape) Unstage(stage *Stage) *SystemShape {
	delete(stage.SystemShapes, systemshape)
	// issue1150
	// delete(stage.SystemShape_stagedOrder, systemshape)
	delete(stage.SystemShapes_mapString, systemshape.Name)

	return systemshape
}

// UnstageVoid removes systemshape off the model stage
func (systemshape *SystemShape) UnstageVoid(stage *Stage) {
	delete(stage.SystemShapes, systemshape)
	// issue1150
	// delete(stage.SystemShape_stagedOrder, systemshape)
	delete(stage.SystemShapes_mapString, systemshape.Name)
}

// commit systemshape to the back repo (if it is already staged)
func (systemshape *SystemShape) Commit(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSystemShape(systemshape)
		}
	}
	return systemshape
}

func (systemshape *SystemShape) CommitVoid(stage *Stage) {
	systemshape.Commit(stage)
}

func (systemshape *SystemShape) StageVoid(stage *Stage) {
	systemshape.Stage(stage)
}

// Checkout systemshape to the back repo (if it is already staged)
func (systemshape *SystemShape) Checkout(stage *Stage) *SystemShape {
	if _, ok := stage.SystemShapes[systemshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSystemShape(systemshape)
		}
	}
	return systemshape
}

// for satisfaction of GongStruct interface
func (systemshape *SystemShape) GetName() (res string) {
	return systemshape.Name
}

// for satisfaction of GongStruct interface
func (systemshape *SystemShape) SetName(name string) {
	systemshape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCompareAnalysis(CompareAnalysis *CompareAnalysis)
	CreateORMComplexity(Complexity *Complexity)
	CreateORMComplexityShape(ComplexityShape *ComplexityShape)
	CreateORMDiagramFloss(DiagramFloss *DiagramFloss)
	CreateORMDiagramFlossEquation(DiagramFlossEquation *DiagramFlossEquation)
	CreateORMEffort(Effort *Effort)
	CreateORMEffortShape(EffortShape *EffortShape)
	CreateORMLibrary(Library *Library)
	CreateORMNote(Note *Note)
	CreateORMNoteComplexityShape(NoteComplexityShape *NoteComplexityShape)
	CreateORMNoteEffortShape(NoteEffortShape *NoteEffortShape)
	CreateORMNotePerformanceShape(NotePerformanceShape *NotePerformanceShape)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMPerformance(Performance *Performance)
	CreateORMPerformanceShape(PerformanceShape *PerformanceShape)
	CreateORMSystem(System *System)
	CreateORMSystemShape(SystemShape *SystemShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCompareAnalysis(CompareAnalysis *CompareAnalysis)
	DeleteORMComplexity(Complexity *Complexity)
	DeleteORMComplexityShape(ComplexityShape *ComplexityShape)
	DeleteORMDiagramFloss(DiagramFloss *DiagramFloss)
	DeleteORMDiagramFlossEquation(DiagramFlossEquation *DiagramFlossEquation)
	DeleteORMEffort(Effort *Effort)
	DeleteORMEffortShape(EffortShape *EffortShape)
	DeleteORMLibrary(Library *Library)
	DeleteORMNote(Note *Note)
	DeleteORMNoteComplexityShape(NoteComplexityShape *NoteComplexityShape)
	DeleteORMNoteEffortShape(NoteEffortShape *NoteEffortShape)
	DeleteORMNotePerformanceShape(NotePerformanceShape *NotePerformanceShape)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMPerformance(Performance *Performance)
	DeleteORMPerformanceShape(PerformanceShape *PerformanceShape)
	DeleteORMSystem(System *System)
	DeleteORMSystemShape(SystemShape *SystemShape)
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

	stage.ComplexityShapes = make(map[*ComplexityShape]struct{})
	stage.ComplexityShapes_mapString = make(map[string]*ComplexityShape)
	stage.ComplexityShape_stagedOrder = make(map[*ComplexityShape]uint)
	stage.ComplexityShapeOrder = 0

	stage.DiagramFlosss = make(map[*DiagramFloss]struct{})
	stage.DiagramFlosss_mapString = make(map[string]*DiagramFloss)
	stage.DiagramFloss_stagedOrder = make(map[*DiagramFloss]uint)
	stage.DiagramFlossOrder = 0

	stage.DiagramFlossEquations = make(map[*DiagramFlossEquation]struct{})
	stage.DiagramFlossEquations_mapString = make(map[string]*DiagramFlossEquation)
	stage.DiagramFlossEquation_stagedOrder = make(map[*DiagramFlossEquation]uint)
	stage.DiagramFlossEquationOrder = 0

	stage.Efforts = make(map[*Effort]struct{})
	stage.Efforts_mapString = make(map[string]*Effort)
	stage.Effort_stagedOrder = make(map[*Effort]uint)
	stage.EffortOrder = 0

	stage.EffortShapes = make(map[*EffortShape]struct{})
	stage.EffortShapes_mapString = make(map[string]*EffortShape)
	stage.EffortShape_stagedOrder = make(map[*EffortShape]uint)
	stage.EffortShapeOrder = 0

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

	stage.PerformanceShapes = make(map[*PerformanceShape]struct{})
	stage.PerformanceShapes_mapString = make(map[string]*PerformanceShape)
	stage.PerformanceShape_stagedOrder = make(map[*PerformanceShape]uint)
	stage.PerformanceShapeOrder = 0

	stage.Systems = make(map[*System]struct{})
	stage.Systems_mapString = make(map[string]*System)
	stage.System_stagedOrder = make(map[*System]uint)
	stage.SystemOrder = 0

	stage.SystemShapes = make(map[*SystemShape]struct{})
	stage.SystemShapes_mapString = make(map[string]*SystemShape)
	stage.SystemShape_stagedOrder = make(map[*SystemShape]uint)
	stage.SystemShapeOrder = 0

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

	stage.ComplexityShapes = nil
	stage.ComplexityShapes_mapString = nil

	stage.DiagramFlosss = nil
	stage.DiagramFlosss_mapString = nil

	stage.DiagramFlossEquations = nil
	stage.DiagramFlossEquations_mapString = nil

	stage.Efforts = nil
	stage.Efforts_mapString = nil

	stage.EffortShapes = nil
	stage.EffortShapes_mapString = nil

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

	stage.PerformanceShapes = nil
	stage.PerformanceShapes_mapString = nil

	stage.Systems = nil
	stage.Systems_mapString = nil

	stage.SystemShapes = nil
	stage.SystemShapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for compareanalysis := range stage.CompareAnalysiss {
		compareanalysis.Unstage(stage)
	}

	for complexity := range stage.Complexitys {
		complexity.Unstage(stage)
	}

	for complexityshape := range stage.ComplexityShapes {
		complexityshape.Unstage(stage)
	}

	for diagramfloss := range stage.DiagramFlosss {
		diagramfloss.Unstage(stage)
	}

	for diagramflossequation := range stage.DiagramFlossEquations {
		diagramflossequation.Unstage(stage)
	}

	for effort := range stage.Efforts {
		effort.Unstage(stage)
	}

	for effortshape := range stage.EffortShapes {
		effortshape.Unstage(stage)
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

	for performanceshape := range stage.PerformanceShapes {
		performanceshape.Unstage(stage)
	}

	for system := range stage.Systems {
		system.Unstage(stage)
	}

	for systemshape := range stage.SystemShapes {
		systemshape.Unstage(stage)
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
	case map[*ComplexityShape]any:
		return any(&stage.ComplexityShapes).(*Type)
	case map[*DiagramFloss]any:
		return any(&stage.DiagramFlosss).(*Type)
	case map[*DiagramFlossEquation]any:
		return any(&stage.DiagramFlossEquations).(*Type)
	case map[*Effort]any:
		return any(&stage.Efforts).(*Type)
	case map[*EffortShape]any:
		return any(&stage.EffortShapes).(*Type)
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
	case map[*PerformanceShape]any:
		return any(&stage.PerformanceShapes).(*Type)
	case map[*System]any:
		return any(&stage.Systems).(*Type)
	case map[*SystemShape]any:
		return any(&stage.SystemShapes).(*Type)
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
	case *ComplexityShape:
		return any(stage.ComplexityShapes_mapString).(map[string]Type)
	case *DiagramFloss:
		return any(stage.DiagramFlosss_mapString).(map[string]Type)
	case *DiagramFlossEquation:
		return any(stage.DiagramFlossEquations_mapString).(map[string]Type)
	case *Effort:
		return any(stage.Efforts_mapString).(map[string]Type)
	case *EffortShape:
		return any(stage.EffortShapes_mapString).(map[string]Type)
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
	case *PerformanceShape:
		return any(stage.PerformanceShapes_mapString).(map[string]Type)
	case *System:
		return any(stage.Systems_mapString).(map[string]Type)
	case *SystemShape:
		return any(stage.SystemShapes_mapString).(map[string]Type)
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
	case ComplexityShape:
		return any(&stage.ComplexityShapes).(*map[*Type]struct{})
	case DiagramFloss:
		return any(&stage.DiagramFlosss).(*map[*Type]struct{})
	case DiagramFlossEquation:
		return any(&stage.DiagramFlossEquations).(*map[*Type]struct{})
	case Effort:
		return any(&stage.Efforts).(*map[*Type]struct{})
	case EffortShape:
		return any(&stage.EffortShapes).(*map[*Type]struct{})
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
	case PerformanceShape:
		return any(&stage.PerformanceShapes).(*map[*Type]struct{})
	case System:
		return any(&stage.Systems).(*map[*Type]struct{})
	case SystemShape:
		return any(&stage.SystemShapes).(*map[*Type]struct{})
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
	case *ComplexityShape:
		return any(&stage.ComplexityShapes).(*map[Type]struct{})
	case *DiagramFloss:
		return any(&stage.DiagramFlosss).(*map[Type]struct{})
	case *DiagramFlossEquation:
		return any(&stage.DiagramFlossEquations).(*map[Type]struct{})
	case *Effort:
		return any(&stage.Efforts).(*map[Type]struct{})
	case *EffortShape:
		return any(&stage.EffortShapes).(*map[Type]struct{})
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
	case *PerformanceShape:
		return any(&stage.PerformanceShapes).(*map[Type]struct{})
	case *System:
		return any(&stage.Systems).(*map[Type]struct{})
	case *SystemShape:
		return any(&stage.SystemShapes).(*map[Type]struct{})
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
	case ComplexityShape:
		return any(&stage.ComplexityShapes_mapString).(*map[string]*Type)
	case DiagramFloss:
		return any(&stage.DiagramFlosss_mapString).(*map[string]*Type)
	case DiagramFlossEquation:
		return any(&stage.DiagramFlossEquations_mapString).(*map[string]*Type)
	case Effort:
		return any(&stage.Efforts_mapString).(*map[string]*Type)
	case EffortShape:
		return any(&stage.EffortShapes_mapString).(*map[string]*Type)
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
	case PerformanceShape:
		return any(&stage.PerformanceShapes_mapString).(*map[string]*Type)
	case System:
		return any(&stage.Systems_mapString).(*map[string]*Type)
	case SystemShape:
		return any(&stage.SystemShapes_mapString).(*map[string]*Type)
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
	case ComplexityShape:
		return any(&ComplexityShape{
			// Initialisation of associations
			// field is initialized with an instance of Complexity with the name of the field
			Complexity: &Complexity{Name: "Complexity"},
		}).(*Type)
	case DiagramFloss:
		return any(&DiagramFloss{
			// Initialisation of associations
			// field is initialized with an instance of SystemShape with the name of the field
			System_Shapes: []*SystemShape{{Name: "System_Shapes"}},
			// field is initialized with an instance of System with the name of the field
			SystemsWhoseNodeIsExpanded: []*System{{Name: "SystemsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ComplexityShape with the name of the field
			Complexity_Shapes: []*ComplexityShape{{Name: "Complexity_Shapes"}},
			// field is initialized with an instance of Complexity with the name of the field
			ComplexitysWhoseNodeIsExpanded: []*Complexity{{Name: "ComplexitysWhoseNodeIsExpanded"}},
			// field is initialized with an instance of PerformanceShape with the name of the field
			Performance_Shapes: []*PerformanceShape{{Name: "Performance_Shapes"}},
			// field is initialized with an instance of Performance with the name of the field
			PerformancesWhoseNodeIsExpanded: []*Performance{{Name: "PerformancesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of EffortShape with the name of the field
			Effort_Shapes: []*EffortShape{{Name: "Effort_Shapes"}},
			// field is initialized with an instance of Effort with the name of the field
			EffortsWhoseNodeIsExpanded: []*Effort{{Name: "EffortsWhoseNodeIsExpanded"}},
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
		}).(*Type)
	case Effort:
		return any(&Effort{
			// Initialisation of associations
		}).(*Type)
	case EffortShape:
		return any(&EffortShape{
			// Initialisation of associations
			// field is initialized with an instance of Effort with the name of the field
			Effort: &Effort{Name: "Effort"},
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
	case PerformanceShape:
		return any(&PerformanceShape{
			// Initialisation of associations
			// field is initialized with an instance of Performance with the name of the field
			Performance: &Performance{Name: "Performance"},
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
			// field is initialized with an instance of DiagramFloss with the name of the field
			DiagramFlosses: []*DiagramFloss{{Name: "DiagramFlosses"}},
			// field is initialized with an instance of DiagramFloss with the name of the field
			DiagramFlossWhoseNodeIsExpanded: []*DiagramFloss{{Name: "DiagramFlossWhoseNodeIsExpanded"}},
			// field is initialized with an instance of System with the name of the field
			SubSystemes: []*System{{Name: "SubSystemes"}},
			// field is initialized with an instance of Complexity with the name of the field
			ComplexitysWhoseNodeIsExpanded: []*Complexity{{Name: "ComplexitysWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Performance with the name of the field
			PerformancesWhoseNodeIsExpanded: []*Performance{{Name: "PerformancesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Effort with the name of the field
			EffortsWhoseNodeIsExpanded: []*Effort{{Name: "EffortsWhoseNodeIsExpanded"}},
		}).(*Type)
	case SystemShape:
		return any(&SystemShape{
			// Initialisation of associations
			// field is initialized with an instance of System with the name of the field
			System: &System{Name: "System"},
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
	// reverse maps of direct associations of ComplexityShape
	case ComplexityShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Complexity":
			res := make(map[*Complexity][]*ComplexityShape)
			for complexityshape := range stage.ComplexityShapes {
				if complexityshape.Complexity != nil {
					complexity_ := complexityshape.Complexity
					var complexityshapes []*ComplexityShape
					_, ok := res[complexity_]
					if ok {
						complexityshapes = res[complexity_]
					} else {
						complexityshapes = make([]*ComplexityShape, 0)
					}
					complexityshapes = append(complexityshapes, complexityshape)
					res[complexity_] = complexityshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DiagramFloss
	case DiagramFloss:
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
	// reverse maps of direct associations of EffortShape
	case EffortShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Effort":
			res := make(map[*Effort][]*EffortShape)
			for effortshape := range stage.EffortShapes {
				if effortshape.Effort != nil {
					effort_ := effortshape.Effort
					var effortshapes []*EffortShape
					_, ok := res[effort_]
					if ok {
						effortshapes = res[effort_]
					} else {
						effortshapes = make([]*EffortShape, 0)
					}
					effortshapes = append(effortshapes, effortshape)
					res[effort_] = effortshapes
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
	// reverse maps of direct associations of PerformanceShape
	case PerformanceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Performance":
			res := make(map[*Performance][]*PerformanceShape)
			for performanceshape := range stage.PerformanceShapes {
				if performanceshape.Performance != nil {
					performance_ := performanceshape.Performance
					var performanceshapes []*PerformanceShape
					_, ok := res[performance_]
					if ok {
						performanceshapes = res[performance_]
					} else {
						performanceshapes = make([]*PerformanceShape, 0)
					}
					performanceshapes = append(performanceshapes, performanceshape)
					res[performance_] = performanceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of System
	case System:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SystemShape
	case SystemShape:
		switch fieldname {
		// insertion point for per direct association field
		case "System":
			res := make(map[*System][]*SystemShape)
			for systemshape := range stage.SystemShapes {
				if systemshape.System != nil {
					system_ := systemshape.System
					var systemshapes []*SystemShape
					_, ok := res[system_]
					if ok {
						systemshapes = res[system_]
					} else {
						systemshapes = make([]*SystemShape, 0)
					}
					systemshapes = append(systemshapes, systemshape)
					res[system_] = systemshapes
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
	// reverse maps of direct associations of ComplexityShape
	case ComplexityShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DiagramFloss
	case DiagramFloss:
		switch fieldname {
		// insertion point for per direct association field
		case "System_Shapes":
			res := make(map[*SystemShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, systemshape_ := range diagramfloss.System_Shapes {
					res[systemshape_] = append(res[systemshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SystemsWhoseNodeIsExpanded":
			res := make(map[*System][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, system_ := range diagramfloss.SystemsWhoseNodeIsExpanded {
					res[system_] = append(res[system_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Complexity_Shapes":
			res := make(map[*ComplexityShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, complexityshape_ := range diagramfloss.Complexity_Shapes {
					res[complexityshape_] = append(res[complexityshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexitysWhoseNodeIsExpanded":
			res := make(map[*Complexity][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, complexity_ := range diagramfloss.ComplexitysWhoseNodeIsExpanded {
					res[complexity_] = append(res[complexity_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Performance_Shapes":
			res := make(map[*PerformanceShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, performanceshape_ := range diagramfloss.Performance_Shapes {
					res[performanceshape_] = append(res[performanceshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "PerformancesWhoseNodeIsExpanded":
			res := make(map[*Performance][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, performance_ := range diagramfloss.PerformancesWhoseNodeIsExpanded {
					res[performance_] = append(res[performance_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Effort_Shapes":
			res := make(map[*EffortShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, effortshape_ := range diagramfloss.Effort_Shapes {
					res[effortshape_] = append(res[effortshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "EffortsWhoseNodeIsExpanded":
			res := make(map[*Effort][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, effort_ := range diagramfloss.EffortsWhoseNodeIsExpanded {
					res[effort_] = append(res[effort_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Note_Shapes":
			res := make(map[*NoteShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, noteshape_ := range diagramfloss.Note_Shapes {
					res[noteshape_] = append(res[noteshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteComplexityShapes":
			res := make(map[*NoteComplexityShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, notecomplexityshape_ := range diagramfloss.NoteComplexityShapes {
					res[notecomplexityshape_] = append(res[notecomplexityshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotePerformanceShapes":
			res := make(map[*NotePerformanceShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, noteperformanceshape_ := range diagramfloss.NotePerformanceShapes {
					res[noteperformanceshape_] = append(res[noteperformanceshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteEffortShapes":
			res := make(map[*NoteEffortShape][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, noteeffortshape_ := range diagramfloss.NoteEffortShapes {
					res[noteeffortshape_] = append(res[noteeffortshape_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotesWhoseNodeIsExpanded":
			res := make(map[*Note][]*DiagramFloss)
			for diagramfloss := range stage.DiagramFlosss {
				for _, note_ := range diagramfloss.NotesWhoseNodeIsExpanded {
					res[note_] = append(res[note_], diagramfloss)
				}
			}
			return any(res).(map[*End][]*Start)
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
		}
	// reverse maps of direct associations of Effort
	case Effort:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EffortShape
	case EffortShape:
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
	// reverse maps of direct associations of PerformanceShape
	case PerformanceShape:
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
		case "DiagramFlosses":
			res := make(map[*DiagramFloss][]*System)
			for system := range stage.Systems {
				for _, diagramfloss_ := range system.DiagramFlosses {
					res[diagramfloss_] = append(res[diagramfloss_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramFlossWhoseNodeIsExpanded":
			res := make(map[*DiagramFloss][]*System)
			for system := range stage.Systems {
				for _, diagramfloss_ := range system.DiagramFlossWhoseNodeIsExpanded {
					res[diagramfloss_] = append(res[diagramfloss_], system)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubSystemes":
			res := make(map[*System][]*System)
			for system := range stage.Systems {
				for _, system_ := range system.SubSystemes {
					res[system_] = append(res[system_], system)
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
	// reverse maps of direct associations of SystemShape
	case SystemShape:
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
	case *CompareAnalysis:
		res = "CompareAnalysis"
	case *Complexity:
		res = "Complexity"
	case *ComplexityShape:
		res = "ComplexityShape"
	case *DiagramFloss:
		res = "DiagramFloss"
	case *DiagramFlossEquation:
		res = "DiagramFlossEquation"
	case *Effort:
		res = "Effort"
	case *EffortShape:
		res = "EffortShape"
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
	case *PerformanceShape:
		res = "PerformanceShape"
	case *System:
		res = "System"
	case *SystemShape:
		res = "SystemShape"
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
		rf.GongstructName = "DiagramFloss"
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
	case *ComplexityShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "Complexity_Shapes"
		res = append(res, rf)
	case *DiagramFloss:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramFlosses"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "DiagramFlossWhoseNodeIsExpanded"
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
	case *Effort:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
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
	case *EffortShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "Effort_Shapes"
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
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "NotesWhoseNodeIsExpanded"
		res = append(res, rf)
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
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "NoteComplexityShapes"
		res = append(res, rf)
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "NoteComplexityShapes"
		res = append(res, rf)
	case *NoteEffortShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "NoteEffortShapes"
		res = append(res, rf)
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "NoteEffortShapes"
		res = append(res, rf)
	case *NotePerformanceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "NotePerformanceShapes"
		res = append(res, rf)
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "NotePerformanceShapes"
		res = append(res, rf)
	case *NoteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "Note_Shapes"
		res = append(res, rf)
		rf.GongstructName = "DiagramFlossEquation"
		rf.Fieldname = "Note_Shapes"
		res = append(res, rf)
	case *Performance:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
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
	case *PerformanceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "Performance_Shapes"
		res = append(res, rf)
	case *System:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "SystemsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootSystems"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SystemsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "System"
		rf.Fieldname = "SubSystemes"
		res = append(res, rf)
	case *SystemShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramFloss"
		rf.Fieldname = "System_Shapes"
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
			Name:               "Alpha",
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

func (complexityshape *ComplexityShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Complexity",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Complexity",
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

func (diagramfloss *DiagramFloss) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "System_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "SystemShape",
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
			Name:                 "Complexity_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ComplexityShape",
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
			Name:                 "Performance_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "PerformanceShape",
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
			Name:                 "Effort_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "EffortShape",
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
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Scale",
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

func (effortshape *EffortShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Effort",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Effort",
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
			Name:                 "Complexities",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Complexity",
		},
		{
			Name:               "IsPerformancesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Performances",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Performance",
		},
		{
			Name:               "IsEffortsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Efforts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Effort",
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

func (performanceshape *PerformanceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Performance",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Performance",
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
			Name:                 "DiagramFlosses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramFloss",
		},
		{
			Name:                 "DiagramFlossWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramFloss",
		},
		{
			Name:               "IsSubSystemNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubSystemes",
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
	}
	return
}

func (systemshape *SystemShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "System",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "System",
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
	case "Alpha":
		res.valueString = fmt.Sprintf("%f", compareanalysis.Alpha)
		res.valueFloat = compareanalysis.Alpha
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
	case "ComputedPrefix":
		res.valueString = complexity.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", complexity.IsExpanded)
		res.valueBool = complexity.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (complexityshape *ComplexityShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = complexityshape.Name
	case "Complexity":
		res.GongFieldValueType = GongFieldValueTypePointer
		if complexityshape.Complexity != nil {
			res.valueString = complexityshape.Complexity.Name
			res.ids = complexityshape.Complexity.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", complexityshape.IsExpanded)
		res.valueBool = complexityshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", complexityshape.X)
		res.valueFloat = complexityshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", complexityshape.Y)
		res.valueFloat = complexityshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", complexityshape.Width)
		res.valueFloat = complexityshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", complexityshape.Height)
		res.valueFloat = complexityshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", complexityshape.IsHidden)
		res.valueBool = complexityshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (diagramfloss *DiagramFloss) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramfloss.Name
	case "Description":
		res.valueString = diagramfloss.Description
	case "ComputedPrefix":
		res.valueString = diagramfloss.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsExpanded)
		res.valueBool = diagramfloss.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsChecked)
		res.valueBool = diagramfloss.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsEditable_)
		res.valueBool = diagramfloss.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsShowPrefix)
		res.valueBool = diagramfloss.IsShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagramfloss.DefaultBoxWidth)
		res.valueFloat = diagramfloss.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagramfloss.DefaultBoxHeigth)
		res.valueFloat = diagramfloss.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramfloss.Width)
		res.valueFloat = diagramfloss.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramfloss.Height)
		res.valueFloat = diagramfloss.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "System_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.System_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSystemsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsSystemsNodeExpanded)
		res.valueBool = diagramfloss.IsSystemsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SystemsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.SystemsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Complexity_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.Complexity_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsComplexitysNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsComplexitysNodeExpanded)
		res.valueBool = diagramfloss.IsComplexitysNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComplexitysWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.ComplexitysWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Performance_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.Performance_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsPerformancesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsPerformancesNodeExpanded)
		res.valueBool = diagramfloss.IsPerformancesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "PerformancesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.PerformancesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Effort_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.Effort_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsEffortsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsEffortsNodeExpanded)
		res.valueBool = diagramfloss.IsEffortsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "EffortsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.EffortsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Note_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.Note_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NoteComplexityShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.NoteComplexityShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NotePerformanceShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.NotePerformanceShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NoteEffortShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.NoteEffortShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsNotesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramfloss.IsNotesNodeExpanded)
		res.valueBool = diagramfloss.IsNotesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NotesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramfloss.NotesWhoseNodeIsExpanded {
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

func (diagramflossequation *DiagramFlossEquation) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramflossequation.Name
	case "Description":
		res.valueString = diagramflossequation.Description
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
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramflossequation.Width)
		res.valueFloat = diagramflossequation.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramflossequation.Height)
		res.valueFloat = diagramflossequation.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Scale":
		res.valueString = fmt.Sprintf("%f", diagramflossequation.Scale)
		res.valueFloat = diagramflossequation.Scale
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
	case "ComputedPrefix":
		res.valueString = effort.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", effort.IsExpanded)
		res.valueBool = effort.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (effortshape *EffortShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = effortshape.Name
	case "Effort":
		res.GongFieldValueType = GongFieldValueTypePointer
		if effortshape.Effort != nil {
			res.valueString = effortshape.Effort.Name
			res.ids = effortshape.Effort.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", effortshape.IsExpanded)
		res.valueBool = effortshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", effortshape.X)
		res.valueFloat = effortshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", effortshape.Y)
		res.valueFloat = effortshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", effortshape.Width)
		res.valueFloat = effortshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", effortshape.Height)
		res.valueFloat = effortshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", effortshape.IsHidden)
		res.valueBool = effortshape.IsHidden
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
	case "IsPerformancesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsPerformancesNodeExpanded)
		res.valueBool = note.IsPerformancesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "IsEffortsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsEffortsNodeExpanded)
		res.valueBool = note.IsEffortsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "ComputedPrefix":
		res.valueString = performance.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", performance.IsExpanded)
		res.valueBool = performance.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (performanceshape *PerformanceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = performanceshape.Name
	case "Performance":
		res.GongFieldValueType = GongFieldValueTypePointer
		if performanceshape.Performance != nil {
			res.valueString = performanceshape.Performance.Name
			res.ids = performanceshape.Performance.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", performanceshape.IsExpanded)
		res.valueBool = performanceshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", performanceshape.X)
		res.valueFloat = performanceshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", performanceshape.Y)
		res.valueFloat = performanceshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", performanceshape.Width)
		res.valueFloat = performanceshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", performanceshape.Height)
		res.valueFloat = performanceshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", performanceshape.IsHidden)
		res.valueBool = performanceshape.IsHidden
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
	case "DiagramFlosses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramFlosses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DiagramFlossWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.DiagramFlossWhoseNodeIsExpanded {
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
	case "SubSystemes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range system.SubSystemes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
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

func (systemshape *SystemShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = systemshape.Name
	case "System":
		res.GongFieldValueType = GongFieldValueTypePointer
		if systemshape.System != nil {
			res.valueString = systemshape.System.Name
			res.ids = systemshape.System.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", systemshape.IsExpanded)
		res.valueBool = systemshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", systemshape.X)
		res.valueFloat = systemshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", systemshape.Y)
		res.valueFloat = systemshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", systemshape.Width)
		res.valueFloat = systemshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", systemshape.Height)
		res.valueFloat = systemshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", systemshape.IsHidden)
		res.valueBool = systemshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "Alpha":
		compareanalysis.Alpha = value.GetValueFloat()
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
	case "ComputedPrefix":
		complexity.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		complexity.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (complexityshape *ComplexityShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		complexityshape.Name = value.GetValueString()
	case "Complexity":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			complexityshape.Complexity = nil
			for __instance__ := range stage.Complexitys {
				if stage.Complexity_stagedOrder[__instance__] == uint(id) {
					complexityshape.Complexity = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		complexityshape.IsExpanded = value.GetValueBool()
	case "X":
		complexityshape.X = value.GetValueFloat()
	case "Y":
		complexityshape.Y = value.GetValueFloat()
	case "Width":
		complexityshape.Width = value.GetValueFloat()
	case "Height":
		complexityshape.Height = value.GetValueFloat()
	case "IsHidden":
		complexityshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (diagramfloss *DiagramFloss) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagramfloss.Name = value.GetValueString()
	case "Description":
		diagramfloss.Description = value.GetValueString()
	case "ComputedPrefix":
		diagramfloss.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagramfloss.IsExpanded = value.GetValueBool()
	case "IsChecked":
		diagramfloss.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagramfloss.IsEditable_ = value.GetValueBool()
	case "IsShowPrefix":
		diagramfloss.IsShowPrefix = value.GetValueBool()
	case "DefaultBoxWidth":
		diagramfloss.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagramfloss.DefaultBoxHeigth = value.GetValueFloat()
	case "Width":
		diagramfloss.Width = value.GetValueFloat()
	case "Height":
		diagramfloss.Height = value.GetValueFloat()
	case "System_Shapes":
		diagramfloss.System_Shapes = make([]*SystemShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.SystemShapes {
					if stage.SystemShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.System_Shapes = append(diagramfloss.System_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsSystemsNodeExpanded":
		diagramfloss.IsSystemsNodeExpanded = value.GetValueBool()
	case "SystemsWhoseNodeIsExpanded":
		diagramfloss.SystemsWhoseNodeIsExpanded = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						diagramfloss.SystemsWhoseNodeIsExpanded = append(diagramfloss.SystemsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Complexity_Shapes":
		diagramfloss.Complexity_Shapes = make([]*ComplexityShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ComplexityShapes {
					if stage.ComplexityShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.Complexity_Shapes = append(diagramfloss.Complexity_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsComplexitysNodeExpanded":
		diagramfloss.IsComplexitysNodeExpanded = value.GetValueBool()
	case "ComplexitysWhoseNodeIsExpanded":
		diagramfloss.ComplexitysWhoseNodeIsExpanded = make([]*Complexity, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Complexitys {
					if stage.Complexity_stagedOrder[__instance__] == uint(id) {
						diagramfloss.ComplexitysWhoseNodeIsExpanded = append(diagramfloss.ComplexitysWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Performance_Shapes":
		diagramfloss.Performance_Shapes = make([]*PerformanceShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.PerformanceShapes {
					if stage.PerformanceShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.Performance_Shapes = append(diagramfloss.Performance_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsPerformancesNodeExpanded":
		diagramfloss.IsPerformancesNodeExpanded = value.GetValueBool()
	case "PerformancesWhoseNodeIsExpanded":
		diagramfloss.PerformancesWhoseNodeIsExpanded = make([]*Performance, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Performances {
					if stage.Performance_stagedOrder[__instance__] == uint(id) {
						diagramfloss.PerformancesWhoseNodeIsExpanded = append(diagramfloss.PerformancesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Effort_Shapes":
		diagramfloss.Effort_Shapes = make([]*EffortShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.EffortShapes {
					if stage.EffortShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.Effort_Shapes = append(diagramfloss.Effort_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsEffortsNodeExpanded":
		diagramfloss.IsEffortsNodeExpanded = value.GetValueBool()
	case "EffortsWhoseNodeIsExpanded":
		diagramfloss.EffortsWhoseNodeIsExpanded = make([]*Effort, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Efforts {
					if stage.Effort_stagedOrder[__instance__] == uint(id) {
						diagramfloss.EffortsWhoseNodeIsExpanded = append(diagramfloss.EffortsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Note_Shapes":
		diagramfloss.Note_Shapes = make([]*NoteShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteShapes {
					if stage.NoteShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.Note_Shapes = append(diagramfloss.Note_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "NoteComplexityShapes":
		diagramfloss.NoteComplexityShapes = make([]*NoteComplexityShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteComplexityShapes {
					if stage.NoteComplexityShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.NoteComplexityShapes = append(diagramfloss.NoteComplexityShapes, __instance__)
						break
					}
				}
			}
		}
	case "NotePerformanceShapes":
		diagramfloss.NotePerformanceShapes = make([]*NotePerformanceShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NotePerformanceShapes {
					if stage.NotePerformanceShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.NotePerformanceShapes = append(diagramfloss.NotePerformanceShapes, __instance__)
						break
					}
				}
			}
		}
	case "NoteEffortShapes":
		diagramfloss.NoteEffortShapes = make([]*NoteEffortShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteEffortShapes {
					if stage.NoteEffortShape_stagedOrder[__instance__] == uint(id) {
						diagramfloss.NoteEffortShapes = append(diagramfloss.NoteEffortShapes, __instance__)
						break
					}
				}
			}
		}
	case "IsNotesNodeExpanded":
		diagramfloss.IsNotesNodeExpanded = value.GetValueBool()
	case "NotesWhoseNodeIsExpanded":
		diagramfloss.NotesWhoseNodeIsExpanded = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						diagramfloss.NotesWhoseNodeIsExpanded = append(diagramfloss.NotesWhoseNodeIsExpanded, __instance__)
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

func (diagramflossequation *DiagramFlossEquation) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagramflossequation.Name = value.GetValueString()
	case "Description":
		diagramflossequation.Description = value.GetValueString()
	case "ComputedPrefix":
		diagramflossequation.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		diagramflossequation.IsExpanded = value.GetValueBool()
	case "IsChecked":
		diagramflossequation.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagramflossequation.IsEditable_ = value.GetValueBool()
	case "Width":
		diagramflossequation.Width = value.GetValueFloat()
	case "Height":
		diagramflossequation.Height = value.GetValueFloat()
	case "Scale":
		diagramflossequation.Scale = value.GetValueFloat()
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
	case "ComputedPrefix":
		effort.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		effort.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (effortshape *EffortShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		effortshape.Name = value.GetValueString()
	case "Effort":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			effortshape.Effort = nil
			for __instance__ := range stage.Efforts {
				if stage.Effort_stagedOrder[__instance__] == uint(id) {
					effortshape.Effort = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		effortshape.IsExpanded = value.GetValueBool()
	case "X":
		effortshape.X = value.GetValueFloat()
	case "Y":
		effortshape.Y = value.GetValueFloat()
	case "Width":
		effortshape.Width = value.GetValueFloat()
	case "Height":
		effortshape.Height = value.GetValueFloat()
	case "IsHidden":
		effortshape.IsHidden = value.GetValueBool()
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
	case "ComputedPrefix":
		note.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		note.IsExpanded = value.GetValueBool()
	case "IsComplexitysNodeExpanded":
		note.IsComplexitysNodeExpanded = value.GetValueBool()
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
	case "IsPerformancesNodeExpanded":
		note.IsPerformancesNodeExpanded = value.GetValueBool()
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
	case "IsEffortsNodeExpanded":
		note.IsEffortsNodeExpanded = value.GetValueBool()
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
	case "ComputedPrefix":
		performance.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		performance.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (performanceshape *PerformanceShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		performanceshape.Name = value.GetValueString()
	case "Performance":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			performanceshape.Performance = nil
			for __instance__ := range stage.Performances {
				if stage.Performance_stagedOrder[__instance__] == uint(id) {
					performanceshape.Performance = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		performanceshape.IsExpanded = value.GetValueBool()
	case "X":
		performanceshape.X = value.GetValueFloat()
	case "Y":
		performanceshape.Y = value.GetValueFloat()
	case "Width":
		performanceshape.Width = value.GetValueFloat()
	case "Height":
		performanceshape.Height = value.GetValueFloat()
	case "IsHidden":
		performanceshape.IsHidden = value.GetValueBool()
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
	case "ComputedPrefix":
		system.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		system.IsExpanded = value.GetValueBool()
	case "SVG_Path":
		system.SVG_Path = value.GetValueString()
	case "InverseAppliedScaling":
		system.InverseAppliedScaling = value.GetValueFloat()
	case "DiagramFlosses":
		system.DiagramFlosses = make([]*DiagramFloss, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramFlosss {
					if stage.DiagramFloss_stagedOrder[__instance__] == uint(id) {
						system.DiagramFlosses = append(system.DiagramFlosses, __instance__)
						break
					}
				}
			}
		}
	case "DiagramFlossWhoseNodeIsExpanded":
		system.DiagramFlossWhoseNodeIsExpanded = make([]*DiagramFloss, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramFlosss {
					if stage.DiagramFloss_stagedOrder[__instance__] == uint(id) {
						system.DiagramFlossWhoseNodeIsExpanded = append(system.DiagramFlossWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsSubSystemNodeExpanded":
		system.IsSubSystemNodeExpanded = value.GetValueBool()
	case "SubSystemes":
		system.SubSystemes = make([]*System, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Systems {
					if stage.System_stagedOrder[__instance__] == uint(id) {
						system.SubSystemes = append(system.SubSystemes, __instance__)
						break
					}
				}
			}
		}
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

func (systemshape *SystemShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		systemshape.Name = value.GetValueString()
	case "System":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			systemshape.System = nil
			for __instance__ := range stage.Systems {
				if stage.System_stagedOrder[__instance__] == uint(id) {
					systemshape.System = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		systemshape.IsExpanded = value.GetValueBool()
	case "X":
		systemshape.X = value.GetValueFloat()
	case "Y":
		systemshape.Y = value.GetValueFloat()
	case "Width":
		systemshape.Width = value.GetValueFloat()
	case "Height":
		systemshape.Height = value.GetValueFloat()
	case "IsHidden":
		systemshape.IsHidden = value.GetValueBool()
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

func (complexityshape *ComplexityShape) GongGetGongstructName() string {
	return "ComplexityShape"
}

func (diagramfloss *DiagramFloss) GongGetGongstructName() string {
	return "DiagramFloss"
}

func (diagramflossequation *DiagramFlossEquation) GongGetGongstructName() string {
	return "DiagramFlossEquation"
}

func (effort *Effort) GongGetGongstructName() string {
	return "Effort"
}

func (effortshape *EffortShape) GongGetGongstructName() string {
	return "EffortShape"
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

func (performanceshape *PerformanceShape) GongGetGongstructName() string {
	return "PerformanceShape"
}

func (system *System) GongGetGongstructName() string {
	return "System"
}

func (systemshape *SystemShape) GongGetGongstructName() string {
	return "SystemShape"
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

	stage.ComplexityShapes_mapString = make(map[string]*ComplexityShape)
	for complexityshape := range stage.ComplexityShapes {
		stage.ComplexityShapes_mapString[complexityshape.Name] = complexityshape
	}

	stage.DiagramFlosss_mapString = make(map[string]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		stage.DiagramFlosss_mapString[diagramfloss.Name] = diagramfloss
	}

	stage.DiagramFlossEquations_mapString = make(map[string]*DiagramFlossEquation)
	for diagramflossequation := range stage.DiagramFlossEquations {
		stage.DiagramFlossEquations_mapString[diagramflossequation.Name] = diagramflossequation
	}

	stage.Efforts_mapString = make(map[string]*Effort)
	for effort := range stage.Efforts {
		stage.Efforts_mapString[effort.Name] = effort
	}

	stage.EffortShapes_mapString = make(map[string]*EffortShape)
	for effortshape := range stage.EffortShapes {
		stage.EffortShapes_mapString[effortshape.Name] = effortshape
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

	stage.PerformanceShapes_mapString = make(map[string]*PerformanceShape)
	for performanceshape := range stage.PerformanceShapes {
		stage.PerformanceShapes_mapString[performanceshape.Name] = performanceshape
	}

	stage.Systems_mapString = make(map[string]*System)
	for system := range stage.Systems {
		stage.Systems_mapString[system.Name] = system
	}

	stage.SystemShapes_mapString = make(map[string]*SystemShape)
	for systemshape := range stage.SystemShapes {
		stage.SystemShapes_mapString[systemshape.Name] = systemshape
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
