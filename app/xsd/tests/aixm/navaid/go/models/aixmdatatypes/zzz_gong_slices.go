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
	res = __gong__appendInstances(res, stage.CodeAirspacePointPositionTypes)

	res = __gong__appendInstances(res, stage.CodeAirspacePointRoleTypes)

	res = __gong__appendInstances(res, stage.CodeAltitudeUseTypes)

	res = __gong__appendInstances(res, stage.CodeAuralMorseTypes)

	res = __gong__appendInstances(res, stage.CodeAuthorityRoleTypes)

	res = __gong__appendInstances(res, stage.CodeBearingTypes)

	res = __gong__appendInstances(res, stage.CodeCardinalDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeCheckpointCategoryTypes)

	res = __gong__appendInstances(res, stage.CodeCourseQualityILSTypes)

	res = __gong__appendInstances(res, stage.CodeDMEChannelTypes)

	res = __gong__appendInstances(res, stage.CodeDMETypes)

	res = __gong__appendInstances(res, stage.CodeDayTypes)

	res = __gong__appendInstances(res, stage.CodeDesignatedPointDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeDesignatedPointTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeDistanceIndicationTypes)

	res = __gong__appendInstances(res, stage.CodeEmissionBandTypes)

	res = __gong__appendInstances(res, stage.CodeILSBackCourseTypes)

	res = __gong__appendInstances(res, stage.CodeIntegrityLevelILSTypes)

	res = __gong__appendInstances(res, stage.CodeMLSAzimuthTypes)

	res = __gong__appendInstances(res, stage.CodeMLSChannelTypes)

	res = __gong__appendInstances(res, stage.CodeMarkerBeaconSignalTypes)

	res = __gong__appendInstances(res, stage.CodeNDBUsageTypes)

	res = __gong__appendInstances(res, stage.CodeNavaidDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeNavaidPurposeTypes)

	res = __gong__appendInstances(res, stage.CodeNavaidServiceTypes)

	res = __gong__appendInstances(res, stage.CodeNorthReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeNotePurposeTypes)

	res = __gong__appendInstances(res, stage.CodePARTypes)

	res = __gong__appendInstances(res, stage.CodePositionInILSTypes)

	res = __gong__appendInstances(res, stage.CodePrimaryRadarTypes)

	res = __gong__appendInstances(res, stage.CodeRadarServiceTypes)

	res = __gong__appendInstances(res, stage.CodeRadioEmissionTypes)

	res = __gong__appendInstances(res, stage.CodeRadioSignalTypes)

	res = __gong__appendInstances(res, stage.CodeReflectorTypes)

	res = __gong__appendInstances(res, stage.CodeSignalPerformanceILSTypes)

	res = __gong__appendInstances(res, stage.CodeSpecialNavigationChainDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeSpecialNavigationStationTypes)

	res = __gong__appendInstances(res, stage.CodeSpecialNavigationSystemTypes)

	res = __gong__appendInstances(res, stage.CodeStandbyPowerTypes)

	res = __gong__appendInstances(res, stage.CodeStatusNavaidTypes)

	res = __gong__appendInstances(res, stage.CodeTACANChannelTypes)

	res = __gong__appendInstances(res, stage.CodeTelecomNetworkTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventCombinationTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventTypes)

	res = __gong__appendInstances(res, stage.CodeTimeReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeTransponderTypes)

	res = __gong__appendInstances(res, stage.CodeVORTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalDatumTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeYesNoTypes)

	res = __gong__appendInstances(res, stage.DateMonthDayTypes)

	res = __gong__appendInstances(res, stage.DateYearTypes)

	res = __gong__appendInstances(res, stage.NoSequenceTypes)

	res = __gong__appendInstances(res, stage.TextAddressTypes)

	res = __gong__appendInstances(res, stage.TextDesignatorTypes)

	res = __gong__appendInstances(res, stage.TextNameTypes)

	res = __gong__appendInstances(res, stage.TextNoteTypes)

	res = __gong__appendInstances(res, stage.TextPhoneTypes)

	res = __gong__appendInstances(res, stage.TextPropertyNameTypes)

	res = __gong__appendInstances(res, stage.TimeTypes)

	res = __gong__appendInstances(res, stage.ValAngleTypes)

	res = __gong__appendInstances(res, stage.ValBearingTypes)

	res = __gong__appendInstances(res, stage.ValDistanceSignedTypes)

	res = __gong__appendInstances(res, stage.ValDistanceTypes)

	res = __gong__appendInstances(res, stage.ValDistanceVerticalTypes)

	res = __gong__appendInstances(res, stage.ValDurationTypes)

	res = __gong__appendInstances(res, stage.ValFrequencyTypes)

	res = __gong__appendInstances(res, stage.ValMagneticVariationTypes)

	return
}

// insertion point per named struct
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

