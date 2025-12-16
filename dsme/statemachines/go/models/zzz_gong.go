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

	statemachines_go "github.com/fullstack-lang/gong/dsme/statemachines/go"
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
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Actions           map[*Action]struct{}
	Actions_mapString map[string]*Action

	// insertion point for slice of pointers maps
	OnAfterActionCreateCallback OnAfterCreateInterface[Action]
	OnAfterActionUpdateCallback OnAfterUpdateInterface[Action]
	OnAfterActionDeleteCallback OnAfterDeleteInterface[Action]
	OnAfterActionReadCallback   OnAfterReadInterface[Action]

	Activitiess           map[*Activities]struct{}
	Activitiess_mapString map[string]*Activities

	// insertion point for slice of pointers maps
	OnAfterActivitiesCreateCallback OnAfterCreateInterface[Activities]
	OnAfterActivitiesUpdateCallback OnAfterUpdateInterface[Activities]
	OnAfterActivitiesDeleteCallback OnAfterDeleteInterface[Activities]
	OnAfterActivitiesReadCallback   OnAfterReadInterface[Activities]

	Architectures           map[*Architecture]struct{}
	Architectures_mapString map[string]*Architecture

	// insertion point for slice of pointers maps
	Architecture_StateMachines_reverseMap map[*StateMachine]*Architecture

	Architecture_Roles_reverseMap map[*Role]*Architecture

	OnAfterArchitectureCreateCallback OnAfterCreateInterface[Architecture]
	OnAfterArchitectureUpdateCallback OnAfterUpdateInterface[Architecture]
	OnAfterArchitectureDeleteCallback OnAfterDeleteInterface[Architecture]
	OnAfterArchitectureReadCallback   OnAfterReadInterface[Architecture]

	Diagrams           map[*Diagram]struct{}
	Diagrams_mapString map[string]*Diagram

	// insertion point for slice of pointers maps
	Diagram_State_Shapes_reverseMap map[*StateShape]*Diagram

	Diagram_Transition_Shapes_reverseMap map[*Transition_Shape]*Diagram

	OnAfterDiagramCreateCallback OnAfterCreateInterface[Diagram]
	OnAfterDiagramUpdateCallback OnAfterUpdateInterface[Diagram]
	OnAfterDiagramDeleteCallback OnAfterDeleteInterface[Diagram]
	OnAfterDiagramReadCallback   OnAfterReadInterface[Diagram]

	Guards           map[*Guard]struct{}
	Guards_mapString map[string]*Guard

	// insertion point for slice of pointers maps
	OnAfterGuardCreateCallback OnAfterCreateInterface[Guard]
	OnAfterGuardUpdateCallback OnAfterUpdateInterface[Guard]
	OnAfterGuardDeleteCallback OnAfterDeleteInterface[Guard]
	OnAfterGuardReadCallback   OnAfterReadInterface[Guard]

	Kills           map[*Kill]struct{}
	Kills_mapString map[string]*Kill

	// insertion point for slice of pointers maps
	OnAfterKillCreateCallback OnAfterCreateInterface[Kill]
	OnAfterKillUpdateCallback OnAfterUpdateInterface[Kill]
	OnAfterKillDeleteCallback OnAfterDeleteInterface[Kill]
	OnAfterKillReadCallback   OnAfterReadInterface[Kill]

	Messages           map[*Message]struct{}
	Messages_mapString map[string]*Message

	// insertion point for slice of pointers maps
	OnAfterMessageCreateCallback OnAfterCreateInterface[Message]
	OnAfterMessageUpdateCallback OnAfterUpdateInterface[Message]
	OnAfterMessageDeleteCallback OnAfterDeleteInterface[Message]
	OnAfterMessageReadCallback   OnAfterReadInterface[Message]

	MessageTypes           map[*MessageType]struct{}
	MessageTypes_mapString map[string]*MessageType

	// insertion point for slice of pointers maps
	OnAfterMessageTypeCreateCallback OnAfterCreateInterface[MessageType]
	OnAfterMessageTypeUpdateCallback OnAfterUpdateInterface[MessageType]
	OnAfterMessageTypeDeleteCallback OnAfterDeleteInterface[MessageType]
	OnAfterMessageTypeReadCallback   OnAfterReadInterface[MessageType]

	Objects           map[*Object]struct{}
	Objects_mapString map[string]*Object

	// insertion point for slice of pointers maps
	Object_Messages_reverseMap map[*Message]*Object

	OnAfterObjectCreateCallback OnAfterCreateInterface[Object]
	OnAfterObjectUpdateCallback OnAfterUpdateInterface[Object]
	OnAfterObjectDeleteCallback OnAfterDeleteInterface[Object]
	OnAfterObjectReadCallback   OnAfterReadInterface[Object]

	Roles           map[*Role]struct{}
	Roles_mapString map[string]*Role

	// insertion point for slice of pointers maps
	Role_RolesWithSamePermissions_reverseMap map[*Role]*Role

	OnAfterRoleCreateCallback OnAfterCreateInterface[Role]
	OnAfterRoleUpdateCallback OnAfterUpdateInterface[Role]
	OnAfterRoleDeleteCallback OnAfterDeleteInterface[Role]
	OnAfterRoleReadCallback   OnAfterReadInterface[Role]

	States           map[*State]struct{}
	States_mapString map[string]*State

	// insertion point for slice of pointers maps
	State_SubStates_reverseMap map[*State]*State

	State_Diagrams_reverseMap map[*Diagram]*State

	State_Activities_reverseMap map[*Activities]*State

	OnAfterStateCreateCallback OnAfterCreateInterface[State]
	OnAfterStateUpdateCallback OnAfterUpdateInterface[State]
	OnAfterStateDeleteCallback OnAfterDeleteInterface[State]
	OnAfterStateReadCallback   OnAfterReadInterface[State]

	StateMachines           map[*StateMachine]struct{}
	StateMachines_mapString map[string]*StateMachine

	// insertion point for slice of pointers maps
	StateMachine_States_reverseMap map[*State]*StateMachine

	StateMachine_Diagrams_reverseMap map[*Diagram]*StateMachine

	OnAfterStateMachineCreateCallback OnAfterCreateInterface[StateMachine]
	OnAfterStateMachineUpdateCallback OnAfterUpdateInterface[StateMachine]
	OnAfterStateMachineDeleteCallback OnAfterDeleteInterface[StateMachine]
	OnAfterStateMachineReadCallback   OnAfterReadInterface[StateMachine]

	StateShapes           map[*StateShape]struct{}
	StateShapes_mapString map[string]*StateShape

	// insertion point for slice of pointers maps
	OnAfterStateShapeCreateCallback OnAfterCreateInterface[StateShape]
	OnAfterStateShapeUpdateCallback OnAfterUpdateInterface[StateShape]
	OnAfterStateShapeDeleteCallback OnAfterDeleteInterface[StateShape]
	OnAfterStateShapeReadCallback   OnAfterReadInterface[StateShape]

	Transitions           map[*Transition]struct{}
	Transitions_mapString map[string]*Transition

	// insertion point for slice of pointers maps
	Transition_RolesWithPermissions_reverseMap map[*Role]*Transition

	Transition_GeneratedMessages_reverseMap map[*MessageType]*Transition

	Transition_Diagrams_reverseMap map[*Diagram]*Transition

	OnAfterTransitionCreateCallback OnAfterCreateInterface[Transition]
	OnAfterTransitionUpdateCallback OnAfterUpdateInterface[Transition]
	OnAfterTransitionDeleteCallback OnAfterDeleteInterface[Transition]
	OnAfterTransitionReadCallback   OnAfterReadInterface[Transition]

	Transition_Shapes           map[*Transition_Shape]struct{}
	Transition_Shapes_mapString map[string]*Transition_Shape

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
	ActionOrder            uint
	ActionMap_Staged_Order map[*Action]uint

	ActivitiesOrder            uint
	ActivitiesMap_Staged_Order map[*Activities]uint

	ArchitectureOrder            uint
	ArchitectureMap_Staged_Order map[*Architecture]uint

	DiagramOrder            uint
	DiagramMap_Staged_Order map[*Diagram]uint

	GuardOrder            uint
	GuardMap_Staged_Order map[*Guard]uint

	KillOrder            uint
	KillMap_Staged_Order map[*Kill]uint

	MessageOrder            uint
	MessageMap_Staged_Order map[*Message]uint

	MessageTypeOrder            uint
	MessageTypeMap_Staged_Order map[*MessageType]uint

	ObjectOrder            uint
	ObjectMap_Staged_Order map[*Object]uint

	RoleOrder            uint
	RoleMap_Staged_Order map[*Role]uint

	StateOrder            uint
	StateMap_Staged_Order map[*State]uint

	StateMachineOrder            uint
	StateMachineMap_Staged_Order map[*StateMachine]uint

	StateShapeOrder            uint
	StateShapeMap_Staged_Order map[*StateShape]uint

	TransitionOrder            uint
	TransitionMap_Staged_Order map[*Transition]uint

	Transition_ShapeOrder            uint
	Transition_ShapeMap_Staged_Order map[*Transition_Shape]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// for the computation of the diff at each commit we need
	reference map[GongstructIF]GongstructIF
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func (stage *Stage) GetReference() map[GongstructIF]GongstructIF {
	return stage.reference
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
	case *Action:
		tmp := GetStructInstancesByOrder(stage.Actions, stage.ActionMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.Activitiess, stage.ActivitiesMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.Architectures, stage.ArchitectureMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.Diagrams, stage.DiagramMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.Guards, stage.GuardMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.Kills, stage.KillMap_Staged_Order)

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
	case *Message:
		tmp := GetStructInstancesByOrder(stage.Messages, stage.MessageMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.MessageTypes, stage.MessageTypeMap_Staged_Order)

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
	case *Object:
		tmp := GetStructInstancesByOrder(stage.Objects, stage.ObjectMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.Roles, stage.RoleMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.States, stage.StateMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.StateMachines, stage.StateMachineMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.StateShapes, stage.StateShapeMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.Transitions, stage.TransitionMap_Staged_Order)

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
		tmp := GetStructInstancesByOrder(stage.Transition_Shapes, stage.Transition_ShapeMap_Staged_Order)

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
		res = GetNamedStructInstances(stage.Actions, stage.ActionMap_Staged_Order)
	case "Activities":
		res = GetNamedStructInstances(stage.Activitiess, stage.ActivitiesMap_Staged_Order)
	case "Architecture":
		res = GetNamedStructInstances(stage.Architectures, stage.ArchitectureMap_Staged_Order)
	case "Diagram":
		res = GetNamedStructInstances(stage.Diagrams, stage.DiagramMap_Staged_Order)
	case "Guard":
		res = GetNamedStructInstances(stage.Guards, stage.GuardMap_Staged_Order)
	case "Kill":
		res = GetNamedStructInstances(stage.Kills, stage.KillMap_Staged_Order)
	case "Message":
		res = GetNamedStructInstances(stage.Messages, stage.MessageMap_Staged_Order)
	case "MessageType":
		res = GetNamedStructInstances(stage.MessageTypes, stage.MessageTypeMap_Staged_Order)
	case "Object":
		res = GetNamedStructInstances(stage.Objects, stage.ObjectMap_Staged_Order)
	case "Role":
		res = GetNamedStructInstances(stage.Roles, stage.RoleMap_Staged_Order)
	case "State":
		res = GetNamedStructInstances(stage.States, stage.StateMap_Staged_Order)
	case "StateMachine":
		res = GetNamedStructInstances(stage.StateMachines, stage.StateMachineMap_Staged_Order)
	case "StateShape":
		res = GetNamedStructInstances(stage.StateShapes, stage.StateShapeMap_Staged_Order)
	case "Transition":
		res = GetNamedStructInstances(stage.Transitions, stage.TransitionMap_Staged_Order)
	case "Transition_Shape":
		res = GetNamedStructInstances(stage.Transition_Shapes, stage.Transition_ShapeMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/dsme/statemachines/go/models"
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
	CommitMessage(message *Message)
	CheckoutMessage(message *Message)
	CommitMessageType(messagetype *MessageType)
	CheckoutMessageType(messagetype *MessageType)
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

		Messages:           make(map[*Message]struct{}),
		Messages_mapString: make(map[string]*Message),

		MessageTypes:           make(map[*MessageType]struct{}),
		MessageTypes_mapString: make(map[string]*MessageType),

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
		ActionMap_Staged_Order: make(map[*Action]uint),

		ActivitiesMap_Staged_Order: make(map[*Activities]uint),

		ArchitectureMap_Staged_Order: make(map[*Architecture]uint),

		DiagramMap_Staged_Order: make(map[*Diagram]uint),

		GuardMap_Staged_Order: make(map[*Guard]uint),

		KillMap_Staged_Order: make(map[*Kill]uint),

		MessageMap_Staged_Order: make(map[*Message]uint),

		MessageTypeMap_Staged_Order: make(map[*MessageType]uint),

		ObjectMap_Staged_Order: make(map[*Object]uint),

		RoleMap_Staged_Order: make(map[*Role]uint),

		StateMap_Staged_Order: make(map[*State]uint),

		StateMachineMap_Staged_Order: make(map[*StateMachine]uint),

		StateShapeMap_Staged_Order: make(map[*StateShape]uint),

		TransitionMap_Staged_Order: make(map[*Transition]uint),

		Transition_ShapeMap_Staged_Order: make(map[*Transition_Shape]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Action"},
			{name: "Activities"},
			{name: "Architecture"},
			{name: "Diagram"},
			{name: "Guard"},
			{name: "Kill"},
			{name: "Message"},
			{name: "MessageType"},
			{name: "Object"},
			{name: "Role"},
			{name: "State"},
			{name: "StateMachine"},
			{name: "StateShape"},
			{name: "Transition"},
			{name: "Transition_Shape"},
		}, // end of insertion point

		reference: make(map[GongstructIF]GongstructIF),
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Action:
		return stage.ActionMap_Staged_Order[instance]
	case *Activities:
		return stage.ActivitiesMap_Staged_Order[instance]
	case *Architecture:
		return stage.ArchitectureMap_Staged_Order[instance]
	case *Diagram:
		return stage.DiagramMap_Staged_Order[instance]
	case *Guard:
		return stage.GuardMap_Staged_Order[instance]
	case *Kill:
		return stage.KillMap_Staged_Order[instance]
	case *Message:
		return stage.MessageMap_Staged_Order[instance]
	case *MessageType:
		return stage.MessageTypeMap_Staged_Order[instance]
	case *Object:
		return stage.ObjectMap_Staged_Order[instance]
	case *Role:
		return stage.RoleMap_Staged_Order[instance]
	case *State:
		return stage.StateMap_Staged_Order[instance]
	case *StateMachine:
		return stage.StateMachineMap_Staged_Order[instance]
	case *StateShape:
		return stage.StateShapeMap_Staged_Order[instance]
	case *Transition:
		return stage.TransitionMap_Staged_Order[instance]
	case *Transition_Shape:
		return stage.Transition_ShapeMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Action:
		return stage.ActionMap_Staged_Order[instance]
	case *Activities:
		return stage.ActivitiesMap_Staged_Order[instance]
	case *Architecture:
		return stage.ArchitectureMap_Staged_Order[instance]
	case *Diagram:
		return stage.DiagramMap_Staged_Order[instance]
	case *Guard:
		return stage.GuardMap_Staged_Order[instance]
	case *Kill:
		return stage.KillMap_Staged_Order[instance]
	case *Message:
		return stage.MessageMap_Staged_Order[instance]
	case *MessageType:
		return stage.MessageTypeMap_Staged_Order[instance]
	case *Object:
		return stage.ObjectMap_Staged_Order[instance]
	case *Role:
		return stage.RoleMap_Staged_Order[instance]
	case *State:
		return stage.StateMap_Staged_Order[instance]
	case *StateMachine:
		return stage.StateMachineMap_Staged_Order[instance]
	case *StateShape:
		return stage.StateShapeMap_Staged_Order[instance]
	case *Transition:
		return stage.TransitionMap_Staged_Order[instance]
	case *Transition_Shape:
		return stage.Transition_ShapeMap_Staged_Order[instance]
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
	stage.ComputeReference()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Action"] = len(stage.Actions)
	stage.Map_GongStructName_InstancesNb["Activities"] = len(stage.Activitiess)
	stage.Map_GongStructName_InstancesNb["Architecture"] = len(stage.Architectures)
	stage.Map_GongStructName_InstancesNb["Diagram"] = len(stage.Diagrams)
	stage.Map_GongStructName_InstancesNb["Guard"] = len(stage.Guards)
	stage.Map_GongStructName_InstancesNb["Kill"] = len(stage.Kills)
	stage.Map_GongStructName_InstancesNb["Message"] = len(stage.Messages)
	stage.Map_GongStructName_InstancesNb["MessageType"] = len(stage.MessageTypes)
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
		stage.ActionMap_Staged_Order[action] = stage.ActionOrder
		stage.ActionOrder++
	}
	stage.Actions_mapString[action.Name] = action

	return action
}

