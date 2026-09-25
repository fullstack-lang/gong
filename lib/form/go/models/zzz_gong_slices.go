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

	// Compute reverse map for named struct FormFieldSelect
	// insertion point per field
	stage.FormFieldSelect_Options_reverseMap = make(map[*Option]*FormFieldSelect)
	for formfieldselect := range stage.FormFieldSelects {
		_ = formfieldselect
		for _, _option := range formfieldselect.Options {
			stage.FormFieldSelect_Options_reverseMap[_option] = formfieldselect
		}
	}

	// Compute reverse map for named struct FormGroup
	// insertion point per field
	stage.FormGroup_FormDivs_reverseMap = make(map[*FormDiv]*FormGroup)
	for formgroup := range stage.FormGroups {
		_ = formgroup
		for _, _formdiv := range formgroup.FormDivs {
			stage.FormGroup_FormDivs_reverseMap[_formdiv] = formgroup
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.CheckBoxs)

	res = __gong__appendInstances(res, stage.FormDivs)

	res = __gong__appendInstances(res, stage.FormEditAssocButtons)

	res = __gong__appendInstances(res, stage.FormFields)

	res = __gong__appendInstances(res, stage.FormFieldDates)

	res = __gong__appendInstances(res, stage.FormFieldDateTimes)

	res = __gong__appendInstances(res, stage.FormFieldFloat64s)

	res = __gong__appendInstances(res, stage.FormFieldInts)

	res = __gong__appendInstances(res, stage.FormFieldSelects)

	res = __gong__appendInstances(res, stage.FormFieldStrings)

	res = __gong__appendInstances(res, stage.FormFieldTimes)

	res = __gong__appendInstances(res, stage.FormGroups)

	res = __gong__appendInstances(res, stage.FormSortAssocButtons)

	res = __gong__appendInstances(res, stage.Options)

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
func (checkbox *CheckBox) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, checkbox)
}

func (formdiv *FormDiv) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formdiv)
}

func (formeditassocbutton *FormEditAssocButton) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formeditassocbutton)
}

func (formfield *FormField) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formfield)
}

func (formfielddate *FormFieldDate) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formfielddate)
}

func (formfielddatetime *FormFieldDateTime) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formfielddatetime)
}

func (formfieldfloat64 *FormFieldFloat64) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formfieldfloat64)
}

func (formfieldint *FormFieldInt) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formfieldint)
}

func (formfieldselect *FormFieldSelect) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formfieldselect)
}

func (formfieldstring *FormFieldString) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formfieldstring)
}

func (formfieldtime *FormFieldTime) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formfieldtime)
}

func (formgroup *FormGroup) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formgroup)
}

func (formsortassocbutton *FormSortAssocButton) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formsortassocbutton)
}

