// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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
	newInstance := *action
	return &newInstance
}

func (activities *Activities) GongCopy() GongstructIF {
	newInstance := *activities
	return &newInstance
}

func (architecture *Architecture) GongCopy() GongstructIF {
	newInstance := *architecture
	return &newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := *diagram
	return &newInstance
}

func (guard *Guard) GongCopy() GongstructIF {
	newInstance := *guard
	return &newInstance
}

func (kill *Kill) GongCopy() GongstructIF {
	newInstance := *kill
	return &newInstance
}

func (message *Message) GongCopy() GongstructIF {
	newInstance := *message
	return &newInstance
}

func (messagetype *MessageType) GongCopy() GongstructIF {
	newInstance := *messagetype
	return &newInstance
}

func (object *Object) GongCopy() GongstructIF {
	newInstance := *object
	return &newInstance
}

func (role *Role) GongCopy() GongstructIF {
	newInstance := *role
	return &newInstance
}

func (state *State) GongCopy() GongstructIF {
	newInstance := *state
	return &newInstance
}

func (statemachine *StateMachine) GongCopy() GongstructIF {
	newInstance := *statemachine
	return &newInstance
}

func (stateshape *StateShape) GongCopy() GongstructIF {
	newInstance := *stateshape
	return &newInstance
}

func (transition *Transition) GongCopy() GongstructIF {
	newInstance := *transition
	return &newInstance
}

