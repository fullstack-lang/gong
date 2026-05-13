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
	AllocatedProcessShapes                map[*AllocatedProcessShape]struct{}
	AllocatedProcessShapes_instance       map[*AllocatedProcessShape]*AllocatedProcessShape
	AllocatedProcessShapes_mapString      map[string]*AllocatedProcessShape
	AllocatedProcessShapeOrder            uint
	AllocatedProcessShape_stagedOrder     map[*AllocatedProcessShape]uint
	AllocatedProcessShape_orderStaged     map[uint]*AllocatedProcessShape
	AllocatedProcessShapes_reference      map[*AllocatedProcessShape]*AllocatedProcessShape
	AllocatedProcessShapes_referenceOrder map[*AllocatedProcessShape]uint

	// insertion point for slice of pointers maps
	OnAfterAllocatedProcessShapeCreateCallback OnAfterCreateInterface[AllocatedProcessShape]
	OnAfterAllocatedProcessShapeUpdateCallback OnAfterUpdateInterface[AllocatedProcessShape]
	OnAfterAllocatedProcessShapeDeleteCallback OnAfterDeleteInterface[AllocatedProcessShape]
	OnAfterAllocatedProcessShapeReadCallback   OnAfterReadInterface[AllocatedProcessShape]

	AllocatedResourceShapes                map[*AllocatedResourceShape]struct{}
	AllocatedResourceShapes_instance       map[*AllocatedResourceShape]*AllocatedResourceShape
	AllocatedResourceShapes_mapString      map[string]*AllocatedResourceShape
	AllocatedResourceShapeOrder            uint
	AllocatedResourceShape_stagedOrder     map[*AllocatedResourceShape]uint
	AllocatedResourceShape_orderStaged     map[uint]*AllocatedResourceShape
	AllocatedResourceShapes_reference      map[*AllocatedResourceShape]*AllocatedResourceShape
	AllocatedResourceShapes_referenceOrder map[*AllocatedResourceShape]uint

	// insertion point for slice of pointers maps
	OnAfterAllocatedResourceShapeCreateCallback OnAfterCreateInterface[AllocatedResourceShape]
	OnAfterAllocatedResourceShapeUpdateCallback OnAfterUpdateInterface[AllocatedResourceShape]
	OnAfterAllocatedResourceShapeDeleteCallback OnAfterDeleteInterface[AllocatedResourceShape]
	OnAfterAllocatedResourceShapeReadCallback   OnAfterReadInterface[AllocatedResourceShape]

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

	Datas                map[*Data]struct{}
	Datas_instance       map[*Data]*Data
	Datas_mapString      map[string]*Data
	DataOrder            uint
	Data_stagedOrder     map[*Data]uint
	Data_orderStaged     map[uint]*Data
	Datas_reference      map[*Data]*Data
	Datas_referenceOrder map[*Data]uint

	// insertion point for slice of pointers maps
	OnAfterDataCreateCallback OnAfterCreateInterface[Data]
	OnAfterDataUpdateCallback OnAfterUpdateInterface[Data]
	OnAfterDataDeleteCallback OnAfterDeleteInterface[Data]
	OnAfterDataReadCallback   OnAfterReadInterface[Data]

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

	DataShapes                map[*DataShape]struct{}
	DataShapes_instance       map[*DataShape]*DataShape
	DataShapes_mapString      map[string]*DataShape
	DataShapeOrder            uint
	DataShape_stagedOrder     map[*DataShape]uint
	DataShape_orderStaged     map[uint]*DataShape
	DataShapes_reference      map[*DataShape]*DataShape
	DataShapes_referenceOrder map[*DataShape]uint

	// insertion point for slice of pointers maps
	OnAfterDataShapeCreateCallback OnAfterCreateInterface[DataShape]
	OnAfterDataShapeUpdateCallback OnAfterUpdateInterface[DataShape]
	OnAfterDataShapeDeleteCallback OnAfterDeleteInterface[DataShape]
	OnAfterDataShapeReadCallback   OnAfterReadInterface[DataShape]

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

	OnAfterDiagramProcessCreateCallback OnAfterCreateInterface[DiagramProcess]
	OnAfterDiagramProcessUpdateCallback OnAfterUpdateInterface[DiagramProcess]
	OnAfterDiagramProcessDeleteCallback OnAfterDeleteInterface[DiagramProcess]
	OnAfterDiagramProcessReadCallback   OnAfterReadInterface[DiagramProcess]

	ExternalParticipantShapes                map[*ExternalParticipantShape]struct{}
	ExternalParticipantShapes_instance       map[*ExternalParticipantShape]*ExternalParticipantShape
	ExternalParticipantShapes_mapString      map[string]*ExternalParticipantShape
	ExternalParticipantShapeOrder            uint
	ExternalParticipantShape_stagedOrder     map[*ExternalParticipantShape]uint
	ExternalParticipantShape_orderStaged     map[uint]*ExternalParticipantShape
	ExternalParticipantShapes_reference      map[*ExternalParticipantShape]*ExternalParticipantShape
	ExternalParticipantShapes_referenceOrder map[*ExternalParticipantShape]uint

	// insertion point for slice of pointers maps
	OnAfterExternalParticipantShapeCreateCallback OnAfterCreateInterface[ExternalParticipantShape]
	OnAfterExternalParticipantShapeUpdateCallback OnAfterUpdateInterface[ExternalParticipantShape]
	OnAfterExternalParticipantShapeDeleteCallback OnAfterDeleteInterface[ExternalParticipantShape]
	OnAfterExternalParticipantShapeReadCallback   OnAfterReadInterface[ExternalParticipantShape]

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

	Library_RootNotes_reverseMap map[*Note]*Library

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
	Note_Tasks_reverseMap map[*Task]*Note

	OnAfterNoteCreateCallback OnAfterCreateInterface[Note]
	OnAfterNoteUpdateCallback OnAfterUpdateInterface[Note]
	OnAfterNoteDeleteCallback OnAfterDeleteInterface[Note]
	OnAfterNoteReadCallback   OnAfterReadInterface[Note]

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

	Process_DataFlows_reverseMap map[*DataFlow]*Process

	Process_ExternalParticipants_reverseMap map[*Participant]*Process

	Process_ExternalParticipantWhoseNodeIsExpanded_reverseMap map[*Participant]*Process

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

	Resources                map[*Resource]struct{}
	Resources_instance       map[*Resource]*Resource
	Resources_mapString      map[string]*Resource
	ResourceOrder            uint
	Resource_stagedOrder     map[*Resource]uint
	Resource_orderStaged     map[uint]*Resource
	Resources_reference      map[*Resource]*Resource
	Resources_referenceOrder map[*Resource]uint

	// insertion point for slice of pointers maps
	OnAfterResourceCreateCallback OnAfterCreateInterface[Resource]
	OnAfterResourceUpdateCallback OnAfterUpdateInterface[Resource]
	OnAfterResourceDeleteCallback OnAfterDeleteInterface[Resource]
	OnAfterResourceReadCallback   OnAfterReadInterface[Resource]

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
	stage.AllocatedProcessShapes_reference = make(map[*AllocatedProcessShape]*AllocatedProcessShape)
	stage.AllocatedProcessShapes_instance = make(map[*AllocatedProcessShape]*AllocatedProcessShape)
	stage.AllocatedProcessShapes_referenceOrder = make(map[*AllocatedProcessShape]uint)

	stage.AllocatedResourceShapes_reference = make(map[*AllocatedResourceShape]*AllocatedResourceShape)
	stage.AllocatedResourceShapes_instance = make(map[*AllocatedResourceShape]*AllocatedResourceShape)
	stage.AllocatedResourceShapes_referenceOrder = make(map[*AllocatedResourceShape]uint)

	stage.ControlFlows_reference = make(map[*ControlFlow]*ControlFlow)
	stage.ControlFlows_instance = make(map[*ControlFlow]*ControlFlow)
	stage.ControlFlows_referenceOrder = make(map[*ControlFlow]uint)

	stage.ControlFlowShapes_reference = make(map[*ControlFlowShape]*ControlFlowShape)
	stage.ControlFlowShapes_instance = make(map[*ControlFlowShape]*ControlFlowShape)
	stage.ControlFlowShapes_referenceOrder = make(map[*ControlFlowShape]uint)

	stage.Datas_reference = make(map[*Data]*Data)
	stage.Datas_instance = make(map[*Data]*Data)
	stage.Datas_referenceOrder = make(map[*Data]uint)

	stage.DataFlows_reference = make(map[*DataFlow]*DataFlow)
	stage.DataFlows_instance = make(map[*DataFlow]*DataFlow)
	stage.DataFlows_referenceOrder = make(map[*DataFlow]uint)

	stage.DataFlowShapes_reference = make(map[*DataFlowShape]*DataFlowShape)
	stage.DataFlowShapes_instance = make(map[*DataFlowShape]*DataFlowShape)
	stage.DataFlowShapes_referenceOrder = make(map[*DataFlowShape]uint)

	stage.DataShapes_reference = make(map[*DataShape]*DataShape)
	stage.DataShapes_instance = make(map[*DataShape]*DataShape)
	stage.DataShapes_referenceOrder = make(map[*DataShape]uint)

	stage.DiagramProcesss_reference = make(map[*DiagramProcess]*DiagramProcess)
	stage.DiagramProcesss_instance = make(map[*DiagramProcess]*DiagramProcess)
	stage.DiagramProcesss_referenceOrder = make(map[*DiagramProcess]uint)

	stage.ExternalParticipantShapes_reference = make(map[*ExternalParticipantShape]*ExternalParticipantShape)
	stage.ExternalParticipantShapes_instance = make(map[*ExternalParticipantShape]*ExternalParticipantShape)
	stage.ExternalParticipantShapes_referenceOrder = make(map[*ExternalParticipantShape]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_instance = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint)

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)

	stage.NoteTaskShapes_reference = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_instance = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint)

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

	stage.Resources_reference = make(map[*Resource]*Resource)
	stage.Resources_instance = make(map[*Resource]*Resource)
	stage.Resources_referenceOrder = make(map[*Resource]uint)

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
	var maxAllocatedProcessShapeOrder uint
	var foundAllocatedProcessShape bool
	for _, order := range stage.AllocatedProcessShape_stagedOrder {
		if !foundAllocatedProcessShape || order > maxAllocatedProcessShapeOrder {
			maxAllocatedProcessShapeOrder = order
			foundAllocatedProcessShape = true
		}
	}
	if foundAllocatedProcessShape {
		stage.AllocatedProcessShapeOrder = maxAllocatedProcessShapeOrder + 1
	} else {
		stage.AllocatedProcessShapeOrder = 0
	}

	var maxAllocatedResourceShapeOrder uint
	var foundAllocatedResourceShape bool
	for _, order := range stage.AllocatedResourceShape_stagedOrder {
		if !foundAllocatedResourceShape || order > maxAllocatedResourceShapeOrder {
			maxAllocatedResourceShapeOrder = order
			foundAllocatedResourceShape = true
		}
	}
	if foundAllocatedResourceShape {
		stage.AllocatedResourceShapeOrder = maxAllocatedResourceShapeOrder + 1
	} else {
		stage.AllocatedResourceShapeOrder = 0
	}

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

	var maxDataOrder uint
	var foundData bool
	for _, order := range stage.Data_stagedOrder {
		if !foundData || order > maxDataOrder {
			maxDataOrder = order
			foundData = true
		}
	}
	if foundData {
		stage.DataOrder = maxDataOrder + 1
	} else {
		stage.DataOrder = 0
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

	var maxDataShapeOrder uint
	var foundDataShape bool
	for _, order := range stage.DataShape_stagedOrder {
		if !foundDataShape || order > maxDataShapeOrder {
			maxDataShapeOrder = order
			foundDataShape = true
		}
	}
	if foundDataShape {
		stage.DataShapeOrder = maxDataShapeOrder + 1
	} else {
		stage.DataShapeOrder = 0
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

	var maxExternalParticipantShapeOrder uint
	var foundExternalParticipantShape bool
	for _, order := range stage.ExternalParticipantShape_stagedOrder {
		if !foundExternalParticipantShape || order > maxExternalParticipantShapeOrder {
			maxExternalParticipantShapeOrder = order
			foundExternalParticipantShape = true
		}
	}
	if foundExternalParticipantShape {
		stage.ExternalParticipantShapeOrder = maxExternalParticipantShapeOrder + 1
	} else {
		stage.ExternalParticipantShapeOrder = 0
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
	case *AllocatedProcessShape:
		tmp := GetStructInstancesByOrder(stage.AllocatedProcessShapes, stage.AllocatedProcessShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AllocatedProcessShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *AllocatedResourceShape:
		tmp := GetStructInstancesByOrder(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *AllocatedResourceShape implements.
			res = append(res, any(v).(T))
		}
		return res
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
	case *Data:
		tmp := GetStructInstancesByOrder(stage.Datas, stage.Data_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Data implements.
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
	case *DataShape:
		tmp := GetStructInstancesByOrder(stage.DataShapes, stage.DataShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DataShape implements.
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
	case *ExternalParticipantShape:
		tmp := GetStructInstancesByOrder(stage.ExternalParticipantShapes, stage.ExternalParticipantShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ExternalParticipantShape implements.
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
	case "AllocatedProcessShape":
		res = GetNamedStructInstances(stage.AllocatedProcessShapes, stage.AllocatedProcessShape_stagedOrder)
	case "AllocatedResourceShape":
		res = GetNamedStructInstances(stage.AllocatedResourceShapes, stage.AllocatedResourceShape_stagedOrder)
	case "ControlFlow":
		res = GetNamedStructInstances(stage.ControlFlows, stage.ControlFlow_stagedOrder)
	case "ControlFlowShape":
		res = GetNamedStructInstances(stage.ControlFlowShapes, stage.ControlFlowShape_stagedOrder)
	case "Data":
		res = GetNamedStructInstances(stage.Datas, stage.Data_stagedOrder)
	case "DataFlow":
		res = GetNamedStructInstances(stage.DataFlows, stage.DataFlow_stagedOrder)
	case "DataFlowShape":
		res = GetNamedStructInstances(stage.DataFlowShapes, stage.DataFlowShape_stagedOrder)
	case "DataShape":
		res = GetNamedStructInstances(stage.DataShapes, stage.DataShape_stagedOrder)
	case "DiagramProcess":
		res = GetNamedStructInstances(stage.DiagramProcesss, stage.DiagramProcess_stagedOrder)
	case "ExternalParticipantShape":
		res = GetNamedStructInstances(stage.ExternalParticipantShapes, stage.ExternalParticipantShape_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Note":
		res = GetNamedStructInstances(stage.Notes, stage.Note_stagedOrder)
	case "NoteShape":
		res = GetNamedStructInstances(stage.NoteShapes, stage.NoteShape_stagedOrder)
	case "NoteTaskShape":
		res = GetNamedStructInstances(stage.NoteTaskShapes, stage.NoteTaskShape_stagedOrder)
	case "Participant":
		res = GetNamedStructInstances(stage.Participants, stage.Participant_stagedOrder)
	case "ParticipantShape":
		res = GetNamedStructInstances(stage.ParticipantShapes, stage.ParticipantShape_stagedOrder)
	case "Process":
		res = GetNamedStructInstances(stage.Processs, stage.Process_stagedOrder)
	case "ProcessShape":
		res = GetNamedStructInstances(stage.ProcessShapes, stage.ProcessShape_stagedOrder)
	case "Resource":
		res = GetNamedStructInstances(stage.Resources, stage.Resource_stagedOrder)
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
	CommitAllocatedProcessShape(allocatedprocessshape *AllocatedProcessShape)
	CheckoutAllocatedProcessShape(allocatedprocessshape *AllocatedProcessShape)
	CommitAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape)
	CheckoutAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape)
	CommitControlFlow(controlflow *ControlFlow)
	CheckoutControlFlow(controlflow *ControlFlow)
	CommitControlFlowShape(controlflowshape *ControlFlowShape)
	CheckoutControlFlowShape(controlflowshape *ControlFlowShape)
	CommitData(data *Data)
	CheckoutData(data *Data)
	CommitDataFlow(dataflow *DataFlow)
	CheckoutDataFlow(dataflow *DataFlow)
	CommitDataFlowShape(dataflowshape *DataFlowShape)
	CheckoutDataFlowShape(dataflowshape *DataFlowShape)
	CommitDataShape(datashape *DataShape)
	CheckoutDataShape(datashape *DataShape)
	CommitDiagramProcess(diagramprocess *DiagramProcess)
	CheckoutDiagramProcess(diagramprocess *DiagramProcess)
	CommitExternalParticipantShape(externalparticipantshape *ExternalParticipantShape)
	CheckoutExternalParticipantShape(externalparticipantshape *ExternalParticipantShape)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitNote(note *Note)
	CheckoutNote(note *Note)
	CommitNoteShape(noteshape *NoteShape)
	CheckoutNoteShape(noteshape *NoteShape)
	CommitNoteTaskShape(notetaskshape *NoteTaskShape)
	CheckoutNoteTaskShape(notetaskshape *NoteTaskShape)
	CommitParticipant(participant *Participant)
	CheckoutParticipant(participant *Participant)
	CommitParticipantShape(participantshape *ParticipantShape)
	CheckoutParticipantShape(participantshape *ParticipantShape)
	CommitProcess(process *Process)
	CheckoutProcess(process *Process)
	CommitProcessShape(processshape *ProcessShape)
	CheckoutProcessShape(processshape *ProcessShape)
	CommitResource(resource *Resource)
	CheckoutResource(resource *Resource)
	CommitTask(task *Task)
	CheckoutTask(task *Task)
	CommitTaskShape(taskshape *TaskShape)
	CheckoutTaskShape(taskshape *TaskShape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "AllocatedProcessShape"},
			{name: "AllocatedResourceShape"},
			{name: "ControlFlow"},
			{name: "ControlFlowShape"},
			{name: "Data"},
			{name: "DataFlow"},
			{name: "DataFlowShape"},
			{name: "DataShape"},
			{name: "DiagramProcess"},
			{name: "ExternalParticipantShape"},
			{name: "Library"},
			{name: "Note"},
			{name: "NoteShape"},
			{name: "NoteTaskShape"},
			{name: "Participant"},
			{name: "ParticipantShape"},
			{name: "Process"},
			{name: "ProcessShape"},
			{name: "Resource"},
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
	case *AllocatedProcessShape:
		return stage.AllocatedProcessShape_stagedOrder[instance]
	case *AllocatedResourceShape:
		return stage.AllocatedResourceShape_stagedOrder[instance]
	case *ControlFlow:
		return stage.ControlFlow_stagedOrder[instance]
	case *ControlFlowShape:
		return stage.ControlFlowShape_stagedOrder[instance]
	case *Data:
		return stage.Data_stagedOrder[instance]
	case *DataFlow:
		return stage.DataFlow_stagedOrder[instance]
	case *DataFlowShape:
		return stage.DataFlowShape_stagedOrder[instance]
	case *DataShape:
		return stage.DataShape_stagedOrder[instance]
	case *DiagramProcess:
		return stage.DiagramProcess_stagedOrder[instance]
	case *ExternalParticipantShape:
		return stage.ExternalParticipantShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *NoteTaskShape:
		return stage.NoteTaskShape_stagedOrder[instance]
	case *Participant:
		return stage.Participant_stagedOrder[instance]
	case *ParticipantShape:
		return stage.ParticipantShape_stagedOrder[instance]
	case *Process:
		return stage.Process_stagedOrder[instance]
	case *ProcessShape:
		return stage.ProcessShape_stagedOrder[instance]
	case *Resource:
		return stage.Resource_stagedOrder[instance]
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

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *AllocatedProcessShape:
		return stage.AllocatedProcessShape_stagedOrder[instance]
	case *AllocatedResourceShape:
		return stage.AllocatedResourceShape_stagedOrder[instance]
	case *ControlFlow:
		return stage.ControlFlow_stagedOrder[instance]
	case *ControlFlowShape:
		return stage.ControlFlowShape_stagedOrder[instance]
	case *Data:
		return stage.Data_stagedOrder[instance]
	case *DataFlow:
		return stage.DataFlow_stagedOrder[instance]
	case *DataFlowShape:
		return stage.DataFlowShape_stagedOrder[instance]
	case *DataShape:
		return stage.DataShape_stagedOrder[instance]
	case *DiagramProcess:
		return stage.DiagramProcess_stagedOrder[instance]
	case *ExternalParticipantShape:
		return stage.ExternalParticipantShape_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *NoteTaskShape:
		return stage.NoteTaskShape_stagedOrder[instance]
	case *Participant:
		return stage.Participant_stagedOrder[instance]
	case *ParticipantShape:
		return stage.ParticipantShape_stagedOrder[instance]
	case *Process:
		return stage.Process_stagedOrder[instance]
	case *ProcessShape:
		return stage.ProcessShape_stagedOrder[instance]
	case *Resource:
		return stage.Resource_stagedOrder[instance]
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
	if _, ok := stage.AllocatedProcessShapes[allocatedprocessshape]; !ok {
		stage.AllocatedProcessShapes[allocatedprocessshape] = struct{}{}
		stage.AllocatedProcessShape_stagedOrder[allocatedprocessshape] = stage.AllocatedProcessShapeOrder
		stage.AllocatedProcessShape_orderStaged[stage.AllocatedProcessShapeOrder] = allocatedprocessshape
		stage.AllocatedProcessShapeOrder++
	}
	stage.AllocatedProcessShapes_mapString[allocatedprocessshape.Name] = allocatedprocessshape

	return allocatedprocessshape
}

// StagePreserveOrder puts allocatedprocessshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AllocatedProcessShapeOrder
// - update stage.AllocatedProcessShapeOrder accordingly
func (allocatedprocessshape *AllocatedProcessShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AllocatedProcessShapes[allocatedprocessshape]; !ok {
		stage.AllocatedProcessShapes[allocatedprocessshape] = struct{}{}

		if order > stage.AllocatedProcessShapeOrder {
			stage.AllocatedProcessShapeOrder = order
		}
		stage.AllocatedProcessShape_stagedOrder[allocatedprocessshape] = order
		stage.AllocatedProcessShape_orderStaged[order] = allocatedprocessshape
		stage.AllocatedProcessShapeOrder++
	}
	stage.AllocatedProcessShapes_mapString[allocatedprocessshape.Name] = allocatedprocessshape
}

// Unstage removes allocatedprocessshape off the model stage
func (allocatedprocessshape *AllocatedProcessShape) Unstage(stage *Stage) *AllocatedProcessShape {
	delete(stage.AllocatedProcessShapes, allocatedprocessshape)
	// issue1150
	// delete(stage.AllocatedProcessShape_stagedOrder, allocatedprocessshape)
	delete(stage.AllocatedProcessShapes_mapString, allocatedprocessshape.Name)

	return allocatedprocessshape
}

// UnstageVoid removes allocatedprocessshape off the model stage
func (allocatedprocessshape *AllocatedProcessShape) UnstageVoid(stage *Stage) {
	delete(stage.AllocatedProcessShapes, allocatedprocessshape)
	// issue1150
	// delete(stage.AllocatedProcessShape_stagedOrder, allocatedprocessshape)
	delete(stage.AllocatedProcessShapes_mapString, allocatedprocessshape.Name)
}

// commit allocatedprocessshape to the back repo (if it is already staged)
func (allocatedprocessshape *AllocatedProcessShape) Commit(stage *Stage) *AllocatedProcessShape {
	if _, ok := stage.AllocatedProcessShapes[allocatedprocessshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAllocatedProcessShape(allocatedprocessshape)
		}
	}
	return allocatedprocessshape
}

func (allocatedprocessshape *AllocatedProcessShape) CommitVoid(stage *Stage) {
	allocatedprocessshape.Commit(stage)
}

func (allocatedprocessshape *AllocatedProcessShape) StageVoid(stage *Stage) {
	allocatedprocessshape.Stage(stage)
}

// Checkout allocatedprocessshape to the back repo (if it is already staged)
func (allocatedprocessshape *AllocatedProcessShape) Checkout(stage *Stage) *AllocatedProcessShape {
	if _, ok := stage.AllocatedProcessShapes[allocatedprocessshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAllocatedProcessShape(allocatedprocessshape)
		}
	}
	return allocatedprocessshape
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
	if _, ok := stage.AllocatedResourceShapes[allocatedresourceshape]; !ok {
		stage.AllocatedResourceShapes[allocatedresourceshape] = struct{}{}
		stage.AllocatedResourceShape_stagedOrder[allocatedresourceshape] = stage.AllocatedResourceShapeOrder
		stage.AllocatedResourceShape_orderStaged[stage.AllocatedResourceShapeOrder] = allocatedresourceshape
		stage.AllocatedResourceShapeOrder++
	}
	stage.AllocatedResourceShapes_mapString[allocatedresourceshape.Name] = allocatedresourceshape

	return allocatedresourceshape
}

// StagePreserveOrder puts allocatedresourceshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.AllocatedResourceShapeOrder
// - update stage.AllocatedResourceShapeOrder accordingly
func (allocatedresourceshape *AllocatedResourceShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.AllocatedResourceShapes[allocatedresourceshape]; !ok {
		stage.AllocatedResourceShapes[allocatedresourceshape] = struct{}{}

		if order > stage.AllocatedResourceShapeOrder {
			stage.AllocatedResourceShapeOrder = order
		}
		stage.AllocatedResourceShape_stagedOrder[allocatedresourceshape] = order
		stage.AllocatedResourceShape_orderStaged[order] = allocatedresourceshape
		stage.AllocatedResourceShapeOrder++
	}
	stage.AllocatedResourceShapes_mapString[allocatedresourceshape.Name] = allocatedresourceshape
}

// Unstage removes allocatedresourceshape off the model stage
func (allocatedresourceshape *AllocatedResourceShape) Unstage(stage *Stage) *AllocatedResourceShape {
	delete(stage.AllocatedResourceShapes, allocatedresourceshape)
	// issue1150
	// delete(stage.AllocatedResourceShape_stagedOrder, allocatedresourceshape)
	delete(stage.AllocatedResourceShapes_mapString, allocatedresourceshape.Name)

	return allocatedresourceshape
}

// UnstageVoid removes allocatedresourceshape off the model stage
func (allocatedresourceshape *AllocatedResourceShape) UnstageVoid(stage *Stage) {
	delete(stage.AllocatedResourceShapes, allocatedresourceshape)
	// issue1150
	// delete(stage.AllocatedResourceShape_stagedOrder, allocatedresourceshape)
	delete(stage.AllocatedResourceShapes_mapString, allocatedresourceshape.Name)
}

// commit allocatedresourceshape to the back repo (if it is already staged)
func (allocatedresourceshape *AllocatedResourceShape) Commit(stage *Stage) *AllocatedResourceShape {
	if _, ok := stage.AllocatedResourceShapes[allocatedresourceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAllocatedResourceShape(allocatedresourceshape)
		}
	}
	return allocatedresourceshape
}

func (allocatedresourceshape *AllocatedResourceShape) CommitVoid(stage *Stage) {
	allocatedresourceshape.Commit(stage)
}

func (allocatedresourceshape *AllocatedResourceShape) StageVoid(stage *Stage) {
	allocatedresourceshape.Stage(stage)
}

// Checkout allocatedresourceshape to the back repo (if it is already staged)
func (allocatedresourceshape *AllocatedResourceShape) Checkout(stage *Stage) *AllocatedResourceShape {
	if _, ok := stage.AllocatedResourceShapes[allocatedresourceshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAllocatedResourceShape(allocatedresourceshape)
		}
	}
	return allocatedresourceshape
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

// Stage puts data to the model stage
func (data *Data) Stage(stage *Stage) *Data {
	if _, ok := stage.Datas[data]; !ok {
		stage.Datas[data] = struct{}{}
		stage.Data_stagedOrder[data] = stage.DataOrder
		stage.Data_orderStaged[stage.DataOrder] = data
		stage.DataOrder++
	}
	stage.Datas_mapString[data.Name] = data

	return data
}

// StagePreserveOrder puts data to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataOrder
// - update stage.DataOrder accordingly
func (data *Data) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Datas[data]; !ok {
		stage.Datas[data] = struct{}{}

		if order > stage.DataOrder {
			stage.DataOrder = order
		}
		stage.Data_stagedOrder[data] = order
		stage.Data_orderStaged[order] = data
		stage.DataOrder++
	}
	stage.Datas_mapString[data.Name] = data
}

// Unstage removes data off the model stage
func (data *Data) Unstage(stage *Stage) *Data {
	delete(stage.Datas, data)
	// issue1150
	// delete(stage.Data_stagedOrder, data)
	delete(stage.Datas_mapString, data.Name)

	return data
}

// UnstageVoid removes data off the model stage
func (data *Data) UnstageVoid(stage *Stage) {
	delete(stage.Datas, data)
	// issue1150
	// delete(stage.Data_stagedOrder, data)
	delete(stage.Datas_mapString, data.Name)
}

// commit data to the back repo (if it is already staged)
func (data *Data) Commit(stage *Stage) *Data {
	if _, ok := stage.Datas[data]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitData(data)
		}
	}
	return data
}

