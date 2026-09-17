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
func (attributeshape *AttributeShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAttributeShapeCreateCallback != nil {
		stage.OnAfterAttributeShapeCreateCallback.OnAfterCreate(stage, attributeshape)
	}
}

func (attributeshape *AttributeShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributeShapeUpdateCallback != nil {
		var frontAttributeShape *AttributeShape
		if front != nil {
			frontAttributeShape, _ = front.(*AttributeShape)
		}
		stage.OnAfterAttributeShapeUpdateCallback.OnAfterUpdate(stage, attributeshape, frontAttributeShape)
	}
}

func (attributeshape *AttributeShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAttributeShapeDeleteCallback != nil {
		var frontAttributeShape *AttributeShape
		if front != nil {
			frontAttributeShape, _ = front.(*AttributeShape)
		}
		stage.OnAfterAttributeShapeDeleteCallback.OnAfterDelete(stage, attributeshape, frontAttributeShape)
	}
}

func (classdiagram *Classdiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterClassdiagramCreateCallback != nil {
		stage.OnAfterClassdiagramCreateCallback.OnAfterCreate(stage, classdiagram)
	}
}

func (classdiagram *Classdiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClassdiagramUpdateCallback != nil {
		var frontClassdiagram *Classdiagram
		if front != nil {
			frontClassdiagram, _ = front.(*Classdiagram)
		}
		stage.OnAfterClassdiagramUpdateCallback.OnAfterUpdate(stage, classdiagram, frontClassdiagram)
	}
}

func (classdiagram *Classdiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClassdiagramDeleteCallback != nil {
		var frontClassdiagram *Classdiagram
		if front != nil {
			frontClassdiagram, _ = front.(*Classdiagram)
		}
		stage.OnAfterClassdiagramDeleteCallback.OnAfterDelete(stage, classdiagram, frontClassdiagram)
	}
}

func (diagrampackage *DiagramPackage) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramPackageCreateCallback != nil {
		stage.OnAfterDiagramPackageCreateCallback.OnAfterCreate(stage, diagrampackage)
	}
}

func (diagrampackage *DiagramPackage) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramPackageUpdateCallback != nil {
		var frontDiagramPackage *DiagramPackage
		if front != nil {
			frontDiagramPackage, _ = front.(*DiagramPackage)
		}
		stage.OnAfterDiagramPackageUpdateCallback.OnAfterUpdate(stage, diagrampackage, frontDiagramPackage)
	}
}

func (diagrampackage *DiagramPackage) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramPackageDeleteCallback != nil {
		var frontDiagramPackage *DiagramPackage
		if front != nil {
			frontDiagramPackage, _ = front.(*DiagramPackage)
		}
		stage.OnAfterDiagramPackageDeleteCallback.OnAfterDelete(stage, diagrampackage, frontDiagramPackage)
	}
}

func (gongenumshape *GongEnumShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongEnumShapeCreateCallback != nil {
		stage.OnAfterGongEnumShapeCreateCallback.OnAfterCreate(stage, gongenumshape)
	}
}

func (gongenumshape *GongEnumShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongEnumShapeUpdateCallback != nil {
		var frontGongEnumShape *GongEnumShape
		if front != nil {
			frontGongEnumShape, _ = front.(*GongEnumShape)
		}
		stage.OnAfterGongEnumShapeUpdateCallback.OnAfterUpdate(stage, gongenumshape, frontGongEnumShape)
	}
}

func (gongenumshape *GongEnumShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongEnumShapeDeleteCallback != nil {
		var frontGongEnumShape *GongEnumShape
		if front != nil {
			frontGongEnumShape, _ = front.(*GongEnumShape)
		}
		stage.OnAfterGongEnumShapeDeleteCallback.OnAfterDelete(stage, gongenumshape, frontGongEnumShape)
	}
}

func (gongenumvalueshape *GongEnumValueShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongEnumValueShapeCreateCallback != nil {
		stage.OnAfterGongEnumValueShapeCreateCallback.OnAfterCreate(stage, gongenumvalueshape)
	}
}

func (gongenumvalueshape *GongEnumValueShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongEnumValueShapeUpdateCallback != nil {
		var frontGongEnumValueShape *GongEnumValueShape
		if front != nil {
			frontGongEnumValueShape, _ = front.(*GongEnumValueShape)
		}
		stage.OnAfterGongEnumValueShapeUpdateCallback.OnAfterUpdate(stage, gongenumvalueshape, frontGongEnumValueShape)
	}
}

