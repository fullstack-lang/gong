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

	statemachines_go "github.com/fullstack-lang/gong/dsm/statemachines/go"
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
	Actions                map[*Action]struct{}
	Actions_instance       map[*Action]*Action
	Actions_mapString      map[string]*Action
	ActionOrder            uint
	Action_stagedOrder     map[*Action]uint
	Action_orderStaged     map[uint]*Action
	Actions_reference      map[*Action]*Action
	Actions_referenceOrder map[*Action]uint

	// insertion point for slice of pointers maps
	OnAfterActionCreateCallback OnAfterCreateInterface[Action]
	OnAfterActionUpdateCallback OnAfterUpdateInterface[Action]
	OnAfterActionDeleteCallback OnAfterDeleteInterface[Action]
	OnAfterActionReadCallback   OnAfterReadInterface[Action]

	Activitiess                map[*Activities]struct{}
	Activitiess_instance       map[*Activities]*Activities
	Activitiess_mapString      map[string]*Activities
	ActivitiesOrder            uint
	Activities_stagedOrder     map[*Activities]uint
	Activities_orderStaged     map[uint]*Activities
	Activitiess_reference      map[*Activities]*Activities
	Activitiess_referenceOrder map[*Activities]uint

	// insertion point for slice of pointers maps
	OnAfterActivitiesCreateCallback OnAfterCreateInterface[Activities]
	OnAfterActivitiesUpdateCallback OnAfterUpdateInterface[Activities]
	OnAfterActivitiesDeleteCallback OnAfterDeleteInterface[Activities]
	OnAfterActivitiesReadCallback   OnAfterReadInterface[Activities]

	Architectures                map[*Architecture]struct{}
	Architectures_instance       map[*Architecture]*Architecture
	Architectures_mapString      map[string]*Architecture
	ArchitectureOrder            uint
	Architecture_stagedOrder     map[*Architecture]uint
	Architecture_orderStaged     map[uint]*Architecture
	Architectures_reference      map[*Architecture]*Architecture
	Architectures_referenceOrder map[*Architecture]uint

	// insertion point for slice of pointers maps
	Architecture_StateMachines_reverseMap map[*StateMachine]*Architecture

	Architecture_Roles_reverseMap map[*Role]*Architecture

	OnAfterArchitectureCreateCallback OnAfterCreateInterface[Architecture]
	OnAfterArchitectureUpdateCallback OnAfterUpdateInterface[Architecture]
	OnAfterArchitectureDeleteCallback OnAfterDeleteInterface[Architecture]
	OnAfterArchitectureReadCallback   OnAfterReadInterface[Architecture]

	Diagrams                map[*Diagram]struct{}
	Diagrams_instance       map[*Diagram]*Diagram
	Diagrams_mapString      map[string]*Diagram
	DiagramOrder            uint
	Diagram_stagedOrder     map[*Diagram]uint
	Diagram_orderStaged     map[uint]*Diagram
	Diagrams_reference      map[*Diagram]*Diagram
	Diagrams_referenceOrder map[*Diagram]uint

	// insertion point for slice of pointers maps
	Diagram_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*Diagram

	Diagram_State_Shapes_reverseMap map[*StateShape]*Diagram

	Diagram_StatesWhoseNodeIsExpanded_reverseMap map[*State]*Diagram

	Diagram_Transition_Shapes_reverseMap map[*Transition_Shape]*Diagram

	Diagram_Note_Shapes_reverseMap map[*NoteShape]*Diagram

	Diagram_NoteState_Shapes_reverseMap map[*NoteStateShape]*Diagram

	Diagram_NoteTransition_Shapes_reverseMap map[*NoteTransitionShape]*Diagram

	OnAfterDiagramCreateCallback OnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback OnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback OnAfterDeleteInterface[Diagram]
	OnAfterDiagramReadCallback   OnAfterReadInterface[Diagram]

	Guards                map[*Guard]struct{}
	Guards_instance       map[*Guard]*Guard
	Guards_mapString      map[string]*Guard
	GuardOrder            uint
	Guard_stagedOrder     map[*Guard]uint
	Guard_orderStaged     map[uint]*Guard
	Guards_reference      map[*Guard]*Guard
	Guards_referenceOrder map[*Guard]uint

	// insertion point for slice of pointers maps
	OnAfterGuardCreateCallback OnAfterCreateInterface[Guard]
	OnAfterGuardUpdateCallback OnAfterUpdateInterface[Guard]
	OnAfterGuardDeleteCallback OnAfterDeleteInterface[Guard]
	OnAfterGuardReadCallback   OnAfterReadInterface[Guard]

	Kills                map[*Kill]struct{}
	Kills_instance       map[*Kill]*Kill
	Kills_mapString      map[string]*Kill
	KillOrder            uint
	Kill_stagedOrder     map[*Kill]uint
	Kill_orderStaged     map[uint]*Kill
	Kills_reference      map[*Kill]*Kill
	Kills_referenceOrder map[*Kill]uint

	// insertion point for slice of pointers maps
	OnAfterKillCreateCallback OnAfterCreateInterface[Kill]
	OnAfterKillUpdateCallback OnAfterUpdateInterface[Kill]
	OnAfterKillDeleteCallback OnAfterDeleteInterface[Kill]
	OnAfterKillReadCallback   OnAfterReadInterface[Kill]

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

	Library_Diagrams_reverseMap map[*Diagram]*Library

	Library_RootStateMachines_reverseMap map[*StateMachine]*Library

	Library_StateMachinesWhoseNodeIsExpanded_reverseMap map[*StateMachine]*Library

	Library_RootNotes_reverseMap map[*Note]*Library

	Library_NotesWhoseNodeIsExpanded_reverseMap map[*Note]*Library

	Library_SubLibrariesWhoseNodeIsExpanded_reverseMap map[*Library]*Library

	OnAfterLibraryCreateCallback OnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback OnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback OnAfterDeleteInterface[Library]
	OnAfterLibraryReadCallback   OnAfterReadInterface[Library]

	Messages                map[*Message]struct{}
	Messages_instance       map[*Message]*Message
	Messages_mapString      map[string]*Message
	MessageOrder            uint
	Message_stagedOrder     map[*Message]uint
	Message_orderStaged     map[uint]*Message
	Messages_reference      map[*Message]*Message
	Messages_referenceOrder map[*Message]uint

	// insertion point for slice of pointers maps
	OnAfterMessageCreateCallback OnAfterCreateInterface[Message]
	OnAfterMessageUpdateCallback OnAfterUpdateInterface[Message]
	OnAfterMessageDeleteCallback OnAfterDeleteInterface[Message]
	OnAfterMessageReadCallback   OnAfterReadInterface[Message]

	MessageTypes                map[*MessageType]struct{}
	MessageTypes_instance       map[*MessageType]*MessageType
	MessageTypes_mapString      map[string]*MessageType
	MessageTypeOrder            uint
	MessageType_stagedOrder     map[*MessageType]uint
	MessageType_orderStaged     map[uint]*MessageType
	MessageTypes_reference      map[*MessageType]*MessageType
	MessageTypes_referenceOrder map[*MessageType]uint

	// insertion point for slice of pointers maps
	OnAfterMessageTypeCreateCallback OnAfterCreateInterface[MessageType]
	OnAfterMessageTypeUpdateCallback OnAfterUpdateInterface[MessageType]
	OnAfterMessageTypeDeleteCallback OnAfterDeleteInterface[MessageType]
	OnAfterMessageTypeReadCallback   OnAfterReadInterface[MessageType]

	Notes                map[*Note]struct{}
	Notes_instance       map[*Note]*Note
	Notes_mapString      map[string]*Note
	NoteOrder            uint
	Note_stagedOrder     map[*Note]uint
	Note_orderStaged     map[uint]*Note
	Notes_reference      map[*Note]*Note
	Notes_referenceOrder map[*Note]uint

	// insertion point for slice of pointers maps
	Note_States_reverseMap map[*State]*Note

	Note_Transitions_reverseMap map[*Transition]*Note

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

	NoteStateShapes                map[*NoteStateShape]struct{}
	NoteStateShapes_instance       map[*NoteStateShape]*NoteStateShape
	NoteStateShapes_mapString      map[string]*NoteStateShape
	NoteStateShapeOrder            uint
	NoteStateShape_stagedOrder     map[*NoteStateShape]uint
	NoteStateShape_orderStaged     map[uint]*NoteStateShape
	NoteStateShapes_reference      map[*NoteStateShape]*NoteStateShape
	NoteStateShapes_referenceOrder map[*NoteStateShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteStateShapeCreateCallback OnAfterCreateInterface[NoteStateShape]
	OnAfterNoteStateShapeUpdateCallback OnAfterUpdateInterface[NoteStateShape]
	OnAfterNoteStateShapeDeleteCallback OnAfterDeleteInterface[NoteStateShape]
	OnAfterNoteStateShapeReadCallback   OnAfterReadInterface[NoteStateShape]

	NoteTransitionShapes                map[*NoteTransitionShape]struct{}
	NoteTransitionShapes_instance       map[*NoteTransitionShape]*NoteTransitionShape
	NoteTransitionShapes_mapString      map[string]*NoteTransitionShape
	NoteTransitionShapeOrder            uint
	NoteTransitionShape_stagedOrder     map[*NoteTransitionShape]uint
	NoteTransitionShape_orderStaged     map[uint]*NoteTransitionShape
	NoteTransitionShapes_reference      map[*NoteTransitionShape]*NoteTransitionShape
	NoteTransitionShapes_referenceOrder map[*NoteTransitionShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteTransitionShapeCreateCallback OnAfterCreateInterface[NoteTransitionShape]
	OnAfterNoteTransitionShapeUpdateCallback OnAfterUpdateInterface[NoteTransitionShape]
	OnAfterNoteTransitionShapeDeleteCallback OnAfterDeleteInterface[NoteTransitionShape]
	OnAfterNoteTransitionShapeReadCallback   OnAfterReadInterface[NoteTransitionShape]

	Objects                map[*Object]struct{}
	Objects_instance       map[*Object]*Object
	Objects_mapString      map[string]*Object
	ObjectOrder            uint
	Object_stagedOrder     map[*Object]uint
	Object_orderStaged     map[uint]*Object
	Objects_reference      map[*Object]*Object
	Objects_referenceOrder map[*Object]uint

	// insertion point for slice of pointers maps
	Object_Messages_reverseMap map[*Message]*Object

	OnAfterObjectCreateCallback OnAfterCreateInterface[Object]
	OnAfterObjectUpdateCallback OnAfterUpdateInterface[Object]
	OnAfterObjectDeleteCallback OnAfterDeleteInterface[Object]
	OnAfterObjectReadCallback   OnAfterReadInterface[Object]

	Roles                map[*Role]struct{}
	Roles_instance       map[*Role]*Role
	Roles_mapString      map[string]*Role
	RoleOrder            uint
	Role_stagedOrder     map[*Role]uint
	Role_orderStaged     map[uint]*Role
	Roles_reference      map[*Role]*Role
	Roles_referenceOrder map[*Role]uint

	// insertion point for slice of pointers maps
	Role_RolesWithSamePermissions_reverseMap map[*Role]*Role

	OnAfterRoleCreateCallback OnAfterCreateInterface[Role]
	OnAfterRoleUpdateCallback OnAfterUpdateInterface[Role]
	OnAfterRoleDeleteCallback OnAfterDeleteInterface[Role]
	OnAfterRoleReadCallback   OnAfterReadInterface[Role]

	States                map[*State]struct{}
	States_instance       map[*State]*State
	States_mapString      map[string]*State
	StateOrder            uint
	State_stagedOrder     map[*State]uint
	State_orderStaged     map[uint]*State
	States_reference      map[*State]*State
	States_referenceOrder map[*State]uint

	// insertion point for slice of pointers maps
	State_SubStates_reverseMap map[*State]*State

	State_Diagrams_reverseMap map[*Diagram]*State

	State_Activities_reverseMap map[*Activities]*State

	OnAfterStateCreateCallback OnAfterCreateInterface[State]
	OnAfterStateUpdateCallback OnAfterUpdateInterface[State]
	OnAfterStateDeleteCallback OnAfterDeleteInterface[State]
	OnAfterStateReadCallback   OnAfterReadInterface[State]

	StateMachines                map[*StateMachine]struct{}
	StateMachines_instance       map[*StateMachine]*StateMachine
	StateMachines_mapString      map[string]*StateMachine
	StateMachineOrder            uint
	StateMachine_stagedOrder     map[*StateMachine]uint
	StateMachine_orderStaged     map[uint]*StateMachine
	StateMachines_reference      map[*StateMachine]*StateMachine
	StateMachines_referenceOrder map[*StateMachine]uint

	// insertion point for slice of pointers maps
	StateMachine_States_reverseMap map[*State]*StateMachine

	StateMachine_Diagrams_reverseMap map[*Diagram]*StateMachine

	OnAfterStateMachineCreateCallback OnAfterCreateInterface[StateMachine]
	OnAfterStateMachineUpdateCallback OnAfterUpdateInterface[StateMachine]
	OnAfterStateMachineDeleteCallback OnAfterDeleteInterface[StateMachine]
	OnAfterStateMachineReadCallback   OnAfterReadInterface[StateMachine]

	StateShapes                map[*StateShape]struct{}
	StateShapes_instance       map[*StateShape]*StateShape
	StateShapes_mapString      map[string]*StateShape
	StateShapeOrder            uint
	StateShape_stagedOrder     map[*StateShape]uint
	StateShape_orderStaged     map[uint]*StateShape
	StateShapes_reference      map[*StateShape]*StateShape
	StateShapes_referenceOrder map[*StateShape]uint

	// insertion point for slice of pointers maps
	OnAfterStateShapeCreateCallback OnAfterCreateInterface[StateShape]
	OnAfterStateShapeUpdateCallback OnAfterUpdateInterface[StateShape]
	OnAfterStateShapeDeleteCallback OnAfterDeleteInterface[StateShape]
	OnAfterStateShapeReadCallback   OnAfterReadInterface[StateShape]

	Transitions                map[*Transition]struct{}
	Transitions_instance       map[*Transition]*Transition
	Transitions_mapString      map[string]*Transition
	TransitionOrder            uint
	Transition_stagedOrder     map[*Transition]uint
	Transition_orderStaged     map[uint]*Transition
	Transitions_reference      map[*Transition]*Transition
	Transitions_referenceOrder map[*Transition]uint

	// insertion point for slice of pointers maps
	Transition_RolesWithPermissions_reverseMap map[*Role]*Transition

	Transition_GeneratedMessages_reverseMap map[*MessageType]*Transition

	Transition_Diagrams_reverseMap map[*Diagram]*Transition

	OnAfterTransitionCreateCallback OnAfterCreateInterface[Transition]
	OnAfterTransitionUpdateCallback OnAfterUpdateInterface[Transition]
	OnAfterTransitionDeleteCallback OnAfterDeleteInterface[Transition]
	OnAfterTransitionReadCallback   OnAfterReadInterface[Transition]

	Transition_Shapes                map[*Transition_Shape]struct{}
	Transition_Shapes_instance       map[*Transition_Shape]*Transition_Shape
	Transition_Shapes_mapString      map[string]*Transition_Shape
	Transition_ShapeOrder            uint
	Transition_Shape_stagedOrder     map[*Transition_Shape]uint
	Transition_Shape_orderStaged     map[uint]*Transition_Shape
	Transition_Shapes_reference      map[*Transition_Shape]*Transition_Shape
	Transition_Shapes_referenceOrder map[*Transition_Shape]uint

	// insertion point for slice of pointers maps
	OnAfterTransition_ShapeCreateCallback OnAfterCreateInterface[Transition_Shape]
	OnAfterTransition_ShapeUpdateCallback OnAfterUpdateInterface[Transition_Shape]
	OnAfterTransition_ShapeDeleteCallback OnAfterDeleteInterface[Transition_Shape]
	OnAfterTransition_ShapeReadCallback   OnAfterReadInterface[Transition_Shape]

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
	stage.Actions_reference = make(map[*Action]*Action)
	stage.Actions_instance = make(map[*Action]*Action)
	stage.Actions_referenceOrder = make(map[*Action]uint)

	stage.Activitiess_reference = make(map[*Activities]*Activities)
	stage.Activitiess_instance = make(map[*Activities]*Activities)
	stage.Activitiess_referenceOrder = make(map[*Activities]uint)

	stage.Architectures_reference = make(map[*Architecture]*Architecture)
	stage.Architectures_instance = make(map[*Architecture]*Architecture)
	stage.Architectures_referenceOrder = make(map[*Architecture]uint)

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint)

	stage.Guards_reference = make(map[*Guard]*Guard)
	stage.Guards_instance = make(map[*Guard]*Guard)
	stage.Guards_referenceOrder = make(map[*Guard]uint)

	stage.Kills_reference = make(map[*Kill]*Kill)
	stage.Kills_instance = make(map[*Kill]*Kill)
	stage.Kills_referenceOrder = make(map[*Kill]uint)

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_instance = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint)

	stage.Messages_reference = make(map[*Message]*Message)
	stage.Messages_instance = make(map[*Message]*Message)
	stage.Messages_referenceOrder = make(map[*Message]uint)

	stage.MessageTypes_reference = make(map[*MessageType]*MessageType)
	stage.MessageTypes_instance = make(map[*MessageType]*MessageType)
	stage.MessageTypes_referenceOrder = make(map[*MessageType]uint)

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_instance = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint)

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)

	stage.NoteStateShapes_reference = make(map[*NoteStateShape]*NoteStateShape)
	stage.NoteStateShapes_instance = make(map[*NoteStateShape]*NoteStateShape)
	stage.NoteStateShapes_referenceOrder = make(map[*NoteStateShape]uint)

	stage.NoteTransitionShapes_reference = make(map[*NoteTransitionShape]*NoteTransitionShape)
	stage.NoteTransitionShapes_instance = make(map[*NoteTransitionShape]*NoteTransitionShape)
	stage.NoteTransitionShapes_referenceOrder = make(map[*NoteTransitionShape]uint)

	stage.Objects_reference = make(map[*Object]*Object)
	stage.Objects_instance = make(map[*Object]*Object)
	stage.Objects_referenceOrder = make(map[*Object]uint)

	stage.Roles_reference = make(map[*Role]*Role)
	stage.Roles_instance = make(map[*Role]*Role)
	stage.Roles_referenceOrder = make(map[*Role]uint)

	stage.States_reference = make(map[*State]*State)
	stage.States_instance = make(map[*State]*State)
	stage.States_referenceOrder = make(map[*State]uint)

	stage.StateMachines_reference = make(map[*StateMachine]*StateMachine)
	stage.StateMachines_instance = make(map[*StateMachine]*StateMachine)
	stage.StateMachines_referenceOrder = make(map[*StateMachine]uint)

	stage.StateShapes_reference = make(map[*StateShape]*StateShape)
	stage.StateShapes_instance = make(map[*StateShape]*StateShape)
	stage.StateShapes_referenceOrder = make(map[*StateShape]uint)

	stage.Transitions_reference = make(map[*Transition]*Transition)
	stage.Transitions_instance = make(map[*Transition]*Transition)
	stage.Transitions_referenceOrder = make(map[*Transition]uint)

	stage.Transition_Shapes_reference = make(map[*Transition_Shape]*Transition_Shape)
	stage.Transition_Shapes_instance = make(map[*Transition_Shape]*Transition_Shape)
	stage.Transition_Shapes_referenceOrder = make(map[*Transition_Shape]uint)

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
	var maxActionOrder uint
	var foundAction bool
	for _, order := range stage.Action_stagedOrder {
		if !foundAction || order > maxActionOrder {
			maxActionOrder = order
			foundAction = true
		}
	}
	if foundAction {
		stage.ActionOrder = maxActionOrder + 1
	} else {
		stage.ActionOrder = 0
	}

	var maxActivitiesOrder uint
	var foundActivities bool
	for _, order := range stage.Activities_stagedOrder {
		if !foundActivities || order > maxActivitiesOrder {
			maxActivitiesOrder = order
			foundActivities = true
		}
	}
	if foundActivities {
		stage.ActivitiesOrder = maxActivitiesOrder + 1
	} else {
		stage.ActivitiesOrder = 0
	}

	var maxArchitectureOrder uint
	var foundArchitecture bool
	for _, order := range stage.Architecture_stagedOrder {
		if !foundArchitecture || order > maxArchitectureOrder {
			maxArchitectureOrder = order
			foundArchitecture = true
		}
	}
	if foundArchitecture {
		stage.ArchitectureOrder = maxArchitectureOrder + 1
	} else {
		stage.ArchitectureOrder = 0
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

	var maxGuardOrder uint
	var foundGuard bool
	for _, order := range stage.Guard_stagedOrder {
		if !foundGuard || order > maxGuardOrder {
			maxGuardOrder = order
			foundGuard = true
		}
	}
	if foundGuard {
		stage.GuardOrder = maxGuardOrder + 1
	} else {
		stage.GuardOrder = 0
	}

	var maxKillOrder uint
	var foundKill bool
	for _, order := range stage.Kill_stagedOrder {
		if !foundKill || order > maxKillOrder {
			maxKillOrder = order
			foundKill = true
		}
	}
	if foundKill {
		stage.KillOrder = maxKillOrder + 1
	} else {
		stage.KillOrder = 0
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

	var maxMessageOrder uint
	var foundMessage bool
	for _, order := range stage.Message_stagedOrder {
		if !foundMessage || order > maxMessageOrder {
			maxMessageOrder = order
			foundMessage = true
		}
	}
	if foundMessage {
		stage.MessageOrder = maxMessageOrder + 1
	} else {
		stage.MessageOrder = 0
	}

	var maxMessageTypeOrder uint
	var foundMessageType bool
	for _, order := range stage.MessageType_stagedOrder {
		if !foundMessageType || order > maxMessageTypeOrder {
			maxMessageTypeOrder = order
			foundMessageType = true
		}
	}
	if foundMessageType {
		stage.MessageTypeOrder = maxMessageTypeOrder + 1
	} else {
		stage.MessageTypeOrder = 0
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

	var maxNoteStateShapeOrder uint
	var foundNoteStateShape bool
	for _, order := range stage.NoteStateShape_stagedOrder {
		if !foundNoteStateShape || order > maxNoteStateShapeOrder {
			maxNoteStateShapeOrder = order
			foundNoteStateShape = true
		}
	}
	if foundNoteStateShape {
		stage.NoteStateShapeOrder = maxNoteStateShapeOrder + 1
	} else {
		stage.NoteStateShapeOrder = 0
	}

	var maxNoteTransitionShapeOrder uint
	var foundNoteTransitionShape bool
	for _, order := range stage.NoteTransitionShape_stagedOrder {
		if !foundNoteTransitionShape || order > maxNoteTransitionShapeOrder {
			maxNoteTransitionShapeOrder = order
			foundNoteTransitionShape = true
		}
	}
	if foundNoteTransitionShape {
		stage.NoteTransitionShapeOrder = maxNoteTransitionShapeOrder + 1
	} else {
		stage.NoteTransitionShapeOrder = 0
	}

	var maxObjectOrder uint
	var foundObject bool
	for _, order := range stage.Object_stagedOrder {
		if !foundObject || order > maxObjectOrder {
			maxObjectOrder = order
			foundObject = true
		}
	}
	if foundObject {
		stage.ObjectOrder = maxObjectOrder + 1
	} else {
		stage.ObjectOrder = 0
	}

	var maxRoleOrder uint
	var foundRole bool
	for _, order := range stage.Role_stagedOrder {
		if !foundRole || order > maxRoleOrder {
			maxRoleOrder = order
			foundRole = true
		}
	}
	if foundRole {
		stage.RoleOrder = maxRoleOrder + 1
	} else {
		stage.RoleOrder = 0
	}

	var maxStateOrder uint
	var foundState bool
	for _, order := range stage.State_stagedOrder {
		if !foundState || order > maxStateOrder {
			maxStateOrder = order
			foundState = true
		}
	}
	if foundState {
		stage.StateOrder = maxStateOrder + 1
	} else {
		stage.StateOrder = 0
	}

	var maxStateMachineOrder uint
	var foundStateMachine bool
	for _, order := range stage.StateMachine_stagedOrder {
		if !foundStateMachine || order > maxStateMachineOrder {
			maxStateMachineOrder = order
			foundStateMachine = true
		}
	}
	if foundStateMachine {
		stage.StateMachineOrder = maxStateMachineOrder + 1
	} else {
		stage.StateMachineOrder = 0
	}

	var maxStateShapeOrder uint
	var foundStateShape bool
	for _, order := range stage.StateShape_stagedOrder {
		if !foundStateShape || order > maxStateShapeOrder {
			maxStateShapeOrder = order
			foundStateShape = true
		}
	}
	if foundStateShape {
		stage.StateShapeOrder = maxStateShapeOrder + 1
	} else {
		stage.StateShapeOrder = 0
	}

	var maxTransitionOrder uint
	var foundTransition bool
	for _, order := range stage.Transition_stagedOrder {
		if !foundTransition || order > maxTransitionOrder {
			maxTransitionOrder = order
			foundTransition = true
		}
	}
	if foundTransition {
		stage.TransitionOrder = maxTransitionOrder + 1
	} else {
		stage.TransitionOrder = 0
	}

	var maxTransition_ShapeOrder uint
	var foundTransition_Shape bool
	for _, order := range stage.Transition_Shape_stagedOrder {
		if !foundTransition_Shape || order > maxTransition_ShapeOrder {
			maxTransition_ShapeOrder = order
			foundTransition_Shape = true
		}
	}
	if foundTransition_Shape {
		stage.Transition_ShapeOrder = maxTransition_ShapeOrder + 1
	} else {
		stage.Transition_ShapeOrder = 0
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
	case *Action:
		tmp := GetStructInstancesByOrder(stage.Actions, stage.Action_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Action implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Activities:
		tmp := GetStructInstancesByOrder(stage.Activitiess, stage.Activities_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Activities implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Architecture:
		tmp := GetStructInstancesByOrder(stage.Architectures, stage.Architecture_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Architecture implements.
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
	case *Guard:
		tmp := GetStructInstancesByOrder(stage.Guards, stage.Guard_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Guard implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Kill:
		tmp := GetStructInstancesByOrder(stage.Kills, stage.Kill_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Kill implements.
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
	case *Message:
		tmp := GetStructInstancesByOrder(stage.Messages, stage.Message_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Message implements.
			res = append(res, any(v).(T))
		}
		return res
	case *MessageType:
		tmp := GetStructInstancesByOrder(stage.MessageTypes, stage.MessageType_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *MessageType implements.
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
	case *NoteStateShape:
		tmp := GetStructInstancesByOrder(stage.NoteStateShapes, stage.NoteStateShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteStateShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *NoteTransitionShape:
		tmp := GetStructInstancesByOrder(stage.NoteTransitionShapes, stage.NoteTransitionShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *NoteTransitionShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Object:
		tmp := GetStructInstancesByOrder(stage.Objects, stage.Object_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Object implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Role:
		tmp := GetStructInstancesByOrder(stage.Roles, stage.Role_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Role implements.
			res = append(res, any(v).(T))
		}
		return res
	case *State:
		tmp := GetStructInstancesByOrder(stage.States, stage.State_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *State implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StateMachine:
		tmp := GetStructInstancesByOrder(stage.StateMachines, stage.StateMachine_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StateMachine implements.
			res = append(res, any(v).(T))
		}
		return res
	case *StateShape:
		tmp := GetStructInstancesByOrder(stage.StateShapes, stage.StateShape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *StateShape implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Transition:
		tmp := GetStructInstancesByOrder(stage.Transitions, stage.Transition_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Transition implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Transition_Shape:
		tmp := GetStructInstancesByOrder(stage.Transition_Shapes, stage.Transition_Shape_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Transition_Shape implements.
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
	case "Action":
		res = GetNamedStructInstances(stage.Actions, stage.Action_stagedOrder)
	case "Activities":
		res = GetNamedStructInstances(stage.Activitiess, stage.Activities_stagedOrder)
	case "Architecture":
		res = GetNamedStructInstances(stage.Architectures, stage.Architecture_stagedOrder)
	case "Diagram":
		res = GetNamedStructInstances(stage.Diagrams, stage.Diagram_stagedOrder)
	case "Guard":
		res = GetNamedStructInstances(stage.Guards, stage.Guard_stagedOrder)
	case "Kill":
		res = GetNamedStructInstances(stage.Kills, stage.Kill_stagedOrder)
	case "Library":
		res = GetNamedStructInstances(stage.Librarys, stage.Library_stagedOrder)
	case "Message":
		res = GetNamedStructInstances(stage.Messages, stage.Message_stagedOrder)
	case "MessageType":
		res = GetNamedStructInstances(stage.MessageTypes, stage.MessageType_stagedOrder)
	case "Note":
		res = GetNamedStructInstances(stage.Notes, stage.Note_stagedOrder)
	case "NoteShape":
		res = GetNamedStructInstances(stage.NoteShapes, stage.NoteShape_stagedOrder)
	case "NoteStateShape":
		res = GetNamedStructInstances(stage.NoteStateShapes, stage.NoteStateShape_stagedOrder)
	case "NoteTransitionShape":
		res = GetNamedStructInstances(stage.NoteTransitionShapes, stage.NoteTransitionShape_stagedOrder)
	case "Object":
		res = GetNamedStructInstances(stage.Objects, stage.Object_stagedOrder)
	case "Role":
		res = GetNamedStructInstances(stage.Roles, stage.Role_stagedOrder)
	case "State":
		res = GetNamedStructInstances(stage.States, stage.State_stagedOrder)
	case "StateMachine":
		res = GetNamedStructInstances(stage.StateMachines, stage.StateMachine_stagedOrder)
	case "StateShape":
		res = GetNamedStructInstances(stage.StateShapes, stage.StateShape_stagedOrder)
	case "Transition":
		res = GetNamedStructInstances(stage.Transitions, stage.Transition_stagedOrder)
	case "Transition_Shape":
		res = GetNamedStructInstances(stage.Transition_Shapes, stage.Transition_Shape_stagedOrder)
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
	return "github.com/fullstack-lang/gong/dsm/statemachines/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return statemachines_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return statemachines_go.GoDiagramsDir
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
	CommitAction(action *Action)
	CheckoutAction(action *Action)
	CommitActivities(activities *Activities)
	CheckoutActivities(activities *Activities)
	CommitArchitecture(architecture *Architecture)
	CheckoutArchitecture(architecture *Architecture)
	CommitDiagram(diagram *Diagram)
	CheckoutDiagram(diagram *Diagram)
	CommitGuard(guard *Guard)
	CheckoutGuard(guard *Guard)
	CommitKill(kill *Kill)
	CheckoutKill(kill *Kill)
	CommitLibrary(library *Library)
	CheckoutLibrary(library *Library)
	CommitMessage(message *Message)
	CheckoutMessage(message *Message)
	CommitMessageType(messagetype *MessageType)
	CheckoutMessageType(messagetype *MessageType)
	CommitNote(note *Note)
	CheckoutNote(note *Note)
	CommitNoteShape(noteshape *NoteShape)
	CheckoutNoteShape(noteshape *NoteShape)
	CommitNoteStateShape(notestateshape *NoteStateShape)
	CheckoutNoteStateShape(notestateshape *NoteStateShape)
	CommitNoteTransitionShape(notetransitionshape *NoteTransitionShape)
	CheckoutNoteTransitionShape(notetransitionshape *NoteTransitionShape)
	CommitObject(object *Object)
	CheckoutObject(object *Object)
	CommitRole(role *Role)
	CheckoutRole(role *Role)
	CommitState(state *State)
	CheckoutState(state *State)
	CommitStateMachine(statemachine *StateMachine)
	CheckoutStateMachine(statemachine *StateMachine)
	CommitStateShape(stateshape *StateShape)
	CheckoutStateShape(stateshape *StateShape)
	CommitTransition(transition *Transition)
	CheckoutTransition(transition *Transition)
	CommitTransition_Shape(transition_shape *Transition_Shape)
	CheckoutTransition_Shape(transition_shape *Transition_Shape)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		Actions:           make(map[*Action]struct{}),
		Actions_mapString: make(map[string]*Action),

		Activitiess:           make(map[*Activities]struct{}),
		Activitiess_mapString: make(map[string]*Activities),

		Architectures:           make(map[*Architecture]struct{}),
		Architectures_mapString: make(map[string]*Architecture),

		Diagrams:           make(map[*Diagram]struct{}),
		Diagrams_mapString: make(map[string]*Diagram),

		Guards:           make(map[*Guard]struct{}),
		Guards_mapString: make(map[string]*Guard),

		Kills:           make(map[*Kill]struct{}),
		Kills_mapString: make(map[string]*Kill),

		Librarys:           make(map[*Library]struct{}),
		Librarys_mapString: make(map[string]*Library),

		Messages:           make(map[*Message]struct{}),
		Messages_mapString: make(map[string]*Message),

		MessageTypes:           make(map[*MessageType]struct{}),
		MessageTypes_mapString: make(map[string]*MessageType),

		Notes:           make(map[*Note]struct{}),
		Notes_mapString: make(map[string]*Note),

		NoteShapes:           make(map[*NoteShape]struct{}),
		NoteShapes_mapString: make(map[string]*NoteShape),

		NoteStateShapes:           make(map[*NoteStateShape]struct{}),
		NoteStateShapes_mapString: make(map[string]*NoteStateShape),

		NoteTransitionShapes:           make(map[*NoteTransitionShape]struct{}),
		NoteTransitionShapes_mapString: make(map[string]*NoteTransitionShape),

		Objects:           make(map[*Object]struct{}),
		Objects_mapString: make(map[string]*Object),

		Roles:           make(map[*Role]struct{}),
		Roles_mapString: make(map[string]*Role),

		States:           make(map[*State]struct{}),
		States_mapString: make(map[string]*State),

		StateMachines:           make(map[*StateMachine]struct{}),
		StateMachines_mapString: make(map[string]*StateMachine),

		StateShapes:           make(map[*StateShape]struct{}),
		StateShapes_mapString: make(map[string]*StateShape),

		Transitions:           make(map[*Transition]struct{}),
		Transitions_mapString: make(map[string]*Transition),

		Transition_Shapes:           make(map[*Transition_Shape]struct{}),
		Transition_Shapes_mapString: make(map[string]*Transition_Shape),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Action_stagedOrder: make(map[*Action]uint),
		Action_orderStaged: make(map[uint]*Action),
		Actions_reference:  make(map[*Action]*Action),

		Activities_stagedOrder: make(map[*Activities]uint),
		Activities_orderStaged: make(map[uint]*Activities),
		Activitiess_reference:  make(map[*Activities]*Activities),

		Architecture_stagedOrder: make(map[*Architecture]uint),
		Architecture_orderStaged: make(map[uint]*Architecture),
		Architectures_reference:  make(map[*Architecture]*Architecture),

		Diagram_stagedOrder: make(map[*Diagram]uint),
		Diagram_orderStaged: make(map[uint]*Diagram),
		Diagrams_reference:  make(map[*Diagram]*Diagram),

		Guard_stagedOrder: make(map[*Guard]uint),
		Guard_orderStaged: make(map[uint]*Guard),
		Guards_reference:  make(map[*Guard]*Guard),

		Kill_stagedOrder: make(map[*Kill]uint),
		Kill_orderStaged: make(map[uint]*Kill),
		Kills_reference:  make(map[*Kill]*Kill),

		Library_stagedOrder: make(map[*Library]uint),
		Library_orderStaged: make(map[uint]*Library),
		Librarys_reference:  make(map[*Library]*Library),

		Message_stagedOrder: make(map[*Message]uint),
		Message_orderStaged: make(map[uint]*Message),
		Messages_reference:  make(map[*Message]*Message),

		MessageType_stagedOrder: make(map[*MessageType]uint),
		MessageType_orderStaged: make(map[uint]*MessageType),
		MessageTypes_reference:  make(map[*MessageType]*MessageType),

		Note_stagedOrder: make(map[*Note]uint),
		Note_orderStaged: make(map[uint]*Note),
		Notes_reference:  make(map[*Note]*Note),

		NoteShape_stagedOrder: make(map[*NoteShape]uint),
		NoteShape_orderStaged: make(map[uint]*NoteShape),
		NoteShapes_reference:  make(map[*NoteShape]*NoteShape),

		NoteStateShape_stagedOrder: make(map[*NoteStateShape]uint),
		NoteStateShape_orderStaged: make(map[uint]*NoteStateShape),
		NoteStateShapes_reference:  make(map[*NoteStateShape]*NoteStateShape),

		NoteTransitionShape_stagedOrder: make(map[*NoteTransitionShape]uint),
		NoteTransitionShape_orderStaged: make(map[uint]*NoteTransitionShape),
		NoteTransitionShapes_reference:  make(map[*NoteTransitionShape]*NoteTransitionShape),

		Object_stagedOrder: make(map[*Object]uint),
		Object_orderStaged: make(map[uint]*Object),
		Objects_reference:  make(map[*Object]*Object),

		Role_stagedOrder: make(map[*Role]uint),
		Role_orderStaged: make(map[uint]*Role),
		Roles_reference:  make(map[*Role]*Role),

		State_stagedOrder: make(map[*State]uint),
		State_orderStaged: make(map[uint]*State),
		States_reference:  make(map[*State]*State),

		StateMachine_stagedOrder: make(map[*StateMachine]uint),
		StateMachine_orderStaged: make(map[uint]*StateMachine),
		StateMachines_reference:  make(map[*StateMachine]*StateMachine),

		StateShape_stagedOrder: make(map[*StateShape]uint),
		StateShape_orderStaged: make(map[uint]*StateShape),
		StateShapes_reference:  make(map[*StateShape]*StateShape),

		Transition_stagedOrder: make(map[*Transition]uint),
		Transition_orderStaged: make(map[uint]*Transition),
		Transitions_reference:  make(map[*Transition]*Transition),

		Transition_Shape_stagedOrder: make(map[*Transition_Shape]uint),
		Transition_Shape_orderStaged: make(map[uint]*Transition_Shape),
		Transition_Shapes_reference:  make(map[*Transition_Shape]*Transition_Shape),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Action": &ActionUnmarshaller{},

			"Activities": &ActivitiesUnmarshaller{},

			"Architecture": &ArchitectureUnmarshaller{},

			"Diagram": &DiagramUnmarshaller{},

			"Guard": &GuardUnmarshaller{},

			"Kill": &KillUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Message": &MessageUnmarshaller{},

			"MessageType": &MessageTypeUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"NoteStateShape": &NoteStateShapeUnmarshaller{},

			"NoteTransitionShape": &NoteTransitionShapeUnmarshaller{},

			"Object": &ObjectUnmarshaller{},

			"Role": &RoleUnmarshaller{},

			"State": &StateUnmarshaller{},

			"StateMachine": &StateMachineUnmarshaller{},

			"StateShape": &StateShapeUnmarshaller{},

			"Transition": &TransitionUnmarshaller{},

			"Transition_Shape": &Transition_ShapeUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Action"},
			{name: "Activities"},
			{name: "Architecture"},
			{name: "Diagram"},
			{name: "Guard"},
			{name: "Kill"},
			{name: "Library"},
			{name: "Message"},
			{name: "MessageType"},
			{name: "Note"},
			{name: "NoteShape"},
			{name: "NoteStateShape"},
			{name: "NoteTransitionShape"},
			{name: "Object"},
			{name: "Role"},
			{name: "State"},
			{name: "StateMachine"},
			{name: "StateShape"},
			{name: "Transition"},
			{name: "Transition_Shape"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Action:
		return stage.Action_stagedOrder[instance]
	case *Activities:
		return stage.Activities_stagedOrder[instance]
	case *Architecture:
		return stage.Architecture_stagedOrder[instance]
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Guard:
		return stage.Guard_stagedOrder[instance]
	case *Kill:
		return stage.Kill_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Message:
		return stage.Message_stagedOrder[instance]
	case *MessageType:
		return stage.MessageType_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *NoteStateShape:
		return stage.NoteStateShape_stagedOrder[instance]
	case *NoteTransitionShape:
		return stage.NoteTransitionShape_stagedOrder[instance]
	case *Object:
		return stage.Object_stagedOrder[instance]
	case *Role:
		return stage.Role_stagedOrder[instance]
	case *State:
		return stage.State_stagedOrder[instance]
	case *StateMachine:
		return stage.StateMachine_stagedOrder[instance]
	case *StateShape:
		return stage.StateShape_stagedOrder[instance]
	case *Transition:
		return stage.Transition_stagedOrder[instance]
	case *Transition_Shape:
		return stage.Transition_Shape_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *Action:
		return any(stage.Action_orderStaged[order]).(Type)
	case *Activities:
		return any(stage.Activities_orderStaged[order]).(Type)
	case *Architecture:
		return any(stage.Architecture_orderStaged[order]).(Type)
	case *Diagram:
		return any(stage.Diagram_orderStaged[order]).(Type)
	case *Guard:
		return any(stage.Guard_orderStaged[order]).(Type)
	case *Kill:
		return any(stage.Kill_orderStaged[order]).(Type)
	case *Library:
		return any(stage.Library_orderStaged[order]).(Type)
	case *Message:
		return any(stage.Message_orderStaged[order]).(Type)
	case *MessageType:
		return any(stage.MessageType_orderStaged[order]).(Type)
	case *Note:
		return any(stage.Note_orderStaged[order]).(Type)
	case *NoteShape:
		return any(stage.NoteShape_orderStaged[order]).(Type)
	case *NoteStateShape:
		return any(stage.NoteStateShape_orderStaged[order]).(Type)
	case *NoteTransitionShape:
		return any(stage.NoteTransitionShape_orderStaged[order]).(Type)
	case *Object:
		return any(stage.Object_orderStaged[order]).(Type)
	case *Role:
		return any(stage.Role_orderStaged[order]).(Type)
	case *State:
		return any(stage.State_orderStaged[order]).(Type)
	case *StateMachine:
		return any(stage.StateMachine_orderStaged[order]).(Type)
	case *StateShape:
		return any(stage.StateShape_orderStaged[order]).(Type)
	case *Transition:
		return any(stage.Transition_orderStaged[order]).(Type)
	case *Transition_Shape:
		return any(stage.Transition_Shape_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Action:
		return stage.Action_stagedOrder[instance]
	case *Activities:
		return stage.Activities_stagedOrder[instance]
	case *Architecture:
		return stage.Architecture_stagedOrder[instance]
	case *Diagram:
		return stage.Diagram_stagedOrder[instance]
	case *Guard:
		return stage.Guard_stagedOrder[instance]
	case *Kill:
		return stage.Kill_stagedOrder[instance]
	case *Library:
		return stage.Library_stagedOrder[instance]
	case *Message:
		return stage.Message_stagedOrder[instance]
	case *MessageType:
		return stage.MessageType_stagedOrder[instance]
	case *Note:
		return stage.Note_stagedOrder[instance]
	case *NoteShape:
		return stage.NoteShape_stagedOrder[instance]
	case *NoteStateShape:
		return stage.NoteStateShape_stagedOrder[instance]
	case *NoteTransitionShape:
		return stage.NoteTransitionShape_stagedOrder[instance]
	case *Object:
		return stage.Object_stagedOrder[instance]
	case *Role:
		return stage.Role_stagedOrder[instance]
	case *State:
		return stage.State_stagedOrder[instance]
	case *StateMachine:
		return stage.StateMachine_stagedOrder[instance]
	case *StateShape:
		return stage.StateShape_stagedOrder[instance]
	case *Transition:
		return stage.Transition_stagedOrder[instance]
	case *Transition_Shape:
		return stage.Transition_Shape_stagedOrder[instance]
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
	stage.Map_GongStructName_InstancesNb["Action"] = len(stage.Actions)
	stage.Map_GongStructName_InstancesNb["Activities"] = len(stage.Activitiess)
	stage.Map_GongStructName_InstancesNb["Architecture"] = len(stage.Architectures)
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Guard"] = len(stage.Guards)
	stage.Map_GongStructName_InstancesNb["Kill"] = len(stage.Kills)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Message"] = len(stage.Messages)
	stage.Map_GongStructName_InstancesNb["MessageType"] = len(stage.MessageTypes)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteStateShape"] = len(stage.NoteStateShapes)
	stage.Map_GongStructName_InstancesNb["NoteTransitionShape"] = len(stage.NoteTransitionShapes)
	stage.Map_GongStructName_InstancesNb["Object"] = len(stage.Objects)
	stage.Map_GongStructName_InstancesNb["Role"] = len(stage.Roles)
	stage.Map_GongStructName_InstancesNb["State"] = len(stage.States)
	stage.Map_GongStructName_InstancesNb["StateMachine"] = len(stage.StateMachines)
	stage.Map_GongStructName_InstancesNb["StateShape"] = len(stage.StateShapes)
	stage.Map_GongStructName_InstancesNb["Transition"] = len(stage.Transitions)
	stage.Map_GongStructName_InstancesNb["Transition_Shape"] = len(stage.Transition_Shapes)
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
// Stage puts action to the model stage
func (action *Action) Stage(stage *Stage) *Action {
	if _, ok := stage.Actions[action]; !ok {
		stage.Actions[action] = struct{}{}
		stage.Action_stagedOrder[action] = stage.ActionOrder
		stage.Action_orderStaged[stage.ActionOrder] = action
		stage.ActionOrder++
	}
	stage.Actions_mapString[action.Name] = action

	return action
}

// StagePreserveOrder puts action to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActionOrder
// - update stage.ActionOrder accordingly
func (action *Action) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Actions[action]; !ok {
		stage.Actions[action] = struct{}{}

		if order > stage.ActionOrder {
			stage.ActionOrder = order
		}
		stage.Action_stagedOrder[action] = order
		stage.Action_orderStaged[order] = action
		stage.ActionOrder++
	}
	stage.Actions_mapString[action.Name] = action
}

// Unstage removes action off the model stage
func (action *Action) Unstage(stage *Stage) *Action {
	delete(stage.Actions, action)
	// issue1150
	// delete(stage.Action_stagedOrder, action)
	delete(stage.Actions_mapString, action.Name)

	return action
}

// UnstageVoid removes action off the model stage
func (action *Action) UnstageVoid(stage *Stage) {
	delete(stage.Actions, action)
	// issue1150
	// delete(stage.Action_stagedOrder, action)
	delete(stage.Actions_mapString, action.Name)
}

// commit action to the back repo (if it is already staged)
func (action *Action) Commit(stage *Stage) *Action {
	if _, ok := stage.Actions[action]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAction(action)
		}
	}
	return action
}

func (action *Action) CommitVoid(stage *Stage) {
	action.Commit(stage)
}

func (action *Action) StageVoid(stage *Stage) {
	action.Stage(stage)
}

// Checkout action to the back repo (if it is already staged)
func (action *Action) Checkout(stage *Stage) *Action {
	if _, ok := stage.Actions[action]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAction(action)
		}
	}
	return action
}

// for satisfaction of GongStruct interface
func (action *Action) GetName() (res string) {
	return action.Name
}

// for satisfaction of GongStruct interface
func (action *Action) SetName(name string) {
	action.Name = name
}

// Stage puts activities to the model stage
func (activities *Activities) Stage(stage *Stage) *Activities {
	if _, ok := stage.Activitiess[activities]; !ok {
		stage.Activitiess[activities] = struct{}{}
		stage.Activities_stagedOrder[activities] = stage.ActivitiesOrder
		stage.Activities_orderStaged[stage.ActivitiesOrder] = activities
		stage.ActivitiesOrder++
	}
	stage.Activitiess_mapString[activities.Name] = activities

	return activities
}

// StagePreserveOrder puts activities to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActivitiesOrder
// - update stage.ActivitiesOrder accordingly
func (activities *Activities) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Activitiess[activities]; !ok {
		stage.Activitiess[activities] = struct{}{}

		if order > stage.ActivitiesOrder {
			stage.ActivitiesOrder = order
		}
		stage.Activities_stagedOrder[activities] = order
		stage.Activities_orderStaged[order] = activities
		stage.ActivitiesOrder++
	}
	stage.Activitiess_mapString[activities.Name] = activities
}

// Unstage removes activities off the model stage
func (activities *Activities) Unstage(stage *Stage) *Activities {
	delete(stage.Activitiess, activities)
	// issue1150
	// delete(stage.Activities_stagedOrder, activities)
	delete(stage.Activitiess_mapString, activities.Name)

	return activities
}

// UnstageVoid removes activities off the model stage
func (activities *Activities) UnstageVoid(stage *Stage) {
	delete(stage.Activitiess, activities)
	// issue1150
	// delete(stage.Activities_stagedOrder, activities)
	delete(stage.Activitiess_mapString, activities.Name)
}

// commit activities to the back repo (if it is already staged)
func (activities *Activities) Commit(stage *Stage) *Activities {
	if _, ok := stage.Activitiess[activities]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitActivities(activities)
		}
	}
	return activities
}

func (activities *Activities) CommitVoid(stage *Stage) {
	activities.Commit(stage)
}

func (activities *Activities) StageVoid(stage *Stage) {
	activities.Stage(stage)
}

// Checkout activities to the back repo (if it is already staged)
func (activities *Activities) Checkout(stage *Stage) *Activities {
	if _, ok := stage.Activitiess[activities]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutActivities(activities)
		}
	}
	return activities
}

// for satisfaction of GongStruct interface
func (activities *Activities) GetName() (res string) {
	return activities.Name
}

// for satisfaction of GongStruct interface
func (activities *Activities) SetName(name string) {
	activities.Name = name
}

// Stage puts architecture to the model stage
func (architecture *Architecture) Stage(stage *Stage) *Architecture {
	if _, ok := stage.Architectures[architecture]; !ok {
		stage.Architectures[architecture] = struct{}{}
		stage.Architecture_stagedOrder[architecture] = stage.ArchitectureOrder
		stage.Architecture_orderStaged[stage.ArchitectureOrder] = architecture
		stage.ArchitectureOrder++
	}
	stage.Architectures_mapString[architecture.Name] = architecture

	return architecture
}

// StagePreserveOrder puts architecture to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArchitectureOrder
// - update stage.ArchitectureOrder accordingly
func (architecture *Architecture) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Architectures[architecture]; !ok {
		stage.Architectures[architecture] = struct{}{}

		if order > stage.ArchitectureOrder {
			stage.ArchitectureOrder = order
		}
		stage.Architecture_stagedOrder[architecture] = order
		stage.Architecture_orderStaged[order] = architecture
		stage.ArchitectureOrder++
	}
	stage.Architectures_mapString[architecture.Name] = architecture
}

// Unstage removes architecture off the model stage
func (architecture *Architecture) Unstage(stage *Stage) *Architecture {
	delete(stage.Architectures, architecture)
	// issue1150
	// delete(stage.Architecture_stagedOrder, architecture)
	delete(stage.Architectures_mapString, architecture.Name)

	return architecture
}

// UnstageVoid removes architecture off the model stage
func (architecture *Architecture) UnstageVoid(stage *Stage) {
	delete(stage.Architectures, architecture)
	// issue1150
	// delete(stage.Architecture_stagedOrder, architecture)
	delete(stage.Architectures_mapString, architecture.Name)
}

// commit architecture to the back repo (if it is already staged)
func (architecture *Architecture) Commit(stage *Stage) *Architecture {
	if _, ok := stage.Architectures[architecture]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitArchitecture(architecture)
		}
	}
	return architecture
}

func (architecture *Architecture) CommitVoid(stage *Stage) {
	architecture.Commit(stage)
}

func (architecture *Architecture) StageVoid(stage *Stage) {
	architecture.Stage(stage)
}

// Checkout architecture to the back repo (if it is already staged)
func (architecture *Architecture) Checkout(stage *Stage) *Architecture {
	if _, ok := stage.Architectures[architecture]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutArchitecture(architecture)
		}
	}
	return architecture
}

// for satisfaction of GongStruct interface
func (architecture *Architecture) GetName() (res string) {
	return architecture.Name
}

// for satisfaction of GongStruct interface
func (architecture *Architecture) SetName(name string) {
	architecture.Name = name
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

// Stage puts guard to the model stage
func (guard *Guard) Stage(stage *Stage) *Guard {
	if _, ok := stage.Guards[guard]; !ok {
		stage.Guards[guard] = struct{}{}
		stage.Guard_stagedOrder[guard] = stage.GuardOrder
		stage.Guard_orderStaged[stage.GuardOrder] = guard
		stage.GuardOrder++
	}
	stage.Guards_mapString[guard.Name] = guard

	return guard
}

// StagePreserveOrder puts guard to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GuardOrder
// - update stage.GuardOrder accordingly
func (guard *Guard) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Guards[guard]; !ok {
		stage.Guards[guard] = struct{}{}

		if order > stage.GuardOrder {
			stage.GuardOrder = order
		}
		stage.Guard_stagedOrder[guard] = order
		stage.Guard_orderStaged[order] = guard
		stage.GuardOrder++
	}
	stage.Guards_mapString[guard.Name] = guard
}

// Unstage removes guard off the model stage
func (guard *Guard) Unstage(stage *Stage) *Guard {
	delete(stage.Guards, guard)
	// issue1150
	// delete(stage.Guard_stagedOrder, guard)
	delete(stage.Guards_mapString, guard.Name)

	return guard
}

// UnstageVoid removes guard off the model stage
func (guard *Guard) UnstageVoid(stage *Stage) {
	delete(stage.Guards, guard)
	// issue1150
	// delete(stage.Guard_stagedOrder, guard)
	delete(stage.Guards_mapString, guard.Name)
}

// commit guard to the back repo (if it is already staged)
func (guard *Guard) Commit(stage *Stage) *Guard {
	if _, ok := stage.Guards[guard]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGuard(guard)
		}
	}
	return guard
}

func (guard *Guard) CommitVoid(stage *Stage) {
	guard.Commit(stage)
}

func (guard *Guard) StageVoid(stage *Stage) {
	guard.Stage(stage)
}

// Checkout guard to the back repo (if it is already staged)
func (guard *Guard) Checkout(stage *Stage) *Guard {
	if _, ok := stage.Guards[guard]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGuard(guard)
		}
	}
	return guard
}

