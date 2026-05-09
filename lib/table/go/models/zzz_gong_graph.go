// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Button:
		ok = stage.IsStagedButton(target)

	case *Cell:
		ok = stage.IsStagedCell(target)

	case *CellBoolean:
		ok = stage.IsStagedCellBoolean(target)

	case *CellFloat64:
		ok = stage.IsStagedCellFloat64(target)

	case *CellIcon:
		ok = stage.IsStagedCellIcon(target)

	case *CellInt:
		ok = stage.IsStagedCellInt(target)

	case *CellString:
		ok = stage.IsStagedCellString(target)

	case *DisplayedColumn:
		ok = stage.IsStagedDisplayedColumn(target)

	case *Row:
		ok = stage.IsStagedRow(target)

	case *SVGIcon:
		ok = stage.IsStagedSVGIcon(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Button:
		ok = stage.IsStagedButton(target)

	case *Cell:
		ok = stage.IsStagedCell(target)

	case *CellBoolean:
		ok = stage.IsStagedCellBoolean(target)

	case *CellFloat64:
		ok = stage.IsStagedCellFloat64(target)

	case *CellIcon:
		ok = stage.IsStagedCellIcon(target)

	case *CellInt:
		ok = stage.IsStagedCellInt(target)

	case *CellString:
		ok = stage.IsStagedCellString(target)

	case *DisplayedColumn:
		ok = stage.IsStagedDisplayedColumn(target)

	case *Row:
		ok = stage.IsStagedRow(target)

	case *SVGIcon:
		ok = stage.IsStagedSVGIcon(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedButton(button *Button) (ok bool) {

	_, ok = stage.Buttons[button]

	return
}

func (stage *Stage) IsStagedCell(cell *Cell) (ok bool) {

	_, ok = stage.Cells[cell]

	return
}

func (stage *Stage) IsStagedCellBoolean(cellboolean *CellBoolean) (ok bool) {

	_, ok = stage.CellBooleans[cellboolean]

	return
}

func (stage *Stage) IsStagedCellFloat64(cellfloat64 *CellFloat64) (ok bool) {

	_, ok = stage.CellFloat64s[cellfloat64]

	return
}

func (stage *Stage) IsStagedCellIcon(cellicon *CellIcon) (ok bool) {

	_, ok = stage.CellIcons[cellicon]

	return
}

func (stage *Stage) IsStagedCellInt(cellint *CellInt) (ok bool) {

	_, ok = stage.CellInts[cellint]

	return
}

func (stage *Stage) IsStagedCellString(cellstring *CellString) (ok bool) {

	_, ok = stage.CellStrings[cellstring]

	return
}

func (stage *Stage) IsStagedDisplayedColumn(displayedcolumn *DisplayedColumn) (ok bool) {

	_, ok = stage.DisplayedColumns[displayedcolumn]

	return
}

func (stage *Stage) IsStagedRow(row *Row) (ok bool) {

	_, ok = stage.Rows[row]

	return
}

func (stage *Stage) IsStagedSVGIcon(svgicon *SVGIcon) (ok bool) {

	_, ok = stage.SVGIcons[svgicon]

	return
}

func (stage *Stage) IsStagedTable(table *Table) (ok bool) {

	_, ok = stage.Tables[table]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Button:
		stage.StageBranchButton(target)

	case *Cell:
		stage.StageBranchCell(target)

	case *CellBoolean:
		stage.StageBranchCellBoolean(target)

	case *CellFloat64:
		stage.StageBranchCellFloat64(target)

	case *CellIcon:
		stage.StageBranchCellIcon(target)

	case *CellInt:
		stage.StageBranchCellInt(target)

	case *CellString:
		stage.StageBranchCellString(target)

	case *DisplayedColumn:
		stage.StageBranchDisplayedColumn(target)

	case *Row:
		stage.StageBranchRow(target)

	case *SVGIcon:
		stage.StageBranchSVGIcon(target)

	case *Table:
		stage.StageBranchTable(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchButton(button *Button) {

	// check if instance is already staged
	if IsStaged(stage, button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if button.SVGIcon != nil {
		StageBranch(stage, button.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCell(cell *Cell) {

	// check if instance is already staged
	if IsStaged(stage, cell) {
		return
	}

	cell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if cell.CellString != nil {
		StageBranch(stage, cell.CellString)
	}
	if cell.CellFloat64 != nil {
		StageBranch(stage, cell.CellFloat64)
	}
	if cell.CellInt != nil {
		StageBranch(stage, cell.CellInt)
	}
	if cell.CellBool != nil {
		StageBranch(stage, cell.CellBool)
	}
	if cell.CellIcon != nil {
		StageBranch(stage, cell.CellIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCellBoolean(cellboolean *CellBoolean) {

	// check if instance is already staged
	if IsStaged(stage, cellboolean) {
		return
	}

	cellboolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCellFloat64(cellfloat64 *CellFloat64) {

	// check if instance is already staged
	if IsStaged(stage, cellfloat64) {
		return
	}

	cellfloat64.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCellIcon(cellicon *CellIcon) {

	// check if instance is already staged
	if IsStaged(stage, cellicon) {
		return
	}

	cellicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCellInt(cellint *CellInt) {

	// check if instance is already staged
	if IsStaged(stage, cellint) {
		return
	}

	cellint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCellString(cellstring *CellString) {

	// check if instance is already staged
	if IsStaged(stage, cellstring) {
		return
	}

	cellstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDisplayedColumn(displayedcolumn *DisplayedColumn) {

	// check if instance is already staged
	if IsStaged(stage, displayedcolumn) {
		return
	}

	displayedcolumn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRow(row *Row) {

	// check if instance is already staged
	if IsStaged(stage, row) {
		return
	}

	row.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range row.Cells {
		StageBranch(stage, _cell)
	}

}

func (stage *Stage) StageBranchSVGIcon(svgicon *SVGIcon) {

	// check if instance is already staged
	if IsStaged(stage, svgicon) {
		return
	}

	svgicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTable(table *Table) {

	// check if instance is already staged
	if IsStaged(stage, table) {
		return
	}

	table.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _displayedcolumn := range table.DisplayedColumns {
		StageBranch(stage, _displayedcolumn)
	}
	for _, _row := range table.Rows {
		StageBranch(stage, _row)
	}
	for _, _row := range table.RowsSelectedForBulkDelete {
		StageBranch(stage, _row)
	}
	for _, _button := range table.Buttons {
		StageBranch(stage, _button)
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
	case *Button:
		toT := CopyBranchButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Cell:
		toT := CopyBranchCell(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellBoolean:
		toT := CopyBranchCellBoolean(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellFloat64:
		toT := CopyBranchCellFloat64(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellIcon:
		toT := CopyBranchCellIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellInt:
		toT := CopyBranchCellInt(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellString:
		toT := CopyBranchCellString(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DisplayedColumn:
		toT := CopyBranchDisplayedColumn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Row:
		toT := CopyBranchRow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SVGIcon:
		toT := CopyBranchSVGIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Table:
		toT := CopyBranchTable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchButton(mapOrigCopy map[any]any, buttonFrom *Button) (buttonTo *Button) {

	// buttonFrom has already been copied
	if _buttonTo, ok := mapOrigCopy[buttonFrom]; ok {
		buttonTo = _buttonTo.(*Button)
		return
	}

	buttonTo = new(Button)
	mapOrigCopy[buttonFrom] = buttonTo
	buttonFrom.CopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers
	if buttonFrom.SVGIcon != nil {
		buttonTo.SVGIcon = CopyBranchSVGIcon(mapOrigCopy, buttonFrom.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCell(mapOrigCopy map[any]any, cellFrom *Cell) (cellTo *Cell) {

	// cellFrom has already been copied
	if _cellTo, ok := mapOrigCopy[cellFrom]; ok {
		cellTo = _cellTo.(*Cell)
		return
	}

	cellTo = new(Cell)
	mapOrigCopy[cellFrom] = cellTo
	cellFrom.CopyBasicFields(cellTo)

	//insertion point for the staging of instances referenced by pointers
	if cellFrom.CellString != nil {
		cellTo.CellString = CopyBranchCellString(mapOrigCopy, cellFrom.CellString)
	}
	if cellFrom.CellFloat64 != nil {
		cellTo.CellFloat64 = CopyBranchCellFloat64(mapOrigCopy, cellFrom.CellFloat64)
	}
	if cellFrom.CellInt != nil {
		cellTo.CellInt = CopyBranchCellInt(mapOrigCopy, cellFrom.CellInt)
	}
	if cellFrom.CellBool != nil {
		cellTo.CellBool = CopyBranchCellBoolean(mapOrigCopy, cellFrom.CellBool)
	}
	if cellFrom.CellIcon != nil {
		cellTo.CellIcon = CopyBranchCellIcon(mapOrigCopy, cellFrom.CellIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCellBoolean(mapOrigCopy map[any]any, cellbooleanFrom *CellBoolean) (cellbooleanTo *CellBoolean) {

	// cellbooleanFrom has already been copied
	if _cellbooleanTo, ok := mapOrigCopy[cellbooleanFrom]; ok {
		cellbooleanTo = _cellbooleanTo.(*CellBoolean)
		return
	}

	cellbooleanTo = new(CellBoolean)
	mapOrigCopy[cellbooleanFrom] = cellbooleanTo
	cellbooleanFrom.CopyBasicFields(cellbooleanTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCellFloat64(mapOrigCopy map[any]any, cellfloat64From *CellFloat64) (cellfloat64To *CellFloat64) {

	// cellfloat64From has already been copied
	if _cellfloat64To, ok := mapOrigCopy[cellfloat64From]; ok {
		cellfloat64To = _cellfloat64To.(*CellFloat64)
		return
	}

	cellfloat64To = new(CellFloat64)
	mapOrigCopy[cellfloat64From] = cellfloat64To
	cellfloat64From.CopyBasicFields(cellfloat64To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCellIcon(mapOrigCopy map[any]any, celliconFrom *CellIcon) (celliconTo *CellIcon) {

	// celliconFrom has already been copied
	if _celliconTo, ok := mapOrigCopy[celliconFrom]; ok {
		celliconTo = _celliconTo.(*CellIcon)
		return
	}

	celliconTo = new(CellIcon)
	mapOrigCopy[celliconFrom] = celliconTo
	celliconFrom.CopyBasicFields(celliconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCellInt(mapOrigCopy map[any]any, cellintFrom *CellInt) (cellintTo *CellInt) {

	// cellintFrom has already been copied
	if _cellintTo, ok := mapOrigCopy[cellintFrom]; ok {
		cellintTo = _cellintTo.(*CellInt)
		return
	}

	cellintTo = new(CellInt)
	mapOrigCopy[cellintFrom] = cellintTo
	cellintFrom.CopyBasicFields(cellintTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCellString(mapOrigCopy map[any]any, cellstringFrom *CellString) (cellstringTo *CellString) {

	// cellstringFrom has already been copied
	if _cellstringTo, ok := mapOrigCopy[cellstringFrom]; ok {
		cellstringTo = _cellstringTo.(*CellString)
		return
	}

	cellstringTo = new(CellString)
	mapOrigCopy[cellstringFrom] = cellstringTo
	cellstringFrom.CopyBasicFields(cellstringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDisplayedColumn(mapOrigCopy map[any]any, displayedcolumnFrom *DisplayedColumn) (displayedcolumnTo *DisplayedColumn) {

	// displayedcolumnFrom has already been copied
	if _displayedcolumnTo, ok := mapOrigCopy[displayedcolumnFrom]; ok {
		displayedcolumnTo = _displayedcolumnTo.(*DisplayedColumn)
		return
	}

	displayedcolumnTo = new(DisplayedColumn)
	mapOrigCopy[displayedcolumnFrom] = displayedcolumnTo
	displayedcolumnFrom.CopyBasicFields(displayedcolumnTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRow(mapOrigCopy map[any]any, rowFrom *Row) (rowTo *Row) {

	// rowFrom has already been copied
	if _rowTo, ok := mapOrigCopy[rowFrom]; ok {
		rowTo = _rowTo.(*Row)
		return
	}

	rowTo = new(Row)
	mapOrigCopy[rowFrom] = rowTo
	rowFrom.CopyBasicFields(rowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range rowFrom.Cells {
		rowTo.Cells = append(rowTo.Cells, CopyBranchCell(mapOrigCopy, _cell))
	}

	return
}

func CopyBranchSVGIcon(mapOrigCopy map[any]any, svgiconFrom *SVGIcon) (svgiconTo *SVGIcon) {

	// svgiconFrom has already been copied
	if _svgiconTo, ok := mapOrigCopy[svgiconFrom]; ok {
		svgiconTo = _svgiconTo.(*SVGIcon)
		return
	}

	svgiconTo = new(SVGIcon)
	mapOrigCopy[svgiconFrom] = svgiconTo
	svgiconFrom.CopyBasicFields(svgiconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {

	// tableFrom has already been copied
	if _tableTo, ok := mapOrigCopy[tableFrom]; ok {
		tableTo = _tableTo.(*Table)
		return
	}

	tableTo = new(Table)
	mapOrigCopy[tableFrom] = tableTo
	tableFrom.CopyBasicFields(tableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _displayedcolumn := range tableFrom.DisplayedColumns {
		tableTo.DisplayedColumns = append(tableTo.DisplayedColumns, CopyBranchDisplayedColumn(mapOrigCopy, _displayedcolumn))
	}
	for _, _row := range tableFrom.Rows {
		tableTo.Rows = append(tableTo.Rows, CopyBranchRow(mapOrigCopy, _row))
	}
	for _, _row := range tableFrom.RowsSelectedForBulkDelete {
		tableTo.RowsSelectedForBulkDelete = append(tableTo.RowsSelectedForBulkDelete, CopyBranchRow(mapOrigCopy, _row))
	}
	for _, _button := range tableFrom.Buttons {
		tableTo.Buttons = append(tableTo.Buttons, CopyBranchButton(mapOrigCopy, _button))
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
	case *Button:
		stage.UnstageBranchButton(target)

	case *Cell:
		stage.UnstageBranchCell(target)

	case *CellBoolean:
		stage.UnstageBranchCellBoolean(target)

	case *CellFloat64:
		stage.UnstageBranchCellFloat64(target)

	case *CellIcon:
		stage.UnstageBranchCellIcon(target)

	case *CellInt:
		stage.UnstageBranchCellInt(target)

	case *CellString:
		stage.UnstageBranchCellString(target)

	case *DisplayedColumn:
		stage.UnstageBranchDisplayedColumn(target)

	case *Row:
		stage.UnstageBranchRow(target)

	case *SVGIcon:
		stage.UnstageBranchSVGIcon(target)

	case *Table:
		stage.UnstageBranchTable(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchButton(button *Button) {

	// check if instance is already staged
	if !IsStaged(stage, button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if button.SVGIcon != nil {
		UnstageBranch(stage, button.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCell(cell *Cell) {

	// check if instance is already staged
	if !IsStaged(stage, cell) {
		return
	}

	cell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if cell.CellString != nil {
		UnstageBranch(stage, cell.CellString)
	}
	if cell.CellFloat64 != nil {
		UnstageBranch(stage, cell.CellFloat64)
	}
	if cell.CellInt != nil {
		UnstageBranch(stage, cell.CellInt)
	}
	if cell.CellBool != nil {
		UnstageBranch(stage, cell.CellBool)
	}
	if cell.CellIcon != nil {
		UnstageBranch(stage, cell.CellIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCellBoolean(cellboolean *CellBoolean) {

	// check if instance is already staged
	if !IsStaged(stage, cellboolean) {
		return
	}

	cellboolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCellFloat64(cellfloat64 *CellFloat64) {

	// check if instance is already staged
	if !IsStaged(stage, cellfloat64) {
		return
	}

	cellfloat64.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCellIcon(cellicon *CellIcon) {

	// check if instance is already staged
	if !IsStaged(stage, cellicon) {
		return
	}

	cellicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCellInt(cellint *CellInt) {

	// check if instance is already staged
	if !IsStaged(stage, cellint) {
		return
	}

	cellint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCellString(cellstring *CellString) {

	// check if instance is already staged
	if !IsStaged(stage, cellstring) {
		return
	}

	cellstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDisplayedColumn(displayedcolumn *DisplayedColumn) {

	// check if instance is already staged
	if !IsStaged(stage, displayedcolumn) {
		return
	}

	displayedcolumn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRow(row *Row) {

	// check if instance is already staged
	if !IsStaged(stage, row) {
		return
	}

	row.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range row.Cells {
		UnstageBranch(stage, _cell)
	}

}

func (stage *Stage) UnstageBranchSVGIcon(svgicon *SVGIcon) {

	// check if instance is already staged
	if !IsStaged(stage, svgicon) {
		return
	}

	svgicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTable(table *Table) {

	// check if instance is already staged
	if !IsStaged(stage, table) {
		return
	}

	table.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _displayedcolumn := range table.DisplayedColumns {
		UnstageBranch(stage, _displayedcolumn)
	}
	for _, _row := range table.Rows {
		UnstageBranch(stage, _row)
	}
	for _, _row := range table.RowsSelectedForBulkDelete {
		UnstageBranch(stage, _row)
	}
	for _, _button := range table.Buttons {
		UnstageBranch(stage, _button)
	}

}

// insertion point for pointer reconstruction from references
func (reference *Button) GongReconstructPointersFromReferences(stage *Stage, instance *Button) {
	// insertion point for pointers field
	if instance.SVGIcon != nil {
		reference.SVGIcon = stage.SVGIcons_reference[instance.SVGIcon]
	}
	// insertion point for slice of pointers field
}

func (reference *Cell) GongReconstructPointersFromReferences(stage *Stage, instance *Cell) {
	// insertion point for pointers field
	if instance.CellString != nil {
		reference.CellString = stage.CellStrings_reference[instance.CellString]
	}
	if instance.CellFloat64 != nil {
		reference.CellFloat64 = stage.CellFloat64s_reference[instance.CellFloat64]
	}
	if instance.CellInt != nil {
		reference.CellInt = stage.CellInts_reference[instance.CellInt]
	}
	if instance.CellBool != nil {
		reference.CellBool = stage.CellBooleans_reference[instance.CellBool]
	}
	if instance.CellIcon != nil {
		reference.CellIcon = stage.CellIcons_reference[instance.CellIcon]
	}
	// insertion point for slice of pointers field
}

func (reference *CellBoolean) GongReconstructPointersFromReferences(stage *Stage, instance *CellBoolean) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CellFloat64) GongReconstructPointersFromReferences(stage *Stage, instance *CellFloat64) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CellIcon) GongReconstructPointersFromReferences(stage *Stage, instance *CellIcon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CellInt) GongReconstructPointersFromReferences(stage *Stage, instance *CellInt) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CellString) GongReconstructPointersFromReferences(stage *Stage, instance *CellString) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DisplayedColumn) GongReconstructPointersFromReferences(stage *Stage, instance *DisplayedColumn) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Row) GongReconstructPointersFromReferences(stage *Stage, instance *Row) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Cells = reference.Cells[:0]
	for _, _b := range instance.Cells {
		reference.Cells = append(reference.Cells, stage.Cells_reference[_b])
	}
}

func (reference *SVGIcon) GongReconstructPointersFromReferences(stage *Stage, instance *SVGIcon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Table) GongReconstructPointersFromReferences(stage *Stage, instance *Table) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.DisplayedColumns = reference.DisplayedColumns[:0]
	for _, _b := range instance.DisplayedColumns {
		reference.DisplayedColumns = append(reference.DisplayedColumns, stage.DisplayedColumns_reference[_b])
	}
	reference.Rows = reference.Rows[:0]
	for _, _b := range instance.Rows {
		reference.Rows = append(reference.Rows, stage.Rows_reference[_b])
	}
	reference.RowsSelectedForBulkDelete = reference.RowsSelectedForBulkDelete[:0]
	for _, _b := range instance.RowsSelectedForBulkDelete {
		reference.RowsSelectedForBulkDelete = append(reference.RowsSelectedForBulkDelete, stage.Rows_reference[_b])
	}
	reference.Buttons = reference.Buttons[:0]
	for _, _b := range instance.Buttons {
		reference.Buttons = append(reference.Buttons, stage.Buttons_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *Button) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.SVGIcon; _reference != nil {
		reference.SVGIcon = nil
		if _instance, ok := stage.SVGIcons_instance[_reference]; ok {
			reference.SVGIcon = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Cell) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.CellString; _reference != nil {
		reference.CellString = nil
		if _instance, ok := stage.CellStrings_instance[_reference]; ok {
			reference.CellString = _instance
		}
	}
	if _reference := reference.CellFloat64; _reference != nil {
		reference.CellFloat64 = nil
		if _instance, ok := stage.CellFloat64s_instance[_reference]; ok {
			reference.CellFloat64 = _instance
		}
	}
	if _reference := reference.CellInt; _reference != nil {
		reference.CellInt = nil
		if _instance, ok := stage.CellInts_instance[_reference]; ok {
			reference.CellInt = _instance
		}
	}
	if _reference := reference.CellBool; _reference != nil {
		reference.CellBool = nil
		if _instance, ok := stage.CellBooleans_instance[_reference]; ok {
			reference.CellBool = _instance
		}
	}
	if _reference := reference.CellIcon; _reference != nil {
		reference.CellIcon = nil
		if _instance, ok := stage.CellIcons_instance[_reference]; ok {
			reference.CellIcon = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *CellBoolean) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CellFloat64) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CellIcon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CellInt) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CellString) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DisplayedColumn) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Row) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Cells []*Cell
	for _, _reference := range reference.Cells {
		if _instance, ok := stage.Cells_instance[_reference]; ok {
			_Cells = append(_Cells, _instance)
		}
	}
	reference.Cells = _Cells
}

func (reference *SVGIcon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Table) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _DisplayedColumns []*DisplayedColumn
	for _, _reference := range reference.DisplayedColumns {
		if _instance, ok := stage.DisplayedColumns_instance[_reference]; ok {
			_DisplayedColumns = append(_DisplayedColumns, _instance)
		}
	}
	reference.DisplayedColumns = _DisplayedColumns
	var _Rows []*Row
	for _, _reference := range reference.Rows {
		if _instance, ok := stage.Rows_instance[_reference]; ok {
			_Rows = append(_Rows, _instance)
		}
	}
	reference.Rows = _Rows
	var _RowsSelectedForBulkDelete []*Row
	for _, _reference := range reference.RowsSelectedForBulkDelete {
		if _instance, ok := stage.Rows_instance[_reference]; ok {
			_RowsSelectedForBulkDelete = append(_RowsSelectedForBulkDelete, _instance)
		}
	}
	reference.RowsSelectedForBulkDelete = _RowsSelectedForBulkDelete
	var _Buttons []*Button
	for _, _reference := range reference.Buttons {
		if _instance, ok := stage.Buttons_instance[_reference]; ok {
			_Buttons = append(_Buttons, _instance)
		}
	}
	reference.Buttons = _Buttons
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (button *Button) GongDiff(stage *Stage, buttonOther *Button) (diffs []string) {
	// insertion point for field diffs
	if button.Name != buttonOther.Name {
		diffs = append(diffs, button.GongMarshallField(stage, "Name"))
	}
	if button.Icon != buttonOther.Icon {
		diffs = append(diffs, button.GongMarshallField(stage, "Icon"))
	}
	if (button.SVGIcon == nil) != (buttonOther.SVGIcon == nil) {
		diffs = append(diffs, button.GongMarshallField(stage, "SVGIcon"))
	} else if button.SVGIcon != nil && buttonOther.SVGIcon != nil {
		if button.SVGIcon != buttonOther.SVGIcon {
			diffs = append(diffs, button.GongMarshallField(stage, "SVGIcon"))
		}
	}
	if button.IsDisabled != buttonOther.IsDisabled {
		diffs = append(diffs, button.GongMarshallField(stage, "IsDisabled"))
	}
	if button.HasToolTip != buttonOther.HasToolTip {
		diffs = append(diffs, button.GongMarshallField(stage, "HasToolTip"))
	}
	if button.ToolTipText != buttonOther.ToolTipText {
		diffs = append(diffs, button.GongMarshallField(stage, "ToolTipText"))
	}
	if button.ToolTipPosition != buttonOther.ToolTipPosition {
		diffs = append(diffs, button.GongMarshallField(stage, "ToolTipPosition"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cell *Cell) GongDiff(stage *Stage, cellOther *Cell) (diffs []string) {
	// insertion point for field diffs
	if cell.Name != cellOther.Name {
		diffs = append(diffs, cell.GongMarshallField(stage, "Name"))
	}
	if (cell.CellString == nil) != (cellOther.CellString == nil) {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellString"))
	} else if cell.CellString != nil && cellOther.CellString != nil {
		if cell.CellString != cellOther.CellString {
			diffs = append(diffs, cell.GongMarshallField(stage, "CellString"))
		}
	}
	if (cell.CellFloat64 == nil) != (cellOther.CellFloat64 == nil) {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellFloat64"))
	} else if cell.CellFloat64 != nil && cellOther.CellFloat64 != nil {
		if cell.CellFloat64 != cellOther.CellFloat64 {
			diffs = append(diffs, cell.GongMarshallField(stage, "CellFloat64"))
		}
	}
	if (cell.CellInt == nil) != (cellOther.CellInt == nil) {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellInt"))
	} else if cell.CellInt != nil && cellOther.CellInt != nil {
		if cell.CellInt != cellOther.CellInt {
			diffs = append(diffs, cell.GongMarshallField(stage, "CellInt"))
		}
	}
	if (cell.CellBool == nil) != (cellOther.CellBool == nil) {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellBool"))
	} else if cell.CellBool != nil && cellOther.CellBool != nil {
		if cell.CellBool != cellOther.CellBool {
			diffs = append(diffs, cell.GongMarshallField(stage, "CellBool"))
		}
	}
	if (cell.CellIcon == nil) != (cellOther.CellIcon == nil) {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellIcon"))
	} else if cell.CellIcon != nil && cellOther.CellIcon != nil {
		if cell.CellIcon != cellOther.CellIcon {
			diffs = append(diffs, cell.GongMarshallField(stage, "CellIcon"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cellboolean *CellBoolean) GongDiff(stage *Stage, cellbooleanOther *CellBoolean) (diffs []string) {
	// insertion point for field diffs
	if cellboolean.Name != cellbooleanOther.Name {
		diffs = append(diffs, cellboolean.GongMarshallField(stage, "Name"))
	}
	if cellboolean.Value != cellbooleanOther.Value {
		diffs = append(diffs, cellboolean.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cellfloat64 *CellFloat64) GongDiff(stage *Stage, cellfloat64Other *CellFloat64) (diffs []string) {
	// insertion point for field diffs
	if cellfloat64.Name != cellfloat64Other.Name {
		diffs = append(diffs, cellfloat64.GongMarshallField(stage, "Name"))
	}
	if cellfloat64.Value != cellfloat64Other.Value {
		diffs = append(diffs, cellfloat64.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cellicon *CellIcon) GongDiff(stage *Stage, celliconOther *CellIcon) (diffs []string) {
	// insertion point for field diffs
	if cellicon.Name != celliconOther.Name {
		diffs = append(diffs, cellicon.GongMarshallField(stage, "Name"))
	}
	if cellicon.Icon != celliconOther.Icon {
		diffs = append(diffs, cellicon.GongMarshallField(stage, "Icon"))
	}
	if cellicon.NeedsConfirmation != celliconOther.NeedsConfirmation {
		diffs = append(diffs, cellicon.GongMarshallField(stage, "NeedsConfirmation"))
	}
	if cellicon.ConfirmationMessage != celliconOther.ConfirmationMessage {
		diffs = append(diffs, cellicon.GongMarshallField(stage, "ConfirmationMessage"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cellint *CellInt) GongDiff(stage *Stage, cellintOther *CellInt) (diffs []string) {
	// insertion point for field diffs
	if cellint.Name != cellintOther.Name {
		diffs = append(diffs, cellint.GongMarshallField(stage, "Name"))
	}
	if cellint.Value != cellintOther.Value {
		diffs = append(diffs, cellint.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cellstring *CellString) GongDiff(stage *Stage, cellstringOther *CellString) (diffs []string) {
	// insertion point for field diffs
	if cellstring.Name != cellstringOther.Name {
		diffs = append(diffs, cellstring.GongMarshallField(stage, "Name"))
	}
	if cellstring.Value != cellstringOther.Value {
		diffs = append(diffs, cellstring.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (displayedcolumn *DisplayedColumn) GongDiff(stage *Stage, displayedcolumnOther *DisplayedColumn) (diffs []string) {
	// insertion point for field diffs
	if displayedcolumn.Name != displayedcolumnOther.Name {
		diffs = append(diffs, displayedcolumn.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (row *Row) GongDiff(stage *Stage, rowOther *Row) (diffs []string) {
	// insertion point for field diffs
	if row.Name != rowOther.Name {
		diffs = append(diffs, row.GongMarshallField(stage, "Name"))
	}
	CellsDifferent := false
	if len(row.Cells) != len(rowOther.Cells) {
		CellsDifferent = true
	} else {
		for i := range row.Cells {
			if (row.Cells[i] == nil) != (rowOther.Cells[i] == nil) {
				CellsDifferent = true
				break
			} else if row.Cells[i] != nil && rowOther.Cells[i] != nil {
				// this is a pointer comparaison
				if row.Cells[i] != rowOther.Cells[i] {
					CellsDifferent = true
					break
				}
			}
		}
	}
	if CellsDifferent {
		ops := Diff(stage, row, rowOther, "Cells", rowOther.Cells, row.Cells)
		diffs = append(diffs, ops)
	}
	if row.IsChecked != rowOther.IsChecked {
		diffs = append(diffs, row.GongMarshallField(stage, "IsChecked"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (svgicon *SVGIcon) GongDiff(stage *Stage, svgiconOther *SVGIcon) (diffs []string) {
	// insertion point for field diffs
	if svgicon.Name != svgiconOther.Name {
		diffs = append(diffs, svgicon.GongMarshallField(stage, "Name"))
	}
	if svgicon.SVG != svgiconOther.SVG {
		diffs = append(diffs, svgicon.GongMarshallField(stage, "SVG"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (table *Table) GongDiff(stage *Stage, tableOther *Table) (diffs []string) {
	// insertion point for field diffs
	if table.Name != tableOther.Name {
		diffs = append(diffs, table.GongMarshallField(stage, "Name"))
	}
	DisplayedColumnsDifferent := false
	if len(table.DisplayedColumns) != len(tableOther.DisplayedColumns) {
		DisplayedColumnsDifferent = true
	} else {
		for i := range table.DisplayedColumns {
			if (table.DisplayedColumns[i] == nil) != (tableOther.DisplayedColumns[i] == nil) {
				DisplayedColumnsDifferent = true
				break
			} else if table.DisplayedColumns[i] != nil && tableOther.DisplayedColumns[i] != nil {
				// this is a pointer comparaison
				if table.DisplayedColumns[i] != tableOther.DisplayedColumns[i] {
					DisplayedColumnsDifferent = true
					break
				}
			}
		}
	}
	if DisplayedColumnsDifferent {
		ops := Diff(stage, table, tableOther, "DisplayedColumns", tableOther.DisplayedColumns, table.DisplayedColumns)
		diffs = append(diffs, ops)
	}
	RowsDifferent := false
	if len(table.Rows) != len(tableOther.Rows) {
		RowsDifferent = true
	} else {
		for i := range table.Rows {
			if (table.Rows[i] == nil) != (tableOther.Rows[i] == nil) {
				RowsDifferent = true
				break
			} else if table.Rows[i] != nil && tableOther.Rows[i] != nil {
				// this is a pointer comparaison
				if table.Rows[i] != tableOther.Rows[i] {
					RowsDifferent = true
					break
				}
			}
		}
	}
	if RowsDifferent {
		ops := Diff(stage, table, tableOther, "Rows", tableOther.Rows, table.Rows)
		diffs = append(diffs, ops)
	}
	if table.HasFiltering != tableOther.HasFiltering {
		diffs = append(diffs, table.GongMarshallField(stage, "HasFiltering"))
	}
	if table.HasColumnSorting != tableOther.HasColumnSorting {
		diffs = append(diffs, table.GongMarshallField(stage, "HasColumnSorting"))
	}
	if table.HasPaginator != tableOther.HasPaginator {
		diffs = append(diffs, table.GongMarshallField(stage, "HasPaginator"))
	}
	if table.HasCheckableRows != tableOther.HasCheckableRows {
		diffs = append(diffs, table.GongMarshallField(stage, "HasCheckableRows"))
	}
	if table.HasSaveButton != tableOther.HasSaveButton {
		diffs = append(diffs, table.GongMarshallField(stage, "HasSaveButton"))
	}
	if table.SaveButtonLabel != tableOther.SaveButtonLabel {
		diffs = append(diffs, table.GongMarshallField(stage, "SaveButtonLabel"))
	}
	if table.HasBulkDeleteButton != tableOther.HasBulkDeleteButton {
		diffs = append(diffs, table.GongMarshallField(stage, "HasBulkDeleteButton"))
	}
	if table.BulkDeleteButtonTooltip != tableOther.BulkDeleteButtonTooltip {
		diffs = append(diffs, table.GongMarshallField(stage, "BulkDeleteButtonTooltip"))
	}
	RowsSelectedForBulkDeleteDifferent := false
	if len(table.RowsSelectedForBulkDelete) != len(tableOther.RowsSelectedForBulkDelete) {
		RowsSelectedForBulkDeleteDifferent = true
	} else {
		for i := range table.RowsSelectedForBulkDelete {
			if (table.RowsSelectedForBulkDelete[i] == nil) != (tableOther.RowsSelectedForBulkDelete[i] == nil) {
				RowsSelectedForBulkDeleteDifferent = true
				break
			} else if table.RowsSelectedForBulkDelete[i] != nil && tableOther.RowsSelectedForBulkDelete[i] != nil {
				// this is a pointer comparaison
				if table.RowsSelectedForBulkDelete[i] != tableOther.RowsSelectedForBulkDelete[i] {
					RowsSelectedForBulkDeleteDifferent = true
					break
				}
			}
		}
	}
	if RowsSelectedForBulkDeleteDifferent {
		ops := Diff(stage, table, tableOther, "RowsSelectedForBulkDelete", tableOther.RowsSelectedForBulkDelete, table.RowsSelectedForBulkDelete)
		diffs = append(diffs, ops)
	}
	if table.CanDragDropRows != tableOther.CanDragDropRows {
		diffs = append(diffs, table.GongMarshallField(stage, "CanDragDropRows"))
	}
	if table.HasCloseButton != tableOther.HasCloseButton {
		diffs = append(diffs, table.GongMarshallField(stage, "HasCloseButton"))
	}
	if table.SavingInProgress != tableOther.SavingInProgress {
		diffs = append(diffs, table.GongMarshallField(stage, "SavingInProgress"))
	}
	if table.NbOfStickyColumns != tableOther.NbOfStickyColumns {
		diffs = append(diffs, table.GongMarshallField(stage, "NbOfStickyColumns"))
	}
	ButtonsDifferent := false
	if len(table.Buttons) != len(tableOther.Buttons) {
		ButtonsDifferent = true
	} else {
		for i := range table.Buttons {
			if (table.Buttons[i] == nil) != (tableOther.Buttons[i] == nil) {
				ButtonsDifferent = true
				break
			} else if table.Buttons[i] != nil && tableOther.Buttons[i] != nil {
				// this is a pointer comparaison
				if table.Buttons[i] != tableOther.Buttons[i] {
					ButtonsDifferent = true
					break
				}
			}
		}
	}
	if ButtonsDifferent {
		ops := Diff(stage, table, tableOther, "Buttons", tableOther.Buttons, table.Buttons)
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
