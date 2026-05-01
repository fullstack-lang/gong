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

	process_go "github.com/fullstack-lang/gong/dsm/process/go"
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
	ControlFlows                map[*ControlFlow]struct{}
	ControlFlows_instance       map[*ControlFlow]*ControlFlow
	ControlFlows_mapString      map[string]*ControlFlow
	ControlFlowOrder            uint
	ControlFlow_stagedOrder     map[*ControlFlow]uint
	ControlFlow_orderStaged     map[uint]*ControlFlow
	ControlFlows_reference      map[*ControlFlow]*ControlFlow
	ControlFlows_referenceOrder map[*ControlFlow]uint

	// insertion point for slice of pointers maps
	OnAfterControlFlowCreateCallback OnAfterCreateInterface[ControlFlow]
	OnAfterControlFlowUpdateCallback OnAfterUpdateInterface[ControlFlow]
	OnAfterControlFlowDeleteCallback OnAfterDeleteInterface[ControlFlow]
	OnAfterControlFlowReadCallback   OnAfterReadInterface[ControlFlow]

	ControlFlowShapes                map[*ControlFlowShape]struct{}
	ControlFlowShapes_instance       map[*ControlFlowShape]*ControlFlowShape
	ControlFlowShapes_mapString      map[string]*ControlFlowShape
	ControlFlowShapeOrder            uint
	ControlFlowShape_stagedOrder     map[*ControlFlowShape]uint
	ControlFlowShape_orderStaged     map[uint]*ControlFlowShape
	ControlFlowShapes_reference      map[*ControlFlowShape]*ControlFlowShape
	ControlFlowShapes_referenceOrder map[*ControlFlowShape]uint

	// insertion point for slice of pointers maps
	OnAfterControlFlowShapeCreateCallback OnAfterCreateInterface[ControlFlowShape]
	OnAfterControlFlowShapeUpdateCallback OnAfterUpdateInterface[ControlFlowShape]
	OnAfterControlFlowShapeDeleteCallback OnAfterDeleteInterface[ControlFlowShape]
	OnAfterControlFlowShapeReadCallback   OnAfterReadInterface[ControlFlowShape]

	DataFlows                map[*DataFlow]struct{}
	DataFlows_instance       map[*DataFlow]*DataFlow
	DataFlows_mapString      map[string]*DataFlow
	DataFlowOrder            uint
	DataFlow_stagedOrder     map[*DataFlow]uint
	DataFlow_orderStaged     map[uint]*DataFlow
	DataFlows_reference      map[*DataFlow]*DataFlow
	DataFlows_referenceOrder map[*DataFlow]uint

	// insertion point for slice of pointers maps
	OnAfterDataFlowCreateCallback OnAfterCreateInterface[DataFlow]
	OnAfterDataFlowUpdateCallback OnAfterUpdateInterface[DataFlow]
	OnAfterDataFlowDeleteCallback OnAfterDeleteInterface[DataFlow]
	OnAfterDataFlowReadCallback   OnAfterReadInterface[DataFlow]

	DataFlowShapes                map[*DataFlowShape]struct{}
	DataFlowShapes_instance       map[*DataFlowShape]*DataFlowShape
	DataFlowShapes_mapString      map[string]*DataFlowShape
	DataFlowShapeOrder            uint
	DataFlowShape_stagedOrder     map[*DataFlowShape]uint
	DataFlowShape_orderStaged     map[uint]*DataFlowShape
	DataFlowShapes_reference      map[*DataFlowShape]*DataFlowShape
	DataFlowShapes_referenceOrder map[*DataFlowShape]uint

	// insertion point for slice of pointers maps
	OnAfterDataFlowShapeCreateCallback OnAfterCreateInterface[DataFlowShape]
	OnAfterDataFlowShapeUpdateCallback OnAfterUpdateInterface[DataFlowShape]
	OnAfterDataFlowShapeDeleteCallback OnAfterDeleteInterface[DataFlowShape]
	OnAfterDataFlowShapeReadCallback   OnAfterReadInterface[DataFlowShape]

	DiagramProcesss                map[*DiagramProcess]struct{}
	DiagramProcesss_instance       map[*DiagramProcess]*DiagramProcess
	DiagramProcesss_mapString      map[string]*DiagramProcess
	DiagramProcessOrder            uint
	DiagramProcess_stagedOrder     map[*DiagramProcess]uint
	DiagramProcess_orderStaged     map[uint]*DiagramProcess
	DiagramProcesss_reference      map[*DiagramProcess]*DiagramProcess
	DiagramProcesss_referenceOrder map[*DiagramProcess]uint

	// insertion point for slice of pointers maps
	DiagramProcess_Process_Shapes_reverseMap map[*ProcessShape]*DiagramProcess

	DiagramProcess_ProcesssWhoseNodeIsExpanded_reverseMap map[*Process]*DiagramProcess

	DiagramProcess_Participant_Shapes_reverseMap map[*ParticipantShape]*DiagramProcess

	DiagramProcess_ParticipantWhoseNodeIsExpanded_reverseMap map[*Participant]*DiagramProcess

	DiagramProcess_TasksWhoseNodeIsExpanded_reverseMap map[*Task]*DiagramProcess

	DiagramProcess_TaskShapes_reverseMap map[*TaskShape]*DiagramProcess

	DiagramProcess_ControlFlowsWhoseNodeIsExpanded_reverseMap map[*ControlFlow]*DiagramProcess

	DiagramProcess_ControlFlowShapes_reverseMap map[*ControlFlowShape]*DiagramProcess

	DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap map[*DataFlow]*DiagramProcess

	DiagramProcess_DataFlowShapes_reverseMap map[*DataFlowShape]*DiagramProcess

	OnAfterDiagramProcessCreateCallback OnAfterCreateInterface[DiagramProcess]
	OnAfterDiagramProcessUpdateCallback OnAfterUpdateInterface[DiagramProcess]
	OnAfterDiagramProcessDeleteCallback OnAfterDeleteInterface[DiagramProcess]
	OnAfterDiagramProcessReadCallback   OnAfterReadInterface[DiagramProcess]

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

	Library_RootProcesses_reverseMap map[*Process]*Library

	Library_ProcesssWhoseNodeIsExpanded_reverseMap map[*Process]*Library

	Library_RootDataFlows_reverseMap map[*DataFlow]*Library

	Library_DataFlowsWhoseNodeIsExpanded_reverseMap map[*DataFlow]*Library

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	Participants                map[*Participant]struct{}
	Participants_instance       map[*Participant]*Participant
	Participants_mapString      map[string]*Participant
	ParticipantOrder            uint
	Participant_stagedOrder     map[*Participant]uint
	Participant_orderStaged     map[uint]*Participant
	Participants_reference      map[*Participant]*Participant
	Participants_referenceOrder map[*Participant]uint

	// insertion point for slice of pointers maps
	Participant_Tasks_reverseMap map[*Task]*Participant

	Participant_ControlFlows_reverseMap map[*ControlFlow]*Participant

	Participant_TaskWhoseOutControlFlowsNodeIsExpanded_reverseMap map[*Task]*Participant

	Participant_TaskWhoseInControlFlowsNodeIsExpanded_reverseMap map[*Task]*Participant

	OnAfterParticipantCreateCallback OnAfterCreateInterface[Participant]
	OnAfterParticipantUpdateCallback OnAfterUpdateInterface[Participant]
	OnAfterParticipantDeleteCallback OnAfterDeleteInterface[Participant]
	OnAfterParticipantReadCallback   OnAfterReadInterface[Participant]

	ParticipantShapes                map[*ParticipantShape]struct{}
	ParticipantShapes_instance       map[*ParticipantShape]*ParticipantShape
	ParticipantShapes_mapString      map[string]*ParticipantShape
	ParticipantShapeOrder            uint
	ParticipantShape_stagedOrder     map[*ParticipantShape]uint
	ParticipantShape_orderStaged     map[uint]*ParticipantShape
	ParticipantShapes_reference      map[*ParticipantShape]*ParticipantShape
	ParticipantShapes_referenceOrder map[*ParticipantShape]uint

	// insertion point for slice of pointers maps
	OnAfterParticipantShapeCreateCallback OnAfterCreateInterface[ParticipantShape]
	OnAfterParticipantShapeUpdateCallback OnAfterUpdateInterface[ParticipantShape]
	OnAfterParticipantShapeDeleteCallback OnAfterDeleteInterface[ParticipantShape]
	OnAfterParticipantShapeReadCallback   OnAfterReadInterface[ParticipantShape]

	Processs                map[*Process]struct{}
	Processs_instance       map[*Process]*Process
	Processs_mapString      map[string]*Process
	ProcessOrder            uint
	Process_stagedOrder     map[*Process]uint
	Process_orderStaged     map[uint]*Process
	Processs_reference      map[*Process]*Process
	Processs_referenceOrder map[*Process]uint

	// insertion point for slice of pointers maps
	Process_DiagramProcesss_reverseMap map[*DiagramProcess]*Process

	Process_DiagramProcessWhoseNodeIsExpanded_reverseMap map[*DiagramProcess]*Process

	Process_SubProcesses_reverseMap map[*Process]*Process

	Process_Participants_reverseMap map[*Participant]*Process

	Process_ParticipantWhoseNodeIsExpanded_reverseMap map[*Participant]*Process

	OnAfterProcessCreateCallback OnAfterCreateInterface[Process]
	OnAfterProcessUpdateCallback OnAfterUpdateInterface[Process]
	OnAfterProcessDeleteCallback OnAfterDeleteInterface[Process]
	OnAfterProcessReadCallback   OnAfterReadInterface[Process]

	ProcessShapes                map[*ProcessShape]struct{}
	ProcessShapes_instance       map[*ProcessShape]*ProcessShape
	ProcessShapes_mapString      map[string]*ProcessShape
	ProcessShapeOrder            uint
	ProcessShape_stagedOrder     map[*ProcessShape]uint
	ProcessShape_orderStaged     map[uint]*ProcessShape
	ProcessShapes_reference      map[*ProcessShape]*ProcessShape
	ProcessShapes_referenceOrder map[*ProcessShape]uint

	// insertion point for slice of pointers maps
	OnAfterProcessShapeCreateCallback OnAfterCreateInterface[ProcessShape]
	OnAfterProcessShapeUpdateCallback OnAfterUpdateInterface[ProcessShape]
	OnAfterProcessShapeDeleteCallback OnAfterDeleteInterface[ProcessShape]
	OnAfterProcessShapeReadCallback   OnAfterReadInterface[ProcessShape]

	Tasks                map[*Task]struct{}
	Tasks_instance       map[*Task]*Task
	Tasks_mapString      map[string]*Task
	TaskOrder            uint
	Task_stagedOrder     map[*Task]uint
	Task_orderStaged     map[uint]*Task
	Tasks_reference      map[*Task]*Task
	Tasks_referenceOrder map[*Task]uint

	// insertion point for slice of pointers maps
	OnAfterTaskCreateCallback OnAfterCreateInterface[Task]
	OnAfterTaskUpdateCallback OnAfterUpdateInterface[Task]
	OnAfterTaskDeleteCallback OnAfterDeleteInterface[Task]
	OnAfterTaskReadCallback   OnAfterReadInterface[Task]

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
	stage.ControlFlows_reference = make(map[*ControlFlow]*ControlFlow)
	stage.ControlFlows_instance = make(map[*ControlFlow]*ControlFlow)
	stage.ControlFlows_referenceOrder = make(map[*ControlFlow]uint)

	stage.ControlFlowShapes_reference = make(map[*ControlFlowShape]*ControlFlowShape)
	stage.ControlFlowShapes_instance = make(map[*ControlFlowShape]*ControlFlowShape)
	stage.ControlFlowShapes_referenceOrder = make(map[*ControlFlowShape]uint)

	stage.DataFlows_reference = make(map[*DataFlow]*DataFlow)
	stage.DataFlows_instance = make(map[*DataFlow]*DataFlow)
	stage.DataFlows_referenceOrder = make(map[*DataFlow]uint)

	stage.DataFlowShapes_reference = make(map[*DataFlowShape]*DataFlowShape)
	stage.DataFlowShapes_instance = make(map[*DataFlowShape]*DataFlowShape)
	stage.DataFlowShapes_referenceOrder = make(map[*DataFlowShape]uint)

	stage.DiagramProcesss_reference = make(map[*DiagramProcess]*DiagramProcess)
	stage.DiagramProcesss_instance = make(map[*DiagramProcess]*DiagramProcess)
	stage.DiagramProcesss_referenceOrder = make(map[*DiagramProcess]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Participants_reference = make(map[*Participant]*Participant)
	stage.Participants_instance = make(map[*Participant]*Participant)
	stage.Participants_referenceOrder = make(map[*Participant]uint)

	stage.ParticipantShapes_reference = make(map[*ParticipantShape]*ParticipantShape)
	stage.ParticipantShapes_instance = make(map[*ParticipantShape]*ParticipantShape)
	stage.ParticipantShapes_referenceOrder = make(map[*ParticipantShape]uint)

	stage.Processs_reference = make(map[*Process]*Process)
	stage.Processs_instance = make(map[*Process]*Process)
	stage.Processs_referenceOrder = make(map[*Process]uint)

	stage.ProcessShapes_reference = make(map[*ProcessShape]*ProcessShape)
	stage.ProcessShapes_instance = make(map[*ProcessShape]*ProcessShape)
	stage.ProcessShapes_referenceOrder = make(map[*ProcessShape]uint)

	stage.Tasks_reference = make(map[*Task]*Task)
	stage.Tasks_instance = make(map[*Task]*Task)
	stage.Tasks_referenceOrder = make(map[*Task]uint)

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
	var maxControlFlowOrder uint
	var foundControlFlow bool
	for _, order := range stage.ControlFlow_stagedOrder {
		if !foundControlFlow || order > maxControlFlowOrder {
			maxControlFlowOrder = order
			foundControlFlow = true
		}
	}
	if foundControlFlow {
		stage.ControlFlowOrder = maxControlFlowOrder + 1
	} else {
		stage.ControlFlowOrder = 0
	}

	var maxControlFlowShapeOrder uint
	var foundControlFlowShape bool
	for _, order := range stage.ControlFlowShape_stagedOrder {
		if !foundControlFlowShape || order > maxControlFlowShapeOrder {
			maxControlFlowShapeOrder = order
			foundControlFlowShape = true
		}
	}
	if foundControlFlowShape {
		stage.ControlFlowShapeOrder = maxControlFlowShapeOrder + 1
	} else {
		stage.ControlFlowShapeOrder = 0
	}

	var maxDataFlowOrder uint
	var foundDataFlow bool
	for _, order := range stage.DataFlow_stagedOrder {
		if !foundDataFlow || order > maxDataFlowOrder {
			maxDataFlowOrder = order
			foundDataFlow = true
		}
	}
	if foundDataFlow {
		stage.DataFlowOrder = maxDataFlowOrder + 1
	} else {
		stage.DataFlowOrder = 0
	}

	var maxDataFlowShapeOrder uint
	var foundDataFlowShape bool
	for _, order := range stage.DataFlowShape_stagedOrder {
		if !foundDataFlowShape || order > maxDataFlowShapeOrder {
			maxDataFlowShapeOrder = order
			foundDataFlowShape = true
		}
	}
	if foundDataFlowShape {
		stage.DataFlowShapeOrder = maxDataFlowShapeOrder + 1
	} else {
		stage.DataFlowShapeOrder = 0
	}

	var maxDiagramProcessOrder uint
	var foundDiagramProcess bool
	for _, order := range stage.DiagramProcess_stagedOrder {
		if !foundDiagramProcess || order > maxDiagramProcessOrder {
			maxDiagramProcessOrder = order
			foundDiagramProcess = true
		}
	}
	if foundDiagramProcess {
		stage.DiagramProcessOrder = maxDiagramProcessOrder + 1
	} else {
		stage.DiagramProcessOrder = 0
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

	var maxParticipantOrder uint
	var foundParticipant bool
	for _, order := range stage.Participant_stagedOrder {
		if !foundParticipant || order > maxParticipantOrder {
			maxParticipantOrder = order
			foundParticipant = true
		}
	}
	if foundParticipant {
		stage.ParticipantOrder = maxParticipantOrder + 1
	} else {
		stage.ParticipantOrder = 0
	}

	var maxParticipantShapeOrder uint
	var foundParticipantShape bool
	for _, order := range stage.ParticipantShape_stagedOrder {
		if !foundParticipantShape || order > maxParticipantShapeOrder {
			maxParticipantShapeOrder = order
			foundParticipantShape = true
		}
	}
	if foundParticipantShape {
		stage.ParticipantShapeOrder = maxParticipantShapeOrder + 1
	} else {
		stage.ParticipantShapeOrder = 0
	}

	var maxProcessOrder uint
	var foundProcess bool
	for _, order := range stage.Process_stagedOrder {
		if !foundProcess || order > maxProcessOrder {
			maxProcessOrder = order
			foundProcess = true
		}
	}
	if foundProcess {
		stage.ProcessOrder = maxProcessOrder + 1
	} else {
		stage.ProcessOrder = 0
	}

	var maxProcessShapeOrder uint
	var foundProcessShape bool
	for _, order := range stage.ProcessShape_stagedOrder {
		if !foundProcessShape || order > maxProcessShapeOrder {
			maxProcessShapeOrder = order
			foundProcessShape = true
		}
	}
	if foundProcessShape {
		stage.ProcessShapeOrder = maxProcessShapeOrder + 1
	} else {
		stage.ProcessShapeOrder = 0
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
	case *ControlFlow:
		tmp := GetStructInstancesByOrder(stage.ControlFlows, stage.ControlFlow_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ControlFlow implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ControlFlowShape:
		tmp := GetStructInstancesByOrder(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ControlFlowShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DataFlow:
		tmp := GetStructInstancesByOrder(stage.DataFlows, stage.DataFlow_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DataFlow implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DataFlowShape:
		tmp := GetStructInstancesByOrder(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DataFlowShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DiagramProcess:
		tmp := GetStructInstancesByOrder(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DiagramProcess implements.
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
	case *Participant:
		tmp := GetStructInstancesByOrder(stage.Participants, stage.Participant_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Participant implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ParticipantShape:
		tmp := GetStructInstancesByOrder(stage.ParticipantShapes, stage.ParticipantShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ParticipantShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Process:
		tmp := GetStructInstancesByOrder(stage.Processs, stage.Process_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Process implements.
			res = append(res, any(v).(T))
		}
		return res
	case *ProcessShape:
		tmp := GetStructInstancesByOrder(stage.ProcessShapes, stage.ProcessShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ProcessShape implements.
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
	case "ControlFlow":
		res = GetNamedStructInstances(stage.ControlFlows, stage.ControlFlow_stagedOrder)
	case "ControlFlowShape":
		res = GetNamedStructInstances(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder)
	case "DataFlow":
		res = GetNamedStructInstances(stage.DataFlows, stage.DataFlow_stagedOrder)
	case "DataFlowShape":
		res = GetNamedStructInstances(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder)
	case "DiagramProcess":
		res = GetNamedStructInstances(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Participant":
		res = GetNamedStructInstances(stage.Participants, stage.Participant_stagedOrder)
	case "ParticipantShape":
		res = GetNamedStructInstances(stage.ParticipantShapes, stage.ParticipantShape_stagedOrder)
	case "Process":
		res = GetNamedStructInstances(stage.Processs, stage.Process_stagedOrder)
	case "ProcessShape":
		res = GetNamedStructInstances(stage.ProcessShapes, stage.ProcessShape_stagedOrder)
	case "Task":
		res = GetNamedStructInstances(stage.Tasks, stage.Task_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/process/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return process_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return process_go.GoDiagramsDir
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
	CommitControlFlow(controlflow *ControlFlow)
	CheckoutControlFlow(controlflow *ControlFlow)
	CommitControlFlowShape(controlflowshape *ControlFlowShape)
	CheckoutControlFlowShape(controlflowshape *ControlFlowShape)
	CommitDataFlow(dataflow *DataFlow)
	CheckoutDataFlow(dataflow *DataFlow)
	CommitDataFlowShape(dataflowshape *DataFlowShape)
	CheckoutDataFlowShape(dataflowshape *DataFlowShape)
	CommitDiagramProcess(diagramprocess *DiagramProcess)
	CheckoutDiagramProcess(diagramprocess *DiagramProcess)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitParticipant(participant *Participant)
	CheckoutParticipant(participant *Participant)
	CommitParticipantShape(participantshape *ParticipantShape)
	CheckoutParticipantShape(participantshape *ParticipantShape)
	CommitProcess(process *Process)
	CheckoutProcess(process *Process)
	CommitProcessShape(processshape *ProcessShape)
	CheckoutProcessShape(processshape *ProcessShape)
	CommitTask(task *Task)
	CheckoutTask(task *Task)
	CommitTaskShape(taskshape *TaskShape)
	CheckoutTaskShape(taskshape *TaskShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		ControlFlows:           make(map[*ControlFlow]struct{}),
		ControlFlows_mapString: make(map[string]*ControlFlow),

		ControlFlowShapes:           make(map[*ControlFlowShape]struct{}),
		ControlFlowShapes_mapString: make(map[string]*ControlFlowShape),

		DataFlows:           make(map[*DataFlow]struct{}),
		DataFlows_mapString: make(map[string]*DataFlow),

		DataFlowShapes:           make(map[*DataFlowShape]struct{}),
		DataFlowShapes_mapString: make(map[string]*DataFlowShape),

		DiagramProcesss:           make(map[*DiagramProcess]struct{}),
		DiagramProcesss_mapString: make(map[string]*DiagramProcess),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Participants:           make(map[*Participant]struct{}),
		Participants_mapString: make(map[string]*Participant),

		ParticipantShapes:           make(map[*ParticipantShape]struct{}),
		ParticipantShapes_mapString: make(map[string]*ParticipantShape),

		Processs:           make(map[*Process]struct{}),
		Processs_mapString: make(map[string]*Process),

		ProcessShapes:           make(map[*ProcessShape]struct{}),
		ProcessShapes_mapString: make(map[string]*ProcessShape),

		Tasks:           make(map[*Task]struct{}),
		Tasks_mapString: make(map[string]*Task),

		TaskShapes:           make(map[*TaskShape]struct{}),
		TaskShapes_mapString: make(map[string]*TaskShape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ControlFlow_stagedOrder: make(map[*ControlFlow]uint),
		ControlFlow_orderStaged: make(map[uint]*ControlFlow),
		ControlFlows_reference: make(map[*ControlFlow]*ControlFlow),

		ControlFlowShape_stagedOrder: make(map[*ControlFlowShape]uint),
		ControlFlowShape_orderStaged: make(map[uint]*ControlFlowShape),
		ControlFlowShapes_reference: make(map[*ControlFlowShape]*ControlFlowShape),

		DataFlow_stagedOrder: make(map[*DataFlow]uint),
		DataFlow_orderStaged: make(map[uint]*DataFlow),
		DataFlows_reference: make(map[*DataFlow]*DataFlow),

		DataFlowShape_stagedOrder: make(map[*DataFlowShape]uint),
		DataFlowShape_orderStaged: make(map[uint]*DataFlowShape),
		DataFlowShapes_reference: make(map[*DataFlowShape]*DataFlowShape),

		DiagramProcess_stagedOrder: make(map[*DiagramProcess]uint),
		DiagramProcess_orderStaged: make(map[uint]*DiagramProcess),
		DiagramProcesss_reference: make(map[*DiagramProcess]*DiagramProcess),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference: make(map[*Library]*Library),

		Participant_stagedOrder: make(map[*Participant]uint),
		Participant_orderStaged: make(map[uint]*Participant),
		Participants_reference: make(map[*Participant]*Participant),

		ParticipantShape_stagedOrder: make(map[*ParticipantShape]uint),
		ParticipantShape_orderStaged: make(map[uint]*ParticipantShape),
		ParticipantShapes_reference: make(map[*ParticipantShape]*ParticipantShape),

		Process_stagedOrder: make(map[*Process]uint),
		Process_orderStaged: make(map[uint]*Process),
		Processs_reference: make(map[*Process]*Process),

		ProcessShape_stagedOrder: make(map[*ProcessShape]uint),
		ProcessShape_orderStaged: make(map[uint]*ProcessShape),
		ProcessShapes_reference: make(map[*ProcessShape]*ProcessShape),

		Task_stagedOrder: make(map[*Task]uint),
		Task_orderStaged: make(map[uint]*Task),
		Tasks_reference: make(map[*Task]*Task),

		TaskShape_stagedOrder: make(map[*TaskShape]uint),
		TaskShape_orderStaged: make(map[uint]*TaskShape),
		TaskShapes_reference: make(map[*TaskShape]*TaskShape),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"ControlFlow": &ControlFlowUnmarshaller{},

			"ControlFlowShape": &ControlFlowShapeUnmarshaller{},

			"DataFlow": &DataFlowUnmarshaller{},

			"DataFlowShape": &DataFlowShapeUnmarshaller{},

			"DiagramProcess": &DiagramProcessUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Participant": &ParticipantUnmarshaller{},

			"ParticipantShape": &ParticipantShapeUnmarshaller{},

			"Process": &ProcessUnmarshaller{},

			"ProcessShape": &ProcessShapeUnmarshaller{},

			"Task": &TaskUnmarshaller{},

			"TaskShape": &TaskShapeUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "ControlFlow"},
			{name: "ControlFlowShape"},
			{name: "DataFlow"},
			{name: "DataFlowShape"},
			{name: "DiagramProcess"},
			{name: "Library"},
			{name: "Participant"},
			{name: "ParticipantShape"},
			{name: "Process"},
			{name: "ProcessShape"},
			{name: "Task"},
			{name: "TaskShape"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ControlFlow:
		return stage.ControlFlow_stagedOrder[instance]
	case *ControlFlowShape:
		return stage.ControlFlowShape_stagedOrder[instance]
	case *DataFlow:
		return stage.DataFlow_stagedOrder[instance]
	case *DataFlowShape:
		return stage.DataFlowShape_stagedOrder[instance]
	case *DiagramProcess:
		return stage.DiagramProcess_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Participant:
		return stage.Participant_stagedOrder[instance]
	case *ParticipantShape:
		return stage.ParticipantShape_stagedOrder[instance]
	case *Process:
		return stage.Process_stagedOrder[instance]
	case *ProcessShape:
		return stage.ProcessShape_stagedOrder[instance]
	case *Task:
		return stage.Task_stagedOrder[instance]
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
	case *ControlFlow:
		return any(stage.ControlFlow_orderStaged[order]).(Type)
	case *ControlFlowShape:
		return any(stage.ControlFlowShape_orderStaged[order]).(Type)
	case *DataFlow:
		return any(stage.DataFlow_orderStaged[order]).(Type)
	case *DataFlowShape:
		return any(stage.DataFlowShape_orderStaged[order]).(Type)
	case *DiagramProcess:
		return any(stage.DiagramProcess_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Participant:
		return any(stage.Participant_orderStaged[order]).(Type)
	case *ParticipantShape:
		return any(stage.ParticipantShape_orderStaged[order]).(Type)
	case *Process:
		return any(stage.Process_orderStaged[order]).(Type)
	case *ProcessShape:
		return any(stage.ProcessShape_orderStaged[order]).(Type)
	case *Task:
		return any(stage.Task_orderStaged[order]).(Type)
	case *TaskShape:
		return any(stage.TaskShape_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *ControlFlow:
		return stage.ControlFlow_stagedOrder[instance]
	case *ControlFlowShape:
		return stage.ControlFlowShape_stagedOrder[instance]
	case *DataFlow:
		return stage.DataFlow_stagedOrder[instance]
	case *DataFlowShape:
		return stage.DataFlowShape_stagedOrder[instance]
	case *DiagramProcess:
		return stage.DiagramProcess_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Participant:
		return stage.Participant_stagedOrder[instance]
	case *ParticipantShape:
		return stage.ParticipantShape_stagedOrder[instance]
	case *Process:
		return stage.Process_stagedOrder[instance]
	case *ProcessShape:
		return stage.ProcessShape_stagedOrder[instance]
	case *Task:
		return stage.Task_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["ControlFlow"] = len(stage.ControlFlows)
	stage.Map_GongStructName_InstancesNb["ControlFlowShape"] = len(stage.ControlFlowShapes)
	stage.Map_GongStructName_InstancesNb["DataFlow"] = len(stage.DataFlows)
	stage.Map_GongStructName_InstancesNb["DataFlowShape"] = len(stage.DataFlowShapes)
	stage.Map_GongStructName_InstancesNb["DiagramProcess"] = len(stage.DiagramProcesss)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Participant"] = len(stage.Participants)
	stage.Map_GongStructName_InstancesNb["ParticipantShape"] = len(stage.ParticipantShapes)
	stage.Map_GongStructName_InstancesNb["Process"] = len(stage.Processs)
	stage.Map_GongStructName_InstancesNb["ProcessShape"] = len(stage.ProcessShapes)
	stage.Map_GongStructName_InstancesNb["Task"] = len(stage.Tasks)
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
// Stage puts controlflow to the model stage
func (controlflow *ControlFlow) Stage(stage *Stage) *ControlFlow {
	if _, ok := stage.ControlFlows[controlflow]; !ok {
		stage.ControlFlows[controlflow] = struct{}{}
		stage.ControlFlow_stagedOrder[controlflow] = stage.ControlFlowOrder
		stage.ControlFlow_orderStaged[stage.ControlFlowOrder] = controlflow
		stage.ControlFlowOrder++
	}
	stage.ControlFlows_mapString[controlflow.Name] = controlflow

	return controlflow
}

// StagePreserveOrder puts controlflow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlFlowOrder
// - update stage.ControlFlowOrder accordingly
func (controlflow *ControlFlow) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ControlFlows[controlflow]; !ok {
		stage.ControlFlows[controlflow] = struct{}{}

		if order > stage.ControlFlowOrder {
			stage.ControlFlowOrder = order
		}
		stage.ControlFlow_stagedOrder[controlflow] = order
		stage.ControlFlow_orderStaged[order] = controlflow
		stage.ControlFlowOrder++
	}
	stage.ControlFlows_mapString[controlflow.Name] = controlflow
}

// Unstage removes controlflow off the model stage
func (controlflow *ControlFlow) Unstage(stage *Stage) *ControlFlow {
	delete(stage.ControlFlows, controlflow)
	// issue1150
	// delete(stage.ControlFlow_stagedOrder, controlflow)
	delete(stage.ControlFlows_mapString, controlflow.Name)

	return controlflow
}

// UnstageVoid removes controlflow off the model stage
func (controlflow *ControlFlow) UnstageVoid(stage *Stage) {
	delete(stage.ControlFlows, controlflow)
	// issue1150
	// delete(stage.ControlFlow_stagedOrder, controlflow)
	delete(stage.ControlFlows_mapString, controlflow.Name)
}

// commit controlflow to the back repo (if it is already staged)
func (controlflow *ControlFlow) Commit(stage *Stage) *ControlFlow {
	if _, ok := stage.ControlFlows[controlflow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitControlFlow(controlflow)
		}
	}
	return controlflow
}

func (controlflow *ControlFlow) CommitVoid(stage *Stage) {
	controlflow.Commit(stage)
}

func (controlflow *ControlFlow) StageVoid(stage *Stage) {
	controlflow.Stage(stage)
}

// Checkout controlflow to the back repo (if it is already staged)
func (controlflow *ControlFlow) Checkout(stage *Stage) *ControlFlow {
	if _, ok := stage.ControlFlows[controlflow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutControlFlow(controlflow)
		}
	}
	return controlflow
}

// for satisfaction of GongStruct interface
func (controlflow *ControlFlow) GetName() (res string) {
	return controlflow.Name
}

// for satisfaction of GongStruct interface
func (controlflow *ControlFlow) SetName(name string) {
	controlflow.Name = name
}

// Stage puts controlflowshape to the model stage
func (controlflowshape *ControlFlowShape) Stage(stage *Stage) *ControlFlowShape {
	if _, ok := stage.ControlFlowShapes[controlflowshape]; !ok {
		stage.ControlFlowShapes[controlflowshape] = struct{}{}
		stage.ControlFlowShape_stagedOrder[controlflowshape] = stage.ControlFlowShapeOrder
		stage.ControlFlowShape_orderStaged[stage.ControlFlowShapeOrder] = controlflowshape
		stage.ControlFlowShapeOrder++
	}
	stage.ControlFlowShapes_mapString[controlflowshape.Name] = controlflowshape

	return controlflowshape
}

// StagePreserveOrder puts controlflowshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlFlowShapeOrder
// - update stage.ControlFlowShapeOrder accordingly
func (controlflowshape *ControlFlowShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ControlFlowShapes[controlflowshape]; !ok {
		stage.ControlFlowShapes[controlflowshape] = struct{}{}

		if order > stage.ControlFlowShapeOrder {
			stage.ControlFlowShapeOrder = order
		}
		stage.ControlFlowShape_stagedOrder[controlflowshape] = order
		stage.ControlFlowShape_orderStaged[order] = controlflowshape
		stage.ControlFlowShapeOrder++
	}
	stage.ControlFlowShapes_mapString[controlflowshape.Name] = controlflowshape
}

// Unstage removes controlflowshape off the model stage
func (controlflowshape *ControlFlowShape) Unstage(stage *Stage) *ControlFlowShape {
	delete(stage.ControlFlowShapes, controlflowshape)
	// issue1150
	// delete(stage.ControlFlowShape_stagedOrder, controlflowshape)
	delete(stage.ControlFlowShapes_mapString, controlflowshape.Name)

	return controlflowshape
}

// UnstageVoid removes controlflowshape off the model stage
func (controlflowshape *ControlFlowShape) UnstageVoid(stage *Stage) {
	delete(stage.ControlFlowShapes, controlflowshape)
	// issue1150
	// delete(stage.ControlFlowShape_stagedOrder, controlflowshape)
	delete(stage.ControlFlowShapes_mapString, controlflowshape.Name)
}

// commit controlflowshape to the back repo (if it is already staged)
func (controlflowshape *ControlFlowShape) Commit(stage *Stage) *ControlFlowShape {
	if _, ok := stage.ControlFlowShapes[controlflowshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitControlFlowShape(controlflowshape)
		}
	}
	return controlflowshape
}

func (controlflowshape *ControlFlowShape) CommitVoid(stage *Stage) {
	controlflowshape.Commit(stage)
}

func (controlflowshape *ControlFlowShape) StageVoid(stage *Stage) {
	controlflowshape.Stage(stage)
}

// Checkout controlflowshape to the back repo (if it is already staged)
func (controlflowshape *ControlFlowShape) Checkout(stage *Stage) *ControlFlowShape {
	if _, ok := stage.ControlFlowShapes[controlflowshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutControlFlowShape(controlflowshape)
		}
	}
	return controlflowshape
}

// for satisfaction of GongStruct interface
func (controlflowshape *ControlFlowShape) GetName() (res string) {
	return controlflowshape.Name
}

// for satisfaction of GongStruct interface
func (controlflowshape *ControlFlowShape) SetName(name string) {
	controlflowshape.Name = name
}

// Stage puts dataflow to the model stage
func (dataflow *DataFlow) Stage(stage *Stage) *DataFlow {
	if _, ok := stage.DataFlows[dataflow]; !ok {
		stage.DataFlows[dataflow] = struct{}{}
		stage.DataFlow_stagedOrder[dataflow] = stage.DataFlowOrder
		stage.DataFlow_orderStaged[stage.DataFlowOrder] = dataflow
		stage.DataFlowOrder++
	}
	stage.DataFlows_mapString[dataflow.Name] = dataflow

	return dataflow
}

// StagePreserveOrder puts dataflow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataFlowOrder
// - update stage.DataFlowOrder accordingly
func (dataflow *DataFlow) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DataFlows[dataflow]; !ok {
		stage.DataFlows[dataflow] = struct{}{}

		if order > stage.DataFlowOrder {
			stage.DataFlowOrder = order
		}
		stage.DataFlow_stagedOrder[dataflow] = order
		stage.DataFlow_orderStaged[order] = dataflow
		stage.DataFlowOrder++
	}
	stage.DataFlows_mapString[dataflow.Name] = dataflow
}

// Unstage removes dataflow off the model stage
func (dataflow *DataFlow) Unstage(stage *Stage) *DataFlow {
	delete(stage.DataFlows, dataflow)
	// issue1150
	// delete(stage.DataFlow_stagedOrder, dataflow)
	delete(stage.DataFlows_mapString, dataflow.Name)

	return dataflow
}

// UnstageVoid removes dataflow off the model stage
func (dataflow *DataFlow) UnstageVoid(stage *Stage) {
	delete(stage.DataFlows, dataflow)
	// issue1150
	// delete(stage.DataFlow_stagedOrder, dataflow)
	delete(stage.DataFlows_mapString, dataflow.Name)
}

// commit dataflow to the back repo (if it is already staged)
func (dataflow *DataFlow) Commit(stage *Stage) *DataFlow {
	if _, ok := stage.DataFlows[dataflow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDataFlow(dataflow)
		}
	}
	return dataflow
}

func (dataflow *DataFlow) CommitVoid(stage *Stage) {
	dataflow.Commit(stage)
}

func (dataflow *DataFlow) StageVoid(stage *Stage) {
	dataflow.Stage(stage)
}

// Checkout dataflow to the back repo (if it is already staged)
func (dataflow *DataFlow) Checkout(stage *Stage) *DataFlow {
	if _, ok := stage.DataFlows[dataflow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDataFlow(dataflow)
		}
	}
	return dataflow
}

// for satisfaction of GongStruct interface
func (dataflow *DataFlow) GetName() (res string) {
	return dataflow.Name
}

// for satisfaction of GongStruct interface
func (dataflow *DataFlow) SetName(name string) {
	dataflow.Name = name
}

// Stage puts dataflowshape to the model stage
func (dataflowshape *DataFlowShape) Stage(stage *Stage) *DataFlowShape {
	if _, ok := stage.DataFlowShapes[dataflowshape]; !ok {
		stage.DataFlowShapes[dataflowshape] = struct{}{}
		stage.DataFlowShape_stagedOrder[dataflowshape] = stage.DataFlowShapeOrder
		stage.DataFlowShape_orderStaged[stage.DataFlowShapeOrder] = dataflowshape
		stage.DataFlowShapeOrder++
	}
	stage.DataFlowShapes_mapString[dataflowshape.Name] = dataflowshape

	return dataflowshape
}

// StagePreserveOrder puts dataflowshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataFlowShapeOrder
// - update stage.DataFlowShapeOrder accordingly
func (dataflowshape *DataFlowShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DataFlowShapes[dataflowshape]; !ok {
		stage.DataFlowShapes[dataflowshape] = struct{}{}

		if order > stage.DataFlowShapeOrder {
			stage.DataFlowShapeOrder = order
		}
		stage.DataFlowShape_stagedOrder[dataflowshape] = order
		stage.DataFlowShape_orderStaged[order] = dataflowshape
		stage.DataFlowShapeOrder++
	}
	stage.DataFlowShapes_mapString[dataflowshape.Name] = dataflowshape
}

// Unstage removes dataflowshape off the model stage
func (dataflowshape *DataFlowShape) Unstage(stage *Stage) *DataFlowShape {
	delete(stage.DataFlowShapes, dataflowshape)
	// issue1150
	// delete(stage.DataFlowShape_stagedOrder, dataflowshape)
	delete(stage.DataFlowShapes_mapString, dataflowshape.Name)

	return dataflowshape
}

// UnstageVoid removes dataflowshape off the model stage
func (dataflowshape *DataFlowShape) UnstageVoid(stage *Stage) {
	delete(stage.DataFlowShapes, dataflowshape)
	// issue1150
	// delete(stage.DataFlowShape_stagedOrder, dataflowshape)
	delete(stage.DataFlowShapes_mapString, dataflowshape.Name)
}

// commit dataflowshape to the back repo (if it is already staged)
func (dataflowshape *DataFlowShape) Commit(stage *Stage) *DataFlowShape {
	if _, ok := stage.DataFlowShapes[dataflowshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDataFlowShape(dataflowshape)
		}
	}
	return dataflowshape
}

func (dataflowshape *DataFlowShape) CommitVoid(stage *Stage) {
	dataflowshape.Commit(stage)
}

func (dataflowshape *DataFlowShape) StageVoid(stage *Stage) {
	dataflowshape.Stage(stage)
}

// Checkout dataflowshape to the back repo (if it is already staged)
func (dataflowshape *DataFlowShape) Checkout(stage *Stage) *DataFlowShape {
	if _, ok := stage.DataFlowShapes[dataflowshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDataFlowShape(dataflowshape)
		}
	}
	return dataflowshape
}

// for satisfaction of GongStruct interface
func (dataflowshape *DataFlowShape) GetName() (res string) {
	return dataflowshape.Name
}

// for satisfaction of GongStruct interface
func (dataflowshape *DataFlowShape) SetName(name string) {
	dataflowshape.Name = name
}

// Stage puts diagramprocess to the model stage
func (diagramprocess *DiagramProcess) Stage(stage *Stage) *DiagramProcess {
	if _, ok := stage.DiagramProcesss[diagramprocess]; !ok {
		stage.DiagramProcesss[diagramprocess] = struct{}{}
		stage.DiagramProcess_stagedOrder[diagramprocess] = stage.DiagramProcessOrder
		stage.DiagramProcess_orderStaged[stage.DiagramProcessOrder] = diagramprocess
		stage.DiagramProcessOrder++
	}
	stage.DiagramProcesss_mapString[diagramprocess.Name] = diagramprocess

	return diagramprocess
}

// StagePreserveOrder puts diagramprocess to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramProcessOrder
// - update stage.DiagramProcessOrder accordingly
func (diagramprocess *DiagramProcess) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DiagramProcesss[diagramprocess]; !ok {
		stage.DiagramProcesss[diagramprocess] = struct{}{}

		if order > stage.DiagramProcessOrder {
			stage.DiagramProcessOrder = order
		}
		stage.DiagramProcess_stagedOrder[diagramprocess] = order
		stage.DiagramProcess_orderStaged[order] = diagramprocess
		stage.DiagramProcessOrder++
	}
	stage.DiagramProcesss_mapString[diagramprocess.Name] = diagramprocess
}

// Unstage removes diagramprocess off the model stage
func (diagramprocess *DiagramProcess) Unstage(stage *Stage) *DiagramProcess {
	delete(stage.DiagramProcesss, diagramprocess)
	// issue1150
	// delete(stage.DiagramProcess_stagedOrder, diagramprocess)
	delete(stage.DiagramProcesss_mapString, diagramprocess.Name)

	return diagramprocess
}

// UnstageVoid removes diagramprocess off the model stage
func (diagramprocess *DiagramProcess) UnstageVoid(stage *Stage) {
	delete(stage.DiagramProcesss, diagramprocess)
	// issue1150
	// delete(stage.DiagramProcess_stagedOrder, diagramprocess)
	delete(stage.DiagramProcesss_mapString, diagramprocess.Name)
}

// commit diagramprocess to the back repo (if it is already staged)
func (diagramprocess *DiagramProcess) Commit(stage *Stage) *DiagramProcess {
	if _, ok := stage.DiagramProcesss[diagramprocess]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDiagramProcess(diagramprocess)
		}
	}
	return diagramprocess
}

func (diagramprocess *DiagramProcess) CommitVoid(stage *Stage) {
	diagramprocess.Commit(stage)
}

func (diagramprocess *DiagramProcess) StageVoid(stage *Stage) {
	diagramprocess.Stage(stage)
}

// Checkout diagramprocess to the back repo (if it is already staged)
func (diagramprocess *DiagramProcess) Checkout(stage *Stage) *DiagramProcess {
	if _, ok := stage.DiagramProcesss[diagramprocess]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDiagramProcess(diagramprocess)
		}
	}
	return diagramprocess
}

// for satisfaction of GongStruct interface
func (diagramprocess *DiagramProcess) GetName() (res string) {
	return diagramprocess.Name
}

// for satisfaction of GongStruct interface
func (diagramprocess *DiagramProcess) SetName(name string) {
	diagramprocess.Name = name
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

// Stage puts participant to the model stage
func (participant *Participant) Stage(stage *Stage) *Participant {
	if _, ok := stage.Participants[participant]; !ok {
		stage.Participants[participant] = struct{}{}
		stage.Participant_stagedOrder[participant] = stage.ParticipantOrder
		stage.Participant_orderStaged[stage.ParticipantOrder] = participant
		stage.ParticipantOrder++
	}
	stage.Participants_mapString[participant.Name] = participant

	return participant
}

// StagePreserveOrder puts participant to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParticipantOrder
// - update stage.ParticipantOrder accordingly
func (participant *Participant) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Participants[participant]; !ok {
		stage.Participants[participant] = struct{}{}

		if order > stage.ParticipantOrder {
			stage.ParticipantOrder = order
		}
		stage.Participant_stagedOrder[participant] = order
		stage.Participant_orderStaged[order] = participant
		stage.ParticipantOrder++
	}
	stage.Participants_mapString[participant.Name] = participant
}

// Unstage removes participant off the model stage
func (participant *Participant) Unstage(stage *Stage) *Participant {
	delete(stage.Participants, participant)
	// issue1150
	// delete(stage.Participant_stagedOrder, participant)
	delete(stage.Participants_mapString, participant.Name)

	return participant
}

// UnstageVoid removes participant off the model stage
func (participant *Participant) UnstageVoid(stage *Stage) {
	delete(stage.Participants, participant)
	// issue1150
	// delete(stage.Participant_stagedOrder, participant)
	delete(stage.Participants_mapString, participant.Name)
}

// commit participant to the back repo (if it is already staged)
func (participant *Participant) Commit(stage *Stage) *Participant {
	if _, ok := stage.Participants[participant]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParticipant(participant)
		}
	}
	return participant
}

