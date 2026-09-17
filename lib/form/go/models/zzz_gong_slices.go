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
	checkbox.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formdiv *FormDiv) GongCopy() GongstructIF {
	newInstance := new(FormDiv)
	formdiv.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formeditassocbutton *FormEditAssocButton) GongCopy() GongstructIF {
	newInstance := new(FormEditAssocButton)
	formeditassocbutton.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formfield *FormField) GongCopy() GongstructIF {
	newInstance := new(FormField)
	formfield.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formfielddate *FormFieldDate) GongCopy() GongstructIF {
	newInstance := new(FormFieldDate)
	formfielddate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formfielddatetime *FormFieldDateTime) GongCopy() GongstructIF {
	newInstance := new(FormFieldDateTime)
	formfielddatetime.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formfieldfloat64 *FormFieldFloat64) GongCopy() GongstructIF {
	newInstance := new(FormFieldFloat64)
	formfieldfloat64.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formfieldint *FormFieldInt) GongCopy() GongstructIF {
	newInstance := new(FormFieldInt)
	formfieldint.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formfieldselect *FormFieldSelect) GongCopy() GongstructIF {
	newInstance := new(FormFieldSelect)
	formfieldselect.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formfieldstring *FormFieldString) GongCopy() GongstructIF {
	newInstance := new(FormFieldString)
	formfieldstring.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formfieldtime *FormFieldTime) GongCopy() GongstructIF {
	newInstance := new(FormFieldTime)
	formfieldtime.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formgroup *FormGroup) GongCopy() GongstructIF {
	newInstance := new(FormGroup)
	formgroup.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formsortassocbutton *FormSortAssocButton) GongCopy() GongstructIF {
	newInstance := new(FormSortAssocButton)
	formsortassocbutton.GongCopyBasicFields(newInstance)
	return newInstance
}

func (option *Option) GongCopy() GongstructIF {
	newInstance := new(Option)
	option.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (checkbox *CheckBox) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(checkbox).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(checkbox), uint64(stage.GetOrder(checkbox)))
	return
}

func (formdiv *FormDiv) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formdiv).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formdiv), uint64(stage.GetOrder(formdiv)))
	return
}

func (formeditassocbutton *FormEditAssocButton) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formeditassocbutton).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formeditassocbutton), uint64(stage.GetOrder(formeditassocbutton)))
	return
}

func (formfield *FormField) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfield).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formfield), uint64(stage.GetOrder(formfield)))
	return
}

func (formfielddate *FormFieldDate) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfielddate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formfielddate), uint64(stage.GetOrder(formfielddate)))
	return
}

func (formfielddatetime *FormFieldDateTime) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfielddatetime).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formfielddatetime), uint64(stage.GetOrder(formfielddatetime)))
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldfloat64).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formfieldfloat64), uint64(stage.GetOrder(formfieldfloat64)))
	return
}

func (formfieldint *FormFieldInt) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldint).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formfieldint), uint64(stage.GetOrder(formfieldint)))
	return
}

func (formfieldselect *FormFieldSelect) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldselect).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formfieldselect), uint64(stage.GetOrder(formfieldselect)))
	return
}

func (formfieldstring *FormFieldString) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldstring).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formfieldstring), uint64(stage.GetOrder(formfieldstring)))
	return
}

func (formfieldtime *FormFieldTime) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formfieldtime).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formfieldtime), uint64(stage.GetOrder(formfieldtime)))
	return
}

func (formgroup *FormGroup) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formgroup).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formgroup), uint64(stage.GetOrder(formgroup)))
	return
}

func (formsortassocbutton *FormSortAssocButton) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formsortassocbutton).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formsortassocbutton), uint64(stage.GetOrder(formsortassocbutton)))
	return
}

func (option *Option) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(option).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(option), uint64(stage.GetOrder(option)))
	return
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.CheckBoxs,
		stage.CheckBox_stagedOrder,
		stage.CheckBoxs_reference,
		&stage.CheckBoxs_referenceOrder,
		stage.CheckBoxs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormDivs,
		stage.FormDiv_stagedOrder,
		stage.FormDivs_reference,
		&stage.FormDivs_referenceOrder,
		stage.FormDivs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormEditAssocButtons,
		stage.FormEditAssocButton_stagedOrder,
		stage.FormEditAssocButtons_reference,
		&stage.FormEditAssocButtons_referenceOrder,
		stage.FormEditAssocButtons_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormFields,
		stage.FormField_stagedOrder,
		stage.FormFields_reference,
		&stage.FormFields_referenceOrder,
		stage.FormFields_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormFieldDates,
		stage.FormFieldDate_stagedOrder,
		stage.FormFieldDates_reference,
		&stage.FormFieldDates_referenceOrder,
		stage.FormFieldDates_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormFieldDateTimes,
		stage.FormFieldDateTime_stagedOrder,
		stage.FormFieldDateTimes_reference,
		&stage.FormFieldDateTimes_referenceOrder,
		stage.FormFieldDateTimes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormFieldFloat64s,
		stage.FormFieldFloat64_stagedOrder,
		stage.FormFieldFloat64s_reference,
		&stage.FormFieldFloat64s_referenceOrder,
		stage.FormFieldFloat64s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormFieldInts,
		stage.FormFieldInt_stagedOrder,
		stage.FormFieldInts_reference,
		&stage.FormFieldInts_referenceOrder,
		stage.FormFieldInts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormFieldSelects,
		stage.FormFieldSelect_stagedOrder,
		stage.FormFieldSelects_reference,
		&stage.FormFieldSelects_referenceOrder,
		stage.FormFieldSelects_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormFieldStrings,
		stage.FormFieldString_stagedOrder,
		stage.FormFieldStrings_reference,
		&stage.FormFieldStrings_referenceOrder,
		stage.FormFieldStrings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormFieldTimes,
		stage.FormFieldTime_stagedOrder,
		stage.FormFieldTimes_reference,
		&stage.FormFieldTimes_referenceOrder,
		stage.FormFieldTimes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormGroups,
		stage.FormGroup_stagedOrder,
		stage.FormGroups_reference,
		&stage.FormGroups_referenceOrder,
		stage.FormGroups_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FormSortAssocButtons,
		stage.FormSortAssocButton_stagedOrder,
		stage.FormSortAssocButtons_reference,
		&stage.FormSortAssocButtons_referenceOrder,
		stage.FormSortAssocButtons_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Options,
		stage.Option_stagedOrder,
		stage.Options_reference,
		&stage.Options_referenceOrder,
		stage.Options_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(checkbox.Name))
	return
}

func (formdiv *FormDiv) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormDiv")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formdiv.Name))
	return
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormEditAssocButton")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formeditassocbutton.Name))
	return
}

func (formfield *FormField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formfield.Name))
	return
}

func (formfielddate *FormFieldDate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formfielddate.Name))
	return
}

func (formfielddatetime *FormFieldDateTime) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDateTime")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formfielddatetime.Name))
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldFloat64")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formfieldfloat64.Name))
	return
}

func (formfieldint *FormFieldInt) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldInt")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formfieldint.Name))
	return
}

func (formfieldselect *FormFieldSelect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldSelect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formfieldselect.Name))
	return
}

func (formfieldstring *FormFieldString) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldString")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formfieldstring.Name))
	return
}

func (formfieldtime *FormFieldTime) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldTime")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formfieldtime.Name))
	return
}

func (formgroup *FormGroup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormGroup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formgroup.Name))
	return
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormSortAssocButton")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formsortassocbutton.Name))
	return
}

func (option *Option) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", option.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Option")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(option.Name))
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

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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
