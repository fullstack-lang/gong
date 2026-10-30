// generated code - do not edit
package gml

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
func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractCurveSegmentTypeCreateCallback != nil {
		stage.OnAfterAbstractCurveSegmentTypeCreateCallback.OnAfterCreate(stage, abstractcurvesegmenttype)
	}
}

func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCurveSegmentTypeUpdateCallback != nil {
		var frontAbstractCurveSegmentType *AbstractCurveSegmentType
		if front != nil {
			frontAbstractCurveSegmentType, _ = front.(*AbstractCurveSegmentType)
		}
		stage.OnAfterAbstractCurveSegmentTypeUpdateCallback.OnAfterUpdate(stage, abstractcurvesegmenttype, frontAbstractCurveSegmentType)
	}
}

func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCurveSegmentTypeDeleteCallback != nil {
		var frontAbstractCurveSegmentType *AbstractCurveSegmentType
		if front != nil {
			frontAbstractCurveSegmentType, _ = front.(*AbstractCurveSegmentType)
		}
		stage.OnAfterAbstractCurveSegmentTypeDeleteCallback.OnAfterDelete(stage, abstractcurvesegmenttype, frontAbstractCurveSegmentType)
	}
}

func (abstractcurvetype *AbstractCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractCurveTypeCreateCallback != nil {
		stage.OnAfterAbstractCurveTypeCreateCallback.OnAfterCreate(stage, abstractcurvetype)
	}
}

func (abstractcurvetype *AbstractCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCurveTypeUpdateCallback != nil {
		var frontAbstractCurveType *AbstractCurveType
		if front != nil {
			frontAbstractCurveType, _ = front.(*AbstractCurveType)
		}
		stage.OnAfterAbstractCurveTypeUpdateCallback.OnAfterUpdate(stage, abstractcurvetype, frontAbstractCurveType)
	}
}

func (abstractcurvetype *AbstractCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCurveTypeDeleteCallback != nil {
		var frontAbstractCurveType *AbstractCurveType
		if front != nil {
			frontAbstractCurveType, _ = front.(*AbstractCurveType)
		}
		stage.OnAfterAbstractCurveTypeDeleteCallback.OnAfterDelete(stage, abstractcurvetype, frontAbstractCurveType)
	}
}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractFeatureMemberTypeCreateCallback != nil {
		stage.OnAfterAbstractFeatureMemberTypeCreateCallback.OnAfterCreate(stage, abstractfeaturemembertype)
	}
}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureMemberTypeUpdateCallback != nil {
		var frontAbstractFeatureMemberType *AbstractFeatureMemberType
		if front != nil {
			frontAbstractFeatureMemberType, _ = front.(*AbstractFeatureMemberType)
		}
		stage.OnAfterAbstractFeatureMemberTypeUpdateCallback.OnAfterUpdate(stage, abstractfeaturemembertype, frontAbstractFeatureMemberType)
	}
}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureMemberTypeDeleteCallback != nil {
		var frontAbstractFeatureMemberType *AbstractFeatureMemberType
		if front != nil {
			frontAbstractFeatureMemberType, _ = front.(*AbstractFeatureMemberType)
		}
		stage.OnAfterAbstractFeatureMemberTypeDeleteCallback.OnAfterDelete(stage, abstractfeaturemembertype, frontAbstractFeatureMemberType)
	}
}

func (abstractfeaturetype *AbstractFeatureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractFeatureTypeCreateCallback != nil {
		stage.OnAfterAbstractFeatureTypeCreateCallback.OnAfterCreate(stage, abstractfeaturetype)
	}
}

func (abstractfeaturetype *AbstractFeatureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureTypeUpdateCallback != nil {
		var frontAbstractFeatureType *AbstractFeatureType
		if front != nil {
			frontAbstractFeatureType, _ = front.(*AbstractFeatureType)
		}
		stage.OnAfterAbstractFeatureTypeUpdateCallback.OnAfterUpdate(stage, abstractfeaturetype, frontAbstractFeatureType)
	}
}

func (abstractfeaturetype *AbstractFeatureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureTypeDeleteCallback != nil {
		var frontAbstractFeatureType *AbstractFeatureType
		if front != nil {
			frontAbstractFeatureType, _ = front.(*AbstractFeatureType)
		}
		stage.OnAfterAbstractFeatureTypeDeleteCallback.OnAfterDelete(stage, abstractfeaturetype, frontAbstractFeatureType)
	}
}

func (abstractgmltype *AbstractGMLType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGMLTypeCreateCallback != nil {
		stage.OnAfterAbstractGMLTypeCreateCallback.OnAfterCreate(stage, abstractgmltype)
	}
}

func (abstractgmltype *AbstractGMLType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGMLTypeUpdateCallback != nil {
		var frontAbstractGMLType *AbstractGMLType
		if front != nil {
			frontAbstractGMLType, _ = front.(*AbstractGMLType)
		}
		stage.OnAfterAbstractGMLTypeUpdateCallback.OnAfterUpdate(stage, abstractgmltype, frontAbstractGMLType)
	}
}

func (abstractgmltype *AbstractGMLType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGMLTypeDeleteCallback != nil {
		var frontAbstractGMLType *AbstractGMLType
		if front != nil {
			frontAbstractGMLType, _ = front.(*AbstractGMLType)
		}
		stage.OnAfterAbstractGMLTypeDeleteCallback.OnAfterDelete(stage, abstractgmltype, frontAbstractGMLType)
	}
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeometricPrimitiveTypeCreateCallback != nil {
		stage.OnAfterAbstractGeometricPrimitiveTypeCreateCallback.OnAfterCreate(stage, abstractgeometricprimitivetype)
	}
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometricPrimitiveTypeUpdateCallback != nil {
		var frontAbstractGeometricPrimitiveType *AbstractGeometricPrimitiveType
		if front != nil {
			frontAbstractGeometricPrimitiveType, _ = front.(*AbstractGeometricPrimitiveType)
		}
		stage.OnAfterAbstractGeometricPrimitiveTypeUpdateCallback.OnAfterUpdate(stage, abstractgeometricprimitivetype, frontAbstractGeometricPrimitiveType)
	}
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometricPrimitiveTypeDeleteCallback != nil {
		var frontAbstractGeometricPrimitiveType *AbstractGeometricPrimitiveType
		if front != nil {
			frontAbstractGeometricPrimitiveType, _ = front.(*AbstractGeometricPrimitiveType)
		}
		stage.OnAfterAbstractGeometricPrimitiveTypeDeleteCallback.OnAfterDelete(stage, abstractgeometricprimitivetype, frontAbstractGeometricPrimitiveType)
	}
}

