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
	AllocatedProcessShapes                map[*AllocatedProcessShape]struct{}
	AllocatedProcessShapes_instance       map[*AllocatedProcessShape]*AllocatedProcessShape
	AllocatedProcessShapes_mapString      map[string]*AllocatedProcessShape
	AllocatedProcessShapeOrder            uint
	AllocatedProcessShape_stagedOrder     map[*AllocatedProcessShape]uint
	AllocatedProcessShape_orderStaged     map[uint]*AllocatedProcessShape
	AllocatedProcessShapes_reference      map[*AllocatedProcessShape]*AllocatedProcessShape
	AllocatedProcessShapes_referenceOrder map[*AllocatedProcessShape]uint

	// insertion point for slice of pointers maps
	OnAfterAllocatedProcessShapeCreateCallback GongOnAfterCreateInterface[AllocatedProcessShape]
	OnAfterAllocatedProcessShapeUpdateCallback GongOnAfterUpdateInterface[AllocatedProcessShape]
	OnAfterAllocatedProcessShapeDeleteCallback GongOnAfterDeleteInterface[AllocatedProcessShape]

	AllocatedResourceShapes                map[*AllocatedResourceShape]struct{}
	AllocatedResourceShapes_instance       map[*AllocatedResourceShape]*AllocatedResourceShape
	AllocatedResourceShapes_mapString      map[string]*AllocatedResourceShape
	AllocatedResourceShapeOrder            uint
	AllocatedResourceShape_stagedOrder     map[*AllocatedResourceShape]uint
	AllocatedResourceShape_orderStaged     map[uint]*AllocatedResourceShape
	AllocatedResourceShapes_reference      map[*AllocatedResourceShape]*AllocatedResourceShape
	AllocatedResourceShapes_referenceOrder map[*AllocatedResourceShape]uint

	// insertion point for slice of pointers maps
	OnAfterAllocatedResourceShapeCreateCallback GongOnAfterCreateInterface[AllocatedResourceShape]
	OnAfterAllocatedResourceShapeUpdateCallback GongOnAfterUpdateInterface[AllocatedResourceShape]
	OnAfterAllocatedResourceShapeDeleteCallback GongOnAfterDeleteInterface[AllocatedResourceShape]

	ControlFlows                map[*ControlFlow]struct{}
	ControlFlows_instance       map[*ControlFlow]*ControlFlow
	ControlFlows_mapString      map[string]*ControlFlow
	ControlFlowOrder            uint
	ControlFlow_stagedOrder     map[*ControlFlow]uint
	ControlFlow_orderStaged     map[uint]*ControlFlow
	ControlFlows_reference      map[*ControlFlow]*ControlFlow
	ControlFlows_referenceOrder map[*ControlFlow]uint

	// insertion point for slice of pointers maps
	OnAfterControlFlowCreateCallback GongOnAfterCreateInterface[ControlFlow]
	OnAfterControlFlowUpdateCallback GongOnAfterUpdateInterface[ControlFlow]
	OnAfterControlFlowDeleteCallback GongOnAfterDeleteInterface[ControlFlow]

	ControlFlowShapes                map[*ControlFlowShape]struct{}
	ControlFlowShapes_instance       map[*ControlFlowShape]*ControlFlowShape
	ControlFlowShapes_mapString      map[string]*ControlFlowShape
	ControlFlowShapeOrder            uint
	ControlFlowShape_stagedOrder     map[*ControlFlowShape]uint
	ControlFlowShape_orderStaged     map[uint]*ControlFlowShape
	ControlFlowShapes_reference      map[*ControlFlowShape]*ControlFlowShape
	ControlFlowShapes_referenceOrder map[*ControlFlowShape]uint

	// insertion point for slice of pointers maps
	OnAfterControlFlowShapeCreateCallback GongOnAfterCreateInterface[ControlFlowShape]
	OnAfterControlFlowShapeUpdateCallback GongOnAfterUpdateInterface[ControlFlowShape]
	OnAfterControlFlowShapeDeleteCallback GongOnAfterDeleteInterface[ControlFlowShape]

	Datas                map[*Data]struct{}
	Datas_instance       map[*Data]*Data
	Datas_mapString      map[string]*Data
	DataOrder            uint
	Data_stagedOrder     map[*Data]uint
	Data_orderStaged     map[uint]*Data
	Datas_reference      map[*Data]*Data
	Datas_referenceOrder map[*Data]uint

	// insertion point for slice of pointers maps
	OnAfterDataCreateCallback GongOnAfterCreateInterface[Data]
	OnAfterDataUpdateCallback GongOnAfterUpdateInterface[Data]
	OnAfterDataDeleteCallback GongOnAfterDeleteInterface[Data]

	DataFlows                map[*DataFlow]struct{}
	DataFlows_instance       map[*DataFlow]*DataFlow
	DataFlows_mapString      map[string]*DataFlow
	DataFlowOrder            uint
	DataFlow_stagedOrder     map[*DataFlow]uint
	DataFlow_orderStaged     map[uint]*DataFlow
	DataFlows_reference      map[*DataFlow]*DataFlow
	DataFlows_referenceOrder map[*DataFlow]uint

	// insertion point for slice of pointers maps
	DataFlow_Datas_reverseMap map[*Data]*DataFlow

	OnAfterDataFlowCreateCallback GongOnAfterCreateInterface[DataFlow]
	OnAfterDataFlowUpdateCallback GongOnAfterUpdateInterface[DataFlow]
	OnAfterDataFlowDeleteCallback GongOnAfterDeleteInterface[DataFlow]

	DataFlowShapes                map[*DataFlowShape]struct{}
	DataFlowShapes_instance       map[*DataFlowShape]*DataFlowShape
	DataFlowShapes_mapString      map[string]*DataFlowShape
	DataFlowShapeOrder            uint
	DataFlowShape_stagedOrder     map[*DataFlowShape]uint
	DataFlowShape_orderStaged     map[uint]*DataFlowShape
	DataFlowShapes_reference      map[*DataFlowShape]*DataFlowShape
	DataFlowShapes_referenceOrder map[*DataFlowShape]uint

	// insertion point for slice of pointers maps
	OnAfterDataFlowShapeCreateCallback GongOnAfterCreateInterface[DataFlowShape]
	OnAfterDataFlowShapeUpdateCallback GongOnAfterUpdateInterface[DataFlowShape]
	OnAfterDataFlowShapeDeleteCallback GongOnAfterDeleteInterface[DataFlowShape]

	DataShapes                map[*DataShape]struct{}
	DataShapes_instance       map[*DataShape]*DataShape
	DataShapes_mapString      map[string]*DataShape
	DataShapeOrder            uint
	DataShape_stagedOrder     map[*DataShape]uint
	DataShape_orderStaged     map[uint]*DataShape
	DataShapes_reference      map[*DataShape]*DataShape
	DataShapes_referenceOrder map[*DataShape]uint

	// insertion point for slice of pointers maps
	OnAfterDataShapeCreateCallback GongOnAfterCreateInterface[DataShape]
	OnAfterDataShapeUpdateCallback GongOnAfterUpdateInterface[DataShape]
	OnAfterDataShapeDeleteCallback GongOnAfterDeleteInterface[DataShape]

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

	DiagramProcess_ExternalParticipant_Shapes_reverseMap map[*ExternalParticipantShape]*DiagramProcess

	DiagramProcess_ExternalParticipantWhoseNodeIsExpanded_reverseMap map[*Participant]*DiagramProcess

	DiagramProcess_ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded_reverseMap map[*Participant]*DiagramProcess

	DiagramProcess_ExternalParticipantsWhoseInDataFlowsNodeIsExpanded_reverseMap map[*Participant]*DiagramProcess

	DiagramProcess_TasksWhoseNodeIsExpanded_reverseMap map[*Task]*DiagramProcess

	DiagramProcess_Task_Shapes_reverseMap map[*TaskShape]*DiagramProcess

	DiagramProcess_ControlFlowsWhoseNodeIsExpanded_reverseMap map[*ControlFlow]*DiagramProcess

	DiagramProcess_ControlFlow_Shapes_reverseMap map[*ControlFlowShape]*DiagramProcess

	DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap map[*DataFlow]*DiagramProcess

	DiagramProcess_DataFlow_Shapes_reverseMap map[*DataFlowShape]*DiagramProcess

	DiagramProcess_DatasWhoseNodeIsExpanded_reverseMap map[*Data]*DiagramProcess

	DiagramProcess_Data_Shapes_reverseMap map[*DataShape]*DiagramProcess

	DiagramProcess_DataFlowsWhoseDataNodeIsExpanded_reverseMap map[*DataFlow]*DiagramProcess

	DiagramProcess_AllocatedResourcesWhoseNodeIsExpanded_reverseMap map[*Resource]*DiagramProcess

	DiagramProcess_AllocatedResourceShapes_reverseMap map[*AllocatedResourceShape]*DiagramProcess

	DiagramProcess_AllocatedProcessesWhoseNodeIsExpanded_reverseMap map[*Process]*DiagramProcess

	DiagramProcess_AllocatedProcessShapes_reverseMap map[*AllocatedProcessShape]*DiagramProcess

	DiagramProcess_Note_Shapes_reverseMap map[*NoteShape]*DiagramProcess

	DiagramProcess_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*DiagramProcess

	DiagramProcess_NoteTaskShapes_reverseMap map[*NoteTaskShape]*DiagramProcess

	OnAfterDiagramProcessCreateCallback GongOnAfterCreateInterface[DiagramProcess]
	OnAfterDiagramProcessUpdateCallback GongOnAfterUpdateInterface[DiagramProcess]
	OnAfterDiagramProcessDeleteCallback GongOnAfterDeleteInterface[DiagramProcess]

	ExternalParticipantShapes                map[*ExternalParticipantShape]struct{}
	ExternalParticipantShapes_instance       map[*ExternalParticipantShape]*ExternalParticipantShape
	ExternalParticipantShapes_mapString      map[string]*ExternalParticipantShape
	ExternalParticipantShapeOrder            uint
	ExternalParticipantShape_stagedOrder     map[*ExternalParticipantShape]uint
	ExternalParticipantShape_orderStaged     map[uint]*ExternalParticipantShape
	ExternalParticipantShapes_reference      map[*ExternalParticipantShape]*ExternalParticipantShape
	ExternalParticipantShapes_referenceOrder map[*ExternalParticipantShape]uint

	// insertion point for slice of pointers maps
	OnAfterExternalParticipantShapeCreateCallback GongOnAfterCreateInterface[ExternalParticipantShape]
	OnAfterExternalParticipantShapeUpdateCallback GongOnAfterUpdateInterface[ExternalParticipantShape]
	OnAfterExternalParticipantShapeDeleteCallback GongOnAfterDeleteInterface[ExternalParticipantShape]

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

	Library_SubLibrariesWhoseNodeIsExpanded_reverseMap map[*Library]*Library

	Library_RootProcesses_reverseMap map[*Process]*Library

	Library_ProcesssWhoseNodeIsExpanded_reverseMap map[*Process]*Library

	Library_RootDataFlows_reverseMap map[*DataFlow]*Library

	Library_DataFlowsWhoseNodeIsExpanded_reverseMap map[*DataFlow]*Library

	Library_RootDatas_reverseMap map[*Data]*Library

	Library_DatasWhoseNodeIsExpanded_reverseMap map[*Data]*Library

	Library_RootResources_reverseMap map[*Resource]*Library

	Library_ResourcesWhoseNodeIsExpanded_reverseMap map[*Resource]*Library

	Library_ParticipantsWhoseNodeIsExpanded_reverseMap map[*Participant]*Library

	Library_RootNotes_reverseMap map[*Note]*Library

	Library_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*Library

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
	Note_Tasks_reverseMap map[*Task]*Note

	OnAfterNoteCreateCallback GongOnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback GongOnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback GongOnAfterDeleteInterface[Note]

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

	Participants                map[*Participant]struct{}
	Participants_instance       map[*Participant]*Participant
	Participants_mapString      map[string]*Participant
	ParticipantOrder            uint
	Participant_stagedOrder     map[*Participant]uint
	Participant_orderStaged     map[uint]*Participant
	Participants_reference      map[*Participant]*Participant
	Participants_referenceOrder map[*Participant]uint

	// insertion point for slice of pointers maps
	Participant_Resources_reverseMap map[*Resource]*Participant

	Participant_Processes_reverseMap map[*Process]*Participant

	Participant_Tasks_reverseMap map[*Task]*Participant

	Participant_ControlFlows_reverseMap map[*ControlFlow]*Participant

	Participant_TaskWhoseOutControlFlowsNodeIsExpanded_reverseMap map[*Task]*Participant

	Participant_TaskWhoseInControlFlowsNodeIsExpanded_reverseMap map[*Task]*Participant

	Participant_TaskWhoseOutDataFlowsNodeIsExpanded_reverseMap map[*Task]*Participant

	Participant_TaskWhoseInDataFlowsNodeIsExpanded_reverseMap map[*Task]*Participant

	OnAfterParticipantCreateCallback GongOnAfterCreateInterface[Participant]
	OnAfterParticipantUpdateCallback GongOnAfterUpdateInterface[Participant]
	OnAfterParticipantDeleteCallback GongOnAfterDeleteInterface[Participant]

	ParticipantShapes                map[*ParticipantShape]struct{}
	ParticipantShapes_instance       map[*ParticipantShape]*ParticipantShape
	ParticipantShapes_mapString      map[string]*ParticipantShape
	ParticipantShapeOrder            uint
	ParticipantShape_stagedOrder     map[*ParticipantShape]uint
	ParticipantShape_orderStaged     map[uint]*ParticipantShape
	ParticipantShapes_reference      map[*ParticipantShape]*ParticipantShape
	ParticipantShapes_referenceOrder map[*ParticipantShape]uint

	// insertion point for slice of pointers maps
	OnAfterParticipantShapeCreateCallback GongOnAfterCreateInterface[ParticipantShape]
	OnAfterParticipantShapeUpdateCallback GongOnAfterUpdateInterface[ParticipantShape]
	OnAfterParticipantShapeDeleteCallback GongOnAfterDeleteInterface[ParticipantShape]

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

	Process_DataFlows_reverseMap map[*DataFlow]*Process

	Process_ExternalParticipants_reverseMap map[*Participant]*Process

	Process_ExternalParticipantWhoseNodeIsExpanded_reverseMap map[*Participant]*Process

	OnAfterProcessCreateCallback GongOnAfterCreateInterface[Process]
	OnAfterProcessUpdateCallback GongOnAfterUpdateInterface[Process]
	OnAfterProcessDeleteCallback GongOnAfterDeleteInterface[Process]

	ProcessShapes                map[*ProcessShape]struct{}
	ProcessShapes_instance       map[*ProcessShape]*ProcessShape
	ProcessShapes_mapString      map[string]*ProcessShape
	ProcessShapeOrder            uint
	ProcessShape_stagedOrder     map[*ProcessShape]uint
	ProcessShape_orderStaged     map[uint]*ProcessShape
	ProcessShapes_reference      map[*ProcessShape]*ProcessShape
	ProcessShapes_referenceOrder map[*ProcessShape]uint

	// insertion point for slice of pointers maps
	OnAfterProcessShapeCreateCallback GongOnAfterCreateInterface[ProcessShape]
	OnAfterProcessShapeUpdateCallback GongOnAfterUpdateInterface[ProcessShape]
	OnAfterProcessShapeDeleteCallback GongOnAfterDeleteInterface[ProcessShape]

	Resources                map[*Resource]struct{}
	Resources_instance       map[*Resource]*Resource
	Resources_mapString      map[string]*Resource
	ResourceOrder            uint
	Resource_stagedOrder     map[*Resource]uint
	Resource_orderStaged     map[uint]*Resource
	Resources_reference      map[*Resource]*Resource
	Resources_referenceOrder map[*Resource]uint

	// insertion point for slice of pointers maps
	OnAfterResourceCreateCallback GongOnAfterCreateInterface[Resource]
	OnAfterResourceUpdateCallback GongOnAfterUpdateInterface[Resource]
	OnAfterResourceDeleteCallback GongOnAfterDeleteInterface[Resource]

	Tasks                map[*Task]struct{}
	Tasks_instance       map[*Task]*Task
	Tasks_mapString      map[string]*Task
	TaskOrder            uint
	Task_stagedOrder     map[*Task]uint
	Task_orderStaged     map[uint]*Task
	Tasks_reference      map[*Task]*Task
	Tasks_referenceOrder map[*Task]uint

	// insertion point for slice of pointers maps
	OnAfterTaskCreateCallback GongOnAfterCreateInterface[Task]
	OnAfterTaskUpdateCallback GongOnAfterUpdateInterface[Task]
	OnAfterTaskDeleteCallback GongOnAfterDeleteInterface[Task]

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
	__gong__clearReferences(&stage.AllocatedProcessShapes_reference, &stage.AllocatedProcessShapes_instance, &stage.AllocatedProcessShapes_referenceOrder)

	__gong__clearReferences(&stage.AllocatedResourceShapes_reference, &stage.AllocatedResourceShapes_instance, &stage.AllocatedResourceShapes_referenceOrder)

	__gong__clearReferences(&stage.ControlFlows_reference, &stage.ControlFlows_instance, &stage.ControlFlows_referenceOrder)

	__gong__clearReferences(&stage.ControlFlowShapes_reference, &stage.ControlFlowShapes_instance, &stage.ControlFlowShapes_referenceOrder)

	__gong__clearReferences(&stage.Datas_reference, &stage.Datas_instance, &stage.Datas_referenceOrder)

	__gong__clearReferences(&stage.DataFlows_reference, &stage.DataFlows_instance, &stage.DataFlows_referenceOrder)

	__gong__clearReferences(&stage.DataFlowShapes_reference, &stage.DataFlowShapes_instance, &stage.DataFlowShapes_referenceOrder)

	__gong__clearReferences(&stage.DataShapes_reference, &stage.DataShapes_instance, &stage.DataShapes_referenceOrder)

	__gong__clearReferences(&stage.DiagramProcesss_reference, &stage.DiagramProcesss_instance, &stage.DiagramProcesss_referenceOrder)

	__gong__clearReferences(&stage.ExternalParticipantShapes_reference, &stage.ExternalParticipantShapes_instance, &stage.ExternalParticipantShapes_referenceOrder)

	__gong__clearReferences(&stage.Librarys_reference, &stage.Librarys_instance, &stage.Librarys_referenceOrder)

	__gong__clearReferences(&stage.Notes_reference, &stage.Notes_instance, &stage.Notes_referenceOrder)

	__gong__clearReferences(&stage.NoteShapes_reference, &stage.NoteShapes_instance, &stage.NoteShapes_referenceOrder)

	__gong__clearReferences(&stage.NoteTaskShapes_reference, &stage.NoteTaskShapes_instance, &stage.NoteTaskShapes_referenceOrder)

	__gong__clearReferences(&stage.Participants_reference, &stage.Participants_instance, &stage.Participants_referenceOrder)

	__gong__clearReferences(&stage.ParticipantShapes_reference, &stage.ParticipantShapes_instance, &stage.ParticipantShapes_referenceOrder)

	__gong__clearReferences(&stage.Processs_reference, &stage.Processs_instance, &stage.Processs_referenceOrder)

	__gong__clearReferences(&stage.ProcessShapes_reference, &stage.ProcessShapes_instance, &stage.ProcessShapes_referenceOrder)

	__gong__clearReferences(&stage.Resources_reference, &stage.Resources_instance, &stage.Resources_referenceOrder)

	__gong__clearReferences(&stage.Tasks_reference, &stage.Tasks_instance, &stage.Tasks_referenceOrder)

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
	stage.AllocatedProcessShapeOrder = __gong__recomputeOrder(stage.AllocatedProcessShape_stagedOrder)

	stage.AllocatedResourceShapeOrder = __gong__recomputeOrder(stage.AllocatedResourceShape_stagedOrder)

	stage.ControlFlowOrder = __gong__recomputeOrder(stage.ControlFlow_stagedOrder)

	stage.ControlFlowShapeOrder = __gong__recomputeOrder(stage.ControlFlowShape_stagedOrder)

	stage.DataOrder = __gong__recomputeOrder(stage.Data_stagedOrder)

	stage.DataFlowOrder = __gong__recomputeOrder(stage.DataFlow_stagedOrder)

	stage.DataFlowShapeOrder = __gong__recomputeOrder(stage.DataFlowShape_stagedOrder)

	stage.DataShapeOrder = __gong__recomputeOrder(stage.DataShape_stagedOrder)

	stage.DiagramProcessOrder = __gong__recomputeOrder(stage.DiagramProcess_stagedOrder)

	stage.ExternalParticipantShapeOrder = __gong__recomputeOrder(stage.ExternalParticipantShape_stagedOrder)

	stage.LibraryOrder = __gong__recomputeOrder(stage.Library_stagedOrder)

	stage.NoteOrder = __gong__recomputeOrder(stage.Note_stagedOrder)

	stage.NoteShapeOrder = __gong__recomputeOrder(stage.NoteShape_stagedOrder)

	stage.NoteTaskShapeOrder = __gong__recomputeOrder(stage.NoteTaskShape_stagedOrder)

	stage.ParticipantOrder = __gong__recomputeOrder(stage.Participant_stagedOrder)

	stage.ParticipantShapeOrder = __gong__recomputeOrder(stage.ParticipantShape_stagedOrder)

	stage.ProcessOrder = __gong__recomputeOrder(stage.Process_stagedOrder)

	stage.ProcessShapeOrder = __gong__recomputeOrder(stage.ProcessShape_stagedOrder)

	stage.ResourceOrder = __gong__recomputeOrder(stage.Resource_stagedOrder)

	stage.TaskOrder = __gong__recomputeOrder(stage.Task_stagedOrder)

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
	case *AllocatedProcessShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.AllocatedProcessShapes, stage.AllocatedProcessShape_stagedOrder))
	case *AllocatedResourceShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder))
	case *ControlFlow:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ControlFlows, stage.ControlFlow_stagedOrder))
	case *ControlFlowShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder))
	case *Data:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Datas, stage.Data_stagedOrder))
	case *DataFlow:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DataFlows, stage.DataFlow_stagedOrder))
	case *DataFlowShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder))
	case *DataShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DataShapes, stage.DataShape_stagedOrder))
	case *DiagramProcess:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder))
	case *ExternalParticipantShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ExternalParticipantShapes, stage.ExternalParticipantShape_stagedOrder))
	case *Library:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder))
	case *Note:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Notes, stage.Note_stagedOrder))
	case *NoteShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteShapes, stage.NoteShape_stagedOrder))
	case *NoteTaskShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder))
	case *Participant:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Participants, stage.Participant_stagedOrder))
	case *ParticipantShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ParticipantShapes, stage.ParticipantShape_stagedOrder))
	case *Process:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Processs, stage.Process_stagedOrder))
	case *ProcessShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.ProcessShapes, stage.ProcessShape_stagedOrder))
	case *Resource:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Resources, stage.Resource_stagedOrder))
	case *Task:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Tasks, stage.Task_stagedOrder))
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
	return "github.com/fullstack-lang/gong/dsm/process/go/models"
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
		AllocatedProcessShapes:           make(map[*AllocatedProcessShape]struct{}),
		AllocatedProcessShapes_mapString: make(map[string]*AllocatedProcessShape),

		AllocatedResourceShapes:           make(map[*AllocatedResourceShape]struct{}),
		AllocatedResourceShapes_mapString: make(map[string]*AllocatedResourceShape),

		ControlFlows:           make(map[*ControlFlow]struct{}),
		ControlFlows_mapString: make(map[string]*ControlFlow),

		ControlFlowShapes:           make(map[*ControlFlowShape]struct{}),
		ControlFlowShapes_mapString: make(map[string]*ControlFlowShape),

		Datas:           make(map[*Data]struct{}),
		Datas_mapString: make(map[string]*Data),

		DataFlows:           make(map[*DataFlow]struct{}),
		DataFlows_mapString: make(map[string]*DataFlow),

		DataFlowShapes:           make(map[*DataFlowShape]struct{}),
		DataFlowShapes_mapString: make(map[string]*DataFlowShape),

		DataShapes:           make(map[*DataShape]struct{}),
		DataShapes_mapString: make(map[string]*DataShape),

		DiagramProcesss:           make(map[*DiagramProcess]struct{}),
		DiagramProcesss_mapString: make(map[string]*DiagramProcess),

		ExternalParticipantShapes:           make(map[*ExternalParticipantShape]struct{}),
		ExternalParticipantShapes_mapString: make(map[string]*ExternalParticipantShape),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		NoteShapes:           make(map[*NoteShape]struct{}),
		NoteShapes_mapString: make(map[string]*NoteShape),

		NoteTaskShapes:           make(map[*NoteTaskShape]struct{}),
		NoteTaskShapes_mapString: make(map[string]*NoteTaskShape),

		Participants:           make(map[*Participant]struct{}),
		Participants_mapString: make(map[string]*Participant),

		ParticipantShapes:           make(map[*ParticipantShape]struct{}),
		ParticipantShapes_mapString: make(map[string]*ParticipantShape),

		Processs:           make(map[*Process]struct{}),
		Processs_mapString: make(map[string]*Process),

		ProcessShapes:           make(map[*ProcessShape]struct{}),
		ProcessShapes_mapString: make(map[string]*ProcessShape),

		Resources:           make(map[*Resource]struct{}),
		Resources_mapString: make(map[string]*Resource),

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
		AllocatedProcessShape_stagedOrder: make(map[*AllocatedProcessShape]uint),
		AllocatedProcessShape_orderStaged: make(map[uint]*AllocatedProcessShape),
		AllocatedProcessShapes_reference:  make(map[*AllocatedProcessShape]*AllocatedProcessShape),

		AllocatedResourceShape_stagedOrder: make(map[*AllocatedResourceShape]uint),
		AllocatedResourceShape_orderStaged: make(map[uint]*AllocatedResourceShape),
		AllocatedResourceShapes_reference:  make(map[*AllocatedResourceShape]*AllocatedResourceShape),

		ControlFlow_stagedOrder: make(map[*ControlFlow]uint),
		ControlFlow_orderStaged: make(map[uint]*ControlFlow),
		ControlFlows_reference:  make(map[*ControlFlow]*ControlFlow),

		ControlFlowShape_stagedOrder: make(map[*ControlFlowShape]uint),
		ControlFlowShape_orderStaged: make(map[uint]*ControlFlowShape),
		ControlFlowShapes_reference:  make(map[*ControlFlowShape]*ControlFlowShape),

		Data_stagedOrder: make(map[*Data]uint),
		Data_orderStaged: make(map[uint]*Data),
		Datas_reference:  make(map[*Data]*Data),

		DataFlow_stagedOrder: make(map[*DataFlow]uint),
		DataFlow_orderStaged: make(map[uint]*DataFlow),
		DataFlows_reference:  make(map[*DataFlow]*DataFlow),

		DataFlowShape_stagedOrder: make(map[*DataFlowShape]uint),
		DataFlowShape_orderStaged: make(map[uint]*DataFlowShape),
		DataFlowShapes_reference:  make(map[*DataFlowShape]*DataFlowShape),

		DataShape_stagedOrder: make(map[*DataShape]uint),
		DataShape_orderStaged: make(map[uint]*DataShape),
		DataShapes_reference:  make(map[*DataShape]*DataShape),

		DiagramProcess_stagedOrder: make(map[*DiagramProcess]uint),
		DiagramProcess_orderStaged: make(map[uint]*DiagramProcess),
		DiagramProcesss_reference:  make(map[*DiagramProcess]*DiagramProcess),

		ExternalParticipantShape_stagedOrder: make(map[*ExternalParticipantShape]uint),
		ExternalParticipantShape_orderStaged: make(map[uint]*ExternalParticipantShape),
		ExternalParticipantShapes_reference:  make(map[*ExternalParticipantShape]*ExternalParticipantShape),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		Note_stagedOrder: make(map[*Note]uint),
		Note_orderStaged: make(map[uint]*Note),
		Notes_reference:  make(map[*Note]*Note),

		NoteShape_stagedOrder: make(map[*NoteShape]uint),
		NoteShape_orderStaged: make(map[uint]*NoteShape),
		NoteShapes_reference:  make(map[*NoteShape]*NoteShape),

		NoteTaskShape_stagedOrder: make(map[*NoteTaskShape]uint),
		NoteTaskShape_orderStaged: make(map[uint]*NoteTaskShape),
		NoteTaskShapes_reference:  make(map[*NoteTaskShape]*NoteTaskShape),

		Participant_stagedOrder: make(map[*Participant]uint),
		Participant_orderStaged: make(map[uint]*Participant),
		Participants_reference:  make(map[*Participant]*Participant),

		ParticipantShape_stagedOrder: make(map[*ParticipantShape]uint),
		ParticipantShape_orderStaged: make(map[uint]*ParticipantShape),
		ParticipantShapes_reference:  make(map[*ParticipantShape]*ParticipantShape),

		Process_stagedOrder: make(map[*Process]uint),
		Process_orderStaged: make(map[uint]*Process),
		Processs_reference:  make(map[*Process]*Process),

		ProcessShape_stagedOrder: make(map[*ProcessShape]uint),
		ProcessShape_orderStaged: make(map[uint]*ProcessShape),
		ProcessShapes_reference:  make(map[*ProcessShape]*ProcessShape),

		Resource_stagedOrder: make(map[*Resource]uint),
		Resource_orderStaged: make(map[uint]*Resource),
		Resources_reference:  make(map[*Resource]*Resource),

		Task_stagedOrder: make(map[*Task]uint),
		Task_orderStaged: make(map[uint]*Task),
		Tasks_reference:  make(map[*Task]*Task),

		TaskShape_stagedOrder: make(map[*TaskShape]uint),
		TaskShape_orderStaged: make(map[uint]*TaskShape),
		TaskShapes_reference:  make(map[*TaskShape]*TaskShape),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"AllocatedProcessShape": &AllocatedProcessShapeUnmarshaller{},

			"AllocatedResourceShape": &AllocatedResourceShapeUnmarshaller{},

			"ControlFlow": &ControlFlowUnmarshaller{},

			"ControlFlowShape": &ControlFlowShapeUnmarshaller{},

			"Data": &DataUnmarshaller{},

			"DataFlow": &DataFlowUnmarshaller{},

			"DataFlowShape": &DataFlowShapeUnmarshaller{},

			"DataShape": &DataShapeUnmarshaller{},

			"DiagramProcess": &DiagramProcessUnmarshaller{},

			"ExternalParticipantShape": &ExternalParticipantShapeUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"NoteTaskShape": &NoteTaskShapeUnmarshaller{},

			"Participant": &ParticipantUnmarshaller{},

			"ParticipantShape": &ParticipantShapeUnmarshaller{},

			"Process": &ProcessUnmarshaller{},

			"ProcessShape": &ProcessShapeUnmarshaller{},

			"Resource": &ResourceUnmarshaller{},

			"Task": &TaskUnmarshaller{},

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
	case *AllocatedProcessShape:
		return any(stage.AllocatedProcessShape_orderStaged[order]).(Type)
	case *AllocatedResourceShape:
		return any(stage.AllocatedResourceShape_orderStaged[order]).(Type)
	case *ControlFlow:
		return any(stage.ControlFlow_orderStaged[order]).(Type)
	case *ControlFlowShape:
		return any(stage.ControlFlowShape_orderStaged[order]).(Type)
	case *Data:
		return any(stage.Data_orderStaged[order]).(Type)
	case *DataFlow:
		return any(stage.DataFlow_orderStaged[order]).(Type)
	case *DataFlowShape:
		return any(stage.DataFlowShape_orderStaged[order]).(Type)
	case *DataShape:
		return any(stage.DataShape_orderStaged[order]).(Type)
	case *DiagramProcess:
		return any(stage.DiagramProcess_orderStaged[order]).(Type)
	case *ExternalParticipantShape:
		return any(stage.ExternalParticipantShape_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Note:
		return any(stage.Note_orderStaged[order]).(Type)
	case *NoteShape:
		return any(stage.NoteShape_orderStaged[order]).(Type)
	case *NoteTaskShape:
		return any(stage.NoteTaskShape_orderStaged[order]).(Type)
	case *Participant:
		return any(stage.Participant_orderStaged[order]).(Type)
	case *ParticipantShape:
		return any(stage.ParticipantShape_orderStaged[order]).(Type)
	case *Process:
		return any(stage.Process_orderStaged[order]).(Type)
	case *ProcessShape:
		return any(stage.ProcessShape_orderStaged[order]).(Type)
	case *Resource:
		return any(stage.Resource_orderStaged[order]).(Type)
	case *Task:
		return any(stage.Task_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["AllocatedProcessShape"] = len(stage.AllocatedProcessShapes)
	stage.Map_GongStructName_InstancesNb["AllocatedResourceShape"] = len(stage.AllocatedResourceShapes)
	stage.Map_GongStructName_InstancesNb["ControlFlow"] = len(stage.ControlFlows)
	stage.Map_GongStructName_InstancesNb["ControlFlowShape"] = len(stage.ControlFlowShapes)
	stage.Map_GongStructName_InstancesNb["Data"] = len(stage.Datas)
	stage.Map_GongStructName_InstancesNb["DataFlow"] = len(stage.DataFlows)
	stage.Map_GongStructName_InstancesNb["DataFlowShape"] = len(stage.DataFlowShapes)
	stage.Map_GongStructName_InstancesNb["DataShape"] = len(stage.DataShapes)
	stage.Map_GongStructName_InstancesNb["DiagramProcess"] = len(stage.DiagramProcesss)
	stage.Map_GongStructName_InstancesNb["ExternalParticipantShape"] = len(stage.ExternalParticipantShapes)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteTaskShape"] = len(stage.NoteTaskShapes)
	stage.Map_GongStructName_InstancesNb["Participant"] = len(stage.Participants)
	stage.Map_GongStructName_InstancesNb["ParticipantShape"] = len(stage.ParticipantShapes)
	stage.Map_GongStructName_InstancesNb["Process"] = len(stage.Processs)
	stage.Map_GongStructName_InstancesNb["ProcessShape"] = len(stage.ProcessShapes)
	stage.Map_GongStructName_InstancesNb["Resource"] = len(stage.Resources)
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
// Stage puts allocatedprocessshape to the model stage
func (allocatedprocessshape *AllocatedProcessShape) Stage(stage *Stage) *AllocatedProcessShape {
	__gong__stage(stage.AllocatedProcessShapes, stage.AllocatedProcessShape_stagedOrder, stage.AllocatedProcessShape_orderStaged, &stage.AllocatedProcessShapeOrder, stage.AllocatedProcessShapes_mapString, allocatedprocessshape, allocatedprocessshape.Name)
	return allocatedprocessshape
}

// StagePreserveOrder puts allocatedprocessshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AllocatedProcessShapeOrder
// - update stage.AllocatedProcessShapeOrder accordingly
func (allocatedprocessshape *AllocatedProcessShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AllocatedProcessShapes, stage.AllocatedProcessShape_stagedOrder, stage.AllocatedProcessShape_orderStaged, &stage.AllocatedProcessShapeOrder, stage.AllocatedProcessShapes_mapString, allocatedprocessshape, order, allocatedprocessshape.Name)
}

// Unstage removes allocatedprocessshape off the model stage
func (allocatedprocessshape *AllocatedProcessShape) Unstage(stage *Stage) *AllocatedProcessShape {
	__gong__unstage(stage.AllocatedProcessShapes, stage.AllocatedProcessShapes_mapString, allocatedprocessshape, allocatedprocessshape.Name)
	return allocatedprocessshape
}

// UnstageVoid removes allocatedprocessshape off the model stage
func (allocatedprocessshape *AllocatedProcessShape) UnstageVoid(stage *Stage) {
	allocatedprocessshape.Unstage(stage)
}

func (allocatedprocessshape *AllocatedProcessShape) StageVoid(stage *Stage) {
	allocatedprocessshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (allocatedprocessshape *AllocatedProcessShape) GetName() (res string) {
	return allocatedprocessshape.Name
}

// for satisfaction of GongStruct interface
func (allocatedprocessshape *AllocatedProcessShape) SetName(name string) {
	allocatedprocessshape.Name = name
}

// Stage puts allocatedresourceshape to the model stage
func (allocatedresourceshape *AllocatedResourceShape) Stage(stage *Stage) *AllocatedResourceShape {
	__gong__stage(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder, stage.AllocatedResourceShape_orderStaged, &stage.AllocatedResourceShapeOrder, stage.AllocatedResourceShapes_mapString, allocatedresourceshape, allocatedresourceshape.Name)
	return allocatedresourceshape
}

// StagePreserveOrder puts allocatedresourceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AllocatedResourceShapeOrder
// - update stage.AllocatedResourceShapeOrder accordingly
func (allocatedresourceshape *AllocatedResourceShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder, stage.AllocatedResourceShape_orderStaged, &stage.AllocatedResourceShapeOrder, stage.AllocatedResourceShapes_mapString, allocatedresourceshape, order, allocatedresourceshape.Name)
}

// Unstage removes allocatedresourceshape off the model stage
func (allocatedresourceshape *AllocatedResourceShape) Unstage(stage *Stage) *AllocatedResourceShape {
	__gong__unstage(stage.AllocatedResourceShapes, stage.AllocatedResourceShapes_mapString, allocatedresourceshape, allocatedresourceshape.Name)
	return allocatedresourceshape
}

// UnstageVoid removes allocatedresourceshape off the model stage
func (allocatedresourceshape *AllocatedResourceShape) UnstageVoid(stage *Stage) {
	allocatedresourceshape.Unstage(stage)
}

func (allocatedresourceshape *AllocatedResourceShape) StageVoid(stage *Stage) {
	allocatedresourceshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (allocatedresourceshape *AllocatedResourceShape) GetName() (res string) {
	return allocatedresourceshape.Name
}

// for satisfaction of GongStruct interface
func (allocatedresourceshape *AllocatedResourceShape) SetName(name string) {
	allocatedresourceshape.Name = name
}

// Stage puts controlflow to the model stage
func (controlflow *ControlFlow) Stage(stage *Stage) *ControlFlow {
	__gong__stage(stage.ControlFlows, stage.ControlFlow_stagedOrder, stage.ControlFlow_orderStaged, &stage.ControlFlowOrder, stage.ControlFlows_mapString, controlflow, controlflow.Name)
	return controlflow
}

// StagePreserveOrder puts controlflow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlFlowOrder
// - update stage.ControlFlowOrder accordingly
func (controlflow *ControlFlow) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ControlFlows, stage.ControlFlow_stagedOrder, stage.ControlFlow_orderStaged, &stage.ControlFlowOrder, stage.ControlFlows_mapString, controlflow, order, controlflow.Name)
}

