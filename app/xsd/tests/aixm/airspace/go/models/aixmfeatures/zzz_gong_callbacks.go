// generated code - do not edit
package aixmfeatures

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
func (a_extension *A_extension) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extensionCreateCallback != nil {
		stage.OnAfterA_extensionCreateCallback.OnAfterCreate(stage, a_extension)
	}
}

func (a_extension *A_extension) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extensionUpdateCallback != nil {
		var frontA_extension *A_extension
		if front != nil {
			frontA_extension, _ = front.(*A_extension)
		}
		stage.OnAfterA_extensionUpdateCallback.OnAfterUpdate(stage, a_extension, frontA_extension)
	}
}

func (a_extension *A_extension) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extensionDeleteCallback != nil {
		var frontA_extension *A_extension
		if front != nil {
			frontA_extension, _ = front.(*A_extension)
		}
		stage.OnAfterA_extensionDeleteCallback.OnAfterDelete(stage, a_extension, frontA_extension)
	}
}

func (a_extension_1 *A_extension_1) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_1CreateCallback != nil {
		stage.OnAfterA_extension_1CreateCallback.OnAfterCreate(stage, a_extension_1)
	}
}

func (a_extension_1 *A_extension_1) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_1UpdateCallback != nil {
		var frontA_extension_1 *A_extension_1
		if front != nil {
			frontA_extension_1, _ = front.(*A_extension_1)
		}
		stage.OnAfterA_extension_1UpdateCallback.OnAfterUpdate(stage, a_extension_1, frontA_extension_1)
	}
}

func (a_extension_1 *A_extension_1) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_1DeleteCallback != nil {
		var frontA_extension_1 *A_extension_1
		if front != nil {
			frontA_extension_1, _ = front.(*A_extension_1)
		}
		stage.OnAfterA_extension_1DeleteCallback.OnAfterDelete(stage, a_extension_1, frontA_extension_1)
	}
}

func (a_extension_10 *A_extension_10) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_10CreateCallback != nil {
		stage.OnAfterA_extension_10CreateCallback.OnAfterCreate(stage, a_extension_10)
	}
}

func (a_extension_10 *A_extension_10) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_10UpdateCallback != nil {
		var frontA_extension_10 *A_extension_10
		if front != nil {
			frontA_extension_10, _ = front.(*A_extension_10)
		}
		stage.OnAfterA_extension_10UpdateCallback.OnAfterUpdate(stage, a_extension_10, frontA_extension_10)
	}
}

func (a_extension_10 *A_extension_10) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_10DeleteCallback != nil {
		var frontA_extension_10 *A_extension_10
		if front != nil {
			frontA_extension_10, _ = front.(*A_extension_10)
		}
		stage.OnAfterA_extension_10DeleteCallback.OnAfterDelete(stage, a_extension_10, frontA_extension_10)
	}
}

func (a_extension_11 *A_extension_11) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_11CreateCallback != nil {
		stage.OnAfterA_extension_11CreateCallback.OnAfterCreate(stage, a_extension_11)
	}
}

func (a_extension_11 *A_extension_11) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_11UpdateCallback != nil {
		var frontA_extension_11 *A_extension_11
		if front != nil {
			frontA_extension_11, _ = front.(*A_extension_11)
		}
		stage.OnAfterA_extension_11UpdateCallback.OnAfterUpdate(stage, a_extension_11, frontA_extension_11)
	}
}

func (a_extension_11 *A_extension_11) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_11DeleteCallback != nil {
		var frontA_extension_11 *A_extension_11
		if front != nil {
			frontA_extension_11, _ = front.(*A_extension_11)
		}
		stage.OnAfterA_extension_11DeleteCallback.OnAfterDelete(stage, a_extension_11, frontA_extension_11)
	}
}

func (a_extension_12 *A_extension_12) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_12CreateCallback != nil {
		stage.OnAfterA_extension_12CreateCallback.OnAfterCreate(stage, a_extension_12)
	}
}

func (a_extension_12 *A_extension_12) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_12UpdateCallback != nil {
		var frontA_extension_12 *A_extension_12
		if front != nil {
			frontA_extension_12, _ = front.(*A_extension_12)
		}
		stage.OnAfterA_extension_12UpdateCallback.OnAfterUpdate(stage, a_extension_12, frontA_extension_12)
	}
}

func (a_extension_12 *A_extension_12) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_12DeleteCallback != nil {
		var frontA_extension_12 *A_extension_12
		if front != nil {
			frontA_extension_12, _ = front.(*A_extension_12)
		}
		stage.OnAfterA_extension_12DeleteCallback.OnAfterDelete(stage, a_extension_12, frontA_extension_12)
	}
}

func (a_extension_13 *A_extension_13) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_13CreateCallback != nil {
		stage.OnAfterA_extension_13CreateCallback.OnAfterCreate(stage, a_extension_13)
	}
}

func (a_extension_13 *A_extension_13) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_13UpdateCallback != nil {
		var frontA_extension_13 *A_extension_13
		if front != nil {
			frontA_extension_13, _ = front.(*A_extension_13)
		}
		stage.OnAfterA_extension_13UpdateCallback.OnAfterUpdate(stage, a_extension_13, frontA_extension_13)
	}
}

func (a_extension_13 *A_extension_13) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_13DeleteCallback != nil {
		var frontA_extension_13 *A_extension_13
		if front != nil {
			frontA_extension_13, _ = front.(*A_extension_13)
		}
		stage.OnAfterA_extension_13DeleteCallback.OnAfterDelete(stage, a_extension_13, frontA_extension_13)
	}
}

func (a_extension_14 *A_extension_14) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_14CreateCallback != nil {
		stage.OnAfterA_extension_14CreateCallback.OnAfterCreate(stage, a_extension_14)
	}
}

func (a_extension_14 *A_extension_14) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_14UpdateCallback != nil {
		var frontA_extension_14 *A_extension_14
		if front != nil {
			frontA_extension_14, _ = front.(*A_extension_14)
		}
		stage.OnAfterA_extension_14UpdateCallback.OnAfterUpdate(stage, a_extension_14, frontA_extension_14)
	}
}

func (a_extension_14 *A_extension_14) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_14DeleteCallback != nil {
		var frontA_extension_14 *A_extension_14
		if front != nil {
			frontA_extension_14, _ = front.(*A_extension_14)
		}
		stage.OnAfterA_extension_14DeleteCallback.OnAfterDelete(stage, a_extension_14, frontA_extension_14)
	}
}

func (a_extension_2 *A_extension_2) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_2CreateCallback != nil {
		stage.OnAfterA_extension_2CreateCallback.OnAfterCreate(stage, a_extension_2)
	}
}

func (a_extension_2 *A_extension_2) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_2UpdateCallback != nil {
		var frontA_extension_2 *A_extension_2
		if front != nil {
			frontA_extension_2, _ = front.(*A_extension_2)
		}
		stage.OnAfterA_extension_2UpdateCallback.OnAfterUpdate(stage, a_extension_2, frontA_extension_2)
	}
}

