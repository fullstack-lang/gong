// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct CheckBox
	// insertion point per field

	// Compute reverse map for named struct FormDiv
	// insertion point per field
	stage.FormDiv_FormFields_reverseMap = make(map[*FormField]*FormDiv)
	for formdiv := range stage.FormDivs {
		_ = formdiv
		for _, _formfield := range formdiv.FormFields {
			stage.FormDiv_FormFields_reverseMap[_formfield] = formdiv
		}
	}
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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.CheckBoxs {
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

	return
}

// insertion point per named struct
func (checkbox *CheckBox) GongCopy() GongstructIF {
	newInstance := new(CheckBox)
	checkbox.CopyBasicFields(newInstance)
	return newInstance
}

func (formdiv *FormDiv) GongCopy() GongstructIF {
	newInstance := new(FormDiv)
	formdiv.CopyBasicFields(newInstance)
	return newInstance
}

func (formeditassocbutton *FormEditAssocButton) GongCopy() GongstructIF {
	newInstance := new(FormEditAssocButton)
	formeditassocbutton.CopyBasicFields(newInstance)
	return newInstance
}

func (formfield *FormField) GongCopy() GongstructIF {
	newInstance := new(FormField)
	formfield.CopyBasicFields(newInstance)
	return newInstance
}

func (formfielddate *FormFieldDate) GongCopy() GongstructIF {
	newInstance := new(FormFieldDate)
	formfielddate.CopyBasicFields(newInstance)
	return newInstance
}

func (formfielddatetime *FormFieldDateTime) GongCopy() GongstructIF {
	newInstance := new(FormFieldDateTime)
	formfielddatetime.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldfloat64 *FormFieldFloat64) GongCopy() GongstructIF {
	newInstance := new(FormFieldFloat64)
	formfieldfloat64.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldint *FormFieldInt) GongCopy() GongstructIF {
	newInstance := new(FormFieldInt)
	formfieldint.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldselect *FormFieldSelect) GongCopy() GongstructIF {
	newInstance := new(FormFieldSelect)
	formfieldselect.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldstring *FormFieldString) GongCopy() GongstructIF {
	newInstance := new(FormFieldString)
	formfieldstring.CopyBasicFields(newInstance)
	return newInstance
}

func (formfieldtime *FormFieldTime) GongCopy() GongstructIF {
	newInstance := new(FormFieldTime)
	formfieldtime.CopyBasicFields(newInstance)
	return newInstance
}

func (formgroup *FormGroup) GongCopy() GongstructIF {
	newInstance := new(FormGroup)
	formgroup.CopyBasicFields(newInstance)
	return newInstance
}

func (formsortassocbutton *FormSortAssocButton) GongCopy() GongstructIF {
	newInstance := new(FormSortAssocButton)
	formsortassocbutton.CopyBasicFields(newInstance)
	return newInstance
}

func (option *Option) GongCopy() GongstructIF {
	newInstance := new(Option)
	option.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (checkbox *CheckBox) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(checkbox).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(checkbox), uint64(GetOrderPointerGongstruct(stage, checkbox)))
	return
}

func (formdiv *FormDiv) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formdiv).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formdiv), uint64(GetOrderPointerGongstruct(stage, formdiv)))
	return
}

func (formeditassocbutton *FormEditAssocButton) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formeditassocbutton).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formeditassocbutton), uint64(GetOrderPointerGongstruct(stage, formeditassocbutton)))
	return
}

func (formfield *FormField) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfield).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formfield), uint64(GetOrderPointerGongstruct(stage, formfield)))
	return
}

func (formfielddate *FormFieldDate) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfielddate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formfielddate), uint64(GetOrderPointerGongstruct(stage, formfielddate)))
	return
}

func (formfielddatetime *FormFieldDateTime) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfielddatetime).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formfielddatetime), uint64(GetOrderPointerGongstruct(stage, formfielddatetime)))
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldfloat64).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formfieldfloat64), uint64(GetOrderPointerGongstruct(stage, formfieldfloat64)))
	return
}

func (formfieldint *FormFieldInt) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldint).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formfieldint), uint64(GetOrderPointerGongstruct(stage, formfieldint)))
	return
}

func (formfieldselect *FormFieldSelect) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldselect).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formfieldselect), uint64(GetOrderPointerGongstruct(stage, formfieldselect)))
	return
}

func (formfieldstring *FormFieldString) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldstring).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formfieldstring), uint64(GetOrderPointerGongstruct(stage, formfieldstring)))
	return
}

func (formfieldtime *FormFieldTime) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldtime).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formfieldtime), uint64(GetOrderPointerGongstruct(stage, formfieldtime)))
	return
}

func (formgroup *FormGroup) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formgroup).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formgroup), uint64(GetOrderPointerGongstruct(stage, formgroup)))
	return
}

func (formsortassocbutton *FormSortAssocButton) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formsortassocbutton).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(formsortassocbutton), uint64(GetOrderPointerGongstruct(stage, formsortassocbutton)))
	return
}