func (abstractgeometrytype *AbstractGeometryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeometryTypeCreateCallback != nil {
		stage.OnAfterAbstractGeometryTypeCreateCallback.OnAfterCreate(stage, abstractgeometrytype)
	}
}

func (abstractgeometrytype *AbstractGeometryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometryTypeUpdateCallback != nil {
		var frontAbstractGeometryType *AbstractGeometryType
		if front != nil {
			frontAbstractGeometryType, _ = front.(*AbstractGeometryType)
		}
		stage.OnAfterAbstractGeometryTypeUpdateCallback.OnAfterUpdate(stage, abstractgeometrytype, frontAbstractGeometryType)
	}
}

func (abstractgeometrytype *AbstractGeometryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometryTypeDeleteCallback != nil {
		var frontAbstractGeometryType *AbstractGeometryType
		if front != nil {
			frontAbstractGeometryType, _ = front.(*AbstractGeometryType)
		}
		stage.OnAfterAbstractGeometryTypeDeleteCallback.OnAfterDelete(stage, abstractgeometrytype, frontAbstractGeometryType)
	}
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractMetadataPropertyTypeCreateCallback != nil {
		stage.OnAfterAbstractMetadataPropertyTypeCreateCallback.OnAfterCreate(stage, abstractmetadatapropertytype)
	}
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMetadataPropertyTypeUpdateCallback != nil {
		var frontAbstractMetadataPropertyType *AbstractMetadataPropertyType
		if front != nil {
			frontAbstractMetadataPropertyType, _ = front.(*AbstractMetadataPropertyType)
		}
		stage.OnAfterAbstractMetadataPropertyTypeUpdateCallback.OnAfterUpdate(stage, abstractmetadatapropertytype, frontAbstractMetadataPropertyType)
	}
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMetadataPropertyTypeDeleteCallback != nil {
		var frontAbstractMetadataPropertyType *AbstractMetadataPropertyType
		if front != nil {
			frontAbstractMetadataPropertyType, _ = front.(*AbstractMetadataPropertyType)
		}
		stage.OnAfterAbstractMetadataPropertyTypeDeleteCallback.OnAfterDelete(stage, abstractmetadatapropertytype, frontAbstractMetadataPropertyType)
	}
}

func (abstractringpropertytype *AbstractRingPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractRingPropertyTypeCreateCallback != nil {
		stage.OnAfterAbstractRingPropertyTypeCreateCallback.OnAfterCreate(stage, abstractringpropertytype)
	}
}

func (abstractringpropertytype *AbstractRingPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRingPropertyTypeUpdateCallback != nil {
		var frontAbstractRingPropertyType *AbstractRingPropertyType
		if front != nil {
			frontAbstractRingPropertyType, _ = front.(*AbstractRingPropertyType)
		}
		stage.OnAfterAbstractRingPropertyTypeUpdateCallback.OnAfterUpdate(stage, abstractringpropertytype, frontAbstractRingPropertyType)
	}
}

func (abstractringpropertytype *AbstractRingPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRingPropertyTypeDeleteCallback != nil {
		var frontAbstractRingPropertyType *AbstractRingPropertyType
		if front != nil {
			frontAbstractRingPropertyType, _ = front.(*AbstractRingPropertyType)
		}
		stage.OnAfterAbstractRingPropertyTypeDeleteCallback.OnAfterDelete(stage, abstractringpropertytype, frontAbstractRingPropertyType)
	}
}

func (abstractringtype *AbstractRingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractRingTypeCreateCallback != nil {
		stage.OnAfterAbstractRingTypeCreateCallback.OnAfterCreate(stage, abstractringtype)
	}
}

func (abstractringtype *AbstractRingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRingTypeUpdateCallback != nil {
		var frontAbstractRingType *AbstractRingType
		if front != nil {
			frontAbstractRingType, _ = front.(*AbstractRingType)
		}
		stage.OnAfterAbstractRingTypeUpdateCallback.OnAfterUpdate(stage, abstractringtype, frontAbstractRingType)
	}
}

func (abstractringtype *AbstractRingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRingTypeDeleteCallback != nil {
		var frontAbstractRingType *AbstractRingType
		if front != nil {
			frontAbstractRingType, _ = front.(*AbstractRingType)
		}
		stage.OnAfterAbstractRingTypeDeleteCallback.OnAfterDelete(stage, abstractringtype, frontAbstractRingType)
	}
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSurfacePatchTypeCreateCallback != nil {
		stage.OnAfterAbstractSurfacePatchTypeCreateCallback.OnAfterCreate(stage, abstractsurfacepatchtype)
	}
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfacePatchTypeUpdateCallback != nil {
		var frontAbstractSurfacePatchType *AbstractSurfacePatchType
		if front != nil {
			frontAbstractSurfacePatchType, _ = front.(*AbstractSurfacePatchType)
		}
		stage.OnAfterAbstractSurfacePatchTypeUpdateCallback.OnAfterUpdate(stage, abstractsurfacepatchtype, frontAbstractSurfacePatchType)
	}
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfacePatchTypeDeleteCallback != nil {
		var frontAbstractSurfacePatchType *AbstractSurfacePatchType
		if front != nil {
			frontAbstractSurfacePatchType, _ = front.(*AbstractSurfacePatchType)
		}
		stage.OnAfterAbstractSurfacePatchTypeDeleteCallback.OnAfterDelete(stage, abstractsurfacepatchtype, frontAbstractSurfacePatchType)
	}
}

func (abstractsurfacetype *AbstractSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSurfaceTypeCreateCallback != nil {
		stage.OnAfterAbstractSurfaceTypeCreateCallback.OnAfterCreate(stage, abstractsurfacetype)
	}
}

