// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (checkbox *CheckBox) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CheckBoxs[checkbox]
	return ok
}

func (formdiv *FormDiv) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormDivs[formdiv]
	return ok
}

func (formeditassocbutton *FormEditAssocButton) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormEditAssocButtons[formeditassocbutton]
	return ok
}

func (formfield *FormField) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormFields[formfield]
	return ok
}

func (formfielddate *FormFieldDate) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormFieldDates[formfielddate]
	return ok
}

func (formfielddatetime *FormFieldDateTime) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormFieldDateTimes[formfielddatetime]
	return ok
}

func (formfieldfloat64 *FormFieldFloat64) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormFieldFloat64s[formfieldfloat64]
	return ok
}

func (formfieldint *FormFieldInt) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormFieldInts[formfieldint]
	return ok
}

func (formfieldselect *FormFieldSelect) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormFieldSelects[formfieldselect]
	return ok
}

func (formfieldstring *FormFieldString) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormFieldStrings[formfieldstring]
	return ok
}

func (formfieldtime *FormFieldTime) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormFieldTimes[formfieldtime]
	return ok
}

func (formgroup *FormGroup) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormGroups[formgroup]
	return ok
}

func (formsortassocbutton *FormSortAssocButton) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FormSortAssocButtons[formsortassocbutton]
	return ok
}

