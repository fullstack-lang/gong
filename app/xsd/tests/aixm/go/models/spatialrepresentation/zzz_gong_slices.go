// generated code - do not edit
package spatialrepresentation

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
	// Compute reverse map for named struct MD_Georectified_Type
	// insertion point per field
	stage.MD_Georectified_Type_AxisDimensionProperties_reverseMap = make(map[*MD_Dimension_PropertyType]*MD_Georectified_Type)
	for md_georectified_type := range stage.MD_Georectified_Types {
		_ = md_georectified_type
		for _, _md_dimension_propertytype := range md_georectified_type.AxisDimensionProperties {
			stage.MD_Georectified_Type_AxisDimensionProperties_reverseMap[_md_dimension_propertytype] = md_georectified_type
		}
	}

	// Compute reverse map for named struct MD_Georeferenceable_Type
	// insertion point per field
	stage.MD_Georeferenceable_Type_AxisDimensionProperties_reverseMap = make(map[*MD_Dimension_PropertyType]*MD_Georeferenceable_Type)
	for md_georeferenceable_type := range stage.MD_Georeferenceable_Types {
		_ = md_georeferenceable_type
		for _, _md_dimension_propertytype := range md_georeferenceable_type.AxisDimensionProperties {
			stage.MD_Georeferenceable_Type_AxisDimensionProperties_reverseMap[_md_dimension_propertytype] = md_georeferenceable_type
		}
	}

	// Compute reverse map for named struct MD_GridSpatialRepresentation_Type
	// insertion point per field
	stage.MD_GridSpatialRepresentation_Type_AxisDimensionProperties_reverseMap = make(map[*MD_Dimension_PropertyType]*MD_GridSpatialRepresentation_Type)
	for md_gridspatialrepresentation_type := range stage.MD_GridSpatialRepresentation_Types {
		_ = md_gridspatialrepresentation_type
		for _, _md_dimension_propertytype := range md_gridspatialrepresentation_type.AxisDimensionProperties {
			stage.MD_GridSpatialRepresentation_Type_AxisDimensionProperties_reverseMap[_md_dimension_propertytype] = md_gridspatialrepresentation_type
		}
	}

	// Compute reverse map for named struct MD_VectorSpatialRepresentation_Type
	// insertion point per field
	stage.MD_VectorSpatialRepresentation_Type_GeometricObjects_reverseMap = make(map[*MD_GeometricObjects_PropertyType]*MD_VectorSpatialRepresentation_Type)
	for md_vectorspatialrepresentation_type := range stage.MD_VectorSpatialRepresentation_Types {
		_ = md_vectorspatialrepresentation_type
		for _, _md_geometricobjects_propertytype := range md_vectorspatialrepresentation_type.GeometricObjects {
			stage.MD_VectorSpatialRepresentation_Type_GeometricObjects_reverseMap[_md_geometricobjects_propertytype] = md_vectorspatialrepresentation_type
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AbstractMD_SpatialRepresentation_Types)

	res = __gong__appendInstances(res, stage.MD_CellGeometryCode_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_DimensionNameTypeCode_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_Dimension_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_Dimension_Types)

	res = __gong__appendInstances(res, stage.MD_GeometricObjectTypeCode_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_GeometricObjects_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_GeometricObjects_Types)

	res = __gong__appendInstances(res, stage.MD_Georectified_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_Georectified_Types)

	res = __gong__appendInstances(res, stage.MD_Georeferenceable_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_Georeferenceable_Types)

	res = __gong__appendInstances(res, stage.MD_GridSpatialRepresentation_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_GridSpatialRepresentation_Types)

	res = __gong__appendInstances(res, stage.MD_PixelOrientationCode_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_SpatialRepresentation_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_TopologyLevelCode_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_VectorSpatialRepresentation_PropertyTypes)

	res = __gong__appendInstances(res, stage.MD_VectorSpatialRepresentation_Types)

	return
}

