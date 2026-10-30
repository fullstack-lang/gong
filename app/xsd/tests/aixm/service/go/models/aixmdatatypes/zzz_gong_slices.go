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
	res = __gong__appendInstances(res, stage.CodeAircraftGroundServiceTypes)

	res = __gong__appendInstances(res, stage.CodeAviationStandardsTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationChannelTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationModeTypes)

	res = __gong__appendInstances(res, stage.CodeDayTypes)

	res = __gong__appendInstances(res, stage.CodeFacilityRankingTypes)

	res = __gong__appendInstances(res, stage.CodeFireFightingTypes)

	res = __gong__appendInstances(res, stage.CodeFlightDestinationTypes)

	res = __gong__appendInstances(res, stage.CodeFlightRuleTypes)

	res = __gong__appendInstances(res, stage.CodeFuelTypes)

	res = __gong__appendInstances(res, stage.CodeIntensityStandByTypes)

	res = __gong__appendInstances(res, stage.CodeLanguageTypes)

	res = __gong__appendInstances(res, stage.CodeLevelSeriesTypes)

	res = __gong__appendInstances(res, stage.CodeLevelTableDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeLightIntensityTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryOperationsTypes)

	res = __gong__appendInstances(res, stage.CodeNitrogenTypes)

	res = __gong__appendInstances(res, stage.CodeNorthReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeNotePurposeTypes)

	res = __gong__appendInstances(res, stage.CodeOilTypes)

	res = __gong__appendInstances(res, stage.CodeOrganisationDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeOrganisationHierarchyTypes)

	res = __gong__appendInstances(res, stage.CodeOrganisationTypes)

	res = __gong__appendInstances(res, stage.CodeOxygenTypes)

	res = __gong__appendInstances(res, stage.CodePassengerServiceTypes)

	res = __gong__appendInstances(res, stage.CodePilotControlledLightingTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMTypes)

	res = __gong__appendInstances(res, stage.CodeRadioEmissionTypes)

	res = __gong__appendInstances(res, stage.CodeServiceATCTypes)

	res = __gong__appendInstances(res, stage.CodeServiceATFMTypes)

	res = __gong__appendInstances(res, stage.CodeServiceGroundControlTypes)

	res = __gong__appendInstances(res, stage.CodeServiceInformationTypes)

	res = __gong__appendInstances(res, stage.CodeServiceSARTypes)

	res = __gong__appendInstances(res, stage.CodeSpecialDateTypes)

	res = __gong__appendInstances(res, stage.CodeStatusOperationsTypes)

	res = __gong__appendInstances(res, stage.CodeStatusServiceTypes)

	res = __gong__appendInstances(res, stage.CodeSystemActivationTypes)

	res = __gong__appendInstances(res, stage.CodeTelecomNetworkTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventCombinationTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventTypes)

	res = __gong__appendInstances(res, stage.CodeTimeReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeUnitDependencyTypes)

	res = __gong__appendInstances(res, stage.CodeUnitTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalDatumTypes)

	res = __gong__appendInstances(res, stage.CodeYesNoTypes)

	res = __gong__appendInstances(res, stage.DateMonthDayTypes)

	res = __gong__appendInstances(res, stage.DateYearTypes)

	res = __gong__appendInstances(res, stage.NoNumberTypes)

	res = __gong__appendInstances(res, stage.TextAddressTypes)

	res = __gong__appendInstances(res, stage.TextDesignatorTypes)

	res = __gong__appendInstances(res, stage.TextInstructionTypes)

	res = __gong__appendInstances(res, stage.TextNameTypes)

	res = __gong__appendInstances(res, stage.TextNoteTypes)

	res = __gong__appendInstances(res, stage.TextPhoneTypes)

	res = __gong__appendInstances(res, stage.TextPropertyNameTypes)

	res = __gong__appendInstances(res, stage.TimeTypes)

	res = __gong__appendInstances(res, stage.ValBearingTypes)

	res = __gong__appendInstances(res, stage.ValDistanceSignedTypes)

	res = __gong__appendInstances(res, stage.ValDistanceTypes)

	res = __gong__appendInstances(res, stage.ValDistanceVerticalTypes)

	res = __gong__appendInstances(res, stage.ValDurationTypes)

	res = __gong__appendInstances(res, stage.ValFrequencyTypes)

	return
}

