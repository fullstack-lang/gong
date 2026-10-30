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

	res = __gong__appendInstances(res, stage.CodeAircraftCategoryTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftEngineNumberTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftEngineTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftICAOTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftTypes)

	res = __gong__appendInstances(res, stage.CodeAircraftWingspanClassTypes)

	res = __gong__appendInstances(res, stage.CodeAltitudeAdjustmentTypes)

	res = __gong__appendInstances(res, stage.CodeAltitudeUseTypes)

	res = __gong__appendInstances(res, stage.CodeApproachEquipmentAdditionalTypes)

	res = __gong__appendInstances(res, stage.CodeApproachGuidanceTypes)

	res = __gong__appendInstances(res, stage.CodeApproachPrefixTypes)

	res = __gong__appendInstances(res, stage.CodeApproachTypes)

	res = __gong__appendInstances(res, stage.CodeApprovalTypes)

	res = __gong__appendInstances(res, stage.CodeArcDirectionTypes)

	res = __gong__appendInstances(res, stage.CodeBearingTypes)

	res = __gong__appendInstances(res, stage.CodeCommunicationModeTypes)

	res = __gong__appendInstances(res, stage.CodeCourseTypes)

	res = __gong__appendInstances(res, stage.CodeDayTypes)

	res = __gong__appendInstances(res, stage.CodeDesignStandardTypes)

	res = __gong__appendInstances(res, stage.CodeDesignatedPointDesignatorTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeDirectionTurnTypes)

	res = __gong__appendInstances(res, stage.CodeEquipmentAntiCollisionTypes)

	res = __gong__appendInstances(res, stage.CodeEquipmentUnavailableTypes)

	res = __gong__appendInstances(res, stage.CodeFinalGuidanceTypes)

	res = __gong__appendInstances(res, stage.CodeFreeFlightTypes)

	res = __gong__appendInstances(res, stage.CodeHeightReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeHoldingUsageTypes)

	res = __gong__appendInstances(res, stage.CodeHoldingUseTypes)

	res = __gong__appendInstances(res, stage.CodeICAOCountryTypes)

	res = __gong__appendInstances(res, stage.CodeMilitaryRoutePointTypes)

	res = __gong__appendInstances(res, stage.CodeMinimaFinalApproachPathTypes)

	res = __gong__appendInstances(res, stage.CodeMinimumAltitudeTypes)

	res = __gong__appendInstances(res, stage.CodeMinimumHeightTypes)

	res = __gong__appendInstances(res, stage.CodeMissedApproachTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationAreaRestrictionTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationAreaTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationEquipmentTypes)

	res = __gong__appendInstances(res, stage.CodeNavigationSpecificationTypes)

	res = __gong__appendInstances(res, stage.CodeNotePurposeTypes)

	res = __gong__appendInstances(res, stage.CodeObstacleAssessmentSurfaceTypes)

	res = __gong__appendInstances(res, stage.CodeObstructionIdSurfaceZoneTypes)

	res = __gong__appendInstances(res, stage.CodeProcedureAvailabilityTypes)

	res = __gong__appendInstances(res, stage.CodeProcedureCodingStandardTypes)

	res = __gong__appendInstances(res, stage.CodeProcedureDistanceTypes)

	res = __gong__appendInstances(res, stage.CodeProcedureFixRoleTypes)

	res = __gong__appendInstances(res, stage.CodeProcedurePhaseTypes)

	res = __gong__appendInstances(res, stage.CodeRNPTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMPointRoleTypes)

	res = __gong__appendInstances(res, stage.CodeRVSMTypes)

	res = __gong__appendInstances(res, stage.CodeReferencePathIdentifierTypes)

	res = __gong__appendInstances(res, stage.CodeReferenceRoleTypes)

	res = __gong__appendInstances(res, stage.CodeRelativePositionTypes)

	res = __gong__appendInstances(res, stage.CodeRouteIndicatorTypes)

	res = __gong__appendInstances(res, stage.CodeSafeAltitudeTypes)

	res = __gong__appendInstances(res, stage.CodeSegmentPathTypes)

	res = __gong__appendInstances(res, stage.CodeSegmentTerminationTypes)

	res = __gong__appendInstances(res, stage.CodeSideTypes)

	res = __gong__appendInstances(res, stage.CodeSpeedReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeTAATypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventCombinationTypes)

	res = __gong__appendInstances(res, stage.CodeTimeEventTypes)

	res = __gong__appendInstances(res, stage.CodeTimeReferenceTypes)

	res = __gong__appendInstances(res, stage.CodeTrajectoryTypes)

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

	res = __gong__appendInstances(res, stage.TextInstructionTypes)

	res = __gong__appendInstances(res, stage.TextNameTypes)

	res = __gong__appendInstances(res, stage.TextNoteTypes)

	res = __gong__appendInstances(res, stage.TextPropertyNameTypes)

	res = __gong__appendInstances(res, stage.TextSIDSTARDesignatorTypes)

	res = __gong__appendInstances(res, stage.TimeTypes)

	res = __gong__appendInstances(res, stage.ValAlarmLimitTypes)

	res = __gong__appendInstances(res, stage.ValAngleTypes)

	res = __gong__appendInstances(res, stage.ValBearingTypes)

	res = __gong__appendInstances(res, stage.ValChannelNumberTypes)

	res = __gong__appendInstances(res, stage.ValDistanceSignedTypes)

	res = __gong__appendInstances(res, stage.ValDistanceTypes)

	res = __gong__appendInstances(res, stage.ValDistanceVerticalTypes)

	res = __gong__appendInstances(res, stage.ValDurationTypes)

	res = __gong__appendInstances(res, stage.ValHexTypes)

	res = __gong__appendInstances(res, stage.ValSlopeTypes)

	res = __gong__appendInstances(res, stage.ValSpeedTypes)

	res = __gong__appendInstances(res, stage.ValTemperatureTypes)

	res = __gong__appendInstances(res, stage.ValWeightTypes)

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

