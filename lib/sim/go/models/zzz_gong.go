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
	Commands                map[*Command]struct{}
	Commands_instance       map[*Command]*Command
	Commands_mapString      map[string]*Command
	CommandOrder            uint
	Command_stagedOrder     map[*Command]uint
	Command_orderStaged     map[uint]*Command
	Commands_reference      map[*Command]*Command
	Commands_referenceOrder map[*Command]uint

	// insertion point for slice of pointers maps
	OnAfterCommandCreateCallback GongOnAfterCreateInterface[Command]
	OnAfterCommandUpdateCallback GongOnAfterUpdateInterface[Command]
	OnAfterCommandDeleteCallback GongOnAfterDeleteInterface[Command]

	DummyAgents                map[*DummyAgent]struct{}
	DummyAgents_instance       map[*DummyAgent]*DummyAgent
	DummyAgents_mapString      map[string]*DummyAgent
	DummyAgentOrder            uint
	DummyAgent_stagedOrder     map[*DummyAgent]uint
	DummyAgent_orderStaged     map[uint]*DummyAgent
	DummyAgents_reference      map[*DummyAgent]*DummyAgent
	DummyAgents_referenceOrder map[*DummyAgent]uint

	// insertion point for slice of pointers maps
	OnAfterDummyAgentCreateCallback GongOnAfterCreateInterface[DummyAgent]
	OnAfterDummyAgentUpdateCallback GongOnAfterUpdateInterface[DummyAgent]
	OnAfterDummyAgentDeleteCallback GongOnAfterDeleteInterface[DummyAgent]

	Engines                map[*Engine]struct{}
	Engines_instance       map[*Engine]*Engine
	Engines_mapString      map[string]*Engine
	EngineOrder            uint
	Engine_stagedOrder     map[*Engine]uint
	Engine_orderStaged     map[uint]*Engine
	Engines_reference      map[*Engine]*Engine
	Engines_referenceOrder map[*Engine]uint

	// insertion point for slice of pointers maps
	OnAfterEngineCreateCallback GongOnAfterCreateInterface[Engine]
	OnAfterEngineUpdateCallback GongOnAfterUpdateInterface[Engine]
	OnAfterEngineDeleteCallback GongOnAfterDeleteInterface[Engine]

	Events                map[*Event]struct{}
	Events_instance       map[*Event]*Event
	Events_mapString      map[string]*Event
	EventOrder            uint
	Event_stagedOrder     map[*Event]uint
	Event_orderStaged     map[uint]*Event
	Events_reference      map[*Event]*Event
	Events_referenceOrder map[*Event]uint

	// insertion point for slice of pointers maps
	OnAfterEventCreateCallback GongOnAfterCreateInterface[Event]
	OnAfterEventUpdateCallback GongOnAfterUpdateInterface[Event]
	OnAfterEventDeleteCallback GongOnAfterDeleteInterface[Event]

	Statuss                map[*Status]struct{}
	Statuss_instance       map[*Status]*Status
	Statuss_mapString      map[string]*Status
	StatusOrder            uint
	Status_stagedOrder     map[*Status]uint
	Status_orderStaged     map[uint]*Status
	Statuss_reference      map[*Status]*Status
	Statuss_referenceOrder map[*Status]uint

	// insertion point for slice of pointers maps
	OnAfterStatusCreateCallback GongOnAfterCreateInterface[Status]
	OnAfterStatusUpdateCallback GongOnAfterUpdateInterface[Status]
	OnAfterStatusDeleteCallback GongOnAfterDeleteInterface[Status]

	UpdateStates                map[*UpdateState]struct{}
	UpdateStates_instance       map[*UpdateState]*UpdateState
	UpdateStates_mapString      map[string]*UpdateState
	UpdateStateOrder            uint
	UpdateState_stagedOrder     map[*UpdateState]uint
	UpdateState_orderStaged     map[uint]*UpdateState
	UpdateStates_reference      map[*UpdateState]*UpdateState
	UpdateStates_referenceOrder map[*UpdateState]uint

	// insertion point for slice of pointers maps
	OnAfterUpdateStateCreateCallback GongOnAfterCreateInterface[UpdateState]
	OnAfterUpdateStateUpdateCallback GongOnAfterUpdateInterface[UpdateState]
	OnAfterUpdateStateDeleteCallback GongOnAfterDeleteInterface[UpdateState]

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
	__gong__clearReferences(&stage.Commands_reference, &stage.Commands_instance, &stage.Commands_referenceOrder)

	__gong__clearReferences(&stage.DummyAgents_reference, &stage.DummyAgents_instance, &stage.DummyAgents_referenceOrder)

	__gong__clearReferences(&stage.Engines_reference, &stage.Engines_instance, &stage.Engines_referenceOrder)

	__gong__clearReferences(&stage.Events_reference, &stage.Events_instance, &stage.Events_referenceOrder)

	__gong__clearReferences(&stage.Statuss_reference, &stage.Statuss_instance, &stage.Statuss_referenceOrder)

	__gong__clearReferences(&stage.UpdateStates_reference, &stage.UpdateStates_instance, &stage.UpdateStates_referenceOrder)

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
	stage.CommandOrder = __gong__recomputeOrder(stage.Command_stagedOrder)

	stage.DummyAgentOrder = __gong__recomputeOrder(stage.DummyAgent_stagedOrder)

	stage.EngineOrder = __gong__recomputeOrder(stage.Engine_stagedOrder)

	stage.EventOrder = __gong__recomputeOrder(stage.Event_stagedOrder)

	stage.StatusOrder = __gong__recomputeOrder(stage.Status_stagedOrder)

	stage.UpdateStateOrder = __gong__recomputeOrder(stage.UpdateState_stagedOrder)

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
	case *Command:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Commands, stage.Command_stagedOrder))
	case *DummyAgent:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DummyAgents, stage.DummyAgent_stagedOrder))
	case *Engine:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Engines, stage.Engine_stagedOrder))
	case *Event:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Events, stage.Event_stagedOrder))
	case *Status:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Statuss, stage.Status_stagedOrder))
	case *UpdateState:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.UpdateStates, stage.UpdateState_stagedOrder))

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
	return "github.com/fullstack-lang/gong/lib/sim/go/models"
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
		Commands:           make(map[*Command]struct{}),
		Commands_mapString: make(map[string]*Command),

		DummyAgents:           make(map[*DummyAgent]struct{}),
		DummyAgents_mapString: make(map[string]*DummyAgent),

		Engines:           make(map[*Engine]struct{}),
		Engines_mapString: make(map[string]*Engine),

		Events:           make(map[*Event]struct{}),
		Events_mapString: make(map[string]*Event),

		Statuss:           make(map[*Status]struct{}),
		Statuss_mapString: make(map[string]*Status),

		UpdateStates:           make(map[*UpdateState]struct{}),
		UpdateStates_mapString: make(map[string]*UpdateState),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Command_stagedOrder: make(map[*Command]uint),
		Command_orderStaged: make(map[uint]*Command),
		Commands_reference:  make(map[*Command]*Command),

		DummyAgent_stagedOrder: make(map[*DummyAgent]uint),
		DummyAgent_orderStaged: make(map[uint]*DummyAgent),
		DummyAgents_reference:  make(map[*DummyAgent]*DummyAgent),

		Engine_stagedOrder: make(map[*Engine]uint),
		Engine_orderStaged: make(map[uint]*Engine),
		Engines_reference:  make(map[*Engine]*Engine),

		Event_stagedOrder: make(map[*Event]uint),
		Event_orderStaged: make(map[uint]*Event),
		Events_reference:  make(map[*Event]*Event),

		Status_stagedOrder: make(map[*Status]uint),
		Status_orderStaged: make(map[uint]*Status),
		Statuss_reference:  make(map[*Status]*Status),

		UpdateState_stagedOrder: make(map[*UpdateState]uint),
		UpdateState_orderStaged: make(map[uint]*UpdateState),
		UpdateStates_reference:  make(map[*UpdateState]*UpdateState),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Command": &CommandUnmarshaller{},

			"DummyAgent": &DummyAgentUnmarshaller{},

			"Engine": &EngineUnmarshaller{},

			"Event": &EventUnmarshaller{},

			"Status": &StatusUnmarshaller{},

			"UpdateState": &UpdateStateUnmarshaller{},

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
	case *Command:
		return any(stage.Command_orderStaged[order]).(Type)
	case *DummyAgent:
		return any(stage.DummyAgent_orderStaged[order]).(Type)
	case *Engine:
		return any(stage.Engine_orderStaged[order]).(Type)
	case *Event:
		return any(stage.Event_orderStaged[order]).(Type)
	case *Status:
		return any(stage.Status_orderStaged[order]).(Type)
	case *UpdateState:
		return any(stage.UpdateState_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Command"] = len(stage.Commands)
	stage.Map_GongStructName_InstancesNb["DummyAgent"] = len(stage.DummyAgents)
	stage.Map_GongStructName_InstancesNb["Engine"] = len(stage.Engines)
	stage.Map_GongStructName_InstancesNb["Event"] = len(stage.Events)
	stage.Map_GongStructName_InstancesNb["Status"] = len(stage.Statuss)
	stage.Map_GongStructName_InstancesNb["UpdateState"] = len(stage.UpdateStates)
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
// Stage puts command to the model stage
func (command *Command) Stage(stage *Stage) *Command {
	__gong__stage(stage.Commands, stage.Command_stagedOrder, stage.Command_orderStaged, &stage.CommandOrder, stage.Commands_mapString, command, command.Name)
	return command
}

// StagePreserveOrder puts command to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CommandOrder
// - update stage.CommandOrder accordingly
func (command *Command) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Commands, stage.Command_stagedOrder, stage.Command_orderStaged, &stage.CommandOrder, stage.Commands_mapString, command, order, command.Name)
}