// insertion point per named struct
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftGroundServiceType)
	codeaircraftgroundservicetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongCopy() GongstructIF {
	newInstance := new(CodeAviationStandardsType)
	codeaviationstandardstype.GongCopyBasicFields(newInstance)
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

func (codedaytype *CodeDayType) GongCopy() GongstructIF {
	newInstance := new(CodeDayType)
	codedaytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongCopy() GongstructIF {
	newInstance := new(CodeFacilityRankingType)
	codefacilityrankingtype.GongCopyBasicFields(newInstance)
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

func (codeflightruletype *CodeFlightRuleType) GongCopy() GongstructIF {
	newInstance := new(CodeFlightRuleType)
	codeflightruletype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefueltype *CodeFuelType) GongCopy() GongstructIF {
	newInstance := new(CodeFuelType)
	codefueltype.GongCopyBasicFields(newInstance)
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

func (codelightintensitytype *CodeLightIntensityType) GongCopy() GongstructIF {
	newInstance := new(CodeLightIntensityType)
	codelightintensitytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongCopy() GongstructIF {
	newInstance := new(CodeMilitaryOperationsType)
	codemilitaryoperationstype.GongCopyBasicFields(newInstance)
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

func (codeoiltype *CodeOilType) GongCopy() GongstructIF {
	newInstance := new(CodeOilType)
	codeoiltype.GongCopyBasicFields(newInstance)
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

func (codervsmtype *CodeRVSMType) GongCopy() GongstructIF {
	newInstance := new(CodeRVSMType)
	codervsmtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderadioemissiontype *CodeRadioEmissionType) GongCopy() GongstructIF {
	newInstance := new(CodeRadioEmissionType)
	coderadioemissiontype.GongCopyBasicFields(newInstance)
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

func (codespecialdatetype *CodeSpecialDateType) GongCopy() GongstructIF {
	newInstance := new(CodeSpecialDateType)
	codespecialdatetype.GongCopyBasicFields(newInstance)
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

func (codesystemactivationtype *CodeSystemActivationType) GongCopy() GongstructIF {
	newInstance := new(CodeSystemActivationType)
	codesystemactivationtype.GongCopyBasicFields(newInstance)
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

func (codeverticaldatumtype *CodeVerticalDatumType) GongCopy() GongstructIF {
	newInstance := new(CodeVerticalDatumType)
	codeverticaldatumtype.GongCopyBasicFields(newInstance)
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

func (valbearingtype *ValBearingType) GongCopy() GongstructIF {
	newInstance := new(ValBearingType)
	valbearingtype.GongCopyBasicFields(newInstance)
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

func (valfrequencytype *ValFrequencyType) GongCopy() GongstructIF {
	newInstance := new(ValFrequencyType)
	valfrequencytype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftgroundservicetype)
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaviationstandardstype)
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

func (codedaytype *CodeDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedaytype)
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefacilityrankingtype)
}

func (codefirefightingtype *CodeFireFightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefirefightingtype)
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightdestinationtype)
}

func (codeflightruletype *CodeFlightRuleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeflightruletype)
}

func (codefueltype *CodeFuelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefueltype)
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

func (codelightintensitytype *CodeLightIntensityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelightintensitytype)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemilitaryoperationstype)
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

func (codeoiltype *CodeOilType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeoiltype)
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

func (codepassengerservicetype *CodePassengerServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepassengerservicetype)
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepilotcontrolledlightingtype)
}

func (codervsmtype *CodeRVSMType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmtype)
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderadioemissiontype)
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

func (codespecialdatetype *CodeSpecialDateType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codespecialdatetype)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusoperationstype)
}

func (codestatusservicetype *CodeStatusServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusservicetype)
}

func (codesystemactivationtype *CodeSystemActivationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesystemactivationtype)
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

func (codeunitdependencytype *CodeUnitDependencyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeunitdependencytype)
}

func (codeunittype *CodeUnitType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeunittype)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticaldatumtype)
}

func (codeyesnotype *CodeYesNoType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeyesnotype)
}

func (datemonthdaytype *DateMonthDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datemonthdaytype)
}

func (dateyeartype *DateYearType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dateyeartype)
}

func (nonumbertype *NoNumberType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nonumbertype)
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

func (valbearingtype *ValBearingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valbearingtype)
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