func (abstractsurfacetype *AbstractSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfaceTypeUpdateCallback != nil {
		var frontAbstractSurfaceType *AbstractSurfaceType
		if front != nil {
			frontAbstractSurfaceType, _ = front.(*AbstractSurfaceType)
		}
		stage.OnAfterAbstractSurfaceTypeUpdateCallback.OnAfterUpdate(stage, abstractsurfacetype, frontAbstractSurfaceType)
	}
}

func (abstractsurfacetype *AbstractSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfaceTypeDeleteCallback != nil {
		var frontAbstractSurfaceType *AbstractSurfaceType
		if front != nil {
			frontAbstractSurfaceType, _ = front.(*AbstractSurfaceType)
		}
		stage.OnAfterAbstractSurfaceTypeDeleteCallback.OnAfterDelete(stage, abstractsurfacetype, frontAbstractSurfaceType)
	}
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimeGeometricPrimitiveTypeCreateCallback != nil {
		stage.OnAfterAbstractTimeGeometricPrimitiveTypeCreateCallback.OnAfterCreate(stage, abstracttimegeometricprimitivetype)
	}
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeGeometricPrimitiveTypeUpdateCallback != nil {
		var frontAbstractTimeGeometricPrimitiveType *AbstractTimeGeometricPrimitiveType
		if front != nil {
			frontAbstractTimeGeometricPrimitiveType, _ = front.(*AbstractTimeGeometricPrimitiveType)
		}
		stage.OnAfterAbstractTimeGeometricPrimitiveTypeUpdateCallback.OnAfterUpdate(stage, abstracttimegeometricprimitivetype, frontAbstractTimeGeometricPrimitiveType)
	}
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeGeometricPrimitiveTypeDeleteCallback != nil {
		var frontAbstractTimeGeometricPrimitiveType *AbstractTimeGeometricPrimitiveType
		if front != nil {
			frontAbstractTimeGeometricPrimitiveType, _ = front.(*AbstractTimeGeometricPrimitiveType)
		}
		stage.OnAfterAbstractTimeGeometricPrimitiveTypeDeleteCallback.OnAfterDelete(stage, abstracttimegeometricprimitivetype, frontAbstractTimeGeometricPrimitiveType)
	}
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimeObjectTypeCreateCallback != nil {
		stage.OnAfterAbstractTimeObjectTypeCreateCallback.OnAfterCreate(stage, abstracttimeobjecttype)
	}
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeObjectTypeUpdateCallback != nil {
		var frontAbstractTimeObjectType *AbstractTimeObjectType
		if front != nil {
			frontAbstractTimeObjectType, _ = front.(*AbstractTimeObjectType)
		}
		stage.OnAfterAbstractTimeObjectTypeUpdateCallback.OnAfterUpdate(stage, abstracttimeobjecttype, frontAbstractTimeObjectType)
	}
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeObjectTypeDeleteCallback != nil {
		var frontAbstractTimeObjectType *AbstractTimeObjectType
		if front != nil {
			frontAbstractTimeObjectType, _ = front.(*AbstractTimeObjectType)
		}
		stage.OnAfterAbstractTimeObjectTypeDeleteCallback.OnAfterDelete(stage, abstracttimeobjecttype, frontAbstractTimeObjectType)
	}
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimePrimitiveTypeCreateCallback != nil {
		stage.OnAfterAbstractTimePrimitiveTypeCreateCallback.OnAfterCreate(stage, abstracttimeprimitivetype)
	}
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimePrimitiveTypeUpdateCallback != nil {
		var frontAbstractTimePrimitiveType *AbstractTimePrimitiveType
		if front != nil {
			frontAbstractTimePrimitiveType, _ = front.(*AbstractTimePrimitiveType)
		}
		stage.OnAfterAbstractTimePrimitiveTypeUpdateCallback.OnAfterUpdate(stage, abstracttimeprimitivetype, frontAbstractTimePrimitiveType)
	}
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimePrimitiveTypeDeleteCallback != nil {
		var frontAbstractTimePrimitiveType *AbstractTimePrimitiveType
		if front != nil {
			frontAbstractTimePrimitiveType, _ = front.(*AbstractTimePrimitiveType)
		}
		stage.OnAfterAbstractTimePrimitiveTypeDeleteCallback.OnAfterDelete(stage, abstracttimeprimitivetype, frontAbstractTimePrimitiveType)
	}
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAbstractTimeSliceTypeCreateCallback.OnAfterCreate(stage, abstracttimeslicetype)
	}
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeSliceTypeUpdateCallback != nil {
		var frontAbstractTimeSliceType *AbstractTimeSliceType
		if front != nil {
			frontAbstractTimeSliceType, _ = front.(*AbstractTimeSliceType)
		}
		stage.OnAfterAbstractTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, abstracttimeslicetype, frontAbstractTimeSliceType)
	}
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeSliceTypeDeleteCallback != nil {
		var frontAbstractTimeSliceType *AbstractTimeSliceType
		if front != nil {
			frontAbstractTimeSliceType, _ = front.(*AbstractTimeSliceType)
		}
		stage.OnAfterAbstractTimeSliceTypeDeleteCallback.OnAfterDelete(stage, abstracttimeslicetype, frontAbstractTimeSliceType)
	}
}

func (angletype *AngleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleTypeCreateCallback != nil {
		stage.OnAfterAngleTypeCreateCallback.OnAfterCreate(stage, angletype)
	}
}

func (angletype *AngleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleTypeUpdateCallback != nil {
		var frontAngleType *AngleType
		if front != nil {
			frontAngleType, _ = front.(*AngleType)
		}
		stage.OnAfterAngleTypeUpdateCallback.OnAfterUpdate(stage, angletype, frontAngleType)
	}
}

func (angletype *AngleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleTypeDeleteCallback != nil {
		var frontAngleType *AngleType
		if front != nil {
			frontAngleType, _ = front.(*AngleType)
		}
		stage.OnAfterAngleTypeDeleteCallback.OnAfterDelete(stage, angletype, frontAngleType)
	}
}

func (arcbycenterpointtype *ArcByCenterPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcByCenterPointTypeCreateCallback != nil {
		stage.OnAfterArcByCenterPointTypeCreateCallback.OnAfterCreate(stage, arcbycenterpointtype)
	}
}

