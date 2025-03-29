// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

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

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

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

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

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

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Checkbox:
		if stage.OnAfterCheckboxReadCallback != nil {
			stage.OnAfterCheckboxReadCallback.OnAfterRead(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupReadCallback != nil {
			stage.OnAfterGroupReadCallback.OnAfterRead(stage, target)
		}
	case *Layout:
		if stage.OnAfterLayoutReadCallback != nil {
			stage.OnAfterLayoutReadCallback.OnAfterRead(stage, target)
		}
	case *Slider:
		if stage.OnAfterSliderReadCallback != nil {
			stage.OnAfterSliderReadCallback.OnAfterRead(stage, target)
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
	case *Checkbox:
		stage.OnAfterCheckboxUpdateCallback = any(callback).(OnAfterUpdateInterface[Checkbox])
	
	case *Group:
		stage.OnAfterGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[Group])
	
	case *Layout:
		stage.OnAfterLayoutUpdateCallback = any(callback).(OnAfterUpdateInterface[Layout])
	
	case *Slider:
		stage.OnAfterSliderUpdateCallback = any(callback).(OnAfterUpdateInterface[Slider])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Checkbox:
		stage.OnAfterCheckboxCreateCallback = any(callback).(OnAfterCreateInterface[Checkbox])
	
	case *Group:
		stage.OnAfterGroupCreateCallback = any(callback).(OnAfterCreateInterface[Group])
	
	case *Layout:
		stage.OnAfterLayoutCreateCallback = any(callback).(OnAfterCreateInterface[Layout])
	
	case *Slider:
		stage.OnAfterSliderCreateCallback = any(callback).(OnAfterCreateInterface[Slider])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Checkbox:
		stage.OnAfterCheckboxDeleteCallback = any(callback).(OnAfterDeleteInterface[Checkbox])
	
	case *Group:
		stage.OnAfterGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[Group])
	
	case *Layout:
		stage.OnAfterLayoutDeleteCallback = any(callback).(OnAfterDeleteInterface[Layout])
	
	case *Slider:
		stage.OnAfterSliderDeleteCallback = any(callback).(OnAfterDeleteInterface[Slider])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Checkbox:
		stage.OnAfterCheckboxReadCallback = any(callback).(OnAfterReadInterface[Checkbox])
	
	case *Group:
		stage.OnAfterGroupReadCallback = any(callback).(OnAfterReadInterface[Group])
	
	case *Layout:
		stage.OnAfterLayoutReadCallback = any(callback).(OnAfterReadInterface[Layout])
	
	case *Slider:
		stage.OnAfterSliderReadCallback = any(callback).(OnAfterReadInterface[Slider])
	
	}
}
