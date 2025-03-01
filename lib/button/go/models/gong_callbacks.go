// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonCreateCallback != nil {
			stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupCreateCallback != nil {
			stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *Layout:
		if stage.OnAfterLayoutCreateCallback != nil {
			stage.OnAfterLayoutCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Button:
		newTarget := any(new).(*Button)
		if stage.OnAfterButtonUpdateCallback != nil {
			stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonDeleteCallback != nil {
			staged := any(staged).(*Button)
			stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, staged, front)
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
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonReadCallback != nil {
			stage.OnAfterButtonReadCallback.OnAfterRead(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupReadCallback != nil {
			stage.OnAfterGroupReadCallback.OnAfterRead(stage, target)
		}
	case *Layout:
		if stage.OnAfterLayoutReadCallback != nil {
			stage.OnAfterLayoutReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Button:
		stage.OnAfterButtonUpdateCallback = any(callback).(OnAfterUpdateInterface[Button])
	
	case *Group:
		stage.OnAfterGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[Group])
	
	case *Layout:
		stage.OnAfterLayoutUpdateCallback = any(callback).(OnAfterUpdateInterface[Layout])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Button:
		stage.OnAfterButtonCreateCallback = any(callback).(OnAfterCreateInterface[Button])
	
	case *Group:
		stage.OnAfterGroupCreateCallback = any(callback).(OnAfterCreateInterface[Group])
	
	case *Layout:
		stage.OnAfterLayoutCreateCallback = any(callback).(OnAfterCreateInterface[Layout])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Button:
		stage.OnAfterButtonDeleteCallback = any(callback).(OnAfterDeleteInterface[Button])
	
	case *Group:
		stage.OnAfterGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[Group])
	
	case *Layout:
		stage.OnAfterLayoutDeleteCallback = any(callback).(OnAfterDeleteInterface[Layout])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Button:
		stage.OnAfterButtonReadCallback = any(callback).(OnAfterReadInterface[Button])
	
	case *Group:
		stage.OnAfterGroupReadCallback = any(callback).(OnAfterReadInterface[Group])
	
	case *Layout:
		stage.OnAfterLayoutReadCallback = any(callback).(OnAfterReadInterface[Layout])
	
	}
}
