// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

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

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

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

func (stage *Stage) IsStagedCheckBox(checkbox *CheckBox) (ok bool) {

	_, ok = stage.CheckBoxs[checkbox]

	return
}

func (stage *Stage) IsStagedDisplayedColumn(displayedcolumn *DisplayedColumn) (ok bool) {

	_, ok = stage.DisplayedColumns[displayedcolumn]

	return
}

func (stage *Stage) IsStagedFormDiv(formdiv *FormDiv) (ok bool) {

	_, ok = stage.FormDivs[formdiv]

	return
}

func (stage *Stage) IsStagedFormEditAssocButton(formeditassocbutton *FormEditAssocButton) (ok bool) {

	_, ok = stage.FormEditAssocButtons[formeditassocbutton]

	return
}

func (stage *Stage) IsStagedFormField(formfield *FormField) (ok bool) {

	_, ok = stage.FormFields[formfield]

	return
}

func (stage *Stage) IsStagedFormFieldDate(formfielddate *FormFieldDate) (ok bool) {

	_, ok = stage.FormFieldDates[formfielddate]

	return
}

func (stage *Stage) IsStagedFormFieldDateTime(formfielddatetime *FormFieldDateTime) (ok bool) {

	_, ok = stage.FormFieldDateTimes[formfielddatetime]

	return
}

func (stage *Stage) IsStagedFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) (ok bool) {

	_, ok = stage.FormFieldFloat64s[formfieldfloat64]

	return
}

func (stage *Stage) IsStagedFormFieldInt(formfieldint *FormFieldInt) (ok bool) {

	_, ok = stage.FormFieldInts[formfieldint]

	return
}

func (stage *Stage) IsStagedFormFieldSelect(formfieldselect *FormFieldSelect) (ok bool) {

	_, ok = stage.FormFieldSelects[formfieldselect]

	return
}

func (stage *Stage) IsStagedFormFieldString(formfieldstring *FormFieldString) (ok bool) {

	_, ok = stage.FormFieldStrings[formfieldstring]

	return
}

func (stage *Stage) IsStagedFormFieldTime(formfieldtime *FormFieldTime) (ok bool) {

	_, ok = stage.FormFieldTimes[formfieldtime]

	return
}

func (stage *Stage) IsStagedFormGroup(formgroup *FormGroup) (ok bool) {

	_, ok = stage.FormGroups[formgroup]

	return
}

func (stage *Stage) IsStagedFormSortAssocButton(formsortassocbutton *FormSortAssocButton) (ok bool) {

	_, ok = stage.FormSortAssocButtons[formsortassocbutton]

	return
}

func (stage *Stage) IsStagedOption(option *Option) (ok bool) {

	_, ok = stage.Options[option]

	return
}