func (arcbycenterpointtype *ArcByCenterPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcByCenterPointTypeUpdateCallback != nil {
		var frontArcByCenterPointType *ArcByCenterPointType
		if front != nil {
			frontArcByCenterPointType, _ = front.(*ArcByCenterPointType)
		}
		stage.OnAfterArcByCenterPointTypeUpdateCallback.OnAfterUpdate(stage, arcbycenterpointtype, frontArcByCenterPointType)
	}
}

func (arcbycenterpointtype *ArcByCenterPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcByCenterPointTypeDeleteCallback != nil {
		var frontArcByCenterPointType *ArcByCenterPointType
		if front != nil {
			frontArcByCenterPointType, _ = front.(*ArcByCenterPointType)
		}
		stage.OnAfterArcByCenterPointTypeDeleteCallback.OnAfterDelete(stage, arcbycenterpointtype, frontArcByCenterPointType)
	}
}

func (arcstringtype *ArcStringType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcStringTypeCreateCallback != nil {
		stage.OnAfterArcStringTypeCreateCallback.OnAfterCreate(stage, arcstringtype)
	}
}

func (arcstringtype *ArcStringType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcStringTypeUpdateCallback != nil {
		var frontArcStringType *ArcStringType
		if front != nil {
			frontArcStringType, _ = front.(*ArcStringType)
		}
		stage.OnAfterArcStringTypeUpdateCallback.OnAfterUpdate(stage, arcstringtype, frontArcStringType)
	}
}

func (arcstringtype *ArcStringType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcStringTypeDeleteCallback != nil {
		var frontArcStringType *ArcStringType
		if front != nil {
			frontArcStringType, _ = front.(*ArcStringType)
		}
		stage.OnAfterArcStringTypeDeleteCallback.OnAfterDelete(stage, arcstringtype, frontArcStringType)
	}
}

func (arctype *ArcType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcTypeCreateCallback != nil {
		stage.OnAfterArcTypeCreateCallback.OnAfterCreate(stage, arctype)
	}
}

func (arctype *ArcType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcTypeUpdateCallback != nil {
		var frontArcType *ArcType
		if front != nil {
			frontArcType, _ = front.(*ArcType)
		}
		stage.OnAfterArcTypeUpdateCallback.OnAfterUpdate(stage, arctype, frontArcType)
	}
}

func (arctype *ArcType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcTypeDeleteCallback != nil {
		var frontArcType *ArcType
		if front != nil {
			frontArcType, _ = front.(*ArcType)
		}
		stage.OnAfterArcTypeDeleteCallback.OnAfterDelete(stage, arctype, frontArcType)
	}
}

func (boundingshapetype *BoundingShapeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBoundingShapeTypeCreateCallback != nil {
		stage.OnAfterBoundingShapeTypeCreateCallback.OnAfterCreate(stage, boundingshapetype)
	}
}

func (boundingshapetype *BoundingShapeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoundingShapeTypeUpdateCallback != nil {
		var frontBoundingShapeType *BoundingShapeType
		if front != nil {
			frontBoundingShapeType, _ = front.(*BoundingShapeType)
		}
		stage.OnAfterBoundingShapeTypeUpdateCallback.OnAfterUpdate(stage, boundingshapetype, frontBoundingShapeType)
	}
}

func (boundingshapetype *BoundingShapeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoundingShapeTypeDeleteCallback != nil {
		var frontBoundingShapeType *BoundingShapeType
		if front != nil {
			frontBoundingShapeType, _ = front.(*BoundingShapeType)
		}
		stage.OnAfterBoundingShapeTypeDeleteCallback.OnAfterDelete(stage, boundingshapetype, frontBoundingShapeType)
	}
}

func (circlebycenterpointtype *CircleByCenterPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircleByCenterPointTypeCreateCallback != nil {
		stage.OnAfterCircleByCenterPointTypeCreateCallback.OnAfterCreate(stage, circlebycenterpointtype)
	}
}

func (circlebycenterpointtype *CircleByCenterPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleByCenterPointTypeUpdateCallback != nil {
		var frontCircleByCenterPointType *CircleByCenterPointType
		if front != nil {
			frontCircleByCenterPointType, _ = front.(*CircleByCenterPointType)
		}
		stage.OnAfterCircleByCenterPointTypeUpdateCallback.OnAfterUpdate(stage, circlebycenterpointtype, frontCircleByCenterPointType)
	}
}

func (circlebycenterpointtype *CircleByCenterPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleByCenterPointTypeDeleteCallback != nil {
		var frontCircleByCenterPointType *CircleByCenterPointType
		if front != nil {
			frontCircleByCenterPointType, _ = front.(*CircleByCenterPointType)
		}
		stage.OnAfterCircleByCenterPointTypeDeleteCallback.OnAfterDelete(stage, circlebycenterpointtype, frontCircleByCenterPointType)
	}
}

func (circletype *CircleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircleTypeCreateCallback != nil {
		stage.OnAfterCircleTypeCreateCallback.OnAfterCreate(stage, circletype)
	}
}

func (circletype *CircleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleTypeUpdateCallback != nil {
		var frontCircleType *CircleType
		if front != nil {
			frontCircleType, _ = front.(*CircleType)
		}
		stage.OnAfterCircleTypeUpdateCallback.OnAfterUpdate(stage, circletype, frontCircleType)
	}
}

func (circletype *CircleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleTypeDeleteCallback != nil {
		var frontCircleType *CircleType
		if front != nil {
			frontCircleType, _ = front.(*CircleType)
		}
		stage.OnAfterCircleTypeDeleteCallback.OnAfterDelete(stage, circletype, frontCircleType)
	}
}

func (codetype *CodeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTypeCreateCallback != nil {
		stage.OnAfterCodeTypeCreateCallback.OnAfterCreate(stage, codetype)
	}
}

func (codetype *CodeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTypeUpdateCallback != nil {
		var frontCodeType *CodeType
		if front != nil {
			frontCodeType, _ = front.(*CodeType)
		}
		stage.OnAfterCodeTypeUpdateCallback.OnAfterUpdate(stage, codetype, frontCodeType)
	}
}

