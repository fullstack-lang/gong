// generated code - do not edit
package content

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
	// Compute reverse map for named struct MD_CoverageDescription_Type
	// insertion point per field
	stage.MD_CoverageDescription_Type_Dimension_reverseMap = make(map[*MD_RangeDimension_PropertyType]*MD_CoverageDescription_Type)
	for md_coveragedescription_type := range stage.MD_CoverageDescription_Types {
		_ = md_coveragedescription_type
		for _, _md_rangedimension_propertytype := range md_coveragedescription_type.Dimension {
			stage.MD_CoverageDescription_Type_Dimension_reverseMap[_md_rangedimension_propertytype] = md_coveragedescription_type
		}
	}

	// Compute reverse map for named struct MD_ImageDescription_Type
	// insertion point per field
	stage.MD_ImageDescription_Type_Dimension_reverseMap = make(map[*MD_RangeDimension_PropertyType]*MD_ImageDescription_Type)
	for md_imagedescription_type := range stage.MD_ImageDescription_Types {
		_ = md_imagedescription_type
		for _, _md_rangedimension_propertytype := range md_imagedescription_type.Dimension {
			stage.MD_ImageDescription_Type_Dimension_reverseMap[_md_rangedimension_propertytype] = md_imagedescription_type
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AbstractMD_ContentInformation_Types)

	res = __gong__appendInstances(res, stage.MD_Band_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_Band_Types)

	res = __gong__appendInstances(res, stage.MD_ContentInformation_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_CoverageContentTypeCode_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_CoverageDescription_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_CoverageDescription_Types)

	res = __gong__appendInstances(res, stage.MD_FeatureCatalogueDescription_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_FeatureCatalogueDescription_Types)

	res = __gong__appendInstances(res, stage.MD_ImageDescription_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_ImageDescription_Types)

	res = __gong__appendInstances(res, stage.MD_ImagingConditionCode_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_RangeDimension_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_RangeDimension_Types)

	return
}

// insertion point per named struct
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongCopy() GongstructIF {
	newInstance := new(AbstractMD_ContentInformation_Type)
	abstractmd_contentinformation_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_band_propertytype *MD_Band_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_Band_PropertyType)
	md_band_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_band_type *MD_Band_Type) GongCopy() GongstructIF {
	newInstance := new(MD_Band_Type)
	md_band_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_ContentInformation_PropertyType)
	md_contentinformation_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_CoverageContentTypeCode_PropertyType)
	md_coveragecontenttypecode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_CoverageDescription_PropertyType)
	md_coveragedescription_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongCopy() GongstructIF {
	newInstance := new(MD_CoverageDescription_Type)
	md_coveragedescription_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_FeatureCatalogueDescription_PropertyType)
	md_featurecataloguedescription_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongCopy() GongstructIF {
	newInstance := new(MD_FeatureCatalogueDescription_Type)
	md_featurecataloguedescription_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_ImageDescription_PropertyType)
	md_imagedescription_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongCopy() GongstructIF {
	newInstance := new(MD_ImageDescription_Type)
	md_imagedescription_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_ImagingConditionCode_PropertyType)
	md_imagingconditioncode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_RangeDimension_PropertyType)
	md_rangedimension_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongCopy() GongstructIF {
	newInstance := new(MD_RangeDimension_Type)
	md_rangedimension_type.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractmd_contentinformation_type)
}

func (md_band_propertytype *MD_Band_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_band_propertytype)
}

