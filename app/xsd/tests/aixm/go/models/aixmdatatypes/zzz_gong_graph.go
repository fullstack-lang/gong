// generated code - do not edit
package aixmdatatypes

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
func (codeatcreportingtype *CodeATCReportingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeATCReportingTypes[codeatcreportingtype]
	return ok
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAerialRefuellingPointTypes[codeaerialrefuellingpointtype]
	return ok
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAerialRefuellingPrefixTypes[codeaerialrefuellingprefixtype]
	return ok
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAerialRefuellingTypes[codeaerialrefuellingtype]
	return ok
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftCategoryTypes[codeaircraftcategorytype]
	return ok
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftEngineNumberTypes[codeaircraftenginenumbertype]
	return ok
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftEngineTypes[codeaircraftenginetype]
	return ok
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftGroundServiceTypes[codeaircraftgroundservicetype]
	return ok
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftICAOTypes[codeaircrafticaotype]
	return ok
}

func (codeaircraftstandtype *CodeAircraftStandType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftStandTypes[codeaircraftstandtype]
	return ok
}

func (codeaircrafttype *CodeAircraftType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftTypes[codeaircrafttype]
	return ok
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftWingspanClassTypes[codeaircraftwingspanclasstype]
	return ok
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirportHeliportCollocationTypes[codeairportheliportcollocationtype]
	return ok
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirportHeliportDesignatorTypes[codeairportheliportdesignatortype]
	return ok
}

func (codeairportheliporttype *CodeAirportHeliportType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirportHeliportTypes[codeairportheliporttype]
	return ok
}

func (codeairportwarningtype *CodeAirportWarningType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirportWarningTypes[codeairportwarningtype]
	return ok
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceActivityTypes[codeairspaceactivitytype]
	return ok
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceAggregationTypes[codeairspaceaggregationtype]
	return ok
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceClassificationTypes[codeairspaceclassificationtype]
	return ok
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceDependencyTypes[codeairspacedependencytype]
	return ok
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceDesignatorTypes[codeairspacedesignatortype]
	return ok
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspacePointPositionTypes[codeairspacepointpositiontype]
	return ok
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspacePointRoleTypes[codeairspacepointroletype]
	return ok
}

func (codeairspacetype *CodeAirspaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceTypes[codeairspacetype]
	return ok
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAltitudeAdjustmentTypes[codealtitudeadjustmenttype]
	return ok
}

func (codealtitudeusetype *CodeAltitudeUseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAltitudeUseTypes[codealtitudeusetype]
	return ok
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApproachEquipmentAdditionalTypes[codeapproachequipmentadditionaltype]
	return ok
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApproachGuidanceTypes[codeapproachguidancetype]
	return ok
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApproachLightingICAOTypes[codeapproachlightingicaotype]
	return ok
}

func (codeapproachlightingtype *CodeApproachLightingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApproachLightingTypes[codeapproachlightingtype]
	return ok
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApproachPrefixTypes[codeapproachprefixtype]
	return ok
}

func (codeapproachtype *CodeApproachType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApproachTypes[codeapproachtype]
	return ok
}

func (codeapprovaltype *CodeApprovalType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApprovalTypes[codeapprovaltype]
	return ok
}

func (codeapronelementtype *CodeApronElementType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApronElementTypes[codeapronelementtype]
	return ok
}

func (codeapronsectiontype *CodeApronSectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApronSectionTypes[codeapronsectiontype]
	return ok
}

func (codearcdirectiontype *CodeArcDirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeArcDirectionTypes[codearcdirectiontype]
	return ok
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeArrestingGearEnergyAbsorbTypes[codearrestinggearenergyabsorbtype]
	return ok
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeArrestingGearEngageDeviceTypes[codearrestinggearengagedevicetype]
	return ok
}

func (codeauralmorsetype *CodeAuralMorseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAuralMorseTypes[codeauralmorsetype]
	return ok
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAuthorityRoleTypes[codeauthorityroletype]
	return ok
}

func (codeauthoritytype *CodeAuthorityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAuthorityTypes[codeauthoritytype]
	return ok
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAviationStandardsTypes[codeaviationstandardstype]
	return ok
}

func (codebearingtype *CodeBearingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeBearingTypes[codebearingtype]
	return ok
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeBuoyDesignatorTypes[codebuoydesignatortype]
	return ok
}

func (codebuoytype *CodeBuoyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeBuoyTypes[codebuoytype]
	return ok
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCardinalDirectionTypes[codecardinaldirectiontype]
	return ok
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCheckpointCategoryTypes[codecheckpointcategorytype]
	return ok
}

func (codecolourtype *CodeColourType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeColourTypes[codecolourtype]
	return ok
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCommunicationChannelTypes[codecommunicationchanneltype]
	return ok
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCommunicationDirectionTypes[codecommunicationdirectiontype]
	return ok
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCommunicationModeTypes[codecommunicationmodetype]
	return ok
}

func (codecomparisontype *CodeComparisonType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeComparisonTypes[codecomparisontype]
	return ok
}

func (codecontaminationtype *CodeContaminationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeContaminationTypes[codecontaminationtype]
	return ok
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCourseQualityILSTypes[codecoursequalityilstype]
	return ok
}

func (codecoursetype *CodeCourseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCourseTypes[codecoursetype]
	return ok
}

func (codedmechanneltype *CodeDMEChannelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDMEChannelTypes[codedmechanneltype]
	return ok
}

func (codedmetype *CodeDMEType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDMETypes[codedmetype]
	return ok
}

func (codedaytype *CodeDayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDayTypes[codedaytype]
	return ok
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDeclaredDistanceTypes[codedeclareddistancetype]
	return ok
}

func (codedesignstandardtype *CodeDesignStandardType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDesignStandardTypes[codedesignstandardtype]
	return ok
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDesignatedPointDesignatorTypes[codedesignatedpointdesignatortype]
	return ok
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDesignatedPointTypes[codedesignatedpointtype]
	return ok
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDirectionReferenceTypes[codedirectionreferencetype]
	return ok
}

func (codedirectionturntype *CodeDirectionTurnType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDirectionTurnTypes[codedirectionturntype]
	return ok
}

func (codedirectiontype *CodeDirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDirectionTypes[codedirectiontype]
	return ok
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDistanceIndicationTypes[codedistanceindicationtype]
	return ok
}

func (codeemissionbandtype *CodeEmissionBandType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeEmissionBandTypes[codeemissionbandtype]
	return ok
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeEquipmentAntiCollisionTypes[codeequipmentanticollisiontype]
	return ok
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeEquipmentUnavailableTypes[codeequipmentunavailabletype]
	return ok
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFacilityRankingTypes[codefacilityrankingtype]
	return ok
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFinalGuidanceTypes[codefinalguidancetype]
	return ok
}

func (codefirefightingtype *CodeFireFightingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFireFightingTypes[codefirefightingtype]
	return ok
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightDestinationTypes[codeflightdestinationtype]
	return ok
}

func (codeflightorigintype *CodeFlightOriginType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightOriginTypes[codeflightorigintype]
	return ok
}

func (codeflightpurposetype *CodeFlightPurposeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightPurposeTypes[codeflightpurposetype]
	return ok
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightRestrictionDesignatorTypes[codeflightrestrictiondesignatortype]
	return ok
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightRestrictionTypes[codeflightrestrictiontype]
	return ok
}

func (codeflightruletype *CodeFlightRuleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightRuleTypes[codeflightruletype]
	return ok
}

func (codeflightstatustype *CodeFlightStatusType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightStatusTypes[codeflightstatustype]
	return ok
}

func (codeflighttype *CodeFlightType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightTypes[codeflighttype]
	return ok
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlowConditionOperationTypes[codeflowconditionoperationtype]
	return ok
}

func (codefreeflighttype *CodeFreeFlightType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFreeFlightTypes[codefreeflighttype]
	return ok
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFrictionDeviceTypes[codefrictiondevicetype]
	return ok
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFrictionEstimateTypes[codefrictionestimatetype]
	return ok
}

func (codefueltype *CodeFuelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFuelTypes[codefueltype]
	return ok
}

func (codegeobordertype *CodeGeoBorderType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeGeoBorderTypes[codegeobordertype]
	return ok
}

func (codegradeseparationtype *CodeGradeSeparationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeGradeSeparationTypes[codegradeseparationtype]
	return ok
}

func (codegroundlightingtype *CodeGroundLightingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeGroundLightingTypes[codegroundlightingtype]
	return ok
}

func (codeguidancelinetype *CodeGuidanceLineType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeGuidanceLineTypes[codeguidancelinetype]
	return ok
}

func (codeheightreferencetype *CodeHeightReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeHeightReferenceTypes[codeheightreferencetype]
	return ok
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeHelicopterPerformanceTypes[codehelicopterperformancetype]
	return ok
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeHoldingCategoryTypes[codeholdingcategorytype]
	return ok
}

func (codeholdingusagetype *CodeHoldingUsageType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeHoldingUsageTypes[codeholdingusagetype]
	return ok
}

func (codeholdingusetype *CodeHoldingUseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeHoldingUseTypes[codeholdingusetype]
	return ok
}

func (codeiatatype *CodeIATAType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeIATATypes[codeiatatype]
	return ok
}

func (codeicaocountrytype *CodeICAOCountryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeICAOCountryTypes[codeicaocountrytype]
	return ok
}

func (codeicaotype *CodeICAOType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeICAOTypes[codeicaotype]
	return ok
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeILSBackCourseTypes[codeilsbackcoursetype]
	return ok
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeIntegrityLevelILSTypes[codeintegritylevelilstype]
	return ok
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeIntensityStandByTypes[codeintensitystandbytype]
	return ok
}

func (codelanguagetype *CodeLanguageType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLanguageTypes[codelanguagetype]
	return ok
}

func (codelevelseriestype *CodeLevelSeriesType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLevelSeriesTypes[codelevelseriestype]
	return ok
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLevelTableDesignatorTypes[codeleveltabledesignatortype]
	return ok
}

func (codeleveltype *CodeLevelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLevelTypes[codeleveltype]
	return ok
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLightHoldingPositionTypes[codelightholdingpositiontype]
	return ok
}

func (codelightintensitytype *CodeLightIntensityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLightIntensityTypes[codelightintensitytype]
	return ok
}

func (codelightsourcetype *CodeLightSourceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLightSourceTypes[codelightsourcetype]
	return ok
}

func (codelightingjartype *CodeLightingJARType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLightingJARTypes[codelightingjartype]
	return ok
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLoadingBridgeTypes[codeloadingbridgetype]
	return ok
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLocationQualifierTypes[codelocationqualifiertype]
	return ok
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLogicalOperatorTypes[codelogicaloperatortype]
	return ok
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMLSAzimuthTypes[codemlsazimuthtype]
	return ok
}

func (codemlschanneltype *CodeMLSChannelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMLSChannelTypes[codemlschanneltype]
	return ok
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMarkerBeaconSignalTypes[codemarkerbeaconsignaltype]
	return ok
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMarkingConditionTypes[codemarkingconditiontype]
	return ok
}

func (codemarkingstyletype *CodeMarkingStyleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMarkingStyleTypes[codemarkingstyletype]
	return ok
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMeteoConditionsTypes[codemeteoconditionstype]
	return ok
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMilitaryOperationsTypes[codemilitaryoperationstype]
	return ok
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMilitaryRoutePointTypes[codemilitaryroutepointtype]
	return ok
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMilitaryStatusTypes[codemilitarystatustype]
	return ok
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMilitaryTrainingTypes[codemilitarytrainingtype]
	return ok
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMinimaFinalApproachPathTypes[codeminimafinalapproachpathtype]
	return ok
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMinimumAltitudeTypes[codeminimumaltitudetype]
	return ok
}

func (codeminimumheighttype *CodeMinimumHeightType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMinimumHeightTypes[codeminimumheighttype]
	return ok
}

func (codemissedapproachtype *CodeMissedApproachType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMissedApproachTypes[codemissedapproachtype]
	return ok
}

func (codendbusagetype *CodeNDBUsageType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNDBUsageTypes[codendbusagetype]
	return ok
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavaidDesignatorTypes[codenavaiddesignatortype]
	return ok
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavaidPurposeTypes[codenavaidpurposetype]
	return ok
}

func (codenavaidservicetype *CodeNavaidServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavaidServiceTypes[codenavaidservicetype]
	return ok
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavigationAreaRestrictionTypes[codenavigationarearestrictiontype]
	return ok
}

func (codenavigationareatype *CodeNavigationAreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavigationAreaTypes[codenavigationareatype]
	return ok
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavigationEquipmentTypes[codenavigationequipmenttype]
	return ok
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavigationSpecificationTypes[codenavigationspecificationtype]
	return ok
}

func (codenitrogentype *CodeNitrogenType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNitrogenTypes[codenitrogentype]
	return ok
}

func (codenorthreferencetype *CodeNorthReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNorthReferenceTypes[codenorthreferencetype]
	return ok
}

func (codenotepurposetype *CodeNotePurposeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNotePurposeTypes[codenotepurposetype]
	return ok
}

func (codeobstacleareatype *CodeObstacleAreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeObstacleAreaTypes[codeobstacleareatype]
	return ok
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeObstacleAssessmentSurfaceTypes[codeobstacleassessmentsurfacetype]
	return ok
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeObstructionIdSurfaceZoneTypes[codeobstructionidsurfacezonetype]
	return ok
}

func (codeoiltype *CodeOilType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOilTypes[codeoiltype]
	return ok
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOperationAirportHeliportTypes[codeoperationairportheliporttype]
	return ok
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOperationManoeuvringAreaTypes[codeoperationmanoeuvringareatype]
	return ok
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOrganisationDesignatorTypes[codeorganisationdesignatortype]
	return ok
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOrganisationHierarchyTypes[codeorganisationhierarchytype]
	return ok
}

func (codeorganisationtype *CodeOrganisationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOrganisationTypes[codeorganisationtype]
	return ok
}

func (codeoxygentype *CodeOxygenType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOxygenTypes[codeoxygentype]
	return ok
}

func (codepartype *CodePARType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePARTypes[codepartype]
	return ok
}

func (codepcnmethodtype *CodePCNMethodType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePCNMethodTypes[codepcnmethodtype]
	return ok
}

func (codepcnpavementtype *CodePCNPavementType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePCNPavementTypes[codepcnpavementtype]
	return ok
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePCNSubgradeTypes[codepcnsubgradetype]
	return ok
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePCNTyrePressureTypes[codepcntyrepressuretype]
	return ok
}

func (codepassengerservicetype *CodePassengerServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePassengerServiceTypes[codepassengerservicetype]
	return ok
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePilotControlledLightingTypes[codepilotcontrolledlightingtype]
	return ok
}

func (codepositioninilstype *CodePositionInILSType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePositionInILSTypes[codepositioninilstype]
	return ok
}

func (codeprimaryradartype *CodePrimaryRadarType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePrimaryRadarTypes[codeprimaryradartype]
	return ok
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeProcedureAvailabilityTypes[codeprocedureavailabilitytype]
	return ok
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeProcedureCodingStandardTypes[codeprocedurecodingstandardtype]
	return ok
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeProcedureDistanceTypes[codeproceduredistancetype]
	return ok
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeProcedureFixRoleTypes[codeprocedurefixroletype]
	return ok
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeProcedurePhaseTypes[codeprocedurephasetype]
	return ok
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeProtectAreaSectionTypes[codeprotectareasectiontype]
	return ok
}

func (codernptype *CodeRNPType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRNPTypes[codernptype]
	return ok
}

func (codervrreadingtype *CodeRVRReadingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRVRReadingTypes[codervrreadingtype]
	return ok
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRVSMPointRoleTypes[codervsmpointroletype]
	return ok
}

func (codervsmtype *CodeRVSMType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRVSMTypes[codervsmtype]
	return ok
}

func (coderadarservicetype *CodeRadarServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRadarServiceTypes[coderadarservicetype]
	return ok
}

func (coderadioemissiontype *CodeRadioEmissionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRadioEmissionTypes[coderadioemissiontype]
	return ok
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRadioFrequencyAreaTypes[coderadiofrequencyareatype]
	return ok
}

func (coderadiosignaltype *CodeRadioSignalType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRadioSignalTypes[coderadiosignaltype]
	return ok
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeReferencePathIdentifierTypes[codereferencepathidentifiertype]
	return ok
}

func (codereferenceroletype *CodeReferenceRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeReferenceRoleTypes[codereferenceroletype]
	return ok
}

func (codereflectortype *CodeReflectorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeReflectorTypes[codereflectortype]
	return ok
}

func (coderelativepositiontype *CodeRelativePositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRelativePositionTypes[coderelativepositiontype]
	return ok
}

func (coderoadtype *CodeRoadType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRoadTypes[coderoadtype]
	return ok
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteAvailabilityTypes[coderouteavailabilitytype]
	return ok
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteDesignatorLetterTypes[coderoutedesignatorlettertype]
	return ok
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteDesignatorPrefixTypes[coderoutedesignatorprefixtype]
	return ok
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteDesignatorSuffixTypes[coderoutedesignatorsuffixtype]
	return ok
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteIndicatorTypes[coderouteindicatortype]
	return ok
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteNavigationTypes[coderoutenavigationtype]
	return ok
}

func (coderouteorigintype *CodeRouteOriginType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteOriginTypes[coderouteorigintype]
	return ok
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteSegmentPathTypes[coderoutesegmentpathtype]
	return ok
}

func (coderoutetype *CodeRouteType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteTypes[coderoutetype]
	return ok
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRuleProcedureTitleTypes[coderuleproceduretitletype]
	return ok
}

func (coderuleproceduretype *CodeRuleProcedureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRuleProcedureTypes[coderuleproceduretype]
	return ok
}

func (coderunwayelementtype *CodeRunwayElementType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRunwayElementTypes[coderunwayelementtype]
	return ok
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRunwayMarkingTypes[coderunwaymarkingtype]
	return ok
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRunwayPointRoleTypes[coderunwaypointroletype]
	return ok
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRunwayProtectionAreaTypes[coderunwayprotectionareatype]
	return ok
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRunwaySectionTypes[coderunwaysectiontype]
	return ok
}

func (coderunwaytype *CodeRunwayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRunwayTypes[coderunwaytype]
	return ok
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSafeAltitudeTypes[codesafealtitudetype]
	return ok
}

func (codesegmentpathtype *CodeSegmentPathType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSegmentPathTypes[codesegmentpathtype]
	return ok
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSegmentTerminationTypes[codesegmentterminationtype]
	return ok
}

func (codeserviceatctype *CodeServiceATCType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeServiceATCTypes[codeserviceatctype]
	return ok
}

func (codeserviceatfmtype *CodeServiceATFMType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeServiceATFMTypes[codeserviceatfmtype]
	return ok
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeServiceGroundControlTypes[codeservicegroundcontroltype]
	return ok
}

func (codeserviceinformationtype *CodeServiceInformationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeServiceInformationTypes[codeserviceinformationtype]
	return ok
}

func (codeservicesartype *CodeServiceSARType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeServiceSARTypes[codeservicesartype]
	return ok
}

func (codesidetype *CodeSideType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSideTypes[codesidetype]
	return ok
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSignalPerformanceILSTypes[codesignalperformanceilstype]
	return ok
}

func (codespecialdatetype *CodeSpecialDateType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpecialDateTypes[codespecialdatetype]
	return ok
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpecialNavigationChainDesignatorTypes[codespecialnavigationchaindesignatortype]
	return ok
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpecialNavigationStationTypes[codespecialnavigationstationtype]
	return ok
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpecialNavigationSystemTypes[codespecialnavigationsystemtype]
	return ok
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpeedReferenceTypes[codespeedreferencetype]
	return ok
}

func (codestandbypowertype *CodeStandbyPowerType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStandbyPowerTypes[codestandbypowertype]
	return ok
}

func (codestatusairporttype *CodeStatusAirportType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusAirportTypes[codestatusairporttype]
	return ok
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusAirspaceTypes[codestatusairspacetype]
	return ok
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusConstructionTypes[codestatusconstructiontype]
	return ok
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusNavaidTypes[codestatusnavaidtype]
	return ok
}

func (codestatusoperationstype *CodeStatusOperationsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusOperationsTypes[codestatusoperationstype]
	return ok
}

func (codestatusservicetype *CodeStatusServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusServiceTypes[codestatusservicetype]
	return ok
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSurfaceCompositionTypes[codesurfacecompositiontype]
	return ok
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSurfaceConditionTypes[codesurfaceconditiontype]
	return ok
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSurfacePreparationTypes[codesurfacepreparationtype]
	return ok
}

func (codesystemactivationtype *CodeSystemActivationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSystemActivationTypes[codesystemactivationtype]
	return ok
}

func (codetaatype *CodeTAAType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTAATypes[codetaatype]
	return ok
}

func (codetacanchanneltype *CodeTACANChannelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTACANChannelTypes[codetacanchanneltype]
	return ok
}

func (codetlofsectiontype *CodeTLOFSectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTLOFSectionTypes[codetlofsectiontype]
	return ok
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTaxiwayElementTypes[codetaxiwayelementtype]
	return ok
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTaxiwaySectionTypes[codetaxiwaysectiontype]
	return ok
}

func (codetaxiwaytype *CodeTaxiwayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTaxiwayTypes[codetaxiwaytype]
	return ok
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTelecomNetworkTypes[codetelecomnetworktype]
	return ok
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTimeEventCombinationTypes[codetimeeventcombinationtype]
	return ok
}

func (codetimeeventtype *CodeTimeEventType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTimeEventTypes[codetimeeventtype]
	return ok
}

func (codetimereferencetype *CodeTimeReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTimeReferenceTypes[codetimereferencetype]
	return ok
}

func (codetrajectorytype *CodeTrajectoryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTrajectoryTypes[codetrajectorytype]
	return ok
}

func (codetranspondertype *CodeTransponderType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTransponderTypes[codetranspondertype]
	return ok
}

func (codeuuidtype *CodeUUIDType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeUUIDTypes[codeuuidtype]
	return ok
}

func (codeunitdependencytype *CodeUnitDependencyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeUnitDependencyTypes[codeunitdependencytype]
	return ok
}

func (codeunittype *CodeUnitType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeUnitTypes[codeunittype]
	return ok
}

func (codeupperalphatype *CodeUpperAlphaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeUpperAlphaTypes[codeupperalphatype]
	return ok
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeUsageLimitationTypes[codeusagelimitationtype]
	return ok
}

func (codevasistype *CodeVASISType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVASISTypes[codevasistype]
	return ok
}

func (codevortype *CodeVORType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVORTypes[codevortype]
	return ok
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeValueInterpretationTypes[codevalueinterpretationtype]
	return ok
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalDatumTypes[codeverticaldatumtype]
	return ok
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalReferenceTypes[codeverticalreferencetype]
	return ok
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalStructureMarkingTypes[codeverticalstructuremarkingtype]
	return ok
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalStructureMaterialTypes[codeverticalstructurematerialtype]
	return ok
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalStructureTypes[codeverticalstructuretype]
	return ok
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVisualDockingGuidanceTypes[codevisualdockingguidancetype]
	return ok
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeWakeTurbulenceTypes[codewaketurbulencetype]
	return ok
}

func (codeworkareatype *CodeWorkAreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeWorkAreaTypes[codeworkareatype]
	return ok
}

func (codeyesnotype *CodeYesNoType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeYesNoTypes[codeyesnotype]
	return ok
}

func (datemonthdaytype *DateMonthDayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DateMonthDayTypes[datemonthdaytype]
	return ok
}

func (datetimetype *DateTimeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DateTimeTypes[datetimetype]
	return ok
}

func (datetype *DateType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DateTypes[datetype]
	return ok
}

func (dateyeartype *DateYearType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DateYearTypes[dateyeartype]
	return ok
}

func (nonumbertype *NoNumberType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoNumberTypes[nonumbertype]
	return ok
}

func (nosequencetype *NoSequenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoSequenceTypes[nosequencetype]
	return ok
}

func (textaddresstype *TextAddressType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextAddressTypes[textaddresstype]
	return ok
}

func (textdesignatortype *TextDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextDesignatorTypes[textdesignatortype]
	return ok
}

func (textinstructiontype *TextInstructionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextInstructionTypes[textinstructiontype]
	return ok
}

func (textnametype *TextNameType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextNameTypes[textnametype]
	return ok
}

func (textnotetype *TextNoteType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextNoteTypes[textnotetype]
	return ok
}

func (textphonetype *TextPhoneType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextPhoneTypes[textphonetype]
	return ok
}

func (textpropertynametype *TextPropertyNameType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextPropertyNameTypes[textpropertynametype]
	return ok
}

func (textremarktype *TextRemarkType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextRemarkTypes[textremarktype]
	return ok
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextSIDSTARDesignatorTypes[textsidstardesignatortype]
	return ok
}

func (timetype *TimeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimeTypes[timetype]
	return ok
}

func (valalarmlimittype *ValAlarmLimitType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValAlarmLimitTypes[valalarmlimittype]
	return ok
}

func (valangletype *ValAngleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValAngleTypes[valangletype]
	return ok
}

func (valbearingtype *ValBearingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValBearingTypes[valbearingtype]
	return ok
}

func (valchannelnumbertype *ValChannelNumberType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValChannelNumberTypes[valchannelnumbertype]
	return ok
}

func (valdepthtype *ValDepthType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValDepthTypes[valdepthtype]
	return ok
}

func (valdistancesignedtype *ValDistanceSignedType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValDistanceSignedTypes[valdistancesignedtype]
	return ok
}

func (valdistancetype *ValDistanceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValDistanceTypes[valdistancetype]
	return ok
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValDistanceVerticalTypes[valdistanceverticaltype]
	return ok
}

func (valdurationtype *ValDurationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValDurationTypes[valdurationtype]
	return ok
}

func (valfltype *ValFLType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValFLTypes[valfltype]
	return ok
}

func (valfrequencytype *ValFrequencyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValFrequencyTypes[valfrequencytype]
	return ok
}

func (valfrictiontype *ValFrictionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValFrictionTypes[valfrictiontype]
	return ok
}

func (valhextype *ValHexType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValHexTypes[valhextype]
	return ok
}

func (vallcntype *ValLCNType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValLCNTypes[vallcntype]
	return ok
}

func (vallightintensitytype *ValLightIntensityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValLightIntensityTypes[vallightintensitytype]
	return ok
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValMagneticVariationChangeTypes[valmagneticvariationchangetype]
	return ok
}

func (valmagneticvariationtype *ValMagneticVariationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValMagneticVariationTypes[valmagneticvariationtype]
	return ok
}

func (valpcntype *ValPCNType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValPCNTypes[valpcntype]
	return ok
}

func (valpercenttype *ValPercentType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValPercentTypes[valpercenttype]
	return ok
}

func (valpressuretype *ValPressureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValPressureTypes[valpressuretype]
	return ok
}

func (valslopetype *ValSlopeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValSlopeTypes[valslopetype]
	return ok
}

func (valspeedtype *ValSpeedType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValSpeedTypes[valspeedtype]
	return ok
}

func (valtemperaturetype *ValTemperatureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValTemperatureTypes[valtemperaturetype]
	return ok
}

func (valweighttype *ValWeightType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValWeightTypes[valweighttype]
	return ok
}

func (xhtmltype *XHTMLType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.XHTMLTypes[xhtmltype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (codeatcreportingtype *CodeATCReportingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeatcreportingtype) {
		return
	}

	codeatcreportingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaerialrefuellingpointtype) {
		return
	}

	codeaerialrefuellingpointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaerialrefuellingprefixtype) {
		return
	}

	codeaerialrefuellingprefixtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaerialrefuellingtype) {
		return
	}

	codeaerialrefuellingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircraftcategorytype) {
		return
	}

	codeaircraftcategorytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircraftenginenumbertype) {
		return
	}

	codeaircraftenginenumbertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftenginetype *CodeAircraftEngineType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircraftenginetype) {
		return
	}

	codeaircraftenginetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircraftgroundservicetype) {
		return
	}

	codeaircraftgroundservicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircrafticaotype *CodeAircraftICAOType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircrafticaotype) {
		return
	}

	codeaircrafticaotype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftstandtype *CodeAircraftStandType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircraftstandtype) {
		return
	}

	codeaircraftstandtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircrafttype *CodeAircraftType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircrafttype) {
		return
	}

	codeaircrafttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircraftwingspanclasstype) {
		return
	}

	codeaircraftwingspanclasstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairportheliportcollocationtype) {
		return
	}

	codeairportheliportcollocationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairportheliportdesignatortype) {
		return
	}

	codeairportheliportdesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairportheliporttype *CodeAirportHeliportType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairportheliporttype) {
		return
	}

	codeairportheliporttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairportwarningtype *CodeAirportWarningType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairportwarningtype) {
		return
	}

	codeairportwarningtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspaceactivitytype) {
		return
	}

	codeairspaceactivitytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspaceaggregationtype) {
		return
	}

	codeairspaceaggregationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspaceclassificationtype) {
		return
	}

	codeairspaceclassificationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacedependencytype) {
		return
	}

	codeairspacedependencytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacedesignatortype) {
		return
	}

	codeairspacedesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacepointpositiontype) {
		return
	}

	codeairspacepointpositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacepointroletype) {
		return
	}

	codeairspacepointroletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacetype *CodeAirspaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacetype) {
		return
	}

	codeairspacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codealtitudeadjustmenttype) {
		return
	}

	codealtitudeadjustmenttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codealtitudeusetype *CodeAltitudeUseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codealtitudeusetype) {
		return
	}

	codealtitudeusetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapproachequipmentadditionaltype) {
		return
	}

	codeapproachequipmentadditionaltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapproachguidancetype) {
		return
	}

	codeapproachguidancetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapproachlightingicaotype) {
		return
	}

	codeapproachlightingicaotype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachlightingtype *CodeApproachLightingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapproachlightingtype) {
		return
	}

	codeapproachlightingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachprefixtype *CodeApproachPrefixType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapproachprefixtype) {
		return
	}

	codeapproachprefixtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachtype *CodeApproachType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapproachtype) {
		return
	}

	codeapproachtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapprovaltype *CodeApprovalType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapprovaltype) {
		return
	}

	codeapprovaltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapronelementtype *CodeApronElementType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapronelementtype) {
		return
	}

	codeapronelementtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapronsectiontype *CodeApronSectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeapronsectiontype) {
		return
	}

	codeapronsectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codearcdirectiontype *CodeArcDirectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codearcdirectiontype) {
		return
	}

	codearcdirectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codearrestinggearenergyabsorbtype) {
		return
	}

	codearrestinggearenergyabsorbtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codearrestinggearengagedevicetype) {
		return
	}

	codearrestinggearengagedevicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeauralmorsetype *CodeAuralMorseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeauralmorsetype) {
		return
	}

	codeauralmorsetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeauthorityroletype *CodeAuthorityRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeauthorityroletype) {
		return
	}

	codeauthorityroletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeauthoritytype *CodeAuthorityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeauthoritytype) {
		return
	}

	codeauthoritytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaviationstandardstype *CodeAviationStandardsType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaviationstandardstype) {
		return
	}

	codeaviationstandardstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codebearingtype *CodeBearingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codebearingtype) {
		return
	}

	codebearingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codebuoydesignatortype) {
		return
	}

	codebuoydesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codebuoytype *CodeBuoyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codebuoytype) {
		return
	}

	codebuoytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecardinaldirectiontype) {
		return
	}

	codecardinaldirectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecheckpointcategorytype) {
		return
	}

	codecheckpointcategorytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecolourtype *CodeColourType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecolourtype) {
		return
	}

	codecolourtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecommunicationchanneltype) {
		return
	}

	codecommunicationchanneltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecommunicationdirectiontype) {
		return
	}

	codecommunicationdirectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecommunicationmodetype *CodeCommunicationModeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecommunicationmodetype) {
		return
	}

	codecommunicationmodetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecomparisontype *CodeComparisonType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecomparisontype) {
		return
	}

	codecomparisontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecontaminationtype *CodeContaminationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecontaminationtype) {
		return
	}

	codecontaminationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecoursequalityilstype) {
		return
	}

	codecoursequalityilstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecoursetype *CodeCourseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecoursetype) {
		return
	}

	codecoursetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedmechanneltype *CodeDMEChannelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedmechanneltype) {
		return
	}

	codedmechanneltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedmetype *CodeDMEType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedmetype) {
		return
	}

	codedmetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedaytype *CodeDayType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedaytype) {
		return
	}

	codedaytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedeclareddistancetype) {
		return
	}

	codedeclareddistancetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedesignstandardtype *CodeDesignStandardType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedesignstandardtype) {
		return
	}

	codedesignstandardtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedesignatedpointdesignatortype) {
		return
	}

	codedesignatedpointdesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedesignatedpointtype *CodeDesignatedPointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedesignatedpointtype) {
		return
	}

	codedesignatedpointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedirectionreferencetype) {
		return
	}

	codedirectionreferencetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedirectionturntype *CodeDirectionTurnType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedirectionturntype) {
		return
	}

	codedirectionturntype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedirectiontype *CodeDirectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedirectiontype) {
		return
	}

	codedirectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedistanceindicationtype) {
		return
	}

	codedistanceindicationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeemissionbandtype *CodeEmissionBandType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeemissionbandtype) {
		return
	}

	codeemissionbandtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeequipmentanticollisiontype) {
		return
	}

	codeequipmentanticollisiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeequipmentunavailabletype) {
		return
	}

	codeequipmentunavailabletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefacilityrankingtype *CodeFacilityRankingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codefacilityrankingtype) {
		return
	}

	codefacilityrankingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefinalguidancetype *CodeFinalGuidanceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codefinalguidancetype) {
		return
	}

	codefinalguidancetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefirefightingtype *CodeFireFightingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codefirefightingtype) {
		return
	}

	codefirefightingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightdestinationtype *CodeFlightDestinationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflightdestinationtype) {
		return
	}

	codeflightdestinationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightorigintype *CodeFlightOriginType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflightorigintype) {
		return
	}

	codeflightorigintype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightpurposetype *CodeFlightPurposeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflightpurposetype) {
		return
	}

	codeflightpurposetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflightrestrictiondesignatortype) {
		return
	}

	codeflightrestrictiondesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflightrestrictiontype) {
		return
	}

	codeflightrestrictiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightruletype *CodeFlightRuleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflightruletype) {
		return
	}

	codeflightruletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightstatustype *CodeFlightStatusType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflightstatustype) {
		return
	}

	codeflightstatustype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflighttype *CodeFlightType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflighttype) {
		return
	}

	codeflighttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflowconditionoperationtype) {
		return
	}

	codeflowconditionoperationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefreeflighttype *CodeFreeFlightType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codefreeflighttype) {
		return
	}

	codefreeflighttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codefrictiondevicetype) {
		return
	}

	codefrictiondevicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codefrictionestimatetype) {
		return
	}

	codefrictionestimatetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefueltype *CodeFuelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codefueltype) {
		return
	}

	codefueltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codegeobordertype *CodeGeoBorderType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codegeobordertype) {
		return
	}

	codegeobordertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codegradeseparationtype *CodeGradeSeparationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codegradeseparationtype) {
		return
	}

	codegradeseparationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codegroundlightingtype *CodeGroundLightingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codegroundlightingtype) {
		return
	}

	codegroundlightingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeguidancelinetype *CodeGuidanceLineType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeguidancelinetype) {
		return
	}

	codeguidancelinetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeheightreferencetype *CodeHeightReferenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeheightreferencetype) {
		return
	}

	codeheightreferencetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codehelicopterperformancetype) {
		return
	}

	codehelicopterperformancetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeholdingcategorytype) {
		return
	}

	codeholdingcategorytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeholdingusagetype *CodeHoldingUsageType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeholdingusagetype) {
		return
	}

	codeholdingusagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeholdingusetype *CodeHoldingUseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeholdingusetype) {
		return
	}

	codeholdingusetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeiatatype *CodeIATAType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeiatatype) {
		return
	}

	codeiatatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeicaocountrytype *CodeICAOCountryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeicaocountrytype) {
		return
	}

	codeicaocountrytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeicaotype *CodeICAOType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeicaotype) {
		return
	}

	codeicaotype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeilsbackcoursetype) {
		return
	}

	codeilsbackcoursetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeintegritylevelilstype) {
		return
	}

	codeintegritylevelilstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeintensitystandbytype *CodeIntensityStandByType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeintensitystandbytype) {
		return
	}

	codeintensitystandbytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelanguagetype *CodeLanguageType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelanguagetype) {
		return
	}

	codelanguagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelevelseriestype *CodeLevelSeriesType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelevelseriestype) {
		return
	}

	codelevelseriestype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeleveltabledesignatortype) {
		return
	}

	codeleveltabledesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeleveltype *CodeLevelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeleveltype) {
		return
	}

	codeleveltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelightholdingpositiontype) {
		return
	}

	codelightholdingpositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelightintensitytype *CodeLightIntensityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelightintensitytype) {
		return
	}

	codelightintensitytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelightsourcetype *CodeLightSourceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelightsourcetype) {
		return
	}

	codelightsourcetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelightingjartype *CodeLightingJARType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelightingjartype) {
		return
	}

	codelightingjartype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeloadingbridgetype) {
		return
	}

	codeloadingbridgetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelocationqualifiertype *CodeLocationQualifierType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelocationqualifiertype) {
		return
	}

	codelocationqualifiertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelogicaloperatortype) {
		return
	}

	codelogicaloperatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemlsazimuthtype) {
		return
	}

	codemlsazimuthtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemlschanneltype *CodeMLSChannelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemlschanneltype) {
		return
	}

	codemlschanneltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemarkerbeaconsignaltype) {
		return
	}

	codemarkerbeaconsignaltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemarkingconditiontype *CodeMarkingConditionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemarkingconditiontype) {
		return
	}

	codemarkingconditiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemarkingstyletype *CodeMarkingStyleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemarkingstyletype) {
		return
	}

	codemarkingstyletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemeteoconditionstype) {
		return
	}

	codemeteoconditionstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemilitaryoperationstype) {
		return
	}

	codemilitaryoperationstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemilitaryroutepointtype) {
		return
	}

	codemilitaryroutepointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemilitarystatustype *CodeMilitaryStatusType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemilitarystatustype) {
		return
	}

	codemilitarystatustype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemilitarytrainingtype) {
		return
	}

	codemilitarytrainingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeminimafinalapproachpathtype) {
		return
	}

	codeminimafinalapproachpathtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeminimumaltitudetype) {
		return
	}

	codeminimumaltitudetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeminimumheighttype *CodeMinimumHeightType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeminimumheighttype) {
		return
	}

	codeminimumheighttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemissedapproachtype *CodeMissedApproachType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemissedapproachtype) {
		return
	}

	codemissedapproachtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codendbusagetype *CodeNDBUsageType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codendbusagetype) {
		return
	}

	codendbusagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavaiddesignatortype) {
		return
	}

	codenavaiddesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavaidpurposetype) {
		return
	}

	codenavaidpurposetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaidservicetype *CodeNavaidServiceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavaidservicetype) {
		return
	}

	codenavaidservicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavigationarearestrictiontype) {
		return
	}

	codenavigationarearestrictiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavigationareatype *CodeNavigationAreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavigationareatype) {
		return
	}

	codenavigationareatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavigationequipmenttype) {
		return
	}

	codenavigationequipmenttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavigationspecificationtype) {
		return
	}

	codenavigationspecificationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenitrogentype *CodeNitrogenType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenitrogentype) {
		return
	}

	codenitrogentype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenorthreferencetype *CodeNorthReferenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenorthreferencetype) {
		return
	}

	codenorthreferencetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenotepurposetype *CodeNotePurposeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenotepurposetype) {
		return
	}

	codenotepurposetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeobstacleareatype *CodeObstacleAreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeobstacleareatype) {
		return
	}

	codeobstacleareatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeobstacleassessmentsurfacetype) {
		return
	}

	codeobstacleassessmentsurfacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeobstructionidsurfacezonetype) {
		return
	}

	codeobstructionidsurfacezonetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeoiltype *CodeOilType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeoiltype) {
		return
	}

	codeoiltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeoperationairportheliporttype) {
		return
	}

	codeoperationairportheliporttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeoperationmanoeuvringareatype) {
		return
	}

	codeoperationmanoeuvringareatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeorganisationdesignatortype) {
		return
	}

	codeorganisationdesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeorganisationhierarchytype) {
		return
	}

	codeorganisationhierarchytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeorganisationtype *CodeOrganisationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeorganisationtype) {
		return
	}

	codeorganisationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeoxygentype *CodeOxygenType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeoxygentype) {
		return
	}

	codeoxygentype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepartype *CodePARType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepartype) {
		return
	}

	codepartype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepcnmethodtype *CodePCNMethodType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepcnmethodtype) {
		return
	}

	codepcnmethodtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepcnpavementtype *CodePCNPavementType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepcnpavementtype) {
		return
	}

	codepcnpavementtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepcnsubgradetype *CodePCNSubgradeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepcnsubgradetype) {
		return
	}

	codepcnsubgradetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepcntyrepressuretype) {
		return
	}

	codepcntyrepressuretype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepassengerservicetype *CodePassengerServiceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepassengerservicetype) {
		return
	}

	codepassengerservicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepilotcontrolledlightingtype) {
		return
	}

	codepilotcontrolledlightingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepositioninilstype *CodePositionInILSType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepositioninilstype) {
		return
	}

	codepositioninilstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprimaryradartype *CodePrimaryRadarType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeprimaryradartype) {
		return
	}

	codeprimaryradartype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeprocedureavailabilitytype) {
		return
	}

	codeprocedureavailabilitytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeprocedurecodingstandardtype) {
		return
	}

	codeprocedurecodingstandardtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeproceduredistancetype) {
		return
	}

	codeproceduredistancetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeprocedurefixroletype) {
		return
	}

	codeprocedurefixroletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeprocedurephasetype) {
		return
	}

	codeprocedurephasetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeprotectareasectiontype) {
		return
	}

	codeprotectareasectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codernptype *CodeRNPType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codernptype) {
		return
	}

	codernptype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codervrreadingtype *CodeRVRReadingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codervrreadingtype) {
		return
	}

	codervrreadingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codervsmpointroletype) {
		return
	}

	codervsmpointroletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codervsmtype *CodeRVSMType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codervsmtype) {
		return
	}

	codervsmtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadarservicetype *CodeRadarServiceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderadarservicetype) {
		return
	}

	coderadarservicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadioemissiontype *CodeRadioEmissionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderadioemissiontype) {
		return
	}

	coderadioemissiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderadiofrequencyareatype) {
		return
	}

	coderadiofrequencyareatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadiosignaltype *CodeRadioSignalType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderadiosignaltype) {
		return
	}

	coderadiosignaltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codereferencepathidentifiertype) {
		return
	}

	codereferencepathidentifiertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codereferenceroletype *CodeReferenceRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codereferenceroletype) {
		return
	}

	codereferenceroletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codereflectortype *CodeReflectorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codereflectortype) {
		return
	}

	codereflectortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderelativepositiontype *CodeRelativePositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderelativepositiontype) {
		return
	}

	coderelativepositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoadtype *CodeRoadType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderoadtype) {
		return
	}

	coderoadtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderouteavailabilitytype) {
		return
	}

	coderouteavailabilitytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderoutedesignatorlettertype) {
		return
	}

	coderoutedesignatorlettertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderoutedesignatorprefixtype) {
		return
	}

	coderoutedesignatorprefixtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderoutedesignatorsuffixtype) {
		return
	}

	coderoutedesignatorsuffixtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderouteindicatortype *CodeRouteIndicatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderouteindicatortype) {
		return
	}

	coderouteindicatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutenavigationtype *CodeRouteNavigationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderoutenavigationtype) {
		return
	}

	coderoutenavigationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderouteorigintype *CodeRouteOriginType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderouteorigintype) {
		return
	}

	coderouteorigintype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderoutesegmentpathtype) {
		return
	}

	coderoutesegmentpathtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutetype *CodeRouteType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderoutetype) {
		return
	}

	coderoutetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderuleproceduretitletype) {
		return
	}

	coderuleproceduretitletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderuleproceduretype *CodeRuleProcedureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderuleproceduretype) {
		return
	}

	coderuleproceduretype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwayelementtype *CodeRunwayElementType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderunwayelementtype) {
		return
	}

	coderunwayelementtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderunwaymarkingtype) {
		return
	}

	coderunwaymarkingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderunwaypointroletype) {
		return
	}

	coderunwaypointroletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderunwayprotectionareatype) {
		return
	}

	coderunwayprotectionareatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwaysectiontype *CodeRunwaySectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderunwaysectiontype) {
		return
	}

	coderunwaysectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwaytype *CodeRunwayType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderunwaytype) {
		return
	}

	coderunwaytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesafealtitudetype *CodeSafeAltitudeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesafealtitudetype) {
		return
	}

	codesafealtitudetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesegmentpathtype *CodeSegmentPathType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesegmentpathtype) {
		return
	}

	codesegmentpathtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesegmentterminationtype) {
		return
	}

	codesegmentterminationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeserviceatctype *CodeServiceATCType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeserviceatctype) {
		return
	}

	codeserviceatctype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeserviceatfmtype *CodeServiceATFMType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeserviceatfmtype) {
		return
	}

	codeserviceatfmtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeservicegroundcontroltype) {
		return
	}

	codeservicegroundcontroltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeserviceinformationtype *CodeServiceInformationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeserviceinformationtype) {
		return
	}

	codeserviceinformationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeservicesartype *CodeServiceSARType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeservicesartype) {
		return
	}

	codeservicesartype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesidetype *CodeSideType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesidetype) {
		return
	}

	codesidetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesignalperformanceilstype) {
		return
	}

	codesignalperformanceilstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialdatetype *CodeSpecialDateType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespecialdatetype) {
		return
	}

	codespecialdatetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespecialnavigationchaindesignatortype) {
		return
	}

	codespecialnavigationchaindesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespecialnavigationstationtype) {
		return
	}

	codespecialnavigationstationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespecialnavigationsystemtype) {
		return
	}

	codespecialnavigationsystemtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespeedreferencetype *CodeSpeedReferenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespeedreferencetype) {
		return
	}

	codespeedreferencetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestandbypowertype *CodeStandbyPowerType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestandbypowertype) {
		return
	}

	codestandbypowertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusairporttype *CodeStatusAirportType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusairporttype) {
		return
	}

	codestatusairporttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusairspacetype *CodeStatusAirspaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusairspacetype) {
		return
	}

	codestatusairspacetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusconstructiontype *CodeStatusConstructionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusconstructiontype) {
		return
	}

	codestatusconstructiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusnavaidtype *CodeStatusNavaidType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusnavaidtype) {
		return
	}

	codestatusnavaidtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusoperationstype *CodeStatusOperationsType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusoperationstype) {
		return
	}

	codestatusoperationstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusservicetype *CodeStatusServiceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusservicetype) {
		return
	}

	codestatusservicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesurfacecompositiontype) {
		return
	}

	codesurfacecompositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesurfaceconditiontype) {
		return
	}

	codesurfaceconditiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesurfacepreparationtype) {
		return
	}

	codesurfacepreparationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesystemactivationtype *CodeSystemActivationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesystemactivationtype) {
		return
	}

	codesystemactivationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetaatype *CodeTAAType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetaatype) {
		return
	}

	codetaatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetacanchanneltype *CodeTACANChannelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetacanchanneltype) {
		return
	}

	codetacanchanneltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetlofsectiontype *CodeTLOFSectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetlofsectiontype) {
		return
	}

	codetlofsectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetaxiwayelementtype) {
		return
	}

	codetaxiwayelementtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetaxiwaysectiontype) {
		return
	}

	codetaxiwaysectiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetaxiwaytype *CodeTaxiwayType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetaxiwaytype) {
		return
	}

	codetaxiwaytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetelecomnetworktype) {
		return
	}

	codetelecomnetworktype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetimeeventcombinationtype) {
		return
	}

	codetimeeventcombinationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetimeeventtype *CodeTimeEventType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetimeeventtype) {
		return
	}

	codetimeeventtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetimereferencetype *CodeTimeReferenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetimereferencetype) {
		return
	}

	codetimereferencetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetrajectorytype *CodeTrajectoryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetrajectorytype) {
		return
	}

	codetrajectorytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetranspondertype *CodeTransponderType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetranspondertype) {
		return
	}

	codetranspondertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeuuidtype *CodeUUIDType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeuuidtype) {
		return
	}

	codeuuidtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeunitdependencytype *CodeUnitDependencyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeunitdependencytype) {
		return
	}

	codeunitdependencytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeunittype *CodeUnitType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeunittype) {
		return
	}

	codeunittype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeupperalphatype *CodeUpperAlphaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeupperalphatype) {
		return
	}

	codeupperalphatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeusagelimitationtype *CodeUsageLimitationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeusagelimitationtype) {
		return
	}

	codeusagelimitationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codevasistype *CodeVASISType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codevasistype) {
		return
	}

	codevasistype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codevortype *CodeVORType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codevortype) {
		return
	}

	codevortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codevalueinterpretationtype) {
		return
	}

	codevalueinterpretationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticaldatumtype *CodeVerticalDatumType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticaldatumtype) {
		return
	}

	codeverticaldatumtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticalreferencetype) {
		return
	}

	codeverticalreferencetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticalstructuremarkingtype) {
		return
	}

	codeverticalstructuremarkingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticalstructurematerialtype) {
		return
	}

	codeverticalstructurematerialtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructuretype *CodeVerticalStructureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticalstructuretype) {
		return
	}

	codeverticalstructuretype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codevisualdockingguidancetype) {
		return
	}

	codevisualdockingguidancetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codewaketurbulencetype) {
		return
	}

	codewaketurbulencetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeworkareatype *CodeWorkAreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeworkareatype) {
		return
	}

	codeworkareatype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeyesnotype *CodeYesNoType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeyesnotype) {
		return
	}

	codeyesnotype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datemonthdaytype *DateMonthDayType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(datemonthdaytype) {
		return
	}

	datemonthdaytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datetimetype *DateTimeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(datetimetype) {
		return
	}

	datetimetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datetype *DateType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(datetype) {
		return
	}

	datetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dateyeartype *DateYearType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(dateyeartype) {
		return
	}

	dateyeartype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (nonumbertype *NoNumberType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(nonumbertype) {
		return
	}

	nonumbertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (nosequencetype *NoSequenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(nosequencetype) {
		return
	}

	nosequencetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textaddresstype *TextAddressType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textaddresstype) {
		return
	}

	textaddresstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textdesignatortype *TextDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textdesignatortype) {
		return
	}

	textdesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textinstructiontype *TextInstructionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textinstructiontype) {
		return
	}

	textinstructiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textnametype *TextNameType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textnametype) {
		return
	}

	textnametype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textnotetype *TextNoteType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textnotetype) {
		return
	}

	textnotetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textphonetype *TextPhoneType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textphonetype) {
		return
	}

	textphonetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textpropertynametype *TextPropertyNameType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textpropertynametype) {
		return
	}

	textpropertynametype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textremarktype *TextRemarkType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textremarktype) {
		return
	}

	textremarktype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textsidstardesignatortype) {
		return
	}

	textsidstardesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timetype *TimeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timetype) {
		return
	}

	timetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valalarmlimittype *ValAlarmLimitType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valalarmlimittype) {
		return
	}

	valalarmlimittype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valangletype *ValAngleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valangletype) {
		return
	}

	valangletype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valbearingtype *ValBearingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valbearingtype) {
		return
	}

	valbearingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valchannelnumbertype *ValChannelNumberType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valchannelnumbertype) {
		return
	}

	valchannelnumbertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdepthtype *ValDepthType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valdepthtype) {
		return
	}

	valdepthtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdistancesignedtype *ValDistanceSignedType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valdistancesignedtype) {
		return
	}

	valdistancesignedtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdistancetype *ValDistanceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valdistancetype) {
		return
	}

	valdistancetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdistanceverticaltype *ValDistanceVerticalType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valdistanceverticaltype) {
		return
	}

	valdistanceverticaltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdurationtype *ValDurationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valdurationtype) {
		return
	}

	valdurationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valfltype *ValFLType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valfltype) {
		return
	}

	valfltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valfrequencytype *ValFrequencyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valfrequencytype) {
		return
	}

	valfrequencytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valfrictiontype *ValFrictionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valfrictiontype) {
		return
	}

	valfrictiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valhextype *ValHexType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valhextype) {
		return
	}

	valhextype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vallcntype *ValLCNType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(vallcntype) {
		return
	}

	vallcntype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vallightintensitytype *ValLightIntensityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(vallightintensitytype) {
		return
	}

	vallightintensitytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valmagneticvariationchangetype) {
		return
	}

	valmagneticvariationchangetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valmagneticvariationtype *ValMagneticVariationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valmagneticvariationtype) {
		return
	}

	valmagneticvariationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valpcntype *ValPCNType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valpcntype) {
		return
	}

	valpcntype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valpercenttype *ValPercentType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valpercenttype) {
		return
	}

	valpercenttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valpressuretype *ValPressureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valpressuretype) {
		return
	}

	valpressuretype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valslopetype *ValSlopeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valslopetype) {
		return
	}

	valslopetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valspeedtype *ValSpeedType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valspeedtype) {
		return
	}

	valspeedtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valtemperaturetype *ValTemperatureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valtemperaturetype) {
		return
	}

	valtemperaturetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valweighttype *ValWeightType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valweighttype) {
		return
	}

	valweighttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xhtmltype *XHTMLType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(xhtmltype) {
		return
	}

	xhtmltype.Stage(stage)

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
	case *CodeATCReportingType:
		toT := GongCopyBranchCodeATCReportingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAerialRefuellingPointType:
		toT := GongCopyBranchCodeAerialRefuellingPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAerialRefuellingPrefixType:
		toT := GongCopyBranchCodeAerialRefuellingPrefixType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAerialRefuellingType:
		toT := GongCopyBranchCodeAerialRefuellingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftCategoryType:
		toT := GongCopyBranchCodeAircraftCategoryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftEngineNumberType:
		toT := GongCopyBranchCodeAircraftEngineNumberType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftEngineType:
		toT := GongCopyBranchCodeAircraftEngineType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftGroundServiceType:
		toT := GongCopyBranchCodeAircraftGroundServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftICAOType:
		toT := GongCopyBranchCodeAircraftICAOType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftStandType:
		toT := GongCopyBranchCodeAircraftStandType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftType:
		toT := GongCopyBranchCodeAircraftType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftWingspanClassType:
		toT := GongCopyBranchCodeAircraftWingspanClassType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirportHeliportCollocationType:
		toT := GongCopyBranchCodeAirportHeliportCollocationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirportHeliportDesignatorType:
		toT := GongCopyBranchCodeAirportHeliportDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirportHeliportType:
		toT := GongCopyBranchCodeAirportHeliportType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirportWarningType:
		toT := GongCopyBranchCodeAirportWarningType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceActivityType:
		toT := GongCopyBranchCodeAirspaceActivityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceAggregationType:
		toT := GongCopyBranchCodeAirspaceAggregationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceClassificationType:
		toT := GongCopyBranchCodeAirspaceClassificationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceDependencyType:
		toT := GongCopyBranchCodeAirspaceDependencyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceDesignatorType:
		toT := GongCopyBranchCodeAirspaceDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspacePointPositionType:
		toT := GongCopyBranchCodeAirspacePointPositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspacePointRoleType:
		toT := GongCopyBranchCodeAirspacePointRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceType:
		toT := GongCopyBranchCodeAirspaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAltitudeAdjustmentType:
		toT := GongCopyBranchCodeAltitudeAdjustmentType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAltitudeUseType:
		toT := GongCopyBranchCodeAltitudeUseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachEquipmentAdditionalType:
		toT := GongCopyBranchCodeApproachEquipmentAdditionalType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachGuidanceType:
		toT := GongCopyBranchCodeApproachGuidanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachLightingICAOType:
		toT := GongCopyBranchCodeApproachLightingICAOType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachLightingType:
		toT := GongCopyBranchCodeApproachLightingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachPrefixType:
		toT := GongCopyBranchCodeApproachPrefixType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachType:
		toT := GongCopyBranchCodeApproachType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApprovalType:
		toT := GongCopyBranchCodeApprovalType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApronElementType:
		toT := GongCopyBranchCodeApronElementType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApronSectionType:
		toT := GongCopyBranchCodeApronSectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeArcDirectionType:
		toT := GongCopyBranchCodeArcDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeArrestingGearEnergyAbsorbType:
		toT := GongCopyBranchCodeArrestingGearEnergyAbsorbType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeArrestingGearEngageDeviceType:
		toT := GongCopyBranchCodeArrestingGearEngageDeviceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAuralMorseType:
		toT := GongCopyBranchCodeAuralMorseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAuthorityRoleType:
		toT := GongCopyBranchCodeAuthorityRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAuthorityType:
		toT := GongCopyBranchCodeAuthorityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAviationStandardsType:
		toT := GongCopyBranchCodeAviationStandardsType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeBearingType:
		toT := GongCopyBranchCodeBearingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeBuoyDesignatorType:
		toT := GongCopyBranchCodeBuoyDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeBuoyType:
		toT := GongCopyBranchCodeBuoyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCardinalDirectionType:
		toT := GongCopyBranchCodeCardinalDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCheckpointCategoryType:
		toT := GongCopyBranchCodeCheckpointCategoryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeColourType:
		toT := GongCopyBranchCodeColourType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCommunicationChannelType:
		toT := GongCopyBranchCodeCommunicationChannelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCommunicationDirectionType:
		toT := GongCopyBranchCodeCommunicationDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCommunicationModeType:
		toT := GongCopyBranchCodeCommunicationModeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeComparisonType:
		toT := GongCopyBranchCodeComparisonType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeContaminationType:
		toT := GongCopyBranchCodeContaminationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCourseQualityILSType:
		toT := GongCopyBranchCodeCourseQualityILSType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCourseType:
		toT := GongCopyBranchCodeCourseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDMEChannelType:
		toT := GongCopyBranchCodeDMEChannelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDMEType:
		toT := GongCopyBranchCodeDMEType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDayType:
		toT := GongCopyBranchCodeDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDeclaredDistanceType:
		toT := GongCopyBranchCodeDeclaredDistanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDesignStandardType:
		toT := GongCopyBranchCodeDesignStandardType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDesignatedPointDesignatorType:
		toT := GongCopyBranchCodeDesignatedPointDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDesignatedPointType:
		toT := GongCopyBranchCodeDesignatedPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDirectionReferenceType:
		toT := GongCopyBranchCodeDirectionReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDirectionTurnType:
		toT := GongCopyBranchCodeDirectionTurnType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDirectionType:
		toT := GongCopyBranchCodeDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDistanceIndicationType:
		toT := GongCopyBranchCodeDistanceIndicationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeEmissionBandType:
		toT := GongCopyBranchCodeEmissionBandType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeEquipmentAntiCollisionType:
		toT := GongCopyBranchCodeEquipmentAntiCollisionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeEquipmentUnavailableType:
		toT := GongCopyBranchCodeEquipmentUnavailableType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFacilityRankingType:
		toT := GongCopyBranchCodeFacilityRankingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFinalGuidanceType:
		toT := GongCopyBranchCodeFinalGuidanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFireFightingType:
		toT := GongCopyBranchCodeFireFightingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightDestinationType:
		toT := GongCopyBranchCodeFlightDestinationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightOriginType:
		toT := GongCopyBranchCodeFlightOriginType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightPurposeType:
		toT := GongCopyBranchCodeFlightPurposeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightRestrictionDesignatorType:
		toT := GongCopyBranchCodeFlightRestrictionDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightRestrictionType:
		toT := GongCopyBranchCodeFlightRestrictionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightRuleType:
		toT := GongCopyBranchCodeFlightRuleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightStatusType:
		toT := GongCopyBranchCodeFlightStatusType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightType:
		toT := GongCopyBranchCodeFlightType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlowConditionOperationType:
		toT := GongCopyBranchCodeFlowConditionOperationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFreeFlightType:
		toT := GongCopyBranchCodeFreeFlightType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFrictionDeviceType:
		toT := GongCopyBranchCodeFrictionDeviceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFrictionEstimateType:
		toT := GongCopyBranchCodeFrictionEstimateType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFuelType:
		toT := GongCopyBranchCodeFuelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeGeoBorderType:
		toT := GongCopyBranchCodeGeoBorderType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeGradeSeparationType:
		toT := GongCopyBranchCodeGradeSeparationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeGroundLightingType:
		toT := GongCopyBranchCodeGroundLightingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeGuidanceLineType:
		toT := GongCopyBranchCodeGuidanceLineType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHeightReferenceType:
		toT := GongCopyBranchCodeHeightReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHelicopterPerformanceType:
		toT := GongCopyBranchCodeHelicopterPerformanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHoldingCategoryType:
		toT := GongCopyBranchCodeHoldingCategoryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHoldingUsageType:
		toT := GongCopyBranchCodeHoldingUsageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHoldingUseType:
		toT := GongCopyBranchCodeHoldingUseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeIATAType:
		toT := GongCopyBranchCodeIATAType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeICAOCountryType:
		toT := GongCopyBranchCodeICAOCountryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeICAOType:
		toT := GongCopyBranchCodeICAOType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeILSBackCourseType:
		toT := GongCopyBranchCodeILSBackCourseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeIntegrityLevelILSType:
		toT := GongCopyBranchCodeIntegrityLevelILSType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeIntensityStandByType:
		toT := GongCopyBranchCodeIntensityStandByType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLanguageType:
		toT := GongCopyBranchCodeLanguageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLevelSeriesType:
		toT := GongCopyBranchCodeLevelSeriesType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLevelTableDesignatorType:
		toT := GongCopyBranchCodeLevelTableDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLevelType:
		toT := GongCopyBranchCodeLevelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLightHoldingPositionType:
		toT := GongCopyBranchCodeLightHoldingPositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLightIntensityType:
		toT := GongCopyBranchCodeLightIntensityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLightSourceType:
		toT := GongCopyBranchCodeLightSourceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLightingJARType:
		toT := GongCopyBranchCodeLightingJARType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLoadingBridgeType:
		toT := GongCopyBranchCodeLoadingBridgeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLocationQualifierType:
		toT := GongCopyBranchCodeLocationQualifierType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLogicalOperatorType:
		toT := GongCopyBranchCodeLogicalOperatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMLSAzimuthType:
		toT := GongCopyBranchCodeMLSAzimuthType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMLSChannelType:
		toT := GongCopyBranchCodeMLSChannelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMarkerBeaconSignalType:
		toT := GongCopyBranchCodeMarkerBeaconSignalType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMarkingConditionType:
		toT := GongCopyBranchCodeMarkingConditionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMarkingStyleType:
		toT := GongCopyBranchCodeMarkingStyleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMeteoConditionsType:
		toT := GongCopyBranchCodeMeteoConditionsType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMilitaryOperationsType:
		toT := GongCopyBranchCodeMilitaryOperationsType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMilitaryRoutePointType:
		toT := GongCopyBranchCodeMilitaryRoutePointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMilitaryStatusType:
		toT := GongCopyBranchCodeMilitaryStatusType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMilitaryTrainingType:
		toT := GongCopyBranchCodeMilitaryTrainingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMinimaFinalApproachPathType:
		toT := GongCopyBranchCodeMinimaFinalApproachPathType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMinimumAltitudeType:
		toT := GongCopyBranchCodeMinimumAltitudeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMinimumHeightType:
		toT := GongCopyBranchCodeMinimumHeightType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMissedApproachType:
		toT := GongCopyBranchCodeMissedApproachType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNDBUsageType:
		toT := GongCopyBranchCodeNDBUsageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavaidDesignatorType:
		toT := GongCopyBranchCodeNavaidDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavaidPurposeType:
		toT := GongCopyBranchCodeNavaidPurposeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavaidServiceType:
		toT := GongCopyBranchCodeNavaidServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavigationAreaRestrictionType:
		toT := GongCopyBranchCodeNavigationAreaRestrictionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavigationAreaType:
		toT := GongCopyBranchCodeNavigationAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavigationEquipmentType:
		toT := GongCopyBranchCodeNavigationEquipmentType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavigationSpecificationType:
		toT := GongCopyBranchCodeNavigationSpecificationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNitrogenType:
		toT := GongCopyBranchCodeNitrogenType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNorthReferenceType:
		toT := GongCopyBranchCodeNorthReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNotePurposeType:
		toT := GongCopyBranchCodeNotePurposeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeObstacleAreaType:
		toT := GongCopyBranchCodeObstacleAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeObstacleAssessmentSurfaceType:
		toT := GongCopyBranchCodeObstacleAssessmentSurfaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeObstructionIdSurfaceZoneType:
		toT := GongCopyBranchCodeObstructionIdSurfaceZoneType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOilType:
		toT := GongCopyBranchCodeOilType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOperationAirportHeliportType:
		toT := GongCopyBranchCodeOperationAirportHeliportType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOperationManoeuvringAreaType:
		toT := GongCopyBranchCodeOperationManoeuvringAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOrganisationDesignatorType:
		toT := GongCopyBranchCodeOrganisationDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOrganisationHierarchyType:
		toT := GongCopyBranchCodeOrganisationHierarchyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOrganisationType:
		toT := GongCopyBranchCodeOrganisationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOxygenType:
		toT := GongCopyBranchCodeOxygenType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePARType:
		toT := GongCopyBranchCodePARType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePCNMethodType:
		toT := GongCopyBranchCodePCNMethodType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePCNPavementType:
		toT := GongCopyBranchCodePCNPavementType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePCNSubgradeType:
		toT := GongCopyBranchCodePCNSubgradeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePCNTyrePressureType:
		toT := GongCopyBranchCodePCNTyrePressureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePassengerServiceType:
		toT := GongCopyBranchCodePassengerServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePilotControlledLightingType:
		toT := GongCopyBranchCodePilotControlledLightingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePositionInILSType:
		toT := GongCopyBranchCodePositionInILSType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePrimaryRadarType:
		toT := GongCopyBranchCodePrimaryRadarType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeProcedureAvailabilityType:
		toT := GongCopyBranchCodeProcedureAvailabilityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeProcedureCodingStandardType:
		toT := GongCopyBranchCodeProcedureCodingStandardType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeProcedureDistanceType:
		toT := GongCopyBranchCodeProcedureDistanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeProcedureFixRoleType:
		toT := GongCopyBranchCodeProcedureFixRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeProcedurePhaseType:
		toT := GongCopyBranchCodeProcedurePhaseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeProtectAreaSectionType:
		toT := GongCopyBranchCodeProtectAreaSectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRNPType:
		toT := GongCopyBranchCodeRNPType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVRReadingType:
		toT := GongCopyBranchCodeRVRReadingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVSMPointRoleType:
		toT := GongCopyBranchCodeRVSMPointRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVSMType:
		toT := GongCopyBranchCodeRVSMType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRadarServiceType:
		toT := GongCopyBranchCodeRadarServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRadioEmissionType:
		toT := GongCopyBranchCodeRadioEmissionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRadioFrequencyAreaType:
		toT := GongCopyBranchCodeRadioFrequencyAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRadioSignalType:
		toT := GongCopyBranchCodeRadioSignalType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeReferencePathIdentifierType:
		toT := GongCopyBranchCodeReferencePathIdentifierType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeReferenceRoleType:
		toT := GongCopyBranchCodeReferenceRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeReflectorType:
		toT := GongCopyBranchCodeReflectorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRelativePositionType:
		toT := GongCopyBranchCodeRelativePositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRoadType:
		toT := GongCopyBranchCodeRoadType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteAvailabilityType:
		toT := GongCopyBranchCodeRouteAvailabilityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteDesignatorLetterType:
		toT := GongCopyBranchCodeRouteDesignatorLetterType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteDesignatorPrefixType:
		toT := GongCopyBranchCodeRouteDesignatorPrefixType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteDesignatorSuffixType:
		toT := GongCopyBranchCodeRouteDesignatorSuffixType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteIndicatorType:
		toT := GongCopyBranchCodeRouteIndicatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteNavigationType:
		toT := GongCopyBranchCodeRouteNavigationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteOriginType:
		toT := GongCopyBranchCodeRouteOriginType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteSegmentPathType:
		toT := GongCopyBranchCodeRouteSegmentPathType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteType:
		toT := GongCopyBranchCodeRouteType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRuleProcedureTitleType:
		toT := GongCopyBranchCodeRuleProcedureTitleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRuleProcedureType:
		toT := GongCopyBranchCodeRuleProcedureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRunwayElementType:
		toT := GongCopyBranchCodeRunwayElementType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRunwayMarkingType:
		toT := GongCopyBranchCodeRunwayMarkingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRunwayPointRoleType:
		toT := GongCopyBranchCodeRunwayPointRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRunwayProtectionAreaType:
		toT := GongCopyBranchCodeRunwayProtectionAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRunwaySectionType:
		toT := GongCopyBranchCodeRunwaySectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRunwayType:
		toT := GongCopyBranchCodeRunwayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSafeAltitudeType:
		toT := GongCopyBranchCodeSafeAltitudeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSegmentPathType:
		toT := GongCopyBranchCodeSegmentPathType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSegmentTerminationType:
		toT := GongCopyBranchCodeSegmentTerminationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeServiceATCType:
		toT := GongCopyBranchCodeServiceATCType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeServiceATFMType:
		toT := GongCopyBranchCodeServiceATFMType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeServiceGroundControlType:
		toT := GongCopyBranchCodeServiceGroundControlType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeServiceInformationType:
		toT := GongCopyBranchCodeServiceInformationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeServiceSARType:
		toT := GongCopyBranchCodeServiceSARType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSideType:
		toT := GongCopyBranchCodeSideType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSignalPerformanceILSType:
		toT := GongCopyBranchCodeSignalPerformanceILSType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpecialDateType:
		toT := GongCopyBranchCodeSpecialDateType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpecialNavigationChainDesignatorType:
		toT := GongCopyBranchCodeSpecialNavigationChainDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpecialNavigationStationType:
		toT := GongCopyBranchCodeSpecialNavigationStationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpecialNavigationSystemType:
		toT := GongCopyBranchCodeSpecialNavigationSystemType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpeedReferenceType:
		toT := GongCopyBranchCodeSpeedReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStandbyPowerType:
		toT := GongCopyBranchCodeStandbyPowerType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusAirportType:
		toT := GongCopyBranchCodeStatusAirportType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusAirspaceType:
		toT := GongCopyBranchCodeStatusAirspaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusConstructionType:
		toT := GongCopyBranchCodeStatusConstructionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusNavaidType:
		toT := GongCopyBranchCodeStatusNavaidType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusOperationsType:
		toT := GongCopyBranchCodeStatusOperationsType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusServiceType:
		toT := GongCopyBranchCodeStatusServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSurfaceCompositionType:
		toT := GongCopyBranchCodeSurfaceCompositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSurfaceConditionType:
		toT := GongCopyBranchCodeSurfaceConditionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSurfacePreparationType:
		toT := GongCopyBranchCodeSurfacePreparationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSystemActivationType:
		toT := GongCopyBranchCodeSystemActivationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTAAType:
		toT := GongCopyBranchCodeTAAType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTACANChannelType:
		toT := GongCopyBranchCodeTACANChannelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTLOFSectionType:
		toT := GongCopyBranchCodeTLOFSectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTaxiwayElementType:
		toT := GongCopyBranchCodeTaxiwayElementType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTaxiwaySectionType:
		toT := GongCopyBranchCodeTaxiwaySectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTaxiwayType:
		toT := GongCopyBranchCodeTaxiwayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTelecomNetworkType:
		toT := GongCopyBranchCodeTelecomNetworkType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTimeEventCombinationType:
		toT := GongCopyBranchCodeTimeEventCombinationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTimeEventType:
		toT := GongCopyBranchCodeTimeEventType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTimeReferenceType:
		toT := GongCopyBranchCodeTimeReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTrajectoryType:
		toT := GongCopyBranchCodeTrajectoryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTransponderType:
		toT := GongCopyBranchCodeTransponderType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeUUIDType:
		toT := GongCopyBranchCodeUUIDType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeUnitDependencyType:
		toT := GongCopyBranchCodeUnitDependencyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeUnitType:
		toT := GongCopyBranchCodeUnitType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeUpperAlphaType:
		toT := GongCopyBranchCodeUpperAlphaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeUsageLimitationType:
		toT := GongCopyBranchCodeUsageLimitationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVASISType:
		toT := GongCopyBranchCodeVASISType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVORType:
		toT := GongCopyBranchCodeVORType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeValueInterpretationType:
		toT := GongCopyBranchCodeValueInterpretationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalDatumType:
		toT := GongCopyBranchCodeVerticalDatumType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalReferenceType:
		toT := GongCopyBranchCodeVerticalReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalStructureMarkingType:
		toT := GongCopyBranchCodeVerticalStructureMarkingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalStructureMaterialType:
		toT := GongCopyBranchCodeVerticalStructureMaterialType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalStructureType:
		toT := GongCopyBranchCodeVerticalStructureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVisualDockingGuidanceType:
		toT := GongCopyBranchCodeVisualDockingGuidanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeWakeTurbulenceType:
		toT := GongCopyBranchCodeWakeTurbulenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeWorkAreaType:
		toT := GongCopyBranchCodeWorkAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeYesNoType:
		toT := GongCopyBranchCodeYesNoType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DateMonthDayType:
		toT := GongCopyBranchDateMonthDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DateTimeType:
		toT := GongCopyBranchDateTimeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DateType:
		toT := GongCopyBranchDateType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DateYearType:
		toT := GongCopyBranchDateYearType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoNumberType:
		toT := GongCopyBranchNoNumberType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoSequenceType:
		toT := GongCopyBranchNoSequenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextAddressType:
		toT := GongCopyBranchTextAddressType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextDesignatorType:
		toT := GongCopyBranchTextDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextInstructionType:
		toT := GongCopyBranchTextInstructionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextNameType:
		toT := GongCopyBranchTextNameType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextNoteType:
		toT := GongCopyBranchTextNoteType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextPhoneType:
		toT := GongCopyBranchTextPhoneType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextPropertyNameType:
		toT := GongCopyBranchTextPropertyNameType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextRemarkType:
		toT := GongCopyBranchTextRemarkType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextSIDSTARDesignatorType:
		toT := GongCopyBranchTextSIDSTARDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimeType:
		toT := GongCopyBranchTimeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValAlarmLimitType:
		toT := GongCopyBranchValAlarmLimitType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValAngleType:
		toT := GongCopyBranchValAngleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValBearingType:
		toT := GongCopyBranchValBearingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValChannelNumberType:
		toT := GongCopyBranchValChannelNumberType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValDepthType:
		toT := GongCopyBranchValDepthType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValDistanceSignedType:
		toT := GongCopyBranchValDistanceSignedType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValDistanceType:
		toT := GongCopyBranchValDistanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValDistanceVerticalType:
		toT := GongCopyBranchValDistanceVerticalType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValDurationType:
		toT := GongCopyBranchValDurationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValFLType:
		toT := GongCopyBranchValFLType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValFrequencyType:
		toT := GongCopyBranchValFrequencyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValFrictionType:
		toT := GongCopyBranchValFrictionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValHexType:
		toT := GongCopyBranchValHexType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValLCNType:
		toT := GongCopyBranchValLCNType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValLightIntensityType:
		toT := GongCopyBranchValLightIntensityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValMagneticVariationChangeType:
		toT := GongCopyBranchValMagneticVariationChangeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValMagneticVariationType:
		toT := GongCopyBranchValMagneticVariationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValPCNType:
		toT := GongCopyBranchValPCNType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValPercentType:
		toT := GongCopyBranchValPercentType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValPressureType:
		toT := GongCopyBranchValPressureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValSlopeType:
		toT := GongCopyBranchValSlopeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValSpeedType:
		toT := GongCopyBranchValSpeedType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValTemperatureType:
		toT := GongCopyBranchValTemperatureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValWeightType:
		toT := GongCopyBranchValWeightType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XHTMLType:
		toT := GongCopyBranchXHTMLType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCodeATCReportingType(mapOrigCopy map[any]any, codeatcreportingtypeFrom *CodeATCReportingType) (codeatcreportingtypeTo *CodeATCReportingType) {
	var alreadyCopied bool
	codeatcreportingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeatcreportingtypeFrom)
	if alreadyCopied {
		return
	}
	codeatcreportingtypeFrom.GongCopyBasicFields(codeatcreportingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAerialRefuellingPointType(mapOrigCopy map[any]any, codeaerialrefuellingpointtypeFrom *CodeAerialRefuellingPointType) (codeaerialrefuellingpointtypeTo *CodeAerialRefuellingPointType) {
	var alreadyCopied bool
	codeaerialrefuellingpointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaerialrefuellingpointtypeFrom)
	if alreadyCopied {
		return
	}
	codeaerialrefuellingpointtypeFrom.GongCopyBasicFields(codeaerialrefuellingpointtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAerialRefuellingPrefixType(mapOrigCopy map[any]any, codeaerialrefuellingprefixtypeFrom *CodeAerialRefuellingPrefixType) (codeaerialrefuellingprefixtypeTo *CodeAerialRefuellingPrefixType) {
	var alreadyCopied bool
	codeaerialrefuellingprefixtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaerialrefuellingprefixtypeFrom)
	if alreadyCopied {
		return
	}
	codeaerialrefuellingprefixtypeFrom.GongCopyBasicFields(codeaerialrefuellingprefixtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAerialRefuellingType(mapOrigCopy map[any]any, codeaerialrefuellingtypeFrom *CodeAerialRefuellingType) (codeaerialrefuellingtypeTo *CodeAerialRefuellingType) {
	var alreadyCopied bool
	codeaerialrefuellingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaerialrefuellingtypeFrom)
	if alreadyCopied {
		return
	}
	codeaerialrefuellingtypeFrom.GongCopyBasicFields(codeaerialrefuellingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAircraftCategoryType(mapOrigCopy map[any]any, codeaircraftcategorytypeFrom *CodeAircraftCategoryType) (codeaircraftcategorytypeTo *CodeAircraftCategoryType) {
	var alreadyCopied bool
	codeaircraftcategorytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaircraftcategorytypeFrom)
	if alreadyCopied {
		return
	}
	codeaircraftcategorytypeFrom.GongCopyBasicFields(codeaircraftcategorytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAircraftEngineNumberType(mapOrigCopy map[any]any, codeaircraftenginenumbertypeFrom *CodeAircraftEngineNumberType) (codeaircraftenginenumbertypeTo *CodeAircraftEngineNumberType) {
	var alreadyCopied bool
	codeaircraftenginenumbertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaircraftenginenumbertypeFrom)
	if alreadyCopied {
		return
	}
	codeaircraftenginenumbertypeFrom.GongCopyBasicFields(codeaircraftenginenumbertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAircraftEngineType(mapOrigCopy map[any]any, codeaircraftenginetypeFrom *CodeAircraftEngineType) (codeaircraftenginetypeTo *CodeAircraftEngineType) {
	var alreadyCopied bool
	codeaircraftenginetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaircraftenginetypeFrom)
	if alreadyCopied {
		return
	}
	codeaircraftenginetypeFrom.GongCopyBasicFields(codeaircraftenginetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAircraftGroundServiceType(mapOrigCopy map[any]any, codeaircraftgroundservicetypeFrom *CodeAircraftGroundServiceType) (codeaircraftgroundservicetypeTo *CodeAircraftGroundServiceType) {
	var alreadyCopied bool
	codeaircraftgroundservicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaircraftgroundservicetypeFrom)
	if alreadyCopied {
		return
	}
	codeaircraftgroundservicetypeFrom.GongCopyBasicFields(codeaircraftgroundservicetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAircraftICAOType(mapOrigCopy map[any]any, codeaircrafticaotypeFrom *CodeAircraftICAOType) (codeaircrafticaotypeTo *CodeAircraftICAOType) {
	var alreadyCopied bool
	codeaircrafticaotypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaircrafticaotypeFrom)
	if alreadyCopied {
		return
	}
	codeaircrafticaotypeFrom.GongCopyBasicFields(codeaircrafticaotypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAircraftStandType(mapOrigCopy map[any]any, codeaircraftstandtypeFrom *CodeAircraftStandType) (codeaircraftstandtypeTo *CodeAircraftStandType) {
	var alreadyCopied bool
	codeaircraftstandtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaircraftstandtypeFrom)
	if alreadyCopied {
		return
	}
	codeaircraftstandtypeFrom.GongCopyBasicFields(codeaircraftstandtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAircraftType(mapOrigCopy map[any]any, codeaircrafttypeFrom *CodeAircraftType) (codeaircrafttypeTo *CodeAircraftType) {
	var alreadyCopied bool
	codeaircrafttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaircrafttypeFrom)
	if alreadyCopied {
		return
	}
	codeaircrafttypeFrom.GongCopyBasicFields(codeaircrafttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAircraftWingspanClassType(mapOrigCopy map[any]any, codeaircraftwingspanclasstypeFrom *CodeAircraftWingspanClassType) (codeaircraftwingspanclasstypeTo *CodeAircraftWingspanClassType) {
	var alreadyCopied bool
	codeaircraftwingspanclasstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaircraftwingspanclasstypeFrom)
	if alreadyCopied {
		return
	}
	codeaircraftwingspanclasstypeFrom.GongCopyBasicFields(codeaircraftwingspanclasstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirportHeliportCollocationType(mapOrigCopy map[any]any, codeairportheliportcollocationtypeFrom *CodeAirportHeliportCollocationType) (codeairportheliportcollocationtypeTo *CodeAirportHeliportCollocationType) {
	var alreadyCopied bool
	codeairportheliportcollocationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairportheliportcollocationtypeFrom)
	if alreadyCopied {
		return
	}
	codeairportheliportcollocationtypeFrom.GongCopyBasicFields(codeairportheliportcollocationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirportHeliportDesignatorType(mapOrigCopy map[any]any, codeairportheliportdesignatortypeFrom *CodeAirportHeliportDesignatorType) (codeairportheliportdesignatortypeTo *CodeAirportHeliportDesignatorType) {
	var alreadyCopied bool
	codeairportheliportdesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairportheliportdesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codeairportheliportdesignatortypeFrom.GongCopyBasicFields(codeairportheliportdesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirportHeliportType(mapOrigCopy map[any]any, codeairportheliporttypeFrom *CodeAirportHeliportType) (codeairportheliporttypeTo *CodeAirportHeliportType) {
	var alreadyCopied bool
	codeairportheliporttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairportheliporttypeFrom)
	if alreadyCopied {
		return
	}
	codeairportheliporttypeFrom.GongCopyBasicFields(codeairportheliporttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirportWarningType(mapOrigCopy map[any]any, codeairportwarningtypeFrom *CodeAirportWarningType) (codeairportwarningtypeTo *CodeAirportWarningType) {
	var alreadyCopied bool
	codeairportwarningtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairportwarningtypeFrom)
	if alreadyCopied {
		return
	}
	codeairportwarningtypeFrom.GongCopyBasicFields(codeairportwarningtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceActivityType(mapOrigCopy map[any]any, codeairspaceactivitytypeFrom *CodeAirspaceActivityType) (codeairspaceactivitytypeTo *CodeAirspaceActivityType) {
	var alreadyCopied bool
	codeairspaceactivitytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspaceactivitytypeFrom)
	if alreadyCopied {
		return
	}
	codeairspaceactivitytypeFrom.GongCopyBasicFields(codeairspaceactivitytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceAggregationType(mapOrigCopy map[any]any, codeairspaceaggregationtypeFrom *CodeAirspaceAggregationType) (codeairspaceaggregationtypeTo *CodeAirspaceAggregationType) {
	var alreadyCopied bool
	codeairspaceaggregationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspaceaggregationtypeFrom)
	if alreadyCopied {
		return
	}
	codeairspaceaggregationtypeFrom.GongCopyBasicFields(codeairspaceaggregationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceClassificationType(mapOrigCopy map[any]any, codeairspaceclassificationtypeFrom *CodeAirspaceClassificationType) (codeairspaceclassificationtypeTo *CodeAirspaceClassificationType) {
	var alreadyCopied bool
	codeairspaceclassificationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspaceclassificationtypeFrom)
	if alreadyCopied {
		return
	}
	codeairspaceclassificationtypeFrom.GongCopyBasicFields(codeairspaceclassificationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceDependencyType(mapOrigCopy map[any]any, codeairspacedependencytypeFrom *CodeAirspaceDependencyType) (codeairspacedependencytypeTo *CodeAirspaceDependencyType) {
	var alreadyCopied bool
	codeairspacedependencytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacedependencytypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacedependencytypeFrom.GongCopyBasicFields(codeairspacedependencytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceDesignatorType(mapOrigCopy map[any]any, codeairspacedesignatortypeFrom *CodeAirspaceDesignatorType) (codeairspacedesignatortypeTo *CodeAirspaceDesignatorType) {
	var alreadyCopied bool
	codeairspacedesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacedesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacedesignatortypeFrom.GongCopyBasicFields(codeairspacedesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspacePointPositionType(mapOrigCopy map[any]any, codeairspacepointpositiontypeFrom *CodeAirspacePointPositionType) (codeairspacepointpositiontypeTo *CodeAirspacePointPositionType) {
	var alreadyCopied bool
	codeairspacepointpositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacepointpositiontypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacepointpositiontypeFrom.GongCopyBasicFields(codeairspacepointpositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspacePointRoleType(mapOrigCopy map[any]any, codeairspacepointroletypeFrom *CodeAirspacePointRoleType) (codeairspacepointroletypeTo *CodeAirspacePointRoleType) {
	var alreadyCopied bool
	codeairspacepointroletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacepointroletypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacepointroletypeFrom.GongCopyBasicFields(codeairspacepointroletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceType(mapOrigCopy map[any]any, codeairspacetypeFrom *CodeAirspaceType) (codeairspacetypeTo *CodeAirspaceType) {
	var alreadyCopied bool
	codeairspacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacetypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacetypeFrom.GongCopyBasicFields(codeairspacetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAltitudeAdjustmentType(mapOrigCopy map[any]any, codealtitudeadjustmenttypeFrom *CodeAltitudeAdjustmentType) (codealtitudeadjustmenttypeTo *CodeAltitudeAdjustmentType) {
	var alreadyCopied bool
	codealtitudeadjustmenttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codealtitudeadjustmenttypeFrom)
	if alreadyCopied {
		return
	}
	codealtitudeadjustmenttypeFrom.GongCopyBasicFields(codealtitudeadjustmenttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAltitudeUseType(mapOrigCopy map[any]any, codealtitudeusetypeFrom *CodeAltitudeUseType) (codealtitudeusetypeTo *CodeAltitudeUseType) {
	var alreadyCopied bool
	codealtitudeusetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codealtitudeusetypeFrom)
	if alreadyCopied {
		return
	}
	codealtitudeusetypeFrom.GongCopyBasicFields(codealtitudeusetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApproachEquipmentAdditionalType(mapOrigCopy map[any]any, codeapproachequipmentadditionaltypeFrom *CodeApproachEquipmentAdditionalType) (codeapproachequipmentadditionaltypeTo *CodeApproachEquipmentAdditionalType) {
	var alreadyCopied bool
	codeapproachequipmentadditionaltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapproachequipmentadditionaltypeFrom)
	if alreadyCopied {
		return
	}
	codeapproachequipmentadditionaltypeFrom.GongCopyBasicFields(codeapproachequipmentadditionaltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApproachGuidanceType(mapOrigCopy map[any]any, codeapproachguidancetypeFrom *CodeApproachGuidanceType) (codeapproachguidancetypeTo *CodeApproachGuidanceType) {
	var alreadyCopied bool
	codeapproachguidancetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapproachguidancetypeFrom)
	if alreadyCopied {
		return
	}
	codeapproachguidancetypeFrom.GongCopyBasicFields(codeapproachguidancetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApproachLightingICAOType(mapOrigCopy map[any]any, codeapproachlightingicaotypeFrom *CodeApproachLightingICAOType) (codeapproachlightingicaotypeTo *CodeApproachLightingICAOType) {
	var alreadyCopied bool
	codeapproachlightingicaotypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapproachlightingicaotypeFrom)
	if alreadyCopied {
		return
	}
	codeapproachlightingicaotypeFrom.GongCopyBasicFields(codeapproachlightingicaotypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApproachLightingType(mapOrigCopy map[any]any, codeapproachlightingtypeFrom *CodeApproachLightingType) (codeapproachlightingtypeTo *CodeApproachLightingType) {
	var alreadyCopied bool
	codeapproachlightingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapproachlightingtypeFrom)
	if alreadyCopied {
		return
	}
	codeapproachlightingtypeFrom.GongCopyBasicFields(codeapproachlightingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApproachPrefixType(mapOrigCopy map[any]any, codeapproachprefixtypeFrom *CodeApproachPrefixType) (codeapproachprefixtypeTo *CodeApproachPrefixType) {
	var alreadyCopied bool
	codeapproachprefixtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapproachprefixtypeFrom)
	if alreadyCopied {
		return
	}
	codeapproachprefixtypeFrom.GongCopyBasicFields(codeapproachprefixtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApproachType(mapOrigCopy map[any]any, codeapproachtypeFrom *CodeApproachType) (codeapproachtypeTo *CodeApproachType) {
	var alreadyCopied bool
	codeapproachtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapproachtypeFrom)
	if alreadyCopied {
		return
	}
	codeapproachtypeFrom.GongCopyBasicFields(codeapproachtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApprovalType(mapOrigCopy map[any]any, codeapprovaltypeFrom *CodeApprovalType) (codeapprovaltypeTo *CodeApprovalType) {
	var alreadyCopied bool
	codeapprovaltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapprovaltypeFrom)
	if alreadyCopied {
		return
	}
	codeapprovaltypeFrom.GongCopyBasicFields(codeapprovaltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApronElementType(mapOrigCopy map[any]any, codeapronelementtypeFrom *CodeApronElementType) (codeapronelementtypeTo *CodeApronElementType) {
	var alreadyCopied bool
	codeapronelementtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapronelementtypeFrom)
	if alreadyCopied {
		return
	}
	codeapronelementtypeFrom.GongCopyBasicFields(codeapronelementtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeApronSectionType(mapOrigCopy map[any]any, codeapronsectiontypeFrom *CodeApronSectionType) (codeapronsectiontypeTo *CodeApronSectionType) {
	var alreadyCopied bool
	codeapronsectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeapronsectiontypeFrom)
	if alreadyCopied {
		return
	}
	codeapronsectiontypeFrom.GongCopyBasicFields(codeapronsectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeArcDirectionType(mapOrigCopy map[any]any, codearcdirectiontypeFrom *CodeArcDirectionType) (codearcdirectiontypeTo *CodeArcDirectionType) {
	var alreadyCopied bool
	codearcdirectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codearcdirectiontypeFrom)
	if alreadyCopied {
		return
	}
	codearcdirectiontypeFrom.GongCopyBasicFields(codearcdirectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeArrestingGearEnergyAbsorbType(mapOrigCopy map[any]any, codearrestinggearenergyabsorbtypeFrom *CodeArrestingGearEnergyAbsorbType) (codearrestinggearenergyabsorbtypeTo *CodeArrestingGearEnergyAbsorbType) {
	var alreadyCopied bool
	codearrestinggearenergyabsorbtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codearrestinggearenergyabsorbtypeFrom)
	if alreadyCopied {
		return
	}
	codearrestinggearenergyabsorbtypeFrom.GongCopyBasicFields(codearrestinggearenergyabsorbtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeArrestingGearEngageDeviceType(mapOrigCopy map[any]any, codearrestinggearengagedevicetypeFrom *CodeArrestingGearEngageDeviceType) (codearrestinggearengagedevicetypeTo *CodeArrestingGearEngageDeviceType) {
	var alreadyCopied bool
	codearrestinggearengagedevicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codearrestinggearengagedevicetypeFrom)
	if alreadyCopied {
		return
	}
	codearrestinggearengagedevicetypeFrom.GongCopyBasicFields(codearrestinggearengagedevicetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAuralMorseType(mapOrigCopy map[any]any, codeauralmorsetypeFrom *CodeAuralMorseType) (codeauralmorsetypeTo *CodeAuralMorseType) {
	var alreadyCopied bool
	codeauralmorsetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeauralmorsetypeFrom)
	if alreadyCopied {
		return
	}
	codeauralmorsetypeFrom.GongCopyBasicFields(codeauralmorsetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAuthorityRoleType(mapOrigCopy map[any]any, codeauthorityroletypeFrom *CodeAuthorityRoleType) (codeauthorityroletypeTo *CodeAuthorityRoleType) {
	var alreadyCopied bool
	codeauthorityroletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeauthorityroletypeFrom)
	if alreadyCopied {
		return
	}
	codeauthorityroletypeFrom.GongCopyBasicFields(codeauthorityroletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAuthorityType(mapOrigCopy map[any]any, codeauthoritytypeFrom *CodeAuthorityType) (codeauthoritytypeTo *CodeAuthorityType) {
	var alreadyCopied bool
	codeauthoritytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeauthoritytypeFrom)
	if alreadyCopied {
		return
	}
	codeauthoritytypeFrom.GongCopyBasicFields(codeauthoritytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAviationStandardsType(mapOrigCopy map[any]any, codeaviationstandardstypeFrom *CodeAviationStandardsType) (codeaviationstandardstypeTo *CodeAviationStandardsType) {
	var alreadyCopied bool
	codeaviationstandardstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeaviationstandardstypeFrom)
	if alreadyCopied {
		return
	}
	codeaviationstandardstypeFrom.GongCopyBasicFields(codeaviationstandardstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeBearingType(mapOrigCopy map[any]any, codebearingtypeFrom *CodeBearingType) (codebearingtypeTo *CodeBearingType) {
	var alreadyCopied bool
	codebearingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codebearingtypeFrom)
	if alreadyCopied {
		return
	}
	codebearingtypeFrom.GongCopyBasicFields(codebearingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeBuoyDesignatorType(mapOrigCopy map[any]any, codebuoydesignatortypeFrom *CodeBuoyDesignatorType) (codebuoydesignatortypeTo *CodeBuoyDesignatorType) {
	var alreadyCopied bool
	codebuoydesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codebuoydesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codebuoydesignatortypeFrom.GongCopyBasicFields(codebuoydesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeBuoyType(mapOrigCopy map[any]any, codebuoytypeFrom *CodeBuoyType) (codebuoytypeTo *CodeBuoyType) {
	var alreadyCopied bool
	codebuoytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codebuoytypeFrom)
	if alreadyCopied {
		return
	}
	codebuoytypeFrom.GongCopyBasicFields(codebuoytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeCardinalDirectionType(mapOrigCopy map[any]any, codecardinaldirectiontypeFrom *CodeCardinalDirectionType) (codecardinaldirectiontypeTo *CodeCardinalDirectionType) {
	var alreadyCopied bool
	codecardinaldirectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecardinaldirectiontypeFrom)
	if alreadyCopied {
		return
	}
	codecardinaldirectiontypeFrom.GongCopyBasicFields(codecardinaldirectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeCheckpointCategoryType(mapOrigCopy map[any]any, codecheckpointcategorytypeFrom *CodeCheckpointCategoryType) (codecheckpointcategorytypeTo *CodeCheckpointCategoryType) {
	var alreadyCopied bool
	codecheckpointcategorytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecheckpointcategorytypeFrom)
	if alreadyCopied {
		return
	}
	codecheckpointcategorytypeFrom.GongCopyBasicFields(codecheckpointcategorytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeColourType(mapOrigCopy map[any]any, codecolourtypeFrom *CodeColourType) (codecolourtypeTo *CodeColourType) {
	var alreadyCopied bool
	codecolourtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecolourtypeFrom)
	if alreadyCopied {
		return
	}
	codecolourtypeFrom.GongCopyBasicFields(codecolourtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeCommunicationChannelType(mapOrigCopy map[any]any, codecommunicationchanneltypeFrom *CodeCommunicationChannelType) (codecommunicationchanneltypeTo *CodeCommunicationChannelType) {
	var alreadyCopied bool
	codecommunicationchanneltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecommunicationchanneltypeFrom)
	if alreadyCopied {
		return
	}
	codecommunicationchanneltypeFrom.GongCopyBasicFields(codecommunicationchanneltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeCommunicationDirectionType(mapOrigCopy map[any]any, codecommunicationdirectiontypeFrom *CodeCommunicationDirectionType) (codecommunicationdirectiontypeTo *CodeCommunicationDirectionType) {
	var alreadyCopied bool
	codecommunicationdirectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecommunicationdirectiontypeFrom)
	if alreadyCopied {
		return
	}
	codecommunicationdirectiontypeFrom.GongCopyBasicFields(codecommunicationdirectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeCommunicationModeType(mapOrigCopy map[any]any, codecommunicationmodetypeFrom *CodeCommunicationModeType) (codecommunicationmodetypeTo *CodeCommunicationModeType) {
	var alreadyCopied bool
	codecommunicationmodetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecommunicationmodetypeFrom)
	if alreadyCopied {
		return
	}
	codecommunicationmodetypeFrom.GongCopyBasicFields(codecommunicationmodetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeComparisonType(mapOrigCopy map[any]any, codecomparisontypeFrom *CodeComparisonType) (codecomparisontypeTo *CodeComparisonType) {
	var alreadyCopied bool
	codecomparisontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecomparisontypeFrom)
	if alreadyCopied {
		return
	}
	codecomparisontypeFrom.GongCopyBasicFields(codecomparisontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeContaminationType(mapOrigCopy map[any]any, codecontaminationtypeFrom *CodeContaminationType) (codecontaminationtypeTo *CodeContaminationType) {
	var alreadyCopied bool
	codecontaminationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecontaminationtypeFrom)
	if alreadyCopied {
		return
	}
	codecontaminationtypeFrom.GongCopyBasicFields(codecontaminationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeCourseQualityILSType(mapOrigCopy map[any]any, codecoursequalityilstypeFrom *CodeCourseQualityILSType) (codecoursequalityilstypeTo *CodeCourseQualityILSType) {
	var alreadyCopied bool
	codecoursequalityilstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecoursequalityilstypeFrom)
	if alreadyCopied {
		return
	}
	codecoursequalityilstypeFrom.GongCopyBasicFields(codecoursequalityilstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeCourseType(mapOrigCopy map[any]any, codecoursetypeFrom *CodeCourseType) (codecoursetypeTo *CodeCourseType) {
	var alreadyCopied bool
	codecoursetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecoursetypeFrom)
	if alreadyCopied {
		return
	}
	codecoursetypeFrom.GongCopyBasicFields(codecoursetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDMEChannelType(mapOrigCopy map[any]any, codedmechanneltypeFrom *CodeDMEChannelType) (codedmechanneltypeTo *CodeDMEChannelType) {
	var alreadyCopied bool
	codedmechanneltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedmechanneltypeFrom)
	if alreadyCopied {
		return
	}
	codedmechanneltypeFrom.GongCopyBasicFields(codedmechanneltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDMEType(mapOrigCopy map[any]any, codedmetypeFrom *CodeDMEType) (codedmetypeTo *CodeDMEType) {
	var alreadyCopied bool
	codedmetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedmetypeFrom)
	if alreadyCopied {
		return
	}
	codedmetypeFrom.GongCopyBasicFields(codedmetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDayType(mapOrigCopy map[any]any, codedaytypeFrom *CodeDayType) (codedaytypeTo *CodeDayType) {
	var alreadyCopied bool
	codedaytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedaytypeFrom)
	if alreadyCopied {
		return
	}
	codedaytypeFrom.GongCopyBasicFields(codedaytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDeclaredDistanceType(mapOrigCopy map[any]any, codedeclareddistancetypeFrom *CodeDeclaredDistanceType) (codedeclareddistancetypeTo *CodeDeclaredDistanceType) {
	var alreadyCopied bool
	codedeclareddistancetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedeclareddistancetypeFrom)
	if alreadyCopied {
		return
	}
	codedeclareddistancetypeFrom.GongCopyBasicFields(codedeclareddistancetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDesignStandardType(mapOrigCopy map[any]any, codedesignstandardtypeFrom *CodeDesignStandardType) (codedesignstandardtypeTo *CodeDesignStandardType) {
	var alreadyCopied bool
	codedesignstandardtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedesignstandardtypeFrom)
	if alreadyCopied {
		return
	}
	codedesignstandardtypeFrom.GongCopyBasicFields(codedesignstandardtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDesignatedPointDesignatorType(mapOrigCopy map[any]any, codedesignatedpointdesignatortypeFrom *CodeDesignatedPointDesignatorType) (codedesignatedpointdesignatortypeTo *CodeDesignatedPointDesignatorType) {
	var alreadyCopied bool
	codedesignatedpointdesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedesignatedpointdesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codedesignatedpointdesignatortypeFrom.GongCopyBasicFields(codedesignatedpointdesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDesignatedPointType(mapOrigCopy map[any]any, codedesignatedpointtypeFrom *CodeDesignatedPointType) (codedesignatedpointtypeTo *CodeDesignatedPointType) {
	var alreadyCopied bool
	codedesignatedpointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedesignatedpointtypeFrom)
	if alreadyCopied {
		return
	}
	codedesignatedpointtypeFrom.GongCopyBasicFields(codedesignatedpointtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDirectionReferenceType(mapOrigCopy map[any]any, codedirectionreferencetypeFrom *CodeDirectionReferenceType) (codedirectionreferencetypeTo *CodeDirectionReferenceType) {
	var alreadyCopied bool
	codedirectionreferencetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedirectionreferencetypeFrom)
	if alreadyCopied {
		return
	}
	codedirectionreferencetypeFrom.GongCopyBasicFields(codedirectionreferencetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDirectionTurnType(mapOrigCopy map[any]any, codedirectionturntypeFrom *CodeDirectionTurnType) (codedirectionturntypeTo *CodeDirectionTurnType) {
	var alreadyCopied bool
	codedirectionturntypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedirectionturntypeFrom)
	if alreadyCopied {
		return
	}
	codedirectionturntypeFrom.GongCopyBasicFields(codedirectionturntypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDirectionType(mapOrigCopy map[any]any, codedirectiontypeFrom *CodeDirectionType) (codedirectiontypeTo *CodeDirectionType) {
	var alreadyCopied bool
	codedirectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedirectiontypeFrom)
	if alreadyCopied {
		return
	}
	codedirectiontypeFrom.GongCopyBasicFields(codedirectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDistanceIndicationType(mapOrigCopy map[any]any, codedistanceindicationtypeFrom *CodeDistanceIndicationType) (codedistanceindicationtypeTo *CodeDistanceIndicationType) {
	var alreadyCopied bool
	codedistanceindicationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedistanceindicationtypeFrom)
	if alreadyCopied {
		return
	}
	codedistanceindicationtypeFrom.GongCopyBasicFields(codedistanceindicationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeEmissionBandType(mapOrigCopy map[any]any, codeemissionbandtypeFrom *CodeEmissionBandType) (codeemissionbandtypeTo *CodeEmissionBandType) {
	var alreadyCopied bool
	codeemissionbandtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeemissionbandtypeFrom)
	if alreadyCopied {
		return
	}
	codeemissionbandtypeFrom.GongCopyBasicFields(codeemissionbandtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeEquipmentAntiCollisionType(mapOrigCopy map[any]any, codeequipmentanticollisiontypeFrom *CodeEquipmentAntiCollisionType) (codeequipmentanticollisiontypeTo *CodeEquipmentAntiCollisionType) {
	var alreadyCopied bool
	codeequipmentanticollisiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeequipmentanticollisiontypeFrom)
	if alreadyCopied {
		return
	}
	codeequipmentanticollisiontypeFrom.GongCopyBasicFields(codeequipmentanticollisiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeEquipmentUnavailableType(mapOrigCopy map[any]any, codeequipmentunavailabletypeFrom *CodeEquipmentUnavailableType) (codeequipmentunavailabletypeTo *CodeEquipmentUnavailableType) {
	var alreadyCopied bool
	codeequipmentunavailabletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeequipmentunavailabletypeFrom)
	if alreadyCopied {
		return
	}
	codeequipmentunavailabletypeFrom.GongCopyBasicFields(codeequipmentunavailabletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFacilityRankingType(mapOrigCopy map[any]any, codefacilityrankingtypeFrom *CodeFacilityRankingType) (codefacilityrankingtypeTo *CodeFacilityRankingType) {
	var alreadyCopied bool
	codefacilityrankingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codefacilityrankingtypeFrom)
	if alreadyCopied {
		return
	}
	codefacilityrankingtypeFrom.GongCopyBasicFields(codefacilityrankingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFinalGuidanceType(mapOrigCopy map[any]any, codefinalguidancetypeFrom *CodeFinalGuidanceType) (codefinalguidancetypeTo *CodeFinalGuidanceType) {
	var alreadyCopied bool
	codefinalguidancetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codefinalguidancetypeFrom)
	if alreadyCopied {
		return
	}
	codefinalguidancetypeFrom.GongCopyBasicFields(codefinalguidancetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFireFightingType(mapOrigCopy map[any]any, codefirefightingtypeFrom *CodeFireFightingType) (codefirefightingtypeTo *CodeFireFightingType) {
	var alreadyCopied bool
	codefirefightingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codefirefightingtypeFrom)
	if alreadyCopied {
		return
	}
	codefirefightingtypeFrom.GongCopyBasicFields(codefirefightingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlightDestinationType(mapOrigCopy map[any]any, codeflightdestinationtypeFrom *CodeFlightDestinationType) (codeflightdestinationtypeTo *CodeFlightDestinationType) {
	var alreadyCopied bool
	codeflightdestinationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflightdestinationtypeFrom)
	if alreadyCopied {
		return
	}
	codeflightdestinationtypeFrom.GongCopyBasicFields(codeflightdestinationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlightOriginType(mapOrigCopy map[any]any, codeflightorigintypeFrom *CodeFlightOriginType) (codeflightorigintypeTo *CodeFlightOriginType) {
	var alreadyCopied bool
	codeflightorigintypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflightorigintypeFrom)
	if alreadyCopied {
		return
	}
	codeflightorigintypeFrom.GongCopyBasicFields(codeflightorigintypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlightPurposeType(mapOrigCopy map[any]any, codeflightpurposetypeFrom *CodeFlightPurposeType) (codeflightpurposetypeTo *CodeFlightPurposeType) {
	var alreadyCopied bool
	codeflightpurposetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflightpurposetypeFrom)
	if alreadyCopied {
		return
	}
	codeflightpurposetypeFrom.GongCopyBasicFields(codeflightpurposetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlightRestrictionDesignatorType(mapOrigCopy map[any]any, codeflightrestrictiondesignatortypeFrom *CodeFlightRestrictionDesignatorType) (codeflightrestrictiondesignatortypeTo *CodeFlightRestrictionDesignatorType) {
	var alreadyCopied bool
	codeflightrestrictiondesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflightrestrictiondesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codeflightrestrictiondesignatortypeFrom.GongCopyBasicFields(codeflightrestrictiondesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlightRestrictionType(mapOrigCopy map[any]any, codeflightrestrictiontypeFrom *CodeFlightRestrictionType) (codeflightrestrictiontypeTo *CodeFlightRestrictionType) {
	var alreadyCopied bool
	codeflightrestrictiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflightrestrictiontypeFrom)
	if alreadyCopied {
		return
	}
	codeflightrestrictiontypeFrom.GongCopyBasicFields(codeflightrestrictiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlightRuleType(mapOrigCopy map[any]any, codeflightruletypeFrom *CodeFlightRuleType) (codeflightruletypeTo *CodeFlightRuleType) {
	var alreadyCopied bool
	codeflightruletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflightruletypeFrom)
	if alreadyCopied {
		return
	}
	codeflightruletypeFrom.GongCopyBasicFields(codeflightruletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlightStatusType(mapOrigCopy map[any]any, codeflightstatustypeFrom *CodeFlightStatusType) (codeflightstatustypeTo *CodeFlightStatusType) {
	var alreadyCopied bool
	codeflightstatustypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflightstatustypeFrom)
	if alreadyCopied {
		return
	}
	codeflightstatustypeFrom.GongCopyBasicFields(codeflightstatustypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlightType(mapOrigCopy map[any]any, codeflighttypeFrom *CodeFlightType) (codeflighttypeTo *CodeFlightType) {
	var alreadyCopied bool
	codeflighttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflighttypeFrom)
	if alreadyCopied {
		return
	}
	codeflighttypeFrom.GongCopyBasicFields(codeflighttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFlowConditionOperationType(mapOrigCopy map[any]any, codeflowconditionoperationtypeFrom *CodeFlowConditionOperationType) (codeflowconditionoperationtypeTo *CodeFlowConditionOperationType) {
	var alreadyCopied bool
	codeflowconditionoperationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeflowconditionoperationtypeFrom)
	if alreadyCopied {
		return
	}
	codeflowconditionoperationtypeFrom.GongCopyBasicFields(codeflowconditionoperationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFreeFlightType(mapOrigCopy map[any]any, codefreeflighttypeFrom *CodeFreeFlightType) (codefreeflighttypeTo *CodeFreeFlightType) {
	var alreadyCopied bool
	codefreeflighttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codefreeflighttypeFrom)
	if alreadyCopied {
		return
	}
	codefreeflighttypeFrom.GongCopyBasicFields(codefreeflighttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFrictionDeviceType(mapOrigCopy map[any]any, codefrictiondevicetypeFrom *CodeFrictionDeviceType) (codefrictiondevicetypeTo *CodeFrictionDeviceType) {
	var alreadyCopied bool
	codefrictiondevicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codefrictiondevicetypeFrom)
	if alreadyCopied {
		return
	}
	codefrictiondevicetypeFrom.GongCopyBasicFields(codefrictiondevicetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFrictionEstimateType(mapOrigCopy map[any]any, codefrictionestimatetypeFrom *CodeFrictionEstimateType) (codefrictionestimatetypeTo *CodeFrictionEstimateType) {
	var alreadyCopied bool
	codefrictionestimatetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codefrictionestimatetypeFrom)
	if alreadyCopied {
		return
	}
	codefrictionestimatetypeFrom.GongCopyBasicFields(codefrictionestimatetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeFuelType(mapOrigCopy map[any]any, codefueltypeFrom *CodeFuelType) (codefueltypeTo *CodeFuelType) {
	var alreadyCopied bool
	codefueltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codefueltypeFrom)
	if alreadyCopied {
		return
	}
	codefueltypeFrom.GongCopyBasicFields(codefueltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeGeoBorderType(mapOrigCopy map[any]any, codegeobordertypeFrom *CodeGeoBorderType) (codegeobordertypeTo *CodeGeoBorderType) {
	var alreadyCopied bool
	codegeobordertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codegeobordertypeFrom)
	if alreadyCopied {
		return
	}
	codegeobordertypeFrom.GongCopyBasicFields(codegeobordertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeGradeSeparationType(mapOrigCopy map[any]any, codegradeseparationtypeFrom *CodeGradeSeparationType) (codegradeseparationtypeTo *CodeGradeSeparationType) {
	var alreadyCopied bool
	codegradeseparationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codegradeseparationtypeFrom)
	if alreadyCopied {
		return
	}
	codegradeseparationtypeFrom.GongCopyBasicFields(codegradeseparationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeGroundLightingType(mapOrigCopy map[any]any, codegroundlightingtypeFrom *CodeGroundLightingType) (codegroundlightingtypeTo *CodeGroundLightingType) {
	var alreadyCopied bool
	codegroundlightingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codegroundlightingtypeFrom)
	if alreadyCopied {
		return
	}
	codegroundlightingtypeFrom.GongCopyBasicFields(codegroundlightingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeGuidanceLineType(mapOrigCopy map[any]any, codeguidancelinetypeFrom *CodeGuidanceLineType) (codeguidancelinetypeTo *CodeGuidanceLineType) {
	var alreadyCopied bool
	codeguidancelinetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeguidancelinetypeFrom)
	if alreadyCopied {
		return
	}
	codeguidancelinetypeFrom.GongCopyBasicFields(codeguidancelinetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeHeightReferenceType(mapOrigCopy map[any]any, codeheightreferencetypeFrom *CodeHeightReferenceType) (codeheightreferencetypeTo *CodeHeightReferenceType) {
	var alreadyCopied bool
	codeheightreferencetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeheightreferencetypeFrom)
	if alreadyCopied {
		return
	}
	codeheightreferencetypeFrom.GongCopyBasicFields(codeheightreferencetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeHelicopterPerformanceType(mapOrigCopy map[any]any, codehelicopterperformancetypeFrom *CodeHelicopterPerformanceType) (codehelicopterperformancetypeTo *CodeHelicopterPerformanceType) {
	var alreadyCopied bool
	codehelicopterperformancetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codehelicopterperformancetypeFrom)
	if alreadyCopied {
		return
	}
	codehelicopterperformancetypeFrom.GongCopyBasicFields(codehelicopterperformancetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeHoldingCategoryType(mapOrigCopy map[any]any, codeholdingcategorytypeFrom *CodeHoldingCategoryType) (codeholdingcategorytypeTo *CodeHoldingCategoryType) {
	var alreadyCopied bool
	codeholdingcategorytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeholdingcategorytypeFrom)
	if alreadyCopied {
		return
	}
	codeholdingcategorytypeFrom.GongCopyBasicFields(codeholdingcategorytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeHoldingUsageType(mapOrigCopy map[any]any, codeholdingusagetypeFrom *CodeHoldingUsageType) (codeholdingusagetypeTo *CodeHoldingUsageType) {
	var alreadyCopied bool
	codeholdingusagetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeholdingusagetypeFrom)
	if alreadyCopied {
		return
	}
	codeholdingusagetypeFrom.GongCopyBasicFields(codeholdingusagetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeHoldingUseType(mapOrigCopy map[any]any, codeholdingusetypeFrom *CodeHoldingUseType) (codeholdingusetypeTo *CodeHoldingUseType) {
	var alreadyCopied bool
	codeholdingusetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeholdingusetypeFrom)
	if alreadyCopied {
		return
	}
	codeholdingusetypeFrom.GongCopyBasicFields(codeholdingusetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeIATAType(mapOrigCopy map[any]any, codeiatatypeFrom *CodeIATAType) (codeiatatypeTo *CodeIATAType) {
	var alreadyCopied bool
	codeiatatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeiatatypeFrom)
	if alreadyCopied {
		return
	}
	codeiatatypeFrom.GongCopyBasicFields(codeiatatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeICAOCountryType(mapOrigCopy map[any]any, codeicaocountrytypeFrom *CodeICAOCountryType) (codeicaocountrytypeTo *CodeICAOCountryType) {
	var alreadyCopied bool
	codeicaocountrytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeicaocountrytypeFrom)
	if alreadyCopied {
		return
	}
	codeicaocountrytypeFrom.GongCopyBasicFields(codeicaocountrytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeICAOType(mapOrigCopy map[any]any, codeicaotypeFrom *CodeICAOType) (codeicaotypeTo *CodeICAOType) {
	var alreadyCopied bool
	codeicaotypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeicaotypeFrom)
	if alreadyCopied {
		return
	}
	codeicaotypeFrom.GongCopyBasicFields(codeicaotypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeILSBackCourseType(mapOrigCopy map[any]any, codeilsbackcoursetypeFrom *CodeILSBackCourseType) (codeilsbackcoursetypeTo *CodeILSBackCourseType) {
	var alreadyCopied bool
	codeilsbackcoursetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeilsbackcoursetypeFrom)
	if alreadyCopied {
		return
	}
	codeilsbackcoursetypeFrom.GongCopyBasicFields(codeilsbackcoursetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeIntegrityLevelILSType(mapOrigCopy map[any]any, codeintegritylevelilstypeFrom *CodeIntegrityLevelILSType) (codeintegritylevelilstypeTo *CodeIntegrityLevelILSType) {
	var alreadyCopied bool
	codeintegritylevelilstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeintegritylevelilstypeFrom)
	if alreadyCopied {
		return
	}
	codeintegritylevelilstypeFrom.GongCopyBasicFields(codeintegritylevelilstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeIntensityStandByType(mapOrigCopy map[any]any, codeintensitystandbytypeFrom *CodeIntensityStandByType) (codeintensitystandbytypeTo *CodeIntensityStandByType) {
	var alreadyCopied bool
	codeintensitystandbytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeintensitystandbytypeFrom)
	if alreadyCopied {
		return
	}
	codeintensitystandbytypeFrom.GongCopyBasicFields(codeintensitystandbytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLanguageType(mapOrigCopy map[any]any, codelanguagetypeFrom *CodeLanguageType) (codelanguagetypeTo *CodeLanguageType) {
	var alreadyCopied bool
	codelanguagetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelanguagetypeFrom)
	if alreadyCopied {
		return
	}
	codelanguagetypeFrom.GongCopyBasicFields(codelanguagetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLevelSeriesType(mapOrigCopy map[any]any, codelevelseriestypeFrom *CodeLevelSeriesType) (codelevelseriestypeTo *CodeLevelSeriesType) {
	var alreadyCopied bool
	codelevelseriestypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelevelseriestypeFrom)
	if alreadyCopied {
		return
	}
	codelevelseriestypeFrom.GongCopyBasicFields(codelevelseriestypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLevelTableDesignatorType(mapOrigCopy map[any]any, codeleveltabledesignatortypeFrom *CodeLevelTableDesignatorType) (codeleveltabledesignatortypeTo *CodeLevelTableDesignatorType) {
	var alreadyCopied bool
	codeleveltabledesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeleveltabledesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codeleveltabledesignatortypeFrom.GongCopyBasicFields(codeleveltabledesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLevelType(mapOrigCopy map[any]any, codeleveltypeFrom *CodeLevelType) (codeleveltypeTo *CodeLevelType) {
	var alreadyCopied bool
	codeleveltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeleveltypeFrom)
	if alreadyCopied {
		return
	}
	codeleveltypeFrom.GongCopyBasicFields(codeleveltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLightHoldingPositionType(mapOrigCopy map[any]any, codelightholdingpositiontypeFrom *CodeLightHoldingPositionType) (codelightholdingpositiontypeTo *CodeLightHoldingPositionType) {
	var alreadyCopied bool
	codelightholdingpositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelightholdingpositiontypeFrom)
	if alreadyCopied {
		return
	}
	codelightholdingpositiontypeFrom.GongCopyBasicFields(codelightholdingpositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLightIntensityType(mapOrigCopy map[any]any, codelightintensitytypeFrom *CodeLightIntensityType) (codelightintensitytypeTo *CodeLightIntensityType) {
	var alreadyCopied bool
	codelightintensitytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelightintensitytypeFrom)
	if alreadyCopied {
		return
	}
	codelightintensitytypeFrom.GongCopyBasicFields(codelightintensitytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLightSourceType(mapOrigCopy map[any]any, codelightsourcetypeFrom *CodeLightSourceType) (codelightsourcetypeTo *CodeLightSourceType) {
	var alreadyCopied bool
	codelightsourcetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelightsourcetypeFrom)
	if alreadyCopied {
		return
	}
	codelightsourcetypeFrom.GongCopyBasicFields(codelightsourcetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLightingJARType(mapOrigCopy map[any]any, codelightingjartypeFrom *CodeLightingJARType) (codelightingjartypeTo *CodeLightingJARType) {
	var alreadyCopied bool
	codelightingjartypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelightingjartypeFrom)
	if alreadyCopied {
		return
	}
	codelightingjartypeFrom.GongCopyBasicFields(codelightingjartypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLoadingBridgeType(mapOrigCopy map[any]any, codeloadingbridgetypeFrom *CodeLoadingBridgeType) (codeloadingbridgetypeTo *CodeLoadingBridgeType) {
	var alreadyCopied bool
	codeloadingbridgetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeloadingbridgetypeFrom)
	if alreadyCopied {
		return
	}
	codeloadingbridgetypeFrom.GongCopyBasicFields(codeloadingbridgetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLocationQualifierType(mapOrigCopy map[any]any, codelocationqualifiertypeFrom *CodeLocationQualifierType) (codelocationqualifiertypeTo *CodeLocationQualifierType) {
	var alreadyCopied bool
	codelocationqualifiertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelocationqualifiertypeFrom)
	if alreadyCopied {
		return
	}
	codelocationqualifiertypeFrom.GongCopyBasicFields(codelocationqualifiertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeLogicalOperatorType(mapOrigCopy map[any]any, codelogicaloperatortypeFrom *CodeLogicalOperatorType) (codelogicaloperatortypeTo *CodeLogicalOperatorType) {
	var alreadyCopied bool
	codelogicaloperatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelogicaloperatortypeFrom)
	if alreadyCopied {
		return
	}
	codelogicaloperatortypeFrom.GongCopyBasicFields(codelogicaloperatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMLSAzimuthType(mapOrigCopy map[any]any, codemlsazimuthtypeFrom *CodeMLSAzimuthType) (codemlsazimuthtypeTo *CodeMLSAzimuthType) {
	var alreadyCopied bool
	codemlsazimuthtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemlsazimuthtypeFrom)
	if alreadyCopied {
		return
	}
	codemlsazimuthtypeFrom.GongCopyBasicFields(codemlsazimuthtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMLSChannelType(mapOrigCopy map[any]any, codemlschanneltypeFrom *CodeMLSChannelType) (codemlschanneltypeTo *CodeMLSChannelType) {
	var alreadyCopied bool
	codemlschanneltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemlschanneltypeFrom)
	if alreadyCopied {
		return
	}
	codemlschanneltypeFrom.GongCopyBasicFields(codemlschanneltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMarkerBeaconSignalType(mapOrigCopy map[any]any, codemarkerbeaconsignaltypeFrom *CodeMarkerBeaconSignalType) (codemarkerbeaconsignaltypeTo *CodeMarkerBeaconSignalType) {
	var alreadyCopied bool
	codemarkerbeaconsignaltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemarkerbeaconsignaltypeFrom)
	if alreadyCopied {
		return
	}
	codemarkerbeaconsignaltypeFrom.GongCopyBasicFields(codemarkerbeaconsignaltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMarkingConditionType(mapOrigCopy map[any]any, codemarkingconditiontypeFrom *CodeMarkingConditionType) (codemarkingconditiontypeTo *CodeMarkingConditionType) {
	var alreadyCopied bool
	codemarkingconditiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemarkingconditiontypeFrom)
	if alreadyCopied {
		return
	}
	codemarkingconditiontypeFrom.GongCopyBasicFields(codemarkingconditiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMarkingStyleType(mapOrigCopy map[any]any, codemarkingstyletypeFrom *CodeMarkingStyleType) (codemarkingstyletypeTo *CodeMarkingStyleType) {
	var alreadyCopied bool
	codemarkingstyletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemarkingstyletypeFrom)
	if alreadyCopied {
		return
	}
	codemarkingstyletypeFrom.GongCopyBasicFields(codemarkingstyletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMeteoConditionsType(mapOrigCopy map[any]any, codemeteoconditionstypeFrom *CodeMeteoConditionsType) (codemeteoconditionstypeTo *CodeMeteoConditionsType) {
	var alreadyCopied bool
	codemeteoconditionstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemeteoconditionstypeFrom)
	if alreadyCopied {
		return
	}
	codemeteoconditionstypeFrom.GongCopyBasicFields(codemeteoconditionstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMilitaryOperationsType(mapOrigCopy map[any]any, codemilitaryoperationstypeFrom *CodeMilitaryOperationsType) (codemilitaryoperationstypeTo *CodeMilitaryOperationsType) {
	var alreadyCopied bool
	codemilitaryoperationstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemilitaryoperationstypeFrom)
	if alreadyCopied {
		return
	}
	codemilitaryoperationstypeFrom.GongCopyBasicFields(codemilitaryoperationstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMilitaryRoutePointType(mapOrigCopy map[any]any, codemilitaryroutepointtypeFrom *CodeMilitaryRoutePointType) (codemilitaryroutepointtypeTo *CodeMilitaryRoutePointType) {
	var alreadyCopied bool
	codemilitaryroutepointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemilitaryroutepointtypeFrom)
	if alreadyCopied {
		return
	}
	codemilitaryroutepointtypeFrom.GongCopyBasicFields(codemilitaryroutepointtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMilitaryStatusType(mapOrigCopy map[any]any, codemilitarystatustypeFrom *CodeMilitaryStatusType) (codemilitarystatustypeTo *CodeMilitaryStatusType) {
	var alreadyCopied bool
	codemilitarystatustypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemilitarystatustypeFrom)
	if alreadyCopied {
		return
	}
	codemilitarystatustypeFrom.GongCopyBasicFields(codemilitarystatustypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMilitaryTrainingType(mapOrigCopy map[any]any, codemilitarytrainingtypeFrom *CodeMilitaryTrainingType) (codemilitarytrainingtypeTo *CodeMilitaryTrainingType) {
	var alreadyCopied bool
	codemilitarytrainingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemilitarytrainingtypeFrom)
	if alreadyCopied {
		return
	}
	codemilitarytrainingtypeFrom.GongCopyBasicFields(codemilitarytrainingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMinimaFinalApproachPathType(mapOrigCopy map[any]any, codeminimafinalapproachpathtypeFrom *CodeMinimaFinalApproachPathType) (codeminimafinalapproachpathtypeTo *CodeMinimaFinalApproachPathType) {
	var alreadyCopied bool
	codeminimafinalapproachpathtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeminimafinalapproachpathtypeFrom)
	if alreadyCopied {
		return
	}
	codeminimafinalapproachpathtypeFrom.GongCopyBasicFields(codeminimafinalapproachpathtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMinimumAltitudeType(mapOrigCopy map[any]any, codeminimumaltitudetypeFrom *CodeMinimumAltitudeType) (codeminimumaltitudetypeTo *CodeMinimumAltitudeType) {
	var alreadyCopied bool
	codeminimumaltitudetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeminimumaltitudetypeFrom)
	if alreadyCopied {
		return
	}
	codeminimumaltitudetypeFrom.GongCopyBasicFields(codeminimumaltitudetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMinimumHeightType(mapOrigCopy map[any]any, codeminimumheighttypeFrom *CodeMinimumHeightType) (codeminimumheighttypeTo *CodeMinimumHeightType) {
	var alreadyCopied bool
	codeminimumheighttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeminimumheighttypeFrom)
	if alreadyCopied {
		return
	}
	codeminimumheighttypeFrom.GongCopyBasicFields(codeminimumheighttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMissedApproachType(mapOrigCopy map[any]any, codemissedapproachtypeFrom *CodeMissedApproachType) (codemissedapproachtypeTo *CodeMissedApproachType) {
	var alreadyCopied bool
	codemissedapproachtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemissedapproachtypeFrom)
	if alreadyCopied {
		return
	}
	codemissedapproachtypeFrom.GongCopyBasicFields(codemissedapproachtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNDBUsageType(mapOrigCopy map[any]any, codendbusagetypeFrom *CodeNDBUsageType) (codendbusagetypeTo *CodeNDBUsageType) {
	var alreadyCopied bool
	codendbusagetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codendbusagetypeFrom)
	if alreadyCopied {
		return
	}
	codendbusagetypeFrom.GongCopyBasicFields(codendbusagetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavaidDesignatorType(mapOrigCopy map[any]any, codenavaiddesignatortypeFrom *CodeNavaidDesignatorType) (codenavaiddesignatortypeTo *CodeNavaidDesignatorType) {
	var alreadyCopied bool
	codenavaiddesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavaiddesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codenavaiddesignatortypeFrom.GongCopyBasicFields(codenavaiddesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavaidPurposeType(mapOrigCopy map[any]any, codenavaidpurposetypeFrom *CodeNavaidPurposeType) (codenavaidpurposetypeTo *CodeNavaidPurposeType) {
	var alreadyCopied bool
	codenavaidpurposetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavaidpurposetypeFrom)
	if alreadyCopied {
		return
	}
	codenavaidpurposetypeFrom.GongCopyBasicFields(codenavaidpurposetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavaidServiceType(mapOrigCopy map[any]any, codenavaidservicetypeFrom *CodeNavaidServiceType) (codenavaidservicetypeTo *CodeNavaidServiceType) {
	var alreadyCopied bool
	codenavaidservicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavaidservicetypeFrom)
	if alreadyCopied {
		return
	}
	codenavaidservicetypeFrom.GongCopyBasicFields(codenavaidservicetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavigationAreaRestrictionType(mapOrigCopy map[any]any, codenavigationarearestrictiontypeFrom *CodeNavigationAreaRestrictionType) (codenavigationarearestrictiontypeTo *CodeNavigationAreaRestrictionType) {
	var alreadyCopied bool
	codenavigationarearestrictiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavigationarearestrictiontypeFrom)
	if alreadyCopied {
		return
	}
	codenavigationarearestrictiontypeFrom.GongCopyBasicFields(codenavigationarearestrictiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavigationAreaType(mapOrigCopy map[any]any, codenavigationareatypeFrom *CodeNavigationAreaType) (codenavigationareatypeTo *CodeNavigationAreaType) {
	var alreadyCopied bool
	codenavigationareatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavigationareatypeFrom)
	if alreadyCopied {
		return
	}
	codenavigationareatypeFrom.GongCopyBasicFields(codenavigationareatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavigationEquipmentType(mapOrigCopy map[any]any, codenavigationequipmenttypeFrom *CodeNavigationEquipmentType) (codenavigationequipmenttypeTo *CodeNavigationEquipmentType) {
	var alreadyCopied bool
	codenavigationequipmenttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavigationequipmenttypeFrom)
	if alreadyCopied {
		return
	}
	codenavigationequipmenttypeFrom.GongCopyBasicFields(codenavigationequipmenttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavigationSpecificationType(mapOrigCopy map[any]any, codenavigationspecificationtypeFrom *CodeNavigationSpecificationType) (codenavigationspecificationtypeTo *CodeNavigationSpecificationType) {
	var alreadyCopied bool
	codenavigationspecificationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavigationspecificationtypeFrom)
	if alreadyCopied {
		return
	}
	codenavigationspecificationtypeFrom.GongCopyBasicFields(codenavigationspecificationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNitrogenType(mapOrigCopy map[any]any, codenitrogentypeFrom *CodeNitrogenType) (codenitrogentypeTo *CodeNitrogenType) {
	var alreadyCopied bool
	codenitrogentypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenitrogentypeFrom)
	if alreadyCopied {
		return
	}
	codenitrogentypeFrom.GongCopyBasicFields(codenitrogentypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNorthReferenceType(mapOrigCopy map[any]any, codenorthreferencetypeFrom *CodeNorthReferenceType) (codenorthreferencetypeTo *CodeNorthReferenceType) {
	var alreadyCopied bool
	codenorthreferencetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenorthreferencetypeFrom)
	if alreadyCopied {
		return
	}
	codenorthreferencetypeFrom.GongCopyBasicFields(codenorthreferencetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNotePurposeType(mapOrigCopy map[any]any, codenotepurposetypeFrom *CodeNotePurposeType) (codenotepurposetypeTo *CodeNotePurposeType) {
	var alreadyCopied bool
	codenotepurposetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenotepurposetypeFrom)
	if alreadyCopied {
		return
	}
	codenotepurposetypeFrom.GongCopyBasicFields(codenotepurposetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeObstacleAreaType(mapOrigCopy map[any]any, codeobstacleareatypeFrom *CodeObstacleAreaType) (codeobstacleareatypeTo *CodeObstacleAreaType) {
	var alreadyCopied bool
	codeobstacleareatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeobstacleareatypeFrom)
	if alreadyCopied {
		return
	}
	codeobstacleareatypeFrom.GongCopyBasicFields(codeobstacleareatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeObstacleAssessmentSurfaceType(mapOrigCopy map[any]any, codeobstacleassessmentsurfacetypeFrom *CodeObstacleAssessmentSurfaceType) (codeobstacleassessmentsurfacetypeTo *CodeObstacleAssessmentSurfaceType) {
	var alreadyCopied bool
	codeobstacleassessmentsurfacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeobstacleassessmentsurfacetypeFrom)
	if alreadyCopied {
		return
	}
	codeobstacleassessmentsurfacetypeFrom.GongCopyBasicFields(codeobstacleassessmentsurfacetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeObstructionIdSurfaceZoneType(mapOrigCopy map[any]any, codeobstructionidsurfacezonetypeFrom *CodeObstructionIdSurfaceZoneType) (codeobstructionidsurfacezonetypeTo *CodeObstructionIdSurfaceZoneType) {
	var alreadyCopied bool
	codeobstructionidsurfacezonetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeobstructionidsurfacezonetypeFrom)
	if alreadyCopied {
		return
	}
	codeobstructionidsurfacezonetypeFrom.GongCopyBasicFields(codeobstructionidsurfacezonetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeOilType(mapOrigCopy map[any]any, codeoiltypeFrom *CodeOilType) (codeoiltypeTo *CodeOilType) {
	var alreadyCopied bool
	codeoiltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeoiltypeFrom)
	if alreadyCopied {
		return
	}
	codeoiltypeFrom.GongCopyBasicFields(codeoiltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeOperationAirportHeliportType(mapOrigCopy map[any]any, codeoperationairportheliporttypeFrom *CodeOperationAirportHeliportType) (codeoperationairportheliporttypeTo *CodeOperationAirportHeliportType) {
	var alreadyCopied bool
	codeoperationairportheliporttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeoperationairportheliporttypeFrom)
	if alreadyCopied {
		return
	}
	codeoperationairportheliporttypeFrom.GongCopyBasicFields(codeoperationairportheliporttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeOperationManoeuvringAreaType(mapOrigCopy map[any]any, codeoperationmanoeuvringareatypeFrom *CodeOperationManoeuvringAreaType) (codeoperationmanoeuvringareatypeTo *CodeOperationManoeuvringAreaType) {
	var alreadyCopied bool
	codeoperationmanoeuvringareatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeoperationmanoeuvringareatypeFrom)
	if alreadyCopied {
		return
	}
	codeoperationmanoeuvringareatypeFrom.GongCopyBasicFields(codeoperationmanoeuvringareatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeOrganisationDesignatorType(mapOrigCopy map[any]any, codeorganisationdesignatortypeFrom *CodeOrganisationDesignatorType) (codeorganisationdesignatortypeTo *CodeOrganisationDesignatorType) {
	var alreadyCopied bool
	codeorganisationdesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeorganisationdesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codeorganisationdesignatortypeFrom.GongCopyBasicFields(codeorganisationdesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeOrganisationHierarchyType(mapOrigCopy map[any]any, codeorganisationhierarchytypeFrom *CodeOrganisationHierarchyType) (codeorganisationhierarchytypeTo *CodeOrganisationHierarchyType) {
	var alreadyCopied bool
	codeorganisationhierarchytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeorganisationhierarchytypeFrom)
	if alreadyCopied {
		return
	}
	codeorganisationhierarchytypeFrom.GongCopyBasicFields(codeorganisationhierarchytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeOrganisationType(mapOrigCopy map[any]any, codeorganisationtypeFrom *CodeOrganisationType) (codeorganisationtypeTo *CodeOrganisationType) {
	var alreadyCopied bool
	codeorganisationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeorganisationtypeFrom)
	if alreadyCopied {
		return
	}
	codeorganisationtypeFrom.GongCopyBasicFields(codeorganisationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeOxygenType(mapOrigCopy map[any]any, codeoxygentypeFrom *CodeOxygenType) (codeoxygentypeTo *CodeOxygenType) {
	var alreadyCopied bool
	codeoxygentypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeoxygentypeFrom)
	if alreadyCopied {
		return
	}
	codeoxygentypeFrom.GongCopyBasicFields(codeoxygentypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePARType(mapOrigCopy map[any]any, codepartypeFrom *CodePARType) (codepartypeTo *CodePARType) {
	var alreadyCopied bool
	codepartypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepartypeFrom)
	if alreadyCopied {
		return
	}
	codepartypeFrom.GongCopyBasicFields(codepartypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePCNMethodType(mapOrigCopy map[any]any, codepcnmethodtypeFrom *CodePCNMethodType) (codepcnmethodtypeTo *CodePCNMethodType) {
	var alreadyCopied bool
	codepcnmethodtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepcnmethodtypeFrom)
	if alreadyCopied {
		return
	}
	codepcnmethodtypeFrom.GongCopyBasicFields(codepcnmethodtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePCNPavementType(mapOrigCopy map[any]any, codepcnpavementtypeFrom *CodePCNPavementType) (codepcnpavementtypeTo *CodePCNPavementType) {
	var alreadyCopied bool
	codepcnpavementtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepcnpavementtypeFrom)
	if alreadyCopied {
		return
	}
	codepcnpavementtypeFrom.GongCopyBasicFields(codepcnpavementtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePCNSubgradeType(mapOrigCopy map[any]any, codepcnsubgradetypeFrom *CodePCNSubgradeType) (codepcnsubgradetypeTo *CodePCNSubgradeType) {
	var alreadyCopied bool
	codepcnsubgradetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepcnsubgradetypeFrom)
	if alreadyCopied {
		return
	}
	codepcnsubgradetypeFrom.GongCopyBasicFields(codepcnsubgradetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePCNTyrePressureType(mapOrigCopy map[any]any, codepcntyrepressuretypeFrom *CodePCNTyrePressureType) (codepcntyrepressuretypeTo *CodePCNTyrePressureType) {
	var alreadyCopied bool
	codepcntyrepressuretypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepcntyrepressuretypeFrom)
	if alreadyCopied {
		return
	}
	codepcntyrepressuretypeFrom.GongCopyBasicFields(codepcntyrepressuretypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePassengerServiceType(mapOrigCopy map[any]any, codepassengerservicetypeFrom *CodePassengerServiceType) (codepassengerservicetypeTo *CodePassengerServiceType) {
	var alreadyCopied bool
	codepassengerservicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepassengerservicetypeFrom)
	if alreadyCopied {
		return
	}
	codepassengerservicetypeFrom.GongCopyBasicFields(codepassengerservicetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePilotControlledLightingType(mapOrigCopy map[any]any, codepilotcontrolledlightingtypeFrom *CodePilotControlledLightingType) (codepilotcontrolledlightingtypeTo *CodePilotControlledLightingType) {
	var alreadyCopied bool
	codepilotcontrolledlightingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepilotcontrolledlightingtypeFrom)
	if alreadyCopied {
		return
	}
	codepilotcontrolledlightingtypeFrom.GongCopyBasicFields(codepilotcontrolledlightingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePositionInILSType(mapOrigCopy map[any]any, codepositioninilstypeFrom *CodePositionInILSType) (codepositioninilstypeTo *CodePositionInILSType) {
	var alreadyCopied bool
	codepositioninilstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepositioninilstypeFrom)
	if alreadyCopied {
		return
	}
	codepositioninilstypeFrom.GongCopyBasicFields(codepositioninilstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePrimaryRadarType(mapOrigCopy map[any]any, codeprimaryradartypeFrom *CodePrimaryRadarType) (codeprimaryradartypeTo *CodePrimaryRadarType) {
	var alreadyCopied bool
	codeprimaryradartypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeprimaryradartypeFrom)
	if alreadyCopied {
		return
	}
	codeprimaryradartypeFrom.GongCopyBasicFields(codeprimaryradartypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeProcedureAvailabilityType(mapOrigCopy map[any]any, codeprocedureavailabilitytypeFrom *CodeProcedureAvailabilityType) (codeprocedureavailabilitytypeTo *CodeProcedureAvailabilityType) {
	var alreadyCopied bool
	codeprocedureavailabilitytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeprocedureavailabilitytypeFrom)
	if alreadyCopied {
		return
	}
	codeprocedureavailabilitytypeFrom.GongCopyBasicFields(codeprocedureavailabilitytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeProcedureCodingStandardType(mapOrigCopy map[any]any, codeprocedurecodingstandardtypeFrom *CodeProcedureCodingStandardType) (codeprocedurecodingstandardtypeTo *CodeProcedureCodingStandardType) {
	var alreadyCopied bool
	codeprocedurecodingstandardtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeprocedurecodingstandardtypeFrom)
	if alreadyCopied {
		return
	}
	codeprocedurecodingstandardtypeFrom.GongCopyBasicFields(codeprocedurecodingstandardtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeProcedureDistanceType(mapOrigCopy map[any]any, codeproceduredistancetypeFrom *CodeProcedureDistanceType) (codeproceduredistancetypeTo *CodeProcedureDistanceType) {
	var alreadyCopied bool
	codeproceduredistancetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeproceduredistancetypeFrom)
	if alreadyCopied {
		return
	}
	codeproceduredistancetypeFrom.GongCopyBasicFields(codeproceduredistancetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeProcedureFixRoleType(mapOrigCopy map[any]any, codeprocedurefixroletypeFrom *CodeProcedureFixRoleType) (codeprocedurefixroletypeTo *CodeProcedureFixRoleType) {
	var alreadyCopied bool
	codeprocedurefixroletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeprocedurefixroletypeFrom)
	if alreadyCopied {
		return
	}
	codeprocedurefixroletypeFrom.GongCopyBasicFields(codeprocedurefixroletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeProcedurePhaseType(mapOrigCopy map[any]any, codeprocedurephasetypeFrom *CodeProcedurePhaseType) (codeprocedurephasetypeTo *CodeProcedurePhaseType) {
	var alreadyCopied bool
	codeprocedurephasetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeprocedurephasetypeFrom)
	if alreadyCopied {
		return
	}
	codeprocedurephasetypeFrom.GongCopyBasicFields(codeprocedurephasetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeProtectAreaSectionType(mapOrigCopy map[any]any, codeprotectareasectiontypeFrom *CodeProtectAreaSectionType) (codeprotectareasectiontypeTo *CodeProtectAreaSectionType) {
	var alreadyCopied bool
	codeprotectareasectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeprotectareasectiontypeFrom)
	if alreadyCopied {
		return
	}
	codeprotectareasectiontypeFrom.GongCopyBasicFields(codeprotectareasectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRNPType(mapOrigCopy map[any]any, codernptypeFrom *CodeRNPType) (codernptypeTo *CodeRNPType) {
	var alreadyCopied bool
	codernptypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codernptypeFrom)
	if alreadyCopied {
		return
	}
	codernptypeFrom.GongCopyBasicFields(codernptypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRVRReadingType(mapOrigCopy map[any]any, codervrreadingtypeFrom *CodeRVRReadingType) (codervrreadingtypeTo *CodeRVRReadingType) {
	var alreadyCopied bool
	codervrreadingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codervrreadingtypeFrom)
	if alreadyCopied {
		return
	}
	codervrreadingtypeFrom.GongCopyBasicFields(codervrreadingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRVSMPointRoleType(mapOrigCopy map[any]any, codervsmpointroletypeFrom *CodeRVSMPointRoleType) (codervsmpointroletypeTo *CodeRVSMPointRoleType) {
	var alreadyCopied bool
	codervsmpointroletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codervsmpointroletypeFrom)
	if alreadyCopied {
		return
	}
	codervsmpointroletypeFrom.GongCopyBasicFields(codervsmpointroletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRVSMType(mapOrigCopy map[any]any, codervsmtypeFrom *CodeRVSMType) (codervsmtypeTo *CodeRVSMType) {
	var alreadyCopied bool
	codervsmtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codervsmtypeFrom)
	if alreadyCopied {
		return
	}
	codervsmtypeFrom.GongCopyBasicFields(codervsmtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRadarServiceType(mapOrigCopy map[any]any, coderadarservicetypeFrom *CodeRadarServiceType) (coderadarservicetypeTo *CodeRadarServiceType) {
	var alreadyCopied bool
	coderadarservicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderadarservicetypeFrom)
	if alreadyCopied {
		return
	}
	coderadarservicetypeFrom.GongCopyBasicFields(coderadarservicetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRadioEmissionType(mapOrigCopy map[any]any, coderadioemissiontypeFrom *CodeRadioEmissionType) (coderadioemissiontypeTo *CodeRadioEmissionType) {
	var alreadyCopied bool
	coderadioemissiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderadioemissiontypeFrom)
	if alreadyCopied {
		return
	}
	coderadioemissiontypeFrom.GongCopyBasicFields(coderadioemissiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRadioFrequencyAreaType(mapOrigCopy map[any]any, coderadiofrequencyareatypeFrom *CodeRadioFrequencyAreaType) (coderadiofrequencyareatypeTo *CodeRadioFrequencyAreaType) {
	var alreadyCopied bool
	coderadiofrequencyareatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderadiofrequencyareatypeFrom)
	if alreadyCopied {
		return
	}
	coderadiofrequencyareatypeFrom.GongCopyBasicFields(coderadiofrequencyareatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRadioSignalType(mapOrigCopy map[any]any, coderadiosignaltypeFrom *CodeRadioSignalType) (coderadiosignaltypeTo *CodeRadioSignalType) {
	var alreadyCopied bool
	coderadiosignaltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderadiosignaltypeFrom)
	if alreadyCopied {
		return
	}
	coderadiosignaltypeFrom.GongCopyBasicFields(coderadiosignaltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeReferencePathIdentifierType(mapOrigCopy map[any]any, codereferencepathidentifiertypeFrom *CodeReferencePathIdentifierType) (codereferencepathidentifiertypeTo *CodeReferencePathIdentifierType) {
	var alreadyCopied bool
	codereferencepathidentifiertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codereferencepathidentifiertypeFrom)
	if alreadyCopied {
		return
	}
	codereferencepathidentifiertypeFrom.GongCopyBasicFields(codereferencepathidentifiertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeReferenceRoleType(mapOrigCopy map[any]any, codereferenceroletypeFrom *CodeReferenceRoleType) (codereferenceroletypeTo *CodeReferenceRoleType) {
	var alreadyCopied bool
	codereferenceroletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codereferenceroletypeFrom)
	if alreadyCopied {
		return
	}
	codereferenceroletypeFrom.GongCopyBasicFields(codereferenceroletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeReflectorType(mapOrigCopy map[any]any, codereflectortypeFrom *CodeReflectorType) (codereflectortypeTo *CodeReflectorType) {
	var alreadyCopied bool
	codereflectortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codereflectortypeFrom)
	if alreadyCopied {
		return
	}
	codereflectortypeFrom.GongCopyBasicFields(codereflectortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRelativePositionType(mapOrigCopy map[any]any, coderelativepositiontypeFrom *CodeRelativePositionType) (coderelativepositiontypeTo *CodeRelativePositionType) {
	var alreadyCopied bool
	coderelativepositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderelativepositiontypeFrom)
	if alreadyCopied {
		return
	}
	coderelativepositiontypeFrom.GongCopyBasicFields(coderelativepositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRoadType(mapOrigCopy map[any]any, coderoadtypeFrom *CodeRoadType) (coderoadtypeTo *CodeRoadType) {
	var alreadyCopied bool
	coderoadtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderoadtypeFrom)
	if alreadyCopied {
		return
	}
	coderoadtypeFrom.GongCopyBasicFields(coderoadtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteAvailabilityType(mapOrigCopy map[any]any, coderouteavailabilitytypeFrom *CodeRouteAvailabilityType) (coderouteavailabilitytypeTo *CodeRouteAvailabilityType) {
	var alreadyCopied bool
	coderouteavailabilitytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderouteavailabilitytypeFrom)
	if alreadyCopied {
		return
	}
	coderouteavailabilitytypeFrom.GongCopyBasicFields(coderouteavailabilitytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteDesignatorLetterType(mapOrigCopy map[any]any, coderoutedesignatorlettertypeFrom *CodeRouteDesignatorLetterType) (coderoutedesignatorlettertypeTo *CodeRouteDesignatorLetterType) {
	var alreadyCopied bool
	coderoutedesignatorlettertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderoutedesignatorlettertypeFrom)
	if alreadyCopied {
		return
	}
	coderoutedesignatorlettertypeFrom.GongCopyBasicFields(coderoutedesignatorlettertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteDesignatorPrefixType(mapOrigCopy map[any]any, coderoutedesignatorprefixtypeFrom *CodeRouteDesignatorPrefixType) (coderoutedesignatorprefixtypeTo *CodeRouteDesignatorPrefixType) {
	var alreadyCopied bool
	coderoutedesignatorprefixtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderoutedesignatorprefixtypeFrom)
	if alreadyCopied {
		return
	}
	coderoutedesignatorprefixtypeFrom.GongCopyBasicFields(coderoutedesignatorprefixtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteDesignatorSuffixType(mapOrigCopy map[any]any, coderoutedesignatorsuffixtypeFrom *CodeRouteDesignatorSuffixType) (coderoutedesignatorsuffixtypeTo *CodeRouteDesignatorSuffixType) {
	var alreadyCopied bool
	coderoutedesignatorsuffixtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderoutedesignatorsuffixtypeFrom)
	if alreadyCopied {
		return
	}
	coderoutedesignatorsuffixtypeFrom.GongCopyBasicFields(coderoutedesignatorsuffixtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteIndicatorType(mapOrigCopy map[any]any, coderouteindicatortypeFrom *CodeRouteIndicatorType) (coderouteindicatortypeTo *CodeRouteIndicatorType) {
	var alreadyCopied bool
	coderouteindicatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderouteindicatortypeFrom)
	if alreadyCopied {
		return
	}
	coderouteindicatortypeFrom.GongCopyBasicFields(coderouteindicatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteNavigationType(mapOrigCopy map[any]any, coderoutenavigationtypeFrom *CodeRouteNavigationType) (coderoutenavigationtypeTo *CodeRouteNavigationType) {
	var alreadyCopied bool
	coderoutenavigationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderoutenavigationtypeFrom)
	if alreadyCopied {
		return
	}
	coderoutenavigationtypeFrom.GongCopyBasicFields(coderoutenavigationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteOriginType(mapOrigCopy map[any]any, coderouteorigintypeFrom *CodeRouteOriginType) (coderouteorigintypeTo *CodeRouteOriginType) {
	var alreadyCopied bool
	coderouteorigintypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderouteorigintypeFrom)
	if alreadyCopied {
		return
	}
	coderouteorigintypeFrom.GongCopyBasicFields(coderouteorigintypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteSegmentPathType(mapOrigCopy map[any]any, coderoutesegmentpathtypeFrom *CodeRouteSegmentPathType) (coderoutesegmentpathtypeTo *CodeRouteSegmentPathType) {
	var alreadyCopied bool
	coderoutesegmentpathtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderoutesegmentpathtypeFrom)
	if alreadyCopied {
		return
	}
	coderoutesegmentpathtypeFrom.GongCopyBasicFields(coderoutesegmentpathtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRouteType(mapOrigCopy map[any]any, coderoutetypeFrom *CodeRouteType) (coderoutetypeTo *CodeRouteType) {
	var alreadyCopied bool
	coderoutetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderoutetypeFrom)
	if alreadyCopied {
		return
	}
	coderoutetypeFrom.GongCopyBasicFields(coderoutetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRuleProcedureTitleType(mapOrigCopy map[any]any, coderuleproceduretitletypeFrom *CodeRuleProcedureTitleType) (coderuleproceduretitletypeTo *CodeRuleProcedureTitleType) {
	var alreadyCopied bool
	coderuleproceduretitletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderuleproceduretitletypeFrom)
	if alreadyCopied {
		return
	}
	coderuleproceduretitletypeFrom.GongCopyBasicFields(coderuleproceduretitletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRuleProcedureType(mapOrigCopy map[any]any, coderuleproceduretypeFrom *CodeRuleProcedureType) (coderuleproceduretypeTo *CodeRuleProcedureType) {
	var alreadyCopied bool
	coderuleproceduretypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderuleproceduretypeFrom)
	if alreadyCopied {
		return
	}
	coderuleproceduretypeFrom.GongCopyBasicFields(coderuleproceduretypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRunwayElementType(mapOrigCopy map[any]any, coderunwayelementtypeFrom *CodeRunwayElementType) (coderunwayelementtypeTo *CodeRunwayElementType) {
	var alreadyCopied bool
	coderunwayelementtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderunwayelementtypeFrom)
	if alreadyCopied {
		return
	}
	coderunwayelementtypeFrom.GongCopyBasicFields(coderunwayelementtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRunwayMarkingType(mapOrigCopy map[any]any, coderunwaymarkingtypeFrom *CodeRunwayMarkingType) (coderunwaymarkingtypeTo *CodeRunwayMarkingType) {
	var alreadyCopied bool
	coderunwaymarkingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderunwaymarkingtypeFrom)
	if alreadyCopied {
		return
	}
	coderunwaymarkingtypeFrom.GongCopyBasicFields(coderunwaymarkingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRunwayPointRoleType(mapOrigCopy map[any]any, coderunwaypointroletypeFrom *CodeRunwayPointRoleType) (coderunwaypointroletypeTo *CodeRunwayPointRoleType) {
	var alreadyCopied bool
	coderunwaypointroletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderunwaypointroletypeFrom)
	if alreadyCopied {
		return
	}
	coderunwaypointroletypeFrom.GongCopyBasicFields(coderunwaypointroletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRunwayProtectionAreaType(mapOrigCopy map[any]any, coderunwayprotectionareatypeFrom *CodeRunwayProtectionAreaType) (coderunwayprotectionareatypeTo *CodeRunwayProtectionAreaType) {
	var alreadyCopied bool
	coderunwayprotectionareatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderunwayprotectionareatypeFrom)
	if alreadyCopied {
		return
	}
	coderunwayprotectionareatypeFrom.GongCopyBasicFields(coderunwayprotectionareatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRunwaySectionType(mapOrigCopy map[any]any, coderunwaysectiontypeFrom *CodeRunwaySectionType) (coderunwaysectiontypeTo *CodeRunwaySectionType) {
	var alreadyCopied bool
	coderunwaysectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderunwaysectiontypeFrom)
	if alreadyCopied {
		return
	}
	coderunwaysectiontypeFrom.GongCopyBasicFields(coderunwaysectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRunwayType(mapOrigCopy map[any]any, coderunwaytypeFrom *CodeRunwayType) (coderunwaytypeTo *CodeRunwayType) {
	var alreadyCopied bool
	coderunwaytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderunwaytypeFrom)
	if alreadyCopied {
		return
	}
	coderunwaytypeFrom.GongCopyBasicFields(coderunwaytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSafeAltitudeType(mapOrigCopy map[any]any, codesafealtitudetypeFrom *CodeSafeAltitudeType) (codesafealtitudetypeTo *CodeSafeAltitudeType) {
	var alreadyCopied bool
	codesafealtitudetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesafealtitudetypeFrom)
	if alreadyCopied {
		return
	}
	codesafealtitudetypeFrom.GongCopyBasicFields(codesafealtitudetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSegmentPathType(mapOrigCopy map[any]any, codesegmentpathtypeFrom *CodeSegmentPathType) (codesegmentpathtypeTo *CodeSegmentPathType) {
	var alreadyCopied bool
	codesegmentpathtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesegmentpathtypeFrom)
	if alreadyCopied {
		return
	}
	codesegmentpathtypeFrom.GongCopyBasicFields(codesegmentpathtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSegmentTerminationType(mapOrigCopy map[any]any, codesegmentterminationtypeFrom *CodeSegmentTerminationType) (codesegmentterminationtypeTo *CodeSegmentTerminationType) {
	var alreadyCopied bool
	codesegmentterminationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesegmentterminationtypeFrom)
	if alreadyCopied {
		return
	}
	codesegmentterminationtypeFrom.GongCopyBasicFields(codesegmentterminationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeServiceATCType(mapOrigCopy map[any]any, codeserviceatctypeFrom *CodeServiceATCType) (codeserviceatctypeTo *CodeServiceATCType) {
	var alreadyCopied bool
	codeserviceatctypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeserviceatctypeFrom)
	if alreadyCopied {
		return
	}
	codeserviceatctypeFrom.GongCopyBasicFields(codeserviceatctypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeServiceATFMType(mapOrigCopy map[any]any, codeserviceatfmtypeFrom *CodeServiceATFMType) (codeserviceatfmtypeTo *CodeServiceATFMType) {
	var alreadyCopied bool
	codeserviceatfmtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeserviceatfmtypeFrom)
	if alreadyCopied {
		return
	}
	codeserviceatfmtypeFrom.GongCopyBasicFields(codeserviceatfmtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeServiceGroundControlType(mapOrigCopy map[any]any, codeservicegroundcontroltypeFrom *CodeServiceGroundControlType) (codeservicegroundcontroltypeTo *CodeServiceGroundControlType) {
	var alreadyCopied bool
	codeservicegroundcontroltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeservicegroundcontroltypeFrom)
	if alreadyCopied {
		return
	}
	codeservicegroundcontroltypeFrom.GongCopyBasicFields(codeservicegroundcontroltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeServiceInformationType(mapOrigCopy map[any]any, codeserviceinformationtypeFrom *CodeServiceInformationType) (codeserviceinformationtypeTo *CodeServiceInformationType) {
	var alreadyCopied bool
	codeserviceinformationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeserviceinformationtypeFrom)
	if alreadyCopied {
		return
	}
	codeserviceinformationtypeFrom.GongCopyBasicFields(codeserviceinformationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeServiceSARType(mapOrigCopy map[any]any, codeservicesartypeFrom *CodeServiceSARType) (codeservicesartypeTo *CodeServiceSARType) {
	var alreadyCopied bool
	codeservicesartypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeservicesartypeFrom)
	if alreadyCopied {
		return
	}
	codeservicesartypeFrom.GongCopyBasicFields(codeservicesartypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSideType(mapOrigCopy map[any]any, codesidetypeFrom *CodeSideType) (codesidetypeTo *CodeSideType) {
	var alreadyCopied bool
	codesidetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesidetypeFrom)
	if alreadyCopied {
		return
	}
	codesidetypeFrom.GongCopyBasicFields(codesidetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSignalPerformanceILSType(mapOrigCopy map[any]any, codesignalperformanceilstypeFrom *CodeSignalPerformanceILSType) (codesignalperformanceilstypeTo *CodeSignalPerformanceILSType) {
	var alreadyCopied bool
	codesignalperformanceilstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesignalperformanceilstypeFrom)
	if alreadyCopied {
		return
	}
	codesignalperformanceilstypeFrom.GongCopyBasicFields(codesignalperformanceilstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSpecialDateType(mapOrigCopy map[any]any, codespecialdatetypeFrom *CodeSpecialDateType) (codespecialdatetypeTo *CodeSpecialDateType) {
	var alreadyCopied bool
	codespecialdatetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codespecialdatetypeFrom)
	if alreadyCopied {
		return
	}
	codespecialdatetypeFrom.GongCopyBasicFields(codespecialdatetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSpecialNavigationChainDesignatorType(mapOrigCopy map[any]any, codespecialnavigationchaindesignatortypeFrom *CodeSpecialNavigationChainDesignatorType) (codespecialnavigationchaindesignatortypeTo *CodeSpecialNavigationChainDesignatorType) {
	var alreadyCopied bool
	codespecialnavigationchaindesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codespecialnavigationchaindesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codespecialnavigationchaindesignatortypeFrom.GongCopyBasicFields(codespecialnavigationchaindesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSpecialNavigationStationType(mapOrigCopy map[any]any, codespecialnavigationstationtypeFrom *CodeSpecialNavigationStationType) (codespecialnavigationstationtypeTo *CodeSpecialNavigationStationType) {
	var alreadyCopied bool
	codespecialnavigationstationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codespecialnavigationstationtypeFrom)
	if alreadyCopied {
		return
	}
	codespecialnavigationstationtypeFrom.GongCopyBasicFields(codespecialnavigationstationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSpecialNavigationSystemType(mapOrigCopy map[any]any, codespecialnavigationsystemtypeFrom *CodeSpecialNavigationSystemType) (codespecialnavigationsystemtypeTo *CodeSpecialNavigationSystemType) {
	var alreadyCopied bool
	codespecialnavigationsystemtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codespecialnavigationsystemtypeFrom)
	if alreadyCopied {
		return
	}
	codespecialnavigationsystemtypeFrom.GongCopyBasicFields(codespecialnavigationsystemtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSpeedReferenceType(mapOrigCopy map[any]any, codespeedreferencetypeFrom *CodeSpeedReferenceType) (codespeedreferencetypeTo *CodeSpeedReferenceType) {
	var alreadyCopied bool
	codespeedreferencetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codespeedreferencetypeFrom)
	if alreadyCopied {
		return
	}
	codespeedreferencetypeFrom.GongCopyBasicFields(codespeedreferencetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStandbyPowerType(mapOrigCopy map[any]any, codestandbypowertypeFrom *CodeStandbyPowerType) (codestandbypowertypeTo *CodeStandbyPowerType) {
	var alreadyCopied bool
	codestandbypowertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestandbypowertypeFrom)
	if alreadyCopied {
		return
	}
	codestandbypowertypeFrom.GongCopyBasicFields(codestandbypowertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStatusAirportType(mapOrigCopy map[any]any, codestatusairporttypeFrom *CodeStatusAirportType) (codestatusairporttypeTo *CodeStatusAirportType) {
	var alreadyCopied bool
	codestatusairporttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusairporttypeFrom)
	if alreadyCopied {
		return
	}
	codestatusairporttypeFrom.GongCopyBasicFields(codestatusairporttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStatusAirspaceType(mapOrigCopy map[any]any, codestatusairspacetypeFrom *CodeStatusAirspaceType) (codestatusairspacetypeTo *CodeStatusAirspaceType) {
	var alreadyCopied bool
	codestatusairspacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusairspacetypeFrom)
	if alreadyCopied {
		return
	}
	codestatusairspacetypeFrom.GongCopyBasicFields(codestatusairspacetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStatusConstructionType(mapOrigCopy map[any]any, codestatusconstructiontypeFrom *CodeStatusConstructionType) (codestatusconstructiontypeTo *CodeStatusConstructionType) {
	var alreadyCopied bool
	codestatusconstructiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusconstructiontypeFrom)
	if alreadyCopied {
		return
	}
	codestatusconstructiontypeFrom.GongCopyBasicFields(codestatusconstructiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStatusNavaidType(mapOrigCopy map[any]any, codestatusnavaidtypeFrom *CodeStatusNavaidType) (codestatusnavaidtypeTo *CodeStatusNavaidType) {
	var alreadyCopied bool
	codestatusnavaidtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusnavaidtypeFrom)
	if alreadyCopied {
		return
	}
	codestatusnavaidtypeFrom.GongCopyBasicFields(codestatusnavaidtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStatusOperationsType(mapOrigCopy map[any]any, codestatusoperationstypeFrom *CodeStatusOperationsType) (codestatusoperationstypeTo *CodeStatusOperationsType) {
	var alreadyCopied bool
	codestatusoperationstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusoperationstypeFrom)
	if alreadyCopied {
		return
	}
	codestatusoperationstypeFrom.GongCopyBasicFields(codestatusoperationstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStatusServiceType(mapOrigCopy map[any]any, codestatusservicetypeFrom *CodeStatusServiceType) (codestatusservicetypeTo *CodeStatusServiceType) {
	var alreadyCopied bool
	codestatusservicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusservicetypeFrom)
	if alreadyCopied {
		return
	}
	codestatusservicetypeFrom.GongCopyBasicFields(codestatusservicetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSurfaceCompositionType(mapOrigCopy map[any]any, codesurfacecompositiontypeFrom *CodeSurfaceCompositionType) (codesurfacecompositiontypeTo *CodeSurfaceCompositionType) {
	var alreadyCopied bool
	codesurfacecompositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesurfacecompositiontypeFrom)
	if alreadyCopied {
		return
	}
	codesurfacecompositiontypeFrom.GongCopyBasicFields(codesurfacecompositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSurfaceConditionType(mapOrigCopy map[any]any, codesurfaceconditiontypeFrom *CodeSurfaceConditionType) (codesurfaceconditiontypeTo *CodeSurfaceConditionType) {
	var alreadyCopied bool
	codesurfaceconditiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesurfaceconditiontypeFrom)
	if alreadyCopied {
		return
	}
	codesurfaceconditiontypeFrom.GongCopyBasicFields(codesurfaceconditiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSurfacePreparationType(mapOrigCopy map[any]any, codesurfacepreparationtypeFrom *CodeSurfacePreparationType) (codesurfacepreparationtypeTo *CodeSurfacePreparationType) {
	var alreadyCopied bool
	codesurfacepreparationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesurfacepreparationtypeFrom)
	if alreadyCopied {
		return
	}
	codesurfacepreparationtypeFrom.GongCopyBasicFields(codesurfacepreparationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSystemActivationType(mapOrigCopy map[any]any, codesystemactivationtypeFrom *CodeSystemActivationType) (codesystemactivationtypeTo *CodeSystemActivationType) {
	var alreadyCopied bool
	codesystemactivationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesystemactivationtypeFrom)
	if alreadyCopied {
		return
	}
	codesystemactivationtypeFrom.GongCopyBasicFields(codesystemactivationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTAAType(mapOrigCopy map[any]any, codetaatypeFrom *CodeTAAType) (codetaatypeTo *CodeTAAType) {
	var alreadyCopied bool
	codetaatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetaatypeFrom)
	if alreadyCopied {
		return
	}
	codetaatypeFrom.GongCopyBasicFields(codetaatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTACANChannelType(mapOrigCopy map[any]any, codetacanchanneltypeFrom *CodeTACANChannelType) (codetacanchanneltypeTo *CodeTACANChannelType) {
	var alreadyCopied bool
	codetacanchanneltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetacanchanneltypeFrom)
	if alreadyCopied {
		return
	}
	codetacanchanneltypeFrom.GongCopyBasicFields(codetacanchanneltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTLOFSectionType(mapOrigCopy map[any]any, codetlofsectiontypeFrom *CodeTLOFSectionType) (codetlofsectiontypeTo *CodeTLOFSectionType) {
	var alreadyCopied bool
	codetlofsectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetlofsectiontypeFrom)
	if alreadyCopied {
		return
	}
	codetlofsectiontypeFrom.GongCopyBasicFields(codetlofsectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTaxiwayElementType(mapOrigCopy map[any]any, codetaxiwayelementtypeFrom *CodeTaxiwayElementType) (codetaxiwayelementtypeTo *CodeTaxiwayElementType) {
	var alreadyCopied bool
	codetaxiwayelementtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetaxiwayelementtypeFrom)
	if alreadyCopied {
		return
	}
	codetaxiwayelementtypeFrom.GongCopyBasicFields(codetaxiwayelementtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTaxiwaySectionType(mapOrigCopy map[any]any, codetaxiwaysectiontypeFrom *CodeTaxiwaySectionType) (codetaxiwaysectiontypeTo *CodeTaxiwaySectionType) {
	var alreadyCopied bool
	codetaxiwaysectiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetaxiwaysectiontypeFrom)
	if alreadyCopied {
		return
	}
	codetaxiwaysectiontypeFrom.GongCopyBasicFields(codetaxiwaysectiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTaxiwayType(mapOrigCopy map[any]any, codetaxiwaytypeFrom *CodeTaxiwayType) (codetaxiwaytypeTo *CodeTaxiwayType) {
	var alreadyCopied bool
	codetaxiwaytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetaxiwaytypeFrom)
	if alreadyCopied {
		return
	}
	codetaxiwaytypeFrom.GongCopyBasicFields(codetaxiwaytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTelecomNetworkType(mapOrigCopy map[any]any, codetelecomnetworktypeFrom *CodeTelecomNetworkType) (codetelecomnetworktypeTo *CodeTelecomNetworkType) {
	var alreadyCopied bool
	codetelecomnetworktypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetelecomnetworktypeFrom)
	if alreadyCopied {
		return
	}
	codetelecomnetworktypeFrom.GongCopyBasicFields(codetelecomnetworktypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTimeEventCombinationType(mapOrigCopy map[any]any, codetimeeventcombinationtypeFrom *CodeTimeEventCombinationType) (codetimeeventcombinationtypeTo *CodeTimeEventCombinationType) {
	var alreadyCopied bool
	codetimeeventcombinationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetimeeventcombinationtypeFrom)
	if alreadyCopied {
		return
	}
	codetimeeventcombinationtypeFrom.GongCopyBasicFields(codetimeeventcombinationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTimeEventType(mapOrigCopy map[any]any, codetimeeventtypeFrom *CodeTimeEventType) (codetimeeventtypeTo *CodeTimeEventType) {
	var alreadyCopied bool
	codetimeeventtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetimeeventtypeFrom)
	if alreadyCopied {
		return
	}
	codetimeeventtypeFrom.GongCopyBasicFields(codetimeeventtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTimeReferenceType(mapOrigCopy map[any]any, codetimereferencetypeFrom *CodeTimeReferenceType) (codetimereferencetypeTo *CodeTimeReferenceType) {
	var alreadyCopied bool
	codetimereferencetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetimereferencetypeFrom)
	if alreadyCopied {
		return
	}
	codetimereferencetypeFrom.GongCopyBasicFields(codetimereferencetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTrajectoryType(mapOrigCopy map[any]any, codetrajectorytypeFrom *CodeTrajectoryType) (codetrajectorytypeTo *CodeTrajectoryType) {
	var alreadyCopied bool
	codetrajectorytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetrajectorytypeFrom)
	if alreadyCopied {
		return
	}
	codetrajectorytypeFrom.GongCopyBasicFields(codetrajectorytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeTransponderType(mapOrigCopy map[any]any, codetranspondertypeFrom *CodeTransponderType) (codetranspondertypeTo *CodeTransponderType) {
	var alreadyCopied bool
	codetranspondertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetranspondertypeFrom)
	if alreadyCopied {
		return
	}
	codetranspondertypeFrom.GongCopyBasicFields(codetranspondertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeUUIDType(mapOrigCopy map[any]any, codeuuidtypeFrom *CodeUUIDType) (codeuuidtypeTo *CodeUUIDType) {
	var alreadyCopied bool
	codeuuidtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeuuidtypeFrom)
	if alreadyCopied {
		return
	}
	codeuuidtypeFrom.GongCopyBasicFields(codeuuidtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeUnitDependencyType(mapOrigCopy map[any]any, codeunitdependencytypeFrom *CodeUnitDependencyType) (codeunitdependencytypeTo *CodeUnitDependencyType) {
	var alreadyCopied bool
	codeunitdependencytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeunitdependencytypeFrom)
	if alreadyCopied {
		return
	}
	codeunitdependencytypeFrom.GongCopyBasicFields(codeunitdependencytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeUnitType(mapOrigCopy map[any]any, codeunittypeFrom *CodeUnitType) (codeunittypeTo *CodeUnitType) {
	var alreadyCopied bool
	codeunittypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeunittypeFrom)
	if alreadyCopied {
		return
	}
	codeunittypeFrom.GongCopyBasicFields(codeunittypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeUpperAlphaType(mapOrigCopy map[any]any, codeupperalphatypeFrom *CodeUpperAlphaType) (codeupperalphatypeTo *CodeUpperAlphaType) {
	var alreadyCopied bool
	codeupperalphatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeupperalphatypeFrom)
	if alreadyCopied {
		return
	}
	codeupperalphatypeFrom.GongCopyBasicFields(codeupperalphatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeUsageLimitationType(mapOrigCopy map[any]any, codeusagelimitationtypeFrom *CodeUsageLimitationType) (codeusagelimitationtypeTo *CodeUsageLimitationType) {
	var alreadyCopied bool
	codeusagelimitationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeusagelimitationtypeFrom)
	if alreadyCopied {
		return
	}
	codeusagelimitationtypeFrom.GongCopyBasicFields(codeusagelimitationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVASISType(mapOrigCopy map[any]any, codevasistypeFrom *CodeVASISType) (codevasistypeTo *CodeVASISType) {
	var alreadyCopied bool
	codevasistypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codevasistypeFrom)
	if alreadyCopied {
		return
	}
	codevasistypeFrom.GongCopyBasicFields(codevasistypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVORType(mapOrigCopy map[any]any, codevortypeFrom *CodeVORType) (codevortypeTo *CodeVORType) {
	var alreadyCopied bool
	codevortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codevortypeFrom)
	if alreadyCopied {
		return
	}
	codevortypeFrom.GongCopyBasicFields(codevortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeValueInterpretationType(mapOrigCopy map[any]any, codevalueinterpretationtypeFrom *CodeValueInterpretationType) (codevalueinterpretationtypeTo *CodeValueInterpretationType) {
	var alreadyCopied bool
	codevalueinterpretationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codevalueinterpretationtypeFrom)
	if alreadyCopied {
		return
	}
	codevalueinterpretationtypeFrom.GongCopyBasicFields(codevalueinterpretationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVerticalDatumType(mapOrigCopy map[any]any, codeverticaldatumtypeFrom *CodeVerticalDatumType) (codeverticaldatumtypeTo *CodeVerticalDatumType) {
	var alreadyCopied bool
	codeverticaldatumtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeverticaldatumtypeFrom)
	if alreadyCopied {
		return
	}
	codeverticaldatumtypeFrom.GongCopyBasicFields(codeverticaldatumtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVerticalReferenceType(mapOrigCopy map[any]any, codeverticalreferencetypeFrom *CodeVerticalReferenceType) (codeverticalreferencetypeTo *CodeVerticalReferenceType) {
	var alreadyCopied bool
	codeverticalreferencetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeverticalreferencetypeFrom)
	if alreadyCopied {
		return
	}
	codeverticalreferencetypeFrom.GongCopyBasicFields(codeverticalreferencetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVerticalStructureMarkingType(mapOrigCopy map[any]any, codeverticalstructuremarkingtypeFrom *CodeVerticalStructureMarkingType) (codeverticalstructuremarkingtypeTo *CodeVerticalStructureMarkingType) {
	var alreadyCopied bool
	codeverticalstructuremarkingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeverticalstructuremarkingtypeFrom)
	if alreadyCopied {
		return
	}
	codeverticalstructuremarkingtypeFrom.GongCopyBasicFields(codeverticalstructuremarkingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVerticalStructureMaterialType(mapOrigCopy map[any]any, codeverticalstructurematerialtypeFrom *CodeVerticalStructureMaterialType) (codeverticalstructurematerialtypeTo *CodeVerticalStructureMaterialType) {
	var alreadyCopied bool
	codeverticalstructurematerialtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeverticalstructurematerialtypeFrom)
	if alreadyCopied {
		return
	}
	codeverticalstructurematerialtypeFrom.GongCopyBasicFields(codeverticalstructurematerialtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVerticalStructureType(mapOrigCopy map[any]any, codeverticalstructuretypeFrom *CodeVerticalStructureType) (codeverticalstructuretypeTo *CodeVerticalStructureType) {
	var alreadyCopied bool
	codeverticalstructuretypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeverticalstructuretypeFrom)
	if alreadyCopied {
		return
	}
	codeverticalstructuretypeFrom.GongCopyBasicFields(codeverticalstructuretypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVisualDockingGuidanceType(mapOrigCopy map[any]any, codevisualdockingguidancetypeFrom *CodeVisualDockingGuidanceType) (codevisualdockingguidancetypeTo *CodeVisualDockingGuidanceType) {
	var alreadyCopied bool
	codevisualdockingguidancetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codevisualdockingguidancetypeFrom)
	if alreadyCopied {
		return
	}
	codevisualdockingguidancetypeFrom.GongCopyBasicFields(codevisualdockingguidancetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeWakeTurbulenceType(mapOrigCopy map[any]any, codewaketurbulencetypeFrom *CodeWakeTurbulenceType) (codewaketurbulencetypeTo *CodeWakeTurbulenceType) {
	var alreadyCopied bool
	codewaketurbulencetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codewaketurbulencetypeFrom)
	if alreadyCopied {
		return
	}
	codewaketurbulencetypeFrom.GongCopyBasicFields(codewaketurbulencetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeWorkAreaType(mapOrigCopy map[any]any, codeworkareatypeFrom *CodeWorkAreaType) (codeworkareatypeTo *CodeWorkAreaType) {
	var alreadyCopied bool
	codeworkareatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeworkareatypeFrom)
	if alreadyCopied {
		return
	}
	codeworkareatypeFrom.GongCopyBasicFields(codeworkareatypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeYesNoType(mapOrigCopy map[any]any, codeyesnotypeFrom *CodeYesNoType) (codeyesnotypeTo *CodeYesNoType) {
	var alreadyCopied bool
	codeyesnotypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeyesnotypeFrom)
	if alreadyCopied {
		return
	}
	codeyesnotypeFrom.GongCopyBasicFields(codeyesnotypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDateMonthDayType(mapOrigCopy map[any]any, datemonthdaytypeFrom *DateMonthDayType) (datemonthdaytypeTo *DateMonthDayType) {
	var alreadyCopied bool
	datemonthdaytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, datemonthdaytypeFrom)
	if alreadyCopied {
		return
	}
	datemonthdaytypeFrom.GongCopyBasicFields(datemonthdaytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDateTimeType(mapOrigCopy map[any]any, datetimetypeFrom *DateTimeType) (datetimetypeTo *DateTimeType) {
	var alreadyCopied bool
	datetimetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, datetimetypeFrom)
	if alreadyCopied {
		return
	}
	datetimetypeFrom.GongCopyBasicFields(datetimetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDateType(mapOrigCopy map[any]any, datetypeFrom *DateType) (datetypeTo *DateType) {
	var alreadyCopied bool
	datetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, datetypeFrom)
	if alreadyCopied {
		return
	}
	datetypeFrom.GongCopyBasicFields(datetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDateYearType(mapOrigCopy map[any]any, dateyeartypeFrom *DateYearType) (dateyeartypeTo *DateYearType) {
	var alreadyCopied bool
	dateyeartypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dateyeartypeFrom)
	if alreadyCopied {
		return
	}
	dateyeartypeFrom.GongCopyBasicFields(dateyeartypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoNumberType(mapOrigCopy map[any]any, nonumbertypeFrom *NoNumberType) (nonumbertypeTo *NoNumberType) {
	var alreadyCopied bool
	nonumbertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, nonumbertypeFrom)
	if alreadyCopied {
		return
	}
	nonumbertypeFrom.GongCopyBasicFields(nonumbertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoSequenceType(mapOrigCopy map[any]any, nosequencetypeFrom *NoSequenceType) (nosequencetypeTo *NoSequenceType) {
	var alreadyCopied bool
	nosequencetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, nosequencetypeFrom)
	if alreadyCopied {
		return
	}
	nosequencetypeFrom.GongCopyBasicFields(nosequencetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextAddressType(mapOrigCopy map[any]any, textaddresstypeFrom *TextAddressType) (textaddresstypeTo *TextAddressType) {
	var alreadyCopied bool
	textaddresstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textaddresstypeFrom)
	if alreadyCopied {
		return
	}
	textaddresstypeFrom.GongCopyBasicFields(textaddresstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextDesignatorType(mapOrigCopy map[any]any, textdesignatortypeFrom *TextDesignatorType) (textdesignatortypeTo *TextDesignatorType) {
	var alreadyCopied bool
	textdesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textdesignatortypeFrom)
	if alreadyCopied {
		return
	}
	textdesignatortypeFrom.GongCopyBasicFields(textdesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextInstructionType(mapOrigCopy map[any]any, textinstructiontypeFrom *TextInstructionType) (textinstructiontypeTo *TextInstructionType) {
	var alreadyCopied bool
	textinstructiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textinstructiontypeFrom)
	if alreadyCopied {
		return
	}
	textinstructiontypeFrom.GongCopyBasicFields(textinstructiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextNameType(mapOrigCopy map[any]any, textnametypeFrom *TextNameType) (textnametypeTo *TextNameType) {
	var alreadyCopied bool
	textnametypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textnametypeFrom)
	if alreadyCopied {
		return
	}
	textnametypeFrom.GongCopyBasicFields(textnametypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextNoteType(mapOrigCopy map[any]any, textnotetypeFrom *TextNoteType) (textnotetypeTo *TextNoteType) {
	var alreadyCopied bool
	textnotetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textnotetypeFrom)
	if alreadyCopied {
		return
	}
	textnotetypeFrom.GongCopyBasicFields(textnotetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextPhoneType(mapOrigCopy map[any]any, textphonetypeFrom *TextPhoneType) (textphonetypeTo *TextPhoneType) {
	var alreadyCopied bool
	textphonetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textphonetypeFrom)
	if alreadyCopied {
		return
	}
	textphonetypeFrom.GongCopyBasicFields(textphonetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextPropertyNameType(mapOrigCopy map[any]any, textpropertynametypeFrom *TextPropertyNameType) (textpropertynametypeTo *TextPropertyNameType) {
	var alreadyCopied bool
	textpropertynametypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textpropertynametypeFrom)
	if alreadyCopied {
		return
	}
	textpropertynametypeFrom.GongCopyBasicFields(textpropertynametypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextRemarkType(mapOrigCopy map[any]any, textremarktypeFrom *TextRemarkType) (textremarktypeTo *TextRemarkType) {
	var alreadyCopied bool
	textremarktypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textremarktypeFrom)
	if alreadyCopied {
		return
	}
	textremarktypeFrom.GongCopyBasicFields(textremarktypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTextSIDSTARDesignatorType(mapOrigCopy map[any]any, textsidstardesignatortypeFrom *TextSIDSTARDesignatorType) (textsidstardesignatortypeTo *TextSIDSTARDesignatorType) {
	var alreadyCopied bool
	textsidstardesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, textsidstardesignatortypeFrom)
	if alreadyCopied {
		return
	}
	textsidstardesignatortypeFrom.GongCopyBasicFields(textsidstardesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTimeType(mapOrigCopy map[any]any, timetypeFrom *TimeType) (timetypeTo *TimeType) {
	var alreadyCopied bool
	timetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timetypeFrom)
	if alreadyCopied {
		return
	}
	timetypeFrom.GongCopyBasicFields(timetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValAlarmLimitType(mapOrigCopy map[any]any, valalarmlimittypeFrom *ValAlarmLimitType) (valalarmlimittypeTo *ValAlarmLimitType) {
	var alreadyCopied bool
	valalarmlimittypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valalarmlimittypeFrom)
	if alreadyCopied {
		return
	}
	valalarmlimittypeFrom.GongCopyBasicFields(valalarmlimittypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValAngleType(mapOrigCopy map[any]any, valangletypeFrom *ValAngleType) (valangletypeTo *ValAngleType) {
	var alreadyCopied bool
	valangletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valangletypeFrom)
	if alreadyCopied {
		return
	}
	valangletypeFrom.GongCopyBasicFields(valangletypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValBearingType(mapOrigCopy map[any]any, valbearingtypeFrom *ValBearingType) (valbearingtypeTo *ValBearingType) {
	var alreadyCopied bool
	valbearingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valbearingtypeFrom)
	if alreadyCopied {
		return
	}
	valbearingtypeFrom.GongCopyBasicFields(valbearingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValChannelNumberType(mapOrigCopy map[any]any, valchannelnumbertypeFrom *ValChannelNumberType) (valchannelnumbertypeTo *ValChannelNumberType) {
	var alreadyCopied bool
	valchannelnumbertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valchannelnumbertypeFrom)
	if alreadyCopied {
		return
	}
	valchannelnumbertypeFrom.GongCopyBasicFields(valchannelnumbertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValDepthType(mapOrigCopy map[any]any, valdepthtypeFrom *ValDepthType) (valdepthtypeTo *ValDepthType) {
	var alreadyCopied bool
	valdepthtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valdepthtypeFrom)
	if alreadyCopied {
		return
	}
	valdepthtypeFrom.GongCopyBasicFields(valdepthtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValDistanceSignedType(mapOrigCopy map[any]any, valdistancesignedtypeFrom *ValDistanceSignedType) (valdistancesignedtypeTo *ValDistanceSignedType) {
	var alreadyCopied bool
	valdistancesignedtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valdistancesignedtypeFrom)
	if alreadyCopied {
		return
	}
	valdistancesignedtypeFrom.GongCopyBasicFields(valdistancesignedtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValDistanceType(mapOrigCopy map[any]any, valdistancetypeFrom *ValDistanceType) (valdistancetypeTo *ValDistanceType) {
	var alreadyCopied bool
	valdistancetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valdistancetypeFrom)
	if alreadyCopied {
		return
	}
	valdistancetypeFrom.GongCopyBasicFields(valdistancetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValDistanceVerticalType(mapOrigCopy map[any]any, valdistanceverticaltypeFrom *ValDistanceVerticalType) (valdistanceverticaltypeTo *ValDistanceVerticalType) {
	var alreadyCopied bool
	valdistanceverticaltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valdistanceverticaltypeFrom)
	if alreadyCopied {
		return
	}
	valdistanceverticaltypeFrom.GongCopyBasicFields(valdistanceverticaltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValDurationType(mapOrigCopy map[any]any, valdurationtypeFrom *ValDurationType) (valdurationtypeTo *ValDurationType) {
	var alreadyCopied bool
	valdurationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valdurationtypeFrom)
	if alreadyCopied {
		return
	}
	valdurationtypeFrom.GongCopyBasicFields(valdurationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValFLType(mapOrigCopy map[any]any, valfltypeFrom *ValFLType) (valfltypeTo *ValFLType) {
	var alreadyCopied bool
	valfltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valfltypeFrom)
	if alreadyCopied {
		return
	}
	valfltypeFrom.GongCopyBasicFields(valfltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValFrequencyType(mapOrigCopy map[any]any, valfrequencytypeFrom *ValFrequencyType) (valfrequencytypeTo *ValFrequencyType) {
	var alreadyCopied bool
	valfrequencytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valfrequencytypeFrom)
	if alreadyCopied {
		return
	}
	valfrequencytypeFrom.GongCopyBasicFields(valfrequencytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValFrictionType(mapOrigCopy map[any]any, valfrictiontypeFrom *ValFrictionType) (valfrictiontypeTo *ValFrictionType) {
	var alreadyCopied bool
	valfrictiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valfrictiontypeFrom)
	if alreadyCopied {
		return
	}
	valfrictiontypeFrom.GongCopyBasicFields(valfrictiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValHexType(mapOrigCopy map[any]any, valhextypeFrom *ValHexType) (valhextypeTo *ValHexType) {
	var alreadyCopied bool
	valhextypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valhextypeFrom)
	if alreadyCopied {
		return
	}
	valhextypeFrom.GongCopyBasicFields(valhextypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValLCNType(mapOrigCopy map[any]any, vallcntypeFrom *ValLCNType) (vallcntypeTo *ValLCNType) {
	var alreadyCopied bool
	vallcntypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, vallcntypeFrom)
	if alreadyCopied {
		return
	}
	vallcntypeFrom.GongCopyBasicFields(vallcntypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValLightIntensityType(mapOrigCopy map[any]any, vallightintensitytypeFrom *ValLightIntensityType) (vallightintensitytypeTo *ValLightIntensityType) {
	var alreadyCopied bool
	vallightintensitytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, vallightintensitytypeFrom)
	if alreadyCopied {
		return
	}
	vallightintensitytypeFrom.GongCopyBasicFields(vallightintensitytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValMagneticVariationChangeType(mapOrigCopy map[any]any, valmagneticvariationchangetypeFrom *ValMagneticVariationChangeType) (valmagneticvariationchangetypeTo *ValMagneticVariationChangeType) {
	var alreadyCopied bool
	valmagneticvariationchangetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valmagneticvariationchangetypeFrom)
	if alreadyCopied {
		return
	}
	valmagneticvariationchangetypeFrom.GongCopyBasicFields(valmagneticvariationchangetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValMagneticVariationType(mapOrigCopy map[any]any, valmagneticvariationtypeFrom *ValMagneticVariationType) (valmagneticvariationtypeTo *ValMagneticVariationType) {
	var alreadyCopied bool
	valmagneticvariationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valmagneticvariationtypeFrom)
	if alreadyCopied {
		return
	}
	valmagneticvariationtypeFrom.GongCopyBasicFields(valmagneticvariationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValPCNType(mapOrigCopy map[any]any, valpcntypeFrom *ValPCNType) (valpcntypeTo *ValPCNType) {
	var alreadyCopied bool
	valpcntypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valpcntypeFrom)
	if alreadyCopied {
		return
	}
	valpcntypeFrom.GongCopyBasicFields(valpcntypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValPercentType(mapOrigCopy map[any]any, valpercenttypeFrom *ValPercentType) (valpercenttypeTo *ValPercentType) {
	var alreadyCopied bool
	valpercenttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valpercenttypeFrom)
	if alreadyCopied {
		return
	}
	valpercenttypeFrom.GongCopyBasicFields(valpercenttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValPressureType(mapOrigCopy map[any]any, valpressuretypeFrom *ValPressureType) (valpressuretypeTo *ValPressureType) {
	var alreadyCopied bool
	valpressuretypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valpressuretypeFrom)
	if alreadyCopied {
		return
	}
	valpressuretypeFrom.GongCopyBasicFields(valpressuretypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValSlopeType(mapOrigCopy map[any]any, valslopetypeFrom *ValSlopeType) (valslopetypeTo *ValSlopeType) {
	var alreadyCopied bool
	valslopetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valslopetypeFrom)
	if alreadyCopied {
		return
	}
	valslopetypeFrom.GongCopyBasicFields(valslopetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValSpeedType(mapOrigCopy map[any]any, valspeedtypeFrom *ValSpeedType) (valspeedtypeTo *ValSpeedType) {
	var alreadyCopied bool
	valspeedtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valspeedtypeFrom)
	if alreadyCopied {
		return
	}
	valspeedtypeFrom.GongCopyBasicFields(valspeedtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValTemperatureType(mapOrigCopy map[any]any, valtemperaturetypeFrom *ValTemperatureType) (valtemperaturetypeTo *ValTemperatureType) {
	var alreadyCopied bool
	valtemperaturetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valtemperaturetypeFrom)
	if alreadyCopied {
		return
	}
	valtemperaturetypeFrom.GongCopyBasicFields(valtemperaturetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchValWeightType(mapOrigCopy map[any]any, valweighttypeFrom *ValWeightType) (valweighttypeTo *ValWeightType) {
	var alreadyCopied bool
	valweighttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, valweighttypeFrom)
	if alreadyCopied {
		return
	}
	valweighttypeFrom.GongCopyBasicFields(valweighttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchXHTMLType(mapOrigCopy map[any]any, xhtmltypeFrom *XHTMLType) (xhtmltypeTo *XHTMLType) {
	var alreadyCopied bool
	xhtmltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, xhtmltypeFrom)
	if alreadyCopied {
		return
	}
	xhtmltypeFrom.GongCopyBasicFields(xhtmltypeTo)

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
func (codeatcreportingtype *CodeATCReportingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeatcreportingtype) {
		return
	}

	codeatcreportingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaerialrefuellingpointtype) {
		return
	}

	codeaerialrefuellingpointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaerialrefuellingprefixtype) {
		return
	}

	codeaerialrefuellingprefixtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaerialrefuellingtype) {
		return
	}

	codeaerialrefuellingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircraftcategorytype) {
		return
	}

	codeaircraftcategorytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircraftenginenumbertype) {
		return
	}

	codeaircraftenginenumbertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftenginetype *CodeAircraftEngineType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircraftenginetype) {
		return
	}

	codeaircraftenginetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircraftgroundservicetype) {
		return
	}

	codeaircraftgroundservicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircrafticaotype *CodeAircraftICAOType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircrafticaotype) {
		return
	}

	codeaircrafticaotype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftstandtype *CodeAircraftStandType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircraftstandtype) {
		return
	}

	codeaircraftstandtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircrafttype *CodeAircraftType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircrafttype) {
		return
	}

	codeaircrafttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircraftwingspanclasstype) {
		return
	}

	codeaircraftwingspanclasstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairportheliportcollocationtype) {
		return
	}

	codeairportheliportcollocationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairportheliportdesignatortype) {
		return
	}

	codeairportheliportdesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairportheliporttype *CodeAirportHeliportType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairportheliporttype) {
		return
	}

	codeairportheliporttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairportwarningtype *CodeAirportWarningType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairportwarningtype) {
		return
	}

	codeairportwarningtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspaceactivitytype) {
		return
	}

	codeairspaceactivitytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspaceaggregationtype) {
		return
	}

	codeairspaceaggregationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspaceclassificationtype) {
		return
	}

	codeairspaceclassificationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacedependencytype) {
		return
	}

	codeairspacedependencytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacedesignatortype) {
		return
	}

	codeairspacedesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacepointpositiontype) {
		return
	}

	codeairspacepointpositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacepointroletype) {
		return
	}

	codeairspacepointroletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacetype *CodeAirspaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacetype) {
		return
	}

	codeairspacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codealtitudeadjustmenttype) {
		return
	}

	codealtitudeadjustmenttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codealtitudeusetype *CodeAltitudeUseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codealtitudeusetype) {
		return
	}

	codealtitudeusetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapproachequipmentadditionaltype) {
		return
	}

	codeapproachequipmentadditionaltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapproachguidancetype) {
		return
	}

	codeapproachguidancetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapproachlightingicaotype) {
		return
	}

	codeapproachlightingicaotype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachlightingtype *CodeApproachLightingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapproachlightingtype) {
		return
	}

	codeapproachlightingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachprefixtype *CodeApproachPrefixType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapproachprefixtype) {
		return
	}

	codeapproachprefixtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapproachtype *CodeApproachType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapproachtype) {
		return
	}

	codeapproachtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapprovaltype *CodeApprovalType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapprovaltype) {
		return
	}

	codeapprovaltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapronelementtype *CodeApronElementType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapronelementtype) {
		return
	}

	codeapronelementtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeapronsectiontype *CodeApronSectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeapronsectiontype) {
		return
	}

	codeapronsectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codearcdirectiontype *CodeArcDirectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codearcdirectiontype) {
		return
	}

	codearcdirectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codearrestinggearenergyabsorbtype) {
		return
	}

	codearrestinggearenergyabsorbtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codearrestinggearengagedevicetype) {
		return
	}

	codearrestinggearengagedevicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeauralmorsetype *CodeAuralMorseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeauralmorsetype) {
		return
	}

	codeauralmorsetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeauthorityroletype *CodeAuthorityRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeauthorityroletype) {
		return
	}

	codeauthorityroletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeauthoritytype *CodeAuthorityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeauthoritytype) {
		return
	}

	codeauthoritytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeaviationstandardstype *CodeAviationStandardsType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaviationstandardstype) {
		return
	}

	codeaviationstandardstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codebearingtype *CodeBearingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codebearingtype) {
		return
	}

	codebearingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codebuoydesignatortype) {
		return
	}

	codebuoydesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codebuoytype *CodeBuoyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codebuoytype) {
		return
	}

	codebuoytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecardinaldirectiontype) {
		return
	}

	codecardinaldirectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecheckpointcategorytype) {
		return
	}

	codecheckpointcategorytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecolourtype *CodeColourType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecolourtype) {
		return
	}

	codecolourtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecommunicationchanneltype) {
		return
	}

	codecommunicationchanneltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecommunicationdirectiontype) {
		return
	}

	codecommunicationdirectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecommunicationmodetype *CodeCommunicationModeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecommunicationmodetype) {
		return
	}

	codecommunicationmodetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecomparisontype *CodeComparisonType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecomparisontype) {
		return
	}

	codecomparisontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecontaminationtype *CodeContaminationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecontaminationtype) {
		return
	}

	codecontaminationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecoursequalityilstype) {
		return
	}

	codecoursequalityilstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecoursetype *CodeCourseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecoursetype) {
		return
	}

	codecoursetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedmechanneltype *CodeDMEChannelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedmechanneltype) {
		return
	}

	codedmechanneltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedmetype *CodeDMEType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedmetype) {
		return
	}

	codedmetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedaytype *CodeDayType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedaytype) {
		return
	}

	codedaytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedeclareddistancetype) {
		return
	}

	codedeclareddistancetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedesignstandardtype *CodeDesignStandardType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedesignstandardtype) {
		return
	}

	codedesignstandardtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedesignatedpointdesignatortype) {
		return
	}

	codedesignatedpointdesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedesignatedpointtype *CodeDesignatedPointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedesignatedpointtype) {
		return
	}

	codedesignatedpointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedirectionreferencetype) {
		return
	}

	codedirectionreferencetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedirectionturntype *CodeDirectionTurnType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedirectionturntype) {
		return
	}

	codedirectionturntype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedirectiontype *CodeDirectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedirectiontype) {
		return
	}

	codedirectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedistanceindicationtype) {
		return
	}

	codedistanceindicationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeemissionbandtype *CodeEmissionBandType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeemissionbandtype) {
		return
	}

	codeemissionbandtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeequipmentanticollisiontype) {
		return
	}

	codeequipmentanticollisiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeequipmentunavailabletype) {
		return
	}

	codeequipmentunavailabletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefacilityrankingtype *CodeFacilityRankingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codefacilityrankingtype) {
		return
	}

	codefacilityrankingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefinalguidancetype *CodeFinalGuidanceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codefinalguidancetype) {
		return
	}

	codefinalguidancetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefirefightingtype *CodeFireFightingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codefirefightingtype) {
		return
	}

	codefirefightingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightdestinationtype *CodeFlightDestinationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflightdestinationtype) {
		return
	}

	codeflightdestinationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightorigintype *CodeFlightOriginType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflightorigintype) {
		return
	}

	codeflightorigintype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightpurposetype *CodeFlightPurposeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflightpurposetype) {
		return
	}

	codeflightpurposetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflightrestrictiondesignatortype) {
		return
	}

	codeflightrestrictiondesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflightrestrictiontype) {
		return
	}

	codeflightrestrictiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightruletype *CodeFlightRuleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflightruletype) {
		return
	}

	codeflightruletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflightstatustype *CodeFlightStatusType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflightstatustype) {
		return
	}

	codeflightstatustype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflighttype *CodeFlightType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflighttype) {
		return
	}

	codeflighttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflowconditionoperationtype) {
		return
	}

	codeflowconditionoperationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefreeflighttype *CodeFreeFlightType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codefreeflighttype) {
		return
	}

	codefreeflighttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codefrictiondevicetype) {
		return
	}

	codefrictiondevicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codefrictionestimatetype) {
		return
	}

	codefrictionestimatetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codefueltype *CodeFuelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codefueltype) {
		return
	}

	codefueltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codegeobordertype *CodeGeoBorderType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codegeobordertype) {
		return
	}

	codegeobordertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codegradeseparationtype *CodeGradeSeparationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codegradeseparationtype) {
		return
	}

	codegradeseparationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codegroundlightingtype *CodeGroundLightingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codegroundlightingtype) {
		return
	}

	codegroundlightingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeguidancelinetype *CodeGuidanceLineType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeguidancelinetype) {
		return
	}

	codeguidancelinetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeheightreferencetype *CodeHeightReferenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeheightreferencetype) {
		return
	}

	codeheightreferencetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codehelicopterperformancetype) {
		return
	}

	codehelicopterperformancetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeholdingcategorytype) {
		return
	}

	codeholdingcategorytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeholdingusagetype *CodeHoldingUsageType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeholdingusagetype) {
		return
	}

	codeholdingusagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeholdingusetype *CodeHoldingUseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeholdingusetype) {
		return
	}

	codeholdingusetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeiatatype *CodeIATAType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeiatatype) {
		return
	}

	codeiatatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeicaocountrytype *CodeICAOCountryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeicaocountrytype) {
		return
	}

	codeicaocountrytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeicaotype *CodeICAOType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeicaotype) {
		return
	}

	codeicaotype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeilsbackcoursetype) {
		return
	}

	codeilsbackcoursetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeintegritylevelilstype) {
		return
	}

	codeintegritylevelilstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeintensitystandbytype *CodeIntensityStandByType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeintensitystandbytype) {
		return
	}

	codeintensitystandbytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelanguagetype *CodeLanguageType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelanguagetype) {
		return
	}

	codelanguagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelevelseriestype *CodeLevelSeriesType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelevelseriestype) {
		return
	}

	codelevelseriestype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeleveltabledesignatortype) {
		return
	}

	codeleveltabledesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeleveltype *CodeLevelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeleveltype) {
		return
	}

	codeleveltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelightholdingpositiontype) {
		return
	}

	codelightholdingpositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelightintensitytype *CodeLightIntensityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelightintensitytype) {
		return
	}

	codelightintensitytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelightsourcetype *CodeLightSourceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelightsourcetype) {
		return
	}

	codelightsourcetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelightingjartype *CodeLightingJARType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelightingjartype) {
		return
	}

	codelightingjartype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeloadingbridgetype) {
		return
	}

	codeloadingbridgetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelocationqualifiertype *CodeLocationQualifierType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelocationqualifiertype) {
		return
	}

	codelocationqualifiertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelogicaloperatortype) {
		return
	}

	codelogicaloperatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemlsazimuthtype) {
		return
	}

	codemlsazimuthtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemlschanneltype *CodeMLSChannelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemlschanneltype) {
		return
	}

	codemlschanneltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemarkerbeaconsignaltype) {
		return
	}

	codemarkerbeaconsignaltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemarkingconditiontype *CodeMarkingConditionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemarkingconditiontype) {
		return
	}

	codemarkingconditiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemarkingstyletype *CodeMarkingStyleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemarkingstyletype) {
		return
	}

	codemarkingstyletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemeteoconditionstype) {
		return
	}

	codemeteoconditionstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemilitaryoperationstype) {
		return
	}

	codemilitaryoperationstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemilitaryroutepointtype) {
		return
	}

	codemilitaryroutepointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemilitarystatustype *CodeMilitaryStatusType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemilitarystatustype) {
		return
	}

	codemilitarystatustype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemilitarytrainingtype) {
		return
	}

	codemilitarytrainingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeminimafinalapproachpathtype) {
		return
	}

	codeminimafinalapproachpathtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeminimumaltitudetype) {
		return
	}

	codeminimumaltitudetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeminimumheighttype *CodeMinimumHeightType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeminimumheighttype) {
		return
	}

	codeminimumheighttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemissedapproachtype *CodeMissedApproachType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemissedapproachtype) {
		return
	}

	codemissedapproachtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codendbusagetype *CodeNDBUsageType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codendbusagetype) {
		return
	}

	codendbusagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavaiddesignatortype) {
		return
	}

	codenavaiddesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavaidpurposetype) {
		return
	}

	codenavaidpurposetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaidservicetype *CodeNavaidServiceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavaidservicetype) {
		return
	}

	codenavaidservicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavigationarearestrictiontype) {
		return
	}

	codenavigationarearestrictiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavigationareatype *CodeNavigationAreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavigationareatype) {
		return
	}

	codenavigationareatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavigationequipmenttype) {
		return
	}

	codenavigationequipmenttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavigationspecificationtype) {
		return
	}

	codenavigationspecificationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenitrogentype *CodeNitrogenType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenitrogentype) {
		return
	}

	codenitrogentype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenorthreferencetype *CodeNorthReferenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenorthreferencetype) {
		return
	}

	codenorthreferencetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenotepurposetype *CodeNotePurposeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenotepurposetype) {
		return
	}

	codenotepurposetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeobstacleareatype *CodeObstacleAreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeobstacleareatype) {
		return
	}

	codeobstacleareatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeobstacleassessmentsurfacetype) {
		return
	}

	codeobstacleassessmentsurfacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeobstructionidsurfacezonetype) {
		return
	}

	codeobstructionidsurfacezonetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeoiltype *CodeOilType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeoiltype) {
		return
	}

	codeoiltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeoperationairportheliporttype) {
		return
	}

	codeoperationairportheliporttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeoperationmanoeuvringareatype) {
		return
	}

	codeoperationmanoeuvringareatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeorganisationdesignatortype) {
		return
	}

	codeorganisationdesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeorganisationhierarchytype) {
		return
	}

	codeorganisationhierarchytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeorganisationtype *CodeOrganisationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeorganisationtype) {
		return
	}

	codeorganisationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeoxygentype *CodeOxygenType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeoxygentype) {
		return
	}

	codeoxygentype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepartype *CodePARType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepartype) {
		return
	}

	codepartype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepcnmethodtype *CodePCNMethodType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepcnmethodtype) {
		return
	}

	codepcnmethodtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepcnpavementtype *CodePCNPavementType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepcnpavementtype) {
		return
	}

	codepcnpavementtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepcnsubgradetype *CodePCNSubgradeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepcnsubgradetype) {
		return
	}

	codepcnsubgradetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepcntyrepressuretype) {
		return
	}

	codepcntyrepressuretype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepassengerservicetype *CodePassengerServiceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepassengerservicetype) {
		return
	}

	codepassengerservicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepilotcontrolledlightingtype) {
		return
	}

	codepilotcontrolledlightingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepositioninilstype *CodePositionInILSType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepositioninilstype) {
		return
	}

	codepositioninilstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprimaryradartype *CodePrimaryRadarType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeprimaryradartype) {
		return
	}

	codeprimaryradartype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeprocedureavailabilitytype) {
		return
	}

	codeprocedureavailabilitytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeprocedurecodingstandardtype) {
		return
	}

	codeprocedurecodingstandardtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeproceduredistancetype) {
		return
	}

	codeproceduredistancetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeprocedurefixroletype) {
		return
	}

	codeprocedurefixroletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeprocedurephasetype) {
		return
	}

	codeprocedurephasetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeprotectareasectiontype) {
		return
	}

	codeprotectareasectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codernptype *CodeRNPType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codernptype) {
		return
	}

	codernptype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codervrreadingtype *CodeRVRReadingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codervrreadingtype) {
		return
	}

	codervrreadingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codervsmpointroletype) {
		return
	}

	codervsmpointroletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codervsmtype *CodeRVSMType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codervsmtype) {
		return
	}

	codervsmtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadarservicetype *CodeRadarServiceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderadarservicetype) {
		return
	}

	coderadarservicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadioemissiontype *CodeRadioEmissionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderadioemissiontype) {
		return
	}

	coderadioemissiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderadiofrequencyareatype) {
		return
	}

	coderadiofrequencyareatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadiosignaltype *CodeRadioSignalType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderadiosignaltype) {
		return
	}

	coderadiosignaltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codereferencepathidentifiertype) {
		return
	}

	codereferencepathidentifiertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codereferenceroletype *CodeReferenceRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codereferenceroletype) {
		return
	}

	codereferenceroletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codereflectortype *CodeReflectorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codereflectortype) {
		return
	}

	codereflectortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderelativepositiontype *CodeRelativePositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderelativepositiontype) {
		return
	}

	coderelativepositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoadtype *CodeRoadType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderoadtype) {
		return
	}

	coderoadtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderouteavailabilitytype) {
		return
	}

	coderouteavailabilitytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderoutedesignatorlettertype) {
		return
	}

	coderoutedesignatorlettertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderoutedesignatorprefixtype) {
		return
	}

	coderoutedesignatorprefixtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderoutedesignatorsuffixtype) {
		return
	}

	coderoutedesignatorsuffixtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderouteindicatortype *CodeRouteIndicatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderouteindicatortype) {
		return
	}

	coderouteindicatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutenavigationtype *CodeRouteNavigationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderoutenavigationtype) {
		return
	}

	coderoutenavigationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderouteorigintype *CodeRouteOriginType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderouteorigintype) {
		return
	}

	coderouteorigintype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderoutesegmentpathtype) {
		return
	}

	coderoutesegmentpathtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderoutetype *CodeRouteType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderoutetype) {
		return
	}

	coderoutetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderuleproceduretitletype) {
		return
	}

	coderuleproceduretitletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderuleproceduretype *CodeRuleProcedureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderuleproceduretype) {
		return
	}

	coderuleproceduretype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwayelementtype *CodeRunwayElementType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderunwayelementtype) {
		return
	}

	coderunwayelementtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderunwaymarkingtype) {
		return
	}

	coderunwaymarkingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderunwaypointroletype) {
		return
	}

	coderunwaypointroletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderunwayprotectionareatype) {
		return
	}

	coderunwayprotectionareatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwaysectiontype *CodeRunwaySectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderunwaysectiontype) {
		return
	}

	coderunwaysectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderunwaytype *CodeRunwayType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderunwaytype) {
		return
	}

	coderunwaytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesafealtitudetype *CodeSafeAltitudeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesafealtitudetype) {
		return
	}

	codesafealtitudetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesegmentpathtype *CodeSegmentPathType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesegmentpathtype) {
		return
	}

	codesegmentpathtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesegmentterminationtype) {
		return
	}

	codesegmentterminationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeserviceatctype *CodeServiceATCType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeserviceatctype) {
		return
	}

	codeserviceatctype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeserviceatfmtype *CodeServiceATFMType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeserviceatfmtype) {
		return
	}

	codeserviceatfmtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeservicegroundcontroltype) {
		return
	}

	codeservicegroundcontroltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeserviceinformationtype *CodeServiceInformationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeserviceinformationtype) {
		return
	}

	codeserviceinformationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeservicesartype *CodeServiceSARType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeservicesartype) {
		return
	}

	codeservicesartype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesidetype *CodeSideType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesidetype) {
		return
	}

	codesidetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesignalperformanceilstype) {
		return
	}

	codesignalperformanceilstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialdatetype *CodeSpecialDateType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespecialdatetype) {
		return
	}

	codespecialdatetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespecialnavigationchaindesignatortype) {
		return
	}

	codespecialnavigationchaindesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespecialnavigationstationtype) {
		return
	}

	codespecialnavigationstationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespecialnavigationsystemtype) {
		return
	}

	codespecialnavigationsystemtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespeedreferencetype *CodeSpeedReferenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespeedreferencetype) {
		return
	}

	codespeedreferencetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestandbypowertype *CodeStandbyPowerType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestandbypowertype) {
		return
	}

	codestandbypowertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusairporttype *CodeStatusAirportType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusairporttype) {
		return
	}

	codestatusairporttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusairspacetype *CodeStatusAirspaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusairspacetype) {
		return
	}

	codestatusairspacetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusconstructiontype *CodeStatusConstructionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusconstructiontype) {
		return
	}

	codestatusconstructiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusnavaidtype *CodeStatusNavaidType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusnavaidtype) {
		return
	}

	codestatusnavaidtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusoperationstype *CodeStatusOperationsType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusoperationstype) {
		return
	}

	codestatusoperationstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusservicetype *CodeStatusServiceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusservicetype) {
		return
	}

	codestatusservicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesurfacecompositiontype) {
		return
	}

	codesurfacecompositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesurfaceconditiontype) {
		return
	}

	codesurfaceconditiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesurfacepreparationtype) {
		return
	}

	codesurfacepreparationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesystemactivationtype *CodeSystemActivationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesystemactivationtype) {
		return
	}

	codesystemactivationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetaatype *CodeTAAType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetaatype) {
		return
	}

	codetaatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetacanchanneltype *CodeTACANChannelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetacanchanneltype) {
		return
	}

	codetacanchanneltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetlofsectiontype *CodeTLOFSectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetlofsectiontype) {
		return
	}

	codetlofsectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetaxiwayelementtype) {
		return
	}

	codetaxiwayelementtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetaxiwaysectiontype) {
		return
	}

	codetaxiwaysectiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetaxiwaytype *CodeTaxiwayType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetaxiwaytype) {
		return
	}

	codetaxiwaytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetelecomnetworktype) {
		return
	}

	codetelecomnetworktype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetimeeventcombinationtype) {
		return
	}

	codetimeeventcombinationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetimeeventtype *CodeTimeEventType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetimeeventtype) {
		return
	}

	codetimeeventtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetimereferencetype *CodeTimeReferenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetimereferencetype) {
		return
	}

	codetimereferencetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetrajectorytype *CodeTrajectoryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetrajectorytype) {
		return
	}

	codetrajectorytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetranspondertype *CodeTransponderType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetranspondertype) {
		return
	}

	codetranspondertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeuuidtype *CodeUUIDType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeuuidtype) {
		return
	}

	codeuuidtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeunitdependencytype *CodeUnitDependencyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeunitdependencytype) {
		return
	}

	codeunitdependencytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeunittype *CodeUnitType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeunittype) {
		return
	}

	codeunittype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeupperalphatype *CodeUpperAlphaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeupperalphatype) {
		return
	}

	codeupperalphatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeusagelimitationtype *CodeUsageLimitationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeusagelimitationtype) {
		return
	}

	codeusagelimitationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codevasistype *CodeVASISType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codevasistype) {
		return
	}

	codevasistype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codevortype *CodeVORType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codevortype) {
		return
	}

	codevortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codevalueinterpretationtype) {
		return
	}

	codevalueinterpretationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticaldatumtype *CodeVerticalDatumType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticaldatumtype) {
		return
	}

	codeverticaldatumtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticalreferencetype) {
		return
	}

	codeverticalreferencetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticalstructuremarkingtype) {
		return
	}

	codeverticalstructuremarkingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticalstructurematerialtype) {
		return
	}

	codeverticalstructurematerialtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructuretype *CodeVerticalStructureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticalstructuretype) {
		return
	}

	codeverticalstructuretype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codevisualdockingguidancetype) {
		return
	}

	codevisualdockingguidancetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codewaketurbulencetype) {
		return
	}

	codewaketurbulencetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeworkareatype *CodeWorkAreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeworkareatype) {
		return
	}

	codeworkareatype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeyesnotype *CodeYesNoType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeyesnotype) {
		return
	}

	codeyesnotype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datemonthdaytype *DateMonthDayType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(datemonthdaytype) {
		return
	}

	datemonthdaytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datetimetype *DateTimeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(datetimetype) {
		return
	}

	datetimetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datetype *DateType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(datetype) {
		return
	}

	datetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dateyeartype *DateYearType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(dateyeartype) {
		return
	}

	dateyeartype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (nonumbertype *NoNumberType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(nonumbertype) {
		return
	}

	nonumbertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (nosequencetype *NoSequenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(nosequencetype) {
		return
	}

	nosequencetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textaddresstype *TextAddressType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textaddresstype) {
		return
	}

	textaddresstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textdesignatortype *TextDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textdesignatortype) {
		return
	}

	textdesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textinstructiontype *TextInstructionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textinstructiontype) {
		return
	}

	textinstructiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textnametype *TextNameType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textnametype) {
		return
	}

	textnametype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textnotetype *TextNoteType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textnotetype) {
		return
	}

	textnotetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textphonetype *TextPhoneType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textphonetype) {
		return
	}

	textphonetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textpropertynametype *TextPropertyNameType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textpropertynametype) {
		return
	}

	textpropertynametype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textremarktype *TextRemarkType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textremarktype) {
		return
	}

	textremarktype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textsidstardesignatortype) {
		return
	}

	textsidstardesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timetype *TimeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timetype) {
		return
	}

	timetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valalarmlimittype *ValAlarmLimitType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valalarmlimittype) {
		return
	}

	valalarmlimittype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valangletype *ValAngleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valangletype) {
		return
	}

	valangletype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valbearingtype *ValBearingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valbearingtype) {
		return
	}

	valbearingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valchannelnumbertype *ValChannelNumberType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valchannelnumbertype) {
		return
	}

	valchannelnumbertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdepthtype *ValDepthType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valdepthtype) {
		return
	}

	valdepthtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdistancesignedtype *ValDistanceSignedType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valdistancesignedtype) {
		return
	}

	valdistancesignedtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdistancetype *ValDistanceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valdistancetype) {
		return
	}

	valdistancetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdistanceverticaltype *ValDistanceVerticalType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valdistanceverticaltype) {
		return
	}

	valdistanceverticaltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valdurationtype *ValDurationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valdurationtype) {
		return
	}

	valdurationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valfltype *ValFLType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valfltype) {
		return
	}

	valfltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valfrequencytype *ValFrequencyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valfrequencytype) {
		return
	}

	valfrequencytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valfrictiontype *ValFrictionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valfrictiontype) {
		return
	}

	valfrictiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valhextype *ValHexType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valhextype) {
		return
	}

	valhextype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vallcntype *ValLCNType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(vallcntype) {
		return
	}

	vallcntype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (vallightintensitytype *ValLightIntensityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(vallightintensitytype) {
		return
	}

	vallightintensitytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valmagneticvariationchangetype) {
		return
	}

	valmagneticvariationchangetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valmagneticvariationtype *ValMagneticVariationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valmagneticvariationtype) {
		return
	}

	valmagneticvariationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valpcntype *ValPCNType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valpcntype) {
		return
	}

	valpcntype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valpercenttype *ValPercentType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valpercenttype) {
		return
	}

	valpercenttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valpressuretype *ValPressureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valpressuretype) {
		return
	}

	valpressuretype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valslopetype *ValSlopeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valslopetype) {
		return
	}

	valslopetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valspeedtype *ValSpeedType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valspeedtype) {
		return
	}

	valspeedtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valtemperaturetype *ValTemperatureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valtemperaturetype) {
		return
	}

	valtemperaturetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (valweighttype *ValWeightType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valweighttype) {
		return
	}

	valweighttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xhtmltype *XHTMLType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(xhtmltype) {
		return
	}

	xhtmltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *CodeATCReportingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeATCReportingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAerialRefuellingPointType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAerialRefuellingPointType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAerialRefuellingPrefixType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAerialRefuellingPrefixType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAerialRefuellingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAerialRefuellingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftCategoryType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftCategoryType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftEngineNumberType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftEngineNumberType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftEngineType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftEngineType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftGroundServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftGroundServiceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftICAOType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftICAOType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftStandType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftStandType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftWingspanClassType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftWingspanClassType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirportHeliportCollocationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirportHeliportCollocationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirportHeliportDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirportHeliportDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirportHeliportType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirportHeliportType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirportWarningType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirportWarningType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceActivityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceActivityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceAggregationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceAggregationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceClassificationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceClassificationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceDependencyType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceDependencyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspacePointPositionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspacePointPositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspacePointRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspacePointRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAltitudeAdjustmentType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAltitudeAdjustmentType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAltitudeUseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAltitudeUseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApproachEquipmentAdditionalType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApproachEquipmentAdditionalType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApproachGuidanceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApproachGuidanceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApproachLightingICAOType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApproachLightingICAOType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApproachLightingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApproachLightingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApproachPrefixType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApproachPrefixType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApproachType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApproachType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApprovalType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApprovalType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApronElementType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApronElementType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApronSectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApronSectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeArcDirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeArcDirectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeArrestingGearEnergyAbsorbType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeArrestingGearEnergyAbsorbType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeArrestingGearEngageDeviceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeArrestingGearEngageDeviceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAuralMorseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAuralMorseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAuthorityRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAuthorityRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAuthorityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAuthorityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAviationStandardsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAviationStandardsType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeBearingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeBearingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeBuoyDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeBuoyDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeBuoyType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeBuoyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCardinalDirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCardinalDirectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCheckpointCategoryType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCheckpointCategoryType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeColourType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeColourType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCommunicationChannelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCommunicationChannelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCommunicationDirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCommunicationDirectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCommunicationModeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCommunicationModeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeComparisonType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeComparisonType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeContaminationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeContaminationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCourseQualityILSType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCourseQualityILSType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCourseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCourseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDMEChannelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDMEChannelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDMEType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDMEType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDayType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDeclaredDistanceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDeclaredDistanceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDesignStandardType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDesignStandardType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDesignatedPointDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDesignatedPointDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDesignatedPointType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDesignatedPointType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDirectionReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDirectionReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDirectionTurnType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDirectionTurnType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDirectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDistanceIndicationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDistanceIndicationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeEmissionBandType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeEmissionBandType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeEquipmentAntiCollisionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeEquipmentAntiCollisionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeEquipmentUnavailableType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeEquipmentUnavailableType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFacilityRankingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFacilityRankingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFinalGuidanceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFinalGuidanceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFireFightingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFireFightingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlightDestinationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightDestinationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlightOriginType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightOriginType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlightPurposeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightPurposeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlightRestrictionDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightRestrictionDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlightRestrictionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightRestrictionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlightRuleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightRuleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlightStatusType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightStatusType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlightType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFlowConditionOperationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlowConditionOperationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFreeFlightType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFreeFlightType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFrictionDeviceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFrictionDeviceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFrictionEstimateType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFrictionEstimateType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFuelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFuelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeGeoBorderType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeGeoBorderType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeGradeSeparationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeGradeSeparationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeGroundLightingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeGroundLightingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeGuidanceLineType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeGuidanceLineType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeHeightReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeHeightReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeHelicopterPerformanceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeHelicopterPerformanceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeHoldingCategoryType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeHoldingCategoryType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeHoldingUsageType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeHoldingUsageType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeHoldingUseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeHoldingUseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeIATAType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeIATAType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeICAOCountryType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeICAOCountryType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeICAOType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeICAOType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeILSBackCourseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeILSBackCourseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeIntegrityLevelILSType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeIntegrityLevelILSType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeIntensityStandByType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeIntensityStandByType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLanguageType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLanguageType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLevelSeriesType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLevelSeriesType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLevelTableDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLevelTableDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLevelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLevelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLightHoldingPositionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLightHoldingPositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLightIntensityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLightIntensityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLightSourceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLightSourceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLightingJARType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLightingJARType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLoadingBridgeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLoadingBridgeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLocationQualifierType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLocationQualifierType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLogicalOperatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLogicalOperatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMLSAzimuthType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMLSAzimuthType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMLSChannelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMLSChannelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMarkerBeaconSignalType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMarkerBeaconSignalType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMarkingConditionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMarkingConditionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMarkingStyleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMarkingStyleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMeteoConditionsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMeteoConditionsType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMilitaryOperationsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMilitaryOperationsType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMilitaryRoutePointType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMilitaryRoutePointType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMilitaryStatusType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMilitaryStatusType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMilitaryTrainingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMilitaryTrainingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMinimaFinalApproachPathType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMinimaFinalApproachPathType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMinimumAltitudeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMinimumAltitudeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMinimumHeightType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMinimumHeightType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMissedApproachType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMissedApproachType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNDBUsageType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNDBUsageType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavaidDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavaidDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavaidPurposeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavaidPurposeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavaidServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavaidServiceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavigationAreaRestrictionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavigationAreaRestrictionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavigationAreaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavigationAreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavigationEquipmentType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavigationEquipmentType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavigationSpecificationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavigationSpecificationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNitrogenType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNitrogenType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNorthReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNorthReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNotePurposeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNotePurposeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeObstacleAreaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeObstacleAreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeObstacleAssessmentSurfaceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeObstacleAssessmentSurfaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeObstructionIdSurfaceZoneType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeObstructionIdSurfaceZoneType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeOilType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOilType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeOperationAirportHeliportType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOperationAirportHeliportType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeOperationManoeuvringAreaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOperationManoeuvringAreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeOrganisationDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOrganisationDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeOrganisationHierarchyType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOrganisationHierarchyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeOrganisationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOrganisationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeOxygenType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOxygenType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePARType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePARType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePCNMethodType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePCNMethodType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePCNPavementType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePCNPavementType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePCNSubgradeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePCNSubgradeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePCNTyrePressureType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePCNTyrePressureType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePassengerServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePassengerServiceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePilotControlledLightingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePilotControlledLightingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePositionInILSType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePositionInILSType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePrimaryRadarType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePrimaryRadarType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeProcedureAvailabilityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeProcedureAvailabilityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeProcedureCodingStandardType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeProcedureCodingStandardType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeProcedureDistanceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeProcedureDistanceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeProcedureFixRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeProcedureFixRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeProcedurePhaseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeProcedurePhaseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeProtectAreaSectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeProtectAreaSectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRNPType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRNPType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRVRReadingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRVRReadingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRVSMPointRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRVSMPointRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRVSMType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRVSMType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRadarServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRadarServiceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRadioEmissionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRadioEmissionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRadioFrequencyAreaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRadioFrequencyAreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRadioSignalType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRadioSignalType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeReferencePathIdentifierType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeReferencePathIdentifierType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeReferenceRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeReferenceRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeReflectorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeReflectorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRelativePositionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRelativePositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRoadType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRoadType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteAvailabilityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteAvailabilityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteDesignatorLetterType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteDesignatorLetterType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteDesignatorPrefixType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteDesignatorPrefixType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteDesignatorSuffixType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteDesignatorSuffixType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteIndicatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteIndicatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteNavigationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteNavigationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteOriginType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteOriginType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteSegmentPathType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteSegmentPathType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRuleProcedureTitleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRuleProcedureTitleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRuleProcedureType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRuleProcedureType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRunwayElementType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRunwayElementType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRunwayMarkingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRunwayMarkingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRunwayPointRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRunwayPointRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRunwayProtectionAreaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRunwayProtectionAreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRunwaySectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRunwaySectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRunwayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRunwayType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSafeAltitudeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSafeAltitudeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSegmentPathType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSegmentPathType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSegmentTerminationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSegmentTerminationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeServiceATCType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeServiceATCType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeServiceATFMType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeServiceATFMType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeServiceGroundControlType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeServiceGroundControlType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeServiceInformationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeServiceInformationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeServiceSARType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeServiceSARType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSideType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSideType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSignalPerformanceILSType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSignalPerformanceILSType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpecialDateType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpecialDateType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpecialNavigationChainDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpecialNavigationChainDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpecialNavigationStationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpecialNavigationStationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpecialNavigationSystemType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpecialNavigationSystemType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpeedReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpeedReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStandbyPowerType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStandbyPowerType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusAirportType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusAirportType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusAirspaceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusAirspaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusConstructionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusConstructionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusNavaidType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusNavaidType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusOperationsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusOperationsType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusServiceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSurfaceCompositionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSurfaceCompositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSurfaceConditionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSurfaceConditionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSurfacePreparationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSurfacePreparationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSystemActivationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSystemActivationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTAAType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTAAType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTACANChannelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTACANChannelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTLOFSectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTLOFSectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTaxiwayElementType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTaxiwayElementType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTaxiwaySectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTaxiwaySectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTaxiwayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTaxiwayType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTelecomNetworkType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTelecomNetworkType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTimeEventCombinationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTimeEventCombinationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTimeEventType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTimeEventType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTimeReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTimeReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTrajectoryType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTrajectoryType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTransponderType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTransponderType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeUUIDType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeUUIDType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeUnitDependencyType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeUnitDependencyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeUnitType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeUnitType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeUpperAlphaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeUpperAlphaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeUsageLimitationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeUsageLimitationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVASISType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVASISType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVORType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVORType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeValueInterpretationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeValueInterpretationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalDatumType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalDatumType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalStructureMarkingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalStructureMarkingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalStructureMaterialType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalStructureMaterialType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalStructureType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalStructureType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVisualDockingGuidanceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVisualDockingGuidanceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeWakeTurbulenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeWakeTurbulenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeWorkAreaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeWorkAreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeYesNoType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeYesNoType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DateMonthDayType) GongReconstructPointersFromReferences(stage *Stage, instance *DateMonthDayType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DateTimeType) GongReconstructPointersFromReferences(stage *Stage, instance *DateTimeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DateType) GongReconstructPointersFromReferences(stage *Stage, instance *DateType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DateYearType) GongReconstructPointersFromReferences(stage *Stage, instance *DateYearType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NoNumberType) GongReconstructPointersFromReferences(stage *Stage, instance *NoNumberType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NoSequenceType) GongReconstructPointersFromReferences(stage *Stage, instance *NoSequenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextAddressType) GongReconstructPointersFromReferences(stage *Stage, instance *TextAddressType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *TextDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextInstructionType) GongReconstructPointersFromReferences(stage *Stage, instance *TextInstructionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextNameType) GongReconstructPointersFromReferences(stage *Stage, instance *TextNameType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextNoteType) GongReconstructPointersFromReferences(stage *Stage, instance *TextNoteType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextPhoneType) GongReconstructPointersFromReferences(stage *Stage, instance *TextPhoneType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextPropertyNameType) GongReconstructPointersFromReferences(stage *Stage, instance *TextPropertyNameType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextRemarkType) GongReconstructPointersFromReferences(stage *Stage, instance *TextRemarkType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TextSIDSTARDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *TextSIDSTARDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TimeType) GongReconstructPointersFromReferences(stage *Stage, instance *TimeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValAlarmLimitType) GongReconstructPointersFromReferences(stage *Stage, instance *ValAlarmLimitType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValAngleType) GongReconstructPointersFromReferences(stage *Stage, instance *ValAngleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValBearingType) GongReconstructPointersFromReferences(stage *Stage, instance *ValBearingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValChannelNumberType) GongReconstructPointersFromReferences(stage *Stage, instance *ValChannelNumberType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValDepthType) GongReconstructPointersFromReferences(stage *Stage, instance *ValDepthType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValDistanceSignedType) GongReconstructPointersFromReferences(stage *Stage, instance *ValDistanceSignedType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValDistanceType) GongReconstructPointersFromReferences(stage *Stage, instance *ValDistanceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValDistanceVerticalType) GongReconstructPointersFromReferences(stage *Stage, instance *ValDistanceVerticalType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValDurationType) GongReconstructPointersFromReferences(stage *Stage, instance *ValDurationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValFLType) GongReconstructPointersFromReferences(stage *Stage, instance *ValFLType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValFrequencyType) GongReconstructPointersFromReferences(stage *Stage, instance *ValFrequencyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValFrictionType) GongReconstructPointersFromReferences(stage *Stage, instance *ValFrictionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValHexType) GongReconstructPointersFromReferences(stage *Stage, instance *ValHexType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValLCNType) GongReconstructPointersFromReferences(stage *Stage, instance *ValLCNType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValLightIntensityType) GongReconstructPointersFromReferences(stage *Stage, instance *ValLightIntensityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValMagneticVariationChangeType) GongReconstructPointersFromReferences(stage *Stage, instance *ValMagneticVariationChangeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValMagneticVariationType) GongReconstructPointersFromReferences(stage *Stage, instance *ValMagneticVariationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValPCNType) GongReconstructPointersFromReferences(stage *Stage, instance *ValPCNType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValPercentType) GongReconstructPointersFromReferences(stage *Stage, instance *ValPercentType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValPressureType) GongReconstructPointersFromReferences(stage *Stage, instance *ValPressureType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValSlopeType) GongReconstructPointersFromReferences(stage *Stage, instance *ValSlopeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValSpeedType) GongReconstructPointersFromReferences(stage *Stage, instance *ValSpeedType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValTemperatureType) GongReconstructPointersFromReferences(stage *Stage, instance *ValTemperatureType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValWeightType) GongReconstructPointersFromReferences(stage *Stage, instance *ValWeightType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *XHTMLType) GongReconstructPointersFromReferences(stage *Stage, instance *XHTMLType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *CodeATCReportingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAerialRefuellingPointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAerialRefuellingPrefixType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAerialRefuellingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftCategoryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftEngineNumberType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftEngineType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftGroundServiceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftICAOType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftStandType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftWingspanClassType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirportHeliportCollocationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirportHeliportDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirportHeliportType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirportWarningType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceActivityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceAggregationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceClassificationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceDependencyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspacePointPositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspacePointRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAltitudeAdjustmentType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAltitudeUseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApproachEquipmentAdditionalType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApproachGuidanceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApproachLightingICAOType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApproachLightingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApproachPrefixType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApproachType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApprovalType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApronElementType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApronSectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeArcDirectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeArrestingGearEnergyAbsorbType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeArrestingGearEngageDeviceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAuralMorseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAuthorityRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAuthorityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAviationStandardsType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeBearingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeBuoyDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeBuoyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCardinalDirectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCheckpointCategoryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeColourType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCommunicationChannelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCommunicationDirectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCommunicationModeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeComparisonType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeContaminationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCourseQualityILSType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCourseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDMEChannelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDMEType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDayType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDeclaredDistanceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDesignStandardType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDesignatedPointDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDesignatedPointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDirectionReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDirectionTurnType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDirectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDistanceIndicationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeEmissionBandType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeEquipmentAntiCollisionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeEquipmentUnavailableType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFacilityRankingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFinalGuidanceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFireFightingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlightDestinationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlightOriginType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlightPurposeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlightRestrictionDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlightRestrictionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlightRuleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlightStatusType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlightType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFlowConditionOperationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFreeFlightType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFrictionDeviceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFrictionEstimateType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFuelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeGeoBorderType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeGradeSeparationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeGroundLightingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeGuidanceLineType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeHeightReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeHelicopterPerformanceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeHoldingCategoryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeHoldingUsageType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeHoldingUseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeIATAType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeICAOCountryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeICAOType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeILSBackCourseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeIntegrityLevelILSType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeIntensityStandByType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLanguageType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLevelSeriesType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLevelTableDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLevelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLightHoldingPositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLightIntensityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLightSourceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLightingJARType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLoadingBridgeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLocationQualifierType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLogicalOperatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMLSAzimuthType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMLSChannelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMarkerBeaconSignalType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMarkingConditionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMarkingStyleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMeteoConditionsType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMilitaryOperationsType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMilitaryRoutePointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMilitaryStatusType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMilitaryTrainingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMinimaFinalApproachPathType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMinimumAltitudeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMinimumHeightType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMissedApproachType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNDBUsageType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavaidDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavaidPurposeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavaidServiceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavigationAreaRestrictionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavigationAreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavigationEquipmentType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavigationSpecificationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNitrogenType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNorthReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNotePurposeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeObstacleAreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeObstacleAssessmentSurfaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeObstructionIdSurfaceZoneType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeOilType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeOperationAirportHeliportType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeOperationManoeuvringAreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeOrganisationDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeOrganisationHierarchyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeOrganisationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeOxygenType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePARType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePCNMethodType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePCNPavementType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePCNSubgradeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePCNTyrePressureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePassengerServiceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePilotControlledLightingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePositionInILSType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePrimaryRadarType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeProcedureAvailabilityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeProcedureCodingStandardType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeProcedureDistanceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeProcedureFixRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeProcedurePhaseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeProtectAreaSectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRNPType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRVRReadingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRVSMPointRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRVSMType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRadarServiceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRadioEmissionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRadioFrequencyAreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRadioSignalType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeReferencePathIdentifierType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeReferenceRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeReflectorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRelativePositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRoadType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteAvailabilityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteDesignatorLetterType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteDesignatorPrefixType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteDesignatorSuffixType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteIndicatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteNavigationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteOriginType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteSegmentPathType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRuleProcedureTitleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRuleProcedureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRunwayElementType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRunwayMarkingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRunwayPointRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRunwayProtectionAreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRunwaySectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRunwayType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSafeAltitudeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSegmentPathType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSegmentTerminationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeServiceATCType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeServiceATFMType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeServiceGroundControlType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeServiceInformationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeServiceSARType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSideType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSignalPerformanceILSType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpecialDateType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpecialNavigationChainDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpecialNavigationStationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpecialNavigationSystemType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpeedReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStandbyPowerType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusAirportType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusAirspaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusConstructionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusNavaidType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusOperationsType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusServiceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSurfaceCompositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSurfaceConditionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSurfacePreparationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSystemActivationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTAAType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTACANChannelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTLOFSectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTaxiwayElementType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTaxiwaySectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTaxiwayType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTelecomNetworkType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTimeEventCombinationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTimeEventType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTimeReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTrajectoryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTransponderType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeUUIDType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeUnitDependencyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeUnitType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeUpperAlphaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeUsageLimitationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVASISType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVORType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeValueInterpretationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalDatumType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalStructureMarkingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalStructureMaterialType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalStructureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVisualDockingGuidanceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeWakeTurbulenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeWorkAreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeYesNoType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DateMonthDayType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DateTimeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DateType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DateYearType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NoNumberType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NoSequenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextAddressType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextInstructionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextNameType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextNoteType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextPhoneType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextPropertyNameType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextRemarkType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TextSIDSTARDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TimeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValAlarmLimitType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValAngleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValBearingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValChannelNumberType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValDepthType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValDistanceSignedType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValDistanceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValDistanceVerticalType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValDurationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValFLType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValFrequencyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValFrictionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValHexType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValLCNType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValLightIntensityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValMagneticVariationChangeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValMagneticVariationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValPCNType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValPercentType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValPressureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValSlopeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValSpeedType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValTemperatureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValWeightType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *XHTMLType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeatcreportingtype *CodeATCReportingType) GongDiff(stage *Stage, codeatcreportingtypeOther *CodeATCReportingType) (diffs []string) {
	// insertion point for field diffs
	if codeatcreportingtype.Name != codeatcreportingtypeOther.Name {
		diffs = append(diffs, codeatcreportingtype.GongMarshallField(stage, "Name"))
	}
	if codeatcreportingtype.NilReason != codeatcreportingtypeOther.NilReason {
		diffs = append(diffs, codeatcreportingtype.GongMarshallField(stage, "NilReason"))
	}
	if codeatcreportingtype.EnclosedText != codeatcreportingtypeOther.EnclosedText {
		diffs = append(diffs, codeatcreportingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongDiff(stage *Stage, codeaerialrefuellingpointtypeOther *CodeAerialRefuellingPointType) (diffs []string) {
	// insertion point for field diffs
	if codeaerialrefuellingpointtype.Name != codeaerialrefuellingpointtypeOther.Name {
		diffs = append(diffs, codeaerialrefuellingpointtype.GongMarshallField(stage, "Name"))
	}
	if codeaerialrefuellingpointtype.NilReason != codeaerialrefuellingpointtypeOther.NilReason {
		diffs = append(diffs, codeaerialrefuellingpointtype.GongMarshallField(stage, "NilReason"))
	}
	if codeaerialrefuellingpointtype.EnclosedText != codeaerialrefuellingpointtypeOther.EnclosedText {
		diffs = append(diffs, codeaerialrefuellingpointtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongDiff(stage *Stage, codeaerialrefuellingprefixtypeOther *CodeAerialRefuellingPrefixType) (diffs []string) {
	// insertion point for field diffs
	if codeaerialrefuellingprefixtype.Name != codeaerialrefuellingprefixtypeOther.Name {
		diffs = append(diffs, codeaerialrefuellingprefixtype.GongMarshallField(stage, "Name"))
	}
	if codeaerialrefuellingprefixtype.NilReason != codeaerialrefuellingprefixtypeOther.NilReason {
		diffs = append(diffs, codeaerialrefuellingprefixtype.GongMarshallField(stage, "NilReason"))
	}
	if codeaerialrefuellingprefixtype.EnclosedText != codeaerialrefuellingprefixtypeOther.EnclosedText {
		diffs = append(diffs, codeaerialrefuellingprefixtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongDiff(stage *Stage, codeaerialrefuellingtypeOther *CodeAerialRefuellingType) (diffs []string) {
	// insertion point for field diffs
	if codeaerialrefuellingtype.Name != codeaerialrefuellingtypeOther.Name {
		diffs = append(diffs, codeaerialrefuellingtype.GongMarshallField(stage, "Name"))
	}
	if codeaerialrefuellingtype.NilReason != codeaerialrefuellingtypeOther.NilReason {
		diffs = append(diffs, codeaerialrefuellingtype.GongMarshallField(stage, "NilReason"))
	}
	if codeaerialrefuellingtype.EnclosedText != codeaerialrefuellingtypeOther.EnclosedText {
		diffs = append(diffs, codeaerialrefuellingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongDiff(stage *Stage, codeaircraftcategorytypeOther *CodeAircraftCategoryType) (diffs []string) {
	// insertion point for field diffs
	if codeaircraftcategorytype.Name != codeaircraftcategorytypeOther.Name {
		diffs = append(diffs, codeaircraftcategorytype.GongMarshallField(stage, "Name"))
	}
	if codeaircraftcategorytype.NilReason != codeaircraftcategorytypeOther.NilReason {
		diffs = append(diffs, codeaircraftcategorytype.GongMarshallField(stage, "NilReason"))
	}
	if codeaircraftcategorytype.EnclosedText != codeaircraftcategorytypeOther.EnclosedText {
		diffs = append(diffs, codeaircraftcategorytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongDiff(stage *Stage, codeaircraftenginenumbertypeOther *CodeAircraftEngineNumberType) (diffs []string) {
	// insertion point for field diffs
	if codeaircraftenginenumbertype.Name != codeaircraftenginenumbertypeOther.Name {
		diffs = append(diffs, codeaircraftenginenumbertype.GongMarshallField(stage, "Name"))
	}
	if codeaircraftenginenumbertype.NilReason != codeaircraftenginenumbertypeOther.NilReason {
		diffs = append(diffs, codeaircraftenginenumbertype.GongMarshallField(stage, "NilReason"))
	}
	if codeaircraftenginenumbertype.EnclosedText != codeaircraftenginenumbertypeOther.EnclosedText {
		diffs = append(diffs, codeaircraftenginenumbertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaircraftenginetype *CodeAircraftEngineType) GongDiff(stage *Stage, codeaircraftenginetypeOther *CodeAircraftEngineType) (diffs []string) {
	// insertion point for field diffs
	if codeaircraftenginetype.Name != codeaircraftenginetypeOther.Name {
		diffs = append(diffs, codeaircraftenginetype.GongMarshallField(stage, "Name"))
	}
	if codeaircraftenginetype.NilReason != codeaircraftenginetypeOther.NilReason {
		diffs = append(diffs, codeaircraftenginetype.GongMarshallField(stage, "NilReason"))
	}
	if codeaircraftenginetype.EnclosedText != codeaircraftenginetypeOther.EnclosedText {
		diffs = append(diffs, codeaircraftenginetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongDiff(stage *Stage, codeaircraftgroundservicetypeOther *CodeAircraftGroundServiceType) (diffs []string) {
	// insertion point for field diffs
	if codeaircraftgroundservicetype.Name != codeaircraftgroundservicetypeOther.Name {
		diffs = append(diffs, codeaircraftgroundservicetype.GongMarshallField(stage, "Name"))
	}
	if codeaircraftgroundservicetype.NilReason != codeaircraftgroundservicetypeOther.NilReason {
		diffs = append(diffs, codeaircraftgroundservicetype.GongMarshallField(stage, "NilReason"))
	}
	if codeaircraftgroundservicetype.EnclosedText != codeaircraftgroundservicetypeOther.EnclosedText {
		diffs = append(diffs, codeaircraftgroundservicetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaircrafticaotype *CodeAircraftICAOType) GongDiff(stage *Stage, codeaircrafticaotypeOther *CodeAircraftICAOType) (diffs []string) {
	// insertion point for field diffs
	if codeaircrafticaotype.Name != codeaircrafticaotypeOther.Name {
		diffs = append(diffs, codeaircrafticaotype.GongMarshallField(stage, "Name"))
	}
	if codeaircrafticaotype.NilReason != codeaircrafticaotypeOther.NilReason {
		diffs = append(diffs, codeaircrafticaotype.GongMarshallField(stage, "NilReason"))
	}
	if codeaircrafticaotype.EnclosedText != codeaircrafticaotypeOther.EnclosedText {
		diffs = append(diffs, codeaircrafticaotype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaircraftstandtype *CodeAircraftStandType) GongDiff(stage *Stage, codeaircraftstandtypeOther *CodeAircraftStandType) (diffs []string) {
	// insertion point for field diffs
	if codeaircraftstandtype.Name != codeaircraftstandtypeOther.Name {
		diffs = append(diffs, codeaircraftstandtype.GongMarshallField(stage, "Name"))
	}
	if codeaircraftstandtype.NilReason != codeaircraftstandtypeOther.NilReason {
		diffs = append(diffs, codeaircraftstandtype.GongMarshallField(stage, "NilReason"))
	}
	if codeaircraftstandtype.EnclosedText != codeaircraftstandtypeOther.EnclosedText {
		diffs = append(diffs, codeaircraftstandtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaircrafttype *CodeAircraftType) GongDiff(stage *Stage, codeaircrafttypeOther *CodeAircraftType) (diffs []string) {
	// insertion point for field diffs
	if codeaircrafttype.Name != codeaircrafttypeOther.Name {
		diffs = append(diffs, codeaircrafttype.GongMarshallField(stage, "Name"))
	}
	if codeaircrafttype.NilReason != codeaircrafttypeOther.NilReason {
		diffs = append(diffs, codeaircrafttype.GongMarshallField(stage, "NilReason"))
	}
	if codeaircrafttype.EnclosedText != codeaircrafttypeOther.EnclosedText {
		diffs = append(diffs, codeaircrafttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongDiff(stage *Stage, codeaircraftwingspanclasstypeOther *CodeAircraftWingspanClassType) (diffs []string) {
	// insertion point for field diffs
	if codeaircraftwingspanclasstype.Name != codeaircraftwingspanclasstypeOther.Name {
		diffs = append(diffs, codeaircraftwingspanclasstype.GongMarshallField(stage, "Name"))
	}
	if codeaircraftwingspanclasstype.NilReason != codeaircraftwingspanclasstypeOther.NilReason {
		diffs = append(diffs, codeaircraftwingspanclasstype.GongMarshallField(stage, "NilReason"))
	}
	if codeaircraftwingspanclasstype.EnclosedText != codeaircraftwingspanclasstypeOther.EnclosedText {
		diffs = append(diffs, codeaircraftwingspanclasstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongDiff(stage *Stage, codeairportheliportcollocationtypeOther *CodeAirportHeliportCollocationType) (diffs []string) {
	// insertion point for field diffs
	if codeairportheliportcollocationtype.Name != codeairportheliportcollocationtypeOther.Name {
		diffs = append(diffs, codeairportheliportcollocationtype.GongMarshallField(stage, "Name"))
	}
	if codeairportheliportcollocationtype.NilReason != codeairportheliportcollocationtypeOther.NilReason {
		diffs = append(diffs, codeairportheliportcollocationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeairportheliportcollocationtype.EnclosedText != codeairportheliportcollocationtypeOther.EnclosedText {
		diffs = append(diffs, codeairportheliportcollocationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongDiff(stage *Stage, codeairportheliportdesignatortypeOther *CodeAirportHeliportDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codeairportheliportdesignatortype.Name != codeairportheliportdesignatortypeOther.Name {
		diffs = append(diffs, codeairportheliportdesignatortype.GongMarshallField(stage, "Name"))
	}
	if codeairportheliportdesignatortype.NilReason != codeairportheliportdesignatortypeOther.NilReason {
		diffs = append(diffs, codeairportheliportdesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codeairportheliportdesignatortype.EnclosedText != codeairportheliportdesignatortypeOther.EnclosedText {
		diffs = append(diffs, codeairportheliportdesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairportheliporttype *CodeAirportHeliportType) GongDiff(stage *Stage, codeairportheliporttypeOther *CodeAirportHeliportType) (diffs []string) {
	// insertion point for field diffs
	if codeairportheliporttype.Name != codeairportheliporttypeOther.Name {
		diffs = append(diffs, codeairportheliporttype.GongMarshallField(stage, "Name"))
	}
	if codeairportheliporttype.NilReason != codeairportheliporttypeOther.NilReason {
		diffs = append(diffs, codeairportheliporttype.GongMarshallField(stage, "NilReason"))
	}
	if codeairportheliporttype.EnclosedText != codeairportheliporttypeOther.EnclosedText {
		diffs = append(diffs, codeairportheliporttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairportwarningtype *CodeAirportWarningType) GongDiff(stage *Stage, codeairportwarningtypeOther *CodeAirportWarningType) (diffs []string) {
	// insertion point for field diffs
	if codeairportwarningtype.Name != codeairportwarningtypeOther.Name {
		diffs = append(diffs, codeairportwarningtype.GongMarshallField(stage, "Name"))
	}
	if codeairportwarningtype.NilReason != codeairportwarningtypeOther.NilReason {
		diffs = append(diffs, codeairportwarningtype.GongMarshallField(stage, "NilReason"))
	}
	if codeairportwarningtype.EnclosedText != codeairportwarningtypeOther.EnclosedText {
		diffs = append(diffs, codeairportwarningtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspaceactivitytype *CodeAirspaceActivityType) GongDiff(stage *Stage, codeairspaceactivitytypeOther *CodeAirspaceActivityType) (diffs []string) {
	// insertion point for field diffs
	if codeairspaceactivitytype.Name != codeairspaceactivitytypeOther.Name {
		diffs = append(diffs, codeairspaceactivitytype.GongMarshallField(stage, "Name"))
	}
	if codeairspaceactivitytype.NilReason != codeairspaceactivitytypeOther.NilReason {
		diffs = append(diffs, codeairspaceactivitytype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspaceactivitytype.EnclosedText != codeairspaceactivitytypeOther.EnclosedText {
		diffs = append(diffs, codeairspaceactivitytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongDiff(stage *Stage, codeairspaceaggregationtypeOther *CodeAirspaceAggregationType) (diffs []string) {
	// insertion point for field diffs
	if codeairspaceaggregationtype.Name != codeairspaceaggregationtypeOther.Name {
		diffs = append(diffs, codeairspaceaggregationtype.GongMarshallField(stage, "Name"))
	}
	if codeairspaceaggregationtype.NilReason != codeairspaceaggregationtypeOther.NilReason {
		diffs = append(diffs, codeairspaceaggregationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspaceaggregationtype.EnclosedText != codeairspaceaggregationtypeOther.EnclosedText {
		diffs = append(diffs, codeairspaceaggregationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongDiff(stage *Stage, codeairspaceclassificationtypeOther *CodeAirspaceClassificationType) (diffs []string) {
	// insertion point for field diffs
	if codeairspaceclassificationtype.Name != codeairspaceclassificationtypeOther.Name {
		diffs = append(diffs, codeairspaceclassificationtype.GongMarshallField(stage, "Name"))
	}
	if codeairspaceclassificationtype.NilReason != codeairspaceclassificationtypeOther.NilReason {
		diffs = append(diffs, codeairspaceclassificationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspaceclassificationtype.EnclosedText != codeairspaceclassificationtypeOther.EnclosedText {
		diffs = append(diffs, codeairspaceclassificationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacedependencytype *CodeAirspaceDependencyType) GongDiff(stage *Stage, codeairspacedependencytypeOther *CodeAirspaceDependencyType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacedependencytype.Name != codeairspacedependencytypeOther.Name {
		diffs = append(diffs, codeairspacedependencytype.GongMarshallField(stage, "Name"))
	}
	if codeairspacedependencytype.NilReason != codeairspacedependencytypeOther.NilReason {
		diffs = append(diffs, codeairspacedependencytype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacedependencytype.EnclosedText != codeairspacedependencytypeOther.EnclosedText {
		diffs = append(diffs, codeairspacedependencytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongDiff(stage *Stage, codeairspacedesignatortypeOther *CodeAirspaceDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacedesignatortype.Name != codeairspacedesignatortypeOther.Name {
		diffs = append(diffs, codeairspacedesignatortype.GongMarshallField(stage, "Name"))
	}
	if codeairspacedesignatortype.NilReason != codeairspacedesignatortypeOther.NilReason {
		diffs = append(diffs, codeairspacedesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacedesignatortype.EnclosedText != codeairspacedesignatortypeOther.EnclosedText {
		diffs = append(diffs, codeairspacedesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongDiff(stage *Stage, codeairspacepointpositiontypeOther *CodeAirspacePointPositionType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacepointpositiontype.Name != codeairspacepointpositiontypeOther.Name {
		diffs = append(diffs, codeairspacepointpositiontype.GongMarshallField(stage, "Name"))
	}
	if codeairspacepointpositiontype.NilReason != codeairspacepointpositiontypeOther.NilReason {
		diffs = append(diffs, codeairspacepointpositiontype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacepointpositiontype.EnclosedText != codeairspacepointpositiontypeOther.EnclosedText {
		diffs = append(diffs, codeairspacepointpositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacepointroletype *CodeAirspacePointRoleType) GongDiff(stage *Stage, codeairspacepointroletypeOther *CodeAirspacePointRoleType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacepointroletype.Name != codeairspacepointroletypeOther.Name {
		diffs = append(diffs, codeairspacepointroletype.GongMarshallField(stage, "Name"))
	}
	if codeairspacepointroletype.NilReason != codeairspacepointroletypeOther.NilReason {
		diffs = append(diffs, codeairspacepointroletype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacepointroletype.EnclosedText != codeairspacepointroletypeOther.EnclosedText {
		diffs = append(diffs, codeairspacepointroletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacetype *CodeAirspaceType) GongDiff(stage *Stage, codeairspacetypeOther *CodeAirspaceType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacetype.Name != codeairspacetypeOther.Name {
		diffs = append(diffs, codeairspacetype.GongMarshallField(stage, "Name"))
	}
	if codeairspacetype.NilReason != codeairspacetypeOther.NilReason {
		diffs = append(diffs, codeairspacetype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacetype.EnclosedText != codeairspacetypeOther.EnclosedText {
		diffs = append(diffs, codeairspacetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongDiff(stage *Stage, codealtitudeadjustmenttypeOther *CodeAltitudeAdjustmentType) (diffs []string) {
	// insertion point for field diffs
	if codealtitudeadjustmenttype.Name != codealtitudeadjustmenttypeOther.Name {
		diffs = append(diffs, codealtitudeadjustmenttype.GongMarshallField(stage, "Name"))
	}
	if codealtitudeadjustmenttype.NilReason != codealtitudeadjustmenttypeOther.NilReason {
		diffs = append(diffs, codealtitudeadjustmenttype.GongMarshallField(stage, "NilReason"))
	}
	if codealtitudeadjustmenttype.EnclosedText != codealtitudeadjustmenttypeOther.EnclosedText {
		diffs = append(diffs, codealtitudeadjustmenttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codealtitudeusetype *CodeAltitudeUseType) GongDiff(stage *Stage, codealtitudeusetypeOther *CodeAltitudeUseType) (diffs []string) {
	// insertion point for field diffs
	if codealtitudeusetype.Name != codealtitudeusetypeOther.Name {
		diffs = append(diffs, codealtitudeusetype.GongMarshallField(stage, "Name"))
	}
	if codealtitudeusetype.NilReason != codealtitudeusetypeOther.NilReason {
		diffs = append(diffs, codealtitudeusetype.GongMarshallField(stage, "NilReason"))
	}
	if codealtitudeusetype.EnclosedText != codealtitudeusetypeOther.EnclosedText {
		diffs = append(diffs, codealtitudeusetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongDiff(stage *Stage, codeapproachequipmentadditionaltypeOther *CodeApproachEquipmentAdditionalType) (diffs []string) {
	// insertion point for field diffs
	if codeapproachequipmentadditionaltype.Name != codeapproachequipmentadditionaltypeOther.Name {
		diffs = append(diffs, codeapproachequipmentadditionaltype.GongMarshallField(stage, "Name"))
	}
	if codeapproachequipmentadditionaltype.NilReason != codeapproachequipmentadditionaltypeOther.NilReason {
		diffs = append(diffs, codeapproachequipmentadditionaltype.GongMarshallField(stage, "NilReason"))
	}
	if codeapproachequipmentadditionaltype.EnclosedText != codeapproachequipmentadditionaltypeOther.EnclosedText {
		diffs = append(diffs, codeapproachequipmentadditionaltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapproachguidancetype *CodeApproachGuidanceType) GongDiff(stage *Stage, codeapproachguidancetypeOther *CodeApproachGuidanceType) (diffs []string) {
	// insertion point for field diffs
	if codeapproachguidancetype.Name != codeapproachguidancetypeOther.Name {
		diffs = append(diffs, codeapproachguidancetype.GongMarshallField(stage, "Name"))
	}
	if codeapproachguidancetype.NilReason != codeapproachguidancetypeOther.NilReason {
		diffs = append(diffs, codeapproachguidancetype.GongMarshallField(stage, "NilReason"))
	}
	if codeapproachguidancetype.EnclosedText != codeapproachguidancetypeOther.EnclosedText {
		diffs = append(diffs, codeapproachguidancetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongDiff(stage *Stage, codeapproachlightingicaotypeOther *CodeApproachLightingICAOType) (diffs []string) {
	// insertion point for field diffs
	if codeapproachlightingicaotype.Name != codeapproachlightingicaotypeOther.Name {
		diffs = append(diffs, codeapproachlightingicaotype.GongMarshallField(stage, "Name"))
	}
	if codeapproachlightingicaotype.NilReason != codeapproachlightingicaotypeOther.NilReason {
		diffs = append(diffs, codeapproachlightingicaotype.GongMarshallField(stage, "NilReason"))
	}
	if codeapproachlightingicaotype.EnclosedText != codeapproachlightingicaotypeOther.EnclosedText {
		diffs = append(diffs, codeapproachlightingicaotype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapproachlightingtype *CodeApproachLightingType) GongDiff(stage *Stage, codeapproachlightingtypeOther *CodeApproachLightingType) (diffs []string) {
	// insertion point for field diffs
	if codeapproachlightingtype.Name != codeapproachlightingtypeOther.Name {
		diffs = append(diffs, codeapproachlightingtype.GongMarshallField(stage, "Name"))
	}
	if codeapproachlightingtype.NilReason != codeapproachlightingtypeOther.NilReason {
		diffs = append(diffs, codeapproachlightingtype.GongMarshallField(stage, "NilReason"))
	}
	if codeapproachlightingtype.EnclosedText != codeapproachlightingtypeOther.EnclosedText {
		diffs = append(diffs, codeapproachlightingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapproachprefixtype *CodeApproachPrefixType) GongDiff(stage *Stage, codeapproachprefixtypeOther *CodeApproachPrefixType) (diffs []string) {
	// insertion point for field diffs
	if codeapproachprefixtype.Name != codeapproachprefixtypeOther.Name {
		diffs = append(diffs, codeapproachprefixtype.GongMarshallField(stage, "Name"))
	}
	if codeapproachprefixtype.NilReason != codeapproachprefixtypeOther.NilReason {
		diffs = append(diffs, codeapproachprefixtype.GongMarshallField(stage, "NilReason"))
	}
	if codeapproachprefixtype.EnclosedText != codeapproachprefixtypeOther.EnclosedText {
		diffs = append(diffs, codeapproachprefixtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapproachtype *CodeApproachType) GongDiff(stage *Stage, codeapproachtypeOther *CodeApproachType) (diffs []string) {
	// insertion point for field diffs
	if codeapproachtype.Name != codeapproachtypeOther.Name {
		diffs = append(diffs, codeapproachtype.GongMarshallField(stage, "Name"))
	}
	if codeapproachtype.NilReason != codeapproachtypeOther.NilReason {
		diffs = append(diffs, codeapproachtype.GongMarshallField(stage, "NilReason"))
	}
	if codeapproachtype.EnclosedText != codeapproachtypeOther.EnclosedText {
		diffs = append(diffs, codeapproachtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapprovaltype *CodeApprovalType) GongDiff(stage *Stage, codeapprovaltypeOther *CodeApprovalType) (diffs []string) {
	// insertion point for field diffs
	if codeapprovaltype.Name != codeapprovaltypeOther.Name {
		diffs = append(diffs, codeapprovaltype.GongMarshallField(stage, "Name"))
	}
	if codeapprovaltype.NilReason != codeapprovaltypeOther.NilReason {
		diffs = append(diffs, codeapprovaltype.GongMarshallField(stage, "NilReason"))
	}
	if codeapprovaltype.EnclosedText != codeapprovaltypeOther.EnclosedText {
		diffs = append(diffs, codeapprovaltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapronelementtype *CodeApronElementType) GongDiff(stage *Stage, codeapronelementtypeOther *CodeApronElementType) (diffs []string) {
	// insertion point for field diffs
	if codeapronelementtype.Name != codeapronelementtypeOther.Name {
		diffs = append(diffs, codeapronelementtype.GongMarshallField(stage, "Name"))
	}
	if codeapronelementtype.NilReason != codeapronelementtypeOther.NilReason {
		diffs = append(diffs, codeapronelementtype.GongMarshallField(stage, "NilReason"))
	}
	if codeapronelementtype.EnclosedText != codeapronelementtypeOther.EnclosedText {
		diffs = append(diffs, codeapronelementtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeapronsectiontype *CodeApronSectionType) GongDiff(stage *Stage, codeapronsectiontypeOther *CodeApronSectionType) (diffs []string) {
	// insertion point for field diffs
	if codeapronsectiontype.Name != codeapronsectiontypeOther.Name {
		diffs = append(diffs, codeapronsectiontype.GongMarshallField(stage, "Name"))
	}
	if codeapronsectiontype.NilReason != codeapronsectiontypeOther.NilReason {
		diffs = append(diffs, codeapronsectiontype.GongMarshallField(stage, "NilReason"))
	}
	if codeapronsectiontype.EnclosedText != codeapronsectiontypeOther.EnclosedText {
		diffs = append(diffs, codeapronsectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codearcdirectiontype *CodeArcDirectionType) GongDiff(stage *Stage, codearcdirectiontypeOther *CodeArcDirectionType) (diffs []string) {
	// insertion point for field diffs
	if codearcdirectiontype.Name != codearcdirectiontypeOther.Name {
		diffs = append(diffs, codearcdirectiontype.GongMarshallField(stage, "Name"))
	}
	if codearcdirectiontype.NilReason != codearcdirectiontypeOther.NilReason {
		diffs = append(diffs, codearcdirectiontype.GongMarshallField(stage, "NilReason"))
	}
	if codearcdirectiontype.EnclosedText != codearcdirectiontypeOther.EnclosedText {
		diffs = append(diffs, codearcdirectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongDiff(stage *Stage, codearrestinggearenergyabsorbtypeOther *CodeArrestingGearEnergyAbsorbType) (diffs []string) {
	// insertion point for field diffs
	if codearrestinggearenergyabsorbtype.Name != codearrestinggearenergyabsorbtypeOther.Name {
		diffs = append(diffs, codearrestinggearenergyabsorbtype.GongMarshallField(stage, "Name"))
	}
	if codearrestinggearenergyabsorbtype.NilReason != codearrestinggearenergyabsorbtypeOther.NilReason {
		diffs = append(diffs, codearrestinggearenergyabsorbtype.GongMarshallField(stage, "NilReason"))
	}
	if codearrestinggearenergyabsorbtype.EnclosedText != codearrestinggearenergyabsorbtypeOther.EnclosedText {
		diffs = append(diffs, codearrestinggearenergyabsorbtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongDiff(stage *Stage, codearrestinggearengagedevicetypeOther *CodeArrestingGearEngageDeviceType) (diffs []string) {
	// insertion point for field diffs
	if codearrestinggearengagedevicetype.Name != codearrestinggearengagedevicetypeOther.Name {
		diffs = append(diffs, codearrestinggearengagedevicetype.GongMarshallField(stage, "Name"))
	}
	if codearrestinggearengagedevicetype.NilReason != codearrestinggearengagedevicetypeOther.NilReason {
		diffs = append(diffs, codearrestinggearengagedevicetype.GongMarshallField(stage, "NilReason"))
	}
	if codearrestinggearengagedevicetype.EnclosedText != codearrestinggearengagedevicetypeOther.EnclosedText {
		diffs = append(diffs, codearrestinggearengagedevicetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeauralmorsetype *CodeAuralMorseType) GongDiff(stage *Stage, codeauralmorsetypeOther *CodeAuralMorseType) (diffs []string) {
	// insertion point for field diffs
	if codeauralmorsetype.Name != codeauralmorsetypeOther.Name {
		diffs = append(diffs, codeauralmorsetype.GongMarshallField(stage, "Name"))
	}
	if codeauralmorsetype.NilReason != codeauralmorsetypeOther.NilReason {
		diffs = append(diffs, codeauralmorsetype.GongMarshallField(stage, "NilReason"))
	}
	if codeauralmorsetype.EnclosedText != codeauralmorsetypeOther.EnclosedText {
		diffs = append(diffs, codeauralmorsetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeauthorityroletype *CodeAuthorityRoleType) GongDiff(stage *Stage, codeauthorityroletypeOther *CodeAuthorityRoleType) (diffs []string) {
	// insertion point for field diffs
	if codeauthorityroletype.Name != codeauthorityroletypeOther.Name {
		diffs = append(diffs, codeauthorityroletype.GongMarshallField(stage, "Name"))
	}
	if codeauthorityroletype.NilReason != codeauthorityroletypeOther.NilReason {
		diffs = append(diffs, codeauthorityroletype.GongMarshallField(stage, "NilReason"))
	}
	if codeauthorityroletype.EnclosedText != codeauthorityroletypeOther.EnclosedText {
		diffs = append(diffs, codeauthorityroletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeauthoritytype *CodeAuthorityType) GongDiff(stage *Stage, codeauthoritytypeOther *CodeAuthorityType) (diffs []string) {
	// insertion point for field diffs
	if codeauthoritytype.Name != codeauthoritytypeOther.Name {
		diffs = append(diffs, codeauthoritytype.GongMarshallField(stage, "Name"))
	}
	if codeauthoritytype.NilReason != codeauthoritytypeOther.NilReason {
		diffs = append(diffs, codeauthoritytype.GongMarshallField(stage, "NilReason"))
	}
	if codeauthoritytype.EnclosedText != codeauthoritytypeOther.EnclosedText {
		diffs = append(diffs, codeauthoritytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeaviationstandardstype *CodeAviationStandardsType) GongDiff(stage *Stage, codeaviationstandardstypeOther *CodeAviationStandardsType) (diffs []string) {
	// insertion point for field diffs
	if codeaviationstandardstype.Name != codeaviationstandardstypeOther.Name {
		diffs = append(diffs, codeaviationstandardstype.GongMarshallField(stage, "Name"))
	}
	if codeaviationstandardstype.NilReason != codeaviationstandardstypeOther.NilReason {
		diffs = append(diffs, codeaviationstandardstype.GongMarshallField(stage, "NilReason"))
	}
	if codeaviationstandardstype.EnclosedText != codeaviationstandardstypeOther.EnclosedText {
		diffs = append(diffs, codeaviationstandardstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codebearingtype *CodeBearingType) GongDiff(stage *Stage, codebearingtypeOther *CodeBearingType) (diffs []string) {
	// insertion point for field diffs
	if codebearingtype.Name != codebearingtypeOther.Name {
		diffs = append(diffs, codebearingtype.GongMarshallField(stage, "Name"))
	}
	if codebearingtype.NilReason != codebearingtypeOther.NilReason {
		diffs = append(diffs, codebearingtype.GongMarshallField(stage, "NilReason"))
	}
	if codebearingtype.EnclosedText != codebearingtypeOther.EnclosedText {
		diffs = append(diffs, codebearingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codebuoydesignatortype *CodeBuoyDesignatorType) GongDiff(stage *Stage, codebuoydesignatortypeOther *CodeBuoyDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codebuoydesignatortype.Name != codebuoydesignatortypeOther.Name {
		diffs = append(diffs, codebuoydesignatortype.GongMarshallField(stage, "Name"))
	}
	if codebuoydesignatortype.NilReason != codebuoydesignatortypeOther.NilReason {
		diffs = append(diffs, codebuoydesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codebuoydesignatortype.EnclosedText != codebuoydesignatortypeOther.EnclosedText {
		diffs = append(diffs, codebuoydesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codebuoytype *CodeBuoyType) GongDiff(stage *Stage, codebuoytypeOther *CodeBuoyType) (diffs []string) {
	// insertion point for field diffs
	if codebuoytype.Name != codebuoytypeOther.Name {
		diffs = append(diffs, codebuoytype.GongMarshallField(stage, "Name"))
	}
	if codebuoytype.NilReason != codebuoytypeOther.NilReason {
		diffs = append(diffs, codebuoytype.GongMarshallField(stage, "NilReason"))
	}
	if codebuoytype.EnclosedText != codebuoytypeOther.EnclosedText {
		diffs = append(diffs, codebuoytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecardinaldirectiontype *CodeCardinalDirectionType) GongDiff(stage *Stage, codecardinaldirectiontypeOther *CodeCardinalDirectionType) (diffs []string) {
	// insertion point for field diffs
	if codecardinaldirectiontype.Name != codecardinaldirectiontypeOther.Name {
		diffs = append(diffs, codecardinaldirectiontype.GongMarshallField(stage, "Name"))
	}
	if codecardinaldirectiontype.NilReason != codecardinaldirectiontypeOther.NilReason {
		diffs = append(diffs, codecardinaldirectiontype.GongMarshallField(stage, "NilReason"))
	}
	if codecardinaldirectiontype.EnclosedText != codecardinaldirectiontypeOther.EnclosedText {
		diffs = append(diffs, codecardinaldirectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongDiff(stage *Stage, codecheckpointcategorytypeOther *CodeCheckpointCategoryType) (diffs []string) {
	// insertion point for field diffs
	if codecheckpointcategorytype.Name != codecheckpointcategorytypeOther.Name {
		diffs = append(diffs, codecheckpointcategorytype.GongMarshallField(stage, "Name"))
	}
	if codecheckpointcategorytype.NilReason != codecheckpointcategorytypeOther.NilReason {
		diffs = append(diffs, codecheckpointcategorytype.GongMarshallField(stage, "NilReason"))
	}
	if codecheckpointcategorytype.EnclosedText != codecheckpointcategorytypeOther.EnclosedText {
		diffs = append(diffs, codecheckpointcategorytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecolourtype *CodeColourType) GongDiff(stage *Stage, codecolourtypeOther *CodeColourType) (diffs []string) {
	// insertion point for field diffs
	if codecolourtype.Name != codecolourtypeOther.Name {
		diffs = append(diffs, codecolourtype.GongMarshallField(stage, "Name"))
	}
	if codecolourtype.NilReason != codecolourtypeOther.NilReason {
		diffs = append(diffs, codecolourtype.GongMarshallField(stage, "NilReason"))
	}
	if codecolourtype.EnclosedText != codecolourtypeOther.EnclosedText {
		diffs = append(diffs, codecolourtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecommunicationchanneltype *CodeCommunicationChannelType) GongDiff(stage *Stage, codecommunicationchanneltypeOther *CodeCommunicationChannelType) (diffs []string) {
	// insertion point for field diffs
	if codecommunicationchanneltype.Name != codecommunicationchanneltypeOther.Name {
		diffs = append(diffs, codecommunicationchanneltype.GongMarshallField(stage, "Name"))
	}
	if codecommunicationchanneltype.NilReason != codecommunicationchanneltypeOther.NilReason {
		diffs = append(diffs, codecommunicationchanneltype.GongMarshallField(stage, "NilReason"))
	}
	if codecommunicationchanneltype.EnclosedText != codecommunicationchanneltypeOther.EnclosedText {
		diffs = append(diffs, codecommunicationchanneltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongDiff(stage *Stage, codecommunicationdirectiontypeOther *CodeCommunicationDirectionType) (diffs []string) {
	// insertion point for field diffs
	if codecommunicationdirectiontype.Name != codecommunicationdirectiontypeOther.Name {
		diffs = append(diffs, codecommunicationdirectiontype.GongMarshallField(stage, "Name"))
	}
	if codecommunicationdirectiontype.NilReason != codecommunicationdirectiontypeOther.NilReason {
		diffs = append(diffs, codecommunicationdirectiontype.GongMarshallField(stage, "NilReason"))
	}
	if codecommunicationdirectiontype.EnclosedText != codecommunicationdirectiontypeOther.EnclosedText {
		diffs = append(diffs, codecommunicationdirectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecommunicationmodetype *CodeCommunicationModeType) GongDiff(stage *Stage, codecommunicationmodetypeOther *CodeCommunicationModeType) (diffs []string) {
	// insertion point for field diffs
	if codecommunicationmodetype.Name != codecommunicationmodetypeOther.Name {
		diffs = append(diffs, codecommunicationmodetype.GongMarshallField(stage, "Name"))
	}
	if codecommunicationmodetype.NilReason != codecommunicationmodetypeOther.NilReason {
		diffs = append(diffs, codecommunicationmodetype.GongMarshallField(stage, "NilReason"))
	}
	if codecommunicationmodetype.EnclosedText != codecommunicationmodetypeOther.EnclosedText {
		diffs = append(diffs, codecommunicationmodetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecomparisontype *CodeComparisonType) GongDiff(stage *Stage, codecomparisontypeOther *CodeComparisonType) (diffs []string) {
	// insertion point for field diffs
	if codecomparisontype.Name != codecomparisontypeOther.Name {
		diffs = append(diffs, codecomparisontype.GongMarshallField(stage, "Name"))
	}
	if codecomparisontype.NilReason != codecomparisontypeOther.NilReason {
		diffs = append(diffs, codecomparisontype.GongMarshallField(stage, "NilReason"))
	}
	if codecomparisontype.EnclosedText != codecomparisontypeOther.EnclosedText {
		diffs = append(diffs, codecomparisontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecontaminationtype *CodeContaminationType) GongDiff(stage *Stage, codecontaminationtypeOther *CodeContaminationType) (diffs []string) {
	// insertion point for field diffs
	if codecontaminationtype.Name != codecontaminationtypeOther.Name {
		diffs = append(diffs, codecontaminationtype.GongMarshallField(stage, "Name"))
	}
	if codecontaminationtype.NilReason != codecontaminationtypeOther.NilReason {
		diffs = append(diffs, codecontaminationtype.GongMarshallField(stage, "NilReason"))
	}
	if codecontaminationtype.EnclosedText != codecontaminationtypeOther.EnclosedText {
		diffs = append(diffs, codecontaminationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecoursequalityilstype *CodeCourseQualityILSType) GongDiff(stage *Stage, codecoursequalityilstypeOther *CodeCourseQualityILSType) (diffs []string) {
	// insertion point for field diffs
	if codecoursequalityilstype.Name != codecoursequalityilstypeOther.Name {
		diffs = append(diffs, codecoursequalityilstype.GongMarshallField(stage, "Name"))
	}
	if codecoursequalityilstype.NilReason != codecoursequalityilstypeOther.NilReason {
		diffs = append(diffs, codecoursequalityilstype.GongMarshallField(stage, "NilReason"))
	}
	if codecoursequalityilstype.EnclosedText != codecoursequalityilstypeOther.EnclosedText {
		diffs = append(diffs, codecoursequalityilstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecoursetype *CodeCourseType) GongDiff(stage *Stage, codecoursetypeOther *CodeCourseType) (diffs []string) {
	// insertion point for field diffs
	if codecoursetype.Name != codecoursetypeOther.Name {
		diffs = append(diffs, codecoursetype.GongMarshallField(stage, "Name"))
	}
	if codecoursetype.NilReason != codecoursetypeOther.NilReason {
		diffs = append(diffs, codecoursetype.GongMarshallField(stage, "NilReason"))
	}
	if codecoursetype.EnclosedText != codecoursetypeOther.EnclosedText {
		diffs = append(diffs, codecoursetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedmechanneltype *CodeDMEChannelType) GongDiff(stage *Stage, codedmechanneltypeOther *CodeDMEChannelType) (diffs []string) {
	// insertion point for field diffs
	if codedmechanneltype.Name != codedmechanneltypeOther.Name {
		diffs = append(diffs, codedmechanneltype.GongMarshallField(stage, "Name"))
	}
	if codedmechanneltype.NilReason != codedmechanneltypeOther.NilReason {
		diffs = append(diffs, codedmechanneltype.GongMarshallField(stage, "NilReason"))
	}
	if codedmechanneltype.EnclosedText != codedmechanneltypeOther.EnclosedText {
		diffs = append(diffs, codedmechanneltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedmetype *CodeDMEType) GongDiff(stage *Stage, codedmetypeOther *CodeDMEType) (diffs []string) {
	// insertion point for field diffs
	if codedmetype.Name != codedmetypeOther.Name {
		diffs = append(diffs, codedmetype.GongMarshallField(stage, "Name"))
	}
	if codedmetype.NilReason != codedmetypeOther.NilReason {
		diffs = append(diffs, codedmetype.GongMarshallField(stage, "NilReason"))
	}
	if codedmetype.EnclosedText != codedmetypeOther.EnclosedText {
		diffs = append(diffs, codedmetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedaytype *CodeDayType) GongDiff(stage *Stage, codedaytypeOther *CodeDayType) (diffs []string) {
	// insertion point for field diffs
	if codedaytype.Name != codedaytypeOther.Name {
		diffs = append(diffs, codedaytype.GongMarshallField(stage, "Name"))
	}
	if codedaytype.NilReason != codedaytypeOther.NilReason {
		diffs = append(diffs, codedaytype.GongMarshallField(stage, "NilReason"))
	}
	if codedaytype.EnclosedText != codedaytypeOther.EnclosedText {
		diffs = append(diffs, codedaytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedeclareddistancetype *CodeDeclaredDistanceType) GongDiff(stage *Stage, codedeclareddistancetypeOther *CodeDeclaredDistanceType) (diffs []string) {
	// insertion point for field diffs
	if codedeclareddistancetype.Name != codedeclareddistancetypeOther.Name {
		diffs = append(diffs, codedeclareddistancetype.GongMarshallField(stage, "Name"))
	}
	if codedeclareddistancetype.NilReason != codedeclareddistancetypeOther.NilReason {
		diffs = append(diffs, codedeclareddistancetype.GongMarshallField(stage, "NilReason"))
	}
	if codedeclareddistancetype.EnclosedText != codedeclareddistancetypeOther.EnclosedText {
		diffs = append(diffs, codedeclareddistancetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedesignstandardtype *CodeDesignStandardType) GongDiff(stage *Stage, codedesignstandardtypeOther *CodeDesignStandardType) (diffs []string) {
	// insertion point for field diffs
	if codedesignstandardtype.Name != codedesignstandardtypeOther.Name {
		diffs = append(diffs, codedesignstandardtype.GongMarshallField(stage, "Name"))
	}
	if codedesignstandardtype.NilReason != codedesignstandardtypeOther.NilReason {
		diffs = append(diffs, codedesignstandardtype.GongMarshallField(stage, "NilReason"))
	}
	if codedesignstandardtype.EnclosedText != codedesignstandardtypeOther.EnclosedText {
		diffs = append(diffs, codedesignstandardtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongDiff(stage *Stage, codedesignatedpointdesignatortypeOther *CodeDesignatedPointDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codedesignatedpointdesignatortype.Name != codedesignatedpointdesignatortypeOther.Name {
		diffs = append(diffs, codedesignatedpointdesignatortype.GongMarshallField(stage, "Name"))
	}
	if codedesignatedpointdesignatortype.NilReason != codedesignatedpointdesignatortypeOther.NilReason {
		diffs = append(diffs, codedesignatedpointdesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codedesignatedpointdesignatortype.EnclosedText != codedesignatedpointdesignatortypeOther.EnclosedText {
		diffs = append(diffs, codedesignatedpointdesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedesignatedpointtype *CodeDesignatedPointType) GongDiff(stage *Stage, codedesignatedpointtypeOther *CodeDesignatedPointType) (diffs []string) {
	// insertion point for field diffs
	if codedesignatedpointtype.Name != codedesignatedpointtypeOther.Name {
		diffs = append(diffs, codedesignatedpointtype.GongMarshallField(stage, "Name"))
	}
	if codedesignatedpointtype.NilReason != codedesignatedpointtypeOther.NilReason {
		diffs = append(diffs, codedesignatedpointtype.GongMarshallField(stage, "NilReason"))
	}
	if codedesignatedpointtype.EnclosedText != codedesignatedpointtypeOther.EnclosedText {
		diffs = append(diffs, codedesignatedpointtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedirectionreferencetype *CodeDirectionReferenceType) GongDiff(stage *Stage, codedirectionreferencetypeOther *CodeDirectionReferenceType) (diffs []string) {
	// insertion point for field diffs
	if codedirectionreferencetype.Name != codedirectionreferencetypeOther.Name {
		diffs = append(diffs, codedirectionreferencetype.GongMarshallField(stage, "Name"))
	}
	if codedirectionreferencetype.NilReason != codedirectionreferencetypeOther.NilReason {
		diffs = append(diffs, codedirectionreferencetype.GongMarshallField(stage, "NilReason"))
	}
	if codedirectionreferencetype.EnclosedText != codedirectionreferencetypeOther.EnclosedText {
		diffs = append(diffs, codedirectionreferencetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedirectionturntype *CodeDirectionTurnType) GongDiff(stage *Stage, codedirectionturntypeOther *CodeDirectionTurnType) (diffs []string) {
	// insertion point for field diffs
	if codedirectionturntype.Name != codedirectionturntypeOther.Name {
		diffs = append(diffs, codedirectionturntype.GongMarshallField(stage, "Name"))
	}
	if codedirectionturntype.NilReason != codedirectionturntypeOther.NilReason {
		diffs = append(diffs, codedirectionturntype.GongMarshallField(stage, "NilReason"))
	}
	if codedirectionturntype.EnclosedText != codedirectionturntypeOther.EnclosedText {
		diffs = append(diffs, codedirectionturntype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedirectiontype *CodeDirectionType) GongDiff(stage *Stage, codedirectiontypeOther *CodeDirectionType) (diffs []string) {
	// insertion point for field diffs
	if codedirectiontype.Name != codedirectiontypeOther.Name {
		diffs = append(diffs, codedirectiontype.GongMarshallField(stage, "Name"))
	}
	if codedirectiontype.NilReason != codedirectiontypeOther.NilReason {
		diffs = append(diffs, codedirectiontype.GongMarshallField(stage, "NilReason"))
	}
	if codedirectiontype.EnclosedText != codedirectiontypeOther.EnclosedText {
		diffs = append(diffs, codedirectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedistanceindicationtype *CodeDistanceIndicationType) GongDiff(stage *Stage, codedistanceindicationtypeOther *CodeDistanceIndicationType) (diffs []string) {
	// insertion point for field diffs
	if codedistanceindicationtype.Name != codedistanceindicationtypeOther.Name {
		diffs = append(diffs, codedistanceindicationtype.GongMarshallField(stage, "Name"))
	}
	if codedistanceindicationtype.NilReason != codedistanceindicationtypeOther.NilReason {
		diffs = append(diffs, codedistanceindicationtype.GongMarshallField(stage, "NilReason"))
	}
	if codedistanceindicationtype.EnclosedText != codedistanceindicationtypeOther.EnclosedText {
		diffs = append(diffs, codedistanceindicationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeemissionbandtype *CodeEmissionBandType) GongDiff(stage *Stage, codeemissionbandtypeOther *CodeEmissionBandType) (diffs []string) {
	// insertion point for field diffs
	if codeemissionbandtype.Name != codeemissionbandtypeOther.Name {
		diffs = append(diffs, codeemissionbandtype.GongMarshallField(stage, "Name"))
	}
	if codeemissionbandtype.NilReason != codeemissionbandtypeOther.NilReason {
		diffs = append(diffs, codeemissionbandtype.GongMarshallField(stage, "NilReason"))
	}
	if codeemissionbandtype.EnclosedText != codeemissionbandtypeOther.EnclosedText {
		diffs = append(diffs, codeemissionbandtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongDiff(stage *Stage, codeequipmentanticollisiontypeOther *CodeEquipmentAntiCollisionType) (diffs []string) {
	// insertion point for field diffs
	if codeequipmentanticollisiontype.Name != codeequipmentanticollisiontypeOther.Name {
		diffs = append(diffs, codeequipmentanticollisiontype.GongMarshallField(stage, "Name"))
	}
	if codeequipmentanticollisiontype.NilReason != codeequipmentanticollisiontypeOther.NilReason {
		diffs = append(diffs, codeequipmentanticollisiontype.GongMarshallField(stage, "NilReason"))
	}
	if codeequipmentanticollisiontype.EnclosedText != codeequipmentanticollisiontypeOther.EnclosedText {
		diffs = append(diffs, codeequipmentanticollisiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongDiff(stage *Stage, codeequipmentunavailabletypeOther *CodeEquipmentUnavailableType) (diffs []string) {
	// insertion point for field diffs
	if codeequipmentunavailabletype.Name != codeequipmentunavailabletypeOther.Name {
		diffs = append(diffs, codeequipmentunavailabletype.GongMarshallField(stage, "Name"))
	}
	if codeequipmentunavailabletype.NilReason != codeequipmentunavailabletypeOther.NilReason {
		diffs = append(diffs, codeequipmentunavailabletype.GongMarshallField(stage, "NilReason"))
	}
	if codeequipmentunavailabletype.EnclosedText != codeequipmentunavailabletypeOther.EnclosedText {
		diffs = append(diffs, codeequipmentunavailabletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codefacilityrankingtype *CodeFacilityRankingType) GongDiff(stage *Stage, codefacilityrankingtypeOther *CodeFacilityRankingType) (diffs []string) {
	// insertion point for field diffs
	if codefacilityrankingtype.Name != codefacilityrankingtypeOther.Name {
		diffs = append(diffs, codefacilityrankingtype.GongMarshallField(stage, "Name"))
	}
	if codefacilityrankingtype.NilReason != codefacilityrankingtypeOther.NilReason {
		diffs = append(diffs, codefacilityrankingtype.GongMarshallField(stage, "NilReason"))
	}
	if codefacilityrankingtype.EnclosedText != codefacilityrankingtypeOther.EnclosedText {
		diffs = append(diffs, codefacilityrankingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codefinalguidancetype *CodeFinalGuidanceType) GongDiff(stage *Stage, codefinalguidancetypeOther *CodeFinalGuidanceType) (diffs []string) {
	// insertion point for field diffs
	if codefinalguidancetype.Name != codefinalguidancetypeOther.Name {
		diffs = append(diffs, codefinalguidancetype.GongMarshallField(stage, "Name"))
	}
	if codefinalguidancetype.NilReason != codefinalguidancetypeOther.NilReason {
		diffs = append(diffs, codefinalguidancetype.GongMarshallField(stage, "NilReason"))
	}
	if codefinalguidancetype.EnclosedText != codefinalguidancetypeOther.EnclosedText {
		diffs = append(diffs, codefinalguidancetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codefirefightingtype *CodeFireFightingType) GongDiff(stage *Stage, codefirefightingtypeOther *CodeFireFightingType) (diffs []string) {
	// insertion point for field diffs
	if codefirefightingtype.Name != codefirefightingtypeOther.Name {
		diffs = append(diffs, codefirefightingtype.GongMarshallField(stage, "Name"))
	}
	if codefirefightingtype.NilReason != codefirefightingtypeOther.NilReason {
		diffs = append(diffs, codefirefightingtype.GongMarshallField(stage, "NilReason"))
	}
	if codefirefightingtype.EnclosedText != codefirefightingtypeOther.EnclosedText {
		diffs = append(diffs, codefirefightingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflightdestinationtype *CodeFlightDestinationType) GongDiff(stage *Stage, codeflightdestinationtypeOther *CodeFlightDestinationType) (diffs []string) {
	// insertion point for field diffs
	if codeflightdestinationtype.Name != codeflightdestinationtypeOther.Name {
		diffs = append(diffs, codeflightdestinationtype.GongMarshallField(stage, "Name"))
	}
	if codeflightdestinationtype.NilReason != codeflightdestinationtypeOther.NilReason {
		diffs = append(diffs, codeflightdestinationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeflightdestinationtype.EnclosedText != codeflightdestinationtypeOther.EnclosedText {
		diffs = append(diffs, codeflightdestinationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflightorigintype *CodeFlightOriginType) GongDiff(stage *Stage, codeflightorigintypeOther *CodeFlightOriginType) (diffs []string) {
	// insertion point for field diffs
	if codeflightorigintype.Name != codeflightorigintypeOther.Name {
		diffs = append(diffs, codeflightorigintype.GongMarshallField(stage, "Name"))
	}
	if codeflightorigintype.NilReason != codeflightorigintypeOther.NilReason {
		diffs = append(diffs, codeflightorigintype.GongMarshallField(stage, "NilReason"))
	}
	if codeflightorigintype.EnclosedText != codeflightorigintypeOther.EnclosedText {
		diffs = append(diffs, codeflightorigintype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflightpurposetype *CodeFlightPurposeType) GongDiff(stage *Stage, codeflightpurposetypeOther *CodeFlightPurposeType) (diffs []string) {
	// insertion point for field diffs
	if codeflightpurposetype.Name != codeflightpurposetypeOther.Name {
		diffs = append(diffs, codeflightpurposetype.GongMarshallField(stage, "Name"))
	}
	if codeflightpurposetype.NilReason != codeflightpurposetypeOther.NilReason {
		diffs = append(diffs, codeflightpurposetype.GongMarshallField(stage, "NilReason"))
	}
	if codeflightpurposetype.EnclosedText != codeflightpurposetypeOther.EnclosedText {
		diffs = append(diffs, codeflightpurposetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongDiff(stage *Stage, codeflightrestrictiondesignatortypeOther *CodeFlightRestrictionDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codeflightrestrictiondesignatortype.Name != codeflightrestrictiondesignatortypeOther.Name {
		diffs = append(diffs, codeflightrestrictiondesignatortype.GongMarshallField(stage, "Name"))
	}
	if codeflightrestrictiondesignatortype.NilReason != codeflightrestrictiondesignatortypeOther.NilReason {
		diffs = append(diffs, codeflightrestrictiondesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codeflightrestrictiondesignatortype.EnclosedText != codeflightrestrictiondesignatortypeOther.EnclosedText {
		diffs = append(diffs, codeflightrestrictiondesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflightrestrictiontype *CodeFlightRestrictionType) GongDiff(stage *Stage, codeflightrestrictiontypeOther *CodeFlightRestrictionType) (diffs []string) {
	// insertion point for field diffs
	if codeflightrestrictiontype.Name != codeflightrestrictiontypeOther.Name {
		diffs = append(diffs, codeflightrestrictiontype.GongMarshallField(stage, "Name"))
	}
	if codeflightrestrictiontype.NilReason != codeflightrestrictiontypeOther.NilReason {
		diffs = append(diffs, codeflightrestrictiontype.GongMarshallField(stage, "NilReason"))
	}
	if codeflightrestrictiontype.EnclosedText != codeflightrestrictiontypeOther.EnclosedText {
		diffs = append(diffs, codeflightrestrictiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflightruletype *CodeFlightRuleType) GongDiff(stage *Stage, codeflightruletypeOther *CodeFlightRuleType) (diffs []string) {
	// insertion point for field diffs
	if codeflightruletype.Name != codeflightruletypeOther.Name {
		diffs = append(diffs, codeflightruletype.GongMarshallField(stage, "Name"))
	}
	if codeflightruletype.NilReason != codeflightruletypeOther.NilReason {
		diffs = append(diffs, codeflightruletype.GongMarshallField(stage, "NilReason"))
	}
	if codeflightruletype.EnclosedText != codeflightruletypeOther.EnclosedText {
		diffs = append(diffs, codeflightruletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflightstatustype *CodeFlightStatusType) GongDiff(stage *Stage, codeflightstatustypeOther *CodeFlightStatusType) (diffs []string) {
	// insertion point for field diffs
	if codeflightstatustype.Name != codeflightstatustypeOther.Name {
		diffs = append(diffs, codeflightstatustype.GongMarshallField(stage, "Name"))
	}
	if codeflightstatustype.NilReason != codeflightstatustypeOther.NilReason {
		diffs = append(diffs, codeflightstatustype.GongMarshallField(stage, "NilReason"))
	}
	if codeflightstatustype.EnclosedText != codeflightstatustypeOther.EnclosedText {
		diffs = append(diffs, codeflightstatustype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflighttype *CodeFlightType) GongDiff(stage *Stage, codeflighttypeOther *CodeFlightType) (diffs []string) {
	// insertion point for field diffs
	if codeflighttype.Name != codeflighttypeOther.Name {
		diffs = append(diffs, codeflighttype.GongMarshallField(stage, "Name"))
	}
	if codeflighttype.NilReason != codeflighttypeOther.NilReason {
		diffs = append(diffs, codeflighttype.GongMarshallField(stage, "NilReason"))
	}
	if codeflighttype.EnclosedText != codeflighttypeOther.EnclosedText {
		diffs = append(diffs, codeflighttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongDiff(stage *Stage, codeflowconditionoperationtypeOther *CodeFlowConditionOperationType) (diffs []string) {
	// insertion point for field diffs
	if codeflowconditionoperationtype.Name != codeflowconditionoperationtypeOther.Name {
		diffs = append(diffs, codeflowconditionoperationtype.GongMarshallField(stage, "Name"))
	}
	if codeflowconditionoperationtype.NilReason != codeflowconditionoperationtypeOther.NilReason {
		diffs = append(diffs, codeflowconditionoperationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeflowconditionoperationtype.EnclosedText != codeflowconditionoperationtypeOther.EnclosedText {
		diffs = append(diffs, codeflowconditionoperationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codefreeflighttype *CodeFreeFlightType) GongDiff(stage *Stage, codefreeflighttypeOther *CodeFreeFlightType) (diffs []string) {
	// insertion point for field diffs
	if codefreeflighttype.Name != codefreeflighttypeOther.Name {
		diffs = append(diffs, codefreeflighttype.GongMarshallField(stage, "Name"))
	}
	if codefreeflighttype.NilReason != codefreeflighttypeOther.NilReason {
		diffs = append(diffs, codefreeflighttype.GongMarshallField(stage, "NilReason"))
	}
	if codefreeflighttype.EnclosedText != codefreeflighttypeOther.EnclosedText {
		diffs = append(diffs, codefreeflighttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codefrictiondevicetype *CodeFrictionDeviceType) GongDiff(stage *Stage, codefrictiondevicetypeOther *CodeFrictionDeviceType) (diffs []string) {
	// insertion point for field diffs
	if codefrictiondevicetype.Name != codefrictiondevicetypeOther.Name {
		diffs = append(diffs, codefrictiondevicetype.GongMarshallField(stage, "Name"))
	}
	if codefrictiondevicetype.NilReason != codefrictiondevicetypeOther.NilReason {
		diffs = append(diffs, codefrictiondevicetype.GongMarshallField(stage, "NilReason"))
	}
	if codefrictiondevicetype.EnclosedText != codefrictiondevicetypeOther.EnclosedText {
		diffs = append(diffs, codefrictiondevicetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codefrictionestimatetype *CodeFrictionEstimateType) GongDiff(stage *Stage, codefrictionestimatetypeOther *CodeFrictionEstimateType) (diffs []string) {
	// insertion point for field diffs
	if codefrictionestimatetype.Name != codefrictionestimatetypeOther.Name {
		diffs = append(diffs, codefrictionestimatetype.GongMarshallField(stage, "Name"))
	}
	if codefrictionestimatetype.NilReason != codefrictionestimatetypeOther.NilReason {
		diffs = append(diffs, codefrictionestimatetype.GongMarshallField(stage, "NilReason"))
	}
	if codefrictionestimatetype.EnclosedText != codefrictionestimatetypeOther.EnclosedText {
		diffs = append(diffs, codefrictionestimatetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codefueltype *CodeFuelType) GongDiff(stage *Stage, codefueltypeOther *CodeFuelType) (diffs []string) {
	// insertion point for field diffs
	if codefueltype.Name != codefueltypeOther.Name {
		diffs = append(diffs, codefueltype.GongMarshallField(stage, "Name"))
	}
	if codefueltype.NilReason != codefueltypeOther.NilReason {
		diffs = append(diffs, codefueltype.GongMarshallField(stage, "NilReason"))
	}
	if codefueltype.EnclosedText != codefueltypeOther.EnclosedText {
		diffs = append(diffs, codefueltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codegeobordertype *CodeGeoBorderType) GongDiff(stage *Stage, codegeobordertypeOther *CodeGeoBorderType) (diffs []string) {
	// insertion point for field diffs
	if codegeobordertype.Name != codegeobordertypeOther.Name {
		diffs = append(diffs, codegeobordertype.GongMarshallField(stage, "Name"))
	}
	if codegeobordertype.NilReason != codegeobordertypeOther.NilReason {
		diffs = append(diffs, codegeobordertype.GongMarshallField(stage, "NilReason"))
	}
	if codegeobordertype.EnclosedText != codegeobordertypeOther.EnclosedText {
		diffs = append(diffs, codegeobordertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codegradeseparationtype *CodeGradeSeparationType) GongDiff(stage *Stage, codegradeseparationtypeOther *CodeGradeSeparationType) (diffs []string) {
	// insertion point for field diffs
	if codegradeseparationtype.Name != codegradeseparationtypeOther.Name {
		diffs = append(diffs, codegradeseparationtype.GongMarshallField(stage, "Name"))
	}
	if codegradeseparationtype.NilReason != codegradeseparationtypeOther.NilReason {
		diffs = append(diffs, codegradeseparationtype.GongMarshallField(stage, "NilReason"))
	}
	if codegradeseparationtype.EnclosedText != codegradeseparationtypeOther.EnclosedText {
		diffs = append(diffs, codegradeseparationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codegroundlightingtype *CodeGroundLightingType) GongDiff(stage *Stage, codegroundlightingtypeOther *CodeGroundLightingType) (diffs []string) {
	// insertion point for field diffs
	if codegroundlightingtype.Name != codegroundlightingtypeOther.Name {
		diffs = append(diffs, codegroundlightingtype.GongMarshallField(stage, "Name"))
	}
	if codegroundlightingtype.NilReason != codegroundlightingtypeOther.NilReason {
		diffs = append(diffs, codegroundlightingtype.GongMarshallField(stage, "NilReason"))
	}
	if codegroundlightingtype.EnclosedText != codegroundlightingtypeOther.EnclosedText {
		diffs = append(diffs, codegroundlightingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeguidancelinetype *CodeGuidanceLineType) GongDiff(stage *Stage, codeguidancelinetypeOther *CodeGuidanceLineType) (diffs []string) {
	// insertion point for field diffs
	if codeguidancelinetype.Name != codeguidancelinetypeOther.Name {
		diffs = append(diffs, codeguidancelinetype.GongMarshallField(stage, "Name"))
	}
	if codeguidancelinetype.NilReason != codeguidancelinetypeOther.NilReason {
		diffs = append(diffs, codeguidancelinetype.GongMarshallField(stage, "NilReason"))
	}
	if codeguidancelinetype.EnclosedText != codeguidancelinetypeOther.EnclosedText {
		diffs = append(diffs, codeguidancelinetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeheightreferencetype *CodeHeightReferenceType) GongDiff(stage *Stage, codeheightreferencetypeOther *CodeHeightReferenceType) (diffs []string) {
	// insertion point for field diffs
	if codeheightreferencetype.Name != codeheightreferencetypeOther.Name {
		diffs = append(diffs, codeheightreferencetype.GongMarshallField(stage, "Name"))
	}
	if codeheightreferencetype.NilReason != codeheightreferencetypeOther.NilReason {
		diffs = append(diffs, codeheightreferencetype.GongMarshallField(stage, "NilReason"))
	}
	if codeheightreferencetype.EnclosedText != codeheightreferencetypeOther.EnclosedText {
		diffs = append(diffs, codeheightreferencetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongDiff(stage *Stage, codehelicopterperformancetypeOther *CodeHelicopterPerformanceType) (diffs []string) {
	// insertion point for field diffs
	if codehelicopterperformancetype.Name != codehelicopterperformancetypeOther.Name {
		diffs = append(diffs, codehelicopterperformancetype.GongMarshallField(stage, "Name"))
	}
	if codehelicopterperformancetype.NilReason != codehelicopterperformancetypeOther.NilReason {
		diffs = append(diffs, codehelicopterperformancetype.GongMarshallField(stage, "NilReason"))
	}
	if codehelicopterperformancetype.EnclosedText != codehelicopterperformancetypeOther.EnclosedText {
		diffs = append(diffs, codehelicopterperformancetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeholdingcategorytype *CodeHoldingCategoryType) GongDiff(stage *Stage, codeholdingcategorytypeOther *CodeHoldingCategoryType) (diffs []string) {
	// insertion point for field diffs
	if codeholdingcategorytype.Name != codeholdingcategorytypeOther.Name {
		diffs = append(diffs, codeholdingcategorytype.GongMarshallField(stage, "Name"))
	}
	if codeholdingcategorytype.NilReason != codeholdingcategorytypeOther.NilReason {
		diffs = append(diffs, codeholdingcategorytype.GongMarshallField(stage, "NilReason"))
	}
	if codeholdingcategorytype.EnclosedText != codeholdingcategorytypeOther.EnclosedText {
		diffs = append(diffs, codeholdingcategorytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeholdingusagetype *CodeHoldingUsageType) GongDiff(stage *Stage, codeholdingusagetypeOther *CodeHoldingUsageType) (diffs []string) {
	// insertion point for field diffs
	if codeholdingusagetype.Name != codeholdingusagetypeOther.Name {
		diffs = append(diffs, codeholdingusagetype.GongMarshallField(stage, "Name"))
	}
	if codeholdingusagetype.NilReason != codeholdingusagetypeOther.NilReason {
		diffs = append(diffs, codeholdingusagetype.GongMarshallField(stage, "NilReason"))
	}
	if codeholdingusagetype.EnclosedText != codeholdingusagetypeOther.EnclosedText {
		diffs = append(diffs, codeholdingusagetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeholdingusetype *CodeHoldingUseType) GongDiff(stage *Stage, codeholdingusetypeOther *CodeHoldingUseType) (diffs []string) {
	// insertion point for field diffs
	if codeholdingusetype.Name != codeholdingusetypeOther.Name {
		diffs = append(diffs, codeholdingusetype.GongMarshallField(stage, "Name"))
	}
	if codeholdingusetype.NilReason != codeholdingusetypeOther.NilReason {
		diffs = append(diffs, codeholdingusetype.GongMarshallField(stage, "NilReason"))
	}
	if codeholdingusetype.EnclosedText != codeholdingusetypeOther.EnclosedText {
		diffs = append(diffs, codeholdingusetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeiatatype *CodeIATAType) GongDiff(stage *Stage, codeiatatypeOther *CodeIATAType) (diffs []string) {
	// insertion point for field diffs
	if codeiatatype.Name != codeiatatypeOther.Name {
		diffs = append(diffs, codeiatatype.GongMarshallField(stage, "Name"))
	}
	if codeiatatype.NilReason != codeiatatypeOther.NilReason {
		diffs = append(diffs, codeiatatype.GongMarshallField(stage, "NilReason"))
	}
	if codeiatatype.EnclosedText != codeiatatypeOther.EnclosedText {
		diffs = append(diffs, codeiatatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeicaocountrytype *CodeICAOCountryType) GongDiff(stage *Stage, codeicaocountrytypeOther *CodeICAOCountryType) (diffs []string) {
	// insertion point for field diffs
	if codeicaocountrytype.Name != codeicaocountrytypeOther.Name {
		diffs = append(diffs, codeicaocountrytype.GongMarshallField(stage, "Name"))
	}
	if codeicaocountrytype.NilReason != codeicaocountrytypeOther.NilReason {
		diffs = append(diffs, codeicaocountrytype.GongMarshallField(stage, "NilReason"))
	}
	if codeicaocountrytype.EnclosedText != codeicaocountrytypeOther.EnclosedText {
		diffs = append(diffs, codeicaocountrytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeicaotype *CodeICAOType) GongDiff(stage *Stage, codeicaotypeOther *CodeICAOType) (diffs []string) {
	// insertion point for field diffs
	if codeicaotype.Name != codeicaotypeOther.Name {
		diffs = append(diffs, codeicaotype.GongMarshallField(stage, "Name"))
	}
	if codeicaotype.NilReason != codeicaotypeOther.NilReason {
		diffs = append(diffs, codeicaotype.GongMarshallField(stage, "NilReason"))
	}
	if codeicaotype.EnclosedText != codeicaotypeOther.EnclosedText {
		diffs = append(diffs, codeicaotype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeilsbackcoursetype *CodeILSBackCourseType) GongDiff(stage *Stage, codeilsbackcoursetypeOther *CodeILSBackCourseType) (diffs []string) {
	// insertion point for field diffs
	if codeilsbackcoursetype.Name != codeilsbackcoursetypeOther.Name {
		diffs = append(diffs, codeilsbackcoursetype.GongMarshallField(stage, "Name"))
	}
	if codeilsbackcoursetype.NilReason != codeilsbackcoursetypeOther.NilReason {
		diffs = append(diffs, codeilsbackcoursetype.GongMarshallField(stage, "NilReason"))
	}
	if codeilsbackcoursetype.EnclosedText != codeilsbackcoursetypeOther.EnclosedText {
		diffs = append(diffs, codeilsbackcoursetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongDiff(stage *Stage, codeintegritylevelilstypeOther *CodeIntegrityLevelILSType) (diffs []string) {
	// insertion point for field diffs
	if codeintegritylevelilstype.Name != codeintegritylevelilstypeOther.Name {
		diffs = append(diffs, codeintegritylevelilstype.GongMarshallField(stage, "Name"))
	}
	if codeintegritylevelilstype.NilReason != codeintegritylevelilstypeOther.NilReason {
		diffs = append(diffs, codeintegritylevelilstype.GongMarshallField(stage, "NilReason"))
	}
	if codeintegritylevelilstype.EnclosedText != codeintegritylevelilstypeOther.EnclosedText {
		diffs = append(diffs, codeintegritylevelilstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeintensitystandbytype *CodeIntensityStandByType) GongDiff(stage *Stage, codeintensitystandbytypeOther *CodeIntensityStandByType) (diffs []string) {
	// insertion point for field diffs
	if codeintensitystandbytype.Name != codeintensitystandbytypeOther.Name {
		diffs = append(diffs, codeintensitystandbytype.GongMarshallField(stage, "Name"))
	}
	if codeintensitystandbytype.NilReason != codeintensitystandbytypeOther.NilReason {
		diffs = append(diffs, codeintensitystandbytype.GongMarshallField(stage, "NilReason"))
	}
	if codeintensitystandbytype.EnclosedText != codeintensitystandbytypeOther.EnclosedText {
		diffs = append(diffs, codeintensitystandbytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelanguagetype *CodeLanguageType) GongDiff(stage *Stage, codelanguagetypeOther *CodeLanguageType) (diffs []string) {
	// insertion point for field diffs
	if codelanguagetype.Name != codelanguagetypeOther.Name {
		diffs = append(diffs, codelanguagetype.GongMarshallField(stage, "Name"))
	}
	if codelanguagetype.NilReason != codelanguagetypeOther.NilReason {
		diffs = append(diffs, codelanguagetype.GongMarshallField(stage, "NilReason"))
	}
	if codelanguagetype.EnclosedText != codelanguagetypeOther.EnclosedText {
		diffs = append(diffs, codelanguagetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelevelseriestype *CodeLevelSeriesType) GongDiff(stage *Stage, codelevelseriestypeOther *CodeLevelSeriesType) (diffs []string) {
	// insertion point for field diffs
	if codelevelseriestype.Name != codelevelseriestypeOther.Name {
		diffs = append(diffs, codelevelseriestype.GongMarshallField(stage, "Name"))
	}
	if codelevelseriestype.NilReason != codelevelseriestypeOther.NilReason {
		diffs = append(diffs, codelevelseriestype.GongMarshallField(stage, "NilReason"))
	}
	if codelevelseriestype.EnclosedText != codelevelseriestypeOther.EnclosedText {
		diffs = append(diffs, codelevelseriestype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongDiff(stage *Stage, codeleveltabledesignatortypeOther *CodeLevelTableDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codeleveltabledesignatortype.Name != codeleveltabledesignatortypeOther.Name {
		diffs = append(diffs, codeleveltabledesignatortype.GongMarshallField(stage, "Name"))
	}
	if codeleveltabledesignatortype.NilReason != codeleveltabledesignatortypeOther.NilReason {
		diffs = append(diffs, codeleveltabledesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codeleveltabledesignatortype.EnclosedText != codeleveltabledesignatortypeOther.EnclosedText {
		diffs = append(diffs, codeleveltabledesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeleveltype *CodeLevelType) GongDiff(stage *Stage, codeleveltypeOther *CodeLevelType) (diffs []string) {
	// insertion point for field diffs
	if codeleveltype.Name != codeleveltypeOther.Name {
		diffs = append(diffs, codeleveltype.GongMarshallField(stage, "Name"))
	}
	if codeleveltype.NilReason != codeleveltypeOther.NilReason {
		diffs = append(diffs, codeleveltype.GongMarshallField(stage, "NilReason"))
	}
	if codeleveltype.EnclosedText != codeleveltypeOther.EnclosedText {
		diffs = append(diffs, codeleveltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongDiff(stage *Stage, codelightholdingpositiontypeOther *CodeLightHoldingPositionType) (diffs []string) {
	// insertion point for field diffs
	if codelightholdingpositiontype.Name != codelightholdingpositiontypeOther.Name {
		diffs = append(diffs, codelightholdingpositiontype.GongMarshallField(stage, "Name"))
	}
	if codelightholdingpositiontype.NilReason != codelightholdingpositiontypeOther.NilReason {
		diffs = append(diffs, codelightholdingpositiontype.GongMarshallField(stage, "NilReason"))
	}
	if codelightholdingpositiontype.EnclosedText != codelightholdingpositiontypeOther.EnclosedText {
		diffs = append(diffs, codelightholdingpositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelightintensitytype *CodeLightIntensityType) GongDiff(stage *Stage, codelightintensitytypeOther *CodeLightIntensityType) (diffs []string) {
	// insertion point for field diffs
	if codelightintensitytype.Name != codelightintensitytypeOther.Name {
		diffs = append(diffs, codelightintensitytype.GongMarshallField(stage, "Name"))
	}
	if codelightintensitytype.NilReason != codelightintensitytypeOther.NilReason {
		diffs = append(diffs, codelightintensitytype.GongMarshallField(stage, "NilReason"))
	}
	if codelightintensitytype.EnclosedText != codelightintensitytypeOther.EnclosedText {
		diffs = append(diffs, codelightintensitytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelightsourcetype *CodeLightSourceType) GongDiff(stage *Stage, codelightsourcetypeOther *CodeLightSourceType) (diffs []string) {
	// insertion point for field diffs
	if codelightsourcetype.Name != codelightsourcetypeOther.Name {
		diffs = append(diffs, codelightsourcetype.GongMarshallField(stage, "Name"))
	}
	if codelightsourcetype.NilReason != codelightsourcetypeOther.NilReason {
		diffs = append(diffs, codelightsourcetype.GongMarshallField(stage, "NilReason"))
	}
	if codelightsourcetype.EnclosedText != codelightsourcetypeOther.EnclosedText {
		diffs = append(diffs, codelightsourcetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelightingjartype *CodeLightingJARType) GongDiff(stage *Stage, codelightingjartypeOther *CodeLightingJARType) (diffs []string) {
	// insertion point for field diffs
	if codelightingjartype.Name != codelightingjartypeOther.Name {
		diffs = append(diffs, codelightingjartype.GongMarshallField(stage, "Name"))
	}
	if codelightingjartype.NilReason != codelightingjartypeOther.NilReason {
		diffs = append(diffs, codelightingjartype.GongMarshallField(stage, "NilReason"))
	}
	if codelightingjartype.EnclosedText != codelightingjartypeOther.EnclosedText {
		diffs = append(diffs, codelightingjartype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeloadingbridgetype *CodeLoadingBridgeType) GongDiff(stage *Stage, codeloadingbridgetypeOther *CodeLoadingBridgeType) (diffs []string) {
	// insertion point for field diffs
	if codeloadingbridgetype.Name != codeloadingbridgetypeOther.Name {
		diffs = append(diffs, codeloadingbridgetype.GongMarshallField(stage, "Name"))
	}
	if codeloadingbridgetype.NilReason != codeloadingbridgetypeOther.NilReason {
		diffs = append(diffs, codeloadingbridgetype.GongMarshallField(stage, "NilReason"))
	}
	if codeloadingbridgetype.EnclosedText != codeloadingbridgetypeOther.EnclosedText {
		diffs = append(diffs, codeloadingbridgetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelocationqualifiertype *CodeLocationQualifierType) GongDiff(stage *Stage, codelocationqualifiertypeOther *CodeLocationQualifierType) (diffs []string) {
	// insertion point for field diffs
	if codelocationqualifiertype.Name != codelocationqualifiertypeOther.Name {
		diffs = append(diffs, codelocationqualifiertype.GongMarshallField(stage, "Name"))
	}
	if codelocationqualifiertype.NilReason != codelocationqualifiertypeOther.NilReason {
		diffs = append(diffs, codelocationqualifiertype.GongMarshallField(stage, "NilReason"))
	}
	if codelocationqualifiertype.EnclosedText != codelocationqualifiertypeOther.EnclosedText {
		diffs = append(diffs, codelocationqualifiertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelogicaloperatortype *CodeLogicalOperatorType) GongDiff(stage *Stage, codelogicaloperatortypeOther *CodeLogicalOperatorType) (diffs []string) {
	// insertion point for field diffs
	if codelogicaloperatortype.Name != codelogicaloperatortypeOther.Name {
		diffs = append(diffs, codelogicaloperatortype.GongMarshallField(stage, "Name"))
	}
	if codelogicaloperatortype.NilReason != codelogicaloperatortypeOther.NilReason {
		diffs = append(diffs, codelogicaloperatortype.GongMarshallField(stage, "NilReason"))
	}
	if codelogicaloperatortype.EnclosedText != codelogicaloperatortypeOther.EnclosedText {
		diffs = append(diffs, codelogicaloperatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemlsazimuthtype *CodeMLSAzimuthType) GongDiff(stage *Stage, codemlsazimuthtypeOther *CodeMLSAzimuthType) (diffs []string) {
	// insertion point for field diffs
	if codemlsazimuthtype.Name != codemlsazimuthtypeOther.Name {
		diffs = append(diffs, codemlsazimuthtype.GongMarshallField(stage, "Name"))
	}
	if codemlsazimuthtype.NilReason != codemlsazimuthtypeOther.NilReason {
		diffs = append(diffs, codemlsazimuthtype.GongMarshallField(stage, "NilReason"))
	}
	if codemlsazimuthtype.EnclosedText != codemlsazimuthtypeOther.EnclosedText {
		diffs = append(diffs, codemlsazimuthtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemlschanneltype *CodeMLSChannelType) GongDiff(stage *Stage, codemlschanneltypeOther *CodeMLSChannelType) (diffs []string) {
	// insertion point for field diffs
	if codemlschanneltype.Name != codemlschanneltypeOther.Name {
		diffs = append(diffs, codemlschanneltype.GongMarshallField(stage, "Name"))
	}
	if codemlschanneltype.NilReason != codemlschanneltypeOther.NilReason {
		diffs = append(diffs, codemlschanneltype.GongMarshallField(stage, "NilReason"))
	}
	if codemlschanneltype.EnclosedText != codemlschanneltypeOther.EnclosedText {
		diffs = append(diffs, codemlschanneltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongDiff(stage *Stage, codemarkerbeaconsignaltypeOther *CodeMarkerBeaconSignalType) (diffs []string) {
	// insertion point for field diffs
	if codemarkerbeaconsignaltype.Name != codemarkerbeaconsignaltypeOther.Name {
		diffs = append(diffs, codemarkerbeaconsignaltype.GongMarshallField(stage, "Name"))
	}
	if codemarkerbeaconsignaltype.NilReason != codemarkerbeaconsignaltypeOther.NilReason {
		diffs = append(diffs, codemarkerbeaconsignaltype.GongMarshallField(stage, "NilReason"))
	}
	if codemarkerbeaconsignaltype.EnclosedText != codemarkerbeaconsignaltypeOther.EnclosedText {
		diffs = append(diffs, codemarkerbeaconsignaltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemarkingconditiontype *CodeMarkingConditionType) GongDiff(stage *Stage, codemarkingconditiontypeOther *CodeMarkingConditionType) (diffs []string) {
	// insertion point for field diffs
	if codemarkingconditiontype.Name != codemarkingconditiontypeOther.Name {
		diffs = append(diffs, codemarkingconditiontype.GongMarshallField(stage, "Name"))
	}
	if codemarkingconditiontype.NilReason != codemarkingconditiontypeOther.NilReason {
		diffs = append(diffs, codemarkingconditiontype.GongMarshallField(stage, "NilReason"))
	}
	if codemarkingconditiontype.EnclosedText != codemarkingconditiontypeOther.EnclosedText {
		diffs = append(diffs, codemarkingconditiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemarkingstyletype *CodeMarkingStyleType) GongDiff(stage *Stage, codemarkingstyletypeOther *CodeMarkingStyleType) (diffs []string) {
	// insertion point for field diffs
	if codemarkingstyletype.Name != codemarkingstyletypeOther.Name {
		diffs = append(diffs, codemarkingstyletype.GongMarshallField(stage, "Name"))
	}
	if codemarkingstyletype.NilReason != codemarkingstyletypeOther.NilReason {
		diffs = append(diffs, codemarkingstyletype.GongMarshallField(stage, "NilReason"))
	}
	if codemarkingstyletype.EnclosedText != codemarkingstyletypeOther.EnclosedText {
		diffs = append(diffs, codemarkingstyletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemeteoconditionstype *CodeMeteoConditionsType) GongDiff(stage *Stage, codemeteoconditionstypeOther *CodeMeteoConditionsType) (diffs []string) {
	// insertion point for field diffs
	if codemeteoconditionstype.Name != codemeteoconditionstypeOther.Name {
		diffs = append(diffs, codemeteoconditionstype.GongMarshallField(stage, "Name"))
	}
	if codemeteoconditionstype.NilReason != codemeteoconditionstypeOther.NilReason {
		diffs = append(diffs, codemeteoconditionstype.GongMarshallField(stage, "NilReason"))
	}
	if codemeteoconditionstype.EnclosedText != codemeteoconditionstypeOther.EnclosedText {
		diffs = append(diffs, codemeteoconditionstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongDiff(stage *Stage, codemilitaryoperationstypeOther *CodeMilitaryOperationsType) (diffs []string) {
	// insertion point for field diffs
	if codemilitaryoperationstype.Name != codemilitaryoperationstypeOther.Name {
		diffs = append(diffs, codemilitaryoperationstype.GongMarshallField(stage, "Name"))
	}
	if codemilitaryoperationstype.NilReason != codemilitaryoperationstypeOther.NilReason {
		diffs = append(diffs, codemilitaryoperationstype.GongMarshallField(stage, "NilReason"))
	}
	if codemilitaryoperationstype.EnclosedText != codemilitaryoperationstypeOther.EnclosedText {
		diffs = append(diffs, codemilitaryoperationstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongDiff(stage *Stage, codemilitaryroutepointtypeOther *CodeMilitaryRoutePointType) (diffs []string) {
	// insertion point for field diffs
	if codemilitaryroutepointtype.Name != codemilitaryroutepointtypeOther.Name {
		diffs = append(diffs, codemilitaryroutepointtype.GongMarshallField(stage, "Name"))
	}
	if codemilitaryroutepointtype.NilReason != codemilitaryroutepointtypeOther.NilReason {
		diffs = append(diffs, codemilitaryroutepointtype.GongMarshallField(stage, "NilReason"))
	}
	if codemilitaryroutepointtype.EnclosedText != codemilitaryroutepointtypeOther.EnclosedText {
		diffs = append(diffs, codemilitaryroutepointtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemilitarystatustype *CodeMilitaryStatusType) GongDiff(stage *Stage, codemilitarystatustypeOther *CodeMilitaryStatusType) (diffs []string) {
	// insertion point for field diffs
	if codemilitarystatustype.Name != codemilitarystatustypeOther.Name {
		diffs = append(diffs, codemilitarystatustype.GongMarshallField(stage, "Name"))
	}
	if codemilitarystatustype.NilReason != codemilitarystatustypeOther.NilReason {
		diffs = append(diffs, codemilitarystatustype.GongMarshallField(stage, "NilReason"))
	}
	if codemilitarystatustype.EnclosedText != codemilitarystatustypeOther.EnclosedText {
		diffs = append(diffs, codemilitarystatustype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongDiff(stage *Stage, codemilitarytrainingtypeOther *CodeMilitaryTrainingType) (diffs []string) {
	// insertion point for field diffs
	if codemilitarytrainingtype.Name != codemilitarytrainingtypeOther.Name {
		diffs = append(diffs, codemilitarytrainingtype.GongMarshallField(stage, "Name"))
	}
	if codemilitarytrainingtype.NilReason != codemilitarytrainingtypeOther.NilReason {
		diffs = append(diffs, codemilitarytrainingtype.GongMarshallField(stage, "NilReason"))
	}
	if codemilitarytrainingtype.EnclosedText != codemilitarytrainingtypeOther.EnclosedText {
		diffs = append(diffs, codemilitarytrainingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongDiff(stage *Stage, codeminimafinalapproachpathtypeOther *CodeMinimaFinalApproachPathType) (diffs []string) {
	// insertion point for field diffs
	if codeminimafinalapproachpathtype.Name != codeminimafinalapproachpathtypeOther.Name {
		diffs = append(diffs, codeminimafinalapproachpathtype.GongMarshallField(stage, "Name"))
	}
	if codeminimafinalapproachpathtype.NilReason != codeminimafinalapproachpathtypeOther.NilReason {
		diffs = append(diffs, codeminimafinalapproachpathtype.GongMarshallField(stage, "NilReason"))
	}
	if codeminimafinalapproachpathtype.EnclosedText != codeminimafinalapproachpathtypeOther.EnclosedText {
		diffs = append(diffs, codeminimafinalapproachpathtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongDiff(stage *Stage, codeminimumaltitudetypeOther *CodeMinimumAltitudeType) (diffs []string) {
	// insertion point for field diffs
	if codeminimumaltitudetype.Name != codeminimumaltitudetypeOther.Name {
		diffs = append(diffs, codeminimumaltitudetype.GongMarshallField(stage, "Name"))
	}
	if codeminimumaltitudetype.NilReason != codeminimumaltitudetypeOther.NilReason {
		diffs = append(diffs, codeminimumaltitudetype.GongMarshallField(stage, "NilReason"))
	}
	if codeminimumaltitudetype.EnclosedText != codeminimumaltitudetypeOther.EnclosedText {
		diffs = append(diffs, codeminimumaltitudetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeminimumheighttype *CodeMinimumHeightType) GongDiff(stage *Stage, codeminimumheighttypeOther *CodeMinimumHeightType) (diffs []string) {
	// insertion point for field diffs
	if codeminimumheighttype.Name != codeminimumheighttypeOther.Name {
		diffs = append(diffs, codeminimumheighttype.GongMarshallField(stage, "Name"))
	}
	if codeminimumheighttype.NilReason != codeminimumheighttypeOther.NilReason {
		diffs = append(diffs, codeminimumheighttype.GongMarshallField(stage, "NilReason"))
	}
	if codeminimumheighttype.EnclosedText != codeminimumheighttypeOther.EnclosedText {
		diffs = append(diffs, codeminimumheighttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemissedapproachtype *CodeMissedApproachType) GongDiff(stage *Stage, codemissedapproachtypeOther *CodeMissedApproachType) (diffs []string) {
	// insertion point for field diffs
	if codemissedapproachtype.Name != codemissedapproachtypeOther.Name {
		diffs = append(diffs, codemissedapproachtype.GongMarshallField(stage, "Name"))
	}
	if codemissedapproachtype.NilReason != codemissedapproachtypeOther.NilReason {
		diffs = append(diffs, codemissedapproachtype.GongMarshallField(stage, "NilReason"))
	}
	if codemissedapproachtype.EnclosedText != codemissedapproachtypeOther.EnclosedText {
		diffs = append(diffs, codemissedapproachtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codendbusagetype *CodeNDBUsageType) GongDiff(stage *Stage, codendbusagetypeOther *CodeNDBUsageType) (diffs []string) {
	// insertion point for field diffs
	if codendbusagetype.Name != codendbusagetypeOther.Name {
		diffs = append(diffs, codendbusagetype.GongMarshallField(stage, "Name"))
	}
	if codendbusagetype.NilReason != codendbusagetypeOther.NilReason {
		diffs = append(diffs, codendbusagetype.GongMarshallField(stage, "NilReason"))
	}
	if codendbusagetype.EnclosedText != codendbusagetypeOther.EnclosedText {
		diffs = append(diffs, codendbusagetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongDiff(stage *Stage, codenavaiddesignatortypeOther *CodeNavaidDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codenavaiddesignatortype.Name != codenavaiddesignatortypeOther.Name {
		diffs = append(diffs, codenavaiddesignatortype.GongMarshallField(stage, "Name"))
	}
	if codenavaiddesignatortype.NilReason != codenavaiddesignatortypeOther.NilReason {
		diffs = append(diffs, codenavaiddesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codenavaiddesignatortype.EnclosedText != codenavaiddesignatortypeOther.EnclosedText {
		diffs = append(diffs, codenavaiddesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavaidpurposetype *CodeNavaidPurposeType) GongDiff(stage *Stage, codenavaidpurposetypeOther *CodeNavaidPurposeType) (diffs []string) {
	// insertion point for field diffs
	if codenavaidpurposetype.Name != codenavaidpurposetypeOther.Name {
		diffs = append(diffs, codenavaidpurposetype.GongMarshallField(stage, "Name"))
	}
	if codenavaidpurposetype.NilReason != codenavaidpurposetypeOther.NilReason {
		diffs = append(diffs, codenavaidpurposetype.GongMarshallField(stage, "NilReason"))
	}
	if codenavaidpurposetype.EnclosedText != codenavaidpurposetypeOther.EnclosedText {
		diffs = append(diffs, codenavaidpurposetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavaidservicetype *CodeNavaidServiceType) GongDiff(stage *Stage, codenavaidservicetypeOther *CodeNavaidServiceType) (diffs []string) {
	// insertion point for field diffs
	if codenavaidservicetype.Name != codenavaidservicetypeOther.Name {
		diffs = append(diffs, codenavaidservicetype.GongMarshallField(stage, "Name"))
	}
	if codenavaidservicetype.NilReason != codenavaidservicetypeOther.NilReason {
		diffs = append(diffs, codenavaidservicetype.GongMarshallField(stage, "NilReason"))
	}
	if codenavaidservicetype.EnclosedText != codenavaidservicetypeOther.EnclosedText {
		diffs = append(diffs, codenavaidservicetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongDiff(stage *Stage, codenavigationarearestrictiontypeOther *CodeNavigationAreaRestrictionType) (diffs []string) {
	// insertion point for field diffs
	if codenavigationarearestrictiontype.Name != codenavigationarearestrictiontypeOther.Name {
		diffs = append(diffs, codenavigationarearestrictiontype.GongMarshallField(stage, "Name"))
	}
	if codenavigationarearestrictiontype.NilReason != codenavigationarearestrictiontypeOther.NilReason {
		diffs = append(diffs, codenavigationarearestrictiontype.GongMarshallField(stage, "NilReason"))
	}
	if codenavigationarearestrictiontype.EnclosedText != codenavigationarearestrictiontypeOther.EnclosedText {
		diffs = append(diffs, codenavigationarearestrictiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavigationareatype *CodeNavigationAreaType) GongDiff(stage *Stage, codenavigationareatypeOther *CodeNavigationAreaType) (diffs []string) {
	// insertion point for field diffs
	if codenavigationareatype.Name != codenavigationareatypeOther.Name {
		diffs = append(diffs, codenavigationareatype.GongMarshallField(stage, "Name"))
	}
	if codenavigationareatype.NilReason != codenavigationareatypeOther.NilReason {
		diffs = append(diffs, codenavigationareatype.GongMarshallField(stage, "NilReason"))
	}
	if codenavigationareatype.EnclosedText != codenavigationareatypeOther.EnclosedText {
		diffs = append(diffs, codenavigationareatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongDiff(stage *Stage, codenavigationequipmenttypeOther *CodeNavigationEquipmentType) (diffs []string) {
	// insertion point for field diffs
	if codenavigationequipmenttype.Name != codenavigationequipmenttypeOther.Name {
		diffs = append(diffs, codenavigationequipmenttype.GongMarshallField(stage, "Name"))
	}
	if codenavigationequipmenttype.NilReason != codenavigationequipmenttypeOther.NilReason {
		diffs = append(diffs, codenavigationequipmenttype.GongMarshallField(stage, "NilReason"))
	}
	if codenavigationequipmenttype.EnclosedText != codenavigationequipmenttypeOther.EnclosedText {
		diffs = append(diffs, codenavigationequipmenttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongDiff(stage *Stage, codenavigationspecificationtypeOther *CodeNavigationSpecificationType) (diffs []string) {
	// insertion point for field diffs
	if codenavigationspecificationtype.Name != codenavigationspecificationtypeOther.Name {
		diffs = append(diffs, codenavigationspecificationtype.GongMarshallField(stage, "Name"))
	}
	if codenavigationspecificationtype.NilReason != codenavigationspecificationtypeOther.NilReason {
		diffs = append(diffs, codenavigationspecificationtype.GongMarshallField(stage, "NilReason"))
	}
	if codenavigationspecificationtype.EnclosedText != codenavigationspecificationtypeOther.EnclosedText {
		diffs = append(diffs, codenavigationspecificationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenitrogentype *CodeNitrogenType) GongDiff(stage *Stage, codenitrogentypeOther *CodeNitrogenType) (diffs []string) {
	// insertion point for field diffs
	if codenitrogentype.Name != codenitrogentypeOther.Name {
		diffs = append(diffs, codenitrogentype.GongMarshallField(stage, "Name"))
	}
	if codenitrogentype.NilReason != codenitrogentypeOther.NilReason {
		diffs = append(diffs, codenitrogentype.GongMarshallField(stage, "NilReason"))
	}
	if codenitrogentype.EnclosedText != codenitrogentypeOther.EnclosedText {
		diffs = append(diffs, codenitrogentype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenorthreferencetype *CodeNorthReferenceType) GongDiff(stage *Stage, codenorthreferencetypeOther *CodeNorthReferenceType) (diffs []string) {
	// insertion point for field diffs
	if codenorthreferencetype.Name != codenorthreferencetypeOther.Name {
		diffs = append(diffs, codenorthreferencetype.GongMarshallField(stage, "Name"))
	}
	if codenorthreferencetype.NilReason != codenorthreferencetypeOther.NilReason {
		diffs = append(diffs, codenorthreferencetype.GongMarshallField(stage, "NilReason"))
	}
	if codenorthreferencetype.EnclosedText != codenorthreferencetypeOther.EnclosedText {
		diffs = append(diffs, codenorthreferencetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenotepurposetype *CodeNotePurposeType) GongDiff(stage *Stage, codenotepurposetypeOther *CodeNotePurposeType) (diffs []string) {
	// insertion point for field diffs
	if codenotepurposetype.Name != codenotepurposetypeOther.Name {
		diffs = append(diffs, codenotepurposetype.GongMarshallField(stage, "Name"))
	}
	if codenotepurposetype.NilReason != codenotepurposetypeOther.NilReason {
		diffs = append(diffs, codenotepurposetype.GongMarshallField(stage, "NilReason"))
	}
	if codenotepurposetype.EnclosedText != codenotepurposetypeOther.EnclosedText {
		diffs = append(diffs, codenotepurposetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeobstacleareatype *CodeObstacleAreaType) GongDiff(stage *Stage, codeobstacleareatypeOther *CodeObstacleAreaType) (diffs []string) {
	// insertion point for field diffs
	if codeobstacleareatype.Name != codeobstacleareatypeOther.Name {
		diffs = append(diffs, codeobstacleareatype.GongMarshallField(stage, "Name"))
	}
	if codeobstacleareatype.NilReason != codeobstacleareatypeOther.NilReason {
		diffs = append(diffs, codeobstacleareatype.GongMarshallField(stage, "NilReason"))
	}
	if codeobstacleareatype.EnclosedText != codeobstacleareatypeOther.EnclosedText {
		diffs = append(diffs, codeobstacleareatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongDiff(stage *Stage, codeobstacleassessmentsurfacetypeOther *CodeObstacleAssessmentSurfaceType) (diffs []string) {
	// insertion point for field diffs
	if codeobstacleassessmentsurfacetype.Name != codeobstacleassessmentsurfacetypeOther.Name {
		diffs = append(diffs, codeobstacleassessmentsurfacetype.GongMarshallField(stage, "Name"))
	}
	if codeobstacleassessmentsurfacetype.NilReason != codeobstacleassessmentsurfacetypeOther.NilReason {
		diffs = append(diffs, codeobstacleassessmentsurfacetype.GongMarshallField(stage, "NilReason"))
	}
	if codeobstacleassessmentsurfacetype.EnclosedText != codeobstacleassessmentsurfacetypeOther.EnclosedText {
		diffs = append(diffs, codeobstacleassessmentsurfacetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongDiff(stage *Stage, codeobstructionidsurfacezonetypeOther *CodeObstructionIdSurfaceZoneType) (diffs []string) {
	// insertion point for field diffs
	if codeobstructionidsurfacezonetype.Name != codeobstructionidsurfacezonetypeOther.Name {
		diffs = append(diffs, codeobstructionidsurfacezonetype.GongMarshallField(stage, "Name"))
	}
	if codeobstructionidsurfacezonetype.NilReason != codeobstructionidsurfacezonetypeOther.NilReason {
		diffs = append(diffs, codeobstructionidsurfacezonetype.GongMarshallField(stage, "NilReason"))
	}
	if codeobstructionidsurfacezonetype.EnclosedText != codeobstructionidsurfacezonetypeOther.EnclosedText {
		diffs = append(diffs, codeobstructionidsurfacezonetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeoiltype *CodeOilType) GongDiff(stage *Stage, codeoiltypeOther *CodeOilType) (diffs []string) {
	// insertion point for field diffs
	if codeoiltype.Name != codeoiltypeOther.Name {
		diffs = append(diffs, codeoiltype.GongMarshallField(stage, "Name"))
	}
	if codeoiltype.NilReason != codeoiltypeOther.NilReason {
		diffs = append(diffs, codeoiltype.GongMarshallField(stage, "NilReason"))
	}
	if codeoiltype.EnclosedText != codeoiltypeOther.EnclosedText {
		diffs = append(diffs, codeoiltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongDiff(stage *Stage, codeoperationairportheliporttypeOther *CodeOperationAirportHeliportType) (diffs []string) {
	// insertion point for field diffs
	if codeoperationairportheliporttype.Name != codeoperationairportheliporttypeOther.Name {
		diffs = append(diffs, codeoperationairportheliporttype.GongMarshallField(stage, "Name"))
	}
	if codeoperationairportheliporttype.NilReason != codeoperationairportheliporttypeOther.NilReason {
		diffs = append(diffs, codeoperationairportheliporttype.GongMarshallField(stage, "NilReason"))
	}
	if codeoperationairportheliporttype.EnclosedText != codeoperationairportheliporttypeOther.EnclosedText {
		diffs = append(diffs, codeoperationairportheliporttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongDiff(stage *Stage, codeoperationmanoeuvringareatypeOther *CodeOperationManoeuvringAreaType) (diffs []string) {
	// insertion point for field diffs
	if codeoperationmanoeuvringareatype.Name != codeoperationmanoeuvringareatypeOther.Name {
		diffs = append(diffs, codeoperationmanoeuvringareatype.GongMarshallField(stage, "Name"))
	}
	if codeoperationmanoeuvringareatype.NilReason != codeoperationmanoeuvringareatypeOther.NilReason {
		diffs = append(diffs, codeoperationmanoeuvringareatype.GongMarshallField(stage, "NilReason"))
	}
	if codeoperationmanoeuvringareatype.EnclosedText != codeoperationmanoeuvringareatypeOther.EnclosedText {
		diffs = append(diffs, codeoperationmanoeuvringareatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongDiff(stage *Stage, codeorganisationdesignatortypeOther *CodeOrganisationDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codeorganisationdesignatortype.Name != codeorganisationdesignatortypeOther.Name {
		diffs = append(diffs, codeorganisationdesignatortype.GongMarshallField(stage, "Name"))
	}
	if codeorganisationdesignatortype.NilReason != codeorganisationdesignatortypeOther.NilReason {
		diffs = append(diffs, codeorganisationdesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codeorganisationdesignatortype.EnclosedText != codeorganisationdesignatortypeOther.EnclosedText {
		diffs = append(diffs, codeorganisationdesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongDiff(stage *Stage, codeorganisationhierarchytypeOther *CodeOrganisationHierarchyType) (diffs []string) {
	// insertion point for field diffs
	if codeorganisationhierarchytype.Name != codeorganisationhierarchytypeOther.Name {
		diffs = append(diffs, codeorganisationhierarchytype.GongMarshallField(stage, "Name"))
	}
	if codeorganisationhierarchytype.NilReason != codeorganisationhierarchytypeOther.NilReason {
		diffs = append(diffs, codeorganisationhierarchytype.GongMarshallField(stage, "NilReason"))
	}
	if codeorganisationhierarchytype.EnclosedText != codeorganisationhierarchytypeOther.EnclosedText {
		diffs = append(diffs, codeorganisationhierarchytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeorganisationtype *CodeOrganisationType) GongDiff(stage *Stage, codeorganisationtypeOther *CodeOrganisationType) (diffs []string) {
	// insertion point for field diffs
	if codeorganisationtype.Name != codeorganisationtypeOther.Name {
		diffs = append(diffs, codeorganisationtype.GongMarshallField(stage, "Name"))
	}
	if codeorganisationtype.NilReason != codeorganisationtypeOther.NilReason {
		diffs = append(diffs, codeorganisationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeorganisationtype.EnclosedText != codeorganisationtypeOther.EnclosedText {
		diffs = append(diffs, codeorganisationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeoxygentype *CodeOxygenType) GongDiff(stage *Stage, codeoxygentypeOther *CodeOxygenType) (diffs []string) {
	// insertion point for field diffs
	if codeoxygentype.Name != codeoxygentypeOther.Name {
		diffs = append(diffs, codeoxygentype.GongMarshallField(stage, "Name"))
	}
	if codeoxygentype.NilReason != codeoxygentypeOther.NilReason {
		diffs = append(diffs, codeoxygentype.GongMarshallField(stage, "NilReason"))
	}
	if codeoxygentype.EnclosedText != codeoxygentypeOther.EnclosedText {
		diffs = append(diffs, codeoxygentype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepartype *CodePARType) GongDiff(stage *Stage, codepartypeOther *CodePARType) (diffs []string) {
	// insertion point for field diffs
	if codepartype.Name != codepartypeOther.Name {
		diffs = append(diffs, codepartype.GongMarshallField(stage, "Name"))
	}
	if codepartype.NilReason != codepartypeOther.NilReason {
		diffs = append(diffs, codepartype.GongMarshallField(stage, "NilReason"))
	}
	if codepartype.EnclosedText != codepartypeOther.EnclosedText {
		diffs = append(diffs, codepartype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepcnmethodtype *CodePCNMethodType) GongDiff(stage *Stage, codepcnmethodtypeOther *CodePCNMethodType) (diffs []string) {
	// insertion point for field diffs
	if codepcnmethodtype.Name != codepcnmethodtypeOther.Name {
		diffs = append(diffs, codepcnmethodtype.GongMarshallField(stage, "Name"))
	}
	if codepcnmethodtype.NilReason != codepcnmethodtypeOther.NilReason {
		diffs = append(diffs, codepcnmethodtype.GongMarshallField(stage, "NilReason"))
	}
	if codepcnmethodtype.EnclosedText != codepcnmethodtypeOther.EnclosedText {
		diffs = append(diffs, codepcnmethodtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepcnpavementtype *CodePCNPavementType) GongDiff(stage *Stage, codepcnpavementtypeOther *CodePCNPavementType) (diffs []string) {
	// insertion point for field diffs
	if codepcnpavementtype.Name != codepcnpavementtypeOther.Name {
		diffs = append(diffs, codepcnpavementtype.GongMarshallField(stage, "Name"))
	}
	if codepcnpavementtype.NilReason != codepcnpavementtypeOther.NilReason {
		diffs = append(diffs, codepcnpavementtype.GongMarshallField(stage, "NilReason"))
	}
	if codepcnpavementtype.EnclosedText != codepcnpavementtypeOther.EnclosedText {
		diffs = append(diffs, codepcnpavementtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepcnsubgradetype *CodePCNSubgradeType) GongDiff(stage *Stage, codepcnsubgradetypeOther *CodePCNSubgradeType) (diffs []string) {
	// insertion point for field diffs
	if codepcnsubgradetype.Name != codepcnsubgradetypeOther.Name {
		diffs = append(diffs, codepcnsubgradetype.GongMarshallField(stage, "Name"))
	}
	if codepcnsubgradetype.NilReason != codepcnsubgradetypeOther.NilReason {
		diffs = append(diffs, codepcnsubgradetype.GongMarshallField(stage, "NilReason"))
	}
	if codepcnsubgradetype.EnclosedText != codepcnsubgradetypeOther.EnclosedText {
		diffs = append(diffs, codepcnsubgradetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepcntyrepressuretype *CodePCNTyrePressureType) GongDiff(stage *Stage, codepcntyrepressuretypeOther *CodePCNTyrePressureType) (diffs []string) {
	// insertion point for field diffs
	if codepcntyrepressuretype.Name != codepcntyrepressuretypeOther.Name {
		diffs = append(diffs, codepcntyrepressuretype.GongMarshallField(stage, "Name"))
	}
	if codepcntyrepressuretype.NilReason != codepcntyrepressuretypeOther.NilReason {
		diffs = append(diffs, codepcntyrepressuretype.GongMarshallField(stage, "NilReason"))
	}
	if codepcntyrepressuretype.EnclosedText != codepcntyrepressuretypeOther.EnclosedText {
		diffs = append(diffs, codepcntyrepressuretype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepassengerservicetype *CodePassengerServiceType) GongDiff(stage *Stage, codepassengerservicetypeOther *CodePassengerServiceType) (diffs []string) {
	// insertion point for field diffs
	if codepassengerservicetype.Name != codepassengerservicetypeOther.Name {
		diffs = append(diffs, codepassengerservicetype.GongMarshallField(stage, "Name"))
	}
	if codepassengerservicetype.NilReason != codepassengerservicetypeOther.NilReason {
		diffs = append(diffs, codepassengerservicetype.GongMarshallField(stage, "NilReason"))
	}
	if codepassengerservicetype.EnclosedText != codepassengerservicetypeOther.EnclosedText {
		diffs = append(diffs, codepassengerservicetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongDiff(stage *Stage, codepilotcontrolledlightingtypeOther *CodePilotControlledLightingType) (diffs []string) {
	// insertion point for field diffs
	if codepilotcontrolledlightingtype.Name != codepilotcontrolledlightingtypeOther.Name {
		diffs = append(diffs, codepilotcontrolledlightingtype.GongMarshallField(stage, "Name"))
	}
	if codepilotcontrolledlightingtype.NilReason != codepilotcontrolledlightingtypeOther.NilReason {
		diffs = append(diffs, codepilotcontrolledlightingtype.GongMarshallField(stage, "NilReason"))
	}
	if codepilotcontrolledlightingtype.EnclosedText != codepilotcontrolledlightingtypeOther.EnclosedText {
		diffs = append(diffs, codepilotcontrolledlightingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepositioninilstype *CodePositionInILSType) GongDiff(stage *Stage, codepositioninilstypeOther *CodePositionInILSType) (diffs []string) {
	// insertion point for field diffs
	if codepositioninilstype.Name != codepositioninilstypeOther.Name {
		diffs = append(diffs, codepositioninilstype.GongMarshallField(stage, "Name"))
	}
	if codepositioninilstype.NilReason != codepositioninilstypeOther.NilReason {
		diffs = append(diffs, codepositioninilstype.GongMarshallField(stage, "NilReason"))
	}
	if codepositioninilstype.EnclosedText != codepositioninilstypeOther.EnclosedText {
		diffs = append(diffs, codepositioninilstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeprimaryradartype *CodePrimaryRadarType) GongDiff(stage *Stage, codeprimaryradartypeOther *CodePrimaryRadarType) (diffs []string) {
	// insertion point for field diffs
	if codeprimaryradartype.Name != codeprimaryradartypeOther.Name {
		diffs = append(diffs, codeprimaryradartype.GongMarshallField(stage, "Name"))
	}
	if codeprimaryradartype.NilReason != codeprimaryradartypeOther.NilReason {
		diffs = append(diffs, codeprimaryradartype.GongMarshallField(stage, "NilReason"))
	}
	if codeprimaryradartype.EnclosedText != codeprimaryradartypeOther.EnclosedText {
		diffs = append(diffs, codeprimaryradartype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongDiff(stage *Stage, codeprocedureavailabilitytypeOther *CodeProcedureAvailabilityType) (diffs []string) {
	// insertion point for field diffs
	if codeprocedureavailabilitytype.Name != codeprocedureavailabilitytypeOther.Name {
		diffs = append(diffs, codeprocedureavailabilitytype.GongMarshallField(stage, "Name"))
	}
	if codeprocedureavailabilitytype.NilReason != codeprocedureavailabilitytypeOther.NilReason {
		diffs = append(diffs, codeprocedureavailabilitytype.GongMarshallField(stage, "NilReason"))
	}
	if codeprocedureavailabilitytype.EnclosedText != codeprocedureavailabilitytypeOther.EnclosedText {
		diffs = append(diffs, codeprocedureavailabilitytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongDiff(stage *Stage, codeprocedurecodingstandardtypeOther *CodeProcedureCodingStandardType) (diffs []string) {
	// insertion point for field diffs
	if codeprocedurecodingstandardtype.Name != codeprocedurecodingstandardtypeOther.Name {
		diffs = append(diffs, codeprocedurecodingstandardtype.GongMarshallField(stage, "Name"))
	}
	if codeprocedurecodingstandardtype.NilReason != codeprocedurecodingstandardtypeOther.NilReason {
		diffs = append(diffs, codeprocedurecodingstandardtype.GongMarshallField(stage, "NilReason"))
	}
	if codeprocedurecodingstandardtype.EnclosedText != codeprocedurecodingstandardtypeOther.EnclosedText {
		diffs = append(diffs, codeprocedurecodingstandardtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeproceduredistancetype *CodeProcedureDistanceType) GongDiff(stage *Stage, codeproceduredistancetypeOther *CodeProcedureDistanceType) (diffs []string) {
	// insertion point for field diffs
	if codeproceduredistancetype.Name != codeproceduredistancetypeOther.Name {
		diffs = append(diffs, codeproceduredistancetype.GongMarshallField(stage, "Name"))
	}
	if codeproceduredistancetype.NilReason != codeproceduredistancetypeOther.NilReason {
		diffs = append(diffs, codeproceduredistancetype.GongMarshallField(stage, "NilReason"))
	}
	if codeproceduredistancetype.EnclosedText != codeproceduredistancetypeOther.EnclosedText {
		diffs = append(diffs, codeproceduredistancetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongDiff(stage *Stage, codeprocedurefixroletypeOther *CodeProcedureFixRoleType) (diffs []string) {
	// insertion point for field diffs
	if codeprocedurefixroletype.Name != codeprocedurefixroletypeOther.Name {
		diffs = append(diffs, codeprocedurefixroletype.GongMarshallField(stage, "Name"))
	}
	if codeprocedurefixroletype.NilReason != codeprocedurefixroletypeOther.NilReason {
		diffs = append(diffs, codeprocedurefixroletype.GongMarshallField(stage, "NilReason"))
	}
	if codeprocedurefixroletype.EnclosedText != codeprocedurefixroletypeOther.EnclosedText {
		diffs = append(diffs, codeprocedurefixroletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeprocedurephasetype *CodeProcedurePhaseType) GongDiff(stage *Stage, codeprocedurephasetypeOther *CodeProcedurePhaseType) (diffs []string) {
	// insertion point for field diffs
	if codeprocedurephasetype.Name != codeprocedurephasetypeOther.Name {
		diffs = append(diffs, codeprocedurephasetype.GongMarshallField(stage, "Name"))
	}
	if codeprocedurephasetype.NilReason != codeprocedurephasetypeOther.NilReason {
		diffs = append(diffs, codeprocedurephasetype.GongMarshallField(stage, "NilReason"))
	}
	if codeprocedurephasetype.EnclosedText != codeprocedurephasetypeOther.EnclosedText {
		diffs = append(diffs, codeprocedurephasetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongDiff(stage *Stage, codeprotectareasectiontypeOther *CodeProtectAreaSectionType) (diffs []string) {
	// insertion point for field diffs
	if codeprotectareasectiontype.Name != codeprotectareasectiontypeOther.Name {
		diffs = append(diffs, codeprotectareasectiontype.GongMarshallField(stage, "Name"))
	}
	if codeprotectareasectiontype.NilReason != codeprotectareasectiontypeOther.NilReason {
		diffs = append(diffs, codeprotectareasectiontype.GongMarshallField(stage, "NilReason"))
	}
	if codeprotectareasectiontype.EnclosedText != codeprotectareasectiontypeOther.EnclosedText {
		diffs = append(diffs, codeprotectareasectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codernptype *CodeRNPType) GongDiff(stage *Stage, codernptypeOther *CodeRNPType) (diffs []string) {
	// insertion point for field diffs
	if codernptype.Name != codernptypeOther.Name {
		diffs = append(diffs, codernptype.GongMarshallField(stage, "Name"))
	}
	if codernptype.NilReason != codernptypeOther.NilReason {
		diffs = append(diffs, codernptype.GongMarshallField(stage, "NilReason"))
	}
	if codernptype.EnclosedText != codernptypeOther.EnclosedText {
		diffs = append(diffs, codernptype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codervrreadingtype *CodeRVRReadingType) GongDiff(stage *Stage, codervrreadingtypeOther *CodeRVRReadingType) (diffs []string) {
	// insertion point for field diffs
	if codervrreadingtype.Name != codervrreadingtypeOther.Name {
		diffs = append(diffs, codervrreadingtype.GongMarshallField(stage, "Name"))
	}
	if codervrreadingtype.NilReason != codervrreadingtypeOther.NilReason {
		diffs = append(diffs, codervrreadingtype.GongMarshallField(stage, "NilReason"))
	}
	if codervrreadingtype.EnclosedText != codervrreadingtypeOther.EnclosedText {
		diffs = append(diffs, codervrreadingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codervsmpointroletype *CodeRVSMPointRoleType) GongDiff(stage *Stage, codervsmpointroletypeOther *CodeRVSMPointRoleType) (diffs []string) {
	// insertion point for field diffs
	if codervsmpointroletype.Name != codervsmpointroletypeOther.Name {
		diffs = append(diffs, codervsmpointroletype.GongMarshallField(stage, "Name"))
	}
	if codervsmpointroletype.NilReason != codervsmpointroletypeOther.NilReason {
		diffs = append(diffs, codervsmpointroletype.GongMarshallField(stage, "NilReason"))
	}
	if codervsmpointroletype.EnclosedText != codervsmpointroletypeOther.EnclosedText {
		diffs = append(diffs, codervsmpointroletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codervsmtype *CodeRVSMType) GongDiff(stage *Stage, codervsmtypeOther *CodeRVSMType) (diffs []string) {
	// insertion point for field diffs
	if codervsmtype.Name != codervsmtypeOther.Name {
		diffs = append(diffs, codervsmtype.GongMarshallField(stage, "Name"))
	}
	if codervsmtype.NilReason != codervsmtypeOther.NilReason {
		diffs = append(diffs, codervsmtype.GongMarshallField(stage, "NilReason"))
	}
	if codervsmtype.EnclosedText != codervsmtypeOther.EnclosedText {
		diffs = append(diffs, codervsmtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderadarservicetype *CodeRadarServiceType) GongDiff(stage *Stage, coderadarservicetypeOther *CodeRadarServiceType) (diffs []string) {
	// insertion point for field diffs
	if coderadarservicetype.Name != coderadarservicetypeOther.Name {
		diffs = append(diffs, coderadarservicetype.GongMarshallField(stage, "Name"))
	}
	if coderadarservicetype.NilReason != coderadarservicetypeOther.NilReason {
		diffs = append(diffs, coderadarservicetype.GongMarshallField(stage, "NilReason"))
	}
	if coderadarservicetype.EnclosedText != coderadarservicetypeOther.EnclosedText {
		diffs = append(diffs, coderadarservicetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderadioemissiontype *CodeRadioEmissionType) GongDiff(stage *Stage, coderadioemissiontypeOther *CodeRadioEmissionType) (diffs []string) {
	// insertion point for field diffs
	if coderadioemissiontype.Name != coderadioemissiontypeOther.Name {
		diffs = append(diffs, coderadioemissiontype.GongMarshallField(stage, "Name"))
	}
	if coderadioemissiontype.NilReason != coderadioemissiontypeOther.NilReason {
		diffs = append(diffs, coderadioemissiontype.GongMarshallField(stage, "NilReason"))
	}
	if coderadioemissiontype.EnclosedText != coderadioemissiontypeOther.EnclosedText {
		diffs = append(diffs, coderadioemissiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongDiff(stage *Stage, coderadiofrequencyareatypeOther *CodeRadioFrequencyAreaType) (diffs []string) {
	// insertion point for field diffs
	if coderadiofrequencyareatype.Name != coderadiofrequencyareatypeOther.Name {
		diffs = append(diffs, coderadiofrequencyareatype.GongMarshallField(stage, "Name"))
	}
	if coderadiofrequencyareatype.NilReason != coderadiofrequencyareatypeOther.NilReason {
		diffs = append(diffs, coderadiofrequencyareatype.GongMarshallField(stage, "NilReason"))
	}
	if coderadiofrequencyareatype.EnclosedText != coderadiofrequencyareatypeOther.EnclosedText {
		diffs = append(diffs, coderadiofrequencyareatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderadiosignaltype *CodeRadioSignalType) GongDiff(stage *Stage, coderadiosignaltypeOther *CodeRadioSignalType) (diffs []string) {
	// insertion point for field diffs
	if coderadiosignaltype.Name != coderadiosignaltypeOther.Name {
		diffs = append(diffs, coderadiosignaltype.GongMarshallField(stage, "Name"))
	}
	if coderadiosignaltype.NilReason != coderadiosignaltypeOther.NilReason {
		diffs = append(diffs, coderadiosignaltype.GongMarshallField(stage, "NilReason"))
	}
	if coderadiosignaltype.EnclosedText != coderadiosignaltypeOther.EnclosedText {
		diffs = append(diffs, coderadiosignaltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongDiff(stage *Stage, codereferencepathidentifiertypeOther *CodeReferencePathIdentifierType) (diffs []string) {
	// insertion point for field diffs
	if codereferencepathidentifiertype.Name != codereferencepathidentifiertypeOther.Name {
		diffs = append(diffs, codereferencepathidentifiertype.GongMarshallField(stage, "Name"))
	}
	if codereferencepathidentifiertype.NilReason != codereferencepathidentifiertypeOther.NilReason {
		diffs = append(diffs, codereferencepathidentifiertype.GongMarshallField(stage, "NilReason"))
	}
	if codereferencepathidentifiertype.EnclosedText != codereferencepathidentifiertypeOther.EnclosedText {
		diffs = append(diffs, codereferencepathidentifiertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codereferenceroletype *CodeReferenceRoleType) GongDiff(stage *Stage, codereferenceroletypeOther *CodeReferenceRoleType) (diffs []string) {
	// insertion point for field diffs
	if codereferenceroletype.Name != codereferenceroletypeOther.Name {
		diffs = append(diffs, codereferenceroletype.GongMarshallField(stage, "Name"))
	}
	if codereferenceroletype.NilReason != codereferenceroletypeOther.NilReason {
		diffs = append(diffs, codereferenceroletype.GongMarshallField(stage, "NilReason"))
	}
	if codereferenceroletype.EnclosedText != codereferenceroletypeOther.EnclosedText {
		diffs = append(diffs, codereferenceroletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codereflectortype *CodeReflectorType) GongDiff(stage *Stage, codereflectortypeOther *CodeReflectorType) (diffs []string) {
	// insertion point for field diffs
	if codereflectortype.Name != codereflectortypeOther.Name {
		diffs = append(diffs, codereflectortype.GongMarshallField(stage, "Name"))
	}
	if codereflectortype.NilReason != codereflectortypeOther.NilReason {
		diffs = append(diffs, codereflectortype.GongMarshallField(stage, "NilReason"))
	}
	if codereflectortype.EnclosedText != codereflectortypeOther.EnclosedText {
		diffs = append(diffs, codereflectortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderelativepositiontype *CodeRelativePositionType) GongDiff(stage *Stage, coderelativepositiontypeOther *CodeRelativePositionType) (diffs []string) {
	// insertion point for field diffs
	if coderelativepositiontype.Name != coderelativepositiontypeOther.Name {
		diffs = append(diffs, coderelativepositiontype.GongMarshallField(stage, "Name"))
	}
	if coderelativepositiontype.NilReason != coderelativepositiontypeOther.NilReason {
		diffs = append(diffs, coderelativepositiontype.GongMarshallField(stage, "NilReason"))
	}
	if coderelativepositiontype.EnclosedText != coderelativepositiontypeOther.EnclosedText {
		diffs = append(diffs, coderelativepositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderoadtype *CodeRoadType) GongDiff(stage *Stage, coderoadtypeOther *CodeRoadType) (diffs []string) {
	// insertion point for field diffs
	if coderoadtype.Name != coderoadtypeOther.Name {
		diffs = append(diffs, coderoadtype.GongMarshallField(stage, "Name"))
	}
	if coderoadtype.NilReason != coderoadtypeOther.NilReason {
		diffs = append(diffs, coderoadtype.GongMarshallField(stage, "NilReason"))
	}
	if coderoadtype.EnclosedText != coderoadtypeOther.EnclosedText {
		diffs = append(diffs, coderoadtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongDiff(stage *Stage, coderouteavailabilitytypeOther *CodeRouteAvailabilityType) (diffs []string) {
	// insertion point for field diffs
	if coderouteavailabilitytype.Name != coderouteavailabilitytypeOther.Name {
		diffs = append(diffs, coderouteavailabilitytype.GongMarshallField(stage, "Name"))
	}
	if coderouteavailabilitytype.NilReason != coderouteavailabilitytypeOther.NilReason {
		diffs = append(diffs, coderouteavailabilitytype.GongMarshallField(stage, "NilReason"))
	}
	if coderouteavailabilitytype.EnclosedText != coderouteavailabilitytypeOther.EnclosedText {
		diffs = append(diffs, coderouteavailabilitytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongDiff(stage *Stage, coderoutedesignatorlettertypeOther *CodeRouteDesignatorLetterType) (diffs []string) {
	// insertion point for field diffs
	if coderoutedesignatorlettertype.Name != coderoutedesignatorlettertypeOther.Name {
		diffs = append(diffs, coderoutedesignatorlettertype.GongMarshallField(stage, "Name"))
	}
	if coderoutedesignatorlettertype.NilReason != coderoutedesignatorlettertypeOther.NilReason {
		diffs = append(diffs, coderoutedesignatorlettertype.GongMarshallField(stage, "NilReason"))
	}
	if coderoutedesignatorlettertype.EnclosedText != coderoutedesignatorlettertypeOther.EnclosedText {
		diffs = append(diffs, coderoutedesignatorlettertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongDiff(stage *Stage, coderoutedesignatorprefixtypeOther *CodeRouteDesignatorPrefixType) (diffs []string) {
	// insertion point for field diffs
	if coderoutedesignatorprefixtype.Name != coderoutedesignatorprefixtypeOther.Name {
		diffs = append(diffs, coderoutedesignatorprefixtype.GongMarshallField(stage, "Name"))
	}
	if coderoutedesignatorprefixtype.NilReason != coderoutedesignatorprefixtypeOther.NilReason {
		diffs = append(diffs, coderoutedesignatorprefixtype.GongMarshallField(stage, "NilReason"))
	}
	if coderoutedesignatorprefixtype.EnclosedText != coderoutedesignatorprefixtypeOther.EnclosedText {
		diffs = append(diffs, coderoutedesignatorprefixtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongDiff(stage *Stage, coderoutedesignatorsuffixtypeOther *CodeRouteDesignatorSuffixType) (diffs []string) {
	// insertion point for field diffs
	if coderoutedesignatorsuffixtype.Name != coderoutedesignatorsuffixtypeOther.Name {
		diffs = append(diffs, coderoutedesignatorsuffixtype.GongMarshallField(stage, "Name"))
	}
	if coderoutedesignatorsuffixtype.NilReason != coderoutedesignatorsuffixtypeOther.NilReason {
		diffs = append(diffs, coderoutedesignatorsuffixtype.GongMarshallField(stage, "NilReason"))
	}
	if coderoutedesignatorsuffixtype.EnclosedText != coderoutedesignatorsuffixtypeOther.EnclosedText {
		diffs = append(diffs, coderoutedesignatorsuffixtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderouteindicatortype *CodeRouteIndicatorType) GongDiff(stage *Stage, coderouteindicatortypeOther *CodeRouteIndicatorType) (diffs []string) {
	// insertion point for field diffs
	if coderouteindicatortype.Name != coderouteindicatortypeOther.Name {
		diffs = append(diffs, coderouteindicatortype.GongMarshallField(stage, "Name"))
	}
	if coderouteindicatortype.NilReason != coderouteindicatortypeOther.NilReason {
		diffs = append(diffs, coderouteindicatortype.GongMarshallField(stage, "NilReason"))
	}
	if coderouteindicatortype.EnclosedText != coderouteindicatortypeOther.EnclosedText {
		diffs = append(diffs, coderouteindicatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderoutenavigationtype *CodeRouteNavigationType) GongDiff(stage *Stage, coderoutenavigationtypeOther *CodeRouteNavigationType) (diffs []string) {
	// insertion point for field diffs
	if coderoutenavigationtype.Name != coderoutenavigationtypeOther.Name {
		diffs = append(diffs, coderoutenavigationtype.GongMarshallField(stage, "Name"))
	}
	if coderoutenavigationtype.NilReason != coderoutenavigationtypeOther.NilReason {
		diffs = append(diffs, coderoutenavigationtype.GongMarshallField(stage, "NilReason"))
	}
	if coderoutenavigationtype.EnclosedText != coderoutenavigationtypeOther.EnclosedText {
		diffs = append(diffs, coderoutenavigationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderouteorigintype *CodeRouteOriginType) GongDiff(stage *Stage, coderouteorigintypeOther *CodeRouteOriginType) (diffs []string) {
	// insertion point for field diffs
	if coderouteorigintype.Name != coderouteorigintypeOther.Name {
		diffs = append(diffs, coderouteorigintype.GongMarshallField(stage, "Name"))
	}
	if coderouteorigintype.NilReason != coderouteorigintypeOther.NilReason {
		diffs = append(diffs, coderouteorigintype.GongMarshallField(stage, "NilReason"))
	}
	if coderouteorigintype.EnclosedText != coderouteorigintypeOther.EnclosedText {
		diffs = append(diffs, coderouteorigintype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongDiff(stage *Stage, coderoutesegmentpathtypeOther *CodeRouteSegmentPathType) (diffs []string) {
	// insertion point for field diffs
	if coderoutesegmentpathtype.Name != coderoutesegmentpathtypeOther.Name {
		diffs = append(diffs, coderoutesegmentpathtype.GongMarshallField(stage, "Name"))
	}
	if coderoutesegmentpathtype.NilReason != coderoutesegmentpathtypeOther.NilReason {
		diffs = append(diffs, coderoutesegmentpathtype.GongMarshallField(stage, "NilReason"))
	}
	if coderoutesegmentpathtype.EnclosedText != coderoutesegmentpathtypeOther.EnclosedText {
		diffs = append(diffs, coderoutesegmentpathtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderoutetype *CodeRouteType) GongDiff(stage *Stage, coderoutetypeOther *CodeRouteType) (diffs []string) {
	// insertion point for field diffs
	if coderoutetype.Name != coderoutetypeOther.Name {
		diffs = append(diffs, coderoutetype.GongMarshallField(stage, "Name"))
	}
	if coderoutetype.NilReason != coderoutetypeOther.NilReason {
		diffs = append(diffs, coderoutetype.GongMarshallField(stage, "NilReason"))
	}
	if coderoutetype.EnclosedText != coderoutetypeOther.EnclosedText {
		diffs = append(diffs, coderoutetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongDiff(stage *Stage, coderuleproceduretitletypeOther *CodeRuleProcedureTitleType) (diffs []string) {
	// insertion point for field diffs
	if coderuleproceduretitletype.Name != coderuleproceduretitletypeOther.Name {
		diffs = append(diffs, coderuleproceduretitletype.GongMarshallField(stage, "Name"))
	}
	if coderuleproceduretitletype.NilReason != coderuleproceduretitletypeOther.NilReason {
		diffs = append(diffs, coderuleproceduretitletype.GongMarshallField(stage, "NilReason"))
	}
	if coderuleproceduretitletype.EnclosedText != coderuleproceduretitletypeOther.EnclosedText {
		diffs = append(diffs, coderuleproceduretitletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderuleproceduretype *CodeRuleProcedureType) GongDiff(stage *Stage, coderuleproceduretypeOther *CodeRuleProcedureType) (diffs []string) {
	// insertion point for field diffs
	if coderuleproceduretype.Name != coderuleproceduretypeOther.Name {
		diffs = append(diffs, coderuleproceduretype.GongMarshallField(stage, "Name"))
	}
	if coderuleproceduretype.NilReason != coderuleproceduretypeOther.NilReason {
		diffs = append(diffs, coderuleproceduretype.GongMarshallField(stage, "NilReason"))
	}
	if coderuleproceduretype.EnclosedText != coderuleproceduretypeOther.EnclosedText {
		diffs = append(diffs, coderuleproceduretype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderunwayelementtype *CodeRunwayElementType) GongDiff(stage *Stage, coderunwayelementtypeOther *CodeRunwayElementType) (diffs []string) {
	// insertion point for field diffs
	if coderunwayelementtype.Name != coderunwayelementtypeOther.Name {
		diffs = append(diffs, coderunwayelementtype.GongMarshallField(stage, "Name"))
	}
	if coderunwayelementtype.NilReason != coderunwayelementtypeOther.NilReason {
		diffs = append(diffs, coderunwayelementtype.GongMarshallField(stage, "NilReason"))
	}
	if coderunwayelementtype.EnclosedText != coderunwayelementtypeOther.EnclosedText {
		diffs = append(diffs, coderunwayelementtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderunwaymarkingtype *CodeRunwayMarkingType) GongDiff(stage *Stage, coderunwaymarkingtypeOther *CodeRunwayMarkingType) (diffs []string) {
	// insertion point for field diffs
	if coderunwaymarkingtype.Name != coderunwaymarkingtypeOther.Name {
		diffs = append(diffs, coderunwaymarkingtype.GongMarshallField(stage, "Name"))
	}
	if coderunwaymarkingtype.NilReason != coderunwaymarkingtypeOther.NilReason {
		diffs = append(diffs, coderunwaymarkingtype.GongMarshallField(stage, "NilReason"))
	}
	if coderunwaymarkingtype.EnclosedText != coderunwaymarkingtypeOther.EnclosedText {
		diffs = append(diffs, coderunwaymarkingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderunwaypointroletype *CodeRunwayPointRoleType) GongDiff(stage *Stage, coderunwaypointroletypeOther *CodeRunwayPointRoleType) (diffs []string) {
	// insertion point for field diffs
	if coderunwaypointroletype.Name != coderunwaypointroletypeOther.Name {
		diffs = append(diffs, coderunwaypointroletype.GongMarshallField(stage, "Name"))
	}
	if coderunwaypointroletype.NilReason != coderunwaypointroletypeOther.NilReason {
		diffs = append(diffs, coderunwaypointroletype.GongMarshallField(stage, "NilReason"))
	}
	if coderunwaypointroletype.EnclosedText != coderunwaypointroletypeOther.EnclosedText {
		diffs = append(diffs, coderunwaypointroletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongDiff(stage *Stage, coderunwayprotectionareatypeOther *CodeRunwayProtectionAreaType) (diffs []string) {
	// insertion point for field diffs
	if coderunwayprotectionareatype.Name != coderunwayprotectionareatypeOther.Name {
		diffs = append(diffs, coderunwayprotectionareatype.GongMarshallField(stage, "Name"))
	}
	if coderunwayprotectionareatype.NilReason != coderunwayprotectionareatypeOther.NilReason {
		diffs = append(diffs, coderunwayprotectionareatype.GongMarshallField(stage, "NilReason"))
	}
	if coderunwayprotectionareatype.EnclosedText != coderunwayprotectionareatypeOther.EnclosedText {
		diffs = append(diffs, coderunwayprotectionareatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderunwaysectiontype *CodeRunwaySectionType) GongDiff(stage *Stage, coderunwaysectiontypeOther *CodeRunwaySectionType) (diffs []string) {
	// insertion point for field diffs
	if coderunwaysectiontype.Name != coderunwaysectiontypeOther.Name {
		diffs = append(diffs, coderunwaysectiontype.GongMarshallField(stage, "Name"))
	}
	if coderunwaysectiontype.NilReason != coderunwaysectiontypeOther.NilReason {
		diffs = append(diffs, coderunwaysectiontype.GongMarshallField(stage, "NilReason"))
	}
	if coderunwaysectiontype.EnclosedText != coderunwaysectiontypeOther.EnclosedText {
		diffs = append(diffs, coderunwaysectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderunwaytype *CodeRunwayType) GongDiff(stage *Stage, coderunwaytypeOther *CodeRunwayType) (diffs []string) {
	// insertion point for field diffs
	if coderunwaytype.Name != coderunwaytypeOther.Name {
		diffs = append(diffs, coderunwaytype.GongMarshallField(stage, "Name"))
	}
	if coderunwaytype.NilReason != coderunwaytypeOther.NilReason {
		diffs = append(diffs, coderunwaytype.GongMarshallField(stage, "NilReason"))
	}
	if coderunwaytype.EnclosedText != coderunwaytypeOther.EnclosedText {
		diffs = append(diffs, coderunwaytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesafealtitudetype *CodeSafeAltitudeType) GongDiff(stage *Stage, codesafealtitudetypeOther *CodeSafeAltitudeType) (diffs []string) {
	// insertion point for field diffs
	if codesafealtitudetype.Name != codesafealtitudetypeOther.Name {
		diffs = append(diffs, codesafealtitudetype.GongMarshallField(stage, "Name"))
	}
	if codesafealtitudetype.NilReason != codesafealtitudetypeOther.NilReason {
		diffs = append(diffs, codesafealtitudetype.GongMarshallField(stage, "NilReason"))
	}
	if codesafealtitudetype.EnclosedText != codesafealtitudetypeOther.EnclosedText {
		diffs = append(diffs, codesafealtitudetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesegmentpathtype *CodeSegmentPathType) GongDiff(stage *Stage, codesegmentpathtypeOther *CodeSegmentPathType) (diffs []string) {
	// insertion point for field diffs
	if codesegmentpathtype.Name != codesegmentpathtypeOther.Name {
		diffs = append(diffs, codesegmentpathtype.GongMarshallField(stage, "Name"))
	}
	if codesegmentpathtype.NilReason != codesegmentpathtypeOther.NilReason {
		diffs = append(diffs, codesegmentpathtype.GongMarshallField(stage, "NilReason"))
	}
	if codesegmentpathtype.EnclosedText != codesegmentpathtypeOther.EnclosedText {
		diffs = append(diffs, codesegmentpathtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesegmentterminationtype *CodeSegmentTerminationType) GongDiff(stage *Stage, codesegmentterminationtypeOther *CodeSegmentTerminationType) (diffs []string) {
	// insertion point for field diffs
	if codesegmentterminationtype.Name != codesegmentterminationtypeOther.Name {
		diffs = append(diffs, codesegmentterminationtype.GongMarshallField(stage, "Name"))
	}
	if codesegmentterminationtype.NilReason != codesegmentterminationtypeOther.NilReason {
		diffs = append(diffs, codesegmentterminationtype.GongMarshallField(stage, "NilReason"))
	}
	if codesegmentterminationtype.EnclosedText != codesegmentterminationtypeOther.EnclosedText {
		diffs = append(diffs, codesegmentterminationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeserviceatctype *CodeServiceATCType) GongDiff(stage *Stage, codeserviceatctypeOther *CodeServiceATCType) (diffs []string) {
	// insertion point for field diffs
	if codeserviceatctype.Name != codeserviceatctypeOther.Name {
		diffs = append(diffs, codeserviceatctype.GongMarshallField(stage, "Name"))
	}
	if codeserviceatctype.NilReason != codeserviceatctypeOther.NilReason {
		diffs = append(diffs, codeserviceatctype.GongMarshallField(stage, "NilReason"))
	}
	if codeserviceatctype.EnclosedText != codeserviceatctypeOther.EnclosedText {
		diffs = append(diffs, codeserviceatctype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeserviceatfmtype *CodeServiceATFMType) GongDiff(stage *Stage, codeserviceatfmtypeOther *CodeServiceATFMType) (diffs []string) {
	// insertion point for field diffs
	if codeserviceatfmtype.Name != codeserviceatfmtypeOther.Name {
		diffs = append(diffs, codeserviceatfmtype.GongMarshallField(stage, "Name"))
	}
	if codeserviceatfmtype.NilReason != codeserviceatfmtypeOther.NilReason {
		diffs = append(diffs, codeserviceatfmtype.GongMarshallField(stage, "NilReason"))
	}
	if codeserviceatfmtype.EnclosedText != codeserviceatfmtypeOther.EnclosedText {
		diffs = append(diffs, codeserviceatfmtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongDiff(stage *Stage, codeservicegroundcontroltypeOther *CodeServiceGroundControlType) (diffs []string) {
	// insertion point for field diffs
	if codeservicegroundcontroltype.Name != codeservicegroundcontroltypeOther.Name {
		diffs = append(diffs, codeservicegroundcontroltype.GongMarshallField(stage, "Name"))
	}
	if codeservicegroundcontroltype.NilReason != codeservicegroundcontroltypeOther.NilReason {
		diffs = append(diffs, codeservicegroundcontroltype.GongMarshallField(stage, "NilReason"))
	}
	if codeservicegroundcontroltype.EnclosedText != codeservicegroundcontroltypeOther.EnclosedText {
		diffs = append(diffs, codeservicegroundcontroltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeserviceinformationtype *CodeServiceInformationType) GongDiff(stage *Stage, codeserviceinformationtypeOther *CodeServiceInformationType) (diffs []string) {
	// insertion point for field diffs
	if codeserviceinformationtype.Name != codeserviceinformationtypeOther.Name {
		diffs = append(diffs, codeserviceinformationtype.GongMarshallField(stage, "Name"))
	}
	if codeserviceinformationtype.NilReason != codeserviceinformationtypeOther.NilReason {
		diffs = append(diffs, codeserviceinformationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeserviceinformationtype.EnclosedText != codeserviceinformationtypeOther.EnclosedText {
		diffs = append(diffs, codeserviceinformationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeservicesartype *CodeServiceSARType) GongDiff(stage *Stage, codeservicesartypeOther *CodeServiceSARType) (diffs []string) {
	// insertion point for field diffs
	if codeservicesartype.Name != codeservicesartypeOther.Name {
		diffs = append(diffs, codeservicesartype.GongMarshallField(stage, "Name"))
	}
	if codeservicesartype.NilReason != codeservicesartypeOther.NilReason {
		diffs = append(diffs, codeservicesartype.GongMarshallField(stage, "NilReason"))
	}
	if codeservicesartype.EnclosedText != codeservicesartypeOther.EnclosedText {
		diffs = append(diffs, codeservicesartype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesidetype *CodeSideType) GongDiff(stage *Stage, codesidetypeOther *CodeSideType) (diffs []string) {
	// insertion point for field diffs
	if codesidetype.Name != codesidetypeOther.Name {
		diffs = append(diffs, codesidetype.GongMarshallField(stage, "Name"))
	}
	if codesidetype.NilReason != codesidetypeOther.NilReason {
		diffs = append(diffs, codesidetype.GongMarshallField(stage, "NilReason"))
	}
	if codesidetype.EnclosedText != codesidetypeOther.EnclosedText {
		diffs = append(diffs, codesidetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongDiff(stage *Stage, codesignalperformanceilstypeOther *CodeSignalPerformanceILSType) (diffs []string) {
	// insertion point for field diffs
	if codesignalperformanceilstype.Name != codesignalperformanceilstypeOther.Name {
		diffs = append(diffs, codesignalperformanceilstype.GongMarshallField(stage, "Name"))
	}
	if codesignalperformanceilstype.NilReason != codesignalperformanceilstypeOther.NilReason {
		diffs = append(diffs, codesignalperformanceilstype.GongMarshallField(stage, "NilReason"))
	}
	if codesignalperformanceilstype.EnclosedText != codesignalperformanceilstypeOther.EnclosedText {
		diffs = append(diffs, codesignalperformanceilstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codespecialdatetype *CodeSpecialDateType) GongDiff(stage *Stage, codespecialdatetypeOther *CodeSpecialDateType) (diffs []string) {
	// insertion point for field diffs
	if codespecialdatetype.Name != codespecialdatetypeOther.Name {
		diffs = append(diffs, codespecialdatetype.GongMarshallField(stage, "Name"))
	}
	if codespecialdatetype.NilReason != codespecialdatetypeOther.NilReason {
		diffs = append(diffs, codespecialdatetype.GongMarshallField(stage, "NilReason"))
	}
	if codespecialdatetype.EnclosedText != codespecialdatetypeOther.EnclosedText {
		diffs = append(diffs, codespecialdatetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongDiff(stage *Stage, codespecialnavigationchaindesignatortypeOther *CodeSpecialNavigationChainDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codespecialnavigationchaindesignatortype.Name != codespecialnavigationchaindesignatortypeOther.Name {
		diffs = append(diffs, codespecialnavigationchaindesignatortype.GongMarshallField(stage, "Name"))
	}
	if codespecialnavigationchaindesignatortype.NilReason != codespecialnavigationchaindesignatortypeOther.NilReason {
		diffs = append(diffs, codespecialnavigationchaindesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codespecialnavigationchaindesignatortype.EnclosedText != codespecialnavigationchaindesignatortypeOther.EnclosedText {
		diffs = append(diffs, codespecialnavigationchaindesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongDiff(stage *Stage, codespecialnavigationstationtypeOther *CodeSpecialNavigationStationType) (diffs []string) {
	// insertion point for field diffs
	if codespecialnavigationstationtype.Name != codespecialnavigationstationtypeOther.Name {
		diffs = append(diffs, codespecialnavigationstationtype.GongMarshallField(stage, "Name"))
	}
	if codespecialnavigationstationtype.NilReason != codespecialnavigationstationtypeOther.NilReason {
		diffs = append(diffs, codespecialnavigationstationtype.GongMarshallField(stage, "NilReason"))
	}
	if codespecialnavigationstationtype.EnclosedText != codespecialnavigationstationtypeOther.EnclosedText {
		diffs = append(diffs, codespecialnavigationstationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongDiff(stage *Stage, codespecialnavigationsystemtypeOther *CodeSpecialNavigationSystemType) (diffs []string) {
	// insertion point for field diffs
	if codespecialnavigationsystemtype.Name != codespecialnavigationsystemtypeOther.Name {
		diffs = append(diffs, codespecialnavigationsystemtype.GongMarshallField(stage, "Name"))
	}
	if codespecialnavigationsystemtype.NilReason != codespecialnavigationsystemtypeOther.NilReason {
		diffs = append(diffs, codespecialnavigationsystemtype.GongMarshallField(stage, "NilReason"))
	}
	if codespecialnavigationsystemtype.EnclosedText != codespecialnavigationsystemtypeOther.EnclosedText {
		diffs = append(diffs, codespecialnavigationsystemtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codespeedreferencetype *CodeSpeedReferenceType) GongDiff(stage *Stage, codespeedreferencetypeOther *CodeSpeedReferenceType) (diffs []string) {
	// insertion point for field diffs
	if codespeedreferencetype.Name != codespeedreferencetypeOther.Name {
		diffs = append(diffs, codespeedreferencetype.GongMarshallField(stage, "Name"))
	}
	if codespeedreferencetype.NilReason != codespeedreferencetypeOther.NilReason {
		diffs = append(diffs, codespeedreferencetype.GongMarshallField(stage, "NilReason"))
	}
	if codespeedreferencetype.EnclosedText != codespeedreferencetypeOther.EnclosedText {
		diffs = append(diffs, codespeedreferencetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestandbypowertype *CodeStandbyPowerType) GongDiff(stage *Stage, codestandbypowertypeOther *CodeStandbyPowerType) (diffs []string) {
	// insertion point for field diffs
	if codestandbypowertype.Name != codestandbypowertypeOther.Name {
		diffs = append(diffs, codestandbypowertype.GongMarshallField(stage, "Name"))
	}
	if codestandbypowertype.NilReason != codestandbypowertypeOther.NilReason {
		diffs = append(diffs, codestandbypowertype.GongMarshallField(stage, "NilReason"))
	}
	if codestandbypowertype.EnclosedText != codestandbypowertypeOther.EnclosedText {
		diffs = append(diffs, codestandbypowertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestatusairporttype *CodeStatusAirportType) GongDiff(stage *Stage, codestatusairporttypeOther *CodeStatusAirportType) (diffs []string) {
	// insertion point for field diffs
	if codestatusairporttype.Name != codestatusairporttypeOther.Name {
		diffs = append(diffs, codestatusairporttype.GongMarshallField(stage, "Name"))
	}
	if codestatusairporttype.NilReason != codestatusairporttypeOther.NilReason {
		diffs = append(diffs, codestatusairporttype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusairporttype.EnclosedText != codestatusairporttypeOther.EnclosedText {
		diffs = append(diffs, codestatusairporttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestatusairspacetype *CodeStatusAirspaceType) GongDiff(stage *Stage, codestatusairspacetypeOther *CodeStatusAirspaceType) (diffs []string) {
	// insertion point for field diffs
	if codestatusairspacetype.Name != codestatusairspacetypeOther.Name {
		diffs = append(diffs, codestatusairspacetype.GongMarshallField(stage, "Name"))
	}
	if codestatusairspacetype.NilReason != codestatusairspacetypeOther.NilReason {
		diffs = append(diffs, codestatusairspacetype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusairspacetype.EnclosedText != codestatusairspacetypeOther.EnclosedText {
		diffs = append(diffs, codestatusairspacetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestatusconstructiontype *CodeStatusConstructionType) GongDiff(stage *Stage, codestatusconstructiontypeOther *CodeStatusConstructionType) (diffs []string) {
	// insertion point for field diffs
	if codestatusconstructiontype.Name != codestatusconstructiontypeOther.Name {
		diffs = append(diffs, codestatusconstructiontype.GongMarshallField(stage, "Name"))
	}
	if codestatusconstructiontype.NilReason != codestatusconstructiontypeOther.NilReason {
		diffs = append(diffs, codestatusconstructiontype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusconstructiontype.EnclosedText != codestatusconstructiontypeOther.EnclosedText {
		diffs = append(diffs, codestatusconstructiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestatusnavaidtype *CodeStatusNavaidType) GongDiff(stage *Stage, codestatusnavaidtypeOther *CodeStatusNavaidType) (diffs []string) {
	// insertion point for field diffs
	if codestatusnavaidtype.Name != codestatusnavaidtypeOther.Name {
		diffs = append(diffs, codestatusnavaidtype.GongMarshallField(stage, "Name"))
	}
	if codestatusnavaidtype.NilReason != codestatusnavaidtypeOther.NilReason {
		diffs = append(diffs, codestatusnavaidtype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusnavaidtype.EnclosedText != codestatusnavaidtypeOther.EnclosedText {
		diffs = append(diffs, codestatusnavaidtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestatusoperationstype *CodeStatusOperationsType) GongDiff(stage *Stage, codestatusoperationstypeOther *CodeStatusOperationsType) (diffs []string) {
	// insertion point for field diffs
	if codestatusoperationstype.Name != codestatusoperationstypeOther.Name {
		diffs = append(diffs, codestatusoperationstype.GongMarshallField(stage, "Name"))
	}
	if codestatusoperationstype.NilReason != codestatusoperationstypeOther.NilReason {
		diffs = append(diffs, codestatusoperationstype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusoperationstype.EnclosedText != codestatusoperationstypeOther.EnclosedText {
		diffs = append(diffs, codestatusoperationstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestatusservicetype *CodeStatusServiceType) GongDiff(stage *Stage, codestatusservicetypeOther *CodeStatusServiceType) (diffs []string) {
	// insertion point for field diffs
	if codestatusservicetype.Name != codestatusservicetypeOther.Name {
		diffs = append(diffs, codestatusservicetype.GongMarshallField(stage, "Name"))
	}
	if codestatusservicetype.NilReason != codestatusservicetypeOther.NilReason {
		diffs = append(diffs, codestatusservicetype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusservicetype.EnclosedText != codestatusservicetypeOther.EnclosedText {
		diffs = append(diffs, codestatusservicetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongDiff(stage *Stage, codesurfacecompositiontypeOther *CodeSurfaceCompositionType) (diffs []string) {
	// insertion point for field diffs
	if codesurfacecompositiontype.Name != codesurfacecompositiontypeOther.Name {
		diffs = append(diffs, codesurfacecompositiontype.GongMarshallField(stage, "Name"))
	}
	if codesurfacecompositiontype.NilReason != codesurfacecompositiontypeOther.NilReason {
		diffs = append(diffs, codesurfacecompositiontype.GongMarshallField(stage, "NilReason"))
	}
	if codesurfacecompositiontype.EnclosedText != codesurfacecompositiontypeOther.EnclosedText {
		diffs = append(diffs, codesurfacecompositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesurfaceconditiontype *CodeSurfaceConditionType) GongDiff(stage *Stage, codesurfaceconditiontypeOther *CodeSurfaceConditionType) (diffs []string) {
	// insertion point for field diffs
	if codesurfaceconditiontype.Name != codesurfaceconditiontypeOther.Name {
		diffs = append(diffs, codesurfaceconditiontype.GongMarshallField(stage, "Name"))
	}
	if codesurfaceconditiontype.NilReason != codesurfaceconditiontypeOther.NilReason {
		diffs = append(diffs, codesurfaceconditiontype.GongMarshallField(stage, "NilReason"))
	}
	if codesurfaceconditiontype.EnclosedText != codesurfaceconditiontypeOther.EnclosedText {
		diffs = append(diffs, codesurfaceconditiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesurfacepreparationtype *CodeSurfacePreparationType) GongDiff(stage *Stage, codesurfacepreparationtypeOther *CodeSurfacePreparationType) (diffs []string) {
	// insertion point for field diffs
	if codesurfacepreparationtype.Name != codesurfacepreparationtypeOther.Name {
		diffs = append(diffs, codesurfacepreparationtype.GongMarshallField(stage, "Name"))
	}
	if codesurfacepreparationtype.NilReason != codesurfacepreparationtypeOther.NilReason {
		diffs = append(diffs, codesurfacepreparationtype.GongMarshallField(stage, "NilReason"))
	}
	if codesurfacepreparationtype.EnclosedText != codesurfacepreparationtypeOther.EnclosedText {
		diffs = append(diffs, codesurfacepreparationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesystemactivationtype *CodeSystemActivationType) GongDiff(stage *Stage, codesystemactivationtypeOther *CodeSystemActivationType) (diffs []string) {
	// insertion point for field diffs
	if codesystemactivationtype.Name != codesystemactivationtypeOther.Name {
		diffs = append(diffs, codesystemactivationtype.GongMarshallField(stage, "Name"))
	}
	if codesystemactivationtype.NilReason != codesystemactivationtypeOther.NilReason {
		diffs = append(diffs, codesystemactivationtype.GongMarshallField(stage, "NilReason"))
	}
	if codesystemactivationtype.EnclosedText != codesystemactivationtypeOther.EnclosedText {
		diffs = append(diffs, codesystemactivationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetaatype *CodeTAAType) GongDiff(stage *Stage, codetaatypeOther *CodeTAAType) (diffs []string) {
	// insertion point for field diffs
	if codetaatype.Name != codetaatypeOther.Name {
		diffs = append(diffs, codetaatype.GongMarshallField(stage, "Name"))
	}
	if codetaatype.NilReason != codetaatypeOther.NilReason {
		diffs = append(diffs, codetaatype.GongMarshallField(stage, "NilReason"))
	}
	if codetaatype.EnclosedText != codetaatypeOther.EnclosedText {
		diffs = append(diffs, codetaatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetacanchanneltype *CodeTACANChannelType) GongDiff(stage *Stage, codetacanchanneltypeOther *CodeTACANChannelType) (diffs []string) {
	// insertion point for field diffs
	if codetacanchanneltype.Name != codetacanchanneltypeOther.Name {
		diffs = append(diffs, codetacanchanneltype.GongMarshallField(stage, "Name"))
	}
	if codetacanchanneltype.NilReason != codetacanchanneltypeOther.NilReason {
		diffs = append(diffs, codetacanchanneltype.GongMarshallField(stage, "NilReason"))
	}
	if codetacanchanneltype.EnclosedText != codetacanchanneltypeOther.EnclosedText {
		diffs = append(diffs, codetacanchanneltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetlofsectiontype *CodeTLOFSectionType) GongDiff(stage *Stage, codetlofsectiontypeOther *CodeTLOFSectionType) (diffs []string) {
	// insertion point for field diffs
	if codetlofsectiontype.Name != codetlofsectiontypeOther.Name {
		diffs = append(diffs, codetlofsectiontype.GongMarshallField(stage, "Name"))
	}
	if codetlofsectiontype.NilReason != codetlofsectiontypeOther.NilReason {
		diffs = append(diffs, codetlofsectiontype.GongMarshallField(stage, "NilReason"))
	}
	if codetlofsectiontype.EnclosedText != codetlofsectiontypeOther.EnclosedText {
		diffs = append(diffs, codetlofsectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetaxiwayelementtype *CodeTaxiwayElementType) GongDiff(stage *Stage, codetaxiwayelementtypeOther *CodeTaxiwayElementType) (diffs []string) {
	// insertion point for field diffs
	if codetaxiwayelementtype.Name != codetaxiwayelementtypeOther.Name {
		diffs = append(diffs, codetaxiwayelementtype.GongMarshallField(stage, "Name"))
	}
	if codetaxiwayelementtype.NilReason != codetaxiwayelementtypeOther.NilReason {
		diffs = append(diffs, codetaxiwayelementtype.GongMarshallField(stage, "NilReason"))
	}
	if codetaxiwayelementtype.EnclosedText != codetaxiwayelementtypeOther.EnclosedText {
		diffs = append(diffs, codetaxiwayelementtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongDiff(stage *Stage, codetaxiwaysectiontypeOther *CodeTaxiwaySectionType) (diffs []string) {
	// insertion point for field diffs
	if codetaxiwaysectiontype.Name != codetaxiwaysectiontypeOther.Name {
		diffs = append(diffs, codetaxiwaysectiontype.GongMarshallField(stage, "Name"))
	}
	if codetaxiwaysectiontype.NilReason != codetaxiwaysectiontypeOther.NilReason {
		diffs = append(diffs, codetaxiwaysectiontype.GongMarshallField(stage, "NilReason"))
	}
	if codetaxiwaysectiontype.EnclosedText != codetaxiwaysectiontypeOther.EnclosedText {
		diffs = append(diffs, codetaxiwaysectiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetaxiwaytype *CodeTaxiwayType) GongDiff(stage *Stage, codetaxiwaytypeOther *CodeTaxiwayType) (diffs []string) {
	// insertion point for field diffs
	if codetaxiwaytype.Name != codetaxiwaytypeOther.Name {
		diffs = append(diffs, codetaxiwaytype.GongMarshallField(stage, "Name"))
	}
	if codetaxiwaytype.NilReason != codetaxiwaytypeOther.NilReason {
		diffs = append(diffs, codetaxiwaytype.GongMarshallField(stage, "NilReason"))
	}
	if codetaxiwaytype.EnclosedText != codetaxiwaytypeOther.EnclosedText {
		diffs = append(diffs, codetaxiwaytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetelecomnetworktype *CodeTelecomNetworkType) GongDiff(stage *Stage, codetelecomnetworktypeOther *CodeTelecomNetworkType) (diffs []string) {
	// insertion point for field diffs
	if codetelecomnetworktype.Name != codetelecomnetworktypeOther.Name {
		diffs = append(diffs, codetelecomnetworktype.GongMarshallField(stage, "Name"))
	}
	if codetelecomnetworktype.NilReason != codetelecomnetworktypeOther.NilReason {
		diffs = append(diffs, codetelecomnetworktype.GongMarshallField(stage, "NilReason"))
	}
	if codetelecomnetworktype.EnclosedText != codetelecomnetworktypeOther.EnclosedText {
		diffs = append(diffs, codetelecomnetworktype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongDiff(stage *Stage, codetimeeventcombinationtypeOther *CodeTimeEventCombinationType) (diffs []string) {
	// insertion point for field diffs
	if codetimeeventcombinationtype.Name != codetimeeventcombinationtypeOther.Name {
		diffs = append(diffs, codetimeeventcombinationtype.GongMarshallField(stage, "Name"))
	}
	if codetimeeventcombinationtype.NilReason != codetimeeventcombinationtypeOther.NilReason {
		diffs = append(diffs, codetimeeventcombinationtype.GongMarshallField(stage, "NilReason"))
	}
	if codetimeeventcombinationtype.EnclosedText != codetimeeventcombinationtypeOther.EnclosedText {
		diffs = append(diffs, codetimeeventcombinationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetimeeventtype *CodeTimeEventType) GongDiff(stage *Stage, codetimeeventtypeOther *CodeTimeEventType) (diffs []string) {
	// insertion point for field diffs
	if codetimeeventtype.Name != codetimeeventtypeOther.Name {
		diffs = append(diffs, codetimeeventtype.GongMarshallField(stage, "Name"))
	}
	if codetimeeventtype.NilReason != codetimeeventtypeOther.NilReason {
		diffs = append(diffs, codetimeeventtype.GongMarshallField(stage, "NilReason"))
	}
	if codetimeeventtype.EnclosedText != codetimeeventtypeOther.EnclosedText {
		diffs = append(diffs, codetimeeventtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetimereferencetype *CodeTimeReferenceType) GongDiff(stage *Stage, codetimereferencetypeOther *CodeTimeReferenceType) (diffs []string) {
	// insertion point for field diffs
	if codetimereferencetype.Name != codetimereferencetypeOther.Name {
		diffs = append(diffs, codetimereferencetype.GongMarshallField(stage, "Name"))
	}
	if codetimereferencetype.NilReason != codetimereferencetypeOther.NilReason {
		diffs = append(diffs, codetimereferencetype.GongMarshallField(stage, "NilReason"))
	}
	if codetimereferencetype.EnclosedText != codetimereferencetypeOther.EnclosedText {
		diffs = append(diffs, codetimereferencetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetrajectorytype *CodeTrajectoryType) GongDiff(stage *Stage, codetrajectorytypeOther *CodeTrajectoryType) (diffs []string) {
	// insertion point for field diffs
	if codetrajectorytype.Name != codetrajectorytypeOther.Name {
		diffs = append(diffs, codetrajectorytype.GongMarshallField(stage, "Name"))
	}
	if codetrajectorytype.NilReason != codetrajectorytypeOther.NilReason {
		diffs = append(diffs, codetrajectorytype.GongMarshallField(stage, "NilReason"))
	}
	if codetrajectorytype.EnclosedText != codetrajectorytypeOther.EnclosedText {
		diffs = append(diffs, codetrajectorytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetranspondertype *CodeTransponderType) GongDiff(stage *Stage, codetranspondertypeOther *CodeTransponderType) (diffs []string) {
	// insertion point for field diffs
	if codetranspondertype.Name != codetranspondertypeOther.Name {
		diffs = append(diffs, codetranspondertype.GongMarshallField(stage, "Name"))
	}
	if codetranspondertype.NilReason != codetranspondertypeOther.NilReason {
		diffs = append(diffs, codetranspondertype.GongMarshallField(stage, "NilReason"))
	}
	if codetranspondertype.EnclosedText != codetranspondertypeOther.EnclosedText {
		diffs = append(diffs, codetranspondertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeuuidtype *CodeUUIDType) GongDiff(stage *Stage, codeuuidtypeOther *CodeUUIDType) (diffs []string) {
	// insertion point for field diffs
	if codeuuidtype.Name != codeuuidtypeOther.Name {
		diffs = append(diffs, codeuuidtype.GongMarshallField(stage, "Name"))
	}
	if codeuuidtype.NilReason != codeuuidtypeOther.NilReason {
		diffs = append(diffs, codeuuidtype.GongMarshallField(stage, "NilReason"))
	}
	if codeuuidtype.EnclosedText != codeuuidtypeOther.EnclosedText {
		diffs = append(diffs, codeuuidtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeunitdependencytype *CodeUnitDependencyType) GongDiff(stage *Stage, codeunitdependencytypeOther *CodeUnitDependencyType) (diffs []string) {
	// insertion point for field diffs
	if codeunitdependencytype.Name != codeunitdependencytypeOther.Name {
		diffs = append(diffs, codeunitdependencytype.GongMarshallField(stage, "Name"))
	}
	if codeunitdependencytype.NilReason != codeunitdependencytypeOther.NilReason {
		diffs = append(diffs, codeunitdependencytype.GongMarshallField(stage, "NilReason"))
	}
	if codeunitdependencytype.EnclosedText != codeunitdependencytypeOther.EnclosedText {
		diffs = append(diffs, codeunitdependencytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeunittype *CodeUnitType) GongDiff(stage *Stage, codeunittypeOther *CodeUnitType) (diffs []string) {
	// insertion point for field diffs
	if codeunittype.Name != codeunittypeOther.Name {
		diffs = append(diffs, codeunittype.GongMarshallField(stage, "Name"))
	}
	if codeunittype.NilReason != codeunittypeOther.NilReason {
		diffs = append(diffs, codeunittype.GongMarshallField(stage, "NilReason"))
	}
	if codeunittype.EnclosedText != codeunittypeOther.EnclosedText {
		diffs = append(diffs, codeunittype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeupperalphatype *CodeUpperAlphaType) GongDiff(stage *Stage, codeupperalphatypeOther *CodeUpperAlphaType) (diffs []string) {
	// insertion point for field diffs
	if codeupperalphatype.Name != codeupperalphatypeOther.Name {
		diffs = append(diffs, codeupperalphatype.GongMarshallField(stage, "Name"))
	}
	if codeupperalphatype.NilReason != codeupperalphatypeOther.NilReason {
		diffs = append(diffs, codeupperalphatype.GongMarshallField(stage, "NilReason"))
	}
	if codeupperalphatype.EnclosedText != codeupperalphatypeOther.EnclosedText {
		diffs = append(diffs, codeupperalphatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeusagelimitationtype *CodeUsageLimitationType) GongDiff(stage *Stage, codeusagelimitationtypeOther *CodeUsageLimitationType) (diffs []string) {
	// insertion point for field diffs
	if codeusagelimitationtype.Name != codeusagelimitationtypeOther.Name {
		diffs = append(diffs, codeusagelimitationtype.GongMarshallField(stage, "Name"))
	}
	if codeusagelimitationtype.NilReason != codeusagelimitationtypeOther.NilReason {
		diffs = append(diffs, codeusagelimitationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeusagelimitationtype.EnclosedText != codeusagelimitationtypeOther.EnclosedText {
		diffs = append(diffs, codeusagelimitationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codevasistype *CodeVASISType) GongDiff(stage *Stage, codevasistypeOther *CodeVASISType) (diffs []string) {
	// insertion point for field diffs
	if codevasistype.Name != codevasistypeOther.Name {
		diffs = append(diffs, codevasistype.GongMarshallField(stage, "Name"))
	}
	if codevasistype.NilReason != codevasistypeOther.NilReason {
		diffs = append(diffs, codevasistype.GongMarshallField(stage, "NilReason"))
	}
	if codevasistype.EnclosedText != codevasistypeOther.EnclosedText {
		diffs = append(diffs, codevasistype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codevortype *CodeVORType) GongDiff(stage *Stage, codevortypeOther *CodeVORType) (diffs []string) {
	// insertion point for field diffs
	if codevortype.Name != codevortypeOther.Name {
		diffs = append(diffs, codevortype.GongMarshallField(stage, "Name"))
	}
	if codevortype.NilReason != codevortypeOther.NilReason {
		diffs = append(diffs, codevortype.GongMarshallField(stage, "NilReason"))
	}
	if codevortype.EnclosedText != codevortypeOther.EnclosedText {
		diffs = append(diffs, codevortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codevalueinterpretationtype *CodeValueInterpretationType) GongDiff(stage *Stage, codevalueinterpretationtypeOther *CodeValueInterpretationType) (diffs []string) {
	// insertion point for field diffs
	if codevalueinterpretationtype.Name != codevalueinterpretationtypeOther.Name {
		diffs = append(diffs, codevalueinterpretationtype.GongMarshallField(stage, "Name"))
	}
	if codevalueinterpretationtype.NilReason != codevalueinterpretationtypeOther.NilReason {
		diffs = append(diffs, codevalueinterpretationtype.GongMarshallField(stage, "NilReason"))
	}
	if codevalueinterpretationtype.EnclosedText != codevalueinterpretationtypeOther.EnclosedText {
		diffs = append(diffs, codevalueinterpretationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeverticaldatumtype *CodeVerticalDatumType) GongDiff(stage *Stage, codeverticaldatumtypeOther *CodeVerticalDatumType) (diffs []string) {
	// insertion point for field diffs
	if codeverticaldatumtype.Name != codeverticaldatumtypeOther.Name {
		diffs = append(diffs, codeverticaldatumtype.GongMarshallField(stage, "Name"))
	}
	if codeverticaldatumtype.NilReason != codeverticaldatumtypeOther.NilReason {
		diffs = append(diffs, codeverticaldatumtype.GongMarshallField(stage, "NilReason"))
	}
	if codeverticaldatumtype.EnclosedText != codeverticaldatumtypeOther.EnclosedText {
		diffs = append(diffs, codeverticaldatumtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeverticalreferencetype *CodeVerticalReferenceType) GongDiff(stage *Stage, codeverticalreferencetypeOther *CodeVerticalReferenceType) (diffs []string) {
	// insertion point for field diffs
	if codeverticalreferencetype.Name != codeverticalreferencetypeOther.Name {
		diffs = append(diffs, codeverticalreferencetype.GongMarshallField(stage, "Name"))
	}
	if codeverticalreferencetype.NilReason != codeverticalreferencetypeOther.NilReason {
		diffs = append(diffs, codeverticalreferencetype.GongMarshallField(stage, "NilReason"))
	}
	if codeverticalreferencetype.EnclosedText != codeverticalreferencetypeOther.EnclosedText {
		diffs = append(diffs, codeverticalreferencetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongDiff(stage *Stage, codeverticalstructuremarkingtypeOther *CodeVerticalStructureMarkingType) (diffs []string) {
	// insertion point for field diffs
	if codeverticalstructuremarkingtype.Name != codeverticalstructuremarkingtypeOther.Name {
		diffs = append(diffs, codeverticalstructuremarkingtype.GongMarshallField(stage, "Name"))
	}
	if codeverticalstructuremarkingtype.NilReason != codeverticalstructuremarkingtypeOther.NilReason {
		diffs = append(diffs, codeverticalstructuremarkingtype.GongMarshallField(stage, "NilReason"))
	}
	if codeverticalstructuremarkingtype.EnclosedText != codeverticalstructuremarkingtypeOther.EnclosedText {
		diffs = append(diffs, codeverticalstructuremarkingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongDiff(stage *Stage, codeverticalstructurematerialtypeOther *CodeVerticalStructureMaterialType) (diffs []string) {
	// insertion point for field diffs
	if codeverticalstructurematerialtype.Name != codeverticalstructurematerialtypeOther.Name {
		diffs = append(diffs, codeverticalstructurematerialtype.GongMarshallField(stage, "Name"))
	}
	if codeverticalstructurematerialtype.NilReason != codeverticalstructurematerialtypeOther.NilReason {
		diffs = append(diffs, codeverticalstructurematerialtype.GongMarshallField(stage, "NilReason"))
	}
	if codeverticalstructurematerialtype.EnclosedText != codeverticalstructurematerialtypeOther.EnclosedText {
		diffs = append(diffs, codeverticalstructurematerialtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeverticalstructuretype *CodeVerticalStructureType) GongDiff(stage *Stage, codeverticalstructuretypeOther *CodeVerticalStructureType) (diffs []string) {
	// insertion point for field diffs
	if codeverticalstructuretype.Name != codeverticalstructuretypeOther.Name {
		diffs = append(diffs, codeverticalstructuretype.GongMarshallField(stage, "Name"))
	}
	if codeverticalstructuretype.NilReason != codeverticalstructuretypeOther.NilReason {
		diffs = append(diffs, codeverticalstructuretype.GongMarshallField(stage, "NilReason"))
	}
	if codeverticalstructuretype.EnclosedText != codeverticalstructuretypeOther.EnclosedText {
		diffs = append(diffs, codeverticalstructuretype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongDiff(stage *Stage, codevisualdockingguidancetypeOther *CodeVisualDockingGuidanceType) (diffs []string) {
	// insertion point for field diffs
	if codevisualdockingguidancetype.Name != codevisualdockingguidancetypeOther.Name {
		diffs = append(diffs, codevisualdockingguidancetype.GongMarshallField(stage, "Name"))
	}
	if codevisualdockingguidancetype.NilReason != codevisualdockingguidancetypeOther.NilReason {
		diffs = append(diffs, codevisualdockingguidancetype.GongMarshallField(stage, "NilReason"))
	}
	if codevisualdockingguidancetype.EnclosedText != codevisualdockingguidancetypeOther.EnclosedText {
		diffs = append(diffs, codevisualdockingguidancetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codewaketurbulencetype *CodeWakeTurbulenceType) GongDiff(stage *Stage, codewaketurbulencetypeOther *CodeWakeTurbulenceType) (diffs []string) {
	// insertion point for field diffs
	if codewaketurbulencetype.Name != codewaketurbulencetypeOther.Name {
		diffs = append(diffs, codewaketurbulencetype.GongMarshallField(stage, "Name"))
	}
	if codewaketurbulencetype.NilReason != codewaketurbulencetypeOther.NilReason {
		diffs = append(diffs, codewaketurbulencetype.GongMarshallField(stage, "NilReason"))
	}
	if codewaketurbulencetype.EnclosedText != codewaketurbulencetypeOther.EnclosedText {
		diffs = append(diffs, codewaketurbulencetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeworkareatype *CodeWorkAreaType) GongDiff(stage *Stage, codeworkareatypeOther *CodeWorkAreaType) (diffs []string) {
	// insertion point for field diffs
	if codeworkareatype.Name != codeworkareatypeOther.Name {
		diffs = append(diffs, codeworkareatype.GongMarshallField(stage, "Name"))
	}
	if codeworkareatype.NilReason != codeworkareatypeOther.NilReason {
		diffs = append(diffs, codeworkareatype.GongMarshallField(stage, "NilReason"))
	}
	if codeworkareatype.EnclosedText != codeworkareatypeOther.EnclosedText {
		diffs = append(diffs, codeworkareatype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeyesnotype *CodeYesNoType) GongDiff(stage *Stage, codeyesnotypeOther *CodeYesNoType) (diffs []string) {
	// insertion point for field diffs
	if codeyesnotype.Name != codeyesnotypeOther.Name {
		diffs = append(diffs, codeyesnotype.GongMarshallField(stage, "Name"))
	}
	if codeyesnotype.NilReason != codeyesnotypeOther.NilReason {
		diffs = append(diffs, codeyesnotype.GongMarshallField(stage, "NilReason"))
	}
	if codeyesnotype.EnclosedText != codeyesnotypeOther.EnclosedText {
		diffs = append(diffs, codeyesnotype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datemonthdaytype *DateMonthDayType) GongDiff(stage *Stage, datemonthdaytypeOther *DateMonthDayType) (diffs []string) {
	// insertion point for field diffs
	if datemonthdaytype.Name != datemonthdaytypeOther.Name {
		diffs = append(diffs, datemonthdaytype.GongMarshallField(stage, "Name"))
	}
	if datemonthdaytype.NilReason != datemonthdaytypeOther.NilReason {
		diffs = append(diffs, datemonthdaytype.GongMarshallField(stage, "NilReason"))
	}
	if datemonthdaytype.EnclosedText != datemonthdaytypeOther.EnclosedText {
		diffs = append(diffs, datemonthdaytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datetimetype *DateTimeType) GongDiff(stage *Stage, datetimetypeOther *DateTimeType) (diffs []string) {
	// insertion point for field diffs
	if datetimetype.Name != datetimetypeOther.Name {
		diffs = append(diffs, datetimetype.GongMarshallField(stage, "Name"))
	}
	if datetimetype.NilReason != datetimetypeOther.NilReason {
		diffs = append(diffs, datetimetype.GongMarshallField(stage, "NilReason"))
	}
	if datetimetype.EnclosedText != datetimetypeOther.EnclosedText {
		diffs = append(diffs, datetimetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datetype *DateType) GongDiff(stage *Stage, datetypeOther *DateType) (diffs []string) {
	// insertion point for field diffs
	if datetype.Name != datetypeOther.Name {
		diffs = append(diffs, datetype.GongMarshallField(stage, "Name"))
	}
	if datetype.NilReason != datetypeOther.NilReason {
		diffs = append(diffs, datetype.GongMarshallField(stage, "NilReason"))
	}
	if datetype.EnclosedText != datetypeOther.EnclosedText {
		diffs = append(diffs, datetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dateyeartype *DateYearType) GongDiff(stage *Stage, dateyeartypeOther *DateYearType) (diffs []string) {
	// insertion point for field diffs
	if dateyeartype.Name != dateyeartypeOther.Name {
		diffs = append(diffs, dateyeartype.GongMarshallField(stage, "Name"))
	}
	if dateyeartype.NilReason != dateyeartypeOther.NilReason {
		diffs = append(diffs, dateyeartype.GongMarshallField(stage, "NilReason"))
	}
	if dateyeartype.EnclosedText != dateyeartypeOther.EnclosedText {
		diffs = append(diffs, dateyeartype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (nonumbertype *NoNumberType) GongDiff(stage *Stage, nonumbertypeOther *NoNumberType) (diffs []string) {
	// insertion point for field diffs
	if nonumbertype.Name != nonumbertypeOther.Name {
		diffs = append(diffs, nonumbertype.GongMarshallField(stage, "Name"))
	}
	if nonumbertype.NilReason != nonumbertypeOther.NilReason {
		diffs = append(diffs, nonumbertype.GongMarshallField(stage, "NilReason"))
	}
	if nonumbertype.EnclosedText != nonumbertypeOther.EnclosedText {
		diffs = append(diffs, nonumbertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (nosequencetype *NoSequenceType) GongDiff(stage *Stage, nosequencetypeOther *NoSequenceType) (diffs []string) {
	// insertion point for field diffs
	if nosequencetype.Name != nosequencetypeOther.Name {
		diffs = append(diffs, nosequencetype.GongMarshallField(stage, "Name"))
	}
	if nosequencetype.NilReason != nosequencetypeOther.NilReason {
		diffs = append(diffs, nosequencetype.GongMarshallField(stage, "NilReason"))
	}
	if nosequencetype.EnclosedText != nosequencetypeOther.EnclosedText {
		diffs = append(diffs, nosequencetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textaddresstype *TextAddressType) GongDiff(stage *Stage, textaddresstypeOther *TextAddressType) (diffs []string) {
	// insertion point for field diffs
	if textaddresstype.Name != textaddresstypeOther.Name {
		diffs = append(diffs, textaddresstype.GongMarshallField(stage, "Name"))
	}
	if textaddresstype.NilReason != textaddresstypeOther.NilReason {
		diffs = append(diffs, textaddresstype.GongMarshallField(stage, "NilReason"))
	}
	if textaddresstype.EnclosedText != textaddresstypeOther.EnclosedText {
		diffs = append(diffs, textaddresstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textdesignatortype *TextDesignatorType) GongDiff(stage *Stage, textdesignatortypeOther *TextDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if textdesignatortype.Name != textdesignatortypeOther.Name {
		diffs = append(diffs, textdesignatortype.GongMarshallField(stage, "Name"))
	}
	if textdesignatortype.NilReason != textdesignatortypeOther.NilReason {
		diffs = append(diffs, textdesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if textdesignatortype.EnclosedText != textdesignatortypeOther.EnclosedText {
		diffs = append(diffs, textdesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textinstructiontype *TextInstructionType) GongDiff(stage *Stage, textinstructiontypeOther *TextInstructionType) (diffs []string) {
	// insertion point for field diffs
	if textinstructiontype.Name != textinstructiontypeOther.Name {
		diffs = append(diffs, textinstructiontype.GongMarshallField(stage, "Name"))
	}
	if textinstructiontype.NilReason != textinstructiontypeOther.NilReason {
		diffs = append(diffs, textinstructiontype.GongMarshallField(stage, "NilReason"))
	}
	if textinstructiontype.EnclosedText != textinstructiontypeOther.EnclosedText {
		diffs = append(diffs, textinstructiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textnametype *TextNameType) GongDiff(stage *Stage, textnametypeOther *TextNameType) (diffs []string) {
	// insertion point for field diffs
	if textnametype.Name != textnametypeOther.Name {
		diffs = append(diffs, textnametype.GongMarshallField(stage, "Name"))
	}
	if textnametype.NilReason != textnametypeOther.NilReason {
		diffs = append(diffs, textnametype.GongMarshallField(stage, "NilReason"))
	}
	if textnametype.EnclosedText != textnametypeOther.EnclosedText {
		diffs = append(diffs, textnametype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textnotetype *TextNoteType) GongDiff(stage *Stage, textnotetypeOther *TextNoteType) (diffs []string) {
	// insertion point for field diffs
	if textnotetype.Name != textnotetypeOther.Name {
		diffs = append(diffs, textnotetype.GongMarshallField(stage, "Name"))
	}
	if textnotetype.NilReason != textnotetypeOther.NilReason {
		diffs = append(diffs, textnotetype.GongMarshallField(stage, "NilReason"))
	}
	if textnotetype.Lang != textnotetypeOther.Lang {
		diffs = append(diffs, textnotetype.GongMarshallField(stage, "Lang"))
	}
	if textnotetype.EnclosedText != textnotetypeOther.EnclosedText {
		diffs = append(diffs, textnotetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textphonetype *TextPhoneType) GongDiff(stage *Stage, textphonetypeOther *TextPhoneType) (diffs []string) {
	// insertion point for field diffs
	if textphonetype.Name != textphonetypeOther.Name {
		diffs = append(diffs, textphonetype.GongMarshallField(stage, "Name"))
	}
	if textphonetype.NilReason != textphonetypeOther.NilReason {
		diffs = append(diffs, textphonetype.GongMarshallField(stage, "NilReason"))
	}
	if textphonetype.EnclosedText != textphonetypeOther.EnclosedText {
		diffs = append(diffs, textphonetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textpropertynametype *TextPropertyNameType) GongDiff(stage *Stage, textpropertynametypeOther *TextPropertyNameType) (diffs []string) {
	// insertion point for field diffs
	if textpropertynametype.Name != textpropertynametypeOther.Name {
		diffs = append(diffs, textpropertynametype.GongMarshallField(stage, "Name"))
	}
	if textpropertynametype.NilReason != textpropertynametypeOther.NilReason {
		diffs = append(diffs, textpropertynametype.GongMarshallField(stage, "NilReason"))
	}
	if textpropertynametype.EnclosedText != textpropertynametypeOther.EnclosedText {
		diffs = append(diffs, textpropertynametype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textremarktype *TextRemarkType) GongDiff(stage *Stage, textremarktypeOther *TextRemarkType) (diffs []string) {
	// insertion point for field diffs
	if textremarktype.Name != textremarktypeOther.Name {
		diffs = append(diffs, textremarktype.GongMarshallField(stage, "Name"))
	}
	if textremarktype.NilReason != textremarktypeOther.NilReason {
		diffs = append(diffs, textremarktype.GongMarshallField(stage, "NilReason"))
	}
	if textremarktype.EnclosedText != textremarktypeOther.EnclosedText {
		diffs = append(diffs, textremarktype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongDiff(stage *Stage, textsidstardesignatortypeOther *TextSIDSTARDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if textsidstardesignatortype.Name != textsidstardesignatortypeOther.Name {
		diffs = append(diffs, textsidstardesignatortype.GongMarshallField(stage, "Name"))
	}
	if textsidstardesignatortype.NilReason != textsidstardesignatortypeOther.NilReason {
		diffs = append(diffs, textsidstardesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if textsidstardesignatortype.EnclosedText != textsidstardesignatortypeOther.EnclosedText {
		diffs = append(diffs, textsidstardesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timetype *TimeType) GongDiff(stage *Stage, timetypeOther *TimeType) (diffs []string) {
	// insertion point for field diffs
	if timetype.Name != timetypeOther.Name {
		diffs = append(diffs, timetype.GongMarshallField(stage, "Name"))
	}
	if timetype.NilReason != timetypeOther.NilReason {
		diffs = append(diffs, timetype.GongMarshallField(stage, "NilReason"))
	}
	if timetype.EnclosedText != timetypeOther.EnclosedText {
		diffs = append(diffs, timetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valalarmlimittype *ValAlarmLimitType) GongDiff(stage *Stage, valalarmlimittypeOther *ValAlarmLimitType) (diffs []string) {
	// insertion point for field diffs
	if valalarmlimittype.Name != valalarmlimittypeOther.Name {
		diffs = append(diffs, valalarmlimittype.GongMarshallField(stage, "Name"))
	}
	if valalarmlimittype.NilReason != valalarmlimittypeOther.NilReason {
		diffs = append(diffs, valalarmlimittype.GongMarshallField(stage, "NilReason"))
	}
	if valalarmlimittype.EnclosedText != valalarmlimittypeOther.EnclosedText {
		diffs = append(diffs, valalarmlimittype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valangletype *ValAngleType) GongDiff(stage *Stage, valangletypeOther *ValAngleType) (diffs []string) {
	// insertion point for field diffs
	if valangletype.Name != valangletypeOther.Name {
		diffs = append(diffs, valangletype.GongMarshallField(stage, "Name"))
	}
	if valangletype.NilReason != valangletypeOther.NilReason {
		diffs = append(diffs, valangletype.GongMarshallField(stage, "NilReason"))
	}
	if valangletype.EnclosedText != valangletypeOther.EnclosedText {
		diffs = append(diffs, valangletype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valbearingtype *ValBearingType) GongDiff(stage *Stage, valbearingtypeOther *ValBearingType) (diffs []string) {
	// insertion point for field diffs
	if valbearingtype.Name != valbearingtypeOther.Name {
		diffs = append(diffs, valbearingtype.GongMarshallField(stage, "Name"))
	}
	if valbearingtype.NilReason != valbearingtypeOther.NilReason {
		diffs = append(diffs, valbearingtype.GongMarshallField(stage, "NilReason"))
	}
	if valbearingtype.EnclosedText != valbearingtypeOther.EnclosedText {
		diffs = append(diffs, valbearingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valchannelnumbertype *ValChannelNumberType) GongDiff(stage *Stage, valchannelnumbertypeOther *ValChannelNumberType) (diffs []string) {
	// insertion point for field diffs
	if valchannelnumbertype.Name != valchannelnumbertypeOther.Name {
		diffs = append(diffs, valchannelnumbertype.GongMarshallField(stage, "Name"))
	}
	if valchannelnumbertype.NilReason != valchannelnumbertypeOther.NilReason {
		diffs = append(diffs, valchannelnumbertype.GongMarshallField(stage, "NilReason"))
	}
	if valchannelnumbertype.EnclosedText != valchannelnumbertypeOther.EnclosedText {
		diffs = append(diffs, valchannelnumbertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valdepthtype *ValDepthType) GongDiff(stage *Stage, valdepthtypeOther *ValDepthType) (diffs []string) {
	// insertion point for field diffs
	if valdepthtype.Name != valdepthtypeOther.Name {
		diffs = append(diffs, valdepthtype.GongMarshallField(stage, "Name"))
	}
	if valdepthtype.Uom != valdepthtypeOther.Uom {
		diffs = append(diffs, valdepthtype.GongMarshallField(stage, "Uom"))
	}
	if valdepthtype.NilReason != valdepthtypeOther.NilReason {
		diffs = append(diffs, valdepthtype.GongMarshallField(stage, "NilReason"))
	}
	if valdepthtype.EnclosedText != valdepthtypeOther.EnclosedText {
		diffs = append(diffs, valdepthtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valdistancesignedtype *ValDistanceSignedType) GongDiff(stage *Stage, valdistancesignedtypeOther *ValDistanceSignedType) (diffs []string) {
	// insertion point for field diffs
	if valdistancesignedtype.Name != valdistancesignedtypeOther.Name {
		diffs = append(diffs, valdistancesignedtype.GongMarshallField(stage, "Name"))
	}
	if valdistancesignedtype.Uom != valdistancesignedtypeOther.Uom {
		diffs = append(diffs, valdistancesignedtype.GongMarshallField(stage, "Uom"))
	}
	if valdistancesignedtype.NilReason != valdistancesignedtypeOther.NilReason {
		diffs = append(diffs, valdistancesignedtype.GongMarshallField(stage, "NilReason"))
	}
	if valdistancesignedtype.EnclosedText != valdistancesignedtypeOther.EnclosedText {
		diffs = append(diffs, valdistancesignedtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valdistancetype *ValDistanceType) GongDiff(stage *Stage, valdistancetypeOther *ValDistanceType) (diffs []string) {
	// insertion point for field diffs
	if valdistancetype.Name != valdistancetypeOther.Name {
		diffs = append(diffs, valdistancetype.GongMarshallField(stage, "Name"))
	}
	if valdistancetype.Uom != valdistancetypeOther.Uom {
		diffs = append(diffs, valdistancetype.GongMarshallField(stage, "Uom"))
	}
	if valdistancetype.NilReason != valdistancetypeOther.NilReason {
		diffs = append(diffs, valdistancetype.GongMarshallField(stage, "NilReason"))
	}
	if valdistancetype.EnclosedText != valdistancetypeOther.EnclosedText {
		diffs = append(diffs, valdistancetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valdistanceverticaltype *ValDistanceVerticalType) GongDiff(stage *Stage, valdistanceverticaltypeOther *ValDistanceVerticalType) (diffs []string) {
	// insertion point for field diffs
	if valdistanceverticaltype.Name != valdistanceverticaltypeOther.Name {
		diffs = append(diffs, valdistanceverticaltype.GongMarshallField(stage, "Name"))
	}
	if valdistanceverticaltype.Uom != valdistanceverticaltypeOther.Uom {
		diffs = append(diffs, valdistanceverticaltype.GongMarshallField(stage, "Uom"))
	}
	if valdistanceverticaltype.NilReason != valdistanceverticaltypeOther.NilReason {
		diffs = append(diffs, valdistanceverticaltype.GongMarshallField(stage, "NilReason"))
	}
	if valdistanceverticaltype.EnclosedText != valdistanceverticaltypeOther.EnclosedText {
		diffs = append(diffs, valdistanceverticaltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valdurationtype *ValDurationType) GongDiff(stage *Stage, valdurationtypeOther *ValDurationType) (diffs []string) {
	// insertion point for field diffs
	if valdurationtype.Name != valdurationtypeOther.Name {
		diffs = append(diffs, valdurationtype.GongMarshallField(stage, "Name"))
	}
	if valdurationtype.Uom != valdurationtypeOther.Uom {
		diffs = append(diffs, valdurationtype.GongMarshallField(stage, "Uom"))
	}
	if valdurationtype.NilReason != valdurationtypeOther.NilReason {
		diffs = append(diffs, valdurationtype.GongMarshallField(stage, "NilReason"))
	}
	if valdurationtype.EnclosedText != valdurationtypeOther.EnclosedText {
		diffs = append(diffs, valdurationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valfltype *ValFLType) GongDiff(stage *Stage, valfltypeOther *ValFLType) (diffs []string) {
	// insertion point for field diffs
	if valfltype.Name != valfltypeOther.Name {
		diffs = append(diffs, valfltype.GongMarshallField(stage, "Name"))
	}
	if valfltype.Uom != valfltypeOther.Uom {
		diffs = append(diffs, valfltype.GongMarshallField(stage, "Uom"))
	}
	if valfltype.NilReason != valfltypeOther.NilReason {
		diffs = append(diffs, valfltype.GongMarshallField(stage, "NilReason"))
	}
	if valfltype.EnclosedText != valfltypeOther.EnclosedText {
		diffs = append(diffs, valfltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valfrequencytype *ValFrequencyType) GongDiff(stage *Stage, valfrequencytypeOther *ValFrequencyType) (diffs []string) {
	// insertion point for field diffs
	if valfrequencytype.Name != valfrequencytypeOther.Name {
		diffs = append(diffs, valfrequencytype.GongMarshallField(stage, "Name"))
	}
	if valfrequencytype.Uom != valfrequencytypeOther.Uom {
		diffs = append(diffs, valfrequencytype.GongMarshallField(stage, "Uom"))
	}
	if valfrequencytype.NilReason != valfrequencytypeOther.NilReason {
		diffs = append(diffs, valfrequencytype.GongMarshallField(stage, "NilReason"))
	}
	if valfrequencytype.EnclosedText != valfrequencytypeOther.EnclosedText {
		diffs = append(diffs, valfrequencytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valfrictiontype *ValFrictionType) GongDiff(stage *Stage, valfrictiontypeOther *ValFrictionType) (diffs []string) {
	// insertion point for field diffs
	if valfrictiontype.Name != valfrictiontypeOther.Name {
		diffs = append(diffs, valfrictiontype.GongMarshallField(stage, "Name"))
	}
	if valfrictiontype.NilReason != valfrictiontypeOther.NilReason {
		diffs = append(diffs, valfrictiontype.GongMarshallField(stage, "NilReason"))
	}
	if valfrictiontype.EnclosedText != valfrictiontypeOther.EnclosedText {
		diffs = append(diffs, valfrictiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valhextype *ValHexType) GongDiff(stage *Stage, valhextypeOther *ValHexType) (diffs []string) {
	// insertion point for field diffs
	if valhextype.Name != valhextypeOther.Name {
		diffs = append(diffs, valhextype.GongMarshallField(stage, "Name"))
	}
	if valhextype.NilReason != valhextypeOther.NilReason {
		diffs = append(diffs, valhextype.GongMarshallField(stage, "NilReason"))
	}
	if valhextype.EnclosedText != valhextypeOther.EnclosedText {
		diffs = append(diffs, valhextype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (vallcntype *ValLCNType) GongDiff(stage *Stage, vallcntypeOther *ValLCNType) (diffs []string) {
	// insertion point for field diffs
	if vallcntype.Name != vallcntypeOther.Name {
		diffs = append(diffs, vallcntype.GongMarshallField(stage, "Name"))
	}
	if vallcntype.NilReason != vallcntypeOther.NilReason {
		diffs = append(diffs, vallcntype.GongMarshallField(stage, "NilReason"))
	}
	if vallcntype.EnclosedText != vallcntypeOther.EnclosedText {
		diffs = append(diffs, vallcntype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (vallightintensitytype *ValLightIntensityType) GongDiff(stage *Stage, vallightintensitytypeOther *ValLightIntensityType) (diffs []string) {
	// insertion point for field diffs
	if vallightintensitytype.Name != vallightintensitytypeOther.Name {
		diffs = append(diffs, vallightintensitytype.GongMarshallField(stage, "Name"))
	}
	if vallightintensitytype.Uom != vallightintensitytypeOther.Uom {
		diffs = append(diffs, vallightintensitytype.GongMarshallField(stage, "Uom"))
	}
	if vallightintensitytype.NilReason != vallightintensitytypeOther.NilReason {
		diffs = append(diffs, vallightintensitytype.GongMarshallField(stage, "NilReason"))
	}
	if vallightintensitytype.EnclosedText != vallightintensitytypeOther.EnclosedText {
		diffs = append(diffs, vallightintensitytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongDiff(stage *Stage, valmagneticvariationchangetypeOther *ValMagneticVariationChangeType) (diffs []string) {
	// insertion point for field diffs
	if valmagneticvariationchangetype.Name != valmagneticvariationchangetypeOther.Name {
		diffs = append(diffs, valmagneticvariationchangetype.GongMarshallField(stage, "Name"))
	}
	if valmagneticvariationchangetype.NilReason != valmagneticvariationchangetypeOther.NilReason {
		diffs = append(diffs, valmagneticvariationchangetype.GongMarshallField(stage, "NilReason"))
	}
	if valmagneticvariationchangetype.EnclosedText != valmagneticvariationchangetypeOther.EnclosedText {
		diffs = append(diffs, valmagneticvariationchangetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valmagneticvariationtype *ValMagneticVariationType) GongDiff(stage *Stage, valmagneticvariationtypeOther *ValMagneticVariationType) (diffs []string) {
	// insertion point for field diffs
	if valmagneticvariationtype.Name != valmagneticvariationtypeOther.Name {
		diffs = append(diffs, valmagneticvariationtype.GongMarshallField(stage, "Name"))
	}
	if valmagneticvariationtype.NilReason != valmagneticvariationtypeOther.NilReason {
		diffs = append(diffs, valmagneticvariationtype.GongMarshallField(stage, "NilReason"))
	}
	if valmagneticvariationtype.EnclosedText != valmagneticvariationtypeOther.EnclosedText {
		diffs = append(diffs, valmagneticvariationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valpcntype *ValPCNType) GongDiff(stage *Stage, valpcntypeOther *ValPCNType) (diffs []string) {
	// insertion point for field diffs
	if valpcntype.Name != valpcntypeOther.Name {
		diffs = append(diffs, valpcntype.GongMarshallField(stage, "Name"))
	}
	if valpcntype.NilReason != valpcntypeOther.NilReason {
		diffs = append(diffs, valpcntype.GongMarshallField(stage, "NilReason"))
	}
	if valpcntype.EnclosedText != valpcntypeOther.EnclosedText {
		diffs = append(diffs, valpcntype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valpercenttype *ValPercentType) GongDiff(stage *Stage, valpercenttypeOther *ValPercentType) (diffs []string) {
	// insertion point for field diffs
	if valpercenttype.Name != valpercenttypeOther.Name {
		diffs = append(diffs, valpercenttype.GongMarshallField(stage, "Name"))
	}
	if valpercenttype.NilReason != valpercenttypeOther.NilReason {
		diffs = append(diffs, valpercenttype.GongMarshallField(stage, "NilReason"))
	}
	if valpercenttype.EnclosedText != valpercenttypeOther.EnclosedText {
		diffs = append(diffs, valpercenttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valpressuretype *ValPressureType) GongDiff(stage *Stage, valpressuretypeOther *ValPressureType) (diffs []string) {
	// insertion point for field diffs
	if valpressuretype.Name != valpressuretypeOther.Name {
		diffs = append(diffs, valpressuretype.GongMarshallField(stage, "Name"))
	}
	if valpressuretype.Uom != valpressuretypeOther.Uom {
		diffs = append(diffs, valpressuretype.GongMarshallField(stage, "Uom"))
	}
	if valpressuretype.NilReason != valpressuretypeOther.NilReason {
		diffs = append(diffs, valpressuretype.GongMarshallField(stage, "NilReason"))
	}
	if valpressuretype.EnclosedText != valpressuretypeOther.EnclosedText {
		diffs = append(diffs, valpressuretype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valslopetype *ValSlopeType) GongDiff(stage *Stage, valslopetypeOther *ValSlopeType) (diffs []string) {
	// insertion point for field diffs
	if valslopetype.Name != valslopetypeOther.Name {
		diffs = append(diffs, valslopetype.GongMarshallField(stage, "Name"))
	}
	if valslopetype.NilReason != valslopetypeOther.NilReason {
		diffs = append(diffs, valslopetype.GongMarshallField(stage, "NilReason"))
	}
	if valslopetype.EnclosedText != valslopetypeOther.EnclosedText {
		diffs = append(diffs, valslopetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valspeedtype *ValSpeedType) GongDiff(stage *Stage, valspeedtypeOther *ValSpeedType) (diffs []string) {
	// insertion point for field diffs
	if valspeedtype.Name != valspeedtypeOther.Name {
		diffs = append(diffs, valspeedtype.GongMarshallField(stage, "Name"))
	}
	if valspeedtype.Uom != valspeedtypeOther.Uom {
		diffs = append(diffs, valspeedtype.GongMarshallField(stage, "Uom"))
	}
	if valspeedtype.NilReason != valspeedtypeOther.NilReason {
		diffs = append(diffs, valspeedtype.GongMarshallField(stage, "NilReason"))
	}
	if valspeedtype.EnclosedText != valspeedtypeOther.EnclosedText {
		diffs = append(diffs, valspeedtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valtemperaturetype *ValTemperatureType) GongDiff(stage *Stage, valtemperaturetypeOther *ValTemperatureType) (diffs []string) {
	// insertion point for field diffs
	if valtemperaturetype.Name != valtemperaturetypeOther.Name {
		diffs = append(diffs, valtemperaturetype.GongMarshallField(stage, "Name"))
	}
	if valtemperaturetype.Uom != valtemperaturetypeOther.Uom {
		diffs = append(diffs, valtemperaturetype.GongMarshallField(stage, "Uom"))
	}
	if valtemperaturetype.NilReason != valtemperaturetypeOther.NilReason {
		diffs = append(diffs, valtemperaturetype.GongMarshallField(stage, "NilReason"))
	}
	if valtemperaturetype.EnclosedText != valtemperaturetypeOther.EnclosedText {
		diffs = append(diffs, valtemperaturetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (valweighttype *ValWeightType) GongDiff(stage *Stage, valweighttypeOther *ValWeightType) (diffs []string) {
	// insertion point for field diffs
	if valweighttype.Name != valweighttypeOther.Name {
		diffs = append(diffs, valweighttype.GongMarshallField(stage, "Name"))
	}
	if valweighttype.Uom != valweighttypeOther.Uom {
		diffs = append(diffs, valweighttype.GongMarshallField(stage, "Uom"))
	}
	if valweighttype.NilReason != valweighttypeOther.NilReason {
		diffs = append(diffs, valweighttype.GongMarshallField(stage, "NilReason"))
	}
	if valweighttype.EnclosedText != valweighttypeOther.EnclosedText {
		diffs = append(diffs, valweighttype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (xhtmltype *XHTMLType) GongDiff(stage *Stage, xhtmltypeOther *XHTMLType) (diffs []string) {
	// insertion point for field diffs
	if xhtmltype.Name != xhtmltypeOther.Name {
		diffs = append(diffs, xhtmltype.GongMarshallField(stage, "Name"))
	}
	if xhtmltype.NilReason != xhtmltypeOther.NilReason {
		diffs = append(diffs, xhtmltype.GongMarshallField(stage, "NilReason"))
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
