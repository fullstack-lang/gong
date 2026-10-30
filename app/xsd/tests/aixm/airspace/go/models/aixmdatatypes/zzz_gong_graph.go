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

func (codeaircrafttype *CodeAircraftType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftTypes[codeaircrafttype]
	return ok
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAircraftWingspanClassTypes[codeaircraftwingspanclasstype]
	return ok
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceActivityTypes[codeairspaceactivitytype]
	return ok
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceAggregationTypes[codeairspaceaggregationtype]
	return ok
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceClassificationTypes[codeairspaceclassificationtype]
	return ok
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceDependencyTypes[codeairspacedependencytype]
	return ok
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceDesignatorTypes[codeairspacedesignatortype]
	return ok
}

func (codeairspacetype *CodeAirspaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeAirspaceTypes[codeairspacetype]
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

func (codecommunicationmodetype *CodeCommunicationModeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeCommunicationModeTypes[codecommunicationmodetype]
	return ok
}

func (codedaytype *CodeDayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDayTypes[codedaytype]
	return ok
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeEquipmentAntiCollisionTypes[codeequipmentanticollisiontype]
	return ok
}

func (codegeobordertype *CodeGeoBorderType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeGeoBorderTypes[codegeobordertype]
	return ok
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeMilitaryOperationsTypes[codemilitaryoperationstype]
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

func (codervsmtype *CodeRVSMType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeRVSMTypes[codervsmtype]
	return ok
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusAirspaceTypes[codestatusairspacetype]
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

func (timetype *TimeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimeTypes[timetype]
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

func (valspeedtype *ValSpeedType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValSpeedTypes[valspeedtype]
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

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspaceactivitytype) {
		return
	}

	codeairspaceactivitytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspaceaggregationtype) {
		return
	}

	codeairspaceaggregationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspaceclassificationtype) {
		return
	}

	codeairspaceclassificationtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacedependencytype) {
		return
	}

	codeairspacedependencytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacedesignatortype) {
		return
	}

	codeairspacedesignatortype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacetype *CodeAirspaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeairspacetype) {
		return
	}

	codeairspacetype.Stage(stage)

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

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeequipmentanticollisiontype) {
		return
	}

	codeequipmentanticollisiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codegeobordertype *CodeGeoBorderType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codegeobordertype) {
		return
	}

	codegeobordertype.Stage(stage)

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

func (codervsmtype *CodeRVSMType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codervsmtype) {
		return
	}

	codervsmtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusairspacetype *CodeStatusAirspaceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusairspacetype) {
		return
	}

	codestatusairspacetype.Stage(stage)

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

func (timetype *TimeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timetype) {
		return
	}

	timetype.Stage(stage)

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

	case *CodeAircraftType:
		toT := GongCopyBranchCodeAircraftType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAircraftWingspanClassType:
		toT := GongCopyBranchCodeAircraftWingspanClassType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceActivityType:
		toT := GongCopyBranchCodeAirspaceActivityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceAggregationType:
		toT := GongCopyBranchCodeAirspaceAggregationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceClassificationType:
		toT := GongCopyBranchCodeAirspaceClassificationType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceDependencyType:
		toT := GongCopyBranchCodeAirspaceDependencyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceDesignatorType:
		toT := GongCopyBranchCodeAirspaceDesignatorType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAirspaceType:
		toT := GongCopyBranchCodeAirspaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAltitudeUseType:
		toT := GongCopyBranchCodeAltitudeUseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeAuthorityType:
		toT := GongCopyBranchCodeAuthorityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeCommunicationModeType:
		toT := GongCopyBranchCodeCommunicationModeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDayType:
		toT := GongCopyBranchCodeDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeEquipmentAntiCollisionType:
		toT := GongCopyBranchCodeEquipmentAntiCollisionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeGeoBorderType:
		toT := GongCopyBranchCodeGeoBorderType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeMilitaryOperationsType:
		toT := GongCopyBranchCodeMilitaryOperationsType(mapOrigCopy, fromT)
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

	case *CodeRVSMType:
		toT := GongCopyBranchCodeRVSMType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusAirspaceType:
		toT := GongCopyBranchCodeStatusAirspaceType(mapOrigCopy, fromT)
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

	case *TextNameType:
		toT := GongCopyBranchTextNameType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextNoteType:
		toT := GongCopyBranchTextNoteType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TextPropertyNameType:
		toT := GongCopyBranchTextPropertyNameType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimeType:
		toT := GongCopyBranchTimeType(mapOrigCopy, fromT)
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

	case *ValSpeedType:
		toT := GongCopyBranchValSpeedType(mapOrigCopy, fromT)
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

