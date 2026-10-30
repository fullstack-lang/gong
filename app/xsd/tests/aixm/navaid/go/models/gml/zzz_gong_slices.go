// generated code - do not edit
package gml

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
	// Compute reverse map for named struct AbstractFeatureType
	// insertion point per field
	stage.AbstractFeatureType_NameXSD_reverseMap = make(map[*CodeType]*AbstractFeatureType)
	for abstractfeaturetype := range stage.AbstractFeatureTypes {
		_ = abstractfeaturetype
		for _, _codetype := range abstractfeaturetype.NameXSD {
			stage.AbstractFeatureType_NameXSD_reverseMap[_codetype] = abstractfeaturetype
		}
	}

	// Compute reverse map for named struct AbstractGMLType
	// insertion point per field
	stage.AbstractGMLType_NameXSD_reverseMap = make(map[*CodeType]*AbstractGMLType)
	for abstractgmltype := range stage.AbstractGMLTypes {
		_ = abstractgmltype
		for _, _codetype := range abstractgmltype.NameXSD {
			stage.AbstractGMLType_NameXSD_reverseMap[_codetype] = abstractgmltype
		}
	}

	// Compute reverse map for named struct AbstractGeometricPrimitiveType
	// insertion point per field
	stage.AbstractGeometricPrimitiveType_NameXSD_reverseMap = make(map[*CodeType]*AbstractGeometricPrimitiveType)
	for abstractgeometricprimitivetype := range stage.AbstractGeometricPrimitiveTypes {
		_ = abstractgeometricprimitivetype
		for _, _codetype := range abstractgeometricprimitivetype.NameXSD {
			stage.AbstractGeometricPrimitiveType_NameXSD_reverseMap[_codetype] = abstractgeometricprimitivetype
		}
	}

	// Compute reverse map for named struct AbstractGeometryType
	// insertion point per field
	stage.AbstractGeometryType_NameXSD_reverseMap = make(map[*CodeType]*AbstractGeometryType)
	for abstractgeometrytype := range stage.AbstractGeometryTypes {
		_ = abstractgeometrytype
		for _, _codetype := range abstractgeometrytype.NameXSD {
			stage.AbstractGeometryType_NameXSD_reverseMap[_codetype] = abstractgeometrytype
		}
	}

	// Compute reverse map for named struct AbstractTimeGeometricPrimitiveType
	// insertion point per field
	stage.AbstractTimeGeometricPrimitiveType_NameXSD_reverseMap = make(map[*CodeType]*AbstractTimeGeometricPrimitiveType)
	for abstracttimegeometricprimitivetype := range stage.AbstractTimeGeometricPrimitiveTypes {
		_ = abstracttimegeometricprimitivetype
		for _, _codetype := range abstracttimegeometricprimitivetype.NameXSD {
			stage.AbstractTimeGeometricPrimitiveType_NameXSD_reverseMap[_codetype] = abstracttimegeometricprimitivetype
		}
	}

	// Compute reverse map for named struct AbstractTimeObjectType
	// insertion point per field
	stage.AbstractTimeObjectType_NameXSD_reverseMap = make(map[*CodeType]*AbstractTimeObjectType)
	for abstracttimeobjecttype := range stage.AbstractTimeObjectTypes {
		_ = abstracttimeobjecttype
		for _, _codetype := range abstracttimeobjecttype.NameXSD {
			stage.AbstractTimeObjectType_NameXSD_reverseMap[_codetype] = abstracttimeobjecttype
		}
	}

	// Compute reverse map for named struct AbstractTimePrimitiveType
	// insertion point per field
	stage.AbstractTimePrimitiveType_NameXSD_reverseMap = make(map[*CodeType]*AbstractTimePrimitiveType)
	for abstracttimeprimitivetype := range stage.AbstractTimePrimitiveTypes {
		_ = abstracttimeprimitivetype
		for _, _codetype := range abstracttimeprimitivetype.NameXSD {
			stage.AbstractTimePrimitiveType_NameXSD_reverseMap[_codetype] = abstracttimeprimitivetype
		}
	}

	// Compute reverse map for named struct AbstractTimeSliceType
	// insertion point per field
	stage.AbstractTimeSliceType_NameXSD_reverseMap = make(map[*CodeType]*AbstractTimeSliceType)
	for abstracttimeslicetype := range stage.AbstractTimeSliceTypes {
		_ = abstracttimeslicetype
		for _, _codetype := range abstracttimeslicetype.NameXSD {
			stage.AbstractTimeSliceType_NameXSD_reverseMap[_codetype] = abstracttimeslicetype
		}
	}

	// Compute reverse map for named struct DynamicFeatureType
	// insertion point per field
	stage.DynamicFeatureType_NameXSD_reverseMap = make(map[*CodeType]*DynamicFeatureType)
	for dynamicfeaturetype := range stage.DynamicFeatureTypes {
		_ = dynamicfeaturetype
		for _, _codetype := range dynamicfeaturetype.NameXSD {
			stage.DynamicFeatureType_NameXSD_reverseMap[_codetype] = dynamicfeaturetype
		}
	}

	// Compute reverse map for named struct PointType
	// insertion point per field
	stage.PointType_NameXSD_reverseMap = make(map[*CodeType]*PointType)
	for pointtype := range stage.PointTypes {
		_ = pointtype
		for _, _codetype := range pointtype.NameXSD {
			stage.PointType_NameXSD_reverseMap[_codetype] = pointtype
		}
	}

	// Compute reverse map for named struct TimeInstantType
	// insertion point per field
	stage.TimeInstantType_NameXSD_reverseMap = make(map[*CodeType]*TimeInstantType)
	for timeinstanttype := range stage.TimeInstantTypes {
		_ = timeinstanttype
		for _, _codetype := range timeinstanttype.NameXSD {
			stage.TimeInstantType_NameXSD_reverseMap[_codetype] = timeinstanttype
		}
	}

	// Compute reverse map for named struct TimePeriodType
	// insertion point per field
	stage.TimePeriodType_NameXSD_reverseMap = make(map[*CodeType]*TimePeriodType)
	for timeperiodtype := range stage.TimePeriodTypes {
		_ = timeperiodtype
		for _, _codetype := range timeperiodtype.NameXSD {
			stage.TimePeriodType_NameXSD_reverseMap[_codetype] = timeperiodtype
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AbstractFeatureMemberTypes)

	res = __gong__appendInstances(res, stage.AbstractFeatureTypes)

	res = __gong__appendInstances(res, stage.AbstractGMLTypes)

	res = __gong__appendInstances(res, stage.AbstractGeometricPrimitiveTypes)

	res = __gong__appendInstances(res, stage.AbstractGeometryTypes)

	res = __gong__appendInstances(res, stage.AbstractMetadataPropertyTypes)

	res = __gong__appendInstances(res, stage.AbstractTimeGeometricPrimitiveTypes)

	res = __gong__appendInstances(res, stage.AbstractTimeObjectTypes)

	res = __gong__appendInstances(res, stage.AbstractTimePrimitiveTypes)

	res = __gong__appendInstances(res, stage.AbstractTimeSliceTypes)

	res = __gong__appendInstances(res, stage.BoundingShapeTypes)

	res = __gong__appendInstances(res, stage.CodeTypes)

	res = __gong__appendInstances(res, stage.CodeWithAuthorityTypes)

	res = __gong__appendInstances(res, stage.DirectPositionTypes)

	res = __gong__appendInstances(res, stage.DynamicFeatureTypes)

	res = __gong__appendInstances(res, stage.EnvelopeTypes)

	res = __gong__appendInstances(res, stage.PointTypes)

	res = __gong__appendInstances(res, stage.TimeInstantTypes)

	res = __gong__appendInstances(res, stage.TimePeriodTypes)

	res = __gong__appendInstances(res, stage.TimePositionTypes)

	res = __gong__appendInstances(res, stage.TimePrimitivePropertyTypes)

	return
}

