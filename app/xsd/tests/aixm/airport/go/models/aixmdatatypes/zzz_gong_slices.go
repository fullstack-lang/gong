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
	res = __gong__appendInstances(res, stage.CodeAircraftCategoryTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftEngineNumberTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftEngineTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftICAOTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftStandTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftWingspanClassTypes)

	res = __gong__appendInstances(res, stage.CodeAirportHeliportCollocationTypes)

	res = __gong__appendInstances(res, stage.CodeAirportHeliportDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeAirportHeliportTypes)

	res = __gong__appendInstances(res, stage.CodeAirportWarningTypes)

	res = __gong__appendInstances(res, stage.CodeApproachGuidanceTypes)

	res = __gong__appendInstances(res, stage.CodeApproachLightingICAOTypes)

	res = __gong__appendInstances(res, stage.CodeApproachLightingTypes)

	res = __gong__appendInstances(res, stage.CodeApronElementTypes)

	res = __gong__appendInstances(res, stage.CodeApronSectionTypes)

	res = __gong__appendInstances(res, stage.CodeArrestingGearEnergyAbsorbTypes)

	res = __gong__appendInstances(res, stage.CodeArrestingGearEngageDeviceTypes)

	res = __gong__appendInstances(res, stage.CodeAuthorityRoleTypes)

	res = __gong__appendInstances(res, stage.CodeBuoyDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeBuoyTypes)

	res = __gong__appendInstances(res, stage.CodeColourTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationModeTypes)

	res = __gong__appendInstances(res, stage.CodeContaminationTypes)

	res = __gong__appendInstances(res, stage.CodeDayTypes)

	res = __gong__appendInstances(res, stage.CodeDeclaredDistanceTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionTurnTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeEquipmentAntiCollisionTypes)

	res = __gong__appendInstances(res, stage.CodeFlightOriginTypes)

	res = __gong__appendInstances(res, stage.CodeFlightPurposeTypes)

	res = __gong__appendInstances(res, stage.CodeFlightRuleTypes)

	res = __gong__appendInstances(res, stage.CodeFlightStatusTypes)

	res = __gong__appendInstances(res, stage.CodeFlightTypes)

	res = __gong__appendInstances(res, stage.CodeFrictionDeviceTypes)

	res = __gong__appendInstances(res, stage.CodeFrictionEstimateTypes)

	res = __gong__appendInstances(res, stage.CodeGradeSeparationTypes)

	res = __gong__appendInstances(res, stage.CodeGuidanceLineTypes)

	res = __gong__appendInstances(res, stage.CodeHelicopterPerformanceTypes)

	res = __gong__appendInstances(res, stage.CodeHoldingCategoryTypes)

	res = __gong__appendInstances(res, stage.CodeIATATypes)

	res = __gong__appendInstances(res, stage.CodeICAOTypes)

	res = __gong__appendInstances(res, stage.CodeLightHoldingPositionTypes)

	res = __gong__appendInstances(res, stage.CodeLightIntensityTypes)

	res = __gong__appendInstances(res, stage.CodeLightSourceTypes)

	res = __gong__appendInstances(res, stage.CodeLightingJARTypes)

	res = __gong__appendInstances(res, stage.CodeLoadingBridgeTypes)

	res = __gong__appendInstances(res, stage.CodeLogicalOperatorTypes)

	res = __gong__appendInstances(res, stage.CodeMarkingConditionTypes)

	res = __gong__appendInstances(res, stage.CodeMarkingStyleTypes)

	res = __gong__appendInstances(res, stage.CodeMeteoConditionsTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryOperationsTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryStatusTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationEquipmentTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationSpecificationTypes)

	res = __gong__appendInstances(res, stage.CodeNotePurposeTypes)

	res = __gong__appendInstances(res, stage.CodeOperationAirportHeliportTypes)

	res = __gong__appendInstances(res, stage.CodeOperationManoeuvringAreaTypes)

	res = __gong__appendInstances(res, stage.CodePCNMethodTypes)

	res = __gong__appendInstances(res, stage.CodePCNPavementTypes)

	res = __gong__appendInstances(res, stage.CodePCNSubgradeTypes)

	res = __gong__appendInstances(res, stage.CodePCNTyrePressureTypes)

	res = __gong__appendInstances(res, stage.CodeProtectAreaSectionTypes)

	res = __gong__appendInstances(res, stage.CodeRVRReadingTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMTypes)

	res = __gong__appendInstances(res, stage.CodeRoadTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayElementTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayMarkingTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayPointRoleTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayProtectionAreaTypes)

	res = __gong__appendInstances(res, stage.CodeRunwaySectionTypes)

	res = __gong__appendInstances(res, stage.CodeRunwayTypes)

	res = __gong__appendInstances(res, stage.CodeSideTypes)

	res = __gong__appendInstances(res, stage.CodeStatusAirportTypes)

	res = __gong__appendInstances(res, stage.CodeStatusOperationsTypes)

	res = __gong__appendInstances(res, stage.CodeSurfaceCompositionTypes)

	res = __gong__appendInstances(res, stage.CodeSurfaceConditionTypes)

	res = __gong__appendInstances(res, stage.CodeSurfacePreparationTypes)

	res = __gong__appendInstances(res, stage.CodeTLOFSectionTypes)

	res = __gong__appendInstances(res, stage.CodeTaxiwayElementTypes)

	res = __gong__appendInstances(res, stage.CodeTaxiwaySectionTypes)

	res = __gong__appendInstances(res, stage.CodeTaxiwayTypes)

	res = __gong__appendInstances(res, stage.CodeTelecomNetworkTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventCombinationTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventTypes)

	res = __gong__appendInstances(res, stage.CodeTimeReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeTransponderTypes)

	res = __gong__appendInstances(res, stage.CodeUsageLimitationTypes)

	res = __gong__appendInstances(res, stage.CodeVASISTypes)

	res = __gong__appendInstances(res, stage.CodeValueInterpretationTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalDatumTypes)

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

	res = __gong__appendInstances(res, stage.TimeTypes)

	res = __gong__appendInstances(res, stage.ValAngleTypes)

	res = __gong__appendInstances(res, stage.ValBearingTypes)

	res = __gong__appendInstances(res, stage.ValDepthTypes)

	res = __gong__appendInstances(res, stage.ValDistanceSignedTypes)

	res = __gong__appendInstances(res, stage.ValDistanceTypes)

	res = __gong__appendInstances(res, stage.ValDistanceVerticalTypes)

	res = __gong__appendInstances(res, stage.ValDurationTypes)

	res = __gong__appendInstances(res, stage.ValFLTypes)

	res = __gong__appendInstances(res, stage.ValFrictionTypes)

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

	return
}

// insertion point per named struct
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

