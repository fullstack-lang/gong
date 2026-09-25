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
	// Compute reverse map for named struct GongEnum
	// insertion point per field
	stage.GongEnum_GongEnumValues_reverseMap = make(map[*GongEnumValue]*GongEnum)
	for gongenum := range stage.GongEnums {
		_ = gongenum
		for _, _gongenumvalue := range gongenum.GongEnumValues {
			stage.GongEnum_GongEnumValues_reverseMap[_gongenumvalue] = gongenum
		}
	}

	// Compute reverse map for named struct GongNote
	// insertion point per field
	stage.GongNote_Links_reverseMap = make(map[*GongLink]*GongNote)
	for gongnote := range stage.GongNotes {
		_ = gongnote
		for _, _gonglink := range gongnote.Links {
			stage.GongNote_Links_reverseMap[_gonglink] = gongnote
		}
	}

	// Compute reverse map for named struct GongStruct
	// insertion point per field
	stage.GongStruct_GongBasicFields_reverseMap = make(map[*GongBasicField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _gongbasicfield := range gongstruct.GongBasicFields {
			stage.GongStruct_GongBasicFields_reverseMap[_gongbasicfield] = gongstruct
		}
	}
	stage.GongStruct_GongTimeFields_reverseMap = make(map[*GongTimeField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _gongtimefield := range gongstruct.GongTimeFields {
			stage.GongStruct_GongTimeFields_reverseMap[_gongtimefield] = gongstruct
		}
	}
	stage.GongStruct_PointerToGongStructFields_reverseMap = make(map[*PointerToGongStructField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
			stage.GongStruct_PointerToGongStructFields_reverseMap[_pointertogongstructfield] = gongstruct
		}
	}
	stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap = make(map[*SliceOfPointerToGongStructField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
			stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap[_sliceofpointertogongstructfield] = gongstruct
		}
	}

	// Compute reverse map for named struct StageSetModel
	// insertion point per field
	stage.StageSetModel_Fields_reverseMap = make(map[*StageSetField]*StageSetModel)
	for stagesetmodel := range stage.StageSetModels {
		_ = stagesetmodel
		for _, _stagesetfield := range stagesetmodel.Fields {
			stage.StageSetModel_Fields_reverseMap[_stagesetfield] = stagesetmodel
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.GongBasicFields)

	res = __gong__appendInstances(res, stage.GongEnums)

	res = __gong__appendInstances(res, stage.GongEnumValues)

	res = __gong__appendInstances(res, stage.GongLinks)

	res = __gong__appendInstances(res, stage.GongNotes)

	res = __gong__appendInstances(res, stage.GongStructs)

	res = __gong__appendInstances(res, stage.GongTimeFields)

	res = __gong__appendInstances(res, stage.MetaReferences)

	res = __gong__appendInstances(res, stage.ModelPkgs)

	res = __gong__appendInstances(res, stage.PointerToGongStructFields)

	res = __gong__appendInstances(res, stage.SliceOfPointerToGongStructFields)

	res = __gong__appendInstances(res, stage.StageSetFields)

	res = __gong__appendInstances(res, stage.StageSetModels)

	return
}

// insertion point per named struct
func (gongbasicfield *GongBasicField) GongCopy() GongstructIF {
	newInstance := new(GongBasicField)
	gongbasicfield.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongenum *GongEnum) GongCopy() GongstructIF {
	newInstance := new(GongEnum)
	gongenum.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongenumvalue *GongEnumValue) GongCopy() GongstructIF {
	newInstance := new(GongEnumValue)
	gongenumvalue.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gonglink *GongLink) GongCopy() GongstructIF {
	newInstance := new(GongLink)
	gonglink.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongnote *GongNote) GongCopy() GongstructIF {
	newInstance := new(GongNote)
	gongnote.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongstruct *GongStruct) GongCopy() GongstructIF {
	newInstance := new(GongStruct)
	gongstruct.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongtimefield *GongTimeField) GongCopy() GongstructIF {
	newInstance := new(GongTimeField)
	gongtimefield.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metareference *MetaReference) GongCopy() GongstructIF {
	newInstance := new(MetaReference)
	metareference.GongCopyBasicFields(newInstance)
	return newInstance
}

func (modelpkg *ModelPkg) GongCopy() GongstructIF {
	newInstance := new(ModelPkg)
	modelpkg.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointertogongstructfield *PointerToGongStructField) GongCopy() GongstructIF {
	newInstance := new(PointerToGongStructField)
	pointertogongstructfield.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongCopy() GongstructIF {
	newInstance := new(SliceOfPointerToGongStructField)
	sliceofpointertogongstructfield.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stagesetfield *StageSetField) GongCopy() GongstructIF {
	newInstance := new(StageSetField)
	stagesetfield.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stagesetmodel *StageSetModel) GongCopy() GongstructIF {
	newInstance := new(StageSetModel)
	stagesetmodel.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (gongbasicfield *GongBasicField) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongbasicfield)
}

func (gongenum *GongEnum) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongenum)
}

