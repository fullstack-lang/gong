// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *CheckBox:
		ok = stage.IsStagedCheckBox(target)

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

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *CheckBox:
		ok = stage.IsStagedCheckBox(target)

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

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedCheckBox(checkbox *CheckBox) (ok bool) {

	_, ok = stage.CheckBoxs[checkbox]

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

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *CheckBox:
		stage.StageBranchCheckBox(target)

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

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchCheckBox(checkbox *CheckBox) {

	// check if instance is already staged
	if IsStaged(stage, checkbox) {
		return
	}

	checkbox.Stage(stage)

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

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *CheckBox:
		toT := CopyBranchCheckBox(mapOrigCopy, fromT)
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

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *CheckBox:
		stage.UnstageBranchCheckBox(target)

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

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchCheckBox(checkbox *CheckBox) {

	// check if instance is already staged
	if !IsStaged(stage, checkbox) {
		return
	}

	checkbox.Unstage(stage)

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

// insertion point for pointer reconstruction from references
func (reference *CheckBox) GongReconstructPointersFromReferences(stage *Stage, instance *CheckBox) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormDiv) GongReconstructPointersFromReferences(stage *Stage, instance *FormDiv) {
	// insertion point for pointers field
	if instance.FormEditAssocButton != nil {
		reference.FormEditAssocButton = stage.FormEditAssocButtons_reference[instance.FormEditAssocButton]
	}
	if instance.FormSortAssocButton != nil {
		reference.FormSortAssocButton = stage.FormSortAssocButtons_reference[instance.FormSortAssocButton]
	}
	// insertion point for slice of pointers field
	reference.FormFields = reference.FormFields[:0]
	for _, _b := range instance.FormFields {
		reference.FormFields = append(reference.FormFields, stage.FormFields_reference[_b])
	}
	reference.CheckBoxs = reference.CheckBoxs[:0]
	for _, _b := range instance.CheckBoxs {
		reference.CheckBoxs = append(reference.CheckBoxs, stage.CheckBoxs_reference[_b])
	}
}

func (reference *FormEditAssocButton) GongReconstructPointersFromReferences(stage *Stage, instance *FormEditAssocButton) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormField) GongReconstructPointersFromReferences(stage *Stage, instance *FormField) {
	// insertion point for pointers field
	if instance.FormFieldString != nil {
		reference.FormFieldString = stage.FormFieldStrings_reference[instance.FormFieldString]
	}
	if instance.FormFieldFloat64 != nil {
		reference.FormFieldFloat64 = stage.FormFieldFloat64s_reference[instance.FormFieldFloat64]
	}
	if instance.FormFieldInt != nil {
		reference.FormFieldInt = stage.FormFieldInts_reference[instance.FormFieldInt]
	}
	if instance.FormFieldDate != nil {
		reference.FormFieldDate = stage.FormFieldDates_reference[instance.FormFieldDate]
	}
	if instance.FormFieldTime != nil {
		reference.FormFieldTime = stage.FormFieldTimes_reference[instance.FormFieldTime]
	}
	if instance.FormFieldDateTime != nil {
		reference.FormFieldDateTime = stage.FormFieldDateTimes_reference[instance.FormFieldDateTime]
	}
	if instance.FormFieldSelect != nil {
		reference.FormFieldSelect = stage.FormFieldSelects_reference[instance.FormFieldSelect]
	}
	// insertion point for slice of pointers field
}

func (reference *FormFieldDate) GongReconstructPointersFromReferences(stage *Stage, instance *FormFieldDate) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormFieldDateTime) GongReconstructPointersFromReferences(stage *Stage, instance *FormFieldDateTime) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormFieldFloat64) GongReconstructPointersFromReferences(stage *Stage, instance *FormFieldFloat64) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormFieldInt) GongReconstructPointersFromReferences(stage *Stage, instance *FormFieldInt) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormFieldSelect) GongReconstructPointersFromReferences(stage *Stage, instance *FormFieldSelect) {
	// insertion point for pointers field
	if instance.Value != nil {
		reference.Value = stage.Options_reference[instance.Value]
	}
	// insertion point for slice of pointers field
	reference.Options = reference.Options[:0]
	for _, _b := range instance.Options {
		reference.Options = append(reference.Options, stage.Options_reference[_b])
	}
}

func (reference *FormFieldString) GongReconstructPointersFromReferences(stage *Stage, instance *FormFieldString) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormFieldTime) GongReconstructPointersFromReferences(stage *Stage, instance *FormFieldTime) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormGroup) GongReconstructPointersFromReferences(stage *Stage, instance *FormGroup) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.FormDivs = reference.FormDivs[:0]
	for _, _b := range instance.FormDivs {
		reference.FormDivs = append(reference.FormDivs, stage.FormDivs_reference[_b])
	}
}

