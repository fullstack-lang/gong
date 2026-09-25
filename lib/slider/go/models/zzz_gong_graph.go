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
func (checkbox *Checkbox) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Checkboxs[checkbox]
	return ok
}

func (group *Group) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Groups[group]
	return ok
}

func (layout *Layout) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Layouts[layout]
	return ok
}

func (slider *Slider) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Sliders[slider]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (checkbox *Checkbox) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(checkbox) {
		return
	}

	checkbox.Stage(stage)

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
	for _, _slider := range group.Sliders {
		stage.StageBranch(_slider)
	}
	for _, _checkbox := range group.Checkboxes {
		stage.StageBranch(_checkbox)
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

}

func (slider *Slider) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(slider) {
		return
	}

	slider.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
	case *Checkbox:
		toT := GongCopyBranchCheckbox(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group:
		toT := GongCopyBranchGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Layout:
		toT := GongCopyBranchLayout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slider:
		toT := GongCopyBranchSlider(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCheckbox(mapOrigCopy map[any]any, checkboxFrom *Checkbox) (checkboxTo *Checkbox) {
	var alreadyCopied bool
	checkboxTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, checkboxFrom)
	if alreadyCopied {
		return
	}
	checkboxFrom.GongCopyBasicFields(checkboxTo)

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
	for _, _slider := range groupFrom.Sliders {
		groupTo.Sliders = append(groupTo.Sliders, GongCopyBranchSlider(mapOrigCopy, _slider))
	}
	for _, _checkbox := range groupFrom.Checkboxes {
		groupTo.Checkboxes = append(groupTo.Checkboxes, GongCopyBranchCheckbox(mapOrigCopy, _checkbox))
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

	return
}

func GongCopyBranchSlider(mapOrigCopy map[any]any, sliderFrom *Slider) (sliderTo *Slider) {
	var alreadyCopied bool
	sliderTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, sliderFrom)
	if alreadyCopied {
		return
	}
	sliderFrom.GongCopyBasicFields(sliderTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
func (checkbox *Checkbox) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(checkbox) {
		return
	}

	checkbox.Unstage(stage)

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
	for _, _slider := range group.Sliders {
		stage.UnstageBranch(_slider)
	}
	for _, _checkbox := range group.Checkboxes {
		stage.UnstageBranch(_checkbox)
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

}

func (slider *Slider) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(slider) {
		return
	}

	slider.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Checkbox) GongReconstructPointersFromReferences(stage *Stage, instance *Checkbox) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Group) GongReconstructPointersFromReferences(stage *Stage, instance *Group) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sliders, stage.Sliders_reference, instance.Sliders)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Checkboxes, stage.Checkboxs_reference, instance.Checkboxes)
}

func (reference *Layout) GongReconstructPointersFromReferences(stage *Stage, instance *Layout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Groups, stage.Groups_reference, instance.Groups)
}

func (reference *Slider) GongReconstructPointersFromReferences(stage *Stage, instance *Slider) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Checkbox) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Group) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sliders, stage.Sliders_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Checkboxes, stage.Checkboxs_instance)
}

func (reference *Layout) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Groups, stage.Groups_instance)
}

func (reference *Slider) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (checkbox *Checkbox) GongDiff(stage *Stage, checkboxOther *Checkbox) (diffs []string) {
	// insertion point for field diffs
	if checkbox.Name != checkboxOther.Name {
		diffs = append(diffs, checkbox.GongMarshallField(stage, "Name"))
	}
	if checkbox.ValueBool != checkboxOther.ValueBool {
		diffs = append(diffs, checkbox.GongMarshallField(stage, "ValueBool"))
	}
	if checkbox.LabelForTrue != checkboxOther.LabelForTrue {
		diffs = append(diffs, checkbox.GongMarshallField(stage, "LabelForTrue"))
	}
	if checkbox.LabelForFalse != checkboxOther.LabelForFalse {
		diffs = append(diffs, checkbox.GongMarshallField(stage, "LabelForFalse"))
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
	if ops := __gong__diffSliceOfPointers(stage, group, "Sliders", groupOther.Sliders, group.Sliders); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, group, "Checkboxes", groupOther.Checkboxes, group.Checkboxes); ops != "" {
		diffs = append(diffs, ops)
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
	if layout.IsWithCustomGutterSize != layoutOther.IsWithCustomGutterSize {
		diffs = append(diffs, layout.GongMarshallField(stage, "IsWithCustomGutterSize"))
	}
	if layout.GutterSize != layoutOther.GutterSize {
		diffs = append(diffs, layout.GongMarshallField(stage, "GutterSize"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (slider *Slider) GongDiff(stage *Stage, sliderOther *Slider) (diffs []string) {
	// insertion point for field diffs
	if slider.Name != sliderOther.Name {
		diffs = append(diffs, slider.GongMarshallField(stage, "Name"))
	}
	if slider.IsFloat64 != sliderOther.IsFloat64 {
		diffs = append(diffs, slider.GongMarshallField(stage, "IsFloat64"))
	}
	if slider.IsInt != sliderOther.IsInt {
		diffs = append(diffs, slider.GongMarshallField(stage, "IsInt"))
	}
	if slider.MinInt != sliderOther.MinInt {
		diffs = append(diffs, slider.GongMarshallField(stage, "MinInt"))
	}
	if slider.MaxInt != sliderOther.MaxInt {
		diffs = append(diffs, slider.GongMarshallField(stage, "MaxInt"))
	}
	if slider.StepInt != sliderOther.StepInt {
		diffs = append(diffs, slider.GongMarshallField(stage, "StepInt"))
	}
	if slider.ValueInt != sliderOther.ValueInt {
		diffs = append(diffs, slider.GongMarshallField(stage, "ValueInt"))
	}
	if slider.MinFloat64 != sliderOther.MinFloat64 {
		diffs = append(diffs, slider.GongMarshallField(stage, "MinFloat64"))
	}
	if slider.MaxFloat64 != sliderOther.MaxFloat64 {
		diffs = append(diffs, slider.GongMarshallField(stage, "MaxFloat64"))
	}
	if slider.StepFloat64 != sliderOther.StepFloat64 {
		diffs = append(diffs, slider.GongMarshallField(stage, "StepFloat64"))
	}
	if slider.ValueFloat64 != sliderOther.ValueFloat64 {
		diffs = append(diffs, slider.GongMarshallField(stage, "ValueFloat64"))
	}
	if slider.IsDisabled != sliderOther.IsDisabled {
		diffs = append(diffs, slider.GongMarshallField(stage, "IsDisabled"))
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
