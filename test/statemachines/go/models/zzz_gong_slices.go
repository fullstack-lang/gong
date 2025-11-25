// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Architecture
	// insertion point per field
	clear(stage.Architecture_StateMachines_reverseMap)
	stage.Architecture_StateMachines_reverseMap = make(map[*StateMachine]*Architecture)
	for architecture := range stage.Architectures {
		_ = architecture
		for _, _statemachine := range architecture.StateMachines {
			stage.Architecture_StateMachines_reverseMap[_statemachine] = architecture
		}
	}
	clear(stage.Architecture_Roles_reverseMap)
	stage.Architecture_Roles_reverseMap = make(map[*Role]*Architecture)
	for architecture := range stage.Architectures {
		_ = architecture
		for _, _role := range architecture.Roles {
			stage.Architecture_Roles_reverseMap[_role] = architecture
		}
	}

	// Compute reverse map for named struct Diagram
	// insertion point per field
	clear(stage.Diagram_State_Shapes_reverseMap)
	stage.Diagram_State_Shapes_reverseMap = make(map[*StateShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _stateshape := range diagram.State_Shapes {
			stage.Diagram_State_Shapes_reverseMap[_stateshape] = diagram
		}
	}
	clear(stage.Diagram_Transition_Shapes_reverseMap)
	stage.Diagram_Transition_Shapes_reverseMap = make(map[*Transition_Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _transition_shape := range diagram.Transition_Shapes {
			stage.Diagram_Transition_Shapes_reverseMap[_transition_shape] = diagram
		}
	}

	// Compute reverse map for named struct Kill
	// insertion point per field

	// Compute reverse map for named struct Message
	// insertion point per field

	// Compute reverse map for named struct MessageType
	// insertion point per field

	// Compute reverse map for named struct Object
	// insertion point per field
	clear(stage.Object_Messages_reverseMap)
	stage.Object_Messages_reverseMap = make(map[*Message]*Object)
	for object := range stage.Objects {
		_ = object
		for _, _message := range object.Messages {
			stage.Object_Messages_reverseMap[_message] = object
		}
	}

	// Compute reverse map for named struct Role
	// insertion point per field
	clear(stage.Role_RolesWithSamePermissions_reverseMap)
	stage.Role_RolesWithSamePermissions_reverseMap = make(map[*Role]*Role)
	for role := range stage.Roles {
		_ = role
		for _, _role := range role.RolesWithSamePermissions {
			stage.Role_RolesWithSamePermissions_reverseMap[_role] = role
		}
	}

	// Compute reverse map for named struct State
	// insertion point per field
	clear(stage.State_SubStates_reverseMap)
	stage.State_SubStates_reverseMap = make(map[*State]*State)
	for state := range stage.States {
		_ = state
		for _, _state := range state.SubStates {
			stage.State_SubStates_reverseMap[_state] = state
		}
	}
	clear(stage.State_Diagrams_reverseMap)
	stage.State_Diagrams_reverseMap = make(map[*Diagram]*State)
	for state := range stage.States {
		_ = state
		for _, _diagram := range state.Diagrams {
			stage.State_Diagrams_reverseMap[_diagram] = state
		}
	}

	// Compute reverse map for named struct StateMachine
	// insertion point per field
	clear(stage.StateMachine_States_reverseMap)
	stage.StateMachine_States_reverseMap = make(map[*State]*StateMachine)
	for statemachine := range stage.StateMachines {
		_ = statemachine
		for _, _state := range statemachine.States {
			stage.StateMachine_States_reverseMap[_state] = statemachine
		}
	}
	clear(stage.StateMachine_Diagrams_reverseMap)
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
	clear(stage.Transition_RolesWithPermissions_reverseMap)
	stage.Transition_RolesWithPermissions_reverseMap = make(map[*Role]*Transition)
	for transition := range stage.Transitions {
		_ = transition
		for _, _role := range transition.RolesWithPermissions {
			stage.Transition_RolesWithPermissions_reverseMap[_role] = transition
		}
	}
	clear(stage.Transition_GeneratedMessages_reverseMap)
	stage.Transition_GeneratedMessages_reverseMap = make(map[*MessageType]*Transition)
	for transition := range stage.Transitions {
		_ = transition
		for _, _messagetype := range transition.GeneratedMessages {
			stage.Transition_GeneratedMessages_reverseMap[_messagetype] = transition
		}
	}
	clear(stage.Transition_Diagrams_reverseMap)
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
	for instance := range stage.Architectures {
		res = append(res, instance)
	}

	for instance := range stage.Diagrams {
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
func (architecture *Architecture) GongCopy() GongstructIF {
	newInstance := *architecture
	return &newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := *diagram
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

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
