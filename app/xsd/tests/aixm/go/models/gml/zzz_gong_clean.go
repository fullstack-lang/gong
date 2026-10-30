// generated code - do not edit
package gml

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by A_Row
func (a_row *A_Row) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&a_row.PosList) || modified
	modified = stage.CleanPointer(&a_row.Pos) || modified
	modified = stage.CleanPointer(&a_row.PointProperty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by A_controlPoint
func (a_controlpoint *A_controlPoint) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&a_controlpoint.PosList) || modified
	modified = stage.CleanPointer(&a_controlpoint.Pos) || modified
	modified = stage.CleanPointer(&a_controlpoint.PointProperty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by A_refLocation
func (a_reflocation *A_refLocation) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&a_reflocation.AffinePlacement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by A_rows
func (a_rows *A_rows) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&a_rows.Row) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractCRSType
func (abstractcrstype *AbstractCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstractcrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractcrstype.Description) || modified
	modified = stage.CleanPointer(&abstractcrstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractContinuousCoverageType
func (abstractcontinuouscoveragetype *AbstractContinuousCoverageType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractcontinuouscoveragetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.Description) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.BoundedBy) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.Location) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.PriorityLocation) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.DomainSet) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.MultiPointDomain) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.MultiCurveDomain) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.MultiSurfaceDomain) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.MultiSolidDomain) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.GridDomain) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.RectifiedGridDomain) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.RangeSet) || modified
	modified = stage.CleanPointer(&abstractcontinuouscoveragetype.CoverageFunction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractCoordinateOperationType
func (abstractcoordinateoperationtype *AbstractCoordinateOperationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractcoordinateoperationtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstractcoordinateoperationtype.CoordinateOperationAccuracy) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractcoordinateoperationtype.Description) || modified
	modified = stage.CleanPointer(&abstractcoordinateoperationtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractcoordinateoperationtype.DomainOfValidity) || modified
	modified = stage.CleanPointer(&abstractcoordinateoperationtype.SourceCRS) || modified
	modified = stage.CleanPointer(&abstractcoordinateoperationtype.TargetCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractCoordinateSystemType
func (abstractcoordinatesystemtype *AbstractCoordinateSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractcoordinatesystemtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstractcoordinatesystemtype.Axis) || modified
	modified = stage.CleanSlice(&abstractcoordinatesystemtype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractcoordinatesystemtype.Description) || modified
	modified = stage.CleanPointer(&abstractcoordinatesystemtype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractCoverageType
func (abstractcoveragetype *AbstractCoverageType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractcoveragetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractcoveragetype.Description) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.BoundedBy) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.Location) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.PriorityLocation) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.DomainSet) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.MultiPointDomain) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.MultiCurveDomain) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.MultiSurfaceDomain) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.MultiSolidDomain) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.GridDomain) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.RectifiedGridDomain) || modified
	modified = stage.CleanPointer(&abstractcoveragetype.RangeSet) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractCurveType
func (abstractcurvetype *AbstractCurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractcurvetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractcurvetype.Description) || modified
	modified = stage.CleanPointer(&abstractcurvetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractDQ_Completeness_Type
func (abstractdq_completeness_type *AbstractDQ_Completeness_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractdq_completeness_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&abstractdq_completeness_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractdq_completeness_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&abstractdq_completeness_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&abstractdq_completeness_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&abstractdq_completeness_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&abstractdq_completeness_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&abstractdq_completeness_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractDQ_Element_Type
func (abstractdq_element_type *AbstractDQ_Element_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractdq_element_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&abstractdq_element_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractdq_element_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&abstractdq_element_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&abstractdq_element_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&abstractdq_element_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&abstractdq_element_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&abstractdq_element_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractDQ_LogicalConsistency_Type
func (abstractdq_logicalconsistency_type *AbstractDQ_LogicalConsistency_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractdq_logicalconsistency_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&abstractdq_logicalconsistency_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractdq_logicalconsistency_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&abstractdq_logicalconsistency_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&abstractdq_logicalconsistency_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&abstractdq_logicalconsistency_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&abstractdq_logicalconsistency_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&abstractdq_logicalconsistency_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractDQ_PositionalAccuracy_Type
func (abstractdq_positionalaccuracy_type *AbstractDQ_PositionalAccuracy_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractdq_positionalaccuracy_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&abstractdq_positionalaccuracy_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractdq_positionalaccuracy_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&abstractdq_positionalaccuracy_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&abstractdq_positionalaccuracy_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&abstractdq_positionalaccuracy_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&abstractdq_positionalaccuracy_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&abstractdq_positionalaccuracy_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractDQ_TemporalAccuracy_Type
func (abstractdq_temporalaccuracy_type *AbstractDQ_TemporalAccuracy_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractdq_temporalaccuracy_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&abstractdq_temporalaccuracy_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractdq_temporalaccuracy_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&abstractdq_temporalaccuracy_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&abstractdq_temporalaccuracy_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&abstractdq_temporalaccuracy_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&abstractdq_temporalaccuracy_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&abstractdq_temporalaccuracy_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractDQ_ThematicAccuracy_Type
func (abstractdq_thematicaccuracy_type *AbstractDQ_ThematicAccuracy_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractdq_thematicaccuracy_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&abstractdq_thematicaccuracy_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractdq_thematicaccuracy_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&abstractdq_thematicaccuracy_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&abstractdq_thematicaccuracy_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&abstractdq_thematicaccuracy_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&abstractdq_thematicaccuracy_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&abstractdq_thematicaccuracy_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractDatumType
func (abstractdatumtype *AbstractDatumType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractdatumtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractdatumtype.Description) || modified
	modified = stage.CleanPointer(&abstractdatumtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractdatumtype.DomainOfValidity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractEX_GeographicExtent_Type
func (abstractex_geographicextent_type *AbstractEX_GeographicExtent_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&abstractex_geographicextent_type.ExtentTypeCode) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractFeatureCollectionType
func (abstractfeaturecollectiontype *AbstractFeatureCollectionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractfeaturecollectiontype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstractfeaturecollectiontype.FeatureMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractfeaturecollectiontype.Description) || modified
	modified = stage.CleanPointer(&abstractfeaturecollectiontype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractfeaturecollectiontype.BoundedBy) || modified
	modified = stage.CleanPointer(&abstractfeaturecollectiontype.Location) || modified
	modified = stage.CleanPointer(&abstractfeaturecollectiontype.PriorityLocation) || modified
	modified = stage.CleanPointer(&abstractfeaturecollectiontype.FeatureMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractFeatureType
func (abstractfeaturetype *AbstractFeatureType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractfeaturetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractfeaturetype.Description) || modified
	modified = stage.CleanPointer(&abstractfeaturetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractfeaturetype.BoundedBy) || modified
	modified = stage.CleanPointer(&abstractfeaturetype.Location) || modified
	modified = stage.CleanPointer(&abstractfeaturetype.PriorityLocation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGMLType
func (abstractgmltype *AbstractGMLType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgmltype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgmltype.Description) || modified
	modified = stage.CleanPointer(&abstractgmltype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeneralConversionType
func (abstractgeneralconversiontype *AbstractGeneralConversionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeneralconversiontype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstractgeneralconversiontype.CoordinateOperationAccuracy) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeneralconversiontype.Description) || modified
	modified = stage.CleanPointer(&abstractgeneralconversiontype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractgeneralconversiontype.DomainOfValidity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeneralDerivedCRSType
func (abstractgeneralderivedcrstype *AbstractGeneralDerivedCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeneralderivedcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstractgeneralderivedcrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeneralderivedcrstype.Description) || modified
	modified = stage.CleanPointer(&abstractgeneralderivedcrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractgeneralderivedcrstype.Conversion) || modified
	modified = stage.CleanPointer(&abstractgeneralderivedcrstype.DefinedByConversion) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeneralOperationParameterPropertyType
func (abstractgeneraloperationparameterpropertytype *AbstractGeneralOperationParameterPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeneraloperationparameterpropertytype.OperationParameter) || modified
	modified = stage.CleanPointer(&abstractgeneraloperationparameterpropertytype.OperationParameterGroup) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeneralOperationParameterType
func (abstractgeneraloperationparametertype *AbstractGeneralOperationParameterType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeneraloperationparametertype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeneraloperationparametertype.Description) || modified
	modified = stage.CleanPointer(&abstractgeneraloperationparametertype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeneralParameterValuePropertyType
func (abstractgeneralparametervaluepropertytype *AbstractGeneralParameterValuePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeneralparametervaluepropertytype.ParameterValue) || modified
	modified = stage.CleanPointer(&abstractgeneralparametervaluepropertytype.ParameterValueGroup) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeneralTransformationType
func (abstractgeneraltransformationtype *AbstractGeneralTransformationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeneraltransformationtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstractgeneraltransformationtype.CoordinateOperationAccuracy) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeneraltransformationtype.Description) || modified
	modified = stage.CleanPointer(&abstractgeneraltransformationtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstractgeneraltransformationtype.DomainOfValidity) || modified
	modified = stage.CleanPointer(&abstractgeneraltransformationtype.SourceCRS) || modified
	modified = stage.CleanPointer(&abstractgeneraltransformationtype.TargetCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeometricAggregateType
func (abstractgeometricaggregatetype *AbstractGeometricAggregateType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeometricaggregatetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeometricaggregatetype.Description) || modified
	modified = stage.CleanPointer(&abstractgeometricaggregatetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeometricPrimitiveType
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeometricprimitivetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeometricprimitivetype.Description) || modified
	modified = stage.CleanPointer(&abstractgeometricprimitivetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGeometryType
func (abstractgeometrytype *AbstractGeometryType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractgeometrytype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractgeometrytype.Description) || modified
	modified = stage.CleanPointer(&abstractgeometrytype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractGriddedSurfaceType
func (abstractgriddedsurfacetype *AbstractGriddedSurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&abstractgriddedsurfacetype.Rows_1) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractMD_Identification_Type
func (abstractmd_identification_type *AbstractMD_Identification_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractmd_identification_type.Credit) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.Status) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.PointOfContact) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.ResourceMaintenance) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.GraphicOverview) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.ResourceFormat) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.DescriptiveKeywords) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.ResourceSpecificUsage) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.ResourceConstraints) || modified
	modified = stage.CleanSlice(&abstractmd_identification_type.AggregationInfo) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractmd_identification_type.Citation) || modified
	modified = stage.CleanPointer(&abstractmd_identification_type.Abstract) || modified
	modified = stage.CleanPointer(&abstractmd_identification_type.Purpose) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractRS_ReferenceSystem_Type
func (abstractrs_referencesystem_type *AbstractRS_ReferenceSystem_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractrs_referencesystem_type.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractrs_referencesystem_type.NameXSD) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractRingPropertyType
func (abstractringpropertytype *AbstractRingPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&abstractringpropertytype.Ring) || modified
	modified = stage.CleanPointer(&abstractringpropertytype.LinearRing) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractRingType
func (abstractringtype *AbstractRingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractringtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractringtype.Description) || modified
	modified = stage.CleanPointer(&abstractringtype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractSolidType
func (abstractsolidtype *AbstractSolidType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractsolidtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractsolidtype.Description) || modified
	modified = stage.CleanPointer(&abstractsolidtype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractSurfaceType
func (abstractsurfacetype *AbstractSurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstractsurfacetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstractsurfacetype.Description) || modified
	modified = stage.CleanPointer(&abstractsurfacetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimeComplexType
func (abstracttimecomplextype *AbstractTimeComplexType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimecomplextype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimecomplextype.Description) || modified
	modified = stage.CleanPointer(&abstracttimecomplextype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimeGeometricPrimitiveType
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimegeometricprimitivetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstracttimegeometricprimitivetype.RelatedTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimegeometricprimitivetype.Description) || modified
	modified = stage.CleanPointer(&abstracttimegeometricprimitivetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimeObjectType
func (abstracttimeobjecttype *AbstractTimeObjectType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimeobjecttype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimeobjecttype.Description) || modified
	modified = stage.CleanPointer(&abstracttimeobjecttype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimePrimitiveType
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimeprimitivetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstracttimeprimitivetype.RelatedTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimeprimitivetype.Description) || modified
	modified = stage.CleanPointer(&abstracttimeprimitivetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimeSliceType
func (abstracttimeslicetype *AbstractTimeSliceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimeslicetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimeslicetype.Description) || modified
	modified = stage.CleanPointer(&abstracttimeslicetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstracttimeslicetype.ValidTime) || modified
	modified = stage.CleanPointer(&abstracttimeslicetype.DataSource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTimeTopologyPrimitiveType
func (abstracttimetopologyprimitivetype *AbstractTimeTopologyPrimitiveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttimetopologyprimitivetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&abstracttimetopologyprimitivetype.RelatedTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttimetopologyprimitivetype.Description) || modified
	modified = stage.CleanPointer(&abstracttimetopologyprimitivetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&abstracttimetopologyprimitivetype.Complex_XSD) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTopoPrimitiveType
func (abstracttopoprimitivetype *AbstractTopoPrimitiveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttopoprimitivetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttopoprimitivetype.Description) || modified
	modified = stage.CleanPointer(&abstracttopoprimitivetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AbstractTopologyType
func (abstracttopologytype *AbstractTopologyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&abstracttopologytype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&abstracttopologytype.Description) || modified
	modified = stage.CleanPointer(&abstracttopologytype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AffineCSPropertyType
func (affinecspropertytype *AffineCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&affinecspropertytype.AffineCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AffineCSType
func (affinecstype *AffineCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&affinecstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&affinecstype.Axis) || modified
	modified = stage.CleanSlice(&affinecstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&affinecstype.Description) || modified
	modified = stage.CleanPointer(&affinecstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AffinePlacementType
func (affineplacementtype *AffinePlacementType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&affineplacementtype.RefDirection) || modified
	// insertion point per field
	modified = stage.CleanPointer(&affineplacementtype.Location) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AngleChoiceType
func (anglechoicetype *AngleChoiceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&anglechoicetype.DmsAngle) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArcByBulgeType
func (arcbybulgetype *ArcByBulgeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&arcbybulgetype.Pos) || modified
	modified = stage.CleanPointer(&arcbybulgetype.PointProperty) || modified
	modified = stage.CleanPointer(&arcbybulgetype.PointRep) || modified
	modified = stage.CleanPointer(&arcbybulgetype.PosList) || modified
	modified = stage.CleanPointer(&arcbybulgetype.Normal) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArcByCenterPointType
func (arcbycenterpointtype *ArcByCenterPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&arcbycenterpointtype.Pos) || modified
	modified = stage.CleanPointer(&arcbycenterpointtype.PointProperty) || modified
	modified = stage.CleanPointer(&arcbycenterpointtype.PointRep) || modified
	modified = stage.CleanPointer(&arcbycenterpointtype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArcStringByBulgeType
func (arcstringbybulgetype *ArcStringByBulgeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arcstringbybulgetype.Pos) || modified
	modified = stage.CleanSlice(&arcstringbybulgetype.PointProperty) || modified
	modified = stage.CleanSlice(&arcstringbybulgetype.PointRep) || modified
	modified = stage.CleanSlice(&arcstringbybulgetype.Normal) || modified
	// insertion point per field
	modified = stage.CleanPointer(&arcstringbybulgetype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArcStringType
func (arcstringtype *ArcStringType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arcstringtype.Pos) || modified
	modified = stage.CleanSlice(&arcstringtype.PointProperty) || modified
	modified = stage.CleanSlice(&arcstringtype.PointRep) || modified
	// insertion point per field
	modified = stage.CleanPointer(&arcstringtype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArcType
func (arctype *ArcType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&arctype.Pos) || modified
	modified = stage.CleanPointer(&arctype.PointProperty) || modified
	modified = stage.CleanPointer(&arctype.PointRep) || modified
	modified = stage.CleanPointer(&arctype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrayAssociationType
func (arrayassociationtype *ArrayAssociationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arrayassociationtype.LineStringSegment) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ArcString) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Arc) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Circle) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ArcStringByBulge) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ArcByBulge) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ArcByCenterPoint) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CircleByCenterPoint) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CubicSpline) || modified
	modified = stage.CleanSlice(&arrayassociationtype.BSpline) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Bezier) || modified
	modified = stage.CleanSlice(&arrayassociationtype.OffsetCurve) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Clothoid) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GeodesicString) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Geodesic) || modified
	modified = stage.CleanSlice(&arrayassociationtype.AffinePlacement) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Envelope) || modified
	modified = stage.CleanSlice(&arrayassociationtype.EnvelopeWithTimePeriod) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MovingObjectStatus) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DynamicFeature) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DynamicFeatureCollection) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiPointCoverage) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiCurveCoverage) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiSurfaceCoverage) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiSolidCoverage) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GridCoverage) || modified
	modified = stage.CleanSlice(&arrayassociationtype.RectifiedGridCoverage) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Observation) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DirectedObservation) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DirectedObservationAtDistance) || modified
	modified = stage.CleanSlice(&arrayassociationtype.FeatureCollection) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiGeometry) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiPoint) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiCurve) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiSurface) || modified
	modified = stage.CleanSlice(&arrayassociationtype.MultiSolid) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Solid) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CompositeSolid) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Surface) || modified
	modified = stage.CleanSlice(&arrayassociationtype.PolyhedralSurface) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TriangulatedSurface) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Tin) || modified
	modified = stage.CleanSlice(&arrayassociationtype.OrientableSurface) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Shell) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Polygon) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CompositeSurface) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Point) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Curve) || modified
	modified = stage.CleanSlice(&arrayassociationtype.OrientableCurve) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Ring) || modified
	modified = stage.CleanSlice(&arrayassociationtype.LinearRing) || modified
	modified = stage.CleanSlice(&arrayassociationtype.LineString) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CompositeCurve) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GeometricComplex) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Grid) || modified
	modified = stage.CleanSlice(&arrayassociationtype.RectifiedGrid) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Definition) || modified
	modified = stage.CleanSlice(&arrayassociationtype.UnitDefinition) || modified
	modified = stage.CleanSlice(&arrayassociationtype.BaseUnit) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DerivedUnit) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ConventionalUnit) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Dictionary) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CoordinateSystemAxis) || modified
	modified = stage.CleanSlice(&arrayassociationtype.EllipsoidalCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CartesianCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.VerticalCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.LinearCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.UserDefinedCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.SphericalCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.PolarCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CylindricalCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.AffineCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TemporalCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ObliqueCartesianCS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ProjectedCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DerivedCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GeodeticCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.VerticalCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.EngineeringCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ImageCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TemporalCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GeographicCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GeocentricCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CompoundCRS) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GeodeticDatum) || modified
	modified = stage.CleanSlice(&arrayassociationtype.EngineeringDatum) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ImageDatum) || modified
	modified = stage.CleanSlice(&arrayassociationtype.VerticalDatum) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TemporalDatum) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Ellipsoid) || modified
	modified = stage.CleanSlice(&arrayassociationtype.PrimeMeridian) || modified
	modified = stage.CleanSlice(&arrayassociationtype.PassThroughOperation) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Conversion) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Transformation) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ConcatenatedOperation) || modified
	modified = stage.CleanSlice(&arrayassociationtype.OperationMethod) || modified
	modified = stage.CleanSlice(&arrayassociationtype.OperationParameter) || modified
	modified = stage.CleanSlice(&arrayassociationtype.OperationParameterGroup) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeReferenceSystem) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeCoordinateSystem) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeCalendar) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeClock) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeOrdinalReferenceSystem) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DefinitionCollection) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DefinitionProxy) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeInstant) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimePeriod) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeNode) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeEdge) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TimeTopologyComplex) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Node) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Edge) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Face) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TopoSolid) || modified
	modified = stage.CleanSlice(&arrayassociationtype.TopoComplex) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Bag) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Array) || modified
	modified = stage.CleanSlice(&arrayassociationtype.DataBlock) || modified
	modified = stage.CleanSlice(&arrayassociationtype.File) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CoverageFunction) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CoverageMappingRule) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GridFunction) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Boolean) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Category) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Count) || modified
	modified = stage.CleanSlice(&arrayassociationtype.Quantity) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CompositeValue) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ValueArray) || modified
	modified = stage.CleanSlice(&arrayassociationtype.CategoryExtent) || modified
	modified = stage.CleanSlice(&arrayassociationtype.QuantityExtent) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ParameterValue) || modified
	modified = stage.CleanSlice(&arrayassociationtype.ParameterValueGroup) || modified
	modified = stage.CleanSlice(&arrayassociationtype.GenericMetaData) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ArrayType
