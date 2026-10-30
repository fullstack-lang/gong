// generated code - do not edit
package content

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
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractMD_ContentInformation_TypeCreateCallback != nil {
		stage.OnAfterAbstractMD_ContentInformation_TypeCreateCallback.OnAfterCreate(stage, abstractmd_contentinformation_type)
	}
}

func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMD_ContentInformation_TypeUpdateCallback != nil {
		var frontAbstractMD_ContentInformation_Type *AbstractMD_ContentInformation_Type
		if front != nil {
			frontAbstractMD_ContentInformation_Type, _ = front.(*AbstractMD_ContentInformation_Type)
		}
		stage.OnAfterAbstractMD_ContentInformation_TypeUpdateCallback.OnAfterUpdate(stage, abstractmd_contentinformation_type, frontAbstractMD_ContentInformation_Type)
	}
}

func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMD_ContentInformation_TypeDeleteCallback != nil {
		var frontAbstractMD_ContentInformation_Type *AbstractMD_ContentInformation_Type
		if front != nil {
			frontAbstractMD_ContentInformation_Type, _ = front.(*AbstractMD_ContentInformation_Type)
		}
		stage.OnAfterAbstractMD_ContentInformation_TypeDeleteCallback.OnAfterDelete(stage, abstractmd_contentinformation_type, frontAbstractMD_ContentInformation_Type)
	}
}

func (md_band_propertytype *MD_Band_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Band_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Band_PropertyTypeCreateCallback.OnAfterCreate(stage, md_band_propertytype)
	}
}

func (md_band_propertytype *MD_Band_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Band_PropertyTypeUpdateCallback != nil {
		var frontMD_Band_PropertyType *MD_Band_PropertyType
		if front != nil {
			frontMD_Band_PropertyType, _ = front.(*MD_Band_PropertyType)
		}
		stage.OnAfterMD_Band_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_band_propertytype, frontMD_Band_PropertyType)
	}
}

func (md_band_propertytype *MD_Band_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Band_PropertyTypeDeleteCallback != nil {
		var frontMD_Band_PropertyType *MD_Band_PropertyType
		if front != nil {
			frontMD_Band_PropertyType, _ = front.(*MD_Band_PropertyType)
		}
		stage.OnAfterMD_Band_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_band_propertytype, frontMD_Band_PropertyType)
	}
}

func (md_band_type *MD_Band_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Band_TypeCreateCallback != nil {
		stage.OnAfterMD_Band_TypeCreateCallback.OnAfterCreate(stage, md_band_type)
	}
}

func (md_band_type *MD_Band_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Band_TypeUpdateCallback != nil {
		var frontMD_Band_Type *MD_Band_Type
		if front != nil {
			frontMD_Band_Type, _ = front.(*MD_Band_Type)
		}
		stage.OnAfterMD_Band_TypeUpdateCallback.OnAfterUpdate(stage, md_band_type, frontMD_Band_Type)
	}
}

