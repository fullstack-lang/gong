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
func (displayselection *DisplaySelection) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DisplaySelections[displayselection]
	return ok
}

func (xlcell *XLCell) GongIsStaged(stage *Stage) bool {
	_, ok := stage.XLCells[xlcell]
	return ok
}

func (xlfile *XLFile) GongIsStaged(stage *Stage) bool {
	_, ok := stage.XLFiles[xlfile]
	return ok
}

func (xlrow *XLRow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.XLRows[xlrow]
	return ok
}

func (xlsheet *XLSheet) GongIsStaged(stage *Stage) bool {
	_, ok := stage.XLSheets[xlsheet]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (displayselection *DisplaySelection) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(displayselection) {
		return
	}

	displayselection.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if displayselection.XLFile != nil {
		stage.StageBranch(displayselection.XLFile)
	}
	if displayselection.XLSheet != nil {
		stage.StageBranch(displayselection.XLSheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xlcell *XLCell) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(xlcell) {
		return
	}

	xlcell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xlfile *XLFile) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(xlfile) {
		return
	}

	xlfile.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlsheet := range xlfile.Sheets {
		stage.StageBranch(_xlsheet)
	}

}

func (xlrow *XLRow) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(xlrow) {
		return
	}

	xlrow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlcell := range xlrow.Cells {
		stage.StageBranch(_xlcell)
	}

}

func (xlsheet *XLSheet) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(xlsheet) {
		return
	}

	xlsheet.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlrow := range xlsheet.Rows {
		stage.StageBranch(_xlrow)
	}
	for _, _xlcell := range xlsheet.SheetCells {
		stage.StageBranch(_xlcell)
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
	case *DisplaySelection:
		toT := GongCopyBranchDisplaySelection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XLCell:
		toT := GongCopyBranchXLCell(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XLFile:
		toT := GongCopyBranchXLFile(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XLRow:
		toT := GongCopyBranchXLRow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XLSheet:
		toT := GongCopyBranchXLSheet(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchDisplaySelection(mapOrigCopy map[any]any, displayselectionFrom *DisplaySelection) (displayselectionTo *DisplaySelection) {
	var alreadyCopied bool
	displayselectionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, displayselectionFrom)
	if alreadyCopied {
		return
	}
	displayselectionFrom.GongCopyBasicFields(displayselectionTo)

	//insertion point for the staging of instances referenced by pointers
	if displayselectionFrom.XLFile != nil {
		displayselectionTo.XLFile = GongCopyBranchXLFile(mapOrigCopy, displayselectionFrom.XLFile)
	}
	if displayselectionFrom.XLSheet != nil {
		displayselectionTo.XLSheet = GongCopyBranchXLSheet(mapOrigCopy, displayselectionFrom.XLSheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchXLCell(mapOrigCopy map[any]any, xlcellFrom *XLCell) (xlcellTo *XLCell) {
	var alreadyCopied bool
	xlcellTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, xlcellFrom)
	if alreadyCopied {
		return
	}
	xlcellFrom.GongCopyBasicFields(xlcellTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchXLFile(mapOrigCopy map[any]any, xlfileFrom *XLFile) (xlfileTo *XLFile) {
	var alreadyCopied bool
	xlfileTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, xlfileFrom)
	if alreadyCopied {
		return
	}
	xlfileFrom.GongCopyBasicFields(xlfileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlsheet := range xlfileFrom.Sheets {
		xlfileTo.Sheets = append(xlfileTo.Sheets, GongCopyBranchXLSheet(mapOrigCopy, _xlsheet))
	}

	return
}

func GongCopyBranchXLRow(mapOrigCopy map[any]any, xlrowFrom *XLRow) (xlrowTo *XLRow) {
	var alreadyCopied bool
	xlrowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, xlrowFrom)
	if alreadyCopied {
		return
	}
	xlrowFrom.GongCopyBasicFields(xlrowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlcell := range xlrowFrom.Cells {
		xlrowTo.Cells = append(xlrowTo.Cells, GongCopyBranchXLCell(mapOrigCopy, _xlcell))
	}

	return
}

func GongCopyBranchXLSheet(mapOrigCopy map[any]any, xlsheetFrom *XLSheet) (xlsheetTo *XLSheet) {
	var alreadyCopied bool
	xlsheetTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, xlsheetFrom)
	if alreadyCopied {
		return
	}
	xlsheetFrom.GongCopyBasicFields(xlsheetTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlrow := range xlsheetFrom.Rows {
		xlsheetTo.Rows = append(xlsheetTo.Rows, GongCopyBranchXLRow(mapOrigCopy, _xlrow))
	}
	for _, _xlcell := range xlsheetFrom.SheetCells {
		xlsheetTo.SheetCells = append(xlsheetTo.SheetCells, GongCopyBranchXLCell(mapOrigCopy, _xlcell))
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
func (displayselection *DisplaySelection) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(displayselection) {
		return
	}

	displayselection.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if displayselection.XLFile != nil {
		stage.UnstageBranch(displayselection.XLFile)
	}
	if displayselection.XLSheet != nil {
		stage.UnstageBranch(displayselection.XLSheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xlcell *XLCell) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(xlcell) {
		return
	}

	xlcell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xlfile *XLFile) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(xlfile) {
		return
	}

	xlfile.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlsheet := range xlfile.Sheets {
		stage.UnstageBranch(_xlsheet)
	}

}

func (xlrow *XLRow) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(xlrow) {
		return
	}

	xlrow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlcell := range xlrow.Cells {
		stage.UnstageBranch(_xlcell)
	}

}

func (xlsheet *XLSheet) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(xlsheet) {
		return
	}

	xlsheet.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlrow := range xlsheet.Rows {
		stage.UnstageBranch(_xlrow)
	}
	for _, _xlcell := range xlsheet.SheetCells {
		stage.UnstageBranch(_xlcell)
	}

}

// insertion point for pointer reconstruction from references
func (reference *DisplaySelection) GongReconstructPointersFromReferences(stage *Stage, instance *DisplaySelection) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.XLFile, stage.XLFiles_reference, instance.XLFile)
	__gong__reconstructPointer(&reference.XLSheet, stage.XLSheets_reference, instance.XLSheet)
	// insertion point for slice of pointers field
}

func (reference *XLCell) GongReconstructPointersFromReferences(stage *Stage, instance *XLCell) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *XLFile) GongReconstructPointersFromReferences(stage *Stage, instance *XLFile) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Sheets, stage.XLSheets_reference, instance.Sheets)
}

func (reference *XLRow) GongReconstructPointersFromReferences(stage *Stage, instance *XLRow) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Cells, stage.XLCells_reference, instance.Cells)
}