func (md_band_type *MD_Band_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_band_type)
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_contentinformation_propertytype)
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_coveragecontenttypecode_propertytype)
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_coveragedescription_propertytype)
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_coveragedescription_type)
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_featurecataloguedescription_propertytype)
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_featurecataloguedescription_type)
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_imagedescription_propertytype)
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_imagedescription_type)
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_imagingconditioncode_propertytype)
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_rangedimension_propertytype)
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_rangedimension_type)
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
		stage.AbstractMD_ContentInformation_Types,
		stage.AbstractMD_ContentInformation_Type_stagedOrder,
		stage.AbstractMD_ContentInformation_Types_reference,
		&stage.AbstractMD_ContentInformation_Types_referenceOrder,
		stage.AbstractMD_ContentInformation_Types_instance,
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
		stage.MD_Band_PropertyTypes,
		stage.MD_Band_PropertyType_stagedOrder,
		stage.MD_Band_PropertyTypes_reference,
		&stage.MD_Band_PropertyTypes_referenceOrder,
		stage.MD_Band_PropertyTypes_instance,
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
		stage.MD_Band_Types,
		stage.MD_Band_Type_stagedOrder,
		stage.MD_Band_Types_reference,
		&stage.MD_Band_Types_referenceOrder,
		stage.MD_Band_Types_instance,
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
		stage.MD_ContentInformation_PropertyTypes,
		stage.MD_ContentInformation_PropertyType_stagedOrder,
		stage.MD_ContentInformation_PropertyTypes_reference,
		&stage.MD_ContentInformation_PropertyTypes_referenceOrder,
		stage.MD_ContentInformation_PropertyTypes_instance,
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
		stage.MD_CoverageContentTypeCode_PropertyTypes,
		stage.MD_CoverageContentTypeCode_PropertyType_stagedOrder,
		stage.MD_CoverageContentTypeCode_PropertyTypes_reference,
		&stage.MD_CoverageContentTypeCode_PropertyTypes_referenceOrder,
		stage.MD_CoverageContentTypeCode_PropertyTypes_instance,
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
		stage.MD_CoverageDescription_PropertyTypes,
		stage.MD_CoverageDescription_PropertyType_stagedOrder,
		stage.MD_CoverageDescription_PropertyTypes_reference,
		&stage.MD_CoverageDescription_PropertyTypes_referenceOrder,
		stage.MD_CoverageDescription_PropertyTypes_instance,
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
		stage.MD_CoverageDescription_Types,
		stage.MD_CoverageDescription_Type_stagedOrder,
		stage.MD_CoverageDescription_Types_reference,
		&stage.MD_CoverageDescription_Types_referenceOrder,
		stage.MD_CoverageDescription_Types_instance,
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
		stage.MD_FeatureCatalogueDescription_PropertyTypes,
		stage.MD_FeatureCatalogueDescription_PropertyType_stagedOrder,
		stage.MD_FeatureCatalogueDescription_PropertyTypes_reference,
		&stage.MD_FeatureCatalogueDescription_PropertyTypes_referenceOrder,
		stage.MD_FeatureCatalogueDescription_PropertyTypes_instance,
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
		stage.MD_FeatureCatalogueDescription_Types,
		stage.MD_FeatureCatalogueDescription_Type_stagedOrder,
		stage.MD_FeatureCatalogueDescription_Types_reference,
		&stage.MD_FeatureCatalogueDescription_Types_referenceOrder,
		stage.MD_FeatureCatalogueDescription_Types_instance,
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
		stage.MD_ImageDescription_PropertyTypes,
		stage.MD_ImageDescription_PropertyType_stagedOrder,
		stage.MD_ImageDescription_PropertyTypes_reference,
		&stage.MD_ImageDescription_PropertyTypes_referenceOrder,
		stage.MD_ImageDescription_PropertyTypes_instance,
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
		stage.MD_ImageDescription_Types,
		stage.MD_ImageDescription_Type_stagedOrder,
		stage.MD_ImageDescription_Types_reference,
		&stage.MD_ImageDescription_Types_referenceOrder,
		stage.MD_ImageDescription_Types_instance,
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
		stage.MD_ImagingConditionCode_PropertyTypes,
		stage.MD_ImagingConditionCode_PropertyType_stagedOrder,
		stage.MD_ImagingConditionCode_PropertyTypes_reference,
		&stage.MD_ImagingConditionCode_PropertyTypes_referenceOrder,
		stage.MD_ImagingConditionCode_PropertyTypes_instance,
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
		stage.MD_RangeDimension_PropertyTypes,
		stage.MD_RangeDimension_PropertyType_stagedOrder,
		stage.MD_RangeDimension_PropertyTypes_reference,
		&stage.MD_RangeDimension_PropertyTypes_referenceOrder,
		stage.MD_RangeDimension_PropertyTypes_instance,
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
		stage.MD_RangeDimension_Types,
		stage.MD_RangeDimension_Type_stagedOrder,
		stage.MD_RangeDimension_Types_reference,
		&stage.MD_RangeDimension_Types_referenceOrder,
		stage.MD_RangeDimension_Types_instance,
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
	__gong__computeReferencePass1(stage, stage.AbstractMD_ContentInformation_Types, &stage.AbstractMD_ContentInformation_Types_reference, &stage.AbstractMD_ContentInformation_Types_referenceOrder, &stage.AbstractMD_ContentInformation_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_Band_PropertyTypes, &stage.MD_Band_PropertyTypes_reference, &stage.MD_Band_PropertyTypes_referenceOrder, &stage.MD_Band_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_Band_Types, &stage.MD_Band_Types_reference, &stage.MD_Band_Types_referenceOrder, &stage.MD_Band_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_ContentInformation_PropertyTypes, &stage.MD_ContentInformation_PropertyTypes_reference, &stage.MD_ContentInformation_PropertyTypes_referenceOrder, &stage.MD_ContentInformation_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_CoverageContentTypeCode_PropertyTypes, &stage.MD_CoverageContentTypeCode_PropertyTypes_reference, &stage.MD_CoverageContentTypeCode_PropertyTypes_referenceOrder, &stage.MD_CoverageContentTypeCode_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_CoverageDescription_PropertyTypes, &stage.MD_CoverageDescription_PropertyTypes_reference, &stage.MD_CoverageDescription_PropertyTypes_referenceOrder, &stage.MD_CoverageDescription_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_CoverageDescription_Types, &stage.MD_CoverageDescription_Types_reference, &stage.MD_CoverageDescription_Types_referenceOrder, &stage.MD_CoverageDescription_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_FeatureCatalogueDescription_PropertyTypes, &stage.MD_FeatureCatalogueDescription_PropertyTypes_reference, &stage.MD_FeatureCatalogueDescription_PropertyTypes_referenceOrder, &stage.MD_FeatureCatalogueDescription_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_FeatureCatalogueDescription_Types, &stage.MD_FeatureCatalogueDescription_Types_reference, &stage.MD_FeatureCatalogueDescription_Types_referenceOrder, &stage.MD_FeatureCatalogueDescription_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_ImageDescription_PropertyTypes, &stage.MD_ImageDescription_PropertyTypes_reference, &stage.MD_ImageDescription_PropertyTypes_referenceOrder, &stage.MD_ImageDescription_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_ImageDescription_Types, &stage.MD_ImageDescription_Types_reference, &stage.MD_ImageDescription_Types_referenceOrder, &stage.MD_ImageDescription_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_ImagingConditionCode_PropertyTypes, &stage.MD_ImagingConditionCode_PropertyTypes_reference, &stage.MD_ImagingConditionCode_PropertyTypes_referenceOrder, &stage.MD_ImagingConditionCode_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_RangeDimension_PropertyTypes, &stage.MD_RangeDimension_PropertyTypes_reference, &stage.MD_RangeDimension_PropertyTypes_referenceOrder, &stage.MD_RangeDimension_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_RangeDimension_Types, &stage.MD_RangeDimension_Types_reference, &stage.MD_RangeDimension_Types_referenceOrder, &stage.MD_RangeDimension_Types_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AbstractMD_ContentInformation_Types, stage.AbstractMD_ContentInformation_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_Band_PropertyTypes, stage.MD_Band_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_Band_Types, stage.MD_Band_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_ContentInformation_PropertyTypes, stage.MD_ContentInformation_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_CoverageContentTypeCode_PropertyTypes, stage.MD_CoverageContentTypeCode_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_CoverageDescription_PropertyTypes, stage.MD_CoverageDescription_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_CoverageDescription_Types, stage.MD_CoverageDescription_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_FeatureCatalogueDescription_PropertyTypes, stage.MD_FeatureCatalogueDescription_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_FeatureCatalogueDescription_Types, stage.MD_FeatureCatalogueDescription_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_ImageDescription_PropertyTypes, stage.MD_ImageDescription_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_ImageDescription_Types, stage.MD_ImageDescription_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_ImagingConditionCode_PropertyTypes, stage.MD_ImagingConditionCode_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_RangeDimension_PropertyTypes, stage.MD_RangeDimension_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_RangeDimension_Types, stage.MD_RangeDimension_Types_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractMD_ContentInformation_Type_stagedOrder, stage.AbstractMD_ContentInformation_Types_referenceOrder, abstractmd_contentinformation_type, "AbstractMD_ContentInformation_Type")
}

