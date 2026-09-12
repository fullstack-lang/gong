// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *A:
		if stage.OnAfterACreateCallback != nil {
			stage.OnAfterACreateCallback.OnAfterCreate(stage, target)
		}
	case *B:
		if stage.OnAfterBCreateCallback != nil {
			stage.OnAfterBCreateCallback.OnAfterCreate(stage, target)
		}
	case *C:
		if stage.OnAfterCCreateCallback != nil {
			stage.OnAfterCCreateCallback.OnAfterCreate(stage, target)
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
	case *A:
		newTarget := any(new).(*A)
		if stage.OnAfterAUpdateCallback != nil {
			stage.OnAfterAUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *B:
		newTarget := any(new).(*B)
		if stage.OnAfterBUpdateCallback != nil {
			stage.OnAfterBUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *C:
		newTarget := any(new).(*C)
		if stage.OnAfterCUpdateCallback != nil {
			stage.OnAfterCUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *A:
		if stage.OnAfterADeleteCallback != nil {
			staged := any(staged).(*A)
			stage.OnAfterADeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *B:
		if stage.OnAfterBDeleteCallback != nil {
			staged := any(staged).(*B)
			stage.OnAfterBDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *C:
		if stage.OnAfterCDeleteCallback != nil {
			staged := any(staged).(*C)
			stage.OnAfterCDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
