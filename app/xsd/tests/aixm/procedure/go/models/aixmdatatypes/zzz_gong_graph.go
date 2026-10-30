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

func (codeaircrafttype *CodeAircraftType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftTypes[codeaircrafttype]
	return ok
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftWingspanClassTypes[codeaircraftwingspanclasstype]
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

func (codearcdirectiontype *CodeArcDirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeArcDirectionTypes[codearcdirectiontype]
	return ok
}

func (codebearingtype *CodeBearingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeBearingTypes[codebearingtype]
	return ok
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCommunicationModeTypes[codecommunicationmodetype]
	return ok
}

func (codecoursetype *CodeCourseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCourseTypes[codecoursetype]
	return ok
}

func (codedaytype *CodeDayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDayTypes[codedaytype]
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

func (codedirectionreferencetype *CodeDirectionReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDirectionReferenceTypes[codedirectionreferencetype]
	return ok
}

func (codedirectionturntype *CodeDirectionTurnType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDirectionTurnTypes[codedirectionturntype]
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

func (codefinalguidancetype *CodeFinalGuidanceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFinalGuidanceTypes[codefinalguidancetype]
	return ok
}

func (codefreeflighttype *CodeFreeFlightType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeFreeFlightTypes[codefreeflighttype]
	return ok
}

func (codeheightreferencetype *CodeHeightReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeHeightReferenceTypes[codeheightreferencetype]
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

func (codeicaocountrytype *CodeICAOCountryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeICAOCountryTypes[codeicaocountrytype]
	return ok
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMilitaryRoutePointTypes[codemilitaryroutepointtype]
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

func (codenotepurposetype *CodeNotePurposeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNotePurposeTypes[codenotepurposetype]
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

func (codernptype *CodeRNPType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRNPTypes[codernptype]
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

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeReferencePathIdentifierTypes[codereferencepathidentifiertype]
	return ok
}

func (codereferenceroletype *CodeReferenceRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeReferenceRoleTypes[codereferenceroletype]
	return ok
}

func (coderelativepositiontype *CodeRelativePositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRelativePositionTypes[coderelativepositiontype]
	return ok
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRouteIndicatorTypes[coderouteindicatortype]
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

func (codesidetype *CodeSideType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSideTypes[codesidetype]
	return ok
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpeedReferenceTypes[codespeedreferencetype]
	return ok
}

func (codetaatype *CodeTAAType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTAATypes[codetaatype]
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

func (codeupperalphatype *CodeUpperAlphaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeUpperAlphaTypes[codeupperalphatype]
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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeWakeTurbulenceTypes[codewaketurbulencetype]
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

func (nonumbertype *NoNumberType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoNumberTypes[nonumbertype]
	return ok
}

func (nosequencetype *NoSequenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoSequenceTypes[nosequencetype]
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

func (textpropertynametype *TextPropertyNameType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TextPropertyNameTypes[textpropertynametype]
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

func (valhextype *ValHexType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValHexTypes[valhextype]
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

func (codearcdirectiontype *CodeArcDirectionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codearcdirectiontype) {
		return
	}

	codearcdirectiontype.Stage(stage)

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

func (codecommunicationmodetype *CodeCommunicationModeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecommunicationmodetype) {
		return
	}

	codecommunicationmodetype.Stage(stage)

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

func (codedaytype *CodeDayType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedaytype) {
		return
	}

	codedaytype.Stage(stage)

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

func (codefinalguidancetype *CodeFinalGuidanceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codefinalguidancetype) {
		return
	}

	codefinalguidancetype.Stage(stage)

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

func (codeheightreferencetype *CodeHeightReferenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeheightreferencetype) {
		return
	}

	codeheightreferencetype.Stage(stage)

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

func (codeicaocountrytype *CodeICAOCountryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeicaocountrytype) {
		return
	}

	codeicaocountrytype.Stage(stage)

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

func (codenotepurposetype *CodeNotePurposeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenotepurposetype) {
		return
	}

	codenotepurposetype.Stage(stage)

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

func (codernptype *CodeRNPType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codernptype) {
		return
	}

	codernptype.Stage(stage)

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

func (coderelativepositiontype *CodeRelativePositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderelativepositiontype) {
		return
	}

	coderelativepositiontype.Stage(stage)

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

func (codesidetype *CodeSideType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesidetype) {
		return
	}

	codesidetype.Stage(stage)

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

func (codetaatype *CodeTAAType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetaatype) {
		return
	}

	codetaatype.Stage(stage)

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

func (codeupperalphatype *CodeUpperAlphaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeupperalphatype) {
		return
	}

	codeupperalphatype.Stage(stage)

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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codewaketurbulencetype) {
		return
	}

	codewaketurbulencetype.Stage(stage)

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

func (textpropertynametype *TextPropertyNameType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(textpropertynametype) {
		return
	}

	textpropertynametype.Stage(stage)

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

func (valhextype *ValHexType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valhextype) {
		return
	}

	valhextype.Stage(stage)

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
	case *CodeATCReportingType:
		toT := GongCopyBranchCodeATCReportingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAerialRefuellingPointType:
		toT := GongCopyBranchCodeAerialRefuellingPointType(mapOrigCopy, fromT)
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

	case *CodeAircraftICAOType:
		toT := GongCopyBranchCodeAircraftICAOType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftType:
		toT := GongCopyBranchCodeAircraftType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftWingspanClassType:
		toT := GongCopyBranchCodeAircraftWingspanClassType(mapOrigCopy, fromT)
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

	case *CodeApproachPrefixType:
		toT := GongCopyBranchCodeApproachPrefixType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApproachType:
		toT := GongCopyBranchCodeApproachType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeApprovalType:
		toT := GongCopyBranchCodeApprovalType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeArcDirectionType:
		toT := GongCopyBranchCodeArcDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeBearingType:
		toT := GongCopyBranchCodeBearingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCommunicationModeType:
		toT := GongCopyBranchCodeCommunicationModeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCourseType:
		toT := GongCopyBranchCodeCourseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDayType:
		toT := GongCopyBranchCodeDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDesignStandardType:
		toT := GongCopyBranchCodeDesignStandardType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDesignatedPointDesignatorType:
		toT := GongCopyBranchCodeDesignatedPointDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDirectionReferenceType:
		toT := GongCopyBranchCodeDirectionReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDirectionTurnType:
		toT := GongCopyBranchCodeDirectionTurnType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeEquipmentAntiCollisionType:
		toT := GongCopyBranchCodeEquipmentAntiCollisionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeEquipmentUnavailableType:
		toT := GongCopyBranchCodeEquipmentUnavailableType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFinalGuidanceType:
		toT := GongCopyBranchCodeFinalGuidanceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeFreeFlightType:
		toT := GongCopyBranchCodeFreeFlightType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHeightReferenceType:
		toT := GongCopyBranchCodeHeightReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHoldingUsageType:
		toT := GongCopyBranchCodeHoldingUsageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeHoldingUseType:
		toT := GongCopyBranchCodeHoldingUseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeICAOCountryType:
		toT := GongCopyBranchCodeICAOCountryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMilitaryRoutePointType:
		toT := GongCopyBranchCodeMilitaryRoutePointType(mapOrigCopy, fromT)
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

	case *CodeNotePurposeType:
		toT := GongCopyBranchCodeNotePurposeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeObstacleAssessmentSurfaceType:
		toT := GongCopyBranchCodeObstacleAssessmentSurfaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeObstructionIdSurfaceZoneType:
		toT := GongCopyBranchCodeObstructionIdSurfaceZoneType(mapOrigCopy, fromT)
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

	case *CodeRNPType:
		toT := GongCopyBranchCodeRNPType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVSMPointRoleType:
		toT := GongCopyBranchCodeRVSMPointRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVSMType:
		toT := GongCopyBranchCodeRVSMType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeReferencePathIdentifierType:
		toT := GongCopyBranchCodeReferencePathIdentifierType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeReferenceRoleType:
		toT := GongCopyBranchCodeReferenceRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRelativePositionType:
		toT := GongCopyBranchCodeRelativePositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRouteIndicatorType:
		toT := GongCopyBranchCodeRouteIndicatorType(mapOrigCopy, fromT)
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

	case *CodeSideType:
		toT := GongCopyBranchCodeSideType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpeedReferenceType:
		toT := GongCopyBranchCodeSpeedReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTAAType:
		toT := GongCopyBranchCodeTAAType(mapOrigCopy, fromT)
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

	case *CodeUpperAlphaType:
		toT := GongCopyBranchCodeUpperAlphaType(mapOrigCopy, fromT)
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

	case *CodeWakeTurbulenceType:
		toT := GongCopyBranchCodeWakeTurbulenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeYesNoType:
		toT := GongCopyBranchCodeYesNoType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DateMonthDayType:
		toT := GongCopyBranchDateMonthDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoNumberType:
		toT := GongCopyBranchNoNumberType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoSequenceType:
		toT := GongCopyBranchNoSequenceType(mapOrigCopy, fromT)
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

	case *TextPropertyNameType:
		toT := GongCopyBranchTextPropertyNameType(mapOrigCopy, fromT)
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

	case *ValHexType:
		toT := GongCopyBranchValHexType(mapOrigCopy, fromT)
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

func (codearcdirectiontype *CodeArcDirectionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codearcdirectiontype) {
		return
	}

	codearcdirectiontype.Unstage(stage)

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

func (codecommunicationmodetype *CodeCommunicationModeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecommunicationmodetype) {
		return
	}

	codecommunicationmodetype.Unstage(stage)

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

func (codedaytype *CodeDayType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedaytype) {
		return
	}

	codedaytype.Unstage(stage)

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

func (codefinalguidancetype *CodeFinalGuidanceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codefinalguidancetype) {
		return
	}

	codefinalguidancetype.Unstage(stage)

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

func (codeheightreferencetype *CodeHeightReferenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeheightreferencetype) {
		return
	}

	codeheightreferencetype.Unstage(stage)

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

func (codeicaocountrytype *CodeICAOCountryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeicaocountrytype) {
		return
	}

	codeicaocountrytype.Unstage(stage)

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

func (codenotepurposetype *CodeNotePurposeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenotepurposetype) {
		return
	}

	codenotepurposetype.Unstage(stage)

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

func (codernptype *CodeRNPType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codernptype) {
		return
	}

	codernptype.Unstage(stage)

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

func (coderelativepositiontype *CodeRelativePositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderelativepositiontype) {
		return
	}

	coderelativepositiontype.Unstage(stage)

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

func (codesidetype *CodeSideType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesidetype) {
		return
	}

	codesidetype.Unstage(stage)

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

func (codetaatype *CodeTAAType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetaatype) {
		return
	}

	codetaatype.Unstage(stage)

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

func (codeupperalphatype *CodeUpperAlphaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeupperalphatype) {
		return
	}

	codeupperalphatype.Unstage(stage)

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

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codewaketurbulencetype) {
		return
	}

	codewaketurbulencetype.Unstage(stage)

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

func (textpropertynametype *TextPropertyNameType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(textpropertynametype) {
		return
	}

	textpropertynametype.Unstage(stage)

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

func (valhextype *ValHexType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valhextype) {
		return
	}

	valhextype.Unstage(stage)

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
func (reference *CodeATCReportingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeATCReportingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAerialRefuellingPointType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAerialRefuellingPointType) {
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

func (reference *CodeAircraftICAOType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftICAOType) {
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

func (reference *CodeArcDirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeArcDirectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeBearingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeBearingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCommunicationModeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCommunicationModeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCourseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCourseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDayType) {
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

func (reference *CodeDirectionReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDirectionReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDirectionTurnType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDirectionTurnType) {
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

func (reference *CodeFinalGuidanceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFinalGuidanceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeFreeFlightType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeFreeFlightType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeHeightReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeHeightReferenceType) {
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

func (reference *CodeICAOCountryType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeICAOCountryType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMilitaryRoutePointType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMilitaryRoutePointType) {
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

func (reference *CodeNotePurposeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNotePurposeType) {
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

func (reference *CodeRNPType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRNPType) {
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

func (reference *CodeReferencePathIdentifierType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeReferencePathIdentifierType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeReferenceRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeReferenceRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRelativePositionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRelativePositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRouteIndicatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRouteIndicatorType) {
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

func (reference *CodeSideType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSideType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpeedReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpeedReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTAAType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTAAType) {
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

func (reference *CodeUpperAlphaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeUpperAlphaType) {
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

func (reference *CodeWakeTurbulenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeWakeTurbulenceType) {
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

func (reference *NoNumberType) GongReconstructPointersFromReferences(stage *Stage, instance *NoNumberType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *NoSequenceType) GongReconstructPointersFromReferences(stage *Stage, instance *NoSequenceType) {
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

func (reference *TextPropertyNameType) GongReconstructPointersFromReferences(stage *Stage, instance *TextPropertyNameType) {
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

func (reference *ValHexType) GongReconstructPointersFromReferences(stage *Stage, instance *ValHexType) {
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
func (reference *CodeATCReportingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAerialRefuellingPointType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeAircraftICAOType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeArcDirectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeBearingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCommunicationModeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCourseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDayType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeDirectionReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDirectionTurnType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeFinalGuidanceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeFreeFlightType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeHeightReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeICAOCountryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMilitaryRoutePointType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeNotePurposeType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeRNPType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeReferencePathIdentifierType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeReferenceRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRelativePositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRouteIndicatorType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeSideType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpeedReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTAAType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeUpperAlphaType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeWakeTurbulenceType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *NoNumberType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *NoSequenceType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *TextPropertyNameType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *ValHexType) GongReconstructPointersFromInstances(stage *Stage) {
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
