// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_State_Shapes_reverseMap = make(map[*StateShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _stateshape := range diagram.State_Shapes {
			stage.Diagram_State_Shapes_reverseMap[_stateshape] = diagram
		}
	}
	stage.Diagram_StatesWhoseNodeIsExpanded_reverseMap = make(map[*State]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _state := range diagram.StatesWhoseNodeIsExpanded {
			stage.Diagram_StatesWhoseNodeIsExpanded_reverseMap[_state] = diagram
		}
	}
	stage.Diagram_Transition_Shapes_reverseMap = make(map[*Transition_Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _transition_shape := range diagram.Transition_Shapes {
			stage.Diagram_Transition_Shapes_reverseMap[_transition_shape] = diagram
		}
	}
	stage.Diagram_Note_Shapes_reverseMap = make(map[*NoteShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _noteshape := range diagram.Note_Shapes {
			stage.Diagram_Note_Shapes_reverseMap[_noteshape] = diagram
		}
	}
	stage.Diagram_NoteState_Shapes_reverseMap = make(map[*NoteStateShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _notestateshape := range diagram.NoteState_Shapes {
			stage.Diagram_NoteState_Shapes_reverseMap[_notestateshape] = diagram
		}
	}

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}
	stage.Library_Diagrams_reverseMap = make(map[*Diagram]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _diagram := range library.Diagrams {
			stage.Library_Diagrams_reverseMap[_diagram] = library
		}
	}
	stage.Library_RootStateMachines_reverseMap = make(map[*StateMachine]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _statemachine := range library.RootStateMachines {
			stage.Library_RootStateMachines_reverseMap[_statemachine] = library
		}
	}
	stage.Library_StateMachinesWhoseNodeIsExpanded_reverseMap = make(map[*StateMachine]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _statemachine := range library.StateMachinesWhoseNodeIsExpanded {
			stage.Library_StateMachinesWhoseNodeIsExpanded_reverseMap[_statemachine] = library
		}
	}
	stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[_library] = library
		}
	}
	stage.Library_Roles_reverseMap = make(map[*Role]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _role := range library.Roles {
			stage.Library_Roles_reverseMap[_role] = library
		}
	}
	stage.Library_MessageTypes_reverseMap = make(map[*MessageType]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _messagetype := range library.MessageTypes {
			stage.Library_MessageTypes_reverseMap[_messagetype] = library
		}
	}

	// Compute reverse map for named struct Object
	// insertion point per field
	stage.Object_Messages_reverseMap = make(map[*Message]*Object)
	for object := range stage.Objects {
		_ = object
		for _, _message := range object.Messages {
			stage.Object_Messages_reverseMap[_message] = object
		}
	}

	// Compute reverse map for named struct Role
	// insertion point per field
	stage.Role_RolesWithSamePermissions_reverseMap = make(map[*Role]*Role)
	for role := range stage.Roles {
		_ = role
		for _, _role := range role.RolesWithSamePermissions {
			stage.Role_RolesWithSamePermissions_reverseMap[_role] = role
		}
	}

	// Compute reverse map for named struct State
	// insertion point per field
	stage.State_SubStates_reverseMap = make(map[*State]*State)
	for state := range stage.States {
		_ = state
		for _, _state := range state.SubStates {
			stage.State_SubStates_reverseMap[_state] = state
		}
	}
	stage.State_Activities_reverseMap = make(map[*Activities]*State)
	for state := range stage.States {
		_ = state
		for _, _activities := range state.Activities {
			stage.State_Activities_reverseMap[_activities] = state
		}
	}
	stage.State_Diagrams_reverseMap = make(map[*Diagram]*State)
	for state := range stage.States {
		_ = state
		for _, _diagram := range state.Diagrams {
			stage.State_Diagrams_reverseMap[_diagram] = state
		}
	}
	stage.State_Notes_reverseMap = make(map[*Note]*State)
	for state := range stage.States {
		_ = state
		for _, _note := range state.Notes {
			stage.State_Notes_reverseMap[_note] = state
		}
	}

	// Compute reverse map for named struct StateMachine
	// insertion point per field
	stage.StateMachine_States_reverseMap = make(map[*State]*StateMachine)
	for statemachine := range stage.StateMachines {
		_ = statemachine
		for _, _state := range statemachine.States {
			stage.StateMachine_States_reverseMap[_state] = statemachine
		}
	}
	stage.StateMachine_Diagrams_reverseMap = make(map[*Diagram]*StateMachine)
	for statemachine := range stage.StateMachines {
		_ = statemachine
		for _, _diagram := range statemachine.Diagrams {
			stage.StateMachine_Diagrams_reverseMap[_diagram] = statemachine
		}
	}

	// Compute reverse map for named struct Transition
	// insertion point per field
	stage.Transition_RolesWithPermissions_reverseMap = make(map[*Role]*Transition)
	for transition := range stage.Transitions {
		_ = transition
		for _, _role := range transition.RolesWithPermissions {
			stage.Transition_RolesWithPermissions_reverseMap[_role] = transition
		}
	}
	stage.Transition_GeneratedMessages_reverseMap = make(map[*MessageType]*Transition)
	for transition := range stage.Transitions {
		_ = transition
		for _, _messagetype := range transition.GeneratedMessages {
			stage.Transition_GeneratedMessages_reverseMap[_messagetype] = transition
		}
	}
	stage.Transition_Diagrams_reverseMap = make(map[*Diagram]*Transition)
	for transition := range stage.Transitions {
		_ = transition
		for _, _diagram := range transition.Diagrams {
			stage.Transition_Diagrams_reverseMap[_diagram] = transition
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Actions)

	res = __gong__appendInstances(res, stage.Activitiess)

	res = __gong__appendInstances(res, stage.Diagrams)

	res = __gong__appendInstances(res, stage.Guards)

	res = __gong__appendInstances(res, stage.Kills)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.Messages)

	res = __gong__appendInstances(res, stage.MessageTypes)

	res = __gong__appendInstances(res, stage.Notes)

	res = __gong__appendInstances(res, stage.NoteShapes)

	res = __gong__appendInstances(res, stage.NoteStateShapes)

	res = __gong__appendInstances(res, stage.Objects)

	res = __gong__appendInstances(res, stage.Roles)

	res = __gong__appendInstances(res, stage.States)

	res = __gong__appendInstances(res, stage.StateMachines)

	res = __gong__appendInstances(res, stage.StateShapes)

	res = __gong__appendInstances(res, stage.Transitions)

	res = __gong__appendInstances(res, stage.Transition_Shapes)

	return
}

