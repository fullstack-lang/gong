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
func (codeatcreportingtype *CodeATCReportingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeATCReportingTypeCreateCallback != nil {
		stage.OnAfterCodeATCReportingTypeCreateCallback.OnAfterCreate(stage, codeatcreportingtype)
	}
}

func (codeatcreportingtype *CodeATCReportingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeATCReportingTypeUpdateCallback != nil {
		var frontCodeATCReportingType *CodeATCReportingType
		if front != nil {
			frontCodeATCReportingType, _ = front.(*CodeATCReportingType)
		}
		stage.OnAfterCodeATCReportingTypeUpdateCallback.OnAfterUpdate(stage, codeatcreportingtype, frontCodeATCReportingType)
	}
}

func (codeatcreportingtype *CodeATCReportingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeATCReportingTypeDeleteCallback != nil {
		var frontCodeATCReportingType *CodeATCReportingType
		if front != nil {
			frontCodeATCReportingType, _ = front.(*CodeATCReportingType)
		}
		stage.OnAfterCodeATCReportingTypeDeleteCallback.OnAfterDelete(stage, codeatcreportingtype, frontCodeATCReportingType)
	}
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAerialRefuellingPointTypeCreateCallback != nil {
		stage.OnAfterCodeAerialRefuellingPointTypeCreateCallback.OnAfterCreate(stage, codeaerialrefuellingpointtype)
	}
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAerialRefuellingPointTypeUpdateCallback != nil {
		var frontCodeAerialRefuellingPointType *CodeAerialRefuellingPointType
		if front != nil {
			frontCodeAerialRefuellingPointType, _ = front.(*CodeAerialRefuellingPointType)
		}
		stage.OnAfterCodeAerialRefuellingPointTypeUpdateCallback.OnAfterUpdate(stage, codeaerialrefuellingpointtype, frontCodeAerialRefuellingPointType)
	}
}

func (codeaerialrefuellingpointtype *CodeAerialRefuellingPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAerialRefuellingPointTypeDeleteCallback != nil {
		var frontCodeAerialRefuellingPointType *CodeAerialRefuellingPointType
		if front != nil {
			frontCodeAerialRefuellingPointType, _ = front.(*CodeAerialRefuellingPointType)
		}
		stage.OnAfterCodeAerialRefuellingPointTypeDeleteCallback.OnAfterDelete(stage, codeaerialrefuellingpointtype, frontCodeAerialRefuellingPointType)
	}
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAerialRefuellingPrefixTypeCreateCallback != nil {
		stage.OnAfterCodeAerialRefuellingPrefixTypeCreateCallback.OnAfterCreate(stage, codeaerialrefuellingprefixtype)
	}
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAerialRefuellingPrefixTypeUpdateCallback != nil {
		var frontCodeAerialRefuellingPrefixType *CodeAerialRefuellingPrefixType
		if front != nil {
			frontCodeAerialRefuellingPrefixType, _ = front.(*CodeAerialRefuellingPrefixType)
		}
		stage.OnAfterCodeAerialRefuellingPrefixTypeUpdateCallback.OnAfterUpdate(stage, codeaerialrefuellingprefixtype, frontCodeAerialRefuellingPrefixType)
	}
}

func (codeaerialrefuellingprefixtype *CodeAerialRefuellingPrefixType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAerialRefuellingPrefixTypeDeleteCallback != nil {
		var frontCodeAerialRefuellingPrefixType *CodeAerialRefuellingPrefixType
		if front != nil {
			frontCodeAerialRefuellingPrefixType, _ = front.(*CodeAerialRefuellingPrefixType)
		}
		stage.OnAfterCodeAerialRefuellingPrefixTypeDeleteCallback.OnAfterDelete(stage, codeaerialrefuellingprefixtype, frontCodeAerialRefuellingPrefixType)
	}
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAerialRefuellingTypeCreateCallback != nil {
		stage.OnAfterCodeAerialRefuellingTypeCreateCallback.OnAfterCreate(stage, codeaerialrefuellingtype)
	}
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAerialRefuellingTypeUpdateCallback != nil {
		var frontCodeAerialRefuellingType *CodeAerialRefuellingType
		if front != nil {
			frontCodeAerialRefuellingType, _ = front.(*CodeAerialRefuellingType)
		}
		stage.OnAfterCodeAerialRefuellingTypeUpdateCallback.OnAfterUpdate(stage, codeaerialrefuellingtype, frontCodeAerialRefuellingType)
	}
}

func (codeaerialrefuellingtype *CodeAerialRefuellingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAerialRefuellingTypeDeleteCallback != nil {
		var frontCodeAerialRefuellingType *CodeAerialRefuellingType
		if front != nil {
			frontCodeAerialRefuellingType, _ = front.(*CodeAerialRefuellingType)
		}
		stage.OnAfterCodeAerialRefuellingTypeDeleteCallback.OnAfterDelete(stage, codeaerialrefuellingtype, frontCodeAerialRefuellingType)
	}
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftCategoryTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftCategoryTypeCreateCallback.OnAfterCreate(stage, codeaircraftcategorytype)
	}
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftCategoryTypeUpdateCallback != nil {
		var frontCodeAircraftCategoryType *CodeAircraftCategoryType
		if front != nil {
			frontCodeAircraftCategoryType, _ = front.(*CodeAircraftCategoryType)
		}
		stage.OnAfterCodeAircraftCategoryTypeUpdateCallback.OnAfterUpdate(stage, codeaircraftcategorytype, frontCodeAircraftCategoryType)
	}
}

func (codeaircraftcategorytype *CodeAircraftCategoryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftCategoryTypeDeleteCallback != nil {
		var frontCodeAircraftCategoryType *CodeAircraftCategoryType
		if front != nil {
			frontCodeAircraftCategoryType, _ = front.(*CodeAircraftCategoryType)
		}
		stage.OnAfterCodeAircraftCategoryTypeDeleteCallback.OnAfterDelete(stage, codeaircraftcategorytype, frontCodeAircraftCategoryType)
	}
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftEngineNumberTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftEngineNumberTypeCreateCallback.OnAfterCreate(stage, codeaircraftenginenumbertype)
	}
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftEngineNumberTypeUpdateCallback != nil {
		var frontCodeAircraftEngineNumberType *CodeAircraftEngineNumberType
		if front != nil {
			frontCodeAircraftEngineNumberType, _ = front.(*CodeAircraftEngineNumberType)
		}
		stage.OnAfterCodeAircraftEngineNumberTypeUpdateCallback.OnAfterUpdate(stage, codeaircraftenginenumbertype, frontCodeAircraftEngineNumberType)
	}
}

func (codeaircraftenginenumbertype *CodeAircraftEngineNumberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftEngineNumberTypeDeleteCallback != nil {
		var frontCodeAircraftEngineNumberType *CodeAircraftEngineNumberType
		if front != nil {
			frontCodeAircraftEngineNumberType, _ = front.(*CodeAircraftEngineNumberType)
		}
		stage.OnAfterCodeAircraftEngineNumberTypeDeleteCallback.OnAfterDelete(stage, codeaircraftenginenumbertype, frontCodeAircraftEngineNumberType)
	}
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftEngineTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftEngineTypeCreateCallback.OnAfterCreate(stage, codeaircraftenginetype)
	}
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftEngineTypeUpdateCallback != nil {
		var frontCodeAircraftEngineType *CodeAircraftEngineType
		if front != nil {
			frontCodeAircraftEngineType, _ = front.(*CodeAircraftEngineType)
		}
		stage.OnAfterCodeAircraftEngineTypeUpdateCallback.OnAfterUpdate(stage, codeaircraftenginetype, frontCodeAircraftEngineType)
	}
}

func (codeaircraftenginetype *CodeAircraftEngineType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftEngineTypeDeleteCallback != nil {
		var frontCodeAircraftEngineType *CodeAircraftEngineType
		if front != nil {
			frontCodeAircraftEngineType, _ = front.(*CodeAircraftEngineType)
		}
		stage.OnAfterCodeAircraftEngineTypeDeleteCallback.OnAfterDelete(stage, codeaircraftenginetype, frontCodeAircraftEngineType)
	}
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftICAOTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftICAOTypeCreateCallback.OnAfterCreate(stage, codeaircrafticaotype)
	}
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftICAOTypeUpdateCallback != nil {
		var frontCodeAircraftICAOType *CodeAircraftICAOType
		if front != nil {
			frontCodeAircraftICAOType, _ = front.(*CodeAircraftICAOType)
		}
		stage.OnAfterCodeAircraftICAOTypeUpdateCallback.OnAfterUpdate(stage, codeaircrafticaotype, frontCodeAircraftICAOType)
	}
}

func (codeaircrafticaotype *CodeAircraftICAOType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftICAOTypeDeleteCallback != nil {
		var frontCodeAircraftICAOType *CodeAircraftICAOType
		if front != nil {
			frontCodeAircraftICAOType, _ = front.(*CodeAircraftICAOType)
		}
		stage.OnAfterCodeAircraftICAOTypeDeleteCallback.OnAfterDelete(stage, codeaircrafticaotype, frontCodeAircraftICAOType)
	}
}

func (codeaircrafttype *CodeAircraftType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftTypeCreateCallback.OnAfterCreate(stage, codeaircrafttype)
	}
}

func (codeaircrafttype *CodeAircraftType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftTypeUpdateCallback != nil {
		var frontCodeAircraftType *CodeAircraftType
		if front != nil {
			frontCodeAircraftType, _ = front.(*CodeAircraftType)
		}
		stage.OnAfterCodeAircraftTypeUpdateCallback.OnAfterUpdate(stage, codeaircrafttype, frontCodeAircraftType)
	}
}