func (participant *Participant) CommitVoid(stage *Stage) {
	participant.Commit(stage)
}

func (participant *Participant) StageVoid(stage *Stage) {
	participant.Stage(stage)
}

// Checkout participant to the back repo (if it is already staged)
func (participant *Participant) Checkout(stage *Stage) *Participant {
	if _, ok := stage.Participants[participant]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParticipant(participant)
		}
	}
	return participant
}

// for satisfaction of GongStruct interface
func (participant *Participant) GetName() (res string) {
	return participant.Name
}

// for satisfaction of GongStruct interface
func (participant *Participant) SetName(name string) {
	participant.Name = name
}

// Stage puts participantshape to the model stage
func (participantshape *ParticipantShape) Stage(stage *Stage) *ParticipantShape {
	if _, ok := stage.ParticipantShapes[participantshape]; !ok {
		stage.ParticipantShapes[participantshape] = struct{}{}
		stage.ParticipantShape_stagedOrder[participantshape] = stage.ParticipantShapeOrder
		stage.ParticipantShape_orderStaged[stage.ParticipantShapeOrder] = participantshape
		stage.ParticipantShapeOrder++
	}
	stage.ParticipantShapes_mapString[participantshape.Name] = participantshape

	return participantshape
}

// StagePreserveOrder puts participantshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParticipantShapeOrder
// - update stage.ParticipantShapeOrder accordingly
func (participantshape *ParticipantShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ParticipantShapes[participantshape]; !ok {
		stage.ParticipantShapes[participantshape] = struct{}{}

		if order > stage.ParticipantShapeOrder {
			stage.ParticipantShapeOrder = order
		}
		stage.ParticipantShape_stagedOrder[participantshape] = order
		stage.ParticipantShape_orderStaged[order] = participantshape
		stage.ParticipantShapeOrder++
	}
	stage.ParticipantShapes_mapString[participantshape.Name] = participantshape
}

