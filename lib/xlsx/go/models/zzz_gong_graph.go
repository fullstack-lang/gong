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
func (displayselection *DisplaySelection) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DisplaySelections[displayselection]

	return
}

func (stage *Stage) IsStagedDisplaySelection(displayselection *DisplaySelection) (ok bool) {

	return displayselection.GongIsStaged(stage)
}

func (xlcell *XLCell) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.XLCells[xlcell]

	return
}

func (stage *Stage) IsStagedXLCell(xlcell *XLCell) (ok bool) {

	return xlcell.GongIsStaged(stage)
}

func (xlfile *XLFile) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.XLFiles[xlfile]

	return
}

func (stage *Stage) IsStagedXLFile(xlfile *XLFile) (ok bool) {

	return xlfile.GongIsStaged(stage)
}

func (xlrow *XLRow) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.XLRows[xlrow]

	return
}

func (stage *Stage) IsStagedXLRow(xlrow *XLRow) (ok bool) {

	return xlrow.GongIsStaged(stage)
}

func (xlsheet *XLSheet) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.XLSheets[xlsheet]

	return
}

func (stage *Stage) IsStagedXLSheet(xlsheet *XLSheet) (ok bool) {

	return xlsheet.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (displayselection *DisplaySelection) GongStageBranch(stage *Stage) {
	stage.StageBranchDisplaySelection(displayselection)
}

func (stage *Stage) StageBranchDisplaySelection(displayselection *DisplaySelection) {

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
	stage.StageBranchXLCell(xlcell)
}

func (stage *Stage) StageBranchXLCell(xlcell *XLCell) {

	// check if instance is already staged
	if stage.IsStaged(xlcell) {
		return
	}

	xlcell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xlfile *XLFile) GongStageBranch(stage *Stage) {
	stage.StageBranchXLFile(xlfile)
}

func (stage *Stage) StageBranchXLFile(xlfile *XLFile) {

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
	stage.StageBranchXLRow(xlrow)
}

func (stage *Stage) StageBranchXLRow(xlrow *XLRow) {

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
	stage.StageBranchXLSheet(xlsheet)
}

func (stage *Stage) StageBranchXLSheet(xlsheet *XLSheet) {

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

	// displayselectionFrom has already been copied
	if _displayselectionTo, ok := mapOrigCopy[displayselectionFrom]; ok {
		displayselectionTo = _displayselectionTo.(*DisplaySelection)
		return
	}

	displayselectionTo = new(DisplaySelection)
	mapOrigCopy[displayselectionFrom] = displayselectionTo
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

	// xlcellFrom has already been copied
	if _xlcellTo, ok := mapOrigCopy[xlcellFrom]; ok {
		xlcellTo = _xlcellTo.(*XLCell)
		return
	}

	xlcellTo = new(XLCell)
	mapOrigCopy[xlcellFrom] = xlcellTo
	xlcellFrom.GongCopyBasicFields(xlcellTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchXLFile(mapOrigCopy map[any]any, xlfileFrom *XLFile) (xlfileTo *XLFile) {

	// xlfileFrom has already been copied
	if _xlfileTo, ok := mapOrigCopy[xlfileFrom]; ok {
		xlfileTo = _xlfileTo.(*XLFile)
		return
	}

	xlfileTo = new(XLFile)
	mapOrigCopy[xlfileFrom] = xlfileTo
	xlfileFrom.GongCopyBasicFields(xlfileTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlsheet := range xlfileFrom.Sheets {
		xlfileTo.Sheets = append(xlfileTo.Sheets, GongCopyBranchXLSheet(mapOrigCopy, _xlsheet))
	}

	return
}

func GongCopyBranchXLRow(mapOrigCopy map[any]any, xlrowFrom *XLRow) (xlrowTo *XLRow) {

	// xlrowFrom has already been copied
	if _xlrowTo, ok := mapOrigCopy[xlrowFrom]; ok {
		xlrowTo = _xlrowTo.(*XLRow)
		return
	}

	xlrowTo = new(XLRow)
	mapOrigCopy[xlrowFrom] = xlrowTo
	xlrowFrom.GongCopyBasicFields(xlrowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xlcell := range xlrowFrom.Cells {
		xlrowTo.Cells = append(xlrowTo.Cells, GongCopyBranchXLCell(mapOrigCopy, _xlcell))
	}

	return
}

func GongCopyBranchXLSheet(mapOrigCopy map[any]any, xlsheetFrom *XLSheet) (xlsheetTo *XLSheet) {

	// xlsheetFrom has already been copied
	if _xlsheetTo, ok := mapOrigCopy[xlsheetFrom]; ok {
		xlsheetTo = _xlsheetTo.(*XLSheet)
		return
	}

	xlsheetTo = new(XLSheet)
	mapOrigCopy[xlsheetFrom] = xlsheetTo
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
	stage.UnstageBranchDisplaySelection(displayselection)
}

func (stage *Stage) UnstageBranchDisplaySelection(displayselection *DisplaySelection) {

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
	stage.UnstageBranchXLCell(xlcell)
}

func (stage *Stage) UnstageBranchXLCell(xlcell *XLCell) {

	// check if instance is already staged
	if !stage.IsStaged(xlcell) {
		return
	}

	xlcell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xlfile *XLFile) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchXLFile(xlfile)
}

func (stage *Stage) UnstageBranchXLFile(xlfile *XLFile) {

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
	stage.UnstageBranchXLRow(xlrow)
}

func (stage *Stage) UnstageBranchXLRow(xlrow *XLRow) {

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
	stage.UnstageBranchXLSheet(xlsheet)
}

func (stage *Stage) UnstageBranchXLSheet(xlsheet *XLSheet) {

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
	if instance.XLFile != nil {
		reference.XLFile = stage.XLFiles_reference[instance.XLFile]
	}
	if instance.XLSheet != nil {
		reference.XLSheet = stage.XLSheets_reference[instance.XLSheet]
	}
	// insertion point for slice of pointers field
}

func (reference *XLCell) GongReconstructPointersFromReferences(stage *Stage, instance *XLCell) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *XLFile) GongReconstructPointersFromReferences(stage *Stage, instance *XLFile) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Sheets = reference.Sheets[:0]
	for _, _b := range instance.Sheets {
		reference.Sheets = append(reference.Sheets, stage.XLSheets_reference[_b])
	}
}

func (reference *XLRow) GongReconstructPointersFromReferences(stage *Stage, instance *XLRow) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Cells = reference.Cells[:0]
	for _, _b := range instance.Cells {
		reference.Cells = append(reference.Cells, stage.XLCells_reference[_b])
	}
}

func (reference *XLSheet) GongReconstructPointersFromReferences(stage *Stage, instance *XLSheet) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Rows = reference.Rows[:0]
	for _, _b := range instance.Rows {
		reference.Rows = append(reference.Rows, stage.XLRows_reference[_b])
	}
	reference.SheetCells = reference.SheetCells[:0]
	for _, _b := range instance.SheetCells {
		reference.SheetCells = append(reference.SheetCells, stage.XLCells_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *DisplaySelection) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.XLFile; _reference != nil {
		reference.XLFile = nil
		if _instance, ok := stage.XLFiles_instance[_reference]; ok {
			reference.XLFile = _instance
		}
	}
	if _reference := reference.XLSheet; _reference != nil {
		reference.XLSheet = nil
		if _instance, ok := stage.XLSheets_instance[_reference]; ok {
			reference.XLSheet = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *XLCell) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *XLFile) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Sheets []*XLSheet
	for _, _reference := range reference.Sheets {
		if _instance, ok := stage.XLSheets_instance[_reference]; ok {
			_Sheets = append(_Sheets, _instance)
		}
	}
	reference.Sheets = _Sheets
}

func (reference *XLRow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Cells []*XLCell
	for _, _reference := range reference.Cells {
		if _instance, ok := stage.XLCells_instance[_reference]; ok {
			_Cells = append(_Cells, _instance)
		}
	}
	reference.Cells = _Cells
}

func (reference *XLSheet) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Rows []*XLRow
	for _, _reference := range reference.Rows {
		if _instance, ok := stage.XLRows_instance[_reference]; ok {
			_Rows = append(_Rows, _instance)
		}
	}
	reference.Rows = _Rows
	var _SheetCells []*XLCell
	for _, _reference := range reference.SheetCells {
		if _instance, ok := stage.XLCells_instance[_reference]; ok {
			_SheetCells = append(_SheetCells, _instance)
		}
	}
	reference.SheetCells = _SheetCells
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
		diffs = append(diffs, displayselection.GongMarshallField(stage, "XLFile"))
	} else if displayselection.XLFile != nil && displayselectionOther.XLFile != nil {
		if displayselection.XLFile != displayselectionOther.XLFile {
			diffs = append(diffs, displayselection.GongMarshallField(stage, "XLFile"))
		}
	}
	if (displayselection.XLSheet == nil) != (displayselectionOther.XLSheet == nil) {
		diffs = append(diffs, displayselection.GongMarshallField(stage, "XLSheet"))
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
		ops := stage.Diff(
			xlfile,
			"Sheets",
			len(xlfileOther.Sheets),
			len(xlfile.Sheets),
			func(i, j int) bool {
				return xlfileOther.Sheets[i] == xlfile.Sheets[j]
			},
			func(j int) string {
				return xlfile.Sheets[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			xlrow,
			"Cells",
			len(xlrowOther.Cells),
			len(xlrow.Cells),
			func(i, j int) bool {
				return xlrowOther.Cells[i] == xlrow.Cells[j]
			},
			func(j int) string {
				return xlrow.Cells[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			xlsheet,
			"Rows",
			len(xlsheetOther.Rows),
			len(xlsheet.Rows),
			func(i, j int) bool {
				return xlsheetOther.Rows[i] == xlsheet.Rows[j]
			},
			func(j int) string {
				return xlsheet.Rows[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			xlsheet,
			"SheetCells",
			len(xlsheetOther.SheetCells),
			len(xlsheet.SheetCells),
			func(i, j int) bool {
				return xlsheetOther.SheetCells[i] == xlsheet.SheetCells[j]
			},
			func(j int) string {
				return xlsheet.SheetCells[j].GongGetIdentifier(stage)
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
