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
	// Compute reverse map for named struct Action
	// insertion point per field

	// Compute reverse map for named struct Activities
	// insertion point per field

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

	// Compute reverse map for named struct Guard
	// insertion point per field

	// Compute reverse map for named struct Kill
	// insertion point per field

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

	// Compute reverse map for named struct Message
	// insertion point per field

	// Compute reverse map for named struct MessageType
	// insertion point per field

	// Compute reverse map for named struct Note
	// insertion point per field

	// Compute reverse map for named struct NoteShape
	// insertion point per field

	// Compute reverse map for named struct NoteStateShape
	// insertion point per field

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

	// Compute reverse map for named struct StateShape
	// insertion point per field

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

	// Compute reverse map for named struct Transition_Shape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Actions {
		res = append(res, instance)
	}

	for instance := range stage.Activitiess {
		res = append(res, instance)
	}

	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Guards {
		res = append(res, instance)
	}

	for instance := range stage.Kills {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Messages {
		res = append(res, instance)
	}

	for instance := range stage.MessageTypes {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.NoteShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteStateShapes {
		res = append(res, instance)
	}

	for instance := range stage.Objects {
		res = append(res, instance)
	}

	for instance := range stage.Roles {
		res = append(res, instance)
	}

	for instance := range stage.States {
		res = append(res, instance)
	}

	for instance := range stage.StateMachines {
		res = append(res, instance)
	}

	for instance := range stage.StateShapes {
		res = append(res, instance)
	}

	for instance := range stage.Transitions {
		res = append(res, instance)
	}

	for instance := range stage.Transition_Shapes {
		res = append(res, instance)
	}

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
func (action *Action) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(action).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(action), uint64(stage.GetOrder(action)))
	return
}

func (activities *Activities) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(activities).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(activities), uint64(stage.GetOrder(activities)))
	return
}

func (diagram *Diagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagram), uint64(stage.GetOrder(diagram)))
	return
}

func (guard *Guard) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(guard).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(guard), uint64(stage.GetOrder(guard)))
	return
}

func (kill *Kill) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(kill).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(kill), uint64(stage.GetOrder(kill)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(library), uint64(stage.GetOrder(library)))
	return
}

func (message *Message) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(message).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(message), uint64(stage.GetOrder(message)))
	return
}

func (messagetype *MessageType) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(messagetype).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(messagetype), uint64(stage.GetOrder(messagetype)))
	return
}

func (note *Note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(note), uint64(stage.GetOrder(note)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(noteshape), uint64(stage.GetOrder(noteshape)))
	return
}

func (notestateshape *NoteStateShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notestateshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notestateshape), uint64(stage.GetOrder(notestateshape)))
	return
}

func (object *Object) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(object).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(object), uint64(stage.GetOrder(object)))
	return
}

func (role *Role) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(role).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(role), uint64(stage.GetOrder(role)))
	return
}

func (state *State) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(state).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(state), uint64(stage.GetOrder(state)))
	return
}

func (statemachine *StateMachine) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(statemachine).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(statemachine), uint64(stage.GetOrder(statemachine)))
	return
}

func (stateshape *StateShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stateshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stateshape), uint64(stage.GetOrder(stateshape)))
	return
}

func (transition *Transition) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(transition).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(transition), uint64(stage.GetOrder(transition)))
	return
}