func (codearcdirectiontype *CodeArcDirectionType) GongCopy() GongstructIF {
	newInstance := new(CodeArcDirectionType)
	codearcdirectiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codebearingtype *CodeBearingType) GongCopy() GongstructIF {
	newInstance := new(CodeBearingType)
	codebearingtype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongCopy() GongstructIF {
	newInstance := new(CodeCommunicationModeType)
	codecommunicationmodetype.GongCopyBasicFields(newInstance)
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

func (codefinalguidancetype *CodeFinalGuidanceType) GongCopy() GongstructIF {
	newInstance := new(CodeFinalGuidanceType)
	codefinalguidancetype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codefreeflighttype *CodeFreeFlightType) GongCopy() GongstructIF {
	newInstance := new(CodeFreeFlightType)
	codefreeflighttype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codeheightreferencetype *CodeHeightReferenceType) GongCopy() GongstructIF {
	newInstance := new(CodeHeightReferenceType)
	codeheightreferencetype.GongCopyBasicFields(newInstance)
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

func (codeicaocountrytype *CodeICAOCountryType) GongCopy() GongstructIF {
	newInstance := new(CodeICAOCountryType)
	codeicaocountrytype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongCopy() GongstructIF {
	newInstance := new(CodeMilitaryRoutePointType)
	codemilitaryroutepointtype.GongCopyBasicFields(newInstance)
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

func (coderelativepositiontype *CodeRelativePositionType) GongCopy() GongstructIF {
	newInstance := new(CodeRelativePositionType)
	coderelativepositiontype.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongCopy() GongstructIF {
	newInstance := new(CodeRouteIndicatorType)
	coderouteindicatortype.GongCopyBasicFields(newInstance)
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

func (codetaatype *CodeTAAType) GongCopy() GongstructIF {
	newInstance := new(CodeTAAType)
	codetaatype.GongCopyBasicFields(newInstance)
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

func (textpropertynametype *TextPropertyNameType) GongCopy() GongstructIF {
	newInstance := new(TextPropertyNameType)
	textpropertynametype.GongCopyBasicFields(newInstance)
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

func (valhextype *ValHexType) GongCopy() GongstructIF {
	newInstance := new(ValHexType)
	valhextype.GongCopyBasicFields(newInstance)
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
func (codeatcreportingtype *CodeATCReportingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeatcreportingtype)
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeaerialrefuellingpointtype)
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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachequipmentadditionaltype)
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeapproachguidancetype)
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

func (codearcdirectiontype *CodeArcDirectionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codearcdirectiontype)
}

func (codebearingtype *CodeBearingType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codebearingtype)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecommunicationmodetype)
}

func (codecoursetype *CodeCourseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codecoursetype)
}

func (codedaytype *CodeDayType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedaytype)
}

func (codedesignstandardtype *CodeDesignStandardType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedesignstandardtype)
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedesignatedpointdesignatortype)
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedirectionreferencetype)
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codedirectionturntype)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeequipmentanticollisiontype)
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeequipmentunavailabletype)
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefinalguidancetype)
}

func (codefreeflighttype *CodeFreeFlightType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codefreeflighttype)
}

func (codeheightreferencetype *CodeHeightReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeheightreferencetype)
}

func (codeholdingusagetype *CodeHoldingUsageType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeholdingusagetype)
}

func (codeholdingusetype *CodeHoldingUseType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeholdingusetype)
}

func (codeicaocountrytype *CodeICAOCountryType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeicaocountrytype)
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codemilitaryroutepointtype)
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

func (codenotepurposetype *CodeNotePurposeType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codenotepurposetype)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstacleassessmentsurfacetype)
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codeobstructionidsurfacezonetype)
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

func (codernptype *CodeRNPType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codernptype)
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmpointroletype)
}

func (codervsmtype *CodeRVSMType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codervsmtype)
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codereferencepathidentifiertype)
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codereferenceroletype)
}

func (coderelativepositiontype *CodeRelativePositionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderelativepositiontype)
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coderouteindicatortype)
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

func (codesidetype *CodeSideType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codesidetype)
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codespeedreferencetype)
}

func (codetaatype *CodeTAAType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, codetaatype)
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

func (textinstructiontype *TextInstructionType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textinstructiontype)
}

func (textnametype *TextNameType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textnametype)
}

func (textnotetype *TextNoteType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textnotetype)
}

func (textpropertynametype *TextPropertyNameType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, textpropertynametype)
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