func (codetype *CodeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTypeDeleteCallback != nil {
		var frontCodeType *CodeType
		if front != nil {
			frontCodeType, _ = front.(*CodeType)
		}
		stage.OnAfterCodeTypeDeleteCallback.OnAfterDelete(stage, codetype, frontCodeType)
	}
}

func (codewithauthoritytype *CodeWithAuthorityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeWithAuthorityTypeCreateCallback != nil {
		stage.OnAfterCodeWithAuthorityTypeCreateCallback.OnAfterCreate(stage, codewithauthoritytype)
	}
}

func (codewithauthoritytype *CodeWithAuthorityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeWithAuthorityTypeUpdateCallback != nil {
		var frontCodeWithAuthorityType *CodeWithAuthorityType
		if front != nil {
			frontCodeWithAuthorityType, _ = front.(*CodeWithAuthorityType)
		}
		stage.OnAfterCodeWithAuthorityTypeUpdateCallback.OnAfterUpdate(stage, codewithauthoritytype, frontCodeWithAuthorityType)
	}
}

func (codewithauthoritytype *CodeWithAuthorityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeWithAuthorityTypeDeleteCallback != nil {
		var frontCodeWithAuthorityType *CodeWithAuthorityType
		if front != nil {
			frontCodeWithAuthorityType, _ = front.(*CodeWithAuthorityType)
		}
		stage.OnAfterCodeWithAuthorityTypeDeleteCallback.OnAfterDelete(stage, codewithauthoritytype, frontCodeWithAuthorityType)
	}
}

func (compositecurvetype *CompositeCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCompositeCurveTypeCreateCallback != nil {
		stage.OnAfterCompositeCurveTypeCreateCallback.OnAfterCreate(stage, compositecurvetype)
	}
}

func (compositecurvetype *CompositeCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeCurveTypeUpdateCallback != nil {
		var frontCompositeCurveType *CompositeCurveType
		if front != nil {
			frontCompositeCurveType, _ = front.(*CompositeCurveType)
		}
		stage.OnAfterCompositeCurveTypeUpdateCallback.OnAfterUpdate(stage, compositecurvetype, frontCompositeCurveType)
	}
}

func (compositecurvetype *CompositeCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeCurveTypeDeleteCallback != nil {
		var frontCompositeCurveType *CompositeCurveType
		if front != nil {
			frontCompositeCurveType, _ = front.(*CompositeCurveType)
		}
		stage.OnAfterCompositeCurveTypeDeleteCallback.OnAfterDelete(stage, compositecurvetype, frontCompositeCurveType)
	}
}

func (curvepropertytype *CurvePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurvePropertyTypeCreateCallback != nil {
		stage.OnAfterCurvePropertyTypeCreateCallback.OnAfterCreate(stage, curvepropertytype)
	}
}

func (curvepropertytype *CurvePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurvePropertyTypeUpdateCallback != nil {
		var frontCurvePropertyType *CurvePropertyType
		if front != nil {
			frontCurvePropertyType, _ = front.(*CurvePropertyType)
		}
		stage.OnAfterCurvePropertyTypeUpdateCallback.OnAfterUpdate(stage, curvepropertytype, frontCurvePropertyType)
	}
}

func (curvepropertytype *CurvePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurvePropertyTypeDeleteCallback != nil {
		var frontCurvePropertyType *CurvePropertyType
		if front != nil {
			frontCurvePropertyType, _ = front.(*CurvePropertyType)
		}
		stage.OnAfterCurvePropertyTypeDeleteCallback.OnAfterDelete(stage, curvepropertytype, frontCurvePropertyType)
	}
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurveSegmentArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterCurveSegmentArrayPropertyTypeCreateCallback.OnAfterCreate(stage, curvesegmentarraypropertytype)
	}
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveSegmentArrayPropertyTypeUpdateCallback != nil {
		var frontCurveSegmentArrayPropertyType *CurveSegmentArrayPropertyType
		if front != nil {
			frontCurveSegmentArrayPropertyType, _ = front.(*CurveSegmentArrayPropertyType)
		}
		stage.OnAfterCurveSegmentArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, curvesegmentarraypropertytype, frontCurveSegmentArrayPropertyType)
	}
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveSegmentArrayPropertyTypeDeleteCallback != nil {
		var frontCurveSegmentArrayPropertyType *CurveSegmentArrayPropertyType
		if front != nil {
			frontCurveSegmentArrayPropertyType, _ = front.(*CurveSegmentArrayPropertyType)
		}
		stage.OnAfterCurveSegmentArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, curvesegmentarraypropertytype, frontCurveSegmentArrayPropertyType)
	}
}

func (curvetype *CurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurveTypeCreateCallback != nil {
		stage.OnAfterCurveTypeCreateCallback.OnAfterCreate(stage, curvetype)
	}
}

func (curvetype *CurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveTypeUpdateCallback != nil {
		var frontCurveType *CurveType
		if front != nil {
			frontCurveType, _ = front.(*CurveType)
		}
		stage.OnAfterCurveTypeUpdateCallback.OnAfterUpdate(stage, curvetype, frontCurveType)
	}
}

func (curvetype *CurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveTypeDeleteCallback != nil {
		var frontCurveType *CurveType
		if front != nil {
			frontCurveType, _ = front.(*CurveType)
		}
		stage.OnAfterCurveTypeDeleteCallback.OnAfterDelete(stage, curvetype, frontCurveType)
	}
}

func (directpositionlisttype *DirectPositionListType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectPositionListTypeCreateCallback != nil {
		stage.OnAfterDirectPositionListTypeCreateCallback.OnAfterCreate(stage, directpositionlisttype)
	}
}

func (directpositionlisttype *DirectPositionListType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectPositionListTypeUpdateCallback != nil {
		var frontDirectPositionListType *DirectPositionListType
		if front != nil {
			frontDirectPositionListType, _ = front.(*DirectPositionListType)
		}
		stage.OnAfterDirectPositionListTypeUpdateCallback.OnAfterUpdate(stage, directpositionlisttype, frontDirectPositionListType)
	}
}

func (directpositionlisttype *DirectPositionListType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectPositionListTypeDeleteCallback != nil {
		var frontDirectPositionListType *DirectPositionListType
		if front != nil {
			frontDirectPositionListType, _ = front.(*DirectPositionListType)
		}
		stage.OnAfterDirectPositionListTypeDeleteCallback.OnAfterDelete(stage, directpositionlisttype, frontDirectPositionListType)
	}
}

