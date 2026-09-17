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
func (button *Button) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterButtonCreateCallback != nil {
		stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, button)
	}
}

func (button *Button) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonUpdateCallback != nil {
		var frontButton *Button
		if front != nil {
			frontButton, _ = front.(*Button)
		}
		stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, button, frontButton)
	}
}

func (button *Button) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonDeleteCallback != nil {
		var frontButton *Button
		if front != nil {
			frontButton, _ = front.(*Button)
		}
		stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, button, frontButton)
	}
}

func (buttontoggle *ButtonToggle) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterButtonToggleCreateCallback != nil {
		stage.OnAfterButtonToggleCreateCallback.OnAfterCreate(stage, buttontoggle)
	}
}

func (buttontoggle *ButtonToggle) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonToggleUpdateCallback != nil {
		var frontButtonToggle *ButtonToggle
		if front != nil {
			frontButtonToggle, _ = front.(*ButtonToggle)
		}
		stage.OnAfterButtonToggleUpdateCallback.OnAfterUpdate(stage, buttontoggle, frontButtonToggle)
	}
}

func (buttontoggle *ButtonToggle) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonToggleDeleteCallback != nil {
		var frontButtonToggle *ButtonToggle
		if front != nil {
			frontButtonToggle, _ = front.(*ButtonToggle)
		}
		stage.OnAfterButtonToggleDeleteCallback.OnAfterDelete(stage, buttontoggle, frontButtonToggle)
	}
}

func (group *Group) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupCreateCallback != nil {
		stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, group)
	}
}

func (group *Group) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupUpdateCallback != nil {
		var frontGroup *Group
		if front != nil {
			frontGroup, _ = front.(*Group)
		}
		stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, group, frontGroup)
	}
}

func (group *Group) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupDeleteCallback != nil {
		var frontGroup *Group
		if front != nil {
			frontGroup, _ = front.(*Group)
		}
		stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, group, frontGroup)
	}
}

func (grouptoogle *GroupToogle) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupToogleCreateCallback != nil {
		stage.OnAfterGroupToogleCreateCallback.OnAfterCreate(stage, grouptoogle)
	}
}

func (grouptoogle *GroupToogle) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupToogleUpdateCallback != nil {
		var frontGroupToogle *GroupToogle
		if front != nil {
			frontGroupToogle, _ = front.(*GroupToogle)
		}
		stage.OnAfterGroupToogleUpdateCallback.OnAfterUpdate(stage, grouptoogle, frontGroupToogle)
	}
}

func (grouptoogle *GroupToogle) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupToogleDeleteCallback != nil {
		var frontGroupToogle *GroupToogle
		if front != nil {
			frontGroupToogle, _ = front.(*GroupToogle)
		}
		stage.OnAfterGroupToogleDeleteCallback.OnAfterDelete(stage, grouptoogle, frontGroupToogle)
	}
}

func (layout *Layout) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLayoutCreateCallback != nil {
		stage.OnAfterLayoutCreateCallback.OnAfterCreate(stage, layout)
	}
}

func (layout *Layout) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLayoutUpdateCallback != nil {
		var frontLayout *Layout
		if front != nil {
			frontLayout, _ = front.(*Layout)
		}
		stage.OnAfterLayoutUpdateCallback.OnAfterUpdate(stage, layout, frontLayout)
	}
}

func (layout *Layout) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLayoutDeleteCallback != nil {
		var frontLayout *Layout
		if front != nil {
			frontLayout, _ = front.(*Layout)
		}
		stage.OnAfterLayoutDeleteCallback.OnAfterDelete(stage, layout, frontLayout)
	}
}