func (codealtitudeusetype *CodeAltitudeUseType) GongCopy() GongstructIF {
	newInstance := new(CodeAltitudeUseType)
	codealtitudeusetype.GongCopyBasicFields(newInstance)
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

func (codebearingtype *CodeBearingType) GongCopy() GongstructIF {
	newInstance := new(CodeBearingType)
	codebearingtype.GongCopyBasicFields(newInstance)
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

func (codecoursequalityilstype *CodeCourseQualityILSType) GongCopy() GongstructIF {
	newInstance := new(CodeCourseQualityILSType)
	codecoursequalityilstype.GongCopyBasicFields(newInstance)
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

func (codepartype *CodePARType) GongCopy() GongstructIF {
	newInstance := new(CodePARType)
	codepartype.GongCopyBasicFields(newInstance)
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

func (coderadiosignaltype *CodeRadioSignalType) GongCopy() GongstructIF {
	newInstance := new(CodeRadioSignalType)
	coderadiosignaltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codereflectortype *CodeReflectorType) GongCopy() GongstructIF {
	newInstance := new(CodeReflectorType)
	codereflectortype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongCopy() GongstructIF {
	newInstance := new(CodeSignalPerformanceILSType)
	codesignalperformanceilstype.GongCopyBasicFields(newInstance)
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

func (codestandbypowertype *CodeStandbyPowerType) GongCopy() GongstructIF {
	newInstance := new(CodeStandbyPowerType)
	codestandbypowertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongCopy() GongstructIF {
	newInstance := new(CodeStatusNavaidType)
	codestatusnavaidtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codetacanchanneltype *CodeTACANChannelType) GongCopy() GongstructIF {
	newInstance := new(CodeTACANChannelType)
	codetacanchanneltype.GongCopyBasicFields(newInstance)
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

func (codevortype *CodeVORType) GongCopy() GongstructIF {
	newInstance := new(CodeVORType)
	codevortype.GongCopyBasicFields(newInstance)
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

func (valmagneticvariationtype *ValMagneticVariationType) GongCopy() GongstructIF {
	newInstance := new(ValMagneticVariationType)
	valmagneticvariationtype.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacepointpositiontype)
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeairspacepointroletype)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codealtitudeusetype)
}

func (codeauralmorsetype *CodeAuralMorseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeauralmorsetype)
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeauthorityroletype)
}

func (codebearingtype *CodeBearingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codebearingtype)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecardinaldirectiontype)
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecheckpointcategorytype)
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecoursequalityilstype)
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

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedesignatedpointdesignatortype)
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedesignatedpointtype)
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedirectionreferencetype)
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedistanceindicationtype)
}

func (codeemissionbandtype *CodeEmissionBandType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeemissionbandtype)
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeilsbackcoursetype)
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeintegritylevelilstype)
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

func (codenorthreferencetype *CodeNorthReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenorthreferencetype)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenotepurposetype)
}

func (codepartype *CodePARType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepartype)
}

func (codepositioninilstype *CodePositionInILSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codepositioninilstype)
}

func (codeprimaryradartype *CodePrimaryRadarType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeprimaryradartype)
}

func (coderadarservicetype *CodeRadarServiceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderadarservicetype)
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderadioemissiontype)
}

func (coderadiosignaltype *CodeRadioSignalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderadiosignaltype)
}

func (codereflectortype *CodeReflectorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codereflectortype)
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesignalperformanceilstype)
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

func (codestandbypowertype *CodeStandbyPowerType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestandbypowertype)
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codestatusnavaidtype)
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetacanchanneltype)
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

func (codevortype *CodeVORType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codevortype)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticaldatumtype)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeverticalreferencetype)
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

func (nosequencetype *NoSequenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, nosequencetype)
}

func (textaddresstype *TextAddressType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textaddresstype)
}

func (textdesignatortype *TextDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textdesignatortype)
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