// StageForceOrder puts action to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActionOrder
// - update stage.ActionOrder accordingly
func (action *Action) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Actions[action]; !ok {
		stage.Actions[action] = struct{}{}

		if order > stage.ActionOrder {
			stage.ActionOrder = order
		}
		stage.ActionMap_Staged_Order[action] = stage.ActionOrder
		stage.ActionOrder++
	}
	stage.Actions_mapString[action.Name] = action
}

// Unstage removes action off the model stage
func (action *Action) Unstage(stage *Stage) *Action {
	delete(stage.Actions, action)
	delete(stage.Actions_mapString, action.Name)

	return action
}

// UnstageVoid removes action off the model stage
func (action *Action) UnstageVoid(stage *Stage) {
	delete(stage.Actions, action)
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
func (action *Action) SetName(name string) (){
	action.Name = name
}

// Stage puts activities to the model stage
func (activities *Activities) Stage(stage *Stage) *Activities {

	if _, ok := stage.Activitiess[activities]; !ok {
		stage.Activitiess[activities] = struct{}{}
		stage.ActivitiesMap_Staged_Order[activities] = stage.ActivitiesOrder
		stage.ActivitiesOrder++
	}
	stage.Activitiess_mapString[activities.Name] = activities

	return activities
}

// StageForceOrder puts activities to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ActivitiesOrder
// - update stage.ActivitiesOrder accordingly
func (activities *Activities) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Activitiess[activities]; !ok {
		stage.Activitiess[activities] = struct{}{}

		if order > stage.ActivitiesOrder {
			stage.ActivitiesOrder = order
		}
		stage.ActivitiesMap_Staged_Order[activities] = stage.ActivitiesOrder
		stage.ActivitiesOrder++
	}
	stage.Activitiess_mapString[activities.Name] = activities
}

