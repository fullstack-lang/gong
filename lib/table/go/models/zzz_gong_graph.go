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
func (button *Button) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Buttons[button]
	return ok
}

func (cell *Cell) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Cells[cell]
	return ok
}

func (cellboolean *CellBoolean) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CellBooleans[cellboolean]
	return ok
}

func (cellfloat64 *CellFloat64) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CellFloat64s[cellfloat64]
	return ok
}

func (cellicon *CellIcon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CellIcons[cellicon]
	return ok
}

func (cellint *CellInt) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CellInts[cellint]
	return ok
}

func (cellstring *CellString) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CellStrings[cellstring]
	return ok
}

func (displayedcolumn *DisplayedColumn) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DisplayedColumns[displayedcolumn]
	return ok
}

func (row *Row) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Rows[row]
	return ok
}

func (svgicon *SVGIcon) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SVGIcons[svgicon]
	return ok
}

func (table *Table) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tables[table]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (button *Button) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(cellboolean) {
		return
	}

	cellboolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellfloat64 *CellFloat64) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(cellfloat64) {
		return
	}

	cellfloat64.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellicon *CellIcon) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(cellicon) {
		return
	}

	cellicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellint *CellInt) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(cellint) {
		return
	}

	cellint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellstring *CellString) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(cellstring) {
		return
	}

	cellstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (displayedcolumn *DisplayedColumn) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(displayedcolumn) {
		return
	}

	displayedcolumn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (row *Row) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(svgicon) {
		return
	}

	svgicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (table *Table) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	buttonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, buttonFrom)
	if alreadyCopied {
		return
	}
	buttonFrom.GongCopyBasicFields(buttonTo)

	//insertion point for the staging of instances referenced by pointers
	if buttonFrom.SVGIcon != nil {
		buttonTo.SVGIcon = GongCopyBranchSVGIcon(mapOrigCopy, buttonFrom.SVGIcon)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCell(mapOrigCopy map[any]any, cellFrom *Cell) (cellTo *Cell) {
	var alreadyCopied bool
	cellTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cellFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	cellbooleanTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cellbooleanFrom)
	if alreadyCopied {
		return
	}
	cellbooleanFrom.GongCopyBasicFields(cellbooleanTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellFloat64(mapOrigCopy map[any]any, cellfloat64From *CellFloat64) (cellfloat64To *CellFloat64) {
	var alreadyCopied bool
	cellfloat64To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cellfloat64From)
	if alreadyCopied {
		return
	}
	cellfloat64From.GongCopyBasicFields(cellfloat64To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellIcon(mapOrigCopy map[any]any, celliconFrom *CellIcon) (celliconTo *CellIcon) {
	var alreadyCopied bool
	celliconTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, celliconFrom)
	if alreadyCopied {
		return
	}
	celliconFrom.GongCopyBasicFields(celliconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellInt(mapOrigCopy map[any]any, cellintFrom *CellInt) (cellintTo *CellInt) {
	var alreadyCopied bool
	cellintTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cellintFrom)
	if alreadyCopied {
		return
	}
	cellintFrom.GongCopyBasicFields(cellintTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCellString(mapOrigCopy map[any]any, cellstringFrom *CellString) (cellstringTo *CellString) {
	var alreadyCopied bool
	cellstringTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, cellstringFrom)
	if alreadyCopied {
		return
	}
	cellstringFrom.GongCopyBasicFields(cellstringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDisplayedColumn(mapOrigCopy map[any]any, displayedcolumnFrom *DisplayedColumn) (displayedcolumnTo *DisplayedColumn) {
	var alreadyCopied bool
	displayedcolumnTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, displayedcolumnFrom)
	if alreadyCopied {
		return
	}
	displayedcolumnFrom.GongCopyBasicFields(displayedcolumnTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRow(mapOrigCopy map[any]any, rowFrom *Row) (rowTo *Row) {
	var alreadyCopied bool
	rowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, rowFrom)
	if alreadyCopied {
		return
	}
	rowFrom.GongCopyBasicFields(rowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range rowFrom.Cells {
		rowTo.Cells = append(rowTo.Cells, GongCopyBranchCell(mapOrigCopy, _cell))
	}

	return
}

func GongCopyBranchSVGIcon(mapOrigCopy map[any]any, svgiconFrom *SVGIcon) (svgiconTo *SVGIcon) {
	var alreadyCopied bool
	svgiconTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, svgiconFrom)
	if alreadyCopied {
		return
	}
	svgiconFrom.GongCopyBasicFields(svgiconTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTable(mapOrigCopy map[any]any, tableFrom *Table) (tableTo *Table) {
	var alreadyCopied bool
	tableTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, tableFrom)
	if alreadyCopied {
		return
	}
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

	// check if instance is already staged
	if !stage.IsStaged(cellboolean) {
		return
	}

	cellboolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellfloat64 *CellFloat64) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(cellfloat64) {
		return
	}

	cellfloat64.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellicon *CellIcon) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(cellicon) {
		return
	}

	cellicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellint *CellInt) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(cellint) {
		return
	}

	cellint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (cellstring *CellString) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(cellstring) {
		return
	}

	cellstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (displayedcolumn *DisplayedColumn) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(displayedcolumn) {
		return
	}

	displayedcolumn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (row *Row) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(svgicon) {
		return
	}

	svgicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (table *Table) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructPointer(&reference.SVGIcon, stage.SVGIcons_reference, instance.SVGIcon)
	// insertion point for slice of pointers field
}

func (reference *Cell) GongReconstructPointersFromReferences(stage *Stage, instance *Cell) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.CellString, stage.CellStrings_reference, instance.CellString)
	__gong__reconstructPointer(&reference.CellFloat64, stage.CellFloat64s_reference, instance.CellFloat64)
	__gong__reconstructPointer(&reference.CellInt, stage.CellInts_reference, instance.CellInt)
	__gong__reconstructPointer(&reference.CellBool, stage.CellBooleans_reference, instance.CellBool)
	__gong__reconstructPointer(&reference.CellIcon, stage.CellIcons_reference, instance.CellIcon)
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
	__gong__reconstructSliceOfPointersFromReferences(&reference.Cells, stage.Cells_reference, instance.Cells)
}

func (reference *SVGIcon) GongReconstructPointersFromReferences(stage *Stage, instance *SVGIcon) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Table) GongReconstructPointersFromReferences(stage *Stage, instance *Table) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.DisplayedColumns, stage.DisplayedColumns_reference, instance.DisplayedColumns)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Rows, stage.Rows_reference, instance.Rows)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RowsSelectedForBulkDelete, stage.Rows_reference, instance.RowsSelectedForBulkDelete)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Buttons, stage.Buttons_reference, instance.Buttons)
}