func (valfrequencytype *ValFrequencyType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valfrequencytype)
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
	__gong__computeReferencePass1(stage, stage.CodeAircraftGroundServiceTypes, &stage.CodeAircraftGroundServiceTypes_reference, &stage.CodeAircraftGroundServiceTypes_referenceOrder, &stage.CodeAircraftGroundServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAviationStandardsTypes, &stage.CodeAviationStandardsTypes_reference, &stage.CodeAviationStandardsTypes_referenceOrder, &stage.CodeAviationStandardsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationChannelTypes, &stage.CodeCommunicationChannelTypes_reference, &stage.CodeCommunicationChannelTypes_referenceOrder, &stage.CodeCommunicationChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationDirectionTypes, &stage.CodeCommunicationDirectionTypes_reference, &stage.CodeCommunicationDirectionTypes_referenceOrder, &stage.CodeCommunicationDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_referenceOrder, &stage.CodeCommunicationModeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDayTypes, &stage.CodeDayTypes_reference, &stage.CodeDayTypes_referenceOrder, &stage.CodeDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFacilityRankingTypes, &stage.CodeFacilityRankingTypes_reference, &stage.CodeFacilityRankingTypes_referenceOrder, &stage.CodeFacilityRankingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFireFightingTypes, &stage.CodeFireFightingTypes_reference, &stage.CodeFireFightingTypes_referenceOrder, &stage.CodeFireFightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightDestinationTypes, &stage.CodeFlightDestinationTypes_reference, &stage.CodeFlightDestinationTypes_referenceOrder, &stage.CodeFlightDestinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_reference, &stage.CodeFlightRuleTypes_referenceOrder, &stage.CodeFlightRuleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFuelTypes, &stage.CodeFuelTypes_reference, &stage.CodeFuelTypes_referenceOrder, &stage.CodeFuelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeIntensityStandByTypes, &stage.CodeIntensityStandByTypes_reference, &stage.CodeIntensityStandByTypes_referenceOrder, &stage.CodeIntensityStandByTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLanguageTypes, &stage.CodeLanguageTypes_reference, &stage.CodeLanguageTypes_referenceOrder, &stage.CodeLanguageTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLevelSeriesTypes, &stage.CodeLevelSeriesTypes_reference, &stage.CodeLevelSeriesTypes_referenceOrder, &stage.CodeLevelSeriesTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLevelTableDesignatorTypes, &stage.CodeLevelTableDesignatorTypes_reference, &stage.CodeLevelTableDesignatorTypes_referenceOrder, &stage.CodeLevelTableDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLightIntensityTypes, &stage.CodeLightIntensityTypes_reference, &stage.CodeLightIntensityTypes_referenceOrder, &stage.CodeLightIntensityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryOperationsTypes, &stage.CodeMilitaryOperationsTypes_reference, &stage.CodeMilitaryOperationsTypes_referenceOrder, &stage.CodeMilitaryOperationsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNitrogenTypes, &stage.CodeNitrogenTypes_reference, &stage.CodeNitrogenTypes_referenceOrder, &stage.CodeNitrogenTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNorthReferenceTypes, &stage.CodeNorthReferenceTypes_reference, &stage.CodeNorthReferenceTypes_referenceOrder, &stage.CodeNorthReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_referenceOrder, &stage.CodeNotePurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOilTypes, &stage.CodeOilTypes_reference, &stage.CodeOilTypes_referenceOrder, &stage.CodeOilTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOrganisationDesignatorTypes, &stage.CodeOrganisationDesignatorTypes_reference, &stage.CodeOrganisationDesignatorTypes_referenceOrder, &stage.CodeOrganisationDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOrganisationHierarchyTypes, &stage.CodeOrganisationHierarchyTypes_reference, &stage.CodeOrganisationHierarchyTypes_referenceOrder, &stage.CodeOrganisationHierarchyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOrganisationTypes, &stage.CodeOrganisationTypes_reference, &stage.CodeOrganisationTypes_referenceOrder, &stage.CodeOrganisationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeOxygenTypes, &stage.CodeOxygenTypes_reference, &stage.CodeOxygenTypes_referenceOrder, &stage.CodeOxygenTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePassengerServiceTypes, &stage.CodePassengerServiceTypes_reference, &stage.CodePassengerServiceTypes_referenceOrder, &stage.CodePassengerServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePilotControlledLightingTypes, &stage.CodePilotControlledLightingTypes_reference, &stage.CodePilotControlledLightingTypes_referenceOrder, &stage.CodePilotControlledLightingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMTypes, &stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_referenceOrder, &stage.CodeRVSMTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRadioEmissionTypes, &stage.CodeRadioEmissionTypes_reference, &stage.CodeRadioEmissionTypes_referenceOrder, &stage.CodeRadioEmissionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceATCTypes, &stage.CodeServiceATCTypes_reference, &stage.CodeServiceATCTypes_referenceOrder, &stage.CodeServiceATCTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceATFMTypes, &stage.CodeServiceATFMTypes_reference, &stage.CodeServiceATFMTypes_referenceOrder, &stage.CodeServiceATFMTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceGroundControlTypes, &stage.CodeServiceGroundControlTypes_reference, &stage.CodeServiceGroundControlTypes_referenceOrder, &stage.CodeServiceGroundControlTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceInformationTypes, &stage.CodeServiceInformationTypes_reference, &stage.CodeServiceInformationTypes_referenceOrder, &stage.CodeServiceInformationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeServiceSARTypes, &stage.CodeServiceSARTypes_reference, &stage.CodeServiceSARTypes_referenceOrder, &stage.CodeServiceSARTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpecialDateTypes, &stage.CodeSpecialDateTypes_reference, &stage.CodeSpecialDateTypes_referenceOrder, &stage.CodeSpecialDateTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusOperationsTypes, &stage.CodeStatusOperationsTypes_reference, &stage.CodeStatusOperationsTypes_referenceOrder, &stage.CodeStatusOperationsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusServiceTypes, &stage.CodeStatusServiceTypes_reference, &stage.CodeStatusServiceTypes_referenceOrder, &stage.CodeStatusServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSystemActivationTypes, &stage.CodeSystemActivationTypes_reference, &stage.CodeSystemActivationTypes_referenceOrder, &stage.CodeSystemActivationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_referenceOrder, &stage.CodeTelecomNetworkTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_referenceOrder, &stage.CodeTimeEventCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_referenceOrder, &stage.CodeTimeEventTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_referenceOrder, &stage.CodeTimeReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUnitDependencyTypes, &stage.CodeUnitDependencyTypes_reference, &stage.CodeUnitDependencyTypes_referenceOrder, &stage.CodeUnitDependencyTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUnitTypes, &stage.CodeUnitTypes_reference, &stage.CodeUnitTypes_referenceOrder, &stage.CodeUnitTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_referenceOrder, &stage.CodeVerticalDatumTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeYesNoTypes, &stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_referenceOrder, &stage.CodeYesNoTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateMonthDayTypes, &stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_referenceOrder, &stage.DateMonthDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateYearTypes, &stage.DateYearTypes_reference, &stage.DateYearTypes_referenceOrder, &stage.DateYearTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoNumberTypes, &stage.NoNumberTypes_reference, &stage.NoNumberTypes_referenceOrder, &stage.NoNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextAddressTypes, &stage.TextAddressTypes_reference, &stage.TextAddressTypes_referenceOrder, &stage.TextAddressTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextDesignatorTypes, &stage.TextDesignatorTypes_reference, &stage.TextDesignatorTypes_referenceOrder, &stage.TextDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextInstructionTypes, &stage.TextInstructionTypes_reference, &stage.TextInstructionTypes_referenceOrder, &stage.TextInstructionTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNameTypes, &stage.TextNameTypes_reference, &stage.TextNameTypes_referenceOrder, &stage.TextNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNoteTypes, &stage.TextNoteTypes_reference, &stage.TextNoteTypes_referenceOrder, &stage.TextNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPhoneTypes, &stage.TextPhoneTypes_reference, &stage.TextPhoneTypes_referenceOrder, &stage.TextPhoneTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_referenceOrder, &stage.TextPropertyNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimeTypes, &stage.TimeTypes_reference, &stage.TimeTypes_referenceOrder, &stage.TimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValBearingTypes, &stage.ValBearingTypes_reference, &stage.ValBearingTypes_referenceOrder, &stage.ValBearingTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_referenceOrder, &stage.ValDistanceSignedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceTypes, &stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_referenceOrder, &stage.ValDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_referenceOrder, &stage.ValDistanceVerticalTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDurationTypes, &stage.ValDurationTypes_reference, &stage.ValDurationTypes_referenceOrder, &stage.ValDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValFrequencyTypes, &stage.ValFrequencyTypes_reference, &stage.ValFrequencyTypes_referenceOrder, &stage.ValFrequencyTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CodeAircraftGroundServiceTypes, stage.CodeAircraftGroundServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAviationStandardsTypes, stage.CodeAviationStandardsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationChannelTypes, stage.CodeCommunicationChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationDirectionTypes, stage.CodeCommunicationDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDayTypes, stage.CodeDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFacilityRankingTypes, stage.CodeFacilityRankingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFireFightingTypes, stage.CodeFireFightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightDestinationTypes, stage.CodeFlightDestinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightRuleTypes, stage.CodeFlightRuleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFuelTypes, stage.CodeFuelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeIntensityStandByTypes, stage.CodeIntensityStandByTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLanguageTypes, stage.CodeLanguageTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLevelSeriesTypes, stage.CodeLevelSeriesTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLevelTableDesignatorTypes, stage.CodeLevelTableDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLightIntensityTypes, stage.CodeLightIntensityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryOperationsTypes, stage.CodeMilitaryOperationsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNitrogenTypes, stage.CodeNitrogenTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNorthReferenceTypes, stage.CodeNorthReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOilTypes, stage.CodeOilTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOrganisationDesignatorTypes, stage.CodeOrganisationDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOrganisationHierarchyTypes, stage.CodeOrganisationHierarchyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOrganisationTypes, stage.CodeOrganisationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeOxygenTypes, stage.CodeOxygenTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePassengerServiceTypes, stage.CodePassengerServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePilotControlledLightingTypes, stage.CodePilotControlledLightingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMTypes, stage.CodeRVSMTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRadioEmissionTypes, stage.CodeRadioEmissionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceATCTypes, stage.CodeServiceATCTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceATFMTypes, stage.CodeServiceATFMTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceGroundControlTypes, stage.CodeServiceGroundControlTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceInformationTypes, stage.CodeServiceInformationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeServiceSARTypes, stage.CodeServiceSARTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpecialDateTypes, stage.CodeSpecialDateTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusOperationsTypes, stage.CodeStatusOperationsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusServiceTypes, stage.CodeStatusServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSystemActivationTypes, stage.CodeSystemActivationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUnitDependencyTypes, stage.CodeUnitDependencyTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUnitTypes, stage.CodeUnitTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeYesNoTypes, stage.CodeYesNoTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateMonthDayTypes, stage.DateMonthDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateYearTypes, stage.DateYearTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoNumberTypes, stage.NoNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextAddressTypes, stage.TextAddressTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextDesignatorTypes, stage.TextDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextInstructionTypes, stage.TextInstructionTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNameTypes, stage.TextNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNoteTypes, stage.TextNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPhoneTypes, stage.TextPhoneTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPropertyNameTypes, stage.TextPropertyNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimeTypes, stage.TimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValBearingTypes, stage.ValBearingTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceTypes, stage.ValDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDurationTypes, stage.ValDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValFrequencyTypes, stage.ValFrequencyTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftGroundServiceType_stagedOrder, stage.CodeAircraftGroundServiceTypes_referenceOrder, codeaircraftgroundservicetype, "CodeAircraftGroundServiceType")
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAviationStandardsType_stagedOrder, stage.CodeAviationStandardsTypes_referenceOrder, codeaviationstandardstype, "CodeAviationStandardsType")
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

