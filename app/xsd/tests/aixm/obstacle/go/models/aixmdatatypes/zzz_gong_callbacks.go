// generated code - do not edit
package aixmdatatypes

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (codecolourtype *CodeColourType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeColourTypeCreateCallback != nil {
		stage.OnAfterCodeColourTypeCreateCallback.OnAfterCreate(stage, codecolourtype)
	}
}

func (codecolourtype *CodeColourType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeColourTypeUpdateCallback != nil {
		var frontCodeColourType *CodeColourType
		if front != nil {
			frontCodeColourType, _ = front.(*CodeColourType)
		}
		stage.OnAfterCodeColourTypeUpdateCallback.OnAfterUpdate(stage, codecolourtype, frontCodeColourType)
	}
}

func (codecolourtype *CodeColourType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeColourTypeDeleteCallback != nil {
		var frontCodeColourType *CodeColourType
		if front != nil {
			frontCodeColourType, _ = front.(*CodeColourType)
		}
		stage.OnAfterCodeColourTypeDeleteCallback.OnAfterDelete(stage, codecolourtype, frontCodeColourType)
	}
}

func (codedaytype *CodeDayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDayTypeCreateCallback != nil {
		stage.OnAfterCodeDayTypeCreateCallback.OnAfterCreate(stage, codedaytype)
	}
}

func (codedaytype *CodeDayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDayTypeUpdateCallback != nil {
		var frontCodeDayType *CodeDayType
		if front != nil {
			frontCodeDayType, _ = front.(*CodeDayType)
		}
		stage.OnAfterCodeDayTypeUpdateCallback.OnAfterUpdate(stage, codedaytype, frontCodeDayType)
	}
}

func (codedaytype *CodeDayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDayTypeDeleteCallback != nil {
		var frontCodeDayType *CodeDayType
		if front != nil {
			frontCodeDayType, _ = front.(*CodeDayType)
		}
		stage.OnAfterCodeDayTypeDeleteCallback.OnAfterDelete(stage, codedaytype, frontCodeDayType)
	}
}

func (codegroundlightingtype *CodeGroundLightingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeGroundLightingTypeCreateCallback != nil {
		stage.OnAfterCodeGroundLightingTypeCreateCallback.OnAfterCreate(stage, codegroundlightingtype)
	}
}

func (codegroundlightingtype *CodeGroundLightingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeGroundLightingTypeUpdateCallback != nil {
		var frontCodeGroundLightingType *CodeGroundLightingType
		if front != nil {
			frontCodeGroundLightingType, _ = front.(*CodeGroundLightingType)
		}
		stage.OnAfterCodeGroundLightingTypeUpdateCallback.OnAfterUpdate(stage, codegroundlightingtype, frontCodeGroundLightingType)
	}
}

func (codegroundlightingtype *CodeGroundLightingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeGroundLightingTypeDeleteCallback != nil {
		var frontCodeGroundLightingType *CodeGroundLightingType
		if front != nil {
			frontCodeGroundLightingType, _ = front.(*CodeGroundLightingType)
		}
		stage.OnAfterCodeGroundLightingTypeDeleteCallback.OnAfterDelete(stage, codegroundlightingtype, frontCodeGroundLightingType)
	}
}

func (codelightintensitytype *CodeLightIntensityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLightIntensityTypeCreateCallback != nil {
		stage.OnAfterCodeLightIntensityTypeCreateCallback.OnAfterCreate(stage, codelightintensitytype)
	}
}

func (codelightintensitytype *CodeLightIntensityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightIntensityTypeUpdateCallback != nil {
		var frontCodeLightIntensityType *CodeLightIntensityType
		if front != nil {
			frontCodeLightIntensityType, _ = front.(*CodeLightIntensityType)
		}
		stage.OnAfterCodeLightIntensityTypeUpdateCallback.OnAfterUpdate(stage, codelightintensitytype, frontCodeLightIntensityType)
	}
}

