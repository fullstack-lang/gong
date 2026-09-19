// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (checkbox *CheckBox) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.CheckBoxs[checkbox]

	return
}

func (stage *Stage) IsStagedCheckBox(checkbox *CheckBox) (ok bool) {

	return checkbox.GongIsStaged(stage)
}

func (formdiv *FormDiv) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormDivs[formdiv]

	return
}

func (stage *Stage) IsStagedFormDiv(formdiv *FormDiv) (ok bool) {

	return formdiv.GongIsStaged(stage)
}

func (formeditassocbutton *FormEditAssocButton) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormEditAssocButtons[formeditassocbutton]

	return
}

func (stage *Stage) IsStagedFormEditAssocButton(formeditassocbutton *FormEditAssocButton) (ok bool) {

	return formeditassocbutton.GongIsStaged(stage)
}

func (formfield *FormField) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormFields[formfield]

	return
}

func (stage *Stage) IsStagedFormField(formfield *FormField) (ok bool) {

	return formfield.GongIsStaged(stage)
}

func (formfielddate *FormFieldDate) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormFieldDates[formfielddate]

	return
}

func (stage *Stage) IsStagedFormFieldDate(formfielddate *FormFieldDate) (ok bool) {

	return formfielddate.GongIsStaged(stage)
}

func (formfielddatetime *FormFieldDateTime) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormFieldDateTimes[formfielddatetime]

	return
}

func (stage *Stage) IsStagedFormFieldDateTime(formfielddatetime *FormFieldDateTime) (ok bool) {

	return formfielddatetime.GongIsStaged(stage)
}

func (formfieldfloat64 *FormFieldFloat64) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormFieldFloat64s[formfieldfloat64]

	return
}

func (stage *Stage) IsStagedFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) (ok bool) {

	return formfieldfloat64.GongIsStaged(stage)
}

func (formfieldint *FormFieldInt) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormFieldInts[formfieldint]

	return
}

func (stage *Stage) IsStagedFormFieldInt(formfieldint *FormFieldInt) (ok bool) {

	return formfieldint.GongIsStaged(stage)
}

func (formfieldselect *FormFieldSelect) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormFieldSelects[formfieldselect]

	return
}

func (stage *Stage) IsStagedFormFieldSelect(formfieldselect *FormFieldSelect) (ok bool) {

	return formfieldselect.GongIsStaged(stage)
}

func (formfieldstring *FormFieldString) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormFieldStrings[formfieldstring]

	return
}

func (stage *Stage) IsStagedFormFieldString(formfieldstring *FormFieldString) (ok bool) {

	return formfieldstring.GongIsStaged(stage)
}

func (formfieldtime *FormFieldTime) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormFieldTimes[formfieldtime]

	return
}

func (stage *Stage) IsStagedFormFieldTime(formfieldtime *FormFieldTime) (ok bool) {

	return formfieldtime.GongIsStaged(stage)
}

func (formgroup *FormGroup) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormGroups[formgroup]

	return
}

func (stage *Stage) IsStagedFormGroup(formgroup *FormGroup) (ok bool) {

	return formgroup.GongIsStaged(stage)
}

func (formsortassocbutton *FormSortAssocButton) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.FormSortAssocButtons[formsortassocbutton]

	return
}

func (stage *Stage) IsStagedFormSortAssocButton(formsortassocbutton *FormSortAssocButton) (ok bool) {

	return formsortassocbutton.GongIsStaged(stage)
}

func (option *Option) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Options[option]

	return
}