// insertion point per named struct
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongCopy() GongstructIF {
	newInstance := new(AbstractFeatureMemberType)
	abstractfeaturemembertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractfeaturetype *AbstractFeatureType) GongCopy() GongstructIF {
	newInstance := new(AbstractFeatureType)
	abstractfeaturetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractgmltype *AbstractGMLType) GongCopy() GongstructIF {
	newInstance := new(AbstractGMLType)
	abstractgmltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongCopy() GongstructIF {
	newInstance := new(AbstractGeometricPrimitiveType)
	abstractgeometricprimitivetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractgeometrytype *AbstractGeometryType) GongCopy() GongstructIF {
	newInstance := new(AbstractGeometryType)
	abstractgeometrytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongCopy() GongstructIF {
	newInstance := new(AbstractMetadataPropertyType)
	abstractmetadatapropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongCopy() GongstructIF {
	newInstance := new(AbstractTimeGeometricPrimitiveType)
	abstracttimegeometricprimitivetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongCopy() GongstructIF {
	newInstance := new(AbstractTimeObjectType)
	abstracttimeobjecttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongCopy() GongstructIF {
	newInstance := new(AbstractTimePrimitiveType)
	abstracttimeprimitivetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongCopy() GongstructIF {
	newInstance := new(AbstractTimeSliceType)
	abstracttimeslicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (boundingshapetype *BoundingShapeType) GongCopy() GongstructIF {
	newInstance := new(BoundingShapeType)
	boundingshapetype.GongCopyBasicFields(newInstance)
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

func (directpositiontype *DirectPositionType) GongCopy() GongstructIF {
	newInstance := new(DirectPositionType)
	directpositiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dynamicfeaturetype *DynamicFeatureType) GongCopy() GongstructIF {
	newInstance := new(DynamicFeatureType)
	dynamicfeaturetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (envelopetype *EnvelopeType) GongCopy() GongstructIF {
	newInstance := new(EnvelopeType)
	envelopetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pointtype *PointType) GongCopy() GongstructIF {
	newInstance := new(PointType)
	pointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timeinstanttype *TimeInstantType) GongCopy() GongstructIF {
	newInstance := new(TimeInstantType)
	timeinstanttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timeperiodtype *TimePeriodType) GongCopy() GongstructIF {
	newInstance := new(TimePeriodType)
	timeperiodtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timepositiontype *TimePositionType) GongCopy() GongstructIF {
	newInstance := new(TimePositionType)
	timepositiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongCopy() GongstructIF {
	newInstance := new(TimePrimitivePropertyType)
	timeprimitivepropertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractfeaturemembertype)
}

func (abstractfeaturetype *AbstractFeatureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractfeaturetype)
}

func (abstractgmltype *AbstractGMLType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractgmltype)
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractgeometricprimitivetype)
}

func (abstractgeometrytype *AbstractGeometryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractgeometrytype)
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractmetadatapropertytype)
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstracttimegeometricprimitivetype)
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstracttimeobjecttype)
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstracttimeprimitivetype)
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstracttimeslicetype)
}