func (a_extension_2 *A_extension_2) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_2DeleteCallback != nil {
		var frontA_extension_2 *A_extension_2
		if front != nil {
			frontA_extension_2, _ = front.(*A_extension_2)
		}
		stage.OnAfterA_extension_2DeleteCallback.OnAfterDelete(stage, a_extension_2, frontA_extension_2)
	}
}

func (a_extension_3 *A_extension_3) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_3CreateCallback != nil {
		stage.OnAfterA_extension_3CreateCallback.OnAfterCreate(stage, a_extension_3)
	}
}

func (a_extension_3 *A_extension_3) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_3UpdateCallback != nil {
		var frontA_extension_3 *A_extension_3
		if front != nil {
			frontA_extension_3, _ = front.(*A_extension_3)
		}
		stage.OnAfterA_extension_3UpdateCallback.OnAfterUpdate(stage, a_extension_3, frontA_extension_3)
	}
}

func (a_extension_3 *A_extension_3) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_3DeleteCallback != nil {
		var frontA_extension_3 *A_extension_3
		if front != nil {
			frontA_extension_3, _ = front.(*A_extension_3)
		}
		stage.OnAfterA_extension_3DeleteCallback.OnAfterDelete(stage, a_extension_3, frontA_extension_3)
	}
}

func (a_extension_4 *A_extension_4) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_4CreateCallback != nil {
		stage.OnAfterA_extension_4CreateCallback.OnAfterCreate(stage, a_extension_4)
	}
}

func (a_extension_4 *A_extension_4) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_4UpdateCallback != nil {
		var frontA_extension_4 *A_extension_4
		if front != nil {
			frontA_extension_4, _ = front.(*A_extension_4)
		}
		stage.OnAfterA_extension_4UpdateCallback.OnAfterUpdate(stage, a_extension_4, frontA_extension_4)
	}
}

func (a_extension_4 *A_extension_4) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_4DeleteCallback != nil {
		var frontA_extension_4 *A_extension_4
		if front != nil {
			frontA_extension_4, _ = front.(*A_extension_4)
		}
		stage.OnAfterA_extension_4DeleteCallback.OnAfterDelete(stage, a_extension_4, frontA_extension_4)
	}
}

func (a_extension_5 *A_extension_5) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_5CreateCallback != nil {
		stage.OnAfterA_extension_5CreateCallback.OnAfterCreate(stage, a_extension_5)
	}
}

func (a_extension_5 *A_extension_5) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_5UpdateCallback != nil {
		var frontA_extension_5 *A_extension_5
		if front != nil {
			frontA_extension_5, _ = front.(*A_extension_5)
		}
		stage.OnAfterA_extension_5UpdateCallback.OnAfterUpdate(stage, a_extension_5, frontA_extension_5)
	}
}

func (a_extension_5 *A_extension_5) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_5DeleteCallback != nil {
		var frontA_extension_5 *A_extension_5
		if front != nil {
			frontA_extension_5, _ = front.(*A_extension_5)
		}
		stage.OnAfterA_extension_5DeleteCallback.OnAfterDelete(stage, a_extension_5, frontA_extension_5)
	}
}

func (a_extension_6 *A_extension_6) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_6CreateCallback != nil {
		stage.OnAfterA_extension_6CreateCallback.OnAfterCreate(stage, a_extension_6)
	}
}

func (a_extension_6 *A_extension_6) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_6UpdateCallback != nil {
		var frontA_extension_6 *A_extension_6
		if front != nil {
			frontA_extension_6, _ = front.(*A_extension_6)
		}
		stage.OnAfterA_extension_6UpdateCallback.OnAfterUpdate(stage, a_extension_6, frontA_extension_6)
	}
}

func (a_extension_6 *A_extension_6) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_6DeleteCallback != nil {
		var frontA_extension_6 *A_extension_6
		if front != nil {
			frontA_extension_6, _ = front.(*A_extension_6)
		}
		stage.OnAfterA_extension_6DeleteCallback.OnAfterDelete(stage, a_extension_6, frontA_extension_6)
	}
}

func (a_extension_7 *A_extension_7) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_7CreateCallback != nil {
		stage.OnAfterA_extension_7CreateCallback.OnAfterCreate(stage, a_extension_7)
	}
}

func (a_extension_7 *A_extension_7) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_7UpdateCallback != nil {
		var frontA_extension_7 *A_extension_7
		if front != nil {
			frontA_extension_7, _ = front.(*A_extension_7)
		}
		stage.OnAfterA_extension_7UpdateCallback.OnAfterUpdate(stage, a_extension_7, frontA_extension_7)
	}
}

func (a_extension_7 *A_extension_7) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_7DeleteCallback != nil {
		var frontA_extension_7 *A_extension_7
		if front != nil {
			frontA_extension_7, _ = front.(*A_extension_7)
		}
		stage.OnAfterA_extension_7DeleteCallback.OnAfterDelete(stage, a_extension_7, frontA_extension_7)
	}
}

func (a_extension_8 *A_extension_8) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_8CreateCallback != nil {
		stage.OnAfterA_extension_8CreateCallback.OnAfterCreate(stage, a_extension_8)
	}
}

func (a_extension_8 *A_extension_8) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_8UpdateCallback != nil {
		var frontA_extension_8 *A_extension_8
		if front != nil {
			frontA_extension_8, _ = front.(*A_extension_8)
		}
		stage.OnAfterA_extension_8UpdateCallback.OnAfterUpdate(stage, a_extension_8, frontA_extension_8)
	}
}

func (a_extension_8 *A_extension_8) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_8DeleteCallback != nil {
		var frontA_extension_8 *A_extension_8
		if front != nil {
			frontA_extension_8, _ = front.(*A_extension_8)
		}
		stage.OnAfterA_extension_8DeleteCallback.OnAfterDelete(stage, a_extension_8, frontA_extension_8)
	}
}

func (a_extension_9 *A_extension_9) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_extension_9CreateCallback != nil {
		stage.OnAfterA_extension_9CreateCallback.OnAfterCreate(stage, a_extension_9)
	}
}

func (a_extension_9 *A_extension_9) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_9UpdateCallback != nil {
		var frontA_extension_9 *A_extension_9
		if front != nil {
			frontA_extension_9, _ = front.(*A_extension_9)
		}
		stage.OnAfterA_extension_9UpdateCallback.OnAfterUpdate(stage, a_extension_9, frontA_extension_9)
	}
}

