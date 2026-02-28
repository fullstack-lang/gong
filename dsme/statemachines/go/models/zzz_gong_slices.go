// generated code - do not edit
package models

import (
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

	// Compute reverse map for named struct Architecture
	// insertion point per field
	stage.Architecture_StateMachines_reverseMap = make(map[*StateMachine]*Architecture)
	for architecture := range stage.Architectures {
		_ = architecture
		for _, _statemachine := range architecture.StateMachines {
			stage.Architecture_StateMachines_reverseMap[_statemachine] = architecture
		}
	}
	stage.Architecture_Roles_reverseMap = make(map[*Role]*Architecture)
	for architecture := range stage.Architectures {
		_ = architecture
		for _, _role := range architecture.Roles {
			stage.Architecture_Roles_reverseMap[_role] = architecture
		}
	}

	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_State_Shapes_reverseMap = make(map[*StateShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _stateshape := range diagram.State_Shapes {
			stage.Diagram_State_Shapes_reverseMap[_stateshape] = diagram
		}
	}
	stage.Diagram_Transition_Shapes_reverseMap = make(map[*Transition_Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _transition_shape := range diagram.Transition_Shapes {
			stage.Diagram_Transition_Shapes_reverseMap[_transition_shape] = diagram
		}
	}

	// Compute reverse map for named struct Guard
	// insertion point per field

	// Compute reverse map for named struct Kill
	// insertion point per field

	// Compute reverse map for named struct Message
	// insertion point per field

	// Compute reverse map for named struct MessageType
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
	stage.State_Diagrams_reverseMap = make(map[*Diagram]*State)
	for state := range stage.States {
		_ = state
		for _, _diagram := range state.Diagrams {
			stage.State_Diagrams_reverseMap[_diagram] = state
		}
	}
	stage.State_Activities_reverseMap = make(map[*Activities]*State)
	for state := range stage.States {
		_ = state
		for _, _activities := range state.Activities {
			stage.State_Activities_reverseMap[_activities] = state
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

	for instance := range stage.Architectures {
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

	for instance := range stage.Messages {
		res = append(res, instance)
	}

	for instance := range stage.MessageTypes {
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
	action.CopyBasicFields(newInstance)
	return newInstance
}

func (activities *Activities) GongCopy() GongstructIF {
	newInstance := new(Activities)
	activities.CopyBasicFields(newInstance)
	return newInstance
}

func (architecture *Architecture) GongCopy() GongstructIF {
	newInstance := new(Architecture)
	architecture.CopyBasicFields(newInstance)
	return newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := new(Diagram)
	diagram.CopyBasicFields(newInstance)
	return newInstance
}

func (guard *Guard) GongCopy() GongstructIF {
	newInstance := new(Guard)
	guard.CopyBasicFields(newInstance)
	return newInstance
}

func (kill *Kill) GongCopy() GongstructIF {
	newInstance := new(Kill)
	kill.CopyBasicFields(newInstance)
	return newInstance
}

func (message *Message) GongCopy() GongstructIF {
	newInstance := new(Message)
	message.CopyBasicFields(newInstance)
	return newInstance
}

func (messagetype *MessageType) GongCopy() GongstructIF {
	newInstance := new(MessageType)
	messagetype.CopyBasicFields(newInstance)
	return newInstance
}

func (object *Object) GongCopy() GongstructIF {
	newInstance := new(Object)
	object.CopyBasicFields(newInstance)
	return newInstance
}

func (role *Role) GongCopy() GongstructIF {
	newInstance := new(Role)
	role.CopyBasicFields(newInstance)
	return newInstance
}

func (state *State) GongCopy() GongstructIF {
	newInstance := new(State)
	state.CopyBasicFields(newInstance)
	return newInstance
}

func (statemachine *StateMachine) GongCopy() GongstructIF {
	newInstance := new(StateMachine)
	statemachine.CopyBasicFields(newInstance)
	return newInstance
}

func (stateshape *StateShape) GongCopy() GongstructIF {
	newInstance := new(StateShape)
	stateshape.CopyBasicFields(newInstance)
	return newInstance
}

func (transition *Transition) GongCopy() GongstructIF {
	newInstance := new(Transition)
	transition.CopyBasicFields(newInstance)
	return newInstance
}

func (transition_shape *Transition_Shape) GongCopy() GongstructIF {
	newInstance := new(Transition_Shape)
	transition_shape.CopyBasicFields(newInstance)
	return newInstance
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
	var actions_newInstances []*Action
	var actions_deletedInstances []*Action

	// parse all staged instances and check if they have a reference
	for action := range stage.Actions {
		if ref, ok := stage.Actions_reference[action]; !ok {
			actions_newInstances = append(actions_newInstances, action)
			newInstancesSlice = append(newInstancesSlice, action.GongMarshallIdentifier(stage))
			if stage.Actions_referenceOrder == nil {
				stage.Actions_referenceOrder = make(map[*Action]uint)
			}
			stage.Actions_referenceOrder[action] = stage.Action_stagedOrder[action]
			newInstancesReverseSlice = append(newInstancesReverseSlice, action.GongMarshallUnstaging(stage))
			// delete(stage.Actions_referenceOrder, action)
			fieldInitializers, pointersInitializations := action.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Action_stagedOrder[ref] = stage.Action_stagedOrder[action]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := action.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, action)
			// delete(stage.Action_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", action.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Actions_reference {
		instance := stage.Actions_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Actions[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			actions_deletedInstances = append(actions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(actions_newInstances)
	lenDeletedInstances += len(actions_deletedInstances)
	var activitiess_newInstances []*Activities
	var activitiess_deletedInstances []*Activities

	// parse all staged instances and check if they have a reference
	for activities := range stage.Activitiess {
		if ref, ok := stage.Activitiess_reference[activities]; !ok {
			activitiess_newInstances = append(activitiess_newInstances, activities)
			newInstancesSlice = append(newInstancesSlice, activities.GongMarshallIdentifier(stage))
			if stage.Activitiess_referenceOrder == nil {
				stage.Activitiess_referenceOrder = make(map[*Activities]uint)
			}
			stage.Activitiess_referenceOrder[activities] = stage.Activities_stagedOrder[activities]
			newInstancesReverseSlice = append(newInstancesReverseSlice, activities.GongMarshallUnstaging(stage))
			// delete(stage.Activitiess_referenceOrder, activities)
			fieldInitializers, pointersInitializations := activities.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Activities_stagedOrder[ref] = stage.Activities_stagedOrder[activities]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := activities.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, activities)
			// delete(stage.Activities_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", activities.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Activitiess_reference {
		instance := stage.Activitiess_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Activitiess[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			activitiess_deletedInstances = append(activitiess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(activitiess_newInstances)
	lenDeletedInstances += len(activitiess_deletedInstances)
	var architectures_newInstances []*Architecture
	var architectures_deletedInstances []*Architecture

	// parse all staged instances and check if they have a reference
	for architecture := range stage.Architectures {
		if ref, ok := stage.Architectures_reference[architecture]; !ok {
			architectures_newInstances = append(architectures_newInstances, architecture)
			newInstancesSlice = append(newInstancesSlice, architecture.GongMarshallIdentifier(stage))
			if stage.Architectures_referenceOrder == nil {
				stage.Architectures_referenceOrder = make(map[*Architecture]uint)
			}
			stage.Architectures_referenceOrder[architecture] = stage.Architecture_stagedOrder[architecture]
			newInstancesReverseSlice = append(newInstancesReverseSlice, architecture.GongMarshallUnstaging(stage))
			// delete(stage.Architectures_referenceOrder, architecture)
			fieldInitializers, pointersInitializations := architecture.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Architecture_stagedOrder[ref] = stage.Architecture_stagedOrder[architecture]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := architecture.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, architecture)
			// delete(stage.Architecture_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", architecture.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Architectures_reference {
		instance := stage.Architectures_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Architectures[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			architectures_deletedInstances = append(architectures_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(architectures_newInstances)
	lenDeletedInstances += len(architectures_deletedInstances)
	var diagrams_newInstances []*Diagram
	var diagrams_deletedInstances []*Diagram

	// parse all staged instances and check if they have a reference
	for diagram := range stage.Diagrams {
		if ref, ok := stage.Diagrams_reference[diagram]; !ok {
			diagrams_newInstances = append(diagrams_newInstances, diagram)
			newInstancesSlice = append(newInstancesSlice, diagram.GongMarshallIdentifier(stage))
			if stage.Diagrams_referenceOrder == nil {
				stage.Diagrams_referenceOrder = make(map[*Diagram]uint)
			}
			stage.Diagrams_referenceOrder[diagram] = stage.Diagram_stagedOrder[diagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagram.GongMarshallUnstaging(stage))
			// delete(stage.Diagrams_referenceOrder, diagram)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Diagram_stagedOrder[ref] = stage.Diagram_stagedOrder[diagram]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagram)
			// delete(stage.Diagram_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", diagram.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Diagrams_reference {
		instance := stage.Diagrams_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Diagrams[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagrams_deletedInstances = append(diagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagrams_newInstances)
	lenDeletedInstances += len(diagrams_deletedInstances)
	var guards_newInstances []*Guard
	var guards_deletedInstances []*Guard

	// parse all staged instances and check if they have a reference
	for guard := range stage.Guards {
		if ref, ok := stage.Guards_reference[guard]; !ok {
			guards_newInstances = append(guards_newInstances, guard)
			newInstancesSlice = append(newInstancesSlice, guard.GongMarshallIdentifier(stage))
			if stage.Guards_referenceOrder == nil {
				stage.Guards_referenceOrder = make(map[*Guard]uint)
			}
			stage.Guards_referenceOrder[guard] = stage.Guard_stagedOrder[guard]
			newInstancesReverseSlice = append(newInstancesReverseSlice, guard.GongMarshallUnstaging(stage))
			// delete(stage.Guards_referenceOrder, guard)
			fieldInitializers, pointersInitializations := guard.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Guard_stagedOrder[ref] = stage.Guard_stagedOrder[guard]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := guard.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, guard)
			// delete(stage.Guard_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", guard.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Guards_reference {
		instance := stage.Guards_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Guards[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			guards_deletedInstances = append(guards_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(guards_newInstances)
	lenDeletedInstances += len(guards_deletedInstances)
	var kills_newInstances []*Kill
	var kills_deletedInstances []*Kill

	// parse all staged instances and check if they have a reference
	for kill := range stage.Kills {
		if ref, ok := stage.Kills_reference[kill]; !ok {
			kills_newInstances = append(kills_newInstances, kill)
			newInstancesSlice = append(newInstancesSlice, kill.GongMarshallIdentifier(stage))
			if stage.Kills_referenceOrder == nil {
				stage.Kills_referenceOrder = make(map[*Kill]uint)
			}
			stage.Kills_referenceOrder[kill] = stage.Kill_stagedOrder[kill]
			newInstancesReverseSlice = append(newInstancesReverseSlice, kill.GongMarshallUnstaging(stage))
			// delete(stage.Kills_referenceOrder, kill)
			fieldInitializers, pointersInitializations := kill.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Kill_stagedOrder[ref] = stage.Kill_stagedOrder[kill]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := kill.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, kill)
			// delete(stage.Kill_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", kill.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Kills_reference {
		instance := stage.Kills_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Kills[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			kills_deletedInstances = append(kills_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(kills_newInstances)
	lenDeletedInstances += len(kills_deletedInstances)
	var messages_newInstances []*Message
	var messages_deletedInstances []*Message

	// parse all staged instances and check if they have a reference
	for message := range stage.Messages {
		if ref, ok := stage.Messages_reference[message]; !ok {
			messages_newInstances = append(messages_newInstances, message)
			newInstancesSlice = append(newInstancesSlice, message.GongMarshallIdentifier(stage))
			if stage.Messages_referenceOrder == nil {
				stage.Messages_referenceOrder = make(map[*Message]uint)
			}
			stage.Messages_referenceOrder[message] = stage.Message_stagedOrder[message]
			newInstancesReverseSlice = append(newInstancesReverseSlice, message.GongMarshallUnstaging(stage))
			// delete(stage.Messages_referenceOrder, message)
			fieldInitializers, pointersInitializations := message.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Message_stagedOrder[ref] = stage.Message_stagedOrder[message]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := message.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, message)
			// delete(stage.Message_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", message.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Messages_reference {
		instance := stage.Messages_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Messages[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			messages_deletedInstances = append(messages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(messages_newInstances)
	lenDeletedInstances += len(messages_deletedInstances)
	var messagetypes_newInstances []*MessageType
	var messagetypes_deletedInstances []*MessageType

	// parse all staged instances and check if they have a reference
	for messagetype := range stage.MessageTypes {
		if ref, ok := stage.MessageTypes_reference[messagetype]; !ok {
			messagetypes_newInstances = append(messagetypes_newInstances, messagetype)
			newInstancesSlice = append(newInstancesSlice, messagetype.GongMarshallIdentifier(stage))
			if stage.MessageTypes_referenceOrder == nil {
				stage.MessageTypes_referenceOrder = make(map[*MessageType]uint)
			}
			stage.MessageTypes_referenceOrder[messagetype] = stage.MessageType_stagedOrder[messagetype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, messagetype.GongMarshallUnstaging(stage))
			// delete(stage.MessageTypes_referenceOrder, messagetype)
			fieldInitializers, pointersInitializations := messagetype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MessageType_stagedOrder[ref] = stage.MessageType_stagedOrder[messagetype]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := messagetype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, messagetype)
			// delete(stage.MessageType_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", messagetype.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.MessageTypes_reference {
		instance := stage.MessageTypes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MessageTypes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			messagetypes_deletedInstances = append(messagetypes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(messagetypes_newInstances)
	lenDeletedInstances += len(messagetypes_deletedInstances)
	var objects_newInstances []*Object
	var objects_deletedInstances []*Object

	// parse all staged instances and check if they have a reference
	for object := range stage.Objects {
		if ref, ok := stage.Objects_reference[object]; !ok {
			objects_newInstances = append(objects_newInstances, object)
			newInstancesSlice = append(newInstancesSlice, object.GongMarshallIdentifier(stage))
			if stage.Objects_referenceOrder == nil {
				stage.Objects_referenceOrder = make(map[*Object]uint)
			}
			stage.Objects_referenceOrder[object] = stage.Object_stagedOrder[object]
			newInstancesReverseSlice = append(newInstancesReverseSlice, object.GongMarshallUnstaging(stage))
			// delete(stage.Objects_referenceOrder, object)
			fieldInitializers, pointersInitializations := object.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Object_stagedOrder[ref] = stage.Object_stagedOrder[object]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := object.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, object)
			// delete(stage.Object_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", object.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Objects_reference {
		instance := stage.Objects_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Objects[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			objects_deletedInstances = append(objects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(objects_newInstances)
	lenDeletedInstances += len(objects_deletedInstances)
	var roles_newInstances []*Role
	var roles_deletedInstances []*Role

	// parse all staged instances and check if they have a reference
	for role := range stage.Roles {
		if ref, ok := stage.Roles_reference[role]; !ok {
			roles_newInstances = append(roles_newInstances, role)
			newInstancesSlice = append(newInstancesSlice, role.GongMarshallIdentifier(stage))
			if stage.Roles_referenceOrder == nil {
				stage.Roles_referenceOrder = make(map[*Role]uint)
			}
			stage.Roles_referenceOrder[role] = stage.Role_stagedOrder[role]
			newInstancesReverseSlice = append(newInstancesReverseSlice, role.GongMarshallUnstaging(stage))
			// delete(stage.Roles_referenceOrder, role)
			fieldInitializers, pointersInitializations := role.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Role_stagedOrder[ref] = stage.Role_stagedOrder[role]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := role.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, role)
			// delete(stage.Role_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", role.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Roles_reference {
		instance := stage.Roles_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Roles[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			roles_deletedInstances = append(roles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(roles_newInstances)
	lenDeletedInstances += len(roles_deletedInstances)
	var states_newInstances []*State
	var states_deletedInstances []*State

	// parse all staged instances and check if they have a reference
	for state := range stage.States {
		if ref, ok := stage.States_reference[state]; !ok {
			states_newInstances = append(states_newInstances, state)
			newInstancesSlice = append(newInstancesSlice, state.GongMarshallIdentifier(stage))
			if stage.States_referenceOrder == nil {
				stage.States_referenceOrder = make(map[*State]uint)
			}
			stage.States_referenceOrder[state] = stage.State_stagedOrder[state]
			newInstancesReverseSlice = append(newInstancesReverseSlice, state.GongMarshallUnstaging(stage))
			// delete(stage.States_referenceOrder, state)
			fieldInitializers, pointersInitializations := state.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.State_stagedOrder[ref] = stage.State_stagedOrder[state]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := state.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, state)
			// delete(stage.State_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", state.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.States_reference {
		instance := stage.States_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.States[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			states_deletedInstances = append(states_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(states_newInstances)
	lenDeletedInstances += len(states_deletedInstances)
	var statemachines_newInstances []*StateMachine
	var statemachines_deletedInstances []*StateMachine

	// parse all staged instances and check if they have a reference
	for statemachine := range stage.StateMachines {
		if ref, ok := stage.StateMachines_reference[statemachine]; !ok {
			statemachines_newInstances = append(statemachines_newInstances, statemachine)
			newInstancesSlice = append(newInstancesSlice, statemachine.GongMarshallIdentifier(stage))
			if stage.StateMachines_referenceOrder == nil {
				stage.StateMachines_referenceOrder = make(map[*StateMachine]uint)
			}
			stage.StateMachines_referenceOrder[statemachine] = stage.StateMachine_stagedOrder[statemachine]
			newInstancesReverseSlice = append(newInstancesReverseSlice, statemachine.GongMarshallUnstaging(stage))
			// delete(stage.StateMachines_referenceOrder, statemachine)
			fieldInitializers, pointersInitializations := statemachine.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StateMachine_stagedOrder[ref] = stage.StateMachine_stagedOrder[statemachine]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := statemachine.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, statemachine)
			// delete(stage.StateMachine_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", statemachine.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StateMachines_reference {
		instance := stage.StateMachines_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StateMachines[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			statemachines_deletedInstances = append(statemachines_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(statemachines_newInstances)
	lenDeletedInstances += len(statemachines_deletedInstances)
	var stateshapes_newInstances []*StateShape
	var stateshapes_deletedInstances []*StateShape

	// parse all staged instances and check if they have a reference
	for stateshape := range stage.StateShapes {
		if ref, ok := stage.StateShapes_reference[stateshape]; !ok {
			stateshapes_newInstances = append(stateshapes_newInstances, stateshape)
			newInstancesSlice = append(newInstancesSlice, stateshape.GongMarshallIdentifier(stage))
			if stage.StateShapes_referenceOrder == nil {
				stage.StateShapes_referenceOrder = make(map[*StateShape]uint)
			}
			stage.StateShapes_referenceOrder[stateshape] = stage.StateShape_stagedOrder[stateshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stateshape.GongMarshallUnstaging(stage))
			// delete(stage.StateShapes_referenceOrder, stateshape)
			fieldInitializers, pointersInitializations := stateshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StateShape_stagedOrder[ref] = stage.StateShape_stagedOrder[stateshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := stateshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stateshape)
			// delete(stage.StateShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", stateshape.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.StateShapes_reference {
		instance := stage.StateShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.StateShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			stateshapes_deletedInstances = append(stateshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(stateshapes_newInstances)
	lenDeletedInstances += len(stateshapes_deletedInstances)
	var transitions_newInstances []*Transition
	var transitions_deletedInstances []*Transition

	// parse all staged instances and check if they have a reference
	for transition := range stage.Transitions {
		if ref, ok := stage.Transitions_reference[transition]; !ok {
			transitions_newInstances = append(transitions_newInstances, transition)
			newInstancesSlice = append(newInstancesSlice, transition.GongMarshallIdentifier(stage))
			if stage.Transitions_referenceOrder == nil {
				stage.Transitions_referenceOrder = make(map[*Transition]uint)
			}
			stage.Transitions_referenceOrder[transition] = stage.Transition_stagedOrder[transition]
			newInstancesReverseSlice = append(newInstancesReverseSlice, transition.GongMarshallUnstaging(stage))
			// delete(stage.Transitions_referenceOrder, transition)
			fieldInitializers, pointersInitializations := transition.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Transition_stagedOrder[ref] = stage.Transition_stagedOrder[transition]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := transition.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, transition)
			// delete(stage.Transition_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", transition.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Transitions_reference {
		instance := stage.Transitions_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Transitions[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			transitions_deletedInstances = append(transitions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(transitions_newInstances)
	lenDeletedInstances += len(transitions_deletedInstances)
	var transition_shapes_newInstances []*Transition_Shape
	var transition_shapes_deletedInstances []*Transition_Shape

	// parse all staged instances and check if they have a reference
	for transition_shape := range stage.Transition_Shapes {
		if ref, ok := stage.Transition_Shapes_reference[transition_shape]; !ok {
			transition_shapes_newInstances = append(transition_shapes_newInstances, transition_shape)
			newInstancesSlice = append(newInstancesSlice, transition_shape.GongMarshallIdentifier(stage))
			if stage.Transition_Shapes_referenceOrder == nil {
				stage.Transition_Shapes_referenceOrder = make(map[*Transition_Shape]uint)
			}
			stage.Transition_Shapes_referenceOrder[transition_shape] = stage.Transition_Shape_stagedOrder[transition_shape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, transition_shape.GongMarshallUnstaging(stage))
			// delete(stage.Transition_Shapes_referenceOrder, transition_shape)
			fieldInitializers, pointersInitializations := transition_shape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Transition_Shape_stagedOrder[ref] = stage.Transition_Shape_stagedOrder[transition_shape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := transition_shape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, transition_shape)
			// delete(stage.Transition_Shape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", transition_shape.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Transition_Shapes_reference {
		instance := stage.Transition_Shapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Transition_Shapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			transition_shapes_deletedInstances = append(transition_shapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(transition_shapes_newInstances)
	lenDeletedInstances += len(transition_shapes_deletedInstances)

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

	stage.Architectures_reference = make(map[*Architecture]*Architecture)
	stage.Architectures_referenceOrder = make(map[*Architecture]uint) // diff Unstage needs the reference order
	stage.Architectures_instance = make(map[*Architecture]*Architecture)
	for instance := range stage.Architectures {
		_copy := instance.GongCopy().(*Architecture)
		stage.Architectures_reference[instance] = _copy
		stage.Architectures_instance[_copy] = instance
		stage.Architectures_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	for instance := range stage.Architectures {
		reference := stage.Architectures_reference[instance]
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

	for instance := range stage.Messages {
		reference := stage.Messages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MessageTypes {
		reference := stage.MessageTypes_reference[instance]
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

func (architecture *Architecture) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Architecture_stagedOrder[architecture]; ok {
		return order
	}
	if order, ok := stage.Architectures_referenceOrder[architecture]; ok {
		return order
	} else {
		log.Printf("instance %p of type Architecture was not staged and does not have a reference order", architecture)
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

func (architecture *Architecture) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", architecture.GongGetGongstructName(), architecture.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (architecture *Architecture) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", architecture.GongGetGongstructName(), architecture.GongGetOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(action.Name))
	return
}

func (activities *Activities) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", activities.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Activities")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(activities.Name))
	return
}

func (architecture *Architecture) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", architecture.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Architecture")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(architecture.Name))
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagram.Name))
	return
}

func (guard *Guard) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", guard.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Guard")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(guard.Name))
	return
}

func (kill *Kill) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", kill.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Kill")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(kill.Name))
	return
}

func (message *Message) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", message.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Message")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(message.Name))
	return
}

func (messagetype *MessageType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MessageType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(messagetype.Name))
	return
}

func (object *Object) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", object.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Object")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(object.Name))
	return
}

func (role *Role) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", role.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Role")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(role.Name))
	return
}

func (state *State) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", state.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "State")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(state.Name))
	return
}

func (statemachine *StateMachine) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StateMachine")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(statemachine.Name))
	return
}

func (stateshape *StateShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StateShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(stateshape.Name))
	return
}

func (transition *Transition) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transition.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transition")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(transition.Name))
	return
}

func (transition_shape *Transition_Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transition_Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(transition_shape.Name))
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

func (architecture *Architecture) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", architecture.GongGetReferenceIdentifier(stage))
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

// end of template