func (codeaircrafttype *CodeAircraftType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftTypeDeleteCallback != nil {
		var frontCodeAircraftType *CodeAircraftType
		if front != nil {
			frontCodeAircraftType, _ = front.(*CodeAircraftType)
		}
		stage.OnAfterCodeAircraftTypeDeleteCallback.OnAfterDelete(stage, codeaircrafttype, frontCodeAircraftType)
	}
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftWingspanClassTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftWingspanClassTypeCreateCallback.OnAfterCreate(stage, codeaircraftwingspanclasstype)
	}
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftWingspanClassTypeUpdateCallback != nil {
		var frontCodeAircraftWingspanClassType *CodeAircraftWingspanClassType
		if front != nil {
			frontCodeAircraftWingspanClassType, _ = front.(*CodeAircraftWingspanClassType)
		}
		stage.OnAfterCodeAircraftWingspanClassTypeUpdateCallback.OnAfterUpdate(stage, codeaircraftwingspanclasstype, frontCodeAircraftWingspanClassType)
	}
}

func (codeaircraftwingspanclasstype *CodeAircraftWingspanClassType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftWingspanClassTypeDeleteCallback != nil {
		var frontCodeAircraftWingspanClassType *CodeAircraftWingspanClassType
		if front != nil {
			frontCodeAircraftWingspanClassType, _ = front.(*CodeAircraftWingspanClassType)
		}
		stage.OnAfterCodeAircraftWingspanClassTypeDeleteCallback.OnAfterDelete(stage, codeaircraftwingspanclasstype, frontCodeAircraftWingspanClassType)
	}
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAltitudeAdjustmentTypeCreateCallback != nil {
		stage.OnAfterCodeAltitudeAdjustmentTypeCreateCallback.OnAfterCreate(stage, codealtitudeadjustmenttype)
	}
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAltitudeAdjustmentTypeUpdateCallback != nil {
		var frontCodeAltitudeAdjustmentType *CodeAltitudeAdjustmentType
		if front != nil {
			frontCodeAltitudeAdjustmentType, _ = front.(*CodeAltitudeAdjustmentType)
		}
		stage.OnAfterCodeAltitudeAdjustmentTypeUpdateCallback.OnAfterUpdate(stage, codealtitudeadjustmenttype, frontCodeAltitudeAdjustmentType)
	}
}

func (codealtitudeadjustmenttype *CodeAltitudeAdjustmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAltitudeAdjustmentTypeDeleteCallback != nil {
		var frontCodeAltitudeAdjustmentType *CodeAltitudeAdjustmentType
		if front != nil {
			frontCodeAltitudeAdjustmentType, _ = front.(*CodeAltitudeAdjustmentType)
		}
		stage.OnAfterCodeAltitudeAdjustmentTypeDeleteCallback.OnAfterDelete(stage, codealtitudeadjustmenttype, frontCodeAltitudeAdjustmentType)
	}
}

func (codealtitudeusetype *CodeAltitudeUseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAltitudeUseTypeCreateCallback != nil {
		stage.OnAfterCodeAltitudeUseTypeCreateCallback.OnAfterCreate(stage, codealtitudeusetype)
	}
}

func (codealtitudeusetype *CodeAltitudeUseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAltitudeUseTypeUpdateCallback != nil {
		var frontCodeAltitudeUseType *CodeAltitudeUseType
		if front != nil {
			frontCodeAltitudeUseType, _ = front.(*CodeAltitudeUseType)
		}
		stage.OnAfterCodeAltitudeUseTypeUpdateCallback.OnAfterUpdate(stage, codealtitudeusetype, frontCodeAltitudeUseType)
	}
}

func (codealtitudeusetype *CodeAltitudeUseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAltitudeUseTypeDeleteCallback != nil {
		var frontCodeAltitudeUseType *CodeAltitudeUseType
		if front != nil {
			frontCodeAltitudeUseType, _ = front.(*CodeAltitudeUseType)
		}
		stage.OnAfterCodeAltitudeUseTypeDeleteCallback.OnAfterDelete(stage, codealtitudeusetype, frontCodeAltitudeUseType)
	}
}

func (codeauthoritytype *CodeAuthorityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAuthorityTypeCreateCallback != nil {
		stage.OnAfterCodeAuthorityTypeCreateCallback.OnAfterCreate(stage, codeauthoritytype)
	}
}

func (codeauthoritytype *CodeAuthorityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuthorityTypeUpdateCallback != nil {
		var frontCodeAuthorityType *CodeAuthorityType
		if front != nil {
			frontCodeAuthorityType, _ = front.(*CodeAuthorityType)
		}
		stage.OnAfterCodeAuthorityTypeUpdateCallback.OnAfterUpdate(stage, codeauthoritytype, frontCodeAuthorityType)
	}
}

func (codeauthoritytype *CodeAuthorityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuthorityTypeDeleteCallback != nil {
		var frontCodeAuthorityType *CodeAuthorityType
		if front != nil {
			frontCodeAuthorityType, _ = front.(*CodeAuthorityType)
		}
		stage.OnAfterCodeAuthorityTypeDeleteCallback.OnAfterDelete(stage, codeauthoritytype, frontCodeAuthorityType)
	}
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCardinalDirectionTypeCreateCallback != nil {
		stage.OnAfterCodeCardinalDirectionTypeCreateCallback.OnAfterCreate(stage, codecardinaldirectiontype)
	}
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCardinalDirectionTypeUpdateCallback != nil {
		var frontCodeCardinalDirectionType *CodeCardinalDirectionType
		if front != nil {
			frontCodeCardinalDirectionType, _ = front.(*CodeCardinalDirectionType)
		}
		stage.OnAfterCodeCardinalDirectionTypeUpdateCallback.OnAfterUpdate(stage, codecardinaldirectiontype, frontCodeCardinalDirectionType)
	}
}

func (codecardinaldirectiontype *CodeCardinalDirectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCardinalDirectionTypeDeleteCallback != nil {
		var frontCodeCardinalDirectionType *CodeCardinalDirectionType
		if front != nil {
			frontCodeCardinalDirectionType, _ = front.(*CodeCardinalDirectionType)
		}
		stage.OnAfterCodeCardinalDirectionTypeDeleteCallback.OnAfterDelete(stage, codecardinaldirectiontype, frontCodeCardinalDirectionType)
	}
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCommunicationModeTypeCreateCallback != nil {
		stage.OnAfterCodeCommunicationModeTypeCreateCallback.OnAfterCreate(stage, codecommunicationmodetype)
	}
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationModeTypeUpdateCallback != nil {
		var frontCodeCommunicationModeType *CodeCommunicationModeType
		if front != nil {
			frontCodeCommunicationModeType, _ = front.(*CodeCommunicationModeType)
		}
		stage.OnAfterCodeCommunicationModeTypeUpdateCallback.OnAfterUpdate(stage, codecommunicationmodetype, frontCodeCommunicationModeType)
	}
}

func (codecommunicationmodetype *CodeCommunicationModeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationModeTypeDeleteCallback != nil {
		var frontCodeCommunicationModeType *CodeCommunicationModeType
		if front != nil {
			frontCodeCommunicationModeType, _ = front.(*CodeCommunicationModeType)
		}
		stage.OnAfterCodeCommunicationModeTypeDeleteCallback.OnAfterDelete(stage, codecommunicationmodetype, frontCodeCommunicationModeType)
	}
}

func (codecomparisontype *CodeComparisonType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeComparisonTypeCreateCallback != nil {
		stage.OnAfterCodeComparisonTypeCreateCallback.OnAfterCreate(stage, codecomparisontype)
	}
}

func (codecomparisontype *CodeComparisonType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeComparisonTypeUpdateCallback != nil {
		var frontCodeComparisonType *CodeComparisonType
		if front != nil {
			frontCodeComparisonType, _ = front.(*CodeComparisonType)
		}
		stage.OnAfterCodeComparisonTypeUpdateCallback.OnAfterUpdate(stage, codecomparisontype, frontCodeComparisonType)
	}
}

func (codecomparisontype *CodeComparisonType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeComparisonTypeDeleteCallback != nil {
		var frontCodeComparisonType *CodeComparisonType
		if front != nil {
			frontCodeComparisonType, _ = front.(*CodeComparisonType)
		}
		stage.OnAfterCodeComparisonTypeDeleteCallback.OnAfterDelete(stage, codecomparisontype, frontCodeComparisonType)
	}
}

func (codecoursetype *CodeCourseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCourseTypeCreateCallback != nil {
		stage.OnAfterCodeCourseTypeCreateCallback.OnAfterCreate(stage, codecoursetype)
	}
}

func (codecoursetype *CodeCourseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCourseTypeUpdateCallback != nil {
		var frontCodeCourseType *CodeCourseType
		if front != nil {
			frontCodeCourseType, _ = front.(*CodeCourseType)
		}
		stage.OnAfterCodeCourseTypeUpdateCallback.OnAfterUpdate(stage, codecoursetype, frontCodeCourseType)
	}
}

