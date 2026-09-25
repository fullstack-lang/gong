// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (action *Action) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Actions[action]
	return ok
}

func (activities *Activities) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Activitiess[activities]
	return ok
}

func (diagram *Diagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Diagrams[diagram]
	return ok
}

func (guard *Guard) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Guards[guard]
	return ok
}

func (kill *Kill) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Kills[kill]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (message *Message) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Messages[message]
	return ok
}

func (messagetype *MessageType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MessageTypes[messagetype]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteShapes[noteshape]
	return ok
}

func (notestateshape *NoteStateShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteStateShapes[notestateshape]
	return ok
}

func (object *Object) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Objects[object]
	return ok
}

func (role *Role) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Roles[role]
	return ok
}

func (state *State) GongIsStaged(stage *Stage) bool {
	_, ok := stage.States[state]
	return ok
}

func (statemachine *StateMachine) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StateMachines[statemachine]
	return ok
}

func (stateshape *StateShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StateShapes[stateshape]
	return ok
}

func (transition *Transition) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Transitions[transition]
	return ok
}

func (transition_shape *Transition_Shape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Transition_Shapes[transition_shape]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (action *Action) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(action) {
		return
	}

	action.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (activities *Activities) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(activities) {
		return
	}

	activities.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(guard) {
		return
	}

	guard.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (kill *Kill) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(kill) {
		return
	}

	kill.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongStageBranch(stage *Stage) {

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
	for _, _messagetype := range library.MessageTypes {
		stage.StageBranch(_messagetype)
	}

}

func (message *Message) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(messagetype) {
		return
	}

	messagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (note *Note) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	actionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, actionFrom)
	if alreadyCopied {
		return
	}
	actionFrom.GongCopyBasicFields(actionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchActivities(mapOrigCopy map[any]any, activitiesFrom *Activities) (activitiesTo *Activities) {
	var alreadyCopied bool
	activitiesTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, activitiesFrom)
	if alreadyCopied {
		return
	}
	activitiesFrom.GongCopyBasicFields(activitiesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {
	var alreadyCopied bool
	diagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	guardTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, guardFrom)
	if alreadyCopied {
		return
	}
	guardFrom.GongCopyBasicFields(guardTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchKill(mapOrigCopy map[any]any, killFrom *Kill) (killTo *Kill) {
	var alreadyCopied bool
	killTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, killFrom)
	if alreadyCopied {
		return
	}
	killFrom.GongCopyBasicFields(killTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {
	var alreadyCopied bool
	libraryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, libraryFrom)
	if alreadyCopied {
		return
	}
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
	for _, _messagetype := range libraryFrom.MessageTypes {
		libraryTo.MessageTypes = append(libraryTo.MessageTypes, GongCopyBranchMessageType(mapOrigCopy, _messagetype))
	}

	return
}

func GongCopyBranchMessage(mapOrigCopy map[any]any, messageFrom *Message) (messageTo *Message) {
	var alreadyCopied bool
	messageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, messageFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	messagetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, messagetypeFrom)
	if alreadyCopied {
		return
	}
	messagetypeFrom.GongCopyBasicFields(messagetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers
	if noteFrom.State != nil {
		noteTo.State = GongCopyBranchState(mapOrigCopy, noteFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {
	var alreadyCopied bool
	noteshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteshapeFrom)
	if alreadyCopied {
		return
	}
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteStateShape(mapOrigCopy map[any]any, notestateshapeFrom *NoteStateShape) (notestateshapeTo *NoteStateShape) {
	var alreadyCopied bool
	notestateshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notestateshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	objectTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, objectFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	roleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, roleFrom)
	if alreadyCopied {
		return
	}
	roleFrom.GongCopyBasicFields(roleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _role := range roleFrom.RolesWithSamePermissions {
		roleTo.RolesWithSamePermissions = append(roleTo.RolesWithSamePermissions, GongCopyBranchRole(mapOrigCopy, _role))
	}

	return
}

func GongCopyBranchState(mapOrigCopy map[any]any, stateFrom *State) (stateTo *State) {
	var alreadyCopied bool
	stateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stateFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	statemachineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, statemachineFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	stateshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stateshapeFrom)
	if alreadyCopied {
		return
	}
	stateshapeFrom.GongCopyBasicFields(stateshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if stateshapeFrom.State != nil {
		stateshapeTo.State = GongCopyBranchState(mapOrigCopy, stateshapeFrom.State)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTransition(mapOrigCopy map[any]any, transitionFrom *Transition) (transitionTo *Transition) {
	var alreadyCopied bool
	transitionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, transitionFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	transition_shapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, transition_shapeFrom)
	if alreadyCopied {
		return
	}
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

// insertion point for unstage branch per struct
func (action *Action) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(action) {
		return
	}

	action.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (activities *Activities) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(activities) {
		return
	}

	activities.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(guard) {
		return
	}

	guard.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (kill *Kill) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(kill) {
		return
	}

	kill.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongUnstageBranch(stage *Stage) {

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
	for _, _messagetype := range library.MessageTypes {
		stage.UnstageBranch(_messagetype)
	}

}

func (message *Message) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(messagetype) {
		return
	}

	messagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (note *Note) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructSliceOfPointersFromReferences(&reference.State_Shapes, stage.StateShapes_reference, instance.State_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.StatesWhoseNodeIsExpanded, stage.States_reference, instance.StatesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Transition_Shapes, stage.Transition_Shapes_reference, instance.Transition_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note_Shapes, stage.NoteShapes_reference, instance.Note_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteState_Shapes, stage.NoteStateShapes_reference, instance.NoteState_Shapes)
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
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Diagrams, stage.Diagrams_reference, instance.Diagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootStateMachines, stage.StateMachines_reference, instance.RootStateMachines)
	__gong__reconstructSliceOfPointersFromReferences(&reference.StateMachinesWhoseNodeIsExpanded, stage.StateMachines_reference, instance.StateMachinesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference, instance.SubLibrariesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Roles, stage.Roles_reference, instance.Roles)
	__gong__reconstructSliceOfPointersFromReferences(&reference.MessageTypes, stage.MessageTypes_reference, instance.MessageTypes)
}

func (reference *Message) GongReconstructPointersFromReferences(stage *Stage, instance *Message) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MessageType, stage.MessageTypes_reference, instance.MessageType)
	__gong__reconstructPointer(&reference.OriginTransition, stage.Transitions_reference, instance.OriginTransition)
	// insertion point for slice of pointers field
}

func (reference *MessageType) GongReconstructPointersFromReferences(stage *Stage, instance *MessageType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.State, stage.States_reference, instance.State)
	// insertion point for slice of pointers field
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	// insertion point for slice of pointers field
}

func (reference *NoteStateShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteStateShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.State, stage.States_reference, instance.State)
	// insertion point for slice of pointers field
}

func (reference *Object) GongReconstructPointersFromReferences(stage *Stage, instance *Object) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.State, stage.States_reference, instance.State)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Messages, stage.Messages_reference, instance.Messages)
}