// Unstage removes controlflow off the model stage
func (controlflow *ControlFlow) Unstage(stage *Stage) *ControlFlow {
	__gong__unstage(stage.ControlFlows, stage.ControlFlows_mapString, controlflow, controlflow.Name)
	return controlflow
}

// UnstageVoid removes controlflow off the model stage
func (controlflow *ControlFlow) UnstageVoid(stage *Stage) {
	controlflow.Unstage(stage)
}

func (controlflow *ControlFlow) StageVoid(stage *Stage) {
	controlflow.Stage(stage)
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
	__gong__stage(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder, stage.ControlFlowShape_orderStaged, &stage.ControlFlowShapeOrder, stage.ControlFlowShapes_mapString, controlflowshape, controlflowshape.Name)
	return controlflowshape
}

// StagePreserveOrder puts controlflowshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ControlFlowShapeOrder
// - update stage.ControlFlowShapeOrder accordingly
func (controlflowshape *ControlFlowShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder, stage.ControlFlowShape_orderStaged, &stage.ControlFlowShapeOrder, stage.ControlFlowShapes_mapString, controlflowshape, order, controlflowshape.Name)
}

// Unstage removes controlflowshape off the model stage
func (controlflowshape *ControlFlowShape) Unstage(stage *Stage) *ControlFlowShape {
	__gong__unstage(stage.ControlFlowShapes, stage.ControlFlowShapes_mapString, controlflowshape, controlflowshape.Name)
	return controlflowshape
}

