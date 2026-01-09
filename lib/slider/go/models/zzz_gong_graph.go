// generated code - do not edit
package models

import "fmt"

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
	SlidersDifferent := false
	if len(group.Sliders) != len(groupOther.Sliders) {
		SlidersDifferent = true
	} else {
		for i := range group.Sliders {
			if (group.Sliders[i] == nil) != (groupOther.Sliders[i] == nil) {
				SlidersDifferent = true
				break
			} else if group.Sliders[i] != nil && groupOther.Sliders[i] != nil {
				// this is a pointer comparaison
				if group.Sliders[i] != groupOther.Sliders[i] {
					SlidersDifferent = true
					break
				}
			}
		}
	}
	if SlidersDifferent {
		ops := Diff(stage, group, groupOther, "Sliders", groupOther.Sliders, group.Sliders)
		diffs = append(diffs, ops)
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
				// this is a pointer comparaison
				if group.Checkboxes[i] != groupOther.Checkboxes[i] {
					CheckboxesDifferent = true
					break
				}
			}
		}
	}
	if CheckboxesDifferent {
		ops := Diff(stage, group, groupOther, "Checkboxes", groupOther.Checkboxes, group.Checkboxes)
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
		ops := Diff(stage, layout, layoutOther, "Groups", layoutOther.Groups, layout.Groups)
		diffs = append(diffs, ops)
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

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
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

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
