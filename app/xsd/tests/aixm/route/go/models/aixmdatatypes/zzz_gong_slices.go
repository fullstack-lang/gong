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

	res = __gong__appendInstances(res, stage.CodeAircraftICAOTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftWingspanClassTypes)

	res = __gong__appendInstances(res, stage.CodeAltitudeAdjustmentTypes)

	res = __gong__appendInstances(res, stage.CodeAltitudeUseTypes)

	res = __gong__appendInstances(res, stage.CodeAuthorityTypes)

	res = __gong__appendInstances(res, stage.CodeCardinalDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationModeTypes)

	res = __gong__appendInstances(res, stage.CodeComparisonTypes)

	res = __gong__appendInstances(res, stage.CodeCourseTypes)

	res = __gong__appendInstances(res, stage.CodeDayTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionTurnTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeEquipmentAntiCollisionTypes)

	res = __gong__appendInstances(res, stage.CodeFlightOriginTypes)

	res = __gong__appendInstances(res, stage.CodeFlightPurposeTypes)

	res = __gong__appendInstances(res, stage.CodeFlightRestrictionDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeFlightRestrictionTypes)

	res = __gong__appendInstances(res, stage.CodeFlightRuleTypes)

	res = __gong__appendInstances(res, stage.CodeFlightStatusTypes)

	res = __gong__appendInstances(res, stage.CodeFlightTypes)

	res = __gong__appendInstances(res, stage.CodeFlowConditionOperationTypes)

	res = __gong__appendInstances(res, stage.CodeFreeFlightTypes)

	res = __gong__appendInstances(res, stage.CodeLevelTypes)

	res = __gong__appendInstances(res, stage.CodeLocationQualifierTypes)

	res = __gong__appendInstances(res, stage.CodeMeteoConditionsTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryRoutePointTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryStatusTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryTrainingTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationEquipmentTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationSpecificationTypes)

	res = __gong__appendInstances(res, stage.CodeNotePurposeTypes)

	res = __gong__appendInstances(res, stage.CodeObstacleAssessmentSurfaceTypes)

	res = __gong__appendInstances(res, stage.CodeObstructionIdSurfaceZoneTypes)

	res = __gong__appendInstances(res, stage.CodeRNPTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMPointRoleTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMTypes)

	res = __gong__appendInstances(res, stage.CodeReferenceRoleTypes)

	res = __gong__appendInstances(res, stage.CodeRouteAvailabilityTypes)

	res = __gong__appendInstances(res, stage.CodeRouteDesignatorLetterTypes)

	res = __gong__appendInstances(res, stage.CodeRouteDesignatorPrefixTypes)

	res = __gong__appendInstances(res, stage.CodeRouteDesignatorSuffixTypes)

	res = __gong__appendInstances(res, stage.CodeRouteNavigationTypes)

	res = __gong__appendInstances(res, stage.CodeRouteOriginTypes)

	res = __gong__appendInstances(res, stage.CodeRouteSegmentPathTypes)

	res = __gong__appendInstances(res, stage.CodeRouteTypes)

	res = __gong__appendInstances(res, stage.CodeRuleProcedureTitleTypes)

	res = __gong__appendInstances(res, stage.CodeRuleProcedureTypes)

	res = __gong__appendInstances(res, stage.CodeSideTypes)

	res = __gong__appendInstances(res, stage.CodeSpeedReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeTACANChannelTypes)

	res = __gong__appendInstances(res, stage.CodeTelecomNetworkTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventCombinationTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventTypes)

	res = __gong__appendInstances(res, stage.CodeTimeReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeTransponderTypes)

	res = __gong__appendInstances(res, stage.CodeUpperAlphaTypes)

	res = __gong__appendInstances(res, stage.CodeValueInterpretationTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalDatumTypes)

	res = __gong__appendInstances(res, stage.CodeVerticalReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeWakeTurbulenceTypes)

	res = __gong__appendInstances(res, stage.CodeYesNoTypes)

	res = __gong__appendInstances(res, stage.DateMonthDayTypes)

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

	res = __gong__appendInstances(res, stage.ValDistanceSignedTypes)

	res = __gong__appendInstances(res, stage.ValDistanceTypes)

	res = __gong__appendInstances(res, stage.ValDistanceVerticalTypes)

	res = __gong__appendInstances(res, stage.ValDurationTypes)

	res = __gong__appendInstances(res, stage.ValSlopeTypes)

	res = __gong__appendInstances(res, stage.ValSpeedTypes)

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