func (data *Data) CommitVoid(stage *Stage) {
	data.Commit(stage)
}

func (data *Data) StageVoid(stage *Stage) {
	data.Stage(stage)
}

// Checkout data to the back repo (if it is already staged)
func (data *Data) Checkout(stage *Stage) *Data {
	if _, ok := stage.Datas[data]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutData(data)
		}
	}
	return data
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

// Stage puts datashape to the model stage
func (datashape *DataShape) Stage(stage *Stage) *DataShape {
	if _, ok := stage.DataShapes[datashape]; !ok {
		stage.DataShapes[datashape] = struct{}{}
		stage.DataShape_stagedOrder[datashape] = stage.DataShapeOrder
		stage.DataShape_orderStaged[stage.DataShapeOrder] = datashape
		stage.DataShapeOrder++
	}
	stage.DataShapes_mapString[datashape.Name] = datashape

	return datashape
}

// StagePreserveOrder puts datashape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DataShapeOrder
// - update stage.DataShapeOrder accordingly
func (datashape *DataShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DataShapes[datashape]; !ok {
		stage.DataShapes[datashape] = struct{}{}

		if order > stage.DataShapeOrder {
			stage.DataShapeOrder = order
		}
		stage.DataShape_stagedOrder[datashape] = order
		stage.DataShape_orderStaged[order] = datashape
		stage.DataShapeOrder++
	}
	stage.DataShapes_mapString[datashape.Name] = datashape
}