func (valhextype *ValHexType) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, valhextype)
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
	__gong__computeReferencePass1(stage, stage.CodeATCReportingTypes, &stage.CodeATCReportingTypes_reference, &stage.CodeATCReportingTypes_referenceOrder, &stage.CodeATCReportingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAerialRefuellingPointTypes, &stage.CodeAerialRefuellingPointTypes_reference, &stage.CodeAerialRefuellingPointTypes_referenceOrder, &stage.CodeAerialRefuellingPointTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftCategoryTypes, &stage.CodeAircraftCategoryTypes_reference, &stage.CodeAircraftCategoryTypes_referenceOrder, &stage.CodeAircraftCategoryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftEngineNumberTypes, &stage.CodeAircraftEngineNumberTypes_reference, &stage.CodeAircraftEngineNumberTypes_referenceOrder, &stage.CodeAircraftEngineNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftEngineTypes, &stage.CodeAircraftEngineTypes_reference, &stage.CodeAircraftEngineTypes_referenceOrder, &stage.CodeAircraftEngineTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftICAOTypes, &stage.CodeAircraftICAOTypes_reference, &stage.CodeAircraftICAOTypes_referenceOrder, &stage.CodeAircraftICAOTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftTypes, &stage.CodeAircraftTypes_reference, &stage.CodeAircraftTypes_referenceOrder, &stage.CodeAircraftTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAircraftWingspanClassTypes, &stage.CodeAircraftWingspanClassTypes_reference, &stage.CodeAircraftWingspanClassTypes_referenceOrder, &stage.CodeAircraftWingspanClassTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAltitudeAdjustmentTypes, &stage.CodeAltitudeAdjustmentTypes_reference, &stage.CodeAltitudeAdjustmentTypes_referenceOrder, &stage.CodeAltitudeAdjustmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeAltitudeUseTypes, &stage.CodeAltitudeUseTypes_reference, &stage.CodeAltitudeUseTypes_referenceOrder, &stage.CodeAltitudeUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachEquipmentAdditionalTypes, &stage.CodeApproachEquipmentAdditionalTypes_reference, &stage.CodeApproachEquipmentAdditionalTypes_referenceOrder, &stage.CodeApproachEquipmentAdditionalTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachGuidanceTypes, &stage.CodeApproachGuidanceTypes_reference, &stage.CodeApproachGuidanceTypes_referenceOrder, &stage.CodeApproachGuidanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachPrefixTypes, &stage.CodeApproachPrefixTypes_reference, &stage.CodeApproachPrefixTypes_referenceOrder, &stage.CodeApproachPrefixTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApproachTypes, &stage.CodeApproachTypes_reference, &stage.CodeApproachTypes_referenceOrder, &stage.CodeApproachTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeApprovalTypes, &stage.CodeApprovalTypes_reference, &stage.CodeApprovalTypes_referenceOrder, &stage.CodeApprovalTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeArcDirectionTypes, &stage.CodeArcDirectionTypes_reference, &stage.CodeArcDirectionTypes_referenceOrder, &stage.CodeArcDirectionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeBearingTypes, &stage.CodeBearingTypes_reference, &stage.CodeBearingTypes_referenceOrder, &stage.CodeBearingTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCommunicationModeTypes, &stage.CodeCommunicationModeTypes_reference, &stage.CodeCommunicationModeTypes_referenceOrder, &stage.CodeCommunicationModeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeCourseTypes, &stage.CodeCourseTypes_reference, &stage.CodeCourseTypes_referenceOrder, &stage.CodeCourseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDayTypes, &stage.CodeDayTypes_reference, &stage.CodeDayTypes_referenceOrder, &stage.CodeDayTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDesignStandardTypes, &stage.CodeDesignStandardTypes_reference, &stage.CodeDesignStandardTypes_referenceOrder, &stage.CodeDesignStandardTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDesignatedPointDesignatorTypes, &stage.CodeDesignatedPointDesignatorTypes_reference, &stage.CodeDesignatedPointDesignatorTypes_referenceOrder, &stage.CodeDesignatedPointDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionReferenceTypes, &stage.CodeDirectionReferenceTypes_reference, &stage.CodeDirectionReferenceTypes_referenceOrder, &stage.CodeDirectionReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeDirectionTurnTypes, &stage.CodeDirectionTurnTypes_reference, &stage.CodeDirectionTurnTypes_referenceOrder, &stage.CodeDirectionTurnTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEquipmentAntiCollisionTypes, &stage.CodeEquipmentAntiCollisionTypes_reference, &stage.CodeEquipmentAntiCollisionTypes_referenceOrder, &stage.CodeEquipmentAntiCollisionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeEquipmentUnavailableTypes, &stage.CodeEquipmentUnavailableTypes_reference, &stage.CodeEquipmentUnavailableTypes_referenceOrder, &stage.CodeEquipmentUnavailableTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFinalGuidanceTypes, &stage.CodeFinalGuidanceTypes_reference, &stage.CodeFinalGuidanceTypes_referenceOrder, &stage.CodeFinalGuidanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeFreeFlightTypes, &stage.CodeFreeFlightTypes_reference, &stage.CodeFreeFlightTypes_referenceOrder, &stage.CodeFreeFlightTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHeightReferenceTypes, &stage.CodeHeightReferenceTypes_reference, &stage.CodeHeightReferenceTypes_referenceOrder, &stage.CodeHeightReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHoldingUsageTypes, &stage.CodeHoldingUsageTypes_reference, &stage.CodeHoldingUsageTypes_referenceOrder, &stage.CodeHoldingUsageTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeHoldingUseTypes, &stage.CodeHoldingUseTypes_reference, &stage.CodeHoldingUseTypes_referenceOrder, &stage.CodeHoldingUseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeICAOCountryTypes, &stage.CodeICAOCountryTypes_reference, &stage.CodeICAOCountryTypes_referenceOrder, &stage.CodeICAOCountryTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMilitaryRoutePointTypes, &stage.CodeMilitaryRoutePointTypes_reference, &stage.CodeMilitaryRoutePointTypes_referenceOrder, &stage.CodeMilitaryRoutePointTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMinimaFinalApproachPathTypes, &stage.CodeMinimaFinalApproachPathTypes_reference, &stage.CodeMinimaFinalApproachPathTypes_referenceOrder, &stage.CodeMinimaFinalApproachPathTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMinimumAltitudeTypes, &stage.CodeMinimumAltitudeTypes_reference, &stage.CodeMinimumAltitudeTypes_referenceOrder, &stage.CodeMinimumAltitudeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMinimumHeightTypes, &stage.CodeMinimumHeightTypes_reference, &stage.CodeMinimumHeightTypes_referenceOrder, &stage.CodeMinimumHeightTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeMissedApproachTypes, &stage.CodeMissedApproachTypes_reference, &stage.CodeMissedApproachTypes_referenceOrder, &stage.CodeMissedApproachTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationAreaRestrictionTypes, &stage.CodeNavigationAreaRestrictionTypes_reference, &stage.CodeNavigationAreaRestrictionTypes_referenceOrder, &stage.CodeNavigationAreaRestrictionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationAreaTypes, &stage.CodeNavigationAreaTypes_reference, &stage.CodeNavigationAreaTypes_referenceOrder, &stage.CodeNavigationAreaTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationEquipmentTypes, &stage.CodeNavigationEquipmentTypes_reference, &stage.CodeNavigationEquipmentTypes_referenceOrder, &stage.CodeNavigationEquipmentTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNavigationSpecificationTypes, &stage.CodeNavigationSpecificationTypes_reference, &stage.CodeNavigationSpecificationTypes_referenceOrder, &stage.CodeNavigationSpecificationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeNotePurposeTypes, &stage.CodeNotePurposeTypes_reference, &stage.CodeNotePurposeTypes_referenceOrder, &stage.CodeNotePurposeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstacleAssessmentSurfaceTypes, &stage.CodeObstacleAssessmentSurfaceTypes_reference, &stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder, &stage.CodeObstacleAssessmentSurfaceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeObstructionIdSurfaceZoneTypes, &stage.CodeObstructionIdSurfaceZoneTypes_reference, &stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder, &stage.CodeObstructionIdSurfaceZoneTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedureAvailabilityTypes, &stage.CodeProcedureAvailabilityTypes_reference, &stage.CodeProcedureAvailabilityTypes_referenceOrder, &stage.CodeProcedureAvailabilityTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedureCodingStandardTypes, &stage.CodeProcedureCodingStandardTypes_reference, &stage.CodeProcedureCodingStandardTypes_referenceOrder, &stage.CodeProcedureCodingStandardTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedureDistanceTypes, &stage.CodeProcedureDistanceTypes_reference, &stage.CodeProcedureDistanceTypes_referenceOrder, &stage.CodeProcedureDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedureFixRoleTypes, &stage.CodeProcedureFixRoleTypes_reference, &stage.CodeProcedureFixRoleTypes_referenceOrder, &stage.CodeProcedureFixRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeProcedurePhaseTypes, &stage.CodeProcedurePhaseTypes_reference, &stage.CodeProcedurePhaseTypes_referenceOrder, &stage.CodeProcedurePhaseTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRNPTypes, &stage.CodeRNPTypes_reference, &stage.CodeRNPTypes_referenceOrder, &stage.CodeRNPTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMPointRoleTypes, &stage.CodeRVSMPointRoleTypes_reference, &stage.CodeRVSMPointRoleTypes_referenceOrder, &stage.CodeRVSMPointRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRVSMTypes, &stage.CodeRVSMTypes_reference, &stage.CodeRVSMTypes_referenceOrder, &stage.CodeRVSMTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeReferencePathIdentifierTypes, &stage.CodeReferencePathIdentifierTypes_reference, &stage.CodeReferencePathIdentifierTypes_referenceOrder, &stage.CodeReferencePathIdentifierTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeReferenceRoleTypes, &stage.CodeReferenceRoleTypes_reference, &stage.CodeReferenceRoleTypes_referenceOrder, &stage.CodeReferenceRoleTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRelativePositionTypes, &stage.CodeRelativePositionTypes_reference, &stage.CodeRelativePositionTypes_referenceOrder, &stage.CodeRelativePositionTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeRouteIndicatorTypes, &stage.CodeRouteIndicatorTypes_reference, &stage.CodeRouteIndicatorTypes_referenceOrder, &stage.CodeRouteIndicatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSafeAltitudeTypes, &stage.CodeSafeAltitudeTypes_reference, &stage.CodeSafeAltitudeTypes_referenceOrder, &stage.CodeSafeAltitudeTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSegmentPathTypes, &stage.CodeSegmentPathTypes_reference, &stage.CodeSegmentPathTypes_referenceOrder, &stage.CodeSegmentPathTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSegmentTerminationTypes, &stage.CodeSegmentTerminationTypes_reference, &stage.CodeSegmentTerminationTypes_referenceOrder, &stage.CodeSegmentTerminationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSideTypes, &stage.CodeSideTypes_reference, &stage.CodeSideTypes_referenceOrder, &stage.CodeSideTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeSpeedReferenceTypes, &stage.CodeSpeedReferenceTypes_reference, &stage.CodeSpeedReferenceTypes_referenceOrder, &stage.CodeSpeedReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTAATypes, &stage.CodeTAATypes_reference, &stage.CodeTAATypes_referenceOrder, &stage.CodeTAATypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventCombinationTypes, &stage.CodeTimeEventCombinationTypes_reference, &stage.CodeTimeEventCombinationTypes_referenceOrder, &stage.CodeTimeEventCombinationTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeEventTypes, &stage.CodeTimeEventTypes_reference, &stage.CodeTimeEventTypes_referenceOrder, &stage.CodeTimeEventTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTimeReferenceTypes, &stage.CodeTimeReferenceTypes_reference, &stage.CodeTimeReferenceTypes_referenceOrder, &stage.CodeTimeReferenceTypes_instance)

	__gong__computeReferencePass1(stage, stage.CodeTrajectoryTypes, &stage.CodeTrajectoryTypes_reference, &stage.CodeTrajectoryTypes_referenceOrder, &stage.CodeTrajectoryTypes_instance)

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

	__gong__computeReferencePass1(stage, stage.TextInstructionTypes, &stage.TextInstructionTypes_reference, &stage.TextInstructionTypes_referenceOrder, &stage.TextInstructionTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNameTypes, &stage.TextNameTypes_reference, &stage.TextNameTypes_referenceOrder, &stage.TextNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextNoteTypes, &stage.TextNoteTypes_reference, &stage.TextNoteTypes_referenceOrder, &stage.TextNoteTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextPropertyNameTypes, &stage.TextPropertyNameTypes_reference, &stage.TextPropertyNameTypes_referenceOrder, &stage.TextPropertyNameTypes_instance)

	__gong__computeReferencePass1(stage, stage.TextSIDSTARDesignatorTypes, &stage.TextSIDSTARDesignatorTypes_reference, &stage.TextSIDSTARDesignatorTypes_referenceOrder, &stage.TextSIDSTARDesignatorTypes_instance)

	__gong__computeReferencePass1(stage, stage.TimeTypes, &stage.TimeTypes_reference, &stage.TimeTypes_referenceOrder, &stage.TimeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValAlarmLimitTypes, &stage.ValAlarmLimitTypes_reference, &stage.ValAlarmLimitTypes_referenceOrder, &stage.ValAlarmLimitTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValAngleTypes, &stage.ValAngleTypes_reference, &stage.ValAngleTypes_referenceOrder, &stage.ValAngleTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValBearingTypes, &stage.ValBearingTypes_reference, &stage.ValBearingTypes_referenceOrder, &stage.ValBearingTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValChannelNumberTypes, &stage.ValChannelNumberTypes_reference, &stage.ValChannelNumberTypes_referenceOrder, &stage.ValChannelNumberTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceSignedTypes, &stage.ValDistanceSignedTypes_reference, &stage.ValDistanceSignedTypes_referenceOrder, &stage.ValDistanceSignedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceTypes, &stage.ValDistanceTypes_reference, &stage.ValDistanceTypes_referenceOrder, &stage.ValDistanceTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDistanceVerticalTypes, &stage.ValDistanceVerticalTypes_reference, &stage.ValDistanceVerticalTypes_referenceOrder, &stage.ValDistanceVerticalTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValDurationTypes, &stage.ValDurationTypes_reference, &stage.ValDurationTypes_referenceOrder, &stage.ValDurationTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValHexTypes, &stage.ValHexTypes_reference, &stage.ValHexTypes_referenceOrder, &stage.ValHexTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValSlopeTypes, &stage.ValSlopeTypes_reference, &stage.ValSlopeTypes_referenceOrder, &stage.ValSlopeTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValSpeedTypes, &stage.ValSpeedTypes_reference, &stage.ValSpeedTypes_referenceOrder, &stage.ValSpeedTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValTemperatureTypes, &stage.ValTemperatureTypes_reference, &stage.ValTemperatureTypes_referenceOrder, &stage.ValTemperatureTypes_instance)

	__gong__computeReferencePass1(stage, stage.ValWeightTypes, &stage.ValWeightTypes_reference, &stage.ValWeightTypes_referenceOrder, &stage.ValWeightTypes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.CodeATCReportingTypes, stage.CodeATCReportingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAerialRefuellingPointTypes, stage.CodeAerialRefuellingPointTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftCategoryTypes, stage.CodeAircraftCategoryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftEngineNumberTypes, stage.CodeAircraftEngineNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftEngineTypes, stage.CodeAircraftEngineTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftICAOTypes, stage.CodeAircraftICAOTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftTypes, stage.CodeAircraftTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAircraftWingspanClassTypes, stage.CodeAircraftWingspanClassTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAltitudeAdjustmentTypes, stage.CodeAltitudeAdjustmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeAltitudeUseTypes, stage.CodeAltitudeUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachEquipmentAdditionalTypes, stage.CodeApproachEquipmentAdditionalTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachGuidanceTypes, stage.CodeApproachGuidanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachPrefixTypes, stage.CodeApproachPrefixTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApproachTypes, stage.CodeApproachTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeApprovalTypes, stage.CodeApprovalTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeArcDirectionTypes, stage.CodeArcDirectionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeBearingTypes, stage.CodeBearingTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCommunicationModeTypes, stage.CodeCommunicationModeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeCourseTypes, stage.CodeCourseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDayTypes, stage.CodeDayTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDesignStandardTypes, stage.CodeDesignStandardTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDesignatedPointDesignatorTypes, stage.CodeDesignatedPointDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionReferenceTypes, stage.CodeDirectionReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeDirectionTurnTypes, stage.CodeDirectionTurnTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEquipmentAntiCollisionTypes, stage.CodeEquipmentAntiCollisionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeEquipmentUnavailableTypes, stage.CodeEquipmentUnavailableTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFinalGuidanceTypes, stage.CodeFinalGuidanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeFreeFlightTypes, stage.CodeFreeFlightTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHeightReferenceTypes, stage.CodeHeightReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHoldingUsageTypes, stage.CodeHoldingUsageTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeHoldingUseTypes, stage.CodeHoldingUseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeICAOCountryTypes, stage.CodeICAOCountryTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMilitaryRoutePointTypes, stage.CodeMilitaryRoutePointTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMinimaFinalApproachPathTypes, stage.CodeMinimaFinalApproachPathTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMinimumAltitudeTypes, stage.CodeMinimumAltitudeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMinimumHeightTypes, stage.CodeMinimumHeightTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeMissedApproachTypes, stage.CodeMissedApproachTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationAreaRestrictionTypes, stage.CodeNavigationAreaRestrictionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationAreaTypes, stage.CodeNavigationAreaTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationEquipmentTypes, stage.CodeNavigationEquipmentTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNavigationSpecificationTypes, stage.CodeNavigationSpecificationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeNotePurposeTypes, stage.CodeNotePurposeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstacleAssessmentSurfaceTypes, stage.CodeObstacleAssessmentSurfaceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeObstructionIdSurfaceZoneTypes, stage.CodeObstructionIdSurfaceZoneTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedureAvailabilityTypes, stage.CodeProcedureAvailabilityTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedureCodingStandardTypes, stage.CodeProcedureCodingStandardTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedureDistanceTypes, stage.CodeProcedureDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedureFixRoleTypes, stage.CodeProcedureFixRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeProcedurePhaseTypes, stage.CodeProcedurePhaseTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRNPTypes, stage.CodeRNPTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMPointRoleTypes, stage.CodeRVSMPointRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRVSMTypes, stage.CodeRVSMTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeReferencePathIdentifierTypes, stage.CodeReferencePathIdentifierTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeReferenceRoleTypes, stage.CodeReferenceRoleTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRelativePositionTypes, stage.CodeRelativePositionTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeRouteIndicatorTypes, stage.CodeRouteIndicatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSafeAltitudeTypes, stage.CodeSafeAltitudeTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSegmentPathTypes, stage.CodeSegmentPathTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSegmentTerminationTypes, stage.CodeSegmentTerminationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSideTypes, stage.CodeSideTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeSpeedReferenceTypes, stage.CodeSpeedReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTAATypes, stage.CodeTAATypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventCombinationTypes, stage.CodeTimeEventCombinationTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeEventTypes, stage.CodeTimeEventTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTimeReferenceTypes, stage.CodeTimeReferenceTypes_reference, stage)

	__gong__computeReferencePass2(stage.CodeTrajectoryTypes, stage.CodeTrajectoryTypes_reference, stage)

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

	__gong__computeReferencePass2(stage.TextInstructionTypes, stage.TextInstructionTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNameTypes, stage.TextNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextNoteTypes, stage.TextNoteTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextPropertyNameTypes, stage.TextPropertyNameTypes_reference, stage)

	__gong__computeReferencePass2(stage.TextSIDSTARDesignatorTypes, stage.TextSIDSTARDesignatorTypes_reference, stage)

	__gong__computeReferencePass2(stage.TimeTypes, stage.TimeTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValAlarmLimitTypes, stage.ValAlarmLimitTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValAngleTypes, stage.ValAngleTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValBearingTypes, stage.ValBearingTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValChannelNumberTypes, stage.ValChannelNumberTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceSignedTypes, stage.ValDistanceSignedTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceTypes, stage.ValDistanceTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDistanceVerticalTypes, stage.ValDistanceVerticalTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValDurationTypes, stage.ValDurationTypes_reference, stage)

	__gong__computeReferencePass2(stage.ValHexTypes, stage.ValHexTypes_reference, stage)

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
func (codeatcreportingtype *CodeATCReportingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeATCReportingType_stagedOrder, stage.CodeATCReportingTypes_referenceOrder, codeatcreportingtype, "CodeATCReportingType")
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeAerialRefuellingPointType_stagedOrder, stage.CodeAerialRefuellingPointTypes_referenceOrder, codeaerialrefuellingpointtype, "CodeAerialRefuellingPointType")
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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachEquipmentAdditionalType_stagedOrder, stage.CodeApproachEquipmentAdditionalTypes_referenceOrder, codeapproachequipmentadditionaltype, "CodeApproachEquipmentAdditionalType")
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeApproachGuidanceType_stagedOrder, stage.CodeApproachGuidanceTypes_referenceOrder, codeapproachguidancetype, "CodeApproachGuidanceType")
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

