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
	"time"

	sim_go "github.com/fullstack-lang/gong/lib/sim/go"
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

var _ = __Gong__Abs
var _ = strings.Clone("")

const ProbeTreeSidebarSuffix = ":sidebar of the probe"
const ProbeTableSuffix = ":table of the probe"
const ProbeNotificationTableSuffix = ":notification table of the probe"
const ProbeFormSuffix = ":form of the probe"
const ProbeSplitSuffix = ":probe of the probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
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
var errUnkownEnum = errors.New("unkown enum")
var _ = errUnkownEnum

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void
var _ = __member

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

	// insertion point for definition of arrays registering instances
	Commands                map[*Command]struct{}
	Commands_reference      map[*Command]*Command
	Commands_referenceOrder map[*Command]uint // diff Unstage needs the reference order
	Commands_mapString      map[string]*Command

	// insertion point for slice of pointers maps
	OnAfterCommandCreateCallback OnAfterCreateInterface[Command]
	OnAfterCommandUpdateCallback OnAfterUpdateInterface[Command]
	OnAfterCommandDeleteCallback OnAfterDeleteInterface[Command]
	OnAfterCommandReadCallback   OnAfterReadInterface[Command]

	DummyAgents                map[*DummyAgent]struct{}
	DummyAgents_reference      map[*DummyAgent]*DummyAgent
	DummyAgents_referenceOrder map[*DummyAgent]uint // diff Unstage needs the reference order
	DummyAgents_mapString      map[string]*DummyAgent

	// insertion point for slice of pointers maps
	OnAfterDummyAgentCreateCallback OnAfterCreateInterface[DummyAgent]
	OnAfterDummyAgentUpdateCallback OnAfterUpdateInterface[DummyAgent]
	OnAfterDummyAgentDeleteCallback OnAfterDeleteInterface[DummyAgent]
	OnAfterDummyAgentReadCallback   OnAfterReadInterface[DummyAgent]

	Engines                map[*Engine]struct{}
	Engines_reference      map[*Engine]*Engine
	Engines_referenceOrder map[*Engine]uint // diff Unstage needs the reference order
	Engines_mapString      map[string]*Engine

	// insertion point for slice of pointers maps
	OnAfterEngineCreateCallback OnAfterCreateInterface[Engine]
	OnAfterEngineUpdateCallback OnAfterUpdateInterface[Engine]
	OnAfterEngineDeleteCallback OnAfterDeleteInterface[Engine]
	OnAfterEngineReadCallback   OnAfterReadInterface[Engine]

	Events                map[*Event]struct{}
	Events_reference      map[*Event]*Event
	Events_referenceOrder map[*Event]uint // diff Unstage needs the reference order
	Events_mapString      map[string]*Event

	// insertion point for slice of pointers maps
	OnAfterEventCreateCallback OnAfterCreateInterface[Event]
	OnAfterEventUpdateCallback OnAfterUpdateInterface[Event]
	OnAfterEventDeleteCallback OnAfterDeleteInterface[Event]
	OnAfterEventReadCallback   OnAfterReadInterface[Event]

	Statuss                map[*Status]struct{}
	Statuss_reference      map[*Status]*Status
	Statuss_referenceOrder map[*Status]uint // diff Unstage needs the reference order
	Statuss_mapString      map[string]*Status

	// insertion point for slice of pointers maps
	OnAfterStatusCreateCallback OnAfterCreateInterface[Status]
	OnAfterStatusUpdateCallback OnAfterUpdateInterface[Status]
	OnAfterStatusDeleteCallback OnAfterDeleteInterface[Status]
	OnAfterStatusReadCallback   OnAfterReadInterface[Status]

	UpdateStates                map[*UpdateState]struct{}
	UpdateStates_reference      map[*UpdateState]*UpdateState
	UpdateStates_referenceOrder map[*UpdateState]uint // diff Unstage needs the reference order
	UpdateStates_mapString      map[string]*UpdateState

	// insertion point for slice of pointers maps
	OnAfterUpdateStateCreateCallback OnAfterCreateInterface[UpdateState]
	OnAfterUpdateStateUpdateCallback OnAfterUpdateInterface[UpdateState]
	OnAfterUpdateStateDeleteCallback OnAfterDeleteInterface[UpdateState]
	OnAfterUpdateStateReadCallback   OnAfterReadInterface[UpdateState]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

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
	CommandOrder            uint
	CommandMap_Staged_Order map[*Command]uint

	DummyAgentOrder            uint
	DummyAgentMap_Staged_Order map[*DummyAgent]uint

	EngineOrder            uint
	EngineMap_Staged_Order map[*Engine]uint

	EventOrder            uint
	EventMap_Staged_Order map[*Event]uint

	StatusOrder            uint
	StatusMap_Staged_Order map[*Status]uint

	UpdateStateOrder            uint
	UpdateStateMap_Staged_Order map[*UpdateState]uint

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

	if stage.commitsBehind >= len(stage.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := stage.backwardCommits[len(stage.backwardCommits)-1-stage.commitsBehind]

	// umarshall the backward commit to the stage

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind++
	err := GongParseAstString(stage, commitToApply, true)
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
	err := GongParseAstString(stage, commitToApply, true)
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
	stage.ComputeReferenceAndOrders()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation 
	var maxCommandOrder uint
	var foundCommand bool
	for _, order := range stage.CommandMap_Staged_Order {
		if !foundCommand || order > maxCommandOrder {
			maxCommandOrder = order
			foundCommand = true
		}
	}
	if foundCommand {
		stage.CommandOrder = maxCommandOrder + 1
	} else {
		stage.CommandOrder = 0
	}

	var maxDummyAgentOrder uint
	var foundDummyAgent bool
	for _, order := range stage.DummyAgentMap_Staged_Order {
		if !foundDummyAgent || order > maxDummyAgentOrder {
			maxDummyAgentOrder = order
			foundDummyAgent = true
		}
	}
	if foundDummyAgent {
		stage.DummyAgentOrder = maxDummyAgentOrder + 1
	} else {
		stage.DummyAgentOrder = 0
	}

	var maxEngineOrder uint
	var foundEngine bool
	for _, order := range stage.EngineMap_Staged_Order {
		if !foundEngine || order > maxEngineOrder {
			maxEngineOrder = order
			foundEngine = true
		}
	}
	if foundEngine {
		stage.EngineOrder = maxEngineOrder + 1
	} else {
		stage.EngineOrder = 0
	}

	var maxEventOrder uint
	var foundEvent bool
	for _, order := range stage.EventMap_Staged_Order {
		if !foundEvent || order > maxEventOrder {
			maxEventOrder = order
			foundEvent = true
		}
	}
	if foundEvent {
		stage.EventOrder = maxEventOrder + 1
	} else {
		stage.EventOrder = 0
	}

	var maxStatusOrder uint
	var foundStatus bool
	for _, order := range stage.StatusMap_Staged_Order {
		if !foundStatus || order > maxStatusOrder {
			maxStatusOrder = order
			foundStatus = true
		}
	}
	if foundStatus {
		stage.StatusOrder = maxStatusOrder + 1
	} else {
		stage.StatusOrder = 0
	}

	var maxUpdateStateOrder uint
	var foundUpdateState bool
	for _, order := range stage.UpdateStateMap_Staged_Order {
		if !foundUpdateState || order > maxUpdateStateOrder {
			maxUpdateStateOrder = order
			foundUpdateState = true
		}
	}
	if foundUpdateState {
		stage.UpdateStateOrder = maxUpdateStateOrder + 1
	} else {
		stage.UpdateStateOrder = 0
	}

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

func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Command:
		tmp := GetStructInstancesByOrder(stage.Commands, stage.CommandMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Command implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DummyAgent:
		tmp := GetStructInstancesByOrder(stage.DummyAgents, stage.DummyAgentMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DummyAgent implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Engine:
		tmp := GetStructInstancesByOrder(stage.Engines, stage.EngineMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Engine implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Event:
		tmp := GetStructInstancesByOrder(stage.Events, stage.EventMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Event implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Status:
		tmp := GetStructInstancesByOrder(stage.Statuss, stage.StatusMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Status implements.
			res = append(res, any(v).(T))
		}
		return res
	case *UpdateState:
		tmp := GetStructInstancesByOrder(stage.UpdateStates, stage.UpdateStateMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *UpdateState implements.
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
	case "Command":
		res = GetNamedStructInstances(stage.Commands, stage.CommandMap_Staged_Order)
	case "DummyAgent":
		res = GetNamedStructInstances(stage.DummyAgents, stage.DummyAgentMap_Staged_Order)
	case "Engine":
		res = GetNamedStructInstances(stage.Engines, stage.EngineMap_Staged_Order)
	case "Event":
		res = GetNamedStructInstances(stage.Events, stage.EventMap_Staged_Order)
	case "Status":
		res = GetNamedStructInstances(stage.Statuss, stage.StatusMap_Staged_Order)
	case "UpdateState":
		res = GetNamedStructInstances(stage.UpdateStates, stage.UpdateStateMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/sim/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return sim_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return sim_go.GoDiagramsDir
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
	CommitCommand(command *Command)
	CheckoutCommand(command *Command)
	CommitDummyAgent(dummyagent *DummyAgent)
	CheckoutDummyAgent(dummyagent *DummyAgent)
	CommitEngine(engine *Engine)
	CheckoutEngine(engine *Engine)
	CommitEvent(event *Event)
	CheckoutEvent(event *Event)
	CommitStatus(status *Status)
	CheckoutStatus(status *Status)
	CommitUpdateState(updatestate *UpdateState)
	CheckoutUpdateState(updatestate *UpdateState)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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
		CommandMap_Staged_Order: make(map[*Command]uint),

		DummyAgentMap_Staged_Order: make(map[*DummyAgent]uint),

		EngineMap_Staged_Order: make(map[*Engine]uint),

		EventMap_Staged_Order: make(map[*Event]uint),

		StatusMap_Staged_Order: make(map[*Status]uint),

		UpdateStateMap_Staged_Order: make(map[*UpdateState]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Command": &CommandUnmarshaller{},

			"DummyAgent": &DummyAgentUnmarshaller{},

			"Engine": &EngineUnmarshaller{},

			"Event": &EventUnmarshaller{},

			"Status": &StatusUnmarshaller{},

			"UpdateState": &UpdateStateUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Command"},
			{name: "DummyAgent"},
			{name: "Engine"},
			{name: "Event"},
			{name: "Status"},
			{name: "UpdateState"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Command:
		return stage.CommandMap_Staged_Order[instance]
	case *DummyAgent:
		return stage.DummyAgentMap_Staged_Order[instance]
	case *Engine:
		return stage.EngineMap_Staged_Order[instance]
	case *Event:
		return stage.EventMap_Staged_Order[instance]
	case *Status:
		return stage.StatusMap_Staged_Order[instance]
	case *UpdateState:
		return stage.UpdateStateMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Command:
		return stage.CommandMap_Staged_Order[instance]
	case *DummyAgent:
		return stage.DummyAgentMap_Staged_Order[instance]
	case *Engine:
		return stage.EngineMap_Staged_Order[instance]
	case *Event:
		return stage.EventMap_Staged_Order[instance]
	case *Status:
		return stage.StatusMap_Staged_Order[instance]
	case *UpdateState:
		return stage.UpdateStateMap_Staged_Order[instance]
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
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
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
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().Refresh()
		}
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

	if _, ok := stage.Commands[command]; !ok {
		stage.Commands[command] = struct{}{}
		stage.CommandMap_Staged_Order[command] = stage.CommandOrder
		stage.CommandOrder++
	}
	stage.Commands_mapString[command.Name] = command

	return command
}

// StagePreserveOrder puts command to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CommandOrder
// - update stage.CommandOrder accordingly
func (command *Command) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Commands[command]; !ok {
		stage.Commands[command] = struct{}{}

		if order > stage.CommandOrder {
			stage.CommandOrder = order
		}
		stage.CommandMap_Staged_Order[command] = order
		stage.CommandOrder++
	}
	stage.Commands_mapString[command.Name] = command
}

// Unstage removes command off the model stage
func (command *Command) Unstage(stage *Stage) *Command {
	delete(stage.Commands, command)
	delete(stage.CommandMap_Staged_Order, command)
	delete(stage.Commands_mapString, command.Name)

	return command
}

// UnstageVoid removes command off the model stage
func (command *Command) UnstageVoid(stage *Stage) {
	delete(stage.Commands, command)
	delete(stage.CommandMap_Staged_Order, command)
	delete(stage.Commands_mapString, command.Name)
}

// commit command to the back repo (if it is already staged)
func (command *Command) Commit(stage *Stage) *Command {
	if _, ok := stage.Commands[command]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCommand(command)
		}
	}
	return command
}

func (command *Command) CommitVoid(stage *Stage) {
	command.Commit(stage)
}

func (command *Command) StageVoid(stage *Stage) {
	command.Stage(stage)
}

// Checkout command to the back repo (if it is already staged)
func (command *Command) Checkout(stage *Stage) *Command {
	if _, ok := stage.Commands[command]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCommand(command)
		}
	}
	return command
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

	if _, ok := stage.DummyAgents[dummyagent]; !ok {
		stage.DummyAgents[dummyagent] = struct{}{}
		stage.DummyAgentMap_Staged_Order[dummyagent] = stage.DummyAgentOrder
		stage.DummyAgentOrder++
	}
	stage.DummyAgents_mapString[dummyagent.Name] = dummyagent

	return dummyagent
}

// StagePreserveOrder puts dummyagent to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DummyAgentOrder
// - update stage.DummyAgentOrder accordingly
func (dummyagent *DummyAgent) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.DummyAgents[dummyagent]; !ok {
		stage.DummyAgents[dummyagent] = struct{}{}

		if order > stage.DummyAgentOrder {
			stage.DummyAgentOrder = order
		}
		stage.DummyAgentMap_Staged_Order[dummyagent] = order
		stage.DummyAgentOrder++
	}
	stage.DummyAgents_mapString[dummyagent.Name] = dummyagent
}

// Unstage removes dummyagent off the model stage
func (dummyagent *DummyAgent) Unstage(stage *Stage) *DummyAgent {
	delete(stage.DummyAgents, dummyagent)
	delete(stage.DummyAgentMap_Staged_Order, dummyagent)
	delete(stage.DummyAgents_mapString, dummyagent.Name)

	return dummyagent
}

// UnstageVoid removes dummyagent off the model stage
func (dummyagent *DummyAgent) UnstageVoid(stage *Stage) {
	delete(stage.DummyAgents, dummyagent)
	delete(stage.DummyAgentMap_Staged_Order, dummyagent)
	delete(stage.DummyAgents_mapString, dummyagent.Name)
}

// commit dummyagent to the back repo (if it is already staged)
func (dummyagent *DummyAgent) Commit(stage *Stage) *DummyAgent {
	if _, ok := stage.DummyAgents[dummyagent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDummyAgent(dummyagent)
		}
	}
	return dummyagent
}

