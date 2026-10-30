// generated code - do not edit
package spatialrepresentation

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
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractMD_SpatialRepresentation_Types[abstractmd_spatialrepresentation_type]
	return ok
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_CellGeometryCode_PropertyTypes[md_cellgeometrycode_propertytype]
	return ok
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_DimensionNameTypeCode_PropertyTypes[md_dimensionnametypecode_propertytype]
	return ok
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Dimension_PropertyTypes[md_dimension_propertytype]
	return ok
}

func (md_dimension_type *MD_Dimension_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Dimension_Types[md_dimension_type]
	return ok
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_GeometricObjectTypeCode_PropertyTypes[md_geometricobjecttypecode_propertytype]
	return ok
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_GeometricObjects_PropertyTypes[md_geometricobjects_propertytype]
	return ok
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_GeometricObjects_Types[md_geometricobjects_type]
	return ok
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Georectified_PropertyTypes[md_georectified_propertytype]
	return ok
}

func (md_georectified_type *MD_Georectified_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Georectified_Types[md_georectified_type]
	return ok
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Georeferenceable_PropertyTypes[md_georeferenceable_propertytype]
	return ok
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Georeferenceable_Types[md_georeferenceable_type]
	return ok
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_GridSpatialRepresentation_PropertyTypes[md_gridspatialrepresentation_propertytype]
	return ok
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_GridSpatialRepresentation_Types[md_gridspatialrepresentation_type]
	return ok
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_PixelOrientationCode_PropertyTypes[md_pixelorientationcode_propertytype]
	return ok
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_SpatialRepresentation_PropertyTypes[md_spatialrepresentation_propertytype]
	return ok
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_TopologyLevelCode_PropertyTypes[md_topologylevelcode_propertytype]
	return ok
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_VectorSpatialRepresentation_PropertyTypes[md_vectorspatialrepresentation_propertytype]
	return ok
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_VectorSpatialRepresentation_Types[md_vectorspatialrepresentation_type]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractmd_spatialrepresentation_type) {
		return
	}

	abstractmd_spatialrepresentation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_cellgeometrycode_propertytype) {
		return
	}

	md_cellgeometrycode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_dimensionnametypecode_propertytype) {
		return
	}

	md_dimensionnametypecode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_dimension_propertytype) {
		return
	}

	md_dimension_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_dimension_propertytype.MD_Dimension != nil {
		stage.StageBranch(md_dimension_propertytype.MD_Dimension)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_dimension_type *MD_Dimension_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_dimension_type) {
		return
	}

	md_dimension_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_dimension_type.DimensionName != nil {
		stage.StageBranch(md_dimension_type.DimensionName)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_geometricobjecttypecode_propertytype) {
		return
	}

	md_geometricobjecttypecode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_geometricobjects_propertytype) {
		return
	}

	md_geometricobjects_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_geometricobjects_propertytype.MD_GeometricObjects != nil {
		stage.StageBranch(md_geometricobjects_propertytype.MD_GeometricObjects)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_geometricobjects_type) {
		return
	}

	md_geometricobjects_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_geometricobjects_type.GeometricObjectType != nil {
		stage.StageBranch(md_geometricobjects_type.GeometricObjectType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_georectified_propertytype) {
		return
	}

	md_georectified_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_georectified_propertytype.MD_Georectified != nil {
		stage.StageBranch(md_georectified_propertytype.MD_Georectified)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_georectified_type *MD_Georectified_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_georectified_type) {
		return
	}

	md_georectified_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_georectified_type.CellGeometry != nil {
		stage.StageBranch(md_georectified_type.CellGeometry)
	}
	if md_georectified_type.PointInPixel != nil {
		stage.StageBranch(md_georectified_type.PointInPixel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_georectified_type.AxisDimensionProperties {
		stage.StageBranch(_md_dimension_propertytype)
	}

}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_georeferenceable_propertytype) {
		return
	}

	md_georeferenceable_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_georeferenceable_propertytype.MD_Georeferenceable != nil {
		stage.StageBranch(md_georeferenceable_propertytype.MD_Georeferenceable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_georeferenceable_type) {
		return
	}

	md_georeferenceable_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_georeferenceable_type.CellGeometry != nil {
		stage.StageBranch(md_georeferenceable_type.CellGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_georeferenceable_type.AxisDimensionProperties {
		stage.StageBranch(_md_dimension_propertytype)
	}

}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_gridspatialrepresentation_propertytype) {
		return
	}

	md_gridspatialrepresentation_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
		stage.StageBranch(md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation)
	}
	if md_gridspatialrepresentation_propertytype.MD_Georeferenceable != nil {
		stage.StageBranch(md_gridspatialrepresentation_propertytype.MD_Georeferenceable)
	}
	if md_gridspatialrepresentation_propertytype.MD_Georectified != nil {
		stage.StageBranch(md_gridspatialrepresentation_propertytype.MD_Georectified)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_gridspatialrepresentation_type) {
		return
	}

	md_gridspatialrepresentation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_gridspatialrepresentation_type.CellGeometry != nil {
		stage.StageBranch(md_gridspatialrepresentation_type.CellGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_gridspatialrepresentation_type.AxisDimensionProperties {
		stage.StageBranch(_md_dimension_propertytype)
	}

}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_pixelorientationcode_propertytype) {
		return
	}

	md_pixelorientationcode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_spatialrepresentation_propertytype) {
		return
	}

	md_spatialrepresentation_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
		stage.StageBranch(md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation)
	}
	if md_spatialrepresentation_propertytype.MD_Georeferenceable != nil {
		stage.StageBranch(md_spatialrepresentation_propertytype.MD_Georeferenceable)
	}
	if md_spatialrepresentation_propertytype.MD_Georectified != nil {
		stage.StageBranch(md_spatialrepresentation_propertytype.MD_Georectified)
	}
	if md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
		stage.StageBranch(md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_topologylevelcode_propertytype) {
		return
	}

	md_topologylevelcode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_vectorspatialrepresentation_propertytype) {
		return
	}

	md_vectorspatialrepresentation_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
		stage.StageBranch(md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_vectorspatialrepresentation_type) {
		return
	}

	md_vectorspatialrepresentation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_vectorspatialrepresentation_type.TopologyLevel != nil {
		stage.StageBranch(md_vectorspatialrepresentation_type.TopologyLevel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_geometricobjects_propertytype := range md_vectorspatialrepresentation_type.GeometricObjects {
		stage.StageBranch(_md_geometricobjects_propertytype)
	}

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
	case *AbstractMD_SpatialRepresentation_Type:
		toT := GongCopyBranchAbstractMD_SpatialRepresentation_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_CellGeometryCode_PropertyType:
		toT := GongCopyBranchMD_CellGeometryCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_DimensionNameTypeCode_PropertyType:
		toT := GongCopyBranchMD_DimensionNameTypeCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Dimension_PropertyType:
		toT := GongCopyBranchMD_Dimension_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Dimension_Type:
		toT := GongCopyBranchMD_Dimension_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_GeometricObjectTypeCode_PropertyType:
		toT := GongCopyBranchMD_GeometricObjectTypeCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_GeometricObjects_PropertyType:
		toT := GongCopyBranchMD_GeometricObjects_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_GeometricObjects_Type:
		toT := GongCopyBranchMD_GeometricObjects_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Georectified_PropertyType:
		toT := GongCopyBranchMD_Georectified_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Georectified_Type:
		toT := GongCopyBranchMD_Georectified_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Georeferenceable_PropertyType:
		toT := GongCopyBranchMD_Georeferenceable_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Georeferenceable_Type:
		toT := GongCopyBranchMD_Georeferenceable_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_GridSpatialRepresentation_PropertyType:
		toT := GongCopyBranchMD_GridSpatialRepresentation_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_GridSpatialRepresentation_Type:
		toT := GongCopyBranchMD_GridSpatialRepresentation_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_PixelOrientationCode_PropertyType:
		toT := GongCopyBranchMD_PixelOrientationCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_SpatialRepresentation_PropertyType:
		toT := GongCopyBranchMD_SpatialRepresentation_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_TopologyLevelCode_PropertyType:
		toT := GongCopyBranchMD_TopologyLevelCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_VectorSpatialRepresentation_PropertyType:
		toT := GongCopyBranchMD_VectorSpatialRepresentation_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_VectorSpatialRepresentation_Type:
		toT := GongCopyBranchMD_VectorSpatialRepresentation_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAbstractMD_SpatialRepresentation_Type(mapOrigCopy map[any]any, abstractmd_spatialrepresentation_typeFrom *AbstractMD_SpatialRepresentation_Type) (abstractmd_spatialrepresentation_typeTo *AbstractMD_SpatialRepresentation_Type) {
	var alreadyCopied bool
	abstractmd_spatialrepresentation_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractmd_spatialrepresentation_typeFrom)
	if alreadyCopied {
		return
	}
	abstractmd_spatialrepresentation_typeFrom.GongCopyBasicFields(abstractmd_spatialrepresentation_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_CellGeometryCode_PropertyType(mapOrigCopy map[any]any, md_cellgeometrycode_propertytypeFrom *MD_CellGeometryCode_PropertyType) (md_cellgeometrycode_propertytypeTo *MD_CellGeometryCode_PropertyType) {
	var alreadyCopied bool
	md_cellgeometrycode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_cellgeometrycode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_cellgeometrycode_propertytypeFrom.GongCopyBasicFields(md_cellgeometrycode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_DimensionNameTypeCode_PropertyType(mapOrigCopy map[any]any, md_dimensionnametypecode_propertytypeFrom *MD_DimensionNameTypeCode_PropertyType) (md_dimensionnametypecode_propertytypeTo *MD_DimensionNameTypeCode_PropertyType) {
	var alreadyCopied bool
	md_dimensionnametypecode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_dimensionnametypecode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_dimensionnametypecode_propertytypeFrom.GongCopyBasicFields(md_dimensionnametypecode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_Dimension_PropertyType(mapOrigCopy map[any]any, md_dimension_propertytypeFrom *MD_Dimension_PropertyType) (md_dimension_propertytypeTo *MD_Dimension_PropertyType) {
	var alreadyCopied bool
	md_dimension_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_dimension_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_dimension_propertytypeFrom.GongCopyBasicFields(md_dimension_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_dimension_propertytypeFrom.MD_Dimension != nil {
		md_dimension_propertytypeTo.MD_Dimension = GongCopyBranchMD_Dimension_Type(mapOrigCopy, md_dimension_propertytypeFrom.MD_Dimension)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_Dimension_Type(mapOrigCopy map[any]any, md_dimension_typeFrom *MD_Dimension_Type) (md_dimension_typeTo *MD_Dimension_Type) {
	var alreadyCopied bool
	md_dimension_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_dimension_typeFrom)
	if alreadyCopied {
		return
	}
	md_dimension_typeFrom.GongCopyBasicFields(md_dimension_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_dimension_typeFrom.DimensionName != nil {
		md_dimension_typeTo.DimensionName = GongCopyBranchMD_DimensionNameTypeCode_PropertyType(mapOrigCopy, md_dimension_typeFrom.DimensionName)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_GeometricObjectTypeCode_PropertyType(mapOrigCopy map[any]any, md_geometricobjecttypecode_propertytypeFrom *MD_GeometricObjectTypeCode_PropertyType) (md_geometricobjecttypecode_propertytypeTo *MD_GeometricObjectTypeCode_PropertyType) {
	var alreadyCopied bool
	md_geometricobjecttypecode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_geometricobjecttypecode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_geometricobjecttypecode_propertytypeFrom.GongCopyBasicFields(md_geometricobjecttypecode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_GeometricObjects_PropertyType(mapOrigCopy map[any]any, md_geometricobjects_propertytypeFrom *MD_GeometricObjects_PropertyType) (md_geometricobjects_propertytypeTo *MD_GeometricObjects_PropertyType) {
	var alreadyCopied bool
	md_geometricobjects_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_geometricobjects_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_geometricobjects_propertytypeFrom.GongCopyBasicFields(md_geometricobjects_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_geometricobjects_propertytypeFrom.MD_GeometricObjects != nil {
		md_geometricobjects_propertytypeTo.MD_GeometricObjects = GongCopyBranchMD_GeometricObjects_Type(mapOrigCopy, md_geometricobjects_propertytypeFrom.MD_GeometricObjects)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_GeometricObjects_Type(mapOrigCopy map[any]any, md_geometricobjects_typeFrom *MD_GeometricObjects_Type) (md_geometricobjects_typeTo *MD_GeometricObjects_Type) {
	var alreadyCopied bool
	md_geometricobjects_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_geometricobjects_typeFrom)
	if alreadyCopied {
		return
	}
	md_geometricobjects_typeFrom.GongCopyBasicFields(md_geometricobjects_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_geometricobjects_typeFrom.GeometricObjectType != nil {
		md_geometricobjects_typeTo.GeometricObjectType = GongCopyBranchMD_GeometricObjectTypeCode_PropertyType(mapOrigCopy, md_geometricobjects_typeFrom.GeometricObjectType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_Georectified_PropertyType(mapOrigCopy map[any]any, md_georectified_propertytypeFrom *MD_Georectified_PropertyType) (md_georectified_propertytypeTo *MD_Georectified_PropertyType) {
	var alreadyCopied bool
	md_georectified_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_georectified_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_georectified_propertytypeFrom.GongCopyBasicFields(md_georectified_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_georectified_propertytypeFrom.MD_Georectified != nil {
		md_georectified_propertytypeTo.MD_Georectified = GongCopyBranchMD_Georectified_Type(mapOrigCopy, md_georectified_propertytypeFrom.MD_Georectified)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_Georectified_Type(mapOrigCopy map[any]any, md_georectified_typeFrom *MD_Georectified_Type) (md_georectified_typeTo *MD_Georectified_Type) {
	var alreadyCopied bool
	md_georectified_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_georectified_typeFrom)
	if alreadyCopied {
		return
	}
	md_georectified_typeFrom.GongCopyBasicFields(md_georectified_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_georectified_typeFrom.CellGeometry != nil {
		md_georectified_typeTo.CellGeometry = GongCopyBranchMD_CellGeometryCode_PropertyType(mapOrigCopy, md_georectified_typeFrom.CellGeometry)
	}
	if md_georectified_typeFrom.PointInPixel != nil {
		md_georectified_typeTo.PointInPixel = GongCopyBranchMD_PixelOrientationCode_PropertyType(mapOrigCopy, md_georectified_typeFrom.PointInPixel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_georectified_typeFrom.AxisDimensionProperties {
		md_georectified_typeTo.AxisDimensionProperties = append(md_georectified_typeTo.AxisDimensionProperties, GongCopyBranchMD_Dimension_PropertyType(mapOrigCopy, _md_dimension_propertytype))
	}

	return
}

func GongCopyBranchMD_Georeferenceable_PropertyType(mapOrigCopy map[any]any, md_georeferenceable_propertytypeFrom *MD_Georeferenceable_PropertyType) (md_georeferenceable_propertytypeTo *MD_Georeferenceable_PropertyType) {
	var alreadyCopied bool
	md_georeferenceable_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_georeferenceable_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_georeferenceable_propertytypeFrom.GongCopyBasicFields(md_georeferenceable_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_georeferenceable_propertytypeFrom.MD_Georeferenceable != nil {
		md_georeferenceable_propertytypeTo.MD_Georeferenceable = GongCopyBranchMD_Georeferenceable_Type(mapOrigCopy, md_georeferenceable_propertytypeFrom.MD_Georeferenceable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_Georeferenceable_Type(mapOrigCopy map[any]any, md_georeferenceable_typeFrom *MD_Georeferenceable_Type) (md_georeferenceable_typeTo *MD_Georeferenceable_Type) {
	var alreadyCopied bool
	md_georeferenceable_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_georeferenceable_typeFrom)
	if alreadyCopied {
		return
	}
	md_georeferenceable_typeFrom.GongCopyBasicFields(md_georeferenceable_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_georeferenceable_typeFrom.CellGeometry != nil {
		md_georeferenceable_typeTo.CellGeometry = GongCopyBranchMD_CellGeometryCode_PropertyType(mapOrigCopy, md_georeferenceable_typeFrom.CellGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_georeferenceable_typeFrom.AxisDimensionProperties {
		md_georeferenceable_typeTo.AxisDimensionProperties = append(md_georeferenceable_typeTo.AxisDimensionProperties, GongCopyBranchMD_Dimension_PropertyType(mapOrigCopy, _md_dimension_propertytype))
	}

	return
}

func GongCopyBranchMD_GridSpatialRepresentation_PropertyType(mapOrigCopy map[any]any, md_gridspatialrepresentation_propertytypeFrom *MD_GridSpatialRepresentation_PropertyType) (md_gridspatialrepresentation_propertytypeTo *MD_GridSpatialRepresentation_PropertyType) {
	var alreadyCopied bool
	md_gridspatialrepresentation_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_gridspatialrepresentation_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_gridspatialrepresentation_propertytypeFrom.GongCopyBasicFields(md_gridspatialrepresentation_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_gridspatialrepresentation_propertytypeFrom.MD_GridSpatialRepresentation != nil {
		md_gridspatialrepresentation_propertytypeTo.MD_GridSpatialRepresentation = GongCopyBranchMD_GridSpatialRepresentation_Type(mapOrigCopy, md_gridspatialrepresentation_propertytypeFrom.MD_GridSpatialRepresentation)
	}
	if md_gridspatialrepresentation_propertytypeFrom.MD_Georeferenceable != nil {
		md_gridspatialrepresentation_propertytypeTo.MD_Georeferenceable = GongCopyBranchMD_Georeferenceable_Type(mapOrigCopy, md_gridspatialrepresentation_propertytypeFrom.MD_Georeferenceable)
	}
	if md_gridspatialrepresentation_propertytypeFrom.MD_Georectified != nil {
		md_gridspatialrepresentation_propertytypeTo.MD_Georectified = GongCopyBranchMD_Georectified_Type(mapOrigCopy, md_gridspatialrepresentation_propertytypeFrom.MD_Georectified)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_GridSpatialRepresentation_Type(mapOrigCopy map[any]any, md_gridspatialrepresentation_typeFrom *MD_GridSpatialRepresentation_Type) (md_gridspatialrepresentation_typeTo *MD_GridSpatialRepresentation_Type) {
	var alreadyCopied bool
	md_gridspatialrepresentation_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_gridspatialrepresentation_typeFrom)
	if alreadyCopied {
		return
	}
	md_gridspatialrepresentation_typeFrom.GongCopyBasicFields(md_gridspatialrepresentation_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_gridspatialrepresentation_typeFrom.CellGeometry != nil {
		md_gridspatialrepresentation_typeTo.CellGeometry = GongCopyBranchMD_CellGeometryCode_PropertyType(mapOrigCopy, md_gridspatialrepresentation_typeFrom.CellGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_gridspatialrepresentation_typeFrom.AxisDimensionProperties {
		md_gridspatialrepresentation_typeTo.AxisDimensionProperties = append(md_gridspatialrepresentation_typeTo.AxisDimensionProperties, GongCopyBranchMD_Dimension_PropertyType(mapOrigCopy, _md_dimension_propertytype))
	}

	return
}

func GongCopyBranchMD_PixelOrientationCode_PropertyType(mapOrigCopy map[any]any, md_pixelorientationcode_propertytypeFrom *MD_PixelOrientationCode_PropertyType) (md_pixelorientationcode_propertytypeTo *MD_PixelOrientationCode_PropertyType) {
	var alreadyCopied bool
	md_pixelorientationcode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_pixelorientationcode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_pixelorientationcode_propertytypeFrom.GongCopyBasicFields(md_pixelorientationcode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_SpatialRepresentation_PropertyType(mapOrigCopy map[any]any, md_spatialrepresentation_propertytypeFrom *MD_SpatialRepresentation_PropertyType) (md_spatialrepresentation_propertytypeTo *MD_SpatialRepresentation_PropertyType) {
	var alreadyCopied bool
	md_spatialrepresentation_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_spatialrepresentation_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_spatialrepresentation_propertytypeFrom.GongCopyBasicFields(md_spatialrepresentation_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_spatialrepresentation_propertytypeFrom.MD_GridSpatialRepresentation != nil {
		md_spatialrepresentation_propertytypeTo.MD_GridSpatialRepresentation = GongCopyBranchMD_GridSpatialRepresentation_Type(mapOrigCopy, md_spatialrepresentation_propertytypeFrom.MD_GridSpatialRepresentation)
	}
	if md_spatialrepresentation_propertytypeFrom.MD_Georeferenceable != nil {
		md_spatialrepresentation_propertytypeTo.MD_Georeferenceable = GongCopyBranchMD_Georeferenceable_Type(mapOrigCopy, md_spatialrepresentation_propertytypeFrom.MD_Georeferenceable)
	}
	if md_spatialrepresentation_propertytypeFrom.MD_Georectified != nil {
		md_spatialrepresentation_propertytypeTo.MD_Georectified = GongCopyBranchMD_Georectified_Type(mapOrigCopy, md_spatialrepresentation_propertytypeFrom.MD_Georectified)
	}
	if md_spatialrepresentation_propertytypeFrom.MD_VectorSpatialRepresentation != nil {
		md_spatialrepresentation_propertytypeTo.MD_VectorSpatialRepresentation = GongCopyBranchMD_VectorSpatialRepresentation_Type(mapOrigCopy, md_spatialrepresentation_propertytypeFrom.MD_VectorSpatialRepresentation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_TopologyLevelCode_PropertyType(mapOrigCopy map[any]any, md_topologylevelcode_propertytypeFrom *MD_TopologyLevelCode_PropertyType) (md_topologylevelcode_propertytypeTo *MD_TopologyLevelCode_PropertyType) {
	var alreadyCopied bool
	md_topologylevelcode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_topologylevelcode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_topologylevelcode_propertytypeFrom.GongCopyBasicFields(md_topologylevelcode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_VectorSpatialRepresentation_PropertyType(mapOrigCopy map[any]any, md_vectorspatialrepresentation_propertytypeFrom *MD_VectorSpatialRepresentation_PropertyType) (md_vectorspatialrepresentation_propertytypeTo *MD_VectorSpatialRepresentation_PropertyType) {
	var alreadyCopied bool
	md_vectorspatialrepresentation_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_vectorspatialrepresentation_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_vectorspatialrepresentation_propertytypeFrom.GongCopyBasicFields(md_vectorspatialrepresentation_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_vectorspatialrepresentation_propertytypeFrom.MD_VectorSpatialRepresentation != nil {
		md_vectorspatialrepresentation_propertytypeTo.MD_VectorSpatialRepresentation = GongCopyBranchMD_VectorSpatialRepresentation_Type(mapOrigCopy, md_vectorspatialrepresentation_propertytypeFrom.MD_VectorSpatialRepresentation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_VectorSpatialRepresentation_Type(mapOrigCopy map[any]any, md_vectorspatialrepresentation_typeFrom *MD_VectorSpatialRepresentation_Type) (md_vectorspatialrepresentation_typeTo *MD_VectorSpatialRepresentation_Type) {
	var alreadyCopied bool
	md_vectorspatialrepresentation_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_vectorspatialrepresentation_typeFrom)
	if alreadyCopied {
		return
	}
	md_vectorspatialrepresentation_typeFrom.GongCopyBasicFields(md_vectorspatialrepresentation_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_vectorspatialrepresentation_typeFrom.TopologyLevel != nil {
		md_vectorspatialrepresentation_typeTo.TopologyLevel = GongCopyBranchMD_TopologyLevelCode_PropertyType(mapOrigCopy, md_vectorspatialrepresentation_typeFrom.TopologyLevel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_geometricobjects_propertytype := range md_vectorspatialrepresentation_typeFrom.GeometricObjects {
		md_vectorspatialrepresentation_typeTo.GeometricObjects = append(md_vectorspatialrepresentation_typeTo.GeometricObjects, GongCopyBranchMD_GeometricObjects_PropertyType(mapOrigCopy, _md_geometricobjects_propertytype))
	}

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
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractmd_spatialrepresentation_type) {
		return
	}

	abstractmd_spatialrepresentation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_cellgeometrycode_propertytype) {
		return
	}

	md_cellgeometrycode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_dimensionnametypecode_propertytype) {
		return
	}

	md_dimensionnametypecode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_dimension_propertytype) {
		return
	}

	md_dimension_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_dimension_propertytype.MD_Dimension != nil {
		stage.UnstageBranch(md_dimension_propertytype.MD_Dimension)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_dimension_type *MD_Dimension_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_dimension_type) {
		return
	}

	md_dimension_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_dimension_type.DimensionName != nil {
		stage.UnstageBranch(md_dimension_type.DimensionName)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_geometricobjecttypecode_propertytype) {
		return
	}

	md_geometricobjecttypecode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_geometricobjects_propertytype) {
		return
	}

	md_geometricobjects_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_geometricobjects_propertytype.MD_GeometricObjects != nil {
		stage.UnstageBranch(md_geometricobjects_propertytype.MD_GeometricObjects)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_geometricobjects_type) {
		return
	}

	md_geometricobjects_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_geometricobjects_type.GeometricObjectType != nil {
		stage.UnstageBranch(md_geometricobjects_type.GeometricObjectType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_georectified_propertytype) {
		return
	}

	md_georectified_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_georectified_propertytype.MD_Georectified != nil {
		stage.UnstageBranch(md_georectified_propertytype.MD_Georectified)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_georectified_type *MD_Georectified_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_georectified_type) {
		return
	}

	md_georectified_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_georectified_type.CellGeometry != nil {
		stage.UnstageBranch(md_georectified_type.CellGeometry)
	}
	if md_georectified_type.PointInPixel != nil {
		stage.UnstageBranch(md_georectified_type.PointInPixel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_georectified_type.AxisDimensionProperties {
		stage.UnstageBranch(_md_dimension_propertytype)
	}

}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_georeferenceable_propertytype) {
		return
	}

	md_georeferenceable_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_georeferenceable_propertytype.MD_Georeferenceable != nil {
		stage.UnstageBranch(md_georeferenceable_propertytype.MD_Georeferenceable)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_georeferenceable_type) {
		return
	}

	md_georeferenceable_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_georeferenceable_type.CellGeometry != nil {
		stage.UnstageBranch(md_georeferenceable_type.CellGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_georeferenceable_type.AxisDimensionProperties {
		stage.UnstageBranch(_md_dimension_propertytype)
	}

}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_gridspatialrepresentation_propertytype) {
		return
	}

	md_gridspatialrepresentation_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
		stage.UnstageBranch(md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation)
	}
	if md_gridspatialrepresentation_propertytype.MD_Georeferenceable != nil {
		stage.UnstageBranch(md_gridspatialrepresentation_propertytype.MD_Georeferenceable)
	}
	if md_gridspatialrepresentation_propertytype.MD_Georectified != nil {
		stage.UnstageBranch(md_gridspatialrepresentation_propertytype.MD_Georectified)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_gridspatialrepresentation_type) {
		return
	}

	md_gridspatialrepresentation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_gridspatialrepresentation_type.CellGeometry != nil {
		stage.UnstageBranch(md_gridspatialrepresentation_type.CellGeometry)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_dimension_propertytype := range md_gridspatialrepresentation_type.AxisDimensionProperties {
		stage.UnstageBranch(_md_dimension_propertytype)
	}

}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_pixelorientationcode_propertytype) {
		return
	}

	md_pixelorientationcode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_spatialrepresentation_propertytype) {
		return
	}

	md_spatialrepresentation_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation != nil {
		stage.UnstageBranch(md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation)
	}
	if md_spatialrepresentation_propertytype.MD_Georeferenceable != nil {
		stage.UnstageBranch(md_spatialrepresentation_propertytype.MD_Georeferenceable)
	}
	if md_spatialrepresentation_propertytype.MD_Georectified != nil {
		stage.UnstageBranch(md_spatialrepresentation_propertytype.MD_Georectified)
	}
	if md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
		stage.UnstageBranch(md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_topologylevelcode_propertytype) {
		return
	}

	md_topologylevelcode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_vectorspatialrepresentation_propertytype) {
		return
	}

	md_vectorspatialrepresentation_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation != nil {
		stage.UnstageBranch(md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_vectorspatialrepresentation_type) {
		return
	}

	md_vectorspatialrepresentation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_vectorspatialrepresentation_type.TopologyLevel != nil {
		stage.UnstageBranch(md_vectorspatialrepresentation_type.TopologyLevel)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_geometricobjects_propertytype := range md_vectorspatialrepresentation_type.GeometricObjects {
		stage.UnstageBranch(_md_geometricobjects_propertytype)
	}

}

// insertion point for pointer reconstruction from references
func (reference *AbstractMD_SpatialRepresentation_Type) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractMD_SpatialRepresentation_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_CellGeometryCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_CellGeometryCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_DimensionNameTypeCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_DimensionNameTypeCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_Dimension_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Dimension_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_Dimension, stage.MD_Dimension_Types_reference, instance.MD_Dimension)
	// insertion point for slice of pointers field
}

func (reference *MD_Dimension_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Dimension_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DimensionName, stage.MD_DimensionNameTypeCode_PropertyTypes_reference, instance.DimensionName)
	// insertion point for slice of pointers field
}

func (reference *MD_GeometricObjectTypeCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_GeometricObjectTypeCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_GeometricObjects_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_GeometricObjects_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_GeometricObjects, stage.MD_GeometricObjects_Types_reference, instance.MD_GeometricObjects)
	// insertion point for slice of pointers field
}

func (reference *MD_GeometricObjects_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_GeometricObjects_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.GeometricObjectType, stage.MD_GeometricObjectTypeCode_PropertyTypes_reference, instance.GeometricObjectType)
	// insertion point for slice of pointers field
}

func (reference *MD_Georectified_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Georectified_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_Georectified, stage.MD_Georectified_Types_reference, instance.MD_Georectified)
	// insertion point for slice of pointers field
}

func (reference *MD_Georectified_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Georectified_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.CellGeometry, stage.MD_CellGeometryCode_PropertyTypes_reference, instance.CellGeometry)
	__gong__reconstructPointer(&reference.PointInPixel, stage.MD_PixelOrientationCode_PropertyTypes_reference, instance.PointInPixel)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.AxisDimensionProperties, stage.MD_Dimension_PropertyTypes_reference, instance.AxisDimensionProperties)
}

func (reference *MD_Georeferenceable_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Georeferenceable_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_Georeferenceable, stage.MD_Georeferenceable_Types_reference, instance.MD_Georeferenceable)
	// insertion point for slice of pointers field
}

func (reference *MD_Georeferenceable_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Georeferenceable_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.CellGeometry, stage.MD_CellGeometryCode_PropertyTypes_reference, instance.CellGeometry)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.AxisDimensionProperties, stage.MD_Dimension_PropertyTypes_reference, instance.AxisDimensionProperties)
}

func (reference *MD_GridSpatialRepresentation_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_GridSpatialRepresentation_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_GridSpatialRepresentation, stage.MD_GridSpatialRepresentation_Types_reference, instance.MD_GridSpatialRepresentation)
	__gong__reconstructPointer(&reference.MD_Georeferenceable, stage.MD_Georeferenceable_Types_reference, instance.MD_Georeferenceable)
	__gong__reconstructPointer(&reference.MD_Georectified, stage.MD_Georectified_Types_reference, instance.MD_Georectified)
	// insertion point for slice of pointers field
}

func (reference *MD_GridSpatialRepresentation_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_GridSpatialRepresentation_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.CellGeometry, stage.MD_CellGeometryCode_PropertyTypes_reference, instance.CellGeometry)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.AxisDimensionProperties, stage.MD_Dimension_PropertyTypes_reference, instance.AxisDimensionProperties)
}

func (reference *MD_PixelOrientationCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_PixelOrientationCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_SpatialRepresentation_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_SpatialRepresentation_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_GridSpatialRepresentation, stage.MD_GridSpatialRepresentation_Types_reference, instance.MD_GridSpatialRepresentation)
	__gong__reconstructPointer(&reference.MD_Georeferenceable, stage.MD_Georeferenceable_Types_reference, instance.MD_Georeferenceable)
	__gong__reconstructPointer(&reference.MD_Georectified, stage.MD_Georectified_Types_reference, instance.MD_Georectified)
	__gong__reconstructPointer(&reference.MD_VectorSpatialRepresentation, stage.MD_VectorSpatialRepresentation_Types_reference, instance.MD_VectorSpatialRepresentation)
	// insertion point for slice of pointers field
}

func (reference *MD_TopologyLevelCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_TopologyLevelCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_VectorSpatialRepresentation_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_VectorSpatialRepresentation_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_VectorSpatialRepresentation, stage.MD_VectorSpatialRepresentation_Types_reference, instance.MD_VectorSpatialRepresentation)
	// insertion point for slice of pointers field
}

func (reference *MD_VectorSpatialRepresentation_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_VectorSpatialRepresentation_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TopologyLevel, stage.MD_TopologyLevelCode_PropertyTypes_reference, instance.TopologyLevel)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GeometricObjects, stage.MD_GeometricObjects_PropertyTypes_reference, instance.GeometricObjects)
}

// insertion point for pointer reconstruction from instances
func (reference *AbstractMD_SpatialRepresentation_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_CellGeometryCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_DimensionNameTypeCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_Dimension_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_Dimension, stage.MD_Dimension_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_Dimension_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DimensionName, stage.MD_DimensionNameTypeCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_GeometricObjectTypeCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_GeometricObjects_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_GeometricObjects, stage.MD_GeometricObjects_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_GeometricObjects_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.GeometricObjectType, stage.MD_GeometricObjectTypeCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_Georectified_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_Georectified, stage.MD_Georectified_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_Georectified_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.CellGeometry, stage.MD_CellGeometryCode_PropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.PointInPixel, stage.MD_PixelOrientationCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.AxisDimensionProperties, stage.MD_Dimension_PropertyTypes_instance)
}

func (reference *MD_Georeferenceable_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_Georeferenceable, stage.MD_Georeferenceable_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_Georeferenceable_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.CellGeometry, stage.MD_CellGeometryCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.AxisDimensionProperties, stage.MD_Dimension_PropertyTypes_instance)
}

func (reference *MD_GridSpatialRepresentation_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_GridSpatialRepresentation, stage.MD_GridSpatialRepresentation_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_Georeferenceable, stage.MD_Georeferenceable_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_Georectified, stage.MD_Georectified_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_GridSpatialRepresentation_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.CellGeometry, stage.MD_CellGeometryCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.AxisDimensionProperties, stage.MD_Dimension_PropertyTypes_instance)
}

func (reference *MD_PixelOrientationCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_SpatialRepresentation_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_GridSpatialRepresentation, stage.MD_GridSpatialRepresentation_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_Georeferenceable, stage.MD_Georeferenceable_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_Georectified, stage.MD_Georectified_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_VectorSpatialRepresentation, stage.MD_VectorSpatialRepresentation_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_TopologyLevelCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_VectorSpatialRepresentation_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_VectorSpatialRepresentation, stage.MD_VectorSpatialRepresentation_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_VectorSpatialRepresentation_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TopologyLevel, stage.MD_TopologyLevelCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GeometricObjects, stage.MD_GeometricObjects_PropertyTypes_instance)
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongDiff(stage *Stage, abstractmd_spatialrepresentation_typeOther *AbstractMD_SpatialRepresentation_Type) (diffs []string) {
	// insertion point for field diffs
	if abstractmd_spatialrepresentation_type.Name != abstractmd_spatialrepresentation_typeOther.Name {
		diffs = append(diffs, abstractmd_spatialrepresentation_type.GongMarshallField(stage, "Name"))
	}
	if abstractmd_spatialrepresentation_type.Id != abstractmd_spatialrepresentation_typeOther.Id {
		diffs = append(diffs, abstractmd_spatialrepresentation_type.GongMarshallField(stage, "Id"))
	}
	if abstractmd_spatialrepresentation_type.Uuid != abstractmd_spatialrepresentation_typeOther.Uuid {
		diffs = append(diffs, abstractmd_spatialrepresentation_type.GongMarshallField(stage, "Uuid"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongDiff(stage *Stage, md_cellgeometrycode_propertytypeOther *MD_CellGeometryCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_cellgeometrycode_propertytype.Name != md_cellgeometrycode_propertytypeOther.Name {
		diffs = append(diffs, md_cellgeometrycode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_cellgeometrycode_propertytype.NilReason != md_cellgeometrycode_propertytypeOther.NilReason {
		diffs = append(diffs, md_cellgeometrycode_propertytype.GongMarshallField(stage, "NilReason"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongDiff(stage *Stage, md_dimensionnametypecode_propertytypeOther *MD_DimensionNameTypeCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_dimensionnametypecode_propertytype.Name != md_dimensionnametypecode_propertytypeOther.Name {
		diffs = append(diffs, md_dimensionnametypecode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_dimensionnametypecode_propertytype.NilReason != md_dimensionnametypecode_propertytypeOther.NilReason {
		diffs = append(diffs, md_dimensionnametypecode_propertytype.GongMarshallField(stage, "NilReason"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_dimension_propertytype *MD_Dimension_PropertyType) GongDiff(stage *Stage, md_dimension_propertytypeOther *MD_Dimension_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_dimension_propertytype.Name != md_dimension_propertytypeOther.Name {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_dimension_propertytype.NilReason != md_dimension_propertytypeOther.NilReason {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_dimension_propertytype.Uuidref != md_dimension_propertytypeOther.Uuidref {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_dimension_propertytype.Type_XSD != md_dimension_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_dimension_propertytype.Href != md_dimension_propertytypeOther.Href {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_dimension_propertytype.Role != md_dimension_propertytypeOther.Role {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_dimension_propertytype.Arcrole != md_dimension_propertytypeOther.Arcrole {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_dimension_propertytype.Title != md_dimension_propertytypeOther.Title {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_dimension_propertytype.Show != md_dimension_propertytypeOther.Show {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_dimension_propertytype.Actuate != md_dimension_propertytypeOther.Actuate {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_dimension_propertytype.MD_Dimension != md_dimension_propertytypeOther.MD_Dimension {
		diffs = append(diffs, md_dimension_propertytype.GongMarshallField(stage, "MD_Dimension"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_dimension_type *MD_Dimension_Type) GongDiff(stage *Stage, md_dimension_typeOther *MD_Dimension_Type) (diffs []string) {
	// insertion point for field diffs
	if md_dimension_type.Name != md_dimension_typeOther.Name {
		diffs = append(diffs, md_dimension_type.GongMarshallField(stage, "Name"))
	}
	if md_dimension_type.Id != md_dimension_typeOther.Id {
		diffs = append(diffs, md_dimension_type.GongMarshallField(stage, "Id"))
	}
	if md_dimension_type.Uuid != md_dimension_typeOther.Uuid {
		diffs = append(diffs, md_dimension_type.GongMarshallField(stage, "Uuid"))
	}
	if md_dimension_type.DimensionName != md_dimension_typeOther.DimensionName {
		diffs = append(diffs, md_dimension_type.GongMarshallField(stage, "DimensionName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongDiff(stage *Stage, md_geometricobjecttypecode_propertytypeOther *MD_GeometricObjectTypeCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_geometricobjecttypecode_propertytype.Name != md_geometricobjecttypecode_propertytypeOther.Name {
		diffs = append(diffs, md_geometricobjecttypecode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_geometricobjecttypecode_propertytype.NilReason != md_geometricobjecttypecode_propertytypeOther.NilReason {
		diffs = append(diffs, md_geometricobjecttypecode_propertytype.GongMarshallField(stage, "NilReason"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongDiff(stage *Stage, md_geometricobjects_propertytypeOther *MD_GeometricObjects_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_geometricobjects_propertytype.Name != md_geometricobjects_propertytypeOther.Name {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_geometricobjects_propertytype.NilReason != md_geometricobjects_propertytypeOther.NilReason {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_geometricobjects_propertytype.Uuidref != md_geometricobjects_propertytypeOther.Uuidref {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_geometricobjects_propertytype.Type_XSD != md_geometricobjects_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_geometricobjects_propertytype.Href != md_geometricobjects_propertytypeOther.Href {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_geometricobjects_propertytype.Role != md_geometricobjects_propertytypeOther.Role {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_geometricobjects_propertytype.Arcrole != md_geometricobjects_propertytypeOther.Arcrole {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_geometricobjects_propertytype.Title != md_geometricobjects_propertytypeOther.Title {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_geometricobjects_propertytype.Show != md_geometricobjects_propertytypeOther.Show {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_geometricobjects_propertytype.Actuate != md_geometricobjects_propertytypeOther.Actuate {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_geometricobjects_propertytype.MD_GeometricObjects != md_geometricobjects_propertytypeOther.MD_GeometricObjects {
		diffs = append(diffs, md_geometricobjects_propertytype.GongMarshallField(stage, "MD_GeometricObjects"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_geometricobjects_type *MD_GeometricObjects_Type) GongDiff(stage *Stage, md_geometricobjects_typeOther *MD_GeometricObjects_Type) (diffs []string) {
	// insertion point for field diffs
	if md_geometricobjects_type.Name != md_geometricobjects_typeOther.Name {
		diffs = append(diffs, md_geometricobjects_type.GongMarshallField(stage, "Name"))
	}
	if md_geometricobjects_type.Id != md_geometricobjects_typeOther.Id {
		diffs = append(diffs, md_geometricobjects_type.GongMarshallField(stage, "Id"))
	}
	if md_geometricobjects_type.Uuid != md_geometricobjects_typeOther.Uuid {
		diffs = append(diffs, md_geometricobjects_type.GongMarshallField(stage, "Uuid"))
	}
	if md_geometricobjects_type.GeometricObjectType != md_geometricobjects_typeOther.GeometricObjectType {
		diffs = append(diffs, md_geometricobjects_type.GongMarshallField(stage, "GeometricObjectType"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_georectified_propertytype *MD_Georectified_PropertyType) GongDiff(stage *Stage, md_georectified_propertytypeOther *MD_Georectified_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_georectified_propertytype.Name != md_georectified_propertytypeOther.Name {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_georectified_propertytype.NilReason != md_georectified_propertytypeOther.NilReason {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_georectified_propertytype.Uuidref != md_georectified_propertytypeOther.Uuidref {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_georectified_propertytype.Type_XSD != md_georectified_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_georectified_propertytype.Href != md_georectified_propertytypeOther.Href {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_georectified_propertytype.Role != md_georectified_propertytypeOther.Role {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_georectified_propertytype.Arcrole != md_georectified_propertytypeOther.Arcrole {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_georectified_propertytype.Title != md_georectified_propertytypeOther.Title {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_georectified_propertytype.Show != md_georectified_propertytypeOther.Show {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_georectified_propertytype.Actuate != md_georectified_propertytypeOther.Actuate {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_georectified_propertytype.MD_Georectified != md_georectified_propertytypeOther.MD_Georectified {
		diffs = append(diffs, md_georectified_propertytype.GongMarshallField(stage, "MD_Georectified"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_georectified_type *MD_Georectified_Type) GongDiff(stage *Stage, md_georectified_typeOther *MD_Georectified_Type) (diffs []string) {
	// insertion point for field diffs
	if md_georectified_type.Name != md_georectified_typeOther.Name {
		diffs = append(diffs, md_georectified_type.GongMarshallField(stage, "Name"))
	}
	if md_georectified_type.Id != md_georectified_typeOther.Id {
		diffs = append(diffs, md_georectified_type.GongMarshallField(stage, "Id"))
	}
	if md_georectified_type.Uuid != md_georectified_typeOther.Uuid {
		diffs = append(diffs, md_georectified_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, md_georectified_type, "AxisDimensionProperties", md_georectified_typeOther.AxisDimensionProperties, md_georectified_type.AxisDimensionProperties); ops != "" {
		diffs = append(diffs, ops)
	}
	if md_georectified_type.CellGeometry != md_georectified_typeOther.CellGeometry {
		diffs = append(diffs, md_georectified_type.GongMarshallField(stage, "CellGeometry"))
	}
	if md_georectified_type.PointInPixel != md_georectified_typeOther.PointInPixel {
		diffs = append(diffs, md_georectified_type.GongMarshallField(stage, "PointInPixel"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongDiff(stage *Stage, md_georeferenceable_propertytypeOther *MD_Georeferenceable_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_georeferenceable_propertytype.Name != md_georeferenceable_propertytypeOther.Name {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_georeferenceable_propertytype.NilReason != md_georeferenceable_propertytypeOther.NilReason {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_georeferenceable_propertytype.Uuidref != md_georeferenceable_propertytypeOther.Uuidref {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_georeferenceable_propertytype.Type_XSD != md_georeferenceable_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_georeferenceable_propertytype.Href != md_georeferenceable_propertytypeOther.Href {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_georeferenceable_propertytype.Role != md_georeferenceable_propertytypeOther.Role {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_georeferenceable_propertytype.Arcrole != md_georeferenceable_propertytypeOther.Arcrole {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_georeferenceable_propertytype.Title != md_georeferenceable_propertytypeOther.Title {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_georeferenceable_propertytype.Show != md_georeferenceable_propertytypeOther.Show {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_georeferenceable_propertytype.Actuate != md_georeferenceable_propertytypeOther.Actuate {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_georeferenceable_propertytype.MD_Georeferenceable != md_georeferenceable_propertytypeOther.MD_Georeferenceable {
		diffs = append(diffs, md_georeferenceable_propertytype.GongMarshallField(stage, "MD_Georeferenceable"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_georeferenceable_type *MD_Georeferenceable_Type) GongDiff(stage *Stage, md_georeferenceable_typeOther *MD_Georeferenceable_Type) (diffs []string) {
	// insertion point for field diffs
	if md_georeferenceable_type.Name != md_georeferenceable_typeOther.Name {
		diffs = append(diffs, md_georeferenceable_type.GongMarshallField(stage, "Name"))
	}
	if md_georeferenceable_type.Id != md_georeferenceable_typeOther.Id {
		diffs = append(diffs, md_georeferenceable_type.GongMarshallField(stage, "Id"))
	}
	if md_georeferenceable_type.Uuid != md_georeferenceable_typeOther.Uuid {
		diffs = append(diffs, md_georeferenceable_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, md_georeferenceable_type, "AxisDimensionProperties", md_georeferenceable_typeOther.AxisDimensionProperties, md_georeferenceable_type.AxisDimensionProperties); ops != "" {
		diffs = append(diffs, ops)
	}
	if md_georeferenceable_type.CellGeometry != md_georeferenceable_typeOther.CellGeometry {
		diffs = append(diffs, md_georeferenceable_type.GongMarshallField(stage, "CellGeometry"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongDiff(stage *Stage, md_gridspatialrepresentation_propertytypeOther *MD_GridSpatialRepresentation_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_gridspatialrepresentation_propertytype.Name != md_gridspatialrepresentation_propertytypeOther.Name {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_gridspatialrepresentation_propertytype.NilReason != md_gridspatialrepresentation_propertytypeOther.NilReason {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_gridspatialrepresentation_propertytype.Uuidref != md_gridspatialrepresentation_propertytypeOther.Uuidref {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_gridspatialrepresentation_propertytype.Type_XSD != md_gridspatialrepresentation_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_gridspatialrepresentation_propertytype.Href != md_gridspatialrepresentation_propertytypeOther.Href {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_gridspatialrepresentation_propertytype.Role != md_gridspatialrepresentation_propertytypeOther.Role {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_gridspatialrepresentation_propertytype.Arcrole != md_gridspatialrepresentation_propertytypeOther.Arcrole {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_gridspatialrepresentation_propertytype.Title != md_gridspatialrepresentation_propertytypeOther.Title {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_gridspatialrepresentation_propertytype.Show != md_gridspatialrepresentation_propertytypeOther.Show {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_gridspatialrepresentation_propertytype.Actuate != md_gridspatialrepresentation_propertytypeOther.Actuate {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_gridspatialrepresentation_propertytype.MD_GridSpatialRepresentation != md_gridspatialrepresentation_propertytypeOther.MD_GridSpatialRepresentation {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "MD_GridSpatialRepresentation"))
	}
	if md_gridspatialrepresentation_propertytype.MD_Georeferenceable != md_gridspatialrepresentation_propertytypeOther.MD_Georeferenceable {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "MD_Georeferenceable"))
	}
	if md_gridspatialrepresentation_propertytype.MD_Georectified != md_gridspatialrepresentation_propertytypeOther.MD_Georectified {
		diffs = append(diffs, md_gridspatialrepresentation_propertytype.GongMarshallField(stage, "MD_Georectified"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongDiff(stage *Stage, md_gridspatialrepresentation_typeOther *MD_GridSpatialRepresentation_Type) (diffs []string) {
	// insertion point for field diffs
	if md_gridspatialrepresentation_type.Name != md_gridspatialrepresentation_typeOther.Name {
		diffs = append(diffs, md_gridspatialrepresentation_type.GongMarshallField(stage, "Name"))
	}
	if md_gridspatialrepresentation_type.Id != md_gridspatialrepresentation_typeOther.Id {
		diffs = append(diffs, md_gridspatialrepresentation_type.GongMarshallField(stage, "Id"))
	}
	if md_gridspatialrepresentation_type.Uuid != md_gridspatialrepresentation_typeOther.Uuid {
		diffs = append(diffs, md_gridspatialrepresentation_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, md_gridspatialrepresentation_type, "AxisDimensionProperties", md_gridspatialrepresentation_typeOther.AxisDimensionProperties, md_gridspatialrepresentation_type.AxisDimensionProperties); ops != "" {
		diffs = append(diffs, ops)
	}
	if md_gridspatialrepresentation_type.CellGeometry != md_gridspatialrepresentation_typeOther.CellGeometry {
		diffs = append(diffs, md_gridspatialrepresentation_type.GongMarshallField(stage, "CellGeometry"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongDiff(stage *Stage, md_pixelorientationcode_propertytypeOther *MD_PixelOrientationCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_pixelorientationcode_propertytype.Name != md_pixelorientationcode_propertytypeOther.Name {
		diffs = append(diffs, md_pixelorientationcode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_pixelorientationcode_propertytype.NilReason != md_pixelorientationcode_propertytypeOther.NilReason {
		diffs = append(diffs, md_pixelorientationcode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_pixelorientationcode_propertytype.MD_PixelOrientationCode != md_pixelorientationcode_propertytypeOther.MD_PixelOrientationCode {
		diffs = append(diffs, md_pixelorientationcode_propertytype.GongMarshallField(stage, "MD_PixelOrientationCode"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongDiff(stage *Stage, md_spatialrepresentation_propertytypeOther *MD_SpatialRepresentation_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_spatialrepresentation_propertytype.Name != md_spatialrepresentation_propertytypeOther.Name {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_spatialrepresentation_propertytype.NilReason != md_spatialrepresentation_propertytypeOther.NilReason {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_spatialrepresentation_propertytype.Uuidref != md_spatialrepresentation_propertytypeOther.Uuidref {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_spatialrepresentation_propertytype.Type_XSD != md_spatialrepresentation_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_spatialrepresentation_propertytype.Href != md_spatialrepresentation_propertytypeOther.Href {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_spatialrepresentation_propertytype.Role != md_spatialrepresentation_propertytypeOther.Role {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_spatialrepresentation_propertytype.Arcrole != md_spatialrepresentation_propertytypeOther.Arcrole {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_spatialrepresentation_propertytype.Title != md_spatialrepresentation_propertytypeOther.Title {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_spatialrepresentation_propertytype.Show != md_spatialrepresentation_propertytypeOther.Show {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_spatialrepresentation_propertytype.Actuate != md_spatialrepresentation_propertytypeOther.Actuate {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_spatialrepresentation_propertytype.MD_GridSpatialRepresentation != md_spatialrepresentation_propertytypeOther.MD_GridSpatialRepresentation {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "MD_GridSpatialRepresentation"))
	}
	if md_spatialrepresentation_propertytype.MD_Georeferenceable != md_spatialrepresentation_propertytypeOther.MD_Georeferenceable {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "MD_Georeferenceable"))
	}
	if md_spatialrepresentation_propertytype.MD_Georectified != md_spatialrepresentation_propertytypeOther.MD_Georectified {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "MD_Georectified"))
	}
	if md_spatialrepresentation_propertytype.MD_VectorSpatialRepresentation != md_spatialrepresentation_propertytypeOther.MD_VectorSpatialRepresentation {
		diffs = append(diffs, md_spatialrepresentation_propertytype.GongMarshallField(stage, "MD_VectorSpatialRepresentation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongDiff(stage *Stage, md_topologylevelcode_propertytypeOther *MD_TopologyLevelCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_topologylevelcode_propertytype.Name != md_topologylevelcode_propertytypeOther.Name {
		diffs = append(diffs, md_topologylevelcode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_topologylevelcode_propertytype.NilReason != md_topologylevelcode_propertytypeOther.NilReason {
		diffs = append(diffs, md_topologylevelcode_propertytype.GongMarshallField(stage, "NilReason"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongDiff(stage *Stage, md_vectorspatialrepresentation_propertytypeOther *MD_VectorSpatialRepresentation_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_vectorspatialrepresentation_propertytype.Name != md_vectorspatialrepresentation_propertytypeOther.Name {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_vectorspatialrepresentation_propertytype.NilReason != md_vectorspatialrepresentation_propertytypeOther.NilReason {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_vectorspatialrepresentation_propertytype.Uuidref != md_vectorspatialrepresentation_propertytypeOther.Uuidref {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_vectorspatialrepresentation_propertytype.Type_XSD != md_vectorspatialrepresentation_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_vectorspatialrepresentation_propertytype.Href != md_vectorspatialrepresentation_propertytypeOther.Href {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_vectorspatialrepresentation_propertytype.Role != md_vectorspatialrepresentation_propertytypeOther.Role {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_vectorspatialrepresentation_propertytype.Arcrole != md_vectorspatialrepresentation_propertytypeOther.Arcrole {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_vectorspatialrepresentation_propertytype.Title != md_vectorspatialrepresentation_propertytypeOther.Title {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_vectorspatialrepresentation_propertytype.Show != md_vectorspatialrepresentation_propertytypeOther.Show {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_vectorspatialrepresentation_propertytype.Actuate != md_vectorspatialrepresentation_propertytypeOther.Actuate {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_vectorspatialrepresentation_propertytype.MD_VectorSpatialRepresentation != md_vectorspatialrepresentation_propertytypeOther.MD_VectorSpatialRepresentation {
		diffs = append(diffs, md_vectorspatialrepresentation_propertytype.GongMarshallField(stage, "MD_VectorSpatialRepresentation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongDiff(stage *Stage, md_vectorspatialrepresentation_typeOther *MD_VectorSpatialRepresentation_Type) (diffs []string) {
	// insertion point for field diffs
	if md_vectorspatialrepresentation_type.Name != md_vectorspatialrepresentation_typeOther.Name {
		diffs = append(diffs, md_vectorspatialrepresentation_type.GongMarshallField(stage, "Name"))
	}
	if md_vectorspatialrepresentation_type.Id != md_vectorspatialrepresentation_typeOther.Id {
		diffs = append(diffs, md_vectorspatialrepresentation_type.GongMarshallField(stage, "Id"))
	}
	if md_vectorspatialrepresentation_type.Uuid != md_vectorspatialrepresentation_typeOther.Uuid {
		diffs = append(diffs, md_vectorspatialrepresentation_type.GongMarshallField(stage, "Uuid"))
	}
	if md_vectorspatialrepresentation_type.TopologyLevel != md_vectorspatialrepresentation_typeOther.TopologyLevel {
		diffs = append(diffs, md_vectorspatialrepresentation_type.GongMarshallField(stage, "TopologyLevel"))
	}
	if ops := __gong__diffSliceOfPointers(stage, md_vectorspatialrepresentation_type, "GeometricObjects", md_vectorspatialrepresentation_typeOther.GeometricObjects, md_vectorspatialrepresentation_type.GeometricObjects); ops != "" {
		diffs = append(diffs, ops)
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
