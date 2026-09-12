// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
