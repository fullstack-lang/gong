// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Cell
	for cell := range stage.Cells {
		_ = cell
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, cell.CellString) {
			cell.CellString = nil
		}
		if !IsStaged(stage, cell.CellFloat64) {
			cell.CellFloat64 = nil
		}
		if !IsStaged(stage, cell.CellInt) {
			cell.CellInt = nil
		}
		if !IsStaged(stage, cell.CellBool) {
			cell.CellBool = nil
		}
		if !IsStaged(stage, cell.CellIcon) {
			cell.CellIcon = nil
		}
	}

	// Compute reverse map for named struct CellBoolean
	for cellboolean := range stage.CellBooleans {
		_ = cellboolean
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct CellFloat64
	for cellfloat64 := range stage.CellFloat64s {
		_ = cellfloat64
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct CellIcon
	for cellicon := range stage.CellIcons {
		_ = cellicon
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct CellInt
	for cellint := range stage.CellInts {
		_ = cellint
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct CellString
	for cellstring := range stage.CellStrings {
		_ = cellstring
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct CheckBox
	for checkbox := range stage.CheckBoxs {
		_ = checkbox
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct DisplayedColumn
	for displayedcolumn := range stage.DisplayedColumns {
		_ = displayedcolumn
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FormDiv
	for formdiv := range stage.FormDivs {
		_ = formdiv
		// insertion point per field
		var _FormFields []*FormField
		for _, _formfield := range formdiv.FormFields {
			if IsStaged(stage, _formfield) {
			 	_FormFields = append(_FormFields, _formfield)
			}
		}
		formdiv.FormFields = _FormFields
		var _CheckBoxs []*CheckBox
		for _, _checkbox := range formdiv.CheckBoxs {
			if IsStaged(stage, _checkbox) {
			 	_CheckBoxs = append(_CheckBoxs, _checkbox)
			}
		}
		formdiv.CheckBoxs = _CheckBoxs
		// insertion point per field
		if !IsStaged(stage, formdiv.FormEditAssocButton) {
			formdiv.FormEditAssocButton = nil
		}
		if !IsStaged(stage, formdiv.FormSortAssocButton) {
			formdiv.FormSortAssocButton = nil
		}
	}

	// Compute reverse map for named struct FormEditAssocButton
	for formeditassocbutton := range stage.FormEditAssocButtons {
		_ = formeditassocbutton
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FormField
	for formfield := range stage.FormFields {
		_ = formfield
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, formfield.FormFieldString) {
			formfield.FormFieldString = nil
		}
		if !IsStaged(stage, formfield.FormFieldFloat64) {
			formfield.FormFieldFloat64 = nil
		}
		if !IsStaged(stage, formfield.FormFieldInt) {
			formfield.FormFieldInt = nil
		}
		if !IsStaged(stage, formfield.FormFieldDate) {
			formfield.FormFieldDate = nil
		}
		if !IsStaged(stage, formfield.FormFieldTime) {
			formfield.FormFieldTime = nil
		}
		if !IsStaged(stage, formfield.FormFieldDateTime) {
			formfield.FormFieldDateTime = nil
		}
		if !IsStaged(stage, formfield.FormFieldSelect) {
			formfield.FormFieldSelect = nil
		}
	}

	// Compute reverse map for named struct FormFieldDate
	for formfielddate := range stage.FormFieldDates {
		_ = formfielddate
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FormFieldDateTime
	for formfielddatetime := range stage.FormFieldDateTimes {
		_ = formfielddatetime
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FormFieldFloat64
	for formfieldfloat64 := range stage.FormFieldFloat64s {
		_ = formfieldfloat64
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FormFieldInt
	for formfieldint := range stage.FormFieldInts {
		_ = formfieldint
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FormFieldSelect
	for formfieldselect := range stage.FormFieldSelects {
		_ = formfieldselect
		// insertion point per field
		var _Options []*Option
		for _, _option := range formfieldselect.Options {
			if IsStaged(stage, _option) {
			 	_Options = append(_Options, _option)
			}
		}
		formfieldselect.Options = _Options
		// insertion point per field
		if !IsStaged(stage, formfieldselect.Value) {
			formfieldselect.Value = nil
		}
	}

	// Compute reverse map for named struct FormFieldString
	for formfieldstring := range stage.FormFieldStrings {
		_ = formfieldstring
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FormFieldTime
	for formfieldtime := range stage.FormFieldTimes {
		_ = formfieldtime
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct FormGroup
	for formgroup := range stage.FormGroups {
		_ = formgroup
		// insertion point per field
		var _FormDivs []*FormDiv
		for _, _formdiv := range formgroup.FormDivs {
			if IsStaged(stage, _formdiv) {
			 	_FormDivs = append(_FormDivs, _formdiv)
			}
		}
		formgroup.FormDivs = _FormDivs
		// insertion point per field
	}

	// Compute reverse map for named struct FormSortAssocButton
	for formsortassocbutton := range stage.FormSortAssocButtons {
		_ = formsortassocbutton
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, formsortassocbutton.FormEditAssocButton) {
			formsortassocbutton.FormEditAssocButton = nil
		}
	}

	// Compute reverse map for named struct Option
	for option := range stage.Options {
		_ = option
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Row
	for row := range stage.Rows {
		_ = row
		// insertion point per field
		var _Cells []*Cell
		for _, _cell := range row.Cells {
			if IsStaged(stage, _cell) {
			 	_Cells = append(_Cells, _cell)
			}
		}
		row.Cells = _Cells
		// insertion point per field
	}

	// Compute reverse map for named struct Table
	for table := range stage.Tables {
		_ = table
		// insertion point per field
		var _DisplayedColumns []*DisplayedColumn
		for _, _displayedcolumn := range table.DisplayedColumns {
			if IsStaged(stage, _displayedcolumn) {
			 	_DisplayedColumns = append(_DisplayedColumns, _displayedcolumn)
			}
		}
		table.DisplayedColumns = _DisplayedColumns
		var _Rows []*Row
		for _, _row := range table.Rows {
			if IsStaged(stage, _row) {
			 	_Rows = append(_Rows, _row)
			}
		}
		table.Rows = _Rows
		// insertion point per field
	}

}