func (a_extension_9 *A_extension_9) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_extension_9DeleteCallback != nil {
		var frontA_extension_9 *A_extension_9
		if front != nil {
			frontA_extension_9, _ = front.(*A_extension_9)
		}
		stage.OnAfterA_extension_9DeleteCallback.OnAfterDelete(stage, a_extension_9, frontA_extension_9)
	}
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractPropertiesWithScheduleTypeCreateCallback != nil {
		stage.OnAfterAbstractPropertiesWithScheduleTypeCreateCallback.OnAfterCreate(stage, abstractpropertieswithscheduletype)
	}
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractPropertiesWithScheduleTypeUpdateCallback != nil {
		var frontAbstractPropertiesWithScheduleType *AbstractPropertiesWithScheduleType
		if front != nil {
			frontAbstractPropertiesWithScheduleType, _ = front.(*AbstractPropertiesWithScheduleType)
		}
		stage.OnAfterAbstractPropertiesWithScheduleTypeUpdateCallback.OnAfterUpdate(stage, abstractpropertieswithscheduletype, frontAbstractPropertiesWithScheduleType)
	}
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractPropertiesWithScheduleTypeDeleteCallback != nil {
		var frontAbstractPropertiesWithScheduleType *AbstractPropertiesWithScheduleType
		if front != nil {
			frontAbstractPropertiesWithScheduleType, _ = front.(*AbstractPropertiesWithScheduleType)
		}
		stage.OnAfterAbstractPropertiesWithScheduleTypeDeleteCallback.OnAfterDelete(stage, abstractpropertieswithscheduletype, frontAbstractPropertiesWithScheduleType)
	}
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftCharacteristicPropertyTypeCreateCallback != nil {
		stage.OnAfterAircraftCharacteristicPropertyTypeCreateCallback.OnAfterCreate(stage, aircraftcharacteristicpropertytype)
	}
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftCharacteristicPropertyTypeUpdateCallback != nil {
		var frontAircraftCharacteristicPropertyType *AircraftCharacteristicPropertyType
		if front != nil {
			frontAircraftCharacteristicPropertyType, _ = front.(*AircraftCharacteristicPropertyType)
		}
		stage.OnAfterAircraftCharacteristicPropertyTypeUpdateCallback.OnAfterUpdate(stage, aircraftcharacteristicpropertytype, frontAircraftCharacteristicPropertyType)
	}
}

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftCharacteristicPropertyTypeDeleteCallback != nil {
		var frontAircraftCharacteristicPropertyType *AircraftCharacteristicPropertyType
		if front != nil {
			frontAircraftCharacteristicPropertyType, _ = front.(*AircraftCharacteristicPropertyType)
		}
		stage.OnAfterAircraftCharacteristicPropertyTypeDeleteCallback.OnAfterDelete(stage, aircraftcharacteristicpropertytype, frontAircraftCharacteristicPropertyType)
	}
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAircraftCharacteristicTypeCreateCallback != nil {
		stage.OnAfterAircraftCharacteristicTypeCreateCallback.OnAfterCreate(stage, aircraftcharacteristictype)
	}
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftCharacteristicTypeUpdateCallback != nil {
		var frontAircraftCharacteristicType *AircraftCharacteristicType
		if front != nil {
			frontAircraftCharacteristicType, _ = front.(*AircraftCharacteristicType)
		}
		stage.OnAfterAircraftCharacteristicTypeUpdateCallback.OnAfterUpdate(stage, aircraftcharacteristictype, frontAircraftCharacteristicType)
	}
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAircraftCharacteristicTypeDeleteCallback != nil {
		var frontAircraftCharacteristicType *AircraftCharacteristicType
		if front != nil {
			frontAircraftCharacteristicType, _ = front.(*AircraftCharacteristicType)
		}
		stage.OnAfterAircraftCharacteristicTypeDeleteCallback.OnAfterDelete(stage, aircraftcharacteristictype, frontAircraftCharacteristicType)
	}
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceActivationPropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceActivationPropertyTypeCreateCallback.OnAfterCreate(stage, airspaceactivationpropertytype)
	}
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceActivationPropertyTypeUpdateCallback != nil {
		var frontAirspaceActivationPropertyType *AirspaceActivationPropertyType
		if front != nil {
			frontAirspaceActivationPropertyType, _ = front.(*AirspaceActivationPropertyType)
		}
		stage.OnAfterAirspaceActivationPropertyTypeUpdateCallback.OnAfterUpdate(stage, airspaceactivationpropertytype, frontAirspaceActivationPropertyType)
	}
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceActivationPropertyTypeDeleteCallback != nil {
		var frontAirspaceActivationPropertyType *AirspaceActivationPropertyType
		if front != nil {
			frontAirspaceActivationPropertyType, _ = front.(*AirspaceActivationPropertyType)
		}
		stage.OnAfterAirspaceActivationPropertyTypeDeleteCallback.OnAfterDelete(stage, airspaceactivationpropertytype, frontAirspaceActivationPropertyType)
	}
}

func (airspaceactivationtype *AirspaceActivationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceActivationTypeCreateCallback != nil {
		stage.OnAfterAirspaceActivationTypeCreateCallback.OnAfterCreate(stage, airspaceactivationtype)
	}
}

func (airspaceactivationtype *AirspaceActivationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceActivationTypeUpdateCallback != nil {
		var frontAirspaceActivationType *AirspaceActivationType
		if front != nil {
			frontAirspaceActivationType, _ = front.(*AirspaceActivationType)
		}
		stage.OnAfterAirspaceActivationTypeUpdateCallback.OnAfterUpdate(stage, airspaceactivationtype, frontAirspaceActivationType)
	}
}