// for satisfaction of GongStruct interface
func (guard *Guard) GetName() (res string) {
	return guard.Name
}

// for satisfaction of GongStruct interface
func (guard *Guard) SetName(name string) {
	guard.Name = name
}

// Stage puts kill to the model stage
func (kill *Kill) Stage(stage *Stage) *Kill {
	if _, ok := stage.Kills[kill]; !ok {
		stage.Kills[kill] = struct{}{}
		stage.Kill_stagedOrder[kill] = stage.KillOrder
		stage.Kill_orderStaged[stage.KillOrder] = kill
		stage.KillOrder++
	}
	stage.Kills_mapString[kill.Name] = kill

	return kill
}

// StagePreserveOrder puts kill to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.KillOrder
// - update stage.KillOrder accordingly
func (kill *Kill) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Kills[kill]; !ok {
		stage.Kills[kill] = struct{}{}

		if order > stage.KillOrder {
			stage.KillOrder = order
		}
		stage.Kill_stagedOrder[kill] = order
		stage.Kill_orderStaged[order] = kill
		stage.KillOrder++
	}
	stage.Kills_mapString[kill.Name] = kill
}

// Unstage removes kill off the model stage
func (kill *Kill) Unstage(stage *Stage) *Kill {
	delete(stage.Kills, kill)
	// issue1150
	// delete(stage.Kill_stagedOrder, kill)
	delete(stage.Kills_mapString, kill.Name)

	return kill
}

