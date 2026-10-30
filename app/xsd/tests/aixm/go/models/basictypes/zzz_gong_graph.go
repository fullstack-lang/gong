// generated code - do not edit
package basictypes

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
func (codelisttype *CodeListType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeListTypes[codelisttype]
	return ok
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOrNilReasonListTypes[codeornilreasonlisttype]
	return ok
}

func (codetype *CodeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTypes[codetype]
	return ok
}

func (codewithauthoritytype *CodeWithAuthorityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeWithAuthorityTypes[codewithauthoritytype]
	return ok
}

func (coordinatestype *CoordinatesType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CoordinatesTypes[coordinatestype]
	return ok
}

func (measurelisttype *MeasureListType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureListTypes[measurelisttype]
	return ok
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureOrNilReasonListTypes[measureornilreasonlisttype]
	return ok
}

func (measuretype *MeasureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MeasureTypes[measuretype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (codelisttype *CodeListType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelisttype) {
		return
	}

	codelisttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeornilreasonlisttype) {
		return
	}

	codeornilreasonlisttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetype *CodeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetype) {
		return
	}

	codetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codewithauthoritytype *CodeWithAuthorityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codewithauthoritytype) {
		return
	}

	codewithauthoritytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coordinatestype *CoordinatesType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coordinatestype) {
		return
	}

	coordinatestype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measurelisttype *MeasureListType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measurelisttype) {
		return
	}

	measurelisttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measureornilreasonlisttype) {
		return
	}

	measureornilreasonlisttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measuretype *MeasureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(measuretype) {
		return
	}

	measuretype.Stage(stage)

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
	case *CodeListType:
		toT := GongCopyBranchCodeListType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOrNilReasonListType:
		toT := GongCopyBranchCodeOrNilReasonListType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeType:
		toT := GongCopyBranchCodeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeWithAuthorityType:
		toT := GongCopyBranchCodeWithAuthorityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CoordinatesType:
		toT := GongCopyBranchCoordinatesType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureListType:
		toT := GongCopyBranchMeasureListType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureOrNilReasonListType:
		toT := GongCopyBranchMeasureOrNilReasonListType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MeasureType:
		toT := GongCopyBranchMeasureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCodeListType(mapOrigCopy map[any]any, codelisttypeFrom *CodeListType) (codelisttypeTo *CodeListType) {
	var alreadyCopied bool
	codelisttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelisttypeFrom)
	if alreadyCopied {
		return
	}
	codelisttypeFrom.GongCopyBasicFields(codelisttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeOrNilReasonListType(mapOrigCopy map[any]any, codeornilreasonlisttypeFrom *CodeOrNilReasonListType) (codeornilreasonlisttypeTo *CodeOrNilReasonListType) {
	var alreadyCopied bool
	codeornilreasonlisttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeornilreasonlisttypeFrom)
	if alreadyCopied {
		return
	}
	codeornilreasonlisttypeFrom.GongCopyBasicFields(codeornilreasonlisttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeType(mapOrigCopy map[any]any, codetypeFrom *CodeType) (codetypeTo *CodeType) {
	var alreadyCopied bool
	codetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetypeFrom)
	if alreadyCopied {
		return
	}
	codetypeFrom.GongCopyBasicFields(codetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeWithAuthorityType(mapOrigCopy map[any]any, codewithauthoritytypeFrom *CodeWithAuthorityType) (codewithauthoritytypeTo *CodeWithAuthorityType) {
	var alreadyCopied bool
	codewithauthoritytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codewithauthoritytypeFrom)
	if alreadyCopied {
		return
	}
	codewithauthoritytypeFrom.GongCopyBasicFields(codewithauthoritytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCoordinatesType(mapOrigCopy map[any]any, coordinatestypeFrom *CoordinatesType) (coordinatestypeTo *CoordinatesType) {
	var alreadyCopied bool
	coordinatestypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coordinatestypeFrom)
	if alreadyCopied {
		return
	}
	coordinatestypeFrom.GongCopyBasicFields(coordinatestypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasureListType(mapOrigCopy map[any]any, measurelisttypeFrom *MeasureListType) (measurelisttypeTo *MeasureListType) {
	var alreadyCopied bool
	measurelisttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measurelisttypeFrom)
	if alreadyCopied {
		return
	}
	measurelisttypeFrom.GongCopyBasicFields(measurelisttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasureOrNilReasonListType(mapOrigCopy map[any]any, measureornilreasonlisttypeFrom *MeasureOrNilReasonListType) (measureornilreasonlisttypeTo *MeasureOrNilReasonListType) {
	var alreadyCopied bool
	measureornilreasonlisttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measureornilreasonlisttypeFrom)
	if alreadyCopied {
		return
	}
	measureornilreasonlisttypeFrom.GongCopyBasicFields(measureornilreasonlisttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMeasureType(mapOrigCopy map[any]any, measuretypeFrom *MeasureType) (measuretypeTo *MeasureType) {
	var alreadyCopied bool
	measuretypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, measuretypeFrom)
	if alreadyCopied {
		return
	}
	measuretypeFrom.GongCopyBasicFields(measuretypeTo)

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
func (codelisttype *CodeListType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelisttype) {
		return
	}

	codelisttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeornilreasonlisttype) {
		return
	}

	codeornilreasonlisttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetype *CodeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetype) {
		return
	}

	codetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codewithauthoritytype *CodeWithAuthorityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codewithauthoritytype) {
		return
	}

	codewithauthoritytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coordinatestype *CoordinatesType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coordinatestype) {
		return
	}

	coordinatestype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measurelisttype *MeasureListType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measurelisttype) {
		return
	}

	measurelisttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measureornilreasonlisttype) {
		return
	}

	measureornilreasonlisttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (measuretype *MeasureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(measuretype) {
		return
	}

	measuretype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *CodeListType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeListType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeOrNilReasonListType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOrNilReasonListType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeWithAuthorityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeWithAuthorityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CoordinatesType) GongReconstructPointersFromReferences(stage *Stage, instance *CoordinatesType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MeasureListType) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureListType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MeasureOrNilReasonListType) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureOrNilReasonListType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MeasureType) GongReconstructPointersFromReferences(stage *Stage, instance *MeasureType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *CodeListType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeOrNilReasonListType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeWithAuthorityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CoordinatesType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MeasureListType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MeasureOrNilReasonListType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MeasureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelisttype *CodeListType) GongDiff(stage *Stage, codelisttypeOther *CodeListType) (diffs []string) {
	// insertion point for field diffs
	if codelisttype.Name != codelisttypeOther.Name {
		diffs = append(diffs, codelisttype.GongMarshallField(stage, "Name"))
	}
	if codelisttype.CodeSpace != codelisttypeOther.CodeSpace {
		diffs = append(diffs, codelisttype.GongMarshallField(stage, "CodeSpace"))
	}
	if codelisttype.EnclosedText != codelisttypeOther.EnclosedText {
		diffs = append(diffs, codelisttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeornilreasonlisttype *CodeOrNilReasonListType) GongDiff(stage *Stage, codeornilreasonlisttypeOther *CodeOrNilReasonListType) (diffs []string) {
	// insertion point for field diffs
	if codeornilreasonlisttype.Name != codeornilreasonlisttypeOther.Name {
		diffs = append(diffs, codeornilreasonlisttype.GongMarshallField(stage, "Name"))
	}
	if codeornilreasonlisttype.CodeSpace != codeornilreasonlisttypeOther.CodeSpace {
		diffs = append(diffs, codeornilreasonlisttype.GongMarshallField(stage, "CodeSpace"))
	}
	if codeornilreasonlisttype.EnclosedText != codeornilreasonlisttypeOther.EnclosedText {
		diffs = append(diffs, codeornilreasonlisttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetype *CodeType) GongDiff(stage *Stage, codetypeOther *CodeType) (diffs []string) {
	// insertion point for field diffs
	if codetype.Name != codetypeOther.Name {
		diffs = append(diffs, codetype.GongMarshallField(stage, "Name"))
	}
	if codetype.CodeSpace != codetypeOther.CodeSpace {
		diffs = append(diffs, codetype.GongMarshallField(stage, "CodeSpace"))
	}
	if codetype.EnclosedText != codetypeOther.EnclosedText {
		diffs = append(diffs, codetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codewithauthoritytype *CodeWithAuthorityType) GongDiff(stage *Stage, codewithauthoritytypeOther *CodeWithAuthorityType) (diffs []string) {
	// insertion point for field diffs
	if codewithauthoritytype.Name != codewithauthoritytypeOther.Name {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "Name"))
	}
	if codewithauthoritytype.CodeSpace != codewithauthoritytypeOther.CodeSpace {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "CodeSpace"))
	}
	if codewithauthoritytype.EnclosedText != codewithauthoritytypeOther.EnclosedText {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coordinatestype *CoordinatesType) GongDiff(stage *Stage, coordinatestypeOther *CoordinatesType) (diffs []string) {
	// insertion point for field diffs
	if coordinatestype.Name != coordinatestypeOther.Name {
		diffs = append(diffs, coordinatestype.GongMarshallField(stage, "Name"))
	}
	if coordinatestype.Decimal != coordinatestypeOther.Decimal {
		diffs = append(diffs, coordinatestype.GongMarshallField(stage, "Decimal"))
	}
	if coordinatestype.Cs != coordinatestypeOther.Cs {
		diffs = append(diffs, coordinatestype.GongMarshallField(stage, "Cs"))
	}
	if coordinatestype.Ts != coordinatestypeOther.Ts {
		diffs = append(diffs, coordinatestype.GongMarshallField(stage, "Ts"))
	}
	if coordinatestype.EnclosedText != coordinatestypeOther.EnclosedText {
		diffs = append(diffs, coordinatestype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measurelisttype *MeasureListType) GongDiff(stage *Stage, measurelisttypeOther *MeasureListType) (diffs []string) {
	// insertion point for field diffs
	if measurelisttype.Name != measurelisttypeOther.Name {
		diffs = append(diffs, measurelisttype.GongMarshallField(stage, "Name"))
	}
	if measurelisttype.Uom != measurelisttypeOther.Uom {
		diffs = append(diffs, measurelisttype.GongMarshallField(stage, "Uom"))
	}
	if measurelisttype.EnclosedText != measurelisttypeOther.EnclosedText {
		diffs = append(diffs, measurelisttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongDiff(stage *Stage, measureornilreasonlisttypeOther *MeasureOrNilReasonListType) (diffs []string) {
	// insertion point for field diffs
	if measureornilreasonlisttype.Name != measureornilreasonlisttypeOther.Name {
		diffs = append(diffs, measureornilreasonlisttype.GongMarshallField(stage, "Name"))
	}
	if measureornilreasonlisttype.Uom != measureornilreasonlisttypeOther.Uom {
		diffs = append(diffs, measureornilreasonlisttype.GongMarshallField(stage, "Uom"))
	}
	if measureornilreasonlisttype.EnclosedText != measureornilreasonlisttypeOther.EnclosedText {
		diffs = append(diffs, measureornilreasonlisttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (measuretype *MeasureType) GongDiff(stage *Stage, measuretypeOther *MeasureType) (diffs []string) {
	// insertion point for field diffs
	if measuretype.Name != measuretypeOther.Name {
		diffs = append(diffs, measuretype.GongMarshallField(stage, "Name"))
	}
	if measuretype.Uom != measuretypeOther.Uom {
		diffs = append(diffs, measuretype.GongMarshallField(stage, "Uom"))
	}
	if measuretype.EnclosedText != measuretypeOther.EnclosedText {
		diffs = append(diffs, measuretype.GongMarshallField(stage, "EnclosedText"))
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
