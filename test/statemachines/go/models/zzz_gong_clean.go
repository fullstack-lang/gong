// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}

	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Architecture
func (architecture *Architecture) GongClean(stage *Stage) {
	// insertion point per field
	architecture.StateMachines = GongCleanSlice(stage, architecture.StateMachines)
	architecture.Roles = GongCleanSlice(stage, architecture.Roles)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) {
	// insertion point per field
	diagram.State_Shapes = GongCleanSlice(stage, diagram.State_Shapes)
	diagram.Transition_Shapes = GongCleanSlice(stage, diagram.Transition_Shapes)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by DoAction
func (doaction *DoAction) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Kill
func (kill *Kill) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Message
func (message *Message) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	message.MessageType = GongCleanPointer(stage, message.MessageType)
	message.OriginTransition = GongCleanPointer(stage, message.OriginTransition)
}

// Clean garbage collect unstaged instances that are referenced by MessageType
func (messagetype *MessageType) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Object
func (object *Object) GongClean(stage *Stage) {
	// insertion point per field
	object.Messages = GongCleanSlice(stage, object.Messages)
	// insertion point per field
	object.State = GongCleanPointer(stage, object.State)
}

// Clean garbage collect unstaged instances that are referenced by Role
func (role *Role) GongClean(stage *Stage) {
	// insertion point per field
	role.RolesWithSamePermissions = GongCleanSlice(stage, role.RolesWithSamePermissions)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by State
func (state *State) GongClean(stage *Stage) {
	// insertion point per field
	state.SubStates = GongCleanSlice(stage, state.SubStates)
	state.Diagrams = GongCleanSlice(stage, state.Diagrams)
	state.DoActions = GongCleanSlice(stage, state.DoActions)
	// insertion point per field
	state.Parent = GongCleanPointer(stage, state.Parent)
}

// Clean garbage collect unstaged instances that are referenced by StateMachine
func (statemachine *StateMachine) GongClean(stage *Stage) {
	// insertion point per field
	statemachine.States = GongCleanSlice(stage, statemachine.States)
	statemachine.Diagrams = GongCleanSlice(stage, statemachine.Diagrams)
	// insertion point per field
	statemachine.InitialState = GongCleanPointer(stage, statemachine.InitialState)
}

// Clean garbage collect unstaged instances that are referenced by StateShape
func (stateshape *StateShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	stateshape.State = GongCleanPointer(stage, stateshape.State)
}

// Clean garbage collect unstaged instances that are referenced by Transition
func (transition *Transition) GongClean(stage *Stage) {
	// insertion point per field
	transition.RolesWithPermissions = GongCleanSlice(stage, transition.RolesWithPermissions)
	transition.GeneratedMessages = GongCleanSlice(stage, transition.GeneratedMessages)
	transition.Diagrams = GongCleanSlice(stage, transition.Diagrams)
	// insertion point per field
	transition.Start = GongCleanPointer(stage, transition.Start)
	transition.End = GongCleanPointer(stage, transition.End)
}

// Clean garbage collect unstaged instances that are referenced by Transition_Shape
func (transition_shape *Transition_Shape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	transition_shape.Transition = GongCleanPointer(stage, transition_shape.Transition)
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
