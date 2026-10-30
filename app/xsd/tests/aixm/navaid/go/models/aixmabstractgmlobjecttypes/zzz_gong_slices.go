// generated code - do not edit
package aixmabstractgmlobjecttypes

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
	res = __gong__appendInstances(res, stage.AbstractAIXMFeatureBaseTypes)

	res = __gong__appendInstances(res, stage.AbstractAIXMFeatureTypes)

	res = __gong__appendInstances(res, stage.AbstractAIXMMessageBaseTypes)

	res = __gong__appendInstances(res, stage.AbstractAIXMMessageTypes)

	res = __gong__appendInstances(res, stage.AbstractAIXMObjectTypes)

	res = __gong__appendInstances(res, stage.AbstractAIXMPropertyTypes)

	res = __gong__appendInstances(res, stage.AbstractAIXMTimeSliceBaseTypes)

	res = __gong__appendInstances(res, stage.AbstractAIXMTimeSliceTypes)

	res = __gong__appendInstances(res, stage.AbstractExtensionTypes)

	res = __gong__appendInstances(res, stage.FeatureMetadataPropertyTypes)

	res = __gong__appendInstances(res, stage.FeatureTimeSliceMetadataPropertyTypes)

	res = __gong__appendInstances(res, stage.MessageMetadataPropertyTypes)

	return
}

// insertion point per named struct
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongCopy() GongstructIF {
	newInstance := new(AbstractAIXMFeatureBaseType)
	abstractaixmfeaturebasetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongCopy() GongstructIF {
	newInstance := new(AbstractAIXMFeatureType)
	abstractaixmfeaturetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongCopy() GongstructIF {
	newInstance := new(AbstractAIXMMessageBaseType)
	abstractaixmmessagebasetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongCopy() GongstructIF {
	newInstance := new(AbstractAIXMMessageType)
	abstractaixmmessagetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongCopy() GongstructIF {
	newInstance := new(AbstractAIXMObjectType)
	abstractaixmobjecttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongCopy() GongstructIF {
	newInstance := new(AbstractAIXMPropertyType)
	abstractaixmpropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongCopy() GongstructIF {
	newInstance := new(AbstractAIXMTimeSliceBaseType)
	abstractaixmtimeslicebasetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AbstractAIXMTimeSliceType)
	abstractaixmtimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractextensiontype *AbstractExtensionType) GongCopy() GongstructIF {
	newInstance := new(AbstractExtensionType)
	abstractextensiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongCopy() GongstructIF {
	newInstance := new(FeatureMetadataPropertyType)
	featuremetadatapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongCopy() GongstructIF {
	newInstance := new(FeatureTimeSliceMetadataPropertyType)
	featuretimeslicemetadatapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongCopy() GongstructIF {
	newInstance := new(MessageMetadataPropertyType)
	messagemetadatapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractaixmfeaturebasetype)
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractaixmfeaturetype)
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractaixmmessagebasetype)
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractaixmmessagetype)
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractaixmobjecttype)
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractaixmpropertytype)
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractaixmtimeslicebasetype)
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractaixmtimeslicetype)
}

