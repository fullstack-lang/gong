// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AttributeShape:
		if stage.OnAfterAttributeShapeCreateCallback != nil {
			stage.OnAfterAttributeShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Classdiagram:
		if stage.OnAfterClassdiagramCreateCallback != nil {
			stage.OnAfterClassdiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageCreateCallback != nil {
			stage.OnAfterDiagramPackageCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongEnumShape:
		if stage.OnAfterGongEnumShapeCreateCallback != nil {
			stage.OnAfterGongEnumShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongEnumValueShape:
		if stage.OnAfterGongEnumValueShapeCreateCallback != nil {
			stage.OnAfterGongEnumValueShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongNoteLinkShape:
		if stage.OnAfterGongNoteLinkShapeCreateCallback != nil {
			stage.OnAfterGongNoteLinkShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongNoteShape:
		if stage.OnAfterGongNoteShapeCreateCallback != nil {
			stage.OnAfterGongNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GongStructShape:
		if stage.OnAfterGongStructShapeCreateCallback != nil {
			stage.OnAfterGongStructShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *LinkShape:
		if stage.OnAfterLinkShapeCreateCallback != nil {
			stage.OnAfterLinkShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *AttributeShape:
		newTarget := any(new).(*AttributeShape)
		if stage.OnAfterAttributeShapeUpdateCallback != nil {
			stage.OnAfterAttributeShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Classdiagram:
		newTarget := any(new).(*Classdiagram)
		if stage.OnAfterClassdiagramUpdateCallback != nil {
			stage.OnAfterClassdiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramPackage:
		newTarget := any(new).(*DiagramPackage)
		if stage.OnAfterDiagramPackageUpdateCallback != nil {
			stage.OnAfterDiagramPackageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongEnumShape:
		newTarget := any(new).(*GongEnumShape)
		if stage.OnAfterGongEnumShapeUpdateCallback != nil {
			stage.OnAfterGongEnumShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongEnumValueShape:
		newTarget := any(new).(*GongEnumValueShape)
		if stage.OnAfterGongEnumValueShapeUpdateCallback != nil {
			stage.OnAfterGongEnumValueShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongNoteLinkShape:
		newTarget := any(new).(*GongNoteLinkShape)
		if stage.OnAfterGongNoteLinkShapeUpdateCallback != nil {
			stage.OnAfterGongNoteLinkShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongNoteShape:
		newTarget := any(new).(*GongNoteShape)
		if stage.OnAfterGongNoteShapeUpdateCallback != nil {
			stage.OnAfterGongNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GongStructShape:
		newTarget := any(new).(*GongStructShape)
		if stage.OnAfterGongStructShapeUpdateCallback != nil {
			stage.OnAfterGongStructShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LinkShape:
		newTarget := any(new).(*LinkShape)
		if stage.OnAfterLinkShapeUpdateCallback != nil {
			stage.OnAfterLinkShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *AttributeShape:
		if stage.OnAfterAttributeShapeDeleteCallback != nil {
			staged := any(staged).(*AttributeShape)
			stage.OnAfterAttributeShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Classdiagram:
		if stage.OnAfterClassdiagramDeleteCallback != nil {
			staged := any(staged).(*Classdiagram)
			stage.OnAfterClassdiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageDeleteCallback != nil {
			staged := any(staged).(*DiagramPackage)
			stage.OnAfterDiagramPackageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongEnumShape:
		if stage.OnAfterGongEnumShapeDeleteCallback != nil {
			staged := any(staged).(*GongEnumShape)
			stage.OnAfterGongEnumShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongEnumValueShape:
		if stage.OnAfterGongEnumValueShapeDeleteCallback != nil {
			staged := any(staged).(*GongEnumValueShape)
			stage.OnAfterGongEnumValueShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongNoteLinkShape:
		if stage.OnAfterGongNoteLinkShapeDeleteCallback != nil {
			staged := any(staged).(*GongNoteLinkShape)
			stage.OnAfterGongNoteLinkShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongNoteShape:
		if stage.OnAfterGongNoteShapeDeleteCallback != nil {
			staged := any(staged).(*GongNoteShape)
			stage.OnAfterGongNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GongStructShape:
		if stage.OnAfterGongStructShapeDeleteCallback != nil {
			staged := any(staged).(*GongStructShape)
			stage.OnAfterGongStructShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LinkShape:
		if stage.OnAfterLinkShapeDeleteCallback != nil {
			staged := any(staged).(*LinkShape)
			stage.OnAfterLinkShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