func (stage *Stage) IsStagedRow(row *Row) (ok bool) {

	_, ok = stage.Rows[row]

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

func (stage *Stage) StageBranchCheckBox(checkbox *CheckBox) {

	// check if instance is already staged
	if IsStaged(stage, checkbox) {
		return
	}

	checkbox.Stage(stage)

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

func (stage *Stage) StageBranchFormDiv(formdiv *FormDiv) {

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

func (stage *Stage) StageBranchFormEditAssocButton(formeditassocbutton *FormEditAssocButton) {

	// check if instance is already staged
	if IsStaged(stage, formeditassocbutton) {
		return
	}

	formeditassocbutton.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFormField(formfield *FormField) {

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

func (stage *Stage) StageBranchFormFieldDate(formfielddate *FormFieldDate) {

	// check if instance is already staged
	if IsStaged(stage, formfielddate) {
		return
	}

	formfielddate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFormFieldDateTime(formfielddatetime *FormFieldDateTime) {

	// check if instance is already staged
	if IsStaged(stage, formfielddatetime) {
		return
	}

	formfielddatetime.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) {

	// check if instance is already staged
	if IsStaged(stage, formfieldfloat64) {
		return
	}

	formfieldfloat64.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFormFieldInt(formfieldint *FormFieldInt) {

	// check if instance is already staged
	if IsStaged(stage, formfieldint) {
		return
	}

	formfieldint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFormFieldSelect(formfieldselect *FormFieldSelect) {

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

func (stage *Stage) StageBranchFormFieldString(formfieldstring *FormFieldString) {

	// check if instance is already staged
	if IsStaged(stage, formfieldstring) {
		return
	}

	formfieldstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFormFieldTime(formfieldtime *FormFieldTime) {

	// check if instance is already staged
	if IsStaged(stage, formfieldtime) {
		return
	}

	formfieldtime.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchFormGroup(formgroup *FormGroup) {

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

func (stage *Stage) StageBranchFormSortAssocButton(formsortassocbutton *FormSortAssocButton) {

	// check if instance is already staged
	if IsStaged(stage, formsortassocbutton) {
		return
	}

	formsortassocbutton.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formsortassocbutton.FormEditAssocButton != nil {
		StageBranch(stage, formsortassocbutton.FormEditAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchOption(option *Option) {

	// check if instance is already staged
	if IsStaged(stage, option) {
		return
	}

	option.Stage(stage)

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

	case *CheckBox:
		toT := CopyBranchCheckBox(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DisplayedColumn:
		toT := CopyBranchDisplayedColumn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormDiv:
		toT := CopyBranchFormDiv(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormEditAssocButton:
		toT := CopyBranchFormEditAssocButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormField:
		toT := CopyBranchFormField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldDate:
		toT := CopyBranchFormFieldDate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldDateTime:
		toT := CopyBranchFormFieldDateTime(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldFloat64:
		toT := CopyBranchFormFieldFloat64(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldInt:
		toT := CopyBranchFormFieldInt(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldSelect:
		toT := CopyBranchFormFieldSelect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldString:
		toT := CopyBranchFormFieldString(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldTime:
		toT := CopyBranchFormFieldTime(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormGroup:
		toT := CopyBranchFormGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormSortAssocButton:
		toT := CopyBranchFormSortAssocButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Option:
		toT := CopyBranchOption(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Row:
		toT := CopyBranchRow(mapOrigCopy, fromT)
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

func CopyBranchCheckBox(mapOrigCopy map[any]any, checkboxFrom *CheckBox) (checkboxTo *CheckBox) {

	// checkboxFrom has already been copied
	if _checkboxTo, ok := mapOrigCopy[checkboxFrom]; ok {
		checkboxTo = _checkboxTo.(*CheckBox)
		return
	}

	checkboxTo = new(CheckBox)
	mapOrigCopy[checkboxFrom] = checkboxTo
	checkboxFrom.CopyBasicFields(checkboxTo)

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

func CopyBranchFormDiv(mapOrigCopy map[any]any, formdivFrom *FormDiv) (formdivTo *FormDiv) {

	// formdivFrom has already been copied
	if _formdivTo, ok := mapOrigCopy[formdivFrom]; ok {
		formdivTo = _formdivTo.(*FormDiv)
		return
	}

	formdivTo = new(FormDiv)
	mapOrigCopy[formdivFrom] = formdivTo
	formdivFrom.CopyBasicFields(formdivTo)

	//insertion point for the staging of instances referenced by pointers
	if formdivFrom.FormEditAssocButton != nil {
		formdivTo.FormEditAssocButton = CopyBranchFormEditAssocButton(mapOrigCopy, formdivFrom.FormEditAssocButton)
	}
	if formdivFrom.FormSortAssocButton != nil {
		formdivTo.FormSortAssocButton = CopyBranchFormSortAssocButton(mapOrigCopy, formdivFrom.FormSortAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formfield := range formdivFrom.FormFields {
		formdivTo.FormFields = append(formdivTo.FormFields, CopyBranchFormField(mapOrigCopy, _formfield))
	}
	for _, _checkbox := range formdivFrom.CheckBoxs {
		formdivTo.CheckBoxs = append(formdivTo.CheckBoxs, CopyBranchCheckBox(mapOrigCopy, _checkbox))
	}

	return
}

func CopyBranchFormEditAssocButton(mapOrigCopy map[any]any, formeditassocbuttonFrom *FormEditAssocButton) (formeditassocbuttonTo *FormEditAssocButton) {

	// formeditassocbuttonFrom has already been copied
	if _formeditassocbuttonTo, ok := mapOrigCopy[formeditassocbuttonFrom]; ok {
		formeditassocbuttonTo = _formeditassocbuttonTo.(*FormEditAssocButton)
		return
	}

	formeditassocbuttonTo = new(FormEditAssocButton)
	mapOrigCopy[formeditassocbuttonFrom] = formeditassocbuttonTo
	formeditassocbuttonFrom.CopyBasicFields(formeditassocbuttonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormField(mapOrigCopy map[any]any, formfieldFrom *FormField) (formfieldTo *FormField) {

	// formfieldFrom has already been copied
	if _formfieldTo, ok := mapOrigCopy[formfieldFrom]; ok {
		formfieldTo = _formfieldTo.(*FormField)
		return
	}

	formfieldTo = new(FormField)
	mapOrigCopy[formfieldFrom] = formfieldTo
	formfieldFrom.CopyBasicFields(formfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if formfieldFrom.FormFieldString != nil {
		formfieldTo.FormFieldString = CopyBranchFormFieldString(mapOrigCopy, formfieldFrom.FormFieldString)
	}
	if formfieldFrom.FormFieldFloat64 != nil {
		formfieldTo.FormFieldFloat64 = CopyBranchFormFieldFloat64(mapOrigCopy, formfieldFrom.FormFieldFloat64)
	}
	if formfieldFrom.FormFieldInt != nil {
		formfieldTo.FormFieldInt = CopyBranchFormFieldInt(mapOrigCopy, formfieldFrom.FormFieldInt)
	}
	if formfieldFrom.FormFieldDate != nil {
		formfieldTo.FormFieldDate = CopyBranchFormFieldDate(mapOrigCopy, formfieldFrom.FormFieldDate)
	}
	if formfieldFrom.FormFieldTime != nil {
		formfieldTo.FormFieldTime = CopyBranchFormFieldTime(mapOrigCopy, formfieldFrom.FormFieldTime)
	}
	if formfieldFrom.FormFieldDateTime != nil {
		formfieldTo.FormFieldDateTime = CopyBranchFormFieldDateTime(mapOrigCopy, formfieldFrom.FormFieldDateTime)
	}
	if formfieldFrom.FormFieldSelect != nil {
		formfieldTo.FormFieldSelect = CopyBranchFormFieldSelect(mapOrigCopy, formfieldFrom.FormFieldSelect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormFieldDate(mapOrigCopy map[any]any, formfielddateFrom *FormFieldDate) (formfielddateTo *FormFieldDate) {

	// formfielddateFrom has already been copied
	if _formfielddateTo, ok := mapOrigCopy[formfielddateFrom]; ok {
		formfielddateTo = _formfielddateTo.(*FormFieldDate)
		return
	}

	formfielddateTo = new(FormFieldDate)
	mapOrigCopy[formfielddateFrom] = formfielddateTo
	formfielddateFrom.CopyBasicFields(formfielddateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormFieldDateTime(mapOrigCopy map[any]any, formfielddatetimeFrom *FormFieldDateTime) (formfielddatetimeTo *FormFieldDateTime) {

	// formfielddatetimeFrom has already been copied
	if _formfielddatetimeTo, ok := mapOrigCopy[formfielddatetimeFrom]; ok {
		formfielddatetimeTo = _formfielddatetimeTo.(*FormFieldDateTime)
		return
	}

	formfielddatetimeTo = new(FormFieldDateTime)
	mapOrigCopy[formfielddatetimeFrom] = formfielddatetimeTo
	formfielddatetimeFrom.CopyBasicFields(formfielddatetimeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormFieldFloat64(mapOrigCopy map[any]any, formfieldfloat64From *FormFieldFloat64) (formfieldfloat64To *FormFieldFloat64) {

	// formfieldfloat64From has already been copied
	if _formfieldfloat64To, ok := mapOrigCopy[formfieldfloat64From]; ok {
		formfieldfloat64To = _formfieldfloat64To.(*FormFieldFloat64)
		return
	}

	formfieldfloat64To = new(FormFieldFloat64)
	mapOrigCopy[formfieldfloat64From] = formfieldfloat64To
	formfieldfloat64From.CopyBasicFields(formfieldfloat64To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormFieldInt(mapOrigCopy map[any]any, formfieldintFrom *FormFieldInt) (formfieldintTo *FormFieldInt) {

	// formfieldintFrom has already been copied
	if _formfieldintTo, ok := mapOrigCopy[formfieldintFrom]; ok {
		formfieldintTo = _formfieldintTo.(*FormFieldInt)
		return
	}

	formfieldintTo = new(FormFieldInt)
	mapOrigCopy[formfieldintFrom] = formfieldintTo
	formfieldintFrom.CopyBasicFields(formfieldintTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormFieldSelect(mapOrigCopy map[any]any, formfieldselectFrom *FormFieldSelect) (formfieldselectTo *FormFieldSelect) {

	// formfieldselectFrom has already been copied
	if _formfieldselectTo, ok := mapOrigCopy[formfieldselectFrom]; ok {
		formfieldselectTo = _formfieldselectTo.(*FormFieldSelect)
		return
	}

	formfieldselectTo = new(FormFieldSelect)
	mapOrigCopy[formfieldselectFrom] = formfieldselectTo
	formfieldselectFrom.CopyBasicFields(formfieldselectTo)

	//insertion point for the staging of instances referenced by pointers
	if formfieldselectFrom.Value != nil {
		formfieldselectTo.Value = CopyBranchOption(mapOrigCopy, formfieldselectFrom.Value)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _option := range formfieldselectFrom.Options {
		formfieldselectTo.Options = append(formfieldselectTo.Options, CopyBranchOption(mapOrigCopy, _option))
	}

	return
}

func CopyBranchFormFieldString(mapOrigCopy map[any]any, formfieldstringFrom *FormFieldString) (formfieldstringTo *FormFieldString) {

	// formfieldstringFrom has already been copied
	if _formfieldstringTo, ok := mapOrigCopy[formfieldstringFrom]; ok {
		formfieldstringTo = _formfieldstringTo.(*FormFieldString)
		return
	}

	formfieldstringTo = new(FormFieldString)
	mapOrigCopy[formfieldstringFrom] = formfieldstringTo
	formfieldstringFrom.CopyBasicFields(formfieldstringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormFieldTime(mapOrigCopy map[any]any, formfieldtimeFrom *FormFieldTime) (formfieldtimeTo *FormFieldTime) {

	// formfieldtimeFrom has already been copied
	if _formfieldtimeTo, ok := mapOrigCopy[formfieldtimeFrom]; ok {
		formfieldtimeTo = _formfieldtimeTo.(*FormFieldTime)
		return
	}

	formfieldtimeTo = new(FormFieldTime)
	mapOrigCopy[formfieldtimeFrom] = formfieldtimeTo
	formfieldtimeFrom.CopyBasicFields(formfieldtimeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormGroup(mapOrigCopy map[any]any, formgroupFrom *FormGroup) (formgroupTo *FormGroup) {

	// formgroupFrom has already been copied
	if _formgroupTo, ok := mapOrigCopy[formgroupFrom]; ok {
		formgroupTo = _formgroupTo.(*FormGroup)
		return
	}

	formgroupTo = new(FormGroup)
	mapOrigCopy[formgroupFrom] = formgroupTo
	formgroupFrom.CopyBasicFields(formgroupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formdiv := range formgroupFrom.FormDivs {
		formgroupTo.FormDivs = append(formgroupTo.FormDivs, CopyBranchFormDiv(mapOrigCopy, _formdiv))
	}

	return
}

func CopyBranchFormSortAssocButton(mapOrigCopy map[any]any, formsortassocbuttonFrom *FormSortAssocButton) (formsortassocbuttonTo *FormSortAssocButton) {

	// formsortassocbuttonFrom has already been copied
	if _formsortassocbuttonTo, ok := mapOrigCopy[formsortassocbuttonFrom]; ok {
		formsortassocbuttonTo = _formsortassocbuttonTo.(*FormSortAssocButton)
		return
	}

	formsortassocbuttonTo = new(FormSortAssocButton)
	mapOrigCopy[formsortassocbuttonFrom] = formsortassocbuttonTo
	formsortassocbuttonFrom.CopyBasicFields(formsortassocbuttonTo)

	//insertion point for the staging of instances referenced by pointers
	if formsortassocbuttonFrom.FormEditAssocButton != nil {
		formsortassocbuttonTo.FormEditAssocButton = CopyBranchFormEditAssocButton(mapOrigCopy, formsortassocbuttonFrom.FormEditAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOption(mapOrigCopy map[any]any, optionFrom *Option) (optionTo *Option) {

	// optionFrom has already been copied
	if _optionTo, ok := mapOrigCopy[optionFrom]; ok {
		optionTo = _optionTo.(*Option)
		return
	}

	optionTo = new(Option)
	mapOrigCopy[optionFrom] = optionTo
	optionFrom.CopyBasicFields(optionTo)

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

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

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

func (stage *Stage) UnstageBranchCheckBox(checkbox *CheckBox) {

	// check if instance is already staged
	if !IsStaged(stage, checkbox) {
		return
	}

	checkbox.Unstage(stage)

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

func (stage *Stage) UnstageBranchFormDiv(formdiv *FormDiv) {

	// check if instance is already staged
	if !IsStaged(stage, formdiv) {
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

func (stage *Stage) UnstageBranchFormEditAssocButton(formeditassocbutton *FormEditAssocButton) {

	// check if instance is already staged
	if !IsStaged(stage, formeditassocbutton) {
		return
	}

	formeditassocbutton.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFormField(formfield *FormField) {

	// check if instance is already staged
	if !IsStaged(stage, formfield) {
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

func (stage *Stage) UnstageBranchFormFieldDate(formfielddate *FormFieldDate) {

	// check if instance is already staged
	if !IsStaged(stage, formfielddate) {
		return
	}

	formfielddate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFormFieldDateTime(formfielddatetime *FormFieldDateTime) {

	// check if instance is already staged
	if !IsStaged(stage, formfielddatetime) {
		return
	}

	formfielddatetime.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) {

	// check if instance is already staged
	if !IsStaged(stage, formfieldfloat64) {
		return
	}

	formfieldfloat64.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFormFieldInt(formfieldint *FormFieldInt) {

	// check if instance is already staged
	if !IsStaged(stage, formfieldint) {
		return
	}

	formfieldint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFormFieldSelect(formfieldselect *FormFieldSelect) {

	// check if instance is already staged
	if !IsStaged(stage, formfieldselect) {
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

func (stage *Stage) UnstageBranchFormFieldString(formfieldstring *FormFieldString) {

	// check if instance is already staged
	if !IsStaged(stage, formfieldstring) {
		return
	}

	formfieldstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFormFieldTime(formfieldtime *FormFieldTime) {

	// check if instance is already staged
	if !IsStaged(stage, formfieldtime) {
		return
	}

	formfieldtime.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchFormGroup(formgroup *FormGroup) {

	// check if instance is already staged
	if !IsStaged(stage, formgroup) {
		return
	}

	formgroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formdiv := range formgroup.FormDivs {
		UnstageBranch(stage, _formdiv)
	}

}

func (stage *Stage) UnstageBranchFormSortAssocButton(formsortassocbutton *FormSortAssocButton) {

	// check if instance is already staged
	if !IsStaged(stage, formsortassocbutton) {
		return
	}

	formsortassocbutton.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formsortassocbutton.FormEditAssocButton != nil {
		UnstageBranch(stage, formsortassocbutton.FormEditAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchOption(option *Option) {

	// check if instance is already staged
	if !IsStaged(stage, option) {
		return
	}

	option.Unstage(stage)

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

}

// insertion point for diff per struct
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
func (checkbox *CheckBox) GongDiff(stage *Stage, checkboxOther *CheckBox) (diffs []string) {
	// insertion point for field diffs
	if checkbox.Name != checkboxOther.Name {
		diffs = append(diffs, checkbox.GongMarshallField(stage, "Name"))
	}
	if checkbox.Value != checkboxOther.Value {
		diffs = append(diffs, checkbox.GongMarshallField(stage, "Value"))
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
func (formdiv *FormDiv) GongDiff(stage *Stage, formdivOther *FormDiv) (diffs []string) {
	// insertion point for field diffs
	if formdiv.Name != formdivOther.Name {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "Name"))
	}
	FormFieldsDifferent := false
	if len(formdiv.FormFields) != len(formdivOther.FormFields) {
		FormFieldsDifferent = true
	} else {
		for i := range formdiv.FormFields {
			if (formdiv.FormFields[i] == nil) != (formdivOther.FormFields[i] == nil) {
				FormFieldsDifferent = true
				break
			} else if formdiv.FormFields[i] != nil && formdivOther.FormFields[i] != nil {
				// this is a pointer comparaison
				if formdiv.FormFields[i] != formdivOther.FormFields[i] {
					FormFieldsDifferent = true
					break
				}
			}
		}
	}
	if FormFieldsDifferent {
		ops := Diff(stage, formdiv, formdivOther, "FormFields", formdivOther.FormFields, formdiv.FormFields)
		diffs = append(diffs, ops)
	}
	CheckBoxsDifferent := false
	if len(formdiv.CheckBoxs) != len(formdivOther.CheckBoxs) {
		CheckBoxsDifferent = true
	} else {
		for i := range formdiv.CheckBoxs {
			if (formdiv.CheckBoxs[i] == nil) != (formdivOther.CheckBoxs[i] == nil) {
				CheckBoxsDifferent = true
				break
			} else if formdiv.CheckBoxs[i] != nil && formdivOther.CheckBoxs[i] != nil {
				// this is a pointer comparaison
				if formdiv.CheckBoxs[i] != formdivOther.CheckBoxs[i] {
					CheckBoxsDifferent = true
					break
				}
			}
		}
	}
	if CheckBoxsDifferent {
		ops := Diff(stage, formdiv, formdivOther, "CheckBoxs", formdivOther.CheckBoxs, formdiv.CheckBoxs)
		diffs = append(diffs, ops)
	}
	if (formdiv.FormEditAssocButton == nil) != (formdivOther.FormEditAssocButton == nil) {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "FormEditAssocButton"))
	} else if formdiv.FormEditAssocButton != nil && formdivOther.FormEditAssocButton != nil {
		if formdiv.FormEditAssocButton != formdivOther.FormEditAssocButton {
			diffs = append(diffs, formdiv.GongMarshallField(stage, "FormEditAssocButton"))
		}
	}
	if (formdiv.FormSortAssocButton == nil) != (formdivOther.FormSortAssocButton == nil) {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "FormSortAssocButton"))
	} else if formdiv.FormSortAssocButton != nil && formdivOther.FormSortAssocButton != nil {
		if formdiv.FormSortAssocButton != formdivOther.FormSortAssocButton {
			diffs = append(diffs, formdiv.GongMarshallField(stage, "FormSortAssocButton"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formeditassocbutton *FormEditAssocButton) GongDiff(stage *Stage, formeditassocbuttonOther *FormEditAssocButton) (diffs []string) {
	// insertion point for field diffs
	if formeditassocbutton.Name != formeditassocbuttonOther.Name {
		diffs = append(diffs, formeditassocbutton.GongMarshallField(stage, "Name"))
	}
	if formeditassocbutton.Label != formeditassocbuttonOther.Label {
		diffs = append(diffs, formeditassocbutton.GongMarshallField(stage, "Label"))
	}
	if formeditassocbutton.AssociationStorage != formeditassocbuttonOther.AssociationStorage {
		diffs = append(diffs, formeditassocbutton.GongMarshallField(stage, "AssociationStorage"))
	}
	if formeditassocbutton.HasChanged != formeditassocbuttonOther.HasChanged {
		diffs = append(diffs, formeditassocbutton.GongMarshallField(stage, "HasChanged"))
	}
	if formeditassocbutton.IsForSavePurpose != formeditassocbuttonOther.IsForSavePurpose {
		diffs = append(diffs, formeditassocbutton.GongMarshallField(stage, "IsForSavePurpose"))
	}
	if formeditassocbutton.HasToolTip != formeditassocbuttonOther.HasToolTip {
		diffs = append(diffs, formeditassocbutton.GongMarshallField(stage, "HasToolTip"))
	}
	if formeditassocbutton.ToolTipText != formeditassocbuttonOther.ToolTipText {
		diffs = append(diffs, formeditassocbutton.GongMarshallField(stage, "ToolTipText"))
	}
	if formeditassocbutton.MatTooltipShowDelay != formeditassocbuttonOther.MatTooltipShowDelay {
		diffs = append(diffs, formeditassocbutton.GongMarshallField(stage, "MatTooltipShowDelay"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formfield *FormField) GongDiff(stage *Stage, formfieldOther *FormField) (diffs []string) {
	// insertion point for field diffs
	if formfield.Name != formfieldOther.Name {
		diffs = append(diffs, formfield.GongMarshallField(stage, "Name"))
	}
	if formfield.InputTypeEnum != formfieldOther.InputTypeEnum {
		diffs = append(diffs, formfield.GongMarshallField(stage, "InputTypeEnum"))
	}
	if formfield.Label != formfieldOther.Label {
		diffs = append(diffs, formfield.GongMarshallField(stage, "Label"))
	}
	if formfield.Placeholder != formfieldOther.Placeholder {
		diffs = append(diffs, formfield.GongMarshallField(stage, "Placeholder"))
	}
	if (formfield.FormFieldString == nil) != (formfieldOther.FormFieldString == nil) {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldString"))
	} else if formfield.FormFieldString != nil && formfieldOther.FormFieldString != nil {
		if formfield.FormFieldString != formfieldOther.FormFieldString {
			diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldString"))
		}
	}
	if (formfield.FormFieldFloat64 == nil) != (formfieldOther.FormFieldFloat64 == nil) {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldFloat64"))
	} else if formfield.FormFieldFloat64 != nil && formfieldOther.FormFieldFloat64 != nil {
		if formfield.FormFieldFloat64 != formfieldOther.FormFieldFloat64 {
			diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldFloat64"))
		}
	}
	if (formfield.FormFieldInt == nil) != (formfieldOther.FormFieldInt == nil) {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldInt"))
	} else if formfield.FormFieldInt != nil && formfieldOther.FormFieldInt != nil {
		if formfield.FormFieldInt != formfieldOther.FormFieldInt {
			diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldInt"))
		}
	}
	if (formfield.FormFieldDate == nil) != (formfieldOther.FormFieldDate == nil) {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldDate"))
	} else if formfield.FormFieldDate != nil && formfieldOther.FormFieldDate != nil {
		if formfield.FormFieldDate != formfieldOther.FormFieldDate {
			diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldDate"))
		}
	}
	if (formfield.FormFieldTime == nil) != (formfieldOther.FormFieldTime == nil) {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldTime"))
	} else if formfield.FormFieldTime != nil && formfieldOther.FormFieldTime != nil {
		if formfield.FormFieldTime != formfieldOther.FormFieldTime {
			diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldTime"))
		}
	}
	if (formfield.FormFieldDateTime == nil) != (formfieldOther.FormFieldDateTime == nil) {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldDateTime"))
	} else if formfield.FormFieldDateTime != nil && formfieldOther.FormFieldDateTime != nil {
		if formfield.FormFieldDateTime != formfieldOther.FormFieldDateTime {
			diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldDateTime"))
		}
	}
	if (formfield.FormFieldSelect == nil) != (formfieldOther.FormFieldSelect == nil) {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldSelect"))
	} else if formfield.FormFieldSelect != nil && formfieldOther.FormFieldSelect != nil {
		if formfield.FormFieldSelect != formfieldOther.FormFieldSelect {
			diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldSelect"))
		}
	}
	if formfield.HasBespokeWidth != formfieldOther.HasBespokeWidth {
		diffs = append(diffs, formfield.GongMarshallField(stage, "HasBespokeWidth"))
	}
	if formfield.BespokeWidthPx != formfieldOther.BespokeWidthPx {
		diffs = append(diffs, formfield.GongMarshallField(stage, "BespokeWidthPx"))
	}
	if formfield.HasBespokeHeight != formfieldOther.HasBespokeHeight {
		diffs = append(diffs, formfield.GongMarshallField(stage, "HasBespokeHeight"))
	}
	if formfield.BespokeHeightPx != formfieldOther.BespokeHeightPx {
		diffs = append(diffs, formfield.GongMarshallField(stage, "BespokeHeightPx"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formfielddate *FormFieldDate) GongDiff(stage *Stage, formfielddateOther *FormFieldDate) (diffs []string) {
	// insertion point for field diffs
	if formfielddate.Name != formfielddateOther.Name {
		diffs = append(diffs, formfielddate.GongMarshallField(stage, "Name"))
	}
	if formfielddate.Value != formfielddateOther.Value {
		diffs = append(diffs, formfielddate.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formfielddatetime *FormFieldDateTime) GongDiff(stage *Stage, formfielddatetimeOther *FormFieldDateTime) (diffs []string) {
	// insertion point for field diffs
	if formfielddatetime.Name != formfielddatetimeOther.Name {
		diffs = append(diffs, formfielddatetime.GongMarshallField(stage, "Name"))
	}
	if formfielddatetime.Value != formfielddatetimeOther.Value {
		diffs = append(diffs, formfielddatetime.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formfieldfloat64 *FormFieldFloat64) GongDiff(stage *Stage, formfieldfloat64Other *FormFieldFloat64) (diffs []string) {
	// insertion point for field diffs
	if formfieldfloat64.Name != formfieldfloat64Other.Name {
		diffs = append(diffs, formfieldfloat64.GongMarshallField(stage, "Name"))
	}
	if formfieldfloat64.Value != formfieldfloat64Other.Value {
		diffs = append(diffs, formfieldfloat64.GongMarshallField(stage, "Value"))
	}
	if formfieldfloat64.HasMinValidator != formfieldfloat64Other.HasMinValidator {
		diffs = append(diffs, formfieldfloat64.GongMarshallField(stage, "HasMinValidator"))
	}
	if formfieldfloat64.MinValue != formfieldfloat64Other.MinValue {
		diffs = append(diffs, formfieldfloat64.GongMarshallField(stage, "MinValue"))
	}
	if formfieldfloat64.HasMaxValidator != formfieldfloat64Other.HasMaxValidator {
		diffs = append(diffs, formfieldfloat64.GongMarshallField(stage, "HasMaxValidator"))
	}
	if formfieldfloat64.MaxValue != formfieldfloat64Other.MaxValue {
		diffs = append(diffs, formfieldfloat64.GongMarshallField(stage, "MaxValue"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formfieldint *FormFieldInt) GongDiff(stage *Stage, formfieldintOther *FormFieldInt) (diffs []string) {
	// insertion point for field diffs
	if formfieldint.Name != formfieldintOther.Name {
		diffs = append(diffs, formfieldint.GongMarshallField(stage, "Name"))
	}
	if formfieldint.Value != formfieldintOther.Value {
		diffs = append(diffs, formfieldint.GongMarshallField(stage, "Value"))
	}
	if formfieldint.HasMinValidator != formfieldintOther.HasMinValidator {
		diffs = append(diffs, formfieldint.GongMarshallField(stage, "HasMinValidator"))
	}
	if formfieldint.MinValue != formfieldintOther.MinValue {
		diffs = append(diffs, formfieldint.GongMarshallField(stage, "MinValue"))
	}
	if formfieldint.HasMaxValidator != formfieldintOther.HasMaxValidator {
		diffs = append(diffs, formfieldint.GongMarshallField(stage, "HasMaxValidator"))
	}
	if formfieldint.MaxValue != formfieldintOther.MaxValue {
		diffs = append(diffs, formfieldint.GongMarshallField(stage, "MaxValue"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formfieldselect *FormFieldSelect) GongDiff(stage *Stage, formfieldselectOther *FormFieldSelect) (diffs []string) {
	// insertion point for field diffs
	if formfieldselect.Name != formfieldselectOther.Name {
		diffs = append(diffs, formfieldselect.GongMarshallField(stage, "Name"))
	}
	if (formfieldselect.Value == nil) != (formfieldselectOther.Value == nil) {
		diffs = append(diffs, formfieldselect.GongMarshallField(stage, "Value"))
	} else if formfieldselect.Value != nil && formfieldselectOther.Value != nil {
		if formfieldselect.Value != formfieldselectOther.Value {
			diffs = append(diffs, formfieldselect.GongMarshallField(stage, "Value"))
		}
	}
	OptionsDifferent := false
	if len(formfieldselect.Options) != len(formfieldselectOther.Options) {
		OptionsDifferent = true
	} else {
		for i := range formfieldselect.Options {
			if (formfieldselect.Options[i] == nil) != (formfieldselectOther.Options[i] == nil) {
				OptionsDifferent = true
				break
			} else if formfieldselect.Options[i] != nil && formfieldselectOther.Options[i] != nil {
				// this is a pointer comparaison
				if formfieldselect.Options[i] != formfieldselectOther.Options[i] {
					OptionsDifferent = true
					break
				}
			}
		}
	}
	if OptionsDifferent {
		ops := Diff(stage, formfieldselect, formfieldselectOther, "Options", formfieldselectOther.Options, formfieldselect.Options)
		diffs = append(diffs, ops)
	}
	if formfieldselect.CanBeEmpty != formfieldselectOther.CanBeEmpty {
		diffs = append(diffs, formfieldselect.GongMarshallField(stage, "CanBeEmpty"))
	}
	if formfieldselect.PreserveInitialOrder != formfieldselectOther.PreserveInitialOrder {
		diffs = append(diffs, formfieldselect.GongMarshallField(stage, "PreserveInitialOrder"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formfieldstring *FormFieldString) GongDiff(stage *Stage, formfieldstringOther *FormFieldString) (diffs []string) {
	// insertion point for field diffs
	if formfieldstring.Name != formfieldstringOther.Name {
		diffs = append(diffs, formfieldstring.GongMarshallField(stage, "Name"))
	}
	if formfieldstring.Value != formfieldstringOther.Value {
		diffs = append(diffs, formfieldstring.GongMarshallField(stage, "Value"))
	}
	if formfieldstring.IsTextArea != formfieldstringOther.IsTextArea {
		diffs = append(diffs, formfieldstring.GongMarshallField(stage, "IsTextArea"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formfieldtime *FormFieldTime) GongDiff(stage *Stage, formfieldtimeOther *FormFieldTime) (diffs []string) {
	// insertion point for field diffs
	if formfieldtime.Name != formfieldtimeOther.Name {
		diffs = append(diffs, formfieldtime.GongMarshallField(stage, "Name"))
	}
	if formfieldtime.Value != formfieldtimeOther.Value {
		diffs = append(diffs, formfieldtime.GongMarshallField(stage, "Value"))
	}
	if formfieldtime.Step != formfieldtimeOther.Step {
		diffs = append(diffs, formfieldtime.GongMarshallField(stage, "Step"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formgroup *FormGroup) GongDiff(stage *Stage, formgroupOther *FormGroup) (diffs []string) {
	// insertion point for field diffs
	if formgroup.Name != formgroupOther.Name {
		diffs = append(diffs, formgroup.GongMarshallField(stage, "Name"))
	}
	if formgroup.Label != formgroupOther.Label {
		diffs = append(diffs, formgroup.GongMarshallField(stage, "Label"))
	}
	FormDivsDifferent := false
	if len(formgroup.FormDivs) != len(formgroupOther.FormDivs) {
		FormDivsDifferent = true
	} else {
		for i := range formgroup.FormDivs {
			if (formgroup.FormDivs[i] == nil) != (formgroupOther.FormDivs[i] == nil) {
				FormDivsDifferent = true
				break
			} else if formgroup.FormDivs[i] != nil && formgroupOther.FormDivs[i] != nil {
				// this is a pointer comparaison
				if formgroup.FormDivs[i] != formgroupOther.FormDivs[i] {
					FormDivsDifferent = true
					break
				}
			}
		}
	}
	if FormDivsDifferent {
		ops := Diff(stage, formgroup, formgroupOther, "FormDivs", formgroupOther.FormDivs, formgroup.FormDivs)
		diffs = append(diffs, ops)
	}
	if formgroup.HasSuppressButton != formgroupOther.HasSuppressButton {
		diffs = append(diffs, formgroup.GongMarshallField(stage, "HasSuppressButton"))
	}
	if formgroup.HasSuppressButtonBeenPressed != formgroupOther.HasSuppressButtonBeenPressed {
		diffs = append(diffs, formgroup.GongMarshallField(stage, "HasSuppressButtonBeenPressed"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (formsortassocbutton *FormSortAssocButton) GongDiff(stage *Stage, formsortassocbuttonOther *FormSortAssocButton) (diffs []string) {
	// insertion point for field diffs
	if formsortassocbutton.Name != formsortassocbuttonOther.Name {
		diffs = append(diffs, formsortassocbutton.GongMarshallField(stage, "Name"))
	}
	if formsortassocbutton.Label != formsortassocbuttonOther.Label {
		diffs = append(diffs, formsortassocbutton.GongMarshallField(stage, "Label"))
	}
	if formsortassocbutton.HasToolTip != formsortassocbuttonOther.HasToolTip {
		diffs = append(diffs, formsortassocbutton.GongMarshallField(stage, "HasToolTip"))
	}
	if formsortassocbutton.ToolTipText != formsortassocbuttonOther.ToolTipText {
		diffs = append(diffs, formsortassocbutton.GongMarshallField(stage, "ToolTipText"))
	}
	if formsortassocbutton.MatTooltipShowDelay != formsortassocbuttonOther.MatTooltipShowDelay {
		diffs = append(diffs, formsortassocbutton.GongMarshallField(stage, "MatTooltipShowDelay"))
	}
	if (formsortassocbutton.FormEditAssocButton == nil) != (formsortassocbuttonOther.FormEditAssocButton == nil) {
		diffs = append(diffs, formsortassocbutton.GongMarshallField(stage, "FormEditAssocButton"))
	} else if formsortassocbutton.FormEditAssocButton != nil && formsortassocbuttonOther.FormEditAssocButton != nil {
		if formsortassocbutton.FormEditAssocButton != formsortassocbuttonOther.FormEditAssocButton {
			diffs = append(diffs, formsortassocbutton.GongMarshallField(stage, "FormEditAssocButton"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (option *Option) GongDiff(stage *Stage, optionOther *Option) (diffs []string) {
	// insertion point for field diffs
	if option.Name != optionOther.Name {
		diffs = append(diffs, option.GongMarshallField(stage, "Name"))
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
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
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