func (abstractextensiontype *AbstractExtensionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractextensiontype)
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, featuremetadatapropertytype)
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, featuretimeslicemetadatapropertytype)
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, messagemetadatapropertytype)
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
		stage.AbstractAIXMFeatureBaseTypes,
		stage.AbstractAIXMFeatureBaseType_stagedOrder,
		stage.AbstractAIXMFeatureBaseTypes_reference,
		&stage.AbstractAIXMFeatureBaseTypes_referenceOrder,
		stage.AbstractAIXMFeatureBaseTypes_instance,
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
		stage.AbstractAIXMFeatureTypes,
		stage.AbstractAIXMFeatureType_stagedOrder,
		stage.AbstractAIXMFeatureTypes_reference,
		&stage.AbstractAIXMFeatureTypes_referenceOrder,
		stage.AbstractAIXMFeatureTypes_instance,
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
		stage.AbstractAIXMMessageBaseTypes,
		stage.AbstractAIXMMessageBaseType_stagedOrder,
		stage.AbstractAIXMMessageBaseTypes_reference,
		&stage.AbstractAIXMMessageBaseTypes_referenceOrder,
		stage.AbstractAIXMMessageBaseTypes_instance,
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
		stage.AbstractAIXMMessageTypes,
		stage.AbstractAIXMMessageType_stagedOrder,
		stage.AbstractAIXMMessageTypes_reference,
		&stage.AbstractAIXMMessageTypes_referenceOrder,
		stage.AbstractAIXMMessageTypes_instance,
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
		stage.AbstractAIXMObjectTypes,
		stage.AbstractAIXMObjectType_stagedOrder,
		stage.AbstractAIXMObjectTypes_reference,
		&stage.AbstractAIXMObjectTypes_referenceOrder,
		stage.AbstractAIXMObjectTypes_instance,
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
		stage.AbstractAIXMPropertyTypes,
		stage.AbstractAIXMPropertyType_stagedOrder,
		stage.AbstractAIXMPropertyTypes_reference,
		&stage.AbstractAIXMPropertyTypes_referenceOrder,
		stage.AbstractAIXMPropertyTypes_instance,
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
		stage.AbstractAIXMTimeSliceBaseTypes,
		stage.AbstractAIXMTimeSliceBaseType_stagedOrder,
		stage.AbstractAIXMTimeSliceBaseTypes_reference,
		&stage.AbstractAIXMTimeSliceBaseTypes_referenceOrder,
		stage.AbstractAIXMTimeSliceBaseTypes_instance,
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
		stage.AbstractAIXMTimeSliceTypes,
		stage.AbstractAIXMTimeSliceType_stagedOrder,
		stage.AbstractAIXMTimeSliceTypes_reference,
		&stage.AbstractAIXMTimeSliceTypes_referenceOrder,
		stage.AbstractAIXMTimeSliceTypes_instance,
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
		stage.AbstractExtensionTypes,
		stage.AbstractExtensionType_stagedOrder,
		stage.AbstractExtensionTypes_reference,
		&stage.AbstractExtensionTypes_referenceOrder,
		stage.AbstractExtensionTypes_instance,
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
		stage.FeatureMetadataPropertyTypes,
		stage.FeatureMetadataPropertyType_stagedOrder,
		stage.FeatureMetadataPropertyTypes_reference,
		&stage.FeatureMetadataPropertyTypes_referenceOrder,
		stage.FeatureMetadataPropertyTypes_instance,
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
		stage.FeatureTimeSliceMetadataPropertyTypes,
		stage.FeatureTimeSliceMetadataPropertyType_stagedOrder,
		stage.FeatureTimeSliceMetadataPropertyTypes_reference,
		&stage.FeatureTimeSliceMetadataPropertyTypes_referenceOrder,
		stage.FeatureTimeSliceMetadataPropertyTypes_instance,
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
		stage.MessageMetadataPropertyTypes,
		stage.MessageMetadataPropertyType_stagedOrder,
		stage.MessageMetadataPropertyTypes_reference,
		&stage.MessageMetadataPropertyTypes_referenceOrder,
		stage.MessageMetadataPropertyTypes_instance,
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
	__gong__computeReferencePass1(stage, stage.AbstractAIXMFeatureBaseTypes, &stage.AbstractAIXMFeatureBaseTypes_reference, &stage.AbstractAIXMFeatureBaseTypes_referenceOrder, &stage.AbstractAIXMFeatureBaseTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractAIXMFeatureTypes, &stage.AbstractAIXMFeatureTypes_reference, &stage.AbstractAIXMFeatureTypes_referenceOrder, &stage.AbstractAIXMFeatureTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractAIXMMessageBaseTypes, &stage.AbstractAIXMMessageBaseTypes_reference, &stage.AbstractAIXMMessageBaseTypes_referenceOrder, &stage.AbstractAIXMMessageBaseTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractAIXMMessageTypes, &stage.AbstractAIXMMessageTypes_reference, &stage.AbstractAIXMMessageTypes_referenceOrder, &stage.AbstractAIXMMessageTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractAIXMObjectTypes, &stage.AbstractAIXMObjectTypes_reference, &stage.AbstractAIXMObjectTypes_referenceOrder, &stage.AbstractAIXMObjectTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractAIXMPropertyTypes, &stage.AbstractAIXMPropertyTypes_reference, &stage.AbstractAIXMPropertyTypes_referenceOrder, &stage.AbstractAIXMPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractAIXMTimeSliceBaseTypes, &stage.AbstractAIXMTimeSliceBaseTypes_reference, &stage.AbstractAIXMTimeSliceBaseTypes_referenceOrder, &stage.AbstractAIXMTimeSliceBaseTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractAIXMTimeSliceTypes, &stage.AbstractAIXMTimeSliceTypes_reference, &stage.AbstractAIXMTimeSliceTypes_referenceOrder, &stage.AbstractAIXMTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractExtensionTypes, &stage.AbstractExtensionTypes_reference, &stage.AbstractExtensionTypes_referenceOrder, &stage.AbstractExtensionTypes_instance)

	__gong__computeReferencePass1(stage, stage.FeatureMetadataPropertyTypes, &stage.FeatureMetadataPropertyTypes_reference, &stage.FeatureMetadataPropertyTypes_referenceOrder, &stage.FeatureMetadataPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.FeatureTimeSliceMetadataPropertyTypes, &stage.FeatureTimeSliceMetadataPropertyTypes_reference, &stage.FeatureTimeSliceMetadataPropertyTypes_referenceOrder, &stage.FeatureTimeSliceMetadataPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MessageMetadataPropertyTypes, &stage.MessageMetadataPropertyTypes_reference, &stage.MessageMetadataPropertyTypes_referenceOrder, &stage.MessageMetadataPropertyTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AbstractAIXMFeatureBaseTypes, stage.AbstractAIXMFeatureBaseTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractAIXMFeatureTypes, stage.AbstractAIXMFeatureTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractAIXMMessageBaseTypes, stage.AbstractAIXMMessageBaseTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractAIXMMessageTypes, stage.AbstractAIXMMessageTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractAIXMObjectTypes, stage.AbstractAIXMObjectTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractAIXMPropertyTypes, stage.AbstractAIXMPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractAIXMTimeSliceBaseTypes, stage.AbstractAIXMTimeSliceBaseTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractAIXMTimeSliceTypes, stage.AbstractAIXMTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractExtensionTypes, stage.AbstractExtensionTypes_reference, stage)

	__gong__computeReferencePass2(stage.FeatureMetadataPropertyTypes, stage.FeatureMetadataPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.FeatureTimeSliceMetadataPropertyTypes, stage.FeatureTimeSliceMetadataPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MessageMetadataPropertyTypes, stage.MessageMetadataPropertyTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAIXMFeatureBaseType_stagedOrder, stage.AbstractAIXMFeatureBaseTypes_referenceOrder, abstractaixmfeaturebasetype, "AbstractAIXMFeatureBaseType")
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAIXMFeatureType_stagedOrder, stage.AbstractAIXMFeatureTypes_referenceOrder, abstractaixmfeaturetype, "AbstractAIXMFeatureType")
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAIXMMessageBaseType_stagedOrder, stage.AbstractAIXMMessageBaseTypes_referenceOrder, abstractaixmmessagebasetype, "AbstractAIXMMessageBaseType")
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAIXMMessageType_stagedOrder, stage.AbstractAIXMMessageTypes_referenceOrder, abstractaixmmessagetype, "AbstractAIXMMessageType")
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAIXMObjectType_stagedOrder, stage.AbstractAIXMObjectTypes_referenceOrder, abstractaixmobjecttype, "AbstractAIXMObjectType")
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAIXMPropertyType_stagedOrder, stage.AbstractAIXMPropertyTypes_referenceOrder, abstractaixmpropertytype, "AbstractAIXMPropertyType")
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAIXMTimeSliceBaseType_stagedOrder, stage.AbstractAIXMTimeSliceBaseTypes_referenceOrder, abstractaixmtimeslicebasetype, "AbstractAIXMTimeSliceBaseType")
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractAIXMTimeSliceType_stagedOrder, stage.AbstractAIXMTimeSliceTypes_referenceOrder, abstractaixmtimeslicetype, "AbstractAIXMTimeSliceType")
}

