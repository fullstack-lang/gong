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

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAircraftGroundServiceTypeCreateCallback != nil {
		stage.OnAfterCodeAircraftGroundServiceTypeCreateCallback.OnAfterCreate(stage, codeaircraftgroundservicetype)
	}
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftGroundServiceTypeUpdateCallback != nil {
		var frontCodeAircraftGroundServiceType *CodeAircraftGroundServiceType
		if front != nil {
			frontCodeAircraftGroundServiceType, _ = front.(*CodeAircraftGroundServiceType)
		}
		stage.OnAfterCodeAircraftGroundServiceTypeUpdateCallback.OnAfterUpdate(stage, codeaircraftgroundservicetype, frontCodeAircraftGroundServiceType)
	}
}

func (codeaircraftgroundservicetype *CodeAircraftGroundServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAircraftGroundServiceTypeDeleteCallback != nil {
		var frontCodeAircraftGroundServiceType *CodeAircraftGroundServiceType
		if front != nil {
			frontCodeAircraftGroundServiceType, _ = front.(*CodeAircraftGroundServiceType)
		}
		stage.OnAfterCodeAircraftGroundServiceTypeDeleteCallback.OnAfterDelete(stage, codeaircraftgroundservicetype, frontCodeAircraftGroundServiceType)
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

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspaceActivityTypeCreateCallback != nil {
		stage.OnAfterCodeAirspaceActivityTypeCreateCallback.OnAfterCreate(stage, codeairspaceactivitytype)
	}
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceActivityTypeUpdateCallback != nil {
		var frontCodeAirspaceActivityType *CodeAirspaceActivityType
		if front != nil {
			frontCodeAirspaceActivityType, _ = front.(*CodeAirspaceActivityType)
		}
		stage.OnAfterCodeAirspaceActivityTypeUpdateCallback.OnAfterUpdate(stage, codeairspaceactivitytype, frontCodeAirspaceActivityType)
	}
}

func (codeairspaceactivitytype *CodeAirspaceActivityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceActivityTypeDeleteCallback != nil {
		var frontCodeAirspaceActivityType *CodeAirspaceActivityType
		if front != nil {
			frontCodeAirspaceActivityType, _ = front.(*CodeAirspaceActivityType)
		}
		stage.OnAfterCodeAirspaceActivityTypeDeleteCallback.OnAfterDelete(stage, codeairspaceactivitytype, frontCodeAirspaceActivityType)
	}
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspaceAggregationTypeCreateCallback != nil {
		stage.OnAfterCodeAirspaceAggregationTypeCreateCallback.OnAfterCreate(stage, codeairspaceaggregationtype)
	}
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceAggregationTypeUpdateCallback != nil {
		var frontCodeAirspaceAggregationType *CodeAirspaceAggregationType
		if front != nil {
			frontCodeAirspaceAggregationType, _ = front.(*CodeAirspaceAggregationType)
		}
		stage.OnAfterCodeAirspaceAggregationTypeUpdateCallback.OnAfterUpdate(stage, codeairspaceaggregationtype, frontCodeAirspaceAggregationType)
	}
}

func (codeairspaceaggregationtype *CodeAirspaceAggregationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceAggregationTypeDeleteCallback != nil {
		var frontCodeAirspaceAggregationType *CodeAirspaceAggregationType
		if front != nil {
			frontCodeAirspaceAggregationType, _ = front.(*CodeAirspaceAggregationType)
		}
		stage.OnAfterCodeAirspaceAggregationTypeDeleteCallback.OnAfterDelete(stage, codeairspaceaggregationtype, frontCodeAirspaceAggregationType)
	}
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspaceClassificationTypeCreateCallback != nil {
		stage.OnAfterCodeAirspaceClassificationTypeCreateCallback.OnAfterCreate(stage, codeairspaceclassificationtype)
	}
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceClassificationTypeUpdateCallback != nil {
		var frontCodeAirspaceClassificationType *CodeAirspaceClassificationType
		if front != nil {
			frontCodeAirspaceClassificationType, _ = front.(*CodeAirspaceClassificationType)
		}
		stage.OnAfterCodeAirspaceClassificationTypeUpdateCallback.OnAfterUpdate(stage, codeairspaceclassificationtype, frontCodeAirspaceClassificationType)
	}
}

func (codeairspaceclassificationtype *CodeAirspaceClassificationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceClassificationTypeDeleteCallback != nil {
		var frontCodeAirspaceClassificationType *CodeAirspaceClassificationType
		if front != nil {
			frontCodeAirspaceClassificationType, _ = front.(*CodeAirspaceClassificationType)
		}
		stage.OnAfterCodeAirspaceClassificationTypeDeleteCallback.OnAfterDelete(stage, codeairspaceclassificationtype, frontCodeAirspaceClassificationType)
	}
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspaceDependencyTypeCreateCallback != nil {
		stage.OnAfterCodeAirspaceDependencyTypeCreateCallback.OnAfterCreate(stage, codeairspacedependencytype)
	}
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceDependencyTypeUpdateCallback != nil {
		var frontCodeAirspaceDependencyType *CodeAirspaceDependencyType
		if front != nil {
			frontCodeAirspaceDependencyType, _ = front.(*CodeAirspaceDependencyType)
		}
		stage.OnAfterCodeAirspaceDependencyTypeUpdateCallback.OnAfterUpdate(stage, codeairspacedependencytype, frontCodeAirspaceDependencyType)
	}
}

func (codeairspacedependencytype *CodeAirspaceDependencyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceDependencyTypeDeleteCallback != nil {
		var frontCodeAirspaceDependencyType *CodeAirspaceDependencyType
		if front != nil {
			frontCodeAirspaceDependencyType, _ = front.(*CodeAirspaceDependencyType)
		}
		stage.OnAfterCodeAirspaceDependencyTypeDeleteCallback.OnAfterDelete(stage, codeairspacedependencytype, frontCodeAirspaceDependencyType)
	}
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspaceDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeAirspaceDesignatorTypeCreateCallback.OnAfterCreate(stage, codeairspacedesignatortype)
	}
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceDesignatorTypeUpdateCallback != nil {
		var frontCodeAirspaceDesignatorType *CodeAirspaceDesignatorType
		if front != nil {
			frontCodeAirspaceDesignatorType, _ = front.(*CodeAirspaceDesignatorType)
		}
		stage.OnAfterCodeAirspaceDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codeairspacedesignatortype, frontCodeAirspaceDesignatorType)
	}
}

func (codeairspacedesignatortype *CodeAirspaceDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceDesignatorTypeDeleteCallback != nil {
		var frontCodeAirspaceDesignatorType *CodeAirspaceDesignatorType
		if front != nil {
			frontCodeAirspaceDesignatorType, _ = front.(*CodeAirspaceDesignatorType)
		}
		stage.OnAfterCodeAirspaceDesignatorTypeDeleteCallback.OnAfterDelete(stage, codeairspacedesignatortype, frontCodeAirspaceDesignatorType)
	}
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspacePointPositionTypeCreateCallback != nil {
		stage.OnAfterCodeAirspacePointPositionTypeCreateCallback.OnAfterCreate(stage, codeairspacepointpositiontype)
	}
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspacePointPositionTypeUpdateCallback != nil {
		var frontCodeAirspacePointPositionType *CodeAirspacePointPositionType
		if front != nil {
			frontCodeAirspacePointPositionType, _ = front.(*CodeAirspacePointPositionType)
		}
		stage.OnAfterCodeAirspacePointPositionTypeUpdateCallback.OnAfterUpdate(stage, codeairspacepointpositiontype, frontCodeAirspacePointPositionType)
	}
}

func (codeairspacepointpositiontype *CodeAirspacePointPositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspacePointPositionTypeDeleteCallback != nil {
		var frontCodeAirspacePointPositionType *CodeAirspacePointPositionType
		if front != nil {
			frontCodeAirspacePointPositionType, _ = front.(*CodeAirspacePointPositionType)
		}
		stage.OnAfterCodeAirspacePointPositionTypeDeleteCallback.OnAfterDelete(stage, codeairspacepointpositiontype, frontCodeAirspacePointPositionType)
	}
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspacePointRoleTypeCreateCallback != nil {
		stage.OnAfterCodeAirspacePointRoleTypeCreateCallback.OnAfterCreate(stage, codeairspacepointroletype)
	}
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspacePointRoleTypeUpdateCallback != nil {
		var frontCodeAirspacePointRoleType *CodeAirspacePointRoleType
		if front != nil {
			frontCodeAirspacePointRoleType, _ = front.(*CodeAirspacePointRoleType)
		}
		stage.OnAfterCodeAirspacePointRoleTypeUpdateCallback.OnAfterUpdate(stage, codeairspacepointroletype, frontCodeAirspacePointRoleType)
	}
}

func (codeairspacepointroletype *CodeAirspacePointRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspacePointRoleTypeDeleteCallback != nil {
		var frontCodeAirspacePointRoleType *CodeAirspacePointRoleType
		if front != nil {
			frontCodeAirspacePointRoleType, _ = front.(*CodeAirspacePointRoleType)
		}
		stage.OnAfterCodeAirspacePointRoleTypeDeleteCallback.OnAfterDelete(stage, codeairspacepointroletype, frontCodeAirspacePointRoleType)
	}
}

func (codeairspacetype *CodeAirspaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAirspaceTypeCreateCallback != nil {
		stage.OnAfterCodeAirspaceTypeCreateCallback.OnAfterCreate(stage, codeairspacetype)
	}
}

func (codeairspacetype *CodeAirspaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceTypeUpdateCallback != nil {
		var frontCodeAirspaceType *CodeAirspaceType
		if front != nil {
			frontCodeAirspaceType, _ = front.(*CodeAirspaceType)
		}
		stage.OnAfterCodeAirspaceTypeUpdateCallback.OnAfterUpdate(stage, codeairspacetype, frontCodeAirspaceType)
	}
}

