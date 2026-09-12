// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
