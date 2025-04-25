// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Chapter:
		if stage.OnAfterChapterCreateCallback != nil {
			stage.OnAfterChapterCreateCallback.OnAfterCreate(stage, target)
		}
	case *Content:
		if stage.OnAfterContentCreateCallback != nil {
			stage.OnAfterContentCreateCallback.OnAfterCreate(stage, target)
		}
	case *Page:
		if stage.OnAfterPageCreateCallback != nil {
			stage.OnAfterPageCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Chapter:
		newTarget := any(new).(*Chapter)
		if stage.OnAfterChapterUpdateCallback != nil {
			stage.OnAfterChapterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Content:
		newTarget := any(new).(*Content)
		if stage.OnAfterContentUpdateCallback != nil {
			stage.OnAfterContentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Page:
		newTarget := any(new).(*Page)
		if stage.OnAfterPageUpdateCallback != nil {
			stage.OnAfterPageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Chapter:
		if stage.OnAfterChapterDeleteCallback != nil {
			staged := any(staged).(*Chapter)
			stage.OnAfterChapterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Content:
		if stage.OnAfterContentDeleteCallback != nil {
			staged := any(staged).(*Content)
			stage.OnAfterContentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Page:
		if stage.OnAfterPageDeleteCallback != nil {
			staged := any(staged).(*Page)
			stage.OnAfterPageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Chapter:
		if stage.OnAfterChapterReadCallback != nil {
			stage.OnAfterChapterReadCallback.OnAfterRead(stage, target)
		}
	case *Content:
		if stage.OnAfterContentReadCallback != nil {
			stage.OnAfterContentReadCallback.OnAfterRead(stage, target)
		}
	case *Page:
		if stage.OnAfterPageReadCallback != nil {
			stage.OnAfterPageReadCallback.OnAfterRead(stage, target)
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
	case *Chapter:
		stage.OnAfterChapterUpdateCallback = any(callback).(OnAfterUpdateInterface[Chapter])
	
	case *Content:
		stage.OnAfterContentUpdateCallback = any(callback).(OnAfterUpdateInterface[Content])
	
	case *Page:
		stage.OnAfterPageUpdateCallback = any(callback).(OnAfterUpdateInterface[Page])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Chapter:
		stage.OnAfterChapterCreateCallback = any(callback).(OnAfterCreateInterface[Chapter])
	
	case *Content:
		stage.OnAfterContentCreateCallback = any(callback).(OnAfterCreateInterface[Content])
	
	case *Page:
		stage.OnAfterPageCreateCallback = any(callback).(OnAfterCreateInterface[Page])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Chapter:
		stage.OnAfterChapterDeleteCallback = any(callback).(OnAfterDeleteInterface[Chapter])
	
	case *Content:
		stage.OnAfterContentDeleteCallback = any(callback).(OnAfterDeleteInterface[Content])
	
	case *Page:
		stage.OnAfterPageDeleteCallback = any(callback).(OnAfterDeleteInterface[Page])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Chapter:
		stage.OnAfterChapterReadCallback = any(callback).(OnAfterReadInterface[Chapter])
	
	case *Content:
		stage.OnAfterContentReadCallback = any(callback).(OnAfterReadInterface[Content])
	
	case *Page:
		stage.OnAfterPageReadCallback = any(callback).(OnAfterReadInterface[Page])
	
	}
}