func (dummyagent *DummyAgent) CommitVoid(stage *Stage) {
	dummyagent.Commit(stage)
}

func (dummyagent *DummyAgent) StageVoid(stage *Stage) {
	dummyagent.Stage(stage)
}

// Checkout dummyagent to the back repo (if it is already staged)
func (dummyagent *DummyAgent) Checkout(stage *Stage) *DummyAgent {
	if _, ok := stage.DummyAgents[dummyagent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDummyAgent(dummyagent)
		}
	}
	return dummyagent
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

	if _, ok := stage.Engines[engine]; !ok {
		stage.Engines[engine] = struct{}{}
		stage.EngineMap_Staged_Order[engine] = stage.EngineOrder
		stage.EngineOrder++
	}
	stage.Engines_mapString[engine.Name] = engine

	return engine
}

// StagePreserveOrder puts engine to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EngineOrder
// - update stage.EngineOrder accordingly
func (engine *Engine) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Engines[engine]; !ok {
		stage.Engines[engine] = struct{}{}

		if order > stage.EngineOrder {
			stage.EngineOrder = order
		}
		stage.EngineMap_Staged_Order[engine] = order
		stage.EngineOrder++
	}
	stage.Engines_mapString[engine.Name] = engine
}

// Unstage removes engine off the model stage
func (engine *Engine) Unstage(stage *Stage) *Engine {
	delete(stage.Engines, engine)
	delete(stage.EngineMap_Staged_Order, engine)
	delete(stage.Engines_mapString, engine.Name)

	return engine
}