func (codecoursetype *CodeCourseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCourseTypeDeleteCallback != nil {
		var frontCodeCourseType *CodeCourseType
		if front != nil {
			frontCodeCourseType, _ = front.(*CodeCourseType)
		}
		stage.OnAfterCodeCourseTypeDeleteCallback.OnAfterDelete(stage, codecoursetype, frontCodeCourseType)
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

func (codedirectionturntype *CodeDirectionTurnType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDirectionTurnTypeCreateCallback != nil {
		stage.OnAfterCodeDirectionTurnTypeCreateCallback.OnAfterCreate(stage, codedirectionturntype)
	}
}

func (codedirectionturntype *CodeDirectionTurnType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDirectionTurnTypeUpdateCallback != nil {
		var frontCodeDirectionTurnType *CodeDirectionTurnType
		if front != nil {
			frontCodeDirectionTurnType, _ = front.(*CodeDirectionTurnType)
		}
		stage.OnAfterCodeDirectionTurnTypeUpdateCallback.OnAfterUpdate(stage, codedirectionturntype, frontCodeDirectionTurnType)
	}
}

func (codedirectionturntype *CodeDirectionTurnType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDirectionTurnTypeDeleteCallback != nil {
		var frontCodeDirectionTurnType *CodeDirectionTurnType
		if front != nil {
			frontCodeDirectionTurnType, _ = front.(*CodeDirectionTurnType)
		}
		stage.OnAfterCodeDirectionTurnTypeDeleteCallback.OnAfterDelete(stage, codedirectionturntype, frontCodeDirectionTurnType)
	}
}

func (codedirectiontype *CodeDirectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDirectionTypeCreateCallback != nil {
		stage.OnAfterCodeDirectionTypeCreateCallback.OnAfterCreate(stage, codedirectiontype)
	}
}

func (codedirectiontype *CodeDirectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDirectionTypeUpdateCallback != nil {
		var frontCodeDirectionType *CodeDirectionType
		if front != nil {
			frontCodeDirectionType, _ = front.(*CodeDirectionType)
		}
		stage.OnAfterCodeDirectionTypeUpdateCallback.OnAfterUpdate(stage, codedirectiontype, frontCodeDirectionType)
	}
}

func (codedirectiontype *CodeDirectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDirectionTypeDeleteCallback != nil {
		var frontCodeDirectionType *CodeDirectionType
		if front != nil {
			frontCodeDirectionType, _ = front.(*CodeDirectionType)
		}
		stage.OnAfterCodeDirectionTypeDeleteCallback.OnAfterDelete(stage, codedirectiontype, frontCodeDirectionType)
	}
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeEquipmentAntiCollisionTypeCreateCallback != nil {
		stage.OnAfterCodeEquipmentAntiCollisionTypeCreateCallback.OnAfterCreate(stage, codeequipmentanticollisiontype)
	}
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeEquipmentAntiCollisionTypeUpdateCallback != nil {
		var frontCodeEquipmentAntiCollisionType *CodeEquipmentAntiCollisionType
		if front != nil {
			frontCodeEquipmentAntiCollisionType, _ = front.(*CodeEquipmentAntiCollisionType)
		}
		stage.OnAfterCodeEquipmentAntiCollisionTypeUpdateCallback.OnAfterUpdate(stage, codeequipmentanticollisiontype, frontCodeEquipmentAntiCollisionType)
	}
}

func (codeequipmentanticollisiontype *CodeEquipmentAntiCollisionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeEquipmentAntiCollisionTypeDeleteCallback != nil {
		var frontCodeEquipmentAntiCollisionType *CodeEquipmentAntiCollisionType
		if front != nil {
			frontCodeEquipmentAntiCollisionType, _ = front.(*CodeEquipmentAntiCollisionType)
		}
		stage.OnAfterCodeEquipmentAntiCollisionTypeDeleteCallback.OnAfterDelete(stage, codeequipmentanticollisiontype, frontCodeEquipmentAntiCollisionType)
	}
}

func (codeflightorigintype *CodeFlightOriginType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightOriginTypeCreateCallback != nil {
		stage.OnAfterCodeFlightOriginTypeCreateCallback.OnAfterCreate(stage, codeflightorigintype)
	}
}

func (codeflightorigintype *CodeFlightOriginType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightOriginTypeUpdateCallback != nil {
		var frontCodeFlightOriginType *CodeFlightOriginType
		if front != nil {
			frontCodeFlightOriginType, _ = front.(*CodeFlightOriginType)
		}
		stage.OnAfterCodeFlightOriginTypeUpdateCallback.OnAfterUpdate(stage, codeflightorigintype, frontCodeFlightOriginType)
	}
}

func (codeflightorigintype *CodeFlightOriginType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightOriginTypeDeleteCallback != nil {
		var frontCodeFlightOriginType *CodeFlightOriginType
		if front != nil {
			frontCodeFlightOriginType, _ = front.(*CodeFlightOriginType)
		}
		stage.OnAfterCodeFlightOriginTypeDeleteCallback.OnAfterDelete(stage, codeflightorigintype, frontCodeFlightOriginType)
	}
}

func (codeflightpurposetype *CodeFlightPurposeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightPurposeTypeCreateCallback != nil {
		stage.OnAfterCodeFlightPurposeTypeCreateCallback.OnAfterCreate(stage, codeflightpurposetype)
	}
}

func (codeflightpurposetype *CodeFlightPurposeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightPurposeTypeUpdateCallback != nil {
		var frontCodeFlightPurposeType *CodeFlightPurposeType
		if front != nil {
			frontCodeFlightPurposeType, _ = front.(*CodeFlightPurposeType)
		}
		stage.OnAfterCodeFlightPurposeTypeUpdateCallback.OnAfterUpdate(stage, codeflightpurposetype, frontCodeFlightPurposeType)
	}
}

func (codeflightpurposetype *CodeFlightPurposeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightPurposeTypeDeleteCallback != nil {
		var frontCodeFlightPurposeType *CodeFlightPurposeType
		if front != nil {
			frontCodeFlightPurposeType, _ = front.(*CodeFlightPurposeType)
		}
		stage.OnAfterCodeFlightPurposeTypeDeleteCallback.OnAfterDelete(stage, codeflightpurposetype, frontCodeFlightPurposeType)
	}
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightRestrictionDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeFlightRestrictionDesignatorTypeCreateCallback.OnAfterCreate(stage, codeflightrestrictiondesignatortype)
	}
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightRestrictionDesignatorTypeUpdateCallback != nil {
		var frontCodeFlightRestrictionDesignatorType *CodeFlightRestrictionDesignatorType
		if front != nil {
			frontCodeFlightRestrictionDesignatorType, _ = front.(*CodeFlightRestrictionDesignatorType)
		}
		stage.OnAfterCodeFlightRestrictionDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codeflightrestrictiondesignatortype, frontCodeFlightRestrictionDesignatorType)
	}
}

func (codeflightrestrictiondesignatortype *CodeFlightRestrictionDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightRestrictionDesignatorTypeDeleteCallback != nil {
		var frontCodeFlightRestrictionDesignatorType *CodeFlightRestrictionDesignatorType
		if front != nil {
			frontCodeFlightRestrictionDesignatorType, _ = front.(*CodeFlightRestrictionDesignatorType)
		}
		stage.OnAfterCodeFlightRestrictionDesignatorTypeDeleteCallback.OnAfterDelete(stage, codeflightrestrictiondesignatortype, frontCodeFlightRestrictionDesignatorType)
	}
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightRestrictionTypeCreateCallback != nil {
		stage.OnAfterCodeFlightRestrictionTypeCreateCallback.OnAfterCreate(stage, codeflightrestrictiontype)
	}
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightRestrictionTypeUpdateCallback != nil {
		var frontCodeFlightRestrictionType *CodeFlightRestrictionType
		if front != nil {
			frontCodeFlightRestrictionType, _ = front.(*CodeFlightRestrictionType)
		}
		stage.OnAfterCodeFlightRestrictionTypeUpdateCallback.OnAfterUpdate(stage, codeflightrestrictiontype, frontCodeFlightRestrictionType)
	}
}

func (codeflightrestrictiontype *CodeFlightRestrictionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightRestrictionTypeDeleteCallback != nil {
		var frontCodeFlightRestrictionType *CodeFlightRestrictionType
		if front != nil {
			frontCodeFlightRestrictionType, _ = front.(*CodeFlightRestrictionType)
		}
		stage.OnAfterCodeFlightRestrictionTypeDeleteCallback.OnAfterDelete(stage, codeflightrestrictiontype, frontCodeFlightRestrictionType)
	}
}

func (codeflightruletype *CodeFlightRuleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightRuleTypeCreateCallback != nil {
		stage.OnAfterCodeFlightRuleTypeCreateCallback.OnAfterCreate(stage, codeflightruletype)
	}
}

func (codeflightruletype *CodeFlightRuleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightRuleTypeUpdateCallback != nil {
		var frontCodeFlightRuleType *CodeFlightRuleType
		if front != nil {
			frontCodeFlightRuleType, _ = front.(*CodeFlightRuleType)
		}
		stage.OnAfterCodeFlightRuleTypeUpdateCallback.OnAfterUpdate(stage, codeflightruletype, frontCodeFlightRuleType)
	}
}

func (codeflightruletype *CodeFlightRuleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightRuleTypeDeleteCallback != nil {
		var frontCodeFlightRuleType *CodeFlightRuleType
		if front != nil {
			frontCodeFlightRuleType, _ = front.(*CodeFlightRuleType)
		}
		stage.OnAfterCodeFlightRuleTypeDeleteCallback.OnAfterDelete(stage, codeflightruletype, frontCodeFlightRuleType)
	}
}

func (codeflightstatustype *CodeFlightStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightStatusTypeCreateCallback != nil {
		stage.OnAfterCodeFlightStatusTypeCreateCallback.OnAfterCreate(stage, codeflightstatustype)
	}
}

func (codeflightstatustype *CodeFlightStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightStatusTypeUpdateCallback != nil {
		var frontCodeFlightStatusType *CodeFlightStatusType
		if front != nil {
			frontCodeFlightStatusType, _ = front.(*CodeFlightStatusType)
		}
		stage.OnAfterCodeFlightStatusTypeUpdateCallback.OnAfterUpdate(stage, codeflightstatustype, frontCodeFlightStatusType)
	}
}

func (codeflightstatustype *CodeFlightStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightStatusTypeDeleteCallback != nil {
		var frontCodeFlightStatusType *CodeFlightStatusType
		if front != nil {
			frontCodeFlightStatusType, _ = front.(*CodeFlightStatusType)
		}
		stage.OnAfterCodeFlightStatusTypeDeleteCallback.OnAfterDelete(stage, codeflightstatustype, frontCodeFlightStatusType)
	}
}

func (codeflighttype *CodeFlightType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightTypeCreateCallback != nil {
		stage.OnAfterCodeFlightTypeCreateCallback.OnAfterCreate(stage, codeflighttype)
	}
}

