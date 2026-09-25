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
	Actions                map[*Action]struct{}
	Actions_instance       map[*Action]*Action
	Actions_mapString      map[string]*Action
	ActionOrder            uint
	Action_stagedOrder     map[*Action]uint
	Action_orderStaged     map[uint]*Action
	Actions_reference      map[*Action]*Action
	Actions_referenceOrder map[*Action]uint

	// insertion point for slice of pointers maps
	OnAfterActionCreateCallback GongOnAfterCreateInterface[Action]
	OnAfterActionUpdateCallback GongOnAfterUpdateInterface[Action]
	OnAfterActionDeleteCallback GongOnAfterDeleteInterface[Action]

	Activitiess                map[*Activities]struct{}
	Activitiess_instance       map[*Activities]*Activities
	Activitiess_mapString      map[string]*Activities
	ActivitiesOrder            uint
	Activities_stagedOrder     map[*Activities]uint
	Activities_orderStaged     map[uint]*Activities
	Activitiess_reference      map[*Activities]*Activities
	Activitiess_referenceOrder map[*Activities]uint

	// insertion point for slice of pointers maps
	OnAfterActivitiesCreateCallback GongOnAfterCreateInterface[Activities]
	OnAfterActivitiesUpdateCallback GongOnAfterUpdateInterface[Activities]
	OnAfterActivitiesDeleteCallback GongOnAfterDeleteInterface[Activities]

	Diagrams                map[*Diagram]struct{}
	Diagrams_instance       map[*Diagram]*Diagram
	Diagrams_mapString      map[string]*Diagram
	DiagramOrder            uint
	Diagram_stagedOrder     map[*Diagram]uint
	Diagram_orderStaged     map[uint]*Diagram
	Diagrams_reference      map[*Diagram]*Diagram
	Diagrams_referenceOrder map[*Diagram]uint

	// insertion point for slice of pointers maps
	Diagram_State_Shapes_reverseMap map[*StateShape]*Diagram

	Diagram_StatesWhoseNodeIsExpanded_reverseMap map[*State]*Diagram

	Diagram_Transition_Shapes_reverseMap map[*Transition_Shape]*Diagram

	Diagram_Note_Shapes_reverseMap map[*NoteShape]*Diagram

	Diagram_NoteState_Shapes_reverseMap map[*NoteStateShape]*Diagram

	OnAfterDiagramCreateCallback GongOnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback GongOnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback GongOnAfterDeleteInterface[Diagram]

	Guards                map[*Guard]struct{}
	Guards_instance       map[*Guard]*Guard
	Guards_mapString      map[string]*Guard
	GuardOrder            uint
	Guard_stagedOrder     map[*Guard]uint
	Guard_orderStaged     map[uint]*Guard
	Guards_reference      map[*Guard]*Guard
	Guards_referenceOrder map[*Guard]uint

	// insertion point for slice of pointers maps
	OnAfterGuardCreateCallback GongOnAfterCreateInterface[Guard]
	OnAfterGuardUpdateCallback GongOnAfterUpdateInterface[Guard]
	OnAfterGuardDeleteCallback GongOnAfterDeleteInterface[Guard]

	Kills                map[*Kill]struct{}
	Kills_instance       map[*Kill]*Kill
	Kills_mapString      map[string]*Kill
	KillOrder            uint
	Kill_stagedOrder     map[*Kill]uint
	Kill_orderStaged     map[uint]*Kill
	Kills_reference      map[*Kill]*Kill
	Kills_referenceOrder map[*Kill]uint

	// insertion point for slice of pointers maps
	OnAfterKillCreateCallback GongOnAfterCreateInterface[Kill]
	OnAfterKillUpdateCallback GongOnAfterUpdateInterface[Kill]
	OnAfterKillDeleteCallback GongOnAfterDeleteInterface[Kill]

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

	Library_SubLibrariesWhoseNodeIsExpanded_reverseMap map[*Library]*Library

	Library_Roles_reverseMap map[*Role]*Library

	Library_MessageTypes_reverseMap map[*MessageType]*Library

	OnAfterLibraryCreateCallback GongOnAfterCreateInterface[Library]
	OnAfterLibraryUpdateCallback GongOnAfterUpdateInterface[Library]
	OnAfterLibraryDeleteCallback GongOnAfterDeleteInterface[Library]

	Messages                map[*Message]struct{}
	Messages_instance       map[*Message]*Message
	Messages_mapString      map[string]*Message
	MessageOrder            uint
	Message_stagedOrder     map[*Message]uint
	Message_orderStaged     map[uint]*Message
	Messages_reference      map[*Message]*Message
	Messages_referenceOrder map[*Message]uint

	// insertion point for slice of pointers maps
	OnAfterMessageCreateCallback GongOnAfterCreateInterface[Message]
	OnAfterMessageUpdateCallback GongOnAfterUpdateInterface[Message]
	OnAfterMessageDeleteCallback GongOnAfterDeleteInterface[Message]

	MessageTypes                map[*MessageType]struct{}
	MessageTypes_instance       map[*MessageType]*MessageType
	MessageTypes_mapString      map[string]*MessageType
	MessageTypeOrder            uint
	MessageType_stagedOrder     map[*MessageType]uint
	MessageType_orderStaged     map[uint]*MessageType
	MessageTypes_reference      map[*MessageType]*MessageType
	MessageTypes_referenceOrder map[*MessageType]uint

	// insertion point for slice of pointers maps
	OnAfterMessageTypeCreateCallback GongOnAfterCreateInterface[MessageType]
	OnAfterMessageTypeUpdateCallback GongOnAfterUpdateInterface[MessageType]
	OnAfterMessageTypeDeleteCallback GongOnAfterDeleteInterface[MessageType]

	Notes                map[*Note]struct{}
	Notes_instance       map[*Note]*Note
	Notes_mapString      map[string]*Note
	NoteOrder            uint
	Note_stagedOrder     map[*Note]uint
	Note_orderStaged     map[uint]*Note
	Notes_reference      map[*Note]*Note
	Notes_referenceOrder map[*Note]uint

	// insertion point for slice of pointers maps
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

	NoteStateShapes                map[*NoteStateShape]struct{}
	NoteStateShapes_instance       map[*NoteStateShape]*NoteStateShape
	NoteStateShapes_mapString      map[string]*NoteStateShape
	NoteStateShapeOrder            uint
	NoteStateShape_stagedOrder     map[*NoteStateShape]uint
	NoteStateShape_orderStaged     map[uint]*NoteStateShape
	NoteStateShapes_reference      map[*NoteStateShape]*NoteStateShape
	NoteStateShapes_referenceOrder map[*NoteStateShape]uint

	// insertion point for slice of pointers maps
	OnAfterNoteStateShapeCreateCallback GongOnAfterCreateInterface[NoteStateShape]
	OnAfterNoteStateShapeUpdateCallback GongOnAfterUpdateInterface[NoteStateShape]
	OnAfterNoteStateShapeDeleteCallback GongOnAfterDeleteInterface[NoteStateShape]

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

	OnAfterObjectCreateCallback GongOnAfterCreateInterface[Object]
	OnAfterObjectUpdateCallback GongOnAfterUpdateInterface[Object]
	OnAfterObjectDeleteCallback GongOnAfterDeleteInterface[Object]

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

	OnAfterRoleCreateCallback GongOnAfterCreateInterface[Role]
	OnAfterRoleUpdateCallback GongOnAfterUpdateInterface[Role]
	OnAfterRoleDeleteCallback GongOnAfterDeleteInterface[Role]

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

	State_Activities_reverseMap map[*Activities]*State

	State_Diagrams_reverseMap map[*Diagram]*State

	State_Notes_reverseMap map[*Note]*State

	OnAfterStateCreateCallback GongOnAfterCreateInterface[State]
	OnAfterStateUpdateCallback GongOnAfterUpdateInterface[State]
	OnAfterStateDeleteCallback GongOnAfterDeleteInterface[State]

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

	OnAfterStateMachineCreateCallback GongOnAfterCreateInterface[StateMachine]
	OnAfterStateMachineUpdateCallback GongOnAfterUpdateInterface[StateMachine]
	OnAfterStateMachineDeleteCallback GongOnAfterDeleteInterface[StateMachine]

	StateShapes                map[*StateShape]struct{}
	StateShapes_instance       map[*StateShape]*StateShape
	StateShapes_mapString      map[string]*StateShape
	StateShapeOrder            uint
	StateShape_stagedOrder     map[*StateShape]uint
	StateShape_orderStaged     map[uint]*StateShape
	StateShapes_reference      map[*StateShape]*StateShape
	StateShapes_referenceOrder map[*StateShape]uint

	// insertion point for slice of pointers maps
	OnAfterStateShapeCreateCallback GongOnAfterCreateInterface[StateShape]
	OnAfterStateShapeUpdateCallback GongOnAfterUpdateInterface[StateShape]
	OnAfterStateShapeDeleteCallback GongOnAfterDeleteInterface[StateShape]

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

	OnAfterTransitionCreateCallback GongOnAfterCreateInterface[Transition]
	OnAfterTransitionUpdateCallback GongOnAfterUpdateInterface[Transition]
	OnAfterTransitionDeleteCallback GongOnAfterDeleteInterface[Transition]

	Transition_Shapes                map[*Transition_Shape]struct{}
	Transition_Shapes_instance       map[*Transition_Shape]*Transition_Shape
	Transition_Shapes_mapString      map[string]*Transition_Shape
	Transition_ShapeOrder            uint
	Transition_Shape_stagedOrder     map[*Transition_Shape]uint
	Transition_Shape_orderStaged     map[uint]*Transition_Shape
	Transition_Shapes_reference      map[*Transition_Shape]*Transition_Shape
	Transition_Shapes_referenceOrder map[*Transition_Shape]uint

	// insertion point for slice of pointers maps
	OnAfterTransition_ShapeCreateCallback GongOnAfterCreateInterface[Transition_Shape]
	OnAfterTransition_ShapeUpdateCallback GongOnAfterUpdateInterface[Transition_Shape]
	OnAfterTransition_ShapeDeleteCallback GongOnAfterDeleteInterface[Transition_Shape]

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
	__gong__clearReferences(&stage.Actions_reference, &stage.Actions_instance, &stage.Actions_referenceOrder)

	__gong__clearReferences(&stage.Activitiess_reference, &stage.Activitiess_instance, &stage.Activitiess_referenceOrder)

	__gong__clearReferences(&stage.Diagrams_reference, &stage.Diagrams_instance, &stage.Diagrams_referenceOrder)

	__gong__clearReferences(&stage.Guards_reference, &stage.Guards_instance, &stage.Guards_referenceOrder)

	__gong__clearReferences(&stage.Kills_reference, &stage.Kills_instance, &stage.Kills_referenceOrder)

	__gong__clearReferences(&stage.Librarys_reference, &stage.Librarys_instance, &stage.Librarys_referenceOrder)

	__gong__clearReferences(&stage.Messages_reference, &stage.Messages_instance, &stage.Messages_referenceOrder)

	__gong__clearReferences(&stage.MessageTypes_reference, &stage.MessageTypes_instance, &stage.MessageTypes_referenceOrder)

	__gong__clearReferences(&stage.Notes_reference, &stage.Notes_instance, &stage.Notes_referenceOrder)

	__gong__clearReferences(&stage.NoteShapes_reference, &stage.NoteShapes_instance, &stage.NoteShapes_referenceOrder)

	__gong__clearReferences(&stage.NoteStateShapes_reference, &stage.NoteStateShapes_instance, &stage.NoteStateShapes_referenceOrder)

	__gong__clearReferences(&stage.Objects_reference, &stage.Objects_instance, &stage.Objects_referenceOrder)

	__gong__clearReferences(&stage.Roles_reference, &stage.Roles_instance, &stage.Roles_referenceOrder)

	__gong__clearReferences(&stage.States_reference, &stage.States_instance, &stage.States_referenceOrder)

	__gong__clearReferences(&stage.StateMachines_reference, &stage.StateMachines_instance, &stage.StateMachines_referenceOrder)

	__gong__clearReferences(&stage.StateShapes_reference, &stage.StateShapes_instance, &stage.StateShapes_referenceOrder)

	__gong__clearReferences(&stage.Transitions_reference, &stage.Transitions_instance, &stage.Transitions_referenceOrder)

	__gong__clearReferences(&stage.Transition_Shapes_reference, &stage.Transition_Shapes_instance, &stage.Transition_Shapes_referenceOrder)

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
	stage.ActionOrder = __gong__recomputeOrder(stage.Action_stagedOrder)

	stage.ActivitiesOrder = __gong__recomputeOrder(stage.Activities_stagedOrder)

	stage.DiagramOrder = __gong__recomputeOrder(stage.Diagram_stagedOrder)

	stage.GuardOrder = __gong__recomputeOrder(stage.Guard_stagedOrder)

	stage.KillOrder = __gong__recomputeOrder(stage.Kill_stagedOrder)

	stage.LibraryOrder = __gong__recomputeOrder(stage.Library_stagedOrder)

	stage.MessageOrder = __gong__recomputeOrder(stage.Message_stagedOrder)

	stage.MessageTypeOrder = __gong__recomputeOrder(stage.MessageType_stagedOrder)

	stage.NoteOrder = __gong__recomputeOrder(stage.Note_stagedOrder)

	stage.NoteShapeOrder = __gong__recomputeOrder(stage.NoteShape_stagedOrder)

	stage.NoteStateShapeOrder = __gong__recomputeOrder(stage.NoteStateShape_stagedOrder)

	stage.ObjectOrder = __gong__recomputeOrder(stage.Object_stagedOrder)

	stage.RoleOrder = __gong__recomputeOrder(stage.Role_stagedOrder)

	stage.StateOrder = __gong__recomputeOrder(stage.State_stagedOrder)

	stage.StateMachineOrder = __gong__recomputeOrder(stage.StateMachine_stagedOrder)

	stage.StateShapeOrder = __gong__recomputeOrder(stage.StateShape_stagedOrder)

	stage.TransitionOrder = __gong__recomputeOrder(stage.Transition_stagedOrder)

	stage.Transition_ShapeOrder = __gong__recomputeOrder(stage.Transition_Shape_stagedOrder)

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
	case *Action:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Actions, stage.Action_stagedOrder))
	case *Activities:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Activitiess, stage.Activities_stagedOrder))
	case *Diagram:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Diagrams, stage.Diagram_stagedOrder))
	case *Guard:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Guards, stage.Guard_stagedOrder))
	case *Kill:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Kills, stage.Kill_stagedOrder))
	case *Library:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Librarys, stage.Library_stagedOrder))
	case *Message:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Messages, stage.Message_stagedOrder))
	case *MessageType:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.MessageTypes, stage.MessageType_stagedOrder))
	case *Note:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Notes, stage.Note_stagedOrder))
	case *NoteShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteShapes, stage.NoteShape_stagedOrder))
	case *NoteStateShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.NoteStateShapes, stage.NoteStateShape_stagedOrder))
	case *Object:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Objects, stage.Object_stagedOrder))
	case *Role:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Roles, stage.Role_stagedOrder))
	case *State:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.States, stage.State_stagedOrder))
	case *StateMachine:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.StateMachines, stage.StateMachine_stagedOrder))
	case *StateShape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.StateShapes, stage.StateShape_stagedOrder))
	case *Transition:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Transitions, stage.Transition_stagedOrder))
	case *Transition_Shape:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Transition_Shapes, stage.Transition_Shape_stagedOrder))

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
	return "github.com/fullstack-lang/gong/dsm/statemachines/go/models"
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
		Actions:           make(map[*Action]struct{}),
		Actions_mapString: make(map[string]*Action),

		Activitiess:           make(map[*Activities]struct{}),
		Activitiess_mapString: make(map[string]*Activities),

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
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Action": &ActionUnmarshaller{},

			"Activities": &ActivitiesUnmarshaller{},

			"Diagram": &DiagramUnmarshaller{},

			"Guard": &GuardUnmarshaller{},

			"Kill": &KillUnmarshaller{},

			"Library": &LibraryUnmarshaller{},

			"Message": &MessageUnmarshaller{},

			"MessageType": &MessageTypeUnmarshaller{},

			"Note": &NoteUnmarshaller{},

			"NoteShape": &NoteShapeUnmarshaller{},

			"NoteStateShape": &NoteStateShapeUnmarshaller{},

			"Object": &ObjectUnmarshaller{},

			"Role": &RoleUnmarshaller{},

			"State": &StateUnmarshaller{},

			"StateMachine": &StateMachineUnmarshaller{},

			"StateShape": &StateShapeUnmarshaller{},

			"Transition": &TransitionUnmarshaller{},

			"Transition_Shape": &Transition_ShapeUnmarshaller{},

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
	case *Action:
		return any(stage.Action_orderStaged[order]).(Type)
	case *Activities:
		return any(stage.Activities_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Guard"] = len(stage.Guards)
	stage.Map_GongStructName_InstancesNb["Kill"] = len(stage.Kills)
	stage.Map_GongStructName_InstancesNb["Library"] = len(stage.Librarys)
	stage.Map_GongStructName_InstancesNb["Message"] = len(stage.Messages)
	stage.Map_GongStructName_InstancesNb["MessageType"] = len(stage.MessageTypes)
	stage.Map_GongStructName_InstancesNb["Note"] = len(stage.Notes)
	stage.Map_GongStructName_InstancesNb["NoteShape"] = len(stage.NoteShapes)
	stage.Map_GongStructName_InstancesNb["NoteStateShape"] = len(stage.NoteStateShapes)
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
	__gong__stage(stage.Actions, stage.Action_stagedOrder, stage.Action_orderStaged, &stage.ActionOrder, stage.Actions_mapString, action, action.Name)
	return action
}