// UnstageVoid removes engine off the model stage
func (engine *Engine) UnstageVoid(stage *Stage) {
	delete(stage.Engines, engine)
	delete(stage.EngineMap_Staged_Order, engine)
	delete(stage.Engines_mapString, engine.Name)
}

// commit engine to the back repo (if it is already staged)
func (engine *Engine) Commit(stage *Stage) *Engine {
	if _, ok := stage.Engines[engine]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEngine(engine)
		}
	}
	return engine
}

func (engine *Engine) CommitVoid(stage *Stage) {
	engine.Commit(stage)
}

func (engine *Engine) StageVoid(stage *Stage) {
	engine.Stage(stage)
}

// Checkout engine to the back repo (if it is already staged)
func (engine *Engine) Checkout(stage *Stage) *Engine {
	if _, ok := stage.Engines[engine]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEngine(engine)
		}
	}
	return engine
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

	if _, ok := stage.Events[event]; !ok {
		stage.Events[event] = struct{}{}
		stage.EventMap_Staged_Order[event] = stage.EventOrder
		stage.EventOrder++
	}
	stage.Events_mapString[event.Name] = event

	return event
}

// StagePreserveOrder puts event to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.EventOrder
// - update stage.EventOrder accordingly
func (event *Event) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Events[event]; !ok {
		stage.Events[event] = struct{}{}

		if order > stage.EventOrder {
			stage.EventOrder = order
		}
		stage.EventMap_Staged_Order[event] = order
		stage.EventOrder++
	}
	stage.Events_mapString[event.Name] = event
}