func (directpositiontype *DirectPositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectPositionTypeCreateCallback != nil {
		stage.OnAfterDirectPositionTypeCreateCallback.OnAfterCreate(stage, directpositiontype)
	}
}

func (directpositiontype *DirectPositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectPositionTypeUpdateCallback != nil {
		var frontDirectPositionType *DirectPositionType
		if front != nil {
			frontDirectPositionType, _ = front.(*DirectPositionType)
		}
		stage.OnAfterDirectPositionTypeUpdateCallback.OnAfterUpdate(stage, directpositiontype, frontDirectPositionType)
	}
}

func (directpositiontype *DirectPositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectPositionTypeDeleteCallback != nil {
		var frontDirectPositionType *DirectPositionType
		if front != nil {
			frontDirectPositionType, _ = front.(*DirectPositionType)
		}
		stage.OnAfterDirectPositionTypeDeleteCallback.OnAfterDelete(stage, directpositiontype, frontDirectPositionType)
	}
}

func (dynamicfeaturetype *DynamicFeatureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDynamicFeatureTypeCreateCallback != nil {
		stage.OnAfterDynamicFeatureTypeCreateCallback.OnAfterCreate(stage, dynamicfeaturetype)
	}
}

func (dynamicfeaturetype *DynamicFeatureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicFeatureTypeUpdateCallback != nil {
		var frontDynamicFeatureType *DynamicFeatureType
		if front != nil {
			frontDynamicFeatureType, _ = front.(*DynamicFeatureType)
		}
		stage.OnAfterDynamicFeatureTypeUpdateCallback.OnAfterUpdate(stage, dynamicfeaturetype, frontDynamicFeatureType)
	}
}

func (dynamicfeaturetype *DynamicFeatureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicFeatureTypeDeleteCallback != nil {
		var frontDynamicFeatureType *DynamicFeatureType
		if front != nil {
			frontDynamicFeatureType, _ = front.(*DynamicFeatureType)
		}
		stage.OnAfterDynamicFeatureTypeDeleteCallback.OnAfterDelete(stage, dynamicfeaturetype, frontDynamicFeatureType)
	}
}

func (envelopetype *EnvelopeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEnvelopeTypeCreateCallback != nil {
		stage.OnAfterEnvelopeTypeCreateCallback.OnAfterCreate(stage, envelopetype)
	}
}

func (envelopetype *EnvelopeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnvelopeTypeUpdateCallback != nil {
		var frontEnvelopeType *EnvelopeType
		if front != nil {
			frontEnvelopeType, _ = front.(*EnvelopeType)
		}
		stage.OnAfterEnvelopeTypeUpdateCallback.OnAfterUpdate(stage, envelopetype, frontEnvelopeType)
	}
}

func (envelopetype *EnvelopeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnvelopeTypeDeleteCallback != nil {
		var frontEnvelopeType *EnvelopeType
		if front != nil {
			frontEnvelopeType, _ = front.(*EnvelopeType)
		}
		stage.OnAfterEnvelopeTypeDeleteCallback.OnAfterDelete(stage, envelopetype, frontEnvelopeType)
	}
}

func (geodesicstringtype *GeodesicStringType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeodesicStringTypeCreateCallback != nil {
		stage.OnAfterGeodesicStringTypeCreateCallback.OnAfterCreate(stage, geodesicstringtype)
	}
}

func (geodesicstringtype *GeodesicStringType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodesicStringTypeUpdateCallback != nil {
		var frontGeodesicStringType *GeodesicStringType
		if front != nil {
			frontGeodesicStringType, _ = front.(*GeodesicStringType)
		}
		stage.OnAfterGeodesicStringTypeUpdateCallback.OnAfterUpdate(stage, geodesicstringtype, frontGeodesicStringType)
	}
}

func (geodesicstringtype *GeodesicStringType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodesicStringTypeDeleteCallback != nil {
		var frontGeodesicStringType *GeodesicStringType
		if front != nil {
			frontGeodesicStringType, _ = front.(*GeodesicStringType)
		}
		stage.OnAfterGeodesicStringTypeDeleteCallback.OnAfterDelete(stage, geodesicstringtype, frontGeodesicStringType)
	}
}

func (geodesictype *GeodesicType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeodesicTypeCreateCallback != nil {
		stage.OnAfterGeodesicTypeCreateCallback.OnAfterCreate(stage, geodesictype)
	}
}

func (geodesictype *GeodesicType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodesicTypeUpdateCallback != nil {
		var frontGeodesicType *GeodesicType
		if front != nil {
			frontGeodesicType, _ = front.(*GeodesicType)
		}
		stage.OnAfterGeodesicTypeUpdateCallback.OnAfterUpdate(stage, geodesictype, frontGeodesicType)
	}
}

func (geodesictype *GeodesicType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodesicTypeDeleteCallback != nil {
		var frontGeodesicType *GeodesicType
		if front != nil {
			frontGeodesicType, _ = front.(*GeodesicType)
		}
		stage.OnAfterGeodesicTypeDeleteCallback.OnAfterDelete(stage, geodesictype, frontGeodesicType)
	}
}

func (lengthtype *LengthType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLengthTypeCreateCallback != nil {
		stage.OnAfterLengthTypeCreateCallback.OnAfterCreate(stage, lengthtype)
	}
}

func (lengthtype *LengthType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLengthTypeUpdateCallback != nil {
		var frontLengthType *LengthType
		if front != nil {
			frontLengthType, _ = front.(*LengthType)
		}
		stage.OnAfterLengthTypeUpdateCallback.OnAfterUpdate(stage, lengthtype, frontLengthType)
	}
}

func (lengthtype *LengthType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLengthTypeDeleteCallback != nil {
		var frontLengthType *LengthType
		if front != nil {
			frontLengthType, _ = front.(*LengthType)
		}
		stage.OnAfterLengthTypeDeleteCallback.OnAfterDelete(stage, lengthtype, frontLengthType)
	}
}

func (linestringsegmenttype *LineStringSegmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLineStringSegmentTypeCreateCallback != nil {
		stage.OnAfterLineStringSegmentTypeCreateCallback.OnAfterCreate(stage, linestringsegmenttype)
	}
}

