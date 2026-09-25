// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (attributeshape *AttributeShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AttributeShapes[attributeshape]
	return ok
}

func (classdiagram *Classdiagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Classdiagrams[classdiagram]
	return ok
}

func (diagrampackage *DiagramPackage) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DiagramPackages[diagrampackage]
	return ok
}

func (gongenumshape *GongEnumShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongEnumShapes[gongenumshape]
	return ok
}

func (gongenumvalueshape *GongEnumValueShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongEnumValueShapes[gongenumvalueshape]
	return ok
}

func (gongnotelinkshape *GongNoteLinkShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongNoteLinkShapes[gongnotelinkshape]
	return ok
}

func (gongnoteshape *GongNoteShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongNoteShapes[gongnoteshape]
	return ok
}

func (gongstructshape *GongStructShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongStructShapes[gongstructshape]
	return ok
}

func (linkshape *LinkShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LinkShapes[linkshape]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (attributeshape *AttributeShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(attributeshape) {
		return
	}

	attributeshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (classdiagram *Classdiagram) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(classdiagram) {
		return
	}

	classdiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongstructshape := range classdiagram.GongStructShapes {
		stage.StageBranch(_gongstructshape)
	}
	for _, _gongenumshape := range classdiagram.GongEnumShapes {
		stage.StageBranch(_gongenumshape)
	}
	for _, _gongnoteshape := range classdiagram.GongNoteShapes {
		stage.StageBranch(_gongnoteshape)
	}

}

func (diagrampackage *DiagramPackage) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(diagrampackage) {
		return
	}

	diagrampackage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagrampackage.SelectedClassdiagram != nil {
		stage.StageBranch(diagrampackage.SelectedClassdiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _classdiagram := range diagrampackage.Classdiagrams {
		stage.StageBranch(_classdiagram)
	}

}

func (gongenumshape *GongEnumShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gongenumshape) {
		return
	}

	gongenumshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
		stage.StageBranch(_gongenumvalueshape)
	}

}

func (gongenumvalueshape *GongEnumValueShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gongenumvalueshape) {
		return
	}

	gongenumvalueshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongnotelinkshape *GongNoteLinkShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gongnotelinkshape) {
		return
	}

	gongnotelinkshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongnoteshape *GongNoteShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gongnoteshape) {
		return
	}

	gongnoteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
		stage.StageBranch(_gongnotelinkshape)
	}

}

func (gongstructshape *GongStructShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gongstructshape) {
		return
	}

	gongstructshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributeshape := range gongstructshape.AttributeShapes {
		stage.StageBranch(_attributeshape)
	}
	for _, _linkshape := range gongstructshape.LinkShapes {
		stage.StageBranch(_linkshape)
	}

}