func (boundingshapetype *BoundingShapeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, boundingshapetype)
}

func (codetype *CodeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetype)
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codewithauthoritytype)
}

func (directpositiontype *DirectPositionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, directpositiontype)
}

func (dynamicfeaturetype *DynamicFeatureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dynamicfeaturetype)
}

func (envelopetype *EnvelopeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, envelopetype)
}

func (pointtype *PointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pointtype)
}

func (timeinstanttype *TimeInstantType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timeinstanttype)
}

func (timeperiodtype *TimePeriodType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timeperiodtype)
}

func (timepositiontype *TimePositionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timepositiontype)
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timeprimitivepropertytype)
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
		stage.AbstractFeatureMemberTypes,
		stage.AbstractFeatureMemberType_stagedOrder,
		stage.AbstractFeatureMemberTypes_reference,
		&stage.AbstractFeatureMemberTypes_referenceOrder,
		stage.AbstractFeatureMemberTypes_instance,
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
		stage.AbstractFeatureTypes,
		stage.AbstractFeatureType_stagedOrder,
		stage.AbstractFeatureTypes_reference,
		&stage.AbstractFeatureTypes_referenceOrder,
		stage.AbstractFeatureTypes_instance,
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
		stage.AbstractGMLTypes,
		stage.AbstractGMLType_stagedOrder,
		stage.AbstractGMLTypes_reference,
		&stage.AbstractGMLTypes_referenceOrder,
		stage.AbstractGMLTypes_instance,
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
		stage.AbstractGeometricPrimitiveTypes,
		stage.AbstractGeometricPrimitiveType_stagedOrder,
		stage.AbstractGeometricPrimitiveTypes_reference,
		&stage.AbstractGeometricPrimitiveTypes_referenceOrder,
		stage.AbstractGeometricPrimitiveTypes_instance,
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
		stage.AbstractGeometryTypes,
		stage.AbstractGeometryType_stagedOrder,
		stage.AbstractGeometryTypes_reference,
		&stage.AbstractGeometryTypes_referenceOrder,
		stage.AbstractGeometryTypes_instance,
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
		stage.AbstractMetadataPropertyTypes,
		stage.AbstractMetadataPropertyType_stagedOrder,
		stage.AbstractMetadataPropertyTypes_reference,
		&stage.AbstractMetadataPropertyTypes_referenceOrder,
		stage.AbstractMetadataPropertyTypes_instance,
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
		stage.AbstractTimeGeometricPrimitiveTypes,
		stage.AbstractTimeGeometricPrimitiveType_stagedOrder,
		stage.AbstractTimeGeometricPrimitiveTypes_reference,
		&stage.AbstractTimeGeometricPrimitiveTypes_referenceOrder,
		stage.AbstractTimeGeometricPrimitiveTypes_instance,
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
		stage.AbstractTimeObjectTypes,
		stage.AbstractTimeObjectType_stagedOrder,
		stage.AbstractTimeObjectTypes_reference,
		&stage.AbstractTimeObjectTypes_referenceOrder,
		stage.AbstractTimeObjectTypes_instance,
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
		stage.AbstractTimePrimitiveTypes,
		stage.AbstractTimePrimitiveType_stagedOrder,
		stage.AbstractTimePrimitiveTypes_reference,
		&stage.AbstractTimePrimitiveTypes_referenceOrder,
		stage.AbstractTimePrimitiveTypes_instance,
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
		stage.AbstractTimeSliceTypes,
		stage.AbstractTimeSliceType_stagedOrder,
		stage.AbstractTimeSliceTypes_reference,
		&stage.AbstractTimeSliceTypes_referenceOrder,
		stage.AbstractTimeSliceTypes_instance,
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
		stage.BoundingShapeTypes,
		stage.BoundingShapeType_stagedOrder,
		stage.BoundingShapeTypes_reference,
		&stage.BoundingShapeTypes_referenceOrder,
		stage.BoundingShapeTypes_instance,
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
		stage.DirectPositionTypes,
		stage.DirectPositionType_stagedOrder,
		stage.DirectPositionTypes_reference,
		&stage.DirectPositionTypes_referenceOrder,
		stage.DirectPositionTypes_instance,
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
		stage.DynamicFeatureTypes,
		stage.DynamicFeatureType_stagedOrder,
		stage.DynamicFeatureTypes_reference,
		&stage.DynamicFeatureTypes_referenceOrder,
		stage.DynamicFeatureTypes_instance,
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
		stage.EnvelopeTypes,
		stage.EnvelopeType_stagedOrder,
		stage.EnvelopeTypes_reference,
		&stage.EnvelopeTypes_referenceOrder,
		stage.EnvelopeTypes_instance,
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
		stage.PointTypes,
		stage.PointType_stagedOrder,
		stage.PointTypes_reference,
		&stage.PointTypes_referenceOrder,
		stage.PointTypes_instance,
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
		stage.TimeInstantTypes,
		stage.TimeInstantType_stagedOrder,
		stage.TimeInstantTypes_reference,
		&stage.TimeInstantTypes_referenceOrder,
		stage.TimeInstantTypes_instance,
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
		stage.TimePeriodTypes,
		stage.TimePeriodType_stagedOrder,
		stage.TimePeriodTypes_reference,
		&stage.TimePeriodTypes_referenceOrder,
		stage.TimePeriodTypes_instance,
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
		stage.TimePositionTypes,
		stage.TimePositionType_stagedOrder,
		stage.TimePositionTypes_reference,
		&stage.TimePositionTypes_referenceOrder,
		stage.TimePositionTypes_instance,
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
		stage.TimePrimitivePropertyTypes,
		stage.TimePrimitivePropertyType_stagedOrder,
		stage.TimePrimitivePropertyTypes_reference,
		&stage.TimePrimitivePropertyTypes_referenceOrder,
		stage.TimePrimitivePropertyTypes_instance,
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
	__gong__computeReferencePass1(stage, stage.AbstractFeatureMemberTypes, &stage.AbstractFeatureMemberTypes_reference, &stage.AbstractFeatureMemberTypes_referenceOrder, &stage.AbstractFeatureMemberTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractFeatureTypes, &stage.AbstractFeatureTypes_reference, &stage.AbstractFeatureTypes_referenceOrder, &stage.AbstractFeatureTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractGMLTypes, &stage.AbstractGMLTypes_reference, &stage.AbstractGMLTypes_referenceOrder, &stage.AbstractGMLTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractGeometricPrimitiveTypes, &stage.AbstractGeometricPrimitiveTypes_reference, &stage.AbstractGeometricPrimitiveTypes_referenceOrder, &stage.AbstractGeometricPrimitiveTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractGeometryTypes, &stage.AbstractGeometryTypes_reference, &stage.AbstractGeometryTypes_referenceOrder, &stage.AbstractGeometryTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractMetadataPropertyTypes, &stage.AbstractMetadataPropertyTypes_reference, &stage.AbstractMetadataPropertyTypes_referenceOrder, &stage.AbstractMetadataPropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractTimeGeometricPrimitiveTypes, &stage.AbstractTimeGeometricPrimitiveTypes_reference, &stage.AbstractTimeGeometricPrimitiveTypes_referenceOrder, &stage.AbstractTimeGeometricPrimitiveTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractTimeObjectTypes, &stage.AbstractTimeObjectTypes_reference, &stage.AbstractTimeObjectTypes_referenceOrder, &stage.AbstractTimeObjectTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractTimePrimitiveTypes, &stage.AbstractTimePrimitiveTypes_reference, &stage.AbstractTimePrimitiveTypes_referenceOrder, &stage.AbstractTimePrimitiveTypes_instance)

	__gong__computeReferencePass1(stage, stage.AbstractTimeSliceTypes, &stage.AbstractTimeSliceTypes_reference, &stage.AbstractTimeSliceTypes_referenceOrder, &stage.AbstractTimeSliceTypes_instance)

	__gong__computeReferencePass1(stage, stage.BoundingShapeTypes, &stage.BoundingShapeTypes_reference, &stage.BoundingShapeTypes_referenceOrder, &stage.BoundingShapeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTypes, &stage.CodeTypes_reference, &stage.CodeTypes_referenceOrder, &stage.CodeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeWithAuthorityTypes, &stage.CodeWithAuthorityTypes_reference, &stage.CodeWithAuthorityTypes_referenceOrder, &stage.CodeWithAuthorityTypes_instance)

	__gong__computeReferencePass1(stage, stage.DirectPositionTypes, &stage.DirectPositionTypes_reference, &stage.DirectPositionTypes_referenceOrder, &stage.DirectPositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.DynamicFeatureTypes, &stage.DynamicFeatureTypes_reference, &stage.DynamicFeatureTypes_referenceOrder, &stage.DynamicFeatureTypes_instance)

	__gong__computeReferencePass1(stage, stage.EnvelopeTypes, &stage.EnvelopeTypes_reference, &stage.EnvelopeTypes_referenceOrder, &stage.EnvelopeTypes_instance)

	__gong__computeReferencePass1(stage, stage.PointTypes, &stage.PointTypes_reference, &stage.PointTypes_referenceOrder, &stage.PointTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimeInstantTypes, &stage.TimeInstantTypes_reference, &stage.TimeInstantTypes_referenceOrder, &stage.TimeInstantTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimePeriodTypes, &stage.TimePeriodTypes_reference, &stage.TimePeriodTypes_referenceOrder, &stage.TimePeriodTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimePositionTypes, &stage.TimePositionTypes_reference, &stage.TimePositionTypes_referenceOrder, &stage.TimePositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimePrimitivePropertyTypes, &stage.TimePrimitivePropertyTypes_reference, &stage.TimePrimitivePropertyTypes_referenceOrder, &stage.TimePrimitivePropertyTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AbstractFeatureMemberTypes, stage.AbstractFeatureMemberTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractFeatureTypes, stage.AbstractFeatureTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractGMLTypes, stage.AbstractGMLTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractGeometricPrimitiveTypes, stage.AbstractGeometricPrimitiveTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractGeometryTypes, stage.AbstractGeometryTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractMetadataPropertyTypes, stage.AbstractMetadataPropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractTimeGeometricPrimitiveTypes, stage.AbstractTimeGeometricPrimitiveTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractTimeObjectTypes, stage.AbstractTimeObjectTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractTimePrimitiveTypes, stage.AbstractTimePrimitiveTypes_reference, stage)

	__gong__computeReferencePass2(stage.AbstractTimeSliceTypes, stage.AbstractTimeSliceTypes_reference, stage)

	__gong__computeReferencePass2(stage.BoundingShapeTypes, stage.BoundingShapeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTypes, stage.CodeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeWithAuthorityTypes, stage.CodeWithAuthorityTypes_reference, stage)

	__gong__computeReferencePass2(stage.DirectPositionTypes, stage.DirectPositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.DynamicFeatureTypes, stage.DynamicFeatureTypes_reference, stage)

	__gong__computeReferencePass2(stage.EnvelopeTypes, stage.EnvelopeTypes_reference, stage)

	__gong__computeReferencePass2(stage.PointTypes, stage.PointTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimeInstantTypes, stage.TimeInstantTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimePeriodTypes, stage.TimePeriodTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimePositionTypes, stage.TimePositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimePrimitivePropertyTypes, stage.TimePrimitivePropertyTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractFeatureMemberType_stagedOrder, stage.AbstractFeatureMemberTypes_referenceOrder, abstractfeaturemembertype, "AbstractFeatureMemberType")
}

