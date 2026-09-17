// generated code - do not edit
package y

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
func (y *Y) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterYCreateCallback != nil {
		stage.OnAfterYCreateCallback.OnAfterCreate(stage, y)
	}
}

func (y *Y) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterYUpdateCallback != nil {
		var frontY *Y
		if front != nil {
			frontY, _ = front.(*Y)
		}
		stage.OnAfterYUpdateCallback.OnAfterUpdate(stage, y, frontY)
	}
}

func (y *Y) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterYDeleteCallback != nil {
		var frontY *Y
		if front != nil {
			frontY, _ = front.(*Y)
		}
		stage.OnAfterYDeleteCallback.OnAfterDelete(stage, y, frontY)
	}
}