func (transition_shape *Transition_Shape) GongCopy() GongstructIF {
	newInstance := *transition_shape
	return &newInstance
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
			stage.Actions_referenceOrder[action] = stage.ActionMap_Staged_Order[action]
			newInstancesReverseSlice = append(newInstancesReverseSlice, action.GongMarshallUnstaging(stage))
			delete(stage.Actions_referenceOrder, action)
			fieldInitializers, pointersInitializations := action.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ActionMap_Staged_Order[ref] = stage.ActionMap_Staged_Order[action]
			diffs := action.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, action)
			delete(stage.ActionMap_Staged_Order, ref)
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
	for ref := range stage.Actions_reference {
		if _, ok := stage.Actions[ref]; !ok {
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
			stage.Activitiess_referenceOrder[activities] = stage.ActivitiesMap_Staged_Order[activities]
			newInstancesReverseSlice = append(newInstancesReverseSlice, activities.GongMarshallUnstaging(stage))
			delete(stage.Activitiess_referenceOrder, activities)
			fieldInitializers, pointersInitializations := activities.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ActivitiesMap_Staged_Order[ref] = stage.ActivitiesMap_Staged_Order[activities]
			diffs := activities.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, activities)
			delete(stage.ActivitiesMap_Staged_Order, ref)
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
	for ref := range stage.Activitiess_reference {
		if _, ok := stage.Activitiess[ref]; !ok {
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
			stage.Architectures_referenceOrder[architecture] = stage.ArchitectureMap_Staged_Order[architecture]
			newInstancesReverseSlice = append(newInstancesReverseSlice, architecture.GongMarshallUnstaging(stage))
			delete(stage.Architectures_referenceOrder, architecture)
			fieldInitializers, pointersInitializations := architecture.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ArchitectureMap_Staged_Order[ref] = stage.ArchitectureMap_Staged_Order[architecture]
			diffs := architecture.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, architecture)
			delete(stage.ArchitectureMap_Staged_Order, ref)
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
	for ref := range stage.Architectures_reference {
		if _, ok := stage.Architectures[ref]; !ok {
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
			stage.Diagrams_referenceOrder[diagram] = stage.DiagramMap_Staged_Order[diagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagram.GongMarshallUnstaging(stage))
			delete(stage.Diagrams_referenceOrder, diagram)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramMap_Staged_Order[ref] = stage.DiagramMap_Staged_Order[diagram]
			diffs := diagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagram)
			delete(stage.DiagramMap_Staged_Order, ref)
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
	for ref := range stage.Diagrams_reference {
		if _, ok := stage.Diagrams[ref]; !ok {
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
			stage.Guards_referenceOrder[guard] = stage.GuardMap_Staged_Order[guard]
			newInstancesReverseSlice = append(newInstancesReverseSlice, guard.GongMarshallUnstaging(stage))
			delete(stage.Guards_referenceOrder, guard)
			fieldInitializers, pointersInitializations := guard.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GuardMap_Staged_Order[ref] = stage.GuardMap_Staged_Order[guard]
			diffs := guard.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, guard)
			delete(stage.GuardMap_Staged_Order, ref)
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
	for ref := range stage.Guards_reference {
		if _, ok := stage.Guards[ref]; !ok {
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
			stage.Kills_referenceOrder[kill] = stage.KillMap_Staged_Order[kill]
			newInstancesReverseSlice = append(newInstancesReverseSlice, kill.GongMarshallUnstaging(stage))
			delete(stage.Kills_referenceOrder, kill)
			fieldInitializers, pointersInitializations := kill.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.KillMap_Staged_Order[ref] = stage.KillMap_Staged_Order[kill]
			diffs := kill.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, kill)
			delete(stage.KillMap_Staged_Order, ref)
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
	for ref := range stage.Kills_reference {
		if _, ok := stage.Kills[ref]; !ok {
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
			stage.Messages_referenceOrder[message] = stage.MessageMap_Staged_Order[message]
			newInstancesReverseSlice = append(newInstancesReverseSlice, message.GongMarshallUnstaging(stage))
			delete(stage.Messages_referenceOrder, message)
			fieldInitializers, pointersInitializations := message.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MessageMap_Staged_Order[ref] = stage.MessageMap_Staged_Order[message]
			diffs := message.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, message)
			delete(stage.MessageMap_Staged_Order, ref)
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
	for ref := range stage.Messages_reference {
		if _, ok := stage.Messages[ref]; !ok {
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
			stage.MessageTypes_referenceOrder[messagetype] = stage.MessageTypeMap_Staged_Order[messagetype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, messagetype.GongMarshallUnstaging(stage))
			delete(stage.MessageTypes_referenceOrder, messagetype)
			fieldInitializers, pointersInitializations := messagetype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MessageTypeMap_Staged_Order[ref] = stage.MessageTypeMap_Staged_Order[messagetype]
			diffs := messagetype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, messagetype)
			delete(stage.MessageTypeMap_Staged_Order, ref)
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
	for ref := range stage.MessageTypes_reference {
		if _, ok := stage.MessageTypes[ref]; !ok {
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
			stage.Objects_referenceOrder[object] = stage.ObjectMap_Staged_Order[object]
			newInstancesReverseSlice = append(newInstancesReverseSlice, object.GongMarshallUnstaging(stage))
			delete(stage.Objects_referenceOrder, object)
			fieldInitializers, pointersInitializations := object.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ObjectMap_Staged_Order[ref] = stage.ObjectMap_Staged_Order[object]
			diffs := object.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, object)
			delete(stage.ObjectMap_Staged_Order, ref)
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
	for ref := range stage.Objects_reference {
		if _, ok := stage.Objects[ref]; !ok {
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
			stage.Roles_referenceOrder[role] = stage.RoleMap_Staged_Order[role]
			newInstancesReverseSlice = append(newInstancesReverseSlice, role.GongMarshallUnstaging(stage))
			delete(stage.Roles_referenceOrder, role)
			fieldInitializers, pointersInitializations := role.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RoleMap_Staged_Order[ref] = stage.RoleMap_Staged_Order[role]
			diffs := role.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, role)
			delete(stage.RoleMap_Staged_Order, ref)
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
	for ref := range stage.Roles_reference {
		if _, ok := stage.Roles[ref]; !ok {
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
			stage.States_referenceOrder[state] = stage.StateMap_Staged_Order[state]
			newInstancesReverseSlice = append(newInstancesReverseSlice, state.GongMarshallUnstaging(stage))
			delete(stage.States_referenceOrder, state)
			fieldInitializers, pointersInitializations := state.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StateMap_Staged_Order[ref] = stage.StateMap_Staged_Order[state]
			diffs := state.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, state)
			delete(stage.StateMap_Staged_Order, ref)
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
	for ref := range stage.States_reference {
		if _, ok := stage.States[ref]; !ok {
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
			stage.StateMachines_referenceOrder[statemachine] = stage.StateMachineMap_Staged_Order[statemachine]
			newInstancesReverseSlice = append(newInstancesReverseSlice, statemachine.GongMarshallUnstaging(stage))
			delete(stage.StateMachines_referenceOrder, statemachine)
			fieldInitializers, pointersInitializations := statemachine.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StateMachineMap_Staged_Order[ref] = stage.StateMachineMap_Staged_Order[statemachine]
			diffs := statemachine.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, statemachine)
			delete(stage.StateMachineMap_Staged_Order, ref)
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
	for ref := range stage.StateMachines_reference {
		if _, ok := stage.StateMachines[ref]; !ok {
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
			stage.StateShapes_referenceOrder[stateshape] = stage.StateShapeMap_Staged_Order[stateshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, stateshape.GongMarshallUnstaging(stage))
			delete(stage.StateShapes_referenceOrder, stateshape)
			fieldInitializers, pointersInitializations := stateshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.StateShapeMap_Staged_Order[ref] = stage.StateShapeMap_Staged_Order[stateshape]
			diffs := stateshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, stateshape)
			delete(stage.StateShapeMap_Staged_Order, ref)
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
	for ref := range stage.StateShapes_reference {
		if _, ok := stage.StateShapes[ref]; !ok {
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
			stage.Transitions_referenceOrder[transition] = stage.TransitionMap_Staged_Order[transition]
			newInstancesReverseSlice = append(newInstancesReverseSlice, transition.GongMarshallUnstaging(stage))
			delete(stage.Transitions_referenceOrder, transition)
			fieldInitializers, pointersInitializations := transition.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TransitionMap_Staged_Order[ref] = stage.TransitionMap_Staged_Order[transition]
			diffs := transition.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, transition)
			delete(stage.TransitionMap_Staged_Order, ref)
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
	for ref := range stage.Transitions_reference {
		if _, ok := stage.Transitions[ref]; !ok {
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
			stage.Transition_Shapes_referenceOrder[transition_shape] = stage.Transition_ShapeMap_Staged_Order[transition_shape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, transition_shape.GongMarshallUnstaging(stage))
			delete(stage.Transition_Shapes_referenceOrder, transition_shape)
			fieldInitializers, pointersInitializations := transition_shape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Transition_ShapeMap_Staged_Order[ref] = stage.Transition_ShapeMap_Staged_Order[transition_shape]
			diffs := transition_shape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, transition_shape)
			delete(stage.Transition_ShapeMap_Staged_Order, ref)
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
	for ref := range stage.Transition_Shapes_reference {
		if _, ok := stage.Transition_Shapes[ref]; !ok {
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

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {

	// insertion point per named struct
	stage.Actions_reference = make(map[*Action]*Action)
	stage.Actions_referenceOrder = make(map[*Action]uint) // diff Unstage needs the reference order
	for instance := range stage.Actions {
		stage.Actions_reference[instance] = instance.GongCopy().(*Action)
		stage.Actions_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Activitiess_reference = make(map[*Activities]*Activities)
	stage.Activitiess_referenceOrder = make(map[*Activities]uint) // diff Unstage needs the reference order
	for instance := range stage.Activitiess {
		stage.Activitiess_reference[instance] = instance.GongCopy().(*Activities)
		stage.Activitiess_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Architectures_reference = make(map[*Architecture]*Architecture)
	stage.Architectures_referenceOrder = make(map[*Architecture]uint) // diff Unstage needs the reference order
	for instance := range stage.Architectures {
		stage.Architectures_reference[instance] = instance.GongCopy().(*Architecture)
		stage.Architectures_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	for instance := range stage.Diagrams {
		stage.Diagrams_reference[instance] = instance.GongCopy().(*Diagram)
		stage.Diagrams_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Guards_reference = make(map[*Guard]*Guard)
	stage.Guards_referenceOrder = make(map[*Guard]uint) // diff Unstage needs the reference order
	for instance := range stage.Guards {
		stage.Guards_reference[instance] = instance.GongCopy().(*Guard)
		stage.Guards_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Kills_reference = make(map[*Kill]*Kill)
	stage.Kills_referenceOrder = make(map[*Kill]uint) // diff Unstage needs the reference order
	for instance := range stage.Kills {
		stage.Kills_reference[instance] = instance.GongCopy().(*Kill)
		stage.Kills_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Messages_reference = make(map[*Message]*Message)
	stage.Messages_referenceOrder = make(map[*Message]uint) // diff Unstage needs the reference order
	for instance := range stage.Messages {
		stage.Messages_reference[instance] = instance.GongCopy().(*Message)
		stage.Messages_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.MessageTypes_reference = make(map[*MessageType]*MessageType)
	stage.MessageTypes_referenceOrder = make(map[*MessageType]uint) // diff Unstage needs the reference order
	for instance := range stage.MessageTypes {
		stage.MessageTypes_reference[instance] = instance.GongCopy().(*MessageType)
		stage.MessageTypes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Objects_reference = make(map[*Object]*Object)
	stage.Objects_referenceOrder = make(map[*Object]uint) // diff Unstage needs the reference order
	for instance := range stage.Objects {
		stage.Objects_reference[instance] = instance.GongCopy().(*Object)
		stage.Objects_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Roles_reference = make(map[*Role]*Role)
	stage.Roles_referenceOrder = make(map[*Role]uint) // diff Unstage needs the reference order
	for instance := range stage.Roles {
		stage.Roles_reference[instance] = instance.GongCopy().(*Role)
		stage.Roles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.States_reference = make(map[*State]*State)
	stage.States_referenceOrder = make(map[*State]uint) // diff Unstage needs the reference order
	for instance := range stage.States {
		stage.States_reference[instance] = instance.GongCopy().(*State)
		stage.States_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.StateMachines_reference = make(map[*StateMachine]*StateMachine)
	stage.StateMachines_referenceOrder = make(map[*StateMachine]uint) // diff Unstage needs the reference order
	for instance := range stage.StateMachines {
		stage.StateMachines_reference[instance] = instance.GongCopy().(*StateMachine)
		stage.StateMachines_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.StateShapes_reference = make(map[*StateShape]*StateShape)
	stage.StateShapes_referenceOrder = make(map[*StateShape]uint) // diff Unstage needs the reference order
	for instance := range stage.StateShapes {
		stage.StateShapes_reference[instance] = instance.GongCopy().(*StateShape)
		stage.StateShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Transitions_reference = make(map[*Transition]*Transition)
	stage.Transitions_referenceOrder = make(map[*Transition]uint) // diff Unstage needs the reference order
	for instance := range stage.Transitions {
		stage.Transitions_reference[instance] = instance.GongCopy().(*Transition)
		stage.Transitions_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Transition_Shapes_reference = make(map[*Transition_Shape]*Transition_Shape)
	stage.Transition_Shapes_referenceOrder = make(map[*Transition_Shape]uint) // diff Unstage needs the reference order
	for instance := range stage.Transition_Shapes {
		stage.Transition_Shapes_reference[instance] = instance.GongCopy().(*Transition_Shape)
		stage.Transition_Shapes_referenceOrder[instance] = instance.GongGetOrder(stage)
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
	if order, ok := stage.ActionMap_Staged_Order[action]; ok {
		return order
	}
	return stage.Actions_referenceOrder[action]
}

func (action *Action) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Actions_referenceOrder[action]
}

func (activities *Activities) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ActivitiesMap_Staged_Order[activities]; ok {
		return order
	}
	return stage.Activitiess_referenceOrder[activities]
}

func (activities *Activities) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Activitiess_referenceOrder[activities]
}

func (architecture *Architecture) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ArchitectureMap_Staged_Order[architecture]; ok {
		return order
	}
	return stage.Architectures_referenceOrder[architecture]
}

func (architecture *Architecture) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Architectures_referenceOrder[architecture]
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramMap_Staged_Order[diagram]; ok {
		return order
	}
	return stage.Diagrams_referenceOrder[diagram]
}

func (diagram *Diagram) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Diagrams_referenceOrder[diagram]
}

func (guard *Guard) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GuardMap_Staged_Order[guard]; ok {
		return order
	}
	return stage.Guards_referenceOrder[guard]
}

func (guard *Guard) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Guards_referenceOrder[guard]
}

func (kill *Kill) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.KillMap_Staged_Order[kill]; ok {
		return order
	}
	return stage.Kills_referenceOrder[kill]
}

func (kill *Kill) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Kills_referenceOrder[kill]
}

func (message *Message) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MessageMap_Staged_Order[message]; ok {
		return order
	}
	return stage.Messages_referenceOrder[message]
}

func (message *Message) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Messages_referenceOrder[message]
}

func (messagetype *MessageType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MessageTypeMap_Staged_Order[messagetype]; ok {
		return order
	}
	return stage.MessageTypes_referenceOrder[messagetype]
}

func (messagetype *MessageType) GongGetReferenceOrder(stage *Stage) uint {
	return stage.MessageTypes_referenceOrder[messagetype]
}

func (object *Object) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ObjectMap_Staged_Order[object]; ok {
		return order
	}
	return stage.Objects_referenceOrder[object]
}

func (object *Object) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Objects_referenceOrder[object]
}

func (role *Role) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RoleMap_Staged_Order[role]; ok {
		return order
	}
	return stage.Roles_referenceOrder[role]
}

func (role *Role) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Roles_referenceOrder[role]
}

func (state *State) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StateMap_Staged_Order[state]; ok {
		return order
	}
	return stage.States_referenceOrder[state]
}

func (state *State) GongGetReferenceOrder(stage *Stage) uint {
	return stage.States_referenceOrder[state]
}

func (statemachine *StateMachine) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StateMachineMap_Staged_Order[statemachine]; ok {
		return order
	}
	return stage.StateMachines_referenceOrder[statemachine]
}

func (statemachine *StateMachine) GongGetReferenceOrder(stage *Stage) uint {
	return stage.StateMachines_referenceOrder[statemachine]
}

func (stateshape *StateShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StateShapeMap_Staged_Order[stateshape]; ok {
		return order
	}
	return stage.StateShapes_referenceOrder[stateshape]
}

func (stateshape *StateShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.StateShapes_referenceOrder[stateshape]
}

func (transition *Transition) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TransitionMap_Staged_Order[transition]; ok {
		return order
	}
	return stage.Transitions_referenceOrder[transition]
}

func (transition *Transition) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Transitions_referenceOrder[transition]
}

func (transition_shape *Transition_Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Transition_ShapeMap_Staged_Order[transition_shape]; ok {
		return order
	}
	return stage.Transition_Shapes_referenceOrder[transition_shape]
}

func (transition_shape *Transition_Shape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Transition_Shapes_referenceOrder[transition_shape]
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
	return fmt.Sprintf("__%s__%08d_", action.GongGetGongstructName(), action.GongGetReferenceOrder(stage))
}

func (activities *Activities) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", activities.GongGetGongstructName(), activities.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (activities *Activities) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", activities.GongGetGongstructName(), activities.GongGetReferenceOrder(stage))
}

func (architecture *Architecture) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", architecture.GongGetGongstructName(), architecture.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (architecture *Architecture) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", architecture.GongGetGongstructName(), architecture.GongGetReferenceOrder(stage))
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetReferenceOrder(stage))
}

func (guard *Guard) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", guard.GongGetGongstructName(), guard.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (guard *Guard) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", guard.GongGetGongstructName(), guard.GongGetReferenceOrder(stage))
}

func (kill *Kill) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", kill.GongGetGongstructName(), kill.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (kill *Kill) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", kill.GongGetGongstructName(), kill.GongGetReferenceOrder(stage))
}

func (message *Message) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", message.GongGetGongstructName(), message.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (message *Message) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", message.GongGetGongstructName(), message.GongGetReferenceOrder(stage))
}