func (abstractfeaturetype *AbstractFeatureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractFeatureType_stagedOrder, stage.AbstractFeatureTypes_referenceOrder, abstractfeaturetype, "AbstractFeatureType")
}

func (abstractgmltype *AbstractGMLType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractGMLType_stagedOrder, stage.AbstractGMLTypes_referenceOrder, abstractgmltype, "AbstractGMLType")
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractGeometricPrimitiveType_stagedOrder, stage.AbstractGeometricPrimitiveTypes_referenceOrder, abstractgeometricprimitivetype, "AbstractGeometricPrimitiveType")
}

func (abstractgeometrytype *AbstractGeometryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractGeometryType_stagedOrder, stage.AbstractGeometryTypes_referenceOrder, abstractgeometrytype, "AbstractGeometryType")
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractMetadataPropertyType_stagedOrder, stage.AbstractMetadataPropertyTypes_referenceOrder, abstractmetadatapropertytype, "AbstractMetadataPropertyType")
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractTimeGeometricPrimitiveType_stagedOrder, stage.AbstractTimeGeometricPrimitiveTypes_referenceOrder, abstracttimegeometricprimitivetype, "AbstractTimeGeometricPrimitiveType")
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractTimeObjectType_stagedOrder, stage.AbstractTimeObjectTypes_referenceOrder, abstracttimeobjecttype, "AbstractTimeObjectType")
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractTimePrimitiveType_stagedOrder, stage.AbstractTimePrimitiveTypes_referenceOrder, abstracttimeprimitivetype, "AbstractTimePrimitiveType")
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractTimeSliceType_stagedOrder, stage.AbstractTimeSliceTypes_referenceOrder, abstracttimeslicetype, "AbstractTimeSliceType")
}

