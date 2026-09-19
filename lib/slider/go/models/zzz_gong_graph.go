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
func (checkbox *Checkbox) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Checkboxs[checkbox]

	return
}

func (stage *Stage) IsStagedCheckbox(checkbox *Checkbox) (ok bool) {

	return checkbox.GongIsStaged(stage)
}

func (group *Group) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Groups[group]

	return
}

func (stage *Stage) IsStagedGroup(group *Group) (ok bool) {

	return group.GongIsStaged(stage)
}

func (layout *Layout) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Layouts[layout]

	return
}

func (stage *Stage) IsStagedLayout(layout *Layout) (ok bool) {

	return layout.GongIsStaged(stage)
}

func (slider *Slider) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Sliders[slider]

	return
}

func (stage *Stage) IsStagedSlider(slider *Slider) (ok bool) {

	return slider.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (checkbox *Checkbox) GongStageBranch(stage *Stage) {
	stage.StageBranchCheckbox(checkbox)
}

func (stage *Stage) StageBranchCheckbox(checkbox *Checkbox) {

	// check if instance is already staged
	if stage.IsStaged(checkbox) {
		return
	}

	checkbox.Stage(stage)

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
	for _, _slider := range group.Sliders {
		stage.StageBranch(_slider)
	}
	for _, _checkbox := range group.Checkboxes {
		stage.StageBranch(_checkbox)
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

}

func (slider *Slider) GongStageBranch(stage *Stage) {
	stage.StageBranchSlider(slider)
}

func (stage *Stage) StageBranchSlider(slider *Slider) {

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

	// checkboxFrom has already been copied
	if _checkboxTo, ok := mapOrigCopy[checkboxFrom]; ok {
		checkboxTo = _checkboxTo.(*Checkbox)
		return
	}

	checkboxTo = new(Checkbox)
	mapOrigCopy[checkboxFrom] = checkboxTo
	checkboxFrom.GongCopyBasicFields(checkboxTo)

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
	for _, _slider := range groupFrom.Sliders {
		groupTo.Sliders = append(groupTo.Sliders, GongCopyBranchSlider(mapOrigCopy, _slider))
	}
	for _, _checkbox := range groupFrom.Checkboxes {
		groupTo.Checkboxes = append(groupTo.Checkboxes, GongCopyBranchCheckbox(mapOrigCopy, _checkbox))
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

	return
}

func GongCopyBranchSlider(mapOrigCopy map[any]any, sliderFrom *Slider) (sliderTo *Slider) {

	// sliderFrom has already been copied
	if _sliderTo, ok := mapOrigCopy[sliderFrom]; ok {
		sliderTo = _sliderTo.(*Slider)
		return
	}

	sliderTo = new(Slider)
	mapOrigCopy[sliderFrom] = sliderTo
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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (checkbox *Checkbox) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCheckbox(checkbox)
}

func (stage *Stage) UnstageBranchCheckbox(checkbox *Checkbox) {

	// check if instance is already staged
	if !stage.IsStaged(checkbox) {
		return
	}

	checkbox.Unstage(stage)

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
	for _, _slider := range group.Sliders {
		stage.UnstageBranch(_slider)
	}
	for _, _checkbox := range group.Checkboxes {
		stage.UnstageBranch(_checkbox)
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

}

func (slider *Slider) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSlider(slider)
}

func (stage *Stage) UnstageBranchSlider(slider *Slider) {

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
	reference.Sliders = reference.Sliders[:0]
	for _, _b := range instance.Sliders {
		reference.Sliders = append(reference.Sliders, stage.Sliders_reference[_b])
	}
	reference.Checkboxes = reference.Checkboxes[:0]
	for _, _b := range instance.Checkboxes {
		reference.Checkboxes = append(reference.Checkboxes, stage.Checkboxs_reference[_b])
	}
}

func (reference *Layout) GongReconstructPointersFromReferences(stage *Stage, instance *Layout) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Groups = reference.Groups[:0]
	for _, _b := range instance.Groups {
		reference.Groups = append(reference.Groups, stage.Groups_reference[_b])
	}
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
	var _Sliders []*Slider
	for _, _reference := range reference.Sliders {
		if _instance, ok := stage.Sliders_instance[_reference]; ok {
			_Sliders = append(_Sliders, _instance)
		}
	}
	reference.Sliders = _Sliders
	var _Checkboxes []*Checkbox
	for _, _reference := range reference.Checkboxes {
		if _instance, ok := stage.Checkboxs_instance[_reference]; ok {
			_Checkboxes = append(_Checkboxes, _instance)
		}
	}
	reference.Checkboxes = _Checkboxes
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
		ops := stage.Diff(
			group,
			"Sliders",
			len(groupOther.Sliders),
			len(group.Sliders),
			func(i, j int) bool {
				return groupOther.Sliders[i] == group.Sliders[j]
			},
			func(j int) string {
				return group.Sliders[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			group,
			"Checkboxes",
			len(groupOther.Checkboxes),
			len(group.Checkboxes),
			func(i, j int) bool {
				return groupOther.Checkboxes[i] == group.Checkboxes[j]
			},
			func(j int) string {
				return group.Checkboxes[j].GongGetIdentifier(stage)
			},
		)
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