// Unstage removes activities off the model stage
func (activities *Activities) Unstage(stage *Stage) *Activities {
	delete(stage.Activitiess, activities)
	delete(stage.Activitiess_mapString, activities.Name)

	return activities
}

// UnstageVoid removes activities off the model stage
func (activities *Activities) UnstageVoid(stage *Stage) {
	delete(stage.Activitiess, activities)
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
func (activities *Activities) SetName(name string) (){
	activities.Name = name
}

// Stage puts architecture to the model stage
func (architecture *Architecture) Stage(stage *Stage) *Architecture {

	if _, ok := stage.Architectures[architecture]; !ok {
		stage.Architectures[architecture] = struct{}{}
		stage.ArchitectureMap_Staged_Order[architecture] = stage.ArchitectureOrder
		stage.ArchitectureOrder++
	}
	stage.Architectures_mapString[architecture.Name] = architecture

	return architecture
}

// StageForceOrder puts architecture to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ArchitectureOrder
// - update stage.ArchitectureOrder accordingly
func (architecture *Architecture) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Architectures[architecture]; !ok {
		stage.Architectures[architecture] = struct{}{}

		if order > stage.ArchitectureOrder {
			stage.ArchitectureOrder = order
		}
		stage.ArchitectureMap_Staged_Order[architecture] = stage.ArchitectureOrder
		stage.ArchitectureOrder++
	}
	stage.Architectures_mapString[architecture.Name] = architecture
}