func (codearcdirectiontype *CodeArcDirectionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeArcDirectionType_stagedOrder, stage.CodeArcDirectionTypes_referenceOrder, codearcdirectiontype, "CodeArcDirectionType")
}

func (codebearingtype *CodeBearingType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeBearingType_stagedOrder, stage.CodeBearingTypes_referenceOrder, codebearingtype, "CodeBearingType")
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCommunicationModeType_stagedOrder, stage.CodeCommunicationModeTypes_referenceOrder, codecommunicationmodetype, "CodeCommunicationModeType")
}

func (codecoursetype *CodeCourseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeCourseType_stagedOrder, stage.CodeCourseTypes_referenceOrder, codecoursetype, "CodeCourseType")
}

func (codedaytype *CodeDayType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDayType_stagedOrder, stage.CodeDayTypes_referenceOrder, codedaytype, "CodeDayType")
}

func (codedesignstandardtype *CodeDesignStandardType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDesignStandardType_stagedOrder, stage.CodeDesignStandardTypes_referenceOrder, codedesignstandardtype, "CodeDesignStandardType")
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDesignatedPointDesignatorType_stagedOrder, stage.CodeDesignatedPointDesignatorTypes_referenceOrder, codedesignatedpointdesignatortype, "CodeDesignatedPointDesignatorType")
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDirectionReferenceType_stagedOrder, stage.CodeDirectionReferenceTypes_referenceOrder, codedirectionreferencetype, "CodeDirectionReferenceType")
}