func (boundingshapetype *BoundingShapeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.BoundingShapeType_stagedOrder, stage.BoundingShapeTypes_referenceOrder, boundingshapetype, "BoundingShapeType")
}

func (codetype *CodeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeType_stagedOrder, stage.CodeTypes_referenceOrder, codetype, "CodeType")
}

func (codewithauthoritytype *CodeWithAuthorityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeWithAuthorityType_stagedOrder, stage.CodeWithAuthorityTypes_referenceOrder, codewithauthoritytype, "CodeWithAuthorityType")
}

func (directpositiontype *DirectPositionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DirectPositionType_stagedOrder, stage.DirectPositionTypes_referenceOrder, directpositiontype, "DirectPositionType")
}

func (dynamicfeaturetype *DynamicFeatureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DynamicFeatureType_stagedOrder, stage.DynamicFeatureTypes_referenceOrder, dynamicfeaturetype, "DynamicFeatureType")
}

func (envelopetype *EnvelopeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EnvelopeType_stagedOrder, stage.EnvelopeTypes_referenceOrder, envelopetype, "EnvelopeType")
}

func (pointtype *PointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PointType_stagedOrder, stage.PointTypes_referenceOrder, pointtype, "PointType")
}

func (timeinstanttype *TimeInstantType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimeInstantType_stagedOrder, stage.TimeInstantTypes_referenceOrder, timeinstanttype, "TimeInstantType")
}

