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
func (checkbox *Checkbox) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCheckboxCreateCallback != nil {
		stage.OnAfterCheckboxCreateCallback.OnAfterCreate(stage, checkbox)
	}
}

func (checkbox *Checkbox) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckboxUpdateCallback != nil {
		var frontCheckbox *Checkbox
		if front != nil {
			frontCheckbox, _ = front.(*Checkbox)
		}
		stage.OnAfterCheckboxUpdateCallback.OnAfterUpdate(stage, checkbox, frontCheckbox)
	}
}

func (checkbox *Checkbox) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCheckboxDeleteCallback != nil {
		var frontCheckbox *Checkbox
		if front != nil {
			frontCheckbox, _ = front.(*Checkbox)
		}
		stage.OnAfterCheckboxDeleteCallback.OnAfterDelete(stage, checkbox, frontCheckbox)
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

func (slider *Slider) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSliderCreateCallback != nil {
		stage.OnAfterSliderCreateCallback.OnAfterCreate(stage, slider)
	}
}

func (slider *Slider) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSliderUpdateCallback != nil {
		var frontSlider *Slider
		if front != nil {
			frontSlider, _ = front.(*Slider)
		}
		stage.OnAfterSliderUpdateCallback.OnAfterUpdate(stage, slider, frontSlider)
	}
}

func (slider *Slider) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSliderDeleteCallback != nil {
		var frontSlider *Slider
		if front != nil {
			frontSlider, _ = front.(*Slider)
		}
		stage.OnAfterSliderDeleteCallback.OnAfterDelete(stage, slider, frontSlider)
	}
}