func (codeairspacetype *CodeAirspaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAirspaceTypeDeleteCallback != nil {
		var frontCodeAirspaceType *CodeAirspaceType
		if front != nil {
			frontCodeAirspaceType, _ = front.(*CodeAirspaceType)
		}
		stage.OnAfterCodeAirspaceTypeDeleteCallback.OnAfterDelete(stage, codeairspacetype, frontCodeAirspaceType)
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

func (codeauralmorsetype *CodeAuralMorseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAuralMorseTypeCreateCallback != nil {
		stage.OnAfterCodeAuralMorseTypeCreateCallback.OnAfterCreate(stage, codeauralmorsetype)
	}
}

func (codeauralmorsetype *CodeAuralMorseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuralMorseTypeUpdateCallback != nil {
		var frontCodeAuralMorseType *CodeAuralMorseType
		if front != nil {
			frontCodeAuralMorseType, _ = front.(*CodeAuralMorseType)
		}
		stage.OnAfterCodeAuralMorseTypeUpdateCallback.OnAfterUpdate(stage, codeauralmorsetype, frontCodeAuralMorseType)
	}
}

func (codeauralmorsetype *CodeAuralMorseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAuralMorseTypeDeleteCallback != nil {
		var frontCodeAuralMorseType *CodeAuralMorseType
		if front != nil {
			frontCodeAuralMorseType, _ = front.(*CodeAuralMorseType)
		}
		stage.OnAfterCodeAuralMorseTypeDeleteCallback.OnAfterDelete(stage, codeauralmorsetype, frontCodeAuralMorseType)
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

func (codeaviationstandardstype *CodeAviationStandardsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeAviationStandardsTypeCreateCallback != nil {
		stage.OnAfterCodeAviationStandardsTypeCreateCallback.OnAfterCreate(stage, codeaviationstandardstype)
	}
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAviationStandardsTypeUpdateCallback != nil {
		var frontCodeAviationStandardsType *CodeAviationStandardsType
		if front != nil {
			frontCodeAviationStandardsType, _ = front.(*CodeAviationStandardsType)
		}
		stage.OnAfterCodeAviationStandardsTypeUpdateCallback.OnAfterUpdate(stage, codeaviationstandardstype, frontCodeAviationStandardsType)
	}
}

func (codeaviationstandardstype *CodeAviationStandardsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeAviationStandardsTypeDeleteCallback != nil {
		var frontCodeAviationStandardsType *CodeAviationStandardsType
		if front != nil {
			frontCodeAviationStandardsType, _ = front.(*CodeAviationStandardsType)
		}
		stage.OnAfterCodeAviationStandardsTypeDeleteCallback.OnAfterDelete(stage, codeaviationstandardstype, frontCodeAviationStandardsType)
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

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCheckpointCategoryTypeCreateCallback != nil {
		stage.OnAfterCodeCheckpointCategoryTypeCreateCallback.OnAfterCreate(stage, codecheckpointcategorytype)
	}
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCheckpointCategoryTypeUpdateCallback != nil {
		var frontCodeCheckpointCategoryType *CodeCheckpointCategoryType
		if front != nil {
			frontCodeCheckpointCategoryType, _ = front.(*CodeCheckpointCategoryType)
		}
		stage.OnAfterCodeCheckpointCategoryTypeUpdateCallback.OnAfterUpdate(stage, codecheckpointcategorytype, frontCodeCheckpointCategoryType)
	}
}

func (codecheckpointcategorytype *CodeCheckpointCategoryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCheckpointCategoryTypeDeleteCallback != nil {
		var frontCodeCheckpointCategoryType *CodeCheckpointCategoryType
		if front != nil {
			frontCodeCheckpointCategoryType, _ = front.(*CodeCheckpointCategoryType)
		}
		stage.OnAfterCodeCheckpointCategoryTypeDeleteCallback.OnAfterDelete(stage, codecheckpointcategorytype, frontCodeCheckpointCategoryType)
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

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCommunicationChannelTypeCreateCallback != nil {
		stage.OnAfterCodeCommunicationChannelTypeCreateCallback.OnAfterCreate(stage, codecommunicationchanneltype)
	}
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationChannelTypeUpdateCallback != nil {
		var frontCodeCommunicationChannelType *CodeCommunicationChannelType
		if front != nil {
			frontCodeCommunicationChannelType, _ = front.(*CodeCommunicationChannelType)
		}
		stage.OnAfterCodeCommunicationChannelTypeUpdateCallback.OnAfterUpdate(stage, codecommunicationchanneltype, frontCodeCommunicationChannelType)
	}
}

func (codecommunicationchanneltype *CodeCommunicationChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationChannelTypeDeleteCallback != nil {
		var frontCodeCommunicationChannelType *CodeCommunicationChannelType
		if front != nil {
			frontCodeCommunicationChannelType, _ = front.(*CodeCommunicationChannelType)
		}
		stage.OnAfterCodeCommunicationChannelTypeDeleteCallback.OnAfterDelete(stage, codecommunicationchanneltype, frontCodeCommunicationChannelType)
	}
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCommunicationDirectionTypeCreateCallback != nil {
		stage.OnAfterCodeCommunicationDirectionTypeCreateCallback.OnAfterCreate(stage, codecommunicationdirectiontype)
	}
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationDirectionTypeUpdateCallback != nil {
		var frontCodeCommunicationDirectionType *CodeCommunicationDirectionType
		if front != nil {
			frontCodeCommunicationDirectionType, _ = front.(*CodeCommunicationDirectionType)
		}
		stage.OnAfterCodeCommunicationDirectionTypeUpdateCallback.OnAfterUpdate(stage, codecommunicationdirectiontype, frontCodeCommunicationDirectionType)
	}
}

func (codecommunicationdirectiontype *CodeCommunicationDirectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCommunicationDirectionTypeDeleteCallback != nil {
		var frontCodeCommunicationDirectionType *CodeCommunicationDirectionType
		if front != nil {
			frontCodeCommunicationDirectionType, _ = front.(*CodeCommunicationDirectionType)
		}
		stage.OnAfterCodeCommunicationDirectionTypeDeleteCallback.OnAfterDelete(stage, codecommunicationdirectiontype, frontCodeCommunicationDirectionType)
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

func (codecoursequalityilstype *CodeCourseQualityILSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeCourseQualityILSTypeCreateCallback != nil {
		stage.OnAfterCodeCourseQualityILSTypeCreateCallback.OnAfterCreate(stage, codecoursequalityilstype)
	}
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCourseQualityILSTypeUpdateCallback != nil {
		var frontCodeCourseQualityILSType *CodeCourseQualityILSType
		if front != nil {
			frontCodeCourseQualityILSType, _ = front.(*CodeCourseQualityILSType)
		}
		stage.OnAfterCodeCourseQualityILSTypeUpdateCallback.OnAfterUpdate(stage, codecoursequalityilstype, frontCodeCourseQualityILSType)
	}
}

func (codecoursequalityilstype *CodeCourseQualityILSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeCourseQualityILSTypeDeleteCallback != nil {
		var frontCodeCourseQualityILSType *CodeCourseQualityILSType
		if front != nil {
			frontCodeCourseQualityILSType, _ = front.(*CodeCourseQualityILSType)
		}
		stage.OnAfterCodeCourseQualityILSTypeDeleteCallback.OnAfterDelete(stage, codecoursequalityilstype, frontCodeCourseQualityILSType)
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

func (codedmechanneltype *CodeDMEChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDMEChannelTypeCreateCallback != nil {
		stage.OnAfterCodeDMEChannelTypeCreateCallback.OnAfterCreate(stage, codedmechanneltype)
	}
}

func (codedmechanneltype *CodeDMEChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDMEChannelTypeUpdateCallback != nil {
		var frontCodeDMEChannelType *CodeDMEChannelType
		if front != nil {
			frontCodeDMEChannelType, _ = front.(*CodeDMEChannelType)
		}
		stage.OnAfterCodeDMEChannelTypeUpdateCallback.OnAfterUpdate(stage, codedmechanneltype, frontCodeDMEChannelType)
	}
}

func (codedmechanneltype *CodeDMEChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDMEChannelTypeDeleteCallback != nil {
		var frontCodeDMEChannelType *CodeDMEChannelType
		if front != nil {
			frontCodeDMEChannelType, _ = front.(*CodeDMEChannelType)
		}
		stage.OnAfterCodeDMEChannelTypeDeleteCallback.OnAfterDelete(stage, codedmechanneltype, frontCodeDMEChannelType)
	}
}

func (codedmetype *CodeDMEType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDMETypeCreateCallback != nil {
		stage.OnAfterCodeDMETypeCreateCallback.OnAfterCreate(stage, codedmetype)
	}
}

func (codedmetype *CodeDMEType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDMETypeUpdateCallback != nil {
		var frontCodeDMEType *CodeDMEType
		if front != nil {
			frontCodeDMEType, _ = front.(*CodeDMEType)
		}
		stage.OnAfterCodeDMETypeUpdateCallback.OnAfterUpdate(stage, codedmetype, frontCodeDMEType)
	}
}

func (codedmetype *CodeDMEType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDMETypeDeleteCallback != nil {
		var frontCodeDMEType *CodeDMEType
		if front != nil {
			frontCodeDMEType, _ = front.(*CodeDMEType)
		}
		stage.OnAfterCodeDMETypeDeleteCallback.OnAfterDelete(stage, codedmetype, frontCodeDMEType)
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

func (codedesignatedpointtype *CodeDesignatedPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDesignatedPointTypeCreateCallback != nil {
		stage.OnAfterCodeDesignatedPointTypeCreateCallback.OnAfterCreate(stage, codedesignatedpointtype)
	}
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignatedPointTypeUpdateCallback != nil {
		var frontCodeDesignatedPointType *CodeDesignatedPointType
		if front != nil {
			frontCodeDesignatedPointType, _ = front.(*CodeDesignatedPointType)
		}
		stage.OnAfterCodeDesignatedPointTypeUpdateCallback.OnAfterUpdate(stage, codedesignatedpointtype, frontCodeDesignatedPointType)
	}
}

func (codedesignatedpointtype *CodeDesignatedPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDesignatedPointTypeDeleteCallback != nil {
		var frontCodeDesignatedPointType *CodeDesignatedPointType
		if front != nil {
			frontCodeDesignatedPointType, _ = front.(*CodeDesignatedPointType)
		}
		stage.OnAfterCodeDesignatedPointTypeDeleteCallback.OnAfterDelete(stage, codedesignatedpointtype, frontCodeDesignatedPointType)
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

func (codedistanceindicationtype *CodeDistanceIndicationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeDistanceIndicationTypeCreateCallback != nil {
		stage.OnAfterCodeDistanceIndicationTypeCreateCallback.OnAfterCreate(stage, codedistanceindicationtype)
	}
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDistanceIndicationTypeUpdateCallback != nil {
		var frontCodeDistanceIndicationType *CodeDistanceIndicationType
		if front != nil {
			frontCodeDistanceIndicationType, _ = front.(*CodeDistanceIndicationType)
		}
		stage.OnAfterCodeDistanceIndicationTypeUpdateCallback.OnAfterUpdate(stage, codedistanceindicationtype, frontCodeDistanceIndicationType)
	}
}

func (codedistanceindicationtype *CodeDistanceIndicationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeDistanceIndicationTypeDeleteCallback != nil {
		var frontCodeDistanceIndicationType *CodeDistanceIndicationType
		if front != nil {
			frontCodeDistanceIndicationType, _ = front.(*CodeDistanceIndicationType)
		}
		stage.OnAfterCodeDistanceIndicationTypeDeleteCallback.OnAfterDelete(stage, codedistanceindicationtype, frontCodeDistanceIndicationType)
	}
}

func (codeemissionbandtype *CodeEmissionBandType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeEmissionBandTypeCreateCallback != nil {
		stage.OnAfterCodeEmissionBandTypeCreateCallback.OnAfterCreate(stage, codeemissionbandtype)
	}
}

func (codeemissionbandtype *CodeEmissionBandType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeEmissionBandTypeUpdateCallback != nil {
		var frontCodeEmissionBandType *CodeEmissionBandType
		if front != nil {
			frontCodeEmissionBandType, _ = front.(*CodeEmissionBandType)
		}
		stage.OnAfterCodeEmissionBandTypeUpdateCallback.OnAfterUpdate(stage, codeemissionbandtype, frontCodeEmissionBandType)
	}
}

func (codeemissionbandtype *CodeEmissionBandType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeEmissionBandTypeDeleteCallback != nil {
		var frontCodeEmissionBandType *CodeEmissionBandType
		if front != nil {
			frontCodeEmissionBandType, _ = front.(*CodeEmissionBandType)
		}
		stage.OnAfterCodeEmissionBandTypeDeleteCallback.OnAfterDelete(stage, codeemissionbandtype, frontCodeEmissionBandType)
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

func (codefacilityrankingtype *CodeFacilityRankingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFacilityRankingTypeCreateCallback != nil {
		stage.OnAfterCodeFacilityRankingTypeCreateCallback.OnAfterCreate(stage, codefacilityrankingtype)
	}
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFacilityRankingTypeUpdateCallback != nil {
		var frontCodeFacilityRankingType *CodeFacilityRankingType
		if front != nil {
			frontCodeFacilityRankingType, _ = front.(*CodeFacilityRankingType)
		}
		stage.OnAfterCodeFacilityRankingTypeUpdateCallback.OnAfterUpdate(stage, codefacilityrankingtype, frontCodeFacilityRankingType)
	}
}

func (codefacilityrankingtype *CodeFacilityRankingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFacilityRankingTypeDeleteCallback != nil {
		var frontCodeFacilityRankingType *CodeFacilityRankingType
		if front != nil {
			frontCodeFacilityRankingType, _ = front.(*CodeFacilityRankingType)
		}
		stage.OnAfterCodeFacilityRankingTypeDeleteCallback.OnAfterDelete(stage, codefacilityrankingtype, frontCodeFacilityRankingType)
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

func (codefirefightingtype *CodeFireFightingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFireFightingTypeCreateCallback != nil {
		stage.OnAfterCodeFireFightingTypeCreateCallback.OnAfterCreate(stage, codefirefightingtype)
	}
}

func (codefirefightingtype *CodeFireFightingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFireFightingTypeUpdateCallback != nil {
		var frontCodeFireFightingType *CodeFireFightingType
		if front != nil {
			frontCodeFireFightingType, _ = front.(*CodeFireFightingType)
		}
		stage.OnAfterCodeFireFightingTypeUpdateCallback.OnAfterUpdate(stage, codefirefightingtype, frontCodeFireFightingType)
	}
}

func (codefirefightingtype *CodeFireFightingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFireFightingTypeDeleteCallback != nil {
		var frontCodeFireFightingType *CodeFireFightingType
		if front != nil {
			frontCodeFireFightingType, _ = front.(*CodeFireFightingType)
		}
		stage.OnAfterCodeFireFightingTypeDeleteCallback.OnAfterDelete(stage, codefirefightingtype, frontCodeFireFightingType)
	}
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFlightDestinationTypeCreateCallback != nil {
		stage.OnAfterCodeFlightDestinationTypeCreateCallback.OnAfterCreate(stage, codeflightdestinationtype)
	}
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightDestinationTypeUpdateCallback != nil {
		var frontCodeFlightDestinationType *CodeFlightDestinationType
		if front != nil {
			frontCodeFlightDestinationType, _ = front.(*CodeFlightDestinationType)
		}
		stage.OnAfterCodeFlightDestinationTypeUpdateCallback.OnAfterUpdate(stage, codeflightdestinationtype, frontCodeFlightDestinationType)
	}
}

func (codeflightdestinationtype *CodeFlightDestinationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFlightDestinationTypeDeleteCallback != nil {
		var frontCodeFlightDestinationType *CodeFlightDestinationType
		if front != nil {
			frontCodeFlightDestinationType, _ = front.(*CodeFlightDestinationType)
		}
		stage.OnAfterCodeFlightDestinationTypeDeleteCallback.OnAfterDelete(stage, codeflightdestinationtype, frontCodeFlightDestinationType)
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

func (codefueltype *CodeFuelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeFuelTypeCreateCallback != nil {
		stage.OnAfterCodeFuelTypeCreateCallback.OnAfterCreate(stage, codefueltype)
	}
}

func (codefueltype *CodeFuelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFuelTypeUpdateCallback != nil {
		var frontCodeFuelType *CodeFuelType
		if front != nil {
			frontCodeFuelType, _ = front.(*CodeFuelType)
		}
		stage.OnAfterCodeFuelTypeUpdateCallback.OnAfterUpdate(stage, codefueltype, frontCodeFuelType)
	}
}

func (codefueltype *CodeFuelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeFuelTypeDeleteCallback != nil {
		var frontCodeFuelType *CodeFuelType
		if front != nil {
			frontCodeFuelType, _ = front.(*CodeFuelType)
		}
		stage.OnAfterCodeFuelTypeDeleteCallback.OnAfterDelete(stage, codefueltype, frontCodeFuelType)
	}
}

func (codegeobordertype *CodeGeoBorderType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeGeoBorderTypeCreateCallback != nil {
		stage.OnAfterCodeGeoBorderTypeCreateCallback.OnAfterCreate(stage, codegeobordertype)
	}
}

func (codegeobordertype *CodeGeoBorderType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeGeoBorderTypeUpdateCallback != nil {
		var frontCodeGeoBorderType *CodeGeoBorderType
		if front != nil {
			frontCodeGeoBorderType, _ = front.(*CodeGeoBorderType)
		}
		stage.OnAfterCodeGeoBorderTypeUpdateCallback.OnAfterUpdate(stage, codegeobordertype, frontCodeGeoBorderType)
	}
}

func (codegeobordertype *CodeGeoBorderType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeGeoBorderTypeDeleteCallback != nil {
		var frontCodeGeoBorderType *CodeGeoBorderType
		if front != nil {
			frontCodeGeoBorderType, _ = front.(*CodeGeoBorderType)
		}
		stage.OnAfterCodeGeoBorderTypeDeleteCallback.OnAfterDelete(stage, codegeobordertype, frontCodeGeoBorderType)
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

func (codeilsbackcoursetype *CodeILSBackCourseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeILSBackCourseTypeCreateCallback != nil {
		stage.OnAfterCodeILSBackCourseTypeCreateCallback.OnAfterCreate(stage, codeilsbackcoursetype)
	}
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeILSBackCourseTypeUpdateCallback != nil {
		var frontCodeILSBackCourseType *CodeILSBackCourseType
		if front != nil {
			frontCodeILSBackCourseType, _ = front.(*CodeILSBackCourseType)
		}
		stage.OnAfterCodeILSBackCourseTypeUpdateCallback.OnAfterUpdate(stage, codeilsbackcoursetype, frontCodeILSBackCourseType)
	}
}

func (codeilsbackcoursetype *CodeILSBackCourseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeILSBackCourseTypeDeleteCallback != nil {
		var frontCodeILSBackCourseType *CodeILSBackCourseType
		if front != nil {
			frontCodeILSBackCourseType, _ = front.(*CodeILSBackCourseType)
		}
		stage.OnAfterCodeILSBackCourseTypeDeleteCallback.OnAfterDelete(stage, codeilsbackcoursetype, frontCodeILSBackCourseType)
	}
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeIntegrityLevelILSTypeCreateCallback != nil {
		stage.OnAfterCodeIntegrityLevelILSTypeCreateCallback.OnAfterCreate(stage, codeintegritylevelilstype)
	}
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIntegrityLevelILSTypeUpdateCallback != nil {
		var frontCodeIntegrityLevelILSType *CodeIntegrityLevelILSType
		if front != nil {
			frontCodeIntegrityLevelILSType, _ = front.(*CodeIntegrityLevelILSType)
		}
		stage.OnAfterCodeIntegrityLevelILSTypeUpdateCallback.OnAfterUpdate(stage, codeintegritylevelilstype, frontCodeIntegrityLevelILSType)
	}
}

func (codeintegritylevelilstype *CodeIntegrityLevelILSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIntegrityLevelILSTypeDeleteCallback != nil {
		var frontCodeIntegrityLevelILSType *CodeIntegrityLevelILSType
		if front != nil {
			frontCodeIntegrityLevelILSType, _ = front.(*CodeIntegrityLevelILSType)
		}
		stage.OnAfterCodeIntegrityLevelILSTypeDeleteCallback.OnAfterDelete(stage, codeintegritylevelilstype, frontCodeIntegrityLevelILSType)
	}
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeIntensityStandByTypeCreateCallback != nil {
		stage.OnAfterCodeIntensityStandByTypeCreateCallback.OnAfterCreate(stage, codeintensitystandbytype)
	}
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIntensityStandByTypeUpdateCallback != nil {
		var frontCodeIntensityStandByType *CodeIntensityStandByType
		if front != nil {
			frontCodeIntensityStandByType, _ = front.(*CodeIntensityStandByType)
		}
		stage.OnAfterCodeIntensityStandByTypeUpdateCallback.OnAfterUpdate(stage, codeintensitystandbytype, frontCodeIntensityStandByType)
	}
}

func (codeintensitystandbytype *CodeIntensityStandByType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeIntensityStandByTypeDeleteCallback != nil {
		var frontCodeIntensityStandByType *CodeIntensityStandByType
		if front != nil {
			frontCodeIntensityStandByType, _ = front.(*CodeIntensityStandByType)
		}
		stage.OnAfterCodeIntensityStandByTypeDeleteCallback.OnAfterDelete(stage, codeintensitystandbytype, frontCodeIntensityStandByType)
	}
}

func (codelanguagetype *CodeLanguageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLanguageTypeCreateCallback != nil {
		stage.OnAfterCodeLanguageTypeCreateCallback.OnAfterCreate(stage, codelanguagetype)
	}
}

func (codelanguagetype *CodeLanguageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLanguageTypeUpdateCallback != nil {
		var frontCodeLanguageType *CodeLanguageType
		if front != nil {
			frontCodeLanguageType, _ = front.(*CodeLanguageType)
		}
		stage.OnAfterCodeLanguageTypeUpdateCallback.OnAfterUpdate(stage, codelanguagetype, frontCodeLanguageType)
	}
}

func (codelanguagetype *CodeLanguageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLanguageTypeDeleteCallback != nil {
		var frontCodeLanguageType *CodeLanguageType
		if front != nil {
			frontCodeLanguageType, _ = front.(*CodeLanguageType)
		}
		stage.OnAfterCodeLanguageTypeDeleteCallback.OnAfterDelete(stage, codelanguagetype, frontCodeLanguageType)
	}
}

func (codelevelseriestype *CodeLevelSeriesType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLevelSeriesTypeCreateCallback != nil {
		stage.OnAfterCodeLevelSeriesTypeCreateCallback.OnAfterCreate(stage, codelevelseriestype)
	}
}

func (codelevelseriestype *CodeLevelSeriesType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelSeriesTypeUpdateCallback != nil {
		var frontCodeLevelSeriesType *CodeLevelSeriesType
		if front != nil {
			frontCodeLevelSeriesType, _ = front.(*CodeLevelSeriesType)
		}
		stage.OnAfterCodeLevelSeriesTypeUpdateCallback.OnAfterUpdate(stage, codelevelseriestype, frontCodeLevelSeriesType)
	}
}

func (codelevelseriestype *CodeLevelSeriesType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelSeriesTypeDeleteCallback != nil {
		var frontCodeLevelSeriesType *CodeLevelSeriesType
		if front != nil {
			frontCodeLevelSeriesType, _ = front.(*CodeLevelSeriesType)
		}
		stage.OnAfterCodeLevelSeriesTypeDeleteCallback.OnAfterDelete(stage, codelevelseriestype, frontCodeLevelSeriesType)
	}
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeLevelTableDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeLevelTableDesignatorTypeCreateCallback.OnAfterCreate(stage, codeleveltabledesignatortype)
	}
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelTableDesignatorTypeUpdateCallback != nil {
		var frontCodeLevelTableDesignatorType *CodeLevelTableDesignatorType
		if front != nil {
			frontCodeLevelTableDesignatorType, _ = front.(*CodeLevelTableDesignatorType)
		}
		stage.OnAfterCodeLevelTableDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codeleveltabledesignatortype, frontCodeLevelTableDesignatorType)
	}
}

func (codeleveltabledesignatortype *CodeLevelTableDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeLevelTableDesignatorTypeDeleteCallback != nil {
		var frontCodeLevelTableDesignatorType *CodeLevelTableDesignatorType
		if front != nil {
			frontCodeLevelTableDesignatorType, _ = front.(*CodeLevelTableDesignatorType)
		}
		stage.OnAfterCodeLevelTableDesignatorTypeDeleteCallback.OnAfterDelete(stage, codeleveltabledesignatortype, frontCodeLevelTableDesignatorType)
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

func (codemlsazimuthtype *CodeMLSAzimuthType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMLSAzimuthTypeCreateCallback != nil {
		stage.OnAfterCodeMLSAzimuthTypeCreateCallback.OnAfterCreate(stage, codemlsazimuthtype)
	}
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMLSAzimuthTypeUpdateCallback != nil {
		var frontCodeMLSAzimuthType *CodeMLSAzimuthType
		if front != nil {
			frontCodeMLSAzimuthType, _ = front.(*CodeMLSAzimuthType)
		}
		stage.OnAfterCodeMLSAzimuthTypeUpdateCallback.OnAfterUpdate(stage, codemlsazimuthtype, frontCodeMLSAzimuthType)
	}
}

func (codemlsazimuthtype *CodeMLSAzimuthType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMLSAzimuthTypeDeleteCallback != nil {
		var frontCodeMLSAzimuthType *CodeMLSAzimuthType
		if front != nil {
			frontCodeMLSAzimuthType, _ = front.(*CodeMLSAzimuthType)
		}
		stage.OnAfterCodeMLSAzimuthTypeDeleteCallback.OnAfterDelete(stage, codemlsazimuthtype, frontCodeMLSAzimuthType)
	}
}

func (codemlschanneltype *CodeMLSChannelType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMLSChannelTypeCreateCallback != nil {
		stage.OnAfterCodeMLSChannelTypeCreateCallback.OnAfterCreate(stage, codemlschanneltype)
	}
}

func (codemlschanneltype *CodeMLSChannelType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMLSChannelTypeUpdateCallback != nil {
		var frontCodeMLSChannelType *CodeMLSChannelType
		if front != nil {
			frontCodeMLSChannelType, _ = front.(*CodeMLSChannelType)
		}
		stage.OnAfterCodeMLSChannelTypeUpdateCallback.OnAfterUpdate(stage, codemlschanneltype, frontCodeMLSChannelType)
	}
}

func (codemlschanneltype *CodeMLSChannelType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMLSChannelTypeDeleteCallback != nil {
		var frontCodeMLSChannelType *CodeMLSChannelType
		if front != nil {
			frontCodeMLSChannelType, _ = front.(*CodeMLSChannelType)
		}
		stage.OnAfterCodeMLSChannelTypeDeleteCallback.OnAfterDelete(stage, codemlschanneltype, frontCodeMLSChannelType)
	}
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeMarkerBeaconSignalTypeCreateCallback != nil {
		stage.OnAfterCodeMarkerBeaconSignalTypeCreateCallback.OnAfterCreate(stage, codemarkerbeaconsignaltype)
	}
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMarkerBeaconSignalTypeUpdateCallback != nil {
		var frontCodeMarkerBeaconSignalType *CodeMarkerBeaconSignalType
		if front != nil {
			frontCodeMarkerBeaconSignalType, _ = front.(*CodeMarkerBeaconSignalType)
		}
		stage.OnAfterCodeMarkerBeaconSignalTypeUpdateCallback.OnAfterUpdate(stage, codemarkerbeaconsignaltype, frontCodeMarkerBeaconSignalType)
	}
}

func (codemarkerbeaconsignaltype *CodeMarkerBeaconSignalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeMarkerBeaconSignalTypeDeleteCallback != nil {
		var frontCodeMarkerBeaconSignalType *CodeMarkerBeaconSignalType
		if front != nil {
			frontCodeMarkerBeaconSignalType, _ = front.(*CodeMarkerBeaconSignalType)
		}
		stage.OnAfterCodeMarkerBeaconSignalTypeDeleteCallback.OnAfterDelete(stage, codemarkerbeaconsignaltype, frontCodeMarkerBeaconSignalType)
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

func (codendbusagetype *CodeNDBUsageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNDBUsageTypeCreateCallback != nil {
		stage.OnAfterCodeNDBUsageTypeCreateCallback.OnAfterCreate(stage, codendbusagetype)
	}
}

func (codendbusagetype *CodeNDBUsageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNDBUsageTypeUpdateCallback != nil {
		var frontCodeNDBUsageType *CodeNDBUsageType
		if front != nil {
			frontCodeNDBUsageType, _ = front.(*CodeNDBUsageType)
		}
		stage.OnAfterCodeNDBUsageTypeUpdateCallback.OnAfterUpdate(stage, codendbusagetype, frontCodeNDBUsageType)
	}
}

func (codendbusagetype *CodeNDBUsageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNDBUsageTypeDeleteCallback != nil {
		var frontCodeNDBUsageType *CodeNDBUsageType
		if front != nil {
			frontCodeNDBUsageType, _ = front.(*CodeNDBUsageType)
		}
		stage.OnAfterCodeNDBUsageTypeDeleteCallback.OnAfterDelete(stage, codendbusagetype, frontCodeNDBUsageType)
	}
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavaidDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeNavaidDesignatorTypeCreateCallback.OnAfterCreate(stage, codenavaiddesignatortype)
	}
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidDesignatorTypeUpdateCallback != nil {
		var frontCodeNavaidDesignatorType *CodeNavaidDesignatorType
		if front != nil {
			frontCodeNavaidDesignatorType, _ = front.(*CodeNavaidDesignatorType)
		}
		stage.OnAfterCodeNavaidDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codenavaiddesignatortype, frontCodeNavaidDesignatorType)
	}
}

func (codenavaiddesignatortype *CodeNavaidDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidDesignatorTypeDeleteCallback != nil {
		var frontCodeNavaidDesignatorType *CodeNavaidDesignatorType
		if front != nil {
			frontCodeNavaidDesignatorType, _ = front.(*CodeNavaidDesignatorType)
		}
		stage.OnAfterCodeNavaidDesignatorTypeDeleteCallback.OnAfterDelete(stage, codenavaiddesignatortype, frontCodeNavaidDesignatorType)
	}
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavaidPurposeTypeCreateCallback != nil {
		stage.OnAfterCodeNavaidPurposeTypeCreateCallback.OnAfterCreate(stage, codenavaidpurposetype)
	}
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidPurposeTypeUpdateCallback != nil {
		var frontCodeNavaidPurposeType *CodeNavaidPurposeType
		if front != nil {
			frontCodeNavaidPurposeType, _ = front.(*CodeNavaidPurposeType)
		}
		stage.OnAfterCodeNavaidPurposeTypeUpdateCallback.OnAfterUpdate(stage, codenavaidpurposetype, frontCodeNavaidPurposeType)
	}
}

func (codenavaidpurposetype *CodeNavaidPurposeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidPurposeTypeDeleteCallback != nil {
		var frontCodeNavaidPurposeType *CodeNavaidPurposeType
		if front != nil {
			frontCodeNavaidPurposeType, _ = front.(*CodeNavaidPurposeType)
		}
		stage.OnAfterCodeNavaidPurposeTypeDeleteCallback.OnAfterDelete(stage, codenavaidpurposetype, frontCodeNavaidPurposeType)
	}
}

func (codenavaidservicetype *CodeNavaidServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNavaidServiceTypeCreateCallback != nil {
		stage.OnAfterCodeNavaidServiceTypeCreateCallback.OnAfterCreate(stage, codenavaidservicetype)
	}
}

func (codenavaidservicetype *CodeNavaidServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidServiceTypeUpdateCallback != nil {
		var frontCodeNavaidServiceType *CodeNavaidServiceType
		if front != nil {
			frontCodeNavaidServiceType, _ = front.(*CodeNavaidServiceType)
		}
		stage.OnAfterCodeNavaidServiceTypeUpdateCallback.OnAfterUpdate(stage, codenavaidservicetype, frontCodeNavaidServiceType)
	}
}

func (codenavaidservicetype *CodeNavaidServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNavaidServiceTypeDeleteCallback != nil {
		var frontCodeNavaidServiceType *CodeNavaidServiceType
		if front != nil {
			frontCodeNavaidServiceType, _ = front.(*CodeNavaidServiceType)
		}
		stage.OnAfterCodeNavaidServiceTypeDeleteCallback.OnAfterDelete(stage, codenavaidservicetype, frontCodeNavaidServiceType)
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

func (codenitrogentype *CodeNitrogenType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNitrogenTypeCreateCallback != nil {
		stage.OnAfterCodeNitrogenTypeCreateCallback.OnAfterCreate(stage, codenitrogentype)
	}
}

func (codenitrogentype *CodeNitrogenType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNitrogenTypeUpdateCallback != nil {
		var frontCodeNitrogenType *CodeNitrogenType
		if front != nil {
			frontCodeNitrogenType, _ = front.(*CodeNitrogenType)
		}
		stage.OnAfterCodeNitrogenTypeUpdateCallback.OnAfterUpdate(stage, codenitrogentype, frontCodeNitrogenType)
	}
}

func (codenitrogentype *CodeNitrogenType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNitrogenTypeDeleteCallback != nil {
		var frontCodeNitrogenType *CodeNitrogenType
		if front != nil {
			frontCodeNitrogenType, _ = front.(*CodeNitrogenType)
		}
		stage.OnAfterCodeNitrogenTypeDeleteCallback.OnAfterDelete(stage, codenitrogentype, frontCodeNitrogenType)
	}
}

func (codenorthreferencetype *CodeNorthReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeNorthReferenceTypeCreateCallback != nil {
		stage.OnAfterCodeNorthReferenceTypeCreateCallback.OnAfterCreate(stage, codenorthreferencetype)
	}
}

func (codenorthreferencetype *CodeNorthReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNorthReferenceTypeUpdateCallback != nil {
		var frontCodeNorthReferenceType *CodeNorthReferenceType
		if front != nil {
			frontCodeNorthReferenceType, _ = front.(*CodeNorthReferenceType)
		}
		stage.OnAfterCodeNorthReferenceTypeUpdateCallback.OnAfterUpdate(stage, codenorthreferencetype, frontCodeNorthReferenceType)
	}
}

func (codenorthreferencetype *CodeNorthReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeNorthReferenceTypeDeleteCallback != nil {
		var frontCodeNorthReferenceType *CodeNorthReferenceType
		if front != nil {
			frontCodeNorthReferenceType, _ = front.(*CodeNorthReferenceType)
		}
		stage.OnAfterCodeNorthReferenceTypeDeleteCallback.OnAfterDelete(stage, codenorthreferencetype, frontCodeNorthReferenceType)
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

func (codeoiltype *CodeOilType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOilTypeCreateCallback != nil {
		stage.OnAfterCodeOilTypeCreateCallback.OnAfterCreate(stage, codeoiltype)
	}
}

func (codeoiltype *CodeOilType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOilTypeUpdateCallback != nil {
		var frontCodeOilType *CodeOilType
		if front != nil {
			frontCodeOilType, _ = front.(*CodeOilType)
		}
		stage.OnAfterCodeOilTypeUpdateCallback.OnAfterUpdate(stage, codeoiltype, frontCodeOilType)
	}
}

func (codeoiltype *CodeOilType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOilTypeDeleteCallback != nil {
		var frontCodeOilType *CodeOilType
		if front != nil {
			frontCodeOilType, _ = front.(*CodeOilType)
		}
		stage.OnAfterCodeOilTypeDeleteCallback.OnAfterDelete(stage, codeoiltype, frontCodeOilType)
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

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOrganisationDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeOrganisationDesignatorTypeCreateCallback.OnAfterCreate(stage, codeorganisationdesignatortype)
	}
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationDesignatorTypeUpdateCallback != nil {
		var frontCodeOrganisationDesignatorType *CodeOrganisationDesignatorType
		if front != nil {
			frontCodeOrganisationDesignatorType, _ = front.(*CodeOrganisationDesignatorType)
		}
		stage.OnAfterCodeOrganisationDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codeorganisationdesignatortype, frontCodeOrganisationDesignatorType)
	}
}

func (codeorganisationdesignatortype *CodeOrganisationDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationDesignatorTypeDeleteCallback != nil {
		var frontCodeOrganisationDesignatorType *CodeOrganisationDesignatorType
		if front != nil {
			frontCodeOrganisationDesignatorType, _ = front.(*CodeOrganisationDesignatorType)
		}
		stage.OnAfterCodeOrganisationDesignatorTypeDeleteCallback.OnAfterDelete(stage, codeorganisationdesignatortype, frontCodeOrganisationDesignatorType)
	}
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOrganisationHierarchyTypeCreateCallback != nil {
		stage.OnAfterCodeOrganisationHierarchyTypeCreateCallback.OnAfterCreate(stage, codeorganisationhierarchytype)
	}
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationHierarchyTypeUpdateCallback != nil {
		var frontCodeOrganisationHierarchyType *CodeOrganisationHierarchyType
		if front != nil {
			frontCodeOrganisationHierarchyType, _ = front.(*CodeOrganisationHierarchyType)
		}
		stage.OnAfterCodeOrganisationHierarchyTypeUpdateCallback.OnAfterUpdate(stage, codeorganisationhierarchytype, frontCodeOrganisationHierarchyType)
	}
}

func (codeorganisationhierarchytype *CodeOrganisationHierarchyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationHierarchyTypeDeleteCallback != nil {
		var frontCodeOrganisationHierarchyType *CodeOrganisationHierarchyType
		if front != nil {
			frontCodeOrganisationHierarchyType, _ = front.(*CodeOrganisationHierarchyType)
		}
		stage.OnAfterCodeOrganisationHierarchyTypeDeleteCallback.OnAfterDelete(stage, codeorganisationhierarchytype, frontCodeOrganisationHierarchyType)
	}
}

func (codeorganisationtype *CodeOrganisationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOrganisationTypeCreateCallback != nil {
		stage.OnAfterCodeOrganisationTypeCreateCallback.OnAfterCreate(stage, codeorganisationtype)
	}
}

func (codeorganisationtype *CodeOrganisationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationTypeUpdateCallback != nil {
		var frontCodeOrganisationType *CodeOrganisationType
		if front != nil {
			frontCodeOrganisationType, _ = front.(*CodeOrganisationType)
		}
		stage.OnAfterCodeOrganisationTypeUpdateCallback.OnAfterUpdate(stage, codeorganisationtype, frontCodeOrganisationType)
	}
}

func (codeorganisationtype *CodeOrganisationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrganisationTypeDeleteCallback != nil {
		var frontCodeOrganisationType *CodeOrganisationType
		if front != nil {
			frontCodeOrganisationType, _ = front.(*CodeOrganisationType)
		}
		stage.OnAfterCodeOrganisationTypeDeleteCallback.OnAfterDelete(stage, codeorganisationtype, frontCodeOrganisationType)
	}
}

func (codeoxygentype *CodeOxygenType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOxygenTypeCreateCallback != nil {
		stage.OnAfterCodeOxygenTypeCreateCallback.OnAfterCreate(stage, codeoxygentype)
	}
}

func (codeoxygentype *CodeOxygenType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOxygenTypeUpdateCallback != nil {
		var frontCodeOxygenType *CodeOxygenType
		if front != nil {
			frontCodeOxygenType, _ = front.(*CodeOxygenType)
		}
		stage.OnAfterCodeOxygenTypeUpdateCallback.OnAfterUpdate(stage, codeoxygentype, frontCodeOxygenType)
	}
}

func (codeoxygentype *CodeOxygenType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOxygenTypeDeleteCallback != nil {
		var frontCodeOxygenType *CodeOxygenType
		if front != nil {
			frontCodeOxygenType, _ = front.(*CodeOxygenType)
		}
		stage.OnAfterCodeOxygenTypeDeleteCallback.OnAfterDelete(stage, codeoxygentype, frontCodeOxygenType)
	}
}

func (codepartype *CodePARType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePARTypeCreateCallback != nil {
		stage.OnAfterCodePARTypeCreateCallback.OnAfterCreate(stage, codepartype)
	}
}

func (codepartype *CodePARType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePARTypeUpdateCallback != nil {
		var frontCodePARType *CodePARType
		if front != nil {
			frontCodePARType, _ = front.(*CodePARType)
		}
		stage.OnAfterCodePARTypeUpdateCallback.OnAfterUpdate(stage, codepartype, frontCodePARType)
	}
}

func (codepartype *CodePARType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePARTypeDeleteCallback != nil {
		var frontCodePARType *CodePARType
		if front != nil {
			frontCodePARType, _ = front.(*CodePARType)
		}
		stage.OnAfterCodePARTypeDeleteCallback.OnAfterDelete(stage, codepartype, frontCodePARType)
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

func (codepassengerservicetype *CodePassengerServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePassengerServiceTypeCreateCallback != nil {
		stage.OnAfterCodePassengerServiceTypeCreateCallback.OnAfterCreate(stage, codepassengerservicetype)
	}
}

func (codepassengerservicetype *CodePassengerServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePassengerServiceTypeUpdateCallback != nil {
		var frontCodePassengerServiceType *CodePassengerServiceType
		if front != nil {
			frontCodePassengerServiceType, _ = front.(*CodePassengerServiceType)
		}
		stage.OnAfterCodePassengerServiceTypeUpdateCallback.OnAfterUpdate(stage, codepassengerservicetype, frontCodePassengerServiceType)
	}
}

func (codepassengerservicetype *CodePassengerServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePassengerServiceTypeDeleteCallback != nil {
		var frontCodePassengerServiceType *CodePassengerServiceType
		if front != nil {
			frontCodePassengerServiceType, _ = front.(*CodePassengerServiceType)
		}
		stage.OnAfterCodePassengerServiceTypeDeleteCallback.OnAfterDelete(stage, codepassengerservicetype, frontCodePassengerServiceType)
	}
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePilotControlledLightingTypeCreateCallback != nil {
		stage.OnAfterCodePilotControlledLightingTypeCreateCallback.OnAfterCreate(stage, codepilotcontrolledlightingtype)
	}
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePilotControlledLightingTypeUpdateCallback != nil {
		var frontCodePilotControlledLightingType *CodePilotControlledLightingType
		if front != nil {
			frontCodePilotControlledLightingType, _ = front.(*CodePilotControlledLightingType)
		}
		stage.OnAfterCodePilotControlledLightingTypeUpdateCallback.OnAfterUpdate(stage, codepilotcontrolledlightingtype, frontCodePilotControlledLightingType)
	}
}

func (codepilotcontrolledlightingtype *CodePilotControlledLightingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePilotControlledLightingTypeDeleteCallback != nil {
		var frontCodePilotControlledLightingType *CodePilotControlledLightingType
		if front != nil {
			frontCodePilotControlledLightingType, _ = front.(*CodePilotControlledLightingType)
		}
		stage.OnAfterCodePilotControlledLightingTypeDeleteCallback.OnAfterDelete(stage, codepilotcontrolledlightingtype, frontCodePilotControlledLightingType)
	}
}

func (codepositioninilstype *CodePositionInILSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePositionInILSTypeCreateCallback != nil {
		stage.OnAfterCodePositionInILSTypeCreateCallback.OnAfterCreate(stage, codepositioninilstype)
	}
}

func (codepositioninilstype *CodePositionInILSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePositionInILSTypeUpdateCallback != nil {
		var frontCodePositionInILSType *CodePositionInILSType
		if front != nil {
			frontCodePositionInILSType, _ = front.(*CodePositionInILSType)
		}
		stage.OnAfterCodePositionInILSTypeUpdateCallback.OnAfterUpdate(stage, codepositioninilstype, frontCodePositionInILSType)
	}
}

func (codepositioninilstype *CodePositionInILSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePositionInILSTypeDeleteCallback != nil {
		var frontCodePositionInILSType *CodePositionInILSType
		if front != nil {
			frontCodePositionInILSType, _ = front.(*CodePositionInILSType)
		}
		stage.OnAfterCodePositionInILSTypeDeleteCallback.OnAfterDelete(stage, codepositioninilstype, frontCodePositionInILSType)
	}
}

func (codeprimaryradartype *CodePrimaryRadarType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodePrimaryRadarTypeCreateCallback != nil {
		stage.OnAfterCodePrimaryRadarTypeCreateCallback.OnAfterCreate(stage, codeprimaryradartype)
	}
}

func (codeprimaryradartype *CodePrimaryRadarType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePrimaryRadarTypeUpdateCallback != nil {
		var frontCodePrimaryRadarType *CodePrimaryRadarType
		if front != nil {
			frontCodePrimaryRadarType, _ = front.(*CodePrimaryRadarType)
		}
		stage.OnAfterCodePrimaryRadarTypeUpdateCallback.OnAfterUpdate(stage, codeprimaryradartype, frontCodePrimaryRadarType)
	}
}

func (codeprimaryradartype *CodePrimaryRadarType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodePrimaryRadarTypeDeleteCallback != nil {
		var frontCodePrimaryRadarType *CodePrimaryRadarType
		if front != nil {
			frontCodePrimaryRadarType, _ = front.(*CodePrimaryRadarType)
		}
		stage.OnAfterCodePrimaryRadarTypeDeleteCallback.OnAfterDelete(stage, codeprimaryradartype, frontCodePrimaryRadarType)
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

func (coderadarservicetype *CodeRadarServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRadarServiceTypeCreateCallback != nil {
		stage.OnAfterCodeRadarServiceTypeCreateCallback.OnAfterCreate(stage, coderadarservicetype)
	}
}

func (coderadarservicetype *CodeRadarServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadarServiceTypeUpdateCallback != nil {
		var frontCodeRadarServiceType *CodeRadarServiceType
		if front != nil {
			frontCodeRadarServiceType, _ = front.(*CodeRadarServiceType)
		}
		stage.OnAfterCodeRadarServiceTypeUpdateCallback.OnAfterUpdate(stage, coderadarservicetype, frontCodeRadarServiceType)
	}
}

func (coderadarservicetype *CodeRadarServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadarServiceTypeDeleteCallback != nil {
		var frontCodeRadarServiceType *CodeRadarServiceType
		if front != nil {
			frontCodeRadarServiceType, _ = front.(*CodeRadarServiceType)
		}
		stage.OnAfterCodeRadarServiceTypeDeleteCallback.OnAfterDelete(stage, coderadarservicetype, frontCodeRadarServiceType)
	}
}

func (coderadioemissiontype *CodeRadioEmissionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRadioEmissionTypeCreateCallback != nil {
		stage.OnAfterCodeRadioEmissionTypeCreateCallback.OnAfterCreate(stage, coderadioemissiontype)
	}
}

func (coderadioemissiontype *CodeRadioEmissionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioEmissionTypeUpdateCallback != nil {
		var frontCodeRadioEmissionType *CodeRadioEmissionType
		if front != nil {
			frontCodeRadioEmissionType, _ = front.(*CodeRadioEmissionType)
		}
		stage.OnAfterCodeRadioEmissionTypeUpdateCallback.OnAfterUpdate(stage, coderadioemissiontype, frontCodeRadioEmissionType)
	}
}

func (coderadioemissiontype *CodeRadioEmissionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioEmissionTypeDeleteCallback != nil {
		var frontCodeRadioEmissionType *CodeRadioEmissionType
		if front != nil {
			frontCodeRadioEmissionType, _ = front.(*CodeRadioEmissionType)
		}
		stage.OnAfterCodeRadioEmissionTypeDeleteCallback.OnAfterDelete(stage, coderadioemissiontype, frontCodeRadioEmissionType)
	}
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRadioFrequencyAreaTypeCreateCallback != nil {
		stage.OnAfterCodeRadioFrequencyAreaTypeCreateCallback.OnAfterCreate(stage, coderadiofrequencyareatype)
	}
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioFrequencyAreaTypeUpdateCallback != nil {
		var frontCodeRadioFrequencyAreaType *CodeRadioFrequencyAreaType
		if front != nil {
			frontCodeRadioFrequencyAreaType, _ = front.(*CodeRadioFrequencyAreaType)
		}
		stage.OnAfterCodeRadioFrequencyAreaTypeUpdateCallback.OnAfterUpdate(stage, coderadiofrequencyareatype, frontCodeRadioFrequencyAreaType)
	}
}

func (coderadiofrequencyareatype *CodeRadioFrequencyAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioFrequencyAreaTypeDeleteCallback != nil {
		var frontCodeRadioFrequencyAreaType *CodeRadioFrequencyAreaType
		if front != nil {
			frontCodeRadioFrequencyAreaType, _ = front.(*CodeRadioFrequencyAreaType)
		}
		stage.OnAfterCodeRadioFrequencyAreaTypeDeleteCallback.OnAfterDelete(stage, coderadiofrequencyareatype, frontCodeRadioFrequencyAreaType)
	}
}

func (coderadiosignaltype *CodeRadioSignalType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeRadioSignalTypeCreateCallback != nil {
		stage.OnAfterCodeRadioSignalTypeCreateCallback.OnAfterCreate(stage, coderadiosignaltype)
	}
}

func (coderadiosignaltype *CodeRadioSignalType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioSignalTypeUpdateCallback != nil {
		var frontCodeRadioSignalType *CodeRadioSignalType
		if front != nil {
			frontCodeRadioSignalType, _ = front.(*CodeRadioSignalType)
		}
		stage.OnAfterCodeRadioSignalTypeUpdateCallback.OnAfterUpdate(stage, coderadiosignaltype, frontCodeRadioSignalType)
	}
}

func (coderadiosignaltype *CodeRadioSignalType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeRadioSignalTypeDeleteCallback != nil {
		var frontCodeRadioSignalType *CodeRadioSignalType
		if front != nil {
			frontCodeRadioSignalType, _ = front.(*CodeRadioSignalType)
		}
		stage.OnAfterCodeRadioSignalTypeDeleteCallback.OnAfterDelete(stage, coderadiosignaltype, frontCodeRadioSignalType)
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

func (codereflectortype *CodeReflectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeReflectorTypeCreateCallback != nil {
		stage.OnAfterCodeReflectorTypeCreateCallback.OnAfterCreate(stage, codereflectortype)
	}
}

func (codereflectortype *CodeReflectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeReflectorTypeUpdateCallback != nil {
		var frontCodeReflectorType *CodeReflectorType
		if front != nil {
			frontCodeReflectorType, _ = front.(*CodeReflectorType)
		}
		stage.OnAfterCodeReflectorTypeUpdateCallback.OnAfterUpdate(stage, codereflectortype, frontCodeReflectorType)
	}
}

func (codereflectortype *CodeReflectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeReflectorTypeDeleteCallback != nil {
		var frontCodeReflectorType *CodeReflectorType
		if front != nil {
			frontCodeReflectorType, _ = front.(*CodeReflectorType)
		}
		stage.OnAfterCodeReflectorTypeDeleteCallback.OnAfterDelete(stage, codereflectortype, frontCodeReflectorType)
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

func (codeserviceatctype *CodeServiceATCType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceATCTypeCreateCallback != nil {
		stage.OnAfterCodeServiceATCTypeCreateCallback.OnAfterCreate(stage, codeserviceatctype)
	}
}

func (codeserviceatctype *CodeServiceATCType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceATCTypeUpdateCallback != nil {
		var frontCodeServiceATCType *CodeServiceATCType
		if front != nil {
			frontCodeServiceATCType, _ = front.(*CodeServiceATCType)
		}
		stage.OnAfterCodeServiceATCTypeUpdateCallback.OnAfterUpdate(stage, codeserviceatctype, frontCodeServiceATCType)
	}
}

func (codeserviceatctype *CodeServiceATCType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceATCTypeDeleteCallback != nil {
		var frontCodeServiceATCType *CodeServiceATCType
		if front != nil {
			frontCodeServiceATCType, _ = front.(*CodeServiceATCType)
		}
		stage.OnAfterCodeServiceATCTypeDeleteCallback.OnAfterDelete(stage, codeserviceatctype, frontCodeServiceATCType)
	}
}

func (codeserviceatfmtype *CodeServiceATFMType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceATFMTypeCreateCallback != nil {
		stage.OnAfterCodeServiceATFMTypeCreateCallback.OnAfterCreate(stage, codeserviceatfmtype)
	}
}

func (codeserviceatfmtype *CodeServiceATFMType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceATFMTypeUpdateCallback != nil {
		var frontCodeServiceATFMType *CodeServiceATFMType
		if front != nil {
			frontCodeServiceATFMType, _ = front.(*CodeServiceATFMType)
		}
		stage.OnAfterCodeServiceATFMTypeUpdateCallback.OnAfterUpdate(stage, codeserviceatfmtype, frontCodeServiceATFMType)
	}
}

func (codeserviceatfmtype *CodeServiceATFMType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceATFMTypeDeleteCallback != nil {
		var frontCodeServiceATFMType *CodeServiceATFMType
		if front != nil {
			frontCodeServiceATFMType, _ = front.(*CodeServiceATFMType)
		}
		stage.OnAfterCodeServiceATFMTypeDeleteCallback.OnAfterDelete(stage, codeserviceatfmtype, frontCodeServiceATFMType)
	}
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceGroundControlTypeCreateCallback != nil {
		stage.OnAfterCodeServiceGroundControlTypeCreateCallback.OnAfterCreate(stage, codeservicegroundcontroltype)
	}
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceGroundControlTypeUpdateCallback != nil {
		var frontCodeServiceGroundControlType *CodeServiceGroundControlType
		if front != nil {
			frontCodeServiceGroundControlType, _ = front.(*CodeServiceGroundControlType)
		}
		stage.OnAfterCodeServiceGroundControlTypeUpdateCallback.OnAfterUpdate(stage, codeservicegroundcontroltype, frontCodeServiceGroundControlType)
	}
}

func (codeservicegroundcontroltype *CodeServiceGroundControlType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceGroundControlTypeDeleteCallback != nil {
		var frontCodeServiceGroundControlType *CodeServiceGroundControlType
		if front != nil {
			frontCodeServiceGroundControlType, _ = front.(*CodeServiceGroundControlType)
		}
		stage.OnAfterCodeServiceGroundControlTypeDeleteCallback.OnAfterDelete(stage, codeservicegroundcontroltype, frontCodeServiceGroundControlType)
	}
}

func (codeserviceinformationtype *CodeServiceInformationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceInformationTypeCreateCallback != nil {
		stage.OnAfterCodeServiceInformationTypeCreateCallback.OnAfterCreate(stage, codeserviceinformationtype)
	}
}

func (codeserviceinformationtype *CodeServiceInformationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceInformationTypeUpdateCallback != nil {
		var frontCodeServiceInformationType *CodeServiceInformationType
		if front != nil {
			frontCodeServiceInformationType, _ = front.(*CodeServiceInformationType)
		}
		stage.OnAfterCodeServiceInformationTypeUpdateCallback.OnAfterUpdate(stage, codeserviceinformationtype, frontCodeServiceInformationType)
	}
}

func (codeserviceinformationtype *CodeServiceInformationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceInformationTypeDeleteCallback != nil {
		var frontCodeServiceInformationType *CodeServiceInformationType
		if front != nil {
			frontCodeServiceInformationType, _ = front.(*CodeServiceInformationType)
		}
		stage.OnAfterCodeServiceInformationTypeDeleteCallback.OnAfterDelete(stage, codeserviceinformationtype, frontCodeServiceInformationType)
	}
}

func (codeservicesartype *CodeServiceSARType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeServiceSARTypeCreateCallback != nil {
		stage.OnAfterCodeServiceSARTypeCreateCallback.OnAfterCreate(stage, codeservicesartype)
	}
}

func (codeservicesartype *CodeServiceSARType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceSARTypeUpdateCallback != nil {
		var frontCodeServiceSARType *CodeServiceSARType
		if front != nil {
			frontCodeServiceSARType, _ = front.(*CodeServiceSARType)
		}
		stage.OnAfterCodeServiceSARTypeUpdateCallback.OnAfterUpdate(stage, codeservicesartype, frontCodeServiceSARType)
	}
}

func (codeservicesartype *CodeServiceSARType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeServiceSARTypeDeleteCallback != nil {
		var frontCodeServiceSARType *CodeServiceSARType
		if front != nil {
			frontCodeServiceSARType, _ = front.(*CodeServiceSARType)
		}
		stage.OnAfterCodeServiceSARTypeDeleteCallback.OnAfterDelete(stage, codeservicesartype, frontCodeServiceSARType)
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

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSignalPerformanceILSTypeCreateCallback != nil {
		stage.OnAfterCodeSignalPerformanceILSTypeCreateCallback.OnAfterCreate(stage, codesignalperformanceilstype)
	}
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSignalPerformanceILSTypeUpdateCallback != nil {
		var frontCodeSignalPerformanceILSType *CodeSignalPerformanceILSType
		if front != nil {
			frontCodeSignalPerformanceILSType, _ = front.(*CodeSignalPerformanceILSType)
		}
		stage.OnAfterCodeSignalPerformanceILSTypeUpdateCallback.OnAfterUpdate(stage, codesignalperformanceilstype, frontCodeSignalPerformanceILSType)
	}
}

func (codesignalperformanceilstype *CodeSignalPerformanceILSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSignalPerformanceILSTypeDeleteCallback != nil {
		var frontCodeSignalPerformanceILSType *CodeSignalPerformanceILSType
		if front != nil {
			frontCodeSignalPerformanceILSType, _ = front.(*CodeSignalPerformanceILSType)
		}
		stage.OnAfterCodeSignalPerformanceILSTypeDeleteCallback.OnAfterDelete(stage, codesignalperformanceilstype, frontCodeSignalPerformanceILSType)
	}
}

func (codespecialdatetype *CodeSpecialDateType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpecialDateTypeCreateCallback != nil {
		stage.OnAfterCodeSpecialDateTypeCreateCallback.OnAfterCreate(stage, codespecialdatetype)
	}
}

func (codespecialdatetype *CodeSpecialDateType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialDateTypeUpdateCallback != nil {
		var frontCodeSpecialDateType *CodeSpecialDateType
		if front != nil {
			frontCodeSpecialDateType, _ = front.(*CodeSpecialDateType)
		}
		stage.OnAfterCodeSpecialDateTypeUpdateCallback.OnAfterUpdate(stage, codespecialdatetype, frontCodeSpecialDateType)
	}
}

func (codespecialdatetype *CodeSpecialDateType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialDateTypeDeleteCallback != nil {
		var frontCodeSpecialDateType *CodeSpecialDateType
		if front != nil {
			frontCodeSpecialDateType, _ = front.(*CodeSpecialDateType)
		}
		stage.OnAfterCodeSpecialDateTypeDeleteCallback.OnAfterDelete(stage, codespecialdatetype, frontCodeSpecialDateType)
	}
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpecialNavigationChainDesignatorTypeCreateCallback != nil {
		stage.OnAfterCodeSpecialNavigationChainDesignatorTypeCreateCallback.OnAfterCreate(stage, codespecialnavigationchaindesignatortype)
	}
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationChainDesignatorTypeUpdateCallback != nil {
		var frontCodeSpecialNavigationChainDesignatorType *CodeSpecialNavigationChainDesignatorType
		if front != nil {
			frontCodeSpecialNavigationChainDesignatorType, _ = front.(*CodeSpecialNavigationChainDesignatorType)
		}
		stage.OnAfterCodeSpecialNavigationChainDesignatorTypeUpdateCallback.OnAfterUpdate(stage, codespecialnavigationchaindesignatortype, frontCodeSpecialNavigationChainDesignatorType)
	}
}

func (codespecialnavigationchaindesignatortype *CodeSpecialNavigationChainDesignatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationChainDesignatorTypeDeleteCallback != nil {
		var frontCodeSpecialNavigationChainDesignatorType *CodeSpecialNavigationChainDesignatorType
		if front != nil {
			frontCodeSpecialNavigationChainDesignatorType, _ = front.(*CodeSpecialNavigationChainDesignatorType)
		}
		stage.OnAfterCodeSpecialNavigationChainDesignatorTypeDeleteCallback.OnAfterDelete(stage, codespecialnavigationchaindesignatortype, frontCodeSpecialNavigationChainDesignatorType)
	}
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpecialNavigationStationTypeCreateCallback != nil {
		stage.OnAfterCodeSpecialNavigationStationTypeCreateCallback.OnAfterCreate(stage, codespecialnavigationstationtype)
	}
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationStationTypeUpdateCallback != nil {
		var frontCodeSpecialNavigationStationType *CodeSpecialNavigationStationType
		if front != nil {
			frontCodeSpecialNavigationStationType, _ = front.(*CodeSpecialNavigationStationType)
		}
		stage.OnAfterCodeSpecialNavigationStationTypeUpdateCallback.OnAfterUpdate(stage, codespecialnavigationstationtype, frontCodeSpecialNavigationStationType)
	}
}

func (codespecialnavigationstationtype *CodeSpecialNavigationStationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationStationTypeDeleteCallback != nil {
		var frontCodeSpecialNavigationStationType *CodeSpecialNavigationStationType
		if front != nil {
			frontCodeSpecialNavigationStationType, _ = front.(*CodeSpecialNavigationStationType)
		}
		stage.OnAfterCodeSpecialNavigationStationTypeDeleteCallback.OnAfterDelete(stage, codespecialnavigationstationtype, frontCodeSpecialNavigationStationType)
	}
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSpecialNavigationSystemTypeCreateCallback != nil {
		stage.OnAfterCodeSpecialNavigationSystemTypeCreateCallback.OnAfterCreate(stage, codespecialnavigationsystemtype)
	}
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationSystemTypeUpdateCallback != nil {
		var frontCodeSpecialNavigationSystemType *CodeSpecialNavigationSystemType
		if front != nil {
			frontCodeSpecialNavigationSystemType, _ = front.(*CodeSpecialNavigationSystemType)
		}
		stage.OnAfterCodeSpecialNavigationSystemTypeUpdateCallback.OnAfterUpdate(stage, codespecialnavigationsystemtype, frontCodeSpecialNavigationSystemType)
	}
}

func (codespecialnavigationsystemtype *CodeSpecialNavigationSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSpecialNavigationSystemTypeDeleteCallback != nil {
		var frontCodeSpecialNavigationSystemType *CodeSpecialNavigationSystemType
		if front != nil {
			frontCodeSpecialNavigationSystemType, _ = front.(*CodeSpecialNavigationSystemType)
		}
		stage.OnAfterCodeSpecialNavigationSystemTypeDeleteCallback.OnAfterDelete(stage, codespecialnavigationsystemtype, frontCodeSpecialNavigationSystemType)
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

func (codestandbypowertype *CodeStandbyPowerType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStandbyPowerTypeCreateCallback != nil {
		stage.OnAfterCodeStandbyPowerTypeCreateCallback.OnAfterCreate(stage, codestandbypowertype)
	}
}

func (codestandbypowertype *CodeStandbyPowerType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStandbyPowerTypeUpdateCallback != nil {
		var frontCodeStandbyPowerType *CodeStandbyPowerType
		if front != nil {
			frontCodeStandbyPowerType, _ = front.(*CodeStandbyPowerType)
		}
		stage.OnAfterCodeStandbyPowerTypeUpdateCallback.OnAfterUpdate(stage, codestandbypowertype, frontCodeStandbyPowerType)
	}
}

func (codestandbypowertype *CodeStandbyPowerType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStandbyPowerTypeDeleteCallback != nil {
		var frontCodeStandbyPowerType *CodeStandbyPowerType
		if front != nil {
			frontCodeStandbyPowerType, _ = front.(*CodeStandbyPowerType)
		}
		stage.OnAfterCodeStandbyPowerTypeDeleteCallback.OnAfterDelete(stage, codestandbypowertype, frontCodeStandbyPowerType)
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

func (codestatusairspacetype *CodeStatusAirspaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusAirspaceTypeCreateCallback != nil {
		stage.OnAfterCodeStatusAirspaceTypeCreateCallback.OnAfterCreate(stage, codestatusairspacetype)
	}
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusAirspaceTypeUpdateCallback != nil {
		var frontCodeStatusAirspaceType *CodeStatusAirspaceType
		if front != nil {
			frontCodeStatusAirspaceType, _ = front.(*CodeStatusAirspaceType)
		}
		stage.OnAfterCodeStatusAirspaceTypeUpdateCallback.OnAfterUpdate(stage, codestatusairspacetype, frontCodeStatusAirspaceType)
	}
}

func (codestatusairspacetype *CodeStatusAirspaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusAirspaceTypeDeleteCallback != nil {
		var frontCodeStatusAirspaceType *CodeStatusAirspaceType
		if front != nil {
			frontCodeStatusAirspaceType, _ = front.(*CodeStatusAirspaceType)
		}
		stage.OnAfterCodeStatusAirspaceTypeDeleteCallback.OnAfterDelete(stage, codestatusairspacetype, frontCodeStatusAirspaceType)
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

func (codestatusnavaidtype *CodeStatusNavaidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusNavaidTypeCreateCallback != nil {
		stage.OnAfterCodeStatusNavaidTypeCreateCallback.OnAfterCreate(stage, codestatusnavaidtype)
	}
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusNavaidTypeUpdateCallback != nil {
		var frontCodeStatusNavaidType *CodeStatusNavaidType
		if front != nil {
			frontCodeStatusNavaidType, _ = front.(*CodeStatusNavaidType)
		}
		stage.OnAfterCodeStatusNavaidTypeUpdateCallback.OnAfterUpdate(stage, codestatusnavaidtype, frontCodeStatusNavaidType)
	}
}

func (codestatusnavaidtype *CodeStatusNavaidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusNavaidTypeDeleteCallback != nil {
		var frontCodeStatusNavaidType *CodeStatusNavaidType
		if front != nil {
			frontCodeStatusNavaidType, _ = front.(*CodeStatusNavaidType)
		}
		stage.OnAfterCodeStatusNavaidTypeDeleteCallback.OnAfterDelete(stage, codestatusnavaidtype, frontCodeStatusNavaidType)
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

func (codestatusservicetype *CodeStatusServiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeStatusServiceTypeCreateCallback != nil {
		stage.OnAfterCodeStatusServiceTypeCreateCallback.OnAfterCreate(stage, codestatusservicetype)
	}
}

func (codestatusservicetype *CodeStatusServiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusServiceTypeUpdateCallback != nil {
		var frontCodeStatusServiceType *CodeStatusServiceType
		if front != nil {
			frontCodeStatusServiceType, _ = front.(*CodeStatusServiceType)
		}
		stage.OnAfterCodeStatusServiceTypeUpdateCallback.OnAfterUpdate(stage, codestatusservicetype, frontCodeStatusServiceType)
	}
}

func (codestatusservicetype *CodeStatusServiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeStatusServiceTypeDeleteCallback != nil {
		var frontCodeStatusServiceType *CodeStatusServiceType
		if front != nil {
			frontCodeStatusServiceType, _ = front.(*CodeStatusServiceType)
		}
		stage.OnAfterCodeStatusServiceTypeDeleteCallback.OnAfterDelete(stage, codestatusservicetype, frontCodeStatusServiceType)
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

func (codesystemactivationtype *CodeSystemActivationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeSystemActivationTypeCreateCallback != nil {
		stage.OnAfterCodeSystemActivationTypeCreateCallback.OnAfterCreate(stage, codesystemactivationtype)
	}
}

func (codesystemactivationtype *CodeSystemActivationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSystemActivationTypeUpdateCallback != nil {
		var frontCodeSystemActivationType *CodeSystemActivationType
		if front != nil {
			frontCodeSystemActivationType, _ = front.(*CodeSystemActivationType)
		}
		stage.OnAfterCodeSystemActivationTypeUpdateCallback.OnAfterUpdate(stage, codesystemactivationtype, frontCodeSystemActivationType)
	}
}

func (codesystemactivationtype *CodeSystemActivationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeSystemActivationTypeDeleteCallback != nil {
		var frontCodeSystemActivationType *CodeSystemActivationType
		if front != nil {
			frontCodeSystemActivationType, _ = front.(*CodeSystemActivationType)
		}
		stage.OnAfterCodeSystemActivationTypeDeleteCallback.OnAfterDelete(stage, codesystemactivationtype, frontCodeSystemActivationType)
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

func (codeuuidtype *CodeUUIDType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeUUIDTypeCreateCallback != nil {
		stage.OnAfterCodeUUIDTypeCreateCallback.OnAfterCreate(stage, codeuuidtype)
	}
}

func (codeuuidtype *CodeUUIDType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUUIDTypeUpdateCallback != nil {
		var frontCodeUUIDType *CodeUUIDType
		if front != nil {
			frontCodeUUIDType, _ = front.(*CodeUUIDType)
		}
		stage.OnAfterCodeUUIDTypeUpdateCallback.OnAfterUpdate(stage, codeuuidtype, frontCodeUUIDType)
	}
}

func (codeuuidtype *CodeUUIDType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUUIDTypeDeleteCallback != nil {
		var frontCodeUUIDType *CodeUUIDType
		if front != nil {
			frontCodeUUIDType, _ = front.(*CodeUUIDType)
		}
		stage.OnAfterCodeUUIDTypeDeleteCallback.OnAfterDelete(stage, codeuuidtype, frontCodeUUIDType)
	}
}

func (codeunitdependencytype *CodeUnitDependencyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeUnitDependencyTypeCreateCallback != nil {
		stage.OnAfterCodeUnitDependencyTypeCreateCallback.OnAfterCreate(stage, codeunitdependencytype)
	}
}

func (codeunitdependencytype *CodeUnitDependencyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUnitDependencyTypeUpdateCallback != nil {
		var frontCodeUnitDependencyType *CodeUnitDependencyType
		if front != nil {
			frontCodeUnitDependencyType, _ = front.(*CodeUnitDependencyType)
		}
		stage.OnAfterCodeUnitDependencyTypeUpdateCallback.OnAfterUpdate(stage, codeunitdependencytype, frontCodeUnitDependencyType)
	}
}

func (codeunitdependencytype *CodeUnitDependencyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUnitDependencyTypeDeleteCallback != nil {
		var frontCodeUnitDependencyType *CodeUnitDependencyType
		if front != nil {
			frontCodeUnitDependencyType, _ = front.(*CodeUnitDependencyType)
		}
		stage.OnAfterCodeUnitDependencyTypeDeleteCallback.OnAfterDelete(stage, codeunitdependencytype, frontCodeUnitDependencyType)
	}
}

func (codeunittype *CodeUnitType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeUnitTypeCreateCallback != nil {
		stage.OnAfterCodeUnitTypeCreateCallback.OnAfterCreate(stage, codeunittype)
	}
}

func (codeunittype *CodeUnitType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUnitTypeUpdateCallback != nil {
		var frontCodeUnitType *CodeUnitType
		if front != nil {
			frontCodeUnitType, _ = front.(*CodeUnitType)
		}
		stage.OnAfterCodeUnitTypeUpdateCallback.OnAfterUpdate(stage, codeunittype, frontCodeUnitType)
	}
}

func (codeunittype *CodeUnitType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeUnitTypeDeleteCallback != nil {
		var frontCodeUnitType *CodeUnitType
		if front != nil {
			frontCodeUnitType, _ = front.(*CodeUnitType)
		}
		stage.OnAfterCodeUnitTypeDeleteCallback.OnAfterDelete(stage, codeunittype, frontCodeUnitType)
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

func (codevortype *CodeVORType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeVORTypeCreateCallback != nil {
		stage.OnAfterCodeVORTypeCreateCallback.OnAfterCreate(stage, codevortype)
	}
}

func (codevortype *CodeVORType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVORTypeUpdateCallback != nil {
		var frontCodeVORType *CodeVORType
		if front != nil {
			frontCodeVORType, _ = front.(*CodeVORType)
		}
		stage.OnAfterCodeVORTypeUpdateCallback.OnAfterUpdate(stage, codevortype, frontCodeVORType)
	}
}

func (codevortype *CodeVORType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeVORTypeDeleteCallback != nil {
		var frontCodeVORType *CodeVORType
		if front != nil {
			frontCodeVORType, _ = front.(*CodeVORType)
		}
		stage.OnAfterCodeVORTypeDeleteCallback.OnAfterDelete(stage, codevortype, frontCodeVORType)
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

func (textremarktype *TextRemarkType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTextRemarkTypeCreateCallback != nil {
		stage.OnAfterTextRemarkTypeCreateCallback.OnAfterCreate(stage, textremarktype)
	}
}

func (textremarktype *TextRemarkType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextRemarkTypeUpdateCallback != nil {
		var frontTextRemarkType *TextRemarkType
		if front != nil {
			frontTextRemarkType, _ = front.(*TextRemarkType)
		}
		stage.OnAfterTextRemarkTypeUpdateCallback.OnAfterUpdate(stage, textremarktype, frontTextRemarkType)
	}
}

func (textremarktype *TextRemarkType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTextRemarkTypeDeleteCallback != nil {
		var frontTextRemarkType *TextRemarkType
		if front != nil {
			frontTextRemarkType, _ = front.(*TextRemarkType)
		}
		stage.OnAfterTextRemarkTypeDeleteCallback.OnAfterDelete(stage, textremarktype, frontTextRemarkType)
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

func (valfrequencytype *ValFrequencyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValFrequencyTypeCreateCallback != nil {
		stage.OnAfterValFrequencyTypeCreateCallback.OnAfterCreate(stage, valfrequencytype)
	}
}

func (valfrequencytype *ValFrequencyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValFrequencyTypeUpdateCallback != nil {
		var frontValFrequencyType *ValFrequencyType
		if front != nil {
			frontValFrequencyType, _ = front.(*ValFrequencyType)
		}
		stage.OnAfterValFrequencyTypeUpdateCallback.OnAfterUpdate(stage, valfrequencytype, frontValFrequencyType)
	}
}

func (valfrequencytype *ValFrequencyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValFrequencyTypeDeleteCallback != nil {
		var frontValFrequencyType *ValFrequencyType
		if front != nil {
			frontValFrequencyType, _ = front.(*ValFrequencyType)
		}
		stage.OnAfterValFrequencyTypeDeleteCallback.OnAfterDelete(stage, valfrequencytype, frontValFrequencyType)
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

