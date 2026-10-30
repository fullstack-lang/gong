// generated code - do not edit
package aixmdatatypes

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.CodeATCReportingTypes)

	res = __gong__appendInstances(res, stage.CodeAerialRefuellingPointTypes)

	res = __gong__appendInstances(res, stage.CodeAerialRefuellingPrefixTypes)

	res = __gong__appendInstances(res, stage.CodeAerialRefuellingTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftCategoryTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftEngineNumberTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftEngineTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftGroundServiceTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftICAOTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftStandTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftWingspanClassTypes)

	res = __gong__appendInstances(res, stage.CodeAirportHeliportCollocationTypes)

	res = __gong__appendInstances(res, stage.CodeAirportHeliportDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeAirportHeliportTypes)

	res = __gong__appendInstances(res, stage.CodeAirportWarningTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceActivityTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceAggregationTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceClassificationTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceDependencyTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeAirspacePointPositionTypes)

	res = __gong__appendInstances(res, stage.CodeAirspacePointRoleTypes)

	res = __gong__appendInstances(res, stage.CodeAirspaceTypes)

	res = __gong__appendInstances(res, stage.CodeAltitudeAdjustmentTypes)

	res = __gong__appendInstances(res, stage.CodeAltitudeUseTypes)

	res = __gong__appendInstances(res, stage.CodeApproachEquipmentAdditionalTypes)

	res = __gong__appendInstances(res, stage.CodeApproachGuidanceTypes)

	res = __gong__appendInstances(res, stage.CodeApproachLightingICAOTypes)

	res = __gong__appendInstances(res, stage.CodeApproachLightingTypes)

	res = __gong__appendInstances(res, stage.CodeApproachPrefixTypes)

	res = __gong__appendInstances(res, stage.CodeApproachTypes)

	res = __gong__appendInstances(res, stage.CodeApprovalTypes)

	res = __gong__appendInstances(res, stage.CodeApronElementTypes)

	res = __gong__appendInstances(res, stage.CodeApronSectionTypes)

	res = __gong__appendInstances(res, stage.CodeArcDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeArrestingGearEnergyAbsorbTypes)

	res = __gong__appendInstances(res, stage.CodeArrestingGearEngageDeviceTypes)

	res = __gong__appendInstances(res, stage.CodeAuralMorseTypes)

	res = __gong__appendInstances(res, stage.CodeAuthorityRoleTypes)

	res = __gong__appendInstances(res, stage.CodeAuthorityTypes)

	res = __gong__appendInstances(res, stage.CodeAviationStandardsTypes)

	res = __gong__appendInstances(res, stage.CodeBearingTypes)

	res = __gong__appendInstances(res, stage.CodeBuoyDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeBuoyTypes)

	res = __gong__appendInstances(res, stage.CodeCardinalDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeCheckpointCategoryTypes)

	res = __gong__appendInstances(res, stage.CodeColourTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationChannelTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationModeTypes)

	res = __gong__appendInstances(res, stage.CodeComparisonTypes)

	res = __gong__appendInstances(res, stage.CodeContaminationTypes)

	res = __gong__appendInstances(res, stage.CodeCourseQualityILSTypes)

	res = __gong__appendInstances(res, stage.CodeCourseTypes)

	res = __gong__appendInstances(res, stage.CodeDMEChannelTypes)

	res = __gong__appendInstances(res, stage.CodeDMETypes)

	res = __gong__appendInstances(res, stage.CodeDayTypes)

	res = __gong__appendInstances(res, stage.CodeDeclaredDistanceTypes)

	res = __gong__appendInstances(res, stage.CodeDesignStandardTypes)

	res = __gong__appendInstances(res, stage.CodeDesignatedPointDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeDesignatedPointTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionTurnTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeDistanceIndicationTypes)

	res = __gong__appendInstances(res, stage.CodeEmissionBandTypes)

	res = __gong__appendInstances(res, stage.CodeEquipmentAntiCollisionTypes)

	res = __gong__appendInstances(res, stage.CodeEquipmentUnavailableTypes)

	res = __gong__appendInstances(res, stage.CodeFacilityRankingTypes)

	res = __gong__appendInstances(res, stage.CodeFinalGuidanceTypes)

	res = __gong__appendInstances(res, stage.CodeFireFightingTypes)

	res = __gong__appendInstances(res, stage.CodeFlightDestinationTypes)

	res = __gong__appendInstances(res, stage.CodeFlightOriginTypes)

	res = __gong__appendInstances(res, stage.CodeFlightPurposeTypes)

	res = __gong__appendInstances(res, stage.CodeFlightRestrictionDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeFlightRestrictionTypes)

	res = __gong__appendInstances(res, stage.CodeFlightRuleTypes)

	res = __gong__appendInstances(res, stage.CodeFlightStatusTypes)

	res = __gong__appendInstances(res, stage.CodeFlightTypes)

	res = __gong__appendInstances(res, stage.CodeFlowConditionOperationTypes)

	res = __gong__appendInstances(res, stage.CodeFreeFlightTypes)

	res = __gong__appendInstances(res, stage.CodeFrictionDeviceTypes)

	res = __gong__appendInstances(res, stage.CodeFrictionEstimateTypes)

	res = __gong__appendInstances(res, stage.CodeFuelTypes)

	res = __gong__appendInstances(res, stage.CodeGeoBorderTypes)

	res = __gong__appendInstances(res, stage.CodeGradeSeparationTypes)

	res = __gong__appendInstances(res, stage.CodeGroundLightingTypes)

	res = __gong__appendInstances(res, stage.CodeGuidanceLineTypes)

	res = __gong__appendInstances(res, stage.CodeHeightReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeHelicopterPerformanceTypes)

	res = __gong__appendInstances(res, stage.CodeHoldingCategoryTypes)

	res = __gong__appendInstances(res, stage.CodeHoldingUsageTypes)

	res = __gong__appendInstances(res, stage.CodeHoldingUseTypes)

	res = __gong__appendInstances(res, stage.CodeIATATypes)

	res = __gong__appendInstances(res, stage.CodeICAOCountryTypes)

	res = __gong__appendInstances(res, stage.CodeICAOTypes)

	res = __gong__appendInstances(res, stage.CodeILSBackCourseTypes)

	res = __gong__appendInstances(res, stage.CodeIntegrityLevelILSTypes)

	res = __gong__appendInstances(res, stage.CodeIntensityStandByTypes)

	res = __gong__appendInstances(res, stage.CodeLanguageTypes)

	res = __gong__appendInstances(res, stage.CodeLevelSeriesTypes)

	res = __gong__appendInstances(res, stage.CodeLevelTableDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeLevelTypes)

	res = __gong__appendInstances(res, stage.CodeLightHoldingPositionTypes)

	res = __gong__appendInstances(res, stage.CodeLightIntensityTypes)

	res = __gong__appendInstances(res, stage.CodeLightSourceTypes)

	res = __gong__appendInstances(res, stage.CodeLightingJARTypes)

	res = __gong__appendInstances(res, stage.CodeLoadingBridgeTypes)

	res = __gong__appendInstances(res, stage.CodeLocationQualifierTypes)

	res = __gong__appendInstances(res, stage.CodeLogicalOperatorTypes)

	res = __gong__appendInstances(res, stage.CodeMLSAzimuthTypes)

	res = __gong__appendInstances(res, stage.CodeMLSChannelTypes)

	res = __gong__appendInstances(res, stage.CodeMarkerBeaconSignalTypes)

	res = __gong__appendInstances(res, stage.CodeMarkingConditionTypes)

	res = __gong__appendInstances(res, stage.CodeMarkingStyleTypes)

	res = __gong__appendInstances(res, stage.CodeMeteoConditionsTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryOperationsTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryRoutePointTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryStatusTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryTrainingTypes)

	res = __gong__appendInstances(res, stage.CodeMinimaFinalApproachPathTypes)

	res = __gong__appendInstances(res, stage.CodeMinimumAltitudeTypes)

	res = __gong__appendInstances(res, stage.CodeMinimumHeightTypes)

	res = __gong__appendInstances(res, stage.CodeMissedApproachTypes)

	res = __gong__appendInstances(res, stage.CodeNDBUsageTypes)

	res = __gong__appendInstances(res, stage.CodeNavaidDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeNavaidPurposeTypes)

	res = __gong__appendInstances(res, stage.CodeNavaidServiceTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationAreaRestrictionTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationAreaTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationEquipmentTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationSpecificationTypes)

	res = __gong__appendInstances(res, stage.CodeNitrogenTypes)

	res = __gong__appendInstances(res, stage.CodeNorthReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeNotePurposeTypes)

	res = __gong__appendInstances(res, stage.CodeObstacleAreaTypes)

	res = __gong__appendInstances(res, stage.CodeObstacleAssessmentSurfaceTypes)

	res = __gong__appendInstances(res, stage.CodeObstructionIdSurfaceZoneTypes)

	res = __gong__appendInstances(res, stage.CodeOilTypes)

	res = __gong__appendInstances(res, stage.CodeOperationAirportHeliportTypes)

	res = __gong__appendInstances(res, stage.CodeOperationManoeuvringAreaTypes)

	res = __gong__appendInstances(res, stage.CodeOrganisationDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeOrganisationHierarchyTypes)

	res = __gong__appendInstances(res, stage.CodeOrganisationTypes)

	res = __gong__appendInstances(res, stage.CodeOxygenTypes)

	res = __gong__appendInstances(res, stage.CodePARTypes)

	res = __gong__appendInstances(res, stage.CodePCNMethodTypes)

	res = __gong__appendInstances(res, stage.CodePCNPavementTypes)

	res = __gong__appendInstances(res, stage.CodePCNSubgradeTypes)

	res = __gong__appendInstances(res, stage.CodePCNTyrePressureTypes)

	res = __gong__appendInstances(res, stage.CodePassengerServiceTypes)

	res = __gong__appendInstances(res, stage.CodePilotControlledLightingTypes)

	res = __gong__appendInstances(res, stage.CodePositionInILSTypes)

	res = __gong__appendInstances(res, stage.CodePrimaryRadarTypes)

	res = __gong__appendInstances(res, stage.CodeProcedureAvailabilityTypes)

	res = __gong__appendInstances(res, stage.CodeProcedureCodingStandardTypes)

	res = __gong__appendInstances(res, stage.CodeProcedureDistanceTypes)

	res = __gong__appendInstances(res, stage.CodeProcedureFixRoleTypes)

	res = __gong__appendInstances(res, stage.CodeProcedurePhaseTypes)

	res = __gong__appendInstances(res, stage.CodeProtectAreaSectionTypes)

	res = __gong__appendInstances(res, stage.CodeRNPTypes)

	res = __gong__appendInstances(res, stage.CodeRVRReadingTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMPointRoleTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMTypes)

	res = __gong__appendInstances(res, stage.CodeRadarServiceTypes)

	res = __gong__appendInstances(res, stage.CodeRadioEmissionTypes)

	res = __gong__appendInstances(res, stage.CodeRadioFrequencyAreaTypes)

	res = __gong__appendInstances(res, stage.CodeRadioSignalTypes)

	res = __gong__appendInstances(res, stage.CodeReferencePathIdentifierTypes)

	res = __gong__appendInstances(res, stage.CodeReferenceRoleTypes)

	res = __gong__appendInstances(res, stage.CodeReflectorTypes)

	res = __gong__appendInstances(res, stage.CodeRelativePositionTypes)

	res = __gong__appendInstances(res, stage.CodeRoadTypes)

	res = __gong__appendInstances(res, stage.CodeRouteAvailabilityTypes)

	res = __gong__appendInstances(res, stage.CodeRouteDesignatorLetterTypes)

	res = __gong__appendInstances(res, stage.CodeRouteDesignatorPrefixTypes)

	res = __gong__appendInstances(res, stage.CodeRouteDesignatorSuffixTypes)

	res = __gong__appendInstances(res, stage.CodeRouteIndicatorTypes)

	res = __gong__appendInstances(res, stage.CodeRouteNavigationTypes)

	res = __gong__appendInstances(res, stage.CodeRouteOriginTypes)

	res = __gong__appendInstances(res, stage.CodeRouteSegmentPathTypes)

	res = __gong__appendInstances(res, stage.CodeRouteTypes)

	res = __gong__appendInstances(res, stage.CodeRuleProcedureTitleTypes)

	res = __gong__appendInstances(res, stage.CodeRuleProcedureTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayElementTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayMarkingTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayPointRoleTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayProtectionAreaTypes)

	res = __gong__appendInstances(res, stage.CodeRunwaySectionTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayTypes)

	res = __gong__appendInstances(res, stage.CodeSafeAltitudeTypes)

	res = __gong__appendInstances(res, stage.CodeSegmentPathTypes)

	res = __gong__appendInstances(res, stage.CodeSegmentTerminationTypes)

	res = __gong__appendInstances(res, stage.CodeServiceATCTypes)

	res = __gong__appendInstances(res, stage.CodeServiceATFMTypes)

	res = __gong__appendInstances(res, stage.CodeServiceGroundControlTypes)

	res = __gong__appendInstances(res, stage.CodeServiceInformationTypes)

	res = __gong__appendInstances(res, stage.CodeServiceSARTypes)

	res = __gong__appendInstances(res, stage.CodeSideTypes)

	res = __gong__appendInstances(res, stage.CodeSignalPerformanceILSTypes)

	res = __gong__appendInstances(res, stage.CodeSpecialDateTypes)

	res = __gong__appendInstances(res, stage.CodeSpecialNavigationChainDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeSpecialNavigationStationTypes)

	res = __gong__appendInstances(res, stage.CodeSpecialNavigationSystemTypes)

	res = __gong__appendInstances(res, stage.CodeSpeedReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeStandbyPowerTypes)

	res = __gong__appendInstances(res, stage.CodeStatusAirportTypes)

	res = __gong__appendInstances(res, stage.CodeStatusAirspaceTypes)

	res = __gong__appendInstances(res, stage.CodeStatusConstructionTypes)

	res = __gong__appendInstances(res, stage.CodeStatusNavaidTypes)

	res = __gong__appendInstances(res, stage.CodeStatusOperationsTypes)

	res = __gong__appendInstances(res, stage.CodeStatusServiceTypes)

	res = __gong__appendInstances(res, stage.CodeSurfaceCompositionTypes)

	res = __gong__appendInstances(res, stage.CodeSurfaceConditionTypes)

	res = __gong__appendInstances(res, stage.CodeSurfacePreparationTypes)

	res = __gong__appendInstances(res, stage.CodeSystemActivationTypes)

	res = __gong__appendInstances(res, stage.CodeTAATypes)

	res = __gong__appendInstances(res, stage.CodeTACANChannelTypes)

	res = __gong__appendInstances(res, stage.CodeTLOFSectionTypes)

	res = __gong__appendInstances(res, stage.CodeTaxiwayElementTypes)

	res = __gong__appendInstances(res, stage.CodeTaxiwaySectionTypes)

	res = __gong__appendInstances(res, stage.CodeTaxiwayTypes)

	res = __gong__appendInstances(res, stage.CodeTelecomNetworkTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventCombinationTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventTypes)

	res = __gong__appendInstances(res, stage.CodeTimeReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeTrajectoryTypes)

	res = __gong__appendInstances(res, stage.CodeTransponderTypes)

	res = __gong__appendInstances(res, stage.CodeUUIDTypes)

	res = __gong__appendInstances(res, stage.CodeUnitDependencyTypes)

	res = __gong__appendInstances(res, stage.CodeUnitTypes)

	res = __gong__appendInstances(res, stage.CodeUpperAlphaTypes)

	res = __gong__appendInstances(res, stage.CodeUsageLimitationTypes)

	res = __gong__appendInstances(res, stage.CodeVASISTypes)

	res = __gong__appendInstances(res, stage.CodeVORTypes)

	res = __gong__appendInstances(res, stage.CodeValueInterpretationTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalDatumTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalStructureMarkingTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalStructureMaterialTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalStructureTypes)

	res = __gong__appendInstances(res, stage.CodeVisualDockingGuidanceTypes)

	res = __gong__appendInstances(res, stage.CodeWakeTurbulenceTypes)

	res = __gong__appendInstances(res, stage.CodeWorkAreaTypes)

	res = __gong__appendInstances(res, stage.CodeYesNoTypes)

	res = __gong__appendInstances(res, stage.DateMonthDayTypes)

	res = __gong__appendInstances(res, stage.DateTimeTypes)

	res = __gong__appendInstances(res, stage.DateTypes)

	res = __gong__appendInstances(res, stage.DateYearTypes)

	res = __gong__appendInstances(res, stage.NoNumberTypes)

	res = __gong__appendInstances(res, stage.NoSequenceTypes)

	res = __gong__appendInstances(res, stage.TextAddressTypes)

	res = __gong__appendInstances(res, stage.TextDesignatorTypes)

	res = __gong__appendInstances(res, stage.TextInstructionTypes)

	res = __gong__appendInstances(res, stage.TextNameTypes)

	res = __gong__appendInstances(res, stage.TextNoteTypes)

	res = __gong__appendInstances(res, stage.TextPhoneTypes)

	res = __gong__appendInstances(res, stage.TextPropertyNameTypes)

	res = __gong__appendInstances(res, stage.TextRemarkTypes)

	res = __gong__appendInstances(res, stage.TextSIDSTARDesignatorTypes)

	res = __gong__appendInstances(res, stage.TimeTypes)

	res = __gong__appendInstances(res, stage.ValAlarmLimitTypes)

	res = __gong__appendInstances(res, stage.ValAngleTypes)

	res = __gong__appendInstances(res, stage.ValBearingTypes)

	res = __gong__appendInstances(res, stage.ValChannelNumberTypes)

	res = __gong__appendInstances(res, stage.ValDepthTypes)

	res = __gong__appendInstances(res, stage.ValDistanceSignedTypes)

	res = __gong__appendInstances(res, stage.ValDistanceTypes)

	res = __gong__appendInstances(res, stage.ValDistanceVerticalTypes)

	res = __gong__appendInstances(res, stage.ValDurationTypes)

	res = __gong__appendInstances(res, stage.ValFLTypes)

	res = __gong__appendInstances(res, stage.ValFrequencyTypes)

	res = __gong__appendInstances(res, stage.ValFrictionTypes)

	res = __gong__appendInstances(res, stage.ValHexTypes)

	res = __gong__appendInstances(res, stage.ValLCNTypes)

	res = __gong__appendInstances(res, stage.ValLightIntensityTypes)

	res = __gong__appendInstances(res, stage.ValMagneticVariationChangeTypes)

	res = __gong__appendInstances(res, stage.ValMagneticVariationTypes)

	res = __gong__appendInstances(res, stage.ValPCNTypes)

	res = __gong__appendInstances(res, stage.ValPercentTypes)

	res = __gong__appendInstances(res, stage.ValPressureTypes)

	res = __gong__appendInstances(res, stage.ValSlopeTypes)

	res = __gong__appendInstances(res, stage.ValSpeedTypes)

	res = __gong__appendInstances(res, stage.ValTemperatureTypes)

	res = __gong__appendInstances(res, stage.ValWeightTypes)

	res = __gong__appendInstances(res, stage.XHTMLTypes)

	return
}

// insertion point per named struct
func (codeatcreportingtype *CodeATCReportingType) GongCopy() GongstructIF {
	newInstance := new(CodeATCReportingType)
	codeatcreportingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongCopy() GongstructIF {
	newInstance := new(CodeAerialRefuellingPointType)
	codeaerialrefuellingpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongCopy() GongstructIF {
	newInstance := new(CodeAerialRefuellingPrefixType)
	codeaerialrefuellingprefixtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongCopy() GongstructIF {
	newInstance := new(CodeAerialRefuellingType)
	codeaerialrefuellingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftCategoryType)
	codeaircraftcategorytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftEngineNumberType)
	codeaircraftenginenumbertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftEngineType)
	codeaircraftenginetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftGroundServiceType)
	codeaircraftgroundservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftICAOType)
	codeaircrafticaotype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftstandtype *CodeAircraftStandType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftStandType)
	codeaircraftstandtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircrafttype *CodeAircraftType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftType)
	codeaircrafttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftWingspanClassType)
	codeaircraftwingspanclasstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongCopy() GongstructIF {
	newInstance := new(CodeAirportHeliportCollocationType)
	codeairportheliportcollocationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeAirportHeliportDesignatorType)
	codeairportheliportdesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairportheliporttype *CodeAirportHeliportType) GongCopy() GongstructIF {
	newInstance := new(CodeAirportHeliportType)
	codeairportheliporttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairportwarningtype *CodeAirportWarningType) GongCopy() GongstructIF {
	newInstance := new(CodeAirportWarningType)
	codeairportwarningtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceActivityType)
	codeairspaceactivitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceAggregationType)
	codeairspaceaggregationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceClassificationType)
	codeairspaceclassificationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceDependencyType)
	codeairspacedependencytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceDesignatorType)
	codeairspacedesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspacePointPositionType)
	codeairspacepointpositiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspacePointRoleType)
	codeairspacepointroletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeairspacetype *CodeAirspaceType) GongCopy() GongstructIF {
	newInstance := new(CodeAirspaceType)
	codeairspacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongCopy() GongstructIF {
	newInstance := new(CodeAltitudeAdjustmentType)
	codealtitudeadjustmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codealtitudeusetype *CodeAltitudeUseType) GongCopy() GongstructIF {
	newInstance := new(CodeAltitudeUseType)
	codealtitudeusetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongCopy() GongstructIF {
	newInstance := new(CodeApproachEquipmentAdditionalType)
	codeapproachequipmentadditionaltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongCopy() GongstructIF {
	newInstance := new(CodeApproachGuidanceType)
	codeapproachguidancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongCopy() GongstructIF {
	newInstance := new(CodeApproachLightingICAOType)
	codeapproachlightingicaotype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapproachlightingtype *CodeApproachLightingType) GongCopy() GongstructIF {
	newInstance := new(CodeApproachLightingType)
	codeapproachlightingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongCopy() GongstructIF {
	newInstance := new(CodeApproachPrefixType)
	codeapproachprefixtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapproachtype *CodeApproachType) GongCopy() GongstructIF {
	newInstance := new(CodeApproachType)
	codeapproachtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapprovaltype *CodeApprovalType) GongCopy() GongstructIF {
	newInstance := new(CodeApprovalType)
	codeapprovaltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapronelementtype *CodeApronElementType) GongCopy() GongstructIF {
	newInstance := new(CodeApronElementType)
	codeapronelementtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeapronsectiontype *CodeApronSectionType) GongCopy() GongstructIF {
	newInstance := new(CodeApronSectionType)
	codeapronsectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codearcdirectiontype *CodeArcDirectionType) GongCopy() GongstructIF {
	newInstance := new(CodeArcDirectionType)
	codearcdirectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongCopy() GongstructIF {
	newInstance := new(CodeArrestingGearEnergyAbsorbType)
	codearrestinggearenergyabsorbtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongCopy() GongstructIF {
	newInstance := new(CodeArrestingGearEngageDeviceType)
	codearrestinggearengagedevicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeauralmorsetype *CodeAuralMorseType) GongCopy() GongstructIF {
	newInstance := new(CodeAuralMorseType)
	codeauralmorsetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongCopy() GongstructIF {
	newInstance := new(CodeAuthorityRoleType)
	codeauthorityroletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeauthoritytype *CodeAuthorityType) GongCopy() GongstructIF {
	newInstance := new(CodeAuthorityType)
	codeauthoritytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongCopy() GongstructIF {
	newInstance := new(CodeAviationStandardsType)
	codeaviationstandardstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codebearingtype *CodeBearingType) GongCopy() GongstructIF {
	newInstance := new(CodeBearingType)
	codebearingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeBuoyDesignatorType)
	codebuoydesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codebuoytype *CodeBuoyType) GongCopy() GongstructIF {
	newInstance := new(CodeBuoyType)
	codebuoytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongCopy() GongstructIF {
	newInstance := new(CodeCardinalDirectionType)
	codecardinaldirectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongCopy() GongstructIF {
	newInstance := new(CodeCheckpointCategoryType)
	codecheckpointcategorytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecolourtype *CodeColourType) GongCopy() GongstructIF {
	newInstance := new(CodeColourType)
	codecolourtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongCopy() GongstructIF {
	newInstance := new(CodeCommunicationChannelType)
	codecommunicationchanneltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongCopy() GongstructIF {
	newInstance := new(CodeCommunicationDirectionType)
	codecommunicationdirectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongCopy() GongstructIF {
	newInstance := new(CodeCommunicationModeType)
	codecommunicationmodetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecomparisontype *CodeComparisonType) GongCopy() GongstructIF {
	newInstance := new(CodeComparisonType)
	codecomparisontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecontaminationtype *CodeContaminationType) GongCopy() GongstructIF {
	newInstance := new(CodeContaminationType)
	codecontaminationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongCopy() GongstructIF {
	newInstance := new(CodeCourseQualityILSType)
	codecoursequalityilstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecoursetype *CodeCourseType) GongCopy() GongstructIF {
	newInstance := new(CodeCourseType)
	codecoursetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedmechanneltype *CodeDMEChannelType) GongCopy() GongstructIF {
	newInstance := new(CodeDMEChannelType)
	codedmechanneltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedmetype *CodeDMEType) GongCopy() GongstructIF {
	newInstance := new(CodeDMEType)
	codedmetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedaytype *CodeDayType) GongCopy() GongstructIF {
	newInstance := new(CodeDayType)
	codedaytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongCopy() GongstructIF {
	newInstance := new(CodeDeclaredDistanceType)
	codedeclareddistancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedesignstandardtype *CodeDesignStandardType) GongCopy() GongstructIF {
	newInstance := new(CodeDesignStandardType)
	codedesignstandardtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeDesignatedPointDesignatorType)
	codedesignatedpointdesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongCopy() GongstructIF {
	newInstance := new(CodeDesignatedPointType)
	codedesignatedpointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeDirectionReferenceType)
	codedirectionreferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedirectionturntype *CodeDirectionTurnType) GongCopy() GongstructIF {
	newInstance := new(CodeDirectionTurnType)
	codedirectionturntype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedirectiontype *CodeDirectionType) GongCopy() GongstructIF {
	newInstance := new(CodeDirectionType)
	codedirectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongCopy() GongstructIF {
	newInstance := new(CodeDistanceIndicationType)
	codedistanceindicationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeemissionbandtype *CodeEmissionBandType) GongCopy() GongstructIF {
	newInstance := new(CodeEmissionBandType)
	codeemissionbandtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongCopy() GongstructIF {
	newInstance := new(CodeEquipmentAntiCollisionType)
	codeequipmentanticollisiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongCopy() GongstructIF {
	newInstance := new(CodeEquipmentUnavailableType)
	codeequipmentunavailabletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongCopy() GongstructIF {
	newInstance := new(CodeFacilityRankingType)
	codefacilityrankingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongCopy() GongstructIF {
	newInstance := new(CodeFinalGuidanceType)
	codefinalguidancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefirefightingtype *CodeFireFightingType) GongCopy() GongstructIF {
	newInstance := new(CodeFireFightingType)
	codefirefightingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightDestinationType)
	codeflightdestinationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflightorigintype *CodeFlightOriginType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightOriginType)
	codeflightorigintype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflightpurposetype *CodeFlightPurposeType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightPurposeType)
	codeflightpurposetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightRestrictionDesignatorType)
	codeflightrestrictiondesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightRestrictionType)
	codeflightrestrictiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflightruletype *CodeFlightRuleType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightRuleType)
	codeflightruletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflightstatustype *CodeFlightStatusType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightStatusType)
	codeflightstatustype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflighttype *CodeFlightType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightType)
	codeflighttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongCopy() GongstructIF {
	newInstance := new(CodeFlowConditionOperationType)
	codeflowconditionoperationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefreeflighttype *CodeFreeFlightType) GongCopy() GongstructIF {
	newInstance := new(CodeFreeFlightType)
	codefreeflighttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongCopy() GongstructIF {
	newInstance := new(CodeFrictionDeviceType)
	codefrictiondevicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongCopy() GongstructIF {
	newInstance := new(CodeFrictionEstimateType)
	codefrictionestimatetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefueltype *CodeFuelType) GongCopy() GongstructIF {
	newInstance := new(CodeFuelType)
	codefueltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codegeobordertype *CodeGeoBorderType) GongCopy() GongstructIF {
	newInstance := new(CodeGeoBorderType)
	codegeobordertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codegradeseparationtype *CodeGradeSeparationType) GongCopy() GongstructIF {
	newInstance := new(CodeGradeSeparationType)
	codegradeseparationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codegroundlightingtype *CodeGroundLightingType) GongCopy() GongstructIF {
	newInstance := new(CodeGroundLightingType)
	codegroundlightingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeguidancelinetype *CodeGuidanceLineType) GongCopy() GongstructIF {
	newInstance := new(CodeGuidanceLineType)
	codeguidancelinetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeheightreferencetype *CodeHeightReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeHeightReferenceType)
	codeheightreferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongCopy() GongstructIF {
	newInstance := new(CodeHelicopterPerformanceType)
	codehelicopterperformancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongCopy() GongstructIF {
	newInstance := new(CodeHoldingCategoryType)
	codeholdingcategorytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeholdingusagetype *CodeHoldingUsageType) GongCopy() GongstructIF {
	newInstance := new(CodeHoldingUsageType)
	codeholdingusagetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeholdingusetype *CodeHoldingUseType) GongCopy() GongstructIF {
	newInstance := new(CodeHoldingUseType)
	codeholdingusetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeiatatype *CodeIATAType) GongCopy() GongstructIF {
	newInstance := new(CodeIATAType)
	codeiatatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeicaocountrytype *CodeICAOCountryType) GongCopy() GongstructIF {
	newInstance := new(CodeICAOCountryType)
	codeicaocountrytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeicaotype *CodeICAOType) GongCopy() GongstructIF {
	newInstance := new(CodeICAOType)
	codeicaotype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongCopy() GongstructIF {
	newInstance := new(CodeILSBackCourseType)
	codeilsbackcoursetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongCopy() GongstructIF {
	newInstance := new(CodeIntegrityLevelILSType)
	codeintegritylevelilstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongCopy() GongstructIF {
	newInstance := new(CodeIntensityStandByType)
	codeintensitystandbytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelanguagetype *CodeLanguageType) GongCopy() GongstructIF {
	newInstance := new(CodeLanguageType)
	codelanguagetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelevelseriestype *CodeLevelSeriesType) GongCopy() GongstructIF {
	newInstance := new(CodeLevelSeriesType)
	codelevelseriestype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeLevelTableDesignatorType)
	codeleveltabledesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeleveltype *CodeLevelType) GongCopy() GongstructIF {
	newInstance := new(CodeLevelType)
	codeleveltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongCopy() GongstructIF {
	newInstance := new(CodeLightHoldingPositionType)
	codelightholdingpositiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelightintensitytype *CodeLightIntensityType) GongCopy() GongstructIF {
	newInstance := new(CodeLightIntensityType)
	codelightintensitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelightsourcetype *CodeLightSourceType) GongCopy() GongstructIF {
	newInstance := new(CodeLightSourceType)
	codelightsourcetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelightingjartype *CodeLightingJARType) GongCopy() GongstructIF {
	newInstance := new(CodeLightingJARType)
	codelightingjartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongCopy() GongstructIF {
	newInstance := new(CodeLoadingBridgeType)
	codeloadingbridgetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongCopy() GongstructIF {
	newInstance := new(CodeLocationQualifierType)
	codelocationqualifiertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongCopy() GongstructIF {
	newInstance := new(CodeLogicalOperatorType)
	codelogicaloperatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongCopy() GongstructIF {
	newInstance := new(CodeMLSAzimuthType)
	codemlsazimuthtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemlschanneltype *CodeMLSChannelType) GongCopy() GongstructIF {
	newInstance := new(CodeMLSChannelType)
	codemlschanneltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongCopy() GongstructIF {
	newInstance := new(CodeMarkerBeaconSignalType)
	codemarkerbeaconsignaltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongCopy() GongstructIF {
	newInstance := new(CodeMarkingConditionType)
	codemarkingconditiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemarkingstyletype *CodeMarkingStyleType) GongCopy() GongstructIF {
	newInstance := new(CodeMarkingStyleType)
	codemarkingstyletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongCopy() GongstructIF {
	newInstance := new(CodeMeteoConditionsType)
	codemeteoconditionstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongCopy() GongstructIF {
	newInstance := new(CodeMilitaryOperationsType)
	codemilitaryoperationstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongCopy() GongstructIF {
	newInstance := new(CodeMilitaryRoutePointType)
	codemilitaryroutepointtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongCopy() GongstructIF {
	newInstance := new(CodeMilitaryStatusType)
	codemilitarystatustype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongCopy() GongstructIF {
	newInstance := new(CodeMilitaryTrainingType)
	codemilitarytrainingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongCopy() GongstructIF {
	newInstance := new(CodeMinimaFinalApproachPathType)
	codeminimafinalapproachpathtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongCopy() GongstructIF {
	newInstance := new(CodeMinimumAltitudeType)
	codeminimumaltitudetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeminimumheighttype *CodeMinimumHeightType) GongCopy() GongstructIF {
	newInstance := new(CodeMinimumHeightType)
	codeminimumheighttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemissedapproachtype *CodeMissedApproachType) GongCopy() GongstructIF {
	newInstance := new(CodeMissedApproachType)
	codemissedapproachtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codendbusagetype *CodeNDBUsageType) GongCopy() GongstructIF {
	newInstance := new(CodeNDBUsageType)
	codendbusagetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeNavaidDesignatorType)
	codenavaiddesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongCopy() GongstructIF {
	newInstance := new(CodeNavaidPurposeType)
	codenavaidpurposetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavaidservicetype *CodeNavaidServiceType) GongCopy() GongstructIF {
	newInstance := new(CodeNavaidServiceType)
	codenavaidservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongCopy() GongstructIF {
	newInstance := new(CodeNavigationAreaRestrictionType)
	codenavigationarearestrictiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavigationareatype *CodeNavigationAreaType) GongCopy() GongstructIF {
	newInstance := new(CodeNavigationAreaType)
	codenavigationareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongCopy() GongstructIF {
	newInstance := new(CodeNavigationEquipmentType)
	codenavigationequipmenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongCopy() GongstructIF {
	newInstance := new(CodeNavigationSpecificationType)
	codenavigationspecificationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenitrogentype *CodeNitrogenType) GongCopy() GongstructIF {
	newInstance := new(CodeNitrogenType)
	codenitrogentype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenorthreferencetype *CodeNorthReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeNorthReferenceType)
	codenorthreferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codenotepurposetype *CodeNotePurposeType) GongCopy() GongstructIF {
	newInstance := new(CodeNotePurposeType)
	codenotepurposetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeobstacleareatype *CodeObstacleAreaType) GongCopy() GongstructIF {
	newInstance := new(CodeObstacleAreaType)
	codeobstacleareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongCopy() GongstructIF {
	newInstance := new(CodeObstacleAssessmentSurfaceType)
	codeobstacleassessmentsurfacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongCopy() GongstructIF {
	newInstance := new(CodeObstructionIdSurfaceZoneType)
	codeobstructionidsurfacezonetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeoiltype *CodeOilType) GongCopy() GongstructIF {
	newInstance := new(CodeOilType)
	codeoiltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongCopy() GongstructIF {
	newInstance := new(CodeOperationAirportHeliportType)
	codeoperationairportheliporttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongCopy() GongstructIF {
	newInstance := new(CodeOperationManoeuvringAreaType)
	codeoperationmanoeuvringareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeOrganisationDesignatorType)
	codeorganisationdesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongCopy() GongstructIF {
	newInstance := new(CodeOrganisationHierarchyType)
	codeorganisationhierarchytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeorganisationtype *CodeOrganisationType) GongCopy() GongstructIF {
	newInstance := new(CodeOrganisationType)
	codeorganisationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeoxygentype *CodeOxygenType) GongCopy() GongstructIF {
	newInstance := new(CodeOxygenType)
	codeoxygentype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codepartype *CodePARType) GongCopy() GongstructIF {
	newInstance := new(CodePARType)
	codepartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codepcnmethodtype *CodePCNMethodType) GongCopy() GongstructIF {
	newInstance := new(CodePCNMethodType)
	codepcnmethodtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codepcnpavementtype *CodePCNPavementType) GongCopy() GongstructIF {
	newInstance := new(CodePCNPavementType)
	codepcnpavementtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongCopy() GongstructIF {
	newInstance := new(CodePCNSubgradeType)
	codepcnsubgradetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongCopy() GongstructIF {
	newInstance := new(CodePCNTyrePressureType)
	codepcntyrepressuretype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codepassengerservicetype *CodePassengerServiceType) GongCopy() GongstructIF {
	newInstance := new(CodePassengerServiceType)
	codepassengerservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongCopy() GongstructIF {
	newInstance := new(CodePilotControlledLightingType)
	codepilotcontrolledlightingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codepositioninilstype *CodePositionInILSType) GongCopy() GongstructIF {
	newInstance := new(CodePositionInILSType)
	codepositioninilstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeprimaryradartype *CodePrimaryRadarType) GongCopy() GongstructIF {
	newInstance := new(CodePrimaryRadarType)
	codeprimaryradartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongCopy() GongstructIF {
	newInstance := new(CodeProcedureAvailabilityType)
	codeprocedureavailabilitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongCopy() GongstructIF {
	newInstance := new(CodeProcedureCodingStandardType)
	codeprocedurecodingstandardtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongCopy() GongstructIF {
	newInstance := new(CodeProcedureDistanceType)
	codeproceduredistancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongCopy() GongstructIF {
	newInstance := new(CodeProcedureFixRoleType)
	codeprocedurefixroletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongCopy() GongstructIF {
	newInstance := new(CodeProcedurePhaseType)
	codeprocedurephasetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongCopy() GongstructIF {
	newInstance := new(CodeProtectAreaSectionType)
	codeprotectareasectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codernptype *CodeRNPType) GongCopy() GongstructIF {
	newInstance := new(CodeRNPType)
	codernptype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codervrreadingtype *CodeRVRReadingType) GongCopy() GongstructIF {
	newInstance := new(CodeRVRReadingType)
	codervrreadingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongCopy() GongstructIF {
	newInstance := new(CodeRVSMPointRoleType)
	codervsmpointroletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codervsmtype *CodeRVSMType) GongCopy() GongstructIF {
	newInstance := new(CodeRVSMType)
	codervsmtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderadarservicetype *CodeRadarServiceType) GongCopy() GongstructIF {
	newInstance := new(CodeRadarServiceType)
	coderadarservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderadioemissiontype *CodeRadioEmissionType) GongCopy() GongstructIF {
	newInstance := new(CodeRadioEmissionType)
	coderadioemissiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongCopy() GongstructIF {
	newInstance := new(CodeRadioFrequencyAreaType)
	coderadiofrequencyareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderadiosignaltype *CodeRadioSignalType) GongCopy() GongstructIF {
	newInstance := new(CodeRadioSignalType)
	coderadiosignaltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongCopy() GongstructIF {
	newInstance := new(CodeReferencePathIdentifierType)
	codereferencepathidentifiertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codereferenceroletype *CodeReferenceRoleType) GongCopy() GongstructIF {
	newInstance := new(CodeReferenceRoleType)
	codereferenceroletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codereflectortype *CodeReflectorType) GongCopy() GongstructIF {
	newInstance := new(CodeReflectorType)
	codereflectortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderelativepositiontype *CodeRelativePositionType) GongCopy() GongstructIF {
	newInstance := new(CodeRelativePositionType)
	coderelativepositiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderoadtype *CodeRoadType) GongCopy() GongstructIF {
	newInstance := new(CodeRoadType)
	coderoadtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteAvailabilityType)
	coderouteavailabilitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteDesignatorLetterType)
	coderoutedesignatorlettertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteDesignatorPrefixType)
	coderoutedesignatorprefixtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteDesignatorSuffixType)
	coderoutedesignatorsuffixtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteIndicatorType)
	coderouteindicatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteNavigationType)
	coderoutenavigationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderouteorigintype *CodeRouteOriginType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteOriginType)
	coderouteorigintype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteSegmentPathType)
	coderoutesegmentpathtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderoutetype *CodeRouteType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteType)
	coderoutetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongCopy() GongstructIF {
	newInstance := new(CodeRuleProcedureTitleType)
	coderuleproceduretitletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderuleproceduretype *CodeRuleProcedureType) GongCopy() GongstructIF {
	newInstance := new(CodeRuleProcedureType)
	coderuleproceduretype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderunwayelementtype *CodeRunwayElementType) GongCopy() GongstructIF {
	newInstance := new(CodeRunwayElementType)
	coderunwayelementtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongCopy() GongstructIF {
	newInstance := new(CodeRunwayMarkingType)
	coderunwaymarkingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongCopy() GongstructIF {
	newInstance := new(CodeRunwayPointRoleType)
	coderunwaypointroletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongCopy() GongstructIF {
	newInstance := new(CodeRunwayProtectionAreaType)
	coderunwayprotectionareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongCopy() GongstructIF {
	newInstance := new(CodeRunwaySectionType)
	coderunwaysectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderunwaytype *CodeRunwayType) GongCopy() GongstructIF {
	newInstance := new(CodeRunwayType)
	coderunwaytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongCopy() GongstructIF {
	newInstance := new(CodeSafeAltitudeType)
	codesafealtitudetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesegmentpathtype *CodeSegmentPathType) GongCopy() GongstructIF {
	newInstance := new(CodeSegmentPathType)
	codesegmentpathtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongCopy() GongstructIF {
	newInstance := new(CodeSegmentTerminationType)
	codesegmentterminationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeserviceatctype *CodeServiceATCType) GongCopy() GongstructIF {
	newInstance := new(CodeServiceATCType)
	codeserviceatctype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeserviceatfmtype *CodeServiceATFMType) GongCopy() GongstructIF {
	newInstance := new(CodeServiceATFMType)
	codeserviceatfmtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongCopy() GongstructIF {
	newInstance := new(CodeServiceGroundControlType)
	codeservicegroundcontroltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeserviceinformationtype *CodeServiceInformationType) GongCopy() GongstructIF {
	newInstance := new(CodeServiceInformationType)
	codeserviceinformationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeservicesartype *CodeServiceSARType) GongCopy() GongstructIF {
	newInstance := new(CodeServiceSARType)
	codeservicesartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesidetype *CodeSideType) GongCopy() GongstructIF {
	newInstance := new(CodeSideType)
	codesidetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongCopy() GongstructIF {
	newInstance := new(CodeSignalPerformanceILSType)
	codesignalperformanceilstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codespecialdatetype *CodeSpecialDateType) GongCopy() GongstructIF {
	newInstance := new(CodeSpecialDateType)
	codespecialdatetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongCopy() GongstructIF {
	newInstance := new(CodeSpecialNavigationChainDesignatorType)
	codespecialnavigationchaindesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongCopy() GongstructIF {
	newInstance := new(CodeSpecialNavigationStationType)
	codespecialnavigationstationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongCopy() GongstructIF {
	newInstance := new(CodeSpecialNavigationSystemType)
	codespecialnavigationsystemtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeSpeedReferenceType)
	codespeedreferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestandbypowertype *CodeStandbyPowerType) GongCopy() GongstructIF {
	newInstance := new(CodeStandbyPowerType)
	codestandbypowertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusairporttype *CodeStatusAirportType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusAirportType)
	codestatusairporttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusAirspaceType)
	codestatusairspacetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusConstructionType)
	codestatusconstructiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusNavaidType)
	codestatusnavaidtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusoperationstype *CodeStatusOperationsType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusOperationsType)
	codestatusoperationstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusservicetype *CodeStatusServiceType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusServiceType)
	codestatusservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongCopy() GongstructIF {
	newInstance := new(CodeSurfaceCompositionType)
	codesurfacecompositiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongCopy() GongstructIF {
	newInstance := new(CodeSurfaceConditionType)
	codesurfaceconditiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongCopy() GongstructIF {
	newInstance := new(CodeSurfacePreparationType)
	codesurfacepreparationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesystemactivationtype *CodeSystemActivationType) GongCopy() GongstructIF {
	newInstance := new(CodeSystemActivationType)
	codesystemactivationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetaatype *CodeTAAType) GongCopy() GongstructIF {
	newInstance := new(CodeTAAType)
	codetaatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetacanchanneltype *CodeTACANChannelType) GongCopy() GongstructIF {
	newInstance := new(CodeTACANChannelType)
	codetacanchanneltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetlofsectiontype *CodeTLOFSectionType) GongCopy() GongstructIF {
	newInstance := new(CodeTLOFSectionType)
	codetlofsectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongCopy() GongstructIF {
	newInstance := new(CodeTaxiwayElementType)
	codetaxiwayelementtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongCopy() GongstructIF {
	newInstance := new(CodeTaxiwaySectionType)
	codetaxiwaysectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetaxiwaytype *CodeTaxiwayType) GongCopy() GongstructIF {
	newInstance := new(CodeTaxiwayType)
	codetaxiwaytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongCopy() GongstructIF {
	newInstance := new(CodeTelecomNetworkType)
	codetelecomnetworktype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongCopy() GongstructIF {
	newInstance := new(CodeTimeEventCombinationType)
	codetimeeventcombinationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetimeeventtype *CodeTimeEventType) GongCopy() GongstructIF {
	newInstance := new(CodeTimeEventType)
	codetimeeventtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetimereferencetype *CodeTimeReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeTimeReferenceType)
	codetimereferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetrajectorytype *CodeTrajectoryType) GongCopy() GongstructIF {
	newInstance := new(CodeTrajectoryType)
	codetrajectorytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetranspondertype *CodeTransponderType) GongCopy() GongstructIF {
	newInstance := new(CodeTransponderType)
	codetranspondertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeuuidtype *CodeUUIDType) GongCopy() GongstructIF {
	newInstance := new(CodeUUIDType)
	codeuuidtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeunitdependencytype *CodeUnitDependencyType) GongCopy() GongstructIF {
	newInstance := new(CodeUnitDependencyType)
	codeunitdependencytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeunittype *CodeUnitType) GongCopy() GongstructIF {
	newInstance := new(CodeUnitType)
	codeunittype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeupperalphatype *CodeUpperAlphaType) GongCopy() GongstructIF {
	newInstance := new(CodeUpperAlphaType)
	codeupperalphatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongCopy() GongstructIF {
	newInstance := new(CodeUsageLimitationType)
	codeusagelimitationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codevasistype *CodeVASISType) GongCopy() GongstructIF {
	newInstance := new(CodeVASISType)
	codevasistype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codevortype *CodeVORType) GongCopy() GongstructIF {
	newInstance := new(CodeVORType)
	codevortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongCopy() GongstructIF {
	newInstance := new(CodeValueInterpretationType)
	codevalueinterpretationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalDatumType)
	codeverticaldatumtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalReferenceType)
	codeverticalreferencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalStructureMarkingType)
	codeverticalstructuremarkingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalStructureMaterialType)
	codeverticalstructurematerialtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalStructureType)
	codeverticalstructuretype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongCopy() GongstructIF {
	newInstance := new(CodeVisualDockingGuidanceType)
	codevisualdockingguidancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongCopy() GongstructIF {
	newInstance := new(CodeWakeTurbulenceType)
	codewaketurbulencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeworkareatype *CodeWorkAreaType) GongCopy() GongstructIF {
	newInstance := new(CodeWorkAreaType)
	codeworkareatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeyesnotype *CodeYesNoType) GongCopy() GongstructIF {
	newInstance := new(CodeYesNoType)
	codeyesnotype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datemonthdaytype *DateMonthDayType) GongCopy() GongstructIF {
	newInstance := new(DateMonthDayType)
	datemonthdaytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datetimetype *DateTimeType) GongCopy() GongstructIF {
	newInstance := new(DateTimeType)
	datetimetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datetype *DateType) GongCopy() GongstructIF {
	newInstance := new(DateType)
	datetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dateyeartype *DateYearType) GongCopy() GongstructIF {
	newInstance := new(DateYearType)
	dateyeartype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (nonumbertype *NoNumberType) GongCopy() GongstructIF {
	newInstance := new(NoNumberType)
	nonumbertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (nosequencetype *NoSequenceType) GongCopy() GongstructIF {
	newInstance := new(NoSequenceType)
	nosequencetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textaddresstype *TextAddressType) GongCopy() GongstructIF {
	newInstance := new(TextAddressType)
	textaddresstype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textdesignatortype *TextDesignatorType) GongCopy() GongstructIF {
	newInstance := new(TextDesignatorType)
	textdesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textinstructiontype *TextInstructionType) GongCopy() GongstructIF {
	newInstance := new(TextInstructionType)
	textinstructiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textnametype *TextNameType) GongCopy() GongstructIF {
	newInstance := new(TextNameType)
	textnametype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textnotetype *TextNoteType) GongCopy() GongstructIF {
	newInstance := new(TextNoteType)
	textnotetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textphonetype *TextPhoneType) GongCopy() GongstructIF {
	newInstance := new(TextPhoneType)
	textphonetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textpropertynametype *TextPropertyNameType) GongCopy() GongstructIF {
	newInstance := new(TextPropertyNameType)
	textpropertynametype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textremarktype *TextRemarkType) GongCopy() GongstructIF {
	newInstance := new(TextRemarkType)
	textremarktype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongCopy() GongstructIF {
	newInstance := new(TextSIDSTARDesignatorType)
	textsidstardesignatortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timetype *TimeType) GongCopy() GongstructIF {
	newInstance := new(TimeType)
	timetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valalarmlimittype *ValAlarmLimitType) GongCopy() GongstructIF {
	newInstance := new(ValAlarmLimitType)
	valalarmlimittype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valangletype *ValAngleType) GongCopy() GongstructIF {
	newInstance := new(ValAngleType)
	valangletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valbearingtype *ValBearingType) GongCopy() GongstructIF {
	newInstance := new(ValBearingType)
	valbearingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valchannelnumbertype *ValChannelNumberType) GongCopy() GongstructIF {
	newInstance := new(ValChannelNumberType)
	valchannelnumbertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdepthtype *ValDepthType) GongCopy() GongstructIF {
	newInstance := new(ValDepthType)
	valdepthtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdistancesignedtype *ValDistanceSignedType) GongCopy() GongstructIF {
	newInstance := new(ValDistanceSignedType)
	valdistancesignedtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdistancetype *ValDistanceType) GongCopy() GongstructIF {
	newInstance := new(ValDistanceType)
	valdistancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongCopy() GongstructIF {
	newInstance := new(ValDistanceVerticalType)
	valdistanceverticaltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valdurationtype *ValDurationType) GongCopy() GongstructIF {
	newInstance := new(ValDurationType)
	valdurationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valfltype *ValFLType) GongCopy() GongstructIF {
	newInstance := new(ValFLType)
	valfltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valfrequencytype *ValFrequencyType) GongCopy() GongstructIF {
	newInstance := new(ValFrequencyType)
	valfrequencytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valfrictiontype *ValFrictionType) GongCopy() GongstructIF {
	newInstance := new(ValFrictionType)
	valfrictiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valhextype *ValHexType) GongCopy() GongstructIF {
	newInstance := new(ValHexType)
	valhextype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vallcntype *ValLCNType) GongCopy() GongstructIF {
	newInstance := new(ValLCNType)
	vallcntype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (vallightintensitytype *ValLightIntensityType) GongCopy() GongstructIF {
	newInstance := new(ValLightIntensityType)
	vallightintensitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongCopy() GongstructIF {
	newInstance := new(ValMagneticVariationChangeType)
	valmagneticvariationchangetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valmagneticvariationtype *ValMagneticVariationType) GongCopy() GongstructIF {
	newInstance := new(ValMagneticVariationType)
	valmagneticvariationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valpcntype *ValPCNType) GongCopy() GongstructIF {
	newInstance := new(ValPCNType)
	valpcntype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valpercenttype *ValPercentType) GongCopy() GongstructIF {
	newInstance := new(ValPercentType)
	valpercenttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valpressuretype *ValPressureType) GongCopy() GongstructIF {
	newInstance := new(ValPressureType)
	valpressuretype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valslopetype *ValSlopeType) GongCopy() GongstructIF {
	newInstance := new(ValSlopeType)
	valslopetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valspeedtype *ValSpeedType) GongCopy() GongstructIF {
	newInstance := new(ValSpeedType)
	valspeedtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valtemperaturetype *ValTemperatureType) GongCopy() GongstructIF {
	newInstance := new(ValTemperatureType)
	valtemperaturetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (valweighttype *ValWeightType) GongCopy() GongstructIF {
	newInstance := new(ValWeightType)
	valweighttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (xhtmltype *XHTMLType) GongCopy() GongstructIF {
	newInstance := new(XHTMLType)
	xhtmltype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (codeatcreportingtype *CodeATCReportingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeatcreportingtype)
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaerialrefuellingpointtype)
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaerialrefuellingprefixtype)
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaerialrefuellingtype)
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftcategorytype)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftenginenumbertype)
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftenginetype)
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftgroundservicetype)
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircrafticaotype)
}

func (codeaircraftstandtype *CodeAircraftStandType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftstandtype)
}

func (codeaircrafttype *CodeAircraftType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircrafttype)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftwingspanclasstype)
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairportheliportcollocationtype)
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairportheliportdesignatortype)
}

func (codeairportheliporttype *CodeAirportHeliportType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairportheliporttype)
}

func (codeairportwarningtype *CodeAirportWarningType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairportwarningtype)
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspaceactivitytype)
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspaceaggregationtype)
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspaceclassificationtype)
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacedependencytype)
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacedesignatortype)
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacepointpositiontype)
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacepointroletype)
}

func (codeairspacetype *CodeAirspaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacetype)
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codealtitudeadjustmenttype)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codealtitudeusetype)
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachequipmentadditionaltype)
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachguidancetype)
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachlightingicaotype)
}

func (codeapproachlightingtype *CodeApproachLightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachlightingtype)
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachprefixtype)
}

func (codeapproachtype *CodeApproachType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachtype)
}

func (codeapprovaltype *CodeApprovalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapprovaltype)
}

func (codeapronelementtype *CodeApronElementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapronelementtype)
}

func (codeapronsectiontype *CodeApronSectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapronsectiontype)
}

func (codearcdirectiontype *CodeArcDirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codearcdirectiontype)
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codearrestinggearenergyabsorbtype)
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codearrestinggearengagedevicetype)
}

func (codeauralmorsetype *CodeAuralMorseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeauralmorsetype)
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeauthorityroletype)
}

func (codeauthoritytype *CodeAuthorityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeauthoritytype)
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaviationstandardstype)
}

func (codebearingtype *CodeBearingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codebearingtype)
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codebuoydesignatortype)
}

func (codebuoytype *CodeBuoyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codebuoytype)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecardinaldirectiontype)
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecheckpointcategorytype)
}

func (codecolourtype *CodeColourType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecolourtype)
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecommunicationchanneltype)
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecommunicationdirectiontype)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecommunicationmodetype)
}

func (codecomparisontype *CodeComparisonType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecomparisontype)
}

func (codecontaminationtype *CodeContaminationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecontaminationtype)
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecoursequalityilstype)
}

func (codecoursetype *CodeCourseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecoursetype)
}

func (codedmechanneltype *CodeDMEChannelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedmechanneltype)
}

func (codedmetype *CodeDMEType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedmetype)
}

func (codedaytype *CodeDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedaytype)
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedeclareddistancetype)
}

func (codedesignstandardtype *CodeDesignStandardType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedesignstandardtype)
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedesignatedpointdesignatortype)
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedesignatedpointtype)
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedirectionreferencetype)
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedirectionturntype)
}

func (codedirectiontype *CodeDirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedirectiontype)
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedistanceindicationtype)
}

func (codeemissionbandtype *CodeEmissionBandType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeemissionbandtype)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeequipmentanticollisiontype)
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeequipmentunavailabletype)
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefacilityrankingtype)
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefinalguidancetype)
}

func (codefirefightingtype *CodeFireFightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefirefightingtype)
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightdestinationtype)
}

func (codeflightorigintype *CodeFlightOriginType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightorigintype)
}

func (codeflightpurposetype *CodeFlightPurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightpurposetype)
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightrestrictiondesignatortype)
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightrestrictiontype)
}

func (codeflightruletype *CodeFlightRuleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightruletype)
}

func (codeflightstatustype *CodeFlightStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightstatustype)
}

func (codeflighttype *CodeFlightType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflighttype)
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflowconditionoperationtype)
}

func (codefreeflighttype *CodeFreeFlightType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefreeflighttype)
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefrictiondevicetype)
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefrictionestimatetype)
}

func (codefueltype *CodeFuelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefueltype)
}

func (codegeobordertype *CodeGeoBorderType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codegeobordertype)
}

func (codegradeseparationtype *CodeGradeSeparationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codegradeseparationtype)
}

func (codegroundlightingtype *CodeGroundLightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codegroundlightingtype)
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeguidancelinetype)
}

func (codeheightreferencetype *CodeHeightReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeheightreferencetype)
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codehelicopterperformancetype)
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeholdingcategorytype)
}

func (codeholdingusagetype *CodeHoldingUsageType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeholdingusagetype)
}

func (codeholdingusetype *CodeHoldingUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeholdingusetype)
}

func (codeiatatype *CodeIATAType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeiatatype)
}

func (codeicaocountrytype *CodeICAOCountryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeicaocountrytype)
}

func (codeicaotype *CodeICAOType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeicaotype)
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeilsbackcoursetype)
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeintegritylevelilstype)
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeintensitystandbytype)
}

func (codelanguagetype *CodeLanguageType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelanguagetype)
}

func (codelevelseriestype *CodeLevelSeriesType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelevelseriestype)
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeleveltabledesignatortype)
}

func (codeleveltype *CodeLevelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeleveltype)
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelightholdingpositiontype)
}

func (codelightintensitytype *CodeLightIntensityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelightintensitytype)
}

func (codelightsourcetype *CodeLightSourceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelightsourcetype)
}

func (codelightingjartype *CodeLightingJARType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelightingjartype)
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeloadingbridgetype)
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelocationqualifiertype)
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelogicaloperatortype)
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemlsazimuthtype)
}

func (codemlschanneltype *CodeMLSChannelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemlschanneltype)
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemarkerbeaconsignaltype)
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemarkingconditiontype)
}

func (codemarkingstyletype *CodeMarkingStyleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemarkingstyletype)
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemeteoconditionstype)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemilitaryoperationstype)
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemilitaryroutepointtype)
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemilitarystatustype)
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemilitarytrainingtype)
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeminimafinalapproachpathtype)
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeminimumaltitudetype)
}

func (codeminimumheighttype *CodeMinimumHeightType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeminimumheighttype)
}

func (codemissedapproachtype *CodeMissedApproachType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemissedapproachtype)
}

func (codendbusagetype *CodeNDBUsageType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codendbusagetype)
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavaiddesignatortype)
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavaidpurposetype)
}

func (codenavaidservicetype *CodeNavaidServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavaidservicetype)
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationarearestrictiontype)
}

func (codenavigationareatype *CodeNavigationAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationareatype)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationequipmenttype)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationspecificationtype)
}

func (codenitrogentype *CodeNitrogenType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenitrogentype)
}

func (codenorthreferencetype *CodeNorthReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenorthreferencetype)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenotepurposetype)
}

func (codeobstacleareatype *CodeObstacleAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstacleareatype)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstacleassessmentsurfacetype)
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstructionidsurfacezonetype)
}

func (codeoiltype *CodeOilType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeoiltype)
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeoperationairportheliporttype)
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeoperationmanoeuvringareatype)
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeorganisationdesignatortype)
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeorganisationhierarchytype)
}

func (codeorganisationtype *CodeOrganisationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeorganisationtype)
}

func (codeoxygentype *CodeOxygenType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeoxygentype)
}

func (codepartype *CodePARType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepartype)
}

func (codepcnmethodtype *CodePCNMethodType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepcnmethodtype)
}

func (codepcnpavementtype *CodePCNPavementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepcnpavementtype)
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepcnsubgradetype)
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepcntyrepressuretype)
}

func (codepassengerservicetype *CodePassengerServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepassengerservicetype)
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepilotcontrolledlightingtype)
}

func (codepositioninilstype *CodePositionInILSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepositioninilstype)
}

func (codeprimaryradartype *CodePrimaryRadarType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeprimaryradartype)
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeprocedureavailabilitytype)
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeprocedurecodingstandardtype)
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeproceduredistancetype)
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeprocedurefixroletype)
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeprocedurephasetype)
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeprotectareasectiontype)
}

func (codernptype *CodeRNPType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codernptype)
}

func (codervrreadingtype *CodeRVRReadingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervrreadingtype)
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmpointroletype)
}

func (codervsmtype *CodeRVSMType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmtype)
}

func (coderadarservicetype *CodeRadarServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderadarservicetype)
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderadioemissiontype)
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderadiofrequencyareatype)
}

func (coderadiosignaltype *CodeRadioSignalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderadiosignaltype)
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codereferencepathidentifiertype)
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codereferenceroletype)
}

func (codereflectortype *CodeReflectorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codereflectortype)
}

func (coderelativepositiontype *CodeRelativePositionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderelativepositiontype)
}

func (coderoadtype *CodeRoadType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderoadtype)
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderouteavailabilitytype)
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderoutedesignatorlettertype)
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderoutedesignatorprefixtype)
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderoutedesignatorsuffixtype)
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderouteindicatortype)
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderoutenavigationtype)
}

func (coderouteorigintype *CodeRouteOriginType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderouteorigintype)
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderoutesegmentpathtype)
}

func (coderoutetype *CodeRouteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderoutetype)
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderuleproceduretitletype)
}

func (coderuleproceduretype *CodeRuleProcedureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderuleproceduretype)
}

func (coderunwayelementtype *CodeRunwayElementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderunwayelementtype)
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderunwaymarkingtype)
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderunwaypointroletype)
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderunwayprotectionareatype)
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderunwaysectiontype)
}

func (coderunwaytype *CodeRunwayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderunwaytype)
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesafealtitudetype)
}

func (codesegmentpathtype *CodeSegmentPathType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesegmentpathtype)
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesegmentterminationtype)
}

func (codeserviceatctype *CodeServiceATCType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeserviceatctype)
}

func (codeserviceatfmtype *CodeServiceATFMType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeserviceatfmtype)
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeservicegroundcontroltype)
}

func (codeserviceinformationtype *CodeServiceInformationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeserviceinformationtype)
}

func (codeservicesartype *CodeServiceSARType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeservicesartype)
}

func (codesidetype *CodeSideType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesidetype)
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesignalperformanceilstype)
}

func (codespecialdatetype *CodeSpecialDateType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codespecialdatetype)
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codespecialnavigationchaindesignatortype)
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codespecialnavigationstationtype)
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codespecialnavigationsystemtype)
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codespeedreferencetype)
}

func (codestandbypowertype *CodeStandbyPowerType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestandbypowertype)
}

func (codestatusairporttype *CodeStatusAirportType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusairporttype)
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusairspacetype)
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusconstructiontype)
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusnavaidtype)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusoperationstype)
}

func (codestatusservicetype *CodeStatusServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusservicetype)
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesurfacecompositiontype)
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesurfaceconditiontype)
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesurfacepreparationtype)
}

func (codesystemactivationtype *CodeSystemActivationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesystemactivationtype)
}

func (codetaatype *CodeTAAType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetaatype)
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetacanchanneltype)
}

func (codetlofsectiontype *CodeTLOFSectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetlofsectiontype)
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetaxiwayelementtype)
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetaxiwaysectiontype)
}

func (codetaxiwaytype *CodeTaxiwayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetaxiwaytype)
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetelecomnetworktype)
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetimeeventcombinationtype)
}

func (codetimeeventtype *CodeTimeEventType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetimeeventtype)
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetimereferencetype)
}

func (codetrajectorytype *CodeTrajectoryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetrajectorytype)
}

func (codetranspondertype *CodeTransponderType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetranspondertype)
}

func (codeuuidtype *CodeUUIDType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeuuidtype)
}

func (codeunitdependencytype *CodeUnitDependencyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeunitdependencytype)
}

func (codeunittype *CodeUnitType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeunittype)
}

func (codeupperalphatype *CodeUpperAlphaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeupperalphatype)
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeusagelimitationtype)
}

func (codevasistype *CodeVASISType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codevasistype)
}

func (codevortype *CodeVORType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codevortype)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codevalueinterpretationtype)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticaldatumtype)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalreferencetype)
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalstructuremarkingtype)
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalstructurematerialtype)
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalstructuretype)
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codevisualdockingguidancetype)
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codewaketurbulencetype)
}

func (codeworkareatype *CodeWorkAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeworkareatype)
}

func (codeyesnotype *CodeYesNoType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeyesnotype)
}

func (datemonthdaytype *DateMonthDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datemonthdaytype)
}

func (datetimetype *DateTimeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datetimetype)
}

func (datetype *DateType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datetype)
}

func (dateyeartype *DateYearType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dateyeartype)
}

func (nonumbertype *NoNumberType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nonumbertype)
}

func (nosequencetype *NoSequenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nosequencetype)
}

func (textaddresstype *TextAddressType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textaddresstype)
}

func (textdesignatortype *TextDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textdesignatortype)
}

func (textinstructiontype *TextInstructionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textinstructiontype)
}

func (textnametype *TextNameType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textnametype)
}

func (textnotetype *TextNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textnotetype)
}

func (textphonetype *TextPhoneType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textphonetype)
}

func (textpropertynametype *TextPropertyNameType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textpropertynametype)
}

func (textremarktype *TextRemarkType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textremarktype)
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textsidstardesignatortype)
}

func (timetype *TimeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timetype)
}

func (valalarmlimittype *ValAlarmLimitType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valalarmlimittype)
}

func (valangletype *ValAngleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valangletype)
}

func (valbearingtype *ValBearingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valbearingtype)
}

func (valchannelnumbertype *ValChannelNumberType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valchannelnumbertype)
}

func (valdepthtype *ValDepthType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdepthtype)
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdistancesignedtype)
}

func (valdistancetype *ValDistanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdistancetype)
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdistanceverticaltype)
}

func (valdurationtype *ValDurationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valdurationtype)
}

func (valfltype *ValFLType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valfltype)
}

func (valfrequencytype *ValFrequencyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valfrequencytype)
}

func (valfrictiontype *ValFrictionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valfrictiontype)
}

func (valhextype *ValHexType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valhextype)
}

func (vallcntype *ValLCNType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vallcntype)
}

func (vallightintensitytype *ValLightIntensityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, vallightintensitytype)
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valmagneticvariationchangetype)
}

func (valmagneticvariationtype *ValMagneticVariationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valmagneticvariationtype)
}

func (valpcntype *ValPCNType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valpcntype)
}

func (valpercenttype *ValPercentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valpercenttype)
}

func (valpressuretype *ValPressureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valpressuretype)
}

func (valslopetype *ValSlopeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valslopetype)
}

func (valspeedtype *ValSpeedType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valspeedtype)
}

func (valtemperaturetype *ValTemperatureType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valtemperaturetype)
}

func (valweighttype *ValWeightType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valweighttype)
}

func (xhtmltype *XHTMLType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, xhtmltype)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.CodeATCReportingTypes,
		stage.CodeATCReportingType_stagedOrder,
		stage.CodeATCReportingTypes_reference,
		&stage.CodeATCReportingTypes_referenceOrder,
		stage.CodeATCReportingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAerialRefuellingPointTypes,
		stage.CodeAerialRefuellingPointType_stagedOrder,
		stage.CodeAerialRefuellingPointTypes_reference,
		&stage.CodeAerialRefuellingPointTypes_referenceOrder,
		stage.CodeAerialRefuellingPointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAerialRefuellingPrefixTypes,
		stage.CodeAerialRefuellingPrefixType_stagedOrder,
		stage.CodeAerialRefuellingPrefixTypes_reference,
		&stage.CodeAerialRefuellingPrefixTypes_referenceOrder,
		stage.CodeAerialRefuellingPrefixTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAerialRefuellingTypes,
		stage.CodeAerialRefuellingType_stagedOrder,
		stage.CodeAerialRefuellingTypes_reference,
		&stage.CodeAerialRefuellingTypes_referenceOrder,
		stage.CodeAerialRefuellingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAircraftCategoryTypes,
		stage.CodeAircraftCategoryType_stagedOrder,
		stage.CodeAircraftCategoryTypes_reference,
		&stage.CodeAircraftCategoryTypes_referenceOrder,
		stage.CodeAircraftCategoryTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAircraftEngineNumberTypes,
		stage.CodeAircraftEngineNumberType_stagedOrder,
		stage.CodeAircraftEngineNumberTypes_reference,
		&stage.CodeAircraftEngineNumberTypes_referenceOrder,
		stage.CodeAircraftEngineNumberTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAircraftEngineTypes,
		stage.CodeAircraftEngineType_stagedOrder,
		stage.CodeAircraftEngineTypes_reference,
		&stage.CodeAircraftEngineTypes_referenceOrder,
		stage.CodeAircraftEngineTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAircraftGroundServiceTypes,
		stage.CodeAircraftGroundServiceType_stagedOrder,
		stage.CodeAircraftGroundServiceTypes_reference,
		&stage.CodeAircraftGroundServiceTypes_referenceOrder,
		stage.CodeAircraftGroundServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAircraftICAOTypes,
		stage.CodeAircraftICAOType_stagedOrder,
		stage.CodeAircraftICAOTypes_reference,
		&stage.CodeAircraftICAOTypes_referenceOrder,
		stage.CodeAircraftICAOTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAircraftStandTypes,
		stage.CodeAircraftStandType_stagedOrder,
		stage.CodeAircraftStandTypes_reference,
		&stage.CodeAircraftStandTypes_referenceOrder,
		stage.CodeAircraftStandTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAircraftTypes,
		stage.CodeAircraftType_stagedOrder,
		stage.CodeAircraftTypes_reference,
		&stage.CodeAircraftTypes_referenceOrder,
		stage.CodeAircraftTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAircraftWingspanClassTypes,
		stage.CodeAircraftWingspanClassType_stagedOrder,
		stage.CodeAircraftWingspanClassTypes_reference,
		&stage.CodeAircraftWingspanClassTypes_referenceOrder,
		stage.CodeAircraftWingspanClassTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirportHeliportCollocationTypes,
		stage.CodeAirportHeliportCollocationType_stagedOrder,
		stage.CodeAirportHeliportCollocationTypes_reference,
		&stage.CodeAirportHeliportCollocationTypes_referenceOrder,
		stage.CodeAirportHeliportCollocationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirportHeliportDesignatorTypes,
		stage.CodeAirportHeliportDesignatorType_stagedOrder,
		stage.CodeAirportHeliportDesignatorTypes_reference,
		&stage.CodeAirportHeliportDesignatorTypes_referenceOrder,
		stage.CodeAirportHeliportDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirportHeliportTypes,
		stage.CodeAirportHeliportType_stagedOrder,
		stage.CodeAirportHeliportTypes_reference,
		&stage.CodeAirportHeliportTypes_referenceOrder,
		stage.CodeAirportHeliportTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirportWarningTypes,
		stage.CodeAirportWarningType_stagedOrder,
		stage.CodeAirportWarningTypes_reference,
		&stage.CodeAirportWarningTypes_referenceOrder,
		stage.CodeAirportWarningTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirspaceActivityTypes,
		stage.CodeAirspaceActivityType_stagedOrder,
		stage.CodeAirspaceActivityTypes_reference,
		&stage.CodeAirspaceActivityTypes_referenceOrder,
		stage.CodeAirspaceActivityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirspaceAggregationTypes,
		stage.CodeAirspaceAggregationType_stagedOrder,
		stage.CodeAirspaceAggregationTypes_reference,
		&stage.CodeAirspaceAggregationTypes_referenceOrder,
		stage.CodeAirspaceAggregationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirspaceClassificationTypes,
		stage.CodeAirspaceClassificationType_stagedOrder,
		stage.CodeAirspaceClassificationTypes_reference,
		&stage.CodeAirspaceClassificationTypes_referenceOrder,
		stage.CodeAirspaceClassificationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirspaceDependencyTypes,
		stage.CodeAirspaceDependencyType_stagedOrder,
		stage.CodeAirspaceDependencyTypes_reference,
		&stage.CodeAirspaceDependencyTypes_referenceOrder,
		stage.CodeAirspaceDependencyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirspaceDesignatorTypes,
		stage.CodeAirspaceDesignatorType_stagedOrder,
		stage.CodeAirspaceDesignatorTypes_reference,
		&stage.CodeAirspaceDesignatorTypes_referenceOrder,
		stage.CodeAirspaceDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirspacePointPositionTypes,
		stage.CodeAirspacePointPositionType_stagedOrder,
		stage.CodeAirspacePointPositionTypes_reference,
		&stage.CodeAirspacePointPositionTypes_referenceOrder,
		stage.CodeAirspacePointPositionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirspacePointRoleTypes,
		stage.CodeAirspacePointRoleType_stagedOrder,
		stage.CodeAirspacePointRoleTypes_reference,
		&stage.CodeAirspacePointRoleTypes_referenceOrder,
		stage.CodeAirspacePointRoleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAirspaceTypes,
		stage.CodeAirspaceType_stagedOrder,
		stage.CodeAirspaceTypes_reference,
		&stage.CodeAirspaceTypes_referenceOrder,
		stage.CodeAirspaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAltitudeAdjustmentTypes,
		stage.CodeAltitudeAdjustmentType_stagedOrder,
		stage.CodeAltitudeAdjustmentTypes_reference,
		&stage.CodeAltitudeAdjustmentTypes_referenceOrder,
		stage.CodeAltitudeAdjustmentTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAltitudeUseTypes,
		stage.CodeAltitudeUseType_stagedOrder,
		stage.CodeAltitudeUseTypes_reference,
		&stage.CodeAltitudeUseTypes_referenceOrder,
		stage.CodeAltitudeUseTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApproachEquipmentAdditionalTypes,
		stage.CodeApproachEquipmentAdditionalType_stagedOrder,
		stage.CodeApproachEquipmentAdditionalTypes_reference,
		&stage.CodeApproachEquipmentAdditionalTypes_referenceOrder,
		stage.CodeApproachEquipmentAdditionalTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApproachGuidanceTypes,
		stage.CodeApproachGuidanceType_stagedOrder,
		stage.CodeApproachGuidanceTypes_reference,
		&stage.CodeApproachGuidanceTypes_referenceOrder,
		stage.CodeApproachGuidanceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApproachLightingICAOTypes,
		stage.CodeApproachLightingICAOType_stagedOrder,
		stage.CodeApproachLightingICAOTypes_reference,
		&stage.CodeApproachLightingICAOTypes_referenceOrder,
		stage.CodeApproachLightingICAOTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApproachLightingTypes,
		stage.CodeApproachLightingType_stagedOrder,
		stage.CodeApproachLightingTypes_reference,
		&stage.CodeApproachLightingTypes_referenceOrder,
		stage.CodeApproachLightingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApproachPrefixTypes,
		stage.CodeApproachPrefixType_stagedOrder,
		stage.CodeApproachPrefixTypes_reference,
		&stage.CodeApproachPrefixTypes_referenceOrder,
		stage.CodeApproachPrefixTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApproachTypes,
		stage.CodeApproachType_stagedOrder,
		stage.CodeApproachTypes_reference,
		&stage.CodeApproachTypes_referenceOrder,
		stage.CodeApproachTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApprovalTypes,
		stage.CodeApprovalType_stagedOrder,
		stage.CodeApprovalTypes_reference,
		&stage.CodeApprovalTypes_referenceOrder,
		stage.CodeApprovalTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApronElementTypes,
		stage.CodeApronElementType_stagedOrder,
		stage.CodeApronElementTypes_reference,
		&stage.CodeApronElementTypes_referenceOrder,
		stage.CodeApronElementTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeApronSectionTypes,
		stage.CodeApronSectionType_stagedOrder,
		stage.CodeApronSectionTypes_reference,
		&stage.CodeApronSectionTypes_referenceOrder,
		stage.CodeApronSectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeArcDirectionTypes,
		stage.CodeArcDirectionType_stagedOrder,
		stage.CodeArcDirectionTypes_reference,
		&stage.CodeArcDirectionTypes_referenceOrder,
		stage.CodeArcDirectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeArrestingGearEnergyAbsorbTypes,
		stage.CodeArrestingGearEnergyAbsorbType_stagedOrder,
		stage.CodeArrestingGearEnergyAbsorbTypes_reference,
		&stage.CodeArrestingGearEnergyAbsorbTypes_referenceOrder,
		stage.CodeArrestingGearEnergyAbsorbTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeArrestingGearEngageDeviceTypes,
		stage.CodeArrestingGearEngageDeviceType_stagedOrder,
		stage.CodeArrestingGearEngageDeviceTypes_reference,
		&stage.CodeArrestingGearEngageDeviceTypes_referenceOrder,
		stage.CodeArrestingGearEngageDeviceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAuralMorseTypes,
		stage.CodeAuralMorseType_stagedOrder,
		stage.CodeAuralMorseTypes_reference,
		&stage.CodeAuralMorseTypes_referenceOrder,
		stage.CodeAuralMorseTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAuthorityRoleTypes,
		stage.CodeAuthorityRoleType_stagedOrder,
		stage.CodeAuthorityRoleTypes_reference,
		&stage.CodeAuthorityRoleTypes_referenceOrder,
		stage.CodeAuthorityRoleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAuthorityTypes,
		stage.CodeAuthorityType_stagedOrder,
		stage.CodeAuthorityTypes_reference,
		&stage.CodeAuthorityTypes_referenceOrder,
		stage.CodeAuthorityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeAviationStandardsTypes,
		stage.CodeAviationStandardsType_stagedOrder,
		stage.CodeAviationStandardsTypes_reference,
		&stage.CodeAviationStandardsTypes_referenceOrder,
		stage.CodeAviationStandardsTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeBearingTypes,
		stage.CodeBearingType_stagedOrder,
		stage.CodeBearingTypes_reference,
		&stage.CodeBearingTypes_referenceOrder,
		stage.CodeBearingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeBuoyDesignatorTypes,
		stage.CodeBuoyDesignatorType_stagedOrder,
		stage.CodeBuoyDesignatorTypes_reference,
		&stage.CodeBuoyDesignatorTypes_referenceOrder,
		stage.CodeBuoyDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeBuoyTypes,
		stage.CodeBuoyType_stagedOrder,
		stage.CodeBuoyTypes_reference,
		&stage.CodeBuoyTypes_referenceOrder,
		stage.CodeBuoyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeCardinalDirectionTypes,
		stage.CodeCardinalDirectionType_stagedOrder,
		stage.CodeCardinalDirectionTypes_reference,
		&stage.CodeCardinalDirectionTypes_referenceOrder,
		stage.CodeCardinalDirectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeCheckpointCategoryTypes,
		stage.CodeCheckpointCategoryType_stagedOrder,
		stage.CodeCheckpointCategoryTypes_reference,
		&stage.CodeCheckpointCategoryTypes_referenceOrder,
		stage.CodeCheckpointCategoryTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeColourTypes,
		stage.CodeColourType_stagedOrder,
		stage.CodeColourTypes_reference,
		&stage.CodeColourTypes_referenceOrder,
		stage.CodeColourTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeCommunicationChannelTypes,
		stage.CodeCommunicationChannelType_stagedOrder,
		stage.CodeCommunicationChannelTypes_reference,
		&stage.CodeCommunicationChannelTypes_referenceOrder,
		stage.CodeCommunicationChannelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeCommunicationDirectionTypes,
		stage.CodeCommunicationDirectionType_stagedOrder,
		stage.CodeCommunicationDirectionTypes_reference,
		&stage.CodeCommunicationDirectionTypes_referenceOrder,
		stage.CodeCommunicationDirectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeCommunicationModeTypes,
		stage.CodeCommunicationModeType_stagedOrder,
		stage.CodeCommunicationModeTypes_reference,
		&stage.CodeCommunicationModeTypes_referenceOrder,
		stage.CodeCommunicationModeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeComparisonTypes,
		stage.CodeComparisonType_stagedOrder,
		stage.CodeComparisonTypes_reference,
		&stage.CodeComparisonTypes_referenceOrder,
		stage.CodeComparisonTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeContaminationTypes,
		stage.CodeContaminationType_stagedOrder,
		stage.CodeContaminationTypes_reference,
		&stage.CodeContaminationTypes_referenceOrder,
		stage.CodeContaminationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeCourseQualityILSTypes,
		stage.CodeCourseQualityILSType_stagedOrder,
		stage.CodeCourseQualityILSTypes_reference,
		&stage.CodeCourseQualityILSTypes_referenceOrder,
		stage.CodeCourseQualityILSTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeCourseTypes,
		stage.CodeCourseType_stagedOrder,
		stage.CodeCourseTypes_reference,
		&stage.CodeCourseTypes_referenceOrder,
		stage.CodeCourseTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDMEChannelTypes,
		stage.CodeDMEChannelType_stagedOrder,
		stage.CodeDMEChannelTypes_reference,
		&stage.CodeDMEChannelTypes_referenceOrder,
		stage.CodeDMEChannelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDMETypes,
		stage.CodeDMEType_stagedOrder,
		stage.CodeDMETypes_reference,
		&stage.CodeDMETypes_referenceOrder,
		stage.CodeDMETypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDayTypes,
		stage.CodeDayType_stagedOrder,
		stage.CodeDayTypes_reference,
		&stage.CodeDayTypes_referenceOrder,
		stage.CodeDayTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDeclaredDistanceTypes,
		stage.CodeDeclaredDistanceType_stagedOrder,
		stage.CodeDeclaredDistanceTypes_reference,
		&stage.CodeDeclaredDistanceTypes_referenceOrder,
		stage.CodeDeclaredDistanceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDesignStandardTypes,
		stage.CodeDesignStandardType_stagedOrder,
		stage.CodeDesignStandardTypes_reference,
		&stage.CodeDesignStandardTypes_referenceOrder,
		stage.CodeDesignStandardTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDesignatedPointDesignatorTypes,
		stage.CodeDesignatedPointDesignatorType_stagedOrder,
		stage.CodeDesignatedPointDesignatorTypes_reference,
		&stage.CodeDesignatedPointDesignatorTypes_referenceOrder,
		stage.CodeDesignatedPointDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDesignatedPointTypes,
		stage.CodeDesignatedPointType_stagedOrder,
		stage.CodeDesignatedPointTypes_reference,
		&stage.CodeDesignatedPointTypes_referenceOrder,
		stage.CodeDesignatedPointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDirectionReferenceTypes,
		stage.CodeDirectionReferenceType_stagedOrder,
		stage.CodeDirectionReferenceTypes_reference,
		&stage.CodeDirectionReferenceTypes_referenceOrder,
		stage.CodeDirectionReferenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDirectionTurnTypes,
		stage.CodeDirectionTurnType_stagedOrder,
		stage.CodeDirectionTurnTypes_reference,
		&stage.CodeDirectionTurnTypes_referenceOrder,
		stage.CodeDirectionTurnTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDirectionTypes,
		stage.CodeDirectionType_stagedOrder,
		stage.CodeDirectionTypes_reference,
		&stage.CodeDirectionTypes_referenceOrder,
		stage.CodeDirectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeDistanceIndicationTypes,
		stage.CodeDistanceIndicationType_stagedOrder,
		stage.CodeDistanceIndicationTypes_reference,
		&stage.CodeDistanceIndicationTypes_referenceOrder,
		stage.CodeDistanceIndicationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeEmissionBandTypes,
		stage.CodeEmissionBandType_stagedOrder,
		stage.CodeEmissionBandTypes_reference,
		&stage.CodeEmissionBandTypes_referenceOrder,
		stage.CodeEmissionBandTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeEquipmentAntiCollisionTypes,
		stage.CodeEquipmentAntiCollisionType_stagedOrder,
		stage.CodeEquipmentAntiCollisionTypes_reference,
		&stage.CodeEquipmentAntiCollisionTypes_referenceOrder,
		stage.CodeEquipmentAntiCollisionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeEquipmentUnavailableTypes,
		stage.CodeEquipmentUnavailableType_stagedOrder,
		stage.CodeEquipmentUnavailableTypes_reference,
		&stage.CodeEquipmentUnavailableTypes_referenceOrder,
		stage.CodeEquipmentUnavailableTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFacilityRankingTypes,
		stage.CodeFacilityRankingType_stagedOrder,
		stage.CodeFacilityRankingTypes_reference,
		&stage.CodeFacilityRankingTypes_referenceOrder,
		stage.CodeFacilityRankingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFinalGuidanceTypes,
		stage.CodeFinalGuidanceType_stagedOrder,
		stage.CodeFinalGuidanceTypes_reference,
		&stage.CodeFinalGuidanceTypes_referenceOrder,
		stage.CodeFinalGuidanceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFireFightingTypes,
		stage.CodeFireFightingType_stagedOrder,
		stage.CodeFireFightingTypes_reference,
		&stage.CodeFireFightingTypes_referenceOrder,
		stage.CodeFireFightingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlightDestinationTypes,
		stage.CodeFlightDestinationType_stagedOrder,
		stage.CodeFlightDestinationTypes_reference,
		&stage.CodeFlightDestinationTypes_referenceOrder,
		stage.CodeFlightDestinationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlightOriginTypes,
		stage.CodeFlightOriginType_stagedOrder,
		stage.CodeFlightOriginTypes_reference,
		&stage.CodeFlightOriginTypes_referenceOrder,
		stage.CodeFlightOriginTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlightPurposeTypes,
		stage.CodeFlightPurposeType_stagedOrder,
		stage.CodeFlightPurposeTypes_reference,
		&stage.CodeFlightPurposeTypes_referenceOrder,
		stage.CodeFlightPurposeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlightRestrictionDesignatorTypes,
		stage.CodeFlightRestrictionDesignatorType_stagedOrder,
		stage.CodeFlightRestrictionDesignatorTypes_reference,
		&stage.CodeFlightRestrictionDesignatorTypes_referenceOrder,
		stage.CodeFlightRestrictionDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlightRestrictionTypes,
		stage.CodeFlightRestrictionType_stagedOrder,
		stage.CodeFlightRestrictionTypes_reference,
		&stage.CodeFlightRestrictionTypes_referenceOrder,
		stage.CodeFlightRestrictionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlightRuleTypes,
		stage.CodeFlightRuleType_stagedOrder,
		stage.CodeFlightRuleTypes_reference,
		&stage.CodeFlightRuleTypes_referenceOrder,
		stage.CodeFlightRuleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlightStatusTypes,
		stage.CodeFlightStatusType_stagedOrder,
		stage.CodeFlightStatusTypes_reference,
		&stage.CodeFlightStatusTypes_referenceOrder,
		stage.CodeFlightStatusTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlightTypes,
		stage.CodeFlightType_stagedOrder,
		stage.CodeFlightTypes_reference,
		&stage.CodeFlightTypes_referenceOrder,
		stage.CodeFlightTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFlowConditionOperationTypes,
		stage.CodeFlowConditionOperationType_stagedOrder,
		stage.CodeFlowConditionOperationTypes_reference,
		&stage.CodeFlowConditionOperationTypes_referenceOrder,
		stage.CodeFlowConditionOperationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFreeFlightTypes,
		stage.CodeFreeFlightType_stagedOrder,
		stage.CodeFreeFlightTypes_reference,
		&stage.CodeFreeFlightTypes_referenceOrder,
		stage.CodeFreeFlightTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFrictionDeviceTypes,
		stage.CodeFrictionDeviceType_stagedOrder,
		stage.CodeFrictionDeviceTypes_reference,
		&stage.CodeFrictionDeviceTypes_referenceOrder,
		stage.CodeFrictionDeviceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFrictionEstimateTypes,
		stage.CodeFrictionEstimateType_stagedOrder,
		stage.CodeFrictionEstimateTypes_reference,
		&stage.CodeFrictionEstimateTypes_referenceOrder,
		stage.CodeFrictionEstimateTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeFuelTypes,
		stage.CodeFuelType_stagedOrder,
		stage.CodeFuelTypes_reference,
		&stage.CodeFuelTypes_referenceOrder,
		stage.CodeFuelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeGeoBorderTypes,
		stage.CodeGeoBorderType_stagedOrder,
		stage.CodeGeoBorderTypes_reference,
		&stage.CodeGeoBorderTypes_referenceOrder,
		stage.CodeGeoBorderTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeGradeSeparationTypes,
		stage.CodeGradeSeparationType_stagedOrder,
		stage.CodeGradeSeparationTypes_reference,
		&stage.CodeGradeSeparationTypes_referenceOrder,
		stage.CodeGradeSeparationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeGroundLightingTypes,
		stage.CodeGroundLightingType_stagedOrder,
		stage.CodeGroundLightingTypes_reference,
		&stage.CodeGroundLightingTypes_referenceOrder,
		stage.CodeGroundLightingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeGuidanceLineTypes,
		stage.CodeGuidanceLineType_stagedOrder,
		stage.CodeGuidanceLineTypes_reference,
		&stage.CodeGuidanceLineTypes_referenceOrder,
		stage.CodeGuidanceLineTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeHeightReferenceTypes,
		stage.CodeHeightReferenceType_stagedOrder,
		stage.CodeHeightReferenceTypes_reference,
		&stage.CodeHeightReferenceTypes_referenceOrder,
		stage.CodeHeightReferenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeHelicopterPerformanceTypes,
		stage.CodeHelicopterPerformanceType_stagedOrder,
		stage.CodeHelicopterPerformanceTypes_reference,
		&stage.CodeHelicopterPerformanceTypes_referenceOrder,
		stage.CodeHelicopterPerformanceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeHoldingCategoryTypes,
		stage.CodeHoldingCategoryType_stagedOrder,
		stage.CodeHoldingCategoryTypes_reference,
		&stage.CodeHoldingCategoryTypes_referenceOrder,
		stage.CodeHoldingCategoryTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeHoldingUsageTypes,
		stage.CodeHoldingUsageType_stagedOrder,
		stage.CodeHoldingUsageTypes_reference,
		&stage.CodeHoldingUsageTypes_referenceOrder,
		stage.CodeHoldingUsageTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeHoldingUseTypes,
		stage.CodeHoldingUseType_stagedOrder,
		stage.CodeHoldingUseTypes_reference,
		&stage.CodeHoldingUseTypes_referenceOrder,
		stage.CodeHoldingUseTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeIATATypes,
		stage.CodeIATAType_stagedOrder,
		stage.CodeIATATypes_reference,
		&stage.CodeIATATypes_referenceOrder,
		stage.CodeIATATypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeICAOCountryTypes,
		stage.CodeICAOCountryType_stagedOrder,
		stage.CodeICAOCountryTypes_reference,
		&stage.CodeICAOCountryTypes_referenceOrder,
		stage.CodeICAOCountryTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeICAOTypes,
		stage.CodeICAOType_stagedOrder,
		stage.CodeICAOTypes_reference,
		&stage.CodeICAOTypes_referenceOrder,
		stage.CodeICAOTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeILSBackCourseTypes,
		stage.CodeILSBackCourseType_stagedOrder,
		stage.CodeILSBackCourseTypes_reference,
		&stage.CodeILSBackCourseTypes_referenceOrder,
		stage.CodeILSBackCourseTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeIntegrityLevelILSTypes,
		stage.CodeIntegrityLevelILSType_stagedOrder,
		stage.CodeIntegrityLevelILSTypes_reference,
		&stage.CodeIntegrityLevelILSTypes_referenceOrder,
		stage.CodeIntegrityLevelILSTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeIntensityStandByTypes,
		stage.CodeIntensityStandByType_stagedOrder,
		stage.CodeIntensityStandByTypes_reference,
		&stage.CodeIntensityStandByTypes_referenceOrder,
		stage.CodeIntensityStandByTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLanguageTypes,
		stage.CodeLanguageType_stagedOrder,
		stage.CodeLanguageTypes_reference,
		&stage.CodeLanguageTypes_referenceOrder,
		stage.CodeLanguageTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLevelSeriesTypes,
		stage.CodeLevelSeriesType_stagedOrder,
		stage.CodeLevelSeriesTypes_reference,
		&stage.CodeLevelSeriesTypes_referenceOrder,
		stage.CodeLevelSeriesTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLevelTableDesignatorTypes,
		stage.CodeLevelTableDesignatorType_stagedOrder,
		stage.CodeLevelTableDesignatorTypes_reference,
		&stage.CodeLevelTableDesignatorTypes_referenceOrder,
		stage.CodeLevelTableDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLevelTypes,
		stage.CodeLevelType_stagedOrder,
		stage.CodeLevelTypes_reference,
		&stage.CodeLevelTypes_referenceOrder,
		stage.CodeLevelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLightHoldingPositionTypes,
		stage.CodeLightHoldingPositionType_stagedOrder,
		stage.CodeLightHoldingPositionTypes_reference,
		&stage.CodeLightHoldingPositionTypes_referenceOrder,
		stage.CodeLightHoldingPositionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLightIntensityTypes,
		stage.CodeLightIntensityType_stagedOrder,
		stage.CodeLightIntensityTypes_reference,
		&stage.CodeLightIntensityTypes_referenceOrder,
		stage.CodeLightIntensityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLightSourceTypes,
		stage.CodeLightSourceType_stagedOrder,
		stage.CodeLightSourceTypes_reference,
		&stage.CodeLightSourceTypes_referenceOrder,
		stage.CodeLightSourceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLightingJARTypes,
		stage.CodeLightingJARType_stagedOrder,
		stage.CodeLightingJARTypes_reference,
		&stage.CodeLightingJARTypes_referenceOrder,
		stage.CodeLightingJARTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLoadingBridgeTypes,
		stage.CodeLoadingBridgeType_stagedOrder,
		stage.CodeLoadingBridgeTypes_reference,
		&stage.CodeLoadingBridgeTypes_referenceOrder,
		stage.CodeLoadingBridgeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLocationQualifierTypes,
		stage.CodeLocationQualifierType_stagedOrder,
		stage.CodeLocationQualifierTypes_reference,
		&stage.CodeLocationQualifierTypes_referenceOrder,
		stage.CodeLocationQualifierTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeLogicalOperatorTypes,
		stage.CodeLogicalOperatorType_stagedOrder,
		stage.CodeLogicalOperatorTypes_reference,
		&stage.CodeLogicalOperatorTypes_referenceOrder,
		stage.CodeLogicalOperatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMLSAzimuthTypes,
		stage.CodeMLSAzimuthType_stagedOrder,
		stage.CodeMLSAzimuthTypes_reference,
		&stage.CodeMLSAzimuthTypes_referenceOrder,
		stage.CodeMLSAzimuthTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMLSChannelTypes,
		stage.CodeMLSChannelType_stagedOrder,
		stage.CodeMLSChannelTypes_reference,
		&stage.CodeMLSChannelTypes_referenceOrder,
		stage.CodeMLSChannelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMarkerBeaconSignalTypes,
		stage.CodeMarkerBeaconSignalType_stagedOrder,
		stage.CodeMarkerBeaconSignalTypes_reference,
		&stage.CodeMarkerBeaconSignalTypes_referenceOrder,
		stage.CodeMarkerBeaconSignalTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMarkingConditionTypes,
		stage.CodeMarkingConditionType_stagedOrder,
		stage.CodeMarkingConditionTypes_reference,
		&stage.CodeMarkingConditionTypes_referenceOrder,
		stage.CodeMarkingConditionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMarkingStyleTypes,
		stage.CodeMarkingStyleType_stagedOrder,
		stage.CodeMarkingStyleTypes_reference,
		&stage.CodeMarkingStyleTypes_referenceOrder,
		stage.CodeMarkingStyleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMeteoConditionsTypes,
		stage.CodeMeteoConditionsType_stagedOrder,
		stage.CodeMeteoConditionsTypes_reference,
		&stage.CodeMeteoConditionsTypes_referenceOrder,
		stage.CodeMeteoConditionsTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMilitaryOperationsTypes,
		stage.CodeMilitaryOperationsType_stagedOrder,
		stage.CodeMilitaryOperationsTypes_reference,
		&stage.CodeMilitaryOperationsTypes_referenceOrder,
		stage.CodeMilitaryOperationsTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMilitaryRoutePointTypes,
		stage.CodeMilitaryRoutePointType_stagedOrder,
		stage.CodeMilitaryRoutePointTypes_reference,
		&stage.CodeMilitaryRoutePointTypes_referenceOrder,
		stage.CodeMilitaryRoutePointTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMilitaryStatusTypes,
		stage.CodeMilitaryStatusType_stagedOrder,
		stage.CodeMilitaryStatusTypes_reference,
		&stage.CodeMilitaryStatusTypes_referenceOrder,
		stage.CodeMilitaryStatusTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMilitaryTrainingTypes,
		stage.CodeMilitaryTrainingType_stagedOrder,
		stage.CodeMilitaryTrainingTypes_reference,
		&stage.CodeMilitaryTrainingTypes_referenceOrder,
		stage.CodeMilitaryTrainingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMinimaFinalApproachPathTypes,
		stage.CodeMinimaFinalApproachPathType_stagedOrder,
		stage.CodeMinimaFinalApproachPathTypes_reference,
		&stage.CodeMinimaFinalApproachPathTypes_referenceOrder,
		stage.CodeMinimaFinalApproachPathTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMinimumAltitudeTypes,
		stage.CodeMinimumAltitudeType_stagedOrder,
		stage.CodeMinimumAltitudeTypes_reference,
		&stage.CodeMinimumAltitudeTypes_referenceOrder,
		stage.CodeMinimumAltitudeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMinimumHeightTypes,
		stage.CodeMinimumHeightType_stagedOrder,
		stage.CodeMinimumHeightTypes_reference,
		&stage.CodeMinimumHeightTypes_referenceOrder,
		stage.CodeMinimumHeightTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeMissedApproachTypes,
		stage.CodeMissedApproachType_stagedOrder,
		stage.CodeMissedApproachTypes_reference,
		&stage.CodeMissedApproachTypes_referenceOrder,
		stage.CodeMissedApproachTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNDBUsageTypes,
		stage.CodeNDBUsageType_stagedOrder,
		stage.CodeNDBUsageTypes_reference,
		&stage.CodeNDBUsageTypes_referenceOrder,
		stage.CodeNDBUsageTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNavaidDesignatorTypes,
		stage.CodeNavaidDesignatorType_stagedOrder,
		stage.CodeNavaidDesignatorTypes_reference,
		&stage.CodeNavaidDesignatorTypes_referenceOrder,
		stage.CodeNavaidDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNavaidPurposeTypes,
		stage.CodeNavaidPurposeType_stagedOrder,
		stage.CodeNavaidPurposeTypes_reference,
		&stage.CodeNavaidPurposeTypes_referenceOrder,
		stage.CodeNavaidPurposeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNavaidServiceTypes,
		stage.CodeNavaidServiceType_stagedOrder,
		stage.CodeNavaidServiceTypes_reference,
		&stage.CodeNavaidServiceTypes_referenceOrder,
		stage.CodeNavaidServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNavigationAreaRestrictionTypes,
		stage.CodeNavigationAreaRestrictionType_stagedOrder,
		stage.CodeNavigationAreaRestrictionTypes_reference,
		&stage.CodeNavigationAreaRestrictionTypes_referenceOrder,
		stage.CodeNavigationAreaRestrictionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNavigationAreaTypes,
		stage.CodeNavigationAreaType_stagedOrder,
		stage.CodeNavigationAreaTypes_reference,
		&stage.CodeNavigationAreaTypes_referenceOrder,
		stage.CodeNavigationAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNavigationEquipmentTypes,
		stage.CodeNavigationEquipmentType_stagedOrder,
		stage.CodeNavigationEquipmentTypes_reference,
		&stage.CodeNavigationEquipmentTypes_referenceOrder,
		stage.CodeNavigationEquipmentTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNavigationSpecificationTypes,
		stage.CodeNavigationSpecificationType_stagedOrder,
		stage.CodeNavigationSpecificationTypes_reference,
		&stage.CodeNavigationSpecificationTypes_referenceOrder,
		stage.CodeNavigationSpecificationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNitrogenTypes,
		stage.CodeNitrogenType_stagedOrder,
		stage.CodeNitrogenTypes_reference,
		&stage.CodeNitrogenTypes_referenceOrder,
		stage.CodeNitrogenTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNorthReferenceTypes,
		stage.CodeNorthReferenceType_stagedOrder,
		stage.CodeNorthReferenceTypes_reference,
		&stage.CodeNorthReferenceTypes_referenceOrder,
		stage.CodeNorthReferenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeNotePurposeTypes,
		stage.CodeNotePurposeType_stagedOrder,
		stage.CodeNotePurposeTypes_reference,
		&stage.CodeNotePurposeTypes_referenceOrder,
		stage.CodeNotePurposeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeObstacleAreaTypes,
		stage.CodeObstacleAreaType_stagedOrder,
		stage.CodeObstacleAreaTypes_reference,
		&stage.CodeObstacleAreaTypes_referenceOrder,
		stage.CodeObstacleAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeObstacleAssessmentSurfaceTypes,
		stage.CodeObstacleAssessmentSurfaceType_stagedOrder,
		stage.CodeObstacleAssessmentSurfaceTypes_reference,
		&stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder,
		stage.CodeObstacleAssessmentSurfaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeObstructionIdSurfaceZoneTypes,
		stage.CodeObstructionIdSurfaceZoneType_stagedOrder,
		stage.CodeObstructionIdSurfaceZoneTypes_reference,
		&stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder,
		stage.CodeObstructionIdSurfaceZoneTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeOilTypes,
		stage.CodeOilType_stagedOrder,
		stage.CodeOilTypes_reference,
		&stage.CodeOilTypes_referenceOrder,
		stage.CodeOilTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeOperationAirportHeliportTypes,
		stage.CodeOperationAirportHeliportType_stagedOrder,
		stage.CodeOperationAirportHeliportTypes_reference,
		&stage.CodeOperationAirportHeliportTypes_referenceOrder,
		stage.CodeOperationAirportHeliportTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeOperationManoeuvringAreaTypes,
		stage.CodeOperationManoeuvringAreaType_stagedOrder,
		stage.CodeOperationManoeuvringAreaTypes_reference,
		&stage.CodeOperationManoeuvringAreaTypes_referenceOrder,
		stage.CodeOperationManoeuvringAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeOrganisationDesignatorTypes,
		stage.CodeOrganisationDesignatorType_stagedOrder,
		stage.CodeOrganisationDesignatorTypes_reference,
		&stage.CodeOrganisationDesignatorTypes_referenceOrder,
		stage.CodeOrganisationDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeOrganisationHierarchyTypes,
		stage.CodeOrganisationHierarchyType_stagedOrder,
		stage.CodeOrganisationHierarchyTypes_reference,
		&stage.CodeOrganisationHierarchyTypes_referenceOrder,
		stage.CodeOrganisationHierarchyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeOrganisationTypes,
		stage.CodeOrganisationType_stagedOrder,
		stage.CodeOrganisationTypes_reference,
		&stage.CodeOrganisationTypes_referenceOrder,
		stage.CodeOrganisationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeOxygenTypes,
		stage.CodeOxygenType_stagedOrder,
		stage.CodeOxygenTypes_reference,
		&stage.CodeOxygenTypes_referenceOrder,
		stage.CodeOxygenTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePARTypes,
		stage.CodePARType_stagedOrder,
		stage.CodePARTypes_reference,
		&stage.CodePARTypes_referenceOrder,
		stage.CodePARTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePCNMethodTypes,
		stage.CodePCNMethodType_stagedOrder,
		stage.CodePCNMethodTypes_reference,
		&stage.CodePCNMethodTypes_referenceOrder,
		stage.CodePCNMethodTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePCNPavementTypes,
		stage.CodePCNPavementType_stagedOrder,
		stage.CodePCNPavementTypes_reference,
		&stage.CodePCNPavementTypes_referenceOrder,
		stage.CodePCNPavementTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePCNSubgradeTypes,
		stage.CodePCNSubgradeType_stagedOrder,
		stage.CodePCNSubgradeTypes_reference,
		&stage.CodePCNSubgradeTypes_referenceOrder,
		stage.CodePCNSubgradeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePCNTyrePressureTypes,
		stage.CodePCNTyrePressureType_stagedOrder,
		stage.CodePCNTyrePressureTypes_reference,
		&stage.CodePCNTyrePressureTypes_referenceOrder,
		stage.CodePCNTyrePressureTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePassengerServiceTypes,
		stage.CodePassengerServiceType_stagedOrder,
		stage.CodePassengerServiceTypes_reference,
		&stage.CodePassengerServiceTypes_referenceOrder,
		stage.CodePassengerServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePilotControlledLightingTypes,
		stage.CodePilotControlledLightingType_stagedOrder,
		stage.CodePilotControlledLightingTypes_reference,
		&stage.CodePilotControlledLightingTypes_referenceOrder,
		stage.CodePilotControlledLightingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePositionInILSTypes,
		stage.CodePositionInILSType_stagedOrder,
		stage.CodePositionInILSTypes_reference,
		&stage.CodePositionInILSTypes_referenceOrder,
		stage.CodePositionInILSTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodePrimaryRadarTypes,
		stage.CodePrimaryRadarType_stagedOrder,
		stage.CodePrimaryRadarTypes_reference,
		&stage.CodePrimaryRadarTypes_referenceOrder,
		stage.CodePrimaryRadarTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeProcedureAvailabilityTypes,
		stage.CodeProcedureAvailabilityType_stagedOrder,
		stage.CodeProcedureAvailabilityTypes_reference,
		&stage.CodeProcedureAvailabilityTypes_referenceOrder,
		stage.CodeProcedureAvailabilityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeProcedureCodingStandardTypes,
		stage.CodeProcedureCodingStandardType_stagedOrder,
		stage.CodeProcedureCodingStandardTypes_reference,
		&stage.CodeProcedureCodingStandardTypes_referenceOrder,
		stage.CodeProcedureCodingStandardTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeProcedureDistanceTypes,
		stage.CodeProcedureDistanceType_stagedOrder,
		stage.CodeProcedureDistanceTypes_reference,
		&stage.CodeProcedureDistanceTypes_referenceOrder,
		stage.CodeProcedureDistanceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeProcedureFixRoleTypes,
		stage.CodeProcedureFixRoleType_stagedOrder,
		stage.CodeProcedureFixRoleTypes_reference,
		&stage.CodeProcedureFixRoleTypes_referenceOrder,
		stage.CodeProcedureFixRoleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeProcedurePhaseTypes,
		stage.CodeProcedurePhaseType_stagedOrder,
		stage.CodeProcedurePhaseTypes_reference,
		&stage.CodeProcedurePhaseTypes_referenceOrder,
		stage.CodeProcedurePhaseTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeProtectAreaSectionTypes,
		stage.CodeProtectAreaSectionType_stagedOrder,
		stage.CodeProtectAreaSectionTypes_reference,
		&stage.CodeProtectAreaSectionTypes_referenceOrder,
		stage.CodeProtectAreaSectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRNPTypes,
		stage.CodeRNPType_stagedOrder,
		stage.CodeRNPTypes_reference,
		&stage.CodeRNPTypes_referenceOrder,
		stage.CodeRNPTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRVRReadingTypes,
		stage.CodeRVRReadingType_stagedOrder,
		stage.CodeRVRReadingTypes_reference,
		&stage.CodeRVRReadingTypes_referenceOrder,
		stage.CodeRVRReadingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRVSMPointRoleTypes,
		stage.CodeRVSMPointRoleType_stagedOrder,
		stage.CodeRVSMPointRoleTypes_reference,
		&stage.CodeRVSMPointRoleTypes_referenceOrder,
		stage.CodeRVSMPointRoleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRVSMTypes,
		stage.CodeRVSMType_stagedOrder,
		stage.CodeRVSMTypes_reference,
		&stage.CodeRVSMTypes_referenceOrder,
		stage.CodeRVSMTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRadarServiceTypes,
		stage.CodeRadarServiceType_stagedOrder,
		stage.CodeRadarServiceTypes_reference,
		&stage.CodeRadarServiceTypes_referenceOrder,
		stage.CodeRadarServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRadioEmissionTypes,
		stage.CodeRadioEmissionType_stagedOrder,
		stage.CodeRadioEmissionTypes_reference,
		&stage.CodeRadioEmissionTypes_referenceOrder,
		stage.CodeRadioEmissionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRadioFrequencyAreaTypes,
		stage.CodeRadioFrequencyAreaType_stagedOrder,
		stage.CodeRadioFrequencyAreaTypes_reference,
		&stage.CodeRadioFrequencyAreaTypes_referenceOrder,
		stage.CodeRadioFrequencyAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRadioSignalTypes,
		stage.CodeRadioSignalType_stagedOrder,
		stage.CodeRadioSignalTypes_reference,
		&stage.CodeRadioSignalTypes_referenceOrder,
		stage.CodeRadioSignalTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeReferencePathIdentifierTypes,
		stage.CodeReferencePathIdentifierType_stagedOrder,
		stage.CodeReferencePathIdentifierTypes_reference,
		&stage.CodeReferencePathIdentifierTypes_referenceOrder,
		stage.CodeReferencePathIdentifierTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeReferenceRoleTypes,
		stage.CodeReferenceRoleType_stagedOrder,
		stage.CodeReferenceRoleTypes_reference,
		&stage.CodeReferenceRoleTypes_referenceOrder,
		stage.CodeReferenceRoleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeReflectorTypes,
		stage.CodeReflectorType_stagedOrder,
		stage.CodeReflectorTypes_reference,
		&stage.CodeReflectorTypes_referenceOrder,
		stage.CodeReflectorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRelativePositionTypes,
		stage.CodeRelativePositionType_stagedOrder,
		stage.CodeRelativePositionTypes_reference,
		&stage.CodeRelativePositionTypes_referenceOrder,
		stage.CodeRelativePositionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRoadTypes,
		stage.CodeRoadType_stagedOrder,
		stage.CodeRoadTypes_reference,
		&stage.CodeRoadTypes_referenceOrder,
		stage.CodeRoadTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteAvailabilityTypes,
		stage.CodeRouteAvailabilityType_stagedOrder,
		stage.CodeRouteAvailabilityTypes_reference,
		&stage.CodeRouteAvailabilityTypes_referenceOrder,
		stage.CodeRouteAvailabilityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteDesignatorLetterTypes,
		stage.CodeRouteDesignatorLetterType_stagedOrder,
		stage.CodeRouteDesignatorLetterTypes_reference,
		&stage.CodeRouteDesignatorLetterTypes_referenceOrder,
		stage.CodeRouteDesignatorLetterTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteDesignatorPrefixTypes,
		stage.CodeRouteDesignatorPrefixType_stagedOrder,
		stage.CodeRouteDesignatorPrefixTypes_reference,
		&stage.CodeRouteDesignatorPrefixTypes_referenceOrder,
		stage.CodeRouteDesignatorPrefixTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteDesignatorSuffixTypes,
		stage.CodeRouteDesignatorSuffixType_stagedOrder,
		stage.CodeRouteDesignatorSuffixTypes_reference,
		&stage.CodeRouteDesignatorSuffixTypes_referenceOrder,
		stage.CodeRouteDesignatorSuffixTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteIndicatorTypes,
		stage.CodeRouteIndicatorType_stagedOrder,
		stage.CodeRouteIndicatorTypes_reference,
		&stage.CodeRouteIndicatorTypes_referenceOrder,
		stage.CodeRouteIndicatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteNavigationTypes,
		stage.CodeRouteNavigationType_stagedOrder,
		stage.CodeRouteNavigationTypes_reference,
		&stage.CodeRouteNavigationTypes_referenceOrder,
		stage.CodeRouteNavigationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteOriginTypes,
		stage.CodeRouteOriginType_stagedOrder,
		stage.CodeRouteOriginTypes_reference,
		&stage.CodeRouteOriginTypes_referenceOrder,
		stage.CodeRouteOriginTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteSegmentPathTypes,
		stage.CodeRouteSegmentPathType_stagedOrder,
		stage.CodeRouteSegmentPathTypes_reference,
		&stage.CodeRouteSegmentPathTypes_referenceOrder,
		stage.CodeRouteSegmentPathTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRouteTypes,
		stage.CodeRouteType_stagedOrder,
		stage.CodeRouteTypes_reference,
		&stage.CodeRouteTypes_referenceOrder,
		stage.CodeRouteTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRuleProcedureTitleTypes,
		stage.CodeRuleProcedureTitleType_stagedOrder,
		stage.CodeRuleProcedureTitleTypes_reference,
		&stage.CodeRuleProcedureTitleTypes_referenceOrder,
		stage.CodeRuleProcedureTitleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRuleProcedureTypes,
		stage.CodeRuleProcedureType_stagedOrder,
		stage.CodeRuleProcedureTypes_reference,
		&stage.CodeRuleProcedureTypes_referenceOrder,
		stage.CodeRuleProcedureTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRunwayElementTypes,
		stage.CodeRunwayElementType_stagedOrder,
		stage.CodeRunwayElementTypes_reference,
		&stage.CodeRunwayElementTypes_referenceOrder,
		stage.CodeRunwayElementTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRunwayMarkingTypes,
		stage.CodeRunwayMarkingType_stagedOrder,
		stage.CodeRunwayMarkingTypes_reference,
		&stage.CodeRunwayMarkingTypes_referenceOrder,
		stage.CodeRunwayMarkingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRunwayPointRoleTypes,
		stage.CodeRunwayPointRoleType_stagedOrder,
		stage.CodeRunwayPointRoleTypes_reference,
		&stage.CodeRunwayPointRoleTypes_referenceOrder,
		stage.CodeRunwayPointRoleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRunwayProtectionAreaTypes,
		stage.CodeRunwayProtectionAreaType_stagedOrder,
		stage.CodeRunwayProtectionAreaTypes_reference,
		&stage.CodeRunwayProtectionAreaTypes_referenceOrder,
		stage.CodeRunwayProtectionAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRunwaySectionTypes,
		stage.CodeRunwaySectionType_stagedOrder,
		stage.CodeRunwaySectionTypes_reference,
		&stage.CodeRunwaySectionTypes_referenceOrder,
		stage.CodeRunwaySectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeRunwayTypes,
		stage.CodeRunwayType_stagedOrder,
		stage.CodeRunwayTypes_reference,
		&stage.CodeRunwayTypes_referenceOrder,
		stage.CodeRunwayTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSafeAltitudeTypes,
		stage.CodeSafeAltitudeType_stagedOrder,
		stage.CodeSafeAltitudeTypes_reference,
		&stage.CodeSafeAltitudeTypes_referenceOrder,
		stage.CodeSafeAltitudeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSegmentPathTypes,
		stage.CodeSegmentPathType_stagedOrder,
		stage.CodeSegmentPathTypes_reference,
		&stage.CodeSegmentPathTypes_referenceOrder,
		stage.CodeSegmentPathTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSegmentTerminationTypes,
		stage.CodeSegmentTerminationType_stagedOrder,
		stage.CodeSegmentTerminationTypes_reference,
		&stage.CodeSegmentTerminationTypes_referenceOrder,
		stage.CodeSegmentTerminationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeServiceATCTypes,
		stage.CodeServiceATCType_stagedOrder,
		stage.CodeServiceATCTypes_reference,
		&stage.CodeServiceATCTypes_referenceOrder,
		stage.CodeServiceATCTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeServiceATFMTypes,
		stage.CodeServiceATFMType_stagedOrder,
		stage.CodeServiceATFMTypes_reference,
		&stage.CodeServiceATFMTypes_referenceOrder,
		stage.CodeServiceATFMTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeServiceGroundControlTypes,
		stage.CodeServiceGroundControlType_stagedOrder,
		stage.CodeServiceGroundControlTypes_reference,
		&stage.CodeServiceGroundControlTypes_referenceOrder,
		stage.CodeServiceGroundControlTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeServiceInformationTypes,
		stage.CodeServiceInformationType_stagedOrder,
		stage.CodeServiceInformationTypes_reference,
		&stage.CodeServiceInformationTypes_referenceOrder,
		stage.CodeServiceInformationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeServiceSARTypes,
		stage.CodeServiceSARType_stagedOrder,
		stage.CodeServiceSARTypes_reference,
		&stage.CodeServiceSARTypes_referenceOrder,
		stage.CodeServiceSARTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSideTypes,
		stage.CodeSideType_stagedOrder,
		stage.CodeSideTypes_reference,
		&stage.CodeSideTypes_referenceOrder,
		stage.CodeSideTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSignalPerformanceILSTypes,
		stage.CodeSignalPerformanceILSType_stagedOrder,
		stage.CodeSignalPerformanceILSTypes_reference,
		&stage.CodeSignalPerformanceILSTypes_referenceOrder,
		stage.CodeSignalPerformanceILSTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSpecialDateTypes,
		stage.CodeSpecialDateType_stagedOrder,
		stage.CodeSpecialDateTypes_reference,
		&stage.CodeSpecialDateTypes_referenceOrder,
		stage.CodeSpecialDateTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSpecialNavigationChainDesignatorTypes,
		stage.CodeSpecialNavigationChainDesignatorType_stagedOrder,
		stage.CodeSpecialNavigationChainDesignatorTypes_reference,
		&stage.CodeSpecialNavigationChainDesignatorTypes_referenceOrder,
		stage.CodeSpecialNavigationChainDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSpecialNavigationStationTypes,
		stage.CodeSpecialNavigationStationType_stagedOrder,
		stage.CodeSpecialNavigationStationTypes_reference,
		&stage.CodeSpecialNavigationStationTypes_referenceOrder,
		stage.CodeSpecialNavigationStationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSpecialNavigationSystemTypes,
		stage.CodeSpecialNavigationSystemType_stagedOrder,
		stage.CodeSpecialNavigationSystemTypes_reference,
		&stage.CodeSpecialNavigationSystemTypes_referenceOrder,
		stage.CodeSpecialNavigationSystemTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSpeedReferenceTypes,
		stage.CodeSpeedReferenceType_stagedOrder,
		stage.CodeSpeedReferenceTypes_reference,
		&stage.CodeSpeedReferenceTypes_referenceOrder,
		stage.CodeSpeedReferenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeStandbyPowerTypes,
		stage.CodeStandbyPowerType_stagedOrder,
		stage.CodeStandbyPowerTypes_reference,
		&stage.CodeStandbyPowerTypes_referenceOrder,
		stage.CodeStandbyPowerTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeStatusAirportTypes,
		stage.CodeStatusAirportType_stagedOrder,
		stage.CodeStatusAirportTypes_reference,
		&stage.CodeStatusAirportTypes_referenceOrder,
		stage.CodeStatusAirportTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeStatusAirspaceTypes,
		stage.CodeStatusAirspaceType_stagedOrder,
		stage.CodeStatusAirspaceTypes_reference,
		&stage.CodeStatusAirspaceTypes_referenceOrder,
		stage.CodeStatusAirspaceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeStatusConstructionTypes,
		stage.CodeStatusConstructionType_stagedOrder,
		stage.CodeStatusConstructionTypes_reference,
		&stage.CodeStatusConstructionTypes_referenceOrder,
		stage.CodeStatusConstructionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeStatusNavaidTypes,
		stage.CodeStatusNavaidType_stagedOrder,
		stage.CodeStatusNavaidTypes_reference,
		&stage.CodeStatusNavaidTypes_referenceOrder,
		stage.CodeStatusNavaidTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeStatusOperationsTypes,
		stage.CodeStatusOperationsType_stagedOrder,
		stage.CodeStatusOperationsTypes_reference,
		&stage.CodeStatusOperationsTypes_referenceOrder,
		stage.CodeStatusOperationsTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeStatusServiceTypes,
		stage.CodeStatusServiceType_stagedOrder,
		stage.CodeStatusServiceTypes_reference,
		&stage.CodeStatusServiceTypes_referenceOrder,
		stage.CodeStatusServiceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSurfaceCompositionTypes,
		stage.CodeSurfaceCompositionType_stagedOrder,
		stage.CodeSurfaceCompositionTypes_reference,
		&stage.CodeSurfaceCompositionTypes_referenceOrder,
		stage.CodeSurfaceCompositionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSurfaceConditionTypes,
		stage.CodeSurfaceConditionType_stagedOrder,
		stage.CodeSurfaceConditionTypes_reference,
		&stage.CodeSurfaceConditionTypes_referenceOrder,
		stage.CodeSurfaceConditionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSurfacePreparationTypes,
		stage.CodeSurfacePreparationType_stagedOrder,
		stage.CodeSurfacePreparationTypes_reference,
		&stage.CodeSurfacePreparationTypes_referenceOrder,
		stage.CodeSurfacePreparationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeSystemActivationTypes,
		stage.CodeSystemActivationType_stagedOrder,
		stage.CodeSystemActivationTypes_reference,
		&stage.CodeSystemActivationTypes_referenceOrder,
		stage.CodeSystemActivationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTAATypes,
		stage.CodeTAAType_stagedOrder,
		stage.CodeTAATypes_reference,
		&stage.CodeTAATypes_referenceOrder,
		stage.CodeTAATypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTACANChannelTypes,
		stage.CodeTACANChannelType_stagedOrder,
		stage.CodeTACANChannelTypes_reference,
		&stage.CodeTACANChannelTypes_referenceOrder,
		stage.CodeTACANChannelTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTLOFSectionTypes,
		stage.CodeTLOFSectionType_stagedOrder,
		stage.CodeTLOFSectionTypes_reference,
		&stage.CodeTLOFSectionTypes_referenceOrder,
		stage.CodeTLOFSectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTaxiwayElementTypes,
		stage.CodeTaxiwayElementType_stagedOrder,
		stage.CodeTaxiwayElementTypes_reference,
		&stage.CodeTaxiwayElementTypes_referenceOrder,
		stage.CodeTaxiwayElementTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTaxiwaySectionTypes,
		stage.CodeTaxiwaySectionType_stagedOrder,
		stage.CodeTaxiwaySectionTypes_reference,
		&stage.CodeTaxiwaySectionTypes_referenceOrder,
		stage.CodeTaxiwaySectionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTaxiwayTypes,
		stage.CodeTaxiwayType_stagedOrder,
		stage.CodeTaxiwayTypes_reference,
		&stage.CodeTaxiwayTypes_referenceOrder,
		stage.CodeTaxiwayTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTelecomNetworkTypes,
		stage.CodeTelecomNetworkType_stagedOrder,
		stage.CodeTelecomNetworkTypes_reference,
		&stage.CodeTelecomNetworkTypes_referenceOrder,
		stage.CodeTelecomNetworkTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTimeEventCombinationTypes,
		stage.CodeTimeEventCombinationType_stagedOrder,
		stage.CodeTimeEventCombinationTypes_reference,
		&stage.CodeTimeEventCombinationTypes_referenceOrder,
		stage.CodeTimeEventCombinationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTimeEventTypes,
		stage.CodeTimeEventType_stagedOrder,
		stage.CodeTimeEventTypes_reference,
		&stage.CodeTimeEventTypes_referenceOrder,
		stage.CodeTimeEventTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTimeReferenceTypes,
		stage.CodeTimeReferenceType_stagedOrder,
		stage.CodeTimeReferenceTypes_reference,
		&stage.CodeTimeReferenceTypes_referenceOrder,
		stage.CodeTimeReferenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTrajectoryTypes,
		stage.CodeTrajectoryType_stagedOrder,
		stage.CodeTrajectoryTypes_reference,
		&stage.CodeTrajectoryTypes_referenceOrder,
		stage.CodeTrajectoryTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeTransponderTypes,
		stage.CodeTransponderType_stagedOrder,
		stage.CodeTransponderTypes_reference,
		&stage.CodeTransponderTypes_referenceOrder,
		stage.CodeTransponderTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeUUIDTypes,
		stage.CodeUUIDType_stagedOrder,
		stage.CodeUUIDTypes_reference,
		&stage.CodeUUIDTypes_referenceOrder,
		stage.CodeUUIDTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeUnitDependencyTypes,
		stage.CodeUnitDependencyType_stagedOrder,
		stage.CodeUnitDependencyTypes_reference,
		&stage.CodeUnitDependencyTypes_referenceOrder,
		stage.CodeUnitDependencyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeUnitTypes,
		stage.CodeUnitType_stagedOrder,
		stage.CodeUnitTypes_reference,
		&stage.CodeUnitTypes_referenceOrder,
		stage.CodeUnitTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeUpperAlphaTypes,
		stage.CodeUpperAlphaType_stagedOrder,
		stage.CodeUpperAlphaTypes_reference,
		&stage.CodeUpperAlphaTypes_referenceOrder,
		stage.CodeUpperAlphaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeUsageLimitationTypes,
		stage.CodeUsageLimitationType_stagedOrder,
		stage.CodeUsageLimitationTypes_reference,
		&stage.CodeUsageLimitationTypes_referenceOrder,
		stage.CodeUsageLimitationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeVASISTypes,
		stage.CodeVASISType_stagedOrder,
		stage.CodeVASISTypes_reference,
		&stage.CodeVASISTypes_referenceOrder,
		stage.CodeVASISTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeVORTypes,
		stage.CodeVORType_stagedOrder,
		stage.CodeVORTypes_reference,
		&stage.CodeVORTypes_referenceOrder,
		stage.CodeVORTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeValueInterpretationTypes,
		stage.CodeValueInterpretationType_stagedOrder,
		stage.CodeValueInterpretationTypes_reference,
		&stage.CodeValueInterpretationTypes_referenceOrder,
		stage.CodeValueInterpretationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeVerticalDatumTypes,
		stage.CodeVerticalDatumType_stagedOrder,
		stage.CodeVerticalDatumTypes_reference,
		&stage.CodeVerticalDatumTypes_referenceOrder,
		stage.CodeVerticalDatumTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeVerticalReferenceTypes,
		stage.CodeVerticalReferenceType_stagedOrder,
		stage.CodeVerticalReferenceTypes_reference,
		&stage.CodeVerticalReferenceTypes_referenceOrder,
		stage.CodeVerticalReferenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeVerticalStructureMarkingTypes,
		stage.CodeVerticalStructureMarkingType_stagedOrder,
		stage.CodeVerticalStructureMarkingTypes_reference,
		&stage.CodeVerticalStructureMarkingTypes_referenceOrder,
		stage.CodeVerticalStructureMarkingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeVerticalStructureMaterialTypes,
		stage.CodeVerticalStructureMaterialType_stagedOrder,
		stage.CodeVerticalStructureMaterialTypes_reference,
		&stage.CodeVerticalStructureMaterialTypes_referenceOrder,
		stage.CodeVerticalStructureMaterialTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeVerticalStructureTypes,
		stage.CodeVerticalStructureType_stagedOrder,
		stage.CodeVerticalStructureTypes_reference,
		&stage.CodeVerticalStructureTypes_referenceOrder,
		stage.CodeVerticalStructureTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeVisualDockingGuidanceTypes,
		stage.CodeVisualDockingGuidanceType_stagedOrder,
		stage.CodeVisualDockingGuidanceTypes_reference,
		&stage.CodeVisualDockingGuidanceTypes_referenceOrder,
		stage.CodeVisualDockingGuidanceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeWakeTurbulenceTypes,
		stage.CodeWakeTurbulenceType_stagedOrder,
		stage.CodeWakeTurbulenceTypes_reference,
		&stage.CodeWakeTurbulenceTypes_referenceOrder,
		stage.CodeWakeTurbulenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeWorkAreaTypes,
		stage.CodeWorkAreaType_stagedOrder,
		stage.CodeWorkAreaTypes_reference,
		&stage.CodeWorkAreaTypes_referenceOrder,
		stage.CodeWorkAreaTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.CodeYesNoTypes,
		stage.CodeYesNoType_stagedOrder,
		stage.CodeYesNoTypes_reference,
		&stage.CodeYesNoTypes_referenceOrder,
		stage.CodeYesNoTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DateMonthDayTypes,
		stage.DateMonthDayType_stagedOrder,
		stage.DateMonthDayTypes_reference,
		&stage.DateMonthDayTypes_referenceOrder,
		stage.DateMonthDayTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DateTimeTypes,
		stage.DateTimeType_stagedOrder,
		stage.DateTimeTypes_reference,
		&stage.DateTimeTypes_referenceOrder,
		stage.DateTimeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DateTypes,
		stage.DateType_stagedOrder,
		stage.DateTypes_reference,
		&stage.DateTypes_referenceOrder,
		stage.DateTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DateYearTypes,
		stage.DateYearType_stagedOrder,
		stage.DateYearTypes_reference,
		&stage.DateYearTypes_referenceOrder,
		stage.DateYearTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoNumberTypes,
		stage.NoNumberType_stagedOrder,
		stage.NoNumberTypes_reference,
		&stage.NoNumberTypes_referenceOrder,
		stage.NoNumberTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoSequenceTypes,
		stage.NoSequenceType_stagedOrder,
		stage.NoSequenceTypes_reference,
		&stage.NoSequenceTypes_referenceOrder,
		stage.NoSequenceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextAddressTypes,
		stage.TextAddressType_stagedOrder,
		stage.TextAddressTypes_reference,
		&stage.TextAddressTypes_referenceOrder,
		stage.TextAddressTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextDesignatorTypes,
		stage.TextDesignatorType_stagedOrder,
		stage.TextDesignatorTypes_reference,
		&stage.TextDesignatorTypes_referenceOrder,
		stage.TextDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextInstructionTypes,
		stage.TextInstructionType_stagedOrder,
		stage.TextInstructionTypes_reference,
		&stage.TextInstructionTypes_referenceOrder,
		stage.TextInstructionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextNameTypes,
		stage.TextNameType_stagedOrder,
		stage.TextNameTypes_reference,
		&stage.TextNameTypes_referenceOrder,
		stage.TextNameTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextNoteTypes,
		stage.TextNoteType_stagedOrder,
		stage.TextNoteTypes_reference,
		&stage.TextNoteTypes_referenceOrder,
		stage.TextNoteTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextPhoneTypes,
		stage.TextPhoneType_stagedOrder,
		stage.TextPhoneTypes_reference,
		&stage.TextPhoneTypes_referenceOrder,
		stage.TextPhoneTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextPropertyNameTypes,
		stage.TextPropertyNameType_stagedOrder,
		stage.TextPropertyNameTypes_reference,
		&stage.TextPropertyNameTypes_referenceOrder,
		stage.TextPropertyNameTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextRemarkTypes,
		stage.TextRemarkType_stagedOrder,
		stage.TextRemarkTypes_reference,
		&stage.TextRemarkTypes_referenceOrder,
		stage.TextRemarkTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TextSIDSTARDesignatorTypes,
		stage.TextSIDSTARDesignatorType_stagedOrder,
		stage.TextSIDSTARDesignatorTypes_reference,
		&stage.TextSIDSTARDesignatorTypes_referenceOrder,
		stage.TextSIDSTARDesignatorTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TimeTypes,
		stage.TimeType_stagedOrder,
		stage.TimeTypes_reference,
		&stage.TimeTypes_referenceOrder,
		stage.TimeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValAlarmLimitTypes,
		stage.ValAlarmLimitType_stagedOrder,
		stage.ValAlarmLimitTypes_reference,
		&stage.ValAlarmLimitTypes_referenceOrder,
		stage.ValAlarmLimitTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValAngleTypes,
		stage.ValAngleType_stagedOrder,
		stage.ValAngleTypes_reference,
		&stage.ValAngleTypes_referenceOrder,
		stage.ValAngleTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValBearingTypes,
		stage.ValBearingType_stagedOrder,
		stage.ValBearingTypes_reference,
		&stage.ValBearingTypes_referenceOrder,
		stage.ValBearingTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValChannelNumberTypes,
		stage.ValChannelNumberType_stagedOrder,
		stage.ValChannelNumberTypes_reference,
		&stage.ValChannelNumberTypes_referenceOrder,
		stage.ValChannelNumberTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValDepthTypes,
		stage.ValDepthType_stagedOrder,
		stage.ValDepthTypes_reference,
		&stage.ValDepthTypes_referenceOrder,
		stage.ValDepthTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValDistanceSignedTypes,
		stage.ValDistanceSignedType_stagedOrder,
		stage.ValDistanceSignedTypes_reference,
		&stage.ValDistanceSignedTypes_referenceOrder,
		stage.ValDistanceSignedTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValDistanceTypes,
		stage.ValDistanceType_stagedOrder,
		stage.ValDistanceTypes_reference,
		&stage.ValDistanceTypes_referenceOrder,
		stage.ValDistanceTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValDistanceVerticalTypes,
		stage.ValDistanceVerticalType_stagedOrder,
		stage.ValDistanceVerticalTypes_reference,
		&stage.ValDistanceVerticalTypes_referenceOrder,
		stage.ValDistanceVerticalTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValDurationTypes,
		stage.ValDurationType_stagedOrder,
		stage.ValDurationTypes_reference,
		&stage.ValDurationTypes_referenceOrder,
		stage.ValDurationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValFLTypes,
		stage.ValFLType_stagedOrder,
		stage.ValFLTypes_reference,
		&stage.ValFLTypes_referenceOrder,
		stage.ValFLTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValFrequencyTypes,
		stage.ValFrequencyType_stagedOrder,
		stage.ValFrequencyTypes_reference,
		&stage.ValFrequencyTypes_referenceOrder,
		stage.ValFrequencyTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValFrictionTypes,
		stage.ValFrictionType_stagedOrder,
		stage.ValFrictionTypes_reference,
		&stage.ValFrictionTypes_referenceOrder,
		stage.ValFrictionTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValHexTypes,
		stage.ValHexType_stagedOrder,
		stage.ValHexTypes_reference,
		&stage.ValHexTypes_referenceOrder,
		stage.ValHexTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValLCNTypes,
		stage.ValLCNType_stagedOrder,
		stage.ValLCNTypes_reference,
		&stage.ValLCNTypes_referenceOrder,
		stage.ValLCNTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValLightIntensityTypes,
		stage.ValLightIntensityType_stagedOrder,
		stage.ValLightIntensityTypes_reference,
		&stage.ValLightIntensityTypes_referenceOrder,
		stage.ValLightIntensityTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValMagneticVariationChangeTypes,
		stage.ValMagneticVariationChangeType_stagedOrder,
		stage.ValMagneticVariationChangeTypes_reference,
		&stage.ValMagneticVariationChangeTypes_referenceOrder,
		stage.ValMagneticVariationChangeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValMagneticVariationTypes,
		stage.ValMagneticVariationType_stagedOrder,
		stage.ValMagneticVariationTypes_reference,
		&stage.ValMagneticVariationTypes_referenceOrder,
		stage.ValMagneticVariationTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValPCNTypes,
		stage.ValPCNType_stagedOrder,
		stage.ValPCNTypes_reference,
		&stage.ValPCNTypes_referenceOrder,
		stage.ValPCNTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValPercentTypes,
		stage.ValPercentType_stagedOrder,
		stage.ValPercentTypes_reference,
		&stage.ValPercentTypes_referenceOrder,
		stage.ValPercentTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValPressureTypes,
		stage.ValPressureType_stagedOrder,
		stage.ValPressureTypes_reference,
		&stage.ValPressureTypes_referenceOrder,
		stage.ValPressureTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValSlopeTypes,
		stage.ValSlopeType_stagedOrder,
		stage.ValSlopeTypes_reference,
		&stage.ValSlopeTypes_referenceOrder,
		stage.ValSlopeTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValSpeedTypes,
		stage.ValSpeedType_stagedOrder,
		stage.ValSpeedTypes_reference,
		&stage.ValSpeedTypes_referenceOrder,
		stage.ValSpeedTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValTemperatureTypes,
		stage.ValTemperatureType_stagedOrder,
		stage.ValTemperatureTypes_reference,
		&stage.ValTemperatureTypes_referenceOrder,
		stage.ValTemperatureTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ValWeightTypes,
		stage.ValWeightType_stagedOrder,
		stage.ValWeightTypes_reference,
		&stage.ValWeightTypes_referenceOrder,
		stage.ValWeightTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.XHTMLTypes,
		stage.XHTMLType_stagedOrder,
		stage.XHTMLTypes_reference,
		&stage.XHTMLTypes_referenceOrder,
		stage.XHTMLTypes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.AppendForwardCommit(forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.AppendBackwardCommit(backwardCommit)
		stage.SetModified(true)
	} else {
		stage.SetModified(false)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.CodeATCReportingTypes, &stage.CodeATCReportingTypes_reference, &stage.CodeATCReportingTypes_referenceOrder, &stage.CodeATCReportingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAerialRefuellingPointTypes, &stage.CodeAerialRefuellingPointTypes_reference, &stage.CodeAerialRefuellingPointTypes_referenceOrder, &stage.CodeAerialRefuellingPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAerialRefuellingPrefixTypes, &stage.CodeAerialRefuellingPrefixTypes_reference, &stage.CodeAerialRefuellingPrefixTypes_referenceOrder, &stage.CodeAerialRefuellingPrefixTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAerialRefuellingTypes, &stage.CodeAerialRefuellingTypes_reference, &stage.CodeAerialRefuellingTypes_referenceOrder, &stage.CodeAerialRefuellingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_reference, &stage.CodeAircraftCategoryTypes_referenceOrder, &stage.CodeAircraftCategoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_reference, &stage.CodeAircraftEngineNumberTypes_referenceOrder, &stage.CodeAircraftEngineNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_reference, &stage.CodeAircraftEngineTypes_referenceOrder, &stage.CodeAircraftEngineTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftGroundServiceTypes, &stage.CodeAircraftGroundServiceTypes_reference, &stage.CodeAircraftGroundServiceTypes_referenceOrder, &stage.CodeAircraftGroundServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_reference, &stage.CodeAircraftICAOTypes_referenceOrder, &stage.CodeAircraftICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftStandTypes, &stage.CodeAircraftStandTypes_reference, &stage.CodeAircraftStandTypes_referenceOrder, &stage.CodeAircraftStandTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftTypes, &stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_referenceOrder, &stage.CodeAircraftTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_referenceOrder, &stage.CodeAircraftWingspanClassTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirportHeliportCollocationTypes, &stage.CodeAirportHeliportCollocationTypes_reference, &stage.CodeAirportHeliportCollocationTypes_referenceOrder, &stage.CodeAirportHeliportCollocationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirportHeliportDesignatorTypes, &stage.CodeAirportHeliportDesignatorTypes_reference, &stage.CodeAirportHeliportDesignatorTypes_referenceOrder, &stage.CodeAirportHeliportDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirportHeliportTypes, &stage.CodeAirportHeliportTypes_reference, &stage.CodeAirportHeliportTypes_referenceOrder, &stage.CodeAirportHeliportTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirportWarningTypes, &stage.CodeAirportWarningTypes_reference, &stage.CodeAirportWarningTypes_referenceOrder, &stage.CodeAirportWarningTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceActivityTypes, &stage.CodeAirspaceActivityTypes_reference, &stage.CodeAirspaceActivityTypes_referenceOrder, &stage.CodeAirspaceActivityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceAggregationTypes, &stage.CodeAirspaceAggregationTypes_reference, &stage.CodeAirspaceAggregationTypes_referenceOrder, &stage.CodeAirspaceAggregationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceClassificationTypes, &stage.CodeAirspaceClassificationTypes_reference, &stage.CodeAirspaceClassificationTypes_referenceOrder, &stage.CodeAirspaceClassificationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceDependencyTypes, &stage.CodeAirspaceDependencyTypes_reference, &stage.CodeAirspaceDependencyTypes_referenceOrder, &stage.CodeAirspaceDependencyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceDesignatorTypes, &stage.CodeAirspaceDesignatorTypes_reference, &stage.CodeAirspaceDesignatorTypes_referenceOrder, &stage.CodeAirspaceDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspacePointPositionTypes, &stage.CodeAirspacePointPositionTypes_reference, &stage.CodeAirspacePointPositionTypes_referenceOrder, &stage.CodeAirspacePointPositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspacePointRoleTypes, &stage.CodeAirspacePointRoleTypes_reference, &stage.CodeAirspacePointRoleTypes_referenceOrder, &stage.CodeAirspacePointRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspaceTypes, &stage.CodeAirspaceTypes_reference, &stage.CodeAirspaceTypes_referenceOrder, &stage.CodeAirspaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAltitudeAdjustmentTypes, &stage.CodeAltitudeAdjustmentTypes_reference, &stage.CodeAltitudeAdjustmentTypes_referenceOrder, &stage.CodeAltitudeAdjustmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_referenceOrder, &stage.CodeAltitudeUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachEquipmentAdditionalTypes, &stage.CodeApproachEquipmentAdditionalTypes_reference, &stage.CodeApproachEquipmentAdditionalTypes_referenceOrder, &stage.CodeApproachEquipmentAdditionalTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachGuidanceTypes, &stage.CodeApproachGuidanceTypes_reference, &stage.CodeApproachGuidanceTypes_referenceOrder, &stage.CodeApproachGuidanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachLightingICAOTypes, &stage.CodeApproachLightingICAOTypes_reference, &stage.CodeApproachLightingICAOTypes_referenceOrder, &stage.CodeApproachLightingICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachLightingTypes, &stage.CodeApproachLightingTypes_reference, &stage.CodeApproachLightingTypes_referenceOrder, &stage.CodeApproachLightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachPrefixTypes, &stage.CodeApproachPrefixTypes_reference, &stage.CodeApproachPrefixTypes_referenceOrder, &stage.CodeApproachPrefixTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachTypes, &stage.CodeApproachTypes_reference, &stage.CodeApproachTypes_referenceOrder, &stage.CodeApproachTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApprovalTypes, &stage.CodeApprovalTypes_reference, &stage.CodeApprovalTypes_referenceOrder, &stage.CodeApprovalTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApronElementTypes, &stage.CodeApronElementTypes_reference, &stage.CodeApronElementTypes_referenceOrder, &stage.CodeApronElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApronSectionTypes, &stage.CodeApronSectionTypes_reference, &stage.CodeApronSectionTypes_referenceOrder, &stage.CodeApronSectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeArcDirectionTypes, &stage.CodeArcDirectionTypes_reference, &stage.CodeArcDirectionTypes_referenceOrder, &stage.CodeArcDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeArrestingGearEnergyAbsorbTypes, &stage.CodeArrestingGearEnergyAbsorbTypes_reference, &stage.CodeArrestingGearEnergyAbsorbTypes_referenceOrder, &stage.CodeArrestingGearEnergyAbsorbTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeArrestingGearEngageDeviceTypes, &stage.CodeArrestingGearEngageDeviceTypes_reference, &stage.CodeArrestingGearEngageDeviceTypes_referenceOrder, &stage.CodeArrestingGearEngageDeviceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAuralMorseTypes, &stage.CodeAuralMorseTypes_reference, &stage.CodeAuralMorseTypes_referenceOrder, &stage.CodeAuralMorseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAuthorityRoleTypes, &stage.CodeAuthorityRoleTypes_reference, &stage.CodeAuthorityRoleTypes_referenceOrder, &stage.CodeAuthorityRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAuthorityTypes, &stage.CodeAuthorityTypes_reference, &stage.CodeAuthorityTypes_referenceOrder, &stage.CodeAuthorityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAviationStandardsTypes, &stage.CodeAviationStandardsTypes_reference, &stage.CodeAviationStandardsTypes_referenceOrder, &stage.CodeAviationStandardsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeBearingTypes, &stage.CodeBearingTypes_reference, &stage.CodeBearingTypes_referenceOrder, &stage.CodeBearingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeBuoyDesignatorTypes, &stage.CodeBuoyDesignatorTypes_reference, &stage.CodeBuoyDesignatorTypes_referenceOrder, &stage.CodeBuoyDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeBuoyTypes, &stage.CodeBuoyTypes_reference, &stage.CodeBuoyTypes_referenceOrder, &stage.CodeBuoyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCardinalDirectionTypes, &stage.CodeCardinalDirectionTypes_reference, &stage.CodeCardinalDirectionTypes_referenceOrder, &stage.CodeCardinalDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCheckpointCategoryTypes, &stage.CodeCheckpointCategoryTypes_reference, &stage.CodeCheckpointCategoryTypes_referenceOrder, &stage.CodeCheckpointCategoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeColourTypes, &stage.CodeColourTypes_reference, &stage.CodeColourTypes_referenceOrder, &stage.CodeColourTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationChannelTypes, &stage.CodeCommunicationChannelTypes_reference, &stage.CodeCommunicationChannelTypes_referenceOrder, &stage.CodeCommunicationChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationDirectionTypes, &stage.CodeCommunicationDirectionTypes_reference, &stage.CodeCommunicationDirectionTypes_referenceOrder, &stage.CodeCommunicationDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_referenceOrder, &stage.CodeCommunicationModeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeComparisonTypes, &stage.CodeComparisonTypes_reference, &stage.CodeComparisonTypes_referenceOrder, &stage.CodeComparisonTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeContaminationTypes, &stage.CodeContaminationTypes_reference, &stage.CodeContaminationTypes_referenceOrder, &stage.CodeContaminationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCourseQualityILSTypes, &stage.CodeCourseQualityILSTypes_reference, &stage.CodeCourseQualityILSTypes_referenceOrder, &stage.CodeCourseQualityILSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCourseTypes, &stage.CodeCourseTypes_reference, &stage.CodeCourseTypes_referenceOrder, &stage.CodeCourseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDMEChannelTypes, &stage.CodeDMEChannelTypes_reference, &stage.CodeDMEChannelTypes_referenceOrder, &stage.CodeDMEChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDMETypes, &stage.CodeDMETypes_reference, &stage.CodeDMETypes_referenceOrder, &stage.CodeDMETypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDayTypes, &stage.CodeDayTypes_reference, &stage.CodeDayTypes_referenceOrder, &stage.CodeDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDeclaredDistanceTypes, &stage.CodeDeclaredDistanceTypes_reference, &stage.CodeDeclaredDistanceTypes_referenceOrder, &stage.CodeDeclaredDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDesignStandardTypes, &stage.CodeDesignStandardTypes_reference, &stage.CodeDesignStandardTypes_referenceOrder, &stage.CodeDesignStandardTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDesignatedPointDesignatorTypes, &stage.CodeDesignatedPointDesignatorTypes_reference, &stage.CodeDesignatedPointDesignatorTypes_referenceOrder, &stage.CodeDesignatedPointDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDesignatedPointTypes, &stage.CodeDesignatedPointTypes_reference, &stage.CodeDesignatedPointTypes_referenceOrder, &stage.CodeDesignatedPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionReferenceTypes, &stage.CodeDirectionReferenceTypes_reference, &stage.CodeDirectionReferenceTypes_referenceOrder, &stage.CodeDirectionReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_reference, &stage.CodeDirectionTurnTypes_referenceOrder, &stage.CodeDirectionTurnTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionTypes, &stage.CodeDirectionTypes_reference, &stage.CodeDirectionTypes_referenceOrder, &stage.CodeDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDistanceIndicationTypes, &stage.CodeDistanceIndicationTypes_reference, &stage.CodeDistanceIndicationTypes_referenceOrder, &stage.CodeDistanceIndicationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEmissionBandTypes, &stage.CodeEmissionBandTypes_reference, &stage.CodeEmissionBandTypes_referenceOrder, &stage.CodeEmissionBandTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder, &stage.CodeEquipmentAntiCollisionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEquipmentUnavailableTypes, &stage.CodeEquipmentUnavailableTypes_reference, &stage.CodeEquipmentUnavailableTypes_referenceOrder, &stage.CodeEquipmentUnavailableTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFacilityRankingTypes, &stage.CodeFacilityRankingTypes_reference, &stage.CodeFacilityRankingTypes_referenceOrder, &stage.CodeFacilityRankingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFinalGuidanceTypes, &stage.CodeFinalGuidanceTypes_reference, &stage.CodeFinalGuidanceTypes_referenceOrder, &stage.CodeFinalGuidanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFireFightingTypes, &stage.CodeFireFightingTypes_reference, &stage.CodeFireFightingTypes_referenceOrder, &stage.CodeFireFightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightDestinationTypes, &stage.CodeFlightDestinationTypes_reference, &stage.CodeFlightDestinationTypes_referenceOrder, &stage.CodeFlightDestinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightOriginTypes, &stage.CodeFlightOriginTypes_reference, &stage.CodeFlightOriginTypes_referenceOrder, &stage.CodeFlightOriginTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightPurposeTypes, &stage.CodeFlightPurposeTypes_reference, &stage.CodeFlightPurposeTypes_referenceOrder, &stage.CodeFlightPurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightRestrictionDesignatorTypes, &stage.CodeFlightRestrictionDesignatorTypes_reference, &stage.CodeFlightRestrictionDesignatorTypes_referenceOrder, &stage.CodeFlightRestrictionDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightRestrictionTypes, &stage.CodeFlightRestrictionTypes_reference, &stage.CodeFlightRestrictionTypes_referenceOrder, &stage.CodeFlightRestrictionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_reference, &stage.CodeFlightRuleTypes_referenceOrder, &stage.CodeFlightRuleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightStatusTypes, &stage.CodeFlightStatusTypes_reference, &stage.CodeFlightStatusTypes_referenceOrder, &stage.CodeFlightStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightTypes, &stage.CodeFlightTypes_reference, &stage.CodeFlightTypes_referenceOrder, &stage.CodeFlightTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlowConditionOperationTypes, &stage.CodeFlowConditionOperationTypes_reference, &stage.CodeFlowConditionOperationTypes_referenceOrder, &stage.CodeFlowConditionOperationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFreeFlightTypes, &stage.CodeFreeFlightTypes_reference, &stage.CodeFreeFlightTypes_referenceOrder, &stage.CodeFreeFlightTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFrictionDeviceTypes, &stage.CodeFrictionDeviceTypes_reference, &stage.CodeFrictionDeviceTypes_referenceOrder, &stage.CodeFrictionDeviceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFrictionEstimateTypes, &stage.CodeFrictionEstimateTypes_reference, &stage.CodeFrictionEstimateTypes_referenceOrder, &stage.CodeFrictionEstimateTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFuelTypes, &stage.CodeFuelTypes_reference, &stage.CodeFuelTypes_referenceOrder, &stage.CodeFuelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeGeoBorderTypes, &stage.CodeGeoBorderTypes_reference, &stage.CodeGeoBorderTypes_referenceOrder, &stage.CodeGeoBorderTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeGradeSeparationTypes, &stage.CodeGradeSeparationTypes_reference, &stage.CodeGradeSeparationTypes_referenceOrder, &stage.CodeGradeSeparationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeGroundLightingTypes, &stage.CodeGroundLightingTypes_reference, &stage.CodeGroundLightingTypes_referenceOrder, &stage.CodeGroundLightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeGuidanceLineTypes, &stage.CodeGuidanceLineTypes_reference, &stage.CodeGuidanceLineTypes_referenceOrder, &stage.CodeGuidanceLineTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHeightReferenceTypes, &stage.CodeHeightReferenceTypes_reference, &stage.CodeHeightReferenceTypes_referenceOrder, &stage.CodeHeightReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHelicopterPerformanceTypes, &stage.CodeHelicopterPerformanceTypes_reference, &stage.CodeHelicopterPerformanceTypes_referenceOrder, &stage.CodeHelicopterPerformanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHoldingCategoryTypes, &stage.CodeHoldingCategoryTypes_reference, &stage.CodeHoldingCategoryTypes_referenceOrder, &stage.CodeHoldingCategoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHoldingUsageTypes, &stage.CodeHoldingUsageTypes_reference, &stage.CodeHoldingUsageTypes_referenceOrder, &stage.CodeHoldingUsageTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHoldingUseTypes, &stage.CodeHoldingUseTypes_reference, &stage.CodeHoldingUseTypes_referenceOrder, &stage.CodeHoldingUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeIATATypes, &stage.CodeIATATypes_reference, &stage.CodeIATATypes_referenceOrder, &stage.CodeIATATypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeICAOCountryTypes, &stage.CodeICAOCountryTypes_reference, &stage.CodeICAOCountryTypes_referenceOrder, &stage.CodeICAOCountryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeICAOTypes, &stage.CodeICAOTypes_reference, &stage.CodeICAOTypes_referenceOrder, &stage.CodeICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeILSBackCourseTypes, &stage.CodeILSBackCourseTypes_reference, &stage.CodeILSBackCourseTypes_referenceOrder, &stage.CodeILSBackCourseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeIntegrityLevelILSTypes, &stage.CodeIntegrityLevelILSTypes_reference, &stage.CodeIntegrityLevelILSTypes_referenceOrder, &stage.CodeIntegrityLevelILSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeIntensityStandByTypes, &stage.CodeIntensityStandByTypes_reference, &stage.CodeIntensityStandByTypes_referenceOrder, &stage.CodeIntensityStandByTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLanguageTypes, &stage.CodeLanguageTypes_reference, &stage.CodeLanguageTypes_referenceOrder, &stage.CodeLanguageTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLevelSeriesTypes, &stage.CodeLevelSeriesTypes_reference, &stage.CodeLevelSeriesTypes_referenceOrder, &stage.CodeLevelSeriesTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLevelTableDesignatorTypes, &stage.CodeLevelTableDesignatorTypes_reference, &stage.CodeLevelTableDesignatorTypes_referenceOrder, &stage.CodeLevelTableDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLevelTypes, &stage.CodeLevelTypes_reference, &stage.CodeLevelTypes_referenceOrder, &stage.CodeLevelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightHoldingPositionTypes, &stage.CodeLightHoldingPositionTypes_reference, &stage.CodeLightHoldingPositionTypes_referenceOrder, &stage.CodeLightHoldingPositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_reference, &stage.CodeLightIntensityTypes_referenceOrder, &stage.CodeLightIntensityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightSourceTypes, &stage.CodeLightSourceTypes_reference, &stage.CodeLightSourceTypes_referenceOrder, &stage.CodeLightSourceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightingJARTypes, &stage.CodeLightingJARTypes_reference, &stage.CodeLightingJARTypes_referenceOrder, &stage.CodeLightingJARTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLoadingBridgeTypes, &stage.CodeLoadingBridgeTypes_reference, &stage.CodeLoadingBridgeTypes_referenceOrder, &stage.CodeLoadingBridgeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLocationQualifierTypes, &stage.CodeLocationQualifierTypes_reference, &stage.CodeLocationQualifierTypes_referenceOrder, &stage.CodeLocationQualifierTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLogicalOperatorTypes, &stage.CodeLogicalOperatorTypes_reference, &stage.CodeLogicalOperatorTypes_referenceOrder, &stage.CodeLogicalOperatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMLSAzimuthTypes, &stage.CodeMLSAzimuthTypes_reference, &stage.CodeMLSAzimuthTypes_referenceOrder, &stage.CodeMLSAzimuthTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMLSChannelTypes, &stage.CodeMLSChannelTypes_reference, &stage.CodeMLSChannelTypes_referenceOrder, &stage.CodeMLSChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMarkerBeaconSignalTypes, &stage.CodeMarkerBeaconSignalTypes_reference, &stage.CodeMarkerBeaconSignalTypes_referenceOrder, &stage.CodeMarkerBeaconSignalTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMarkingConditionTypes, &stage.CodeMarkingConditionTypes_reference, &stage.CodeMarkingConditionTypes_referenceOrder, &stage.CodeMarkingConditionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMarkingStyleTypes, &stage.CodeMarkingStyleTypes_reference, &stage.CodeMarkingStyleTypes_referenceOrder, &stage.CodeMarkingStyleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMeteoConditionsTypes, &stage.CodeMeteoConditionsTypes_reference, &stage.CodeMeteoConditionsTypes_referenceOrder, &stage.CodeMeteoConditionsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_reference, &stage.CodeMilitaryOperationsTypes_referenceOrder, &stage.CodeMilitaryOperationsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryRoutePointTypes, &stage.CodeMilitaryRoutePointTypes_reference, &stage.CodeMilitaryRoutePointTypes_referenceOrder, &stage.CodeMilitaryRoutePointTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryStatusTypes, &stage.CodeMilitaryStatusTypes_reference, &stage.CodeMilitaryStatusTypes_referenceOrder, &stage.CodeMilitaryStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryTrainingTypes, &stage.CodeMilitaryTrainingTypes_reference, &stage.CodeMilitaryTrainingTypes_referenceOrder, &stage.CodeMilitaryTrainingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMinimaFinalApproachPathTypes, &stage.CodeMinimaFinalApproachPathTypes_reference, &stage.CodeMinimaFinalApproachPathTypes_referenceOrder, &stage.CodeMinimaFinalApproachPathTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMinimumAltitudeTypes, &stage.CodeMinimumAltitudeTypes_reference, &stage.CodeMinimumAltitudeTypes_referenceOrder, &stage.CodeMinimumAltitudeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMinimumHeightTypes, &stage.CodeMinimumHeightTypes_reference, &stage.CodeMinimumHeightTypes_referenceOrder, &stage.CodeMinimumHeightTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMissedApproachTypes, &stage.CodeMissedApproachTypes_reference, &stage.CodeMissedApproachTypes_referenceOrder, &stage.CodeMissedApproachTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNDBUsageTypes, &stage.CodeNDBUsageTypes_reference, &stage.CodeNDBUsageTypes_referenceOrder, &stage.CodeNDBUsageTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavaidDesignatorTypes, &stage.CodeNavaidDesignatorTypes_reference, &stage.CodeNavaidDesignatorTypes_referenceOrder, &stage.CodeNavaidDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavaidPurposeTypes, &stage.CodeNavaidPurposeTypes_reference, &stage.CodeNavaidPurposeTypes_referenceOrder, &stage.CodeNavaidPurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavaidServiceTypes, &stage.CodeNavaidServiceTypes_reference, &stage.CodeNavaidServiceTypes_referenceOrder, &stage.CodeNavaidServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationAreaRestrictionTypes, &stage.CodeNavigationAreaRestrictionTypes_reference, &stage.CodeNavigationAreaRestrictionTypes_referenceOrder, &stage.CodeNavigationAreaRestrictionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationAreaTypes, &stage.CodeNavigationAreaTypes_reference, &stage.CodeNavigationAreaTypes_referenceOrder, &stage.CodeNavigationAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_referenceOrder, &stage.CodeNavigationEquipmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_referenceOrder, &stage.CodeNavigationSpecificationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNitrogenTypes, &stage.CodeNitrogenTypes_reference, &stage.CodeNitrogenTypes_referenceOrder, &stage.CodeNitrogenTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNorthReferenceTypes, &stage.CodeNorthReferenceTypes_reference, &stage.CodeNorthReferenceTypes_referenceOrder, &stage.CodeNorthReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_referenceOrder, &stage.CodeNotePurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstacleAreaTypes, &stage.CodeObstacleAreaTypes_reference, &stage.CodeObstacleAreaTypes_referenceOrder, &stage.CodeObstacleAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_reference, &stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder, &stage.CodeObstacleAssessmentSurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstructionIdSurfaceZoneTypes, &stage.CodeObstructionIdSurfaceZoneTypes_reference, &stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder, &stage.CodeObstructionIdSurfaceZoneTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOilTypes, &stage.CodeOilTypes_reference, &stage.CodeOilTypes_referenceOrder, &stage.CodeOilTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOperationAirportHeliportTypes, &stage.CodeOperationAirportHeliportTypes_reference, &stage.CodeOperationAirportHeliportTypes_referenceOrder, &stage.CodeOperationAirportHeliportTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOperationManoeuvringAreaTypes, &stage.CodeOperationManoeuvringAreaTypes_reference, &stage.CodeOperationManoeuvringAreaTypes_referenceOrder, &stage.CodeOperationManoeuvringAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOrganisationDesignatorTypes, &stage.CodeOrganisationDesignatorTypes_reference, &stage.CodeOrganisationDesignatorTypes_referenceOrder, &stage.CodeOrganisationDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOrganisationHierarchyTypes, &stage.CodeOrganisationHierarchyTypes_reference, &stage.CodeOrganisationHierarchyTypes_referenceOrder, &stage.CodeOrganisationHierarchyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOrganisationTypes, &stage.CodeOrganisationTypes_reference, &stage.CodeOrganisationTypes_referenceOrder, &stage.CodeOrganisationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOxygenTypes, &stage.CodeOxygenTypes_reference, &stage.CodeOxygenTypes_referenceOrder, &stage.CodeOxygenTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePARTypes, &stage.CodePARTypes_reference, &stage.CodePARTypes_referenceOrder, &stage.CodePARTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePCNMethodTypes, &stage.CodePCNMethodTypes_reference, &stage.CodePCNMethodTypes_referenceOrder, &stage.CodePCNMethodTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePCNPavementTypes, &stage.CodePCNPavementTypes_reference, &stage.CodePCNPavementTypes_referenceOrder, &stage.CodePCNPavementTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePCNSubgradeTypes, &stage.CodePCNSubgradeTypes_reference, &stage.CodePCNSubgradeTypes_referenceOrder, &stage.CodePCNSubgradeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePCNTyrePressureTypes, &stage.CodePCNTyrePressureTypes_reference, &stage.CodePCNTyrePressureTypes_referenceOrder, &stage.CodePCNTyrePressureTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePassengerServiceTypes, &stage.CodePassengerServiceTypes_reference, &stage.CodePassengerServiceTypes_referenceOrder, &stage.CodePassengerServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePilotControlledLightingTypes, &stage.CodePilotControlledLightingTypes_reference, &stage.CodePilotControlledLightingTypes_referenceOrder, &stage.CodePilotControlledLightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePositionInILSTypes, &stage.CodePositionInILSTypes_reference, &stage.CodePositionInILSTypes_referenceOrder, &stage.CodePositionInILSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePrimaryRadarTypes, &stage.CodePrimaryRadarTypes_reference, &stage.CodePrimaryRadarTypes_referenceOrder, &stage.CodePrimaryRadarTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedureAvailabilityTypes, &stage.CodeProcedureAvailabilityTypes_reference, &stage.CodeProcedureAvailabilityTypes_referenceOrder, &stage.CodeProcedureAvailabilityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedureCodingStandardTypes, &stage.CodeProcedureCodingStandardTypes_reference, &stage.CodeProcedureCodingStandardTypes_referenceOrder, &stage.CodeProcedureCodingStandardTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedureDistanceTypes, &stage.CodeProcedureDistanceTypes_reference, &stage.CodeProcedureDistanceTypes_referenceOrder, &stage.CodeProcedureDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedureFixRoleTypes, &stage.CodeProcedureFixRoleTypes_reference, &stage.CodeProcedureFixRoleTypes_referenceOrder, &stage.CodeProcedureFixRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedurePhaseTypes, &stage.CodeProcedurePhaseTypes_reference, &stage.CodeProcedurePhaseTypes_referenceOrder, &stage.CodeProcedurePhaseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProtectAreaSectionTypes, &stage.CodeProtectAreaSectionTypes_reference, &stage.CodeProtectAreaSectionTypes_referenceOrder, &stage.CodeProtectAreaSectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRNPTypes, &stage.CodeRNPTypes_reference, &stage.CodeRNPTypes_referenceOrder, &stage.CodeRNPTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVRReadingTypes, &stage.CodeRVRReadingTypes_reference, &stage.CodeRVRReadingTypes_referenceOrder, &stage.CodeRVRReadingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMPointRoleTypes, &stage.CodeRVSMPointRoleTypes_reference, &stage.CodeRVSMPointRoleTypes_referenceOrder, &stage.CodeRVSMPointRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMTypes, &stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_referenceOrder, &stage.CodeRVSMTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRadarServiceTypes, &stage.CodeRadarServiceTypes_reference, &stage.CodeRadarServiceTypes_referenceOrder, &stage.CodeRadarServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRadioEmissionTypes, &stage.CodeRadioEmissionTypes_reference, &stage.CodeRadioEmissionTypes_referenceOrder, &stage.CodeRadioEmissionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRadioFrequencyAreaTypes, &stage.CodeRadioFrequencyAreaTypes_reference, &stage.CodeRadioFrequencyAreaTypes_referenceOrder, &stage.CodeRadioFrequencyAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRadioSignalTypes, &stage.CodeRadioSignalTypes_reference, &stage.CodeRadioSignalTypes_referenceOrder, &stage.CodeRadioSignalTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeReferencePathIdentifierTypes, &stage.CodeReferencePathIdentifierTypes_reference, &stage.CodeReferencePathIdentifierTypes_referenceOrder, &stage.CodeReferencePathIdentifierTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeReferenceRoleTypes, &stage.CodeReferenceRoleTypes_reference, &stage.CodeReferenceRoleTypes_referenceOrder, &stage.CodeReferenceRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeReflectorTypes, &stage.CodeReflectorTypes_reference, &stage.CodeReflectorTypes_referenceOrder, &stage.CodeReflectorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRelativePositionTypes, &stage.CodeRelativePositionTypes_reference, &stage.CodeRelativePositionTypes_referenceOrder, &stage.CodeRelativePositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRoadTypes, &stage.CodeRoadTypes_reference, &stage.CodeRoadTypes_referenceOrder, &stage.CodeRoadTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteAvailabilityTypes, &stage.CodeRouteAvailabilityTypes_reference, &stage.CodeRouteAvailabilityTypes_referenceOrder, &stage.CodeRouteAvailabilityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteDesignatorLetterTypes, &stage.CodeRouteDesignatorLetterTypes_reference, &stage.CodeRouteDesignatorLetterTypes_referenceOrder, &stage.CodeRouteDesignatorLetterTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteDesignatorPrefixTypes, &stage.CodeRouteDesignatorPrefixTypes_reference, &stage.CodeRouteDesignatorPrefixTypes_referenceOrder, &stage.CodeRouteDesignatorPrefixTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteDesignatorSuffixTypes, &stage.CodeRouteDesignatorSuffixTypes_reference, &stage.CodeRouteDesignatorSuffixTypes_referenceOrder, &stage.CodeRouteDesignatorSuffixTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteIndicatorTypes, &stage.CodeRouteIndicatorTypes_reference, &stage.CodeRouteIndicatorTypes_referenceOrder, &stage.CodeRouteIndicatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteNavigationTypes, &stage.CodeRouteNavigationTypes_reference, &stage.CodeRouteNavigationTypes_referenceOrder, &stage.CodeRouteNavigationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteOriginTypes, &stage.CodeRouteOriginTypes_reference, &stage.CodeRouteOriginTypes_referenceOrder, &stage.CodeRouteOriginTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteSegmentPathTypes, &stage.CodeRouteSegmentPathTypes_reference, &stage.CodeRouteSegmentPathTypes_referenceOrder, &stage.CodeRouteSegmentPathTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteTypes, &stage.CodeRouteTypes_reference, &stage.CodeRouteTypes_referenceOrder, &stage.CodeRouteTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRuleProcedureTitleTypes, &stage.CodeRuleProcedureTitleTypes_reference, &stage.CodeRuleProcedureTitleTypes_referenceOrder, &stage.CodeRuleProcedureTitleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRuleProcedureTypes, &stage.CodeRuleProcedureTypes_reference, &stage.CodeRuleProcedureTypes_referenceOrder, &stage.CodeRuleProcedureTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayElementTypes, &stage.CodeRunwayElementTypes_reference, &stage.CodeRunwayElementTypes_referenceOrder, &stage.CodeRunwayElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayMarkingTypes, &stage.CodeRunwayMarkingTypes_reference, &stage.CodeRunwayMarkingTypes_referenceOrder, &stage.CodeRunwayMarkingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayPointRoleTypes, &stage.CodeRunwayPointRoleTypes_reference, &stage.CodeRunwayPointRoleTypes_referenceOrder, &stage.CodeRunwayPointRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayProtectionAreaTypes, &stage.CodeRunwayProtectionAreaTypes_reference, &stage.CodeRunwayProtectionAreaTypes_referenceOrder, &stage.CodeRunwayProtectionAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwaySectionTypes, &stage.CodeRunwaySectionTypes_reference, &stage.CodeRunwaySectionTypes_referenceOrder, &stage.CodeRunwaySectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayTypes, &stage.CodeRunwayTypes_reference, &stage.CodeRunwayTypes_referenceOrder, &stage.CodeRunwayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSafeAltitudeTypes, &stage.CodeSafeAltitudeTypes_reference, &stage.CodeSafeAltitudeTypes_referenceOrder, &stage.CodeSafeAltitudeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSegmentPathTypes, &stage.CodeSegmentPathTypes_reference, &stage.CodeSegmentPathTypes_referenceOrder, &stage.CodeSegmentPathTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSegmentTerminationTypes, &stage.CodeSegmentTerminationTypes_reference, &stage.CodeSegmentTerminationTypes_referenceOrder, &stage.CodeSegmentTerminationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceATCTypes, &stage.CodeServiceATCTypes_reference, &stage.CodeServiceATCTypes_referenceOrder, &stage.CodeServiceATCTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceATFMTypes, &stage.CodeServiceATFMTypes_reference, &stage.CodeServiceATFMTypes_referenceOrder, &stage.CodeServiceATFMTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceGroundControlTypes, &stage.CodeServiceGroundControlTypes_reference, &stage.CodeServiceGroundControlTypes_referenceOrder, &stage.CodeServiceGroundControlTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceInformationTypes, &stage.CodeServiceInformationTypes_reference, &stage.CodeServiceInformationTypes_referenceOrder, &stage.CodeServiceInformationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceSARTypes, &stage.CodeServiceSARTypes_reference, &stage.CodeServiceSARTypes_referenceOrder, &stage.CodeServiceSARTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSideTypes, &stage.CodeSideTypes_reference, &stage.CodeSideTypes_referenceOrder, &stage.CodeSideTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSignalPerformanceILSTypes, &stage.CodeSignalPerformanceILSTypes_reference, &stage.CodeSignalPerformanceILSTypes_referenceOrder, &stage.CodeSignalPerformanceILSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpecialDateTypes, &stage.CodeSpecialDateTypes_reference, &stage.CodeSpecialDateTypes_referenceOrder, &stage.CodeSpecialDateTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpecialNavigationChainDesignatorTypes, &stage.CodeSpecialNavigationChainDesignatorTypes_reference, &stage.CodeSpecialNavigationChainDesignatorTypes_referenceOrder, &stage.CodeSpecialNavigationChainDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpecialNavigationStationTypes, &stage.CodeSpecialNavigationStationTypes_reference, &stage.CodeSpecialNavigationStationTypes_referenceOrder, &stage.CodeSpecialNavigationStationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpecialNavigationSystemTypes, &stage.CodeSpecialNavigationSystemTypes_reference, &stage.CodeSpecialNavigationSystemTypes_referenceOrder, &stage.CodeSpecialNavigationSystemTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpeedReferenceTypes, &stage.CodeSpeedReferenceTypes_reference, &stage.CodeSpeedReferenceTypes_referenceOrder, &stage.CodeSpeedReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStandbyPowerTypes, &stage.CodeStandbyPowerTypes_reference, &stage.CodeStandbyPowerTypes_referenceOrder, &stage.CodeStandbyPowerTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusAirportTypes, &stage.CodeStatusAirportTypes_reference, &stage.CodeStatusAirportTypes_referenceOrder, &stage.CodeStatusAirportTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusAirspaceTypes, &stage.CodeStatusAirspaceTypes_reference, &stage.CodeStatusAirspaceTypes_referenceOrder, &stage.CodeStatusAirspaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusConstructionTypes, &stage.CodeStatusConstructionTypes_reference, &stage.CodeStatusConstructionTypes_referenceOrder, &stage.CodeStatusConstructionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusNavaidTypes, &stage.CodeStatusNavaidTypes_reference, &stage.CodeStatusNavaidTypes_referenceOrder, &stage.CodeStatusNavaidTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_reference, &stage.CodeStatusOperationsTypes_referenceOrder, &stage.CodeStatusOperationsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusServiceTypes, &stage.CodeStatusServiceTypes_reference, &stage.CodeStatusServiceTypes_referenceOrder, &stage.CodeStatusServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSurfaceCompositionTypes, &stage.CodeSurfaceCompositionTypes_reference, &stage.CodeSurfaceCompositionTypes_referenceOrder, &stage.CodeSurfaceCompositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSurfaceConditionTypes, &stage.CodeSurfaceConditionTypes_reference, &stage.CodeSurfaceConditionTypes_referenceOrder, &stage.CodeSurfaceConditionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSurfacePreparationTypes, &stage.CodeSurfacePreparationTypes_reference, &stage.CodeSurfacePreparationTypes_referenceOrder, &stage.CodeSurfacePreparationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSystemActivationTypes, &stage.CodeSystemActivationTypes_reference, &stage.CodeSystemActivationTypes_referenceOrder, &stage.CodeSystemActivationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTAATypes, &stage.CodeTAATypes_reference, &stage.CodeTAATypes_referenceOrder, &stage.CodeTAATypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTACANChannelTypes, &stage.CodeTACANChannelTypes_reference, &stage.CodeTACANChannelTypes_referenceOrder, &stage.CodeTACANChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTLOFSectionTypes, &stage.CodeTLOFSectionTypes_reference, &stage.CodeTLOFSectionTypes_referenceOrder, &stage.CodeTLOFSectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTaxiwayElementTypes, &stage.CodeTaxiwayElementTypes_reference, &stage.CodeTaxiwayElementTypes_referenceOrder, &stage.CodeTaxiwayElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTaxiwaySectionTypes, &stage.CodeTaxiwaySectionTypes_reference, &stage.CodeTaxiwaySectionTypes_referenceOrder, &stage.CodeTaxiwaySectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTaxiwayTypes, &stage.CodeTaxiwayTypes_reference, &stage.CodeTaxiwayTypes_referenceOrder, &stage.CodeTaxiwayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_referenceOrder, &stage.CodeTelecomNetworkTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_referenceOrder, &stage.CodeTimeEventCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_referenceOrder, &stage.CodeTimeEventTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_referenceOrder, &stage.CodeTimeReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTrajectoryTypes, &stage.CodeTrajectoryTypes_reference, &stage.CodeTrajectoryTypes_referenceOrder, &stage.CodeTrajectoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTransponderTypes, &stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_referenceOrder, &stage.CodeTransponderTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUUIDTypes, &stage.CodeUUIDTypes_reference, &stage.CodeUUIDTypes_referenceOrder, &stage.CodeUUIDTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUnitDependencyTypes, &stage.CodeUnitDependencyTypes_reference, &stage.CodeUnitDependencyTypes_referenceOrder, &stage.CodeUnitDependencyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUnitTypes, &stage.CodeUnitTypes_reference, &stage.CodeUnitTypes_referenceOrder, &stage.CodeUnitTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUpperAlphaTypes, &stage.CodeUpperAlphaTypes_reference, &stage.CodeUpperAlphaTypes_referenceOrder, &stage.CodeUpperAlphaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUsageLimitationTypes, &stage.CodeUsageLimitationTypes_reference, &stage.CodeUsageLimitationTypes_referenceOrder, &stage.CodeUsageLimitationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVASISTypes, &stage.CodeVASISTypes_reference, &stage.CodeVASISTypes_referenceOrder, &stage.CodeVASISTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVORTypes, &stage.CodeVORTypes_reference, &stage.CodeVORTypes_referenceOrder, &stage.CodeVORTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_reference, &stage.CodeValueInterpretationTypes_referenceOrder, &stage.CodeValueInterpretationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_referenceOrder, &stage.CodeVerticalDatumTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_reference, &stage.CodeVerticalReferenceTypes_referenceOrder, &stage.CodeVerticalReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalStructureMarkingTypes, &stage.CodeVerticalStructureMarkingTypes_reference, &stage.CodeVerticalStructureMarkingTypes_referenceOrder, &stage.CodeVerticalStructureMarkingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalStructureMaterialTypes, &stage.CodeVerticalStructureMaterialTypes_reference, &stage.CodeVerticalStructureMaterialTypes_referenceOrder, &stage.CodeVerticalStructureMaterialTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalStructureTypes, &stage.CodeVerticalStructureTypes_reference, &stage.CodeVerticalStructureTypes_referenceOrder, &stage.CodeVerticalStructureTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVisualDockingGuidanceTypes, &stage.CodeVisualDockingGuidanceTypes_reference, &stage.CodeVisualDockingGuidanceTypes_referenceOrder, &stage.CodeVisualDockingGuidanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_reference, &stage.CodeWakeTurbulenceTypes_referenceOrder, &stage.CodeWakeTurbulenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeWorkAreaTypes, &stage.CodeWorkAreaTypes_reference, &stage.CodeWorkAreaTypes_referenceOrder, &stage.CodeWorkAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeYesNoTypes, &stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_referenceOrder, &stage.CodeYesNoTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateMonthDayTypes, &stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_referenceOrder, &stage.DateMonthDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateTimeTypes, &stage.DateTimeTypes_reference, &stage.DateTimeTypes_referenceOrder, &stage.DateTimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateTypes, &stage.DateTypes_reference, &stage.DateTypes_referenceOrder, &stage.DateTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateYearTypes, &stage.DateYearTypes_reference, &stage.DateYearTypes_referenceOrder, &stage.DateYearTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoNumberTypes, &stage.NoNumberTypes_reference, &stage.NoNumberTypes_referenceOrder, &stage.NoNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoSequenceTypes, &stage.NoSequenceTypes_reference, &stage.NoSequenceTypes_referenceOrder, &stage.NoSequenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextAddressTypes, &stage.TextAddressTypes_reference, &stage.TextAddressTypes_referenceOrder, &stage.TextAddressTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextDesignatorTypes, &stage.TextDesignatorTypes_reference, &stage.TextDesignatorTypes_referenceOrder, &stage.TextDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextInstructionTypes, &stage.TextInstructionTypes_reference, &stage.TextInstructionTypes_referenceOrder, &stage.TextInstructionTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNameTypes, &stage.TextNameTypes_reference, &stage.TextNameTypes_referenceOrder, &stage.TextNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNoteTypes, &stage.TextNoteTypes_reference, &stage.TextNoteTypes_referenceOrder, &stage.TextNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPhoneTypes, &stage.TextPhoneTypes_reference, &stage.TextPhoneTypes_referenceOrder, &stage.TextPhoneTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_referenceOrder, &stage.TextPropertyNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextRemarkTypes, &stage.TextRemarkTypes_reference, &stage.TextRemarkTypes_referenceOrder, &stage.TextRemarkTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextSIDSTARDesignatorTypes, &stage.TextSIDSTARDesignatorTypes_reference, &stage.TextSIDSTARDesignatorTypes_referenceOrder, &stage.TextSIDSTARDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimeTypes, &stage.TimeTypes_reference, &stage.TimeTypes_referenceOrder, &stage.TimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValAlarmLimitTypes, &stage.ValAlarmLimitTypes_reference, &stage.ValAlarmLimitTypes_referenceOrder, &stage.ValAlarmLimitTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValAngleTypes, &stage.ValAngleTypes_reference, &stage.ValAngleTypes_referenceOrder, &stage.ValAngleTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValBearingTypes, &stage.ValBearingTypes_reference, &stage.ValBearingTypes_referenceOrder, &stage.ValBearingTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValChannelNumberTypes, &stage.ValChannelNumberTypes_reference, &stage.ValChannelNumberTypes_referenceOrder, &stage.ValChannelNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDepthTypes, &stage.ValDepthTypes_reference, &stage.ValDepthTypes_referenceOrder, &stage.ValDepthTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_referenceOrder, &stage.ValDistanceSignedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceTypes, &stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_referenceOrder, &stage.ValDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_referenceOrder, &stage.ValDistanceVerticalTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDurationTypes, &stage.ValDurationTypes_reference, &stage.ValDurationTypes_referenceOrder, &stage.ValDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValFLTypes, &stage.ValFLTypes_reference, &stage.ValFLTypes_referenceOrder, &stage.ValFLTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValFrequencyTypes, &stage.ValFrequencyTypes_reference, &stage.ValFrequencyTypes_referenceOrder, &stage.ValFrequencyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValFrictionTypes, &stage.ValFrictionTypes_reference, &stage.ValFrictionTypes_referenceOrder, &stage.ValFrictionTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValHexTypes, &stage.ValHexTypes_reference, &stage.ValHexTypes_referenceOrder, &stage.ValHexTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValLCNTypes, &stage.ValLCNTypes_reference, &stage.ValLCNTypes_referenceOrder, &stage.ValLCNTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValLightIntensityTypes, &stage.ValLightIntensityTypes_reference, &stage.ValLightIntensityTypes_referenceOrder, &stage.ValLightIntensityTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValMagneticVariationChangeTypes, &stage.ValMagneticVariationChangeTypes_reference, &stage.ValMagneticVariationChangeTypes_referenceOrder, &stage.ValMagneticVariationChangeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValMagneticVariationTypes, &stage.ValMagneticVariationTypes_reference, &stage.ValMagneticVariationTypes_referenceOrder, &stage.ValMagneticVariationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValPCNTypes, &stage.ValPCNTypes_reference, &stage.ValPCNTypes_referenceOrder, &stage.ValPCNTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValPercentTypes, &stage.ValPercentTypes_reference, &stage.ValPercentTypes_referenceOrder, &stage.ValPercentTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValPressureTypes, &stage.ValPressureTypes_reference, &stage.ValPressureTypes_referenceOrder, &stage.ValPressureTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValSlopeTypes, &stage.ValSlopeTypes_reference, &stage.ValSlopeTypes_referenceOrder, &stage.ValSlopeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValSpeedTypes, &stage.ValSpeedTypes_reference, &stage.ValSpeedTypes_referenceOrder, &stage.ValSpeedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValTemperatureTypes, &stage.ValTemperatureTypes_reference, &stage.ValTemperatureTypes_referenceOrder, &stage.ValTemperatureTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValWeightTypes, &stage.ValWeightTypes_reference, &stage.ValWeightTypes_referenceOrder, &stage.ValWeightTypes_instance)

	__gong__computeReferencePass1(stage, stage.XHTMLTypes, &stage.XHTMLTypes_reference, &stage.XHTMLTypes_referenceOrder, &stage.XHTMLTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CodeATCReportingTypes, stage.CodeATCReportingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAerialRefuellingPointTypes, stage.CodeAerialRefuellingPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAerialRefuellingPrefixTypes, stage.CodeAerialRefuellingPrefixTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAerialRefuellingTypes, stage.CodeAerialRefuellingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftGroundServiceTypes, stage.CodeAircraftGroundServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftStandTypes, stage.CodeAircraftStandTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftTypes, stage.CodeAircraftTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirportHeliportCollocationTypes, stage.CodeAirportHeliportCollocationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirportHeliportDesignatorTypes, stage.CodeAirportHeliportDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirportHeliportTypes, stage.CodeAirportHeliportTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirportWarningTypes, stage.CodeAirportWarningTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceActivityTypes, stage.CodeAirspaceActivityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceAggregationTypes, stage.CodeAirspaceAggregationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceClassificationTypes, stage.CodeAirspaceClassificationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceDependencyTypes, stage.CodeAirspaceDependencyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceDesignatorTypes, stage.CodeAirspaceDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspacePointPositionTypes, stage.CodeAirspacePointPositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspacePointRoleTypes, stage.CodeAirspacePointRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspaceTypes, stage.CodeAirspaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAltitudeAdjustmentTypes, stage.CodeAltitudeAdjustmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachEquipmentAdditionalTypes, stage.CodeApproachEquipmentAdditionalTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachGuidanceTypes, stage.CodeApproachGuidanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachLightingICAOTypes, stage.CodeApproachLightingICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachLightingTypes, stage.CodeApproachLightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachPrefixTypes, stage.CodeApproachPrefixTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachTypes, stage.CodeApproachTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApprovalTypes, stage.CodeApprovalTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApronElementTypes, stage.CodeApronElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApronSectionTypes, stage.CodeApronSectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeArcDirectionTypes, stage.CodeArcDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeArrestingGearEnergyAbsorbTypes, stage.CodeArrestingGearEnergyAbsorbTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeArrestingGearEngageDeviceTypes, stage.CodeArrestingGearEngageDeviceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAuralMorseTypes, stage.CodeAuralMorseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAuthorityRoleTypes, stage.CodeAuthorityRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAuthorityTypes, stage.CodeAuthorityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAviationStandardsTypes, stage.CodeAviationStandardsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeBearingTypes, stage.CodeBearingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeBuoyDesignatorTypes, stage.CodeBuoyDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeBuoyTypes, stage.CodeBuoyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCardinalDirectionTypes, stage.CodeCardinalDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCheckpointCategoryTypes, stage.CodeCheckpointCategoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeColourTypes, stage.CodeColourTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationChannelTypes, stage.CodeCommunicationChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationDirectionTypes, stage.CodeCommunicationDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeComparisonTypes, stage.CodeComparisonTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeContaminationTypes, stage.CodeContaminationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCourseQualityILSTypes, stage.CodeCourseQualityILSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCourseTypes, stage.CodeCourseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDMEChannelTypes, stage.CodeDMEChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDMETypes, stage.CodeDMETypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDayTypes, stage.CodeDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDeclaredDistanceTypes, stage.CodeDeclaredDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDesignStandardTypes, stage.CodeDesignStandardTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDesignatedPointDesignatorTypes, stage.CodeDesignatedPointDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDesignatedPointTypes, stage.CodeDesignatedPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionReferenceTypes, stage.CodeDirectionReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionTypes, stage.CodeDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDistanceIndicationTypes, stage.CodeDistanceIndicationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEmissionBandTypes, stage.CodeEmissionBandTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEquipmentUnavailableTypes, stage.CodeEquipmentUnavailableTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFacilityRankingTypes, stage.CodeFacilityRankingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFinalGuidanceTypes, stage.CodeFinalGuidanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFireFightingTypes, stage.CodeFireFightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightDestinationTypes, stage.CodeFlightDestinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightOriginTypes, stage.CodeFlightOriginTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightPurposeTypes, stage.CodeFlightPurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightRestrictionDesignatorTypes, stage.CodeFlightRestrictionDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightRestrictionTypes, stage.CodeFlightRestrictionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightRuleTypes, stage.CodeFlightRuleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightStatusTypes, stage.CodeFlightStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightTypes, stage.CodeFlightTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlowConditionOperationTypes, stage.CodeFlowConditionOperationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFreeFlightTypes, stage.CodeFreeFlightTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFrictionDeviceTypes, stage.CodeFrictionDeviceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFrictionEstimateTypes, stage.CodeFrictionEstimateTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFuelTypes, stage.CodeFuelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeGeoBorderTypes, stage.CodeGeoBorderTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeGradeSeparationTypes, stage.CodeGradeSeparationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeGroundLightingTypes, stage.CodeGroundLightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeGuidanceLineTypes, stage.CodeGuidanceLineTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHeightReferenceTypes, stage.CodeHeightReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHelicopterPerformanceTypes, stage.CodeHelicopterPerformanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHoldingCategoryTypes, stage.CodeHoldingCategoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHoldingUsageTypes, stage.CodeHoldingUsageTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHoldingUseTypes, stage.CodeHoldingUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeIATATypes, stage.CodeIATATypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeICAOCountryTypes, stage.CodeICAOCountryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeICAOTypes, stage.CodeICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeILSBackCourseTypes, stage.CodeILSBackCourseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeIntegrityLevelILSTypes, stage.CodeIntegrityLevelILSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeIntensityStandByTypes, stage.CodeIntensityStandByTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLanguageTypes, stage.CodeLanguageTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLevelSeriesTypes, stage.CodeLevelSeriesTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLevelTableDesignatorTypes, stage.CodeLevelTableDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLevelTypes, stage.CodeLevelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightHoldingPositionTypes, stage.CodeLightHoldingPositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightIntensityTypes, stage.CodeLightIntensityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightSourceTypes, stage.CodeLightSourceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightingJARTypes, stage.CodeLightingJARTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLoadingBridgeTypes, stage.CodeLoadingBridgeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLocationQualifierTypes, stage.CodeLocationQualifierTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLogicalOperatorTypes, stage.CodeLogicalOperatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMLSAzimuthTypes, stage.CodeMLSAzimuthTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMLSChannelTypes, stage.CodeMLSChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMarkerBeaconSignalTypes, stage.CodeMarkerBeaconSignalTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMarkingConditionTypes, stage.CodeMarkingConditionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMarkingStyleTypes, stage.CodeMarkingStyleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMeteoConditionsTypes, stage.CodeMeteoConditionsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryRoutePointTypes, stage.CodeMilitaryRoutePointTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryStatusTypes, stage.CodeMilitaryStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryTrainingTypes, stage.CodeMilitaryTrainingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMinimaFinalApproachPathTypes, stage.CodeMinimaFinalApproachPathTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMinimumAltitudeTypes, stage.CodeMinimumAltitudeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMinimumHeightTypes, stage.CodeMinimumHeightTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMissedApproachTypes, stage.CodeMissedApproachTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNDBUsageTypes, stage.CodeNDBUsageTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavaidDesignatorTypes, stage.CodeNavaidDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavaidPurposeTypes, stage.CodeNavaidPurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavaidServiceTypes, stage.CodeNavaidServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationAreaRestrictionTypes, stage.CodeNavigationAreaRestrictionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationAreaTypes, stage.CodeNavigationAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNitrogenTypes, stage.CodeNitrogenTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNorthReferenceTypes, stage.CodeNorthReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstacleAreaTypes, stage.CodeObstacleAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstructionIdSurfaceZoneTypes, stage.CodeObstructionIdSurfaceZoneTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOilTypes, stage.CodeOilTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOperationAirportHeliportTypes, stage.CodeOperationAirportHeliportTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOperationManoeuvringAreaTypes, stage.CodeOperationManoeuvringAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOrganisationDesignatorTypes, stage.CodeOrganisationDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOrganisationHierarchyTypes, stage.CodeOrganisationHierarchyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOrganisationTypes, stage.CodeOrganisationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOxygenTypes, stage.CodeOxygenTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePARTypes, stage.CodePARTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePCNMethodTypes, stage.CodePCNMethodTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePCNPavementTypes, stage.CodePCNPavementTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePCNSubgradeTypes, stage.CodePCNSubgradeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePCNTyrePressureTypes, stage.CodePCNTyrePressureTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePassengerServiceTypes, stage.CodePassengerServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePilotControlledLightingTypes, stage.CodePilotControlledLightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePositionInILSTypes, stage.CodePositionInILSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePrimaryRadarTypes, stage.CodePrimaryRadarTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedureAvailabilityTypes, stage.CodeProcedureAvailabilityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedureCodingStandardTypes, stage.CodeProcedureCodingStandardTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedureDistanceTypes, stage.CodeProcedureDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedureFixRoleTypes, stage.CodeProcedureFixRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedurePhaseTypes, stage.CodeProcedurePhaseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProtectAreaSectionTypes, stage.CodeProtectAreaSectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRNPTypes, stage.CodeRNPTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVRReadingTypes, stage.CodeRVRReadingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMPointRoleTypes, stage.CodeRVSMPointRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMTypes, stage.CodeRVSMTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRadarServiceTypes, stage.CodeRadarServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRadioEmissionTypes, stage.CodeRadioEmissionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRadioFrequencyAreaTypes, stage.CodeRadioFrequencyAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRadioSignalTypes, stage.CodeRadioSignalTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeReferencePathIdentifierTypes, stage.CodeReferencePathIdentifierTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeReferenceRoleTypes, stage.CodeReferenceRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeReflectorTypes, stage.CodeReflectorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRelativePositionTypes, stage.CodeRelativePositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRoadTypes, stage.CodeRoadTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteAvailabilityTypes, stage.CodeRouteAvailabilityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteDesignatorLetterTypes, stage.CodeRouteDesignatorLetterTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteDesignatorPrefixTypes, stage.CodeRouteDesignatorPrefixTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteDesignatorSuffixTypes, stage.CodeRouteDesignatorSuffixTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteIndicatorTypes, stage.CodeRouteIndicatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteNavigationTypes, stage.CodeRouteNavigationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteOriginTypes, stage.CodeRouteOriginTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteSegmentPathTypes, stage.CodeRouteSegmentPathTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteTypes, stage.CodeRouteTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRuleProcedureTitleTypes, stage.CodeRuleProcedureTitleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRuleProcedureTypes, stage.CodeRuleProcedureTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayElementTypes, stage.CodeRunwayElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayMarkingTypes, stage.CodeRunwayMarkingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayPointRoleTypes, stage.CodeRunwayPointRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayProtectionAreaTypes, stage.CodeRunwayProtectionAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwaySectionTypes, stage.CodeRunwaySectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayTypes, stage.CodeRunwayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSafeAltitudeTypes, stage.CodeSafeAltitudeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSegmentPathTypes, stage.CodeSegmentPathTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSegmentTerminationTypes, stage.CodeSegmentTerminationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceATCTypes, stage.CodeServiceATCTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceATFMTypes, stage.CodeServiceATFMTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceGroundControlTypes, stage.CodeServiceGroundControlTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceInformationTypes, stage.CodeServiceInformationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceSARTypes, stage.CodeServiceSARTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSideTypes, stage.CodeSideTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSignalPerformanceILSTypes, stage.CodeSignalPerformanceILSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpecialDateTypes, stage.CodeSpecialDateTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpecialNavigationChainDesignatorTypes, stage.CodeSpecialNavigationChainDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpecialNavigationStationTypes, stage.CodeSpecialNavigationStationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpecialNavigationSystemTypes, stage.CodeSpecialNavigationSystemTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpeedReferenceTypes, stage.CodeSpeedReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStandbyPowerTypes, stage.CodeStandbyPowerTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusAirportTypes, stage.CodeStatusAirportTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusAirspaceTypes, stage.CodeStatusAirspaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusConstructionTypes, stage.CodeStatusConstructionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusNavaidTypes, stage.CodeStatusNavaidTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusServiceTypes, stage.CodeStatusServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSurfaceCompositionTypes, stage.CodeSurfaceCompositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSurfaceConditionTypes, stage.CodeSurfaceConditionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSurfacePreparationTypes, stage.CodeSurfacePreparationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSystemActivationTypes, stage.CodeSystemActivationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTAATypes, stage.CodeTAATypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTACANChannelTypes, stage.CodeTACANChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTLOFSectionTypes, stage.CodeTLOFSectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTaxiwayElementTypes, stage.CodeTaxiwayElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTaxiwaySectionTypes, stage.CodeTaxiwaySectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTaxiwayTypes, stage.CodeTaxiwayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTrajectoryTypes, stage.CodeTrajectoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTransponderTypes, stage.CodeTransponderTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUUIDTypes, stage.CodeUUIDTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUnitDependencyTypes, stage.CodeUnitDependencyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUnitTypes, stage.CodeUnitTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUpperAlphaTypes, stage.CodeUpperAlphaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUsageLimitationTypes, stage.CodeUsageLimitationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVASISTypes, stage.CodeVASISTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVORTypes, stage.CodeVORTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalStructureMarkingTypes, stage.CodeVerticalStructureMarkingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalStructureMaterialTypes, stage.CodeVerticalStructureMaterialTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalStructureTypes, stage.CodeVerticalStructureTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVisualDockingGuidanceTypes, stage.CodeVisualDockingGuidanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeWakeTurbulenceTypes, stage.CodeWakeTurbulenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeWorkAreaTypes, stage.CodeWorkAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeYesNoTypes, stage.CodeYesNoTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateMonthDayTypes, stage.DateMonthDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateTimeTypes, stage.DateTimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateTypes, stage.DateTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateYearTypes, stage.DateYearTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoNumberTypes, stage.NoNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoSequenceTypes, stage.NoSequenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextAddressTypes, stage.TextAddressTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextDesignatorTypes, stage.TextDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextInstructionTypes, stage.TextInstructionTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNameTypes, stage.TextNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNoteTypes, stage.TextNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPhoneTypes, stage.TextPhoneTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPropertyNameTypes, stage.TextPropertyNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextRemarkTypes, stage.TextRemarkTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextSIDSTARDesignatorTypes, stage.TextSIDSTARDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimeTypes, stage.TimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValAlarmLimitTypes, stage.ValAlarmLimitTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValAngleTypes, stage.ValAngleTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValBearingTypes, stage.ValBearingTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValChannelNumberTypes, stage.ValChannelNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDepthTypes, stage.ValDepthTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceTypes, stage.ValDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDurationTypes, stage.ValDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValFLTypes, stage.ValFLTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValFrequencyTypes, stage.ValFrequencyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValFrictionTypes, stage.ValFrictionTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValHexTypes, stage.ValHexTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValLCNTypes, stage.ValLCNTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValLightIntensityTypes, stage.ValLightIntensityTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValMagneticVariationChangeTypes, stage.ValMagneticVariationChangeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValMagneticVariationTypes, stage.ValMagneticVariationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValPCNTypes, stage.ValPCNTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValPercentTypes, stage.ValPercentTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValPressureTypes, stage.ValPressureTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValSlopeTypes, stage.ValSlopeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValSpeedTypes, stage.ValSpeedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValTemperatureTypes, stage.ValTemperatureTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValWeightTypes, stage.ValWeightTypes_reference, stage)

	__gong__computeReferencePass2(stage.XHTMLTypes, stage.XHTMLTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (codeatcreportingtype *CodeATCReportingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeATCReportingType_stagedOrder, stage.CodeATCReportingTypes_referenceOrder, codeatcreportingtype, "CodeATCReportingType")
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAerialRefuellingPointType_stagedOrder, stage.CodeAerialRefuellingPointTypes_referenceOrder, codeaerialrefuellingpointtype, "CodeAerialRefuellingPointType")
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAerialRefuellingPrefixType_stagedOrder, stage.CodeAerialRefuellingPrefixTypes_referenceOrder, codeaerialrefuellingprefixtype, "CodeAerialRefuellingPrefixType")
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAerialRefuellingType_stagedOrder, stage.CodeAerialRefuellingTypes_referenceOrder, codeaerialrefuellingtype, "CodeAerialRefuellingType")
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftCategoryType_stagedOrder, stage.CodeAircraftCategoryTypes_referenceOrder, codeaircraftcategorytype, "CodeAircraftCategoryType")
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftEngineNumberType_stagedOrder, stage.CodeAircraftEngineNumberTypes_referenceOrder, codeaircraftenginenumbertype, "CodeAircraftEngineNumberType")
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftEngineType_stagedOrder, stage.CodeAircraftEngineTypes_referenceOrder, codeaircraftenginetype, "CodeAircraftEngineType")
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftGroundServiceType_stagedOrder, stage.CodeAircraftGroundServiceTypes_referenceOrder, codeaircraftgroundservicetype, "CodeAircraftGroundServiceType")
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftICAOType_stagedOrder, stage.CodeAircraftICAOTypes_referenceOrder, codeaircrafticaotype, "CodeAircraftICAOType")
}

func (codeaircraftstandtype *CodeAircraftStandType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftStandType_stagedOrder, stage.CodeAircraftStandTypes_referenceOrder, codeaircraftstandtype, "CodeAircraftStandType")
}

func (codeaircrafttype *CodeAircraftType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftType_stagedOrder, stage.CodeAircraftTypes_referenceOrder, codeaircrafttype, "CodeAircraftType")
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftWingspanClassType_stagedOrder, stage.CodeAircraftWingspanClassTypes_referenceOrder, codeaircraftwingspanclasstype, "CodeAircraftWingspanClassType")
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirportHeliportCollocationType_stagedOrder, stage.CodeAirportHeliportCollocationTypes_referenceOrder, codeairportheliportcollocationtype, "CodeAirportHeliportCollocationType")
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirportHeliportDesignatorType_stagedOrder, stage.CodeAirportHeliportDesignatorTypes_referenceOrder, codeairportheliportdesignatortype, "CodeAirportHeliportDesignatorType")
}

func (codeairportheliporttype *CodeAirportHeliportType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirportHeliportType_stagedOrder, stage.CodeAirportHeliportTypes_referenceOrder, codeairportheliporttype, "CodeAirportHeliportType")
}

func (codeairportwarningtype *CodeAirportWarningType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirportWarningType_stagedOrder, stage.CodeAirportWarningTypes_referenceOrder, codeairportwarningtype, "CodeAirportWarningType")
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceActivityType_stagedOrder, stage.CodeAirspaceActivityTypes_referenceOrder, codeairspaceactivitytype, "CodeAirspaceActivityType")
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceAggregationType_stagedOrder, stage.CodeAirspaceAggregationTypes_referenceOrder, codeairspaceaggregationtype, "CodeAirspaceAggregationType")
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceClassificationType_stagedOrder, stage.CodeAirspaceClassificationTypes_referenceOrder, codeairspaceclassificationtype, "CodeAirspaceClassificationType")
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceDependencyType_stagedOrder, stage.CodeAirspaceDependencyTypes_referenceOrder, codeairspacedependencytype, "CodeAirspaceDependencyType")
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceDesignatorType_stagedOrder, stage.CodeAirspaceDesignatorTypes_referenceOrder, codeairspacedesignatortype, "CodeAirspaceDesignatorType")
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspacePointPositionType_stagedOrder, stage.CodeAirspacePointPositionTypes_referenceOrder, codeairspacepointpositiontype, "CodeAirspacePointPositionType")
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspacePointRoleType_stagedOrder, stage.CodeAirspacePointRoleTypes_referenceOrder, codeairspacepointroletype, "CodeAirspacePointRoleType")
}

func (codeairspacetype *CodeAirspaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspaceType_stagedOrder, stage.CodeAirspaceTypes_referenceOrder, codeairspacetype, "CodeAirspaceType")
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAltitudeAdjustmentType_stagedOrder, stage.CodeAltitudeAdjustmentTypes_referenceOrder, codealtitudeadjustmenttype, "CodeAltitudeAdjustmentType")
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAltitudeUseType_stagedOrder, stage.CodeAltitudeUseTypes_referenceOrder, codealtitudeusetype, "CodeAltitudeUseType")
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachEquipmentAdditionalType_stagedOrder, stage.CodeApproachEquipmentAdditionalTypes_referenceOrder, codeapproachequipmentadditionaltype, "CodeApproachEquipmentAdditionalType")
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachGuidanceType_stagedOrder, stage.CodeApproachGuidanceTypes_referenceOrder, codeapproachguidancetype, "CodeApproachGuidanceType")
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachLightingICAOType_stagedOrder, stage.CodeApproachLightingICAOTypes_referenceOrder, codeapproachlightingicaotype, "CodeApproachLightingICAOType")
}

func (codeapproachlightingtype *CodeApproachLightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachLightingType_stagedOrder, stage.CodeApproachLightingTypes_referenceOrder, codeapproachlightingtype, "CodeApproachLightingType")
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachPrefixType_stagedOrder, stage.CodeApproachPrefixTypes_referenceOrder, codeapproachprefixtype, "CodeApproachPrefixType")
}

func (codeapproachtype *CodeApproachType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachType_stagedOrder, stage.CodeApproachTypes_referenceOrder, codeapproachtype, "CodeApproachType")
}

func (codeapprovaltype *CodeApprovalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApprovalType_stagedOrder, stage.CodeApprovalTypes_referenceOrder, codeapprovaltype, "CodeApprovalType")
}

func (codeapronelementtype *CodeApronElementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApronElementType_stagedOrder, stage.CodeApronElementTypes_referenceOrder, codeapronelementtype, "CodeApronElementType")
}

func (codeapronsectiontype *CodeApronSectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApronSectionType_stagedOrder, stage.CodeApronSectionTypes_referenceOrder, codeapronsectiontype, "CodeApronSectionType")
}

func (codearcdirectiontype *CodeArcDirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeArcDirectionType_stagedOrder, stage.CodeArcDirectionTypes_referenceOrder, codearcdirectiontype, "CodeArcDirectionType")
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeArrestingGearEnergyAbsorbType_stagedOrder, stage.CodeArrestingGearEnergyAbsorbTypes_referenceOrder, codearrestinggearenergyabsorbtype, "CodeArrestingGearEnergyAbsorbType")
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeArrestingGearEngageDeviceType_stagedOrder, stage.CodeArrestingGearEngageDeviceTypes_referenceOrder, codearrestinggearengagedevicetype, "CodeArrestingGearEngageDeviceType")
}

func (codeauralmorsetype *CodeAuralMorseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAuralMorseType_stagedOrder, stage.CodeAuralMorseTypes_referenceOrder, codeauralmorsetype, "CodeAuralMorseType")
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAuthorityRoleType_stagedOrder, stage.CodeAuthorityRoleTypes_referenceOrder, codeauthorityroletype, "CodeAuthorityRoleType")
}

func (codeauthoritytype *CodeAuthorityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAuthorityType_stagedOrder, stage.CodeAuthorityTypes_referenceOrder, codeauthoritytype, "CodeAuthorityType")
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAviationStandardsType_stagedOrder, stage.CodeAviationStandardsTypes_referenceOrder, codeaviationstandardstype, "CodeAviationStandardsType")
}

func (codebearingtype *CodeBearingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeBearingType_stagedOrder, stage.CodeBearingTypes_referenceOrder, codebearingtype, "CodeBearingType")
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeBuoyDesignatorType_stagedOrder, stage.CodeBuoyDesignatorTypes_referenceOrder, codebuoydesignatortype, "CodeBuoyDesignatorType")
}

func (codebuoytype *CodeBuoyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeBuoyType_stagedOrder, stage.CodeBuoyTypes_referenceOrder, codebuoytype, "CodeBuoyType")
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCardinalDirectionType_stagedOrder, stage.CodeCardinalDirectionTypes_referenceOrder, codecardinaldirectiontype, "CodeCardinalDirectionType")
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCheckpointCategoryType_stagedOrder, stage.CodeCheckpointCategoryTypes_referenceOrder, codecheckpointcategorytype, "CodeCheckpointCategoryType")
}

func (codecolourtype *CodeColourType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeColourType_stagedOrder, stage.CodeColourTypes_referenceOrder, codecolourtype, "CodeColourType")
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCommunicationChannelType_stagedOrder, stage.CodeCommunicationChannelTypes_referenceOrder, codecommunicationchanneltype, "CodeCommunicationChannelType")
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCommunicationDirectionType_stagedOrder, stage.CodeCommunicationDirectionTypes_referenceOrder, codecommunicationdirectiontype, "CodeCommunicationDirectionType")
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCommunicationModeType_stagedOrder, stage.CodeCommunicationModeTypes_referenceOrder, codecommunicationmodetype, "CodeCommunicationModeType")
}

func (codecomparisontype *CodeComparisonType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeComparisonType_stagedOrder, stage.CodeComparisonTypes_referenceOrder, codecomparisontype, "CodeComparisonType")
}

func (codecontaminationtype *CodeContaminationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeContaminationType_stagedOrder, stage.CodeContaminationTypes_referenceOrder, codecontaminationtype, "CodeContaminationType")
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCourseQualityILSType_stagedOrder, stage.CodeCourseQualityILSTypes_referenceOrder, codecoursequalityilstype, "CodeCourseQualityILSType")
}

func (codecoursetype *CodeCourseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCourseType_stagedOrder, stage.CodeCourseTypes_referenceOrder, codecoursetype, "CodeCourseType")
}

func (codedmechanneltype *CodeDMEChannelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDMEChannelType_stagedOrder, stage.CodeDMEChannelTypes_referenceOrder, codedmechanneltype, "CodeDMEChannelType")
}

func (codedmetype *CodeDMEType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDMEType_stagedOrder, stage.CodeDMETypes_referenceOrder, codedmetype, "CodeDMEType")
}

func (codedaytype *CodeDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDayType_stagedOrder, stage.CodeDayTypes_referenceOrder, codedaytype, "CodeDayType")
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDeclaredDistanceType_stagedOrder, stage.CodeDeclaredDistanceTypes_referenceOrder, codedeclareddistancetype, "CodeDeclaredDistanceType")
}

func (codedesignstandardtype *CodeDesignStandardType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDesignStandardType_stagedOrder, stage.CodeDesignStandardTypes_referenceOrder, codedesignstandardtype, "CodeDesignStandardType")
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDesignatedPointDesignatorType_stagedOrder, stage.CodeDesignatedPointDesignatorTypes_referenceOrder, codedesignatedpointdesignatortype, "CodeDesignatedPointDesignatorType")
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDesignatedPointType_stagedOrder, stage.CodeDesignatedPointTypes_referenceOrder, codedesignatedpointtype, "CodeDesignatedPointType")
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDirectionReferenceType_stagedOrder, stage.CodeDirectionReferenceTypes_referenceOrder, codedirectionreferencetype, "CodeDirectionReferenceType")
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDirectionTurnType_stagedOrder, stage.CodeDirectionTurnTypes_referenceOrder, codedirectionturntype, "CodeDirectionTurnType")
}

func (codedirectiontype *CodeDirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDirectionType_stagedOrder, stage.CodeDirectionTypes_referenceOrder, codedirectiontype, "CodeDirectionType")
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDistanceIndicationType_stagedOrder, stage.CodeDistanceIndicationTypes_referenceOrder, codedistanceindicationtype, "CodeDistanceIndicationType")
}

func (codeemissionbandtype *CodeEmissionBandType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeEmissionBandType_stagedOrder, stage.CodeEmissionBandTypes_referenceOrder, codeemissionbandtype, "CodeEmissionBandType")
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeEquipmentAntiCollisionType_stagedOrder, stage.CodeEquipmentAntiCollisionTypes_referenceOrder, codeequipmentanticollisiontype, "CodeEquipmentAntiCollisionType")
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeEquipmentUnavailableType_stagedOrder, stage.CodeEquipmentUnavailableTypes_referenceOrder, codeequipmentunavailabletype, "CodeEquipmentUnavailableType")
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFacilityRankingType_stagedOrder, stage.CodeFacilityRankingTypes_referenceOrder, codefacilityrankingtype, "CodeFacilityRankingType")
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFinalGuidanceType_stagedOrder, stage.CodeFinalGuidanceTypes_referenceOrder, codefinalguidancetype, "CodeFinalGuidanceType")
}

func (codefirefightingtype *CodeFireFightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFireFightingType_stagedOrder, stage.CodeFireFightingTypes_referenceOrder, codefirefightingtype, "CodeFireFightingType")
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightDestinationType_stagedOrder, stage.CodeFlightDestinationTypes_referenceOrder, codeflightdestinationtype, "CodeFlightDestinationType")
}

func (codeflightorigintype *CodeFlightOriginType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightOriginType_stagedOrder, stage.CodeFlightOriginTypes_referenceOrder, codeflightorigintype, "CodeFlightOriginType")
}

func (codeflightpurposetype *CodeFlightPurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightPurposeType_stagedOrder, stage.CodeFlightPurposeTypes_referenceOrder, codeflightpurposetype, "CodeFlightPurposeType")
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightRestrictionDesignatorType_stagedOrder, stage.CodeFlightRestrictionDesignatorTypes_referenceOrder, codeflightrestrictiondesignatortype, "CodeFlightRestrictionDesignatorType")
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightRestrictionType_stagedOrder, stage.CodeFlightRestrictionTypes_referenceOrder, codeflightrestrictiontype, "CodeFlightRestrictionType")
}

func (codeflightruletype *CodeFlightRuleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightRuleType_stagedOrder, stage.CodeFlightRuleTypes_referenceOrder, codeflightruletype, "CodeFlightRuleType")
}

func (codeflightstatustype *CodeFlightStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightStatusType_stagedOrder, stage.CodeFlightStatusTypes_referenceOrder, codeflightstatustype, "CodeFlightStatusType")
}

func (codeflighttype *CodeFlightType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightType_stagedOrder, stage.CodeFlightTypes_referenceOrder, codeflighttype, "CodeFlightType")
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlowConditionOperationType_stagedOrder, stage.CodeFlowConditionOperationTypes_referenceOrder, codeflowconditionoperationtype, "CodeFlowConditionOperationType")
}

func (codefreeflighttype *CodeFreeFlightType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFreeFlightType_stagedOrder, stage.CodeFreeFlightTypes_referenceOrder, codefreeflighttype, "CodeFreeFlightType")
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFrictionDeviceType_stagedOrder, stage.CodeFrictionDeviceTypes_referenceOrder, codefrictiondevicetype, "CodeFrictionDeviceType")
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFrictionEstimateType_stagedOrder, stage.CodeFrictionEstimateTypes_referenceOrder, codefrictionestimatetype, "CodeFrictionEstimateType")
}

func (codefueltype *CodeFuelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFuelType_stagedOrder, stage.CodeFuelTypes_referenceOrder, codefueltype, "CodeFuelType")
}

func (codegeobordertype *CodeGeoBorderType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeGeoBorderType_stagedOrder, stage.CodeGeoBorderTypes_referenceOrder, codegeobordertype, "CodeGeoBorderType")
}

func (codegradeseparationtype *CodeGradeSeparationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeGradeSeparationType_stagedOrder, stage.CodeGradeSeparationTypes_referenceOrder, codegradeseparationtype, "CodeGradeSeparationType")
}

func (codegroundlightingtype *CodeGroundLightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeGroundLightingType_stagedOrder, stage.CodeGroundLightingTypes_referenceOrder, codegroundlightingtype, "CodeGroundLightingType")
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeGuidanceLineType_stagedOrder, stage.CodeGuidanceLineTypes_referenceOrder, codeguidancelinetype, "CodeGuidanceLineType")
}

func (codeheightreferencetype *CodeHeightReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHeightReferenceType_stagedOrder, stage.CodeHeightReferenceTypes_referenceOrder, codeheightreferencetype, "CodeHeightReferenceType")
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHelicopterPerformanceType_stagedOrder, stage.CodeHelicopterPerformanceTypes_referenceOrder, codehelicopterperformancetype, "CodeHelicopterPerformanceType")
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHoldingCategoryType_stagedOrder, stage.CodeHoldingCategoryTypes_referenceOrder, codeholdingcategorytype, "CodeHoldingCategoryType")
}

func (codeholdingusagetype *CodeHoldingUsageType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHoldingUsageType_stagedOrder, stage.CodeHoldingUsageTypes_referenceOrder, codeholdingusagetype, "CodeHoldingUsageType")
}

func (codeholdingusetype *CodeHoldingUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHoldingUseType_stagedOrder, stage.CodeHoldingUseTypes_referenceOrder, codeholdingusetype, "CodeHoldingUseType")
}

func (codeiatatype *CodeIATAType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeIATAType_stagedOrder, stage.CodeIATATypes_referenceOrder, codeiatatype, "CodeIATAType")
}

func (codeicaocountrytype *CodeICAOCountryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeICAOCountryType_stagedOrder, stage.CodeICAOCountryTypes_referenceOrder, codeicaocountrytype, "CodeICAOCountryType")
}

func (codeicaotype *CodeICAOType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeICAOType_stagedOrder, stage.CodeICAOTypes_referenceOrder, codeicaotype, "CodeICAOType")
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeILSBackCourseType_stagedOrder, stage.CodeILSBackCourseTypes_referenceOrder, codeilsbackcoursetype, "CodeILSBackCourseType")
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeIntegrityLevelILSType_stagedOrder, stage.CodeIntegrityLevelILSTypes_referenceOrder, codeintegritylevelilstype, "CodeIntegrityLevelILSType")
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeIntensityStandByType_stagedOrder, stage.CodeIntensityStandByTypes_referenceOrder, codeintensitystandbytype, "CodeIntensityStandByType")
}

func (codelanguagetype *CodeLanguageType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLanguageType_stagedOrder, stage.CodeLanguageTypes_referenceOrder, codelanguagetype, "CodeLanguageType")
}

func (codelevelseriestype *CodeLevelSeriesType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLevelSeriesType_stagedOrder, stage.CodeLevelSeriesTypes_referenceOrder, codelevelseriestype, "CodeLevelSeriesType")
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLevelTableDesignatorType_stagedOrder, stage.CodeLevelTableDesignatorTypes_referenceOrder, codeleveltabledesignatortype, "CodeLevelTableDesignatorType")
}

func (codeleveltype *CodeLevelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLevelType_stagedOrder, stage.CodeLevelTypes_referenceOrder, codeleveltype, "CodeLevelType")
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLightHoldingPositionType_stagedOrder, stage.CodeLightHoldingPositionTypes_referenceOrder, codelightholdingpositiontype, "CodeLightHoldingPositionType")
}

func (codelightintensitytype *CodeLightIntensityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLightIntensityType_stagedOrder, stage.CodeLightIntensityTypes_referenceOrder, codelightintensitytype, "CodeLightIntensityType")
}

func (codelightsourcetype *CodeLightSourceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLightSourceType_stagedOrder, stage.CodeLightSourceTypes_referenceOrder, codelightsourcetype, "CodeLightSourceType")
}

func (codelightingjartype *CodeLightingJARType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLightingJARType_stagedOrder, stage.CodeLightingJARTypes_referenceOrder, codelightingjartype, "CodeLightingJARType")
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLoadingBridgeType_stagedOrder, stage.CodeLoadingBridgeTypes_referenceOrder, codeloadingbridgetype, "CodeLoadingBridgeType")
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLocationQualifierType_stagedOrder, stage.CodeLocationQualifierTypes_referenceOrder, codelocationqualifiertype, "CodeLocationQualifierType")
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLogicalOperatorType_stagedOrder, stage.CodeLogicalOperatorTypes_referenceOrder, codelogicaloperatortype, "CodeLogicalOperatorType")
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMLSAzimuthType_stagedOrder, stage.CodeMLSAzimuthTypes_referenceOrder, codemlsazimuthtype, "CodeMLSAzimuthType")
}

func (codemlschanneltype *CodeMLSChannelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMLSChannelType_stagedOrder, stage.CodeMLSChannelTypes_referenceOrder, codemlschanneltype, "CodeMLSChannelType")
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMarkerBeaconSignalType_stagedOrder, stage.CodeMarkerBeaconSignalTypes_referenceOrder, codemarkerbeaconsignaltype, "CodeMarkerBeaconSignalType")
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMarkingConditionType_stagedOrder, stage.CodeMarkingConditionTypes_referenceOrder, codemarkingconditiontype, "CodeMarkingConditionType")
}

func (codemarkingstyletype *CodeMarkingStyleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMarkingStyleType_stagedOrder, stage.CodeMarkingStyleTypes_referenceOrder, codemarkingstyletype, "CodeMarkingStyleType")
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMeteoConditionsType_stagedOrder, stage.CodeMeteoConditionsTypes_referenceOrder, codemeteoconditionstype, "CodeMeteoConditionsType")
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMilitaryOperationsType_stagedOrder, stage.CodeMilitaryOperationsTypes_referenceOrder, codemilitaryoperationstype, "CodeMilitaryOperationsType")
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMilitaryRoutePointType_stagedOrder, stage.CodeMilitaryRoutePointTypes_referenceOrder, codemilitaryroutepointtype, "CodeMilitaryRoutePointType")
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMilitaryStatusType_stagedOrder, stage.CodeMilitaryStatusTypes_referenceOrder, codemilitarystatustype, "CodeMilitaryStatusType")
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMilitaryTrainingType_stagedOrder, stage.CodeMilitaryTrainingTypes_referenceOrder, codemilitarytrainingtype, "CodeMilitaryTrainingType")
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMinimaFinalApproachPathType_stagedOrder, stage.CodeMinimaFinalApproachPathTypes_referenceOrder, codeminimafinalapproachpathtype, "CodeMinimaFinalApproachPathType")
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMinimumAltitudeType_stagedOrder, stage.CodeMinimumAltitudeTypes_referenceOrder, codeminimumaltitudetype, "CodeMinimumAltitudeType")
}

func (codeminimumheighttype *CodeMinimumHeightType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMinimumHeightType_stagedOrder, stage.CodeMinimumHeightTypes_referenceOrder, codeminimumheighttype, "CodeMinimumHeightType")
}

func (codemissedapproachtype *CodeMissedApproachType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMissedApproachType_stagedOrder, stage.CodeMissedApproachTypes_referenceOrder, codemissedapproachtype, "CodeMissedApproachType")
}

func (codendbusagetype *CodeNDBUsageType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNDBUsageType_stagedOrder, stage.CodeNDBUsageTypes_referenceOrder, codendbusagetype, "CodeNDBUsageType")
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavaidDesignatorType_stagedOrder, stage.CodeNavaidDesignatorTypes_referenceOrder, codenavaiddesignatortype, "CodeNavaidDesignatorType")
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavaidPurposeType_stagedOrder, stage.CodeNavaidPurposeTypes_referenceOrder, codenavaidpurposetype, "CodeNavaidPurposeType")
}

func (codenavaidservicetype *CodeNavaidServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavaidServiceType_stagedOrder, stage.CodeNavaidServiceTypes_referenceOrder, codenavaidservicetype, "CodeNavaidServiceType")
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationAreaRestrictionType_stagedOrder, stage.CodeNavigationAreaRestrictionTypes_referenceOrder, codenavigationarearestrictiontype, "CodeNavigationAreaRestrictionType")
}

func (codenavigationareatype *CodeNavigationAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationAreaType_stagedOrder, stage.CodeNavigationAreaTypes_referenceOrder, codenavigationareatype, "CodeNavigationAreaType")
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationEquipmentType_stagedOrder, stage.CodeNavigationEquipmentTypes_referenceOrder, codenavigationequipmenttype, "CodeNavigationEquipmentType")
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationSpecificationType_stagedOrder, stage.CodeNavigationSpecificationTypes_referenceOrder, codenavigationspecificationtype, "CodeNavigationSpecificationType")
}

func (codenitrogentype *CodeNitrogenType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNitrogenType_stagedOrder, stage.CodeNitrogenTypes_referenceOrder, codenitrogentype, "CodeNitrogenType")
}

func (codenorthreferencetype *CodeNorthReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNorthReferenceType_stagedOrder, stage.CodeNorthReferenceTypes_referenceOrder, codenorthreferencetype, "CodeNorthReferenceType")
}

func (codenotepurposetype *CodeNotePurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeTypes_referenceOrder, codenotepurposetype, "CodeNotePurposeType")
}

func (codeobstacleareatype *CodeObstacleAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstacleAreaType_stagedOrder, stage.CodeObstacleAreaTypes_referenceOrder, codeobstacleareatype, "CodeObstacleAreaType")
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstacleAssessmentSurfaceType_stagedOrder, stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder, codeobstacleassessmentsurfacetype, "CodeObstacleAssessmentSurfaceType")
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstructionIdSurfaceZoneType_stagedOrder, stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder, codeobstructionidsurfacezonetype, "CodeObstructionIdSurfaceZoneType")
}

func (codeoiltype *CodeOilType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOilType_stagedOrder, stage.CodeOilTypes_referenceOrder, codeoiltype, "CodeOilType")
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOperationAirportHeliportType_stagedOrder, stage.CodeOperationAirportHeliportTypes_referenceOrder, codeoperationairportheliporttype, "CodeOperationAirportHeliportType")
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOperationManoeuvringAreaType_stagedOrder, stage.CodeOperationManoeuvringAreaTypes_referenceOrder, codeoperationmanoeuvringareatype, "CodeOperationManoeuvringAreaType")
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOrganisationDesignatorType_stagedOrder, stage.CodeOrganisationDesignatorTypes_referenceOrder, codeorganisationdesignatortype, "CodeOrganisationDesignatorType")
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOrganisationHierarchyType_stagedOrder, stage.CodeOrganisationHierarchyTypes_referenceOrder, codeorganisationhierarchytype, "CodeOrganisationHierarchyType")
}

func (codeorganisationtype *CodeOrganisationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOrganisationType_stagedOrder, stage.CodeOrganisationTypes_referenceOrder, codeorganisationtype, "CodeOrganisationType")
}

func (codeoxygentype *CodeOxygenType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOxygenType_stagedOrder, stage.CodeOxygenTypes_referenceOrder, codeoxygentype, "CodeOxygenType")
}

func (codepartype *CodePARType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePARType_stagedOrder, stage.CodePARTypes_referenceOrder, codepartype, "CodePARType")
}

func (codepcnmethodtype *CodePCNMethodType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePCNMethodType_stagedOrder, stage.CodePCNMethodTypes_referenceOrder, codepcnmethodtype, "CodePCNMethodType")
}

func (codepcnpavementtype *CodePCNPavementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePCNPavementType_stagedOrder, stage.CodePCNPavementTypes_referenceOrder, codepcnpavementtype, "CodePCNPavementType")
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePCNSubgradeType_stagedOrder, stage.CodePCNSubgradeTypes_referenceOrder, codepcnsubgradetype, "CodePCNSubgradeType")
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePCNTyrePressureType_stagedOrder, stage.CodePCNTyrePressureTypes_referenceOrder, codepcntyrepressuretype, "CodePCNTyrePressureType")
}

func (codepassengerservicetype *CodePassengerServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePassengerServiceType_stagedOrder, stage.CodePassengerServiceTypes_referenceOrder, codepassengerservicetype, "CodePassengerServiceType")
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePilotControlledLightingType_stagedOrder, stage.CodePilotControlledLightingTypes_referenceOrder, codepilotcontrolledlightingtype, "CodePilotControlledLightingType")
}

func (codepositioninilstype *CodePositionInILSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePositionInILSType_stagedOrder, stage.CodePositionInILSTypes_referenceOrder, codepositioninilstype, "CodePositionInILSType")
}

func (codeprimaryradartype *CodePrimaryRadarType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePrimaryRadarType_stagedOrder, stage.CodePrimaryRadarTypes_referenceOrder, codeprimaryradartype, "CodePrimaryRadarType")
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeProcedureAvailabilityType_stagedOrder, stage.CodeProcedureAvailabilityTypes_referenceOrder, codeprocedureavailabilitytype, "CodeProcedureAvailabilityType")
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeProcedureCodingStandardType_stagedOrder, stage.CodeProcedureCodingStandardTypes_referenceOrder, codeprocedurecodingstandardtype, "CodeProcedureCodingStandardType")
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeProcedureDistanceType_stagedOrder, stage.CodeProcedureDistanceTypes_referenceOrder, codeproceduredistancetype, "CodeProcedureDistanceType")
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeProcedureFixRoleType_stagedOrder, stage.CodeProcedureFixRoleTypes_referenceOrder, codeprocedurefixroletype, "CodeProcedureFixRoleType")
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeProcedurePhaseType_stagedOrder, stage.CodeProcedurePhaseTypes_referenceOrder, codeprocedurephasetype, "CodeProcedurePhaseType")
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeProtectAreaSectionType_stagedOrder, stage.CodeProtectAreaSectionTypes_referenceOrder, codeprotectareasectiontype, "CodeProtectAreaSectionType")
}

func (codernptype *CodeRNPType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRNPType_stagedOrder, stage.CodeRNPTypes_referenceOrder, codernptype, "CodeRNPType")
}

func (codervrreadingtype *CodeRVRReadingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVRReadingType_stagedOrder, stage.CodeRVRReadingTypes_referenceOrder, codervrreadingtype, "CodeRVRReadingType")
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMPointRoleType_stagedOrder, stage.CodeRVSMPointRoleTypes_referenceOrder, codervsmpointroletype, "CodeRVSMPointRoleType")
}

func (codervsmtype *CodeRVSMType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMType_stagedOrder, stage.CodeRVSMTypes_referenceOrder, codervsmtype, "CodeRVSMType")
}

func (coderadarservicetype *CodeRadarServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRadarServiceType_stagedOrder, stage.CodeRadarServiceTypes_referenceOrder, coderadarservicetype, "CodeRadarServiceType")
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRadioEmissionType_stagedOrder, stage.CodeRadioEmissionTypes_referenceOrder, coderadioemissiontype, "CodeRadioEmissionType")
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRadioFrequencyAreaType_stagedOrder, stage.CodeRadioFrequencyAreaTypes_referenceOrder, coderadiofrequencyareatype, "CodeRadioFrequencyAreaType")
}

func (coderadiosignaltype *CodeRadioSignalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRadioSignalType_stagedOrder, stage.CodeRadioSignalTypes_referenceOrder, coderadiosignaltype, "CodeRadioSignalType")
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeReferencePathIdentifierType_stagedOrder, stage.CodeReferencePathIdentifierTypes_referenceOrder, codereferencepathidentifiertype, "CodeReferencePathIdentifierType")
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeReferenceRoleType_stagedOrder, stage.CodeReferenceRoleTypes_referenceOrder, codereferenceroletype, "CodeReferenceRoleType")
}

func (codereflectortype *CodeReflectorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeReflectorType_stagedOrder, stage.CodeReflectorTypes_referenceOrder, codereflectortype, "CodeReflectorType")
}

func (coderelativepositiontype *CodeRelativePositionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRelativePositionType_stagedOrder, stage.CodeRelativePositionTypes_referenceOrder, coderelativepositiontype, "CodeRelativePositionType")
}

func (coderoadtype *CodeRoadType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRoadType_stagedOrder, stage.CodeRoadTypes_referenceOrder, coderoadtype, "CodeRoadType")
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteAvailabilityType_stagedOrder, stage.CodeRouteAvailabilityTypes_referenceOrder, coderouteavailabilitytype, "CodeRouteAvailabilityType")
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteDesignatorLetterType_stagedOrder, stage.CodeRouteDesignatorLetterTypes_referenceOrder, coderoutedesignatorlettertype, "CodeRouteDesignatorLetterType")
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteDesignatorPrefixType_stagedOrder, stage.CodeRouteDesignatorPrefixTypes_referenceOrder, coderoutedesignatorprefixtype, "CodeRouteDesignatorPrefixType")
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteDesignatorSuffixType_stagedOrder, stage.CodeRouteDesignatorSuffixTypes_referenceOrder, coderoutedesignatorsuffixtype, "CodeRouteDesignatorSuffixType")
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteIndicatorType_stagedOrder, stage.CodeRouteIndicatorTypes_referenceOrder, coderouteindicatortype, "CodeRouteIndicatorType")
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteNavigationType_stagedOrder, stage.CodeRouteNavigationTypes_referenceOrder, coderoutenavigationtype, "CodeRouteNavigationType")
}

func (coderouteorigintype *CodeRouteOriginType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteOriginType_stagedOrder, stage.CodeRouteOriginTypes_referenceOrder, coderouteorigintype, "CodeRouteOriginType")
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteSegmentPathType_stagedOrder, stage.CodeRouteSegmentPathTypes_referenceOrder, coderoutesegmentpathtype, "CodeRouteSegmentPathType")
}

func (coderoutetype *CodeRouteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteType_stagedOrder, stage.CodeRouteTypes_referenceOrder, coderoutetype, "CodeRouteType")
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRuleProcedureTitleType_stagedOrder, stage.CodeRuleProcedureTitleTypes_referenceOrder, coderuleproceduretitletype, "CodeRuleProcedureTitleType")
}

func (coderuleproceduretype *CodeRuleProcedureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRuleProcedureType_stagedOrder, stage.CodeRuleProcedureTypes_referenceOrder, coderuleproceduretype, "CodeRuleProcedureType")
}

func (coderunwayelementtype *CodeRunwayElementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRunwayElementType_stagedOrder, stage.CodeRunwayElementTypes_referenceOrder, coderunwayelementtype, "CodeRunwayElementType")
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRunwayMarkingType_stagedOrder, stage.CodeRunwayMarkingTypes_referenceOrder, coderunwaymarkingtype, "CodeRunwayMarkingType")
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRunwayPointRoleType_stagedOrder, stage.CodeRunwayPointRoleTypes_referenceOrder, coderunwaypointroletype, "CodeRunwayPointRoleType")
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRunwayProtectionAreaType_stagedOrder, stage.CodeRunwayProtectionAreaTypes_referenceOrder, coderunwayprotectionareatype, "CodeRunwayProtectionAreaType")
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRunwaySectionType_stagedOrder, stage.CodeRunwaySectionTypes_referenceOrder, coderunwaysectiontype, "CodeRunwaySectionType")
}

func (coderunwaytype *CodeRunwayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRunwayType_stagedOrder, stage.CodeRunwayTypes_referenceOrder, coderunwaytype, "CodeRunwayType")
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSafeAltitudeType_stagedOrder, stage.CodeSafeAltitudeTypes_referenceOrder, codesafealtitudetype, "CodeSafeAltitudeType")
}

func (codesegmentpathtype *CodeSegmentPathType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSegmentPathType_stagedOrder, stage.CodeSegmentPathTypes_referenceOrder, codesegmentpathtype, "CodeSegmentPathType")
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSegmentTerminationType_stagedOrder, stage.CodeSegmentTerminationTypes_referenceOrder, codesegmentterminationtype, "CodeSegmentTerminationType")
}

func (codeserviceatctype *CodeServiceATCType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeServiceATCType_stagedOrder, stage.CodeServiceATCTypes_referenceOrder, codeserviceatctype, "CodeServiceATCType")
}

func (codeserviceatfmtype *CodeServiceATFMType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeServiceATFMType_stagedOrder, stage.CodeServiceATFMTypes_referenceOrder, codeserviceatfmtype, "CodeServiceATFMType")
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeServiceGroundControlType_stagedOrder, stage.CodeServiceGroundControlTypes_referenceOrder, codeservicegroundcontroltype, "CodeServiceGroundControlType")
}

func (codeserviceinformationtype *CodeServiceInformationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeServiceInformationType_stagedOrder, stage.CodeServiceInformationTypes_referenceOrder, codeserviceinformationtype, "CodeServiceInformationType")
}

func (codeservicesartype *CodeServiceSARType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeServiceSARType_stagedOrder, stage.CodeServiceSARTypes_referenceOrder, codeservicesartype, "CodeServiceSARType")
}

func (codesidetype *CodeSideType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSideType_stagedOrder, stage.CodeSideTypes_referenceOrder, codesidetype, "CodeSideType")
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSignalPerformanceILSType_stagedOrder, stage.CodeSignalPerformanceILSTypes_referenceOrder, codesignalperformanceilstype, "CodeSignalPerformanceILSType")
}

func (codespecialdatetype *CodeSpecialDateType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSpecialDateType_stagedOrder, stage.CodeSpecialDateTypes_referenceOrder, codespecialdatetype, "CodeSpecialDateType")
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSpecialNavigationChainDesignatorType_stagedOrder, stage.CodeSpecialNavigationChainDesignatorTypes_referenceOrder, codespecialnavigationchaindesignatortype, "CodeSpecialNavigationChainDesignatorType")
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSpecialNavigationStationType_stagedOrder, stage.CodeSpecialNavigationStationTypes_referenceOrder, codespecialnavigationstationtype, "CodeSpecialNavigationStationType")
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSpecialNavigationSystemType_stagedOrder, stage.CodeSpecialNavigationSystemTypes_referenceOrder, codespecialnavigationsystemtype, "CodeSpecialNavigationSystemType")
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSpeedReferenceType_stagedOrder, stage.CodeSpeedReferenceTypes_referenceOrder, codespeedreferencetype, "CodeSpeedReferenceType")
}

func (codestandbypowertype *CodeStandbyPowerType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStandbyPowerType_stagedOrder, stage.CodeStandbyPowerTypes_referenceOrder, codestandbypowertype, "CodeStandbyPowerType")
}

func (codestatusairporttype *CodeStatusAirportType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusAirportType_stagedOrder, stage.CodeStatusAirportTypes_referenceOrder, codestatusairporttype, "CodeStatusAirportType")
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusAirspaceType_stagedOrder, stage.CodeStatusAirspaceTypes_referenceOrder, codestatusairspacetype, "CodeStatusAirspaceType")
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusConstructionType_stagedOrder, stage.CodeStatusConstructionTypes_referenceOrder, codestatusconstructiontype, "CodeStatusConstructionType")
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusNavaidType_stagedOrder, stage.CodeStatusNavaidTypes_referenceOrder, codestatusnavaidtype, "CodeStatusNavaidType")
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusOperationsType_stagedOrder, stage.CodeStatusOperationsTypes_referenceOrder, codestatusoperationstype, "CodeStatusOperationsType")
}

func (codestatusservicetype *CodeStatusServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusServiceType_stagedOrder, stage.CodeStatusServiceTypes_referenceOrder, codestatusservicetype, "CodeStatusServiceType")
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSurfaceCompositionType_stagedOrder, stage.CodeSurfaceCompositionTypes_referenceOrder, codesurfacecompositiontype, "CodeSurfaceCompositionType")
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSurfaceConditionType_stagedOrder, stage.CodeSurfaceConditionTypes_referenceOrder, codesurfaceconditiontype, "CodeSurfaceConditionType")
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSurfacePreparationType_stagedOrder, stage.CodeSurfacePreparationTypes_referenceOrder, codesurfacepreparationtype, "CodeSurfacePreparationType")
}

func (codesystemactivationtype *CodeSystemActivationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSystemActivationType_stagedOrder, stage.CodeSystemActivationTypes_referenceOrder, codesystemactivationtype, "CodeSystemActivationType")
}

func (codetaatype *CodeTAAType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTAAType_stagedOrder, stage.CodeTAATypes_referenceOrder, codetaatype, "CodeTAAType")
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTACANChannelType_stagedOrder, stage.CodeTACANChannelTypes_referenceOrder, codetacanchanneltype, "CodeTACANChannelType")
}

func (codetlofsectiontype *CodeTLOFSectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTLOFSectionType_stagedOrder, stage.CodeTLOFSectionTypes_referenceOrder, codetlofsectiontype, "CodeTLOFSectionType")
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTaxiwayElementType_stagedOrder, stage.CodeTaxiwayElementTypes_referenceOrder, codetaxiwayelementtype, "CodeTaxiwayElementType")
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTaxiwaySectionType_stagedOrder, stage.CodeTaxiwaySectionTypes_referenceOrder, codetaxiwaysectiontype, "CodeTaxiwaySectionType")
}

func (codetaxiwaytype *CodeTaxiwayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTaxiwayType_stagedOrder, stage.CodeTaxiwayTypes_referenceOrder, codetaxiwaytype, "CodeTaxiwayType")
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTelecomNetworkType_stagedOrder, stage.CodeTelecomNetworkTypes_referenceOrder, codetelecomnetworktype, "CodeTelecomNetworkType")
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTimeEventCombinationType_stagedOrder, stage.CodeTimeEventCombinationTypes_referenceOrder, codetimeeventcombinationtype, "CodeTimeEventCombinationType")
}

func (codetimeeventtype *CodeTimeEventType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTimeEventType_stagedOrder, stage.CodeTimeEventTypes_referenceOrder, codetimeeventtype, "CodeTimeEventType")
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTimeReferenceType_stagedOrder, stage.CodeTimeReferenceTypes_referenceOrder, codetimereferencetype, "CodeTimeReferenceType")
}

func (codetrajectorytype *CodeTrajectoryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTrajectoryType_stagedOrder, stage.CodeTrajectoryTypes_referenceOrder, codetrajectorytype, "CodeTrajectoryType")
}

func (codetranspondertype *CodeTransponderType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTransponderType_stagedOrder, stage.CodeTransponderTypes_referenceOrder, codetranspondertype, "CodeTransponderType")
}

func (codeuuidtype *CodeUUIDType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUUIDType_stagedOrder, stage.CodeUUIDTypes_referenceOrder, codeuuidtype, "CodeUUIDType")
}

func (codeunitdependencytype *CodeUnitDependencyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUnitDependencyType_stagedOrder, stage.CodeUnitDependencyTypes_referenceOrder, codeunitdependencytype, "CodeUnitDependencyType")
}

func (codeunittype *CodeUnitType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUnitType_stagedOrder, stage.CodeUnitTypes_referenceOrder, codeunittype, "CodeUnitType")
}

func (codeupperalphatype *CodeUpperAlphaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUpperAlphaType_stagedOrder, stage.CodeUpperAlphaTypes_referenceOrder, codeupperalphatype, "CodeUpperAlphaType")
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUsageLimitationType_stagedOrder, stage.CodeUsageLimitationTypes_referenceOrder, codeusagelimitationtype, "CodeUsageLimitationType")
}

func (codevasistype *CodeVASISType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVASISType_stagedOrder, stage.CodeVASISTypes_referenceOrder, codevasistype, "CodeVASISType")
}

func (codevortype *CodeVORType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVORType_stagedOrder, stage.CodeVORTypes_referenceOrder, codevortype, "CodeVORType")
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeValueInterpretationType_stagedOrder, stage.CodeValueInterpretationTypes_referenceOrder, codevalueinterpretationtype, "CodeValueInterpretationType")
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalDatumType_stagedOrder, stage.CodeVerticalDatumTypes_referenceOrder, codeverticaldatumtype, "CodeVerticalDatumType")
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalReferenceType_stagedOrder, stage.CodeVerticalReferenceTypes_referenceOrder, codeverticalreferencetype, "CodeVerticalReferenceType")
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalStructureMarkingType_stagedOrder, stage.CodeVerticalStructureMarkingTypes_referenceOrder, codeverticalstructuremarkingtype, "CodeVerticalStructureMarkingType")
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalStructureMaterialType_stagedOrder, stage.CodeVerticalStructureMaterialTypes_referenceOrder, codeverticalstructurematerialtype, "CodeVerticalStructureMaterialType")
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalStructureType_stagedOrder, stage.CodeVerticalStructureTypes_referenceOrder, codeverticalstructuretype, "CodeVerticalStructureType")
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVisualDockingGuidanceType_stagedOrder, stage.CodeVisualDockingGuidanceTypes_referenceOrder, codevisualdockingguidancetype, "CodeVisualDockingGuidanceType")
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeWakeTurbulenceType_stagedOrder, stage.CodeWakeTurbulenceTypes_referenceOrder, codewaketurbulencetype, "CodeWakeTurbulenceType")
}

func (codeworkareatype *CodeWorkAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeWorkAreaType_stagedOrder, stage.CodeWorkAreaTypes_referenceOrder, codeworkareatype, "CodeWorkAreaType")
}

func (codeyesnotype *CodeYesNoType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeYesNoType_stagedOrder, stage.CodeYesNoTypes_referenceOrder, codeyesnotype, "CodeYesNoType")
}

func (datemonthdaytype *DateMonthDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateMonthDayType_stagedOrder, stage.DateMonthDayTypes_referenceOrder, datemonthdaytype, "DateMonthDayType")
}

func (datetimetype *DateTimeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateTimeType_stagedOrder, stage.DateTimeTypes_referenceOrder, datetimetype, "DateTimeType")
}

func (datetype *DateType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateType_stagedOrder, stage.DateTypes_referenceOrder, datetype, "DateType")
}

func (dateyeartype *DateYearType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateYearType_stagedOrder, stage.DateYearTypes_referenceOrder, dateyeartype, "DateYearType")
}

func (nonumbertype *NoNumberType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoNumberType_stagedOrder, stage.NoNumberTypes_referenceOrder, nonumbertype, "NoNumberType")
}

func (nosequencetype *NoSequenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoSequenceType_stagedOrder, stage.NoSequenceTypes_referenceOrder, nosequencetype, "NoSequenceType")
}

func (textaddresstype *TextAddressType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextAddressType_stagedOrder, stage.TextAddressTypes_referenceOrder, textaddresstype, "TextAddressType")
}

func (textdesignatortype *TextDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextDesignatorType_stagedOrder, stage.TextDesignatorTypes_referenceOrder, textdesignatortype, "TextDesignatorType")
}

func (textinstructiontype *TextInstructionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextInstructionType_stagedOrder, stage.TextInstructionTypes_referenceOrder, textinstructiontype, "TextInstructionType")
}

func (textnametype *TextNameType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextNameType_stagedOrder, stage.TextNameTypes_referenceOrder, textnametype, "TextNameType")
}

func (textnotetype *TextNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextNoteType_stagedOrder, stage.TextNoteTypes_referenceOrder, textnotetype, "TextNoteType")
}

func (textphonetype *TextPhoneType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextPhoneType_stagedOrder, stage.TextPhoneTypes_referenceOrder, textphonetype, "TextPhoneType")
}

func (textpropertynametype *TextPropertyNameType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextPropertyNameType_stagedOrder, stage.TextPropertyNameTypes_referenceOrder, textpropertynametype, "TextPropertyNameType")
}

func (textremarktype *TextRemarkType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextRemarkType_stagedOrder, stage.TextRemarkTypes_referenceOrder, textremarktype, "TextRemarkType")
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextSIDSTARDesignatorType_stagedOrder, stage.TextSIDSTARDesignatorTypes_referenceOrder, textsidstardesignatortype, "TextSIDSTARDesignatorType")
}

func (timetype *TimeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimeType_stagedOrder, stage.TimeTypes_referenceOrder, timetype, "TimeType")
}

func (valalarmlimittype *ValAlarmLimitType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValAlarmLimitType_stagedOrder, stage.ValAlarmLimitTypes_referenceOrder, valalarmlimittype, "ValAlarmLimitType")
}

func (valangletype *ValAngleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValAngleType_stagedOrder, stage.ValAngleTypes_referenceOrder, valangletype, "ValAngleType")
}

func (valbearingtype *ValBearingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValBearingType_stagedOrder, stage.ValBearingTypes_referenceOrder, valbearingtype, "ValBearingType")
}

func (valchannelnumbertype *ValChannelNumberType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValChannelNumberType_stagedOrder, stage.ValChannelNumberTypes_referenceOrder, valchannelnumbertype, "ValChannelNumberType")
}

func (valdepthtype *ValDepthType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDepthType_stagedOrder, stage.ValDepthTypes_referenceOrder, valdepthtype, "ValDepthType")
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDistanceSignedType_stagedOrder, stage.ValDistanceSignedTypes_referenceOrder, valdistancesignedtype, "ValDistanceSignedType")
}

func (valdistancetype *ValDistanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDistanceType_stagedOrder, stage.ValDistanceTypes_referenceOrder, valdistancetype, "ValDistanceType")
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDistanceVerticalType_stagedOrder, stage.ValDistanceVerticalTypes_referenceOrder, valdistanceverticaltype, "ValDistanceVerticalType")
}

func (valdurationtype *ValDurationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValDurationType_stagedOrder, stage.ValDurationTypes_referenceOrder, valdurationtype, "ValDurationType")
}

func (valfltype *ValFLType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValFLType_stagedOrder, stage.ValFLTypes_referenceOrder, valfltype, "ValFLType")
}

func (valfrequencytype *ValFrequencyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValFrequencyType_stagedOrder, stage.ValFrequencyTypes_referenceOrder, valfrequencytype, "ValFrequencyType")
}

func (valfrictiontype *ValFrictionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValFrictionType_stagedOrder, stage.ValFrictionTypes_referenceOrder, valfrictiontype, "ValFrictionType")
}

func (valhextype *ValHexType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValHexType_stagedOrder, stage.ValHexTypes_referenceOrder, valhextype, "ValHexType")
}

func (vallcntype *ValLCNType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValLCNType_stagedOrder, stage.ValLCNTypes_referenceOrder, vallcntype, "ValLCNType")
}

func (vallightintensitytype *ValLightIntensityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValLightIntensityType_stagedOrder, stage.ValLightIntensityTypes_referenceOrder, vallightintensitytype, "ValLightIntensityType")
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValMagneticVariationChangeType_stagedOrder, stage.ValMagneticVariationChangeTypes_referenceOrder, valmagneticvariationchangetype, "ValMagneticVariationChangeType")
}

func (valmagneticvariationtype *ValMagneticVariationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValMagneticVariationType_stagedOrder, stage.ValMagneticVariationTypes_referenceOrder, valmagneticvariationtype, "ValMagneticVariationType")
}

func (valpcntype *ValPCNType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValPCNType_stagedOrder, stage.ValPCNTypes_referenceOrder, valpcntype, "ValPCNType")
}

func (valpercenttype *ValPercentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValPercentType_stagedOrder, stage.ValPercentTypes_referenceOrder, valpercenttype, "ValPercentType")
}

func (valpressuretype *ValPressureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValPressureType_stagedOrder, stage.ValPressureTypes_referenceOrder, valpressuretype, "ValPressureType")
}

func (valslopetype *ValSlopeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValSlopeType_stagedOrder, stage.ValSlopeTypes_referenceOrder, valslopetype, "ValSlopeType")
}

func (valspeedtype *ValSpeedType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValSpeedType_stagedOrder, stage.ValSpeedTypes_referenceOrder, valspeedtype, "ValSpeedType")
}

func (valtemperaturetype *ValTemperatureType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValTemperatureType_stagedOrder, stage.ValTemperatureTypes_referenceOrder, valtemperaturetype, "ValTemperatureType")
}

func (valweighttype *ValWeightType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValWeightType_stagedOrder, stage.ValWeightTypes_referenceOrder, valweighttype, "ValWeightType")
}

func (xhtmltype *XHTMLType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.XHTMLType_stagedOrder, stage.XHTMLTypes_referenceOrder, xhtmltype, "XHTMLType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (codeatcreportingtype *CodeATCReportingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeatcreportingtype, codeatcreportingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeatcreportingtype *CodeATCReportingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeatcreportingtype.GongGetIdentifier(stage)
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaerialrefuellingpointtype, codeaerialrefuellingpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaerialrefuellingpointtype.GongGetIdentifier(stage)
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaerialrefuellingprefixtype, codeaerialrefuellingprefixtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaerialrefuellingprefixtype.GongGetIdentifier(stage)
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaerialrefuellingtype, codeaerialrefuellingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaerialrefuellingtype.GongGetIdentifier(stage)
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftcategorytype, codeaircraftcategorytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftcategorytype.GongGetIdentifier(stage)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftenginenumbertype, codeaircraftenginenumbertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftenginenumbertype.GongGetIdentifier(stage)
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftenginetype, codeaircraftenginetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftenginetype *CodeAircraftEngineType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftenginetype.GongGetIdentifier(stage)
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftgroundservicetype, codeaircraftgroundservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftgroundservicetype.GongGetIdentifier(stage)
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircrafticaotype, codeaircrafticaotype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircrafticaotype *CodeAircraftICAOType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircrafticaotype.GongGetIdentifier(stage)
}

func (codeaircraftstandtype *CodeAircraftStandType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftstandtype, codeaircraftstandtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftstandtype *CodeAircraftStandType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftstandtype.GongGetIdentifier(stage)
}

func (codeaircrafttype *CodeAircraftType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircrafttype, codeaircrafttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircrafttype *CodeAircraftType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircrafttype.GongGetIdentifier(stage)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftwingspanclasstype, codeaircraftwingspanclasstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftwingspanclasstype.GongGetIdentifier(stage)
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairportheliportcollocationtype, codeairportheliportcollocationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairportheliportcollocationtype.GongGetIdentifier(stage)
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairportheliportdesignatortype, codeairportheliportdesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairportheliportdesignatortype.GongGetIdentifier(stage)
}

func (codeairportheliporttype *CodeAirportHeliportType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairportheliporttype, codeairportheliporttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairportheliporttype *CodeAirportHeliportType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairportheliporttype.GongGetIdentifier(stage)
}

func (codeairportwarningtype *CodeAirportWarningType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairportwarningtype, codeairportwarningtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairportwarningtype *CodeAirportWarningType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairportwarningtype.GongGetIdentifier(stage)
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspaceactivitytype, codeairspaceactivitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspaceactivitytype *CodeAirspaceActivityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspaceactivitytype.GongGetIdentifier(stage)
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspaceaggregationtype, codeairspaceaggregationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspaceaggregationtype.GongGetIdentifier(stage)
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspaceclassificationtype, codeairspaceclassificationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspaceclassificationtype.GongGetIdentifier(stage)
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspacedependencytype, codeairspacedependencytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspacedependencytype *CodeAirspaceDependencyType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspacedependencytype.GongGetIdentifier(stage)
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspacedesignatortype, codeairspacedesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspacedesignatortype.GongGetIdentifier(stage)
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspacepointpositiontype, codeairspacepointpositiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspacepointpositiontype.GongGetIdentifier(stage)
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspacepointroletype, codeairspacepointroletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspacepointroletype.GongGetIdentifier(stage)
}

func (codeairspacetype *CodeAirspaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeairspacetype, codeairspacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeairspacetype *CodeAirspaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeairspacetype.GongGetIdentifier(stage)
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codealtitudeadjustmenttype, codealtitudeadjustmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return codealtitudeadjustmenttype.GongGetIdentifier(stage)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codealtitudeusetype, codealtitudeusetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codealtitudeusetype *CodeAltitudeUseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codealtitudeusetype.GongGetIdentifier(stage)
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapproachequipmentadditionaltype, codeapproachequipmentadditionaltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapproachequipmentadditionaltype.GongGetIdentifier(stage)
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapproachguidancetype, codeapproachguidancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapproachguidancetype.GongGetIdentifier(stage)
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapproachlightingicaotype, codeapproachlightingicaotype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapproachlightingicaotype.GongGetIdentifier(stage)
}

func (codeapproachlightingtype *CodeApproachLightingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapproachlightingtype, codeapproachlightingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapproachlightingtype *CodeApproachLightingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapproachlightingtype.GongGetIdentifier(stage)
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapproachprefixtype, codeapproachprefixtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapproachprefixtype *CodeApproachPrefixType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapproachprefixtype.GongGetIdentifier(stage)
}

func (codeapproachtype *CodeApproachType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapproachtype, codeapproachtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapproachtype *CodeApproachType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapproachtype.GongGetIdentifier(stage)
}

func (codeapprovaltype *CodeApprovalType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapprovaltype, codeapprovaltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapprovaltype *CodeApprovalType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapprovaltype.GongGetIdentifier(stage)
}

func (codeapronelementtype *CodeApronElementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapronelementtype, codeapronelementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapronelementtype *CodeApronElementType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapronelementtype.GongGetIdentifier(stage)
}

func (codeapronsectiontype *CodeApronSectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeapronsectiontype, codeapronsectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeapronsectiontype *CodeApronSectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeapronsectiontype.GongGetIdentifier(stage)
}

func (codearcdirectiontype *CodeArcDirectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codearcdirectiontype, codearcdirectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codearcdirectiontype *CodeArcDirectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codearcdirectiontype.GongGetIdentifier(stage)
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codearrestinggearenergyabsorbtype, codearrestinggearenergyabsorbtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetReferenceIdentifier(stage *Stage) string {
	return codearrestinggearenergyabsorbtype.GongGetIdentifier(stage)
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codearrestinggearengagedevicetype, codearrestinggearengagedevicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codearrestinggearengagedevicetype.GongGetIdentifier(stage)
}

func (codeauralmorsetype *CodeAuralMorseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeauralmorsetype, codeauralmorsetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeauralmorsetype *CodeAuralMorseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeauralmorsetype.GongGetIdentifier(stage)
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeauthorityroletype, codeauthorityroletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeauthorityroletype *CodeAuthorityRoleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeauthorityroletype.GongGetIdentifier(stage)
}

func (codeauthoritytype *CodeAuthorityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeauthoritytype, codeauthoritytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeauthoritytype *CodeAuthorityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeauthoritytype.GongGetIdentifier(stage)
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaviationstandardstype, codeaviationstandardstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaviationstandardstype *CodeAviationStandardsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaviationstandardstype.GongGetIdentifier(stage)
}

func (codebearingtype *CodeBearingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codebearingtype, codebearingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codebearingtype *CodeBearingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codebearingtype.GongGetIdentifier(stage)
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codebuoydesignatortype, codebuoydesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codebuoydesignatortype.GongGetIdentifier(stage)
}

func (codebuoytype *CodeBuoyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codebuoytype, codebuoytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codebuoytype *CodeBuoyType) GongGetReferenceIdentifier(stage *Stage) string {
	return codebuoytype.GongGetIdentifier(stage)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecardinaldirectiontype, codecardinaldirectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecardinaldirectiontype.GongGetIdentifier(stage)
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecheckpointcategorytype, codecheckpointcategorytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecheckpointcategorytype.GongGetIdentifier(stage)
}

func (codecolourtype *CodeColourType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecolourtype, codecolourtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecolourtype *CodeColourType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecolourtype.GongGetIdentifier(stage)
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecommunicationchanneltype, codecommunicationchanneltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecommunicationchanneltype *CodeCommunicationChannelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecommunicationchanneltype.GongGetIdentifier(stage)
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecommunicationdirectiontype, codecommunicationdirectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecommunicationdirectiontype.GongGetIdentifier(stage)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecommunicationmodetype, codecommunicationmodetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecommunicationmodetype *CodeCommunicationModeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecommunicationmodetype.GongGetIdentifier(stage)
}

func (codecomparisontype *CodeComparisonType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecomparisontype, codecomparisontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecomparisontype *CodeComparisonType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecomparisontype.GongGetIdentifier(stage)
}

func (codecontaminationtype *CodeContaminationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecontaminationtype, codecontaminationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecontaminationtype *CodeContaminationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecontaminationtype.GongGetIdentifier(stage)
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecoursequalityilstype, codecoursequalityilstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecoursequalityilstype.GongGetIdentifier(stage)
}

func (codecoursetype *CodeCourseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecoursetype, codecoursetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecoursetype *CodeCourseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecoursetype.GongGetIdentifier(stage)
}

func (codedmechanneltype *CodeDMEChannelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedmechanneltype, codedmechanneltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedmechanneltype *CodeDMEChannelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedmechanneltype.GongGetIdentifier(stage)
}

func (codedmetype *CodeDMEType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedmetype, codedmetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedmetype *CodeDMEType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedmetype.GongGetIdentifier(stage)
}

func (codedaytype *CodeDayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedaytype, codedaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedaytype *CodeDayType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedaytype.GongGetIdentifier(stage)
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedeclareddistancetype, codedeclareddistancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedeclareddistancetype.GongGetIdentifier(stage)
}

func (codedesignstandardtype *CodeDesignStandardType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedesignstandardtype, codedesignstandardtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedesignstandardtype *CodeDesignStandardType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedesignstandardtype.GongGetIdentifier(stage)
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedesignatedpointdesignatortype, codedesignatedpointdesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedesignatedpointdesignatortype.GongGetIdentifier(stage)
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedesignatedpointtype, codedesignatedpointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedesignatedpointtype *CodeDesignatedPointType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedesignatedpointtype.GongGetIdentifier(stage)
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedirectionreferencetype, codedirectionreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedirectionreferencetype.GongGetIdentifier(stage)
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedirectionturntype, codedirectionturntype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedirectionturntype *CodeDirectionTurnType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedirectionturntype.GongGetIdentifier(stage)
}

func (codedirectiontype *CodeDirectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedirectiontype, codedirectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedirectiontype *CodeDirectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedirectiontype.GongGetIdentifier(stage)
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedistanceindicationtype, codedistanceindicationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedistanceindicationtype.GongGetIdentifier(stage)
}

func (codeemissionbandtype *CodeEmissionBandType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeemissionbandtype, codeemissionbandtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeemissionbandtype *CodeEmissionBandType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeemissionbandtype.GongGetIdentifier(stage)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeequipmentanticollisiontype, codeequipmentanticollisiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeequipmentanticollisiontype.GongGetIdentifier(stage)
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeequipmentunavailabletype, codeequipmentunavailabletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeequipmentunavailabletype.GongGetIdentifier(stage)
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefacilityrankingtype, codefacilityrankingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefacilityrankingtype *CodeFacilityRankingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefacilityrankingtype.GongGetIdentifier(stage)
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefinalguidancetype, codefinalguidancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefinalguidancetype *CodeFinalGuidanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefinalguidancetype.GongGetIdentifier(stage)
}

func (codefirefightingtype *CodeFireFightingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefirefightingtype, codefirefightingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefirefightingtype *CodeFireFightingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefirefightingtype.GongGetIdentifier(stage)
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflightdestinationtype, codeflightdestinationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflightdestinationtype *CodeFlightDestinationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflightdestinationtype.GongGetIdentifier(stage)
}

func (codeflightorigintype *CodeFlightOriginType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflightorigintype, codeflightorigintype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflightorigintype *CodeFlightOriginType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflightorigintype.GongGetIdentifier(stage)
}

func (codeflightpurposetype *CodeFlightPurposeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflightpurposetype, codeflightpurposetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflightpurposetype *CodeFlightPurposeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflightpurposetype.GongGetIdentifier(stage)
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflightrestrictiondesignatortype, codeflightrestrictiondesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflightrestrictiondesignatortype.GongGetIdentifier(stage)
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflightrestrictiontype, codeflightrestrictiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflightrestrictiontype *CodeFlightRestrictionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflightrestrictiontype.GongGetIdentifier(stage)
}

func (codeflightruletype *CodeFlightRuleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflightruletype, codeflightruletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflightruletype *CodeFlightRuleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflightruletype.GongGetIdentifier(stage)
}

func (codeflightstatustype *CodeFlightStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflightstatustype, codeflightstatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflightstatustype *CodeFlightStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflightstatustype.GongGetIdentifier(stage)
}

func (codeflighttype *CodeFlightType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflighttype, codeflighttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflighttype *CodeFlightType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflighttype.GongGetIdentifier(stage)
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflowconditionoperationtype, codeflowconditionoperationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflowconditionoperationtype.GongGetIdentifier(stage)
}

func (codefreeflighttype *CodeFreeFlightType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefreeflighttype, codefreeflighttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefreeflighttype *CodeFreeFlightType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefreeflighttype.GongGetIdentifier(stage)
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefrictiondevicetype, codefrictiondevicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefrictiondevicetype.GongGetIdentifier(stage)
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefrictionestimatetype, codefrictionestimatetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefrictionestimatetype.GongGetIdentifier(stage)
}

func (codefueltype *CodeFuelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefueltype, codefueltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefueltype *CodeFuelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefueltype.GongGetIdentifier(stage)
}

func (codegeobordertype *CodeGeoBorderType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codegeobordertype, codegeobordertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codegeobordertype *CodeGeoBorderType) GongGetReferenceIdentifier(stage *Stage) string {
	return codegeobordertype.GongGetIdentifier(stage)
}

func (codegradeseparationtype *CodeGradeSeparationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codegradeseparationtype, codegradeseparationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codegradeseparationtype *CodeGradeSeparationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codegradeseparationtype.GongGetIdentifier(stage)
}

func (codegroundlightingtype *CodeGroundLightingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codegroundlightingtype, codegroundlightingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codegroundlightingtype *CodeGroundLightingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codegroundlightingtype.GongGetIdentifier(stage)
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeguidancelinetype, codeguidancelinetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeguidancelinetype *CodeGuidanceLineType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeguidancelinetype.GongGetIdentifier(stage)
}

func (codeheightreferencetype *CodeHeightReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeheightreferencetype, codeheightreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeheightreferencetype *CodeHeightReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeheightreferencetype.GongGetIdentifier(stage)
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codehelicopterperformancetype, codehelicopterperformancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codehelicopterperformancetype.GongGetIdentifier(stage)
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeholdingcategorytype, codeholdingcategorytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeholdingcategorytype.GongGetIdentifier(stage)
}

func (codeholdingusagetype *CodeHoldingUsageType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeholdingusagetype, codeholdingusagetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeholdingusagetype *CodeHoldingUsageType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeholdingusagetype.GongGetIdentifier(stage)
}

func (codeholdingusetype *CodeHoldingUseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeholdingusetype, codeholdingusetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeholdingusetype *CodeHoldingUseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeholdingusetype.GongGetIdentifier(stage)
}

func (codeiatatype *CodeIATAType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeiatatype, codeiatatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeiatatype *CodeIATAType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeiatatype.GongGetIdentifier(stage)
}

func (codeicaocountrytype *CodeICAOCountryType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeicaocountrytype, codeicaocountrytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeicaocountrytype *CodeICAOCountryType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeicaocountrytype.GongGetIdentifier(stage)
}

func (codeicaotype *CodeICAOType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeicaotype, codeicaotype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeicaotype *CodeICAOType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeicaotype.GongGetIdentifier(stage)
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeilsbackcoursetype, codeilsbackcoursetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeilsbackcoursetype.GongGetIdentifier(stage)
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeintegritylevelilstype, codeintegritylevelilstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeintegritylevelilstype.GongGetIdentifier(stage)
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeintensitystandbytype, codeintensitystandbytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeintensitystandbytype *CodeIntensityStandByType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeintensitystandbytype.GongGetIdentifier(stage)
}

func (codelanguagetype *CodeLanguageType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelanguagetype, codelanguagetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelanguagetype *CodeLanguageType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelanguagetype.GongGetIdentifier(stage)
}

func (codelevelseriestype *CodeLevelSeriesType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelevelseriestype, codelevelseriestype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelevelseriestype *CodeLevelSeriesType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelevelseriestype.GongGetIdentifier(stage)
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeleveltabledesignatortype, codeleveltabledesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeleveltabledesignatortype.GongGetIdentifier(stage)
}

func (codeleveltype *CodeLevelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeleveltype, codeleveltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeleveltype *CodeLevelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeleveltype.GongGetIdentifier(stage)
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelightholdingpositiontype, codelightholdingpositiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelightholdingpositiontype.GongGetIdentifier(stage)
}

func (codelightintensitytype *CodeLightIntensityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelightintensitytype, codelightintensitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelightintensitytype *CodeLightIntensityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelightintensitytype.GongGetIdentifier(stage)
}

func (codelightsourcetype *CodeLightSourceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelightsourcetype, codelightsourcetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelightsourcetype *CodeLightSourceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelightsourcetype.GongGetIdentifier(stage)
}

func (codelightingjartype *CodeLightingJARType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelightingjartype, codelightingjartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelightingjartype *CodeLightingJARType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelightingjartype.GongGetIdentifier(stage)
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeloadingbridgetype, codeloadingbridgetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeloadingbridgetype *CodeLoadingBridgeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeloadingbridgetype.GongGetIdentifier(stage)
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelocationqualifiertype, codelocationqualifiertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelocationqualifiertype *CodeLocationQualifierType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelocationqualifiertype.GongGetIdentifier(stage)
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelogicaloperatortype, codelogicaloperatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelogicaloperatortype.GongGetIdentifier(stage)
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemlsazimuthtype, codemlsazimuthtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemlsazimuthtype *CodeMLSAzimuthType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemlsazimuthtype.GongGetIdentifier(stage)
}

func (codemlschanneltype *CodeMLSChannelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemlschanneltype, codemlschanneltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemlschanneltype *CodeMLSChannelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemlschanneltype.GongGetIdentifier(stage)
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemarkerbeaconsignaltype, codemarkerbeaconsignaltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemarkerbeaconsignaltype.GongGetIdentifier(stage)
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemarkingconditiontype, codemarkingconditiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemarkingconditiontype *CodeMarkingConditionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemarkingconditiontype.GongGetIdentifier(stage)
}

func (codemarkingstyletype *CodeMarkingStyleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemarkingstyletype, codemarkingstyletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemarkingstyletype *CodeMarkingStyleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemarkingstyletype.GongGetIdentifier(stage)
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemeteoconditionstype, codemeteoconditionstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemeteoconditionstype.GongGetIdentifier(stage)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemilitaryoperationstype, codemilitaryoperationstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemilitaryoperationstype.GongGetIdentifier(stage)
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemilitaryroutepointtype, codemilitaryroutepointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemilitaryroutepointtype.GongGetIdentifier(stage)
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemilitarystatustype, codemilitarystatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemilitarystatustype *CodeMilitaryStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemilitarystatustype.GongGetIdentifier(stage)
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemilitarytrainingtype, codemilitarytrainingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemilitarytrainingtype.GongGetIdentifier(stage)
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeminimafinalapproachpathtype, codeminimafinalapproachpathtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeminimafinalapproachpathtype.GongGetIdentifier(stage)
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeminimumaltitudetype, codeminimumaltitudetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeminimumaltitudetype.GongGetIdentifier(stage)
}

func (codeminimumheighttype *CodeMinimumHeightType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeminimumheighttype, codeminimumheighttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeminimumheighttype *CodeMinimumHeightType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeminimumheighttype.GongGetIdentifier(stage)
}

func (codemissedapproachtype *CodeMissedApproachType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemissedapproachtype, codemissedapproachtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemissedapproachtype *CodeMissedApproachType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemissedapproachtype.GongGetIdentifier(stage)
}

func (codendbusagetype *CodeNDBUsageType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codendbusagetype, codendbusagetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codendbusagetype *CodeNDBUsageType) GongGetReferenceIdentifier(stage *Stage) string {
	return codendbusagetype.GongGetIdentifier(stage)
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavaiddesignatortype, codenavaiddesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavaiddesignatortype.GongGetIdentifier(stage)
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavaidpurposetype, codenavaidpurposetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavaidpurposetype *CodeNavaidPurposeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavaidpurposetype.GongGetIdentifier(stage)
}

func (codenavaidservicetype *CodeNavaidServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavaidservicetype, codenavaidservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavaidservicetype *CodeNavaidServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavaidservicetype.GongGetIdentifier(stage)
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavigationarearestrictiontype, codenavigationarearestrictiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavigationarearestrictiontype.GongGetIdentifier(stage)
}

func (codenavigationareatype *CodeNavigationAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavigationareatype, codenavigationareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavigationareatype *CodeNavigationAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavigationareatype.GongGetIdentifier(stage)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavigationequipmenttype, codenavigationequipmenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavigationequipmenttype.GongGetIdentifier(stage)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenavigationspecificationtype, codenavigationspecificationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenavigationspecificationtype.GongGetIdentifier(stage)
}

func (codenitrogentype *CodeNitrogenType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenitrogentype, codenitrogentype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenitrogentype *CodeNitrogenType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenitrogentype.GongGetIdentifier(stage)
}

func (codenorthreferencetype *CodeNorthReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenorthreferencetype, codenorthreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenorthreferencetype *CodeNorthReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenorthreferencetype.GongGetIdentifier(stage)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenotepurposetype, codenotepurposetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenotepurposetype *CodeNotePurposeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenotepurposetype.GongGetIdentifier(stage)
}

func (codeobstacleareatype *CodeObstacleAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeobstacleareatype, codeobstacleareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeobstacleareatype *CodeObstacleAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeobstacleareatype.GongGetIdentifier(stage)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeobstacleassessmentsurfacetype, codeobstacleassessmentsurfacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeobstacleassessmentsurfacetype.GongGetIdentifier(stage)
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeobstructionidsurfacezonetype, codeobstructionidsurfacezonetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeobstructionidsurfacezonetype.GongGetIdentifier(stage)
}

func (codeoiltype *CodeOilType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeoiltype, codeoiltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeoiltype *CodeOilType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeoiltype.GongGetIdentifier(stage)
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeoperationairportheliporttype, codeoperationairportheliporttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeoperationairportheliporttype.GongGetIdentifier(stage)
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeoperationmanoeuvringareatype, codeoperationmanoeuvringareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeoperationmanoeuvringareatype.GongGetIdentifier(stage)
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeorganisationdesignatortype, codeorganisationdesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeorganisationdesignatortype.GongGetIdentifier(stage)
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeorganisationhierarchytype, codeorganisationhierarchytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeorganisationhierarchytype.GongGetIdentifier(stage)
}

func (codeorganisationtype *CodeOrganisationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeorganisationtype, codeorganisationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeorganisationtype *CodeOrganisationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeorganisationtype.GongGetIdentifier(stage)
}

func (codeoxygentype *CodeOxygenType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeoxygentype, codeoxygentype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeoxygentype *CodeOxygenType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeoxygentype.GongGetIdentifier(stage)
}

func (codepartype *CodePARType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepartype, codepartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepartype *CodePARType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepartype.GongGetIdentifier(stage)
}

func (codepcnmethodtype *CodePCNMethodType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepcnmethodtype, codepcnmethodtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepcnmethodtype *CodePCNMethodType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepcnmethodtype.GongGetIdentifier(stage)
}

func (codepcnpavementtype *CodePCNPavementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepcnpavementtype, codepcnpavementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepcnpavementtype *CodePCNPavementType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepcnpavementtype.GongGetIdentifier(stage)
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepcnsubgradetype, codepcnsubgradetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepcnsubgradetype *CodePCNSubgradeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepcnsubgradetype.GongGetIdentifier(stage)
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepcntyrepressuretype, codepcntyrepressuretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepcntyrepressuretype *CodePCNTyrePressureType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepcntyrepressuretype.GongGetIdentifier(stage)
}

func (codepassengerservicetype *CodePassengerServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepassengerservicetype, codepassengerservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepassengerservicetype *CodePassengerServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepassengerservicetype.GongGetIdentifier(stage)
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepilotcontrolledlightingtype, codepilotcontrolledlightingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepilotcontrolledlightingtype.GongGetIdentifier(stage)
}

func (codepositioninilstype *CodePositionInILSType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepositioninilstype, codepositioninilstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepositioninilstype *CodePositionInILSType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepositioninilstype.GongGetIdentifier(stage)
}

func (codeprimaryradartype *CodePrimaryRadarType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeprimaryradartype, codeprimaryradartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeprimaryradartype *CodePrimaryRadarType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeprimaryradartype.GongGetIdentifier(stage)
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeprocedureavailabilitytype, codeprocedureavailabilitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeprocedureavailabilitytype.GongGetIdentifier(stage)
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeprocedurecodingstandardtype, codeprocedurecodingstandardtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeprocedurecodingstandardtype.GongGetIdentifier(stage)
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeproceduredistancetype, codeproceduredistancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeproceduredistancetype *CodeProcedureDistanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeproceduredistancetype.GongGetIdentifier(stage)
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeprocedurefixroletype, codeprocedurefixroletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeprocedurefixroletype.GongGetIdentifier(stage)
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeprocedurephasetype, codeprocedurephasetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeprocedurephasetype *CodeProcedurePhaseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeprocedurephasetype.GongGetIdentifier(stage)
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeprotectareasectiontype, codeprotectareasectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeprotectareasectiontype.GongGetIdentifier(stage)
}

func (codernptype *CodeRNPType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codernptype, codernptype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codernptype *CodeRNPType) GongGetReferenceIdentifier(stage *Stage) string {
	return codernptype.GongGetIdentifier(stage)
}

func (codervrreadingtype *CodeRVRReadingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codervrreadingtype, codervrreadingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codervrreadingtype *CodeRVRReadingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codervrreadingtype.GongGetIdentifier(stage)
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codervsmpointroletype, codervsmpointroletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codervsmpointroletype.GongGetIdentifier(stage)
}

func (codervsmtype *CodeRVSMType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codervsmtype, codervsmtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codervsmtype *CodeRVSMType) GongGetReferenceIdentifier(stage *Stage) string {
	return codervsmtype.GongGetIdentifier(stage)
}

func (coderadarservicetype *CodeRadarServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderadarservicetype, coderadarservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderadarservicetype *CodeRadarServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderadarservicetype.GongGetIdentifier(stage)
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderadioemissiontype, coderadioemissiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderadioemissiontype *CodeRadioEmissionType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderadioemissiontype.GongGetIdentifier(stage)
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderadiofrequencyareatype, coderadiofrequencyareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderadiofrequencyareatype.GongGetIdentifier(stage)
}

func (coderadiosignaltype *CodeRadioSignalType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderadiosignaltype, coderadiosignaltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderadiosignaltype *CodeRadioSignalType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderadiosignaltype.GongGetIdentifier(stage)
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codereferencepathidentifiertype, codereferencepathidentifiertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetReferenceIdentifier(stage *Stage) string {
	return codereferencepathidentifiertype.GongGetIdentifier(stage)
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codereferenceroletype, codereferenceroletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codereferenceroletype *CodeReferenceRoleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codereferenceroletype.GongGetIdentifier(stage)
}

func (codereflectortype *CodeReflectorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codereflectortype, codereflectortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codereflectortype *CodeReflectorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codereflectortype.GongGetIdentifier(stage)
}

func (coderelativepositiontype *CodeRelativePositionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderelativepositiontype, coderelativepositiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderelativepositiontype *CodeRelativePositionType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderelativepositiontype.GongGetIdentifier(stage)
}

func (coderoadtype *CodeRoadType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderoadtype, coderoadtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderoadtype *CodeRoadType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderoadtype.GongGetIdentifier(stage)
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderouteavailabilitytype, coderouteavailabilitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderouteavailabilitytype.GongGetIdentifier(stage)
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderoutedesignatorlettertype, coderoutedesignatorlettertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderoutedesignatorlettertype.GongGetIdentifier(stage)
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderoutedesignatorprefixtype, coderoutedesignatorprefixtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderoutedesignatorprefixtype.GongGetIdentifier(stage)
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderoutedesignatorsuffixtype, coderoutedesignatorsuffixtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderoutedesignatorsuffixtype.GongGetIdentifier(stage)
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderouteindicatortype, coderouteindicatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderouteindicatortype *CodeRouteIndicatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderouteindicatortype.GongGetIdentifier(stage)
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderoutenavigationtype, coderoutenavigationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderoutenavigationtype *CodeRouteNavigationType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderoutenavigationtype.GongGetIdentifier(stage)
}

func (coderouteorigintype *CodeRouteOriginType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderouteorigintype, coderouteorigintype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderouteorigintype *CodeRouteOriginType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderouteorigintype.GongGetIdentifier(stage)
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderoutesegmentpathtype, coderoutesegmentpathtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderoutesegmentpathtype.GongGetIdentifier(stage)
}

func (coderoutetype *CodeRouteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderoutetype, coderoutetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderoutetype *CodeRouteType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderoutetype.GongGetIdentifier(stage)
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderuleproceduretitletype, coderuleproceduretitletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderuleproceduretitletype.GongGetIdentifier(stage)
}

func (coderuleproceduretype *CodeRuleProcedureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderuleproceduretype, coderuleproceduretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderuleproceduretype *CodeRuleProcedureType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderuleproceduretype.GongGetIdentifier(stage)
}

func (coderunwayelementtype *CodeRunwayElementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderunwayelementtype, coderunwayelementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderunwayelementtype *CodeRunwayElementType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderunwayelementtype.GongGetIdentifier(stage)
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderunwaymarkingtype, coderunwaymarkingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderunwaymarkingtype *CodeRunwayMarkingType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderunwaymarkingtype.GongGetIdentifier(stage)
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderunwaypointroletype, coderunwaypointroletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderunwaypointroletype *CodeRunwayPointRoleType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderunwaypointroletype.GongGetIdentifier(stage)
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderunwayprotectionareatype, coderunwayprotectionareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderunwayprotectionareatype.GongGetIdentifier(stage)
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderunwaysectiontype, coderunwaysectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderunwaysectiontype *CodeRunwaySectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderunwaysectiontype.GongGetIdentifier(stage)
}

func (coderunwaytype *CodeRunwayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderunwaytype, coderunwaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderunwaytype *CodeRunwayType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderunwaytype.GongGetIdentifier(stage)
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesafealtitudetype, codesafealtitudetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesafealtitudetype *CodeSafeAltitudeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesafealtitudetype.GongGetIdentifier(stage)
}

func (codesegmentpathtype *CodeSegmentPathType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesegmentpathtype, codesegmentpathtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesegmentpathtype *CodeSegmentPathType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesegmentpathtype.GongGetIdentifier(stage)
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesegmentterminationtype, codesegmentterminationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesegmentterminationtype *CodeSegmentTerminationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesegmentterminationtype.GongGetIdentifier(stage)
}

func (codeserviceatctype *CodeServiceATCType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeserviceatctype, codeserviceatctype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeserviceatctype *CodeServiceATCType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeserviceatctype.GongGetIdentifier(stage)
}

func (codeserviceatfmtype *CodeServiceATFMType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeserviceatfmtype, codeserviceatfmtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeserviceatfmtype *CodeServiceATFMType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeserviceatfmtype.GongGetIdentifier(stage)
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeservicegroundcontroltype, codeservicegroundcontroltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeservicegroundcontroltype.GongGetIdentifier(stage)
}

func (codeserviceinformationtype *CodeServiceInformationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeserviceinformationtype, codeserviceinformationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeserviceinformationtype *CodeServiceInformationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeserviceinformationtype.GongGetIdentifier(stage)
}

func (codeservicesartype *CodeServiceSARType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeservicesartype, codeservicesartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeservicesartype *CodeServiceSARType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeservicesartype.GongGetIdentifier(stage)
}

func (codesidetype *CodeSideType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesidetype, codesidetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesidetype *CodeSideType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesidetype.GongGetIdentifier(stage)
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesignalperformanceilstype, codesignalperformanceilstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesignalperformanceilstype.GongGetIdentifier(stage)
}

func (codespecialdatetype *CodeSpecialDateType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codespecialdatetype, codespecialdatetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codespecialdatetype *CodeSpecialDateType) GongGetReferenceIdentifier(stage *Stage) string {
	return codespecialdatetype.GongGetIdentifier(stage)
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codespecialnavigationchaindesignatortype, codespecialnavigationchaindesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codespecialnavigationchaindesignatortype.GongGetIdentifier(stage)
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codespecialnavigationstationtype, codespecialnavigationstationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codespecialnavigationstationtype.GongGetIdentifier(stage)
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codespecialnavigationsystemtype, codespecialnavigationsystemtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongGetReferenceIdentifier(stage *Stage) string {
	return codespecialnavigationsystemtype.GongGetIdentifier(stage)
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codespeedreferencetype, codespeedreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codespeedreferencetype *CodeSpeedReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codespeedreferencetype.GongGetIdentifier(stage)
}

func (codestandbypowertype *CodeStandbyPowerType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestandbypowertype, codestandbypowertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestandbypowertype *CodeStandbyPowerType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestandbypowertype.GongGetIdentifier(stage)
}

func (codestatusairporttype *CodeStatusAirportType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusairporttype, codestatusairporttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusairporttype *CodeStatusAirportType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusairporttype.GongGetIdentifier(stage)
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusairspacetype, codestatusairspacetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusairspacetype *CodeStatusAirspaceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusairspacetype.GongGetIdentifier(stage)
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusconstructiontype, codestatusconstructiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusconstructiontype *CodeStatusConstructionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusconstructiontype.GongGetIdentifier(stage)
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusnavaidtype, codestatusnavaidtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusnavaidtype *CodeStatusNavaidType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusnavaidtype.GongGetIdentifier(stage)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusoperationstype, codestatusoperationstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusoperationstype *CodeStatusOperationsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusoperationstype.GongGetIdentifier(stage)
}

func (codestatusservicetype *CodeStatusServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusservicetype, codestatusservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusservicetype *CodeStatusServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusservicetype.GongGetIdentifier(stage)
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesurfacecompositiontype, codesurfacecompositiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesurfacecompositiontype.GongGetIdentifier(stage)
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesurfaceconditiontype, codesurfaceconditiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesurfaceconditiontype *CodeSurfaceConditionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesurfaceconditiontype.GongGetIdentifier(stage)
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesurfacepreparationtype, codesurfacepreparationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesurfacepreparationtype *CodeSurfacePreparationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesurfacepreparationtype.GongGetIdentifier(stage)
}

func (codesystemactivationtype *CodeSystemActivationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesystemactivationtype, codesystemactivationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesystemactivationtype *CodeSystemActivationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesystemactivationtype.GongGetIdentifier(stage)
}

func (codetaatype *CodeTAAType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetaatype, codetaatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetaatype *CodeTAAType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetaatype.GongGetIdentifier(stage)
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetacanchanneltype, codetacanchanneltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetacanchanneltype *CodeTACANChannelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetacanchanneltype.GongGetIdentifier(stage)
}

func (codetlofsectiontype *CodeTLOFSectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetlofsectiontype, codetlofsectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetlofsectiontype *CodeTLOFSectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetlofsectiontype.GongGetIdentifier(stage)
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetaxiwayelementtype, codetaxiwayelementtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetaxiwayelementtype *CodeTaxiwayElementType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetaxiwayelementtype.GongGetIdentifier(stage)
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetaxiwaysectiontype, codetaxiwaysectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetaxiwaysectiontype.GongGetIdentifier(stage)
}

func (codetaxiwaytype *CodeTaxiwayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetaxiwaytype, codetaxiwaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetaxiwaytype *CodeTaxiwayType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetaxiwaytype.GongGetIdentifier(stage)
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetelecomnetworktype, codetelecomnetworktype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetelecomnetworktype *CodeTelecomNetworkType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetelecomnetworktype.GongGetIdentifier(stage)
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetimeeventcombinationtype, codetimeeventcombinationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetimeeventcombinationtype.GongGetIdentifier(stage)
}

func (codetimeeventtype *CodeTimeEventType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetimeeventtype, codetimeeventtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetimeeventtype *CodeTimeEventType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetimeeventtype.GongGetIdentifier(stage)
}

func (codetimereferencetype *CodeTimeReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetimereferencetype, codetimereferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetimereferencetype *CodeTimeReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetimereferencetype.GongGetIdentifier(stage)
}

func (codetrajectorytype *CodeTrajectoryType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetrajectorytype, codetrajectorytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetrajectorytype *CodeTrajectoryType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetrajectorytype.GongGetIdentifier(stage)
}

func (codetranspondertype *CodeTransponderType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetranspondertype, codetranspondertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetranspondertype *CodeTransponderType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetranspondertype.GongGetIdentifier(stage)
}

func (codeuuidtype *CodeUUIDType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeuuidtype, codeuuidtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeuuidtype *CodeUUIDType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeuuidtype.GongGetIdentifier(stage)
}

func (codeunitdependencytype *CodeUnitDependencyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeunitdependencytype, codeunitdependencytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeunitdependencytype *CodeUnitDependencyType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeunitdependencytype.GongGetIdentifier(stage)
}

func (codeunittype *CodeUnitType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeunittype, codeunittype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeunittype *CodeUnitType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeunittype.GongGetIdentifier(stage)
}

func (codeupperalphatype *CodeUpperAlphaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeupperalphatype, codeupperalphatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeupperalphatype *CodeUpperAlphaType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeupperalphatype.GongGetIdentifier(stage)
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeusagelimitationtype, codeusagelimitationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeusagelimitationtype *CodeUsageLimitationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeusagelimitationtype.GongGetIdentifier(stage)
}

func (codevasistype *CodeVASISType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codevasistype, codevasistype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codevasistype *CodeVASISType) GongGetReferenceIdentifier(stage *Stage) string {
	return codevasistype.GongGetIdentifier(stage)
}

func (codevortype *CodeVORType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codevortype, codevortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codevortype *CodeVORType) GongGetReferenceIdentifier(stage *Stage) string {
	return codevortype.GongGetIdentifier(stage)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codevalueinterpretationtype, codevalueinterpretationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codevalueinterpretationtype.GongGetIdentifier(stage)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticaldatumtype, codeverticaldatumtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticaldatumtype *CodeVerticalDatumType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticaldatumtype.GongGetIdentifier(stage)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticalreferencetype, codeverticalreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticalreferencetype.GongGetIdentifier(stage)
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticalstructuremarkingtype, codeverticalstructuremarkingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticalstructuremarkingtype.GongGetIdentifier(stage)
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticalstructurematerialtype, codeverticalstructurematerialtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticalstructurematerialtype.GongGetIdentifier(stage)
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticalstructuretype, codeverticalstructuretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticalstructuretype *CodeVerticalStructureType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticalstructuretype.GongGetIdentifier(stage)
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codevisualdockingguidancetype, codevisualdockingguidancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codevisualdockingguidancetype.GongGetIdentifier(stage)
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codewaketurbulencetype, codewaketurbulencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codewaketurbulencetype.GongGetIdentifier(stage)
}

func (codeworkareatype *CodeWorkAreaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeworkareatype, codeworkareatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeworkareatype *CodeWorkAreaType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeworkareatype.GongGetIdentifier(stage)
}

func (codeyesnotype *CodeYesNoType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeyesnotype, codeyesnotype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeyesnotype *CodeYesNoType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeyesnotype.GongGetIdentifier(stage)
}

func (datemonthdaytype *DateMonthDayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datemonthdaytype, datemonthdaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datemonthdaytype *DateMonthDayType) GongGetReferenceIdentifier(stage *Stage) string {
	return datemonthdaytype.GongGetIdentifier(stage)
}

func (datetimetype *DateTimeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datetimetype, datetimetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datetimetype *DateTimeType) GongGetReferenceIdentifier(stage *Stage) string {
	return datetimetype.GongGetIdentifier(stage)
}

func (datetype *DateType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datetype, datetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datetype *DateType) GongGetReferenceIdentifier(stage *Stage) string {
	return datetype.GongGetIdentifier(stage)
}

func (dateyeartype *DateYearType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dateyeartype, dateyeartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dateyeartype *DateYearType) GongGetReferenceIdentifier(stage *Stage) string {
	return dateyeartype.GongGetIdentifier(stage)
}

func (nonumbertype *NoNumberType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(nonumbertype, nonumbertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (nonumbertype *NoNumberType) GongGetReferenceIdentifier(stage *Stage) string {
	return nonumbertype.GongGetIdentifier(stage)
}

func (nosequencetype *NoSequenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(nosequencetype, nosequencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (nosequencetype *NoSequenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return nosequencetype.GongGetIdentifier(stage)
}

func (textaddresstype *TextAddressType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textaddresstype, textaddresstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textaddresstype *TextAddressType) GongGetReferenceIdentifier(stage *Stage) string {
	return textaddresstype.GongGetIdentifier(stage)
}

func (textdesignatortype *TextDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textdesignatortype, textdesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textdesignatortype *TextDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return textdesignatortype.GongGetIdentifier(stage)
}

func (textinstructiontype *TextInstructionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textinstructiontype, textinstructiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textinstructiontype *TextInstructionType) GongGetReferenceIdentifier(stage *Stage) string {
	return textinstructiontype.GongGetIdentifier(stage)
}

func (textnametype *TextNameType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textnametype, textnametype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textnametype *TextNameType) GongGetReferenceIdentifier(stage *Stage) string {
	return textnametype.GongGetIdentifier(stage)
}

func (textnotetype *TextNoteType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textnotetype, textnotetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textnotetype *TextNoteType) GongGetReferenceIdentifier(stage *Stage) string {
	return textnotetype.GongGetIdentifier(stage)
}

func (textphonetype *TextPhoneType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textphonetype, textphonetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textphonetype *TextPhoneType) GongGetReferenceIdentifier(stage *Stage) string {
	return textphonetype.GongGetIdentifier(stage)
}

func (textpropertynametype *TextPropertyNameType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textpropertynametype, textpropertynametype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textpropertynametype *TextPropertyNameType) GongGetReferenceIdentifier(stage *Stage) string {
	return textpropertynametype.GongGetIdentifier(stage)
}

func (textremarktype *TextRemarkType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textremarktype, textremarktype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textremarktype *TextRemarkType) GongGetReferenceIdentifier(stage *Stage) string {
	return textremarktype.GongGetIdentifier(stage)
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textsidstardesignatortype, textsidstardesignatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return textsidstardesignatortype.GongGetIdentifier(stage)
}

func (timetype *TimeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timetype, timetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timetype *TimeType) GongGetReferenceIdentifier(stage *Stage) string {
	return timetype.GongGetIdentifier(stage)
}

func (valalarmlimittype *ValAlarmLimitType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valalarmlimittype, valalarmlimittype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valalarmlimittype *ValAlarmLimitType) GongGetReferenceIdentifier(stage *Stage) string {
	return valalarmlimittype.GongGetIdentifier(stage)
}

func (valangletype *ValAngleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valangletype, valangletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valangletype *ValAngleType) GongGetReferenceIdentifier(stage *Stage) string {
	return valangletype.GongGetIdentifier(stage)
}

func (valbearingtype *ValBearingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valbearingtype, valbearingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valbearingtype *ValBearingType) GongGetReferenceIdentifier(stage *Stage) string {
	return valbearingtype.GongGetIdentifier(stage)
}

func (valchannelnumbertype *ValChannelNumberType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valchannelnumbertype, valchannelnumbertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valchannelnumbertype *ValChannelNumberType) GongGetReferenceIdentifier(stage *Stage) string {
	return valchannelnumbertype.GongGetIdentifier(stage)
}

func (valdepthtype *ValDepthType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdepthtype, valdepthtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdepthtype *ValDepthType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdepthtype.GongGetIdentifier(stage)
}

func (valdistancesignedtype *ValDistanceSignedType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdistancesignedtype, valdistancesignedtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdistancesignedtype *ValDistanceSignedType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdistancesignedtype.GongGetIdentifier(stage)
}

func (valdistancetype *ValDistanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdistancetype, valdistancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdistancetype *ValDistanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdistancetype.GongGetIdentifier(stage)
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdistanceverticaltype, valdistanceverticaltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdistanceverticaltype *ValDistanceVerticalType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdistanceverticaltype.GongGetIdentifier(stage)
}

func (valdurationtype *ValDurationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valdurationtype, valdurationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valdurationtype *ValDurationType) GongGetReferenceIdentifier(stage *Stage) string {
	return valdurationtype.GongGetIdentifier(stage)
}

func (valfltype *ValFLType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valfltype, valfltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valfltype *ValFLType) GongGetReferenceIdentifier(stage *Stage) string {
	return valfltype.GongGetIdentifier(stage)
}

func (valfrequencytype *ValFrequencyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valfrequencytype, valfrequencytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valfrequencytype *ValFrequencyType) GongGetReferenceIdentifier(stage *Stage) string {
	return valfrequencytype.GongGetIdentifier(stage)
}

func (valfrictiontype *ValFrictionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valfrictiontype, valfrictiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valfrictiontype *ValFrictionType) GongGetReferenceIdentifier(stage *Stage) string {
	return valfrictiontype.GongGetIdentifier(stage)
}

func (valhextype *ValHexType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valhextype, valhextype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valhextype *ValHexType) GongGetReferenceIdentifier(stage *Stage) string {
	return valhextype.GongGetIdentifier(stage)
}

func (vallcntype *ValLCNType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vallcntype, vallcntype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vallcntype *ValLCNType) GongGetReferenceIdentifier(stage *Stage) string {
	return vallcntype.GongGetIdentifier(stage)
}

func (vallightintensitytype *ValLightIntensityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(vallightintensitytype, vallightintensitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (vallightintensitytype *ValLightIntensityType) GongGetReferenceIdentifier(stage *Stage) string {
	return vallightintensitytype.GongGetIdentifier(stage)
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valmagneticvariationchangetype, valmagneticvariationchangetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongGetReferenceIdentifier(stage *Stage) string {
	return valmagneticvariationchangetype.GongGetIdentifier(stage)
}

func (valmagneticvariationtype *ValMagneticVariationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valmagneticvariationtype, valmagneticvariationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valmagneticvariationtype *ValMagneticVariationType) GongGetReferenceIdentifier(stage *Stage) string {
	return valmagneticvariationtype.GongGetIdentifier(stage)
}

func (valpcntype *ValPCNType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valpcntype, valpcntype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valpcntype *ValPCNType) GongGetReferenceIdentifier(stage *Stage) string {
	return valpcntype.GongGetIdentifier(stage)
}

func (valpercenttype *ValPercentType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valpercenttype, valpercenttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valpercenttype *ValPercentType) GongGetReferenceIdentifier(stage *Stage) string {
	return valpercenttype.GongGetIdentifier(stage)
}

func (valpressuretype *ValPressureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valpressuretype, valpressuretype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valpressuretype *ValPressureType) GongGetReferenceIdentifier(stage *Stage) string {
	return valpressuretype.GongGetIdentifier(stage)
}

func (valslopetype *ValSlopeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valslopetype, valslopetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valslopetype *ValSlopeType) GongGetReferenceIdentifier(stage *Stage) string {
	return valslopetype.GongGetIdentifier(stage)
}

func (valspeedtype *ValSpeedType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valspeedtype, valspeedtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valspeedtype *ValSpeedType) GongGetReferenceIdentifier(stage *Stage) string {
	return valspeedtype.GongGetIdentifier(stage)
}

func (valtemperaturetype *ValTemperatureType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valtemperaturetype, valtemperaturetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valtemperaturetype *ValTemperatureType) GongGetReferenceIdentifier(stage *Stage) string {
	return valtemperaturetype.GongGetIdentifier(stage)
}

func (valweighttype *ValWeightType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valweighttype, valweighttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valweighttype *ValWeightType) GongGetReferenceIdentifier(stage *Stage) string {
	return valweighttype.GongGetIdentifier(stage)
}

func (xhtmltype *XHTMLType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(xhtmltype, xhtmltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xhtmltype *XHTMLType) GongGetReferenceIdentifier(stage *Stage) string {
	return xhtmltype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (codeatcreportingtype *CodeATCReportingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeatcreportingtype.GongGetIdentifier(stage), "CodeATCReportingType", codeatcreportingtype.Name)
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaerialrefuellingpointtype.GongGetIdentifier(stage), "CodeAerialRefuellingPointType", codeaerialrefuellingpointtype.Name)
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaerialrefuellingprefixtype.GongGetIdentifier(stage), "CodeAerialRefuellingPrefixType", codeaerialrefuellingprefixtype.Name)
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaerialrefuellingtype.GongGetIdentifier(stage), "CodeAerialRefuellingType", codeaerialrefuellingtype.Name)
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftcategorytype.GongGetIdentifier(stage), "CodeAircraftCategoryType", codeaircraftcategorytype.Name)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftenginenumbertype.GongGetIdentifier(stage), "CodeAircraftEngineNumberType", codeaircraftenginenumbertype.Name)
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftenginetype.GongGetIdentifier(stage), "CodeAircraftEngineType", codeaircraftenginetype.Name)
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftgroundservicetype.GongGetIdentifier(stage), "CodeAircraftGroundServiceType", codeaircraftgroundservicetype.Name)
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircrafticaotype.GongGetIdentifier(stage), "CodeAircraftICAOType", codeaircrafticaotype.Name)
}

func (codeaircraftstandtype *CodeAircraftStandType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftstandtype.GongGetIdentifier(stage), "CodeAircraftStandType", codeaircraftstandtype.Name)
}

func (codeaircrafttype *CodeAircraftType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircrafttype.GongGetIdentifier(stage), "CodeAircraftType", codeaircrafttype.Name)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftwingspanclasstype.GongGetIdentifier(stage), "CodeAircraftWingspanClassType", codeaircraftwingspanclasstype.Name)
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairportheliportcollocationtype.GongGetIdentifier(stage), "CodeAirportHeliportCollocationType", codeairportheliportcollocationtype.Name)
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairportheliportdesignatortype.GongGetIdentifier(stage), "CodeAirportHeliportDesignatorType", codeairportheliportdesignatortype.Name)
}

func (codeairportheliporttype *CodeAirportHeliportType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairportheliporttype.GongGetIdentifier(stage), "CodeAirportHeliportType", codeairportheliporttype.Name)
}

func (codeairportwarningtype *CodeAirportWarningType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairportwarningtype.GongGetIdentifier(stage), "CodeAirportWarningType", codeairportwarningtype.Name)
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspaceactivitytype.GongGetIdentifier(stage), "CodeAirspaceActivityType", codeairspaceactivitytype.Name)
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspaceaggregationtype.GongGetIdentifier(stage), "CodeAirspaceAggregationType", codeairspaceaggregationtype.Name)
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspaceclassificationtype.GongGetIdentifier(stage), "CodeAirspaceClassificationType", codeairspaceclassificationtype.Name)
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacedependencytype.GongGetIdentifier(stage), "CodeAirspaceDependencyType", codeairspacedependencytype.Name)
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacedesignatortype.GongGetIdentifier(stage), "CodeAirspaceDesignatorType", codeairspacedesignatortype.Name)
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacepointpositiontype.GongGetIdentifier(stage), "CodeAirspacePointPositionType", codeairspacepointpositiontype.Name)
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacepointroletype.GongGetIdentifier(stage), "CodeAirspacePointRoleType", codeairspacepointroletype.Name)
}

func (codeairspacetype *CodeAirspaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacetype.GongGetIdentifier(stage), "CodeAirspaceType", codeairspacetype.Name)
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codealtitudeadjustmenttype.GongGetIdentifier(stage), "CodeAltitudeAdjustmentType", codealtitudeadjustmenttype.Name)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codealtitudeusetype.GongGetIdentifier(stage), "CodeAltitudeUseType", codealtitudeusetype.Name)
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachequipmentadditionaltype.GongGetIdentifier(stage), "CodeApproachEquipmentAdditionalType", codeapproachequipmentadditionaltype.Name)
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachguidancetype.GongGetIdentifier(stage), "CodeApproachGuidanceType", codeapproachguidancetype.Name)
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachlightingicaotype.GongGetIdentifier(stage), "CodeApproachLightingICAOType", codeapproachlightingicaotype.Name)
}

func (codeapproachlightingtype *CodeApproachLightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachlightingtype.GongGetIdentifier(stage), "CodeApproachLightingType", codeapproachlightingtype.Name)
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachprefixtype.GongGetIdentifier(stage), "CodeApproachPrefixType", codeapproachprefixtype.Name)
}

func (codeapproachtype *CodeApproachType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachtype.GongGetIdentifier(stage), "CodeApproachType", codeapproachtype.Name)
}

func (codeapprovaltype *CodeApprovalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapprovaltype.GongGetIdentifier(stage), "CodeApprovalType", codeapprovaltype.Name)
}

func (codeapronelementtype *CodeApronElementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapronelementtype.GongGetIdentifier(stage), "CodeApronElementType", codeapronelementtype.Name)
}

func (codeapronsectiontype *CodeApronSectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapronsectiontype.GongGetIdentifier(stage), "CodeApronSectionType", codeapronsectiontype.Name)
}

func (codearcdirectiontype *CodeArcDirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codearcdirectiontype.GongGetIdentifier(stage), "CodeArcDirectionType", codearcdirectiontype.Name)
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codearrestinggearenergyabsorbtype.GongGetIdentifier(stage), "CodeArrestingGearEnergyAbsorbType", codearrestinggearenergyabsorbtype.Name)
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codearrestinggearengagedevicetype.GongGetIdentifier(stage), "CodeArrestingGearEngageDeviceType", codearrestinggearengagedevicetype.Name)
}

func (codeauralmorsetype *CodeAuralMorseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeauralmorsetype.GongGetIdentifier(stage), "CodeAuralMorseType", codeauralmorsetype.Name)
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeauthorityroletype.GongGetIdentifier(stage), "CodeAuthorityRoleType", codeauthorityroletype.Name)
}

func (codeauthoritytype *CodeAuthorityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeauthoritytype.GongGetIdentifier(stage), "CodeAuthorityType", codeauthoritytype.Name)
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaviationstandardstype.GongGetIdentifier(stage), "CodeAviationStandardsType", codeaviationstandardstype.Name)
}

func (codebearingtype *CodeBearingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codebearingtype.GongGetIdentifier(stage), "CodeBearingType", codebearingtype.Name)
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codebuoydesignatortype.GongGetIdentifier(stage), "CodeBuoyDesignatorType", codebuoydesignatortype.Name)
}

func (codebuoytype *CodeBuoyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codebuoytype.GongGetIdentifier(stage), "CodeBuoyType", codebuoytype.Name)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecardinaldirectiontype.GongGetIdentifier(stage), "CodeCardinalDirectionType", codecardinaldirectiontype.Name)
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecheckpointcategorytype.GongGetIdentifier(stage), "CodeCheckpointCategoryType", codecheckpointcategorytype.Name)
}

func (codecolourtype *CodeColourType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecolourtype.GongGetIdentifier(stage), "CodeColourType", codecolourtype.Name)
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecommunicationchanneltype.GongGetIdentifier(stage), "CodeCommunicationChannelType", codecommunicationchanneltype.Name)
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecommunicationdirectiontype.GongGetIdentifier(stage), "CodeCommunicationDirectionType", codecommunicationdirectiontype.Name)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecommunicationmodetype.GongGetIdentifier(stage), "CodeCommunicationModeType", codecommunicationmodetype.Name)
}

func (codecomparisontype *CodeComparisonType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecomparisontype.GongGetIdentifier(stage), "CodeComparisonType", codecomparisontype.Name)
}

func (codecontaminationtype *CodeContaminationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecontaminationtype.GongGetIdentifier(stage), "CodeContaminationType", codecontaminationtype.Name)
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecoursequalityilstype.GongGetIdentifier(stage), "CodeCourseQualityILSType", codecoursequalityilstype.Name)
}

func (codecoursetype *CodeCourseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecoursetype.GongGetIdentifier(stage), "CodeCourseType", codecoursetype.Name)
}

func (codedmechanneltype *CodeDMEChannelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedmechanneltype.GongGetIdentifier(stage), "CodeDMEChannelType", codedmechanneltype.Name)
}

func (codedmetype *CodeDMEType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedmetype.GongGetIdentifier(stage), "CodeDMEType", codedmetype.Name)
}

func (codedaytype *CodeDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedaytype.GongGetIdentifier(stage), "CodeDayType", codedaytype.Name)
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedeclareddistancetype.GongGetIdentifier(stage), "CodeDeclaredDistanceType", codedeclareddistancetype.Name)
}

func (codedesignstandardtype *CodeDesignStandardType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedesignstandardtype.GongGetIdentifier(stage), "CodeDesignStandardType", codedesignstandardtype.Name)
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedesignatedpointdesignatortype.GongGetIdentifier(stage), "CodeDesignatedPointDesignatorType", codedesignatedpointdesignatortype.Name)
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedesignatedpointtype.GongGetIdentifier(stage), "CodeDesignatedPointType", codedesignatedpointtype.Name)
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedirectionreferencetype.GongGetIdentifier(stage), "CodeDirectionReferenceType", codedirectionreferencetype.Name)
}

func (codedirectionturntype *CodeDirectionTurnType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedirectionturntype.GongGetIdentifier(stage), "CodeDirectionTurnType", codedirectionturntype.Name)
}

func (codedirectiontype *CodeDirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedirectiontype.GongGetIdentifier(stage), "CodeDirectionType", codedirectiontype.Name)
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedistanceindicationtype.GongGetIdentifier(stage), "CodeDistanceIndicationType", codedistanceindicationtype.Name)
}

func (codeemissionbandtype *CodeEmissionBandType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeemissionbandtype.GongGetIdentifier(stage), "CodeEmissionBandType", codeemissionbandtype.Name)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeequipmentanticollisiontype.GongGetIdentifier(stage), "CodeEquipmentAntiCollisionType", codeequipmentanticollisiontype.Name)
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeequipmentunavailabletype.GongGetIdentifier(stage), "CodeEquipmentUnavailableType", codeequipmentunavailabletype.Name)
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefacilityrankingtype.GongGetIdentifier(stage), "CodeFacilityRankingType", codefacilityrankingtype.Name)
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefinalguidancetype.GongGetIdentifier(stage), "CodeFinalGuidanceType", codefinalguidancetype.Name)
}

func (codefirefightingtype *CodeFireFightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefirefightingtype.GongGetIdentifier(stage), "CodeFireFightingType", codefirefightingtype.Name)
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightdestinationtype.GongGetIdentifier(stage), "CodeFlightDestinationType", codeflightdestinationtype.Name)
}

func (codeflightorigintype *CodeFlightOriginType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightorigintype.GongGetIdentifier(stage), "CodeFlightOriginType", codeflightorigintype.Name)
}

func (codeflightpurposetype *CodeFlightPurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightpurposetype.GongGetIdentifier(stage), "CodeFlightPurposeType", codeflightpurposetype.Name)
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightrestrictiondesignatortype.GongGetIdentifier(stage), "CodeFlightRestrictionDesignatorType", codeflightrestrictiondesignatortype.Name)
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightrestrictiontype.GongGetIdentifier(stage), "CodeFlightRestrictionType", codeflightrestrictiontype.Name)
}

func (codeflightruletype *CodeFlightRuleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightruletype.GongGetIdentifier(stage), "CodeFlightRuleType", codeflightruletype.Name)
}

func (codeflightstatustype *CodeFlightStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightstatustype.GongGetIdentifier(stage), "CodeFlightStatusType", codeflightstatustype.Name)
}

func (codeflighttype *CodeFlightType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflighttype.GongGetIdentifier(stage), "CodeFlightType", codeflighttype.Name)
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflowconditionoperationtype.GongGetIdentifier(stage), "CodeFlowConditionOperationType", codeflowconditionoperationtype.Name)
}

func (codefreeflighttype *CodeFreeFlightType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefreeflighttype.GongGetIdentifier(stage), "CodeFreeFlightType", codefreeflighttype.Name)
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefrictiondevicetype.GongGetIdentifier(stage), "CodeFrictionDeviceType", codefrictiondevicetype.Name)
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefrictionestimatetype.GongGetIdentifier(stage), "CodeFrictionEstimateType", codefrictionestimatetype.Name)
}

func (codefueltype *CodeFuelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefueltype.GongGetIdentifier(stage), "CodeFuelType", codefueltype.Name)
}

func (codegeobordertype *CodeGeoBorderType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codegeobordertype.GongGetIdentifier(stage), "CodeGeoBorderType", codegeobordertype.Name)
}

func (codegradeseparationtype *CodeGradeSeparationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codegradeseparationtype.GongGetIdentifier(stage), "CodeGradeSeparationType", codegradeseparationtype.Name)
}

func (codegroundlightingtype *CodeGroundLightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codegroundlightingtype.GongGetIdentifier(stage), "CodeGroundLightingType", codegroundlightingtype.Name)
}

func (codeguidancelinetype *CodeGuidanceLineType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeguidancelinetype.GongGetIdentifier(stage), "CodeGuidanceLineType", codeguidancelinetype.Name)
}

func (codeheightreferencetype *CodeHeightReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeheightreferencetype.GongGetIdentifier(stage), "CodeHeightReferenceType", codeheightreferencetype.Name)
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codehelicopterperformancetype.GongGetIdentifier(stage), "CodeHelicopterPerformanceType", codehelicopterperformancetype.Name)
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeholdingcategorytype.GongGetIdentifier(stage), "CodeHoldingCategoryType", codeholdingcategorytype.Name)
}

func (codeholdingusagetype *CodeHoldingUsageType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeholdingusagetype.GongGetIdentifier(stage), "CodeHoldingUsageType", codeholdingusagetype.Name)
}

func (codeholdingusetype *CodeHoldingUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeholdingusetype.GongGetIdentifier(stage), "CodeHoldingUseType", codeholdingusetype.Name)
}

func (codeiatatype *CodeIATAType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeiatatype.GongGetIdentifier(stage), "CodeIATAType", codeiatatype.Name)
}

func (codeicaocountrytype *CodeICAOCountryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeicaocountrytype.GongGetIdentifier(stage), "CodeICAOCountryType", codeicaocountrytype.Name)
}

func (codeicaotype *CodeICAOType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeicaotype.GongGetIdentifier(stage), "CodeICAOType", codeicaotype.Name)
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeilsbackcoursetype.GongGetIdentifier(stage), "CodeILSBackCourseType", codeilsbackcoursetype.Name)
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeintegritylevelilstype.GongGetIdentifier(stage), "CodeIntegrityLevelILSType", codeintegritylevelilstype.Name)
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeintensitystandbytype.GongGetIdentifier(stage), "CodeIntensityStandByType", codeintensitystandbytype.Name)
}

func (codelanguagetype *CodeLanguageType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelanguagetype.GongGetIdentifier(stage), "CodeLanguageType", codelanguagetype.Name)
}

func (codelevelseriestype *CodeLevelSeriesType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelevelseriestype.GongGetIdentifier(stage), "CodeLevelSeriesType", codelevelseriestype.Name)
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeleveltabledesignatortype.GongGetIdentifier(stage), "CodeLevelTableDesignatorType", codeleveltabledesignatortype.Name)
}

func (codeleveltype *CodeLevelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeleveltype.GongGetIdentifier(stage), "CodeLevelType", codeleveltype.Name)
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelightholdingpositiontype.GongGetIdentifier(stage), "CodeLightHoldingPositionType", codelightholdingpositiontype.Name)
}

func (codelightintensitytype *CodeLightIntensityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelightintensitytype.GongGetIdentifier(stage), "CodeLightIntensityType", codelightintensitytype.Name)
}

func (codelightsourcetype *CodeLightSourceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelightsourcetype.GongGetIdentifier(stage), "CodeLightSourceType", codelightsourcetype.Name)
}

func (codelightingjartype *CodeLightingJARType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelightingjartype.GongGetIdentifier(stage), "CodeLightingJARType", codelightingjartype.Name)
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeloadingbridgetype.GongGetIdentifier(stage), "CodeLoadingBridgeType", codeloadingbridgetype.Name)
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelocationqualifiertype.GongGetIdentifier(stage), "CodeLocationQualifierType", codelocationqualifiertype.Name)
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelogicaloperatortype.GongGetIdentifier(stage), "CodeLogicalOperatorType", codelogicaloperatortype.Name)
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemlsazimuthtype.GongGetIdentifier(stage), "CodeMLSAzimuthType", codemlsazimuthtype.Name)
}

func (codemlschanneltype *CodeMLSChannelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemlschanneltype.GongGetIdentifier(stage), "CodeMLSChannelType", codemlschanneltype.Name)
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemarkerbeaconsignaltype.GongGetIdentifier(stage), "CodeMarkerBeaconSignalType", codemarkerbeaconsignaltype.Name)
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemarkingconditiontype.GongGetIdentifier(stage), "CodeMarkingConditionType", codemarkingconditiontype.Name)
}

func (codemarkingstyletype *CodeMarkingStyleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemarkingstyletype.GongGetIdentifier(stage), "CodeMarkingStyleType", codemarkingstyletype.Name)
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemeteoconditionstype.GongGetIdentifier(stage), "CodeMeteoConditionsType", codemeteoconditionstype.Name)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemilitaryoperationstype.GongGetIdentifier(stage), "CodeMilitaryOperationsType", codemilitaryoperationstype.Name)
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemilitaryroutepointtype.GongGetIdentifier(stage), "CodeMilitaryRoutePointType", codemilitaryroutepointtype.Name)
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemilitarystatustype.GongGetIdentifier(stage), "CodeMilitaryStatusType", codemilitarystatustype.Name)
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemilitarytrainingtype.GongGetIdentifier(stage), "CodeMilitaryTrainingType", codemilitarytrainingtype.Name)
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeminimafinalapproachpathtype.GongGetIdentifier(stage), "CodeMinimaFinalApproachPathType", codeminimafinalapproachpathtype.Name)
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeminimumaltitudetype.GongGetIdentifier(stage), "CodeMinimumAltitudeType", codeminimumaltitudetype.Name)
}

func (codeminimumheighttype *CodeMinimumHeightType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeminimumheighttype.GongGetIdentifier(stage), "CodeMinimumHeightType", codeminimumheighttype.Name)
}

func (codemissedapproachtype *CodeMissedApproachType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemissedapproachtype.GongGetIdentifier(stage), "CodeMissedApproachType", codemissedapproachtype.Name)
}

func (codendbusagetype *CodeNDBUsageType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codendbusagetype.GongGetIdentifier(stage), "CodeNDBUsageType", codendbusagetype.Name)
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavaiddesignatortype.GongGetIdentifier(stage), "CodeNavaidDesignatorType", codenavaiddesignatortype.Name)
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavaidpurposetype.GongGetIdentifier(stage), "CodeNavaidPurposeType", codenavaidpurposetype.Name)
}

func (codenavaidservicetype *CodeNavaidServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavaidservicetype.GongGetIdentifier(stage), "CodeNavaidServiceType", codenavaidservicetype.Name)
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationarearestrictiontype.GongGetIdentifier(stage), "CodeNavigationAreaRestrictionType", codenavigationarearestrictiontype.Name)
}

func (codenavigationareatype *CodeNavigationAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationareatype.GongGetIdentifier(stage), "CodeNavigationAreaType", codenavigationareatype.Name)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationequipmenttype.GongGetIdentifier(stage), "CodeNavigationEquipmentType", codenavigationequipmenttype.Name)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationspecificationtype.GongGetIdentifier(stage), "CodeNavigationSpecificationType", codenavigationspecificationtype.Name)
}

func (codenitrogentype *CodeNitrogenType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenitrogentype.GongGetIdentifier(stage), "CodeNitrogenType", codenitrogentype.Name)
}

func (codenorthreferencetype *CodeNorthReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenorthreferencetype.GongGetIdentifier(stage), "CodeNorthReferenceType", codenorthreferencetype.Name)
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenotepurposetype.GongGetIdentifier(stage), "CodeNotePurposeType", codenotepurposetype.Name)
}

func (codeobstacleareatype *CodeObstacleAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstacleareatype.GongGetIdentifier(stage), "CodeObstacleAreaType", codeobstacleareatype.Name)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstacleassessmentsurfacetype.GongGetIdentifier(stage), "CodeObstacleAssessmentSurfaceType", codeobstacleassessmentsurfacetype.Name)
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstructionidsurfacezonetype.GongGetIdentifier(stage), "CodeObstructionIdSurfaceZoneType", codeobstructionidsurfacezonetype.Name)
}

func (codeoiltype *CodeOilType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeoiltype.GongGetIdentifier(stage), "CodeOilType", codeoiltype.Name)
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeoperationairportheliporttype.GongGetIdentifier(stage), "CodeOperationAirportHeliportType", codeoperationairportheliporttype.Name)
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeoperationmanoeuvringareatype.GongGetIdentifier(stage), "CodeOperationManoeuvringAreaType", codeoperationmanoeuvringareatype.Name)
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeorganisationdesignatortype.GongGetIdentifier(stage), "CodeOrganisationDesignatorType", codeorganisationdesignatortype.Name)
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeorganisationhierarchytype.GongGetIdentifier(stage), "CodeOrganisationHierarchyType", codeorganisationhierarchytype.Name)
}

func (codeorganisationtype *CodeOrganisationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeorganisationtype.GongGetIdentifier(stage), "CodeOrganisationType", codeorganisationtype.Name)
}

func (codeoxygentype *CodeOxygenType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeoxygentype.GongGetIdentifier(stage), "CodeOxygenType", codeoxygentype.Name)
}

func (codepartype *CodePARType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepartype.GongGetIdentifier(stage), "CodePARType", codepartype.Name)
}

func (codepcnmethodtype *CodePCNMethodType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepcnmethodtype.GongGetIdentifier(stage), "CodePCNMethodType", codepcnmethodtype.Name)
}

func (codepcnpavementtype *CodePCNPavementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepcnpavementtype.GongGetIdentifier(stage), "CodePCNPavementType", codepcnpavementtype.Name)
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepcnsubgradetype.GongGetIdentifier(stage), "CodePCNSubgradeType", codepcnsubgradetype.Name)
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepcntyrepressuretype.GongGetIdentifier(stage), "CodePCNTyrePressureType", codepcntyrepressuretype.Name)
}

func (codepassengerservicetype *CodePassengerServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepassengerservicetype.GongGetIdentifier(stage), "CodePassengerServiceType", codepassengerservicetype.Name)
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepilotcontrolledlightingtype.GongGetIdentifier(stage), "CodePilotControlledLightingType", codepilotcontrolledlightingtype.Name)
}

func (codepositioninilstype *CodePositionInILSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepositioninilstype.GongGetIdentifier(stage), "CodePositionInILSType", codepositioninilstype.Name)
}

func (codeprimaryradartype *CodePrimaryRadarType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeprimaryradartype.GongGetIdentifier(stage), "CodePrimaryRadarType", codeprimaryradartype.Name)
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeprocedureavailabilitytype.GongGetIdentifier(stage), "CodeProcedureAvailabilityType", codeprocedureavailabilitytype.Name)
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeprocedurecodingstandardtype.GongGetIdentifier(stage), "CodeProcedureCodingStandardType", codeprocedurecodingstandardtype.Name)
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeproceduredistancetype.GongGetIdentifier(stage), "CodeProcedureDistanceType", codeproceduredistancetype.Name)
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeprocedurefixroletype.GongGetIdentifier(stage), "CodeProcedureFixRoleType", codeprocedurefixroletype.Name)
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeprocedurephasetype.GongGetIdentifier(stage), "CodeProcedurePhaseType", codeprocedurephasetype.Name)
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeprotectareasectiontype.GongGetIdentifier(stage), "CodeProtectAreaSectionType", codeprotectareasectiontype.Name)
}

func (codernptype *CodeRNPType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codernptype.GongGetIdentifier(stage), "CodeRNPType", codernptype.Name)
}

func (codervrreadingtype *CodeRVRReadingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervrreadingtype.GongGetIdentifier(stage), "CodeRVRReadingType", codervrreadingtype.Name)
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmpointroletype.GongGetIdentifier(stage), "CodeRVSMPointRoleType", codervsmpointroletype.Name)
}

func (codervsmtype *CodeRVSMType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmtype.GongGetIdentifier(stage), "CodeRVSMType", codervsmtype.Name)
}

func (coderadarservicetype *CodeRadarServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderadarservicetype.GongGetIdentifier(stage), "CodeRadarServiceType", coderadarservicetype.Name)
}

func (coderadioemissiontype *CodeRadioEmissionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderadioemissiontype.GongGetIdentifier(stage), "CodeRadioEmissionType", coderadioemissiontype.Name)
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderadiofrequencyareatype.GongGetIdentifier(stage), "CodeRadioFrequencyAreaType", coderadiofrequencyareatype.Name)
}

func (coderadiosignaltype *CodeRadioSignalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderadiosignaltype.GongGetIdentifier(stage), "CodeRadioSignalType", coderadiosignaltype.Name)
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codereferencepathidentifiertype.GongGetIdentifier(stage), "CodeReferencePathIdentifierType", codereferencepathidentifiertype.Name)
}

func (codereferenceroletype *CodeReferenceRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codereferenceroletype.GongGetIdentifier(stage), "CodeReferenceRoleType", codereferenceroletype.Name)
}

func (codereflectortype *CodeReflectorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codereflectortype.GongGetIdentifier(stage), "CodeReflectorType", codereflectortype.Name)
}

func (coderelativepositiontype *CodeRelativePositionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderelativepositiontype.GongGetIdentifier(stage), "CodeRelativePositionType", coderelativepositiontype.Name)
}

func (coderoadtype *CodeRoadType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderoadtype.GongGetIdentifier(stage), "CodeRoadType", coderoadtype.Name)
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderouteavailabilitytype.GongGetIdentifier(stage), "CodeRouteAvailabilityType", coderouteavailabilitytype.Name)
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderoutedesignatorlettertype.GongGetIdentifier(stage), "CodeRouteDesignatorLetterType", coderoutedesignatorlettertype.Name)
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderoutedesignatorprefixtype.GongGetIdentifier(stage), "CodeRouteDesignatorPrefixType", coderoutedesignatorprefixtype.Name)
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderoutedesignatorsuffixtype.GongGetIdentifier(stage), "CodeRouteDesignatorSuffixType", coderoutedesignatorsuffixtype.Name)
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderouteindicatortype.GongGetIdentifier(stage), "CodeRouteIndicatorType", coderouteindicatortype.Name)
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderoutenavigationtype.GongGetIdentifier(stage), "CodeRouteNavigationType", coderoutenavigationtype.Name)
}

func (coderouteorigintype *CodeRouteOriginType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderouteorigintype.GongGetIdentifier(stage), "CodeRouteOriginType", coderouteorigintype.Name)
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderoutesegmentpathtype.GongGetIdentifier(stage), "CodeRouteSegmentPathType", coderoutesegmentpathtype.Name)
}

func (coderoutetype *CodeRouteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderoutetype.GongGetIdentifier(stage), "CodeRouteType", coderoutetype.Name)
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderuleproceduretitletype.GongGetIdentifier(stage), "CodeRuleProcedureTitleType", coderuleproceduretitletype.Name)
}

func (coderuleproceduretype *CodeRuleProcedureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderuleproceduretype.GongGetIdentifier(stage), "CodeRuleProcedureType", coderuleproceduretype.Name)
}

func (coderunwayelementtype *CodeRunwayElementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderunwayelementtype.GongGetIdentifier(stage), "CodeRunwayElementType", coderunwayelementtype.Name)
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderunwaymarkingtype.GongGetIdentifier(stage), "CodeRunwayMarkingType", coderunwaymarkingtype.Name)
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderunwaypointroletype.GongGetIdentifier(stage), "CodeRunwayPointRoleType", coderunwaypointroletype.Name)
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderunwayprotectionareatype.GongGetIdentifier(stage), "CodeRunwayProtectionAreaType", coderunwayprotectionareatype.Name)
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderunwaysectiontype.GongGetIdentifier(stage), "CodeRunwaySectionType", coderunwaysectiontype.Name)
}

func (coderunwaytype *CodeRunwayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderunwaytype.GongGetIdentifier(stage), "CodeRunwayType", coderunwaytype.Name)
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesafealtitudetype.GongGetIdentifier(stage), "CodeSafeAltitudeType", codesafealtitudetype.Name)
}

func (codesegmentpathtype *CodeSegmentPathType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesegmentpathtype.GongGetIdentifier(stage), "CodeSegmentPathType", codesegmentpathtype.Name)
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesegmentterminationtype.GongGetIdentifier(stage), "CodeSegmentTerminationType", codesegmentterminationtype.Name)
}

func (codeserviceatctype *CodeServiceATCType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeserviceatctype.GongGetIdentifier(stage), "CodeServiceATCType", codeserviceatctype.Name)
}

func (codeserviceatfmtype *CodeServiceATFMType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeserviceatfmtype.GongGetIdentifier(stage), "CodeServiceATFMType", codeserviceatfmtype.Name)
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeservicegroundcontroltype.GongGetIdentifier(stage), "CodeServiceGroundControlType", codeservicegroundcontroltype.Name)
}

func (codeserviceinformationtype *CodeServiceInformationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeserviceinformationtype.GongGetIdentifier(stage), "CodeServiceInformationType", codeserviceinformationtype.Name)
}

func (codeservicesartype *CodeServiceSARType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeservicesartype.GongGetIdentifier(stage), "CodeServiceSARType", codeservicesartype.Name)
}

func (codesidetype *CodeSideType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesidetype.GongGetIdentifier(stage), "CodeSideType", codesidetype.Name)
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesignalperformanceilstype.GongGetIdentifier(stage), "CodeSignalPerformanceILSType", codesignalperformanceilstype.Name)
}

func (codespecialdatetype *CodeSpecialDateType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codespecialdatetype.GongGetIdentifier(stage), "CodeSpecialDateType", codespecialdatetype.Name)
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codespecialnavigationchaindesignatortype.GongGetIdentifier(stage), "CodeSpecialNavigationChainDesignatorType", codespecialnavigationchaindesignatortype.Name)
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codespecialnavigationstationtype.GongGetIdentifier(stage), "CodeSpecialNavigationStationType", codespecialnavigationstationtype.Name)
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codespecialnavigationsystemtype.GongGetIdentifier(stage), "CodeSpecialNavigationSystemType", codespecialnavigationsystemtype.Name)
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codespeedreferencetype.GongGetIdentifier(stage), "CodeSpeedReferenceType", codespeedreferencetype.Name)
}

func (codestandbypowertype *CodeStandbyPowerType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestandbypowertype.GongGetIdentifier(stage), "CodeStandbyPowerType", codestandbypowertype.Name)
}

func (codestatusairporttype *CodeStatusAirportType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusairporttype.GongGetIdentifier(stage), "CodeStatusAirportType", codestatusairporttype.Name)
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusairspacetype.GongGetIdentifier(stage), "CodeStatusAirspaceType", codestatusairspacetype.Name)
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusconstructiontype.GongGetIdentifier(stage), "CodeStatusConstructionType", codestatusconstructiontype.Name)
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusnavaidtype.GongGetIdentifier(stage), "CodeStatusNavaidType", codestatusnavaidtype.Name)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusoperationstype.GongGetIdentifier(stage), "CodeStatusOperationsType", codestatusoperationstype.Name)
}

func (codestatusservicetype *CodeStatusServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusservicetype.GongGetIdentifier(stage), "CodeStatusServiceType", codestatusservicetype.Name)
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesurfacecompositiontype.GongGetIdentifier(stage), "CodeSurfaceCompositionType", codesurfacecompositiontype.Name)
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesurfaceconditiontype.GongGetIdentifier(stage), "CodeSurfaceConditionType", codesurfaceconditiontype.Name)
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesurfacepreparationtype.GongGetIdentifier(stage), "CodeSurfacePreparationType", codesurfacepreparationtype.Name)
}

func (codesystemactivationtype *CodeSystemActivationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesystemactivationtype.GongGetIdentifier(stage), "CodeSystemActivationType", codesystemactivationtype.Name)
}

func (codetaatype *CodeTAAType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetaatype.GongGetIdentifier(stage), "CodeTAAType", codetaatype.Name)
}

func (codetacanchanneltype *CodeTACANChannelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetacanchanneltype.GongGetIdentifier(stage), "CodeTACANChannelType", codetacanchanneltype.Name)
}

func (codetlofsectiontype *CodeTLOFSectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetlofsectiontype.GongGetIdentifier(stage), "CodeTLOFSectionType", codetlofsectiontype.Name)
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetaxiwayelementtype.GongGetIdentifier(stage), "CodeTaxiwayElementType", codetaxiwayelementtype.Name)
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetaxiwaysectiontype.GongGetIdentifier(stage), "CodeTaxiwaySectionType", codetaxiwaysectiontype.Name)
}

func (codetaxiwaytype *CodeTaxiwayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetaxiwaytype.GongGetIdentifier(stage), "CodeTaxiwayType", codetaxiwaytype.Name)
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetelecomnetworktype.GongGetIdentifier(stage), "CodeTelecomNetworkType", codetelecomnetworktype.Name)
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetimeeventcombinationtype.GongGetIdentifier(stage), "CodeTimeEventCombinationType", codetimeeventcombinationtype.Name)
}

func (codetimeeventtype *CodeTimeEventType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetimeeventtype.GongGetIdentifier(stage), "CodeTimeEventType", codetimeeventtype.Name)
}

func (codetimereferencetype *CodeTimeReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetimereferencetype.GongGetIdentifier(stage), "CodeTimeReferenceType", codetimereferencetype.Name)
}

func (codetrajectorytype *CodeTrajectoryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetrajectorytype.GongGetIdentifier(stage), "CodeTrajectoryType", codetrajectorytype.Name)
}

func (codetranspondertype *CodeTransponderType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetranspondertype.GongGetIdentifier(stage), "CodeTransponderType", codetranspondertype.Name)
}

func (codeuuidtype *CodeUUIDType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeuuidtype.GongGetIdentifier(stage), "CodeUUIDType", codeuuidtype.Name)
}

func (codeunitdependencytype *CodeUnitDependencyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeunitdependencytype.GongGetIdentifier(stage), "CodeUnitDependencyType", codeunitdependencytype.Name)
}

func (codeunittype *CodeUnitType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeunittype.GongGetIdentifier(stage), "CodeUnitType", codeunittype.Name)
}

func (codeupperalphatype *CodeUpperAlphaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeupperalphatype.GongGetIdentifier(stage), "CodeUpperAlphaType", codeupperalphatype.Name)
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeusagelimitationtype.GongGetIdentifier(stage), "CodeUsageLimitationType", codeusagelimitationtype.Name)
}

func (codevasistype *CodeVASISType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codevasistype.GongGetIdentifier(stage), "CodeVASISType", codevasistype.Name)
}

func (codevortype *CodeVORType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codevortype.GongGetIdentifier(stage), "CodeVORType", codevortype.Name)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codevalueinterpretationtype.GongGetIdentifier(stage), "CodeValueInterpretationType", codevalueinterpretationtype.Name)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticaldatumtype.GongGetIdentifier(stage), "CodeVerticalDatumType", codeverticaldatumtype.Name)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalreferencetype.GongGetIdentifier(stage), "CodeVerticalReferenceType", codeverticalreferencetype.Name)
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalstructuremarkingtype.GongGetIdentifier(stage), "CodeVerticalStructureMarkingType", codeverticalstructuremarkingtype.Name)
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalstructurematerialtype.GongGetIdentifier(stage), "CodeVerticalStructureMaterialType", codeverticalstructurematerialtype.Name)
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalstructuretype.GongGetIdentifier(stage), "CodeVerticalStructureType", codeverticalstructuretype.Name)
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codevisualdockingguidancetype.GongGetIdentifier(stage), "CodeVisualDockingGuidanceType", codevisualdockingguidancetype.Name)
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codewaketurbulencetype.GongGetIdentifier(stage), "CodeWakeTurbulenceType", codewaketurbulencetype.Name)
}

func (codeworkareatype *CodeWorkAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeworkareatype.GongGetIdentifier(stage), "CodeWorkAreaType", codeworkareatype.Name)
}

func (codeyesnotype *CodeYesNoType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeyesnotype.GongGetIdentifier(stage), "CodeYesNoType", codeyesnotype.Name)
}

func (datemonthdaytype *DateMonthDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datemonthdaytype.GongGetIdentifier(stage), "DateMonthDayType", datemonthdaytype.Name)
}

func (datetimetype *DateTimeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datetimetype.GongGetIdentifier(stage), "DateTimeType", datetimetype.Name)
}

func (datetype *DateType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datetype.GongGetIdentifier(stage), "DateType", datetype.Name)
}

func (dateyeartype *DateYearType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dateyeartype.GongGetIdentifier(stage), "DateYearType", dateyeartype.Name)
}

func (nonumbertype *NoNumberType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nonumbertype.GongGetIdentifier(stage), "NoNumberType", nonumbertype.Name)
}

func (nosequencetype *NoSequenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nosequencetype.GongGetIdentifier(stage), "NoSequenceType", nosequencetype.Name)
}

func (textaddresstype *TextAddressType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textaddresstype.GongGetIdentifier(stage), "TextAddressType", textaddresstype.Name)
}

func (textdesignatortype *TextDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textdesignatortype.GongGetIdentifier(stage), "TextDesignatorType", textdesignatortype.Name)
}

func (textinstructiontype *TextInstructionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textinstructiontype.GongGetIdentifier(stage), "TextInstructionType", textinstructiontype.Name)
}

func (textnametype *TextNameType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textnametype.GongGetIdentifier(stage), "TextNameType", textnametype.Name)
}

func (textnotetype *TextNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textnotetype.GongGetIdentifier(stage), "TextNoteType", textnotetype.Name)
}

func (textphonetype *TextPhoneType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textphonetype.GongGetIdentifier(stage), "TextPhoneType", textphonetype.Name)
}

func (textpropertynametype *TextPropertyNameType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textpropertynametype.GongGetIdentifier(stage), "TextPropertyNameType", textpropertynametype.Name)
}

func (textremarktype *TextRemarkType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textremarktype.GongGetIdentifier(stage), "TextRemarkType", textremarktype.Name)
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textsidstardesignatortype.GongGetIdentifier(stage), "TextSIDSTARDesignatorType", textsidstardesignatortype.Name)
}

func (timetype *TimeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timetype.GongGetIdentifier(stage), "TimeType", timetype.Name)
}

func (valalarmlimittype *ValAlarmLimitType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valalarmlimittype.GongGetIdentifier(stage), "ValAlarmLimitType", valalarmlimittype.Name)
}

func (valangletype *ValAngleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valangletype.GongGetIdentifier(stage), "ValAngleType", valangletype.Name)
}

func (valbearingtype *ValBearingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valbearingtype.GongGetIdentifier(stage), "ValBearingType", valbearingtype.Name)
}

func (valchannelnumbertype *ValChannelNumberType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valchannelnumbertype.GongGetIdentifier(stage), "ValChannelNumberType", valchannelnumbertype.Name)
}

func (valdepthtype *ValDepthType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdepthtype.GongGetIdentifier(stage), "ValDepthType", valdepthtype.Name)
}

func (valdistancesignedtype *ValDistanceSignedType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdistancesignedtype.GongGetIdentifier(stage), "ValDistanceSignedType", valdistancesignedtype.Name)
}

func (valdistancetype *ValDistanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdistancetype.GongGetIdentifier(stage), "ValDistanceType", valdistancetype.Name)
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdistanceverticaltype.GongGetIdentifier(stage), "ValDistanceVerticalType", valdistanceverticaltype.Name)
}

func (valdurationtype *ValDurationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valdurationtype.GongGetIdentifier(stage), "ValDurationType", valdurationtype.Name)
}

func (valfltype *ValFLType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valfltype.GongGetIdentifier(stage), "ValFLType", valfltype.Name)
}

func (valfrequencytype *ValFrequencyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valfrequencytype.GongGetIdentifier(stage), "ValFrequencyType", valfrequencytype.Name)
}

func (valfrictiontype *ValFrictionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valfrictiontype.GongGetIdentifier(stage), "ValFrictionType", valfrictiontype.Name)
}

func (valhextype *ValHexType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valhextype.GongGetIdentifier(stage), "ValHexType", valhextype.Name)
}

func (vallcntype *ValLCNType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vallcntype.GongGetIdentifier(stage), "ValLCNType", vallcntype.Name)
}

func (vallightintensitytype *ValLightIntensityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(vallightintensitytype.GongGetIdentifier(stage), "ValLightIntensityType", vallightintensitytype.Name)
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valmagneticvariationchangetype.GongGetIdentifier(stage), "ValMagneticVariationChangeType", valmagneticvariationchangetype.Name)
}

func (valmagneticvariationtype *ValMagneticVariationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valmagneticvariationtype.GongGetIdentifier(stage), "ValMagneticVariationType", valmagneticvariationtype.Name)
}

func (valpcntype *ValPCNType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valpcntype.GongGetIdentifier(stage), "ValPCNType", valpcntype.Name)
}

func (valpercenttype *ValPercentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valpercenttype.GongGetIdentifier(stage), "ValPercentType", valpercenttype.Name)
}

func (valpressuretype *ValPressureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valpressuretype.GongGetIdentifier(stage), "ValPressureType", valpressuretype.Name)
}

func (valslopetype *ValSlopeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valslopetype.GongGetIdentifier(stage), "ValSlopeType", valslopetype.Name)
}

func (valspeedtype *ValSpeedType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valspeedtype.GongGetIdentifier(stage), "ValSpeedType", valspeedtype.Name)
}

func (valtemperaturetype *ValTemperatureType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valtemperaturetype.GongGetIdentifier(stage), "ValTemperatureType", valtemperaturetype.Name)
}

func (valweighttype *ValWeightType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valweighttype.GongGetIdentifier(stage), "ValWeightType", valweighttype.Name)
}

func (xhtmltype *XHTMLType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(xhtmltype.GongGetIdentifier(stage), "XHTMLType", xhtmltype.Name)
}

// insertion point for unstaging
func (codeatcreportingtype *CodeATCReportingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeatcreportingtype.GongGetReferenceIdentifier(stage))
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaerialrefuellingpointtype.GongGetReferenceIdentifier(stage))
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaerialrefuellingprefixtype.GongGetReferenceIdentifier(stage))
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaerialrefuellingtype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftcategorytype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftenginenumbertype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftenginetype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftgroundservicetype.GongGetReferenceIdentifier(stage))
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircrafticaotype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftstandtype *CodeAircraftStandType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftstandtype.GongGetReferenceIdentifier(stage))
}

func (codeaircrafttype *CodeAircraftType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircrafttype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftwingspanclasstype.GongGetReferenceIdentifier(stage))
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairportheliportcollocationtype.GongGetReferenceIdentifier(stage))
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairportheliportdesignatortype.GongGetReferenceIdentifier(stage))
}

func (codeairportheliporttype *CodeAirportHeliportType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairportheliporttype.GongGetReferenceIdentifier(stage))
}

func (codeairportwarningtype *CodeAirportWarningType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairportwarningtype.GongGetReferenceIdentifier(stage))
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspaceactivitytype.GongGetReferenceIdentifier(stage))
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspaceaggregationtype.GongGetReferenceIdentifier(stage))
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspaceclassificationtype.GongGetReferenceIdentifier(stage))
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacedependencytype.GongGetReferenceIdentifier(stage))
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacedesignatortype.GongGetReferenceIdentifier(stage))
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacepointpositiontype.GongGetReferenceIdentifier(stage))
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacepointroletype.GongGetReferenceIdentifier(stage))
}

func (codeairspacetype *CodeAirspaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacetype.GongGetReferenceIdentifier(stage))
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codealtitudeadjustmenttype.GongGetReferenceIdentifier(stage))
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codealtitudeusetype.GongGetReferenceIdentifier(stage))
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachequipmentadditionaltype.GongGetReferenceIdentifier(stage))
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachguidancetype.GongGetReferenceIdentifier(stage))
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachlightingicaotype.GongGetReferenceIdentifier(stage))
}

func (codeapproachlightingtype *CodeApproachLightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachlightingtype.GongGetReferenceIdentifier(stage))
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachprefixtype.GongGetReferenceIdentifier(stage))
}

func (codeapproachtype *CodeApproachType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachtype.GongGetReferenceIdentifier(stage))
}

func (codeapprovaltype *CodeApprovalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapprovaltype.GongGetReferenceIdentifier(stage))
}

func (codeapronelementtype *CodeApronElementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapronelementtype.GongGetReferenceIdentifier(stage))
}

func (codeapronsectiontype *CodeApronSectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapronsectiontype.GongGetReferenceIdentifier(stage))
}

func (codearcdirectiontype *CodeArcDirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codearcdirectiontype.GongGetReferenceIdentifier(stage))
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codearrestinggearenergyabsorbtype.GongGetReferenceIdentifier(stage))
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codearrestinggearengagedevicetype.GongGetReferenceIdentifier(stage))
}

func (codeauralmorsetype *CodeAuralMorseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeauralmorsetype.GongGetReferenceIdentifier(stage))
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeauthorityroletype.GongGetReferenceIdentifier(stage))
}

func (codeauthoritytype *CodeAuthorityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeauthoritytype.GongGetReferenceIdentifier(stage))
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaviationstandardstype.GongGetReferenceIdentifier(stage))
}

func (codebearingtype *CodeBearingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codebearingtype.GongGetReferenceIdentifier(stage))
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codebuoydesignatortype.GongGetReferenceIdentifier(stage))
}

func (codebuoytype *CodeBuoyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codebuoytype.GongGetReferenceIdentifier(stage))
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecardinaldirectiontype.GongGetReferenceIdentifier(stage))
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecheckpointcategorytype.GongGetReferenceIdentifier(stage))
}

func (codecolourtype *CodeColourType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecolourtype.GongGetReferenceIdentifier(stage))
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecommunicationchanneltype.GongGetReferenceIdentifier(stage))
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecommunicationdirectiontype.GongGetReferenceIdentifier(stage))
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecommunicationmodetype.GongGetReferenceIdentifier(stage))
}

func (codecomparisontype *CodeComparisonType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecomparisontype.GongGetReferenceIdentifier(stage))
}

func (codecontaminationtype *CodeContaminationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecontaminationtype.GongGetReferenceIdentifier(stage))
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecoursequalityilstype.GongGetReferenceIdentifier(stage))
}

func (codecoursetype *CodeCourseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecoursetype.GongGetReferenceIdentifier(stage))
}

func (codedmechanneltype *CodeDMEChannelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedmechanneltype.GongGetReferenceIdentifier(stage))
}

func (codedmetype *CodeDMEType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedmetype.GongGetReferenceIdentifier(stage))
}

func (codedaytype *CodeDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedaytype.GongGetReferenceIdentifier(stage))
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedeclareddistancetype.GongGetReferenceIdentifier(stage))
}

func (codedesignstandardtype *CodeDesignStandardType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedesignstandardtype.GongGetReferenceIdentifier(stage))
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedesignatedpointdesignatortype.GongGetReferenceIdentifier(stage))
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedesignatedpointtype.GongGetReferenceIdentifier(stage))
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedirectionreferencetype.GongGetReferenceIdentifier(stage))
}

func (codedirectionturntype *CodeDirectionTurnType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedirectionturntype.GongGetReferenceIdentifier(stage))
}

func (codedirectiontype *CodeDirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedirectiontype.GongGetReferenceIdentifier(stage))
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedistanceindicationtype.GongGetReferenceIdentifier(stage))
}

func (codeemissionbandtype *CodeEmissionBandType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeemissionbandtype.GongGetReferenceIdentifier(stage))
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeequipmentanticollisiontype.GongGetReferenceIdentifier(stage))
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeequipmentunavailabletype.GongGetReferenceIdentifier(stage))
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefacilityrankingtype.GongGetReferenceIdentifier(stage))
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefinalguidancetype.GongGetReferenceIdentifier(stage))
}

func (codefirefightingtype *CodeFireFightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefirefightingtype.GongGetReferenceIdentifier(stage))
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightdestinationtype.GongGetReferenceIdentifier(stage))
}

func (codeflightorigintype *CodeFlightOriginType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightorigintype.GongGetReferenceIdentifier(stage))
}

func (codeflightpurposetype *CodeFlightPurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightpurposetype.GongGetReferenceIdentifier(stage))
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightrestrictiondesignatortype.GongGetReferenceIdentifier(stage))
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightrestrictiontype.GongGetReferenceIdentifier(stage))
}

func (codeflightruletype *CodeFlightRuleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightruletype.GongGetReferenceIdentifier(stage))
}

func (codeflightstatustype *CodeFlightStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightstatustype.GongGetReferenceIdentifier(stage))
}

func (codeflighttype *CodeFlightType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflighttype.GongGetReferenceIdentifier(stage))
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflowconditionoperationtype.GongGetReferenceIdentifier(stage))
}

func (codefreeflighttype *CodeFreeFlightType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefreeflighttype.GongGetReferenceIdentifier(stage))
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefrictiondevicetype.GongGetReferenceIdentifier(stage))
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefrictionestimatetype.GongGetReferenceIdentifier(stage))
}

func (codefueltype *CodeFuelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefueltype.GongGetReferenceIdentifier(stage))
}

func (codegeobordertype *CodeGeoBorderType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codegeobordertype.GongGetReferenceIdentifier(stage))
}

func (codegradeseparationtype *CodeGradeSeparationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codegradeseparationtype.GongGetReferenceIdentifier(stage))
}

func (codegroundlightingtype *CodeGroundLightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codegroundlightingtype.GongGetReferenceIdentifier(stage))
}

func (codeguidancelinetype *CodeGuidanceLineType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeguidancelinetype.GongGetReferenceIdentifier(stage))
}

func (codeheightreferencetype *CodeHeightReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeheightreferencetype.GongGetReferenceIdentifier(stage))
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codehelicopterperformancetype.GongGetReferenceIdentifier(stage))
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeholdingcategorytype.GongGetReferenceIdentifier(stage))
}

func (codeholdingusagetype *CodeHoldingUsageType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeholdingusagetype.GongGetReferenceIdentifier(stage))
}

func (codeholdingusetype *CodeHoldingUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeholdingusetype.GongGetReferenceIdentifier(stage))
}

func (codeiatatype *CodeIATAType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeiatatype.GongGetReferenceIdentifier(stage))
}

func (codeicaocountrytype *CodeICAOCountryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeicaocountrytype.GongGetReferenceIdentifier(stage))
}

func (codeicaotype *CodeICAOType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeicaotype.GongGetReferenceIdentifier(stage))
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeilsbackcoursetype.GongGetReferenceIdentifier(stage))
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeintegritylevelilstype.GongGetReferenceIdentifier(stage))
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeintensitystandbytype.GongGetReferenceIdentifier(stage))
}

func (codelanguagetype *CodeLanguageType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelanguagetype.GongGetReferenceIdentifier(stage))
}

func (codelevelseriestype *CodeLevelSeriesType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelevelseriestype.GongGetReferenceIdentifier(stage))
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeleveltabledesignatortype.GongGetReferenceIdentifier(stage))
}

func (codeleveltype *CodeLevelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeleveltype.GongGetReferenceIdentifier(stage))
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelightholdingpositiontype.GongGetReferenceIdentifier(stage))
}

func (codelightintensitytype *CodeLightIntensityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelightintensitytype.GongGetReferenceIdentifier(stage))
}

func (codelightsourcetype *CodeLightSourceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelightsourcetype.GongGetReferenceIdentifier(stage))
}

func (codelightingjartype *CodeLightingJARType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelightingjartype.GongGetReferenceIdentifier(stage))
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeloadingbridgetype.GongGetReferenceIdentifier(stage))
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelocationqualifiertype.GongGetReferenceIdentifier(stage))
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelogicaloperatortype.GongGetReferenceIdentifier(stage))
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemlsazimuthtype.GongGetReferenceIdentifier(stage))
}

func (codemlschanneltype *CodeMLSChannelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemlschanneltype.GongGetReferenceIdentifier(stage))
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemarkerbeaconsignaltype.GongGetReferenceIdentifier(stage))
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemarkingconditiontype.GongGetReferenceIdentifier(stage))
}

func (codemarkingstyletype *CodeMarkingStyleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemarkingstyletype.GongGetReferenceIdentifier(stage))
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemeteoconditionstype.GongGetReferenceIdentifier(stage))
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemilitaryoperationstype.GongGetReferenceIdentifier(stage))
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemilitaryroutepointtype.GongGetReferenceIdentifier(stage))
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemilitarystatustype.GongGetReferenceIdentifier(stage))
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemilitarytrainingtype.GongGetReferenceIdentifier(stage))
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeminimafinalapproachpathtype.GongGetReferenceIdentifier(stage))
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeminimumaltitudetype.GongGetReferenceIdentifier(stage))
}

func (codeminimumheighttype *CodeMinimumHeightType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeminimumheighttype.GongGetReferenceIdentifier(stage))
}

func (codemissedapproachtype *CodeMissedApproachType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemissedapproachtype.GongGetReferenceIdentifier(stage))
}

func (codendbusagetype *CodeNDBUsageType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codendbusagetype.GongGetReferenceIdentifier(stage))
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavaiddesignatortype.GongGetReferenceIdentifier(stage))
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavaidpurposetype.GongGetReferenceIdentifier(stage))
}

func (codenavaidservicetype *CodeNavaidServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavaidservicetype.GongGetReferenceIdentifier(stage))
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationarearestrictiontype.GongGetReferenceIdentifier(stage))
}

func (codenavigationareatype *CodeNavigationAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationareatype.GongGetReferenceIdentifier(stage))
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationequipmenttype.GongGetReferenceIdentifier(stage))
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationspecificationtype.GongGetReferenceIdentifier(stage))
}

func (codenitrogentype *CodeNitrogenType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenitrogentype.GongGetReferenceIdentifier(stage))
}

func (codenorthreferencetype *CodeNorthReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenorthreferencetype.GongGetReferenceIdentifier(stage))
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenotepurposetype.GongGetReferenceIdentifier(stage))
}

func (codeobstacleareatype *CodeObstacleAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstacleareatype.GongGetReferenceIdentifier(stage))
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstacleassessmentsurfacetype.GongGetReferenceIdentifier(stage))
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstructionidsurfacezonetype.GongGetReferenceIdentifier(stage))
}

func (codeoiltype *CodeOilType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeoiltype.GongGetReferenceIdentifier(stage))
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeoperationairportheliporttype.GongGetReferenceIdentifier(stage))
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeoperationmanoeuvringareatype.GongGetReferenceIdentifier(stage))
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeorganisationdesignatortype.GongGetReferenceIdentifier(stage))
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeorganisationhierarchytype.GongGetReferenceIdentifier(stage))
}

func (codeorganisationtype *CodeOrganisationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeorganisationtype.GongGetReferenceIdentifier(stage))
}

func (codeoxygentype *CodeOxygenType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeoxygentype.GongGetReferenceIdentifier(stage))
}

func (codepartype *CodePARType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepartype.GongGetReferenceIdentifier(stage))
}

func (codepcnmethodtype *CodePCNMethodType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepcnmethodtype.GongGetReferenceIdentifier(stage))
}

func (codepcnpavementtype *CodePCNPavementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepcnpavementtype.GongGetReferenceIdentifier(stage))
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepcnsubgradetype.GongGetReferenceIdentifier(stage))
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepcntyrepressuretype.GongGetReferenceIdentifier(stage))
}

func (codepassengerservicetype *CodePassengerServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepassengerservicetype.GongGetReferenceIdentifier(stage))
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepilotcontrolledlightingtype.GongGetReferenceIdentifier(stage))
}

func (codepositioninilstype *CodePositionInILSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepositioninilstype.GongGetReferenceIdentifier(stage))
}

func (codeprimaryradartype *CodePrimaryRadarType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeprimaryradartype.GongGetReferenceIdentifier(stage))
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeprocedureavailabilitytype.GongGetReferenceIdentifier(stage))
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeprocedurecodingstandardtype.GongGetReferenceIdentifier(stage))
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeproceduredistancetype.GongGetReferenceIdentifier(stage))
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeprocedurefixroletype.GongGetReferenceIdentifier(stage))
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeprocedurephasetype.GongGetReferenceIdentifier(stage))
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeprotectareasectiontype.GongGetReferenceIdentifier(stage))
}

func (codernptype *CodeRNPType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codernptype.GongGetReferenceIdentifier(stage))
}

func (codervrreadingtype *CodeRVRReadingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervrreadingtype.GongGetReferenceIdentifier(stage))
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmpointroletype.GongGetReferenceIdentifier(stage))
}

func (codervsmtype *CodeRVSMType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmtype.GongGetReferenceIdentifier(stage))
}

func (coderadarservicetype *CodeRadarServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderadarservicetype.GongGetReferenceIdentifier(stage))
}

func (coderadioemissiontype *CodeRadioEmissionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderadioemissiontype.GongGetReferenceIdentifier(stage))
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderadiofrequencyareatype.GongGetReferenceIdentifier(stage))
}

func (coderadiosignaltype *CodeRadioSignalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderadiosignaltype.GongGetReferenceIdentifier(stage))
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codereferencepathidentifiertype.GongGetReferenceIdentifier(stage))
}

func (codereferenceroletype *CodeReferenceRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codereferenceroletype.GongGetReferenceIdentifier(stage))
}

func (codereflectortype *CodeReflectorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codereflectortype.GongGetReferenceIdentifier(stage))
}

func (coderelativepositiontype *CodeRelativePositionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderelativepositiontype.GongGetReferenceIdentifier(stage))
}

func (coderoadtype *CodeRoadType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderoadtype.GongGetReferenceIdentifier(stage))
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderouteavailabilitytype.GongGetReferenceIdentifier(stage))
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderoutedesignatorlettertype.GongGetReferenceIdentifier(stage))
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderoutedesignatorprefixtype.GongGetReferenceIdentifier(stage))
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderoutedesignatorsuffixtype.GongGetReferenceIdentifier(stage))
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderouteindicatortype.GongGetReferenceIdentifier(stage))
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderoutenavigationtype.GongGetReferenceIdentifier(stage))
}

func (coderouteorigintype *CodeRouteOriginType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderouteorigintype.GongGetReferenceIdentifier(stage))
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderoutesegmentpathtype.GongGetReferenceIdentifier(stage))
}

func (coderoutetype *CodeRouteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderoutetype.GongGetReferenceIdentifier(stage))
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderuleproceduretitletype.GongGetReferenceIdentifier(stage))
}

func (coderuleproceduretype *CodeRuleProcedureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderuleproceduretype.GongGetReferenceIdentifier(stage))
}

func (coderunwayelementtype *CodeRunwayElementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderunwayelementtype.GongGetReferenceIdentifier(stage))
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderunwaymarkingtype.GongGetReferenceIdentifier(stage))
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderunwaypointroletype.GongGetReferenceIdentifier(stage))
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderunwayprotectionareatype.GongGetReferenceIdentifier(stage))
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderunwaysectiontype.GongGetReferenceIdentifier(stage))
}

func (coderunwaytype *CodeRunwayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderunwaytype.GongGetReferenceIdentifier(stage))
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesafealtitudetype.GongGetReferenceIdentifier(stage))
}

func (codesegmentpathtype *CodeSegmentPathType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesegmentpathtype.GongGetReferenceIdentifier(stage))
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesegmentterminationtype.GongGetReferenceIdentifier(stage))
}

func (codeserviceatctype *CodeServiceATCType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeserviceatctype.GongGetReferenceIdentifier(stage))
}

func (codeserviceatfmtype *CodeServiceATFMType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeserviceatfmtype.GongGetReferenceIdentifier(stage))
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeservicegroundcontroltype.GongGetReferenceIdentifier(stage))
}

func (codeserviceinformationtype *CodeServiceInformationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeserviceinformationtype.GongGetReferenceIdentifier(stage))
}

func (codeservicesartype *CodeServiceSARType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeservicesartype.GongGetReferenceIdentifier(stage))
}

func (codesidetype *CodeSideType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesidetype.GongGetReferenceIdentifier(stage))
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesignalperformanceilstype.GongGetReferenceIdentifier(stage))
}

func (codespecialdatetype *CodeSpecialDateType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codespecialdatetype.GongGetReferenceIdentifier(stage))
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codespecialnavigationchaindesignatortype.GongGetReferenceIdentifier(stage))
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codespecialnavigationstationtype.GongGetReferenceIdentifier(stage))
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codespecialnavigationsystemtype.GongGetReferenceIdentifier(stage))
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codespeedreferencetype.GongGetReferenceIdentifier(stage))
}

func (codestandbypowertype *CodeStandbyPowerType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestandbypowertype.GongGetReferenceIdentifier(stage))
}

func (codestatusairporttype *CodeStatusAirportType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusairporttype.GongGetReferenceIdentifier(stage))
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusairspacetype.GongGetReferenceIdentifier(stage))
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusconstructiontype.GongGetReferenceIdentifier(stage))
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusnavaidtype.GongGetReferenceIdentifier(stage))
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusoperationstype.GongGetReferenceIdentifier(stage))
}

func (codestatusservicetype *CodeStatusServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusservicetype.GongGetReferenceIdentifier(stage))
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesurfacecompositiontype.GongGetReferenceIdentifier(stage))
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesurfaceconditiontype.GongGetReferenceIdentifier(stage))
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesurfacepreparationtype.GongGetReferenceIdentifier(stage))
}

func (codesystemactivationtype *CodeSystemActivationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesystemactivationtype.GongGetReferenceIdentifier(stage))
}

func (codetaatype *CodeTAAType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetaatype.GongGetReferenceIdentifier(stage))
}

func (codetacanchanneltype *CodeTACANChannelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetacanchanneltype.GongGetReferenceIdentifier(stage))
}

func (codetlofsectiontype *CodeTLOFSectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetlofsectiontype.GongGetReferenceIdentifier(stage))
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetaxiwayelementtype.GongGetReferenceIdentifier(stage))
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetaxiwaysectiontype.GongGetReferenceIdentifier(stage))
}

func (codetaxiwaytype *CodeTaxiwayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetaxiwaytype.GongGetReferenceIdentifier(stage))
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetelecomnetworktype.GongGetReferenceIdentifier(stage))
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetimeeventcombinationtype.GongGetReferenceIdentifier(stage))
}

func (codetimeeventtype *CodeTimeEventType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetimeeventtype.GongGetReferenceIdentifier(stage))
}

func (codetimereferencetype *CodeTimeReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetimereferencetype.GongGetReferenceIdentifier(stage))
}

func (codetrajectorytype *CodeTrajectoryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetrajectorytype.GongGetReferenceIdentifier(stage))
}

func (codetranspondertype *CodeTransponderType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetranspondertype.GongGetReferenceIdentifier(stage))
}

func (codeuuidtype *CodeUUIDType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeuuidtype.GongGetReferenceIdentifier(stage))
}

func (codeunitdependencytype *CodeUnitDependencyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeunitdependencytype.GongGetReferenceIdentifier(stage))
}

func (codeunittype *CodeUnitType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeunittype.GongGetReferenceIdentifier(stage))
}

func (codeupperalphatype *CodeUpperAlphaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeupperalphatype.GongGetReferenceIdentifier(stage))
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeusagelimitationtype.GongGetReferenceIdentifier(stage))
}

func (codevasistype *CodeVASISType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codevasistype.GongGetReferenceIdentifier(stage))
}

func (codevortype *CodeVORType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codevortype.GongGetReferenceIdentifier(stage))
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codevalueinterpretationtype.GongGetReferenceIdentifier(stage))
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticaldatumtype.GongGetReferenceIdentifier(stage))
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalreferencetype.GongGetReferenceIdentifier(stage))
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalstructuremarkingtype.GongGetReferenceIdentifier(stage))
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalstructurematerialtype.GongGetReferenceIdentifier(stage))
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalstructuretype.GongGetReferenceIdentifier(stage))
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codevisualdockingguidancetype.GongGetReferenceIdentifier(stage))
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codewaketurbulencetype.GongGetReferenceIdentifier(stage))
}

func (codeworkareatype *CodeWorkAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeworkareatype.GongGetReferenceIdentifier(stage))
}

func (codeyesnotype *CodeYesNoType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeyesnotype.GongGetReferenceIdentifier(stage))
}

func (datemonthdaytype *DateMonthDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datemonthdaytype.GongGetReferenceIdentifier(stage))
}

func (datetimetype *DateTimeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datetimetype.GongGetReferenceIdentifier(stage))
}

func (datetype *DateType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datetype.GongGetReferenceIdentifier(stage))
}

func (dateyeartype *DateYearType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dateyeartype.GongGetReferenceIdentifier(stage))
}

func (nonumbertype *NoNumberType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nonumbertype.GongGetReferenceIdentifier(stage))
}

func (nosequencetype *NoSequenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nosequencetype.GongGetReferenceIdentifier(stage))
}

func (textaddresstype *TextAddressType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textaddresstype.GongGetReferenceIdentifier(stage))
}

func (textdesignatortype *TextDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textdesignatortype.GongGetReferenceIdentifier(stage))
}

func (textinstructiontype *TextInstructionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textinstructiontype.GongGetReferenceIdentifier(stage))
}

func (textnametype *TextNameType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textnametype.GongGetReferenceIdentifier(stage))
}

func (textnotetype *TextNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textnotetype.GongGetReferenceIdentifier(stage))
}

func (textphonetype *TextPhoneType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textphonetype.GongGetReferenceIdentifier(stage))
}

func (textpropertynametype *TextPropertyNameType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textpropertynametype.GongGetReferenceIdentifier(stage))
}

func (textremarktype *TextRemarkType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textremarktype.GongGetReferenceIdentifier(stage))
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textsidstardesignatortype.GongGetReferenceIdentifier(stage))
}

func (timetype *TimeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timetype.GongGetReferenceIdentifier(stage))
}

func (valalarmlimittype *ValAlarmLimitType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valalarmlimittype.GongGetReferenceIdentifier(stage))
}

func (valangletype *ValAngleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valangletype.GongGetReferenceIdentifier(stage))
}

func (valbearingtype *ValBearingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valbearingtype.GongGetReferenceIdentifier(stage))
}

func (valchannelnumbertype *ValChannelNumberType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valchannelnumbertype.GongGetReferenceIdentifier(stage))
}

func (valdepthtype *ValDepthType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdepthtype.GongGetReferenceIdentifier(stage))
}

func (valdistancesignedtype *ValDistanceSignedType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdistancesignedtype.GongGetReferenceIdentifier(stage))
}

func (valdistancetype *ValDistanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdistancetype.GongGetReferenceIdentifier(stage))
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdistanceverticaltype.GongGetReferenceIdentifier(stage))
}

func (valdurationtype *ValDurationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valdurationtype.GongGetReferenceIdentifier(stage))
}

func (valfltype *ValFLType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valfltype.GongGetReferenceIdentifier(stage))
}

func (valfrequencytype *ValFrequencyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valfrequencytype.GongGetReferenceIdentifier(stage))
}

func (valfrictiontype *ValFrictionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valfrictiontype.GongGetReferenceIdentifier(stage))
}

func (valhextype *ValHexType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valhextype.GongGetReferenceIdentifier(stage))
}

func (vallcntype *ValLCNType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vallcntype.GongGetReferenceIdentifier(stage))
}

func (vallightintensitytype *ValLightIntensityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(vallightintensitytype.GongGetReferenceIdentifier(stage))
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valmagneticvariationchangetype.GongGetReferenceIdentifier(stage))
}

func (valmagneticvariationtype *ValMagneticVariationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valmagneticvariationtype.GongGetReferenceIdentifier(stage))
}

func (valpcntype *ValPCNType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valpcntype.GongGetReferenceIdentifier(stage))
}

func (valpercenttype *ValPercentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valpercenttype.GongGetReferenceIdentifier(stage))
}

func (valpressuretype *ValPressureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valpressuretype.GongGetReferenceIdentifier(stage))
}

func (valslopetype *ValSlopeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valslopetype.GongGetReferenceIdentifier(stage))
}

func (valspeedtype *ValSpeedType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valspeedtype.GongGetReferenceIdentifier(stage))
}

func (valtemperaturetype *ValTemperatureType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valtemperaturetype.GongGetReferenceIdentifier(stage))
}

func (valweighttype *ValWeightType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valweighttype.GongGetReferenceIdentifier(stage))
}

func (xhtmltype *XHTMLType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(xhtmltype.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__aixmdatatypes__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
