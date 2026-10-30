// generated code - do not edit
package gml

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func (stage *Stage) SerializeStage(filename string) {
	stage.SerializeStage2(filename, false)
}

func (stage *Stage) SerializeStage2(filename string, addIDs bool) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.A_Rows {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_Row", instances, (*A_Row)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_controlPoints {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_controlPoint", instances, (*A_controlPoint)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_refLocations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_refLocation", instances, (*A_refLocation)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_rowss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_rows", instances, (*A_rows)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractCRSType", instances, (*AbstractCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractContinuousCoverageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractContinuousCoverageType", instances, (*AbstractContinuousCoverageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractCoordinateOperationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractCoordinateOperationType", instances, (*AbstractCoordinateOperationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractCoordinateSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractCoordinateSystemType", instances, (*AbstractCoordinateSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractCoverageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractCoverageType", instances, (*AbstractCoverageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractCurveSegmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractCurveSegmentType", instances, (*AbstractCurveSegmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractCurveType", instances, (*AbstractCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDQ_Completeness_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDQ_Completeness_Type", instances, (*AbstractDQ_Completeness_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDQ_Element_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDQ_Element_Type", instances, (*AbstractDQ_Element_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDQ_LogicalConsistency_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDQ_LogicalConsistency_Type", instances, (*AbstractDQ_LogicalConsistency_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDQ_PositionalAccuracy_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDQ_PositionalAccuracy_Type", instances, (*AbstractDQ_PositionalAccuracy_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDQ_Result_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDQ_Result_Type", instances, (*AbstractDQ_Result_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDQ_TemporalAccuracy_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDQ_TemporalAccuracy_Type", instances, (*AbstractDQ_TemporalAccuracy_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDQ_ThematicAccuracy_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDQ_ThematicAccuracy_Type", instances, (*AbstractDQ_ThematicAccuracy_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractDatumTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractDatumType", instances, (*AbstractDatumType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractEX_GeographicExtent_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractEX_GeographicExtent_Type", instances, (*AbstractEX_GeographicExtent_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractFeatureCollectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractFeatureCollectionType", instances, (*AbstractFeatureCollectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractFeatureMemberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractFeatureMemberType", instances, (*AbstractFeatureMemberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractFeatureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractFeatureType", instances, (*AbstractFeatureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGMLTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGMLType", instances, (*AbstractGMLType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeneralConversionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeneralConversionType", instances, (*AbstractGeneralConversionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeneralDerivedCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeneralDerivedCRSType", instances, (*AbstractGeneralDerivedCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeneralOperationParameterPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeneralOperationParameterPropertyType", instances, (*AbstractGeneralOperationParameterPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeneralOperationParameterTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeneralOperationParameterType", instances, (*AbstractGeneralOperationParameterType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeneralParameterValuePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeneralParameterValuePropertyType", instances, (*AbstractGeneralParameterValuePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeneralParameterValueTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeneralParameterValueType", instances, (*AbstractGeneralParameterValueType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeneralTransformationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeneralTransformationType", instances, (*AbstractGeneralTransformationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeometricAggregateTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeometricAggregateType", instances, (*AbstractGeometricAggregateType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeometricPrimitiveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeometricPrimitiveType", instances, (*AbstractGeometricPrimitiveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGeometryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGeometryType", instances, (*AbstractGeometryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractGriddedSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractGriddedSurfaceType", instances, (*AbstractGriddedSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractMD_Identification_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractMD_Identification_Type", instances, (*AbstractMD_Identification_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractMemberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractMemberType", instances, (*AbstractMemberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractMetaDataTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractMetaDataType", instances, (*AbstractMetaDataType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractMetadataPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractMetadataPropertyType", instances, (*AbstractMetadataPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractParametricCurveSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractParametricCurveSurfaceType", instances, (*AbstractParametricCurveSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractRS_ReferenceSystem_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractRS_ReferenceSystem_Type", instances, (*AbstractRS_ReferenceSystem_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractRingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractRingPropertyType", instances, (*AbstractRingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractRingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractRingType", instances, (*AbstractRingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSolidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSolidType", instances, (*AbstractSolidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSurfacePatchTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSurfacePatchType", instances, (*AbstractSurfacePatchType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractSurfaceType", instances, (*AbstractSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTimeComplexTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTimeComplexType", instances, (*AbstractTimeComplexType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTimeGeometricPrimitiveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTimeGeometricPrimitiveType", instances, (*AbstractTimeGeometricPrimitiveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTimeObjectTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTimeObjectType", instances, (*AbstractTimeObjectType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTimePrimitiveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTimePrimitiveType", instances, (*AbstractTimePrimitiveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTimeSliceType", instances, (*AbstractTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTimeTopologyPrimitiveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTimeTopologyPrimitiveType", instances, (*AbstractTimeTopologyPrimitiveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTopoPrimitiveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTopoPrimitiveType", instances, (*AbstractTopoPrimitiveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractTopologyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractTopologyType", instances, (*AbstractTopologyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AffineCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AffineCSPropertyType", instances, (*AffineCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AffineCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AffineCSType", instances, (*AffineCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AffinePlacementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AffinePlacementType", instances, (*AffinePlacementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AngleChoiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AngleChoiceType", instances, (*AngleChoiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Angle_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Angle_PropertyType", instances, (*Angle_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArcByBulgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArcByBulgeType", instances, (*ArcByBulgeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArcByCenterPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArcByCenterPointType", instances, (*ArcByCenterPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArcStringByBulgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArcStringByBulgeType", instances, (*ArcStringByBulgeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArcStringTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArcStringType", instances, (*ArcStringType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArcTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArcType", instances, (*ArcType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrayAssociationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrayAssociationType", instances, (*ArrayAssociationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ArrayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ArrayType", instances, (*ArrayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AssociationRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AssociationRoleType", instances, (*AssociationRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BSplineTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BSplineType", instances, (*BSplineType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BagTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BagType", instances, (*BagType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BaseUnitTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BaseUnitType", instances, (*BaseUnitType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BezierTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BezierType", instances, (*BezierType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Binary_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Binary_PropertyType", instances, (*Binary_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Binary_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Binary_Type", instances, (*Binary_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Booleans {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Boolean", instances, (*Boolean)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BooleanPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BooleanPropertyType", instances, (*BooleanPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Boolean_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Boolean_PropertyType", instances, (*Boolean_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BoundedFeatureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BoundedFeatureType", instances, (*BoundedFeatureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BoundingShapeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BoundingShapeType", instances, (*BoundingShapeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Address_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Address_PropertyType", instances, (*CI_Address_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Address_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Address_Type", instances, (*CI_Address_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Citation_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Citation_PropertyType", instances, (*CI_Citation_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Citation_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Citation_Type", instances, (*CI_Citation_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Contact_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Contact_PropertyType", instances, (*CI_Contact_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Contact_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Contact_Type", instances, (*CI_Contact_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_DateTypeCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_DateTypeCode_PropertyType", instances, (*CI_DateTypeCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Date_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Date_PropertyType", instances, (*CI_Date_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Date_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Date_Type", instances, (*CI_Date_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_OnLineFunctionCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_OnLineFunctionCode_PropertyType", instances, (*CI_OnLineFunctionCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_OnlineResource_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_OnlineResource_PropertyType", instances, (*CI_OnlineResource_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_OnlineResource_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_OnlineResource_Type", instances, (*CI_OnlineResource_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_PresentationFormCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_PresentationFormCode_PropertyType", instances, (*CI_PresentationFormCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_ResponsibleParty_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_ResponsibleParty_PropertyType", instances, (*CI_ResponsibleParty_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_ResponsibleParty_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_ResponsibleParty_Type", instances, (*CI_ResponsibleParty_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_RoleCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_RoleCode_PropertyType", instances, (*CI_RoleCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Series_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Series_PropertyType", instances, (*CI_Series_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Series_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Series_Type", instances, (*CI_Series_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Telephone_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Telephone_PropertyType", instances, (*CI_Telephone_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CI_Telephone_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CI_Telephone_Type", instances, (*CI_Telephone_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CRSPropertyType", instances, (*CRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CartesianCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CartesianCSPropertyType", instances, (*CartesianCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CartesianCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CartesianCSType", instances, (*CartesianCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Categorys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Category", instances, (*Category)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CategoryExtentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CategoryExtentType", instances, (*CategoryExtentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CategoryPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CategoryPropertyType", instances, (*CategoryPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CharacterString_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CharacterString_PropertyType", instances, (*CharacterString_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CircleByCenterPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CircleByCenterPointType", instances, (*CircleByCenterPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CircleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CircleType", instances, (*CircleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ClothoidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ClothoidType", instances, (*ClothoidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CompositeCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CompositeCurveType", instances, (*CompositeCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CompositeSolidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CompositeSolidType", instances, (*CompositeSolidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CompositeSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CompositeSurfaceType", instances, (*CompositeSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CompositeValueTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CompositeValueType", instances, (*CompositeValueType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CompoundCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CompoundCRSPropertyType", instances, (*CompoundCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CompoundCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CompoundCRSType", instances, (*CompoundCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConcatenatedOperationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConcatenatedOperationPropertyType", instances, (*ConcatenatedOperationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConcatenatedOperationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConcatenatedOperationType", instances, (*ConcatenatedOperationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConeType", instances, (*ConeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConventionalUnitTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConventionalUnitType", instances, (*ConventionalUnitType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConversionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConversionPropertyType", instances, (*ConversionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConversionToPreferredUnitTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConversionToPreferredUnitType", instances, (*ConversionToPreferredUnitType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ConversionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ConversionType", instances, (*ConversionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CoordinateOperationAccuracys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CoordinateOperationAccuracy", instances, (*CoordinateOperationAccuracy)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CoordinateOperationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CoordinateOperationPropertyType", instances, (*CoordinateOperationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CoordinateSystemAxisPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CoordinateSystemAxisPropertyType", instances, (*CoordinateSystemAxisPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CoordinateSystemAxisTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CoordinateSystemAxisType", instances, (*CoordinateSystemAxisType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CoordinateSystemPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CoordinateSystemPropertyType", instances, (*CoordinateSystemPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Counts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Count", instances, (*Count)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CountPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CountPropertyType", instances, (*CountPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Country_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Country_PropertyType", instances, (*Country_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CoverageFunctionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CoverageFunctionType", instances, (*CoverageFunctionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CubicSplineTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CubicSplineType", instances, (*CubicSplineType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurveArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurveArrayPropertyType", instances, (*CurveArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurvePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurvePropertyType", instances, (*CurvePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurveSegmentArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurveSegmentArrayPropertyType", instances, (*CurveSegmentArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurveType", instances, (*CurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CylinderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CylinderType", instances, (*CylinderType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CylindricalCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CylindricalCSPropertyType", instances, (*CylindricalCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CylindricalCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CylindricalCSType", instances, (*CylindricalCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DMSAngleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DMSAngleType", instances, (*DMSAngleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_AbsoluteExternalPositionalAccuracy_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_AbsoluteExternalPositionalAccuracy_PropertyType", instances, (*DQ_AbsoluteExternalPositionalAccuracy_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_AbsoluteExternalPositionalAccuracy_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_AbsoluteExternalPositionalAccuracy_Type", instances, (*DQ_AbsoluteExternalPositionalAccuracy_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_AccuracyOfATimeMeasurement_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_AccuracyOfATimeMeasurement_PropertyType", instances, (*DQ_AccuracyOfATimeMeasurement_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_AccuracyOfATimeMeasurement_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_AccuracyOfATimeMeasurement_Type", instances, (*DQ_AccuracyOfATimeMeasurement_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_CompletenessCommission_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_CompletenessCommission_PropertyType", instances, (*DQ_CompletenessCommission_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_CompletenessCommission_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_CompletenessCommission_Type", instances, (*DQ_CompletenessCommission_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_CompletenessOmission_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_CompletenessOmission_PropertyType", instances, (*DQ_CompletenessOmission_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_CompletenessOmission_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_CompletenessOmission_Type", instances, (*DQ_CompletenessOmission_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_Completeness_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_Completeness_PropertyType", instances, (*DQ_Completeness_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_ConceptualConsistency_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_ConceptualConsistency_PropertyType", instances, (*DQ_ConceptualConsistency_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_ConceptualConsistency_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_ConceptualConsistency_Type", instances, (*DQ_ConceptualConsistency_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_ConformanceResult_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_ConformanceResult_PropertyType", instances, (*DQ_ConformanceResult_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_ConformanceResult_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_ConformanceResult_Type", instances, (*DQ_ConformanceResult_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_DataQuality_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_DataQuality_PropertyType", instances, (*DQ_DataQuality_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_DataQuality_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_DataQuality_Type", instances, (*DQ_DataQuality_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_DomainConsistency_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_DomainConsistency_PropertyType", instances, (*DQ_DomainConsistency_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_DomainConsistency_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_DomainConsistency_Type", instances, (*DQ_DomainConsistency_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_Element_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_Element_PropertyType", instances, (*DQ_Element_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_EvaluationMethodTypeCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_EvaluationMethodTypeCode_PropertyType", instances, (*DQ_EvaluationMethodTypeCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_FormatConsistency_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_FormatConsistency_PropertyType", instances, (*DQ_FormatConsistency_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_FormatConsistency_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_FormatConsistency_Type", instances, (*DQ_FormatConsistency_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_GriddedDataPositionalAccuracy_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_GriddedDataPositionalAccuracy_PropertyType", instances, (*DQ_GriddedDataPositionalAccuracy_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_GriddedDataPositionalAccuracy_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_GriddedDataPositionalAccuracy_Type", instances, (*DQ_GriddedDataPositionalAccuracy_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_LogicalConsistency_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_LogicalConsistency_PropertyType", instances, (*DQ_LogicalConsistency_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_NonQuantitativeAttributeAccuracy_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_NonQuantitativeAttributeAccuracy_PropertyType", instances, (*DQ_NonQuantitativeAttributeAccuracy_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_NonQuantitativeAttributeAccuracy_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_NonQuantitativeAttributeAccuracy_Type", instances, (*DQ_NonQuantitativeAttributeAccuracy_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_PositionalAccuracy_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_PositionalAccuracy_PropertyType", instances, (*DQ_PositionalAccuracy_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_QuantitativeAttributeAccuracy_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_QuantitativeAttributeAccuracy_PropertyType", instances, (*DQ_QuantitativeAttributeAccuracy_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_QuantitativeAttributeAccuracy_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_QuantitativeAttributeAccuracy_Type", instances, (*DQ_QuantitativeAttributeAccuracy_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_QuantitativeResult_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_QuantitativeResult_PropertyType", instances, (*DQ_QuantitativeResult_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_QuantitativeResult_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_QuantitativeResult_Type", instances, (*DQ_QuantitativeResult_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_RelativeInternalPositionalAccuracy_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_RelativeInternalPositionalAccuracy_PropertyType", instances, (*DQ_RelativeInternalPositionalAccuracy_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_RelativeInternalPositionalAccuracy_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_RelativeInternalPositionalAccuracy_Type", instances, (*DQ_RelativeInternalPositionalAccuracy_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_Result_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_Result_PropertyType", instances, (*DQ_Result_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_Scope_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_Scope_PropertyType", instances, (*DQ_Scope_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_Scope_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_Scope_Type", instances, (*DQ_Scope_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_TemporalAccuracy_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_TemporalAccuracy_PropertyType", instances, (*DQ_TemporalAccuracy_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_TemporalConsistency_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_TemporalConsistency_PropertyType", instances, (*DQ_TemporalConsistency_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_TemporalConsistency_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_TemporalConsistency_Type", instances, (*DQ_TemporalConsistency_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_TemporalValidity_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_TemporalValidity_PropertyType", instances, (*DQ_TemporalValidity_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_TemporalValidity_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_TemporalValidity_Type", instances, (*DQ_TemporalValidity_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_ThematicAccuracy_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_ThematicAccuracy_PropertyType", instances, (*DQ_ThematicAccuracy_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_ThematicClassificationCorrectness_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_ThematicClassificationCorrectness_PropertyType", instances, (*DQ_ThematicClassificationCorrectness_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_ThematicClassificationCorrectness_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_ThematicClassificationCorrectness_Type", instances, (*DQ_ThematicClassificationCorrectness_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_TopologicalConsistency_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_TopologicalConsistency_PropertyType", instances, (*DQ_TopologicalConsistency_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DQ_TopologicalConsistency_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DQ_TopologicalConsistency_Type", instances, (*DQ_TopologicalConsistency_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DS_AssociationTypeCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DS_AssociationTypeCode_PropertyType", instances, (*DS_AssociationTypeCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DS_Association_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DS_Association_PropertyType", instances, (*DS_Association_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DS_Association_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DS_Association_Type", instances, (*DS_Association_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DS_InitiativeTypeCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DS_InitiativeTypeCode_PropertyType", instances, (*DS_InitiativeTypeCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DataBlockTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DataBlockType", instances, (*DataBlockType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DateTime_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DateTime_PropertyType", instances, (*DateTime_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Date_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Date_PropertyType", instances, (*Date_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DatumPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DatumPropertyType", instances, (*DatumPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Decimal_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Decimal_PropertyType", instances, (*Decimal_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DefinitionBaseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DefinitionBaseType", instances, (*DefinitionBaseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DefinitionProxyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DefinitionProxyType", instances, (*DefinitionProxyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DefinitionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DefinitionType", instances, (*DefinitionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DegreesTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DegreesType", instances, (*DegreesType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DerivationUnitTermTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DerivationUnitTermType", instances, (*DerivationUnitTermType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DerivedCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DerivedCRSPropertyType", instances, (*DerivedCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DerivedCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DerivedCRSType", instances, (*DerivedCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DerivedUnitTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DerivedUnitType", instances, (*DerivedUnitType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DictionaryEntryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DictionaryEntryType", instances, (*DictionaryEntryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DictionaryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DictionaryType", instances, (*DictionaryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectPositionListTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectPositionListType", instances, (*DirectPositionListType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectPositionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectPositionType", instances, (*DirectPositionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectedEdgePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectedEdgePropertyType", instances, (*DirectedEdgePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectedFacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectedFacePropertyType", instances, (*DirectedFacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectedNodePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectedNodePropertyType", instances, (*DirectedNodePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectedObservationAtDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectedObservationAtDistanceType", instances, (*DirectedObservationAtDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectedObservationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectedObservationType", instances, (*DirectedObservationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectedTopoSolidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectedTopoSolidPropertyType", instances, (*DirectedTopoSolidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionDescriptionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionDescriptionType", instances, (*DirectionDescriptionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionPropertyType", instances, (*DirectionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionVectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionVectorType", instances, (*DirectionVectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DiscreteCoverageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DiscreteCoverageType", instances, (*DiscreteCoverageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Distance_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Distance_PropertyType", instances, (*Distance_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DomainOfValiditys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DomainOfValidity", instances, (*DomainOfValidity)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DomainSetTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DomainSetType", instances, (*DomainSetType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DynamicFeatureCollectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DynamicFeatureCollectionType", instances, (*DynamicFeatureCollectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DynamicFeatureMemberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DynamicFeatureMemberType", instances, (*DynamicFeatureMemberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DynamicFeatureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DynamicFeatureType", instances, (*DynamicFeatureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_BoundingPolygon_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_BoundingPolygon_PropertyType", instances, (*EX_BoundingPolygon_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_BoundingPolygon_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_BoundingPolygon_Type", instances, (*EX_BoundingPolygon_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_Extent_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_Extent_PropertyType", instances, (*EX_Extent_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_Extent_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_Extent_Type", instances, (*EX_Extent_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_GeographicBoundingBox_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_GeographicBoundingBox_PropertyType", instances, (*EX_GeographicBoundingBox_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_GeographicBoundingBox_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_GeographicBoundingBox_Type", instances, (*EX_GeographicBoundingBox_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_GeographicDescription_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_GeographicDescription_PropertyType", instances, (*EX_GeographicDescription_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_GeographicDescription_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_GeographicDescription_Type", instances, (*EX_GeographicDescription_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_GeographicExtent_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_GeographicExtent_PropertyType", instances, (*EX_GeographicExtent_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_SpatialTemporalExtent_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_SpatialTemporalExtent_PropertyType", instances, (*EX_SpatialTemporalExtent_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_SpatialTemporalExtent_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_SpatialTemporalExtent_Type", instances, (*EX_SpatialTemporalExtent_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_TemporalExtent_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_TemporalExtent_PropertyType", instances, (*EX_TemporalExtent_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_TemporalExtent_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_TemporalExtent_Type", instances, (*EX_TemporalExtent_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_VerticalExtent_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_VerticalExtent_PropertyType", instances, (*EX_VerticalExtent_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EX_VerticalExtent_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EX_VerticalExtent_Type", instances, (*EX_VerticalExtent_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EdgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EdgeType", instances, (*EdgeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EllipsoidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EllipsoidPropertyType", instances, (*EllipsoidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EllipsoidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EllipsoidType", instances, (*EllipsoidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EllipsoidalCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EllipsoidalCSPropertyType", instances, (*EllipsoidalCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EllipsoidalCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EllipsoidalCSType", instances, (*EllipsoidalCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EngineeringCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EngineeringCRSPropertyType", instances, (*EngineeringCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EngineeringCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EngineeringCRSType", instances, (*EngineeringCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EngineeringDatumPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EngineeringDatumPropertyType", instances, (*EngineeringDatumPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EngineeringDatumTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EngineeringDatumType", instances, (*EngineeringDatumType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EnvelopeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EnvelopeType", instances, (*EnvelopeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EnvelopeWithTimePeriodTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EnvelopeWithTimePeriodType", instances, (*EnvelopeWithTimePeriodType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FaceOrTopoSolidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FaceOrTopoSolidPropertyType", instances, (*FaceOrTopoSolidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FaceType", instances, (*FaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FeatureArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FeatureArrayPropertyType", instances, (*FeatureArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FeatureCollectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FeatureCollectionType", instances, (*FeatureCollectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FeaturePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FeaturePropertyType", instances, (*FeaturePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FileTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FileType", instances, (*FileType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FormulaCitations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FormulaCitation", instances, (*FormulaCitation)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FormulaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FormulaType", instances, (*FormulaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GM_Object_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GM_Object_PropertyType", instances, (*GM_Object_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GM_Point_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GM_Point_PropertyType", instances, (*GM_Point_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeneralConversionPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeneralConversionPropertyType", instances, (*GeneralConversionPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeneralTransformationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeneralTransformationPropertyType", instances, (*GeneralTransformationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GenericMetaDataTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GenericMetaDataType", instances, (*GenericMetaDataType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GenericName_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GenericName_PropertyType", instances, (*GenericName_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeocentricCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeocentricCRSPropertyType", instances, (*GeocentricCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeocentricCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeocentricCRSType", instances, (*GeocentricCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeodesicStringTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeodesicStringType", instances, (*GeodesicStringType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeodesicTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeodesicType", instances, (*GeodesicType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeodeticCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeodeticCRSPropertyType", instances, (*GeodeticCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeodeticCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeodeticCRSType", instances, (*GeodeticCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeodeticDatumPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeodeticDatumPropertyType", instances, (*GeodeticDatumPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeodeticDatumTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeodeticDatumType", instances, (*GeodeticDatumType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeographicCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeographicCRSPropertyType", instances, (*GeographicCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeographicCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeographicCRSType", instances, (*GeographicCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeometricComplexPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeometricComplexPropertyType", instances, (*GeometricComplexPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeometricComplexTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeometricComplexType", instances, (*GeometricComplexType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeometricPrimitivePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeometricPrimitivePropertyType", instances, (*GeometricPrimitivePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeometryArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeometryArrayPropertyType", instances, (*GeometryArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeometryPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeometryPropertyType", instances, (*GeometryPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GridEnvelopeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GridEnvelopeType", instances, (*GridEnvelopeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GridFunctionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GridFunctionType", instances, (*GridFunctionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GridLimitsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GridLimitsType", instances, (*GridLimitsType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GridTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GridType", instances, (*GridType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HistoryPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HistoryPropertyType", instances, (*HistoryPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IdentifiedObjectTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IdentifiedObjectType", instances, (*IdentifiedObjectType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ImageCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ImageCRSPropertyType", instances, (*ImageCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ImageCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ImageCRSType", instances, (*ImageCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ImageDatumPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ImageDatumPropertyType", instances, (*ImageDatumPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ImageDatumTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ImageDatumType", instances, (*ImageDatumType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.IndirectEntryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "IndirectEntryType", instances, (*IndirectEntryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InlinePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InlinePropertyType", instances, (*InlinePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Integer_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Integer_PropertyType", instances, (*Integer_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.KnotPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "KnotPropertyType", instances, (*KnotPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.KnotTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "KnotType", instances, (*KnotType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LI_Lineage_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LI_Lineage_PropertyType", instances, (*LI_Lineage_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LI_Lineage_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LI_Lineage_Type", instances, (*LI_Lineage_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LI_ProcessStep_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LI_ProcessStep_PropertyType", instances, (*LI_ProcessStep_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LI_ProcessStep_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LI_ProcessStep_Type", instances, (*LI_ProcessStep_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LI_Source_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LI_Source_PropertyType", instances, (*LI_Source_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LI_Source_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LI_Source_Type", instances, (*LI_Source_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LanguageCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LanguageCode_PropertyType", instances, (*LanguageCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Length_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Length_PropertyType", instances, (*Length_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LineStringSegmentArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LineStringSegmentArrayPropertyType", instances, (*LineStringSegmentArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LineStringSegmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LineStringSegmentType", instances, (*LineStringSegmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LineStringTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LineStringType", instances, (*LineStringType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinearCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinearCSPropertyType", instances, (*LinearCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinearCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinearCSType", instances, (*LinearCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinearRingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinearRingPropertyType", instances, (*LinearRingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinearRingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinearRingType", instances, (*LinearRingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalName_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalName_PropertyType", instances, (*LocalName_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalisedCharacterString_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalisedCharacterString_PropertyType", instances, (*LocalisedCharacterString_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocalisedCharacterString_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocalisedCharacterString_Type", instances, (*LocalisedCharacterString_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LocationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LocationPropertyType", instances, (*LocationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_AggregateInformation_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_AggregateInformation_PropertyType", instances, (*MD_AggregateInformation_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_AggregateInformation_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_AggregateInformation_Type", instances, (*MD_AggregateInformation_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_BrowseGraphic_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_BrowseGraphic_PropertyType", instances, (*MD_BrowseGraphic_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_BrowseGraphic_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_BrowseGraphic_Type", instances, (*MD_BrowseGraphic_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_CharacterSetCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_CharacterSetCode_PropertyType", instances, (*MD_CharacterSetCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ClassificationCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ClassificationCode_PropertyType", instances, (*MD_ClassificationCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Constraints_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Constraints_PropertyType", instances, (*MD_Constraints_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Constraints_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Constraints_Type", instances, (*MD_Constraints_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_DataIdentification_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_DataIdentification_PropertyType", instances, (*MD_DataIdentification_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_DataIdentification_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_DataIdentification_Type", instances, (*MD_DataIdentification_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_DigitalTransferOptions_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_DigitalTransferOptions_PropertyType", instances, (*MD_DigitalTransferOptions_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_DigitalTransferOptions_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_DigitalTransferOptions_Type", instances, (*MD_DigitalTransferOptions_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_DistributionUnits_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_DistributionUnits_PropertyType", instances, (*MD_DistributionUnits_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Distribution_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Distribution_PropertyType", instances, (*MD_Distribution_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Distribution_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Distribution_Type", instances, (*MD_Distribution_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Distributor_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Distributor_PropertyType", instances, (*MD_Distributor_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Distributor_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Distributor_Type", instances, (*MD_Distributor_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Format_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Format_PropertyType", instances, (*MD_Format_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Format_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Format_Type", instances, (*MD_Format_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Identification_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Identification_PropertyType", instances, (*MD_Identification_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Identifier_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Identifier_PropertyType", instances, (*MD_Identifier_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Identifier_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Identifier_Type", instances, (*MD_Identifier_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_KeywordTypeCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_KeywordTypeCode_PropertyType", instances, (*MD_KeywordTypeCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Keywords_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Keywords_PropertyType", instances, (*MD_Keywords_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Keywords_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Keywords_Type", instances, (*MD_Keywords_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_LegalConstraints_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_LegalConstraints_PropertyType", instances, (*MD_LegalConstraints_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_LegalConstraints_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_LegalConstraints_Type", instances, (*MD_LegalConstraints_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_MaintenanceFrequencyCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_MaintenanceFrequencyCode_PropertyType", instances, (*MD_MaintenanceFrequencyCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_MaintenanceInformation_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_MaintenanceInformation_PropertyType", instances, (*MD_MaintenanceInformation_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_MaintenanceInformation_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_MaintenanceInformation_Type", instances, (*MD_MaintenanceInformation_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_MediumFormatCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_MediumFormatCode_PropertyType", instances, (*MD_MediumFormatCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_MediumNameCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_MediumNameCode_PropertyType", instances, (*MD_MediumNameCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Medium_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Medium_PropertyType", instances, (*MD_Medium_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Medium_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Medium_Type", instances, (*MD_Medium_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ProgressCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ProgressCode_PropertyType", instances, (*MD_ProgressCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ReferenceSystem_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ReferenceSystem_PropertyType", instances, (*MD_ReferenceSystem_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ReferenceSystem_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ReferenceSystem_Type", instances, (*MD_ReferenceSystem_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_RepresentativeFraction_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_RepresentativeFraction_PropertyType", instances, (*MD_RepresentativeFraction_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_RepresentativeFraction_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_RepresentativeFraction_Type", instances, (*MD_RepresentativeFraction_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Resolution_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Resolution_PropertyType", instances, (*MD_Resolution_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Resolution_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Resolution_Type", instances, (*MD_Resolution_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_RestrictionCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_RestrictionCode_PropertyType", instances, (*MD_RestrictionCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ScopeCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ScopeCode_PropertyType", instances, (*MD_ScopeCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ScopeDescription_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ScopeDescription_PropertyType", instances, (*MD_ScopeDescription_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ScopeDescription_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ScopeDescription_Type", instances, (*MD_ScopeDescription_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_SecurityConstraints_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_SecurityConstraints_PropertyType", instances, (*MD_SecurityConstraints_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_SecurityConstraints_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_SecurityConstraints_Type", instances, (*MD_SecurityConstraints_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ServiceIdentification_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ServiceIdentification_PropertyType", instances, (*MD_ServiceIdentification_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_ServiceIdentification_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_ServiceIdentification_Type", instances, (*MD_ServiceIdentification_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_SpatialRepresentationTypeCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_SpatialRepresentationTypeCode_PropertyType", instances, (*MD_SpatialRepresentationTypeCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_StandardOrderProcess_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_StandardOrderProcess_PropertyType", instances, (*MD_StandardOrderProcess_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_StandardOrderProcess_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_StandardOrderProcess_Type", instances, (*MD_StandardOrderProcess_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_TopicCategoryCode_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_TopicCategoryCode_PropertyType", instances, (*MD_TopicCategoryCode_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Usage_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Usage_PropertyType", instances, (*MD_Usage_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MD_Usage_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MD_Usage_Type", instances, (*MD_Usage_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MappingRuleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MappingRuleType", instances, (*MappingRuleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Measure_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Measure_PropertyType", instances, (*Measure_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MemberName_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MemberName_PropertyType", instances, (*MemberName_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MemberName_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MemberName_Type", instances, (*MemberName_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MetaDataPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MetaDataPropertyType", instances, (*MetaDataPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MovingObjectStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MovingObjectStatusType", instances, (*MovingObjectStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiCurvePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiCurvePropertyType", instances, (*MultiCurvePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiCurveType", instances, (*MultiCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiGeometryPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiGeometryPropertyType", instances, (*MultiGeometryPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiGeometryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiGeometryType", instances, (*MultiGeometryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiPointPropertyType", instances, (*MultiPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiPointType", instances, (*MultiPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiSolidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiSolidPropertyType", instances, (*MultiSolidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiSolidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiSolidType", instances, (*MultiSolidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiSurfacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiSurfacePropertyType", instances, (*MultiSurfacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiSurfaceType", instances, (*MultiSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiplicityRange_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiplicityRange_PropertyType", instances, (*MultiplicityRange_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MultiplicityRange_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MultiplicityRange_Type", instances, (*MultiplicityRange_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Multiplicity_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Multiplicity_PropertyType", instances, (*Multiplicity_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Multiplicity_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Multiplicity_Type", instances, (*Multiplicity_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NodeOrEdgePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NodeOrEdgePropertyType", instances, (*NodeOrEdgePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NodePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NodePropertyType", instances, (*NodePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NodeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NodeType", instances, (*NodeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Number_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Number_PropertyType", instances, (*Number_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObliqueCartesianCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObliqueCartesianCSPropertyType", instances, (*ObliqueCartesianCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObliqueCartesianCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObliqueCartesianCSType", instances, (*ObliqueCartesianCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ObservationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ObservationType", instances, (*ObservationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OffsetCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OffsetCurveType", instances, (*OffsetCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OperationMethodPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OperationMethodPropertyType", instances, (*OperationMethodPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OperationMethodTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OperationMethodType", instances, (*OperationMethodType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OperationParameterGroupPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OperationParameterGroupPropertyType", instances, (*OperationParameterGroupPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OperationParameterGroupTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OperationParameterGroupType", instances, (*OperationParameterGroupType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OperationParameterPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OperationParameterPropertyType", instances, (*OperationParameterPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OperationParameterTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OperationParameterType", instances, (*OperationParameterType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OperationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OperationPropertyType", instances, (*OperationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrientableCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrientableCurveType", instances, (*OrientableCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OrientableSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrientableSurfaceType", instances, (*OrientableSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PT_FreeText_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PT_FreeText_PropertyType", instances, (*PT_FreeText_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PT_FreeText_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PT_FreeText_Type", instances, (*PT_FreeText_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PT_LocaleContainer_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PT_LocaleContainer_PropertyType", instances, (*PT_LocaleContainer_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PT_LocaleContainer_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PT_LocaleContainer_Type", instances, (*PT_LocaleContainer_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PT_Locale_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PT_Locale_PropertyType", instances, (*PT_Locale_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PT_Locale_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PT_Locale_Type", instances, (*PT_Locale_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ParameterValueGroupTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ParameterValueGroupType", instances, (*ParameterValueGroupType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ParameterValueTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ParameterValueType", instances, (*ParameterValueType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassThroughOperationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassThroughOperationPropertyType", instances, (*PassThroughOperationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PassThroughOperationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PassThroughOperationType", instances, (*PassThroughOperationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointArrayPropertyType", instances, (*PointArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointPropertyType", instances, (*PointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PointType", instances, (*PointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PolarCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PolarCSPropertyType", instances, (*PolarCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PolarCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PolarCSType", instances, (*PolarCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PolygonPatchTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PolygonPatchType", instances, (*PolygonPatchType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PolygonTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PolygonType", instances, (*PolygonType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimeMeridianPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimeMeridianPropertyType", instances, (*PrimeMeridianPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrimeMeridianTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrimeMeridianType", instances, (*PrimeMeridianType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PriorityLocationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PriorityLocationPropertyType", instances, (*PriorityLocationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProcedurePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProcedurePropertyType", instances, (*ProcedurePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProjectedCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProjectedCRSPropertyType", instances, (*ProjectedCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ProjectedCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ProjectedCRSType", instances, (*ProjectedCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Quantitys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Quantity", instances, (*Quantity)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.QuantityExtentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "QuantityExtentType", instances, (*QuantityExtentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.QuantityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "QuantityPropertyType", instances, (*QuantityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RS_Identifier_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RS_Identifier_PropertyType", instances, (*RS_Identifier_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RS_Identifier_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RS_Identifier_Type", instances, (*RS_Identifier_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RS_ReferenceSystem_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RS_ReferenceSystem_PropertyType", instances, (*RS_ReferenceSystem_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RangeSetTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RangeSetType", instances, (*RangeSetType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Real_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Real_PropertyType", instances, (*Real_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RecordType_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RecordType_PropertyType", instances, (*RecordType_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RecordType_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RecordType_Type", instances, (*RecordType_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Record_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Record_PropertyType", instances, (*Record_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RectangleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RectangleType", instances, (*RectangleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RectifiedGridTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RectifiedGridType", instances, (*RectifiedGridType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ReferenceType", instances, (*ReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RelatedTimeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RelatedTimeType", instances, (*RelatedTimeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ResultTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ResultType", instances, (*ResultType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RingPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RingPropertyType", instances, (*RingPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RingType", instances, (*RingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SC_CRS_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SC_CRS_PropertyType", instances, (*SC_CRS_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Scale_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Scale_PropertyType", instances, (*Scale_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ScopedName_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ScopedName_PropertyType", instances, (*ScopedName_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondDefiningParameters {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondDefiningParameter", instances, (*SecondDefiningParameter)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SecondDefiningParameter_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SecondDefiningParameter_1", instances, (*SecondDefiningParameter_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SequenceRuleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SequenceRuleType", instances, (*SequenceRuleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShellPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShellPropertyType", instances, (*ShellPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ShellTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ShellType", instances, (*ShellType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SingleCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SingleCRSPropertyType", instances, (*SingleCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SingleOperationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SingleOperationPropertyType", instances, (*SingleOperationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SolidArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SolidArrayPropertyType", instances, (*SolidArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SolidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SolidPropertyType", instances, (*SolidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SolidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SolidType", instances, (*SolidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SphereTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SphereType", instances, (*SphereType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SphericalCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SphericalCSPropertyType", instances, (*SphericalCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SphericalCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SphericalCSType", instances, (*SphericalCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StringOrRefTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StringOrRefType", instances, (*StringOrRefType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceArrayPropertyType", instances, (*SurfaceArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfacePatchArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfacePatchArrayPropertyType", instances, (*SurfacePatchArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfacePropertyType", instances, (*SurfacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceType", instances, (*SurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TM_PeriodDuration_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TM_PeriodDuration_PropertyType", instances, (*TM_PeriodDuration_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TM_Primitive_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TM_Primitive_PropertyType", instances, (*TM_Primitive_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TargetPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TargetPropertyType", instances, (*TargetPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TemporalCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TemporalCRSPropertyType", instances, (*TemporalCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TemporalCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TemporalCRSType", instances, (*TemporalCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TemporalCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TemporalCSPropertyType", instances, (*TemporalCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TemporalCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TemporalCSType", instances, (*TemporalCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TemporalDatumBaseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TemporalDatumBaseType", instances, (*TemporalDatumBaseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TemporalDatumPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TemporalDatumPropertyType", instances, (*TemporalDatumPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TemporalDatumTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TemporalDatumType", instances, (*TemporalDatumType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeCSPropertyType", instances, (*TimeCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeCSType", instances, (*TimeCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeCalendarEraPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeCalendarEraPropertyType", instances, (*TimeCalendarEraPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeCalendarEraTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeCalendarEraType", instances, (*TimeCalendarEraType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeCalendarPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeCalendarPropertyType", instances, (*TimeCalendarPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeCalendarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeCalendarType", instances, (*TimeCalendarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeClockPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeClockPropertyType", instances, (*TimeClockPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeClockTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeClockType", instances, (*TimeClockType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeCoordinateSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeCoordinateSystemType", instances, (*TimeCoordinateSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeEdgePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeEdgePropertyType", instances, (*TimeEdgePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeEdgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeEdgeType", instances, (*TimeEdgeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeInstantPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeInstantPropertyType", instances, (*TimeInstantPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeInstantTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeInstantType", instances, (*TimeInstantType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeIntervalLengthTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeIntervalLengthType", instances, (*TimeIntervalLengthType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeNodePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeNodePropertyType", instances, (*TimeNodePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeNodeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeNodeType", instances, (*TimeNodeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeOrdinalEraPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeOrdinalEraPropertyType", instances, (*TimeOrdinalEraPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeOrdinalEraTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeOrdinalEraType", instances, (*TimeOrdinalEraType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeOrdinalReferenceSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeOrdinalReferenceSystemType", instances, (*TimeOrdinalReferenceSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimePeriodPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimePeriodPropertyType", instances, (*TimePeriodPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimePeriodTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimePeriodType", instances, (*TimePeriodType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimePositionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimePositionType", instances, (*TimePositionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimePrimitivePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimePrimitivePropertyType", instances, (*TimePrimitivePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeReferenceSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeReferenceSystemType", instances, (*TimeReferenceSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeTopologyComplexPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeTopologyComplexPropertyType", instances, (*TimeTopologyComplexPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeTopologyComplexTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeTopologyComplexType", instances, (*TimeTopologyComplexType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeTopologyPrimitivePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeTopologyPrimitivePropertyType", instances, (*TimeTopologyPrimitivePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TinTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TinType", instances, (*TinType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoComplexPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoComplexPropertyType", instances, (*TopoComplexPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoComplexTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoComplexType", instances, (*TopoComplexType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoCurvePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoCurvePropertyType", instances, (*TopoCurvePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoCurveType", instances, (*TopoCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoPointPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoPointPropertyType", instances, (*TopoPointPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoPointType", instances, (*TopoPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoPrimitiveArrayAssociationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoPrimitiveArrayAssociationType", instances, (*TopoPrimitiveArrayAssociationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoPrimitiveMemberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoPrimitiveMemberType", instances, (*TopoPrimitiveMemberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoSolidPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoSolidPropertyType", instances, (*TopoSolidPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoSolidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoSolidType", instances, (*TopoSolidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoSurfacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoSurfacePropertyType", instances, (*TopoSurfacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoSurfaceType", instances, (*TopoSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoVolumePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoVolumePropertyType", instances, (*TopoVolumePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TopoVolumeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TopoVolumeType", instances, (*TopoVolumeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TransformationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TransformationPropertyType", instances, (*TransformationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TransformationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TransformationType", instances, (*TransformationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TriangleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TriangleType", instances, (*TriangleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TypeName_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TypeName_PropertyType", instances, (*TypeName_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TypeName_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TypeName_Type", instances, (*TypeName_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.URL_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "URL_PropertyType", instances, (*URL_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitDefinitionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitDefinitionType", instances, (*UnitDefinitionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitOfMeasureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitOfMeasureType", instances, (*UnitOfMeasureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnitOfMeasure_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnitOfMeasure_PropertyType", instances, (*UnitOfMeasure_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnlimitedInteger_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnlimitedInteger_PropertyType", instances, (*UnlimitedInteger_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UnlimitedInteger_Types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UnlimitedInteger_Type", instances, (*UnlimitedInteger_Type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UomAngle_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UomAngle_PropertyType", instances, (*UomAngle_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UomArea_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UomArea_PropertyType", instances, (*UomArea_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UomLength_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UomLength_PropertyType", instances, (*UomLength_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UomScale_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UomScale_PropertyType", instances, (*UomScale_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UomTime_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UomTime_PropertyType", instances, (*UomTime_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UomVelocity_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UomVelocity_PropertyType", instances, (*UomVelocity_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UomVolume_PropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UomVolume_PropertyType", instances, (*UomVolume_PropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UserDefinedCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UserDefinedCSPropertyType", instances, (*UserDefinedCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.UserDefinedCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "UserDefinedCSType", instances, (*UserDefinedCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValueArrayPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValueArrayPropertyType", instances, (*ValueArrayPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValueArrayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValueArrayType", instances, (*ValueArrayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValuePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValuePropertyType", instances, (*ValuePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VectorType", instances, (*VectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalCRSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalCRSPropertyType", instances, (*VerticalCRSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalCRSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalCRSType", instances, (*VerticalCRSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalCSPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalCSPropertyType", instances, (*VerticalCSPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalCSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalCSType", instances, (*VerticalCSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalDatumPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalDatumPropertyType", instances, (*VerticalDatumPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VerticalDatumTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VerticalDatumType", instances, (*VerticalDatumType)(nil).GongGetFieldHeaders(), addIDs)
		}
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return f
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return f
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Use a lazy iterator instead of loading all rows into memory
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows iterator for sheet %q: %v\n", sheet, err)
			continue
		}

		// Check if there is at least one row, and move the iterator to it
		if !rows.Next() {
			rows.Close() // Always close iterators
			continue
		}

		// Read ONLY the first row
		firstRow, err := rows.Columns()

		// Close the iterator immediately since we don't need the rest of the sheet
		rows.Close()

		if err != nil {
			fmt.Printf("failed to get columns for sheet %q: %v\n", sheet, err)
			continue
		}

		// If the first row is completely empty, skip
		if len(firstRow) == 0 {
			continue
		}

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return f
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return f
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
	}

	var tab GongExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func (stage *Stage) SerializeStageAsBytes(addIDs bool) ([]byte, error) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func __gong__shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// GongTabulator is an interface for writing to a table strings
type GongTabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

type Tabulator = GongTabulator

type GongExcelizeTabulator struct {
	f *excelize.File
}

type ExcelizeTabulator = GongExcelizeTabulator

func (tab *GongExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *GongExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *GongExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *GongExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

// SerializeExcelize is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelize(f *excelize.File, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) {
	sheetName := __gong__shortenString(name)

	// Create a new sheet.
	f.NewSheet(sheetName)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range fields {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), "ID")
				} else {
					// one have to put the type of the cell
					header := fieldHeader.Name
					switch fieldHeader.GongFieldValueType {
					case GongFieldValueTypeInt:
						header += ":int"
					case GongFieldValueTypeIntDuration:
						header += ":duration"
					case GongFieldValueTypeFloat:
						header += ":float"
					case GongFieldValueTypeBool:
						header += ":bool"
					case GongFieldValueTypeString:
						header += ":string"
					case GongFieldValueTypeDate:
						header += ":date"
					default:
						header += ":basicType"
					}
					header += ":noID"
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", GongIntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range fields {
			fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}
}

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type GongstructPtr](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type GongstructPtr](f *excelize.File, addIDs bool) {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	stage.SerializeExcelize(f, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
