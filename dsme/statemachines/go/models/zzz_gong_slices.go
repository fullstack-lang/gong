// generated code - do not edit
package models

import (
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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	// insertion point per named struct
	var actions_newInstances []*Action
	var actions_deletedInstances []*Action

	// parse all staged instances and check if they have a reference
	for action := range stage.Actions {
		if ref, ok := stage.Actions_reference[action]; !ok {
			actions_newInstances = append(actions_newInstances, action)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Action "+action.Name,
				)
			}
		} else {
			diffs := action.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Action \""+action.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for action := range stage.Actions_reference {
		if _, ok := stage.Actions[action]; !ok {
			actions_deletedInstances = append(actions_deletedInstances, action)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Action "+action.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Activities "+activities.Name,
				)
			}
		} else {
			diffs := activities.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Activities \""+activities.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for activities := range stage.Activitiess_reference {
		if _, ok := stage.Activitiess[activities]; !ok {
			activitiess_deletedInstances = append(activitiess_deletedInstances, activities)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Activities "+activities.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Architecture "+architecture.Name,
				)
			}
		} else {
			diffs := architecture.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Architecture \""+architecture.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for architecture := range stage.Architectures_reference {
		if _, ok := stage.Architectures[architecture]; !ok {
			architectures_deletedInstances = append(architectures_deletedInstances, architecture)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Architecture "+architecture.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Diagram "+diagram.Name,
				)
			}
		} else {
			diffs := diagram.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Diagram \""+diagram.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for diagram := range stage.Diagrams_reference {
		if _, ok := stage.Diagrams[diagram]; !ok {
			diagrams_deletedInstances = append(diagrams_deletedInstances, diagram)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Diagram "+diagram.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Guard "+guard.Name,
				)
			}
		} else {
			diffs := guard.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Guard \""+guard.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for guard := range stage.Guards_reference {
		if _, ok := stage.Guards[guard]; !ok {
			guards_deletedInstances = append(guards_deletedInstances, guard)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Guard "+guard.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Kill "+kill.Name,
				)
			}
		} else {
			diffs := kill.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Kill \""+kill.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for kill := range stage.Kills_reference {
		if _, ok := stage.Kills[kill]; !ok {
			kills_deletedInstances = append(kills_deletedInstances, kill)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Kill "+kill.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Message "+message.Name,
				)
			}
		} else {
			diffs := message.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Message \""+message.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for message := range stage.Messages_reference {
		if _, ok := stage.Messages[message]; !ok {
			messages_deletedInstances = append(messages_deletedInstances, message)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Message "+message.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of MessageType "+messagetype.Name,
				)
			}
		} else {
			diffs := messagetype.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of MessageType \""+messagetype.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for messagetype := range stage.MessageTypes_reference {
		if _, ok := stage.MessageTypes[messagetype]; !ok {
			messagetypes_deletedInstances = append(messagetypes_deletedInstances, messagetype)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of MessageType "+messagetype.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Object "+object.Name,
				)
			}
		} else {
			diffs := object.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Object \""+object.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for object := range stage.Objects_reference {
		if _, ok := stage.Objects[object]; !ok {
			objects_deletedInstances = append(objects_deletedInstances, object)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Object "+object.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Role "+role.Name,
				)
			}
		} else {
			diffs := role.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Role \""+role.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for role := range stage.Roles_reference {
		if _, ok := stage.Roles[role]; !ok {
			roles_deletedInstances = append(roles_deletedInstances, role)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Role "+role.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of State "+state.Name,
				)
			}
		} else {
			diffs := state.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of State \""+state.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for state := range stage.States_reference {
		if _, ok := stage.States[state]; !ok {
			states_deletedInstances = append(states_deletedInstances, state)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of State "+state.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of StateMachine "+statemachine.Name,
				)
			}
		} else {
			diffs := statemachine.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of StateMachine \""+statemachine.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for statemachine := range stage.StateMachines_reference {
		if _, ok := stage.StateMachines[statemachine]; !ok {
			statemachines_deletedInstances = append(statemachines_deletedInstances, statemachine)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of StateMachine "+statemachine.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of StateShape "+stateshape.Name,
				)
			}
		} else {
			diffs := stateshape.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of StateShape \""+stateshape.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for stateshape := range stage.StateShapes_reference {
		if _, ok := stage.StateShapes[stateshape]; !ok {
			stateshapes_deletedInstances = append(stateshapes_deletedInstances, stateshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of StateShape "+stateshape.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Transition "+transition.Name,
				)
			}
		} else {
			diffs := transition.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Transition \""+transition.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for transition := range stage.Transitions_reference {
		if _, ok := stage.Transitions[transition]; !ok {
			transitions_deletedInstances = append(transitions_deletedInstances, transition)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Transition "+transition.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Transition_Shape "+transition_shape.Name,
				)
			}
		} else {
			diffs := transition_shape.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Transition_Shape \""+transition_shape.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for transition_shape := range stage.Transition_Shapes_reference {
		if _, ok := stage.Transition_Shapes[transition_shape]; !ok {
			transition_shapes_deletedInstances = append(transition_shapes_deletedInstances, transition_shape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Transition_Shape "+transition_shape.Name,
				)
			}
		}
	}

	lenNewInstances += len(transition_shapes_newInstances)
	lenDeletedInstances += len(transition_shapes_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Actions_reference = make(map[*Action]*Action)
	for instance := range stage.Actions {
		stage.Actions_reference[instance] = instance.GongCopy().(*Action)
	}

	stage.Activitiess_reference = make(map[*Activities]*Activities)
	for instance := range stage.Activitiess {
		stage.Activitiess_reference[instance] = instance.GongCopy().(*Activities)
	}

	stage.Architectures_reference = make(map[*Architecture]*Architecture)
	for instance := range stage.Architectures {
		stage.Architectures_reference[instance] = instance.GongCopy().(*Architecture)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		stage.Diagrams_reference[instance] = instance.GongCopy().(*Diagram)
	}

	stage.Guards_reference = make(map[*Guard]*Guard)
	for instance := range stage.Guards {
		stage.Guards_reference[instance] = instance.GongCopy().(*Guard)
	}

	stage.Kills_reference = make(map[*Kill]*Kill)
	for instance := range stage.Kills {
		stage.Kills_reference[instance] = instance.GongCopy().(*Kill)
	}

	stage.Messages_reference = make(map[*Message]*Message)
	for instance := range stage.Messages {
		stage.Messages_reference[instance] = instance.GongCopy().(*Message)
	}

	stage.MessageTypes_reference = make(map[*MessageType]*MessageType)
	for instance := range stage.MessageTypes {
		stage.MessageTypes_reference[instance] = instance.GongCopy().(*MessageType)
	}

	stage.Objects_reference = make(map[*Object]*Object)
	for instance := range stage.Objects {
		stage.Objects_reference[instance] = instance.GongCopy().(*Object)
	}

	stage.Roles_reference = make(map[*Role]*Role)
	for instance := range stage.Roles {
		stage.Roles_reference[instance] = instance.GongCopy().(*Role)
	}

	stage.States_reference = make(map[*State]*State)
	for instance := range stage.States {
		stage.States_reference[instance] = instance.GongCopy().(*State)
	}

	stage.StateMachines_reference = make(map[*StateMachine]*StateMachine)
	for instance := range stage.StateMachines {
		stage.StateMachines_reference[instance] = instance.GongCopy().(*StateMachine)
	}

	stage.StateShapes_reference = make(map[*StateShape]*StateShape)
	for instance := range stage.StateShapes {
		stage.StateShapes_reference[instance] = instance.GongCopy().(*StateShape)
	}

	stage.Transitions_reference = make(map[*Transition]*Transition)
	for instance := range stage.Transitions {
		stage.Transitions_reference[instance] = instance.GongCopy().(*Transition)
	}

	stage.Transition_Shapes_reference = make(map[*Transition_Shape]*Transition_Shape)
	for instance := range stage.Transition_Shapes {
		stage.Transition_Shapes_reference[instance] = instance.GongCopy().(*Transition_Shape)
	}

}