func (codeauthorityroletype *CodeAuthorityRoleType) GongCopy() GongstructIF {
	newInstance := new(CodeAuthorityRoleType)
	codeauthorityroletype.GongCopyBasicFields(newInstance)
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

func (codecolourtype *CodeColourType) GongCopy() GongstructIF {
	newInstance := new(CodeColourType)
	codecolourtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongCopy() GongstructIF {
	newInstance := new(CodeCommunicationModeType)
	codecommunicationmodetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecontaminationtype *CodeContaminationType) GongCopy() GongstructIF {
	newInstance := new(CodeContaminationType)
	codecontaminationtype.GongCopyBasicFields(newInstance)
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

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongCopy() GongstructIF {
	newInstance := new(CodeEquipmentAntiCollisionType)
	codeequipmentanticollisiontype.GongCopyBasicFields(newInstance)
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

func (codegradeseparationtype *CodeGradeSeparationType) GongCopy() GongstructIF {
	newInstance := new(CodeGradeSeparationType)
	codegradeseparationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeguidancelinetype *CodeGuidanceLineType) GongCopy() GongstructIF {
	newInstance := new(CodeGuidanceLineType)
	codeguidancelinetype.GongCopyBasicFields(newInstance)
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

func (codeiatatype *CodeIATAType) GongCopy() GongstructIF {
	newInstance := new(CodeIATAType)
	codeiatatype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeicaotype *CodeICAOType) GongCopy() GongstructIF {
	newInstance := new(CodeICAOType)
	codeicaotype.GongCopyBasicFields(newInstance)
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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongCopy() GongstructIF {
	newInstance := new(CodeLogicalOperatorType)
	codelogicaloperatortype.GongCopyBasicFields(newInstance)
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

func (codemilitarystatustype *CodeMilitaryStatusType) GongCopy() GongstructIF {
	newInstance := new(CodeMilitaryStatusType)
	codemilitarystatustype.GongCopyBasicFields(newInstance)
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

func (codenotepurposetype *CodeNotePurposeType) GongCopy() GongstructIF {
	newInstance := new(CodeNotePurposeType)
	codenotepurposetype.GongCopyBasicFields(newInstance)
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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongCopy() GongstructIF {
	newInstance := new(CodeProtectAreaSectionType)
	codeprotectareasectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codervrreadingtype *CodeRVRReadingType) GongCopy() GongstructIF {
	newInstance := new(CodeRVRReadingType)
	codervrreadingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codervsmtype *CodeRVSMType) GongCopy() GongstructIF {
	newInstance := new(CodeRVSMType)
	codervsmtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderoadtype *CodeRoadType) GongCopy() GongstructIF {
	newInstance := new(CodeRoadType)
	coderoadtype.GongCopyBasicFields(newInstance)
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

func (codesidetype *CodeSideType) GongCopy() GongstructIF {
	newInstance := new(CodeSideType)
	codesidetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusairporttype *CodeStatusAirportType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusAirportType)
	codestatusairporttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusoperationstype *CodeStatusOperationsType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusOperationsType)
	codestatusoperationstype.GongCopyBasicFields(newInstance)
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

func (codetranspondertype *CodeTransponderType) GongCopy() GongstructIF {
	newInstance := new(CodeTransponderType)
	codetranspondertype.GongCopyBasicFields(newInstance)
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

func (timetype *TimeType) GongCopy() GongstructIF {
	newInstance := new(TimeType)
	timetype.GongCopyBasicFields(newInstance)
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

func (valfrictiontype *ValFrictionType) GongCopy() GongstructIF {
	newInstance := new(ValFrictionType)
	valfrictiontype.GongCopyBasicFields(newInstance)
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

// insertion point per named struct
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftcategorytype)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftenginenumbertype)
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftenginetype)
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

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachguidancetype)
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachlightingicaotype)
}

func (codeapproachlightingtype *CodeApproachLightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachlightingtype)
}

func (codeapronelementtype *CodeApronElementType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapronelementtype)
}

func (codeapronsectiontype *CodeApronSectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapronsectiontype)
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codearrestinggearenergyabsorbtype)
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codearrestinggearengagedevicetype)
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeauthorityroletype)
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codebuoydesignatortype)
}

func (codebuoytype *CodeBuoyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codebuoytype)
}

func (codecolourtype *CodeColourType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecolourtype)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecommunicationmodetype)
}

func (codecontaminationtype *CodeContaminationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecontaminationtype)
}

func (codedaytype *CodeDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedaytype)
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedeclareddistancetype)
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedirectionturntype)
}

func (codedirectiontype *CodeDirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedirectiontype)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeequipmentanticollisiontype)
}

func (codeflightorigintype *CodeFlightOriginType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightorigintype)
}

func (codeflightpurposetype *CodeFlightPurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightpurposetype)
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefrictiondevicetype)
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefrictionestimatetype)
}

func (codegradeseparationtype *CodeGradeSeparationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codegradeseparationtype)
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeguidancelinetype)
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codehelicopterperformancetype)
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeholdingcategorytype)
}

func (codeiatatype *CodeIATAType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeiatatype)
}

func (codeicaotype *CodeICAOType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeicaotype)
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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelogicaloperatortype)
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

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemilitarystatustype)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationequipmenttype)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationspecificationtype)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenotepurposetype)
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeoperationairportheliporttype)
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeoperationmanoeuvringareatype)
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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeprotectareasectiontype)
}

func (codervrreadingtype *CodeRVRReadingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervrreadingtype)
}

func (codervsmtype *CodeRVSMType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmtype)
}

func (coderoadtype *CodeRoadType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderoadtype)
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

func (codesidetype *CodeSideType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesidetype)
}

func (codestatusairporttype *CodeStatusAirportType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusairporttype)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusoperationstype)
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

func (codetranspondertype *CodeTransponderType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetranspondertype)
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeusagelimitationtype)
}

func (codevasistype *CodeVASISType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codevasistype)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codevalueinterpretationtype)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticaldatumtype)
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

func (timetype *TimeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timetype)
}

func (valangletype *ValAngleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valangletype)
}

func (valbearingtype *ValBearingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valbearingtype)
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