// Unstage removes event off the model stage
func (event *Event) Unstage(stage *Stage) *Event {
	delete(stage.Events, event)
	delete(stage.EventMap_Staged_Order, event)
	delete(stage.Events_mapString, event.Name)

	return event
}

// UnstageVoid removes event off the model stage
func (event *Event) UnstageVoid(stage *Stage) {
	delete(stage.Events, event)
	delete(stage.EventMap_Staged_Order, event)
	delete(stage.Events_mapString, event.Name)
}

// commit event to the back repo (if it is already staged)
func (event *Event) Commit(stage *Stage) *Event {
	if _, ok := stage.Events[event]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEvent(event)
		}
	}
	return event
}

func (event *Event) CommitVoid(stage *Stage) {
	event.Commit(stage)
}

func (event *Event) StageVoid(stage *Stage) {
	event.Stage(stage)
}

// Checkout event to the back repo (if it is already staged)
func (event *Event) Checkout(stage *Stage) *Event {
	if _, ok := stage.Events[event]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEvent(event)
		}
	}
	return event
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

	if _, ok := stage.Statuss[status]; !ok {
		stage.Statuss[status] = struct{}{}
		stage.StatusMap_Staged_Order[status] = stage.StatusOrder
		stage.StatusOrder++
	}
	stage.Statuss_mapString[status.Name] = status

	return status
}

