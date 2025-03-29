// generated code - do not edit
package models

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
