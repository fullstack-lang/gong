// generated code - do not edit
package spatialrepresentation

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractMD_SpatialRepresentation_TypeCreateCallback != nil {
		stage.OnAfterAbstractMD_SpatialRepresentation_TypeCreateCallback.OnAfterCreate(stage, abstractmd_spatialrepresentation_type)
	}
}

func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMD_SpatialRepresentation_TypeUpdateCallback != nil {
		var frontAbstractMD_SpatialRepresentation_Type *AbstractMD_SpatialRepresentation_Type
		if front != nil {
			frontAbstractMD_SpatialRepresentation_Type, _ = front.(*AbstractMD_SpatialRepresentation_Type)
		}
		stage.OnAfterAbstractMD_SpatialRepresentation_TypeUpdateCallback.OnAfterUpdate(stage, abstractmd_spatialrepresentation_type, frontAbstractMD_SpatialRepresentation_Type)
	}
}

func (abstractmd_spatialrepresentation_type *AbstractMD_SpatialRepresentation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMD_SpatialRepresentation_TypeDeleteCallback != nil {
		var frontAbstractMD_SpatialRepresentation_Type *AbstractMD_SpatialRepresentation_Type
		if front != nil {
			frontAbstractMD_SpatialRepresentation_Type, _ = front.(*AbstractMD_SpatialRepresentation_Type)
		}
		stage.OnAfterAbstractMD_SpatialRepresentation_TypeDeleteCallback.OnAfterDelete(stage, abstractmd_spatialrepresentation_type, frontAbstractMD_SpatialRepresentation_Type)
	}
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_CellGeometryCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_CellGeometryCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_cellgeometrycode_propertytype)
	}
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CellGeometryCode_PropertyTypeUpdateCallback != nil {
		var frontMD_CellGeometryCode_PropertyType *MD_CellGeometryCode_PropertyType
		if front != nil {
			frontMD_CellGeometryCode_PropertyType, _ = front.(*MD_CellGeometryCode_PropertyType)
		}
		stage.OnAfterMD_CellGeometryCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_cellgeometrycode_propertytype, frontMD_CellGeometryCode_PropertyType)
	}
}

func (md_cellgeometrycode_propertytype *MD_CellGeometryCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CellGeometryCode_PropertyTypeDeleteCallback != nil {
		var frontMD_CellGeometryCode_PropertyType *MD_CellGeometryCode_PropertyType
		if front != nil {
			frontMD_CellGeometryCode_PropertyType, _ = front.(*MD_CellGeometryCode_PropertyType)
		}
		stage.OnAfterMD_CellGeometryCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_cellgeometrycode_propertytype, frontMD_CellGeometryCode_PropertyType)
	}
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_DimensionNameTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_DimensionNameTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_dimensionnametypecode_propertytype)
	}
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DimensionNameTypeCode_PropertyTypeUpdateCallback != nil {
		var frontMD_DimensionNameTypeCode_PropertyType *MD_DimensionNameTypeCode_PropertyType
		if front != nil {
			frontMD_DimensionNameTypeCode_PropertyType, _ = front.(*MD_DimensionNameTypeCode_PropertyType)
		}
		stage.OnAfterMD_DimensionNameTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_dimensionnametypecode_propertytype, frontMD_DimensionNameTypeCode_PropertyType)
	}
}

func (md_dimensionnametypecode_propertytype *MD_DimensionNameTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DimensionNameTypeCode_PropertyTypeDeleteCallback != nil {
		var frontMD_DimensionNameTypeCode_PropertyType *MD_DimensionNameTypeCode_PropertyType
		if front != nil {
			frontMD_DimensionNameTypeCode_PropertyType, _ = front.(*MD_DimensionNameTypeCode_PropertyType)
		}
		stage.OnAfterMD_DimensionNameTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_dimensionnametypecode_propertytype, frontMD_DimensionNameTypeCode_PropertyType)
	}
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Dimension_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Dimension_PropertyTypeCreateCallback.OnAfterCreate(stage, md_dimension_propertytype)
	}
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Dimension_PropertyTypeUpdateCallback != nil {
		var frontMD_Dimension_PropertyType *MD_Dimension_PropertyType
		if front != nil {
			frontMD_Dimension_PropertyType, _ = front.(*MD_Dimension_PropertyType)
		}
		stage.OnAfterMD_Dimension_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_dimension_propertytype, frontMD_Dimension_PropertyType)
	}
}