func (codeaircrafticaotype *CodeAircraftICAOType) GongCopy() GongstructIF {
	newInstance := new(CodeAircraftICAOType)
	codeaircrafticaotype.GongCopyBasicFields(newInstance)
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

func (codeauthoritytype *CodeAuthorityType) GongCopy() GongstructIF {
	newInstance := new(CodeAuthorityType)
	codeauthoritytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongCopy() GongstructIF {
	newInstance := new(CodeCardinalDirectionType)
	codecardinaldirectiontype.GongCopyBasicFields(newInstance)
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

func (codecoursetype *CodeCourseType) GongCopy() GongstructIF {
	newInstance := new(CodeCourseType)
	codecoursetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codedaytype *CodeDayType) GongCopy() GongstructIF {
	newInstance := new(CodeDayType)
	codedaytype.GongCopyBasicFields(newInstance)
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

func (codeleveltype *CodeLevelType) GongCopy() GongstructIF {
	newInstance := new(CodeLevelType)
	codeleveltype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongCopy() GongstructIF {
	newInstance := new(CodeLocationQualifierType)
	codelocationqualifiertype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongCopy() GongstructIF {
	newInstance := new(CodeMeteoConditionsType)
	codemeteoconditionstype.GongCopyBasicFields(newInstance)
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

func (codernptype *CodeRNPType) GongCopy() GongstructIF {
	newInstance := new(CodeRNPType)
	codernptype.GongCopyBasicFields(newInstance)
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

func (codereferenceroletype *CodeReferenceRoleType) GongCopy() GongstructIF {
	newInstance := new(CodeReferenceRoleType)
	codereferenceroletype.GongCopyBasicFields(newInstance)
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

func (codesidetype *CodeSideType) GongCopy() GongstructIF {
	newInstance := new(CodeSideType)
	codesidetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeSpeedReferenceType)
	codespeedreferencetype.GongCopyBasicFields(newInstance)
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

func (codeupperalphatype *CodeUpperAlphaType) GongCopy() GongstructIF {
	newInstance := new(CodeUpperAlphaType)
	codeupperalphatype.GongCopyBasicFields(newInstance)
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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongCopy() GongstructIF {
	newInstance := new(CodeWakeTurbulenceType)
	codewaketurbulencetype.GongCopyBasicFields(newInstance)
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

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircrafticaotype)
}

func (codeaircrafttype *CodeAircraftType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircrafttype)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaircraftwingspanclasstype)
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codealtitudeadjustmenttype)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codealtitudeusetype)
}

func (codeauthoritytype *CodeAuthorityType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeauthoritytype)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecardinaldirectiontype)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecommunicationmodetype)
}

func (codecomparisontype *CodeComparisonType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecomparisontype)
}

func (codecoursetype *CodeCourseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecoursetype)
}

func (codedaytype *CodeDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedaytype)
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

func (codeleveltype *CodeLevelType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeleveltype)
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codelocationqualifiertype)
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemeteoconditionstype)
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

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationequipmenttype)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenavigationspecificationtype)
}

func (codenotepurposetype *CodeNotePurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenotepurposetype)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstacleassessmentsurfacetype)
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstructionidsurfacezonetype)
}

func (codernptype *CodeRNPType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codernptype)
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmpointroletype)
}

func (codervsmtype *CodeRVSMType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmtype)
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codereferenceroletype)
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

func (codesidetype *CodeSideType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesidetype)
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codespeedreferencetype)
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

func (codeupperalphatype *CodeUpperAlphaType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeupperalphatype)
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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codewaketurbulencetype)
}

func (codeyesnotype *CodeYesNoType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeyesnotype)
}

func (datemonthdaytype *DateMonthDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datemonthdaytype)
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

func (valslopetype *ValSlopeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valslopetype)
}