// Unstage removes command off the model stage
func (command *Command) Unstage(stage *Stage) *Command {
	__gong__unstage(stage.Commands, stage.Commands_mapString, command, command.Name)
	return command
}

// UnstageVoid removes command off the model stage
func (command *Command) UnstageVoid(stage *Stage) {
	command.Unstage(stage)
}

func (command *Command) StageVoid(stage *Stage) {
	command.Stage(stage)
}

// for satisfaction of GongStruct interface
func (command *Command) GetName() (res string) {
	return command.Name
}

// for satisfaction of GongStruct interface
func (command *Command) SetName(name string) {
	command.Name = name
}

// Stage puts dummyagent to the model stage
func (dummyagent *DummyAgent) Stage(stage *Stage) *DummyAgent {
	__gong__stage(stage.DummyAgents, stage.DummyAgent_stagedOrder, stage.DummyAgent_orderStaged, &stage.DummyAgentOrder, stage.DummyAgents_mapString, dummyagent, dummyagent.Name)
	return dummyagent
}

// StagePreserveOrder puts dummyagent to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DummyAgentOrder
// - update stage.DummyAgentOrder accordingly
func (dummyagent *DummyAgent) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DummyAgents, stage.DummyAgent_stagedOrder, stage.DummyAgent_orderStaged, &stage.DummyAgentOrder, stage.DummyAgents_mapString, dummyagent, order, dummyagent.Name)
}