func (valfrictiontype *ValFrictionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valfrictiontype)
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
	__gong__computeReferencePass1(stage, stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_reference, &stage.CodeAircraftCategoryTypes_referenceOrder, &stage.CodeAircraftCategoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_reference, &stage.CodeAircraftEngineNumberTypes_referenceOrder, &stage.CodeAircraftEngineNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_reference, &stage.CodeAircraftEngineTypes_referenceOrder, &stage.CodeAircraftEngineTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_reference, &stage.CodeAircraftICAOTypes_referenceOrder, &stage.CodeAircraftICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftStandTypes, &stage.CodeAircraftStandTypes_reference, &stage.CodeAircraftStandTypes_referenceOrder, &stage.CodeAircraftStandTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftTypes, &stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_referenceOrder, &stage.CodeAircraftTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_referenceOrder, &stage.CodeAircraftWingspanClassTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirportHeliportCollocationTypes, &stage.CodeAirportHeliportCollocationTypes_reference, &stage.CodeAirportHeliportCollocationTypes_referenceOrder, &stage.CodeAirportHeliportCollocationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirportHeliportDesignatorTypes, &stage.CodeAirportHeliportDesignatorTypes_reference, &stage.CodeAirportHeliportDesignatorTypes_referenceOrder, &stage.CodeAirportHeliportDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirportHeliportTypes, &stage.CodeAirportHeliportTypes_reference, &stage.CodeAirportHeliportTypes_referenceOrder, &stage.CodeAirportHeliportTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirportWarningTypes, &stage.CodeAirportWarningTypes_reference, &stage.CodeAirportWarningTypes_referenceOrder, &stage.CodeAirportWarningTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachGuidanceTypes, &stage.CodeApproachGuidanceTypes_reference, &stage.CodeApproachGuidanceTypes_referenceOrder, &stage.CodeApproachGuidanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachLightingICAOTypes, &stage.CodeApproachLightingICAOTypes_reference, &stage.CodeApproachLightingICAOTypes_referenceOrder, &stage.CodeApproachLightingICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachLightingTypes, &stage.CodeApproachLightingTypes_reference, &stage.CodeApproachLightingTypes_referenceOrder, &stage.CodeApproachLightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApronElementTypes, &stage.CodeApronElementTypes_reference, &stage.CodeApronElementTypes_referenceOrder, &stage.CodeApronElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApronSectionTypes, &stage.CodeApronSectionTypes_reference, &stage.CodeApronSectionTypes_referenceOrder, &stage.CodeApronSectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeArrestingGearEnergyAbsorbTypes, &stage.CodeArrestingGearEnergyAbsorbTypes_reference, &stage.CodeArrestingGearEnergyAbsorbTypes_referenceOrder, &stage.CodeArrestingGearEnergyAbsorbTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeArrestingGearEngageDeviceTypes, &stage.CodeArrestingGearEngageDeviceTypes_reference, &stage.CodeArrestingGearEngageDeviceTypes_referenceOrder, &stage.CodeArrestingGearEngageDeviceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAuthorityRoleTypes, &stage.CodeAuthorityRoleTypes_reference, &stage.CodeAuthorityRoleTypes_referenceOrder, &stage.CodeAuthorityRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeBuoyDesignatorTypes, &stage.CodeBuoyDesignatorTypes_reference, &stage.CodeBuoyDesignatorTypes_referenceOrder, &stage.CodeBuoyDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeBuoyTypes, &stage.CodeBuoyTypes_reference, &stage.CodeBuoyTypes_referenceOrder, &stage.CodeBuoyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeColourTypes, &stage.CodeColourTypes_reference, &stage.CodeColourTypes_referenceOrder, &stage.CodeColourTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_referenceOrder, &stage.CodeCommunicationModeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeContaminationTypes, &stage.CodeContaminationTypes_reference, &stage.CodeContaminationTypes_referenceOrder, &stage.CodeContaminationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDayTypes, &stage.CodeDayTypes_reference, &stage.CodeDayTypes_referenceOrder, &stage.CodeDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDeclaredDistanceTypes, &stage.CodeDeclaredDistanceTypes_reference, &stage.CodeDeclaredDistanceTypes_referenceOrder, &stage.CodeDeclaredDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_reference, &stage.CodeDirectionTurnTypes_referenceOrder, &stage.CodeDirectionTurnTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionTypes, &stage.CodeDirectionTypes_reference, &stage.CodeDirectionTypes_referenceOrder, &stage.CodeDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder, &stage.CodeEquipmentAntiCollisionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightOriginTypes, &stage.CodeFlightOriginTypes_reference, &stage.CodeFlightOriginTypes_referenceOrder, &stage.CodeFlightOriginTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightPurposeTypes, &stage.CodeFlightPurposeTypes_reference, &stage.CodeFlightPurposeTypes_referenceOrder, &stage.CodeFlightPurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_reference, &stage.CodeFlightRuleTypes_referenceOrder, &stage.CodeFlightRuleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightStatusTypes, &stage.CodeFlightStatusTypes_reference, &stage.CodeFlightStatusTypes_referenceOrder, &stage.CodeFlightStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightTypes, &stage.CodeFlightTypes_reference, &stage.CodeFlightTypes_referenceOrder, &stage.CodeFlightTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFrictionDeviceTypes, &stage.CodeFrictionDeviceTypes_reference, &stage.CodeFrictionDeviceTypes_referenceOrder, &stage.CodeFrictionDeviceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFrictionEstimateTypes, &stage.CodeFrictionEstimateTypes_reference, &stage.CodeFrictionEstimateTypes_referenceOrder, &stage.CodeFrictionEstimateTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeGradeSeparationTypes, &stage.CodeGradeSeparationTypes_reference, &stage.CodeGradeSeparationTypes_referenceOrder, &stage.CodeGradeSeparationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeGuidanceLineTypes, &stage.CodeGuidanceLineTypes_reference, &stage.CodeGuidanceLineTypes_referenceOrder, &stage.CodeGuidanceLineTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHelicopterPerformanceTypes, &stage.CodeHelicopterPerformanceTypes_reference, &stage.CodeHelicopterPerformanceTypes_referenceOrder, &stage.CodeHelicopterPerformanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHoldingCategoryTypes, &stage.CodeHoldingCategoryTypes_reference, &stage.CodeHoldingCategoryTypes_referenceOrder, &stage.CodeHoldingCategoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeIATATypes, &stage.CodeIATATypes_reference, &stage.CodeIATATypes_referenceOrder, &stage.CodeIATATypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeICAOTypes, &stage.CodeICAOTypes_reference, &stage.CodeICAOTypes_referenceOrder, &stage.CodeICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightHoldingPositionTypes, &stage.CodeLightHoldingPositionTypes_reference, &stage.CodeLightHoldingPositionTypes_referenceOrder, &stage.CodeLightHoldingPositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_reference, &stage.CodeLightIntensityTypes_referenceOrder, &stage.CodeLightIntensityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightSourceTypes, &stage.CodeLightSourceTypes_reference, &stage.CodeLightSourceTypes_referenceOrder, &stage.CodeLightSourceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightingJARTypes, &stage.CodeLightingJARTypes_reference, &stage.CodeLightingJARTypes_referenceOrder, &stage.CodeLightingJARTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLoadingBridgeTypes, &stage.CodeLoadingBridgeTypes_reference, &stage.CodeLoadingBridgeTypes_referenceOrder, &stage.CodeLoadingBridgeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLogicalOperatorTypes, &stage.CodeLogicalOperatorTypes_reference, &stage.CodeLogicalOperatorTypes_referenceOrder, &stage.CodeLogicalOperatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMarkingConditionTypes, &stage.CodeMarkingConditionTypes_reference, &stage.CodeMarkingConditionTypes_referenceOrder, &stage.CodeMarkingConditionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMarkingStyleTypes, &stage.CodeMarkingStyleTypes_reference, &stage.CodeMarkingStyleTypes_referenceOrder, &stage.CodeMarkingStyleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMeteoConditionsTypes, &stage.CodeMeteoConditionsTypes_reference, &stage.CodeMeteoConditionsTypes_referenceOrder, &stage.CodeMeteoConditionsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_reference, &stage.CodeMilitaryOperationsTypes_referenceOrder, &stage.CodeMilitaryOperationsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryStatusTypes, &stage.CodeMilitaryStatusTypes_reference, &stage.CodeMilitaryStatusTypes_referenceOrder, &stage.CodeMilitaryStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_referenceOrder, &stage.CodeNavigationEquipmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_referenceOrder, &stage.CodeNavigationSpecificationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_referenceOrder, &stage.CodeNotePurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOperationAirportHeliportTypes, &stage.CodeOperationAirportHeliportTypes_reference, &stage.CodeOperationAirportHeliportTypes_referenceOrder, &stage.CodeOperationAirportHeliportTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOperationManoeuvringAreaTypes, &stage.CodeOperationManoeuvringAreaTypes_reference, &stage.CodeOperationManoeuvringAreaTypes_referenceOrder, &stage.CodeOperationManoeuvringAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePCNMethodTypes, &stage.CodePCNMethodTypes_reference, &stage.CodePCNMethodTypes_referenceOrder, &stage.CodePCNMethodTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePCNPavementTypes, &stage.CodePCNPavementTypes_reference, &stage.CodePCNPavementTypes_referenceOrder, &stage.CodePCNPavementTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePCNSubgradeTypes, &stage.CodePCNSubgradeTypes_reference, &stage.CodePCNSubgradeTypes_referenceOrder, &stage.CodePCNSubgradeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePCNTyrePressureTypes, &stage.CodePCNTyrePressureTypes_reference, &stage.CodePCNTyrePressureTypes_referenceOrder, &stage.CodePCNTyrePressureTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProtectAreaSectionTypes, &stage.CodeProtectAreaSectionTypes_reference, &stage.CodeProtectAreaSectionTypes_referenceOrder, &stage.CodeProtectAreaSectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVRReadingTypes, &stage.CodeRVRReadingTypes_reference, &stage.CodeRVRReadingTypes_referenceOrder, &stage.CodeRVRReadingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMTypes, &stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_referenceOrder, &stage.CodeRVSMTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRoadTypes, &stage.CodeRoadTypes_reference, &stage.CodeRoadTypes_referenceOrder, &stage.CodeRoadTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayElementTypes, &stage.CodeRunwayElementTypes_reference, &stage.CodeRunwayElementTypes_referenceOrder, &stage.CodeRunwayElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayMarkingTypes, &stage.CodeRunwayMarkingTypes_reference, &stage.CodeRunwayMarkingTypes_referenceOrder, &stage.CodeRunwayMarkingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayPointRoleTypes, &stage.CodeRunwayPointRoleTypes_reference, &stage.CodeRunwayPointRoleTypes_referenceOrder, &stage.CodeRunwayPointRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayProtectionAreaTypes, &stage.CodeRunwayProtectionAreaTypes_reference, &stage.CodeRunwayProtectionAreaTypes_referenceOrder, &stage.CodeRunwayProtectionAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwaySectionTypes, &stage.CodeRunwaySectionTypes_reference, &stage.CodeRunwaySectionTypes_referenceOrder, &stage.CodeRunwaySectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRunwayTypes, &stage.CodeRunwayTypes_reference, &stage.CodeRunwayTypes_referenceOrder, &stage.CodeRunwayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSideTypes, &stage.CodeSideTypes_reference, &stage.CodeSideTypes_referenceOrder, &stage.CodeSideTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusAirportTypes, &stage.CodeStatusAirportTypes_reference, &stage.CodeStatusAirportTypes_referenceOrder, &stage.CodeStatusAirportTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_reference, &stage.CodeStatusOperationsTypes_referenceOrder, &stage.CodeStatusOperationsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSurfaceCompositionTypes, &stage.CodeSurfaceCompositionTypes_reference, &stage.CodeSurfaceCompositionTypes_referenceOrder, &stage.CodeSurfaceCompositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSurfaceConditionTypes, &stage.CodeSurfaceConditionTypes_reference, &stage.CodeSurfaceConditionTypes_referenceOrder, &stage.CodeSurfaceConditionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSurfacePreparationTypes, &stage.CodeSurfacePreparationTypes_reference, &stage.CodeSurfacePreparationTypes_referenceOrder, &stage.CodeSurfacePreparationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTLOFSectionTypes, &stage.CodeTLOFSectionTypes_reference, &stage.CodeTLOFSectionTypes_referenceOrder, &stage.CodeTLOFSectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTaxiwayElementTypes, &stage.CodeTaxiwayElementTypes_reference, &stage.CodeTaxiwayElementTypes_referenceOrder, &stage.CodeTaxiwayElementTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTaxiwaySectionTypes, &stage.CodeTaxiwaySectionTypes_reference, &stage.CodeTaxiwaySectionTypes_referenceOrder, &stage.CodeTaxiwaySectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTaxiwayTypes, &stage.CodeTaxiwayTypes_reference, &stage.CodeTaxiwayTypes_referenceOrder, &stage.CodeTaxiwayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_referenceOrder, &stage.CodeTelecomNetworkTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_referenceOrder, &stage.CodeTimeEventCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_referenceOrder, &stage.CodeTimeEventTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_referenceOrder, &stage.CodeTimeReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTransponderTypes, &stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_referenceOrder, &stage.CodeTransponderTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUsageLimitationTypes, &stage.CodeUsageLimitationTypes_reference, &stage.CodeUsageLimitationTypes_referenceOrder, &stage.CodeUsageLimitationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVASISTypes, &stage.CodeVASISTypes_reference, &stage.CodeVASISTypes_referenceOrder, &stage.CodeVASISTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_reference, &stage.CodeValueInterpretationTypes_referenceOrder, &stage.CodeValueInterpretationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_referenceOrder, &stage.CodeVerticalDatumTypes_instance)

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

	__gong__computeReferencePass1(stage, stage.TimeTypes, &stage.TimeTypes_reference, &stage.TimeTypes_referenceOrder, &stage.TimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValAngleTypes, &stage.ValAngleTypes_reference, &stage.ValAngleTypes_referenceOrder, &stage.ValAngleTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValBearingTypes, &stage.ValBearingTypes_reference, &stage.ValBearingTypes_referenceOrder, &stage.ValBearingTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDepthTypes, &stage.ValDepthTypes_reference, &stage.ValDepthTypes_referenceOrder, &stage.ValDepthTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_referenceOrder, &stage.ValDistanceSignedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceTypes, &stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_referenceOrder, &stage.ValDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_referenceOrder, &stage.ValDistanceVerticalTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDurationTypes, &stage.ValDurationTypes_reference, &stage.ValDurationTypes_referenceOrder, &stage.ValDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValFLTypes, &stage.ValFLTypes_reference, &stage.ValFLTypes_referenceOrder, &stage.ValFLTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValFrictionTypes, &stage.ValFrictionTypes_reference, &stage.ValFrictionTypes_referenceOrder, &stage.ValFrictionTypes_instance)

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

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftStandTypes, stage.CodeAircraftStandTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftTypes, stage.CodeAircraftTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirportHeliportCollocationTypes, stage.CodeAirportHeliportCollocationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirportHeliportDesignatorTypes, stage.CodeAirportHeliportDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirportHeliportTypes, stage.CodeAirportHeliportTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirportWarningTypes, stage.CodeAirportWarningTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachGuidanceTypes, stage.CodeApproachGuidanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachLightingICAOTypes, stage.CodeApproachLightingICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachLightingTypes, stage.CodeApproachLightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApronElementTypes, stage.CodeApronElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApronSectionTypes, stage.CodeApronSectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeArrestingGearEnergyAbsorbTypes, stage.CodeArrestingGearEnergyAbsorbTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeArrestingGearEngageDeviceTypes, stage.CodeArrestingGearEngageDeviceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAuthorityRoleTypes, stage.CodeAuthorityRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeBuoyDesignatorTypes, stage.CodeBuoyDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeBuoyTypes, stage.CodeBuoyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeColourTypes, stage.CodeColourTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeContaminationTypes, stage.CodeContaminationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDayTypes, stage.CodeDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDeclaredDistanceTypes, stage.CodeDeclaredDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionTypes, stage.CodeDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightOriginTypes, stage.CodeFlightOriginTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightPurposeTypes, stage.CodeFlightPurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightRuleTypes, stage.CodeFlightRuleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightStatusTypes, stage.CodeFlightStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightTypes, stage.CodeFlightTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFrictionDeviceTypes, stage.CodeFrictionDeviceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFrictionEstimateTypes, stage.CodeFrictionEstimateTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeGradeSeparationTypes, stage.CodeGradeSeparationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeGuidanceLineTypes, stage.CodeGuidanceLineTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHelicopterPerformanceTypes, stage.CodeHelicopterPerformanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHoldingCategoryTypes, stage.CodeHoldingCategoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeIATATypes, stage.CodeIATATypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeICAOTypes, stage.CodeICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightHoldingPositionTypes, stage.CodeLightHoldingPositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightIntensityTypes, stage.CodeLightIntensityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightSourceTypes, stage.CodeLightSourceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightingJARTypes, stage.CodeLightingJARTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLoadingBridgeTypes, stage.CodeLoadingBridgeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLogicalOperatorTypes, stage.CodeLogicalOperatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMarkingConditionTypes, stage.CodeMarkingConditionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMarkingStyleTypes, stage.CodeMarkingStyleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMeteoConditionsTypes, stage.CodeMeteoConditionsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryStatusTypes, stage.CodeMilitaryStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOperationAirportHeliportTypes, stage.CodeOperationAirportHeliportTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOperationManoeuvringAreaTypes, stage.CodeOperationManoeuvringAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePCNMethodTypes, stage.CodePCNMethodTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePCNPavementTypes, stage.CodePCNPavementTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePCNSubgradeTypes, stage.CodePCNSubgradeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePCNTyrePressureTypes, stage.CodePCNTyrePressureTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProtectAreaSectionTypes, stage.CodeProtectAreaSectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVRReadingTypes, stage.CodeRVRReadingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMTypes, stage.CodeRVSMTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRoadTypes, stage.CodeRoadTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayElementTypes, stage.CodeRunwayElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayMarkingTypes, stage.CodeRunwayMarkingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayPointRoleTypes, stage.CodeRunwayPointRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayProtectionAreaTypes, stage.CodeRunwayProtectionAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwaySectionTypes, stage.CodeRunwaySectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRunwayTypes, stage.CodeRunwayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSideTypes, stage.CodeSideTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusAirportTypes, stage.CodeStatusAirportTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSurfaceCompositionTypes, stage.CodeSurfaceCompositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSurfaceConditionTypes, stage.CodeSurfaceConditionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSurfacePreparationTypes, stage.CodeSurfacePreparationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTLOFSectionTypes, stage.CodeTLOFSectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTaxiwayElementTypes, stage.CodeTaxiwayElementTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTaxiwaySectionTypes, stage.CodeTaxiwaySectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTaxiwayTypes, stage.CodeTaxiwayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTransponderTypes, stage.CodeTransponderTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUsageLimitationTypes, stage.CodeUsageLimitationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVASISTypes, stage.CodeVASISTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumTypes_reference, stage)

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

	__gong__computeReferencePass2(stage.TimeTypes, stage.TimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValAngleTypes, stage.ValAngleTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValBearingTypes, stage.ValBearingTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDepthTypes, stage.ValDepthTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceTypes, stage.ValDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDurationTypes, stage.ValDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValFLTypes, stage.ValFLTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValFrictionTypes, stage.ValFrictionTypes_reference, stage)

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

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftCategoryType_stagedOrder, stage.CodeAircraftCategoryTypes_referenceOrder, codeaircraftcategorytype, "CodeAircraftCategoryType")
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftEngineNumberType_stagedOrder, stage.CodeAircraftEngineNumberTypes_referenceOrder, codeaircraftenginenumbertype, "CodeAircraftEngineNumberType")
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftEngineType_stagedOrder, stage.CodeAircraftEngineTypes_referenceOrder, codeaircraftenginetype, "CodeAircraftEngineType")
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

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachGuidanceType_stagedOrder, stage.CodeApproachGuidanceTypes_referenceOrder, codeapproachguidancetype, "CodeApproachGuidanceType")
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachLightingICAOType_stagedOrder, stage.CodeApproachLightingICAOTypes_referenceOrder, codeapproachlightingicaotype, "CodeApproachLightingICAOType")
}