// insertion point per named struct
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongCopy() GongstructIF {
	newInstance := new(AbstractMD_SpatialRepresentation_Type)
	abstractmd_spatialrepresentation_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_CellGeometryCode_PropertyType)
	md_cellgeometrycode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_DimensionNameTypeCode_PropertyType)
	md_dimensionnametypecode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_Dimension_PropertyType)
	md_dimension_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_dimension_type *MD_Dimension_Type) GongCopy() GongstructIF {
	newInstance := new(MD_Dimension_Type)
	md_dimension_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_GeometricObjectTypeCode_PropertyType)
	md_geometricobjecttypecode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_GeometricObjects_PropertyType)
	md_geometricobjects_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongCopy() GongstructIF {
	newInstance := new(MD_GeometricObjects_Type)
	md_geometricobjects_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_Georectified_PropertyType)
	md_georectified_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_georectified_type *MD_Georectified_Type) GongCopy() GongstructIF {
	newInstance := new(MD_Georectified_Type)
	md_georectified_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_Georeferenceable_PropertyType)
	md_georeferenceable_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongCopy() GongstructIF {
	newInstance := new(MD_Georeferenceable_Type)
	md_georeferenceable_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_GridSpatialRepresentation_PropertyType)
	md_gridspatialrepresentation_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongCopy() GongstructIF {
	newInstance := new(MD_GridSpatialRepresentation_Type)
	md_gridspatialrepresentation_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_PixelOrientationCode_PropertyType)
	md_pixelorientationcode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_SpatialRepresentation_PropertyType)
	md_spatialrepresentation_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_TopologyLevelCode_PropertyType)
	md_topologylevelcode_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongCopy() GongstructIF {
	newInstance := new(MD_VectorSpatialRepresentation_PropertyType)
	md_vectorspatialrepresentation_propertytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongCopy() GongstructIF {
	newInstance := new(MD_VectorSpatialRepresentation_Type)
	md_vectorspatialrepresentation_type.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, abstractmd_spatialrepresentation_type)
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_cellgeometrycode_propertytype)
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_dimensionnametypecode_propertytype)
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_dimension_propertytype)
}

func (md_dimension_type *MD_Dimension_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_dimension_type)
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_geometricobjecttypecode_propertytype)
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_geometricobjects_propertytype)
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_geometricobjects_type)
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_georectified_propertytype)
}