// Unstage removes dummyagent off the model stage
func (dummyagent *DummyAgent) Unstage(stage *Stage) *DummyAgent {
	__gong__unstage(stage.DummyAgents, stage.DummyAgents_mapString, dummyagent, dummyagent.Name)
	return dummyagent
}

// UnstageVoid removes dummyagent off the model stage
func (dummyagent *DummyAgent) UnstageVoid(stage *Stage) {
	dummyagent.Unstage(stage)
}

func (dummyagent *DummyAgent) StageVoid(stage *Stage) {
	dummyagent.Stage(stage)
}

// for satisfaction of GongStruct interface
func (dummyagent *DummyAgent) GetName() (res string) {
	return dummyagent.Name
}

// for satisfaction of GongStruct interface
func (dummyagent *DummyAgent) SetName(name string) {
	dummyagent.Name = name
}

// Stage puts engine to the model stage
func (engine *Engine) Stage(stage *Stage) *Engine {
	__gong__stage(stage.Engines, stage.Engine_stagedOrder, stage.Engine_orderStaged, &stage.EngineOrder, stage.Engines_mapString, engine, engine.Name)
	return engine
}

// StagePreserveOrder puts engine to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EngineOrder
// - update stage.EngineOrder accordingly
func (engine *Engine) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Engines, stage.Engine_stagedOrder, stage.Engine_orderStaged, &stage.EngineOrder, stage.Engines_mapString, engine, order, engine.Name)
}

// Unstage removes engine off the model stage
func (engine *Engine) Unstage(stage *Stage) *Engine {
	__gong__unstage(stage.Engines, stage.Engines_mapString, engine, engine.Name)
	return engine
}

// UnstageVoid removes engine off the model stage
func (engine *Engine) UnstageVoid(stage *Stage) {
	engine.Unstage(stage)
}

func (engine *Engine) StageVoid(stage *Stage) {
	engine.Stage(stage)
}

// for satisfaction of GongStruct interface
func (engine *Engine) GetName() (res string) {
	return engine.Name
}

