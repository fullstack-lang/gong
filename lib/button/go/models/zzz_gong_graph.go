// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Button:
		ok = stage.IsStagedButton(target)

	case *ButtonToggle:
		ok = stage.IsStagedButtonToggle(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *GroupToogle:
		ok = stage.IsStagedGroupToogle(target)

	case *Layout:
		ok = stage.IsStagedLayout(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Button:
		ok = stage.IsStagedButton(target)

	case *ButtonToggle:
		ok = stage.IsStagedButtonToggle(target)

	case *Group:
		ok = stage.IsStagedGroup(target)

	case *GroupToogle:
		ok = stage.IsStagedGroupToogle(target)

	case *Layout:
		ok = stage.IsStagedLayout(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedButton(button *Button) (ok bool) {

	_, ok = stage.Buttons[button]

	return
}

func (stage *Stage) IsStagedButtonToggle(buttontoggle *ButtonToggle) (ok bool) {

	_, ok = stage.ButtonToggles[buttontoggle]

	return
}

func (stage *Stage) IsStagedGroup(group *Group) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *Stage) IsStagedGroupToogle(grouptoogle *GroupToogle) (ok bool) {

	_, ok = stage.GroupToogles[grouptoogle]

	return
}

func (stage *Stage) IsStagedLayout(layout *Layout) (ok bool) {

	_, ok = stage.Layouts[layout]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Button:
		stage.StageBranchButton(target)

	case *ButtonToggle:
		stage.StageBranchButtonToggle(target)

	case *Group:
		stage.StageBranchGroup(target)

	case *GroupToogle:
		stage.StageBranchGroupToogle(target)

	case *Layout:
		stage.StageBranchLayout(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchButton(button *Button) {

	// check if instance is already staged
	if IsStaged(stage, button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchButtonToggle(buttontoggle *ButtonToggle) {

	// check if instance is already staged
	if IsStaged(stage, buttontoggle) {
		return
	}

	buttontoggle.Stage(stage)

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
	for _, _button := range group.Buttons {
		StageBranch(stage, _button)
	}

}

func (stage *Stage) StageBranchGroupToogle(grouptoogle *GroupToogle) {

	// check if instance is already staged
	if IsStaged(stage, grouptoogle) {
		return
	}

	grouptoogle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _buttontoggle := range grouptoogle.ButtonToggles {
		StageBranch(stage, _buttontoggle)
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
	for _, _grouptoogle := range layout.GroupToogles {
		StageBranch(stage, _grouptoogle)
	}

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
	case *Button:
		toT := CopyBranchButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ButtonToggle:
		toT := CopyBranchButtonToggle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := CopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GroupToogle:
		toT := CopyBranchGroupToogle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Layout:
		toT := CopyBranchLayout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchButton(mapOrigCopy map[any]any, buttonFrom *Button) (buttonTo *Button) {

	// buttonFrom has already been copied
	if _buttonTo, ok := mapOrigCopy[buttonFrom]; ok {
		buttonTo = _buttonTo.(*Button)
		return
	}

	buttonTo = new(Button)
	mapOrigCopy[buttonFrom] = buttonTo
	buttonFrom.CopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchButtonToggle(mapOrigCopy map[any]any, buttontoggleFrom *ButtonToggle) (buttontoggleTo *ButtonToggle) {

	// buttontoggleFrom has already been copied
	if _buttontoggleTo, ok := mapOrigCopy[buttontoggleFrom]; ok {
		buttontoggleTo = _buttontoggleTo.(*ButtonToggle)
		return
	}

	buttontoggleTo = new(ButtonToggle)
	mapOrigCopy[buttontoggleFrom] = buttontoggleTo
	buttontoggleFrom.CopyBasicFields(buttontoggleTo)

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
	for _, _button := range groupFrom.Buttons {
		groupTo.Buttons = append(groupTo.Buttons, CopyBranchButton(mapOrigCopy, _button))
	}

	return
}

func CopyBranchGroupToogle(mapOrigCopy map[any]any, grouptoogleFrom *GroupToogle) (grouptoogleTo *GroupToogle) {

	// grouptoogleFrom has already been copied
	if _grouptoogleTo, ok := mapOrigCopy[grouptoogleFrom]; ok {
		grouptoogleTo = _grouptoogleTo.(*GroupToogle)
		return
	}

	grouptoogleTo = new(GroupToogle)
	mapOrigCopy[grouptoogleFrom] = grouptoogleTo
	grouptoogleFrom.CopyBasicFields(grouptoogleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _buttontoggle := range grouptoogleFrom.ButtonToggles {
		grouptoogleTo.ButtonToggles = append(grouptoogleTo.ButtonToggles, CopyBranchButtonToggle(mapOrigCopy, _buttontoggle))
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
	for _, _grouptoogle := range layoutFrom.GroupToogles {
		layoutTo.GroupToogles = append(layoutTo.GroupToogles, CopyBranchGroupToogle(mapOrigCopy, _grouptoogle))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Button:
		stage.UnstageBranchButton(target)

	case *ButtonToggle:
		stage.UnstageBranchButtonToggle(target)

	case *Group:
		stage.UnstageBranchGroup(target)

	case *GroupToogle:
		stage.UnstageBranchGroupToogle(target)

	case *Layout:
		stage.UnstageBranchLayout(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchButton(button *Button) {

	// check if instance is already staged
	if !IsStaged(stage, button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchButtonToggle(buttontoggle *ButtonToggle) {

	// check if instance is already staged
	if !IsStaged(stage, buttontoggle) {
		return
	}

	buttontoggle.Unstage(stage)

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
	for _, _button := range group.Buttons {
		UnstageBranch(stage, _button)
	}

}

func (stage *Stage) UnstageBranchGroupToogle(grouptoogle *GroupToogle) {

	// check if instance is already staged
	if !IsStaged(stage, grouptoogle) {
		return
	}

	grouptoogle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _buttontoggle := range grouptoogle.ButtonToggles {
		UnstageBranch(stage, _buttontoggle)
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
	for _, _grouptoogle := range layout.GroupToogles {
		UnstageBranch(stage, _grouptoogle)
	}

}


// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (button *Button) GongDiff(buttonOther *Button) (diffs []string) {
	// insertion point for field diffs
	if button.Name != buttonOther.Name {
		diffs = append(diffs, "Name")
	}
	if button.Label != buttonOther.Label {
		diffs = append(diffs, "Label")
	}
	if button.Icon != buttonOther.Icon {
		diffs = append(diffs, "Icon")
	}
	if button.IsDisabled != buttonOther.IsDisabled {
		diffs = append(diffs, "IsDisabled")
	}
	if button.Color != buttonOther.Color {
		diffs = append(diffs, "Color")
	}
	if button.MatButtonType != buttonOther.MatButtonType {
		diffs = append(diffs, "MatButtonType")
	}
	if button.MatButtonAppearance != buttonOther.MatButtonAppearance {
		diffs = append(diffs, "MatButtonAppearance")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (buttontoggle *ButtonToggle) GongDiff(buttontoggleOther *ButtonToggle) (diffs []string) {
	// insertion point for field diffs
	if buttontoggle.Name != buttontoggleOther.Name {
		diffs = append(diffs, "Name")
	}
	if buttontoggle.Label != buttontoggleOther.Label {
		diffs = append(diffs, "Label")
	}
	if buttontoggle.Icon != buttontoggleOther.Icon {
		diffs = append(diffs, "Icon")
	}
	if buttontoggle.IsDisabled != buttontoggleOther.IsDisabled {
		diffs = append(diffs, "IsDisabled")
	}
	if buttontoggle.IsChecked != buttontoggleOther.IsChecked {
		diffs = append(diffs, "IsChecked")
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
	ButtonsDifferent := false
    if len(group.Buttons) != len(groupOther.Buttons) {
        ButtonsDifferent = true
    } else {
        for i := range group.Buttons {
            if (group.Buttons[i] == nil) != (groupOther.Buttons[i] == nil) {
                ButtonsDifferent = true
                break
            } else if group.Buttons[i] != nil && groupOther.Buttons[i] != nil {
                if len(group.Buttons[i].GongDiff(groupOther.Buttons[i])) > 0 {
                    ButtonsDifferent = true
                    break
                }
            }
        }
    }
    if ButtonsDifferent {
        diffs = append(diffs, "Buttons")
    }
	if group.NbColumns != groupOther.NbColumns {
		diffs = append(diffs, "NbColumns")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (grouptoogle *GroupToogle) GongDiff(grouptoogleOther *GroupToogle) (diffs []string) {
	// insertion point for field diffs
	if grouptoogle.Name != grouptoogleOther.Name {
		diffs = append(diffs, "Name")
	}
	if grouptoogle.Percentage != grouptoogleOther.Percentage {
		diffs = append(diffs, "Percentage")
	}
	ButtonTogglesDifferent := false
    if len(grouptoogle.ButtonToggles) != len(grouptoogleOther.ButtonToggles) {
        ButtonTogglesDifferent = true
    } else {
        for i := range grouptoogle.ButtonToggles {
            if (grouptoogle.ButtonToggles[i] == nil) != (grouptoogleOther.ButtonToggles[i] == nil) {
                ButtonTogglesDifferent = true
                break
            } else if grouptoogle.ButtonToggles[i] != nil && grouptoogleOther.ButtonToggles[i] != nil {
                if len(grouptoogle.ButtonToggles[i].GongDiff(grouptoogleOther.ButtonToggles[i])) > 0 {
                    ButtonTogglesDifferent = true
                    break
                }
            }
        }
    }
    if ButtonTogglesDifferent {
        diffs = append(diffs, "ButtonToggles")
    }
	if grouptoogle.IsSingleSelector != grouptoogleOther.IsSingleSelector {
		diffs = append(diffs, "IsSingleSelector")
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
	GroupTooglesDifferent := false
    if len(layout.GroupToogles) != len(layoutOther.GroupToogles) {
        GroupTooglesDifferent = true
    } else {
        for i := range layout.GroupToogles {
            if (layout.GroupToogles[i] == nil) != (layoutOther.GroupToogles[i] == nil) {
                GroupTooglesDifferent = true
                break
            } else if layout.GroupToogles[i] != nil && layoutOther.GroupToogles[i] != nil {
                if len(layout.GroupToogles[i].GongDiff(layoutOther.GroupToogles[i])) > 0 {
                    GroupTooglesDifferent = true
                    break
                }
            }
        }
    }
    if GroupTooglesDifferent {
        diffs = append(diffs, "GroupToogles")
    }

	return
}