// UnstageVoid removes kill off the model stage
func (kill *Kill) UnstageVoid(stage *Stage) {
	delete(stage.Kills, kill)
	// issue1150
	// delete(stage.Kill_stagedOrder, kill)
	delete(stage.Kills_mapString, kill.Name)
}

// commit kill to the back repo (if it is already staged)
func (kill *Kill) Commit(stage *Stage) *Kill {
	if _, ok := stage.Kills[kill]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitKill(kill)
		}
	}
	return kill
}

func (kill *Kill) CommitVoid(stage *Stage) {
	kill.Commit(stage)
}

func (kill *Kill) StageVoid(stage *Stage) {
	kill.Stage(stage)
}

// Checkout kill to the back repo (if it is already staged)
func (kill *Kill) Checkout(stage *Stage) *Kill {
	if _, ok := stage.Kills[kill]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutKill(kill)
		}
	}
	return kill
}

// for satisfaction of GongStruct interface
func (kill *Kill) GetName() (res string) {
	return kill.Name
}

// for satisfaction of GongStruct interface
func (kill *Kill) SetName(name string) {
	kill.Name = name
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

// Stage puts message to the model stage
func (message *Message) Stage(stage *Stage) *Message {
	if _, ok := stage.Messages[message]; !ok {
		stage.Messages[message] = struct{}{}
		stage.Message_stagedOrder[message] = stage.MessageOrder
		stage.Message_orderStaged[stage.MessageOrder] = message
		stage.MessageOrder++
	}
	stage.Messages_mapString[message.Name] = message

	return message
}

// StagePreserveOrder puts message to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MessageOrder
// - update stage.MessageOrder accordingly
func (message *Message) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Messages[message]; !ok {
		stage.Messages[message] = struct{}{}

		if order > stage.MessageOrder {
			stage.MessageOrder = order
		}
		stage.Message_stagedOrder[message] = order
		stage.Message_orderStaged[order] = message
		stage.MessageOrder++
	}
	stage.Messages_mapString[message.Name] = message
}

// Unstage removes message off the model stage
func (message *Message) Unstage(stage *Stage) *Message {
	delete(stage.Messages, message)
	// issue1150
	// delete(stage.Message_stagedOrder, message)
	delete(stage.Messages_mapString, message.Name)

	return message
}

// UnstageVoid removes message off the model stage
func (message *Message) UnstageVoid(stage *Stage) {
	delete(stage.Messages, message)
	// issue1150
	// delete(stage.Message_stagedOrder, message)
	delete(stage.Messages_mapString, message.Name)
}

// commit message to the back repo (if it is already staged)
func (message *Message) Commit(stage *Stage) *Message {
	if _, ok := stage.Messages[message]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMessage(message)
		}
	}
	return message
}

func (message *Message) CommitVoid(stage *Stage) {
	message.Commit(stage)
}

func (message *Message) StageVoid(stage *Stage) {
	message.Stage(stage)
}

