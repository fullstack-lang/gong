// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	*slice = cleanedSlice
	return len(cleanedSlice) != len(*slice)
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	if !IsStagedPointerToGongstruct(stage, *element) {
		var zero T
		*element = zero
		return true
	}
	return false
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Action
func (action *Action) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Activities
func (activities *Activities) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Architecture
func (architecture *Architecture) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &architecture.StateMachines)  || modified
	modified = GongCleanSlice(stage, &architecture.Roles)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagram.State_Shapes)  || modified
	modified = GongCleanSlice(stage, &diagram.Transition_Shapes)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Guard
func (guard *Guard) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Kill
func (kill *Kill) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Message
func (message *Message) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &message.MessageType)  || modified
	modified = GongCleanPointer(stage, &message.OriginTransition)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MessageType
func (messagetype *MessageType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Object
func (object *Object) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &object.Messages)  || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &object.State)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Role
func (role *Role) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &role.RolesWithSamePermissions)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by State
func (state *State) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &state.SubStates)  || modified
	modified = GongCleanSlice(stage, &state.Diagrams)  || modified
	modified = GongCleanSlice(stage, &state.Activities)  || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &state.Parent)  || modified
	modified = GongCleanPointer(stage, &state.Entry)  || modified
	modified = GongCleanPointer(stage, &state.Exit)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StateMachine
func (statemachine *StateMachine) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &statemachine.States)  || modified
	modified = GongCleanSlice(stage, &statemachine.Diagrams)  || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &statemachine.InitialState)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StateShape
func (stateshape *StateShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &stateshape.State)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Transition
func (transition *Transition) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &transition.RolesWithPermissions)  || modified
	modified = GongCleanSlice(stage, &transition.GeneratedMessages)  || modified
	modified = GongCleanSlice(stage, &transition.Diagrams)  || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &transition.Start)  || modified
	modified = GongCleanPointer(stage, &transition.End)  || modified
	modified = GongCleanPointer(stage, &transition.Guard)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Transition_Shape
func (transition_shape *Transition_Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &transition_shape.Transition)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