// Unstage removes architecture off the model stage
func (architecture *Architecture) Unstage(stage *Stage) *Architecture {
	delete(stage.Architectures, architecture)
	delete(stage.Architectures_mapString, architecture.Name)

	return architecture
}

// UnstageVoid removes architecture off the model stage
func (architecture *Architecture) UnstageVoid(stage *Stage) {
	delete(stage.Architectures, architecture)
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
func (architecture *Architecture) SetName(name string) (){
	architecture.Name = name
}

// Stage puts diagram to the model stage
func (diagram *Diagram) Stage(stage *Stage) *Diagram {

	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}
		stage.DiagramMap_Staged_Order[diagram] = stage.DiagramOrder
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram

	return diagram
}

// StageForceOrder puts diagram to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DiagramOrder
// - update stage.DiagramOrder accordingly
func (diagram *Diagram) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Diagrams[diagram]; !ok {
		stage.Diagrams[diagram] = struct{}{}

		if order > stage.DiagramOrder {
			stage.DiagramOrder = order
		}
		stage.DiagramMap_Staged_Order[diagram] = stage.DiagramOrder
		stage.DiagramOrder++
	}
	stage.Diagrams_mapString[diagram.Name] = diagram
}

// Unstage removes diagram off the model stage
func (diagram *Diagram) Unstage(stage *Stage) *Diagram {
	delete(stage.Diagrams, diagram)
	delete(stage.Diagrams_mapString, diagram.Name)

	return diagram
}

// UnstageVoid removes diagram off the model stage
func (diagram *Diagram) UnstageVoid(stage *Stage) {
	delete(stage.Diagrams, diagram)
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
func (diagram *Diagram) SetName(name string) (){
	diagram.Name = name
}

// Stage puts guard to the model stage
func (guard *Guard) Stage(stage *Stage) *Guard {

	if _, ok := stage.Guards[guard]; !ok {
		stage.Guards[guard] = struct{}{}
		stage.GuardMap_Staged_Order[guard] = stage.GuardOrder
		stage.GuardOrder++
	}
	stage.Guards_mapString[guard.Name] = guard

	return guard
}

// StageForceOrder puts guard to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GuardOrder
// - update stage.GuardOrder accordingly
func (guard *Guard) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Guards[guard]; !ok {
		stage.Guards[guard] = struct{}{}

		if order > stage.GuardOrder {
			stage.GuardOrder = order
		}
		stage.GuardMap_Staged_Order[guard] = stage.GuardOrder
		stage.GuardOrder++
	}
	stage.Guards_mapString[guard.Name] = guard
}

// Unstage removes guard off the model stage
func (guard *Guard) Unstage(stage *Stage) *Guard {
	delete(stage.Guards, guard)
	delete(stage.Guards_mapString, guard.Name)

	return guard
}

// UnstageVoid removes guard off the model stage
func (guard *Guard) UnstageVoid(stage *Stage) {
	delete(stage.Guards, guard)
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
func (guard *Guard) SetName(name string) (){
	guard.Name = name
}

// Stage puts kill to the model stage
func (kill *Kill) Stage(stage *Stage) *Kill {

	if _, ok := stage.Kills[kill]; !ok {
		stage.Kills[kill] = struct{}{}
		stage.KillMap_Staged_Order[kill] = stage.KillOrder
		stage.KillOrder++
	}
	stage.Kills_mapString[kill.Name] = kill

	return kill
}

// StageForceOrder puts kill to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.KillOrder
// - update stage.KillOrder accordingly
func (kill *Kill) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Kills[kill]; !ok {
		stage.Kills[kill] = struct{}{}

		if order > stage.KillOrder {
			stage.KillOrder = order
		}
		stage.KillMap_Staged_Order[kill] = stage.KillOrder
		stage.KillOrder++
	}
	stage.Kills_mapString[kill.Name] = kill
}

// Unstage removes kill off the model stage
func (kill *Kill) Unstage(stage *Stage) *Kill {
	delete(stage.Kills, kill)
	delete(stage.Kills_mapString, kill.Name)

	return kill
}

