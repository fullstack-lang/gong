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
func (button *Button) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Buttons[button]

	return
}

func (stage *Stage) IsStagedButton(button *Button) (ok bool) {

	return button.GongIsStaged(stage)
}

func (cell *Cell) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Cells[cell]

	return
}

func (stage *Stage) IsStagedCell(cell *Cell) (ok bool) {

	return cell.GongIsStaged(stage)
}

func (cellboolean *CellBoolean) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.CellBooleans[cellboolean]

	return
}

func (stage *Stage) IsStagedCellBoolean(cellboolean *CellBoolean) (ok bool) {

	return cellboolean.GongIsStaged(stage)
}

func (cellfloat64 *CellFloat64) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.CellFloat64s[cellfloat64]

	return
}

func (stage *Stage) IsStagedCellFloat64(cellfloat64 *CellFloat64) (ok bool) {

	return cellfloat64.GongIsStaged(stage)
}

func (cellicon *CellIcon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.CellIcons[cellicon]

	return
}

func (stage *Stage) IsStagedCellIcon(cellicon *CellIcon) (ok bool) {

	return cellicon.GongIsStaged(stage)
}

func (cellint *CellInt) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.CellInts[cellint]

	return
}

func (stage *Stage) IsStagedCellInt(cellint *CellInt) (ok bool) {

	return cellint.GongIsStaged(stage)
}

func (cellstring *CellString) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.CellStrings[cellstring]

	return
}

func (stage *Stage) IsStagedCellString(cellstring *CellString) (ok bool) {

	return cellstring.GongIsStaged(stage)
}

func (displayedcolumn *DisplayedColumn) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DisplayedColumns[displayedcolumn]

	return
}

func (stage *Stage) IsStagedDisplayedColumn(displayedcolumn *DisplayedColumn) (ok bool) {

	return displayedcolumn.GongIsStaged(stage)
}

func (row *Row) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Rows[row]

	return
}

func (stage *Stage) IsStagedRow(row *Row) (ok bool) {

	return row.GongIsStaged(stage)
}

func (svgicon *SVGIcon) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SVGIcons[svgicon]

	return
}

func (stage *Stage) IsStagedSVGIcon(svgicon *SVGIcon) (ok bool) {

	return svgicon.GongIsStaged(stage)
}

func (table *Table) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Tables[table]

	return
}