func (md_georectified_type *MD_Georectified_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_georectified_type)
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_georeferenceable_propertytype)
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_georeferenceable_type)
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_gridspatialrepresentation_propertytype)
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_gridspatialrepresentation_type)
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_pixelorientationcode_propertytype)
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_spatialrepresentation_propertytype)
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_topologylevelcode_propertytype)
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_vectorspatialrepresentation_propertytype)
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, md_vectorspatialrepresentation_type)
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
		stage.AbstractMD_SpatialRepresentation_Types,
		stage.AbstractMD_SpatialRepresentation_Type_stagedOrder,
		stage.AbstractMD_SpatialRepresentation_Types_reference,
		&stage.AbstractMD_SpatialRepresentation_Types_referenceOrder,
		stage.AbstractMD_SpatialRepresentation_Types_instance,
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
		stage.MD_CellGeometryCode_PropertyTypes,
		stage.MD_CellGeometryCode_PropertyType_stagedOrder,
		stage.MD_CellGeometryCode_PropertyTypes_reference,
		&stage.MD_CellGeometryCode_PropertyTypes_referenceOrder,
		stage.MD_CellGeometryCode_PropertyTypes_instance,
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
		stage.MD_DimensionNameTypeCode_PropertyTypes,
		stage.MD_DimensionNameTypeCode_PropertyType_stagedOrder,
		stage.MD_DimensionNameTypeCode_PropertyTypes_reference,
		&stage.MD_DimensionNameTypeCode_PropertyTypes_referenceOrder,
		stage.MD_DimensionNameTypeCode_PropertyTypes_instance,
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
		stage.MD_Dimension_PropertyTypes,
		stage.MD_Dimension_PropertyType_stagedOrder,
		stage.MD_Dimension_PropertyTypes_reference,
		&stage.MD_Dimension_PropertyTypes_referenceOrder,
		stage.MD_Dimension_PropertyTypes_instance,
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
		stage.MD_Dimension_Types,
		stage.MD_Dimension_Type_stagedOrder,
		stage.MD_Dimension_Types_reference,
		&stage.MD_Dimension_Types_referenceOrder,
		stage.MD_Dimension_Types_instance,
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
		stage.MD_GeometricObjectTypeCode_PropertyTypes,
		stage.MD_GeometricObjectTypeCode_PropertyType_stagedOrder,
		stage.MD_GeometricObjectTypeCode_PropertyTypes_reference,
		&stage.MD_GeometricObjectTypeCode_PropertyTypes_referenceOrder,
		stage.MD_GeometricObjectTypeCode_PropertyTypes_instance,
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
		stage.MD_GeometricObjects_PropertyTypes,
		stage.MD_GeometricObjects_PropertyType_stagedOrder,
		stage.MD_GeometricObjects_PropertyTypes_reference,
		&stage.MD_GeometricObjects_PropertyTypes_referenceOrder,
		stage.MD_GeometricObjects_PropertyTypes_instance,
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
		stage.MD_GeometricObjects_Types,
		stage.MD_GeometricObjects_Type_stagedOrder,
		stage.MD_GeometricObjects_Types_reference,
		&stage.MD_GeometricObjects_Types_referenceOrder,
		stage.MD_GeometricObjects_Types_instance,
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
		stage.MD_Georectified_PropertyTypes,
		stage.MD_Georectified_PropertyType_stagedOrder,
		stage.MD_Georectified_PropertyTypes_reference,
		&stage.MD_Georectified_PropertyTypes_referenceOrder,
		stage.MD_Georectified_PropertyTypes_instance,
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
		stage.MD_Georectified_Types,
		stage.MD_Georectified_Type_stagedOrder,
		stage.MD_Georectified_Types_reference,
		&stage.MD_Georectified_Types_referenceOrder,
		stage.MD_Georectified_Types_instance,
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
		stage.MD_Georeferenceable_PropertyTypes,
		stage.MD_Georeferenceable_PropertyType_stagedOrder,
		stage.MD_Georeferenceable_PropertyTypes_reference,
		&stage.MD_Georeferenceable_PropertyTypes_referenceOrder,
		stage.MD_Georeferenceable_PropertyTypes_instance,
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
		stage.MD_Georeferenceable_Types,
		stage.MD_Georeferenceable_Type_stagedOrder,
		stage.MD_Georeferenceable_Types_reference,
		&stage.MD_Georeferenceable_Types_referenceOrder,
		stage.MD_Georeferenceable_Types_instance,
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
		stage.MD_GridSpatialRepresentation_PropertyTypes,
		stage.MD_GridSpatialRepresentation_PropertyType_stagedOrder,
		stage.MD_GridSpatialRepresentation_PropertyTypes_reference,
		&stage.MD_GridSpatialRepresentation_PropertyTypes_referenceOrder,
		stage.MD_GridSpatialRepresentation_PropertyTypes_instance,
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
		stage.MD_GridSpatialRepresentation_Types,
		stage.MD_GridSpatialRepresentation_Type_stagedOrder,
		stage.MD_GridSpatialRepresentation_Types_reference,
		&stage.MD_GridSpatialRepresentation_Types_referenceOrder,
		stage.MD_GridSpatialRepresentation_Types_instance,
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
		stage.MD_PixelOrientationCode_PropertyTypes,
		stage.MD_PixelOrientationCode_PropertyType_stagedOrder,
		stage.MD_PixelOrientationCode_PropertyTypes_reference,
		&stage.MD_PixelOrientationCode_PropertyTypes_referenceOrder,
		stage.MD_PixelOrientationCode_PropertyTypes_instance,
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
		stage.MD_SpatialRepresentation_PropertyTypes,
		stage.MD_SpatialRepresentation_PropertyType_stagedOrder,
		stage.MD_SpatialRepresentation_PropertyTypes_reference,
		&stage.MD_SpatialRepresentation_PropertyTypes_referenceOrder,
		stage.MD_SpatialRepresentation_PropertyTypes_instance,
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
		stage.MD_TopologyLevelCode_PropertyTypes,
		stage.MD_TopologyLevelCode_PropertyType_stagedOrder,
		stage.MD_TopologyLevelCode_PropertyTypes_reference,
		&stage.MD_TopologyLevelCode_PropertyTypes_referenceOrder,
		stage.MD_TopologyLevelCode_PropertyTypes_instance,
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
		stage.MD_VectorSpatialRepresentation_PropertyTypes,
		stage.MD_VectorSpatialRepresentation_PropertyType_stagedOrder,
		stage.MD_VectorSpatialRepresentation_PropertyTypes_reference,
		&stage.MD_VectorSpatialRepresentation_PropertyTypes_referenceOrder,
		stage.MD_VectorSpatialRepresentation_PropertyTypes_instance,
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
		stage.MD_VectorSpatialRepresentation_Types,
		stage.MD_VectorSpatialRepresentation_Type_stagedOrder,
		stage.MD_VectorSpatialRepresentation_Types_reference,
		&stage.MD_VectorSpatialRepresentation_Types_referenceOrder,
		stage.MD_VectorSpatialRepresentation_Types_instance,
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
	__gong__computeReferencePass1(stage, stage.AbstractMD_SpatialRepresentation_Types, &stage.AbstractMD_SpatialRepresentation_Types_reference, &stage.AbstractMD_SpatialRepresentation_Types_referenceOrder, &stage.AbstractMD_SpatialRepresentation_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_CellGeometryCode_PropertyTypes, &stage.MD_CellGeometryCode_PropertyTypes_reference, &stage.MD_CellGeometryCode_PropertyTypes_referenceOrder, &stage.MD_CellGeometryCode_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_DimensionNameTypeCode_PropertyTypes, &stage.MD_DimensionNameTypeCode_PropertyTypes_reference, &stage.MD_DimensionNameTypeCode_PropertyTypes_referenceOrder, &stage.MD_DimensionNameTypeCode_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_Dimension_PropertyTypes, &stage.MD_Dimension_PropertyTypes_reference, &stage.MD_Dimension_PropertyTypes_referenceOrder, &stage.MD_Dimension_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_Dimension_Types, &stage.MD_Dimension_Types_reference, &stage.MD_Dimension_Types_referenceOrder, &stage.MD_Dimension_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_GeometricObjectTypeCode_PropertyTypes, &stage.MD_GeometricObjectTypeCode_PropertyTypes_reference, &stage.MD_GeometricObjectTypeCode_PropertyTypes_referenceOrder, &stage.MD_GeometricObjectTypeCode_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_GeometricObjects_PropertyTypes, &stage.MD_GeometricObjects_PropertyTypes_reference, &stage.MD_GeometricObjects_PropertyTypes_referenceOrder, &stage.MD_GeometricObjects_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_GeometricObjects_Types, &stage.MD_GeometricObjects_Types_reference, &stage.MD_GeometricObjects_Types_referenceOrder, &stage.MD_GeometricObjects_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_Georectified_PropertyTypes, &stage.MD_Georectified_PropertyTypes_reference, &stage.MD_Georectified_PropertyTypes_referenceOrder, &stage.MD_Georectified_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_Georectified_Types, &stage.MD_Georectified_Types_reference, &stage.MD_Georectified_Types_referenceOrder, &stage.MD_Georectified_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_Georeferenceable_PropertyTypes, &stage.MD_Georeferenceable_PropertyTypes_reference, &stage.MD_Georeferenceable_PropertyTypes_referenceOrder, &stage.MD_Georeferenceable_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_Georeferenceable_Types, &stage.MD_Georeferenceable_Types_reference, &stage.MD_Georeferenceable_Types_referenceOrder, &stage.MD_Georeferenceable_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_GridSpatialRepresentation_PropertyTypes, &stage.MD_GridSpatialRepresentation_PropertyTypes_reference, &stage.MD_GridSpatialRepresentation_PropertyTypes_referenceOrder, &stage.MD_GridSpatialRepresentation_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_GridSpatialRepresentation_Types, &stage.MD_GridSpatialRepresentation_Types_reference, &stage.MD_GridSpatialRepresentation_Types_referenceOrder, &stage.MD_GridSpatialRepresentation_Types_instance)

	__gong__computeReferencePass1(stage, stage.MD_PixelOrientationCode_PropertyTypes, &stage.MD_PixelOrientationCode_PropertyTypes_reference, &stage.MD_PixelOrientationCode_PropertyTypes_referenceOrder, &stage.MD_PixelOrientationCode_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_SpatialRepresentation_PropertyTypes, &stage.MD_SpatialRepresentation_PropertyTypes_reference, &stage.MD_SpatialRepresentation_PropertyTypes_referenceOrder, &stage.MD_SpatialRepresentation_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_TopologyLevelCode_PropertyTypes, &stage.MD_TopologyLevelCode_PropertyTypes_reference, &stage.MD_TopologyLevelCode_PropertyTypes_referenceOrder, &stage.MD_TopologyLevelCode_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_VectorSpatialRepresentation_PropertyTypes, &stage.MD_VectorSpatialRepresentation_PropertyTypes_reference, &stage.MD_VectorSpatialRepresentation_PropertyTypes_referenceOrder, &stage.MD_VectorSpatialRepresentation_PropertyTypes_instance)

	__gong__computeReferencePass1(stage, stage.MD_VectorSpatialRepresentation_Types, &stage.MD_VectorSpatialRepresentation_Types_reference, &stage.MD_VectorSpatialRepresentation_Types_referenceOrder, &stage.MD_VectorSpatialRepresentation_Types_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AbstractMD_SpatialRepresentation_Types, stage.AbstractMD_SpatialRepresentation_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_CellGeometryCode_PropertyTypes, stage.MD_CellGeometryCode_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_DimensionNameTypeCode_PropertyTypes, stage.MD_DimensionNameTypeCode_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_Dimension_PropertyTypes, stage.MD_Dimension_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_Dimension_Types, stage.MD_Dimension_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_GeometricObjectTypeCode_PropertyTypes, stage.MD_GeometricObjectTypeCode_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_GeometricObjects_PropertyTypes, stage.MD_GeometricObjects_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_GeometricObjects_Types, stage.MD_GeometricObjects_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_Georectified_PropertyTypes, stage.MD_Georectified_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_Georectified_Types, stage.MD_Georectified_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_Georeferenceable_PropertyTypes, stage.MD_Georeferenceable_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_Georeferenceable_Types, stage.MD_Georeferenceable_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_GridSpatialRepresentation_PropertyTypes, stage.MD_GridSpatialRepresentation_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_GridSpatialRepresentation_Types, stage.MD_GridSpatialRepresentation_Types_reference, stage)

	__gong__computeReferencePass2(stage.MD_PixelOrientationCode_PropertyTypes, stage.MD_PixelOrientationCode_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_SpatialRepresentation_PropertyTypes, stage.MD_SpatialRepresentation_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_TopologyLevelCode_PropertyTypes, stage.MD_TopologyLevelCode_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_VectorSpatialRepresentation_PropertyTypes, stage.MD_VectorSpatialRepresentation_PropertyTypes_reference, stage)

	__gong__computeReferencePass2(stage.MD_VectorSpatialRepresentation_Types, stage.MD_VectorSpatialRepresentation_Types_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AbstractMD_SpatialRepresentation_Type_stagedOrder, stage.AbstractMD_SpatialRepresentation_Types_referenceOrder, abstractmd_spatialrepresentation_type, "AbstractMD_SpatialRepresentation_Type")
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_CellGeometryCode_PropertyType_stagedOrder, stage.MD_CellGeometryCode_PropertyTypes_referenceOrder, md_cellgeometrycode_propertytype, "MD_CellGeometryCode_PropertyType")
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_DimensionNameTypeCode_PropertyType_stagedOrder, stage.MD_DimensionNameTypeCode_PropertyTypes_referenceOrder, md_dimensionnametypecode_propertytype, "MD_DimensionNameTypeCode_PropertyType")
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Dimension_PropertyType_stagedOrder, stage.MD_Dimension_PropertyTypes_referenceOrder, md_dimension_propertytype, "MD_Dimension_PropertyType")
}