// Checkout message to the back repo (if it is already staged)
func (message *Message) Checkout(stage *Stage) *Message {
	if _, ok := stage.Messages[message]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMessage(message)
		}
	}
	return message
}

// for satisfaction of GongStruct interface
func (message *Message) GetName() (res string) {
	return message.Name
}

// for satisfaction of GongStruct interface
func (message *Message) SetName(name string) {
	message.Name = name
}

// Stage puts messagetype to the model stage
func (messagetype *MessageType) Stage(stage *Stage) *MessageType {
	if _, ok := stage.MessageTypes[messagetype]; !ok {
		stage.MessageTypes[messagetype] = struct{}{}
		stage.MessageType_stagedOrder[messagetype] = stage.MessageTypeOrder
		stage.MessageType_orderStaged[stage.MessageTypeOrder] = messagetype
		stage.MessageTypeOrder++
	}
	stage.MessageTypes_mapString[messagetype.Name] = messagetype

	return messagetype
}

// StagePreserveOrder puts messagetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MessageTypeOrder
// - update stage.MessageTypeOrder accordingly
func (messagetype *MessageType) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.MessageTypes[messagetype]; !ok {
		stage.MessageTypes[messagetype] = struct{}{}

		if order > stage.MessageTypeOrder {
			stage.MessageTypeOrder = order
		}
		stage.MessageType_stagedOrder[messagetype] = order
		stage.MessageType_orderStaged[order] = messagetype
		stage.MessageTypeOrder++
	}
	stage.MessageTypes_mapString[messagetype.Name] = messagetype
}

// Unstage removes messagetype off the model stage
func (messagetype *MessageType) Unstage(stage *Stage) *MessageType {
	delete(stage.MessageTypes, messagetype)
	// issue1150
	// delete(stage.MessageType_stagedOrder, messagetype)
	delete(stage.MessageTypes_mapString, messagetype.Name)

	return messagetype
}

// UnstageVoid removes messagetype off the model stage
func (messagetype *MessageType) UnstageVoid(stage *Stage) {
	delete(stage.MessageTypes, messagetype)
	// issue1150
	// delete(stage.MessageType_stagedOrder, messagetype)
	delete(stage.MessageTypes_mapString, messagetype.Name)
}

// commit messagetype to the back repo (if it is already staged)
func (messagetype *MessageType) Commit(stage *Stage) *MessageType {
	if _, ok := stage.MessageTypes[messagetype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMessageType(messagetype)
		}
	}
	return messagetype
}

func (messagetype *MessageType) CommitVoid(stage *Stage) {
	messagetype.Commit(stage)
}

func (messagetype *MessageType) StageVoid(stage *Stage) {
	messagetype.Stage(stage)
}

// Checkout messagetype to the back repo (if it is already staged)
func (messagetype *MessageType) Checkout(stage *Stage) *MessageType {
	if _, ok := stage.MessageTypes[messagetype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMessageType(messagetype)
		}
	}
	return messagetype
}

// for satisfaction of GongStruct interface
func (messagetype *MessageType) GetName() (res string) {
	return messagetype.Name
}

// for satisfaction of GongStruct interface
func (messagetype *MessageType) SetName(name string) {
	messagetype.Name = name
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

// Stage puts notestateshape to the model stage
func (notestateshape *NoteStateShape) Stage(stage *Stage) *NoteStateShape {
	if _, ok := stage.NoteStateShapes[notestateshape]; !ok {
		stage.NoteStateShapes[notestateshape] = struct{}{}
		stage.NoteStateShape_stagedOrder[notestateshape] = stage.NoteStateShapeOrder
		stage.NoteStateShape_orderStaged[stage.NoteStateShapeOrder] = notestateshape
		stage.NoteStateShapeOrder++
	}
	stage.NoteStateShapes_mapString[notestateshape.Name] = notestateshape

	return notestateshape
}

// StagePreserveOrder puts notestateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteStateShapeOrder
// - update stage.NoteStateShapeOrder accordingly
func (notestateshape *NoteStateShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteStateShapes[notestateshape]; !ok {
		stage.NoteStateShapes[notestateshape] = struct{}{}

		if order > stage.NoteStateShapeOrder {
			stage.NoteStateShapeOrder = order
		}
		stage.NoteStateShape_stagedOrder[notestateshape] = order
		stage.NoteStateShape_orderStaged[order] = notestateshape
		stage.NoteStateShapeOrder++
	}
	stage.NoteStateShapes_mapString[notestateshape.Name] = notestateshape
}

// Unstage removes notestateshape off the model stage
func (notestateshape *NoteStateShape) Unstage(stage *Stage) *NoteStateShape {
	delete(stage.NoteStateShapes, notestateshape)
	// issue1150
	// delete(stage.NoteStateShape_stagedOrder, notestateshape)
	delete(stage.NoteStateShapes_mapString, notestateshape.Name)

	return notestateshape
}

// UnstageVoid removes notestateshape off the model stage
func (notestateshape *NoteStateShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteStateShapes, notestateshape)
	// issue1150
	// delete(stage.NoteStateShape_stagedOrder, notestateshape)
	delete(stage.NoteStateShapes_mapString, notestateshape.Name)
}

// commit notestateshape to the back repo (if it is already staged)
func (notestateshape *NoteStateShape) Commit(stage *Stage) *NoteStateShape {
	if _, ok := stage.NoteStateShapes[notestateshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteStateShape(notestateshape)
		}
	}
	return notestateshape
}

func (notestateshape *NoteStateShape) CommitVoid(stage *Stage) {
	notestateshape.Commit(stage)
}

func (notestateshape *NoteStateShape) StageVoid(stage *Stage) {
	notestateshape.Stage(stage)
}

// Checkout notestateshape to the back repo (if it is already staged)
func (notestateshape *NoteStateShape) Checkout(stage *Stage) *NoteStateShape {
	if _, ok := stage.NoteStateShapes[notestateshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteStateShape(notestateshape)
		}
	}
	return notestateshape
}

// for satisfaction of GongStruct interface
func (notestateshape *NoteStateShape) GetName() (res string) {
	return notestateshape.Name
}

// for satisfaction of GongStruct interface
func (notestateshape *NoteStateShape) SetName(name string) {
	notestateshape.Name = name
}

// Stage puts notetransitionshape to the model stage
func (notetransitionshape *NoteTransitionShape) Stage(stage *Stage) *NoteTransitionShape {
	if _, ok := stage.NoteTransitionShapes[notetransitionshape]; !ok {
		stage.NoteTransitionShapes[notetransitionshape] = struct{}{}
		stage.NoteTransitionShape_stagedOrder[notetransitionshape] = stage.NoteTransitionShapeOrder
		stage.NoteTransitionShape_orderStaged[stage.NoteTransitionShapeOrder] = notetransitionshape
		stage.NoteTransitionShapeOrder++
	}
	stage.NoteTransitionShapes_mapString[notetransitionshape.Name] = notetransitionshape

	return notetransitionshape
}

// StagePreserveOrder puts notetransitionshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteTransitionShapeOrder
// - update stage.NoteTransitionShapeOrder accordingly
func (notetransitionshape *NoteTransitionShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.NoteTransitionShapes[notetransitionshape]; !ok {
		stage.NoteTransitionShapes[notetransitionshape] = struct{}{}

		if order > stage.NoteTransitionShapeOrder {
			stage.NoteTransitionShapeOrder = order
		}
		stage.NoteTransitionShape_stagedOrder[notetransitionshape] = order
		stage.NoteTransitionShape_orderStaged[order] = notetransitionshape
		stage.NoteTransitionShapeOrder++
	}
	stage.NoteTransitionShapes_mapString[notetransitionshape.Name] = notetransitionshape
}

// Unstage removes notetransitionshape off the model stage
func (notetransitionshape *NoteTransitionShape) Unstage(stage *Stage) *NoteTransitionShape {
	delete(stage.NoteTransitionShapes, notetransitionshape)
	// issue1150
	// delete(stage.NoteTransitionShape_stagedOrder, notetransitionshape)
	delete(stage.NoteTransitionShapes_mapString, notetransitionshape.Name)

	return notetransitionshape
}

// UnstageVoid removes notetransitionshape off the model stage
func (notetransitionshape *NoteTransitionShape) UnstageVoid(stage *Stage) {
	delete(stage.NoteTransitionShapes, notetransitionshape)
	// issue1150
	// delete(stage.NoteTransitionShape_stagedOrder, notetransitionshape)
	delete(stage.NoteTransitionShapes_mapString, notetransitionshape.Name)
}

// commit notetransitionshape to the back repo (if it is already staged)
func (notetransitionshape *NoteTransitionShape) Commit(stage *Stage) *NoteTransitionShape {
	if _, ok := stage.NoteTransitionShapes[notetransitionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNoteTransitionShape(notetransitionshape)
		}
	}
	return notetransitionshape
}

func (notetransitionshape *NoteTransitionShape) CommitVoid(stage *Stage) {
	notetransitionshape.Commit(stage)
}

func (notetransitionshape *NoteTransitionShape) StageVoid(stage *Stage) {
	notetransitionshape.Stage(stage)
}

// Checkout notetransitionshape to the back repo (if it is already staged)
func (notetransitionshape *NoteTransitionShape) Checkout(stage *Stage) *NoteTransitionShape {
	if _, ok := stage.NoteTransitionShapes[notetransitionshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNoteTransitionShape(notetransitionshape)
		}
	}
	return notetransitionshape
}

// for satisfaction of GongStruct interface
func (notetransitionshape *NoteTransitionShape) GetName() (res string) {
	return notetransitionshape.Name
}

// for satisfaction of GongStruct interface
func (notetransitionshape *NoteTransitionShape) SetName(name string) {
	notetransitionshape.Name = name
}

// Stage puts object to the model stage
func (object *Object) Stage(stage *Stage) *Object {
	if _, ok := stage.Objects[object]; !ok {
		stage.Objects[object] = struct{}{}
		stage.Object_stagedOrder[object] = stage.ObjectOrder
		stage.Object_orderStaged[stage.ObjectOrder] = object
		stage.ObjectOrder++
	}
	stage.Objects_mapString[object.Name] = object

	return object
}

// StagePreserveOrder puts object to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ObjectOrder
// - update stage.ObjectOrder accordingly
func (object *Object) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Objects[object]; !ok {
		stage.Objects[object] = struct{}{}

		if order > stage.ObjectOrder {
			stage.ObjectOrder = order
		}
		stage.Object_stagedOrder[object] = order
		stage.Object_orderStaged[order] = object
		stage.ObjectOrder++
	}
	stage.Objects_mapString[object.Name] = object
}

// Unstage removes object off the model stage
func (object *Object) Unstage(stage *Stage) *Object {
	delete(stage.Objects, object)
	// issue1150
	// delete(stage.Object_stagedOrder, object)
	delete(stage.Objects_mapString, object.Name)

	return object
}

// UnstageVoid removes object off the model stage
func (object *Object) UnstageVoid(stage *Stage) {
	delete(stage.Objects, object)
	// issue1150
	// delete(stage.Object_stagedOrder, object)
	delete(stage.Objects_mapString, object.Name)
}

// commit object to the back repo (if it is already staged)
func (object *Object) Commit(stage *Stage) *Object {
	if _, ok := stage.Objects[object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitObject(object)
		}
	}
	return object
}

func (object *Object) CommitVoid(stage *Stage) {
	object.Commit(stage)
}

func (object *Object) StageVoid(stage *Stage) {
	object.Stage(stage)
}

// Checkout object to the back repo (if it is already staged)
func (object *Object) Checkout(stage *Stage) *Object {
	if _, ok := stage.Objects[object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutObject(object)
		}
	}
	return object
}

// for satisfaction of GongStruct interface
func (object *Object) GetName() (res string) {
	return object.Name
}

// for satisfaction of GongStruct interface
func (object *Object) SetName(name string) {
	object.Name = name
}

// Stage puts role to the model stage
func (role *Role) Stage(stage *Stage) *Role {
	if _, ok := stage.Roles[role]; !ok {
		stage.Roles[role] = struct{}{}
		stage.Role_stagedOrder[role] = stage.RoleOrder
		stage.Role_orderStaged[stage.RoleOrder] = role
		stage.RoleOrder++
	}
	stage.Roles_mapString[role.Name] = role

	return role
}

// StagePreserveOrder puts role to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RoleOrder
// - update stage.RoleOrder accordingly
func (role *Role) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Roles[role]; !ok {
		stage.Roles[role] = struct{}{}

		if order > stage.RoleOrder {
			stage.RoleOrder = order
		}
		stage.Role_stagedOrder[role] = order
		stage.Role_orderStaged[order] = role
		stage.RoleOrder++
	}
	stage.Roles_mapString[role.Name] = role
}

// Unstage removes role off the model stage
func (role *Role) Unstage(stage *Stage) *Role {
	delete(stage.Roles, role)
	// issue1150
	// delete(stage.Role_stagedOrder, role)
	delete(stage.Roles_mapString, role.Name)

	return role
}

// UnstageVoid removes role off the model stage
func (role *Role) UnstageVoid(stage *Stage) {
	delete(stage.Roles, role)
	// issue1150
	// delete(stage.Role_stagedOrder, role)
	delete(stage.Roles_mapString, role.Name)
}

// commit role to the back repo (if it is already staged)
func (role *Role) Commit(stage *Stage) *Role {
	if _, ok := stage.Roles[role]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRole(role)
		}
	}
	return role
}

func (role *Role) CommitVoid(stage *Stage) {
	role.Commit(stage)
}

func (role *Role) StageVoid(stage *Stage) {
	role.Stage(stage)
}

// Checkout role to the back repo (if it is already staged)
func (role *Role) Checkout(stage *Stage) *Role {
	if _, ok := stage.Roles[role]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRole(role)
		}
	}
	return role
}

// for satisfaction of GongStruct interface
func (role *Role) GetName() (res string) {
	return role.Name
}

// for satisfaction of GongStruct interface
func (role *Role) SetName(name string) {
	role.Name = name
}

// Stage puts state to the model stage
func (state *State) Stage(stage *Stage) *State {
	if _, ok := stage.States[state]; !ok {
		stage.States[state] = struct{}{}
		stage.State_stagedOrder[state] = stage.StateOrder
		stage.State_orderStaged[stage.StateOrder] = state
		stage.StateOrder++
	}
	stage.States_mapString[state.Name] = state

	return state
}

// StagePreserveOrder puts state to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateOrder
// - update stage.StateOrder accordingly
func (state *State) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.States[state]; !ok {
		stage.States[state] = struct{}{}

		if order > stage.StateOrder {
			stage.StateOrder = order
		}
		stage.State_stagedOrder[state] = order
		stage.State_orderStaged[order] = state
		stage.StateOrder++
	}
	stage.States_mapString[state.Name] = state
}

// Unstage removes state off the model stage
func (state *State) Unstage(stage *Stage) *State {
	delete(stage.States, state)
	// issue1150
	// delete(stage.State_stagedOrder, state)
	delete(stage.States_mapString, state.Name)

	return state
}

// UnstageVoid removes state off the model stage
func (state *State) UnstageVoid(stage *Stage) {
	delete(stage.States, state)
	// issue1150
	// delete(stage.State_stagedOrder, state)
	delete(stage.States_mapString, state.Name)
}

// commit state to the back repo (if it is already staged)
func (state *State) Commit(stage *Stage) *State {
	if _, ok := stage.States[state]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitState(state)
		}
	}
	return state
}

func (state *State) CommitVoid(stage *Stage) {
	state.Commit(stage)
}

func (state *State) StageVoid(stage *Stage) {
	state.Stage(stage)
}

// Checkout state to the back repo (if it is already staged)
func (state *State) Checkout(stage *Stage) *State {
	if _, ok := stage.States[state]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutState(state)
		}
	}
	return state
}

// for satisfaction of GongStruct interface
func (state *State) GetName() (res string) {
	return state.Name
}

// for satisfaction of GongStruct interface
func (state *State) SetName(name string) {
	state.Name = name
}

// Stage puts statemachine to the model stage
func (statemachine *StateMachine) Stage(stage *Stage) *StateMachine {
	if _, ok := stage.StateMachines[statemachine]; !ok {
		stage.StateMachines[statemachine] = struct{}{}
		stage.StateMachine_stagedOrder[statemachine] = stage.StateMachineOrder
		stage.StateMachine_orderStaged[stage.StateMachineOrder] = statemachine
		stage.StateMachineOrder++
	}
	stage.StateMachines_mapString[statemachine.Name] = statemachine

	return statemachine
}

// StagePreserveOrder puts statemachine to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateMachineOrder
// - update stage.StateMachineOrder accordingly
func (statemachine *StateMachine) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StateMachines[statemachine]; !ok {
		stage.StateMachines[statemachine] = struct{}{}

		if order > stage.StateMachineOrder {
			stage.StateMachineOrder = order
		}
		stage.StateMachine_stagedOrder[statemachine] = order
		stage.StateMachine_orderStaged[order] = statemachine
		stage.StateMachineOrder++
	}
	stage.StateMachines_mapString[statemachine.Name] = statemachine
}

// Unstage removes statemachine off the model stage
func (statemachine *StateMachine) Unstage(stage *Stage) *StateMachine {
	delete(stage.StateMachines, statemachine)
	// issue1150
	// delete(stage.StateMachine_stagedOrder, statemachine)
	delete(stage.StateMachines_mapString, statemachine.Name)

	return statemachine
}

// UnstageVoid removes statemachine off the model stage
func (statemachine *StateMachine) UnstageVoid(stage *Stage) {
	delete(stage.StateMachines, statemachine)
	// issue1150
	// delete(stage.StateMachine_stagedOrder, statemachine)
	delete(stage.StateMachines_mapString, statemachine.Name)
}

// commit statemachine to the back repo (if it is already staged)
func (statemachine *StateMachine) Commit(stage *Stage) *StateMachine {
	if _, ok := stage.StateMachines[statemachine]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStateMachine(statemachine)
		}
	}
	return statemachine
}

func (statemachine *StateMachine) CommitVoid(stage *Stage) {
	statemachine.Commit(stage)
}

func (statemachine *StateMachine) StageVoid(stage *Stage) {
	statemachine.Stage(stage)
}

// Checkout statemachine to the back repo (if it is already staged)
func (statemachine *StateMachine) Checkout(stage *Stage) *StateMachine {
	if _, ok := stage.StateMachines[statemachine]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStateMachine(statemachine)
		}
	}
	return statemachine
}

// for satisfaction of GongStruct interface
func (statemachine *StateMachine) GetName() (res string) {
	return statemachine.Name
}

// for satisfaction of GongStruct interface
func (statemachine *StateMachine) SetName(name string) {
	statemachine.Name = name
}

// Stage puts stateshape to the model stage
func (stateshape *StateShape) Stage(stage *Stage) *StateShape {
	if _, ok := stage.StateShapes[stateshape]; !ok {
		stage.StateShapes[stateshape] = struct{}{}
		stage.StateShape_stagedOrder[stateshape] = stage.StateShapeOrder
		stage.StateShape_orderStaged[stage.StateShapeOrder] = stateshape
		stage.StateShapeOrder++
	}
	stage.StateShapes_mapString[stateshape.Name] = stateshape

	return stateshape
}

// StagePreserveOrder puts stateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateShapeOrder
// - update stage.StateShapeOrder accordingly
func (stateshape *StateShape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.StateShapes[stateshape]; !ok {
		stage.StateShapes[stateshape] = struct{}{}

		if order > stage.StateShapeOrder {
			stage.StateShapeOrder = order
		}
		stage.StateShape_stagedOrder[stateshape] = order
		stage.StateShape_orderStaged[order] = stateshape
		stage.StateShapeOrder++
	}
	stage.StateShapes_mapString[stateshape.Name] = stateshape
}