func (gongenumvalue *GongEnumValue) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongenumvalue)
}

func (gonglink *GongLink) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gonglink)
}

func (gongnote *GongNote) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongnote)
}

func (gongstruct *GongStruct) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongstruct)
}

func (gongtimefield *GongTimeField) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongtimefield)
}

func (metareference *MetaReference) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metareference)
}

func (modelpkg *ModelPkg) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, modelpkg)
}

func (pointertogongstructfield *PointerToGongStructField) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointertogongstructfield)
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sliceofpointertogongstructfield)
}

func (stagesetfield *StageSetField) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stagesetfield)
}

func (stagesetmodel *StageSetModel) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stagesetmodel)
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
		stage.GongBasicFields,
		stage.GongBasicField_stagedOrder,
		stage.GongBasicFields_reference,
		&stage.GongBasicFields_referenceOrder,
		stage.GongBasicFields_instance,
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
		stage.GongEnums,
		stage.GongEnum_stagedOrder,
		stage.GongEnums_reference,
		&stage.GongEnums_referenceOrder,
		stage.GongEnums_instance,
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
		stage.GongEnumValues,
		stage.GongEnumValue_stagedOrder,
		stage.GongEnumValues_reference,
		&stage.GongEnumValues_referenceOrder,
		stage.GongEnumValues_instance,
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
		stage.GongLinks,
		stage.GongLink_stagedOrder,
		stage.GongLinks_reference,
		&stage.GongLinks_referenceOrder,
		stage.GongLinks_instance,
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
		stage.GongNotes,
		stage.GongNote_stagedOrder,
		stage.GongNotes_reference,
		&stage.GongNotes_referenceOrder,
		stage.GongNotes_instance,
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
		stage.GongStructs,
		stage.GongStruct_stagedOrder,
		stage.GongStructs_reference,
		&stage.GongStructs_referenceOrder,
		stage.GongStructs_instance,
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
		stage.GongTimeFields,
		stage.GongTimeField_stagedOrder,
		stage.GongTimeFields_reference,
		&stage.GongTimeFields_referenceOrder,
		stage.GongTimeFields_instance,
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
		stage.MetaReferences,
		stage.MetaReference_stagedOrder,
		stage.MetaReferences_reference,
		&stage.MetaReferences_referenceOrder,
		stage.MetaReferences_instance,
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
		stage.ModelPkgs,
		stage.ModelPkg_stagedOrder,
		stage.ModelPkgs_reference,
		&stage.ModelPkgs_referenceOrder,
		stage.ModelPkgs_instance,
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
		stage.PointerToGongStructFields,
		stage.PointerToGongStructField_stagedOrder,
		stage.PointerToGongStructFields_reference,
		&stage.PointerToGongStructFields_referenceOrder,
		stage.PointerToGongStructFields_instance,
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
		stage.SliceOfPointerToGongStructFields,
		stage.SliceOfPointerToGongStructField_stagedOrder,
		stage.SliceOfPointerToGongStructFields_reference,
		&stage.SliceOfPointerToGongStructFields_referenceOrder,
		stage.SliceOfPointerToGongStructFields_instance,
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
		stage.StageSetFields,
		stage.StageSetField_stagedOrder,
		stage.StageSetFields_reference,
		&stage.StageSetFields_referenceOrder,
		stage.StageSetFields_instance,
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
		stage.StageSetModels,
		stage.StageSetModel_stagedOrder,
		stage.StageSetModels_reference,
		&stage.StageSetModels_referenceOrder,
		stage.StageSetModels_instance,
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
	__gong__computeReferencePass1(stage, stage.GongBasicFields, &stage.GongBasicFields_reference, &stage.GongBasicFields_referenceOrder, &stage.GongBasicFields_instance)

	__gong__computeReferencePass1(stage, stage.GongEnums, &stage.GongEnums_reference, &stage.GongEnums_referenceOrder, &stage.GongEnums_instance)

	__gong__computeReferencePass1(stage, stage.GongEnumValues, &stage.GongEnumValues_reference, &stage.GongEnumValues_referenceOrder, &stage.GongEnumValues_instance)

	__gong__computeReferencePass1(stage, stage.GongLinks, &stage.GongLinks_reference, &stage.GongLinks_referenceOrder, &stage.GongLinks_instance)

	__gong__computeReferencePass1(stage, stage.GongNotes, &stage.GongNotes_reference, &stage.GongNotes_referenceOrder, &stage.GongNotes_instance)

	__gong__computeReferencePass1(stage, stage.GongStructs, &stage.GongStructs_reference, &stage.GongStructs_referenceOrder, &stage.GongStructs_instance)

	__gong__computeReferencePass1(stage, stage.GongTimeFields, &stage.GongTimeFields_reference, &stage.GongTimeFields_referenceOrder, &stage.GongTimeFields_instance)

	__gong__computeReferencePass1(stage, stage.MetaReferences, &stage.MetaReferences_reference, &stage.MetaReferences_referenceOrder, &stage.MetaReferences_instance)

	__gong__computeReferencePass1(stage, stage.ModelPkgs, &stage.ModelPkgs_reference, &stage.ModelPkgs_referenceOrder, &stage.ModelPkgs_instance)

	__gong__computeReferencePass1(stage, stage.PointerToGongStructFields, &stage.PointerToGongStructFields_reference, &stage.PointerToGongStructFields_referenceOrder, &stage.PointerToGongStructFields_instance)

	__gong__computeReferencePass1(stage, stage.SliceOfPointerToGongStructFields, &stage.SliceOfPointerToGongStructFields_reference, &stage.SliceOfPointerToGongStructFields_referenceOrder, &stage.SliceOfPointerToGongStructFields_instance)

	__gong__computeReferencePass1(stage, stage.StageSetFields, &stage.StageSetFields_reference, &stage.StageSetFields_referenceOrder, &stage.StageSetFields_instance)

	__gong__computeReferencePass1(stage, stage.StageSetModels, &stage.StageSetModels_reference, &stage.StageSetModels_referenceOrder, &stage.StageSetModels_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.GongBasicFields, stage.GongBasicFields_reference, stage)

	__gong__computeReferencePass2(stage.GongEnums, stage.GongEnums_reference, stage)

	__gong__computeReferencePass2(stage.GongEnumValues, stage.GongEnumValues_reference, stage)

	__gong__computeReferencePass2(stage.GongLinks, stage.GongLinks_reference, stage)

	__gong__computeReferencePass2(stage.GongNotes, stage.GongNotes_reference, stage)

	__gong__computeReferencePass2(stage.GongStructs, stage.GongStructs_reference, stage)

	__gong__computeReferencePass2(stage.GongTimeFields, stage.GongTimeFields_reference, stage)

	__gong__computeReferencePass2(stage.MetaReferences, stage.MetaReferences_reference, stage)

	__gong__computeReferencePass2(stage.ModelPkgs, stage.ModelPkgs_reference, stage)

	__gong__computeReferencePass2(stage.PointerToGongStructFields, stage.PointerToGongStructFields_reference, stage)

	__gong__computeReferencePass2(stage.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructFields_reference, stage)

	__gong__computeReferencePass2(stage.StageSetFields, stage.StageSetFields_reference, stage)

	__gong__computeReferencePass2(stage.StageSetModels, stage.StageSetModels_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (gongbasicfield *GongBasicField) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongBasicField_stagedOrder, stage.GongBasicFields_referenceOrder, gongbasicfield, "GongBasicField")
}