func (transition_shape *Transition_Shape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(transition_shape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(transition_shape), uint64(stage.GetOrder(transition_shape)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
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
	stage.Actions_reference = make(map[*Action]*Action)
	stage.Actions_referenceOrder = make(map[*Action]uint) // diff Unstage needs the reference order
	stage.Actions_instance = make(map[*Action]*Action)
	for instance := range stage.Actions {
		_copy := instance.GongCopy().(*Action)
		stage.Actions_reference[instance] = _copy
		stage.Actions_instance[_copy] = instance
		stage.Actions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Activitiess_reference = make(map[*Activities]*Activities)
	stage.Activitiess_referenceOrder = make(map[*Activities]uint) // diff Unstage needs the reference order
	stage.Activitiess_instance = make(map[*Activities]*Activities)
	for instance := range stage.Activitiess {
		_copy := instance.GongCopy().(*Activities)
		stage.Activitiess_reference[instance] = _copy
		stage.Activitiess_instance[_copy] = instance
		stage.Activitiess_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	stage.Diagrams_instance = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		_copy := instance.GongCopy().(*Diagram)
		stage.Diagrams_reference[instance] = _copy
		stage.Diagrams_instance[_copy] = instance
		stage.Diagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Guards_reference = make(map[*Guard]*Guard)
	stage.Guards_referenceOrder = make(map[*Guard]uint) // diff Unstage needs the reference order
	stage.Guards_instance = make(map[*Guard]*Guard)
	for instance := range stage.Guards {
		_copy := instance.GongCopy().(*Guard)
		stage.Guards_reference[instance] = _copy
		stage.Guards_instance[_copy] = instance
		stage.Guards_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Kills_reference = make(map[*Kill]*Kill)
	stage.Kills_referenceOrder = make(map[*Kill]uint) // diff Unstage needs the reference order
	stage.Kills_instance = make(map[*Kill]*Kill)
	for instance := range stage.Kills {
		_copy := instance.GongCopy().(*Kill)
		stage.Kills_reference[instance] = _copy
		stage.Kills_instance[_copy] = instance
		stage.Kills_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint) // diff Unstage needs the reference order
	stage.Librarys_instance = make(map[*Library]*Library)
	for instance := range stage.Librarys {
		_copy := instance.GongCopy().(*Library)
		stage.Librarys_reference[instance] = _copy
		stage.Librarys_instance[_copy] = instance
		stage.Librarys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Messages_reference = make(map[*Message]*Message)
	stage.Messages_referenceOrder = make(map[*Message]uint) // diff Unstage needs the reference order
	stage.Messages_instance = make(map[*Message]*Message)
	for instance := range stage.Messages {
		_copy := instance.GongCopy().(*Message)
		stage.Messages_reference[instance] = _copy
		stage.Messages_instance[_copy] = instance
		stage.Messages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MessageTypes_reference = make(map[*MessageType]*MessageType)
	stage.MessageTypes_referenceOrder = make(map[*MessageType]uint) // diff Unstage needs the reference order
	stage.MessageTypes_instance = make(map[*MessageType]*MessageType)
	for instance := range stage.MessageTypes {
		_copy := instance.GongCopy().(*MessageType)
		stage.MessageTypes_reference[instance] = _copy
		stage.MessageTypes_instance[_copy] = instance
		stage.MessageTypes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint) // diff Unstage needs the reference order
	stage.Notes_instance = make(map[*Note]*Note)
	for instance := range stage.Notes {
		_copy := instance.GongCopy().(*Note)
		stage.Notes_reference[instance] = _copy
		stage.Notes_instance[_copy] = instance
		stage.Notes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint) // diff Unstage needs the reference order
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	for instance := range stage.NoteShapes {
		_copy := instance.GongCopy().(*NoteShape)
		stage.NoteShapes_reference[instance] = _copy
		stage.NoteShapes_instance[_copy] = instance
		stage.NoteShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteStateShapes_reference = make(map[*NoteStateShape]*NoteStateShape)
	stage.NoteStateShapes_referenceOrder = make(map[*NoteStateShape]uint) // diff Unstage needs the reference order
	stage.NoteStateShapes_instance = make(map[*NoteStateShape]*NoteStateShape)
	for instance := range stage.NoteStateShapes {
		_copy := instance.GongCopy().(*NoteStateShape)
		stage.NoteStateShapes_reference[instance] = _copy
		stage.NoteStateShapes_instance[_copy] = instance
		stage.NoteStateShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Objects_reference = make(map[*Object]*Object)
	stage.Objects_referenceOrder = make(map[*Object]uint) // diff Unstage needs the reference order
	stage.Objects_instance = make(map[*Object]*Object)
	for instance := range stage.Objects {
		_copy := instance.GongCopy().(*Object)
		stage.Objects_reference[instance] = _copy
		stage.Objects_instance[_copy] = instance
		stage.Objects_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Roles_reference = make(map[*Role]*Role)
	stage.Roles_referenceOrder = make(map[*Role]uint) // diff Unstage needs the reference order
	stage.Roles_instance = make(map[*Role]*Role)
	for instance := range stage.Roles {
		_copy := instance.GongCopy().(*Role)
		stage.Roles_reference[instance] = _copy
		stage.Roles_instance[_copy] = instance
		stage.Roles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.States_reference = make(map[*State]*State)
	stage.States_referenceOrder = make(map[*State]uint) // diff Unstage needs the reference order
	stage.States_instance = make(map[*State]*State)
	for instance := range stage.States {
		_copy := instance.GongCopy().(*State)
		stage.States_reference[instance] = _copy
		stage.States_instance[_copy] = instance
		stage.States_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StateMachines_reference = make(map[*StateMachine]*StateMachine)
	stage.StateMachines_referenceOrder = make(map[*StateMachine]uint) // diff Unstage needs the reference order
	stage.StateMachines_instance = make(map[*StateMachine]*StateMachine)
	for instance := range stage.StateMachines {
		_copy := instance.GongCopy().(*StateMachine)
		stage.StateMachines_reference[instance] = _copy
		stage.StateMachines_instance[_copy] = instance
		stage.StateMachines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StateShapes_reference = make(map[*StateShape]*StateShape)
	stage.StateShapes_referenceOrder = make(map[*StateShape]uint) // diff Unstage needs the reference order
	stage.StateShapes_instance = make(map[*StateShape]*StateShape)
	for instance := range stage.StateShapes {
		_copy := instance.GongCopy().(*StateShape)
		stage.StateShapes_reference[instance] = _copy
		stage.StateShapes_instance[_copy] = instance
		stage.StateShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Transitions_reference = make(map[*Transition]*Transition)
	stage.Transitions_referenceOrder = make(map[*Transition]uint) // diff Unstage needs the reference order
	stage.Transitions_instance = make(map[*Transition]*Transition)
	for instance := range stage.Transitions {
		_copy := instance.GongCopy().(*Transition)
		stage.Transitions_reference[instance] = _copy
		stage.Transitions_instance[_copy] = instance
		stage.Transitions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Transition_Shapes_reference = make(map[*Transition_Shape]*Transition_Shape)
	stage.Transition_Shapes_referenceOrder = make(map[*Transition_Shape]uint) // diff Unstage needs the reference order
	stage.Transition_Shapes_instance = make(map[*Transition_Shape]*Transition_Shape)
	for instance := range stage.Transition_Shapes {
		_copy := instance.GongCopy().(*Transition_Shape)
		stage.Transition_Shapes_reference[instance] = _copy
		stage.Transition_Shapes_instance[_copy] = instance
		stage.Transition_Shapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Actions {
		reference := stage.Actions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Activitiess {
		reference := stage.Activitiess_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Diagrams {
		reference := stage.Diagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Guards {
		reference := stage.Guards_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Kills {
		reference := stage.Kills_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Messages {
		reference := stage.Messages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MessageTypes {
		reference := stage.MessageTypes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Notes {
		reference := stage.Notes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteShapes {
		reference := stage.NoteShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteStateShapes {
		reference := stage.NoteStateShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Objects {
		reference := stage.Objects_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Roles {
		reference := stage.Roles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.States {
		reference := stage.States_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StateMachines {
		reference := stage.StateMachines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StateShapes {
		reference := stage.StateShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Transitions {
		reference := stage.Transitions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Transition_Shapes {
		reference := stage.Transition_Shapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

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
	if order, ok := stage.Action_stagedOrder[action]; ok {
		return order
	}
	if order, ok := stage.Actions_referenceOrder[action]; ok {
		return order
	} else {
		log.Printf("instance %p of type Action was not staged and does not have a reference order", action)
		return 0
	}
}

func (activities *Activities) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Activities_stagedOrder[activities]; ok {
		return order
	}
	if order, ok := stage.Activitiess_referenceOrder[activities]; ok {
		return order
	} else {
		log.Printf("instance %p of type Activities was not staged and does not have a reference order", activities)
		return 0
	}
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Diagram_stagedOrder[diagram]; ok {
		return order
	}
	if order, ok := stage.Diagrams_referenceOrder[diagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Diagram was not staged and does not have a reference order", diagram)
		return 0
	}
}

func (guard *Guard) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Guard_stagedOrder[guard]; ok {
		return order
	}
	if order, ok := stage.Guards_referenceOrder[guard]; ok {
		return order
	} else {
		log.Printf("instance %p of type Guard was not staged and does not have a reference order", guard)
		return 0
	}
}

func (kill *Kill) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Kill_stagedOrder[kill]; ok {
		return order
	}
	if order, ok := stage.Kills_referenceOrder[kill]; ok {
		return order
	} else {
		log.Printf("instance %p of type Kill was not staged and does not have a reference order", kill)
		return 0
	}
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Library_stagedOrder[library]; ok {
		return order
	}
	if order, ok := stage.Librarys_referenceOrder[library]; ok {
		return order
	} else {
		log.Printf("instance %p of type Library was not staged and does not have a reference order", library)
		return 0
	}
}

func (message *Message) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Message_stagedOrder[message]; ok {
		return order
	}
	if order, ok := stage.Messages_referenceOrder[message]; ok {
		return order
	} else {
		log.Printf("instance %p of type Message was not staged and does not have a reference order", message)
		return 0
	}
}

func (messagetype *MessageType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MessageType_stagedOrder[messagetype]; ok {
		return order
	}
	if order, ok := stage.MessageTypes_referenceOrder[messagetype]; ok {
		return order
	} else {
		log.Printf("instance %p of type MessageType was not staged and does not have a reference order", messagetype)
		return 0
	}
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Note_stagedOrder[note]; ok {
		return order
	}
	if order, ok := stage.Notes_referenceOrder[note]; ok {
		return order
	} else {
		log.Printf("instance %p of type Note was not staged and does not have a reference order", note)
		return 0
	}
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteShape_stagedOrder[noteshape]; ok {
		return order
	}
	if order, ok := stage.NoteShapes_referenceOrder[noteshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteShape was not staged and does not have a reference order", noteshape)
		return 0
	}
}

func (notestateshape *NoteStateShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteStateShape_stagedOrder[notestateshape]; ok {
		return order
	}
	if order, ok := stage.NoteStateShapes_referenceOrder[notestateshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteStateShape was not staged and does not have a reference order", notestateshape)
		return 0
	}
}

func (object *Object) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Object_stagedOrder[object]; ok {
		return order
	}
	if order, ok := stage.Objects_referenceOrder[object]; ok {
		return order
	} else {
		log.Printf("instance %p of type Object was not staged and does not have a reference order", object)
		return 0
	}
}

func (role *Role) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Role_stagedOrder[role]; ok {
		return order
	}
	if order, ok := stage.Roles_referenceOrder[role]; ok {
		return order
	} else {
		log.Printf("instance %p of type Role was not staged and does not have a reference order", role)
		return 0
	}
}

func (state *State) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.State_stagedOrder[state]; ok {
		return order
	}
	if order, ok := stage.States_referenceOrder[state]; ok {
		return order
	} else {
		log.Printf("instance %p of type State was not staged and does not have a reference order", state)
		return 0
	}
}

func (statemachine *StateMachine) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StateMachine_stagedOrder[statemachine]; ok {
		return order
	}
	if order, ok := stage.StateMachines_referenceOrder[statemachine]; ok {
		return order
	} else {
		log.Printf("instance %p of type StateMachine was not staged and does not have a reference order", statemachine)
		return 0
	}
}

func (stateshape *StateShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StateShape_stagedOrder[stateshape]; ok {
		return order
	}
	if order, ok := stage.StateShapes_referenceOrder[stateshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type StateShape was not staged and does not have a reference order", stateshape)
		return 0
	}
}

func (transition *Transition) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Transition_stagedOrder[transition]; ok {
		return order
	}
	if order, ok := stage.Transitions_referenceOrder[transition]; ok {
		return order
	} else {
		log.Printf("instance %p of type Transition was not staged and does not have a reference order", transition)
		return 0
	}
}

func (transition_shape *Transition_Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Transition_Shape_stagedOrder[transition_shape]; ok {
		return order
	}
	if order, ok := stage.Transition_Shapes_referenceOrder[transition_shape]; ok {
		return order
	} else {
		log.Printf("instance %p of type Transition_Shape was not staged and does not have a reference order", transition_shape)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (action *Action) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", action.GongGetGongstructName(), action.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (action *Action) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", action.GongGetGongstructName(), action.GongGetOrder(stage))
}

func (activities *Activities) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", activities.GongGetGongstructName(), activities.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (activities *Activities) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", activities.GongGetGongstructName(), activities.GongGetOrder(stage))
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

func (guard *Guard) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", guard.GongGetGongstructName(), guard.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (guard *Guard) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", guard.GongGetGongstructName(), guard.GongGetOrder(stage))
}

func (kill *Kill) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", kill.GongGetGongstructName(), kill.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (kill *Kill) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", kill.GongGetGongstructName(), kill.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (message *Message) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", message.GongGetGongstructName(), message.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (message *Message) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", message.GongGetGongstructName(), message.GongGetOrder(stage))
}

func (messagetype *MessageType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", messagetype.GongGetGongstructName(), messagetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (messagetype *MessageType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", messagetype.GongGetGongstructName(), messagetype.GongGetOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

func (notestateshape *NoteStateShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notestateshape.GongGetGongstructName(), notestateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notestateshape *NoteStateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notestateshape.GongGetGongstructName(), notestateshape.GongGetOrder(stage))
}

func (object *Object) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", object.GongGetGongstructName(), object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (object *Object) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", object.GongGetGongstructName(), object.GongGetOrder(stage))
}

func (role *Role) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", role.GongGetGongstructName(), role.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (role *Role) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", role.GongGetGongstructName(), role.GongGetOrder(stage))
}

func (state *State) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", state.GongGetGongstructName(), state.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (state *State) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", state.GongGetGongstructName(), state.GongGetOrder(stage))
}

func (statemachine *StateMachine) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", statemachine.GongGetGongstructName(), statemachine.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (statemachine *StateMachine) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", statemachine.GongGetGongstructName(), statemachine.GongGetOrder(stage))
}

func (stateshape *StateShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stateshape.GongGetGongstructName(), stateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stateshape *StateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stateshape.GongGetGongstructName(), stateshape.GongGetOrder(stage))
}

func (transition *Transition) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transition.GongGetGongstructName(), transition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transition *Transition) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transition.GongGetGongstructName(), transition.GongGetOrder(stage))
}

func (transition_shape *Transition_Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transition_shape.GongGetGongstructName(), transition_shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transition_shape *Transition_Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transition_shape.GongGetGongstructName(), transition_shape.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (action *Action) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", action.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Action")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(action.Name))
	return
}

func (activities *Activities) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", activities.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Activities")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(activities.Name))
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagram.Name))
	return
}

