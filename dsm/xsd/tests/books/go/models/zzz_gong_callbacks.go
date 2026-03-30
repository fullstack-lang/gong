// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *BookType:
		if stage.OnAfterBookTypeCreateCallback != nil {
			stage.OnAfterBookTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Books:
		if stage.OnAfterBooksCreateCallback != nil {
			stage.OnAfterBooksCreateCallback.OnAfterCreate(stage, target)
		}
	case *Credit:
		if stage.OnAfterCreditCreateCallback != nil {
			stage.OnAfterCreditCreateCallback.OnAfterCreate(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkCreateCallback != nil {
			stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, target)
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
	case *BookType:
		newTarget := any(new).(*BookType)
		if stage.OnAfterBookTypeUpdateCallback != nil {
			stage.OnAfterBookTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Books:
		newTarget := any(new).(*Books)
		if stage.OnAfterBooksUpdateCallback != nil {
			stage.OnAfterBooksUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Credit:
		newTarget := any(new).(*Credit)
		if stage.OnAfterCreditUpdateCallback != nil {
			stage.OnAfterCreditUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Link:
		newTarget := any(new).(*Link)
		if stage.OnAfterLinkUpdateCallback != nil {
			stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *BookType:
		if stage.OnAfterBookTypeDeleteCallback != nil {
			staged := any(staged).(*BookType)
			stage.OnAfterBookTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Books:
		if stage.OnAfterBooksDeleteCallback != nil {
			staged := any(staged).(*Books)
			stage.OnAfterBooksDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Credit:
		if stage.OnAfterCreditDeleteCallback != nil {
			staged := any(staged).(*Credit)
			stage.OnAfterCreditDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Link:
		if stage.OnAfterLinkDeleteCallback != nil {
			staged := any(staged).(*Link)
			stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *BookType:
		if stage.OnAfterBookTypeReadCallback != nil {
			stage.OnAfterBookTypeReadCallback.OnAfterRead(stage, target)
		}
	case *Books:
		if stage.OnAfterBooksReadCallback != nil {
			stage.OnAfterBooksReadCallback.OnAfterRead(stage, target)
		}
	case *Credit:
		if stage.OnAfterCreditReadCallback != nil {
			stage.OnAfterCreditReadCallback.OnAfterRead(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkReadCallback != nil {
			stage.OnAfterLinkReadCallback.OnAfterRead(stage, target)
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
	case *BookType:
		stage.OnAfterBookTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[BookType])
	case *Books:
		stage.OnAfterBooksUpdateCallback = any(callback).(OnAfterUpdateInterface[Books])
	case *Credit:
		stage.OnAfterCreditUpdateCallback = any(callback).(OnAfterUpdateInterface[Credit])
	case *Link:
		stage.OnAfterLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[Link])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BookType:
		stage.OnAfterBookTypeCreateCallback = any(callback).(OnAfterCreateInterface[BookType])
	case *Books:
		stage.OnAfterBooksCreateCallback = any(callback).(OnAfterCreateInterface[Books])
	case *Credit:
		stage.OnAfterCreditCreateCallback = any(callback).(OnAfterCreateInterface[Credit])
	case *Link:
		stage.OnAfterLinkCreateCallback = any(callback).(OnAfterCreateInterface[Link])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BookType:
		stage.OnAfterBookTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[BookType])
	case *Books:
		stage.OnAfterBooksDeleteCallback = any(callback).(OnAfterDeleteInterface[Books])
	case *Credit:
		stage.OnAfterCreditDeleteCallback = any(callback).(OnAfterDeleteInterface[Credit])
	case *Link:
		stage.OnAfterLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[Link])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BookType:
		stage.OnAfterBookTypeReadCallback = any(callback).(OnAfterReadInterface[BookType])
	case *Books:
		stage.OnAfterBooksReadCallback = any(callback).(OnAfterReadInterface[Books])
	case *Credit:
		stage.OnAfterCreditReadCallback = any(callback).(OnAfterReadInterface[Credit])
	case *Link:
		stage.OnAfterLinkReadCallback = any(callback).(OnAfterReadInterface[Link])
	}
}
