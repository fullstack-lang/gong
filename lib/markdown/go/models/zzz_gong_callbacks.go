// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Content:
		if stage.OnAfterContentCreateCallback != nil {
			stage.OnAfterContentCreateCallback.OnAfterCreate(stage, target)
		}
	case *SvgImage:
		if stage.OnAfterSvgImageCreateCallback != nil {
			stage.OnAfterSvgImageCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type, mouseEvent *Gong__MouseEvent) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Content:
		newTarget := any(new).(*Content)
		if stage.OnAfterContentUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterContentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterContentUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterContentUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *SvgImage:
		newTarget := any(new).(*SvgImage)
		if stage.OnAfterSvgImageUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSvgImageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSvgImageUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSvgImageUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Content:
		if stage.OnAfterContentDeleteCallback != nil {
			staged := any(staged).(*Content)
			stage.OnAfterContentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SvgImage:
		if stage.OnAfterSvgImageDeleteCallback != nil {
			staged := any(staged).(*SvgImage)
			stage.OnAfterSvgImageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Content:
		if stage.OnAfterContentReadCallback != nil {
			stage.OnAfterContentReadCallback.OnAfterRead(stage, target)
		}
	case *SvgImage:
		if stage.OnAfterSvgImageReadCallback != nil {
			stage.OnAfterSvgImageReadCallback.OnAfterRead(stage, target)
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
	case *Content:
		stage.OnAfterContentUpdateCallback = any(callback).(OnAfterUpdateInterface[Content])
	
	case *SvgImage:
		stage.OnAfterSvgImageUpdateCallback = any(callback).(OnAfterUpdateInterface[SvgImage])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Content:
		stage.OnAfterContentCreateCallback = any(callback).(OnAfterCreateInterface[Content])
	
	case *SvgImage:
		stage.OnAfterSvgImageCreateCallback = any(callback).(OnAfterCreateInterface[SvgImage])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Content:
		stage.OnAfterContentDeleteCallback = any(callback).(OnAfterDeleteInterface[Content])
	
	case *SvgImage:
		stage.OnAfterSvgImageDeleteCallback = any(callback).(OnAfterDeleteInterface[SvgImage])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Content:
		stage.OnAfterContentReadCallback = any(callback).(OnAfterReadInterface[Content])
	
	case *SvgImage:
		stage.OnAfterSvgImageReadCallback = any(callback).(OnAfterReadInterface[SvgImage])
	
	}
}
