// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Classdiagram:
		ok = stage.IsStagedClassdiagram(target)

	case *DiagramPackage:
		ok = stage.IsStagedDiagramPackage(target)

	case *AttributeShape:
		ok = stage.IsStagedFieldShape(target)

	case *GongEnumShape:
		ok = stage.IsStagedGongEnumShape(target)

	case *GongEnumValueEntry:
		ok = stage.IsStagedGongEnumValueEntry(target)

	case *GongStructShape:
		ok = stage.IsStagedGongStructShape(target)

	case *LinkShape:
		ok = stage.IsStagedLinkShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteShapeLink:
		ok = stage.IsStagedNoteShapeLink(target)

	case *Position:
		ok = stage.IsStagedPosition(target)

	case *UmlState:
		ok = stage.IsStagedUmlState(target)

	case *Umlsc:
		ok = stage.IsStagedUmlsc(target)

	case *Vertice:
		ok = stage.IsStagedVertice(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedClassdiagram(classdiagram *Classdiagram) (ok bool) {

	_, ok = stage.Classdiagrams[classdiagram]

	return
}

func (stage *Stage) IsStagedDiagramPackage(diagrampackage *DiagramPackage) (ok bool) {

	_, ok = stage.DiagramPackages[diagrampackage]

	return
}

func (stage *Stage) IsStagedFieldShape(fieldshape *AttributeShape) (ok bool) {

	_, ok = stage.FieldShapes[fieldshape]

	return
}

func (stage *Stage) IsStagedGongEnumShape(gongenumshape *GongEnumShape) (ok bool) {

	_, ok = stage.GongEnumShapes[gongenumshape]

	return
}

func (stage *Stage) IsStagedGongEnumValueEntry(gongenumvalueentry *GongEnumValueEntry) (ok bool) {

	_, ok = stage.GongEnumValueEntrys[gongenumvalueentry]

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

func (stage *Stage) IsStagedPosition(position *Position) (ok bool) {

	_, ok = stage.Positions[position]

	return
}

func (stage *Stage) IsStagedUmlState(umlstate *UmlState) (ok bool) {

	_, ok = stage.UmlStates[umlstate]

	return
}

func (stage *Stage) IsStagedUmlsc(umlsc *Umlsc) (ok bool) {

	_, ok = stage.Umlscs[umlsc]

	return
}

func (stage *Stage) IsStagedVertice(vertice *Vertice) (ok bool) {

	_, ok = stage.Vertices[vertice]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Classdiagram:
		stage.StageBranchClassdiagram(target)

	case *DiagramPackage:
		stage.StageBranchDiagramPackage(target)

	case *AttributeShape:
		stage.StageBranchFieldShape(target)

	case *GongEnumShape:
		stage.StageBranchGongEnumShape(target)

	case *GongEnumValueEntry:
		stage.StageBranchGongEnumValueEntry(target)

	case *GongStructShape:
		stage.StageBranchGongStructShape(target)

	case *LinkShape:
		stage.StageBranchLinkShape(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *NoteShapeLink:
		stage.StageBranchNoteShapeLink(target)

	case *Position:
		stage.StageBranchPosition(target)

	case *UmlState:
		stage.StageBranchUmlState(target)

	case *Umlsc:
		stage.StageBranchUmlsc(target)

	case *Vertice:
		stage.StageBranchVertice(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
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
	for _, _umlsc := range diagrampackage.Umlscs {
		StageBranch(stage, _umlsc)
	}

}

func (stage *Stage) StageBranchFieldShape(fieldshape *AttributeShape) {

	// check if instance is already staged
	if IsStaged(stage, fieldshape) {
		return
	}

	fieldshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGongEnumShape(gongenumshape *GongEnumShape) {

	// check if instance is already staged
	if IsStaged(stage, gongenumshape) {
		return
	}

	gongenumshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongenumshape.Position != nil {
		StageBranch(stage, gongenumshape.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueentry := range gongenumshape.GongEnumValueEntrys {
		StageBranch(stage, _gongenumvalueentry)
	}

}

func (stage *Stage) StageBranchGongEnumValueEntry(gongenumvalueentry *GongEnumValueEntry) {

	// check if instance is already staged
	if IsStaged(stage, gongenumvalueentry) {
		return
	}

	gongenumvalueentry.Stage(stage)

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
	if gongstructshape.Position != nil {
		StageBranch(stage, gongstructshape.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _fieldshape := range gongstructshape.AttributeShapes {
		StageBranch(stage, _fieldshape)
	}
	for _, _linkshape := range gongstructshape.Links {
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
	if linkshape.Middlevertice != nil {
		StageBranch(stage, linkshape.Middlevertice)
	}

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

func (stage *Stage) StageBranchPosition(position *Position) {

	// check if instance is already staged
	if IsStaged(stage, position) {
		return
	}

	position.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchUmlState(umlstate *UmlState) {

	// check if instance is already staged
	if IsStaged(stage, umlstate) {
		return
	}

	umlstate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchUmlsc(umlsc *Umlsc) {

	// check if instance is already staged
	if IsStaged(stage, umlsc) {
		return
	}

	umlsc.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _umlstate := range umlsc.States {
		StageBranch(stage, _umlstate)
	}

}

func (stage *Stage) StageBranchVertice(vertice *Vertice) {

	// check if instance is already staged
	if IsStaged(stage, vertice) {
		return
	}

	vertice.Stage(stage)

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
	case *Classdiagram:
		toT := CopyBranchClassdiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramPackage:
		toT := CopyBranchDiagramPackage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AttributeShape:
		toT := CopyBranchFieldShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnumShape:
		toT := CopyBranchGongEnumShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnumValueEntry:
		toT := CopyBranchGongEnumValueEntry(mapOrigCopy, fromT)
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

	case *Position:
		toT := CopyBranchPosition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *UmlState:
		toT := CopyBranchUmlState(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Umlsc:
		toT := CopyBranchUmlsc(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Vertice:
		toT := CopyBranchVertice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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
	for _, _umlsc := range diagrampackageFrom.Umlscs {
		diagrampackageTo.Umlscs = append(diagrampackageTo.Umlscs, CopyBranchUmlsc(mapOrigCopy, _umlsc))
	}

	return
}

func CopyBranchFieldShape(mapOrigCopy map[any]any, fieldshapeFrom *AttributeShape) (fieldshapeTo *AttributeShape) {

	// fieldshapeFrom has already been copied
	if _fieldshapeTo, ok := mapOrigCopy[fieldshapeFrom]; ok {
		fieldshapeTo = _fieldshapeTo.(*AttributeShape)
		return
	}

	fieldshapeTo = new(AttributeShape)
	mapOrigCopy[fieldshapeFrom] = fieldshapeTo
	fieldshapeFrom.CopyBasicFields(fieldshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
	if gongenumshapeFrom.Position != nil {
		gongenumshapeTo.Position = CopyBranchPosition(mapOrigCopy, gongenumshapeFrom.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueentry := range gongenumshapeFrom.GongEnumValueEntrys {
		gongenumshapeTo.GongEnumValueEntrys = append(gongenumshapeTo.GongEnumValueEntrys, CopyBranchGongEnumValueEntry(mapOrigCopy, _gongenumvalueentry))
	}

	return
}

func CopyBranchGongEnumValueEntry(mapOrigCopy map[any]any, gongenumvalueentryFrom *GongEnumValueEntry) (gongenumvalueentryTo *GongEnumValueEntry) {

	// gongenumvalueentryFrom has already been copied
	if _gongenumvalueentryTo, ok := mapOrigCopy[gongenumvalueentryFrom]; ok {
		gongenumvalueentryTo = _gongenumvalueentryTo.(*GongEnumValueEntry)
		return
	}

	gongenumvalueentryTo = new(GongEnumValueEntry)
	mapOrigCopy[gongenumvalueentryFrom] = gongenumvalueentryTo
	gongenumvalueentryFrom.CopyBasicFields(gongenumvalueentryTo)

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
	if gongstructshapeFrom.Position != nil {
		gongstructshapeTo.Position = CopyBranchPosition(mapOrigCopy, gongstructshapeFrom.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _fieldshape := range gongstructshapeFrom.AttributeShapes {
		gongstructshapeTo.AttributeShapes = append(gongstructshapeTo.AttributeShapes, CopyBranchFieldShape(mapOrigCopy, _fieldshape))
	}
	for _, _linkshape := range gongstructshapeFrom.Links {
		gongstructshapeTo.Links = append(gongstructshapeTo.Links, CopyBranchLinkShape(mapOrigCopy, _linkshape))
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
	if linkshapeFrom.Middlevertice != nil {
		linkshapeTo.Middlevertice = CopyBranchVertice(mapOrigCopy, linkshapeFrom.Middlevertice)
	}

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

func CopyBranchPosition(mapOrigCopy map[any]any, positionFrom *Position) (positionTo *Position) {

	// positionFrom has already been copied
	if _positionTo, ok := mapOrigCopy[positionFrom]; ok {
		positionTo = _positionTo.(*Position)
		return
	}

	positionTo = new(Position)
	mapOrigCopy[positionFrom] = positionTo
	positionFrom.CopyBasicFields(positionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchUmlState(mapOrigCopy map[any]any, umlstateFrom *UmlState) (umlstateTo *UmlState) {

	// umlstateFrom has already been copied
	if _umlstateTo, ok := mapOrigCopy[umlstateFrom]; ok {
		umlstateTo = _umlstateTo.(*UmlState)
		return
	}

	umlstateTo = new(UmlState)
	mapOrigCopy[umlstateFrom] = umlstateTo
	umlstateFrom.CopyBasicFields(umlstateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchUmlsc(mapOrigCopy map[any]any, umlscFrom *Umlsc) (umlscTo *Umlsc) {

	// umlscFrom has already been copied
	if _umlscTo, ok := mapOrigCopy[umlscFrom]; ok {
		umlscTo = _umlscTo.(*Umlsc)
		return
	}

	umlscTo = new(Umlsc)
	mapOrigCopy[umlscFrom] = umlscTo
	umlscFrom.CopyBasicFields(umlscTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _umlstate := range umlscFrom.States {
		umlscTo.States = append(umlscTo.States, CopyBranchUmlState(mapOrigCopy, _umlstate))
	}

	return
}

func CopyBranchVertice(mapOrigCopy map[any]any, verticeFrom *Vertice) (verticeTo *Vertice) {

	// verticeFrom has already been copied
	if _verticeTo, ok := mapOrigCopy[verticeFrom]; ok {
		verticeTo = _verticeTo.(*Vertice)
		return
	}

	verticeTo = new(Vertice)
	mapOrigCopy[verticeFrom] = verticeTo
	verticeFrom.CopyBasicFields(verticeTo)

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
	case *Classdiagram:
		stage.UnstageBranchClassdiagram(target)

	case *DiagramPackage:
		stage.UnstageBranchDiagramPackage(target)

	case *AttributeShape:
		stage.UnstageBranchFieldShape(target)

	case *GongEnumShape:
		stage.UnstageBranchGongEnumShape(target)

	case *GongEnumValueEntry:
		stage.UnstageBranchGongEnumValueEntry(target)

	case *GongStructShape:
		stage.UnstageBranchGongStructShape(target)

	case *LinkShape:
		stage.UnstageBranchLinkShape(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *NoteShapeLink:
		stage.UnstageBranchNoteShapeLink(target)

	case *Position:
		stage.UnstageBranchPosition(target)

	case *UmlState:
		stage.UnstageBranchUmlState(target)

	case *Umlsc:
		stage.UnstageBranchUmlsc(target)

	case *Vertice:
		stage.UnstageBranchVertice(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
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
	for _, _umlsc := range diagrampackage.Umlscs {
		UnstageBranch(stage, _umlsc)
	}

}

func (stage *Stage) UnstageBranchFieldShape(fieldshape *AttributeShape) {

	// check if instance is already staged
	if !IsStaged(stage, fieldshape) {
		return
	}

	fieldshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGongEnumShape(gongenumshape *GongEnumShape) {

	// check if instance is already staged
	if !IsStaged(stage, gongenumshape) {
		return
	}

	gongenumshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongenumshape.Position != nil {
		UnstageBranch(stage, gongenumshape.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueentry := range gongenumshape.GongEnumValueEntrys {
		UnstageBranch(stage, _gongenumvalueentry)
	}

}

func (stage *Stage) UnstageBranchGongEnumValueEntry(gongenumvalueentry *GongEnumValueEntry) {

	// check if instance is already staged
	if !IsStaged(stage, gongenumvalueentry) {
		return
	}

	gongenumvalueentry.Unstage(stage)

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
	if gongstructshape.Position != nil {
		UnstageBranch(stage, gongstructshape.Position)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _fieldshape := range gongstructshape.AttributeShapes {
		UnstageBranch(stage, _fieldshape)
	}
	for _, _linkshape := range gongstructshape.Links {
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
	if linkshape.Middlevertice != nil {
		UnstageBranch(stage, linkshape.Middlevertice)
	}

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

func (stage *Stage) UnstageBranchPosition(position *Position) {

	// check if instance is already staged
	if !IsStaged(stage, position) {
		return
	}

	position.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchUmlState(umlstate *UmlState) {

	// check if instance is already staged
	if !IsStaged(stage, umlstate) {
		return
	}

	umlstate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchUmlsc(umlsc *Umlsc) {

	// check if instance is already staged
	if !IsStaged(stage, umlsc) {
		return
	}

	umlsc.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _umlstate := range umlsc.States {
		UnstageBranch(stage, _umlstate)
	}

}

func (stage *Stage) UnstageBranchVertice(vertice *Vertice) {

	// check if instance is already staged
	if !IsStaged(stage, vertice) {
		return
	}

	vertice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