// StagePreserveOrder puts action to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActionOrder
// - update stage.ActionOrder accordingly
func (action *Action) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Actions, stage.Action_stagedOrder, stage.Action_orderStaged, &stage.ActionOrder, stage.Actions_mapString, action, order, action.Name)
}

// Unstage removes action off the model stage
func (action *Action) Unstage(stage *Stage) *Action {
	__gong__unstage(stage.Actions, stage.Actions_mapString, action, action.Name)
	return action
}

// UnstageVoid removes action off the model stage
func (action *Action) UnstageVoid(stage *Stage) {
	action.Unstage(stage)
}

func (action *Action) StageVoid(stage *Stage) {
	action.Stage(stage)
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
	__gong__stage(stage.Activitiess, stage.Activities_stagedOrder, stage.Activities_orderStaged, &stage.ActivitiesOrder, stage.Activitiess_mapString, activities, activities.Name)
	return activities
}

// StagePreserveOrder puts activities to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActivitiesOrder
// - update stage.ActivitiesOrder accordingly
func (activities *Activities) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Activitiess, stage.Activities_stagedOrder, stage.Activities_orderStaged, &stage.ActivitiesOrder, stage.Activitiess_mapString, activities, order, activities.Name)
}

// Unstage removes activities off the model stage
func (activities *Activities) Unstage(stage *Stage) *Activities {
	__gong__unstage(stage.Activitiess, stage.Activitiess_mapString, activities, activities.Name)
	return activities
}