func (reference *Role) GongReconstructPointersFromReferences(stage *Stage, instance *Role) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.RolesWithSamePermissions, stage.Roles_reference, instance.RolesWithSamePermissions)
}

func (reference *State) GongReconstructPointersFromReferences(stage *Stage, instance *State) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Entry, stage.Actions_reference, instance.Entry)
	__gong__reconstructPointer(&reference.Exit, stage.Actions_reference, instance.Exit)
	__gong__reconstructPointer(&reference.Parent, stage.States_reference, instance.Parent)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubStates, stage.States_reference, instance.SubStates)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Activities, stage.Activitiess_reference, instance.Activities)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Diagrams, stage.Diagrams_reference, instance.Diagrams)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Notes, stage.Notes_reference, instance.Notes)
}

func (reference *StateMachine) GongReconstructPointersFromReferences(stage *Stage, instance *StateMachine) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.InitialState, stage.States_reference, instance.InitialState)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.States, stage.States_reference, instance.States)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Diagrams, stage.Diagrams_reference, instance.Diagrams)
}

func (reference *StateShape) GongReconstructPointersFromReferences(stage *Stage, instance *StateShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.State, stage.States_reference, instance.State)
	// insertion point for slice of pointers field
}

func (reference *Transition) GongReconstructPointersFromReferences(stage *Stage, instance *Transition) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Start, stage.States_reference, instance.Start)
	__gong__reconstructPointer(&reference.End, stage.States_reference, instance.End)
	__gong__reconstructPointer(&reference.Guard, stage.Guards_reference, instance.Guard)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.RolesWithPermissions, stage.Roles_reference, instance.RolesWithPermissions)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GeneratedMessages, stage.MessageTypes_reference, instance.GeneratedMessages)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Diagrams, stage.Diagrams_reference, instance.Diagrams)
}

