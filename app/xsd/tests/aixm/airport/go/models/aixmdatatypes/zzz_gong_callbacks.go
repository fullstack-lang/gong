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

func (codeaircraftstandtype *CodeAircraftStandType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftStandTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftStandTypeCreateCallback.OnAfterCreate(stage, codeaircraftstandtype)
	}
}

func (codeaircraftstandtype *CodeAircraftStandType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftStandTypeUpdateCallback != nil {
		var frontCodeAircraftStandType *CodeAircraftStandType
		if front != nil {
			frontCodeAircraftStandType, _ = front.(*CodeAircraftStandType)
		}
		stage.OnAfterCodeAircraftStandTypeUpdateCallback.OnAfterUpdate(stage, codeaircraftstandtype, frontCodeAircraftStandType)
	}
}

func (codeaircraftstandtype *CodeAircraftStandType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftStandTypeDeleteCallback != nil {
		var frontCodeAircraftStandType *CodeAircraftStandType
		if front != nil {
			frontCodeAircraftStandType, _ = front.(*CodeAircraftStandType)
		}
		stage.OnAfterCodeAircraftStandTypeDeleteCallback.OnAfterDelete(stage, codeaircraftstandtype, frontCodeAircraftStandType)
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

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirportHeliportCollocationTypeCreateCallback != nil {
		stage.OnAfterCodeAirportHeliportCollocationTypeCreateCallback.OnAfterCreate(stage, codeairportheliportcollocationtype)
	}
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirportHeliportCollocationTypeUpdateCallback != nil {
		var frontCodeAirportHeliportCollocationType *CodeAirportHeliportCollocationType
		if front != nil {
			frontCodeAirportHeliportCollocationType, _ = front.(*CodeAirportHeliportCollocationType)
		}
		stage.OnAfterCodeAirportHeliportCollocationTypeUpdateCallback.OnAfterUpdate(stage, codeairportheliportcollocationtype, frontCodeAirportHeliportCollocationType)
	}
}

func (codeairportheliportcollocationtype *CodeAirportHeliportCollocationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirportHeliportCollocationTypeDeleteCallback != nil {
		var frontCodeAirportHeliportCollocationType *CodeAirportHeliportCollocationType
		if front != nil {
			frontCodeAirportHeliportCollocationType, _ = front.(*CodeAirportHeliportCollocationType)
		}
		stage.OnAfterCodeAirportHeliportCollocationTypeDeleteCallback.OnAfterDelete(stage, codeairportheliportcollocationtype, frontCodeAirportHeliportCollocationType)
	}
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirportHeliportDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeAirportHeliportDesignatorTypeCreateCallback.OnAfterCreate(stage, codeairportheliportdesignatortype)
	}
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirportHeliportDesignatorTypeUpdateCallback != nil {
		var frontCodeAirportHeliportDesignatorType *CodeAirportHeliportDesignatorType
		if front != nil {
			frontCodeAirportHeliportDesignatorType, _ = front.(*CodeAirportHeliportDesignatorType)
		}
		stage.OnAfterCodeAirportHeliportDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codeairportheliportdesignatortype, frontCodeAirportHeliportDesignatorType)
	}
}

func (codeairportheliportdesignatortype *CodeAirportHeliportDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirportHeliportDesignatorTypeDeleteCallback != nil {
		var frontCodeAirportHeliportDesignatorType *CodeAirportHeliportDesignatorType
		if front != nil {
			frontCodeAirportHeliportDesignatorType, _ = front.(*CodeAirportHeliportDesignatorType)
		}
		stage.OnAfterCodeAirportHeliportDesignatorTypeDeleteCallback.OnAfterDelete(stage, codeairportheliportdesignatortype, frontCodeAirportHeliportDesignatorType)
	}
}

func (codeairportheliporttype *CodeAirportHeliportType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirportHeliportTypeCreateCallback != nil {
		stage.OnAfterCodeAirportHeliportTypeCreateCallback.OnAfterCreate(stage, codeairportheliporttype)
	}
}

func (codeairportheliporttype *CodeAirportHeliportType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirportHeliportTypeUpdateCallback != nil {
		var frontCodeAirportHeliportType *CodeAirportHeliportType
		if front != nil {
			frontCodeAirportHeliportType, _ = front.(*CodeAirportHeliportType)
		}
		stage.OnAfterCodeAirportHeliportTypeUpdateCallback.OnAfterUpdate(stage, codeairportheliporttype, frontCodeAirportHeliportType)
	}
}

func (codeairportheliporttype *CodeAirportHeliportType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirportHeliportTypeDeleteCallback != nil {
		var frontCodeAirportHeliportType *CodeAirportHeliportType
		if front != nil {
			frontCodeAirportHeliportType, _ = front.(*CodeAirportHeliportType)
		}
		stage.OnAfterCodeAirportHeliportTypeDeleteCallback.OnAfterDelete(stage, codeairportheliporttype, frontCodeAirportHeliportType)
	}
}

func (codeairportwarningtype *CodeAirportWarningType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirportWarningTypeCreateCallback != nil {
		stage.OnAfterCodeAirportWarningTypeCreateCallback.OnAfterCreate(stage, codeairportwarningtype)
	}
}

func (codeairportwarningtype *CodeAirportWarningType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirportWarningTypeUpdateCallback != nil {
		var frontCodeAirportWarningType *CodeAirportWarningType
		if front != nil {
			frontCodeAirportWarningType, _ = front.(*CodeAirportWarningType)
		}
		stage.OnAfterCodeAirportWarningTypeUpdateCallback.OnAfterUpdate(stage, codeairportwarningtype, frontCodeAirportWarningType)
	}
}