func (codedaytype *CodeDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDayType_stagedOrder, stage.CodeDayTypes_referenceOrder, codedaytype, "CodeDayType")
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFacilityRankingType_stagedOrder, stage.CodeFacilityRankingTypes_referenceOrder, codefacilityrankingtype, "CodeFacilityRankingType")
}

func (codefirefightingtype *CodeFireFightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFireFightingType_stagedOrder, stage.CodeFireFightingTypes_referenceOrder, codefirefightingtype, "CodeFireFightingType")
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightDestinationType_stagedOrder, stage.CodeFlightDestinationTypes_referenceOrder, codeflightdestinationtype, "CodeFlightDestinationType")
}

func (codeflightruletype *CodeFlightRuleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFlightRuleType_stagedOrder, stage.CodeFlightRuleTypes_referenceOrder, codeflightruletype, "CodeFlightRuleType")
}

func (codefueltype *CodeFuelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFuelType_stagedOrder, stage.CodeFuelTypes_referenceOrder, codefueltype, "CodeFuelType")
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

func (codelightintensitytype *CodeLightIntensityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLightIntensityType_stagedOrder, stage.CodeLightIntensityTypes_referenceOrder, codelightintensitytype, "CodeLightIntensityType")
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMilitaryOperationsType_stagedOrder, stage.CodeMilitaryOperationsTypes_referenceOrder, codemilitaryoperationstype, "CodeMilitaryOperationsType")
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