func (codedirectionturntype *CodeDirectionTurnType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeDirectionTurnType_stagedOrder, stage.CodeDirectionTurnTypes_referenceOrder, codedirectionturntype, "CodeDirectionTurnType")
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeEquipmentAntiCollisionType_stagedOrder, stage.CodeEquipmentAntiCollisionTypes_referenceOrder, codeequipmentanticollisiontype, "CodeEquipmentAntiCollisionType")
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeEquipmentUnavailableType_stagedOrder, stage.CodeEquipmentUnavailableTypes_referenceOrder, codeequipmentunavailabletype, "CodeEquipmentUnavailableType")
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFinalGuidanceType_stagedOrder, stage.CodeFinalGuidanceTypes_referenceOrder, codefinalguidancetype, "CodeFinalGuidanceType")
}

func (codefreeflighttype *CodeFreeFlightType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeFreeFlightType_stagedOrder, stage.CodeFreeFlightTypes_referenceOrder, codefreeflighttype, "CodeFreeFlightType")
}

func (codeheightreferencetype *CodeHeightReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHeightReferenceType_stagedOrder, stage.CodeHeightReferenceTypes_referenceOrder, codeheightreferencetype, "CodeHeightReferenceType")
}

func (codeholdingusagetype *CodeHoldingUsageType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHoldingUsageType_stagedOrder, stage.CodeHoldingUsageTypes_referenceOrder, codeholdingusagetype, "CodeHoldingUsageType")
}

