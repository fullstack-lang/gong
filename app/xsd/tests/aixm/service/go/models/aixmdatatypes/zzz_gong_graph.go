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
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftGroundServiceTypes[codeaircraftgroundservicetype]
	return ok
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAviationStandardsTypes[codeaviationstandardstype]
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

func (codedaytype *CodeDayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDayTypes[codedaytype]
	return ok
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFacilityRankingTypes[codefacilityrankingtype]
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

func (codeflightruletype *CodeFlightRuleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFlightRuleTypes[codeflightruletype]
	return ok
}

func (codefueltype *CodeFuelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFuelTypes[codefueltype]
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

func (codelightintensitytype *CodeLightIntensityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLightIntensityTypes[codelightintensitytype]
	return ok
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMilitaryOperationsTypes[codemilitaryoperationstype]
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

func (codeoiltype *CodeOilType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeOilTypes[codeoiltype]
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

func (codepassengerservicetype *CodePassengerServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePassengerServiceTypes[codepassengerservicetype]
	return ok
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePilotControlledLightingTypes[codepilotcontrolledlightingtype]
	return ok
}

func (codervsmtype *CodeRVSMType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRVSMTypes[codervsmtype]
	return ok
}

func (coderadioemissiontype *CodeRadioEmissionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRadioEmissionTypes[coderadioemissiontype]
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

func (codespecialdatetype *CodeSpecialDateType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpecialDateTypes[codespecialdatetype]
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

func (codesystemactivationtype *CodeSystemActivationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSystemActivationTypes[codesystemactivationtype]
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

func (codeunitdependencytype *CodeUnitDependencyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeUnitDependencyTypes[codeunitdependencytype]
	return ok
}

func (codeunittype *CodeUnitType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeUnitTypes[codeunittype]
	return ok
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalDatumTypes[codeverticaldatumtype]
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

func (dateyeartype *DateYearType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DateYearTypes[dateyeartype]
	return ok
}

func (nonumbertype *NoNumberType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoNumberTypes[nonumbertype]
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

func (valbearingtype *ValBearingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValBearingTypes[valbearingtype]
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

func (valfrequencytype *ValFrequencyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValFrequencyTypes[valfrequencytype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeaircraftgroundservicetype) {
		return
	}

	codeaircraftgroundservicetype.Stage(stage)

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

func (codedaytype *CodeDayType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedaytype) {
		return
	}

	codedaytype.Stage(stage)

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

func (codeflightruletype *CodeFlightRuleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeflightruletype) {
		return
	}

	codeflightruletype.Stage(stage)

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

func (codelightintensitytype *CodeLightIntensityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelightintensitytype) {
		return
	}

	codelightintensitytype.Stage(stage)

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

func (codeoiltype *CodeOilType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeoiltype) {
		return
	}

	codeoiltype.Stage(stage)

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

func (codervsmtype *CodeRVSMType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codervsmtype) {
		return
	}

	codervsmtype.Stage(stage)

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

func (codespecialdatetype *CodeSpecialDateType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespecialdatetype) {
		return
	}

	codespecialdatetype.Stage(stage)

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

func (codesystemactivationtype *CodeSystemActivationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesystemactivationtype) {
		return
	}

	codesystemactivationtype.Stage(stage)

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

func (codeverticaldatumtype *CodeVerticalDatumType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticaldatumtype) {
		return
	}

	codeverticaldatumtype.Stage(stage)

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

func (valbearingtype *ValBearingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valbearingtype) {
		return
	}

	valbearingtype.Stage(stage)

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

func (valfrequencytype *ValFrequencyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valfrequencytype) {
		return
	}

	valfrequencytype.Stage(stage)

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
	case *CodeAircraftGroundServiceType:
		toT := GongCopyBranchCodeAircraftGroundServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAviationStandardsType:
		toT := GongCopyBranchCodeAviationStandardsType(mapOrigCopy, fromT)
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

	case *CodeDayType:
		toT := GongCopyBranchCodeDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFacilityRankingType:
		toT := GongCopyBranchCodeFacilityRankingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFireFightingType:
		toT := GongCopyBranchCodeFireFightingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightDestinationType:
		toT := GongCopyBranchCodeFlightDestinationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFlightRuleType:
		toT := GongCopyBranchCodeFlightRuleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFuelType:
		toT := GongCopyBranchCodeFuelType(mapOrigCopy, fromT)
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

	case *CodeLightIntensityType:
		toT := GongCopyBranchCodeLightIntensityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMilitaryOperationsType:
		toT := GongCopyBranchCodeMilitaryOperationsType(mapOrigCopy, fromT)
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

	case *CodeOilType:
		toT := GongCopyBranchCodeOilType(mapOrigCopy, fromT)
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

	case *CodePassengerServiceType:
		toT := GongCopyBranchCodePassengerServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePilotControlledLightingType:
		toT := GongCopyBranchCodePilotControlledLightingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVSMType:
		toT := GongCopyBranchCodeRVSMType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRadioEmissionType:
		toT := GongCopyBranchCodeRadioEmissionType(mapOrigCopy, fromT)
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

	case *CodeSpecialDateType:
		toT := GongCopyBranchCodeSpecialDateType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusOperationsType:
		toT := GongCopyBranchCodeStatusOperationsType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusServiceType:
		toT := GongCopyBranchCodeStatusServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSystemActivationType:
		toT := GongCopyBranchCodeSystemActivationType(mapOrigCopy, fromT)
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

	case *CodeUnitDependencyType:
		toT := GongCopyBranchCodeUnitDependencyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeUnitType:
		toT := GongCopyBranchCodeUnitType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalDatumType:
		toT := GongCopyBranchCodeVerticalDatumType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeYesNoType:
		toT := GongCopyBranchCodeYesNoType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DateMonthDayType:
		toT := GongCopyBranchDateMonthDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DateYearType:
		toT := GongCopyBranchDateYearType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoNumberType:
		toT := GongCopyBranchNoNumberType(mapOrigCopy, fromT)
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

	case *ValBearingType:
		toT := GongCopyBranchValBearingType(mapOrigCopy, fromT)
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

	case *ValFrequencyType:
		toT := GongCopyBranchValFrequencyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeaircraftgroundservicetype) {
		return
	}

	codeaircraftgroundservicetype.Unstage(stage)

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

func (codedaytype *CodeDayType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedaytype) {
		return
	}

	codedaytype.Unstage(stage)

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

func (codeflightruletype *CodeFlightRuleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeflightruletype) {
		return
	}

	codeflightruletype.Unstage(stage)

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

func (codelightintensitytype *CodeLightIntensityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelightintensitytype) {
		return
	}

	codelightintensitytype.Unstage(stage)

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

func (codeoiltype *CodeOilType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeoiltype) {
		return
	}

	codeoiltype.Unstage(stage)

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

func (codervsmtype *CodeRVSMType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codervsmtype) {
		return
	}

	codervsmtype.Unstage(stage)

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

func (codespecialdatetype *CodeSpecialDateType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespecialdatetype) {
		return
	}

	codespecialdatetype.Unstage(stage)

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

func (codesystemactivationtype *CodeSystemActivationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesystemactivationtype) {
		return
	}

	codesystemactivationtype.Unstage(stage)

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

func (codeverticaldatumtype *CodeVerticalDatumType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticaldatumtype) {
		return
	}

	codeverticaldatumtype.Unstage(stage)

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

func (valbearingtype *ValBearingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valbearingtype) {
		return
	}

	valbearingtype.Unstage(stage)

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

func (valfrequencytype *ValFrequencyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valfrequencytype) {
		return
	}

	valfrequencytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *CodeAircraftGroundServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftGroundServiceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAviationStandardsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAviationStandardsType) {
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

func (reference *CodeDayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDayType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFacilityRankingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFacilityRankingType) {
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

func (reference *CodeFlightRuleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFlightRuleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFuelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFuelType) {
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

func (reference *CodeLightIntensityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLightIntensityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMilitaryOperationsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMilitaryOperationsType) {
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

func (reference *CodeOilType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeOilType) {
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

func (reference *CodePassengerServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePassengerServiceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePilotControlledLightingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePilotControlledLightingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRVSMType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRVSMType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRadioEmissionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRadioEmissionType) {
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

func (reference *CodeSpecialDateType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpecialDateType) {
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

func (reference *CodeSystemActivationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSystemActivationType) {
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

func (reference *CodeUnitDependencyType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeUnitDependencyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeUnitType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeUnitType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalDatumType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalDatumType) {
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

func (reference *DateYearType) GongReconstructPointersFromReferences(stage *Stage, instance *DateYearType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NoNumberType) GongReconstructPointersFromReferences(stage *Stage, instance *NoNumberType) {
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

func (reference *ValBearingType) GongReconstructPointersFromReferences(stage *Stage, instance *ValBearingType) {
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

func (reference *ValFrequencyType) GongReconstructPointersFromReferences(stage *Stage, instance *ValFrequencyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *CodeAircraftGroundServiceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAviationStandardsType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeDayType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFacilityRankingType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeFlightRuleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFuelType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeLightIntensityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMilitaryOperationsType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeOilType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodePassengerServiceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePilotControlledLightingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRVSMType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRadioEmissionType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeSpecialDateType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeSystemActivationType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeUnitDependencyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeUnitType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalDatumType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *DateYearType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NoNumberType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *ValBearingType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *ValFrequencyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
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