// UnstageVoid removes activities off the model stage
func (activities *Activities) UnstageVoid(stage *Stage) {
	activities.Unstage(stage)
}

func (activities *Activities) StageVoid(stage *Stage) {
	activities.Stage(stage)
}

// for satisfaction of GongStruct interface
func (activities *Activities) GetName() (res string) {
	return activities.Name
}

// for satisfaction of GongStruct interface
func (activities *Activities) SetName(name string) {
	activities.Name = name
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

// Stage puts guard to the model stage
func (guard *Guard) Stage(stage *Stage) *Guard {
	__gong__stage(stage.Guards, stage.Guard_stagedOrder, stage.Guard_orderStaged, &stage.GuardOrder, stage.Guards_mapString, guard, guard.Name)
	return guard
}

// StagePreserveOrder puts guard to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GuardOrder
// - update stage.GuardOrder accordingly
func (guard *Guard) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Guards, stage.Guard_stagedOrder, stage.Guard_orderStaged, &stage.GuardOrder, stage.Guards_mapString, guard, order, guard.Name)
}

// Unstage removes guard off the model stage
func (guard *Guard) Unstage(stage *Stage) *Guard {
	__gong__unstage(stage.Guards, stage.Guards_mapString, guard, guard.Name)
	return guard
}

// UnstageVoid removes guard off the model stage
func (guard *Guard) UnstageVoid(stage *Stage) {
	guard.Unstage(stage)
}