func (md_band_propertytype *MD_Band_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Band_PropertyType_stagedOrder, stage.MD_Band_PropertyTypes_referenceOrder, md_band_propertytype, "MD_Band_PropertyType")
}

func (md_band_type *MD_Band_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Band_Type_stagedOrder, stage.MD_Band_Types_referenceOrder, md_band_type, "MD_Band_Type")
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_ContentInformation_PropertyType_stagedOrder, stage.MD_ContentInformation_PropertyTypes_referenceOrder, md_contentinformation_propertytype, "MD_ContentInformation_PropertyType")
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_CoverageContentTypeCode_PropertyType_stagedOrder, stage.MD_CoverageContentTypeCode_PropertyTypes_referenceOrder, md_coveragecontenttypecode_propertytype, "MD_CoverageContentTypeCode_PropertyType")
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_CoverageDescription_PropertyType_stagedOrder, stage.MD_CoverageDescription_PropertyTypes_referenceOrder, md_coveragedescription_propertytype, "MD_CoverageDescription_PropertyType")
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_CoverageDescription_Type_stagedOrder, stage.MD_CoverageDescription_Types_referenceOrder, md_coveragedescription_type, "MD_CoverageDescription_Type")
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_FeatureCatalogueDescription_PropertyType_stagedOrder, stage.MD_FeatureCatalogueDescription_PropertyTypes_referenceOrder, md_featurecataloguedescription_propertytype, "MD_FeatureCatalogueDescription_PropertyType")
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_FeatureCatalogueDescription_Type_stagedOrder, stage.MD_FeatureCatalogueDescription_Types_referenceOrder, md_featurecataloguedescription_type, "MD_FeatureCatalogueDescription_Type")
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_ImageDescription_PropertyType_stagedOrder, stage.MD_ImageDescription_PropertyTypes_referenceOrder, md_imagedescription_propertytype, "MD_ImageDescription_PropertyType")
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_ImageDescription_Type_stagedOrder, stage.MD_ImageDescription_Types_referenceOrder, md_imagedescription_type, "MD_ImageDescription_Type")
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_ImagingConditionCode_PropertyType_stagedOrder, stage.MD_ImagingConditionCode_PropertyTypes_referenceOrder, md_imagingconditioncode_propertytype, "MD_ImagingConditionCode_PropertyType")
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_RangeDimension_PropertyType_stagedOrder, stage.MD_RangeDimension_PropertyTypes_referenceOrder, md_rangedimension_propertytype, "MD_RangeDimension_PropertyType")
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_RangeDimension_Type_stagedOrder, stage.MD_RangeDimension_Types_referenceOrder, md_rangedimension_type, "MD_RangeDimension_Type")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractmd_contentinformation_type, abstractmd_contentinformation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractmd_contentinformation_type.GongGetIdentifier(stage)
}