func (md_dimension_propertytype *MD_Dimension_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Dimension_PropertyTypeDeleteCallback != nil {
		var frontMD_Dimension_PropertyType *MD_Dimension_PropertyType
		if front != nil {
			frontMD_Dimension_PropertyType, _ = front.(*MD_Dimension_PropertyType)
		}
		stage.OnAfterMD_Dimension_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_dimension_propertytype, frontMD_Dimension_PropertyType)
	}
}

func (md_dimension_type *MD_Dimension_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Dimension_TypeCreateCallback != nil {
		stage.OnAfterMD_Dimension_TypeCreateCallback.OnAfterCreate(stage, md_dimension_type)
	}
}

func (md_dimension_type *MD_Dimension_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Dimension_TypeUpdateCallback != nil {
		var frontMD_Dimension_Type *MD_Dimension_Type
		if front != nil {
			frontMD_Dimension_Type, _ = front.(*MD_Dimension_Type)
		}
		stage.OnAfterMD_Dimension_TypeUpdateCallback.OnAfterUpdate(stage, md_dimension_type, frontMD_Dimension_Type)
	}
}

func (md_dimension_type *MD_Dimension_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Dimension_TypeDeleteCallback != nil {
		var frontMD_Dimension_Type *MD_Dimension_Type
		if front != nil {
			frontMD_Dimension_Type, _ = front.(*MD_Dimension_Type)
		}
		stage.OnAfterMD_Dimension_TypeDeleteCallback.OnAfterDelete(stage, md_dimension_type, frontMD_Dimension_Type)
	}
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_GeometricObjectTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_GeometricObjectTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_geometricobjecttypecode_propertytype)
	}
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GeometricObjectTypeCode_PropertyTypeUpdateCallback != nil {
		var frontMD_GeometricObjectTypeCode_PropertyType *MD_GeometricObjectTypeCode_PropertyType
		if front != nil {
			frontMD_GeometricObjectTypeCode_PropertyType, _ = front.(*MD_GeometricObjectTypeCode_PropertyType)
		}
		stage.OnAfterMD_GeometricObjectTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_geometricobjecttypecode_propertytype, frontMD_GeometricObjectTypeCode_PropertyType)
	}
}

func (md_geometricobjecttypecode_propertytype *MD_GeometricObjectTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GeometricObjectTypeCode_PropertyTypeDeleteCallback != nil {
		var frontMD_GeometricObjectTypeCode_PropertyType *MD_GeometricObjectTypeCode_PropertyType
		if front != nil {
			frontMD_GeometricObjectTypeCode_PropertyType, _ = front.(*MD_GeometricObjectTypeCode_PropertyType)
		}
		stage.OnAfterMD_GeometricObjectTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_geometricobjecttypecode_propertytype, frontMD_GeometricObjectTypeCode_PropertyType)
	}
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_GeometricObjects_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_GeometricObjects_PropertyTypeCreateCallback.OnAfterCreate(stage, md_geometricobjects_propertytype)
	}
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GeometricObjects_PropertyTypeUpdateCallback != nil {
		var frontMD_GeometricObjects_PropertyType *MD_GeometricObjects_PropertyType
		if front != nil {
			frontMD_GeometricObjects_PropertyType, _ = front.(*MD_GeometricObjects_PropertyType)
		}
		stage.OnAfterMD_GeometricObjects_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_geometricobjects_propertytype, frontMD_GeometricObjects_PropertyType)
	}
}

func (md_geometricobjects_propertytype *MD_GeometricObjects_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GeometricObjects_PropertyTypeDeleteCallback != nil {
		var frontMD_GeometricObjects_PropertyType *MD_GeometricObjects_PropertyType
		if front != nil {
			frontMD_GeometricObjects_PropertyType, _ = front.(*MD_GeometricObjects_PropertyType)
		}
		stage.OnAfterMD_GeometricObjects_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_geometricobjects_propertytype, frontMD_GeometricObjects_PropertyType)
	}
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_GeometricObjects_TypeCreateCallback != nil {
		stage.OnAfterMD_GeometricObjects_TypeCreateCallback.OnAfterCreate(stage, md_geometricobjects_type)
	}
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GeometricObjects_TypeUpdateCallback != nil {
		var frontMD_GeometricObjects_Type *MD_GeometricObjects_Type
		if front != nil {
			frontMD_GeometricObjects_Type, _ = front.(*MD_GeometricObjects_Type)
		}
		stage.OnAfterMD_GeometricObjects_TypeUpdateCallback.OnAfterUpdate(stage, md_geometricobjects_type, frontMD_GeometricObjects_Type)
	}
}