func (arraytype *ArrayType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&arraytype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&arraytype.Description) || modified
	modified = stage.CleanPointer(&arraytype.DescriptionReference) || modified
	modified = stage.CleanPointer(&arraytype.Members) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BSplineType
func (bsplinetype *BSplineType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&bsplinetype.Pos) || modified
	modified = stage.CleanSlice(&bsplinetype.PointProperty) || modified
	modified = stage.CleanSlice(&bsplinetype.PointRep) || modified
	modified = stage.CleanSlice(&bsplinetype.Knot) || modified
	// insertion point per field
	modified = stage.CleanPointer(&bsplinetype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BagType
func (bagtype *BagType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&bagtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&bagtype.Member) || modified
	// insertion point per field
	modified = stage.CleanPointer(&bagtype.Description) || modified
	modified = stage.CleanPointer(&bagtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&bagtype.Members) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BaseUnitType
func (baseunittype *BaseUnitType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&baseunittype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&baseunittype.Description) || modified
	modified = stage.CleanPointer(&baseunittype.DescriptionReference) || modified
	modified = stage.CleanPointer(&baseunittype.QuantityType) || modified
	modified = stage.CleanPointer(&baseunittype.QuantityTypeReference) || modified
	modified = stage.CleanPointer(&baseunittype.UnitsSystem) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BezierType
func (beziertype *BezierType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&beziertype.Pos) || modified
	modified = stage.CleanSlice(&beziertype.PointProperty) || modified
	modified = stage.CleanSlice(&beziertype.PointRep) || modified
	// insertion point per field
	modified = stage.CleanPointer(&beziertype.PosList) || modified
	modified = stage.CleanPointer(&beziertype.Knot) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Binary_PropertyType
func (binary_propertytype *Binary_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&binary_propertytype.Binary) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BooleanPropertyType
func (booleanpropertytype *BooleanPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&booleanpropertytype.Boolean) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BoundedFeatureType
func (boundedfeaturetype *BoundedFeatureType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&boundedfeaturetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&boundedfeaturetype.Description) || modified
	modified = stage.CleanPointer(&boundedfeaturetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&boundedfeaturetype.BoundedBy) || modified
	modified = stage.CleanPointer(&boundedfeaturetype.Location) || modified
	modified = stage.CleanPointer(&boundedfeaturetype.PriorityLocation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by BoundingShapeType
func (boundingshapetype *BoundingShapeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&boundingshapetype.Envelope) || modified
	modified = stage.CleanPointer(&boundingshapetype.EnvelopeWithTimePeriod) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Address_PropertyType
func (ci_address_propertytype *CI_Address_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_address_propertytype.CI_Address) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Address_Type
func (ci_address_type *CI_Address_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ci_address_type.DeliveryPoint) || modified
	modified = stage.CleanSlice(&ci_address_type.ElectronicMailAddress) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ci_address_type.City) || modified
	modified = stage.CleanPointer(&ci_address_type.AdministrativeArea) || modified
	modified = stage.CleanPointer(&ci_address_type.PostalCode) || modified
	modified = stage.CleanPointer(&ci_address_type.Country) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Citation_PropertyType
func (ci_citation_propertytype *CI_Citation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_citation_propertytype.CI_Citation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Citation_Type
func (ci_citation_type *CI_Citation_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ci_citation_type.AlternateTitle) || modified
	modified = stage.CleanSlice(&ci_citation_type.Date) || modified
	modified = stage.CleanSlice(&ci_citation_type.Identifier) || modified
	modified = stage.CleanSlice(&ci_citation_type.CitedResponsibleParty) || modified
	modified = stage.CleanSlice(&ci_citation_type.PresentationForm) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ci_citation_type.Title) || modified
	modified = stage.CleanPointer(&ci_citation_type.Edition) || modified
	modified = stage.CleanPointer(&ci_citation_type.EditionDate) || modified
	modified = stage.CleanPointer(&ci_citation_type.Series) || modified
	modified = stage.CleanPointer(&ci_citation_type.OtherCitationDetails) || modified
	modified = stage.CleanPointer(&ci_citation_type.CollectiveTitle) || modified
	modified = stage.CleanPointer(&ci_citation_type.ISBN) || modified
	modified = stage.CleanPointer(&ci_citation_type.ISSN) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Contact_PropertyType
func (ci_contact_propertytype *CI_Contact_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_contact_propertytype.CI_Contact) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Contact_Type
func (ci_contact_type *CI_Contact_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_contact_type.Phone) || modified
	modified = stage.CleanPointer(&ci_contact_type.Address) || modified
	modified = stage.CleanPointer(&ci_contact_type.OnlineResource) || modified
	modified = stage.CleanPointer(&ci_contact_type.HoursOfService) || modified
	modified = stage.CleanPointer(&ci_contact_type.ContactInstructions) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Date_PropertyType
func (ci_date_propertytype *CI_Date_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_date_propertytype.CI_Date) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Date_Type
func (ci_date_type *CI_Date_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_date_type.Date) || modified
	modified = stage.CleanPointer(&ci_date_type.DateType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_OnlineResource_PropertyType
func (ci_onlineresource_propertytype *CI_OnlineResource_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_onlineresource_propertytype.CI_OnlineResource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_OnlineResource_Type
func (ci_onlineresource_type *CI_OnlineResource_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_onlineresource_type.Linkage) || modified
	modified = stage.CleanPointer(&ci_onlineresource_type.Protocol) || modified
	modified = stage.CleanPointer(&ci_onlineresource_type.ApplicationProfile) || modified
	modified = stage.CleanPointer(&ci_onlineresource_type.NameXSD) || modified
	modified = stage.CleanPointer(&ci_onlineresource_type.Description) || modified
	modified = stage.CleanPointer(&ci_onlineresource_type.Function) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_ResponsibleParty_PropertyType
func (ci_responsibleparty_propertytype *CI_ResponsibleParty_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_responsibleparty_propertytype.CI_ResponsibleParty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_ResponsibleParty_Type
func (ci_responsibleparty_type *CI_ResponsibleParty_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_responsibleparty_type.IndividualName) || modified
	modified = stage.CleanPointer(&ci_responsibleparty_type.OrganisationName) || modified
	modified = stage.CleanPointer(&ci_responsibleparty_type.PositionName) || modified
	modified = stage.CleanPointer(&ci_responsibleparty_type.ContactInfo) || modified
	modified = stage.CleanPointer(&ci_responsibleparty_type.Role) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Series_PropertyType
func (ci_series_propertytype *CI_Series_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_series_propertytype.CI_Series) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Series_Type
func (ci_series_type *CI_Series_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_series_type.NameXSD) || modified
	modified = stage.CleanPointer(&ci_series_type.IssueIdentification) || modified
	modified = stage.CleanPointer(&ci_series_type.Page) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Telephone_PropertyType
func (ci_telephone_propertytype *CI_Telephone_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ci_telephone_propertytype.CI_Telephone) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CI_Telephone_Type
func (ci_telephone_type *CI_Telephone_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ci_telephone_type.Voice) || modified
	modified = stage.CleanSlice(&ci_telephone_type.Facsimile) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CRSPropertyType
func (crspropertytype *CRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&crspropertytype.ProjectedCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.DerivedCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.GeodeticCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.VerticalCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.EngineeringCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.ImageCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.TemporalCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.GeographicCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.GeocentricCRS) || modified
	modified = stage.CleanPointer(&crspropertytype.CompoundCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CartesianCSPropertyType
func (cartesiancspropertytype *CartesianCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&cartesiancspropertytype.CartesianCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CartesianCSType
func (cartesiancstype *CartesianCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&cartesiancstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&cartesiancstype.Axis) || modified
	modified = stage.CleanSlice(&cartesiancstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&cartesiancstype.Description) || modified
	modified = stage.CleanPointer(&cartesiancstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CategoryPropertyType
func (categorypropertytype *CategoryPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&categorypropertytype.Category) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CharacterString_PropertyType
func (characterstring_propertytype *CharacterString_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&characterstring_propertytype.LocalisedCharacterString) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CircleByCenterPointType
func (circlebycenterpointtype *CircleByCenterPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&circlebycenterpointtype.Pos) || modified
	modified = stage.CleanPointer(&circlebycenterpointtype.PointProperty) || modified
	modified = stage.CleanPointer(&circlebycenterpointtype.PointRep) || modified
	modified = stage.CleanPointer(&circlebycenterpointtype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CircleType
func (circletype *CircleType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&circletype.Pos) || modified
	modified = stage.CleanSlice(&circletype.PointProperty) || modified
	modified = stage.CleanSlice(&circletype.PointRep) || modified
	// insertion point per field
	modified = stage.CleanPointer(&circletype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ClothoidType
func (clothoidtype *ClothoidType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&clothoidtype.RefLocation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CompositeCurveType
func (compositecurvetype *CompositeCurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&compositecurvetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&compositecurvetype.CurveMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&compositecurvetype.Description) || modified
	modified = stage.CleanPointer(&compositecurvetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CompositeSolidType
func (compositesolidtype *CompositeSolidType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&compositesolidtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&compositesolidtype.SolidMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&compositesolidtype.Description) || modified
	modified = stage.CleanPointer(&compositesolidtype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CompositeSurfaceType
func (compositesurfacetype *CompositeSurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&compositesurfacetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&compositesurfacetype.SurfaceMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&compositesurfacetype.Description) || modified
	modified = stage.CleanPointer(&compositesurfacetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CompositeValueType
func (compositevaluetype *CompositeValueType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&compositevaluetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&compositevaluetype.ValueComponent) || modified
	// insertion point per field
	modified = stage.CleanPointer(&compositevaluetype.Description) || modified
	modified = stage.CleanPointer(&compositevaluetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&compositevaluetype.ValueComponents) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CompoundCRSPropertyType
func (compoundcrspropertytype *CompoundCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&compoundcrspropertytype.CompoundCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CompoundCRSType
func (compoundcrstype *CompoundCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&compoundcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&compoundcrstype.DomainOfValidity) || modified
	modified = stage.CleanSlice(&compoundcrstype.ComponentReferenceSystem) || modified
	modified = stage.CleanSlice(&compoundcrstype.IncludesSingleCRS) || modified
	// insertion point per field
	modified = stage.CleanPointer(&compoundcrstype.Description) || modified
	modified = stage.CleanPointer(&compoundcrstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcatenatedOperationPropertyType
func (concatenatedoperationpropertytype *ConcatenatedOperationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&concatenatedoperationpropertytype.ConcatenatedOperation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConcatenatedOperationType
func (concatenatedoperationtype *ConcatenatedOperationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&concatenatedoperationtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&concatenatedoperationtype.CoordinateOperationAccuracy) || modified
	modified = stage.CleanSlice(&concatenatedoperationtype.CoordOperation) || modified
	modified = stage.CleanSlice(&concatenatedoperationtype.UsesSingleOperation) || modified
	modified = stage.CleanSlice(&concatenatedoperationtype.UsesOperation) || modified
	// insertion point per field
	modified = stage.CleanPointer(&concatenatedoperationtype.Description) || modified
	modified = stage.CleanPointer(&concatenatedoperationtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&concatenatedoperationtype.DomainOfValidity) || modified
	modified = stage.CleanPointer(&concatenatedoperationtype.SourceCRS) || modified
	modified = stage.CleanPointer(&concatenatedoperationtype.TargetCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConeType
func (conetype *ConeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&conetype.Rows_1) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConventionalUnitType
func (conventionalunittype *ConventionalUnitType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&conventionalunittype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&conventionalunittype.DerivationUnitTerm) || modified
	// insertion point per field
	modified = stage.CleanPointer(&conventionalunittype.Description) || modified
	modified = stage.CleanPointer(&conventionalunittype.DescriptionReference) || modified
	modified = stage.CleanPointer(&conventionalunittype.QuantityType) || modified
	modified = stage.CleanPointer(&conventionalunittype.QuantityTypeReference) || modified
	modified = stage.CleanPointer(&conventionalunittype.ConversionToPreferredUnit) || modified
	modified = stage.CleanPointer(&conventionalunittype.RoughConversionToPreferredUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConversionPropertyType
func (conversionpropertytype *ConversionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&conversionpropertytype.Conversion) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConversionToPreferredUnitType
func (conversiontopreferredunittype *ConversionToPreferredUnitType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&conversiontopreferredunittype.Formula) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ConversionType
func (conversiontype *ConversionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&conversiontype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&conversiontype.CoordinateOperationAccuracy) || modified
	modified = stage.CleanSlice(&conversiontype.ParameterValue) || modified
	modified = stage.CleanSlice(&conversiontype.UsesValue) || modified
	modified = stage.CleanSlice(&conversiontype.IncludesValue) || modified
	// insertion point per field
	modified = stage.CleanPointer(&conversiontype.Description) || modified
	modified = stage.CleanPointer(&conversiontype.DescriptionReference) || modified
	modified = stage.CleanPointer(&conversiontype.DomainOfValidity) || modified
	modified = stage.CleanPointer(&conversiontype.SourceCRS) || modified
	modified = stage.CleanPointer(&conversiontype.TargetCRS) || modified
	modified = stage.CleanPointer(&conversiontype.Method) || modified
	modified = stage.CleanPointer(&conversiontype.UsesMethod) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CoordinateOperationAccuracy
func (coordinateoperationaccuracy *CoordinateOperationAccuracy) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&coordinateoperationaccuracy.DQ_RelativeInternalPositionalAccuracy) || modified
	modified = stage.CleanPointer(&coordinateoperationaccuracy.DQ_GriddedDataPositionalAccuracy) || modified
	modified = stage.CleanPointer(&coordinateoperationaccuracy.DQ_AbsoluteExternalPositionalAccuracy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CoordinateOperationPropertyType
func (coordinateoperationpropertytype *CoordinateOperationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&coordinateoperationpropertytype.PassThroughOperation) || modified
	modified = stage.CleanPointer(&coordinateoperationpropertytype.Conversion) || modified
	modified = stage.CleanPointer(&coordinateoperationpropertytype.Transformation) || modified
	modified = stage.CleanPointer(&coordinateoperationpropertytype.ConcatenatedOperation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CoordinateSystemAxisPropertyType
func (coordinatesystemaxispropertytype *CoordinateSystemAxisPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&coordinatesystemaxispropertytype.CoordinateSystemAxis) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CoordinateSystemAxisType
func (coordinatesystemaxistype *CoordinateSystemAxisType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&coordinatesystemaxistype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&coordinatesystemaxistype.Description) || modified
	modified = stage.CleanPointer(&coordinatesystemaxistype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CoordinateSystemPropertyType
func (coordinatesystempropertytype *CoordinateSystemPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&coordinatesystempropertytype.EllipsoidalCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.CartesianCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.VerticalCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.TimeCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.LinearCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.UserDefinedCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.SphericalCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.PolarCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.CylindricalCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.AffineCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.TemporalCS) || modified
	modified = stage.CleanPointer(&coordinatesystempropertytype.ObliqueCartesianCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CountPropertyType
func (countpropertytype *CountPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&countpropertytype.Count) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CoverageFunctionType
func (coveragefunctiontype *CoverageFunctionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&coveragefunctiontype.MappingRule) || modified
	modified = stage.CleanPointer(&coveragefunctiontype.CoverageMappingRule) || modified
	modified = stage.CleanPointer(&coveragefunctiontype.GridFunction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CubicSplineType
func (cubicsplinetype *CubicSplineType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&cubicsplinetype.Pos) || modified
	modified = stage.CleanSlice(&cubicsplinetype.PointProperty) || modified
	modified = stage.CleanSlice(&cubicsplinetype.PointRep) || modified
	// insertion point per field
	modified = stage.CleanPointer(&cubicsplinetype.PosList) || modified
	modified = stage.CleanPointer(&cubicsplinetype.VectorAtStart) || modified
	modified = stage.CleanPointer(&cubicsplinetype.VectorAtEnd) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CurveArrayPropertyType
func (curvearraypropertytype *CurveArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&curvearraypropertytype.Curve) || modified
	modified = stage.CleanSlice(&curvearraypropertytype.OrientableCurve) || modified
	modified = stage.CleanSlice(&curvearraypropertytype.Ring) || modified
	modified = stage.CleanSlice(&curvearraypropertytype.LinearRing) || modified
	modified = stage.CleanSlice(&curvearraypropertytype.LineString) || modified
	modified = stage.CleanSlice(&curvearraypropertytype.CompositeCurve) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CurvePropertyType
func (curvepropertytype *CurvePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&curvepropertytype.Curve) || modified
	modified = stage.CleanPointer(&curvepropertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&curvepropertytype.Ring) || modified
	modified = stage.CleanPointer(&curvepropertytype.LinearRing) || modified
	modified = stage.CleanPointer(&curvepropertytype.LineString) || modified
	modified = stage.CleanPointer(&curvepropertytype.CompositeCurve) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CurveSegmentArrayPropertyType
func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.LineStringSegment) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.ArcString) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.Arc) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.Circle) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.ArcStringByBulge) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.ArcByBulge) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.ArcByCenterPoint) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.CircleByCenterPoint) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.CubicSpline) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.BSpline) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.Bezier) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.OffsetCurve) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.Clothoid) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.GeodesicString) || modified
	modified = stage.CleanSlice(&curvesegmentarraypropertytype.Geodesic) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by CurveType
func (curvetype *CurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&curvetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&curvetype.Description) || modified
	modified = stage.CleanPointer(&curvetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&curvetype.Segments) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CylinderType
func (cylindertype *CylinderType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&cylindertype.Rows_1) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CylindricalCSPropertyType
func (cylindricalcspropertytype *CylindricalCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&cylindricalcspropertytype.CylindricalCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by CylindricalCSType
func (cylindricalcstype *CylindricalCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&cylindricalcstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&cylindricalcstype.Axis) || modified
	modified = stage.CleanSlice(&cylindricalcstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&cylindricalcstype.Description) || modified
	modified = stage.CleanPointer(&cylindricalcstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DMSAngleType
func (dmsangletype *DMSAngleType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dmsangletype.Degrees) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_AbsoluteExternalPositionalAccuracy_PropertyType
func (dq_absoluteexternalpositionalaccuracy_propertytype *DQ_AbsoluteExternalPositionalAccuracy_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_absoluteexternalpositionalaccuracy_propertytype.DQ_AbsoluteExternalPositionalAccuracy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_AbsoluteExternalPositionalAccuracy_Type
func (dq_absoluteexternalpositionalaccuracy_type *DQ_AbsoluteExternalPositionalAccuracy_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_absoluteexternalpositionalaccuracy_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_absoluteexternalpositionalaccuracy_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_absoluteexternalpositionalaccuracy_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_absoluteexternalpositionalaccuracy_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_absoluteexternalpositionalaccuracy_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_absoluteexternalpositionalaccuracy_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_absoluteexternalpositionalaccuracy_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_absoluteexternalpositionalaccuracy_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_AccuracyOfATimeMeasurement_PropertyType
func (dq_accuracyofatimemeasurement_propertytype *DQ_AccuracyOfATimeMeasurement_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_accuracyofatimemeasurement_propertytype.DQ_AccuracyOfATimeMeasurement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_AccuracyOfATimeMeasurement_Type
func (dq_accuracyofatimemeasurement_type *DQ_AccuracyOfATimeMeasurement_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_accuracyofatimemeasurement_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_accuracyofatimemeasurement_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_accuracyofatimemeasurement_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_accuracyofatimemeasurement_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_accuracyofatimemeasurement_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_accuracyofatimemeasurement_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_accuracyofatimemeasurement_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_accuracyofatimemeasurement_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_CompletenessCommission_PropertyType
func (dq_completenesscommission_propertytype *DQ_CompletenessCommission_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_completenesscommission_propertytype.DQ_CompletenessCommission) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_CompletenessCommission_Type
func (dq_completenesscommission_type *DQ_CompletenessCommission_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_completenesscommission_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_completenesscommission_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_completenesscommission_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_completenesscommission_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_completenesscommission_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_completenesscommission_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_completenesscommission_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_completenesscommission_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_CompletenessOmission_PropertyType
func (dq_completenessomission_propertytype *DQ_CompletenessOmission_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_completenessomission_propertytype.DQ_CompletenessOmission) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_CompletenessOmission_Type
func (dq_completenessomission_type *DQ_CompletenessOmission_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_completenessomission_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_completenessomission_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_completenessomission_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_completenessomission_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_completenessomission_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_completenessomission_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_completenessomission_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_completenessomission_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_Completeness_PropertyType
func (dq_completeness_propertytype *DQ_Completeness_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_completeness_propertytype.DQ_CompletenessOmission) || modified
	modified = stage.CleanPointer(&dq_completeness_propertytype.DQ_CompletenessCommission) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_ConceptualConsistency_PropertyType
func (dq_conceptualconsistency_propertytype *DQ_ConceptualConsistency_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_conceptualconsistency_propertytype.DQ_ConceptualConsistency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_ConceptualConsistency_Type
func (dq_conceptualconsistency_type *DQ_ConceptualConsistency_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_conceptualconsistency_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_conceptualconsistency_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_conceptualconsistency_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_conceptualconsistency_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_conceptualconsistency_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_conceptualconsistency_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_conceptualconsistency_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_conceptualconsistency_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_ConformanceResult_PropertyType
func (dq_conformanceresult_propertytype *DQ_ConformanceResult_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_conformanceresult_propertytype.DQ_ConformanceResult) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_ConformanceResult_Type
func (dq_conformanceresult_type *DQ_ConformanceResult_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_conformanceresult_type.Specification) || modified
	modified = stage.CleanPointer(&dq_conformanceresult_type.Explanation) || modified
	modified = stage.CleanPointer(&dq_conformanceresult_type.Pass) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_DataQuality_PropertyType
func (dq_dataquality_propertytype *DQ_DataQuality_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_dataquality_propertytype.DQ_DataQuality) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_DataQuality_Type
func (dq_dataquality_type *DQ_DataQuality_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_dataquality_type.Report) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_dataquality_type.Scope) || modified
	modified = stage.CleanPointer(&dq_dataquality_type.Lineage) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_DomainConsistency_PropertyType
func (dq_domainconsistency_propertytype *DQ_DomainConsistency_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_domainconsistency_propertytype.DQ_DomainConsistency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_DomainConsistency_Type
func (dq_domainconsistency_type *DQ_DomainConsistency_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_domainconsistency_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_domainconsistency_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_domainconsistency_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_domainconsistency_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_domainconsistency_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_domainconsistency_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_domainconsistency_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_domainconsistency_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_Element_PropertyType
func (dq_element_propertytype *DQ_Element_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_TemporalValidity) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_TemporalConsistency) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_AccuracyOfATimeMeasurement) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_QuantitativeAttributeAccuracy) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_NonQuantitativeAttributeAccuracy) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_ThematicClassificationCorrectness) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_RelativeInternalPositionalAccuracy) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_GriddedDataPositionalAccuracy) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_AbsoluteExternalPositionalAccuracy) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_TopologicalConsistency) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_FormatConsistency) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_DomainConsistency) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_ConceptualConsistency) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_CompletenessOmission) || modified
	modified = stage.CleanPointer(&dq_element_propertytype.DQ_CompletenessCommission) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_FormatConsistency_PropertyType
func (dq_formatconsistency_propertytype *DQ_FormatConsistency_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_formatconsistency_propertytype.DQ_FormatConsistency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_FormatConsistency_Type
func (dq_formatconsistency_type *DQ_FormatConsistency_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_formatconsistency_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_formatconsistency_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_formatconsistency_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_formatconsistency_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_formatconsistency_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_formatconsistency_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_formatconsistency_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_formatconsistency_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_GriddedDataPositionalAccuracy_PropertyType
func (dq_griddeddatapositionalaccuracy_propertytype *DQ_GriddedDataPositionalAccuracy_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_griddeddatapositionalaccuracy_propertytype.DQ_GriddedDataPositionalAccuracy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_GriddedDataPositionalAccuracy_Type
func (dq_griddeddatapositionalaccuracy_type *DQ_GriddedDataPositionalAccuracy_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_griddeddatapositionalaccuracy_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_griddeddatapositionalaccuracy_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_griddeddatapositionalaccuracy_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_griddeddatapositionalaccuracy_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_griddeddatapositionalaccuracy_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_griddeddatapositionalaccuracy_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_griddeddatapositionalaccuracy_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_griddeddatapositionalaccuracy_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_LogicalConsistency_PropertyType
func (dq_logicalconsistency_propertytype *DQ_LogicalConsistency_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_logicalconsistency_propertytype.DQ_TopologicalConsistency) || modified
	modified = stage.CleanPointer(&dq_logicalconsistency_propertytype.DQ_FormatConsistency) || modified
	modified = stage.CleanPointer(&dq_logicalconsistency_propertytype.DQ_DomainConsistency) || modified
	modified = stage.CleanPointer(&dq_logicalconsistency_propertytype.DQ_ConceptualConsistency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_NonQuantitativeAttributeAccuracy_PropertyType
func (dq_nonquantitativeattributeaccuracy_propertytype *DQ_NonQuantitativeAttributeAccuracy_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_nonquantitativeattributeaccuracy_propertytype.DQ_NonQuantitativeAttributeAccuracy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_NonQuantitativeAttributeAccuracy_Type
func (dq_nonquantitativeattributeaccuracy_type *DQ_NonQuantitativeAttributeAccuracy_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_nonquantitativeattributeaccuracy_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_nonquantitativeattributeaccuracy_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_nonquantitativeattributeaccuracy_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_nonquantitativeattributeaccuracy_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_nonquantitativeattributeaccuracy_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_nonquantitativeattributeaccuracy_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_nonquantitativeattributeaccuracy_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_nonquantitativeattributeaccuracy_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_PositionalAccuracy_PropertyType
func (dq_positionalaccuracy_propertytype *DQ_PositionalAccuracy_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_positionalaccuracy_propertytype.DQ_RelativeInternalPositionalAccuracy) || modified
	modified = stage.CleanPointer(&dq_positionalaccuracy_propertytype.DQ_GriddedDataPositionalAccuracy) || modified
	modified = stage.CleanPointer(&dq_positionalaccuracy_propertytype.DQ_AbsoluteExternalPositionalAccuracy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_QuantitativeAttributeAccuracy_PropertyType
func (dq_quantitativeattributeaccuracy_propertytype *DQ_QuantitativeAttributeAccuracy_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_quantitativeattributeaccuracy_propertytype.DQ_QuantitativeAttributeAccuracy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_QuantitativeAttributeAccuracy_Type
func (dq_quantitativeattributeaccuracy_type *DQ_QuantitativeAttributeAccuracy_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_quantitativeattributeaccuracy_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_quantitativeattributeaccuracy_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_quantitativeattributeaccuracy_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_quantitativeattributeaccuracy_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_quantitativeattributeaccuracy_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_quantitativeattributeaccuracy_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_quantitativeattributeaccuracy_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_quantitativeattributeaccuracy_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_QuantitativeResult_PropertyType
func (dq_quantitativeresult_propertytype *DQ_QuantitativeResult_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_quantitativeresult_propertytype.DQ_QuantitativeResult) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_QuantitativeResult_Type
func (dq_quantitativeresult_type *DQ_QuantitativeResult_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_quantitativeresult_type.Value) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_quantitativeresult_type.ValueType) || modified
	modified = stage.CleanPointer(&dq_quantitativeresult_type.ValueUnit) || modified
	modified = stage.CleanPointer(&dq_quantitativeresult_type.ErrorStatistic) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_RelativeInternalPositionalAccuracy_PropertyType
func (dq_relativeinternalpositionalaccuracy_propertytype *DQ_RelativeInternalPositionalAccuracy_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_relativeinternalpositionalaccuracy_propertytype.DQ_RelativeInternalPositionalAccuracy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_RelativeInternalPositionalAccuracy_Type
func (dq_relativeinternalpositionalaccuracy_type *DQ_RelativeInternalPositionalAccuracy_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_relativeinternalpositionalaccuracy_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_relativeinternalpositionalaccuracy_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_relativeinternalpositionalaccuracy_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_relativeinternalpositionalaccuracy_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_relativeinternalpositionalaccuracy_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_relativeinternalpositionalaccuracy_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_relativeinternalpositionalaccuracy_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_relativeinternalpositionalaccuracy_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_Result_PropertyType
func (dq_result_propertytype *DQ_Result_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_result_propertytype.DQ_ConformanceResult) || modified
	modified = stage.CleanPointer(&dq_result_propertytype.DQ_QuantitativeResult) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_Scope_PropertyType
func (dq_scope_propertytype *DQ_Scope_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_scope_propertytype.DQ_Scope) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_Scope_Type
func (dq_scope_type *DQ_Scope_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_scope_type.LevelDescription) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_scope_type.Level) || modified
	modified = stage.CleanPointer(&dq_scope_type.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_TemporalAccuracy_PropertyType
func (dq_temporalaccuracy_propertytype *DQ_TemporalAccuracy_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_temporalaccuracy_propertytype.DQ_TemporalValidity) || modified
	modified = stage.CleanPointer(&dq_temporalaccuracy_propertytype.DQ_TemporalConsistency) || modified
	modified = stage.CleanPointer(&dq_temporalaccuracy_propertytype.DQ_AccuracyOfATimeMeasurement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_TemporalConsistency_PropertyType
func (dq_temporalconsistency_propertytype *DQ_TemporalConsistency_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_temporalconsistency_propertytype.DQ_TemporalConsistency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_TemporalConsistency_Type
func (dq_temporalconsistency_type *DQ_TemporalConsistency_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_temporalconsistency_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_temporalconsistency_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_temporalconsistency_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_temporalconsistency_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_temporalconsistency_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_temporalconsistency_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_temporalconsistency_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_temporalconsistency_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_TemporalValidity_PropertyType
func (dq_temporalvalidity_propertytype *DQ_TemporalValidity_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_temporalvalidity_propertytype.DQ_TemporalValidity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_TemporalValidity_Type
func (dq_temporalvalidity_type *DQ_TemporalValidity_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_temporalvalidity_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_temporalvalidity_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_temporalvalidity_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_temporalvalidity_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_temporalvalidity_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_temporalvalidity_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_temporalvalidity_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_temporalvalidity_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_ThematicAccuracy_PropertyType
func (dq_thematicaccuracy_propertytype *DQ_ThematicAccuracy_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_thematicaccuracy_propertytype.DQ_QuantitativeAttributeAccuracy) || modified
	modified = stage.CleanPointer(&dq_thematicaccuracy_propertytype.DQ_NonQuantitativeAttributeAccuracy) || modified
	modified = stage.CleanPointer(&dq_thematicaccuracy_propertytype.DQ_ThematicClassificationCorrectness) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_ThematicClassificationCorrectness_PropertyType
func (dq_thematicclassificationcorrectness_propertytype *DQ_ThematicClassificationCorrectness_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_thematicclassificationcorrectness_propertytype.DQ_ThematicClassificationCorrectness) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_ThematicClassificationCorrectness_Type
func (dq_thematicclassificationcorrectness_type *DQ_ThematicClassificationCorrectness_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_thematicclassificationcorrectness_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_thematicclassificationcorrectness_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_thematicclassificationcorrectness_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_thematicclassificationcorrectness_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_thematicclassificationcorrectness_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_thematicclassificationcorrectness_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_thematicclassificationcorrectness_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_thematicclassificationcorrectness_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_TopologicalConsistency_PropertyType
func (dq_topologicalconsistency_propertytype *DQ_TopologicalConsistency_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dq_topologicalconsistency_propertytype.DQ_TopologicalConsistency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DQ_TopologicalConsistency_Type
func (dq_topologicalconsistency_type *DQ_TopologicalConsistency_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dq_topologicalconsistency_type.NameOfMeasure) || modified
	modified = stage.CleanSlice(&dq_topologicalconsistency_type.DateTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dq_topologicalconsistency_type.MeasureIdentification) || modified
	modified = stage.CleanPointer(&dq_topologicalconsistency_type.MeasureDescription) || modified
	modified = stage.CleanPointer(&dq_topologicalconsistency_type.EvaluationMethodType) || modified
	modified = stage.CleanPointer(&dq_topologicalconsistency_type.EvaluationMethodDescription) || modified
	modified = stage.CleanPointer(&dq_topologicalconsistency_type.EvaluationProcedure) || modified
	modified = stage.CleanPointer(&dq_topologicalconsistency_type.Result) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DS_Association_PropertyType
func (ds_association_propertytype *DS_Association_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ds_association_propertytype.DS_Association) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataBlockType
func (datablocktype *DataBlockType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&datablocktype.RangeParameters) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DatumPropertyType
func (datumpropertytype *DatumPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&datumpropertytype.GeodeticDatum) || modified
	modified = stage.CleanPointer(&datumpropertytype.EngineeringDatum) || modified
	modified = stage.CleanPointer(&datumpropertytype.ImageDatum) || modified
	modified = stage.CleanPointer(&datumpropertytype.VerticalDatum) || modified
	modified = stage.CleanPointer(&datumpropertytype.TemporalDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DefinitionBaseType
func (definitionbasetype *DefinitionBaseType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&definitionbasetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&definitionbasetype.Description) || modified
	modified = stage.CleanPointer(&definitionbasetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DefinitionProxyType
func (definitionproxytype *DefinitionProxyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&definitionproxytype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&definitionproxytype.Description) || modified
	modified = stage.CleanPointer(&definitionproxytype.DescriptionReference) || modified
	modified = stage.CleanPointer(&definitionproxytype.DefinitionRef) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DefinitionType
func (definitiontype *DefinitionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&definitiontype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&definitiontype.Description) || modified
	modified = stage.CleanPointer(&definitiontype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DerivedCRSPropertyType
func (derivedcrspropertytype *DerivedCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&derivedcrspropertytype.DerivedCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DerivedCRSType
func (derivedcrstype *DerivedCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&derivedcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&derivedcrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&derivedcrstype.Description) || modified
	modified = stage.CleanPointer(&derivedcrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&derivedcrstype.Conversion) || modified
	modified = stage.CleanPointer(&derivedcrstype.DefinedByConversion) || modified
	modified = stage.CleanPointer(&derivedcrstype.BaseCRS) || modified
	modified = stage.CleanPointer(&derivedcrstype.CoordinateSystem) || modified
	modified = stage.CleanPointer(&derivedcrstype.UsesCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DerivedUnitType
func (derivedunittype *DerivedUnitType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&derivedunittype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&derivedunittype.DerivationUnitTerm) || modified
	// insertion point per field
	modified = stage.CleanPointer(&derivedunittype.Description) || modified
	modified = stage.CleanPointer(&derivedunittype.DescriptionReference) || modified
	modified = stage.CleanPointer(&derivedunittype.QuantityType) || modified
	modified = stage.CleanPointer(&derivedunittype.QuantityTypeReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DictionaryEntryType
func (dictionaryentrytype *DictionaryEntryType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dictionaryentrytype.Definition) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.BaseUnit) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.ConventionalUnit) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.Dictionary) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.CoordinateSystemAxis) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.EllipsoidalCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.CartesianCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.VerticalCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TimeCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.LinearCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.UserDefinedCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.SphericalCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.PolarCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.CylindricalCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.AffineCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TemporalCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.ObliqueCartesianCS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.ProjectedCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.DerivedCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.GeodeticCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.VerticalCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.EngineeringCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.ImageCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TemporalCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.GeographicCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.GeocentricCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.CompoundCRS) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.GeodeticDatum) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.EngineeringDatum) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.ImageDatum) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.VerticalDatum) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TemporalDatum) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.Ellipsoid) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.PrimeMeridian) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.PassThroughOperation) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.Conversion) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.Transformation) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.ConcatenatedOperation) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.OperationMethod) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.OperationParameter) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.OperationParameterGroup) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TimeReferenceSystem) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TimeCoordinateSystem) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TimeCalendar) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TimeClock) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.TimeOrdinalReferenceSystem) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.DefinitionCollection) || modified
	modified = stage.CleanPointer(&dictionaryentrytype.DefinitionProxy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DictionaryType
func (dictionarytype *DictionaryType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dictionarytype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&dictionarytype.DictionaryEntry) || modified
	modified = stage.CleanSlice(&dictionarytype.DefinitionMember) || modified
	modified = stage.CleanSlice(&dictionarytype.IndirectEntry) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dictionarytype.Description) || modified
	modified = stage.CleanPointer(&dictionarytype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectedEdgePropertyType
func (directededgepropertytype *DirectedEdgePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directededgepropertytype.Edge) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectedFacePropertyType
func (directedfacepropertytype *DirectedFacePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directedfacepropertytype.Face) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectedNodePropertyType
func (directednodepropertytype *DirectedNodePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directednodepropertytype.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectedObservationAtDistanceType
func (directedobservationatdistancetype *DirectedObservationAtDistanceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&directedobservationatdistancetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&directedobservationatdistancetype.Description) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.BoundedBy) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.Location) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.PriorityLocation) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.ValidTime) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.Using) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.Target) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.Subject) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.ResultOf) || modified
	modified = stage.CleanPointer(&directedobservationatdistancetype.Direction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectedObservationType
func (directedobservationtype *DirectedObservationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&directedobservationtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&directedobservationtype.Description) || modified
	modified = stage.CleanPointer(&directedobservationtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&directedobservationtype.BoundedBy) || modified
	modified = stage.CleanPointer(&directedobservationtype.Location) || modified
	modified = stage.CleanPointer(&directedobservationtype.PriorityLocation) || modified
	modified = stage.CleanPointer(&directedobservationtype.ValidTime) || modified
	modified = stage.CleanPointer(&directedobservationtype.Using) || modified
	modified = stage.CleanPointer(&directedobservationtype.Target) || modified
	modified = stage.CleanPointer(&directedobservationtype.Subject) || modified
	modified = stage.CleanPointer(&directedobservationtype.ResultOf) || modified
	modified = stage.CleanPointer(&directedobservationtype.Direction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectedTopoSolidPropertyType
func (directedtoposolidpropertytype *DirectedTopoSolidPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directedtoposolidpropertytype.TopoSolid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectionDescriptionType
func (directiondescriptiontype *DirectionDescriptionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directiondescriptiontype.Reference_XSD) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectionPropertyType
func (directionpropertytype *DirectionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directionpropertytype.DirectionVector) || modified
	modified = stage.CleanPointer(&directionpropertytype.DirectionDescription) || modified
	modified = stage.CleanPointer(&directionpropertytype.DirectionString) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DirectionVectorType
func (directionvectortype *DirectionVectorType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&directionvectortype.Vector) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DiscreteCoverageType
func (discretecoveragetype *DiscreteCoverageType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&discretecoveragetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&discretecoveragetype.Description) || modified
	modified = stage.CleanPointer(&discretecoveragetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&discretecoveragetype.BoundedBy) || modified
	modified = stage.CleanPointer(&discretecoveragetype.Location) || modified
	modified = stage.CleanPointer(&discretecoveragetype.PriorityLocation) || modified
	modified = stage.CleanPointer(&discretecoveragetype.DomainSet) || modified
	modified = stage.CleanPointer(&discretecoveragetype.MultiPointDomain) || modified
	modified = stage.CleanPointer(&discretecoveragetype.MultiCurveDomain) || modified
	modified = stage.CleanPointer(&discretecoveragetype.MultiSurfaceDomain) || modified
	modified = stage.CleanPointer(&discretecoveragetype.MultiSolidDomain) || modified
	modified = stage.CleanPointer(&discretecoveragetype.GridDomain) || modified
	modified = stage.CleanPointer(&discretecoveragetype.RectifiedGridDomain) || modified
	modified = stage.CleanPointer(&discretecoveragetype.RangeSet) || modified
	modified = stage.CleanPointer(&discretecoveragetype.CoverageFunction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DomainOfValidity
func (domainofvalidity *DomainOfValidity) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&domainofvalidity.EX_Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DomainSetType
func (domainsettype *DomainSetType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&domainsettype.MultiGeometry) || modified
	modified = stage.CleanPointer(&domainsettype.MultiPoint) || modified
	modified = stage.CleanPointer(&domainsettype.MultiCurve) || modified
	modified = stage.CleanPointer(&domainsettype.MultiSurface) || modified
	modified = stage.CleanPointer(&domainsettype.MultiSolid) || modified
	modified = stage.CleanPointer(&domainsettype.Solid) || modified
	modified = stage.CleanPointer(&domainsettype.CompositeSolid) || modified
	modified = stage.CleanPointer(&domainsettype.Surface) || modified
	modified = stage.CleanPointer(&domainsettype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&domainsettype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&domainsettype.Tin) || modified
	modified = stage.CleanPointer(&domainsettype.OrientableSurface) || modified
	modified = stage.CleanPointer(&domainsettype.Shell) || modified
	modified = stage.CleanPointer(&domainsettype.Polygon) || modified
	modified = stage.CleanPointer(&domainsettype.CompositeSurface) || modified
	modified = stage.CleanPointer(&domainsettype.Point) || modified
	modified = stage.CleanPointer(&domainsettype.Curve) || modified
	modified = stage.CleanPointer(&domainsettype.OrientableCurve) || modified
	modified = stage.CleanPointer(&domainsettype.Ring) || modified
	modified = stage.CleanPointer(&domainsettype.LinearRing) || modified
	modified = stage.CleanPointer(&domainsettype.LineString) || modified
	modified = stage.CleanPointer(&domainsettype.CompositeCurve) || modified
	modified = stage.CleanPointer(&domainsettype.GeometricComplex) || modified
	modified = stage.CleanPointer(&domainsettype.Grid) || modified
	modified = stage.CleanPointer(&domainsettype.RectifiedGrid) || modified
	modified = stage.CleanPointer(&domainsettype.TimeInstant) || modified
	modified = stage.CleanPointer(&domainsettype.TimePeriod) || modified
	modified = stage.CleanPointer(&domainsettype.TimeNode) || modified
	modified = stage.CleanPointer(&domainsettype.TimeEdge) || modified
	modified = stage.CleanPointer(&domainsettype.TimeTopologyComplex) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DynamicFeatureCollectionType
func (dynamicfeaturecollectiontype *DynamicFeatureCollectionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dynamicfeaturecollectiontype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.Description) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.DescriptionReference) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.BoundedBy) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.Location) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.PriorityLocation) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.ValidTime) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.History) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.Track) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.DataSource) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.DataSourceReference) || modified
	modified = stage.CleanPointer(&dynamicfeaturecollectiontype.DynamicMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DynamicFeatureMemberType
func (dynamicfeaturemembertype *DynamicFeatureMemberType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dynamicfeaturemembertype.DynamicFeature) || modified
	modified = stage.CleanSlice(&dynamicfeaturemembertype.DynamicFeatureCollection) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DynamicFeatureType
func (dynamicfeaturetype *DynamicFeatureType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dynamicfeaturetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dynamicfeaturetype.Description) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.BoundedBy) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.Location) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.PriorityLocation) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.ValidTime) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.History) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.Track) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.DataSource) || modified
	modified = stage.CleanPointer(&dynamicfeaturetype.DataSourceReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_BoundingPolygon_PropertyType
func (ex_boundingpolygon_propertytype *EX_BoundingPolygon_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_boundingpolygon_propertytype.EX_BoundingPolygon) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_BoundingPolygon_Type
func (ex_boundingpolygon_type *EX_BoundingPolygon_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ex_boundingpolygon_type.Polygon) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ex_boundingpolygon_type.ExtentTypeCode) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_Extent_PropertyType
func (ex_extent_propertytype *EX_Extent_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_extent_propertytype.EX_Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_Extent_Type
func (ex_extent_type *EX_Extent_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ex_extent_type.GeographicElement) || modified
	modified = stage.CleanSlice(&ex_extent_type.TemporalElement) || modified
	modified = stage.CleanSlice(&ex_extent_type.VerticalElement) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ex_extent_type.Description) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_GeographicBoundingBox_PropertyType
