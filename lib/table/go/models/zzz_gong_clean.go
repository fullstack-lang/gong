// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}
    
	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// Clean computes the reverse map, for all intances, for all clean to pointers field
func (stage *Stage) Clean() {
	// insertion point per named struct
	// clean up Cell
	for cell := range stage.Cells {
		_ = cell
		// insertion point per field
		// insertion point per field
		cell.CellString = GongCleanPointer(stage, cell.CellString)
		cell.CellFloat64 = GongCleanPointer(stage, cell.CellFloat64)
		cell.CellInt = GongCleanPointer(stage, cell.CellInt)
		cell.CellBool = GongCleanPointer(stage, cell.CellBool)
		cell.CellIcon = GongCleanPointer(stage, cell.CellIcon)
	}

	// clean up CellBoolean
	for cellboolean := range stage.CellBooleans {
		_ = cellboolean
		// insertion point per field
		// insertion point per field
	}

	// clean up CellFloat64
	for cellfloat64 := range stage.CellFloat64s {
		_ = cellfloat64
		// insertion point per field
		// insertion point per field
	}

	// clean up CellIcon
	for cellicon := range stage.CellIcons {
		_ = cellicon
		// insertion point per field
		// insertion point per field
	}

	// clean up CellInt
	for cellint := range stage.CellInts {
		_ = cellint
		// insertion point per field
		// insertion point per field
	}

	// clean up CellString
	for cellstring := range stage.CellStrings {
		_ = cellstring
		// insertion point per field
		// insertion point per field
	}

	// clean up CheckBox
	for checkbox := range stage.CheckBoxs {
		_ = checkbox
		// insertion point per field
		// insertion point per field
	}

	// clean up DisplayedColumn
	for displayedcolumn := range stage.DisplayedColumns {
		_ = displayedcolumn
		// insertion point per field
		// insertion point per field
	}

	// clean up FormDiv
	for formdiv := range stage.FormDivs {
		_ = formdiv
		// insertion point per field
		formdiv.FormFields = GongCleanSlice(stage, formdiv.FormFields)
		formdiv.CheckBoxs = GongCleanSlice(stage, formdiv.CheckBoxs)
		// insertion point per field
		formdiv.FormEditAssocButton = GongCleanPointer(stage, formdiv.FormEditAssocButton)
		formdiv.FormSortAssocButton = GongCleanPointer(stage, formdiv.FormSortAssocButton)
	}

	// clean up FormEditAssocButton
	for formeditassocbutton := range stage.FormEditAssocButtons {
		_ = formeditassocbutton
		// insertion point per field
		// insertion point per field
	}

	// clean up FormField
	for formfield := range stage.FormFields {
		_ = formfield
		// insertion point per field
		// insertion point per field
		formfield.FormFieldString = GongCleanPointer(stage, formfield.FormFieldString)
		formfield.FormFieldFloat64 = GongCleanPointer(stage, formfield.FormFieldFloat64)
		formfield.FormFieldInt = GongCleanPointer(stage, formfield.FormFieldInt)
		formfield.FormFieldDate = GongCleanPointer(stage, formfield.FormFieldDate)
		formfield.FormFieldTime = GongCleanPointer(stage, formfield.FormFieldTime)
		formfield.FormFieldDateTime = GongCleanPointer(stage, formfield.FormFieldDateTime)
		formfield.FormFieldSelect = GongCleanPointer(stage, formfield.FormFieldSelect)
	}

	// clean up FormFieldDate
	for formfielddate := range stage.FormFieldDates {
		_ = formfielddate
		// insertion point per field
		// insertion point per field
	}

	// clean up FormFieldDateTime
	for formfielddatetime := range stage.FormFieldDateTimes {
		_ = formfielddatetime
		// insertion point per field
		// insertion point per field
	}

	// clean up FormFieldFloat64
	for formfieldfloat64 := range stage.FormFieldFloat64s {
		_ = formfieldfloat64
		// insertion point per field
		// insertion point per field
	}

	// clean up FormFieldInt
	for formfieldint := range stage.FormFieldInts {
		_ = formfieldint
		// insertion point per field
		// insertion point per field
	}

	// clean up FormFieldSelect
	for formfieldselect := range stage.FormFieldSelects {
		_ = formfieldselect
		// insertion point per field
		formfieldselect.Options = GongCleanSlice(stage, formfieldselect.Options)
		// insertion point per field
		formfieldselect.Value = GongCleanPointer(stage, formfieldselect.Value)
	}

	// clean up FormFieldString
	for formfieldstring := range stage.FormFieldStrings {
		_ = formfieldstring
		// insertion point per field
		// insertion point per field
	}

	// clean up FormFieldTime
	for formfieldtime := range stage.FormFieldTimes {
		_ = formfieldtime
		// insertion point per field
		// insertion point per field
	}

	// clean up FormGroup
	for formgroup := range stage.FormGroups {
		_ = formgroup
		// insertion point per field
		formgroup.FormDivs = GongCleanSlice(stage, formgroup.FormDivs)
		// insertion point per field
	}

	// clean up FormSortAssocButton
	for formsortassocbutton := range stage.FormSortAssocButtons {
		_ = formsortassocbutton
		// insertion point per field
		// insertion point per field
		formsortassocbutton.FormEditAssocButton = GongCleanPointer(stage, formsortassocbutton.FormEditAssocButton)
	}

	// clean up Option
	for option := range stage.Options {
		_ = option
		// insertion point per field
		// insertion point per field
	}

	// clean up Row
	for row := range stage.Rows {
		_ = row
		// insertion point per field
		row.Cells = GongCleanSlice(stage, row.Cells)
		// insertion point per field
	}

	// clean up Table
	for table := range stage.Tables {
		_ = table
		// insertion point per field
		table.DisplayedColumns = GongCleanSlice(stage, table.DisplayedColumns)
		table.Rows = GongCleanSlice(stage, table.Rows)
		// insertion point per field
	}

}
