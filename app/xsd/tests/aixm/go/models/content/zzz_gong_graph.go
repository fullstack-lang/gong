// generated code - do not edit
package content

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
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractMD_ContentInformation_Types[abstractmd_contentinformation_type]
	return ok
}

func (md_band_propertytype *MD_Band_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Band_PropertyTypes[md_band_propertytype]
	return ok
}

func (md_band_type *MD_Band_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_Band_Types[md_band_type]
	return ok
}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_ContentInformation_PropertyTypes[md_contentinformation_propertytype]
	return ok
}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_CoverageContentTypeCode_PropertyTypes[md_coveragecontenttypecode_propertytype]
	return ok
}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_CoverageDescription_PropertyTypes[md_coveragedescription_propertytype]
	return ok
}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_CoverageDescription_Types[md_coveragedescription_type]
	return ok
}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_FeatureCatalogueDescription_PropertyTypes[md_featurecataloguedescription_propertytype]
	return ok
}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_FeatureCatalogueDescription_Types[md_featurecataloguedescription_type]
	return ok
}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_ImageDescription_PropertyTypes[md_imagedescription_propertytype]
	return ok
}

func (md_imagedescription_type *MD_ImageDescription_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_ImageDescription_Types[md_imagedescription_type]
	return ok
}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_ImagingConditionCode_PropertyTypes[md_imagingconditioncode_propertytype]
	return ok
}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_RangeDimension_PropertyTypes[md_rangedimension_propertytype]
	return ok
}