func (linkshape *LinkShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(linkshape) {
		return
	}

	linkshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AttributeShape:
		toT := GongCopyBranchAttributeShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Classdiagram:
		toT := GongCopyBranchClassdiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramPackage:
		toT := GongCopyBranchDiagramPackage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnumShape:
		toT := GongCopyBranchGongEnumShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnumValueShape:
		toT := GongCopyBranchGongEnumValueShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongNoteLinkShape:
		toT := GongCopyBranchGongNoteLinkShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongNoteShape:
		toT := GongCopyBranchGongNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongStructShape:
		toT := GongCopyBranchGongStructShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinkShape:
		toT := GongCopyBranchLinkShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAttributeShape(mapOrigCopy map[any]any, attributeshapeFrom *AttributeShape) (attributeshapeTo *AttributeShape) {
	var alreadyCopied bool
	attributeshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, attributeshapeFrom)
	if alreadyCopied {
		return
	}
	attributeshapeFrom.GongCopyBasicFields(attributeshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchClassdiagram(mapOrigCopy map[any]any, classdiagramFrom *Classdiagram) (classdiagramTo *Classdiagram) {
	var alreadyCopied bool
	classdiagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, classdiagramFrom)
	if alreadyCopied {
		return
	}
	classdiagramFrom.GongCopyBasicFields(classdiagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongstructshape := range classdiagramFrom.GongStructShapes {
		classdiagramTo.GongStructShapes = append(classdiagramTo.GongStructShapes, GongCopyBranchGongStructShape(mapOrigCopy, _gongstructshape))
	}
	for _, _gongenumshape := range classdiagramFrom.GongEnumShapes {
		classdiagramTo.GongEnumShapes = append(classdiagramTo.GongEnumShapes, GongCopyBranchGongEnumShape(mapOrigCopy, _gongenumshape))
	}
	for _, _gongnoteshape := range classdiagramFrom.GongNoteShapes {
		classdiagramTo.GongNoteShapes = append(classdiagramTo.GongNoteShapes, GongCopyBranchGongNoteShape(mapOrigCopy, _gongnoteshape))
	}

	return
}

func GongCopyBranchDiagramPackage(mapOrigCopy map[any]any, diagrampackageFrom *DiagramPackage) (diagrampackageTo *DiagramPackage) {
	var alreadyCopied bool
	diagrampackageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagrampackageFrom)
	if alreadyCopied {
		return
	}
	diagrampackageFrom.GongCopyBasicFields(diagrampackageTo)

	//insertion point for the staging of instances referenced by pointers
	if diagrampackageFrom.SelectedClassdiagram != nil {
		diagrampackageTo.SelectedClassdiagram = GongCopyBranchClassdiagram(mapOrigCopy, diagrampackageFrom.SelectedClassdiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _classdiagram := range diagrampackageFrom.Classdiagrams {
		diagrampackageTo.Classdiagrams = append(diagrampackageTo.Classdiagrams, GongCopyBranchClassdiagram(mapOrigCopy, _classdiagram))
	}

	return
}

func GongCopyBranchGongEnumShape(mapOrigCopy map[any]any, gongenumshapeFrom *GongEnumShape) (gongenumshapeTo *GongEnumShape) {
	var alreadyCopied bool
	gongenumshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongenumshapeFrom)
	if alreadyCopied {
		return
	}
	gongenumshapeFrom.GongCopyBasicFields(gongenumshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueshape := range gongenumshapeFrom.GongEnumValueShapes {
		gongenumshapeTo.GongEnumValueShapes = append(gongenumshapeTo.GongEnumValueShapes, GongCopyBranchGongEnumValueShape(mapOrigCopy, _gongenumvalueshape))
	}

	return
}

func GongCopyBranchGongEnumValueShape(mapOrigCopy map[any]any, gongenumvalueshapeFrom *GongEnumValueShape) (gongenumvalueshapeTo *GongEnumValueShape) {
	var alreadyCopied bool
	gongenumvalueshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongenumvalueshapeFrom)
	if alreadyCopied {
		return
	}
	gongenumvalueshapeFrom.GongCopyBasicFields(gongenumvalueshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGongNoteLinkShape(mapOrigCopy map[any]any, gongnotelinkshapeFrom *GongNoteLinkShape) (gongnotelinkshapeTo *GongNoteLinkShape) {
	var alreadyCopied bool
	gongnotelinkshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongnotelinkshapeFrom)
	if alreadyCopied {
		return
	}
	gongnotelinkshapeFrom.GongCopyBasicFields(gongnotelinkshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGongNoteShape(mapOrigCopy map[any]any, gongnoteshapeFrom *GongNoteShape) (gongnoteshapeTo *GongNoteShape) {
	var alreadyCopied bool
	gongnoteshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongnoteshapeFrom)
	if alreadyCopied {
		return
	}
	gongnoteshapeFrom.GongCopyBasicFields(gongnoteshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongnotelinkshape := range gongnoteshapeFrom.GongNoteLinkShapes {
		gongnoteshapeTo.GongNoteLinkShapes = append(gongnoteshapeTo.GongNoteLinkShapes, GongCopyBranchGongNoteLinkShape(mapOrigCopy, _gongnotelinkshape))
	}

	return
}

func GongCopyBranchGongStructShape(mapOrigCopy map[any]any, gongstructshapeFrom *GongStructShape) (gongstructshapeTo *GongStructShape) {
	var alreadyCopied bool
	gongstructshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongstructshapeFrom)
	if alreadyCopied {
		return
	}
	gongstructshapeFrom.GongCopyBasicFields(gongstructshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributeshape := range gongstructshapeFrom.AttributeShapes {
		gongstructshapeTo.AttributeShapes = append(gongstructshapeTo.AttributeShapes, GongCopyBranchAttributeShape(mapOrigCopy, _attributeshape))
	}
	for _, _linkshape := range gongstructshapeFrom.LinkShapes {
		gongstructshapeTo.LinkShapes = append(gongstructshapeTo.LinkShapes, GongCopyBranchLinkShape(mapOrigCopy, _linkshape))
	}

	return
}

func GongCopyBranchLinkShape(mapOrigCopy map[any]any, linkshapeFrom *LinkShape) (linkshapeTo *LinkShape) {
	var alreadyCopied bool
	linkshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, linkshapeFrom)
	if alreadyCopied {
		return
	}
	linkshapeFrom.GongCopyBasicFields(linkshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (attributeshape *AttributeShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(attributeshape) {
		return
	}

	attributeshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (classdiagram *Classdiagram) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(classdiagram) {
		return
	}

	classdiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongstructshape := range classdiagram.GongStructShapes {
		stage.UnstageBranch(_gongstructshape)
	}
	for _, _gongenumshape := range classdiagram.GongEnumShapes {
		stage.UnstageBranch(_gongenumshape)
	}
	for _, _gongnoteshape := range classdiagram.GongNoteShapes {
		stage.UnstageBranch(_gongnoteshape)
	}

}

func (diagrampackage *DiagramPackage) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(diagrampackage) {
		return
	}

	diagrampackage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagrampackage.SelectedClassdiagram != nil {
		stage.UnstageBranch(diagrampackage.SelectedClassdiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _classdiagram := range diagrampackage.Classdiagrams {
		stage.UnstageBranch(_classdiagram)
	}

}

func (gongenumshape *GongEnumShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gongenumshape) {
		return
	}

	gongenumshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
		stage.UnstageBranch(_gongenumvalueshape)
	}

}

func (gongenumvalueshape *GongEnumValueShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gongenumvalueshape) {
		return
	}

	gongenumvalueshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongnotelinkshape *GongNoteLinkShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gongnotelinkshape) {
		return
	}

	gongnotelinkshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongnoteshape *GongNoteShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gongnoteshape) {
		return
	}

	gongnoteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
		stage.UnstageBranch(_gongnotelinkshape)
	}

}

func (gongstructshape *GongStructShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gongstructshape) {
		return
	}

	gongstructshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attributeshape := range gongstructshape.AttributeShapes {
		stage.UnstageBranch(_attributeshape)
	}
	for _, _linkshape := range gongstructshape.LinkShapes {
		stage.UnstageBranch(_linkshape)
	}

}