func (gongenum *GongEnum) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongEnum_stagedOrder, stage.GongEnums_referenceOrder, gongenum, "GongEnum")
}

func (gongenumvalue *GongEnumValue) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongEnumValue_stagedOrder, stage.GongEnumValues_referenceOrder, gongenumvalue, "GongEnumValue")
}

func (gonglink *GongLink) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongLink_stagedOrder, stage.GongLinks_referenceOrder, gonglink, "GongLink")
}

func (gongnote *GongNote) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongNote_stagedOrder, stage.GongNotes_referenceOrder, gongnote, "GongNote")
}

func (gongstruct *GongStruct) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongStruct_stagedOrder, stage.GongStructs_referenceOrder, gongstruct, "GongStruct")
}

func (gongtimefield *GongTimeField) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongTimeField_stagedOrder, stage.GongTimeFields_referenceOrder, gongtimefield, "GongTimeField")
}

func (metareference *MetaReference) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MetaReference_stagedOrder, stage.MetaReferences_referenceOrder, metareference, "MetaReference")
}

func (modelpkg *ModelPkg) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ModelPkg_stagedOrder, stage.ModelPkgs_referenceOrder, modelpkg, "ModelPkg")
}

func (pointertogongstructfield *PointerToGongStructField) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointerToGongStructField_stagedOrder, stage.PointerToGongStructFields_referenceOrder, pointertogongstructfield, "PointerToGongStructField")
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SliceOfPointerToGongStructField_stagedOrder, stage.SliceOfPointerToGongStructFields_referenceOrder, sliceofpointertogongstructfield, "SliceOfPointerToGongStructField")
}

