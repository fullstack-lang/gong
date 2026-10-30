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
func (codecolourtype *CodeColourType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeColourTypes[codecolourtype]
	return ok
}

func (codedaytype *CodeDayType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeDayTypes[codedaytype]
	return ok
}

func (codegroundlightingtype *CodeGroundLightingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeGroundLightingTypes[codegroundlightingtype]
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

func (codenotepurposetype *CodeNotePurposeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeNotePurposeTypes[codenotepurposetype]
	return ok
}

func (codeobstacleareatype *CodeObstacleAreaType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeObstacleAreaTypes[codeobstacleareatype]
	return ok
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeObstacleAssessmentSurfaceTypes[codeobstacleassessmentsurfacetype]
	return ok
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusConstructionTypes[codestatusconstructiontype]
	return ok
}

func (codestatusoperationstype *CodeStatusOperationsType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeStatusOperationsTypes[codestatusoperationstype]
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

func (codeverticaldatumtype *CodeVerticalDatumType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalDatumTypes[codeverticaldatumtype]
	return ok
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalStructureMarkingTypes[codeverticalstructuremarkingtype]
	return ok
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalStructureMaterialTypes[codeverticalstructurematerialtype]
	return ok
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeVerticalStructureTypes[codeverticalstructuretype]
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

func (vallightintensitytype *ValLightIntensityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ValLightIntensityTypes[vallightintensitytype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (codecolourtype *CodeColourType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codecolourtype) {
		return
	}

	codecolourtype.Stage(stage)

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

func (codegroundlightingtype *CodeGroundLightingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codegroundlightingtype) {
		return
	}

	codegroundlightingtype.Stage(stage)

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

func (codenotepurposetype *CodeNotePurposeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codenotepurposetype) {
		return
	}

	codenotepurposetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeobstacleareatype *CodeObstacleAreaType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeobstacleareatype) {
		return
	}

	codeobstacleareatype.Stage(stage)

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

func (codestatusconstructiontype *CodeStatusConstructionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codestatusconstructiontype) {
		return
	}

	codestatusconstructiontype.Stage(stage)

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

func (codeverticaldatumtype *CodeVerticalDatumType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticaldatumtype) {
		return
	}

	codeverticaldatumtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticalstructuremarkingtype) {
		return
	}

	codeverticalstructuremarkingtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticalstructurematerialtype) {
		return
	}

	codeverticalstructurematerialtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructuretype *CodeVerticalStructureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codeverticalstructuretype) {
		return
	}

	codeverticalstructuretype.Stage(stage)

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

func (vallightintensitytype *ValLightIntensityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(vallightintensitytype) {
		return
	}

	vallightintensitytype.Stage(stage)

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
	case *CodeColourType:
		toT := GongCopyBranchCodeColourType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeDayType:
		toT := GongCopyBranchCodeDayType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeGroundLightingType:
		toT := GongCopyBranchCodeGroundLightingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLightIntensityType:
		toT := GongCopyBranchCodeLightIntensityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeLightSourceType:
		toT := GongCopyBranchCodeLightSourceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeNotePurposeType:
		toT := GongCopyBranchCodeNotePurposeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeObstacleAreaType:
		toT := GongCopyBranchCodeObstacleAreaType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeObstacleAssessmentSurfaceType:
		toT := GongCopyBranchCodeObstacleAssessmentSurfaceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusConstructionType:
		toT := GongCopyBranchCodeStatusConstructionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeStatusOperationsType:
		toT := GongCopyBranchCodeStatusOperationsType(mapOrigCopy, fromT)
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

	case *CodeVerticalDatumType:
		toT := GongCopyBranchCodeVerticalDatumType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalStructureMarkingType:
		toT := GongCopyBranchCodeVerticalStructureMarkingType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalStructureMaterialType:
		toT := GongCopyBranchCodeVerticalStructureMaterialType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeVerticalStructureType:
		toT := GongCopyBranchCodeVerticalStructureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeYesNoType:
		toT := GongCopyBranchCodeYesNoType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DateMonthDayType:
		toT := GongCopyBranchDateMonthDayType(mapOrigCopy, fromT)
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

	case *ValLightIntensityType:
		toT := GongCopyBranchValLightIntensityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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

func GongCopyBranchCodeGroundLightingType(mapOrigCopy map[any]any, codegroundlightingtypeFrom *CodeGroundLightingType) (codegroundlightingtypeTo *CodeGroundLightingType) {
	var alreadyCopied bool
	codegroundlightingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codegroundlightingtypeFrom)
	if alreadyCopied {
		return
	}
	codegroundlightingtypeFrom.GongCopyBasicFields(codegroundlightingtypeTo)

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

func GongCopyBranchCodeObstacleAreaType(mapOrigCopy map[any]any, codeobstacleareatypeFrom *CodeObstacleAreaType) (codeobstacleareatypeTo *CodeObstacleAreaType) {
	var alreadyCopied bool
	codeobstacleareatypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeobstacleareatypeFrom)
	if alreadyCopied {
		return
	}
	codeobstacleareatypeFrom.GongCopyBasicFields(codeobstacleareatypeTo)

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

func GongCopyBranchCodeStatusConstructionType(mapOrigCopy map[any]any, codestatusconstructiontypeFrom *CodeStatusConstructionType) (codestatusconstructiontypeTo *CodeStatusConstructionType) {
	var alreadyCopied bool
	codestatusconstructiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codestatusconstructiontypeFrom)
	if alreadyCopied {
		return
	}
	codestatusconstructiontypeFrom.GongCopyBasicFields(codestatusconstructiontypeTo)

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

func GongCopyBranchCodeVerticalStructureMarkingType(mapOrigCopy map[any]any, codeverticalstructuremarkingtypeFrom *CodeVerticalStructureMarkingType) (codeverticalstructuremarkingtypeTo *CodeVerticalStructureMarkingType) {
	var alreadyCopied bool
	codeverticalstructuremarkingtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeverticalstructuremarkingtypeFrom)
	if alreadyCopied {
		return
	}
	codeverticalstructuremarkingtypeFrom.GongCopyBasicFields(codeverticalstructuremarkingtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVerticalStructureMaterialType(mapOrigCopy map[any]any, codeverticalstructurematerialtypeFrom *CodeVerticalStructureMaterialType) (codeverticalstructurematerialtypeTo *CodeVerticalStructureMaterialType) {
	var alreadyCopied bool
	codeverticalstructurematerialtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeverticalstructurematerialtypeFrom)
	if alreadyCopied {
		return
	}
	codeverticalstructurematerialtypeFrom.GongCopyBasicFields(codeverticalstructurematerialtypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeVerticalStructureType(mapOrigCopy map[any]any, codeverticalstructuretypeFrom *CodeVerticalStructureType) (codeverticalstructuretypeTo *CodeVerticalStructureType) {
	var alreadyCopied bool
	codeverticalstructuretypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codeverticalstructuretypeFrom)
	if alreadyCopied {
		return
	}
	codeverticalstructuretypeFrom.GongCopyBasicFields(codeverticalstructuretypeTo)

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
func (codecolourtype *CodeColourType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codecolourtype) {
		return
	}

	codecolourtype.Unstage(stage)

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

func (codegroundlightingtype *CodeGroundLightingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codegroundlightingtype) {
		return
	}

	codegroundlightingtype.Unstage(stage)

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

func (codenotepurposetype *CodeNotePurposeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codenotepurposetype) {
		return
	}

	codenotepurposetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeobstacleareatype *CodeObstacleAreaType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeobstacleareatype) {
		return
	}

	codeobstacleareatype.Unstage(stage)

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

func (codestatusconstructiontype *CodeStatusConstructionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codestatusconstructiontype) {
		return
	}

	codestatusconstructiontype.Unstage(stage)

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

func (codeverticaldatumtype *CodeVerticalDatumType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticaldatumtype) {
		return
	}

	codeverticaldatumtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticalstructuremarkingtype) {
		return
	}

	codeverticalstructuremarkingtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticalstructurematerialtype) {
		return
	}

	codeverticalstructurematerialtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codeverticalstructuretype *CodeVerticalStructureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codeverticalstructuretype) {
		return
	}

	codeverticalstructuretype.Unstage(stage)

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