func (codeairportwarningtype *CodeAirportWarningType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirportWarningTypeDeleteCallback != nil {
		var frontCodeAirportWarningType *CodeAirportWarningType
		if front != nil {
			frontCodeAirportWarningType, _ = front.(*CodeAirportWarningType)
		}
		stage.OnAfterCodeAirportWarningTypeDeleteCallback.OnAfterDelete(stage, codeairportwarningtype, frontCodeAirportWarningType)
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

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApproachLightingICAOTypeCreateCallback != nil {
		stage.OnAfterCodeApproachLightingICAOTypeCreateCallback.OnAfterCreate(stage, codeapproachlightingicaotype)
	}
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachLightingICAOTypeUpdateCallback != nil {
		var frontCodeApproachLightingICAOType *CodeApproachLightingICAOType
		if front != nil {
			frontCodeApproachLightingICAOType, _ = front.(*CodeApproachLightingICAOType)
		}
		stage.OnAfterCodeApproachLightingICAOTypeUpdateCallback.OnAfterUpdate(stage, codeapproachlightingicaotype, frontCodeApproachLightingICAOType)
	}
}

func (codeapproachlightingicaotype *CodeApproachLightingICAOType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachLightingICAOTypeDeleteCallback != nil {
		var frontCodeApproachLightingICAOType *CodeApproachLightingICAOType
		if front != nil {
			frontCodeApproachLightingICAOType, _ = front.(*CodeApproachLightingICAOType)
		}
		stage.OnAfterCodeApproachLightingICAOTypeDeleteCallback.OnAfterDelete(stage, codeapproachlightingicaotype, frontCodeApproachLightingICAOType)
	}
}

func (codeapproachlightingtype *CodeApproachLightingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApproachLightingTypeCreateCallback != nil {
		stage.OnAfterCodeApproachLightingTypeCreateCallback.OnAfterCreate(stage, codeapproachlightingtype)
	}
}

func (codeapproachlightingtype *CodeApproachLightingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachLightingTypeUpdateCallback != nil {
		var frontCodeApproachLightingType *CodeApproachLightingType
		if front != nil {
			frontCodeApproachLightingType, _ = front.(*CodeApproachLightingType)
		}
		stage.OnAfterCodeApproachLightingTypeUpdateCallback.OnAfterUpdate(stage, codeapproachlightingtype, frontCodeApproachLightingType)
	}
}

func (codeapproachlightingtype *CodeApproachLightingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApproachLightingTypeDeleteCallback != nil {
		var frontCodeApproachLightingType *CodeApproachLightingType
		if front != nil {
			frontCodeApproachLightingType, _ = front.(*CodeApproachLightingType)
		}
		stage.OnAfterCodeApproachLightingTypeDeleteCallback.OnAfterDelete(stage, codeapproachlightingtype, frontCodeApproachLightingType)
	}
}

func (codeapronelementtype *CodeApronElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApronElementTypeCreateCallback != nil {
		stage.OnAfterCodeApronElementTypeCreateCallback.OnAfterCreate(stage, codeapronelementtype)
	}
}

func (codeapronelementtype *CodeApronElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApronElementTypeUpdateCallback != nil {
		var frontCodeApronElementType *CodeApronElementType
		if front != nil {
			frontCodeApronElementType, _ = front.(*CodeApronElementType)
		}
		stage.OnAfterCodeApronElementTypeUpdateCallback.OnAfterUpdate(stage, codeapronelementtype, frontCodeApronElementType)
	}
}

func (codeapronelementtype *CodeApronElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApronElementTypeDeleteCallback != nil {
		var frontCodeApronElementType *CodeApronElementType
		if front != nil {
			frontCodeApronElementType, _ = front.(*CodeApronElementType)
		}
		stage.OnAfterCodeApronElementTypeDeleteCallback.OnAfterDelete(stage, codeapronelementtype, frontCodeApronElementType)
	}
}

func (codeapronsectiontype *CodeApronSectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeApronSectionTypeCreateCallback != nil {
		stage.OnAfterCodeApronSectionTypeCreateCallback.OnAfterCreate(stage, codeapronsectiontype)
	}
}

func (codeapronsectiontype *CodeApronSectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApronSectionTypeUpdateCallback != nil {
		var frontCodeApronSectionType *CodeApronSectionType
		if front != nil {
			frontCodeApronSectionType, _ = front.(*CodeApronSectionType)
		}
		stage.OnAfterCodeApronSectionTypeUpdateCallback.OnAfterUpdate(stage, codeapronsectiontype, frontCodeApronSectionType)
	}
}

func (codeapronsectiontype *CodeApronSectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeApronSectionTypeDeleteCallback != nil {
		var frontCodeApronSectionType *CodeApronSectionType
		if front != nil {
			frontCodeApronSectionType, _ = front.(*CodeApronSectionType)
		}
		stage.OnAfterCodeApronSectionTypeDeleteCallback.OnAfterDelete(stage, codeapronsectiontype, frontCodeApronSectionType)
	}
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeArrestingGearEnergyAbsorbTypeCreateCallback != nil {
		stage.OnAfterCodeArrestingGearEnergyAbsorbTypeCreateCallback.OnAfterCreate(stage, codearrestinggearenergyabsorbtype)
	}
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeArrestingGearEnergyAbsorbTypeUpdateCallback != nil {
		var frontCodeArrestingGearEnergyAbsorbType *CodeArrestingGearEnergyAbsorbType
		if front != nil {
			frontCodeArrestingGearEnergyAbsorbType, _ = front.(*CodeArrestingGearEnergyAbsorbType)
		}
		stage.OnAfterCodeArrestingGearEnergyAbsorbTypeUpdateCallback.OnAfterUpdate(stage, codearrestinggearenergyabsorbtype, frontCodeArrestingGearEnergyAbsorbType)
	}
}

func (codearrestinggearenergyabsorbtype *CodeArrestingGearEnergyAbsorbType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeArrestingGearEnergyAbsorbTypeDeleteCallback != nil {
		var frontCodeArrestingGearEnergyAbsorbType *CodeArrestingGearEnergyAbsorbType
		if front != nil {
			frontCodeArrestingGearEnergyAbsorbType, _ = front.(*CodeArrestingGearEnergyAbsorbType)
		}
		stage.OnAfterCodeArrestingGearEnergyAbsorbTypeDeleteCallback.OnAfterDelete(stage, codearrestinggearenergyabsorbtype, frontCodeArrestingGearEnergyAbsorbType)
	}
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeArrestingGearEngageDeviceTypeCreateCallback != nil {
		stage.OnAfterCodeArrestingGearEngageDeviceTypeCreateCallback.OnAfterCreate(stage, codearrestinggearengagedevicetype)
	}
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeArrestingGearEngageDeviceTypeUpdateCallback != nil {
		var frontCodeArrestingGearEngageDeviceType *CodeArrestingGearEngageDeviceType
		if front != nil {
			frontCodeArrestingGearEngageDeviceType, _ = front.(*CodeArrestingGearEngageDeviceType)
		}
		stage.OnAfterCodeArrestingGearEngageDeviceTypeUpdateCallback.OnAfterUpdate(stage, codearrestinggearengagedevicetype, frontCodeArrestingGearEngageDeviceType)
	}
}

func (codearrestinggearengagedevicetype *CodeArrestingGearEngageDeviceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeArrestingGearEngageDeviceTypeDeleteCallback != nil {
		var frontCodeArrestingGearEngageDeviceType *CodeArrestingGearEngageDeviceType
		if front != nil {
			frontCodeArrestingGearEngageDeviceType, _ = front.(*CodeArrestingGearEngageDeviceType)
		}
		stage.OnAfterCodeArrestingGearEngageDeviceTypeDeleteCallback.OnAfterDelete(stage, codearrestinggearengagedevicetype, frontCodeArrestingGearEngageDeviceType)
	}
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAuthorityRoleTypeCreateCallback != nil {
		stage.OnAfterCodeAuthorityRoleTypeCreateCallback.OnAfterCreate(stage, codeauthorityroletype)
	}
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuthorityRoleTypeUpdateCallback != nil {
		var frontCodeAuthorityRoleType *CodeAuthorityRoleType
		if front != nil {
			frontCodeAuthorityRoleType, _ = front.(*CodeAuthorityRoleType)
		}
		stage.OnAfterCodeAuthorityRoleTypeUpdateCallback.OnAfterUpdate(stage, codeauthorityroletype, frontCodeAuthorityRoleType)
	}
}

func (codeauthorityroletype *CodeAuthorityRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuthorityRoleTypeDeleteCallback != nil {
		var frontCodeAuthorityRoleType *CodeAuthorityRoleType
		if front != nil {
			frontCodeAuthorityRoleType, _ = front.(*CodeAuthorityRoleType)
		}
		stage.OnAfterCodeAuthorityRoleTypeDeleteCallback.OnAfterDelete(stage, codeauthorityroletype, frontCodeAuthorityRoleType)
	}
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeBuoyDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeBuoyDesignatorTypeCreateCallback.OnAfterCreate(stage, codebuoydesignatortype)
	}
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeBuoyDesignatorTypeUpdateCallback != nil {
		var frontCodeBuoyDesignatorType *CodeBuoyDesignatorType
		if front != nil {
			frontCodeBuoyDesignatorType, _ = front.(*CodeBuoyDesignatorType)
		}
		stage.OnAfterCodeBuoyDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codebuoydesignatortype, frontCodeBuoyDesignatorType)
	}
}

func (codebuoydesignatortype *CodeBuoyDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeBuoyDesignatorTypeDeleteCallback != nil {
		var frontCodeBuoyDesignatorType *CodeBuoyDesignatorType
		if front != nil {
			frontCodeBuoyDesignatorType, _ = front.(*CodeBuoyDesignatorType)
		}
		stage.OnAfterCodeBuoyDesignatorTypeDeleteCallback.OnAfterDelete(stage, codebuoydesignatortype, frontCodeBuoyDesignatorType)
	}
}

func (codebuoytype *CodeBuoyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeBuoyTypeCreateCallback != nil {
		stage.OnAfterCodeBuoyTypeCreateCallback.OnAfterCreate(stage, codebuoytype)
	}
}

func (codebuoytype *CodeBuoyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeBuoyTypeUpdateCallback != nil {
		var frontCodeBuoyType *CodeBuoyType
		if front != nil {
			frontCodeBuoyType, _ = front.(*CodeBuoyType)
		}
		stage.OnAfterCodeBuoyTypeUpdateCallback.OnAfterUpdate(stage, codebuoytype, frontCodeBuoyType)
	}
}

func (codebuoytype *CodeBuoyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeBuoyTypeDeleteCallback != nil {
		var frontCodeBuoyType *CodeBuoyType
		if front != nil {
			frontCodeBuoyType, _ = front.(*CodeBuoyType)
		}
		stage.OnAfterCodeBuoyTypeDeleteCallback.OnAfterDelete(stage, codebuoytype, frontCodeBuoyType)
	}
}

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

func (codecontaminationtype *CodeContaminationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeContaminationTypeCreateCallback != nil {
		stage.OnAfterCodeContaminationTypeCreateCallback.OnAfterCreate(stage, codecontaminationtype)
	}
}

func (codecontaminationtype *CodeContaminationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeContaminationTypeUpdateCallback != nil {
		var frontCodeContaminationType *CodeContaminationType
		if front != nil {
			frontCodeContaminationType, _ = front.(*CodeContaminationType)
		}
		stage.OnAfterCodeContaminationTypeUpdateCallback.OnAfterUpdate(stage, codecontaminationtype, frontCodeContaminationType)
	}
}

