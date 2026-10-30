// generated code - do not edit
package basictypes

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
func (codelisttype *CodeListType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeListTypeCreateCallback != nil {
		stage.OnAfterCodeListTypeCreateCallback.OnAfterCreate(stage, codelisttype)
	}
}

func (codelisttype *CodeListType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeListTypeUpdateCallback != nil {
		var frontCodeListType *CodeListType
		if front != nil {
			frontCodeListType, _ = front.(*CodeListType)
		}
		stage.OnAfterCodeListTypeUpdateCallback.OnAfterUpdate(stage, codelisttype, frontCodeListType)
	}
}

func (codelisttype *CodeListType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeListTypeDeleteCallback != nil {
		var frontCodeListType *CodeListType
		if front != nil {
			frontCodeListType, _ = front.(*CodeListType)
		}
		stage.OnAfterCodeListTypeDeleteCallback.OnAfterDelete(stage, codelisttype, frontCodeListType)
	}
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeOrNilReasonListTypeCreateCallback != nil {
		stage.OnAfterCodeOrNilReasonListTypeCreateCallback.OnAfterCreate(stage, codeornilreasonlisttype)
	}
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrNilReasonListTypeUpdateCallback != nil {
		var frontCodeOrNilReasonListType *CodeOrNilReasonListType
		if front != nil {
			frontCodeOrNilReasonListType, _ = front.(*CodeOrNilReasonListType)
		}
		stage.OnAfterCodeOrNilReasonListTypeUpdateCallback.OnAfterUpdate(stage, codeornilreasonlisttype, frontCodeOrNilReasonListType)
	}
}

func (codeornilreasonlisttype *CodeOrNilReasonListType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeOrNilReasonListTypeDeleteCallback != nil {
		var frontCodeOrNilReasonListType *CodeOrNilReasonListType
		if front != nil {
			frontCodeOrNilReasonListType, _ = front.(*CodeOrNilReasonListType)
		}
		stage.OnAfterCodeOrNilReasonListTypeDeleteCallback.OnAfterDelete(stage, codeornilreasonlisttype, frontCodeOrNilReasonListType)
	}
}

func (codetype *CodeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeTypeCreateCallback != nil {
		stage.OnAfterCodeTypeCreateCallback.OnAfterCreate(stage, codetype)
	}
}

func (codetype *CodeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTypeUpdateCallback != nil {
		var frontCodeType *CodeType
		if front != nil {
			frontCodeType, _ = front.(*CodeType)
		}
		stage.OnAfterCodeTypeUpdateCallback.OnAfterUpdate(stage, codetype, frontCodeType)
	}
}

func (codetype *CodeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeTypeDeleteCallback != nil {
		var frontCodeType *CodeType
		if front != nil {
			frontCodeType, _ = front.(*CodeType)
		}
		stage.OnAfterCodeTypeDeleteCallback.OnAfterDelete(stage, codetype, frontCodeType)
	}
}

func (codewithauthoritytype *CodeWithAuthorityType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeWithAuthorityTypeCreateCallback != nil {
		stage.OnAfterCodeWithAuthorityTypeCreateCallback.OnAfterCreate(stage, codewithauthoritytype)
	}
}

func (codewithauthoritytype *CodeWithAuthorityType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeWithAuthorityTypeUpdateCallback != nil {
		var frontCodeWithAuthorityType *CodeWithAuthorityType
		if front != nil {
			frontCodeWithAuthorityType, _ = front.(*CodeWithAuthorityType)
		}
		stage.OnAfterCodeWithAuthorityTypeUpdateCallback.OnAfterUpdate(stage, codewithauthoritytype, frontCodeWithAuthorityType)
	}
}

func (codewithauthoritytype *CodeWithAuthorityType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeWithAuthorityTypeDeleteCallback != nil {
		var frontCodeWithAuthorityType *CodeWithAuthorityType
		if front != nil {
			frontCodeWithAuthorityType, _ = front.(*CodeWithAuthorityType)
		}
		stage.OnAfterCodeWithAuthorityTypeDeleteCallback.OnAfterDelete(stage, codewithauthoritytype, frontCodeWithAuthorityType)
	}
}

