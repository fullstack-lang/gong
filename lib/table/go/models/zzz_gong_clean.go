// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	*slice = cleanedSlice
	return len(cleanedSlice) != len(*slice)
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	if !IsStagedPointerToGongstruct(stage, *element) {
		var zero T
		*element = zero
		return true
	}
	return false
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Cell
func (cell *Cell) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &cell.CellString)  || modified
	modified = GongCleanPointer(stage, &cell.CellFloat64)  || modified
	modified = GongCleanPointer(stage, &cell.CellInt)  || modified
	modified = GongCleanPointer(stage, &cell.CellBool)  || modified
	modified = GongCleanPointer(stage, &cell.CellIcon)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CellBoolean
func (cellboolean *CellBoolean) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CellFloat64
func (cellfloat64 *CellFloat64) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CellIcon
func (cellicon *CellIcon) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CellInt
func (cellint *CellInt) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CellString
func (cellstring *CellString) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CheckBox
func (checkbox *CheckBox) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DisplayedColumn
func (displayedcolumn *DisplayedColumn) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormDiv
func (formdiv *FormDiv) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &formdiv.FormFields)  || modified
	modified = GongCleanSlice(stage, &formdiv.CheckBoxs)  || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &formdiv.FormEditAssocButton)  || modified
	modified = GongCleanPointer(stage, &formdiv.FormSortAssocButton)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FormEditAssocButton
func (formeditassocbutton *FormEditAssocButton) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormField
func (formfield *FormField) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &formfield.FormFieldString)  || modified
	modified = GongCleanPointer(stage, &formfield.FormFieldFloat64)  || modified
	modified = GongCleanPointer(stage, &formfield.FormFieldInt)  || modified
	modified = GongCleanPointer(stage, &formfield.FormFieldDate)  || modified
	modified = GongCleanPointer(stage, &formfield.FormFieldTime)  || modified
	modified = GongCleanPointer(stage, &formfield.FormFieldDateTime)  || modified
	modified = GongCleanPointer(stage, &formfield.FormFieldSelect)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FormFieldDate
func (formfielddate *FormFieldDate) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormFieldDateTime
func (formfielddatetime *FormFieldDateTime) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormFieldFloat64
func (formfieldfloat64 *FormFieldFloat64) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormFieldInt
func (formfieldint *FormFieldInt) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormFieldSelect
func (formfieldselect *FormFieldSelect) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &formfieldselect.Options)  || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &formfieldselect.Value)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FormFieldString
func (formfieldstring *FormFieldString) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormFieldTime
func (formfieldtime *FormFieldTime) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormGroup
func (formgroup *FormGroup) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &formgroup.FormDivs)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FormSortAssocButton
func (formsortassocbutton *FormSortAssocButton) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &formsortassocbutton.FormEditAssocButton)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Option
func (option *Option) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Row
func (row *Row) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &row.Cells)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Table
func (table *Table) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &table.DisplayedColumns)  || modified
	modified = GongCleanSlice(stage, &table.Rows)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
