package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *GongBasicField:
		if stage.OnAfterGongBasicFieldCreateCallback != nil {
			stage.OnAfterGongBasicFieldCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongEnum:
		if stage.OnAfterGongEnumCreateCallback != nil {
			stage.OnAfterGongEnumCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongEnumValue:
		if stage.OnAfterGongEnumValueCreateCallback != nil {
			stage.OnAfterGongEnumValueCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongLink:
		if stage.OnAfterGongLinkCreateCallback != nil {
			stage.OnAfterGongLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongNote:
		if stage.OnAfterGongNoteCreateCallback != nil {
			stage.OnAfterGongNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongStruct:
		if stage.OnAfterGongStructCreateCallback != nil {
			stage.OnAfterGongStructCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongTimeField:
		if stage.OnAfterGongTimeFieldCreateCallback != nil {
			stage.OnAfterGongTimeFieldCreateCallback.OnAfterCreate(stage, target)
		}
	case *Meta:
		if stage.OnAfterMetaCreateCallback != nil {
			stage.OnAfterMetaCreateCallback.OnAfterCreate(stage, target)
		}
	case *MetaReference:
		if stage.OnAfterMetaReferenceCreateCallback != nil {
			stage.OnAfterMetaReferenceCreateCallback.OnAfterCreate(stage, target)
		}
	case *ModelPkg:
		if stage.OnAfterModelPkgCreateCallback != nil {
			stage.OnAfterModelPkgCreateCallback.OnAfterCreate(stage, target)
		}
	case *PointerToGongStructField:
		if stage.OnAfterPointerToGongStructFieldCreateCallback != nil {
			stage.OnAfterPointerToGongStructFieldCreateCallback.OnAfterCreate(stage, target)
		}
	case *SliceOfPointerToGongStructField:
		if stage.OnAfterSliceOfPointerToGongStructFieldCreateCallback != nil {
			stage.OnAfterSliceOfPointerToGongStructFieldCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *GongBasicField:
		newTarget := any(new).(*GongBasicField)
		if stage.OnAfterGongBasicFieldUpdateCallback != nil {
			stage.OnAfterGongBasicFieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongEnum:
		newTarget := any(new).(*GongEnum)
		if stage.OnAfterGongEnumUpdateCallback != nil {
			stage.OnAfterGongEnumUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongEnumValue:
		newTarget := any(new).(*GongEnumValue)
		if stage.OnAfterGongEnumValueUpdateCallback != nil {
			stage.OnAfterGongEnumValueUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongLink:
		newTarget := any(new).(*GongLink)
		if stage.OnAfterGongLinkUpdateCallback != nil {
			stage.OnAfterGongLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongNote:
		newTarget := any(new).(*GongNote)
		if stage.OnAfterGongNoteUpdateCallback != nil {
			stage.OnAfterGongNoteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongStruct:
		newTarget := any(new).(*GongStruct)
		if stage.OnAfterGongStructUpdateCallback != nil {
			stage.OnAfterGongStructUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongTimeField:
		newTarget := any(new).(*GongTimeField)
		if stage.OnAfterGongTimeFieldUpdateCallback != nil {
			stage.OnAfterGongTimeFieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Meta:
		newTarget := any(new).(*Meta)
		if stage.OnAfterMetaUpdateCallback != nil {
			stage.OnAfterMetaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MetaReference:
		newTarget := any(new).(*MetaReference)
		if stage.OnAfterMetaReferenceUpdateCallback != nil {
			stage.OnAfterMetaReferenceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ModelPkg:
		newTarget := any(new).(*ModelPkg)
		if stage.OnAfterModelPkgUpdateCallback != nil {
			stage.OnAfterModelPkgUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PointerToGongStructField:
		newTarget := any(new).(*PointerToGongStructField)
		if stage.OnAfterPointerToGongStructFieldUpdateCallback != nil {
			stage.OnAfterPointerToGongStructFieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SliceOfPointerToGongStructField:
		newTarget := any(new).(*SliceOfPointerToGongStructField)
		if stage.OnAfterSliceOfPointerToGongStructFieldUpdateCallback != nil {
			stage.OnAfterSliceOfPointerToGongStructFieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *GongBasicField:
		if stage.OnAfterGongBasicFieldDeleteCallback != nil {
			staged := any(staged).(*GongBasicField)
			stage.OnAfterGongBasicFieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongEnum:
		if stage.OnAfterGongEnumDeleteCallback != nil {
			staged := any(staged).(*GongEnum)
			stage.OnAfterGongEnumDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongEnumValue:
		if stage.OnAfterGongEnumValueDeleteCallback != nil {
			staged := any(staged).(*GongEnumValue)
			stage.OnAfterGongEnumValueDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongLink:
		if stage.OnAfterGongLinkDeleteCallback != nil {
			staged := any(staged).(*GongLink)
			stage.OnAfterGongLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongNote:
		if stage.OnAfterGongNoteDeleteCallback != nil {
			staged := any(staged).(*GongNote)
			stage.OnAfterGongNoteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongStruct:
		if stage.OnAfterGongStructDeleteCallback != nil {
			staged := any(staged).(*GongStruct)
			stage.OnAfterGongStructDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongTimeField:
		if stage.OnAfterGongTimeFieldDeleteCallback != nil {
			staged := any(staged).(*GongTimeField)
			stage.OnAfterGongTimeFieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Meta:
		if stage.OnAfterMetaDeleteCallback != nil {
			staged := any(staged).(*Meta)
			stage.OnAfterMetaDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MetaReference:
		if stage.OnAfterMetaReferenceDeleteCallback != nil {
			staged := any(staged).(*MetaReference)
			stage.OnAfterMetaReferenceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ModelPkg:
		if stage.OnAfterModelPkgDeleteCallback != nil {
			staged := any(staged).(*ModelPkg)
			stage.OnAfterModelPkgDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PointerToGongStructField:
		if stage.OnAfterPointerToGongStructFieldDeleteCallback != nil {
			staged := any(staged).(*PointerToGongStructField)
			stage.OnAfterPointerToGongStructFieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SliceOfPointerToGongStructField:
		if stage.OnAfterSliceOfPointerToGongStructFieldDeleteCallback != nil {
			staged := any(staged).(*SliceOfPointerToGongStructField)
			stage.OnAfterSliceOfPointerToGongStructFieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *GongBasicField:
		if stage.OnAfterGongBasicFieldReadCallback != nil {
			stage.OnAfterGongBasicFieldReadCallback.OnAfterRead(stage, target)
		}
	case *GongEnum:
		if stage.OnAfterGongEnumReadCallback != nil {
			stage.OnAfterGongEnumReadCallback.OnAfterRead(stage, target)
		}
	case *GongEnumValue:
		if stage.OnAfterGongEnumValueReadCallback != nil {
			stage.OnAfterGongEnumValueReadCallback.OnAfterRead(stage, target)
		}
	case *GongLink:
		if stage.OnAfterGongLinkReadCallback != nil {
			stage.OnAfterGongLinkReadCallback.OnAfterRead(stage, target)
		}
	case *GongNote:
		if stage.OnAfterGongNoteReadCallback != nil {
			stage.OnAfterGongNoteReadCallback.OnAfterRead(stage, target)
		}
	case *GongStruct:
		if stage.OnAfterGongStructReadCallback != nil {
			stage.OnAfterGongStructReadCallback.OnAfterRead(stage, target)
		}
	case *GongTimeField:
		if stage.OnAfterGongTimeFieldReadCallback != nil {
			stage.OnAfterGongTimeFieldReadCallback.OnAfterRead(stage, target)
		}
	case *Meta:
		if stage.OnAfterMetaReadCallback != nil {
			stage.OnAfterMetaReadCallback.OnAfterRead(stage, target)
		}
	case *MetaReference:
		if stage.OnAfterMetaReferenceReadCallback != nil {
			stage.OnAfterMetaReferenceReadCallback.OnAfterRead(stage, target)
		}
	case *ModelPkg:
		if stage.OnAfterModelPkgReadCallback != nil {
			stage.OnAfterModelPkgReadCallback.OnAfterRead(stage, target)
		}
	case *PointerToGongStructField:
		if stage.OnAfterPointerToGongStructFieldReadCallback != nil {
			stage.OnAfterPointerToGongStructFieldReadCallback.OnAfterRead(stage, target)
		}
	case *SliceOfPointerToGongStructField:
		if stage.OnAfterSliceOfPointerToGongStructFieldReadCallback != nil {
			stage.OnAfterSliceOfPointerToGongStructFieldReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *GongBasicField:
		stage.OnAfterGongBasicFieldUpdateCallback = any(callback).(OnAfterUpdateInterface[GongBasicField])
	
	case *GongEnum:
		stage.OnAfterGongEnumUpdateCallback = any(callback).(OnAfterUpdateInterface[GongEnum])
	
	case *GongEnumValue:
		stage.OnAfterGongEnumValueUpdateCallback = any(callback).(OnAfterUpdateInterface[GongEnumValue])
	
	case *GongLink:
		stage.OnAfterGongLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[GongLink])
	
	case *GongNote:
		stage.OnAfterGongNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[GongNote])
	
	case *GongStruct:
		stage.OnAfterGongStructUpdateCallback = any(callback).(OnAfterUpdateInterface[GongStruct])
	
	case *GongTimeField:
		stage.OnAfterGongTimeFieldUpdateCallback = any(callback).(OnAfterUpdateInterface[GongTimeField])
	
	case *Meta:
		stage.OnAfterMetaUpdateCallback = any(callback).(OnAfterUpdateInterface[Meta])
	
	case *MetaReference:
		stage.OnAfterMetaReferenceUpdateCallback = any(callback).(OnAfterUpdateInterface[MetaReference])
	
	case *ModelPkg:
		stage.OnAfterModelPkgUpdateCallback = any(callback).(OnAfterUpdateInterface[ModelPkg])
	
	case *PointerToGongStructField:
		stage.OnAfterPointerToGongStructFieldUpdateCallback = any(callback).(OnAfterUpdateInterface[PointerToGongStructField])
	
	case *SliceOfPointerToGongStructField:
		stage.OnAfterSliceOfPointerToGongStructFieldUpdateCallback = any(callback).(OnAfterUpdateInterface[SliceOfPointerToGongStructField])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *GongBasicField:
		stage.OnAfterGongBasicFieldCreateCallback = any(callback).(OnAfterCreateInterface[GongBasicField])
	
	case *GongEnum:
		stage.OnAfterGongEnumCreateCallback = any(callback).(OnAfterCreateInterface[GongEnum])
	
	case *GongEnumValue:
		stage.OnAfterGongEnumValueCreateCallback = any(callback).(OnAfterCreateInterface[GongEnumValue])
	
	case *GongLink:
		stage.OnAfterGongLinkCreateCallback = any(callback).(OnAfterCreateInterface[GongLink])
	
	case *GongNote:
		stage.OnAfterGongNoteCreateCallback = any(callback).(OnAfterCreateInterface[GongNote])
	
	case *GongStruct:
		stage.OnAfterGongStructCreateCallback = any(callback).(OnAfterCreateInterface[GongStruct])
	
	case *GongTimeField:
		stage.OnAfterGongTimeFieldCreateCallback = any(callback).(OnAfterCreateInterface[GongTimeField])
	
	case *Meta:
		stage.OnAfterMetaCreateCallback = any(callback).(OnAfterCreateInterface[Meta])
	
	case *MetaReference:
		stage.OnAfterMetaReferenceCreateCallback = any(callback).(OnAfterCreateInterface[MetaReference])
	
	case *ModelPkg:
		stage.OnAfterModelPkgCreateCallback = any(callback).(OnAfterCreateInterface[ModelPkg])
	
	case *PointerToGongStructField:
		stage.OnAfterPointerToGongStructFieldCreateCallback = any(callback).(OnAfterCreateInterface[PointerToGongStructField])
	
	case *SliceOfPointerToGongStructField:
		stage.OnAfterSliceOfPointerToGongStructFieldCreateCallback = any(callback).(OnAfterCreateInterface[SliceOfPointerToGongStructField])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *GongBasicField:
		stage.OnAfterGongBasicFieldDeleteCallback = any(callback).(OnAfterDeleteInterface[GongBasicField])
	
	case *GongEnum:
		stage.OnAfterGongEnumDeleteCallback = any(callback).(OnAfterDeleteInterface[GongEnum])
	
	case *GongEnumValue:
		stage.OnAfterGongEnumValueDeleteCallback = any(callback).(OnAfterDeleteInterface[GongEnumValue])
	
	case *GongLink:
		stage.OnAfterGongLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[GongLink])
	
	case *GongNote:
		stage.OnAfterGongNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[GongNote])
	
	case *GongStruct:
		stage.OnAfterGongStructDeleteCallback = any(callback).(OnAfterDeleteInterface[GongStruct])
	
	case *GongTimeField:
		stage.OnAfterGongTimeFieldDeleteCallback = any(callback).(OnAfterDeleteInterface[GongTimeField])
	
	case *Meta:
		stage.OnAfterMetaDeleteCallback = any(callback).(OnAfterDeleteInterface[Meta])
	
	case *MetaReference:
		stage.OnAfterMetaReferenceDeleteCallback = any(callback).(OnAfterDeleteInterface[MetaReference])
	
	case *ModelPkg:
		stage.OnAfterModelPkgDeleteCallback = any(callback).(OnAfterDeleteInterface[ModelPkg])
	
	case *PointerToGongStructField:
		stage.OnAfterPointerToGongStructFieldDeleteCallback = any(callback).(OnAfterDeleteInterface[PointerToGongStructField])
	
	case *SliceOfPointerToGongStructField:
		stage.OnAfterSliceOfPointerToGongStructFieldDeleteCallback = any(callback).(OnAfterDeleteInterface[SliceOfPointerToGongStructField])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *GongBasicField:
		stage.OnAfterGongBasicFieldReadCallback = any(callback).(OnAfterReadInterface[GongBasicField])
	
	case *GongEnum:
		stage.OnAfterGongEnumReadCallback = any(callback).(OnAfterReadInterface[GongEnum])
	
	case *GongEnumValue:
		stage.OnAfterGongEnumValueReadCallback = any(callback).(OnAfterReadInterface[GongEnumValue])
	
	case *GongLink:
		stage.OnAfterGongLinkReadCallback = any(callback).(OnAfterReadInterface[GongLink])
	
	case *GongNote:
		stage.OnAfterGongNoteReadCallback = any(callback).(OnAfterReadInterface[GongNote])
	
	case *GongStruct:
		stage.OnAfterGongStructReadCallback = any(callback).(OnAfterReadInterface[GongStruct])
	
	case *GongTimeField:
		stage.OnAfterGongTimeFieldReadCallback = any(callback).(OnAfterReadInterface[GongTimeField])
	
	case *Meta:
		stage.OnAfterMetaReadCallback = any(callback).(OnAfterReadInterface[Meta])
	
	case *MetaReference:
		stage.OnAfterMetaReferenceReadCallback = any(callback).(OnAfterReadInterface[MetaReference])
	
	case *ModelPkg:
		stage.OnAfterModelPkgReadCallback = any(callback).(OnAfterReadInterface[ModelPkg])
	
	case *PointerToGongStructField:
		stage.OnAfterPointerToGongStructFieldReadCallback = any(callback).(OnAfterReadInterface[PointerToGongStructField])
	
	case *SliceOfPointerToGongStructField:
		stage.OnAfterSliceOfPointerToGongStructFieldReadCallback = any(callback).(OnAfterReadInterface[SliceOfPointerToGongStructField])
	
	}
}