func (codeoiltype *CodeOilType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeOilType_stagedOrder, stage.CodeOilTypes_referenceOrder, codeoiltype, "CodeOilType")
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

func (codepassengerservicetype *CodePassengerServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePassengerServiceType_stagedOrder, stage.CodePassengerServiceTypes_referenceOrder, codepassengerservicetype, "CodePassengerServiceType")
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePilotControlledLightingType_stagedOrder, stage.CodePilotControlledLightingTypes_referenceOrder, codepilotcontrolledlightingtype, "CodePilotControlledLightingType")
}

func (codervsmtype *CodeRVSMType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMType_stagedOrder, stage.CodeRVSMTypes_referenceOrder, codervsmtype, "CodeRVSMType")
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRadioEmissionType_stagedOrder, stage.CodeRadioEmissionTypes_referenceOrder, coderadioemissiontype, "CodeRadioEmissionType")
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

func (codespecialdatetype *CodeSpecialDateType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSpecialDateType_stagedOrder, stage.CodeSpecialDateTypes_referenceOrder, codespecialdatetype, "CodeSpecialDateType")
}

func (codestatusoperationstype *CodeStatusOperationsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusOperationsType_stagedOrder, stage.CodeStatusOperationsTypes_referenceOrder, codestatusoperationstype, "CodeStatusOperationsType")
}