func (md_dimension_type *MD_Dimension_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Dimension_Type_stagedOrder, stage.MD_Dimension_Types_referenceOrder, md_dimension_type, "MD_Dimension_Type")
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_GeometricObjectTypeCode_PropertyType_stagedOrder, stage.MD_GeometricObjectTypeCode_PropertyTypes_referenceOrder, md_geometricobjecttypecode_propertytype, "MD_GeometricObjectTypeCode_PropertyType")
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_GeometricObjects_PropertyType_stagedOrder, stage.MD_GeometricObjects_PropertyTypes_referenceOrder, md_geometricobjects_propertytype, "MD_GeometricObjects_PropertyType")
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_GeometricObjects_Type_stagedOrder, stage.MD_GeometricObjects_Types_referenceOrder, md_geometricobjects_type, "MD_GeometricObjects_Type")
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Georectified_PropertyType_stagedOrder, stage.MD_Georectified_PropertyTypes_referenceOrder, md_georectified_propertytype, "MD_Georectified_PropertyType")
}

func (md_georectified_type *MD_Georectified_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Georectified_Type_stagedOrder, stage.MD_Georectified_Types_referenceOrder, md_georectified_type, "MD_Georectified_Type")
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Georeferenceable_PropertyType_stagedOrder, stage.MD_Georeferenceable_PropertyTypes_referenceOrder, md_georeferenceable_propertytype, "MD_Georeferenceable_PropertyType")
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_Georeferenceable_Type_stagedOrder, stage.MD_Georeferenceable_Types_referenceOrder, md_georeferenceable_type, "MD_Georeferenceable_Type")
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_GridSpatialRepresentation_PropertyType_stagedOrder, stage.MD_GridSpatialRepresentation_PropertyTypes_referenceOrder, md_gridspatialrepresentation_propertytype, "MD_GridSpatialRepresentation_PropertyType")
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_GridSpatialRepresentation_Type_stagedOrder, stage.MD_GridSpatialRepresentation_Types_referenceOrder, md_gridspatialrepresentation_type, "MD_GridSpatialRepresentation_Type")
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_PixelOrientationCode_PropertyType_stagedOrder, stage.MD_PixelOrientationCode_PropertyTypes_referenceOrder, md_pixelorientationcode_propertytype, "MD_PixelOrientationCode_PropertyType")
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_SpatialRepresentation_PropertyType_stagedOrder, stage.MD_SpatialRepresentation_PropertyTypes_referenceOrder, md_spatialrepresentation_propertytype, "MD_SpatialRepresentation_PropertyType")
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_TopologyLevelCode_PropertyType_stagedOrder, stage.MD_TopologyLevelCode_PropertyTypes_referenceOrder, md_topologylevelcode_propertytype, "MD_TopologyLevelCode_PropertyType")
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_VectorSpatialRepresentation_PropertyType_stagedOrder, stage.MD_VectorSpatialRepresentation_PropertyTypes_referenceOrder, md_vectorspatialrepresentation_propertytype, "MD_VectorSpatialRepresentation_PropertyType")
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.MD_VectorSpatialRepresentation_Type_stagedOrder, stage.MD_VectorSpatialRepresentation_Types_referenceOrder, md_vectorspatialrepresentation_type, "MD_VectorSpatialRepresentation_Type")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(abstractmd_spatialrepresentation_type, abstractmd_spatialrepresentation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return abstractmd_spatialrepresentation_type.GongGetIdentifier(stage)
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_cellgeometrycode_propertytype, md_cellgeometrycode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_cellgeometrycode_propertytype.GongGetIdentifier(stage)
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_dimensionnametypecode_propertytype, md_dimensionnametypecode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_dimensionnametypecode_propertytype.GongGetIdentifier(stage)
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_dimension_propertytype, md_dimension_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_dimension_propertytype *MD_Dimension_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_dimension_propertytype.GongGetIdentifier(stage)
}