func (timeperiodtype *TimePeriodType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimePeriodType_stagedOrder, stage.TimePeriodTypes_referenceOrder, timeperiodtype, "TimePeriodType")
}

func (timepositiontype *TimePositionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimePositionType_stagedOrder, stage.TimePositionTypes_referenceOrder, timepositiontype, "TimePositionType")
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimePrimitivePropertyType_stagedOrder, stage.TimePrimitivePropertyTypes_referenceOrder, timeprimitivepropertytype, "TimePrimitivePropertyType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractfeaturemembertype, abstractfeaturemembertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractfeaturemembertype.GongGetIdentifier(stage)
}

func (abstractfeaturetype *AbstractFeatureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractfeaturetype, abstractfeaturetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractfeaturetype *AbstractFeatureType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractfeaturetype.GongGetIdentifier(stage)
}

func (abstractgmltype *AbstractGMLType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractgmltype, abstractgmltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractgmltype *AbstractGMLType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractgmltype.GongGetIdentifier(stage)
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractgeometricprimitivetype, abstractgeometricprimitivetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractgeometricprimitivetype.GongGetIdentifier(stage)
}

func (abstractgeometrytype *AbstractGeometryType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractgeometrytype, abstractgeometrytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractgeometrytype *AbstractGeometryType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractgeometrytype.GongGetIdentifier(stage)
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractmetadatapropertytype, abstractmetadatapropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractmetadatapropertytype.GongGetIdentifier(stage)
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstracttimegeometricprimitivetype, abstracttimegeometricprimitivetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstracttimegeometricprimitivetype.GongGetIdentifier(stage)
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstracttimeobjecttype, abstracttimeobjecttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstracttimeobjecttype *AbstractTimeObjectType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstracttimeobjecttype.GongGetIdentifier(stage)
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstracttimeprimitivetype, abstracttimeprimitivetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstracttimeprimitivetype.GongGetIdentifier(stage)
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstracttimeslicetype, abstracttimeslicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstracttimeslicetype *AbstractTimeSliceType) GongGetReferenceIdentifier(stage *Stage) string {
	return abstracttimeslicetype.GongGetIdentifier(stage)
}