func (codeapproachlightingtype *CodeApproachLightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachLightingType_stagedOrder, stage.CodeApproachLightingTypes_referenceOrder, codeapproachlightingtype, "CodeApproachLightingType")
}

func (codeapronelementtype *CodeApronElementType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApronElementType_stagedOrder, stage.CodeApronElementTypes_referenceOrder, codeapronelementtype, "CodeApronElementType")
}

func (codeapronsectiontype *CodeApronSectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApronSectionType_stagedOrder, stage.CodeApronSectionTypes_referenceOrder, codeapronsectiontype, "CodeApronSectionType")
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeArrestingGearEnergyAbsorbType_stagedOrder, stage.CodeArrestingGearEnergyAbsorbTypes_referenceOrder, codearrestinggearenergyabsorbtype, "CodeArrestingGearEnergyAbsorbType")
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeArrestingGearEngageDeviceType_stagedOrder, stage.CodeArrestingGearEngageDeviceTypes_referenceOrder, codearrestinggearengagedevicetype, "CodeArrestingGearEngageDeviceType")
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAuthorityRoleType_stagedOrder, stage.CodeAuthorityRoleTypes_referenceOrder, codeauthorityroletype, "CodeAuthorityRoleType")
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeBuoyDesignatorType_stagedOrder, stage.CodeBuoyDesignatorTypes_referenceOrder, codebuoydesignatortype, "CodeBuoyDesignatorType")
}

