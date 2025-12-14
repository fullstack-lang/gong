// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Action:
		if stage.OnAfterActionCreateCallback != nil {
			stage.OnAfterActionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Activities:
		if stage.OnAfterActivitiesCreateCallback != nil {
			stage.OnAfterActivitiesCreateCallback.OnAfterCreate(stage, target)
		}
	case *Architecture:
		if stage.OnAfterArchitectureCreateCallback != nil {
			stage.OnAfterArchitectureCreateCallback.OnAfterCreate(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Kill:
		if stage.OnAfterKillCreateCallback != nil {
			stage.OnAfterKillCreateCallback.OnAfterCreate(stage, target)
		}
	case *Message:
		if stage.OnAfterMessageCreateCallback != nil {
			stage.OnAfterMessageCreateCallback.OnAfterCreate(stage, target)
		}
	case *MessageType:
		if stage.OnAfterMessageTypeCreateCallback != nil {
			stage.OnAfterMessageTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Object:
		if stage.OnAfterObjectCreateCallback != nil {
			stage.OnAfterObjectCreateCallback.OnAfterCreate(stage, target)
		}
	case *Role:
		if stage.OnAfterRoleCreateCallback != nil {
			stage.OnAfterRoleCreateCallback.OnAfterCreate(stage, target)
		}
	case *State:
		if stage.OnAfterStateCreateCallback != nil {
			stage.OnAfterStateCreateCallback.OnAfterCreate(stage, target)
		}
	case *StateMachine:
		if stage.OnAfterStateMachineCreateCallback != nil {
			stage.OnAfterStateMachineCreateCallback.OnAfterCreate(stage, target)
		}
	case *StateShape:
		if stage.OnAfterStateShapeCreateCallback != nil {
			stage.OnAfterStateShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Transition:
		if stage.OnAfterTransitionCreateCallback != nil {
			stage.OnAfterTransitionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Transition_Shape:
		if stage.OnAfterTransition_ShapeCreateCallback != nil {
			stage.OnAfterTransition_ShapeCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Action:
		newTarget := any(new).(*Action)
		if stage.OnAfterActionUpdateCallback != nil {
			stage.OnAfterActionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Activities:
		newTarget := any(new).(*Activities)
		if stage.OnAfterActivitiesUpdateCallback != nil {
			stage.OnAfterActivitiesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Architecture:
		newTarget := any(new).(*Architecture)
		if stage.OnAfterArchitectureUpdateCallback != nil {
			stage.OnAfterArchitectureUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Diagram:
		newTarget := any(new).(*Diagram)
		if stage.OnAfterDiagramUpdateCallback != nil {
			stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Kill:
		newTarget := any(new).(*Kill)
		if stage.OnAfterKillUpdateCallback != nil {
			stage.OnAfterKillUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Message:
		newTarget := any(new).(*Message)
		if stage.OnAfterMessageUpdateCallback != nil {
			stage.OnAfterMessageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MessageType:
		newTarget := any(new).(*MessageType)
		if stage.OnAfterMessageTypeUpdateCallback != nil {
			stage.OnAfterMessageTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Object:
		newTarget := any(new).(*Object)
		if stage.OnAfterObjectUpdateCallback != nil {
			stage.OnAfterObjectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Role:
		newTarget := any(new).(*Role)
		if stage.OnAfterRoleUpdateCallback != nil {
			stage.OnAfterRoleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *State:
		newTarget := any(new).(*State)
		if stage.OnAfterStateUpdateCallback != nil {
			stage.OnAfterStateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StateMachine:
		newTarget := any(new).(*StateMachine)
		if stage.OnAfterStateMachineUpdateCallback != nil {
			stage.OnAfterStateMachineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StateShape:
		newTarget := any(new).(*StateShape)
		if stage.OnAfterStateShapeUpdateCallback != nil {
			stage.OnAfterStateShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Transition:
		newTarget := any(new).(*Transition)
		if stage.OnAfterTransitionUpdateCallback != nil {
			stage.OnAfterTransitionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Transition_Shape:
		newTarget := any(new).(*Transition_Shape)
		if stage.OnAfterTransition_ShapeUpdateCallback != nil {
			stage.OnAfterTransition_ShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Action:
		if stage.OnAfterActionDeleteCallback != nil {
			staged := any(staged).(*Action)
			stage.OnAfterActionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Activities:
		if stage.OnAfterActivitiesDeleteCallback != nil {
			staged := any(staged).(*Activities)
			stage.OnAfterActivitiesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Architecture:
		if stage.OnAfterArchitectureDeleteCallback != nil {
			staged := any(staged).(*Architecture)
			stage.OnAfterArchitectureDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Diagram:
		if stage.OnAfterDiagramDeleteCallback != nil {
			staged := any(staged).(*Diagram)
			stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Kill:
		if stage.OnAfterKillDeleteCallback != nil {
			staged := any(staged).(*Kill)
			stage.OnAfterKillDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Message:
		if stage.OnAfterMessageDeleteCallback != nil {
			staged := any(staged).(*Message)
			stage.OnAfterMessageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MessageType:
		if stage.OnAfterMessageTypeDeleteCallback != nil {
			staged := any(staged).(*MessageType)
			stage.OnAfterMessageTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Object:
		if stage.OnAfterObjectDeleteCallback != nil {
			staged := any(staged).(*Object)
			stage.OnAfterObjectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Role:
		if stage.OnAfterRoleDeleteCallback != nil {
			staged := any(staged).(*Role)
			stage.OnAfterRoleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *State:
		if stage.OnAfterStateDeleteCallback != nil {
			staged := any(staged).(*State)
			stage.OnAfterStateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StateMachine:
		if stage.OnAfterStateMachineDeleteCallback != nil {
			staged := any(staged).(*StateMachine)
			stage.OnAfterStateMachineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StateShape:
		if stage.OnAfterStateShapeDeleteCallback != nil {
			staged := any(staged).(*StateShape)
			stage.OnAfterStateShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Transition:
		if stage.OnAfterTransitionDeleteCallback != nil {
			staged := any(staged).(*Transition)
			stage.OnAfterTransitionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Transition_Shape:
		if stage.OnAfterTransition_ShapeDeleteCallback != nil {
			staged := any(staged).(*Transition_Shape)
			stage.OnAfterTransition_ShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Action:
		if stage.OnAfterActionReadCallback != nil {
			stage.OnAfterActionReadCallback.OnAfterRead(stage, target)
		}
	case *Activities:
		if stage.OnAfterActivitiesReadCallback != nil {
			stage.OnAfterActivitiesReadCallback.OnAfterRead(stage, target)
		}
	case *Architecture:
		if stage.OnAfterArchitectureReadCallback != nil {
			stage.OnAfterArchitectureReadCallback.OnAfterRead(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramReadCallback != nil {
			stage.OnAfterDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *Kill:
		if stage.OnAfterKillReadCallback != nil {
			stage.OnAfterKillReadCallback.OnAfterRead(stage, target)
		}
	case *Message:
		if stage.OnAfterMessageReadCallback != nil {
			stage.OnAfterMessageReadCallback.OnAfterRead(stage, target)
		}
	case *MessageType:
		if stage.OnAfterMessageTypeReadCallback != nil {
			stage.OnAfterMessageTypeReadCallback.OnAfterRead(stage, target)
		}
	case *Object:
		if stage.OnAfterObjectReadCallback != nil {
			stage.OnAfterObjectReadCallback.OnAfterRead(stage, target)
		}
	case *Role:
		if stage.OnAfterRoleReadCallback != nil {
			stage.OnAfterRoleReadCallback.OnAfterRead(stage, target)
		}
	case *State:
		if stage.OnAfterStateReadCallback != nil {
			stage.OnAfterStateReadCallback.OnAfterRead(stage, target)
		}
	case *StateMachine:
		if stage.OnAfterStateMachineReadCallback != nil {
			stage.OnAfterStateMachineReadCallback.OnAfterRead(stage, target)
		}
	case *StateShape:
		if stage.OnAfterStateShapeReadCallback != nil {
			stage.OnAfterStateShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Transition:
		if stage.OnAfterTransitionReadCallback != nil {
			stage.OnAfterTransitionReadCallback.OnAfterRead(stage, target)
		}
	case *Transition_Shape:
		if stage.OnAfterTransition_ShapeReadCallback != nil {
			stage.OnAfterTransition_ShapeReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Action:
		stage.OnAfterActionUpdateCallback = any(callback).(OnAfterUpdateInterface[Action])
	
	case *Activities:
		stage.OnAfterActivitiesUpdateCallback = any(callback).(OnAfterUpdateInterface[Activities])
	
	case *Architecture:
		stage.OnAfterArchitectureUpdateCallback = any(callback).(OnAfterUpdateInterface[Architecture])
	
	case *Diagram:
		stage.OnAfterDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Diagram])
	
	case *Kill:
		stage.OnAfterKillUpdateCallback = any(callback).(OnAfterUpdateInterface[Kill])
	
	case *Message:
		stage.OnAfterMessageUpdateCallback = any(callback).(OnAfterUpdateInterface[Message])
	
	case *MessageType:
		stage.OnAfterMessageTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[MessageType])
	
	case *Object:
		stage.OnAfterObjectUpdateCallback = any(callback).(OnAfterUpdateInterface[Object])
	
	case *Role:
		stage.OnAfterRoleUpdateCallback = any(callback).(OnAfterUpdateInterface[Role])
	
	case *State:
		stage.OnAfterStateUpdateCallback = any(callback).(OnAfterUpdateInterface[State])
	
	case *StateMachine:
		stage.OnAfterStateMachineUpdateCallback = any(callback).(OnAfterUpdateInterface[StateMachine])
	
	case *StateShape:
		stage.OnAfterStateShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StateShape])
	
	case *Transition:
		stage.OnAfterTransitionUpdateCallback = any(callback).(OnAfterUpdateInterface[Transition])
	
	case *Transition_Shape:
		stage.OnAfterTransition_ShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[Transition_Shape])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Action:
		stage.OnAfterActionCreateCallback = any(callback).(OnAfterCreateInterface[Action])
	
	case *Activities:
		stage.OnAfterActivitiesCreateCallback = any(callback).(OnAfterCreateInterface[Activities])
	
	case *Architecture:
		stage.OnAfterArchitectureCreateCallback = any(callback).(OnAfterCreateInterface[Architecture])
	
	case *Diagram:
		stage.OnAfterDiagramCreateCallback = any(callback).(OnAfterCreateInterface[Diagram])
	
	case *Kill:
		stage.OnAfterKillCreateCallback = any(callback).(OnAfterCreateInterface[Kill])
	
	case *Message:
		stage.OnAfterMessageCreateCallback = any(callback).(OnAfterCreateInterface[Message])
	
	case *MessageType:
		stage.OnAfterMessageTypeCreateCallback = any(callback).(OnAfterCreateInterface[MessageType])
	
	case *Object:
		stage.OnAfterObjectCreateCallback = any(callback).(OnAfterCreateInterface[Object])
	
	case *Role:
		stage.OnAfterRoleCreateCallback = any(callback).(OnAfterCreateInterface[Role])
	
	case *State:
		stage.OnAfterStateCreateCallback = any(callback).(OnAfterCreateInterface[State])
	
	case *StateMachine:
		stage.OnAfterStateMachineCreateCallback = any(callback).(OnAfterCreateInterface[StateMachine])
	
	case *StateShape:
		stage.OnAfterStateShapeCreateCallback = any(callback).(OnAfterCreateInterface[StateShape])
	
	case *Transition:
		stage.OnAfterTransitionCreateCallback = any(callback).(OnAfterCreateInterface[Transition])
	
	case *Transition_Shape:
		stage.OnAfterTransition_ShapeCreateCallback = any(callback).(OnAfterCreateInterface[Transition_Shape])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Action:
		stage.OnAfterActionDeleteCallback = any(callback).(OnAfterDeleteInterface[Action])
	
	case *Activities:
		stage.OnAfterActivitiesDeleteCallback = any(callback).(OnAfterDeleteInterface[Activities])
	
	case *Architecture:
		stage.OnAfterArchitectureDeleteCallback = any(callback).(OnAfterDeleteInterface[Architecture])
	
	case *Diagram:
		stage.OnAfterDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Diagram])
	
	case *Kill:
		stage.OnAfterKillDeleteCallback = any(callback).(OnAfterDeleteInterface[Kill])
	
	case *Message:
		stage.OnAfterMessageDeleteCallback = any(callback).(OnAfterDeleteInterface[Message])
	
	case *MessageType:
		stage.OnAfterMessageTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[MessageType])
	
	case *Object:
		stage.OnAfterObjectDeleteCallback = any(callback).(OnAfterDeleteInterface[Object])
	
	case *Role:
		stage.OnAfterRoleDeleteCallback = any(callback).(OnAfterDeleteInterface[Role])
	
	case *State:
		stage.OnAfterStateDeleteCallback = any(callback).(OnAfterDeleteInterface[State])
	
	case *StateMachine:
		stage.OnAfterStateMachineDeleteCallback = any(callback).(OnAfterDeleteInterface[StateMachine])
	
	case *StateShape:
		stage.OnAfterStateShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StateShape])
	
	case *Transition:
		stage.OnAfterTransitionDeleteCallback = any(callback).(OnAfterDeleteInterface[Transition])
	
	case *Transition_Shape:
		stage.OnAfterTransition_ShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[Transition_Shape])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Action:
		stage.OnAfterActionReadCallback = any(callback).(OnAfterReadInterface[Action])
	
	case *Activities:
		stage.OnAfterActivitiesReadCallback = any(callback).(OnAfterReadInterface[Activities])
	
	case *Architecture:
		stage.OnAfterArchitectureReadCallback = any(callback).(OnAfterReadInterface[Architecture])
	
	case *Diagram:
		stage.OnAfterDiagramReadCallback = any(callback).(OnAfterReadInterface[Diagram])
	
	case *Kill:
		stage.OnAfterKillReadCallback = any(callback).(OnAfterReadInterface[Kill])
	
	case *Message:
		stage.OnAfterMessageReadCallback = any(callback).(OnAfterReadInterface[Message])
	
	case *MessageType:
		stage.OnAfterMessageTypeReadCallback = any(callback).(OnAfterReadInterface[MessageType])
	
	case *Object:
		stage.OnAfterObjectReadCallback = any(callback).(OnAfterReadInterface[Object])
	
	case *Role:
		stage.OnAfterRoleReadCallback = any(callback).(OnAfterReadInterface[Role])
	
	case *State:
		stage.OnAfterStateReadCallback = any(callback).(OnAfterReadInterface[State])
	
	case *StateMachine:
		stage.OnAfterStateMachineReadCallback = any(callback).(OnAfterReadInterface[StateMachine])
	
	case *StateShape:
		stage.OnAfterStateShapeReadCallback = any(callback).(OnAfterReadInterface[StateShape])
	
	case *Transition:
		stage.OnAfterTransitionReadCallback = any(callback).(OnAfterReadInterface[Transition])
	
	case *Transition_Shape:
		stage.OnAfterTransition_ShapeReadCallback = any(callback).(OnAfterReadInterface[Transition_Shape])
	
	}
}