func (reference *XLSheet) GongReconstructPointersFromReferences(stage *Stage, instance *XLSheet) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Rows, stage.XLRows_reference, instance.Rows)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SheetCells, stage.XLCells_reference, instance.SheetCells)
}

// insertion point for pointer reconstruction from instances
func (reference *DisplaySelection) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.XLFile, stage.XLFiles_instance)
	__gong__reconstructPointerFromInstance(&reference.XLSheet, stage.XLSheets_instance)
	// insertion point for slice of pointers fields
}

func (reference *XLCell) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *XLFile) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Sheets, stage.XLSheets_instance)
}

func (reference *XLRow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Cells, stage.XLCells_instance)
}

func (reference *XLSheet) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Rows, stage.XLRows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SheetCells, stage.XLCells_instance)
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (displayselection *DisplaySelection) GongDiff(stage *Stage, displayselectionOther *DisplaySelection) (diffs []string) {
	// insertion point for field diffs
	if displayselection.Name != displayselectionOther.Name {
		diffs = append(diffs, displayselection.GongMarshallField(stage, "Name"))
	}
	if displayselection.XLFile != displayselectionOther.XLFile {
		diffs = append(diffs, displayselection.GongMarshallField(stage, "XLFile"))
	}
	if displayselection.XLSheet != displayselectionOther.XLSheet {
		diffs = append(diffs, displayselection.GongMarshallField(stage, "XLSheet"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (xlcell *XLCell) GongDiff(stage *Stage, xlcellOther *XLCell) (diffs []string) {
	// insertion point for field diffs
	if xlcell.Name != xlcellOther.Name {
		diffs = append(diffs, xlcell.GongMarshallField(stage, "Name"))
	}
	if xlcell.X != xlcellOther.X {
		diffs = append(diffs, xlcell.GongMarshallField(stage, "X"))
	}
	if xlcell.Y != xlcellOther.Y {
		diffs = append(diffs, xlcell.GongMarshallField(stage, "Y"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (xlfile *XLFile) GongDiff(stage *Stage, xlfileOther *XLFile) (diffs []string) {
	// insertion point for field diffs
	if xlfile.Name != xlfileOther.Name {
		diffs = append(diffs, xlfile.GongMarshallField(stage, "Name"))
	}
	if xlfile.NbSheets != xlfileOther.NbSheets {
		diffs = append(diffs, xlfile.GongMarshallField(stage, "NbSheets"))
	}
	if ops := __gong__diffSliceOfPointers(stage, xlfile, "Sheets", xlfileOther.Sheets, xlfile.Sheets); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (xlrow *XLRow) GongDiff(stage *Stage, xlrowOther *XLRow) (diffs []string) {
	// insertion point for field diffs
	if xlrow.Name != xlrowOther.Name {
		diffs = append(diffs, xlrow.GongMarshallField(stage, "Name"))
	}
	if xlrow.RowIndex != xlrowOther.RowIndex {
		diffs = append(diffs, xlrow.GongMarshallField(stage, "RowIndex"))
	}
	if ops := __gong__diffSliceOfPointers(stage, xlrow, "Cells", xlrowOther.Cells, xlrow.Cells); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (xlsheet *XLSheet) GongDiff(stage *Stage, xlsheetOther *XLSheet) (diffs []string) {
	// insertion point for field diffs
	if xlsheet.Name != xlsheetOther.Name {
		diffs = append(diffs, xlsheet.GongMarshallField(stage, "Name"))
	}
	if xlsheet.MaxRow != xlsheetOther.MaxRow {
		diffs = append(diffs, xlsheet.GongMarshallField(stage, "MaxRow"))
	}
	if xlsheet.MaxCol != xlsheetOther.MaxCol {
		diffs = append(diffs, xlsheet.GongMarshallField(stage, "MaxCol"))
	}
	if xlsheet.NbRows != xlsheetOther.NbRows {
		diffs = append(diffs, xlsheet.GongMarshallField(stage, "NbRows"))
	}
	if ops := __gong__diffSliceOfPointers(stage, xlsheet, "Rows", xlsheetOther.Rows, xlsheet.Rows); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, xlsheet, "SheetCells", xlsheetOther.SheetCells, xlsheet.SheetCells); ops != "" {
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