// UnstageVoid removes kill off the model stage
func (kill *Kill) UnstageVoid(stage *Stage) {
	delete(stage.Kills, kill)
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
func (kill *Kill) SetName(name string) (){
	kill.Name = name
}

// Stage puts message to the model stage
func (message *Message) Stage(stage *Stage) *Message {

	if _, ok := stage.Messages[message]; !ok {
		stage.Messages[message] = struct{}{}
		stage.MessageMap_Staged_Order[message] = stage.MessageOrder
		stage.MessageOrder++
	}
	stage.Messages_mapString[message.Name] = message

	return message
}

// StageForceOrder puts message to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MessageOrder
// - update stage.MessageOrder accordingly
func (message *Message) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Messages[message]; !ok {
		stage.Messages[message] = struct{}{}

		if order > stage.MessageOrder {
			stage.MessageOrder = order
		}
		stage.MessageMap_Staged_Order[message] = stage.MessageOrder
		stage.MessageOrder++
	}
	stage.Messages_mapString[message.Name] = message
}

// Unstage removes message off the model stage
func (message *Message) Unstage(stage *Stage) *Message {
	delete(stage.Messages, message)
	delete(stage.Messages_mapString, message.Name)

	return message
}

// UnstageVoid removes message off the model stage
func (message *Message) UnstageVoid(stage *Stage) {
	delete(stage.Messages, message)
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
func (message *Message) SetName(name string) (){
	message.Name = name
}

// Stage puts messagetype to the model stage
func (messagetype *MessageType) Stage(stage *Stage) *MessageType {

	if _, ok := stage.MessageTypes[messagetype]; !ok {
		stage.MessageTypes[messagetype] = struct{}{}
		stage.MessageTypeMap_Staged_Order[messagetype] = stage.MessageTypeOrder
		stage.MessageTypeOrder++
	}
	stage.MessageTypes_mapString[messagetype.Name] = messagetype

	return messagetype
}

// StageForceOrder puts messagetype to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.MessageTypeOrder
// - update stage.MessageTypeOrder accordingly
func (messagetype *MessageType) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.MessageTypes[messagetype]; !ok {
		stage.MessageTypes[messagetype] = struct{}{}

		if order > stage.MessageTypeOrder {
			stage.MessageTypeOrder = order
		}
		stage.MessageTypeMap_Staged_Order[messagetype] = stage.MessageTypeOrder
		stage.MessageTypeOrder++
	}
	stage.MessageTypes_mapString[messagetype.Name] = messagetype
}

// Unstage removes messagetype off the model stage
func (messagetype *MessageType) Unstage(stage *Stage) *MessageType {
	delete(stage.MessageTypes, messagetype)
	delete(stage.MessageTypes_mapString, messagetype.Name)

	return messagetype
}

// UnstageVoid removes messagetype off the model stage
func (messagetype *MessageType) UnstageVoid(stage *Stage) {
	delete(stage.MessageTypes, messagetype)
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
func (messagetype *MessageType) SetName(name string) (){
	messagetype.Name = name
}

// Stage puts object to the model stage
func (object *Object) Stage(stage *Stage) *Object {

	if _, ok := stage.Objects[object]; !ok {
		stage.Objects[object] = struct{}{}
		stage.ObjectMap_Staged_Order[object] = stage.ObjectOrder
		stage.ObjectOrder++
	}
	stage.Objects_mapString[object.Name] = object

	return object
}

// StageForceOrder puts object to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ObjectOrder
// - update stage.ObjectOrder accordingly
func (object *Object) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Objects[object]; !ok {
		stage.Objects[object] = struct{}{}

		if order > stage.ObjectOrder {
			stage.ObjectOrder = order
		}
		stage.ObjectMap_Staged_Order[object] = stage.ObjectOrder
		stage.ObjectOrder++
	}
	stage.Objects_mapString[object.Name] = object
}

// Unstage removes object off the model stage
func (object *Object) Unstage(stage *Stage) *Object {
	delete(stage.Objects, object)
	delete(stage.Objects_mapString, object.Name)

	return object
}

// UnstageVoid removes object off the model stage
func (object *Object) UnstageVoid(stage *Stage) {
	delete(stage.Objects, object)
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
func (object *Object) SetName(name string) (){
	object.Name = name
}

// Stage puts role to the model stage
func (role *Role) Stage(stage *Stage) *Role {

	if _, ok := stage.Roles[role]; !ok {
		stage.Roles[role] = struct{}{}
		stage.RoleMap_Staged_Order[role] = stage.RoleOrder
		stage.RoleOrder++
	}
	stage.Roles_mapString[role.Name] = role

	return role
}

// StageForceOrder puts role to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RoleOrder
// - update stage.RoleOrder accordingly
func (role *Role) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Roles[role]; !ok {
		stage.Roles[role] = struct{}{}

		if order > stage.RoleOrder {
			stage.RoleOrder = order
		}
		stage.RoleMap_Staged_Order[role] = stage.RoleOrder
		stage.RoleOrder++
	}
	stage.Roles_mapString[role.Name] = role
}

// Unstage removes role off the model stage
func (role *Role) Unstage(stage *Stage) *Role {
	delete(stage.Roles, role)
	delete(stage.Roles_mapString, role.Name)

	return role
}

// UnstageVoid removes role off the model stage
func (role *Role) UnstageVoid(stage *Stage) {
	delete(stage.Roles, role)
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
func (role *Role) SetName(name string) (){
	role.Name = name
}

// Stage puts state to the model stage
func (state *State) Stage(stage *Stage) *State {

	if _, ok := stage.States[state]; !ok {
		stage.States[state] = struct{}{}
		stage.StateMap_Staged_Order[state] = stage.StateOrder
		stage.StateOrder++
	}
	stage.States_mapString[state.Name] = state

	return state
}

// StageForceOrder puts state to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateOrder
// - update stage.StateOrder accordingly
func (state *State) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.States[state]; !ok {
		stage.States[state] = struct{}{}

		if order > stage.StateOrder {
			stage.StateOrder = order
		}
		stage.StateMap_Staged_Order[state] = stage.StateOrder
		stage.StateOrder++
	}
	stage.States_mapString[state.Name] = state
}

// Unstage removes state off the model stage
func (state *State) Unstage(stage *Stage) *State {
	delete(stage.States, state)
	delete(stage.States_mapString, state.Name)

	return state
}