func (vallightintensitytype *ValLightIntensityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(vallightintensitytype) {
		return
	}

	vallightintensitytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *CodeColourType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeColourType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeDayType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeDayType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeGroundLightingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeGroundLightingType) {
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

func (reference *CodeNotePurposeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeNotePurposeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeObstacleAreaType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeObstacleAreaType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeObstacleAssessmentSurfaceType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeObstacleAssessmentSurfaceType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusConstructionType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusConstructionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeStatusOperationsType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeStatusOperationsType) {
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

func (reference *CodeVerticalDatumType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalDatumType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalStructureMarkingType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalStructureMarkingType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalStructureMaterialType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalStructureMaterialType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeVerticalStructureType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeVerticalStructureType) {
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

func (reference *ValLightIntensityType) GongReconstructPointersFromReferences(stage *Stage, instance *ValLightIntensityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *CodeColourType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeDayType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeGroundLightingType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeNotePurposeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeObstacleAreaType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeObstacleAssessmentSurfaceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusConstructionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeStatusOperationsType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *CodeVerticalDatumType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalStructureMarkingType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalStructureMaterialType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeVerticalStructureType) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *ValLightIntensityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
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
func (codegroundlightingtype *CodeGroundLightingType) GongDiff(stage *Stage, codegroundlightingtypeOther *CodeGroundLightingType) (diffs []string) {
	// insertion point for field diffs
	if codegroundlightingtype.Name != codegroundlightingtypeOther.Name {
		diffs = append(diffs, codegroundlightingtype.GongMarshallField(stage, "Name"))
	}
	if codegroundlightingtype.NilReason != codegroundlightingtypeOther.NilReason {
		diffs = append(diffs, codegroundlightingtype.GongMarshallField(stage, "NilReason"))
	}
	if codegroundlightingtype.EnclosedText != codegroundlightingtypeOther.EnclosedText {
		diffs = append(diffs, codegroundlightingtype.GongMarshallField(stage, "EnclosedText"))
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
func (codeobstacleareatype *CodeObstacleAreaType) GongDiff(stage *Stage, codeobstacleareatypeOther *CodeObstacleAreaType) (diffs []string) {
	// insertion point for field diffs
	if codeobstacleareatype.Name != codeobstacleareatypeOther.Name {
		diffs = append(diffs, codeobstacleareatype.GongMarshallField(stage, "Name"))
	}
	if codeobstacleareatype.NilReason != codeobstacleareatypeOther.NilReason {
		diffs = append(diffs, codeobstacleareatype.GongMarshallField(stage, "NilReason"))
	}
	if codeobstacleareatype.EnclosedText != codeobstacleareatypeOther.EnclosedText {
		diffs = append(diffs, codeobstacleareatype.GongMarshallField(stage, "EnclosedText"))
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
func (codestatusconstructiontype *CodeStatusConstructionType) GongDiff(stage *Stage, codestatusconstructiontypeOther *CodeStatusConstructionType) (diffs []string) {
	// insertion point for field diffs
	if codestatusconstructiontype.Name != codestatusconstructiontypeOther.Name {
		diffs = append(diffs, codestatusconstructiontype.GongMarshallField(stage, "Name"))
	}
	if codestatusconstructiontype.NilReason != codestatusconstructiontypeOther.NilReason {
		diffs = append(diffs, codestatusconstructiontype.GongMarshallField(stage, "NilReason"))
	}
	if codestatusconstructiontype.EnclosedText != codestatusconstructiontypeOther.EnclosedText {
		diffs = append(diffs, codestatusconstructiontype.GongMarshallField(stage, "EnclosedText"))
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
func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongDiff(stage *Stage, codeverticalstructuremarkingtypeOther *CodeVerticalStructureMarkingType) (diffs []string) {
	// insertion point for field diffs
	if codeverticalstructuremarkingtype.Name != codeverticalstructuremarkingtypeOther.Name {
		diffs = append(diffs, codeverticalstructuremarkingtype.GongMarshallField(stage, "Name"))
	}
	if codeverticalstructuremarkingtype.NilReason != codeverticalstructuremarkingtypeOther.NilReason {
		diffs = append(diffs, codeverticalstructuremarkingtype.GongMarshallField(stage, "NilReason"))
	}
	if codeverticalstructuremarkingtype.EnclosedText != codeverticalstructuremarkingtypeOther.EnclosedText {
		diffs = append(diffs, codeverticalstructuremarkingtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongDiff(stage *Stage, codeverticalstructurematerialtypeOther *CodeVerticalStructureMaterialType) (diffs []string) {
	// insertion point for field diffs
	if codeverticalstructurematerialtype.Name != codeverticalstructurematerialtypeOther.Name {
		diffs = append(diffs, codeverticalstructurematerialtype.GongMarshallField(stage, "Name"))
	}
	if codeverticalstructurematerialtype.NilReason != codeverticalstructurematerialtypeOther.NilReason {
		diffs = append(diffs, codeverticalstructurematerialtype.GongMarshallField(stage, "NilReason"))
	}
	if codeverticalstructurematerialtype.EnclosedText != codeverticalstructurematerialtypeOther.EnclosedText {
		diffs = append(diffs, codeverticalstructurematerialtype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codeverticalstructuretype *CodeVerticalStructureType) GongDiff(stage *Stage, codeverticalstructuretypeOther *CodeVerticalStructureType) (diffs []string) {
	// insertion point for field diffs
	if codeverticalstructuretype.Name != codeverticalstructuretypeOther.Name {
		diffs = append(diffs, codeverticalstructuretype.GongMarshallField(stage, "Name"))
	}
	if codeverticalstructuretype.NilReason != codeverticalstructuretypeOther.NilReason {
		diffs = append(diffs, codeverticalstructuretype.GongMarshallField(stage, "NilReason"))
	}
	if codeverticalstructuretype.EnclosedText != codeverticalstructuretypeOther.EnclosedText {
		diffs = append(diffs, codeverticalstructuretype.GongMarshallField(stage, "EnclosedText"))
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