func (messagetype *MessageType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", messagetype.GongGetGongstructName(), messagetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (messagetype *MessageType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", messagetype.GongGetGongstructName(), messagetype.GongGetReferenceOrder(stage))
}

func (object *Object) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", object.GongGetGongstructName(), object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (object *Object) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", object.GongGetGongstructName(), object.GongGetReferenceOrder(stage))
}

func (role *Role) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", role.GongGetGongstructName(), role.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (role *Role) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", role.GongGetGongstructName(), role.GongGetReferenceOrder(stage))
}

func (state *State) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", state.GongGetGongstructName(), state.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (state *State) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", state.GongGetGongstructName(), state.GongGetReferenceOrder(stage))
}

func (statemachine *StateMachine) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", statemachine.GongGetGongstructName(), statemachine.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (statemachine *StateMachine) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", statemachine.GongGetGongstructName(), statemachine.GongGetReferenceOrder(stage))
}

func (stateshape *StateShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stateshape.GongGetGongstructName(), stateshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stateshape *StateShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stateshape.GongGetGongstructName(), stateshape.GongGetReferenceOrder(stage))
}

func (transition *Transition) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transition.GongGetGongstructName(), transition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transition *Transition) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transition.GongGetGongstructName(), transition.GongGetReferenceOrder(stage))
}