func (codeflighttype *CodeFlightType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightTypeUpdateCallback != nil {
		var frontCodeFlightType *CodeFlightType
		if front != nil {
			frontCodeFlightType, _ = front.(*CodeFlightType)
		}
		stage.OnAfterCodeFlightTypeUpdateCallback.OnAfterUpdate(stage, codeflighttype, frontCodeFlightType)
	}
}

func (codeflighttype *CodeFlightType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightTypeDeleteCallback != nil {
		var frontCodeFlightType *CodeFlightType
		if front != nil {
			frontCodeFlightType, _ = front.(*CodeFlightType)
		}
		stage.OnAfterCodeFlightTypeDeleteCallback.OnAfterDelete(stage, codeflighttype, frontCodeFlightType)
	}
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlowConditionOperationTypeCreateCallback != nil {
		stage.OnAfterCodeFlowConditionOperationTypeCreateCallback.OnAfterCreate(stage, codeflowconditionoperationtype)
	}
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlowConditionOperationTypeUpdateCallback != nil {
		var frontCodeFlowConditionOperationType *CodeFlowConditionOperationType
		if front != nil {
			frontCodeFlowConditionOperationType, _ = front.(*CodeFlowConditionOperationType)
		}
		stage.OnAfterCodeFlowConditionOperationTypeUpdateCallback.OnAfterUpdate(stage, codeflowconditionoperationtype, frontCodeFlowConditionOperationType)
	}
}

func (codeflowconditionoperationtype *CodeFlowConditionOperationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlowConditionOperationTypeDeleteCallback != nil {
		var frontCodeFlowConditionOperationType *CodeFlowConditionOperationType
		if front != nil {
			frontCodeFlowConditionOperationType, _ = front.(*CodeFlowConditionOperationType)
		}
		stage.OnAfterCodeFlowConditionOperationTypeDeleteCallback.OnAfterDelete(stage, codeflowconditionoperationtype, frontCodeFlowConditionOperationType)
	}
}

func (codefreeflighttype *CodeFreeFlightType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFreeFlightTypeCreateCallback != nil {
		stage.OnAfterCodeFreeFlightTypeCreateCallback.OnAfterCreate(stage, codefreeflighttype)
	}
}

func (codefreeflighttype *CodeFreeFlightType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFreeFlightTypeUpdateCallback != nil {
		var frontCodeFreeFlightType *CodeFreeFlightType
		if front != nil {
			frontCodeFreeFlightType, _ = front.(*CodeFreeFlightType)
		}
		stage.OnAfterCodeFreeFlightTypeUpdateCallback.OnAfterUpdate(stage, codefreeflighttype, frontCodeFreeFlightType)
	}
}

func (codefreeflighttype *CodeFreeFlightType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFreeFlightTypeDeleteCallback != nil {
		var frontCodeFreeFlightType *CodeFreeFlightType
		if front != nil {
			frontCodeFreeFlightType, _ = front.(*CodeFreeFlightType)
		}
		stage.OnAfterCodeFreeFlightTypeDeleteCallback.OnAfterDelete(stage, codefreeflighttype, frontCodeFreeFlightType)
	}
}

func (codeleveltype *CodeLevelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLevelTypeCreateCallback != nil {
		stage.OnAfterCodeLevelTypeCreateCallback.OnAfterCreate(stage, codeleveltype)
	}
}

func (codeleveltype *CodeLevelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelTypeUpdateCallback != nil {
		var frontCodeLevelType *CodeLevelType
		if front != nil {
			frontCodeLevelType, _ = front.(*CodeLevelType)
		}
		stage.OnAfterCodeLevelTypeUpdateCallback.OnAfterUpdate(stage, codeleveltype, frontCodeLevelType)
	}
}

func (codeleveltype *CodeLevelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelTypeDeleteCallback != nil {
		var frontCodeLevelType *CodeLevelType
		if front != nil {
			frontCodeLevelType, _ = front.(*CodeLevelType)
		}
		stage.OnAfterCodeLevelTypeDeleteCallback.OnAfterDelete(stage, codeleveltype, frontCodeLevelType)
	}
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLocationQualifierTypeCreateCallback != nil {
		stage.OnAfterCodeLocationQualifierTypeCreateCallback.OnAfterCreate(stage, codelocationqualifiertype)
	}
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLocationQualifierTypeUpdateCallback != nil {
		var frontCodeLocationQualifierType *CodeLocationQualifierType
		if front != nil {
			frontCodeLocationQualifierType, _ = front.(*CodeLocationQualifierType)
		}
		stage.OnAfterCodeLocationQualifierTypeUpdateCallback.OnAfterUpdate(stage, codelocationqualifiertype, frontCodeLocationQualifierType)
	}
}

func (codelocationqualifiertype *CodeLocationQualifierType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLocationQualifierTypeDeleteCallback != nil {
		var frontCodeLocationQualifierType *CodeLocationQualifierType
		if front != nil {
			frontCodeLocationQualifierType, _ = front.(*CodeLocationQualifierType)
		}
		stage.OnAfterCodeLocationQualifierTypeDeleteCallback.OnAfterDelete(stage, codelocationqualifiertype, frontCodeLocationQualifierType)
	}
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMeteoConditionsTypeCreateCallback != nil {
		stage.OnAfterCodeMeteoConditionsTypeCreateCallback.OnAfterCreate(stage, codemeteoconditionstype)
	}
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMeteoConditionsTypeUpdateCallback != nil {
		var frontCodeMeteoConditionsType *CodeMeteoConditionsType
		if front != nil {
			frontCodeMeteoConditionsType, _ = front.(*CodeMeteoConditionsType)
		}
		stage.OnAfterCodeMeteoConditionsTypeUpdateCallback.OnAfterUpdate(stage, codemeteoconditionstype, frontCodeMeteoConditionsType)
	}
}

func (codemeteoconditionstype *CodeMeteoConditionsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMeteoConditionsTypeDeleteCallback != nil {
		var frontCodeMeteoConditionsType *CodeMeteoConditionsType
		if front != nil {
			frontCodeMeteoConditionsType, _ = front.(*CodeMeteoConditionsType)
		}
		stage.OnAfterCodeMeteoConditionsTypeDeleteCallback.OnAfterDelete(stage, codemeteoconditionstype, frontCodeMeteoConditionsType)
	}
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMilitaryRoutePointTypeCreateCallback != nil {
		stage.OnAfterCodeMilitaryRoutePointTypeCreateCallback.OnAfterCreate(stage, codemilitaryroutepointtype)
	}
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryRoutePointTypeUpdateCallback != nil {
		var frontCodeMilitaryRoutePointType *CodeMilitaryRoutePointType
		if front != nil {
			frontCodeMilitaryRoutePointType, _ = front.(*CodeMilitaryRoutePointType)
		}
		stage.OnAfterCodeMilitaryRoutePointTypeUpdateCallback.OnAfterUpdate(stage, codemilitaryroutepointtype, frontCodeMilitaryRoutePointType)
	}
}

func (codemilitaryroutepointtype *CodeMilitaryRoutePointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryRoutePointTypeDeleteCallback != nil {
		var frontCodeMilitaryRoutePointType *CodeMilitaryRoutePointType
		if front != nil {
			frontCodeMilitaryRoutePointType, _ = front.(*CodeMilitaryRoutePointType)
		}
		stage.OnAfterCodeMilitaryRoutePointTypeDeleteCallback.OnAfterDelete(stage, codemilitaryroutepointtype, frontCodeMilitaryRoutePointType)
	}
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMilitaryStatusTypeCreateCallback != nil {
		stage.OnAfterCodeMilitaryStatusTypeCreateCallback.OnAfterCreate(stage, codemilitarystatustype)
	}
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryStatusTypeUpdateCallback != nil {
		var frontCodeMilitaryStatusType *CodeMilitaryStatusType
		if front != nil {
			frontCodeMilitaryStatusType, _ = front.(*CodeMilitaryStatusType)
		}
		stage.OnAfterCodeMilitaryStatusTypeUpdateCallback.OnAfterUpdate(stage, codemilitarystatustype, frontCodeMilitaryStatusType)
	}
}

func (codemilitarystatustype *CodeMilitaryStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryStatusTypeDeleteCallback != nil {
		var frontCodeMilitaryStatusType *CodeMilitaryStatusType
		if front != nil {
			frontCodeMilitaryStatusType, _ = front.(*CodeMilitaryStatusType)
		}
		stage.OnAfterCodeMilitaryStatusTypeDeleteCallback.OnAfterDelete(stage, codemilitarystatustype, frontCodeMilitaryStatusType)
	}
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMilitaryTrainingTypeCreateCallback != nil {
		stage.OnAfterCodeMilitaryTrainingTypeCreateCallback.OnAfterCreate(stage, codemilitarytrainingtype)
	}
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryTrainingTypeUpdateCallback != nil {
		var frontCodeMilitaryTrainingType *CodeMilitaryTrainingType
		if front != nil {
			frontCodeMilitaryTrainingType, _ = front.(*CodeMilitaryTrainingType)
		}
		stage.OnAfterCodeMilitaryTrainingTypeUpdateCallback.OnAfterUpdate(stage, codemilitarytrainingtype, frontCodeMilitaryTrainingType)
	}
}

func (codemilitarytrainingtype *CodeMilitaryTrainingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryTrainingTypeDeleteCallback != nil {
		var frontCodeMilitaryTrainingType *CodeMilitaryTrainingType
		if front != nil {
			frontCodeMilitaryTrainingType, _ = front.(*CodeMilitaryTrainingType)
		}
		stage.OnAfterCodeMilitaryTrainingTypeDeleteCallback.OnAfterDelete(stage, codemilitarytrainingtype, frontCodeMilitaryTrainingType)
	}
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavigationEquipmentTypeCreateCallback != nil {
		stage.OnAfterCodeNavigationEquipmentTypeCreateCallback.OnAfterCreate(stage, codenavigationequipmenttype)
	}
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavigationEquipmentTypeUpdateCallback != nil {
		var frontCodeNavigationEquipmentType *CodeNavigationEquipmentType
		if front != nil {
			frontCodeNavigationEquipmentType, _ = front.(*CodeNavigationEquipmentType)
		}
		stage.OnAfterCodeNavigationEquipmentTypeUpdateCallback.OnAfterUpdate(stage, codenavigationequipmenttype, frontCodeNavigationEquipmentType)
	}
}

