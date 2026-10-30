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