func (codebuoytype *CodeBuoyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeBuoyType_stagedOrder, stage.CodeBuoyTypes_referenceOrder, codebuoytype, "CodeBuoyType")
}

func (codecolourtype *CodeColourType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeColourType_stagedOrder, stage.CodeColourTypes_referenceOrder, codecolourtype, "CodeColourType")
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCommunicationModeType_stagedOrder, stage.CodeCommunicationModeTypes_referenceOrder, codecommunicationmodetype, "CodeCommunicationModeType")
}

func (codecontaminationtype *CodeContaminationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeContaminationType_stagedOrder, stage.CodeContaminationTypes_referenceOrder, codecontaminationtype, "CodeContaminationType")
}

func (codedaytype *CodeDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDayType_stagedOrder, stage.CodeDayTypes_referenceOrder, codedaytype, "CodeDayType")
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDeclaredDistanceType_stagedOrder, stage.CodeDeclaredDistanceTypes_referenceOrder, codedeclareddistancetype, "CodeDeclaredDistanceType")
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDirectionTurnType_stagedOrder, stage.CodeDirectionTurnTypes_referenceOrder, codedirectionturntype, "CodeDirectionTurnType")
}

func (codedirectiontype *CodeDirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDirectionType_stagedOrder, stage.CodeDirectionTypes_referenceOrder, codedirectiontype, "CodeDirectionType")
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeEquipmentAntiCollisionType_stagedOrder, stage.CodeEquipmentAntiCollisionTypes_referenceOrder, codeequipmentanticollisiontype, "CodeEquipmentAntiCollisionType")
}

func (codeflightorigintype *CodeFlightOriginType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightOriginType_stagedOrder, stage.CodeFlightOriginTypes_referenceOrder, codeflightorigintype, "CodeFlightOriginType")
}

func (codeflightpurposetype *CodeFlightPurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightPurposeType_stagedOrder, stage.CodeFlightPurposeTypes_referenceOrder, codeflightpurposetype, "CodeFlightPurposeType")
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFrictionDeviceType_stagedOrder, stage.CodeFrictionDeviceTypes_referenceOrder, codefrictiondevicetype, "CodeFrictionDeviceType")
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFrictionEstimateType_stagedOrder, stage.CodeFrictionEstimateTypes_referenceOrder, codefrictionestimatetype, "CodeFrictionEstimateType")
}