func (codenavigationequipmenttype *CodeNavigationEquipmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavigationEquipmentTypeDeleteCallback != nil {
		var frontCodeNavigationEquipmentType *CodeNavigationEquipmentType
		if front != nil {
			frontCodeNavigationEquipmentType, _ = front.(*CodeNavigationEquipmentType)
		}
		stage.OnAfterCodeNavigationEquipmentTypeDeleteCallback.OnAfterDelete(stage, codenavigationequipmenttype, frontCodeNavigationEquipmentType)
	}
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavigationSpecificationTypeCreateCallback != nil {
		stage.OnAfterCodeNavigationSpecificationTypeCreateCallback.OnAfterCreate(stage, codenavigationspecificationtype)
	}
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavigationSpecificationTypeUpdateCallback != nil {
		var frontCodeNavigationSpecificationType *CodeNavigationSpecificationType
		if front != nil {
			frontCodeNavigationSpecificationType, _ = front.(*CodeNavigationSpecificationType)
		}
		stage.OnAfterCodeNavigationSpecificationTypeUpdateCallback.OnAfterUpdate(stage, codenavigationspecificationtype, frontCodeNavigationSpecificationType)
	}
}

func (codenavigationspecificationtype *CodeNavigationSpecificationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavigationSpecificationTypeDeleteCallback != nil {
		var frontCodeNavigationSpecificationType *CodeNavigationSpecificationType
		if front != nil {
			frontCodeNavigationSpecificationType, _ = front.(*CodeNavigationSpecificationType)
		}
		stage.OnAfterCodeNavigationSpecificationTypeDeleteCallback.OnAfterDelete(stage, codenavigationspecificationtype, frontCodeNavigationSpecificationType)
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

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeObstructionIdSurfaceZoneTypeCreateCallback != nil {
		stage.OnAfterCodeObstructionIdSurfaceZoneTypeCreateCallback.OnAfterCreate(stage, codeobstructionidsurfacezonetype)
	}
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeObstructionIdSurfaceZoneTypeUpdateCallback != nil {
		var frontCodeObstructionIdSurfaceZoneType *CodeObstructionIdSurfaceZoneType
		if front != nil {
			frontCodeObstructionIdSurfaceZoneType, _ = front.(*CodeObstructionIdSurfaceZoneType)
		}
		stage.OnAfterCodeObstructionIdSurfaceZoneTypeUpdateCallback.OnAfterUpdate(stage, codeobstructionidsurfacezonetype, frontCodeObstructionIdSurfaceZoneType)
	}
}

func (codeobstructionidsurfacezonetype *CodeObstructionIdSurfaceZoneType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeObstructionIdSurfaceZoneTypeDeleteCallback != nil {
		var frontCodeObstructionIdSurfaceZoneType *CodeObstructionIdSurfaceZoneType
		if front != nil {
			frontCodeObstructionIdSurfaceZoneType, _ = front.(*CodeObstructionIdSurfaceZoneType)
		}
		stage.OnAfterCodeObstructionIdSurfaceZoneTypeDeleteCallback.OnAfterDelete(stage, codeobstructionidsurfacezonetype, frontCodeObstructionIdSurfaceZoneType)
	}
}

func (codernptype *CodeRNPType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRNPTypeCreateCallback != nil {
		stage.OnAfterCodeRNPTypeCreateCallback.OnAfterCreate(stage, codernptype)
	}
}

func (codernptype *CodeRNPType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRNPTypeUpdateCallback != nil {
		var frontCodeRNPType *CodeRNPType
		if front != nil {
			frontCodeRNPType, _ = front.(*CodeRNPType)
		}
		stage.OnAfterCodeRNPTypeUpdateCallback.OnAfterUpdate(stage, codernptype, frontCodeRNPType)
	}
}

func (codernptype *CodeRNPType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRNPTypeDeleteCallback != nil {
		var frontCodeRNPType *CodeRNPType
		if front != nil {
			frontCodeRNPType, _ = front.(*CodeRNPType)
		}
		stage.OnAfterCodeRNPTypeDeleteCallback.OnAfterDelete(stage, codernptype, frontCodeRNPType)
	}
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRVSMPointRoleTypeCreateCallback != nil {
		stage.OnAfterCodeRVSMPointRoleTypeCreateCallback.OnAfterCreate(stage, codervsmpointroletype)
	}
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRVSMPointRoleTypeUpdateCallback != nil {
		var frontCodeRVSMPointRoleType *CodeRVSMPointRoleType
		if front != nil {
			frontCodeRVSMPointRoleType, _ = front.(*CodeRVSMPointRoleType)
		}
		stage.OnAfterCodeRVSMPointRoleTypeUpdateCallback.OnAfterUpdate(stage, codervsmpointroletype, frontCodeRVSMPointRoleType)
	}
}

func (codervsmpointroletype *CodeRVSMPointRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRVSMPointRoleTypeDeleteCallback != nil {
		var frontCodeRVSMPointRoleType *CodeRVSMPointRoleType
		if front != nil {
			frontCodeRVSMPointRoleType, _ = front.(*CodeRVSMPointRoleType)
		}
		stage.OnAfterCodeRVSMPointRoleTypeDeleteCallback.OnAfterDelete(stage, codervsmpointroletype, frontCodeRVSMPointRoleType)
	}
}

func (codervsmtype *CodeRVSMType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRVSMTypeCreateCallback != nil {
		stage.OnAfterCodeRVSMTypeCreateCallback.OnAfterCreate(stage, codervsmtype)
	}
}

func (codervsmtype *CodeRVSMType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRVSMTypeUpdateCallback != nil {
		var frontCodeRVSMType *CodeRVSMType
		if front != nil {
			frontCodeRVSMType, _ = front.(*CodeRVSMType)
		}
		stage.OnAfterCodeRVSMTypeUpdateCallback.OnAfterUpdate(stage, codervsmtype, frontCodeRVSMType)
	}
}

func (codervsmtype *CodeRVSMType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRVSMTypeDeleteCallback != nil {
		var frontCodeRVSMType *CodeRVSMType
		if front != nil {
			frontCodeRVSMType, _ = front.(*CodeRVSMType)
		}
		stage.OnAfterCodeRVSMTypeDeleteCallback.OnAfterDelete(stage, codervsmtype, frontCodeRVSMType)
	}
}

func (codereferenceroletype *CodeReferenceRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeReferenceRoleTypeCreateCallback != nil {
		stage.OnAfterCodeReferenceRoleTypeCreateCallback.OnAfterCreate(stage, codereferenceroletype)
	}
}

func (codereferenceroletype *CodeReferenceRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeReferenceRoleTypeUpdateCallback != nil {
		var frontCodeReferenceRoleType *CodeReferenceRoleType
		if front != nil {
			frontCodeReferenceRoleType, _ = front.(*CodeReferenceRoleType)
		}
		stage.OnAfterCodeReferenceRoleTypeUpdateCallback.OnAfterUpdate(stage, codereferenceroletype, frontCodeReferenceRoleType)
	}
}

func (codereferenceroletype *CodeReferenceRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeReferenceRoleTypeDeleteCallback != nil {
		var frontCodeReferenceRoleType *CodeReferenceRoleType
		if front != nil {
			frontCodeReferenceRoleType, _ = front.(*CodeReferenceRoleType)
		}
		stage.OnAfterCodeReferenceRoleTypeDeleteCallback.OnAfterDelete(stage, codereferenceroletype, frontCodeReferenceRoleType)
	}
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteAvailabilityTypeCreateCallback != nil {
		stage.OnAfterCodeRouteAvailabilityTypeCreateCallback.OnAfterCreate(stage, coderouteavailabilitytype)
	}
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteAvailabilityTypeUpdateCallback != nil {
		var frontCodeRouteAvailabilityType *CodeRouteAvailabilityType
		if front != nil {
			frontCodeRouteAvailabilityType, _ = front.(*CodeRouteAvailabilityType)
		}
		stage.OnAfterCodeRouteAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, coderouteavailabilitytype, frontCodeRouteAvailabilityType)
	}
}

func (coderouteavailabilitytype *CodeRouteAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteAvailabilityTypeDeleteCallback != nil {
		var frontCodeRouteAvailabilityType *CodeRouteAvailabilityType
		if front != nil {
			frontCodeRouteAvailabilityType, _ = front.(*CodeRouteAvailabilityType)
		}
		stage.OnAfterCodeRouteAvailabilityTypeDeleteCallback.OnAfterDelete(stage, coderouteavailabilitytype, frontCodeRouteAvailabilityType)
	}
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteDesignatorLetterTypeCreateCallback != nil {
		stage.OnAfterCodeRouteDesignatorLetterTypeCreateCallback.OnAfterCreate(stage, coderoutedesignatorlettertype)
	}
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteDesignatorLetterTypeUpdateCallback != nil {
		var frontCodeRouteDesignatorLetterType *CodeRouteDesignatorLetterType
		if front != nil {
			frontCodeRouteDesignatorLetterType, _ = front.(*CodeRouteDesignatorLetterType)
		}
		stage.OnAfterCodeRouteDesignatorLetterTypeUpdateCallback.OnAfterUpdate(stage, coderoutedesignatorlettertype, frontCodeRouteDesignatorLetterType)
	}
}

