// generated code - do not edit
package models

// insertion point
func (inst *Action) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Activities) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "State":
			switch reverseField.Fieldname {
			case "Activities":
				if _state, ok := stage.State_Activities_reverseMap[inst]; ok {
					res = _state.Name
				}
			}
	}
	return
}

func (inst *Architecture) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "State":
			switch reverseField.Fieldname {
			case "Diagrams":
				if _state, ok := stage.State_Diagrams_reverseMap[inst]; ok {
					res = _state.Name
				}
			}
		case "StateMachine":
			switch reverseField.Fieldname {
			case "Diagrams":
				if _statemachine, ok := stage.StateMachine_Diagrams_reverseMap[inst]; ok {
					res = _statemachine.Name
				}
			}
		case "Transition":
			switch reverseField.Fieldname {
			case "Diagrams":
				if _transition, ok := stage.Transition_Diagrams_reverseMap[inst]; ok {
					res = _transition.Name
				}
			}
	}
	return
}

func (inst *Guard) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Kill) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Message) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Object":
			switch reverseField.Fieldname {
			case "Messages":
				if _object, ok := stage.Object_Messages_reverseMap[inst]; ok {
					res = _object.Name
				}
			}
	}
	return
}

func (inst *MessageType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Transition":
			switch reverseField.Fieldname {
			case "GeneratedMessages":
				if _transition, ok := stage.Transition_GeneratedMessages_reverseMap[inst]; ok {
					res = _transition.Name
				}
			}
	}
	return
}

func (inst *Object) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Role) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Architecture":
			switch reverseField.Fieldname {
			case "Roles":
				if _architecture, ok := stage.Architecture_Roles_reverseMap[inst]; ok {
					res = _architecture.Name
				}
			}
		case "Role":
			switch reverseField.Fieldname {
			case "RolesWithSamePermissions":
				if _role, ok := stage.Role_RolesWithSamePermissions_reverseMap[inst]; ok {
					res = _role.Name
				}
			}
		case "Transition":
			switch reverseField.Fieldname {
			case "RolesWithPermissions":
				if _transition, ok := stage.Transition_RolesWithPermissions_reverseMap[inst]; ok {
					res = _transition.Name
				}
			}
	}
	return
}

func (inst *State) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "State":
			switch reverseField.Fieldname {
			case "SubStates":
				if _state, ok := stage.State_SubStates_reverseMap[inst]; ok {
					res = _state.Name
				}
			}
		case "StateMachine":
			switch reverseField.Fieldname {
			case "States":
				if _statemachine, ok := stage.StateMachine_States_reverseMap[inst]; ok {
					res = _statemachine.Name
				}
			}
	}
	return
}

func (inst *StateMachine) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Architecture":
			switch reverseField.Fieldname {
			case "StateMachines":
				if _architecture, ok := stage.Architecture_StateMachines_reverseMap[inst]; ok {
					res = _architecture.Name
				}
			}
	}
	return
}

func (inst *StateShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "State_Shapes":
				if _diagram, ok := stage.Diagram_State_Shapes_reverseMap[inst]; ok {
					res = _diagram.Name
				}
			}
	}
	return
}

func (inst *Transition) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Transition_Shape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Transition_Shapes":
				if _diagram, ok := stage.Diagram_Transition_Shapes_reverseMap[inst]; ok {
					res = _diagram.Name
				}
			}
	}
	return
}


// insertion point
func (inst *Action) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Activities) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "State":
			switch reverseField.Fieldname {
			case "Activities":
				res = stage.State_Activities_reverseMap[inst]
			}
	}
	return res
}

func (inst *Architecture) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Diagram) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "State":
			switch reverseField.Fieldname {
			case "Diagrams":
				res = stage.State_Diagrams_reverseMap[inst]
			}
		case "StateMachine":
			switch reverseField.Fieldname {
			case "Diagrams":
				res = stage.StateMachine_Diagrams_reverseMap[inst]
			}
		case "Transition":
			switch reverseField.Fieldname {
			case "Diagrams":
				res = stage.Transition_Diagrams_reverseMap[inst]
			}
	}
	return res
}

func (inst *Guard) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Kill) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Message) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Object":
			switch reverseField.Fieldname {
			case "Messages":
				res = stage.Object_Messages_reverseMap[inst]
			}
	}
	return res
}

func (inst *MessageType) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Transition":
			switch reverseField.Fieldname {
			case "GeneratedMessages":
				res = stage.Transition_GeneratedMessages_reverseMap[inst]
			}
	}
	return res
}

func (inst *Object) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Role) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Architecture":
			switch reverseField.Fieldname {
			case "Roles":
				res = stage.Architecture_Roles_reverseMap[inst]
			}
		case "Role":
			switch reverseField.Fieldname {
			case "RolesWithSamePermissions":
				res = stage.Role_RolesWithSamePermissions_reverseMap[inst]
			}
		case "Transition":
			switch reverseField.Fieldname {
			case "RolesWithPermissions":
				res = stage.Transition_RolesWithPermissions_reverseMap[inst]
			}
	}
	return res
}

func (inst *State) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "State":
			switch reverseField.Fieldname {
			case "SubStates":
				res = stage.State_SubStates_reverseMap[inst]
			}
		case "StateMachine":
			switch reverseField.Fieldname {
			case "States":
				res = stage.StateMachine_States_reverseMap[inst]
			}
	}
	return res
}

func (inst *StateMachine) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Architecture":
			switch reverseField.Fieldname {
			case "StateMachines":
				res = stage.Architecture_StateMachines_reverseMap[inst]
			}
	}
	return res
}

func (inst *StateShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "State_Shapes":
				res = stage.Diagram_State_Shapes_reverseMap[inst]
			}
	}
	return res
}

func (inst *Transition) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Transition_Shape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Transition_Shapes":
				res = stage.Diagram_Transition_Shapes_reverseMap[inst]
			}
	}
	return res
}