func (valmagneticvariationtype *ValMagneticVariationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valmagneticvariationtype)
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
	__gong__computeReferencePass1(stage, stage.CodeAirspacePointPositionTypes, &stage.CodeAirspacePointPositionTypes_reference, &stage.CodeAirspacePointPositionTypes_referenceOrder, &stage.CodeAirspacePointPositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAirspacePointRoleTypes, &stage.CodeAirspacePointRoleTypes_reference, &stage.CodeAirspacePointRoleTypes_referenceOrder, &stage.CodeAirspacePointRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_referenceOrder, &stage.CodeAltitudeUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAuralMorseTypes, &stage.CodeAuralMorseTypes_reference, &stage.CodeAuralMorseTypes_referenceOrder, &stage.CodeAuralMorseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAuthorityRoleTypes, &stage.CodeAuthorityRoleTypes_reference, &stage.CodeAuthorityRoleTypes_referenceOrder, &stage.CodeAuthorityRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeBearingTypes, &stage.CodeBearingTypes_reference, &stage.CodeBearingTypes_referenceOrder, &stage.CodeBearingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCardinalDirectionTypes, &stage.CodeCardinalDirectionTypes_reference, &stage.CodeCardinalDirectionTypes_referenceOrder, &stage.CodeCardinalDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCheckpointCategoryTypes, &stage.CodeCheckpointCategoryTypes_reference, &stage.CodeCheckpointCategoryTypes_referenceOrder, &stage.CodeCheckpointCategoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCourseQualityILSTypes, &stage.CodeCourseQualityILSTypes_reference, &stage.CodeCourseQualityILSTypes_referenceOrder, &stage.CodeCourseQualityILSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDMEChannelTypes, &stage.CodeDMEChannelTypes_reference, &stage.CodeDMEChannelTypes_referenceOrder, &stage.CodeDMEChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDMETypes, &stage.CodeDMETypes_reference, &stage.CodeDMETypes_referenceOrder, &stage.CodeDMETypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDayTypes, &stage.CodeDayTypes_reference, &stage.CodeDayTypes_referenceOrder, &stage.CodeDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDesignatedPointDesignatorTypes, &stage.CodeDesignatedPointDesignatorTypes_reference, &stage.CodeDesignatedPointDesignatorTypes_referenceOrder, &stage.CodeDesignatedPointDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDesignatedPointTypes, &stage.CodeDesignatedPointTypes_reference, &stage.CodeDesignatedPointTypes_referenceOrder, &stage.CodeDesignatedPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionReferenceTypes, &stage.CodeDirectionReferenceTypes_reference, &stage.CodeDirectionReferenceTypes_referenceOrder, &stage.CodeDirectionReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDistanceIndicationTypes, &stage.CodeDistanceIndicationTypes_reference, &stage.CodeDistanceIndicationTypes_referenceOrder, &stage.CodeDistanceIndicationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEmissionBandTypes, &stage.CodeEmissionBandTypes_reference, &stage.CodeEmissionBandTypes_referenceOrder, &stage.CodeEmissionBandTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeILSBackCourseTypes, &stage.CodeILSBackCourseTypes_reference, &stage.CodeILSBackCourseTypes_referenceOrder, &stage.CodeILSBackCourseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeIntegrityLevelILSTypes, &stage.CodeIntegrityLevelILSTypes_reference, &stage.CodeIntegrityLevelILSTypes_referenceOrder, &stage.CodeIntegrityLevelILSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMLSAzimuthTypes, &stage.CodeMLSAzimuthTypes_reference, &stage.CodeMLSAzimuthTypes_referenceOrder, &stage.CodeMLSAzimuthTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMLSChannelTypes, &stage.CodeMLSChannelTypes_reference, &stage.CodeMLSChannelTypes_referenceOrder, &stage.CodeMLSChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMarkerBeaconSignalTypes, &stage.CodeMarkerBeaconSignalTypes_reference, &stage.CodeMarkerBeaconSignalTypes_referenceOrder, &stage.CodeMarkerBeaconSignalTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNDBUsageTypes, &stage.CodeNDBUsageTypes_reference, &stage.CodeNDBUsageTypes_referenceOrder, &stage.CodeNDBUsageTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavaidDesignatorTypes, &stage.CodeNavaidDesignatorTypes_reference, &stage.CodeNavaidDesignatorTypes_referenceOrder, &stage.CodeNavaidDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavaidPurposeTypes, &stage.CodeNavaidPurposeTypes_reference, &stage.CodeNavaidPurposeTypes_referenceOrder, &stage.CodeNavaidPurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavaidServiceTypes, &stage.CodeNavaidServiceTypes_reference, &stage.CodeNavaidServiceTypes_referenceOrder, &stage.CodeNavaidServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNorthReferenceTypes, &stage.CodeNorthReferenceTypes_reference, &stage.CodeNorthReferenceTypes_referenceOrder, &stage.CodeNorthReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_referenceOrder, &stage.CodeNotePurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePARTypes, &stage.CodePARTypes_reference, &stage.CodePARTypes_referenceOrder, &stage.CodePARTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePositionInILSTypes, &stage.CodePositionInILSTypes_reference, &stage.CodePositionInILSTypes_referenceOrder, &stage.CodePositionInILSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodePrimaryRadarTypes, &stage.CodePrimaryRadarTypes_reference, &stage.CodePrimaryRadarTypes_referenceOrder, &stage.CodePrimaryRadarTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRadarServiceTypes, &stage.CodeRadarServiceTypes_reference, &stage.CodeRadarServiceTypes_referenceOrder, &stage.CodeRadarServiceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRadioEmissionTypes, &stage.CodeRadioEmissionTypes_reference, &stage.CodeRadioEmissionTypes_referenceOrder, &stage.CodeRadioEmissionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRadioSignalTypes, &stage.CodeRadioSignalTypes_reference, &stage.CodeRadioSignalTypes_referenceOrder, &stage.CodeRadioSignalTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeReflectorTypes, &stage.CodeReflectorTypes_reference, &stage.CodeReflectorTypes_referenceOrder, &stage.CodeReflectorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSignalPerformanceILSTypes, &stage.CodeSignalPerformanceILSTypes_reference, &stage.CodeSignalPerformanceILSTypes_referenceOrder, &stage.CodeSignalPerformanceILSTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpecialNavigationChainDesignatorTypes, &stage.CodeSpecialNavigationChainDesignatorTypes_reference, &stage.CodeSpecialNavigationChainDesignatorTypes_referenceOrder, &stage.CodeSpecialNavigationChainDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpecialNavigationStationTypes, &stage.CodeSpecialNavigationStationTypes_reference, &stage.CodeSpecialNavigationStationTypes_referenceOrder, &stage.CodeSpecialNavigationStationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpecialNavigationSystemTypes, &stage.CodeSpecialNavigationSystemTypes_reference, &stage.CodeSpecialNavigationSystemTypes_referenceOrder, &stage.CodeSpecialNavigationSystemTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStandbyPowerTypes, &stage.CodeStandbyPowerTypes_reference, &stage.CodeStandbyPowerTypes_referenceOrder, &stage.CodeStandbyPowerTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeStatusNavaidTypes, &stage.CodeStatusNavaidTypes_reference, &stage.CodeStatusNavaidTypes_referenceOrder, &stage.CodeStatusNavaidTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTACANChannelTypes, &stage.CodeTACANChannelTypes_reference, &stage.CodeTACANChannelTypes_referenceOrder, &stage.CodeTACANChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_referenceOrder, &stage.CodeTelecomNetworkTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_referenceOrder, &stage.CodeTimeEventCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_referenceOrder, &stage.CodeTimeEventTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_referenceOrder, &stage.CodeTimeReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTransponderTypes, &stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_referenceOrder, &stage.CodeTransponderTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVORTypes, &stage.CodeVORTypes_reference, &stage.CodeVORTypes_referenceOrder, &stage.CodeVORTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_referenceOrder, &stage.CodeVerticalDatumTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_reference, &stage.CodeVerticalReferenceTypes_referenceOrder, &stage.CodeVerticalReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeYesNoTypes, &stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_referenceOrder, &stage.CodeYesNoTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateMonthDayTypes, &stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_referenceOrder, &stage.DateMonthDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateYearTypes, &stage.DateYearTypes_reference, &stage.DateYearTypes_referenceOrder, &stage.DateYearTypes_instance)

	__gong__computeReferencePass1(stage, stage.NoSequenceTypes, &stage.NoSequenceTypes_reference, &stage.NoSequenceTypes_referenceOrder, &stage.NoSequenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextAddressTypes, &stage.TextAddressTypes_reference, &stage.TextAddressTypes_referenceOrder, &stage.TextAddressTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextDesignatorTypes, &stage.TextDesignatorTypes_reference, &stage.TextDesignatorTypes_referenceOrder, &stage.TextDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNameTypes, &stage.TextNameTypes_reference, &stage.TextNameTypes_referenceOrder, &stage.TextNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNoteTypes, &stage.TextNoteTypes_reference, &stage.TextNoteTypes_referenceOrder, &stage.TextNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPhoneTypes, &stage.TextPhoneTypes_reference, &stage.TextPhoneTypes_referenceOrder, &stage.TextPhoneTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_referenceOrder, &stage.TextPropertyNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimeTypes, &stage.TimeTypes_reference, &stage.TimeTypes_referenceOrder, &stage.TimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValAngleTypes, &stage.ValAngleTypes_reference, &stage.ValAngleTypes_referenceOrder, &stage.ValAngleTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValBearingTypes, &stage.ValBearingTypes_reference, &stage.ValBearingTypes_referenceOrder, &stage.ValBearingTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_referenceOrder, &stage.ValDistanceSignedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceTypes, &stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_referenceOrder, &stage.ValDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_referenceOrder, &stage.ValDistanceVerticalTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDurationTypes, &stage.ValDurationTypes_reference, &stage.ValDurationTypes_referenceOrder, &stage.ValDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValFrequencyTypes, &stage.ValFrequencyTypes_reference, &stage.ValFrequencyTypes_referenceOrder, &stage.ValFrequencyTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValMagneticVariationTypes, &stage.ValMagneticVariationTypes_reference, &stage.ValMagneticVariationTypes_referenceOrder, &stage.ValMagneticVariationTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CodeAirspacePointPositionTypes, stage.CodeAirspacePointPositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAirspacePointRoleTypes, stage.CodeAirspacePointRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAuralMorseTypes, stage.CodeAuralMorseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAuthorityRoleTypes, stage.CodeAuthorityRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeBearingTypes, stage.CodeBearingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCardinalDirectionTypes, stage.CodeCardinalDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCheckpointCategoryTypes, stage.CodeCheckpointCategoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCourseQualityILSTypes, stage.CodeCourseQualityILSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDMEChannelTypes, stage.CodeDMEChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDMETypes, stage.CodeDMETypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDayTypes, stage.CodeDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDesignatedPointDesignatorTypes, stage.CodeDesignatedPointDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDesignatedPointTypes, stage.CodeDesignatedPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionReferenceTypes, stage.CodeDirectionReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDistanceIndicationTypes, stage.CodeDistanceIndicationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEmissionBandTypes, stage.CodeEmissionBandTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeILSBackCourseTypes, stage.CodeILSBackCourseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeIntegrityLevelILSTypes, stage.CodeIntegrityLevelILSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMLSAzimuthTypes, stage.CodeMLSAzimuthTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMLSChannelTypes, stage.CodeMLSChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMarkerBeaconSignalTypes, stage.CodeMarkerBeaconSignalTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNDBUsageTypes, stage.CodeNDBUsageTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavaidDesignatorTypes, stage.CodeNavaidDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavaidPurposeTypes, stage.CodeNavaidPurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavaidServiceTypes, stage.CodeNavaidServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNorthReferenceTypes, stage.CodeNorthReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePARTypes, stage.CodePARTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePositionInILSTypes, stage.CodePositionInILSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodePrimaryRadarTypes, stage.CodePrimaryRadarTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRadarServiceTypes, stage.CodeRadarServiceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRadioEmissionTypes, stage.CodeRadioEmissionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRadioSignalTypes, stage.CodeRadioSignalTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeReflectorTypes, stage.CodeReflectorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSignalPerformanceILSTypes, stage.CodeSignalPerformanceILSTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpecialNavigationChainDesignatorTypes, stage.CodeSpecialNavigationChainDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpecialNavigationStationTypes, stage.CodeSpecialNavigationStationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpecialNavigationSystemTypes, stage.CodeSpecialNavigationSystemTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStandbyPowerTypes, stage.CodeStandbyPowerTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeStatusNavaidTypes, stage.CodeStatusNavaidTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTACANChannelTypes, stage.CodeTACANChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTransponderTypes, stage.CodeTransponderTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVORTypes, stage.CodeVORTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeYesNoTypes, stage.CodeYesNoTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateMonthDayTypes, stage.DateMonthDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateYearTypes, stage.DateYearTypes_reference, stage)

	__gong__computeReferencePass2(stage.NoSequenceTypes, stage.NoSequenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextAddressTypes, stage.TextAddressTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextDesignatorTypes, stage.TextDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNameTypes, stage.TextNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNoteTypes, stage.TextNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPhoneTypes, stage.TextPhoneTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPropertyNameTypes, stage.TextPropertyNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimeTypes, stage.TimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValAngleTypes, stage.ValAngleTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValBearingTypes, stage.ValBearingTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceTypes, stage.ValDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDurationTypes, stage.ValDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValFrequencyTypes, stage.ValFrequencyTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValMagneticVariationTypes, stage.ValMagneticVariationTypes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspacePointPositionType_stagedOrder, stage.CodeAirspacePointPositionTypes_referenceOrder, codeairspacepointpositiontype, "CodeAirspacePointPositionType")
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAirspacePointRoleType_stagedOrder, stage.CodeAirspacePointRoleTypes_referenceOrder, codeairspacepointroletype, "CodeAirspacePointRoleType")
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAltitudeUseType_stagedOrder, stage.CodeAltitudeUseTypes_referenceOrder, codealtitudeusetype, "CodeAltitudeUseType")
}