func (coderoutedesignatorlettertype *CodeRouteDesignatorLetterType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteDesignatorLetterTypeDeleteCallback != nil {
		var frontCodeRouteDesignatorLetterType *CodeRouteDesignatorLetterType
		if front != nil {
			frontCodeRouteDesignatorLetterType, _ = front.(*CodeRouteDesignatorLetterType)
		}
		stage.OnAfterCodeRouteDesignatorLetterTypeDeleteCallback.OnAfterDelete(stage, coderoutedesignatorlettertype, frontCodeRouteDesignatorLetterType)
	}
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteDesignatorPrefixTypeCreateCallback != nil {
		stage.OnAfterCodeRouteDesignatorPrefixTypeCreateCallback.OnAfterCreate(stage, coderoutedesignatorprefixtype)
	}
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteDesignatorPrefixTypeUpdateCallback != nil {
		var frontCodeRouteDesignatorPrefixType *CodeRouteDesignatorPrefixType
		if front != nil {
			frontCodeRouteDesignatorPrefixType, _ = front.(*CodeRouteDesignatorPrefixType)
		}
		stage.OnAfterCodeRouteDesignatorPrefixTypeUpdateCallback.OnAfterUpdate(stage, coderoutedesignatorprefixtype, frontCodeRouteDesignatorPrefixType)
	}
}

func (coderoutedesignatorprefixtype *CodeRouteDesignatorPrefixType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteDesignatorPrefixTypeDeleteCallback != nil {
		var frontCodeRouteDesignatorPrefixType *CodeRouteDesignatorPrefixType
		if front != nil {
			frontCodeRouteDesignatorPrefixType, _ = front.(*CodeRouteDesignatorPrefixType)
		}
		stage.OnAfterCodeRouteDesignatorPrefixTypeDeleteCallback.OnAfterDelete(stage, coderoutedesignatorprefixtype, frontCodeRouteDesignatorPrefixType)
	}
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteDesignatorSuffixTypeCreateCallback != nil {
		stage.OnAfterCodeRouteDesignatorSuffixTypeCreateCallback.OnAfterCreate(stage, coderoutedesignatorsuffixtype)
	}
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteDesignatorSuffixTypeUpdateCallback != nil {
		var frontCodeRouteDesignatorSuffixType *CodeRouteDesignatorSuffixType
		if front != nil {
			frontCodeRouteDesignatorSuffixType, _ = front.(*CodeRouteDesignatorSuffixType)
		}
		stage.OnAfterCodeRouteDesignatorSuffixTypeUpdateCallback.OnAfterUpdate(stage, coderoutedesignatorsuffixtype, frontCodeRouteDesignatorSuffixType)
	}
}

func (coderoutedesignatorsuffixtype *CodeRouteDesignatorSuffixType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteDesignatorSuffixTypeDeleteCallback != nil {
		var frontCodeRouteDesignatorSuffixType *CodeRouteDesignatorSuffixType
		if front != nil {
			frontCodeRouteDesignatorSuffixType, _ = front.(*CodeRouteDesignatorSuffixType)
		}
		stage.OnAfterCodeRouteDesignatorSuffixTypeDeleteCallback.OnAfterDelete(stage, coderoutedesignatorsuffixtype, frontCodeRouteDesignatorSuffixType)
	}
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteNavigationTypeCreateCallback != nil {
		stage.OnAfterCodeRouteNavigationTypeCreateCallback.OnAfterCreate(stage, coderoutenavigationtype)
	}
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteNavigationTypeUpdateCallback != nil {
		var frontCodeRouteNavigationType *CodeRouteNavigationType
		if front != nil {
			frontCodeRouteNavigationType, _ = front.(*CodeRouteNavigationType)
		}
		stage.OnAfterCodeRouteNavigationTypeUpdateCallback.OnAfterUpdate(stage, coderoutenavigationtype, frontCodeRouteNavigationType)
	}
}

func (coderoutenavigationtype *CodeRouteNavigationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteNavigationTypeDeleteCallback != nil {
		var frontCodeRouteNavigationType *CodeRouteNavigationType
		if front != nil {
			frontCodeRouteNavigationType, _ = front.(*CodeRouteNavigationType)
		}
		stage.OnAfterCodeRouteNavigationTypeDeleteCallback.OnAfterDelete(stage, coderoutenavigationtype, frontCodeRouteNavigationType)
	}
}

func (coderouteorigintype *CodeRouteOriginType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteOriginTypeCreateCallback != nil {
		stage.OnAfterCodeRouteOriginTypeCreateCallback.OnAfterCreate(stage, coderouteorigintype)
	}
}

func (coderouteorigintype *CodeRouteOriginType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteOriginTypeUpdateCallback != nil {
		var frontCodeRouteOriginType *CodeRouteOriginType
		if front != nil {
			frontCodeRouteOriginType, _ = front.(*CodeRouteOriginType)
		}
		stage.OnAfterCodeRouteOriginTypeUpdateCallback.OnAfterUpdate(stage, coderouteorigintype, frontCodeRouteOriginType)
	}
}

func (coderouteorigintype *CodeRouteOriginType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteOriginTypeDeleteCallback != nil {
		var frontCodeRouteOriginType *CodeRouteOriginType
		if front != nil {
			frontCodeRouteOriginType, _ = front.(*CodeRouteOriginType)
		}
		stage.OnAfterCodeRouteOriginTypeDeleteCallback.OnAfterDelete(stage, coderouteorigintype, frontCodeRouteOriginType)
	}
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteSegmentPathTypeCreateCallback != nil {
		stage.OnAfterCodeRouteSegmentPathTypeCreateCallback.OnAfterCreate(stage, coderoutesegmentpathtype)
	}
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteSegmentPathTypeUpdateCallback != nil {
		var frontCodeRouteSegmentPathType *CodeRouteSegmentPathType
		if front != nil {
			frontCodeRouteSegmentPathType, _ = front.(*CodeRouteSegmentPathType)
		}
		stage.OnAfterCodeRouteSegmentPathTypeUpdateCallback.OnAfterUpdate(stage, coderoutesegmentpathtype, frontCodeRouteSegmentPathType)
	}
}

func (coderoutesegmentpathtype *CodeRouteSegmentPathType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteSegmentPathTypeDeleteCallback != nil {
		var frontCodeRouteSegmentPathType *CodeRouteSegmentPathType
		if front != nil {
			frontCodeRouteSegmentPathType, _ = front.(*CodeRouteSegmentPathType)
		}
		stage.OnAfterCodeRouteSegmentPathTypeDeleteCallback.OnAfterDelete(stage, coderoutesegmentpathtype, frontCodeRouteSegmentPathType)
	}
}

func (coderoutetype *CodeRouteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteTypeCreateCallback != nil {
		stage.OnAfterCodeRouteTypeCreateCallback.OnAfterCreate(stage, coderoutetype)
	}
}

func (coderoutetype *CodeRouteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteTypeUpdateCallback != nil {
		var frontCodeRouteType *CodeRouteType
		if front != nil {
			frontCodeRouteType, _ = front.(*CodeRouteType)
		}
		stage.OnAfterCodeRouteTypeUpdateCallback.OnAfterUpdate(stage, coderoutetype, frontCodeRouteType)
	}
}

func (coderoutetype *CodeRouteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteTypeDeleteCallback != nil {
		var frontCodeRouteType *CodeRouteType
		if front != nil {
			frontCodeRouteType, _ = front.(*CodeRouteType)
		}
		stage.OnAfterCodeRouteTypeDeleteCallback.OnAfterDelete(stage, coderoutetype, frontCodeRouteType)
	}
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRuleProcedureTitleTypeCreateCallback != nil {
		stage.OnAfterCodeRuleProcedureTitleTypeCreateCallback.OnAfterCreate(stage, coderuleproceduretitletype)
	}
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRuleProcedureTitleTypeUpdateCallback != nil {
		var frontCodeRuleProcedureTitleType *CodeRuleProcedureTitleType
		if front != nil {
			frontCodeRuleProcedureTitleType, _ = front.(*CodeRuleProcedureTitleType)
		}
		stage.OnAfterCodeRuleProcedureTitleTypeUpdateCallback.OnAfterUpdate(stage, coderuleproceduretitletype, frontCodeRuleProcedureTitleType)
	}
}

func (coderuleproceduretitletype *CodeRuleProcedureTitleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRuleProcedureTitleTypeDeleteCallback != nil {
		var frontCodeRuleProcedureTitleType *CodeRuleProcedureTitleType
		if front != nil {
			frontCodeRuleProcedureTitleType, _ = front.(*CodeRuleProcedureTitleType)
		}
		stage.OnAfterCodeRuleProcedureTitleTypeDeleteCallback.OnAfterDelete(stage, coderuleproceduretitletype, frontCodeRuleProcedureTitleType)
	}
}

func (coderuleproceduretype *CodeRuleProcedureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRuleProcedureTypeCreateCallback != nil {
		stage.OnAfterCodeRuleProcedureTypeCreateCallback.OnAfterCreate(stage, coderuleproceduretype)
	}
}

func (coderuleproceduretype *CodeRuleProcedureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRuleProcedureTypeUpdateCallback != nil {
		var frontCodeRuleProcedureType *CodeRuleProcedureType
		if front != nil {
			frontCodeRuleProcedureType, _ = front.(*CodeRuleProcedureType)
		}
		stage.OnAfterCodeRuleProcedureTypeUpdateCallback.OnAfterUpdate(stage, coderuleproceduretype, frontCodeRuleProcedureType)
	}
}

func (coderuleproceduretype *CodeRuleProcedureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRuleProcedureTypeDeleteCallback != nil {
		var frontCodeRuleProcedureType *CodeRuleProcedureType
		if front != nil {
			frontCodeRuleProcedureType, _ = front.(*CodeRuleProcedureType)
		}
		stage.OnAfterCodeRuleProcedureTypeDeleteCallback.OnAfterDelete(stage, coderuleproceduretype, frontCodeRuleProcedureType)
	}
}

func (codesidetype *CodeSideType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSideTypeCreateCallback != nil {
		stage.OnAfterCodeSideTypeCreateCallback.OnAfterCreate(stage, codesidetype)
	}
}