func (stage *Stage) IsStagedTable(table *Table) (ok bool) {

	return table.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (button *Button) GongStageBranch(stage *Stage) {
	stage.StageBranchButton(button)
}

func (stage *Stage) StageBranchButton(button *Button) {

	// check if instance is already staged
	if stage.IsStaged(button) {
		return
	}

	button.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if button.SVGIcon != nil {
		stage.StageBranch(button.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cell *Cell) GongStageBranch(stage *Stage) {
	stage.StageBranchCell(cell)
}

func (stage *Stage) StageBranchCell(cell *Cell) {

	// check if instance is already staged
	if stage.IsStaged(cell) {
		return
	}

	cell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if cell.CellString != nil {
		stage.StageBranch(cell.CellString)
	}
	if cell.CellFloat64 != nil {
		stage.StageBranch(cell.CellFloat64)
	}
	if cell.CellInt != nil {
		stage.StageBranch(cell.CellInt)
	}
	if cell.CellBool != nil {
		stage.StageBranch(cell.CellBool)
	}
	if cell.CellIcon != nil {
		stage.StageBranch(cell.CellIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellboolean *CellBoolean) GongStageBranch(stage *Stage) {
	stage.StageBranchCellBoolean(cellboolean)
}

func (stage *Stage) StageBranchCellBoolean(cellboolean *CellBoolean) {

	// check if instance is already staged
	if stage.IsStaged(cellboolean) {
		return
	}

	cellboolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellfloat64 *CellFloat64) GongStageBranch(stage *Stage) {
	stage.StageBranchCellFloat64(cellfloat64)
}

func (stage *Stage) StageBranchCellFloat64(cellfloat64 *CellFloat64) {

	// check if instance is already staged
	if stage.IsStaged(cellfloat64) {
		return
	}

	cellfloat64.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellicon *CellIcon) GongStageBranch(stage *Stage) {
	stage.StageBranchCellIcon(cellicon)
}

func (stage *Stage) StageBranchCellIcon(cellicon *CellIcon) {

	// check if instance is already staged
	if stage.IsStaged(cellicon) {
		return
	}

	cellicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellint *CellInt) GongStageBranch(stage *Stage) {
	stage.StageBranchCellInt(cellint)
}

func (stage *Stage) StageBranchCellInt(cellint *CellInt) {

	// check if instance is already staged
	if stage.IsStaged(cellint) {
		return
	}

	cellint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellstring *CellString) GongStageBranch(stage *Stage) {
	stage.StageBranchCellString(cellstring)
}

func (stage *Stage) StageBranchCellString(cellstring *CellString) {

	// check if instance is already staged
	if stage.IsStaged(cellstring) {
		return
	}

	cellstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (displayedcolumn *DisplayedColumn) GongStageBranch(stage *Stage) {
	stage.StageBranchDisplayedColumn(displayedcolumn)
}

func (stage *Stage) StageBranchDisplayedColumn(displayedcolumn *DisplayedColumn) {

	// check if instance is already staged
	if stage.IsStaged(displayedcolumn) {
		return
	}

	displayedcolumn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (row *Row) GongStageBranch(stage *Stage) {
	stage.StageBranchRow(row)
}

func (stage *Stage) StageBranchRow(row *Row) {

	// check if instance is already staged
	if stage.IsStaged(row) {
		return
	}

	row.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range row.Cells {
		stage.StageBranch(_cell)
	}

}

func (svgicon *SVGIcon) GongStageBranch(stage *Stage) {
	stage.StageBranchSVGIcon(svgicon)
}

func (stage *Stage) StageBranchSVGIcon(svgicon *SVGIcon) {

	// check if instance is already staged
	if stage.IsStaged(svgicon) {
		return
	}

	svgicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (table *Table) GongStageBranch(stage *Stage) {
	stage.StageBranchTable(table)
}

func (stage *Stage) StageBranchTable(table *Table) {

	// check if instance is already staged
	if stage.IsStaged(table) {
		return
	}

	table.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _displayedcolumn := range table.DisplayedColumns {
		stage.StageBranch(_displayedcolumn)
	}
	for _, _row := range table.Rows {
		stage.StageBranch(_row)
	}
	for _, _row := range table.RowsSelectedForBulkDelete {
		stage.StageBranch(_row)
	}
	for _, _button := range table.Buttons {
		stage.StageBranch(_button)
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
	case *Button:
		toT := GongCopyBranchButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Cell:
		toT := GongCopyBranchCell(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellBoolean:
		toT := GongCopyBranchCellBoolean(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellFloat64:
		toT := GongCopyBranchCellFloat64(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellIcon:
		toT := GongCopyBranchCellIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellInt:
		toT := GongCopyBranchCellInt(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CellString:
		toT := GongCopyBranchCellString(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DisplayedColumn:
		toT := GongCopyBranchDisplayedColumn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Row:
		toT := GongCopyBranchRow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SVGIcon:
		toT := GongCopyBranchSVGIcon(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Table:
		toT := GongCopyBranchTable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchButton(mapOrigCopy map[any]any, buttonFrom *Button) (buttonTo *Button) {

	// buttonFrom has already been copied
	if _buttonTo, ok := mapOrigCopy[buttonFrom]; ok {
		buttonTo = _buttonTo.(*Button)
		return
	}

	buttonTo = new(Button)
	mapOrigCopy[buttonFrom] = buttonTo
	buttonFrom.GongCopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers
	if buttonFrom.SVGIcon != nil {
		buttonTo.SVGIcon = GongCopyBranchSVGIcon(mapOrigCopy, buttonFrom.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCell(mapOrigCopy map[any]any, cellFrom *Cell) (cellTo *Cell) {

	// cellFrom has already been copied
	if _cellTo, ok := mapOrigCopy[cellFrom]; ok {
		cellTo = _cellTo.(*Cell)
		return
	}

	cellTo = new(Cell)
	mapOrigCopy[cellFrom] = cellTo
	cellFrom.GongCopyBasicFields(cellTo)

	//insertion point for the staging of instances referenced by pointers
	if cellFrom.CellString != nil {
		cellTo.CellString = GongCopyBranchCellString(mapOrigCopy, cellFrom.CellString)
	}
	if cellFrom.CellFloat64 != nil {
		cellTo.CellFloat64 = GongCopyBranchCellFloat64(mapOrigCopy, cellFrom.CellFloat64)
	}
	if cellFrom.CellInt != nil {
		cellTo.CellInt = GongCopyBranchCellInt(mapOrigCopy, cellFrom.CellInt)
	}
	if cellFrom.CellBool != nil {
		cellTo.CellBool = GongCopyBranchCellBoolean(mapOrigCopy, cellFrom.CellBool)
	}
	if cellFrom.CellIcon != nil {
		cellTo.CellIcon = GongCopyBranchCellIcon(mapOrigCopy, cellFrom.CellIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellBoolean(mapOrigCopy map[any]any, cellbooleanFrom *CellBoolean) (cellbooleanTo *CellBoolean) {

	// cellbooleanFrom has already been copied
	if _cellbooleanTo, ok := mapOrigCopy[cellbooleanFrom]; ok {
		cellbooleanTo = _cellbooleanTo.(*CellBoolean)
		return
	}

	cellbooleanTo = new(CellBoolean)
	mapOrigCopy[cellbooleanFrom] = cellbooleanTo
	cellbooleanFrom.GongCopyBasicFields(cellbooleanTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellFloat64(mapOrigCopy map[any]any, cellfloat64From *CellFloat64) (cellfloat64To *CellFloat64) {

	// cellfloat64From has already been copied
	if _cellfloat64To, ok := mapOrigCopy[cellfloat64From]; ok {
		cellfloat64To = _cellfloat64To.(*CellFloat64)
		return
	}

	cellfloat64To = new(CellFloat64)
	mapOrigCopy[cellfloat64From] = cellfloat64To
	cellfloat64From.GongCopyBasicFields(cellfloat64To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellIcon(mapOrigCopy map[any]any, celliconFrom *CellIcon) (celliconTo *CellIcon) {

	// celliconFrom has already been copied
	if _celliconTo, ok := mapOrigCopy[celliconFrom]; ok {
		celliconTo = _celliconTo.(*CellIcon)
		return
	}

	celliconTo = new(CellIcon)
	mapOrigCopy[celliconFrom] = celliconTo
	celliconFrom.GongCopyBasicFields(celliconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellInt(mapOrigCopy map[any]any, cellintFrom *CellInt) (cellintTo *CellInt) {

	// cellintFrom has already been copied
	if _cellintTo, ok := mapOrigCopy[cellintFrom]; ok {
		cellintTo = _cellintTo.(*CellInt)
		return
	}

	cellintTo = new(CellInt)
	mapOrigCopy[cellintFrom] = cellintTo
	cellintFrom.GongCopyBasicFields(cellintTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellString(mapOrigCopy map[any]any, cellstringFrom *CellString) (cellstringTo *CellString) {

	// cellstringFrom has already been copied
	if _cellstringTo, ok := mapOrigCopy[cellstringFrom]; ok {
		cellstringTo = _cellstringTo.(*CellString)
		return
	}

	cellstringTo = new(CellString)
	mapOrigCopy[cellstringFrom] = cellstringTo
	cellstringFrom.GongCopyBasicFields(cellstringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDisplayedColumn(mapOrigCopy map[any]any, displayedcolumnFrom *DisplayedColumn) (displayedcolumnTo *DisplayedColumn) {

	// displayedcolumnFrom has already been copied
	if _displayedcolumnTo, ok := mapOrigCopy[displayedcolumnFrom]; ok {
		displayedcolumnTo = _displayedcolumnTo.(*DisplayedColumn)
		return
	}

	displayedcolumnTo = new(DisplayedColumn)
	mapOrigCopy[displayedcolumnFrom] = displayedcolumnTo
	displayedcolumnFrom.GongCopyBasicFields(displayedcolumnTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRow(mapOrigCopy map[any]any, rowFrom *Row) (rowTo *Row) {

	// rowFrom has already been copied
	if _rowTo, ok := mapOrigCopy[rowFrom]; ok {
		rowTo = _rowTo.(*Row)
		return
	}

	rowTo = new(Row)
	mapOrigCopy[rowFrom] = rowTo
	rowFrom.GongCopyBasicFields(rowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range rowFrom.Cells {
		rowTo.Cells = append(rowTo.Cells, GongCopyBranchCell(mapOrigCopy, _cell))
	}

	return
}

func GongCopyBranchSVGIcon(mapOrigCopy map[any]any, svgiconFrom *SVGIcon) (svgiconTo *SVGIcon) {

	// svgiconFrom has already been copied
	if _svgiconTo, ok := mapOrigCopy[svgiconFrom]; ok {
		svgiconTo = _svgiconTo.(*SVGIcon)
		return
	}

	svgiconTo = new(SVGIcon)
	mapOrigCopy[svgiconFrom] = svgiconTo
	svgiconFrom.GongCopyBasicFields(svgiconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {

	// tableFrom has already been copied
	if _tableTo, ok := mapOrigCopy[tableFrom]; ok {
		tableTo = _tableTo.(*Table)
		return
	}

	tableTo = new(Table)
	mapOrigCopy[tableFrom] = tableTo
	tableFrom.GongCopyBasicFields(tableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _displayedcolumn := range tableFrom.DisplayedColumns {
		tableTo.DisplayedColumns = append(tableTo.DisplayedColumns, GongCopyBranchDisplayedColumn(mapOrigCopy, _displayedcolumn))
	}
	for _, _row := range tableFrom.Rows {
		tableTo.Rows = append(tableTo.Rows, GongCopyBranchRow(mapOrigCopy, _row))
	}
	for _, _row := range tableFrom.RowsSelectedForBulkDelete {
		tableTo.RowsSelectedForBulkDelete = append(tableTo.RowsSelectedForBulkDelete, GongCopyBranchRow(mapOrigCopy, _row))
	}
	for _, _button := range tableFrom.Buttons {
		tableTo.Buttons = append(tableTo.Buttons, GongCopyBranchButton(mapOrigCopy, _button))
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
func (button *Button) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchButton(button)
}

func (stage *Stage) UnstageBranchButton(button *Button) {

	// check if instance is already staged
	if !stage.IsStaged(button) {
		return
	}

	button.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if button.SVGIcon != nil {
		stage.UnstageBranch(button.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cell *Cell) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCell(cell)
}

func (stage *Stage) UnstageBranchCell(cell *Cell) {

	// check if instance is already staged
	if !stage.IsStaged(cell) {
		return
	}

	cell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if cell.CellString != nil {
		stage.UnstageBranch(cell.CellString)
	}
	if cell.CellFloat64 != nil {
		stage.UnstageBranch(cell.CellFloat64)
	}
	if cell.CellInt != nil {
		stage.UnstageBranch(cell.CellInt)
	}
	if cell.CellBool != nil {
		stage.UnstageBranch(cell.CellBool)
	}
	if cell.CellIcon != nil {
		stage.UnstageBranch(cell.CellIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellboolean *CellBoolean) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCellBoolean(cellboolean)
}

func (stage *Stage) UnstageBranchCellBoolean(cellboolean *CellBoolean) {

	// check if instance is already staged
	if !stage.IsStaged(cellboolean) {
		return
	}

	cellboolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellfloat64 *CellFloat64) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCellFloat64(cellfloat64)
}

func (stage *Stage) UnstageBranchCellFloat64(cellfloat64 *CellFloat64) {

	// check if instance is already staged
	if !stage.IsStaged(cellfloat64) {
		return
	}

	cellfloat64.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellicon *CellIcon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCellIcon(cellicon)
}

func (stage *Stage) UnstageBranchCellIcon(cellicon *CellIcon) {

	// check if instance is already staged
	if !stage.IsStaged(cellicon) {
		return
	}

	cellicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellint *CellInt) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCellInt(cellint)
}

func (stage *Stage) UnstageBranchCellInt(cellint *CellInt) {

	// check if instance is already staged
	if !stage.IsStaged(cellint) {
		return
	}

	cellint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellstring *CellString) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCellString(cellstring)
}

func (stage *Stage) UnstageBranchCellString(cellstring *CellString) {

	// check if instance is already staged
	if !stage.IsStaged(cellstring) {
		return
	}

	cellstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (displayedcolumn *DisplayedColumn) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDisplayedColumn(displayedcolumn)
}

func (stage *Stage) UnstageBranchDisplayedColumn(displayedcolumn *DisplayedColumn) {

	// check if instance is already staged
	if !stage.IsStaged(displayedcolumn) {
		return
	}

	displayedcolumn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (row *Row) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRow(row)
}

func (stage *Stage) UnstageBranchRow(row *Row) {

	// check if instance is already staged
	if !stage.IsStaged(row) {
		return
	}

	row.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range row.Cells {
		stage.UnstageBranch(_cell)
	}

}

func (svgicon *SVGIcon) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSVGIcon(svgicon)
}

func (stage *Stage) UnstageBranchSVGIcon(svgicon *SVGIcon) {

	// check if instance is already staged
	if !stage.IsStaged(svgicon) {
		return
	}

	svgicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (table *Table) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTable(table)
}

func (stage *Stage) UnstageBranchTable(table *Table) {

	// check if instance is already staged
	if !stage.IsStaged(table) {
		return
	}

	table.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _displayedcolumn := range table.DisplayedColumns {
		stage.UnstageBranch(_displayedcolumn)
	}
	for _, _row := range table.Rows {
		stage.UnstageBranch(_row)
	}
	for _, _row := range table.RowsSelectedForBulkDelete {
		stage.UnstageBranch(_row)
	}
	for _, _button := range table.Buttons {
		stage.UnstageBranch(_button)
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
		ops := stage.Diff(
			row,
			"Cells",
			len(rowOther.Cells),
			len(row.Cells),
			func(i, j int) bool {
				return rowOther.Cells[i] == row.Cells[j]
			},
			func(j int) string {
				return row.Cells[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			table,
			"DisplayedColumns",
			len(tableOther.DisplayedColumns),
			len(table.DisplayedColumns),
			func(i, j int) bool {
				return tableOther.DisplayedColumns[i] == table.DisplayedColumns[j]
			},
			func(j int) string {
				return table.DisplayedColumns[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			table,
			"Rows",
			len(tableOther.Rows),
			len(table.Rows),
			func(i, j int) bool {
				return tableOther.Rows[i] == table.Rows[j]
			},
			func(j int) string {
				return table.Rows[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			table,
			"RowsSelectedForBulkDelete",
			len(tableOther.RowsSelectedForBulkDelete),
			len(table.RowsSelectedForBulkDelete),
			func(i, j int) bool {
				return tableOther.RowsSelectedForBulkDelete[i] == table.RowsSelectedForBulkDelete[j]
			},
			func(j int) string {
				return table.RowsSelectedForBulkDelete[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			table,
			"Buttons",
			len(tableOther.Buttons),
			len(table.Buttons),
			func(i, j int) bool {
				return tableOther.Buttons[i] == table.Buttons[j]
			},
			func(j int) string {
				return table.Buttons[j].GongGetIdentifier(stage)
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