func (codeauralmorsetype *CodeAuralMorseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAuralMorseType_stagedOrder, stage.CodeAuralMorseTypes_referenceOrder, codeauralmorsetype, "CodeAuralMorseType")
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAuthorityRoleType_stagedOrder, stage.CodeAuthorityRoleTypes_referenceOrder, codeauthorityroletype, "CodeAuthorityRoleType")
}

func (codebearingtype *CodeBearingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeBearingType_stagedOrder, stage.CodeBearingTypes_referenceOrder, codebearingtype, "CodeBearingType")
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCardinalDirectionType_stagedOrder, stage.CodeCardinalDirectionTypes_referenceOrder, codecardinaldirectiontype, "CodeCardinalDirectionType")
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCheckpointCategoryType_stagedOrder, stage.CodeCheckpointCategoryTypes_referenceOrder, codecheckpointcategorytype, "CodeCheckpointCategoryType")
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCourseQualityILSType_stagedOrder, stage.CodeCourseQualityILSTypes_referenceOrder, codecoursequalityilstype, "CodeCourseQualityILSType")
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

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDesignatedPointDesignatorType_stagedOrder, stage.CodeDesignatedPointDesignatorTypes_referenceOrder, codedesignatedpointdesignatortype, "CodeDesignatedPointDesignatorType")
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDesignatedPointType_stagedOrder, stage.CodeDesignatedPointTypes_referenceOrder, codedesignatedpointtype, "CodeDesignatedPointType")
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDirectionReferenceType_stagedOrder, stage.CodeDirectionReferenceTypes_referenceOrder, codedirectionreferencetype, "CodeDirectionReferenceType")
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDistanceIndicationType_stagedOrder, stage.CodeDistanceIndicationTypes_referenceOrder, codedistanceindicationtype, "CodeDistanceIndicationType")
}