func (codegradeseparationtype *CodeGradeSeparationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeGradeSeparationType_stagedOrder, stage.CodeGradeSeparationTypes_referenceOrder, codegradeseparationtype, "CodeGradeSeparationType")
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeGuidanceLineType_stagedOrder, stage.CodeGuidanceLineTypes_referenceOrder, codeguidancelinetype, "CodeGuidanceLineType")
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHelicopterPerformanceType_stagedOrder, stage.CodeHelicopterPerformanceTypes_referenceOrder, codehelicopterperformancetype, "CodeHelicopterPerformanceType")
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHoldingCategoryType_stagedOrder, stage.CodeHoldingCategoryTypes_referenceOrder, codeholdingcategorytype, "CodeHoldingCategoryType")
}

func (codeiatatype *CodeIATAType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeIATAType_stagedOrder, stage.CodeIATATypes_referenceOrder, codeiatatype, "CodeIATAType")
}

func (codeicaotype *CodeICAOType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeICAOType_stagedOrder, stage.CodeICAOTypes_referenceOrder, codeicaotype, "CodeICAOType")
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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLogicalOperatorType_stagedOrder, stage.CodeLogicalOperatorTypes_referenceOrder, codelogicaloperatortype, "CodeLogicalOperatorType")
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

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMilitaryStatusType_stagedOrder, stage.CodeMilitaryStatusTypes_referenceOrder, codemilitarystatustype, "CodeMilitaryStatusType")
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationEquipmentType_stagedOrder, stage.CodeNavigationEquipmentTypes_referenceOrder, codenavigationequipmenttype, "CodeNavigationEquipmentType")
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationSpecificationType_stagedOrder, stage.CodeNavigationSpecificationTypes_referenceOrder, codenavigationspecificationtype, "CodeNavigationSpecificationType")
}

func (codenotepurposetype *CodeNotePurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeTypes_referenceOrder, codenotepurposetype, "CodeNotePurposeType")
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOperationAirportHeliportType_stagedOrder, stage.CodeOperationAirportHeliportTypes_referenceOrder, codeoperationairportheliporttype, "CodeOperationAirportHeliportType")
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOperationManoeuvringAreaType_stagedOrder, stage.CodeOperationManoeuvringAreaTypes_referenceOrder, codeoperationmanoeuvringareatype, "CodeOperationManoeuvringAreaType")
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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeProtectAreaSectionType_stagedOrder, stage.CodeProtectAreaSectionTypes_referenceOrder, codeprotectareasectiontype, "CodeProtectAreaSectionType")
}

func (codervrreadingtype *CodeRVRReadingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVRReadingType_stagedOrder, stage.CodeRVRReadingTypes_referenceOrder, codervrreadingtype, "CodeRVRReadingType")
}

func (codervsmtype *CodeRVSMType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMType_stagedOrder, stage.CodeRVSMTypes_referenceOrder, codervsmtype, "CodeRVSMType")
}

func (coderoadtype *CodeRoadType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRoadType_stagedOrder, stage.CodeRoadTypes_referenceOrder, coderoadtype, "CodeRoadType")
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

func (codesidetype *CodeSideType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSideType_stagedOrder, stage.CodeSideTypes_referenceOrder, codesidetype, "CodeSideType")
}

func (codestatusairporttype *CodeStatusAirportType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusAirportType_stagedOrder, stage.CodeStatusAirportTypes_referenceOrder, codestatusairporttype, "CodeStatusAirportType")
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusOperationsType_stagedOrder, stage.CodeStatusOperationsTypes_referenceOrder, codestatusoperationstype, "CodeStatusOperationsType")
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

func (codetranspondertype *CodeTransponderType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTransponderType_stagedOrder, stage.CodeTransponderTypes_referenceOrder, codetranspondertype, "CodeTransponderType")
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUsageLimitationType_stagedOrder, stage.CodeUsageLimitationTypes_referenceOrder, codeusagelimitationtype, "CodeUsageLimitationType")
}

func (codevasistype *CodeVASISType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVASISType_stagedOrder, stage.CodeVASISTypes_referenceOrder, codevasistype, "CodeVASISType")
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeValueInterpretationType_stagedOrder, stage.CodeValueInterpretationTypes_referenceOrder, codevalueinterpretationtype, "CodeValueInterpretationType")
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalDatumType_stagedOrder, stage.CodeVerticalDatumTypes_referenceOrder, codeverticaldatumtype, "CodeVerticalDatumType")
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

func (timetype *TimeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TimeType_stagedOrder, stage.TimeTypes_referenceOrder, timetype, "TimeType")
}

func (valangletype *ValAngleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValAngleType_stagedOrder, stage.ValAngleTypes_referenceOrder, valangletype, "ValAngleType")
}

func (valbearingtype *ValBearingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValBearingType_stagedOrder, stage.ValBearingTypes_referenceOrder, valbearingtype, "ValBearingType")
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

func (valfrictiontype *ValFrictionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValFrictionType_stagedOrder, stage.ValFrictionTypes_referenceOrder, valfrictiontype, "ValFrictionType")
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

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
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

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeauthorityroletype, codeauthorityroletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeauthorityroletype *CodeAuthorityRoleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeauthorityroletype.GongGetIdentifier(stage)
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

func (codecolourtype *CodeColourType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecolourtype, codecolourtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecolourtype *CodeColourType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecolourtype.GongGetIdentifier(stage)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecommunicationmodetype, codecommunicationmodetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecommunicationmodetype *CodeCommunicationModeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecommunicationmodetype.GongGetIdentifier(stage)
}

