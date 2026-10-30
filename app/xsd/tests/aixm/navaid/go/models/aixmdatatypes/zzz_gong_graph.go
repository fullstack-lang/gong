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
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspacePointPositionTypes[codeairspacepointpositiontype]
	return ok
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspacePointRoleTypes[codeairspacepointroletype]
	return ok
}

func (codealtitudeusetype *CodeAltitudeUseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAltitudeUseTypes[codealtitudeusetype]
	return ok
}

func (codeauralmorsetype *CodeAuralMorseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAuralMorseTypes[codeauralmorsetype]
	return ok
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAuthorityRoleTypes[codeauthorityroletype]
	return ok
}

func (codebearingtype *CodeBearingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeBearingTypes[codebearingtype]
	return ok
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCardinalDirectionTypes[codecardinaldirectiontype]
	return ok
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCheckpointCategoryTypes[codecheckpointcategorytype]
	return ok
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCourseQualityILSTypes[codecoursequalityilstype]
	return ok
}

func (codedmechanneltype *CodeDMEChannelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDMEChannelTypes[codedmechanneltype]
	return ok
}

func (codedmetype *CodeDMEType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDMETypes[codedmetype]
	return ok
}

func (codedaytype *CodeDayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDayTypes[codedaytype]
	return ok
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDesignatedPointDesignatorTypes[codedesignatedpointdesignatortype]
	return ok
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDesignatedPointTypes[codedesignatedpointtype]
	return ok
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDirectionReferenceTypes[codedirectionreferencetype]
	return ok
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDistanceIndicationTypes[codedistanceindicationtype]
	return ok
}

func (codeemissionbandtype *CodeEmissionBandType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeEmissionBandTypes[codeemissionbandtype]
	return ok
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeILSBackCourseTypes[codeilsbackcoursetype]
	return ok
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeIntegrityLevelILSTypes[codeintegritylevelilstype]
	return ok
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMLSAzimuthTypes[codemlsazimuthtype]
	return ok
}

func (codemlschanneltype *CodeMLSChannelType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMLSChannelTypes[codemlschanneltype]
	return ok
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMarkerBeaconSignalTypes[codemarkerbeaconsignaltype]
	return ok
}

func (codendbusagetype *CodeNDBUsageType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNDBUsageTypes[codendbusagetype]
	return ok
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavaidDesignatorTypes[codenavaiddesignatortype]
	return ok
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavaidPurposeTypes[codenavaidpurposetype]
	return ok
}

func (codenavaidservicetype *CodeNavaidServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNavaidServiceTypes[codenavaidservicetype]
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

func (codepartype *CodePARType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePARTypes[codepartype]
	return ok
}

func (codepositioninilstype *CodePositionInILSType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePositionInILSTypes[codepositioninilstype]
	return ok
}

func (codeprimaryradartype *CodePrimaryRadarType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodePrimaryRadarTypes[codeprimaryradartype]
	return ok
}

func (coderadarservicetype *CodeRadarServiceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRadarServiceTypes[coderadarservicetype]
	return ok
}

func (coderadioemissiontype *CodeRadioEmissionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRadioEmissionTypes[coderadioemissiontype]
	return ok
}

func (coderadiosignaltype *CodeRadioSignalType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRadioSignalTypes[coderadiosignaltype]
	return ok
}

func (codereflectortype *CodeReflectorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeReflectorTypes[codereflectortype]
	return ok
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSignalPerformanceILSTypes[codesignalperformanceilstype]
	return ok
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpecialNavigationChainDesignatorTypes[codespecialnavigationchaindesignatortype]
	return ok
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpecialNavigationStationTypes[codespecialnavigationstationtype]
	return ok
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeSpecialNavigationSystemTypes[codespecialnavigationsystemtype]
	return ok
}

func (codestandbypowertype *CodeStandbyPowerType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStandbyPowerTypes[codestandbypowertype]
	return ok
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusNavaidTypes[codestatusnavaidtype]
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

func (codevortype *CodeVORType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVORTypes[codevortype]
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

func (valfrequencytype *ValFrequencyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValFrequencyTypes[valfrequencytype]
	return ok
}

func (valmagneticvariationtype *ValMagneticVariationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValMagneticVariationTypes[valmagneticvariationtype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacepointpositiontype) {
		return
	}

	codeairspacepointpositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacepointroletype) {
		return
	}

	codeairspacepointroletype.Stage(stage)

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

func (codeauralmorsetype *CodeAuralMorseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeauralmorsetype) {
		return
	}

	codeauralmorsetype.Stage(stage)

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

func (codebearingtype *CodeBearingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codebearingtype) {
		return
	}

	codebearingtype.Stage(stage)

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

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecheckpointcategorytype) {
		return
	}

	codecheckpointcategorytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecoursequalityilstype) {
		return
	}

	codecoursequalityilstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedmechanneltype *CodeDMEChannelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedmechanneltype) {
		return
	}

	codedmechanneltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedmetype *CodeDMEType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedmetype) {
		return
	}

	codedmetype.Stage(stage)

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

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedesignatedpointdesignatortype) {
		return
	}

	codedesignatedpointdesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedesignatedpointtype *CodeDesignatedPointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedesignatedpointtype) {
		return
	}

	codedesignatedpointtype.Stage(stage)

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

func (codedistanceindicationtype *CodeDistanceIndicationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codedistanceindicationtype) {
		return
	}

	codedistanceindicationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeemissionbandtype *CodeEmissionBandType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeemissionbandtype) {
		return
	}

	codeemissionbandtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeilsbackcoursetype) {
		return
	}

	codeilsbackcoursetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeintegritylevelilstype) {
		return
	}

	codeintegritylevelilstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemlsazimuthtype) {
		return
	}

	codemlsazimuthtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemlschanneltype *CodeMLSChannelType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemlschanneltype) {
		return
	}

	codemlschanneltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codemarkerbeaconsignaltype) {
		return
	}

	codemarkerbeaconsignaltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codendbusagetype *CodeNDBUsageType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codendbusagetype) {
		return
	}

	codendbusagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavaiddesignatortype) {
		return
	}

	codenavaiddesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavaidpurposetype) {
		return
	}

	codenavaidpurposetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaidservicetype *CodeNavaidServiceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenavaidservicetype) {
		return
	}

	codenavaidservicetype.Stage(stage)

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

func (codepartype *CodePARType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepartype) {
		return
	}

	codepartype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepositioninilstype *CodePositionInILSType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codepositioninilstype) {
		return
	}

	codepositioninilstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprimaryradartype *CodePrimaryRadarType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeprimaryradartype) {
		return
	}

	codeprimaryradartype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadarservicetype *CodeRadarServiceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderadarservicetype) {
		return
	}

	coderadarservicetype.Stage(stage)

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