func (valspeedtype *ValSpeedType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valspeedtype)
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

	__gong__computeReferencePass1(stage, stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_reference, &stage.CodeAircraftICAOTypes_referenceOrder, &stage.CodeAircraftICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftTypes, &stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_referenceOrder, &stage.CodeAircraftTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_referenceOrder, &stage.CodeAircraftWingspanClassTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAltitudeAdjustmentTypes, &stage.CodeAltitudeAdjustmentTypes_reference, &stage.CodeAltitudeAdjustmentTypes_referenceOrder, &stage.CodeAltitudeAdjustmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_referenceOrder, &stage.CodeAltitudeUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAuthorityTypes, &stage.CodeAuthorityTypes_reference, &stage.CodeAuthorityTypes_referenceOrder, &stage.CodeAuthorityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCardinalDirectionTypes, &stage.CodeCardinalDirectionTypes_reference, &stage.CodeCardinalDirectionTypes_referenceOrder, &stage.CodeCardinalDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_referenceOrder, &stage.CodeCommunicationModeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeComparisonTypes, &stage.CodeComparisonTypes_reference, &stage.CodeComparisonTypes_referenceOrder, &stage.CodeComparisonTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCourseTypes, &stage.CodeCourseTypes_reference, &stage.CodeCourseTypes_referenceOrder, &stage.CodeCourseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDayTypes, &stage.CodeDayTypes_reference, &stage.CodeDayTypes_referenceOrder, &stage.CodeDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_reference, &stage.CodeDirectionTurnTypes_referenceOrder, &stage.CodeDirectionTurnTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionTypes, &stage.CodeDirectionTypes_reference, &stage.CodeDirectionTypes_referenceOrder, &stage.CodeDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder, &stage.CodeEquipmentAntiCollisionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightOriginTypes, &stage.CodeFlightOriginTypes_reference, &stage.CodeFlightOriginTypes_referenceOrder, &stage.CodeFlightOriginTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightPurposeTypes, &stage.CodeFlightPurposeTypes_reference, &stage.CodeFlightPurposeTypes_referenceOrder, &stage.CodeFlightPurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightRestrictionDesignatorTypes, &stage.CodeFlightRestrictionDesignatorTypes_reference, &stage.CodeFlightRestrictionDesignatorTypes_referenceOrder, &stage.CodeFlightRestrictionDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightRestrictionTypes, &stage.CodeFlightRestrictionTypes_reference, &stage.CodeFlightRestrictionTypes_referenceOrder, &stage.CodeFlightRestrictionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightRuleTypes, &stage.CodeFlightRuleTypes_reference, &stage.CodeFlightRuleTypes_referenceOrder, &stage.CodeFlightRuleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightStatusTypes, &stage.CodeFlightStatusTypes_reference, &stage.CodeFlightStatusTypes_referenceOrder, &stage.CodeFlightStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlightTypes, &stage.CodeFlightTypes_reference, &stage.CodeFlightTypes_referenceOrder, &stage.CodeFlightTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFlowConditionOperationTypes, &stage.CodeFlowConditionOperationTypes_reference, &stage.CodeFlowConditionOperationTypes_referenceOrder, &stage.CodeFlowConditionOperationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFreeFlightTypes, &stage.CodeFreeFlightTypes_reference, &stage.CodeFreeFlightTypes_referenceOrder, &stage.CodeFreeFlightTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLevelTypes, &stage.CodeLevelTypes_reference, &stage.CodeLevelTypes_referenceOrder, &stage.CodeLevelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeLocationQualifierTypes, &stage.CodeLocationQualifierTypes_reference, &stage.CodeLocationQualifierTypes_referenceOrder, &stage.CodeLocationQualifierTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMeteoConditionsTypes, &stage.CodeMeteoConditionsTypes_reference, &stage.CodeMeteoConditionsTypes_referenceOrder, &stage.CodeMeteoConditionsTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryRoutePointTypes, &stage.CodeMilitaryRoutePointTypes_reference, &stage.CodeMilitaryRoutePointTypes_referenceOrder, &stage.CodeMilitaryRoutePointTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryStatusTypes, &stage.CodeMilitaryStatusTypes_reference, &stage.CodeMilitaryStatusTypes_referenceOrder, &stage.CodeMilitaryStatusTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryTrainingTypes, &stage.CodeMilitaryTrainingTypes_reference, &stage.CodeMilitaryTrainingTypes_referenceOrder, &stage.CodeMilitaryTrainingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_referenceOrder, &stage.CodeNavigationEquipmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_referenceOrder, &stage.CodeNavigationSpecificationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_referenceOrder, &stage.CodeNotePurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_reference, &stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder, &stage.CodeObstacleAssessmentSurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstructionIdSurfaceZoneTypes, &stage.CodeObstructionIdSurfaceZoneTypes_reference, &stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder, &stage.CodeObstructionIdSurfaceZoneTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRNPTypes, &stage.CodeRNPTypes_reference, &stage.CodeRNPTypes_referenceOrder, &stage.CodeRNPTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMPointRoleTypes, &stage.CodeRVSMPointRoleTypes_reference, &stage.CodeRVSMPointRoleTypes_referenceOrder, &stage.CodeRVSMPointRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMTypes, &stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_referenceOrder, &stage.CodeRVSMTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeReferenceRoleTypes, &stage.CodeReferenceRoleTypes_reference, &stage.CodeReferenceRoleTypes_referenceOrder, &stage.CodeReferenceRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteAvailabilityTypes, &stage.CodeRouteAvailabilityTypes_reference, &stage.CodeRouteAvailabilityTypes_referenceOrder, &stage.CodeRouteAvailabilityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteDesignatorLetterTypes, &stage.CodeRouteDesignatorLetterTypes_reference, &stage.CodeRouteDesignatorLetterTypes_referenceOrder, &stage.CodeRouteDesignatorLetterTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteDesignatorPrefixTypes, &stage.CodeRouteDesignatorPrefixTypes_reference, &stage.CodeRouteDesignatorPrefixTypes_referenceOrder, &stage.CodeRouteDesignatorPrefixTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteDesignatorSuffixTypes, &stage.CodeRouteDesignatorSuffixTypes_reference, &stage.CodeRouteDesignatorSuffixTypes_referenceOrder, &stage.CodeRouteDesignatorSuffixTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteNavigationTypes, &stage.CodeRouteNavigationTypes_reference, &stage.CodeRouteNavigationTypes_referenceOrder, &stage.CodeRouteNavigationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteOriginTypes, &stage.CodeRouteOriginTypes_reference, &stage.CodeRouteOriginTypes_referenceOrder, &stage.CodeRouteOriginTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteSegmentPathTypes, &stage.CodeRouteSegmentPathTypes_reference, &stage.CodeRouteSegmentPathTypes_referenceOrder, &stage.CodeRouteSegmentPathTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteTypes, &stage.CodeRouteTypes_reference, &stage.CodeRouteTypes_referenceOrder, &stage.CodeRouteTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRuleProcedureTitleTypes, &stage.CodeRuleProcedureTitleTypes_reference, &stage.CodeRuleProcedureTitleTypes_referenceOrder, &stage.CodeRuleProcedureTitleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRuleProcedureTypes, &stage.CodeRuleProcedureTypes_reference, &stage.CodeRuleProcedureTypes_referenceOrder, &stage.CodeRuleProcedureTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSideTypes, &stage.CodeSideTypes_reference, &stage.CodeSideTypes_referenceOrder, &stage.CodeSideTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpeedReferenceTypes, &stage.CodeSpeedReferenceTypes_reference, &stage.CodeSpeedReferenceTypes_referenceOrder, &stage.CodeSpeedReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTACANChannelTypes, &stage.CodeTACANChannelTypes_reference, &stage.CodeTACANChannelTypes_referenceOrder, &stage.CodeTACANChannelTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTelecomNetworkTypes, &stage.CodeTelecomNetworkTypes_reference, &stage.CodeTelecomNetworkTypes_referenceOrder, &stage.CodeTelecomNetworkTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_referenceOrder, &stage.CodeTimeEventCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_referenceOrder, &stage.CodeTimeEventTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_referenceOrder, &stage.CodeTimeReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTransponderTypes, &stage.CodeTransponderTypes_reference, &stage.CodeTransponderTypes_referenceOrder, &stage.CodeTransponderTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeUpperAlphaTypes, &stage.CodeUpperAlphaTypes_reference, &stage.CodeUpperAlphaTypes_referenceOrder, &stage.CodeUpperAlphaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeValueInterpretationTypes, &stage.CodeValueInterpretationTypes_reference, &stage.CodeValueInterpretationTypes_referenceOrder, &stage.CodeValueInterpretationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalDatumTypes, &stage.CodeVerticalDatumTypes_reference, &stage.CodeVerticalDatumTypes_referenceOrder, &stage.CodeVerticalDatumTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeVerticalReferenceTypes, &stage.CodeVerticalReferenceTypes_reference, &stage.CodeVerticalReferenceTypes_referenceOrder, &stage.CodeVerticalReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeWakeTurbulenceTypes, &stage.CodeWakeTurbulenceTypes_reference, &stage.CodeWakeTurbulenceTypes_referenceOrder, &stage.CodeWakeTurbulenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeYesNoTypes, &stage.CodeYesNoTypes_reference, &stage.CodeYesNoTypes_referenceOrder, &stage.CodeYesNoTypes_instance)

	__gong__computeReferencePass1(stage, stage.DateMonthDayTypes, &stage.DateMonthDayTypes_reference, &stage.DateMonthDayTypes_referenceOrder, &stage.DateMonthDayTypes_instance)

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

	__gong__computeReferencePass1(stage, stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_referenceOrder, &stage.ValDistanceSignedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceTypes, &stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_referenceOrder, &stage.ValDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_referenceOrder, &stage.ValDistanceVerticalTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDurationTypes, &stage.ValDurationTypes_reference, &stage.ValDurationTypes_referenceOrder, &stage.ValDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValSlopeTypes, &stage.ValSlopeTypes_reference, &stage.ValSlopeTypes_referenceOrder, &stage.ValSlopeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValSpeedTypes, &stage.ValSpeedTypes_reference, &stage.ValSpeedTypes_referenceOrder, &stage.ValSpeedTypes_instance)

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

	__gong__computeReferencePass2(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftTypes, stage.CodeAircraftTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAltitudeAdjustmentTypes, stage.CodeAltitudeAdjustmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAuthorityTypes, stage.CodeAuthorityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCardinalDirectionTypes, stage.CodeCardinalDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeComparisonTypes, stage.CodeComparisonTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCourseTypes, stage.CodeCourseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDayTypes, stage.CodeDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionTypes, stage.CodeDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightOriginTypes, stage.CodeFlightOriginTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightPurposeTypes, stage.CodeFlightPurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightRestrictionDesignatorTypes, stage.CodeFlightRestrictionDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightRestrictionTypes, stage.CodeFlightRestrictionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightRuleTypes, stage.CodeFlightRuleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightStatusTypes, stage.CodeFlightStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlightTypes, stage.CodeFlightTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFlowConditionOperationTypes, stage.CodeFlowConditionOperationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFreeFlightTypes, stage.CodeFreeFlightTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLevelTypes, stage.CodeLevelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeLocationQualifierTypes, stage.CodeLocationQualifierTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMeteoConditionsTypes, stage.CodeMeteoConditionsTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryRoutePointTypes, stage.CodeMilitaryRoutePointTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryStatusTypes, stage.CodeMilitaryStatusTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryTrainingTypes, stage.CodeMilitaryTrainingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstructionIdSurfaceZoneTypes, stage.CodeObstructionIdSurfaceZoneTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRNPTypes, stage.CodeRNPTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMPointRoleTypes, stage.CodeRVSMPointRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMTypes, stage.CodeRVSMTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeReferenceRoleTypes, stage.CodeReferenceRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteAvailabilityTypes, stage.CodeRouteAvailabilityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteDesignatorLetterTypes, stage.CodeRouteDesignatorLetterTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteDesignatorPrefixTypes, stage.CodeRouteDesignatorPrefixTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteDesignatorSuffixTypes, stage.CodeRouteDesignatorSuffixTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteNavigationTypes, stage.CodeRouteNavigationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteOriginTypes, stage.CodeRouteOriginTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteSegmentPathTypes, stage.CodeRouteSegmentPathTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteTypes, stage.CodeRouteTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRuleProcedureTitleTypes, stage.CodeRuleProcedureTitleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRuleProcedureTypes, stage.CodeRuleProcedureTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSideTypes, stage.CodeSideTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpeedReferenceTypes, stage.CodeSpeedReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTACANChannelTypes, stage.CodeTACANChannelTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTelecomNetworkTypes, stage.CodeTelecomNetworkTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTransponderTypes, stage.CodeTransponderTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeUpperAlphaTypes, stage.CodeUpperAlphaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeValueInterpretationTypes, stage.CodeValueInterpretationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalDatumTypes, stage.CodeVerticalDatumTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeVerticalReferenceTypes, stage.CodeVerticalReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeWakeTurbulenceTypes, stage.CodeWakeTurbulenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeYesNoTypes, stage.CodeYesNoTypes_reference, stage)

	__gong__computeReferencePass2(stage.DateMonthDayTypes, stage.DateMonthDayTypes_reference, stage)

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

	__gong__computeReferencePass2(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceTypes, stage.ValDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDurationTypes, stage.ValDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValSlopeTypes, stage.ValSlopeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValSpeedTypes, stage.ValSpeedTypes_reference, stage)

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

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftICAOType_stagedOrder, stage.CodeAircraftICAOTypes_referenceOrder, codeaircrafticaotype, "CodeAircraftICAOType")
}

