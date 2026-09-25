// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (button *Button) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Buttons[button]
	return ok
}

func (buttontoggle *ButtonToggle) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ButtonToggles[buttontoggle]
	return ok
}

func (group *Group) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Groups[group]
	return ok
}

func (grouptoogle *GroupToogle) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GroupToogles[grouptoogle]
	return ok
}

func (layout *Layout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Layouts[layout]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (button *Button) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (buttontoggle *ButtonToggle) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(buttontoggle) {
		return
	}

	buttontoggle.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group *Group) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	buttonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, buttonFrom)
	if alreadyCopied {
		return
	}
	buttonFrom.GongCopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchButtonToggle(mapOrigCopy map[any]any, buttontoggleFrom *ButtonToggle) (buttontoggleTo *ButtonToggle) {
	var alreadyCopied bool
	buttontoggleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, buttontoggleFrom)
	if alreadyCopied {
		return
	}
	buttontoggleFrom.GongCopyBasicFields(buttontoggleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGroup(mapOrigCopy map[any]any, groupFrom *Group) (groupTo *Group) {
	var alreadyCopied bool
	groupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, groupFrom)
	if alreadyCopied {
		return
	}
	groupFrom.GongCopyBasicFields(groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _button := range groupFrom.Buttons {
		groupTo.Buttons = append(groupTo.Buttons, GongCopyBranchButton(mapOrigCopy, _button))
	}

	return
}

func GongCopyBranchGroupToogle(mapOrigCopy map[any]any, grouptoogleFrom *GroupToogle) (grouptoogleTo *GroupToogle) {
	var alreadyCopied bool
	grouptoogleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, grouptoogleFrom)
	if alreadyCopied {
		return
	}
	grouptoogleFrom.GongCopyBasicFields(grouptoogleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _buttontoggle := range grouptoogleFrom.ButtonToggles {
		grouptoogleTo.ButtonToggles = append(grouptoogleTo.ButtonToggles, GongCopyBranchButtonToggle(mapOrigCopy, _buttontoggle))
	}

	return
}

func GongCopyBranchLayout(mapOrigCopy map[any]any, layoutFrom *Layout) (layoutTo *Layout) {
	var alreadyCopied bool
	layoutTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, layoutFrom)
	if alreadyCopied {
		return
	}
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

	// check if instance is already staged
	if !stage.IsStaged(button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (buttontoggle *ButtonToggle) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(buttontoggle) {
		return
	}

	buttontoggle.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (group *Group) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructSliceOfPointersFromReferences(&reference.Buttons, stage.Buttons_reference, instance.Buttons)
}

func (reference *GroupToogle) GongReconstructPointersFromReferences(stage *Stage, instance *GroupToogle) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ButtonToggles, stage.ButtonToggles_reference, instance.ButtonToggles)
}

func (reference *Layout) GongReconstructPointersFromReferences(stage *Stage, instance *Layout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GroupToogles, stage.GroupToogles_reference, instance.GroupToogles)
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.Buttons, stage.Buttons_instance)
}

func (reference *GroupToogle) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ButtonToggles, stage.ButtonToggles_instance)
}

func (reference *Layout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GroupToogles, stage.GroupToogles_instance)
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
	if ops := __gong__diffSliceOfPointers(stage, group, "Buttons", groupOther.Buttons, group.Buttons); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, grouptoogle, "ButtonToggles", grouptoogleOther.ButtonToggles, grouptoogle.ButtonToggles); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, layout, "Groups", layoutOther.Groups, layout.Groups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, layout, "GroupToogles", layoutOther.GroupToogles, layout.GroupToogles); ops != "" {
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

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