// Unstage removes datashape off the model stage
func (datashape *DataShape) Unstage(stage *Stage) *DataShape {
	delete(stage.DataShapes, datashape)
	// issue1150
	// delete(stage.DataShape_stagedOrder, datashape)
	delete(stage.DataShapes_mapString, datashape.Name)

	return datashape
}

// UnstageVoid removes datashape off the model stage
func (datashape *DataShape) UnstageVoid(stage *Stage) {
	delete(stage.DataShapes, datashape)
	// issue1150
	// delete(stage.DataShape_stagedOrder, datashape)
	delete(stage.DataShapes_mapString, datashape.Name)
}

// commit datashape to the back repo (if it is already staged)
func (datashape *DataShape) Commit(stage *Stage) *DataShape {
	if _, ok := stage.DataShapes[datashape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDataShape(datashape)
		}
	}
	return datashape
}

func (datashape *DataShape) CommitVoid(stage *Stage) {
	datashape.Commit(stage)
}

func (datashape *DataShape) StageVoid(stage *Stage) {
	datashape.Stage(stage)
}

// Checkout datashape to the back repo (if it is already staged)
func (datashape *DataShape) Checkout(stage *Stage) *DataShape {
	if _, ok := stage.DataShapes[datashape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDataShape(datashape)
		}
	}
	return datashape
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