func (option *Option) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, option)
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
	__gong__computeReferencePass1(stage, stage.CheckBoxs, &stage.CheckBoxs_reference, &stage.CheckBoxs_referenceOrder, &stage.CheckBoxs_instance)

	__gong__computeReferencePass1(stage, stage.FormDivs, &stage.FormDivs_reference, &stage.FormDivs_referenceOrder, &stage.FormDivs_instance)

	__gong__computeReferencePass1(stage, stage.FormEditAssocButtons, &stage.FormEditAssocButtons_reference, &stage.FormEditAssocButtons_referenceOrder, &stage.FormEditAssocButtons_instance)

	__gong__computeReferencePass1(stage, stage.FormFields, &stage.FormFields_reference, &stage.FormFields_referenceOrder, &stage.FormFields_instance)

	__gong__computeReferencePass1(stage, stage.FormFieldDates, &stage.FormFieldDates_reference, &stage.FormFieldDates_referenceOrder, &stage.FormFieldDates_instance)

	__gong__computeReferencePass1(stage, stage.FormFieldDateTimes, &stage.FormFieldDateTimes_reference, &stage.FormFieldDateTimes_referenceOrder, &stage.FormFieldDateTimes_instance)

	__gong__computeReferencePass1(stage, stage.FormFieldFloat64s, &stage.FormFieldFloat64s_reference, &stage.FormFieldFloat64s_referenceOrder, &stage.FormFieldFloat64s_instance)

	__gong__computeReferencePass1(stage, stage.FormFieldInts, &stage.FormFieldInts_reference, &stage.FormFieldInts_referenceOrder, &stage.FormFieldInts_instance)

	__gong__computeReferencePass1(stage, stage.FormFieldSelects, &stage.FormFieldSelects_reference, &stage.FormFieldSelects_referenceOrder, &stage.FormFieldSelects_instance)

	__gong__computeReferencePass1(stage, stage.FormFieldStrings, &stage.FormFieldStrings_reference, &stage.FormFieldStrings_referenceOrder, &stage.FormFieldStrings_instance)

	__gong__computeReferencePass1(stage, stage.FormFieldTimes, &stage.FormFieldTimes_reference, &stage.FormFieldTimes_referenceOrder, &stage.FormFieldTimes_instance)

	__gong__computeReferencePass1(stage, stage.FormGroups, &stage.FormGroups_reference, &stage.FormGroups_referenceOrder, &stage.FormGroups_instance)

	__gong__computeReferencePass1(stage, stage.FormSortAssocButtons, &stage.FormSortAssocButtons_reference, &stage.FormSortAssocButtons_referenceOrder, &stage.FormSortAssocButtons_instance)

	__gong__computeReferencePass1(stage, stage.Options, &stage.Options_reference, &stage.Options_referenceOrder, &stage.Options_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CheckBoxs, stage.CheckBoxs_reference, stage)

	__gong__computeReferencePass2(stage.FormDivs, stage.FormDivs_reference, stage)

	__gong__computeReferencePass2(stage.FormEditAssocButtons, stage.FormEditAssocButtons_reference, stage)

	__gong__computeReferencePass2(stage.FormFields, stage.FormFields_reference, stage)

	__gong__computeReferencePass2(stage.FormFieldDates, stage.FormFieldDates_reference, stage)

	__gong__computeReferencePass2(stage.FormFieldDateTimes, stage.FormFieldDateTimes_reference, stage)

	__gong__computeReferencePass2(stage.FormFieldFloat64s, stage.FormFieldFloat64s_reference, stage)

	__gong__computeReferencePass2(stage.FormFieldInts, stage.FormFieldInts_reference, stage)

	__gong__computeReferencePass2(stage.FormFieldSelects, stage.FormFieldSelects_reference, stage)

	__gong__computeReferencePass2(stage.FormFieldStrings, stage.FormFieldStrings_reference, stage)

	__gong__computeReferencePass2(stage.FormFieldTimes, stage.FormFieldTimes_reference, stage)

	__gong__computeReferencePass2(stage.FormGroups, stage.FormGroups_reference, stage)

	__gong__computeReferencePass2(stage.FormSortAssocButtons, stage.FormSortAssocButtons_reference, stage)

	__gong__computeReferencePass2(stage.Options, stage.Options_reference, stage)

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
	return __gong__getOrder(stage.CheckBox_stagedOrder, stage.CheckBoxs_referenceOrder, checkbox, "CheckBox")
}

func (formdiv *FormDiv) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormDiv_stagedOrder, stage.FormDivs_referenceOrder, formdiv, "FormDiv")
}

func (formeditassocbutton *FormEditAssocButton) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormEditAssocButton_stagedOrder, stage.FormEditAssocButtons_referenceOrder, formeditassocbutton, "FormEditAssocButton")
}

func (formfield *FormField) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormField_stagedOrder, stage.FormFields_referenceOrder, formfield, "FormField")
}

func (formfielddate *FormFieldDate) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormFieldDate_stagedOrder, stage.FormFieldDates_referenceOrder, formfielddate, "FormFieldDate")
}