func (codelightintensitytype *CodeLightIntensityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightIntensityTypeDeleteCallback != nil {
		var frontCodeLightIntensityType *CodeLightIntensityType
		if front != nil {
			frontCodeLightIntensityType, _ = front.(*CodeLightIntensityType)
		}
		stage.OnAfterCodeLightIntensityTypeDeleteCallback.OnAfterDelete(stage, codelightintensitytype, frontCodeLightIntensityType)
	}
}

func (codelightsourcetype *CodeLightSourceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLightSourceTypeCreateCallback != nil {
		stage.OnAfterCodeLightSourceTypeCreateCallback.OnAfterCreate(stage, codelightsourcetype)
	}
}

func (codelightsourcetype *CodeLightSourceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightSourceTypeUpdateCallback != nil {
		var frontCodeLightSourceType *CodeLightSourceType
		if front != nil {
			frontCodeLightSourceType, _ = front.(*CodeLightSourceType)
		}
		stage.OnAfterCodeLightSourceTypeUpdateCallback.OnAfterUpdate(stage, codelightsourcetype, frontCodeLightSourceType)
	}
}

func (codelightsourcetype *CodeLightSourceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightSourceTypeDeleteCallback != nil {
		var frontCodeLightSourceType *CodeLightSourceType
		if front != nil {
			frontCodeLightSourceType, _ = front.(*CodeLightSourceType)
		}
		stage.OnAfterCodeLightSourceTypeDeleteCallback.OnAfterDelete(stage, codelightsourcetype, frontCodeLightSourceType)
	}
}

func (codenotepurposetype *CodeNotePurposeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNotePurposeTypeCreateCallback != nil {
		stage.OnAfterCodeNotePurposeTypeCreateCallback.OnAfterCreate(stage, codenotepurposetype)
	}
}

func (codenotepurposetype *CodeNotePurposeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNotePurposeTypeUpdateCallback != nil {
		var frontCodeNotePurposeType *CodeNotePurposeType
		if front != nil {
			frontCodeNotePurposeType, _ = front.(*CodeNotePurposeType)
		}
		stage.OnAfterCodeNotePurposeTypeUpdateCallback.OnAfterUpdate(stage, codenotepurposetype, frontCodeNotePurposeType)
	}
}

func (codenotepurposetype *CodeNotePurposeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNotePurposeTypeDeleteCallback != nil {
		var frontCodeNotePurposeType *CodeNotePurposeType
		if front != nil {
			frontCodeNotePurposeType, _ = front.(*CodeNotePurposeType)
		}
		stage.OnAfterCodeNotePurposeTypeDeleteCallback.OnAfterDelete(stage, codenotepurposetype, frontCodeNotePurposeType)
	}
}

func (codeobstacleareatype *CodeObstacleAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeObstacleAreaTypeCreateCallback != nil {
		stage.OnAfterCodeObstacleAreaTypeCreateCallback.OnAfterCreate(stage, codeobstacleareatype)
	}
}

func (codeobstacleareatype *CodeObstacleAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeObstacleAreaTypeUpdateCallback != nil {
		var frontCodeObstacleAreaType *CodeObstacleAreaType
		if front != nil {
			frontCodeObstacleAreaType, _ = front.(*CodeObstacleAreaType)
		}
		stage.OnAfterCodeObstacleAreaTypeUpdateCallback.OnAfterUpdate(stage, codeobstacleareatype, frontCodeObstacleAreaType)
	}
}

