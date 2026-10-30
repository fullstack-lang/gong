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

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApproachEquipmentAdditionalTypeCreateCallback != nil {
		stage.OnAfterCodeApproachEquipmentAdditionalTypeCreateCallback.OnAfterCreate(stage, codeapproachequipmentadditionaltype)
	}
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachEquipmentAdditionalTypeUpdateCallback != nil {
		var frontCodeApproachEquipmentAdditionalType *CodeApproachEquipmentAdditionalType
		if front != nil {
			frontCodeApproachEquipmentAdditionalType, _ = front.(*CodeApproachEquipmentAdditionalType)
		}
		stage.OnAfterCodeApproachEquipmentAdditionalTypeUpdateCallback.OnAfterUpdate(stage, codeapproachequipmentadditionaltype, frontCodeApproachEquipmentAdditionalType)
	}
}

func (codeapproachequipmentadditionaltype *CodeApproachEquipmentAdditionalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachEquipmentAdditionalTypeDeleteCallback != nil {
		var frontCodeApproachEquipmentAdditionalType *CodeApproachEquipmentAdditionalType
		if front != nil {
			frontCodeApproachEquipmentAdditionalType, _ = front.(*CodeApproachEquipmentAdditionalType)
		}
		stage.OnAfterCodeApproachEquipmentAdditionalTypeDeleteCallback.OnAfterDelete(stage, codeapproachequipmentadditionaltype, frontCodeApproachEquipmentAdditionalType)
	}
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApproachGuidanceTypeCreateCallback != nil {
		stage.OnAfterCodeApproachGuidanceTypeCreateCallback.OnAfterCreate(stage, codeapproachguidancetype)
	}
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachGuidanceTypeUpdateCallback != nil {
		var frontCodeApproachGuidanceType *CodeApproachGuidanceType
		if front != nil {
			frontCodeApproachGuidanceType, _ = front.(*CodeApproachGuidanceType)
		}
		stage.OnAfterCodeApproachGuidanceTypeUpdateCallback.OnAfterUpdate(stage, codeapproachguidancetype, frontCodeApproachGuidanceType)
	}
}

func (codeapproachguidancetype *CodeApproachGuidanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachGuidanceTypeDeleteCallback != nil {
		var frontCodeApproachGuidanceType *CodeApproachGuidanceType
		if front != nil {
			frontCodeApproachGuidanceType, _ = front.(*CodeApproachGuidanceType)
		}
		stage.OnAfterCodeApproachGuidanceTypeDeleteCallback.OnAfterDelete(stage, codeapproachguidancetype, frontCodeApproachGuidanceType)
	}
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApproachPrefixTypeCreateCallback != nil {
		stage.OnAfterCodeApproachPrefixTypeCreateCallback.OnAfterCreate(stage, codeapproachprefixtype)
	}
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachPrefixTypeUpdateCallback != nil {
		var frontCodeApproachPrefixType *CodeApproachPrefixType
		if front != nil {
			frontCodeApproachPrefixType, _ = front.(*CodeApproachPrefixType)
		}
		stage.OnAfterCodeApproachPrefixTypeUpdateCallback.OnAfterUpdate(stage, codeapproachprefixtype, frontCodeApproachPrefixType)
	}
}

func (codeapproachprefixtype *CodeApproachPrefixType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachPrefixTypeDeleteCallback != nil {
		var frontCodeApproachPrefixType *CodeApproachPrefixType
		if front != nil {
			frontCodeApproachPrefixType, _ = front.(*CodeApproachPrefixType)
		}
		stage.OnAfterCodeApproachPrefixTypeDeleteCallback.OnAfterDelete(stage, codeapproachprefixtype, frontCodeApproachPrefixType)
	}
}

func (codeapproachtype *CodeApproachType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApproachTypeCreateCallback != nil {
		stage.OnAfterCodeApproachTypeCreateCallback.OnAfterCreate(stage, codeapproachtype)
	}
}

func (codeapproachtype *CodeApproachType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachTypeUpdateCallback != nil {
		var frontCodeApproachType *CodeApproachType
		if front != nil {
			frontCodeApproachType, _ = front.(*CodeApproachType)
		}
		stage.OnAfterCodeApproachTypeUpdateCallback.OnAfterUpdate(stage, codeapproachtype, frontCodeApproachType)
	}
}

func (codeapproachtype *CodeApproachType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachTypeDeleteCallback != nil {
		var frontCodeApproachType *CodeApproachType
		if front != nil {
			frontCodeApproachType, _ = front.(*CodeApproachType)
		}
		stage.OnAfterCodeApproachTypeDeleteCallback.OnAfterDelete(stage, codeapproachtype, frontCodeApproachType)
	}
}

func (codeapprovaltype *CodeApprovalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApprovalTypeCreateCallback != nil {
		stage.OnAfterCodeApprovalTypeCreateCallback.OnAfterCreate(stage, codeapprovaltype)
	}
}

func (codeapprovaltype *CodeApprovalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApprovalTypeUpdateCallback != nil {
		var frontCodeApprovalType *CodeApprovalType
		if front != nil {
			frontCodeApprovalType, _ = front.(*CodeApprovalType)
		}
		stage.OnAfterCodeApprovalTypeUpdateCallback.OnAfterUpdate(stage, codeapprovaltype, frontCodeApprovalType)
	}
}

func (codeapprovaltype *CodeApprovalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApprovalTypeDeleteCallback != nil {
		var frontCodeApprovalType *CodeApprovalType
		if front != nil {
			frontCodeApprovalType, _ = front.(*CodeApprovalType)
		}
		stage.OnAfterCodeApprovalTypeDeleteCallback.OnAfterDelete(stage, codeapprovaltype, frontCodeApprovalType)
	}
}