// Unstage removes participantshape off the model stage
func (participantshape *ParticipantShape) Unstage(stage *Stage) *ParticipantShape {
	delete(stage.ParticipantShapes, participantshape)
	// issue1150
	// delete(stage.ParticipantShape_stagedOrder, participantshape)
	delete(stage.ParticipantShapes_mapString, participantshape.Name)

	return participantshape
}

// UnstageVoid removes participantshape off the model stage
func (participantshape *ParticipantShape) UnstageVoid(stage *Stage) {
	delete(stage.ParticipantShapes, participantshape)
	// issue1150
	// delete(stage.ParticipantShape_stagedOrder, participantshape)
	delete(stage.ParticipantShapes_mapString, participantshape.Name)
}

// commit participantshape to the back repo (if it is already staged)
func (participantshape *ParticipantShape) Commit(stage *Stage) *ParticipantShape {
	if _, ok := stage.ParticipantShapes[participantshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitParticipantShape(participantshape)
		}
	}
	return participantshape
}

func (participantshape *ParticipantShape) CommitVoid(stage *Stage) {
	participantshape.Commit(stage)
}

func (participantshape *ParticipantShape) StageVoid(stage *Stage) {
	participantshape.Stage(stage)
}

// Checkout participantshape to the back repo (if it is already staged)
func (participantshape *ParticipantShape) Checkout(stage *Stage) *ParticipantShape {
	if _, ok := stage.ParticipantShapes[participantshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutParticipantShape(participantshape)
		}
	}
	return participantshape
}