func (codestatusservicetype *CodeStatusServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusServiceType_stagedOrder, stage.CodeStatusServiceTypes_referenceOrder, codestatusservicetype, "CodeStatusServiceType")
}

func (codesystemactivationtype *CodeSystemActivationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSystemActivationType_stagedOrder, stage.CodeSystemActivationTypes_referenceOrder, codesystemactivationtype, "CodeSystemActivationType")
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

func (codeunitdependencytype *CodeUnitDependencyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUnitDependencyType_stagedOrder, stage.CodeUnitDependencyTypes_referenceOrder, codeunitdependencytype, "CodeUnitDependencyType")
}

func (codeunittype *CodeUnitType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUnitType_stagedOrder, stage.CodeUnitTypes_referenceOrder, codeunittype, "CodeUnitType")
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalDatumType_stagedOrder, stage.CodeVerticalDatumTypes_referenceOrder, codeverticaldatumtype, "CodeVerticalDatumType")
}

func (codeyesnotype *CodeYesNoType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeYesNoType_stagedOrder, stage.CodeYesNoTypes_referenceOrder, codeyesnotype, "CodeYesNoType")
}

func (datemonthdaytype *DateMonthDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateMonthDayType_stagedOrder, stage.DateMonthDayTypes_referenceOrder, datemonthdaytype, "DateMonthDayType")
}

func (dateyeartype *DateYearType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateYearType_stagedOrder, stage.DateYearTypes_referenceOrder, dateyeartype, "DateYearType")
}

func (nonumbertype *NoNumberType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoNumberType_stagedOrder, stage.NoNumberTypes_referenceOrder, nonumbertype, "NoNumberType")
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

func (valbearingtype *ValBearingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValBearingType_stagedOrder, stage.ValBearingTypes_referenceOrder, valbearingtype, "ValBearingType")
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

