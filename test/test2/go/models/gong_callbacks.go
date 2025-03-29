// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

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
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

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
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

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
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *A:
		if stage.OnAfterAReadCallback != nil {
			stage.OnAfterAReadCallback.OnAfterRead(stage, target)
		}
	case *B:
		if stage.OnAfterBReadCallback != nil {
			stage.OnAfterBReadCallback.OnAfterRead(stage, target)
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
	case *A:
		stage.OnAfterAUpdateCallback = any(callback).(OnAfterUpdateInterface[A])
	
	case *B:
		stage.OnAfterBUpdateCallback = any(callback).(OnAfterUpdateInterface[B])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *A:
		stage.OnAfterACreateCallback = any(callback).(OnAfterCreateInterface[A])
	
	case *B:
		stage.OnAfterBCreateCallback = any(callback).(OnAfterCreateInterface[B])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *A:
		stage.OnAfterADeleteCallback = any(callback).(OnAfterDeleteInterface[A])
	
	case *B:
		stage.OnAfterBDeleteCallback = any(callback).(OnAfterDeleteInterface[B])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *A:
		stage.OnAfterAReadCallback = any(callback).(OnAfterReadInterface[A])
	
	case *B:
		stage.OnAfterBReadCallback = any(callback).(OnAfterReadInterface[B])
	
	}
}
