// generated code - do not edit
package basictypes

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
	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.CodeListTypes)

	res = __gong__appendInstances(res, stage.CodeOrNilReasonListTypes)

	res = __gong__appendInstances(res, stage.CodeTypes)

	res = __gong__appendInstances(res, stage.CodeWithAuthorityTypes)

	res = __gong__appendInstances(res, stage.CoordinatesTypes)

	res = __gong__appendInstances(res, stage.MeasureListTypes)

	res = __gong__appendInstances(res, stage.MeasureOrNilReasonListTypes)

	res = __gong__appendInstances(res, stage.MeasureTypes)

	return
}

// insertion point per named struct
func (codelisttype *CodeListType) GongCopy() GongstructIF {
	newInstance := new(CodeListType)
	codelisttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongCopy() GongstructIF {
	newInstance := new(CodeOrNilReasonListType)
	codeornilreasonlisttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetype *CodeType) GongCopy() GongstructIF {
	newInstance := new(CodeType)
	codetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codewithauthoritytype *CodeWithAuthorityType) GongCopy() GongstructIF {
	newInstance := new(CodeWithAuthorityType)
	codewithauthoritytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coordinatestype *CoordinatesType) GongCopy() GongstructIF {
	newInstance := new(CoordinatesType)
	coordinatestype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measurelisttype *MeasureListType) GongCopy() GongstructIF {
	newInstance := new(MeasureListType)
	measurelisttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongCopy() GongstructIF {
	newInstance := new(MeasureOrNilReasonListType)
	measureornilreasonlisttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measuretype *MeasureType) GongCopy() GongstructIF {
	newInstance := new(MeasureType)
	measuretype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (codelisttype *CodeListType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelisttype)
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeornilreasonlisttype)
}

func (codetype *CodeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetype)
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codewithauthoritytype)
}

func (coordinatestype *CoordinatesType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coordinatestype)
}

func (measurelisttype *MeasureListType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measurelisttype)
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measureornilreasonlisttype)
}

func (measuretype *MeasureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measuretype)
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
		stage.CodeListTypes,
		stage.CodeListType_stagedOrder,
		stage.CodeListTypes_reference,
		&stage.CodeListTypes_referenceOrder,
		stage.CodeListTypes_instance,
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
		stage.CodeOrNilReasonListTypes,
		stage.CodeOrNilReasonListType_stagedOrder,
		stage.CodeOrNilReasonListTypes_reference,
		&stage.CodeOrNilReasonListTypes_referenceOrder,
		stage.CodeOrNilReasonListTypes_instance,
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
		stage.CodeTypes,
		stage.CodeType_stagedOrder,
		stage.CodeTypes_reference,
		&stage.CodeTypes_referenceOrder,
		stage.CodeTypes_instance,
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
		stage.CodeWithAuthorityTypes,
		stage.CodeWithAuthorityType_stagedOrder,
		stage.CodeWithAuthorityTypes_reference,
		&stage.CodeWithAuthorityTypes_referenceOrder,
		stage.CodeWithAuthorityTypes_instance,
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
		stage.CoordinatesTypes,
		stage.CoordinatesType_stagedOrder,
		stage.CoordinatesTypes_reference,
		&stage.CoordinatesTypes_referenceOrder,
		stage.CoordinatesTypes_instance,
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
		stage.MeasureListTypes,
		stage.MeasureListType_stagedOrder,
		stage.MeasureListTypes_reference,
		&stage.MeasureListTypes_referenceOrder,
		stage.MeasureListTypes_instance,
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
		stage.MeasureOrNilReasonListTypes,
		stage.MeasureOrNilReasonListType_stagedOrder,
		stage.MeasureOrNilReasonListTypes_reference,
		&stage.MeasureOrNilReasonListTypes_referenceOrder,
		stage.MeasureOrNilReasonListTypes_instance,
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
		stage.MeasureTypes,
		stage.MeasureType_stagedOrder,
		stage.MeasureTypes_reference,
		&stage.MeasureTypes_referenceOrder,
		stage.MeasureTypes_instance,
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
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.CodeListTypes, &stage.CodeListTypes_reference, &stage.CodeListTypes_referenceOrder, &stage.CodeListTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOrNilReasonListTypes, &stage.CodeOrNilReasonListTypes_reference, &stage.CodeOrNilReasonListTypes_referenceOrder, &stage.CodeOrNilReasonListTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTypes, &stage.CodeTypes_reference, &stage.CodeTypes_referenceOrder, &stage.CodeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeWithAuthorityTypes, &stage.CodeWithAuthorityTypes_reference, &stage.CodeWithAuthorityTypes_referenceOrder, &stage.CodeWithAuthorityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CoordinatesTypes, &stage.CoordinatesTypes_reference, &stage.CoordinatesTypes_referenceOrder, &stage.CoordinatesTypes_instance)

	__gong__computeReferencePass1(stage, stage.MeasureListTypes, &stage.MeasureListTypes_reference, &stage.MeasureListTypes_referenceOrder, &stage.MeasureListTypes_instance)

	__gong__computeReferencePass1(stage, stage.MeasureOrNilReasonListTypes, &stage.MeasureOrNilReasonListTypes_reference, &stage.MeasureOrNilReasonListTypes_referenceOrder, &stage.MeasureOrNilReasonListTypes_instance)

	__gong__computeReferencePass1(stage, stage.MeasureTypes, &stage.MeasureTypes_reference, &stage.MeasureTypes_referenceOrder, &stage.MeasureTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CodeListTypes, stage.CodeListTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOrNilReasonListTypes, stage.CodeOrNilReasonListTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTypes, stage.CodeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeWithAuthorityTypes, stage.CodeWithAuthorityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CoordinatesTypes, stage.CoordinatesTypes_reference, stage)

	__gong__computeReferencePass2(stage.MeasureListTypes, stage.MeasureListTypes_reference, stage)

	__gong__computeReferencePass2(stage.MeasureOrNilReasonListTypes, stage.MeasureOrNilReasonListTypes_reference, stage)

	__gong__computeReferencePass2(stage.MeasureTypes, stage.MeasureTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (codelisttype *CodeListType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeListType_stagedOrder, stage.CodeListTypes_referenceOrder, codelisttype, "CodeListType")
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOrNilReasonListType_stagedOrder, stage.CodeOrNilReasonListTypes_referenceOrder, codeornilreasonlisttype, "CodeOrNilReasonListType")
}

