// generated code - do not edit
package models

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
func (gongbasicfield *GongBasicField) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongBasicFieldCreateCallback != nil {
		stage.OnAfterGongBasicFieldCreateCallback.OnAfterCreate(stage, gongbasicfield)
	}
}

func (gongbasicfield *GongBasicField) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongBasicFieldUpdateCallback != nil {
		var frontGongBasicField *GongBasicField
		if front != nil {
			frontGongBasicField, _ = front.(*GongBasicField)
		}
		stage.OnAfterGongBasicFieldUpdateCallback.OnAfterUpdate(stage, gongbasicfield, frontGongBasicField)
	}
}

func (gongbasicfield *GongBasicField) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongBasicFieldDeleteCallback != nil {
		var frontGongBasicField *GongBasicField
		if front != nil {
			frontGongBasicField, _ = front.(*GongBasicField)
		}
		stage.OnAfterGongBasicFieldDeleteCallback.OnAfterDelete(stage, gongbasicfield, frontGongBasicField)
	}
}

func (gongenum *GongEnum) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongEnumCreateCallback != nil {
		stage.OnAfterGongEnumCreateCallback.OnAfterCreate(stage, gongenum)
	}
}

func (gongenum *GongEnum) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongEnumUpdateCallback != nil {
		var frontGongEnum *GongEnum
		if front != nil {
			frontGongEnum, _ = front.(*GongEnum)
		}
		stage.OnAfterGongEnumUpdateCallback.OnAfterUpdate(stage, gongenum, frontGongEnum)
	}
}

func (gongenum *GongEnum) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongEnumDeleteCallback != nil {
		var frontGongEnum *GongEnum
		if front != nil {
			frontGongEnum, _ = front.(*GongEnum)
		}
		stage.OnAfterGongEnumDeleteCallback.OnAfterDelete(stage, gongenum, frontGongEnum)
	}
}

func (gongenumvalue *GongEnumValue) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongEnumValueCreateCallback != nil {
		stage.OnAfterGongEnumValueCreateCallback.OnAfterCreate(stage, gongenumvalue)
	}
}

func (gongenumvalue *GongEnumValue) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongEnumValueUpdateCallback != nil {
		var frontGongEnumValue *GongEnumValue
		if front != nil {
			frontGongEnumValue, _ = front.(*GongEnumValue)
		}
		stage.OnAfterGongEnumValueUpdateCallback.OnAfterUpdate(stage, gongenumvalue, frontGongEnumValue)
	}
}

func (gongenumvalue *GongEnumValue) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongEnumValueDeleteCallback != nil {
		var frontGongEnumValue *GongEnumValue
		if front != nil {
			frontGongEnumValue, _ = front.(*GongEnumValue)
		}
		stage.OnAfterGongEnumValueDeleteCallback.OnAfterDelete(stage, gongenumvalue, frontGongEnumValue)
	}
}

func (gonglink *GongLink) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongLinkCreateCallback != nil {
		stage.OnAfterGongLinkCreateCallback.OnAfterCreate(stage, gonglink)
	}
}

func (gonglink *GongLink) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongLinkUpdateCallback != nil {
		var frontGongLink *GongLink
		if front != nil {
			frontGongLink, _ = front.(*GongLink)
		}
		stage.OnAfterGongLinkUpdateCallback.OnAfterUpdate(stage, gonglink, frontGongLink)
	}
}

func (gonglink *GongLink) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongLinkDeleteCallback != nil {
		var frontGongLink *GongLink
		if front != nil {
			frontGongLink, _ = front.(*GongLink)
		}
		stage.OnAfterGongLinkDeleteCallback.OnAfterDelete(stage, gonglink, frontGongLink)
	}
}

func (gongnote *GongNote) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongNoteCreateCallback != nil {
		stage.OnAfterGongNoteCreateCallback.OnAfterCreate(stage, gongnote)
	}
}

func (gongnote *GongNote) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongNoteUpdateCallback != nil {
		var frontGongNote *GongNote
		if front != nil {
			frontGongNote, _ = front.(*GongNote)
		}
		stage.OnAfterGongNoteUpdateCallback.OnAfterUpdate(stage, gongnote, frontGongNote)
	}
}

func (gongnote *GongNote) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongNoteDeleteCallback != nil {
		var frontGongNote *GongNote
		if front != nil {
			frontGongNote, _ = front.(*GongNote)
		}
		stage.OnAfterGongNoteDeleteCallback.OnAfterDelete(stage, gongnote, frontGongNote)
	}
}

func (gongstruct *GongStruct) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongStructCreateCallback != nil {
		stage.OnAfterGongStructCreateCallback.OnAfterCreate(stage, gongstruct)
	}
}

func (gongstruct *GongStruct) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongStructUpdateCallback != nil {
		var frontGongStruct *GongStruct
		if front != nil {
			frontGongStruct, _ = front.(*GongStruct)
		}
		stage.OnAfterGongStructUpdateCallback.OnAfterUpdate(stage, gongstruct, frontGongStruct)
	}
}

func (gongstruct *GongStruct) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongStructDeleteCallback != nil {
		var frontGongStruct *GongStruct
		if front != nil {
			frontGongStruct, _ = front.(*GongStruct)
		}
		stage.OnAfterGongStructDeleteCallback.OnAfterDelete(stage, gongstruct, frontGongStruct)
	}
}