func (guard *Guard) StageVoid(stage *Stage) {
	guard.Stage(stage)
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
	__gong__stage(stage.Kills, stage.Kill_stagedOrder, stage.Kill_orderStaged, &stage.KillOrder, stage.Kills_mapString, kill, kill.Name)
	return kill
}

// StagePreserveOrder puts kill to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.KillOrder
// - update stage.KillOrder accordingly
func (kill *Kill) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Kills, stage.Kill_stagedOrder, stage.Kill_orderStaged, &stage.KillOrder, stage.Kills_mapString, kill, order, kill.Name)
}

// Unstage removes kill off the model stage
func (kill *Kill) Unstage(stage *Stage) *Kill {
	__gong__unstage(stage.Kills, stage.Kills_mapString, kill, kill.Name)
	return kill
}

// UnstageVoid removes kill off the model stage
func (kill *Kill) UnstageVoid(stage *Stage) {
	kill.Unstage(stage)
}

func (kill *Kill) StageVoid(stage *Stage) {
	kill.Stage(stage)
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

// Stage puts message to the model stage
func (message *Message) Stage(stage *Stage) *Message {
	__gong__stage(stage.Messages, stage.Message_stagedOrder, stage.Message_orderStaged, &stage.MessageOrder, stage.Messages_mapString, message, message.Name)
	return message
}

// StagePreserveOrder puts message to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MessageOrder
// - update stage.MessageOrder accordingly
func (message *Message) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Messages, stage.Message_stagedOrder, stage.Message_orderStaged, &stage.MessageOrder, stage.Messages_mapString, message, order, message.Name)
}

// Unstage removes message off the model stage
func (message *Message) Unstage(stage *Stage) *Message {
	__gong__unstage(stage.Messages, stage.Messages_mapString, message, message.Name)
	return message
}

// UnstageVoid removes message off the model stage
func (message *Message) UnstageVoid(stage *Stage) {
	message.Unstage(stage)
}

func (message *Message) StageVoid(stage *Stage) {
	message.Stage(stage)
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
	__gong__stage(stage.MessageTypes, stage.MessageType_stagedOrder, stage.MessageType_orderStaged, &stage.MessageTypeOrder, stage.MessageTypes_mapString, messagetype, messagetype.Name)
	return messagetype
}

// StagePreserveOrder puts messagetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MessageTypeOrder
// - update stage.MessageTypeOrder accordingly
func (messagetype *MessageType) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.MessageTypes, stage.MessageType_stagedOrder, stage.MessageType_orderStaged, &stage.MessageTypeOrder, stage.MessageTypes_mapString, messagetype, order, messagetype.Name)
}

// Unstage removes messagetype off the model stage
func (messagetype *MessageType) Unstage(stage *Stage) *MessageType {
	__gong__unstage(stage.MessageTypes, stage.MessageTypes_mapString, messagetype, messagetype.Name)
	return messagetype
}

// UnstageVoid removes messagetype off the model stage
func (messagetype *MessageType) UnstageVoid(stage *Stage) {
	messagetype.Unstage(stage)
}