// Unstage removes stateshape off the model stage
func (stateshape *StateShape) Unstage(stage *Stage) *StateShape {
	delete(stage.StateShapes, stateshape)
	// issue1150
	// delete(stage.StateShape_stagedOrder, stateshape)
	delete(stage.StateShapes_mapString, stateshape.Name)

	return stateshape
}

// UnstageVoid removes stateshape off the model stage
func (stateshape *StateShape) UnstageVoid(stage *Stage) {
	delete(stage.StateShapes, stateshape)
	// issue1150
	// delete(stage.StateShape_stagedOrder, stateshape)
	delete(stage.StateShapes_mapString, stateshape.Name)
}

// commit stateshape to the back repo (if it is already staged)
func (stateshape *StateShape) Commit(stage *Stage) *StateShape {
	if _, ok := stage.StateShapes[stateshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStateShape(stateshape)
		}
	}
	return stateshape
}

func (stateshape *StateShape) CommitVoid(stage *Stage) {
	stateshape.Commit(stage)
}

func (stateshape *StateShape) StageVoid(stage *Stage) {
	stateshape.Stage(stage)
}

// Checkout stateshape to the back repo (if it is already staged)
func (stateshape *StateShape) Checkout(stage *Stage) *StateShape {
	if _, ok := stage.StateShapes[stateshape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStateShape(stateshape)
		}
	}
	return stateshape
}

// for satisfaction of GongStruct interface
func (stateshape *StateShape) GetName() (res string) {
	return stateshape.Name
}

// for satisfaction of GongStruct interface
func (stateshape *StateShape) SetName(name string) {
	stateshape.Name = name
}

// Stage puts transition to the model stage
func (transition *Transition) Stage(stage *Stage) *Transition {
	if _, ok := stage.Transitions[transition]; !ok {
		stage.Transitions[transition] = struct{}{}
		stage.Transition_stagedOrder[transition] = stage.TransitionOrder
		stage.Transition_orderStaged[stage.TransitionOrder] = transition
		stage.TransitionOrder++
	}
	stage.Transitions_mapString[transition.Name] = transition

	return transition
}

// StagePreserveOrder puts transition to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TransitionOrder
// - update stage.TransitionOrder accordingly
func (transition *Transition) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Transitions[transition]; !ok {
		stage.Transitions[transition] = struct{}{}

		if order > stage.TransitionOrder {
			stage.TransitionOrder = order
		}
		stage.Transition_stagedOrder[transition] = order
		stage.Transition_orderStaged[order] = transition
		stage.TransitionOrder++
	}
	stage.Transitions_mapString[transition.Name] = transition
}

// Unstage removes transition off the model stage
func (transition *Transition) Unstage(stage *Stage) *Transition {
	delete(stage.Transitions, transition)
	// issue1150
	// delete(stage.Transition_stagedOrder, transition)
	delete(stage.Transitions_mapString, transition.Name)

	return transition
}

// UnstageVoid removes transition off the model stage
func (transition *Transition) UnstageVoid(stage *Stage) {
	delete(stage.Transitions, transition)
	// issue1150
	// delete(stage.Transition_stagedOrder, transition)
	delete(stage.Transitions_mapString, transition.Name)
}

// commit transition to the back repo (if it is already staged)
func (transition *Transition) Commit(stage *Stage) *Transition {
	if _, ok := stage.Transitions[transition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTransition(transition)
		}
	}
	return transition
}

func (transition *Transition) CommitVoid(stage *Stage) {
	transition.Commit(stage)
}

func (transition *Transition) StageVoid(stage *Stage) {
	transition.Stage(stage)
}

// Checkout transition to the back repo (if it is already staged)
func (transition *Transition) Checkout(stage *Stage) *Transition {
	if _, ok := stage.Transitions[transition]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTransition(transition)
		}
	}
	return transition
}

// for satisfaction of GongStruct interface
func (transition *Transition) GetName() (res string) {
	return transition.Name
}

// for satisfaction of GongStruct interface
func (transition *Transition) SetName(name string) {
	transition.Name = name
}

// Stage puts transition_shape to the model stage
func (transition_shape *Transition_Shape) Stage(stage *Stage) *Transition_Shape {
	if _, ok := stage.Transition_Shapes[transition_shape]; !ok {
		stage.Transition_Shapes[transition_shape] = struct{}{}
		stage.Transition_Shape_stagedOrder[transition_shape] = stage.Transition_ShapeOrder
		stage.Transition_Shape_orderStaged[stage.Transition_ShapeOrder] = transition_shape
		stage.Transition_ShapeOrder++
	}
	stage.Transition_Shapes_mapString[transition_shape.Name] = transition_shape

	return transition_shape
}

// StagePreserveOrder puts transition_shape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Transition_ShapeOrder
// - update stage.Transition_ShapeOrder accordingly
func (transition_shape *Transition_Shape) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Transition_Shapes[transition_shape]; !ok {
		stage.Transition_Shapes[transition_shape] = struct{}{}

		if order > stage.Transition_ShapeOrder {
			stage.Transition_ShapeOrder = order
		}
		stage.Transition_Shape_stagedOrder[transition_shape] = order
		stage.Transition_Shape_orderStaged[order] = transition_shape
		stage.Transition_ShapeOrder++
	}
	stage.Transition_Shapes_mapString[transition_shape.Name] = transition_shape
}

// Unstage removes transition_shape off the model stage
func (transition_shape *Transition_Shape) Unstage(stage *Stage) *Transition_Shape {
	delete(stage.Transition_Shapes, transition_shape)
	// issue1150
	// delete(stage.Transition_Shape_stagedOrder, transition_shape)
	delete(stage.Transition_Shapes_mapString, transition_shape.Name)

	return transition_shape
}

// UnstageVoid removes transition_shape off the model stage
func (transition_shape *Transition_Shape) UnstageVoid(stage *Stage) {
	delete(stage.Transition_Shapes, transition_shape)
	// issue1150
	// delete(stage.Transition_Shape_stagedOrder, transition_shape)
	delete(stage.Transition_Shapes_mapString, transition_shape.Name)
}

// commit transition_shape to the back repo (if it is already staged)
func (transition_shape *Transition_Shape) Commit(stage *Stage) *Transition_Shape {
	if _, ok := stage.Transition_Shapes[transition_shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTransition_Shape(transition_shape)
		}
	}
	return transition_shape
}

func (transition_shape *Transition_Shape) CommitVoid(stage *Stage) {
	transition_shape.Commit(stage)
}

func (transition_shape *Transition_Shape) StageVoid(stage *Stage) {
	transition_shape.Stage(stage)
}

// Checkout transition_shape to the back repo (if it is already staged)
func (transition_shape *Transition_Shape) Checkout(stage *Stage) *Transition_Shape {
	if _, ok := stage.Transition_Shapes[transition_shape]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTransition_Shape(transition_shape)
		}
	}
	return transition_shape
}

// for satisfaction of GongStruct interface
func (transition_shape *Transition_Shape) GetName() (res string) {
	return transition_shape.Name
}