// UnstageVoid removes state off the model stage
func (state *State) UnstageVoid(stage *Stage) {
	delete(stage.States, state)
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
func (state *State) SetName(name string) (){
	state.Name = name
}

// Stage puts statemachine to the model stage
func (statemachine *StateMachine) Stage(stage *Stage) *StateMachine {

	if _, ok := stage.StateMachines[statemachine]; !ok {
		stage.StateMachines[statemachine] = struct{}{}
		stage.StateMachineMap_Staged_Order[statemachine] = stage.StateMachineOrder
		stage.StateMachineOrder++
	}
	stage.StateMachines_mapString[statemachine.Name] = statemachine

	return statemachine
}

// StageForceOrder puts statemachine to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateMachineOrder
// - update stage.StateMachineOrder accordingly
func (statemachine *StateMachine) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.StateMachines[statemachine]; !ok {
		stage.StateMachines[statemachine] = struct{}{}

		if order > stage.StateMachineOrder {
			stage.StateMachineOrder = order
		}
		stage.StateMachineMap_Staged_Order[statemachine] = stage.StateMachineOrder
		stage.StateMachineOrder++
	}
	stage.StateMachines_mapString[statemachine.Name] = statemachine
}

// Unstage removes statemachine off the model stage
func (statemachine *StateMachine) Unstage(stage *Stage) *StateMachine {
	delete(stage.StateMachines, statemachine)
	delete(stage.StateMachines_mapString, statemachine.Name)

	return statemachine
}

// UnstageVoid removes statemachine off the model stage
func (statemachine *StateMachine) UnstageVoid(stage *Stage) {
	delete(stage.StateMachines, statemachine)
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
func (statemachine *StateMachine) SetName(name string) (){
	statemachine.Name = name
}

// Stage puts stateshape to the model stage
func (stateshape *StateShape) Stage(stage *Stage) *StateShape {

	if _, ok := stage.StateShapes[stateshape]; !ok {
		stage.StateShapes[stateshape] = struct{}{}
		stage.StateShapeMap_Staged_Order[stateshape] = stage.StateShapeOrder
		stage.StateShapeOrder++
	}
	stage.StateShapes_mapString[stateshape.Name] = stateshape

	return stateshape
}

// StageForceOrder puts stateshape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.StateShapeOrder
// - update stage.StateShapeOrder accordingly
func (stateshape *StateShape) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.StateShapes[stateshape]; !ok {
		stage.StateShapes[stateshape] = struct{}{}

		if order > stage.StateShapeOrder {
			stage.StateShapeOrder = order
		}
		stage.StateShapeMap_Staged_Order[stateshape] = stage.StateShapeOrder
		stage.StateShapeOrder++
	}
	stage.StateShapes_mapString[stateshape.Name] = stateshape
}

// Unstage removes stateshape off the model stage
func (stateshape *StateShape) Unstage(stage *Stage) *StateShape {
	delete(stage.StateShapes, stateshape)
	delete(stage.StateShapes_mapString, stateshape.Name)

	return stateshape
}

// UnstageVoid removes stateshape off the model stage
func (stateshape *StateShape) UnstageVoid(stage *Stage) {
	delete(stage.StateShapes, stateshape)
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
func (stateshape *StateShape) SetName(name string) (){
	stateshape.Name = name
}

// Stage puts transition to the model stage
func (transition *Transition) Stage(stage *Stage) *Transition {

	if _, ok := stage.Transitions[transition]; !ok {
		stage.Transitions[transition] = struct{}{}
		stage.TransitionMap_Staged_Order[transition] = stage.TransitionOrder
		stage.TransitionOrder++
	}
	stage.Transitions_mapString[transition.Name] = transition

	return transition
}

// StageForceOrder puts transition to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TransitionOrder
// - update stage.TransitionOrder accordingly
func (transition *Transition) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Transitions[transition]; !ok {
		stage.Transitions[transition] = struct{}{}

		if order > stage.TransitionOrder {
			stage.TransitionOrder = order
		}
		stage.TransitionMap_Staged_Order[transition] = stage.TransitionOrder
		stage.TransitionOrder++
	}
	stage.Transitions_mapString[transition.Name] = transition
}

// Unstage removes transition off the model stage
func (transition *Transition) Unstage(stage *Stage) *Transition {
	delete(stage.Transitions, transition)
	delete(stage.Transitions_mapString, transition.Name)

	return transition
}