func (codeaircrafttype *CodeAircraftType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftType_stagedOrder, stage.CodeAircraftTypes_referenceOrder, codeaircrafttype, "CodeAircraftType")
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAircraftWingspanClassType_stagedOrder, stage.CodeAircraftWingspanClassTypes_referenceOrder, codeaircraftwingspanclasstype, "CodeAircraftWingspanClassType")
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAltitudeAdjustmentType_stagedOrder, stage.CodeAltitudeAdjustmentTypes_referenceOrder, codealtitudeadjustmenttype, "CodeAltitudeAdjustmentType")
}

func (codealtitudeusetype *CodeAltitudeUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAltitudeUseType_stagedOrder, stage.CodeAltitudeUseTypes_referenceOrder, codealtitudeusetype, "CodeAltitudeUseType")
}

func (codeauthoritytype *CodeAuthorityType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAuthorityType_stagedOrder, stage.CodeAuthorityTypes_referenceOrder, codeauthoritytype, "CodeAuthorityType")
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCardinalDirectionType_stagedOrder, stage.CodeCardinalDirectionTypes_referenceOrder, codecardinaldirectiontype, "CodeCardinalDirectionType")
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCommunicationModeType_stagedOrder, stage.CodeCommunicationModeTypes_referenceOrder, codecommunicationmodetype, "CodeCommunicationModeType")
}