// for satisfaction of GongStruct interface
func (participantshape *ParticipantShape) GetName() (res string) {
	return participantshape.Name
}

// for satisfaction of GongStruct interface
func (participantshape *ParticipantShape) SetName(name string) {
	participantshape.Name = name
}

// Stage puts process to the model stage
func (process *Process) Stage(stage *Stage) *Process {
	if _, ok := stage.Processs[process]; !ok {
		stage.Processs[process] = struct{}{}
		stage.Process_stagedOrder[process] = stage.ProcessOrder
		stage.Process_orderStaged[stage.ProcessOrder] = process
		stage.ProcessOrder++
	}
	stage.Processs_mapString[process.Name] = process

	return process
}

// StagePreserveOrder puts process to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProcessOrder
// - update stage.ProcessOrder accordingly
func (process *Process) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Processs[process]; !ok {
		stage.Processs[process] = struct{}{}

		if order > stage.ProcessOrder {
			stage.ProcessOrder = order
		}
		stage.Process_stagedOrder[process] = order
		stage.Process_orderStaged[order] = process
		stage.ProcessOrder++
	}
	stage.Processs_mapString[process.Name] = process
}

// Unstage removes process off the model stage
func (process *Process) Unstage(stage *Stage) *Process {
	delete(stage.Processs, process)
	// issue1150
	// delete(stage.Process_stagedOrder, process)
	delete(stage.Processs_mapString, process.Name)

	return process
}