func (coordinatestype *CoordinatesType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCoordinatesTypeCreateCallback != nil {
		stage.OnAfterCoordinatesTypeCreateCallback.OnAfterCreate(stage, coordinatestype)
	}
}

func (coordinatestype *CoordinatesType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinatesTypeUpdateCallback != nil {
		var frontCoordinatesType *CoordinatesType
		if front != nil {
			frontCoordinatesType, _ = front.(*CoordinatesType)
		}
		stage.OnAfterCoordinatesTypeUpdateCallback.OnAfterUpdate(stage, coordinatestype, frontCoordinatesType)
	}
}

func (coordinatestype *CoordinatesType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinatesTypeDeleteCallback != nil {
		var frontCoordinatesType *CoordinatesType
		if front != nil {
			frontCoordinatesType, _ = front.(*CoordinatesType)
		}
		stage.OnAfterCoordinatesTypeDeleteCallback.OnAfterDelete(stage, coordinatestype, frontCoordinatesType)
	}
}

func (measurelisttype *MeasureListType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureListTypeCreateCallback != nil {
		stage.OnAfterMeasureListTypeCreateCallback.OnAfterCreate(stage, measurelisttype)
	}
}

func (measurelisttype *MeasureListType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureListTypeUpdateCallback != nil {
		var frontMeasureListType *MeasureListType
		if front != nil {
			frontMeasureListType, _ = front.(*MeasureListType)
		}
		stage.OnAfterMeasureListTypeUpdateCallback.OnAfterUpdate(stage, measurelisttype, frontMeasureListType)
	}
}

func (measurelisttype *MeasureListType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureListTypeDeleteCallback != nil {
		var frontMeasureListType *MeasureListType
		if front != nil {
			frontMeasureListType, _ = front.(*MeasureListType)
		}
		stage.OnAfterMeasureListTypeDeleteCallback.OnAfterDelete(stage, measurelisttype, frontMeasureListType)
	}
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureOrNilReasonListTypeCreateCallback != nil {
		stage.OnAfterMeasureOrNilReasonListTypeCreateCallback.OnAfterCreate(stage, measureornilreasonlisttype)
	}
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureOrNilReasonListTypeUpdateCallback != nil {
		var frontMeasureOrNilReasonListType *MeasureOrNilReasonListType
		if front != nil {
			frontMeasureOrNilReasonListType, _ = front.(*MeasureOrNilReasonListType)
		}
		stage.OnAfterMeasureOrNilReasonListTypeUpdateCallback.OnAfterUpdate(stage, measureornilreasonlisttype, frontMeasureOrNilReasonListType)
	}
}

func (measureornilreasonlisttype *MeasureOrNilReasonListType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureOrNilReasonListTypeDeleteCallback != nil {
		var frontMeasureOrNilReasonListType *MeasureOrNilReasonListType
		if front != nil {
			frontMeasureOrNilReasonListType, _ = front.(*MeasureOrNilReasonListType)
		}
		stage.OnAfterMeasureOrNilReasonListTypeDeleteCallback.OnAfterDelete(stage, measureornilreasonlisttype, frontMeasureOrNilReasonListType)
	}
}

func (measuretype *MeasureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasureTypeCreateCallback != nil {
		stage.OnAfterMeasureTypeCreateCallback.OnAfterCreate(stage, measuretype)
	}
}

func (measuretype *MeasureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureTypeUpdateCallback != nil {
		var frontMeasureType *MeasureType
		if front != nil {
			frontMeasureType, _ = front.(*MeasureType)
		}
		stage.OnAfterMeasureTypeUpdateCallback.OnAfterUpdate(stage, measuretype, frontMeasureType)
	}
}

func (measuretype *MeasureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasureTypeDeleteCallback != nil {
		var frontMeasureType *MeasureType
		if front != nil {
			frontMeasureType, _ = front.(*MeasureType)
		}
		stage.OnAfterMeasureTypeDeleteCallback.OnAfterDelete(stage, measuretype, frontMeasureType)
	}
}