// UnstageVoid removes controlflowshape off the model stage
func (controlflowshape *ControlFlowShape) UnstageVoid(stage *Stage) {
	controlflowshape.Unstage(stage)
}

func (controlflowshape *ControlFlowShape) StageVoid(stage *Stage) {
	controlflowshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (controlflowshape *ControlFlowShape) GetName() (res string) {
	return controlflowshape.Name
}

// for satisfaction of GongStruct interface
func (controlflowshape *ControlFlowShape) SetName(name string) {
	controlflowshape.Name = name
}

// Stage puts data to the model stage
func (data *Data) Stage(stage *Stage) *Data {
	__gong__stage(stage.Datas, stage.Data_stagedOrder, stage.Data_orderStaged, &stage.DataOrder, stage.Datas_mapString, data, data.Name)
	return data
}

// StagePreserveOrder puts data to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataOrder
// - update stage.DataOrder accordingly
func (data *Data) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Datas, stage.Data_stagedOrder, stage.Data_orderStaged, &stage.DataOrder, stage.Datas_mapString, data, order, data.Name)
}

// Unstage removes data off the model stage
func (data *Data) Unstage(stage *Stage) *Data {
	__gong__unstage(stage.Datas, stage.Datas_mapString, data, data.Name)
	return data
}