// insertion point per named struct
func (action *Action) GongCopy() GongstructIF {
	newInstance := new(Action)
	action.GongCopyBasicFields(newInstance)
	return newInstance
}

func (activities *Activities) GongCopy() GongstructIF {
	newInstance := new(Activities)
	activities.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (guard *Guard) GongCopy() GongstructIF {
	newInstance := new(Guard)
	guard.GongCopyBasicFields(newInstance)
	return newInstance
}

func (kill *Kill) GongCopy() GongstructIF {
	newInstance := new(Kill)
	kill.GongCopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.GongCopyBasicFields(newInstance)
	return newInstance
}

func (message *Message) GongCopy() GongstructIF {
	newInstance := new(Message)
	message.GongCopyBasicFields(newInstance)
	return newInstance
}

func (messagetype *MessageType) GongCopy() GongstructIF {
	newInstance := new(MessageType)
	messagetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notestateshape *NoteStateShape) GongCopy() GongstructIF {
	newInstance := new(NoteStateShape)
	notestateshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (object *Object) GongCopy() GongstructIF {
	newInstance := new(Object)
	object.GongCopyBasicFields(newInstance)
	return newInstance
}

func (role *Role) GongCopy() GongstructIF {
	newInstance := new(Role)
	role.GongCopyBasicFields(newInstance)
	return newInstance
}

func (state *State) GongCopy() GongstructIF {
	newInstance := new(State)
	state.GongCopyBasicFields(newInstance)
	return newInstance
}

func (statemachine *StateMachine) GongCopy() GongstructIF {
	newInstance := new(StateMachine)
	statemachine.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stateshape *StateShape) GongCopy() GongstructIF {
	newInstance := new(StateShape)
	stateshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (transition *Transition) GongCopy() GongstructIF {
	newInstance := new(Transition)
	transition.GongCopyBasicFields(newInstance)
	return newInstance
}

func (transition_shape *Transition_Shape) GongCopy() GongstructIF {
	newInstance := new(Transition_Shape)
	transition_shape.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (action *Action) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, action)
}

func (activities *Activities) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, activities)
}

