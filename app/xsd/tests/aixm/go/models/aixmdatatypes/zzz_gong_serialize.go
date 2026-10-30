// generated code - do not edit
package aixmdatatypes

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
			for instance := range stage.CodeATCReportingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeATCReportingType", instances, (*CodeATCReportingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAerialRefuellingPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAerialRefuellingPointType", instances, (*CodeAerialRefuellingPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAerialRefuellingPrefixTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAerialRefuellingPrefixType", instances, (*CodeAerialRefuellingPrefixType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAerialRefuellingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAerialRefuellingType", instances, (*CodeAerialRefuellingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftCategoryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftCategoryType", instances, (*CodeAircraftCategoryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftEngineNumberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftEngineNumberType", instances, (*CodeAircraftEngineNumberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftEngineTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftEngineType", instances, (*CodeAircraftEngineType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftGroundServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftGroundServiceType", instances, (*CodeAircraftGroundServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftICAOTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftICAOType", instances, (*CodeAircraftICAOType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftStandTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftStandType", instances, (*CodeAircraftStandType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftType", instances, (*CodeAircraftType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAircraftWingspanClassTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAircraftWingspanClassType", instances, (*CodeAircraftWingspanClassType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirportHeliportCollocationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirportHeliportCollocationType", instances, (*CodeAirportHeliportCollocationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirportHeliportDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirportHeliportDesignatorType", instances, (*CodeAirportHeliportDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirportHeliportTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirportHeliportType", instances, (*CodeAirportHeliportType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirportWarningTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirportWarningType", instances, (*CodeAirportWarningType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirspaceActivityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirspaceActivityType", instances, (*CodeAirspaceActivityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirspaceAggregationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirspaceAggregationType", instances, (*CodeAirspaceAggregationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirspaceClassificationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirspaceClassificationType", instances, (*CodeAirspaceClassificationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirspaceDependencyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirspaceDependencyType", instances, (*CodeAirspaceDependencyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirspaceDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirspaceDesignatorType", instances, (*CodeAirspaceDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirspacePointPositionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirspacePointPositionType", instances, (*CodeAirspacePointPositionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirspacePointRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirspacePointRoleType", instances, (*CodeAirspacePointRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAirspaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAirspaceType", instances, (*CodeAirspaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAltitudeAdjustmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAltitudeAdjustmentType", instances, (*CodeAltitudeAdjustmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAltitudeUseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAltitudeUseType", instances, (*CodeAltitudeUseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApproachEquipmentAdditionalTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApproachEquipmentAdditionalType", instances, (*CodeApproachEquipmentAdditionalType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApproachGuidanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApproachGuidanceType", instances, (*CodeApproachGuidanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApproachLightingICAOTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApproachLightingICAOType", instances, (*CodeApproachLightingICAOType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApproachLightingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApproachLightingType", instances, (*CodeApproachLightingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApproachPrefixTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApproachPrefixType", instances, (*CodeApproachPrefixType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApproachTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApproachType", instances, (*CodeApproachType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApprovalTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApprovalType", instances, (*CodeApprovalType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApronElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApronElementType", instances, (*CodeApronElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeApronSectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeApronSectionType", instances, (*CodeApronSectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeArcDirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeArcDirectionType", instances, (*CodeArcDirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeArrestingGearEnergyAbsorbTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeArrestingGearEnergyAbsorbType", instances, (*CodeArrestingGearEnergyAbsorbType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeArrestingGearEngageDeviceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeArrestingGearEngageDeviceType", instances, (*CodeArrestingGearEngageDeviceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAuralMorseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAuralMorseType", instances, (*CodeAuralMorseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAuthorityRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAuthorityRoleType", instances, (*CodeAuthorityRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAuthorityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAuthorityType", instances, (*CodeAuthorityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeAviationStandardsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAviationStandardsType", instances, (*CodeAviationStandardsType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeBearingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeBearingType", instances, (*CodeBearingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeBuoyDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeBuoyDesignatorType", instances, (*CodeBuoyDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeBuoyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeBuoyType", instances, (*CodeBuoyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCardinalDirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCardinalDirectionType", instances, (*CodeCardinalDirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCheckpointCategoryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCheckpointCategoryType", instances, (*CodeCheckpointCategoryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeColourTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeColourType", instances, (*CodeColourType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCommunicationChannelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCommunicationChannelType", instances, (*CodeCommunicationChannelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCommunicationDirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCommunicationDirectionType", instances, (*CodeCommunicationDirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCommunicationModeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCommunicationModeType", instances, (*CodeCommunicationModeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeComparisonTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeComparisonType", instances, (*CodeComparisonType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeContaminationType", instances, (*CodeContaminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCourseQualityILSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCourseQualityILSType", instances, (*CodeCourseQualityILSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeCourseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeCourseType", instances, (*CodeCourseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDMEChannelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDMEChannelType", instances, (*CodeDMEChannelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDMETypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDMEType", instances, (*CodeDMEType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDayType", instances, (*CodeDayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDeclaredDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDeclaredDistanceType", instances, (*CodeDeclaredDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDesignStandardTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDesignStandardType", instances, (*CodeDesignStandardType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDesignatedPointDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDesignatedPointDesignatorType", instances, (*CodeDesignatedPointDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDesignatedPointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDesignatedPointType", instances, (*CodeDesignatedPointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDirectionReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDirectionReferenceType", instances, (*CodeDirectionReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDirectionTurnTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDirectionTurnType", instances, (*CodeDirectionTurnType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDirectionType", instances, (*CodeDirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeDistanceIndicationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeDistanceIndicationType", instances, (*CodeDistanceIndicationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeEmissionBandTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeEmissionBandType", instances, (*CodeEmissionBandType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeEquipmentAntiCollisionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeEquipmentAntiCollisionType", instances, (*CodeEquipmentAntiCollisionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeEquipmentUnavailableTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeEquipmentUnavailableType", instances, (*CodeEquipmentUnavailableType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFacilityRankingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFacilityRankingType", instances, (*CodeFacilityRankingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFinalGuidanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFinalGuidanceType", instances, (*CodeFinalGuidanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFireFightingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFireFightingType", instances, (*CodeFireFightingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightDestinationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightDestinationType", instances, (*CodeFlightDestinationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightOriginTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightOriginType", instances, (*CodeFlightOriginType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightPurposeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightPurposeType", instances, (*CodeFlightPurposeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightRestrictionDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightRestrictionDesignatorType", instances, (*CodeFlightRestrictionDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightRestrictionType", instances, (*CodeFlightRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightRuleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightRuleType", instances, (*CodeFlightRuleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightStatusType", instances, (*CodeFlightStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlightType", instances, (*CodeFlightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFlowConditionOperationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFlowConditionOperationType", instances, (*CodeFlowConditionOperationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFreeFlightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFreeFlightType", instances, (*CodeFreeFlightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFrictionDeviceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFrictionDeviceType", instances, (*CodeFrictionDeviceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFrictionEstimateTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFrictionEstimateType", instances, (*CodeFrictionEstimateType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeFuelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeFuelType", instances, (*CodeFuelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeGeoBorderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeGeoBorderType", instances, (*CodeGeoBorderType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeGradeSeparationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeGradeSeparationType", instances, (*CodeGradeSeparationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeGroundLightingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeGroundLightingType", instances, (*CodeGroundLightingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeGuidanceLineTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeGuidanceLineType", instances, (*CodeGuidanceLineType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeHeightReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeHeightReferenceType", instances, (*CodeHeightReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeHelicopterPerformanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeHelicopterPerformanceType", instances, (*CodeHelicopterPerformanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeHoldingCategoryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeHoldingCategoryType", instances, (*CodeHoldingCategoryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeHoldingUsageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeHoldingUsageType", instances, (*CodeHoldingUsageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeHoldingUseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeHoldingUseType", instances, (*CodeHoldingUseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeIATATypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeIATAType", instances, (*CodeIATAType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeICAOCountryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeICAOCountryType", instances, (*CodeICAOCountryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeICAOTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeICAOType", instances, (*CodeICAOType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeILSBackCourseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeILSBackCourseType", instances, (*CodeILSBackCourseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeIntegrityLevelILSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeIntegrityLevelILSType", instances, (*CodeIntegrityLevelILSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeIntensityStandByTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeIntensityStandByType", instances, (*CodeIntensityStandByType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLanguageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLanguageType", instances, (*CodeLanguageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLevelSeriesTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLevelSeriesType", instances, (*CodeLevelSeriesType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLevelTableDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLevelTableDesignatorType", instances, (*CodeLevelTableDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLevelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLevelType", instances, (*CodeLevelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLightHoldingPositionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLightHoldingPositionType", instances, (*CodeLightHoldingPositionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLightIntensityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLightIntensityType", instances, (*CodeLightIntensityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLightSourceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLightSourceType", instances, (*CodeLightSourceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLightingJARTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLightingJARType", instances, (*CodeLightingJARType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLoadingBridgeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLoadingBridgeType", instances, (*CodeLoadingBridgeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLocationQualifierTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLocationQualifierType", instances, (*CodeLocationQualifierType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeLogicalOperatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLogicalOperatorType", instances, (*CodeLogicalOperatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMLSAzimuthTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMLSAzimuthType", instances, (*CodeMLSAzimuthType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMLSChannelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMLSChannelType", instances, (*CodeMLSChannelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMarkerBeaconSignalTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMarkerBeaconSignalType", instances, (*CodeMarkerBeaconSignalType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMarkingConditionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMarkingConditionType", instances, (*CodeMarkingConditionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMarkingStyleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMarkingStyleType", instances, (*CodeMarkingStyleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMeteoConditionsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMeteoConditionsType", instances, (*CodeMeteoConditionsType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMilitaryOperationsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMilitaryOperationsType", instances, (*CodeMilitaryOperationsType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMilitaryRoutePointTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMilitaryRoutePointType", instances, (*CodeMilitaryRoutePointType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMilitaryStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMilitaryStatusType", instances, (*CodeMilitaryStatusType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMilitaryTrainingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMilitaryTrainingType", instances, (*CodeMilitaryTrainingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMinimaFinalApproachPathTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMinimaFinalApproachPathType", instances, (*CodeMinimaFinalApproachPathType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMinimumAltitudeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMinimumAltitudeType", instances, (*CodeMinimumAltitudeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMinimumHeightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMinimumHeightType", instances, (*CodeMinimumHeightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeMissedApproachTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMissedApproachType", instances, (*CodeMissedApproachType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNDBUsageTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNDBUsageType", instances, (*CodeNDBUsageType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavaidDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavaidDesignatorType", instances, (*CodeNavaidDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavaidPurposeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavaidPurposeType", instances, (*CodeNavaidPurposeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavaidServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavaidServiceType", instances, (*CodeNavaidServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavigationAreaRestrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavigationAreaRestrictionType", instances, (*CodeNavigationAreaRestrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavigationAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavigationAreaType", instances, (*CodeNavigationAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavigationEquipmentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavigationEquipmentType", instances, (*CodeNavigationEquipmentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNavigationSpecificationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNavigationSpecificationType", instances, (*CodeNavigationSpecificationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNitrogenTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNitrogenType", instances, (*CodeNitrogenType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNorthReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNorthReferenceType", instances, (*CodeNorthReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeNotePurposeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNotePurposeType", instances, (*CodeNotePurposeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeObstacleAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeObstacleAreaType", instances, (*CodeObstacleAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeObstacleAssessmentSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeObstacleAssessmentSurfaceType", instances, (*CodeObstacleAssessmentSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeObstructionIdSurfaceZoneTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeObstructionIdSurfaceZoneType", instances, (*CodeObstructionIdSurfaceZoneType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeOilTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeOilType", instances, (*CodeOilType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeOperationAirportHeliportTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeOperationAirportHeliportType", instances, (*CodeOperationAirportHeliportType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeOperationManoeuvringAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeOperationManoeuvringAreaType", instances, (*CodeOperationManoeuvringAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeOrganisationDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeOrganisationDesignatorType", instances, (*CodeOrganisationDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeOrganisationHierarchyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeOrganisationHierarchyType", instances, (*CodeOrganisationHierarchyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeOrganisationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeOrganisationType", instances, (*CodeOrganisationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeOxygenTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeOxygenType", instances, (*CodeOxygenType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePARTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePARType", instances, (*CodePARType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePCNMethodTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePCNMethodType", instances, (*CodePCNMethodType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePCNPavementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePCNPavementType", instances, (*CodePCNPavementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePCNSubgradeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePCNSubgradeType", instances, (*CodePCNSubgradeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePCNTyrePressureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePCNTyrePressureType", instances, (*CodePCNTyrePressureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePassengerServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePassengerServiceType", instances, (*CodePassengerServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePilotControlledLightingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePilotControlledLightingType", instances, (*CodePilotControlledLightingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePositionInILSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePositionInILSType", instances, (*CodePositionInILSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodePrimaryRadarTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodePrimaryRadarType", instances, (*CodePrimaryRadarType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeProcedureAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeProcedureAvailabilityType", instances, (*CodeProcedureAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeProcedureCodingStandardTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeProcedureCodingStandardType", instances, (*CodeProcedureCodingStandardType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeProcedureDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeProcedureDistanceType", instances, (*CodeProcedureDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeProcedureFixRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeProcedureFixRoleType", instances, (*CodeProcedureFixRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeProcedurePhaseTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeProcedurePhaseType", instances, (*CodeProcedurePhaseType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeProtectAreaSectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeProtectAreaSectionType", instances, (*CodeProtectAreaSectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRNPTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRNPType", instances, (*CodeRNPType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRVRReadingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRVRReadingType", instances, (*CodeRVRReadingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRVSMPointRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRVSMPointRoleType", instances, (*CodeRVSMPointRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRVSMTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRVSMType", instances, (*CodeRVSMType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRadarServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRadarServiceType", instances, (*CodeRadarServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRadioEmissionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRadioEmissionType", instances, (*CodeRadioEmissionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRadioFrequencyAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRadioFrequencyAreaType", instances, (*CodeRadioFrequencyAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRadioSignalTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRadioSignalType", instances, (*CodeRadioSignalType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeReferencePathIdentifierTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeReferencePathIdentifierType", instances, (*CodeReferencePathIdentifierType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeReferenceRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeReferenceRoleType", instances, (*CodeReferenceRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeReflectorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeReflectorType", instances, (*CodeReflectorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRelativePositionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRelativePositionType", instances, (*CodeRelativePositionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRoadTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRoadType", instances, (*CodeRoadType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteAvailabilityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteAvailabilityType", instances, (*CodeRouteAvailabilityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteDesignatorLetterTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteDesignatorLetterType", instances, (*CodeRouteDesignatorLetterType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteDesignatorPrefixTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteDesignatorPrefixType", instances, (*CodeRouteDesignatorPrefixType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteDesignatorSuffixTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteDesignatorSuffixType", instances, (*CodeRouteDesignatorSuffixType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteIndicatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteIndicatorType", instances, (*CodeRouteIndicatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteNavigationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteNavigationType", instances, (*CodeRouteNavigationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteOriginTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteOriginType", instances, (*CodeRouteOriginType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteSegmentPathTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteSegmentPathType", instances, (*CodeRouteSegmentPathType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRouteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRouteType", instances, (*CodeRouteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRuleProcedureTitleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRuleProcedureTitleType", instances, (*CodeRuleProcedureTitleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRuleProcedureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRuleProcedureType", instances, (*CodeRuleProcedureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRunwayElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRunwayElementType", instances, (*CodeRunwayElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRunwayMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRunwayMarkingType", instances, (*CodeRunwayMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRunwayPointRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRunwayPointRoleType", instances, (*CodeRunwayPointRoleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRunwayProtectionAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRunwayProtectionAreaType", instances, (*CodeRunwayProtectionAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRunwaySectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRunwaySectionType", instances, (*CodeRunwaySectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeRunwayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRunwayType", instances, (*CodeRunwayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSafeAltitudeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSafeAltitudeType", instances, (*CodeSafeAltitudeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSegmentPathTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSegmentPathType", instances, (*CodeSegmentPathType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSegmentTerminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSegmentTerminationType", instances, (*CodeSegmentTerminationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeServiceATCTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeServiceATCType", instances, (*CodeServiceATCType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeServiceATFMTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeServiceATFMType", instances, (*CodeServiceATFMType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeServiceGroundControlTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeServiceGroundControlType", instances, (*CodeServiceGroundControlType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeServiceInformationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeServiceInformationType", instances, (*CodeServiceInformationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeServiceSARTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeServiceSARType", instances, (*CodeServiceSARType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSideTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSideType", instances, (*CodeSideType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSignalPerformanceILSTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSignalPerformanceILSType", instances, (*CodeSignalPerformanceILSType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSpecialDateTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSpecialDateType", instances, (*CodeSpecialDateType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSpecialNavigationChainDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSpecialNavigationChainDesignatorType", instances, (*CodeSpecialNavigationChainDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSpecialNavigationStationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSpecialNavigationStationType", instances, (*CodeSpecialNavigationStationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSpecialNavigationSystemTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSpecialNavigationSystemType", instances, (*CodeSpecialNavigationSystemType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSpeedReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSpeedReferenceType", instances, (*CodeSpeedReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeStandbyPowerTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeStandbyPowerType", instances, (*CodeStandbyPowerType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeStatusAirportTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeStatusAirportType", instances, (*CodeStatusAirportType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeStatusAirspaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeStatusAirspaceType", instances, (*CodeStatusAirspaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeStatusConstructionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeStatusConstructionType", instances, (*CodeStatusConstructionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeStatusNavaidTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeStatusNavaidType", instances, (*CodeStatusNavaidType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeStatusOperationsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeStatusOperationsType", instances, (*CodeStatusOperationsType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeStatusServiceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeStatusServiceType", instances, (*CodeStatusServiceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSurfaceCompositionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSurfaceCompositionType", instances, (*CodeSurfaceCompositionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSurfaceConditionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSurfaceConditionType", instances, (*CodeSurfaceConditionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSurfacePreparationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSurfacePreparationType", instances, (*CodeSurfacePreparationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeSystemActivationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSystemActivationType", instances, (*CodeSystemActivationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTAATypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTAAType", instances, (*CodeTAAType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTACANChannelTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTACANChannelType", instances, (*CodeTACANChannelType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTLOFSectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTLOFSectionType", instances, (*CodeTLOFSectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTaxiwayElementTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTaxiwayElementType", instances, (*CodeTaxiwayElementType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTaxiwaySectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTaxiwaySectionType", instances, (*CodeTaxiwaySectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTaxiwayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTaxiwayType", instances, (*CodeTaxiwayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTelecomNetworkTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTelecomNetworkType", instances, (*CodeTelecomNetworkType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTimeEventCombinationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTimeEventCombinationType", instances, (*CodeTimeEventCombinationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTimeEventTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTimeEventType", instances, (*CodeTimeEventType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTimeReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTimeReferenceType", instances, (*CodeTimeReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTrajectoryTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTrajectoryType", instances, (*CodeTrajectoryType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeTransponderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTransponderType", instances, (*CodeTransponderType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeUUIDTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeUUIDType", instances, (*CodeUUIDType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeUnitDependencyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeUnitDependencyType", instances, (*CodeUnitDependencyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeUnitTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeUnitType", instances, (*CodeUnitType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeUpperAlphaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeUpperAlphaType", instances, (*CodeUpperAlphaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeUsageLimitationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeUsageLimitationType", instances, (*CodeUsageLimitationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVASISTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVASISType", instances, (*CodeVASISType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVORTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVORType", instances, (*CodeVORType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeValueInterpretationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeValueInterpretationType", instances, (*CodeValueInterpretationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVerticalDatumTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVerticalDatumType", instances, (*CodeVerticalDatumType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVerticalReferenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVerticalReferenceType", instances, (*CodeVerticalReferenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVerticalStructureMarkingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVerticalStructureMarkingType", instances, (*CodeVerticalStructureMarkingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVerticalStructureMaterialTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVerticalStructureMaterialType", instances, (*CodeVerticalStructureMaterialType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVerticalStructureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVerticalStructureType", instances, (*CodeVerticalStructureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeVisualDockingGuidanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeVisualDockingGuidanceType", instances, (*CodeVisualDockingGuidanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeWakeTurbulenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeWakeTurbulenceType", instances, (*CodeWakeTurbulenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeWorkAreaTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeWorkAreaType", instances, (*CodeWorkAreaType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CodeYesNoTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeYesNoType", instances, (*CodeYesNoType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DateMonthDayTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DateMonthDayType", instances, (*DateMonthDayType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DateTimeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DateTimeType", instances, (*DateTimeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DateTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DateType", instances, (*DateType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DateYearTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DateYearType", instances, (*DateYearType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoNumberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoNumberType", instances, (*NoNumberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoSequenceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoSequenceType", instances, (*NoSequenceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextAddressTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextAddressType", instances, (*TextAddressType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextDesignatorType", instances, (*TextDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextInstructionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextInstructionType", instances, (*TextInstructionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextNameTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextNameType", instances, (*TextNameType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextNoteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextNoteType", instances, (*TextNoteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextPhoneTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextPhoneType", instances, (*TextPhoneType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextPropertyNameTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextPropertyNameType", instances, (*TextPropertyNameType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextRemarkTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextRemarkType", instances, (*TextRemarkType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextSIDSTARDesignatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextSIDSTARDesignatorType", instances, (*TextSIDSTARDesignatorType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeType", instances, (*TimeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValAlarmLimitTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValAlarmLimitType", instances, (*ValAlarmLimitType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValAngleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValAngleType", instances, (*ValAngleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValBearingTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValBearingType", instances, (*ValBearingType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValChannelNumberTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValChannelNumberType", instances, (*ValChannelNumberType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDepthTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDepthType", instances, (*ValDepthType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDistanceSignedTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDistanceSignedType", instances, (*ValDistanceSignedType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDistanceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDistanceType", instances, (*ValDistanceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDistanceVerticalTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDistanceVerticalType", instances, (*ValDistanceVerticalType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValDurationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValDurationType", instances, (*ValDurationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValFLTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValFLType", instances, (*ValFLType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValFrequencyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValFrequencyType", instances, (*ValFrequencyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValFrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValFrictionType", instances, (*ValFrictionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValHexTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValHexType", instances, (*ValHexType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValLCNTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValLCNType", instances, (*ValLCNType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValLightIntensityTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValLightIntensityType", instances, (*ValLightIntensityType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValMagneticVariationChangeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValMagneticVariationChangeType", instances, (*ValMagneticVariationChangeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValMagneticVariationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValMagneticVariationType", instances, (*ValMagneticVariationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValPCNTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValPCNType", instances, (*ValPCNType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValPercentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValPercentType", instances, (*ValPercentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValPressureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValPressureType", instances, (*ValPressureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValSlopeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValSlopeType", instances, (*ValSlopeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValSpeedTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValSpeedType", instances, (*ValSpeedType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValTemperatureTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValTemperatureType", instances, (*ValTemperatureType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ValWeightTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValWeightType", instances, (*ValWeightType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.XHTMLTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "XHTMLType", instances, (*XHTMLType)(nil).GongGetFieldHeaders(), addIDs)
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