func (codecontaminationtype *CodeContaminationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeContaminationTypeDeleteCallback != nil {
		var frontCodeContaminationType *CodeContaminationType
		if front != nil {
			frontCodeContaminationType, _ = front.(*CodeContaminationType)
		}
		stage.OnAfterCodeContaminationTypeDeleteCallback.OnAfterDelete(stage, codecontaminationtype, frontCodeContaminationType)
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

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDeclaredDistanceTypeCreateCallback != nil {
		stage.OnAfterCodeDeclaredDistanceTypeCreateCallback.OnAfterCreate(stage, codedeclareddistancetype)
	}
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDeclaredDistanceTypeUpdateCallback != nil {
		var frontCodeDeclaredDistanceType *CodeDeclaredDistanceType
		if front != nil {
			frontCodeDeclaredDistanceType, _ = front.(*CodeDeclaredDistanceType)
		}
		stage.OnAfterCodeDeclaredDistanceTypeUpdateCallback.OnAfterUpdate(stage, codedeclareddistancetype, frontCodeDeclaredDistanceType)
	}
}

func (codedeclareddistancetype *CodeDeclaredDistanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDeclaredDistanceTypeDeleteCallback != nil {
		var frontCodeDeclaredDistanceType *CodeDeclaredDistanceType
		if front != nil {
			frontCodeDeclaredDistanceType, _ = front.(*CodeDeclaredDistanceType)
		}
		stage.OnAfterCodeDeclaredDistanceTypeDeleteCallback.OnAfterDelete(stage, codedeclareddistancetype, frontCodeDeclaredDistanceType)
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

func (codefrictiondevicetype *CodeFrictionDeviceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFrictionDeviceTypeCreateCallback != nil {
		stage.OnAfterCodeFrictionDeviceTypeCreateCallback.OnAfterCreate(stage, codefrictiondevicetype)
	}
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFrictionDeviceTypeUpdateCallback != nil {
		var frontCodeFrictionDeviceType *CodeFrictionDeviceType
		if front != nil {
			frontCodeFrictionDeviceType, _ = front.(*CodeFrictionDeviceType)
		}
		stage.OnAfterCodeFrictionDeviceTypeUpdateCallback.OnAfterUpdate(stage, codefrictiondevicetype, frontCodeFrictionDeviceType)
	}
}

func (codefrictiondevicetype *CodeFrictionDeviceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFrictionDeviceTypeDeleteCallback != nil {
		var frontCodeFrictionDeviceType *CodeFrictionDeviceType
		if front != nil {
			frontCodeFrictionDeviceType, _ = front.(*CodeFrictionDeviceType)
		}
		stage.OnAfterCodeFrictionDeviceTypeDeleteCallback.OnAfterDelete(stage, codefrictiondevicetype, frontCodeFrictionDeviceType)
	}
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFrictionEstimateTypeCreateCallback != nil {
		stage.OnAfterCodeFrictionEstimateTypeCreateCallback.OnAfterCreate(stage, codefrictionestimatetype)
	}
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFrictionEstimateTypeUpdateCallback != nil {
		var frontCodeFrictionEstimateType *CodeFrictionEstimateType
		if front != nil {
			frontCodeFrictionEstimateType, _ = front.(*CodeFrictionEstimateType)
		}
		stage.OnAfterCodeFrictionEstimateTypeUpdateCallback.OnAfterUpdate(stage, codefrictionestimatetype, frontCodeFrictionEstimateType)
	}
}

func (codefrictionestimatetype *CodeFrictionEstimateType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFrictionEstimateTypeDeleteCallback != nil {
		var frontCodeFrictionEstimateType *CodeFrictionEstimateType
		if front != nil {
			frontCodeFrictionEstimateType, _ = front.(*CodeFrictionEstimateType)
		}
		stage.OnAfterCodeFrictionEstimateTypeDeleteCallback.OnAfterDelete(stage, codefrictionestimatetype, frontCodeFrictionEstimateType)
	}
}

func (codegradeseparationtype *CodeGradeSeparationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeGradeSeparationTypeCreateCallback != nil {
		stage.OnAfterCodeGradeSeparationTypeCreateCallback.OnAfterCreate(stage, codegradeseparationtype)
	}
}

func (codegradeseparationtype *CodeGradeSeparationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeGradeSeparationTypeUpdateCallback != nil {
		var frontCodeGradeSeparationType *CodeGradeSeparationType
		if front != nil {
			frontCodeGradeSeparationType, _ = front.(*CodeGradeSeparationType)
		}
		stage.OnAfterCodeGradeSeparationTypeUpdateCallback.OnAfterUpdate(stage, codegradeseparationtype, frontCodeGradeSeparationType)
	}
}

func (codegradeseparationtype *CodeGradeSeparationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeGradeSeparationTypeDeleteCallback != nil {
		var frontCodeGradeSeparationType *CodeGradeSeparationType
		if front != nil {
			frontCodeGradeSeparationType, _ = front.(*CodeGradeSeparationType)
		}
		stage.OnAfterCodeGradeSeparationTypeDeleteCallback.OnAfterDelete(stage, codegradeseparationtype, frontCodeGradeSeparationType)
	}
}

func (codeguidancelinetype *CodeGuidanceLineType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeGuidanceLineTypeCreateCallback != nil {
		stage.OnAfterCodeGuidanceLineTypeCreateCallback.OnAfterCreate(stage, codeguidancelinetype)
	}
}

func (codeguidancelinetype *CodeGuidanceLineType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeGuidanceLineTypeUpdateCallback != nil {
		var frontCodeGuidanceLineType *CodeGuidanceLineType
		if front != nil {
			frontCodeGuidanceLineType, _ = front.(*CodeGuidanceLineType)
		}
		stage.OnAfterCodeGuidanceLineTypeUpdateCallback.OnAfterUpdate(stage, codeguidancelinetype, frontCodeGuidanceLineType)
	}
}

func (codeguidancelinetype *CodeGuidanceLineType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeGuidanceLineTypeDeleteCallback != nil {
		var frontCodeGuidanceLineType *CodeGuidanceLineType
		if front != nil {
			frontCodeGuidanceLineType, _ = front.(*CodeGuidanceLineType)
		}
		stage.OnAfterCodeGuidanceLineTypeDeleteCallback.OnAfterDelete(stage, codeguidancelinetype, frontCodeGuidanceLineType)
	}
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeHelicopterPerformanceTypeCreateCallback != nil {
		stage.OnAfterCodeHelicopterPerformanceTypeCreateCallback.OnAfterCreate(stage, codehelicopterperformancetype)
	}
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHelicopterPerformanceTypeUpdateCallback != nil {
		var frontCodeHelicopterPerformanceType *CodeHelicopterPerformanceType
		if front != nil {
			frontCodeHelicopterPerformanceType, _ = front.(*CodeHelicopterPerformanceType)
		}
		stage.OnAfterCodeHelicopterPerformanceTypeUpdateCallback.OnAfterUpdate(stage, codehelicopterperformancetype, frontCodeHelicopterPerformanceType)
	}
}

func (codehelicopterperformancetype *CodeHelicopterPerformanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHelicopterPerformanceTypeDeleteCallback != nil {
		var frontCodeHelicopterPerformanceType *CodeHelicopterPerformanceType
		if front != nil {
			frontCodeHelicopterPerformanceType, _ = front.(*CodeHelicopterPerformanceType)
		}
		stage.OnAfterCodeHelicopterPerformanceTypeDeleteCallback.OnAfterDelete(stage, codehelicopterperformancetype, frontCodeHelicopterPerformanceType)
	}
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeHoldingCategoryTypeCreateCallback != nil {
		stage.OnAfterCodeHoldingCategoryTypeCreateCallback.OnAfterCreate(stage, codeholdingcategorytype)
	}
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHoldingCategoryTypeUpdateCallback != nil {
		var frontCodeHoldingCategoryType *CodeHoldingCategoryType
		if front != nil {
			frontCodeHoldingCategoryType, _ = front.(*CodeHoldingCategoryType)
		}
		stage.OnAfterCodeHoldingCategoryTypeUpdateCallback.OnAfterUpdate(stage, codeholdingcategorytype, frontCodeHoldingCategoryType)
	}
}

func (codeholdingcategorytype *CodeHoldingCategoryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeHoldingCategoryTypeDeleteCallback != nil {
		var frontCodeHoldingCategoryType *CodeHoldingCategoryType
		if front != nil {
			frontCodeHoldingCategoryType, _ = front.(*CodeHoldingCategoryType)
		}
		stage.OnAfterCodeHoldingCategoryTypeDeleteCallback.OnAfterDelete(stage, codeholdingcategorytype, frontCodeHoldingCategoryType)
	}
}

func (codeiatatype *CodeIATAType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeIATATypeCreateCallback != nil {
		stage.OnAfterCodeIATATypeCreateCallback.OnAfterCreate(stage, codeiatatype)
	}
}

func (codeiatatype *CodeIATAType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIATATypeUpdateCallback != nil {
		var frontCodeIATAType *CodeIATAType
		if front != nil {
			frontCodeIATAType, _ = front.(*CodeIATAType)
		}
		stage.OnAfterCodeIATATypeUpdateCallback.OnAfterUpdate(stage, codeiatatype, frontCodeIATAType)
	}
}

func (codeiatatype *CodeIATAType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIATATypeDeleteCallback != nil {
		var frontCodeIATAType *CodeIATAType
		if front != nil {
			frontCodeIATAType, _ = front.(*CodeIATAType)
		}
		stage.OnAfterCodeIATATypeDeleteCallback.OnAfterDelete(stage, codeiatatype, frontCodeIATAType)
	}
}

func (codeicaotype *CodeICAOType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeICAOTypeCreateCallback != nil {
		stage.OnAfterCodeICAOTypeCreateCallback.OnAfterCreate(stage, codeicaotype)
	}
}

func (codeicaotype *CodeICAOType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeICAOTypeUpdateCallback != nil {
		var frontCodeICAOType *CodeICAOType
		if front != nil {
			frontCodeICAOType, _ = front.(*CodeICAOType)
		}
		stage.OnAfterCodeICAOTypeUpdateCallback.OnAfterUpdate(stage, codeicaotype, frontCodeICAOType)
	}
}

func (codeicaotype *CodeICAOType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeICAOTypeDeleteCallback != nil {
		var frontCodeICAOType *CodeICAOType
		if front != nil {
			frontCodeICAOType, _ = front.(*CodeICAOType)
		}
		stage.OnAfterCodeICAOTypeDeleteCallback.OnAfterDelete(stage, codeicaotype, frontCodeICAOType)
	}
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLightHoldingPositionTypeCreateCallback != nil {
		stage.OnAfterCodeLightHoldingPositionTypeCreateCallback.OnAfterCreate(stage, codelightholdingpositiontype)
	}
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightHoldingPositionTypeUpdateCallback != nil {
		var frontCodeLightHoldingPositionType *CodeLightHoldingPositionType
		if front != nil {
			frontCodeLightHoldingPositionType, _ = front.(*CodeLightHoldingPositionType)
		}
		stage.OnAfterCodeLightHoldingPositionTypeUpdateCallback.OnAfterUpdate(stage, codelightholdingpositiontype, frontCodeLightHoldingPositionType)
	}
}

func (codelightholdingpositiontype *CodeLightHoldingPositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightHoldingPositionTypeDeleteCallback != nil {
		var frontCodeLightHoldingPositionType *CodeLightHoldingPositionType
		if front != nil {
			frontCodeLightHoldingPositionType, _ = front.(*CodeLightHoldingPositionType)
		}
		stage.OnAfterCodeLightHoldingPositionTypeDeleteCallback.OnAfterDelete(stage, codelightholdingpositiontype, frontCodeLightHoldingPositionType)
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

func (codelightingjartype *CodeLightingJARType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLightingJARTypeCreateCallback != nil {
		stage.OnAfterCodeLightingJARTypeCreateCallback.OnAfterCreate(stage, codelightingjartype)
	}
}

func (codelightingjartype *CodeLightingJARType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightingJARTypeUpdateCallback != nil {
		var frontCodeLightingJARType *CodeLightingJARType
		if front != nil {
			frontCodeLightingJARType, _ = front.(*CodeLightingJARType)
		}
		stage.OnAfterCodeLightingJARTypeUpdateCallback.OnAfterUpdate(stage, codelightingjartype, frontCodeLightingJARType)
	}
}

func (codelightingjartype *CodeLightingJARType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLightingJARTypeDeleteCallback != nil {
		var frontCodeLightingJARType *CodeLightingJARType
		if front != nil {
			frontCodeLightingJARType, _ = front.(*CodeLightingJARType)
		}
		stage.OnAfterCodeLightingJARTypeDeleteCallback.OnAfterDelete(stage, codelightingjartype, frontCodeLightingJARType)
	}
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLoadingBridgeTypeCreateCallback != nil {
		stage.OnAfterCodeLoadingBridgeTypeCreateCallback.OnAfterCreate(stage, codeloadingbridgetype)
	}
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLoadingBridgeTypeUpdateCallback != nil {
		var frontCodeLoadingBridgeType *CodeLoadingBridgeType
		if front != nil {
			frontCodeLoadingBridgeType, _ = front.(*CodeLoadingBridgeType)
		}
		stage.OnAfterCodeLoadingBridgeTypeUpdateCallback.OnAfterUpdate(stage, codeloadingbridgetype, frontCodeLoadingBridgeType)
	}
}

func (codeloadingbridgetype *CodeLoadingBridgeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLoadingBridgeTypeDeleteCallback != nil {
		var frontCodeLoadingBridgeType *CodeLoadingBridgeType
		if front != nil {
			frontCodeLoadingBridgeType, _ = front.(*CodeLoadingBridgeType)
		}
		stage.OnAfterCodeLoadingBridgeTypeDeleteCallback.OnAfterDelete(stage, codeloadingbridgetype, frontCodeLoadingBridgeType)
	}
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLogicalOperatorTypeCreateCallback != nil {
		stage.OnAfterCodeLogicalOperatorTypeCreateCallback.OnAfterCreate(stage, codelogicaloperatortype)
	}
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLogicalOperatorTypeUpdateCallback != nil {
		var frontCodeLogicalOperatorType *CodeLogicalOperatorType
		if front != nil {
			frontCodeLogicalOperatorType, _ = front.(*CodeLogicalOperatorType)
		}
		stage.OnAfterCodeLogicalOperatorTypeUpdateCallback.OnAfterUpdate(stage, codelogicaloperatortype, frontCodeLogicalOperatorType)
	}
}

func (codelogicaloperatortype *CodeLogicalOperatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLogicalOperatorTypeDeleteCallback != nil {
		var frontCodeLogicalOperatorType *CodeLogicalOperatorType
		if front != nil {
			frontCodeLogicalOperatorType, _ = front.(*CodeLogicalOperatorType)
		}
		stage.OnAfterCodeLogicalOperatorTypeDeleteCallback.OnAfterDelete(stage, codelogicaloperatortype, frontCodeLogicalOperatorType)
	}
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMarkingConditionTypeCreateCallback != nil {
		stage.OnAfterCodeMarkingConditionTypeCreateCallback.OnAfterCreate(stage, codemarkingconditiontype)
	}
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMarkingConditionTypeUpdateCallback != nil {
		var frontCodeMarkingConditionType *CodeMarkingConditionType
		if front != nil {
			frontCodeMarkingConditionType, _ = front.(*CodeMarkingConditionType)
		}
		stage.OnAfterCodeMarkingConditionTypeUpdateCallback.OnAfterUpdate(stage, codemarkingconditiontype, frontCodeMarkingConditionType)
	}
}

func (codemarkingconditiontype *CodeMarkingConditionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMarkingConditionTypeDeleteCallback != nil {
		var frontCodeMarkingConditionType *CodeMarkingConditionType
		if front != nil {
			frontCodeMarkingConditionType, _ = front.(*CodeMarkingConditionType)
		}
		stage.OnAfterCodeMarkingConditionTypeDeleteCallback.OnAfterDelete(stage, codemarkingconditiontype, frontCodeMarkingConditionType)
	}
}

func (codemarkingstyletype *CodeMarkingStyleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMarkingStyleTypeCreateCallback != nil {
		stage.OnAfterCodeMarkingStyleTypeCreateCallback.OnAfterCreate(stage, codemarkingstyletype)
	}
}

func (codemarkingstyletype *CodeMarkingStyleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMarkingStyleTypeUpdateCallback != nil {
		var frontCodeMarkingStyleType *CodeMarkingStyleType
		if front != nil {
			frontCodeMarkingStyleType, _ = front.(*CodeMarkingStyleType)
		}
		stage.OnAfterCodeMarkingStyleTypeUpdateCallback.OnAfterUpdate(stage, codemarkingstyletype, frontCodeMarkingStyleType)
	}
}

func (codemarkingstyletype *CodeMarkingStyleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMarkingStyleTypeDeleteCallback != nil {
		var frontCodeMarkingStyleType *CodeMarkingStyleType
		if front != nil {
			frontCodeMarkingStyleType, _ = front.(*CodeMarkingStyleType)
		}
		stage.OnAfterCodeMarkingStyleTypeDeleteCallback.OnAfterDelete(stage, codemarkingstyletype, frontCodeMarkingStyleType)
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

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMilitaryOperationsTypeCreateCallback != nil {
		stage.OnAfterCodeMilitaryOperationsTypeCreateCallback.OnAfterCreate(stage, codemilitaryoperationstype)
	}
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryOperationsTypeUpdateCallback != nil {
		var frontCodeMilitaryOperationsType *CodeMilitaryOperationsType
		if front != nil {
			frontCodeMilitaryOperationsType, _ = front.(*CodeMilitaryOperationsType)
		}
		stage.OnAfterCodeMilitaryOperationsTypeUpdateCallback.OnAfterUpdate(stage, codemilitaryoperationstype, frontCodeMilitaryOperationsType)
	}
}

func (codemilitaryoperationstype *CodeMilitaryOperationsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMilitaryOperationsTypeDeleteCallback != nil {
		var frontCodeMilitaryOperationsType *CodeMilitaryOperationsType
		if front != nil {
			frontCodeMilitaryOperationsType, _ = front.(*CodeMilitaryOperationsType)
		}
		stage.OnAfterCodeMilitaryOperationsTypeDeleteCallback.OnAfterDelete(stage, codemilitaryoperationstype, frontCodeMilitaryOperationsType)
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

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOperationAirportHeliportTypeCreateCallback != nil {
		stage.OnAfterCodeOperationAirportHeliportTypeCreateCallback.OnAfterCreate(stage, codeoperationairportheliporttype)
	}
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOperationAirportHeliportTypeUpdateCallback != nil {
		var frontCodeOperationAirportHeliportType *CodeOperationAirportHeliportType
		if front != nil {
			frontCodeOperationAirportHeliportType, _ = front.(*CodeOperationAirportHeliportType)
		}
		stage.OnAfterCodeOperationAirportHeliportTypeUpdateCallback.OnAfterUpdate(stage, codeoperationairportheliporttype, frontCodeOperationAirportHeliportType)
	}
}

func (codeoperationairportheliporttype *CodeOperationAirportHeliportType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOperationAirportHeliportTypeDeleteCallback != nil {
		var frontCodeOperationAirportHeliportType *CodeOperationAirportHeliportType
		if front != nil {
			frontCodeOperationAirportHeliportType, _ = front.(*CodeOperationAirportHeliportType)
		}
		stage.OnAfterCodeOperationAirportHeliportTypeDeleteCallback.OnAfterDelete(stage, codeoperationairportheliporttype, frontCodeOperationAirportHeliportType)
	}
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOperationManoeuvringAreaTypeCreateCallback != nil {
		stage.OnAfterCodeOperationManoeuvringAreaTypeCreateCallback.OnAfterCreate(stage, codeoperationmanoeuvringareatype)
	}
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOperationManoeuvringAreaTypeUpdateCallback != nil {
		var frontCodeOperationManoeuvringAreaType *CodeOperationManoeuvringAreaType
		if front != nil {
			frontCodeOperationManoeuvringAreaType, _ = front.(*CodeOperationManoeuvringAreaType)
		}
		stage.OnAfterCodeOperationManoeuvringAreaTypeUpdateCallback.OnAfterUpdate(stage, codeoperationmanoeuvringareatype, frontCodeOperationManoeuvringAreaType)
	}
}

func (codeoperationmanoeuvringareatype *CodeOperationManoeuvringAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOperationManoeuvringAreaTypeDeleteCallback != nil {
		var frontCodeOperationManoeuvringAreaType *CodeOperationManoeuvringAreaType
		if front != nil {
			frontCodeOperationManoeuvringAreaType, _ = front.(*CodeOperationManoeuvringAreaType)
		}
		stage.OnAfterCodeOperationManoeuvringAreaTypeDeleteCallback.OnAfterDelete(stage, codeoperationmanoeuvringareatype, frontCodeOperationManoeuvringAreaType)
	}
}

func (codepcnmethodtype *CodePCNMethodType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePCNMethodTypeCreateCallback != nil {
		stage.OnAfterCodePCNMethodTypeCreateCallback.OnAfterCreate(stage, codepcnmethodtype)
	}
}

func (codepcnmethodtype *CodePCNMethodType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePCNMethodTypeUpdateCallback != nil {
		var frontCodePCNMethodType *CodePCNMethodType
		if front != nil {
			frontCodePCNMethodType, _ = front.(*CodePCNMethodType)
		}
		stage.OnAfterCodePCNMethodTypeUpdateCallback.OnAfterUpdate(stage, codepcnmethodtype, frontCodePCNMethodType)
	}
}

func (codepcnmethodtype *CodePCNMethodType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePCNMethodTypeDeleteCallback != nil {
		var frontCodePCNMethodType *CodePCNMethodType
		if front != nil {
			frontCodePCNMethodType, _ = front.(*CodePCNMethodType)
		}
		stage.OnAfterCodePCNMethodTypeDeleteCallback.OnAfterDelete(stage, codepcnmethodtype, frontCodePCNMethodType)
	}
}

func (codepcnpavementtype *CodePCNPavementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePCNPavementTypeCreateCallback != nil {
		stage.OnAfterCodePCNPavementTypeCreateCallback.OnAfterCreate(stage, codepcnpavementtype)
	}
}

func (codepcnpavementtype *CodePCNPavementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePCNPavementTypeUpdateCallback != nil {
		var frontCodePCNPavementType *CodePCNPavementType
		if front != nil {
			frontCodePCNPavementType, _ = front.(*CodePCNPavementType)
		}
		stage.OnAfterCodePCNPavementTypeUpdateCallback.OnAfterUpdate(stage, codepcnpavementtype, frontCodePCNPavementType)
	}
}

func (codepcnpavementtype *CodePCNPavementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePCNPavementTypeDeleteCallback != nil {
		var frontCodePCNPavementType *CodePCNPavementType
		if front != nil {
			frontCodePCNPavementType, _ = front.(*CodePCNPavementType)
		}
		stage.OnAfterCodePCNPavementTypeDeleteCallback.OnAfterDelete(stage, codepcnpavementtype, frontCodePCNPavementType)
	}
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePCNSubgradeTypeCreateCallback != nil {
		stage.OnAfterCodePCNSubgradeTypeCreateCallback.OnAfterCreate(stage, codepcnsubgradetype)
	}
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePCNSubgradeTypeUpdateCallback != nil {
		var frontCodePCNSubgradeType *CodePCNSubgradeType
		if front != nil {
			frontCodePCNSubgradeType, _ = front.(*CodePCNSubgradeType)
		}
		stage.OnAfterCodePCNSubgradeTypeUpdateCallback.OnAfterUpdate(stage, codepcnsubgradetype, frontCodePCNSubgradeType)
	}
}

func (codepcnsubgradetype *CodePCNSubgradeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePCNSubgradeTypeDeleteCallback != nil {
		var frontCodePCNSubgradeType *CodePCNSubgradeType
		if front != nil {
			frontCodePCNSubgradeType, _ = front.(*CodePCNSubgradeType)
		}
		stage.OnAfterCodePCNSubgradeTypeDeleteCallback.OnAfterDelete(stage, codepcnsubgradetype, frontCodePCNSubgradeType)
	}
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePCNTyrePressureTypeCreateCallback != nil {
		stage.OnAfterCodePCNTyrePressureTypeCreateCallback.OnAfterCreate(stage, codepcntyrepressuretype)
	}
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePCNTyrePressureTypeUpdateCallback != nil {
		var frontCodePCNTyrePressureType *CodePCNTyrePressureType
		if front != nil {
			frontCodePCNTyrePressureType, _ = front.(*CodePCNTyrePressureType)
		}
		stage.OnAfterCodePCNTyrePressureTypeUpdateCallback.OnAfterUpdate(stage, codepcntyrepressuretype, frontCodePCNTyrePressureType)
	}
}

func (codepcntyrepressuretype *CodePCNTyrePressureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePCNTyrePressureTypeDeleteCallback != nil {
		var frontCodePCNTyrePressureType *CodePCNTyrePressureType
		if front != nil {
			frontCodePCNTyrePressureType, _ = front.(*CodePCNTyrePressureType)
		}
		stage.OnAfterCodePCNTyrePressureTypeDeleteCallback.OnAfterDelete(stage, codepcntyrepressuretype, frontCodePCNTyrePressureType)
	}
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeProtectAreaSectionTypeCreateCallback != nil {
		stage.OnAfterCodeProtectAreaSectionTypeCreateCallback.OnAfterCreate(stage, codeprotectareasectiontype)
	}
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProtectAreaSectionTypeUpdateCallback != nil {
		var frontCodeProtectAreaSectionType *CodeProtectAreaSectionType
		if front != nil {
			frontCodeProtectAreaSectionType, _ = front.(*CodeProtectAreaSectionType)
		}
		stage.OnAfterCodeProtectAreaSectionTypeUpdateCallback.OnAfterUpdate(stage, codeprotectareasectiontype, frontCodeProtectAreaSectionType)
	}
}

func (codeprotectareasectiontype *CodeProtectAreaSectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeProtectAreaSectionTypeDeleteCallback != nil {
		var frontCodeProtectAreaSectionType *CodeProtectAreaSectionType
		if front != nil {
			frontCodeProtectAreaSectionType, _ = front.(*CodeProtectAreaSectionType)
		}
		stage.OnAfterCodeProtectAreaSectionTypeDeleteCallback.OnAfterDelete(stage, codeprotectareasectiontype, frontCodeProtectAreaSectionType)
	}
}

func (codervrreadingtype *CodeRVRReadingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRVRReadingTypeCreateCallback != nil {
		stage.OnAfterCodeRVRReadingTypeCreateCallback.OnAfterCreate(stage, codervrreadingtype)
	}
}

func (codervrreadingtype *CodeRVRReadingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRVRReadingTypeUpdateCallback != nil {
		var frontCodeRVRReadingType *CodeRVRReadingType
		if front != nil {
			frontCodeRVRReadingType, _ = front.(*CodeRVRReadingType)
		}
		stage.OnAfterCodeRVRReadingTypeUpdateCallback.OnAfterUpdate(stage, codervrreadingtype, frontCodeRVRReadingType)
	}
}

func (codervrreadingtype *CodeRVRReadingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRVRReadingTypeDeleteCallback != nil {
		var frontCodeRVRReadingType *CodeRVRReadingType
		if front != nil {
			frontCodeRVRReadingType, _ = front.(*CodeRVRReadingType)
		}
		stage.OnAfterCodeRVRReadingTypeDeleteCallback.OnAfterDelete(stage, codervrreadingtype, frontCodeRVRReadingType)
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

func (coderoadtype *CodeRoadType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRoadTypeCreateCallback != nil {
		stage.OnAfterCodeRoadTypeCreateCallback.OnAfterCreate(stage, coderoadtype)
	}
}

func (coderoadtype *CodeRoadType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRoadTypeUpdateCallback != nil {
		var frontCodeRoadType *CodeRoadType
		if front != nil {
			frontCodeRoadType, _ = front.(*CodeRoadType)
		}
		stage.OnAfterCodeRoadTypeUpdateCallback.OnAfterUpdate(stage, coderoadtype, frontCodeRoadType)
	}
}

func (coderoadtype *CodeRoadType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRoadTypeDeleteCallback != nil {
		var frontCodeRoadType *CodeRoadType
		if front != nil {
			frontCodeRoadType, _ = front.(*CodeRoadType)
		}
		stage.OnAfterCodeRoadTypeDeleteCallback.OnAfterDelete(stage, coderoadtype, frontCodeRoadType)
	}
}

func (coderunwayelementtype *CodeRunwayElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRunwayElementTypeCreateCallback != nil {
		stage.OnAfterCodeRunwayElementTypeCreateCallback.OnAfterCreate(stage, coderunwayelementtype)
	}
}

func (coderunwayelementtype *CodeRunwayElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayElementTypeUpdateCallback != nil {
		var frontCodeRunwayElementType *CodeRunwayElementType
		if front != nil {
			frontCodeRunwayElementType, _ = front.(*CodeRunwayElementType)
		}
		stage.OnAfterCodeRunwayElementTypeUpdateCallback.OnAfterUpdate(stage, coderunwayelementtype, frontCodeRunwayElementType)
	}
}

func (coderunwayelementtype *CodeRunwayElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayElementTypeDeleteCallback != nil {
		var frontCodeRunwayElementType *CodeRunwayElementType
		if front != nil {
			frontCodeRunwayElementType, _ = front.(*CodeRunwayElementType)
		}
		stage.OnAfterCodeRunwayElementTypeDeleteCallback.OnAfterDelete(stage, coderunwayelementtype, frontCodeRunwayElementType)
	}
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRunwayMarkingTypeCreateCallback != nil {
		stage.OnAfterCodeRunwayMarkingTypeCreateCallback.OnAfterCreate(stage, coderunwaymarkingtype)
	}
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayMarkingTypeUpdateCallback != nil {
		var frontCodeRunwayMarkingType *CodeRunwayMarkingType
		if front != nil {
			frontCodeRunwayMarkingType, _ = front.(*CodeRunwayMarkingType)
		}
		stage.OnAfterCodeRunwayMarkingTypeUpdateCallback.OnAfterUpdate(stage, coderunwaymarkingtype, frontCodeRunwayMarkingType)
	}
}

func (coderunwaymarkingtype *CodeRunwayMarkingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayMarkingTypeDeleteCallback != nil {
		var frontCodeRunwayMarkingType *CodeRunwayMarkingType
		if front != nil {
			frontCodeRunwayMarkingType, _ = front.(*CodeRunwayMarkingType)
		}
		stage.OnAfterCodeRunwayMarkingTypeDeleteCallback.OnAfterDelete(stage, coderunwaymarkingtype, frontCodeRunwayMarkingType)
	}
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRunwayPointRoleTypeCreateCallback != nil {
		stage.OnAfterCodeRunwayPointRoleTypeCreateCallback.OnAfterCreate(stage, coderunwaypointroletype)
	}
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayPointRoleTypeUpdateCallback != nil {
		var frontCodeRunwayPointRoleType *CodeRunwayPointRoleType
		if front != nil {
			frontCodeRunwayPointRoleType, _ = front.(*CodeRunwayPointRoleType)
		}
		stage.OnAfterCodeRunwayPointRoleTypeUpdateCallback.OnAfterUpdate(stage, coderunwaypointroletype, frontCodeRunwayPointRoleType)
	}
}

func (coderunwaypointroletype *CodeRunwayPointRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayPointRoleTypeDeleteCallback != nil {
		var frontCodeRunwayPointRoleType *CodeRunwayPointRoleType
		if front != nil {
			frontCodeRunwayPointRoleType, _ = front.(*CodeRunwayPointRoleType)
		}
		stage.OnAfterCodeRunwayPointRoleTypeDeleteCallback.OnAfterDelete(stage, coderunwaypointroletype, frontCodeRunwayPointRoleType)
	}
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRunwayProtectionAreaTypeCreateCallback != nil {
		stage.OnAfterCodeRunwayProtectionAreaTypeCreateCallback.OnAfterCreate(stage, coderunwayprotectionareatype)
	}
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayProtectionAreaTypeUpdateCallback != nil {
		var frontCodeRunwayProtectionAreaType *CodeRunwayProtectionAreaType
		if front != nil {
			frontCodeRunwayProtectionAreaType, _ = front.(*CodeRunwayProtectionAreaType)
		}
		stage.OnAfterCodeRunwayProtectionAreaTypeUpdateCallback.OnAfterUpdate(stage, coderunwayprotectionareatype, frontCodeRunwayProtectionAreaType)
	}
}

func (coderunwayprotectionareatype *CodeRunwayProtectionAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayProtectionAreaTypeDeleteCallback != nil {
		var frontCodeRunwayProtectionAreaType *CodeRunwayProtectionAreaType
		if front != nil {
			frontCodeRunwayProtectionAreaType, _ = front.(*CodeRunwayProtectionAreaType)
		}
		stage.OnAfterCodeRunwayProtectionAreaTypeDeleteCallback.OnAfterDelete(stage, coderunwayprotectionareatype, frontCodeRunwayProtectionAreaType)
	}
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRunwaySectionTypeCreateCallback != nil {
		stage.OnAfterCodeRunwaySectionTypeCreateCallback.OnAfterCreate(stage, coderunwaysectiontype)
	}
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwaySectionTypeUpdateCallback != nil {
		var frontCodeRunwaySectionType *CodeRunwaySectionType
		if front != nil {
			frontCodeRunwaySectionType, _ = front.(*CodeRunwaySectionType)
		}
		stage.OnAfterCodeRunwaySectionTypeUpdateCallback.OnAfterUpdate(stage, coderunwaysectiontype, frontCodeRunwaySectionType)
	}
}

func (coderunwaysectiontype *CodeRunwaySectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwaySectionTypeDeleteCallback != nil {
		var frontCodeRunwaySectionType *CodeRunwaySectionType
		if front != nil {
			frontCodeRunwaySectionType, _ = front.(*CodeRunwaySectionType)
		}
		stage.OnAfterCodeRunwaySectionTypeDeleteCallback.OnAfterDelete(stage, coderunwaysectiontype, frontCodeRunwaySectionType)
	}
}

func (coderunwaytype *CodeRunwayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRunwayTypeCreateCallback != nil {
		stage.OnAfterCodeRunwayTypeCreateCallback.OnAfterCreate(stage, coderunwaytype)
	}
}

func (coderunwaytype *CodeRunwayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayTypeUpdateCallback != nil {
		var frontCodeRunwayType *CodeRunwayType
		if front != nil {
			frontCodeRunwayType, _ = front.(*CodeRunwayType)
		}
		stage.OnAfterCodeRunwayTypeUpdateCallback.OnAfterUpdate(stage, coderunwaytype, frontCodeRunwayType)
	}
}

func (coderunwaytype *CodeRunwayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRunwayTypeDeleteCallback != nil {
		var frontCodeRunwayType *CodeRunwayType
		if front != nil {
			frontCodeRunwayType, _ = front.(*CodeRunwayType)
		}
		stage.OnAfterCodeRunwayTypeDeleteCallback.OnAfterDelete(stage, coderunwaytype, frontCodeRunwayType)
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

func (codestatusairporttype *CodeStatusAirportType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusAirportTypeCreateCallback != nil {
		stage.OnAfterCodeStatusAirportTypeCreateCallback.OnAfterCreate(stage, codestatusairporttype)
	}
}

func (codestatusairporttype *CodeStatusAirportType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusAirportTypeUpdateCallback != nil {
		var frontCodeStatusAirportType *CodeStatusAirportType
		if front != nil {
			frontCodeStatusAirportType, _ = front.(*CodeStatusAirportType)
		}
		stage.OnAfterCodeStatusAirportTypeUpdateCallback.OnAfterUpdate(stage, codestatusairporttype, frontCodeStatusAirportType)
	}
}

func (codestatusairporttype *CodeStatusAirportType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusAirportTypeDeleteCallback != nil {
		var frontCodeStatusAirportType *CodeStatusAirportType
		if front != nil {
			frontCodeStatusAirportType, _ = front.(*CodeStatusAirportType)
		}
		stage.OnAfterCodeStatusAirportTypeDeleteCallback.OnAfterDelete(stage, codestatusairporttype, frontCodeStatusAirportType)
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

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSurfaceCompositionTypeCreateCallback != nil {
		stage.OnAfterCodeSurfaceCompositionTypeCreateCallback.OnAfterCreate(stage, codesurfacecompositiontype)
	}
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSurfaceCompositionTypeUpdateCallback != nil {
		var frontCodeSurfaceCompositionType *CodeSurfaceCompositionType
		if front != nil {
			frontCodeSurfaceCompositionType, _ = front.(*CodeSurfaceCompositionType)
		}
		stage.OnAfterCodeSurfaceCompositionTypeUpdateCallback.OnAfterUpdate(stage, codesurfacecompositiontype, frontCodeSurfaceCompositionType)
	}
}

func (codesurfacecompositiontype *CodeSurfaceCompositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSurfaceCompositionTypeDeleteCallback != nil {
		var frontCodeSurfaceCompositionType *CodeSurfaceCompositionType
		if front != nil {
			frontCodeSurfaceCompositionType, _ = front.(*CodeSurfaceCompositionType)
		}
		stage.OnAfterCodeSurfaceCompositionTypeDeleteCallback.OnAfterDelete(stage, codesurfacecompositiontype, frontCodeSurfaceCompositionType)
	}
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSurfaceConditionTypeCreateCallback != nil {
		stage.OnAfterCodeSurfaceConditionTypeCreateCallback.OnAfterCreate(stage, codesurfaceconditiontype)
	}
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSurfaceConditionTypeUpdateCallback != nil {
		var frontCodeSurfaceConditionType *CodeSurfaceConditionType
		if front != nil {
			frontCodeSurfaceConditionType, _ = front.(*CodeSurfaceConditionType)
		}
		stage.OnAfterCodeSurfaceConditionTypeUpdateCallback.OnAfterUpdate(stage, codesurfaceconditiontype, frontCodeSurfaceConditionType)
	}
}

func (codesurfaceconditiontype *CodeSurfaceConditionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSurfaceConditionTypeDeleteCallback != nil {
		var frontCodeSurfaceConditionType *CodeSurfaceConditionType
		if front != nil {
			frontCodeSurfaceConditionType, _ = front.(*CodeSurfaceConditionType)
		}
		stage.OnAfterCodeSurfaceConditionTypeDeleteCallback.OnAfterDelete(stage, codesurfaceconditiontype, frontCodeSurfaceConditionType)
	}
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSurfacePreparationTypeCreateCallback != nil {
		stage.OnAfterCodeSurfacePreparationTypeCreateCallback.OnAfterCreate(stage, codesurfacepreparationtype)
	}
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSurfacePreparationTypeUpdateCallback != nil {
		var frontCodeSurfacePreparationType *CodeSurfacePreparationType
		if front != nil {
			frontCodeSurfacePreparationType, _ = front.(*CodeSurfacePreparationType)
		}
		stage.OnAfterCodeSurfacePreparationTypeUpdateCallback.OnAfterUpdate(stage, codesurfacepreparationtype, frontCodeSurfacePreparationType)
	}
}

func (codesurfacepreparationtype *CodeSurfacePreparationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSurfacePreparationTypeDeleteCallback != nil {
		var frontCodeSurfacePreparationType *CodeSurfacePreparationType
		if front != nil {
			frontCodeSurfacePreparationType, _ = front.(*CodeSurfacePreparationType)
		}
		stage.OnAfterCodeSurfacePreparationTypeDeleteCallback.OnAfterDelete(stage, codesurfacepreparationtype, frontCodeSurfacePreparationType)
	}
}

func (codetlofsectiontype *CodeTLOFSectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTLOFSectionTypeCreateCallback != nil {
		stage.OnAfterCodeTLOFSectionTypeCreateCallback.OnAfterCreate(stage, codetlofsectiontype)
	}
}

func (codetlofsectiontype *CodeTLOFSectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTLOFSectionTypeUpdateCallback != nil {
		var frontCodeTLOFSectionType *CodeTLOFSectionType
		if front != nil {
			frontCodeTLOFSectionType, _ = front.(*CodeTLOFSectionType)
		}
		stage.OnAfterCodeTLOFSectionTypeUpdateCallback.OnAfterUpdate(stage, codetlofsectiontype, frontCodeTLOFSectionType)
	}
}

func (codetlofsectiontype *CodeTLOFSectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTLOFSectionTypeDeleteCallback != nil {
		var frontCodeTLOFSectionType *CodeTLOFSectionType
		if front != nil {
			frontCodeTLOFSectionType, _ = front.(*CodeTLOFSectionType)
		}
		stage.OnAfterCodeTLOFSectionTypeDeleteCallback.OnAfterDelete(stage, codetlofsectiontype, frontCodeTLOFSectionType)
	}
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTaxiwayElementTypeCreateCallback != nil {
		stage.OnAfterCodeTaxiwayElementTypeCreateCallback.OnAfterCreate(stage, codetaxiwayelementtype)
	}
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTaxiwayElementTypeUpdateCallback != nil {
		var frontCodeTaxiwayElementType *CodeTaxiwayElementType
		if front != nil {
			frontCodeTaxiwayElementType, _ = front.(*CodeTaxiwayElementType)
		}
		stage.OnAfterCodeTaxiwayElementTypeUpdateCallback.OnAfterUpdate(stage, codetaxiwayelementtype, frontCodeTaxiwayElementType)
	}
}

func (codetaxiwayelementtype *CodeTaxiwayElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTaxiwayElementTypeDeleteCallback != nil {
		var frontCodeTaxiwayElementType *CodeTaxiwayElementType
		if front != nil {
			frontCodeTaxiwayElementType, _ = front.(*CodeTaxiwayElementType)
		}
		stage.OnAfterCodeTaxiwayElementTypeDeleteCallback.OnAfterDelete(stage, codetaxiwayelementtype, frontCodeTaxiwayElementType)
	}
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTaxiwaySectionTypeCreateCallback != nil {
		stage.OnAfterCodeTaxiwaySectionTypeCreateCallback.OnAfterCreate(stage, codetaxiwaysectiontype)
	}
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTaxiwaySectionTypeUpdateCallback != nil {
		var frontCodeTaxiwaySectionType *CodeTaxiwaySectionType
		if front != nil {
			frontCodeTaxiwaySectionType, _ = front.(*CodeTaxiwaySectionType)
		}
		stage.OnAfterCodeTaxiwaySectionTypeUpdateCallback.OnAfterUpdate(stage, codetaxiwaysectiontype, frontCodeTaxiwaySectionType)
	}
}

func (codetaxiwaysectiontype *CodeTaxiwaySectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTaxiwaySectionTypeDeleteCallback != nil {
		var frontCodeTaxiwaySectionType *CodeTaxiwaySectionType
		if front != nil {
			frontCodeTaxiwaySectionType, _ = front.(*CodeTaxiwaySectionType)
		}
		stage.OnAfterCodeTaxiwaySectionTypeDeleteCallback.OnAfterDelete(stage, codetaxiwaysectiontype, frontCodeTaxiwaySectionType)
	}
}

func (codetaxiwaytype *CodeTaxiwayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTaxiwayTypeCreateCallback != nil {
		stage.OnAfterCodeTaxiwayTypeCreateCallback.OnAfterCreate(stage, codetaxiwaytype)
	}
}

func (codetaxiwaytype *CodeTaxiwayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTaxiwayTypeUpdateCallback != nil {
		var frontCodeTaxiwayType *CodeTaxiwayType
		if front != nil {
			frontCodeTaxiwayType, _ = front.(*CodeTaxiwayType)
		}
		stage.OnAfterCodeTaxiwayTypeUpdateCallback.OnAfterUpdate(stage, codetaxiwaytype, frontCodeTaxiwayType)
	}
}

func (codetaxiwaytype *CodeTaxiwayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTaxiwayTypeDeleteCallback != nil {
		var frontCodeTaxiwayType *CodeTaxiwayType
		if front != nil {
			frontCodeTaxiwayType, _ = front.(*CodeTaxiwayType)
		}
		stage.OnAfterCodeTaxiwayTypeDeleteCallback.OnAfterDelete(stage, codetaxiwaytype, frontCodeTaxiwayType)
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

func (codeusagelimitationtype *CodeUsageLimitationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeUsageLimitationTypeCreateCallback != nil {
		stage.OnAfterCodeUsageLimitationTypeCreateCallback.OnAfterCreate(stage, codeusagelimitationtype)
	}
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUsageLimitationTypeUpdateCallback != nil {
		var frontCodeUsageLimitationType *CodeUsageLimitationType
		if front != nil {
			frontCodeUsageLimitationType, _ = front.(*CodeUsageLimitationType)
		}
		stage.OnAfterCodeUsageLimitationTypeUpdateCallback.OnAfterUpdate(stage, codeusagelimitationtype, frontCodeUsageLimitationType)
	}
}

func (codeusagelimitationtype *CodeUsageLimitationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUsageLimitationTypeDeleteCallback != nil {
		var frontCodeUsageLimitationType *CodeUsageLimitationType
		if front != nil {
			frontCodeUsageLimitationType, _ = front.(*CodeUsageLimitationType)
		}
		stage.OnAfterCodeUsageLimitationTypeDeleteCallback.OnAfterDelete(stage, codeusagelimitationtype, frontCodeUsageLimitationType)
	}
}

func (codevasistype *CodeVASISType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVASISTypeCreateCallback != nil {
		stage.OnAfterCodeVASISTypeCreateCallback.OnAfterCreate(stage, codevasistype)
	}
}

func (codevasistype *CodeVASISType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVASISTypeUpdateCallback != nil {
		var frontCodeVASISType *CodeVASISType
		if front != nil {
			frontCodeVASISType, _ = front.(*CodeVASISType)
		}
		stage.OnAfterCodeVASISTypeUpdateCallback.OnAfterUpdate(stage, codevasistype, frontCodeVASISType)
	}
}

func (codevasistype *CodeVASISType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVASISTypeDeleteCallback != nil {
		var frontCodeVASISType *CodeVASISType
		if front != nil {
			frontCodeVASISType, _ = front.(*CodeVASISType)
		}
		stage.OnAfterCodeVASISTypeDeleteCallback.OnAfterDelete(stage, codevasistype, frontCodeVASISType)
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

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVisualDockingGuidanceTypeCreateCallback != nil {
		stage.OnAfterCodeVisualDockingGuidanceTypeCreateCallback.OnAfterCreate(stage, codevisualdockingguidancetype)
	}
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVisualDockingGuidanceTypeUpdateCallback != nil {
		var frontCodeVisualDockingGuidanceType *CodeVisualDockingGuidanceType
		if front != nil {
			frontCodeVisualDockingGuidanceType, _ = front.(*CodeVisualDockingGuidanceType)
		}
		stage.OnAfterCodeVisualDockingGuidanceTypeUpdateCallback.OnAfterUpdate(stage, codevisualdockingguidancetype, frontCodeVisualDockingGuidanceType)
	}
}

func (codevisualdockingguidancetype *CodeVisualDockingGuidanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVisualDockingGuidanceTypeDeleteCallback != nil {
		var frontCodeVisualDockingGuidanceType *CodeVisualDockingGuidanceType
		if front != nil {
			frontCodeVisualDockingGuidanceType, _ = front.(*CodeVisualDockingGuidanceType)
		}
		stage.OnAfterCodeVisualDockingGuidanceTypeDeleteCallback.OnAfterDelete(stage, codevisualdockingguidancetype, frontCodeVisualDockingGuidanceType)
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

func (codeworkareatype *CodeWorkAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeWorkAreaTypeCreateCallback != nil {
		stage.OnAfterCodeWorkAreaTypeCreateCallback.OnAfterCreate(stage, codeworkareatype)
	}
}

func (codeworkareatype *CodeWorkAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeWorkAreaTypeUpdateCallback != nil {
		var frontCodeWorkAreaType *CodeWorkAreaType
		if front != nil {
			frontCodeWorkAreaType, _ = front.(*CodeWorkAreaType)
		}
		stage.OnAfterCodeWorkAreaTypeUpdateCallback.OnAfterUpdate(stage, codeworkareatype, frontCodeWorkAreaType)
	}
}

func (codeworkareatype *CodeWorkAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeWorkAreaTypeDeleteCallback != nil {
		var frontCodeWorkAreaType *CodeWorkAreaType
		if front != nil {
			frontCodeWorkAreaType, _ = front.(*CodeWorkAreaType)
		}
		stage.OnAfterCodeWorkAreaTypeDeleteCallback.OnAfterDelete(stage, codeworkareatype, frontCodeWorkAreaType)
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

func (datetimetype *DateTimeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDateTimeTypeCreateCallback != nil {
		stage.OnAfterDateTimeTypeCreateCallback.OnAfterCreate(stage, datetimetype)
	}
}

func (datetimetype *DateTimeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateTimeTypeUpdateCallback != nil {
		var frontDateTimeType *DateTimeType
		if front != nil {
			frontDateTimeType, _ = front.(*DateTimeType)
		}
		stage.OnAfterDateTimeTypeUpdateCallback.OnAfterUpdate(stage, datetimetype, frontDateTimeType)
	}
}

func (datetimetype *DateTimeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateTimeTypeDeleteCallback != nil {
		var frontDateTimeType *DateTimeType
		if front != nil {
			frontDateTimeType, _ = front.(*DateTimeType)
		}
		stage.OnAfterDateTimeTypeDeleteCallback.OnAfterDelete(stage, datetimetype, frontDateTimeType)
	}
}

func (datetype *DateType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDateTypeCreateCallback != nil {
		stage.OnAfterDateTypeCreateCallback.OnAfterCreate(stage, datetype)
	}
}

func (datetype *DateType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateTypeUpdateCallback != nil {
		var frontDateType *DateType
		if front != nil {
			frontDateType, _ = front.(*DateType)
		}
		stage.OnAfterDateTypeUpdateCallback.OnAfterUpdate(stage, datetype, frontDateType)
	}
}

func (datetype *DateType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateTypeDeleteCallback != nil {
		var frontDateType *DateType
		if front != nil {
			frontDateType, _ = front.(*DateType)
		}
		stage.OnAfterDateTypeDeleteCallback.OnAfterDelete(stage, datetype, frontDateType)
	}
}

func (dateyeartype *DateYearType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDateYearTypeCreateCallback != nil {
		stage.OnAfterDateYearTypeCreateCallback.OnAfterCreate(stage, dateyeartype)
	}
}

func (dateyeartype *DateYearType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateYearTypeUpdateCallback != nil {
		var frontDateYearType *DateYearType
		if front != nil {
			frontDateYearType, _ = front.(*DateYearType)
		}
		stage.OnAfterDateYearTypeUpdateCallback.OnAfterUpdate(stage, dateyeartype, frontDateYearType)
	}
}

func (dateyeartype *DateYearType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateYearTypeDeleteCallback != nil {
		var frontDateYearType *DateYearType
		if front != nil {
			frontDateYearType, _ = front.(*DateYearType)
		}
		stage.OnAfterDateYearTypeDeleteCallback.OnAfterDelete(stage, dateyeartype, frontDateYearType)
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

func (valdepthtype *ValDepthType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValDepthTypeCreateCallback != nil {
		stage.OnAfterValDepthTypeCreateCallback.OnAfterCreate(stage, valdepthtype)
	}
}

func (valdepthtype *ValDepthType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDepthTypeUpdateCallback != nil {
		var frontValDepthType *ValDepthType
		if front != nil {
			frontValDepthType, _ = front.(*ValDepthType)
		}
		stage.OnAfterValDepthTypeUpdateCallback.OnAfterUpdate(stage, valdepthtype, frontValDepthType)
	}
}

func (valdepthtype *ValDepthType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValDepthTypeDeleteCallback != nil {
		var frontValDepthType *ValDepthType
		if front != nil {
			frontValDepthType, _ = front.(*ValDepthType)
		}
		stage.OnAfterValDepthTypeDeleteCallback.OnAfterDelete(stage, valdepthtype, frontValDepthType)
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

func (valfltype *ValFLType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValFLTypeCreateCallback != nil {
		stage.OnAfterValFLTypeCreateCallback.OnAfterCreate(stage, valfltype)
	}
}

func (valfltype *ValFLType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValFLTypeUpdateCallback != nil {
		var frontValFLType *ValFLType
		if front != nil {
			frontValFLType, _ = front.(*ValFLType)
		}
		stage.OnAfterValFLTypeUpdateCallback.OnAfterUpdate(stage, valfltype, frontValFLType)
	}
}

func (valfltype *ValFLType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValFLTypeDeleteCallback != nil {
		var frontValFLType *ValFLType
		if front != nil {
			frontValFLType, _ = front.(*ValFLType)
		}
		stage.OnAfterValFLTypeDeleteCallback.OnAfterDelete(stage, valfltype, frontValFLType)
	}
}

func (valfrictiontype *ValFrictionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValFrictionTypeCreateCallback != nil {
		stage.OnAfterValFrictionTypeCreateCallback.OnAfterCreate(stage, valfrictiontype)
	}
}

func (valfrictiontype *ValFrictionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValFrictionTypeUpdateCallback != nil {
		var frontValFrictionType *ValFrictionType
		if front != nil {
			frontValFrictionType, _ = front.(*ValFrictionType)
		}
		stage.OnAfterValFrictionTypeUpdateCallback.OnAfterUpdate(stage, valfrictiontype, frontValFrictionType)
	}
}

func (valfrictiontype *ValFrictionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValFrictionTypeDeleteCallback != nil {
		var frontValFrictionType *ValFrictionType
		if front != nil {
			frontValFrictionType, _ = front.(*ValFrictionType)
		}
		stage.OnAfterValFrictionTypeDeleteCallback.OnAfterDelete(stage, valfrictiontype, frontValFrictionType)
	}
}

func (vallcntype *ValLCNType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValLCNTypeCreateCallback != nil {
		stage.OnAfterValLCNTypeCreateCallback.OnAfterCreate(stage, vallcntype)
	}
}

func (vallcntype *ValLCNType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValLCNTypeUpdateCallback != nil {
		var frontValLCNType *ValLCNType
		if front != nil {
			frontValLCNType, _ = front.(*ValLCNType)
		}
		stage.OnAfterValLCNTypeUpdateCallback.OnAfterUpdate(stage, vallcntype, frontValLCNType)
	}
}

func (vallcntype *ValLCNType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValLCNTypeDeleteCallback != nil {
		var frontValLCNType *ValLCNType
		if front != nil {
			frontValLCNType, _ = front.(*ValLCNType)
		}
		stage.OnAfterValLCNTypeDeleteCallback.OnAfterDelete(stage, vallcntype, frontValLCNType)
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

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValMagneticVariationChangeTypeCreateCallback != nil {
		stage.OnAfterValMagneticVariationChangeTypeCreateCallback.OnAfterCreate(stage, valmagneticvariationchangetype)
	}
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValMagneticVariationChangeTypeUpdateCallback != nil {
		var frontValMagneticVariationChangeType *ValMagneticVariationChangeType
		if front != nil {
			frontValMagneticVariationChangeType, _ = front.(*ValMagneticVariationChangeType)
		}
		stage.OnAfterValMagneticVariationChangeTypeUpdateCallback.OnAfterUpdate(stage, valmagneticvariationchangetype, frontValMagneticVariationChangeType)
	}
}

func (valmagneticvariationchangetype *ValMagneticVariationChangeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValMagneticVariationChangeTypeDeleteCallback != nil {
		var frontValMagneticVariationChangeType *ValMagneticVariationChangeType
		if front != nil {
			frontValMagneticVariationChangeType, _ = front.(*ValMagneticVariationChangeType)
		}
		stage.OnAfterValMagneticVariationChangeTypeDeleteCallback.OnAfterDelete(stage, valmagneticvariationchangetype, frontValMagneticVariationChangeType)
	}
}

func (valmagneticvariationtype *ValMagneticVariationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValMagneticVariationTypeCreateCallback != nil {
		stage.OnAfterValMagneticVariationTypeCreateCallback.OnAfterCreate(stage, valmagneticvariationtype)
	}
}

func (valmagneticvariationtype *ValMagneticVariationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValMagneticVariationTypeUpdateCallback != nil {
		var frontValMagneticVariationType *ValMagneticVariationType
		if front != nil {
			frontValMagneticVariationType, _ = front.(*ValMagneticVariationType)
		}
		stage.OnAfterValMagneticVariationTypeUpdateCallback.OnAfterUpdate(stage, valmagneticvariationtype, frontValMagneticVariationType)
	}
}

func (valmagneticvariationtype *ValMagneticVariationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValMagneticVariationTypeDeleteCallback != nil {
		var frontValMagneticVariationType *ValMagneticVariationType
		if front != nil {
			frontValMagneticVariationType, _ = front.(*ValMagneticVariationType)
		}
		stage.OnAfterValMagneticVariationTypeDeleteCallback.OnAfterDelete(stage, valmagneticvariationtype, frontValMagneticVariationType)
	}
}

func (valpcntype *ValPCNType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValPCNTypeCreateCallback != nil {
		stage.OnAfterValPCNTypeCreateCallback.OnAfterCreate(stage, valpcntype)
	}
}

func (valpcntype *ValPCNType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValPCNTypeUpdateCallback != nil {
		var frontValPCNType *ValPCNType
		if front != nil {
			frontValPCNType, _ = front.(*ValPCNType)
		}
		stage.OnAfterValPCNTypeUpdateCallback.OnAfterUpdate(stage, valpcntype, frontValPCNType)
	}
}

func (valpcntype *ValPCNType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValPCNTypeDeleteCallback != nil {
		var frontValPCNType *ValPCNType
		if front != nil {
			frontValPCNType, _ = front.(*ValPCNType)
		}
		stage.OnAfterValPCNTypeDeleteCallback.OnAfterDelete(stage, valpcntype, frontValPCNType)
	}
}

func (valpercenttype *ValPercentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValPercentTypeCreateCallback != nil {
		stage.OnAfterValPercentTypeCreateCallback.OnAfterCreate(stage, valpercenttype)
	}
}

func (valpercenttype *ValPercentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValPercentTypeUpdateCallback != nil {
		var frontValPercentType *ValPercentType
		if front != nil {
			frontValPercentType, _ = front.(*ValPercentType)
		}
		stage.OnAfterValPercentTypeUpdateCallback.OnAfterUpdate(stage, valpercenttype, frontValPercentType)
	}
}

func (valpercenttype *ValPercentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValPercentTypeDeleteCallback != nil {
		var frontValPercentType *ValPercentType
		if front != nil {
			frontValPercentType, _ = front.(*ValPercentType)
		}
		stage.OnAfterValPercentTypeDeleteCallback.OnAfterDelete(stage, valpercenttype, frontValPercentType)
	}
}

func (valpressuretype *ValPressureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValPressureTypeCreateCallback != nil {
		stage.OnAfterValPressureTypeCreateCallback.OnAfterCreate(stage, valpressuretype)
	}
}

func (valpressuretype *ValPressureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValPressureTypeUpdateCallback != nil {
		var frontValPressureType *ValPressureType
		if front != nil {
			frontValPressureType, _ = front.(*ValPressureType)
		}
		stage.OnAfterValPressureTypeUpdateCallback.OnAfterUpdate(stage, valpressuretype, frontValPressureType)
	}
}

func (valpressuretype *ValPressureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValPressureTypeDeleteCallback != nil {
		var frontValPressureType *ValPressureType
		if front != nil {
			frontValPressureType, _ = front.(*ValPressureType)
		}
		stage.OnAfterValPressureTypeDeleteCallback.OnAfterDelete(stage, valpressuretype, frontValPressureType)
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