func (codecomparisontype *CodeComparisonType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeComparisonType_stagedOrder, stage.CodeComparisonTypes_referenceOrder, codecomparisontype, "CodeComparisonType")
}

func (codecoursetype *CodeCourseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCourseType_stagedOrder, stage.CodeCourseTypes_referenceOrder, codecoursetype, "CodeCourseType")
}

func (codedaytype *CodeDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDayType_stagedOrder, stage.CodeDayTypes_referenceOrder, codedaytype, "CodeDayType")
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

func (codeleveltype *CodeLevelType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLevelType_stagedOrder, stage.CodeLevelTypes_referenceOrder, codeleveltype, "CodeLevelType")
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeLocationQualifierType_stagedOrder, stage.CodeLocationQualifierTypes_referenceOrder, codelocationqualifiertype, "CodeLocationQualifierType")
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMeteoConditionsType_stagedOrder, stage.CodeMeteoConditionsTypes_referenceOrder, codemeteoconditionstype, "CodeMeteoConditionsType")
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

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationEquipmentType_stagedOrder, stage.CodeNavigationEquipmentTypes_referenceOrder, codenavigationequipmenttype, "CodeNavigationEquipmentType")
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNavigationSpecificationType_stagedOrder, stage.CodeNavigationSpecificationTypes_referenceOrder, codenavigationspecificationtype, "CodeNavigationSpecificationType")
}

