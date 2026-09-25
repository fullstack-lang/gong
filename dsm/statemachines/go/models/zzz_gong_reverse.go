// generated code - do not edit
package models

// insertion point
func (inst *Action) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Activities) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Diagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Diagrams":
			if _library, ok := stage.Library_Diagrams_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
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

func (inst *Guard) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Kill) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			if _library, ok := stage.Library_SubLibraries_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "SubLibrariesWhoseNodeIsExpanded":
			if _library, ok := stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *Message) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *MessageType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "MessageTypes":
			if _library, ok := stage.Library_MessageTypes_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
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

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "State":
		switch reverseField.Fieldname {
		case "Notes":
			if _state, ok := stage.State_Notes_reverseMap[inst]; ok {
				res = _state.Name
			}
		}
	}
	return
}

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			if _diagram, ok := stage.Diagram_Note_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *NoteStateShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "NoteState_Shapes":
			if _diagram, ok := stage.Diagram_NoteState_Shapes_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
	}
	return
}

func (inst *Object) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Role) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Roles":
			if _library, ok := stage.Library_Roles_reverseMap[inst]; ok {
				res = _library.Name
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

func (inst *State) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Diagram":
		switch reverseField.Fieldname {
		case "StatesWhoseNodeIsExpanded":
			if _diagram, ok := stage.Diagram_StatesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagram.Name
			}
		}
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

func (inst *StateMachine) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootStateMachines":
			if _library, ok := stage.Library_RootStateMachines_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "StateMachinesWhoseNodeIsExpanded":
			if _library, ok := stage.Library_StateMachinesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *StateShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Transition) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Transition_Shape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