func (codeobstacleareatype *CodeObstacleAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeObstacleAreaTypeDeleteCallback != nil {
		var frontCodeObstacleAreaType *CodeObstacleAreaType
		if front != nil {
			frontCodeObstacleAreaType, _ = front.(*CodeObstacleAreaType)
		}
		stage.OnAfterCodeObstacleAreaTypeDeleteCallback.OnAfterDelete(stage, codeobstacleareatype, frontCodeObstacleAreaType)
	}
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeObstacleAssessmentSurfaceTypeCreateCallback != nil {
		stage.OnAfterCodeObstacleAssessmentSurfaceTypeCreateCallback.OnAfterCreate(stage, codeobstacleassessmentsurfacetype)
	}
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeObstacleAssessmentSurfaceTypeUpdateCallback != nil {
		var frontCodeObstacleAssessmentSurfaceType *CodeObstacleAssessmentSurfaceType
		if front != nil {
			frontCodeObstacleAssessmentSurfaceType, _ = front.(*CodeObstacleAssessmentSurfaceType)
		}
		stage.OnAfterCodeObstacleAssessmentSurfaceTypeUpdateCallback.OnAfterUpdate(stage, codeobstacleassessmentsurfacetype, frontCodeObstacleAssessmentSurfaceType)
	}
}

func (codeobstacleassessmentsurfacetype *CodeObstacleAssessmentSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeObstacleAssessmentSurfaceTypeDeleteCallback != nil {
		var frontCodeObstacleAssessmentSurfaceType *CodeObstacleAssessmentSurfaceType
		if front != nil {
			frontCodeObstacleAssessmentSurfaceType, _ = front.(*CodeObstacleAssessmentSurfaceType)
		}
		stage.OnAfterCodeObstacleAssessmentSurfaceTypeDeleteCallback.OnAfterDelete(stage, codeobstacleassessmentsurfacetype, frontCodeObstacleAssessmentSurfaceType)
	}
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusConstructionTypeCreateCallback != nil {
		stage.OnAfterCodeStatusConstructionTypeCreateCallback.OnAfterCreate(stage, codestatusconstructiontype)
	}
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusConstructionTypeUpdateCallback != nil {
		var frontCodeStatusConstructionType *CodeStatusConstructionType
		if front != nil {
			frontCodeStatusConstructionType, _ = front.(*CodeStatusConstructionType)
		}
		stage.OnAfterCodeStatusConstructionTypeUpdateCallback.OnAfterUpdate(stage, codestatusconstructiontype, frontCodeStatusConstructionType)
	}
}

func (codestatusconstructiontype *CodeStatusConstructionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusConstructionTypeDeleteCallback != nil {
		var frontCodeStatusConstructionType *CodeStatusConstructionType
		if front != nil {
			frontCodeStatusConstructionType, _ = front.(*CodeStatusConstructionType)
		}
		stage.OnAfterCodeStatusConstructionTypeDeleteCallback.OnAfterDelete(stage, codestatusconstructiontype, frontCodeStatusConstructionType)
	}
}

func (codestatusoperationstype *CodeStatusOperationsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusOperationsTypeCreateCallback != nil {
		stage.OnAfterCodeStatusOperationsTypeCreateCallback.OnAfterCreate(stage, codestatusoperationstype)
	}
}

func (codestatusoperationstype *CodeStatusOperationsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusOperationsTypeUpdateCallback != nil {
		var frontCodeStatusOperationsType *CodeStatusOperationsType
		if front != nil {
			frontCodeStatusOperationsType, _ = front.(*CodeStatusOperationsType)
		}
		stage.OnAfterCodeStatusOperationsTypeUpdateCallback.OnAfterUpdate(stage, codestatusoperationstype, frontCodeStatusOperationsType)
	}
}

func (codestatusoperationstype *CodeStatusOperationsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusOperationsTypeDeleteCallback != nil {
		var frontCodeStatusOperationsType *CodeStatusOperationsType
		if front != nil {
			frontCodeStatusOperationsType, _ = front.(*CodeStatusOperationsType)
		}
		stage.OnAfterCodeStatusOperationsTypeDeleteCallback.OnAfterDelete(stage, codestatusoperationstype, frontCodeStatusOperationsType)
	}
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTimeEventCombinationTypeCreateCallback != nil {
		stage.OnAfterCodeTimeEventCombinationTypeCreateCallback.OnAfterCreate(stage, codetimeeventcombinationtype)
	}
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeEventCombinationTypeUpdateCallback != nil {
		var frontCodeTimeEventCombinationType *CodeTimeEventCombinationType
		if front != nil {
			frontCodeTimeEventCombinationType, _ = front.(*CodeTimeEventCombinationType)
		}
		stage.OnAfterCodeTimeEventCombinationTypeUpdateCallback.OnAfterUpdate(stage, codetimeeventcombinationtype, frontCodeTimeEventCombinationType)
	}
}