func GongCopyBranchCodeAirspaceActivityType(mapOrigCopy map[any]any, codeairspaceactivitytypeFrom *CodeAirspaceActivityType) (codeairspaceactivitytypeTo *CodeAirspaceActivityType) {
	var alreadyCopied bool
	codeairspaceactivitytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspaceactivitytypeFrom)
	if alreadyCopied {
		return
	}
	codeairspaceactivitytypeFrom.GongCopyBasicFields(codeairspaceactivitytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceAggregationType(mapOrigCopy map[any]any, codeairspaceaggregationtypeFrom *CodeAirspaceAggregationType) (codeairspaceaggregationtypeTo *CodeAirspaceAggregationType) {
	var alreadyCopied bool
	codeairspaceaggregationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspaceaggregationtypeFrom)
	if alreadyCopied {
		return
	}
	codeairspaceaggregationtypeFrom.GongCopyBasicFields(codeairspaceaggregationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceClassificationType(mapOrigCopy map[any]any, codeairspaceclassificationtypeFrom *CodeAirspaceClassificationType) (codeairspaceclassificationtypeTo *CodeAirspaceClassificationType) {
	var alreadyCopied bool
	codeairspaceclassificationtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspaceclassificationtypeFrom)
	if alreadyCopied {
		return
	}
	codeairspaceclassificationtypeFrom.GongCopyBasicFields(codeairspaceclassificationtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceDependencyType(mapOrigCopy map[any]any, codeairspacedependencytypeFrom *CodeAirspaceDependencyType) (codeairspacedependencytypeTo *CodeAirspaceDependencyType) {
	var alreadyCopied bool
	codeairspacedependencytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacedependencytypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacedependencytypeFrom.GongCopyBasicFields(codeairspacedependencytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceDesignatorType(mapOrigCopy map[any]any, codeairspacedesignatortypeFrom *CodeAirspaceDesignatorType) (codeairspacedesignatortypeTo *CodeAirspaceDesignatorType) {
	var alreadyCopied bool
	codeairspacedesignatortypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacedesignatortypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacedesignatortypeFrom.GongCopyBasicFields(codeairspacedesignatortypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeAirspaceType(mapOrigCopy map[any]any, codeairspacetypeFrom *CodeAirspaceType) (codeairspacetypeTo *CodeAirspaceType) {
	var alreadyCopied bool
	codeairspacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeairspacetypeFrom)
	if alreadyCopied {
		return
	}
	codeairspacetypeFrom.GongCopyBasicFields(codeairspacetypeTo)

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

func GongCopyBranchCodeGeoBorderType(mapOrigCopy map[any]any, codegeobordertypeFrom *CodeGeoBorderType) (codegeobordertypeTo *CodeGeoBorderType) {
	var alreadyCopied bool
	codegeobordertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codegeobordertypeFrom)
	if alreadyCopied {
		return
	}
	codegeobordertypeFrom.GongCopyBasicFields(codegeobordertypeTo)

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

func GongCopyBranchCodeStatusAirspaceType(mapOrigCopy map[any]any, codestatusairspacetypeFrom *CodeStatusAirspaceType) (codestatusairspacetypeTo *CodeStatusAirspaceType) {
	var alreadyCopied bool
	codestatusairspacetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusairspacetypeFrom)
	if alreadyCopied {
		return
	}
	codestatusairspacetypeFrom.GongCopyBasicFields(codestatusairspacetypeTo)

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

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspaceactivitytype) {
		return
	}

	codeairspaceactivitytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspaceaggregationtype) {
		return
	}

	codeairspaceaggregationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspaceclassificationtype) {
		return
	}

	codeairspaceclassificationtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacedependencytype) {
		return
	}

	codeairspacedependencytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacedesignatortype) {
		return
	}

	codeairspacedesignatortype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeairspacetype *CodeAirspaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeairspacetype) {
		return
	}

	codeairspacetype.Unstage(stage)

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

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeequipmentanticollisiontype) {
		return
	}

	codeequipmentanticollisiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codegeobordertype *CodeGeoBorderType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codegeobordertype) {
		return
	}

	codegeobordertype.Unstage(stage)

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

