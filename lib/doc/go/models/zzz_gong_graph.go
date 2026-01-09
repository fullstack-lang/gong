// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

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

	case *GongNoteLinkShape:
		ok = stage.IsStagedGongNoteLinkShape(target)

	case *GongNoteShape:
		ok = stage.IsStagedGongNoteShape(target)

	case *GongStructShape:
		ok = stage.IsStagedGongStructShape(target)

	case *LinkShape:
		ok = stage.IsStagedLinkShape(target)

	default:
		_ = target
	}
	return
}

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

	case *GongNoteLinkShape:
		ok = stage.IsStagedGongNoteLinkShape(target)

	case *GongNoteShape:
		ok = stage.IsStagedGongNoteShape(target)

	case *GongStructShape:
		ok = stage.IsStagedGongStructShape(target)

	case *LinkShape:
		ok = stage.IsStagedLinkShape(target)

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

func (stage *Stage) IsStagedGongNoteLinkShape(gongnotelinkshape *GongNoteLinkShape) (ok bool) {

	_, ok = stage.GongNoteLinkShapes[gongnotelinkshape]

	return
}

func (stage *Stage) IsStagedGongNoteShape(gongnoteshape *GongNoteShape) (ok bool) {

	_, ok = stage.GongNoteShapes[gongnoteshape]

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

	case *GongNoteLinkShape:
		stage.StageBranchGongNoteLinkShape(target)

	case *GongNoteShape:
		stage.StageBranchGongNoteShape(target)

	case *GongStructShape:
		stage.StageBranchGongStructShape(target)

	case *LinkShape:
		stage.StageBranchLinkShape(target)

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
	for _, _gongnoteshape := range classdiagram.GongNoteShapes {
		StageBranch(stage, _gongnoteshape)
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

func (stage *Stage) StageBranchGongNoteLinkShape(gongnotelinkshape *GongNoteLinkShape) {

	// check if instance is already staged
	if IsStaged(stage, gongnotelinkshape) {
		return
	}

	gongnotelinkshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGongNoteShape(gongnoteshape *GongNoteShape) {

	// check if instance is already staged
	if IsStaged(stage, gongnoteshape) {
		return
	}

	gongnoteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
		StageBranch(stage, _gongnotelinkshape)
	}

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

	case *GongNoteLinkShape:
		toT := CopyBranchGongNoteLinkShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongNoteShape:
		toT := CopyBranchGongNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongStructShape:
		toT := CopyBranchGongStructShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinkShape:
		toT := CopyBranchLinkShape(mapOrigCopy, fromT)
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
	for _, _gongnoteshape := range classdiagramFrom.GongNoteShapes {
		classdiagramTo.GongNoteShapes = append(classdiagramTo.GongNoteShapes, CopyBranchGongNoteShape(mapOrigCopy, _gongnoteshape))
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

func CopyBranchGongNoteLinkShape(mapOrigCopy map[any]any, gongnotelinkshapeFrom *GongNoteLinkShape) (gongnotelinkshapeTo *GongNoteLinkShape) {

	// gongnotelinkshapeFrom has already been copied
	if _gongnotelinkshapeTo, ok := mapOrigCopy[gongnotelinkshapeFrom]; ok {
		gongnotelinkshapeTo = _gongnotelinkshapeTo.(*GongNoteLinkShape)
		return
	}

	gongnotelinkshapeTo = new(GongNoteLinkShape)
	mapOrigCopy[gongnotelinkshapeFrom] = gongnotelinkshapeTo
	gongnotelinkshapeFrom.CopyBasicFields(gongnotelinkshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGongNoteShape(mapOrigCopy map[any]any, gongnoteshapeFrom *GongNoteShape) (gongnoteshapeTo *GongNoteShape) {

	// gongnoteshapeFrom has already been copied
	if _gongnoteshapeTo, ok := mapOrigCopy[gongnoteshapeFrom]; ok {
		gongnoteshapeTo = _gongnoteshapeTo.(*GongNoteShape)
		return
	}

	gongnoteshapeTo = new(GongNoteShape)
	mapOrigCopy[gongnoteshapeFrom] = gongnoteshapeTo
	gongnoteshapeFrom.CopyBasicFields(gongnoteshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongnotelinkshape := range gongnoteshapeFrom.GongNoteLinkShapes {
		gongnoteshapeTo.GongNoteLinkShapes = append(gongnoteshapeTo.GongNoteLinkShapes, CopyBranchGongNoteLinkShape(mapOrigCopy, _gongnotelinkshape))
	}

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

	case *GongNoteLinkShape:
		stage.UnstageBranchGongNoteLinkShape(target)

	case *GongNoteShape:
		stage.UnstageBranchGongNoteShape(target)

	case *GongStructShape:
		stage.UnstageBranchGongStructShape(target)

	case *LinkShape:
		stage.UnstageBranchLinkShape(target)

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
	for _, _gongnoteshape := range classdiagram.GongNoteShapes {
		UnstageBranch(stage, _gongnoteshape)
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

func (stage *Stage) UnstageBranchGongNoteLinkShape(gongnotelinkshape *GongNoteLinkShape) {

	// check if instance is already staged
	if !IsStaged(stage, gongnotelinkshape) {
		return
	}

	gongnotelinkshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGongNoteShape(gongnoteshape *GongNoteShape) {

	// check if instance is already staged
	if !IsStaged(stage, gongnoteshape) {
		return
	}

	gongnoteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
		UnstageBranch(stage, _gongnotelinkshape)
	}

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

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attributeshape *AttributeShape) GongDiff(stage *Stage, attributeshapeOther *AttributeShape) (diffs []string) {
	// insertion point for field diffs
	if attributeshape.Name != attributeshapeOther.Name {
		diffs = append(diffs, attributeshape.GongMarshallField(stage, "Name"))
	}
	if attributeshape.IdentifierMeta != attributeshapeOther.IdentifierMeta {
		diffs = append(diffs, attributeshape.GongMarshallField(stage, "IdentifierMeta"))
	}
	if attributeshape.FieldTypeAsString != attributeshapeOther.FieldTypeAsString {
		diffs = append(diffs, attributeshape.GongMarshallField(stage, "FieldTypeAsString"))
	}
	if attributeshape.Structname != attributeshapeOther.Structname {
		diffs = append(diffs, attributeshape.GongMarshallField(stage, "Structname"))
	}
	if attributeshape.Fieldtypename != attributeshapeOther.Fieldtypename {
		diffs = append(diffs, attributeshape.GongMarshallField(stage, "Fieldtypename"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (classdiagram *Classdiagram) GongDiff(stage *Stage, classdiagramOther *Classdiagram) (diffs []string) {
	// insertion point for field diffs
	if classdiagram.Name != classdiagramOther.Name {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "Name"))
	}
	if classdiagram.Description != classdiagramOther.Description {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "Description"))
	}
	if classdiagram.IsIncludedInStaticWebSite != classdiagramOther.IsIncludedInStaticWebSite {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "IsIncludedInStaticWebSite"))
	}
	GongStructShapesDifferent := false
	if len(classdiagram.GongStructShapes) != len(classdiagramOther.GongStructShapes) {
		GongStructShapesDifferent = true
	} else {
		for i := range classdiagram.GongStructShapes {
			if (classdiagram.GongStructShapes[i] == nil) != (classdiagramOther.GongStructShapes[i] == nil) {
				GongStructShapesDifferent = true
				break
			} else if classdiagram.GongStructShapes[i] != nil && classdiagramOther.GongStructShapes[i] != nil {
				// this is a pointer comparaison
				if classdiagram.GongStructShapes[i] != classdiagramOther.GongStructShapes[i] {
					GongStructShapesDifferent = true
					break
				}
			}
		}
	}
	if GongStructShapesDifferent {
		ops := Diff(stage, classdiagram, classdiagramOther, "GongStructShapes", classdiagramOther.GongStructShapes, classdiagram.GongStructShapes)
		diffs = append(diffs, ops)
	}
	GongEnumShapesDifferent := false
	if len(classdiagram.GongEnumShapes) != len(classdiagramOther.GongEnumShapes) {
		GongEnumShapesDifferent = true
	} else {
		for i := range classdiagram.GongEnumShapes {
			if (classdiagram.GongEnumShapes[i] == nil) != (classdiagramOther.GongEnumShapes[i] == nil) {
				GongEnumShapesDifferent = true
				break
			} else if classdiagram.GongEnumShapes[i] != nil && classdiagramOther.GongEnumShapes[i] != nil {
				// this is a pointer comparaison
				if classdiagram.GongEnumShapes[i] != classdiagramOther.GongEnumShapes[i] {
					GongEnumShapesDifferent = true
					break
				}
			}
		}
	}
	if GongEnumShapesDifferent {
		ops := Diff(stage, classdiagram, classdiagramOther, "GongEnumShapes", classdiagramOther.GongEnumShapes, classdiagram.GongEnumShapes)
		diffs = append(diffs, ops)
	}
	GongNoteShapesDifferent := false
	if len(classdiagram.GongNoteShapes) != len(classdiagramOther.GongNoteShapes) {
		GongNoteShapesDifferent = true
	} else {
		for i := range classdiagram.GongNoteShapes {
			if (classdiagram.GongNoteShapes[i] == nil) != (classdiagramOther.GongNoteShapes[i] == nil) {
				GongNoteShapesDifferent = true
				break
			} else if classdiagram.GongNoteShapes[i] != nil && classdiagramOther.GongNoteShapes[i] != nil {
				// this is a pointer comparaison
				if classdiagram.GongNoteShapes[i] != classdiagramOther.GongNoteShapes[i] {
					GongNoteShapesDifferent = true
					break
				}
			}
		}
	}
	if GongNoteShapesDifferent {
		ops := Diff(stage, classdiagram, classdiagramOther, "GongNoteShapes", classdiagramOther.GongNoteShapes, classdiagram.GongNoteShapes)
		diffs = append(diffs, ops)
	}
	if classdiagram.ShowNbInstances != classdiagramOther.ShowNbInstances {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "ShowNbInstances"))
	}
	if classdiagram.ShowMultiplicity != classdiagramOther.ShowMultiplicity {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "ShowMultiplicity"))
	}
	if classdiagram.ShowLinkNames != classdiagramOther.ShowLinkNames {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "ShowLinkNames"))
	}
	if classdiagram.IsInRenameMode != classdiagramOther.IsInRenameMode {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "IsInRenameMode"))
	}
	if classdiagram.IsExpanded != classdiagramOther.IsExpanded {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "IsExpanded"))
	}
	if classdiagram.NodeGongStructsIsExpanded != classdiagramOther.NodeGongStructsIsExpanded {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "NodeGongStructsIsExpanded"))
	}
	if classdiagram.NodeGongStructNodeExpansion != classdiagramOther.NodeGongStructNodeExpansion {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "NodeGongStructNodeExpansion"))
	}
	if classdiagram.NodeGongEnumsIsExpanded != classdiagramOther.NodeGongEnumsIsExpanded {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "NodeGongEnumsIsExpanded"))
	}
	if classdiagram.NodeGongEnumNodeExpansion != classdiagramOther.NodeGongEnumNodeExpansion {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "NodeGongEnumNodeExpansion"))
	}
	if classdiagram.NodeGongNotesIsExpanded != classdiagramOther.NodeGongNotesIsExpanded {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "NodeGongNotesIsExpanded"))
	}
	if classdiagram.NodeGongNoteNodeExpansion != classdiagramOther.NodeGongNoteNodeExpansion {
		diffs = append(diffs, classdiagram.GongMarshallField(stage, "NodeGongNoteNodeExpansion"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagrampackage *DiagramPackage) GongDiff(stage *Stage, diagrampackageOther *DiagramPackage) (diffs []string) {
	// insertion point for field diffs
	if diagrampackage.Name != diagrampackageOther.Name {
		diffs = append(diffs, diagrampackage.GongMarshallField(stage, "Name"))
	}
	if diagrampackage.Path != diagrampackageOther.Path {
		diffs = append(diffs, diagrampackage.GongMarshallField(stage, "Path"))
	}
	if diagrampackage.GongModelPath != diagrampackageOther.GongModelPath {
		diffs = append(diffs, diagrampackage.GongMarshallField(stage, "GongModelPath"))
	}
	ClassdiagramsDifferent := false
	if len(diagrampackage.Classdiagrams) != len(diagrampackageOther.Classdiagrams) {
		ClassdiagramsDifferent = true
	} else {
		for i := range diagrampackage.Classdiagrams {
			if (diagrampackage.Classdiagrams[i] == nil) != (diagrampackageOther.Classdiagrams[i] == nil) {
				ClassdiagramsDifferent = true
				break
			} else if diagrampackage.Classdiagrams[i] != nil && diagrampackageOther.Classdiagrams[i] != nil {
				// this is a pointer comparaison
				if diagrampackage.Classdiagrams[i] != diagrampackageOther.Classdiagrams[i] {
					ClassdiagramsDifferent = true
					break
				}
			}
		}
	}
	if ClassdiagramsDifferent {
		ops := Diff(stage, diagrampackage, diagrampackageOther, "Classdiagrams", diagrampackageOther.Classdiagrams, diagrampackage.Classdiagrams)
		diffs = append(diffs, ops)
	}
	if (diagrampackage.SelectedClassdiagram == nil) != (diagrampackageOther.SelectedClassdiagram == nil) {
		diffs = append(diffs, diagrampackage.GongMarshallField(stage, "SelectedClassdiagram"))
	} else if diagrampackage.SelectedClassdiagram != nil && diagrampackageOther.SelectedClassdiagram != nil {
		if diagrampackage.SelectedClassdiagram != diagrampackageOther.SelectedClassdiagram {
			diffs = append(diffs, diagrampackage.GongMarshallField(stage, "SelectedClassdiagram"))
		}
	}
	if diagrampackage.AbsolutePathToDiagramPackage != diagrampackageOther.AbsolutePathToDiagramPackage {
		diffs = append(diffs, diagrampackage.GongMarshallField(stage, "AbsolutePathToDiagramPackage"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongenumshape *GongEnumShape) GongDiff(stage *Stage, gongenumshapeOther *GongEnumShape) (diffs []string) {
	// insertion point for field diffs
	if gongenumshape.Name != gongenumshapeOther.Name {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "Name"))
	}
	if gongenumshape.X != gongenumshapeOther.X {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "X"))
	}
	if gongenumshape.Y != gongenumshapeOther.Y {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "Y"))
	}
	if gongenumshape.IdentifierMeta != gongenumshapeOther.IdentifierMeta {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "IdentifierMeta"))
	}
	GongEnumValueShapesDifferent := false
	if len(gongenumshape.GongEnumValueShapes) != len(gongenumshapeOther.GongEnumValueShapes) {
		GongEnumValueShapesDifferent = true
	} else {
		for i := range gongenumshape.GongEnumValueShapes {
			if (gongenumshape.GongEnumValueShapes[i] == nil) != (gongenumshapeOther.GongEnumValueShapes[i] == nil) {
				GongEnumValueShapesDifferent = true
				break
			} else if gongenumshape.GongEnumValueShapes[i] != nil && gongenumshapeOther.GongEnumValueShapes[i] != nil {
				// this is a pointer comparaison
				if gongenumshape.GongEnumValueShapes[i] != gongenumshapeOther.GongEnumValueShapes[i] {
					GongEnumValueShapesDifferent = true
					break
				}
			}
		}
	}
	if GongEnumValueShapesDifferent {
		ops := Diff(stage, gongenumshape, gongenumshapeOther, "GongEnumValueShapes", gongenumshapeOther.GongEnumValueShapes, gongenumshape.GongEnumValueShapes)
		diffs = append(diffs, ops)
	}
	if gongenumshape.Width != gongenumshapeOther.Width {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "Width"))
	}
	if gongenumshape.Height != gongenumshapeOther.Height {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "Height"))
	}
	if gongenumshape.IsExpanded != gongenumshapeOther.IsExpanded {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongenumvalueshape *GongEnumValueShape) GongDiff(stage *Stage, gongenumvalueshapeOther *GongEnumValueShape) (diffs []string) {
	// insertion point for field diffs
	if gongenumvalueshape.Name != gongenumvalueshapeOther.Name {
		diffs = append(diffs, gongenumvalueshape.GongMarshallField(stage, "Name"))
	}
	if gongenumvalueshape.IdentifierMeta != gongenumvalueshapeOther.IdentifierMeta {
		diffs = append(diffs, gongenumvalueshape.GongMarshallField(stage, "IdentifierMeta"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongnotelinkshape *GongNoteLinkShape) GongDiff(stage *Stage, gongnotelinkshapeOther *GongNoteLinkShape) (diffs []string) {
	// insertion point for field diffs
	if gongnotelinkshape.Name != gongnotelinkshapeOther.Name {
		diffs = append(diffs, gongnotelinkshape.GongMarshallField(stage, "Name"))
	}
	if gongnotelinkshape.Identifier != gongnotelinkshapeOther.Identifier {
		diffs = append(diffs, gongnotelinkshape.GongMarshallField(stage, "Identifier"))
	}
	if gongnotelinkshape.Type != gongnotelinkshapeOther.Type {
		diffs = append(diffs, gongnotelinkshape.GongMarshallField(stage, "Type"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongnoteshape *GongNoteShape) GongDiff(stage *Stage, gongnoteshapeOther *GongNoteShape) (diffs []string) {
	// insertion point for field diffs
	if gongnoteshape.Name != gongnoteshapeOther.Name {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "Name"))
	}
	if gongnoteshape.Identifier != gongnoteshapeOther.Identifier {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "Identifier"))
	}
	if gongnoteshape.Body != gongnoteshapeOther.Body {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "Body"))
	}
	if gongnoteshape.BodyHTML != gongnoteshapeOther.BodyHTML {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "BodyHTML"))
	}
	if gongnoteshape.X != gongnoteshapeOther.X {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "X"))
	}
	if gongnoteshape.Y != gongnoteshapeOther.Y {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "Y"))
	}
	if gongnoteshape.Width != gongnoteshapeOther.Width {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "Width"))
	}
	if gongnoteshape.Height != gongnoteshapeOther.Height {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "Height"))
	}
	if gongnoteshape.Matched != gongnoteshapeOther.Matched {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "Matched"))
	}
	GongNoteLinkShapesDifferent := false
	if len(gongnoteshape.GongNoteLinkShapes) != len(gongnoteshapeOther.GongNoteLinkShapes) {
		GongNoteLinkShapesDifferent = true
	} else {
		for i := range gongnoteshape.GongNoteLinkShapes {
			if (gongnoteshape.GongNoteLinkShapes[i] == nil) != (gongnoteshapeOther.GongNoteLinkShapes[i] == nil) {
				GongNoteLinkShapesDifferent = true
				break
			} else if gongnoteshape.GongNoteLinkShapes[i] != nil && gongnoteshapeOther.GongNoteLinkShapes[i] != nil {
				// this is a pointer comparaison
				if gongnoteshape.GongNoteLinkShapes[i] != gongnoteshapeOther.GongNoteLinkShapes[i] {
					GongNoteLinkShapesDifferent = true
					break
				}
			}
		}
	}
	if GongNoteLinkShapesDifferent {
		ops := Diff(stage, gongnoteshape, gongnoteshapeOther, "GongNoteLinkShapes", gongnoteshapeOther.GongNoteLinkShapes, gongnoteshape.GongNoteLinkShapes)
		diffs = append(diffs, ops)
	}
	if gongnoteshape.IsExpanded != gongnoteshapeOther.IsExpanded {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongstructshape *GongStructShape) GongDiff(stage *Stage, gongstructshapeOther *GongStructShape) (diffs []string) {
	// insertion point for field diffs
	if gongstructshape.Name != gongstructshapeOther.Name {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "Name"))
	}
	if gongstructshape.X != gongstructshapeOther.X {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "X"))
	}
	if gongstructshape.Y != gongstructshapeOther.Y {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "Y"))
	}
	if gongstructshape.IdentifierMeta != gongstructshapeOther.IdentifierMeta {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "IdentifierMeta"))
	}
	AttributeShapesDifferent := false
	if len(gongstructshape.AttributeShapes) != len(gongstructshapeOther.AttributeShapes) {
		AttributeShapesDifferent = true
	} else {
		for i := range gongstructshape.AttributeShapes {
			if (gongstructshape.AttributeShapes[i] == nil) != (gongstructshapeOther.AttributeShapes[i] == nil) {
				AttributeShapesDifferent = true
				break
			} else if gongstructshape.AttributeShapes[i] != nil && gongstructshapeOther.AttributeShapes[i] != nil {
				// this is a pointer comparaison
				if gongstructshape.AttributeShapes[i] != gongstructshapeOther.AttributeShapes[i] {
					AttributeShapesDifferent = true
					break
				}
			}
		}
	}
	if AttributeShapesDifferent {
		ops := Diff(stage, gongstructshape, gongstructshapeOther, "AttributeShapes", gongstructshapeOther.AttributeShapes, gongstructshape.AttributeShapes)
		diffs = append(diffs, ops)
	}
	LinkShapesDifferent := false
	if len(gongstructshape.LinkShapes) != len(gongstructshapeOther.LinkShapes) {
		LinkShapesDifferent = true
	} else {
		for i := range gongstructshape.LinkShapes {
			if (gongstructshape.LinkShapes[i] == nil) != (gongstructshapeOther.LinkShapes[i] == nil) {
				LinkShapesDifferent = true
				break
			} else if gongstructshape.LinkShapes[i] != nil && gongstructshapeOther.LinkShapes[i] != nil {
				// this is a pointer comparaison
				if gongstructshape.LinkShapes[i] != gongstructshapeOther.LinkShapes[i] {
					LinkShapesDifferent = true
					break
				}
			}
		}
	}
	if LinkShapesDifferent {
		ops := Diff(stage, gongstructshape, gongstructshapeOther, "LinkShapes", gongstructshapeOther.LinkShapes, gongstructshape.LinkShapes)
		diffs = append(diffs, ops)
	}
	if gongstructshape.Width != gongstructshapeOther.Width {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "Width"))
	}
	if gongstructshape.Height != gongstructshapeOther.Height {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "Height"))
	}
	if gongstructshape.IsSelected != gongstructshapeOther.IsSelected {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "IsSelected"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (linkshape *LinkShape) GongDiff(stage *Stage, linkshapeOther *LinkShape) (diffs []string) {
	// insertion point for field diffs
	if linkshape.Name != linkshapeOther.Name {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "Name"))
	}
	if linkshape.IdentifierMeta != linkshapeOther.IdentifierMeta {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "IdentifierMeta"))
	}
	if linkshape.FieldTypeIdentifierMeta != linkshapeOther.FieldTypeIdentifierMeta {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "FieldTypeIdentifierMeta"))
	}
	if linkshape.FieldOffsetX != linkshapeOther.FieldOffsetX {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "FieldOffsetX"))
	}
	if linkshape.FieldOffsetY != linkshapeOther.FieldOffsetY {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "FieldOffsetY"))
	}
	if linkshape.TargetMultiplicity != linkshapeOther.TargetMultiplicity {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "TargetMultiplicity"))
	}
	if linkshape.TargetMultiplicityOffsetX != linkshapeOther.TargetMultiplicityOffsetX {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "TargetMultiplicityOffsetX"))
	}
	if linkshape.TargetMultiplicityOffsetY != linkshapeOther.TargetMultiplicityOffsetY {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "TargetMultiplicityOffsetY"))
	}
	if linkshape.SourceMultiplicity != linkshapeOther.SourceMultiplicity {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "SourceMultiplicity"))
	}
	if linkshape.SourceMultiplicityOffsetX != linkshapeOther.SourceMultiplicityOffsetX {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "SourceMultiplicityOffsetX"))
	}
	if linkshape.SourceMultiplicityOffsetY != linkshapeOther.SourceMultiplicityOffsetY {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "SourceMultiplicityOffsetY"))
	}
	if linkshape.X != linkshapeOther.X {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "X"))
	}
	if linkshape.Y != linkshapeOther.Y {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "Y"))
	}
	if linkshape.StartOrientation != linkshapeOther.StartOrientation {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "StartOrientation"))
	}
	if linkshape.StartRatio != linkshapeOther.StartRatio {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "StartRatio"))
	}
	if linkshape.EndOrientation != linkshapeOther.EndOrientation {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "EndOrientation"))
	}
	if linkshape.EndRatio != linkshapeOther.EndRatio {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "EndRatio"))
	}
	if linkshape.CornerOffsetRatio != linkshapeOther.CornerOffsetRatio {
		diffs = append(diffs, linkshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