// for satisfaction of GongStruct interface
func (transition_shape *Transition_Shape) SetName(name string) {
	transition_shape.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAction(Action *Action)
	CreateORMActivities(Activities *Activities)
	CreateORMArchitecture(Architecture *Architecture)
	CreateORMDiagram(Diagram *Diagram)
	CreateORMGuard(Guard *Guard)
	CreateORMKill(Kill *Kill)
	CreateORMLibrary(Library *Library)
	CreateORMMessage(Message *Message)
	CreateORMMessageType(MessageType *MessageType)
	CreateORMNote(Note *Note)
	CreateORMNoteShape(NoteShape *NoteShape)
	CreateORMNoteStateShape(NoteStateShape *NoteStateShape)
	CreateORMNoteTransitionShape(NoteTransitionShape *NoteTransitionShape)
	CreateORMObject(Object *Object)
	CreateORMRole(Role *Role)
	CreateORMState(State *State)
	CreateORMStateMachine(StateMachine *StateMachine)
	CreateORMStateShape(StateShape *StateShape)
	CreateORMTransition(Transition *Transition)
	CreateORMTransition_Shape(Transition_Shape *Transition_Shape)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAction(Action *Action)
	DeleteORMActivities(Activities *Activities)
	DeleteORMArchitecture(Architecture *Architecture)
	DeleteORMDiagram(Diagram *Diagram)
	DeleteORMGuard(Guard *Guard)
	DeleteORMKill(Kill *Kill)
	DeleteORMLibrary(Library *Library)
	DeleteORMMessage(Message *Message)
	DeleteORMMessageType(MessageType *MessageType)
	DeleteORMNote(Note *Note)
	DeleteORMNoteShape(NoteShape *NoteShape)
	DeleteORMNoteStateShape(NoteStateShape *NoteStateShape)
	DeleteORMNoteTransitionShape(NoteTransitionShape *NoteTransitionShape)
	DeleteORMObject(Object *Object)
	DeleteORMRole(Role *Role)
	DeleteORMState(State *State)
	DeleteORMStateMachine(StateMachine *StateMachine)
	DeleteORMStateShape(StateShape *StateShape)
	DeleteORMTransition(Transition *Transition)
	DeleteORMTransition_Shape(Transition_Shape *Transition_Shape)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Actions = make(map[*Action]struct{})
	stage.Actions_mapString = make(map[string]*Action)
	stage.Action_stagedOrder = make(map[*Action]uint)
	stage.ActionOrder = 0

	stage.Activitiess = make(map[*Activities]struct{})
	stage.Activitiess_mapString = make(map[string]*Activities)
	stage.Activities_stagedOrder = make(map[*Activities]uint)
	stage.ActivitiesOrder = 0

	stage.Architectures = make(map[*Architecture]struct{})
	stage.Architectures_mapString = make(map[string]*Architecture)
	stage.Architecture_stagedOrder = make(map[*Architecture]uint)
	stage.ArchitectureOrder = 0

	stage.Diagrams = make(map[*Diagram]struct{})
	stage.Diagrams_mapString = make(map[string]*Diagram)
	stage.Diagram_stagedOrder = make(map[*Diagram]uint)
	stage.DiagramOrder = 0

	stage.Guards = make(map[*Guard]struct{})
	stage.Guards_mapString = make(map[string]*Guard)
	stage.Guard_stagedOrder = make(map[*Guard]uint)
	stage.GuardOrder = 0

	stage.Kills = make(map[*Kill]struct{})
	stage.Kills_mapString = make(map[string]*Kill)
	stage.Kill_stagedOrder = make(map[*Kill]uint)
	stage.KillOrder = 0

	stage.Librarys = make(map[*Library]struct{})
	stage.Librarys_mapString = make(map[string]*Library)
	stage.Library_stagedOrder = make(map[*Library]uint)
	stage.LibraryOrder = 0

	stage.Messages = make(map[*Message]struct{})
	stage.Messages_mapString = make(map[string]*Message)
	stage.Message_stagedOrder = make(map[*Message]uint)
	stage.MessageOrder = 0

	stage.MessageTypes = make(map[*MessageType]struct{})
	stage.MessageTypes_mapString = make(map[string]*MessageType)
	stage.MessageType_stagedOrder = make(map[*MessageType]uint)
	stage.MessageTypeOrder = 0

	stage.Notes = make(map[*Note]struct{})
	stage.Notes_mapString = make(map[string]*Note)
	stage.Note_stagedOrder = make(map[*Note]uint)
	stage.NoteOrder = 0

	stage.NoteShapes = make(map[*NoteShape]struct{})
	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	stage.NoteShape_stagedOrder = make(map[*NoteShape]uint)
	stage.NoteShapeOrder = 0

	stage.NoteStateShapes = make(map[*NoteStateShape]struct{})
	stage.NoteStateShapes_mapString = make(map[string]*NoteStateShape)
	stage.NoteStateShape_stagedOrder = make(map[*NoteStateShape]uint)
	stage.NoteStateShapeOrder = 0

	stage.NoteTransitionShapes = make(map[*NoteTransitionShape]struct{})
	stage.NoteTransitionShapes_mapString = make(map[string]*NoteTransitionShape)
	stage.NoteTransitionShape_stagedOrder = make(map[*NoteTransitionShape]uint)
	stage.NoteTransitionShapeOrder = 0

	stage.Objects = make(map[*Object]struct{})
	stage.Objects_mapString = make(map[string]*Object)
	stage.Object_stagedOrder = make(map[*Object]uint)
	stage.ObjectOrder = 0

	stage.Roles = make(map[*Role]struct{})
	stage.Roles_mapString = make(map[string]*Role)
	stage.Role_stagedOrder = make(map[*Role]uint)
	stage.RoleOrder = 0

	stage.States = make(map[*State]struct{})
	stage.States_mapString = make(map[string]*State)
	stage.State_stagedOrder = make(map[*State]uint)
	stage.StateOrder = 0

	stage.StateMachines = make(map[*StateMachine]struct{})
	stage.StateMachines_mapString = make(map[string]*StateMachine)
	stage.StateMachine_stagedOrder = make(map[*StateMachine]uint)
	stage.StateMachineOrder = 0

	stage.StateShapes = make(map[*StateShape]struct{})
	stage.StateShapes_mapString = make(map[string]*StateShape)
	stage.StateShape_stagedOrder = make(map[*StateShape]uint)
	stage.StateShapeOrder = 0

	stage.Transitions = make(map[*Transition]struct{})
	stage.Transitions_mapString = make(map[string]*Transition)
	stage.Transition_stagedOrder = make(map[*Transition]uint)
	stage.TransitionOrder = 0

	stage.Transition_Shapes = make(map[*Transition_Shape]struct{})
	stage.Transition_Shapes_mapString = make(map[string]*Transition_Shape)
	stage.Transition_Shape_stagedOrder = make(map[*Transition_Shape]uint)
	stage.Transition_ShapeOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Actions = nil
	stage.Actions_mapString = nil

	stage.Activitiess = nil
	stage.Activitiess_mapString = nil

	stage.Architectures = nil
	stage.Architectures_mapString = nil

	stage.Diagrams = nil
	stage.Diagrams_mapString = nil

	stage.Guards = nil
	stage.Guards_mapString = nil

	stage.Kills = nil
	stage.Kills_mapString = nil

	stage.Librarys = nil
	stage.Librarys_mapString = nil

	stage.Messages = nil
	stage.Messages_mapString = nil

	stage.MessageTypes = nil
	stage.MessageTypes_mapString = nil

	stage.Notes = nil
	stage.Notes_mapString = nil

	stage.NoteShapes = nil
	stage.NoteShapes_mapString = nil

	stage.NoteStateShapes = nil
	stage.NoteStateShapes_mapString = nil

	stage.NoteTransitionShapes = nil
	stage.NoteTransitionShapes_mapString = nil

	stage.Objects = nil
	stage.Objects_mapString = nil

	stage.Roles = nil
	stage.Roles_mapString = nil

	stage.States = nil
	stage.States_mapString = nil

	stage.StateMachines = nil
	stage.StateMachines_mapString = nil

	stage.StateShapes = nil
	stage.StateShapes_mapString = nil

	stage.Transitions = nil
	stage.Transitions_mapString = nil

	stage.Transition_Shapes = nil
	stage.Transition_Shapes_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for action := range stage.Actions {
		action.Unstage(stage)
	}

	for activities := range stage.Activitiess {
		activities.Unstage(stage)
	}

	for architecture := range stage.Architectures {
		architecture.Unstage(stage)
	}

	for diagram := range stage.Diagrams {
		diagram.Unstage(stage)
	}

	for guard := range stage.Guards {
		guard.Unstage(stage)
	}

	for kill := range stage.Kills {
		kill.Unstage(stage)
	}

	for library := range stage.Librarys {
		library.Unstage(stage)
	}

	for message := range stage.Messages {
		message.Unstage(stage)
	}

	for messagetype := range stage.MessageTypes {
		messagetype.Unstage(stage)
	}

	for note := range stage.Notes {
		note.Unstage(stage)
	}

	for noteshape := range stage.NoteShapes {
		noteshape.Unstage(stage)
	}

	for notestateshape := range stage.NoteStateShapes {
		notestateshape.Unstage(stage)
	}

	for notetransitionshape := range stage.NoteTransitionShapes {
		notetransitionshape.Unstage(stage)
	}

	for object := range stage.Objects {
		object.Unstage(stage)
	}

	for role := range stage.Roles {
		role.Unstage(stage)
	}

	for state := range stage.States {
		state.Unstage(stage)
	}

	for statemachine := range stage.StateMachines {
		statemachine.Unstage(stage)
	}

	for stateshape := range stage.StateShapes {
		stateshape.Unstage(stage)
	}

	for transition := range stage.Transitions {
		transition.Unstage(stage)
	}

	for transition_shape := range stage.Transition_Shapes {
		transition_shape.Unstage(stage)
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
	case map[*Action]any:
		return any(&stage.Actions).(*Type)
	case map[*Activities]any:
		return any(&stage.Activitiess).(*Type)
	case map[*Architecture]any:
		return any(&stage.Architectures).(*Type)
	case map[*Diagram]any:
		return any(&stage.Diagrams).(*Type)
	case map[*Guard]any:
		return any(&stage.Guards).(*Type)
	case map[*Kill]any:
		return any(&stage.Kills).(*Type)
	case map[*Library]any:
		return any(&stage.Librarys).(*Type)
	case map[*Message]any:
		return any(&stage.Messages).(*Type)
	case map[*MessageType]any:
		return any(&stage.MessageTypes).(*Type)
	case map[*Note]any:
		return any(&stage.Notes).(*Type)
	case map[*NoteShape]any:
		return any(&stage.NoteShapes).(*Type)
	case map[*NoteStateShape]any:
		return any(&stage.NoteStateShapes).(*Type)
	case map[*NoteTransitionShape]any:
		return any(&stage.NoteTransitionShapes).(*Type)
	case map[*Object]any:
		return any(&stage.Objects).(*Type)
	case map[*Role]any:
		return any(&stage.Roles).(*Type)
	case map[*State]any:
		return any(&stage.States).(*Type)
	case map[*StateMachine]any:
		return any(&stage.StateMachines).(*Type)
	case map[*StateShape]any:
		return any(&stage.StateShapes).(*Type)
	case map[*Transition]any:
		return any(&stage.Transitions).(*Type)
	case map[*Transition_Shape]any:
		return any(&stage.Transition_Shapes).(*Type)
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
	case *Action:
		return any(stage.Actions_mapString).(map[string]Type)
	case *Activities:
		return any(stage.Activitiess_mapString).(map[string]Type)
	case *Architecture:
		return any(stage.Architectures_mapString).(map[string]Type)
	case *Diagram:
		return any(stage.Diagrams_mapString).(map[string]Type)
	case *Guard:
		return any(stage.Guards_mapString).(map[string]Type)
	case *Kill:
		return any(stage.Kills_mapString).(map[string]Type)
	case *Library:
		return any(stage.Librarys_mapString).(map[string]Type)
	case *Message:
		return any(stage.Messages_mapString).(map[string]Type)
	case *MessageType:
		return any(stage.MessageTypes_mapString).(map[string]Type)
	case *Note:
		return any(stage.Notes_mapString).(map[string]Type)
	case *NoteShape:
		return any(stage.NoteShapes_mapString).(map[string]Type)
	case *NoteStateShape:
		return any(stage.NoteStateShapes_mapString).(map[string]Type)
	case *NoteTransitionShape:
		return any(stage.NoteTransitionShapes_mapString).(map[string]Type)
	case *Object:
		return any(stage.Objects_mapString).(map[string]Type)
	case *Role:
		return any(stage.Roles_mapString).(map[string]Type)
	case *State:
		return any(stage.States_mapString).(map[string]Type)
	case *StateMachine:
		return any(stage.StateMachines_mapString).(map[string]Type)
	case *StateShape:
		return any(stage.StateShapes_mapString).(map[string]Type)
	case *Transition:
		return any(stage.Transitions_mapString).(map[string]Type)
	case *Transition_Shape:
		return any(stage.Transition_Shapes_mapString).(map[string]Type)
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
	case Action:
		return any(&stage.Actions).(*map[*Type]struct{})
	case Activities:
		return any(&stage.Activitiess).(*map[*Type]struct{})
	case Architecture:
		return any(&stage.Architectures).(*map[*Type]struct{})
	case Diagram:
		return any(&stage.Diagrams).(*map[*Type]struct{})
	case Guard:
		return any(&stage.Guards).(*map[*Type]struct{})
	case Kill:
		return any(&stage.Kills).(*map[*Type]struct{})
	case Library:
		return any(&stage.Librarys).(*map[*Type]struct{})
	case Message:
		return any(&stage.Messages).(*map[*Type]struct{})
	case MessageType:
		return any(&stage.MessageTypes).(*map[*Type]struct{})
	case Note:
		return any(&stage.Notes).(*map[*Type]struct{})
	case NoteShape:
		return any(&stage.NoteShapes).(*map[*Type]struct{})
	case NoteStateShape:
		return any(&stage.NoteStateShapes).(*map[*Type]struct{})
	case NoteTransitionShape:
		return any(&stage.NoteTransitionShapes).(*map[*Type]struct{})
	case Object:
		return any(&stage.Objects).(*map[*Type]struct{})
	case Role:
		return any(&stage.Roles).(*map[*Type]struct{})
	case State:
		return any(&stage.States).(*map[*Type]struct{})
	case StateMachine:
		return any(&stage.StateMachines).(*map[*Type]struct{})
	case StateShape:
		return any(&stage.StateShapes).(*map[*Type]struct{})
	case Transition:
		return any(&stage.Transitions).(*map[*Type]struct{})
	case Transition_Shape:
		return any(&stage.Transition_Shapes).(*map[*Type]struct{})
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
	case *Action:
		return any(&stage.Actions).(*map[Type]struct{})
	case *Activities:
		return any(&stage.Activitiess).(*map[Type]struct{})
	case *Architecture:
		return any(&stage.Architectures).(*map[Type]struct{})
	case *Diagram:
		return any(&stage.Diagrams).(*map[Type]struct{})
	case *Guard:
		return any(&stage.Guards).(*map[Type]struct{})
	case *Kill:
		return any(&stage.Kills).(*map[Type]struct{})
	case *Library:
		return any(&stage.Librarys).(*map[Type]struct{})
	case *Message:
		return any(&stage.Messages).(*map[Type]struct{})
	case *MessageType:
		return any(&stage.MessageTypes).(*map[Type]struct{})
	case *Note:
		return any(&stage.Notes).(*map[Type]struct{})
	case *NoteShape:
		return any(&stage.NoteShapes).(*map[Type]struct{})
	case *NoteStateShape:
		return any(&stage.NoteStateShapes).(*map[Type]struct{})
	case *NoteTransitionShape:
		return any(&stage.NoteTransitionShapes).(*map[Type]struct{})
	case *Object:
		return any(&stage.Objects).(*map[Type]struct{})
	case *Role:
		return any(&stage.Roles).(*map[Type]struct{})
	case *State:
		return any(&stage.States).(*map[Type]struct{})
	case *StateMachine:
		return any(&stage.StateMachines).(*map[Type]struct{})
	case *StateShape:
		return any(&stage.StateShapes).(*map[Type]struct{})
	case *Transition:
		return any(&stage.Transitions).(*map[Type]struct{})
	case *Transition_Shape:
		return any(&stage.Transition_Shapes).(*map[Type]struct{})
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
	case Action:
		return any(&stage.Actions_mapString).(*map[string]*Type)
	case Activities:
		return any(&stage.Activitiess_mapString).(*map[string]*Type)
	case Architecture:
		return any(&stage.Architectures_mapString).(*map[string]*Type)
	case Diagram:
		return any(&stage.Diagrams_mapString).(*map[string]*Type)
	case Guard:
		return any(&stage.Guards_mapString).(*map[string]*Type)
	case Kill:
		return any(&stage.Kills_mapString).(*map[string]*Type)
	case Library:
		return any(&stage.Librarys_mapString).(*map[string]*Type)
	case Message:
		return any(&stage.Messages_mapString).(*map[string]*Type)
	case MessageType:
		return any(&stage.MessageTypes_mapString).(*map[string]*Type)
	case Note:
		return any(&stage.Notes_mapString).(*map[string]*Type)
	case NoteShape:
		return any(&stage.NoteShapes_mapString).(*map[string]*Type)
	case NoteStateShape:
		return any(&stage.NoteStateShapes_mapString).(*map[string]*Type)
	case NoteTransitionShape:
		return any(&stage.NoteTransitionShapes_mapString).(*map[string]*Type)
	case Object:
		return any(&stage.Objects_mapString).(*map[string]*Type)
	case Role:
		return any(&stage.Roles_mapString).(*map[string]*Type)
	case State:
		return any(&stage.States_mapString).(*map[string]*Type)
	case StateMachine:
		return any(&stage.StateMachines_mapString).(*map[string]*Type)
	case StateShape:
		return any(&stage.StateShapes_mapString).(*map[string]*Type)
	case Transition:
		return any(&stage.Transitions_mapString).(*map[string]*Type)
	case Transition_Shape:
		return any(&stage.Transition_Shapes_mapString).(*map[string]*Type)
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
	case Action:
		return any(&Action{
			// Initialisation of associations
		}).(*Type)
	case Activities:
		return any(&Activities{
			// Initialisation of associations
		}).(*Type)
	case Architecture:
		return any(&Architecture{
			// Initialisation of associations
			// field is initialized with an instance of StateMachine with the name of the field
			StateMachines: []*StateMachine{{Name: "StateMachines"}},
			// field is initialized with an instance of Role with the name of the field
			Roles: []*Role{{Name: "Roles"}},
		}).(*Type)
	case Diagram:
		return any(&Diagram{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of StateShape with the name of the field
			State_Shapes: []*StateShape{{Name: "State_Shapes"}},
			// field is initialized with an instance of State with the name of the field
			StatesWhoseNodeIsExpanded: []*State{{Name: "StatesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Transition_Shape with the name of the field
			Transition_Shapes: []*Transition_Shape{{Name: "Transition_Shapes"}},
			// field is initialized with an instance of NoteShape with the name of the field
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			// field is initialized with an instance of NoteStateShape with the name of the field
			NoteState_Shapes: []*NoteStateShape{{Name: "NoteState_Shapes"}},
			// field is initialized with an instance of NoteTransitionShape with the name of the field
			NoteTransition_Shapes: []*NoteTransitionShape{{Name: "NoteTransition_Shapes"}},
		}).(*Type)
	case Guard:
		return any(&Guard{
			// Initialisation of associations
		}).(*Type)
	case Kill:
		return any(&Kill{
			// Initialisation of associations
		}).(*Type)
	case Library:
		return any(&Library{
			// Initialisation of associations
			// field is initialized with an instance of Library with the name of the field
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			// field is initialized with an instance of Diagram with the name of the field
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			// field is initialized with an instance of StateMachine with the name of the field
			RootStateMachines: []*StateMachine{{Name: "RootStateMachines"}},
			// field is initialized with an instance of StateMachine with the name of the field
			StateMachinesWhoseNodeIsExpanded: []*StateMachine{{Name: "StateMachinesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Note with the name of the field
			RootNotes: []*Note{{Name: "RootNotes"}},
			// field is initialized with an instance of Note with the name of the field
			NotesWhoseNodeIsExpanded: []*Note{{Name: "NotesWhoseNodeIsExpanded"}},
			// field is initialized with an instance of Library with the name of the field
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
		}).(*Type)
	case Message:
		return any(&Message{
			// Initialisation of associations
			// field is initialized with an instance of MessageType with the name of the field
			MessageType: &MessageType{Name: "MessageType"},
			// field is initialized with an instance of Transition with the name of the field
			OriginTransition: &Transition{Name: "OriginTransition"},
		}).(*Type)
	case MessageType:
		return any(&MessageType{
			// Initialisation of associations
		}).(*Type)
	case Note:
		return any(&Note{
			// Initialisation of associations
			// field is initialized with an instance of State with the name of the field
			States: []*State{{Name: "States"}},
			// field is initialized with an instance of Transition with the name of the field
			Transitions: []*Transition{{Name: "Transitions"}},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
		}).(*Type)
	case NoteStateShape:
		return any(&NoteStateShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of State with the name of the field
			State: &State{Name: "State"},
		}).(*Type)
	case NoteTransitionShape:
		return any(&NoteTransitionShape{
			// Initialisation of associations
			// field is initialized with an instance of Note with the name of the field
			Note: &Note{Name: "Note"},
			// field is initialized with an instance of Transition with the name of the field
			Transition: &Transition{Name: "Transition"},
		}).(*Type)
	case Object:
		return any(&Object{
			// Initialisation of associations
			// field is initialized with an instance of State with the name of the field
			State: &State{Name: "State"},
			// field is initialized with an instance of Message with the name of the field
			Messages: []*Message{{Name: "Messages"}},
		}).(*Type)
	case Role:
		return any(&Role{
			// Initialisation of associations
			// field is initialized with an instance of Role with the name of the field
			RolesWithSamePermissions: []*Role{{Name: "RolesWithSamePermissions"}},
		}).(*Type)
	case State:
		return any(&State{
			// Initialisation of associations
			// field is initialized with an instance of State with the name of the field
			Parent: &State{Name: "Parent"},
			// field is initialized with an instance of State with the name of the field
			SubStates: []*State{{Name: "SubStates"}},
			// field is initialized with an instance of Diagram with the name of the field
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			// field is initialized with an instance of Action with the name of the field
			Entry: &Action{Name: "Entry"},
			// field is initialized with an instance of Activities with the name of the field
			Activities: []*Activities{{Name: "Activities"}},
			// field is initialized with an instance of Action with the name of the field
			Exit: &Action{Name: "Exit"},
		}).(*Type)
	case StateMachine:
		return any(&StateMachine{
			// Initialisation of associations
			// field is initialized with an instance of State with the name of the field
			States: []*State{{Name: "States"}},
			// field is initialized with an instance of Diagram with the name of the field
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			// field is initialized with an instance of State with the name of the field
			InitialState: &State{Name: "InitialState"},
		}).(*Type)
	case StateShape:
		return any(&StateShape{
			// Initialisation of associations
			// field is initialized with an instance of State with the name of the field
			State: &State{Name: "State"},
		}).(*Type)
	case Transition:
		return any(&Transition{
			// Initialisation of associations
			// field is initialized with an instance of State with the name of the field
			Start: &State{Name: "Start"},
			// field is initialized with an instance of State with the name of the field
			End: &State{Name: "End"},
			// field is initialized with an instance of Role with the name of the field
			RolesWithPermissions: []*Role{{Name: "RolesWithPermissions"}},
			// field is initialized with an instance of MessageType with the name of the field
			GeneratedMessages: []*MessageType{{Name: "GeneratedMessages"}},
			// field is initialized with an instance of Guard with the name of the field
			Guard: &Guard{Name: "Guard"},
			// field is initialized with an instance of Diagram with the name of the field
			Diagrams: []*Diagram{{Name: "Diagrams"}},
		}).(*Type)
	case Transition_Shape:
		return any(&Transition_Shape{
			// Initialisation of associations
			// field is initialized with an instance of Transition with the name of the field
			Transition: &Transition{Name: "Transition"},
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
	// reverse maps of direct associations of Action
	case Action:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Activities
	case Activities:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Architecture
	case Architecture:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Guard
	case Guard:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Kill
	case Kill:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Library
	case Library:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Message
	case Message:
		switch fieldname {
		// insertion point for per direct association field
		case "MessageType":
			res := make(map[*MessageType][]*Message)
			for message := range stage.Messages {
				if message.MessageType != nil {
					messagetype_ := message.MessageType
					var messages []*Message
					_, ok := res[messagetype_]
					if ok {
						messages = res[messagetype_]
					} else {
						messages = make([]*Message, 0)
					}
					messages = append(messages, message)
					res[messagetype_] = messages
				}
			}
			return any(res).(map[*End][]*Start)
		case "OriginTransition":
			res := make(map[*Transition][]*Message)
			for message := range stage.Messages {
				if message.OriginTransition != nil {
					transition_ := message.OriginTransition
					var messages []*Message
					_, ok := res[transition_]
					if ok {
						messages = res[transition_]
					} else {
						messages = make([]*Message, 0)
					}
					messages = append(messages, message)
					res[transition_] = messages
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MessageType
	case MessageType:
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
	// reverse maps of direct associations of NoteStateShape
	case NoteStateShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteStateShape)
			for notestateshape := range stage.NoteStateShapes {
				if notestateshape.Note != nil {
					note_ := notestateshape.Note
					var notestateshapes []*NoteStateShape
					_, ok := res[note_]
					if ok {
						notestateshapes = res[note_]
					} else {
						notestateshapes = make([]*NoteStateShape, 0)
					}
					notestateshapes = append(notestateshapes, notestateshape)
					res[note_] = notestateshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "State":
			res := make(map[*State][]*NoteStateShape)
			for notestateshape := range stage.NoteStateShapes {
				if notestateshape.State != nil {
					state_ := notestateshape.State
					var notestateshapes []*NoteStateShape
					_, ok := res[state_]
					if ok {
						notestateshapes = res[state_]
					} else {
						notestateshapes = make([]*NoteStateShape, 0)
					}
					notestateshapes = append(notestateshapes, notestateshape)
					res[state_] = notestateshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteTransitionShape
	case NoteTransitionShape:
		switch fieldname {
		// insertion point for per direct association field
		case "Note":
			res := make(map[*Note][]*NoteTransitionShape)
			for notetransitionshape := range stage.NoteTransitionShapes {
				if notetransitionshape.Note != nil {
					note_ := notetransitionshape.Note
					var notetransitionshapes []*NoteTransitionShape
					_, ok := res[note_]
					if ok {
						notetransitionshapes = res[note_]
					} else {
						notetransitionshapes = make([]*NoteTransitionShape, 0)
					}
					notetransitionshapes = append(notetransitionshapes, notetransitionshape)
					res[note_] = notetransitionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		case "Transition":
			res := make(map[*Transition][]*NoteTransitionShape)
			for notetransitionshape := range stage.NoteTransitionShapes {
				if notetransitionshape.Transition != nil {
					transition_ := notetransitionshape.Transition
					var notetransitionshapes []*NoteTransitionShape
					_, ok := res[transition_]
					if ok {
						notetransitionshapes = res[transition_]
					} else {
						notetransitionshapes = make([]*NoteTransitionShape, 0)
					}
					notetransitionshapes = append(notetransitionshapes, notetransitionshape)
					res[transition_] = notetransitionshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Object
	case Object:
		switch fieldname {
		// insertion point for per direct association field
		case "State":
			res := make(map[*State][]*Object)
			for object := range stage.Objects {
				if object.State != nil {
					state_ := object.State
					var objects []*Object
					_, ok := res[state_]
					if ok {
						objects = res[state_]
					} else {
						objects = make([]*Object, 0)
					}
					objects = append(objects, object)
					res[state_] = objects
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Role
	case Role:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of State
	case State:
		switch fieldname {
		// insertion point for per direct association field
		case "Parent":
			res := make(map[*State][]*State)
			for state := range stage.States {
				if state.Parent != nil {
					state_ := state.Parent
					var states []*State
					_, ok := res[state_]
					if ok {
						states = res[state_]
					} else {
						states = make([]*State, 0)
					}
					states = append(states, state)
					res[state_] = states
				}
			}
			return any(res).(map[*End][]*Start)
		case "Entry":
			res := make(map[*Action][]*State)
			for state := range stage.States {
				if state.Entry != nil {
					action_ := state.Entry
					var states []*State
					_, ok := res[action_]
					if ok {
						states = res[action_]
					} else {
						states = make([]*State, 0)
					}
					states = append(states, state)
					res[action_] = states
				}
			}
			return any(res).(map[*End][]*Start)
		case "Exit":
			res := make(map[*Action][]*State)
			for state := range stage.States {
				if state.Exit != nil {
					action_ := state.Exit
					var states []*State
					_, ok := res[action_]
					if ok {
						states = res[action_]
					} else {
						states = make([]*State, 0)
					}
					states = append(states, state)
					res[action_] = states
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StateMachine
	case StateMachine:
		switch fieldname {
		// insertion point for per direct association field
		case "InitialState":
			res := make(map[*State][]*StateMachine)
			for statemachine := range stage.StateMachines {
				if statemachine.InitialState != nil {
					state_ := statemachine.InitialState
					var statemachines []*StateMachine
					_, ok := res[state_]
					if ok {
						statemachines = res[state_]
					} else {
						statemachines = make([]*StateMachine, 0)
					}
					statemachines = append(statemachines, statemachine)
					res[state_] = statemachines
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StateShape
	case StateShape:
		switch fieldname {
		// insertion point for per direct association field
		case "State":
			res := make(map[*State][]*StateShape)
			for stateshape := range stage.StateShapes {
				if stateshape.State != nil {
					state_ := stateshape.State
					var stateshapes []*StateShape
					_, ok := res[state_]
					if ok {
						stateshapes = res[state_]
					} else {
						stateshapes = make([]*StateShape, 0)
					}
					stateshapes = append(stateshapes, stateshape)
					res[state_] = stateshapes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Transition
	case Transition:
		switch fieldname {
		// insertion point for per direct association field
		case "Start":
			res := make(map[*State][]*Transition)
			for transition := range stage.Transitions {
				if transition.Start != nil {
					state_ := transition.Start
					var transitions []*Transition
					_, ok := res[state_]
					if ok {
						transitions = res[state_]
					} else {
						transitions = make([]*Transition, 0)
					}
					transitions = append(transitions, transition)
					res[state_] = transitions
				}
			}
			return any(res).(map[*End][]*Start)
		case "End":
			res := make(map[*State][]*Transition)
			for transition := range stage.Transitions {
				if transition.End != nil {
					state_ := transition.End
					var transitions []*Transition
					_, ok := res[state_]
					if ok {
						transitions = res[state_]
					} else {
						transitions = make([]*Transition, 0)
					}
					transitions = append(transitions, transition)
					res[state_] = transitions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Guard":
			res := make(map[*Guard][]*Transition)
			for transition := range stage.Transitions {
				if transition.Guard != nil {
					guard_ := transition.Guard
					var transitions []*Transition
					_, ok := res[guard_]
					if ok {
						transitions = res[guard_]
					} else {
						transitions = make([]*Transition, 0)
					}
					transitions = append(transitions, transition)
					res[guard_] = transitions
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Transition_Shape
	case Transition_Shape:
		switch fieldname {
		// insertion point for per direct association field
		case "Transition":
			res := make(map[*Transition][]*Transition_Shape)
			for transition_shape := range stage.Transition_Shapes {
				if transition_shape.Transition != nil {
					transition_ := transition_shape.Transition
					var transition_shapes []*Transition_Shape
					_, ok := res[transition_]
					if ok {
						transition_shapes = res[transition_]
					} else {
						transition_shapes = make([]*Transition_Shape, 0)
					}
					transition_shapes = append(transition_shapes, transition_shape)
					res[transition_] = transition_shapes
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
	// reverse maps of direct associations of Action
	case Action:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Activities
	case Activities:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Architecture
	case Architecture:
		switch fieldname {
		// insertion point for per direct association field
		case "StateMachines":
			res := make(map[*StateMachine][]*Architecture)
			for architecture := range stage.Architectures {
				for _, statemachine_ := range architecture.StateMachines {
					res[statemachine_] = append(res[statemachine_], architecture)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Roles":
			res := make(map[*Role][]*Architecture)
			for architecture := range stage.Architectures {
				for _, role_ := range architecture.Roles {
					res[role_] = append(res[role_], architecture)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
		case "NotesWhoseNodeIsExpanded":
			res := make(map[*Note][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, note_ := range diagram.NotesWhoseNodeIsExpanded {
					res[note_] = append(res[note_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "State_Shapes":
			res := make(map[*StateShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, stateshape_ := range diagram.State_Shapes {
					res[stateshape_] = append(res[stateshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "StatesWhoseNodeIsExpanded":
			res := make(map[*State][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, state_ := range diagram.StatesWhoseNodeIsExpanded {
					res[state_] = append(res[state_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Transition_Shapes":
			res := make(map[*Transition_Shape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, transition_shape_ := range diagram.Transition_Shapes {
					res[transition_shape_] = append(res[transition_shape_], diagram)
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
		case "NoteState_Shapes":
			res := make(map[*NoteStateShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, notestateshape_ := range diagram.NoteState_Shapes {
					res[notestateshape_] = append(res[notestateshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		case "NoteTransition_Shapes":
			res := make(map[*NoteTransitionShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, notetransitionshape_ := range diagram.NoteTransition_Shapes {
					res[notetransitionshape_] = append(res[notetransitionshape_], diagram)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Guard
	case Guard:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Kill
	case Kill:
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
		case "Diagrams":
			res := make(map[*Diagram][]*Library)
			for library := range stage.Librarys {
				for _, diagram_ := range library.Diagrams {
					res[diagram_] = append(res[diagram_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootStateMachines":
			res := make(map[*StateMachine][]*Library)
			for library := range stage.Librarys {
				for _, statemachine_ := range library.RootStateMachines {
					res[statemachine_] = append(res[statemachine_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "StateMachinesWhoseNodeIsExpanded":
			res := make(map[*StateMachine][]*Library)
			for library := range stage.Librarys {
				for _, statemachine_ := range library.StateMachinesWhoseNodeIsExpanded {
					res[statemachine_] = append(res[statemachine_], library)
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
		case "SubLibrariesWhoseNodeIsExpanded":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibrariesWhoseNodeIsExpanded {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Message
	case Message:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MessageType
	case MessageType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Note
	case Note:
		switch fieldname {
		// insertion point for per direct association field
		case "States":
			res := make(map[*State][]*Note)
			for note := range stage.Notes {
				for _, state_ := range note.States {
					res[state_] = append(res[state_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Transitions":
			res := make(map[*Transition][]*Note)
			for note := range stage.Notes {
				for _, transition_ := range note.Transitions {
					res[transition_] = append(res[transition_], note)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of NoteShape
	case NoteShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteStateShape
	case NoteStateShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of NoteTransitionShape
	case NoteTransitionShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Object
	case Object:
		switch fieldname {
		// insertion point for per direct association field
		case "Messages":
			res := make(map[*Message][]*Object)
			for object := range stage.Objects {
				for _, message_ := range object.Messages {
					res[message_] = append(res[message_], object)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Role
	case Role:
		switch fieldname {
		// insertion point for per direct association field
		case "RolesWithSamePermissions":
			res := make(map[*Role][]*Role)
			for role := range stage.Roles {
				for _, role_ := range role.RolesWithSamePermissions {
					res[role_] = append(res[role_], role)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of State
	case State:
		switch fieldname {
		// insertion point for per direct association field
		case "SubStates":
			res := make(map[*State][]*State)
			for state := range stage.States {
				for _, state_ := range state.SubStates {
					res[state_] = append(res[state_], state)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Diagrams":
			res := make(map[*Diagram][]*State)
			for state := range stage.States {
				for _, diagram_ := range state.Diagrams {
					res[diagram_] = append(res[diagram_], state)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Activities":
			res := make(map[*Activities][]*State)
			for state := range stage.States {
				for _, activities_ := range state.Activities {
					res[activities_] = append(res[activities_], state)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StateMachine
	case StateMachine:
		switch fieldname {
		// insertion point for per direct association field
		case "States":
			res := make(map[*State][]*StateMachine)
			for statemachine := range stage.StateMachines {
				for _, state_ := range statemachine.States {
					res[state_] = append(res[state_], statemachine)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Diagrams":
			res := make(map[*Diagram][]*StateMachine)
			for statemachine := range stage.StateMachines {
				for _, diagram_ := range statemachine.Diagrams {
					res[diagram_] = append(res[diagram_], statemachine)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of StateShape
	case StateShape:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Transition
	case Transition:
		switch fieldname {
		// insertion point for per direct association field
		case "RolesWithPermissions":
			res := make(map[*Role][]*Transition)
			for transition := range stage.Transitions {
				for _, role_ := range transition.RolesWithPermissions {
					res[role_] = append(res[role_], transition)
				}
			}
			return any(res).(map[*End][]*Start)
		case "GeneratedMessages":
			res := make(map[*MessageType][]*Transition)
			for transition := range stage.Transitions {
				for _, messagetype_ := range transition.GeneratedMessages {
					res[messagetype_] = append(res[messagetype_], transition)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Diagrams":
			res := make(map[*Diagram][]*Transition)
			for transition := range stage.Transitions {
				for _, diagram_ := range transition.Diagrams {
					res[diagram_] = append(res[diagram_], transition)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Transition_Shape
	case Transition_Shape:
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
	case *Action:
		res = "Action"
	case *Activities:
		res = "Activities"
	case *Architecture:
		res = "Architecture"
	case *Diagram:
		res = "Diagram"
	case *Guard:
		res = "Guard"
	case *Kill:
		res = "Kill"
	case *Library:
		res = "Library"
	case *Message:
		res = "Message"
	case *MessageType:
		res = "MessageType"
	case *Note:
		res = "Note"
	case *NoteShape:
		res = "NoteShape"
	case *NoteStateShape:
		res = "NoteStateShape"
	case *NoteTransitionShape:
		res = "NoteTransitionShape"
	case *Object:
		res = "Object"
	case *Role:
		res = "Role"
	case *State:
		res = "State"
	case *StateMachine:
		res = "StateMachine"
	case *StateShape:
		res = "StateShape"
	case *Transition:
		res = "Transition"
	case *Transition_Shape:
		res = "Transition_Shape"
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
	case *Action:
		var rf ReverseField
		_ = rf
	case *Activities:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "State"
		rf.Fieldname = "Activities"
		res = append(res, rf)
	case *Architecture:
		var rf ReverseField
		_ = rf
	case *Diagram:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
		rf.GongstructName = "State"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
		rf.GongstructName = "StateMachine"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
		rf.GongstructName = "Transition"
		rf.Fieldname = "Diagrams"
		res = append(res, rf)
	case *Guard:
		var rf ReverseField
		_ = rf
	case *Kill:
		var rf ReverseField
		_ = rf
	case *Library:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibraries"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "SubLibrariesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *Message:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Object"
		rf.Fieldname = "Messages"
		res = append(res, rf)
	case *MessageType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Transition"
		rf.Fieldname = "GeneratedMessages"
		res = append(res, rf)
	case *Note:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
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
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Note_Shapes"
		res = append(res, rf)
	case *NoteStateShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NoteState_Shapes"
		res = append(res, rf)
	case *NoteTransitionShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "NoteTransition_Shapes"
		res = append(res, rf)
	case *Object:
		var rf ReverseField
		_ = rf
	case *Role:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Architecture"
		rf.Fieldname = "Roles"
		res = append(res, rf)
		rf.GongstructName = "Role"
		rf.Fieldname = "RolesWithSamePermissions"
		res = append(res, rf)
		rf.GongstructName = "Transition"
		rf.Fieldname = "RolesWithPermissions"
		res = append(res, rf)
	case *State:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "StatesWhoseNodeIsExpanded"
		res = append(res, rf)
		rf.GongstructName = "Note"
		rf.Fieldname = "States"
		res = append(res, rf)
		rf.GongstructName = "State"
		rf.Fieldname = "SubStates"
		res = append(res, rf)
		rf.GongstructName = "StateMachine"
		rf.Fieldname = "States"
		res = append(res, rf)
	case *StateMachine:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Architecture"
		rf.Fieldname = "StateMachines"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "RootStateMachines"
		res = append(res, rf)
		rf.GongstructName = "Library"
		rf.Fieldname = "StateMachinesWhoseNodeIsExpanded"
		res = append(res, rf)
	case *StateShape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "State_Shapes"
		res = append(res, rf)
	case *Transition:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Note"
		rf.Fieldname = "Transitions"
		res = append(res, rf)
	case *Transition_Shape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Transition_Shapes"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (action *Action) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Criticality",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Criticality",
		},
	}
	return
}

func (activities *Activities) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Criticality",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "Criticality",
		},
	}
	return
}

func (architecture *Architecture) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "StateMachines",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StateMachine",
		},
		{
			Name:                 "Roles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Role",
		},
		{
			Name:               "NbPixPerCharacter",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEditable_",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsStatesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
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
			Name:                 "State_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StateShape",
		},
		{
			Name:                 "StatesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "State",
		},
		{
			Name:                 "Transition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Transition_Shape",
		},
		{
			Name:                 "Note_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteShape",
		},
		{
			Name:                 "NoteState_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteStateShape",
		},
		{
			Name:                 "NoteTransition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "NoteTransitionShape",
		},
	}
	return
}

func (guard *Guard) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (kill *Kill) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
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
			Name:                 "Diagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
		{
			Name:                 "RootStateMachines",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StateMachine",
		},
		{
			Name:               "IsStateMachinesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "StateMachinesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StateMachine",
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
			Name:               "IsSubLibrariesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubLibrariesWhoseNodeIsExpanded",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Library",
		},
		{
			Name:               "IsExpandedTmp",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (message *Message) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsSelected",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "MessageType",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "MessageType",
		},
		{
			Name:                 "OriginTransition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Transition",
		},
	}
	return
}

func (messagetype *MessageType) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "States",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "State",
		},
		{
			Name:                 "Transitions",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Transition",
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

func (notestateshape *NoteStateShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "State",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
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

func (notetransitionshape *NoteTransitionShape) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "Transition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Transition",
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

func (object *Object) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "State",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
		{
			Name:               "IsSelected",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "Rank",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "DOF",
			GongFieldValueType: GongFieldValueTypeDate,
		},
		{
			Name:                 "Messages",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Message",
		},
	}
	return
}

func (role *Role) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "RolesWithSamePermissions",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Role",
		},
	}
	return
}

func (state *State) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Parent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
		{
			Name:               "IsDecisionNode",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsFictious",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsEndState",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubStates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "State",
		},
		{
			Name:                 "Diagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
		{
			Name:                 "Entry",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Action",
		},
		{
			Name:                 "Activities",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Activities",
		},
		{
			Name:                 "Exit",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Action",
		},
	}
	return
}

func (statemachine *StateMachine) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "States",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "State",
		},
		{
			Name:                 "Diagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
		{
			Name:                 "InitialState",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
	}
	return
}

func (stateshape *StateShape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "State",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
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

func (transition *Transition) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Start",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
		{
			Name:                 "End",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
		{
			Name:                 "RolesWithPermissions",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Role",
		},
		{
			Name:                 "GeneratedMessages",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MessageType",
		},
		{
			Name:                 "Guard",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Guard",
		},
		{
			Name:                 "Diagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
	}
	return
}

func (transition_shape *Transition_Shape) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Transition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Transition",
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
func (action *Action) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = action.Name
	case "Criticality":
		enum := action.Criticality
		res.valueString = enum.ToCodeString()
	}
	return
}

func (activities *Activities) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = activities.Name
	case "Criticality":
		enum := activities.Criticality
		res.valueString = enum.ToCodeString()
	}
	return
}

func (architecture *Architecture) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = architecture.Name
	case "StateMachines":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range architecture.StateMachines {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Roles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range architecture.Roles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NbPixPerCharacter":
		res.valueString = fmt.Sprintf("%f", architecture.NbPixPerCharacter)
		res.valueFloat = architecture.NbPixPerCharacter
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (diagram *Diagram) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = diagram.Name
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", diagram.IsChecked)
		res.valueBool = diagram.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsExpanded)
		res.valueBool = diagram.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEditable_":
		res.valueString = fmt.Sprintf("%t", diagram.IsEditable_)
		res.valueBool = diagram.IsEditable_
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsStatesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsStatesNodeExpanded)
		res.valueBool = diagram.IsStatesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsNotesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", diagram.IsNotesNodeExpanded)
		res.valueBool = diagram.IsNotesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "State_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.State_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "StatesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.StatesWhoseNodeIsExpanded {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Transition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Transition_Shapes {
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
	case "NoteState_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.NoteState_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "NoteTransition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.NoteTransition_Shapes {
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

func (guard *Guard) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = guard.Name
	}
	return
}

func (kill *Kill) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = kill.Name
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
	case "RootStateMachines":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.RootStateMachines {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsStateMachinesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsStateMachinesNodeExpanded)
		res.valueBool = library.IsStateMachinesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "StateMachinesWhoseNodeIsExpanded":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.StateMachinesWhoseNodeIsExpanded {
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
	case "IsExpandedTmp":
		res.valueString = fmt.Sprintf("%t", library.IsExpandedTmp)
		res.valueBool = library.IsExpandedTmp
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (message *Message) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = message.Name
	case "IsSelected":
		res.valueString = fmt.Sprintf("%t", message.IsSelected)
		res.valueBool = message.IsSelected
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MessageType":
		res.GongFieldValueType = GongFieldValueTypePointer
		if message.MessageType != nil {
			res.valueString = message.MessageType.Name
			res.ids = message.MessageType.GongGetUUID(stage)
		}
	case "OriginTransition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if message.OriginTransition != nil {
			res.valueString = message.OriginTransition.Name
			res.ids = message.OriginTransition.GongGetUUID(stage)
		}
	}
	return
}

func (messagetype *MessageType) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = messagetype.Name
	case "Description":
		res.valueString = messagetype.Description
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
	case "States":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.States {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Transitions":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range note.Transitions {
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

func (notestateshape *NoteStateShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notestateshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notestateshape.Note != nil {
			res.valueString = notestateshape.Note.Name
			res.ids = notestateshape.Note.GongGetUUID(stage)
		}
	case "State":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notestateshape.State != nil {
			res.valueString = notestateshape.State.Name
			res.ids = notestateshape.State.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", notestateshape.StartRatio)
		res.valueFloat = notestateshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", notestateshape.EndRatio)
		res.valueFloat = notestateshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := notestateshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := notestateshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", notestateshape.CornerOffsetRatio)
		res.valueFloat = notestateshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", notestateshape.IsHidden)
		res.valueBool = notestateshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (notetransitionshape *NoteTransitionShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = notetransitionshape.Name
	case "Note":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notetransitionshape.Note != nil {
			res.valueString = notetransitionshape.Note.Name
			res.ids = notetransitionshape.Note.GongGetUUID(stage)
		}
	case "Transition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if notetransitionshape.Transition != nil {
			res.valueString = notetransitionshape.Transition.Name
			res.ids = notetransitionshape.Transition.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", notetransitionshape.StartRatio)
		res.valueFloat = notetransitionshape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", notetransitionshape.EndRatio)
		res.valueFloat = notetransitionshape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := notetransitionshape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := notetransitionshape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", notetransitionshape.CornerOffsetRatio)
		res.valueFloat = notetransitionshape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", notetransitionshape.IsHidden)
		res.valueBool = notetransitionshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (object *Object) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = object.Name
	case "State":
		res.GongFieldValueType = GongFieldValueTypePointer
		if object.State != nil {
			res.valueString = object.State.Name
			res.ids = object.State.GongGetUUID(stage)
		}
	case "IsSelected":
		res.valueString = fmt.Sprintf("%t", object.IsSelected)
		res.valueBool = object.IsSelected
		res.GongFieldValueType = GongFieldValueTypeBool
	case "Rank":
		res.valueString = fmt.Sprintf("%d", object.Rank)
		res.valueInt = object.Rank
		res.GongFieldValueType = GongFieldValueTypeInt
	case "DOF":
		res.valueString = object.DOF.String()
	case "Messages":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range object.Messages {
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

func (role *Role) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = role.Name
	case "Acronym":
		res.valueString = role.Acronym
	case "RolesWithSamePermissions":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range role.RolesWithSamePermissions {
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

func (state *State) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = state.Name
	case "Parent":
		res.GongFieldValueType = GongFieldValueTypePointer
		if state.Parent != nil {
			res.valueString = state.Parent.Name
			res.ids = state.Parent.GongGetUUID(stage)
		}
	case "IsDecisionNode":
		res.valueString = fmt.Sprintf("%t", state.IsDecisionNode)
		res.valueBool = state.IsDecisionNode
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsFictious":
		res.valueString = fmt.Sprintf("%t", state.IsFictious)
		res.valueBool = state.IsFictious
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsEndState":
		res.valueString = fmt.Sprintf("%t", state.IsEndState)
		res.valueBool = state.IsEndState
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SubStates":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range state.SubStates {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range state.Diagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Entry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if state.Entry != nil {
			res.valueString = state.Entry.Name
			res.ids = state.Entry.GongGetUUID(stage)
		}
	case "Activities":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range state.Activities {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Exit":
		res.GongFieldValueType = GongFieldValueTypePointer
		if state.Exit != nil {
			res.valueString = state.Exit.Name
			res.ids = state.Exit.GongGetUUID(stage)
		}
	}
	return
}

func (statemachine *StateMachine) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = statemachine.Name
	case "ComputedPrefix":
		res.valueString = statemachine.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", statemachine.IsExpanded)
		res.valueBool = statemachine.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LayoutDirection":
		enum := statemachine.LayoutDirection
		res.valueString = enum.ToCodeString()
	case "States":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range statemachine.States {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range statemachine.Diagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "InitialState":
		res.GongFieldValueType = GongFieldValueTypePointer
		if statemachine.InitialState != nil {
			res.valueString = statemachine.InitialState.Name
			res.ids = statemachine.InitialState.GongGetUUID(stage)
		}
	}
	return
}

func (stateshape *StateShape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = stateshape.Name
	case "State":
		res.GongFieldValueType = GongFieldValueTypePointer
		if stateshape.State != nil {
			res.valueString = stateshape.State.Name
			res.ids = stateshape.State.GongGetUUID(stage)
		}
	case "X":
		res.valueString = fmt.Sprintf("%f", stateshape.X)
		res.valueFloat = stateshape.X
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Y":
		res.valueString = fmt.Sprintf("%f", stateshape.Y)
		res.valueFloat = stateshape.Y
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Width":
		res.valueString = fmt.Sprintf("%f", stateshape.Width)
		res.valueFloat = stateshape.Width
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Height":
		res.valueString = fmt.Sprintf("%f", stateshape.Height)
		res.valueFloat = stateshape.Height
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", stateshape.IsHidden)
		res.valueBool = stateshape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (transition *Transition) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = transition.Name
	case "Start":
		res.GongFieldValueType = GongFieldValueTypePointer
		if transition.Start != nil {
			res.valueString = transition.Start.Name
			res.ids = transition.Start.GongGetUUID(stage)
		}
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if transition.End != nil {
			res.valueString = transition.End.Name
			res.ids = transition.End.GongGetUUID(stage)
		}
	case "RolesWithPermissions":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range transition.RolesWithPermissions {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "GeneratedMessages":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range transition.GeneratedMessages {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Guard":
		res.GongFieldValueType = GongFieldValueTypePointer
		if transition.Guard != nil {
			res.valueString = transition.Guard.Name
			res.ids = transition.Guard.GongGetUUID(stage)
		}
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range transition.Diagrams {
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

func (transition_shape *Transition_Shape) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = transition_shape.Name
	case "Transition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if transition_shape.Transition != nil {
			res.valueString = transition_shape.Transition.Name
			res.ids = transition_shape.Transition.GongGetUUID(stage)
		}
	case "StartRatio":
		res.valueString = fmt.Sprintf("%f", transition_shape.StartRatio)
		res.valueFloat = transition_shape.StartRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "EndRatio":
		res.valueString = fmt.Sprintf("%f", transition_shape.EndRatio)
		res.valueFloat = transition_shape.EndRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StartOrientation":
		enum := transition_shape.StartOrientation
		res.valueString = enum.ToCodeString()
	case "EndOrientation":
		enum := transition_shape.EndOrientation
		res.valueString = enum.ToCodeString()
	case "CornerOffsetRatio":
		res.valueString = fmt.Sprintf("%f", transition_shape.CornerOffsetRatio)
		res.valueFloat = transition_shape.CornerOffsetRatio
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "IsHidden":
		res.valueString = fmt.Sprintf("%t", transition_shape.IsHidden)
		res.valueBool = transition_shape.IsHidden
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (action *Action) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		action.Name = value.GetValueString()
	case "Criticality":
		action.Criticality.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (activities *Activities) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		activities.Name = value.GetValueString()
	case "Criticality":
		activities.Criticality.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (architecture *Architecture) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		architecture.Name = value.GetValueString()
	case "StateMachines":
		architecture.StateMachines = make([]*StateMachine, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StateMachines {
					if stage.StateMachine_stagedOrder[__instance__] == uint(id) {
						architecture.StateMachines = append(architecture.StateMachines, __instance__)
						break
					}
				}
			}
		}
	case "Roles":
		architecture.Roles = make([]*Role, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Roles {
					if stage.Role_stagedOrder[__instance__] == uint(id) {
						architecture.Roles = append(architecture.Roles, __instance__)
						break
					}
				}
			}
		}
	case "NbPixPerCharacter":
		architecture.NbPixPerCharacter = value.GetValueFloat()
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
	case "IsChecked":
		diagram.IsChecked = value.GetValueBool()
	case "IsExpanded":
		diagram.IsExpanded = value.GetValueBool()
	case "IsEditable_":
		diagram.IsEditable_ = value.GetValueBool()
	case "IsStatesNodeExpanded":
		diagram.IsStatesNodeExpanded = value.GetValueBool()
	case "IsNotesNodeExpanded":
		diagram.IsNotesNodeExpanded = value.GetValueBool()
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
	case "State_Shapes":
		diagram.State_Shapes = make([]*StateShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StateShapes {
					if stage.StateShape_stagedOrder[__instance__] == uint(id) {
						diagram.State_Shapes = append(diagram.State_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "StatesWhoseNodeIsExpanded":
		diagram.StatesWhoseNodeIsExpanded = make([]*State, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.States {
					if stage.State_stagedOrder[__instance__] == uint(id) {
						diagram.StatesWhoseNodeIsExpanded = append(diagram.StatesWhoseNodeIsExpanded, __instance__)
						break
					}
				}
			}
		}
	case "Transition_Shapes":
		diagram.Transition_Shapes = make([]*Transition_Shape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Transition_Shapes {
					if stage.Transition_Shape_stagedOrder[__instance__] == uint(id) {
						diagram.Transition_Shapes = append(diagram.Transition_Shapes, __instance__)
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
	case "NoteState_Shapes":
		diagram.NoteState_Shapes = make([]*NoteStateShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteStateShapes {
					if stage.NoteStateShape_stagedOrder[__instance__] == uint(id) {
						diagram.NoteState_Shapes = append(diagram.NoteState_Shapes, __instance__)
						break
					}
				}
			}
		}
	case "NoteTransition_Shapes":
		diagram.NoteTransition_Shapes = make([]*NoteTransitionShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.NoteTransitionShapes {
					if stage.NoteTransitionShape_stagedOrder[__instance__] == uint(id) {
						diagram.NoteTransition_Shapes = append(diagram.NoteTransition_Shapes, __instance__)
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

func (guard *Guard) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		guard.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (kill *Kill) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		kill.Name = value.GetValueString()
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
	case "RootStateMachines":
		library.RootStateMachines = make([]*StateMachine, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StateMachines {
					if stage.StateMachine_stagedOrder[__instance__] == uint(id) {
						library.RootStateMachines = append(library.RootStateMachines, __instance__)
						break
					}
				}
			}
		}
	case "IsStateMachinesNodeExpanded":
		library.IsStateMachinesNodeExpanded = value.GetValueBool()
	case "StateMachinesWhoseNodeIsExpanded":
		library.StateMachinesWhoseNodeIsExpanded = make([]*StateMachine, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StateMachines {
					if stage.StateMachine_stagedOrder[__instance__] == uint(id) {
						library.StateMachinesWhoseNodeIsExpanded = append(library.StateMachinesWhoseNodeIsExpanded, __instance__)
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
	case "IsExpandedTmp":
		library.IsExpandedTmp = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (message *Message) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		message.Name = value.GetValueString()
	case "IsSelected":
		message.IsSelected = value.GetValueBool()
	case "MessageType":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			message.MessageType = nil
			for __instance__ := range stage.MessageTypes {
				if stage.MessageType_stagedOrder[__instance__] == uint(id) {
					message.MessageType = __instance__
					break
				}
			}
		}
	case "OriginTransition":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			message.OriginTransition = nil
			for __instance__ := range stage.Transitions {
				if stage.Transition_stagedOrder[__instance__] == uint(id) {
					message.OriginTransition = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (messagetype *MessageType) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		messagetype.Name = value.GetValueString()
	case "Description":
		messagetype.Description = value.GetValueString()
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
	case "States":
		note.States = make([]*State, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.States {
					if stage.State_stagedOrder[__instance__] == uint(id) {
						note.States = append(note.States, __instance__)
						break
					}
				}
			}
		}
	case "Transitions":
		note.Transitions = make([]*Transition, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Transitions {
					if stage.Transition_stagedOrder[__instance__] == uint(id) {
						note.Transitions = append(note.Transitions, __instance__)
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

func (notestateshape *NoteStateShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		notestateshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notestateshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					notestateshape.Note = __instance__
					break
				}
			}
		}
	case "State":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notestateshape.State = nil
			for __instance__ := range stage.States {
				if stage.State_stagedOrder[__instance__] == uint(id) {
					notestateshape.State = __instance__
					break
				}
			}
		}
	case "StartRatio":
		notestateshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		notestateshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		notestateshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		notestateshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		notestateshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		notestateshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (notetransitionshape *NoteTransitionShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		notetransitionshape.Name = value.GetValueString()
	case "Note":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notetransitionshape.Note = nil
			for __instance__ := range stage.Notes {
				if stage.Note_stagedOrder[__instance__] == uint(id) {
					notetransitionshape.Note = __instance__
					break
				}
			}
		}
	case "Transition":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			notetransitionshape.Transition = nil
			for __instance__ := range stage.Transitions {
				if stage.Transition_stagedOrder[__instance__] == uint(id) {
					notetransitionshape.Transition = __instance__
					break
				}
			}
		}
	case "StartRatio":
		notetransitionshape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		notetransitionshape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		notetransitionshape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		notetransitionshape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		notetransitionshape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		notetransitionshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (object *Object) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		object.Name = value.GetValueString()
	case "State":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			object.State = nil
			for __instance__ := range stage.States {
				if stage.State_stagedOrder[__instance__] == uint(id) {
					object.State = __instance__
					break
				}
			}
		}
	case "IsSelected":
		object.IsSelected = value.GetValueBool()
	case "Rank":
		object.Rank = int(value.GetValueInt())
	case "Messages":
		object.Messages = make([]*Message, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Messages {
					if stage.Message_stagedOrder[__instance__] == uint(id) {
						object.Messages = append(object.Messages, __instance__)
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

func (role *Role) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		role.Name = value.GetValueString()
	case "Acronym":
		role.Acronym = value.GetValueString()
	case "RolesWithSamePermissions":
		role.RolesWithSamePermissions = make([]*Role, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Roles {
					if stage.Role_stagedOrder[__instance__] == uint(id) {
						role.RolesWithSamePermissions = append(role.RolesWithSamePermissions, __instance__)
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

func (state *State) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		state.Name = value.GetValueString()
	case "Parent":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			state.Parent = nil
			for __instance__ := range stage.States {
				if stage.State_stagedOrder[__instance__] == uint(id) {
					state.Parent = __instance__
					break
				}
			}
		}
	case "IsDecisionNode":
		state.IsDecisionNode = value.GetValueBool()
	case "IsFictious":
		state.IsFictious = value.GetValueBool()
	case "IsEndState":
		state.IsEndState = value.GetValueBool()
	case "SubStates":
		state.SubStates = make([]*State, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.States {
					if stage.State_stagedOrder[__instance__] == uint(id) {
						state.SubStates = append(state.SubStates, __instance__)
						break
					}
				}
			}
		}
	case "Diagrams":
		state.Diagrams = make([]*Diagram, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Diagrams {
					if stage.Diagram_stagedOrder[__instance__] == uint(id) {
						state.Diagrams = append(state.Diagrams, __instance__)
						break
					}
				}
			}
		}
	case "Entry":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			state.Entry = nil
			for __instance__ := range stage.Actions {
				if stage.Action_stagedOrder[__instance__] == uint(id) {
					state.Entry = __instance__
					break
				}
			}
		}
	case "Activities":
		state.Activities = make([]*Activities, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Activitiess {
					if stage.Activities_stagedOrder[__instance__] == uint(id) {
						state.Activities = append(state.Activities, __instance__)
						break
					}
				}
			}
		}
	case "Exit":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			state.Exit = nil
			for __instance__ := range stage.Actions {
				if stage.Action_stagedOrder[__instance__] == uint(id) {
					state.Exit = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (statemachine *StateMachine) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		statemachine.Name = value.GetValueString()
	case "ComputedPrefix":
		statemachine.ComputedPrefix = value.GetValueString()
	case "IsExpanded":
		statemachine.IsExpanded = value.GetValueBool()
	case "LayoutDirection":
		statemachine.LayoutDirection.FromCodeString(value.GetValueString())
	case "States":
		statemachine.States = make([]*State, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.States {
					if stage.State_stagedOrder[__instance__] == uint(id) {
						statemachine.States = append(statemachine.States, __instance__)
						break
					}
				}
			}
		}
	case "Diagrams":
		statemachine.Diagrams = make([]*Diagram, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Diagrams {
					if stage.Diagram_stagedOrder[__instance__] == uint(id) {
						statemachine.Diagrams = append(statemachine.Diagrams, __instance__)
						break
					}
				}
			}
		}
	case "InitialState":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			statemachine.InitialState = nil
			for __instance__ := range stage.States {
				if stage.State_stagedOrder[__instance__] == uint(id) {
					statemachine.InitialState = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (stateshape *StateShape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		stateshape.Name = value.GetValueString()
	case "State":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			stateshape.State = nil
			for __instance__ := range stage.States {
				if stage.State_stagedOrder[__instance__] == uint(id) {
					stateshape.State = __instance__
					break
				}
			}
		}
	case "X":
		stateshape.X = value.GetValueFloat()
	case "Y":
		stateshape.Y = value.GetValueFloat()
	case "Width":
		stateshape.Width = value.GetValueFloat()
	case "Height":
		stateshape.Height = value.GetValueFloat()
	case "IsHidden":
		stateshape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (transition *Transition) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		transition.Name = value.GetValueString()
	case "Start":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			transition.Start = nil
			for __instance__ := range stage.States {
				if stage.State_stagedOrder[__instance__] == uint(id) {
					transition.Start = __instance__
					break
				}
			}
		}
	case "End":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			transition.End = nil
			for __instance__ := range stage.States {
				if stage.State_stagedOrder[__instance__] == uint(id) {
					transition.End = __instance__
					break
				}
			}
		}
	case "RolesWithPermissions":
		transition.RolesWithPermissions = make([]*Role, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Roles {
					if stage.Role_stagedOrder[__instance__] == uint(id) {
						transition.RolesWithPermissions = append(transition.RolesWithPermissions, __instance__)
						break
					}
				}
			}
		}
	case "GeneratedMessages":
		transition.GeneratedMessages = make([]*MessageType, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.MessageTypes {
					if stage.MessageType_stagedOrder[__instance__] == uint(id) {
						transition.GeneratedMessages = append(transition.GeneratedMessages, __instance__)
						break
					}
				}
			}
		}
	case "Guard":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			transition.Guard = nil
			for __instance__ := range stage.Guards {
				if stage.Guard_stagedOrder[__instance__] == uint(id) {
					transition.Guard = __instance__
					break
				}
			}
		}
	case "Diagrams":
		transition.Diagrams = make([]*Diagram, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Diagrams {
					if stage.Diagram_stagedOrder[__instance__] == uint(id) {
						transition.Diagrams = append(transition.Diagrams, __instance__)
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

func (transition_shape *Transition_Shape) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		transition_shape.Name = value.GetValueString()
	case "Transition":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			transition_shape.Transition = nil
			for __instance__ := range stage.Transitions {
				if stage.Transition_stagedOrder[__instance__] == uint(id) {
					transition_shape.Transition = __instance__
					break
				}
			}
		}
	case "StartRatio":
		transition_shape.StartRatio = value.GetValueFloat()
	case "EndRatio":
		transition_shape.EndRatio = value.GetValueFloat()
	case "StartOrientation":
		transition_shape.StartOrientation.FromCodeString(value.GetValueString())
	case "EndOrientation":
		transition_shape.EndOrientation.FromCodeString(value.GetValueString())
	case "CornerOffsetRatio":
		transition_shape.CornerOffsetRatio = value.GetValueFloat()
	case "IsHidden":
		transition_shape.IsHidden = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (action *Action) GongGetGongstructName() string {
	return "Action"
}

func (activities *Activities) GongGetGongstructName() string {
	return "Activities"
}

func (architecture *Architecture) GongGetGongstructName() string {
	return "Architecture"
}

func (diagram *Diagram) GongGetGongstructName() string {
	return "Diagram"
}

func (guard *Guard) GongGetGongstructName() string {
	return "Guard"
}

func (kill *Kill) GongGetGongstructName() string {
	return "Kill"
}

func (library *Library) GongGetGongstructName() string {
	return "Library"
}

func (message *Message) GongGetGongstructName() string {
	return "Message"
}

func (messagetype *MessageType) GongGetGongstructName() string {
	return "MessageType"
}

func (note *Note) GongGetGongstructName() string {
	return "Note"
}

func (noteshape *NoteShape) GongGetGongstructName() string {
	return "NoteShape"
}

func (notestateshape *NoteStateShape) GongGetGongstructName() string {
	return "NoteStateShape"
}

func (notetransitionshape *NoteTransitionShape) GongGetGongstructName() string {
	return "NoteTransitionShape"
}

func (object *Object) GongGetGongstructName() string {
	return "Object"
}

func (role *Role) GongGetGongstructName() string {
	return "Role"
}

func (state *State) GongGetGongstructName() string {
	return "State"
}

func (statemachine *StateMachine) GongGetGongstructName() string {
	return "StateMachine"
}

func (stateshape *StateShape) GongGetGongstructName() string {
	return "StateShape"
}

func (transition *Transition) GongGetGongstructName() string {
	return "Transition"
}

func (transition_shape *Transition_Shape) GongGetGongstructName() string {
	return "Transition_Shape"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.Actions_mapString = make(map[string]*Action)
	for action := range stage.Actions {
		stage.Actions_mapString[action.Name] = action
	}

	stage.Activitiess_mapString = make(map[string]*Activities)
	for activities := range stage.Activitiess {
		stage.Activitiess_mapString[activities.Name] = activities
	}

	stage.Architectures_mapString = make(map[string]*Architecture)
	for architecture := range stage.Architectures {
		stage.Architectures_mapString[architecture.Name] = architecture
	}

	stage.Diagrams_mapString = make(map[string]*Diagram)
	for diagram := range stage.Diagrams {
		stage.Diagrams_mapString[diagram.Name] = diagram
	}

	stage.Guards_mapString = make(map[string]*Guard)
	for guard := range stage.Guards {
		stage.Guards_mapString[guard.Name] = guard
	}

	stage.Kills_mapString = make(map[string]*Kill)
	for kill := range stage.Kills {
		stage.Kills_mapString[kill.Name] = kill
	}

	stage.Librarys_mapString = make(map[string]*Library)
	for library := range stage.Librarys {
		stage.Librarys_mapString[library.Name] = library
	}

	stage.Messages_mapString = make(map[string]*Message)
	for message := range stage.Messages {
		stage.Messages_mapString[message.Name] = message
	}

	stage.MessageTypes_mapString = make(map[string]*MessageType)
	for messagetype := range stage.MessageTypes {
		stage.MessageTypes_mapString[messagetype.Name] = messagetype
	}

	stage.Notes_mapString = make(map[string]*Note)
	for note := range stage.Notes {
		stage.Notes_mapString[note.Name] = note
	}

	stage.NoteShapes_mapString = make(map[string]*NoteShape)
	for noteshape := range stage.NoteShapes {
		stage.NoteShapes_mapString[noteshape.Name] = noteshape
	}

	stage.NoteStateShapes_mapString = make(map[string]*NoteStateShape)
	for notestateshape := range stage.NoteStateShapes {
		stage.NoteStateShapes_mapString[notestateshape.Name] = notestateshape
	}

	stage.NoteTransitionShapes_mapString = make(map[string]*NoteTransitionShape)
	for notetransitionshape := range stage.NoteTransitionShapes {
		stage.NoteTransitionShapes_mapString[notetransitionshape.Name] = notetransitionshape
	}

	stage.Objects_mapString = make(map[string]*Object)
	for object := range stage.Objects {
		stage.Objects_mapString[object.Name] = object
	}

	stage.Roles_mapString = make(map[string]*Role)
	for role := range stage.Roles {
		stage.Roles_mapString[role.Name] = role
	}

	stage.States_mapString = make(map[string]*State)
	for state := range stage.States {
		stage.States_mapString[state.Name] = state
	}

	stage.StateMachines_mapString = make(map[string]*StateMachine)
	for statemachine := range stage.StateMachines {
		stage.StateMachines_mapString[statemachine.Name] = statemachine
	}

	stage.StateShapes_mapString = make(map[string]*StateShape)
	for stateshape := range stage.StateShapes {
		stage.StateShapes_mapString[stateshape.Name] = stateshape
	}

	stage.Transitions_mapString = make(map[string]*Transition)
	for transition := range stage.Transitions {
		stage.Transitions_mapString[transition.Name] = transition
	}

	stage.Transition_Shapes_mapString = make(map[string]*Transition_Shape)
	for transition_shape := range stage.Transition_Shapes {
		stage.Transition_Shapes_mapString[transition_shape.Name] = transition_shape
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