func (codesidetype *CodeSideType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSideTypeUpdateCallback != nil {
		var frontCodeSideType *CodeSideType
		if front != nil {
			frontCodeSideType, _ = front.(*CodeSideType)
		}
		stage.OnAfterCodeSideTypeUpdateCallback.OnAfterUpdate(stage, codesidetype, frontCodeSideType)
	}
}

func (codesidetype *CodeSideType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSideTypeDeleteCallback != nil {
		var frontCodeSideType *CodeSideType
		if front != nil {
			frontCodeSideType, _ = front.(*CodeSideType)
		}
		stage.OnAfterCodeSideTypeDeleteCallback.OnAfterDelete(stage, codesidetype, frontCodeSideType)
	}
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpeedReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeSpeedReferenceTypeCreateCallback.OnAfterCreate(stage, codespeedreferencetype)
	}
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpeedReferenceTypeUpdateCallback != nil {
		var frontCodeSpeedReferenceType *CodeSpeedReferenceType
		if front != nil {
			frontCodeSpeedReferenceType, _ = front.(*CodeSpeedReferenceType)
		}
		stage.OnAfterCodeSpeedReferenceTypeUpdateCallback.OnAfterUpdate(stage, codespeedreferencetype, frontCodeSpeedReferenceType)
	}
}

func (codespeedreferencetype *CodeSpeedReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpeedReferenceTypeDeleteCallback != nil {
		var frontCodeSpeedReferenceType *CodeSpeedReferenceType
		if front != nil {
			frontCodeSpeedReferenceType, _ = front.(*CodeSpeedReferenceType)
		}
		stage.OnAfterCodeSpeedReferenceTypeDeleteCallback.OnAfterDelete(stage, codespeedreferencetype, frontCodeSpeedReferenceType)
	}
}

func (codetacanchanneltype *CodeTACANChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTACANChannelTypeCreateCallback != nil {
		stage.OnAfterCodeTACANChannelTypeCreateCallback.OnAfterCreate(stage, codetacanchanneltype)
	}
}

func (codetacanchanneltype *CodeTACANChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTACANChannelTypeUpdateCallback != nil {
		var frontCodeTACANChannelType *CodeTACANChannelType
		if front != nil {
			frontCodeTACANChannelType, _ = front.(*CodeTACANChannelType)
		}
		stage.OnAfterCodeTACANChannelTypeUpdateCallback.OnAfterUpdate(stage, codetacanchanneltype, frontCodeTACANChannelType)
	}
}

func (codetacanchanneltype *CodeTACANChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTACANChannelTypeDeleteCallback != nil {
		var frontCodeTACANChannelType *CodeTACANChannelType
		if front != nil {
			frontCodeTACANChannelType, _ = front.(*CodeTACANChannelType)
		}
		stage.OnAfterCodeTACANChannelTypeDeleteCallback.OnAfterDelete(stage, codetacanchanneltype, frontCodeTACANChannelType)
	}
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTelecomNetworkTypeCreateCallback != nil {
		stage.OnAfterCodeTelecomNetworkTypeCreateCallback.OnAfterCreate(stage, codetelecomnetworktype)
	}
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTelecomNetworkTypeUpdateCallback != nil {
		var frontCodeTelecomNetworkType *CodeTelecomNetworkType
		if front != nil {
			frontCodeTelecomNetworkType, _ = front.(*CodeTelecomNetworkType)
		}
		stage.OnAfterCodeTelecomNetworkTypeUpdateCallback.OnAfterUpdate(stage, codetelecomnetworktype, frontCodeTelecomNetworkType)
	}
}