func (diagram *Diagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagram)
}

func (guard *Guard) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, guard)
}

func (kill *Kill) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, kill)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (message *Message) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, message)
}

func (messagetype *MessageType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, messagetype)
}

func (note *Note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note)
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteshape)
}

func (notestateshape *NoteStateShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notestateshape)
}

func (object *Object) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, object)
}

func (role *Role) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, role)
}

func (state *State) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, state)
}

func (statemachine *StateMachine) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, statemachine)
}

func (stateshape *StateShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stateshape)
}

func (transition *Transition) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, transition)
}

func (transition_shape *Transition_Shape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, transition_shape)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.Actions,
		stage.Action_stagedOrder,
		stage.Actions_reference,
		&stage.Actions_referenceOrder,
		stage.Actions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Activitiess,
		stage.Activities_stagedOrder,
		stage.Activitiess_reference,
		&stage.Activitiess_referenceOrder,
		stage.Activitiess_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Diagrams,
		stage.Diagram_stagedOrder,
		stage.Diagrams_reference,
		&stage.Diagrams_referenceOrder,
		stage.Diagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Guards,
		stage.Guard_stagedOrder,
		stage.Guards_reference,
		&stage.Guards_referenceOrder,
		stage.Guards_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Kills,
		stage.Kill_stagedOrder,
		stage.Kills_reference,
		&stage.Kills_referenceOrder,
		stage.Kills_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Librarys,
		stage.Library_stagedOrder,
		stage.Librarys_reference,
		&stage.Librarys_referenceOrder,
		stage.Librarys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Messages,
		stage.Message_stagedOrder,
		stage.Messages_reference,
		&stage.Messages_referenceOrder,
		stage.Messages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.MessageTypes,
		stage.MessageType_stagedOrder,
		stage.MessageTypes_reference,
		&stage.MessageTypes_referenceOrder,
		stage.MessageTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Notes,
		stage.Note_stagedOrder,
		stage.Notes_reference,
		&stage.Notes_referenceOrder,
		stage.Notes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteShapes,
		stage.NoteShape_stagedOrder,
		stage.NoteShapes_reference,
		&stage.NoteShapes_referenceOrder,
		stage.NoteShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteStateShapes,
		stage.NoteStateShape_stagedOrder,
		stage.NoteStateShapes_reference,
		&stage.NoteStateShapes_referenceOrder,
		stage.NoteStateShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Objects,
		stage.Object_stagedOrder,
		stage.Objects_reference,
		&stage.Objects_referenceOrder,
		stage.Objects_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Roles,
		stage.Role_stagedOrder,
		stage.Roles_reference,
		&stage.Roles_referenceOrder,
		stage.Roles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.States,
		stage.State_stagedOrder,
		stage.States_reference,
		&stage.States_referenceOrder,
		stage.States_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StateMachines,
		stage.StateMachine_stagedOrder,
		stage.StateMachines_reference,
		&stage.StateMachines_referenceOrder,
		stage.StateMachines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.StateShapes,
		stage.StateShape_stagedOrder,
		stage.StateShapes_reference,
		&stage.StateShapes_referenceOrder,
		stage.StateShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Transitions,
		stage.Transition_stagedOrder,
		stage.Transitions_reference,
		&stage.Transitions_referenceOrder,
		stage.Transitions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Transition_Shapes,
		stage.Transition_Shape_stagedOrder,
		stage.Transition_Shapes_reference,
		&stage.Transition_Shapes_referenceOrder,
		stage.Transition_Shapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.Actions, &stage.Actions_reference, &stage.Actions_referenceOrder, &stage.Actions_instance)

	__gong__computeReferencePass1(stage, stage.Activitiess, &stage.Activitiess_reference, &stage.Activitiess_referenceOrder, &stage.Activitiess_instance)

	__gong__computeReferencePass1(stage, stage.Diagrams, &stage.Diagrams_reference, &stage.Diagrams_referenceOrder, &stage.Diagrams_instance)

	__gong__computeReferencePass1(stage, stage.Guards, &stage.Guards_reference, &stage.Guards_referenceOrder, &stage.Guards_instance)

	__gong__computeReferencePass1(stage, stage.Kills, &stage.Kills_reference, &stage.Kills_referenceOrder, &stage.Kills_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.Messages, &stage.Messages_reference, &stage.Messages_referenceOrder, &stage.Messages_instance)

	__gong__computeReferencePass1(stage, stage.MessageTypes, &stage.MessageTypes_reference, &stage.MessageTypes_referenceOrder, &stage.MessageTypes_instance)

	__gong__computeReferencePass1(stage, stage.Notes, &stage.Notes_reference, &stage.Notes_referenceOrder, &stage.Notes_instance)

	__gong__computeReferencePass1(stage, stage.NoteShapes, &stage.NoteShapes_reference, &stage.NoteShapes_referenceOrder, &stage.NoteShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteStateShapes, &stage.NoteStateShapes_reference, &stage.NoteStateShapes_referenceOrder, &stage.NoteStateShapes_instance)

	__gong__computeReferencePass1(stage, stage.Objects, &stage.Objects_reference, &stage.Objects_referenceOrder, &stage.Objects_instance)

	__gong__computeReferencePass1(stage, stage.Roles, &stage.Roles_reference, &stage.Roles_referenceOrder, &stage.Roles_instance)

	__gong__computeReferencePass1(stage, stage.States, &stage.States_reference, &stage.States_referenceOrder, &stage.States_instance)

	__gong__computeReferencePass1(stage, stage.StateMachines, &stage.StateMachines_reference, &stage.StateMachines_referenceOrder, &stage.StateMachines_instance)

	__gong__computeReferencePass1(stage, stage.StateShapes, &stage.StateShapes_reference, &stage.StateShapes_referenceOrder, &stage.StateShapes_instance)

	__gong__computeReferencePass1(stage, stage.Transitions, &stage.Transitions_reference, &stage.Transitions_referenceOrder, &stage.Transitions_instance)

	__gong__computeReferencePass1(stage, stage.Transition_Shapes, &stage.Transition_Shapes_reference, &stage.Transition_Shapes_referenceOrder, &stage.Transition_Shapes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Actions, stage.Actions_reference, stage)

	__gong__computeReferencePass2(stage.Activitiess, stage.Activitiess_reference, stage)

	__gong__computeReferencePass2(stage.Diagrams, stage.Diagrams_reference, stage)

	__gong__computeReferencePass2(stage.Guards, stage.Guards_reference, stage)

	__gong__computeReferencePass2(stage.Kills, stage.Kills_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.Messages, stage.Messages_reference, stage)

	__gong__computeReferencePass2(stage.MessageTypes, stage.MessageTypes_reference, stage)

	__gong__computeReferencePass2(stage.Notes, stage.Notes_reference, stage)

	__gong__computeReferencePass2(stage.NoteShapes, stage.NoteShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteStateShapes, stage.NoteStateShapes_reference, stage)

	__gong__computeReferencePass2(stage.Objects, stage.Objects_reference, stage)

	__gong__computeReferencePass2(stage.Roles, stage.Roles_reference, stage)

	__gong__computeReferencePass2(stage.States, stage.States_reference, stage)

	__gong__computeReferencePass2(stage.StateMachines, stage.StateMachines_reference, stage)

	__gong__computeReferencePass2(stage.StateShapes, stage.StateShapes_reference, stage)

	__gong__computeReferencePass2(stage.Transitions, stage.Transitions_reference, stage)

	__gong__computeReferencePass2(stage.Transition_Shapes, stage.Transition_Shapes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (action *Action) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Action_stagedOrder, stage.Actions_referenceOrder, action, "Action")
}