// for satisfaction of GongStruct interface
func (engine *Engine) SetName(name string) {
	engine.Name = name
}

// Stage puts event to the model stage
func (event *Event) Stage(stage *Stage) *Event {
	__gong__stage(stage.Events, stage.Event_stagedOrder, stage.Event_orderStaged, &stage.EventOrder, stage.Events_mapString, event, event.Name)
	return event
}

// StagePreserveOrder puts event to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EventOrder
// - update stage.EventOrder accordingly
func (event *Event) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Events, stage.Event_stagedOrder, stage.Event_orderStaged, &stage.EventOrder, stage.Events_mapString, event, order, event.Name)
}

// Unstage removes event off the model stage
func (event *Event) Unstage(stage *Stage) *Event {
	__gong__unstage(stage.Events, stage.Events_mapString, event, event.Name)
	return event
}

// UnstageVoid removes event off the model stage
func (event *Event) UnstageVoid(stage *Stage) {
	event.Unstage(stage)
}

func (event *Event) StageVoid(stage *Stage) {
	event.Stage(stage)
}

// for satisfaction of GongStruct interface
func (event *Event) GetName() (res string) {
	return event.Name
}

// for satisfaction of GongStruct interface
func (event *Event) SetName(name string) {
	event.Name = name
}

// Stage puts status to the model stage
func (status *Status) Stage(stage *Stage) *Status {
	__gong__stage(stage.Statuss, stage.Status_stagedOrder, stage.Status_orderStaged, &stage.StatusOrder, stage.Statuss_mapString, status, status.Name)
	return status
}

// StagePreserveOrder puts status to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StatusOrder
// - update stage.StatusOrder accordingly
func (status *Status) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Statuss, stage.Status_stagedOrder, stage.Status_orderStaged, &stage.StatusOrder, stage.Statuss_mapString, status, order, status.Name)
}

// Unstage removes status off the model stage
func (status *Status) Unstage(stage *Stage) *Status {
	__gong__unstage(stage.Statuss, stage.Statuss_mapString, status, status.Name)
	return status
}

// UnstageVoid removes status off the model stage
func (status *Status) UnstageVoid(stage *Stage) {
	status.Unstage(stage)
}

func (status *Status) StageVoid(stage *Stage) {
	status.Stage(stage)
}

// for satisfaction of GongStruct interface
func (status *Status) GetName() (res string) {
	return status.Name
}

// for satisfaction of GongStruct interface
func (status *Status) SetName(name string) {
	status.Name = name
}

// Stage puts updatestate to the model stage
func (updatestate *UpdateState) Stage(stage *Stage) *UpdateState {
	__gong__stage(stage.UpdateStates, stage.UpdateState_stagedOrder, stage.UpdateState_orderStaged, &stage.UpdateStateOrder, stage.UpdateStates_mapString, updatestate, updatestate.Name)
	return updatestate
}

// StagePreserveOrder puts updatestate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UpdateStateOrder
// - update stage.UpdateStateOrder accordingly
func (updatestate *UpdateState) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.UpdateStates, stage.UpdateState_stagedOrder, stage.UpdateState_orderStaged, &stage.UpdateStateOrder, stage.UpdateStates_mapString, updatestate, order, updatestate.Name)
}

// Unstage removes updatestate off the model stage
func (updatestate *UpdateState) Unstage(stage *Stage) *UpdateState {
	__gong__unstage(stage.UpdateStates, stage.UpdateStates_mapString, updatestate, updatestate.Name)
	return updatestate
}

// UnstageVoid removes updatestate off the model stage
func (updatestate *UpdateState) UnstageVoid(stage *Stage) {
	updatestate.Unstage(stage)
}

func (updatestate *UpdateState) StageVoid(stage *Stage) {
	updatestate.Stage(stage)
}

// for satisfaction of GongStruct interface
func (updatestate *UpdateState) GetName() (res string) {
	return updatestate.Name
}