func (messagetype *MessageType) StageVoid(stage *Stage) {
	messagetype.Stage(stage)
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

// Stage puts notestateshape to the model stage
func (notestateshape *NoteStateShape) Stage(stage *Stage) *NoteStateShape {
	__gong__stage(stage.NoteStateShapes, stage.NoteStateShape_stagedOrder, stage.NoteStateShape_orderStaged, &stage.NoteStateShapeOrder, stage.NoteStateShapes_mapString, notestateshape, notestateshape.Name)
	return notestateshape
}

// StagePreserveOrder puts notestateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.NoteStateShapeOrder
// - update stage.NoteStateShapeOrder accordingly
func (notestateshape *NoteStateShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.NoteStateShapes, stage.NoteStateShape_stagedOrder, stage.NoteStateShape_orderStaged, &stage.NoteStateShapeOrder, stage.NoteStateShapes_mapString, notestateshape, order, notestateshape.Name)
}

// Unstage removes notestateshape off the model stage
func (notestateshape *NoteStateShape) Unstage(stage *Stage) *NoteStateShape {
	__gong__unstage(stage.NoteStateShapes, stage.NoteStateShapes_mapString, notestateshape, notestateshape.Name)
	return notestateshape
}

// UnstageVoid removes notestateshape off the model stage
func (notestateshape *NoteStateShape) UnstageVoid(stage *Stage) {
	notestateshape.Unstage(stage)
}

func (notestateshape *NoteStateShape) StageVoid(stage *Stage) {
	notestateshape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (notestateshape *NoteStateShape) GetName() (res string) {
	return notestateshape.Name
}

// for satisfaction of GongStruct interface
func (notestateshape *NoteStateShape) SetName(name string) {
	notestateshape.Name = name
}

// Stage puts object to the model stage
func (object *Object) Stage(stage *Stage) *Object {
	__gong__stage(stage.Objects, stage.Object_stagedOrder, stage.Object_orderStaged, &stage.ObjectOrder, stage.Objects_mapString, object, object.Name)
	return object
}

// StagePreserveOrder puts object to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ObjectOrder
// - update stage.ObjectOrder accordingly
func (object *Object) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Objects, stage.Object_stagedOrder, stage.Object_orderStaged, &stage.ObjectOrder, stage.Objects_mapString, object, order, object.Name)
}

// Unstage removes object off the model stage
func (object *Object) Unstage(stage *Stage) *Object {
	__gong__unstage(stage.Objects, stage.Objects_mapString, object, object.Name)
	return object
}

// UnstageVoid removes object off the model stage
func (object *Object) UnstageVoid(stage *Stage) {
	object.Unstage(stage)
}

func (object *Object) StageVoid(stage *Stage) {
	object.Stage(stage)
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
	__gong__stage(stage.Roles, stage.Role_stagedOrder, stage.Role_orderStaged, &stage.RoleOrder, stage.Roles_mapString, role, role.Name)
	return role
}

// StagePreserveOrder puts role to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RoleOrder
// - update stage.RoleOrder accordingly
func (role *Role) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Roles, stage.Role_stagedOrder, stage.Role_orderStaged, &stage.RoleOrder, stage.Roles_mapString, role, order, role.Name)
}

// Unstage removes role off the model stage
func (role *Role) Unstage(stage *Stage) *Role {
	__gong__unstage(stage.Roles, stage.Roles_mapString, role, role.Name)
	return role
}

// UnstageVoid removes role off the model stage
func (role *Role) UnstageVoid(stage *Stage) {
	role.Unstage(stage)
}

func (role *Role) StageVoid(stage *Stage) {
	role.Stage(stage)
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
	__gong__stage(stage.States, stage.State_stagedOrder, stage.State_orderStaged, &stage.StateOrder, stage.States_mapString, state, state.Name)
	return state
}

// StagePreserveOrder puts state to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateOrder
// - update stage.StateOrder accordingly
func (state *State) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.States, stage.State_stagedOrder, stage.State_orderStaged, &stage.StateOrder, stage.States_mapString, state, order, state.Name)
}

// Unstage removes state off the model stage
func (state *State) Unstage(stage *Stage) *State {
	__gong__unstage(stage.States, stage.States_mapString, state, state.Name)
	return state
}

// UnstageVoid removes state off the model stage
func (state *State) UnstageVoid(stage *Stage) {
	state.Unstage(stage)
}

func (state *State) StageVoid(stage *Stage) {
	state.Stage(stage)
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
	__gong__stage(stage.StateMachines, stage.StateMachine_stagedOrder, stage.StateMachine_orderStaged, &stage.StateMachineOrder, stage.StateMachines_mapString, statemachine, statemachine.Name)
	return statemachine
}

// StagePreserveOrder puts statemachine to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateMachineOrder
// - update stage.StateMachineOrder accordingly
func (statemachine *StateMachine) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StateMachines, stage.StateMachine_stagedOrder, stage.StateMachine_orderStaged, &stage.StateMachineOrder, stage.StateMachines_mapString, statemachine, order, statemachine.Name)
}

// Unstage removes statemachine off the model stage
func (statemachine *StateMachine) Unstage(stage *Stage) *StateMachine {
	__gong__unstage(stage.StateMachines, stage.StateMachines_mapString, statemachine, statemachine.Name)
	return statemachine
}

// UnstageVoid removes statemachine off the model stage
func (statemachine *StateMachine) UnstageVoid(stage *Stage) {
	statemachine.Unstage(stage)
}

func (statemachine *StateMachine) StageVoid(stage *Stage) {
	statemachine.Stage(stage)
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
	__gong__stage(stage.StateShapes, stage.StateShape_stagedOrder, stage.StateShape_orderStaged, &stage.StateShapeOrder, stage.StateShapes_mapString, stateshape, stateshape.Name)
	return stateshape
}

// StagePreserveOrder puts stateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateShapeOrder
// - update stage.StateShapeOrder accordingly
func (stateshape *StateShape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.StateShapes, stage.StateShape_stagedOrder, stage.StateShape_orderStaged, &stage.StateShapeOrder, stage.StateShapes_mapString, stateshape, order, stateshape.Name)
}

// Unstage removes stateshape off the model stage
func (stateshape *StateShape) Unstage(stage *Stage) *StateShape {
	__gong__unstage(stage.StateShapes, stage.StateShapes_mapString, stateshape, stateshape.Name)
	return stateshape
}