func (codeholdingusetype *CodeHoldingUseType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeHoldingUseType_stagedOrder, stage.CodeHoldingUseTypes_referenceOrder, codeholdingusetype, "CodeHoldingUseType")
}

func (codeicaocountrytype *CodeICAOCountryType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeICAOCountryType_stagedOrder, stage.CodeICAOCountryTypes_referenceOrder, codeicaocountrytype, "CodeICAOCountryType")
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeMilitaryRoutePointType_stagedOrder, stage.CodeMilitaryRoutePointTypes_referenceOrder, codemilitaryroutepointtype, "CodeMilitaryRoutePointType")
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

func (codenotepurposetype *CodeNotePurposeType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeNotePurposeType_stagedOrder, stage.CodeNotePurposeTypes_referenceOrder, codenotepurposetype, "CodeNotePurposeType")
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstacleAssessmentSurfaceType_stagedOrder, stage.CodeObstacleAssessmentSurfaceTypes_referenceOrder, codeobstacleassessmentsurfacetype, "CodeObstacleAssessmentSurfaceType")
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeObstructionIdSurfaceZoneType_stagedOrder, stage.CodeObstructionIdSurfaceZoneTypes_referenceOrder, codeobstructionidsurfacezonetype, "CodeObstructionIdSurfaceZoneType")
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

func (codernptype *CodeRNPType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRNPType_stagedOrder, stage.CodeRNPTypes_referenceOrder, codernptype, "CodeRNPType")
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMPointRoleType_stagedOrder, stage.CodeRVSMPointRoleTypes_referenceOrder, codervsmpointroletype, "CodeRVSMPointRoleType")
}

func (codervsmtype *CodeRVSMType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRVSMType_stagedOrder, stage.CodeRVSMTypes_referenceOrder, codervsmtype, "CodeRVSMType")
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeReferencePathIdentifierType_stagedOrder, stage.CodeReferencePathIdentifierTypes_referenceOrder, codereferencepathidentifiertype, "CodeReferencePathIdentifierType")
}

func (codereferenceroletype *CodeReferenceRoleType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeReferenceRoleType_stagedOrder, stage.CodeReferenceRoleTypes_referenceOrder, codereferenceroletype, "CodeReferenceRoleType")
}

func (coderelativepositiontype *CodeRelativePositionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRelativePositionType_stagedOrder, stage.CodeRelativePositionTypes_referenceOrder, coderelativepositiontype, "CodeRelativePositionType")
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeRouteIndicatorType_stagedOrder, stage.CodeRouteIndicatorTypes_referenceOrder, coderouteindicatortype, "CodeRouteIndicatorType")
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

func (codesidetype *CodeSideType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSideType_stagedOrder, stage.CodeSideTypes_referenceOrder, codesidetype, "CodeSideType")
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeSpeedReferenceType_stagedOrder, stage.CodeSpeedReferenceTypes_referenceOrder, codespeedreferencetype, "CodeSpeedReferenceType")
}

func (codetaatype *CodeTAAType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.CodeTAAType_stagedOrder, stage.CodeTAATypes_referenceOrder, codetaatype, "CodeTAAType")
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

func (textinstructiontype *TextInstructionType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextInstructionType_stagedOrder, stage.TextInstructionTypes_referenceOrder, textinstructiontype, "TextInstructionType")
}

func (textnametype *TextNameType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextNameType_stagedOrder, stage.TextNameTypes_referenceOrder, textnametype, "TextNameType")
}

func (textnotetype *TextNoteType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextNoteType_stagedOrder, stage.TextNoteTypes_referenceOrder, textnotetype, "TextNoteType")
}

func (textpropertynametype *TextPropertyNameType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TextPropertyNameType_stagedOrder, stage.TextPropertyNameTypes_referenceOrder, textpropertynametype, "TextPropertyNameType")
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

func (valhextype *ValHexType) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ValHexType_stagedOrder, stage.ValHexTypes_referenceOrder, valhextype, "ValHexType")
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

func (codearcdirectiontype *CodeArcDirectionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codearcdirectiontype, codearcdirectiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codearcdirectiontype *CodeArcDirectionType) GongGetReferenceIdentifier(stage *Stage) string {
	return codearcdirectiontype.GongGetIdentifier(stage)
}