func (md_rangedimension_type *MD_RangeDimension_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_RangeDimension_Types[md_rangedimension_type]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractmd_contentinformation_type) {
		return
	}

	abstractmd_contentinformation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_band_propertytype *MD_Band_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_band_propertytype) {
		return
	}

	md_band_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_band_propertytype.MD_Band != nil {
		stage.StageBranch(md_band_propertytype.MD_Band)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_band_type *MD_Band_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_band_type) {
		return
	}

	md_band_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_contentinformation_propertytype) {
		return
	}

	md_contentinformation_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_contentinformation_propertytype.MD_FeatureCatalogueDescription != nil {
		stage.StageBranch(md_contentinformation_propertytype.MD_FeatureCatalogueDescription)
	}
	if md_contentinformation_propertytype.MD_CoverageDescription != nil {
		stage.StageBranch(md_contentinformation_propertytype.MD_CoverageDescription)
	}
	if md_contentinformation_propertytype.MD_ImageDescription != nil {
		stage.StageBranch(md_contentinformation_propertytype.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_coveragecontenttypecode_propertytype) {
		return
	}

	md_coveragecontenttypecode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_coveragedescription_propertytype) {
		return
	}

	md_coveragedescription_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_coveragedescription_propertytype.MD_CoverageDescription != nil {
		stage.StageBranch(md_coveragedescription_propertytype.MD_CoverageDescription)
	}
	if md_coveragedescription_propertytype.MD_ImageDescription != nil {
		stage.StageBranch(md_coveragedescription_propertytype.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_coveragedescription_type) {
		return
	}

	md_coveragedescription_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_coveragedescription_type.ContentType != nil {
		stage.StageBranch(md_coveragedescription_type.ContentType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_rangedimension_propertytype := range md_coveragedescription_type.Dimension {
		stage.StageBranch(_md_rangedimension_propertytype)
	}

}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_featurecataloguedescription_propertytype) {
		return
	}

	md_featurecataloguedescription_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription != nil {
		stage.StageBranch(md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_featurecataloguedescription_type) {
		return
	}

	md_featurecataloguedescription_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_imagedescription_propertytype) {
		return
	}

	md_imagedescription_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_imagedescription_propertytype.MD_ImageDescription != nil {
		stage.StageBranch(md_imagedescription_propertytype.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_imagedescription_type *MD_ImageDescription_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_imagedescription_type) {
		return
	}

	md_imagedescription_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_imagedescription_type.ContentType != nil {
		stage.StageBranch(md_imagedescription_type.ContentType)
	}
	if md_imagedescription_type.ImagingCondition != nil {
		stage.StageBranch(md_imagedescription_type.ImagingCondition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_rangedimension_propertytype := range md_imagedescription_type.Dimension {
		stage.StageBranch(_md_rangedimension_propertytype)
	}

}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_imagingconditioncode_propertytype) {
		return
	}

	md_imagingconditioncode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_rangedimension_propertytype) {
		return
	}

	md_rangedimension_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_rangedimension_propertytype.MD_RangeDimension != nil {
		stage.StageBranch(md_rangedimension_propertytype.MD_RangeDimension)
	}
	if md_rangedimension_propertytype.MD_Band != nil {
		stage.StageBranch(md_rangedimension_propertytype.MD_Band)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_rangedimension_type *MD_RangeDimension_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_rangedimension_type) {
		return
	}

	md_rangedimension_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
	case *AbstractMD_ContentInformation_Type:
		toT := GongCopyBranchAbstractMD_ContentInformation_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Band_PropertyType:
		toT := GongCopyBranchMD_Band_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_Band_Type:
		toT := GongCopyBranchMD_Band_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_ContentInformation_PropertyType:
		toT := GongCopyBranchMD_ContentInformation_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_CoverageContentTypeCode_PropertyType:
		toT := GongCopyBranchMD_CoverageContentTypeCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_CoverageDescription_PropertyType:
		toT := GongCopyBranchMD_CoverageDescription_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_CoverageDescription_Type:
		toT := GongCopyBranchMD_CoverageDescription_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_FeatureCatalogueDescription_PropertyType:
		toT := GongCopyBranchMD_FeatureCatalogueDescription_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_FeatureCatalogueDescription_Type:
		toT := GongCopyBranchMD_FeatureCatalogueDescription_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_ImageDescription_PropertyType:
		toT := GongCopyBranchMD_ImageDescription_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_ImageDescription_Type:
		toT := GongCopyBranchMD_ImageDescription_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_ImagingConditionCode_PropertyType:
		toT := GongCopyBranchMD_ImagingConditionCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_RangeDimension_PropertyType:
		toT := GongCopyBranchMD_RangeDimension_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_RangeDimension_Type:
		toT := GongCopyBranchMD_RangeDimension_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAbstractMD_ContentInformation_Type(mapOrigCopy map[any]any, abstractmd_contentinformation_typeFrom *AbstractMD_ContentInformation_Type) (abstractmd_contentinformation_typeTo *AbstractMD_ContentInformation_Type) {
	var alreadyCopied bool
	abstractmd_contentinformation_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractmd_contentinformation_typeFrom)
	if alreadyCopied {
		return
	}
	abstractmd_contentinformation_typeFrom.GongCopyBasicFields(abstractmd_contentinformation_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_Band_PropertyType(mapOrigCopy map[any]any, md_band_propertytypeFrom *MD_Band_PropertyType) (md_band_propertytypeTo *MD_Band_PropertyType) {
	var alreadyCopied bool
	md_band_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_band_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_band_propertytypeFrom.GongCopyBasicFields(md_band_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_band_propertytypeFrom.MD_Band != nil {
		md_band_propertytypeTo.MD_Band = GongCopyBranchMD_Band_Type(mapOrigCopy, md_band_propertytypeFrom.MD_Band)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_Band_Type(mapOrigCopy map[any]any, md_band_typeFrom *MD_Band_Type) (md_band_typeTo *MD_Band_Type) {
	var alreadyCopied bool
	md_band_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_band_typeFrom)
	if alreadyCopied {
		return
	}
	md_band_typeFrom.GongCopyBasicFields(md_band_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_ContentInformation_PropertyType(mapOrigCopy map[any]any, md_contentinformation_propertytypeFrom *MD_ContentInformation_PropertyType) (md_contentinformation_propertytypeTo *MD_ContentInformation_PropertyType) {
	var alreadyCopied bool
	md_contentinformation_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_contentinformation_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_contentinformation_propertytypeFrom.GongCopyBasicFields(md_contentinformation_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_contentinformation_propertytypeFrom.MD_FeatureCatalogueDescription != nil {
		md_contentinformation_propertytypeTo.MD_FeatureCatalogueDescription = GongCopyBranchMD_FeatureCatalogueDescription_Type(mapOrigCopy, md_contentinformation_propertytypeFrom.MD_FeatureCatalogueDescription)
	}
	if md_contentinformation_propertytypeFrom.MD_CoverageDescription != nil {
		md_contentinformation_propertytypeTo.MD_CoverageDescription = GongCopyBranchMD_CoverageDescription_Type(mapOrigCopy, md_contentinformation_propertytypeFrom.MD_CoverageDescription)
	}
	if md_contentinformation_propertytypeFrom.MD_ImageDescription != nil {
		md_contentinformation_propertytypeTo.MD_ImageDescription = GongCopyBranchMD_ImageDescription_Type(mapOrigCopy, md_contentinformation_propertytypeFrom.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_CoverageContentTypeCode_PropertyType(mapOrigCopy map[any]any, md_coveragecontenttypecode_propertytypeFrom *MD_CoverageContentTypeCode_PropertyType) (md_coveragecontenttypecode_propertytypeTo *MD_CoverageContentTypeCode_PropertyType) {
	var alreadyCopied bool
	md_coveragecontenttypecode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_coveragecontenttypecode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_coveragecontenttypecode_propertytypeFrom.GongCopyBasicFields(md_coveragecontenttypecode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_CoverageDescription_PropertyType(mapOrigCopy map[any]any, md_coveragedescription_propertytypeFrom *MD_CoverageDescription_PropertyType) (md_coveragedescription_propertytypeTo *MD_CoverageDescription_PropertyType) {
	var alreadyCopied bool
	md_coveragedescription_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_coveragedescription_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_coveragedescription_propertytypeFrom.GongCopyBasicFields(md_coveragedescription_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_coveragedescription_propertytypeFrom.MD_CoverageDescription != nil {
		md_coveragedescription_propertytypeTo.MD_CoverageDescription = GongCopyBranchMD_CoverageDescription_Type(mapOrigCopy, md_coveragedescription_propertytypeFrom.MD_CoverageDescription)
	}
	if md_coveragedescription_propertytypeFrom.MD_ImageDescription != nil {
		md_coveragedescription_propertytypeTo.MD_ImageDescription = GongCopyBranchMD_ImageDescription_Type(mapOrigCopy, md_coveragedescription_propertytypeFrom.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_CoverageDescription_Type(mapOrigCopy map[any]any, md_coveragedescription_typeFrom *MD_CoverageDescription_Type) (md_coveragedescription_typeTo *MD_CoverageDescription_Type) {
	var alreadyCopied bool
	md_coveragedescription_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_coveragedescription_typeFrom)
	if alreadyCopied {
		return
	}
	md_coveragedescription_typeFrom.GongCopyBasicFields(md_coveragedescription_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_coveragedescription_typeFrom.ContentType != nil {
		md_coveragedescription_typeTo.ContentType = GongCopyBranchMD_CoverageContentTypeCode_PropertyType(mapOrigCopy, md_coveragedescription_typeFrom.ContentType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_rangedimension_propertytype := range md_coveragedescription_typeFrom.Dimension {
		md_coveragedescription_typeTo.Dimension = append(md_coveragedescription_typeTo.Dimension, GongCopyBranchMD_RangeDimension_PropertyType(mapOrigCopy, _md_rangedimension_propertytype))
	}

	return
}

func GongCopyBranchMD_FeatureCatalogueDescription_PropertyType(mapOrigCopy map[any]any, md_featurecataloguedescription_propertytypeFrom *MD_FeatureCatalogueDescription_PropertyType) (md_featurecataloguedescription_propertytypeTo *MD_FeatureCatalogueDescription_PropertyType) {
	var alreadyCopied bool
	md_featurecataloguedescription_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_featurecataloguedescription_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_featurecataloguedescription_propertytypeFrom.GongCopyBasicFields(md_featurecataloguedescription_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_featurecataloguedescription_propertytypeFrom.MD_FeatureCatalogueDescription != nil {
		md_featurecataloguedescription_propertytypeTo.MD_FeatureCatalogueDescription = GongCopyBranchMD_FeatureCatalogueDescription_Type(mapOrigCopy, md_featurecataloguedescription_propertytypeFrom.MD_FeatureCatalogueDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_FeatureCatalogueDescription_Type(mapOrigCopy map[any]any, md_featurecataloguedescription_typeFrom *MD_FeatureCatalogueDescription_Type) (md_featurecataloguedescription_typeTo *MD_FeatureCatalogueDescription_Type) {
	var alreadyCopied bool
	md_featurecataloguedescription_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_featurecataloguedescription_typeFrom)
	if alreadyCopied {
		return
	}
	md_featurecataloguedescription_typeFrom.GongCopyBasicFields(md_featurecataloguedescription_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_ImageDescription_PropertyType(mapOrigCopy map[any]any, md_imagedescription_propertytypeFrom *MD_ImageDescription_PropertyType) (md_imagedescription_propertytypeTo *MD_ImageDescription_PropertyType) {
	var alreadyCopied bool
	md_imagedescription_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_imagedescription_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_imagedescription_propertytypeFrom.GongCopyBasicFields(md_imagedescription_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_imagedescription_propertytypeFrom.MD_ImageDescription != nil {
		md_imagedescription_propertytypeTo.MD_ImageDescription = GongCopyBranchMD_ImageDescription_Type(mapOrigCopy, md_imagedescription_propertytypeFrom.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_ImageDescription_Type(mapOrigCopy map[any]any, md_imagedescription_typeFrom *MD_ImageDescription_Type) (md_imagedescription_typeTo *MD_ImageDescription_Type) {
	var alreadyCopied bool
	md_imagedescription_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_imagedescription_typeFrom)
	if alreadyCopied {
		return
	}
	md_imagedescription_typeFrom.GongCopyBasicFields(md_imagedescription_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_imagedescription_typeFrom.ContentType != nil {
		md_imagedescription_typeTo.ContentType = GongCopyBranchMD_CoverageContentTypeCode_PropertyType(mapOrigCopy, md_imagedescription_typeFrom.ContentType)
	}
	if md_imagedescription_typeFrom.ImagingCondition != nil {
		md_imagedescription_typeTo.ImagingCondition = GongCopyBranchMD_ImagingConditionCode_PropertyType(mapOrigCopy, md_imagedescription_typeFrom.ImagingCondition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_rangedimension_propertytype := range md_imagedescription_typeFrom.Dimension {
		md_imagedescription_typeTo.Dimension = append(md_imagedescription_typeTo.Dimension, GongCopyBranchMD_RangeDimension_PropertyType(mapOrigCopy, _md_rangedimension_propertytype))
	}

	return
}

func GongCopyBranchMD_ImagingConditionCode_PropertyType(mapOrigCopy map[any]any, md_imagingconditioncode_propertytypeFrom *MD_ImagingConditionCode_PropertyType) (md_imagingconditioncode_propertytypeTo *MD_ImagingConditionCode_PropertyType) {
	var alreadyCopied bool
	md_imagingconditioncode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_imagingconditioncode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_imagingconditioncode_propertytypeFrom.GongCopyBasicFields(md_imagingconditioncode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_RangeDimension_PropertyType(mapOrigCopy map[any]any, md_rangedimension_propertytypeFrom *MD_RangeDimension_PropertyType) (md_rangedimension_propertytypeTo *MD_RangeDimension_PropertyType) {
	var alreadyCopied bool
	md_rangedimension_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_rangedimension_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_rangedimension_propertytypeFrom.GongCopyBasicFields(md_rangedimension_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_rangedimension_propertytypeFrom.MD_RangeDimension != nil {
		md_rangedimension_propertytypeTo.MD_RangeDimension = GongCopyBranchMD_RangeDimension_Type(mapOrigCopy, md_rangedimension_propertytypeFrom.MD_RangeDimension)
	}
	if md_rangedimension_propertytypeFrom.MD_Band != nil {
		md_rangedimension_propertytypeTo.MD_Band = GongCopyBranchMD_Band_Type(mapOrigCopy, md_rangedimension_propertytypeFrom.MD_Band)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_RangeDimension_Type(mapOrigCopy map[any]any, md_rangedimension_typeFrom *MD_RangeDimension_Type) (md_rangedimension_typeTo *MD_RangeDimension_Type) {
	var alreadyCopied bool
	md_rangedimension_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_rangedimension_typeFrom)
	if alreadyCopied {
		return
	}
	md_rangedimension_typeFrom.GongCopyBasicFields(md_rangedimension_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractmd_contentinformation_type) {
		return
	}

	abstractmd_contentinformation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_band_propertytype *MD_Band_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_band_propertytype) {
		return
	}

	md_band_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_band_propertytype.MD_Band != nil {
		stage.UnstageBranch(md_band_propertytype.MD_Band)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_band_type *MD_Band_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_band_type) {
		return
	}

	md_band_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_contentinformation_propertytype) {
		return
	}

	md_contentinformation_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_contentinformation_propertytype.MD_FeatureCatalogueDescription != nil {
		stage.UnstageBranch(md_contentinformation_propertytype.MD_FeatureCatalogueDescription)
	}
	if md_contentinformation_propertytype.MD_CoverageDescription != nil {
		stage.UnstageBranch(md_contentinformation_propertytype.MD_CoverageDescription)
	}
	if md_contentinformation_propertytype.MD_ImageDescription != nil {
		stage.UnstageBranch(md_contentinformation_propertytype.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_coveragecontenttypecode_propertytype) {
		return
	}

	md_coveragecontenttypecode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_coveragedescription_propertytype) {
		return
	}

	md_coveragedescription_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_coveragedescription_propertytype.MD_CoverageDescription != nil {
		stage.UnstageBranch(md_coveragedescription_propertytype.MD_CoverageDescription)
	}
	if md_coveragedescription_propertytype.MD_ImageDescription != nil {
		stage.UnstageBranch(md_coveragedescription_propertytype.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_coveragedescription_type *MD_CoverageDescription_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_coveragedescription_type) {
		return
	}

	md_coveragedescription_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_coveragedescription_type.ContentType != nil {
		stage.UnstageBranch(md_coveragedescription_type.ContentType)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_rangedimension_propertytype := range md_coveragedescription_type.Dimension {
		stage.UnstageBranch(_md_rangedimension_propertytype)
	}

}

func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_featurecataloguedescription_propertytype) {
		return
	}

	md_featurecataloguedescription_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription != nil {
		stage.UnstageBranch(md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_featurecataloguedescription_type) {
		return
	}

	md_featurecataloguedescription_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_imagedescription_propertytype) {
		return
	}

	md_imagedescription_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_imagedescription_propertytype.MD_ImageDescription != nil {
		stage.UnstageBranch(md_imagedescription_propertytype.MD_ImageDescription)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_imagedescription_type *MD_ImageDescription_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_imagedescription_type) {
		return
	}

	md_imagedescription_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_imagedescription_type.ContentType != nil {
		stage.UnstageBranch(md_imagedescription_type.ContentType)
	}
	if md_imagedescription_type.ImagingCondition != nil {
		stage.UnstageBranch(md_imagedescription_type.ImagingCondition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_rangedimension_propertytype := range md_imagedescription_type.Dimension {
		stage.UnstageBranch(_md_rangedimension_propertytype)
	}

}

func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_imagingconditioncode_propertytype) {
		return
	}

	md_imagingconditioncode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_rangedimension_propertytype) {
		return
	}

	md_rangedimension_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_rangedimension_propertytype.MD_RangeDimension != nil {
		stage.UnstageBranch(md_rangedimension_propertytype.MD_RangeDimension)
	}
	if md_rangedimension_propertytype.MD_Band != nil {
		stage.UnstageBranch(md_rangedimension_propertytype.MD_Band)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_rangedimension_type *MD_RangeDimension_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_rangedimension_type) {
		return
	}

	md_rangedimension_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AbstractMD_ContentInformation_Type) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractMD_ContentInformation_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_Band_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Band_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_Band, stage.MD_Band_Types_reference, instance.MD_Band)
	// insertion point for slice of pointers field
}

func (reference *MD_Band_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_Band_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_ContentInformation_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_ContentInformation_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_FeatureCatalogueDescription, stage.MD_FeatureCatalogueDescription_Types_reference, instance.MD_FeatureCatalogueDescription)
	__gong__reconstructPointer(&reference.MD_CoverageDescription, stage.MD_CoverageDescription_Types_reference, instance.MD_CoverageDescription)
	__gong__reconstructPointer(&reference.MD_ImageDescription, stage.MD_ImageDescription_Types_reference, instance.MD_ImageDescription)
	// insertion point for slice of pointers field
}

func (reference *MD_CoverageContentTypeCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_CoverageContentTypeCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_CoverageDescription_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_CoverageDescription_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_CoverageDescription, stage.MD_CoverageDescription_Types_reference, instance.MD_CoverageDescription)
	__gong__reconstructPointer(&reference.MD_ImageDescription, stage.MD_ImageDescription_Types_reference, instance.MD_ImageDescription)
	// insertion point for slice of pointers field
}

func (reference *MD_CoverageDescription_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_CoverageDescription_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ContentType, stage.MD_CoverageContentTypeCode_PropertyTypes_reference, instance.ContentType)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Dimension, stage.MD_RangeDimension_PropertyTypes_reference, instance.Dimension)
}

func (reference *MD_FeatureCatalogueDescription_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_FeatureCatalogueDescription_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_FeatureCatalogueDescription, stage.MD_FeatureCatalogueDescription_Types_reference, instance.MD_FeatureCatalogueDescription)
	// insertion point for slice of pointers field
}

func (reference *MD_FeatureCatalogueDescription_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_FeatureCatalogueDescription_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_ImageDescription_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_ImageDescription_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_ImageDescription, stage.MD_ImageDescription_Types_reference, instance.MD_ImageDescription)
	// insertion point for slice of pointers field
}

func (reference *MD_ImageDescription_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_ImageDescription_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ContentType, stage.MD_CoverageContentTypeCode_PropertyTypes_reference, instance.ContentType)
	__gong__reconstructPointer(&reference.ImagingCondition, stage.MD_ImagingConditionCode_PropertyTypes_reference, instance.ImagingCondition)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Dimension, stage.MD_RangeDimension_PropertyTypes_reference, instance.Dimension)
}

func (reference *MD_ImagingConditionCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_ImagingConditionCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_RangeDimension_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_RangeDimension_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_RangeDimension, stage.MD_RangeDimension_Types_reference, instance.MD_RangeDimension)
	__gong__reconstructPointer(&reference.MD_Band, stage.MD_Band_Types_reference, instance.MD_Band)
	// insertion point for slice of pointers field
}

func (reference *MD_RangeDimension_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_RangeDimension_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AbstractMD_ContentInformation_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_Band_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_Band, stage.MD_Band_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_Band_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_ContentInformation_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_FeatureCatalogueDescription, stage.MD_FeatureCatalogueDescription_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_CoverageDescription, stage.MD_CoverageDescription_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_ImageDescription, stage.MD_ImageDescription_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_CoverageContentTypeCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_CoverageDescription_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_CoverageDescription, stage.MD_CoverageDescription_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_ImageDescription, stage.MD_ImageDescription_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_CoverageDescription_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ContentType, stage.MD_CoverageContentTypeCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Dimension, stage.MD_RangeDimension_PropertyTypes_instance)
}

func (reference *MD_FeatureCatalogueDescription_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_FeatureCatalogueDescription, stage.MD_FeatureCatalogueDescription_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_FeatureCatalogueDescription_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_ImageDescription_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_ImageDescription, stage.MD_ImageDescription_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_ImageDescription_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ContentType, stage.MD_CoverageContentTypeCode_PropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.ImagingCondition, stage.MD_ImagingConditionCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Dimension, stage.MD_RangeDimension_PropertyTypes_instance)
}

func (reference *MD_ImagingConditionCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_RangeDimension_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_RangeDimension, stage.MD_RangeDimension_Types_instance)
	__gong__reconstructPointerFromInstance(&reference.MD_Band, stage.MD_Band_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_RangeDimension_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractmd_contentinformation_type *AbstractMD_ContentInformation_Type) GongDiff(stage *Stage, abstractmd_contentinformation_typeOther *AbstractMD_ContentInformation_Type) (diffs []string) {
	// insertion point for field diffs
	if abstractmd_contentinformation_type.Name != abstractmd_contentinformation_typeOther.Name {
		diffs = append(diffs, abstractmd_contentinformation_type.GongMarshallField(stage, "Name"))
	}
	if abstractmd_contentinformation_type.Id != abstractmd_contentinformation_typeOther.Id {
		diffs = append(diffs, abstractmd_contentinformation_type.GongMarshallField(stage, "Id"))
	}
	if abstractmd_contentinformation_type.Uuid != abstractmd_contentinformation_typeOther.Uuid {
		diffs = append(diffs, abstractmd_contentinformation_type.GongMarshallField(stage, "Uuid"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_band_propertytype *MD_Band_PropertyType) GongDiff(stage *Stage, md_band_propertytypeOther *MD_Band_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_band_propertytype.Name != md_band_propertytypeOther.Name {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_band_propertytype.NilReason != md_band_propertytypeOther.NilReason {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_band_propertytype.Uuidref != md_band_propertytypeOther.Uuidref {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_band_propertytype.Type_XSD != md_band_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_band_propertytype.Href != md_band_propertytypeOther.Href {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_band_propertytype.Role != md_band_propertytypeOther.Role {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_band_propertytype.Arcrole != md_band_propertytypeOther.Arcrole {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_band_propertytype.Title != md_band_propertytypeOther.Title {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_band_propertytype.Show != md_band_propertytypeOther.Show {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_band_propertytype.Actuate != md_band_propertytypeOther.Actuate {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_band_propertytype.MD_Band != md_band_propertytypeOther.MD_Band {
		diffs = append(diffs, md_band_propertytype.GongMarshallField(stage, "MD_Band"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_band_type *MD_Band_Type) GongDiff(stage *Stage, md_band_typeOther *MD_Band_Type) (diffs []string) {
	// insertion point for field diffs
	if md_band_type.Name != md_band_typeOther.Name {
		diffs = append(diffs, md_band_type.GongMarshallField(stage, "Name"))
	}
	if md_band_type.Id != md_band_typeOther.Id {
		diffs = append(diffs, md_band_type.GongMarshallField(stage, "Id"))
	}
	if md_band_type.Uuid != md_band_typeOther.Uuid {
		diffs = append(diffs, md_band_type.GongMarshallField(stage, "Uuid"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_contentinformation_propertytype *MD_ContentInformation_PropertyType) GongDiff(stage *Stage, md_contentinformation_propertytypeOther *MD_ContentInformation_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_contentinformation_propertytype.Name != md_contentinformation_propertytypeOther.Name {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_contentinformation_propertytype.NilReason != md_contentinformation_propertytypeOther.NilReason {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_contentinformation_propertytype.Uuidref != md_contentinformation_propertytypeOther.Uuidref {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_contentinformation_propertytype.Type_XSD != md_contentinformation_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_contentinformation_propertytype.Href != md_contentinformation_propertytypeOther.Href {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_contentinformation_propertytype.Role != md_contentinformation_propertytypeOther.Role {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_contentinformation_propertytype.Arcrole != md_contentinformation_propertytypeOther.Arcrole {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_contentinformation_propertytype.Title != md_contentinformation_propertytypeOther.Title {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_contentinformation_propertytype.Show != md_contentinformation_propertytypeOther.Show {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_contentinformation_propertytype.Actuate != md_contentinformation_propertytypeOther.Actuate {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_contentinformation_propertytype.MD_FeatureCatalogueDescription != md_contentinformation_propertytypeOther.MD_FeatureCatalogueDescription {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "MD_FeatureCatalogueDescription"))
	}
	if md_contentinformation_propertytype.MD_CoverageDescription != md_contentinformation_propertytypeOther.MD_CoverageDescription {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "MD_CoverageDescription"))
	}
	if md_contentinformation_propertytype.MD_ImageDescription != md_contentinformation_propertytypeOther.MD_ImageDescription {
		diffs = append(diffs, md_contentinformation_propertytype.GongMarshallField(stage, "MD_ImageDescription"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_coveragecontenttypecode_propertytype *MD_CoverageContentTypeCode_PropertyType) GongDiff(stage *Stage, md_coveragecontenttypecode_propertytypeOther *MD_CoverageContentTypeCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_coveragecontenttypecode_propertytype.Name != md_coveragecontenttypecode_propertytypeOther.Name {
		diffs = append(diffs, md_coveragecontenttypecode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_coveragecontenttypecode_propertytype.NilReason != md_coveragecontenttypecode_propertytypeOther.NilReason {
		diffs = append(diffs, md_coveragecontenttypecode_propertytype.GongMarshallField(stage, "NilReason"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_coveragedescription_propertytype *MD_CoverageDescription_PropertyType) GongDiff(stage *Stage, md_coveragedescription_propertytypeOther *MD_CoverageDescription_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_coveragedescription_propertytype.Name != md_coveragedescription_propertytypeOther.Name {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_coveragedescription_propertytype.NilReason != md_coveragedescription_propertytypeOther.NilReason {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_coveragedescription_propertytype.Uuidref != md_coveragedescription_propertytypeOther.Uuidref {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_coveragedescription_propertytype.Type_XSD != md_coveragedescription_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_coveragedescription_propertytype.Href != md_coveragedescription_propertytypeOther.Href {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_coveragedescription_propertytype.Role != md_coveragedescription_propertytypeOther.Role {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_coveragedescription_propertytype.Arcrole != md_coveragedescription_propertytypeOther.Arcrole {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_coveragedescription_propertytype.Title != md_coveragedescription_propertytypeOther.Title {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_coveragedescription_propertytype.Show != md_coveragedescription_propertytypeOther.Show {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_coveragedescription_propertytype.Actuate != md_coveragedescription_propertytypeOther.Actuate {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_coveragedescription_propertytype.MD_CoverageDescription != md_coveragedescription_propertytypeOther.MD_CoverageDescription {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "MD_CoverageDescription"))
	}
	if md_coveragedescription_propertytype.MD_ImageDescription != md_coveragedescription_propertytypeOther.MD_ImageDescription {
		diffs = append(diffs, md_coveragedescription_propertytype.GongMarshallField(stage, "MD_ImageDescription"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_coveragedescription_type *MD_CoverageDescription_Type) GongDiff(stage *Stage, md_coveragedescription_typeOther *MD_CoverageDescription_Type) (diffs []string) {
	// insertion point for field diffs
	if md_coveragedescription_type.Name != md_coveragedescription_typeOther.Name {
		diffs = append(diffs, md_coveragedescription_type.GongMarshallField(stage, "Name"))
	}
	if md_coveragedescription_type.Id != md_coveragedescription_typeOther.Id {
		diffs = append(diffs, md_coveragedescription_type.GongMarshallField(stage, "Id"))
	}
	if md_coveragedescription_type.Uuid != md_coveragedescription_typeOther.Uuid {
		diffs = append(diffs, md_coveragedescription_type.GongMarshallField(stage, "Uuid"))
	}
	if md_coveragedescription_type.ContentType != md_coveragedescription_typeOther.ContentType {
		diffs = append(diffs, md_coveragedescription_type.GongMarshallField(stage, "ContentType"))
	}
	if ops := __gong__diffSliceOfPointers(stage, md_coveragedescription_type, "Dimension", md_coveragedescription_typeOther.Dimension, md_coveragedescription_type.Dimension); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_featurecataloguedescription_propertytype *MD_FeatureCatalogueDescription_PropertyType) GongDiff(stage *Stage, md_featurecataloguedescription_propertytypeOther *MD_FeatureCatalogueDescription_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_featurecataloguedescription_propertytype.Name != md_featurecataloguedescription_propertytypeOther.Name {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_featurecataloguedescription_propertytype.NilReason != md_featurecataloguedescription_propertytypeOther.NilReason {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_featurecataloguedescription_propertytype.Uuidref != md_featurecataloguedescription_propertytypeOther.Uuidref {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_featurecataloguedescription_propertytype.Type_XSD != md_featurecataloguedescription_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_featurecataloguedescription_propertytype.Href != md_featurecataloguedescription_propertytypeOther.Href {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_featurecataloguedescription_propertytype.Role != md_featurecataloguedescription_propertytypeOther.Role {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_featurecataloguedescription_propertytype.Arcrole != md_featurecataloguedescription_propertytypeOther.Arcrole {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_featurecataloguedescription_propertytype.Title != md_featurecataloguedescription_propertytypeOther.Title {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_featurecataloguedescription_propertytype.Show != md_featurecataloguedescription_propertytypeOther.Show {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_featurecataloguedescription_propertytype.Actuate != md_featurecataloguedescription_propertytypeOther.Actuate {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_featurecataloguedescription_propertytype.MD_FeatureCatalogueDescription != md_featurecataloguedescription_propertytypeOther.MD_FeatureCatalogueDescription {
		diffs = append(diffs, md_featurecataloguedescription_propertytype.GongMarshallField(stage, "MD_FeatureCatalogueDescription"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_featurecataloguedescription_type *MD_FeatureCatalogueDescription_Type) GongDiff(stage *Stage, md_featurecataloguedescription_typeOther *MD_FeatureCatalogueDescription_Type) (diffs []string) {
	// insertion point for field diffs
	if md_featurecataloguedescription_type.Name != md_featurecataloguedescription_typeOther.Name {
		diffs = append(diffs, md_featurecataloguedescription_type.GongMarshallField(stage, "Name"))
	}
	if md_featurecataloguedescription_type.Id != md_featurecataloguedescription_typeOther.Id {
		diffs = append(diffs, md_featurecataloguedescription_type.GongMarshallField(stage, "Id"))
	}
	if md_featurecataloguedescription_type.Uuid != md_featurecataloguedescription_typeOther.Uuid {
		diffs = append(diffs, md_featurecataloguedescription_type.GongMarshallField(stage, "Uuid"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_imagedescription_propertytype *MD_ImageDescription_PropertyType) GongDiff(stage *Stage, md_imagedescription_propertytypeOther *MD_ImageDescription_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_imagedescription_propertytype.Name != md_imagedescription_propertytypeOther.Name {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_imagedescription_propertytype.NilReason != md_imagedescription_propertytypeOther.NilReason {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_imagedescription_propertytype.Uuidref != md_imagedescription_propertytypeOther.Uuidref {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_imagedescription_propertytype.Type_XSD != md_imagedescription_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_imagedescription_propertytype.Href != md_imagedescription_propertytypeOther.Href {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_imagedescription_propertytype.Role != md_imagedescription_propertytypeOther.Role {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_imagedescription_propertytype.Arcrole != md_imagedescription_propertytypeOther.Arcrole {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_imagedescription_propertytype.Title != md_imagedescription_propertytypeOther.Title {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_imagedescription_propertytype.Show != md_imagedescription_propertytypeOther.Show {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_imagedescription_propertytype.Actuate != md_imagedescription_propertytypeOther.Actuate {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_imagedescription_propertytype.MD_ImageDescription != md_imagedescription_propertytypeOther.MD_ImageDescription {
		diffs = append(diffs, md_imagedescription_propertytype.GongMarshallField(stage, "MD_ImageDescription"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_imagedescription_type *MD_ImageDescription_Type) GongDiff(stage *Stage, md_imagedescription_typeOther *MD_ImageDescription_Type) (diffs []string) {
	// insertion point for field diffs
	if md_imagedescription_type.Name != md_imagedescription_typeOther.Name {
		diffs = append(diffs, md_imagedescription_type.GongMarshallField(stage, "Name"))
	}
	if md_imagedescription_type.Id != md_imagedescription_typeOther.Id {
		diffs = append(diffs, md_imagedescription_type.GongMarshallField(stage, "Id"))
	}
	if md_imagedescription_type.Uuid != md_imagedescription_typeOther.Uuid {
		diffs = append(diffs, md_imagedescription_type.GongMarshallField(stage, "Uuid"))
	}
	if md_imagedescription_type.ContentType != md_imagedescription_typeOther.ContentType {
		diffs = append(diffs, md_imagedescription_type.GongMarshallField(stage, "ContentType"))
	}
	if ops := __gong__diffSliceOfPointers(stage, md_imagedescription_type, "Dimension", md_imagedescription_typeOther.Dimension, md_imagedescription_type.Dimension); ops != "" {
		diffs = append(diffs, ops)
	}
	if md_imagedescription_type.ImagingCondition != md_imagedescription_typeOther.ImagingCondition {
		diffs = append(diffs, md_imagedescription_type.GongMarshallField(stage, "ImagingCondition"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_imagingconditioncode_propertytype *MD_ImagingConditionCode_PropertyType) GongDiff(stage *Stage, md_imagingconditioncode_propertytypeOther *MD_ImagingConditionCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_imagingconditioncode_propertytype.Name != md_imagingconditioncode_propertytypeOther.Name {
		diffs = append(diffs, md_imagingconditioncode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_imagingconditioncode_propertytype.NilReason != md_imagingconditioncode_propertytypeOther.NilReason {
		diffs = append(diffs, md_imagingconditioncode_propertytype.GongMarshallField(stage, "NilReason"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_rangedimension_propertytype *MD_RangeDimension_PropertyType) GongDiff(stage *Stage, md_rangedimension_propertytypeOther *MD_RangeDimension_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_rangedimension_propertytype.Name != md_rangedimension_propertytypeOther.Name {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_rangedimension_propertytype.NilReason != md_rangedimension_propertytypeOther.NilReason {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_rangedimension_propertytype.Uuidref != md_rangedimension_propertytypeOther.Uuidref {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_rangedimension_propertytype.Type_XSD != md_rangedimension_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_rangedimension_propertytype.Href != md_rangedimension_propertytypeOther.Href {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_rangedimension_propertytype.Role != md_rangedimension_propertytypeOther.Role {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_rangedimension_propertytype.Arcrole != md_rangedimension_propertytypeOther.Arcrole {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_rangedimension_propertytype.Title != md_rangedimension_propertytypeOther.Title {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_rangedimension_propertytype.Show != md_rangedimension_propertytypeOther.Show {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_rangedimension_propertytype.Actuate != md_rangedimension_propertytypeOther.Actuate {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_rangedimension_propertytype.MD_RangeDimension != md_rangedimension_propertytypeOther.MD_RangeDimension {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "MD_RangeDimension"))
	}
	if md_rangedimension_propertytype.MD_Band != md_rangedimension_propertytypeOther.MD_Band {
		diffs = append(diffs, md_rangedimension_propertytype.GongMarshallField(stage, "MD_Band"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_rangedimension_type *MD_RangeDimension_Type) GongDiff(stage *Stage, md_rangedimension_typeOther *MD_RangeDimension_Type) (diffs []string) {
	// insertion point for field diffs
	if md_rangedimension_type.Name != md_rangedimension_typeOther.Name {
		diffs = append(diffs, md_rangedimension_type.GongMarshallField(stage, "Name"))
	}
	if md_rangedimension_type.Id != md_rangedimension_typeOther.Id {
		diffs = append(diffs, md_rangedimension_type.GongMarshallField(stage, "Id"))
	}
	if md_rangedimension_type.Uuid != md_rangedimension_typeOther.Uuid {
		diffs = append(diffs, md_rangedimension_type.GongMarshallField(stage, "Uuid"))
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