// UnstageVoid removes stateshape off the model stage
func (stateshape *StateShape) UnstageVoid(stage *Stage) {
	stateshape.Unstage(stage)
}

func (stateshape *StateShape) StageVoid(stage *Stage) {
	stateshape.Stage(stage)
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
	__gong__stage(stage.Transitions, stage.Transition_stagedOrder, stage.Transition_orderStaged, &stage.TransitionOrder, stage.Transitions_mapString, transition, transition.Name)
	return transition
}

// StagePreserveOrder puts transition to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TransitionOrder
// - update stage.TransitionOrder accordingly
func (transition *Transition) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Transitions, stage.Transition_stagedOrder, stage.Transition_orderStaged, &stage.TransitionOrder, stage.Transitions_mapString, transition, order, transition.Name)
}

// Unstage removes transition off the model stage
func (transition *Transition) Unstage(stage *Stage) *Transition {
	__gong__unstage(stage.Transitions, stage.Transitions_mapString, transition, transition.Name)
	return transition
}

// UnstageVoid removes transition off the model stage
func (transition *Transition) UnstageVoid(stage *Stage) {
	transition.Unstage(stage)
}

func (transition *Transition) StageVoid(stage *Stage) {
	transition.Stage(stage)
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
	__gong__stage(stage.Transition_Shapes, stage.Transition_Shape_stagedOrder, stage.Transition_Shape_orderStaged, &stage.Transition_ShapeOrder, stage.Transition_Shapes_mapString, transition_shape, transition_shape.Name)
	return transition_shape
}

// StagePreserveOrder puts transition_shape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Transition_ShapeOrder
// - update stage.Transition_ShapeOrder accordingly
func (transition_shape *Transition_Shape) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Transition_Shapes, stage.Transition_Shape_stagedOrder, stage.Transition_Shape_orderStaged, &stage.Transition_ShapeOrder, stage.Transition_Shapes_mapString, transition_shape, order, transition_shape.Name)
}

// Unstage removes transition_shape off the model stage
func (transition_shape *Transition_Shape) Unstage(stage *Stage) *Transition_Shape {
	__gong__unstage(stage.Transition_Shapes, stage.Transition_Shapes_mapString, transition_shape, transition_shape.Name)
	return transition_shape
}

// UnstageVoid removes transition_shape off the model stage
func (transition_shape *Transition_Shape) UnstageVoid(stage *Stage) {
	transition_shape.Unstage(stage)
}

func (transition_shape *Transition_Shape) StageVoid(stage *Stage) {
	transition_shape.Stage(stage)
}

// for satisfaction of GongStruct interface
func (transition_shape *Transition_Shape) GetName() (res string) {
	return transition_shape.Name
}

