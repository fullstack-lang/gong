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