func (codetimeeventcombinationtype *CodeTimeEventCombinationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeEventCombinationTypeDeleteCallback != nil {
		var frontCodeTimeEventCombinationType *CodeTimeEventCombinationType
		if front != nil {
			frontCodeTimeEventCombinationType, _ = front.(*CodeTimeEventCombinationType)
		}
		stage.OnAfterCodeTimeEventCombinationTypeDeleteCallback.OnAfterDelete(stage, codetimeeventcombinationtype, frontCodeTimeEventCombinationType)
	}
}

func (codetimeeventtype *CodeTimeEventType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTimeEventTypeCreateCallback != nil {
		stage.OnAfterCodeTimeEventTypeCreateCallback.OnAfterCreate(stage, codetimeeventtype)
	}
}

func (codetimeeventtype *CodeTimeEventType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeEventTypeUpdateCallback != nil {
		var frontCodeTimeEventType *CodeTimeEventType
		if front != nil {
			frontCodeTimeEventType, _ = front.(*CodeTimeEventType)
		}
		stage.OnAfterCodeTimeEventTypeUpdateCallback.OnAfterUpdate(stage, codetimeeventtype, frontCodeTimeEventType)
	}
}

func (codetimeeventtype *CodeTimeEventType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeEventTypeDeleteCallback != nil {
		var frontCodeTimeEventType *CodeTimeEventType
		if front != nil {
			frontCodeTimeEventType, _ = front.(*CodeTimeEventType)
		}
		stage.OnAfterCodeTimeEventTypeDeleteCallback.OnAfterDelete(stage, codetimeeventtype, frontCodeTimeEventType)
	}
}

func (codetimereferencetype *CodeTimeReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTimeReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeTimeReferenceTypeCreateCallback.OnAfterCreate(stage, codetimereferencetype)
	}
}

func (codetimereferencetype *CodeTimeReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeReferenceTypeUpdateCallback != nil {
		var frontCodeTimeReferenceType *CodeTimeReferenceType
		if front != nil {
			frontCodeTimeReferenceType, _ = front.(*CodeTimeReferenceType)
		}
		stage.OnAfterCodeTimeReferenceTypeUpdateCallback.OnAfterUpdate(stage, codetimereferencetype, frontCodeTimeReferenceType)
	}
}

func (codetimereferencetype *CodeTimeReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTimeReferenceTypeDeleteCallback != nil {
		var frontCodeTimeReferenceType *CodeTimeReferenceType
		if front != nil {
			frontCodeTimeReferenceType, _ = front.(*CodeTimeReferenceType)
		}
		stage.OnAfterCodeTimeReferenceTypeDeleteCallback.OnAfterDelete(stage, codetimereferencetype, frontCodeTimeReferenceType)
	}
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVerticalDatumTypeCreateCallback != nil {
		stage.OnAfterCodeVerticalDatumTypeCreateCallback.OnAfterCreate(stage, codeverticaldatumtype)
	}
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalDatumTypeUpdateCallback != nil {
		var frontCodeVerticalDatumType *CodeVerticalDatumType
		if front != nil {
			frontCodeVerticalDatumType, _ = front.(*CodeVerticalDatumType)
		}
		stage.OnAfterCodeVerticalDatumTypeUpdateCallback.OnAfterUpdate(stage, codeverticaldatumtype, frontCodeVerticalDatumType)
	}
}