func (stagesetfield *StageSetField) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StageSetField_stagedOrder, stage.StageSetFields_referenceOrder, stagesetfield, "StageSetField")
}

func (stagesetmodel *StageSetModel) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.StageSetModel_stagedOrder, stage.StageSetModels_referenceOrder, stagesetmodel, "StageSetModel")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (gongbasicfield *GongBasicField) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongbasicfield, gongbasicfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongbasicfield *GongBasicField) GongGetReferenceIdentifier(stage *Stage) string {
	return gongbasicfield.GongGetIdentifier(stage)
}

func (gongenum *GongEnum) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongenum, gongenum.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenum *GongEnum) GongGetReferenceIdentifier(stage *Stage) string {
	return gongenum.GongGetIdentifier(stage)
}

func (gongenumvalue *GongEnumValue) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongenumvalue, gongenumvalue.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumvalue *GongEnumValue) GongGetReferenceIdentifier(stage *Stage) string {
	return gongenumvalue.GongGetIdentifier(stage)
}

func (gonglink *GongLink) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gonglink, gonglink.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gonglink *GongLink) GongGetReferenceIdentifier(stage *Stage) string {
	return gonglink.GongGetIdentifier(stage)
}

func (gongnote *GongNote) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongnote, gongnote.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnote *GongNote) GongGetReferenceIdentifier(stage *Stage) string {
	return gongnote.GongGetIdentifier(stage)
}

func (gongstruct *GongStruct) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongstruct, gongstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongstruct *GongStruct) GongGetReferenceIdentifier(stage *Stage) string {
	return gongstruct.GongGetIdentifier(stage)
}

func (gongtimefield *GongTimeField) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongtimefield, gongtimefield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongtimefield *GongTimeField) GongGetReferenceIdentifier(stage *Stage) string {
	return gongtimefield.GongGetIdentifier(stage)
}

func (metareference *MetaReference) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metareference, metareference.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metareference *MetaReference) GongGetReferenceIdentifier(stage *Stage) string {
	return metareference.GongGetIdentifier(stage)
}

