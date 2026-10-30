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
			for instance := range stage.CodeAuthorityRoleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeAuthorityRoleType", instances, (*CodeAuthorityRoleType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeColourTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeColourType", instances, (*CodeColourType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeContaminationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeContaminationType", instances, (*CodeContaminationType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeEquipmentAntiCollisionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeEquipmentAntiCollisionType", instances, (*CodeEquipmentAntiCollisionType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeGradeSeparationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeGradeSeparationType", instances, (*CodeGradeSeparationType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeIATATypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeIATAType", instances, (*CodeIATAType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeLogicalOperatorTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeLogicalOperatorType", instances, (*CodeLogicalOperatorType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeMilitaryStatusTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeMilitaryStatusType", instances, (*CodeMilitaryStatusType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeNotePurposeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeNotePurposeType", instances, (*CodeNotePurposeType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeProtectAreaSectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeProtectAreaSectionType", instances, (*CodeProtectAreaSectionType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeRVSMTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeRVSMType", instances, (*CodeRVSMType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeSideTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeSideType", instances, (*CodeSideType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeStatusOperationsTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeStatusOperationsType", instances, (*CodeStatusOperationsType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.CodeTransponderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CodeTransponderType", instances, (*CodeTransponderType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.TimeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeType", instances, (*TimeType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.ValFrictionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ValFrictionType", instances, (*ValFrictionType)(nil).GongGetFieldHeaders(), addIDs)
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