// Stage puts externalparticipantshape to the model stage
func (externalparticipantshape *ExternalParticipantShape) Stage(stage *Stage) *ExternalParticipantShape {
	if _, ok := stage.ExternalParticipantShapes[externalparticipantshape]; !ok {
		stage.ExternalParticipantShapes[externalparticipantshape] = struct{}{}
		stage.ExternalParticipantShape_stagedOrder[externalparticipantshape] = stage.ExternalParticipantShapeOrder
		stage.ExternalParticipantShape_orderStaged[stage.ExternalParticipantShapeOrder] = externalparticipantshape
		stage.ExternalParticipantShapeOrder++
	}
	stage.ExternalParticipantShapes_mapString[externalparticipantshape.Name] = externalparticipantshape

	return externalparticipantshape
}

// StagePreserveOrder puts externalparticipantshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ExternalParticipantShapeOrder
// - update stage.ExternalParticipantShapeOrder accordingly
func (externalparticipantshape *ExternalParticipantShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.ExternalParticipantShapes[externalparticipantshape]; !ok {
		stage.ExternalParticipantShapes[externalparticipantshape] = struct{}{}

		if order > stage.ExternalParticipantShapeOrder {
			stage.ExternalParticipantShapeOrder = order
		}
		stage.ExternalParticipantShape_stagedOrder[externalparticipantshape] = order
		stage.ExternalParticipantShape_orderStaged[order] = externalparticipantshape
		stage.ExternalParticipantShapeOrder++
	}
	stage.ExternalParticipantShapes_mapString[externalparticipantshape.Name] = externalparticipantshape
}

// Unstage removes externalparticipantshape off the model stage
func (externalparticipantshape *ExternalParticipantShape) Unstage(stage *Stage) *ExternalParticipantShape {
	delete(stage.ExternalParticipantShapes, externalparticipantshape)
	// issue1150
	// delete(stage.ExternalParticipantShape_stagedOrder, externalparticipantshape)
	delete(stage.ExternalParticipantShapes_mapString, externalparticipantshape.Name)

	return externalparticipantshape
}

// UnstageVoid removes externalparticipantshape off the model stage
func (externalparticipantshape *ExternalParticipantShape) UnstageVoid(stage *Stage) {
	delete(stage.ExternalParticipantShapes, externalparticipantshape)
	// issue1150
	// delete(stage.ExternalParticipantShape_stagedOrder, externalparticipantshape)
	delete(stage.ExternalParticipantShapes_mapString, externalparticipantshape.Name)
}

// commit externalparticipantshape to the back repo (if it is already staged)
func (externalparticipantshape *ExternalParticipantShape) Commit(stage *Stage) *ExternalParticipantShape {
	if _, ok := stage.ExternalParticipantShapes[externalparticipantshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitExternalParticipantShape(externalparticipantshape)
		}
	}
	return externalparticipantshape
}

func (externalparticipantshape *ExternalParticipantShape) CommitVoid(stage *Stage) {
	externalparticipantshape.Commit(stage)
}

func (externalparticipantshape *ExternalParticipantShape) StageVoid(stage *Stage) {
	externalparticipantshape.Stage(stage)
}