func (boundingshapetype *BoundingShapeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(boundingshapetype, boundingshapetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (boundingshapetype *BoundingShapeType) GongGetReferenceIdentifier(stage *Stage) string {
	return boundingshapetype.GongGetIdentifier(stage)
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

func (directpositiontype *DirectPositionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(directpositiontype, directpositiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (directpositiontype *DirectPositionType) GongGetReferenceIdentifier(stage *Stage) string {
	return directpositiontype.GongGetIdentifier(stage)
}

func (dynamicfeaturetype *DynamicFeatureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dynamicfeaturetype, dynamicfeaturetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dynamicfeaturetype *DynamicFeatureType) GongGetReferenceIdentifier(stage *Stage) string {
	return dynamicfeaturetype.GongGetIdentifier(stage)
}

func (envelopetype *EnvelopeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(envelopetype, envelopetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (envelopetype *EnvelopeType) GongGetReferenceIdentifier(stage *Stage) string {
	return envelopetype.GongGetIdentifier(stage)
}

func (pointtype *PointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pointtype, pointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointtype *PointType) GongGetReferenceIdentifier(stage *Stage) string {
	return pointtype.GongGetIdentifier(stage)
}

func (timeinstanttype *TimeInstantType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timeinstanttype, timeinstanttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timeinstanttype *TimeInstantType) GongGetReferenceIdentifier(stage *Stage) string {
	return timeinstanttype.GongGetIdentifier(stage)
}

func (timeperiodtype *TimePeriodType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timeperiodtype, timeperiodtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timeperiodtype *TimePeriodType) GongGetReferenceIdentifier(stage *Stage) string {
	return timeperiodtype.GongGetIdentifier(stage)
}

func (timepositiontype *TimePositionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timepositiontype, timepositiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timepositiontype *TimePositionType) GongGetReferenceIdentifier(stage *Stage) string {
	return timepositiontype.GongGetIdentifier(stage)
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timeprimitivepropertytype, timeprimitivepropertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timeprimitivepropertytype *TimePrimitivePropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return timeprimitivepropertytype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractfeaturemembertype.GongGetIdentifier(stage), "AbstractFeatureMemberType", abstractfeaturemembertype.Name)
}

func (abstractfeaturetype *AbstractFeatureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractfeaturetype.GongGetIdentifier(stage), "AbstractFeatureType", abstractfeaturetype.Name)
}

func (abstractgmltype *AbstractGMLType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractgmltype.GongGetIdentifier(stage), "AbstractGMLType", abstractgmltype.Name)
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractgeometricprimitivetype.GongGetIdentifier(stage), "AbstractGeometricPrimitiveType", abstractgeometricprimitivetype.Name)
}

func (abstractgeometrytype *AbstractGeometryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractgeometrytype.GongGetIdentifier(stage), "AbstractGeometryType", abstractgeometrytype.Name)
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractmetadatapropertytype.GongGetIdentifier(stage), "AbstractMetadataPropertyType", abstractmetadatapropertytype.Name)
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstracttimegeometricprimitivetype.GongGetIdentifier(stage), "AbstractTimeGeometricPrimitiveType", abstracttimegeometricprimitivetype.Name)
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstracttimeobjecttype.GongGetIdentifier(stage), "AbstractTimeObjectType", abstracttimeobjecttype.Name)
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstracttimeprimitivetype.GongGetIdentifier(stage), "AbstractTimePrimitiveType", abstracttimeprimitivetype.Name)
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstracttimeslicetype.GongGetIdentifier(stage), "AbstractTimeSliceType", abstracttimeslicetype.Name)
}

