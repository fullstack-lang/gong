// generated code - do not edit
package barline

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
func (barstylecolor *BarStyleColor) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BarStyleColors[barstylecolor]
	return ok
}

func (barline *Barline) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Barlines[barline]
	return ok
}

func (ending *Ending) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Endings[ending]
	return ok
}

func (repeat *Repeat) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Repeats[repeat]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (barstylecolor *BarStyleColor) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(barstylecolor) {
		return
	}

	barstylecolor.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (barline *Barline) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(barline) {
		return
	}

	barline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if barline.BarStyle != nil {
		stage.StageBranch(barline.BarStyle)
	}
	if barline.Ending != nil {
		stage.StageBranch(barline.Ending)
	}
	if barline.Repeat != nil {
		stage.StageBranch(barline.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ending *Ending) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(ending) {
		return
	}

	ending.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (repeat *Repeat) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(repeat) {
		return
	}

	repeat.Stage(stage)

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
	case *BarStyleColor:
		toT := GongCopyBranchBarStyleColor(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Barline:
		toT := GongCopyBranchBarline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Ending:
		toT := GongCopyBranchEnding(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Repeat:
		toT := GongCopyBranchRepeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchBarStyleColor(mapOrigCopy map[any]any, barstylecolorFrom *BarStyleColor) (barstylecolorTo *BarStyleColor) {
	var alreadyCopied bool
	barstylecolorTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, barstylecolorFrom)
	if alreadyCopied {
		return
	}
	barstylecolorFrom.GongCopyBasicFields(barstylecolorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBarline(mapOrigCopy map[any]any, barlineFrom *Barline) (barlineTo *Barline) {
	var alreadyCopied bool
	barlineTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, barlineFrom)
	if alreadyCopied {
		return
	}
	barlineFrom.GongCopyBasicFields(barlineTo)

	//insertion point for the staging of instances referenced by pointers
	if barlineFrom.BarStyle != nil {
		barlineTo.BarStyle = GongCopyBranchBarStyleColor(mapOrigCopy, barlineFrom.BarStyle)
	}
	if barlineFrom.Ending != nil {
		barlineTo.Ending = GongCopyBranchEnding(mapOrigCopy, barlineFrom.Ending)
	}
	if barlineFrom.Repeat != nil {
		barlineTo.Repeat = GongCopyBranchRepeat(mapOrigCopy, barlineFrom.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEnding(mapOrigCopy map[any]any, endingFrom *Ending) (endingTo *Ending) {
	var alreadyCopied bool
	endingTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, endingFrom)
	if alreadyCopied {
		return
	}
	endingFrom.GongCopyBasicFields(endingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRepeat(mapOrigCopy map[any]any, repeatFrom *Repeat) (repeatTo *Repeat) {
	var alreadyCopied bool
	repeatTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, repeatFrom)
	if alreadyCopied {
		return
	}
	repeatFrom.GongCopyBasicFields(repeatTo)

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
func (barstylecolor *BarStyleColor) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(barstylecolor) {
		return
	}

	barstylecolor.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (barline *Barline) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(barline) {
		return
	}

	barline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if barline.BarStyle != nil {
		stage.UnstageBranch(barline.BarStyle)
	}
	if barline.Ending != nil {
		stage.UnstageBranch(barline.Ending)
	}
	if barline.Repeat != nil {
		stage.UnstageBranch(barline.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (ending *Ending) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(ending) {
		return
	}

	ending.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (repeat *Repeat) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(repeat) {
		return
	}

	repeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *BarStyleColor) GongReconstructPointersFromReferences(stage *Stage, instance *BarStyleColor) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Barline) GongReconstructPointersFromReferences(stage *Stage, instance *Barline) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.BarStyle, stage.BarStyleColors_reference, instance.BarStyle)
	__gong__reconstructPointer(&reference.Ending, stage.Endings_reference, instance.Ending)
	__gong__reconstructPointer(&reference.Repeat, stage.Repeats_reference, instance.Repeat)
	// insertion point for slice of pointers field
}

func (reference *Ending) GongReconstructPointersFromReferences(stage *Stage, instance *Ending) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Repeat) GongReconstructPointersFromReferences(stage *Stage, instance *Repeat) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *BarStyleColor) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Barline) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.BarStyle, stage.BarStyleColors_instance)
	__gong__reconstructPointerFromInstance(&reference.Ending, stage.Endings_instance)
	__gong__reconstructPointerFromInstance(&reference.Repeat, stage.Repeats_instance)
	// insertion point for slice of pointers fields
}