func (airspaceactivationtype *AirspaceActivationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceActivationTypeDeleteCallback != nil {
		var frontAirspaceActivationType *AirspaceActivationType
		if front != nil {
			frontAirspaceActivationType, _ = front.(*AirspaceActivationType)
		}
		stage.OnAfterAirspaceActivationTypeDeleteCallback.OnAfterDelete(stage, airspaceactivationtype, frontAirspaceActivationType)
	}
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceGeometryComponentPropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceGeometryComponentPropertyTypeCreateCallback.OnAfterCreate(stage, airspacegeometrycomponentpropertytype)
	}
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceGeometryComponentPropertyTypeUpdateCallback != nil {
		var frontAirspaceGeometryComponentPropertyType *AirspaceGeometryComponentPropertyType
		if front != nil {
			frontAirspaceGeometryComponentPropertyType, _ = front.(*AirspaceGeometryComponentPropertyType)
		}
		stage.OnAfterAirspaceGeometryComponentPropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacegeometrycomponentpropertytype, frontAirspaceGeometryComponentPropertyType)
	}
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceGeometryComponentPropertyTypeDeleteCallback != nil {
		var frontAirspaceGeometryComponentPropertyType *AirspaceGeometryComponentPropertyType
		if front != nil {
			frontAirspaceGeometryComponentPropertyType, _ = front.(*AirspaceGeometryComponentPropertyType)
		}
		stage.OnAfterAirspaceGeometryComponentPropertyTypeDeleteCallback.OnAfterDelete(stage, airspacegeometrycomponentpropertytype, frontAirspaceGeometryComponentPropertyType)
	}
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceGeometryComponentTypeCreateCallback != nil {
		stage.OnAfterAirspaceGeometryComponentTypeCreateCallback.OnAfterCreate(stage, airspacegeometrycomponenttype)
	}
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceGeometryComponentTypeUpdateCallback != nil {
		var frontAirspaceGeometryComponentType *AirspaceGeometryComponentType
		if front != nil {
			frontAirspaceGeometryComponentType, _ = front.(*AirspaceGeometryComponentType)
		}
		stage.OnAfterAirspaceGeometryComponentTypeUpdateCallback.OnAfterUpdate(stage, airspacegeometrycomponenttype, frontAirspaceGeometryComponentType)
	}
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceGeometryComponentTypeDeleteCallback != nil {
		var frontAirspaceGeometryComponentType *AirspaceGeometryComponentType
		if front != nil {
			frontAirspaceGeometryComponentType, _ = front.(*AirspaceGeometryComponentType)
		}
		stage.OnAfterAirspaceGeometryComponentTypeDeleteCallback.OnAfterDelete(stage, airspacegeometrycomponenttype, frontAirspaceGeometryComponentType)
	}
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceLayerClassPropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceLayerClassPropertyTypeCreateCallback.OnAfterCreate(stage, airspacelayerclasspropertytype)
	}
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerClassPropertyTypeUpdateCallback != nil {
		var frontAirspaceLayerClassPropertyType *AirspaceLayerClassPropertyType
		if front != nil {
			frontAirspaceLayerClassPropertyType, _ = front.(*AirspaceLayerClassPropertyType)
		}
		stage.OnAfterAirspaceLayerClassPropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacelayerclasspropertytype, frontAirspaceLayerClassPropertyType)
	}
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerClassPropertyTypeDeleteCallback != nil {
		var frontAirspaceLayerClassPropertyType *AirspaceLayerClassPropertyType
		if front != nil {
			frontAirspaceLayerClassPropertyType, _ = front.(*AirspaceLayerClassPropertyType)
		}
		stage.OnAfterAirspaceLayerClassPropertyTypeDeleteCallback.OnAfterDelete(stage, airspacelayerclasspropertytype, frontAirspaceLayerClassPropertyType)
	}
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceLayerClassTypeCreateCallback != nil {
		stage.OnAfterAirspaceLayerClassTypeCreateCallback.OnAfterCreate(stage, airspacelayerclasstype)
	}
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerClassTypeUpdateCallback != nil {
		var frontAirspaceLayerClassType *AirspaceLayerClassType
		if front != nil {
			frontAirspaceLayerClassType, _ = front.(*AirspaceLayerClassType)
		}
		stage.OnAfterAirspaceLayerClassTypeUpdateCallback.OnAfterUpdate(stage, airspacelayerclasstype, frontAirspaceLayerClassType)
	}
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerClassTypeDeleteCallback != nil {
		var frontAirspaceLayerClassType *AirspaceLayerClassType
		if front != nil {
			frontAirspaceLayerClassType, _ = front.(*AirspaceLayerClassType)
		}
		stage.OnAfterAirspaceLayerClassTypeDeleteCallback.OnAfterDelete(stage, airspacelayerclasstype, frontAirspaceLayerClassType)
	}
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceLayerPropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceLayerPropertyTypeCreateCallback.OnAfterCreate(stage, airspacelayerpropertytype)
	}
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerPropertyTypeUpdateCallback != nil {
		var frontAirspaceLayerPropertyType *AirspaceLayerPropertyType
		if front != nil {
			frontAirspaceLayerPropertyType, _ = front.(*AirspaceLayerPropertyType)
		}
		stage.OnAfterAirspaceLayerPropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacelayerpropertytype, frontAirspaceLayerPropertyType)
	}
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerPropertyTypeDeleteCallback != nil {
		var frontAirspaceLayerPropertyType *AirspaceLayerPropertyType
		if front != nil {
			frontAirspaceLayerPropertyType, _ = front.(*AirspaceLayerPropertyType)
		}
		stage.OnAfterAirspaceLayerPropertyTypeDeleteCallback.OnAfterDelete(stage, airspacelayerpropertytype, frontAirspaceLayerPropertyType)
	}
}

func (airspacelayertype *AirspaceLayerType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceLayerTypeCreateCallback != nil {
		stage.OnAfterAirspaceLayerTypeCreateCallback.OnAfterCreate(stage, airspacelayertype)
	}
}

func (airspacelayertype *AirspaceLayerType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerTypeUpdateCallback != nil {
		var frontAirspaceLayerType *AirspaceLayerType
		if front != nil {
			frontAirspaceLayerType, _ = front.(*AirspaceLayerType)
		}
		stage.OnAfterAirspaceLayerTypeUpdateCallback.OnAfterUpdate(stage, airspacelayertype, frontAirspaceLayerType)
	}
}

func (airspacelayertype *AirspaceLayerType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceLayerTypeDeleteCallback != nil {
		var frontAirspaceLayerType *AirspaceLayerType
		if front != nil {
			frontAirspaceLayerType, _ = front.(*AirspaceLayerType)
		}
		stage.OnAfterAirspaceLayerTypeDeleteCallback.OnAfterDelete(stage, airspacelayertype, frontAirspaceLayerType)
	}
}

func (airspacepropertytype *AirspacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspacePropertyTypeCreateCallback != nil {
		stage.OnAfterAirspacePropertyTypeCreateCallback.OnAfterCreate(stage, airspacepropertytype)
	}
}

func (airspacepropertytype *AirspacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspacePropertyTypeUpdateCallback != nil {
		var frontAirspacePropertyType *AirspacePropertyType
		if front != nil {
			frontAirspacePropertyType, _ = front.(*AirspacePropertyType)
		}
		stage.OnAfterAirspacePropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacepropertytype, frontAirspacePropertyType)
	}
}

func (airspacepropertytype *AirspacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspacePropertyTypeDeleteCallback != nil {
		var frontAirspacePropertyType *AirspacePropertyType
		if front != nil {
			frontAirspacePropertyType, _ = front.(*AirspacePropertyType)
		}
		stage.OnAfterAirspacePropertyTypeDeleteCallback.OnAfterDelete(stage, airspacepropertytype, frontAirspacePropertyType)
	}
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, airspacetimeslicepropertytype)
	}
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAirspaceTimeSlicePropertyType *AirspaceTimeSlicePropertyType
		if front != nil {
			frontAirspaceTimeSlicePropertyType, _ = front.(*AirspaceTimeSlicePropertyType)
		}
		stage.OnAfterAirspaceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacetimeslicepropertytype, frontAirspaceTimeSlicePropertyType)
	}
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAirspaceTimeSlicePropertyType *AirspaceTimeSlicePropertyType
		if front != nil {
			frontAirspaceTimeSlicePropertyType, _ = front.(*AirspaceTimeSlicePropertyType)
		}
		stage.OnAfterAirspaceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, airspacetimeslicepropertytype, frontAirspaceTimeSlicePropertyType)
	}
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAirspaceTimeSliceTypeCreateCallback.OnAfterCreate(stage, airspacetimeslicetype)
	}
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceTimeSliceTypeUpdateCallback != nil {
		var frontAirspaceTimeSliceType *AirspaceTimeSliceType
		if front != nil {
			frontAirspaceTimeSliceType, _ = front.(*AirspaceTimeSliceType)
		}
		stage.OnAfterAirspaceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, airspacetimeslicetype, frontAirspaceTimeSliceType)
	}
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceTimeSliceTypeDeleteCallback != nil {
		var frontAirspaceTimeSliceType *AirspaceTimeSliceType
		if front != nil {
			frontAirspaceTimeSliceType, _ = front.(*AirspaceTimeSliceType)
		}
		stage.OnAfterAirspaceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, airspacetimeslicetype, frontAirspaceTimeSliceType)
	}
}