func (reference *Transition_Shape) GongReconstructPointersFromReferences(stage *Stage, instance *Transition_Shape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Transition, stage.Transitions_reference, instance.Transition)
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.State_Shapes, stage.StateShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.StatesWhoseNodeIsExpanded, stage.States_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Transition_Shapes, stage.Transition_Shapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note_Shapes, stage.NoteShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteState_Shapes, stage.NoteStateShapes_instance)
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Diagrams, stage.Diagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootStateMachines, stage.StateMachines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.StateMachinesWhoseNodeIsExpanded, stage.StateMachines_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Roles, stage.Roles_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.MessageTypes, stage.MessageTypes_instance)
}

func (reference *Message) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MessageType, stage.MessageTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.OriginTransition, stage.Transitions_instance)
	// insertion point for slice of pointers fields
}

func (reference *MessageType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.State, stage.States_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteStateShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.State, stage.States_instance)
	// insertion point for slice of pointers fields
}

func (reference *Object) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.State, stage.States_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Messages, stage.Messages_instance)
}

func (reference *Role) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.RolesWithSamePermissions, stage.Roles_instance)
}

func (reference *State) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Entry, stage.Actions_instance)
	__gong__reconstructPointerFromInstance(&reference.Exit, stage.Actions_instance)
	__gong__reconstructPointerFromInstance(&reference.Parent, stage.States_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubStates, stage.States_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Activities, stage.Activitiess_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Diagrams, stage.Diagrams_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Notes, stage.Notes_instance)
}

func (reference *StateMachine) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.InitialState, stage.States_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.States, stage.States_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Diagrams, stage.Diagrams_instance)
}

func (reference *StateShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.State, stage.States_instance)
	// insertion point for slice of pointers fields
}

func (reference *Transition) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Start, stage.States_instance)
	__gong__reconstructPointerFromInstance(&reference.End, stage.States_instance)
	__gong__reconstructPointerFromInstance(&reference.Guard, stage.Guards_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.RolesWithPermissions, stage.Roles_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GeneratedMessages, stage.MessageTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Diagrams, stage.Diagrams_instance)
}