func (ex_geographicboundingbox_propertytype *EX_GeographicBoundingBox_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_geographicboundingbox_propertytype.EX_GeographicBoundingBox) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_GeographicBoundingBox_Type
func (ex_geographicboundingbox_type *EX_GeographicBoundingBox_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_geographicboundingbox_type.ExtentTypeCode) || modified
	modified = stage.CleanPointer(&ex_geographicboundingbox_type.WestBoundLongitude) || modified
	modified = stage.CleanPointer(&ex_geographicboundingbox_type.EastBoundLongitude) || modified
	modified = stage.CleanPointer(&ex_geographicboundingbox_type.SouthBoundLatitude) || modified
	modified = stage.CleanPointer(&ex_geographicboundingbox_type.NorthBoundLatitude) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_GeographicDescription_PropertyType
func (ex_geographicdescription_propertytype *EX_GeographicDescription_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_geographicdescription_propertytype.EX_GeographicDescription) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_GeographicDescription_Type
func (ex_geographicdescription_type *EX_GeographicDescription_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_geographicdescription_type.ExtentTypeCode) || modified
	modified = stage.CleanPointer(&ex_geographicdescription_type.GeographicIdentifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_GeographicExtent_PropertyType
func (ex_geographicextent_propertytype *EX_GeographicExtent_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_geographicextent_propertytype.EX_BoundingPolygon) || modified
	modified = stage.CleanPointer(&ex_geographicextent_propertytype.EX_GeographicBoundingBox) || modified
	modified = stage.CleanPointer(&ex_geographicextent_propertytype.EX_GeographicDescription) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_SpatialTemporalExtent_PropertyType
func (ex_spatialtemporalextent_propertytype *EX_SpatialTemporalExtent_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_spatialtemporalextent_propertytype.EX_SpatialTemporalExtent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_SpatialTemporalExtent_Type
func (ex_spatialtemporalextent_type *EX_SpatialTemporalExtent_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ex_spatialtemporalextent_type.SpatialExtent) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ex_spatialtemporalextent_type.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_TemporalExtent_PropertyType
func (ex_temporalextent_propertytype *EX_TemporalExtent_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_temporalextent_propertytype.EX_TemporalExtent) || modified
	modified = stage.CleanPointer(&ex_temporalextent_propertytype.EX_SpatialTemporalExtent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_TemporalExtent_Type
func (ex_temporalextent_type *EX_TemporalExtent_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_temporalextent_type.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_VerticalExtent_PropertyType
func (ex_verticalextent_propertytype *EX_VerticalExtent_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_verticalextent_propertytype.EX_VerticalExtent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EX_VerticalExtent_Type
func (ex_verticalextent_type *EX_VerticalExtent_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ex_verticalextent_type.MinimumValue) || modified
	modified = stage.CleanPointer(&ex_verticalextent_type.MaximumValue) || modified
	modified = stage.CleanPointer(&ex_verticalextent_type.VerticalCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EdgeType
func (edgetype *EdgeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&edgetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&edgetype.DirectedFace) || modified
	// insertion point per field
	modified = stage.CleanPointer(&edgetype.Description) || modified
	modified = stage.CleanPointer(&edgetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&edgetype.Container) || modified
	modified = stage.CleanPointer(&edgetype.DirectedNode) || modified
	modified = stage.CleanPointer(&edgetype.CurveProperty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EllipsoidPropertyType
func (ellipsoidpropertytype *EllipsoidPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ellipsoidpropertytype.Ellipsoid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EllipsoidType
func (ellipsoidtype *EllipsoidType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ellipsoidtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ellipsoidtype.Description) || modified
	modified = stage.CleanPointer(&ellipsoidtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&ellipsoidtype.SecondDefiningParameter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EllipsoidalCSPropertyType
func (ellipsoidalcspropertytype *EllipsoidalCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ellipsoidalcspropertytype.EllipsoidalCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EllipsoidalCSType
func (ellipsoidalcstype *EllipsoidalCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ellipsoidalcstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&ellipsoidalcstype.Axis) || modified
	modified = stage.CleanSlice(&ellipsoidalcstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ellipsoidalcstype.Description) || modified
	modified = stage.CleanPointer(&ellipsoidalcstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EngineeringCRSPropertyType
func (engineeringcrspropertytype *EngineeringCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&engineeringcrspropertytype.EngineeringCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EngineeringCRSType
func (engineeringcrstype *EngineeringCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&engineeringcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&engineeringcrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&engineeringcrstype.Description) || modified
	modified = stage.CleanPointer(&engineeringcrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&engineeringcrstype.AffineCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.UsesAffineCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.CartesianCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.UsesCartesianCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.CylindricalCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.LinearCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.PolarCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.SphericalCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.UsesSphericalCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.UserDefinedCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.CoordinateSystem) || modified
	modified = stage.CleanPointer(&engineeringcrstype.UsesCS) || modified
	modified = stage.CleanPointer(&engineeringcrstype.EngineeringDatum) || modified
	modified = stage.CleanPointer(&engineeringcrstype.UsesEngineeringDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EngineeringDatumPropertyType
func (engineeringdatumpropertytype *EngineeringDatumPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&engineeringdatumpropertytype.EngineeringDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EngineeringDatumType
func (engineeringdatumtype *EngineeringDatumType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&engineeringdatumtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&engineeringdatumtype.Description) || modified
	modified = stage.CleanPointer(&engineeringdatumtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&engineeringdatumtype.DomainOfValidity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EnvelopeType
func (envelopetype *EnvelopeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&envelopetype.LowerCorner) || modified
	modified = stage.CleanPointer(&envelopetype.UpperCorner) || modified
	modified = stage.CleanPointer(&envelopetype.Pos) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EnvelopeWithTimePeriodType
func (envelopewithtimeperiodtype *EnvelopeWithTimePeriodType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&envelopewithtimeperiodtype.LowerCorner) || modified
	modified = stage.CleanPointer(&envelopewithtimeperiodtype.UpperCorner) || modified
	modified = stage.CleanPointer(&envelopewithtimeperiodtype.Pos) || modified
	modified = stage.CleanPointer(&envelopewithtimeperiodtype.BeginPosition) || modified
	modified = stage.CleanPointer(&envelopewithtimeperiodtype.EndPosition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FaceOrTopoSolidPropertyType
func (faceortoposolidpropertytype *FaceOrTopoSolidPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&faceortoposolidpropertytype.Face) || modified
	modified = stage.CleanPointer(&faceortoposolidpropertytype.TopoSolid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FaceType
func (facetype *FaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&facetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&facetype.Isolated) || modified
	modified = stage.CleanSlice(&facetype.DirectedEdge) || modified
	// insertion point per field
	modified = stage.CleanPointer(&facetype.Description) || modified
	modified = stage.CleanPointer(&facetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&facetype.DirectedTopoSolid) || modified
	modified = stage.CleanPointer(&facetype.SurfaceProperty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FeatureArrayPropertyType
func (featurearraypropertytype *FeatureArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&featurearraypropertytype.DynamicFeature) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.DynamicFeatureCollection) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.MultiPointCoverage) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.MultiCurveCoverage) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.MultiSurfaceCoverage) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.MultiSolidCoverage) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.GridCoverage) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.RectifiedGridCoverage) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.Observation) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.DirectedObservation) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.DirectedObservationAtDistance) || modified
	modified = stage.CleanSlice(&featurearraypropertytype.FeatureCollection) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by FeatureCollectionType
func (featurecollectiontype *FeatureCollectionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&featurecollectiontype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&featurecollectiontype.FeatureMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&featurecollectiontype.Description) || modified
	modified = stage.CleanPointer(&featurecollectiontype.DescriptionReference) || modified
	modified = stage.CleanPointer(&featurecollectiontype.BoundedBy) || modified
	modified = stage.CleanPointer(&featurecollectiontype.Location) || modified
	modified = stage.CleanPointer(&featurecollectiontype.PriorityLocation) || modified
	modified = stage.CleanPointer(&featurecollectiontype.FeatureMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FeaturePropertyType
func (featurepropertytype *FeaturePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&featurepropertytype.DynamicFeature) || modified
	modified = stage.CleanPointer(&featurepropertytype.DynamicFeatureCollection) || modified
	modified = stage.CleanPointer(&featurepropertytype.MultiPointCoverage) || modified
	modified = stage.CleanPointer(&featurepropertytype.MultiCurveCoverage) || modified
	modified = stage.CleanPointer(&featurepropertytype.MultiSurfaceCoverage) || modified
	modified = stage.CleanPointer(&featurepropertytype.MultiSolidCoverage) || modified
	modified = stage.CleanPointer(&featurepropertytype.GridCoverage) || modified
	modified = stage.CleanPointer(&featurepropertytype.RectifiedGridCoverage) || modified
	modified = stage.CleanPointer(&featurepropertytype.Observation) || modified
	modified = stage.CleanPointer(&featurepropertytype.DirectedObservation) || modified
	modified = stage.CleanPointer(&featurepropertytype.DirectedObservationAtDistance) || modified
	modified = stage.CleanPointer(&featurepropertytype.FeatureCollection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FileType
func (filetype *FileType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&filetype.RangeParameters) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by FormulaCitation
func (formulacitation *FormulaCitation) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&formulacitation.CI_Citation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GM_Object_PropertyType
func (gm_object_propertytype *GM_Object_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&gm_object_propertytype.MultiGeometry) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.MultiPoint) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.MultiCurve) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.MultiSurface) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.MultiSolid) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Solid) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.CompositeSolid) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Surface) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Tin) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Shell) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Polygon) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Point) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Curve) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Ring) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.LinearRing) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.LineString) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.CompositeCurve) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.GeometricComplex) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.Grid) || modified
	modified = stage.CleanPointer(&gm_object_propertytype.RectifiedGrid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GM_Point_PropertyType
func (gm_point_propertytype *GM_Point_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&gm_point_propertytype.Point) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeneralConversionPropertyType
func (generalconversionpropertytype *GeneralConversionPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&generalconversionpropertytype.Conversion) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeneralTransformationPropertyType
func (generaltransformationpropertytype *GeneralTransformationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&generaltransformationpropertytype.Transformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeocentricCRSPropertyType
func (geocentriccrspropertytype *GeocentricCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geocentriccrspropertytype.GeocentricCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeocentricCRSType
func (geocentriccrstype *GeocentricCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&geocentriccrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&geocentriccrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&geocentriccrstype.Description) || modified
	modified = stage.CleanPointer(&geocentriccrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&geocentriccrstype.UsesCartesianCS) || modified
	modified = stage.CleanPointer(&geocentriccrstype.UsesSphericalCS) || modified
	modified = stage.CleanPointer(&geocentriccrstype.UsesGeodeticDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeodesicStringType
func (geodesicstringtype *GeodesicStringType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geodesicstringtype.PosList) || modified
	modified = stage.CleanPointer(&geodesicstringtype.Pos) || modified
	modified = stage.CleanPointer(&geodesicstringtype.PointProperty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeodesicType
func (geodesictype *GeodesicType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geodesictype.PosList) || modified
	modified = stage.CleanPointer(&geodesictype.Pos) || modified
	modified = stage.CleanPointer(&geodesictype.PointProperty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeodeticCRSPropertyType
func (geodeticcrspropertytype *GeodeticCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geodeticcrspropertytype.GeodeticCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeodeticCRSType
func (geodeticcrstype *GeodeticCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&geodeticcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&geodeticcrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&geodeticcrstype.Description) || modified
	modified = stage.CleanPointer(&geodeticcrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&geodeticcrstype.EllipsoidalCS) || modified
	modified = stage.CleanPointer(&geodeticcrstype.UsesEllipsoidalCS) || modified
	modified = stage.CleanPointer(&geodeticcrstype.CartesianCS) || modified
	modified = stage.CleanPointer(&geodeticcrstype.UsesCartesianCS) || modified
	modified = stage.CleanPointer(&geodeticcrstype.SphericalCS) || modified
	modified = stage.CleanPointer(&geodeticcrstype.UsesSphericalCS) || modified
	modified = stage.CleanPointer(&geodeticcrstype.GeodeticDatum) || modified
	modified = stage.CleanPointer(&geodeticcrstype.UsesGeodeticDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeodeticDatumPropertyType
func (geodeticdatumpropertytype *GeodeticDatumPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geodeticdatumpropertytype.GeodeticDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeodeticDatumType
func (geodeticdatumtype *GeodeticDatumType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&geodeticdatumtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&geodeticdatumtype.Description) || modified
	modified = stage.CleanPointer(&geodeticdatumtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&geodeticdatumtype.DomainOfValidity) || modified
	modified = stage.CleanPointer(&geodeticdatumtype.PrimeMeridian) || modified
	modified = stage.CleanPointer(&geodeticdatumtype.UsesPrimeMeridian) || modified
	modified = stage.CleanPointer(&geodeticdatumtype.Ellipsoid) || modified
	modified = stage.CleanPointer(&geodeticdatumtype.UsesEllipsoid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeographicCRSPropertyType
func (geographiccrspropertytype *GeographicCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geographiccrspropertytype.GeographicCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeographicCRSType
func (geographiccrstype *GeographicCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&geographiccrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&geographiccrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&geographiccrstype.Description) || modified
	modified = stage.CleanPointer(&geographiccrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&geographiccrstype.UsesEllipsoidalCS) || modified
	modified = stage.CleanPointer(&geographiccrstype.UsesGeodeticDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeometricComplexPropertyType
func (geometriccomplexpropertytype *GeometricComplexPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geometriccomplexpropertytype.GeometricComplex) || modified
	modified = stage.CleanPointer(&geometriccomplexpropertytype.CompositeCurve) || modified
	modified = stage.CleanPointer(&geometriccomplexpropertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&geometriccomplexpropertytype.CompositeSolid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeometricComplexType
func (geometriccomplextype *GeometricComplexType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&geometriccomplextype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&geometriccomplextype.Element) || modified
	// insertion point per field
	modified = stage.CleanPointer(&geometriccomplextype.Description) || modified
	modified = stage.CleanPointer(&geometriccomplextype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeometricPrimitivePropertyType
func (geometricprimitivepropertytype *GeometricPrimitivePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geometricprimitivepropertytype.Solid) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.CompositeSolid) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.Surface) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.Tin) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.Shell) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.Polygon) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.Point) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.Curve) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.Ring) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.LinearRing) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.LineString) || modified
	modified = stage.CleanPointer(&geometricprimitivepropertytype.CompositeCurve) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GeometryArrayPropertyType
func (geometryarraypropertytype *GeometryArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&geometryarraypropertytype.MultiGeometry) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.MultiPoint) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.MultiCurve) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.MultiSurface) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.MultiSolid) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Solid) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.CompositeSolid) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Surface) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.PolyhedralSurface) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.TriangulatedSurface) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Tin) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.OrientableSurface) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Shell) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Polygon) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.CompositeSurface) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Point) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Curve) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.OrientableCurve) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Ring) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.LinearRing) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.LineString) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.CompositeCurve) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.GeometricComplex) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.Grid) || modified
	modified = stage.CleanSlice(&geometryarraypropertytype.RectifiedGrid) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GeometryPropertyType