func (activities *Activities) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Activities_stagedOrder, stage.Activitiess_referenceOrder, activities, "Activities")
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Diagram_stagedOrder, stage.Diagrams_referenceOrder, diagram, "Diagram")
}

func (guard *Guard) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Guard_stagedOrder, stage.Guards_referenceOrder, guard, "Guard")
}

func (kill *Kill) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Kill_stagedOrder, stage.Kills_referenceOrder, kill, "Kill")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (message *Message) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Message_stagedOrder, stage.Messages_referenceOrder, message, "Message")
}

func (messagetype *MessageType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MessageType_stagedOrder, stage.MessageTypes_referenceOrder, messagetype, "MessageType")
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_stagedOrder, stage.Notes_referenceOrder, note, "Note")
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteShape_stagedOrder, stage.NoteShapes_referenceOrder, noteshape, "NoteShape")
}

func (notestateshape *NoteStateShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteStateShape_stagedOrder, stage.NoteStateShapes_referenceOrder, notestateshape, "NoteStateShape")
}

func (object *Object) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Object_stagedOrder, stage.Objects_referenceOrder, object, "Object")
}

func (role *Role) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Role_stagedOrder, stage.Roles_referenceOrder, role, "Role")
}

func (state *State) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.State_stagedOrder, stage.States_referenceOrder, state, "State")
}