func (codeemissionbandtype *CodeEmissionBandType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeEmissionBandType_stagedOrder, stage.CodeEmissionBandTypes_referenceOrder, codeemissionbandtype, "CodeEmissionBandType")
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeILSBackCourseType_stagedOrder, stage.CodeILSBackCourseTypes_referenceOrder, codeilsbackcoursetype, "CodeILSBackCourseType")
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeIntegrityLevelILSType_stagedOrder, stage.CodeIntegrityLevelILSTypes_referenceOrder, codeintegritylevelilstype, "CodeIntegrityLevelILSType")
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

func (codenorthreferencetype *CodeNorthReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNorthReferenceType_stagedOrder, stage.CodeNorthReferenceTypes_referenceOrder, codenorthreferencetype, "CodeNorthReferenceType")
}

func (codenotepurposetype *CodeNotePurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeTypes_referenceOrder, codenotepurposetype, "CodeNotePurposeType")
}

func (codepartype *CodePARType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePARType_stagedOrder, stage.CodePARTypes_referenceOrder, codepartype, "CodePARType")
}

func (codepositioninilstype *CodePositionInILSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePositionInILSType_stagedOrder, stage.CodePositionInILSTypes_referenceOrder, codepositioninilstype, "CodePositionInILSType")
}

