// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Checkbox:
		if stage.OnAfterCheckboxCreateCallback != nil {
			stage.OnAfterCheckboxCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupCreateCallback != nil {
			stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *Layout:
		if stage.OnAfterLayoutCreateCallback != nil {
			stage.OnAfterLayoutCreateCallback.OnAfterCreate(stage, target)
		}
	case *Slider:
		if stage.OnAfterSliderCreateCallback != nil {
			stage.OnAfterSliderCreateCallback.OnAfterCreate(stage, target)
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
	case *Checkbox:
		newTarget := any(new).(*Checkbox)
		if stage.OnAfterCheckboxUpdateCallback != nil {
			stage.OnAfterCheckboxUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group:
		newTarget := any(new).(*Group)
		if stage.OnAfterGroupUpdateCallback != nil {
			stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Layout:
		newTarget := any(new).(*Layout)
		if stage.OnAfterLayoutUpdateCallback != nil {
			stage.OnAfterLayoutUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Slider:
		newTarget := any(new).(*Slider)
		if stage.OnAfterSliderUpdateCallback != nil {
			stage.OnAfterSliderUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Checkbox:
		if stage.OnAfterCheckboxDeleteCallback != nil {
			staged := any(staged).(*Checkbox)
			stage.OnAfterCheckboxDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group:
		if stage.OnAfterGroupDeleteCallback != nil {
			staged := any(staged).(*Group)
			stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Layout:
		if stage.OnAfterLayoutDeleteCallback != nil {
			staged := any(staged).(*Layout)
			stage.OnAfterLayoutDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Slider:
		if stage.OnAfterSliderDeleteCallback != nil {
			staged := any(staged).(*Slider)
			stage.OnAfterSliderDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