func (md_dimension_type *MD_Dimension_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_dimension_type, md_dimension_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_dimension_type *MD_Dimension_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_dimension_type.GongGetIdentifier(stage)
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_geometricobjecttypecode_propertytype, md_geometricobjecttypecode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_geometricobjecttypecode_propertytype.GongGetIdentifier(stage)
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_geometricobjects_propertytype, md_geometricobjects_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_geometricobjects_propertytype.GongGetIdentifier(stage)
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_geometricobjects_type, md_geometricobjects_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_geometricobjects_type *MD_GeometricObjects_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_geometricobjects_type.GongGetIdentifier(stage)
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_georectified_propertytype, md_georectified_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_georectified_propertytype *MD_Georectified_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_georectified_propertytype.GongGetIdentifier(stage)
}

func (md_georectified_type *MD_Georectified_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_georectified_type, md_georectified_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_georectified_type *MD_Georectified_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_georectified_type.GongGetIdentifier(stage)
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_georeferenceable_propertytype, md_georeferenceable_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_georeferenceable_propertytype.GongGetIdentifier(stage)
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_georeferenceable_type, md_georeferenceable_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_georeferenceable_type *MD_Georeferenceable_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_georeferenceable_type.GongGetIdentifier(stage)
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_gridspatialrepresentation_propertytype, md_gridspatialrepresentation_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_gridspatialrepresentation_propertytype.GongGetIdentifier(stage)
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_gridspatialrepresentation_type, md_gridspatialrepresentation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_gridspatialrepresentation_type.GongGetIdentifier(stage)
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_pixelorientationcode_propertytype, md_pixelorientationcode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_pixelorientationcode_propertytype.GongGetIdentifier(stage)
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_spatialrepresentation_propertytype, md_spatialrepresentation_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_spatialrepresentation_propertytype.GongGetIdentifier(stage)
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_topologylevelcode_propertytype, md_topologylevelcode_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_topologylevelcode_propertytype.GongGetIdentifier(stage)
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_vectorspatialrepresentation_propertytype, md_vectorspatialrepresentation_propertytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongGetReferenceIdentifier(stage *Stage) string {
	return md_vectorspatialrepresentation_propertytype.GongGetIdentifier(stage)
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(md_vectorspatialrepresentation_type, md_vectorspatialrepresentation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongGetReferenceIdentifier(stage *Stage) string {
	return md_vectorspatialrepresentation_type.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(abstractmd_spatialrepresentation_type.GongGetIdentifier(stage), "AbstractMD_SpatialRepresentation_Type", abstractmd_spatialrepresentation_type.Name)
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_cellgeometrycode_propertytype.GongGetIdentifier(stage), "MD_CellGeometryCode_PropertyType", md_cellgeometrycode_propertytype.Name)
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_dimensionnametypecode_propertytype.GongGetIdentifier(stage), "MD_DimensionNameTypeCode_PropertyType", md_dimensionnametypecode_propertytype.Name)
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_dimension_propertytype.GongGetIdentifier(stage), "MD_Dimension_PropertyType", md_dimension_propertytype.Name)
}

