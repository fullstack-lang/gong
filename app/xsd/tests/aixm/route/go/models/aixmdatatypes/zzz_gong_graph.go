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

func (codeauthoritytype *CodeAuthorityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAuthorityTypes[codeauthoritytype]
	return ok
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCardinalDirectionTypes[codecardinaldirectiontype]
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

func (codecoursetype *CodeCourseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCourseTypes[codecoursetype]
	return ok
}

func (codedaytype *CodeDayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDayTypes[codedaytype]
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

func (codeleveltype *CodeLevelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLevelTypes[codeleveltype]
	return ok
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeLocationQualifierTypes[codelocationqualifiertype]
	return ok
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMeteoConditionsTypes[codemeteoconditionstype]
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

func (codereferenceroletype *CodeReferenceRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeReferenceRoleTypes[codereferenceroletype]
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

func (codesidetype *CodeSideType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSideTypes[codesidetype]
	return ok
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpeedReferenceTypes[codespeedreferencetype]
	return ok
}

func (codetacanchanneltype *CodeTACANChannelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTACANChannelTypes[codetacanchanneltype]
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

func (valslopetype *ValSlopeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValSlopeTypes[valslopetype]
	return ok
}

func (valspeedtype *ValSpeedType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValSpeedTypes[valspeedtype]
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

func (codeauthoritytype *CodeAuthorityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeauthoritytype) {
		return
	}

	codeauthoritytype.Stage(stage)

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

func (codeleveltype *CodeLevelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeleveltype) {
		return
	}

	codeleveltype.Stage(stage)

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

func (codemeteoconditionstype *CodeMeteoConditionsType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemeteoconditionstype) {
		return
	}

	codemeteoconditionstype.Stage(stage)

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

func (codereferenceroletype *CodeReferenceRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codereferenceroletype) {
		return
	}

	codereferenceroletype.Stage(stage)

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

func (codetacanchanneltype *CodeTACANChannelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetacanchanneltype) {
		return
	}

	codetacanchanneltype.Stage(stage)

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

	case *CodeAuthorityType:
		toT := GongCopyBranchCodeAuthorityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCardinalDirectionType:
		toT := GongCopyBranchCodeCardinalDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCommunicationModeType:
		toT := GongCopyBranchCodeCommunicationModeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeComparisonType:
		toT := GongCopyBranchCodeComparisonType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCourseType:
		toT := GongCopyBranchCodeCourseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDayType:
		toT := GongCopyBranchCodeDayType(mapOrigCopy, fromT)
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

	case *CodeLevelType:
		toT := GongCopyBranchCodeLevelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLocationQualifierType:
		toT := GongCopyBranchCodeLocationQualifierType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMeteoConditionsType:
		toT := GongCopyBranchCodeMeteoConditionsType(mapOrigCopy, fromT)
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

	case *CodeRNPType:
		toT := GongCopyBranchCodeRNPType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVSMPointRoleType:
		toT := GongCopyBranchCodeRVSMPointRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRVSMType:
		toT := GongCopyBranchCodeRVSMType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeReferenceRoleType:
		toT := GongCopyBranchCodeReferenceRoleType(mapOrigCopy, fromT)
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

	case *CodeSideType:
		toT := GongCopyBranchCodeSideType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpeedReferenceType:
		toT := GongCopyBranchCodeSpeedReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeTACANChannelType:
		toT := GongCopyBranchCodeTACANChannelType(mapOrigCopy, fromT)
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

	case *ValSlopeType:
		toT := GongCopyBranchValSlopeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValSpeedType:
		toT := GongCopyBranchValSpeedType(mapOrigCopy, fromT)
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

func (codeauthoritytype *CodeAuthorityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeauthoritytype) {
		return
	}

	codeauthoritytype.Unstage(stage)

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

func (codeleveltype *CodeLevelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeleveltype) {
		return
	}

	codeleveltype.Unstage(stage)

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

func (codemeteoconditionstype *CodeMeteoConditionsType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemeteoconditionstype) {
		return
	}

	codemeteoconditionstype.Unstage(stage)

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

func (codereferenceroletype *CodeReferenceRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codereferenceroletype) {
		return
	}

	codereferenceroletype.Unstage(stage)

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

func (codetacanchanneltype *CodeTACANChannelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetacanchanneltype) {
		return
	}

	codetacanchanneltype.Unstage(stage)

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

func (reference *CodeAuthorityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAuthorityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCardinalDirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCardinalDirectionType) {
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

func (reference *CodeCourseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCourseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDayType) {
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

func (reference *CodeLevelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLevelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeLocationQualifierType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeLocationQualifierType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMeteoConditionsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMeteoConditionsType) {
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

func (reference *CodeReferenceRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeReferenceRoleType) {
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

func (reference *CodeSideType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSideType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpeedReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpeedReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeTACANChannelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeTACANChannelType) {
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

func (reference *ValSlopeType) GongReconstructPointersFromReferences(stage *Stage, instance *ValSlopeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValSpeedType) GongReconstructPointersFromReferences(stage *Stage, instance *ValSpeedType) {
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

func (reference *CodeAuthorityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCardinalDirectionType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeCourseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDayType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeLevelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeLocationQualifierType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMeteoConditionsType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeReferenceRoleType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeSideType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpeedReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeTACANChannelType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *ValSlopeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValSpeedType) GongReconstructPointersFromInstances(stage *Stage) {
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