func (linestringsegmenttype *LineStringSegmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineStringSegmentTypeUpdateCallback != nil {
		var frontLineStringSegmentType *LineStringSegmentType
		if front != nil {
			frontLineStringSegmentType, _ = front.(*LineStringSegmentType)
		}
		stage.OnAfterLineStringSegmentTypeUpdateCallback.OnAfterUpdate(stage, linestringsegmenttype, frontLineStringSegmentType)
	}
}

func (linestringsegmenttype *LineStringSegmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineStringSegmentTypeDeleteCallback != nil {
		var frontLineStringSegmentType *LineStringSegmentType
		if front != nil {
			frontLineStringSegmentType, _ = front.(*LineStringSegmentType)
		}
		stage.OnAfterLineStringSegmentTypeDeleteCallback.OnAfterDelete(stage, linestringsegmenttype, frontLineStringSegmentType)
	}
}

func (measuretype *MeasureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureTypeCreateCallback != nil {
		stage.OnAfterMeasureTypeCreateCallback.OnAfterCreate(stage, measuretype)
	}
}

func (measuretype *MeasureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureTypeUpdateCallback != nil {
		var frontMeasureType *MeasureType
		if front != nil {
			frontMeasureType, _ = front.(*MeasureType)
		}
		stage.OnAfterMeasureTypeUpdateCallback.OnAfterUpdate(stage, measuretype, frontMeasureType)
	}
}

func (measuretype *MeasureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureTypeDeleteCallback != nil {
		var frontMeasureType *MeasureType
		if front != nil {
			frontMeasureType, _ = front.(*MeasureType)
		}
		stage.OnAfterMeasureTypeDeleteCallback.OnAfterDelete(stage, measuretype, frontMeasureType)
	}
}

func (orientablecurvetype *OrientableCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrientableCurveTypeCreateCallback != nil {
		stage.OnAfterOrientableCurveTypeCreateCallback.OnAfterCreate(stage, orientablecurvetype)
	}
}

func (orientablecurvetype *OrientableCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrientableCurveTypeUpdateCallback != nil {
		var frontOrientableCurveType *OrientableCurveType
		if front != nil {
			frontOrientableCurveType, _ = front.(*OrientableCurveType)
		}
		stage.OnAfterOrientableCurveTypeUpdateCallback.OnAfterUpdate(stage, orientablecurvetype, frontOrientableCurveType)
	}
}

func (orientablecurvetype *OrientableCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrientableCurveTypeDeleteCallback != nil {
		var frontOrientableCurveType *OrientableCurveType
		if front != nil {
			frontOrientableCurveType, _ = front.(*OrientableCurveType)
		}
		stage.OnAfterOrientableCurveTypeDeleteCallback.OnAfterDelete(stage, orientablecurvetype, frontOrientableCurveType)
	}
}

func (pointpropertytype *PointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointPropertyTypeCreateCallback != nil {
		stage.OnAfterPointPropertyTypeCreateCallback.OnAfterCreate(stage, pointpropertytype)
	}
}

func (pointpropertytype *PointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointPropertyTypeUpdateCallback != nil {
		var frontPointPropertyType *PointPropertyType
		if front != nil {
			frontPointPropertyType, _ = front.(*PointPropertyType)
		}
		stage.OnAfterPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, pointpropertytype, frontPointPropertyType)
	}
}

func (pointpropertytype *PointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointPropertyTypeDeleteCallback != nil {
		var frontPointPropertyType *PointPropertyType
		if front != nil {
			frontPointPropertyType, _ = front.(*PointPropertyType)
		}
		stage.OnAfterPointPropertyTypeDeleteCallback.OnAfterDelete(stage, pointpropertytype, frontPointPropertyType)
	}
}

func (pointtype *PointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointTypeCreateCallback != nil {
		stage.OnAfterPointTypeCreateCallback.OnAfterCreate(stage, pointtype)
	}
}

func (pointtype *PointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointTypeUpdateCallback != nil {
		var frontPointType *PointType
		if front != nil {
			frontPointType, _ = front.(*PointType)
		}
		stage.OnAfterPointTypeUpdateCallback.OnAfterUpdate(stage, pointtype, frontPointType)
	}
}

func (pointtype *PointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointTypeDeleteCallback != nil {
		var frontPointType *PointType
		if front != nil {
			frontPointType, _ = front.(*PointType)
		}
		stage.OnAfterPointTypeDeleteCallback.OnAfterDelete(stage, pointtype, frontPointType)
	}
}

func (polygonpatchtype *PolygonPatchType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolygonPatchTypeCreateCallback != nil {
		stage.OnAfterPolygonPatchTypeCreateCallback.OnAfterCreate(stage, polygonpatchtype)
	}
}

func (polygonpatchtype *PolygonPatchType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygonPatchTypeUpdateCallback != nil {
		var frontPolygonPatchType *PolygonPatchType
		if front != nil {
			frontPolygonPatchType, _ = front.(*PolygonPatchType)
		}
		stage.OnAfterPolygonPatchTypeUpdateCallback.OnAfterUpdate(stage, polygonpatchtype, frontPolygonPatchType)
	}
}

func (polygonpatchtype *PolygonPatchType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygonPatchTypeDeleteCallback != nil {
		var frontPolygonPatchType *PolygonPatchType
		if front != nil {
			frontPolygonPatchType, _ = front.(*PolygonPatchType)
		}
		stage.OnAfterPolygonPatchTypeDeleteCallback.OnAfterDelete(stage, polygonpatchtype, frontPolygonPatchType)
	}
}

func (ringtype *RingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRingTypeCreateCallback != nil {
		stage.OnAfterRingTypeCreateCallback.OnAfterCreate(stage, ringtype)
	}
}

func (ringtype *RingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRingTypeUpdateCallback != nil {
		var frontRingType *RingType
		if front != nil {
			frontRingType, _ = front.(*RingType)
		}
		stage.OnAfterRingTypeUpdateCallback.OnAfterUpdate(stage, ringtype, frontRingType)
	}
}