func (geometrypropertytype *GeometryPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&geometrypropertytype.MultiGeometry) || modified
	modified = stage.CleanPointer(&geometrypropertytype.MultiPoint) || modified
	modified = stage.CleanPointer(&geometrypropertytype.MultiCurve) || modified
	modified = stage.CleanPointer(&geometrypropertytype.MultiSurface) || modified
	modified = stage.CleanPointer(&geometrypropertytype.MultiSolid) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Solid) || modified
	modified = stage.CleanPointer(&geometrypropertytype.CompositeSolid) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Surface) || modified
	modified = stage.CleanPointer(&geometrypropertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&geometrypropertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Tin) || modified
	modified = stage.CleanPointer(&geometrypropertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Shell) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Polygon) || modified
	modified = stage.CleanPointer(&geometrypropertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Point) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Curve) || modified
	modified = stage.CleanPointer(&geometrypropertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Ring) || modified
	modified = stage.CleanPointer(&geometrypropertytype.LinearRing) || modified
	modified = stage.CleanPointer(&geometrypropertytype.LineString) || modified
	modified = stage.CleanPointer(&geometrypropertytype.CompositeCurve) || modified
	modified = stage.CleanPointer(&geometrypropertytype.GeometricComplex) || modified
	modified = stage.CleanPointer(&geometrypropertytype.Grid) || modified
	modified = stage.CleanPointer(&geometrypropertytype.RectifiedGrid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GridFunctionType
func (gridfunctiontype *GridFunctionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&gridfunctiontype.SequenceRule) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GridLimitsType
func (gridlimitstype *GridLimitsType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&gridlimitstype.GridEnvelope) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GridType
func (gridtype *GridType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&gridtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&gridtype.Description) || modified
	modified = stage.CleanPointer(&gridtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&gridtype.Limits) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by HistoryPropertyType
func (historypropertytype *HistoryPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&historypropertytype.MovingObjectStatus) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by IdentifiedObjectType
func (identifiedobjecttype *IdentifiedObjectType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&identifiedobjecttype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&identifiedobjecttype.Description) || modified
	modified = stage.CleanPointer(&identifiedobjecttype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ImageCRSPropertyType
func (imagecrspropertytype *ImageCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&imagecrspropertytype.ImageCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ImageCRSType
func (imagecrstype *ImageCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&imagecrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&imagecrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&imagecrstype.Description) || modified
	modified = stage.CleanPointer(&imagecrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&imagecrstype.CartesianCS) || modified
	modified = stage.CleanPointer(&imagecrstype.UsesCartesianCS) || modified
	modified = stage.CleanPointer(&imagecrstype.AffineCS) || modified
	modified = stage.CleanPointer(&imagecrstype.UsesAffineCS) || modified
	modified = stage.CleanPointer(&imagecrstype.UsesObliqueCartesianCS) || modified
	modified = stage.CleanPointer(&imagecrstype.ImageDatum) || modified
	modified = stage.CleanPointer(&imagecrstype.UsesImageDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ImageDatumPropertyType
func (imagedatumpropertytype *ImageDatumPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&imagedatumpropertytype.ImageDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ImageDatumType
func (imagedatumtype *ImageDatumType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&imagedatumtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&imagedatumtype.Description) || modified
	modified = stage.CleanPointer(&imagedatumtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&imagedatumtype.DomainOfValidity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by IndirectEntryType
func (indirectentrytype *IndirectEntryType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&indirectentrytype.DefinitionProxy) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by KnotPropertyType
func (knotpropertytype *KnotPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&knotpropertytype.Knot) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LI_Lineage_PropertyType
func (li_lineage_propertytype *LI_Lineage_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&li_lineage_propertytype.LI_Lineage) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LI_Lineage_Type
func (li_lineage_type *LI_Lineage_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&li_lineage_type.ProcessStep) || modified
	modified = stage.CleanSlice(&li_lineage_type.Source) || modified
	// insertion point per field
	modified = stage.CleanPointer(&li_lineage_type.Statement) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LI_ProcessStep_PropertyType
func (li_processstep_propertytype *LI_ProcessStep_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&li_processstep_propertytype.LI_ProcessStep) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LI_ProcessStep_Type
func (li_processstep_type *LI_ProcessStep_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&li_processstep_type.Processor) || modified
	modified = stage.CleanSlice(&li_processstep_type.Source) || modified
	// insertion point per field
	modified = stage.CleanPointer(&li_processstep_type.Description) || modified
	modified = stage.CleanPointer(&li_processstep_type.Rationale) || modified
	modified = stage.CleanPointer(&li_processstep_type.DateTime) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LI_Source_PropertyType
func (li_source_propertytype *LI_Source_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&li_source_propertytype.LI_Source) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LI_Source_Type
func (li_source_type *LI_Source_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&li_source_type.SourceExtent) || modified
	modified = stage.CleanSlice(&li_source_type.SourceStep) || modified
	// insertion point per field
	modified = stage.CleanPointer(&li_source_type.Description) || modified
	modified = stage.CleanPointer(&li_source_type.ScaleDenominator) || modified
	modified = stage.CleanPointer(&li_source_type.SourceReferenceSystem) || modified
	modified = stage.CleanPointer(&li_source_type.SourceCitation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LineStringSegmentArrayPropertyType
func (linestringsegmentarraypropertytype *LineStringSegmentArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&linestringsegmentarraypropertytype.LineStringSegment) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by LineStringSegmentType
func (linestringsegmenttype *LineStringSegmentType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&linestringsegmenttype.Pos) || modified
	modified = stage.CleanSlice(&linestringsegmenttype.PointProperty) || modified
	modified = stage.CleanSlice(&linestringsegmenttype.PointRep) || modified
	// insertion point per field
	modified = stage.CleanPointer(&linestringsegmenttype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LineStringType
func (linestringtype *LineStringType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&linestringtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&linestringtype.Pos) || modified
	modified = stage.CleanSlice(&linestringtype.PointProperty) || modified
	modified = stage.CleanSlice(&linestringtype.PointRep) || modified
	// insertion point per field
	modified = stage.CleanPointer(&linestringtype.Description) || modified
	modified = stage.CleanPointer(&linestringtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&linestringtype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LinearCSPropertyType
func (linearcspropertytype *LinearCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&linearcspropertytype.LinearCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LinearCSType
func (linearcstype *LinearCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&linearcstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&linearcstype.Axis) || modified
	modified = stage.CleanSlice(&linearcstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&linearcstype.Description) || modified
	modified = stage.CleanPointer(&linearcstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LinearRingPropertyType
func (linearringpropertytype *LinearRingPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&linearringpropertytype.LinearRing) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LinearRingType
func (linearringtype *LinearRingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&linearringtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&linearringtype.Pos) || modified
	modified = stage.CleanSlice(&linearringtype.PointProperty) || modified
	modified = stage.CleanSlice(&linearringtype.PointRep) || modified
	// insertion point per field
	modified = stage.CleanPointer(&linearringtype.Description) || modified
	modified = stage.CleanPointer(&linearringtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&linearringtype.PosList) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LocalisedCharacterString_PropertyType
func (localisedcharacterstring_propertytype *LocalisedCharacterString_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&localisedcharacterstring_propertytype.LocalisedCharacterString) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LocationPropertyType
func (locationpropertytype *LocationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&locationpropertytype.MultiGeometry) || modified
	modified = stage.CleanPointer(&locationpropertytype.MultiPoint) || modified
	modified = stage.CleanPointer(&locationpropertytype.MultiCurve) || modified
	modified = stage.CleanPointer(&locationpropertytype.MultiSurface) || modified
	modified = stage.CleanPointer(&locationpropertytype.MultiSolid) || modified
	modified = stage.CleanPointer(&locationpropertytype.Solid) || modified
	modified = stage.CleanPointer(&locationpropertytype.CompositeSolid) || modified
	modified = stage.CleanPointer(&locationpropertytype.Surface) || modified
	modified = stage.CleanPointer(&locationpropertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&locationpropertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&locationpropertytype.Tin) || modified
	modified = stage.CleanPointer(&locationpropertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&locationpropertytype.Shell) || modified
	modified = stage.CleanPointer(&locationpropertytype.Polygon) || modified
	modified = stage.CleanPointer(&locationpropertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&locationpropertytype.Point) || modified
	modified = stage.CleanPointer(&locationpropertytype.Curve) || modified
	modified = stage.CleanPointer(&locationpropertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&locationpropertytype.Ring) || modified
	modified = stage.CleanPointer(&locationpropertytype.LinearRing) || modified
	modified = stage.CleanPointer(&locationpropertytype.LineString) || modified
	modified = stage.CleanPointer(&locationpropertytype.CompositeCurve) || modified
	modified = stage.CleanPointer(&locationpropertytype.GeometricComplex) || modified
	modified = stage.CleanPointer(&locationpropertytype.Grid) || modified
	modified = stage.CleanPointer(&locationpropertytype.RectifiedGrid) || modified
	modified = stage.CleanPointer(&locationpropertytype.LocationString) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_AggregateInformation_PropertyType
func (md_aggregateinformation_propertytype *MD_AggregateInformation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_aggregateinformation_propertytype.MD_AggregateInformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_AggregateInformation_Type
func (md_aggregateinformation_type *MD_AggregateInformation_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_aggregateinformation_type.AggregateDataSetName) || modified
	modified = stage.CleanPointer(&md_aggregateinformation_type.AggregateDataSetIdentifier) || modified
	modified = stage.CleanPointer(&md_aggregateinformation_type.AssociationType) || modified
	modified = stage.CleanPointer(&md_aggregateinformation_type.InitiativeType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_BrowseGraphic_PropertyType
func (md_browsegraphic_propertytype *MD_BrowseGraphic_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_browsegraphic_propertytype.MD_BrowseGraphic) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_BrowseGraphic_Type
func (md_browsegraphic_type *MD_BrowseGraphic_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_browsegraphic_type.FileName) || modified
	modified = stage.CleanPointer(&md_browsegraphic_type.FileDescription) || modified
	modified = stage.CleanPointer(&md_browsegraphic_type.FileType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Constraints_PropertyType
func (md_constraints_propertytype *MD_Constraints_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_constraints_propertytype.MD_Constraints) || modified
	modified = stage.CleanPointer(&md_constraints_propertytype.MD_LegalConstraints) || modified
	modified = stage.CleanPointer(&md_constraints_propertytype.MD_SecurityConstraints) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Constraints_Type
func (md_constraints_type *MD_Constraints_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_constraints_type.UseLimitation) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_DataIdentification_PropertyType
func (md_dataidentification_propertytype *MD_DataIdentification_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_dataidentification_propertytype.MD_DataIdentification) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_DataIdentification_Type
func (md_dataidentification_type *MD_DataIdentification_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_dataidentification_type.Credit) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.Status) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.PointOfContact) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.ResourceMaintenance) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.GraphicOverview) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.ResourceFormat) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.DescriptiveKeywords) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.ResourceSpecificUsage) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.ResourceConstraints) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.AggregationInfo) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.SpatialRepresentationType) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.SpatialResolution) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.Language) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.CharacterSet) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.TopicCategory) || modified
	modified = stage.CleanSlice(&md_dataidentification_type.Extent) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_dataidentification_type.Citation) || modified
	modified = stage.CleanPointer(&md_dataidentification_type.Abstract) || modified
	modified = stage.CleanPointer(&md_dataidentification_type.Purpose) || modified
	modified = stage.CleanPointer(&md_dataidentification_type.EnvironmentDescription) || modified
	modified = stage.CleanPointer(&md_dataidentification_type.SupplementalInformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_DigitalTransferOptions_PropertyType
func (md_digitaltransferoptions_propertytype *MD_DigitalTransferOptions_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_digitaltransferoptions_propertytype.MD_DigitalTransferOptions) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_DigitalTransferOptions_Type
func (md_digitaltransferoptions_type *MD_DigitalTransferOptions_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_digitaltransferoptions_type.OnLine) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_digitaltransferoptions_type.UnitsOfDistribution) || modified
	modified = stage.CleanPointer(&md_digitaltransferoptions_type.TransferSize) || modified
	modified = stage.CleanPointer(&md_digitaltransferoptions_type.OffLine) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Distribution_PropertyType