func (option *Option) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(option).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(option), uint64(GetOrderPointerGongstruct(stage, option)))
	return
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var checkboxs_newInstances []*CheckBox
	var checkboxs_deletedInstances []*CheckBox

	// parse all staged instances and check if they have a reference
	for checkbox := range stage.CheckBoxs {
		if ref, ok := stage.CheckBoxs_reference[checkbox]; !ok {
			checkboxs_newInstances = append(checkboxs_newInstances, checkbox)
			newInstancesSlice = append(newInstancesSlice, checkbox.GongMarshallIdentifier(stage))
			if stage.CheckBoxs_referenceOrder == nil {
				stage.CheckBoxs_referenceOrder = make(map[*CheckBox]uint)
			}
			stage.CheckBoxs_referenceOrder[checkbox] = stage.CheckBox_stagedOrder[checkbox]
			newInstancesReverseSlice = append(newInstancesReverseSlice, checkbox.GongMarshallUnstaging(stage))
			// delete(stage.CheckBoxs_referenceOrder, checkbox)
			fieldInitializers, pointersInitializations := checkbox.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CheckBox_stagedOrder[ref] = stage.CheckBox_stagedOrder[checkbox]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := checkbox.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, checkbox)
			// delete(stage.CheckBox_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", checkbox.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.CheckBoxs_reference {
		instance := stage.CheckBoxs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.CheckBoxs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			checkboxs_deletedInstances = append(checkboxs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(checkboxs_newInstances)
	lenDeletedInstances += len(checkboxs_deletedInstances)
	var formdivs_newInstances []*FormDiv
	var formdivs_deletedInstances []*FormDiv

	// parse all staged instances and check if they have a reference
	for formdiv := range stage.FormDivs {
		if ref, ok := stage.FormDivs_reference[formdiv]; !ok {
			formdivs_newInstances = append(formdivs_newInstances, formdiv)
			newInstancesSlice = append(newInstancesSlice, formdiv.GongMarshallIdentifier(stage))
			if stage.FormDivs_referenceOrder == nil {
				stage.FormDivs_referenceOrder = make(map[*FormDiv]uint)
			}
			stage.FormDivs_referenceOrder[formdiv] = stage.FormDiv_stagedOrder[formdiv]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formdiv.GongMarshallUnstaging(stage))
			// delete(stage.FormDivs_referenceOrder, formdiv)
			fieldInitializers, pointersInitializations := formdiv.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormDiv_stagedOrder[ref] = stage.FormDiv_stagedOrder[formdiv]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formdiv.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formdiv)
			// delete(stage.FormDiv_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formdiv.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormDivs_reference {
		instance := stage.FormDivs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormDivs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formdivs_deletedInstances = append(formdivs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formdivs_newInstances)
	lenDeletedInstances += len(formdivs_deletedInstances)
	var formeditassocbuttons_newInstances []*FormEditAssocButton
	var formeditassocbuttons_deletedInstances []*FormEditAssocButton

	// parse all staged instances and check if they have a reference
	for formeditassocbutton := range stage.FormEditAssocButtons {
		if ref, ok := stage.FormEditAssocButtons_reference[formeditassocbutton]; !ok {
			formeditassocbuttons_newInstances = append(formeditassocbuttons_newInstances, formeditassocbutton)
			newInstancesSlice = append(newInstancesSlice, formeditassocbutton.GongMarshallIdentifier(stage))
			if stage.FormEditAssocButtons_referenceOrder == nil {
				stage.FormEditAssocButtons_referenceOrder = make(map[*FormEditAssocButton]uint)
			}
			stage.FormEditAssocButtons_referenceOrder[formeditassocbutton] = stage.FormEditAssocButton_stagedOrder[formeditassocbutton]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formeditassocbutton.GongMarshallUnstaging(stage))
			// delete(stage.FormEditAssocButtons_referenceOrder, formeditassocbutton)
			fieldInitializers, pointersInitializations := formeditassocbutton.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormEditAssocButton_stagedOrder[ref] = stage.FormEditAssocButton_stagedOrder[formeditassocbutton]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formeditassocbutton.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formeditassocbutton)
			// delete(stage.FormEditAssocButton_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formeditassocbutton.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormEditAssocButtons_reference {
		instance := stage.FormEditAssocButtons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormEditAssocButtons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formeditassocbuttons_deletedInstances = append(formeditassocbuttons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formeditassocbuttons_newInstances)
	lenDeletedInstances += len(formeditassocbuttons_deletedInstances)
	var formfields_newInstances []*FormField
	var formfields_deletedInstances []*FormField

	// parse all staged instances and check if they have a reference
	for formfield := range stage.FormFields {
		if ref, ok := stage.FormFields_reference[formfield]; !ok {
			formfields_newInstances = append(formfields_newInstances, formfield)
			newInstancesSlice = append(newInstancesSlice, formfield.GongMarshallIdentifier(stage))
			if stage.FormFields_referenceOrder == nil {
				stage.FormFields_referenceOrder = make(map[*FormField]uint)
			}
			stage.FormFields_referenceOrder[formfield] = stage.FormField_stagedOrder[formfield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfield.GongMarshallUnstaging(stage))
			// delete(stage.FormFields_referenceOrder, formfield)
			fieldInitializers, pointersInitializations := formfield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormField_stagedOrder[ref] = stage.FormField_stagedOrder[formfield]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfield)
			// delete(stage.FormField_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfield.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormFields_reference {
		instance := stage.FormFields_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFields[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfields_deletedInstances = append(formfields_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfields_newInstances)
	lenDeletedInstances += len(formfields_deletedInstances)
	var formfielddates_newInstances []*FormFieldDate
	var formfielddates_deletedInstances []*FormFieldDate

	// parse all staged instances and check if they have a reference
	for formfielddate := range stage.FormFieldDates {
		if ref, ok := stage.FormFieldDates_reference[formfielddate]; !ok {
			formfielddates_newInstances = append(formfielddates_newInstances, formfielddate)
			newInstancesSlice = append(newInstancesSlice, formfielddate.GongMarshallIdentifier(stage))
			if stage.FormFieldDates_referenceOrder == nil {
				stage.FormFieldDates_referenceOrder = make(map[*FormFieldDate]uint)
			}
			stage.FormFieldDates_referenceOrder[formfielddate] = stage.FormFieldDate_stagedOrder[formfielddate]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfielddate.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldDates_referenceOrder, formfielddate)
			fieldInitializers, pointersInitializations := formfielddate.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldDate_stagedOrder[ref] = stage.FormFieldDate_stagedOrder[formfielddate]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfielddate.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfielddate)
			// delete(stage.FormFieldDate_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfielddate.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormFieldDates_reference {
		instance := stage.FormFieldDates_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldDates[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfielddates_deletedInstances = append(formfielddates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfielddates_newInstances)
	lenDeletedInstances += len(formfielddates_deletedInstances)
	var formfielddatetimes_newInstances []*FormFieldDateTime
	var formfielddatetimes_deletedInstances []*FormFieldDateTime

	// parse all staged instances and check if they have a reference
	for formfielddatetime := range stage.FormFieldDateTimes {
		if ref, ok := stage.FormFieldDateTimes_reference[formfielddatetime]; !ok {
			formfielddatetimes_newInstances = append(formfielddatetimes_newInstances, formfielddatetime)
			newInstancesSlice = append(newInstancesSlice, formfielddatetime.GongMarshallIdentifier(stage))
			if stage.FormFieldDateTimes_referenceOrder == nil {
				stage.FormFieldDateTimes_referenceOrder = make(map[*FormFieldDateTime]uint)
			}
			stage.FormFieldDateTimes_referenceOrder[formfielddatetime] = stage.FormFieldDateTime_stagedOrder[formfielddatetime]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfielddatetime.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldDateTimes_referenceOrder, formfielddatetime)
			fieldInitializers, pointersInitializations := formfielddatetime.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldDateTime_stagedOrder[ref] = stage.FormFieldDateTime_stagedOrder[formfielddatetime]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfielddatetime.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfielddatetime)
			// delete(stage.FormFieldDateTime_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfielddatetime.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormFieldDateTimes_reference {
		instance := stage.FormFieldDateTimes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldDateTimes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfielddatetimes_deletedInstances = append(formfielddatetimes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfielddatetimes_newInstances)
	lenDeletedInstances += len(formfielddatetimes_deletedInstances)
	var formfieldfloat64s_newInstances []*FormFieldFloat64
	var formfieldfloat64s_deletedInstances []*FormFieldFloat64

	// parse all staged instances and check if they have a reference
	for formfieldfloat64 := range stage.FormFieldFloat64s {
		if ref, ok := stage.FormFieldFloat64s_reference[formfieldfloat64]; !ok {
			formfieldfloat64s_newInstances = append(formfieldfloat64s_newInstances, formfieldfloat64)
			newInstancesSlice = append(newInstancesSlice, formfieldfloat64.GongMarshallIdentifier(stage))
			if stage.FormFieldFloat64s_referenceOrder == nil {
				stage.FormFieldFloat64s_referenceOrder = make(map[*FormFieldFloat64]uint)
			}
			stage.FormFieldFloat64s_referenceOrder[formfieldfloat64] = stage.FormFieldFloat64_stagedOrder[formfieldfloat64]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldfloat64.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldFloat64s_referenceOrder, formfieldfloat64)
			fieldInitializers, pointersInitializations := formfieldfloat64.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldFloat64_stagedOrder[ref] = stage.FormFieldFloat64_stagedOrder[formfieldfloat64]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldfloat64.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldfloat64)
			// delete(stage.FormFieldFloat64_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldfloat64.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormFieldFloat64s_reference {
		instance := stage.FormFieldFloat64s_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldFloat64s[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldfloat64s_deletedInstances = append(formfieldfloat64s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldfloat64s_newInstances)
	lenDeletedInstances += len(formfieldfloat64s_deletedInstances)
	var formfieldints_newInstances []*FormFieldInt
	var formfieldints_deletedInstances []*FormFieldInt

	// parse all staged instances and check if they have a reference
	for formfieldint := range stage.FormFieldInts {
		if ref, ok := stage.FormFieldInts_reference[formfieldint]; !ok {
			formfieldints_newInstances = append(formfieldints_newInstances, formfieldint)
			newInstancesSlice = append(newInstancesSlice, formfieldint.GongMarshallIdentifier(stage))
			if stage.FormFieldInts_referenceOrder == nil {
				stage.FormFieldInts_referenceOrder = make(map[*FormFieldInt]uint)
			}
			stage.FormFieldInts_referenceOrder[formfieldint] = stage.FormFieldInt_stagedOrder[formfieldint]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldint.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldInts_referenceOrder, formfieldint)
			fieldInitializers, pointersInitializations := formfieldint.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldInt_stagedOrder[ref] = stage.FormFieldInt_stagedOrder[formfieldint]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldint.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldint)
			// delete(stage.FormFieldInt_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldint.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormFieldInts_reference {
		instance := stage.FormFieldInts_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldInts[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldints_deletedInstances = append(formfieldints_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldints_newInstances)
	lenDeletedInstances += len(formfieldints_deletedInstances)
	var formfieldselects_newInstances []*FormFieldSelect
	var formfieldselects_deletedInstances []*FormFieldSelect

	// parse all staged instances and check if they have a reference
	for formfieldselect := range stage.FormFieldSelects {
		if ref, ok := stage.FormFieldSelects_reference[formfieldselect]; !ok {
			formfieldselects_newInstances = append(formfieldselects_newInstances, formfieldselect)
			newInstancesSlice = append(newInstancesSlice, formfieldselect.GongMarshallIdentifier(stage))
			if stage.FormFieldSelects_referenceOrder == nil {
				stage.FormFieldSelects_referenceOrder = make(map[*FormFieldSelect]uint)
			}
			stage.FormFieldSelects_referenceOrder[formfieldselect] = stage.FormFieldSelect_stagedOrder[formfieldselect]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldselect.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldSelects_referenceOrder, formfieldselect)
			fieldInitializers, pointersInitializations := formfieldselect.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldSelect_stagedOrder[ref] = stage.FormFieldSelect_stagedOrder[formfieldselect]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldselect.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldselect)
			// delete(stage.FormFieldSelect_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldselect.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormFieldSelects_reference {
		instance := stage.FormFieldSelects_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldSelects[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldselects_deletedInstances = append(formfieldselects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldselects_newInstances)
	lenDeletedInstances += len(formfieldselects_deletedInstances)
	var formfieldstrings_newInstances []*FormFieldString
	var formfieldstrings_deletedInstances []*FormFieldString

	// parse all staged instances and check if they have a reference
	for formfieldstring := range stage.FormFieldStrings {
		if ref, ok := stage.FormFieldStrings_reference[formfieldstring]; !ok {
			formfieldstrings_newInstances = append(formfieldstrings_newInstances, formfieldstring)
			newInstancesSlice = append(newInstancesSlice, formfieldstring.GongMarshallIdentifier(stage))
			if stage.FormFieldStrings_referenceOrder == nil {
				stage.FormFieldStrings_referenceOrder = make(map[*FormFieldString]uint)
			}
			stage.FormFieldStrings_referenceOrder[formfieldstring] = stage.FormFieldString_stagedOrder[formfieldstring]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldstring.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldStrings_referenceOrder, formfieldstring)
			fieldInitializers, pointersInitializations := formfieldstring.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldString_stagedOrder[ref] = stage.FormFieldString_stagedOrder[formfieldstring]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldstring.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldstring)
			// delete(stage.FormFieldString_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldstring.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormFieldStrings_reference {
		instance := stage.FormFieldStrings_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldStrings[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldstrings_deletedInstances = append(formfieldstrings_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldstrings_newInstances)
	lenDeletedInstances += len(formfieldstrings_deletedInstances)
	var formfieldtimes_newInstances []*FormFieldTime
	var formfieldtimes_deletedInstances []*FormFieldTime

	// parse all staged instances and check if they have a reference
	for formfieldtime := range stage.FormFieldTimes {
		if ref, ok := stage.FormFieldTimes_reference[formfieldtime]; !ok {
			formfieldtimes_newInstances = append(formfieldtimes_newInstances, formfieldtime)
			newInstancesSlice = append(newInstancesSlice, formfieldtime.GongMarshallIdentifier(stage))
			if stage.FormFieldTimes_referenceOrder == nil {
				stage.FormFieldTimes_referenceOrder = make(map[*FormFieldTime]uint)
			}
			stage.FormFieldTimes_referenceOrder[formfieldtime] = stage.FormFieldTime_stagedOrder[formfieldtime]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formfieldtime.GongMarshallUnstaging(stage))
			// delete(stage.FormFieldTimes_referenceOrder, formfieldtime)
			fieldInitializers, pointersInitializations := formfieldtime.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormFieldTime_stagedOrder[ref] = stage.FormFieldTime_stagedOrder[formfieldtime]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formfieldtime.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formfieldtime)
			// delete(stage.FormFieldTime_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formfieldtime.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormFieldTimes_reference {
		instance := stage.FormFieldTimes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormFieldTimes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formfieldtimes_deletedInstances = append(formfieldtimes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formfieldtimes_newInstances)
	lenDeletedInstances += len(formfieldtimes_deletedInstances)
	var formgroups_newInstances []*FormGroup
	var formgroups_deletedInstances []*FormGroup

	// parse all staged instances and check if they have a reference
	for formgroup := range stage.FormGroups {
		if ref, ok := stage.FormGroups_reference[formgroup]; !ok {
			formgroups_newInstances = append(formgroups_newInstances, formgroup)
			newInstancesSlice = append(newInstancesSlice, formgroup.GongMarshallIdentifier(stage))
			if stage.FormGroups_referenceOrder == nil {
				stage.FormGroups_referenceOrder = make(map[*FormGroup]uint)
			}
			stage.FormGroups_referenceOrder[formgroup] = stage.FormGroup_stagedOrder[formgroup]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formgroup.GongMarshallUnstaging(stage))
			// delete(stage.FormGroups_referenceOrder, formgroup)
			fieldInitializers, pointersInitializations := formgroup.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormGroup_stagedOrder[ref] = stage.FormGroup_stagedOrder[formgroup]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formgroup.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formgroup)
			// delete(stage.FormGroup_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formgroup.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormGroups_reference {
		instance := stage.FormGroups_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormGroups[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formgroups_deletedInstances = append(formgroups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formgroups_newInstances)
	lenDeletedInstances += len(formgroups_deletedInstances)
	var formsortassocbuttons_newInstances []*FormSortAssocButton
	var formsortassocbuttons_deletedInstances []*FormSortAssocButton

	// parse all staged instances and check if they have a reference
	for formsortassocbutton := range stage.FormSortAssocButtons {
		if ref, ok := stage.FormSortAssocButtons_reference[formsortassocbutton]; !ok {
			formsortassocbuttons_newInstances = append(formsortassocbuttons_newInstances, formsortassocbutton)
			newInstancesSlice = append(newInstancesSlice, formsortassocbutton.GongMarshallIdentifier(stage))
			if stage.FormSortAssocButtons_referenceOrder == nil {
				stage.FormSortAssocButtons_referenceOrder = make(map[*FormSortAssocButton]uint)
			}
			stage.FormSortAssocButtons_referenceOrder[formsortassocbutton] = stage.FormSortAssocButton_stagedOrder[formsortassocbutton]
			newInstancesReverseSlice = append(newInstancesReverseSlice, formsortassocbutton.GongMarshallUnstaging(stage))
			// delete(stage.FormSortAssocButtons_referenceOrder, formsortassocbutton)
			fieldInitializers, pointersInitializations := formsortassocbutton.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormSortAssocButton_stagedOrder[ref] = stage.FormSortAssocButton_stagedOrder[formsortassocbutton]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := formsortassocbutton.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, formsortassocbutton)
			// delete(stage.FormSortAssocButton_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", formsortassocbutton.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.FormSortAssocButtons_reference {
		instance := stage.FormSortAssocButtons_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.FormSortAssocButtons[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			formsortassocbuttons_deletedInstances = append(formsortassocbuttons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(formsortassocbuttons_newInstances)
	lenDeletedInstances += len(formsortassocbuttons_deletedInstances)
	var options_newInstances []*Option
	var options_deletedInstances []*Option

	// parse all staged instances and check if they have a reference
	for option := range stage.Options {
		if ref, ok := stage.Options_reference[option]; !ok {
			options_newInstances = append(options_newInstances, option)
			newInstancesSlice = append(newInstancesSlice, option.GongMarshallIdentifier(stage))
			if stage.Options_referenceOrder == nil {
				stage.Options_referenceOrder = make(map[*Option]uint)
			}
			stage.Options_referenceOrder[option] = stage.Option_stagedOrder[option]
			newInstancesReverseSlice = append(newInstancesReverseSlice, option.GongMarshallUnstaging(stage))
			// delete(stage.Options_referenceOrder, option)
			fieldInitializers, pointersInitializations := option.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Option_stagedOrder[ref] = stage.Option_stagedOrder[option]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := option.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, option)
			// delete(stage.Option_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", option.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Options_reference {
		instance := stage.Options_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Options[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			options_deletedInstances = append(options_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(options_newInstances)
	lenDeletedInstances += len(options_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.CheckBoxs_reference = make(map[*CheckBox]*CheckBox)
	stage.CheckBoxs_referenceOrder = make(map[*CheckBox]uint) // diff Unstage needs the reference order
	stage.CheckBoxs_instance = make(map[*CheckBox]*CheckBox)
	for instance := range stage.CheckBoxs {
		_copy := instance.GongCopy().(*CheckBox)
		stage.CheckBoxs_reference[instance] = _copy
		stage.CheckBoxs_instance[_copy] = instance
		stage.CheckBoxs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormDivs_reference = make(map[*FormDiv]*FormDiv)
	stage.FormDivs_referenceOrder = make(map[*FormDiv]uint) // diff Unstage needs the reference order
	stage.FormDivs_instance = make(map[*FormDiv]*FormDiv)
	for instance := range stage.FormDivs {
		_copy := instance.GongCopy().(*FormDiv)
		stage.FormDivs_reference[instance] = _copy
		stage.FormDivs_instance[_copy] = instance
		stage.FormDivs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormEditAssocButtons_reference = make(map[*FormEditAssocButton]*FormEditAssocButton)
	stage.FormEditAssocButtons_referenceOrder = make(map[*FormEditAssocButton]uint) // diff Unstage needs the reference order
	stage.FormEditAssocButtons_instance = make(map[*FormEditAssocButton]*FormEditAssocButton)
	for instance := range stage.FormEditAssocButtons {
		_copy := instance.GongCopy().(*FormEditAssocButton)
		stage.FormEditAssocButtons_reference[instance] = _copy
		stage.FormEditAssocButtons_instance[_copy] = instance
		stage.FormEditAssocButtons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFields_reference = make(map[*FormField]*FormField)
	stage.FormFields_referenceOrder = make(map[*FormField]uint) // diff Unstage needs the reference order
	stage.FormFields_instance = make(map[*FormField]*FormField)
	for instance := range stage.FormFields {
		_copy := instance.GongCopy().(*FormField)
		stage.FormFields_reference[instance] = _copy
		stage.FormFields_instance[_copy] = instance
		stage.FormFields_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldDates_reference = make(map[*FormFieldDate]*FormFieldDate)
	stage.FormFieldDates_referenceOrder = make(map[*FormFieldDate]uint) // diff Unstage needs the reference order
	stage.FormFieldDates_instance = make(map[*FormFieldDate]*FormFieldDate)
	for instance := range stage.FormFieldDates {
		_copy := instance.GongCopy().(*FormFieldDate)
		stage.FormFieldDates_reference[instance] = _copy
		stage.FormFieldDates_instance[_copy] = instance
		stage.FormFieldDates_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldDateTimes_reference = make(map[*FormFieldDateTime]*FormFieldDateTime)
	stage.FormFieldDateTimes_referenceOrder = make(map[*FormFieldDateTime]uint) // diff Unstage needs the reference order
	stage.FormFieldDateTimes_instance = make(map[*FormFieldDateTime]*FormFieldDateTime)
	for instance := range stage.FormFieldDateTimes {
		_copy := instance.GongCopy().(*FormFieldDateTime)
		stage.FormFieldDateTimes_reference[instance] = _copy
		stage.FormFieldDateTimes_instance[_copy] = instance
		stage.FormFieldDateTimes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldFloat64s_reference = make(map[*FormFieldFloat64]*FormFieldFloat64)
	stage.FormFieldFloat64s_referenceOrder = make(map[*FormFieldFloat64]uint) // diff Unstage needs the reference order
	stage.FormFieldFloat64s_instance = make(map[*FormFieldFloat64]*FormFieldFloat64)
	for instance := range stage.FormFieldFloat64s {
		_copy := instance.GongCopy().(*FormFieldFloat64)
		stage.FormFieldFloat64s_reference[instance] = _copy
		stage.FormFieldFloat64s_instance[_copy] = instance
		stage.FormFieldFloat64s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldInts_reference = make(map[*FormFieldInt]*FormFieldInt)
	stage.FormFieldInts_referenceOrder = make(map[*FormFieldInt]uint) // diff Unstage needs the reference order
	stage.FormFieldInts_instance = make(map[*FormFieldInt]*FormFieldInt)
	for instance := range stage.FormFieldInts {
		_copy := instance.GongCopy().(*FormFieldInt)
		stage.FormFieldInts_reference[instance] = _copy
		stage.FormFieldInts_instance[_copy] = instance
		stage.FormFieldInts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldSelects_reference = make(map[*FormFieldSelect]*FormFieldSelect)
	stage.FormFieldSelects_referenceOrder = make(map[*FormFieldSelect]uint) // diff Unstage needs the reference order
	stage.FormFieldSelects_instance = make(map[*FormFieldSelect]*FormFieldSelect)
	for instance := range stage.FormFieldSelects {
		_copy := instance.GongCopy().(*FormFieldSelect)
		stage.FormFieldSelects_reference[instance] = _copy
		stage.FormFieldSelects_instance[_copy] = instance
		stage.FormFieldSelects_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldStrings_reference = make(map[*FormFieldString]*FormFieldString)
	stage.FormFieldStrings_referenceOrder = make(map[*FormFieldString]uint) // diff Unstage needs the reference order
	stage.FormFieldStrings_instance = make(map[*FormFieldString]*FormFieldString)
	for instance := range stage.FormFieldStrings {
		_copy := instance.GongCopy().(*FormFieldString)
		stage.FormFieldStrings_reference[instance] = _copy
		stage.FormFieldStrings_instance[_copy] = instance
		stage.FormFieldStrings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormFieldTimes_reference = make(map[*FormFieldTime]*FormFieldTime)
	stage.FormFieldTimes_referenceOrder = make(map[*FormFieldTime]uint) // diff Unstage needs the reference order
	stage.FormFieldTimes_instance = make(map[*FormFieldTime]*FormFieldTime)
	for instance := range stage.FormFieldTimes {
		_copy := instance.GongCopy().(*FormFieldTime)
		stage.FormFieldTimes_reference[instance] = _copy
		stage.FormFieldTimes_instance[_copy] = instance
		stage.FormFieldTimes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormGroups_reference = make(map[*FormGroup]*FormGroup)
	stage.FormGroups_referenceOrder = make(map[*FormGroup]uint) // diff Unstage needs the reference order
	stage.FormGroups_instance = make(map[*FormGroup]*FormGroup)
	for instance := range stage.FormGroups {
		_copy := instance.GongCopy().(*FormGroup)
		stage.FormGroups_reference[instance] = _copy
		stage.FormGroups_instance[_copy] = instance
		stage.FormGroups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FormSortAssocButtons_reference = make(map[*FormSortAssocButton]*FormSortAssocButton)
	stage.FormSortAssocButtons_referenceOrder = make(map[*FormSortAssocButton]uint) // diff Unstage needs the reference order
	stage.FormSortAssocButtons_instance = make(map[*FormSortAssocButton]*FormSortAssocButton)
	for instance := range stage.FormSortAssocButtons {
		_copy := instance.GongCopy().(*FormSortAssocButton)
		stage.FormSortAssocButtons_reference[instance] = _copy
		stage.FormSortAssocButtons_instance[_copy] = instance
		stage.FormSortAssocButtons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Options_reference = make(map[*Option]*Option)
	stage.Options_referenceOrder = make(map[*Option]uint) // diff Unstage needs the reference order
	stage.Options_instance = make(map[*Option]*Option)
	for instance := range stage.Options {
		_copy := instance.GongCopy().(*Option)
		stage.Options_reference[instance] = _copy
		stage.Options_instance[_copy] = instance
		stage.Options_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.CheckBoxs {
		reference := stage.CheckBoxs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormDivs {
		reference := stage.FormDivs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormEditAssocButtons {
		reference := stage.FormEditAssocButtons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFields {
		reference := stage.FormFields_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldDates {
		reference := stage.FormFieldDates_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldDateTimes {
		reference := stage.FormFieldDateTimes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldFloat64s {
		reference := stage.FormFieldFloat64s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldInts {
		reference := stage.FormFieldInts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldSelects {
		reference := stage.FormFieldSelects_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldStrings {
		reference := stage.FormFieldStrings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormFieldTimes {
		reference := stage.FormFieldTimes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormGroups {
		reference := stage.FormGroups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FormSortAssocButtons {
		reference := stage.FormSortAssocButtons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Options {
		reference := stage.Options_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (checkbox *CheckBox) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CheckBox_stagedOrder[checkbox]; ok {
		return order
	}
	if order, ok := stage.CheckBoxs_referenceOrder[checkbox]; ok {
		return order
	} else {
		log.Printf("instance %p of type CheckBox was not staged and does not have a reference order", checkbox)
		return 0
	}
}

func (formdiv *FormDiv) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormDiv_stagedOrder[formdiv]; ok {
		return order
	}
	if order, ok := stage.FormDivs_referenceOrder[formdiv]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormDiv was not staged and does not have a reference order", formdiv)
		return 0
	}
}

func (formeditassocbutton *FormEditAssocButton) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormEditAssocButton_stagedOrder[formeditassocbutton]; ok {
		return order
	}
	if order, ok := stage.FormEditAssocButtons_referenceOrder[formeditassocbutton]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormEditAssocButton was not staged and does not have a reference order", formeditassocbutton)
		return 0
	}
}

func (formfield *FormField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormField_stagedOrder[formfield]; ok {
		return order
	}
	if order, ok := stage.FormFields_referenceOrder[formfield]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormField was not staged and does not have a reference order", formfield)
		return 0
	}
}

func (formfielddate *FormFieldDate) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldDate_stagedOrder[formfielddate]; ok {
		return order
	}
	if order, ok := stage.FormFieldDates_referenceOrder[formfielddate]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldDate was not staged and does not have a reference order", formfielddate)
		return 0
	}
}

func (formfielddatetime *FormFieldDateTime) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldDateTime_stagedOrder[formfielddatetime]; ok {
		return order
	}
	if order, ok := stage.FormFieldDateTimes_referenceOrder[formfielddatetime]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldDateTime was not staged and does not have a reference order", formfielddatetime)
		return 0
	}
}

func (formfieldfloat64 *FormFieldFloat64) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldFloat64_stagedOrder[formfieldfloat64]; ok {
		return order
	}
	if order, ok := stage.FormFieldFloat64s_referenceOrder[formfieldfloat64]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldFloat64 was not staged and does not have a reference order", formfieldfloat64)
		return 0
	}
}

func (formfieldint *FormFieldInt) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldInt_stagedOrder[formfieldint]; ok {
		return order
	}
	if order, ok := stage.FormFieldInts_referenceOrder[formfieldint]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldInt was not staged and does not have a reference order", formfieldint)
		return 0
	}
}

func (formfieldselect *FormFieldSelect) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldSelect_stagedOrder[formfieldselect]; ok {
		return order
	}
	if order, ok := stage.FormFieldSelects_referenceOrder[formfieldselect]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldSelect was not staged and does not have a reference order", formfieldselect)
		return 0
	}
}

func (formfieldstring *FormFieldString) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldString_stagedOrder[formfieldstring]; ok {
		return order
	}
	if order, ok := stage.FormFieldStrings_referenceOrder[formfieldstring]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldString was not staged and does not have a reference order", formfieldstring)
		return 0
	}
}

func (formfieldtime *FormFieldTime) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormFieldTime_stagedOrder[formfieldtime]; ok {
		return order
	}
	if order, ok := stage.FormFieldTimes_referenceOrder[formfieldtime]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormFieldTime was not staged and does not have a reference order", formfieldtime)
		return 0
	}
}

func (formgroup *FormGroup) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormGroup_stagedOrder[formgroup]; ok {
		return order
	}
	if order, ok := stage.FormGroups_referenceOrder[formgroup]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormGroup was not staged and does not have a reference order", formgroup)
		return 0
	}
}

func (formsortassocbutton *FormSortAssocButton) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormSortAssocButton_stagedOrder[formsortassocbutton]; ok {
		return order
	}
	if order, ok := stage.FormSortAssocButtons_referenceOrder[formsortassocbutton]; ok {
		return order
	} else {
		log.Printf("instance %p of type FormSortAssocButton was not staged and does not have a reference order", formsortassocbutton)
		return 0
	}
}

func (option *Option) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Option_stagedOrder[option]; ok {
		return order
	}
	if order, ok := stage.Options_referenceOrder[option]; ok {
		return order
	} else {
		log.Printf("instance %p of type Option was not staged and does not have a reference order", option)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (checkbox *CheckBox) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkbox *CheckBox) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetOrder(stage))
}

func (formdiv *FormDiv) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formdiv.GongGetGongstructName(), formdiv.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formdiv *FormDiv) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formdiv.GongGetGongstructName(), formdiv.GongGetOrder(stage))
}

func (formeditassocbutton *FormEditAssocButton) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formeditassocbutton.GongGetGongstructName(), formeditassocbutton.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formeditassocbutton *FormEditAssocButton) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formeditassocbutton.GongGetGongstructName(), formeditassocbutton.GongGetOrder(stage))
}

func (formfield *FormField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfield.GongGetGongstructName(), formfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfield *FormField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfield.GongGetGongstructName(), formfield.GongGetOrder(stage))
}

func (formfielddate *FormFieldDate) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddate.GongGetGongstructName(), formfielddate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfielddate *FormFieldDate) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddate.GongGetGongstructName(), formfielddate.GongGetOrder(stage))
}

func (formfielddatetime *FormFieldDateTime) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddatetime.GongGetGongstructName(), formfielddatetime.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfielddatetime *FormFieldDateTime) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfielddatetime.GongGetGongstructName(), formfielddatetime.GongGetOrder(stage))
}

func (formfieldfloat64 *FormFieldFloat64) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldfloat64.GongGetGongstructName(), formfieldfloat64.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldfloat64 *FormFieldFloat64) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldfloat64.GongGetGongstructName(), formfieldfloat64.GongGetOrder(stage))
}

func (formfieldint *FormFieldInt) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldint.GongGetGongstructName(), formfieldint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldint *FormFieldInt) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldint.GongGetGongstructName(), formfieldint.GongGetOrder(stage))
}

func (formfieldselect *FormFieldSelect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldselect.GongGetGongstructName(), formfieldselect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldselect *FormFieldSelect) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldselect.GongGetGongstructName(), formfieldselect.GongGetOrder(stage))
}