// for satisfaction of GongStruct interface
func (updatestate *UpdateState) SetName(name string) {
	updatestate.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Commands, &stage.Commands_mapString, &stage.Command_stagedOrder, &stage.CommandOrder)

	__gong__resetStageType(&stage.DummyAgents, &stage.DummyAgents_mapString, &stage.DummyAgent_stagedOrder, &stage.DummyAgentOrder)

	__gong__resetStageType(&stage.Engines, &stage.Engines_mapString, &stage.Engine_stagedOrder, &stage.EngineOrder)

	__gong__resetStageType(&stage.Events, &stage.Events_mapString, &stage.Event_stagedOrder, &stage.EventOrder)

	__gong__resetStageType(&stage.Statuss, &stage.Statuss_mapString, &stage.Status_stagedOrder, &stage.StatusOrder)

	__gong__resetStageType(&stage.UpdateStates, &stage.UpdateStates_mapString, &stage.UpdateState_stagedOrder, &stage.UpdateStateOrder)

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
	case *Command:
		return any(stage.Commands_mapString).(map[string]Type)
	case *DummyAgent:
		return any(stage.DummyAgents_mapString).(map[string]Type)
	case *Engine:
		return any(stage.Engines_mapString).(map[string]Type)
	case *Event:
		return any(stage.Events_mapString).(map[string]Type)
	case *Status:
		return any(stage.Statuss_mapString).(map[string]Type)
	case *UpdateState:
		return any(stage.UpdateStates_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Command:
		return any(&stage.Commands).(*map[Type]struct{})
	case *DummyAgent:
		return any(&stage.DummyAgents).(*map[Type]struct{})
	case *Engine:
		return any(&stage.Engines).(*map[Type]struct{})
	case *Event:
		return any(&stage.Events).(*map[Type]struct{})
	case *Status:
		return any(&stage.Statuss).(*map[Type]struct{})
	case *UpdateState:
		return any(&stage.UpdateStates).(*map[Type]struct{})
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
	case Command:
		return any(&Command{
			Engine: &Engine{Name: "Engine"},
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
	// reverse maps of direct associations of Command
	case Command:
		switch fieldname {
		// insertion point for per direct association field
		case "Engine":
			res := make(map[*Engine][]*Command)
			for command := range stage.Commands {
				if command.Engine != nil {
					engine_ := command.Engine
					var commands []*Command
					_, ok := res[engine_]
					if ok {
						commands = res[engine_]
					} else {
						commands = make([]*Command, 0)
					}
					commands = append(commands, command)
					res[engine_] = commands
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DummyAgent
	case DummyAgent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Engine
	case Engine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Event
	case Event:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Status
	case Status:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UpdateState
	case UpdateState:
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
	// reverse maps of direct associations of Command
	case Command:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DummyAgent
	case DummyAgent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Engine
	case Engine:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Event
	case Event:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Status
	case Status:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of UpdateState
	case UpdateState:
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
	case *Command:
		res = any(new(Command)).(Type)
	case *DummyAgent:
		res = any(new(DummyAgent)).(Type)
	case *Engine:
		res = any(new(Engine)).(Type)
	case *Event:
		res = any(new(Event)).(Type)
	case *Status:
		res = any(new(Status)).(Type)
	case *UpdateState:
		res = any(new(UpdateState)).(Type)
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
	case *Command:
		res = "Command"
	case *DummyAgent:
		res = "DummyAgent"
	case *Engine:
		res = "Engine"
	case *Event:
		res = "Event"
	case *Status:
		res = "Status"
	case *UpdateState:
		res = "UpdateState"
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
	case *Command:
		var rf ReverseField
		_ = rf
	case *DummyAgent:
		var rf ReverseField
		_ = rf
	case *Engine:
		var rf ReverseField
		_ = rf
	case *Event:
		var rf ReverseField
		_ = rf
	case *Status:
		var rf ReverseField
		_ = rf
	case *UpdateState:
		var rf ReverseField
		_ = rf
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (command *Command) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Command",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "CommandType",
		},
		{
			Name:               "CommandDate",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Engine",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Engine",
		},
	}
	return
}

func (dummyagent *DummyAgent) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "TechName",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (engine *Engine) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "EndTime",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "CurrentTime",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "DisplayFormat",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SecondsSinceStart",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "Fired",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "ControlMode",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "ControlMode",
		},
		{
			Name:                 "State",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "EngineState",
		},
		{
			Name:               "Speed",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (event *Event) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeIntDuration,
		},
	}
	return
}

func (status *Status) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "CurrentCommand",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "CommandType",
		},
		{
			Name:               "CompletionDate",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "CurrentSpeedCommand",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "SpeedCommandType",
		},
		{
			Name:               "SpeedCommandCompletionDate",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (updatestate *UpdateState) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeIntDuration,
		},
		{
			Name:               "Period",
			GongFieldValueType: GongFieldValueTypeIntDuration,
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
func (command *Command) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = command.Name
	case "Command":
		enum := command.Command
		res.valueString = enum.ToCodeString()
	case "CommandDate":
		res.valueString = command.CommandDate
	case "Engine":
		res.GongFieldValueType = GongFieldValueTypePointer
		if command.Engine != nil {
			res.valueString = command.Engine.Name
			res.ids = command.Engine.GongGetUUID(stage)
		}
	}
	return
}