func (md_geometricobjects_type *MD_GeometricObjects_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GeometricObjects_TypeDeleteCallback != nil {
		var frontMD_GeometricObjects_Type *MD_GeometricObjects_Type
		if front != nil {
			frontMD_GeometricObjects_Type, _ = front.(*MD_GeometricObjects_Type)
		}
		stage.OnAfterMD_GeometricObjects_TypeDeleteCallback.OnAfterDelete(stage, md_geometricobjects_type, frontMD_GeometricObjects_Type)
	}
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Georectified_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Georectified_PropertyTypeCreateCallback.OnAfterCreate(stage, md_georectified_propertytype)
	}
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Georectified_PropertyTypeUpdateCallback != nil {
		var frontMD_Georectified_PropertyType *MD_Georectified_PropertyType
		if front != nil {
			frontMD_Georectified_PropertyType, _ = front.(*MD_Georectified_PropertyType)
		}
		stage.OnAfterMD_Georectified_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_georectified_propertytype, frontMD_Georectified_PropertyType)
	}
}

func (md_georectified_propertytype *MD_Georectified_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Georectified_PropertyTypeDeleteCallback != nil {
		var frontMD_Georectified_PropertyType *MD_Georectified_PropertyType
		if front != nil {
			frontMD_Georectified_PropertyType, _ = front.(*MD_Georectified_PropertyType)
		}
		stage.OnAfterMD_Georectified_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_georectified_propertytype, frontMD_Georectified_PropertyType)
	}
}

func (md_georectified_type *MD_Georectified_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Georectified_TypeCreateCallback != nil {
		stage.OnAfterMD_Georectified_TypeCreateCallback.OnAfterCreate(stage, md_georectified_type)
	}
}

func (md_georectified_type *MD_Georectified_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Georectified_TypeUpdateCallback != nil {
		var frontMD_Georectified_Type *MD_Georectified_Type
		if front != nil {
			frontMD_Georectified_Type, _ = front.(*MD_Georectified_Type)
		}
		stage.OnAfterMD_Georectified_TypeUpdateCallback.OnAfterUpdate(stage, md_georectified_type, frontMD_Georectified_Type)
	}
}

func (md_georectified_type *MD_Georectified_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Georectified_TypeDeleteCallback != nil {
		var frontMD_Georectified_Type *MD_Georectified_Type
		if front != nil {
			frontMD_Georectified_Type, _ = front.(*MD_Georectified_Type)
		}
		stage.OnAfterMD_Georectified_TypeDeleteCallback.OnAfterDelete(stage, md_georectified_type, frontMD_Georectified_Type)
	}
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Georeferenceable_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Georeferenceable_PropertyTypeCreateCallback.OnAfterCreate(stage, md_georeferenceable_propertytype)
	}
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Georeferenceable_PropertyTypeUpdateCallback != nil {
		var frontMD_Georeferenceable_PropertyType *MD_Georeferenceable_PropertyType
		if front != nil {
			frontMD_Georeferenceable_PropertyType, _ = front.(*MD_Georeferenceable_PropertyType)
		}
		stage.OnAfterMD_Georeferenceable_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_georeferenceable_propertytype, frontMD_Georeferenceable_PropertyType)
	}
}

func (md_georeferenceable_propertytype *MD_Georeferenceable_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Georeferenceable_PropertyTypeDeleteCallback != nil {
		var frontMD_Georeferenceable_PropertyType *MD_Georeferenceable_PropertyType
		if front != nil {
			frontMD_Georeferenceable_PropertyType, _ = front.(*MD_Georeferenceable_PropertyType)
		}
		stage.OnAfterMD_Georeferenceable_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_georeferenceable_propertytype, frontMD_Georeferenceable_PropertyType)
	}
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Georeferenceable_TypeCreateCallback != nil {
		stage.OnAfterMD_Georeferenceable_TypeCreateCallback.OnAfterCreate(stage, md_georeferenceable_type)
	}
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Georeferenceable_TypeUpdateCallback != nil {
		var frontMD_Georeferenceable_Type *MD_Georeferenceable_Type
		if front != nil {
			frontMD_Georeferenceable_Type, _ = front.(*MD_Georeferenceable_Type)
		}
		stage.OnAfterMD_Georeferenceable_TypeUpdateCallback.OnAfterUpdate(stage, md_georeferenceable_type, frontMD_Georeferenceable_Type)
	}
}

