// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *DisplaySelection:
		ok = stage.IsStagedDisplaySelection(target)

	case *XLCell:
		ok = stage.IsStagedXLCell(target)

	case *XLFile:
		ok = stage.IsStagedXLFile(target)

	case *XLRow:
		ok = stage.IsStagedXLRow(target)

	case *XLSheet:
		ok = stage.IsStagedXLSheet(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *DisplaySelection:
		ok = stage.IsStagedDisplaySelection(target)

	case *XLCell:
		ok = stage.IsStagedXLCell(target)

	case *XLFile:
		ok = stage.IsStagedXLFile(target)

	case *XLRow:
		ok = stage.IsStagedXLRow(target)

	case *XLSheet:
		ok = stage.IsStagedXLSheet(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedDisplaySelection(displayselection *DisplaySelection) (ok bool) {

	_, ok = stage.DisplaySelections[displayselection]

	return
}

func (stage *Stage) IsStagedXLCell(xlcell *XLCell) (ok bool) {

	_, ok = stage.XLCells[xlcell]

	return
}

func (stage *Stage) IsStagedXLFile(xlfile *XLFile) (ok bool) {

	_, ok = stage.XLFiles[xlfile]

	return
}

func (stage *Stage) IsStagedXLRow(xlrow *XLRow) (ok bool) {

	_, ok = stage.XLRows[xlrow]

	return
}

func (stage *Stage) IsStagedXLSheet(xlsheet *XLSheet) (ok bool) {

	_, ok = stage.XLSheets[xlsheet]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *DisplaySelection:
		stage.StageBranchDisplaySelection(target)

	case *XLCell:
		stage.StageBranchXLCell(target)

	case *XLFile:
		stage.StageBranchXLFile(target)

	case *XLRow:
		stage.StageBranchXLRow(target)

	case *XLSheet:
		stage.StageBranchXLSheet(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchDisplaySelection(displayselection *DisplaySelection) {

	// check if instance is already staged
	if IsStaged(stage, displayselection) {
		return
	}

	displayselection.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if displayselection.XLFile != nil {
		StageBranch(stage, displayselection.XLFile)
	}
	if displayselection.XLSheet != nil {
		StageBranch(stage, displayselection.XLSheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchXLCell(xlcell *XLCell) {

	// check if instance is already staged
	if IsStaged(stage, xlcell) {
		return
	}

	xlcell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchXLFile(xlfile *XLFile) {

	// check if instance is already staged
	if IsStaged(stage, xlfile) {
		return
	}

	xlfile.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlsheet := range xlfile.Sheets {
		StageBranch(stage, _xlsheet)
	}

}

func (stage *Stage) StageBranchXLRow(xlrow *XLRow) {

	// check if instance is already staged
	if IsStaged(stage, xlrow) {
		return
	}

	xlrow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlcell := range xlrow.Cells {
		StageBranch(stage, _xlcell)
	}

}

func (stage *Stage) StageBranchXLSheet(xlsheet *XLSheet) {

	// check if instance is already staged
	if IsStaged(stage, xlsheet) {
		return
	}

	xlsheet.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlrow := range xlsheet.Rows {
		StageBranch(stage, _xlrow)
	}
	for _, _xlcell := range xlsheet.SheetCells {
		StageBranch(stage, _xlcell)
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
	case *DisplaySelection:
		toT := CopyBranchDisplaySelection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XLCell:
		toT := CopyBranchXLCell(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XLFile:
		toT := CopyBranchXLFile(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XLRow:
		toT := CopyBranchXLRow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XLSheet:
		toT := CopyBranchXLSheet(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchDisplaySelection(mapOrigCopy map[any]any, displayselectionFrom *DisplaySelection) (displayselectionTo *DisplaySelection) {

	// displayselectionFrom has already been copied
	if _displayselectionTo, ok := mapOrigCopy[displayselectionFrom]; ok {
		displayselectionTo = _displayselectionTo.(*DisplaySelection)
		return
	}

	displayselectionTo = new(DisplaySelection)
	mapOrigCopy[displayselectionFrom] = displayselectionTo
	displayselectionFrom.CopyBasicFields(displayselectionTo)

	//insertion point for the staging of instances referenced by pointers
	if displayselectionFrom.XLFile != nil {
		displayselectionTo.XLFile = CopyBranchXLFile(mapOrigCopy, displayselectionFrom.XLFile)
	}
	if displayselectionFrom.XLSheet != nil {
		displayselectionTo.XLSheet = CopyBranchXLSheet(mapOrigCopy, displayselectionFrom.XLSheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXLCell(mapOrigCopy map[any]any, xlcellFrom *XLCell) (xlcellTo *XLCell) {

	// xlcellFrom has already been copied
	if _xlcellTo, ok := mapOrigCopy[xlcellFrom]; ok {
		xlcellTo = _xlcellTo.(*XLCell)
		return
	}

	xlcellTo = new(XLCell)
	mapOrigCopy[xlcellFrom] = xlcellTo
	xlcellFrom.CopyBasicFields(xlcellTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXLFile(mapOrigCopy map[any]any, xlfileFrom *XLFile) (xlfileTo *XLFile) {

	// xlfileFrom has already been copied
	if _xlfileTo, ok := mapOrigCopy[xlfileFrom]; ok {
		xlfileTo = _xlfileTo.(*XLFile)
		return
	}

	xlfileTo = new(XLFile)
	mapOrigCopy[xlfileFrom] = xlfileTo
	xlfileFrom.CopyBasicFields(xlfileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlsheet := range xlfileFrom.Sheets {
		xlfileTo.Sheets = append(xlfileTo.Sheets, CopyBranchXLSheet(mapOrigCopy, _xlsheet))
	}

	return
}

func CopyBranchXLRow(mapOrigCopy map[any]any, xlrowFrom *XLRow) (xlrowTo *XLRow) {

	// xlrowFrom has already been copied
	if _xlrowTo, ok := mapOrigCopy[xlrowFrom]; ok {
		xlrowTo = _xlrowTo.(*XLRow)
		return
	}

	xlrowTo = new(XLRow)
	mapOrigCopy[xlrowFrom] = xlrowTo
	xlrowFrom.CopyBasicFields(xlrowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlcell := range xlrowFrom.Cells {
		xlrowTo.Cells = append(xlrowTo.Cells, CopyBranchXLCell(mapOrigCopy, _xlcell))
	}

	return
}

func CopyBranchXLSheet(mapOrigCopy map[any]any, xlsheetFrom *XLSheet) (xlsheetTo *XLSheet) {

	// xlsheetFrom has already been copied
	if _xlsheetTo, ok := mapOrigCopy[xlsheetFrom]; ok {
		xlsheetTo = _xlsheetTo.(*XLSheet)
		return
	}

	xlsheetTo = new(XLSheet)
	mapOrigCopy[xlsheetFrom] = xlsheetTo
	xlsheetFrom.CopyBasicFields(xlsheetTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlrow := range xlsheetFrom.Rows {
		xlsheetTo.Rows = append(xlsheetTo.Rows, CopyBranchXLRow(mapOrigCopy, _xlrow))
	}
	for _, _xlcell := range xlsheetFrom.SheetCells {
		xlsheetTo.SheetCells = append(xlsheetTo.SheetCells, CopyBranchXLCell(mapOrigCopy, _xlcell))
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
	case *DisplaySelection:
		stage.UnstageBranchDisplaySelection(target)

	case *XLCell:
		stage.UnstageBranchXLCell(target)

	case *XLFile:
		stage.UnstageBranchXLFile(target)

	case *XLRow:
		stage.UnstageBranchXLRow(target)

	case *XLSheet:
		stage.UnstageBranchXLSheet(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchDisplaySelection(displayselection *DisplaySelection) {

	// check if instance is already staged
	if !IsStaged(stage, displayselection) {
		return
	}

	displayselection.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if displayselection.XLFile != nil {
		UnstageBranch(stage, displayselection.XLFile)
	}
	if displayselection.XLSheet != nil {
		UnstageBranch(stage, displayselection.XLSheet)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchXLCell(xlcell *XLCell) {

	// check if instance is already staged
	if !IsStaged(stage, xlcell) {
		return
	}

	xlcell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchXLFile(xlfile *XLFile) {

	// check if instance is already staged
	if !IsStaged(stage, xlfile) {
		return
	}

	xlfile.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlsheet := range xlfile.Sheets {
		UnstageBranch(stage, _xlsheet)
	}

}

func (stage *Stage) UnstageBranchXLRow(xlrow *XLRow) {

	// check if instance is already staged
	if !IsStaged(stage, xlrow) {
		return
	}

	xlrow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlcell := range xlrow.Cells {
		UnstageBranch(stage, _xlcell)
	}

}

func (stage *Stage) UnstageBranchXLSheet(xlsheet *XLSheet) {

	// check if instance is already staged
	if !IsStaged(stage, xlsheet) {
		return
	}

	xlsheet.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlrow := range xlsheet.Rows {
		UnstageBranch(stage, _xlrow)
	}
	for _, _xlcell := range xlsheet.SheetCells {
		UnstageBranch(stage, _xlcell)
	}

}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (displayselection *DisplaySelection) GongDiff(stage *Stage, displayselectionOther *DisplaySelection) (diffs []string) {
	// insertion point for field diffs
	if displayselection.Name != displayselectionOther.Name {
		diffs = append(diffs, displayselection.GongMarshallField(stage, "Name"))
	}
	if (displayselection.XLFile == nil) != (displayselectionOther.XLFile == nil) {
		diffs = append(diffs, "XLFile")
	} else if displayselection.XLFile != nil && displayselectionOther.XLFile != nil {
		if displayselection.XLFile != displayselectionOther.XLFile {
			diffs = append(diffs, displayselection.GongMarshallField(stage, "XLFile"))
		}
	}
	if (displayselection.XLSheet == nil) != (displayselectionOther.XLSheet == nil) {
		diffs = append(diffs, "XLSheet")
	} else if displayselection.XLSheet != nil && displayselectionOther.XLSheet != nil {
		if displayselection.XLSheet != displayselectionOther.XLSheet {
			diffs = append(diffs, displayselection.GongMarshallField(stage, "XLSheet"))
		}
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
	SheetsDifferent := false
	if len(xlfile.Sheets) != len(xlfileOther.Sheets) {
		SheetsDifferent = true
	} else {
		for i := range xlfile.Sheets {
			if (xlfile.Sheets[i] == nil) != (xlfileOther.Sheets[i] == nil) {
				SheetsDifferent = true
				break
			} else if xlfile.Sheets[i] != nil && xlfileOther.Sheets[i] != nil {
				// this is a pointer comparaison
				if xlfile.Sheets[i] != xlfileOther.Sheets[i] {
					SheetsDifferent = true
					break
				}
			}
		}
	}
	if SheetsDifferent {
		ops := Diff(stage, xlfile, xlfileOther, "Sheets", xlfileOther.Sheets, xlfile.Sheets)
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
	CellsDifferent := false
	if len(xlrow.Cells) != len(xlrowOther.Cells) {
		CellsDifferent = true
	} else {
		for i := range xlrow.Cells {
			if (xlrow.Cells[i] == nil) != (xlrowOther.Cells[i] == nil) {
				CellsDifferent = true
				break
			} else if xlrow.Cells[i] != nil && xlrowOther.Cells[i] != nil {
				// this is a pointer comparaison
				if xlrow.Cells[i] != xlrowOther.Cells[i] {
					CellsDifferent = true
					break
				}
			}
		}
	}
	if CellsDifferent {
		ops := Diff(stage, xlrow, xlrowOther, "Cells", xlrowOther.Cells, xlrow.Cells)
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
	RowsDifferent := false
	if len(xlsheet.Rows) != len(xlsheetOther.Rows) {
		RowsDifferent = true
	} else {
		for i := range xlsheet.Rows {
			if (xlsheet.Rows[i] == nil) != (xlsheetOther.Rows[i] == nil) {
				RowsDifferent = true
				break
			} else if xlsheet.Rows[i] != nil && xlsheetOther.Rows[i] != nil {
				// this is a pointer comparaison
				if xlsheet.Rows[i] != xlsheetOther.Rows[i] {
					RowsDifferent = true
					break
				}
			}
		}
	}
	if RowsDifferent {
		ops := Diff(stage, xlsheet, xlsheetOther, "Rows", xlsheetOther.Rows, xlsheet.Rows)
		diffs = append(diffs, ops)
	}
	SheetCellsDifferent := false
	if len(xlsheet.SheetCells) != len(xlsheetOther.SheetCells) {
		SheetCellsDifferent = true
	} else {
		for i := range xlsheet.SheetCells {
			if (xlsheet.SheetCells[i] == nil) != (xlsheetOther.SheetCells[i] == nil) {
				SheetCellsDifferent = true
				break
			} else if xlsheet.SheetCells[i] != nil && xlsheetOther.SheetCells[i] != nil {
				// this is a pointer comparaison
				if xlsheet.SheetCells[i] != xlsheetOther.SheetCells[i] {
					SheetCellsDifferent = true
					break
				}
			}
		}
	}
	if SheetCellsDifferent {
		ops := Diff(stage, xlsheet, xlsheetOther, "SheetCells", xlsheetOther.SheetCells, xlsheet.SheetCells)
		diffs = append(diffs, ops)
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
			ops += fmt.Sprintf("\t%s.%s = slices.Delete( %s.%s, %d, %d)\n", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
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
			ops += fmt.Sprintf("\t%s.%s = slices.Insert( %s.%s, %d, %s)\n",  a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