func (codeverticaldatumtype *CodeVerticalDatumType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalDatumTypeDeleteCallback != nil {
		var frontCodeVerticalDatumType *CodeVerticalDatumType
		if front != nil {
			frontCodeVerticalDatumType, _ = front.(*CodeVerticalDatumType)
		}
		stage.OnAfterCodeVerticalDatumTypeDeleteCallback.OnAfterDelete(stage, codeverticaldatumtype, frontCodeVerticalDatumType)
	}
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVerticalStructureMarkingTypeCreateCallback != nil {
		stage.OnAfterCodeVerticalStructureMarkingTypeCreateCallback.OnAfterCreate(stage, codeverticalstructuremarkingtype)
	}
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalStructureMarkingTypeUpdateCallback != nil {
		var frontCodeVerticalStructureMarkingType *CodeVerticalStructureMarkingType
		if front != nil {
			frontCodeVerticalStructureMarkingType, _ = front.(*CodeVerticalStructureMarkingType)
		}
		stage.OnAfterCodeVerticalStructureMarkingTypeUpdateCallback.OnAfterUpdate(stage, codeverticalstructuremarkingtype, frontCodeVerticalStructureMarkingType)
	}
}

func (codeverticalstructuremarkingtype *CodeVerticalStructureMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalStructureMarkingTypeDeleteCallback != nil {
		var frontCodeVerticalStructureMarkingType *CodeVerticalStructureMarkingType
		if front != nil {
			frontCodeVerticalStructureMarkingType, _ = front.(*CodeVerticalStructureMarkingType)
		}
		stage.OnAfterCodeVerticalStructureMarkingTypeDeleteCallback.OnAfterDelete(stage, codeverticalstructuremarkingtype, frontCodeVerticalStructureMarkingType)
	}
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVerticalStructureMaterialTypeCreateCallback != nil {
		stage.OnAfterCodeVerticalStructureMaterialTypeCreateCallback.OnAfterCreate(stage, codeverticalstructurematerialtype)
	}
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalStructureMaterialTypeUpdateCallback != nil {
		var frontCodeVerticalStructureMaterialType *CodeVerticalStructureMaterialType
		if front != nil {
			frontCodeVerticalStructureMaterialType, _ = front.(*CodeVerticalStructureMaterialType)
		}
		stage.OnAfterCodeVerticalStructureMaterialTypeUpdateCallback.OnAfterUpdate(stage, codeverticalstructurematerialtype, frontCodeVerticalStructureMaterialType)
	}
}

func (codeverticalstructurematerialtype *CodeVerticalStructureMaterialType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalStructureMaterialTypeDeleteCallback != nil {
		var frontCodeVerticalStructureMaterialType *CodeVerticalStructureMaterialType
		if front != nil {
			frontCodeVerticalStructureMaterialType, _ = front.(*CodeVerticalStructureMaterialType)
		}
		stage.OnAfterCodeVerticalStructureMaterialTypeDeleteCallback.OnAfterDelete(stage, codeverticalstructurematerialtype, frontCodeVerticalStructureMaterialType)
	}
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVerticalStructureTypeCreateCallback != nil {
		stage.OnAfterCodeVerticalStructureTypeCreateCallback.OnAfterCreate(stage, codeverticalstructuretype)
	}
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalStructureTypeUpdateCallback != nil {
		var frontCodeVerticalStructureType *CodeVerticalStructureType
		if front != nil {
			frontCodeVerticalStructureType, _ = front.(*CodeVerticalStructureType)
		}
		stage.OnAfterCodeVerticalStructureTypeUpdateCallback.OnAfterUpdate(stage, codeverticalstructuretype, frontCodeVerticalStructureType)
	}
}

func (codeverticalstructuretype *CodeVerticalStructureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalStructureTypeDeleteCallback != nil {
		var frontCodeVerticalStructureType *CodeVerticalStructureType
		if front != nil {
			frontCodeVerticalStructureType, _ = front.(*CodeVerticalStructureType)
		}
		stage.OnAfterCodeVerticalStructureTypeDeleteCallback.OnAfterDelete(stage, codeverticalstructuretype, frontCodeVerticalStructureType)
	}
}

func (codeyesnotype *CodeYesNoType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeYesNoTypeCreateCallback != nil {
		stage.OnAfterCodeYesNoTypeCreateCallback.OnAfterCreate(stage, codeyesnotype)
	}
}