// insertion point for pointer reconstruction from instances
func (reference *Button) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.SVGIcon, stage.SVGIcons_instance)
	// insertion point for slice of pointers fields
}

func (reference *Cell) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.CellString, stage.CellStrings_instance)
	__gong__reconstructPointerFromInstance(&reference.CellFloat64, stage.CellFloat64s_instance)
	__gong__reconstructPointerFromInstance(&reference.CellInt, stage.CellInts_instance)
	__gong__reconstructPointerFromInstance(&reference.CellBool, stage.CellBooleans_instance)
	__gong__reconstructPointerFromInstance(&reference.CellIcon, stage.CellIcons_instance)
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.Cells, stage.Cells_instance)
}

func (reference *SVGIcon) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Table) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.DisplayedColumns, stage.DisplayedColumns_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Rows, stage.Rows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RowsSelectedForBulkDelete, stage.Rows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Buttons, stage.Buttons_instance)
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
	if button.SVGIcon != buttonOther.SVGIcon {
		diffs = append(diffs, button.GongMarshallField(stage, "SVGIcon"))
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
	if cell.CellString != cellOther.CellString {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellString"))
	}
	if cell.CellFloat64 != cellOther.CellFloat64 {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellFloat64"))
	}
	if cell.CellInt != cellOther.CellInt {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellInt"))
	}
	if cell.CellBool != cellOther.CellBool {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellBool"))
	}
	if cell.CellIcon != cellOther.CellIcon {
		diffs = append(diffs, cell.GongMarshallField(stage, "CellIcon"))
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
	if ops := __gong__diffSliceOfPointers(stage, row, "Cells", rowOther.Cells, row.Cells); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, table, "DisplayedColumns", tableOther.DisplayedColumns, table.DisplayedColumns); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, table, "Rows", tableOther.Rows, table.Rows); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, table, "RowsSelectedForBulkDelete", tableOther.RowsSelectedForBulkDelete, table.RowsSelectedForBulkDelete); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, table, "Buttons", tableOther.Buttons, table.Buttons); ops != "" {
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