func (formfielddatetime *FormFieldDateTime) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormFieldDateTime_stagedOrder, stage.FormFieldDateTimes_referenceOrder, formfielddatetime, "FormFieldDateTime")
}

func (formfieldfloat64 *FormFieldFloat64) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormFieldFloat64_stagedOrder, stage.FormFieldFloat64s_referenceOrder, formfieldfloat64, "FormFieldFloat64")
}

func (formfieldint *FormFieldInt) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormFieldInt_stagedOrder, stage.FormFieldInts_referenceOrder, formfieldint, "FormFieldInt")
}

func (formfieldselect *FormFieldSelect) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormFieldSelect_stagedOrder, stage.FormFieldSelects_referenceOrder, formfieldselect, "FormFieldSelect")
}

func (formfieldstring *FormFieldString) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormFieldString_stagedOrder, stage.FormFieldStrings_referenceOrder, formfieldstring, "FormFieldString")
}

func (formfieldtime *FormFieldTime) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormFieldTime_stagedOrder, stage.FormFieldTimes_referenceOrder, formfieldtime, "FormFieldTime")
}

func (formgroup *FormGroup) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormGroup_stagedOrder, stage.FormGroups_referenceOrder, formgroup, "FormGroup")
}

func (formsortassocbutton *FormSortAssocButton) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FormSortAssocButton_stagedOrder, stage.FormSortAssocButtons_referenceOrder, formsortassocbutton, "FormSortAssocButton")
}

func (option *Option) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Option_stagedOrder, stage.Options_referenceOrder, option, "Option")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (checkbox *CheckBox) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(checkbox, checkbox.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkbox *CheckBox) GongGetReferenceIdentifier(stage *Stage) string {
	return checkbox.GongGetIdentifier(stage)
}

func (formdiv *FormDiv) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formdiv, formdiv.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formdiv *FormDiv) GongGetReferenceIdentifier(stage *Stage) string {
	return formdiv.GongGetIdentifier(stage)
}

func (formeditassocbutton *FormEditAssocButton) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formeditassocbutton, formeditassocbutton.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formeditassocbutton *FormEditAssocButton) GongGetReferenceIdentifier(stage *Stage) string {
	return formeditassocbutton.GongGetIdentifier(stage)
}

func (formfield *FormField) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formfield, formfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfield *FormField) GongGetReferenceIdentifier(stage *Stage) string {
	return formfield.GongGetIdentifier(stage)
}

func (formfielddate *FormFieldDate) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formfielddate, formfielddate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfielddate *FormFieldDate) GongGetReferenceIdentifier(stage *Stage) string {
	return formfielddate.GongGetIdentifier(stage)
}

func (formfielddatetime *FormFieldDateTime) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formfielddatetime, formfielddatetime.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfielddatetime *FormFieldDateTime) GongGetReferenceIdentifier(stage *Stage) string {
	return formfielddatetime.GongGetIdentifier(stage)
}

func (formfieldfloat64 *FormFieldFloat64) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formfieldfloat64, formfieldfloat64.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldfloat64 *FormFieldFloat64) GongGetReferenceIdentifier(stage *Stage) string {
	return formfieldfloat64.GongGetIdentifier(stage)
}

func (formfieldint *FormFieldInt) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formfieldint, formfieldint.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldint *FormFieldInt) GongGetReferenceIdentifier(stage *Stage) string {
	return formfieldint.GongGetIdentifier(stage)
}

func (formfieldselect *FormFieldSelect) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formfieldselect, formfieldselect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldselect *FormFieldSelect) GongGetReferenceIdentifier(stage *Stage) string {
	return formfieldselect.GongGetIdentifier(stage)
}

func (formfieldstring *FormFieldString) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formfieldstring, formfieldstring.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldstring *FormFieldString) GongGetReferenceIdentifier(stage *Stage) string {
	return formfieldstring.GongGetIdentifier(stage)
}

