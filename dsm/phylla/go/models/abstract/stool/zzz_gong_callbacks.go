// generated code - do not edit
package stool

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
func (stoolabstract *StoolAbstract) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStoolAbstractCreateCallback != nil {
		stage.OnAfterStoolAbstractCreateCallback.OnAfterCreate(stage, stoolabstract)
	}
}

func (stoolabstract *StoolAbstract) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStoolAbstractUpdateCallback != nil {
		var frontStoolAbstract *StoolAbstract
		if front != nil {
			frontStoolAbstract, _ = front.(*StoolAbstract)
		}
		stage.OnAfterStoolAbstractUpdateCallback.OnAfterUpdate(stage, stoolabstract, frontStoolAbstract)
	}
}

func (stoolabstract *StoolAbstract) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStoolAbstractDeleteCallback != nil {
		var frontStoolAbstract *StoolAbstract
		if front != nil {
			frontStoolAbstract, _ = front.(*StoolAbstract)
		}
		stage.OnAfterStoolAbstractDeleteCallback.OnAfterDelete(stage, stoolabstract, frontStoolAbstract)
	}
}