func (codeyesnotype *CodeYesNoType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeYesNoTypeUpdateCallback != nil {
		var frontCodeYesNoType *CodeYesNoType
		if front != nil {
			frontCodeYesNoType, _ = front.(*CodeYesNoType)
		}
		stage.OnAfterCodeYesNoTypeUpdateCallback.OnAfterUpdate(stage, codeyesnotype, frontCodeYesNoType)
	}
}

func (codeyesnotype *CodeYesNoType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeYesNoTypeDeleteCallback != nil {
		var frontCodeYesNoType *CodeYesNoType
		if front != nil {
			frontCodeYesNoType, _ = front.(*CodeYesNoType)
		}
		stage.OnAfterCodeYesNoTypeDeleteCallback.OnAfterDelete(stage, codeyesnotype, frontCodeYesNoType)
	}
}

func (datemonthdaytype *DateMonthDayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDateMonthDayTypeCreateCallback != nil {
		stage.OnAfterDateMonthDayTypeCreateCallback.OnAfterCreate(stage, datemonthdaytype)
	}
}

func (datemonthdaytype *DateMonthDayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateMonthDayTypeUpdateCallback != nil {
		var frontDateMonthDayType *DateMonthDayType
		if front != nil {
			frontDateMonthDayType, _ = front.(*DateMonthDayType)
		}
		stage.OnAfterDateMonthDayTypeUpdateCallback.OnAfterUpdate(stage, datemonthdaytype, frontDateMonthDayType)
	}
}

func (datemonthdaytype *DateMonthDayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateMonthDayTypeDeleteCallback != nil {
		var frontDateMonthDayType *DateMonthDayType
		if front != nil {
			frontDateMonthDayType, _ = front.(*DateMonthDayType)
		}
		stage.OnAfterDateMonthDayTypeDeleteCallback.OnAfterDelete(stage, datemonthdaytype, frontDateMonthDayType)
	}
}

func (textdesignatortype *TextDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextDesignatorTypeCreateCallback != nil {
		stage.OnAfterTextDesignatorTypeCreateCallback.OnAfterCreate(stage, textdesignatortype)
	}
}

func (textdesignatortype *TextDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextDesignatorTypeUpdateCallback != nil {
		var frontTextDesignatorType *TextDesignatorType
		if front != nil {
			frontTextDesignatorType, _ = front.(*TextDesignatorType)
		}
		stage.OnAfterTextDesignatorTypeUpdateCallback.OnAfterUpdate(stage, textdesignatortype, frontTextDesignatorType)
	}
}

func (textdesignatortype *TextDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextDesignatorTypeDeleteCallback != nil {
		var frontTextDesignatorType *TextDesignatorType
		if front != nil {
			frontTextDesignatorType, _ = front.(*TextDesignatorType)
		}
		stage.OnAfterTextDesignatorTypeDeleteCallback.OnAfterDelete(stage, textdesignatortype, frontTextDesignatorType)
	}
}

func (textnametype *TextNameType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextNameTypeCreateCallback != nil {
		stage.OnAfterTextNameTypeCreateCallback.OnAfterCreate(stage, textnametype)
	}
}

func (textnametype *TextNameType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextNameTypeUpdateCallback != nil {
		var frontTextNameType *TextNameType
		if front != nil {
			frontTextNameType, _ = front.(*TextNameType)
		}
		stage.OnAfterTextNameTypeUpdateCallback.OnAfterUpdate(stage, textnametype, frontTextNameType)
	}
}

func (textnametype *TextNameType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextNameTypeDeleteCallback != nil {
		var frontTextNameType *TextNameType
		if front != nil {
			frontTextNameType, _ = front.(*TextNameType)
		}
		stage.OnAfterTextNameTypeDeleteCallback.OnAfterDelete(stage, textnametype, frontTextNameType)
	}
}