func (md_distribution_propertytype *MD_Distribution_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_distribution_propertytype.MD_Distribution) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Distribution_Type
func (md_distribution_type *MD_Distribution_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_distribution_type.DistributionFormat) || modified
	modified = stage.CleanSlice(&md_distribution_type.Distributor) || modified
	modified = stage.CleanSlice(&md_distribution_type.TransferOptions) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Distributor_PropertyType
func (md_distributor_propertytype *MD_Distributor_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_distributor_propertytype.MD_Distributor) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Distributor_Type
func (md_distributor_type *MD_Distributor_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_distributor_type.DistributionOrderProcess) || modified
	modified = stage.CleanSlice(&md_distributor_type.DistributorFormat) || modified
	modified = stage.CleanSlice(&md_distributor_type.DistributorTransferOptions) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_distributor_type.DistributorContact) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Format_PropertyType
func (md_format_propertytype *MD_Format_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_format_propertytype.MD_Format) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Format_Type
func (md_format_type *MD_Format_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_format_type.FormatDistributor) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_format_type.NameXSD) || modified
	modified = stage.CleanPointer(&md_format_type.Version) || modified
	modified = stage.CleanPointer(&md_format_type.AmendmentNumber) || modified
	modified = stage.CleanPointer(&md_format_type.Specification) || modified
	modified = stage.CleanPointer(&md_format_type.FileDecompressionTechnique) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Identification_PropertyType