func (formfieldstring *FormFieldString) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldstring.GongGetGongstructName(), formfieldstring.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldstring *FormFieldString) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldstring.GongGetGongstructName(), formfieldstring.GongGetOrder(stage))
}

func (formfieldtime *FormFieldTime) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldtime.GongGetGongstructName(), formfieldtime.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldtime *FormFieldTime) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formfieldtime.GongGetGongstructName(), formfieldtime.GongGetOrder(stage))
}

func (formgroup *FormGroup) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formgroup.GongGetGongstructName(), formgroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formgroup *FormGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formgroup.GongGetGongstructName(), formgroup.GongGetOrder(stage))
}

func (formsortassocbutton *FormSortAssocButton) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formsortassocbutton.GongGetGongstructName(), formsortassocbutton.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formsortassocbutton *FormSortAssocButton) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formsortassocbutton.GongGetGongstructName(), formsortassocbutton.GongGetOrder(stage))
}

func (option *Option) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", option.GongGetGongstructName(), option.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (option *Option) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", option.GongGetGongstructName(), option.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (checkbox *CheckBox) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CheckBox")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(checkbox.Name))
	return
}

func (formdiv *FormDiv) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormDiv")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formdiv.Name))
	return
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormEditAssocButton")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formeditassocbutton.Name))
	return
}