func (option *Option) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Options[option]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (checkbox *CheckBox) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(checkbox) {
		return
	}

	checkbox.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formdiv *FormDiv) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(formeditassocbutton) {
		return
	}

	formeditassocbutton.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfield *FormField) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(formfielddate) {
		return
	}

	formfielddate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfielddatetime *FormFieldDateTime) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formfielddatetime) {
		return
	}

	formfielddatetime.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldfloat64 *FormFieldFloat64) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formfieldfloat64) {
		return
	}

	formfieldfloat64.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldint *FormFieldInt) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formfieldint) {
		return
	}

	formfieldint.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldselect *FormFieldSelect) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(formfieldstring) {
		return
	}

	formfieldstring.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldtime *FormFieldTime) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(formfieldtime) {
		return
	}

	formfieldtime.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formgroup *FormGroup) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	checkboxTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, checkboxFrom)
	if alreadyCopied {
		return
	}
	checkboxFrom.GongCopyBasicFields(checkboxTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormDiv(mapOrigCopy map[any]any, formdivFrom *FormDiv) (formdivTo *FormDiv) {
	var alreadyCopied bool
	formdivTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formdivFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	formeditassocbuttonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formeditassocbuttonFrom)
	if alreadyCopied {
		return
	}
	formeditassocbuttonFrom.GongCopyBasicFields(formeditassocbuttonTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormField(mapOrigCopy map[any]any, formfieldFrom *FormField) (formfieldTo *FormField) {
	var alreadyCopied bool
	formfieldTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formfieldFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	formfielddateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formfielddateFrom)
	if alreadyCopied {
		return
	}
	formfielddateFrom.GongCopyBasicFields(formfielddateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldDateTime(mapOrigCopy map[any]any, formfielddatetimeFrom *FormFieldDateTime) (formfielddatetimeTo *FormFieldDateTime) {
	var alreadyCopied bool
	formfielddatetimeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formfielddatetimeFrom)
	if alreadyCopied {
		return
	}
	formfielddatetimeFrom.GongCopyBasicFields(formfielddatetimeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldFloat64(mapOrigCopy map[any]any, formfieldfloat64From *FormFieldFloat64) (formfieldfloat64To *FormFieldFloat64) {
	var alreadyCopied bool
	formfieldfloat64To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formfieldfloat64From)
	if alreadyCopied {
		return
	}
	formfieldfloat64From.GongCopyBasicFields(formfieldfloat64To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldInt(mapOrigCopy map[any]any, formfieldintFrom *FormFieldInt) (formfieldintTo *FormFieldInt) {
	var alreadyCopied bool
	formfieldintTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formfieldintFrom)
	if alreadyCopied {
		return
	}
	formfieldintFrom.GongCopyBasicFields(formfieldintTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldSelect(mapOrigCopy map[any]any, formfieldselectFrom *FormFieldSelect) (formfieldselectTo *FormFieldSelect) {
	var alreadyCopied bool
	formfieldselectTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formfieldselectFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	formfieldstringTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formfieldstringFrom)
	if alreadyCopied {
		return
	}
	formfieldstringFrom.GongCopyBasicFields(formfieldstringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormFieldTime(mapOrigCopy map[any]any, formfieldtimeFrom *FormFieldTime) (formfieldtimeTo *FormFieldTime) {
	var alreadyCopied bool
	formfieldtimeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formfieldtimeFrom)
	if alreadyCopied {
		return
	}
	formfieldtimeFrom.GongCopyBasicFields(formfieldtimeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFormGroup(mapOrigCopy map[any]any, formgroupFrom *FormGroup) (formgroupTo *FormGroup) {
	var alreadyCopied bool
	formgroupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formgroupFrom)
	if alreadyCopied {
		return
	}
	formgroupFrom.GongCopyBasicFields(formgroupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _formdiv := range formgroupFrom.FormDivs {
		formgroupTo.FormDivs = append(formgroupTo.FormDivs, GongCopyBranchFormDiv(mapOrigCopy, _formdiv))
	}

	return
}

func GongCopyBranchFormSortAssocButton(mapOrigCopy map[any]any, formsortassocbuttonFrom *FormSortAssocButton) (formsortassocbuttonTo *FormSortAssocButton) {
	var alreadyCopied bool
	formsortassocbuttonTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, formsortassocbuttonFrom)
	if alreadyCopied {
		return
	}
	formsortassocbuttonFrom.GongCopyBasicFields(formsortassocbuttonTo)

	//insertion point for the staging of instances referenced by pointers
	if formsortassocbuttonFrom.FormEditAssocButton != nil {
		formsortassocbuttonTo.FormEditAssocButton = GongCopyBranchFormEditAssocButton(mapOrigCopy, formsortassocbuttonFrom.FormEditAssocButton)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchOption(mapOrigCopy map[any]any, optionFrom *Option) (optionTo *Option) {
	var alreadyCopied bool
	optionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, optionFrom)
	if alreadyCopied {
		return
	}
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

// insertion point for unstage branch per struct
func (checkbox *CheckBox) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(checkbox) {
		return
	}

	checkbox.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formdiv *FormDiv) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(formeditassocbutton) {
		return
	}

	formeditassocbutton.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfield *FormField) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(formfielddate) {
		return
	}

	formfielddate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfielddatetime *FormFieldDateTime) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formfielddatetime) {
		return
	}

	formfielddatetime.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldfloat64 *FormFieldFloat64) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formfieldfloat64) {
		return
	}

	formfieldfloat64.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldint *FormFieldInt) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formfieldint) {
		return
	}

	formfieldint.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldselect *FormFieldSelect) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(formfieldstring) {
		return
	}

	formfieldstring.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formfieldtime *FormFieldTime) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(formfieldtime) {
		return
	}

	formfieldtime.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (formgroup *FormGroup) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructPointer(&reference.FormEditAssocButton, stage.FormEditAssocButtons_reference, instance.FormEditAssocButton)
	__gong__reconstructPointer(&reference.FormSortAssocButton, stage.FormSortAssocButtons_reference, instance.FormSortAssocButton)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.FormFields, stage.FormFields_reference, instance.FormFields)
	__gong__reconstructSliceOfPointersFromReferences(&reference.CheckBoxs, stage.CheckBoxs_reference, instance.CheckBoxs)
}

func (reference *FormEditAssocButton) GongReconstructPointersFromReferences(stage *Stage, instance *FormEditAssocButton) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FormField) GongReconstructPointersFromReferences(stage *Stage, instance *FormField) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.FormFieldString, stage.FormFieldStrings_reference, instance.FormFieldString)
	__gong__reconstructPointer(&reference.FormFieldFloat64, stage.FormFieldFloat64s_reference, instance.FormFieldFloat64)
	__gong__reconstructPointer(&reference.FormFieldInt, stage.FormFieldInts_reference, instance.FormFieldInt)
	__gong__reconstructPointer(&reference.FormFieldDate, stage.FormFieldDates_reference, instance.FormFieldDate)
	__gong__reconstructPointer(&reference.FormFieldTime, stage.FormFieldTimes_reference, instance.FormFieldTime)
	__gong__reconstructPointer(&reference.FormFieldDateTime, stage.FormFieldDateTimes_reference, instance.FormFieldDateTime)
	__gong__reconstructPointer(&reference.FormFieldSelect, stage.FormFieldSelects_reference, instance.FormFieldSelect)
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
	__gong__reconstructPointer(&reference.Value, stage.Options_reference, instance.Value)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Options, stage.Options_reference, instance.Options)
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
	__gong__reconstructSliceOfPointersFromReferences(&reference.FormDivs, stage.FormDivs_reference, instance.FormDivs)
}

