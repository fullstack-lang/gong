// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Cell
	// insertion point per field

	// Compute reverse map for named struct CellBoolean
	// insertion point per field

	// Compute reverse map for named struct CellFloat64
	// insertion point per field

	// Compute reverse map for named struct CellIcon
	// insertion point per field

	// Compute reverse map for named struct CellInt
	// insertion point per field

	// Compute reverse map for named struct CellString
	// insertion point per field

	// Compute reverse map for named struct CheckBox
	// insertion point per field

	// Compute reverse map for named struct DisplayedColumn
	// insertion point per field

	// Compute reverse map for named struct FormDiv
	// insertion point per field
	clear(stage.FormDiv_FormFields_reverseMap)
	stage.FormDiv_FormFields_reverseMap = make(map[*FormField]*FormDiv)
	for formdiv := range stage.FormDivs {
		_ = formdiv
		for _, _formfield := range formdiv.FormFields {
			stage.FormDiv_FormFields_reverseMap[_formfield] = formdiv
		}
	}
	clear(stage.FormDiv_CheckBoxs_reverseMap)
	stage.FormDiv_CheckBoxs_reverseMap = make(map[*CheckBox]*FormDiv)
	for formdiv := range stage.FormDivs {
		_ = formdiv
		for _, _checkbox := range formdiv.CheckBoxs {
			stage.FormDiv_CheckBoxs_reverseMap[_checkbox] = formdiv
		}
	}

	// Compute reverse map for named struct FormEditAssocButton
	// insertion point per field

	// Compute reverse map for named struct FormField
	// insertion point per field

	// Compute reverse map for named struct FormFieldDate
	// insertion point per field

	// Compute reverse map for named struct FormFieldDateTime
	// insertion point per field

	// Compute reverse map for named struct FormFieldFloat64
	// insertion point per field

	// Compute reverse map for named struct FormFieldInt
	// insertion point per field

	// Compute reverse map for named struct FormFieldSelect
	// insertion point per field
	clear(stage.FormFieldSelect_Options_reverseMap)
	stage.FormFieldSelect_Options_reverseMap = make(map[*Option]*FormFieldSelect)
	for formfieldselect := range stage.FormFieldSelects {
		_ = formfieldselect
		for _, _option := range formfieldselect.Options {
			stage.FormFieldSelect_Options_reverseMap[_option] = formfieldselect
		}
	}

	// Compute reverse map for named struct FormFieldString
	// insertion point per field

	// Compute reverse map for named struct FormFieldTime
	// insertion point per field

	// Compute reverse map for named struct FormGroup
	// insertion point per field
	clear(stage.FormGroup_FormDivs_reverseMap)
	stage.FormGroup_FormDivs_reverseMap = make(map[*FormDiv]*FormGroup)
	for formgroup := range stage.FormGroups {
		_ = formgroup
		for _, _formdiv := range formgroup.FormDivs {
			stage.FormGroup_FormDivs_reverseMap[_formdiv] = formgroup
		}
	}

	// Compute reverse map for named struct FormSortAssocButton
	// insertion point per field

	// Compute reverse map for named struct Option
	// insertion point per field

	// Compute reverse map for named struct Row
	// insertion point per field
	clear(stage.Row_Cells_reverseMap)
	stage.Row_Cells_reverseMap = make(map[*Cell]*Row)
	for row := range stage.Rows {
		_ = row
		for _, _cell := range row.Cells {
			stage.Row_Cells_reverseMap[_cell] = row
		}
	}

	// Compute reverse map for named struct Table
	// insertion point per field
	clear(stage.Table_DisplayedColumns_reverseMap)
	stage.Table_DisplayedColumns_reverseMap = make(map[*DisplayedColumn]*Table)
	for table := range stage.Tables {
		_ = table
		for _, _displayedcolumn := range table.DisplayedColumns {
			stage.Table_DisplayedColumns_reverseMap[_displayedcolumn] = table
		}
	}
	clear(stage.Table_Rows_reverseMap)
	stage.Table_Rows_reverseMap = make(map[*Row]*Table)
	for table := range stage.Tables {
		_ = table
		for _, _row := range table.Rows {
			stage.Table_Rows_reverseMap[_row] = table
		}
	}

}