func (transition_shape *Transition_Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transition_shape.GongGetGongstructName(), transition_shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transition_shape *Transition_Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transition_shape.GongGetGongstructName(), transition_shape.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (action *Action) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", action.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Action")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", action.Name)
	return
}
func (activities *Activities) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", activities.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Activities")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", activities.Name)
	return
}
func (architecture *Architecture) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", architecture.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Architecture")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", architecture.Name)
	return
}
func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", diagram.Name)
	return
}
func (guard *Guard) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", guard.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Guard")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", guard.Name)
	return
}
func (kill *Kill) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", kill.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Kill")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", kill.Name)
	return
}
func (message *Message) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", message.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Message")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", message.Name)
	return
}
func (messagetype *MessageType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", messagetype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MessageType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", messagetype.Name)
	return
}
func (object *Object) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", object.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Object")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", object.Name)
	return
}
func (role *Role) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", role.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Role")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", role.Name)
	return
}
func (state *State) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", state.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "State")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", state.Name)
	return
}
func (statemachine *StateMachine) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", statemachine.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StateMachine")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", statemachine.Name)
	return
}
func (stateshape *StateShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stateshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StateShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", stateshape.Name)
	return
}
func (transition *Transition) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transition.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transition")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", transition.Name)
	return
}
func (transition_shape *Transition_Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transition_shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transition_Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", transition_shape.Name)
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
