// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Action:
		ok = stage.IsStagedAction(target)

	case *Activities:
		ok = stage.IsStagedActivities(target)

	case *Architecture:
		ok = stage.IsStagedArchitecture(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Guard:
		ok = stage.IsStagedGuard(target)

	case *Kill:
		ok = stage.IsStagedKill(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Message:
		ok = stage.IsStagedMessage(target)

	case *MessageType:
		ok = stage.IsStagedMessageType(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteStateShape:
		ok = stage.IsStagedNoteStateShape(target)

	case *NoteTransitionShape:
		ok = stage.IsStagedNoteTransitionShape(target)

	case *Object:
		ok = stage.IsStagedObject(target)

	case *Role:
		ok = stage.IsStagedRole(target)

	case *State:
		ok = stage.IsStagedState(target)

	case *StateMachine:
		ok = stage.IsStagedStateMachine(target)

	case *StateShape:
		ok = stage.IsStagedStateShape(target)

	case *Transition:
		ok = stage.IsStagedTransition(target)

	case *Transition_Shape:
		ok = stage.IsStagedTransition_Shape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Action:
		ok = stage.IsStagedAction(target)

	case *Activities:
		ok = stage.IsStagedActivities(target)

	case *Architecture:
		ok = stage.IsStagedArchitecture(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Guard:
		ok = stage.IsStagedGuard(target)

	case *Kill:
		ok = stage.IsStagedKill(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Message:
		ok = stage.IsStagedMessage(target)

	case *MessageType:
		ok = stage.IsStagedMessageType(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteStateShape:
		ok = stage.IsStagedNoteStateShape(target)

	case *NoteTransitionShape:
		ok = stage.IsStagedNoteTransitionShape(target)

	case *Object:
		ok = stage.IsStagedObject(target)

	case *Role:
		ok = stage.IsStagedRole(target)

	case *State:
		ok = stage.IsStagedState(target)

	case *StateMachine:
		ok = stage.IsStagedStateMachine(target)

	case *StateShape:
		ok = stage.IsStagedStateShape(target)

	case *Transition:
		ok = stage.IsStagedTransition(target)

	case *Transition_Shape:
		ok = stage.IsStagedTransition_Shape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAction(action *Action) (ok bool) {

	_, ok = stage.Actions[action]

	return
}

func (stage *Stage) IsStagedActivities(activities *Activities) (ok bool) {

	_, ok = stage.Activitiess[activities]

	return
}

func (stage *Stage) IsStagedArchitecture(architecture *Architecture) (ok bool) {

	_, ok = stage.Architectures[architecture]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedGuard(guard *Guard) (ok bool) {

	_, ok = stage.Guards[guard]

	return
}

func (stage *Stage) IsStagedKill(kill *Kill) (ok bool) {

	_, ok = stage.Kills[kill]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedMessage(message *Message) (ok bool) {

	_, ok = stage.Messages[message]

	return
}

func (stage *Stage) IsStagedMessageType(messagetype *MessageType) (ok bool) {

	_, ok = stage.MessageTypes[messagetype]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteStateShape(notestateshape *NoteStateShape) (ok bool) {

	_, ok = stage.NoteStateShapes[notestateshape]

	return
}

func (stage *Stage) IsStagedNoteTransitionShape(notetransitionshape *NoteTransitionShape) (ok bool) {

	_, ok = stage.NoteTransitionShapes[notetransitionshape]

	return
}

func (stage *Stage) IsStagedObject(object *Object) (ok bool) {

	_, ok = stage.Objects[object]

	return
}

func (stage *Stage) IsStagedRole(role *Role) (ok bool) {

	_, ok = stage.Roles[role]

	return
}

func (stage *Stage) IsStagedState(state *State) (ok bool) {

	_, ok = stage.States[state]

	return
}

func (stage *Stage) IsStagedStateMachine(statemachine *StateMachine) (ok bool) {

	_, ok = stage.StateMachines[statemachine]

	return
}

func (stage *Stage) IsStagedStateShape(stateshape *StateShape) (ok bool) {

	_, ok = stage.StateShapes[stateshape]

	return
}

func (stage *Stage) IsStagedTransition(transition *Transition) (ok bool) {

	_, ok = stage.Transitions[transition]

	return
}

func (stage *Stage) IsStagedTransition_Shape(transition_shape *Transition_Shape) (ok bool) {

	_, ok = stage.Transition_Shapes[transition_shape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Action:
		stage.StageBranchAction(target)

	case *Activities:
		stage.StageBranchActivities(target)

	case *Architecture:
		stage.StageBranchArchitecture(target)

	case *Diagram:
		stage.StageBranchDiagram(target)

	case *Guard:
		stage.StageBranchGuard(target)

	case *Kill:
		stage.StageBranchKill(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Message:
		stage.StageBranchMessage(target)

	case *MessageType:
		stage.StageBranchMessageType(target)

	case *Note:
		stage.StageBranchNote(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *NoteStateShape:
		stage.StageBranchNoteStateShape(target)

	case *NoteTransitionShape:
		stage.StageBranchNoteTransitionShape(target)

	case *Object:
		stage.StageBranchObject(target)

	case *Role:
		stage.StageBranchRole(target)

	case *State:
		stage.StageBranchState(target)

	case *StateMachine:
		stage.StageBranchStateMachine(target)

	case *StateShape:
		stage.StageBranchStateShape(target)

	case *Transition:
		stage.StageBranchTransition(target)

	case *Transition_Shape:
		stage.StageBranchTransition_Shape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAction(action *Action) {

	// check if instance is already staged
	if IsStaged(stage, action) {
		return
	}

	action.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchActivities(activities *Activities) {

	// check if instance is already staged
	if IsStaged(stage, activities) {
		return
	}

	activities.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchArchitecture(architecture *Architecture) {

	// check if instance is already staged
	if IsStaged(stage, architecture) {
		return
	}

	architecture.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _statemachine := range architecture.StateMachines {
		StageBranch(stage, _statemachine)
	}
	for _, _role := range architecture.Roles {
		StageBranch(stage, _role)
	}

}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if IsStaged(stage, diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}
	for _, _stateshape := range diagram.State_Shapes {
		StageBranch(stage, _stateshape)
	}
	for _, _state := range diagram.StatesWhoseNodeIsExpanded {
		StageBranch(stage, _state)
	}
	for _, _transition_shape := range diagram.Transition_Shapes {
		StageBranch(stage, _transition_shape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		StageBranch(stage, _noteshape)
	}
	for _, _notestateshape := range diagram.NoteState_Shapes {
		StageBranch(stage, _notestateshape)
	}
	for _, _notetransitionshape := range diagram.NoteTransition_Shapes {
		StageBranch(stage, _notetransitionshape)
	}

}

func (stage *Stage) StageBranchGuard(guard *Guard) {

	// check if instance is already staged
	if IsStaged(stage, guard) {
		return
	}

	guard.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchKill(kill *Kill) {

	// check if instance is already staged
	if IsStaged(stage, kill) {
		return
	}

	kill.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if IsStaged(stage, library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		StageBranch(stage, _library)
	}
	for _, _diagram := range library.Diagrams {
		StageBranch(stage, _diagram)
	}
	for _, _statemachine := range library.RootStateMachines {
		StageBranch(stage, _statemachine)
	}
	for _, _statemachine := range library.StateMachinesWhoseNodeIsExpanded {
		StageBranch(stage, _statemachine)
	}
	for _, _note := range library.RootNotes {
		StageBranch(stage, _note)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		StageBranch(stage, _library)
	}

}

func (stage *Stage) StageBranchMessage(message *Message) {

	// check if instance is already staged
	if IsStaged(stage, message) {
		return
	}

	message.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if message.MessageType != nil {
		StageBranch(stage, message.MessageType)
	}
	if message.OriginTransition != nil {
		StageBranch(stage, message.OriginTransition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMessageType(messagetype *MessageType) {

	// check if instance is already staged
	if IsStaged(stage, messagetype) {
		return
	}

	messagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if IsStaged(stage, note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range note.States {
		StageBranch(stage, _state)
	}
	for _, _transition := range note.Transitions {
		StageBranch(stage, _transition)
	}

}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if IsStaged(stage, noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		StageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteStateShape(notestateshape *NoteStateShape) {

	// check if instance is already staged
	if IsStaged(stage, notestateshape) {
		return
	}

	notestateshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notestateshape.Note != nil {
		StageBranch(stage, notestateshape.Note)
	}
	if notestateshape.State != nil {
		StageBranch(stage, notestateshape.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteTransitionShape(notetransitionshape *NoteTransitionShape) {

	// check if instance is already staged
	if IsStaged(stage, notetransitionshape) {
		return
	}

	notetransitionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetransitionshape.Note != nil {
		StageBranch(stage, notetransitionshape.Note)
	}
	if notetransitionshape.Transition != nil {
		StageBranch(stage, notetransitionshape.Transition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchObject(object *Object) {

	// check if instance is already staged
	if IsStaged(stage, object) {
		return
	}

	object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if object.State != nil {
		StageBranch(stage, object.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _message := range object.Messages {
		StageBranch(stage, _message)
	}

}

func (stage *Stage) StageBranchRole(role *Role) {

	// check if instance is already staged
	if IsStaged(stage, role) {
		return
	}

	role.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range role.RolesWithSamePermissions {
		StageBranch(stage, _role)
	}

}

func (stage *Stage) StageBranchState(state *State) {

	// check if instance is already staged
	if IsStaged(stage, state) {
		return
	}

	state.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if state.Parent != nil {
		StageBranch(stage, state.Parent)
	}
	if state.Entry != nil {
		StageBranch(stage, state.Entry)
	}
	if state.Exit != nil {
		StageBranch(stage, state.Exit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range state.SubStates {
		StageBranch(stage, _state)
	}
	for _, _diagram := range state.Diagrams {
		StageBranch(stage, _diagram)
	}
	for _, _activities := range state.Activities {
		StageBranch(stage, _activities)
	}

}

func (stage *Stage) StageBranchStateMachine(statemachine *StateMachine) {

	// check if instance is already staged
	if IsStaged(stage, statemachine) {
		return
	}

	statemachine.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if statemachine.InitialState != nil {
		StageBranch(stage, statemachine.InitialState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range statemachine.States {
		StageBranch(stage, _state)
	}
	for _, _diagram := range statemachine.Diagrams {
		StageBranch(stage, _diagram)
	}

}

func (stage *Stage) StageBranchStateShape(stateshape *StateShape) {

	// check if instance is already staged
	if IsStaged(stage, stateshape) {
		return
	}

	stateshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stateshape.State != nil {
		StageBranch(stage, stateshape.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTransition(transition *Transition) {

	// check if instance is already staged
	if IsStaged(stage, transition) {
		return
	}

	transition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transition.Start != nil {
		StageBranch(stage, transition.Start)
	}
	if transition.End != nil {
		StageBranch(stage, transition.End)
	}
	if transition.Guard != nil {
		StageBranch(stage, transition.Guard)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range transition.RolesWithPermissions {
		StageBranch(stage, _role)
	}
	for _, _messagetype := range transition.GeneratedMessages {
		StageBranch(stage, _messagetype)
	}
	for _, _diagram := range transition.Diagrams {
		StageBranch(stage, _diagram)
	}

}

func (stage *Stage) StageBranchTransition_Shape(transition_shape *Transition_Shape) {

	// check if instance is already staged
	if IsStaged(stage, transition_shape) {
		return
	}

	transition_shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transition_shape.Transition != nil {
		StageBranch(stage, transition_shape.Transition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Action:
		toT := CopyBranchAction(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Activities:
		toT := CopyBranchActivities(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Architecture:
		toT := CopyBranchArchitecture(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := CopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Guard:
		toT := CopyBranchGuard(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Kill:
		toT := CopyBranchKill(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Message:
		toT := CopyBranchMessage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MessageType:
		toT := CopyBranchMessageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := CopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := CopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteStateShape:
		toT := CopyBranchNoteStateShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteTransitionShape:
		toT := CopyBranchNoteTransitionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Object:
		toT := CopyBranchObject(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Role:
		toT := CopyBranchRole(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *State:
		toT := CopyBranchState(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StateMachine:
		toT := CopyBranchStateMachine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StateShape:
		toT := CopyBranchStateShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Transition:
		toT := CopyBranchTransition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Transition_Shape:
		toT := CopyBranchTransition_Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAction(mapOrigCopy map[any]any, actionFrom *Action) (actionTo *Action) {

	// actionFrom has already been copied
	if _actionTo, ok := mapOrigCopy[actionFrom]; ok {
		actionTo = _actionTo.(*Action)
		return
	}

	actionTo = new(Action)
	mapOrigCopy[actionFrom] = actionTo
	actionFrom.CopyBasicFields(actionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchActivities(mapOrigCopy map[any]any, activitiesFrom *Activities) (activitiesTo *Activities) {

	// activitiesFrom has already been copied
	if _activitiesTo, ok := mapOrigCopy[activitiesFrom]; ok {
		activitiesTo = _activitiesTo.(*Activities)
		return
	}

	activitiesTo = new(Activities)
	mapOrigCopy[activitiesFrom] = activitiesTo
	activitiesFrom.CopyBasicFields(activitiesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchArchitecture(mapOrigCopy map[any]any, architectureFrom *Architecture) (architectureTo *Architecture) {

	// architectureFrom has already been copied
	if _architectureTo, ok := mapOrigCopy[architectureFrom]; ok {
		architectureTo = _architectureTo.(*Architecture)
		return
	}

	architectureTo = new(Architecture)
	mapOrigCopy[architectureFrom] = architectureTo
	architectureFrom.CopyBasicFields(architectureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _statemachine := range architectureFrom.StateMachines {
		architectureTo.StateMachines = append(architectureTo.StateMachines, CopyBranchStateMachine(mapOrigCopy, _statemachine))
	}
	for _, _role := range architectureFrom.Roles {
		architectureTo.Roles = append(architectureTo.Roles, CopyBranchRole(mapOrigCopy, _role))
	}

	return
}

func CopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.CopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range diagramFrom.NotesWhoseNodeIsExpanded {
		diagramTo.NotesWhoseNodeIsExpanded = append(diagramTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _stateshape := range diagramFrom.State_Shapes {
		diagramTo.State_Shapes = append(diagramTo.State_Shapes, CopyBranchStateShape(mapOrigCopy, _stateshape))
	}
	for _, _state := range diagramFrom.StatesWhoseNodeIsExpanded {
		diagramTo.StatesWhoseNodeIsExpanded = append(diagramTo.StatesWhoseNodeIsExpanded, CopyBranchState(mapOrigCopy, _state))
	}
	for _, _transition_shape := range diagramFrom.Transition_Shapes {
		diagramTo.Transition_Shapes = append(diagramTo.Transition_Shapes, CopyBranchTransition_Shape(mapOrigCopy, _transition_shape))
	}
	for _, _noteshape := range diagramFrom.Note_Shapes {
		diagramTo.Note_Shapes = append(diagramTo.Note_Shapes, CopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _notestateshape := range diagramFrom.NoteState_Shapes {
		diagramTo.NoteState_Shapes = append(diagramTo.NoteState_Shapes, CopyBranchNoteStateShape(mapOrigCopy, _notestateshape))
	}
	for _, _notetransitionshape := range diagramFrom.NoteTransition_Shapes {
		diagramTo.NoteTransition_Shapes = append(diagramTo.NoteTransition_Shapes, CopyBranchNoteTransitionShape(mapOrigCopy, _notetransitionshape))
	}

	return
}

func CopyBranchGuard(mapOrigCopy map[any]any, guardFrom *Guard) (guardTo *Guard) {

	// guardFrom has already been copied
	if _guardTo, ok := mapOrigCopy[guardFrom]; ok {
		guardTo = _guardTo.(*Guard)
		return
	}

	guardTo = new(Guard)
	mapOrigCopy[guardFrom] = guardTo
	guardFrom.CopyBasicFields(guardTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchKill(mapOrigCopy map[any]any, killFrom *Kill) (killTo *Kill) {

	// killFrom has already been copied
	if _killTo, ok := mapOrigCopy[killFrom]; ok {
		killTo = _killTo.(*Kill)
		return
	}

	killTo = new(Kill)
	mapOrigCopy[killFrom] = killTo
	killFrom.CopyBasicFields(killTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.CopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _diagram := range libraryFrom.Diagrams {
		libraryTo.Diagrams = append(libraryTo.Diagrams, CopyBranchDiagram(mapOrigCopy, _diagram))
	}
	for _, _statemachine := range libraryFrom.RootStateMachines {
		libraryTo.RootStateMachines = append(libraryTo.RootStateMachines, CopyBranchStateMachine(mapOrigCopy, _statemachine))
	}
	for _, _statemachine := range libraryFrom.StateMachinesWhoseNodeIsExpanded {
		libraryTo.StateMachinesWhoseNodeIsExpanded = append(libraryTo.StateMachinesWhoseNodeIsExpanded, CopyBranchStateMachine(mapOrigCopy, _statemachine))
	}
	for _, _note := range libraryFrom.RootNotes {
		libraryTo.RootNotes = append(libraryTo.RootNotes, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _note := range libraryFrom.NotesWhoseNodeIsExpanded {
		libraryTo.NotesWhoseNodeIsExpanded = append(libraryTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, CopyBranchLibrary(mapOrigCopy, _library))
	}

	return
}

func CopyBranchMessage(mapOrigCopy map[any]any, messageFrom *Message) (messageTo *Message) {

	// messageFrom has already been copied
	if _messageTo, ok := mapOrigCopy[messageFrom]; ok {
		messageTo = _messageTo.(*Message)
		return
	}

	messageTo = new(Message)
	mapOrigCopy[messageFrom] = messageTo
	messageFrom.CopyBasicFields(messageTo)

	//insertion point for the staging of instances referenced by pointers
	if messageFrom.MessageType != nil {
		messageTo.MessageType = CopyBranchMessageType(mapOrigCopy, messageFrom.MessageType)
	}
	if messageFrom.OriginTransition != nil {
		messageTo.OriginTransition = CopyBranchTransition(mapOrigCopy, messageFrom.OriginTransition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMessageType(mapOrigCopy map[any]any, messagetypeFrom *MessageType) (messagetypeTo *MessageType) {

	// messagetypeFrom has already been copied
	if _messagetypeTo, ok := mapOrigCopy[messagetypeFrom]; ok {
		messagetypeTo = _messagetypeTo.(*MessageType)
		return
	}

	messagetypeTo = new(MessageType)
	mapOrigCopy[messagetypeFrom] = messagetypeTo
	messagetypeFrom.CopyBasicFields(messagetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.CopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range noteFrom.States {
		noteTo.States = append(noteTo.States, CopyBranchState(mapOrigCopy, _state))
	}
	for _, _transition := range noteFrom.Transitions {
		noteTo.Transitions = append(noteTo.Transitions, CopyBranchTransition(mapOrigCopy, _transition))
	}

	return
}

func CopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.CopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = CopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteStateShape(mapOrigCopy map[any]any, notestateshapeFrom *NoteStateShape) (notestateshapeTo *NoteStateShape) {

	// notestateshapeFrom has already been copied
	if _notestateshapeTo, ok := mapOrigCopy[notestateshapeFrom]; ok {
		notestateshapeTo = _notestateshapeTo.(*NoteStateShape)
		return
	}

	notestateshapeTo = new(NoteStateShape)
	mapOrigCopy[notestateshapeFrom] = notestateshapeTo
	notestateshapeFrom.CopyBasicFields(notestateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notestateshapeFrom.Note != nil {
		notestateshapeTo.Note = CopyBranchNote(mapOrigCopy, notestateshapeFrom.Note)
	}
	if notestateshapeFrom.State != nil {
		notestateshapeTo.State = CopyBranchState(mapOrigCopy, notestateshapeFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteTransitionShape(mapOrigCopy map[any]any, notetransitionshapeFrom *NoteTransitionShape) (notetransitionshapeTo *NoteTransitionShape) {

	// notetransitionshapeFrom has already been copied
	if _notetransitionshapeTo, ok := mapOrigCopy[notetransitionshapeFrom]; ok {
		notetransitionshapeTo = _notetransitionshapeTo.(*NoteTransitionShape)
		return
	}

	notetransitionshapeTo = new(NoteTransitionShape)
	mapOrigCopy[notetransitionshapeFrom] = notetransitionshapeTo
	notetransitionshapeFrom.CopyBasicFields(notetransitionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notetransitionshapeFrom.Note != nil {
		notetransitionshapeTo.Note = CopyBranchNote(mapOrigCopy, notetransitionshapeFrom.Note)
	}
	if notetransitionshapeFrom.Transition != nil {
		notetransitionshapeTo.Transition = CopyBranchTransition(mapOrigCopy, notetransitionshapeFrom.Transition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchObject(mapOrigCopy map[any]any, objectFrom *Object) (objectTo *Object) {

	// objectFrom has already been copied
	if _objectTo, ok := mapOrigCopy[objectFrom]; ok {
		objectTo = _objectTo.(*Object)
		return
	}

	objectTo = new(Object)
	mapOrigCopy[objectFrom] = objectTo
	objectFrom.CopyBasicFields(objectTo)

	//insertion point for the staging of instances referenced by pointers
	if objectFrom.State != nil {
		objectTo.State = CopyBranchState(mapOrigCopy, objectFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _message := range objectFrom.Messages {
		objectTo.Messages = append(objectTo.Messages, CopyBranchMessage(mapOrigCopy, _message))
	}

	return
}

func CopyBranchRole(mapOrigCopy map[any]any, roleFrom *Role) (roleTo *Role) {

	// roleFrom has already been copied
	if _roleTo, ok := mapOrigCopy[roleFrom]; ok {
		roleTo = _roleTo.(*Role)
		return
	}

	roleTo = new(Role)
	mapOrigCopy[roleFrom] = roleTo
	roleFrom.CopyBasicFields(roleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range roleFrom.RolesWithSamePermissions {
		roleTo.RolesWithSamePermissions = append(roleTo.RolesWithSamePermissions, CopyBranchRole(mapOrigCopy, _role))
	}

	return
}

func CopyBranchState(mapOrigCopy map[any]any, stateFrom *State) (stateTo *State) {

	// stateFrom has already been copied
	if _stateTo, ok := mapOrigCopy[stateFrom]; ok {
		stateTo = _stateTo.(*State)
		return
	}

	stateTo = new(State)
	mapOrigCopy[stateFrom] = stateTo
	stateFrom.CopyBasicFields(stateTo)

	//insertion point for the staging of instances referenced by pointers
	if stateFrom.Parent != nil {
		stateTo.Parent = CopyBranchState(mapOrigCopy, stateFrom.Parent)
	}
	if stateFrom.Entry != nil {
		stateTo.Entry = CopyBranchAction(mapOrigCopy, stateFrom.Entry)
	}
	if stateFrom.Exit != nil {
		stateTo.Exit = CopyBranchAction(mapOrigCopy, stateFrom.Exit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range stateFrom.SubStates {
		stateTo.SubStates = append(stateTo.SubStates, CopyBranchState(mapOrigCopy, _state))
	}
	for _, _diagram := range stateFrom.Diagrams {
		stateTo.Diagrams = append(stateTo.Diagrams, CopyBranchDiagram(mapOrigCopy, _diagram))
	}
	for _, _activities := range stateFrom.Activities {
		stateTo.Activities = append(stateTo.Activities, CopyBranchActivities(mapOrigCopy, _activities))
	}

	return
}

func CopyBranchStateMachine(mapOrigCopy map[any]any, statemachineFrom *StateMachine) (statemachineTo *StateMachine) {

	// statemachineFrom has already been copied
	if _statemachineTo, ok := mapOrigCopy[statemachineFrom]; ok {
		statemachineTo = _statemachineTo.(*StateMachine)
		return
	}

	statemachineTo = new(StateMachine)
	mapOrigCopy[statemachineFrom] = statemachineTo
	statemachineFrom.CopyBasicFields(statemachineTo)

	//insertion point for the staging of instances referenced by pointers
	if statemachineFrom.InitialState != nil {
		statemachineTo.InitialState = CopyBranchState(mapOrigCopy, statemachineFrom.InitialState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range statemachineFrom.States {
		statemachineTo.States = append(statemachineTo.States, CopyBranchState(mapOrigCopy, _state))
	}
	for _, _diagram := range statemachineFrom.Diagrams {
		statemachineTo.Diagrams = append(statemachineTo.Diagrams, CopyBranchDiagram(mapOrigCopy, _diagram))
	}

	return
}

func CopyBranchStateShape(mapOrigCopy map[any]any, stateshapeFrom *StateShape) (stateshapeTo *StateShape) {

	// stateshapeFrom has already been copied
	if _stateshapeTo, ok := mapOrigCopy[stateshapeFrom]; ok {
		stateshapeTo = _stateshapeTo.(*StateShape)
		return
	}

	stateshapeTo = new(StateShape)
	mapOrigCopy[stateshapeFrom] = stateshapeTo
	stateshapeFrom.CopyBasicFields(stateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stateshapeFrom.State != nil {
		stateshapeTo.State = CopyBranchState(mapOrigCopy, stateshapeFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTransition(mapOrigCopy map[any]any, transitionFrom *Transition) (transitionTo *Transition) {

	// transitionFrom has already been copied
	if _transitionTo, ok := mapOrigCopy[transitionFrom]; ok {
		transitionTo = _transitionTo.(*Transition)
		return
	}

	transitionTo = new(Transition)
	mapOrigCopy[transitionFrom] = transitionTo
	transitionFrom.CopyBasicFields(transitionTo)

	//insertion point for the staging of instances referenced by pointers
	if transitionFrom.Start != nil {
		transitionTo.Start = CopyBranchState(mapOrigCopy, transitionFrom.Start)
	}
	if transitionFrom.End != nil {
		transitionTo.End = CopyBranchState(mapOrigCopy, transitionFrom.End)
	}
	if transitionFrom.Guard != nil {
		transitionTo.Guard = CopyBranchGuard(mapOrigCopy, transitionFrom.Guard)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range transitionFrom.RolesWithPermissions {
		transitionTo.RolesWithPermissions = append(transitionTo.RolesWithPermissions, CopyBranchRole(mapOrigCopy, _role))
	}
	for _, _messagetype := range transitionFrom.GeneratedMessages {
		transitionTo.GeneratedMessages = append(transitionTo.GeneratedMessages, CopyBranchMessageType(mapOrigCopy, _messagetype))
	}
	for _, _diagram := range transitionFrom.Diagrams {
		transitionTo.Diagrams = append(transitionTo.Diagrams, CopyBranchDiagram(mapOrigCopy, _diagram))
	}

	return
}

func CopyBranchTransition_Shape(mapOrigCopy map[any]any, transition_shapeFrom *Transition_Shape) (transition_shapeTo *Transition_Shape) {

	// transition_shapeFrom has already been copied
	if _transition_shapeTo, ok := mapOrigCopy[transition_shapeFrom]; ok {
		transition_shapeTo = _transition_shapeTo.(*Transition_Shape)
		return
	}

	transition_shapeTo = new(Transition_Shape)
	mapOrigCopy[transition_shapeFrom] = transition_shapeTo
	transition_shapeFrom.CopyBasicFields(transition_shapeTo)

	//insertion point for the staging of instances referenced by pointers
	if transition_shapeFrom.Transition != nil {
		transition_shapeTo.Transition = CopyBranchTransition(mapOrigCopy, transition_shapeFrom.Transition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Action:
		stage.UnstageBranchAction(target)

	case *Activities:
		stage.UnstageBranchActivities(target)

	case *Architecture:
		stage.UnstageBranchArchitecture(target)

	case *Diagram:
		stage.UnstageBranchDiagram(target)

	case *Guard:
		stage.UnstageBranchGuard(target)

	case *Kill:
		stage.UnstageBranchKill(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Message:
		stage.UnstageBranchMessage(target)

	case *MessageType:
		stage.UnstageBranchMessageType(target)

	case *Note:
		stage.UnstageBranchNote(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *NoteStateShape:
		stage.UnstageBranchNoteStateShape(target)

	case *NoteTransitionShape:
		stage.UnstageBranchNoteTransitionShape(target)

	case *Object:
		stage.UnstageBranchObject(target)

	case *Role:
		stage.UnstageBranchRole(target)

	case *State:
		stage.UnstageBranchState(target)

	case *StateMachine:
		stage.UnstageBranchStateMachine(target)

	case *StateShape:
		stage.UnstageBranchStateShape(target)

	case *Transition:
		stage.UnstageBranchTransition(target)

	case *Transition_Shape:
		stage.UnstageBranchTransition_Shape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAction(action *Action) {

	// check if instance is already staged
	if !IsStaged(stage, action) {
		return
	}

	action.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchActivities(activities *Activities) {

	// check if instance is already staged
	if !IsStaged(stage, activities) {
		return
	}

	activities.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchArchitecture(architecture *Architecture) {

	// check if instance is already staged
	if !IsStaged(stage, architecture) {
		return
	}

	architecture.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _statemachine := range architecture.StateMachines {
		UnstageBranch(stage, _statemachine)
	}
	for _, _role := range architecture.Roles {
		UnstageBranch(stage, _role)
	}

}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !IsStaged(stage, diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}
	for _, _stateshape := range diagram.State_Shapes {
		UnstageBranch(stage, _stateshape)
	}
	for _, _state := range diagram.StatesWhoseNodeIsExpanded {
		UnstageBranch(stage, _state)
	}
	for _, _transition_shape := range diagram.Transition_Shapes {
		UnstageBranch(stage, _transition_shape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		UnstageBranch(stage, _noteshape)
	}
	for _, _notestateshape := range diagram.NoteState_Shapes {
		UnstageBranch(stage, _notestateshape)
	}
	for _, _notetransitionshape := range diagram.NoteTransition_Shapes {
		UnstageBranch(stage, _notetransitionshape)
	}

}

func (stage *Stage) UnstageBranchGuard(guard *Guard) {

	// check if instance is already staged
	if !IsStaged(stage, guard) {
		return
	}

	guard.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchKill(kill *Kill) {

	// check if instance is already staged
	if !IsStaged(stage, kill) {
		return
	}

	kill.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !IsStaged(stage, library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		UnstageBranch(stage, _library)
	}
	for _, _diagram := range library.Diagrams {
		UnstageBranch(stage, _diagram)
	}
	for _, _statemachine := range library.RootStateMachines {
		UnstageBranch(stage, _statemachine)
	}
	for _, _statemachine := range library.StateMachinesWhoseNodeIsExpanded {
		UnstageBranch(stage, _statemachine)
	}
	for _, _note := range library.RootNotes {
		UnstageBranch(stage, _note)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		UnstageBranch(stage, _library)
	}

}

func (stage *Stage) UnstageBranchMessage(message *Message) {

	// check if instance is already staged
	if !IsStaged(stage, message) {
		return
	}

	message.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if message.MessageType != nil {
		UnstageBranch(stage, message.MessageType)
	}
	if message.OriginTransition != nil {
		UnstageBranch(stage, message.OriginTransition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMessageType(messagetype *MessageType) {

	// check if instance is already staged
	if !IsStaged(stage, messagetype) {
		return
	}

	messagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !IsStaged(stage, note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range note.States {
		UnstageBranch(stage, _state)
	}
	for _, _transition := range note.Transitions {
		UnstageBranch(stage, _transition)
	}

}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		UnstageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteStateShape(notestateshape *NoteStateShape) {

	// check if instance is already staged
	if !IsStaged(stage, notestateshape) {
		return
	}

	notestateshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notestateshape.Note != nil {
		UnstageBranch(stage, notestateshape.Note)
	}
	if notestateshape.State != nil {
		UnstageBranch(stage, notestateshape.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteTransitionShape(notetransitionshape *NoteTransitionShape) {

	// check if instance is already staged
	if !IsStaged(stage, notetransitionshape) {
		return
	}

	notetransitionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetransitionshape.Note != nil {
		UnstageBranch(stage, notetransitionshape.Note)
	}
	if notetransitionshape.Transition != nil {
		UnstageBranch(stage, notetransitionshape.Transition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchObject(object *Object) {

	// check if instance is already staged
	if !IsStaged(stage, object) {
		return
	}

	object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if object.State != nil {
		UnstageBranch(stage, object.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _message := range object.Messages {
		UnstageBranch(stage, _message)
	}

}

func (stage *Stage) UnstageBranchRole(role *Role) {

	// check if instance is already staged
	if !IsStaged(stage, role) {
		return
	}

	role.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range role.RolesWithSamePermissions {
		UnstageBranch(stage, _role)
	}

}

func (stage *Stage) UnstageBranchState(state *State) {

	// check if instance is already staged
	if !IsStaged(stage, state) {
		return
	}

	state.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if state.Parent != nil {
		UnstageBranch(stage, state.Parent)
	}
	if state.Entry != nil {
		UnstageBranch(stage, state.Entry)
	}
	if state.Exit != nil {
		UnstageBranch(stage, state.Exit)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range state.SubStates {
		UnstageBranch(stage, _state)
	}
	for _, _diagram := range state.Diagrams {
		UnstageBranch(stage, _diagram)
	}
	for _, _activities := range state.Activities {
		UnstageBranch(stage, _activities)
	}

}

func (stage *Stage) UnstageBranchStateMachine(statemachine *StateMachine) {

	// check if instance is already staged
	if !IsStaged(stage, statemachine) {
		return
	}

	statemachine.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if statemachine.InitialState != nil {
		UnstageBranch(stage, statemachine.InitialState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range statemachine.States {
		UnstageBranch(stage, _state)
	}
	for _, _diagram := range statemachine.Diagrams {
		UnstageBranch(stage, _diagram)
	}

}

func (stage *Stage) UnstageBranchStateShape(stateshape *StateShape) {

	// check if instance is already staged
	if !IsStaged(stage, stateshape) {
		return
	}

	stateshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stateshape.State != nil {
		UnstageBranch(stage, stateshape.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTransition(transition *Transition) {

	// check if instance is already staged
	if !IsStaged(stage, transition) {
		return
	}

	transition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transition.Start != nil {
		UnstageBranch(stage, transition.Start)
	}
	if transition.End != nil {
		UnstageBranch(stage, transition.End)
	}
	if transition.Guard != nil {
		UnstageBranch(stage, transition.Guard)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range transition.RolesWithPermissions {
		UnstageBranch(stage, _role)
	}
	for _, _messagetype := range transition.GeneratedMessages {
		UnstageBranch(stage, _messagetype)
	}
	for _, _diagram := range transition.Diagrams {
		UnstageBranch(stage, _diagram)
	}

}

func (stage *Stage) UnstageBranchTransition_Shape(transition_shape *Transition_Shape) {

	// check if instance is already staged
	if !IsStaged(stage, transition_shape) {
		return
	}

	transition_shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transition_shape.Transition != nil {
		UnstageBranch(stage, transition_shape.Transition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Action) GongReconstructPointersFromReferences(stage *Stage, instance *Action) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Activities) GongReconstructPointersFromReferences(stage *Stage, instance *Activities) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Architecture) GongReconstructPointersFromReferences(stage *Stage, instance *Architecture) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.StateMachines = reference.StateMachines[:0]
	for _, _b := range instance.StateMachines {
		reference.StateMachines = append(reference.StateMachines, stage.StateMachines_reference[_b])
	}
	reference.Roles = reference.Roles[:0]
	for _, _b := range instance.Roles {
		reference.Roles = append(reference.Roles, stage.Roles_reference[_b])
	}
}

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
	reference.State_Shapes = reference.State_Shapes[:0]
	for _, _b := range instance.State_Shapes {
		reference.State_Shapes = append(reference.State_Shapes, stage.StateShapes_reference[_b])
	}
	reference.StatesWhoseNodeIsExpanded = reference.StatesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.StatesWhoseNodeIsExpanded {
		reference.StatesWhoseNodeIsExpanded = append(reference.StatesWhoseNodeIsExpanded, stage.States_reference[_b])
	}
	reference.Transition_Shapes = reference.Transition_Shapes[:0]
	for _, _b := range instance.Transition_Shapes {
		reference.Transition_Shapes = append(reference.Transition_Shapes, stage.Transition_Shapes_reference[_b])
	}
	reference.Note_Shapes = reference.Note_Shapes[:0]
	for _, _b := range instance.Note_Shapes {
		reference.Note_Shapes = append(reference.Note_Shapes, stage.NoteShapes_reference[_b])
	}
	reference.NoteState_Shapes = reference.NoteState_Shapes[:0]
	for _, _b := range instance.NoteState_Shapes {
		reference.NoteState_Shapes = append(reference.NoteState_Shapes, stage.NoteStateShapes_reference[_b])
	}
	reference.NoteTransition_Shapes = reference.NoteTransition_Shapes[:0]
	for _, _b := range instance.NoteTransition_Shapes {
		reference.NoteTransition_Shapes = append(reference.NoteTransition_Shapes, stage.NoteTransitionShapes_reference[_b])
	}
}

func (reference *Guard) GongReconstructPointersFromReferences(stage *Stage, instance *Guard) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Kill) GongReconstructPointersFromReferences(stage *Stage, instance *Kill) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.Diagrams = reference.Diagrams[:0]
	for _, _b := range instance.Diagrams {
		reference.Diagrams = append(reference.Diagrams, stage.Diagrams_reference[_b])
	}
	reference.RootStateMachines = reference.RootStateMachines[:0]
	for _, _b := range instance.RootStateMachines {
		reference.RootStateMachines = append(reference.RootStateMachines, stage.StateMachines_reference[_b])
	}
	reference.StateMachinesWhoseNodeIsExpanded = reference.StateMachinesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.StateMachinesWhoseNodeIsExpanded {
		reference.StateMachinesWhoseNodeIsExpanded = append(reference.StateMachinesWhoseNodeIsExpanded, stage.StateMachines_reference[_b])
	}
	reference.RootNotes = reference.RootNotes[:0]
	for _, _b := range instance.RootNotes {
		reference.RootNotes = append(reference.RootNotes, stage.Notes_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
	reference.SubLibrariesWhoseNodeIsExpanded = reference.SubLibrariesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SubLibrariesWhoseNodeIsExpanded {
		reference.SubLibrariesWhoseNodeIsExpanded = append(reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference[_b])
	}
}

func (reference *Message) GongReconstructPointersFromReferences(stage *Stage, instance *Message) {
	// insertion point for pointers field
	if instance.MessageType != nil {
		reference.MessageType = stage.MessageTypes_reference[instance.MessageType]
	}
	if instance.OriginTransition != nil {
		reference.OriginTransition = stage.Transitions_reference[instance.OriginTransition]
	}
	// insertion point for slice of pointers field
}

func (reference *MessageType) GongReconstructPointersFromReferences(stage *Stage, instance *MessageType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.States = reference.States[:0]
	for _, _b := range instance.States {
		reference.States = append(reference.States, stage.States_reference[_b])
	}
	reference.Transitions = reference.Transitions[:0]
	for _, _b := range instance.Transitions {
		reference.Transitions = append(reference.Transitions, stage.Transitions_reference[_b])
	}
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteStateShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteStateShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.State != nil {
		reference.State = stage.States_reference[instance.State]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteTransitionShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteTransitionShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Transition != nil {
		reference.Transition = stage.Transitions_reference[instance.Transition]
	}
	// insertion point for slice of pointers field
}

func (reference *Object) GongReconstructPointersFromReferences(stage *Stage, instance *Object) {
	// insertion point for pointers field
	if instance.State != nil {
		reference.State = stage.States_reference[instance.State]
	}
	// insertion point for slice of pointers field
	reference.Messages = reference.Messages[:0]
	for _, _b := range instance.Messages {
		reference.Messages = append(reference.Messages, stage.Messages_reference[_b])
	}
}

func (reference *Role) GongReconstructPointersFromReferences(stage *Stage, instance *Role) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.RolesWithSamePermissions = reference.RolesWithSamePermissions[:0]
	for _, _b := range instance.RolesWithSamePermissions {
		reference.RolesWithSamePermissions = append(reference.RolesWithSamePermissions, stage.Roles_reference[_b])
	}
}

func (reference *State) GongReconstructPointersFromReferences(stage *Stage, instance *State) {
	// insertion point for pointers field
	if instance.Parent != nil {
		reference.Parent = stage.States_reference[instance.Parent]
	}
	if instance.Entry != nil {
		reference.Entry = stage.Actions_reference[instance.Entry]
	}
	if instance.Exit != nil {
		reference.Exit = stage.Actions_reference[instance.Exit]
	}
	// insertion point for slice of pointers field
	reference.SubStates = reference.SubStates[:0]
	for _, _b := range instance.SubStates {
		reference.SubStates = append(reference.SubStates, stage.States_reference[_b])
	}
	reference.Diagrams = reference.Diagrams[:0]
	for _, _b := range instance.Diagrams {
		reference.Diagrams = append(reference.Diagrams, stage.Diagrams_reference[_b])
	}
	reference.Activities = reference.Activities[:0]
	for _, _b := range instance.Activities {
		reference.Activities = append(reference.Activities, stage.Activitiess_reference[_b])
	}
}

func (reference *StateMachine) GongReconstructPointersFromReferences(stage *Stage, instance *StateMachine) {
	// insertion point for pointers field
	if instance.InitialState != nil {
		reference.InitialState = stage.States_reference[instance.InitialState]
	}
	// insertion point for slice of pointers field
	reference.States = reference.States[:0]
	for _, _b := range instance.States {
		reference.States = append(reference.States, stage.States_reference[_b])
	}
	reference.Diagrams = reference.Diagrams[:0]
	for _, _b := range instance.Diagrams {
		reference.Diagrams = append(reference.Diagrams, stage.Diagrams_reference[_b])
	}
}

func (reference *StateShape) GongReconstructPointersFromReferences(stage *Stage, instance *StateShape) {
	// insertion point for pointers field
	if instance.State != nil {
		reference.State = stage.States_reference[instance.State]
	}
	// insertion point for slice of pointers field
}

func (reference *Transition) GongReconstructPointersFromReferences(stage *Stage, instance *Transition) {
	// insertion point for pointers field
	if instance.Start != nil {
		reference.Start = stage.States_reference[instance.Start]
	}
	if instance.End != nil {
		reference.End = stage.States_reference[instance.End]
	}
	if instance.Guard != nil {
		reference.Guard = stage.Guards_reference[instance.Guard]
	}
	// insertion point for slice of pointers field
	reference.RolesWithPermissions = reference.RolesWithPermissions[:0]
	for _, _b := range instance.RolesWithPermissions {
		reference.RolesWithPermissions = append(reference.RolesWithPermissions, stage.Roles_reference[_b])
	}
	reference.GeneratedMessages = reference.GeneratedMessages[:0]
	for _, _b := range instance.GeneratedMessages {
		reference.GeneratedMessages = append(reference.GeneratedMessages, stage.MessageTypes_reference[_b])
	}
	reference.Diagrams = reference.Diagrams[:0]
	for _, _b := range instance.Diagrams {
		reference.Diagrams = append(reference.Diagrams, stage.Diagrams_reference[_b])
	}
}

func (reference *Transition_Shape) GongReconstructPointersFromReferences(stage *Stage, instance *Transition_Shape) {
	// insertion point for pointers field
	if instance.Transition != nil {
		reference.Transition = stage.Transitions_reference[instance.Transition]
	}
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Action) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Activities) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Architecture) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _StateMachines []*StateMachine
	for _, _reference := range reference.StateMachines {
		if _instance, ok := stage.StateMachines_instance[_reference]; ok {
			_StateMachines = append(_StateMachines, _instance)
		}
	}
	reference.StateMachines = _StateMachines
	var _Roles []*Role
	for _, _reference := range reference.Roles {
		if _instance, ok := stage.Roles_instance[_reference]; ok {
			_Roles = append(_Roles, _instance)
		}
	}
	reference.Roles = _Roles
}

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
	var _State_Shapes []*StateShape
	for _, _reference := range reference.State_Shapes {
		if _instance, ok := stage.StateShapes_instance[_reference]; ok {
			_State_Shapes = append(_State_Shapes, _instance)
		}
	}
	reference.State_Shapes = _State_Shapes
	var _StatesWhoseNodeIsExpanded []*State
	for _, _reference := range reference.StatesWhoseNodeIsExpanded {
		if _instance, ok := stage.States_instance[_reference]; ok {
			_StatesWhoseNodeIsExpanded = append(_StatesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.StatesWhoseNodeIsExpanded = _StatesWhoseNodeIsExpanded
	var _Transition_Shapes []*Transition_Shape
	for _, _reference := range reference.Transition_Shapes {
		if _instance, ok := stage.Transition_Shapes_instance[_reference]; ok {
			_Transition_Shapes = append(_Transition_Shapes, _instance)
		}
	}
	reference.Transition_Shapes = _Transition_Shapes
	var _Note_Shapes []*NoteShape
	for _, _reference := range reference.Note_Shapes {
		if _instance, ok := stage.NoteShapes_instance[_reference]; ok {
			_Note_Shapes = append(_Note_Shapes, _instance)
		}
	}
	reference.Note_Shapes = _Note_Shapes
	var _NoteState_Shapes []*NoteStateShape
	for _, _reference := range reference.NoteState_Shapes {
		if _instance, ok := stage.NoteStateShapes_instance[_reference]; ok {
			_NoteState_Shapes = append(_NoteState_Shapes, _instance)
		}
	}
	reference.NoteState_Shapes = _NoteState_Shapes
	var _NoteTransition_Shapes []*NoteTransitionShape
	for _, _reference := range reference.NoteTransition_Shapes {
		if _instance, ok := stage.NoteTransitionShapes_instance[_reference]; ok {
			_NoteTransition_Shapes = append(_NoteTransition_Shapes, _instance)
		}
	}
	reference.NoteTransition_Shapes = _NoteTransition_Shapes
}

func (reference *Guard) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Kill) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
	var _Diagrams []*Diagram
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
	var _RootStateMachines []*StateMachine
	for _, _reference := range reference.RootStateMachines {
		if _instance, ok := stage.StateMachines_instance[_reference]; ok {
			_RootStateMachines = append(_RootStateMachines, _instance)
		}
	}
	reference.RootStateMachines = _RootStateMachines
	var _StateMachinesWhoseNodeIsExpanded []*StateMachine
	for _, _reference := range reference.StateMachinesWhoseNodeIsExpanded {
		if _instance, ok := stage.StateMachines_instance[_reference]; ok {
			_StateMachinesWhoseNodeIsExpanded = append(_StateMachinesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.StateMachinesWhoseNodeIsExpanded = _StateMachinesWhoseNodeIsExpanded
	var _RootNotes []*Note
	for _, _reference := range reference.RootNotes {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_RootNotes = append(_RootNotes, _instance)
		}
	}
	reference.RootNotes = _RootNotes
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
	var _SubLibrariesWhoseNodeIsExpanded []*Library
	for _, _reference := range reference.SubLibrariesWhoseNodeIsExpanded {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibrariesWhoseNodeIsExpanded = append(_SubLibrariesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SubLibrariesWhoseNodeIsExpanded = _SubLibrariesWhoseNodeIsExpanded
}

func (reference *Message) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.MessageType; _reference != nil {
		reference.MessageType = nil
		if _instance, ok := stage.MessageTypes_instance[_reference]; ok {
			reference.MessageType = _instance
		}
	}
	if _reference := reference.OriginTransition; _reference != nil {
		reference.OriginTransition = nil
		if _instance, ok := stage.Transitions_instance[_reference]; ok {
			reference.OriginTransition = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *MessageType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _States []*State
	for _, _reference := range reference.States {
		if _instance, ok := stage.States_instance[_reference]; ok {
			_States = append(_States, _instance)
		}
	}
	reference.States = _States
	var _Transitions []*Transition
	for _, _reference := range reference.Transitions {
		if _instance, ok := stage.Transitions_instance[_reference]; ok {
			_Transitions = append(_Transitions, _instance)
		}
	}
	reference.Transitions = _Transitions
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteStateShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.State; _reference != nil {
		reference.State = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.State = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteTransitionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Transition; _reference != nil {
		reference.Transition = nil
		if _instance, ok := stage.Transitions_instance[_reference]; ok {
			reference.Transition = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Object) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.State; _reference != nil {
		reference.State = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.State = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Messages []*Message
	for _, _reference := range reference.Messages {
		if _instance, ok := stage.Messages_instance[_reference]; ok {
			_Messages = append(_Messages, _instance)
		}
	}
	reference.Messages = _Messages
}

func (reference *Role) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _RolesWithSamePermissions []*Role
	for _, _reference := range reference.RolesWithSamePermissions {
		if _instance, ok := stage.Roles_instance[_reference]; ok {
			_RolesWithSamePermissions = append(_RolesWithSamePermissions, _instance)
		}
	}
	reference.RolesWithSamePermissions = _RolesWithSamePermissions
}

func (reference *State) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Parent; _reference != nil {
		reference.Parent = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.Parent = _instance
		}
	}
	if _reference := reference.Entry; _reference != nil {
		reference.Entry = nil
		if _instance, ok := stage.Actions_instance[_reference]; ok {
			reference.Entry = _instance
		}
	}
	if _reference := reference.Exit; _reference != nil {
		reference.Exit = nil
		if _instance, ok := stage.Actions_instance[_reference]; ok {
			reference.Exit = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _SubStates []*State
	for _, _reference := range reference.SubStates {
		if _instance, ok := stage.States_instance[_reference]; ok {
			_SubStates = append(_SubStates, _instance)
		}
	}
	reference.SubStates = _SubStates
	var _Diagrams []*Diagram
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
	var _Activities []*Activities
	for _, _reference := range reference.Activities {
		if _instance, ok := stage.Activitiess_instance[_reference]; ok {
			_Activities = append(_Activities, _instance)
		}
	}
	reference.Activities = _Activities
}

func (reference *StateMachine) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.InitialState; _reference != nil {
		reference.InitialState = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.InitialState = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _States []*State
	for _, _reference := range reference.States {
		if _instance, ok := stage.States_instance[_reference]; ok {
			_States = append(_States, _instance)
		}
	}
	reference.States = _States
	var _Diagrams []*Diagram
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
}

func (reference *StateShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.State; _reference != nil {
		reference.State = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.State = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Transition) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Start; _reference != nil {
		reference.Start = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.Start = _instance
		}
	}
	if _reference := reference.End; _reference != nil {
		reference.End = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.End = _instance
		}
	}
	if _reference := reference.Guard; _reference != nil {
		reference.Guard = nil
		if _instance, ok := stage.Guards_instance[_reference]; ok {
			reference.Guard = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _RolesWithPermissions []*Role
	for _, _reference := range reference.RolesWithPermissions {
		if _instance, ok := stage.Roles_instance[_reference]; ok {
			_RolesWithPermissions = append(_RolesWithPermissions, _instance)
		}
	}
	reference.RolesWithPermissions = _RolesWithPermissions
	var _GeneratedMessages []*MessageType
	for _, _reference := range reference.GeneratedMessages {
		if _instance, ok := stage.MessageTypes_instance[_reference]; ok {
			_GeneratedMessages = append(_GeneratedMessages, _instance)
		}
	}
	reference.GeneratedMessages = _GeneratedMessages
	var _Diagrams []*Diagram
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
}

func (reference *Transition_Shape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Transition; _reference != nil {
		reference.Transition = nil
		if _instance, ok := stage.Transitions_instance[_reference]; ok {
			reference.Transition = _instance
		}
	}
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (action *Action) GongDiff(stage *Stage, actionOther *Action) (diffs []string) {
	// insertion point for field diffs
	if action.Name != actionOther.Name {
		diffs = append(diffs, action.GongMarshallField(stage, "Name"))
	}
	if action.Criticality != actionOther.Criticality {
		diffs = append(diffs, action.GongMarshallField(stage, "Criticality"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (activities *Activities) GongDiff(stage *Stage, activitiesOther *Activities) (diffs []string) {
	// insertion point for field diffs
	if activities.Name != activitiesOther.Name {
		diffs = append(diffs, activities.GongMarshallField(stage, "Name"))
	}
	if activities.Criticality != activitiesOther.Criticality {
		diffs = append(diffs, activities.GongMarshallField(stage, "Criticality"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (architecture *Architecture) GongDiff(stage *Stage, architectureOther *Architecture) (diffs []string) {
	// insertion point for field diffs
	if architecture.Name != architectureOther.Name {
		diffs = append(diffs, architecture.GongMarshallField(stage, "Name"))
	}
	StateMachinesDifferent := false
	if len(architecture.StateMachines) != len(architectureOther.StateMachines) {
		StateMachinesDifferent = true
	} else {
		for i := range architecture.StateMachines {
			if (architecture.StateMachines[i] == nil) != (architectureOther.StateMachines[i] == nil) {
				StateMachinesDifferent = true
				break
			} else if architecture.StateMachines[i] != nil && architectureOther.StateMachines[i] != nil {
				// this is a pointer comparaison
				if architecture.StateMachines[i] != architectureOther.StateMachines[i] {
					StateMachinesDifferent = true
					break
				}
			}
		}
	}
	if StateMachinesDifferent {
		ops := Diff(stage, architecture, architectureOther, "StateMachines", architectureOther.StateMachines, architecture.StateMachines)
		diffs = append(diffs, ops)
	}
	RolesDifferent := false
	if len(architecture.Roles) != len(architectureOther.Roles) {
		RolesDifferent = true
	} else {
		for i := range architecture.Roles {
			if (architecture.Roles[i] == nil) != (architectureOther.Roles[i] == nil) {
				RolesDifferent = true
				break
			} else if architecture.Roles[i] != nil && architectureOther.Roles[i] != nil {
				// this is a pointer comparaison
				if architecture.Roles[i] != architectureOther.Roles[i] {
					RolesDifferent = true
					break
				}
			}
		}
	}
	if RolesDifferent {
		ops := Diff(stage, architecture, architectureOther, "Roles", architectureOther.Roles, architecture.Roles)
		diffs = append(diffs, ops)
	}
	if architecture.NbPixPerCharacter != architectureOther.NbPixPerCharacter {
		diffs = append(diffs, architecture.GongMarshallField(stage, "NbPixPerCharacter"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagram *Diagram) GongDiff(stage *Stage, diagramOther *Diagram) (diffs []string) {
	// insertion point for field diffs
	if diagram.Name != diagramOther.Name {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Name"))
	}
	if diagram.IsChecked != diagramOther.IsChecked {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsChecked"))
	}
	if diagram.IsExpanded != diagramOther.IsExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsExpanded"))
	}
	if diagram.IsEditable_ != diagramOther.IsEditable_ {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsEditable_"))
	}
	if diagram.IsStatesNodeExpanded != diagramOther.IsStatesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsStatesNodeExpanded"))
	}
	if diagram.IsNotesNodeExpanded != diagramOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(diagram.NotesWhoseNodeIsExpanded) != len(diagramOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.NotesWhoseNodeIsExpanded {
			if (diagram.NotesWhoseNodeIsExpanded[i] == nil) != (diagramOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.NotesWhoseNodeIsExpanded[i] != nil && diagramOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.NotesWhoseNodeIsExpanded[i] != diagramOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "NotesWhoseNodeIsExpanded", diagramOther.NotesWhoseNodeIsExpanded, diagram.NotesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	State_ShapesDifferent := false
	if len(diagram.State_Shapes) != len(diagramOther.State_Shapes) {
		State_ShapesDifferent = true
	} else {
		for i := range diagram.State_Shapes {
			if (diagram.State_Shapes[i] == nil) != (diagramOther.State_Shapes[i] == nil) {
				State_ShapesDifferent = true
				break
			} else if diagram.State_Shapes[i] != nil && diagramOther.State_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.State_Shapes[i] != diagramOther.State_Shapes[i] {
					State_ShapesDifferent = true
					break
				}
			}
		}
	}
	if State_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "State_Shapes", diagramOther.State_Shapes, diagram.State_Shapes)
		diffs = append(diffs, ops)
	}
	StatesWhoseNodeIsExpandedDifferent := false
	if len(diagram.StatesWhoseNodeIsExpanded) != len(diagramOther.StatesWhoseNodeIsExpanded) {
		StatesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.StatesWhoseNodeIsExpanded {
			if (diagram.StatesWhoseNodeIsExpanded[i] == nil) != (diagramOther.StatesWhoseNodeIsExpanded[i] == nil) {
				StatesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.StatesWhoseNodeIsExpanded[i] != nil && diagramOther.StatesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.StatesWhoseNodeIsExpanded[i] != diagramOther.StatesWhoseNodeIsExpanded[i] {
					StatesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if StatesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagram, diagramOther, "StatesWhoseNodeIsExpanded", diagramOther.StatesWhoseNodeIsExpanded, diagram.StatesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	Transition_ShapesDifferent := false
	if len(diagram.Transition_Shapes) != len(diagramOther.Transition_Shapes) {
		Transition_ShapesDifferent = true
	} else {
		for i := range diagram.Transition_Shapes {
			if (diagram.Transition_Shapes[i] == nil) != (diagramOther.Transition_Shapes[i] == nil) {
				Transition_ShapesDifferent = true
				break
			} else if diagram.Transition_Shapes[i] != nil && diagramOther.Transition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Transition_Shapes[i] != diagramOther.Transition_Shapes[i] {
					Transition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Transition_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Transition_Shapes", diagramOther.Transition_Shapes, diagram.Transition_Shapes)
		diffs = append(diffs, ops)
	}
	Note_ShapesDifferent := false
	if len(diagram.Note_Shapes) != len(diagramOther.Note_Shapes) {
		Note_ShapesDifferent = true
	} else {
		for i := range diagram.Note_Shapes {
			if (diagram.Note_Shapes[i] == nil) != (diagramOther.Note_Shapes[i] == nil) {
				Note_ShapesDifferent = true
				break
			} else if diagram.Note_Shapes[i] != nil && diagramOther.Note_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Note_Shapes[i] != diagramOther.Note_Shapes[i] {
					Note_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Note_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Note_Shapes", diagramOther.Note_Shapes, diagram.Note_Shapes)
		diffs = append(diffs, ops)
	}
	NoteState_ShapesDifferent := false
	if len(diagram.NoteState_Shapes) != len(diagramOther.NoteState_Shapes) {
		NoteState_ShapesDifferent = true
	} else {
		for i := range diagram.NoteState_Shapes {
			if (diagram.NoteState_Shapes[i] == nil) != (diagramOther.NoteState_Shapes[i] == nil) {
				NoteState_ShapesDifferent = true
				break
			} else if diagram.NoteState_Shapes[i] != nil && diagramOther.NoteState_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.NoteState_Shapes[i] != diagramOther.NoteState_Shapes[i] {
					NoteState_ShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteState_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "NoteState_Shapes", diagramOther.NoteState_Shapes, diagram.NoteState_Shapes)
		diffs = append(diffs, ops)
	}
	NoteTransition_ShapesDifferent := false
	if len(diagram.NoteTransition_Shapes) != len(diagramOther.NoteTransition_Shapes) {
		NoteTransition_ShapesDifferent = true
	} else {
		for i := range diagram.NoteTransition_Shapes {
			if (diagram.NoteTransition_Shapes[i] == nil) != (diagramOther.NoteTransition_Shapes[i] == nil) {
				NoteTransition_ShapesDifferent = true
				break
			} else if diagram.NoteTransition_Shapes[i] != nil && diagramOther.NoteTransition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.NoteTransition_Shapes[i] != diagramOther.NoteTransition_Shapes[i] {
					NoteTransition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteTransition_ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "NoteTransition_Shapes", diagramOther.NoteTransition_Shapes, diagram.NoteTransition_Shapes)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (guard *Guard) GongDiff(stage *Stage, guardOther *Guard) (diffs []string) {
	// insertion point for field diffs
	if guard.Name != guardOther.Name {
		diffs = append(diffs, guard.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (kill *Kill) GongDiff(stage *Stage, killOther *Kill) (diffs []string) {
	// insertion point for field diffs
	if kill.Name != killOther.Name {
		diffs = append(diffs, kill.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (library *Library) GongDiff(stage *Stage, libraryOther *Library) (diffs []string) {
	// insertion point for field diffs
	if library.Name != libraryOther.Name {
		diffs = append(diffs, library.GongMarshallField(stage, "Name"))
	}
	SubLibrariesDifferent := false
	if len(library.SubLibraries) != len(libraryOther.SubLibraries) {
		SubLibrariesDifferent = true
	} else {
		for i := range library.SubLibraries {
			if (library.SubLibraries[i] == nil) != (libraryOther.SubLibraries[i] == nil) {
				SubLibrariesDifferent = true
				break
			} else if library.SubLibraries[i] != nil && libraryOther.SubLibraries[i] != nil {
				// this is a pointer comparaison
				if library.SubLibraries[i] != libraryOther.SubLibraries[i] {
					SubLibrariesDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries)
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
	}
	if library.IsExpanded != libraryOther.IsExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpanded"))
	}
	if library.LayoutDirection != libraryOther.LayoutDirection {
		diffs = append(diffs, library.GongMarshallField(stage, "LayoutDirection"))
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	DiagramsDifferent := false
	if len(library.Diagrams) != len(libraryOther.Diagrams) {
		DiagramsDifferent = true
	} else {
		for i := range library.Diagrams {
			if (library.Diagrams[i] == nil) != (libraryOther.Diagrams[i] == nil) {
				DiagramsDifferent = true
				break
			} else if library.Diagrams[i] != nil && libraryOther.Diagrams[i] != nil {
				// this is a pointer comparaison
				if library.Diagrams[i] != libraryOther.Diagrams[i] {
					DiagramsDifferent = true
					break
				}
			}
		}
	}
	if DiagramsDifferent {
		ops := Diff(stage, library, libraryOther, "Diagrams", libraryOther.Diagrams, library.Diagrams)
		diffs = append(diffs, ops)
	}
	RootStateMachinesDifferent := false
	if len(library.RootStateMachines) != len(libraryOther.RootStateMachines) {
		RootStateMachinesDifferent = true
	} else {
		for i := range library.RootStateMachines {
			if (library.RootStateMachines[i] == nil) != (libraryOther.RootStateMachines[i] == nil) {
				RootStateMachinesDifferent = true
				break
			} else if library.RootStateMachines[i] != nil && libraryOther.RootStateMachines[i] != nil {
				// this is a pointer comparaison
				if library.RootStateMachines[i] != libraryOther.RootStateMachines[i] {
					RootStateMachinesDifferent = true
					break
				}
			}
		}
	}
	if RootStateMachinesDifferent {
		ops := Diff(stage, library, libraryOther, "RootStateMachines", libraryOther.RootStateMachines, library.RootStateMachines)
		diffs = append(diffs, ops)
	}
	if library.IsStateMachinesNodeExpanded != libraryOther.IsStateMachinesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsStateMachinesNodeExpanded"))
	}
	StateMachinesWhoseNodeIsExpandedDifferent := false
	if len(library.StateMachinesWhoseNodeIsExpanded) != len(libraryOther.StateMachinesWhoseNodeIsExpanded) {
		StateMachinesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.StateMachinesWhoseNodeIsExpanded {
			if (library.StateMachinesWhoseNodeIsExpanded[i] == nil) != (libraryOther.StateMachinesWhoseNodeIsExpanded[i] == nil) {
				StateMachinesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.StateMachinesWhoseNodeIsExpanded[i] != nil && libraryOther.StateMachinesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.StateMachinesWhoseNodeIsExpanded[i] != libraryOther.StateMachinesWhoseNodeIsExpanded[i] {
					StateMachinesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if StateMachinesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "StateMachinesWhoseNodeIsExpanded", libraryOther.StateMachinesWhoseNodeIsExpanded, library.StateMachinesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	RootNotesDifferent := false
	if len(library.RootNotes) != len(libraryOther.RootNotes) {
		RootNotesDifferent = true
	} else {
		for i := range library.RootNotes {
			if (library.RootNotes[i] == nil) != (libraryOther.RootNotes[i] == nil) {
				RootNotesDifferent = true
				break
			} else if library.RootNotes[i] != nil && libraryOther.RootNotes[i] != nil {
				// this is a pointer comparaison
				if library.RootNotes[i] != libraryOther.RootNotes[i] {
					RootNotesDifferent = true
					break
				}
			}
		}
	}
	if RootNotesDifferent {
		ops := Diff(stage, library, libraryOther, "RootNotes", libraryOther.RootNotes, library.RootNotes)
		diffs = append(diffs, ops)
	}
	if library.IsNotesNodeExpanded != libraryOther.IsNotesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(library.NotesWhoseNodeIsExpanded) != len(libraryOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.NotesWhoseNodeIsExpanded {
			if (library.NotesWhoseNodeIsExpanded[i] == nil) != (libraryOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.NotesWhoseNodeIsExpanded[i] != nil && libraryOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.NotesWhoseNodeIsExpanded[i] != libraryOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "NotesWhoseNodeIsExpanded", libraryOther.NotesWhoseNodeIsExpanded, library.NotesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	SubLibrariesWhoseNodeIsExpandedDifferent := false
	if len(library.SubLibrariesWhoseNodeIsExpanded) != len(libraryOther.SubLibrariesWhoseNodeIsExpanded) {
		SubLibrariesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.SubLibrariesWhoseNodeIsExpanded {
			if (library.SubLibrariesWhoseNodeIsExpanded[i] == nil) != (libraryOther.SubLibrariesWhoseNodeIsExpanded[i] == nil) {
				SubLibrariesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.SubLibrariesWhoseNodeIsExpanded[i] != nil && libraryOther.SubLibrariesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.SubLibrariesWhoseNodeIsExpanded[i] != libraryOther.SubLibrariesWhoseNodeIsExpanded[i] {
					SubLibrariesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (message *Message) GongDiff(stage *Stage, messageOther *Message) (diffs []string) {
	// insertion point for field diffs
	if message.Name != messageOther.Name {
		diffs = append(diffs, message.GongMarshallField(stage, "Name"))
	}
	if message.IsSelected != messageOther.IsSelected {
		diffs = append(diffs, message.GongMarshallField(stage, "IsSelected"))
	}
	if (message.MessageType == nil) != (messageOther.MessageType == nil) {
		diffs = append(diffs, message.GongMarshallField(stage, "MessageType"))
	} else if message.MessageType != nil && messageOther.MessageType != nil {
		if message.MessageType != messageOther.MessageType {
			diffs = append(diffs, message.GongMarshallField(stage, "MessageType"))
		}
	}
	if (message.OriginTransition == nil) != (messageOther.OriginTransition == nil) {
		diffs = append(diffs, message.GongMarshallField(stage, "OriginTransition"))
	} else if message.OriginTransition != nil && messageOther.OriginTransition != nil {
		if message.OriginTransition != messageOther.OriginTransition {
			diffs = append(diffs, message.GongMarshallField(stage, "OriginTransition"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (messagetype *MessageType) GongDiff(stage *Stage, messagetypeOther *MessageType) (diffs []string) {
	// insertion point for field diffs
	if messagetype.Name != messagetypeOther.Name {
		diffs = append(diffs, messagetype.GongMarshallField(stage, "Name"))
	}
	if messagetype.Description != messagetypeOther.Description {
		diffs = append(diffs, messagetype.GongMarshallField(stage, "Description"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if note.ComputedPrefix != noteOther.ComputedPrefix {
		diffs = append(diffs, note.GongMarshallField(stage, "ComputedPrefix"))
	}
	if note.IsExpanded != noteOther.IsExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsExpanded"))
	}
	if note.LayoutDirection != noteOther.LayoutDirection {
		diffs = append(diffs, note.GongMarshallField(stage, "LayoutDirection"))
	}
	StatesDifferent := false
	if len(note.States) != len(noteOther.States) {
		StatesDifferent = true
	} else {
		for i := range note.States {
			if (note.States[i] == nil) != (noteOther.States[i] == nil) {
				StatesDifferent = true
				break
			} else if note.States[i] != nil && noteOther.States[i] != nil {
				// this is a pointer comparaison
				if note.States[i] != noteOther.States[i] {
					StatesDifferent = true
					break
				}
			}
		}
	}
	if StatesDifferent {
		ops := Diff(stage, note, noteOther, "States", noteOther.States, note.States)
		diffs = append(diffs, ops)
	}
	TransitionsDifferent := false
	if len(note.Transitions) != len(noteOther.Transitions) {
		TransitionsDifferent = true
	} else {
		for i := range note.Transitions {
			if (note.Transitions[i] == nil) != (noteOther.Transitions[i] == nil) {
				TransitionsDifferent = true
				break
			} else if note.Transitions[i] != nil && noteOther.Transitions[i] != nil {
				// this is a pointer comparaison
				if note.Transitions[i] != noteOther.Transitions[i] {
					TransitionsDifferent = true
					break
				}
			}
		}
	}
	if TransitionsDifferent {
		ops := Diff(stage, note, noteOther, "Transitions", noteOther.Transitions, note.Transitions)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteshape *NoteShape) GongDiff(stage *Stage, noteshapeOther *NoteShape) (diffs []string) {
	// insertion point for field diffs
	if noteshape.Name != noteshapeOther.Name {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Name"))
	}
	if (noteshape.Note == nil) != (noteshapeOther.Note == nil) {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
	} else if noteshape.Note != nil && noteshapeOther.Note != nil {
		if noteshape.Note != noteshapeOther.Note {
			diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
		}
	}
	if noteshape.OverideLayoutDirection != noteshapeOther.OverideLayoutDirection {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "OverideLayoutDirection"))
	}
	if noteshape.LayoutDirection != noteshapeOther.LayoutDirection {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "LayoutDirection"))
	}
	if noteshape.X != noteshapeOther.X {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "X"))
	}
	if noteshape.Y != noteshapeOther.Y {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Y"))
	}
	if noteshape.Width != noteshapeOther.Width {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Width"))
	}
	if noteshape.Height != noteshapeOther.Height {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Height"))
	}
	if noteshape.IsHidden != noteshapeOther.IsHidden {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notestateshape *NoteStateShape) GongDiff(stage *Stage, notestateshapeOther *NoteStateShape) (diffs []string) {
	// insertion point for field diffs
	if notestateshape.Name != notestateshapeOther.Name {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "Name"))
	}
	if (notestateshape.Note == nil) != (notestateshapeOther.Note == nil) {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "Note"))
	} else if notestateshape.Note != nil && notestateshapeOther.Note != nil {
		if notestateshape.Note != notestateshapeOther.Note {
			diffs = append(diffs, notestateshape.GongMarshallField(stage, "Note"))
		}
	}
	if (notestateshape.State == nil) != (notestateshapeOther.State == nil) {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "State"))
	} else if notestateshape.State != nil && notestateshapeOther.State != nil {
		if notestateshape.State != notestateshapeOther.State {
			diffs = append(diffs, notestateshape.GongMarshallField(stage, "State"))
		}
	}
	if notestateshape.StartRatio != notestateshapeOther.StartRatio {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "StartRatio"))
	}
	if notestateshape.EndRatio != notestateshapeOther.EndRatio {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "EndRatio"))
	}
	if notestateshape.StartOrientation != notestateshapeOther.StartOrientation {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "StartOrientation"))
	}
	if notestateshape.EndOrientation != notestateshapeOther.EndOrientation {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "EndOrientation"))
	}
	if notestateshape.CornerOffsetRatio != notestateshapeOther.CornerOffsetRatio {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notestateshape.IsHidden != notestateshapeOther.IsHidden {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notetransitionshape *NoteTransitionShape) GongDiff(stage *Stage, notetransitionshapeOther *NoteTransitionShape) (diffs []string) {
	// insertion point for field diffs
	if notetransitionshape.Name != notetransitionshapeOther.Name {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "Name"))
	}
	if (notetransitionshape.Note == nil) != (notetransitionshapeOther.Note == nil) {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "Note"))
	} else if notetransitionshape.Note != nil && notetransitionshapeOther.Note != nil {
		if notetransitionshape.Note != notetransitionshapeOther.Note {
			diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "Note"))
		}
	}
	if (notetransitionshape.Transition == nil) != (notetransitionshapeOther.Transition == nil) {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "Transition"))
	} else if notetransitionshape.Transition != nil && notetransitionshapeOther.Transition != nil {
		if notetransitionshape.Transition != notetransitionshapeOther.Transition {
			diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "Transition"))
		}
	}
	if notetransitionshape.StartRatio != notetransitionshapeOther.StartRatio {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "StartRatio"))
	}
	if notetransitionshape.EndRatio != notetransitionshapeOther.EndRatio {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "EndRatio"))
	}
	if notetransitionshape.StartOrientation != notetransitionshapeOther.StartOrientation {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if notetransitionshape.EndOrientation != notetransitionshapeOther.EndOrientation {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if notetransitionshape.CornerOffsetRatio != notetransitionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notetransitionshape.IsHidden != notetransitionshapeOther.IsHidden {
		diffs = append(diffs, notetransitionshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (object *Object) GongDiff(stage *Stage, objectOther *Object) (diffs []string) {
	// insertion point for field diffs
	if object.Name != objectOther.Name {
		diffs = append(diffs, object.GongMarshallField(stage, "Name"))
	}
	if (object.State == nil) != (objectOther.State == nil) {
		diffs = append(diffs, object.GongMarshallField(stage, "State"))
	} else if object.State != nil && objectOther.State != nil {
		if object.State != objectOther.State {
			diffs = append(diffs, object.GongMarshallField(stage, "State"))
		}
	}
	if object.IsSelected != objectOther.IsSelected {
		diffs = append(diffs, object.GongMarshallField(stage, "IsSelected"))
	}
	if object.Rank != objectOther.Rank {
		diffs = append(diffs, object.GongMarshallField(stage, "Rank"))
	}
	if object.DOF != objectOther.DOF {
		diffs = append(diffs, object.GongMarshallField(stage, "DOF"))
	}
	MessagesDifferent := false
	if len(object.Messages) != len(objectOther.Messages) {
		MessagesDifferent = true
	} else {
		for i := range object.Messages {
			if (object.Messages[i] == nil) != (objectOther.Messages[i] == nil) {
				MessagesDifferent = true
				break
			} else if object.Messages[i] != nil && objectOther.Messages[i] != nil {
				// this is a pointer comparaison
				if object.Messages[i] != objectOther.Messages[i] {
					MessagesDifferent = true
					break
				}
			}
		}
	}
	if MessagesDifferent {
		ops := Diff(stage, object, objectOther, "Messages", objectOther.Messages, object.Messages)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (role *Role) GongDiff(stage *Stage, roleOther *Role) (diffs []string) {
	// insertion point for field diffs
	if role.Name != roleOther.Name {
		diffs = append(diffs, role.GongMarshallField(stage, "Name"))
	}
	if role.Acronym != roleOther.Acronym {
		diffs = append(diffs, role.GongMarshallField(stage, "Acronym"))
	}
	RolesWithSamePermissionsDifferent := false
	if len(role.RolesWithSamePermissions) != len(roleOther.RolesWithSamePermissions) {
		RolesWithSamePermissionsDifferent = true
	} else {
		for i := range role.RolesWithSamePermissions {
			if (role.RolesWithSamePermissions[i] == nil) != (roleOther.RolesWithSamePermissions[i] == nil) {
				RolesWithSamePermissionsDifferent = true
				break
			} else if role.RolesWithSamePermissions[i] != nil && roleOther.RolesWithSamePermissions[i] != nil {
				// this is a pointer comparaison
				if role.RolesWithSamePermissions[i] != roleOther.RolesWithSamePermissions[i] {
					RolesWithSamePermissionsDifferent = true
					break
				}
			}
		}
	}
	if RolesWithSamePermissionsDifferent {
		ops := Diff(stage, role, roleOther, "RolesWithSamePermissions", roleOther.RolesWithSamePermissions, role.RolesWithSamePermissions)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (state *State) GongDiff(stage *Stage, stateOther *State) (diffs []string) {
	// insertion point for field diffs
	if state.Name != stateOther.Name {
		diffs = append(diffs, state.GongMarshallField(stage, "Name"))
	}
	if (state.Parent == nil) != (stateOther.Parent == nil) {
		diffs = append(diffs, state.GongMarshallField(stage, "Parent"))
	} else if state.Parent != nil && stateOther.Parent != nil {
		if state.Parent != stateOther.Parent {
			diffs = append(diffs, state.GongMarshallField(stage, "Parent"))
		}
	}
	if state.IsDecisionNode != stateOther.IsDecisionNode {
		diffs = append(diffs, state.GongMarshallField(stage, "IsDecisionNode"))
	}
	if state.IsFictious != stateOther.IsFictious {
		diffs = append(diffs, state.GongMarshallField(stage, "IsFictious"))
	}
	if state.IsEndState != stateOther.IsEndState {
		diffs = append(diffs, state.GongMarshallField(stage, "IsEndState"))
	}
	SubStatesDifferent := false
	if len(state.SubStates) != len(stateOther.SubStates) {
		SubStatesDifferent = true
	} else {
		for i := range state.SubStates {
			if (state.SubStates[i] == nil) != (stateOther.SubStates[i] == nil) {
				SubStatesDifferent = true
				break
			} else if state.SubStates[i] != nil && stateOther.SubStates[i] != nil {
				// this is a pointer comparaison
				if state.SubStates[i] != stateOther.SubStates[i] {
					SubStatesDifferent = true
					break
				}
			}
		}
	}
	if SubStatesDifferent {
		ops := Diff(stage, state, stateOther, "SubStates", stateOther.SubStates, state.SubStates)
		diffs = append(diffs, ops)
	}
	DiagramsDifferent := false
	if len(state.Diagrams) != len(stateOther.Diagrams) {
		DiagramsDifferent = true
	} else {
		for i := range state.Diagrams {
			if (state.Diagrams[i] == nil) != (stateOther.Diagrams[i] == nil) {
				DiagramsDifferent = true
				break
			} else if state.Diagrams[i] != nil && stateOther.Diagrams[i] != nil {
				// this is a pointer comparaison
				if state.Diagrams[i] != stateOther.Diagrams[i] {
					DiagramsDifferent = true
					break
				}
			}
		}
	}
	if DiagramsDifferent {
		ops := Diff(stage, state, stateOther, "Diagrams", stateOther.Diagrams, state.Diagrams)
		diffs = append(diffs, ops)
	}
	if (state.Entry == nil) != (stateOther.Entry == nil) {
		diffs = append(diffs, state.GongMarshallField(stage, "Entry"))
	} else if state.Entry != nil && stateOther.Entry != nil {
		if state.Entry != stateOther.Entry {
			diffs = append(diffs, state.GongMarshallField(stage, "Entry"))
		}
	}
	ActivitiesDifferent := false
	if len(state.Activities) != len(stateOther.Activities) {
		ActivitiesDifferent = true
	} else {
		for i := range state.Activities {
			if (state.Activities[i] == nil) != (stateOther.Activities[i] == nil) {
				ActivitiesDifferent = true
				break
			} else if state.Activities[i] != nil && stateOther.Activities[i] != nil {
				// this is a pointer comparaison
				if state.Activities[i] != stateOther.Activities[i] {
					ActivitiesDifferent = true
					break
				}
			}
		}
	}
	if ActivitiesDifferent {
		ops := Diff(stage, state, stateOther, "Activities", stateOther.Activities, state.Activities)
		diffs = append(diffs, ops)
	}
	if (state.Exit == nil) != (stateOther.Exit == nil) {
		diffs = append(diffs, state.GongMarshallField(stage, "Exit"))
	} else if state.Exit != nil && stateOther.Exit != nil {
		if state.Exit != stateOther.Exit {
			diffs = append(diffs, state.GongMarshallField(stage, "Exit"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (statemachine *StateMachine) GongDiff(stage *Stage, statemachineOther *StateMachine) (diffs []string) {
	// insertion point for field diffs
	if statemachine.Name != statemachineOther.Name {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "Name"))
	}
	if statemachine.ComputedPrefix != statemachineOther.ComputedPrefix {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "ComputedPrefix"))
	}
	if statemachine.IsExpanded != statemachineOther.IsExpanded {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "IsExpanded"))
	}
	if statemachine.LayoutDirection != statemachineOther.LayoutDirection {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "LayoutDirection"))
	}
	StatesDifferent := false
	if len(statemachine.States) != len(statemachineOther.States) {
		StatesDifferent = true
	} else {
		for i := range statemachine.States {
			if (statemachine.States[i] == nil) != (statemachineOther.States[i] == nil) {
				StatesDifferent = true
				break
			} else if statemachine.States[i] != nil && statemachineOther.States[i] != nil {
				// this is a pointer comparaison
				if statemachine.States[i] != statemachineOther.States[i] {
					StatesDifferent = true
					break
				}
			}
		}
	}
	if StatesDifferent {
		ops := Diff(stage, statemachine, statemachineOther, "States", statemachineOther.States, statemachine.States)
		diffs = append(diffs, ops)
	}
	DiagramsDifferent := false
	if len(statemachine.Diagrams) != len(statemachineOther.Diagrams) {
		DiagramsDifferent = true
	} else {
		for i := range statemachine.Diagrams {
			if (statemachine.Diagrams[i] == nil) != (statemachineOther.Diagrams[i] == nil) {
				DiagramsDifferent = true
				break
			} else if statemachine.Diagrams[i] != nil && statemachineOther.Diagrams[i] != nil {
				// this is a pointer comparaison
				if statemachine.Diagrams[i] != statemachineOther.Diagrams[i] {
					DiagramsDifferent = true
					break
				}
			}
		}
	}
	if DiagramsDifferent {
		ops := Diff(stage, statemachine, statemachineOther, "Diagrams", statemachineOther.Diagrams, statemachine.Diagrams)
		diffs = append(diffs, ops)
	}
	if (statemachine.InitialState == nil) != (statemachineOther.InitialState == nil) {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "InitialState"))
	} else if statemachine.InitialState != nil && statemachineOther.InitialState != nil {
		if statemachine.InitialState != statemachineOther.InitialState {
			diffs = append(diffs, statemachine.GongMarshallField(stage, "InitialState"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stateshape *StateShape) GongDiff(stage *Stage, stateshapeOther *StateShape) (diffs []string) {
	// insertion point for field diffs
	if stateshape.Name != stateshapeOther.Name {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "Name"))
	}
	if (stateshape.State == nil) != (stateshapeOther.State == nil) {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "State"))
	} else if stateshape.State != nil && stateshapeOther.State != nil {
		if stateshape.State != stateshapeOther.State {
			diffs = append(diffs, stateshape.GongMarshallField(stage, "State"))
		}
	}
	if stateshape.X != stateshapeOther.X {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "X"))
	}
	if stateshape.Y != stateshapeOther.Y {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "Y"))
	}
	if stateshape.Width != stateshapeOther.Width {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "Width"))
	}
	if stateshape.Height != stateshapeOther.Height {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "Height"))
	}
	if stateshape.IsHidden != stateshapeOther.IsHidden {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (transition *Transition) GongDiff(stage *Stage, transitionOther *Transition) (diffs []string) {
	// insertion point for field diffs
	if transition.Name != transitionOther.Name {
		diffs = append(diffs, transition.GongMarshallField(stage, "Name"))
	}
	if (transition.Start == nil) != (transitionOther.Start == nil) {
		diffs = append(diffs, transition.GongMarshallField(stage, "Start"))
	} else if transition.Start != nil && transitionOther.Start != nil {
		if transition.Start != transitionOther.Start {
			diffs = append(diffs, transition.GongMarshallField(stage, "Start"))
		}
	}
	if (transition.End == nil) != (transitionOther.End == nil) {
		diffs = append(diffs, transition.GongMarshallField(stage, "End"))
	} else if transition.End != nil && transitionOther.End != nil {
		if transition.End != transitionOther.End {
			diffs = append(diffs, transition.GongMarshallField(stage, "End"))
		}
	}
	RolesWithPermissionsDifferent := false
	if len(transition.RolesWithPermissions) != len(transitionOther.RolesWithPermissions) {
		RolesWithPermissionsDifferent = true
	} else {
		for i := range transition.RolesWithPermissions {
			if (transition.RolesWithPermissions[i] == nil) != (transitionOther.RolesWithPermissions[i] == nil) {
				RolesWithPermissionsDifferent = true
				break
			} else if transition.RolesWithPermissions[i] != nil && transitionOther.RolesWithPermissions[i] != nil {
				// this is a pointer comparaison
				if transition.RolesWithPermissions[i] != transitionOther.RolesWithPermissions[i] {
					RolesWithPermissionsDifferent = true
					break
				}
			}
		}
	}
	if RolesWithPermissionsDifferent {
		ops := Diff(stage, transition, transitionOther, "RolesWithPermissions", transitionOther.RolesWithPermissions, transition.RolesWithPermissions)
		diffs = append(diffs, ops)
	}
	GeneratedMessagesDifferent := false
	if len(transition.GeneratedMessages) != len(transitionOther.GeneratedMessages) {
		GeneratedMessagesDifferent = true
	} else {
		for i := range transition.GeneratedMessages {
			if (transition.GeneratedMessages[i] == nil) != (transitionOther.GeneratedMessages[i] == nil) {
				GeneratedMessagesDifferent = true
				break
			} else if transition.GeneratedMessages[i] != nil && transitionOther.GeneratedMessages[i] != nil {
				// this is a pointer comparaison
				if transition.GeneratedMessages[i] != transitionOther.GeneratedMessages[i] {
					GeneratedMessagesDifferent = true
					break
				}
			}
		}
	}
	if GeneratedMessagesDifferent {
		ops := Diff(stage, transition, transitionOther, "GeneratedMessages", transitionOther.GeneratedMessages, transition.GeneratedMessages)
		diffs = append(diffs, ops)
	}
	if (transition.Guard == nil) != (transitionOther.Guard == nil) {
		diffs = append(diffs, transition.GongMarshallField(stage, "Guard"))
	} else if transition.Guard != nil && transitionOther.Guard != nil {
		if transition.Guard != transitionOther.Guard {
			diffs = append(diffs, transition.GongMarshallField(stage, "Guard"))
		}
	}
	DiagramsDifferent := false
	if len(transition.Diagrams) != len(transitionOther.Diagrams) {
		DiagramsDifferent = true
	} else {
		for i := range transition.Diagrams {
			if (transition.Diagrams[i] == nil) != (transitionOther.Diagrams[i] == nil) {
				DiagramsDifferent = true
				break
			} else if transition.Diagrams[i] != nil && transitionOther.Diagrams[i] != nil {
				// this is a pointer comparaison
				if transition.Diagrams[i] != transitionOther.Diagrams[i] {
					DiagramsDifferent = true
					break
				}
			}
		}
	}
	if DiagramsDifferent {
		ops := Diff(stage, transition, transitionOther, "Diagrams", transitionOther.Diagrams, transition.Diagrams)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (transition_shape *Transition_Shape) GongDiff(stage *Stage, transition_shapeOther *Transition_Shape) (diffs []string) {
	// insertion point for field diffs
	if transition_shape.Name != transition_shapeOther.Name {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "Name"))
	}
	if (transition_shape.Transition == nil) != (transition_shapeOther.Transition == nil) {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "Transition"))
	} else if transition_shape.Transition != nil && transition_shapeOther.Transition != nil {
		if transition_shape.Transition != transition_shapeOther.Transition {
			diffs = append(diffs, transition_shape.GongMarshallField(stage, "Transition"))
		}
	}
	if transition_shape.StartRatio != transition_shapeOther.StartRatio {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "StartRatio"))
	}
	if transition_shape.EndRatio != transition_shapeOther.EndRatio {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "EndRatio"))
	}
	if transition_shape.StartOrientation != transition_shapeOther.StartOrientation {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "StartOrientation"))
	}
	if transition_shape.EndOrientation != transition_shapeOther.EndOrientation {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "EndOrientation"))
	}
	if transition_shape.CornerOffsetRatio != transition_shapeOther.CornerOffsetRatio {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if transition_shape.IsHidden != transition_shapeOther.IsHidden {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
