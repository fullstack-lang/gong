// generated code - do not edit
package barline

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
func (barstylecolor *BarStyleColor) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBarStyleColorCreateCallback != nil {
		stage.OnAfterBarStyleColorCreateCallback.OnAfterCreate(stage, barstylecolor)
	}
}

func (barstylecolor *BarStyleColor) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarStyleColorUpdateCallback != nil {
		var frontBarStyleColor *BarStyleColor
		if front != nil {
			frontBarStyleColor, _ = front.(*BarStyleColor)
		}
		stage.OnAfterBarStyleColorUpdateCallback.OnAfterUpdate(stage, barstylecolor, frontBarStyleColor)
	}
}

func (barstylecolor *BarStyleColor) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarStyleColorDeleteCallback != nil {
		var frontBarStyleColor *BarStyleColor
		if front != nil {
			frontBarStyleColor, _ = front.(*BarStyleColor)
		}
		stage.OnAfterBarStyleColorDeleteCallback.OnAfterDelete(stage, barstylecolor, frontBarStyleColor)
	}
}

func (barline *Barline) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBarlineCreateCallback != nil {
		stage.OnAfterBarlineCreateCallback.OnAfterCreate(stage, barline)
	}
}

func (barline *Barline) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarlineUpdateCallback != nil {
		var frontBarline *Barline
		if front != nil {
			frontBarline, _ = front.(*Barline)
		}
		stage.OnAfterBarlineUpdateCallback.OnAfterUpdate(stage, barline, frontBarline)
	}
}

func (barline *Barline) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarlineDeleteCallback != nil {
		var frontBarline *Barline
		if front != nil {
			frontBarline, _ = front.(*Barline)
		}
		stage.OnAfterBarlineDeleteCallback.OnAfterDelete(stage, barline, frontBarline)
	}
}

func (ending *Ending) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEndingCreateCallback != nil {
		stage.OnAfterEndingCreateCallback.OnAfterCreate(stage, ending)
	}
}

func (ending *Ending) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndingUpdateCallback != nil {
		var frontEnding *Ending
		if front != nil {
			frontEnding, _ = front.(*Ending)
		}
		stage.OnAfterEndingUpdateCallback.OnAfterUpdate(stage, ending, frontEnding)
	}
}

func (ending *Ending) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndingDeleteCallback != nil {
		var frontEnding *Ending
		if front != nil {
			frontEnding, _ = front.(*Ending)
		}
		stage.OnAfterEndingDeleteCallback.OnAfterDelete(stage, ending, frontEnding)
	}
}

func (repeat *Repeat) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRepeatCreateCallback != nil {
		stage.OnAfterRepeatCreateCallback.OnAfterCreate(stage, repeat)
	}
}

func (repeat *Repeat) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRepeatUpdateCallback != nil {
		var frontRepeat *Repeat
		if front != nil {
			frontRepeat, _ = front.(*Repeat)
		}
		stage.OnAfterRepeatUpdateCallback.OnAfterUpdate(stage, repeat, frontRepeat)
	}
}

func (repeat *Repeat) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRepeatDeleteCallback != nil {
		var frontRepeat *Repeat
		if front != nil {
			frontRepeat, _ = front.(*Repeat)
		}
		stage.OnAfterRepeatDeleteCallback.OnAfterDelete(stage, repeat, frontRepeat)
	}
}

