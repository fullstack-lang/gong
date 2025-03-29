// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Command:
		if stage.OnAfterCommandCreateCallback != nil {
			stage.OnAfterCommandCreateCallback.OnAfterCreate(stage, target)
		}
	case *DummyAgent:
		if stage.OnAfterDummyAgentCreateCallback != nil {
			stage.OnAfterDummyAgentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Engine:
		if stage.OnAfterEngineCreateCallback != nil {
			stage.OnAfterEngineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Event:
		if stage.OnAfterEventCreateCallback != nil {
			stage.OnAfterEventCreateCallback.OnAfterCreate(stage, target)
		}
	case *Status:
		if stage.OnAfterStatusCreateCallback != nil {
			stage.OnAfterStatusCreateCallback.OnAfterCreate(stage, target)
		}
	case *UpdateState:
		if stage.OnAfterUpdateStateCreateCallback != nil {
			stage.OnAfterUpdateStateCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Command:
		newTarget := any(new).(*Command)
		if stage.OnAfterCommandUpdateCallback != nil {
			stage.OnAfterCommandUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DummyAgent:
		newTarget := any(new).(*DummyAgent)
		if stage.OnAfterDummyAgentUpdateCallback != nil {
			stage.OnAfterDummyAgentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Engine:
		newTarget := any(new).(*Engine)
		if stage.OnAfterEngineUpdateCallback != nil {
			stage.OnAfterEngineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Event:
		newTarget := any(new).(*Event)
		if stage.OnAfterEventUpdateCallback != nil {
			stage.OnAfterEventUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Status:
		newTarget := any(new).(*Status)
		if stage.OnAfterStatusUpdateCallback != nil {
			stage.OnAfterStatusUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *UpdateState:
		newTarget := any(new).(*UpdateState)
		if stage.OnAfterUpdateStateUpdateCallback != nil {
			stage.OnAfterUpdateStateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Command:
		if stage.OnAfterCommandDeleteCallback != nil {
			staged := any(staged).(*Command)
			stage.OnAfterCommandDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DummyAgent:
		if stage.OnAfterDummyAgentDeleteCallback != nil {
			staged := any(staged).(*DummyAgent)
			stage.OnAfterDummyAgentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Engine:
		if stage.OnAfterEngineDeleteCallback != nil {
			staged := any(staged).(*Engine)
			stage.OnAfterEngineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Event:
		if stage.OnAfterEventDeleteCallback != nil {
			staged := any(staged).(*Event)
			stage.OnAfterEventDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Status:
		if stage.OnAfterStatusDeleteCallback != nil {
			staged := any(staged).(*Status)
			stage.OnAfterStatusDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *UpdateState:
		if stage.OnAfterUpdateStateDeleteCallback != nil {
			staged := any(staged).(*UpdateState)
			stage.OnAfterUpdateStateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Command:
		if stage.OnAfterCommandReadCallback != nil {
			stage.OnAfterCommandReadCallback.OnAfterRead(stage, target)
		}
	case *DummyAgent:
		if stage.OnAfterDummyAgentReadCallback != nil {
			stage.OnAfterDummyAgentReadCallback.OnAfterRead(stage, target)
		}
	case *Engine:
		if stage.OnAfterEngineReadCallback != nil {
			stage.OnAfterEngineReadCallback.OnAfterRead(stage, target)
		}
	case *Event:
		if stage.OnAfterEventReadCallback != nil {
			stage.OnAfterEventReadCallback.OnAfterRead(stage, target)
		}
	case *Status:
		if stage.OnAfterStatusReadCallback != nil {
			stage.OnAfterStatusReadCallback.OnAfterRead(stage, target)
		}
	case *UpdateState:
		if stage.OnAfterUpdateStateReadCallback != nil {
			stage.OnAfterUpdateStateReadCallback.OnAfterRead(stage, target)
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
	case *Command:
		stage.OnAfterCommandUpdateCallback = any(callback).(OnAfterUpdateInterface[Command])
	
	case *DummyAgent:
		stage.OnAfterDummyAgentUpdateCallback = any(callback).(OnAfterUpdateInterface[DummyAgent])
	
	case *Engine:
		stage.OnAfterEngineUpdateCallback = any(callback).(OnAfterUpdateInterface[Engine])
	
	case *Event:
		stage.OnAfterEventUpdateCallback = any(callback).(OnAfterUpdateInterface[Event])
	
	case *Status:
		stage.OnAfterStatusUpdateCallback = any(callback).(OnAfterUpdateInterface[Status])
	
	case *UpdateState:
		stage.OnAfterUpdateStateUpdateCallback = any(callback).(OnAfterUpdateInterface[UpdateState])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Command:
		stage.OnAfterCommandCreateCallback = any(callback).(OnAfterCreateInterface[Command])
	
	case *DummyAgent:
		stage.OnAfterDummyAgentCreateCallback = any(callback).(OnAfterCreateInterface[DummyAgent])
	
	case *Engine:
		stage.OnAfterEngineCreateCallback = any(callback).(OnAfterCreateInterface[Engine])
	
	case *Event:
		stage.OnAfterEventCreateCallback = any(callback).(OnAfterCreateInterface[Event])
	
	case *Status:
		stage.OnAfterStatusCreateCallback = any(callback).(OnAfterCreateInterface[Status])
	
	case *UpdateState:
		stage.OnAfterUpdateStateCreateCallback = any(callback).(OnAfterCreateInterface[UpdateState])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Command:
		stage.OnAfterCommandDeleteCallback = any(callback).(OnAfterDeleteInterface[Command])
	
	case *DummyAgent:
		stage.OnAfterDummyAgentDeleteCallback = any(callback).(OnAfterDeleteInterface[DummyAgent])
	
	case *Engine:
		stage.OnAfterEngineDeleteCallback = any(callback).(OnAfterDeleteInterface[Engine])
	
	case *Event:
		stage.OnAfterEventDeleteCallback = any(callback).(OnAfterDeleteInterface[Event])
	
	case *Status:
		stage.OnAfterStatusDeleteCallback = any(callback).(OnAfterDeleteInterface[Status])
	
	case *UpdateState:
		stage.OnAfterUpdateStateDeleteCallback = any(callback).(OnAfterDeleteInterface[UpdateState])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Command:
		stage.OnAfterCommandReadCallback = any(callback).(OnAfterReadInterface[Command])
	
	case *DummyAgent:
		stage.OnAfterDummyAgentReadCallback = any(callback).(OnAfterReadInterface[DummyAgent])
	
	case *Engine:
		stage.OnAfterEngineReadCallback = any(callback).(OnAfterReadInterface[Engine])
	
	case *Event:
		stage.OnAfterEventReadCallback = any(callback).(OnAfterReadInterface[Event])
	
	case *Status:
		stage.OnAfterStatusReadCallback = any(callback).(OnAfterReadInterface[Status])
	
	case *UpdateState:
		stage.OnAfterUpdateStateReadCallback = any(callback).(OnAfterReadInterface[UpdateState])
	
	}
}