func (airspacetype *AirspaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceTypeCreateCallback != nil {
		stage.OnAfterAirspaceTypeCreateCallback.OnAfterCreate(stage, airspacetype)
	}
}

func (airspacetype *AirspaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceTypeUpdateCallback != nil {
		var frontAirspaceType *AirspaceType
		if front != nil {
			frontAirspaceType, _ = front.(*AirspaceType)
		}
		stage.OnAfterAirspaceTypeUpdateCallback.OnAfterUpdate(stage, airspacetype, frontAirspaceType)
	}
}

func (airspacetype *AirspaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceTypeDeleteCallback != nil {
		var frontAirspaceType *AirspaceType
		if front != nil {
			frontAirspaceType, _ = front.(*AirspaceType)
		}
		stage.OnAfterAirspaceTypeDeleteCallback.OnAfterDelete(stage, airspacetype, frontAirspaceType)
	}
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceVolumeDependencyPropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceVolumeDependencyPropertyTypeCreateCallback.OnAfterCreate(stage, airspacevolumedependencypropertytype)
	}
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceVolumeDependencyPropertyTypeUpdateCallback != nil {
		var frontAirspaceVolumeDependencyPropertyType *AirspaceVolumeDependencyPropertyType
		if front != nil {
			frontAirspaceVolumeDependencyPropertyType, _ = front.(*AirspaceVolumeDependencyPropertyType)
		}
		stage.OnAfterAirspaceVolumeDependencyPropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacevolumedependencypropertytype, frontAirspaceVolumeDependencyPropertyType)
	}
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceVolumeDependencyPropertyTypeDeleteCallback != nil {
		var frontAirspaceVolumeDependencyPropertyType *AirspaceVolumeDependencyPropertyType
		if front != nil {
			frontAirspaceVolumeDependencyPropertyType, _ = front.(*AirspaceVolumeDependencyPropertyType)
		}
		stage.OnAfterAirspaceVolumeDependencyPropertyTypeDeleteCallback.OnAfterDelete(stage, airspacevolumedependencypropertytype, frontAirspaceVolumeDependencyPropertyType)
	}
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceVolumeDependencyTypeCreateCallback != nil {
		stage.OnAfterAirspaceVolumeDependencyTypeCreateCallback.OnAfterCreate(stage, airspacevolumedependencytype)
	}
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceVolumeDependencyTypeUpdateCallback != nil {
		var frontAirspaceVolumeDependencyType *AirspaceVolumeDependencyType
		if front != nil {
			frontAirspaceVolumeDependencyType, _ = front.(*AirspaceVolumeDependencyType)
		}
		stage.OnAfterAirspaceVolumeDependencyTypeUpdateCallback.OnAfterUpdate(stage, airspacevolumedependencytype, frontAirspaceVolumeDependencyType)
	}
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceVolumeDependencyTypeDeleteCallback != nil {
		var frontAirspaceVolumeDependencyType *AirspaceVolumeDependencyType
		if front != nil {
			frontAirspaceVolumeDependencyType, _ = front.(*AirspaceVolumeDependencyType)
		}
		stage.OnAfterAirspaceVolumeDependencyTypeDeleteCallback.OnAfterDelete(stage, airspacevolumedependencytype, frontAirspaceVolumeDependencyType)
	}
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceVolumePropertyTypeCreateCallback != nil {
		stage.OnAfterAirspaceVolumePropertyTypeCreateCallback.OnAfterCreate(stage, airspacevolumepropertytype)
	}
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceVolumePropertyTypeUpdateCallback != nil {
		var frontAirspaceVolumePropertyType *AirspaceVolumePropertyType
		if front != nil {
			frontAirspaceVolumePropertyType, _ = front.(*AirspaceVolumePropertyType)
		}
		stage.OnAfterAirspaceVolumePropertyTypeUpdateCallback.OnAfterUpdate(stage, airspacevolumepropertytype, frontAirspaceVolumePropertyType)
	}
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceVolumePropertyTypeDeleteCallback != nil {
		var frontAirspaceVolumePropertyType *AirspaceVolumePropertyType
		if front != nil {
			frontAirspaceVolumePropertyType, _ = front.(*AirspaceVolumePropertyType)
		}
		stage.OnAfterAirspaceVolumePropertyTypeDeleteCallback.OnAfterDelete(stage, airspacevolumepropertytype, frontAirspaceVolumePropertyType)
	}
}

func (airspacevolumetype *AirspaceVolumeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirspaceVolumeTypeCreateCallback != nil {
		stage.OnAfterAirspaceVolumeTypeCreateCallback.OnAfterCreate(stage, airspacevolumetype)
	}
}

func (airspacevolumetype *AirspaceVolumeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceVolumeTypeUpdateCallback != nil {
		var frontAirspaceVolumeType *AirspaceVolumeType
		if front != nil {
			frontAirspaceVolumeType, _ = front.(*AirspaceVolumeType)
		}
		stage.OnAfterAirspaceVolumeTypeUpdateCallback.OnAfterUpdate(stage, airspacevolumetype, frontAirspaceVolumeType)
	}
}