func (codeprimaryradartype *CodePrimaryRadarType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodePrimaryRadarType_stagedOrder, stage.CodePrimaryRadarTypes_referenceOrder, codeprimaryradartype, "CodePrimaryRadarType")
}

func (coderadarservicetype *CodeRadarServiceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRadarServiceType_stagedOrder, stage.CodeRadarServiceTypes_referenceOrder, coderadarservicetype, "CodeRadarServiceType")
}

func (coderadioemissiontype *CodeRadioEmissionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRadioEmissionType_stagedOrder, stage.CodeRadioEmissionTypes_referenceOrder, coderadioemissiontype, "CodeRadioEmissionType")
}

func (coderadiosignaltype *CodeRadioSignalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRadioSignalType_stagedOrder, stage.CodeRadioSignalTypes_referenceOrder, coderadiosignaltype, "CodeRadioSignalType")
}

func (codereflectortype *CodeReflectorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeReflectorType_stagedOrder, stage.CodeReflectorTypes_referenceOrder, codereflectortype, "CodeReflectorType")
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSignalPerformanceILSType_stagedOrder, stage.CodeSignalPerformanceILSTypes_referenceOrder, codesignalperformanceilstype, "CodeSignalPerformanceILSType")
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

func (codestandbypowertype *CodeStandbyPowerType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStandbyPowerType_stagedOrder, stage.CodeStandbyPowerTypes_referenceOrder, codestandbypowertype, "CodeStandbyPowerType")
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeStatusNavaidType_stagedOrder, stage.CodeStatusNavaidTypes_referenceOrder, codestatusnavaidtype, "CodeStatusNavaidType")
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTACANChannelType_stagedOrder, stage.CodeTACANChannelTypes_referenceOrder, codetacanchanneltype, "CodeTACANChannelType")
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

func (codevortype *CodeVORType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVORType_stagedOrder, stage.CodeVORTypes_referenceOrder, codevortype, "CodeVORType")
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalDatumType_stagedOrder, stage.CodeVerticalDatumTypes_referenceOrder, codeverticaldatumtype, "CodeVerticalDatumType")
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeVerticalReferenceType_stagedOrder, stage.CodeVerticalReferenceTypes_referenceOrder, codeverticalreferencetype, "CodeVerticalReferenceType")
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

func (nosequencetype *NoSequenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoSequenceType_stagedOrder, stage.NoSequenceTypes_referenceOrder, nosequencetype, "NoSequenceType")
}

func (textaddresstype *TextAddressType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextAddressType_stagedOrder, stage.TextAddressTypes_referenceOrder, textaddresstype, "TextAddressType")
}

func (textdesignatortype *TextDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextDesignatorType_stagedOrder, stage.TextDesignatorTypes_referenceOrder, textdesignatortype, "TextDesignatorType")
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

func (valmagneticvariationtype *ValMagneticVariationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValMagneticVariationType_stagedOrder, stage.ValMagneticVariationTypes_referenceOrder, valmagneticvariationtype, "ValMagneticVariationType")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
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