func (formfieldtime *FormFieldTime) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formfieldtime, formfieldtime.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formfieldtime *FormFieldTime) GongGetReferenceIdentifier(stage *Stage) string {
	return formfieldtime.GongGetIdentifier(stage)
}

func (formgroup *FormGroup) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formgroup, formgroup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formgroup *FormGroup) GongGetReferenceIdentifier(stage *Stage) string {
	return formgroup.GongGetIdentifier(stage)
}

func (formsortassocbutton *FormSortAssocButton) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formsortassocbutton, formsortassocbutton.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formsortassocbutton *FormSortAssocButton) GongGetReferenceIdentifier(stage *Stage) string {
	return formsortassocbutton.GongGetIdentifier(stage)
}

func (option *Option) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(option, option.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (option *Option) GongGetReferenceIdentifier(stage *Stage) string {
	return option.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (checkbox *CheckBox) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(checkbox.GongGetIdentifier(stage), "CheckBox", checkbox.Name)
}

func (formdiv *FormDiv) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formdiv.GongGetIdentifier(stage), "FormDiv", formdiv.Name)
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formeditassocbutton.GongGetIdentifier(stage), "FormEditAssocButton", formeditassocbutton.Name)
}

func (formfield *FormField) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formfield.GongGetIdentifier(stage), "FormField", formfield.Name)
}

func (formfielddate *FormFieldDate) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formfielddate.GongGetIdentifier(stage), "FormFieldDate", formfielddate.Name)
}

func (formfielddatetime *FormFieldDateTime) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formfielddatetime.GongGetIdentifier(stage), "FormFieldDateTime", formfielddatetime.Name)
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formfieldfloat64.GongGetIdentifier(stage), "FormFieldFloat64", formfieldfloat64.Name)
}

func (formfieldint *FormFieldInt) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formfieldint.GongGetIdentifier(stage), "FormFieldInt", formfieldint.Name)
}

func (formfieldselect *FormFieldSelect) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formfieldselect.GongGetIdentifier(stage), "FormFieldSelect", formfieldselect.Name)
}

func (formfieldstring *FormFieldString) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formfieldstring.GongGetIdentifier(stage), "FormFieldString", formfieldstring.Name)
}

func (formfieldtime *FormFieldTime) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formfieldtime.GongGetIdentifier(stage), "FormFieldTime", formfieldtime.Name)
}

func (formgroup *FormGroup) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formgroup.GongGetIdentifier(stage), "FormGroup", formgroup.Name)
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formsortassocbutton.GongGetIdentifier(stage), "FormSortAssocButton", formsortassocbutton.Name)
}

func (option *Option) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(option.GongGetIdentifier(stage), "Option", option.Name)
}

// insertion point for unstaging
func (checkbox *CheckBox) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(checkbox.GongGetReferenceIdentifier(stage))
}

func (formdiv *FormDiv) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formdiv.GongGetReferenceIdentifier(stage))
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formeditassocbutton.GongGetReferenceIdentifier(stage))
}

func (formfield *FormField) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formfield.GongGetReferenceIdentifier(stage))
}

func (formfielddate *FormFieldDate) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formfielddate.GongGetReferenceIdentifier(stage))
}

func (formfielddatetime *FormFieldDateTime) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formfielddatetime.GongGetReferenceIdentifier(stage))
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formfieldfloat64.GongGetReferenceIdentifier(stage))
}

func (formfieldint *FormFieldInt) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formfieldint.GongGetReferenceIdentifier(stage))
}

func (formfieldselect *FormFieldSelect) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formfieldselect.GongGetReferenceIdentifier(stage))
}

func (formfieldstring *FormFieldString) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formfieldstring.GongGetReferenceIdentifier(stage))
}

func (formfieldtime *FormFieldTime) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formfieldtime.GongGetReferenceIdentifier(stage))
}

func (formgroup *FormGroup) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formgroup.GongGetReferenceIdentifier(stage))
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formsortassocbutton.GongGetReferenceIdentifier(stage))
}

func (option *Option) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(option.GongGetReferenceIdentifier(stage))
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

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