func (airspacevolumetype *AirspaceVolumeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirspaceVolumeTypeDeleteCallback != nil {
		var frontAirspaceVolumeType *AirspaceVolumeType
		if front != nil {
			frontAirspaceVolumeType, _ = front.(*AirspaceVolumeType)
		}
		stage.OnAfterAirspaceVolumeTypeDeleteCallback.OnAfterDelete(stage, airspacevolumetype, frontAirspaceVolumeType)
	}
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForAirspaceTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAuthorityForAirspaceTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, authorityforairspacetimeslicepropertytype)
	}
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAirspaceTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAuthorityForAirspaceTimeSlicePropertyType *AuthorityForAirspaceTimeSlicePropertyType
		if front != nil {
			frontAuthorityForAirspaceTimeSlicePropertyType, _ = front.(*AuthorityForAirspaceTimeSlicePropertyType)
		}
		stage.OnAfterAuthorityForAirspaceTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, authorityforairspacetimeslicepropertytype, frontAuthorityForAirspaceTimeSlicePropertyType)
	}
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAirspaceTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAuthorityForAirspaceTimeSlicePropertyType *AuthorityForAirspaceTimeSlicePropertyType
		if front != nil {
			frontAuthorityForAirspaceTimeSlicePropertyType, _ = front.(*AuthorityForAirspaceTimeSlicePropertyType)
		}
		stage.OnAfterAuthorityForAirspaceTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, authorityforairspacetimeslicepropertytype, frontAuthorityForAirspaceTimeSlicePropertyType)
	}
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForAirspaceTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAuthorityForAirspaceTimeSliceTypeCreateCallback.OnAfterCreate(stage, authorityforairspacetimeslicetype)
	}
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAirspaceTimeSliceTypeUpdateCallback != nil {
		var frontAuthorityForAirspaceTimeSliceType *AuthorityForAirspaceTimeSliceType
		if front != nil {
			frontAuthorityForAirspaceTimeSliceType, _ = front.(*AuthorityForAirspaceTimeSliceType)
		}
		stage.OnAfterAuthorityForAirspaceTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, authorityforairspacetimeslicetype, frontAuthorityForAirspaceTimeSliceType)
	}
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAirspaceTimeSliceTypeDeleteCallback != nil {
		var frontAuthorityForAirspaceTimeSliceType *AuthorityForAirspaceTimeSliceType
		if front != nil {
			frontAuthorityForAirspaceTimeSliceType, _ = front.(*AuthorityForAirspaceTimeSliceType)
		}
		stage.OnAfterAuthorityForAirspaceTimeSliceTypeDeleteCallback.OnAfterDelete(stage, authorityforairspacetimeslicetype, frontAuthorityForAirspaceTimeSliceType)
	}
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAuthorityForAirspaceTypeCreateCallback != nil {
		stage.OnAfterAuthorityForAirspaceTypeCreateCallback.OnAfterCreate(stage, authorityforairspacetype)
	}
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAirspaceTypeUpdateCallback != nil {
		var frontAuthorityForAirspaceType *AuthorityForAirspaceType
		if front != nil {
			frontAuthorityForAirspaceType, _ = front.(*AuthorityForAirspaceType)
		}
		stage.OnAfterAuthorityForAirspaceTypeUpdateCallback.OnAfterUpdate(stage, authorityforairspacetype, frontAuthorityForAirspaceType)
	}
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAuthorityForAirspaceTypeDeleteCallback != nil {
		var frontAuthorityForAirspaceType *AuthorityForAirspaceType
		if front != nil {
			frontAuthorityForAirspaceType, _ = front.(*AuthorityForAirspaceType)
		}
		stage.OnAfterAuthorityForAirspaceTypeDeleteCallback.OnAfterDelete(stage, authorityforairspacetype, frontAuthorityForAirspaceType)
	}
}

func (curvepropertytype *CurvePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurvePropertyTypeCreateCallback != nil {
		stage.OnAfterCurvePropertyTypeCreateCallback.OnAfterCreate(stage, curvepropertytype)
	}
}

func (curvepropertytype *CurvePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurvePropertyTypeUpdateCallback != nil {
		var frontCurvePropertyType *CurvePropertyType
		if front != nil {
			frontCurvePropertyType, _ = front.(*CurvePropertyType)
		}
		stage.OnAfterCurvePropertyTypeUpdateCallback.OnAfterUpdate(stage, curvepropertytype, frontCurvePropertyType)
	}
}

func (curvepropertytype *CurvePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurvePropertyTypeDeleteCallback != nil {
		var frontCurvePropertyType *CurvePropertyType
		if front != nil {
			frontCurvePropertyType, _ = front.(*CurvePropertyType)
		}
		stage.OnAfterCurvePropertyTypeDeleteCallback.OnAfterDelete(stage, curvepropertytype, frontCurvePropertyType)
	}
}

func (curvetype *CurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurveTypeCreateCallback != nil {
		stage.OnAfterCurveTypeCreateCallback.OnAfterCreate(stage, curvetype)
	}
}

func (curvetype *CurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveTypeUpdateCallback != nil {
		var frontCurveType *CurveType
		if front != nil {
			frontCurveType, _ = front.(*CurveType)
		}
		stage.OnAfterCurveTypeUpdateCallback.OnAfterUpdate(stage, curvetype, frontCurveType)
	}
}

func (curvetype *CurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveTypeDeleteCallback != nil {
		var frontCurveType *CurveType
		if front != nil {
			frontCurveType, _ = front.(*CurveType)
		}
		stage.OnAfterCurveTypeDeleteCallback.OnAfterDelete(stage, curvetype, frontCurveType)
	}
}

func (elevatedcurvetype *ElevatedCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedCurveTypeCreateCallback != nil {
		stage.OnAfterElevatedCurveTypeCreateCallback.OnAfterCreate(stage, elevatedcurvetype)
	}
}

func (elevatedcurvetype *ElevatedCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedCurveTypeUpdateCallback != nil {
		var frontElevatedCurveType *ElevatedCurveType
		if front != nil {
			frontElevatedCurveType, _ = front.(*ElevatedCurveType)
		}
		stage.OnAfterElevatedCurveTypeUpdateCallback.OnAfterUpdate(stage, elevatedcurvetype, frontElevatedCurveType)
	}
}

func (elevatedcurvetype *ElevatedCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedCurveTypeDeleteCallback != nil {
		var frontElevatedCurveType *ElevatedCurveType
		if front != nil {
			frontElevatedCurveType, _ = front.(*ElevatedCurveType)
		}
		stage.OnAfterElevatedCurveTypeDeleteCallback.OnAfterDelete(stage, elevatedcurvetype, frontElevatedCurveType)
	}
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedSurfaceTypeCreateCallback != nil {
		stage.OnAfterElevatedSurfaceTypeCreateCallback.OnAfterCreate(stage, elevatedsurfacetype)
	}
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedSurfaceTypeUpdateCallback != nil {
		var frontElevatedSurfaceType *ElevatedSurfaceType
		if front != nil {
			frontElevatedSurfaceType, _ = front.(*ElevatedSurfaceType)
		}
		stage.OnAfterElevatedSurfaceTypeUpdateCallback.OnAfterUpdate(stage, elevatedsurfacetype, frontElevatedSurfaceType)
	}
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedSurfaceTypeDeleteCallback != nil {
		var frontElevatedSurfaceType *ElevatedSurfaceType
		if front != nil {
			frontElevatedSurfaceType, _ = front.(*ElevatedSurfaceType)
		}
		stage.OnAfterElevatedSurfaceTypeDeleteCallback.OnAfterDelete(stage, elevatedsurfacetype, frontElevatedSurfaceType)
	}
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeoBorderTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterGeoBorderTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, geobordertimeslicepropertytype)
	}
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoBorderTimeSlicePropertyTypeUpdateCallback != nil {
		var frontGeoBorderTimeSlicePropertyType *GeoBorderTimeSlicePropertyType
		if front != nil {
			frontGeoBorderTimeSlicePropertyType, _ = front.(*GeoBorderTimeSlicePropertyType)
		}
		stage.OnAfterGeoBorderTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, geobordertimeslicepropertytype, frontGeoBorderTimeSlicePropertyType)
	}
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoBorderTimeSlicePropertyTypeDeleteCallback != nil {
		var frontGeoBorderTimeSlicePropertyType *GeoBorderTimeSlicePropertyType
		if front != nil {
			frontGeoBorderTimeSlicePropertyType, _ = front.(*GeoBorderTimeSlicePropertyType)
		}
		stage.OnAfterGeoBorderTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, geobordertimeslicepropertytype, frontGeoBorderTimeSlicePropertyType)
	}
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeoBorderTimeSliceTypeCreateCallback != nil {
		stage.OnAfterGeoBorderTimeSliceTypeCreateCallback.OnAfterCreate(stage, geobordertimeslicetype)
	}
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoBorderTimeSliceTypeUpdateCallback != nil {
		var frontGeoBorderTimeSliceType *GeoBorderTimeSliceType
		if front != nil {
			frontGeoBorderTimeSliceType, _ = front.(*GeoBorderTimeSliceType)
		}
		stage.OnAfterGeoBorderTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, geobordertimeslicetype, frontGeoBorderTimeSliceType)
	}
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoBorderTimeSliceTypeDeleteCallback != nil {
		var frontGeoBorderTimeSliceType *GeoBorderTimeSliceType
		if front != nil {
			frontGeoBorderTimeSliceType, _ = front.(*GeoBorderTimeSliceType)
		}
		stage.OnAfterGeoBorderTimeSliceTypeDeleteCallback.OnAfterDelete(stage, geobordertimeslicetype, frontGeoBorderTimeSliceType)
	}
}