// StagePreserveOrder puts status to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StatusOrder
// - update stage.StatusOrder accordingly
func (status *Status) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Statuss[status]; !ok {
		stage.Statuss[status] = struct{}{}

		if order > stage.StatusOrder {
			stage.StatusOrder = order
		}
		stage.StatusMap_Staged_Order[status] = order
		stage.StatusOrder++
	}
	stage.Statuss_mapString[status.Name] = status
}

// Unstage removes status off the model stage
func (status *Status) Unstage(stage *Stage) *Status {
	delete(stage.Statuss, status)
	delete(stage.StatusMap_Staged_Order, status)
	delete(stage.Statuss_mapString, status.Name)

	return status
}

// UnstageVoid removes status off the model stage
func (status *Status) UnstageVoid(stage *Stage) {
	delete(stage.Statuss, status)
	delete(stage.StatusMap_Staged_Order, status)
	delete(stage.Statuss_mapString, status.Name)
}

// commit status to the back repo (if it is already staged)
func (status *Status) Commit(stage *Stage) *Status {
	if _, ok := stage.Statuss[status]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitStatus(status)
		}
	}
	return status
}

func (status *Status) CommitVoid(stage *Stage) {
	status.Commit(stage)
}

func (status *Status) StageVoid(stage *Stage) {
	status.Stage(stage)
}

