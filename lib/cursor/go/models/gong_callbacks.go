// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Cursor:
		if stage.OnAfterCursorCreateCallback != nil {
			stage.OnAfterCursorCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Cursor:
		newTarget := any(new).(*Cursor)
		if stage.OnAfterCursorUpdateCallback != nil {
			stage.OnAfterCursorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Cursor:
		if stage.OnAfterCursorDeleteCallback != nil {
			staged := any(staged).(*Cursor)
			stage.OnAfterCursorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Cursor:
		if stage.OnAfterCursorReadCallback != nil {
			stage.OnAfterCursorReadCallback.OnAfterRead(stage, target)
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
	case *Cursor:
		stage.OnAfterCursorUpdateCallback = any(callback).(OnAfterUpdateInterface[Cursor])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Cursor:
		stage.OnAfterCursorCreateCallback = any(callback).(OnAfterCreateInterface[Cursor])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Cursor:
		stage.OnAfterCursorDeleteCallback = any(callback).(OnAfterDeleteInterface[Cursor])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Cursor:
		stage.OnAfterCursorReadCallback = any(callback).(OnAfterReadInterface[Cursor])
	
	}
}
