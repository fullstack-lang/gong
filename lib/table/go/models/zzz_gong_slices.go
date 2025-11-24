// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
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

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Cells {
		res = append(res, instance)
	}

	for instance := range stage.CellBooleans {
		res = append(res, instance)
	}

	for instance := range stage.CellFloat64s {
		res = append(res, instance)
	}

	for instance := range stage.CellIcons {
		res = append(res, instance)
	}

	for instance := range stage.CellInts {
		res = append(res, instance)
	}

	for instance := range stage.CellStrings {
		res = append(res, instance)
	}

	for instance := range stage.CheckBoxs {
		res = append(res, instance)
	}

	for instance := range stage.DisplayedColumns {
		res = append(res, instance)
	}

	for instance := range stage.FormDivs {
		res = append(res, instance)
	}

	for instance := range stage.FormEditAssocButtons {
		res = append(res, instance)
	}

	for instance := range stage.FormFields {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldDates {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldDateTimes {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldFloat64s {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldInts {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldSelects {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldStrings {
		res = append(res, instance)
	}

	for instance := range stage.FormFieldTimes {
		res = append(res, instance)
	}

	for instance := range stage.FormGroups {
		res = append(res, instance)
	}

	for instance := range stage.FormSortAssocButtons {
		res = append(res, instance)
	}

	for instance := range stage.Options {
		res = append(res, instance)
	}

	for instance := range stage.Rows {
		res = append(res, instance)
	}

	for instance := range stage.Tables {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (cell *Cell) GongCopy() GongstructIF {
	var newInstance Cell
	newInstance = *cell
	return &newInstance
}

func (cellboolean *CellBoolean) GongCopy() GongstructIF {
	var newInstance CellBoolean
	newInstance = *cellboolean
	return &newInstance
}

func (cellfloat64 *CellFloat64) GongCopy() GongstructIF {
	var newInstance CellFloat64
	newInstance = *cellfloat64
	return &newInstance
}

func (cellicon *CellIcon) GongCopy() GongstructIF {
	var newInstance CellIcon
	newInstance = *cellicon
	return &newInstance
}

func (cellint *CellInt) GongCopy() GongstructIF {
	var newInstance CellInt
	newInstance = *cellint
	return &newInstance
}

func (cellstring *CellString) GongCopy() GongstructIF {
	var newInstance CellString
	newInstance = *cellstring
	return &newInstance
}

func (checkbox *CheckBox) GongCopy() GongstructIF {
	var newInstance CheckBox
	newInstance = *checkbox
	return &newInstance
}

func (displayedcolumn *DisplayedColumn) GongCopy() GongstructIF {
	var newInstance DisplayedColumn
	newInstance = *displayedcolumn
	return &newInstance
}

func (formdiv *FormDiv) GongCopy() GongstructIF {
	var newInstance FormDiv
	newInstance = *formdiv
	return &newInstance
}

func (formeditassocbutton *FormEditAssocButton) GongCopy() GongstructIF {
	var newInstance FormEditAssocButton
	newInstance = *formeditassocbutton
	return &newInstance
}

func (formfield *FormField) GongCopy() GongstructIF {
	var newInstance FormField
	newInstance = *formfield
	return &newInstance
}

func (formfielddate *FormFieldDate) GongCopy() GongstructIF {
	var newInstance FormFieldDate
	newInstance = *formfielddate
	return &newInstance
}

func (formfielddatetime *FormFieldDateTime) GongCopy() GongstructIF {
	var newInstance FormFieldDateTime
	newInstance = *formfielddatetime
	return &newInstance
}

func (formfieldfloat64 *FormFieldFloat64) GongCopy() GongstructIF {
	var newInstance FormFieldFloat64
	newInstance = *formfieldfloat64
	return &newInstance
}

func (formfieldint *FormFieldInt) GongCopy() GongstructIF {
	var newInstance FormFieldInt
	newInstance = *formfieldint
	return &newInstance
}

func (formfieldselect *FormFieldSelect) GongCopy() GongstructIF {
	var newInstance FormFieldSelect
	newInstance = *formfieldselect
	return &newInstance
}

func (formfieldstring *FormFieldString) GongCopy() GongstructIF {
	var newInstance FormFieldString
	newInstance = *formfieldstring
	return &newInstance
}

func (formfieldtime *FormFieldTime) GongCopy() GongstructIF {
	var newInstance FormFieldTime
	newInstance = *formfieldtime
	return &newInstance
}

func (formgroup *FormGroup) GongCopy() GongstructIF {
	var newInstance FormGroup
	newInstance = *formgroup
	return &newInstance
}

func (formsortassocbutton *FormSortAssocButton) GongCopy() GongstructIF {
	var newInstance FormSortAssocButton
	newInstance = *formsortassocbutton
	return &newInstance
}

func (option *Option) GongCopy() GongstructIF {
	var newInstance Option
	newInstance = *option
	return &newInstance
}

func (row *Row) GongCopy() GongstructIF {
	var newInstance Row
	newInstance = *row
	return &newInstance
}

func (table *Table) GongCopy() GongstructIF {
	var newInstance Table
	newInstance = *table
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