func (codearcdirectiontype *CodeArcDirectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeArcDirectionTypeCreateCallback != nil {
		stage.OnAfterCodeArcDirectionTypeCreateCallback.OnAfterCreate(stage, codearcdirectiontype)
	}
}

func (codearcdirectiontype *CodeArcDirectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeArcDirectionTypeUpdateCallback != nil {
		var frontCodeArcDirectionType *CodeArcDirectionType
		if front != nil {
			frontCodeArcDirectionType, _ = front.(*CodeArcDirectionType)
		}
		stage.OnAfterCodeArcDirectionTypeUpdateCallback.OnAfterUpdate(stage, codearcdirectiontype, frontCodeArcDirectionType)
	}
}

func (codearcdirectiontype *CodeArcDirectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeArcDirectionTypeDeleteCallback != nil {
		var frontCodeArcDirectionType *CodeArcDirectionType
		if front != nil {
			frontCodeArcDirectionType, _ = front.(*CodeArcDirectionType)
		}
		stage.OnAfterCodeArcDirectionTypeDeleteCallback.OnAfterDelete(stage, codearcdirectiontype, frontCodeArcDirectionType)
	}
}

func (codebearingtype *CodeBearingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeBearingTypeCreateCallback != nil {
		stage.OnAfterCodeBearingTypeCreateCallback.OnAfterCreate(stage, codebearingtype)
	}
}

func (codebearingtype *CodeBearingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeBearingTypeUpdateCallback != nil {
		var frontCodeBearingType *CodeBearingType
		if front != nil {
			frontCodeBearingType, _ = front.(*CodeBearingType)
		}
		stage.OnAfterCodeBearingTypeUpdateCallback.OnAfterUpdate(stage, codebearingtype, frontCodeBearingType)
	}
}

func (codebearingtype *CodeBearingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeBearingTypeDeleteCallback != nil {
		var frontCodeBearingType *CodeBearingType
		if front != nil {
			frontCodeBearingType, _ = front.(*CodeBearingType)
		}
		stage.OnAfterCodeBearingTypeDeleteCallback.OnAfterDelete(stage, codebearingtype, frontCodeBearingType)
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

func (codedesignstandardtype *CodeDesignStandardType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDesignStandardTypeCreateCallback != nil {
		stage.OnAfterCodeDesignStandardTypeCreateCallback.OnAfterCreate(stage, codedesignstandardtype)
	}
}

func (codedesignstandardtype *CodeDesignStandardType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignStandardTypeUpdateCallback != nil {
		var frontCodeDesignStandardType *CodeDesignStandardType
		if front != nil {
			frontCodeDesignStandardType, _ = front.(*CodeDesignStandardType)
		}
		stage.OnAfterCodeDesignStandardTypeUpdateCallback.OnAfterUpdate(stage, codedesignstandardtype, frontCodeDesignStandardType)
	}
}

func (codedesignstandardtype *CodeDesignStandardType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignStandardTypeDeleteCallback != nil {
		var frontCodeDesignStandardType *CodeDesignStandardType
		if front != nil {
			frontCodeDesignStandardType, _ = front.(*CodeDesignStandardType)
		}
		stage.OnAfterCodeDesignStandardTypeDeleteCallback.OnAfterDelete(stage, codedesignstandardtype, frontCodeDesignStandardType)
	}
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDesignatedPointDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeDesignatedPointDesignatorTypeCreateCallback.OnAfterCreate(stage, codedesignatedpointdesignatortype)
	}
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignatedPointDesignatorTypeUpdateCallback != nil {
		var frontCodeDesignatedPointDesignatorType *CodeDesignatedPointDesignatorType
		if front != nil {
			frontCodeDesignatedPointDesignatorType, _ = front.(*CodeDesignatedPointDesignatorType)
		}
		stage.OnAfterCodeDesignatedPointDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codedesignatedpointdesignatortype, frontCodeDesignatedPointDesignatorType)
	}
}

func (codedesignatedpointdesignatortype *CodeDesignatedPointDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignatedPointDesignatorTypeDeleteCallback != nil {
		var frontCodeDesignatedPointDesignatorType *CodeDesignatedPointDesignatorType
		if front != nil {
			frontCodeDesignatedPointDesignatorType, _ = front.(*CodeDesignatedPointDesignatorType)
		}
		stage.OnAfterCodeDesignatedPointDesignatorTypeDeleteCallback.OnAfterDelete(stage, codedesignatedpointdesignatortype, frontCodeDesignatedPointDesignatorType)
	}
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDirectionReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeDirectionReferenceTypeCreateCallback.OnAfterCreate(stage, codedirectionreferencetype)
	}
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDirectionReferenceTypeUpdateCallback != nil {
		var frontCodeDirectionReferenceType *CodeDirectionReferenceType
		if front != nil {
			frontCodeDirectionReferenceType, _ = front.(*CodeDirectionReferenceType)
		}
		stage.OnAfterCodeDirectionReferenceTypeUpdateCallback.OnAfterUpdate(stage, codedirectionreferencetype, frontCodeDirectionReferenceType)
	}
}

