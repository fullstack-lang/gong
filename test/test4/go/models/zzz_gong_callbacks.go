// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Astruct:
		if stage.OnAfterAstructCreateCallback != nil {
			stage.OnAfterAstructCreateCallback.OnAfterCreate(stage, target)
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
	case *Astruct:
		newTarget := any(new).(*Astruct)
		if stage.OnAfterAstructUpdateCallback != nil {
			stage.OnAfterAstructUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Astruct:
		if stage.OnAfterAstructDeleteCallback != nil {
			staged := any(staged).(*Astruct)
			stage.OnAfterAstructDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Astruct:
		if stage.OnAfterAstructReadCallback != nil {
			stage.OnAfterAstructReadCallback.OnAfterRead(stage, target)
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
	case *Astruct:
		stage.OnAfterAstructUpdateCallback = any(callback).(OnAfterUpdateInterface[Astruct])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Astruct:
		stage.OnAfterAstructCreateCallback = any(callback).(OnAfterCreateInterface[Astruct])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Astruct:
		stage.OnAfterAstructDeleteCallback = any(callback).(OnAfterDeleteInterface[Astruct])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Astruct:
		stage.OnAfterAstructReadCallback = any(callback).(OnAfterReadInterface[Astruct])
	}
}
