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

func (codeapronelementtype *CodeApronElementType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApronElementTypes[codeapronelementtype]
	return ok
}

func (codeapronsectiontype *CodeApronSectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeApronSectionTypes[codeapronsectiontype]
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

func (codeauthorityroletype *CodeAuthorityRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAuthorityRoleTypes[codeauthorityroletype]
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

func (codecolourtype *CodeColourType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeColourTypes[codecolourtype]
	return ok
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCommunicationModeTypes[codecommunicationmodetype]
	return ok
}

func (codecontaminationtype *CodeContaminationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeContaminationTypes[codecontaminationtype]
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

func (codedirectionturntype *CodeDirectionTurnType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDirectionTurnTypes[codedirectionturntype]
	return ok
}

func (codedirectiontype *CodeDirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDirectionTypes[codedirectiontype]
	return ok
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeEquipmentAntiCollisionTypes[codeequipmentanticollisiontype]
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFrictionDeviceTypes[codefrictiondevicetype]
	return ok
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFrictionEstimateTypes[codefrictionestimatetype]
	return ok
}

func (codegradeseparationtype *CodeGradeSeparationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeGradeSeparationTypes[codegradeseparationtype]
	return ok
}

func (codeguidancelinetype *CodeGuidanceLineType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeGuidanceLineTypes[codeguidancelinetype]
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

func (codeiatatype *CodeIATAType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeIATATypes[codeiatatype]
	return ok
}

func (codeicaotype *CodeICAOType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeICAOTypes[codeicaotype]
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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLogicalOperatorTypes[codelogicaloperatortype]
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

func (codemilitarystatustype *CodeMilitaryStatusType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMilitaryStatusTypes[codemilitarystatustype]
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

func (codenotepurposetype *CodeNotePurposeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNotePurposeTypes[codenotepurposetype]
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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeProtectAreaSectionTypes[codeprotectareasectiontype]
	return ok
}

func (codervrreadingtype *CodeRVRReadingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRVRReadingTypes[codervrreadingtype]
	return ok
}

func (codervsmtype *CodeRVSMType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRVSMTypes[codervsmtype]
	return ok
}

func (coderoadtype *CodeRoadType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRoadTypes[coderoadtype]
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

func (codesidetype *CodeSideType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSideTypes[codesidetype]
	return ok
}

func (codestatusairporttype *CodeStatusAirportType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusAirportTypes[codestatusairporttype]
	return ok
}

func (codestatusoperationstype *CodeStatusOperationsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusOperationsTypes[codestatusoperationstype]
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

func (codetranspondertype *CodeTransponderType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTransponderTypes[codetranspondertype]
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

func (codevalueinterpretationtype *CodeValueInterpretationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeValueInterpretationTypes[codevalueinterpretationtype]
	return ok
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalDatumTypes[codeverticaldatumtype]
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

func (timetype *TimeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimeTypes[timetype]
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

func (valfrictiontype *ValFrictionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValFrictionTypes[valfrictiontype]
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

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
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

func (codeauthorityroletype *CodeAuthorityRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeauthorityroletype) {
		return
	}

	codeauthorityroletype.Stage(stage)

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

func (codecolourtype *CodeColourType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecolourtype) {
		return
	}

	codecolourtype.Stage(stage)

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

func (codecontaminationtype *CodeContaminationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecontaminationtype) {
		return
	}

	codecontaminationtype.Stage(stage)

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

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeequipmentanticollisiontype) {
		return
	}

	codeequipmentanticollisiontype.Stage(stage)

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

func (codegradeseparationtype *CodeGradeSeparationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codegradeseparationtype) {
		return
	}

	codegradeseparationtype.Stage(stage)

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

func (codeiatatype *CodeIATAType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeiatatype) {
		return
	}

	codeiatatype.Stage(stage)

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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelogicaloperatortype) {
		return
	}

	codelogicaloperatortype.Stage(stage)

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

func (codemilitarystatustype *CodeMilitaryStatusType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemilitarystatustype) {
		return
	}

	codemilitarystatustype.Stage(stage)

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

func (codenotepurposetype *CodeNotePurposeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenotepurposetype) {
		return
	}

	codenotepurposetype.Stage(stage)

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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeprotectareasectiontype) {
		return
	}

	codeprotectareasectiontype.Stage(stage)

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

func (codervsmtype *CodeRVSMType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codervsmtype) {
		return
	}

	codervsmtype.Stage(stage)

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

func (codesidetype *CodeSideType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesidetype) {
		return
	}

	codesidetype.Stage(stage)

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

func (codestatusoperationstype *CodeStatusOperationsType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusoperationstype) {
		return
	}

	codestatusoperationstype.Stage(stage)

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

func (codetranspondertype *CodeTransponderType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetranspondertype) {
		return
	}

	codetranspondertype.Stage(stage)

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

func (timetype *TimeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timetype) {
		return
	}

	timetype.Stage(stage)

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

func (valfrictiontype *ValFrictionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valfrictiontype) {
		return
	}

	valfrictiontype.Stage(stage)

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

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *CodeAircraftCategoryType:
		toT := GongCopyBranchCodeAircraftCategoryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftEngineNumberType:
		toT := GongCopyBranchCodeAircraftEngineNumberType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftEngineType:
		toT := GongCopyBranchCodeAircraftEngineType(mapOrigCopy, fromT)
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

	case *CodeApproachGuidanceType:
		toT := GongCopyBranchCodeApproachGuidanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachLightingICAOType:
		toT := GongCopyBranchCodeApproachLightingICAOType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachLightingType:
		toT := GongCopyBranchCodeApproachLightingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApronElementType:
		toT := GongCopyBranchCodeApronElementType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApronSectionType:
		toT := GongCopyBranchCodeApronSectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeArrestingGearEnergyAbsorbType:
		toT := GongCopyBranchCodeArrestingGearEnergyAbsorbType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeArrestingGearEngageDeviceType:
		toT := GongCopyBranchCodeArrestingGearEngageDeviceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAuthorityRoleType:
		toT := GongCopyBranchCodeAuthorityRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeBuoyDesignatorType:
		toT := GongCopyBranchCodeBuoyDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeBuoyType:
		toT := GongCopyBranchCodeBuoyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeColourType:
		toT := GongCopyBranchCodeColourType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCommunicationModeType:
		toT := GongCopyBranchCodeCommunicationModeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeContaminationType:
		toT := GongCopyBranchCodeContaminationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDayType:
		toT := GongCopyBranchCodeDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDeclaredDistanceType:
		toT := GongCopyBranchCodeDeclaredDistanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDirectionTurnType:
		toT := GongCopyBranchCodeDirectionTurnType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDirectionType:
		toT := GongCopyBranchCodeDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeEquipmentAntiCollisionType:
		toT := GongCopyBranchCodeEquipmentAntiCollisionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightOriginType:
		toT := GongCopyBranchCodeFlightOriginType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightPurposeType:
		toT := GongCopyBranchCodeFlightPurposeType(mapOrigCopy, fromT)
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

	case *CodeFrictionDeviceType:
		toT := GongCopyBranchCodeFrictionDeviceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFrictionEstimateType:
		toT := GongCopyBranchCodeFrictionEstimateType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeGradeSeparationType:
		toT := GongCopyBranchCodeGradeSeparationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeGuidanceLineType:
		toT := GongCopyBranchCodeGuidanceLineType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHelicopterPerformanceType:
		toT := GongCopyBranchCodeHelicopterPerformanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHoldingCategoryType:
		toT := GongCopyBranchCodeHoldingCategoryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeIATAType:
		toT := GongCopyBranchCodeIATAType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeICAOType:
		toT := GongCopyBranchCodeICAOType(mapOrigCopy, fromT)
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

	case *CodeLogicalOperatorType:
		toT := GongCopyBranchCodeLogicalOperatorType(mapOrigCopy, fromT)
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

	case *CodeMilitaryStatusType:
		toT := GongCopyBranchCodeMilitaryStatusType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavigationEquipmentType:
		toT := GongCopyBranchCodeNavigationEquipmentType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavigationSpecificationType:
		toT := GongCopyBranchCodeNavigationSpecificationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNotePurposeType:
		toT := GongCopyBranchCodeNotePurposeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOperationAirportHeliportType:
		toT := GongCopyBranchCodeOperationAirportHeliportType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeOperationManoeuvringAreaType:
		toT := GongCopyBranchCodeOperationManoeuvringAreaType(mapOrigCopy, fromT)
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

	case *CodeProtectAreaSectionType:
		toT := GongCopyBranchCodeProtectAreaSectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVRReadingType:
		toT := GongCopyBranchCodeRVRReadingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVSMType:
		toT := GongCopyBranchCodeRVSMType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRoadType:
		toT := GongCopyBranchCodeRoadType(mapOrigCopy, fromT)
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

	case *CodeSideType:
		toT := GongCopyBranchCodeSideType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusAirportType:
		toT := GongCopyBranchCodeStatusAirportType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusOperationsType:
		toT := GongCopyBranchCodeStatusOperationsType(mapOrigCopy, fromT)
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

	case *CodeTransponderType:
		toT := GongCopyBranchCodeTransponderType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeUsageLimitationType:
		toT := GongCopyBranchCodeUsageLimitationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVASISType:
		toT := GongCopyBranchCodeVASISType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeValueInterpretationType:
		toT := GongCopyBranchCodeValueInterpretationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalDatumType:
		toT := GongCopyBranchCodeVerticalDatumType(mapOrigCopy, fromT)
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

	case *TimeType:
		toT := GongCopyBranchTimeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValAngleType:
		toT := GongCopyBranchValAngleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValBearingType:
		toT := GongCopyBranchValBearingType(mapOrigCopy, fromT)
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

	case *ValFrictionType:
		toT := GongCopyBranchValFrictionType(mapOrigCopy, fromT)
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

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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

func (codeauthorityroletype *CodeAuthorityRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeauthorityroletype) {
		return
	}

	codeauthorityroletype.Unstage(stage)

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

func (codecolourtype *CodeColourType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecolourtype) {
		return
	}

	codecolourtype.Unstage(stage)

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

func (codecontaminationtype *CodeContaminationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecontaminationtype) {
		return
	}

	codecontaminationtype.Unstage(stage)

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

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeequipmentanticollisiontype) {
		return
	}

	codeequipmentanticollisiontype.Unstage(stage)

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

func (codegradeseparationtype *CodeGradeSeparationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codegradeseparationtype) {
		return
	}

	codegradeseparationtype.Unstage(stage)

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

func (codeiatatype *CodeIATAType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeiatatype) {
		return
	}

	codeiatatype.Unstage(stage)

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

func (codelogicaloperatortype *CodeLogicalOperatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelogicaloperatortype) {
		return
	}

	codelogicaloperatortype.Unstage(stage)

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

func (codemilitarystatustype *CodeMilitaryStatusType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemilitarystatustype) {
		return
	}

	codemilitarystatustype.Unstage(stage)

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

func (codenotepurposetype *CodeNotePurposeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenotepurposetype) {
		return
	}

	codenotepurposetype.Unstage(stage)

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

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeprotectareasectiontype) {
		return
	}

	codeprotectareasectiontype.Unstage(stage)

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

func (codervsmtype *CodeRVSMType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codervsmtype) {
		return
	}

	codervsmtype.Unstage(stage)

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

func (codesidetype *CodeSideType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesidetype) {
		return
	}

	codesidetype.Unstage(stage)

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

func (codestatusoperationstype *CodeStatusOperationsType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusoperationstype) {
		return
	}

	codestatusoperationstype.Unstage(stage)

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

func (codetranspondertype *CodeTransponderType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetranspondertype) {
		return
	}

	codetranspondertype.Unstage(stage)

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

func (timetype *TimeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timetype) {
		return
	}

	timetype.Unstage(stage)

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

func (valfrictiontype *ValFrictionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valfrictiontype) {
		return
	}

	valfrictiontype.Unstage(stage)

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

// insertion point for pointer reconstruction from references
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

func (reference *CodeApronElementType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApronElementType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeApronSectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeApronSectionType) {
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

func (reference *CodeAuthorityRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAuthorityRoleType) {
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

func (reference *CodeColourType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeColourType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCommunicationModeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCommunicationModeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeContaminationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeContaminationType) {
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

func (reference *CodeDirectionTurnType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDirectionTurnType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDirectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeEquipmentAntiCollisionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeEquipmentAntiCollisionType) {
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

func (reference *CodeFrictionDeviceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFrictionDeviceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFrictionEstimateType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFrictionEstimateType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeGradeSeparationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeGradeSeparationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeGuidanceLineType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeGuidanceLineType) {
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

func (reference *CodeIATAType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeIATAType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeICAOType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeICAOType) {
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

func (reference *CodeLogicalOperatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLogicalOperatorType) {
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

func (reference *CodeMilitaryStatusType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMilitaryStatusType) {
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

func (reference *CodeNotePurposeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNotePurposeType) {
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

func (reference *CodeProtectAreaSectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeProtectAreaSectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRVRReadingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRVRReadingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRVSMType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRVSMType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRoadType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRoadType) {
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

func (reference *CodeSideType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSideType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusAirportType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusAirportType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusOperationsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusOperationsType) {
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

func (reference *CodeTransponderType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTransponderType) {
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

func (reference *CodeValueInterpretationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeValueInterpretationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalDatumType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalDatumType) {
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

func (reference *TimeType) GongReconstructPointersFromReferences(stage *Stage, instance *TimeType) {
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

func (reference *ValFrictionType) GongReconstructPointersFromReferences(stage *Stage, instance *ValFrictionType) {
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

// insertion point for pointer reconstruction from instances
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

func (reference *CodeApronElementType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeApronSectionType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeAuthorityRoleType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeColourType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCommunicationModeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeContaminationType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeDirectionTurnType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDirectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeEquipmentAntiCollisionType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeFrictionDeviceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFrictionEstimateType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeGradeSeparationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeGuidanceLineType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeIATAType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeICAOType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeLogicalOperatorType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeMilitaryStatusType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeNotePurposeType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeProtectAreaSectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRVRReadingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRVSMType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRoadType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeSideType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusAirportType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusOperationsType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeTransponderType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeValueInterpretationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalDatumType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *TimeType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *ValFrictionType) GongReconstructPointersFromInstances(stage *Stage) {
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

// insertion point for diff per struct
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
