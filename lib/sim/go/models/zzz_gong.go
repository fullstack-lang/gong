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

const ProbeTreeSidebarSuffix = "-sidebar"
const ProbeTableSuffix = "-table"
const ProbeFormSuffix = "-form"
const ProbeSplitSuffix = "-probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

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

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Commands           map[*Command]any
	Commands_mapString map[string]*Command

	// insertion point for slice of pointers maps
	OnAfterCommandCreateCallback OnAfterCreateInterface[Command]
	OnAfterCommandUpdateCallback OnAfterUpdateInterface[Command]
	OnAfterCommandDeleteCallback OnAfterDeleteInterface[Command]
	OnAfterCommandReadCallback   OnAfterReadInterface[Command]

	DummyAgents           map[*DummyAgent]any
	DummyAgents_mapString map[string]*DummyAgent

	// insertion point for slice of pointers maps
	OnAfterDummyAgentCreateCallback OnAfterCreateInterface[DummyAgent]
	OnAfterDummyAgentUpdateCallback OnAfterUpdateInterface[DummyAgent]
	OnAfterDummyAgentDeleteCallback OnAfterDeleteInterface[DummyAgent]
	OnAfterDummyAgentReadCallback   OnAfterReadInterface[DummyAgent]

	Engines           map[*Engine]any
	Engines_mapString map[string]*Engine

	// insertion point for slice of pointers maps
	OnAfterEngineCreateCallback OnAfterCreateInterface[Engine]
	OnAfterEngineUpdateCallback OnAfterUpdateInterface[Engine]
	OnAfterEngineDeleteCallback OnAfterDeleteInterface[Engine]
	OnAfterEngineReadCallback   OnAfterReadInterface[Engine]

	Events           map[*Event]any
	Events_mapString map[string]*Event

	// insertion point for slice of pointers maps
	OnAfterEventCreateCallback OnAfterCreateInterface[Event]
	OnAfterEventUpdateCallback OnAfterUpdateInterface[Event]
	OnAfterEventDeleteCallback OnAfterDeleteInterface[Event]
	OnAfterEventReadCallback   OnAfterReadInterface[Event]

	Statuss           map[*Status]any
	Statuss_mapString map[string]*Status

	// insertion point for slice of pointers maps
	OnAfterStatusCreateCallback OnAfterCreateInterface[Status]
	OnAfterStatusUpdateCallback OnAfterUpdateInterface[Status]
	OnAfterStatusDeleteCallback OnAfterDeleteInterface[Status]
	OnAfterStatusReadCallback   OnAfterReadInterface[Status]

	UpdateStates           map[*UpdateState]any
	UpdateStates_mapString map[string]*UpdateState

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
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []string) {

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
		Commands:           make(map[*Command]any),
		Commands_mapString: make(map[string]*Command),

		DummyAgents:           make(map[*DummyAgent]any),
		DummyAgents_mapString: make(map[string]*DummyAgent),

		Engines:           make(map[*Engine]any),
		Engines_mapString: make(map[string]*Engine),

		Events:           make(map[*Event]any),
		Events_mapString: make(map[string]*Event),

		Statuss:           make(map[*Status]any),
		Statuss_mapString: make(map[string]*Status),

		UpdateStates:           make(map[*UpdateState]any),
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Command"},
			{name: "DummyAgent"},
			{name: "Engine"},
			{name: "Event"},
			{name: "Status"},
			{name: "UpdateState"},
		}, // end of insertion point
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

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

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
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Command"] = len(stage.Commands)
	stage.Map_GongStructName_InstancesNb["DummyAgent"] = len(stage.DummyAgents)
	stage.Map_GongStructName_InstancesNb["Engine"] = len(stage.Engines)
	stage.Map_GongStructName_InstancesNb["Event"] = len(stage.Events)
	stage.Map_GongStructName_InstancesNb["Status"] = len(stage.Statuss)
	stage.Map_GongStructName_InstancesNb["UpdateState"] = len(stage.UpdateStates)

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
		stage.Commands[command] = __member
		stage.CommandMap_Staged_Order[command] = stage.CommandOrder
		stage.CommandOrder++
	}
	stage.Commands_mapString[command.Name] = command

	return command
}

// Unstage removes command off the model stage
func (command *Command) Unstage(stage *Stage) *Command {
	delete(stage.Commands, command)
	delete(stage.Commands_mapString, command.Name)
	return command
}

