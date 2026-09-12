// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