func (valfrequencytype *ValFrequencyType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValFrequencyType_stagedOrder, stage.ValFrequencyTypes_referenceOrder, valfrequencytype, "ValFrequencyType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircraftgroundservicetype, codeaircraftgroundservicetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircraftgroundservicetype.GongGetIdentifier(stage)
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaviationstandardstype, codeaviationstandardstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaviationstandardstype *CodeAviationStandardsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaviationstandardstype.GongGetIdentifier(stage)
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

func (codedaytype *CodeDayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedaytype, codedaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedaytype *CodeDayType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedaytype.GongGetIdentifier(stage)
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefacilityrankingtype, codefacilityrankingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefacilityrankingtype *CodeFacilityRankingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefacilityrankingtype.GongGetIdentifier(stage)
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

func (codeflightruletype *CodeFlightRuleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeflightruletype, codeflightruletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeflightruletype *CodeFlightRuleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeflightruletype.GongGetIdentifier(stage)
}

func (codefueltype *CodeFuelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefueltype, codefueltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefueltype *CodeFuelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefueltype.GongGetIdentifier(stage)
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

func (codelightintensitytype *CodeLightIntensityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelightintensitytype, codelightintensitytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelightintensitytype *CodeLightIntensityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelightintensitytype.GongGetIdentifier(stage)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemilitaryoperationstype, codemilitaryoperationstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemilitaryoperationstype.GongGetIdentifier(stage)
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

func (codeoiltype *CodeOilType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeoiltype, codeoiltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeoiltype *CodeOilType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeoiltype.GongGetIdentifier(stage)
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

func (codervsmtype *CodeRVSMType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codervsmtype, codervsmtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codervsmtype *CodeRVSMType) GongGetReferenceIdentifier(stage *Stage) string {
	return codervsmtype.GongGetIdentifier(stage)
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderadioemissiontype, coderadioemissiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderadioemissiontype *CodeRadioEmissionType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderadioemissiontype.GongGetIdentifier(stage)
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

func (codespecialdatetype *CodeSpecialDateType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codespecialdatetype, codespecialdatetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codespecialdatetype *CodeSpecialDateType) GongGetReferenceIdentifier(stage *Stage) string {
	return codespecialdatetype.GongGetIdentifier(stage)
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

func (codesystemactivationtype *CodeSystemActivationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesystemactivationtype, codesystemactivationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesystemactivationtype *CodeSystemActivationType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesystemactivationtype.GongGetIdentifier(stage)
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

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeverticaldatumtype, codeverticaldatumtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeverticaldatumtype *CodeVerticalDatumType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeverticaldatumtype.GongGetIdentifier(stage)
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

func (valbearingtype *ValBearingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valbearingtype, valbearingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valbearingtype *ValBearingType) GongGetReferenceIdentifier(stage *Stage) string {
	return valbearingtype.GongGetIdentifier(stage)
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

func (valfrequencytype *ValFrequencyType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valfrequencytype, valfrequencytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valfrequencytype *ValFrequencyType) GongGetReferenceIdentifier(stage *Stage) string {
	return valfrequencytype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftgroundservicetype.GongGetIdentifier(stage), "CodeAircraftGroundServiceType", codeaircraftgroundservicetype.Name)
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaviationstandardstype.GongGetIdentifier(stage), "CodeAviationStandardsType", codeaviationstandardstype.Name)
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

func (codedaytype *CodeDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedaytype.GongGetIdentifier(stage), "CodeDayType", codedaytype.Name)
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefacilityrankingtype.GongGetIdentifier(stage), "CodeFacilityRankingType", codefacilityrankingtype.Name)
}

func (codefirefightingtype *CodeFireFightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefirefightingtype.GongGetIdentifier(stage), "CodeFireFightingType", codefirefightingtype.Name)
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightdestinationtype.GongGetIdentifier(stage), "CodeFlightDestinationType", codeflightdestinationtype.Name)
}

func (codeflightruletype *CodeFlightRuleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeflightruletype.GongGetIdentifier(stage), "CodeFlightRuleType", codeflightruletype.Name)
}

func (codefueltype *CodeFuelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefueltype.GongGetIdentifier(stage), "CodeFuelType", codefueltype.Name)
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

func (codelightintensitytype *CodeLightIntensityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelightintensitytype.GongGetIdentifier(stage), "CodeLightIntensityType", codelightintensitytype.Name)
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemilitaryoperationstype.GongGetIdentifier(stage), "CodeMilitaryOperationsType", codemilitaryoperationstype.Name)
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

func (codeoiltype *CodeOilType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeoiltype.GongGetIdentifier(stage), "CodeOilType", codeoiltype.Name)
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

func (codepassengerservicetype *CodePassengerServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepassengerservicetype.GongGetIdentifier(stage), "CodePassengerServiceType", codepassengerservicetype.Name)
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepilotcontrolledlightingtype.GongGetIdentifier(stage), "CodePilotControlledLightingType", codepilotcontrolledlightingtype.Name)
}

func (codervsmtype *CodeRVSMType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmtype.GongGetIdentifier(stage), "CodeRVSMType", codervsmtype.Name)
}

func (coderadioemissiontype *CodeRadioEmissionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderadioemissiontype.GongGetIdentifier(stage), "CodeRadioEmissionType", coderadioemissiontype.Name)
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

func (codespecialdatetype *CodeSpecialDateType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codespecialdatetype.GongGetIdentifier(stage), "CodeSpecialDateType", codespecialdatetype.Name)
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusoperationstype.GongGetIdentifier(stage), "CodeStatusOperationsType", codestatusoperationstype.Name)
}

func (codestatusservicetype *CodeStatusServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusservicetype.GongGetIdentifier(stage), "CodeStatusServiceType", codestatusservicetype.Name)
}