// UnstageVoid removes process off the model stage
func (process *Process) UnstageVoid(stage *Stage) {
	delete(stage.Processs, process)
	// issue1150
	// delete(stage.Process_stagedOrder, process)
	delete(stage.Processs_mapString, process.Name)
}

// commit process to the back repo (if it is already staged)
func (process *Process) Commit(stage *Stage) *Process {
	if _, ok := stage.Processs[process]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProcess(process)
		}
	}
	return process
}

func (process *Process) CommitVoid(stage *Stage) {
	process.Commit(stage)
}

func (process *Process) StageVoid(stage *Stage) {
	process.Stage(stage)
}

// Checkout process to the back repo (if it is already staged)
func (process *Process) Checkout(stage *Stage) *Process {
	if _, ok := stage.Processs[process]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProcess(process)
		}
	}
	return process
}

// for satisfaction of GongStruct interface
func (process *Process) GetName() (res string) {
	return process.Name
}

// for satisfaction of GongStruct interface
func (process *Process) SetName(name string) {
	process.Name = name
}

// Stage puts processshape to the model stage
func (processshape *ProcessShape) Stage(stage *Stage) *ProcessShape {
	if _, ok := stage.ProcessShapes[processshape]; !ok {
		stage.ProcessShapes[processshape] = struct{}{}
		stage.ProcessShape_stagedOrder[processshape] = stage.ProcessShapeOrder
		stage.ProcessShape_orderStaged[stage.ProcessShapeOrder] = processshape
		stage.ProcessShapeOrder++
	}
	stage.ProcessShapes_mapString[processshape.Name] = processshape

	return processshape
}

// StagePreserveOrder puts processshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProcessShapeOrder
// - update stage.ProcessShapeOrder accordingly
func (processshape *ProcessShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ProcessShapes[processshape]; !ok {
		stage.ProcessShapes[processshape] = struct{}{}

		if order > stage.ProcessShapeOrder {
			stage.ProcessShapeOrder = order
		}
		stage.ProcessShape_stagedOrder[processshape] = order
		stage.ProcessShape_orderStaged[order] = processshape
		stage.ProcessShapeOrder++
	}
	stage.ProcessShapes_mapString[processshape.Name] = processshape
}

// Unstage removes processshape off the model stage
func (processshape *ProcessShape) Unstage(stage *Stage) *ProcessShape {
	delete(stage.ProcessShapes, processshape)
	// issue1150
	// delete(stage.ProcessShape_stagedOrder, processshape)
	delete(stage.ProcessShapes_mapString, processshape.Name)

	return processshape
}

// UnstageVoid removes processshape off the model stage
func (processshape *ProcessShape) UnstageVoid(stage *Stage) {
	delete(stage.ProcessShapes, processshape)
	// issue1150
	// delete(stage.ProcessShape_stagedOrder, processshape)
	delete(stage.ProcessShapes_mapString, processshape.Name)
}

// commit processshape to the back repo (if it is already staged)
func (processshape *ProcessShape) Commit(stage *Stage) *ProcessShape {
	if _, ok := stage.ProcessShapes[processshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProcessShape(processshape)
		}
	}
	return processshape
}

func (processshape *ProcessShape) CommitVoid(stage *Stage) {
	processshape.Commit(stage)
}

func (processshape *ProcessShape) StageVoid(stage *Stage) {
	processshape.Stage(stage)
}

// Checkout processshape to the back repo (if it is already staged)
func (processshape *ProcessShape) Checkout(stage *Stage) *ProcessShape {
	if _, ok := stage.ProcessShapes[processshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProcessShape(processshape)
		}
	}
	return processshape
}

// for satisfaction of GongStruct interface
func (processshape *ProcessShape) GetName() (res string) {
	return processshape.Name
}

