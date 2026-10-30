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

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAeronauticalGroundLightTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterAeronauticalGroundLightTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, aeronauticalgroundlighttimeslicepropertytype)
	}
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAeronauticalGroundLightTimeSlicePropertyTypeUpdateCallback != nil {
		var frontAeronauticalGroundLightTimeSlicePropertyType *AeronauticalGroundLightTimeSlicePropertyType
		if front != nil {
			frontAeronauticalGroundLightTimeSlicePropertyType, _ = front.(*AeronauticalGroundLightTimeSlicePropertyType)
		}
		stage.OnAfterAeronauticalGroundLightTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, aeronauticalgroundlighttimeslicepropertytype, frontAeronauticalGroundLightTimeSlicePropertyType)
	}
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAeronauticalGroundLightTimeSlicePropertyTypeDeleteCallback != nil {
		var frontAeronauticalGroundLightTimeSlicePropertyType *AeronauticalGroundLightTimeSlicePropertyType
		if front != nil {
			frontAeronauticalGroundLightTimeSlicePropertyType, _ = front.(*AeronauticalGroundLightTimeSlicePropertyType)
		}
		stage.OnAfterAeronauticalGroundLightTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, aeronauticalgroundlighttimeslicepropertytype, frontAeronauticalGroundLightTimeSlicePropertyType)
	}
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAeronauticalGroundLightTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAeronauticalGroundLightTimeSliceTypeCreateCallback.OnAfterCreate(stage, aeronauticalgroundlighttimeslicetype)
	}
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAeronauticalGroundLightTimeSliceTypeUpdateCallback != nil {
		var frontAeronauticalGroundLightTimeSliceType *AeronauticalGroundLightTimeSliceType
		if front != nil {
			frontAeronauticalGroundLightTimeSliceType, _ = front.(*AeronauticalGroundLightTimeSliceType)
		}
		stage.OnAfterAeronauticalGroundLightTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, aeronauticalgroundlighttimeslicetype, frontAeronauticalGroundLightTimeSliceType)
	}
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAeronauticalGroundLightTimeSliceTypeDeleteCallback != nil {
		var frontAeronauticalGroundLightTimeSliceType *AeronauticalGroundLightTimeSliceType
		if front != nil {
			frontAeronauticalGroundLightTimeSliceType, _ = front.(*AeronauticalGroundLightTimeSliceType)
		}
		stage.OnAfterAeronauticalGroundLightTimeSliceTypeDeleteCallback.OnAfterDelete(stage, aeronauticalgroundlighttimeslicetype, frontAeronauticalGroundLightTimeSliceType)
	}
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAeronauticalGroundLightTypeCreateCallback != nil {
		stage.OnAfterAeronauticalGroundLightTypeCreateCallback.OnAfterCreate(stage, aeronauticalgroundlighttype)
	}
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAeronauticalGroundLightTypeUpdateCallback != nil {
		var frontAeronauticalGroundLightType *AeronauticalGroundLightType
		if front != nil {
			frontAeronauticalGroundLightType, _ = front.(*AeronauticalGroundLightType)
		}
		stage.OnAfterAeronauticalGroundLightTypeUpdateCallback.OnAfterUpdate(stage, aeronauticalgroundlighttype, frontAeronauticalGroundLightType)
	}
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAeronauticalGroundLightTypeDeleteCallback != nil {
		var frontAeronauticalGroundLightType *AeronauticalGroundLightType
		if front != nil {
			frontAeronauticalGroundLightType, _ = front.(*AeronauticalGroundLightType)
		}
		stage.OnAfterAeronauticalGroundLightTypeDeleteCallback.OnAfterDelete(stage, aeronauticalgroundlighttype, frontAeronauticalGroundLightType)
	}
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAirportHeliportPropertyTypeCreateCallback != nil {
		stage.OnAfterAirportHeliportPropertyTypeCreateCallback.OnAfterCreate(stage, airportheliportpropertytype)
	}
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportPropertyTypeUpdateCallback != nil {
		var frontAirportHeliportPropertyType *AirportHeliportPropertyType
		if front != nil {
			frontAirportHeliportPropertyType, _ = front.(*AirportHeliportPropertyType)
		}
		stage.OnAfterAirportHeliportPropertyTypeUpdateCallback.OnAfterUpdate(stage, airportheliportpropertytype, frontAirportHeliportPropertyType)
	}
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAirportHeliportPropertyTypeDeleteCallback != nil {
		var frontAirportHeliportPropertyType *AirportHeliportPropertyType
		if front != nil {
			frontAirportHeliportPropertyType, _ = front.(*AirportHeliportPropertyType)
		}
		stage.OnAfterAirportHeliportPropertyTypeDeleteCallback.OnAfterDelete(stage, airportheliportpropertytype, frontAirportHeliportPropertyType)
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

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedCurvePropertyTypeCreateCallback != nil {
		stage.OnAfterElevatedCurvePropertyTypeCreateCallback.OnAfterCreate(stage, elevatedcurvepropertytype)
	}
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedCurvePropertyTypeUpdateCallback != nil {
		var frontElevatedCurvePropertyType *ElevatedCurvePropertyType
		if front != nil {
			frontElevatedCurvePropertyType, _ = front.(*ElevatedCurvePropertyType)
		}
		stage.OnAfterElevatedCurvePropertyTypeUpdateCallback.OnAfterUpdate(stage, elevatedcurvepropertytype, frontElevatedCurvePropertyType)
	}
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedCurvePropertyTypeDeleteCallback != nil {
		var frontElevatedCurvePropertyType *ElevatedCurvePropertyType
		if front != nil {
			frontElevatedCurvePropertyType, _ = front.(*ElevatedCurvePropertyType)
		}
		stage.OnAfterElevatedCurvePropertyTypeDeleteCallback.OnAfterDelete(stage, elevatedcurvepropertytype, frontElevatedCurvePropertyType)
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

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedPointPropertyTypeCreateCallback != nil {
		stage.OnAfterElevatedPointPropertyTypeCreateCallback.OnAfterCreate(stage, elevatedpointpropertytype)
	}
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedPointPropertyTypeUpdateCallback != nil {
		var frontElevatedPointPropertyType *ElevatedPointPropertyType
		if front != nil {
			frontElevatedPointPropertyType, _ = front.(*ElevatedPointPropertyType)
		}
		stage.OnAfterElevatedPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, elevatedpointpropertytype, frontElevatedPointPropertyType)
	}
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedPointPropertyTypeDeleteCallback != nil {
		var frontElevatedPointPropertyType *ElevatedPointPropertyType
		if front != nil {
			frontElevatedPointPropertyType, _ = front.(*ElevatedPointPropertyType)
		}
		stage.OnAfterElevatedPointPropertyTypeDeleteCallback.OnAfterDelete(stage, elevatedpointpropertytype, frontElevatedPointPropertyType)
	}
}