func (coderadiosignaltype *CodeRadioSignalType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(coderadiosignaltype) {
		return
	}

	coderadiosignaltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codereflectortype *CodeReflectorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codereflectortype) {
		return
	}

	codereflectortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codesignalperformanceilstype) {
		return
	}

	codesignalperformanceilstype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespecialnavigationchaindesignatortype) {
		return
	}

	codespecialnavigationchaindesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespecialnavigationstationtype) {
		return
	}

	codespecialnavigationstationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codespecialnavigationsystemtype) {
		return
	}

	codespecialnavigationsystemtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestandbypowertype *CodeStandbyPowerType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestandbypowertype) {
		return
	}

	codestandbypowertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusnavaidtype *CodeStatusNavaidType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusnavaidtype) {
		return
	}

	codestatusnavaidtype.Stage(stage)

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

func (codevortype *CodeVORType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codevortype) {
		return
	}

	codevortype.Stage(stage)

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

func (valfrequencytype *ValFrequencyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(valfrequencytype) {
		return
	}

	valfrequencytype.Stage(stage)

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

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *CodeAirspacePointPositionType:
		toT := GongCopyBranchCodeAirspacePointPositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspacePointRoleType:
		toT := GongCopyBranchCodeAirspacePointRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAltitudeUseType:
		toT := GongCopyBranchCodeAltitudeUseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAuralMorseType:
		toT := GongCopyBranchCodeAuralMorseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAuthorityRoleType:
		toT := GongCopyBranchCodeAuthorityRoleType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeBearingType:
		toT := GongCopyBranchCodeBearingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCardinalDirectionType:
		toT := GongCopyBranchCodeCardinalDirectionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCheckpointCategoryType:
		toT := GongCopyBranchCodeCheckpointCategoryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCourseQualityILSType:
		toT := GongCopyBranchCodeCourseQualityILSType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDMEChannelType:
		toT := GongCopyBranchCodeDMEChannelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDMEType:
		toT := GongCopyBranchCodeDMEType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDayType:
		toT := GongCopyBranchCodeDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDesignatedPointDesignatorType:
		toT := GongCopyBranchCodeDesignatedPointDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDesignatedPointType:
		toT := GongCopyBranchCodeDesignatedPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDirectionReferenceType:
		toT := GongCopyBranchCodeDirectionReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDistanceIndicationType:
		toT := GongCopyBranchCodeDistanceIndicationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeEmissionBandType:
		toT := GongCopyBranchCodeEmissionBandType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeILSBackCourseType:
		toT := GongCopyBranchCodeILSBackCourseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeIntegrityLevelILSType:
		toT := GongCopyBranchCodeIntegrityLevelILSType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMLSAzimuthType:
		toT := GongCopyBranchCodeMLSAzimuthType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMLSChannelType:
		toT := GongCopyBranchCodeMLSChannelType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMarkerBeaconSignalType:
		toT := GongCopyBranchCodeMarkerBeaconSignalType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNDBUsageType:
		toT := GongCopyBranchCodeNDBUsageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavaidDesignatorType:
		toT := GongCopyBranchCodeNavaidDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavaidPurposeType:
		toT := GongCopyBranchCodeNavaidPurposeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNavaidServiceType:
		toT := GongCopyBranchCodeNavaidServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNorthReferenceType:
		toT := GongCopyBranchCodeNorthReferenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNotePurposeType:
		toT := GongCopyBranchCodeNotePurposeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePARType:
		toT := GongCopyBranchCodePARType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePositionInILSType:
		toT := GongCopyBranchCodePositionInILSType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodePrimaryRadarType:
		toT := GongCopyBranchCodePrimaryRadarType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRadarServiceType:
		toT := GongCopyBranchCodeRadarServiceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRadioEmissionType:
		toT := GongCopyBranchCodeRadioEmissionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeRadioSignalType:
		toT := GongCopyBranchCodeRadioSignalType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeReflectorType:
		toT := GongCopyBranchCodeReflectorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSignalPerformanceILSType:
		toT := GongCopyBranchCodeSignalPerformanceILSType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpecialNavigationChainDesignatorType:
		toT := GongCopyBranchCodeSpecialNavigationChainDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpecialNavigationStationType:
		toT := GongCopyBranchCodeSpecialNavigationStationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeSpecialNavigationSystemType:
		toT := GongCopyBranchCodeSpecialNavigationSystemType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStandbyPowerType:
		toT := GongCopyBranchCodeStandbyPowerType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusNavaidType:
		toT := GongCopyBranchCodeStatusNavaidType(mapOrigCopy, fromT)
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

	case *CodeVORType:
		toT := GongCopyBranchCodeVORType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalDatumType:
		toT := GongCopyBranchCodeVerticalDatumType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalReferenceType:
		toT := GongCopyBranchCodeVerticalReferenceType(mapOrigCopy, fromT)
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

	case *NoSequenceType:
		toT := GongCopyBranchNoSequenceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextAddressType:
		toT := GongCopyBranchTextAddressType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextDesignatorType:
		toT := GongCopyBranchTextDesignatorType(mapOrigCopy, fromT)
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

	case *ValFrequencyType:
		toT := GongCopyBranchValFrequencyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ValMagneticVariationType:
		toT := GongCopyBranchValMagneticVariationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCodeAirspacePointPositionType(mapOrigCopy map[any]any, codeairspacepointpositiontypeFrom *CodeAirspacePointPositionType) (codeairspacepointpositiontypeTo *CodeAirspacePointPositionType) {
	var alreadyCopied bool
	codeairspacepointpositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacepointpositiontypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacepointpositiontypeFrom.GongCopyBasicFields(codeairspacepointpositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspacePointRoleType(mapOrigCopy map[any]any, codeairspacepointroletypeFrom *CodeAirspacePointRoleType) (codeairspacepointroletypeTo *CodeAirspacePointRoleType) {
	var alreadyCopied bool
	codeairspacepointroletypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacepointroletypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacepointroletypeFrom.GongCopyBasicFields(codeairspacepointroletypeTo)

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

func GongCopyBranchCodeAuralMorseType(mapOrigCopy map[any]any, codeauralmorsetypeFrom *CodeAuralMorseType) (codeauralmorsetypeTo *CodeAuralMorseType) {
	var alreadyCopied bool
	codeauralmorsetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeauralmorsetypeFrom)
	if alreadyCopied {
		return
	}
	codeauralmorsetypeFrom.GongCopyBasicFields(codeauralmorsetypeTo)

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

func GongCopyBranchCodeCheckpointCategoryType(mapOrigCopy map[any]any, codecheckpointcategorytypeFrom *CodeCheckpointCategoryType) (codecheckpointcategorytypeTo *CodeCheckpointCategoryType) {
	var alreadyCopied bool
	codecheckpointcategorytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecheckpointcategorytypeFrom)
	if alreadyCopied {
		return
	}
	codecheckpointcategorytypeFrom.GongCopyBasicFields(codecheckpointcategorytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeCourseQualityILSType(mapOrigCopy map[any]any, codecoursequalityilstypeFrom *CodeCourseQualityILSType) (codecoursequalityilstypeTo *CodeCourseQualityILSType) {
	var alreadyCopied bool
	codecoursequalityilstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codecoursequalityilstypeFrom)
	if alreadyCopied {
		return
	}
	codecoursequalityilstypeFrom.GongCopyBasicFields(codecoursequalityilstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDMEChannelType(mapOrigCopy map[any]any, codedmechanneltypeFrom *CodeDMEChannelType) (codedmechanneltypeTo *CodeDMEChannelType) {
	var alreadyCopied bool
	codedmechanneltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedmechanneltypeFrom)
	if alreadyCopied {
		return
	}
	codedmechanneltypeFrom.GongCopyBasicFields(codedmechanneltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeDMEType(mapOrigCopy map[any]any, codedmetypeFrom *CodeDMEType) (codedmetypeTo *CodeDMEType) {
	var alreadyCopied bool
	codedmetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedmetypeFrom)
	if alreadyCopied {
		return
	}
	codedmetypeFrom.GongCopyBasicFields(codedmetypeTo)

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

func GongCopyBranchCodeDesignatedPointType(mapOrigCopy map[any]any, codedesignatedpointtypeFrom *CodeDesignatedPointType) (codedesignatedpointtypeTo *CodeDesignatedPointType) {
	var alreadyCopied bool
	codedesignatedpointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedesignatedpointtypeFrom)
	if alreadyCopied {
		return
	}
	codedesignatedpointtypeFrom.GongCopyBasicFields(codedesignatedpointtypeTo)

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

func GongCopyBranchCodeDistanceIndicationType(mapOrigCopy map[any]any, codedistanceindicationtypeFrom *CodeDistanceIndicationType) (codedistanceindicationtypeTo *CodeDistanceIndicationType) {
	var alreadyCopied bool
	codedistanceindicationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codedistanceindicationtypeFrom)
	if alreadyCopied {
		return
	}
	codedistanceindicationtypeFrom.GongCopyBasicFields(codedistanceindicationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeEmissionBandType(mapOrigCopy map[any]any, codeemissionbandtypeFrom *CodeEmissionBandType) (codeemissionbandtypeTo *CodeEmissionBandType) {
	var alreadyCopied bool
	codeemissionbandtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeemissionbandtypeFrom)
	if alreadyCopied {
		return
	}
	codeemissionbandtypeFrom.GongCopyBasicFields(codeemissionbandtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeILSBackCourseType(mapOrigCopy map[any]any, codeilsbackcoursetypeFrom *CodeILSBackCourseType) (codeilsbackcoursetypeTo *CodeILSBackCourseType) {
	var alreadyCopied bool
	codeilsbackcoursetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeilsbackcoursetypeFrom)
	if alreadyCopied {
		return
	}
	codeilsbackcoursetypeFrom.GongCopyBasicFields(codeilsbackcoursetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeIntegrityLevelILSType(mapOrigCopy map[any]any, codeintegritylevelilstypeFrom *CodeIntegrityLevelILSType) (codeintegritylevelilstypeTo *CodeIntegrityLevelILSType) {
	var alreadyCopied bool
	codeintegritylevelilstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeintegritylevelilstypeFrom)
	if alreadyCopied {
		return
	}
	codeintegritylevelilstypeFrom.GongCopyBasicFields(codeintegritylevelilstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMLSAzimuthType(mapOrigCopy map[any]any, codemlsazimuthtypeFrom *CodeMLSAzimuthType) (codemlsazimuthtypeTo *CodeMLSAzimuthType) {
	var alreadyCopied bool
	codemlsazimuthtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemlsazimuthtypeFrom)
	if alreadyCopied {
		return
	}
	codemlsazimuthtypeFrom.GongCopyBasicFields(codemlsazimuthtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMLSChannelType(mapOrigCopy map[any]any, codemlschanneltypeFrom *CodeMLSChannelType) (codemlschanneltypeTo *CodeMLSChannelType) {
	var alreadyCopied bool
	codemlschanneltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemlschanneltypeFrom)
	if alreadyCopied {
		return
	}
	codemlschanneltypeFrom.GongCopyBasicFields(codemlschanneltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeMarkerBeaconSignalType(mapOrigCopy map[any]any, codemarkerbeaconsignaltypeFrom *CodeMarkerBeaconSignalType) (codemarkerbeaconsignaltypeTo *CodeMarkerBeaconSignalType) {
	var alreadyCopied bool
	codemarkerbeaconsignaltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codemarkerbeaconsignaltypeFrom)
	if alreadyCopied {
		return
	}
	codemarkerbeaconsignaltypeFrom.GongCopyBasicFields(codemarkerbeaconsignaltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNDBUsageType(mapOrigCopy map[any]any, codendbusagetypeFrom *CodeNDBUsageType) (codendbusagetypeTo *CodeNDBUsageType) {
	var alreadyCopied bool
	codendbusagetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codendbusagetypeFrom)
	if alreadyCopied {
		return
	}
	codendbusagetypeFrom.GongCopyBasicFields(codendbusagetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavaidDesignatorType(mapOrigCopy map[any]any, codenavaiddesignatortypeFrom *CodeNavaidDesignatorType) (codenavaiddesignatortypeTo *CodeNavaidDesignatorType) {
	var alreadyCopied bool
	codenavaiddesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavaiddesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codenavaiddesignatortypeFrom.GongCopyBasicFields(codenavaiddesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavaidPurposeType(mapOrigCopy map[any]any, codenavaidpurposetypeFrom *CodeNavaidPurposeType) (codenavaidpurposetypeTo *CodeNavaidPurposeType) {
	var alreadyCopied bool
	codenavaidpurposetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavaidpurposetypeFrom)
	if alreadyCopied {
		return
	}
	codenavaidpurposetypeFrom.GongCopyBasicFields(codenavaidpurposetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeNavaidServiceType(mapOrigCopy map[any]any, codenavaidservicetypeFrom *CodeNavaidServiceType) (codenavaidservicetypeTo *CodeNavaidServiceType) {
	var alreadyCopied bool
	codenavaidservicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codenavaidservicetypeFrom)
	if alreadyCopied {
		return
	}
	codenavaidservicetypeFrom.GongCopyBasicFields(codenavaidservicetypeTo)

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

func GongCopyBranchCodePARType(mapOrigCopy map[any]any, codepartypeFrom *CodePARType) (codepartypeTo *CodePARType) {
	var alreadyCopied bool
	codepartypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepartypeFrom)
	if alreadyCopied {
		return
	}
	codepartypeFrom.GongCopyBasicFields(codepartypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePositionInILSType(mapOrigCopy map[any]any, codepositioninilstypeFrom *CodePositionInILSType) (codepositioninilstypeTo *CodePositionInILSType) {
	var alreadyCopied bool
	codepositioninilstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codepositioninilstypeFrom)
	if alreadyCopied {
		return
	}
	codepositioninilstypeFrom.GongCopyBasicFields(codepositioninilstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodePrimaryRadarType(mapOrigCopy map[any]any, codeprimaryradartypeFrom *CodePrimaryRadarType) (codeprimaryradartypeTo *CodePrimaryRadarType) {
	var alreadyCopied bool
	codeprimaryradartypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeprimaryradartypeFrom)
	if alreadyCopied {
		return
	}
	codeprimaryradartypeFrom.GongCopyBasicFields(codeprimaryradartypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeRadarServiceType(mapOrigCopy map[any]any, coderadarservicetypeFrom *CodeRadarServiceType) (coderadarservicetypeTo *CodeRadarServiceType) {
	var alreadyCopied bool
	coderadarservicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderadarservicetypeFrom)
	if alreadyCopied {
		return
	}
	coderadarservicetypeFrom.GongCopyBasicFields(coderadarservicetypeTo)

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

func GongCopyBranchCodeRadioSignalType(mapOrigCopy map[any]any, coderadiosignaltypeFrom *CodeRadioSignalType) (coderadiosignaltypeTo *CodeRadioSignalType) {
	var alreadyCopied bool
	coderadiosignaltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, coderadiosignaltypeFrom)
	if alreadyCopied {
		return
	}
	coderadiosignaltypeFrom.GongCopyBasicFields(coderadiosignaltypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeReflectorType(mapOrigCopy map[any]any, codereflectortypeFrom *CodeReflectorType) (codereflectortypeTo *CodeReflectorType) {
	var alreadyCopied bool
	codereflectortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codereflectortypeFrom)
	if alreadyCopied {
		return
	}
	codereflectortypeFrom.GongCopyBasicFields(codereflectortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSignalPerformanceILSType(mapOrigCopy map[any]any, codesignalperformanceilstypeFrom *CodeSignalPerformanceILSType) (codesignalperformanceilstypeTo *CodeSignalPerformanceILSType) {
	var alreadyCopied bool
	codesignalperformanceilstypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codesignalperformanceilstypeFrom)
	if alreadyCopied {
		return
	}
	codesignalperformanceilstypeFrom.GongCopyBasicFields(codesignalperformanceilstypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSpecialNavigationChainDesignatorType(mapOrigCopy map[any]any, codespecialnavigationchaindesignatortypeFrom *CodeSpecialNavigationChainDesignatorType) (codespecialnavigationchaindesignatortypeTo *CodeSpecialNavigationChainDesignatorType) {
	var alreadyCopied bool
	codespecialnavigationchaindesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codespecialnavigationchaindesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codespecialnavigationchaindesignatortypeFrom.GongCopyBasicFields(codespecialnavigationchaindesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSpecialNavigationStationType(mapOrigCopy map[any]any, codespecialnavigationstationtypeFrom *CodeSpecialNavigationStationType) (codespecialnavigationstationtypeTo *CodeSpecialNavigationStationType) {
	var alreadyCopied bool
	codespecialnavigationstationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codespecialnavigationstationtypeFrom)
	if alreadyCopied {
		return
	}
	codespecialnavigationstationtypeFrom.GongCopyBasicFields(codespecialnavigationstationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeSpecialNavigationSystemType(mapOrigCopy map[any]any, codespecialnavigationsystemtypeFrom *CodeSpecialNavigationSystemType) (codespecialnavigationsystemtypeTo *CodeSpecialNavigationSystemType) {
	var alreadyCopied bool
	codespecialnavigationsystemtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codespecialnavigationsystemtypeFrom)
	if alreadyCopied {
		return
	}
	codespecialnavigationsystemtypeFrom.GongCopyBasicFields(codespecialnavigationsystemtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStandbyPowerType(mapOrigCopy map[any]any, codestandbypowertypeFrom *CodeStandbyPowerType) (codestandbypowertypeTo *CodeStandbyPowerType) {
	var alreadyCopied bool
	codestandbypowertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestandbypowertypeFrom)
	if alreadyCopied {
		return
	}
	codestandbypowertypeFrom.GongCopyBasicFields(codestandbypowertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeStatusNavaidType(mapOrigCopy map[any]any, codestatusnavaidtypeFrom *CodeStatusNavaidType) (codestatusnavaidtypeTo *CodeStatusNavaidType) {
	var alreadyCopied bool
	codestatusnavaidtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusnavaidtypeFrom)
	if alreadyCopied {
		return
	}
	codestatusnavaidtypeFrom.GongCopyBasicFields(codestatusnavaidtypeTo)

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

func GongCopyBranchCodeVORType(mapOrigCopy map[any]any, codevortypeFrom *CodeVORType) (codevortypeTo *CodeVORType) {
	var alreadyCopied bool
	codevortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codevortypeFrom)
	if alreadyCopied {
		return
	}
	codevortypeFrom.GongCopyBasicFields(codevortypeTo)

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
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacepointpositiontype) {
		return
	}

	codeairspacepointpositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacepointroletype) {
		return
	}

	codeairspacepointroletype.Unstage(stage)

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

func (codeauralmorsetype *CodeAuralMorseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeauralmorsetype) {
		return
	}

	codeauralmorsetype.Unstage(stage)

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

func (codebearingtype *CodeBearingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codebearingtype) {
		return
	}

	codebearingtype.Unstage(stage)

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

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecheckpointcategorytype) {
		return
	}

	codecheckpointcategorytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecoursequalityilstype) {
		return
	}

	codecoursequalityilstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedmechanneltype *CodeDMEChannelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedmechanneltype) {
		return
	}

	codedmechanneltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedmetype *CodeDMEType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedmetype) {
		return
	}

	codedmetype.Unstage(stage)

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

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedesignatedpointdesignatortype) {
		return
	}

	codedesignatedpointdesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codedesignatedpointtype *CodeDesignatedPointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedesignatedpointtype) {
		return
	}

	codedesignatedpointtype.Unstage(stage)

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

func (codedistanceindicationtype *CodeDistanceIndicationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codedistanceindicationtype) {
		return
	}

	codedistanceindicationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeemissionbandtype *CodeEmissionBandType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeemissionbandtype) {
		return
	}

	codeemissionbandtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeilsbackcoursetype) {
		return
	}

	codeilsbackcoursetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeintegritylevelilstype) {
		return
	}

	codeintegritylevelilstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemlsazimuthtype) {
		return
	}

	codemlsazimuthtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemlschanneltype *CodeMLSChannelType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemlschanneltype) {
		return
	}

	codemlschanneltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codemarkerbeaconsignaltype) {
		return
	}

	codemarkerbeaconsignaltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codendbusagetype *CodeNDBUsageType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codendbusagetype) {
		return
	}

	codendbusagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavaiddesignatortype) {
		return
	}

	codenavaiddesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavaidpurposetype) {
		return
	}

	codenavaidpurposetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codenavaidservicetype *CodeNavaidServiceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenavaidservicetype) {
		return
	}

	codenavaidservicetype.Unstage(stage)

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

func (codepartype *CodePARType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepartype) {
		return
	}

	codepartype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codepositioninilstype *CodePositionInILSType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codepositioninilstype) {
		return
	}

	codepositioninilstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeprimaryradartype *CodePrimaryRadarType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeprimaryradartype) {
		return
	}

	codeprimaryradartype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (coderadarservicetype *CodeRadarServiceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderadarservicetype) {
		return
	}

	coderadarservicetype.Unstage(stage)

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

func (coderadiosignaltype *CodeRadioSignalType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(coderadiosignaltype) {
		return
	}

	coderadiosignaltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codereflectortype *CodeReflectorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codereflectortype) {
		return
	}

	codereflectortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codesignalperformanceilstype) {
		return
	}

	codesignalperformanceilstype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespecialnavigationchaindesignatortype) {
		return
	}

	codespecialnavigationchaindesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespecialnavigationstationtype) {
		return
	}

	codespecialnavigationstationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codespecialnavigationsystemtype) {
		return
	}

	codespecialnavigationsystemtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestandbypowertype *CodeStandbyPowerType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestandbypowertype) {
		return
	}

	codestandbypowertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusnavaidtype *CodeStatusNavaidType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusnavaidtype) {
		return
	}

	codestatusnavaidtype.Unstage(stage)

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

func (codevortype *CodeVORType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codevortype) {
		return
	}

	codevortype.Unstage(stage)

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

func (valfrequencytype *ValFrequencyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(valfrequencytype) {
		return
	}

	valfrequencytype.Unstage(stage)

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

// insertion point for pointer reconstruction from references
func (reference *CodeAirspacePointPositionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspacePointPositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspacePointRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspacePointRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAltitudeUseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAltitudeUseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAuralMorseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAuralMorseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAuthorityRoleType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAuthorityRoleType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeBearingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeBearingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCardinalDirectionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCardinalDirectionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCheckpointCategoryType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCheckpointCategoryType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeCourseQualityILSType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCourseQualityILSType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDMEChannelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDMEChannelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDMEType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDMEType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDayType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDesignatedPointDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDesignatedPointDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDesignatedPointType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDesignatedPointType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDirectionReferenceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDirectionReferenceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDistanceIndicationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDistanceIndicationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeEmissionBandType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeEmissionBandType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeILSBackCourseType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeILSBackCourseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeIntegrityLevelILSType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeIntegrityLevelILSType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMLSAzimuthType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMLSAzimuthType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMLSChannelType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMLSChannelType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMarkerBeaconSignalType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMarkerBeaconSignalType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNDBUsageType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNDBUsageType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavaidDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavaidDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavaidPurposeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavaidPurposeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeNavaidServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNavaidServiceType) {
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

func (reference *CodePARType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePARType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePositionInILSType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePositionInILSType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodePrimaryRadarType) GongReconstructPointersFromReferences(stage *Stage, instance *CodePrimaryRadarType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRadarServiceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRadarServiceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRadioEmissionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRadioEmissionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeRadioSignalType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRadioSignalType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeReflectorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeReflectorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSignalPerformanceILSType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSignalPerformanceILSType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpecialNavigationChainDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpecialNavigationChainDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpecialNavigationStationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpecialNavigationStationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeSpecialNavigationSystemType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeSpecialNavigationSystemType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStandbyPowerType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStandbyPowerType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusNavaidType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusNavaidType) {
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

func (reference *CodeVORType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVORType) {
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

func (reference *ValFrequencyType) GongReconstructPointersFromReferences(stage *Stage, instance *ValFrequencyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ValMagneticVariationType) GongReconstructPointersFromReferences(stage *Stage, instance *ValMagneticVariationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *CodeAirspacePointPositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspacePointRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAltitudeUseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAuralMorseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAuthorityRoleType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeBearingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCardinalDirectionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCheckpointCategoryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeCourseQualityILSType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDMEChannelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDMEType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDayType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDesignatedPointDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDesignatedPointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDirectionReferenceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDistanceIndicationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeEmissionBandType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeILSBackCourseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeIntegrityLevelILSType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMLSAzimuthType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMLSChannelType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMarkerBeaconSignalType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNDBUsageType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavaidDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavaidPurposeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeNavaidServiceType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodePARType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePositionInILSType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodePrimaryRadarType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRadarServiceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRadioEmissionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeRadioSignalType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeReflectorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSignalPerformanceILSType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpecialNavigationChainDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpecialNavigationStationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeSpecialNavigationSystemType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStandbyPowerType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusNavaidType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeVORType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *ValFrequencyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ValMagneticVariationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongDiff(stage *Stage, codeairspacepointpositiontypeOther *CodeAirspacePointPositionType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacepointpositiontype.Name != codeairspacepointpositiontypeOther.Name {
		diffs = append(diffs, codeairspacepointpositiontype.GongMarshallField(stage, "Name"))
	}
	if codeairspacepointpositiontype.NilReason != codeairspacepointpositiontypeOther.NilReason {
		diffs = append(diffs, codeairspacepointpositiontype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacepointpositiontype.EnclosedText != codeairspacepointpositiontypeOther.EnclosedText {
		diffs = append(diffs, codeairspacepointpositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacepointroletype *CodeAirspacePointRoleType) GongDiff(stage *Stage, codeairspacepointroletypeOther *CodeAirspacePointRoleType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacepointroletype.Name != codeairspacepointroletypeOther.Name {
		diffs = append(diffs, codeairspacepointroletype.GongMarshallField(stage, "Name"))
	}
	if codeairspacepointroletype.NilReason != codeairspacepointroletypeOther.NilReason {
		diffs = append(diffs, codeairspacepointroletype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacepointroletype.EnclosedText != codeairspacepointroletypeOther.EnclosedText {
		diffs = append(diffs, codeairspacepointroletype.GongMarshallField(stage, "EnclosedText"))
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
func (codeauralmorsetype *CodeAuralMorseType) GongDiff(stage *Stage, codeauralmorsetypeOther *CodeAuralMorseType) (diffs []string) {
	// insertion point for field diffs
	if codeauralmorsetype.Name != codeauralmorsetypeOther.Name {
		diffs = append(diffs, codeauralmorsetype.GongMarshallField(stage, "Name"))
	}
	if codeauralmorsetype.NilReason != codeauralmorsetypeOther.NilReason {
		diffs = append(diffs, codeauralmorsetype.GongMarshallField(stage, "NilReason"))
	}
	if codeauralmorsetype.EnclosedText != codeauralmorsetypeOther.EnclosedText {
		diffs = append(diffs, codeauralmorsetype.GongMarshallField(stage, "EnclosedText"))
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
func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongDiff(stage *Stage, codecheckpointcategorytypeOther *CodeCheckpointCategoryType) (diffs []string) {
	// insertion point for field diffs
	if codecheckpointcategorytype.Name != codecheckpointcategorytypeOther.Name {
		diffs = append(diffs, codecheckpointcategorytype.GongMarshallField(stage, "Name"))
	}
	if codecheckpointcategorytype.NilReason != codecheckpointcategorytypeOther.NilReason {
		diffs = append(diffs, codecheckpointcategorytype.GongMarshallField(stage, "NilReason"))
	}
	if codecheckpointcategorytype.EnclosedText != codecheckpointcategorytypeOther.EnclosedText {
		diffs = append(diffs, codecheckpointcategorytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codecoursequalityilstype *CodeCourseQualityILSType) GongDiff(stage *Stage, codecoursequalityilstypeOther *CodeCourseQualityILSType) (diffs []string) {
	// insertion point for field diffs
	if codecoursequalityilstype.Name != codecoursequalityilstypeOther.Name {
		diffs = append(diffs, codecoursequalityilstype.GongMarshallField(stage, "Name"))
	}
	if codecoursequalityilstype.NilReason != codecoursequalityilstypeOther.NilReason {
		diffs = append(diffs, codecoursequalityilstype.GongMarshallField(stage, "NilReason"))
	}
	if codecoursequalityilstype.EnclosedText != codecoursequalityilstypeOther.EnclosedText {
		diffs = append(diffs, codecoursequalityilstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedmechanneltype *CodeDMEChannelType) GongDiff(stage *Stage, codedmechanneltypeOther *CodeDMEChannelType) (diffs []string) {
	// insertion point for field diffs
	if codedmechanneltype.Name != codedmechanneltypeOther.Name {
		diffs = append(diffs, codedmechanneltype.GongMarshallField(stage, "Name"))
	}
	if codedmechanneltype.NilReason != codedmechanneltypeOther.NilReason {
		diffs = append(diffs, codedmechanneltype.GongMarshallField(stage, "NilReason"))
	}
	if codedmechanneltype.EnclosedText != codedmechanneltypeOther.EnclosedText {
		diffs = append(diffs, codedmechanneltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codedmetype *CodeDMEType) GongDiff(stage *Stage, codedmetypeOther *CodeDMEType) (diffs []string) {
	// insertion point for field diffs
	if codedmetype.Name != codedmetypeOther.Name {
		diffs = append(diffs, codedmetype.GongMarshallField(stage, "Name"))
	}
	if codedmetype.NilReason != codedmetypeOther.NilReason {
		diffs = append(diffs, codedmetype.GongMarshallField(stage, "NilReason"))
	}
	if codedmetype.EnclosedText != codedmetypeOther.EnclosedText {
		diffs = append(diffs, codedmetype.GongMarshallField(stage, "EnclosedText"))
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
func (codedesignatedpointtype *CodeDesignatedPointType) GongDiff(stage *Stage, codedesignatedpointtypeOther *CodeDesignatedPointType) (diffs []string) {
	// insertion point for field diffs
	if codedesignatedpointtype.Name != codedesignatedpointtypeOther.Name {
		diffs = append(diffs, codedesignatedpointtype.GongMarshallField(stage, "Name"))
	}
	if codedesignatedpointtype.NilReason != codedesignatedpointtypeOther.NilReason {
		diffs = append(diffs, codedesignatedpointtype.GongMarshallField(stage, "NilReason"))
	}
	if codedesignatedpointtype.EnclosedText != codedesignatedpointtypeOther.EnclosedText {
		diffs = append(diffs, codedesignatedpointtype.GongMarshallField(stage, "EnclosedText"))
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
func (codedistanceindicationtype *CodeDistanceIndicationType) GongDiff(stage *Stage, codedistanceindicationtypeOther *CodeDistanceIndicationType) (diffs []string) {
	// insertion point for field diffs
	if codedistanceindicationtype.Name != codedistanceindicationtypeOther.Name {
		diffs = append(diffs, codedistanceindicationtype.GongMarshallField(stage, "Name"))
	}
	if codedistanceindicationtype.NilReason != codedistanceindicationtypeOther.NilReason {
		diffs = append(diffs, codedistanceindicationtype.GongMarshallField(stage, "NilReason"))
	}
	if codedistanceindicationtype.EnclosedText != codedistanceindicationtypeOther.EnclosedText {
		diffs = append(diffs, codedistanceindicationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeemissionbandtype *CodeEmissionBandType) GongDiff(stage *Stage, codeemissionbandtypeOther *CodeEmissionBandType) (diffs []string) {
	// insertion point for field diffs
	if codeemissionbandtype.Name != codeemissionbandtypeOther.Name {
		diffs = append(diffs, codeemissionbandtype.GongMarshallField(stage, "Name"))
	}
	if codeemissionbandtype.NilReason != codeemissionbandtypeOther.NilReason {
		diffs = append(diffs, codeemissionbandtype.GongMarshallField(stage, "NilReason"))
	}
	if codeemissionbandtype.EnclosedText != codeemissionbandtypeOther.EnclosedText {
		diffs = append(diffs, codeemissionbandtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeilsbackcoursetype *CodeILSBackCourseType) GongDiff(stage *Stage, codeilsbackcoursetypeOther *CodeILSBackCourseType) (diffs []string) {
	// insertion point for field diffs
	if codeilsbackcoursetype.Name != codeilsbackcoursetypeOther.Name {
		diffs = append(diffs, codeilsbackcoursetype.GongMarshallField(stage, "Name"))
	}
	if codeilsbackcoursetype.NilReason != codeilsbackcoursetypeOther.NilReason {
		diffs = append(diffs, codeilsbackcoursetype.GongMarshallField(stage, "NilReason"))
	}
	if codeilsbackcoursetype.EnclosedText != codeilsbackcoursetypeOther.EnclosedText {
		diffs = append(diffs, codeilsbackcoursetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongDiff(stage *Stage, codeintegritylevelilstypeOther *CodeIntegrityLevelILSType) (diffs []string) {
	// insertion point for field diffs
	if codeintegritylevelilstype.Name != codeintegritylevelilstypeOther.Name {
		diffs = append(diffs, codeintegritylevelilstype.GongMarshallField(stage, "Name"))
	}
	if codeintegritylevelilstype.NilReason != codeintegritylevelilstypeOther.NilReason {
		diffs = append(diffs, codeintegritylevelilstype.GongMarshallField(stage, "NilReason"))
	}
	if codeintegritylevelilstype.EnclosedText != codeintegritylevelilstypeOther.EnclosedText {
		diffs = append(diffs, codeintegritylevelilstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemlsazimuthtype *CodeMLSAzimuthType) GongDiff(stage *Stage, codemlsazimuthtypeOther *CodeMLSAzimuthType) (diffs []string) {
	// insertion point for field diffs
	if codemlsazimuthtype.Name != codemlsazimuthtypeOther.Name {
		diffs = append(diffs, codemlsazimuthtype.GongMarshallField(stage, "Name"))
	}
	if codemlsazimuthtype.NilReason != codemlsazimuthtypeOther.NilReason {
		diffs = append(diffs, codemlsazimuthtype.GongMarshallField(stage, "NilReason"))
	}
	if codemlsazimuthtype.EnclosedText != codemlsazimuthtypeOther.EnclosedText {
		diffs = append(diffs, codemlsazimuthtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemlschanneltype *CodeMLSChannelType) GongDiff(stage *Stage, codemlschanneltypeOther *CodeMLSChannelType) (diffs []string) {
	// insertion point for field diffs
	if codemlschanneltype.Name != codemlschanneltypeOther.Name {
		diffs = append(diffs, codemlschanneltype.GongMarshallField(stage, "Name"))
	}
	if codemlschanneltype.NilReason != codemlschanneltypeOther.NilReason {
		diffs = append(diffs, codemlschanneltype.GongMarshallField(stage, "NilReason"))
	}
	if codemlschanneltype.EnclosedText != codemlschanneltypeOther.EnclosedText {
		diffs = append(diffs, codemlschanneltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongDiff(stage *Stage, codemarkerbeaconsignaltypeOther *CodeMarkerBeaconSignalType) (diffs []string) {
	// insertion point for field diffs
	if codemarkerbeaconsignaltype.Name != codemarkerbeaconsignaltypeOther.Name {
		diffs = append(diffs, codemarkerbeaconsignaltype.GongMarshallField(stage, "Name"))
	}
	if codemarkerbeaconsignaltype.NilReason != codemarkerbeaconsignaltypeOther.NilReason {
		diffs = append(diffs, codemarkerbeaconsignaltype.GongMarshallField(stage, "NilReason"))
	}
	if codemarkerbeaconsignaltype.EnclosedText != codemarkerbeaconsignaltypeOther.EnclosedText {
		diffs = append(diffs, codemarkerbeaconsignaltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codendbusagetype *CodeNDBUsageType) GongDiff(stage *Stage, codendbusagetypeOther *CodeNDBUsageType) (diffs []string) {
	// insertion point for field diffs
	if codendbusagetype.Name != codendbusagetypeOther.Name {
		diffs = append(diffs, codendbusagetype.GongMarshallField(stage, "Name"))
	}
	if codendbusagetype.NilReason != codendbusagetypeOther.NilReason {
		diffs = append(diffs, codendbusagetype.GongMarshallField(stage, "NilReason"))
	}
	if codendbusagetype.EnclosedText != codendbusagetypeOther.EnclosedText {
		diffs = append(diffs, codendbusagetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongDiff(stage *Stage, codenavaiddesignatortypeOther *CodeNavaidDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codenavaiddesignatortype.Name != codenavaiddesignatortypeOther.Name {
		diffs = append(diffs, codenavaiddesignatortype.GongMarshallField(stage, "Name"))
	}
	if codenavaiddesignatortype.NilReason != codenavaiddesignatortypeOther.NilReason {
		diffs = append(diffs, codenavaiddesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codenavaiddesignatortype.EnclosedText != codenavaiddesignatortypeOther.EnclosedText {
		diffs = append(diffs, codenavaiddesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavaidpurposetype *CodeNavaidPurposeType) GongDiff(stage *Stage, codenavaidpurposetypeOther *CodeNavaidPurposeType) (diffs []string) {
	// insertion point for field diffs
	if codenavaidpurposetype.Name != codenavaidpurposetypeOther.Name {
		diffs = append(diffs, codenavaidpurposetype.GongMarshallField(stage, "Name"))
	}
	if codenavaidpurposetype.NilReason != codenavaidpurposetypeOther.NilReason {
		diffs = append(diffs, codenavaidpurposetype.GongMarshallField(stage, "NilReason"))
	}
	if codenavaidpurposetype.EnclosedText != codenavaidpurposetypeOther.EnclosedText {
		diffs = append(diffs, codenavaidpurposetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codenavaidservicetype *CodeNavaidServiceType) GongDiff(stage *Stage, codenavaidservicetypeOther *CodeNavaidServiceType) (diffs []string) {
	// insertion point for field diffs
	if codenavaidservicetype.Name != codenavaidservicetypeOther.Name {
		diffs = append(diffs, codenavaidservicetype.GongMarshallField(stage, "Name"))
	}
	if codenavaidservicetype.NilReason != codenavaidservicetypeOther.NilReason {
		diffs = append(diffs, codenavaidservicetype.GongMarshallField(stage, "NilReason"))
	}
	if codenavaidservicetype.EnclosedText != codenavaidservicetypeOther.EnclosedText {
		diffs = append(diffs, codenavaidservicetype.GongMarshallField(stage, "EnclosedText"))
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
func (codepartype *CodePARType) GongDiff(stage *Stage, codepartypeOther *CodePARType) (diffs []string) {
	// insertion point for field diffs
	if codepartype.Name != codepartypeOther.Name {
		diffs = append(diffs, codepartype.GongMarshallField(stage, "Name"))
	}
	if codepartype.NilReason != codepartypeOther.NilReason {
		diffs = append(diffs, codepartype.GongMarshallField(stage, "NilReason"))
	}
	if codepartype.EnclosedText != codepartypeOther.EnclosedText {
		diffs = append(diffs, codepartype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codepositioninilstype *CodePositionInILSType) GongDiff(stage *Stage, codepositioninilstypeOther *CodePositionInILSType) (diffs []string) {
	// insertion point for field diffs
	if codepositioninilstype.Name != codepositioninilstypeOther.Name {
		diffs = append(diffs, codepositioninilstype.GongMarshallField(stage, "Name"))
	}
	if codepositioninilstype.NilReason != codepositioninilstypeOther.NilReason {
		diffs = append(diffs, codepositioninilstype.GongMarshallField(stage, "NilReason"))
	}
	if codepositioninilstype.EnclosedText != codepositioninilstypeOther.EnclosedText {
		diffs = append(diffs, codepositioninilstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeprimaryradartype *CodePrimaryRadarType) GongDiff(stage *Stage, codeprimaryradartypeOther *CodePrimaryRadarType) (diffs []string) {
	// insertion point for field diffs
	if codeprimaryradartype.Name != codeprimaryradartypeOther.Name {
		diffs = append(diffs, codeprimaryradartype.GongMarshallField(stage, "Name"))
	}
	if codeprimaryradartype.NilReason != codeprimaryradartypeOther.NilReason {
		diffs = append(diffs, codeprimaryradartype.GongMarshallField(stage, "NilReason"))
	}
	if codeprimaryradartype.EnclosedText != codeprimaryradartypeOther.EnclosedText {
		diffs = append(diffs, codeprimaryradartype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (coderadarservicetype *CodeRadarServiceType) GongDiff(stage *Stage, coderadarservicetypeOther *CodeRadarServiceType) (diffs []string) {
	// insertion point for field diffs
	if coderadarservicetype.Name != coderadarservicetypeOther.Name {
		diffs = append(diffs, coderadarservicetype.GongMarshallField(stage, "Name"))
	}
	if coderadarservicetype.NilReason != coderadarservicetypeOther.NilReason {
		diffs = append(diffs, coderadarservicetype.GongMarshallField(stage, "NilReason"))
	}
	if coderadarservicetype.EnclosedText != coderadarservicetypeOther.EnclosedText {
		diffs = append(diffs, coderadarservicetype.GongMarshallField(stage, "EnclosedText"))
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
func (coderadiosignaltype *CodeRadioSignalType) GongDiff(stage *Stage, coderadiosignaltypeOther *CodeRadioSignalType) (diffs []string) {
	// insertion point for field diffs
	if coderadiosignaltype.Name != coderadiosignaltypeOther.Name {
		diffs = append(diffs, coderadiosignaltype.GongMarshallField(stage, "Name"))
	}
	if coderadiosignaltype.NilReason != coderadiosignaltypeOther.NilReason {
		diffs = append(diffs, coderadiosignaltype.GongMarshallField(stage, "NilReason"))
	}
	if coderadiosignaltype.EnclosedText != coderadiosignaltypeOther.EnclosedText {
		diffs = append(diffs, coderadiosignaltype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codereflectortype *CodeReflectorType) GongDiff(stage *Stage, codereflectortypeOther *CodeReflectorType) (diffs []string) {
	// insertion point for field diffs
	if codereflectortype.Name != codereflectortypeOther.Name {
		diffs = append(diffs, codereflectortype.GongMarshallField(stage, "Name"))
	}
	if codereflectortype.NilReason != codereflectortypeOther.NilReason {
		diffs = append(diffs, codereflectortype.GongMarshallField(stage, "NilReason"))
	}
	if codereflectortype.EnclosedText != codereflectortypeOther.EnclosedText {
		diffs = append(diffs, codereflectortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongDiff(stage *Stage, codesignalperformanceilstypeOther *CodeSignalPerformanceILSType) (diffs []string) {
	// insertion point for field diffs
	if codesignalperformanceilstype.Name != codesignalperformanceilstypeOther.Name {
		diffs = append(diffs, codesignalperformanceilstype.GongMarshallField(stage, "Name"))
	}
	if codesignalperformanceilstype.NilReason != codesignalperformanceilstypeOther.NilReason {
		diffs = append(diffs, codesignalperformanceilstype.GongMarshallField(stage, "NilReason"))
	}
	if codesignalperformanceilstype.EnclosedText != codesignalperformanceilstypeOther.EnclosedText {
		diffs = append(diffs, codesignalperformanceilstype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongDiff(stage *Stage, codespecialnavigationchaindesignatortypeOther *CodeSpecialNavigationChainDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codespecialnavigationchaindesignatortype.Name != codespecialnavigationchaindesignatortypeOther.Name {
		diffs = append(diffs, codespecialnavigationchaindesignatortype.GongMarshallField(stage, "Name"))
	}
	if codespecialnavigationchaindesignatortype.NilReason != codespecialnavigationchaindesignatortypeOther.NilReason {
		diffs = append(diffs, codespecialnavigationchaindesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codespecialnavigationchaindesignatortype.EnclosedText != codespecialnavigationchaindesignatortypeOther.EnclosedText {
		diffs = append(diffs, codespecialnavigationchaindesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongDiff(stage *Stage, codespecialnavigationstationtypeOther *CodeSpecialNavigationStationType) (diffs []string) {
	// insertion point for field diffs
	if codespecialnavigationstationtype.Name != codespecialnavigationstationtypeOther.Name {
		diffs = append(diffs, codespecialnavigationstationtype.GongMarshallField(stage, "Name"))
	}
	if codespecialnavigationstationtype.NilReason != codespecialnavigationstationtypeOther.NilReason {
		diffs = append(diffs, codespecialnavigationstationtype.GongMarshallField(stage, "NilReason"))
	}
	if codespecialnavigationstationtype.EnclosedText != codespecialnavigationstationtypeOther.EnclosedText {
		diffs = append(diffs, codespecialnavigationstationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongDiff(stage *Stage, codespecialnavigationsystemtypeOther *CodeSpecialNavigationSystemType) (diffs []string) {
	// insertion point for field diffs
	if codespecialnavigationsystemtype.Name != codespecialnavigationsystemtypeOther.Name {
		diffs = append(diffs, codespecialnavigationsystemtype.GongMarshallField(stage, "Name"))
	}
	if codespecialnavigationsystemtype.NilReason != codespecialnavigationsystemtypeOther.NilReason {
		diffs = append(diffs, codespecialnavigationsystemtype.GongMarshallField(stage, "NilReason"))
	}
	if codespecialnavigationsystemtype.EnclosedText != codespecialnavigationsystemtypeOther.EnclosedText {
		diffs = append(diffs, codespecialnavigationsystemtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestandbypowertype *CodeStandbyPowerType) GongDiff(stage *Stage, codestandbypowertypeOther *CodeStandbyPowerType) (diffs []string) {
	// insertion point for field diffs
	if codestandbypowertype.Name != codestandbypowertypeOther.Name {
		diffs = append(diffs, codestandbypowertype.GongMarshallField(stage, "Name"))
	}
	if codestandbypowertype.NilReason != codestandbypowertypeOther.NilReason {
		diffs = append(diffs, codestandbypowertype.GongMarshallField(stage, "NilReason"))
	}
	if codestandbypowertype.EnclosedText != codestandbypowertypeOther.EnclosedText {
		diffs = append(diffs, codestandbypowertype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codestatusnavaidtype *CodeStatusNavaidType) GongDiff(stage *Stage, codestatusnavaidtypeOther *CodeStatusNavaidType) (diffs []string) {
	// insertion point for field diffs
	if codestatusnavaidtype.Name != codestatusnavaidtypeOther.Name {
		diffs = append(diffs, codestatusnavaidtype.GongMarshallField(stage, "Name"))
	}
	if codestatusnavaidtype.NilReason != codestatusnavaidtypeOther.NilReason {
		diffs = append(diffs, codestatusnavaidtype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusnavaidtype.EnclosedText != codestatusnavaidtypeOther.EnclosedText {
		diffs = append(diffs, codestatusnavaidtype.GongMarshallField(stage, "EnclosedText"))
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
func (codevortype *CodeVORType) GongDiff(stage *Stage, codevortypeOther *CodeVORType) (diffs []string) {
	// insertion point for field diffs
	if codevortype.Name != codevortypeOther.Name {
		diffs = append(diffs, codevortype.GongMarshallField(stage, "Name"))
	}
	if codevortype.NilReason != codevortypeOther.NilReason {
		diffs = append(diffs, codevortype.GongMarshallField(stage, "NilReason"))
	}
	if codevortype.EnclosedText != codevortypeOther.EnclosedText {
		diffs = append(diffs, codevortype.GongMarshallField(stage, "EnclosedText"))
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
