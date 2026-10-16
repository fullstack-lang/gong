// generated code - do not edit
package xlink

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
func (arctype *ArcType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArcTypes[arctype]
	return ok
}

func (extended *Extended) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Extendeds[extended]
	return ok
}

func (locatortype *LocatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LocatorTypes[locatortype]
	return ok
}

func (resourcetype *ResourceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ResourceTypes[resourcetype]
	return ok
}

func (simple *Simple) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Simples[simple]
	return ok
}

func (titleelttype *TitleEltType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TitleEltTypes[titleelttype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (arctype *ArcType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(arctype) {
		return
	}

	arctype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range arctype.Title_1 {
		stage.StageBranch(_titleelttype)
	}

}

func (extended *Extended) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(extended) {
		return
	}

	extended.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range extended.Title_1 {
		stage.StageBranch(_titleelttype)
	}
	for _, _resourcetype := range extended.Resource {
		stage.StageBranch(_resourcetype)
	}
	for _, _locatortype := range extended.Locator {
		stage.StageBranch(_locatortype)
	}
	for _, _arctype := range extended.Arc {
		stage.StageBranch(_arctype)
	}

}

func (locatortype *LocatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(locatortype) {
		return
	}

	locatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range locatortype.Title_1 {
		stage.StageBranch(_titleelttype)
	}

}

func (resourcetype *ResourceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(resourcetype) {
		return
	}

	resourcetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (simple *Simple) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(simple) {
		return
	}

	simple.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (titleelttype *TitleEltType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(titleelttype) {
		return
	}

	titleelttype.Stage(stage)

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
	case *ArcType:
		toT := GongCopyBranchArcType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Extended:
		toT := GongCopyBranchExtended(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LocatorType:
		toT := GongCopyBranchLocatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ResourceType:
		toT := GongCopyBranchResourceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Simple:
		toT := GongCopyBranchSimple(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TitleEltType:
		toT := GongCopyBranchTitleEltType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchArcType(mapOrigCopy map[any]any, arctypeFrom *ArcType) (arctypeTo *ArcType) {
	var alreadyCopied bool
	arctypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, arctypeFrom)
	if alreadyCopied {
		return
	}
	arctypeFrom.GongCopyBasicFields(arctypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range arctypeFrom.Title_1 {
		arctypeTo.Title_1 = append(arctypeTo.Title_1, GongCopyBranchTitleEltType(mapOrigCopy, _titleelttype))
	}

	return
}

func GongCopyBranchExtended(mapOrigCopy map[any]any, extendedFrom *Extended) (extendedTo *Extended) {
	var alreadyCopied bool
	extendedTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, extendedFrom)
	if alreadyCopied {
		return
	}
	extendedFrom.GongCopyBasicFields(extendedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range extendedFrom.Title_1 {
		extendedTo.Title_1 = append(extendedTo.Title_1, GongCopyBranchTitleEltType(mapOrigCopy, _titleelttype))
	}
	for _, _resourcetype := range extendedFrom.Resource {
		extendedTo.Resource = append(extendedTo.Resource, GongCopyBranchResourceType(mapOrigCopy, _resourcetype))
	}
	for _, _locatortype := range extendedFrom.Locator {
		extendedTo.Locator = append(extendedTo.Locator, GongCopyBranchLocatorType(mapOrigCopy, _locatortype))
	}
	for _, _arctype := range extendedFrom.Arc {
		extendedTo.Arc = append(extendedTo.Arc, GongCopyBranchArcType(mapOrigCopy, _arctype))
	}

	return
}

func GongCopyBranchLocatorType(mapOrigCopy map[any]any, locatortypeFrom *LocatorType) (locatortypeTo *LocatorType) {
	var alreadyCopied bool
	locatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, locatortypeFrom)
	if alreadyCopied {
		return
	}
	locatortypeFrom.GongCopyBasicFields(locatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range locatortypeFrom.Title_1 {
		locatortypeTo.Title_1 = append(locatortypeTo.Title_1, GongCopyBranchTitleEltType(mapOrigCopy, _titleelttype))
	}

	return
}

func GongCopyBranchResourceType(mapOrigCopy map[any]any, resourcetypeFrom *ResourceType) (resourcetypeTo *ResourceType) {
	var alreadyCopied bool
	resourcetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, resourcetypeFrom)
	if alreadyCopied {
		return
	}
	resourcetypeFrom.GongCopyBasicFields(resourcetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSimple(mapOrigCopy map[any]any, simpleFrom *Simple) (simpleTo *Simple) {
	var alreadyCopied bool
	simpleTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, simpleFrom)
	if alreadyCopied {
		return
	}
	simpleFrom.GongCopyBasicFields(simpleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTitleEltType(mapOrigCopy map[any]any, titleelttypeFrom *TitleEltType) (titleelttypeTo *TitleEltType) {
	var alreadyCopied bool
	titleelttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, titleelttypeFrom)
	if alreadyCopied {
		return
	}
	titleelttypeFrom.GongCopyBasicFields(titleelttypeTo)

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
func (arctype *ArcType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(arctype) {
		return
	}

	arctype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range arctype.Title_1 {
		stage.UnstageBranch(_titleelttype)
	}

}

func (extended *Extended) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(extended) {
		return
	}

	extended.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range extended.Title_1 {
		stage.UnstageBranch(_titleelttype)
	}
	for _, _resourcetype := range extended.Resource {
		stage.UnstageBranch(_resourcetype)
	}
	for _, _locatortype := range extended.Locator {
		stage.UnstageBranch(_locatortype)
	}
	for _, _arctype := range extended.Arc {
		stage.UnstageBranch(_arctype)
	}

}

func (locatortype *LocatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(locatortype) {
		return
	}

	locatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _titleelttype := range locatortype.Title_1 {
		stage.UnstageBranch(_titleelttype)
	}

}

func (resourcetype *ResourceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(resourcetype) {
		return
	}

	resourcetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (simple *Simple) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(simple) {
		return
	}

	simple.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (titleelttype *TitleEltType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(titleelttype) {
		return
	}

	titleelttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *ArcType) GongReconstructPointersFromReferences(stage *Stage, instance *ArcType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Title_1, stage.TitleEltTypes_reference, instance.Title_1)
}

func (reference *Extended) GongReconstructPointersFromReferences(stage *Stage, instance *Extended) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Title_1, stage.TitleEltTypes_reference, instance.Title_1)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Resource, stage.ResourceTypes_reference, instance.Resource)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Locator, stage.LocatorTypes_reference, instance.Locator)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Arc, stage.ArcTypes_reference, instance.Arc)
}

func (reference *LocatorType) GongReconstructPointersFromReferences(stage *Stage, instance *LocatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Title_1, stage.TitleEltTypes_reference, instance.Title_1)
}

