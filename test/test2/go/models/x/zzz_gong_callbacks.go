// generated code - do not edit
package x

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
func (x *X) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXCreateCallback != nil {
		stage.OnAfterXCreateCallback.OnAfterCreate(stage, x)
	}
}

func (x *X) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXUpdateCallback != nil {
		var frontX *X
		if front != nil {
			frontX, _ = front.(*X)
		}
		stage.OnAfterXUpdateCallback.OnAfterUpdate(stage, x, frontX)
	}
}

func (x *X) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXDeleteCallback != nil {
		var frontX *X
		if front != nil {
			frontX, _ = front.(*X)
		}
		stage.OnAfterXDeleteCallback.OnAfterDelete(stage, x, frontX)
	}
}