func (reference *Transition_Shape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Transition, stage.Transitions_instance)
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
	if ops := __gong__diffSliceOfPointers(stage, diagram, "State_Shapes", diagramOther.State_Shapes, diagram.State_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "StatesWhoseNodeIsExpanded", diagramOther.StatesWhoseNodeIsExpanded, diagram.StatesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Transition_Shapes", diagramOther.Transition_Shapes, diagram.Transition_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Note_Shapes", diagramOther.Note_Shapes, diagram.Note_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NoteState_Shapes", diagramOther.NoteState_Shapes, diagram.NoteState_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.ShowRoles != diagramOther.ShowRoles {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ShowRoles"))
	}
	if diagram.ShowMessages != diagramOther.ShowMessages {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ShowMessages"))
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
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, library, "Diagrams", libraryOther.Diagrams, library.Diagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootStateMachines", libraryOther.RootStateMachines, library.RootStateMachines); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsStateMachinesNodeExpanded != libraryOther.IsStateMachinesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsStateMachinesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "StateMachinesWhoseNodeIsExpanded", libraryOther.StateMachinesWhoseNodeIsExpanded, library.StateMachinesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "Roles", libraryOther.Roles, library.Roles); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsRolesNodeExpanded != libraryOther.IsRolesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRolesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "MessageTypes", libraryOther.MessageTypes, library.MessageTypes); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsMessageTypesNodeExpanded != libraryOther.IsMessageTypesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsMessageTypesNodeExpanded"))
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
	if message.MessageType != messageOther.MessageType {
		diffs = append(diffs, message.GongMarshallField(stage, "MessageType"))
	}
	if message.OriginTransition != messageOther.OriginTransition {
		diffs = append(diffs, message.GongMarshallField(stage, "OriginTransition"))
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
	if note.State != noteOther.State {
		diffs = append(diffs, note.GongMarshallField(stage, "State"))
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
	if noteshape.Note != noteshapeOther.Note {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
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
	if notestateshape.Note != notestateshapeOther.Note {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "Note"))
	}
	if notestateshape.State != notestateshapeOther.State {
		diffs = append(diffs, notestateshape.GongMarshallField(stage, "State"))
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
	if object.State != objectOther.State {
		diffs = append(diffs, object.GongMarshallField(stage, "State"))
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
	if ops := __gong__diffSliceOfPointers(stage, object, "Messages", objectOther.Messages, object.Messages); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, role, "RolesWithSamePermissions", roleOther.RolesWithSamePermissions, role.RolesWithSamePermissions); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, state, "SubStates", stateOther.SubStates, state.SubStates); ops != "" {
		diffs = append(diffs, ops)
	}
	if state.Entry != stateOther.Entry {
		diffs = append(diffs, state.GongMarshallField(stage, "Entry"))
	}
	if ops := __gong__diffSliceOfPointers(stage, state, "Activities", stateOther.Activities, state.Activities); ops != "" {
		diffs = append(diffs, ops)
	}
	if state.Exit != stateOther.Exit {
		diffs = append(diffs, state.GongMarshallField(stage, "Exit"))
	}
	if state.Parent != stateOther.Parent {
		diffs = append(diffs, state.GongMarshallField(stage, "Parent"))
	}
	if state.IsFictious != stateOther.IsFictious {
		diffs = append(diffs, state.GongMarshallField(stage, "IsFictious"))
	}
	if ops := __gong__diffSliceOfPointers(stage, state, "Diagrams", stateOther.Diagrams, state.Diagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, state, "Notes", stateOther.Notes, state.Notes); ops != "" {
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
	if statemachine.InitialState != statemachineOther.InitialState {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "InitialState"))
	}
	if ops := __gong__diffSliceOfPointers(stage, statemachine, "States", statemachineOther.States, statemachine.States); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, statemachine, "Diagrams", statemachineOther.Diagrams, statemachine.Diagrams); ops != "" {
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
	if stateshape.State != stateshapeOther.State {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "State"))
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
	if transition.Start != transitionOther.Start {
		diffs = append(diffs, transition.GongMarshallField(stage, "Start"))
	}
	if transition.End != transitionOther.End {
		diffs = append(diffs, transition.GongMarshallField(stage, "End"))
	}
	if ops := __gong__diffSliceOfPointers(stage, transition, "RolesWithPermissions", transitionOther.RolesWithPermissions, transition.RolesWithPermissions); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, transition, "GeneratedMessages", transitionOther.GeneratedMessages, transition.GeneratedMessages); ops != "" {
		diffs = append(diffs, ops)
	}
	if transition.Guard != transitionOther.Guard {
		diffs = append(diffs, transition.GongMarshallField(stage, "Guard"))
	}
	if ops := __gong__diffSliceOfPointers(stage, transition, "Diagrams", transitionOther.Diagrams, transition.Diagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if transition.IsExpanded != transitionOther.IsExpanded {
		diffs = append(diffs, transition.GongMarshallField(stage, "IsExpanded"))
	}
	if transition.IsRolesNodeExpanded != transitionOther.IsRolesNodeExpanded {
		diffs = append(diffs, transition.GongMarshallField(stage, "IsRolesNodeExpanded"))
	}
	if transition.IsMessagesNodeExpanded != transitionOther.IsMessagesNodeExpanded {
		diffs = append(diffs, transition.GongMarshallField(stage, "IsMessagesNodeExpanded"))
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
	if transition_shape.Transition != transition_shapeOther.Transition {
		diffs = append(diffs, transition_shape.GongMarshallField(stage, "Transition"))
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

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