func (modelpkg *ModelPkg) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(modelpkg, modelpkg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (modelpkg *ModelPkg) GongGetReferenceIdentifier(stage *Stage) string {
	return modelpkg.GongGetIdentifier(stage)
}

func (pointertogongstructfield *PointerToGongStructField) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointertogongstructfield, pointertogongstructfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointertogongstructfield *PointerToGongStructField) GongGetReferenceIdentifier(stage *Stage) string {
	return pointertogongstructfield.GongGetIdentifier(stage)
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sliceofpointertogongstructfield, sliceofpointertogongstructfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetReferenceIdentifier(stage *Stage) string {
	return sliceofpointertogongstructfield.GongGetIdentifier(stage)
}

func (stagesetfield *StageSetField) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stagesetfield, stagesetfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stagesetfield *StageSetField) GongGetReferenceIdentifier(stage *Stage) string {
	return stagesetfield.GongGetIdentifier(stage)
}

func (stagesetmodel *StageSetModel) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stagesetmodel, stagesetmodel.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stagesetmodel *StageSetModel) GongGetReferenceIdentifier(stage *Stage) string {
	return stagesetmodel.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (gongbasicfield *GongBasicField) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongbasicfield.GongGetIdentifier(stage), "GongBasicField", gongbasicfield.Name)
}

func (gongenum *GongEnum) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongenum.GongGetIdentifier(stage), "GongEnum", gongenum.Name)
}

func (gongenumvalue *GongEnumValue) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongenumvalue.GongGetIdentifier(stage), "GongEnumValue", gongenumvalue.Name)
}

func (gonglink *GongLink) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gonglink.GongGetIdentifier(stage), "GongLink", gonglink.Name)
}

func (gongnote *GongNote) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongnote.GongGetIdentifier(stage), "GongNote", gongnote.Name)
}

func (gongstruct *GongStruct) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongstruct.GongGetIdentifier(stage), "GongStruct", gongstruct.Name)
}

func (gongtimefield *GongTimeField) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongtimefield.GongGetIdentifier(stage), "GongTimeField", gongtimefield.Name)
}

func (metareference *MetaReference) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metareference.GongGetIdentifier(stage), "MetaReference", metareference.Name)
}

func (modelpkg *ModelPkg) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(modelpkg.GongGetIdentifier(stage), "ModelPkg", modelpkg.Name)
}

func (pointertogongstructfield *PointerToGongStructField) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointertogongstructfield.GongGetIdentifier(stage), "PointerToGongStructField", pointertogongstructfield.Name)
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sliceofpointertogongstructfield.GongGetIdentifier(stage), "SliceOfPointerToGongStructField", sliceofpointertogongstructfield.Name)
}

func (stagesetfield *StageSetField) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stagesetfield.GongGetIdentifier(stage), "StageSetField", stagesetfield.Name)
}

func (stagesetmodel *StageSetModel) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stagesetmodel.GongGetIdentifier(stage), "StageSetModel", stagesetmodel.Name)
}

// insertion point for unstaging
func (gongbasicfield *GongBasicField) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongbasicfield.GongGetReferenceIdentifier(stage))
}

func (gongenum *GongEnum) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongenum.GongGetReferenceIdentifier(stage))
}

func (gongenumvalue *GongEnumValue) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongenumvalue.GongGetReferenceIdentifier(stage))
}

func (gonglink *GongLink) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gonglink.GongGetReferenceIdentifier(stage))
}

func (gongnote *GongNote) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongnote.GongGetReferenceIdentifier(stage))
}

func (gongstruct *GongStruct) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongstruct.GongGetReferenceIdentifier(stage))
}

func (gongtimefield *GongTimeField) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongtimefield.GongGetReferenceIdentifier(stage))
}

func (metareference *MetaReference) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metareference.GongGetReferenceIdentifier(stage))
}

func (modelpkg *ModelPkg) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(modelpkg.GongGetReferenceIdentifier(stage))
}

func (pointertogongstructfield *PointerToGongStructField) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointertogongstructfield.GongGetReferenceIdentifier(stage))
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sliceofpointertogongstructfield.GongGetReferenceIdentifier(stage))
}

func (stagesetfield *StageSetField) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stagesetfield.GongGetReferenceIdentifier(stage))
}

func (stagesetmodel *StageSetModel) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stagesetmodel.GongGetReferenceIdentifier(stage))
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