// for satisfaction of GongStruct interface
func (processshape *ProcessShape) SetName(name string) {
	processshape.Name = name
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
	CreateORMControlFlow(ControlFlow *ControlFlow)
	CreateORMControlFlowShape(ControlFlowShape *ControlFlowShape)
	CreateORMDataFlow(DataFlow *DataFlow)
	CreateORMDataFlowShape(DataFlowShape *DataFlowShape)
	CreateORMDiagramProcess(DiagramProcess *DiagramProcess)
	CreateORMLibrary(Library *Library)
	CreateORMParticipant(Participant *Participant)
	CreateORMParticipantShape(ParticipantShape *ParticipantShape)
	CreateORMProcess(Process *Process)
	CreateORMProcessShape(ProcessShape *ProcessShape)
	CreateORMTask(Task *Task)
	CreateORMTaskShape(TaskShape *TaskShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMControlFlow(ControlFlow *ControlFlow)
	DeleteORMControlFlowShape(ControlFlowShape *ControlFlowShape)
	DeleteORMDataFlow(DataFlow *DataFlow)
	DeleteORMDataFlowShape(DataFlowShape *DataFlowShape)
	DeleteORMDiagramProcess(DiagramProcess *DiagramProcess)
	DeleteORMLibrary(Library *Library)
	DeleteORMParticipant(Participant *Participant)
	DeleteORMParticipantShape(ParticipantShape *ParticipantShape)
	DeleteORMProcess(Process *Process)
	DeleteORMProcessShape(ProcessShape *ProcessShape)
	DeleteORMTask(Task *Task)
	DeleteORMTaskShape(TaskShape *TaskShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.ControlFlows = make(map[*ControlFlow]struct{})
	stage.ControlFlows_mapString = make(map[string]*ControlFlow)
	stage.ControlFlow_stagedOrder = make(map[*ControlFlow]uint)
	stage.ControlFlowOrder = 0

	stage.ControlFlowShapes = make(map[*ControlFlowShape]struct{})
	stage.ControlFlowShapes_mapString = make(map[string]*ControlFlowShape)
	stage.ControlFlowShape_stagedOrder = make(map[*ControlFlowShape]uint)
	stage.ControlFlowShapeOrder = 0

	stage.DataFlows = make(map[*DataFlow]struct{})
	stage.DataFlows_mapString = make(map[string]*DataFlow)
	stage.DataFlow_stagedOrder = make(map[*DataFlow]uint)
	stage.DataFlowOrder = 0

	stage.DataFlowShapes = make(map[*DataFlowShape]struct{})
	stage.DataFlowShapes_mapString = make(map[string]*DataFlowShape)
	stage.DataFlowShape_stagedOrder = make(map[*DataFlowShape]uint)
	stage.DataFlowShapeOrder = 0

	stage.DiagramProcesss = make(map[*DiagramProcess]struct{})
	stage.DiagramProcesss_mapString = make(map[string]*DiagramProcess)
	stage.DiagramProcess_stagedOrder = make(map[*DiagramProcess]uint)
	stage.DiagramProcessOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Participants = make(map[*Participant]struct{})
	stage.Participants_mapString = make(map[string]*Participant)
	stage.Participant_stagedOrder = make(map[*Participant]uint)
	stage.ParticipantOrder = 0

	stage.ParticipantShapes = make(map[*ParticipantShape]struct{})
	stage.ParticipantShapes_mapString = make(map[string]*ParticipantShape)
	stage.ParticipantShape_stagedOrder = make(map[*ParticipantShape]uint)
	stage.ParticipantShapeOrder = 0

	stage.Processs = make(map[*Process]struct{})
	stage.Processs_mapString = make(map[string]*Process)
	stage.Process_stagedOrder = make(map[*Process]uint)
	stage.ProcessOrder = 0

	stage.ProcessShapes = make(map[*ProcessShape]struct{})
	stage.ProcessShapes_mapString = make(map[string]*ProcessShape)
	stage.ProcessShape_stagedOrder = make(map[*ProcessShape]uint)
	stage.ProcessShapeOrder = 0

	stage.Tasks = make(map[*Task]struct{})
	stage.Tasks_mapString = make(map[string]*Task)
	stage.Task_stagedOrder = make(map[*Task]uint)
	stage.TaskOrder = 0

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
	stage.ControlFlows = nil
	stage.ControlFlows_mapString = nil

	stage.ControlFlowShapes = nil
	stage.ControlFlowShapes_mapString = nil

	stage.DataFlows = nil
	stage.DataFlows_mapString = nil

	stage.DataFlowShapes = nil
	stage.DataFlowShapes_mapString = nil

	stage.DiagramProcesss = nil
	stage.DiagramProcesss_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Participants = nil
	stage.Participants_mapString = nil

	stage.ParticipantShapes = nil
	stage.ParticipantShapes_mapString = nil

	stage.Processs = nil
	stage.Processs_mapString = nil

	stage.ProcessShapes = nil
	stage.ProcessShapes_mapString = nil

	stage.Tasks = nil
	stage.Tasks_mapString = nil

	stage.TaskShapes = nil
	stage.TaskShapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for controlflow := range stage.ControlFlows {
		controlflow.Unstage(stage)
	}

	for controlflowshape := range stage.ControlFlowShapes {
		controlflowshape.Unstage(stage)
	}

	for dataflow := range stage.DataFlows {
		dataflow.Unstage(stage)
	}

	for dataflowshape := range stage.DataFlowShapes {
		dataflowshape.Unstage(stage)
	}

	for diagramprocess := range stage.DiagramProcesss {
		diagramprocess.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for participant := range stage.Participants {
		participant.Unstage(stage)
	}

	for participantshape := range stage.ParticipantShapes {
		participantshape.Unstage(stage)
	}

	for process := range stage.Processs {
		process.Unstage(stage)
	}

	for processshape := range stage.ProcessShapes {
		processshape.Unstage(stage)
	}

	for task := range stage.Tasks {
		task.Unstage(stage)
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
	case map[*ControlFlow]any:
		return any(&stage.ControlFlows).(*Type)
	case map[*ControlFlowShape]any:
		return any(&stage.ControlFlowShapes).(*Type)
	case map[*DataFlow]any:
		return any(&stage.DataFlows).(*Type)
	case map[*DataFlowShape]any:
		return any(&stage.DataFlowShapes).(*Type)
	case map[*DiagramProcess]any:
		return any(&stage.DiagramProcesss).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Participant]any:
		return any(&stage.Participants).(*Type)
	case map[*ParticipantShape]any:
		return any(&stage.ParticipantShapes).(*Type)
	case map[*Process]any:
		return any(&stage.Processs).(*Type)
	case map[*ProcessShape]any:
		return any(&stage.ProcessShapes).(*Type)
	case map[*Task]any:
		return any(&stage.Tasks).(*Type)
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
	case *ControlFlow:
		return any(stage.ControlFlows_mapString).(map[string]Type)
	case *ControlFlowShape:
		return any(stage.ControlFlowShapes_mapString).(map[string]Type)
	case *DataFlow:
		return any(stage.DataFlows_mapString).(map[string]Type)
	case *DataFlowShape:
		return any(stage.DataFlowShapes_mapString).(map[string]Type)
	case *DiagramProcess:
		return any(stage.DiagramProcesss_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Participant:
		return any(stage.Participants_mapString).(map[string]Type)
	case *ParticipantShape:
		return any(stage.ParticipantShapes_mapString).(map[string]Type)
	case *Process:
		return any(stage.Processs_mapString).(map[string]Type)
	case *ProcessShape:
		return any(stage.ProcessShapes_mapString).(map[string]Type)
	case *Task:
		return any(stage.Tasks_mapString).(map[string]Type)
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
	case ControlFlow:
		return any(&stage.ControlFlows).(*map[*Type]struct{})
	case ControlFlowShape:
		return any(&stage.ControlFlowShapes).(*map[*Type]struct{})
	case DataFlow:
		return any(&stage.DataFlows).(*map[*Type]struct{})
	case DataFlowShape:
		return any(&stage.DataFlowShapes).(*map[*Type]struct{})
	case DiagramProcess:
		return any(&stage.DiagramProcesss).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Participant:
		return any(&stage.Participants).(*map[*Type]struct{})
	case ParticipantShape:
		return any(&stage.ParticipantShapes).(*map[*Type]struct{})
	case Process:
		return any(&stage.Processs).(*map[*Type]struct{})
	case ProcessShape:
		return any(&stage.ProcessShapes).(*map[*Type]struct{})
	case Task:
		return any(&stage.Tasks).(*map[*Type]struct{})
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
	case *ControlFlow:
		return any(&stage.ControlFlows).(*map[Type]struct{})
	case *ControlFlowShape:
		return any(&stage.ControlFlowShapes).(*map[Type]struct{})
	case *DataFlow:
		return any(&stage.DataFlows).(*map[Type]struct{})
	case *DataFlowShape:
		return any(&stage.DataFlowShapes).(*map[Type]struct{})
	case *DiagramProcess:
		return any(&stage.DiagramProcesss).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Participant:
		return any(&stage.Participants).(*map[Type]struct{})
	case *ParticipantShape:
		return any(&stage.ParticipantShapes).(*map[Type]struct{})
	case *Process:
		return any(&stage.Processs).(*map[Type]struct{})
	case *ProcessShape:
		return any(&stage.ProcessShapes).(*map[Type]struct{})
	case *Task:
		return any(&stage.Tasks).(*map[Type]struct{})
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
	case ControlFlow:
		return any(&stage.ControlFlows_mapString).(*map[string]*Type)
	case ControlFlowShape:
		return any(&stage.ControlFlowShapes_mapString).(*map[string]*Type)
	case DataFlow:
		return any(&stage.DataFlows_mapString).(*map[string]*Type)
	case DataFlowShape:
		return any(&stage.DataFlowShapes_mapString).(*map[string]*Type)
	case DiagramProcess:
		return any(&stage.DiagramProcesss_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Participant:
		return any(&stage.Participants_mapString).(*map[string]*Type)
	case ParticipantShape:
		return any(&stage.ParticipantShapes_mapString).(*map[string]*Type)
	case Process:
		return any(&stage.Processs_mapString).(*map[string]*Type)
	case ProcessShape:
		return any(&stage.ProcessShapes_mapString).(*map[string]*Type)
	case Task:
		return any(&stage.Tasks_mapString).(*map[string]*Type)
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
	case ControlFlow:
		return any(&ControlFlow{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Start: &Task{Name: "Start"},
			// field is initialized with an instance of Task with the name of the field
			End: &Task{Name: "End"},
		}).(*Type)
	case ControlFlowShape:
		return any(&ControlFlowShape{
			// Initialisation of associations
			// field is initialized with an instance of ControlFlow with the name of the field
			ControlFlow: &ControlFlow{Name: "ControlFlow"},
		}).(*Type)
	case DataFlow:
		return any(&DataFlow{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Start: &Task{Name: "Start"},
			// field is initialized with an instance of Task with the name of the field
			End: &Task{Name: "End"},
		}).(*Type)
	case DataFlowShape:
		return any(&DataFlowShape{
			// Initialisation of associations
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlow: &DataFlow{Name: "DataFlow"},
		}).(*Type)
	case DiagramProcess:
		return any(&DiagramProcess{
			// Initialisation of associations
			// field is initialized with an instance of ProcessShape with the name of the field
			Process_Shapes: []*ProcessShape{{Name: "Process_Shapes"}},
			// field is initialized with an instance of Process with the name of the field
			ProcesssWhoseNodeIsExpanded: []*Process{{Name: "ProcesssWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ParticipantShape with the name of the field
			Participant_Shapes: []*ParticipantShape{{Name: "Participant_Shapes"}},
			// field is initialized with an instance of Participant with the name of the field
			ParticipantWhoseNodeIsExpanded: []*Participant{{Name: "ParticipantWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TasksWhoseNodeIsExpanded: []*Task{{Name: "TasksWhoseNodeIsExpanded"}},
			// field is initialized with an instance of TaskShape with the name of the field
			TaskShapes: []*TaskShape{{Name: "TaskShapes"}},
			// field is initialized with an instance of ControlFlow with the name of the field
			ControlFlowsWhoseNodeIsExpanded: []*ControlFlow{{Name: "ControlFlowsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ControlFlowShape with the name of the field
			ControlFlowShapes: []*ControlFlowShape{{Name: "ControlFlowShapes"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlowsWhoseNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataFlowShape with the name of the field
			DataFlowShapes: []*DataFlowShape{{Name: "DataFlowShapes"}},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Process with the name of the field
			RootProcesses: []*Process{{Name: "RootProcesses"}},
			// field is initialized with an instance of Process with the name of the field
			ProcesssWhoseNodeIsExpanded: []*Process{{Name: "ProcesssWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataFlow with the name of the field
			RootDataFlows: []*DataFlow{{Name: "RootDataFlows"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlowsWhoseNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseNodeIsExpanded"}},
		}).(*Type)
	case Participant:
		return any(&Participant{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Tasks: []*Task{{Name: "Tasks"}},
			// field is initialized with an instance of ControlFlow with the name of the field
			ControlFlows: []*ControlFlow{{Name: "ControlFlows"}},
			// field is initialized with an instance of Task with the name of the field
			TaskWhoseOutControlFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseOutControlFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TaskWhoseInControlFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseInControlFlowsNodeIsExpanded"}},
		}).(*Type)
	case ParticipantShape:
		return any(&ParticipantShape{
			// Initialisation of associations
			// field is initialized with an instance of Participant with the name of the field
			Participant: &Participant{Name: "Participant"},
		}).(*Type)
	case Process:
		return any(&Process{
			// Initialisation of associations
			// field is initialized with an instance of DiagramProcess with the name of the field
			DiagramProcesss: []*DiagramProcess{{Name: "DiagramProcesss"}},
			// field is initialized with an instance of DiagramProcess with the name of the field
			DiagramProcessWhoseNodeIsExpanded: []*DiagramProcess{{Name: "DiagramProcessWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Process with the name of the field
			SubProcesses: []*Process{{Name: "SubProcesses"}},
			// field is initialized with an instance of Participant with the name of the field
			Participants: []*Participant{{Name: "Participants"}},
			// field is initialized with an instance of Participant with the name of the field
			ParticipantWhoseNodeIsExpanded: []*Participant{{Name: "ParticipantWhoseNodeIsExpanded"}},
		}).(*Type)
	case ProcessShape:
		return any(&ProcessShape{
			// Initialisation of associations
			// field is initialized with an instance of Process with the name of the field
			Process: &Process{Name: "Process"},
		}).(*Type)
	case Task:
		return any(&Task{
			// Initialisation of associations
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
	// reverse maps of direct associations of ControlFlow
	case ControlFlow:
		switch fieldname {
		// insertion point for per direct association field
		case "Start":
			res := make(map[*Task][]*ControlFlow)
			for controlflow := range stage.ControlFlows {
				if controlflow.Start != nil {
					task_ := controlflow.Start
					var controlflows []*ControlFlow
					_, ok := res[task_]
					if ok {
						controlflows = res[task_]
					} else {
						controlflows = make([]*ControlFlow, 0)
					}
					controlflows = append(controlflows, controlflow)
					res[task_] = controlflows
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*Task][]*ControlFlow)
			for controlflow := range stage.ControlFlows {
				if controlflow.End != nil {
					task_ := controlflow.End
					var controlflows []*ControlFlow
					_, ok := res[task_]
					if ok {
						controlflows = res[task_]
					} else {
						controlflows = make([]*ControlFlow, 0)
					}
					controlflows = append(controlflows, controlflow)
					res[task_] = controlflows
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ControlFlowShape
	case ControlFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		case "ControlFlow":
			res := make(map[*ControlFlow][]*ControlFlowShape)
			for controlflowshape := range stage.ControlFlowShapes {
				if controlflowshape.ControlFlow != nil {
					controlflow_ := controlflowshape.ControlFlow
					var controlflowshapes []*ControlFlowShape
					_, ok := res[controlflow_]
					if ok {
						controlflowshapes = res[controlflow_]
					} else {
						controlflowshapes = make([]*ControlFlowShape, 0)
					}
					controlflowshapes = append(controlflowshapes, controlflowshape)
					res[controlflow_] = controlflowshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DataFlow
	case DataFlow:
		switch fieldname {
		// insertion point for per direct association field
		case "Start":
			res := make(map[*Task][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.Start != nil {
					task_ := dataflow.Start
					var dataflows []*DataFlow
					_, ok := res[task_]
					if ok {
						dataflows = res[task_]
					} else {
						dataflows = make([]*DataFlow, 0)
					}
					dataflows = append(dataflows, dataflow)
					res[task_] = dataflows
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*Task][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.End != nil {
					task_ := dataflow.End
					var dataflows []*DataFlow
					_, ok := res[task_]
					if ok {
						dataflows = res[task_]
					} else {
						dataflows = make([]*DataFlow, 0)
					}
					dataflows = append(dataflows, dataflow)
					res[task_] = dataflows
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DataFlowShape
	case DataFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		case "DataFlow":
			res := make(map[*DataFlow][]*DataFlowShape)
			for dataflowshape := range stage.DataFlowShapes {
				if dataflowshape.DataFlow != nil {
					dataflow_ := dataflowshape.DataFlow
					var dataflowshapes []*DataFlowShape
					_, ok := res[dataflow_]
					if ok {
						dataflowshapes = res[dataflow_]
					} else {
						dataflowshapes = make([]*DataFlowShape, 0)
					}
					dataflowshapes = append(dataflowshapes, dataflowshape)
					res[dataflow_] = dataflowshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DiagramProcess
	case DiagramProcess:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Participant
	case Participant:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ParticipantShape
	case ParticipantShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Participant":
			res := make(map[*Participant][]*ParticipantShape)
			for participantshape := range stage.ParticipantShapes {
				if participantshape.Participant != nil {
					participant_ := participantshape.Participant
					var participantshapes []*ParticipantShape
					_, ok := res[participant_]
					if ok {
						participantshapes = res[participant_]
					} else {
						participantshapes = make([]*ParticipantShape, 0)
					}
					participantshapes = append(participantshapes, participantshape)
					res[participant_] = participantshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Process
	case Process:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ProcessShape
	case ProcessShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Process":
			res := make(map[*Process][]*ProcessShape)
			for processshape := range stage.ProcessShapes {
				if processshape.Process != nil {
					process_ := processshape.Process
					var processshapes []*ProcessShape
					_, ok := res[process_]
					if ok {
						processshapes = res[process_]
					} else {
						processshapes = make([]*ProcessShape, 0)
					}
					processshapes = append(processshapes, processshape)
					res[process_] = processshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of ControlFlow
	case ControlFlow:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ControlFlowShape
	case ControlFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DataFlow
	case DataFlow:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DataFlowShape
	case DataFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DiagramProcess
	case DiagramProcess:
		switch fieldname {
		// insertion point for per direct association field
		case "Process_Shapes":
			res := make(map[*ProcessShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, processshape_ := range diagramprocess.Process_Shapes {
					res[processshape_] = append(res[processshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProcesssWhoseNodeIsExpanded":
			res := make(map[*Process][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, process_ := range diagramprocess.ProcesssWhoseNodeIsExpanded {
					res[process_] = append(res[process_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Participant_Shapes":
			res := make(map[*ParticipantShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, participantshape_ := range diagramprocess.Participant_Shapes {
					res[participantshape_] = append(res[participantshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParticipantWhoseNodeIsExpanded":
			res := make(map[*Participant][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, participant_ := range diagramprocess.ParticipantWhoseNodeIsExpanded {
					res[participant_] = append(res[participant_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TasksWhoseNodeIsExpanded":
			res := make(map[*Task][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, task_ := range diagramprocess.TasksWhoseNodeIsExpanded {
					res[task_] = append(res[task_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskShapes":
			res := make(map[*TaskShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, taskshape_ := range diagramprocess.TaskShapes {
					res[taskshape_] = append(res[taskshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlFlowsWhoseNodeIsExpanded":
			res := make(map[*ControlFlow][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, controlflow_ := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
					res[controlflow_] = append(res[controlflow_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlFlowShapes":
			res := make(map[*ControlFlowShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, controlflowshape_ := range diagramprocess.ControlFlowShapes {
					res[controlflowshape_] = append(res[controlflowshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlowsWhoseNodeIsExpanded":
			res := make(map[*DataFlow][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, dataflow_ := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
					res[dataflow_] = append(res[dataflow_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlowShapes":
			res := make(map[*DataFlowShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, dataflowshape_ := range diagramprocess.DataFlowShapes {
					res[dataflowshape_] = append(res[dataflowshape_], diagramprocess)
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
		case "RootProcesses":
			res := make(map[*Process][]*Library)
			for library := range stage.Librarys {
				for _, process_ := range library.RootProcesses {
					res[process_] = append(res[process_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ProcesssWhoseNodeIsExpanded":
			res := make(map[*Process][]*Library)
			for library := range stage.Librarys {
				for _, process_ := range library.ProcesssWhoseNodeIsExpanded {
					res[process_] = append(res[process_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootDataFlows":
			res := make(map[*DataFlow][]*Library)
			for library := range stage.Librarys {
				for _, dataflow_ := range library.RootDataFlows {
					res[dataflow_] = append(res[dataflow_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlowsWhoseNodeIsExpanded":
			res := make(map[*DataFlow][]*Library)
			for library := range stage.Librarys {
				for _, dataflow_ := range library.DataFlowsWhoseNodeIsExpanded {
					res[dataflow_] = append(res[dataflow_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Participant
	case Participant:
		switch fieldname {
		// insertion point for per direct association field
		case "Tasks":
			res := make(map[*Task][]*Participant)
			for participant := range stage.Participants {
				for _, task_ := range participant.Tasks {
					res[task_] = append(res[task_], participant)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ControlFlows":
			res := make(map[*ControlFlow][]*Participant)
			for participant := range stage.Participants {
				for _, controlflow_ := range participant.ControlFlows {
					res[controlflow_] = append(res[controlflow_], participant)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskWhoseOutControlFlowsNodeIsExpanded":
			res := make(map[*Task][]*Participant)
			for participant := range stage.Participants {
				for _, task_ := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
					res[task_] = append(res[task_], participant)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskWhoseInControlFlowsNodeIsExpanded":
			res := make(map[*Task][]*Participant)
			for participant := range stage.Participants {
				for _, task_ := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
					res[task_] = append(res[task_], participant)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ParticipantShape
	case ParticipantShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Process
	case Process:
		switch fieldname {
		// insertion point for per direct association field
		case "DiagramProcesss":
			res := make(map[*DiagramProcess][]*Process)
			for process := range stage.Processs {
				for _, diagramprocess_ := range process.DiagramProcesss {
					res[diagramprocess_] = append(res[diagramprocess_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DiagramProcessWhoseNodeIsExpanded":
			res := make(map[*DiagramProcess][]*Process)
			for process := range stage.Processs {
				for _, diagramprocess_ := range process.DiagramProcessWhoseNodeIsExpanded {
					res[diagramprocess_] = append(res[diagramprocess_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SubProcesses":
			res := make(map[*Process][]*Process)
			for process := range stage.Processs {
				for _, process_ := range process.SubProcesses {
					res[process_] = append(res[process_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Participants":
			res := make(map[*Participant][]*Process)
			for process := range stage.Processs {
				for _, participant_ := range process.Participants {
					res[participant_] = append(res[participant_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParticipantWhoseNodeIsExpanded":
			res := make(map[*Participant][]*Process)
			for process := range stage.Processs {
				for _, participant_ := range process.ParticipantWhoseNodeIsExpanded {
					res[participant_] = append(res[participant_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ProcessShape
	case ProcessShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Task
	case Task:
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
	case *ControlFlow:
		res = "ControlFlow"
	case *ControlFlowShape:
		res = "ControlFlowShape"
	case *DataFlow:
		res = "DataFlow"
	case *DataFlowShape:
		res = "DataFlowShape"
	case *DiagramProcess:
		res = "DiagramProcess"
	case *Library:
		res = "Library"
	case *Participant:
		res = "Participant"
	case *ParticipantShape:
		res = "ParticipantShape"
	case *Process:
		res = "Process"
	case *ProcessShape:
		res = "ProcessShape"
	case *Task:
		res = "Task"
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
	case *ControlFlow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ControlFlowsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Participant"
		rf.Fieldname = "ControlFlows"
		res = append(res, rf)
	case *ControlFlowShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ControlFlowShapes"
		res = append(res, rf)
	case *DataFlow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "DataFlowsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootDataFlows"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "DataFlowsWhoseNodeIsExpanded"
		res = append(res, rf)
	case *DataFlowShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "DataFlowShapes"
		res = append(res, rf)
	case *DiagramProcess:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Process"
		rf.Fieldname = "DiagramProcesss"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "DiagramProcessWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
	case *Participant:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ParticipantWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "Participants"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "ParticipantWhoseNodeIsExpanded"
		res = append(res, rf)
	case *ParticipantShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "Participant_Shapes"
		res = append(res, rf)
	case *Process:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootProcesses"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "SubProcesses"
		res = append(res, rf)
	case *ProcessShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "Process_Shapes"
		res = append(res, rf)
	case *Task:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "TasksWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Participant"
		rf.Fieldname = "Tasks"
		res = append(res, rf)
		rf.GongstructName = "Participant"
		rf.Fieldname = "TaskWhoseOutControlFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Participant"
		rf.Fieldname = "TaskWhoseInControlFlowsNodeIsExpanded"
		res = append(res, rf)
	case *TaskShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "TaskShapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (controlflow *ControlFlow) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Start",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "End",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
	}
	return
}

func (controlflowshape *ControlFlowShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ControlFlow",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "ControlFlow",
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

func (dataflow *DataFlow) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Start",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "End",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
	}
	return
}

func (dataflowshape *DataFlowShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "DataFlow",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DataFlow",
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

func (diagramprocess *DiagramProcess) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Process_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ProcessShape",
		},
		{
			Name:               "IsProcesssNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ProcesssWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:                 "Participant_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ParticipantShape",
		},
		{
			Name:               "IsParticipantsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ParticipantWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Participant",
		},
		{
			Name:                 "TasksWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "TaskShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskShape",
		},
		{
			Name:                 "ControlFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlFlow",
		},
		{
			Name:                 "ControlFlowShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlFlowShape",
		},
		{
			Name:                 "DataFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:                 "DataFlowShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlowShape",
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
			Name:               "ComputedPrefix",
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
			Name:                 "RootProcesses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:                 "ProcesssWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:                 "RootDataFlows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:                 "DataFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:               "IsExpandedTmp",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (participant *Participant) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "IsTasksNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Tasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:               "IsControlFlowsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ControlFlows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlFlow",
		},
		{
			Name:                 "TaskWhoseOutControlFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "TaskWhoseInControlFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
	}
	return
}

func (participantshape *ParticipantShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Participant",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Participant",
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

func (process *Process) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "DiagramProcesss",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramProcess",
		},
		{
			Name:                 "DiagramProcessWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DiagramProcess",
		},
		{
			Name:               "IsSubProcessNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubProcesses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:                 "Participants",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Participant",
		},
		{
			Name:                 "ParticipantWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Participant",
		},
	}
	return
}

func (processshape *ProcessShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Process",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Process",
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
			Name:               "IsStartTask",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEndTask",
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
func (controlflow *ControlFlow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlflow.Name
	case "ComputedPrefix":
		res.valueString = controlflow.ComputedPrefix
	case "Start":
		res.GongFieldValueType = GongFieldValueTypePointer
		if controlflow.Start != nil {
			res.valueString = controlflow.Start.Name
			res.ids = controlflow.Start.GongGetUUID(stage)
		}
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if controlflow.End != nil {
			res.valueString = controlflow.End.Name
			res.ids = controlflow.End.GongGetUUID(stage)
		}
	}
	return
}

func (controlflowshape *ControlFlowShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlflowshape.Name
	case "ControlFlow":
		res.GongFieldValueType = GongFieldValueTypePointer
		if controlflowshape.ControlFlow != nil {
			res.valueString = controlflowshape.ControlFlow.Name
			res.ids = controlflowshape.ControlFlow.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", controlflowshape.StartRatio)
		res.valueFloat = controlflowshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", controlflowshape.EndRatio)
		res.valueFloat = controlflowshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := controlflowshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := controlflowshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", controlflowshape.CornerOffsetRatio)
		res.valueFloat = controlflowshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", controlflowshape.IsHidden)
		res.valueBool = controlflowshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (dataflow *DataFlow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dataflow.Name
	case "ComputedPrefix":
		res.valueString = dataflow.ComputedPrefix
	case "Start":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.Start != nil {
			res.valueString = dataflow.Start.Name
			res.ids = dataflow.Start.GongGetUUID(stage)
		}
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.End != nil {
			res.valueString = dataflow.End.Name
			res.ids = dataflow.End.GongGetUUID(stage)
		}
	}
	return
}

func (dataflowshape *DataFlowShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dataflowshape.Name
	case "DataFlow":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflowshape.DataFlow != nil {
			res.valueString = dataflowshape.DataFlow.Name
			res.ids = dataflowshape.DataFlow.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", dataflowshape.StartRatio)
		res.valueFloat = dataflowshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", dataflowshape.EndRatio)
		res.valueFloat = dataflowshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := dataflowshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := dataflowshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", dataflowshape.CornerOffsetRatio)
		res.valueFloat = dataflowshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", dataflowshape.IsHidden)
		res.valueBool = dataflowshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (diagramprocess *DiagramProcess) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramprocess.Name
	case "ComputedPrefix":
		res.valueString = diagramprocess.ComputedPrefix
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsChecked)
		res.valueBool = diagramprocess.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsEditable_)
		res.valueBool = diagramprocess.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsShowPrefix":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsShowPrefix)
		res.valueBool = diagramprocess.IsShowPrefix
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DefaultBoxWidth":
		res.valueString = fmt.Sprintf("%f", diagramprocess.DefaultBoxWidth)
		res.valueFloat = diagramprocess.DefaultBoxWidth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "DefaultBoxHeigth":
		res.valueString = fmt.Sprintf("%f", diagramprocess.DefaultBoxHeigth)
		res.valueFloat = diagramprocess.DefaultBoxHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", diagramprocess.Width)
		res.valueFloat = diagramprocess.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", diagramprocess.Height)
		res.valueFloat = diagramprocess.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Process_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.Process_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsProcesssNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsProcesssNodeExpanded)
		res.valueBool = diagramprocess.IsProcesssNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ProcesssWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ProcesssWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Participant_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.Participant_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsParticipantsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsParticipantsNodeExpanded)
		res.valueBool = diagramprocess.IsParticipantsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ParticipantWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ParticipantWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TasksWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.TasksWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TaskShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.TaskShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ControlFlowsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ControlFlowShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ControlFlowShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DataFlowsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DataFlowShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.DataFlowShapes {
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
	case "ComputedPrefix":
		res.valueString = library.ComputedPrefix
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
	case "RootProcesses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootProcesses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ProcesssWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.ProcesssWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "RootDataFlows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootDataFlows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DataFlowsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.DataFlowsWhoseNodeIsExpanded {
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

func (participant *Participant) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = participant.Name
	case "ComputedPrefix":
		res.valueString = participant.ComputedPrefix
	case "IsTasksNodeExpanded":
		res.valueString = fmt.Sprintf("%t", participant.IsTasksNodeExpanded)
		res.valueBool = participant.IsTasksNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Tasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range participant.Tasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsControlFlowsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", participant.IsControlFlowsNodeExpanded)
		res.valueBool = participant.IsControlFlowsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ControlFlows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range participant.ControlFlows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TaskWhoseOutControlFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TaskWhoseInControlFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
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

func (participantshape *ParticipantShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = participantshape.Name
	case "Participant":
		res.GongFieldValueType = GongFieldValueTypePointer
		if participantshape.Participant != nil {
			res.valueString = participantshape.Participant.Name
			res.ids = participantshape.Participant.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", participantshape.IsExpanded)
		res.valueBool = participantshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", participantshape.X)
		res.valueFloat = participantshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", participantshape.Y)
		res.valueFloat = participantshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", participantshape.Width)
		res.valueFloat = participantshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", participantshape.Height)
		res.valueFloat = participantshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", participantshape.IsHidden)
		res.valueBool = participantshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (process *Process) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = process.Name
	case "ComputedPrefix":
		res.valueString = process.ComputedPrefix
	case "DiagramProcesss":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.DiagramProcesss {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DiagramProcessWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.DiagramProcessWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsSubProcessNodeExpanded":
		res.valueString = fmt.Sprintf("%t", process.IsSubProcessNodeExpanded)
		res.valueBool = process.IsSubProcessNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubProcesses":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.SubProcesses {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Participants":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.Participants {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ParticipantWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.ParticipantWhoseNodeIsExpanded {
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

func (processshape *ProcessShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = processshape.Name
	case "Process":
		res.GongFieldValueType = GongFieldValueTypePointer
		if processshape.Process != nil {
			res.valueString = processshape.Process.Name
			res.ids = processshape.Process.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", processshape.IsExpanded)
		res.valueBool = processshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", processshape.X)
		res.valueFloat = processshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", processshape.Y)
		res.valueFloat = processshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", processshape.Width)
		res.valueFloat = processshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", processshape.Height)
		res.valueFloat = processshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", processshape.IsHidden)
		res.valueBool = processshape.IsHidden
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
	case "IsStartTask":
		res.valueString = fmt.Sprintf("%t", task.IsStartTask)
		res.valueBool = task.IsStartTask
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEndTask":
		res.valueString = fmt.Sprintf("%t", task.IsEndTask)
		res.valueBool = task.IsEndTask
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
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", taskshape.IsExpanded)
		res.valueBool = taskshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
func (controlflow *ControlFlow) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		controlflow.Name = value.GetValueString()
	case "ComputedPrefix":
		controlflow.ComputedPrefix = value.GetValueString()
	case "Start":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			controlflow.Start = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					controlflow.Start = __instance__
					break
				}
			}
		}
	case "End":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			controlflow.End = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					controlflow.End = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (controlflowshape *ControlFlowShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		controlflowshape.Name = value.GetValueString()
	case "ControlFlow":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			controlflowshape.ControlFlow = nil
			for __instance__ := range stage.ControlFlows {
				if stage.ControlFlow_stagedOrder[__instance__] == uint(id) {
					controlflowshape.ControlFlow = __instance__
					break
				}
			}
		}
	case "StartRatio":
		controlflowshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		controlflowshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		controlflowshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		controlflowshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		controlflowshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		controlflowshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (dataflow *DataFlow) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		dataflow.Name = value.GetValueString()
	case "ComputedPrefix":
		dataflow.ComputedPrefix = value.GetValueString()
	case "Start":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.Start = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					dataflow.Start = __instance__
					break
				}
			}
		}
	case "End":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.End = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					dataflow.End = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (dataflowshape *DataFlowShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		dataflowshape.Name = value.GetValueString()
	case "DataFlow":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflowshape.DataFlow = nil
			for __instance__ := range stage.DataFlows {
				if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
					dataflowshape.DataFlow = __instance__
					break
				}
			}
		}
	case "StartRatio":
		dataflowshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		dataflowshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		dataflowshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		dataflowshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		dataflowshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		dataflowshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (diagramprocess *DiagramProcess) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		diagramprocess.Name = value.GetValueString()
	case "ComputedPrefix":
		diagramprocess.ComputedPrefix = value.GetValueString()
	case "IsChecked":
		diagramprocess.IsChecked = value.GetValueBool()
	case "IsEditable_":
		diagramprocess.IsEditable_ = value.GetValueBool()
	case "IsShowPrefix":
		diagramprocess.IsShowPrefix = value.GetValueBool()
	case "DefaultBoxWidth":
		diagramprocess.DefaultBoxWidth = value.GetValueFloat()
	case "DefaultBoxHeigth":
		diagramprocess.DefaultBoxHeigth = value.GetValueFloat()
	case "Width":
		diagramprocess.Width = value.GetValueFloat()
	case "Height":
		diagramprocess.Height = value.GetValueFloat()
	case "Process_Shapes":
		diagramprocess.Process_Shapes = make([]*ProcessShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ProcessShapes {
					if stage.ProcessShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.Process_Shapes = append(diagramprocess.Process_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsProcesssNodeExpanded":
		diagramprocess.IsProcesssNodeExpanded = value.GetValueBool()
	case "ProcesssWhoseNodeIsExpanded":
		diagramprocess.ProcesssWhoseNodeIsExpanded = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ProcesssWhoseNodeIsExpanded = append(diagramprocess.ProcesssWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Participant_Shapes":
		diagramprocess.Participant_Shapes = make([]*ParticipantShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ParticipantShapes {
					if stage.ParticipantShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.Participant_Shapes = append(diagramprocess.Participant_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsParticipantsNodeExpanded":
		diagramprocess.IsParticipantsNodeExpanded = value.GetValueBool()
	case "ParticipantWhoseNodeIsExpanded":
		diagramprocess.ParticipantWhoseNodeIsExpanded = make([]*Participant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Participants {
					if stage.Participant_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ParticipantWhoseNodeIsExpanded = append(diagramprocess.ParticipantWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "TasksWhoseNodeIsExpanded":
		diagramprocess.TasksWhoseNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						diagramprocess.TasksWhoseNodeIsExpanded = append(diagramprocess.TasksWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "TaskShapes":
		diagramprocess.TaskShapes = make([]*TaskShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskShapes {
					if stage.TaskShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.TaskShapes = append(diagramprocess.TaskShapes, __instance__)
						break
					}
				}
			}
		}
	case "ControlFlowsWhoseNodeIsExpanded":
		diagramprocess.ControlFlowsWhoseNodeIsExpanded = make([]*ControlFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlFlows {
					if stage.ControlFlow_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ControlFlowsWhoseNodeIsExpanded = append(diagramprocess.ControlFlowsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ControlFlowShapes":
		diagramprocess.ControlFlowShapes = make([]*ControlFlowShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlFlowShapes {
					if stage.ControlFlowShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ControlFlowShapes = append(diagramprocess.ControlFlowShapes, __instance__)
						break
					}
				}
			}
		}
	case "DataFlowsWhoseNodeIsExpanded":
		diagramprocess.DataFlowsWhoseNodeIsExpanded = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						diagramprocess.DataFlowsWhoseNodeIsExpanded = append(diagramprocess.DataFlowsWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "DataFlowShapes":
		diagramprocess.DataFlowShapes = make([]*DataFlowShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlowShapes {
					if stage.DataFlowShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.DataFlowShapes = append(diagramprocess.DataFlowShapes, __instance__)
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
	case "ComputedPrefix":
		library.ComputedPrefix = value.GetValueString()
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
	case "RootProcesses":
		library.RootProcesses = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						library.RootProcesses = append(library.RootProcesses, __instance__)
						break
					}
				}
			}
		}
	case "ProcesssWhoseNodeIsExpanded":
		library.ProcesssWhoseNodeIsExpanded = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						library.ProcesssWhoseNodeIsExpanded = append(library.ProcesssWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "RootDataFlows":
		library.RootDataFlows = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						library.RootDataFlows = append(library.RootDataFlows, __instance__)
						break
					}
				}
			}
		}
	case "DataFlowsWhoseNodeIsExpanded":
		library.DataFlowsWhoseNodeIsExpanded = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						library.DataFlowsWhoseNodeIsExpanded = append(library.DataFlowsWhoseNodeIsExpanded, __instance__)
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

func (participant *Participant) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		participant.Name = value.GetValueString()
	case "ComputedPrefix":
		participant.ComputedPrefix = value.GetValueString()
	case "IsTasksNodeExpanded":
		participant.IsTasksNodeExpanded = value.GetValueBool()
	case "Tasks":
		participant.Tasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						participant.Tasks = append(participant.Tasks, __instance__)
						break
					}
				}
			}
		}
	case "IsControlFlowsNodeExpanded":
		participant.IsControlFlowsNodeExpanded = value.GetValueBool()
	case "ControlFlows":
		participant.ControlFlows = make([]*ControlFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlFlows {
					if stage.ControlFlow_stagedOrder[__instance__] == uint(id) {
						participant.ControlFlows = append(participant.ControlFlows, __instance__)
						break
					}
				}
			}
		}
	case "TaskWhoseOutControlFlowsNodeIsExpanded":
		participant.TaskWhoseOutControlFlowsNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						participant.TaskWhoseOutControlFlowsNodeIsExpanded = append(participant.TaskWhoseOutControlFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "TaskWhoseInControlFlowsNodeIsExpanded":
		participant.TaskWhoseInControlFlowsNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						participant.TaskWhoseInControlFlowsNodeIsExpanded = append(participant.TaskWhoseInControlFlowsNodeIsExpanded, __instance__)
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

func (participantshape *ParticipantShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		participantshape.Name = value.GetValueString()
	case "Participant":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			participantshape.Participant = nil
			for __instance__ := range stage.Participants {
				if stage.Participant_stagedOrder[__instance__] == uint(id) {
					participantshape.Participant = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		participantshape.IsExpanded = value.GetValueBool()
	case "X":
		participantshape.X = value.GetValueFloat()
	case "Y":
		participantshape.Y = value.GetValueFloat()
	case "Width":
		participantshape.Width = value.GetValueFloat()
	case "Height":
		participantshape.Height = value.GetValueFloat()
	case "IsHidden":
		participantshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (process *Process) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		process.Name = value.GetValueString()
	case "ComputedPrefix":
		process.ComputedPrefix = value.GetValueString()
	case "DiagramProcesss":
		process.DiagramProcesss = make([]*DiagramProcess, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramProcesss {
					if stage.DiagramProcess_stagedOrder[__instance__] == uint(id) {
						process.DiagramProcesss = append(process.DiagramProcesss, __instance__)
						break
					}
				}
			}
		}
	case "DiagramProcessWhoseNodeIsExpanded":
		process.DiagramProcessWhoseNodeIsExpanded = make([]*DiagramProcess, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DiagramProcesss {
					if stage.DiagramProcess_stagedOrder[__instance__] == uint(id) {
						process.DiagramProcessWhoseNodeIsExpanded = append(process.DiagramProcessWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsSubProcessNodeExpanded":
		process.IsSubProcessNodeExpanded = value.GetValueBool()
	case "SubProcesses":
		process.SubProcesses = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						process.SubProcesses = append(process.SubProcesses, __instance__)
						break
					}
				}
			}
		}
	case "Participants":
		process.Participants = make([]*Participant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Participants {
					if stage.Participant_stagedOrder[__instance__] == uint(id) {
						process.Participants = append(process.Participants, __instance__)
						break
					}
				}
			}
		}
	case "ParticipantWhoseNodeIsExpanded":
		process.ParticipantWhoseNodeIsExpanded = make([]*Participant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Participants {
					if stage.Participant_stagedOrder[__instance__] == uint(id) {
						process.ParticipantWhoseNodeIsExpanded = append(process.ParticipantWhoseNodeIsExpanded, __instance__)
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

func (processshape *ProcessShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		processshape.Name = value.GetValueString()
	case "Process":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			processshape.Process = nil
			for __instance__ := range stage.Processs {
				if stage.Process_stagedOrder[__instance__] == uint(id) {
					processshape.Process = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		processshape.IsExpanded = value.GetValueBool()
	case "X":
		processshape.X = value.GetValueFloat()
	case "Y":
		processshape.Y = value.GetValueFloat()
	case "Width":
		processshape.Width = value.GetValueFloat()
	case "Height":
		processshape.Height = value.GetValueFloat()
	case "IsHidden":
		processshape.IsHidden = value.GetValueBool()
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
	case "IsStartTask":
		task.IsStartTask = value.GetValueBool()
	case "IsEndTask":
		task.IsEndTask = value.GetValueBool()
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
	case "IsExpanded":
		taskshape.IsExpanded = value.GetValueBool()
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
func (controlflow *ControlFlow) GongGetGongstructName() string {
	return "ControlFlow"
}

func (controlflowshape *ControlFlowShape) GongGetGongstructName() string {
	return "ControlFlowShape"
}

func (dataflow *DataFlow) GongGetGongstructName() string {
	return "DataFlow"
}

func (dataflowshape *DataFlowShape) GongGetGongstructName() string {
	return "DataFlowShape"
}

func (diagramprocess *DiagramProcess) GongGetGongstructName() string {
	return "DiagramProcess"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (participant *Participant) GongGetGongstructName() string {
	return "Participant"
}

func (participantshape *ParticipantShape) GongGetGongstructName() string {
	return "ParticipantShape"
}

func (process *Process) GongGetGongstructName() string {
	return "Process"
}

func (processshape *ProcessShape) GongGetGongstructName() string {
	return "ProcessShape"
}

func (task *Task) GongGetGongstructName() string {
	return "Task"
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
	stage.ControlFlows_mapString = make(map[string]*ControlFlow)
	for controlflow := range stage.ControlFlows {
		stage.ControlFlows_mapString[controlflow.Name] = controlflow
	}

	stage.ControlFlowShapes_mapString = make(map[string]*ControlFlowShape)
	for controlflowshape := range stage.ControlFlowShapes {
		stage.ControlFlowShapes_mapString[controlflowshape.Name] = controlflowshape
	}

	stage.DataFlows_mapString = make(map[string]*DataFlow)
	for dataflow := range stage.DataFlows {
		stage.DataFlows_mapString[dataflow.Name] = dataflow
	}

	stage.DataFlowShapes_mapString = make(map[string]*DataFlowShape)
	for dataflowshape := range stage.DataFlowShapes {
		stage.DataFlowShapes_mapString[dataflowshape.Name] = dataflowshape
	}

	stage.DiagramProcesss_mapString = make(map[string]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		stage.DiagramProcesss_mapString[diagramprocess.Name] = diagramprocess
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Participants_mapString = make(map[string]*Participant)
	for participant := range stage.Participants {
		stage.Participants_mapString[participant.Name] = participant
	}

	stage.ParticipantShapes_mapString = make(map[string]*ParticipantShape)
	for participantshape := range stage.ParticipantShapes {
		stage.ParticipantShapes_mapString[participantshape.Name] = participantshape
	}

	stage.Processs_mapString = make(map[string]*Process)
	for process := range stage.Processs {
		stage.Processs_mapString[process.Name] = process
	}

	stage.ProcessShapes_mapString = make(map[string]*ProcessShape)
	for processshape := range stage.ProcessShapes {
		stage.ProcessShapes_mapString[processshape.Name] = processshape
	}

	stage.Tasks_mapString = make(map[string]*Task)
	for task := range stage.Tasks {
		stage.Tasks_mapString[task.Name] = task
	}

	stage.TaskShapes_mapString = make(map[string]*TaskShape)
	for taskshape := range stage.TaskShapes {
		stage.TaskShapes_mapString[taskshape.Name] = taskshape
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