// Checkout externalparticipantshape to the back repo (if it is already staged)
func (externalparticipantshape *ExternalParticipantShape) Checkout(stage *Stage) *ExternalParticipantShape {
	if _, ok := stage.ExternalParticipantShapes[externalparticipantshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutExternalParticipantShape(externalparticipantshape)
		}
	}
	return externalparticipantshape
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
	CreateORMAllocatedProcessShape(AllocatedProcessShape *AllocatedProcessShape)
	CreateORMAllocatedResourceShape(AllocatedResourceShape *AllocatedResourceShape)
	CreateORMControlFlow(ControlFlow *ControlFlow)
	CreateORMControlFlowShape(ControlFlowShape *ControlFlowShape)
	CreateORMData(Data *Data)
	CreateORMDataFlow(DataFlow *DataFlow)
	CreateORMDataFlowShape(DataFlowShape *DataFlowShape)
	CreateORMDataShape(DataShape *DataShape)
	CreateORMDiagramProcess(DiagramProcess *DiagramProcess)
	CreateORMExternalParticipantShape(ExternalParticipantShape *ExternalParticipantShape)
	CreateORMLibrary(Library *Library)
	CreateORMNote(Note *Note)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMNoteTaskShape(NoteTaskShape *NoteTaskShape)
	CreateORMParticipant(Participant *Participant)
	CreateORMParticipantShape(ParticipantShape *ParticipantShape)
	CreateORMProcess(Process *Process)
	CreateORMProcessShape(ProcessShape *ProcessShape)
	CreateORMResource(Resource *Resource)
	CreateORMTask(Task *Task)
	CreateORMTaskShape(TaskShape *TaskShape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAllocatedProcessShape(AllocatedProcessShape *AllocatedProcessShape)
	DeleteORMAllocatedResourceShape(AllocatedResourceShape *AllocatedResourceShape)
	DeleteORMControlFlow(ControlFlow *ControlFlow)
	DeleteORMControlFlowShape(ControlFlowShape *ControlFlowShape)
	DeleteORMData(Data *Data)
	DeleteORMDataFlow(DataFlow *DataFlow)
	DeleteORMDataFlowShape(DataFlowShape *DataFlowShape)
	DeleteORMDataShape(DataShape *DataShape)
	DeleteORMDiagramProcess(DiagramProcess *DiagramProcess)
	DeleteORMExternalParticipantShape(ExternalParticipantShape *ExternalParticipantShape)
	DeleteORMLibrary(Library *Library)
	DeleteORMNote(Note *Note)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMNoteTaskShape(NoteTaskShape *NoteTaskShape)
	DeleteORMParticipant(Participant *Participant)
	DeleteORMParticipantShape(ParticipantShape *ParticipantShape)
	DeleteORMProcess(Process *Process)
	DeleteORMProcessShape(ProcessShape *ProcessShape)
	DeleteORMResource(Resource *Resource)
	DeleteORMTask(Task *Task)
	DeleteORMTaskShape(TaskShape *TaskShape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.AllocatedProcessShapes = make(map[*AllocatedProcessShape]struct{})
	stage.AllocatedProcessShapes_mapString = make(map[string]*AllocatedProcessShape)
	stage.AllocatedProcessShape_stagedOrder = make(map[*AllocatedProcessShape]uint)
	stage.AllocatedProcessShapeOrder = 0

	stage.AllocatedResourceShapes = make(map[*AllocatedResourceShape]struct{})
	stage.AllocatedResourceShapes_mapString = make(map[string]*AllocatedResourceShape)
	stage.AllocatedResourceShape_stagedOrder = make(map[*AllocatedResourceShape]uint)
	stage.AllocatedResourceShapeOrder = 0

	stage.ControlFlows = make(map[*ControlFlow]struct{})
	stage.ControlFlows_mapString = make(map[string]*ControlFlow)
	stage.ControlFlow_stagedOrder = make(map[*ControlFlow]uint)
	stage.ControlFlowOrder = 0

	stage.ControlFlowShapes = make(map[*ControlFlowShape]struct{})
	stage.ControlFlowShapes_mapString = make(map[string]*ControlFlowShape)
	stage.ControlFlowShape_stagedOrder = make(map[*ControlFlowShape]uint)
	stage.ControlFlowShapeOrder = 0

	stage.Datas = make(map[*Data]struct{})
	stage.Datas_mapString = make(map[string]*Data)
	stage.Data_stagedOrder = make(map[*Data]uint)
	stage.DataOrder = 0

	stage.DataFlows = make(map[*DataFlow]struct{})
	stage.DataFlows_mapString = make(map[string]*DataFlow)
	stage.DataFlow_stagedOrder = make(map[*DataFlow]uint)
	stage.DataFlowOrder = 0

	stage.DataFlowShapes = make(map[*DataFlowShape]struct{})
	stage.DataFlowShapes_mapString = make(map[string]*DataFlowShape)
	stage.DataFlowShape_stagedOrder = make(map[*DataFlowShape]uint)
	stage.DataFlowShapeOrder = 0

	stage.DataShapes = make(map[*DataShape]struct{})
	stage.DataShapes_mapString = make(map[string]*DataShape)
	stage.DataShape_stagedOrder = make(map[*DataShape]uint)
	stage.DataShapeOrder = 0

	stage.DiagramProcesss = make(map[*DiagramProcess]struct{})
	stage.DiagramProcesss_mapString = make(map[string]*DiagramProcess)
	stage.DiagramProcess_stagedOrder = make(map[*DiagramProcess]uint)
	stage.DiagramProcessOrder = 0

	stage.ExternalParticipantShapes = make(map[*ExternalParticipantShape]struct{})
	stage.ExternalParticipantShapes_mapString = make(map[string]*ExternalParticipantShape)
	stage.ExternalParticipantShape_stagedOrder = make(map[*ExternalParticipantShape]uint)
	stage.ExternalParticipantShapeOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Notes = make(map[*Note]struct{})
	stage.Notes_mapString = make(map[string]*Note)
	stage.Note_stagedOrder = make(map[*Note]uint)
	stage.NoteOrder = 0

	stage.NoteShapes = make(map[*NoteShape]struct{})
	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	stage.NoteShape_stagedOrder = make(map[*NoteShape]uint)
	stage.NoteShapeOrder = 0

	stage.NoteTaskShapes = make(map[*NoteTaskShape]struct{})
	stage.NoteTaskShapes_mapString = make(map[string]*NoteTaskShape)
	stage.NoteTaskShape_stagedOrder = make(map[*NoteTaskShape]uint)
	stage.NoteTaskShapeOrder = 0

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

	stage.Resources = make(map[*Resource]struct{})
	stage.Resources_mapString = make(map[string]*Resource)
	stage.Resource_stagedOrder = make(map[*Resource]uint)
	stage.ResourceOrder = 0

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
	stage.AllocatedProcessShapes = nil
	stage.AllocatedProcessShapes_mapString = nil

	stage.AllocatedResourceShapes = nil
	stage.AllocatedResourceShapes_mapString = nil

	stage.ControlFlows = nil
	stage.ControlFlows_mapString = nil

	stage.ControlFlowShapes = nil
	stage.ControlFlowShapes_mapString = nil

	stage.Datas = nil
	stage.Datas_mapString = nil

	stage.DataFlows = nil
	stage.DataFlows_mapString = nil

	stage.DataFlowShapes = nil
	stage.DataFlowShapes_mapString = nil

	stage.DataShapes = nil
	stage.DataShapes_mapString = nil

	stage.DiagramProcesss = nil
	stage.DiagramProcesss_mapString = nil

	stage.ExternalParticipantShapes = nil
	stage.ExternalParticipantShapes_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Notes = nil
	stage.Notes_mapString = nil

	stage.NoteShapes = nil
	stage.NoteShapes_mapString = nil

	stage.NoteTaskShapes = nil
	stage.NoteTaskShapes_mapString = nil

	stage.Participants = nil
	stage.Participants_mapString = nil

	stage.ParticipantShapes = nil
	stage.ParticipantShapes_mapString = nil

	stage.Processs = nil
	stage.Processs_mapString = nil

	stage.ProcessShapes = nil
	stage.ProcessShapes_mapString = nil

	stage.Resources = nil
	stage.Resources_mapString = nil

	stage.Tasks = nil
	stage.Tasks_mapString = nil

	stage.TaskShapes = nil
	stage.TaskShapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for allocatedprocessshape := range stage.AllocatedProcessShapes {
		allocatedprocessshape.Unstage(stage)
	}

	for allocatedresourceshape := range stage.AllocatedResourceShapes {
		allocatedresourceshape.Unstage(stage)
	}

	for controlflow := range stage.ControlFlows {
		controlflow.Unstage(stage)
	}

	for controlflowshape := range stage.ControlFlowShapes {
		controlflowshape.Unstage(stage)
	}

	for data := range stage.Datas {
		data.Unstage(stage)
	}

	for dataflow := range stage.DataFlows {
		dataflow.Unstage(stage)
	}

	for dataflowshape := range stage.DataFlowShapes {
		dataflowshape.Unstage(stage)
	}

	for datashape := range stage.DataShapes {
		datashape.Unstage(stage)
	}

	for diagramprocess := range stage.DiagramProcesss {
		diagramprocess.Unstage(stage)
	}

	for externalparticipantshape := range stage.ExternalParticipantShapes {
		externalparticipantshape.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for note := range stage.Notes {
		note.Unstage(stage)
	}

	for noteshape := range stage.NoteShapes {
		noteshape.Unstage(stage)
	}

	for notetaskshape := range stage.NoteTaskShapes {
		notetaskshape.Unstage(stage)
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

	for resource := range stage.Resources {
		resource.Unstage(stage)
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
	case map[*AllocatedProcessShape]any:
		return any(&stage.AllocatedProcessShapes).(*Type)
	case map[*AllocatedResourceShape]any:
		return any(&stage.AllocatedResourceShapes).(*Type)
	case map[*ControlFlow]any:
		return any(&stage.ControlFlows).(*Type)
	case map[*ControlFlowShape]any:
		return any(&stage.ControlFlowShapes).(*Type)
	case map[*Data]any:
		return any(&stage.Datas).(*Type)
	case map[*DataFlow]any:
		return any(&stage.DataFlows).(*Type)
	case map[*DataFlowShape]any:
		return any(&stage.DataFlowShapes).(*Type)
	case map[*DataShape]any:
		return any(&stage.DataShapes).(*Type)
	case map[*DiagramProcess]any:
		return any(&stage.DiagramProcesss).(*Type)
	case map[*ExternalParticipantShape]any:
		return any(&stage.ExternalParticipantShapes).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Note]any:
		return any(&stage.Notes).(*Type)
	case map[*NoteShape]any:
		return any(&stage.NoteShapes).(*Type)
	case map[*NoteTaskShape]any:
		return any(&stage.NoteTaskShapes).(*Type)
	case map[*Participant]any:
		return any(&stage.Participants).(*Type)
	case map[*ParticipantShape]any:
		return any(&stage.ParticipantShapes).(*Type)
	case map[*Process]any:
		return any(&stage.Processs).(*Type)
	case map[*ProcessShape]any:
		return any(&stage.ProcessShapes).(*Type)
	case map[*Resource]any:
		return any(&stage.Resources).(*Type)
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

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AllocatedProcessShape:
		return any(&stage.AllocatedProcessShapes).(*map[*Type]struct{})
	case AllocatedResourceShape:
		return any(&stage.AllocatedResourceShapes).(*map[*Type]struct{})
	case ControlFlow:
		return any(&stage.ControlFlows).(*map[*Type]struct{})
	case ControlFlowShape:
		return any(&stage.ControlFlowShapes).(*map[*Type]struct{})
	case Data:
		return any(&stage.Datas).(*map[*Type]struct{})
	case DataFlow:
		return any(&stage.DataFlows).(*map[*Type]struct{})
	case DataFlowShape:
		return any(&stage.DataFlowShapes).(*map[*Type]struct{})
	case DataShape:
		return any(&stage.DataShapes).(*map[*Type]struct{})
	case DiagramProcess:
		return any(&stage.DiagramProcesss).(*map[*Type]struct{})
	case ExternalParticipantShape:
		return any(&stage.ExternalParticipantShapes).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Note:
		return any(&stage.Notes).(*map[*Type]struct{})
	case NoteShape:
		return any(&stage.NoteShapes).(*map[*Type]struct{})
	case NoteTaskShape:
		return any(&stage.NoteTaskShapes).(*map[*Type]struct{})
	case Participant:
		return any(&stage.Participants).(*map[*Type]struct{})
	case ParticipantShape:
		return any(&stage.ParticipantShapes).(*map[*Type]struct{})
	case Process:
		return any(&stage.Processs).(*map[*Type]struct{})
	case ProcessShape:
		return any(&stage.ProcessShapes).(*map[*Type]struct{})
	case Resource:
		return any(&stage.Resources).(*map[*Type]struct{})
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AllocatedProcessShape:
		return any(&stage.AllocatedProcessShapes_mapString).(*map[string]*Type)
	case AllocatedResourceShape:
		return any(&stage.AllocatedResourceShapes_mapString).(*map[string]*Type)
	case ControlFlow:
		return any(&stage.ControlFlows_mapString).(*map[string]*Type)
	case ControlFlowShape:
		return any(&stage.ControlFlowShapes_mapString).(*map[string]*Type)
	case Data:
		return any(&stage.Datas_mapString).(*map[string]*Type)
	case DataFlow:
		return any(&stage.DataFlows_mapString).(*map[string]*Type)
	case DataFlowShape:
		return any(&stage.DataFlowShapes_mapString).(*map[string]*Type)
	case DataShape:
		return any(&stage.DataShapes_mapString).(*map[string]*Type)
	case DiagramProcess:
		return any(&stage.DiagramProcesss_mapString).(*map[string]*Type)
	case ExternalParticipantShape:
		return any(&stage.ExternalParticipantShapes_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Note:
		return any(&stage.Notes_mapString).(*map[string]*Type)
	case NoteShape:
		return any(&stage.NoteShapes_mapString).(*map[string]*Type)
	case NoteTaskShape:
		return any(&stage.NoteTaskShapes_mapString).(*map[string]*Type)
	case Participant:
		return any(&stage.Participants_mapString).(*map[string]*Type)
	case ParticipantShape:
		return any(&stage.ParticipantShapes_mapString).(*map[string]*Type)
	case Process:
		return any(&stage.Processs_mapString).(*map[string]*Type)
	case ProcessShape:
		return any(&stage.ProcessShapes_mapString).(*map[string]*Type)
	case Resource:
		return any(&stage.Resources_mapString).(*map[string]*Type)
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
	case AllocatedProcessShape:
		return any(&AllocatedProcessShape{
			// Initialisation of associations
			// field is initialized with an instance of Participant with the name of the field
			Participant: &Participant{Name: "Participant"},
			// field is initialized with an instance of Process with the name of the field
			Process: &Process{Name: "Process"},
		}).(*Type)
	case AllocatedResourceShape:
		return any(&AllocatedResourceShape{
			// Initialisation of associations
			// field is initialized with an instance of Participant with the name of the field
			Participant: &Participant{Name: "Participant"},
			// field is initialized with an instance of Resource with the name of the field
			Resource: &Resource{Name: "Resource"},
		}).(*Type)
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
	case Data:
		return any(&Data{
			// Initialisation of associations
		}).(*Type)
	case DataFlow:
		return any(&DataFlow{
			// Initialisation of associations
			// field is initialized with an instance of Data with the name of the field
			Datas: []*Data{{Name: "Datas"}},
			// field is initialized with an instance of Task with the name of the field
			StartTask: &Task{Name: "StartTask"},
			// field is initialized with an instance of Task with the name of the field
			EndTask: &Task{Name: "EndTask"},
			// field is initialized with an instance of Participant with the name of the field
			StartExternalParticipant: &Participant{Name: "StartExternalParticipant"},
			// field is initialized with an instance of Participant with the name of the field
			EndExternalParticipant: &Participant{Name: "EndExternalParticipant"},
		}).(*Type)
	case DataFlowShape:
		return any(&DataFlowShape{
			// Initialisation of associations
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlow: &DataFlow{Name: "DataFlow"},
		}).(*Type)
	case DataShape:
		return any(&DataShape{
			// Initialisation of associations
			// field is initialized with an instance of Data with the name of the field
			Data: &Data{Name: "Data"},
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
			// field is initialized with an instance of ExternalParticipantShape with the name of the field
			ExternalParticipant_Shapes: []*ExternalParticipantShape{{Name: "ExternalParticipant_Shapes"}},
			// field is initialized with an instance of Participant with the name of the field
			ExternalParticipantWhoseNodeIsExpanded: []*Participant{{Name: "ExternalParticipantWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Participant with the name of the field
			ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded: []*Participant{{Name: "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Participant with the name of the field
			ExternalParticipantsWhoseInDataFlowsNodeIsExpanded: []*Participant{{Name: "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TasksWhoseNodeIsExpanded: []*Task{{Name: "TasksWhoseNodeIsExpanded"}},
			// field is initialized with an instance of TaskShape with the name of the field
			Task_Shapes: []*TaskShape{{Name: "Task_Shapes"}},
			// field is initialized with an instance of ControlFlow with the name of the field
			ControlFlowsWhoseNodeIsExpanded: []*ControlFlow{{Name: "ControlFlowsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of ControlFlowShape with the name of the field
			ControlFlow_Shapes: []*ControlFlowShape{{Name: "ControlFlow_Shapes"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlowsWhoseNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataFlowShape with the name of the field
			DataFlow_Shapes: []*DataFlowShape{{Name: "DataFlow_Shapes"}},
			// field is initialized with an instance of Data with the name of the field
			DatasWhoseNodeIsExpanded: []*Data{{Name: "DatasWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataShape with the name of the field
			Data_Shapes: []*DataShape{{Name: "Data_Shapes"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlowsWhoseDataNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseDataNodeIsExpanded"}},
			// field is initialized with an instance of Resource with the name of the field
			AllocatedResourcesWhoseNodeIsExpanded: []*Resource{{Name: "AllocatedResourcesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of AllocatedResourceShape with the name of the field
			AllocatedResourceShapes: []*AllocatedResourceShape{{Name: "AllocatedResourceShapes"}},
			// field is initialized with an instance of Process with the name of the field
			AllocatedProcessesWhoseNodeIsExpanded: []*Process{{Name: "AllocatedProcessesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of AllocatedProcessShape with the name of the field
			AllocatedProcessShapes: []*AllocatedProcessShape{{Name: "AllocatedProcessShapes"}},
			// field is initialized with an instance of NoteShape with the name of the field
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of NoteTaskShape with the name of the field
			NoteTaskShapes: []*NoteTaskShape{{Name: "NoteTaskShapes"}},
		}).(*Type)
	case ExternalParticipantShape:
		return any(&ExternalParticipantShape{
			// Initialisation of associations
			// field is initialized with an instance of Participant with the name of the field
			Participant: &Participant{Name: "Participant"},
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Process with the name of the field
			RootProcesses: []*Process{{Name: "RootProcesses"}},
			// field is initialized with an instance of Process with the name of the field
			ProcesssWhoseNodeIsExpanded: []*Process{{Name: "ProcesssWhoseNodeIsExpanded"}},
			// field is initialized with an instance of DataFlow with the name of the field
			RootDataFlows: []*DataFlow{{Name: "RootDataFlows"}},
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlowsWhoseNodeIsExpanded: []*DataFlow{{Name: "DataFlowsWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Data with the name of the field
			RootDatas: []*Data{{Name: "RootDatas"}},
			// field is initialized with an instance of Data with the name of the field
			DatasWhoseNodeIsExpanded: []*Data{{Name: "DatasWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Resource with the name of the field
			RootResources: []*Resource{{Name: "RootResources"}},
			// field is initialized with an instance of Resource with the name of the field
			ResourcesWhoseNodeIsExpanded: []*Resource{{Name: "ResourcesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Note with the name of the field
			RootNotes: []*Note{{Name: "RootNotes"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
		}).(*Type)
	case Note:
		return any(&Note{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Tasks: []*Task{{Name: "Tasks"}},
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
	case Participant:
		return any(&Participant{
			// Initialisation of associations
			// field is initialized with an instance of Resource with the name of the field
			Resources: []*Resource{{Name: "Resources"}},
			// field is initialized with an instance of Process with the name of the field
			Processes: []*Process{{Name: "Processes"}},
			// field is initialized with an instance of Task with the name of the field
			Tasks: []*Task{{Name: "Tasks"}},
			// field is initialized with an instance of ControlFlow with the name of the field
			ControlFlows: []*ControlFlow{{Name: "ControlFlows"}},
			// field is initialized with an instance of Task with the name of the field
			TaskWhoseOutControlFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseOutControlFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TaskWhoseInControlFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseInControlFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TaskWhoseOutDataFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseOutDataFlowsNodeIsExpanded"}},
			// field is initialized with an instance of Task with the name of the field
			TaskWhoseInDataFlowsNodeIsExpanded: []*Task{{Name: "TaskWhoseInDataFlowsNodeIsExpanded"}},
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
			// field is initialized with an instance of DataFlow with the name of the field
			DataFlows: []*DataFlow{{Name: "DataFlows"}},
			// field is initialized with an instance of Participant with the name of the field
			ExternalParticipants: []*Participant{{Name: "ExternalParticipants"}},
			// field is initialized with an instance of Participant with the name of the field
			ExternalParticipantWhoseNodeIsExpanded: []*Participant{{Name: "ExternalParticipantWhoseNodeIsExpanded"}},
		}).(*Type)
	case ProcessShape:
		return any(&ProcessShape{
			// Initialisation of associations
			// field is initialized with an instance of Process with the name of the field
			Process: &Process{Name: "Process"},
		}).(*Type)
	case Resource:
		return any(&Resource{
			// Initialisation of associations
		}).(*Type)
	case Task:
		return any(&Task{
			// Initialisation of associations
			// field is initialized with an instance of Process with the name of the field
			Type: &Process{Name: "Type"},
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {
	res = make([]ReverseField, 0)

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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:               "SVG_Path",
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
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ComputedPrefix",
			GongFieldValueType: GongFieldValueTypeString,
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
	case "Description":
		res.valueString = data.Description
	case "ComputedPrefix":
		res.valueString = data.ComputedPrefix
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
	case "ComputedPrefix":
		res.valueString = note.ComputedPrefix
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
	case "SVG_Path":
		res.valueString = process.SVG_Path
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
	case "Description":
		res.valueString = resource.Description
	case "ComputedPrefix":
		res.valueString = resource.ComputedPrefix
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

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (allocatedprocessshape *AllocatedProcessShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		allocatedprocessshape.Name = value.GetValueString()
	case "Participant":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			allocatedprocessshape.Participant = nil
			for __instance__ := range stage.Participants {
				if stage.Participant_stagedOrder[__instance__] == uint(id) {
					allocatedprocessshape.Participant = __instance__
					break
				}
			}
		}
	case "Process":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			allocatedprocessshape.Process = nil
			for __instance__ := range stage.Processs {
				if stage.Process_stagedOrder[__instance__] == uint(id) {
					allocatedprocessshape.Process = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (allocatedresourceshape *AllocatedResourceShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		allocatedresourceshape.Name = value.GetValueString()
	case "Participant":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			allocatedresourceshape.Participant = nil
			for __instance__ := range stage.Participants {
				if stage.Participant_stagedOrder[__instance__] == uint(id) {
					allocatedresourceshape.Participant = __instance__
					break
				}
			}
		}
	case "Resource":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			allocatedresourceshape.Resource = nil
			for __instance__ := range stage.Resources {
				if stage.Resource_stagedOrder[__instance__] == uint(id) {
					allocatedresourceshape.Resource = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (controlflow *ControlFlow) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		controlflow.Name = value.GetValueString()
	case "Description":
		controlflow.Description = value.GetValueString()
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

func (data *Data) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		data.Name = value.GetValueString()
	case "Description":
		data.Description = value.GetValueString()
	case "ComputedPrefix":
		data.ComputedPrefix = value.GetValueString()
	case "SVG_Path":
		data.SVG_Path = value.GetValueString()
	case "InverseAppliedScaling":
		data.InverseAppliedScaling = value.GetValueFloat()
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
	case "Datas":
		dataflow.Datas = make([]*Data, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Datas {
					if stage.Data_stagedOrder[__instance__] == uint(id) {
						dataflow.Datas = append(dataflow.Datas, __instance__)
						break
					}
				}
			}
		}
	case "Description":
		dataflow.Description = value.GetValueString()
	case "ComputedPrefix":
		dataflow.ComputedPrefix = value.GetValueString()
	case "Type":
		dataflow.Type.FromCodeString(value.GetValueString())
	case "StartTask":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.StartTask = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					dataflow.StartTask = __instance__
					break
				}
			}
		}
	case "EndTask":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.EndTask = nil
			for __instance__ := range stage.Tasks {
				if stage.Task_stagedOrder[__instance__] == uint(id) {
					dataflow.EndTask = __instance__
					break
				}
			}
		}
	case "StartExternalParticipant":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.StartExternalParticipant = nil
			for __instance__ := range stage.Participants {
				if stage.Participant_stagedOrder[__instance__] == uint(id) {
					dataflow.StartExternalParticipant = __instance__
					break
				}
			}
		}
	case "EndExternalParticipant":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			dataflow.EndExternalParticipant = nil
			for __instance__ := range stage.Participants {
				if stage.Participant_stagedOrder[__instance__] == uint(id) {
					dataflow.EndExternalParticipant = __instance__
					break
				}
			}
		}
	case "IsDatasNodeExpanded":
		dataflow.IsDatasNodeExpanded = value.GetValueBool()
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

func (datashape *DataShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		datashape.Name = value.GetValueString()
	case "Data":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			datashape.Data = nil
			for __instance__ := range stage.Datas {
				if stage.Data_stagedOrder[__instance__] == uint(id) {
					datashape.Data = __instance__
					break
				}
			}
		}
	case "DataFlow":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			datashape.DataFlow = nil
			for __instance__ := range stage.DataFlows {
				if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
					datashape.DataFlow = __instance__
					break
				}
			}
		}
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
	case "Description":
		diagramprocess.Description = value.GetValueString()
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
	case "ExternalParticipant_Shapes":
		diagramprocess.ExternalParticipant_Shapes = make([]*ExternalParticipantShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ExternalParticipantShapes {
					if stage.ExternalParticipantShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ExternalParticipant_Shapes = append(diagramprocess.ExternalParticipant_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "IsExternalParticipantsNodeExpanded":
		diagramprocess.IsExternalParticipantsNodeExpanded = value.GetValueBool()
	case "ExternalParticipantWhoseNodeIsExpanded":
		diagramprocess.ExternalParticipantWhoseNodeIsExpanded = make([]*Participant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Participants {
					if stage.Participant_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ExternalParticipantWhoseNodeIsExpanded = append(diagramprocess.ExternalParticipantWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
		diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = make([]*Participant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Participants {
					if stage.Participant_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
		diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = make([]*Participant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Participants {
					if stage.Participant_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, __instance__)
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
	case "Task_Shapes":
		diagramprocess.Task_Shapes = make([]*TaskShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.TaskShapes {
					if stage.TaskShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.Task_Shapes = append(diagramprocess.Task_Shapes, __instance__)
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
	case "ControlFlow_Shapes":
		diagramprocess.ControlFlow_Shapes = make([]*ControlFlowShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ControlFlowShapes {
					if stage.ControlFlowShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.ControlFlow_Shapes = append(diagramprocess.ControlFlow_Shapes, __instance__)
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
	case "DataFlow_Shapes":
		diagramprocess.DataFlow_Shapes = make([]*DataFlowShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlowShapes {
					if stage.DataFlowShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.DataFlow_Shapes = append(diagramprocess.DataFlow_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "DatasWhoseNodeIsExpanded":
		diagramprocess.DatasWhoseNodeIsExpanded = make([]*Data, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Datas {
					if stage.Data_stagedOrder[__instance__] == uint(id) {
						diagramprocess.DatasWhoseNodeIsExpanded = append(diagramprocess.DatasWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Data_Shapes":
		diagramprocess.Data_Shapes = make([]*DataShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataShapes {
					if stage.DataShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.Data_Shapes = append(diagramprocess.Data_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "DataFlowsWhoseDataNodeIsExpanded":
		diagramprocess.DataFlowsWhoseDataNodeIsExpanded = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						diagramprocess.DataFlowsWhoseDataNodeIsExpanded = append(diagramprocess.DataFlowsWhoseDataNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "AllocatedResourcesWhoseNodeIsExpanded":
		diagramprocess.AllocatedResourcesWhoseNodeIsExpanded = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						diagramprocess.AllocatedResourcesWhoseNodeIsExpanded = append(diagramprocess.AllocatedResourcesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "AllocatedResourceShapes":
		diagramprocess.AllocatedResourceShapes = make([]*AllocatedResourceShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AllocatedResourceShapes {
					if stage.AllocatedResourceShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.AllocatedResourceShapes = append(diagramprocess.AllocatedResourceShapes, __instance__)
						break
					}
				}
			}
		}
	case "AllocatedProcessesWhoseNodeIsExpanded":
		diagramprocess.AllocatedProcessesWhoseNodeIsExpanded = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						diagramprocess.AllocatedProcessesWhoseNodeIsExpanded = append(diagramprocess.AllocatedProcessesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "AllocatedProcessShapes":
		diagramprocess.AllocatedProcessShapes = make([]*AllocatedProcessShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.AllocatedProcessShapes {
					if stage.AllocatedProcessShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.AllocatedProcessShapes = append(diagramprocess.AllocatedProcessShapes, __instance__)
						break
					}
				}
			}
		}
	case "Note_Shapes":
		diagramprocess.Note_Shapes = make([]*NoteShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteShapes {
					if stage.NoteShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.Note_Shapes = append(diagramprocess.Note_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "NotesWhoseNodeIsExpanded":
		diagramprocess.NotesWhoseNodeIsExpanded = make([]*Note, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Notes {
					if stage.Note_stagedOrder[__instance__] == uint(id) {
						diagramprocess.NotesWhoseNodeIsExpanded = append(diagramprocess.NotesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "IsNotesNodeExpanded":
		diagramprocess.IsNotesNodeExpanded = value.GetValueBool()
	case "NoteTaskShapes":
		diagramprocess.NoteTaskShapes = make([]*NoteTaskShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteTaskShapes {
					if stage.NoteTaskShape_stagedOrder[__instance__] == uint(id) {
						diagramprocess.NoteTaskShapes = append(diagramprocess.NoteTaskShapes, __instance__)
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

func (externalparticipantshape *ExternalParticipantShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		externalparticipantshape.Name = value.GetValueString()
	case "Participant":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			externalparticipantshape.Participant = nil
			for __instance__ := range stage.Participants {
				if stage.Participant_stagedOrder[__instance__] == uint(id) {
					externalparticipantshape.Participant = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		externalparticipantshape.IsExpanded = value.GetValueBool()
	case "X":
		externalparticipantshape.X = value.GetValueFloat()
	case "Y":
		externalparticipantshape.Y = value.GetValueFloat()
	case "Width":
		externalparticipantshape.Width = value.GetValueFloat()
	case "Height":
		externalparticipantshape.Height = value.GetValueFloat()
	case "IsHidden":
		externalparticipantshape.IsHidden = value.GetValueBool()
	case "TailHeigth":
		externalparticipantshape.TailHeigth = value.GetValueFloat()
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
	case "IsProcessesNodeExpanded":
		library.IsProcessesNodeExpanded = value.GetValueBool()
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
	case "IsDataFlowsNodeExpanded":
		library.IsDataFlowsNodeExpanded = value.GetValueBool()
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
	case "RootDatas":
		library.RootDatas = make([]*Data, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Datas {
					if stage.Data_stagedOrder[__instance__] == uint(id) {
						library.RootDatas = append(library.RootDatas, __instance__)
						break
					}
				}
			}
		}
	case "IsDatasNodeExpanded":
		library.IsDatasNodeExpanded = value.GetValueBool()
	case "DatasWhoseNodeIsExpanded":
		library.DatasWhoseNodeIsExpanded = make([]*Data, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Datas {
					if stage.Data_stagedOrder[__instance__] == uint(id) {
						library.DatasWhoseNodeIsExpanded = append(library.DatasWhoseNodeIsExpanded, __instance__)
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
	case "IsResourcesNodeExpanded":
		library.IsResourcesNodeExpanded = value.GetValueBool()
	case "ResourcesWhoseNodeIsExpanded":
		library.ResourcesWhoseNodeIsExpanded = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						library.ResourcesWhoseNodeIsExpanded = append(library.ResourcesWhoseNodeIsExpanded, __instance__)
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
	case "ComputedPrefix":
		note.ComputedPrefix = value.GetValueString()
	case "IsTasksNodeExpanded":
		note.IsTasksNodeExpanded = value.GetValueBool()
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

func (participant *Participant) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		participant.Name = value.GetValueString()
	case "Description":
		participant.Description = value.GetValueString()
	case "Resources":
		participant.Resources = make([]*Resource, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Resources {
					if stage.Resource_stagedOrder[__instance__] == uint(id) {
						participant.Resources = append(participant.Resources, __instance__)
						break
					}
				}
			}
		}
	case "IsResourcesNodeExpanded":
		participant.IsResourcesNodeExpanded = value.GetValueBool()
	case "Processes":
		participant.Processes = make([]*Process, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Processs {
					if stage.Process_stagedOrder[__instance__] == uint(id) {
						participant.Processes = append(participant.Processes, __instance__)
						break
					}
				}
			}
		}
	case "IsProcessesNodeExpanded":
		participant.IsProcessesNodeExpanded = value.GetValueBool()
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
	case "IsDataFlowsNodeExpanded":
		participant.IsDataFlowsNodeExpanded = value.GetValueBool()
	case "TaskWhoseOutDataFlowsNodeIsExpanded":
		participant.TaskWhoseOutDataFlowsNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						participant.TaskWhoseOutDataFlowsNodeIsExpanded = append(participant.TaskWhoseOutDataFlowsNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "TaskWhoseInDataFlowsNodeIsExpanded":
		participant.TaskWhoseInDataFlowsNodeIsExpanded = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.Task_stagedOrder[__instance__] == uint(id) {
						participant.TaskWhoseInDataFlowsNodeIsExpanded = append(participant.TaskWhoseInDataFlowsNodeIsExpanded, __instance__)
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
	case "WidthWeight":
		participantshape.WidthWeight = value.GetValueFloat()
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
	case "Description":
		process.Description = value.GetValueString()
	case "ComputedPrefix":
		process.ComputedPrefix = value.GetValueString()
	case "SVG_Path":
		process.SVG_Path = value.GetValueString()
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
	case "DataFlows":
		process.DataFlows = make([]*DataFlow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DataFlows {
					if stage.DataFlow_stagedOrder[__instance__] == uint(id) {
						process.DataFlows = append(process.DataFlows, __instance__)
						break
					}
				}
			}
		}
	case "IsDataFlowsNodeExpanded":
		process.IsDataFlowsNodeExpanded = value.GetValueBool()
	case "ExternalParticipants":
		process.ExternalParticipants = make([]*Participant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Participants {
					if stage.Participant_stagedOrder[__instance__] == uint(id) {
						process.ExternalParticipants = append(process.ExternalParticipants, __instance__)
						break
					}
				}
			}
		}
	case "ExternalParticipantWhoseNodeIsExpanded":
		process.ExternalParticipantWhoseNodeIsExpanded = make([]*Participant, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Participants {
					if stage.Participant_stagedOrder[__instance__] == uint(id) {
						process.ExternalParticipantWhoseNodeIsExpanded = append(process.ExternalParticipantWhoseNodeIsExpanded, __instance__)
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

func (resource *Resource) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		resource.Name = value.GetValueString()
	case "Description":
		resource.Description = value.GetValueString()
	case "ComputedPrefix":
		resource.ComputedPrefix = value.GetValueString()
	case "SVG_Path":
		resource.SVG_Path = value.GetValueString()
	case "InverseAppliedScaling":
		resource.InverseAppliedScaling = value.GetValueFloat()
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
	case "Description":
		task.Description = value.GetValueString()
	case "ComputedPrefix":
		task.ComputedPrefix = value.GetValueString()
	case "IsStartTask":
		task.IsStartTask = value.GetValueBool()
	case "IsEndTask":
		task.IsEndTask = value.GetValueBool()
	case "Type":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			task.Type = nil
			for __instance__ := range stage.Processs {
				if stage.Process_stagedOrder[__instance__] == uint(id) {
					task.Type = __instance__
					break
				}
			}
		}
	case "IsTaskNameNotProcessName":
		task.IsTaskNameNotProcessName = value.GetValueBool()
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.AllocatedProcessShapes_mapString = make(map[string]*AllocatedProcessShape)
	for allocatedprocessshape := range stage.AllocatedProcessShapes {
		stage.AllocatedProcessShapes_mapString[allocatedprocessshape.Name] = allocatedprocessshape
	}

	stage.AllocatedResourceShapes_mapString = make(map[string]*AllocatedResourceShape)
	for allocatedresourceshape := range stage.AllocatedResourceShapes {
		stage.AllocatedResourceShapes_mapString[allocatedresourceshape.Name] = allocatedresourceshape
	}

	stage.ControlFlows_mapString = make(map[string]*ControlFlow)
	for controlflow := range stage.ControlFlows {
		stage.ControlFlows_mapString[controlflow.Name] = controlflow
	}

	stage.ControlFlowShapes_mapString = make(map[string]*ControlFlowShape)
	for controlflowshape := range stage.ControlFlowShapes {
		stage.ControlFlowShapes_mapString[controlflowshape.Name] = controlflowshape
	}

	stage.Datas_mapString = make(map[string]*Data)
	for data := range stage.Datas {
		stage.Datas_mapString[data.Name] = data
	}

	stage.DataFlows_mapString = make(map[string]*DataFlow)
	for dataflow := range stage.DataFlows {
		stage.DataFlows_mapString[dataflow.Name] = dataflow
	}

	stage.DataFlowShapes_mapString = make(map[string]*DataFlowShape)
	for dataflowshape := range stage.DataFlowShapes {
		stage.DataFlowShapes_mapString[dataflowshape.Name] = dataflowshape
	}

	stage.DataShapes_mapString = make(map[string]*DataShape)
	for datashape := range stage.DataShapes {
		stage.DataShapes_mapString[datashape.Name] = datashape
	}

	stage.DiagramProcesss_mapString = make(map[string]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		stage.DiagramProcesss_mapString[diagramprocess.Name] = diagramprocess
	}

	stage.ExternalParticipantShapes_mapString = make(map[string]*ExternalParticipantShape)
	for externalparticipantshape := range stage.ExternalParticipantShapes {
		stage.ExternalParticipantShapes_mapString[externalparticipantshape.Name] = externalparticipantshape
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Notes_mapString = make(map[string]*Note)
	for note := range stage.Notes {
		stage.Notes_mapString[note.Name] = note
	}

	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	for noteshape := range stage.NoteShapes {
		stage.NoteShapes_mapString[noteshape.Name] = noteshape
	}

	stage.NoteTaskShapes_mapString = make(map[string]*NoteTaskShape)
	for notetaskshape := range stage.NoteTaskShapes {
		stage.NoteTaskShapes_mapString[notetaskshape.Name] = notetaskshape
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

	stage.Resources_mapString = make(map[string]*Resource)
	for resource := range stage.Resources {
		stage.Resources_mapString[resource.Name] = resource
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