func (codebearingtype *CodeBearingType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codebearingtype, codebearingtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codebearingtype *CodeBearingType) GongGetReferenceIdentifier(stage *Stage) string {
	return codebearingtype.GongGetIdentifier(stage)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codecommunicationmodetype, codecommunicationmodetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codecommunicationmodetype *CodeCommunicationModeType) GongGetReferenceIdentifier(stage *Stage) string {
	return codecommunicationmodetype.GongGetIdentifier(stage)
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

func (codefinalguidancetype *CodeFinalGuidanceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefinalguidancetype, codefinalguidancetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefinalguidancetype *CodeFinalGuidanceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefinalguidancetype.GongGetIdentifier(stage)
}

func (codefreeflighttype *CodeFreeFlightType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codefreeflighttype, codefreeflighttype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codefreeflighttype *CodeFreeFlightType) GongGetReferenceIdentifier(stage *Stage) string {
	return codefreeflighttype.GongGetIdentifier(stage)
}

func (codeheightreferencetype *CodeHeightReferenceType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeheightreferencetype, codeheightreferencetype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeheightreferencetype *CodeHeightReferenceType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeheightreferencetype.GongGetIdentifier(stage)
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

func (codeicaocountrytype *CodeICAOCountryType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codeicaocountrytype, codeicaocountrytype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codeicaocountrytype *CodeICAOCountryType) GongGetReferenceIdentifier(stage *Stage) string {
	return codeicaocountrytype.GongGetIdentifier(stage)
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codemilitaryroutepointtype, codemilitaryroutepointtype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongGetReferenceIdentifier(stage *Stage) string {
	return codemilitaryroutepointtype.GongGetIdentifier(stage)
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

func (coderelativepositiontype *CodeRelativePositionType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderelativepositiontype, coderelativepositiontype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderelativepositiontype *CodeRelativePositionType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderelativepositiontype.GongGetIdentifier(stage)
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coderouteindicatortype, coderouteindicatortype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coderouteindicatortype *CodeRouteIndicatorType) GongGetReferenceIdentifier(stage *Stage) string {
	return coderouteindicatortype.GongGetIdentifier(stage)
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

func (codetaatype *CodeTAAType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(codetaatype, codetaatype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (codetaatype *CodeTAAType) GongGetReferenceIdentifier(stage *Stage) string {
	return codetaatype.GongGetIdentifier(stage)
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

func (textpropertynametype *TextPropertyNameType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(textpropertynametype, textpropertynametype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (textpropertynametype *TextPropertyNameType) GongGetReferenceIdentifier(stage *Stage) string {
	return textpropertynametype.GongGetIdentifier(stage)
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

func (valhextype *ValHexType) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(valhextype, valhextype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (valhextype *ValHexType) GongGetReferenceIdentifier(stage *Stage) string {
	return valhextype.GongGetIdentifier(stage)
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
func (codeatcreportingtype *CodeATCReportingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeatcreportingtype.GongGetIdentifier(stage), "CodeATCReportingType", codeatcreportingtype.Name)
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeaerialrefuellingpointtype.GongGetIdentifier(stage), "CodeAerialRefuellingPointType", codeaerialrefuellingpointtype.Name)
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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachequipmentadditionaltype.GongGetIdentifier(stage), "CodeApproachEquipmentAdditionalType", codeapproachequipmentadditionaltype.Name)
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeapproachguidancetype.GongGetIdentifier(stage), "CodeApproachGuidanceType", codeapproachguidancetype.Name)
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

func (codearcdirectiontype *CodeArcDirectionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codearcdirectiontype.GongGetIdentifier(stage), "CodeArcDirectionType", codearcdirectiontype.Name)
}

func (codebearingtype *CodeBearingType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codebearingtype.GongGetIdentifier(stage), "CodeBearingType", codebearingtype.Name)
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecommunicationmodetype.GongGetIdentifier(stage), "CodeCommunicationModeType", codecommunicationmodetype.Name)
}

func (codecoursetype *CodeCourseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codecoursetype.GongGetIdentifier(stage), "CodeCourseType", codecoursetype.Name)
}

func (codedaytype *CodeDayType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedaytype.GongGetIdentifier(stage), "CodeDayType", codedaytype.Name)
}

func (codedesignstandardtype *CodeDesignStandardType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedesignstandardtype.GongGetIdentifier(stage), "CodeDesignStandardType", codedesignstandardtype.Name)
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedesignatedpointdesignatortype.GongGetIdentifier(stage), "CodeDesignatedPointDesignatorType", codedesignatedpointdesignatortype.Name)
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedirectionreferencetype.GongGetIdentifier(stage), "CodeDirectionReferenceType", codedirectionreferencetype.Name)
}

func (codedirectionturntype *CodeDirectionTurnType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codedirectionturntype.GongGetIdentifier(stage), "CodeDirectionTurnType", codedirectionturntype.Name)
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeequipmentanticollisiontype.GongGetIdentifier(stage), "CodeEquipmentAntiCollisionType", codeequipmentanticollisiontype.Name)
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeequipmentunavailabletype.GongGetIdentifier(stage), "CodeEquipmentUnavailableType", codeequipmentunavailabletype.Name)
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefinalguidancetype.GongGetIdentifier(stage), "CodeFinalGuidanceType", codefinalguidancetype.Name)
}

func (codefreeflighttype *CodeFreeFlightType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codefreeflighttype.GongGetIdentifier(stage), "CodeFreeFlightType", codefreeflighttype.Name)
}

func (codeheightreferencetype *CodeHeightReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeheightreferencetype.GongGetIdentifier(stage), "CodeHeightReferenceType", codeheightreferencetype.Name)
}

func (codeholdingusagetype *CodeHoldingUsageType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeholdingusagetype.GongGetIdentifier(stage), "CodeHoldingUsageType", codeholdingusagetype.Name)
}

func (codeholdingusetype *CodeHoldingUseType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeholdingusetype.GongGetIdentifier(stage), "CodeHoldingUseType", codeholdingusetype.Name)
}

func (codeicaocountrytype *CodeICAOCountryType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeicaocountrytype.GongGetIdentifier(stage), "CodeICAOCountryType", codeicaocountrytype.Name)
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codemilitaryroutepointtype.GongGetIdentifier(stage), "CodeMilitaryRoutePointType", codemilitaryroutepointtype.Name)
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