// Checkout status to the back repo (if it is already staged)
func (status *Status) Checkout(stage *Stage) *Status {
	if _, ok := stage.Statuss[status]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutStatus(status)
		}
	}
	return status
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

	if _, ok := stage.UpdateStates[updatestate]; !ok {
		stage.UpdateStates[updatestate] = struct{}{}
		stage.UpdateStateMap_Staged_Order[updatestate] = stage.UpdateStateOrder
		stage.UpdateStateOrder++
	}
	stage.UpdateStates_mapString[updatestate.Name] = updatestate

	return updatestate
}

// StagePreserveOrder puts updatestate to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.UpdateStateOrder
// - update stage.UpdateStateOrder accordingly
func (updatestate *UpdateState) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.UpdateStates[updatestate]; !ok {
		stage.UpdateStates[updatestate] = struct{}{}

		if order > stage.UpdateStateOrder {
			stage.UpdateStateOrder = order
		}
		stage.UpdateStateMap_Staged_Order[updatestate] = order
		stage.UpdateStateOrder++
	}
	stage.UpdateStates_mapString[updatestate.Name] = updatestate
}

// Unstage removes updatestate off the model stage
func (updatestate *UpdateState) Unstage(stage *Stage) *UpdateState {
	delete(stage.UpdateStates, updatestate)
	delete(stage.UpdateStateMap_Staged_Order, updatestate)
	delete(stage.UpdateStates_mapString, updatestate.Name)

	return updatestate
}

// UnstageVoid removes updatestate off the model stage
func (updatestate *UpdateState) UnstageVoid(stage *Stage) {
	delete(stage.UpdateStates, updatestate)
	delete(stage.UpdateStateMap_Staged_Order, updatestate)
	delete(stage.UpdateStates_mapString, updatestate.Name)
}

// commit updatestate to the back repo (if it is already staged)
func (updatestate *UpdateState) Commit(stage *Stage) *UpdateState {
	if _, ok := stage.UpdateStates[updatestate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUpdateState(updatestate)
		}
	}
	return updatestate
}

func (updatestate *UpdateState) CommitVoid(stage *Stage) {
	updatestate.Commit(stage)
}

func (updatestate *UpdateState) StageVoid(stage *Stage) {
	updatestate.Stage(stage)
}

// Checkout updatestate to the back repo (if it is already staged)
func (updatestate *UpdateState) Checkout(stage *Stage) *UpdateState {
	if _, ok := stage.UpdateStates[updatestate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUpdateState(updatestate)
		}
	}
	return updatestate
}

// for satisfaction of GongStruct interface
func (updatestate *UpdateState) GetName() (res string) {
	return updatestate.Name
}

