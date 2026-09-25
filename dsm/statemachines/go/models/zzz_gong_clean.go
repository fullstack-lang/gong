// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&diagram.State_Shapes) || modified
	modified = stage.CleanSlice(&diagram.StatesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.Transition_Shapes) || modified
	modified = stage.CleanSlice(&diagram.Note_Shapes) || modified
	modified = stage.CleanSlice(&diagram.NoteState_Shapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	modified = stage.CleanSlice(&library.Diagrams) || modified
	modified = stage.CleanSlice(&library.RootStateMachines) || modified
	modified = stage.CleanSlice(&library.StateMachinesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.Roles) || modified
	modified = stage.CleanSlice(&library.MessageTypes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Message
func (message *Message) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&message.MessageType) || modified
	modified = stage.CleanPointer(&message.OriginTransition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&note.State) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteStateShape
func (notestateshape *NoteStateShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&notestateshape.Note) || modified
	modified = stage.CleanPointer(&notestateshape.State) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Object
func (object *Object) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&object.Messages) || modified
	// insertion point per field
	modified = stage.CleanPointer(&object.State) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Role
func (role *Role) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&role.RolesWithSamePermissions) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by State
func (state *State) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&state.SubStates) || modified
	modified = stage.CleanSlice(&state.Activities) || modified
	modified = stage.CleanSlice(&state.Diagrams) || modified
	modified = stage.CleanSlice(&state.Notes) || modified
	// insertion point per field
	modified = stage.CleanPointer(&state.Entry) || modified
	modified = stage.CleanPointer(&state.Exit) || modified
	modified = stage.CleanPointer(&state.Parent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StateMachine
func (statemachine *StateMachine) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&statemachine.States) || modified
	modified = stage.CleanSlice(&statemachine.Diagrams) || modified
	// insertion point per field
	modified = stage.CleanPointer(&statemachine.InitialState) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by StateShape
func (stateshape *StateShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&stateshape.State) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Transition
func (transition *Transition) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&transition.RolesWithPermissions) || modified
	modified = stage.CleanSlice(&transition.GeneratedMessages) || modified
	modified = stage.CleanSlice(&transition.Diagrams) || modified
	// insertion point per field
	modified = stage.CleanPointer(&transition.Start) || modified
	modified = stage.CleanPointer(&transition.End) || modified
	modified = stage.CleanPointer(&transition.Guard) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Transition_Shape
func (transition_shape *Transition_Shape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&transition_shape.Transition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