func (codenotepurposetype *CodeNotePurposeType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codenotepurposetype.GongGetIdentifier(stage), "CodeNotePurposeType", codenotepurposetype.Name)
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstacleassessmentsurfacetype.GongGetIdentifier(stage), "CodeObstacleAssessmentSurfaceType", codeobstacleassessmentsurfacetype.Name)
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codeobstructionidsurfacezonetype.GongGetIdentifier(stage), "CodeObstructionIdSurfaceZoneType", codeobstructionidsurfacezonetype.Name)
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

func (codernptype *CodeRNPType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codernptype.GongGetIdentifier(stage), "CodeRNPType", codernptype.Name)
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmpointroletype.GongGetIdentifier(stage), "CodeRVSMPointRoleType", codervsmpointroletype.Name)
}

func (codervsmtype *CodeRVSMType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codervsmtype.GongGetIdentifier(stage), "CodeRVSMType", codervsmtype.Name)
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codereferencepathidentifiertype.GongGetIdentifier(stage), "CodeReferencePathIdentifierType", codereferencepathidentifiertype.Name)
}

func (codereferenceroletype *CodeReferenceRoleType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codereferenceroletype.GongGetIdentifier(stage), "CodeReferenceRoleType", codereferenceroletype.Name)
}

func (coderelativepositiontype *CodeRelativePositionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderelativepositiontype.GongGetIdentifier(stage), "CodeRelativePositionType", coderelativepositiontype.Name)
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coderouteindicatortype.GongGetIdentifier(stage), "CodeRouteIndicatorType", coderouteindicatortype.Name)
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

func (codesidetype *CodeSideType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codesidetype.GongGetIdentifier(stage), "CodeSideType", codesidetype.Name)
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codespeedreferencetype.GongGetIdentifier(stage), "CodeSpeedReferenceType", codespeedreferencetype.Name)
}

func (codetaatype *CodeTAAType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(codetaatype.GongGetIdentifier(stage), "CodeTAAType", codetaatype.Name)
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

func (textinstructiontype *TextInstructionType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textinstructiontype.GongGetIdentifier(stage), "TextInstructionType", textinstructiontype.Name)
}

func (textnametype *TextNameType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textnametype.GongGetIdentifier(stage), "TextNameType", textnametype.Name)
}

func (textnotetype *TextNoteType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textnotetype.GongGetIdentifier(stage), "TextNoteType", textnotetype.Name)
}

func (textpropertynametype *TextPropertyNameType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(textpropertynametype.GongGetIdentifier(stage), "TextPropertyNameType", textpropertynametype.Name)
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

func (valhextype *ValHexType) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(valhextype.GongGetIdentifier(stage), "ValHexType", valhextype.Name)
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
func (codeatcreportingtype *CodeATCReportingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeatcreportingtype.GongGetReferenceIdentifier(stage))
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeaerialrefuellingpointtype.GongGetReferenceIdentifier(stage))
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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachequipmentadditionaltype.GongGetReferenceIdentifier(stage))
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeapproachguidancetype.GongGetReferenceIdentifier(stage))
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

func (codearcdirectiontype *CodeArcDirectionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codearcdirectiontype.GongGetReferenceIdentifier(stage))
}

func (codebearingtype *CodeBearingType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codebearingtype.GongGetReferenceIdentifier(stage))
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecommunicationmodetype.GongGetReferenceIdentifier(stage))
}

func (codecoursetype *CodeCourseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codecoursetype.GongGetReferenceIdentifier(stage))
}

func (codedaytype *CodeDayType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedaytype.GongGetReferenceIdentifier(stage))
}

func (codedesignstandardtype *CodeDesignStandardType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedesignstandardtype.GongGetReferenceIdentifier(stage))
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedesignatedpointdesignatortype.GongGetReferenceIdentifier(stage))
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedirectionreferencetype.GongGetReferenceIdentifier(stage))
}

func (codedirectionturntype *CodeDirectionTurnType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codedirectionturntype.GongGetReferenceIdentifier(stage))
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeequipmentanticollisiontype.GongGetReferenceIdentifier(stage))
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeequipmentunavailabletype.GongGetReferenceIdentifier(stage))
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefinalguidancetype.GongGetReferenceIdentifier(stage))
}

func (codefreeflighttype *CodeFreeFlightType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codefreeflighttype.GongGetReferenceIdentifier(stage))
}

func (codeheightreferencetype *CodeHeightReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeheightreferencetype.GongGetReferenceIdentifier(stage))
}

func (codeholdingusagetype *CodeHoldingUsageType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeholdingusagetype.GongGetReferenceIdentifier(stage))
}

func (codeholdingusetype *CodeHoldingUseType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeholdingusetype.GongGetReferenceIdentifier(stage))
}

func (codeicaocountrytype *CodeICAOCountryType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeicaocountrytype.GongGetReferenceIdentifier(stage))
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codemilitaryroutepointtype.GongGetReferenceIdentifier(stage))
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

func (codenotepurposetype *CodeNotePurposeType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codenotepurposetype.GongGetReferenceIdentifier(stage))
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstacleassessmentsurfacetype.GongGetReferenceIdentifier(stage))
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codeobstructionidsurfacezonetype.GongGetReferenceIdentifier(stage))
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

func (codernptype *CodeRNPType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codernptype.GongGetReferenceIdentifier(stage))
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmpointroletype.GongGetReferenceIdentifier(stage))
}

func (codervsmtype *CodeRVSMType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codervsmtype.GongGetReferenceIdentifier(stage))
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codereferencepathidentifiertype.GongGetReferenceIdentifier(stage))
}

func (codereferenceroletype *CodeReferenceRoleType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codereferenceroletype.GongGetReferenceIdentifier(stage))
}

func (coderelativepositiontype *CodeRelativePositionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderelativepositiontype.GongGetReferenceIdentifier(stage))
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coderouteindicatortype.GongGetReferenceIdentifier(stage))
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

func (codesidetype *CodeSideType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codesidetype.GongGetReferenceIdentifier(stage))
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codespeedreferencetype.GongGetReferenceIdentifier(stage))
}

func (codetaatype *CodeTAAType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(codetaatype.GongGetReferenceIdentifier(stage))
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

func (textinstructiontype *TextInstructionType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textinstructiontype.GongGetReferenceIdentifier(stage))
}

func (textnametype *TextNameType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textnametype.GongGetReferenceIdentifier(stage))
}

func (textnotetype *TextNoteType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textnotetype.GongGetReferenceIdentifier(stage))
}

func (textpropertynametype *TextPropertyNameType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(textpropertynametype.GongGetReferenceIdentifier(stage))
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

func (valhextype *ValHexType) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(valhextype.GongGetReferenceIdentifier(stage))
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
