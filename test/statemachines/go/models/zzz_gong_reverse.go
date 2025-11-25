// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *Architecture:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Diagram:
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

	case *Kill:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Message:
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

	case *MessageType:
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

	case *Object:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Role:
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

	case *State:
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

	case *StateMachine:
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

	case *StateShape:
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

	case *Transition:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Transition_Shape:
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

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *Architecture:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Diagram:
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

	case *Kill:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Message:
		switch reverseField.GongstructName {
		// insertion point
		case "Object":
			switch reverseField.Fieldname {
			case "Messages":
				res = stage.Object_Messages_reverseMap[inst]
			}
		}

	case *MessageType:
		switch reverseField.GongstructName {
		// insertion point
		case "Transition":
			switch reverseField.Fieldname {
			case "GeneratedMessages":
				res = stage.Transition_GeneratedMessages_reverseMap[inst]
			}
		}

	case *Object:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Role:
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

	case *State:
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

	case *StateMachine:
		switch reverseField.GongstructName {
		// insertion point
		case "Architecture":
			switch reverseField.Fieldname {
			case "StateMachines":
				res = stage.Architecture_StateMachines_reverseMap[inst]
			}
		}

	case *StateShape:
		switch reverseField.GongstructName {
		// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "State_Shapes":
				res = stage.Diagram_State_Shapes_reverseMap[inst]
			}
		}

	case *Transition:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Transition_Shape:
		switch reverseField.GongstructName {
		// insertion point
		case "Diagram":
			switch reverseField.Fieldname {
			case "Transition_Shapes":
				res = stage.Diagram_Transition_Shapes_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
