// generated code - do not edit
package models

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
