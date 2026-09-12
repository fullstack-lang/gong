// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonCreateCallback != nil {
			stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *ButtonToggle:
		if stage.OnAfterButtonToggleCreateCallback != nil {
			stage.OnAfterButtonToggleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupCreateCallback != nil {
			stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *GroupToogle:
		if stage.OnAfterGroupToogleCreateCallback != nil {
			stage.OnAfterGroupToogleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Layout:
		if stage.OnAfterLayoutCreateCallback != nil {
			stage.OnAfterLayoutCreateCallback.OnAfterCreate(stage, target)
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
	case *Button:
		newTarget := any(new).(*Button)
		if stage.OnAfterButtonUpdateCallback != nil {
			stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ButtonToggle:
		newTarget := any(new).(*ButtonToggle)
		if stage.OnAfterButtonToggleUpdateCallback != nil {
			stage.OnAfterButtonToggleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group:
		newTarget := any(new).(*Group)
		if stage.OnAfterGroupUpdateCallback != nil {
			stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GroupToogle:
		newTarget := any(new).(*GroupToogle)
		if stage.OnAfterGroupToogleUpdateCallback != nil {
			stage.OnAfterGroupToogleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Layout:
		newTarget := any(new).(*Layout)
		if stage.OnAfterLayoutUpdateCallback != nil {
			stage.OnAfterLayoutUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Button:
		if stage.OnAfterButtonDeleteCallback != nil {
			staged := any(staged).(*Button)
			stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ButtonToggle:
		if stage.OnAfterButtonToggleDeleteCallback != nil {
			staged := any(staged).(*ButtonToggle)
			stage.OnAfterButtonToggleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group:
		if stage.OnAfterGroupDeleteCallback != nil {
			staged := any(staged).(*Group)
			stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GroupToogle:
		if stage.OnAfterGroupToogleDeleteCallback != nil {
			staged := any(staged).(*GroupToogle)
			stage.OnAfterGroupToogleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Layout:
		if stage.OnAfterLayoutDeleteCallback != nil {
			staged := any(staged).(*Layout)
			stage.OnAfterLayoutDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
