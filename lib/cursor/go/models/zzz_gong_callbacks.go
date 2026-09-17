// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (cursor *Cursor) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCursorCreateCallback != nil {
		stage.OnAfterCursorCreateCallback.OnAfterCreate(stage, cursor)
	}
}

func (cursor *Cursor) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCursorUpdateCallback != nil {
		var frontCursor *Cursor
		if front != nil {
			frontCursor, _ = front.(*Cursor)
		}
		stage.OnAfterCursorUpdateCallback.OnAfterUpdate(stage, cursor, frontCursor)
	}
}

func (cursor *Cursor) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCursorDeleteCallback != nil {
		var frontCursor *Cursor
		if front != nil {
			frontCursor, _ = front.(*Cursor)
		}
		stage.OnAfterCursorDeleteCallback.OnAfterDelete(stage, cursor, frontCursor)
	}
}