func (statemachine *StateMachine) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StateMachine_stagedOrder, stage.StateMachines_referenceOrder, statemachine, "StateMachine")
}

func (stateshape *StateShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StateShape_stagedOrder, stage.StateShapes_referenceOrder, stateshape, "StateShape")
}

func (transition *Transition) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Transition_stagedOrder, stage.Transitions_referenceOrder, transition, "Transition")
}

func (transition_shape *Transition_Shape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Transition_Shape_stagedOrder, stage.Transition_Shapes_referenceOrder, transition_shape, "Transition_Shape")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (action *Action) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(action, action.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (action *Action) GongGetReferenceIdentifier(stage *Stage) string {
	return action.GongGetIdentifier(stage)
}

func (activities *Activities) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(activities, activities.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (activities *Activities) GongGetReferenceIdentifier(stage *Stage) string {
	return activities.GongGetIdentifier(stage)
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagram, diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return diagram.GongGetIdentifier(stage)
}

func (guard *Guard) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(guard, guard.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (guard *Guard) GongGetReferenceIdentifier(stage *Stage) string {
	return guard.GongGetIdentifier(stage)
}

func (kill *Kill) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(kill, kill.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (kill *Kill) GongGetReferenceIdentifier(stage *Stage) string {
	return kill.GongGetIdentifier(stage)
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(library, library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return library.GongGetIdentifier(stage)
}

func (message *Message) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(message, message.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (message *Message) GongGetReferenceIdentifier(stage *Stage) string {
	return message.GongGetIdentifier(stage)
}

func (messagetype *MessageType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(messagetype, messagetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (messagetype *MessageType) GongGetReferenceIdentifier(stage *Stage) string {
	return messagetype.GongGetIdentifier(stage)
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(note, note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return note.GongGetIdentifier(stage)
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteshape, noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteshape.GongGetIdentifier(stage)
}

func (notestateshape *NoteStateShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notestateshape, notestateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notestateshape *NoteStateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return notestateshape.GongGetIdentifier(stage)
}

func (object *Object) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(object, object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (object *Object) GongGetReferenceIdentifier(stage *Stage) string {
	return object.GongGetIdentifier(stage)
}

func (role *Role) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(role, role.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (role *Role) GongGetReferenceIdentifier(stage *Stage) string {
	return role.GongGetIdentifier(stage)
}

func (state *State) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(state, state.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (state *State) GongGetReferenceIdentifier(stage *Stage) string {
	return state.GongGetIdentifier(stage)
}

func (statemachine *StateMachine) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(statemachine, statemachine.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (statemachine *StateMachine) GongGetReferenceIdentifier(stage *Stage) string {
	return statemachine.GongGetIdentifier(stage)
}

func (stateshape *StateShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stateshape, stateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stateshape *StateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return stateshape.GongGetIdentifier(stage)
}

func (transition *Transition) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(transition, transition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transition *Transition) GongGetReferenceIdentifier(stage *Stage) string {
	return transition.GongGetIdentifier(stage)
}

func (transition_shape *Transition_Shape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(transition_shape, transition_shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transition_shape *Transition_Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return transition_shape.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (action *Action) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(action.GongGetIdentifier(stage), "Action", action.Name)
}

func (activities *Activities) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(activities.GongGetIdentifier(stage), "Activities", activities.Name)
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagram.GongGetIdentifier(stage), "Diagram", diagram.Name)
}

func (guard *Guard) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(guard.GongGetIdentifier(stage), "Guard", guard.Name)
}

func (kill *Kill) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(kill.GongGetIdentifier(stage), "Kill", kill.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (message *Message) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(message.GongGetIdentifier(stage), "Message", message.Name)
}

func (messagetype *MessageType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(messagetype.GongGetIdentifier(stage), "MessageType", messagetype.Name)
}

func (note *Note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note.GongGetIdentifier(stage), "Note", note.Name)
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteshape.GongGetIdentifier(stage), "NoteShape", noteshape.Name)
}

func (notestateshape *NoteStateShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notestateshape.GongGetIdentifier(stage), "NoteStateShape", notestateshape.Name)
}

func (object *Object) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(object.GongGetIdentifier(stage), "Object", object.Name)
}

func (role *Role) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(role.GongGetIdentifier(stage), "Role", role.Name)
}

func (state *State) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(state.GongGetIdentifier(stage), "State", state.Name)
}

func (statemachine *StateMachine) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(statemachine.GongGetIdentifier(stage), "StateMachine", statemachine.Name)
}

func (stateshape *StateShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stateshape.GongGetIdentifier(stage), "StateShape", stateshape.Name)
}

func (transition *Transition) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(transition.GongGetIdentifier(stage), "Transition", transition.Name)
}

func (transition_shape *Transition_Shape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(transition_shape.GongGetIdentifier(stage), "Transition_Shape", transition_shape.Name)
}

// insertion point for unstaging
func (action *Action) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(action.GongGetReferenceIdentifier(stage))
}

func (activities *Activities) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(activities.GongGetReferenceIdentifier(stage))
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagram.GongGetReferenceIdentifier(stage))
}

func (guard *Guard) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(guard.GongGetReferenceIdentifier(stage))
}

func (kill *Kill) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(kill.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (message *Message) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(message.GongGetReferenceIdentifier(stage))
}

func (messagetype *MessageType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(messagetype.GongGetReferenceIdentifier(stage))
}

func (note *Note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note.GongGetReferenceIdentifier(stage))
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteshape.GongGetReferenceIdentifier(stage))
}

func (notestateshape *NoteStateShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notestateshape.GongGetReferenceIdentifier(stage))
}

func (object *Object) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(object.GongGetReferenceIdentifier(stage))
}

func (role *Role) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(role.GongGetReferenceIdentifier(stage))
}

func (state *State) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(state.GongGetReferenceIdentifier(stage))
}

func (statemachine *StateMachine) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(statemachine.GongGetReferenceIdentifier(stage))
}

func (stateshape *StateShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stateshape.GongGetReferenceIdentifier(stage))
}

func (transition *Transition) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(transition.GongGetReferenceIdentifier(stage))
}

func (transition_shape *Transition_Shape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(transition_shape.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