func (ringtype *RingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRingTypeDeleteCallback != nil {
		var frontRingType *RingType
		if front != nil {
			frontRingType, _ = front.(*RingType)
		}
		stage.OnAfterRingTypeDeleteCallback.OnAfterDelete(stage, ringtype, frontRingType)
	}
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfacePatchArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterSurfacePatchArrayPropertyTypeCreateCallback.OnAfterCreate(stage, surfacepatcharraypropertytype)
	}
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePatchArrayPropertyTypeUpdateCallback != nil {
		var frontSurfacePatchArrayPropertyType *SurfacePatchArrayPropertyType
		if front != nil {
			frontSurfacePatchArrayPropertyType, _ = front.(*SurfacePatchArrayPropertyType)
		}
		stage.OnAfterSurfacePatchArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, surfacepatcharraypropertytype, frontSurfacePatchArrayPropertyType)
	}
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePatchArrayPropertyTypeDeleteCallback != nil {
		var frontSurfacePatchArrayPropertyType *SurfacePatchArrayPropertyType
		if front != nil {
			frontSurfacePatchArrayPropertyType, _ = front.(*SurfacePatchArrayPropertyType)
		}
		stage.OnAfterSurfacePatchArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, surfacepatcharraypropertytype, frontSurfacePatchArrayPropertyType)
	}
}

func (surfacetype *SurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceTypeCreateCallback != nil {
		stage.OnAfterSurfaceTypeCreateCallback.OnAfterCreate(stage, surfacetype)
	}
}

func (surfacetype *SurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceTypeUpdateCallback != nil {
		var frontSurfaceType *SurfaceType
		if front != nil {
			frontSurfaceType, _ = front.(*SurfaceType)
		}
		stage.OnAfterSurfaceTypeUpdateCallback.OnAfterUpdate(stage, surfacetype, frontSurfaceType)
	}
}

func (surfacetype *SurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceTypeDeleteCallback != nil {
		var frontSurfaceType *SurfaceType
		if front != nil {
			frontSurfaceType, _ = front.(*SurfaceType)
		}
		stage.OnAfterSurfaceTypeDeleteCallback.OnAfterDelete(stage, surfacetype, frontSurfaceType)
	}
}

func (timeinstanttype *TimeInstantType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeInstantTypeCreateCallback != nil {
		stage.OnAfterTimeInstantTypeCreateCallback.OnAfterCreate(stage, timeinstanttype)
	}
}

func (timeinstanttype *TimeInstantType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeInstantTypeUpdateCallback != nil {
		var frontTimeInstantType *TimeInstantType
		if front != nil {
			frontTimeInstantType, _ = front.(*TimeInstantType)
		}
		stage.OnAfterTimeInstantTypeUpdateCallback.OnAfterUpdate(stage, timeinstanttype, frontTimeInstantType)
	}
}

func (timeinstanttype *TimeInstantType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeInstantTypeDeleteCallback != nil {
		var frontTimeInstantType *TimeInstantType
		if front != nil {
			frontTimeInstantType, _ = front.(*TimeInstantType)
		}
		stage.OnAfterTimeInstantTypeDeleteCallback.OnAfterDelete(stage, timeinstanttype, frontTimeInstantType)
	}
}

func (timeperiodtype *TimePeriodType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimePeriodTypeCreateCallback != nil {
		stage.OnAfterTimePeriodTypeCreateCallback.OnAfterCreate(stage, timeperiodtype)
	}
}

func (timeperiodtype *TimePeriodType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePeriodTypeUpdateCallback != nil {
		var frontTimePeriodType *TimePeriodType
		if front != nil {
			frontTimePeriodType, _ = front.(*TimePeriodType)
		}
		stage.OnAfterTimePeriodTypeUpdateCallback.OnAfterUpdate(stage, timeperiodtype, frontTimePeriodType)
	}
}

func (timeperiodtype *TimePeriodType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePeriodTypeDeleteCallback != nil {
		var frontTimePeriodType *TimePeriodType
		if front != nil {
			frontTimePeriodType, _ = front.(*TimePeriodType)
		}
		stage.OnAfterTimePeriodTypeDeleteCallback.OnAfterDelete(stage, timeperiodtype, frontTimePeriodType)
	}
}

func (timepositiontype *TimePositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimePositionTypeCreateCallback != nil {
		stage.OnAfterTimePositionTypeCreateCallback.OnAfterCreate(stage, timepositiontype)
	}
}

func (timepositiontype *TimePositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePositionTypeUpdateCallback != nil {
		var frontTimePositionType *TimePositionType
		if front != nil {
			frontTimePositionType, _ = front.(*TimePositionType)
		}
		stage.OnAfterTimePositionTypeUpdateCallback.OnAfterUpdate(stage, timepositiontype, frontTimePositionType)
	}
}

func (timepositiontype *TimePositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePositionTypeDeleteCallback != nil {
		var frontTimePositionType *TimePositionType
		if front != nil {
			frontTimePositionType, _ = front.(*TimePositionType)
		}
		stage.OnAfterTimePositionTypeDeleteCallback.OnAfterDelete(stage, timepositiontype, frontTimePositionType)
	}
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimePrimitivePropertyTypeCreateCallback != nil {
		stage.OnAfterTimePrimitivePropertyTypeCreateCallback.OnAfterCreate(stage, timeprimitivepropertytype)
	}
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePrimitivePropertyTypeUpdateCallback != nil {
		var frontTimePrimitivePropertyType *TimePrimitivePropertyType
		if front != nil {
			frontTimePrimitivePropertyType, _ = front.(*TimePrimitivePropertyType)
		}
		stage.OnAfterTimePrimitivePropertyTypeUpdateCallback.OnAfterUpdate(stage, timeprimitivepropertytype, frontTimePrimitivePropertyType)
	}
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePrimitivePropertyTypeDeleteCallback != nil {
		var frontTimePrimitivePropertyType *TimePrimitivePropertyType
		if front != nil {
			frontTimePrimitivePropertyType, _ = front.(*TimePrimitivePropertyType)
		}
		stage.OnAfterTimePrimitivePropertyTypeDeleteCallback.OnAfterDelete(stage, timeprimitivepropertytype, frontTimePrimitivePropertyType)
	}
}