func (reference *FormSortAssocButton) GongReconstructPointersFromReferences(stage *Stage, instance *FormSortAssocButton) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.FormEditAssocButton, stage.FormEditAssocButtons_reference, instance.FormEditAssocButton)
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
	__gong__reconstructPointerFromInstance(&reference.FormEditAssocButton, stage.FormEditAssocButtons_instance)
	__gong__reconstructPointerFromInstance(&reference.FormSortAssocButton, stage.FormSortAssocButtons_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.FormFields, stage.FormFields_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.CheckBoxs, stage.CheckBoxs_instance)
}

func (reference *FormEditAssocButton) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FormField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.FormFieldString, stage.FormFieldStrings_instance)
	__gong__reconstructPointerFromInstance(&reference.FormFieldFloat64, stage.FormFieldFloat64s_instance)
	__gong__reconstructPointerFromInstance(&reference.FormFieldInt, stage.FormFieldInts_instance)
	__gong__reconstructPointerFromInstance(&reference.FormFieldDate, stage.FormFieldDates_instance)
	__gong__reconstructPointerFromInstance(&reference.FormFieldTime, stage.FormFieldTimes_instance)
	__gong__reconstructPointerFromInstance(&reference.FormFieldDateTime, stage.FormFieldDateTimes_instance)
	__gong__reconstructPointerFromInstance(&reference.FormFieldSelect, stage.FormFieldSelects_instance)
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
	__gong__reconstructPointerFromInstance(&reference.Value, stage.Options_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Options, stage.Options_instance)
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.FormDivs, stage.FormDivs_instance)
}

func (reference *FormSortAssocButton) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.FormEditAssocButton, stage.FormEditAssocButtons_instance)
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
	if ops := __gong__diffSliceOfPointers(stage, formdiv, "FormFields", formdivOther.FormFields, formdiv.FormFields); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, formdiv, "CheckBoxs", formdivOther.CheckBoxs, formdiv.CheckBoxs); ops != "" {
		diffs = append(diffs, ops)
	}
	if formdiv.FormEditAssocButton != formdivOther.FormEditAssocButton {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "FormEditAssocButton"))
	}
	if formdiv.FormSortAssocButton != formdivOther.FormSortAssocButton {
		diffs = append(diffs, formdiv.GongMarshallField(stage, "FormSortAssocButton"))
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
	if formfield.FormFieldString != formfieldOther.FormFieldString {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldString"))
	}
	if formfield.FormFieldFloat64 != formfieldOther.FormFieldFloat64 {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldFloat64"))
	}
	if formfield.FormFieldInt != formfieldOther.FormFieldInt {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldInt"))
	}
	if formfield.FormFieldDate != formfieldOther.FormFieldDate {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldDate"))
	}
	if formfield.FormFieldTime != formfieldOther.FormFieldTime {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldTime"))
	}
	if formfield.FormFieldDateTime != formfieldOther.FormFieldDateTime {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldDateTime"))
	}
	if formfield.FormFieldSelect != formfieldOther.FormFieldSelect {
		diffs = append(diffs, formfield.GongMarshallField(stage, "FormFieldSelect"))
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
	if formfieldselect.Value != formfieldselectOther.Value {
		diffs = append(diffs, formfieldselect.GongMarshallField(stage, "Value"))
	}
	if ops := __gong__diffSliceOfPointers(stage, formfieldselect, "Options", formfieldselectOther.Options, formfieldselect.Options); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, formgroup, "FormDivs", formgroupOther.FormDivs, formgroup.FormDivs); ops != "" {
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
	if formsortassocbutton.FormEditAssocButton != formsortassocbuttonOther.FormEditAssocButton {
		diffs = append(diffs, formsortassocbutton.GongMarshallField(stage, "FormEditAssocButton"))
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

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