// UnstageVoid removes data off the model stage
func (data *Data) UnstageVoid(stage *Stage) {
	data.Unstage(stage)
}

func (data *Data) StageVoid(stage *Stage) {
	data.Stage(stage)
}

// for satisfaction of GongStruct interface
func (data *Data) GetName() (res string) {
	return data.Name
}

// for satisfaction of GongStruct interface
func (data *Data) SetName(name string) {
	data.Name = name
}

// Stage puts dataflow to the model stage
func (dataflow *DataFlow) Stage(stage *Stage) *DataFlow {
	__gong__stage(stage.DataFlows, stage.DataFlow_stagedOrder, stage.DataFlow_orderStaged, &stage.DataFlowOrder, stage.DataFlows_mapString, dataflow, dataflow.Name)
	return dataflow
}

// StagePreserveOrder puts dataflow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataFlowOrder
// - update stage.DataFlowOrder accordingly
func (dataflow *DataFlow) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DataFlows, stage.DataFlow_stagedOrder, stage.DataFlow_orderStaged, &stage.DataFlowOrder, stage.DataFlows_mapString, dataflow, order, dataflow.Name)
}

// Unstage removes dataflow off the model stage
func (dataflow *DataFlow) Unstage(stage *Stage) *DataFlow {
	__gong__unstage(stage.DataFlows, stage.DataFlows_mapString, dataflow, dataflow.Name)
	return dataflow
}

// UnstageVoid removes dataflow off the model stage
func (dataflow *DataFlow) UnstageVoid(stage *Stage) {
	dataflow.Unstage(stage)
}

func (dataflow *DataFlow) StageVoid(stage *Stage) {
	dataflow.Stage(stage)
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
	__gong__stage(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder, stage.DataFlowShape_orderStaged, &stage.DataFlowShapeOrder, stage.DataFlowShapes_mapString, dataflowshape, dataflowshape.Name)
	return dataflowshape
}

// StagePreserveOrder puts dataflowshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataFlowShapeOrder
// - update stage.DataFlowShapeOrder accordingly
func (dataflowshape *DataFlowShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder, stage.DataFlowShape_orderStaged, &stage.DataFlowShapeOrder, stage.DataFlowShapes_mapString, dataflowshape, order, dataflowshape.Name)
}

// Unstage removes dataflowshape off the model stage
func (dataflowshape *DataFlowShape) Unstage(stage *Stage) *DataFlowShape {
	__gong__unstage(stage.DataFlowShapes, stage.DataFlowShapes_mapString, dataflowshape, dataflowshape.Name)
	return dataflowshape
}

// UnstageVoid removes dataflowshape off the model stage
func (dataflowshape *DataFlowShape) UnstageVoid(stage *Stage) {
	dataflowshape.Unstage(stage)
}

func (dataflowshape *DataFlowShape) StageVoid(stage *Stage) {
	dataflowshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (dataflowshape *DataFlowShape) GetName() (res string) {
	return dataflowshape.Name
}

// for satisfaction of GongStruct interface
func (dataflowshape *DataFlowShape) SetName(name string) {
	dataflowshape.Name = name
}

// Stage puts datashape to the model stage
func (datashape *DataShape) Stage(stage *Stage) *DataShape {
	__gong__stage(stage.DataShapes, stage.DataShape_stagedOrder, stage.DataShape_orderStaged, &stage.DataShapeOrder, stage.DataShapes_mapString, datashape, datashape.Name)
	return datashape
}

// StagePreserveOrder puts datashape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataShapeOrder
// - update stage.DataShapeOrder accordingly
func (datashape *DataShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DataShapes, stage.DataShape_stagedOrder, stage.DataShape_orderStaged, &stage.DataShapeOrder, stage.DataShapes_mapString, datashape, order, datashape.Name)
}

// Unstage removes datashape off the model stage
func (datashape *DataShape) Unstage(stage *Stage) *DataShape {
	__gong__unstage(stage.DataShapes, stage.DataShapes_mapString, datashape, datashape.Name)
	return datashape
}

// UnstageVoid removes datashape off the model stage
func (datashape *DataShape) UnstageVoid(stage *Stage) {
	datashape.Unstage(stage)
}

func (datashape *DataShape) StageVoid(stage *Stage) {
	datashape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (datashape *DataShape) GetName() (res string) {
	return datashape.Name
}

// for satisfaction of GongStruct interface
func (datashape *DataShape) SetName(name string) {
	datashape.Name = name
}

// Stage puts diagramprocess to the model stage
func (diagramprocess *DiagramProcess) Stage(stage *Stage) *DiagramProcess {
	__gong__stage(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder, stage.DiagramProcess_orderStaged, &stage.DiagramProcessOrder, stage.DiagramProcesss_mapString, diagramprocess, diagramprocess.Name)
	return diagramprocess
}

// StagePreserveOrder puts diagramprocess to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramProcessOrder
// - update stage.DiagramProcessOrder accordingly
func (diagramprocess *DiagramProcess) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder, stage.DiagramProcess_orderStaged, &stage.DiagramProcessOrder, stage.DiagramProcesss_mapString, diagramprocess, order, diagramprocess.Name)
}

// Unstage removes diagramprocess off the model stage
func (diagramprocess *DiagramProcess) Unstage(stage *Stage) *DiagramProcess {
	__gong__unstage(stage.DiagramProcesss, stage.DiagramProcesss_mapString, diagramprocess, diagramprocess.Name)
	return diagramprocess
}

// UnstageVoid removes diagramprocess off the model stage
func (diagramprocess *DiagramProcess) UnstageVoid(stage *Stage) {
	diagramprocess.Unstage(stage)
}

func (diagramprocess *DiagramProcess) StageVoid(stage *Stage) {
	diagramprocess.Stage(stage)
}

// for satisfaction of GongStruct interface
func (diagramprocess *DiagramProcess) GetName() (res string) {
	return diagramprocess.Name
}

// for satisfaction of GongStruct interface
func (diagramprocess *DiagramProcess) SetName(name string) {
	diagramprocess.Name = name
}

// Stage puts externalparticipantshape to the model stage
func (externalparticipantshape *ExternalParticipantShape) Stage(stage *Stage) *ExternalParticipantShape {
	__gong__stage(stage.ExternalParticipantShapes, stage.ExternalParticipantShape_stagedOrder, stage.ExternalParticipantShape_orderStaged, &stage.ExternalParticipantShapeOrder, stage.ExternalParticipantShapes_mapString, externalparticipantshape, externalparticipantshape.Name)
	return externalparticipantshape
}

// StagePreserveOrder puts externalparticipantshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExternalParticipantShapeOrder
// - update stage.ExternalParticipantShapeOrder accordingly
func (externalparticipantshape *ExternalParticipantShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ExternalParticipantShapes, stage.ExternalParticipantShape_stagedOrder, stage.ExternalParticipantShape_orderStaged, &stage.ExternalParticipantShapeOrder, stage.ExternalParticipantShapes_mapString, externalparticipantshape, order, externalparticipantshape.Name)
}

// Unstage removes externalparticipantshape off the model stage
func (externalparticipantshape *ExternalParticipantShape) Unstage(stage *Stage) *ExternalParticipantShape {
	__gong__unstage(stage.ExternalParticipantShapes, stage.ExternalParticipantShapes_mapString, externalparticipantshape, externalparticipantshape.Name)
	return externalparticipantshape
}

// UnstageVoid removes externalparticipantshape off the model stage
func (externalparticipantshape *ExternalParticipantShape) UnstageVoid(stage *Stage) {
	externalparticipantshape.Unstage(stage)
}

func (externalparticipantshape *ExternalParticipantShape) StageVoid(stage *Stage) {
	externalparticipantshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (externalparticipantshape *ExternalParticipantShape) GetName() (res string) {
	return externalparticipantshape.Name
}

// for satisfaction of GongStruct interface
func (externalparticipantshape *ExternalParticipantShape) SetName(name string) {
	externalparticipantshape.Name = name
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

// Stage puts participant to the model stage
func (participant *Participant) Stage(stage *Stage) *Participant {
	__gong__stage(stage.Participants, stage.Participant_stagedOrder, stage.Participant_orderStaged, &stage.ParticipantOrder, stage.Participants_mapString, participant, participant.Name)
	return participant
}

// StagePreserveOrder puts participant to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParticipantOrder
// - update stage.ParticipantOrder accordingly
func (participant *Participant) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Participants, stage.Participant_stagedOrder, stage.Participant_orderStaged, &stage.ParticipantOrder, stage.Participants_mapString, participant, order, participant.Name)
}

// Unstage removes participant off the model stage
func (participant *Participant) Unstage(stage *Stage) *Participant {
	__gong__unstage(stage.Participants, stage.Participants_mapString, participant, participant.Name)
	return participant
}

// UnstageVoid removes participant off the model stage
func (participant *Participant) UnstageVoid(stage *Stage) {
	participant.Unstage(stage)
}

func (participant *Participant) StageVoid(stage *Stage) {
	participant.Stage(stage)
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
	__gong__stage(stage.ParticipantShapes, stage.ParticipantShape_stagedOrder, stage.ParticipantShape_orderStaged, &stage.ParticipantShapeOrder, stage.ParticipantShapes_mapString, participantshape, participantshape.Name)
	return participantshape
}