func (codenotepurposetype *CodeNotePurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeTypes_referenceOrder, codenotepurposetype, "CodeNotePurposeType")
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstacleAssessmentSurfaceType_stagedOrder, stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder, codeobstacleassessmentsurfacetype, "CodeObstacleAssessmentSurfaceType")
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstructionIdSurfaceZoneType_stagedOrder, stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder, codeobstructionidsurfacezonetype, "CodeObstructionIdSurfaceZoneType")
}

func (codernptype *CodeRNPType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRNPType_stagedOrder, stage.CodeRNPTypes_referenceOrder, codernptype, "CodeRNPType")
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMPointRoleType_stagedOrder, stage.CodeRVSMPointRoleTypes_referenceOrder, codervsmpointroletype, "CodeRVSMPointRoleType")
}

func (codervsmtype *CodeRVSMType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMType_stagedOrder, stage.CodeRVSMTypes_referenceOrder, codervsmtype, "CodeRVSMType")
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeReferenceRoleType_stagedOrder, stage.CodeReferenceRoleTypes_referenceOrder, codereferenceroletype, "CodeReferenceRoleType")
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

func (codesidetype *CodeSideType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSideType_stagedOrder, stage.CodeSideTypes_referenceOrder, codesidetype, "CodeSideType")
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSpeedReferenceType_stagedOrder, stage.CodeSpeedReferenceTypes_referenceOrder, codespeedreferencetype, "CodeSpeedReferenceType")
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

func (codeupperalphatype *CodeUpperAlphaType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeUpperAlphaType_stagedOrder, stage.CodeUpperAlphaTypes_referenceOrder, codeupperalphatype, "CodeUpperAlphaType")
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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeWakeTurbulenceType_stagedOrder, stage.CodeWakeTurbulenceTypes_referenceOrder, codewaketurbulencetype, "CodeWakeTurbulenceType")
}

func (codeyesnotype *CodeYesNoType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeYesNoType_stagedOrder, stage.CodeYesNoTypes_referenceOrder, codeyesnotype, "CodeYesNoType")
}

func (datemonthdaytype *DateMonthDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DateMonthDayType_stagedOrder, stage.DateMonthDayTypes_referenceOrder, datemonthdaytype, "DateMonthDayType")
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

func (valslopetype *ValSlopeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValSlopeType_stagedOrder, stage.ValSlopeTypes_referenceOrder, valslopetype, "ValSlopeType")
}

func (valspeedtype *ValSpeedType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValSpeedType_stagedOrder, stage.ValSpeedTypes_referenceOrder, valspeedtype, "ValSpeedType")
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