func (elevatedpointtype *ElevatedPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedPointTypeCreateCallback != nil {
		stage.OnAfterElevatedPointTypeCreateCallback.OnAfterCreate(stage, elevatedpointtype)
	}
}

func (elevatedpointtype *ElevatedPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedPointTypeUpdateCallback != nil {
		var frontElevatedPointType *ElevatedPointType
		if front != nil {
			frontElevatedPointType, _ = front.(*ElevatedPointType)
		}
		stage.OnAfterElevatedPointTypeUpdateCallback.OnAfterUpdate(stage, elevatedpointtype, frontElevatedPointType)
	}
}

func (elevatedpointtype *ElevatedPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedPointTypeDeleteCallback != nil {
		var frontElevatedPointType *ElevatedPointType
		if front != nil {
			frontElevatedPointType, _ = front.(*ElevatedPointType)
		}
		stage.OnAfterElevatedPointTypeDeleteCallback.OnAfterDelete(stage, elevatedpointtype, frontElevatedPointType)
	}
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterElevatedSurfacePropertyTypeCreateCallback != nil {
		stage.OnAfterElevatedSurfacePropertyTypeCreateCallback.OnAfterCreate(stage, elevatedsurfacepropertytype)
	}
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedSurfacePropertyTypeUpdateCallback != nil {
		var frontElevatedSurfacePropertyType *ElevatedSurfacePropertyType
		if front != nil {
			frontElevatedSurfacePropertyType, _ = front.(*ElevatedSurfacePropertyType)
		}
		stage.OnAfterElevatedSurfacePropertyTypeUpdateCallback.OnAfterUpdate(stage, elevatedsurfacepropertytype, frontElevatedSurfacePropertyType)
	}
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterElevatedSurfacePropertyTypeDeleteCallback != nil {
		var frontElevatedSurfacePropertyType *ElevatedSurfacePropertyType
		if front != nil {
			frontElevatedSurfacePropertyType, _ = front.(*ElevatedSurfacePropertyType)
		}
		stage.OnAfterElevatedSurfacePropertyTypeDeleteCallback.OnAfterDelete(stage, elevatedsurfacepropertytype, frontElevatedSurfacePropertyType)
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

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroundLightSystemPropertyTypeCreateCallback != nil {
		stage.OnAfterGroundLightSystemPropertyTypeCreateCallback.OnAfterCreate(stage, groundlightsystempropertytype)
	}
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundLightSystemPropertyTypeUpdateCallback != nil {
		var frontGroundLightSystemPropertyType *GroundLightSystemPropertyType
		if front != nil {
			frontGroundLightSystemPropertyType, _ = front.(*GroundLightSystemPropertyType)
		}
		stage.OnAfterGroundLightSystemPropertyTypeUpdateCallback.OnAfterUpdate(stage, groundlightsystempropertytype, frontGroundLightSystemPropertyType)
	}
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroundLightSystemPropertyTypeDeleteCallback != nil {
		var frontGroundLightSystemPropertyType *GroundLightSystemPropertyType
		if front != nil {
			frontGroundLightSystemPropertyType, _ = front.(*GroundLightSystemPropertyType)
		}
		stage.OnAfterGroundLightSystemPropertyTypeDeleteCallback.OnAfterDelete(stage, groundlightsystempropertytype, frontGroundLightSystemPropertyType)
	}
}

func (lightelementpropertytype *LightElementPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightElementPropertyTypeCreateCallback != nil {
		stage.OnAfterLightElementPropertyTypeCreateCallback.OnAfterCreate(stage, lightelementpropertytype)
	}
}

func (lightelementpropertytype *LightElementPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementPropertyTypeUpdateCallback != nil {
		var frontLightElementPropertyType *LightElementPropertyType
		if front != nil {
			frontLightElementPropertyType, _ = front.(*LightElementPropertyType)
		}
		stage.OnAfterLightElementPropertyTypeUpdateCallback.OnAfterUpdate(stage, lightelementpropertytype, frontLightElementPropertyType)
	}
}

func (lightelementpropertytype *LightElementPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementPropertyTypeDeleteCallback != nil {
		var frontLightElementPropertyType *LightElementPropertyType
		if front != nil {
			frontLightElementPropertyType, _ = front.(*LightElementPropertyType)
		}
		stage.OnAfterLightElementPropertyTypeDeleteCallback.OnAfterDelete(stage, lightelementpropertytype, frontLightElementPropertyType)
	}
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightElementStatusPropertyTypeCreateCallback != nil {
		stage.OnAfterLightElementStatusPropertyTypeCreateCallback.OnAfterCreate(stage, lightelementstatuspropertytype)
	}
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementStatusPropertyTypeUpdateCallback != nil {
		var frontLightElementStatusPropertyType *LightElementStatusPropertyType
		if front != nil {
			frontLightElementStatusPropertyType, _ = front.(*LightElementStatusPropertyType)
		}
		stage.OnAfterLightElementStatusPropertyTypeUpdateCallback.OnAfterUpdate(stage, lightelementstatuspropertytype, frontLightElementStatusPropertyType)
	}
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementStatusPropertyTypeDeleteCallback != nil {
		var frontLightElementStatusPropertyType *LightElementStatusPropertyType
		if front != nil {
			frontLightElementStatusPropertyType, _ = front.(*LightElementStatusPropertyType)
		}
		stage.OnAfterLightElementStatusPropertyTypeDeleteCallback.OnAfterDelete(stage, lightelementstatuspropertytype, frontLightElementStatusPropertyType)
	}
}

func (lightelementstatustype *LightElementStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightElementStatusTypeCreateCallback != nil {
		stage.OnAfterLightElementStatusTypeCreateCallback.OnAfterCreate(stage, lightelementstatustype)
	}
}

func (lightelementstatustype *LightElementStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementStatusTypeUpdateCallback != nil {
		var frontLightElementStatusType *LightElementStatusType
		if front != nil {
			frontLightElementStatusType, _ = front.(*LightElementStatusType)
		}
		stage.OnAfterLightElementStatusTypeUpdateCallback.OnAfterUpdate(stage, lightelementstatustype, frontLightElementStatusType)
	}
}

func (lightelementstatustype *LightElementStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementStatusTypeDeleteCallback != nil {
		var frontLightElementStatusType *LightElementStatusType
		if front != nil {
			frontLightElementStatusType, _ = front.(*LightElementStatusType)
		}
		stage.OnAfterLightElementStatusTypeDeleteCallback.OnAfterDelete(stage, lightelementstatustype, frontLightElementStatusType)
	}
}

func (lightelementtype *LightElementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLightElementTypeCreateCallback != nil {
		stage.OnAfterLightElementTypeCreateCallback.OnAfterCreate(stage, lightelementtype)
	}
}

func (lightelementtype *LightElementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementTypeUpdateCallback != nil {
		var frontLightElementType *LightElementType
		if front != nil {
			frontLightElementType, _ = front.(*LightElementType)
		}
		stage.OnAfterLightElementTypeUpdateCallback.OnAfterUpdate(stage, lightelementtype, frontLightElementType)
	}
}

func (lightelementtype *LightElementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLightElementTypeDeleteCallback != nil {
		var frontLightElementType *LightElementType
		if front != nil {
			frontLightElementType, _ = front.(*LightElementType)
		}
		stage.OnAfterLightElementTypeDeleteCallback.OnAfterDelete(stage, lightelementtype, frontLightElementType)
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

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMarkerBeaconPropertyTypeCreateCallback != nil {
		stage.OnAfterMarkerBeaconPropertyTypeCreateCallback.OnAfterCreate(stage, markerbeaconpropertytype)
	}
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerBeaconPropertyTypeUpdateCallback != nil {
		var frontMarkerBeaconPropertyType *MarkerBeaconPropertyType
		if front != nil {
			frontMarkerBeaconPropertyType, _ = front.(*MarkerBeaconPropertyType)
		}
		stage.OnAfterMarkerBeaconPropertyTypeUpdateCallback.OnAfterUpdate(stage, markerbeaconpropertytype, frontMarkerBeaconPropertyType)
	}
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMarkerBeaconPropertyTypeDeleteCallback != nil {
		var frontMarkerBeaconPropertyType *MarkerBeaconPropertyType
		if front != nil {
			frontMarkerBeaconPropertyType, _ = front.(*MarkerBeaconPropertyType)
		}
		stage.OnAfterMarkerBeaconPropertyTypeDeleteCallback.OnAfterDelete(stage, markerbeaconpropertytype, frontMarkerBeaconPropertyType)
	}
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNavaidEquipmentPropertyTypeCreateCallback != nil {
		stage.OnAfterNavaidEquipmentPropertyTypeCreateCallback.OnAfterCreate(stage, navaidequipmentpropertytype)
	}
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentPropertyTypeUpdateCallback != nil {
		var frontNavaidEquipmentPropertyType *NavaidEquipmentPropertyType
		if front != nil {
			frontNavaidEquipmentPropertyType, _ = front.(*NavaidEquipmentPropertyType)
		}
		stage.OnAfterNavaidEquipmentPropertyTypeUpdateCallback.OnAfterUpdate(stage, navaidequipmentpropertytype, frontNavaidEquipmentPropertyType)
	}
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNavaidEquipmentPropertyTypeDeleteCallback != nil {
		var frontNavaidEquipmentPropertyType *NavaidEquipmentPropertyType
		if front != nil {
			frontNavaidEquipmentPropertyType, _ = front.(*NavaidEquipmentPropertyType)
		}
		stage.OnAfterNavaidEquipmentPropertyTypeDeleteCallback.OnAfterDelete(stage, navaidequipmentpropertytype, frontNavaidEquipmentPropertyType)
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

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstacleAreaTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterObstacleAreaTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, obstacleareatimeslicepropertytype)
	}
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAreaTimeSlicePropertyTypeUpdateCallback != nil {
		var frontObstacleAreaTimeSlicePropertyType *ObstacleAreaTimeSlicePropertyType
		if front != nil {
			frontObstacleAreaTimeSlicePropertyType, _ = front.(*ObstacleAreaTimeSlicePropertyType)
		}
		stage.OnAfterObstacleAreaTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, obstacleareatimeslicepropertytype, frontObstacleAreaTimeSlicePropertyType)
	}
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAreaTimeSlicePropertyTypeDeleteCallback != nil {
		var frontObstacleAreaTimeSlicePropertyType *ObstacleAreaTimeSlicePropertyType
		if front != nil {
			frontObstacleAreaTimeSlicePropertyType, _ = front.(*ObstacleAreaTimeSlicePropertyType)
		}
		stage.OnAfterObstacleAreaTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, obstacleareatimeslicepropertytype, frontObstacleAreaTimeSlicePropertyType)
	}
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstacleAreaTimeSliceTypeCreateCallback != nil {
		stage.OnAfterObstacleAreaTimeSliceTypeCreateCallback.OnAfterCreate(stage, obstacleareatimeslicetype)
	}
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAreaTimeSliceTypeUpdateCallback != nil {
		var frontObstacleAreaTimeSliceType *ObstacleAreaTimeSliceType
		if front != nil {
			frontObstacleAreaTimeSliceType, _ = front.(*ObstacleAreaTimeSliceType)
		}
		stage.OnAfterObstacleAreaTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, obstacleareatimeslicetype, frontObstacleAreaTimeSliceType)
	}
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAreaTimeSliceTypeDeleteCallback != nil {
		var frontObstacleAreaTimeSliceType *ObstacleAreaTimeSliceType
		if front != nil {
			frontObstacleAreaTimeSliceType, _ = front.(*ObstacleAreaTimeSliceType)
		}
		stage.OnAfterObstacleAreaTimeSliceTypeDeleteCallback.OnAfterDelete(stage, obstacleareatimeslicetype, frontObstacleAreaTimeSliceType)
	}
}