func (md_georeferenceable_type *MD_Georeferenceable_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Georeferenceable_TypeDeleteCallback != nil {
		var frontMD_Georeferenceable_Type *MD_Georeferenceable_Type
		if front != nil {
			frontMD_Georeferenceable_Type, _ = front.(*MD_Georeferenceable_Type)
		}
		stage.OnAfterMD_Georeferenceable_TypeDeleteCallback.OnAfterDelete(stage, md_georeferenceable_type, frontMD_Georeferenceable_Type)
	}
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_GridSpatialRepresentation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_GridSpatialRepresentation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_gridspatialrepresentation_propertytype)
	}
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GridSpatialRepresentation_PropertyTypeUpdateCallback != nil {
		var frontMD_GridSpatialRepresentation_PropertyType *MD_GridSpatialRepresentation_PropertyType
		if front != nil {
			frontMD_GridSpatialRepresentation_PropertyType, _ = front.(*MD_GridSpatialRepresentation_PropertyType)
		}
		stage.OnAfterMD_GridSpatialRepresentation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_gridspatialrepresentation_propertytype, frontMD_GridSpatialRepresentation_PropertyType)
	}
}

func (md_gridspatialrepresentation_propertytype *MD_GridSpatialRepresentation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GridSpatialRepresentation_PropertyTypeDeleteCallback != nil {
		var frontMD_GridSpatialRepresentation_PropertyType *MD_GridSpatialRepresentation_PropertyType
		if front != nil {
			frontMD_GridSpatialRepresentation_PropertyType, _ = front.(*MD_GridSpatialRepresentation_PropertyType)
		}
		stage.OnAfterMD_GridSpatialRepresentation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_gridspatialrepresentation_propertytype, frontMD_GridSpatialRepresentation_PropertyType)
	}
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_GridSpatialRepresentation_TypeCreateCallback != nil {
		stage.OnAfterMD_GridSpatialRepresentation_TypeCreateCallback.OnAfterCreate(stage, md_gridspatialrepresentation_type)
	}
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GridSpatialRepresentation_TypeUpdateCallback != nil {
		var frontMD_GridSpatialRepresentation_Type *MD_GridSpatialRepresentation_Type
		if front != nil {
			frontMD_GridSpatialRepresentation_Type, _ = front.(*MD_GridSpatialRepresentation_Type)
		}
		stage.OnAfterMD_GridSpatialRepresentation_TypeUpdateCallback.OnAfterUpdate(stage, md_gridspatialrepresentation_type, frontMD_GridSpatialRepresentation_Type)
	}
}

func (md_gridspatialrepresentation_type *MD_GridSpatialRepresentation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_GridSpatialRepresentation_TypeDeleteCallback != nil {
		var frontMD_GridSpatialRepresentation_Type *MD_GridSpatialRepresentation_Type
		if front != nil {
			frontMD_GridSpatialRepresentation_Type, _ = front.(*MD_GridSpatialRepresentation_Type)
		}
		stage.OnAfterMD_GridSpatialRepresentation_TypeDeleteCallback.OnAfterDelete(stage, md_gridspatialrepresentation_type, frontMD_GridSpatialRepresentation_Type)
	}
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_PixelOrientationCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_PixelOrientationCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_pixelorientationcode_propertytype)
	}
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_PixelOrientationCode_PropertyTypeUpdateCallback != nil {
		var frontMD_PixelOrientationCode_PropertyType *MD_PixelOrientationCode_PropertyType
		if front != nil {
			frontMD_PixelOrientationCode_PropertyType, _ = front.(*MD_PixelOrientationCode_PropertyType)
		}
		stage.OnAfterMD_PixelOrientationCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_pixelorientationcode_propertytype, frontMD_PixelOrientationCode_PropertyType)
	}
}

func (md_pixelorientationcode_propertytype *MD_PixelOrientationCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_PixelOrientationCode_PropertyTypeDeleteCallback != nil {
		var frontMD_PixelOrientationCode_PropertyType *MD_PixelOrientationCode_PropertyType
		if front != nil {
			frontMD_PixelOrientationCode_PropertyType, _ = front.(*MD_PixelOrientationCode_PropertyType)
		}
		stage.OnAfterMD_PixelOrientationCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_pixelorientationcode_propertytype, frontMD_PixelOrientationCode_PropertyType)
	}
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_SpatialRepresentation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_SpatialRepresentation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_spatialrepresentation_propertytype)
	}
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_SpatialRepresentation_PropertyTypeUpdateCallback != nil {
		var frontMD_SpatialRepresentation_PropertyType *MD_SpatialRepresentation_PropertyType
		if front != nil {
			frontMD_SpatialRepresentation_PropertyType, _ = front.(*MD_SpatialRepresentation_PropertyType)
		}
		stage.OnAfterMD_SpatialRepresentation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_spatialrepresentation_propertytype, frontMD_SpatialRepresentation_PropertyType)
	}
}

