// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
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

	case *CheckBox:
		ok = stage.IsStagedCheckBox(target)

	case *DisplayedColumn:
		ok = stage.IsStagedDisplayedColumn(target)

	case *FormDiv:
		ok = stage.IsStagedFormDiv(target)

	case *FormEditAssocButton:
		ok = stage.IsStagedFormEditAssocButton(target)

	case *FormField:
		ok = stage.IsStagedFormField(target)

	case *FormFieldDate:
		ok = stage.IsStagedFormFieldDate(target)

	case *FormFieldDateTime:
		ok = stage.IsStagedFormFieldDateTime(target)

	case *FormFieldFloat64:
		ok = stage.IsStagedFormFieldFloat64(target)

	case *FormFieldInt:
		ok = stage.IsStagedFormFieldInt(target)

	case *FormFieldSelect:
		ok = stage.IsStagedFormFieldSelect(target)

	case *FormFieldString:
		ok = stage.IsStagedFormFieldString(target)

	case *FormFieldTime:
		ok = stage.IsStagedFormFieldTime(target)

	case *FormGroup:
		ok = stage.IsStagedFormGroup(target)

	case *FormSortAssocButton:
		ok = stage.IsStagedFormSortAssocButton(target)

	case *Option:
		ok = stage.IsStagedOption(target)

	case *Row:
		ok = stage.IsStagedRow(target)

	case *Table:
		ok = stage.IsStagedTable(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
	func (stage *StageStruct) IsStagedCell(cell *Cell) (ok bool) {

		_, ok = stage.Cells[cell]
	
		return
	}

	func (stage *StageStruct) IsStagedCellBoolean(cellboolean *CellBoolean) (ok bool) {

		_, ok = stage.CellBooleans[cellboolean]
	
		return
	}

	func (stage *StageStruct) IsStagedCellFloat64(cellfloat64 *CellFloat64) (ok bool) {

		_, ok = stage.CellFloat64s[cellfloat64]
	
		return
	}

	func (stage *StageStruct) IsStagedCellIcon(cellicon *CellIcon) (ok bool) {

		_, ok = stage.CellIcons[cellicon]
	
		return
	}

	func (stage *StageStruct) IsStagedCellInt(cellint *CellInt) (ok bool) {

		_, ok = stage.CellInts[cellint]
	
		return
	}

	func (stage *StageStruct) IsStagedCellString(cellstring *CellString) (ok bool) {

		_, ok = stage.CellStrings[cellstring]
	
		return
	}

	func (stage *StageStruct) IsStagedCheckBox(checkbox *CheckBox) (ok bool) {

		_, ok = stage.CheckBoxs[checkbox]
	
		return
	}

	func (stage *StageStruct) IsStagedDisplayedColumn(displayedcolumn *DisplayedColumn) (ok bool) {

		_, ok = stage.DisplayedColumns[displayedcolumn]
	
		return
	}

	func (stage *StageStruct) IsStagedFormDiv(formdiv *FormDiv) (ok bool) {

		_, ok = stage.FormDivs[formdiv]
	
		return
	}

	func (stage *StageStruct) IsStagedFormEditAssocButton(formeditassocbutton *FormEditAssocButton) (ok bool) {

		_, ok = stage.FormEditAssocButtons[formeditassocbutton]
	
		return
	}

	func (stage *StageStruct) IsStagedFormField(formfield *FormField) (ok bool) {

		_, ok = stage.FormFields[formfield]
	
		return
	}

	func (stage *StageStruct) IsStagedFormFieldDate(formfielddate *FormFieldDate) (ok bool) {

		_, ok = stage.FormFieldDates[formfielddate]
	
		return
	}

	func (stage *StageStruct) IsStagedFormFieldDateTime(formfielddatetime *FormFieldDateTime) (ok bool) {

		_, ok = stage.FormFieldDateTimes[formfielddatetime]
	
		return
	}

	func (stage *StageStruct) IsStagedFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) (ok bool) {

		_, ok = stage.FormFieldFloat64s[formfieldfloat64]
	
		return
	}

	func (stage *StageStruct) IsStagedFormFieldInt(formfieldint *FormFieldInt) (ok bool) {

		_, ok = stage.FormFieldInts[formfieldint]
	
		return
	}

	func (stage *StageStruct) IsStagedFormFieldSelect(formfieldselect *FormFieldSelect) (ok bool) {

		_, ok = stage.FormFieldSelects[formfieldselect]
	
		return
	}

	func (stage *StageStruct) IsStagedFormFieldString(formfieldstring *FormFieldString) (ok bool) {

		_, ok = stage.FormFieldStrings[formfieldstring]
	
		return
	}

	func (stage *StageStruct) IsStagedFormFieldTime(formfieldtime *FormFieldTime) (ok bool) {

		_, ok = stage.FormFieldTimes[formfieldtime]
	
		return
	}

	func (stage *StageStruct) IsStagedFormGroup(formgroup *FormGroup) (ok bool) {

		_, ok = stage.FormGroups[formgroup]
	
		return
	}

	func (stage *StageStruct) IsStagedFormSortAssocButton(formsortassocbutton *FormSortAssocButton) (ok bool) {

		_, ok = stage.FormSortAssocButtons[formsortassocbutton]
	
		return
	}

	func (stage *StageStruct) IsStagedOption(option *Option) (ok bool) {

		_, ok = stage.Options[option]
	
		return
	}

	func (stage *StageStruct) IsStagedRow(row *Row) (ok bool) {

		_, ok = stage.Rows[row]
	
		return
	}

	func (stage *StageStruct) IsStagedTable(table *Table) (ok bool) {

		_, ok = stage.Tables[table]
	
		return
	}


// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
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

	case *CheckBox:
		stage.StageBranchCheckBox(target)

	case *DisplayedColumn:
		stage.StageBranchDisplayedColumn(target)

	case *FormDiv:
		stage.StageBranchFormDiv(target)

	case *FormEditAssocButton:
		stage.StageBranchFormEditAssocButton(target)

	case *FormField:
		stage.StageBranchFormField(target)

	case *FormFieldDate:
		stage.StageBranchFormFieldDate(target)

	case *FormFieldDateTime:
		stage.StageBranchFormFieldDateTime(target)

	case *FormFieldFloat64:
		stage.StageBranchFormFieldFloat64(target)

	case *FormFieldInt:
		stage.StageBranchFormFieldInt(target)

	case *FormFieldSelect:
		stage.StageBranchFormFieldSelect(target)

	case *FormFieldString:
		stage.StageBranchFormFieldString(target)

	case *FormFieldTime:
		stage.StageBranchFormFieldTime(target)

	case *FormGroup:
		stage.StageBranchFormGroup(target)

	case *FormSortAssocButton:
		stage.StageBranchFormSortAssocButton(target)

	case *Option:
		stage.StageBranchOption(target)

	case *Row:
		stage.StageBranchRow(target)

	case *Table:
		stage.StageBranchTable(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchCell(cell *Cell) {

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

func (stage *StageStruct) StageBranchCellBoolean(cellboolean *CellBoolean) {

	// check if instance is already staged
	if IsStaged(stage, cellboolean) {
		return
	}

	cellboolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCellFloat64(cellfloat64 *CellFloat64) {

	// check if instance is already staged
	if IsStaged(stage, cellfloat64) {
		return
	}

	cellfloat64.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCellIcon(cellicon *CellIcon) {

	// check if instance is already staged
	if IsStaged(stage, cellicon) {
		return
	}

	cellicon.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCellInt(cellint *CellInt) {

	// check if instance is already staged
	if IsStaged(stage, cellint) {
		return
	}

	cellint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCellString(cellstring *CellString) {

	// check if instance is already staged
	if IsStaged(stage, cellstring) {
		return
	}

	cellstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCheckBox(checkbox *CheckBox) {

	// check if instance is already staged
	if IsStaged(stage, checkbox) {
		return
	}

	checkbox.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDisplayedColumn(displayedcolumn *DisplayedColumn) {

	// check if instance is already staged
	if IsStaged(stage, displayedcolumn) {
		return
	}

	displayedcolumn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormDiv(formdiv *FormDiv) {

	// check if instance is already staged
	if IsStaged(stage, formdiv) {
		return
	}

	formdiv.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formdiv.FormEditAssocButton != nil {
		StageBranch(stage, formdiv.FormEditAssocButton)
	}
	if formdiv.FormSortAssocButton != nil {
		StageBranch(stage, formdiv.FormSortAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formfield := range formdiv.FormFields {
		StageBranch(stage, _formfield)
	}
	for _, _checkbox := range formdiv.CheckBoxs {
		StageBranch(stage, _checkbox)
	}

}

func (stage *StageStruct) StageBranchFormEditAssocButton(formeditassocbutton *FormEditAssocButton) {

	// check if instance is already staged
	if IsStaged(stage, formeditassocbutton) {
		return
	}

	formeditassocbutton.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormField(formfield *FormField) {

	// check if instance is already staged
	if IsStaged(stage, formfield) {
		return
	}

	formfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formfield.FormFieldString != nil {
		StageBranch(stage, formfield.FormFieldString)
	}
	if formfield.FormFieldFloat64 != nil {
		StageBranch(stage, formfield.FormFieldFloat64)
	}
	if formfield.FormFieldInt != nil {
		StageBranch(stage, formfield.FormFieldInt)
	}
	if formfield.FormFieldDate != nil {
		StageBranch(stage, formfield.FormFieldDate)
	}
	if formfield.FormFieldTime != nil {
		StageBranch(stage, formfield.FormFieldTime)
	}
	if formfield.FormFieldDateTime != nil {
		StageBranch(stage, formfield.FormFieldDateTime)
	}
	if formfield.FormFieldSelect != nil {
		StageBranch(stage, formfield.FormFieldSelect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormFieldDate(formfielddate *FormFieldDate) {

	// check if instance is already staged
	if IsStaged(stage, formfielddate) {
		return
	}

	formfielddate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormFieldDateTime(formfielddatetime *FormFieldDateTime) {

	// check if instance is already staged
	if IsStaged(stage, formfielddatetime) {
		return
	}

	formfielddatetime.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) {

	// check if instance is already staged
	if IsStaged(stage, formfieldfloat64) {
		return
	}

	formfieldfloat64.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormFieldInt(formfieldint *FormFieldInt) {

	// check if instance is already staged
	if IsStaged(stage, formfieldint) {
		return
	}

	formfieldint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormFieldSelect(formfieldselect *FormFieldSelect) {

	// check if instance is already staged
	if IsStaged(stage, formfieldselect) {
		return
	}

	formfieldselect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formfieldselect.Value != nil {
		StageBranch(stage, formfieldselect.Value)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _option := range formfieldselect.Options {
		StageBranch(stage, _option)
	}

}

func (stage *StageStruct) StageBranchFormFieldString(formfieldstring *FormFieldString) {

	// check if instance is already staged
	if IsStaged(stage, formfieldstring) {
		return
	}

	formfieldstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormFieldTime(formfieldtime *FormFieldTime) {

	// check if instance is already staged
	if IsStaged(stage, formfieldtime) {
		return
	}

	formfieldtime.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormGroup(formgroup *FormGroup) {

	// check if instance is already staged
	if IsStaged(stage, formgroup) {
		return
	}

	formgroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formdiv := range formgroup.FormDivs {
		StageBranch(stage, _formdiv)
	}

}

func (stage *StageStruct) StageBranchFormSortAssocButton(formsortassocbutton *FormSortAssocButton) {

	// check if instance is already staged
	if IsStaged(stage, formsortassocbutton) {
		return
	}

	formsortassocbutton.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOption(option *Option) {

	// check if instance is already staged
	if IsStaged(stage, option) {
		return
	}

	option.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRow(row *Row) {

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

func (stage *StageStruct) StageBranchTable(table *Table) {

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

}


// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
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

	case *CheckBox:
		stage.UnstageBranchCheckBox(target)

	case *DisplayedColumn:
		stage.UnstageBranchDisplayedColumn(target)

	case *FormDiv:
		stage.UnstageBranchFormDiv(target)

	case *FormEditAssocButton:
		stage.UnstageBranchFormEditAssocButton(target)

	case *FormField:
		stage.UnstageBranchFormField(target)

	case *FormFieldDate:
		stage.UnstageBranchFormFieldDate(target)

	case *FormFieldDateTime:
		stage.UnstageBranchFormFieldDateTime(target)

	case *FormFieldFloat64:
		stage.UnstageBranchFormFieldFloat64(target)

	case *FormFieldInt:
		stage.UnstageBranchFormFieldInt(target)

	case *FormFieldSelect:
		stage.UnstageBranchFormFieldSelect(target)

	case *FormFieldString:
		stage.UnstageBranchFormFieldString(target)

	case *FormFieldTime:
		stage.UnstageBranchFormFieldTime(target)

	case *FormGroup:
		stage.UnstageBranchFormGroup(target)

	case *FormSortAssocButton:
		stage.UnstageBranchFormSortAssocButton(target)

	case *Option:
		stage.UnstageBranchOption(target)

	case *Row:
		stage.UnstageBranchRow(target)

	case *Table:
		stage.UnstageBranchTable(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchCell(cell *Cell) {

	// check if instance is already staged
	if ! IsStaged(stage, cell) {
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

func (stage *StageStruct) UnstageBranchCellBoolean(cellboolean *CellBoolean) {

	// check if instance is already staged
	if ! IsStaged(stage, cellboolean) {
		return
	}

	cellboolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCellFloat64(cellfloat64 *CellFloat64) {

	// check if instance is already staged
	if ! IsStaged(stage, cellfloat64) {
		return
	}

	cellfloat64.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCellIcon(cellicon *CellIcon) {

	// check if instance is already staged
	if ! IsStaged(stage, cellicon) {
		return
	}

	cellicon.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCellInt(cellint *CellInt) {

	// check if instance is already staged
	if ! IsStaged(stage, cellint) {
		return
	}

	cellint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCellString(cellstring *CellString) {

	// check if instance is already staged
	if ! IsStaged(stage, cellstring) {
		return
	}

	cellstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCheckBox(checkbox *CheckBox) {

	// check if instance is already staged
	if ! IsStaged(stage, checkbox) {
		return
	}

	checkbox.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDisplayedColumn(displayedcolumn *DisplayedColumn) {

	// check if instance is already staged
	if ! IsStaged(stage, displayedcolumn) {
		return
	}

	displayedcolumn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormDiv(formdiv *FormDiv) {

	// check if instance is already staged
	if ! IsStaged(stage, formdiv) {
		return
	}

	formdiv.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formdiv.FormEditAssocButton != nil {
		UnstageBranch(stage, formdiv.FormEditAssocButton)
	}
	if formdiv.FormSortAssocButton != nil {
		UnstageBranch(stage, formdiv.FormSortAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formfield := range formdiv.FormFields {
		UnstageBranch(stage, _formfield)
	}
	for _, _checkbox := range formdiv.CheckBoxs {
		UnstageBranch(stage, _checkbox)
	}

}

func (stage *StageStruct) UnstageBranchFormEditAssocButton(formeditassocbutton *FormEditAssocButton) {

	// check if instance is already staged
	if ! IsStaged(stage, formeditassocbutton) {
		return
	}

	formeditassocbutton.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormField(formfield *FormField) {

	// check if instance is already staged
	if ! IsStaged(stage, formfield) {
		return
	}

	formfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formfield.FormFieldString != nil {
		UnstageBranch(stage, formfield.FormFieldString)
	}
	if formfield.FormFieldFloat64 != nil {
		UnstageBranch(stage, formfield.FormFieldFloat64)
	}
	if formfield.FormFieldInt != nil {
		UnstageBranch(stage, formfield.FormFieldInt)
	}
	if formfield.FormFieldDate != nil {
		UnstageBranch(stage, formfield.FormFieldDate)
	}
	if formfield.FormFieldTime != nil {
		UnstageBranch(stage, formfield.FormFieldTime)
	}
	if formfield.FormFieldDateTime != nil {
		UnstageBranch(stage, formfield.FormFieldDateTime)
	}
	if formfield.FormFieldSelect != nil {
		UnstageBranch(stage, formfield.FormFieldSelect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormFieldDate(formfielddate *FormFieldDate) {

	// check if instance is already staged
	if ! IsStaged(stage, formfielddate) {
		return
	}

	formfielddate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormFieldDateTime(formfielddatetime *FormFieldDateTime) {

	// check if instance is already staged
	if ! IsStaged(stage, formfielddatetime) {
		return
	}

	formfielddatetime.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) {

	// check if instance is already staged
	if ! IsStaged(stage, formfieldfloat64) {
		return
	}

	formfieldfloat64.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormFieldInt(formfieldint *FormFieldInt) {

	// check if instance is already staged
	if ! IsStaged(stage, formfieldint) {
		return
	}

	formfieldint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormFieldSelect(formfieldselect *FormFieldSelect) {

	// check if instance is already staged
	if ! IsStaged(stage, formfieldselect) {
		return
	}

	formfieldselect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formfieldselect.Value != nil {
		UnstageBranch(stage, formfieldselect.Value)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _option := range formfieldselect.Options {
		UnstageBranch(stage, _option)
	}

}

func (stage *StageStruct) UnstageBranchFormFieldString(formfieldstring *FormFieldString) {

	// check if instance is already staged
	if ! IsStaged(stage, formfieldstring) {
		return
	}

	formfieldstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormFieldTime(formfieldtime *FormFieldTime) {

	// check if instance is already staged
	if ! IsStaged(stage, formfieldtime) {
		return
	}

	formfieldtime.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormGroup(formgroup *FormGroup) {

	// check if instance is already staged
	if ! IsStaged(stage, formgroup) {
		return
	}

	formgroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formdiv := range formgroup.FormDivs {
		UnstageBranch(stage, _formdiv)
	}

}

func (stage *StageStruct) UnstageBranchFormSortAssocButton(formsortassocbutton *FormSortAssocButton) {

	// check if instance is already staged
	if ! IsStaged(stage, formsortassocbutton) {
		return
	}

	formsortassocbutton.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOption(option *Option) {

	// check if instance is already staged
	if ! IsStaged(stage, option) {
		return
	}

	option.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRow(row *Row) {

	// check if instance is already staged
	if ! IsStaged(stage, row) {
		return
	}

	row.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _cell := range row.Cells {
		UnstageBranch(stage, _cell)
	}

}

func (stage *StageStruct) UnstageBranchTable(table *Table) {

	// check if instance is already staged
	if ! IsStaged(stage, table) {
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

}