func (linkshape *LinkShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(linkshape) {
		return
	}

	linkshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AttributeShape) GongReconstructPointersFromReferences(stage *Stage, instance *AttributeShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Classdiagram) GongReconstructPointersFromReferences(stage *Stage, instance *Classdiagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GongStructShapes, stage.GongStructShapes_reference, instance.GongStructShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GongEnumShapes, stage.GongEnumShapes_reference, instance.GongEnumShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GongNoteShapes, stage.GongNoteShapes_reference, instance.GongNoteShapes)
}

func (reference *DiagramPackage) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramPackage) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.SelectedClassdiagram, stage.Classdiagrams_reference, instance.SelectedClassdiagram)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Classdiagrams, stage.Classdiagrams_reference, instance.Classdiagrams)
}

func (reference *GongEnumShape) GongReconstructPointersFromReferences(stage *Stage, instance *GongEnumShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GongEnumValueShapes, stage.GongEnumValueShapes_reference, instance.GongEnumValueShapes)
}

func (reference *GongEnumValueShape) GongReconstructPointersFromReferences(stage *Stage, instance *GongEnumValueShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GongNoteLinkShape) GongReconstructPointersFromReferences(stage *Stage, instance *GongNoteLinkShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GongNoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *GongNoteShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GongNoteLinkShapes, stage.GongNoteLinkShapes_reference, instance.GongNoteLinkShapes)
}

