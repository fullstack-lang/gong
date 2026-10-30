// generated code - do not edit
package measures

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
func (angletype *AngleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleTypeCreateCallback != nil {
		stage.OnAfterAngleTypeCreateCallback.OnAfterCreate(stage, angletype)
	}
}

func (angletype *AngleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleTypeUpdateCallback != nil {
		var frontAngleType *AngleType
		if front != nil {
			frontAngleType, _ = front.(*AngleType)
		}
		stage.OnAfterAngleTypeUpdateCallback.OnAfterUpdate(stage, angletype, frontAngleType)
	}
}

func (angletype *AngleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleTypeDeleteCallback != nil {
		var frontAngleType *AngleType
		if front != nil {
			frontAngleType, _ = front.(*AngleType)
		}
		stage.OnAfterAngleTypeDeleteCallback.OnAfterDelete(stage, angletype, frontAngleType)
	}
}

func (areatype *AreaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAreaTypeCreateCallback != nil {
		stage.OnAfterAreaTypeCreateCallback.OnAfterCreate(stage, areatype)
	}
}

func (areatype *AreaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAreaTypeUpdateCallback != nil {
		var frontAreaType *AreaType
		if front != nil {
			frontAreaType, _ = front.(*AreaType)
		}
		stage.OnAfterAreaTypeUpdateCallback.OnAfterUpdate(stage, areatype, frontAreaType)
	}
}

func (areatype *AreaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAreaTypeDeleteCallback != nil {
		var frontAreaType *AreaType
		if front != nil {
			frontAreaType, _ = front.(*AreaType)
		}
		stage.OnAfterAreaTypeDeleteCallback.OnAfterDelete(stage, areatype, frontAreaType)
	}
}

func (gridlengthtype *GridLengthType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGridLengthTypeCreateCallback != nil {
		stage.OnAfterGridLengthTypeCreateCallback.OnAfterCreate(stage, gridlengthtype)
	}
}

func (gridlengthtype *GridLengthType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridLengthTypeUpdateCallback != nil {
		var frontGridLengthType *GridLengthType
		if front != nil {
			frontGridLengthType, _ = front.(*GridLengthType)
		}
		stage.OnAfterGridLengthTypeUpdateCallback.OnAfterUpdate(stage, gridlengthtype, frontGridLengthType)
	}
}

func (gridlengthtype *GridLengthType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridLengthTypeDeleteCallback != nil {
		var frontGridLengthType *GridLengthType
		if front != nil {
			frontGridLengthType, _ = front.(*GridLengthType)
		}
		stage.OnAfterGridLengthTypeDeleteCallback.OnAfterDelete(stage, gridlengthtype, frontGridLengthType)
	}
}

func (lengthtype *LengthType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLengthTypeCreateCallback != nil {
		stage.OnAfterLengthTypeCreateCallback.OnAfterCreate(stage, lengthtype)
	}
}

func (lengthtype *LengthType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLengthTypeUpdateCallback != nil {
		var frontLengthType *LengthType
		if front != nil {
			frontLengthType, _ = front.(*LengthType)
		}
		stage.OnAfterLengthTypeUpdateCallback.OnAfterUpdate(stage, lengthtype, frontLengthType)
	}
}

func (lengthtype *LengthType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLengthTypeDeleteCallback != nil {
		var frontLengthType *LengthType
		if front != nil {
			frontLengthType, _ = front.(*LengthType)
		}
		stage.OnAfterLengthTypeDeleteCallback.OnAfterDelete(stage, lengthtype, frontLengthType)
	}
}

func (scaletype *ScaleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScaleTypeCreateCallback != nil {
		stage.OnAfterScaleTypeCreateCallback.OnAfterCreate(stage, scaletype)
	}
}

func (scaletype *ScaleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScaleTypeUpdateCallback != nil {
		var frontScaleType *ScaleType
		if front != nil {
			frontScaleType, _ = front.(*ScaleType)
		}
		stage.OnAfterScaleTypeUpdateCallback.OnAfterUpdate(stage, scaletype, frontScaleType)
	}
}

func (scaletype *ScaleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScaleTypeDeleteCallback != nil {
		var frontScaleType *ScaleType
		if front != nil {
			frontScaleType, _ = front.(*ScaleType)
		}
		stage.OnAfterScaleTypeDeleteCallback.OnAfterDelete(stage, scaletype, frontScaleType)
	}
}

func (speedtype *SpeedType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSpeedTypeCreateCallback != nil {
		stage.OnAfterSpeedTypeCreateCallback.OnAfterCreate(stage, speedtype)
	}
}

func (speedtype *SpeedType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpeedTypeUpdateCallback != nil {
		var frontSpeedType *SpeedType
		if front != nil {
			frontSpeedType, _ = front.(*SpeedType)
		}
		stage.OnAfterSpeedTypeUpdateCallback.OnAfterUpdate(stage, speedtype, frontSpeedType)
	}
}

func (speedtype *SpeedType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSpeedTypeDeleteCallback != nil {
		var frontSpeedType *SpeedType
		if front != nil {
			frontSpeedType, _ = front.(*SpeedType)
		}
		stage.OnAfterSpeedTypeDeleteCallback.OnAfterDelete(stage, speedtype, frontSpeedType)
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

func (volumetype *VolumeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVolumeTypeCreateCallback != nil {
		stage.OnAfterVolumeTypeCreateCallback.OnAfterCreate(stage, volumetype)
	}
}

func (volumetype *VolumeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVolumeTypeUpdateCallback != nil {
		var frontVolumeType *VolumeType
		if front != nil {
			frontVolumeType, _ = front.(*VolumeType)
		}
		stage.OnAfterVolumeTypeUpdateCallback.OnAfterUpdate(stage, volumetype, frontVolumeType)
	}
}

func (volumetype *VolumeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVolumeTypeDeleteCallback != nil {
		var frontVolumeType *VolumeType
		if front != nil {
			frontVolumeType, _ = front.(*VolumeType)
		}
		stage.OnAfterVolumeTypeDeleteCallback.OnAfterDelete(stage, volumetype, frontVolumeType)
	}
}