func (md_identification_propertytype *MD_Identification_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_identification_propertytype.MD_DataIdentification) || modified
	modified = stage.CleanPointer(&md_identification_propertytype.MD_ServiceIdentification) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Identifier_PropertyType
func (md_identifier_propertytype *MD_Identifier_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_identifier_propertytype.MD_Identifier) || modified
	modified = stage.CleanPointer(&md_identifier_propertytype.RS_Identifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Identifier_Type
func (md_identifier_type *MD_Identifier_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_identifier_type.Authority) || modified
	modified = stage.CleanPointer(&md_identifier_type.Code) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Keywords_PropertyType
func (md_keywords_propertytype *MD_Keywords_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_keywords_propertytype.MD_Keywords) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Keywords_Type
func (md_keywords_type *MD_Keywords_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_keywords_type.Keyword) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_keywords_type.Type_XSD) || modified
	modified = stage.CleanPointer(&md_keywords_type.ThesaurusName) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_LegalConstraints_PropertyType
func (md_legalconstraints_propertytype *MD_LegalConstraints_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_legalconstraints_propertytype.MD_LegalConstraints) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_LegalConstraints_Type
func (md_legalconstraints_type *MD_LegalConstraints_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_legalconstraints_type.UseLimitation) || modified
	modified = stage.CleanSlice(&md_legalconstraints_type.AccessConstraints) || modified
	modified = stage.CleanSlice(&md_legalconstraints_type.UseConstraints) || modified
	modified = stage.CleanSlice(&md_legalconstraints_type.OtherConstraints) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_MaintenanceInformation_PropertyType
func (md_maintenanceinformation_propertytype *MD_MaintenanceInformation_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_maintenanceinformation_propertytype.MD_MaintenanceInformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_MaintenanceInformation_Type
func (md_maintenanceinformation_type *MD_MaintenanceInformation_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_maintenanceinformation_type.UpdateScope) || modified
	modified = stage.CleanSlice(&md_maintenanceinformation_type.UpdateScopeDescription) || modified
	modified = stage.CleanSlice(&md_maintenanceinformation_type.MaintenanceNote) || modified
	modified = stage.CleanSlice(&md_maintenanceinformation_type.Contact) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_maintenanceinformation_type.MaintenanceAndUpdateFrequency) || modified
	modified = stage.CleanPointer(&md_maintenanceinformation_type.DateOfNextUpdate) || modified
	modified = stage.CleanPointer(&md_maintenanceinformation_type.UserDefinedMaintenanceFrequency) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Medium_PropertyType
func (md_medium_propertytype *MD_Medium_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_medium_propertytype.MD_Medium) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Medium_Type
func (md_medium_type *MD_Medium_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_medium_type.Density) || modified
	modified = stage.CleanSlice(&md_medium_type.MediumFormat) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_medium_type.NameXSD) || modified
	modified = stage.CleanPointer(&md_medium_type.DensityUnits) || modified
	modified = stage.CleanPointer(&md_medium_type.Volumes) || modified
	modified = stage.CleanPointer(&md_medium_type.MediumNote) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ReferenceSystem_PropertyType
func (md_referencesystem_propertytype *MD_ReferenceSystem_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_referencesystem_propertytype.MD_ReferenceSystem) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ReferenceSystem_Type
func (md_referencesystem_type *MD_ReferenceSystem_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_referencesystem_type.ReferenceSystemIdentifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_RepresentativeFraction_PropertyType
func (md_representativefraction_propertytype *MD_RepresentativeFraction_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_representativefraction_propertytype.MD_RepresentativeFraction) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_RepresentativeFraction_Type
func (md_representativefraction_type *MD_RepresentativeFraction_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_representativefraction_type.Denominator) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Resolution_PropertyType
func (md_resolution_propertytype *MD_Resolution_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_resolution_propertytype.MD_Resolution) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Resolution_Type
func (md_resolution_type *MD_Resolution_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_resolution_type.EquivalentScale) || modified
	modified = stage.CleanPointer(&md_resolution_type.Distance) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ScopeDescription_PropertyType
func (md_scopedescription_propertytype *MD_ScopeDescription_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_scopedescription_propertytype.MD_ScopeDescription) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ScopeDescription_Type
func (md_scopedescription_type *MD_ScopeDescription_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_scopedescription_type.Dataset) || modified
	modified = stage.CleanPointer(&md_scopedescription_type.Other) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_SecurityConstraints_PropertyType
func (md_securityconstraints_propertytype *MD_SecurityConstraints_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_securityconstraints_propertytype.MD_SecurityConstraints) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_SecurityConstraints_Type
func (md_securityconstraints_type *MD_SecurityConstraints_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_securityconstraints_type.UseLimitation) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_securityconstraints_type.Classification) || modified
	modified = stage.CleanPointer(&md_securityconstraints_type.UserNote) || modified
	modified = stage.CleanPointer(&md_securityconstraints_type.ClassificationSystem) || modified
	modified = stage.CleanPointer(&md_securityconstraints_type.HandlingDescription) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ServiceIdentification_PropertyType
func (md_serviceidentification_propertytype *MD_ServiceIdentification_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_serviceidentification_propertytype.MD_ServiceIdentification) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_ServiceIdentification_Type
func (md_serviceidentification_type *MD_ServiceIdentification_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_serviceidentification_type.Credit) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.Status) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.PointOfContact) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.ResourceMaintenance) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.GraphicOverview) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.ResourceFormat) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.DescriptiveKeywords) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.ResourceSpecificUsage) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.ResourceConstraints) || modified
	modified = stage.CleanSlice(&md_serviceidentification_type.AggregationInfo) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_serviceidentification_type.Citation) || modified
	modified = stage.CleanPointer(&md_serviceidentification_type.Abstract) || modified
	modified = stage.CleanPointer(&md_serviceidentification_type.Purpose) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_StandardOrderProcess_PropertyType
func (md_standardorderprocess_propertytype *MD_StandardOrderProcess_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_standardorderprocess_propertytype.MD_StandardOrderProcess) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_StandardOrderProcess_Type
func (md_standardorderprocess_type *MD_StandardOrderProcess_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_standardorderprocess_type.Fees) || modified
	modified = stage.CleanPointer(&md_standardorderprocess_type.PlannedAvailableDateTime) || modified
	modified = stage.CleanPointer(&md_standardorderprocess_type.OrderingInstructions) || modified
	modified = stage.CleanPointer(&md_standardorderprocess_type.Turnaround) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Usage_PropertyType
func (md_usage_propertytype *MD_Usage_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&md_usage_propertytype.MD_Usage) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MD_Usage_Type
func (md_usage_type *MD_Usage_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&md_usage_type.UserContactInfo) || modified
	// insertion point per field
	modified = stage.CleanPointer(&md_usage_type.SpecificUsage) || modified
	modified = stage.CleanPointer(&md_usage_type.UsageDateTime) || modified
	modified = stage.CleanPointer(&md_usage_type.UserDeterminedLimitations) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MappingRuleType
