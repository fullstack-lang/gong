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
func (a *A) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterACreateCallback != nil {
		stage.OnAfterACreateCallback.OnAfterCreate(stage, a)
	}
}

func (a *A) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAUpdateCallback != nil {
		var frontA *A
		if front != nil {
			frontA, _ = front.(*A)
		}
		stage.OnAfterAUpdateCallback.OnAfterUpdate(stage, a, frontA)
	}
}

func (a *A) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterADeleteCallback != nil {
		var frontA *A
		if front != nil {
			frontA, _ = front.(*A)
		}
		stage.OnAfterADeleteCallback.OnAfterDelete(stage, a, frontA)
	}
}

func (b *B) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBCreateCallback != nil {
		stage.OnAfterBCreateCallback.OnAfterCreate(stage, b)
	}
}

func (b *B) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBUpdateCallback != nil {
		var frontB *B
		if front != nil {
			frontB, _ = front.(*B)
		}
		stage.OnAfterBUpdateCallback.OnAfterUpdate(stage, b, frontB)
	}
}

func (b *B) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBDeleteCallback != nil {
		var frontB *B
		if front != nil {
			frontB, _ = front.(*B)
		}
		stage.OnAfterBDeleteCallback.OnAfterDelete(stage, b, frontB)
	}
}

func (c *C) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCCreateCallback != nil {
		stage.OnAfterCCreateCallback.OnAfterCreate(stage, c)
	}
}

func (c *C) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCUpdateCallback != nil {
		var frontC *C
		if front != nil {
			frontC, _ = front.(*C)
		}
		stage.OnAfterCUpdateCallback.OnAfterUpdate(stage, c, frontC)
	}
}

func (c *C) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCDeleteCallback != nil {
		var frontC *C
		if front != nil {
			frontC, _ = front.(*C)
		}
		stage.OnAfterCDeleteCallback.OnAfterDelete(stage, c, frontC)
	}
}