func (md_band_propertytype *MD_Band_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_band_propertytype, md_band_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_band_propertytype *MD_Band_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_band_propertytype.GongGetIdentifier(stage)
}

func (md_band_type *MD_Band_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_band_type, md_band_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_band_type *MD_Band_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_band_type.GongGetIdentifier(stage)
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_contentinformation_propertytype, md_contentinformation_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_contentinformation_propertytype.GongGetIdentifier(stage)
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_coveragecontenttypecode_propertytype, md_coveragecontenttypecode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_coveragecontenttypecode_propertytype.GongGetIdentifier(stage)
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_coveragedescription_propertytype, md_coveragedescription_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_coveragedescription_propertytype.GongGetIdentifier(stage)
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_coveragedescription_type, md_coveragedescription_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_coveragedescription_type *MD_CoverageDescription_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_coveragedescription_type.GongGetIdentifier(stage)
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_featurecataloguedescription_propertytype, md_featurecataloguedescription_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_featurecataloguedescription_propertytype.GongGetIdentifier(stage)
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_featurecataloguedescription_type, md_featurecataloguedescription_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_featurecataloguedescription_type.GongGetIdentifier(stage)
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_imagedescription_propertytype, md_imagedescription_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_imagedescription_propertytype.GongGetIdentifier(stage)
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_imagedescription_type, md_imagedescription_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_imagedescription_type *MD_ImageDescription_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_imagedescription_type.GongGetIdentifier(stage)
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_imagingconditioncode_propertytype, md_imagingconditioncode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_imagingconditioncode_propertytype.GongGetIdentifier(stage)
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_rangedimension_propertytype, md_rangedimension_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_rangedimension_propertytype.GongGetIdentifier(stage)
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_rangedimension_type, md_rangedimension_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_rangedimension_type *MD_RangeDimension_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_rangedimension_type.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractmd_contentinformation_type.GongGetIdentifier(stage), "AbstractMD_ContentInformation_Type", abstractmd_contentinformation_type.Name)
}

