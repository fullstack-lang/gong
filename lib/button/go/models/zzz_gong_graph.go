// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (button *Button) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Buttons[button]

	return
}

func (stage *Stage) IsStagedButton(button *Button) (ok bool) {

	return button.GongIsStaged(stage)
}

func (buttontoggle *ButtonToggle) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ButtonToggles[buttontoggle]

	return
}

func (stage *Stage) IsStagedButtonToggle(buttontoggle *ButtonToggle) (ok bool) {

	return buttontoggle.GongIsStaged(stage)
}

func (group *Group) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *Stage) IsStagedGroup(group *Group) (ok bool) {

	return group.GongIsStaged(stage)
}

func (grouptoogle *GroupToogle) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GroupToogles[grouptoogle]

	return
}

func (stage *Stage) IsStagedGroupToogle(grouptoogle *GroupToogle) (ok bool) {

	return grouptoogle.GongIsStaged(stage)
}

func (layout *Layout) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Layouts[layout]

	return
}

func (stage *Stage) IsStagedLayout(layout *Layout) (ok bool) {

	return layout.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (button *Button) GongStageBranch(stage *Stage) {
	stage.StageBranchButton(button)
}

func (stage *Stage) StageBranchButton(button *Button) {

	// check if instance is already staged
	if stage.IsStaged(button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (buttontoggle *ButtonToggle) GongStageBranch(stage *Stage) {
	stage.StageBranchButtonToggle(buttontoggle)
}

func (stage *Stage) StageBranchButtonToggle(buttontoggle *ButtonToggle) {

	// check if instance is already staged
	if stage.IsStaged(buttontoggle) {
		return
	}

	buttontoggle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group *Group) GongStageBranch(stage *Stage) {
	stage.StageBranchGroup(group)
}

func (stage *Stage) StageBranchGroup(group *Group) {

	// check if instance is already staged
	if stage.IsStaged(group) {
		return
	}

	group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _button := range group.Buttons {
		stage.StageBranch(_button)
	}

}

func (grouptoogle *GroupToogle) GongStageBranch(stage *Stage) {
	stage.StageBranchGroupToogle(grouptoogle)
}

func (stage *Stage) StageBranchGroupToogle(grouptoogle *GroupToogle) {

	// check if instance is already staged
	if stage.IsStaged(grouptoogle) {
		return
	}

	grouptoogle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _buttontoggle := range grouptoogle.ButtonToggles {
		stage.StageBranch(_buttontoggle)
	}

}

func (layout *Layout) GongStageBranch(stage *Stage) {
	stage.StageBranchLayout(layout)
}

func (stage *Stage) StageBranchLayout(layout *Layout) {

	// check if instance is already staged
	if stage.IsStaged(layout) {
		return
	}

	layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range layout.Groups {
		stage.StageBranch(_group)
	}
	for _, _grouptoogle := range layout.GroupToogles {
		stage.StageBranch(_grouptoogle)
	}

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Button:
		toT := GongCopyBranchButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ButtonToggle:
		toT := GongCopyBranchButtonToggle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := GongCopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GroupToogle:
		toT := GongCopyBranchGroupToogle(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Layout:
		toT := GongCopyBranchLayout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchButton(mapOrigCopy map[any]any, buttonFrom *Button) (buttonTo *Button) {

	// buttonFrom has already been copied
	if _buttonTo, ok := mapOrigCopy[buttonFrom]; ok {
		buttonTo = _buttonTo.(*Button)
		return
	}

	buttonTo = new(Button)
	mapOrigCopy[buttonFrom] = buttonTo
	buttonFrom.GongCopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchButtonToggle(mapOrigCopy map[any]any, buttontoggleFrom *ButtonToggle) (buttontoggleTo *ButtonToggle) {

	// buttontoggleFrom has already been copied
	if _buttontoggleTo, ok := mapOrigCopy[buttontoggleFrom]; ok {
		buttontoggleTo = _buttontoggleTo.(*ButtonToggle)
		return
	}

	buttontoggleTo = new(ButtonToggle)
	mapOrigCopy[buttontoggleFrom] = buttontoggleTo
	buttontoggleFrom.GongCopyBasicFields(buttontoggleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {

	// groupFrom has already been copied
	if _groupTo, ok := mapOrigCopy[groupFrom]; ok {
		groupTo = _groupTo.(*Group)
		return
	}

	groupTo = new(Group)
	mapOrigCopy[groupFrom] = groupTo
	groupFrom.GongCopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _button := range groupFrom.Buttons {
		groupTo.Buttons = append(groupTo.Buttons, GongCopyBranchButton(mapOrigCopy, _button))
	}

	return
}

func GongCopyBranchGroupToogle(mapOrigCopy map[any]any, grouptoogleFrom *GroupToogle) (grouptoogleTo *GroupToogle) {

	// grouptoogleFrom has already been copied
	if _grouptoogleTo, ok := mapOrigCopy[grouptoogleFrom]; ok {
		grouptoogleTo = _grouptoogleTo.(*GroupToogle)
		return
	}

	grouptoogleTo = new(GroupToogle)
	mapOrigCopy[grouptoogleFrom] = grouptoogleTo
	grouptoogleFrom.GongCopyBasicFields(grouptoogleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _buttontoggle := range grouptoogleFrom.ButtonToggles {
		grouptoogleTo.ButtonToggles = append(grouptoogleTo.ButtonToggles, GongCopyBranchButtonToggle(mapOrigCopy, _buttontoggle))
	}

	return
}

func GongCopyBranchLayout(mapOrigCopy map[any]any, layoutFrom *Layout) (layoutTo *Layout) {

	// layoutFrom has already been copied
	if _layoutTo, ok := mapOrigCopy[layoutFrom]; ok {
		layoutTo = _layoutTo.(*Layout)
		return
	}

	layoutTo = new(Layout)
	mapOrigCopy[layoutFrom] = layoutTo
	layoutFrom.GongCopyBasicFields(layoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range layoutFrom.Groups {
		layoutTo.Groups = append(layoutTo.Groups, GongCopyBranchGroup(mapOrigCopy, _group))
	}
	for _, _grouptoogle := range layoutFrom.GroupToogles {
		layoutTo.GroupToogles = append(layoutTo.GroupToogles, GongCopyBranchGroupToogle(mapOrigCopy, _grouptoogle))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (button *Button) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchButton(button)
}

func (stage *Stage) UnstageBranchButton(button *Button) {

	// check if instance is already staged
	if !stage.IsStaged(button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (buttontoggle *ButtonToggle) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchButtonToggle(buttontoggle)
}

func (stage *Stage) UnstageBranchButtonToggle(buttontoggle *ButtonToggle) {

	// check if instance is already staged
	if !stage.IsStaged(buttontoggle) {
		return
	}

	buttontoggle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group *Group) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGroup(group)
}

func (stage *Stage) UnstageBranchGroup(group *Group) {

	// check if instance is already staged
	if !stage.IsStaged(group) {
		return
	}

	group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _button := range group.Buttons {
		stage.UnstageBranch(_button)
	}

}

func (grouptoogle *GroupToogle) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGroupToogle(grouptoogle)
}

func (stage *Stage) UnstageBranchGroupToogle(grouptoogle *GroupToogle) {

	// check if instance is already staged
	if !stage.IsStaged(grouptoogle) {
		return
	}

	grouptoogle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _buttontoggle := range grouptoogle.ButtonToggles {
		stage.UnstageBranch(_buttontoggle)
	}

}

func (layout *Layout) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLayout(layout)
}

func (stage *Stage) UnstageBranchLayout(layout *Layout) {

	// check if instance is already staged
	if !stage.IsStaged(layout) {
		return
	}

	layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _group := range layout.Groups {
		stage.UnstageBranch(_group)
	}
	for _, _grouptoogle := range layout.GroupToogles {
		stage.UnstageBranch(_grouptoogle)
	}

}

// insertion point for pointer reconstruction from references
func (reference *Button) GongReconstructPointersFromReferences(stage *Stage, instance *Button) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ButtonToggle) GongReconstructPointersFromReferences(stage *Stage, instance *ButtonToggle) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Group) GongReconstructPointersFromReferences(stage *Stage, instance *Group) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Buttons = reference.Buttons[:0]
	for _, _b := range instance.Buttons {
		reference.Buttons = append(reference.Buttons, stage.Buttons_reference[_b])
	}
}

func (reference *GroupToogle) GongReconstructPointersFromReferences(stage *Stage, instance *GroupToogle) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ButtonToggles = reference.ButtonToggles[:0]
	for _, _b := range instance.ButtonToggles {
		reference.ButtonToggles = append(reference.ButtonToggles, stage.ButtonToggles_reference[_b])
	}
}

func (reference *Layout) GongReconstructPointersFromReferences(stage *Stage, instance *Layout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
	reference.GroupToogles = reference.GroupToogles[:0]
	for _, _b := range instance.GroupToogles {
		reference.GroupToogles = append(reference.GroupToogles, stage.GroupToogles_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *Button) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ButtonToggle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Buttons []*Button
	for _, _reference := range reference.Buttons {
		if _instance, ok := stage.Buttons_instance[_reference]; ok {
			_Buttons = append(_Buttons, _instance)
		}
	}
	reference.Buttons = _Buttons
}

func (reference *GroupToogle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ButtonToggles []*ButtonToggle
	for _, _reference := range reference.ButtonToggles {
		if _instance, ok := stage.ButtonToggles_instance[_reference]; ok {
			_ButtonToggles = append(_ButtonToggles, _instance)
		}
	}
	reference.ButtonToggles = _ButtonToggles
}

func (reference *Layout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Groups []*Group
	for _, _reference := range reference.Groups {
		if _instance, ok := stage.Groups_instance[_reference]; ok {
			_Groups = append(_Groups, _instance)
		}
	}
	reference.Groups = _Groups
	var _GroupToogles []*GroupToogle
	for _, _reference := range reference.GroupToogles {
		if _instance, ok := stage.GroupToogles_instance[_reference]; ok {
			_GroupToogles = append(_GroupToogles, _instance)
		}
	}
	reference.GroupToogles = _GroupToogles
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (button *Button) GongDiff(stage *Stage, buttonOther *Button) (diffs []string) {
	// insertion point for field diffs
	if button.Name != buttonOther.Name {
		diffs = append(diffs, button.GongMarshallField(stage, "Name"))
	}
	if button.Label != buttonOther.Label {
		diffs = append(diffs, button.GongMarshallField(stage, "Label"))
	}
	if button.Icon != buttonOther.Icon {
		diffs = append(diffs, button.GongMarshallField(stage, "Icon"))
	}
	if button.IsDisabled != buttonOther.IsDisabled {
		diffs = append(diffs, button.GongMarshallField(stage, "IsDisabled"))
	}
	if button.Color != buttonOther.Color {
		diffs = append(diffs, button.GongMarshallField(stage, "Color"))
	}
	if button.MatButtonType != buttonOther.MatButtonType {
		diffs = append(diffs, button.GongMarshallField(stage, "MatButtonType"))
	}
	if button.MatButtonAppearance != buttonOther.MatButtonAppearance {
		diffs = append(diffs, button.GongMarshallField(stage, "MatButtonAppearance"))
	}
	if button.HasToolTip != buttonOther.HasToolTip {
		diffs = append(diffs, button.GongMarshallField(stage, "HasToolTip"))
	}
	if button.ToolTipText != buttonOther.ToolTipText {
		diffs = append(diffs, button.GongMarshallField(stage, "ToolTipText"))
	}
	if button.ToolTipPosition != buttonOther.ToolTipPosition {
		diffs = append(diffs, button.GongMarshallField(stage, "ToolTipPosition"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (buttontoggle *ButtonToggle) GongDiff(stage *Stage, buttontoggleOther *ButtonToggle) (diffs []string) {
	// insertion point for field diffs
	if buttontoggle.Name != buttontoggleOther.Name {
		diffs = append(diffs, buttontoggle.GongMarshallField(stage, "Name"))
	}
	if buttontoggle.Label != buttontoggleOther.Label {
		diffs = append(diffs, buttontoggle.GongMarshallField(stage, "Label"))
	}
	if buttontoggle.Icon != buttontoggleOther.Icon {
		diffs = append(diffs, buttontoggle.GongMarshallField(stage, "Icon"))
	}
	if buttontoggle.IsDisabled != buttontoggleOther.IsDisabled {
		diffs = append(diffs, buttontoggle.GongMarshallField(stage, "IsDisabled"))
	}
	if buttontoggle.IsChecked != buttontoggleOther.IsChecked {
		diffs = append(diffs, buttontoggle.GongMarshallField(stage, "IsChecked"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (group *Group) GongDiff(stage *Stage, groupOther *Group) (diffs []string) {
	// insertion point for field diffs
	if group.Name != groupOther.Name {
		diffs = append(diffs, group.GongMarshallField(stage, "Name"))
	}
	if group.Percentage != groupOther.Percentage {
		diffs = append(diffs, group.GongMarshallField(stage, "Percentage"))
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
				// this is a pointer comparaison
				if group.Buttons[i] != groupOther.Buttons[i] {
					ButtonsDifferent = true
					break
				}
			}
		}
	}
	if ButtonsDifferent {
		ops := stage.Diff(
			group,
			"Buttons",
			len(groupOther.Buttons),
			len(group.Buttons),
			func(i, j int) bool {
				return groupOther.Buttons[i] == group.Buttons[j]
			},
			func(j int) string {
				return group.Buttons[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if group.NbColumns != groupOther.NbColumns {
		diffs = append(diffs, group.GongMarshallField(stage, "NbColumns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (grouptoogle *GroupToogle) GongDiff(stage *Stage, grouptoogleOther *GroupToogle) (diffs []string) {
	// insertion point for field diffs
	if grouptoogle.Name != grouptoogleOther.Name {
		diffs = append(diffs, grouptoogle.GongMarshallField(stage, "Name"))
	}
	if grouptoogle.Percentage != grouptoogleOther.Percentage {
		diffs = append(diffs, grouptoogle.GongMarshallField(stage, "Percentage"))
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
				// this is a pointer comparaison
				if grouptoogle.ButtonToggles[i] != grouptoogleOther.ButtonToggles[i] {
					ButtonTogglesDifferent = true
					break
				}
			}
		}
	}
	if ButtonTogglesDifferent {
		ops := stage.Diff(
			grouptoogle,
			"ButtonToggles",
			len(grouptoogleOther.ButtonToggles),
			len(grouptoogle.ButtonToggles),
			func(i, j int) bool {
				return grouptoogleOther.ButtonToggles[i] == grouptoogle.ButtonToggles[j]
			},
			func(j int) string {
				return grouptoogle.ButtonToggles[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if grouptoogle.IsSingleSelector != grouptoogleOther.IsSingleSelector {
		diffs = append(diffs, grouptoogle.GongMarshallField(stage, "IsSingleSelector"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (layout *Layout) GongDiff(stage *Stage, layoutOther *Layout) (diffs []string) {
	// insertion point for field diffs
	if layout.Name != layoutOther.Name {
		diffs = append(diffs, layout.GongMarshallField(stage, "Name"))
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
				// this is a pointer comparaison
				if layout.Groups[i] != layoutOther.Groups[i] {
					GroupsDifferent = true
					break
				}
			}
		}
	}
	if GroupsDifferent {
		ops := stage.Diff(
			layout,
			"Groups",
			len(layoutOther.Groups),
			len(layout.Groups),
			func(i, j int) bool {
				return layoutOther.Groups[i] == layout.Groups[j]
			},
			func(j int) string {
				return layout.Groups[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
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
				// this is a pointer comparaison
				if layout.GroupToogles[i] != layoutOther.GroupToogles[i] {
					GroupTooglesDifferent = true
					break
				}
			}
		}
	}
	if GroupTooglesDifferent {
		ops := stage.Diff(
			layout,
			"GroupToogles",
			len(layoutOther.GroupToogles),
			len(layout.GroupToogles),
			func(i, j int) bool {
				return layoutOther.GroupToogles[i] == layout.GroupToogles[j]
			},
			func(j int) string {
				return layout.GroupToogles[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