func (guard *Guard) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", guard.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Guard")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(guard.Name))
	return
}

func (kill *Kill) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", kill.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Kill")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(kill.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(library.Name))
	return
}

func (message *Message) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", message.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Message")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(message.Name))
	return
}

func (messagetype *MessageType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MessageType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(messagetype.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(note.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(noteshape.Name))
	return
}

func (notestateshape *NoteStateShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notestateshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteStateShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notestateshape.Name))
	return
}

func (object *Object) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", object.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Object")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(object.Name))
	return
}

func (role *Role) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", role.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Role")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(role.Name))
	return
}

func (state *State) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", state.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "State")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(state.Name))
	return
}

func (statemachine *StateMachine) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StateMachine")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(statemachine.Name))
	return
}

func (stateshape *StateShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StateShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stateshape.Name))
	return
}

func (transition *Transition) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transition.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transition")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(transition.Name))
	return
}

func (transition_shape *Transition_Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transition_Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(transition_shape.Name))
	return
}

// insertion point for unstaging
func (action *Action) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", action.GongGetReferenceIdentifier(stage))
	return
}

func (activities *Activities) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", activities.GongGetReferenceIdentifier(stage))
	return
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetReferenceIdentifier(stage))
	return
}

func (guard *Guard) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", guard.GongGetReferenceIdentifier(stage))
	return
}

func (kill *Kill) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", kill.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (message *Message) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", message.GongGetReferenceIdentifier(stage))
	return
}

func (messagetype *MessageType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", messagetype.GongGetReferenceIdentifier(stage))
	return
}

func (note *Note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetReferenceIdentifier(stage))
	return
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetReferenceIdentifier(stage))
	return
}

func (notestateshape *NoteStateShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notestateshape.GongGetReferenceIdentifier(stage))
	return
}

func (object *Object) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", object.GongGetReferenceIdentifier(stage))
	return
}

func (role *Role) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", role.GongGetReferenceIdentifier(stage))
	return
}

func (state *State) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", state.GongGetReferenceIdentifier(stage))
	return
}

func (statemachine *StateMachine) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", statemachine.GongGetReferenceIdentifier(stage))
	return
}

func (stateshape *StateShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stateshape.GongGetReferenceIdentifier(stage))
	return
}

func (transition *Transition) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transition.GongGetReferenceIdentifier(stage))
	return
}

func (transition_shape *Transition_Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transition_shape.GongGetReferenceIdentifier(stage))
	return
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

// end of template