func (codeaircrafticaotype *CodeAircraftICAOType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeaircrafticaotype, codeaircrafticaotype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeaircrafticaotype *CodeAircraftICAOType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeaircrafticaotype.GongGetIdentifier(stage)
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

func (codeauthoritytype *CodeAuthorityType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeauthoritytype, codeauthoritytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeauthoritytype *CodeAuthorityType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeauthoritytype.GongGetIdentifier(stage)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecardinaldirectiontype, codecardinaldirectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecardinaldirectiontype *CodeCardinalDirectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecardinaldirectiontype.GongGetIdentifier(stage)
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

func (codecoursetype *CodeCourseType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecoursetype, codecoursetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecoursetype *CodeCourseType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecoursetype.GongGetIdentifier(stage)
}

func (codedaytype *CodeDayType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codedaytype, codedaytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codedaytype *CodeDayType) GongGetReferenceIdentifier(stage *Stage) string {
	return codedaytype.GongGetIdentifier(stage)
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

func (codeleveltype *CodeLevelType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeleveltype, codeleveltype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeleveltype *CodeLevelType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeleveltype.GongGetIdentifier(stage)
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codelocationqualifiertype, codelocationqualifiertype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codelocationqualifiertype *CodeLocationQualifierType) GongGetReferenceIdentifier(stage *Stage) string {
	return codelocationqualifiertype.GongGetIdentifier(stage)
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemeteoconditionstype, codemeteoconditionstype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemeteoconditionstype *CodeMeteoConditionsType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemeteoconditionstype.GongGetIdentifier(stage)
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

func (codernptype *CodeRNPType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codernptype, codernptype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codernptype *CodeRNPType) GongGetReferenceIdentifier(stage *Stage) string {
	return codernptype.GongGetIdentifier(stage)
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

func (codereferenceroletype *CodeReferenceRoleType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codereferenceroletype, codereferenceroletype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codereferenceroletype *CodeReferenceRoleType) GongGetReferenceIdentifier(stage *Stage) string {
	return codereferenceroletype.GongGetIdentifier(stage)
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

func (codesidetype *CodeSideType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codesidetype, codesidetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codesidetype *CodeSideType) GongGetReferenceIdentifier(stage *Stage) string {
	return codesidetype.GongGetIdentifier(stage)
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codespeedreferencetype, codespeedreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codespeedreferencetype *CodeSpeedReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codespeedreferencetype.GongGetIdentifier(stage)
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

func (codeupperalphatype *CodeUpperAlphaType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeupperalphatype, codeupperalphatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeupperalphatype *CodeUpperAlphaType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeupperalphatype.GongGetIdentifier(stage)
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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codewaketurbulencetype, codewaketurbulencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codewaketurbulencetype *CodeWakeTurbulenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codewaketurbulencetype.GongGetIdentifier(stage)
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

func (codeaircrafticaotype *CodeAircraftICAOType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircrafticaotype.GongGetIdentifier(stage), "CodeAircraftICAOType", codeaircrafticaotype.Name)
}

func (codeaircrafttype *CodeAircraftType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircrafttype.GongGetIdentifier(stage), "CodeAircraftType", codeaircrafttype.Name)
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaircraftwingspanclasstype.GongGetIdentifier(stage), "CodeAircraftWingspanClassType", codeaircraftwingspanclasstype.Name)
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codealtitudeadjustmenttype.GongGetIdentifier(stage), "CodeAltitudeAdjustmentType", codealtitudeadjustmenttype.Name)
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codealtitudeusetype.GongGetIdentifier(stage), "CodeAltitudeUseType", codealtitudeusetype.Name)
}

func (codeauthoritytype *CodeAuthorityType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeauthoritytype.GongGetIdentifier(stage), "CodeAuthorityType", codeauthoritytype.Name)
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecardinaldirectiontype.GongGetIdentifier(stage), "CodeCardinalDirectionType", codecardinaldirectiontype.Name)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecommunicationmodetype.GongGetIdentifier(stage), "CodeCommunicationModeType", codecommunicationmodetype.Name)
}

func (codecomparisontype *CodeComparisonType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecomparisontype.GongGetIdentifier(stage), "CodeComparisonType", codecomparisontype.Name)
}

func (codecoursetype *CodeCourseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecoursetype.GongGetIdentifier(stage), "CodeCourseType", codecoursetype.Name)
}

func (codedaytype *CodeDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedaytype.GongGetIdentifier(stage), "CodeDayType", codedaytype.Name)
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

func (codeleveltype *CodeLevelType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeleveltype.GongGetIdentifier(stage), "CodeLevelType", codeleveltype.Name)
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codelocationqualifiertype.GongGetIdentifier(stage), "CodeLocationQualifierType", codelocationqualifiertype.Name)
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemeteoconditionstype.GongGetIdentifier(stage), "CodeMeteoConditionsType", codemeteoconditionstype.Name)
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

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationequipmenttype.GongGetIdentifier(stage), "CodeNavigationEquipmentType", codenavigationequipmenttype.Name)
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenavigationspecificationtype.GongGetIdentifier(stage), "CodeNavigationSpecificationType", codenavigationspecificationtype.Name)
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenotepurposetype.GongGetIdentifier(stage), "CodeNotePurposeType", codenotepurposetype.Name)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstacleassessmentsurfacetype.GongGetIdentifier(stage), "CodeObstacleAssessmentSurfaceType", codeobstacleassessmentsurfacetype.Name)
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstructionidsurfacezonetype.GongGetIdentifier(stage), "CodeObstructionIdSurfaceZoneType", codeobstructionidsurfacezonetype.Name)
}

