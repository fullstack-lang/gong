// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (action *Action) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Actions[action]

	return
}

func (stage *Stage) IsStagedAction(action *Action) (ok bool) {

	return action.GongIsStaged(stage)
}

func (activities *Activities) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Activitiess[activities]

	return
}

func (stage *Stage) IsStagedActivities(activities *Activities) (ok bool) {

	return activities.GongIsStaged(stage)
}

func (diagram *Diagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	return diagram.GongIsStaged(stage)
}

func (guard *Guard) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Guards[guard]

	return
}

func (stage *Stage) IsStagedGuard(guard *Guard) (ok bool) {

	return guard.GongIsStaged(stage)
}

func (kill *Kill) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Kills[kill]

	return
}

func (stage *Stage) IsStagedKill(kill *Kill) (ok bool) {

	return kill.GongIsStaged(stage)
}

func (library *Library) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	return library.GongIsStaged(stage)
}

func (message *Message) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Messages[message]

	return
}

func (stage *Stage) IsStagedMessage(message *Message) (ok bool) {

	return message.GongIsStaged(stage)
}

func (messagetype *MessageType) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.MessageTypes[messagetype]

	return
}

func (stage *Stage) IsStagedMessageType(messagetype *MessageType) (ok bool) {

	return messagetype.GongIsStaged(stage)
}

func (note *Note) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	return note.GongIsStaged(stage)
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	return noteshape.GongIsStaged(stage)
}

func (notestateshape *NoteStateShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteStateShapes[notestateshape]

	return
}

func (stage *Stage) IsStagedNoteStateShape(notestateshape *NoteStateShape) (ok bool) {

	return notestateshape.GongIsStaged(stage)
}

func (object *Object) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Objects[object]

	return
}

func (stage *Stage) IsStagedObject(object *Object) (ok bool) {

	return object.GongIsStaged(stage)
}

func (role *Role) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Roles[role]

	return
}

func (stage *Stage) IsStagedRole(role *Role) (ok bool) {

	return role.GongIsStaged(stage)
}

func (state *State) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.States[state]

	return
}

func (stage *Stage) IsStagedState(state *State) (ok bool) {

	return state.GongIsStaged(stage)
}

func (statemachine *StateMachine) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StateMachines[statemachine]

	return
}

func (stage *Stage) IsStagedStateMachine(statemachine *StateMachine) (ok bool) {

	return statemachine.GongIsStaged(stage)
}

func (stateshape *StateShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StateShapes[stateshape]

	return
}

func (stage *Stage) IsStagedStateShape(stateshape *StateShape) (ok bool) {

	return stateshape.GongIsStaged(stage)
}

func (transition *Transition) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Transitions[transition]

	return
}

func (stage *Stage) IsStagedTransition(transition *Transition) (ok bool) {

	return transition.GongIsStaged(stage)
}

func (transition_shape *Transition_Shape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Transition_Shapes[transition_shape]

	return
}