func (codecontaminationtype *CodeContaminationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecontaminationtype, codecontaminationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecontaminationtype *CodeContaminationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecontaminationtype.GongGetIdentifier(stage)
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

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeequipmentanticollisiontype, codeequipmentanticollisiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeequipmentanticollisiontype.GongGetIdentifier(stage)
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

func (codegradeseparationtype *CodeGradeSeparationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codegradeseparationtype, codegradeseparationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codegradeseparationtype *CodeGradeSeparationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codegradeseparationtype.GongGetIdentifier(stage)
}

func (codeguidancelinetype *CodeGuidanceLineType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeguidancelinetype, codeguidancelinetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeguidancelinetype *CodeGuidanceLineType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeguidancelinetype.GongGetIdentifier(stage)
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

func (codeiatatype *CodeIATAType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeiatatype, codeiatatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeiatatype *CodeIATAType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeiatatype.GongGetIdentifier(stage)
}

func (codeicaotype *CodeICAOType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeicaotype, codeicaotype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeicaotype *CodeICAOType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeicaotype.GongGetIdentifier(stage)
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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelogicaloperatortype, codelogicaloperatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelogicaloperatortype *CodeLogicalOperatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelogicaloperatortype.GongGetIdentifier(stage)
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

func (codemilitarystatustype *CodeMilitaryStatusType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemilitarystatustype, codemilitarystatustype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemilitarystatustype *CodeMilitaryStatusType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemilitarystatustype.GongGetIdentifier(stage)
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

func (codenotepurposetype *CodeNotePurposeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codenotepurposetype, codenotepurposetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codenotepurposetype *CodeNotePurposeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codenotepurposetype.GongGetIdentifier(stage)
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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeprotectareasectiontype, codeprotectareasectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeprotectareasectiontype.GongGetIdentifier(stage)
}

func (codervrreadingtype *CodeRVRReadingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codervrreadingtype, codervrreadingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codervrreadingtype *CodeRVRReadingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codervrreadingtype.GongGetIdentifier(stage)
}

func (codervsmtype *CodeRVSMType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codervsmtype, codervsmtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codervsmtype *CodeRVSMType) GongGetReferenceIdentifier(stage *Stage) string {
	return codervsmtype.GongGetIdentifier(stage)
}

func (coderoadtype *CodeRoadType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderoadtype, coderoadtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderoadtype *CodeRoadType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderoadtype.GongGetIdentifier(stage)
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

func (codesidetype *CodeSideType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesidetype, codesidetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesidetype *CodeSideType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesidetype.GongGetIdentifier(stage)
}

func (codestatusairporttype *CodeStatusAirportType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusairporttype, codestatusairporttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusairporttype *CodeStatusAirportType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusairporttype.GongGetIdentifier(stage)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusoperationstype, codestatusoperationstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusoperationstype *CodeStatusOperationsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusoperationstype.GongGetIdentifier(stage)
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

func (codetranspondertype *CodeTransponderType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetranspondertype, codetranspondertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetranspondertype *CodeTransponderType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetranspondertype.GongGetIdentifier(stage)
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

func (timetype *TimeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timetype, timetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timetype *TimeType) GongGetReferenceIdentifier(stage *Stage) string {
	return timetype.GongGetIdentifier(stage)
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

func (valfrictiontype *ValFrictionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valfrictiontype, valfrictiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valfrictiontype *ValFrictionType) GongGetReferenceIdentifier(stage *Stage) string {
	return valfrictiontype.GongGetIdentifier(stage)
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

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftcategorytype.GongGetIdentifier(stage), "CodeAircraftCategoryType", codeaircraftcategorytype.Name)
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftenginenumbertype.GongGetIdentifier(stage), "CodeAircraftEngineNumberType", codeaircraftenginenumbertype.Name)
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftenginetype.GongGetIdentifier(stage), "CodeAircraftEngineType", codeaircraftenginetype.Name)
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

func (codeapproachguidancetype *CodeApproachGuidanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachguidancetype.GongGetIdentifier(stage), "CodeApproachGuidanceType", codeapproachguidancetype.Name)
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachlightingicaotype.GongGetIdentifier(stage), "CodeApproachLightingICAOType", codeapproachlightingicaotype.Name)
}

func (codeapproachlightingtype *CodeApproachLightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachlightingtype.GongGetIdentifier(stage), "CodeApproachLightingType", codeapproachlightingtype.Name)
}

func (codeapronelementtype *CodeApronElementType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapronelementtype.GongGetIdentifier(stage), "CodeApronElementType", codeapronelementtype.Name)
}

func (codeapronsectiontype *CodeApronSectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapronsectiontype.GongGetIdentifier(stage), "CodeApronSectionType", codeapronsectiontype.Name)
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codearrestinggearenergyabsorbtype.GongGetIdentifier(stage), "CodeArrestingGearEnergyAbsorbType", codearrestinggearenergyabsorbtype.Name)
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codearrestinggearengagedevicetype.GongGetIdentifier(stage), "CodeArrestingGearEngageDeviceType", codearrestinggearengagedevicetype.Name)
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeauthorityroletype.GongGetIdentifier(stage), "CodeAuthorityRoleType", codeauthorityroletype.Name)
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codebuoydesignatortype.GongGetIdentifier(stage), "CodeBuoyDesignatorType", codebuoydesignatortype.Name)
}

func (codebuoytype *CodeBuoyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codebuoytype.GongGetIdentifier(stage), "CodeBuoyType", codebuoytype.Name)
}

func (codecolourtype *CodeColourType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecolourtype.GongGetIdentifier(stage), "CodeColourType", codecolourtype.Name)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecommunicationmodetype.GongGetIdentifier(stage), "CodeCommunicationModeType", codecommunicationmodetype.Name)
}

func (codecontaminationtype *CodeContaminationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecontaminationtype.GongGetIdentifier(stage), "CodeContaminationType", codecontaminationtype.Name)
}

func (codedaytype *CodeDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedaytype.GongGetIdentifier(stage), "CodeDayType", codedaytype.Name)
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedeclareddistancetype.GongGetIdentifier(stage), "CodeDeclaredDistanceType", codedeclareddistancetype.Name)
}

func (codedirectionturntype *CodeDirectionTurnType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedirectionturntype.GongGetIdentifier(stage), "CodeDirectionTurnType", codedirectionturntype.Name)
}

func (codedirectiontype *CodeDirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedirectiontype.GongGetIdentifier(stage), "CodeDirectionType", codedirectiontype.Name)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeequipmentanticollisiontype.GongGetIdentifier(stage), "CodeEquipmentAntiCollisionType", codeequipmentanticollisiontype.Name)
}

func (codeflightorigintype *CodeFlightOriginType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightorigintype.GongGetIdentifier(stage), "CodeFlightOriginType", codeflightorigintype.Name)
}

func (codeflightpurposetype *CodeFlightPurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightpurposetype.GongGetIdentifier(stage), "CodeFlightPurposeType", codeflightpurposetype.Name)
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefrictiondevicetype.GongGetIdentifier(stage), "CodeFrictionDeviceType", codefrictiondevicetype.Name)
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefrictionestimatetype.GongGetIdentifier(stage), "CodeFrictionEstimateType", codefrictionestimatetype.Name)
}

func (codegradeseparationtype *CodeGradeSeparationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codegradeseparationtype.GongGetIdentifier(stage), "CodeGradeSeparationType", codegradeseparationtype.Name)
}

func (codeguidancelinetype *CodeGuidanceLineType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeguidancelinetype.GongGetIdentifier(stage), "CodeGuidanceLineType", codeguidancelinetype.Name)
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codehelicopterperformancetype.GongGetIdentifier(stage), "CodeHelicopterPerformanceType", codehelicopterperformancetype.Name)
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeholdingcategorytype.GongGetIdentifier(stage), "CodeHoldingCategoryType", codeholdingcategorytype.Name)
}

func (codeiatatype *CodeIATAType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeiatatype.GongGetIdentifier(stage), "CodeIATAType", codeiatatype.Name)
}

func (codeicaotype *CodeICAOType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeicaotype.GongGetIdentifier(stage), "CodeICAOType", codeicaotype.Name)
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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelogicaloperatortype.GongGetIdentifier(stage), "CodeLogicalOperatorType", codelogicaloperatortype.Name)
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

func (codemilitarystatustype *CodeMilitaryStatusType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemilitarystatustype.GongGetIdentifier(stage), "CodeMilitaryStatusType", codemilitarystatustype.Name)
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationequipmenttype.GongGetIdentifier(stage), "CodeNavigationEquipmentType", codenavigationequipmenttype.Name)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationspecificationtype.GongGetIdentifier(stage), "CodeNavigationSpecificationType", codenavigationspecificationtype.Name)
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenotepurposetype.GongGetIdentifier(stage), "CodeNotePurposeType", codenotepurposetype.Name)
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeoperationairportheliporttype.GongGetIdentifier(stage), "CodeOperationAirportHeliportType", codeoperationairportheliporttype.Name)
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeoperationmanoeuvringareatype.GongGetIdentifier(stage), "CodeOperationManoeuvringAreaType", codeoperationmanoeuvringareatype.Name)
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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeprotectareasectiontype.GongGetIdentifier(stage), "CodeProtectAreaSectionType", codeprotectareasectiontype.Name)
}