func (codetelecomnetworktype *CodeTelecomNetworkType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTelecomNetworkTypeDeleteCallback != nil {
		var frontCodeTelecomNetworkType *CodeTelecomNetworkType
		if front != nil {
			frontCodeTelecomNetworkType, _ = front.(*CodeTelecomNetworkType)
		}
		stage.OnAfterCodeTelecomNetworkTypeDeleteCallback.OnAfterDelete(stage, codetelecomnetworktype, frontCodeTelecomNetworkType)
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

func (codetranspondertype *CodeTransponderType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTransponderTypeCreateCallback != nil {
		stage.OnAfterCodeTransponderTypeCreateCallback.OnAfterCreate(stage, codetranspondertype)
	}
}

func (codetranspondertype *CodeTransponderType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTransponderTypeUpdateCallback != nil {
		var frontCodeTransponderType *CodeTransponderType
		if front != nil {
			frontCodeTransponderType, _ = front.(*CodeTransponderType)
		}
		stage.OnAfterCodeTransponderTypeUpdateCallback.OnAfterUpdate(stage, codetranspondertype, frontCodeTransponderType)
	}
}

func (codetranspondertype *CodeTransponderType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTransponderTypeDeleteCallback != nil {
		var frontCodeTransponderType *CodeTransponderType
		if front != nil {
			frontCodeTransponderType, _ = front.(*CodeTransponderType)
		}
		stage.OnAfterCodeTransponderTypeDeleteCallback.OnAfterDelete(stage, codetranspondertype, frontCodeTransponderType)
	}
}

func (codeupperalphatype *CodeUpperAlphaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeUpperAlphaTypeCreateCallback != nil {
		stage.OnAfterCodeUpperAlphaTypeCreateCallback.OnAfterCreate(stage, codeupperalphatype)
	}
}

func (codeupperalphatype *CodeUpperAlphaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUpperAlphaTypeUpdateCallback != nil {
		var frontCodeUpperAlphaType *CodeUpperAlphaType
		if front != nil {
			frontCodeUpperAlphaType, _ = front.(*CodeUpperAlphaType)
		}
		stage.OnAfterCodeUpperAlphaTypeUpdateCallback.OnAfterUpdate(stage, codeupperalphatype, frontCodeUpperAlphaType)
	}
}

func (codeupperalphatype *CodeUpperAlphaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUpperAlphaTypeDeleteCallback != nil {
		var frontCodeUpperAlphaType *CodeUpperAlphaType
		if front != nil {
			frontCodeUpperAlphaType, _ = front.(*CodeUpperAlphaType)
		}
		stage.OnAfterCodeUpperAlphaTypeDeleteCallback.OnAfterDelete(stage, codeupperalphatype, frontCodeUpperAlphaType)
	}
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeValueInterpretationTypeCreateCallback != nil {
		stage.OnAfterCodeValueInterpretationTypeCreateCallback.OnAfterCreate(stage, codevalueinterpretationtype)
	}
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeValueInterpretationTypeUpdateCallback != nil {
		var frontCodeValueInterpretationType *CodeValueInterpretationType
		if front != nil {
			frontCodeValueInterpretationType, _ = front.(*CodeValueInterpretationType)
		}
		stage.OnAfterCodeValueInterpretationTypeUpdateCallback.OnAfterUpdate(stage, codevalueinterpretationtype, frontCodeValueInterpretationType)
	}
}

func (codevalueinterpretationtype *CodeValueInterpretationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeValueInterpretationTypeDeleteCallback != nil {
		var frontCodeValueInterpretationType *CodeValueInterpretationType
		if front != nil {
			frontCodeValueInterpretationType, _ = front.(*CodeValueInterpretationType)
		}
		stage.OnAfterCodeValueInterpretationTypeDeleteCallback.OnAfterDelete(stage, codevalueinterpretationtype, frontCodeValueInterpretationType)
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

func (codeverticalreferencetype *CodeVerticalReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVerticalReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeVerticalReferenceTypeCreateCallback.OnAfterCreate(stage, codeverticalreferencetype)
	}
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalReferenceTypeUpdateCallback != nil {
		var frontCodeVerticalReferenceType *CodeVerticalReferenceType
		if front != nil {
			frontCodeVerticalReferenceType, _ = front.(*CodeVerticalReferenceType)
		}
		stage.OnAfterCodeVerticalReferenceTypeUpdateCallback.OnAfterUpdate(stage, codeverticalreferencetype, frontCodeVerticalReferenceType)
	}
}

func (codeverticalreferencetype *CodeVerticalReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVerticalReferenceTypeDeleteCallback != nil {
		var frontCodeVerticalReferenceType *CodeVerticalReferenceType
		if front != nil {
			frontCodeVerticalReferenceType, _ = front.(*CodeVerticalReferenceType)
		}
		stage.OnAfterCodeVerticalReferenceTypeDeleteCallback.OnAfterDelete(stage, codeverticalreferencetype, frontCodeVerticalReferenceType)
	}
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeWakeTurbulenceTypeCreateCallback != nil {
		stage.OnAfterCodeWakeTurbulenceTypeCreateCallback.OnAfterCreate(stage, codewaketurbulencetype)
	}
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeWakeTurbulenceTypeUpdateCallback != nil {
		var frontCodeWakeTurbulenceType *CodeWakeTurbulenceType
		if front != nil {
			frontCodeWakeTurbulenceType, _ = front.(*CodeWakeTurbulenceType)
		}
		stage.OnAfterCodeWakeTurbulenceTypeUpdateCallback.OnAfterUpdate(stage, codewaketurbulencetype, frontCodeWakeTurbulenceType)
	}
}

func (codewaketurbulencetype *CodeWakeTurbulenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeWakeTurbulenceTypeDeleteCallback != nil {
		var frontCodeWakeTurbulenceType *CodeWakeTurbulenceType
		if front != nil {
			frontCodeWakeTurbulenceType, _ = front.(*CodeWakeTurbulenceType)
		}
		stage.OnAfterCodeWakeTurbulenceTypeDeleteCallback.OnAfterDelete(stage, codewaketurbulencetype, frontCodeWakeTurbulenceType)
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

func (nonumbertype *NoNumberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoNumberTypeCreateCallback != nil {
		stage.OnAfterNoNumberTypeCreateCallback.OnAfterCreate(stage, nonumbertype)
	}
}

func (nonumbertype *NoNumberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoNumberTypeUpdateCallback != nil {
		var frontNoNumberType *NoNumberType
		if front != nil {
			frontNoNumberType, _ = front.(*NoNumberType)
		}
		stage.OnAfterNoNumberTypeUpdateCallback.OnAfterUpdate(stage, nonumbertype, frontNoNumberType)
	}
}

func (nonumbertype *NoNumberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoNumberTypeDeleteCallback != nil {
		var frontNoNumberType *NoNumberType
		if front != nil {
			frontNoNumberType, _ = front.(*NoNumberType)
		}
		stage.OnAfterNoNumberTypeDeleteCallback.OnAfterDelete(stage, nonumbertype, frontNoNumberType)
	}
}

func (nosequencetype *NoSequenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoSequenceTypeCreateCallback != nil {
		stage.OnAfterNoSequenceTypeCreateCallback.OnAfterCreate(stage, nosequencetype)
	}
}

func (nosequencetype *NoSequenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoSequenceTypeUpdateCallback != nil {
		var frontNoSequenceType *NoSequenceType
		if front != nil {
			frontNoSequenceType, _ = front.(*NoSequenceType)
		}
		stage.OnAfterNoSequenceTypeUpdateCallback.OnAfterUpdate(stage, nosequencetype, frontNoSequenceType)
	}
}

func (nosequencetype *NoSequenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoSequenceTypeDeleteCallback != nil {
		var frontNoSequenceType *NoSequenceType
		if front != nil {
			frontNoSequenceType, _ = front.(*NoSequenceType)
		}
		stage.OnAfterNoSequenceTypeDeleteCallback.OnAfterDelete(stage, nosequencetype, frontNoSequenceType)
	}
}

func (textaddresstype *TextAddressType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextAddressTypeCreateCallback != nil {
		stage.OnAfterTextAddressTypeCreateCallback.OnAfterCreate(stage, textaddresstype)
	}
}

func (textaddresstype *TextAddressType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextAddressTypeUpdateCallback != nil {
		var frontTextAddressType *TextAddressType
		if front != nil {
			frontTextAddressType, _ = front.(*TextAddressType)
		}
		stage.OnAfterTextAddressTypeUpdateCallback.OnAfterUpdate(stage, textaddresstype, frontTextAddressType)
	}
}

func (textaddresstype *TextAddressType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextAddressTypeDeleteCallback != nil {
		var frontTextAddressType *TextAddressType
		if front != nil {
			frontTextAddressType, _ = front.(*TextAddressType)
		}
		stage.OnAfterTextAddressTypeDeleteCallback.OnAfterDelete(stage, textaddresstype, frontTextAddressType)
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

func (textinstructiontype *TextInstructionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextInstructionTypeCreateCallback != nil {
		stage.OnAfterTextInstructionTypeCreateCallback.OnAfterCreate(stage, textinstructiontype)
	}
}

func (textinstructiontype *TextInstructionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextInstructionTypeUpdateCallback != nil {
		var frontTextInstructionType *TextInstructionType
		if front != nil {
			frontTextInstructionType, _ = front.(*TextInstructionType)
		}
		stage.OnAfterTextInstructionTypeUpdateCallback.OnAfterUpdate(stage, textinstructiontype, frontTextInstructionType)
	}
}

func (textinstructiontype *TextInstructionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextInstructionTypeDeleteCallback != nil {
		var frontTextInstructionType *TextInstructionType
		if front != nil {
			frontTextInstructionType, _ = front.(*TextInstructionType)
		}
		stage.OnAfterTextInstructionTypeDeleteCallback.OnAfterDelete(stage, textinstructiontype, frontTextInstructionType)
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

func (textphonetype *TextPhoneType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextPhoneTypeCreateCallback != nil {
		stage.OnAfterTextPhoneTypeCreateCallback.OnAfterCreate(stage, textphonetype)
	}
}

func (textphonetype *TextPhoneType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextPhoneTypeUpdateCallback != nil {
		var frontTextPhoneType *TextPhoneType
		if front != nil {
			frontTextPhoneType, _ = front.(*TextPhoneType)
		}
		stage.OnAfterTextPhoneTypeUpdateCallback.OnAfterUpdate(stage, textphonetype, frontTextPhoneType)
	}
}

func (textphonetype *TextPhoneType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextPhoneTypeDeleteCallback != nil {
		var frontTextPhoneType *TextPhoneType
		if front != nil {
			frontTextPhoneType, _ = front.(*TextPhoneType)
		}
		stage.OnAfterTextPhoneTypeDeleteCallback.OnAfterDelete(stage, textphonetype, frontTextPhoneType)
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

func (valangletype *ValAngleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValAngleTypeCreateCallback != nil {
		stage.OnAfterValAngleTypeCreateCallback.OnAfterCreate(stage, valangletype)
	}
}

func (valangletype *ValAngleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValAngleTypeUpdateCallback != nil {
		var frontValAngleType *ValAngleType
		if front != nil {
			frontValAngleType, _ = front.(*ValAngleType)
		}
		stage.OnAfterValAngleTypeUpdateCallback.OnAfterUpdate(stage, valangletype, frontValAngleType)
	}
}

func (valangletype *ValAngleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValAngleTypeDeleteCallback != nil {
		var frontValAngleType *ValAngleType
		if front != nil {
			frontValAngleType, _ = front.(*ValAngleType)
		}
		stage.OnAfterValAngleTypeDeleteCallback.OnAfterDelete(stage, valangletype, frontValAngleType)
	}
}

func (valbearingtype *ValBearingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValBearingTypeCreateCallback != nil {
		stage.OnAfterValBearingTypeCreateCallback.OnAfterCreate(stage, valbearingtype)
	}
}

func (valbearingtype *ValBearingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValBearingTypeUpdateCallback != nil {
		var frontValBearingType *ValBearingType
		if front != nil {
			frontValBearingType, _ = front.(*ValBearingType)
		}
		stage.OnAfterValBearingTypeUpdateCallback.OnAfterUpdate(stage, valbearingtype, frontValBearingType)
	}
}

func (valbearingtype *ValBearingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValBearingTypeDeleteCallback != nil {
		var frontValBearingType *ValBearingType
		if front != nil {
			frontValBearingType, _ = front.(*ValBearingType)
		}
		stage.OnAfterValBearingTypeDeleteCallback.OnAfterDelete(stage, valbearingtype, frontValBearingType)
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

func (valslopetype *ValSlopeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValSlopeTypeCreateCallback != nil {
		stage.OnAfterValSlopeTypeCreateCallback.OnAfterCreate(stage, valslopetype)
	}
}

func (valslopetype *ValSlopeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValSlopeTypeUpdateCallback != nil {
		var frontValSlopeType *ValSlopeType
		if front != nil {
			frontValSlopeType, _ = front.(*ValSlopeType)
		}
		stage.OnAfterValSlopeTypeUpdateCallback.OnAfterUpdate(stage, valslopetype, frontValSlopeType)
	}
}

func (valslopetype *ValSlopeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValSlopeTypeDeleteCallback != nil {
		var frontValSlopeType *ValSlopeType
		if front != nil {
			frontValSlopeType, _ = front.(*ValSlopeType)
		}
		stage.OnAfterValSlopeTypeDeleteCallback.OnAfterDelete(stage, valslopetype, frontValSlopeType)
	}
}

func (valspeedtype *ValSpeedType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValSpeedTypeCreateCallback != nil {
		stage.OnAfterValSpeedTypeCreateCallback.OnAfterCreate(stage, valspeedtype)
	}
}

func (valspeedtype *ValSpeedType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValSpeedTypeUpdateCallback != nil {
		var frontValSpeedType *ValSpeedType
		if front != nil {
			frontValSpeedType, _ = front.(*ValSpeedType)
		}
		stage.OnAfterValSpeedTypeUpdateCallback.OnAfterUpdate(stage, valspeedtype, frontValSpeedType)
	}
}

func (valspeedtype *ValSpeedType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValSpeedTypeDeleteCallback != nil {
		var frontValSpeedType *ValSpeedType
		if front != nil {
			frontValSpeedType, _ = front.(*ValSpeedType)
		}
		stage.OnAfterValSpeedTypeDeleteCallback.OnAfterDelete(stage, valspeedtype, frontValSpeedType)
	}
}

func (valweighttype *ValWeightType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValWeightTypeCreateCallback != nil {
		stage.OnAfterValWeightTypeCreateCallback.OnAfterCreate(stage, valweighttype)
	}
}

func (valweighttype *ValWeightType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValWeightTypeUpdateCallback != nil {
		var frontValWeightType *ValWeightType
		if front != nil {
			frontValWeightType, _ = front.(*ValWeightType)
		}
		stage.OnAfterValWeightTypeUpdateCallback.OnAfterUpdate(stage, valweighttype, frontValWeightType)
	}
}

func (valweighttype *ValWeightType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValWeightTypeDeleteCallback != nil {
		var frontValWeightType *ValWeightType
		if front != nil {
			frontValWeightType, _ = front.(*ValWeightType)
		}
		stage.OnAfterValWeightTypeDeleteCallback.OnAfterDelete(stage, valweighttype, frontValWeightType)
	}
}

func (xhtmltype *XHTMLType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXHTMLTypeCreateCallback != nil {
		stage.OnAfterXHTMLTypeCreateCallback.OnAfterCreate(stage, xhtmltype)
	}
}

func (xhtmltype *XHTMLType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXHTMLTypeUpdateCallback != nil {
		var frontXHTMLType *XHTMLType
		if front != nil {
			frontXHTMLType, _ = front.(*XHTMLType)
		}
		stage.OnAfterXHTMLTypeUpdateCallback.OnAfterUpdate(stage, xhtmltype, frontXHTMLType)
	}
}

func (xhtmltype *XHTMLType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXHTMLTypeDeleteCallback != nil {
		var frontXHTMLType *XHTMLType
		if front != nil {
			frontXHTMLType, _ = front.(*XHTMLType)
		}
		stage.OnAfterXHTMLTypeDeleteCallback.OnAfterDelete(stage, xhtmltype, frontXHTMLType)
	}
}