func (codedirectionreferencetype *CodeDirectionReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDirectionReferenceTypeDeleteCallback != nil {
		var frontCodeDirectionReferenceType *CodeDirectionReferenceType
		if front != nil {
			frontCodeDirectionReferenceType, _ = front.(*CodeDirectionReferenceType)
		}
		stage.OnAfterCodeDirectionReferenceTypeDeleteCallback.OnAfterDelete(stage, codedirectionreferencetype, frontCodeDirectionReferenceType)
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

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeEquipmentUnavailableTypeCreateCallback != nil {
		stage.OnAfterCodeEquipmentUnavailableTypeCreateCallback.OnAfterCreate(stage, codeequipmentunavailabletype)
	}
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeEquipmentUnavailableTypeUpdateCallback != nil {
		var frontCodeEquipmentUnavailableType *CodeEquipmentUnavailableType
		if front != nil {
			frontCodeEquipmentUnavailableType, _ = front.(*CodeEquipmentUnavailableType)
		}
		stage.OnAfterCodeEquipmentUnavailableTypeUpdateCallback.OnAfterUpdate(stage, codeequipmentunavailabletype, frontCodeEquipmentUnavailableType)
	}
}

func (codeequipmentunavailabletype *CodeEquipmentUnavailableType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeEquipmentUnavailableTypeDeleteCallback != nil {
		var frontCodeEquipmentUnavailableType *CodeEquipmentUnavailableType
		if front != nil {
			frontCodeEquipmentUnavailableType, _ = front.(*CodeEquipmentUnavailableType)
		}
		stage.OnAfterCodeEquipmentUnavailableTypeDeleteCallback.OnAfterDelete(stage, codeequipmentunavailabletype, frontCodeEquipmentUnavailableType)
	}
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFinalGuidanceTypeCreateCallback != nil {
		stage.OnAfterCodeFinalGuidanceTypeCreateCallback.OnAfterCreate(stage, codefinalguidancetype)
	}
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFinalGuidanceTypeUpdateCallback != nil {
		var frontCodeFinalGuidanceType *CodeFinalGuidanceType
		if front != nil {
			frontCodeFinalGuidanceType, _ = front.(*CodeFinalGuidanceType)
		}
		stage.OnAfterCodeFinalGuidanceTypeUpdateCallback.OnAfterUpdate(stage, codefinalguidancetype, frontCodeFinalGuidanceType)
	}
}

func (codefinalguidancetype *CodeFinalGuidanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFinalGuidanceTypeDeleteCallback != nil {
		var frontCodeFinalGuidanceType *CodeFinalGuidanceType
		if front != nil {
			frontCodeFinalGuidanceType, _ = front.(*CodeFinalGuidanceType)
		}
		stage.OnAfterCodeFinalGuidanceTypeDeleteCallback.OnAfterDelete(stage, codefinalguidancetype, frontCodeFinalGuidanceType)
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

func (codeheightreferencetype *CodeHeightReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeHeightReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeHeightReferenceTypeCreateCallback.OnAfterCreate(stage, codeheightreferencetype)
	}
}

func (codeheightreferencetype *CodeHeightReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHeightReferenceTypeUpdateCallback != nil {
		var frontCodeHeightReferenceType *CodeHeightReferenceType
		if front != nil {
			frontCodeHeightReferenceType, _ = front.(*CodeHeightReferenceType)
		}
		stage.OnAfterCodeHeightReferenceTypeUpdateCallback.OnAfterUpdate(stage, codeheightreferencetype, frontCodeHeightReferenceType)
	}
}

func (codeheightreferencetype *CodeHeightReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHeightReferenceTypeDeleteCallback != nil {
		var frontCodeHeightReferenceType *CodeHeightReferenceType
		if front != nil {
			frontCodeHeightReferenceType, _ = front.(*CodeHeightReferenceType)
		}
		stage.OnAfterCodeHeightReferenceTypeDeleteCallback.OnAfterDelete(stage, codeheightreferencetype, frontCodeHeightReferenceType)
	}
}

func (codeholdingusagetype *CodeHoldingUsageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeHoldingUsageTypeCreateCallback != nil {
		stage.OnAfterCodeHoldingUsageTypeCreateCallback.OnAfterCreate(stage, codeholdingusagetype)
	}
}

func (codeholdingusagetype *CodeHoldingUsageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHoldingUsageTypeUpdateCallback != nil {
		var frontCodeHoldingUsageType *CodeHoldingUsageType
		if front != nil {
			frontCodeHoldingUsageType, _ = front.(*CodeHoldingUsageType)
		}
		stage.OnAfterCodeHoldingUsageTypeUpdateCallback.OnAfterUpdate(stage, codeholdingusagetype, frontCodeHoldingUsageType)
	}
}

func (codeholdingusagetype *CodeHoldingUsageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHoldingUsageTypeDeleteCallback != nil {
		var frontCodeHoldingUsageType *CodeHoldingUsageType
		if front != nil {
			frontCodeHoldingUsageType, _ = front.(*CodeHoldingUsageType)
		}
		stage.OnAfterCodeHoldingUsageTypeDeleteCallback.OnAfterDelete(stage, codeholdingusagetype, frontCodeHoldingUsageType)
	}
}

func (codeholdingusetype *CodeHoldingUseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeHoldingUseTypeCreateCallback != nil {
		stage.OnAfterCodeHoldingUseTypeCreateCallback.OnAfterCreate(stage, codeholdingusetype)
	}
}

func (codeholdingusetype *CodeHoldingUseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHoldingUseTypeUpdateCallback != nil {
		var frontCodeHoldingUseType *CodeHoldingUseType
		if front != nil {
			frontCodeHoldingUseType, _ = front.(*CodeHoldingUseType)
		}
		stage.OnAfterCodeHoldingUseTypeUpdateCallback.OnAfterUpdate(stage, codeholdingusetype, frontCodeHoldingUseType)
	}
}

func (codeholdingusetype *CodeHoldingUseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHoldingUseTypeDeleteCallback != nil {
		var frontCodeHoldingUseType *CodeHoldingUseType
		if front != nil {
			frontCodeHoldingUseType, _ = front.(*CodeHoldingUseType)
		}
		stage.OnAfterCodeHoldingUseTypeDeleteCallback.OnAfterDelete(stage, codeholdingusetype, frontCodeHoldingUseType)
	}
}

func (codeicaocountrytype *CodeICAOCountryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeICAOCountryTypeCreateCallback != nil {
		stage.OnAfterCodeICAOCountryTypeCreateCallback.OnAfterCreate(stage, codeicaocountrytype)
	}
}

func (codeicaocountrytype *CodeICAOCountryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeICAOCountryTypeUpdateCallback != nil {
		var frontCodeICAOCountryType *CodeICAOCountryType
		if front != nil {
			frontCodeICAOCountryType, _ = front.(*CodeICAOCountryType)
		}
		stage.OnAfterCodeICAOCountryTypeUpdateCallback.OnAfterUpdate(stage, codeicaocountrytype, frontCodeICAOCountryType)
	}
}

func (codeicaocountrytype *CodeICAOCountryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeICAOCountryTypeDeleteCallback != nil {
		var frontCodeICAOCountryType *CodeICAOCountryType
		if front != nil {
			frontCodeICAOCountryType, _ = front.(*CodeICAOCountryType)
		}
		stage.OnAfterCodeICAOCountryTypeDeleteCallback.OnAfterDelete(stage, codeicaocountrytype, frontCodeICAOCountryType)
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

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMinimaFinalApproachPathTypeCreateCallback != nil {
		stage.OnAfterCodeMinimaFinalApproachPathTypeCreateCallback.OnAfterCreate(stage, codeminimafinalapproachpathtype)
	}
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMinimaFinalApproachPathTypeUpdateCallback != nil {
		var frontCodeMinimaFinalApproachPathType *CodeMinimaFinalApproachPathType
		if front != nil {
			frontCodeMinimaFinalApproachPathType, _ = front.(*CodeMinimaFinalApproachPathType)
		}
		stage.OnAfterCodeMinimaFinalApproachPathTypeUpdateCallback.OnAfterUpdate(stage, codeminimafinalapproachpathtype, frontCodeMinimaFinalApproachPathType)
	}
}

func (codeminimafinalapproachpathtype *CodeMinimaFinalApproachPathType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMinimaFinalApproachPathTypeDeleteCallback != nil {
		var frontCodeMinimaFinalApproachPathType *CodeMinimaFinalApproachPathType
		if front != nil {
			frontCodeMinimaFinalApproachPathType, _ = front.(*CodeMinimaFinalApproachPathType)
		}
		stage.OnAfterCodeMinimaFinalApproachPathTypeDeleteCallback.OnAfterDelete(stage, codeminimafinalapproachpathtype, frontCodeMinimaFinalApproachPathType)
	}
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMinimumAltitudeTypeCreateCallback != nil {
		stage.OnAfterCodeMinimumAltitudeTypeCreateCallback.OnAfterCreate(stage, codeminimumaltitudetype)
	}
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMinimumAltitudeTypeUpdateCallback != nil {
		var frontCodeMinimumAltitudeType *CodeMinimumAltitudeType
		if front != nil {
			frontCodeMinimumAltitudeType, _ = front.(*CodeMinimumAltitudeType)
		}
		stage.OnAfterCodeMinimumAltitudeTypeUpdateCallback.OnAfterUpdate(stage, codeminimumaltitudetype, frontCodeMinimumAltitudeType)
	}
}

func (codeminimumaltitudetype *CodeMinimumAltitudeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMinimumAltitudeTypeDeleteCallback != nil {
		var frontCodeMinimumAltitudeType *CodeMinimumAltitudeType
		if front != nil {
			frontCodeMinimumAltitudeType, _ = front.(*CodeMinimumAltitudeType)
		}
		stage.OnAfterCodeMinimumAltitudeTypeDeleteCallback.OnAfterDelete(stage, codeminimumaltitudetype, frontCodeMinimumAltitudeType)
	}
}

func (codeminimumheighttype *CodeMinimumHeightType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMinimumHeightTypeCreateCallback != nil {
		stage.OnAfterCodeMinimumHeightTypeCreateCallback.OnAfterCreate(stage, codeminimumheighttype)
	}
}

func (codeminimumheighttype *CodeMinimumHeightType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMinimumHeightTypeUpdateCallback != nil {
		var frontCodeMinimumHeightType *CodeMinimumHeightType
		if front != nil {
			frontCodeMinimumHeightType, _ = front.(*CodeMinimumHeightType)
		}
		stage.OnAfterCodeMinimumHeightTypeUpdateCallback.OnAfterUpdate(stage, codeminimumheighttype, frontCodeMinimumHeightType)
	}
}

func (codeminimumheighttype *CodeMinimumHeightType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMinimumHeightTypeDeleteCallback != nil {
		var frontCodeMinimumHeightType *CodeMinimumHeightType
		if front != nil {
			frontCodeMinimumHeightType, _ = front.(*CodeMinimumHeightType)
		}
		stage.OnAfterCodeMinimumHeightTypeDeleteCallback.OnAfterDelete(stage, codeminimumheighttype, frontCodeMinimumHeightType)
	}
}

func (codemissedapproachtype *CodeMissedApproachType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMissedApproachTypeCreateCallback != nil {
		stage.OnAfterCodeMissedApproachTypeCreateCallback.OnAfterCreate(stage, codemissedapproachtype)
	}
}

func (codemissedapproachtype *CodeMissedApproachType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMissedApproachTypeUpdateCallback != nil {
		var frontCodeMissedApproachType *CodeMissedApproachType
		if front != nil {
			frontCodeMissedApproachType, _ = front.(*CodeMissedApproachType)
		}
		stage.OnAfterCodeMissedApproachTypeUpdateCallback.OnAfterUpdate(stage, codemissedapproachtype, frontCodeMissedApproachType)
	}
}

func (codemissedapproachtype *CodeMissedApproachType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMissedApproachTypeDeleteCallback != nil {
		var frontCodeMissedApproachType *CodeMissedApproachType
		if front != nil {
			frontCodeMissedApproachType, _ = front.(*CodeMissedApproachType)
		}
		stage.OnAfterCodeMissedApproachTypeDeleteCallback.OnAfterDelete(stage, codemissedapproachtype, frontCodeMissedApproachType)
	}
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavigationAreaRestrictionTypeCreateCallback != nil {
		stage.OnAfterCodeNavigationAreaRestrictionTypeCreateCallback.OnAfterCreate(stage, codenavigationarearestrictiontype)
	}
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavigationAreaRestrictionTypeUpdateCallback != nil {
		var frontCodeNavigationAreaRestrictionType *CodeNavigationAreaRestrictionType
		if front != nil {
			frontCodeNavigationAreaRestrictionType, _ = front.(*CodeNavigationAreaRestrictionType)
		}
		stage.OnAfterCodeNavigationAreaRestrictionTypeUpdateCallback.OnAfterUpdate(stage, codenavigationarearestrictiontype, frontCodeNavigationAreaRestrictionType)
	}
}

func (codenavigationarearestrictiontype *CodeNavigationAreaRestrictionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavigationAreaRestrictionTypeDeleteCallback != nil {
		var frontCodeNavigationAreaRestrictionType *CodeNavigationAreaRestrictionType
		if front != nil {
			frontCodeNavigationAreaRestrictionType, _ = front.(*CodeNavigationAreaRestrictionType)
		}
		stage.OnAfterCodeNavigationAreaRestrictionTypeDeleteCallback.OnAfterDelete(stage, codenavigationarearestrictiontype, frontCodeNavigationAreaRestrictionType)
	}
}

func (codenavigationareatype *CodeNavigationAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavigationAreaTypeCreateCallback != nil {
		stage.OnAfterCodeNavigationAreaTypeCreateCallback.OnAfterCreate(stage, codenavigationareatype)
	}
}

func (codenavigationareatype *CodeNavigationAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavigationAreaTypeUpdateCallback != nil {
		var frontCodeNavigationAreaType *CodeNavigationAreaType
		if front != nil {
			frontCodeNavigationAreaType, _ = front.(*CodeNavigationAreaType)
		}
		stage.OnAfterCodeNavigationAreaTypeUpdateCallback.OnAfterUpdate(stage, codenavigationareatype, frontCodeNavigationAreaType)
	}
}

func (codenavigationareatype *CodeNavigationAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavigationAreaTypeDeleteCallback != nil {
		var frontCodeNavigationAreaType *CodeNavigationAreaType
		if front != nil {
			frontCodeNavigationAreaType, _ = front.(*CodeNavigationAreaType)
		}
		stage.OnAfterCodeNavigationAreaTypeDeleteCallback.OnAfterDelete(stage, codenavigationareatype, frontCodeNavigationAreaType)
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

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeProcedureAvailabilityTypeCreateCallback != nil {
		stage.OnAfterCodeProcedureAvailabilityTypeCreateCallback.OnAfterCreate(stage, codeprocedureavailabilitytype)
	}
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedureAvailabilityTypeUpdateCallback != nil {
		var frontCodeProcedureAvailabilityType *CodeProcedureAvailabilityType
		if front != nil {
			frontCodeProcedureAvailabilityType, _ = front.(*CodeProcedureAvailabilityType)
		}
		stage.OnAfterCodeProcedureAvailabilityTypeUpdateCallback.OnAfterUpdate(stage, codeprocedureavailabilitytype, frontCodeProcedureAvailabilityType)
	}
}

func (codeprocedureavailabilitytype *CodeProcedureAvailabilityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedureAvailabilityTypeDeleteCallback != nil {
		var frontCodeProcedureAvailabilityType *CodeProcedureAvailabilityType
		if front != nil {
			frontCodeProcedureAvailabilityType, _ = front.(*CodeProcedureAvailabilityType)
		}
		stage.OnAfterCodeProcedureAvailabilityTypeDeleteCallback.OnAfterDelete(stage, codeprocedureavailabilitytype, frontCodeProcedureAvailabilityType)
	}
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeProcedureCodingStandardTypeCreateCallback != nil {
		stage.OnAfterCodeProcedureCodingStandardTypeCreateCallback.OnAfterCreate(stage, codeprocedurecodingstandardtype)
	}
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedureCodingStandardTypeUpdateCallback != nil {
		var frontCodeProcedureCodingStandardType *CodeProcedureCodingStandardType
		if front != nil {
			frontCodeProcedureCodingStandardType, _ = front.(*CodeProcedureCodingStandardType)
		}
		stage.OnAfterCodeProcedureCodingStandardTypeUpdateCallback.OnAfterUpdate(stage, codeprocedurecodingstandardtype, frontCodeProcedureCodingStandardType)
	}
}

func (codeprocedurecodingstandardtype *CodeProcedureCodingStandardType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedureCodingStandardTypeDeleteCallback != nil {
		var frontCodeProcedureCodingStandardType *CodeProcedureCodingStandardType
		if front != nil {
			frontCodeProcedureCodingStandardType, _ = front.(*CodeProcedureCodingStandardType)
		}
		stage.OnAfterCodeProcedureCodingStandardTypeDeleteCallback.OnAfterDelete(stage, codeprocedurecodingstandardtype, frontCodeProcedureCodingStandardType)
	}
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeProcedureDistanceTypeCreateCallback != nil {
		stage.OnAfterCodeProcedureDistanceTypeCreateCallback.OnAfterCreate(stage, codeproceduredistancetype)
	}
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedureDistanceTypeUpdateCallback != nil {
		var frontCodeProcedureDistanceType *CodeProcedureDistanceType
		if front != nil {
			frontCodeProcedureDistanceType, _ = front.(*CodeProcedureDistanceType)
		}
		stage.OnAfterCodeProcedureDistanceTypeUpdateCallback.OnAfterUpdate(stage, codeproceduredistancetype, frontCodeProcedureDistanceType)
	}
}

func (codeproceduredistancetype *CodeProcedureDistanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedureDistanceTypeDeleteCallback != nil {
		var frontCodeProcedureDistanceType *CodeProcedureDistanceType
		if front != nil {
			frontCodeProcedureDistanceType, _ = front.(*CodeProcedureDistanceType)
		}
		stage.OnAfterCodeProcedureDistanceTypeDeleteCallback.OnAfterDelete(stage, codeproceduredistancetype, frontCodeProcedureDistanceType)
	}
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeProcedureFixRoleTypeCreateCallback != nil {
		stage.OnAfterCodeProcedureFixRoleTypeCreateCallback.OnAfterCreate(stage, codeprocedurefixroletype)
	}
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedureFixRoleTypeUpdateCallback != nil {
		var frontCodeProcedureFixRoleType *CodeProcedureFixRoleType
		if front != nil {
			frontCodeProcedureFixRoleType, _ = front.(*CodeProcedureFixRoleType)
		}
		stage.OnAfterCodeProcedureFixRoleTypeUpdateCallback.OnAfterUpdate(stage, codeprocedurefixroletype, frontCodeProcedureFixRoleType)
	}
}

func (codeprocedurefixroletype *CodeProcedureFixRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedureFixRoleTypeDeleteCallback != nil {
		var frontCodeProcedureFixRoleType *CodeProcedureFixRoleType
		if front != nil {
			frontCodeProcedureFixRoleType, _ = front.(*CodeProcedureFixRoleType)
		}
		stage.OnAfterCodeProcedureFixRoleTypeDeleteCallback.OnAfterDelete(stage, codeprocedurefixroletype, frontCodeProcedureFixRoleType)
	}
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeProcedurePhaseTypeCreateCallback != nil {
		stage.OnAfterCodeProcedurePhaseTypeCreateCallback.OnAfterCreate(stage, codeprocedurephasetype)
	}
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedurePhaseTypeUpdateCallback != nil {
		var frontCodeProcedurePhaseType *CodeProcedurePhaseType
		if front != nil {
			frontCodeProcedurePhaseType, _ = front.(*CodeProcedurePhaseType)
		}
		stage.OnAfterCodeProcedurePhaseTypeUpdateCallback.OnAfterUpdate(stage, codeprocedurephasetype, frontCodeProcedurePhaseType)
	}
}

func (codeprocedurephasetype *CodeProcedurePhaseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProcedurePhaseTypeDeleteCallback != nil {
		var frontCodeProcedurePhaseType *CodeProcedurePhaseType
		if front != nil {
			frontCodeProcedurePhaseType, _ = front.(*CodeProcedurePhaseType)
		}
		stage.OnAfterCodeProcedurePhaseTypeDeleteCallback.OnAfterDelete(stage, codeprocedurephasetype, frontCodeProcedurePhaseType)
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

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeReferencePathIdentifierTypeCreateCallback != nil {
		stage.OnAfterCodeReferencePathIdentifierTypeCreateCallback.OnAfterCreate(stage, codereferencepathidentifiertype)
	}
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeReferencePathIdentifierTypeUpdateCallback != nil {
		var frontCodeReferencePathIdentifierType *CodeReferencePathIdentifierType
		if front != nil {
			frontCodeReferencePathIdentifierType, _ = front.(*CodeReferencePathIdentifierType)
		}
		stage.OnAfterCodeReferencePathIdentifierTypeUpdateCallback.OnAfterUpdate(stage, codereferencepathidentifiertype, frontCodeReferencePathIdentifierType)
	}
}

func (codereferencepathidentifiertype *CodeReferencePathIdentifierType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeReferencePathIdentifierTypeDeleteCallback != nil {
		var frontCodeReferencePathIdentifierType *CodeReferencePathIdentifierType
		if front != nil {
			frontCodeReferencePathIdentifierType, _ = front.(*CodeReferencePathIdentifierType)
		}
		stage.OnAfterCodeReferencePathIdentifierTypeDeleteCallback.OnAfterDelete(stage, codereferencepathidentifiertype, frontCodeReferencePathIdentifierType)
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

func (coderelativepositiontype *CodeRelativePositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRelativePositionTypeCreateCallback != nil {
		stage.OnAfterCodeRelativePositionTypeCreateCallback.OnAfterCreate(stage, coderelativepositiontype)
	}
}

func (coderelativepositiontype *CodeRelativePositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRelativePositionTypeUpdateCallback != nil {
		var frontCodeRelativePositionType *CodeRelativePositionType
		if front != nil {
			frontCodeRelativePositionType, _ = front.(*CodeRelativePositionType)
		}
		stage.OnAfterCodeRelativePositionTypeUpdateCallback.OnAfterUpdate(stage, coderelativepositiontype, frontCodeRelativePositionType)
	}
}

func (coderelativepositiontype *CodeRelativePositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRelativePositionTypeDeleteCallback != nil {
		var frontCodeRelativePositionType *CodeRelativePositionType
		if front != nil {
			frontCodeRelativePositionType, _ = front.(*CodeRelativePositionType)
		}
		stage.OnAfterCodeRelativePositionTypeDeleteCallback.OnAfterDelete(stage, coderelativepositiontype, frontCodeRelativePositionType)
	}
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRouteIndicatorTypeCreateCallback != nil {
		stage.OnAfterCodeRouteIndicatorTypeCreateCallback.OnAfterCreate(stage, coderouteindicatortype)
	}
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteIndicatorTypeUpdateCallback != nil {
		var frontCodeRouteIndicatorType *CodeRouteIndicatorType
		if front != nil {
			frontCodeRouteIndicatorType, _ = front.(*CodeRouteIndicatorType)
		}
		stage.OnAfterCodeRouteIndicatorTypeUpdateCallback.OnAfterUpdate(stage, coderouteindicatortype, frontCodeRouteIndicatorType)
	}
}

func (coderouteindicatortype *CodeRouteIndicatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRouteIndicatorTypeDeleteCallback != nil {
		var frontCodeRouteIndicatorType *CodeRouteIndicatorType
		if front != nil {
			frontCodeRouteIndicatorType, _ = front.(*CodeRouteIndicatorType)
		}
		stage.OnAfterCodeRouteIndicatorTypeDeleteCallback.OnAfterDelete(stage, coderouteindicatortype, frontCodeRouteIndicatorType)
	}
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSafeAltitudeTypeCreateCallback != nil {
		stage.OnAfterCodeSafeAltitudeTypeCreateCallback.OnAfterCreate(stage, codesafealtitudetype)
	}
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSafeAltitudeTypeUpdateCallback != nil {
		var frontCodeSafeAltitudeType *CodeSafeAltitudeType
		if front != nil {
			frontCodeSafeAltitudeType, _ = front.(*CodeSafeAltitudeType)
		}
		stage.OnAfterCodeSafeAltitudeTypeUpdateCallback.OnAfterUpdate(stage, codesafealtitudetype, frontCodeSafeAltitudeType)
	}
}

func (codesafealtitudetype *CodeSafeAltitudeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSafeAltitudeTypeDeleteCallback != nil {
		var frontCodeSafeAltitudeType *CodeSafeAltitudeType
		if front != nil {
			frontCodeSafeAltitudeType, _ = front.(*CodeSafeAltitudeType)
		}
		stage.OnAfterCodeSafeAltitudeTypeDeleteCallback.OnAfterDelete(stage, codesafealtitudetype, frontCodeSafeAltitudeType)
	}
}

func (codesegmentpathtype *CodeSegmentPathType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSegmentPathTypeCreateCallback != nil {
		stage.OnAfterCodeSegmentPathTypeCreateCallback.OnAfterCreate(stage, codesegmentpathtype)
	}
}

func (codesegmentpathtype *CodeSegmentPathType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSegmentPathTypeUpdateCallback != nil {
		var frontCodeSegmentPathType *CodeSegmentPathType
		if front != nil {
			frontCodeSegmentPathType, _ = front.(*CodeSegmentPathType)
		}
		stage.OnAfterCodeSegmentPathTypeUpdateCallback.OnAfterUpdate(stage, codesegmentpathtype, frontCodeSegmentPathType)
	}
}

func (codesegmentpathtype *CodeSegmentPathType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSegmentPathTypeDeleteCallback != nil {
		var frontCodeSegmentPathType *CodeSegmentPathType
		if front != nil {
			frontCodeSegmentPathType, _ = front.(*CodeSegmentPathType)
		}
		stage.OnAfterCodeSegmentPathTypeDeleteCallback.OnAfterDelete(stage, codesegmentpathtype, frontCodeSegmentPathType)
	}
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSegmentTerminationTypeCreateCallback != nil {
		stage.OnAfterCodeSegmentTerminationTypeCreateCallback.OnAfterCreate(stage, codesegmentterminationtype)
	}
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSegmentTerminationTypeUpdateCallback != nil {
		var frontCodeSegmentTerminationType *CodeSegmentTerminationType
		if front != nil {
			frontCodeSegmentTerminationType, _ = front.(*CodeSegmentTerminationType)
		}
		stage.OnAfterCodeSegmentTerminationTypeUpdateCallback.OnAfterUpdate(stage, codesegmentterminationtype, frontCodeSegmentTerminationType)
	}
}

func (codesegmentterminationtype *CodeSegmentTerminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSegmentTerminationTypeDeleteCallback != nil {
		var frontCodeSegmentTerminationType *CodeSegmentTerminationType
		if front != nil {
			frontCodeSegmentTerminationType, _ = front.(*CodeSegmentTerminationType)
		}
		stage.OnAfterCodeSegmentTerminationTypeDeleteCallback.OnAfterDelete(stage, codesegmentterminationtype, frontCodeSegmentTerminationType)
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

func (codetaatype *CodeTAAType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTAATypeCreateCallback != nil {
		stage.OnAfterCodeTAATypeCreateCallback.OnAfterCreate(stage, codetaatype)
	}
}

func (codetaatype *CodeTAAType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTAATypeUpdateCallback != nil {
		var frontCodeTAAType *CodeTAAType
		if front != nil {
			frontCodeTAAType, _ = front.(*CodeTAAType)
		}
		stage.OnAfterCodeTAATypeUpdateCallback.OnAfterUpdate(stage, codetaatype, frontCodeTAAType)
	}
}

func (codetaatype *CodeTAAType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTAATypeDeleteCallback != nil {
		var frontCodeTAAType *CodeTAAType
		if front != nil {
			frontCodeTAAType, _ = front.(*CodeTAAType)
		}
		stage.OnAfterCodeTAATypeDeleteCallback.OnAfterDelete(stage, codetaatype, frontCodeTAAType)
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

func (codetrajectorytype *CodeTrajectoryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTrajectoryTypeCreateCallback != nil {
		stage.OnAfterCodeTrajectoryTypeCreateCallback.OnAfterCreate(stage, codetrajectorytype)
	}
}

func (codetrajectorytype *CodeTrajectoryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTrajectoryTypeUpdateCallback != nil {
		var frontCodeTrajectoryType *CodeTrajectoryType
		if front != nil {
			frontCodeTrajectoryType, _ = front.(*CodeTrajectoryType)
		}
		stage.OnAfterCodeTrajectoryTypeUpdateCallback.OnAfterUpdate(stage, codetrajectorytype, frontCodeTrajectoryType)
	}
}

func (codetrajectorytype *CodeTrajectoryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTrajectoryTypeDeleteCallback != nil {
		var frontCodeTrajectoryType *CodeTrajectoryType
		if front != nil {
			frontCodeTrajectoryType, _ = front.(*CodeTrajectoryType)
		}
		stage.OnAfterCodeTrajectoryTypeDeleteCallback.OnAfterDelete(stage, codetrajectorytype, frontCodeTrajectoryType)
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

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextSIDSTARDesignatorTypeCreateCallback != nil {
		stage.OnAfterTextSIDSTARDesignatorTypeCreateCallback.OnAfterCreate(stage, textsidstardesignatortype)
	}
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextSIDSTARDesignatorTypeUpdateCallback != nil {
		var frontTextSIDSTARDesignatorType *TextSIDSTARDesignatorType
		if front != nil {
			frontTextSIDSTARDesignatorType, _ = front.(*TextSIDSTARDesignatorType)
		}
		stage.OnAfterTextSIDSTARDesignatorTypeUpdateCallback.OnAfterUpdate(stage, textsidstardesignatortype, frontTextSIDSTARDesignatorType)
	}
}

func (textsidstardesignatortype *TextSIDSTARDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextSIDSTARDesignatorTypeDeleteCallback != nil {
		var frontTextSIDSTARDesignatorType *TextSIDSTARDesignatorType
		if front != nil {
			frontTextSIDSTARDesignatorType, _ = front.(*TextSIDSTARDesignatorType)
		}
		stage.OnAfterTextSIDSTARDesignatorTypeDeleteCallback.OnAfterDelete(stage, textsidstardesignatortype, frontTextSIDSTARDesignatorType)
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

func (valalarmlimittype *ValAlarmLimitType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValAlarmLimitTypeCreateCallback != nil {
		stage.OnAfterValAlarmLimitTypeCreateCallback.OnAfterCreate(stage, valalarmlimittype)
	}
}

func (valalarmlimittype *ValAlarmLimitType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValAlarmLimitTypeUpdateCallback != nil {
		var frontValAlarmLimitType *ValAlarmLimitType
		if front != nil {
			frontValAlarmLimitType, _ = front.(*ValAlarmLimitType)
		}
		stage.OnAfterValAlarmLimitTypeUpdateCallback.OnAfterUpdate(stage, valalarmlimittype, frontValAlarmLimitType)
	}
}

func (valalarmlimittype *ValAlarmLimitType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValAlarmLimitTypeDeleteCallback != nil {
		var frontValAlarmLimitType *ValAlarmLimitType
		if front != nil {
			frontValAlarmLimitType, _ = front.(*ValAlarmLimitType)
		}
		stage.OnAfterValAlarmLimitTypeDeleteCallback.OnAfterDelete(stage, valalarmlimittype, frontValAlarmLimitType)
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

func (valchannelnumbertype *ValChannelNumberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValChannelNumberTypeCreateCallback != nil {
		stage.OnAfterValChannelNumberTypeCreateCallback.OnAfterCreate(stage, valchannelnumbertype)
	}
}

func (valchannelnumbertype *ValChannelNumberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValChannelNumberTypeUpdateCallback != nil {
		var frontValChannelNumberType *ValChannelNumberType
		if front != nil {
			frontValChannelNumberType, _ = front.(*ValChannelNumberType)
		}
		stage.OnAfterValChannelNumberTypeUpdateCallback.OnAfterUpdate(stage, valchannelnumbertype, frontValChannelNumberType)
	}
}

func (valchannelnumbertype *ValChannelNumberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValChannelNumberTypeDeleteCallback != nil {
		var frontValChannelNumberType *ValChannelNumberType
		if front != nil {
			frontValChannelNumberType, _ = front.(*ValChannelNumberType)
		}
		stage.OnAfterValChannelNumberTypeDeleteCallback.OnAfterDelete(stage, valchannelnumbertype, frontValChannelNumberType)
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

func (valhextype *ValHexType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValHexTypeCreateCallback != nil {
		stage.OnAfterValHexTypeCreateCallback.OnAfterCreate(stage, valhextype)
	}
}

func (valhextype *ValHexType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValHexTypeUpdateCallback != nil {
		var frontValHexType *ValHexType
		if front != nil {
			frontValHexType, _ = front.(*ValHexType)
		}
		stage.OnAfterValHexTypeUpdateCallback.OnAfterUpdate(stage, valhextype, frontValHexType)
	}
}

func (valhextype *ValHexType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValHexTypeDeleteCallback != nil {
		var frontValHexType *ValHexType
		if front != nil {
			frontValHexType, _ = front.(*ValHexType)
		}
		stage.OnAfterValHexTypeDeleteCallback.OnAfterDelete(stage, valhextype, frontValHexType)
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

func (valtemperaturetype *ValTemperatureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValTemperatureTypeCreateCallback != nil {
		stage.OnAfterValTemperatureTypeCreateCallback.OnAfterCreate(stage, valtemperaturetype)
	}
}

func (valtemperaturetype *ValTemperatureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValTemperatureTypeUpdateCallback != nil {
		var frontValTemperatureType *ValTemperatureType
		if front != nil {
			frontValTemperatureType, _ = front.(*ValTemperatureType)
		}
		stage.OnAfterValTemperatureTypeUpdateCallback.OnAfterUpdate(stage, valtemperaturetype, frontValTemperatureType)
	}
}

func (valtemperaturetype *ValTemperatureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValTemperatureTypeDeleteCallback != nil {
		var frontValTemperatureType *ValTemperatureType
		if front != nil {
			frontValTemperatureType, _ = front.(*ValTemperatureType)
		}
		stage.OnAfterValTemperatureTypeDeleteCallback.OnAfterDelete(stage, valtemperaturetype, frontValTemperatureType)
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