func (codealtitudeusetype *CodeAltitudeUseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codealtitudeusetype, codealtitudeusetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codealtitudeusetype *CodeAltitudeUseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codealtitudeusetype.GongGetIdentifier(stage)
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

func (codebearingtype *CodeBearingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codebearingtype, codebearingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codebearingtype *CodeBearingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codebearingtype.GongGetIdentifier(stage)
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

func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecoursequalityilstype, codecoursequalityilstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecoursequalityilstype *CodeCourseQualityILSType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecoursequalityilstype.GongGetIdentifier(stage)
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

func (codepartype *CodePARType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codepartype, codepartype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codepartype *CodePARType) GongGetReferenceIdentifier(stage *Stage) string {
	return codepartype.GongGetIdentifier(stage)
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

func (coderadiosignaltype *CodeRadioSignalType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderadiosignaltype, coderadiosignaltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderadiosignaltype *CodeRadioSignalType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderadiosignaltype.GongGetIdentifier(stage)
}

func (codereflectortype *CodeReflectorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codereflectortype, codereflectortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codereflectortype *CodeReflectorType) GongGetReferenceIdentifier(stage *Stage) string {
	return codereflectortype.GongGetIdentifier(stage)
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesignalperformanceilstype, codesignalperformanceilstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesignalperformanceilstype.GongGetIdentifier(stage)
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

func (codestandbypowertype *CodeStandbyPowerType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestandbypowertype, codestandbypowertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestandbypowertype *CodeStandbyPowerType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestandbypowertype.GongGetIdentifier(stage)
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codestatusnavaidtype, codestatusnavaidtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codestatusnavaidtype *CodeStatusNavaidType) GongGetReferenceIdentifier(stage *Stage) string {
	return codestatusnavaidtype.GongGetIdentifier(stage)
}

func (codetacanchanneltype *CodeTACANChannelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetacanchanneltype, codetacanchanneltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetacanchanneltype *CodeTACANChannelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetacanchanneltype.GongGetIdentifier(stage)
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

func (codevortype *CodeVORType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codevortype, codevortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codevortype *CodeVORType) GongGetReferenceIdentifier(stage *Stage) string {
	return codevortype.GongGetIdentifier(stage)
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

func (valmagneticvariationtype *ValMagneticVariationType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valmagneticvariationtype, valmagneticvariationtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valmagneticvariationtype *ValMagneticVariationType) GongGetReferenceIdentifier(stage *Stage) string {
	return valmagneticvariationtype.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacepointpositiontype.GongGetIdentifier(stage), "CodeAirspacePointPositionType", codeairspacepointpositiontype.Name)
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeairspacepointroletype.GongGetIdentifier(stage), "CodeAirspacePointRoleType", codeairspacepointroletype.Name)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codealtitudeusetype.GongGetIdentifier(stage), "CodeAltitudeUseType", codealtitudeusetype.Name)
}

func (codeauralmorsetype *CodeAuralMorseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeauralmorsetype.GongGetIdentifier(stage), "CodeAuralMorseType", codeauralmorsetype.Name)
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeauthorityroletype.GongGetIdentifier(stage), "CodeAuthorityRoleType", codeauthorityroletype.Name)
}

func (codebearingtype *CodeBearingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codebearingtype.GongGetIdentifier(stage), "CodeBearingType", codebearingtype.Name)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecardinaldirectiontype.GongGetIdentifier(stage), "CodeCardinalDirectionType", codecardinaldirectiontype.Name)
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecheckpointcategorytype.GongGetIdentifier(stage), "CodeCheckpointCategoryType", codecheckpointcategorytype.Name)
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecoursequalityilstype.GongGetIdentifier(stage), "CodeCourseQualityILSType", codecoursequalityilstype.Name)
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

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedesignatedpointdesignatortype.GongGetIdentifier(stage), "CodeDesignatedPointDesignatorType", codedesignatedpointdesignatortype.Name)
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedesignatedpointtype.GongGetIdentifier(stage), "CodeDesignatedPointType", codedesignatedpointtype.Name)
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedirectionreferencetype.GongGetIdentifier(stage), "CodeDirectionReferenceType", codedirectionreferencetype.Name)
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedistanceindicationtype.GongGetIdentifier(stage), "CodeDistanceIndicationType", codedistanceindicationtype.Name)
}

func (codeemissionbandtype *CodeEmissionBandType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeemissionbandtype.GongGetIdentifier(stage), "CodeEmissionBandType", codeemissionbandtype.Name)
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeilsbackcoursetype.GongGetIdentifier(stage), "CodeILSBackCourseType", codeilsbackcoursetype.Name)
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeintegritylevelilstype.GongGetIdentifier(stage), "CodeIntegrityLevelILSType", codeintegritylevelilstype.Name)
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

func (codenorthreferencetype *CodeNorthReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenorthreferencetype.GongGetIdentifier(stage), "CodeNorthReferenceType", codenorthreferencetype.Name)
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenotepurposetype.GongGetIdentifier(stage), "CodeNotePurposeType", codenotepurposetype.Name)
}

func (codepartype *CodePARType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepartype.GongGetIdentifier(stage), "CodePARType", codepartype.Name)
}

func (codepositioninilstype *CodePositionInILSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codepositioninilstype.GongGetIdentifier(stage), "CodePositionInILSType", codepositioninilstype.Name)
}

func (codeprimaryradartype *CodePrimaryRadarType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeprimaryradartype.GongGetIdentifier(stage), "CodePrimaryRadarType", codeprimaryradartype.Name)
}

func (coderadarservicetype *CodeRadarServiceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderadarservicetype.GongGetIdentifier(stage), "CodeRadarServiceType", coderadarservicetype.Name)
}

func (coderadioemissiontype *CodeRadioEmissionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderadioemissiontype.GongGetIdentifier(stage), "CodeRadioEmissionType", coderadioemissiontype.Name)
}

func (coderadiosignaltype *CodeRadioSignalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderadiosignaltype.GongGetIdentifier(stage), "CodeRadioSignalType", coderadiosignaltype.Name)
}

func (codereflectortype *CodeReflectorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codereflectortype.GongGetIdentifier(stage), "CodeReflectorType", codereflectortype.Name)
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesignalperformanceilstype.GongGetIdentifier(stage), "CodeSignalPerformanceILSType", codesignalperformanceilstype.Name)
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