// UnstageVoid removes command off the model stage
func (command *Command) UnstageVoid(stage *Stage) {
	delete(stage.Commands, command)
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

// Stage puts dummyagent to the model stage
func (dummyagent *DummyAgent) Stage(stage *Stage) *DummyAgent {

	if _, ok := stage.DummyAgents[dummyagent]; !ok {
		stage.DummyAgents[dummyagent] = __member
		stage.DummyAgentMap_Staged_Order[dummyagent] = stage.DummyAgentOrder
		stage.DummyAgentOrder++
	}
	stage.DummyAgents_mapString[dummyagent.Name] = dummyagent

	return dummyagent
}

// Unstage removes dummyagent off the model stage
func (dummyagent *DummyAgent) Unstage(stage *Stage) *DummyAgent {
	delete(stage.DummyAgents, dummyagent)
	delete(stage.DummyAgents_mapString, dummyagent.Name)
	return dummyagent
}

// UnstageVoid removes dummyagent off the model stage
func (dummyagent *DummyAgent) UnstageVoid(stage *Stage) {
	delete(stage.DummyAgents, dummyagent)
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

// Stage puts engine to the model stage
func (engine *Engine) Stage(stage *Stage) *Engine {

	if _, ok := stage.Engines[engine]; !ok {
		stage.Engines[engine] = __member
		stage.EngineMap_Staged_Order[engine] = stage.EngineOrder
		stage.EngineOrder++
	}
	stage.Engines_mapString[engine.Name] = engine

	return engine
}

// Unstage removes engine off the model stage
func (engine *Engine) Unstage(stage *Stage) *Engine {
	delete(stage.Engines, engine)
	delete(stage.Engines_mapString, engine.Name)
	return engine
}

// UnstageVoid removes engine off the model stage
func (engine *Engine) UnstageVoid(stage *Stage) {
	delete(stage.Engines, engine)
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

// Stage puts event to the model stage
func (event *Event) Stage(stage *Stage) *Event {

	if _, ok := stage.Events[event]; !ok {
		stage.Events[event] = __member
		stage.EventMap_Staged_Order[event] = stage.EventOrder
		stage.EventOrder++
	}
	stage.Events_mapString[event.Name] = event

	return event
}

// Unstage removes event off the model stage
func (event *Event) Unstage(stage *Stage) *Event {
	delete(stage.Events, event)
	delete(stage.Events_mapString, event.Name)
	return event
}

// UnstageVoid removes event off the model stage
func (event *Event) UnstageVoid(stage *Stage) {
	delete(stage.Events, event)
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

// Stage puts status to the model stage
func (status *Status) Stage(stage *Stage) *Status {

	if _, ok := stage.Statuss[status]; !ok {
		stage.Statuss[status] = __member
		stage.StatusMap_Staged_Order[status] = stage.StatusOrder
		stage.StatusOrder++
	}
	stage.Statuss_mapString[status.Name] = status

	return status
}

// Unstage removes status off the model stage
func (status *Status) Unstage(stage *Stage) *Status {
	delete(stage.Statuss, status)
	delete(stage.Statuss_mapString, status.Name)
	return status
}

// UnstageVoid removes status off the model stage
func (status *Status) UnstageVoid(stage *Stage) {
	delete(stage.Statuss, status)
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

// Stage puts updatestate to the model stage
func (updatestate *UpdateState) Stage(stage *Stage) *UpdateState {

	if _, ok := stage.UpdateStates[updatestate]; !ok {
		stage.UpdateStates[updatestate] = __member
		stage.UpdateStateMap_Staged_Order[updatestate] = stage.UpdateStateOrder
		stage.UpdateStateOrder++
	}
	stage.UpdateStates_mapString[updatestate.Name] = updatestate

	return updatestate
}

// Unstage removes updatestate off the model stage
func (updatestate *UpdateState) Unstage(stage *Stage) *UpdateState {
	delete(stage.UpdateStates, updatestate)
	delete(stage.UpdateStates_mapString, updatestate.Name)
	return updatestate
}

// UnstageVoid removes updatestate off the model stage
func (updatestate *UpdateState) UnstageVoid(stage *Stage) {
	delete(stage.UpdateStates, updatestate)
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
	stage.Commands = make(map[*Command]any)
	stage.Commands_mapString = make(map[string]*Command)
	stage.CommandMap_Staged_Order = make(map[*Command]uint)
	stage.CommandOrder = 0

	stage.DummyAgents = make(map[*DummyAgent]any)
	stage.DummyAgents_mapString = make(map[string]*DummyAgent)
	stage.DummyAgentMap_Staged_Order = make(map[*DummyAgent]uint)
	stage.DummyAgentOrder = 0

	stage.Engines = make(map[*Engine]any)
	stage.Engines_mapString = make(map[string]*Engine)
	stage.EngineMap_Staged_Order = make(map[*Engine]uint)
	stage.EngineOrder = 0

	stage.Events = make(map[*Event]any)
	stage.Events_mapString = make(map[string]*Event)
	stage.EventMap_Staged_Order = make(map[*Event]uint)
	stage.EventOrder = 0

	stage.Statuss = make(map[*Status]any)
	stage.Statuss_mapString = make(map[string]*Status)
	stage.StatusMap_Staged_Order = make(map[*Status]uint)
	stage.StatusOrder = 0

	stage.UpdateStates = make(map[*UpdateState]any)
	stage.UpdateStates_mapString = make(map[string]*UpdateState)
	stage.UpdateStateMap_Staged_Order = make(map[*UpdateState]uint)
	stage.UpdateStateOrder = 0

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
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*Stage)
	UnstageVoid(stage *Stage)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

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

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Command:
		return any(&stage.Commands_mapString).(*Type)
	case map[string]*DummyAgent:
		return any(&stage.DummyAgents_mapString).(*Type)
	case map[string]*Engine:
		return any(&stage.Engines_mapString).(*Type)
	case map[string]*Event:
		return any(&stage.Events_mapString).(*Type)
	case map[string]*Status:
		return any(&stage.Statuss_mapString).(*Type)
	case map[string]*UpdateState:
		return any(&stage.UpdateStates_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Command:
		return any(&stage.Commands).(*map[*Type]any)
	case DummyAgent:
		return any(&stage.DummyAgents).(*map[*Type]any)
	case Engine:
		return any(&stage.Engines).(*map[*Type]any)
	case Event:
		return any(&stage.Events).(*map[*Type]any)
	case Status:
		return any(&stage.Statuss).(*map[*Type]any)
	case UpdateState:
		return any(&stage.UpdateStates).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Command:
		return any(&stage.Commands).(*map[Type]any)
	case *DummyAgent:
		return any(&stage.DummyAgents).(*map[Type]any)
	case *Engine:
		return any(&stage.Engines).(*map[Type]any)
	case *Event:
		return any(&stage.Events).(*map[Type]any)
	case *Status:
		return any(&stage.Statuss).(*map[Type]any)
	case *UpdateState:
		return any(&stage.UpdateStates).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
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

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Command:
		res = "Command"
	case DummyAgent:
		res = "DummyAgent"
	case Engine:
		res = "Engine"
	case Event:
		res = "Event"
	case Status:
		res = "Status"
	case UpdateState:
		res = "UpdateState"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

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

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Command:
		res = []string{"Name", "Command", "CommandDate", "Engine"}
	case DummyAgent:
		res = []string{"TechName", "Name"}
	case Engine:
		res = []string{"Name", "EndTime", "CurrentTime", "DisplayFormat", "SecondsSinceStart", "Fired", "ControlMode", "State", "Speed"}
	case Event:
		res = []string{"Name", "Duration"}
	case Status:
		res = []string{"Name", "CurrentCommand", "CompletionDate", "CurrentSpeedCommand", "SpeedCommandCompletionDate"}
	case UpdateState:
		res = []string{"Name", "Duration", "Period"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Command:
		var rf ReverseField
		_ = rf
	case DummyAgent:
		var rf ReverseField
		_ = rf
	case Engine:
		var rf ReverseField
		_ = rf
	case Event:
		var rf ReverseField
		_ = rf
	case Status:
		var rf ReverseField
		_ = rf
	case UpdateState:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Command:
		res = []string{"Name", "Command", "CommandDate", "Engine"}
	case *DummyAgent:
		res = []string{"TechName", "Name"}
	case *Engine:
		res = []string{"Name", "EndTime", "CurrentTime", "DisplayFormat", "SecondsSinceStart", "Fired", "ControlMode", "State", "Speed"}
	case *Event:
		res = []string{"Name", "Duration"}
	case *Status:
		res = []string{"Name", "CurrentCommand", "CompletionDate", "CurrentSpeedCommand", "SpeedCommandCompletionDate"}
	case *UpdateState:
		res = []string{"Name", "Duration", "Period"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt    GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat  GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool   GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
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

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Command:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Command":
			enum := inferedInstance.Command
			res.valueString = enum.ToCodeString()
		case "CommandDate":
			res.valueString = inferedInstance.CommandDate
		case "Engine":
			if inferedInstance.Engine != nil {
				res.valueString = inferedInstance.Engine.Name
			}
		}
	case *DummyAgent:
		switch fieldName {
		// string value of fields
		case "TechName":
			res.valueString = inferedInstance.TechName
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case *Engine:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "EndTime":
			res.valueString = inferedInstance.EndTime
		case "CurrentTime":
			res.valueString = inferedInstance.CurrentTime
		case "DisplayFormat":
			res.valueString = inferedInstance.DisplayFormat
		case "SecondsSinceStart":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SecondsSinceStart)
			res.valueFloat = inferedInstance.SecondsSinceStart
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Fired":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Fired)
			res.valueInt = inferedInstance.Fired
			res.GongFieldValueType = GongFieldValueTypeInt
		case "ControlMode":
			enum := inferedInstance.ControlMode
			res.valueString = enum.ToCodeString()
		case "State":
			enum := inferedInstance.State
			res.valueString = enum.ToCodeString()
		case "Speed":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Speed)
			res.valueFloat = inferedInstance.Speed
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case *Event:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Duration":
			if math.Abs(inferedInstance.Duration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Duration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Duration.Hours()) % 24
				remainingMinutes := int(inferedInstance.Duration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Duration.Seconds()) % 60

				if inferedInstance.Duration.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.Duration.String())
			}
		}
	case *Status:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "CurrentCommand":
			enum := inferedInstance.CurrentCommand
			res.valueString = enum.ToCodeString()
		case "CompletionDate":
			res.valueString = inferedInstance.CompletionDate
		case "CurrentSpeedCommand":
			enum := inferedInstance.CurrentSpeedCommand
			res.valueString = enum.ToCodeString()
		case "SpeedCommandCompletionDate":
			res.valueString = inferedInstance.SpeedCommandCompletionDate
		}
	case *UpdateState:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Duration":
			if math.Abs(inferedInstance.Duration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Duration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Duration.Hours()) % 24
				remainingMinutes := int(inferedInstance.Duration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Duration.Seconds()) % 60

				if inferedInstance.Duration.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.Duration.String())
			}
		case "Period":
			if math.Abs(inferedInstance.Period.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Period.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Period.Hours()) % 24
				remainingMinutes := int(inferedInstance.Period.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Period.Seconds()) % 60

				if inferedInstance.Period.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.Period.String())
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Command:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Command":
			enum := inferedInstance.Command
			res.valueString = enum.ToCodeString()
		case "CommandDate":
			res.valueString = inferedInstance.CommandDate
		case "Engine":
			if inferedInstance.Engine != nil {
				res.valueString = inferedInstance.Engine.Name
			}
		}
	case DummyAgent:
		switch fieldName {
		// string value of fields
		case "TechName":
			res.valueString = inferedInstance.TechName
		case "Name":
			res.valueString = inferedInstance.Name
		}
	case Engine:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "EndTime":
			res.valueString = inferedInstance.EndTime
		case "CurrentTime":
			res.valueString = inferedInstance.CurrentTime
		case "DisplayFormat":
			res.valueString = inferedInstance.DisplayFormat
		case "SecondsSinceStart":
			res.valueString = fmt.Sprintf("%f", inferedInstance.SecondsSinceStart)
			res.valueFloat = inferedInstance.SecondsSinceStart
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Fired":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Fired)
			res.valueInt = inferedInstance.Fired
			res.GongFieldValueType = GongFieldValueTypeInt
		case "ControlMode":
			enum := inferedInstance.ControlMode
			res.valueString = enum.ToCodeString()
		case "State":
			enum := inferedInstance.State
			res.valueString = enum.ToCodeString()
		case "Speed":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Speed)
			res.valueFloat = inferedInstance.Speed
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	case Event:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Duration":
			if math.Abs(inferedInstance.Duration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Duration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Duration.Hours()) % 24
				remainingMinutes := int(inferedInstance.Duration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Duration.Seconds()) % 60

				if inferedInstance.Duration.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.Duration.String())
			}
		}
	case Status:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "CurrentCommand":
			enum := inferedInstance.CurrentCommand
			res.valueString = enum.ToCodeString()
		case "CompletionDate":
			res.valueString = inferedInstance.CompletionDate
		case "CurrentSpeedCommand":
			enum := inferedInstance.CurrentSpeedCommand
			res.valueString = enum.ToCodeString()
		case "SpeedCommandCompletionDate":
			res.valueString = inferedInstance.SpeedCommandCompletionDate
		}
	case UpdateState:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Duration":
			if math.Abs(inferedInstance.Duration.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Duration.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Duration.Hours()) % 24
				remainingMinutes := int(inferedInstance.Duration.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Duration.Seconds()) % 60

				if inferedInstance.Duration.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.Duration.String())
			}
		case "Period":
			if math.Abs(inferedInstance.Period.Hours()) >= 24 {
				days := __Gong__Abs(int(int(inferedInstance.Period.Hours()) / 24))
				months := int(days / 31)
				days = days - months*31

				remainingHours := int(inferedInstance.Period.Hours()) % 24
				remainingMinutes := int(inferedInstance.Period.Minutes()) % 60
				remainingSeconds := int(inferedInstance.Period.Seconds()) % 60

				if inferedInstance.Period.Hours() < 0 {
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
				res.valueString = fmt.Sprintf("%s\n", inferedInstance.Period.String())
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