func (abstractextensiontype *AbstractExtensionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractExtensionType_stagedOrder, stage.AbstractExtensionTypes_referenceOrder, abstractextensiontype, "AbstractExtensionType")
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FeatureMetadataPropertyType_stagedOrder, stage.FeatureMetadataPropertyTypes_referenceOrder, featuremetadatapropertytype, "FeatureMetadataPropertyType")
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FeatureTimeSliceMetadataPropertyType_stagedOrder, stage.FeatureTimeSliceMetadataPropertyTypes_referenceOrder, featuretimeslicemetadatapropertytype, "FeatureTimeSliceMetadataPropertyType")
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MessageMetadataPropertyType_stagedOrder, stage.MessageMetadataPropertyTypes_referenceOrder, messagemetadatapropertytype, "MessageMetadataPropertyType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractaixmfeaturebasetype, abstractaixmfeaturebasetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractaixmfeaturebasetype.GongGetIdentifier(stage)
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractaixmfeaturetype, abstractaixmfeaturetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractaixmfeaturetype.GongGetIdentifier(stage)
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractaixmmessagebasetype, abstractaixmmessagebasetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractaixmmessagebasetype.GongGetIdentifier(stage)
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractaixmmessagetype, abstractaixmmessagetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractaixmmessagetype *AbstractAIXMMessageType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractaixmmessagetype.GongGetIdentifier(stage)
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractaixmobjecttype, abstractaixmobjecttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractaixmobjecttype *AbstractAIXMObjectType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractaixmobjecttype.GongGetIdentifier(stage)
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractaixmpropertytype, abstractaixmpropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractaixmpropertytype.GongGetIdentifier(stage)
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractaixmtimeslicebasetype, abstractaixmtimeslicebasetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractaixmtimeslicebasetype.GongGetIdentifier(stage)
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractaixmtimeslicetype, abstractaixmtimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractaixmtimeslicetype.GongGetIdentifier(stage)
}