func (mappingruletype *MappingRuleType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&mappingruletype.RuleReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MemberName_PropertyType
func (membername_propertytype *MemberName_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&membername_propertytype.MemberName) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MemberName_Type
func (membername_type *MemberName_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&membername_type.AName) || modified
	modified = stage.CleanPointer(&membername_type.AttributeType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MetaDataPropertyType
func (metadatapropertytype *MetaDataPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&metadatapropertytype.GenericMetaData) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MovingObjectStatusType
func (movingobjectstatustype *MovingObjectStatusType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&movingobjectstatustype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&movingobjectstatustype.Description) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.DescriptionReference) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.ValidTime) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.DataSource) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.Position) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.Pos) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.LocationReference) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.Location) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.PriorityLocation) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.Bearing) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.Status) || modified
	modified = stage.CleanPointer(&movingobjectstatustype.StatusReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiCurvePropertyType
func (multicurvepropertytype *MultiCurvePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&multicurvepropertytype.MultiCurve) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiCurveType
func (multicurvetype *MultiCurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&multicurvetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&multicurvetype.CurveMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&multicurvetype.Description) || modified
	modified = stage.CleanPointer(&multicurvetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&multicurvetype.CurveMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiGeometryPropertyType
func (multigeometrypropertytype *MultiGeometryPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&multigeometrypropertytype.MultiGeometry) || modified
	modified = stage.CleanPointer(&multigeometrypropertytype.MultiPoint) || modified
	modified = stage.CleanPointer(&multigeometrypropertytype.MultiCurve) || modified
	modified = stage.CleanPointer(&multigeometrypropertytype.MultiSurface) || modified
	modified = stage.CleanPointer(&multigeometrypropertytype.MultiSolid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiGeometryType
func (multigeometrytype *MultiGeometryType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&multigeometrytype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&multigeometrytype.GeometryMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&multigeometrytype.Description) || modified
	modified = stage.CleanPointer(&multigeometrytype.DescriptionReference) || modified
	modified = stage.CleanPointer(&multigeometrytype.GeometryMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiPointPropertyType
func (multipointpropertytype *MultiPointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&multipointpropertytype.MultiPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiPointType
func (multipointtype *MultiPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&multipointtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&multipointtype.PointMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&multipointtype.Description) || modified
	modified = stage.CleanPointer(&multipointtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&multipointtype.PointMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiSolidPropertyType
func (multisolidpropertytype *MultiSolidPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&multisolidpropertytype.MultiSolid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiSolidType
func (multisolidtype *MultiSolidType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&multisolidtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&multisolidtype.SolidMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&multisolidtype.Description) || modified
	modified = stage.CleanPointer(&multisolidtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&multisolidtype.SolidMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiSurfacePropertyType
func (multisurfacepropertytype *MultiSurfacePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&multisurfacepropertytype.MultiSurface) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiSurfaceType
func (multisurfacetype *MultiSurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&multisurfacetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&multisurfacetype.SurfaceMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&multisurfacetype.Description) || modified
	modified = stage.CleanPointer(&multisurfacetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&multisurfacetype.SurfaceMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiplicityRange_PropertyType
func (multiplicityrange_propertytype *MultiplicityRange_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&multiplicityrange_propertytype.MultiplicityRange) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by MultiplicityRange_Type
func (multiplicityrange_type *MultiplicityRange_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&multiplicityrange_type.Lower) || modified
	modified = stage.CleanPointer(&multiplicityrange_type.Upper) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Multiplicity_PropertyType
func (multiplicity_propertytype *Multiplicity_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&multiplicity_propertytype.Multiplicity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Multiplicity_Type
func (multiplicity_type *Multiplicity_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&multiplicity_type.Range_XSD) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NodeOrEdgePropertyType
func (nodeoredgepropertytype *NodeOrEdgePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&nodeoredgepropertytype.Node) || modified
	modified = stage.CleanPointer(&nodeoredgepropertytype.Edge) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NodePropertyType
func (nodepropertytype *NodePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&nodepropertytype.Node) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NodeType
func (nodetype *NodeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&nodetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&nodetype.DirectedEdge) || modified
	// insertion point per field
	modified = stage.CleanPointer(&nodetype.Description) || modified
	modified = stage.CleanPointer(&nodetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&nodetype.Container) || modified
	modified = stage.CleanPointer(&nodetype.PointProperty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObliqueCartesianCSPropertyType
func (obliquecartesiancspropertytype *ObliqueCartesianCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&obliquecartesiancspropertytype.ObliqueCartesianCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObliqueCartesianCSType
func (obliquecartesiancstype *ObliqueCartesianCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&obliquecartesiancstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&obliquecartesiancstype.Axis) || modified
	modified = stage.CleanSlice(&obliquecartesiancstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&obliquecartesiancstype.Description) || modified
	modified = stage.CleanPointer(&obliquecartesiancstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ObservationType
func (observationtype *ObservationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&observationtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&observationtype.Description) || modified
	modified = stage.CleanPointer(&observationtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&observationtype.BoundedBy) || modified
	modified = stage.CleanPointer(&observationtype.Location) || modified
	modified = stage.CleanPointer(&observationtype.PriorityLocation) || modified
	modified = stage.CleanPointer(&observationtype.ValidTime) || modified
	modified = stage.CleanPointer(&observationtype.Using) || modified
	modified = stage.CleanPointer(&observationtype.Target) || modified
	modified = stage.CleanPointer(&observationtype.Subject) || modified
	modified = stage.CleanPointer(&observationtype.ResultOf) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OffsetCurveType
func (offsetcurvetype *OffsetCurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&offsetcurvetype.OffsetBase) || modified
	modified = stage.CleanPointer(&offsetcurvetype.RefDirection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OperationMethodPropertyType
func (operationmethodpropertytype *OperationMethodPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&operationmethodpropertytype.OperationMethod) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OperationMethodType
func (operationmethodtype *OperationMethodType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&operationmethodtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&operationmethodtype.Parameter) || modified
	modified = stage.CleanSlice(&operationmethodtype.IncludesParameter) || modified
	modified = stage.CleanSlice(&operationmethodtype.GeneralOperationParameter) || modified
	modified = stage.CleanSlice(&operationmethodtype.UsesParameter) || modified
	// insertion point per field
	modified = stage.CleanPointer(&operationmethodtype.Description) || modified
	modified = stage.CleanPointer(&operationmethodtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&operationmethodtype.FormulaCitation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OperationParameterGroupPropertyType
func (operationparametergrouppropertytype *OperationParameterGroupPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&operationparametergrouppropertytype.OperationParameterGroup) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OperationParameterGroupType
func (operationparametergrouptype *OperationParameterGroupType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&operationparametergrouptype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&operationparametergrouptype.Parameter) || modified
	modified = stage.CleanSlice(&operationparametergrouptype.IncludesParameter) || modified
	modified = stage.CleanSlice(&operationparametergrouptype.GeneralOperationParameter) || modified
	modified = stage.CleanSlice(&operationparametergrouptype.UsesParameter) || modified
	// insertion point per field
	modified = stage.CleanPointer(&operationparametergrouptype.Description) || modified
	modified = stage.CleanPointer(&operationparametergrouptype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OperationParameterPropertyType
func (operationparameterpropertytype *OperationParameterPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&operationparameterpropertytype.OperationParameter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OperationParameterType
func (operationparametertype *OperationParameterType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&operationparametertype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&operationparametertype.Description) || modified
	modified = stage.CleanPointer(&operationparametertype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OperationPropertyType
func (operationpropertytype *OperationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&operationpropertytype.Conversion) || modified
	modified = stage.CleanPointer(&operationpropertytype.Transformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OrientableCurveType
func (orientablecurvetype *OrientableCurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&orientablecurvetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&orientablecurvetype.Description) || modified
	modified = stage.CleanPointer(&orientablecurvetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&orientablecurvetype.BaseCurve) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by OrientableSurfaceType
func (orientablesurfacetype *OrientableSurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&orientablesurfacetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&orientablesurfacetype.Description) || modified
	modified = stage.CleanPointer(&orientablesurfacetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&orientablesurfacetype.BaseSurface) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PT_FreeText_PropertyType
func (pt_freetext_propertytype *PT_FreeText_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pt_freetext_propertytype.LocalisedCharacterString) || modified
	modified = stage.CleanPointer(&pt_freetext_propertytype.PT_FreeText) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PT_FreeText_Type
func (pt_freetext_type *PT_FreeText_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pt_freetext_type.TextGroup) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PT_LocaleContainer_PropertyType
func (pt_localecontainer_propertytype *PT_LocaleContainer_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pt_localecontainer_propertytype.PT_LocaleContainer) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PT_LocaleContainer_Type
func (pt_localecontainer_type *PT_LocaleContainer_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pt_localecontainer_type.Date) || modified
	modified = stage.CleanSlice(&pt_localecontainer_type.ResponsibleParty) || modified
	modified = stage.CleanSlice(&pt_localecontainer_type.LocalisedString) || modified
	// insertion point per field
	modified = stage.CleanPointer(&pt_localecontainer_type.Description) || modified
	modified = stage.CleanPointer(&pt_localecontainer_type.Locale) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PT_Locale_PropertyType
func (pt_locale_propertytype *PT_Locale_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pt_locale_propertytype.PT_Locale) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PT_Locale_Type
func (pt_locale_type *PT_Locale_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pt_locale_type.LanguageCode) || modified
	modified = stage.CleanPointer(&pt_locale_type.Country) || modified
	modified = stage.CleanPointer(&pt_locale_type.CharacterEncoding) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParameterValueGroupType
func (parametervaluegrouptype *ParameterValueGroupType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&parametervaluegrouptype.ParameterValue) || modified
	modified = stage.CleanSlice(&parametervaluegrouptype.UsesValue) || modified
	modified = stage.CleanSlice(&parametervaluegrouptype.IncludesValue) || modified
	// insertion point per field
	modified = stage.CleanPointer(&parametervaluegrouptype.Group) || modified
	modified = stage.CleanPointer(&parametervaluegrouptype.ValuesOfGroup) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ParameterValueType
func (parametervaluetype *ParameterValueType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&parametervaluetype.DmsAngleValue) || modified
	modified = stage.CleanPointer(&parametervaluetype.OperationParameter) || modified
	modified = stage.CleanPointer(&parametervaluetype.ValueOfParameter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PassThroughOperationPropertyType
func (passthroughoperationpropertytype *PassThroughOperationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&passthroughoperationpropertytype.PassThroughOperation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PassThroughOperationType
func (passthroughoperationtype *PassThroughOperationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&passthroughoperationtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&passthroughoperationtype.CoordinateOperationAccuracy) || modified
	// insertion point per field
	modified = stage.CleanPointer(&passthroughoperationtype.Description) || modified
	modified = stage.CleanPointer(&passthroughoperationtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&passthroughoperationtype.DomainOfValidity) || modified
	modified = stage.CleanPointer(&passthroughoperationtype.SourceCRS) || modified
	modified = stage.CleanPointer(&passthroughoperationtype.TargetCRS) || modified
	modified = stage.CleanPointer(&passthroughoperationtype.CoordOperation) || modified
	modified = stage.CleanPointer(&passthroughoperationtype.UsesSingleOperation) || modified
	modified = stage.CleanPointer(&passthroughoperationtype.UsesOperation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PointArrayPropertyType
func (pointarraypropertytype *PointArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pointarraypropertytype.Point) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PointPropertyType
func (pointpropertytype *PointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&pointpropertytype.Point) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PointType
func (pointtype *PointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&pointtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&pointtype.Description) || modified
	modified = stage.CleanPointer(&pointtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&pointtype.Pos) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PolarCSPropertyType
func (polarcspropertytype *PolarCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&polarcspropertytype.PolarCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PolarCSType
func (polarcstype *PolarCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&polarcstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&polarcstype.Axis) || modified
	modified = stage.CleanSlice(&polarcstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&polarcstype.Description) || modified
	modified = stage.CleanPointer(&polarcstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PolygonPatchType
func (polygonpatchtype *PolygonPatchType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&polygonpatchtype.Interior) || modified
	// insertion point per field
	modified = stage.CleanPointer(&polygonpatchtype.Exterior) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PolygonType
func (polygontype *PolygonType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&polygontype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&polygontype.Interior) || modified
	// insertion point per field
	modified = stage.CleanPointer(&polygontype.Description) || modified
	modified = stage.CleanPointer(&polygontype.DescriptionReference) || modified
	modified = stage.CleanPointer(&polygontype.Exterior) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PrimeMeridianPropertyType
func (primemeridianpropertytype *PrimeMeridianPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&primemeridianpropertytype.PrimeMeridian) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PrimeMeridianType
func (primemeridiantype *PrimeMeridianType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&primemeridiantype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&primemeridiantype.Description) || modified
	modified = stage.CleanPointer(&primemeridiantype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PriorityLocationPropertyType
func (prioritylocationpropertytype *PriorityLocationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&prioritylocationpropertytype.MultiGeometry) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.MultiPoint) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.MultiCurve) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.MultiSurface) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.MultiSolid) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Solid) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.CompositeSolid) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Surface) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Tin) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Shell) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Polygon) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Point) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Curve) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Ring) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.LinearRing) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.LineString) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.CompositeCurve) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.GeometricComplex) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.Grid) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.RectifiedGrid) || modified
	modified = stage.CleanPointer(&prioritylocationpropertytype.LocationString) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcedurePropertyType
func (procedurepropertytype *ProcedurePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&procedurepropertytype.DynamicFeature) || modified
	modified = stage.CleanPointer(&procedurepropertytype.DynamicFeatureCollection) || modified
	modified = stage.CleanPointer(&procedurepropertytype.MultiPointCoverage) || modified
	modified = stage.CleanPointer(&procedurepropertytype.MultiCurveCoverage) || modified
	modified = stage.CleanPointer(&procedurepropertytype.MultiSurfaceCoverage) || modified
	modified = stage.CleanPointer(&procedurepropertytype.MultiSolidCoverage) || modified
	modified = stage.CleanPointer(&procedurepropertytype.GridCoverage) || modified
	modified = stage.CleanPointer(&procedurepropertytype.RectifiedGridCoverage) || modified
	modified = stage.CleanPointer(&procedurepropertytype.Observation) || modified
	modified = stage.CleanPointer(&procedurepropertytype.DirectedObservation) || modified
	modified = stage.CleanPointer(&procedurepropertytype.DirectedObservationAtDistance) || modified
	modified = stage.CleanPointer(&procedurepropertytype.FeatureCollection) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProjectedCRSPropertyType
func (projectedcrspropertytype *ProjectedCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&projectedcrspropertytype.ProjectedCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProjectedCRSType
func (projectedcrstype *ProjectedCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&projectedcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&projectedcrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&projectedcrstype.Description) || modified
	modified = stage.CleanPointer(&projectedcrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&projectedcrstype.Conversion) || modified
	modified = stage.CleanPointer(&projectedcrstype.DefinedByConversion) || modified
	modified = stage.CleanPointer(&projectedcrstype.BaseGeodeticCRS) || modified
	modified = stage.CleanPointer(&projectedcrstype.BaseGeographicCRS) || modified
	modified = stage.CleanPointer(&projectedcrstype.CartesianCS) || modified
	modified = stage.CleanPointer(&projectedcrstype.UsesCartesianCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by QuantityPropertyType
func (quantitypropertytype *QuantityPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&quantitypropertytype.Quantity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RS_Identifier_PropertyType
func (rs_identifier_propertytype *RS_Identifier_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&rs_identifier_propertytype.RS_Identifier) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RS_Identifier_Type
func (rs_identifier_type *RS_Identifier_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&rs_identifier_type.Authority) || modified
	modified = stage.CleanPointer(&rs_identifier_type.Code) || modified
	modified = stage.CleanPointer(&rs_identifier_type.CodeSpace) || modified
	modified = stage.CleanPointer(&rs_identifier_type.Version) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RS_ReferenceSystem_PropertyType
func (rs_referencesystem_propertytype *RS_ReferenceSystem_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&rs_referencesystem_propertytype.AbstractRS_ReferenceSystem) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RangeSetType
func (rangesettype *RangeSetType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&rangesettype.ValueArray) || modified
	// insertion point per field
	modified = stage.CleanPointer(&rangesettype.DataBlock) || modified
	modified = stage.CleanPointer(&rangesettype.File) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RecordType_PropertyType
func (recordtype_propertytype *RecordType_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&recordtype_propertytype.RecordType) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RectangleType
func (rectangletype *RectangleType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&rectangletype.Exterior) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RectifiedGridType
func (rectifiedgridtype *RectifiedGridType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&rectifiedgridtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&rectifiedgridtype.OffsetVector) || modified
	// insertion point per field
	modified = stage.CleanPointer(&rectifiedgridtype.Description) || modified
	modified = stage.CleanPointer(&rectifiedgridtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&rectifiedgridtype.Limits) || modified
	modified = stage.CleanPointer(&rectifiedgridtype.Origin) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RelatedTimeType
func (relatedtimetype *RelatedTimeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&relatedtimetype.TimeInstant) || modified
	modified = stage.CleanPointer(&relatedtimetype.TimePeriod) || modified
	modified = stage.CleanPointer(&relatedtimetype.TimeNode) || modified
	modified = stage.CleanPointer(&relatedtimetype.TimeEdge) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RingPropertyType
func (ringpropertytype *RingPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&ringpropertytype.Ring) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RingType
func (ringtype *RingType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&ringtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&ringtype.CurveMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&ringtype.Description) || modified
	modified = stage.CleanPointer(&ringtype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SC_CRS_PropertyType
func (sc_crs_propertytype *SC_CRS_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&sc_crs_propertytype.ProjectedCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.DerivedCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.GeodeticCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.VerticalCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.EngineeringCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.ImageCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.TemporalCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.GeographicCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.GeocentricCRS) || modified
	modified = stage.CleanPointer(&sc_crs_propertytype.CompoundCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SecondDefiningParameter
func (seconddefiningparameter *SecondDefiningParameter) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&seconddefiningparameter.SecondDefiningParameter) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ShellPropertyType
func (shellpropertytype *ShellPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&shellpropertytype.Shell) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ShellType
func (shelltype *ShellType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&shelltype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&shelltype.SurfaceMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&shelltype.Description) || modified
	modified = stage.CleanPointer(&shelltype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SingleCRSPropertyType
func (singlecrspropertytype *SingleCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&singlecrspropertytype.ProjectedCRS) || modified
	modified = stage.CleanPointer(&singlecrspropertytype.DerivedCRS) || modified
	modified = stage.CleanPointer(&singlecrspropertytype.GeodeticCRS) || modified
	modified = stage.CleanPointer(&singlecrspropertytype.VerticalCRS) || modified
	modified = stage.CleanPointer(&singlecrspropertytype.EngineeringCRS) || modified
	modified = stage.CleanPointer(&singlecrspropertytype.ImageCRS) || modified
	modified = stage.CleanPointer(&singlecrspropertytype.TemporalCRS) || modified
	modified = stage.CleanPointer(&singlecrspropertytype.GeographicCRS) || modified
	modified = stage.CleanPointer(&singlecrspropertytype.GeocentricCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SingleOperationPropertyType
func (singleoperationpropertytype *SingleOperationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&singleoperationpropertytype.PassThroughOperation) || modified
	modified = stage.CleanPointer(&singleoperationpropertytype.Conversion) || modified
	modified = stage.CleanPointer(&singleoperationpropertytype.Transformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SolidArrayPropertyType
func (solidarraypropertytype *SolidArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&solidarraypropertytype.Solid) || modified
	modified = stage.CleanSlice(&solidarraypropertytype.CompositeSolid) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SolidPropertyType
func (solidpropertytype *SolidPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&solidpropertytype.Solid) || modified
	modified = stage.CleanPointer(&solidpropertytype.CompositeSolid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SolidType
func (solidtype *SolidType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&solidtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&solidtype.Interior) || modified
	// insertion point per field
	modified = stage.CleanPointer(&solidtype.Description) || modified
	modified = stage.CleanPointer(&solidtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&solidtype.Exterior) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SphereType
func (spheretype *SphereType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&spheretype.Rows_1) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SphericalCSPropertyType
func (sphericalcspropertytype *SphericalCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&sphericalcspropertytype.SphericalCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SphericalCSType
func (sphericalcstype *SphericalCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&sphericalcstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&sphericalcstype.Axis) || modified
	modified = stage.CleanSlice(&sphericalcstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&sphericalcstype.Description) || modified
	modified = stage.CleanPointer(&sphericalcstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfaceArrayPropertyType
func (surfacearraypropertytype *SurfaceArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surfacearraypropertytype.Surface) || modified
	modified = stage.CleanSlice(&surfacearraypropertytype.PolyhedralSurface) || modified
	modified = stage.CleanSlice(&surfacearraypropertytype.TriangulatedSurface) || modified
	modified = stage.CleanSlice(&surfacearraypropertytype.Tin) || modified
	modified = stage.CleanSlice(&surfacearraypropertytype.OrientableSurface) || modified
	modified = stage.CleanSlice(&surfacearraypropertytype.Shell) || modified
	modified = stage.CleanSlice(&surfacearraypropertytype.Polygon) || modified
	modified = stage.CleanSlice(&surfacearraypropertytype.CompositeSurface) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfacePatchArrayPropertyType
func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surfacepatcharraypropertytype.PolygonPatch) || modified
	modified = stage.CleanSlice(&surfacepatcharraypropertytype.Triangle) || modified
	modified = stage.CleanSlice(&surfacepatcharraypropertytype.Rectangle) || modified
	modified = stage.CleanSlice(&surfacepatcharraypropertytype.Cone) || modified
	modified = stage.CleanSlice(&surfacepatcharraypropertytype.Cylinder) || modified
	modified = stage.CleanSlice(&surfacepatcharraypropertytype.Sphere) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfacePropertyType
func (surfacepropertytype *SurfacePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&surfacepropertytype.Surface) || modified
	modified = stage.CleanPointer(&surfacepropertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&surfacepropertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&surfacepropertytype.Tin) || modified
	modified = stage.CleanPointer(&surfacepropertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&surfacepropertytype.Shell) || modified
	modified = stage.CleanPointer(&surfacepropertytype.Polygon) || modified
	modified = stage.CleanPointer(&surfacepropertytype.CompositeSurface) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SurfaceType
func (surfacetype *SurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&surfacetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&surfacetype.Description) || modified
	modified = stage.CleanPointer(&surfacetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&surfacetype.Patches) || modified
	modified = stage.CleanPointer(&surfacetype.PolygonPatches) || modified
	modified = stage.CleanPointer(&surfacetype.TrianglePatches) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TM_Primitive_PropertyType
func (tm_primitive_propertytype *TM_Primitive_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&tm_primitive_propertytype.TimeInstant) || modified
	modified = stage.CleanPointer(&tm_primitive_propertytype.TimePeriod) || modified
	modified = stage.CleanPointer(&tm_primitive_propertytype.TimeNode) || modified
	modified = stage.CleanPointer(&tm_primitive_propertytype.TimeEdge) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TargetPropertyType
func (targetpropertytype *TargetPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&targetpropertytype.DynamicFeature) || modified
	modified = stage.CleanPointer(&targetpropertytype.DynamicFeatureCollection) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiPointCoverage) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiCurveCoverage) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiSurfaceCoverage) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiSolidCoverage) || modified
	modified = stage.CleanPointer(&targetpropertytype.GridCoverage) || modified
	modified = stage.CleanPointer(&targetpropertytype.RectifiedGridCoverage) || modified
	modified = stage.CleanPointer(&targetpropertytype.Observation) || modified
	modified = stage.CleanPointer(&targetpropertytype.DirectedObservation) || modified
	modified = stage.CleanPointer(&targetpropertytype.DirectedObservationAtDistance) || modified
	modified = stage.CleanPointer(&targetpropertytype.FeatureCollection) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiGeometry) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiPoint) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiCurve) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiSurface) || modified
	modified = stage.CleanPointer(&targetpropertytype.MultiSolid) || modified
	modified = stage.CleanPointer(&targetpropertytype.Solid) || modified
	modified = stage.CleanPointer(&targetpropertytype.CompositeSolid) || modified
	modified = stage.CleanPointer(&targetpropertytype.Surface) || modified
	modified = stage.CleanPointer(&targetpropertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&targetpropertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&targetpropertytype.Tin) || modified
	modified = stage.CleanPointer(&targetpropertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&targetpropertytype.Shell) || modified
	modified = stage.CleanPointer(&targetpropertytype.Polygon) || modified
	modified = stage.CleanPointer(&targetpropertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&targetpropertytype.Point) || modified
	modified = stage.CleanPointer(&targetpropertytype.Curve) || modified
	modified = stage.CleanPointer(&targetpropertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&targetpropertytype.Ring) || modified
	modified = stage.CleanPointer(&targetpropertytype.LinearRing) || modified
	modified = stage.CleanPointer(&targetpropertytype.LineString) || modified
	modified = stage.CleanPointer(&targetpropertytype.CompositeCurve) || modified
	modified = stage.CleanPointer(&targetpropertytype.GeometricComplex) || modified
	modified = stage.CleanPointer(&targetpropertytype.Grid) || modified
	modified = stage.CleanPointer(&targetpropertytype.RectifiedGrid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TemporalCRSPropertyType
func (temporalcrspropertytype *TemporalCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&temporalcrspropertytype.TemporalCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TemporalCRSType
func (temporalcrstype *TemporalCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&temporalcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&temporalcrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&temporalcrstype.Description) || modified
	modified = stage.CleanPointer(&temporalcrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&temporalcrstype.TimeCS) || modified
	modified = stage.CleanPointer(&temporalcrstype.UsesTimeCS) || modified
	modified = stage.CleanPointer(&temporalcrstype.UsesTemporalCS) || modified
	modified = stage.CleanPointer(&temporalcrstype.TemporalDatum) || modified
	modified = stage.CleanPointer(&temporalcrstype.UsesTemporalDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TemporalCSPropertyType
func (temporalcspropertytype *TemporalCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&temporalcspropertytype.TemporalCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TemporalCSType
func (temporalcstype *TemporalCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&temporalcstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&temporalcstype.Axis) || modified
	modified = stage.CleanSlice(&temporalcstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&temporalcstype.Description) || modified
	modified = stage.CleanPointer(&temporalcstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TemporalDatumBaseType
func (temporaldatumbasetype *TemporalDatumBaseType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&temporaldatumbasetype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&temporaldatumbasetype.Description) || modified
	modified = stage.CleanPointer(&temporaldatumbasetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&temporaldatumbasetype.DomainOfValidity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TemporalDatumPropertyType
func (temporaldatumpropertytype *TemporalDatumPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&temporaldatumpropertytype.TemporalDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TemporalDatumType
func (temporaldatumtype *TemporalDatumType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&temporaldatumtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&temporaldatumtype.Description) || modified
	modified = stage.CleanPointer(&temporaldatumtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&temporaldatumtype.DomainOfValidity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeCSPropertyType
func (timecspropertytype *TimeCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timecspropertytype.TimeCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeCSType
func (timecstype *TimeCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timecstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timecstype.Axis) || modified
	modified = stage.CleanSlice(&timecstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timecstype.Description) || modified
	modified = stage.CleanPointer(&timecstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeCalendarEraPropertyType
func (timecalendarerapropertytype *TimeCalendarEraPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timecalendarerapropertytype.TimeCalendarEra) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeCalendarEraType
func (timecalendareratype *TimeCalendarEraType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timecalendareratype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timecalendareratype.Description) || modified
	modified = stage.CleanPointer(&timecalendareratype.DescriptionReference) || modified
	modified = stage.CleanPointer(&timecalendareratype.ReferenceEvent) || modified
	modified = stage.CleanPointer(&timecalendareratype.EpochOfUse) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeCalendarPropertyType
func (timecalendarpropertytype *TimeCalendarPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timecalendarpropertytype.TimeCalendar) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeCalendarType
func (timecalendartype *TimeCalendarType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timecalendartype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timecalendartype.ReferenceFrame) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timecalendartype.Description) || modified
	modified = stage.CleanPointer(&timecalendartype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeClockPropertyType
func (timeclockpropertytype *TimeClockPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timeclockpropertytype.TimeClock) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeClockType
func (timeclocktype *TimeClockType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timeclocktype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timeclocktype.DateBasis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timeclocktype.Description) || modified
	modified = stage.CleanPointer(&timeclocktype.DescriptionReference) || modified
	modified = stage.CleanPointer(&timeclocktype.ReferenceEvent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeCoordinateSystemType
func (timecoordinatesystemtype *TimeCoordinateSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timecoordinatesystemtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timecoordinatesystemtype.Description) || modified
	modified = stage.CleanPointer(&timecoordinatesystemtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&timecoordinatesystemtype.OriginPosition) || modified
	modified = stage.CleanPointer(&timecoordinatesystemtype.Origin) || modified
	modified = stage.CleanPointer(&timecoordinatesystemtype.Interval) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeEdgePropertyType
func (timeedgepropertytype *TimeEdgePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timeedgepropertytype.TimeEdge) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeEdgeType
func (timeedgetype *TimeEdgeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timeedgetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timeedgetype.RelatedTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timeedgetype.Description) || modified
	modified = stage.CleanPointer(&timeedgetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&timeedgetype.Complex_XSD) || modified
	modified = stage.CleanPointer(&timeedgetype.Start) || modified
	modified = stage.CleanPointer(&timeedgetype.End) || modified
	modified = stage.CleanPointer(&timeedgetype.Extent) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeInstantPropertyType
func (timeinstantpropertytype *TimeInstantPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timeinstantpropertytype.TimeInstant) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeInstantType
func (timeinstanttype *TimeInstantType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timeinstanttype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timeinstanttype.RelatedTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timeinstanttype.Description) || modified
	modified = stage.CleanPointer(&timeinstanttype.DescriptionReference) || modified
	modified = stage.CleanPointer(&timeinstanttype.TimePosition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeNodePropertyType
func (timenodepropertytype *TimeNodePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timenodepropertytype.TimeNode) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeNodeType
func (timenodetype *TimeNodeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timenodetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timenodetype.RelatedTime) || modified
	modified = stage.CleanSlice(&timenodetype.PreviousEdge) || modified
	modified = stage.CleanSlice(&timenodetype.NextEdge) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timenodetype.Description) || modified
	modified = stage.CleanPointer(&timenodetype.DescriptionReference) || modified
	modified = stage.CleanPointer(&timenodetype.Complex_XSD) || modified
	modified = stage.CleanPointer(&timenodetype.Position) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeOrdinalEraPropertyType
func (timeordinalerapropertytype *TimeOrdinalEraPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timeordinalerapropertytype.TimeOrdinalEra) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeOrdinalEraType
func (timeordinaleratype *TimeOrdinalEraType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timeordinaleratype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timeordinaleratype.RelatedTime) || modified
	modified = stage.CleanSlice(&timeordinaleratype.Member) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timeordinaleratype.Description) || modified
	modified = stage.CleanPointer(&timeordinaleratype.DescriptionReference) || modified
	modified = stage.CleanPointer(&timeordinaleratype.Start) || modified
	modified = stage.CleanPointer(&timeordinaleratype.End) || modified
	modified = stage.CleanPointer(&timeordinaleratype.Extent) || modified
	modified = stage.CleanPointer(&timeordinaleratype.Group) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeOrdinalReferenceSystemType
func (timeordinalreferencesystemtype *TimeOrdinalReferenceSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timeordinalreferencesystemtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timeordinalreferencesystemtype.Component) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timeordinalreferencesystemtype.Description) || modified
	modified = stage.CleanPointer(&timeordinalreferencesystemtype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimePeriodPropertyType
func (timeperiodpropertytype *TimePeriodPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timeperiodpropertytype.TimePeriod) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimePeriodType
func (timeperiodtype *TimePeriodType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timeperiodtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timeperiodtype.RelatedTime) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timeperiodtype.Description) || modified
	modified = stage.CleanPointer(&timeperiodtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&timeperiodtype.BeginPosition) || modified
	modified = stage.CleanPointer(&timeperiodtype.Begin) || modified
	modified = stage.CleanPointer(&timeperiodtype.EndPosition) || modified
	modified = stage.CleanPointer(&timeperiodtype.End) || modified
	modified = stage.CleanPointer(&timeperiodtype.TimeInterval) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimePrimitivePropertyType
func (timeprimitivepropertytype *TimePrimitivePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timeprimitivepropertytype.TimeInstant) || modified
	modified = stage.CleanPointer(&timeprimitivepropertytype.TimePeriod) || modified
	modified = stage.CleanPointer(&timeprimitivepropertytype.TimeNode) || modified
	modified = stage.CleanPointer(&timeprimitivepropertytype.TimeEdge) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeReferenceSystemType
func (timereferencesystemtype *TimeReferenceSystemType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timereferencesystemtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timereferencesystemtype.Description) || modified
	modified = stage.CleanPointer(&timereferencesystemtype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeTopologyComplexPropertyType
func (timetopologycomplexpropertytype *TimeTopologyComplexPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timetopologycomplexpropertytype.TimeTopologyComplex) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeTopologyComplexType
func (timetopologycomplextype *TimeTopologyComplexType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&timetopologycomplextype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&timetopologycomplextype.Primitive) || modified
	// insertion point per field
	modified = stage.CleanPointer(&timetopologycomplextype.Description) || modified
	modified = stage.CleanPointer(&timetopologycomplextype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TimeTopologyPrimitivePropertyType
func (timetopologyprimitivepropertytype *TimeTopologyPrimitivePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&timetopologyprimitivepropertytype.TimeNode) || modified
	modified = stage.CleanPointer(&timetopologyprimitivepropertytype.TimeEdge) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TinType
func (tintype *TinType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&tintype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&tintype.StopLines) || modified
	modified = stage.CleanSlice(&tintype.BreakLines) || modified
	// insertion point per field
	modified = stage.CleanPointer(&tintype.Description) || modified
	modified = stage.CleanPointer(&tintype.DescriptionReference) || modified
	modified = stage.CleanPointer(&tintype.Patches) || modified
	modified = stage.CleanPointer(&tintype.PolygonPatches) || modified
	modified = stage.CleanPointer(&tintype.TrianglePatches) || modified
	modified = stage.CleanPointer(&tintype.ControlPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoComplexPropertyType
func (topocomplexpropertytype *TopoComplexPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&topocomplexpropertytype.TopoComplex) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoComplexType
func (topocomplextype *TopoComplexType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topocomplextype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&topocomplextype.SuperComplex) || modified
	modified = stage.CleanSlice(&topocomplextype.SubComplex) || modified
	modified = stage.CleanSlice(&topocomplextype.TopoPrimitiveMember) || modified
	// insertion point per field
	modified = stage.CleanPointer(&topocomplextype.Description) || modified
	modified = stage.CleanPointer(&topocomplextype.DescriptionReference) || modified
	modified = stage.CleanPointer(&topocomplextype.MaximalComplex) || modified
	modified = stage.CleanPointer(&topocomplextype.TopoPrimitiveMembers) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoCurvePropertyType
func (topocurvepropertytype *TopoCurvePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&topocurvepropertytype.TopoCurve) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoCurveType
func (topocurvetype *TopoCurveType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topocurvetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&topocurvetype.DirectedEdge) || modified
	// insertion point per field
	modified = stage.CleanPointer(&topocurvetype.Description) || modified
	modified = stage.CleanPointer(&topocurvetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoPointPropertyType
func (topopointpropertytype *TopoPointPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&topopointpropertytype.TopoPoint) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoPointType
func (topopointtype *TopoPointType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topopointtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&topopointtype.Description) || modified
	modified = stage.CleanPointer(&topopointtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&topopointtype.DirectedNode) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoPrimitiveArrayAssociationType
func (topoprimitivearrayassociationtype *TopoPrimitiveArrayAssociationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topoprimitivearrayassociationtype.Node) || modified
	modified = stage.CleanSlice(&topoprimitivearrayassociationtype.Edge) || modified
	modified = stage.CleanSlice(&topoprimitivearrayassociationtype.Face) || modified
	modified = stage.CleanSlice(&topoprimitivearrayassociationtype.TopoSolid) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoPrimitiveMemberType
func (topoprimitivemembertype *TopoPrimitiveMemberType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&topoprimitivemembertype.Node) || modified
	modified = stage.CleanPointer(&topoprimitivemembertype.Edge) || modified
	modified = stage.CleanPointer(&topoprimitivemembertype.Face) || modified
	modified = stage.CleanPointer(&topoprimitivemembertype.TopoSolid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoSolidPropertyType
func (toposolidpropertytype *TopoSolidPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&toposolidpropertytype.TopoSolid) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoSolidType
func (toposolidtype *TopoSolidType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&toposolidtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&toposolidtype.Isolated) || modified
	modified = stage.CleanSlice(&toposolidtype.DirectedFace) || modified
	// insertion point per field
	modified = stage.CleanPointer(&toposolidtype.Description) || modified
	modified = stage.CleanPointer(&toposolidtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&toposolidtype.SolidProperty) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoSurfacePropertyType
func (toposurfacepropertytype *TopoSurfacePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&toposurfacepropertytype.TopoSurface) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoSurfaceType
func (toposurfacetype *TopoSurfaceType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&toposurfacetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&toposurfacetype.DirectedFace) || modified
	// insertion point per field
	modified = stage.CleanPointer(&toposurfacetype.Description) || modified
	modified = stage.CleanPointer(&toposurfacetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoVolumePropertyType
func (topovolumepropertytype *TopoVolumePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&topovolumepropertytype.TopoVolume) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TopoVolumeType
func (topovolumetype *TopoVolumeType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&topovolumetype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&topovolumetype.DirectedTopoSolid) || modified
	// insertion point per field
	modified = stage.CleanPointer(&topovolumetype.Description) || modified
	modified = stage.CleanPointer(&topovolumetype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TransformationPropertyType
func (transformationpropertytype *TransformationPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&transformationpropertytype.Transformation) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TransformationType
func (transformationtype *TransformationType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&transformationtype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&transformationtype.CoordinateOperationAccuracy) || modified
	modified = stage.CleanSlice(&transformationtype.ParameterValue) || modified
	modified = stage.CleanSlice(&transformationtype.UsesValue) || modified
	modified = stage.CleanSlice(&transformationtype.IncludesValue) || modified
	// insertion point per field
	modified = stage.CleanPointer(&transformationtype.Description) || modified
	modified = stage.CleanPointer(&transformationtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&transformationtype.DomainOfValidity) || modified
	modified = stage.CleanPointer(&transformationtype.SourceCRS) || modified
	modified = stage.CleanPointer(&transformationtype.TargetCRS) || modified
	modified = stage.CleanPointer(&transformationtype.Method) || modified
	modified = stage.CleanPointer(&transformationtype.UsesMethod) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TriangleType
func (triangletype *TriangleType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&triangletype.Exterior) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TypeName_PropertyType
func (typename_propertytype *TypeName_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&typename_propertytype.TypeName) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TypeName_Type
func (typename_type *TypeName_Type) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&typename_type.AName) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnitDefinitionType
func (unitdefinitiontype *UnitDefinitionType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&unitdefinitiontype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&unitdefinitiontype.Description) || modified
	modified = stage.CleanPointer(&unitdefinitiontype.DescriptionReference) || modified
	modified = stage.CleanPointer(&unitdefinitiontype.QuantityType) || modified
	modified = stage.CleanPointer(&unitdefinitiontype.QuantityTypeReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnitOfMeasure_PropertyType
func (unitofmeasure_propertytype *UnitOfMeasure_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&unitofmeasure_propertytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&unitofmeasure_propertytype.BaseUnit) || modified
	modified = stage.CleanPointer(&unitofmeasure_propertytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&unitofmeasure_propertytype.ConventionalUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UnlimitedInteger_PropertyType
func (unlimitedinteger_propertytype *UnlimitedInteger_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&unlimitedinteger_propertytype.UnlimitedInteger) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UomAngle_PropertyType
func (uomangle_propertytype *UomAngle_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&uomangle_propertytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&uomangle_propertytype.BaseUnit) || modified
	modified = stage.CleanPointer(&uomangle_propertytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&uomangle_propertytype.ConventionalUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UomArea_PropertyType
func (uomarea_propertytype *UomArea_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&uomarea_propertytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&uomarea_propertytype.BaseUnit) || modified
	modified = stage.CleanPointer(&uomarea_propertytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&uomarea_propertytype.ConventionalUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UomLength_PropertyType
func (uomlength_propertytype *UomLength_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&uomlength_propertytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&uomlength_propertytype.BaseUnit) || modified
	modified = stage.CleanPointer(&uomlength_propertytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&uomlength_propertytype.ConventionalUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UomScale_PropertyType
func (uomscale_propertytype *UomScale_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&uomscale_propertytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&uomscale_propertytype.BaseUnit) || modified
	modified = stage.CleanPointer(&uomscale_propertytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&uomscale_propertytype.ConventionalUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UomTime_PropertyType
func (uomtime_propertytype *UomTime_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&uomtime_propertytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&uomtime_propertytype.BaseUnit) || modified
	modified = stage.CleanPointer(&uomtime_propertytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&uomtime_propertytype.ConventionalUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UomVelocity_PropertyType
func (uomvelocity_propertytype *UomVelocity_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&uomvelocity_propertytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&uomvelocity_propertytype.BaseUnit) || modified
	modified = stage.CleanPointer(&uomvelocity_propertytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&uomvelocity_propertytype.ConventionalUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UomVolume_PropertyType
func (uomvolume_propertytype *UomVolume_PropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&uomvolume_propertytype.UnitDefinition) || modified
	modified = stage.CleanPointer(&uomvolume_propertytype.BaseUnit) || modified
	modified = stage.CleanPointer(&uomvolume_propertytype.DerivedUnit) || modified
	modified = stage.CleanPointer(&uomvolume_propertytype.ConventionalUnit) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UserDefinedCSPropertyType
func (userdefinedcspropertytype *UserDefinedCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&userdefinedcspropertytype.UserDefinedCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by UserDefinedCSType
func (userdefinedcstype *UserDefinedCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&userdefinedcstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&userdefinedcstype.Axis) || modified
	modified = stage.CleanSlice(&userdefinedcstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&userdefinedcstype.Description) || modified
	modified = stage.CleanPointer(&userdefinedcstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ValueArrayPropertyType
func (valuearraypropertytype *ValueArrayPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&valuearraypropertytype.Boolean) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Category) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Count) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Quantity) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.CompositeValue) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.ValueArray) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.CategoryExtent) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.QuantityExtent) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.MultiGeometry) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.MultiPoint) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.MultiCurve) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.MultiSurface) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.MultiSolid) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Solid) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.CompositeSolid) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Surface) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Tin) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Shell) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Polygon) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Point) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Curve) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Ring) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.LinearRing) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.LineString) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.CompositeCurve) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.GeometricComplex) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.Grid) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.RectifiedGrid) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.TimeInstant) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.TimePeriod) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.TimeNode) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.TimeEdge) || modified
	modified = stage.CleanPointer(&valuearraypropertytype.TimeTopologyComplex) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ValueArrayType
func (valuearraytype *ValueArrayType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&valuearraytype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&valuearraytype.ValueComponent) || modified
	// insertion point per field
	modified = stage.CleanPointer(&valuearraytype.Description) || modified
	modified = stage.CleanPointer(&valuearraytype.DescriptionReference) || modified
	modified = stage.CleanPointer(&valuearraytype.ValueComponents) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ValuePropertyType
func (valuepropertytype *ValuePropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&valuepropertytype.Boolean) || modified
	modified = stage.CleanPointer(&valuepropertytype.Category) || modified
	modified = stage.CleanPointer(&valuepropertytype.Count) || modified
	modified = stage.CleanPointer(&valuepropertytype.Quantity) || modified
	modified = stage.CleanPointer(&valuepropertytype.CompositeValue) || modified
	modified = stage.CleanPointer(&valuepropertytype.ValueArray) || modified
	modified = stage.CleanPointer(&valuepropertytype.CategoryExtent) || modified
	modified = stage.CleanPointer(&valuepropertytype.QuantityExtent) || modified
	modified = stage.CleanPointer(&valuepropertytype.MultiGeometry) || modified
	modified = stage.CleanPointer(&valuepropertytype.MultiPoint) || modified
	modified = stage.CleanPointer(&valuepropertytype.MultiCurve) || modified
	modified = stage.CleanPointer(&valuepropertytype.MultiSurface) || modified
	modified = stage.CleanPointer(&valuepropertytype.MultiSolid) || modified
	modified = stage.CleanPointer(&valuepropertytype.Solid) || modified
	modified = stage.CleanPointer(&valuepropertytype.CompositeSolid) || modified
	modified = stage.CleanPointer(&valuepropertytype.Surface) || modified
	modified = stage.CleanPointer(&valuepropertytype.PolyhedralSurface) || modified
	modified = stage.CleanPointer(&valuepropertytype.TriangulatedSurface) || modified
	modified = stage.CleanPointer(&valuepropertytype.Tin) || modified
	modified = stage.CleanPointer(&valuepropertytype.OrientableSurface) || modified
	modified = stage.CleanPointer(&valuepropertytype.Shell) || modified
	modified = stage.CleanPointer(&valuepropertytype.Polygon) || modified
	modified = stage.CleanPointer(&valuepropertytype.CompositeSurface) || modified
	modified = stage.CleanPointer(&valuepropertytype.Point) || modified
	modified = stage.CleanPointer(&valuepropertytype.Curve) || modified
	modified = stage.CleanPointer(&valuepropertytype.OrientableCurve) || modified
	modified = stage.CleanPointer(&valuepropertytype.Ring) || modified
	modified = stage.CleanPointer(&valuepropertytype.LinearRing) || modified
	modified = stage.CleanPointer(&valuepropertytype.LineString) || modified
	modified = stage.CleanPointer(&valuepropertytype.CompositeCurve) || modified
	modified = stage.CleanPointer(&valuepropertytype.GeometricComplex) || modified
	modified = stage.CleanPointer(&valuepropertytype.Grid) || modified
	modified = stage.CleanPointer(&valuepropertytype.RectifiedGrid) || modified
	modified = stage.CleanPointer(&valuepropertytype.TimeInstant) || modified
	modified = stage.CleanPointer(&valuepropertytype.TimePeriod) || modified
	modified = stage.CleanPointer(&valuepropertytype.TimeNode) || modified
	modified = stage.CleanPointer(&valuepropertytype.TimeEdge) || modified
	modified = stage.CleanPointer(&valuepropertytype.TimeTopologyComplex) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalCRSPropertyType
func (verticalcrspropertytype *VerticalCRSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&verticalcrspropertytype.VerticalCRS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalCRSType
func (verticalcrstype *VerticalCRSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&verticalcrstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&verticalcrstype.DomainOfValidity) || modified
	// insertion point per field
	modified = stage.CleanPointer(&verticalcrstype.Description) || modified
	modified = stage.CleanPointer(&verticalcrstype.DescriptionReference) || modified
	modified = stage.CleanPointer(&verticalcrstype.VerticalCS) || modified
	modified = stage.CleanPointer(&verticalcrstype.UsesVerticalCS) || modified
	modified = stage.CleanPointer(&verticalcrstype.VerticalDatum) || modified
	modified = stage.CleanPointer(&verticalcrstype.UsesVerticalDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalCSPropertyType
func (verticalcspropertytype *VerticalCSPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&verticalcspropertytype.VerticalCS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalCSType
func (verticalcstype *VerticalCSType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&verticalcstype.MetaDataProperty) || modified
	modified = stage.CleanSlice(&verticalcstype.Axis) || modified
	modified = stage.CleanSlice(&verticalcstype.UsesAxis) || modified
	// insertion point per field
	modified = stage.CleanPointer(&verticalcstype.Description) || modified
	modified = stage.CleanPointer(&verticalcstype.DescriptionReference) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalDatumPropertyType
func (verticaldatumpropertytype *VerticalDatumPropertyType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&verticaldatumpropertytype.VerticalDatum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by VerticalDatumType
func (verticaldatumtype *VerticalDatumType) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&verticaldatumtype.MetaDataProperty) || modified
	// insertion point per field
	modified = stage.CleanPointer(&verticaldatumtype.Description) || modified
	modified = stage.CleanPointer(&verticaldatumtype.DescriptionReference) || modified
	modified = stage.CleanPointer(&verticaldatumtype.DomainOfValidity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