func (codernptype *CodeRNPType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codernptype.GongGetIdentifier(stage), "CodeRNPType", codernptype.Name)
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmpointroletype.GongGetIdentifier(stage), "CodeRVSMPointRoleType", codervsmpointroletype.Name)
}

func (codervsmtype *CodeRVSMType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmtype.GongGetIdentifier(stage), "CodeRVSMType", codervsmtype.Name)
}

func (codereferenceroletype *CodeReferenceRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codereferenceroletype.GongGetIdentifier(stage), "CodeReferenceRoleType", codereferenceroletype.Name)
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

func (codesidetype *CodeSideType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesidetype.GongGetIdentifier(stage), "CodeSideType", codesidetype.Name)
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codespeedreferencetype.GongGetIdentifier(stage), "CodeSpeedReferenceType", codespeedreferencetype.Name)
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

func (codeupperalphatype *CodeUpperAlphaType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeupperalphatype.GongGetIdentifier(stage), "CodeUpperAlphaType", codeupperalphatype.Name)
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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codewaketurbulencetype.GongGetIdentifier(stage), "CodeWakeTurbulenceType", codewaketurbulencetype.Name)
}

func (codeyesnotype *CodeYesNoType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeyesnotype.GongGetIdentifier(stage), "CodeYesNoType", codeyesnotype.Name)
}

func (datemonthdaytype *DateMonthDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datemonthdaytype.GongGetIdentifier(stage), "DateMonthDayType", datemonthdaytype.Name)
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

func (valslopetype *ValSlopeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valslopetype.GongGetIdentifier(stage), "ValSlopeType", valslopetype.Name)
}

func (valspeedtype *ValSpeedType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valspeedtype.GongGetIdentifier(stage), "ValSpeedType", valspeedtype.Name)
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

func (codeaircrafticaotype *CodeAircraftICAOType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircrafticaotype.GongGetReferenceIdentifier(stage))
}

func (codeaircrafttype *CodeAircraftType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircrafttype.GongGetReferenceIdentifier(stage))
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaircraftwingspanclasstype.GongGetReferenceIdentifier(stage))
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codealtitudeadjustmenttype.GongGetReferenceIdentifier(stage))
}

func (codealtitudeusetype *CodeAltitudeUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codealtitudeusetype.GongGetReferenceIdentifier(stage))
}

func (codeauthoritytype *CodeAuthorityType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeauthoritytype.GongGetReferenceIdentifier(stage))
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecardinaldirectiontype.GongGetReferenceIdentifier(stage))
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecommunicationmodetype.GongGetReferenceIdentifier(stage))
}

func (codecomparisontype *CodeComparisonType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecomparisontype.GongGetReferenceIdentifier(stage))
}

func (codecoursetype *CodeCourseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecoursetype.GongGetReferenceIdentifier(stage))
}

func (codedaytype *CodeDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedaytype.GongGetReferenceIdentifier(stage))
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

func (codeleveltype *CodeLevelType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeleveltype.GongGetReferenceIdentifier(stage))
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codelocationqualifiertype.GongGetReferenceIdentifier(stage))
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemeteoconditionstype.GongGetReferenceIdentifier(stage))
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

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationequipmenttype.GongGetReferenceIdentifier(stage))
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenavigationspecificationtype.GongGetReferenceIdentifier(stage))
}

func (codenotepurposetype *CodeNotePurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenotepurposetype.GongGetReferenceIdentifier(stage))
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstacleassessmentsurfacetype.GongGetReferenceIdentifier(stage))
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstructionidsurfacezonetype.GongGetReferenceIdentifier(stage))
}

func (codernptype *CodeRNPType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codernptype.GongGetReferenceIdentifier(stage))
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmpointroletype.GongGetReferenceIdentifier(stage))
}

func (codervsmtype *CodeRVSMType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmtype.GongGetReferenceIdentifier(stage))
}

func (codereferenceroletype *CodeReferenceRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codereferenceroletype.GongGetReferenceIdentifier(stage))
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

func (codesidetype *CodeSideType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesidetype.GongGetReferenceIdentifier(stage))
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codespeedreferencetype.GongGetReferenceIdentifier(stage))
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

func (codeupperalphatype *CodeUpperAlphaType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeupperalphatype.GongGetReferenceIdentifier(stage))
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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codewaketurbulencetype.GongGetReferenceIdentifier(stage))
}

func (codeyesnotype *CodeYesNoType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeyesnotype.GongGetReferenceIdentifier(stage))
}

func (datemonthdaytype *DateMonthDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datemonthdaytype.GongGetReferenceIdentifier(stage))
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

func (valslopetype *ValSlopeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valslopetype.GongGetReferenceIdentifier(stage))
}

func (valspeedtype *ValSpeedType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valspeedtype.GongGetReferenceIdentifier(stage))
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