func (textnotetype *TextNoteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextNoteTypeCreateCallback != nil {
		stage.OnAfterTextNoteTypeCreateCallback.OnAfterCreate(stage, textnotetype)
	}
}

func (textnotetype *TextNoteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextNoteTypeUpdateCallback != nil {
		var frontTextNoteType *TextNoteType
		if front != nil {
			frontTextNoteType, _ = front.(*TextNoteType)
		}
		stage.OnAfterTextNoteTypeUpdateCallback.OnAfterUpdate(stage, textnotetype, frontTextNoteType)
	}
}

func (textnotetype *TextNoteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextNoteTypeDeleteCallback != nil {
		var frontTextNoteType *TextNoteType
		if front != nil {
			frontTextNoteType, _ = front.(*TextNoteType)
		}
		stage.OnAfterTextNoteTypeDeleteCallback.OnAfterDelete(stage, textnotetype, frontTextNoteType)
	}
}

func (textpropertynametype *TextPropertyNameType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextPropertyNameTypeCreateCallback != nil {
		stage.OnAfterTextPropertyNameTypeCreateCallback.OnAfterCreate(stage, textpropertynametype)
	}
}

func (textpropertynametype *TextPropertyNameType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextPropertyNameTypeUpdateCallback != nil {
		var frontTextPropertyNameType *TextPropertyNameType
		if front != nil {
			frontTextPropertyNameType, _ = front.(*TextPropertyNameType)
		}
		stage.OnAfterTextPropertyNameTypeUpdateCallback.OnAfterUpdate(stage, textpropertynametype, frontTextPropertyNameType)
	}
}

func (textpropertynametype *TextPropertyNameType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextPropertyNameTypeDeleteCallback != nil {
		var frontTextPropertyNameType *TextPropertyNameType
		if front != nil {
			frontTextPropertyNameType, _ = front.(*TextPropertyNameType)
		}
		stage.OnAfterTextPropertyNameTypeDeleteCallback.OnAfterDelete(stage, textpropertynametype, frontTextPropertyNameType)
	}
}

func (timetype *TimeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeTypeCreateCallback != nil {
		stage.OnAfterTimeTypeCreateCallback.OnAfterCreate(stage, timetype)
	}
}

func (timetype *TimeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTypeUpdateCallback != nil {
		var frontTimeType *TimeType
		if front != nil {
			frontTimeType, _ = front.(*TimeType)
		}
		stage.OnAfterTimeTypeUpdateCallback.OnAfterUpdate(stage, timetype, frontTimeType)
	}
}

func (timetype *TimeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTypeDeleteCallback != nil {
		var frontTimeType *TimeType
		if front != nil {
			frontTimeType, _ = front.(*TimeType)
		}
		stage.OnAfterTimeTypeDeleteCallback.OnAfterDelete(stage, timetype, frontTimeType)
	}
}

func (valdistancesignedtype *ValDistanceSignedType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDistanceSignedTypeCreateCallback != nil {
		stage.OnAfterValDistanceSignedTypeCreateCallback.OnAfterCreate(stage, valdistancesignedtype)
	}
}

func (valdistancesignedtype *ValDistanceSignedType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceSignedTypeUpdateCallback != nil {
		var frontValDistanceSignedType *ValDistanceSignedType
		if front != nil {
			frontValDistanceSignedType, _ = front.(*ValDistanceSignedType)
		}
		stage.OnAfterValDistanceSignedTypeUpdateCallback.OnAfterUpdate(stage, valdistancesignedtype, frontValDistanceSignedType)
	}
}

func (valdistancesignedtype *ValDistanceSignedType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceSignedTypeDeleteCallback != nil {
		var frontValDistanceSignedType *ValDistanceSignedType
		if front != nil {
			frontValDistanceSignedType, _ = front.(*ValDistanceSignedType)
		}
		stage.OnAfterValDistanceSignedTypeDeleteCallback.OnAfterDelete(stage, valdistancesignedtype, frontValDistanceSignedType)
	}
}

