// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Checkbox:
		ok = stage.IsStagedCheckbox(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *Layout:
		ok = stage.IsStagedLayout(target)

	case *Slider:
		ok = stage.IsStagedSlider(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Checkbox:
		ok = stage.IsStagedCheckbox(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *Layout:
		ok = stage.IsStagedLayout(target)

	case *Slider:
		ok = stage.IsStagedSlider(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedCheckbox(checkbox *Checkbox) (ok bool) {

	_, ok = stage.Checkboxs[checkbox]

	return
}

func (stage *Stage) IsStagedGroup(group *Group) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *Stage) IsStagedLayout(layout *Layout) (ok bool) {

	_, ok = stage.Layouts[layout]

	return
}

func (stage *Stage) IsStagedSlider(slider *Slider) (ok bool) {

	_, ok = stage.Sliders[slider]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Checkbox:
		stage.StageBranchCheckbox(target)

	case *Group:
		stage.StageBranchGroup(target)

	case *Layout:
		stage.StageBranchLayout(target)

	case *Slider:
		stage.StageBranchSlider(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchCheckbox(checkbox *Checkbox) {

	// check if instance is already staged
	if IsStaged(stage, checkbox) {
		return
	}

	checkbox.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGroup(group *Group) {

	// check if instance is already staged
	if IsStaged(stage, group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _slider := range group.Sliders {
		StageBranch(stage, _slider)
	}
	for _, _checkbox := range group.Checkboxes {
		StageBranch(stage, _checkbox)
	}

}

func (stage *Stage) StageBranchLayout(layout *Layout) {

	// check if instance is already staged
	if IsStaged(stage, layout) {
		return
	}

	layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range layout.Groups {
		StageBranch(stage, _group)
	}

}

func (stage *Stage) StageBranchSlider(slider *Slider) {

	// check if instance is already staged
	if IsStaged(stage, slider) {
		return
	}

	slider.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Checkbox:
		toT := CopyBranchCheckbox(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := CopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Layout:
		toT := CopyBranchLayout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slider:
		toT := CopyBranchSlider(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchCheckbox(mapOrigCopy map[any]any, checkboxFrom *Checkbox) (checkboxTo *Checkbox) {

	// checkboxFrom has already been copied
	if _checkboxTo, ok := mapOrigCopy[checkboxFrom]; ok {
		checkboxTo = _checkboxTo.(*Checkbox)
		return
	}

	checkboxTo = new(Checkbox)
	mapOrigCopy[checkboxFrom] = checkboxTo
	checkboxFrom.CopyBasicFields(checkboxTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {

	// groupFrom has already been copied
	if _groupTo, ok := mapOrigCopy[groupFrom]; ok {
		groupTo = _groupTo.(*Group)
		return
	}

	groupTo = new(Group)
	mapOrigCopy[groupFrom] = groupTo
	groupFrom.CopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _slider := range groupFrom.Sliders {
		groupTo.Sliders = append(groupTo.Sliders, CopyBranchSlider(mapOrigCopy, _slider))
	}
	for _, _checkbox := range groupFrom.Checkboxes {
		groupTo.Checkboxes = append(groupTo.Checkboxes, CopyBranchCheckbox(mapOrigCopy, _checkbox))
	}

	return
}

func CopyBranchLayout(mapOrigCopy map[any]any, layoutFrom *Layout) (layoutTo *Layout) {

	// layoutFrom has already been copied
	if _layoutTo, ok := mapOrigCopy[layoutFrom]; ok {
		layoutTo = _layoutTo.(*Layout)
		return
	}

	layoutTo = new(Layout)
	mapOrigCopy[layoutFrom] = layoutTo
	layoutFrom.CopyBasicFields(layoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range layoutFrom.Groups {
		layoutTo.Groups = append(layoutTo.Groups, CopyBranchGroup(mapOrigCopy, _group))
	}

	return
}

func CopyBranchSlider(mapOrigCopy map[any]any, sliderFrom *Slider) (sliderTo *Slider) {

	// sliderFrom has already been copied
	if _sliderTo, ok := mapOrigCopy[sliderFrom]; ok {
		sliderTo = _sliderTo.(*Slider)
		return
	}

	sliderTo = new(Slider)
	mapOrigCopy[sliderFrom] = sliderTo
	sliderFrom.CopyBasicFields(sliderTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Checkbox:
		stage.UnstageBranchCheckbox(target)

	case *Group:
		stage.UnstageBranchGroup(target)

	case *Layout:
		stage.UnstageBranchLayout(target)

	case *Slider:
		stage.UnstageBranchSlider(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchCheckbox(checkbox *Checkbox) {

	// check if instance is already staged
	if !IsStaged(stage, checkbox) {
		return
	}

	checkbox.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGroup(group *Group) {

	// check if instance is already staged
	if !IsStaged(stage, group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _slider := range group.Sliders {
		UnstageBranch(stage, _slider)
	}
	for _, _checkbox := range group.Checkboxes {
		UnstageBranch(stage, _checkbox)
	}

}

func (stage *Stage) UnstageBranchLayout(layout *Layout) {

	// check if instance is already staged
	if !IsStaged(stage, layout) {
		return
	}

	layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range layout.Groups {
		UnstageBranch(stage, _group)
	}

}

func (stage *Stage) UnstageBranchSlider(slider *Slider) {

	// check if instance is already staged
	if !IsStaged(stage, slider) {
		return
	}

	slider.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}


// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (checkbox *Checkbox) GongDiff(checkboxOther *Checkbox) (diffs []string) {
	// insertion point for field diffs
	if checkbox.Name != checkboxOther.Name {
		diffs = append(diffs, "Name")
	}
	if checkbox.ValueBool != checkboxOther.ValueBool {
		diffs = append(diffs, "ValueBool")
	}
	if checkbox.LabelForTrue != checkboxOther.LabelForTrue {
		diffs = append(diffs, "LabelForTrue")
	}
	if checkbox.LabelForFalse != checkboxOther.LabelForFalse {
		diffs = append(diffs, "LabelForFalse")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group *Group) GongDiff(groupOther *Group) (diffs []string) {
	// insertion point for field diffs
	if group.Name != groupOther.Name {
		diffs = append(diffs, "Name")
	}
	if group.Percentage != groupOther.Percentage {
		diffs = append(diffs, "Percentage")
	}
	SlidersDifferent := false
    if len(group.Sliders) != len(groupOther.Sliders) {
        SlidersDifferent = true
    } else {
        for i := range group.Sliders {
            if (group.Sliders[i] == nil) != (groupOther.Sliders[i] == nil) {
                SlidersDifferent = true
                break
            } else if group.Sliders[i] != nil && groupOther.Sliders[i] != nil {
                if len(group.Sliders[i].GongDiff(groupOther.Sliders[i])) > 0 {
                    SlidersDifferent = true
                    break
                }
            }
        }
    }
    if SlidersDifferent {
        diffs = append(diffs, "Sliders")
    }
	CheckboxesDifferent := false
    if len(group.Checkboxes) != len(groupOther.Checkboxes) {
        CheckboxesDifferent = true
    } else {
        for i := range group.Checkboxes {
            if (group.Checkboxes[i] == nil) != (groupOther.Checkboxes[i] == nil) {
                CheckboxesDifferent = true
                break
            } else if group.Checkboxes[i] != nil && groupOther.Checkboxes[i] != nil {
                if len(group.Checkboxes[i].GongDiff(groupOther.Checkboxes[i])) > 0 {
                    CheckboxesDifferent = true
                    break
                }
            }
        }
    }
    if CheckboxesDifferent {
        diffs = append(diffs, "Checkboxes")
    }

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (layout *Layout) GongDiff(layoutOther *Layout) (diffs []string) {
	// insertion point for field diffs
	if layout.Name != layoutOther.Name {
		diffs = append(diffs, "Name")
	}
	GroupsDifferent := false
    if len(layout.Groups) != len(layoutOther.Groups) {
        GroupsDifferent = true
    } else {
        for i := range layout.Groups {
            if (layout.Groups[i] == nil) != (layoutOther.Groups[i] == nil) {
                GroupsDifferent = true
                break
            } else if layout.Groups[i] != nil && layoutOther.Groups[i] != nil {
                if len(layout.Groups[i].GongDiff(layoutOther.Groups[i])) > 0 {
                    GroupsDifferent = true
                    break
                }
            }
        }
    }
    if GroupsDifferent {
        diffs = append(diffs, "Groups")
    }

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slider *Slider) GongDiff(sliderOther *Slider) (diffs []string) {
	// insertion point for field diffs
	if slider.Name != sliderOther.Name {
		diffs = append(diffs, "Name")
	}
	if slider.IsFloat64 != sliderOther.IsFloat64 {
		diffs = append(diffs, "IsFloat64")
	}
	if slider.IsInt != sliderOther.IsInt {
		diffs = append(diffs, "IsInt")
	}
	if slider.MinInt != sliderOther.MinInt {
		diffs = append(diffs, "MinInt")
	}
	if slider.MaxInt != sliderOther.MaxInt {
		diffs = append(diffs, "MaxInt")
	}
	if slider.StepInt != sliderOther.StepInt {
		diffs = append(diffs, "StepInt")
	}
	if slider.ValueInt != sliderOther.ValueInt {
		diffs = append(diffs, "ValueInt")
	}
	if slider.MinFloat64 != sliderOther.MinFloat64 {
		diffs = append(diffs, "MinFloat64")
	}
	if slider.MaxFloat64 != sliderOther.MaxFloat64 {
		diffs = append(diffs, "MaxFloat64")
	}
	if slider.StepFloat64 != sliderOther.StepFloat64 {
		diffs = append(diffs, "StepFloat64")
	}
	if slider.ValueFloat64 != sliderOther.ValueFloat64 {
		diffs = append(diffs, "ValueFloat64")
	}

	return
}
