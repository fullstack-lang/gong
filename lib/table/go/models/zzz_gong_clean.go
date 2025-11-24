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

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Cell
func (cell *Cell) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	cell.CellString = GongCleanPointer(stage, cell.CellString)
	cell.CellFloat64 = GongCleanPointer(stage, cell.CellFloat64)
	cell.CellInt = GongCleanPointer(stage, cell.CellInt)
	cell.CellBool = GongCleanPointer(stage, cell.CellBool)
	cell.CellIcon = GongCleanPointer(stage, cell.CellIcon)
}

// Clean garbage collect unstaged instances that are referenced by CellBoolean
func (cellboolean *CellBoolean) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by CellFloat64
func (cellfloat64 *CellFloat64) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by CellIcon
func (cellicon *CellIcon) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by CellInt
func (cellint *CellInt) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by CellString
func (cellstring *CellString) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by CheckBox
func (checkbox *CheckBox) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by DisplayedColumn
func (displayedcolumn *DisplayedColumn) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormDiv
func (formdiv *FormDiv) GongClean(stage *Stage) {
	// insertion point per field
	formdiv.FormFields = GongCleanSlice(stage, formdiv.FormFields)
	formdiv.CheckBoxs = GongCleanSlice(stage, formdiv.CheckBoxs)
	// insertion point per field
	formdiv.FormEditAssocButton = GongCleanPointer(stage, formdiv.FormEditAssocButton)
	formdiv.FormSortAssocButton = GongCleanPointer(stage, formdiv.FormSortAssocButton)
}

// Clean garbage collect unstaged instances that are referenced by FormEditAssocButton
func (formeditassocbutton *FormEditAssocButton) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormField
func (formfield *FormField) GongClean(stage *Stage) {
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

// Clean garbage collect unstaged instances that are referenced by FormFieldDate
func (formfielddate *FormFieldDate) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormFieldDateTime
func (formfielddatetime *FormFieldDateTime) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormFieldFloat64
func (formfieldfloat64 *FormFieldFloat64) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormFieldInt
func (formfieldint *FormFieldInt) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormFieldSelect
func (formfieldselect *FormFieldSelect) GongClean(stage *Stage) {
	// insertion point per field
	formfieldselect.Options = GongCleanSlice(stage, formfieldselect.Options)
	// insertion point per field
	formfieldselect.Value = GongCleanPointer(stage, formfieldselect.Value)
}

// Clean garbage collect unstaged instances that are referenced by FormFieldString
func (formfieldstring *FormFieldString) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormFieldTime
func (formfieldtime *FormFieldTime) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormGroup
func (formgroup *FormGroup) GongClean(stage *Stage) {
	// insertion point per field
	formgroup.FormDivs = GongCleanSlice(stage, formgroup.FormDivs)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by FormSortAssocButton
func (formsortassocbutton *FormSortAssocButton) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	formsortassocbutton.FormEditAssocButton = GongCleanPointer(stage, formsortassocbutton.FormEditAssocButton)
}

// Clean garbage collect unstaged instances that are referenced by Option
func (option *Option) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Row
func (row *Row) GongClean(stage *Stage) {
	// insertion point per field
	row.Cells = GongCleanSlice(stage, row.Cells)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Table
func (table *Table) GongClean(stage *Stage) {
	// insertion point per field
	table.DisplayedColumns = GongCleanSlice(stage, table.DisplayedColumns)
	table.Rows = GongCleanSlice(stage, table.Rows)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
