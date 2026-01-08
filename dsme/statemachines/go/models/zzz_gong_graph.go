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

	case *Message:
		ok = stage.IsStagedMessage(target)

	case *MessageType:
		ok = stage.IsStagedMessageType(target)

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

	case *Message:
		ok = stage.IsStagedMessage(target)

	case *MessageType:
		ok = stage.IsStagedMessageType(target)

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

func (stage *Stage) IsStagedMessage(message *Message) (ok bool) {

	_, ok = stage.Messages[message]

	return
}

func (stage *Stage) IsStagedMessageType(messagetype *MessageType) (ok bool) {

	_, ok = stage.MessageTypes[messagetype]

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

	case *Message:
		stage.StageBranchMessage(target)

	case *MessageType:
		stage.StageBranchMessageType(target)

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
	for _, _stateshape := range diagram.State_Shapes {
		StageBranch(stage, _stateshape)
	}
	for _, _transition_shape := range diagram.Transition_Shapes {
		StageBranch(stage, _transition_shape)
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

	case *Message:
		toT := CopyBranchMessage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MessageType:
		toT := CopyBranchMessageType(mapOrigCopy, fromT)
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
	for _, _stateshape := range diagramFrom.State_Shapes {
		diagramTo.State_Shapes = append(diagramTo.State_Shapes, CopyBranchStateShape(mapOrigCopy, _stateshape))
	}
	for _, _transition_shape := range diagramFrom.Transition_Shapes {
		diagramTo.Transition_Shapes = append(diagramTo.Transition_Shapes, CopyBranchTransition_Shape(mapOrigCopy, _transition_shape))
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

	case *Message:
		stage.UnstageBranchMessage(target)

	case *MessageType:
		stage.UnstageBranchMessageType(target)

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
	for _, _stateshape := range diagram.State_Shapes {
		UnstageBranch(stage, _stateshape)
	}
	for _, _transition_shape := range diagram.Transition_Shapes {
		UnstageBranch(stage, _transition_shape)
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
	if diagram.IsInRenameMode != diagramOther.IsInRenameMode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsInRenameMode"))
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
	if state.IsFictif != stateOther.IsFictif {
		diffs = append(diffs, state.GongMarshallField(stage, "IsFictif"))
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
	if statemachine.IsNodeExpanded != statemachineOther.IsNodeExpanded {
		diffs = append(diffs, statemachine.GongMarshallField(stage, "IsNodeExpanded"))
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
	if stateshape.IsExpanded != stateshapeOther.IsExpanded {
		diffs = append(diffs, stateshape.GongMarshallField(stage, "IsExpanded"))
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
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
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