// for satisfaction of GongStruct interface
func (transition_shape *Transition_Shape) SetName(name string) {
	transition_shape.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Actions, &stage.Actions_mapString, &stage.Action_stagedOrder, &stage.ActionOrder)

	__gong__resetStageType(&stage.Activitiess, &stage.Activitiess_mapString, &stage.Activities_stagedOrder, &stage.ActivitiesOrder)

	__gong__resetStageType(&stage.Diagrams, &stage.Diagrams_mapString, &stage.Diagram_stagedOrder, &stage.DiagramOrder)

	__gong__resetStageType(&stage.Guards, &stage.Guards_mapString, &stage.Guard_stagedOrder, &stage.GuardOrder)

	__gong__resetStageType(&stage.Kills, &stage.Kills_mapString, &stage.Kill_stagedOrder, &stage.KillOrder)

	__gong__resetStageType(&stage.Librarys, &stage.Librarys_mapString, &stage.Library_stagedOrder, &stage.LibraryOrder)

	__gong__resetStageType(&stage.Messages, &stage.Messages_mapString, &stage.Message_stagedOrder, &stage.MessageOrder)

	__gong__resetStageType(&stage.MessageTypes, &stage.MessageTypes_mapString, &stage.MessageType_stagedOrder, &stage.MessageTypeOrder)

	__gong__resetStageType(&stage.Notes, &stage.Notes_mapString, &stage.Note_stagedOrder, &stage.NoteOrder)

	__gong__resetStageType(&stage.NoteShapes, &stage.NoteShapes_mapString, &stage.NoteShape_stagedOrder, &stage.NoteShapeOrder)

	__gong__resetStageType(&stage.NoteStateShapes, &stage.NoteStateShapes_mapString, &stage.NoteStateShape_stagedOrder, &stage.NoteStateShapeOrder)

	__gong__resetStageType(&stage.Objects, &stage.Objects_mapString, &stage.Object_stagedOrder, &stage.ObjectOrder)

	__gong__resetStageType(&stage.Roles, &stage.Roles_mapString, &stage.Role_stagedOrder, &stage.RoleOrder)

	__gong__resetStageType(&stage.States, &stage.States_mapString, &stage.State_stagedOrder, &stage.StateOrder)

	__gong__resetStageType(&stage.StateMachines, &stage.StateMachines_mapString, &stage.StateMachine_stagedOrder, &stage.StateMachineOrder)

	__gong__resetStageType(&stage.StateShapes, &stage.StateShapes_mapString, &stage.StateShape_stagedOrder, &stage.StateShapeOrder)

	__gong__resetStageType(&stage.Transitions, &stage.Transitions_mapString, &stage.Transition_stagedOrder, &stage.TransitionOrder)

	__gong__resetStageType(&stage.Transition_Shapes, &stage.Transition_Shapes_mapString, &stage.Transition_Shape_stagedOrder, &stage.Transition_ShapeOrder)

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
	case *Action:
		return any(stage.Actions_mapString).(map[string]Type)
	case *Activities:
		return any(stage.Activitiess_mapString).(map[string]Type)
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

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Action:
		return any(&stage.Actions).(*map[Type]struct{})
	case *Activities:
		return any(&stage.Activitiess).(*map[Type]struct{})
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
			State_Shapes: []*StateShape{{Name: "State_Shapes"}},
			StatesWhoseNodeIsExpanded: []*State{{Name: "StatesWhoseNodeIsExpanded"}},
			Transition_Shapes: []*Transition_Shape{{Name: "Transition_Shapes"}},
			Note_Shapes: []*NoteShape{{Name: "Note_Shapes"}},
			NoteState_Shapes: []*NoteStateShape{{Name: "NoteState_Shapes"}},
		}).(*Type)
	case Library:
		return any(&Library{
			SubLibraries: []*Library{{Name: "SubLibraries"}},
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			RootStateMachines: []*StateMachine{{Name: "RootStateMachines"}},
			StateMachinesWhoseNodeIsExpanded: []*StateMachine{{Name: "StateMachinesWhoseNodeIsExpanded"}},
			SubLibrariesWhoseNodeIsExpanded: []*Library{{Name: "SubLibrariesWhoseNodeIsExpanded"}},
			Roles: []*Role{{Name: "Roles"}},
			MessageTypes: []*MessageType{{Name: "MessageTypes"}},
		}).(*Type)
	case Message:
		return any(&Message{
			MessageType: &MessageType{Name: "MessageType"},
			OriginTransition: &Transition{Name: "OriginTransition"},
		}).(*Type)
	case Note:
		return any(&Note{
			State: &State{Name: "State"},
		}).(*Type)
	case NoteShape:
		return any(&NoteShape{
			Note: &Note{Name: "Note"},
		}).(*Type)
	case NoteStateShape:
		return any(&NoteStateShape{
			Note: &Note{Name: "Note"},
			State: &State{Name: "State"},
		}).(*Type)
	case Object:
		return any(&Object{
			State: &State{Name: "State"},
			Messages: []*Message{{Name: "Messages"}},
		}).(*Type)
	case Role:
		return any(&Role{
			RolesWithSamePermissions: []*Role{{Name: "RolesWithSamePermissions"}},
		}).(*Type)
	case State:
		return any(&State{
			SubStates: []*State{{Name: "SubStates"}},
			Entry: &Action{Name: "Entry"},
			Activities: []*Activities{{Name: "Activities"}},
			Exit: &Action{Name: "Exit"},
			Parent: &State{Name: "Parent"},
			Diagrams: []*Diagram{{Name: "Diagrams"}},
			Notes: []*Note{{Name: "Notes"}},
		}).(*Type)
	case StateMachine:
		return any(&StateMachine{
			InitialState: &State{Name: "InitialState"},
			States: []*State{{Name: "States"}},
			Diagrams: []*Diagram{{Name: "Diagrams"}},
		}).(*Type)
	case StateShape:
		return any(&StateShape{
			State: &State{Name: "State"},
		}).(*Type)
	case Transition:
		return any(&Transition{
			Start: &State{Name: "Start"},
			End: &State{Name: "End"},
			RolesWithPermissions: []*Role{{Name: "RolesWithPermissions"}},
			GeneratedMessages: []*MessageType{{Name: "GeneratedMessages"}},
			Guard: &Guard{Name: "Guard"},
			Diagrams: []*Diagram{{Name: "Diagrams"}},
		}).(*Type)
	case Transition_Shape:
		return any(&Transition_Shape{
			Transition: &Transition{Name: "Transition"},
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
		case "State":
			res := make(map[*State][]*Note)
			for note := range stage.Notes {
				if note.State != nil {
					state_ := note.State
					var notes []*Note
					_, ok := res[state_]
					if ok {
						notes = res[state_]
					} else {
						notes = make([]*Note, 0)
					}
					notes = append(notes, note)
					res[state_] = notes
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

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
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
	// reverse maps of direct associations of Diagram
	case Diagram:
		switch fieldname {
		// insertion point for per direct association field
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
		case "SubLibrariesWhoseNodeIsExpanded":
			res := make(map[*Library][]*Library)
			for library := range stage.Librarys {
				for _, library_ := range library.SubLibrariesWhoseNodeIsExpanded {
					res[library_] = append(res[library_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Roles":
			res := make(map[*Role][]*Library)
			for library := range stage.Librarys {
				for _, role_ := range library.Roles {
					res[role_] = append(res[role_], library)
				}
			}
			return any(res).(map[*End][]*Start)
		case "MessageTypes":
			res := make(map[*MessageType][]*Library)
			for library := range stage.Librarys {
				for _, messagetype_ := range library.MessageTypes {
					res[messagetype_] = append(res[messagetype_], library)
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
		case "Activities":
			res := make(map[*Activities][]*State)
			for state := range stage.States {
				for _, activities_ := range state.Activities {
					res[activities_] = append(res[activities_], state)
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
		case "Notes":
			res := make(map[*Note][]*State)
			for state := range stage.States {
				for _, note_ := range state.Notes {
					res[note_] = append(res[note_], state)
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

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *Action:
		res = any(new(Action)).(Type)
	case *Activities:
		res = any(new(Activities)).(Type)
	case *Diagram:
		res = any(new(Diagram)).(Type)
	case *Guard:
		res = any(new(Guard)).(Type)
	case *Kill:
		res = any(new(Kill)).(Type)
	case *Library:
		res = any(new(Library)).(Type)
	case *Message:
		res = any(new(Message)).(Type)
	case *MessageType:
		res = any(new(MessageType)).(Type)
	case *Note:
		res = any(new(Note)).(Type)
	case *NoteShape:
		res = any(new(NoteShape)).(Type)
	case *NoteStateShape:
		res = any(new(NoteStateShape)).(Type)
	case *Object:
		res = any(new(Object)).(Type)
	case *Role:
		res = any(new(Role)).(Type)
	case *State:
		res = any(new(State)).(Type)
	case *StateMachine:
		res = any(new(StateMachine)).(Type)
	case *StateShape:
		res = any(new(StateShape)).(Type)
	case *Transition:
		res = any(new(Transition)).(Type)
	case *Transition_Shape:
		res = any(new(Transition_Shape)).(Type)
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
	case *Action:
		res = "Action"
	case *Activities:
		res = "Activities"
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
	case *Action:
		var rf ReverseField
		_ = rf
	case *Activities:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "State"
		rf.Fieldname = "Activities"
		res = append(res, rf)
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
		rf.GongstructName = "Library"
		rf.Fieldname = "MessageTypes"
		res = append(res, rf)
		rf.GongstructName = "Transition"
		rf.Fieldname = "GeneratedMessages"
		res = append(res, rf)
	case *Note:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "State"
		rf.Fieldname = "Notes"
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
	case *Object:
		var rf ReverseField
		_ = rf
	case *Role:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Library"
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
		rf.GongstructName = "State"
		rf.Fieldname = "SubStates"
		res = append(res, rf)
		rf.GongstructName = "StateMachine"
		rf.Fieldname = "States"
		res = append(res, rf)
	case *StateMachine:
		var rf ReverseField
		_ = rf
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
	case *Transition_Shape:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Diagram"
		rf.Fieldname = "Transition_Shapes"
		res = append(res, rf)
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
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
			Name:               "ShowRoles",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ShowMessages",
			GongFieldValueType: GongFieldValueTypeBool,
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
		{
			Name:                 "Roles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Role",
		},
		{
			Name:               "IsRolesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "MessageTypes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "MessageType",
		},
		{
			Name:               "IsMessageTypesNodeExpanded",
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
			Name:                 "State",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
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
			Name:               "IsEndState",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsDecisionNode",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "SubStates",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "State",
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
		{
			Name:                 "Parent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
		{
			Name:               "IsFictious",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:                 "Diagrams",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Diagram",
		},
		{
			Name:                 "Notes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Note",
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
			Name:                 "InitialState",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
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
			Name:               "IsWithTransitionNameAutonamticalyGenerated",
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
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsRolesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsMessagesNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBool,
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
	case "ShowRoles":
		res.valueString = fmt.Sprintf("%t", diagram.ShowRoles)
		res.valueBool = diagram.ShowRoles
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ShowMessages":
		res.valueString = fmt.Sprintf("%t", diagram.ShowMessages)
		res.valueBool = diagram.ShowMessages
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "Roles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.Roles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsRolesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsRolesNodeExpanded)
		res.valueBool = library.IsRolesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MessageTypes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range library.MessageTypes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsMessageTypesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", library.IsMessageTypesNodeExpanded)
		res.valueBool = library.IsMessageTypesNodeExpanded
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
	case "State":
		res.GongFieldValueType = GongFieldValueTypePointer
		if note.State != nil {
			res.valueString = note.State.Name
			res.ids = note.State.GongGetUUID(stage)
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
	case "IsEndState":
		res.valueString = fmt.Sprintf("%t", state.IsEndState)
		res.valueBool = state.IsEndState
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsDecisionNode":
		res.valueString = fmt.Sprintf("%t", state.IsDecisionNode)
		res.valueBool = state.IsDecisionNode
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
	case "Parent":
		res.GongFieldValueType = GongFieldValueTypePointer
		if state.Parent != nil {
			res.valueString = state.Parent.Name
			res.ids = state.Parent.GongGetUUID(stage)
		}
	case "IsFictious":
		res.valueString = fmt.Sprintf("%t", state.IsFictious)
		res.valueBool = state.IsFictious
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "Notes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range state.Notes {
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

func (statemachine *StateMachine) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = statemachine.Name
	case "InitialState":
		res.GongFieldValueType = GongFieldValueTypePointer
		if statemachine.InitialState != nil {
			res.valueString = statemachine.InitialState.Name
			res.ids = statemachine.InitialState.GongGetUUID(stage)
		}
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
	case "IsWithTransitionNameAutonamticalyGenerated":
		res.valueString = fmt.Sprintf("%t", statemachine.IsWithTransitionNameAutonamticalyGenerated)
		res.valueBool = statemachine.IsWithTransitionNameAutonamticalyGenerated
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ComputedPrefix":
		res.valueString = statemachine.ComputedPrefix
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", statemachine.IsExpanded)
		res.valueBool = statemachine.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", transition.IsExpanded)
		res.valueBool = transition.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsRolesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", transition.IsRolesNodeExpanded)
		res.valueBool = transition.IsRolesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsMessagesNodeExpanded":
		res.valueString = fmt.Sprintf("%t", transition.IsMessagesNodeExpanded)
		res.valueBool = transition.IsMessagesNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (action *Action) GongGetGongstructName() string {
	return "Action"
}

func (activities *Activities) GongGetGongstructName() string {
	return "Activities"
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

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	__gong__rebuildMapString(stage.Actions, &stage.Actions_mapString)

	__gong__rebuildMapString(stage.Activitiess, &stage.Activitiess_mapString)

	__gong__rebuildMapString(stage.Diagrams, &stage.Diagrams_mapString)

	__gong__rebuildMapString(stage.Guards, &stage.Guards_mapString)

	__gong__rebuildMapString(stage.Kills, &stage.Kills_mapString)

	__gong__rebuildMapString(stage.Librarys, &stage.Librarys_mapString)

	__gong__rebuildMapString(stage.Messages, &stage.Messages_mapString)

	__gong__rebuildMapString(stage.MessageTypes, &stage.MessageTypes_mapString)

	__gong__rebuildMapString(stage.Notes, &stage.Notes_mapString)

	__gong__rebuildMapString(stage.NoteShapes, &stage.NoteShapes_mapString)

	__gong__rebuildMapString(stage.NoteStateShapes, &stage.NoteStateShapes_mapString)

	__gong__rebuildMapString(stage.Objects, &stage.Objects_mapString)

	__gong__rebuildMapString(stage.Roles, &stage.Roles_mapString)

	__gong__rebuildMapString(stage.States, &stage.States_mapString)

	__gong__rebuildMapString(stage.StateMachines, &stage.StateMachines_mapString)

	__gong__rebuildMapString(stage.StateShapes, &stage.StateShapes_mapString)

	__gong__rebuildMapString(stage.Transitions, &stage.Transitions_mapString)

	__gong__rebuildMapString(stage.Transition_Shapes, &stage.Transition_Shapes_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