func (gongenumvalueshape *GongEnumValueShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongEnumValueShapeDeleteCallback != nil {
		var frontGongEnumValueShape *GongEnumValueShape
		if front != nil {
			frontGongEnumValueShape, _ = front.(*GongEnumValueShape)
		}
		stage.OnAfterGongEnumValueShapeDeleteCallback.OnAfterDelete(stage, gongenumvalueshape, frontGongEnumValueShape)
	}
}

func (gongnotelinkshape *GongNoteLinkShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongNoteLinkShapeCreateCallback != nil {
		stage.OnAfterGongNoteLinkShapeCreateCallback.OnAfterCreate(stage, gongnotelinkshape)
	}
}

func (gongnotelinkshape *GongNoteLinkShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongNoteLinkShapeUpdateCallback != nil {
		var frontGongNoteLinkShape *GongNoteLinkShape
		if front != nil {
			frontGongNoteLinkShape, _ = front.(*GongNoteLinkShape)
		}
		stage.OnAfterGongNoteLinkShapeUpdateCallback.OnAfterUpdate(stage, gongnotelinkshape, frontGongNoteLinkShape)
	}
}

func (gongnotelinkshape *GongNoteLinkShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongNoteLinkShapeDeleteCallback != nil {
		var frontGongNoteLinkShape *GongNoteLinkShape
		if front != nil {
			frontGongNoteLinkShape, _ = front.(*GongNoteLinkShape)
		}
		stage.OnAfterGongNoteLinkShapeDeleteCallback.OnAfterDelete(stage, gongnotelinkshape, frontGongNoteLinkShape)
	}
}

func (gongnoteshape *GongNoteShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongNoteShapeCreateCallback != nil {
		stage.OnAfterGongNoteShapeCreateCallback.OnAfterCreate(stage, gongnoteshape)
	}
}

func (gongnoteshape *GongNoteShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongNoteShapeUpdateCallback != nil {
		var frontGongNoteShape *GongNoteShape
		if front != nil {
			frontGongNoteShape, _ = front.(*GongNoteShape)
		}
		stage.OnAfterGongNoteShapeUpdateCallback.OnAfterUpdate(stage, gongnoteshape, frontGongNoteShape)
	}
}

func (gongnoteshape *GongNoteShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongNoteShapeDeleteCallback != nil {
		var frontGongNoteShape *GongNoteShape
		if front != nil {
			frontGongNoteShape, _ = front.(*GongNoteShape)
		}
		stage.OnAfterGongNoteShapeDeleteCallback.OnAfterDelete(stage, gongnoteshape, frontGongNoteShape)
	}
}

func (gongstructshape *GongStructShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGongStructShapeCreateCallback != nil {
		stage.OnAfterGongStructShapeCreateCallback.OnAfterCreate(stage, gongstructshape)
	}
}

func (gongstructshape *GongStructShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongStructShapeUpdateCallback != nil {
		var frontGongStructShape *GongStructShape
		if front != nil {
			frontGongStructShape, _ = front.(*GongStructShape)
		}
		stage.OnAfterGongStructShapeUpdateCallback.OnAfterUpdate(stage, gongstructshape, frontGongStructShape)
	}
}

func (gongstructshape *GongStructShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGongStructShapeDeleteCallback != nil {
		var frontGongStructShape *GongStructShape
		if front != nil {
			frontGongStructShape, _ = front.(*GongStructShape)
		}
		stage.OnAfterGongStructShapeDeleteCallback.OnAfterDelete(stage, gongstructshape, frontGongStructShape)
	}
}

func (linkshape *LinkShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinkShapeCreateCallback != nil {
		stage.OnAfterLinkShapeCreateCallback.OnAfterCreate(stage, linkshape)
	}
}

func (linkshape *LinkShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkShapeUpdateCallback != nil {
		var frontLinkShape *LinkShape
		if front != nil {
			frontLinkShape, _ = front.(*LinkShape)
		}
		stage.OnAfterLinkShapeUpdateCallback.OnAfterUpdate(stage, linkshape, frontLinkShape)
	}
}

func (linkshape *LinkShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinkShapeDeleteCallback != nil {
		var frontLinkShape *LinkShape
		if front != nil {
			frontLinkShape, _ = front.(*LinkShape)
		}
		stage.OnAfterLinkShapeDeleteCallback.OnAfterDelete(stage, linkshape, frontLinkShape)
	}
}