// StagePreserveOrder puts participantshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ParticipantShapeOrder
// - update stage.ParticipantShapeOrder accordingly
func (participantshape *ParticipantShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ParticipantShapes, stage.ParticipantShape_stagedOrder, stage.ParticipantShape_orderStaged, &stage.ParticipantShapeOrder, stage.ParticipantShapes_mapString, participantshape, order, participantshape.Name)
}

// Unstage removes participantshape off the model stage
func (participantshape *ParticipantShape) Unstage(stage *Stage) *ParticipantShape {
	__gong__unstage(stage.ParticipantShapes, stage.ParticipantShapes_mapString, participantshape, participantshape.Name)
	return participantshape
}

// UnstageVoid removes participantshape off the model stage
func (participantshape *ParticipantShape) UnstageVoid(stage *Stage) {
	participantshape.Unstage(stage)
}

func (participantshape *ParticipantShape) StageVoid(stage *Stage) {
	participantshape.Stage(stage)
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
	__gong__stage(stage.Processs, stage.Process_stagedOrder, stage.Process_orderStaged, &stage.ProcessOrder, stage.Processs_mapString, process, process.Name)
	return process
}

// StagePreserveOrder puts process to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProcessOrder
// - update stage.ProcessOrder accordingly
func (process *Process) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Processs, stage.Process_stagedOrder, stage.Process_orderStaged, &stage.ProcessOrder, stage.Processs_mapString, process, order, process.Name)
}

// Unstage removes process off the model stage
func (process *Process) Unstage(stage *Stage) *Process {
	__gong__unstage(stage.Processs, stage.Processs_mapString, process, process.Name)
	return process
}

// UnstageVoid removes process off the model stage
func (process *Process) UnstageVoid(stage *Stage) {
	process.Unstage(stage)
}

func (process *Process) StageVoid(stage *Stage) {
	process.Stage(stage)
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
	__gong__stage(stage.ProcessShapes, stage.ProcessShape_stagedOrder, stage.ProcessShape_orderStaged, &stage.ProcessShapeOrder, stage.ProcessShapes_mapString, processshape, processshape.Name)
	return processshape
}

// StagePreserveOrder puts processshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProcessShapeOrder
// - update stage.ProcessShapeOrder accordingly
func (processshape *ProcessShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.ProcessShapes, stage.ProcessShape_stagedOrder, stage.ProcessShape_orderStaged, &stage.ProcessShapeOrder, stage.ProcessShapes_mapString, processshape, order, processshape.Name)
}

// Unstage removes processshape off the model stage
func (processshape *ProcessShape) Unstage(stage *Stage) *ProcessShape {
	__gong__unstage(stage.ProcessShapes, stage.ProcessShapes_mapString, processshape, processshape.Name)
	return processshape
}

// UnstageVoid removes processshape off the model stage
func (processshape *ProcessShape) UnstageVoid(stage *Stage) {
	processshape.Unstage(stage)
}