func (reference *GongStructShape) GongReconstructPointersFromReferences(stage *Stage, instance *GongStructShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.AttributeShapes, stage.AttributeShapes_reference, instance.AttributeShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.LinkShapes, stage.LinkShapes_reference, instance.LinkShapes)
}

func (reference *LinkShape) GongReconstructPointersFromReferences(stage *Stage, instance *LinkShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AttributeShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Classdiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GongStructShapes, stage.GongStructShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GongEnumShapes, stage.GongEnumShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GongNoteShapes, stage.GongNoteShapes_instance)
}

func (reference *DiagramPackage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.SelectedClassdiagram, stage.Classdiagrams_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Classdiagrams, stage.Classdiagrams_instance)
}

func (reference *GongEnumShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GongEnumValueShapes, stage.GongEnumValueShapes_instance)
}

func (reference *GongEnumValueShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GongNoteLinkShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GongNoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GongNoteLinkShapes, stage.GongNoteLinkShapes_instance)
}

func (reference *GongStructShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.AttributeShapes, stage.AttributeShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.LinkShapes, stage.LinkShapes_instance)
}

func (reference *LinkShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
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
	if ops := __gong__diffSliceOfPointers(stage, classdiagram, "GongStructShapes", classdiagramOther.GongStructShapes, classdiagram.GongStructShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, classdiagram, "GongEnumShapes", classdiagramOther.GongEnumShapes, classdiagram.GongEnumShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, classdiagram, "GongNoteShapes", classdiagramOther.GongNoteShapes, classdiagram.GongNoteShapes); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, diagrampackage, "Classdiagrams", diagrampackageOther.Classdiagrams, diagrampackage.Classdiagrams); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagrampackage.SelectedClassdiagram != diagrampackageOther.SelectedClassdiagram {
		diffs = append(diffs, diagrampackage.GongMarshallField(stage, "SelectedClassdiagram"))
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
	if gongenumshape.Width != gongenumshapeOther.Width {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "Width"))
	}
	if gongenumshape.Height != gongenumshapeOther.Height {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "Height"))
	}
	if gongenumshape.IsHidden != gongenumshapeOther.IsHidden {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "IsHidden"))
	}
	if gongenumshape.IdentifierMeta != gongenumshapeOther.IdentifierMeta {
		diffs = append(diffs, gongenumshape.GongMarshallField(stage, "IdentifierMeta"))
	}
	if ops := __gong__diffSliceOfPointers(stage, gongenumshape, "GongEnumValueShapes", gongenumshapeOther.GongEnumValueShapes, gongenumshape.GongEnumValueShapes); ops != "" {
		diffs = append(diffs, ops)
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
	if gongnoteshape.IsHidden != gongnoteshapeOther.IsHidden {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "IsHidden"))
	}
	if gongnoteshape.Matched != gongnoteshapeOther.Matched {
		diffs = append(diffs, gongnoteshape.GongMarshallField(stage, "Matched"))
	}
	if ops := __gong__diffSliceOfPointers(stage, gongnoteshape, "GongNoteLinkShapes", gongnoteshapeOther.GongNoteLinkShapes, gongnoteshape.GongNoteLinkShapes); ops != "" {
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
	if gongstructshape.Width != gongstructshapeOther.Width {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "Width"))
	}
	if gongstructshape.Height != gongstructshapeOther.Height {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "Height"))
	}
	if gongstructshape.IsHidden != gongstructshapeOther.IsHidden {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "IsHidden"))
	}
	if gongstructshape.IdentifierMeta != gongstructshapeOther.IdentifierMeta {
		diffs = append(diffs, gongstructshape.GongMarshallField(stage, "IdentifierMeta"))
	}
	if ops := __gong__diffSliceOfPointers(stage, gongstructshape, "AttributeShapes", gongstructshapeOther.AttributeShapes, gongstructshape.AttributeShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, gongstructshape, "LinkShapes", gongstructshapeOther.LinkShapes, gongstructshape.LinkShapes); ops != "" {
		diffs = append(diffs, ops)
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

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