func (geobordertype *GeoBorderType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeoBorderTypeCreateCallback != nil {
		stage.OnAfterGeoBorderTypeCreateCallback.OnAfterCreate(stage, geobordertype)
	}
}

func (geobordertype *GeoBorderType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoBorderTypeUpdateCallback != nil {
		var frontGeoBorderType *GeoBorderType
		if front != nil {
			frontGeoBorderType, _ = front.(*GeoBorderType)
		}
		stage.OnAfterGeoBorderTypeUpdateCallback.OnAfterUpdate(stage, geobordertype, frontGeoBorderType)
	}
}

func (geobordertype *GeoBorderType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeoBorderTypeDeleteCallback != nil {
		var frontGeoBorderType *GeoBorderType
		if front != nil {
			frontGeoBorderType, _ = front.(*GeoBorderType)
		}
		stage.OnAfterGeoBorderTypeDeleteCallback.OnAfterDelete(stage, geobordertype, frontGeoBorderType)
	}
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinguisticNotePropertyTypeCreateCallback != nil {
		stage.OnAfterLinguisticNotePropertyTypeCreateCallback.OnAfterCreate(stage, linguisticnotepropertytype)
	}
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinguisticNotePropertyTypeUpdateCallback != nil {
		var frontLinguisticNotePropertyType *LinguisticNotePropertyType
		if front != nil {
			frontLinguisticNotePropertyType, _ = front.(*LinguisticNotePropertyType)
		}
		stage.OnAfterLinguisticNotePropertyTypeUpdateCallback.OnAfterUpdate(stage, linguisticnotepropertytype, frontLinguisticNotePropertyType)
	}
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinguisticNotePropertyTypeDeleteCallback != nil {
		var frontLinguisticNotePropertyType *LinguisticNotePropertyType
		if front != nil {
			frontLinguisticNotePropertyType, _ = front.(*LinguisticNotePropertyType)
		}
		stage.OnAfterLinguisticNotePropertyTypeDeleteCallback.OnAfterDelete(stage, linguisticnotepropertytype, frontLinguisticNotePropertyType)
	}
}

func (linguisticnotetype *LinguisticNoteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinguisticNoteTypeCreateCallback != nil {
		stage.OnAfterLinguisticNoteTypeCreateCallback.OnAfterCreate(stage, linguisticnotetype)
	}
}

func (linguisticnotetype *LinguisticNoteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinguisticNoteTypeUpdateCallback != nil {
		var frontLinguisticNoteType *LinguisticNoteType
		if front != nil {
			frontLinguisticNoteType, _ = front.(*LinguisticNoteType)
		}
		stage.OnAfterLinguisticNoteTypeUpdateCallback.OnAfterUpdate(stage, linguisticnotetype, frontLinguisticNoteType)
	}
}

func (linguisticnotetype *LinguisticNoteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinguisticNoteTypeDeleteCallback != nil {
		var frontLinguisticNoteType *LinguisticNoteType
		if front != nil {
			frontLinguisticNoteType, _ = front.(*LinguisticNoteType)
		}
		stage.OnAfterLinguisticNoteTypeDeleteCallback.OnAfterDelete(stage, linguisticnotetype, frontLinguisticNoteType)
	}
}

func (notepropertytype *NotePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNotePropertyTypeCreateCallback != nil {
		stage.OnAfterNotePropertyTypeCreateCallback.OnAfterCreate(stage, notepropertytype)
	}
}

func (notepropertytype *NotePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePropertyTypeUpdateCallback != nil {
		var frontNotePropertyType *NotePropertyType
		if front != nil {
			frontNotePropertyType, _ = front.(*NotePropertyType)
		}
		stage.OnAfterNotePropertyTypeUpdateCallback.OnAfterUpdate(stage, notepropertytype, frontNotePropertyType)
	}
}

func (notepropertytype *NotePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePropertyTypeDeleteCallback != nil {
		var frontNotePropertyType *NotePropertyType
		if front != nil {
			frontNotePropertyType, _ = front.(*NotePropertyType)
		}
		stage.OnAfterNotePropertyTypeDeleteCallback.OnAfterDelete(stage, notepropertytype, frontNotePropertyType)
	}
}

func (notetype *NoteType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteTypeCreateCallback != nil {
		stage.OnAfterNoteTypeCreateCallback.OnAfterCreate(stage, notetype)
	}
}

func (notetype *NoteType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTypeUpdateCallback != nil {
		var frontNoteType *NoteType
		if front != nil {
			frontNoteType, _ = front.(*NoteType)
		}
		stage.OnAfterNoteTypeUpdateCallback.OnAfterUpdate(stage, notetype, frontNoteType)
	}
}