// for satisfaction of GongStruct interface
func (updatestate *UpdateState) SetName(name string) {
	updatestate.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCommand(Command *Command)
	CreateORMDummyAgent(DummyAgent *DummyAgent)
	CreateORMEngine(Engine *Engine)
	CreateORMEvent(Event *Event)
	CreateORMStatus(Status *Status)
	CreateORMUpdateState(UpdateState *UpdateState)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCommand(Command *Command)
	DeleteORMDummyAgent(DummyAgent *DummyAgent)
	DeleteORMEngine(Engine *Engine)
	DeleteORMEvent(Event *Event)
	DeleteORMStatus(Status *Status)
	DeleteORMUpdateState(UpdateState *UpdateState)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Commands = make(map[*Command]struct{})
	stage.Commands_mapString = make(map[string]*Command)
	stage.CommandMap_Staged_Order = make(map[*Command]uint)
	stage.CommandOrder = 0

	stage.DummyAgents = make(map[*DummyAgent]struct{})
	stage.DummyAgents_mapString = make(map[string]*DummyAgent)
	stage.DummyAgentMap_Staged_Order = make(map[*DummyAgent]uint)
	stage.DummyAgentOrder = 0

	stage.Engines = make(map[*Engine]struct{})
	stage.Engines_mapString = make(map[string]*Engine)
	stage.EngineMap_Staged_Order = make(map[*Engine]uint)
	stage.EngineOrder = 0

	stage.Events = make(map[*Event]struct{})
	stage.Events_mapString = make(map[string]*Event)
	stage.EventMap_Staged_Order = make(map[*Event]uint)
	stage.EventOrder = 0

	stage.Statuss = make(map[*Status]struct{})
	stage.Statuss_mapString = make(map[string]*Status)
	stage.StatusMap_Staged_Order = make(map[*Status]uint)
	stage.StatusOrder = 0

	stage.UpdateStates = make(map[*UpdateState]struct{})
	stage.UpdateStates_mapString = make(map[string]*UpdateState)
	stage.UpdateStateMap_Staged_Order = make(map[*UpdateState]uint)
	stage.UpdateStateOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Commands = nil
	stage.Commands_mapString = nil

	stage.DummyAgents = nil
	stage.DummyAgents_mapString = nil

	stage.Engines = nil
	stage.Engines_mapString = nil

	stage.Events = nil
	stage.Events_mapString = nil

	stage.Statuss = nil
	stage.Statuss_mapString = nil

	stage.UpdateStates = nil
	stage.UpdateStates_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for command := range stage.Commands {
		command.Unstage(stage)
	}

	for dummyagent := range stage.DummyAgents {
		dummyagent.Unstage(stage)
	}

	for engine := range stage.Engines {
		engine.Unstage(stage)
	}

	for event := range stage.Events {
		event.Unstage(stage)
	}

	for status := range stage.Statuss {
		status.Unstage(stage)
	}

	for updatestate := range stage.UpdateStates {
		updatestate.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

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
	case map[*Command]any:
		return any(&stage.Commands).(*Type)
	case map[*DummyAgent]any:
		return any(&stage.DummyAgents).(*Type)
	case map[*Engine]any:
		return any(&stage.Engines).(*Type)
	case map[*Event]any:
		return any(&stage.Events).(*Type)
	case map[*Status]any:
		return any(&stage.Statuss).(*Type)
	case map[*UpdateState]any:
		return any(&stage.UpdateStates).(*Type)
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

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Command:
		return any(&stage.Commands).(*map[*Type]struct{})
	case DummyAgent:
		return any(&stage.DummyAgents).(*map[*Type]struct{})
	case Engine:
		return any(&stage.Engines).(*map[*Type]struct{})
	case Event:
		return any(&stage.Events).(*map[*Type]struct{})
	case Status:
		return any(&stage.Statuss).(*map[*Type]struct{})
	case UpdateState:
		return any(&stage.UpdateStates).(*map[*Type]struct{})
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Command:
		return any(&stage.Commands_mapString).(*map[string]*Type)
	case DummyAgent:
		return any(&stage.DummyAgents_mapString).(*map[string]*Type)
	case Engine:
		return any(&stage.Engines_mapString).(*map[string]*Type)
	case Event:
		return any(&stage.Events_mapString).(*map[string]*Type)
	case Status:
		return any(&stage.Statuss_mapString).(*map[string]*Type)
	case UpdateState:
		return any(&stage.UpdateStates_mapString).(*map[string]*Type)
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
	case Command:
		return any(&Command{
			// Initialisation of associations
			// field is initialized with an instance of Engine with the name of the field
			Engine: &Engine{Name: "Engine"},
		}).(*Type)
	case DummyAgent:
		return any(&DummyAgent{
			// Initialisation of associations
		}).(*Type)
	case Engine:
		return any(&Engine{
			// Initialisation of associations
		}).(*Type)
	case Event:
		return any(&Event{
			// Initialisation of associations
		}).(*Type)
	case Status:
		return any(&Status{
			// Initialisation of associations
		}).(*Type)
	case UpdateState:
		return any(&UpdateState{
			// Initialisation of associations
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {

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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {

	res = make([]ReverseField, 0)

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

// insertion point for get fields header method
func (command *Command) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Command",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CommandDate",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (engine *Engine) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndTime",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CurrentTime",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DisplayFormat",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "SecondsSinceStart",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Fired",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ControlMode",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "State",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Speed",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (event *Event) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (status *Status) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CurrentCommand",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CompletionDate",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CurrentSpeedCommand",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "SpeedCommandCompletionDate",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (updatestate *UpdateState) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Duration",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Period",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, command.Engine))
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
			days := __Gong__Abs(int(int(event.Duration.Hours()) / 24))
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
			days := __Gong__Abs(int(int(updatestate.Duration.Hours()) / 24))
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
			days := __Gong__Abs(int(int(updatestate.Period.Hours()) / 24))
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
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (command *Command) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		command.Name = value.GetValueString()
	case "Command":
		command.Command.FromCodeString(value.GetValueString())
	case "CommandDate":
		command.CommandDate = value.GetValueString()
	case "Engine":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			command.Engine = nil
			for __instance__ := range stage.Engines {
				if stage.EngineMap_Staged_Order[__instance__] == uint(id) {
					command.Engine = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (dummyagent *DummyAgent) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "TechName":
		dummyagent.TechName = value.GetValueString()
	case "Name":
		dummyagent.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (engine *Engine) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		engine.Name = value.GetValueString()
	case "EndTime":
		engine.EndTime = value.GetValueString()
	case "CurrentTime":
		engine.CurrentTime = value.GetValueString()
	case "DisplayFormat":
		engine.DisplayFormat = value.GetValueString()
	case "SecondsSinceStart":
		engine.SecondsSinceStart = value.GetValueFloat()
	case "Fired":
		engine.Fired = int(value.GetValueInt())
	case "ControlMode":
		engine.ControlMode.FromCodeString(value.GetValueString())
	case "State":
		engine.State.FromCodeString(value.GetValueString())
	case "Speed":
		engine.Speed = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (event *Event) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		event.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (status *Status) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		status.Name = value.GetValueString()
	case "CurrentCommand":
		status.CurrentCommand.FromCodeString(value.GetValueString())
	case "CompletionDate":
		status.CompletionDate = value.GetValueString()
	case "CurrentSpeedCommand":
		status.CurrentSpeedCommand.FromCodeString(value.GetValueString())
	case "SpeedCommandCompletionDate":
		status.SpeedCommandCompletionDate = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (updatestate *UpdateState) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		updatestate.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Commands_mapString = make(map[string]*Command)
	for command := range stage.Commands {
		stage.Commands_mapString[command.Name] = command
	}

	stage.DummyAgents_mapString = make(map[string]*DummyAgent)
	for dummyagent := range stage.DummyAgents {
		stage.DummyAgents_mapString[dummyagent.Name] = dummyagent
	}

	stage.Engines_mapString = make(map[string]*Engine)
	for engine := range stage.Engines {
		stage.Engines_mapString[engine.Name] = engine
	}

	stage.Events_mapString = make(map[string]*Event)
	for event := range stage.Events {
		stage.Events_mapString[event.Name] = event
	}

	stage.Statuss_mapString = make(map[string]*Status)
	for status := range stage.Statuss {
		stage.Statuss_mapString[status.Name] = status
	}

	stage.UpdateStates_mapString = make(map[string]*UpdateState)
	for updatestate := range stage.UpdateStates {
		stage.UpdateStates_mapString[updatestate.Name] = updatestate
	}

}

// Last line of the template