func (md_dimension_type *MD_Dimension_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_dimension_type.GongGetIdentifier(stage), "MD_Dimension_Type", md_dimension_type.Name)
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_geometricobjecttypecode_propertytype.GongGetIdentifier(stage), "MD_GeometricObjectTypeCode_PropertyType", md_geometricobjecttypecode_propertytype.Name)
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_geometricobjects_propertytype.GongGetIdentifier(stage), "MD_GeometricObjects_PropertyType", md_geometricobjects_propertytype.Name)
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_geometricobjects_type.GongGetIdentifier(stage), "MD_GeometricObjects_Type", md_geometricobjects_type.Name)
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_georectified_propertytype.GongGetIdentifier(stage), "MD_Georectified_PropertyType", md_georectified_propertytype.Name)
}

func (md_georectified_type *MD_Georectified_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_georectified_type.GongGetIdentifier(stage), "MD_Georectified_Type", md_georectified_type.Name)
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_georeferenceable_propertytype.GongGetIdentifier(stage), "MD_Georeferenceable_PropertyType", md_georeferenceable_propertytype.Name)
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_georeferenceable_type.GongGetIdentifier(stage), "MD_Georeferenceable_Type", md_georeferenceable_type.Name)
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_gridspatialrepresentation_propertytype.GongGetIdentifier(stage), "MD_GridSpatialRepresentation_PropertyType", md_gridspatialrepresentation_propertytype.Name)
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_gridspatialrepresentation_type.GongGetIdentifier(stage), "MD_GridSpatialRepresentation_Type", md_gridspatialrepresentation_type.Name)
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_pixelorientationcode_propertytype.GongGetIdentifier(stage), "MD_PixelOrientationCode_PropertyType", md_pixelorientationcode_propertytype.Name)
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_spatialrepresentation_propertytype.GongGetIdentifier(stage), "MD_SpatialRepresentation_PropertyType", md_spatialrepresentation_propertytype.Name)
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_topologylevelcode_propertytype.GongGetIdentifier(stage), "MD_TopologyLevelCode_PropertyType", md_topologylevelcode_propertytype.Name)
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_vectorspatialrepresentation_propertytype.GongGetIdentifier(stage), "MD_VectorSpatialRepresentation_PropertyType", md_vectorspatialrepresentation_propertytype.Name)
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(md_vectorspatialrepresentation_type.GongGetIdentifier(stage), "MD_VectorSpatialRepresentation_Type", md_vectorspatialrepresentation_type.Name)
}

// insertion point for unstaging
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(abstractmd_spatialrepresentation_type.GongGetReferenceIdentifier(stage))
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_cellgeometrycode_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_dimensionnametypecode_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_dimension_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_dimension_type *MD_Dimension_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_dimension_type.GongGetReferenceIdentifier(stage))
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_geometricobjecttypecode_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_geometricobjects_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_geometricobjects_type.GongGetReferenceIdentifier(stage))
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_georectified_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_georectified_type *MD_Georectified_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_georectified_type.GongGetReferenceIdentifier(stage))
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_georeferenceable_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_georeferenceable_type.GongGetReferenceIdentifier(stage))
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_gridspatialrepresentation_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_gridspatialrepresentation_type.GongGetReferenceIdentifier(stage))
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_pixelorientationcode_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_spatialrepresentation_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_topologylevelcode_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_vectorspatialrepresentation_propertytype.GongGetReferenceIdentifier(stage))
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(md_vectorspatialrepresentation_type.GongGetReferenceIdentifier(stage))
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
	return fmt.Sprintf("__spatialrepresentation__%s__%08d_", s.GongGetGongstructName(), order)
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