func (notetype *NoteType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTypeDeleteCallback != nil {
		var frontNoteType *NoteType
		if front != nil {
			frontNoteType, _ = front.(*NoteType)
		}
		stage.OnAfterNoteTypeDeleteCallback.OnAfterDelete(stage, notetype, frontNoteType)
	}
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrganisationAuthorityPropertyTypeCreateCallback != nil {
		stage.OnAfterOrganisationAuthorityPropertyTypeCreateCallback.OnAfterCreate(stage, organisationauthoritypropertytype)
	}
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityPropertyTypeUpdateCallback != nil {
		var frontOrganisationAuthorityPropertyType *OrganisationAuthorityPropertyType
		if front != nil {
			frontOrganisationAuthorityPropertyType, _ = front.(*OrganisationAuthorityPropertyType)
		}
		stage.OnAfterOrganisationAuthorityPropertyTypeUpdateCallback.OnAfterUpdate(stage, organisationauthoritypropertytype, frontOrganisationAuthorityPropertyType)
	}
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrganisationAuthorityPropertyTypeDeleteCallback != nil {
		var frontOrganisationAuthorityPropertyType *OrganisationAuthorityPropertyType
		if front != nil {
			frontOrganisationAuthorityPropertyType, _ = front.(*OrganisationAuthorityPropertyType)
		}
		stage.OnAfterOrganisationAuthorityPropertyTypeDeleteCallback.OnAfterDelete(stage, organisationauthoritypropertytype, frontOrganisationAuthorityPropertyType)
	}
}

func (routepropertytype *RoutePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRoutePropertyTypeCreateCallback != nil {
		stage.OnAfterRoutePropertyTypeCreateCallback.OnAfterCreate(stage, routepropertytype)
	}
}

func (routepropertytype *RoutePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoutePropertyTypeUpdateCallback != nil {
		var frontRoutePropertyType *RoutePropertyType
		if front != nil {
			frontRoutePropertyType, _ = front.(*RoutePropertyType)
		}
		stage.OnAfterRoutePropertyTypeUpdateCallback.OnAfterUpdate(stage, routepropertytype, frontRoutePropertyType)
	}
}

func (routepropertytype *RoutePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRoutePropertyTypeDeleteCallback != nil {
		var frontRoutePropertyType *RoutePropertyType
		if front != nil {
			frontRoutePropertyType, _ = front.(*RoutePropertyType)
		}
		stage.OnAfterRoutePropertyTypeDeleteCallback.OnAfterDelete(stage, routepropertytype, frontRoutePropertyType)
	}
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStandardLevelColumnPropertyTypeCreateCallback != nil {
		stage.OnAfterStandardLevelColumnPropertyTypeCreateCallback.OnAfterCreate(stage, standardlevelcolumnpropertytype)
	}
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnPropertyTypeUpdateCallback != nil {
		var frontStandardLevelColumnPropertyType *StandardLevelColumnPropertyType
		if front != nil {
			frontStandardLevelColumnPropertyType, _ = front.(*StandardLevelColumnPropertyType)
		}
		stage.OnAfterStandardLevelColumnPropertyTypeUpdateCallback.OnAfterUpdate(stage, standardlevelcolumnpropertytype, frontStandardLevelColumnPropertyType)
	}
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStandardLevelColumnPropertyTypeDeleteCallback != nil {
		var frontStandardLevelColumnPropertyType *StandardLevelColumnPropertyType
		if front != nil {
			frontStandardLevelColumnPropertyType, _ = front.(*StandardLevelColumnPropertyType)
		}
		stage.OnAfterStandardLevelColumnPropertyTypeDeleteCallback.OnAfterDelete(stage, standardlevelcolumnpropertytype, frontStandardLevelColumnPropertyType)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfacePropertyTypeCreateCallback != nil {
		stage.OnAfterSurfacePropertyTypeCreateCallback.OnAfterCreate(stage, surfacepropertytype)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePropertyTypeUpdateCallback != nil {
		var frontSurfacePropertyType *SurfacePropertyType
		if front != nil {
			frontSurfacePropertyType, _ = front.(*SurfacePropertyType)
		}
		stage.OnAfterSurfacePropertyTypeUpdateCallback.OnAfterUpdate(stage, surfacepropertytype, frontSurfacePropertyType)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePropertyTypeDeleteCallback != nil {
		var frontSurfacePropertyType *SurfacePropertyType
		if front != nil {
			frontSurfacePropertyType, _ = front.(*SurfacePropertyType)
		}
		stage.OnAfterSurfacePropertyTypeDeleteCallback.OnAfterDelete(stage, surfacepropertytype, frontSurfacePropertyType)
	}
}

func (surfacetype *SurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceTypeCreateCallback != nil {
		stage.OnAfterSurfaceTypeCreateCallback.OnAfterCreate(stage, surfacetype)
	}
}

func (surfacetype *SurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceTypeUpdateCallback != nil {
		var frontSurfaceType *SurfaceType
		if front != nil {
			frontSurfaceType, _ = front.(*SurfaceType)
		}
		stage.OnAfterSurfaceTypeUpdateCallback.OnAfterUpdate(stage, surfacetype, frontSurfaceType)
	}
}

func (surfacetype *SurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceTypeDeleteCallback != nil {
		var frontSurfaceType *SurfaceType
		if front != nil {
			frontSurfaceType, _ = front.(*SurfaceType)
		}
		stage.OnAfterSurfaceTypeDeleteCallback.OnAfterDelete(stage, surfacetype, frontSurfaceType)
	}
}

func (timesheetpropertytype *TimesheetPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimesheetPropertyTypeCreateCallback != nil {
		stage.OnAfterTimesheetPropertyTypeCreateCallback.OnAfterCreate(stage, timesheetpropertytype)
	}
}

func (timesheetpropertytype *TimesheetPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimesheetPropertyTypeUpdateCallback != nil {
		var frontTimesheetPropertyType *TimesheetPropertyType
		if front != nil {
			frontTimesheetPropertyType, _ = front.(*TimesheetPropertyType)
		}
		stage.OnAfterTimesheetPropertyTypeUpdateCallback.OnAfterUpdate(stage, timesheetpropertytype, frontTimesheetPropertyType)
	}
}

func (timesheetpropertytype *TimesheetPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimesheetPropertyTypeDeleteCallback != nil {
		var frontTimesheetPropertyType *TimesheetPropertyType
		if front != nil {
			frontTimesheetPropertyType, _ = front.(*TimesheetPropertyType)
		}
		stage.OnAfterTimesheetPropertyTypeDeleteCallback.OnAfterDelete(stage, timesheetpropertytype, frontTimesheetPropertyType)
	}
}

func (timesheettype *TimesheetType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimesheetTypeCreateCallback != nil {
		stage.OnAfterTimesheetTypeCreateCallback.OnAfterCreate(stage, timesheettype)
	}
}

func (timesheettype *TimesheetType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimesheetTypeUpdateCallback != nil {
		var frontTimesheetType *TimesheetType
		if front != nil {
			frontTimesheetType, _ = front.(*TimesheetType)
		}
		stage.OnAfterTimesheetTypeUpdateCallback.OnAfterUpdate(stage, timesheettype, frontTimesheetType)
	}
}

func (timesheettype *TimesheetType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimesheetTypeDeleteCallback != nil {
		var frontTimesheetType *TimesheetType
		if front != nil {
			frontTimesheetType, _ = front.(*TimesheetType)
		}
		stage.OnAfterTimesheetTypeDeleteCallback.OnAfterDelete(stage, timesheettype, frontTimesheetType)
	}
}