func (gongtimefield *GongTimeField) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongTimeFieldCreateCallback != nil {
		stage.OnAfterGongTimeFieldCreateCallback.OnAfterCreate(stage, gongtimefield)
	}
}

func (gongtimefield *GongTimeField) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongTimeFieldUpdateCallback != nil {
		var frontGongTimeField *GongTimeField
		if front != nil {
			frontGongTimeField, _ = front.(*GongTimeField)
		}
		stage.OnAfterGongTimeFieldUpdateCallback.OnAfterUpdate(stage, gongtimefield, frontGongTimeField)
	}
}

func (gongtimefield *GongTimeField) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongTimeFieldDeleteCallback != nil {
		var frontGongTimeField *GongTimeField
		if front != nil {
			frontGongTimeField, _ = front.(*GongTimeField)
		}
		stage.OnAfterGongTimeFieldDeleteCallback.OnAfterDelete(stage, gongtimefield, frontGongTimeField)
	}
}

func (metareference *MetaReference) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetaReferenceCreateCallback != nil {
		stage.OnAfterMetaReferenceCreateCallback.OnAfterCreate(stage, metareference)
	}
}

func (metareference *MetaReference) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetaReferenceUpdateCallback != nil {
		var frontMetaReference *MetaReference
		if front != nil {
			frontMetaReference, _ = front.(*MetaReference)
		}
		stage.OnAfterMetaReferenceUpdateCallback.OnAfterUpdate(stage, metareference, frontMetaReference)
	}
}

func (metareference *MetaReference) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetaReferenceDeleteCallback != nil {
		var frontMetaReference *MetaReference
		if front != nil {
			frontMetaReference, _ = front.(*MetaReference)
		}
		stage.OnAfterMetaReferenceDeleteCallback.OnAfterDelete(stage, metareference, frontMetaReference)
	}
}

func (modelpkg *ModelPkg) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterModelPkgCreateCallback != nil {
		stage.OnAfterModelPkgCreateCallback.OnAfterCreate(stage, modelpkg)
	}
}

func (modelpkg *ModelPkg) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterModelPkgUpdateCallback != nil {
		var frontModelPkg *ModelPkg
		if front != nil {
			frontModelPkg, _ = front.(*ModelPkg)
		}
		stage.OnAfterModelPkgUpdateCallback.OnAfterUpdate(stage, modelpkg, frontModelPkg)
	}
}

func (modelpkg *ModelPkg) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterModelPkgDeleteCallback != nil {
		var frontModelPkg *ModelPkg
		if front != nil {
			frontModelPkg, _ = front.(*ModelPkg)
		}
		stage.OnAfterModelPkgDeleteCallback.OnAfterDelete(stage, modelpkg, frontModelPkg)
	}
}

func (pointertogongstructfield *PointerToGongStructField) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointerToGongStructFieldCreateCallback != nil {
		stage.OnAfterPointerToGongStructFieldCreateCallback.OnAfterCreate(stage, pointertogongstructfield)
	}
}

func (pointertogongstructfield *PointerToGongStructField) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointerToGongStructFieldUpdateCallback != nil {
		var frontPointerToGongStructField *PointerToGongStructField
		if front != nil {
			frontPointerToGongStructField, _ = front.(*PointerToGongStructField)
		}
		stage.OnAfterPointerToGongStructFieldUpdateCallback.OnAfterUpdate(stage, pointertogongstructfield, frontPointerToGongStructField)
	}
}

func (pointertogongstructfield *PointerToGongStructField) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointerToGongStructFieldDeleteCallback != nil {
		var frontPointerToGongStructField *PointerToGongStructField
		if front != nil {
			frontPointerToGongStructField, _ = front.(*PointerToGongStructField)
		}
		stage.OnAfterPointerToGongStructFieldDeleteCallback.OnAfterDelete(stage, pointertogongstructfield, frontPointerToGongStructField)
	}
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSliceOfPointerToGongStructFieldCreateCallback != nil {
		stage.OnAfterSliceOfPointerToGongStructFieldCreateCallback.OnAfterCreate(stage, sliceofpointertogongstructfield)
	}
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSliceOfPointerToGongStructFieldUpdateCallback != nil {
		var frontSliceOfPointerToGongStructField *SliceOfPointerToGongStructField
		if front != nil {
			frontSliceOfPointerToGongStructField, _ = front.(*SliceOfPointerToGongStructField)
		}
		stage.OnAfterSliceOfPointerToGongStructFieldUpdateCallback.OnAfterUpdate(stage, sliceofpointertogongstructfield, frontSliceOfPointerToGongStructField)
	}
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSliceOfPointerToGongStructFieldDeleteCallback != nil {
		var frontSliceOfPointerToGongStructField *SliceOfPointerToGongStructField
		if front != nil {
			frontSliceOfPointerToGongStructField, _ = front.(*SliceOfPointerToGongStructField)
		}
		stage.OnAfterSliceOfPointerToGongStructFieldDeleteCallback.OnAfterDelete(stage, sliceofpointertogongstructfield, frontSliceOfPointerToGongStructField)
	}
}