func (stage *Stage) IsStagedTransition_Shape(transition_shape *Transition_Shape) (ok bool) {

	return transition_shape.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (action *Action) GongStageBranch(stage *Stage) {
	stage.StageBranchAction(action)
}

func (stage *Stage) StageBranchAction(action *Action) {

	// check if instance is already staged
	if stage.IsStaged(action) {
		return
	}

	action.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (activities *Activities) GongStageBranch(stage *Stage) {
	stage.StageBranchActivities(activities)
}

func (stage *Stage) StageBranchActivities(activities *Activities) {

	// check if instance is already staged
	if stage.IsStaged(activities) {
		return
	}

	activities.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongStageBranch(stage *Stage) {
	stage.StageBranchDiagram(diagram)
}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if stage.IsStaged(diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stateshape := range diagram.State_Shapes {
		stage.StageBranch(_stateshape)
	}
	for _, _state := range diagram.StatesWhoseNodeIsExpanded {
		stage.StageBranch(_state)
	}
	for _, _transition_shape := range diagram.Transition_Shapes {
		stage.StageBranch(_transition_shape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		stage.StageBranch(_noteshape)
	}
	for _, _notestateshape := range diagram.NoteState_Shapes {
		stage.StageBranch(_notestateshape)
	}

}

func (guard *Guard) GongStageBranch(stage *Stage) {
	stage.StageBranchGuard(guard)
}

func (stage *Stage) StageBranchGuard(guard *Guard) {

	// check if instance is already staged
	if stage.IsStaged(guard) {
		return
	}

	guard.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (kill *Kill) GongStageBranch(stage *Stage) {
	stage.StageBranchKill(kill)
}

func (stage *Stage) StageBranchKill(kill *Kill) {

	// check if instance is already staged
	if stage.IsStaged(kill) {
		return
	}

	kill.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongStageBranch(stage *Stage) {
	stage.StageBranchLibrary(library)
}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if stage.IsStaged(library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.StageBranch(_library)
	}
	for _, _diagram := range library.Diagrams {
		stage.StageBranch(_diagram)
	}
	for _, _statemachine := range library.RootStateMachines {
		stage.StageBranch(_statemachine)
	}
	for _, _statemachine := range library.StateMachinesWhoseNodeIsExpanded {
		stage.StageBranch(_statemachine)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.StageBranch(_library)
	}
	for _, _role := range library.Roles {
		stage.StageBranch(_role)
	}

}

func (message *Message) GongStageBranch(stage *Stage) {
	stage.StageBranchMessage(message)
}

func (stage *Stage) StageBranchMessage(message *Message) {

	// check if instance is already staged
	if stage.IsStaged(message) {
		return
	}

	message.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if message.MessageType != nil {
		stage.StageBranch(message.MessageType)
	}
	if message.OriginTransition != nil {
		stage.StageBranch(message.OriginTransition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (messagetype *MessageType) GongStageBranch(stage *Stage) {
	stage.StageBranchMessageType(messagetype)
}

func (stage *Stage) StageBranchMessageType(messagetype *MessageType) {

	// check if instance is already staged
	if stage.IsStaged(messagetype) {
		return
	}

	messagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (note *Note) GongStageBranch(stage *Stage) {
	stage.StageBranchNote(note)
}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if note.State != nil {
		stage.StageBranch(note.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteshape *NoteShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteShape(noteshape)
}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if stage.IsStaged(noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.StageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notestateshape *NoteStateShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteStateShape(notestateshape)
}

func (stage *Stage) StageBranchNoteStateShape(notestateshape *NoteStateShape) {

	// check if instance is already staged
	if stage.IsStaged(notestateshape) {
		return
	}

	notestateshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notestateshape.Note != nil {
		stage.StageBranch(notestateshape.Note)
	}
	if notestateshape.State != nil {
		stage.StageBranch(notestateshape.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (object *Object) GongStageBranch(stage *Stage) {
	stage.StageBranchObject(object)
}

func (stage *Stage) StageBranchObject(object *Object) {

	// check if instance is already staged
	if stage.IsStaged(object) {
		return
	}

	object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if object.State != nil {
		stage.StageBranch(object.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _message := range object.Messages {
		stage.StageBranch(_message)
	}

}

func (role *Role) GongStageBranch(stage *Stage) {
	stage.StageBranchRole(role)
}

func (stage *Stage) StageBranchRole(role *Role) {

	// check if instance is already staged
	if stage.IsStaged(role) {
		return
	}

	role.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range role.RolesWithSamePermissions {
		stage.StageBranch(_role)
	}

}

func (state *State) GongStageBranch(stage *Stage) {
	stage.StageBranchState(state)
}

func (stage *Stage) StageBranchState(state *State) {

	// check if instance is already staged
	if stage.IsStaged(state) {
		return
	}

	state.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if state.Entry != nil {
		stage.StageBranch(state.Entry)
	}
	if state.Exit != nil {
		stage.StageBranch(state.Exit)
	}
	if state.Parent != nil {
		stage.StageBranch(state.Parent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range state.SubStates {
		stage.StageBranch(_state)
	}
	for _, _activities := range state.Activities {
		stage.StageBranch(_activities)
	}
	for _, _diagram := range state.Diagrams {
		stage.StageBranch(_diagram)
	}
	for _, _note := range state.Notes {
		stage.StageBranch(_note)
	}

}

func (statemachine *StateMachine) GongStageBranch(stage *Stage) {
	stage.StageBranchStateMachine(statemachine)
}

func (stage *Stage) StageBranchStateMachine(statemachine *StateMachine) {

	// check if instance is already staged
	if stage.IsStaged(statemachine) {
		return
	}

	statemachine.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if statemachine.InitialState != nil {
		stage.StageBranch(statemachine.InitialState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range statemachine.States {
		stage.StageBranch(_state)
	}
	for _, _diagram := range statemachine.Diagrams {
		stage.StageBranch(_diagram)
	}

}

func (stateshape *StateShape) GongStageBranch(stage *Stage) {
	stage.StageBranchStateShape(stateshape)
}

func (stage *Stage) StageBranchStateShape(stateshape *StateShape) {

	// check if instance is already staged
	if stage.IsStaged(stateshape) {
		return
	}

	stateshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stateshape.State != nil {
		stage.StageBranch(stateshape.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (transition *Transition) GongStageBranch(stage *Stage) {
	stage.StageBranchTransition(transition)
}

func (stage *Stage) StageBranchTransition(transition *Transition) {

	// check if instance is already staged
	if stage.IsStaged(transition) {
		return
	}

	transition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transition.Start != nil {
		stage.StageBranch(transition.Start)
	}
	if transition.End != nil {
		stage.StageBranch(transition.End)
	}
	if transition.Guard != nil {
		stage.StageBranch(transition.Guard)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range transition.RolesWithPermissions {
		stage.StageBranch(_role)
	}
	for _, _messagetype := range transition.GeneratedMessages {
		stage.StageBranch(_messagetype)
	}
	for _, _diagram := range transition.Diagrams {
		stage.StageBranch(_diagram)
	}

}

func (transition_shape *Transition_Shape) GongStageBranch(stage *Stage) {
	stage.StageBranchTransition_Shape(transition_shape)
}

func (stage *Stage) StageBranchTransition_Shape(transition_shape *Transition_Shape) {

	// check if instance is already staged
	if stage.IsStaged(transition_shape) {
		return
	}

	transition_shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transition_shape.Transition != nil {
		stage.StageBranch(transition_shape.Transition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Action:
		toT := GongCopyBranchAction(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Activities:
		toT := GongCopyBranchActivities(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := GongCopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Guard:
		toT := GongCopyBranchGuard(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Kill:
		toT := GongCopyBranchKill(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Message:
		toT := GongCopyBranchMessage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MessageType:
		toT := GongCopyBranchMessageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := GongCopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteStateShape:
		toT := GongCopyBranchNoteStateShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Object:
		toT := GongCopyBranchObject(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Role:
		toT := GongCopyBranchRole(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *State:
		toT := GongCopyBranchState(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StateMachine:
		toT := GongCopyBranchStateMachine(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StateShape:
		toT := GongCopyBranchStateShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Transition:
		toT := GongCopyBranchTransition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Transition_Shape:
		toT := GongCopyBranchTransition_Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAction(mapOrigCopy map[any]any, actionFrom *Action) (actionTo *Action) {

	// actionFrom has already been copied
	if _actionTo, ok := mapOrigCopy[actionFrom]; ok {
		actionTo = _actionTo.(*Action)
		return
	}

	actionTo = new(Action)
	mapOrigCopy[actionFrom] = actionTo
	actionFrom.GongCopyBasicFields(actionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchActivities(mapOrigCopy map[any]any, activitiesFrom *Activities) (activitiesTo *Activities) {

	// activitiesFrom has already been copied
	if _activitiesTo, ok := mapOrigCopy[activitiesFrom]; ok {
		activitiesTo = _activitiesTo.(*Activities)
		return
	}

	activitiesTo = new(Activities)
	mapOrigCopy[activitiesFrom] = activitiesTo
	activitiesFrom.GongCopyBasicFields(activitiesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.GongCopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stateshape := range diagramFrom.State_Shapes {
		diagramTo.State_Shapes = append(diagramTo.State_Shapes, GongCopyBranchStateShape(mapOrigCopy, _stateshape))
	}
	for _, _state := range diagramFrom.StatesWhoseNodeIsExpanded {
		diagramTo.StatesWhoseNodeIsExpanded = append(diagramTo.StatesWhoseNodeIsExpanded, GongCopyBranchState(mapOrigCopy, _state))
	}
	for _, _transition_shape := range diagramFrom.Transition_Shapes {
		diagramTo.Transition_Shapes = append(diagramTo.Transition_Shapes, GongCopyBranchTransition_Shape(mapOrigCopy, _transition_shape))
	}
	for _, _noteshape := range diagramFrom.Note_Shapes {
		diagramTo.Note_Shapes = append(diagramTo.Note_Shapes, GongCopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _notestateshape := range diagramFrom.NoteState_Shapes {
		diagramTo.NoteState_Shapes = append(diagramTo.NoteState_Shapes, GongCopyBranchNoteStateShape(mapOrigCopy, _notestateshape))
	}

	return
}

func GongCopyBranchGuard(mapOrigCopy map[any]any, guardFrom *Guard) (guardTo *Guard) {

	// guardFrom has already been copied
	if _guardTo, ok := mapOrigCopy[guardFrom]; ok {
		guardTo = _guardTo.(*Guard)
		return
	}

	guardTo = new(Guard)
	mapOrigCopy[guardFrom] = guardTo
	guardFrom.GongCopyBasicFields(guardTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKill(mapOrigCopy map[any]any, killFrom *Kill) (killTo *Kill) {

	// killFrom has already been copied
	if _killTo, ok := mapOrigCopy[killFrom]; ok {
		killTo = _killTo.(*Kill)
		return
	}

	killTo = new(Kill)
	mapOrigCopy[killFrom] = killTo
	killFrom.GongCopyBasicFields(killTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.GongCopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _diagram := range libraryFrom.Diagrams {
		libraryTo.Diagrams = append(libraryTo.Diagrams, GongCopyBranchDiagram(mapOrigCopy, _diagram))
	}
	for _, _statemachine := range libraryFrom.RootStateMachines {
		libraryTo.RootStateMachines = append(libraryTo.RootStateMachines, GongCopyBranchStateMachine(mapOrigCopy, _statemachine))
	}
	for _, _statemachine := range libraryFrom.StateMachinesWhoseNodeIsExpanded {
		libraryTo.StateMachinesWhoseNodeIsExpanded = append(libraryTo.StateMachinesWhoseNodeIsExpanded, GongCopyBranchStateMachine(mapOrigCopy, _statemachine))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _role := range libraryFrom.Roles {
		libraryTo.Roles = append(libraryTo.Roles, GongCopyBranchRole(mapOrigCopy, _role))
	}

	return
}

func GongCopyBranchMessage(mapOrigCopy map[any]any, messageFrom *Message) (messageTo *Message) {

	// messageFrom has already been copied
	if _messageTo, ok := mapOrigCopy[messageFrom]; ok {
		messageTo = _messageTo.(*Message)
		return
	}

	messageTo = new(Message)
	mapOrigCopy[messageFrom] = messageTo
	messageFrom.GongCopyBasicFields(messageTo)

	//insertion point for the staging of instances referenced by pointers
	if messageFrom.MessageType != nil {
		messageTo.MessageType = GongCopyBranchMessageType(mapOrigCopy, messageFrom.MessageType)
	}
	if messageFrom.OriginTransition != nil {
		messageTo.OriginTransition = GongCopyBranchTransition(mapOrigCopy, messageFrom.OriginTransition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMessageType(mapOrigCopy map[any]any, messagetypeFrom *MessageType) (messagetypeTo *MessageType) {

	// messagetypeFrom has already been copied
	if _messagetypeTo, ok := mapOrigCopy[messagetypeFrom]; ok {
		messagetypeTo = _messagetypeTo.(*MessageType)
		return
	}

	messagetypeTo = new(MessageType)
	mapOrigCopy[messagetypeFrom] = messagetypeTo
	messagetypeFrom.GongCopyBasicFields(messagetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers
	if noteFrom.State != nil {
		noteTo.State = GongCopyBranchState(mapOrigCopy, noteFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteStateShape(mapOrigCopy map[any]any, notestateshapeFrom *NoteStateShape) (notestateshapeTo *NoteStateShape) {

	// notestateshapeFrom has already been copied
	if _notestateshapeTo, ok := mapOrigCopy[notestateshapeFrom]; ok {
		notestateshapeTo = _notestateshapeTo.(*NoteStateShape)
		return
	}

	notestateshapeTo = new(NoteStateShape)
	mapOrigCopy[notestateshapeFrom] = notestateshapeTo
	notestateshapeFrom.GongCopyBasicFields(notestateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notestateshapeFrom.Note != nil {
		notestateshapeTo.Note = GongCopyBranchNote(mapOrigCopy, notestateshapeFrom.Note)
	}
	if notestateshapeFrom.State != nil {
		notestateshapeTo.State = GongCopyBranchState(mapOrigCopy, notestateshapeFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchObject(mapOrigCopy map[any]any, objectFrom *Object) (objectTo *Object) {

	// objectFrom has already been copied
	if _objectTo, ok := mapOrigCopy[objectFrom]; ok {
		objectTo = _objectTo.(*Object)
		return
	}

	objectTo = new(Object)
	mapOrigCopy[objectFrom] = objectTo
	objectFrom.GongCopyBasicFields(objectTo)

	//insertion point for the staging of instances referenced by pointers
	if objectFrom.State != nil {
		objectTo.State = GongCopyBranchState(mapOrigCopy, objectFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _message := range objectFrom.Messages {
		objectTo.Messages = append(objectTo.Messages, GongCopyBranchMessage(mapOrigCopy, _message))
	}

	return
}

func GongCopyBranchRole(mapOrigCopy map[any]any, roleFrom *Role) (roleTo *Role) {

	// roleFrom has already been copied
	if _roleTo, ok := mapOrigCopy[roleFrom]; ok {
		roleTo = _roleTo.(*Role)
		return
	}

	roleTo = new(Role)
	mapOrigCopy[roleFrom] = roleTo
	roleFrom.GongCopyBasicFields(roleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range roleFrom.RolesWithSamePermissions {
		roleTo.RolesWithSamePermissions = append(roleTo.RolesWithSamePermissions, GongCopyBranchRole(mapOrigCopy, _role))
	}

	return
}

func GongCopyBranchState(mapOrigCopy map[any]any, stateFrom *State) (stateTo *State) {

	// stateFrom has already been copied
	if _stateTo, ok := mapOrigCopy[stateFrom]; ok {
		stateTo = _stateTo.(*State)
		return
	}

	stateTo = new(State)
	mapOrigCopy[stateFrom] = stateTo
	stateFrom.GongCopyBasicFields(stateTo)

	//insertion point for the staging of instances referenced by pointers
	if stateFrom.Entry != nil {
		stateTo.Entry = GongCopyBranchAction(mapOrigCopy, stateFrom.Entry)
	}
	if stateFrom.Exit != nil {
		stateTo.Exit = GongCopyBranchAction(mapOrigCopy, stateFrom.Exit)
	}
	if stateFrom.Parent != nil {
		stateTo.Parent = GongCopyBranchState(mapOrigCopy, stateFrom.Parent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range stateFrom.SubStates {
		stateTo.SubStates = append(stateTo.SubStates, GongCopyBranchState(mapOrigCopy, _state))
	}
	for _, _activities := range stateFrom.Activities {
		stateTo.Activities = append(stateTo.Activities, GongCopyBranchActivities(mapOrigCopy, _activities))
	}
	for _, _diagram := range stateFrom.Diagrams {
		stateTo.Diagrams = append(stateTo.Diagrams, GongCopyBranchDiagram(mapOrigCopy, _diagram))
	}
	for _, _note := range stateFrom.Notes {
		stateTo.Notes = append(stateTo.Notes, GongCopyBranchNote(mapOrigCopy, _note))
	}

	return
}

func GongCopyBranchStateMachine(mapOrigCopy map[any]any, statemachineFrom *StateMachine) (statemachineTo *StateMachine) {

	// statemachineFrom has already been copied
	if _statemachineTo, ok := mapOrigCopy[statemachineFrom]; ok {
		statemachineTo = _statemachineTo.(*StateMachine)
		return
	}

	statemachineTo = new(StateMachine)
	mapOrigCopy[statemachineFrom] = statemachineTo
	statemachineFrom.GongCopyBasicFields(statemachineTo)

	//insertion point for the staging of instances referenced by pointers
	if statemachineFrom.InitialState != nil {
		statemachineTo.InitialState = GongCopyBranchState(mapOrigCopy, statemachineFrom.InitialState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range statemachineFrom.States {
		statemachineTo.States = append(statemachineTo.States, GongCopyBranchState(mapOrigCopy, _state))
	}
	for _, _diagram := range statemachineFrom.Diagrams {
		statemachineTo.Diagrams = append(statemachineTo.Diagrams, GongCopyBranchDiagram(mapOrigCopy, _diagram))
	}

	return
}

func GongCopyBranchStateShape(mapOrigCopy map[any]any, stateshapeFrom *StateShape) (stateshapeTo *StateShape) {

	// stateshapeFrom has already been copied
	if _stateshapeTo, ok := mapOrigCopy[stateshapeFrom]; ok {
		stateshapeTo = _stateshapeTo.(*StateShape)
		return
	}

	stateshapeTo = new(StateShape)
	mapOrigCopy[stateshapeFrom] = stateshapeTo
	stateshapeFrom.GongCopyBasicFields(stateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stateshapeFrom.State != nil {
		stateshapeTo.State = GongCopyBranchState(mapOrigCopy, stateshapeFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTransition(mapOrigCopy map[any]any, transitionFrom *Transition) (transitionTo *Transition) {

	// transitionFrom has already been copied
	if _transitionTo, ok := mapOrigCopy[transitionFrom]; ok {
		transitionTo = _transitionTo.(*Transition)
		return
	}

	transitionTo = new(Transition)
	mapOrigCopy[transitionFrom] = transitionTo
	transitionFrom.GongCopyBasicFields(transitionTo)

	//insertion point for the staging of instances referenced by pointers
	if transitionFrom.Start != nil {
		transitionTo.Start = GongCopyBranchState(mapOrigCopy, transitionFrom.Start)
	}
	if transitionFrom.End != nil {
		transitionTo.End = GongCopyBranchState(mapOrigCopy, transitionFrom.End)
	}
	if transitionFrom.Guard != nil {
		transitionTo.Guard = GongCopyBranchGuard(mapOrigCopy, transitionFrom.Guard)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range transitionFrom.RolesWithPermissions {
		transitionTo.RolesWithPermissions = append(transitionTo.RolesWithPermissions, GongCopyBranchRole(mapOrigCopy, _role))
	}
	for _, _messagetype := range transitionFrom.GeneratedMessages {
		transitionTo.GeneratedMessages = append(transitionTo.GeneratedMessages, GongCopyBranchMessageType(mapOrigCopy, _messagetype))
	}
	for _, _diagram := range transitionFrom.Diagrams {
		transitionTo.Diagrams = append(transitionTo.Diagrams, GongCopyBranchDiagram(mapOrigCopy, _diagram))
	}

	return
}

func GongCopyBranchTransition_Shape(mapOrigCopy map[any]any, transition_shapeFrom *Transition_Shape) (transition_shapeTo *Transition_Shape) {

	// transition_shapeFrom has already been copied
	if _transition_shapeTo, ok := mapOrigCopy[transition_shapeFrom]; ok {
		transition_shapeTo = _transition_shapeTo.(*Transition_Shape)
		return
	}

	transition_shapeTo = new(Transition_Shape)
	mapOrigCopy[transition_shapeFrom] = transition_shapeTo
	transition_shapeFrom.GongCopyBasicFields(transition_shapeTo)

	//insertion point for the staging of instances referenced by pointers
	if transition_shapeFrom.Transition != nil {
		transition_shapeTo.Transition = GongCopyBranchTransition(mapOrigCopy, transition_shapeFrom.Transition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (action *Action) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAction(action)
}

func (stage *Stage) UnstageBranchAction(action *Action) {

	// check if instance is already staged
	if !stage.IsStaged(action) {
		return
	}

	action.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (activities *Activities) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchActivities(activities)
}

func (stage *Stage) UnstageBranchActivities(activities *Activities) {

	// check if instance is already staged
	if !stage.IsStaged(activities) {
		return
	}

	activities.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDiagram(diagram)
}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !stage.IsStaged(diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stateshape := range diagram.State_Shapes {
		stage.UnstageBranch(_stateshape)
	}
	for _, _state := range diagram.StatesWhoseNodeIsExpanded {
		stage.UnstageBranch(_state)
	}
	for _, _transition_shape := range diagram.Transition_Shapes {
		stage.UnstageBranch(_transition_shape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		stage.UnstageBranch(_noteshape)
	}
	for _, _notestateshape := range diagram.NoteState_Shapes {
		stage.UnstageBranch(_notestateshape)
	}

}

func (guard *Guard) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGuard(guard)
}

func (stage *Stage) UnstageBranchGuard(guard *Guard) {

	// check if instance is already staged
	if !stage.IsStaged(guard) {
		return
	}

	guard.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (kill *Kill) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchKill(kill)
}

func (stage *Stage) UnstageBranchKill(kill *Kill) {

	// check if instance is already staged
	if !stage.IsStaged(kill) {
		return
	}

	kill.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLibrary(library)
}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !stage.IsStaged(library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.UnstageBranch(_library)
	}
	for _, _diagram := range library.Diagrams {
		stage.UnstageBranch(_diagram)
	}
	for _, _statemachine := range library.RootStateMachines {
		stage.UnstageBranch(_statemachine)
	}
	for _, _statemachine := range library.StateMachinesWhoseNodeIsExpanded {
		stage.UnstageBranch(_statemachine)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.UnstageBranch(_library)
	}
	for _, _role := range library.Roles {
		stage.UnstageBranch(_role)
	}

}

func (message *Message) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMessage(message)
}

func (stage *Stage) UnstageBranchMessage(message *Message) {

	// check if instance is already staged
	if !stage.IsStaged(message) {
		return
	}

	message.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if message.MessageType != nil {
		stage.UnstageBranch(message.MessageType)
	}
	if message.OriginTransition != nil {
		stage.UnstageBranch(message.OriginTransition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (messagetype *MessageType) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMessageType(messagetype)
}

func (stage *Stage) UnstageBranchMessageType(messagetype *MessageType) {

	// check if instance is already staged
	if !stage.IsStaged(messagetype) {
		return
	}

	messagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (note *Note) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNote(note)
}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if note.State != nil {
		stage.UnstageBranch(note.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteshape *NoteShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteShape(noteshape)
}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !stage.IsStaged(noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.UnstageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notestateshape *NoteStateShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteStateShape(notestateshape)
}

func (stage *Stage) UnstageBranchNoteStateShape(notestateshape *NoteStateShape) {

	// check if instance is already staged
	if !stage.IsStaged(notestateshape) {
		return
	}

	notestateshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notestateshape.Note != nil {
		stage.UnstageBranch(notestateshape.Note)
	}
	if notestateshape.State != nil {
		stage.UnstageBranch(notestateshape.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (object *Object) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchObject(object)
}

func (stage *Stage) UnstageBranchObject(object *Object) {

	// check if instance is already staged
	if !stage.IsStaged(object) {
		return
	}

	object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if object.State != nil {
		stage.UnstageBranch(object.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _message := range object.Messages {
		stage.UnstageBranch(_message)
	}

}

func (role *Role) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRole(role)
}

func (stage *Stage) UnstageBranchRole(role *Role) {

	// check if instance is already staged
	if !stage.IsStaged(role) {
		return
	}

	role.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range role.RolesWithSamePermissions {
		stage.UnstageBranch(_role)
	}

}

func (state *State) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchState(state)
}

func (stage *Stage) UnstageBranchState(state *State) {

	// check if instance is already staged
	if !stage.IsStaged(state) {
		return
	}

	state.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if state.Entry != nil {
		stage.UnstageBranch(state.Entry)
	}
	if state.Exit != nil {
		stage.UnstageBranch(state.Exit)
	}
	if state.Parent != nil {
		stage.UnstageBranch(state.Parent)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range state.SubStates {
		stage.UnstageBranch(_state)
	}
	for _, _activities := range state.Activities {
		stage.UnstageBranch(_activities)
	}
	for _, _diagram := range state.Diagrams {
		stage.UnstageBranch(_diagram)
	}
	for _, _note := range state.Notes {
		stage.UnstageBranch(_note)
	}

}

func (statemachine *StateMachine) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStateMachine(statemachine)
}

func (stage *Stage) UnstageBranchStateMachine(statemachine *StateMachine) {

	// check if instance is already staged
	if !stage.IsStaged(statemachine) {
		return
	}

	statemachine.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if statemachine.InitialState != nil {
		stage.UnstageBranch(statemachine.InitialState)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _state := range statemachine.States {
		stage.UnstageBranch(_state)
	}
	for _, _diagram := range statemachine.Diagrams {
		stage.UnstageBranch(_diagram)
	}

}

func (stateshape *StateShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStateShape(stateshape)
}

func (stage *Stage) UnstageBranchStateShape(stateshape *StateShape) {

	// check if instance is already staged
	if !stage.IsStaged(stateshape) {
		return
	}

	stateshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if stateshape.State != nil {
		stage.UnstageBranch(stateshape.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (transition *Transition) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTransition(transition)
}

func (stage *Stage) UnstageBranchTransition(transition *Transition) {

	// check if instance is already staged
	if !stage.IsStaged(transition) {
		return
	}

	transition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transition.Start != nil {
		stage.UnstageBranch(transition.Start)
	}
	if transition.End != nil {
		stage.UnstageBranch(transition.End)
	}
	if transition.Guard != nil {
		stage.UnstageBranch(transition.Guard)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range transition.RolesWithPermissions {
		stage.UnstageBranch(_role)
	}
	for _, _messagetype := range transition.GeneratedMessages {
		stage.UnstageBranch(_messagetype)
	}
	for _, _diagram := range transition.Diagrams {
		stage.UnstageBranch(_diagram)
	}

}

func (transition_shape *Transition_Shape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTransition_Shape(transition_shape)
}

func (stage *Stage) UnstageBranchTransition_Shape(transition_shape *Transition_Shape) {

	// check if instance is already staged
	if !stage.IsStaged(transition_shape) {
		return
	}

	transition_shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if transition_shape.Transition != nil {
		stage.UnstageBranch(transition_shape.Transition)
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

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
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
	reference.SubLibrariesWhoseNodeIsExpanded = reference.SubLibrariesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SubLibrariesWhoseNodeIsExpanded {
		reference.SubLibrariesWhoseNodeIsExpanded = append(reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference[_b])
	}
	reference.Roles = reference.Roles[:0]
	for _, _b := range instance.Roles {
		reference.Roles = append(reference.Roles, stage.Roles_reference[_b])
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
	if instance.State != nil {
		reference.State = stage.States_reference[instance.State]
	}
	// insertion point for slice of pointers field
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
	if instance.Entry != nil {
		reference.Entry = stage.Actions_reference[instance.Entry]
	}
	if instance.Exit != nil {
		reference.Exit = stage.Actions_reference[instance.Exit]
	}
	if instance.Parent != nil {
		reference.Parent = stage.States_reference[instance.Parent]
	}
	// insertion point for slice of pointers field
	reference.SubStates = reference.SubStates[:0]
	for _, _b := range instance.SubStates {
		reference.SubStates = append(reference.SubStates, stage.States_reference[_b])
	}
	reference.Activities = reference.Activities[:0]
	for _, _b := range instance.Activities {
		reference.Activities = append(reference.Activities, stage.Activitiess_reference[_b])
	}
	reference.Diagrams = reference.Diagrams[:0]
	for _, _b := range instance.Diagrams {
		reference.Diagrams = append(reference.Diagrams, stage.Diagrams_reference[_b])
	}
	reference.Notes = reference.Notes[:0]
	for _, _b := range instance.Notes {
		reference.Notes = append(reference.Notes, stage.Notes_reference[_b])
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

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
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
	var _SubLibrariesWhoseNodeIsExpanded []*Library
	for _, _reference := range reference.SubLibrariesWhoseNodeIsExpanded {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibrariesWhoseNodeIsExpanded = append(_SubLibrariesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SubLibrariesWhoseNodeIsExpanded = _SubLibrariesWhoseNodeIsExpanded
	var _Roles []*Role
	for _, _reference := range reference.Roles {
		if _instance, ok := stage.Roles_instance[_reference]; ok {
			_Roles = append(_Roles, _instance)
		}
	}
	reference.Roles = _Roles
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
	if _reference := reference.State; _reference != nil {
		reference.State = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.State = _instance
		}
	}
	// insertion point for slice of pointers fields
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
	if _reference := reference.Parent; _reference != nil {
		reference.Parent = nil
		if _instance, ok := stage.States_instance[_reference]; ok {
			reference.Parent = _instance
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
	var _Activities []*Activities
	for _, _reference := range reference.Activities {
		if _instance, ok := stage.Activitiess_instance[_reference]; ok {
			_Activities = append(_Activities, _instance)
		}
	}
	reference.Activities = _Activities
	var _Diagrams []*Diagram
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
	var _Notes []*Note
	for _, _reference := range reference.Notes {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_Notes = append(_Notes, _instance)
		}
	}
	reference.Notes = _Notes
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
		ops := stage.Diff(
			diagram,
			"State_Shapes",
			len(diagramOther.State_Shapes),
			len(diagram.State_Shapes),
			func(i, j int) bool {
				return diagramOther.State_Shapes[i] == diagram.State_Shapes[j]
			},
			func(j int) string {
				return diagram.State_Shapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagram,
			"StatesWhoseNodeIsExpanded",
			len(diagramOther.StatesWhoseNodeIsExpanded),
			len(diagram.StatesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.StatesWhoseNodeIsExpanded[i] == diagram.StatesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.StatesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagram,
			"Transition_Shapes",
			len(diagramOther.Transition_Shapes),
			len(diagram.Transition_Shapes),
			func(i, j int) bool {
				return diagramOther.Transition_Shapes[i] == diagram.Transition_Shapes[j]
			},
			func(j int) string {
				return diagram.Transition_Shapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagram,
			"Note_Shapes",
			len(diagramOther.Note_Shapes),
			len(diagram.Note_Shapes),
			func(i, j int) bool {
				return diagramOther.Note_Shapes[i] == diagram.Note_Shapes[j]
			},
			func(j int) string {
				return diagram.Note_Shapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagram,
			"NoteState_Shapes",
			len(diagramOther.NoteState_Shapes),
			len(diagram.NoteState_Shapes),
			func(i, j int) bool {
				return diagramOther.NoteState_Shapes[i] == diagram.NoteState_Shapes[j]
			},
			func(j int) string {
				return diagram.NoteState_Shapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			library,
			"SubLibraries",
			len(libraryOther.SubLibraries),
			len(library.SubLibraries),
			func(i, j int) bool {
				return libraryOther.SubLibraries[i] == library.SubLibraries[j]
			},
			func(j int) string {
				return library.SubLibraries[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			library,
			"Diagrams",
			len(libraryOther.Diagrams),
			len(library.Diagrams),
			func(i, j int) bool {
				return libraryOther.Diagrams[i] == library.Diagrams[j]
			},
			func(j int) string {
				return library.Diagrams[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			library,
			"RootStateMachines",
			len(libraryOther.RootStateMachines),
			len(library.RootStateMachines),
			func(i, j int) bool {
				return libraryOther.RootStateMachines[i] == library.RootStateMachines[j]
			},
			func(j int) string {
				return library.RootStateMachines[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			library,
			"StateMachinesWhoseNodeIsExpanded",
			len(libraryOther.StateMachinesWhoseNodeIsExpanded),
			len(library.StateMachinesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.StateMachinesWhoseNodeIsExpanded[i] == library.StateMachinesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.StateMachinesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			library,
			"SubLibrariesWhoseNodeIsExpanded",
			len(libraryOther.SubLibrariesWhoseNodeIsExpanded),
			len(library.SubLibrariesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.SubLibrariesWhoseNodeIsExpanded[i] == library.SubLibrariesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.SubLibrariesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
	}
	RolesDifferent := false
	if len(library.Roles) != len(libraryOther.Roles) {
		RolesDifferent = true
	} else {
		for i := range library.Roles {
			if (library.Roles[i] == nil) != (libraryOther.Roles[i] == nil) {
				RolesDifferent = true
				break
			} else if library.Roles[i] != nil && libraryOther.Roles[i] != nil {
				// this is a pointer comparaison
				if library.Roles[i] != libraryOther.Roles[i] {
					RolesDifferent = true
					break
				}
			}
		}
	}
	if RolesDifferent {
		ops := stage.Diff(
			library,
			"Roles",
			len(libraryOther.Roles),
			len(library.Roles),
			func(i, j int) bool {
				return libraryOther.Roles[i] == library.Roles[j]
			},
			func(j int) string {
				return library.Roles[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
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
	if (note.State == nil) != (noteOther.State == nil) {
		diffs = append(diffs, note.GongMarshallField(stage, "State"))
	} else if note.State != nil && noteOther.State != nil {
		if note.State != noteOther.State {
			diffs = append(diffs, note.GongMarshallField(stage, "State"))
		}
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
		ops := stage.Diff(
			object,
			"Messages",
			len(objectOther.Messages),
			len(object.Messages),
			func(i, j int) bool {
				return objectOther.Messages[i] == object.Messages[j]
			},
			func(j int) string {
				return object.Messages[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			role,
			"RolesWithSamePermissions",
			len(roleOther.RolesWithSamePermissions),
			len(role.RolesWithSamePermissions),
			func(i, j int) bool {
				return roleOther.RolesWithSamePermissions[i] == role.RolesWithSamePermissions[j]
			},
			func(j int) string {
				return role.RolesWithSamePermissions[j].GongGetIdentifier(stage)
			},
		)
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
	if state.IsEndState != stateOther.IsEndState {
		diffs = append(diffs, state.GongMarshallField(stage, "IsEndState"))
	}
	if state.IsDecisionNode != stateOther.IsDecisionNode {
		diffs = append(diffs, state.GongMarshallField(stage, "IsDecisionNode"))
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
		ops := stage.Diff(
			state,
			"SubStates",
			len(stateOther.SubStates),
			len(state.SubStates),
			func(i, j int) bool {
				return stateOther.SubStates[i] == state.SubStates[j]
			},
			func(j int) string {
				return state.SubStates[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			state,
			"Activities",
			len(stateOther.Activities),
			len(state.Activities),
			func(i, j int) bool {
				return stateOther.Activities[i] == state.Activities[j]
			},
			func(j int) string {
				return state.Activities[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if (state.Exit == nil) != (stateOther.Exit == nil) {
		diffs = append(diffs, state.GongMarshallField(stage, "Exit"))
	} else if state.Exit != nil && stateOther.Exit != nil {
		if state.Exit != stateOther.Exit {
			diffs = append(diffs, state.GongMarshallField(stage, "Exit"))
		}
	}
	if (state.Parent == nil) != (stateOther.Parent == nil) {
		diffs = append(diffs, state.GongMarshallField(stage, "Parent"))
	} else if state.Parent != nil && stateOther.Parent != nil {
		if state.Parent != stateOther.Parent {
			diffs = append(diffs, state.GongMarshallField(stage, "Parent"))
		}
	}
	if state.IsFictious != stateOther.IsFictious {
		diffs = append(diffs, state.GongMarshallField(stage, "IsFictious"))
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
		ops := stage.Diff(
			state,
			"Diagrams",
			len(stateOther.Diagrams),
			len(state.Diagrams),
			func(i, j int) bool {
				return stateOther.Diagrams[i] == state.Diagrams[j]
			},
			func(j int) string {
				return state.Diagrams[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	NotesDifferent := false
	if len(state.Notes) != len(stateOther.Notes) {
		NotesDifferent = true
	} else {
		for i := range state.Notes {
			if (state.Notes[i] == nil) != (stateOther.Notes[i] == nil) {
				NotesDifferent = true
				break
			} else if state.Notes[i] != nil && stateOther.Notes[i] != nil {
				// this is a pointer comparaison
				if state.Notes[i] != stateOther.Notes[i] {
					NotesDifferent = true
					break
				}
			}
		}
	}
	if NotesDifferent {
		ops := stage.Diff(
			state,
			"Notes",
			len(stateOther.Notes),
			len(state.Notes),
			func(i, j int) bool {
				return stateOther.Notes[i] == state.Notes[j]
			},
			func(j int) string {
				return state.Notes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
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
	if (statemachine.InitialState == nil) != (statemachineOther.InitialState == nil) {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "InitialState"))
	} else if statemachine.InitialState != nil && statemachineOther.InitialState != nil {
		if statemachine.InitialState != statemachineOther.InitialState {
			diffs = append(diffs, statemachine.GongMarshallField(stage, "InitialState"))
		}
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
		ops := stage.Diff(
			statemachine,
			"States",
			len(statemachineOther.States),
			len(statemachine.States),
			func(i, j int) bool {
				return statemachineOther.States[i] == statemachine.States[j]
			},
			func(j int) string {
				return statemachine.States[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			statemachine,
			"Diagrams",
			len(statemachineOther.Diagrams),
			len(statemachine.Diagrams),
			func(i, j int) bool {
				return statemachineOther.Diagrams[i] == statemachine.Diagrams[j]
			},
			func(j int) string {
				return statemachine.Diagrams[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if statemachine.IsWithTransitionNameAutonamticalyGenerated != statemachineOther.IsWithTransitionNameAutonamticalyGenerated {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "IsWithTransitionNameAutonamticalyGenerated"))
	}
	if statemachine.ComputedPrefix != statemachineOther.ComputedPrefix {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "ComputedPrefix"))
	}
	if statemachine.IsExpanded != statemachineOther.IsExpanded {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "IsExpanded"))
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
		ops := stage.Diff(
			transition,
			"RolesWithPermissions",
			len(transitionOther.RolesWithPermissions),
			len(transition.RolesWithPermissions),
			func(i, j int) bool {
				return transitionOther.RolesWithPermissions[i] == transition.RolesWithPermissions[j]
			},
			func(j int) string {
				return transition.RolesWithPermissions[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			transition,
			"GeneratedMessages",
			len(transitionOther.GeneratedMessages),
			len(transition.GeneratedMessages),
			func(i, j int) bool {
				return transitionOther.GeneratedMessages[i] == transition.GeneratedMessages[j]
			},
			func(j int) string {
				return transition.GeneratedMessages[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			transition,
			"Diagrams",
			len(transitionOther.Diagrams),
			len(transition.Diagrams),
			func(i, j int) bool {
				return transitionOther.Diagrams[i] == transition.Diagrams[j]
			},
			func(j int) string {
				return transition.Diagrams[j].GongGetIdentifier(stage)
			},
		)
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

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