func (formfield *FormField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfield.Name))
	return
}

func (formfielddate *FormFieldDate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfielddate.Name))
	return
}

func (formfielddatetime *FormFieldDateTime) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDateTime")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfielddatetime.Name))
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldFloat64")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldfloat64.Name))
	return
}

func (formfieldint *FormFieldInt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldInt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldint.Name))
	return
}

func (formfieldselect *FormFieldSelect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldSelect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldselect.Name))
	return
}

func (formfieldstring *FormFieldString) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldString")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldstring.Name))
	return
}

func (formfieldtime *FormFieldTime) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldTime")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formfieldtime.Name))
	return
}

func (formgroup *FormGroup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormGroup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formgroup.Name))
	return
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormSortAssocButton")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(formsortassocbutton.Name))
	return
}

func (option *Option) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", option.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Option")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(option.Name))
	return
}

// insertion point for unstaging
func (checkbox *CheckBox) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetReferenceIdentifier(stage))
	return
}

func (formdiv *FormDiv) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formdiv.GongGetReferenceIdentifier(stage))
	return
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formeditassocbutton.GongGetReferenceIdentifier(stage))
	return
}

func (formfield *FormField) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfield.GongGetReferenceIdentifier(stage))
	return
}

func (formfielddate *FormFieldDate) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddate.GongGetReferenceIdentifier(stage))
	return
}

func (formfielddatetime *FormFieldDateTime) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddatetime.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldfloat64.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldint *FormFieldInt) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldint.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldselect *FormFieldSelect) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldselect.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldstring *FormFieldString) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldstring.GongGetReferenceIdentifier(stage))
	return
}

func (formfieldtime *FormFieldTime) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldtime.GongGetReferenceIdentifier(stage))
	return
}

func (formgroup *FormGroup) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formgroup.GongGetReferenceIdentifier(stage))
	return
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formsortassocbutton.GongGetReferenceIdentifier(stage))
	return
}

func (option *Option) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", option.GongGetReferenceIdentifier(stage))
	return
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// end of template