func (processshape *ProcessShape) StageVoid(stage *Stage) {
	processshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (processshape *ProcessShape) GetName() (res string) {
	return processshape.Name
}

// for satisfaction of GongStruct interface
func (processshape *ProcessShape) SetName(name string) {
	processshape.Name = name
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
	__gong__resetStageType(&stage.AllocatedProcessShapes, &stage.AllocatedProcessShapes_mapString, &stage.AllocatedProcessShape_stagedOrder, &stage.AllocatedProcessShapeOrder)

	__gong__resetStageType(&stage.AllocatedResourceShapes, &stage.AllocatedResourceShapes_mapString, &stage.AllocatedResourceShape_stagedOrder, &stage.AllocatedResourceShapeOrder)

	__gong__resetStageType(&stage.ControlFlows, &stage.ControlFlows_mapString, &stage.ControlFlow_stagedOrder, &stage.ControlFlowOrder)

	__gong__resetStageType(&stage.ControlFlowShapes, &stage.ControlFlowShapes_mapString, &stage.ControlFlowShape_stagedOrder, &stage.ControlFlowShapeOrder)

	__gong__resetStageType(&stage.Datas, &stage.Datas_mapString, &stage.Data_stagedOrder, &stage.DataOrder)

	__gong__resetStageType(&stage.DataFlows, &stage.DataFlows_mapString, &stage.DataFlow_stagedOrder, &stage.DataFlowOrder)

	__gong__resetStageType(&stage.DataFlowShapes, &stage.DataFlowShapes_mapString, &stage.DataFlowShape_stagedOrder, &stage.DataFlowShapeOrder)

	__gong__resetStageType(&stage.DataShapes, &stage.DataShapes_mapString, &stage.DataShape_stagedOrder, &stage.DataShapeOrder)

	__gong__resetStageType(&stage.DiagramProcesss, &stage.DiagramProcesss_mapString, &stage.DiagramProcess_stagedOrder, &stage.DiagramProcessOrder)

	__gong__resetStageType(&stage.ExternalParticipantShapes, &stage.ExternalParticipantShapes_mapString, &stage.ExternalParticipantShape_stagedOrder, &stage.ExternalParticipantShapeOrder)

	__gong__resetStageType(&stage.Librarys, &stage.Librarys_mapString, &stage.Library_stagedOrder, &stage.LibraryOrder)

	__gong__resetStageType(&stage.Notes, &stage.Notes_mapString, &stage.Note_stagedOrder, &stage.NoteOrder)

	__gong__resetStageType(&stage.NoteShapes, &stage.NoteShapes_mapString, &stage.NoteShape_stagedOrder, &stage.NoteShapeOrder)

	__gong__resetStageType(&stage.NoteTaskShapes, &stage.NoteTaskShapes_mapString, &stage.NoteTaskShape_stagedOrder, &stage.NoteTaskShapeOrder)

	__gong__resetStageType(&stage.Participants, &stage.Participants_mapString, &stage.Participant_stagedOrder, &stage.ParticipantOrder)

	__gong__resetStageType(&stage.ParticipantShapes, &stage.ParticipantShapes_mapString, &stage.ParticipantShape_stagedOrder, &stage.ParticipantShapeOrder)

	__gong__resetStageType(&stage.Processs, &stage.Processs_mapString, &stage.Process_stagedOrder, &stage.ProcessOrder)

	__gong__resetStageType(&stage.ProcessShapes, &stage.ProcessShapes_mapString, &stage.ProcessShape_stagedOrder, &stage.ProcessShapeOrder)

	__gong__resetStageType(&stage.Resources, &stage.Resources_mapString, &stage.Resource_stagedOrder, &stage.ResourceOrder)

	__gong__resetStageType(&stage.Tasks, &stage.Tasks_mapString, &stage.Task_stagedOrder, &stage.TaskOrder)

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
	case *AllocatedProcessShape:
		return any(stage.AllocatedProcessShapes_mapString).(map[string]Type)
	case *AllocatedResourceShape:
		return any(stage.AllocatedResourceShapes_mapString).(map[string]Type)
	case *ControlFlow:
		return any(stage.ControlFlows_mapString).(map[string]Type)
	case *ControlFlowShape:
		return any(stage.ControlFlowShapes_mapString).(map[string]Type)
	case *Data:
		return any(stage.Datas_mapString).(map[string]Type)
	case *DataFlow:
		return any(stage.DataFlows_mapString).(map[string]Type)
	case *DataFlowShape:
		return any(stage.DataFlowShapes_mapString).(map[string]Type)
	case *DataShape:
		return any(stage.DataShapes_mapString).(map[string]Type)
	case *DiagramProcess:
		return any(stage.DiagramProcesss_mapString).(map[string]Type)
	case *ExternalParticipantShape:
		return any(stage.ExternalParticipantShapes_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Note:
		return any(stage.Notes_mapString).(map[string]Type)
	case *NoteShape:
		return any(stage.NoteShapes_mapString).(map[string]Type)
	case *NoteTaskShape:
		return any(stage.NoteTaskShapes_mapString).(map[string]Type)
	case *Participant:
		return any(stage.Participants_mapString).(map[string]Type)
	case *ParticipantShape:
		return any(stage.ParticipantShapes_mapString).(map[string]Type)
	case *Process:
		return any(stage.Processs_mapString).(map[string]Type)
	case *ProcessShape:
		return any(stage.ProcessShapes_mapString).(map[string]Type)
	case *Resource:
		return any(stage.Resources_mapString).(map[string]Type)
	case *Task:
		return any(stage.Tasks_mapString).(map[string]Type)
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
	case *AllocatedProcessShape:
		return any(&stage.AllocatedProcessShapes).(*map[Type]struct{})
	case *AllocatedResourceShape:
		return any(&stage.AllocatedResourceShapes).(*map[Type]struct{})
	case *ControlFlow:
		return any(&stage.ControlFlows).(*map[Type]struct{})
	case *ControlFlowShape:
		return any(&stage.ControlFlowShapes).(*map[Type]struct{})
	case *Data:
		return any(&stage.Datas).(*map[Type]struct{})
	case *DataFlow:
		return any(&stage.DataFlows).(*map[Type]struct{})
	case *DataFlowShape:
		return any(&stage.DataFlowShapes).(*map[Type]struct{})
	case *DataShape:
		return any(&stage.DataShapes).(*map[Type]struct{})
	case *DiagramProcess:
		return any(&stage.DiagramProcesss).(*map[Type]struct{})
	case *ExternalParticipantShape:
		return any(&stage.ExternalParticipantShapes).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Note:
		return any(&stage.Notes).(*map[Type]struct{})
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]struct{})
	case *NoteTaskShape:
		return any(&stage.NoteTaskShapes).(*map[Type]struct{})
	case *Participant:
		return any(&stage.Participants).(*map[Type]struct{})
	case *ParticipantShape:
		return any(&stage.ParticipantShapes).(*map[Type]struct{})
	case *Process:
		return any(&stage.Processs).(*map[Type]struct{})
	case *ProcessShape:
		return any(&stage.ProcessShapes).(*map[Type]struct{})
	case *Resource:
		return any(&stage.Resources).(*map[Type]struct{})
	case *Task:
		return any(&stage.Tasks).(*map[Type]struct{})
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
	case AllocatedProcessShape:
		return any(&AllocatedProcessShape{
			Participant: &Participant{Name: "Participant"},
			Process: &Process{Name: "Process"},
		}).(*Type)
	case AllocatedResourceShape:
		return any(&AllocatedResourceShape{
			Participant: &Participant{Name: "Participant"},
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
	case ControlFlow:
		return any(&ControlFlow{
			Start: &Task{Name: "Start"},
			End: &Task{Name: "End"},
		}).(*Type)
	case ControlFlowShape:
		return any(&ControlFlowShape{
			ControlFlow: &ControlFlow{Name: "ControlFlow"},
		}).(*Type)
	case DataFlow:
		return any(&DataFlow{
			Datas: []*Data{{Name: "Datas"}},
			StartTask: &Task{Name: "StartTask"},
			EndTask: &Task{Name: "EndTask"},
			StartExternalParticipant: &Participant{Name: "StartExternalParticipant"},
			EndExternalParticipant: &Participant{Name: "EndExternalParticipant"},
		}).(*Type)
	case DataFlowShape:
		return any(&DataFlowShape{
			DataFlow: &DataFlow{Name: "DataFlow"},
		}).(*Type)
	case DataShape:
		return any(&DataShape{
			Data: &Data{Name: "Data"},
			DataFlow: &DataFlow{Name: "DataFlow"},
		}).(*Type)
	case DiagramProcess:
		return any(&DiagramProcess{
			Process_Shapes: []*ProcessShape{{Name: "Process_Shapes"}},
			ProcesssWhoseNodeIsExpanded: []*Process{{Name: "ProcesssWhoseNodeIsExpanded"}},
			Participant_Shapes: []*ParticipantShape{{Name: "Participant_Shapes"}},
			ParticipantWhoseNodeIsExpanded: []*Participant{{Name: "ParticipantWhoseNodeIsExpanded"}},
			ExternalParticipant_Shapes: []*ExternalParticipantShape{{Name: "ExternalParticipant_Shapes"}},
			ExternalParticipantWhoseNodeIsExpanded: []*Participant{{Name: "ExternalParticipantWhoseNodeIsExpanded"}},
			ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded: []*Participant{{Name: "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded"}},
			ExternalParticipantsWhoseInDataFlowsNodeIsExpanded: []*Participant{{Name: "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded"}},
			TasksWhoseNodeIsExpanded: []*Task{{Name: "TasksWhoseNodeIsExpanded"}},
			Task_Shapes: []*TaskShape{{Name: "Task_Shapes"}},
			ControlFlowsWhoseNodeIsExpanded: []*ControlFlow{{Name: "ControlFlowsWhoseNodeIsExpanded"}},
			ControlFlow_Shapes: []*ControlFlowShape{{Name: "ControlFlow_Shapes"}},
			DataFlowsWhoseNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseNodeIsExpanded"}},
			DataFlow_Shapes: []*DataFlowShape{{Name: "DataFlow_Shapes"}},
			DatasWhoseNodeIsExpanded: []*Data{{Name: "DatasWhoseNodeIsExpanded"}},
			Data_Shapes: []*DataShape{{Name: "Data_Shapes"}},
			DataFlowsWhoseDataNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseDataNodeIsExpanded"}},
			AllocatedResourcesWhoseNodeIsExpanded: []*Resource{{Name: "AllocatedResourcesWhoseNodeIsExpanded"}},
			AllocatedResourceShapes: []*AllocatedResourceShape{{Name: "AllocatedResourceShapes"}},
			AllocatedProcessesWhoseNodeIsExpanded: []*Process{{Name: "AllocatedProcessesWhoseNodeIsExpanded"}},
			AllocatedProcessShapes: []*AllocatedProcessShape{{Name: "AllocatedProcessShapes"}},
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			NoteTaskShapes: []*NoteTaskShape{{Name: "NoteTaskShapes"}},
		}).(*Type)
	case ExternalParticipantShape:
		return any(&ExternalParticipantShape{
			Participant: &Participant{Name: "Participant"},
		}).(*Type)
	case Library:
		return any(&Library{
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
			RootProcesses: []*Process{{Name: "RootProcesses"}},
			ProcesssWhoseNodeIsExpanded: []*Process{{Name: "ProcesssWhoseNodeIsExpanded"}},
			RootDataFlows: []*DataFlow{{Name: "RootDataFlows"}},
			DataFlowsWhoseNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseNodeIsExpanded"}},
			RootDatas: []*Data{{Name: "RootDatas"}},
			DatasWhoseNodeIsExpanded: []*Data{{Name: "DatasWhoseNodeIsExpanded"}},
			RootResources: []*Resource{{Name: "RootResources"}},
			ResourcesWhoseNodeIsExpanded: []*Resource{{Name: "ResourcesWhoseNodeIsExpanded"}},
			ParticipantsWhoseNodeIsExpanded: []*Participant{{Name: "ParticipantsWhoseNodeIsExpanded"}},
			RootNotes: []*Note{{Name: "RootNotes"}},
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
		}).(*Type)
	case Note:
		return any(&Note{
			Tasks: []*Task{{Name: "Tasks"}},
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
	case Participant:
		return any(&Participant{
			Resources: []*Resource{{Name: "Resources"}},
			Processes: []*Process{{Name: "Processes"}},
			Tasks: []*Task{{Name: "Tasks"}},
			ControlFlows: []*ControlFlow{{Name: "ControlFlows"}},
			TaskWhoseOutControlFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseOutControlFlowsNodeIsExpanded"}},
			TaskWhoseInControlFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseInControlFlowsNodeIsExpanded"}},
			TaskWhoseOutDataFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseOutDataFlowsNodeIsExpanded"}},
			TaskWhoseInDataFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseInDataFlowsNodeIsExpanded"}},
		}).(*Type)
	case ParticipantShape:
		return any(&ParticipantShape{
			Participant: &Participant{Name: "Participant"},
		}).(*Type)
	case Process:
		return any(&Process{
			DiagramProcesss: []*DiagramProcess{{Name: "DiagramProcesss"}},
			DiagramProcessWhoseNodeIsExpanded: []*DiagramProcess{{Name: "DiagramProcessWhoseNodeIsExpanded"}},
			SubProcesses: []*Process{{Name: "SubProcesses"}},
			Participants: []*Participant{{Name: "Participants"}},
			ParticipantWhoseNodeIsExpanded: []*Participant{{Name: "ParticipantWhoseNodeIsExpanded"}},
			DataFlows: []*DataFlow{{Name: "DataFlows"}},
			ExternalParticipants: []*Participant{{Name: "ExternalParticipants"}},
			ExternalParticipantWhoseNodeIsExpanded: []*Participant{{Name: "ExternalParticipantWhoseNodeIsExpanded"}},
		}).(*Type)
	case ProcessShape:
		return any(&ProcessShape{
			Process: &Process{Name: "Process"},
		}).(*Type)
	case Task:
		return any(&Task{
			Type: &Process{Name: "Type"},
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
	// reverse maps of direct associations of AllocatedProcessShape
	case AllocatedProcessShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Participant":
			res := make(map[*Participant][]*AllocatedProcessShape)
			for allocatedprocessshape := range stage.AllocatedProcessShapes {
				if allocatedprocessshape.Participant != nil {
					participant_ := allocatedprocessshape.Participant
					var allocatedprocessshapes []*AllocatedProcessShape
					_, ok := res[participant_]
					if ok {
						allocatedprocessshapes = res[participant_]
					} else {
						allocatedprocessshapes = make([]*AllocatedProcessShape, 0)
					}
					allocatedprocessshapes = append(allocatedprocessshapes, allocatedprocessshape)
					res[participant_] = allocatedprocessshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Process":
			res := make(map[*Process][]*AllocatedProcessShape)
			for allocatedprocessshape := range stage.AllocatedProcessShapes {
				if allocatedprocessshape.Process != nil {
					process_ := allocatedprocessshape.Process
					var allocatedprocessshapes []*AllocatedProcessShape
					_, ok := res[process_]
					if ok {
						allocatedprocessshapes = res[process_]
					} else {
						allocatedprocessshapes = make([]*AllocatedProcessShape, 0)
					}
					allocatedprocessshapes = append(allocatedprocessshapes, allocatedprocessshape)
					res[process_] = allocatedprocessshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AllocatedResourceShape
	case AllocatedResourceShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Participant":
			res := make(map[*Participant][]*AllocatedResourceShape)
			for allocatedresourceshape := range stage.AllocatedResourceShapes {
				if allocatedresourceshape.Participant != nil {
					participant_ := allocatedresourceshape.Participant
					var allocatedresourceshapes []*AllocatedResourceShape
					_, ok := res[participant_]
					if ok {
						allocatedresourceshapes = res[participant_]
					} else {
						allocatedresourceshapes = make([]*AllocatedResourceShape, 0)
					}
					allocatedresourceshapes = append(allocatedresourceshapes, allocatedresourceshape)
					res[participant_] = allocatedresourceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Resource":
			res := make(map[*Resource][]*AllocatedResourceShape)
			for allocatedresourceshape := range stage.AllocatedResourceShapes {
				if allocatedresourceshape.Resource != nil {
					resource_ := allocatedresourceshape.Resource
					var allocatedresourceshapes []*AllocatedResourceShape
					_, ok := res[resource_]
					if ok {
						allocatedresourceshapes = res[resource_]
					} else {
						allocatedresourceshapes = make([]*AllocatedResourceShape, 0)
					}
					allocatedresourceshapes = append(allocatedresourceshapes, allocatedresourceshape)
					res[resource_] = allocatedresourceshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
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
	// reverse maps of direct associations of Data
	case Data:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DataFlow
	case DataFlow:
		switch fieldname {
		// insertion point for per direct association field
		case "StartTask":
			res := make(map[*Task][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.StartTask != nil {
					task_ := dataflow.StartTask
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
		case "EndTask":
			res := make(map[*Task][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.EndTask != nil {
					task_ := dataflow.EndTask
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
		case "StartExternalParticipant":
			res := make(map[*Participant][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.StartExternalParticipant != nil {
					participant_ := dataflow.StartExternalParticipant
					var dataflows []*DataFlow
					_, ok := res[participant_]
					if ok {
						dataflows = res[participant_]
					} else {
						dataflows = make([]*DataFlow, 0)
					}
					dataflows = append(dataflows, dataflow)
					res[participant_] = dataflows
				}
			}
			return any(res).(map[*End][]*Start)
		case "EndExternalParticipant":
			res := make(map[*Participant][]*DataFlow)
			for dataflow := range stage.DataFlows {
				if dataflow.EndExternalParticipant != nil {
					participant_ := dataflow.EndExternalParticipant
					var dataflows []*DataFlow
					_, ok := res[participant_]
					if ok {
						dataflows = res[participant_]
					} else {
						dataflows = make([]*DataFlow, 0)
					}
					dataflows = append(dataflows, dataflow)
					res[participant_] = dataflows
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
	// reverse maps of direct associations of DataShape
	case DataShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Data":
			res := make(map[*Data][]*DataShape)
			for datashape := range stage.DataShapes {
				if datashape.Data != nil {
					data_ := datashape.Data
					var datashapes []*DataShape
					_, ok := res[data_]
					if ok {
						datashapes = res[data_]
					} else {
						datashapes = make([]*DataShape, 0)
					}
					datashapes = append(datashapes, datashape)
					res[data_] = datashapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlow":
			res := make(map[*DataFlow][]*DataShape)
			for datashape := range stage.DataShapes {
				if datashape.DataFlow != nil {
					dataflow_ := datashape.DataFlow
					var datashapes []*DataShape
					_, ok := res[dataflow_]
					if ok {
						datashapes = res[dataflow_]
					} else {
						datashapes = make([]*DataShape, 0)
					}
					datashapes = append(datashapes, datashape)
					res[dataflow_] = datashapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DiagramProcess
	case DiagramProcess:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ExternalParticipantShape
	case ExternalParticipantShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Participant":
			res := make(map[*Participant][]*ExternalParticipantShape)
			for externalparticipantshape := range stage.ExternalParticipantShapes {
				if externalparticipantshape.Participant != nil {
					participant_ := externalparticipantshape.Participant
					var externalparticipantshapes []*ExternalParticipantShape
					_, ok := res[participant_]
					if ok {
						externalparticipantshapes = res[participant_]
					} else {
						externalparticipantshapes = make([]*ExternalParticipantShape, 0)
					}
					externalparticipantshapes = append(externalparticipantshapes, externalparticipantshape)
					res[participant_] = externalparticipantshapes
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
	// reverse maps of direct associations of Resource
	case Resource:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
		case "Type":
			res := make(map[*Process][]*Task)
			for task := range stage.Tasks {
				if task.Type != nil {
					process_ := task.Type
					var tasks []*Task
					_, ok := res[process_]
					if ok {
						tasks = res[process_]
					} else {
						tasks = make([]*Task, 0)
					}
					tasks = append(tasks, task)
					res[process_] = tasks
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
	// reverse maps of direct associations of AllocatedProcessShape
	case AllocatedProcessShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AllocatedResourceShape
	case AllocatedResourceShape:
		switch fieldname {
		// insertion point for per direct association field
		}
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
	// reverse maps of direct associations of Data
	case Data:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DataFlow
	case DataFlow:
		switch fieldname {
		// insertion point for per direct association field
		case "Datas":
			res := make(map[*Data][]*DataFlow)
			for dataflow := range stage.DataFlows {
				for _, data_ := range dataflow.Datas {
					res[data_] = append(res[data_], dataflow)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DataFlowShape
	case DataFlowShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DataShape
	case DataShape:
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
		case "ExternalParticipant_Shapes":
			res := make(map[*ExternalParticipantShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, externalparticipantshape_ := range diagramprocess.ExternalParticipant_Shapes {
					res[externalparticipantshape_] = append(res[externalparticipantshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalParticipantWhoseNodeIsExpanded":
			res := make(map[*Participant][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, participant_ := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
					res[participant_] = append(res[participant_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
			res := make(map[*Participant][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, participant_ := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
					res[participant_] = append(res[participant_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
			res := make(map[*Participant][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, participant_ := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
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
		case "Task_Shapes":
			res := make(map[*TaskShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, taskshape_ := range diagramprocess.Task_Shapes {
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
		case "ControlFlow_Shapes":
			res := make(map[*ControlFlowShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, controlflowshape_ := range diagramprocess.ControlFlow_Shapes {
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
		case "DataFlow_Shapes":
			res := make(map[*DataFlowShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, dataflowshape_ := range diagramprocess.DataFlow_Shapes {
					res[dataflowshape_] = append(res[dataflowshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DatasWhoseNodeIsExpanded":
			res := make(map[*Data][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, data_ := range diagramprocess.DatasWhoseNodeIsExpanded {
					res[data_] = append(res[data_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Data_Shapes":
			res := make(map[*DataShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, datashape_ := range diagramprocess.Data_Shapes {
					res[datashape_] = append(res[datashape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DataFlowsWhoseDataNodeIsExpanded":
			res := make(map[*DataFlow][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, dataflow_ := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
					res[dataflow_] = append(res[dataflow_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AllocatedResourcesWhoseNodeIsExpanded":
			res := make(map[*Resource][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, resource_ := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
					res[resource_] = append(res[resource_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AllocatedResourceShapes":
			res := make(map[*AllocatedResourceShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, allocatedresourceshape_ := range diagramprocess.AllocatedResourceShapes {
					res[allocatedresourceshape_] = append(res[allocatedresourceshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AllocatedProcessesWhoseNodeIsExpanded":
			res := make(map[*Process][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, process_ := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
					res[process_] = append(res[process_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "AllocatedProcessShapes":
			res := make(map[*AllocatedProcessShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, allocatedprocessshape_ := range diagramprocess.AllocatedProcessShapes {
					res[allocatedprocessshape_] = append(res[allocatedprocessshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Note_Shapes":
			res := make(map[*NoteShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, noteshape_ := range diagramprocess.Note_Shapes {
					res[noteshape_] = append(res[noteshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NotesWhoseNodeIsExpanded":
			res := make(map[*Note][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, note_ := range diagramprocess.NotesWhoseNodeIsExpanded {
					res[note_] = append(res[note_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteTaskShapes":
			res := make(map[*NoteTaskShape][]*DiagramProcess)
			for diagramprocess := range stage.DiagramProcesss {
				for _, notetaskshape_ := range diagramprocess.NoteTaskShapes {
					res[notetaskshape_] = append(res[notetaskshape_], diagramprocess)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ExternalParticipantShape
	case ExternalParticipantShape:
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
		case "SubLibrariesWhoseNodeIsExpanded":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibrariesWhoseNodeIsExpanded {
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
		case "RootDatas":
			res := make(map[*Data][]*Library)
			for library := range stage.Librarys {
				for _, data_ := range library.RootDatas {
					res[data_] = append(res[data_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "DatasWhoseNodeIsExpanded":
			res := make(map[*Data][]*Library)
			for library := range stage.Librarys {
				for _, data_ := range library.DatasWhoseNodeIsExpanded {
					res[data_] = append(res[data_], library)
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
		case "ResourcesWhoseNodeIsExpanded":
			res := make(map[*Resource][]*Library)
			for library := range stage.Librarys {
				for _, resource_ := range library.ResourcesWhoseNodeIsExpanded {
					res[resource_] = append(res[resource_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ParticipantsWhoseNodeIsExpanded":
			res := make(map[*Participant][]*Library)
			for library := range stage.Librarys {
				for _, participant_ := range library.ParticipantsWhoseNodeIsExpanded {
					res[participant_] = append(res[participant_], library)
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
		case "Tasks":
			res := make(map[*Task][]*Note)
			for note := range stage.Notes {
				for _, task_ := range note.Tasks {
					res[task_] = append(res[task_], note)
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of Participant
	case Participant:
		switch fieldname {
		// insertion point for per direct association field
		case "Resources":
			res := make(map[*Resource][]*Participant)
			for participant := range stage.Participants {
				for _, resource_ := range participant.Resources {
					res[resource_] = append(res[resource_], participant)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Processes":
			res := make(map[*Process][]*Participant)
			for participant := range stage.Participants {
				for _, process_ := range participant.Processes {
					res[process_] = append(res[process_], participant)
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "TaskWhoseOutDataFlowsNodeIsExpanded":
			res := make(map[*Task][]*Participant)
			for participant := range stage.Participants {
				for _, task_ := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
					res[task_] = append(res[task_], participant)
				}
			}
			return any(res).(map[*End][]*Start)
		case "TaskWhoseInDataFlowsNodeIsExpanded":
			res := make(map[*Task][]*Participant)
			for participant := range stage.Participants {
				for _, task_ := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
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
		case "DataFlows":
			res := make(map[*DataFlow][]*Process)
			for process := range stage.Processs {
				for _, dataflow_ := range process.DataFlows {
					res[dataflow_] = append(res[dataflow_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalParticipants":
			res := make(map[*Participant][]*Process)
			for process := range stage.Processs {
				for _, participant_ := range process.ExternalParticipants {
					res[participant_] = append(res[participant_], process)
				}
			}
			return any(res).(map[*End][]*Start)
		case "ExternalParticipantWhoseNodeIsExpanded":
			res := make(map[*Participant][]*Process)
			for process := range stage.Processs {
				for _, participant_ := range process.ExternalParticipantWhoseNodeIsExpanded {
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
	// reverse maps of direct associations of Resource
	case Resource:
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

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *AllocatedProcessShape:
		res = any(new(AllocatedProcessShape)).(Type)
	case *AllocatedResourceShape:
		res = any(new(AllocatedResourceShape)).(Type)
	case *ControlFlow:
		res = any(new(ControlFlow)).(Type)
	case *ControlFlowShape:
		res = any(new(ControlFlowShape)).(Type)
	case *Data:
		res = any(new(Data)).(Type)
	case *DataFlow:
		res = any(new(DataFlow)).(Type)
	case *DataFlowShape:
		res = any(new(DataFlowShape)).(Type)
	case *DataShape:
		res = any(new(DataShape)).(Type)
	case *DiagramProcess:
		res = any(new(DiagramProcess)).(Type)
	case *ExternalParticipantShape:
		res = any(new(ExternalParticipantShape)).(Type)
	case *Library:
		res = any(new(Library)).(Type)
	case *Note:
		res = any(new(Note)).(Type)
	case *NoteShape:
		res = any(new(NoteShape)).(Type)
	case *NoteTaskShape:
		res = any(new(NoteTaskShape)).(Type)
	case *Participant:
		res = any(new(Participant)).(Type)
	case *ParticipantShape:
		res = any(new(ParticipantShape)).(Type)
	case *Process:
		res = any(new(Process)).(Type)
	case *ProcessShape:
		res = any(new(ProcessShape)).(Type)
	case *Resource:
		res = any(new(Resource)).(Type)
	case *Task:
		res = any(new(Task)).(Type)
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
	case *AllocatedProcessShape:
		res = "AllocatedProcessShape"
	case *AllocatedResourceShape:
		res = "AllocatedResourceShape"
	case *ControlFlow:
		res = "ControlFlow"
	case *ControlFlowShape:
		res = "ControlFlowShape"
	case *Data:
		res = "Data"
	case *DataFlow:
		res = "DataFlow"
	case *DataFlowShape:
		res = "DataFlowShape"
	case *DataShape:
		res = "DataShape"
	case *DiagramProcess:
		res = "DiagramProcess"
	case *ExternalParticipantShape:
		res = "ExternalParticipantShape"
	case *Library:
		res = "Library"
	case *Note:
		res = "Note"
	case *NoteShape:
		res = "NoteShape"
	case *NoteTaskShape:
		res = "NoteTaskShape"
	case *Participant:
		res = "Participant"
	case *ParticipantShape:
		res = "ParticipantShape"
	case *Process:
		res = "Process"
	case *ProcessShape:
		res = "ProcessShape"
	case *Resource:
		res = "Resource"
	case *Task:
		res = "Task"
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
	case *AllocatedProcessShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "AllocatedProcessShapes"
		res = append(res, rf)
	case *AllocatedResourceShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "AllocatedResourceShapes"
		res = append(res, rf)
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
		rf.Fieldname = "ControlFlow_Shapes"
		res = append(res, rf)
	case *Data:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DataFlow"
		rf.Fieldname = "Datas"
		res = append(res, rf)
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "DatasWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootDatas"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "DatasWhoseNodeIsExpanded"
		res = append(res, rf)
	case *DataFlow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "DataFlowsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "DataFlowsWhoseDataNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootDataFlows"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "DataFlowsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "DataFlows"
		res = append(res, rf)
	case *DataFlowShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "DataFlow_Shapes"
		res = append(res, rf)
	case *DataShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "Data_Shapes"
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
	case *ExternalParticipantShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ExternalParticipant_Shapes"
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
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "NotesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootNotes"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "NotesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *NoteShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "Note_Shapes"
		res = append(res, rf)
	case *NoteTaskShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "NoteTaskShapes"
		res = append(res, rf)
	case *Participant:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ParticipantWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ExternalParticipantWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "ParticipantsWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "Participants"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "ParticipantWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "ExternalParticipants"
		res = append(res, rf)
		rf.GongstructName = "Process"
		rf.Fieldname = "ExternalParticipantWhoseNodeIsExpanded"
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
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "AllocatedProcessesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootProcesses"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "ProcesssWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Participant"
		rf.Fieldname = "Processes"
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
	case *Resource:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "AllocatedResourcesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootResources"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "ResourcesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Participant"
		rf.Fieldname = "Resources"
		res = append(res, rf)
	case *Task:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "TasksWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "Tasks"
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
		rf.GongstructName = "Participant"
		rf.Fieldname = "TaskWhoseOutDataFlowsNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Participant"
		rf.Fieldname = "TaskWhoseInDataFlowsNodeIsExpanded"
		res = append(res, rf)
	case *TaskShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "DiagramProcess"
		rf.Fieldname = "Task_Shapes"
		res = append(res, rf)
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (allocatedprocessshape *AllocatedProcessShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Process",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Process",
		},
	}
	return
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Resource",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Resource",
		},
	}
	return
}

func (controlflow *ControlFlow) GongGetFieldHeaders() (res []GongFieldHeader) {
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

func (data *Data) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Acronym",
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
			Name:               "SVG_Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InverseAppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			Name:                 "Datas",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Data",
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
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "DataFlowType",
		},
		{
			Name:                 "StartTask",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "EndTask",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "StartExternalParticipant",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Participant",
		},
		{
			Name:                 "EndExternalParticipant",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Participant",
		},
		{
			Name:               "IsDatasNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
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

func (datashape *DataShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Data",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Data",
		},
		{
			Name:                 "DataFlow",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "DataFlow",
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
			Name:                 "ExternalParticipant_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ExternalParticipantShape",
		},
		{
			Name:               "IsExternalParticipantsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ExternalParticipantWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Participant",
		},
		{
			Name:                 "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Participant",
		},
		{
			Name:                 "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Participant",
		},
		{
			Name:                 "TasksWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "Task_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "TaskShape",
		},
		{
			Name:                 "ControlFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlFlow",
		},
		{
			Name:                 "ControlFlow_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ControlFlowShape",
		},
		{
			Name:                 "DataFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:                 "DataFlow_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlowShape",
		},
		{
			Name:                 "DatasWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Data",
		},
		{
			Name:                 "Data_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataShape",
		},
		{
			Name:                 "DataFlowsWhoseDataNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:                 "AllocatedResourcesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
		{
			Name:                 "AllocatedResourceShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AllocatedResourceShape",
		},
		{
			Name:                 "AllocatedProcessesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:                 "AllocatedProcessShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "AllocatedProcessShape",
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
			Name:                 "NoteTaskShapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteTaskShape",
		},
	}
	return
}

func (externalparticipantshape *ExternalParticipantShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
		{
			Name:               "TailHeigth",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			Name:               "IsRootLibrary",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibraries",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
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
			Name:                 "RootProcesses",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:               "IsProcessesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
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
			Name:               "IsDataFlowsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "DataFlowsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:                 "RootDatas",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Data",
		},
		{
			Name:               "IsDatasNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "DatasWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Data",
		},
		{
			Name:                 "RootResources",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
		{
			Name:               "IsResourcesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ResourcesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
		{
			Name:                 "ParticipantsWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Participant",
		},
		{
			Name:                 "RootNotes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
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
			Name:               "IsTasksNodeExpanded",
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

func (participant *Participant) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsProcessResource",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Resources",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Resource",
		},
		{
			Name:               "IsResourcesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Processes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Process",
		},
		{
			Name:               "IsProcessesNodeExpanded",
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
		{
			Name:               "IsDataFlowsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "TaskWhoseOutDataFlowsNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "TaskWhoseInDataFlowsNodeIsExpanded",
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
		{
			Name:               "WidthWeight",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			Name:               "SVG_Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InverseAppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
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
		{
			Name:                 "DataFlows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DataFlow",
		},
		{
			Name:               "IsDataFlowsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "ExternalParticipants",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Participant",
		},
		{
			Name:                 "ExternalParticipantWhoseNodeIsExpanded",
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

func (resource *Resource) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Acronym",
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
			Name:               "SVG_Path",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "InverseAppliedScaling",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsStartTask",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEndTask",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Type",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Process",
		},
		{
			Name:               "IsTaskNameNotProcessName",
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
func (allocatedprocessshape *AllocatedProcessShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = allocatedprocessshape.Name
	case "Participant":
		res.GongFieldValueType = GongFieldValueTypePointer
		if allocatedprocessshape.Participant != nil {
			res.valueString = allocatedprocessshape.Participant.Name
			res.ids = allocatedprocessshape.Participant.GongGetUUID(stage)
		}
	case "Process":
		res.GongFieldValueType = GongFieldValueTypePointer
		if allocatedprocessshape.Process != nil {
			res.valueString = allocatedprocessshape.Process.Name
			res.ids = allocatedprocessshape.Process.GongGetUUID(stage)
		}
	}
	return
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = allocatedresourceshape.Name
	case "Participant":
		res.GongFieldValueType = GongFieldValueTypePointer
		if allocatedresourceshape.Participant != nil {
			res.valueString = allocatedresourceshape.Participant.Name
			res.ids = allocatedresourceshape.Participant.GongGetUUID(stage)
		}
	case "Resource":
		res.GongFieldValueType = GongFieldValueTypePointer
		if allocatedresourceshape.Resource != nil {
			res.valueString = allocatedresourceshape.Resource.Name
			res.ids = allocatedresourceshape.Resource.GongGetUUID(stage)
		}
	}
	return
}

func (controlflow *ControlFlow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = controlflow.Name
	case "Description":
		res.valueString = controlflow.Description
	case "ComputedPrefix":
		res.valueString = controlflow.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", controlflow.IsExpanded)
		res.valueBool = controlflow.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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

func (data *Data) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = data.Name
	case "Acronym":
		res.valueString = data.Acronym
	case "Description":
		res.valueString = data.Description
	case "ComputedPrefix":
		res.valueString = data.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", data.IsExpanded)
		res.valueBool = data.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SVG_Path":
		res.valueString = data.SVG_Path
	case "InverseAppliedScaling":
		res.valueString = fmt.Sprintf("%f", data.InverseAppliedScaling)
		res.valueFloat = data.InverseAppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (dataflow *DataFlow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = dataflow.Name
	case "Datas":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range dataflow.Datas {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Description":
		res.valueString = dataflow.Description
	case "ComputedPrefix":
		res.valueString = dataflow.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", dataflow.IsExpanded)
		res.valueBool = dataflow.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Type":
		enum := dataflow.Type
		res.valueString = enum.ToCodeString()
	case "StartTask":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.StartTask != nil {
			res.valueString = dataflow.StartTask.Name
			res.ids = dataflow.StartTask.GongGetUUID(stage)
		}
	case "EndTask":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.EndTask != nil {
			res.valueString = dataflow.EndTask.Name
			res.ids = dataflow.EndTask.GongGetUUID(stage)
		}
	case "StartExternalParticipant":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.StartExternalParticipant != nil {
			res.valueString = dataflow.StartExternalParticipant.Name
			res.ids = dataflow.StartExternalParticipant.GongGetUUID(stage)
		}
	case "EndExternalParticipant":
		res.GongFieldValueType = GongFieldValueTypePointer
		if dataflow.EndExternalParticipant != nil {
			res.valueString = dataflow.EndExternalParticipant.Name
			res.ids = dataflow.EndExternalParticipant.GongGetUUID(stage)
		}
	case "IsDatasNodeExpanded":
		res.valueString = fmt.Sprintf("%t", dataflow.IsDatasNodeExpanded)
		res.valueBool = dataflow.IsDatasNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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

func (datashape *DataShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = datashape.Name
	case "Data":
		res.GongFieldValueType = GongFieldValueTypePointer
		if datashape.Data != nil {
			res.valueString = datashape.Data.Name
			res.ids = datashape.Data.GongGetUUID(stage)
		}
	case "DataFlow":
		res.GongFieldValueType = GongFieldValueTypePointer
		if datashape.DataFlow != nil {
			res.valueString = datashape.DataFlow.Name
			res.ids = datashape.DataFlow.GongGetUUID(stage)
		}
	}
	return
}

func (diagramprocess *DiagramProcess) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagramprocess.Name
	case "Description":
		res.valueString = diagramprocess.Description
	case "ComputedPrefix":
		res.valueString = diagramprocess.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsExpanded)
		res.valueBool = diagramprocess.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "ExternalParticipant_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ExternalParticipant_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsExternalParticipantsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsExternalParticipantsNodeExpanded)
		res.valueBool = diagramprocess.IsExternalParticipantsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ExternalParticipantWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
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
	case "Task_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.Task_Shapes {
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
	case "ControlFlow_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.ControlFlow_Shapes {
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
	case "DataFlow_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.DataFlow_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DatasWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.DatasWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Data_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.Data_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "DataFlowsWhoseDataNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AllocatedResourcesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AllocatedResourceShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.AllocatedResourceShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AllocatedProcessesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "AllocatedProcessShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.AllocatedProcessShapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Note_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.Note_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NotesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.NotesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsNotesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagramprocess.IsNotesNodeExpanded)
		res.valueBool = diagramprocess.IsNotesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NoteTaskShapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagramprocess.NoteTaskShapes {
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

func (externalparticipantshape *ExternalParticipantShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = externalparticipantshape.Name
	case "Participant":
		res.GongFieldValueType = GongFieldValueTypePointer
		if externalparticipantshape.Participant != nil {
			res.valueString = externalparticipantshape.Participant.Name
			res.ids = externalparticipantshape.Participant.GongGetUUID(stage)
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", externalparticipantshape.IsExpanded)
		res.valueBool = externalparticipantshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "X":
		res.valueString = fmt.Sprintf("%f", externalparticipantshape.X)
		res.valueFloat = externalparticipantshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", externalparticipantshape.Y)
		res.valueFloat = externalparticipantshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", externalparticipantshape.Width)
		res.valueFloat = externalparticipantshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", externalparticipantshape.Height)
		res.valueFloat = externalparticipantshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", externalparticipantshape.IsHidden)
		res.valueBool = externalparticipantshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	case "TailHeigth":
		res.valueString = fmt.Sprintf("%f", externalparticipantshape.TailHeigth)
		res.valueFloat = externalparticipantshape.TailHeigth
		res.GongFieldValueType = GongFieldValueTypeFloat
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
	case "IsRootLibrary":
		res.valueString = fmt.Sprintf("%t", library.IsRootLibrary)
		res.valueBool = library.IsRootLibrary
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
	case "IsProcessesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsProcessesNodeExpanded)
		res.valueBool = library.IsProcessesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "IsDataFlowsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsDataFlowsNodeExpanded)
		res.valueBool = library.IsDataFlowsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "RootDatas":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootDatas {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDatasNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsDatasNodeExpanded)
		res.valueBool = library.IsDatasNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "DatasWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.DatasWhoseNodeIsExpanded {
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
	case "IsResourcesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsResourcesNodeExpanded)
		res.valueBool = library.IsResourcesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ResourcesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.ResourcesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ParticipantsWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.ParticipantsWhoseNodeIsExpanded {
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
	case "IsTasksNodeExpanded":
		res.valueString = fmt.Sprintf("%t", note.IsTasksNodeExpanded)
		res.valueBool = note.IsTasksNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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

func (participant *Participant) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = participant.Name
	case "IsProcessResource":
		res.valueString = fmt.Sprintf("%t", participant.IsProcessResource)
		res.valueBool = participant.IsProcessResource
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Description":
		res.valueString = participant.Description
	case "Resources":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range participant.Resources {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsResourcesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", participant.IsResourcesNodeExpanded)
		res.valueBool = participant.IsResourcesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Processes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range participant.Processes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsProcessesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", participant.IsProcessesNodeExpanded)
		res.valueBool = participant.IsProcessesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = participant.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", participant.IsExpanded)
		res.valueBool = participant.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "IsDataFlowsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", participant.IsDataFlowsNodeExpanded)
		res.valueBool = participant.IsDataFlowsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "TaskWhoseOutDataFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "TaskWhoseInDataFlowsNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
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
	case "WidthWeight":
		res.valueString = fmt.Sprintf("%f", participantshape.WidthWeight)
		res.valueFloat = participantshape.WidthWeight
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (process *Process) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = process.Name
	case "Description":
		res.valueString = process.Description
	case "ComputedPrefix":
		res.valueString = process.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", process.IsExpanded)
		res.valueBool = process.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SVG_Path":
		res.valueString = process.SVG_Path
	case "InverseAppliedScaling":
		res.valueString = fmt.Sprintf("%f", process.InverseAppliedScaling)
		res.valueFloat = process.InverseAppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
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
	case "DataFlows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.DataFlows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsDataFlowsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", process.IsDataFlowsNodeExpanded)
		res.valueBool = process.IsDataFlowsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ExternalParticipants":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.ExternalParticipants {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "ExternalParticipantWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range process.ExternalParticipantWhoseNodeIsExpanded {
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

func (resource *Resource) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = resource.Name
	case "Acronym":
		res.valueString = resource.Acronym
	case "Description":
		res.valueString = resource.Description
	case "ComputedPrefix":
		res.valueString = resource.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", resource.IsExpanded)
		res.valueBool = resource.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SVG_Path":
		res.valueString = resource.SVG_Path
	case "InverseAppliedScaling":
		res.valueString = fmt.Sprintf("%f", resource.InverseAppliedScaling)
		res.valueFloat = resource.InverseAppliedScaling
		res.GongFieldValueType = GongFieldValueTypeFloat
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
	case "ComputedPrefix":
		res.valueString = task.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", task.IsExpanded)
		res.valueBool = task.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsStartTask":
		res.valueString = fmt.Sprintf("%t", task.IsStartTask)
		res.valueBool = task.IsStartTask
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEndTask":
		res.valueString = fmt.Sprintf("%t", task.IsEndTask)
		res.valueBool = task.IsEndTask
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Type":
		res.GongFieldValueType = GongFieldValueTypePointer
		if task.Type != nil {
			res.valueString = task.Type.Name
			res.ids = task.Type.GongGetUUID(stage)
		}
	case "IsTaskNameNotProcessName":
		res.valueString = fmt.Sprintf("%t", task.IsTaskNameNotProcessName)
		res.valueBool = task.IsTaskNameNotProcessName
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (allocatedprocessshape *AllocatedProcessShape) GongGetGongstructName() string {
	return "AllocatedProcessShape"
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetGongstructName() string {
	return "AllocatedResourceShape"
}

func (controlflow *ControlFlow) GongGetGongstructName() string {
	return "ControlFlow"
}

func (controlflowshape *ControlFlowShape) GongGetGongstructName() string {
	return "ControlFlowShape"
}

func (data *Data) GongGetGongstructName() string {
	return "Data"
}

func (dataflow *DataFlow) GongGetGongstructName() string {
	return "DataFlow"
}

func (dataflowshape *DataFlowShape) GongGetGongstructName() string {
	return "DataFlowShape"
}

func (datashape *DataShape) GongGetGongstructName() string {
	return "DataShape"
}

func (diagramprocess *DiagramProcess) GongGetGongstructName() string {
	return "DiagramProcess"
}

func (externalparticipantshape *ExternalParticipantShape) GongGetGongstructName() string {
	return "ExternalParticipantShape"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (note *Note) GongGetGongstructName() string {
	return "Note"
}

func (noteshape *NoteShape) GongGetGongstructName() string {
	return "NoteShape"
}

func (notetaskshape *NoteTaskShape) GongGetGongstructName() string {
	return "NoteTaskShape"
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

func (resource *Resource) GongGetGongstructName() string {
	return "Resource"
}

func (task *Task) GongGetGongstructName() string {
	return "Task"
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
	__gong__rebuildMapString(stage.AllocatedProcessShapes, &stage.AllocatedProcessShapes_mapString)

	__gong__rebuildMapString(stage.AllocatedResourceShapes, &stage.AllocatedResourceShapes_mapString)

	__gong__rebuildMapString(stage.ControlFlows, &stage.ControlFlows_mapString)

	__gong__rebuildMapString(stage.ControlFlowShapes, &stage.ControlFlowShapes_mapString)

	__gong__rebuildMapString(stage.Datas, &stage.Datas_mapString)

	__gong__rebuildMapString(stage.DataFlows, &stage.DataFlows_mapString)

	__gong__rebuildMapString(stage.DataFlowShapes, &stage.DataFlowShapes_mapString)

	__gong__rebuildMapString(stage.DataShapes, &stage.DataShapes_mapString)

	__gong__rebuildMapString(stage.DiagramProcesss, &stage.DiagramProcesss_mapString)

	__gong__rebuildMapString(stage.ExternalParticipantShapes, &stage.ExternalParticipantShapes_mapString)

	__gong__rebuildMapString(stage.Librarys, &stage.Librarys_mapString)

	__gong__rebuildMapString(stage.Notes, &stage.Notes_mapString)

	__gong__rebuildMapString(stage.NoteShapes, &stage.NoteShapes_mapString)

	__gong__rebuildMapString(stage.NoteTaskShapes, &stage.NoteTaskShapes_mapString)

	__gong__rebuildMapString(stage.Participants, &stage.Participants_mapString)

	__gong__rebuildMapString(stage.ParticipantShapes, &stage.ParticipantShapes_mapString)

	__gong__rebuildMapString(stage.Processs, &stage.Processs_mapString)

	__gong__rebuildMapString(stage.ProcessShapes, &stage.ProcessShapes_mapString)

	__gong__rebuildMapString(stage.Resources, &stage.Resources_mapString)

	__gong__rebuildMapString(stage.Tasks, &stage.Tasks_mapString)

	__gong__rebuildMapString(stage.TaskShapes, &stage.TaskShapes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