func (abstractextensiontype *AbstractExtensionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractextensiontype, abstractextensiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractextensiontype *AbstractExtensionType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractextensiontype.GongGetIdentifier(stage)
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(featuremetadatapropertytype, featuremetadatapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return featuremetadatapropertytype.GongGetIdentifier(stage)
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(featuretimeslicemetadatapropertytype, featuretimeslicemetadatapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return featuretimeslicemetadatapropertytype.GongGetIdentifier(stage)
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(messagemetadatapropertytype, messagemetadatapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (messagemetadatapropertytype *MessageMetadataPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return messagemetadatapropertytype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractaixmfeaturebasetype.GongGetIdentifier(stage), "AbstractAIXMFeatureBaseType", abstractaixmfeaturebasetype.Name)
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractaixmfeaturetype.GongGetIdentifier(stage), "AbstractAIXMFeatureType", abstractaixmfeaturetype.Name)
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractaixmmessagebasetype.GongGetIdentifier(stage), "AbstractAIXMMessageBaseType", abstractaixmmessagebasetype.Name)
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractaixmmessagetype.GongGetIdentifier(stage), "AbstractAIXMMessageType", abstractaixmmessagetype.Name)
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractaixmobjecttype.GongGetIdentifier(stage), "AbstractAIXMObjectType", abstractaixmobjecttype.Name)
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractaixmpropertytype.GongGetIdentifier(stage), "AbstractAIXMPropertyType", abstractaixmpropertytype.Name)
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractaixmtimeslicebasetype.GongGetIdentifier(stage), "AbstractAIXMTimeSliceBaseType", abstractaixmtimeslicebasetype.Name)
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractaixmtimeslicetype.GongGetIdentifier(stage), "AbstractAIXMTimeSliceType", abstractaixmtimeslicetype.Name)
}

func (abstractextensiontype *AbstractExtensionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractextensiontype.GongGetIdentifier(stage), "AbstractExtensionType", abstractextensiontype.Name)
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(featuremetadatapropertytype.GongGetIdentifier(stage), "FeatureMetadataPropertyType", featuremetadatapropertytype.Name)
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(featuretimeslicemetadatapropertytype.GongGetIdentifier(stage), "FeatureTimeSliceMetadataPropertyType", featuretimeslicemetadatapropertytype.Name)
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(messagemetadatapropertytype.GongGetIdentifier(stage), "MessageMetadataPropertyType", messagemetadatapropertytype.Name)
}

// insertion point for unstaging
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractaixmfeaturebasetype.GongGetReferenceIdentifier(stage))
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractaixmfeaturetype.GongGetReferenceIdentifier(stage))
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractaixmmessagebasetype.GongGetReferenceIdentifier(stage))
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractaixmmessagetype.GongGetReferenceIdentifier(stage))
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractaixmobjecttype.GongGetReferenceIdentifier(stage))
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractaixmpropertytype.GongGetReferenceIdentifier(stage))
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractaixmtimeslicebasetype.GongGetReferenceIdentifier(stage))
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractaixmtimeslicetype.GongGetReferenceIdentifier(stage))
}

func (abstractextensiontype *AbstractExtensionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractextensiontype.GongGetReferenceIdentifier(stage))
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(featuremetadatapropertytype.GongGetReferenceIdentifier(stage))
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(featuretimeslicemetadatapropertytype.GongGetReferenceIdentifier(stage))
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(messagemetadatapropertytype.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__aixmabstractgmlobjecttypes__%s__%08d_", s.GongGetGongstructName(), order)
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