func (codestandbypowertype *CodeStandbyPowerType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestandbypowertype.GongGetIdentifier(stage), "CodeStandbyPowerType", codestandbypowertype.Name)
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codestatusnavaidtype.GongGetIdentifier(stage), "CodeStatusNavaidType", codestatusnavaidtype.Name)
}

func (codetacanchanneltype *CodeTACANChannelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetacanchanneltype.GongGetIdentifier(stage), "CodeTACANChannelType", codetacanchanneltype.Name)
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

func (codevortype *CodeVORType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codevortype.GongGetIdentifier(stage), "CodeVORType", codevortype.Name)
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticaldatumtype.GongGetIdentifier(stage), "CodeVerticalDatumType", codeverticaldatumtype.Name)
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeverticalreferencetype.GongGetIdentifier(stage), "CodeVerticalReferenceType", codeverticalreferencetype.Name)
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

func (nosequencetype *NoSequenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(nosequencetype.GongGetIdentifier(stage), "NoSequenceType", nosequencetype.Name)
}

func (textaddresstype *TextAddressType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textaddresstype.GongGetIdentifier(stage), "TextAddressType", textaddresstype.Name)
}

func (textdesignatortype *TextDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textdesignatortype.GongGetIdentifier(stage), "TextDesignatorType", textdesignatortype.Name)
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

func (valmagneticvariationtype *ValMagneticVariationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valmagneticvariationtype.GongGetIdentifier(stage), "ValMagneticVariationType", valmagneticvariationtype.Name)
}

// insertion point for unstaging
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacepointpositiontype.GongGetReferenceIdentifier(stage))
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeairspacepointroletype.GongGetReferenceIdentifier(stage))
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codealtitudeusetype.GongGetReferenceIdentifier(stage))
}

func (codeauralmorsetype *CodeAuralMorseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeauralmorsetype.GongGetReferenceIdentifier(stage))
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeauthorityroletype.GongGetReferenceIdentifier(stage))
}

func (codebearingtype *CodeBearingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codebearingtype.GongGetReferenceIdentifier(stage))
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecardinaldirectiontype.GongGetReferenceIdentifier(stage))
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecheckpointcategorytype.GongGetReferenceIdentifier(stage))
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecoursequalityilstype.GongGetReferenceIdentifier(stage))
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

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedesignatedpointdesignatortype.GongGetReferenceIdentifier(stage))
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedesignatedpointtype.GongGetReferenceIdentifier(stage))
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedirectionreferencetype.GongGetReferenceIdentifier(stage))
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedistanceindicationtype.GongGetReferenceIdentifier(stage))
}

func (codeemissionbandtype *CodeEmissionBandType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeemissionbandtype.GongGetReferenceIdentifier(stage))
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeilsbackcoursetype.GongGetReferenceIdentifier(stage))
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeintegritylevelilstype.GongGetReferenceIdentifier(stage))
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

func (codenorthreferencetype *CodeNorthReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenorthreferencetype.GongGetReferenceIdentifier(stage))
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenotepurposetype.GongGetReferenceIdentifier(stage))
}

func (codepartype *CodePARType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepartype.GongGetReferenceIdentifier(stage))
}

func (codepositioninilstype *CodePositionInILSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codepositioninilstype.GongGetReferenceIdentifier(stage))
}

func (codeprimaryradartype *CodePrimaryRadarType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeprimaryradartype.GongGetReferenceIdentifier(stage))
}

func (coderadarservicetype *CodeRadarServiceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderadarservicetype.GongGetReferenceIdentifier(stage))
}

func (coderadioemissiontype *CodeRadioEmissionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderadioemissiontype.GongGetReferenceIdentifier(stage))
}

func (coderadiosignaltype *CodeRadioSignalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderadiosignaltype.GongGetReferenceIdentifier(stage))
}

func (codereflectortype *CodeReflectorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codereflectortype.GongGetReferenceIdentifier(stage))
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesignalperformanceilstype.GongGetReferenceIdentifier(stage))
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

func (codestandbypowertype *CodeStandbyPowerType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestandbypowertype.GongGetReferenceIdentifier(stage))
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codestatusnavaidtype.GongGetReferenceIdentifier(stage))
}

func (codetacanchanneltype *CodeTACANChannelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetacanchanneltype.GongGetReferenceIdentifier(stage))
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

func (codevortype *CodeVORType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codevortype.GongGetReferenceIdentifier(stage))
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticaldatumtype.GongGetReferenceIdentifier(stage))
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeverticalreferencetype.GongGetReferenceIdentifier(stage))
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

func (nosequencetype *NoSequenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(nosequencetype.GongGetReferenceIdentifier(stage))
}

func (textaddresstype *TextAddressType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textaddresstype.GongGetReferenceIdentifier(stage))
}

func (textdesignatortype *TextDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textdesignatortype.GongGetReferenceIdentifier(stage))
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

func (valmagneticvariationtype *ValMagneticVariationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valmagneticvariationtype.GongGetReferenceIdentifier(stage))
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