func (reference *ResourceType) GongReconstructPointersFromReferences(stage *Stage, instance *ResourceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Simple) GongReconstructPointersFromReferences(stage *Stage, instance *Simple) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TitleEltType) GongReconstructPointersFromReferences(stage *Stage, instance *TitleEltType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *ArcType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Title_1, stage.TitleEltTypes_instance)
}

func (reference *Extended) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Title_1, stage.TitleEltTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Resource, stage.ResourceTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Locator, stage.LocatorTypes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Arc, stage.ArcTypes_instance)
}

func (reference *LocatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Title_1, stage.TitleEltTypes_instance)
}

func (reference *ResourceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Simple) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TitleEltType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (arctype *ArcType) GongDiff(stage *Stage, arctypeOther *ArcType) (diffs []string) {
	// insertion point for field diffs
	if arctype.Name != arctypeOther.Name {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Name"))
	}
	if arctype.Type != arctypeOther.Type {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Type"))
	}
	if arctype.Arcrole != arctypeOther.Arcrole {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Arcrole"))
	}
	if arctype.Title != arctypeOther.Title {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Title"))
	}
	if arctype.Show != arctypeOther.Show {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Show"))
	}
	if arctype.Actuate != arctypeOther.Actuate {
		diffs = append(diffs, arctype.GongMarshallField(stage, "Actuate"))
	}
	if arctype.From != arctypeOther.From {
		diffs = append(diffs, arctype.GongMarshallField(stage, "From"))
	}
	if arctype.To != arctypeOther.To {
		diffs = append(diffs, arctype.GongMarshallField(stage, "To"))
	}
	if ops := __gong__diffSliceOfPointers(stage, arctype, "Title_1", arctypeOther.Title_1, arctype.Title_1); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (extended *Extended) GongDiff(stage *Stage, extendedOther *Extended) (diffs []string) {
	// insertion point for field diffs
	if extended.Name != extendedOther.Name {
		diffs = append(diffs, extended.GongMarshallField(stage, "Name"))
	}
	if extended.Type != extendedOther.Type {
		diffs = append(diffs, extended.GongMarshallField(stage, "Type"))
	}
	if extended.Role != extendedOther.Role {
		diffs = append(diffs, extended.GongMarshallField(stage, "Role"))
	}
	if extended.Title != extendedOther.Title {
		diffs = append(diffs, extended.GongMarshallField(stage, "Title"))
	}
	if ops := __gong__diffSliceOfPointers(stage, extended, "Title_1", extendedOther.Title_1, extended.Title_1); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, extended, "Resource", extendedOther.Resource, extended.Resource); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, extended, "Locator", extendedOther.Locator, extended.Locator); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, extended, "Arc", extendedOther.Arc, extended.Arc); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (locatortype *LocatorType) GongDiff(stage *Stage, locatortypeOther *LocatorType) (diffs []string) {
	// insertion point for field diffs
	if locatortype.Name != locatortypeOther.Name {
		diffs = append(diffs, locatortype.GongMarshallField(stage, "Name"))
	}
	if locatortype.Type != locatortypeOther.Type {
		diffs = append(diffs, locatortype.GongMarshallField(stage, "Type"))
	}
	if locatortype.Href != locatortypeOther.Href {
		diffs = append(diffs, locatortype.GongMarshallField(stage, "Href"))
	}
	if locatortype.Role != locatortypeOther.Role {
		diffs = append(diffs, locatortype.GongMarshallField(stage, "Role"))
	}
	if locatortype.Title != locatortypeOther.Title {
		diffs = append(diffs, locatortype.GongMarshallField(stage, "Title"))
	}
	if locatortype.Label != locatortypeOther.Label {
		diffs = append(diffs, locatortype.GongMarshallField(stage, "Label"))
	}
	if ops := __gong__diffSliceOfPointers(stage, locatortype, "Title_1", locatortypeOther.Title_1, locatortype.Title_1); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (resourcetype *ResourceType) GongDiff(stage *Stage, resourcetypeOther *ResourceType) (diffs []string) {
	// insertion point for field diffs
	if resourcetype.Name != resourcetypeOther.Name {
		diffs = append(diffs, resourcetype.GongMarshallField(stage, "Name"))
	}
	if resourcetype.Type != resourcetypeOther.Type {
		diffs = append(diffs, resourcetype.GongMarshallField(stage, "Type"))
	}
	if resourcetype.Role != resourcetypeOther.Role {
		diffs = append(diffs, resourcetype.GongMarshallField(stage, "Role"))
	}
	if resourcetype.Title != resourcetypeOther.Title {
		diffs = append(diffs, resourcetype.GongMarshallField(stage, "Title"))
	}
	if resourcetype.Label != resourcetypeOther.Label {
		diffs = append(diffs, resourcetype.GongMarshallField(stage, "Label"))
	}
	if resourcetype.Value != resourcetypeOther.Value {
		diffs = append(diffs, resourcetype.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (simple *Simple) GongDiff(stage *Stage, simpleOther *Simple) (diffs []string) {
	// insertion point for field diffs
	if simple.Name != simpleOther.Name {
		diffs = append(diffs, simple.GongMarshallField(stage, "Name"))
	}
	if simple.Type != simpleOther.Type {
		diffs = append(diffs, simple.GongMarshallField(stage, "Type"))
	}
	if simple.Href != simpleOther.Href {
		diffs = append(diffs, simple.GongMarshallField(stage, "Href"))
	}
	if simple.Role != simpleOther.Role {
		diffs = append(diffs, simple.GongMarshallField(stage, "Role"))
	}
	if simple.Arcrole != simpleOther.Arcrole {
		diffs = append(diffs, simple.GongMarshallField(stage, "Arcrole"))
	}
	if simple.Title != simpleOther.Title {
		diffs = append(diffs, simple.GongMarshallField(stage, "Title"))
	}
	if simple.Show != simpleOther.Show {
		diffs = append(diffs, simple.GongMarshallField(stage, "Show"))
	}
	if simple.Actuate != simpleOther.Actuate {
		diffs = append(diffs, simple.GongMarshallField(stage, "Actuate"))
	}
	if simple.Value != simpleOther.Value {
		diffs = append(diffs, simple.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (titleelttype *TitleEltType) GongDiff(stage *Stage, titleelttypeOther *TitleEltType) (diffs []string) {
	// insertion point for field diffs
	if titleelttype.Name != titleelttypeOther.Name {
		diffs = append(diffs, titleelttype.GongMarshallField(stage, "Name"))
	}
	if titleelttype.Type != titleelttypeOther.Type {
		diffs = append(diffs, titleelttype.GongMarshallField(stage, "Type"))
	}
	if titleelttype.Lang != titleelttypeOther.Lang {
		diffs = append(diffs, titleelttype.GongMarshallField(stage, "Lang"))
	}
	if titleelttype.Value != titleelttypeOther.Value {
		diffs = append(diffs, titleelttype.GongMarshallField(stage, "Value"))
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