func (dummyagent *DummyAgent) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "TechName":
		res.valueString = dummyagent.TechName
	case "Name":
		res.valueString = dummyagent.Name
	}
	return
}

func (engine *Engine) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = engine.Name
	case "EndTime":
		res.valueString = engine.EndTime
	case "CurrentTime":
		res.valueString = engine.CurrentTime
	case "DisplayFormat":
		res.valueString = engine.DisplayFormat
	case "SecondsSinceStart":
		res.valueString = fmt.Sprintf("%f", engine.SecondsSinceStart)
		res.valueFloat = engine.SecondsSinceStart
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Fired":
		res.valueString = fmt.Sprintf("%d", engine.Fired)
		res.valueInt = engine.Fired
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ControlMode":
		enum := engine.ControlMode
		res.valueString = enum.ToCodeString()
	case "State":
		enum := engine.State
		res.valueString = enum.ToCodeString()
	case "Speed":
		res.valueString = fmt.Sprintf("%f", engine.Speed)
		res.valueFloat = engine.Speed
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (event *Event) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = event.Name
	case "Duration":
		if math.Abs(event.Duration.Hours()) >= 24 {
			days := __gong__abs(int(int(event.Duration.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(event.Duration.Hours()) % 24
			remainingMinutes := int(event.Duration.Minutes()) % 60
			remainingSeconds := int(event.Duration.Seconds()) % 60

			if event.Duration.Hours() < 0 {
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
			res.valueString = fmt.Sprintf("%s\n", event.Duration.String())
		}
	}
	return
}

func (status *Status) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = status.Name
	case "CurrentCommand":
		enum := status.CurrentCommand
		res.valueString = enum.ToCodeString()
	case "CompletionDate":
		res.valueString = status.CompletionDate
	case "CurrentSpeedCommand":
		enum := status.CurrentSpeedCommand
		res.valueString = enum.ToCodeString()
	case "SpeedCommandCompletionDate":
		res.valueString = status.SpeedCommandCompletionDate
	}
	return
}

func (updatestate *UpdateState) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = updatestate.Name
	case "Duration":
		if math.Abs(updatestate.Duration.Hours()) >= 24 {
			days := __gong__abs(int(int(updatestate.Duration.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(updatestate.Duration.Hours()) % 24
			remainingMinutes := int(updatestate.Duration.Minutes()) % 60
			remainingSeconds := int(updatestate.Duration.Seconds()) % 60

			if updatestate.Duration.Hours() < 0 {
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
			res.valueString = fmt.Sprintf("%s\n", updatestate.Duration.String())
		}
	case "Period":
		if math.Abs(updatestate.Period.Hours()) >= 24 {
			days := __gong__abs(int(int(updatestate.Period.Hours()) / 24))
			months := int(days / 31)
			days = days - months*31

			remainingHours := int(updatestate.Period.Hours()) % 24
			remainingMinutes := int(updatestate.Period.Minutes()) % 60
			remainingSeconds := int(updatestate.Period.Seconds()) % 60

			if updatestate.Period.Hours() < 0 {
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
			res.valueString = fmt.Sprintf("%s\n", updatestate.Period.String())
		}
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
func (command *Command) GongGetGongstructName() string {
	return "Command"
}

func (dummyagent *DummyAgent) GongGetGongstructName() string {
	return "DummyAgent"
}

func (engine *Engine) GongGetGongstructName() string {
	return "Engine"
}

func (event *Event) GongGetGongstructName() string {
	return "Event"
}

func (status *Status) GongGetGongstructName() string {
	return "Status"
}

func (updatestate *UpdateState) GongGetGongstructName() string {
	return "UpdateState"
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
	__gong__rebuildMapString(stage.Commands, &stage.Commands_mapString)

	__gong__rebuildMapString(stage.DummyAgents, &stage.DummyAgents_mapString)

	__gong__rebuildMapString(stage.Engines, &stage.Engines_mapString)

	__gong__rebuildMapString(stage.Events, &stage.Events_mapString)

	__gong__rebuildMapString(stage.Statuss, &stage.Statuss_mapString)

	__gong__rebuildMapString(stage.UpdateStates, &stage.UpdateStates_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