func (boundingshapetype *BoundingShapeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(boundingshapetype.GongGetIdentifier(stage), "BoundingShapeType", boundingshapetype.Name)
}

func (codetype *CodeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetype.GongGetIdentifier(stage), "CodeType", codetype.Name)
}

func (codewithauthoritytype *CodeWithAuthorityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codewithauthoritytype.GongGetIdentifier(stage), "CodeWithAuthorityType", codewithauthoritytype.Name)
}

func (directpositiontype *DirectPositionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(directpositiontype.GongGetIdentifier(stage), "DirectPositionType", directpositiontype.Name)
}

func (dynamicfeaturetype *DynamicFeatureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dynamicfeaturetype.GongGetIdentifier(stage), "DynamicFeatureType", dynamicfeaturetype.Name)
}

func (envelopetype *EnvelopeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(envelopetype.GongGetIdentifier(stage), "EnvelopeType", envelopetype.Name)
}

func (pointtype *PointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pointtype.GongGetIdentifier(stage), "PointType", pointtype.Name)
}

func (timeinstanttype *TimeInstantType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timeinstanttype.GongGetIdentifier(stage), "TimeInstantType", timeinstanttype.Name)
}

func (timeperiodtype *TimePeriodType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timeperiodtype.GongGetIdentifier(stage), "TimePeriodType", timeperiodtype.Name)
}

func (timepositiontype *TimePositionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timepositiontype.GongGetIdentifier(stage), "TimePositionType", timepositiontype.Name)
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timeprimitivepropertytype.GongGetIdentifier(stage), "TimePrimitivePropertyType", timeprimitivepropertytype.Name)
}

// insertion point for unstaging
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractfeaturemembertype.GongGetReferenceIdentifier(stage))
}

func (abstractfeaturetype *AbstractFeatureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractfeaturetype.GongGetReferenceIdentifier(stage))
}

func (abstractgmltype *AbstractGMLType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractgmltype.GongGetReferenceIdentifier(stage))
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractgeometricprimitivetype.GongGetReferenceIdentifier(stage))
}

func (abstractgeometrytype *AbstractGeometryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractgeometrytype.GongGetReferenceIdentifier(stage))
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractmetadatapropertytype.GongGetReferenceIdentifier(stage))
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstracttimegeometricprimitivetype.GongGetReferenceIdentifier(stage))
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstracttimeobjecttype.GongGetReferenceIdentifier(stage))
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstracttimeprimitivetype.GongGetReferenceIdentifier(stage))
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstracttimeslicetype.GongGetReferenceIdentifier(stage))
}

func (boundingshapetype *BoundingShapeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(boundingshapetype.GongGetReferenceIdentifier(stage))
}

func (codetype *CodeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetype.GongGetReferenceIdentifier(stage))
}

func (codewithauthoritytype *CodeWithAuthorityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codewithauthoritytype.GongGetReferenceIdentifier(stage))
}

func (directpositiontype *DirectPositionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(directpositiontype.GongGetReferenceIdentifier(stage))
}

func (dynamicfeaturetype *DynamicFeatureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dynamicfeaturetype.GongGetReferenceIdentifier(stage))
}

func (envelopetype *EnvelopeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(envelopetype.GongGetReferenceIdentifier(stage))
}

func (pointtype *PointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pointtype.GongGetReferenceIdentifier(stage))
}

func (timeinstanttype *TimeInstantType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timeinstanttype.GongGetReferenceIdentifier(stage))
}

func (timeperiodtype *TimePeriodType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timeperiodtype.GongGetReferenceIdentifier(stage))
}

func (timepositiontype *TimePositionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timepositiontype.GongGetReferenceIdentifier(stage))
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timeprimitivepropertytype.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__gml__%s__%08d_", s.GongGetGongstructName(), order)
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