func (codetype *CodeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeType_stagedOrder, stage.CodeTypes_referenceOrder, codetype, "CodeType")
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeWithAuthorityType_stagedOrder, stage.CodeWithAuthorityTypes_referenceOrder, codewithauthoritytype, "CodeWithAuthorityType")
}

func (coordinatestype *CoordinatesType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CoordinatesType_stagedOrder, stage.CoordinatesTypes_referenceOrder, coordinatestype, "CoordinatesType")
}

func (measurelisttype *MeasureListType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeasureListType_stagedOrder, stage.MeasureListTypes_referenceOrder, measurelisttype, "MeasureListType")
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeasureOrNilReasonListType_stagedOrder, stage.MeasureOrNilReasonListTypes_referenceOrder, measureornilreasonlisttype, "MeasureOrNilReasonListType")
}

func (measuretype *MeasureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MeasureType_stagedOrder, stage.MeasureTypes_referenceOrder, measuretype, "MeasureType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (codelisttype *CodeListType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelisttype, codelisttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelisttype *CodeListType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelisttype.GongGetIdentifier(stage)
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeornilreasonlisttype, codeornilreasonlisttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeornilreasonlisttype *CodeOrNilReasonListType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeornilreasonlisttype.GongGetIdentifier(stage)
}

func (codetype *CodeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetype, codetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetype *CodeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetype.GongGetIdentifier(stage)
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codewithauthoritytype, codewithauthoritytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codewithauthoritytype *CodeWithAuthorityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codewithauthoritytype.GongGetIdentifier(stage)
}

func (coordinatestype *CoordinatesType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coordinatestype, coordinatestype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coordinatestype *CoordinatesType) GongGetReferenceIdentifier(stage *Stage) string {
	return coordinatestype.GongGetIdentifier(stage)
}

func (measurelisttype *MeasureListType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measurelisttype, measurelisttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measurelisttype *MeasureListType) GongGetReferenceIdentifier(stage *Stage) string {
	return measurelisttype.GongGetIdentifier(stage)
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measureornilreasonlisttype, measureornilreasonlisttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongGetReferenceIdentifier(stage *Stage) string {
	return measureornilreasonlisttype.GongGetIdentifier(stage)
}

func (measuretype *MeasureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measuretype, measuretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measuretype *MeasureType) GongGetReferenceIdentifier(stage *Stage) string {
	return measuretype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (codelisttype *CodeListType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelisttype.GongGetIdentifier(stage), "CodeListType", codelisttype.Name)
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeornilreasonlisttype.GongGetIdentifier(stage), "CodeOrNilReasonListType", codeornilreasonlisttype.Name)
}

func (codetype *CodeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetype.GongGetIdentifier(stage), "CodeType", codetype.Name)
}

func (codewithauthoritytype *CodeWithAuthorityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codewithauthoritytype.GongGetIdentifier(stage), "CodeWithAuthorityType", codewithauthoritytype.Name)
}

func (coordinatestype *CoordinatesType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coordinatestype.GongGetIdentifier(stage), "CoordinatesType", coordinatestype.Name)
}

func (measurelisttype *MeasureListType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measurelisttype.GongGetIdentifier(stage), "MeasureListType", measurelisttype.Name)
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measureornilreasonlisttype.GongGetIdentifier(stage), "MeasureOrNilReasonListType", measureornilreasonlisttype.Name)
}

func (measuretype *MeasureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measuretype.GongGetIdentifier(stage), "MeasureType", measuretype.Name)
}

// insertion point for unstaging
func (codelisttype *CodeListType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelisttype.GongGetReferenceIdentifier(stage))
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeornilreasonlisttype.GongGetReferenceIdentifier(stage))
}

func (codetype *CodeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetype.GongGetReferenceIdentifier(stage))
}

func (codewithauthoritytype *CodeWithAuthorityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codewithauthoritytype.GongGetReferenceIdentifier(stage))
}

func (coordinatestype *CoordinatesType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coordinatestype.GongGetReferenceIdentifier(stage))
}

func (measurelisttype *MeasureListType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measurelisttype.GongGetReferenceIdentifier(stage))
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measureornilreasonlisttype.GongGetReferenceIdentifier(stage))
}

func (measuretype *MeasureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measuretype.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__basictypes__%s__%08d_", s.GongGetGongstructName(), order)
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