func (md_spatialrepresentation_propertytype *MD_SpatialRepresentation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_SpatialRepresentation_PropertyTypeDeleteCallback != nil {
		var frontMD_SpatialRepresentation_PropertyType *MD_SpatialRepresentation_PropertyType
		if front != nil {
			frontMD_SpatialRepresentation_PropertyType, _ = front.(*MD_SpatialRepresentation_PropertyType)
		}
		stage.OnAfterMD_SpatialRepresentation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_spatialrepresentation_propertytype, frontMD_SpatialRepresentation_PropertyType)
	}
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_TopologyLevelCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_TopologyLevelCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_topologylevelcode_propertytype)
	}
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_TopologyLevelCode_PropertyTypeUpdateCallback != nil {
		var frontMD_TopologyLevelCode_PropertyType *MD_TopologyLevelCode_PropertyType
		if front != nil {
			frontMD_TopologyLevelCode_PropertyType, _ = front.(*MD_TopologyLevelCode_PropertyType)
		}
		stage.OnAfterMD_TopologyLevelCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_topologylevelcode_propertytype, frontMD_TopologyLevelCode_PropertyType)
	}
}

func (md_topologylevelcode_propertytype *MD_TopologyLevelCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_TopologyLevelCode_PropertyTypeDeleteCallback != nil {
		var frontMD_TopologyLevelCode_PropertyType *MD_TopologyLevelCode_PropertyType
		if front != nil {
			frontMD_TopologyLevelCode_PropertyType, _ = front.(*MD_TopologyLevelCode_PropertyType)
		}
		stage.OnAfterMD_TopologyLevelCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_topologylevelcode_propertytype, frontMD_TopologyLevelCode_PropertyType)
	}
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_VectorSpatialRepresentation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_VectorSpatialRepresentation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_vectorspatialrepresentation_propertytype)
	}
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_VectorSpatialRepresentation_PropertyTypeUpdateCallback != nil {
		var frontMD_VectorSpatialRepresentation_PropertyType *MD_VectorSpatialRepresentation_PropertyType
		if front != nil {
			frontMD_VectorSpatialRepresentation_PropertyType, _ = front.(*MD_VectorSpatialRepresentation_PropertyType)
		}
		stage.OnAfterMD_VectorSpatialRepresentation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_vectorspatialrepresentation_propertytype, frontMD_VectorSpatialRepresentation_PropertyType)
	}
}

func (md_vectorspatialrepresentation_propertytype *MD_VectorSpatialRepresentation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_VectorSpatialRepresentation_PropertyTypeDeleteCallback != nil {
		var frontMD_VectorSpatialRepresentation_PropertyType *MD_VectorSpatialRepresentation_PropertyType
		if front != nil {
			frontMD_VectorSpatialRepresentation_PropertyType, _ = front.(*MD_VectorSpatialRepresentation_PropertyType)
		}
		stage.OnAfterMD_VectorSpatialRepresentation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_vectorspatialrepresentation_propertytype, frontMD_VectorSpatialRepresentation_PropertyType)
	}
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_VectorSpatialRepresentation_TypeCreateCallback != nil {
		stage.OnAfterMD_VectorSpatialRepresentation_TypeCreateCallback.OnAfterCreate(stage, md_vectorspatialrepresentation_type)
	}
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_VectorSpatialRepresentation_TypeUpdateCallback != nil {
		var frontMD_VectorSpatialRepresentation_Type *MD_VectorSpatialRepresentation_Type
		if front != nil {
			frontMD_VectorSpatialRepresentation_Type, _ = front.(*MD_VectorSpatialRepresentation_Type)
		}
		stage.OnAfterMD_VectorSpatialRepresentation_TypeUpdateCallback.OnAfterUpdate(stage, md_vectorspatialrepresentation_type, frontMD_VectorSpatialRepresentation_Type)
	}
}

func (md_vectorspatialrepresentation_type *MD_VectorSpatialRepresentation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_VectorSpatialRepresentation_TypeDeleteCallback != nil {
		var frontMD_VectorSpatialRepresentation_Type *MD_VectorSpatialRepresentation_Type
		if front != nil {
			frontMD_VectorSpatialRepresentation_Type, _ = front.(*MD_VectorSpatialRepresentation_Type)
		}
		stage.OnAfterMD_VectorSpatialRepresentation_TypeDeleteCallback.OnAfterDelete(stage, md_vectorspatialrepresentation_type, frontMD_VectorSpatialRepresentation_Type)
	}
}