func (obstacleareatype *ObstacleAreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObstacleAreaTypeCreateCallback != nil {
		stage.OnAfterObstacleAreaTypeCreateCallback.OnAfterCreate(stage, obstacleareatype)
	}
}

func (obstacleareatype *ObstacleAreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAreaTypeUpdateCallback != nil {
		var frontObstacleAreaType *ObstacleAreaType
		if front != nil {
			frontObstacleAreaType, _ = front.(*ObstacleAreaType)
		}
		stage.OnAfterObstacleAreaTypeUpdateCallback.OnAfterUpdate(stage, obstacleareatype, frontObstacleAreaType)
	}
}

func (obstacleareatype *ObstacleAreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObstacleAreaTypeDeleteCallback != nil {
		var frontObstacleAreaType *ObstacleAreaType
		if front != nil {
			frontObstacleAreaType, _ = front.(*ObstacleAreaType)
		}
		stage.OnAfterObstacleAreaTypeDeleteCallback.OnAfterDelete(stage, obstacleareatype, frontObstacleAreaType)
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

func (passengerservicepropertytype *PassengerServicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassengerServicePropertyTypeCreateCallback != nil {
		stage.OnAfterPassengerServicePropertyTypeCreateCallback.OnAfterCreate(stage, passengerservicepropertytype)
	}
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerServicePropertyTypeUpdateCallback != nil {
		var frontPassengerServicePropertyType *PassengerServicePropertyType
		if front != nil {
			frontPassengerServicePropertyType, _ = front.(*PassengerServicePropertyType)
		}
		stage.OnAfterPassengerServicePropertyTypeUpdateCallback.OnAfterUpdate(stage, passengerservicepropertytype, frontPassengerServicePropertyType)
	}
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassengerServicePropertyTypeDeleteCallback != nil {
		var frontPassengerServicePropertyType *PassengerServicePropertyType
		if front != nil {
			frontPassengerServicePropertyType, _ = front.(*PassengerServicePropertyType)
		}
		stage.OnAfterPassengerServicePropertyTypeDeleteCallback.OnAfterDelete(stage, passengerservicepropertytype, frontPassengerServicePropertyType)
	}
}

func (pointtype *PointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointTypeCreateCallback != nil {
		stage.OnAfterPointTypeCreateCallback.OnAfterCreate(stage, pointtype)
	}
}

func (pointtype *PointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointTypeUpdateCallback != nil {
		var frontPointType *PointType
		if front != nil {
			frontPointType, _ = front.(*PointType)
		}
		stage.OnAfterPointTypeUpdateCallback.OnAfterUpdate(stage, pointtype, frontPointType)
	}
}

func (pointtype *PointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointTypeDeleteCallback != nil {
		var frontPointType *PointType
		if front != nil {
			frontPointType, _ = front.(*PointType)
		}
		stage.OnAfterPointTypeDeleteCallback.OnAfterDelete(stage, pointtype, frontPointType)
	}
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRunwayDirectionPropertyTypeCreateCallback != nil {
		stage.OnAfterRunwayDirectionPropertyTypeCreateCallback.OnAfterCreate(stage, runwaydirectionpropertytype)
	}
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionPropertyTypeUpdateCallback != nil {
		var frontRunwayDirectionPropertyType *RunwayDirectionPropertyType
		if front != nil {
			frontRunwayDirectionPropertyType, _ = front.(*RunwayDirectionPropertyType)
		}
		stage.OnAfterRunwayDirectionPropertyTypeUpdateCallback.OnAfterUpdate(stage, runwaydirectionpropertytype, frontRunwayDirectionPropertyType)
	}
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRunwayDirectionPropertyTypeDeleteCallback != nil {
		var frontRunwayDirectionPropertyType *RunwayDirectionPropertyType
		if front != nil {
			frontRunwayDirectionPropertyType, _ = front.(*RunwayDirectionPropertyType)
		}
		stage.OnAfterRunwayDirectionPropertyTypeDeleteCallback.OnAfterDelete(stage, runwaydirectionpropertytype, frontRunwayDirectionPropertyType)
	}
}

func (servicepropertytype *ServicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterServicePropertyTypeCreateCallback != nil {
		stage.OnAfterServicePropertyTypeCreateCallback.OnAfterCreate(stage, servicepropertytype)
	}
}

func (servicepropertytype *ServicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterServicePropertyTypeUpdateCallback != nil {
		var frontServicePropertyType *ServicePropertyType
		if front != nil {
			frontServicePropertyType, _ = front.(*ServicePropertyType)
		}
		stage.OnAfterServicePropertyTypeUpdateCallback.OnAfterUpdate(stage, servicepropertytype, frontServicePropertyType)
	}
}

func (servicepropertytype *ServicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterServicePropertyTypeDeleteCallback != nil {
		var frontServicePropertyType *ServicePropertyType
		if front != nil {
			frontServicePropertyType, _ = front.(*ServicePropertyType)
		}
		stage.OnAfterServicePropertyTypeDeleteCallback.OnAfterDelete(stage, servicepropertytype, frontServicePropertyType)
	}
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpecialNavigationStationPropertyTypeCreateCallback != nil {
		stage.OnAfterSpecialNavigationStationPropertyTypeCreateCallback.OnAfterCreate(stage, specialnavigationstationpropertytype)
	}
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationPropertyTypeUpdateCallback != nil {
		var frontSpecialNavigationStationPropertyType *SpecialNavigationStationPropertyType
		if front != nil {
			frontSpecialNavigationStationPropertyType, _ = front.(*SpecialNavigationStationPropertyType)
		}
		stage.OnAfterSpecialNavigationStationPropertyTypeUpdateCallback.OnAfterUpdate(stage, specialnavigationstationpropertytype, frontSpecialNavigationStationPropertyType)
	}
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpecialNavigationStationPropertyTypeDeleteCallback != nil {
		var frontSpecialNavigationStationPropertyType *SpecialNavigationStationPropertyType
		if front != nil {
			frontSpecialNavigationStationPropertyType, _ = front.(*SpecialNavigationStationPropertyType)
		}
		stage.OnAfterSpecialNavigationStationPropertyTypeDeleteCallback.OnAfterDelete(stage, specialnavigationstationpropertytype, frontSpecialNavigationStationPropertyType)
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

func (unitpropertytype *UnitPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitPropertyTypeCreateCallback != nil {
		stage.OnAfterUnitPropertyTypeCreateCallback.OnAfterCreate(stage, unitpropertytype)
	}
}

func (unitpropertytype *UnitPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitPropertyTypeUpdateCallback != nil {
		var frontUnitPropertyType *UnitPropertyType
		if front != nil {
			frontUnitPropertyType, _ = front.(*UnitPropertyType)
		}
		stage.OnAfterUnitPropertyTypeUpdateCallback.OnAfterUpdate(stage, unitpropertytype, frontUnitPropertyType)
	}
}

func (unitpropertytype *UnitPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitPropertyTypeDeleteCallback != nil {
		var frontUnitPropertyType *UnitPropertyType
		if front != nil {
			frontUnitPropertyType, _ = front.(*UnitPropertyType)
		}
		stage.OnAfterUnitPropertyTypeDeleteCallback.OnAfterDelete(stage, unitpropertytype, frontUnitPropertyType)
	}
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructureLightingStatusPropertyTypeCreateCallback != nil {
		stage.OnAfterVerticalStructureLightingStatusPropertyTypeCreateCallback.OnAfterCreate(stage, verticalstructurelightingstatuspropertytype)
	}
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureLightingStatusPropertyTypeUpdateCallback != nil {
		var frontVerticalStructureLightingStatusPropertyType *VerticalStructureLightingStatusPropertyType
		if front != nil {
			frontVerticalStructureLightingStatusPropertyType, _ = front.(*VerticalStructureLightingStatusPropertyType)
		}
		stage.OnAfterVerticalStructureLightingStatusPropertyTypeUpdateCallback.OnAfterUpdate(stage, verticalstructurelightingstatuspropertytype, frontVerticalStructureLightingStatusPropertyType)
	}
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureLightingStatusPropertyTypeDeleteCallback != nil {
		var frontVerticalStructureLightingStatusPropertyType *VerticalStructureLightingStatusPropertyType
		if front != nil {
			frontVerticalStructureLightingStatusPropertyType, _ = front.(*VerticalStructureLightingStatusPropertyType)
		}
		stage.OnAfterVerticalStructureLightingStatusPropertyTypeDeleteCallback.OnAfterDelete(stage, verticalstructurelightingstatuspropertytype, frontVerticalStructureLightingStatusPropertyType)
	}
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructureLightingStatusTypeCreateCallback != nil {
		stage.OnAfterVerticalStructureLightingStatusTypeCreateCallback.OnAfterCreate(stage, verticalstructurelightingstatustype)
	}
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureLightingStatusTypeUpdateCallback != nil {
		var frontVerticalStructureLightingStatusType *VerticalStructureLightingStatusType
		if front != nil {
			frontVerticalStructureLightingStatusType, _ = front.(*VerticalStructureLightingStatusType)
		}
		stage.OnAfterVerticalStructureLightingStatusTypeUpdateCallback.OnAfterUpdate(stage, verticalstructurelightingstatustype, frontVerticalStructureLightingStatusType)
	}
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureLightingStatusTypeDeleteCallback != nil {
		var frontVerticalStructureLightingStatusType *VerticalStructureLightingStatusType
		if front != nil {
			frontVerticalStructureLightingStatusType, _ = front.(*VerticalStructureLightingStatusType)
		}
		stage.OnAfterVerticalStructureLightingStatusTypeDeleteCallback.OnAfterDelete(stage, verticalstructurelightingstatustype, frontVerticalStructureLightingStatusType)
	}
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructurePartPropertyTypeCreateCallback != nil {
		stage.OnAfterVerticalStructurePartPropertyTypeCreateCallback.OnAfterCreate(stage, verticalstructurepartpropertytype)
	}
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructurePartPropertyTypeUpdateCallback != nil {
		var frontVerticalStructurePartPropertyType *VerticalStructurePartPropertyType
		if front != nil {
			frontVerticalStructurePartPropertyType, _ = front.(*VerticalStructurePartPropertyType)
		}
		stage.OnAfterVerticalStructurePartPropertyTypeUpdateCallback.OnAfterUpdate(stage, verticalstructurepartpropertytype, frontVerticalStructurePartPropertyType)
	}
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructurePartPropertyTypeDeleteCallback != nil {
		var frontVerticalStructurePartPropertyType *VerticalStructurePartPropertyType
		if front != nil {
			frontVerticalStructurePartPropertyType, _ = front.(*VerticalStructurePartPropertyType)
		}
		stage.OnAfterVerticalStructurePartPropertyTypeDeleteCallback.OnAfterDelete(stage, verticalstructurepartpropertytype, frontVerticalStructurePartPropertyType)
	}
}

func (verticalstructureparttype *VerticalStructurePartType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructurePartTypeCreateCallback != nil {
		stage.OnAfterVerticalStructurePartTypeCreateCallback.OnAfterCreate(stage, verticalstructureparttype)
	}
}

func (verticalstructureparttype *VerticalStructurePartType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructurePartTypeUpdateCallback != nil {
		var frontVerticalStructurePartType *VerticalStructurePartType
		if front != nil {
			frontVerticalStructurePartType, _ = front.(*VerticalStructurePartType)
		}
		stage.OnAfterVerticalStructurePartTypeUpdateCallback.OnAfterUpdate(stage, verticalstructureparttype, frontVerticalStructurePartType)
	}
}

func (verticalstructureparttype *VerticalStructurePartType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructurePartTypeDeleteCallback != nil {
		var frontVerticalStructurePartType *VerticalStructurePartType
		if front != nil {
			frontVerticalStructurePartType, _ = front.(*VerticalStructurePartType)
		}
		stage.OnAfterVerticalStructurePartTypeDeleteCallback.OnAfterDelete(stage, verticalstructureparttype, frontVerticalStructurePartType)
	}
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructurePropertyTypeCreateCallback != nil {
		stage.OnAfterVerticalStructurePropertyTypeCreateCallback.OnAfterCreate(stage, verticalstructurepropertytype)
	}
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructurePropertyTypeUpdateCallback != nil {
		var frontVerticalStructurePropertyType *VerticalStructurePropertyType
		if front != nil {
			frontVerticalStructurePropertyType, _ = front.(*VerticalStructurePropertyType)
		}
		stage.OnAfterVerticalStructurePropertyTypeUpdateCallback.OnAfterUpdate(stage, verticalstructurepropertytype, frontVerticalStructurePropertyType)
	}
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructurePropertyTypeDeleteCallback != nil {
		var frontVerticalStructurePropertyType *VerticalStructurePropertyType
		if front != nil {
			frontVerticalStructurePropertyType, _ = front.(*VerticalStructurePropertyType)
		}
		stage.OnAfterVerticalStructurePropertyTypeDeleteCallback.OnAfterDelete(stage, verticalstructurepropertytype, frontVerticalStructurePropertyType)
	}
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructureTimeSlicePropertyTypeCreateCallback != nil {
		stage.OnAfterVerticalStructureTimeSlicePropertyTypeCreateCallback.OnAfterCreate(stage, verticalstructuretimeslicepropertytype)
	}
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureTimeSlicePropertyTypeUpdateCallback != nil {
		var frontVerticalStructureTimeSlicePropertyType *VerticalStructureTimeSlicePropertyType
		if front != nil {
			frontVerticalStructureTimeSlicePropertyType, _ = front.(*VerticalStructureTimeSlicePropertyType)
		}
		stage.OnAfterVerticalStructureTimeSlicePropertyTypeUpdateCallback.OnAfterUpdate(stage, verticalstructuretimeslicepropertytype, frontVerticalStructureTimeSlicePropertyType)
	}
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureTimeSlicePropertyTypeDeleteCallback != nil {
		var frontVerticalStructureTimeSlicePropertyType *VerticalStructureTimeSlicePropertyType
		if front != nil {
			frontVerticalStructureTimeSlicePropertyType, _ = front.(*VerticalStructureTimeSlicePropertyType)
		}
		stage.OnAfterVerticalStructureTimeSlicePropertyTypeDeleteCallback.OnAfterDelete(stage, verticalstructuretimeslicepropertytype, frontVerticalStructureTimeSlicePropertyType)
	}
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructureTimeSliceTypeCreateCallback != nil {
		stage.OnAfterVerticalStructureTimeSliceTypeCreateCallback.OnAfterCreate(stage, verticalstructuretimeslicetype)
	}
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureTimeSliceTypeUpdateCallback != nil {
		var frontVerticalStructureTimeSliceType *VerticalStructureTimeSliceType
		if front != nil {
			frontVerticalStructureTimeSliceType, _ = front.(*VerticalStructureTimeSliceType)
		}
		stage.OnAfterVerticalStructureTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, verticalstructuretimeslicetype, frontVerticalStructureTimeSliceType)
	}
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureTimeSliceTypeDeleteCallback != nil {
		var frontVerticalStructureTimeSliceType *VerticalStructureTimeSliceType
		if front != nil {
			frontVerticalStructureTimeSliceType, _ = front.(*VerticalStructureTimeSliceType)
		}
		stage.OnAfterVerticalStructureTimeSliceTypeDeleteCallback.OnAfterDelete(stage, verticalstructuretimeslicetype, frontVerticalStructureTimeSliceType)
	}
}

func (verticalstructuretype *VerticalStructureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalStructureTypeCreateCallback != nil {
		stage.OnAfterVerticalStructureTypeCreateCallback.OnAfterCreate(stage, verticalstructuretype)
	}
}

func (verticalstructuretype *VerticalStructureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureTypeUpdateCallback != nil {
		var frontVerticalStructureType *VerticalStructureType
		if front != nil {
			frontVerticalStructureType, _ = front.(*VerticalStructureType)
		}
		stage.OnAfterVerticalStructureTypeUpdateCallback.OnAfterUpdate(stage, verticalstructuretype, frontVerticalStructureType)
	}
}

func (verticalstructuretype *VerticalStructureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalStructureTypeDeleteCallback != nil {
		var frontVerticalStructureType *VerticalStructureType
		if front != nil {
			frontVerticalStructureType, _ = front.(*VerticalStructureType)
		}
		stage.OnAfterVerticalStructureTypeDeleteCallback.OnAfterDelete(stage, verticalstructuretype, frontVerticalStructureType)
	}
}