func (stage *Stage) IsStagedOption(option *Option) (ok bool) {

	return option.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (checkbox *CheckBox) GongStageBranch(stage *Stage) {
	stage.StageBranchCheckBox(checkbox)
}

func (stage *Stage) StageBranchCheckBox(checkbox *CheckBox) {

	// check if instance is already staged
	if stage.IsStaged(checkbox) {
		return
	}

	checkbox.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formdiv *FormDiv) GongStageBranch(stage *Stage) {
	stage.StageBranchFormDiv(formdiv)
}

func (stage *Stage) StageBranchFormDiv(formdiv *FormDiv) {

	// check if instance is already staged
	if stage.IsStaged(formdiv) {
		return
	}

	formdiv.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formdiv.FormEditAssocButton != nil {
		stage.StageBranch(formdiv.FormEditAssocButton)
	}
	if formdiv.FormSortAssocButton != nil {
		stage.StageBranch(formdiv.FormSortAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formfield := range formdiv.FormFields {
		stage.StageBranch(_formfield)
	}
	for _, _checkbox := range formdiv.CheckBoxs {
		stage.StageBranch(_checkbox)
	}

}

func (formeditassocbutton *FormEditAssocButton) GongStageBranch(stage *Stage) {
	stage.StageBranchFormEditAssocButton(formeditassocbutton)
}

func (stage *Stage) StageBranchFormEditAssocButton(formeditassocbutton *FormEditAssocButton) {

	// check if instance is already staged
	if stage.IsStaged(formeditassocbutton) {
		return
	}

	formeditassocbutton.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfield *FormField) GongStageBranch(stage *Stage) {
	stage.StageBranchFormField(formfield)
}

func (stage *Stage) StageBranchFormField(formfield *FormField) {

	// check if instance is already staged
	if stage.IsStaged(formfield) {
		return
	}

	formfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formfield.FormFieldString != nil {
		stage.StageBranch(formfield.FormFieldString)
	}
	if formfield.FormFieldFloat64 != nil {
		stage.StageBranch(formfield.FormFieldFloat64)
	}
	if formfield.FormFieldInt != nil {
		stage.StageBranch(formfield.FormFieldInt)
	}
	if formfield.FormFieldDate != nil {
		stage.StageBranch(formfield.FormFieldDate)
	}
	if formfield.FormFieldTime != nil {
		stage.StageBranch(formfield.FormFieldTime)
	}
	if formfield.FormFieldDateTime != nil {
		stage.StageBranch(formfield.FormFieldDateTime)
	}
	if formfield.FormFieldSelect != nil {
		stage.StageBranch(formfield.FormFieldSelect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfielddate *FormFieldDate) GongStageBranch(stage *Stage) {
	stage.StageBranchFormFieldDate(formfielddate)
}

func (stage *Stage) StageBranchFormFieldDate(formfielddate *FormFieldDate) {

	// check if instance is already staged
	if stage.IsStaged(formfielddate) {
		return
	}

	formfielddate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfielddatetime *FormFieldDateTime) GongStageBranch(stage *Stage) {
	stage.StageBranchFormFieldDateTime(formfielddatetime)
}

func (stage *Stage) StageBranchFormFieldDateTime(formfielddatetime *FormFieldDateTime) {

	// check if instance is already staged
	if stage.IsStaged(formfielddatetime) {
		return
	}

	formfielddatetime.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldfloat64 *FormFieldFloat64) GongStageBranch(stage *Stage) {
	stage.StageBranchFormFieldFloat64(formfieldfloat64)
}

func (stage *Stage) StageBranchFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) {

	// check if instance is already staged
	if stage.IsStaged(formfieldfloat64) {
		return
	}

	formfieldfloat64.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldint *FormFieldInt) GongStageBranch(stage *Stage) {
	stage.StageBranchFormFieldInt(formfieldint)
}

func (stage *Stage) StageBranchFormFieldInt(formfieldint *FormFieldInt) {

	// check if instance is already staged
	if stage.IsStaged(formfieldint) {
		return
	}

	formfieldint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldselect *FormFieldSelect) GongStageBranch(stage *Stage) {
	stage.StageBranchFormFieldSelect(formfieldselect)
}

func (stage *Stage) StageBranchFormFieldSelect(formfieldselect *FormFieldSelect) {

	// check if instance is already staged
	if stage.IsStaged(formfieldselect) {
		return
	}

	formfieldselect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formfieldselect.Value != nil {
		stage.StageBranch(formfieldselect.Value)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _option := range formfieldselect.Options {
		stage.StageBranch(_option)
	}

}

func (formfieldstring *FormFieldString) GongStageBranch(stage *Stage) {
	stage.StageBranchFormFieldString(formfieldstring)
}

func (stage *Stage) StageBranchFormFieldString(formfieldstring *FormFieldString) {

	// check if instance is already staged
	if stage.IsStaged(formfieldstring) {
		return
	}

	formfieldstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldtime *FormFieldTime) GongStageBranch(stage *Stage) {
	stage.StageBranchFormFieldTime(formfieldtime)
}

func (stage *Stage) StageBranchFormFieldTime(formfieldtime *FormFieldTime) {

	// check if instance is already staged
	if stage.IsStaged(formfieldtime) {
		return
	}

	formfieldtime.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formgroup *FormGroup) GongStageBranch(stage *Stage) {
	stage.StageBranchFormGroup(formgroup)
}

func (stage *Stage) StageBranchFormGroup(formgroup *FormGroup) {

	// check if instance is already staged
	if stage.IsStaged(formgroup) {
		return
	}

	formgroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formdiv := range formgroup.FormDivs {
		stage.StageBranch(_formdiv)
	}

}

func (formsortassocbutton *FormSortAssocButton) GongStageBranch(stage *Stage) {
	stage.StageBranchFormSortAssocButton(formsortassocbutton)
}

func (stage *Stage) StageBranchFormSortAssocButton(formsortassocbutton *FormSortAssocButton) {

	// check if instance is already staged
	if stage.IsStaged(formsortassocbutton) {
		return
	}

	formsortassocbutton.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formsortassocbutton.FormEditAssocButton != nil {
		stage.StageBranch(formsortassocbutton.FormEditAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (option *Option) GongStageBranch(stage *Stage) {
	stage.StageBranchOption(option)
}

func (stage *Stage) StageBranchOption(option *Option) {

	// check if instance is already staged
	if stage.IsStaged(option) {
		return
	}

	option.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
	case *CheckBox:
		toT := GongCopyBranchCheckBox(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormDiv:
		toT := GongCopyBranchFormDiv(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormEditAssocButton:
		toT := GongCopyBranchFormEditAssocButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormField:
		toT := GongCopyBranchFormField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldDate:
		toT := GongCopyBranchFormFieldDate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldDateTime:
		toT := GongCopyBranchFormFieldDateTime(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldFloat64:
		toT := GongCopyBranchFormFieldFloat64(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldInt:
		toT := GongCopyBranchFormFieldInt(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldSelect:
		toT := GongCopyBranchFormFieldSelect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldString:
		toT := GongCopyBranchFormFieldString(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormFieldTime:
		toT := GongCopyBranchFormFieldTime(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormGroup:
		toT := GongCopyBranchFormGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FormSortAssocButton:
		toT := GongCopyBranchFormSortAssocButton(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Option:
		toT := GongCopyBranchOption(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCheckBox(mapOrigCopy map[any]any, checkboxFrom *CheckBox) (checkboxTo *CheckBox) {

	// checkboxFrom has already been copied
	if _checkboxTo, ok := mapOrigCopy[checkboxFrom]; ok {
		checkboxTo = _checkboxTo.(*CheckBox)
		return
	}

	checkboxTo = new(CheckBox)
	mapOrigCopy[checkboxFrom] = checkboxTo
	checkboxFrom.GongCopyBasicFields(checkboxTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormDiv(mapOrigCopy map[any]any, formdivFrom *FormDiv) (formdivTo *FormDiv) {

	// formdivFrom has already been copied
	if _formdivTo, ok := mapOrigCopy[formdivFrom]; ok {
		formdivTo = _formdivTo.(*FormDiv)
		return
	}

	formdivTo = new(FormDiv)
	mapOrigCopy[formdivFrom] = formdivTo
	formdivFrom.GongCopyBasicFields(formdivTo)

	//insertion point for the staging of instances referenced by pointers
	if formdivFrom.FormEditAssocButton != nil {
		formdivTo.FormEditAssocButton = GongCopyBranchFormEditAssocButton(mapOrigCopy, formdivFrom.FormEditAssocButton)
	}
	if formdivFrom.FormSortAssocButton != nil {
		formdivTo.FormSortAssocButton = GongCopyBranchFormSortAssocButton(mapOrigCopy, formdivFrom.FormSortAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formfield := range formdivFrom.FormFields {
		formdivTo.FormFields = append(formdivTo.FormFields, GongCopyBranchFormField(mapOrigCopy, _formfield))
	}
	for _, _checkbox := range formdivFrom.CheckBoxs {
		formdivTo.CheckBoxs = append(formdivTo.CheckBoxs, GongCopyBranchCheckBox(mapOrigCopy, _checkbox))
	}

	return
}

func GongCopyBranchFormEditAssocButton(mapOrigCopy map[any]any, formeditassocbuttonFrom *FormEditAssocButton) (formeditassocbuttonTo *FormEditAssocButton) {

	// formeditassocbuttonFrom has already been copied
	if _formeditassocbuttonTo, ok := mapOrigCopy[formeditassocbuttonFrom]; ok {
		formeditassocbuttonTo = _formeditassocbuttonTo.(*FormEditAssocButton)
		return
	}

	formeditassocbuttonTo = new(FormEditAssocButton)
	mapOrigCopy[formeditassocbuttonFrom] = formeditassocbuttonTo
	formeditassocbuttonFrom.GongCopyBasicFields(formeditassocbuttonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormField(mapOrigCopy map[any]any, formfieldFrom *FormField) (formfieldTo *FormField) {

	// formfieldFrom has already been copied
	if _formfieldTo, ok := mapOrigCopy[formfieldFrom]; ok {
		formfieldTo = _formfieldTo.(*FormField)
		return
	}

	formfieldTo = new(FormField)
	mapOrigCopy[formfieldFrom] = formfieldTo
	formfieldFrom.GongCopyBasicFields(formfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if formfieldFrom.FormFieldString != nil {
		formfieldTo.FormFieldString = GongCopyBranchFormFieldString(mapOrigCopy, formfieldFrom.FormFieldString)
	}
	if formfieldFrom.FormFieldFloat64 != nil {
		formfieldTo.FormFieldFloat64 = GongCopyBranchFormFieldFloat64(mapOrigCopy, formfieldFrom.FormFieldFloat64)
	}
	if formfieldFrom.FormFieldInt != nil {
		formfieldTo.FormFieldInt = GongCopyBranchFormFieldInt(mapOrigCopy, formfieldFrom.FormFieldInt)
	}
	if formfieldFrom.FormFieldDate != nil {
		formfieldTo.FormFieldDate = GongCopyBranchFormFieldDate(mapOrigCopy, formfieldFrom.FormFieldDate)
	}
	if formfieldFrom.FormFieldTime != nil {
		formfieldTo.FormFieldTime = GongCopyBranchFormFieldTime(mapOrigCopy, formfieldFrom.FormFieldTime)
	}
	if formfieldFrom.FormFieldDateTime != nil {
		formfieldTo.FormFieldDateTime = GongCopyBranchFormFieldDateTime(mapOrigCopy, formfieldFrom.FormFieldDateTime)
	}
	if formfieldFrom.FormFieldSelect != nil {
		formfieldTo.FormFieldSelect = GongCopyBranchFormFieldSelect(mapOrigCopy, formfieldFrom.FormFieldSelect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldDate(mapOrigCopy map[any]any, formfielddateFrom *FormFieldDate) (formfielddateTo *FormFieldDate) {

	// formfielddateFrom has already been copied
	if _formfielddateTo, ok := mapOrigCopy[formfielddateFrom]; ok {
		formfielddateTo = _formfielddateTo.(*FormFieldDate)
		return
	}

	formfielddateTo = new(FormFieldDate)
	mapOrigCopy[formfielddateFrom] = formfielddateTo
	formfielddateFrom.GongCopyBasicFields(formfielddateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldDateTime(mapOrigCopy map[any]any, formfielddatetimeFrom *FormFieldDateTime) (formfielddatetimeTo *FormFieldDateTime) {

	// formfielddatetimeFrom has already been copied
	if _formfielddatetimeTo, ok := mapOrigCopy[formfielddatetimeFrom]; ok {
		formfielddatetimeTo = _formfielddatetimeTo.(*FormFieldDateTime)
		return
	}

	formfielddatetimeTo = new(FormFieldDateTime)
	mapOrigCopy[formfielddatetimeFrom] = formfielddatetimeTo
	formfielddatetimeFrom.GongCopyBasicFields(formfielddatetimeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldFloat64(mapOrigCopy map[any]any, formfieldfloat64From *FormFieldFloat64) (formfieldfloat64To *FormFieldFloat64) {

	// formfieldfloat64From has already been copied
	if _formfieldfloat64To, ok := mapOrigCopy[formfieldfloat64From]; ok {
		formfieldfloat64To = _formfieldfloat64To.(*FormFieldFloat64)
		return
	}

	formfieldfloat64To = new(FormFieldFloat64)
	mapOrigCopy[formfieldfloat64From] = formfieldfloat64To
	formfieldfloat64From.GongCopyBasicFields(formfieldfloat64To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldInt(mapOrigCopy map[any]any, formfieldintFrom *FormFieldInt) (formfieldintTo *FormFieldInt) {

	// formfieldintFrom has already been copied
	if _formfieldintTo, ok := mapOrigCopy[formfieldintFrom]; ok {
		formfieldintTo = _formfieldintTo.(*FormFieldInt)
		return
	}

	formfieldintTo = new(FormFieldInt)
	mapOrigCopy[formfieldintFrom] = formfieldintTo
	formfieldintFrom.GongCopyBasicFields(formfieldintTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldSelect(mapOrigCopy map[any]any, formfieldselectFrom *FormFieldSelect) (formfieldselectTo *FormFieldSelect) {

	// formfieldselectFrom has already been copied
	if _formfieldselectTo, ok := mapOrigCopy[formfieldselectFrom]; ok {
		formfieldselectTo = _formfieldselectTo.(*FormFieldSelect)
		return
	}

	formfieldselectTo = new(FormFieldSelect)
	mapOrigCopy[formfieldselectFrom] = formfieldselectTo
	formfieldselectFrom.GongCopyBasicFields(formfieldselectTo)

	//insertion point for the staging of instances referenced by pointers
	if formfieldselectFrom.Value != nil {
		formfieldselectTo.Value = GongCopyBranchOption(mapOrigCopy, formfieldselectFrom.Value)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _option := range formfieldselectFrom.Options {
		formfieldselectTo.Options = append(formfieldselectTo.Options, GongCopyBranchOption(mapOrigCopy, _option))
	}

	return
}

func GongCopyBranchFormFieldString(mapOrigCopy map[any]any, formfieldstringFrom *FormFieldString) (formfieldstringTo *FormFieldString) {

	// formfieldstringFrom has already been copied
	if _formfieldstringTo, ok := mapOrigCopy[formfieldstringFrom]; ok {
		formfieldstringTo = _formfieldstringTo.(*FormFieldString)
		return
	}

	formfieldstringTo = new(FormFieldString)
	mapOrigCopy[formfieldstringFrom] = formfieldstringTo
	formfieldstringFrom.GongCopyBasicFields(formfieldstringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldTime(mapOrigCopy map[any]any, formfieldtimeFrom *FormFieldTime) (formfieldtimeTo *FormFieldTime) {

	// formfieldtimeFrom has already been copied
	if _formfieldtimeTo, ok := mapOrigCopy[formfieldtimeFrom]; ok {
		formfieldtimeTo = _formfieldtimeTo.(*FormFieldTime)
		return
	}

	formfieldtimeTo = new(FormFieldTime)
	mapOrigCopy[formfieldtimeFrom] = formfieldtimeTo
	formfieldtimeFrom.GongCopyBasicFields(formfieldtimeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormGroup(mapOrigCopy map[any]any, formgroupFrom *FormGroup) (formgroupTo *FormGroup) {

	// formgroupFrom has already been copied
	if _formgroupTo, ok := mapOrigCopy[formgroupFrom]; ok {
		formgroupTo = _formgroupTo.(*FormGroup)
		return
	}

	formgroupTo = new(FormGroup)
	mapOrigCopy[formgroupFrom] = formgroupTo
	formgroupFrom.GongCopyBasicFields(formgroupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formdiv := range formgroupFrom.FormDivs {
		formgroupTo.FormDivs = append(formgroupTo.FormDivs, GongCopyBranchFormDiv(mapOrigCopy, _formdiv))
	}

	return
}

func GongCopyBranchFormSortAssocButton(mapOrigCopy map[any]any, formsortassocbuttonFrom *FormSortAssocButton) (formsortassocbuttonTo *FormSortAssocButton) {

	// formsortassocbuttonFrom has already been copied
	if _formsortassocbuttonTo, ok := mapOrigCopy[formsortassocbuttonFrom]; ok {
		formsortassocbuttonTo = _formsortassocbuttonTo.(*FormSortAssocButton)
		return
	}

	formsortassocbuttonTo = new(FormSortAssocButton)
	mapOrigCopy[formsortassocbuttonFrom] = formsortassocbuttonTo
	formsortassocbuttonFrom.GongCopyBasicFields(formsortassocbuttonTo)

	//insertion point for the staging of instances referenced by pointers
	if formsortassocbuttonFrom.FormEditAssocButton != nil {
		formsortassocbuttonTo.FormEditAssocButton = GongCopyBranchFormEditAssocButton(mapOrigCopy, formsortassocbuttonFrom.FormEditAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOption(mapOrigCopy map[any]any, optionFrom *Option) (optionTo *Option) {

	// optionFrom has already been copied
	if _optionTo, ok := mapOrigCopy[optionFrom]; ok {
		optionTo = _optionTo.(*Option)
		return
	}

	optionTo = new(Option)
	mapOrigCopy[optionFrom] = optionTo
	optionFrom.GongCopyBasicFields(optionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (checkbox *CheckBox) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCheckBox(checkbox)
}

func (stage *Stage) UnstageBranchCheckBox(checkbox *CheckBox) {

	// check if instance is already staged
	if !stage.IsStaged(checkbox) {
		return
	}

	checkbox.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formdiv *FormDiv) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormDiv(formdiv)
}

func (stage *Stage) UnstageBranchFormDiv(formdiv *FormDiv) {

	// check if instance is already staged
	if !stage.IsStaged(formdiv) {
		return
	}

	formdiv.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formdiv.FormEditAssocButton != nil {
		stage.UnstageBranch(formdiv.FormEditAssocButton)
	}
	if formdiv.FormSortAssocButton != nil {
		stage.UnstageBranch(formdiv.FormSortAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formfield := range formdiv.FormFields {
		stage.UnstageBranch(_formfield)
	}
	for _, _checkbox := range formdiv.CheckBoxs {
		stage.UnstageBranch(_checkbox)
	}

}

func (formeditassocbutton *FormEditAssocButton) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormEditAssocButton(formeditassocbutton)
}

func (stage *Stage) UnstageBranchFormEditAssocButton(formeditassocbutton *FormEditAssocButton) {

	// check if instance is already staged
	if !stage.IsStaged(formeditassocbutton) {
		return
	}

	formeditassocbutton.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfield *FormField) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormField(formfield)
}

func (stage *Stage) UnstageBranchFormField(formfield *FormField) {

	// check if instance is already staged
	if !stage.IsStaged(formfield) {
		return
	}

	formfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formfield.FormFieldString != nil {
		stage.UnstageBranch(formfield.FormFieldString)
	}
	if formfield.FormFieldFloat64 != nil {
		stage.UnstageBranch(formfield.FormFieldFloat64)
	}
	if formfield.FormFieldInt != nil {
		stage.UnstageBranch(formfield.FormFieldInt)
	}
	if formfield.FormFieldDate != nil {
		stage.UnstageBranch(formfield.FormFieldDate)
	}
	if formfield.FormFieldTime != nil {
		stage.UnstageBranch(formfield.FormFieldTime)
	}
	if formfield.FormFieldDateTime != nil {
		stage.UnstageBranch(formfield.FormFieldDateTime)
	}
	if formfield.FormFieldSelect != nil {
		stage.UnstageBranch(formfield.FormFieldSelect)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfielddate *FormFieldDate) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormFieldDate(formfielddate)
}

func (stage *Stage) UnstageBranchFormFieldDate(formfielddate *FormFieldDate) {

	// check if instance is already staged
	if !stage.IsStaged(formfielddate) {
		return
	}

	formfielddate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfielddatetime *FormFieldDateTime) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormFieldDateTime(formfielddatetime)
}

func (stage *Stage) UnstageBranchFormFieldDateTime(formfielddatetime *FormFieldDateTime) {

	// check if instance is already staged
	if !stage.IsStaged(formfielddatetime) {
		return
	}

	formfielddatetime.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldfloat64 *FormFieldFloat64) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormFieldFloat64(formfieldfloat64)
}

func (stage *Stage) UnstageBranchFormFieldFloat64(formfieldfloat64 *FormFieldFloat64) {

	// check if instance is already staged
	if !stage.IsStaged(formfieldfloat64) {
		return
	}

	formfieldfloat64.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldint *FormFieldInt) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormFieldInt(formfieldint)
}

func (stage *Stage) UnstageBranchFormFieldInt(formfieldint *FormFieldInt) {

	// check if instance is already staged
	if !stage.IsStaged(formfieldint) {
		return
	}

	formfieldint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldselect *FormFieldSelect) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormFieldSelect(formfieldselect)
}

func (stage *Stage) UnstageBranchFormFieldSelect(formfieldselect *FormFieldSelect) {

	// check if instance is already staged
	if !stage.IsStaged(formfieldselect) {
		return
	}

	formfieldselect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formfieldselect.Value != nil {
		stage.UnstageBranch(formfieldselect.Value)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _option := range formfieldselect.Options {
		stage.UnstageBranch(_option)
	}

}

func (formfieldstring *FormFieldString) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormFieldString(formfieldstring)
}

func (stage *Stage) UnstageBranchFormFieldString(formfieldstring *FormFieldString) {

	// check if instance is already staged
	if !stage.IsStaged(formfieldstring) {
		return
	}

	formfieldstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldtime *FormFieldTime) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormFieldTime(formfieldtime)
}

func (stage *Stage) UnstageBranchFormFieldTime(formfieldtime *FormFieldTime) {

	// check if instance is already staged
	if !stage.IsStaged(formfieldtime) {
		return
	}

	formfieldtime.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formgroup *FormGroup) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormGroup(formgroup)
}

func (stage *Stage) UnstageBranchFormGroup(formgroup *FormGroup) {

	// check if instance is already staged
	if !stage.IsStaged(formgroup) {
		return
	}

	formgroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formdiv := range formgroup.FormDivs {
		stage.UnstageBranch(_formdiv)
	}

}

func (formsortassocbutton *FormSortAssocButton) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchFormSortAssocButton(formsortassocbutton)
}

func (stage *Stage) UnstageBranchFormSortAssocButton(formsortassocbutton *FormSortAssocButton) {

	// check if instance is already staged
	if !stage.IsStaged(formsortassocbutton) {
		return
	}

	formsortassocbutton.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if formsortassocbutton.FormEditAssocButton != nil {
		stage.UnstageBranch(formsortassocbutton.FormEditAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (option *Option) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchOption(option)
}

func (stage *Stage) UnstageBranchOption(option *Option) {

	// check if instance is already staged
	if !stage.IsStaged(option) {
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
		ops := stage.Diff(
			formdiv,
			"FormFields",
			len(formdivOther.FormFields),
			len(formdiv.FormFields),
			func(i, j int) bool {
				return formdivOther.FormFields[i] == formdiv.FormFields[j]
			},
			func(j int) string {
				return formdiv.FormFields[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			formdiv,
			"CheckBoxs",
			len(formdivOther.CheckBoxs),
			len(formdiv.CheckBoxs),
			func(i, j int) bool {
				return formdivOther.CheckBoxs[i] == formdiv.CheckBoxs[j]
			},
			func(j int) string {
				return formdiv.CheckBoxs[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			formfieldselect,
			"Options",
			len(formfieldselectOther.Options),
			len(formfieldselect.Options),
			func(i, j int) bool {
				return formfieldselectOther.Options[i] == formfieldselect.Options[j]
			},
			func(j int) string {
				return formfieldselect.Options[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			formgroup,
			"FormDivs",
			len(formgroupOther.FormDivs),
			len(formgroup.FormDivs),
			func(i, j int) bool {
				return formgroupOther.FormDivs[i] == formgroup.FormDivs[j]
			},
			func(j int) string {
				return formgroup.FormDivs[j].GongGetIdentifier(stage)
			},
		)
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