func (valdistancetype *ValDistanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDistanceTypeCreateCallback != nil {
		stage.OnAfterValDistanceTypeCreateCallback.OnAfterCreate(stage, valdistancetype)
	}
}

func (valdistancetype *ValDistanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceTypeUpdateCallback != nil {
		var frontValDistanceType *ValDistanceType
		if front != nil {
			frontValDistanceType, _ = front.(*ValDistanceType)
		}
		stage.OnAfterValDistanceTypeUpdateCallback.OnAfterUpdate(stage, valdistancetype, frontValDistanceType)
	}
}

func (valdistancetype *ValDistanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceTypeDeleteCallback != nil {
		var frontValDistanceType *ValDistanceType
		if front != nil {
			frontValDistanceType, _ = front.(*ValDistanceType)
		}
		stage.OnAfterValDistanceTypeDeleteCallback.OnAfterDelete(stage, valdistancetype, frontValDistanceType)
	}
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDistanceVerticalTypeCreateCallback != nil {
		stage.OnAfterValDistanceVerticalTypeCreateCallback.OnAfterCreate(stage, valdistanceverticaltype)
	}
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceVerticalTypeUpdateCallback != nil {
		var frontValDistanceVerticalType *ValDistanceVerticalType
		if front != nil {
			frontValDistanceVerticalType, _ = front.(*ValDistanceVerticalType)
		}
		stage.OnAfterValDistanceVerticalTypeUpdateCallback.OnAfterUpdate(stage, valdistanceverticaltype, frontValDistanceVerticalType)
	}
}

func (valdistanceverticaltype *ValDistanceVerticalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDistanceVerticalTypeDeleteCallback != nil {
		var frontValDistanceVerticalType *ValDistanceVerticalType
		if front != nil {
			frontValDistanceVerticalType, _ = front.(*ValDistanceVerticalType)
		}
		stage.OnAfterValDistanceVerticalTypeDeleteCallback.OnAfterDelete(stage, valdistanceverticaltype, frontValDistanceVerticalType)
	}
}

func (valdurationtype *ValDurationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDurationTypeCreateCallback != nil {
		stage.OnAfterValDurationTypeCreateCallback.OnAfterCreate(stage, valdurationtype)
	}
}

func (valdurationtype *ValDurationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDurationTypeUpdateCallback != nil {
		var frontValDurationType *ValDurationType
		if front != nil {
			frontValDurationType, _ = front.(*ValDurationType)
		}
		stage.OnAfterValDurationTypeUpdateCallback.OnAfterUpdate(stage, valdurationtype, frontValDurationType)
	}
}

func (valdurationtype *ValDurationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDurationTypeDeleteCallback != nil {
		var frontValDurationType *ValDurationType
		if front != nil {
			frontValDurationType, _ = front.(*ValDurationType)
		}
		stage.OnAfterValDurationTypeDeleteCallback.OnAfterDelete(stage, valdurationtype, frontValDurationType)
	}
}

func (vallightintensitytype *ValLightIntensityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValLightIntensityTypeCreateCallback != nil {
		stage.OnAfterValLightIntensityTypeCreateCallback.OnAfterCreate(stage, vallightintensitytype)
	}
}

func (vallightintensitytype *ValLightIntensityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValLightIntensityTypeUpdateCallback != nil {
		var frontValLightIntensityType *ValLightIntensityType
		if front != nil {
			frontValLightIntensityType, _ = front.(*ValLightIntensityType)
		}
		stage.OnAfterValLightIntensityTypeUpdateCallback.OnAfterUpdate(stage, vallightintensitytype, frontValLightIntensityType)
	}
}

func (vallightintensitytype *ValLightIntensityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValLightIntensityTypeDeleteCallback != nil {
		var frontValLightIntensityType *ValLightIntensityType
		if front != nil {
			frontValLightIntensityType, _ = front.(*ValLightIntensityType)
		}
		stage.OnAfterValLightIntensityTypeDeleteCallback.OnAfterDelete(stage, vallightintensitytype, frontValLightIntensityType)
	}
}