func (codesystemactivationtype *CodeSystemActivationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesystemactivationtype.GongGetIdentifier(stage), "CodeSystemActivationType", codesystemactivationtype.Name)
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

func (codeunitdependencytype *CodeUnitDependencyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeunitdependencytype.GongGetIdentifier(stage), "CodeUnitDependencyType", codeunitdependencytype.Name)
}

func (codeunittype *CodeUnitType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeunittype.GongGetIdentifier(stage), "CodeUnitType", codeunittype.Name)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticaldatumtype.GongGetIdentifier(stage), "CodeVerticalDatumType", codeverticaldatumtype.Name)
}

func (codeyesnotype *CodeYesNoType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeyesnotype.GongGetIdentifier(stage), "CodeYesNoType", codeyesnotype.Name)
}

func (datemonthdaytype *DateMonthDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datemonthdaytype.GongGetIdentifier(stage), "DateMonthDayType", datemonthdaytype.Name)
}

func (dateyeartype *DateYearType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dateyeartype.GongGetIdentifier(stage), "DateYearType", dateyeartype.Name)
}

func (nonumbertype *NoNumberType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nonumbertype.GongGetIdentifier(stage), "NoNumberType", nonumbertype.Name)
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

func (valbearingtype *ValBearingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valbearingtype.GongGetIdentifier(stage), "ValBearingType", valbearingtype.Name)
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

func (valfrequencytype *ValFrequencyType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valfrequencytype.GongGetIdentifier(stage), "ValFrequencyType", valfrequencytype.Name)
}

// insertion point for unstaging
func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftgroundservicetype.GongGetReferenceIdentifier(stage))
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaviationstandardstype.GongGetReferenceIdentifier(stage))
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

func (codedaytype *CodeDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedaytype.GongGetReferenceIdentifier(stage))
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefacilityrankingtype.GongGetReferenceIdentifier(stage))
}

func (codefirefightingtype *CodeFireFightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefirefightingtype.GongGetReferenceIdentifier(stage))
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightdestinationtype.GongGetReferenceIdentifier(stage))
}

func (codeflightruletype *CodeFlightRuleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeflightruletype.GongGetReferenceIdentifier(stage))
}

func (codefueltype *CodeFuelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefueltype.GongGetReferenceIdentifier(stage))
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

func (codelightintensitytype *CodeLightIntensityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelightintensitytype.GongGetReferenceIdentifier(stage))
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemilitaryoperationstype.GongGetReferenceIdentifier(stage))
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

func (codeoiltype *CodeOilType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeoiltype.GongGetReferenceIdentifier(stage))
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

func (codepassengerservicetype *CodePassengerServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepassengerservicetype.GongGetReferenceIdentifier(stage))
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepilotcontrolledlightingtype.GongGetReferenceIdentifier(stage))
}

func (codervsmtype *CodeRVSMType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmtype.GongGetReferenceIdentifier(stage))
}

func (coderadioemissiontype *CodeRadioEmissionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderadioemissiontype.GongGetReferenceIdentifier(stage))
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

func (codespecialdatetype *CodeSpecialDateType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codespecialdatetype.GongGetReferenceIdentifier(stage))
}

func (codestatusoperationstype *CodeStatusOperationsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusoperationstype.GongGetReferenceIdentifier(stage))
}

func (codestatusservicetype *CodeStatusServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusservicetype.GongGetReferenceIdentifier(stage))
}

func (codesystemactivationtype *CodeSystemActivationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesystemactivationtype.GongGetReferenceIdentifier(stage))
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

func (codeunitdependencytype *CodeUnitDependencyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeunitdependencytype.GongGetReferenceIdentifier(stage))
}

func (codeunittype *CodeUnitType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeunittype.GongGetReferenceIdentifier(stage))
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticaldatumtype.GongGetReferenceIdentifier(stage))
}

func (codeyesnotype *CodeYesNoType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeyesnotype.GongGetReferenceIdentifier(stage))
}

func (datemonthdaytype *DateMonthDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datemonthdaytype.GongGetReferenceIdentifier(stage))
}

func (dateyeartype *DateYearType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dateyeartype.GongGetReferenceIdentifier(stage))
}

func (nonumbertype *NoNumberType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nonumbertype.GongGetReferenceIdentifier(stage))
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

func (valbearingtype *ValBearingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valbearingtype.GongGetReferenceIdentifier(stage))
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

func (valfrequencytype *ValFrequencyType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valfrequencytype.GongGetReferenceIdentifier(stage))
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