func (codervrreadingtype *CodeRVRReadingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervrreadingtype.GongGetIdentifier(stage), "CodeRVRReadingType", codervrreadingtype.Name)
}

func (codervsmtype *CodeRVSMType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmtype.GongGetIdentifier(stage), "CodeRVSMType", codervsmtype.Name)
}

func (coderoadtype *CodeRoadType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderoadtype.GongGetIdentifier(stage), "CodeRoadType", coderoadtype.Name)
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

func (codesidetype *CodeSideType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesidetype.GongGetIdentifier(stage), "CodeSideType", codesidetype.Name)
}

func (codestatusairporttype *CodeStatusAirportType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusairporttype.GongGetIdentifier(stage), "CodeStatusAirportType", codestatusairporttype.Name)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusoperationstype.GongGetIdentifier(stage), "CodeStatusOperationsType", codestatusoperationstype.Name)
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

func (codetranspondertype *CodeTransponderType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetranspondertype.GongGetIdentifier(stage), "CodeTransponderType", codetranspondertype.Name)
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeusagelimitationtype.GongGetIdentifier(stage), "CodeUsageLimitationType", codeusagelimitationtype.Name)
}

func (codevasistype *CodeVASISType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codevasistype.GongGetIdentifier(stage), "CodeVASISType", codevasistype.Name)
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codevalueinterpretationtype.GongGetIdentifier(stage), "CodeValueInterpretationType", codevalueinterpretationtype.Name)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticaldatumtype.GongGetIdentifier(stage), "CodeVerticalDatumType", codeverticaldatumtype.Name)
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

func (timetype *TimeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timetype.GongGetIdentifier(stage), "TimeType", timetype.Name)
}

func (valangletype *ValAngleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valangletype.GongGetIdentifier(stage), "ValAngleType", valangletype.Name)
}

func (valbearingtype *ValBearingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valbearingtype.GongGetIdentifier(stage), "ValBearingType", valbearingtype.Name)
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

func (valfrictiontype *ValFrictionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valfrictiontype.GongGetIdentifier(stage), "ValFrictionType", valfrictiontype.Name)
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

// insertion point for unstaging
func (codeaircraftcategorytype *CodeAircraftCategoryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftcategorytype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftenginenumbertype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftenginetype.GongGetReferenceIdentifier(stage))
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

func (codeapproachguidancetype *CodeApproachGuidanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachguidancetype.GongGetReferenceIdentifier(stage))
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachlightingicaotype.GongGetReferenceIdentifier(stage))
}

func (codeapproachlightingtype *CodeApproachLightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachlightingtype.GongGetReferenceIdentifier(stage))
}

func (codeapronelementtype *CodeApronElementType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapronelementtype.GongGetReferenceIdentifier(stage))
}

func (codeapronsectiontype *CodeApronSectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapronsectiontype.GongGetReferenceIdentifier(stage))
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codearrestinggearenergyabsorbtype.GongGetReferenceIdentifier(stage))
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codearrestinggearengagedevicetype.GongGetReferenceIdentifier(stage))
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeauthorityroletype.GongGetReferenceIdentifier(stage))
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codebuoydesignatortype.GongGetReferenceIdentifier(stage))
}

func (codebuoytype *CodeBuoyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codebuoytype.GongGetReferenceIdentifier(stage))
}

func (codecolourtype *CodeColourType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecolourtype.GongGetReferenceIdentifier(stage))
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecommunicationmodetype.GongGetReferenceIdentifier(stage))
}

func (codecontaminationtype *CodeContaminationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecontaminationtype.GongGetReferenceIdentifier(stage))
}

func (codedaytype *CodeDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedaytype.GongGetReferenceIdentifier(stage))
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedeclareddistancetype.GongGetReferenceIdentifier(stage))
}

func (codedirectionturntype *CodeDirectionTurnType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedirectionturntype.GongGetReferenceIdentifier(stage))
}

func (codedirectiontype *CodeDirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedirectiontype.GongGetReferenceIdentifier(stage))
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeequipmentanticollisiontype.GongGetReferenceIdentifier(stage))
}

func (codeflightorigintype *CodeFlightOriginType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightorigintype.GongGetReferenceIdentifier(stage))
}

func (codeflightpurposetype *CodeFlightPurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightpurposetype.GongGetReferenceIdentifier(stage))
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefrictiondevicetype.GongGetReferenceIdentifier(stage))
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefrictionestimatetype.GongGetReferenceIdentifier(stage))
}

func (codegradeseparationtype *CodeGradeSeparationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codegradeseparationtype.GongGetReferenceIdentifier(stage))
}

func (codeguidancelinetype *CodeGuidanceLineType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeguidancelinetype.GongGetReferenceIdentifier(stage))
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codehelicopterperformancetype.GongGetReferenceIdentifier(stage))
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeholdingcategorytype.GongGetReferenceIdentifier(stage))
}

func (codeiatatype *CodeIATAType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeiatatype.GongGetReferenceIdentifier(stage))
}

func (codeicaotype *CodeICAOType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeicaotype.GongGetReferenceIdentifier(stage))
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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelogicaloperatortype.GongGetReferenceIdentifier(stage))
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

func (codemilitarystatustype *CodeMilitaryStatusType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemilitarystatustype.GongGetReferenceIdentifier(stage))
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationequipmenttype.GongGetReferenceIdentifier(stage))
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationspecificationtype.GongGetReferenceIdentifier(stage))
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenotepurposetype.GongGetReferenceIdentifier(stage))
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeoperationairportheliporttype.GongGetReferenceIdentifier(stage))
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeoperationmanoeuvringareatype.GongGetReferenceIdentifier(stage))
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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeprotectareasectiontype.GongGetReferenceIdentifier(stage))
}

func (codervrreadingtype *CodeRVRReadingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervrreadingtype.GongGetReferenceIdentifier(stage))
}

func (codervsmtype *CodeRVSMType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmtype.GongGetReferenceIdentifier(stage))
}

func (coderoadtype *CodeRoadType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderoadtype.GongGetReferenceIdentifier(stage))
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

func (codesidetype *CodeSideType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesidetype.GongGetReferenceIdentifier(stage))
}

func (codestatusairporttype *CodeStatusAirportType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusairporttype.GongGetReferenceIdentifier(stage))
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusoperationstype.GongGetReferenceIdentifier(stage))
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

func (codetranspondertype *CodeTransponderType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetranspondertype.GongGetReferenceIdentifier(stage))
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeusagelimitationtype.GongGetReferenceIdentifier(stage))
}

func (codevasistype *CodeVASISType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codevasistype.GongGetReferenceIdentifier(stage))
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codevalueinterpretationtype.GongGetReferenceIdentifier(stage))
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticaldatumtype.GongGetReferenceIdentifier(stage))
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

func (timetype *TimeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timetype.GongGetReferenceIdentifier(stage))
}

func (valangletype *ValAngleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valangletype.GongGetReferenceIdentifier(stage))
}

func (valbearingtype *ValBearingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valbearingtype.GongGetReferenceIdentifier(stage))
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

func (valfrictiontype *ValFrictionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valfrictiontype.GongGetReferenceIdentifier(stage))
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