func (reference *Ending) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Repeat) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (barstylecolor *BarStyleColor) GongDiff(stage *Stage, barstylecolorOther *BarStyleColor) (diffs []string) {
	// insertion point for field diffs
	if barstylecolor.Name != barstylecolorOther.Name {
		diffs = append(diffs, barstylecolor.GongMarshallField(stage, "Name"))
	}
	if barstylecolor.Color != barstylecolorOther.Color {
		diffs = append(diffs, barstylecolor.GongMarshallField(stage, "Color"))
	}
	if barstylecolor.Value != barstylecolorOther.Value {
		diffs = append(diffs, barstylecolor.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (barline *Barline) GongDiff(stage *Stage, barlineOther *Barline) (diffs []string) {
	// insertion point for field diffs
	if barline.Name != barlineOther.Name {
		diffs = append(diffs, barline.GongMarshallField(stage, "Name"))
	}
	if barline.Location != barlineOther.Location {
		diffs = append(diffs, barline.GongMarshallField(stage, "Location"))
	}
	if barline.Segno != barlineOther.Segno {
		diffs = append(diffs, barline.GongMarshallField(stage, "Segno"))
	}
	if barline.Coda != barlineOther.Coda {
		diffs = append(diffs, barline.GongMarshallField(stage, "Coda"))
	}
	if barline.Divisions != barlineOther.Divisions {
		diffs = append(diffs, barline.GongMarshallField(stage, "Divisions"))
	}
	if barline.Id != barlineOther.Id {
		diffs = append(diffs, barline.GongMarshallField(stage, "Id"))
	}
	if barline.BarStyle != barlineOther.BarStyle {
		diffs = append(diffs, barline.GongMarshallField(stage, "BarStyle"))
	}
	if barline.Ending != barlineOther.Ending {
		diffs = append(diffs, barline.GongMarshallField(stage, "Ending"))
	}
	if barline.Repeat != barlineOther.Repeat {
		diffs = append(diffs, barline.GongMarshallField(stage, "Repeat"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (ending *Ending) GongDiff(stage *Stage, endingOther *Ending) (diffs []string) {
	// insertion point for field diffs
	if ending.Name != endingOther.Name {
		diffs = append(diffs, ending.GongMarshallField(stage, "Name"))
	}
	if ending.Number != endingOther.Number {
		diffs = append(diffs, ending.GongMarshallField(stage, "Number"))
	}
	if ending.Type != endingOther.Type {
		diffs = append(diffs, ending.GongMarshallField(stage, "Type"))
	}
	if ending.EndLength != endingOther.EndLength {
		diffs = append(diffs, ending.GongMarshallField(stage, "EndLength"))
	}
	if ending.TextX != endingOther.TextX {
		diffs = append(diffs, ending.GongMarshallField(stage, "TextX"))
	}
	if ending.TextY != endingOther.TextY {
		diffs = append(diffs, ending.GongMarshallField(stage, "TextY"))
	}
	if ending.DefaultX != endingOther.DefaultX {
		diffs = append(diffs, ending.GongMarshallField(stage, "DefaultX"))
	}
	if ending.DefaultY != endingOther.DefaultY {
		diffs = append(diffs, ending.GongMarshallField(stage, "DefaultY"))
	}
	if ending.RelativeX != endingOther.RelativeX {
		diffs = append(diffs, ending.GongMarshallField(stage, "RelativeX"))
	}
	if ending.RelativeY != endingOther.RelativeY {
		diffs = append(diffs, ending.GongMarshallField(stage, "RelativeY"))
	}
	if ending.FontFamily != endingOther.FontFamily {
		diffs = append(diffs, ending.GongMarshallField(stage, "FontFamily"))
	}
	if ending.FontSize != endingOther.FontSize {
		diffs = append(diffs, ending.GongMarshallField(stage, "FontSize"))
	}
	if ending.Color != endingOther.Color {
		diffs = append(diffs, ending.GongMarshallField(stage, "Color"))
	}
	if ending.System != endingOther.System {
		diffs = append(diffs, ending.GongMarshallField(stage, "System"))
	}
	if ending.Value != endingOther.Value {
		diffs = append(diffs, ending.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (repeat *Repeat) GongDiff(stage *Stage, repeatOther *Repeat) (diffs []string) {
	// insertion point for field diffs
	if repeat.Name != repeatOther.Name {
		diffs = append(diffs, repeat.GongMarshallField(stage, "Name"))
	}
	if repeat.Direction != repeatOther.Direction {
		diffs = append(diffs, repeat.GongMarshallField(stage, "Direction"))
	}
	if repeat.Times != repeatOther.Times {
		diffs = append(diffs, repeat.GongMarshallField(stage, "Times"))
	}
	if repeat.Winged != repeatOther.Winged {
		diffs = append(diffs, repeat.GongMarshallField(stage, "Winged"))
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