func (codervsmtype *CodeRVSMType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codervsmtype) {
		return
	}

	codervsmtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codestatusairspacetype *CodeStatusAirspaceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusairspacetype) {
		return
	}

	codestatusairspacetype.Unstage(stage)

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

func (timetype *TimeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timetype) {
		return
	}

	timetype.Unstage(stage)

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

func (reference *CodeAircraftType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAircraftWingspanClassType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAircraftWingspanClassType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceActivityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceActivityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceAggregationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceAggregationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceClassificationType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceClassificationType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceDependencyType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceDependencyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceDesignatorType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceDesignatorType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeAirspaceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeAirspaceType) {
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

func (reference *CodeCommunicationModeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeCommunicationModeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDayType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeEquipmentAntiCollisionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeEquipmentAntiCollisionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeGeoBorderType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeGeoBorderType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeMilitaryOperationsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeMilitaryOperationsType) {
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

func (reference *CodeRVSMType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeRVSMType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusAirspaceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusAirspaceType) {
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

func (reference *TimeType) GongReconstructPointersFromReferences(stage *Stage, instance *TimeType) {
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

func (reference *ValSpeedType) GongReconstructPointersFromReferences(stage *Stage, instance *ValSpeedType) {
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

func (reference *CodeAircraftType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAircraftWingspanClassType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceActivityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceAggregationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceClassificationType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceDependencyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceDesignatorType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeAirspaceType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeCommunicationModeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDayType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeEquipmentAntiCollisionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeGeoBorderType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeMilitaryOperationsType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeRVSMType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusAirspaceType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *TimeType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *ValSpeedType) GongReconstructPointersFromInstances(stage *Stage) {
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
func (codeairspaceactivitytype *CodeAirspaceActivityType) GongDiff(stage *Stage, codeairspaceactivitytypeOther *CodeAirspaceActivityType) (diffs []string) {
	// insertion point for field diffs
	if codeairspaceactivitytype.Name != codeairspaceactivitytypeOther.Name {
		diffs = append(diffs, codeairspaceactivitytype.GongMarshallField(stage, "Name"))
	}
	if codeairspaceactivitytype.NilReason != codeairspaceactivitytypeOther.NilReason {
		diffs = append(diffs, codeairspaceactivitytype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspaceactivitytype.EnclosedText != codeairspaceactivitytypeOther.EnclosedText {
		diffs = append(diffs, codeairspaceactivitytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongDiff(stage *Stage, codeairspaceaggregationtypeOther *CodeAirspaceAggregationType) (diffs []string) {
	// insertion point for field diffs
	if codeairspaceaggregationtype.Name != codeairspaceaggregationtypeOther.Name {
		diffs = append(diffs, codeairspaceaggregationtype.GongMarshallField(stage, "Name"))
	}
	if codeairspaceaggregationtype.NilReason != codeairspaceaggregationtypeOther.NilReason {
		diffs = append(diffs, codeairspaceaggregationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspaceaggregationtype.EnclosedText != codeairspaceaggregationtypeOther.EnclosedText {
		diffs = append(diffs, codeairspaceaggregationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongDiff(stage *Stage, codeairspaceclassificationtypeOther *CodeAirspaceClassificationType) (diffs []string) {
	// insertion point for field diffs
	if codeairspaceclassificationtype.Name != codeairspaceclassificationtypeOther.Name {
		diffs = append(diffs, codeairspaceclassificationtype.GongMarshallField(stage, "Name"))
	}
	if codeairspaceclassificationtype.NilReason != codeairspaceclassificationtypeOther.NilReason {
		diffs = append(diffs, codeairspaceclassificationtype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspaceclassificationtype.EnclosedText != codeairspaceclassificationtypeOther.EnclosedText {
		diffs = append(diffs, codeairspaceclassificationtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacedependencytype *CodeAirspaceDependencyType) GongDiff(stage *Stage, codeairspacedependencytypeOther *CodeAirspaceDependencyType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacedependencytype.Name != codeairspacedependencytypeOther.Name {
		diffs = append(diffs, codeairspacedependencytype.GongMarshallField(stage, "Name"))
	}
	if codeairspacedependencytype.NilReason != codeairspacedependencytypeOther.NilReason {
		diffs = append(diffs, codeairspacedependencytype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacedependencytype.EnclosedText != codeairspacedependencytypeOther.EnclosedText {
		diffs = append(diffs, codeairspacedependencytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongDiff(stage *Stage, codeairspacedesignatortypeOther *CodeAirspaceDesignatorType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacedesignatortype.Name != codeairspacedesignatortypeOther.Name {
		diffs = append(diffs, codeairspacedesignatortype.GongMarshallField(stage, "Name"))
	}
	if codeairspacedesignatortype.NilReason != codeairspacedesignatortypeOther.NilReason {
		diffs = append(diffs, codeairspacedesignatortype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacedesignatortype.EnclosedText != codeairspacedesignatortypeOther.EnclosedText {
		diffs = append(diffs, codeairspacedesignatortype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeairspacetype *CodeAirspaceType) GongDiff(stage *Stage, codeairspacetypeOther *CodeAirspaceType) (diffs []string) {
	// insertion point for field diffs
	if codeairspacetype.Name != codeairspacetypeOther.Name {
		diffs = append(diffs, codeairspacetype.GongMarshallField(stage, "Name"))
	}
	if codeairspacetype.NilReason != codeairspacetypeOther.NilReason {
		diffs = append(diffs, codeairspacetype.GongMarshallField(stage, "NilReason"))
	}
	if codeairspacetype.EnclosedText != codeairspacetypeOther.EnclosedText {
		diffs = append(diffs, codeairspacetype.GongMarshallField(stage, "EnclosedText"))
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
func (codegeobordertype *CodeGeoBorderType) GongDiff(stage *Stage, codegeobordertypeOther *CodeGeoBorderType) (diffs []string) {
	// insertion point for field diffs
	if codegeobordertype.Name != codegeobordertypeOther.Name {
		diffs = append(diffs, codegeobordertype.GongMarshallField(stage, "Name"))
	}
	if codegeobordertype.NilReason != codegeobordertypeOther.NilReason {
		diffs = append(diffs, codegeobordertype.GongMarshallField(stage, "NilReason"))
	}
	if codegeobordertype.EnclosedText != codegeobordertypeOther.EnclosedText {
		diffs = append(diffs, codegeobordertype.GongMarshallField(stage, "EnclosedText"))
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
func (codestatusairspacetype *CodeStatusAirspaceType) GongDiff(stage *Stage, codestatusairspacetypeOther *CodeStatusAirspaceType) (diffs []string) {
	// insertion point for field diffs
	if codestatusairspacetype.Name != codestatusairspacetypeOther.Name {
		diffs = append(diffs, codestatusairspacetype.GongMarshallField(stage, "Name"))
	}
	if codestatusairspacetype.NilReason != codestatusairspacetypeOther.NilReason {
		diffs = append(diffs, codestatusairspacetype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusairspacetype.EnclosedText != codestatusairspacetypeOther.EnclosedText {
		diffs = append(diffs, codestatusairspacetype.GongMarshallField(stage, "EnclosedText"))
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