// UnstageVoid removes transition off the model stage
func (transition *Transition) UnstageVoid(stage *Stage) {
	delete(stage.Transitions, transition)
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
func (transition *Transition) SetName(name string) (){
	transition.Name = name
}

// Stage puts transition_shape to the model stage
func (transition_shape *Transition_Shape) Stage(stage *Stage) *Transition_Shape {

	if _, ok := stage.Transition_Shapes[transition_shape]; !ok {
		stage.Transition_Shapes[transition_shape] = struct{}{}
		stage.Transition_ShapeMap_Staged_Order[transition_shape] = stage.Transition_ShapeOrder
		stage.Transition_ShapeOrder++
	}
	stage.Transition_Shapes_mapString[transition_shape.Name] = transition_shape

	return transition_shape
}

// StageForceOrder puts transition_shape to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.Transition_ShapeOrder
// - update stage.Transition_ShapeOrder accordingly
func (transition_shape *Transition_Shape) StageForceOrder(stage *Stage, order uint) {

	if _, ok := stage.Transition_Shapes[transition_shape]; !ok {
		stage.Transition_Shapes[transition_shape] = struct{}{}

		if order > stage.Transition_ShapeOrder {
			stage.Transition_ShapeOrder = order
		}
		stage.Transition_ShapeMap_Staged_Order[transition_shape] = stage.Transition_ShapeOrder
		stage.Transition_ShapeOrder++
	}
	stage.Transition_Shapes_mapString[transition_shape.Name] = transition_shape
}

// Unstage removes transition_shape off the model stage
func (transition_shape *Transition_Shape) Unstage(stage *Stage) *Transition_Shape {
	delete(stage.Transition_Shapes, transition_shape)
	delete(stage.Transition_Shapes_mapString, transition_shape.Name)

	return transition_shape
}

// UnstageVoid removes transition_shape off the model stage
func (transition_shape *Transition_Shape) UnstageVoid(stage *Stage) {
	delete(stage.Transition_Shapes, transition_shape)
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
func (transition_shape *Transition_Shape) SetName(name string) (){
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
	CreateORMMessage(Message *Message)
	CreateORMMessageType(MessageType *MessageType)
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
	DeleteORMMessage(Message *Message)
	DeleteORMMessageType(MessageType *MessageType)
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
	stage.ActionMap_Staged_Order = make(map[*Action]uint)
	stage.ActionOrder = 0

	stage.Activitiess = make(map[*Activities]struct{})
	stage.Activitiess_mapString = make(map[string]*Activities)
	stage.ActivitiesMap_Staged_Order = make(map[*Activities]uint)
	stage.ActivitiesOrder = 0

	stage.Architectures = make(map[*Architecture]struct{})
	stage.Architectures_mapString = make(map[string]*Architecture)
	stage.ArchitectureMap_Staged_Order = make(map[*Architecture]uint)
	stage.ArchitectureOrder = 0

	stage.Diagrams = make(map[*Diagram]struct{})
	stage.Diagrams_mapString = make(map[string]*Diagram)
	stage.DiagramMap_Staged_Order = make(map[*Diagram]uint)
	stage.DiagramOrder = 0

	stage.Guards = make(map[*Guard]struct{})
	stage.Guards_mapString = make(map[string]*Guard)
	stage.GuardMap_Staged_Order = make(map[*Guard]uint)
	stage.GuardOrder = 0

	stage.Kills = make(map[*Kill]struct{})
	stage.Kills_mapString = make(map[string]*Kill)
	stage.KillMap_Staged_Order = make(map[*Kill]uint)
	stage.KillOrder = 0

	stage.Messages = make(map[*Message]struct{})
	stage.Messages_mapString = make(map[string]*Message)
	stage.MessageMap_Staged_Order = make(map[*Message]uint)
	stage.MessageOrder = 0

	stage.MessageTypes = make(map[*MessageType]struct{})
	stage.MessageTypes_mapString = make(map[string]*MessageType)
	stage.MessageTypeMap_Staged_Order = make(map[*MessageType]uint)
	stage.MessageTypeOrder = 0

	stage.Objects = make(map[*Object]struct{})
	stage.Objects_mapString = make(map[string]*Object)
	stage.ObjectMap_Staged_Order = make(map[*Object]uint)
	stage.ObjectOrder = 0

	stage.Roles = make(map[*Role]struct{})
	stage.Roles_mapString = make(map[string]*Role)
	stage.RoleMap_Staged_Order = make(map[*Role]uint)
	stage.RoleOrder = 0

	stage.States = make(map[*State]struct{})
	stage.States_mapString = make(map[string]*State)
	stage.StateMap_Staged_Order = make(map[*State]uint)
	stage.StateOrder = 0

	stage.StateMachines = make(map[*StateMachine]struct{})
	stage.StateMachines_mapString = make(map[string]*StateMachine)
	stage.StateMachineMap_Staged_Order = make(map[*StateMachine]uint)
	stage.StateMachineOrder = 0

	stage.StateShapes = make(map[*StateShape]struct{})
	stage.StateShapes_mapString = make(map[string]*StateShape)
	stage.StateShapeMap_Staged_Order = make(map[*StateShape]uint)
	stage.StateShapeOrder = 0

	stage.Transitions = make(map[*Transition]struct{})
	stage.Transitions_mapString = make(map[string]*Transition)
	stage.TransitionMap_Staged_Order = make(map[*Transition]uint)
	stage.TransitionOrder = 0

	stage.Transition_Shapes = make(map[*Transition_Shape]struct{})
	stage.Transition_Shapes_mapString = make(map[string]*Transition_Shape)
	stage.Transition_ShapeMap_Staged_Order = make(map[*Transition_Shape]uint)
	stage.Transition_ShapeOrder = 0

	stage.ComputeReference()
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

	stage.Messages = nil
	stage.Messages_mapString = nil

	stage.MessageTypes = nil
	stage.MessageTypes_mapString = nil

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

	for message := range stage.Messages {
		message.Unstage(stage)
	}

	for messagetype := range stage.MessageTypes {
		messagetype.Unstage(stage)
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
	GongClean(stage *Stage)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
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
	case map[*Message]any:
		return any(&stage.Messages).(*Type)
	case map[*MessageType]any:
		return any(&stage.MessageTypes).(*Type)
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
	case *Message:
		return any(stage.Messages_mapString).(map[string]Type)
	case *MessageType:
		return any(stage.MessageTypes_mapString).(map[string]Type)
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
	case Message:
		return any(&stage.Messages).(*map[*Type]struct{})
	case MessageType:
		return any(&stage.MessageTypes).(*map[*Type]struct{})
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
	case *Message:
		return any(&stage.Messages).(*map[Type]struct{})
	case *MessageType:
		return any(&stage.MessageTypes).(*map[Type]struct{})
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
	case Message:
		return any(&stage.Messages_mapString).(*map[string]*Type)
	case MessageType:
		return any(&stage.MessageTypes_mapString).(*map[string]*Type)
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
			// field is initialized with an instance of StateShape with the name of the field
			State_Shapes: []*StateShape{{Name: "State_Shapes"}},
			// field is initialized with an instance of Transition_Shape with the name of the field
			Transition_Shapes: []*Transition_Shape{{Name: "Transition_Shapes"}},
		}).(*Type)
	case Guard:
		return any(&Guard{
			// Initialisation of associations
		}).(*Type)
	case Kill:
		return any(&Kill{
			// Initialisation of associations
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
		case "State_Shapes":
			res := make(map[*StateShape][]*Diagram)
			for diagram := range stage.Diagrams {
				for _, stateshape_ := range diagram.State_Shapes {
					res[stateshape_] = append(res[stateshape_], diagram)
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
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

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
	case *Message:
		res = "Message"
	case *MessageType:
		res = "MessageType"
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

func GetReverseFields[Type PointerToGongstruct]() (res []ReverseField) {

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

// insertion point for get fields header method
func (action *Action) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Criticality",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (activities *Activities) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Criticality",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (architecture *Architecture) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (diagram *Diagram) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsEditable_",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsInRenameMode",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "State_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "StateShape",
		},
		{
			Name:                 "Transition_Shapes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Transition_Shape",
		},
	}
	return
}

func (guard *Guard) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (kill *Kill) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (message *Message) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Description",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (object *Object) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "State",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
		{
			Name:               "IsSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Rank",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DOF",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Acronym",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Parent",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
		{
			Name:               "IsDecisionNode",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsFictif",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsEndState",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsNodeExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "State",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "State",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Width",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Height",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (transition *Transition) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Transition",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Transition",
		},
		{
			Name:               "StartRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndRatio",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndOrientation",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "CornerOffsetRatio",
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Roles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range architecture.Roles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
	case "IsInRenameMode":
		res.valueString = fmt.Sprintf("%t", diagram.IsInRenameMode)
		res.valueBool = diagram.IsInRenameMode
		res.GongFieldValueType = GongFieldValueTypeBool
	case "State_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.State_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Transition_Shapes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range diagram.Transition_Shapes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, message.MessageType))
		}
	case "OriginTransition":
		res.GongFieldValueType = GongFieldValueTypePointer
		if message.OriginTransition != nil {
			res.valueString = message.OriginTransition.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, message.OriginTransition))
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
func (object *Object) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = object.Name
	case "State":
		res.GongFieldValueType = GongFieldValueTypePointer
		if object.State != nil {
			res.valueString = object.State.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, object.State))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, state.Parent))
		}
	case "IsDecisionNode":
		res.valueString = fmt.Sprintf("%t", state.IsDecisionNode)
		res.valueBool = state.IsDecisionNode
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsFictif":
		res.valueString = fmt.Sprintf("%t", state.IsFictif)
		res.valueBool = state.IsFictif
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range state.Diagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Entry":
		res.GongFieldValueType = GongFieldValueTypePointer
		if state.Entry != nil {
			res.valueString = state.Entry.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, state.Entry))
		}
	case "Activities":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range state.Activities {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Exit":
		res.GongFieldValueType = GongFieldValueTypePointer
		if state.Exit != nil {
			res.valueString = state.Exit.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, state.Exit))
		}
	}
	return
}
func (statemachine *StateMachine) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = statemachine.Name
	case "IsNodeExpanded":
		res.valueString = fmt.Sprintf("%t", statemachine.IsNodeExpanded)
		res.valueBool = statemachine.IsNodeExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	case "States":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range statemachine.States {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range statemachine.Diagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "InitialState":
		res.GongFieldValueType = GongFieldValueTypePointer
		if statemachine.InitialState != nil {
			res.valueString = statemachine.InitialState.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, statemachine.InitialState))
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, stateshape.State))
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", stateshape.IsExpanded)
		res.valueBool = stateshape.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, transition.Start))
		}
	case "End":
		res.GongFieldValueType = GongFieldValueTypePointer
		if transition.End != nil {
			res.valueString = transition.End.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, transition.End))
		}
	case "RolesWithPermissions":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range transition.RolesWithPermissions {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "GeneratedMessages":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range transition.GeneratedMessages {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Guard":
		res.GongFieldValueType = GongFieldValueTypePointer
		if transition.Guard != nil {
			res.valueString = transition.Guard.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, transition.Guard))
		}
	case "Diagrams":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range transition.Diagrams {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, transition_shape.Transition))
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
					if stage.StateMachineMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.RoleMap_Staged_Order[__instance__] == uint(id) {
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
	case "IsInRenameMode":
		diagram.IsInRenameMode = value.GetValueBool()
	case "State_Shapes":
		diagram.State_Shapes = make([]*StateShape, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.StateShapes {
					if stage.StateShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.State_Shapes = append(diagram.State_Shapes, __instance__)
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
					if stage.Transition_ShapeMap_Staged_Order[__instance__] == uint(id) {
						diagram.Transition_Shapes = append(diagram.Transition_Shapes, __instance__)
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
				if stage.MessageTypeMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.TransitionMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.StateMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.MessageMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.RoleMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.StateMap_Staged_Order[__instance__] == uint(id) {
					state.Parent = __instance__
					break
				}
			}
		}
	case "IsDecisionNode":
		state.IsDecisionNode = value.GetValueBool()
	case "IsFictif":
		state.IsFictif = value.GetValueBool()
	case "IsEndState":
		state.IsEndState = value.GetValueBool()
	case "SubStates":
		state.SubStates = make([]*State, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.States {
					if stage.StateMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.DiagramMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.ActionMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.ActivitiesMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.ActionMap_Staged_Order[__instance__] == uint(id) {
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
	case "IsNodeExpanded":
		statemachine.IsNodeExpanded = value.GetValueBool()
	case "States":
		statemachine.States = make([]*State, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.States {
					if stage.StateMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.DiagramMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.StateMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.StateMap_Staged_Order[__instance__] == uint(id) {
					stateshape.State = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		stateshape.IsExpanded = value.GetValueBool()
	case "X":
		stateshape.X = value.GetValueFloat()
	case "Y":
		stateshape.Y = value.GetValueFloat()
	case "Width":
		stateshape.Width = value.GetValueFloat()
	case "Height":
		stateshape.Height = value.GetValueFloat()
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
				if stage.StateMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.StateMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.RoleMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.MessageTypeMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.GuardMap_Staged_Order[__instance__] == uint(id) {
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
					if stage.DiagramMap_Staged_Order[__instance__] == uint(id) {
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
				if stage.TransitionMap_Staged_Order[__instance__] == uint(id) {
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

func (message *Message) GongGetGongstructName() string {
	return "Message"
}

func (messagetype *MessageType) GongGetGongstructName() string {
	return "MessageType"
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

	stage.Messages_mapString = make(map[string]*Message)
	for message := range stage.Messages {
		stage.Messages_mapString[message.Name] = message
	}

	stage.MessageTypes_mapString = make(map[string]*MessageType)
	for messagetype := range stage.MessageTypes {
		stage.MessageTypes_mapString[messagetype.Name] = messagetype
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

}
// Last line of the template