func (md_band_propertytype *MD_Band_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_band_propertytype.GongGetIdentifier(stage), "MD_Band_PropertyType", md_band_propertytype.Name)
}

func (md_band_type *MD_Band_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_band_type.GongGetIdentifier(stage), "MD_Band_Type", md_band_type.Name)
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_contentinformation_propertytype.GongGetIdentifier(stage), "MD_ContentInformation_PropertyType", md_contentinformation_propertytype.Name)
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_coveragecontenttypecode_propertytype.GongGetIdentifier(stage), "MD_CoverageContentTypeCode_PropertyType", md_coveragecontenttypecode_propertytype.Name)
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_coveragedescription_propertytype.GongGetIdentifier(stage), "MD_CoverageDescription_PropertyType", md_coveragedescription_propertytype.Name)
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_coveragedescription_type.GongGetIdentifier(stage), "MD_CoverageDescription_Type", md_coveragedescription_type.Name)
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_featurecataloguedescription_propertytype.GongGetIdentifier(stage), "MD_FeatureCatalogueDescription_PropertyType", md_featurecataloguedescription_propertytype.Name)
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_featurecataloguedescription_type.GongGetIdentifier(stage), "MD_FeatureCatalogueDescription_Type", md_featurecataloguedescription_type.Name)
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_imagedescription_propertytype.GongGetIdentifier(stage), "MD_ImageDescription_PropertyType", md_imagedescription_propertytype.Name)
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_imagedescription_type.GongGetIdentifier(stage), "MD_ImageDescription_Type", md_imagedescription_type.Name)
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_imagingconditioncode_propertytype.GongGetIdentifier(stage), "MD_ImagingConditionCode_PropertyType", md_imagingconditioncode_propertytype.Name)
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_rangedimension_propertytype.GongGetIdentifier(stage), "MD_RangeDimension_PropertyType", md_rangedimension_propertytype.Name)
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_rangedimension_type.GongGetIdentifier(stage), "MD_RangeDimension_Type", md_rangedimension_type.Name)
}

// insertion point for unstaging
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractmd_contentinformation_type.GongGetReferenceIdentifier(stage))
}

func (md_band_propertytype *MD_Band_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_band_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_band_type *MD_Band_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_band_type.GongGetReferenceIdentifier(stage))
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_contentinformation_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_coveragecontenttypecode_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_coveragedescription_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_coveragedescription_type.GongGetReferenceIdentifier(stage))
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_featurecataloguedescription_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_featurecataloguedescription_type.GongGetReferenceIdentifier(stage))
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_imagedescription_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_imagedescription_type.GongGetReferenceIdentifier(stage))
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_imagingconditioncode_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_rangedimension_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_rangedimension_type.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__content__%s__%08d_", s.GongGetGongstructName(), order)
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