func (md_band_type *MD_Band_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Band_TypeDeleteCallback != nil {
		var frontMD_Band_Type *MD_Band_Type
		if front != nil {
			frontMD_Band_Type, _ = front.(*MD_Band_Type)
		}
		stage.OnAfterMD_Band_TypeDeleteCallback.OnAfterDelete(stage, md_band_type, frontMD_Band_Type)
	}
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ContentInformation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ContentInformation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_contentinformation_propertytype)
	}
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ContentInformation_PropertyTypeUpdateCallback != nil {
		var frontMD_ContentInformation_PropertyType *MD_ContentInformation_PropertyType
		if front != nil {
			frontMD_ContentInformation_PropertyType, _ = front.(*MD_ContentInformation_PropertyType)
		}
		stage.OnAfterMD_ContentInformation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_contentinformation_propertytype, frontMD_ContentInformation_PropertyType)
	}
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ContentInformation_PropertyTypeDeleteCallback != nil {
		var frontMD_ContentInformation_PropertyType *MD_ContentInformation_PropertyType
		if front != nil {
			frontMD_ContentInformation_PropertyType, _ = front.(*MD_ContentInformation_PropertyType)
		}
		stage.OnAfterMD_ContentInformation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_contentinformation_propertytype, frontMD_ContentInformation_PropertyType)
	}
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_CoverageContentTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_CoverageContentTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_coveragecontenttypecode_propertytype)
	}
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CoverageContentTypeCode_PropertyTypeUpdateCallback != nil {
		var frontMD_CoverageContentTypeCode_PropertyType *MD_CoverageContentTypeCode_PropertyType
		if front != nil {
			frontMD_CoverageContentTypeCode_PropertyType, _ = front.(*MD_CoverageContentTypeCode_PropertyType)
		}
		stage.OnAfterMD_CoverageContentTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_coveragecontenttypecode_propertytype, frontMD_CoverageContentTypeCode_PropertyType)
	}
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CoverageContentTypeCode_PropertyTypeDeleteCallback != nil {
		var frontMD_CoverageContentTypeCode_PropertyType *MD_CoverageContentTypeCode_PropertyType
		if front != nil {
			frontMD_CoverageContentTypeCode_PropertyType, _ = front.(*MD_CoverageContentTypeCode_PropertyType)
		}
		stage.OnAfterMD_CoverageContentTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_coveragecontenttypecode_propertytype, frontMD_CoverageContentTypeCode_PropertyType)
	}
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_CoverageDescription_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_CoverageDescription_PropertyTypeCreateCallback.OnAfterCreate(stage, md_coveragedescription_propertytype)
	}
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CoverageDescription_PropertyTypeUpdateCallback != nil {
		var frontMD_CoverageDescription_PropertyType *MD_CoverageDescription_PropertyType
		if front != nil {
			frontMD_CoverageDescription_PropertyType, _ = front.(*MD_CoverageDescription_PropertyType)
		}
		stage.OnAfterMD_CoverageDescription_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_coveragedescription_propertytype, frontMD_CoverageDescription_PropertyType)
	}
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CoverageDescription_PropertyTypeDeleteCallback != nil {
		var frontMD_CoverageDescription_PropertyType *MD_CoverageDescription_PropertyType
		if front != nil {
			frontMD_CoverageDescription_PropertyType, _ = front.(*MD_CoverageDescription_PropertyType)
		}
		stage.OnAfterMD_CoverageDescription_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_coveragedescription_propertytype, frontMD_CoverageDescription_PropertyType)
	}
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_CoverageDescription_TypeCreateCallback != nil {
		stage.OnAfterMD_CoverageDescription_TypeCreateCallback.OnAfterCreate(stage, md_coveragedescription_type)
	}
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CoverageDescription_TypeUpdateCallback != nil {
		var frontMD_CoverageDescription_Type *MD_CoverageDescription_Type
		if front != nil {
			frontMD_CoverageDescription_Type, _ = front.(*MD_CoverageDescription_Type)
		}
		stage.OnAfterMD_CoverageDescription_TypeUpdateCallback.OnAfterUpdate(stage, md_coveragedescription_type, frontMD_CoverageDescription_Type)
	}
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CoverageDescription_TypeDeleteCallback != nil {
		var frontMD_CoverageDescription_Type *MD_CoverageDescription_Type
		if front != nil {
			frontMD_CoverageDescription_Type, _ = front.(*MD_CoverageDescription_Type)
		}
		stage.OnAfterMD_CoverageDescription_TypeDeleteCallback.OnAfterDelete(stage, md_coveragedescription_type, frontMD_CoverageDescription_Type)
	}
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_FeatureCatalogueDescription_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_FeatureCatalogueDescription_PropertyTypeCreateCallback.OnAfterCreate(stage, md_featurecataloguedescription_propertytype)
	}
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_FeatureCatalogueDescription_PropertyTypeUpdateCallback != nil {
		var frontMD_FeatureCatalogueDescription_PropertyType *MD_FeatureCatalogueDescription_PropertyType
		if front != nil {
			frontMD_FeatureCatalogueDescription_PropertyType, _ = front.(*MD_FeatureCatalogueDescription_PropertyType)
		}
		stage.OnAfterMD_FeatureCatalogueDescription_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_featurecataloguedescription_propertytype, frontMD_FeatureCatalogueDescription_PropertyType)
	}
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_FeatureCatalogueDescription_PropertyTypeDeleteCallback != nil {
		var frontMD_FeatureCatalogueDescription_PropertyType *MD_FeatureCatalogueDescription_PropertyType
		if front != nil {
			frontMD_FeatureCatalogueDescription_PropertyType, _ = front.(*MD_FeatureCatalogueDescription_PropertyType)
		}
		stage.OnAfterMD_FeatureCatalogueDescription_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_featurecataloguedescription_propertytype, frontMD_FeatureCatalogueDescription_PropertyType)
	}
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_FeatureCatalogueDescription_TypeCreateCallback != nil {
		stage.OnAfterMD_FeatureCatalogueDescription_TypeCreateCallback.OnAfterCreate(stage, md_featurecataloguedescription_type)
	}
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_FeatureCatalogueDescription_TypeUpdateCallback != nil {
		var frontMD_FeatureCatalogueDescription_Type *MD_FeatureCatalogueDescription_Type
		if front != nil {
			frontMD_FeatureCatalogueDescription_Type, _ = front.(*MD_FeatureCatalogueDescription_Type)
		}
		stage.OnAfterMD_FeatureCatalogueDescription_TypeUpdateCallback.OnAfterUpdate(stage, md_featurecataloguedescription_type, frontMD_FeatureCatalogueDescription_Type)
	}
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_FeatureCatalogueDescription_TypeDeleteCallback != nil {
		var frontMD_FeatureCatalogueDescription_Type *MD_FeatureCatalogueDescription_Type
		if front != nil {
			frontMD_FeatureCatalogueDescription_Type, _ = front.(*MD_FeatureCatalogueDescription_Type)
		}
		stage.OnAfterMD_FeatureCatalogueDescription_TypeDeleteCallback.OnAfterDelete(stage, md_featurecataloguedescription_type, frontMD_FeatureCatalogueDescription_Type)
	}
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ImageDescription_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ImageDescription_PropertyTypeCreateCallback.OnAfterCreate(stage, md_imagedescription_propertytype)
	}
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ImageDescription_PropertyTypeUpdateCallback != nil {
		var frontMD_ImageDescription_PropertyType *MD_ImageDescription_PropertyType
		if front != nil {
			frontMD_ImageDescription_PropertyType, _ = front.(*MD_ImageDescription_PropertyType)
		}
		stage.OnAfterMD_ImageDescription_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_imagedescription_propertytype, frontMD_ImageDescription_PropertyType)
	}
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ImageDescription_PropertyTypeDeleteCallback != nil {
		var frontMD_ImageDescription_PropertyType *MD_ImageDescription_PropertyType
		if front != nil {
			frontMD_ImageDescription_PropertyType, _ = front.(*MD_ImageDescription_PropertyType)
		}
		stage.OnAfterMD_ImageDescription_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_imagedescription_propertytype, frontMD_ImageDescription_PropertyType)
	}
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ImageDescription_TypeCreateCallback != nil {
		stage.OnAfterMD_ImageDescription_TypeCreateCallback.OnAfterCreate(stage, md_imagedescription_type)
	}
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ImageDescription_TypeUpdateCallback != nil {
		var frontMD_ImageDescription_Type *MD_ImageDescription_Type
		if front != nil {
			frontMD_ImageDescription_Type, _ = front.(*MD_ImageDescription_Type)
		}
		stage.OnAfterMD_ImageDescription_TypeUpdateCallback.OnAfterUpdate(stage, md_imagedescription_type, frontMD_ImageDescription_Type)
	}
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ImageDescription_TypeDeleteCallback != nil {
		var frontMD_ImageDescription_Type *MD_ImageDescription_Type
		if front != nil {
			frontMD_ImageDescription_Type, _ = front.(*MD_ImageDescription_Type)
		}
		stage.OnAfterMD_ImageDescription_TypeDeleteCallback.OnAfterDelete(stage, md_imagedescription_type, frontMD_ImageDescription_Type)
	}
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ImagingConditionCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ImagingConditionCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_imagingconditioncode_propertytype)
	}
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ImagingConditionCode_PropertyTypeUpdateCallback != nil {
		var frontMD_ImagingConditionCode_PropertyType *MD_ImagingConditionCode_PropertyType
		if front != nil {
			frontMD_ImagingConditionCode_PropertyType, _ = front.(*MD_ImagingConditionCode_PropertyType)
		}
		stage.OnAfterMD_ImagingConditionCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_imagingconditioncode_propertytype, frontMD_ImagingConditionCode_PropertyType)
	}
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ImagingConditionCode_PropertyTypeDeleteCallback != nil {
		var frontMD_ImagingConditionCode_PropertyType *MD_ImagingConditionCode_PropertyType
		if front != nil {
			frontMD_ImagingConditionCode_PropertyType, _ = front.(*MD_ImagingConditionCode_PropertyType)
		}
		stage.OnAfterMD_ImagingConditionCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_imagingconditioncode_propertytype, frontMD_ImagingConditionCode_PropertyType)
	}
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_RangeDimension_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_RangeDimension_PropertyTypeCreateCallback.OnAfterCreate(stage, md_rangedimension_propertytype)
	}
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RangeDimension_PropertyTypeUpdateCallback != nil {
		var frontMD_RangeDimension_PropertyType *MD_RangeDimension_PropertyType
		if front != nil {
			frontMD_RangeDimension_PropertyType, _ = front.(*MD_RangeDimension_PropertyType)
		}
		stage.OnAfterMD_RangeDimension_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_rangedimension_propertytype, frontMD_RangeDimension_PropertyType)
	}
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RangeDimension_PropertyTypeDeleteCallback != nil {
		var frontMD_RangeDimension_PropertyType *MD_RangeDimension_PropertyType
		if front != nil {
			frontMD_RangeDimension_PropertyType, _ = front.(*MD_RangeDimension_PropertyType)
		}
		stage.OnAfterMD_RangeDimension_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_rangedimension_propertytype, frontMD_RangeDimension_PropertyType)
	}
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_RangeDimension_TypeCreateCallback != nil {
		stage.OnAfterMD_RangeDimension_TypeCreateCallback.OnAfterCreate(stage, md_rangedimension_type)
	}
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RangeDimension_TypeUpdateCallback != nil {
		var frontMD_RangeDimension_Type *MD_RangeDimension_Type
		if front != nil {
			frontMD_RangeDimension_Type, _ = front.(*MD_RangeDimension_Type)
		}
		stage.OnAfterMD_RangeDimension_TypeUpdateCallback.OnAfterUpdate(stage, md_rangedimension_type, frontMD_RangeDimension_Type)
	}
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RangeDimension_TypeDeleteCallback != nil {
		var frontMD_RangeDimension_Type *MD_RangeDimension_Type
		if front != nil {
			frontMD_RangeDimension_Type, _ = front.(*MD_RangeDimension_Type)
		}
		stage.OnAfterMD_RangeDimension_TypeDeleteCallback.OnAfterDelete(stage, md_rangedimension_type, frontMD_RangeDimension_Type)
	}
}