func (reference *FormSortAssocButton) GongReconstructPointersFromReferences(stage *Stage, instance *FormSortAssocButton) {
	// insertion point for pointers field
	if instance.FormEditAssocButton != nil {
		reference.FormEditAssocButton = stage.FormEditAssocButtons_reference[instance.FormEditAssocButton]
	}
	// insertion point for slice of pointers field
}

func (reference *Option) GongReconstructPointersFromReferences(stage *Stage, instance *Option) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *CheckBox) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormDiv) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.FormEditAssocButton; _reference != nil {
		reference.FormEditAssocButton = nil
		if _instance, ok := stage.FormEditAssocButtons_instance[_reference]; ok {
			reference.FormEditAssocButton = _instance
		}
	}
	if _reference := reference.FormSortAssocButton; _reference != nil {
		reference.FormSortAssocButton = nil
		if _instance, ok := stage.FormSortAssocButtons_instance[_reference]; ok {
			reference.FormSortAssocButton = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _FormFields []*FormField
	for _, _reference := range reference.FormFields {
		if _instance, ok := stage.FormFields_instance[_reference]; ok {
			_FormFields = append(_FormFields, _instance)
		}
	}
	reference.FormFields = _FormFields
	var _CheckBoxs []*CheckBox
	for _, _reference := range reference.CheckBoxs {
		if _instance, ok := stage.CheckBoxs_instance[_reference]; ok {
			_CheckBoxs = append(_CheckBoxs, _instance)
		}
	}
	reference.CheckBoxs = _CheckBoxs
}

func (reference *FormEditAssocButton) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.FormFieldString; _reference != nil {
		reference.FormFieldString = nil
		if _instance, ok := stage.FormFieldStrings_instance[_reference]; ok {
			reference.FormFieldString = _instance
		}
	}
	if _reference := reference.FormFieldFloat64; _reference != nil {
		reference.FormFieldFloat64 = nil
		if _instance, ok := stage.FormFieldFloat64s_instance[_reference]; ok {
			reference.FormFieldFloat64 = _instance
		}
	}
	if _reference := reference.FormFieldInt; _reference != nil {
		reference.FormFieldInt = nil
		if _instance, ok := stage.FormFieldInts_instance[_reference]; ok {
			reference.FormFieldInt = _instance
		}
	}
	if _reference := reference.FormFieldDate; _reference != nil {
		reference.FormFieldDate = nil
		if _instance, ok := stage.FormFieldDates_instance[_reference]; ok {
			reference.FormFieldDate = _instance
		}
	}
	if _reference := reference.FormFieldTime; _reference != nil {
		reference.FormFieldTime = nil
		if _instance, ok := stage.FormFieldTimes_instance[_reference]; ok {
			reference.FormFieldTime = _instance
		}
	}
	if _reference := reference.FormFieldDateTime; _reference != nil {
		reference.FormFieldDateTime = nil
		if _instance, ok := stage.FormFieldDateTimes_instance[_reference]; ok {
			reference.FormFieldDateTime = _instance
		}
	}
	if _reference := reference.FormFieldSelect; _reference != nil {
		reference.FormFieldSelect = nil
		if _instance, ok := stage.FormFieldSelects_instance[_reference]; ok {
			reference.FormFieldSelect = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *FormFieldDate) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormFieldDateTime) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormFieldFloat64) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormFieldInt) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormFieldSelect) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Value; _reference != nil {
		reference.Value = nil
		if _instance, ok := stage.Options_instance[_reference]; ok {
			reference.Value = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Options []*Option
	for _, _reference := range reference.Options {
		if _instance, ok := stage.Options_instance[_reference]; ok {
			_Options = append(_Options, _instance)
		}
	}
	reference.Options = _Options
}

func (reference *FormFieldString) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormFieldTime) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormGroup) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _FormDivs []*FormDiv
	for _, _reference := range reference.FormDivs {
		if _instance, ok := stage.FormDivs_instance[_reference]; ok {
			_FormDivs = append(_FormDivs, _instance)
		}
	}
	reference.FormDivs = _FormDivs
}

func (reference *FormSortAssocButton) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.FormEditAssocButton; _reference != nil {
		reference.FormEditAssocButton = nil
		if _instance, ok := stage.FormEditAssocButtons_instance[_reference]; ok {
			reference.FormEditAssocButton = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Option) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
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
	if formdiv.IsADivider != formdivOther.IsADivider {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "IsADivider"))
	}
	if formdiv.IsAStartAccordionGroup != formdivOther.IsAStartAccordionGroup {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "IsAStartAccordionGroup"))
	}
	if formdiv.AccordionGroupName != formdivOther.AccordionGroupName {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "AccordionGroupName"))
	}
	if formdiv.IsAEndAccordionGroup != formdivOther.IsAEndAccordionGroup {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "IsAEndAccordionGroup"))
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
	if formgroup.TypeLabel != formgroupOther.TypeLabel {
		diffs = append(diffs, formgroup.GongMarshallField(stage, "TypeLabel"))
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
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
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
