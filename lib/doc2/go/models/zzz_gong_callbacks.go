// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

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

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type, mouseEvent *Gong__MouseEvent) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AttributeShape:
		newTarget := any(new).(*AttributeShape)
		if stage.OnAfterAttributeShapeUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterAttributeShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterAttributeShapeUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterAttributeShapeUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Classdiagram:
		newTarget := any(new).(*Classdiagram)
		if stage.OnAfterClassdiagramUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterClassdiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterClassdiagramUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterClassdiagramUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *DiagramPackage:
		newTarget := any(new).(*DiagramPackage)
		if stage.OnAfterDiagramPackageUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterDiagramPackageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterDiagramPackageUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterDiagramPackageUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *GongEnumShape:
		newTarget := any(new).(*GongEnumShape)
		if stage.OnAfterGongEnumShapeUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterGongEnumShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterGongEnumShapeUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterGongEnumShapeUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *GongEnumValueShape:
		newTarget := any(new).(*GongEnumValueShape)
		if stage.OnAfterGongEnumValueShapeUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterGongEnumValueShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterGongEnumValueShapeUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterGongEnumValueShapeUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *GongNoteLinkShape:
		newTarget := any(new).(*GongNoteLinkShape)
		if stage.OnAfterGongNoteLinkShapeUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterGongNoteLinkShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterGongNoteLinkShapeUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterGongNoteLinkShapeUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *GongNoteShape:
		newTarget := any(new).(*GongNoteShape)
		if stage.OnAfterGongNoteShapeUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterGongNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterGongNoteShapeUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterGongNoteShapeUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *GongStructShape:
		newTarget := any(new).(*GongStructShape)
		if stage.OnAfterGongStructShapeUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterGongStructShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterGongStructShapeUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterGongStructShapeUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *LinkShape:
		newTarget := any(new).(*LinkShape)
		if stage.OnAfterLinkShapeUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterLinkShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterLinkShapeUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterLinkShapeUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

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

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AttributeShape:
		if stage.OnAfterAttributeShapeReadCallback != nil {
			stage.OnAfterAttributeShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Classdiagram:
		if stage.OnAfterClassdiagramReadCallback != nil {
			stage.OnAfterClassdiagramReadCallback.OnAfterRead(stage, target)
		}
	case *DiagramPackage:
		if stage.OnAfterDiagramPackageReadCallback != nil {
			stage.OnAfterDiagramPackageReadCallback.OnAfterRead(stage, target)
		}
	case *GongEnumShape:
		if stage.OnAfterGongEnumShapeReadCallback != nil {
			stage.OnAfterGongEnumShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GongEnumValueShape:
		if stage.OnAfterGongEnumValueShapeReadCallback != nil {
			stage.OnAfterGongEnumValueShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GongNoteLinkShape:
		if stage.OnAfterGongNoteLinkShapeReadCallback != nil {
			stage.OnAfterGongNoteLinkShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GongNoteShape:
		if stage.OnAfterGongNoteShapeReadCallback != nil {
			stage.OnAfterGongNoteShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GongStructShape:
		if stage.OnAfterGongStructShapeReadCallback != nil {
			stage.OnAfterGongStructShapeReadCallback.OnAfterRead(stage, target)
		}
	case *LinkShape:
		if stage.OnAfterLinkShapeReadCallback != nil {
			stage.OnAfterLinkShapeReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AttributeShape:
		stage.OnAfterAttributeShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[AttributeShape])
	
	case *Classdiagram:
		stage.OnAfterClassdiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Classdiagram])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramPackage])
	
	case *GongEnumShape:
		stage.OnAfterGongEnumShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GongEnumShape])
	
	case *GongEnumValueShape:
		stage.OnAfterGongEnumValueShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GongEnumValueShape])
	
	case *GongNoteLinkShape:
		stage.OnAfterGongNoteLinkShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GongNoteLinkShape])
	
	case *GongNoteShape:
		stage.OnAfterGongNoteShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GongNoteShape])
	
	case *GongStructShape:
		stage.OnAfterGongStructShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GongStructShape])
	
	case *LinkShape:
		stage.OnAfterLinkShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[LinkShape])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AttributeShape:
		stage.OnAfterAttributeShapeCreateCallback = any(callback).(OnAfterCreateInterface[AttributeShape])
	
	case *Classdiagram:
		stage.OnAfterClassdiagramCreateCallback = any(callback).(OnAfterCreateInterface[Classdiagram])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageCreateCallback = any(callback).(OnAfterCreateInterface[DiagramPackage])
	
	case *GongEnumShape:
		stage.OnAfterGongEnumShapeCreateCallback = any(callback).(OnAfterCreateInterface[GongEnumShape])
	
	case *GongEnumValueShape:
		stage.OnAfterGongEnumValueShapeCreateCallback = any(callback).(OnAfterCreateInterface[GongEnumValueShape])
	
	case *GongNoteLinkShape:
		stage.OnAfterGongNoteLinkShapeCreateCallback = any(callback).(OnAfterCreateInterface[GongNoteLinkShape])
	
	case *GongNoteShape:
		stage.OnAfterGongNoteShapeCreateCallback = any(callback).(OnAfterCreateInterface[GongNoteShape])
	
	case *GongStructShape:
		stage.OnAfterGongStructShapeCreateCallback = any(callback).(OnAfterCreateInterface[GongStructShape])
	
	case *LinkShape:
		stage.OnAfterLinkShapeCreateCallback = any(callback).(OnAfterCreateInterface[LinkShape])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AttributeShape:
		stage.OnAfterAttributeShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[AttributeShape])
	
	case *Classdiagram:
		stage.OnAfterClassdiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Classdiagram])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramPackage])
	
	case *GongEnumShape:
		stage.OnAfterGongEnumShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GongEnumShape])
	
	case *GongEnumValueShape:
		stage.OnAfterGongEnumValueShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GongEnumValueShape])
	
	case *GongNoteLinkShape:
		stage.OnAfterGongNoteLinkShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GongNoteLinkShape])
	
	case *GongNoteShape:
		stage.OnAfterGongNoteShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GongNoteShape])
	
	case *GongStructShape:
		stage.OnAfterGongStructShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GongStructShape])
	
	case *LinkShape:
		stage.OnAfterLinkShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[LinkShape])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AttributeShape:
		stage.OnAfterAttributeShapeReadCallback = any(callback).(OnAfterReadInterface[AttributeShape])
	
	case *Classdiagram:
		stage.OnAfterClassdiagramReadCallback = any(callback).(OnAfterReadInterface[Classdiagram])
	
	case *DiagramPackage:
		stage.OnAfterDiagramPackageReadCallback = any(callback).(OnAfterReadInterface[DiagramPackage])
	
	case *GongEnumShape:
		stage.OnAfterGongEnumShapeReadCallback = any(callback).(OnAfterReadInterface[GongEnumShape])
	
	case *GongEnumValueShape:
		stage.OnAfterGongEnumValueShapeReadCallback = any(callback).(OnAfterReadInterface[GongEnumValueShape])
	
	case *GongNoteLinkShape:
		stage.OnAfterGongNoteLinkShapeReadCallback = any(callback).(OnAfterReadInterface[GongNoteLinkShape])
	
	case *GongNoteShape:
		stage.OnAfterGongNoteShapeReadCallback = any(callback).(OnAfterReadInterface[GongNoteShape])
	
	case *GongStructShape:
		stage.OnAfterGongStructShapeReadCallback = any(callback).(OnAfterReadInterface[GongStructShape])
	
	case *LinkShape:
		stage.OnAfterLinkShapeReadCallback = any(callback).(OnAfterReadInterface[LinkShape])
	
	}
}
