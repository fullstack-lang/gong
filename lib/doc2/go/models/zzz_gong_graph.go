// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AttributeShape:
		ok = stage.IsStagedAttributeShape(target)

	case *Classdiagram:
		ok = stage.IsStagedClassdiagram(target)

	case *DiagramPackage:
		ok = stage.IsStagedDiagramPackage(target)

	case *GongEnumShape:
		ok = stage.IsStagedGongEnumShape(target)

	case *GongEnumValueShape:
		ok = stage.IsStagedGongEnumValueShape(target)

	case *GongStructShape:
		ok = stage.IsStagedGongStructShape(target)

	case *LinkShape:
		ok = stage.IsStagedLinkShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteShapeLink:
		ok = stage.IsStagedNoteShapeLink(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAttributeShape(attributeshape *AttributeShape) (ok bool) {

	_, ok = stage.AttributeShapes[attributeshape]

	return
}

func (stage *Stage) IsStagedClassdiagram(classdiagram *Classdiagram) (ok bool) {

	_, ok = stage.Classdiagrams[classdiagram]

	return
}

func (stage *Stage) IsStagedDiagramPackage(diagrampackage *DiagramPackage) (ok bool) {

	_, ok = stage.DiagramPackages[diagrampackage]

	return
}

func (stage *Stage) IsStagedGongEnumShape(gongenumshape *GongEnumShape) (ok bool) {

	_, ok = stage.GongEnumShapes[gongenumshape]

	return
}

func (stage *Stage) IsStagedGongEnumValueShape(gongenumvalueshape *GongEnumValueShape) (ok bool) {

	_, ok = stage.GongEnumValueShapes[gongenumvalueshape]

	return
}

func (stage *Stage) IsStagedGongStructShape(gongstructshape *GongStructShape) (ok bool) {

	_, ok = stage.GongStructShapes[gongstructshape]

	return
}

func (stage *Stage) IsStagedLinkShape(linkshape *LinkShape) (ok bool) {

	_, ok = stage.LinkShapes[linkshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteShapeLink(noteshapelink *NoteShapeLink) (ok bool) {

	_, ok = stage.NoteShapeLinks[noteshapelink]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AttributeShape:
		stage.StageBranchAttributeShape(target)

	case *Classdiagram:
		stage.StageBranchClassdiagram(target)

	case *DiagramPackage:
		stage.StageBranchDiagramPackage(target)

	case *GongEnumShape:
		stage.StageBranchGongEnumShape(target)

	case *GongEnumValueShape:
		stage.StageBranchGongEnumValueShape(target)

	case *GongStructShape:
		stage.StageBranchGongStructShape(target)

	case *LinkShape:
		stage.StageBranchLinkShape(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *NoteShapeLink:
		stage.StageBranchNoteShapeLink(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAttributeShape(attributeshape *AttributeShape) {

	// check if instance is already staged
	if IsStaged(stage, attributeshape) {
		return
	}

	attributeshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchClassdiagram(classdiagram *Classdiagram) {

	// check if instance is already staged
	if IsStaged(stage, classdiagram) {
		return
	}

	classdiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongstructshape := range classdiagram.GongStructShapes {
		StageBranch(stage, _gongstructshape)
	}
	for _, _gongenumshape := range classdiagram.GongEnumShapes {
		StageBranch(stage, _gongenumshape)
	}
	for _, _noteshape := range classdiagram.NoteShapes {
		StageBranch(stage, _noteshape)
	}

}

func (stage *Stage) StageBranchDiagramPackage(diagrampackage *DiagramPackage) {

	// check if instance is already staged
	if IsStaged(stage, diagrampackage) {
		return
	}

	diagrampackage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagrampackage.SelectedClassdiagram != nil {
		StageBranch(stage, diagrampackage.SelectedClassdiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _classdiagram := range diagrampackage.Classdiagrams {
		StageBranch(stage, _classdiagram)
	}

}

func (stage *Stage) StageBranchGongEnumShape(gongenumshape *GongEnumShape) {

	// check if instance is already staged
	if IsStaged(stage, gongenumshape) {
		return
	}

	gongenumshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
		StageBranch(stage, _gongenumvalueshape)
	}

}

func (stage *Stage) StageBranchGongEnumValueShape(gongenumvalueshape *GongEnumValueShape) {

	// check if instance is already staged
	if IsStaged(stage, gongenumvalueshape) {
		return
	}

	gongenumvalueshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGongStructShape(gongstructshape *GongStructShape) {

	// check if instance is already staged
	if IsStaged(stage, gongstructshape) {
		return
	}

	gongstructshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributeshape := range gongstructshape.AttributeShapes {
		StageBranch(stage, _attributeshape)
	}
	for _, _linkshape := range gongstructshape.LinkShapes {
		StageBranch(stage, _linkshape)
	}

}

func (stage *Stage) StageBranchLinkShape(linkshape *LinkShape) {

	// check if instance is already staged
	if IsStaged(stage, linkshape) {
		return
	}

	linkshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if IsStaged(stage, noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshapelink := range noteshape.NoteShapeLinks {
		StageBranch(stage, _noteshapelink)
	}

}

func (stage *Stage) StageBranchNoteShapeLink(noteshapelink *NoteShapeLink) {

	// check if instance is already staged
	if IsStaged(stage, noteshapelink) {
		return
	}

	noteshapelink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AttributeShape:
		toT := CopyBranchAttributeShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Classdiagram:
		toT := CopyBranchClassdiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramPackage:
		toT := CopyBranchDiagramPackage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnumShape:
		toT := CopyBranchGongEnumShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnumValueShape:
		toT := CopyBranchGongEnumValueShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongStructShape:
		toT := CopyBranchGongStructShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinkShape:
		toT := CopyBranchLinkShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := CopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShapeLink:
		toT := CopyBranchNoteShapeLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAttributeShape(mapOrigCopy map[any]any, attributeshapeFrom *AttributeShape) (attributeshapeTo *AttributeShape) {

	// attributeshapeFrom has already been copied
	if _attributeshapeTo, ok := mapOrigCopy[attributeshapeFrom]; ok {
		attributeshapeTo = _attributeshapeTo.(*AttributeShape)
		return
	}

	attributeshapeTo = new(AttributeShape)
	mapOrigCopy[attributeshapeFrom] = attributeshapeTo
	attributeshapeFrom.CopyBasicFields(attributeshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchClassdiagram(mapOrigCopy map[any]any, classdiagramFrom *Classdiagram) (classdiagramTo *Classdiagram) {

	// classdiagramFrom has already been copied
	if _classdiagramTo, ok := mapOrigCopy[classdiagramFrom]; ok {
		classdiagramTo = _classdiagramTo.(*Classdiagram)
		return
	}

	classdiagramTo = new(Classdiagram)
	mapOrigCopy[classdiagramFrom] = classdiagramTo
	classdiagramFrom.CopyBasicFields(classdiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongstructshape := range classdiagramFrom.GongStructShapes {
		classdiagramTo.GongStructShapes = append(classdiagramTo.GongStructShapes, CopyBranchGongStructShape(mapOrigCopy, _gongstructshape))
	}
	for _, _gongenumshape := range classdiagramFrom.GongEnumShapes {
		classdiagramTo.GongEnumShapes = append(classdiagramTo.GongEnumShapes, CopyBranchGongEnumShape(mapOrigCopy, _gongenumshape))
	}
	for _, _noteshape := range classdiagramFrom.NoteShapes {
		classdiagramTo.NoteShapes = append(classdiagramTo.NoteShapes, CopyBranchNoteShape(mapOrigCopy, _noteshape))
	}

	return
}

func CopyBranchDiagramPackage(mapOrigCopy map[any]any, diagrampackageFrom *DiagramPackage) (diagrampackageTo *DiagramPackage) {

	// diagrampackageFrom has already been copied
	if _diagrampackageTo, ok := mapOrigCopy[diagrampackageFrom]; ok {
		diagrampackageTo = _diagrampackageTo.(*DiagramPackage)
		return
	}

	diagrampackageTo = new(DiagramPackage)
	mapOrigCopy[diagrampackageFrom] = diagrampackageTo
	diagrampackageFrom.CopyBasicFields(diagrampackageTo)

	//insertion point for the staging of instances referenced by pointers
	if diagrampackageFrom.SelectedClassdiagram != nil {
		diagrampackageTo.SelectedClassdiagram = CopyBranchClassdiagram(mapOrigCopy, diagrampackageFrom.SelectedClassdiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _classdiagram := range diagrampackageFrom.Classdiagrams {
		diagrampackageTo.Classdiagrams = append(diagrampackageTo.Classdiagrams, CopyBranchClassdiagram(mapOrigCopy, _classdiagram))
	}

	return
}

func CopyBranchGongEnumShape(mapOrigCopy map[any]any, gongenumshapeFrom *GongEnumShape) (gongenumshapeTo *GongEnumShape) {

	// gongenumshapeFrom has already been copied
	if _gongenumshapeTo, ok := mapOrigCopy[gongenumshapeFrom]; ok {
		gongenumshapeTo = _gongenumshapeTo.(*GongEnumShape)
		return
	}

	gongenumshapeTo = new(GongEnumShape)
	mapOrigCopy[gongenumshapeFrom] = gongenumshapeTo
	gongenumshapeFrom.CopyBasicFields(gongenumshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueshape := range gongenumshapeFrom.GongEnumValueShapes {
		gongenumshapeTo.GongEnumValueShapes = append(gongenumshapeTo.GongEnumValueShapes, CopyBranchGongEnumValueShape(mapOrigCopy, _gongenumvalueshape))
	}

	return
}

func CopyBranchGongEnumValueShape(mapOrigCopy map[any]any, gongenumvalueshapeFrom *GongEnumValueShape) (gongenumvalueshapeTo *GongEnumValueShape) {

	// gongenumvalueshapeFrom has already been copied
	if _gongenumvalueshapeTo, ok := mapOrigCopy[gongenumvalueshapeFrom]; ok {
		gongenumvalueshapeTo = _gongenumvalueshapeTo.(*GongEnumValueShape)
		return
	}

	gongenumvalueshapeTo = new(GongEnumValueShape)
	mapOrigCopy[gongenumvalueshapeFrom] = gongenumvalueshapeTo
	gongenumvalueshapeFrom.CopyBasicFields(gongenumvalueshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGongStructShape(mapOrigCopy map[any]any, gongstructshapeFrom *GongStructShape) (gongstructshapeTo *GongStructShape) {

	// gongstructshapeFrom has already been copied
	if _gongstructshapeTo, ok := mapOrigCopy[gongstructshapeFrom]; ok {
		gongstructshapeTo = _gongstructshapeTo.(*GongStructShape)
		return
	}

	gongstructshapeTo = new(GongStructShape)
	mapOrigCopy[gongstructshapeFrom] = gongstructshapeTo
	gongstructshapeFrom.CopyBasicFields(gongstructshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributeshape := range gongstructshapeFrom.AttributeShapes {
		gongstructshapeTo.AttributeShapes = append(gongstructshapeTo.AttributeShapes, CopyBranchAttributeShape(mapOrigCopy, _attributeshape))
	}
	for _, _linkshape := range gongstructshapeFrom.LinkShapes {
		gongstructshapeTo.LinkShapes = append(gongstructshapeTo.LinkShapes, CopyBranchLinkShape(mapOrigCopy, _linkshape))
	}

	return
}

func CopyBranchLinkShape(mapOrigCopy map[any]any, linkshapeFrom *LinkShape) (linkshapeTo *LinkShape) {

	// linkshapeFrom has already been copied
	if _linkshapeTo, ok := mapOrigCopy[linkshapeFrom]; ok {
		linkshapeTo = _linkshapeTo.(*LinkShape)
		return
	}

	linkshapeTo = new(LinkShape)
	mapOrigCopy[linkshapeFrom] = linkshapeTo
	linkshapeFrom.CopyBasicFields(linkshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.CopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshapelink := range noteshapeFrom.NoteShapeLinks {
		noteshapeTo.NoteShapeLinks = append(noteshapeTo.NoteShapeLinks, CopyBranchNoteShapeLink(mapOrigCopy, _noteshapelink))
	}

	return
}

func CopyBranchNoteShapeLink(mapOrigCopy map[any]any, noteshapelinkFrom *NoteShapeLink) (noteshapelinkTo *NoteShapeLink) {

	// noteshapelinkFrom has already been copied
	if _noteshapelinkTo, ok := mapOrigCopy[noteshapelinkFrom]; ok {
		noteshapelinkTo = _noteshapelinkTo.(*NoteShapeLink)
		return
	}

	noteshapelinkTo = new(NoteShapeLink)
	mapOrigCopy[noteshapelinkFrom] = noteshapelinkTo
	noteshapelinkFrom.CopyBasicFields(noteshapelinkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *AttributeShape:
		stage.UnstageBranchAttributeShape(target)

	case *Classdiagram:
		stage.UnstageBranchClassdiagram(target)

	case *DiagramPackage:
		stage.UnstageBranchDiagramPackage(target)

	case *GongEnumShape:
		stage.UnstageBranchGongEnumShape(target)

	case *GongEnumValueShape:
		stage.UnstageBranchGongEnumValueShape(target)

	case *GongStructShape:
		stage.UnstageBranchGongStructShape(target)

	case *LinkShape:
		stage.UnstageBranchLinkShape(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *NoteShapeLink:
		stage.UnstageBranchNoteShapeLink(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAttributeShape(attributeshape *AttributeShape) {

	// check if instance is already staged
	if !IsStaged(stage, attributeshape) {
		return
	}

	attributeshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchClassdiagram(classdiagram *Classdiagram) {

	// check if instance is already staged
	if !IsStaged(stage, classdiagram) {
		return
	}

	classdiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongstructshape := range classdiagram.GongStructShapes {
		UnstageBranch(stage, _gongstructshape)
	}
	for _, _gongenumshape := range classdiagram.GongEnumShapes {
		UnstageBranch(stage, _gongenumshape)
	}
	for _, _noteshape := range classdiagram.NoteShapes {
		UnstageBranch(stage, _noteshape)
	}

}

func (stage *Stage) UnstageBranchDiagramPackage(diagrampackage *DiagramPackage) {

	// check if instance is already staged
	if !IsStaged(stage, diagrampackage) {
		return
	}

	diagrampackage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagrampackage.SelectedClassdiagram != nil {
		UnstageBranch(stage, diagrampackage.SelectedClassdiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _classdiagram := range diagrampackage.Classdiagrams {
		UnstageBranch(stage, _classdiagram)
	}

}

func (stage *Stage) UnstageBranchGongEnumShape(gongenumshape *GongEnumShape) {

	// check if instance is already staged
	if !IsStaged(stage, gongenumshape) {
		return
	}

	gongenumshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
		UnstageBranch(stage, _gongenumvalueshape)
	}

}

func (stage *Stage) UnstageBranchGongEnumValueShape(gongenumvalueshape *GongEnumValueShape) {

	// check if instance is already staged
	if !IsStaged(stage, gongenumvalueshape) {
		return
	}

	gongenumvalueshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGongStructShape(gongstructshape *GongStructShape) {

	// check if instance is already staged
	if !IsStaged(stage, gongstructshape) {
		return
	}

	gongstructshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributeshape := range gongstructshape.AttributeShapes {
		UnstageBranch(stage, _attributeshape)
	}
	for _, _linkshape := range gongstructshape.LinkShapes {
		UnstageBranch(stage, _linkshape)
	}

}

func (stage *Stage) UnstageBranchLinkShape(linkshape *LinkShape) {

	// check if instance is already staged
	if !IsStaged(stage, linkshape) {
		return
	}

	linkshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshapelink := range noteshape.NoteShapeLinks {
		UnstageBranch(stage, _noteshapelink)
	}

}

func (stage *Stage) UnstageBranchNoteShapeLink(noteshapelink *NoteShapeLink) {

	// check if instance is already staged
	if !IsStaged(stage, noteshapelink) {
		return
	}

	noteshapelink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
