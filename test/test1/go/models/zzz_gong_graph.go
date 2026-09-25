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
func (astruct *Astruct) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Astructs[astruct]
	return ok
}

func (astructbstruct2use *AstructBstruct2Use) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AstructBstruct2Uses[astructbstruct2use]
	return ok
}

func (astructbstructuse *AstructBstructUse) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AstructBstructUses[astructbstructuse]
	return ok
}

func (bstruct *Bstruct) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Bstructs[bstruct]
	return ok
}

func (dstruct *Dstruct) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Dstructs[dstruct]
	return ok
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongIsStaged(stage *Stage) bool {
	_, ok := stage.F0123456789012345678901234567890s[f0123456789012345678901234567890]
	return ok
}

func (gstruct *Gstruct) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Gstructs[gstruct]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (astruct *Astruct) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(astruct) {
		return
	}

	astruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astruct.Associationtob != nil {
		stage.StageBranch(astruct.Associationtob)
	}
	if astruct.Anotherassociationtob_2 != nil {
		stage.StageBranch(astruct.Anotherassociationtob_2)
	}
	if astruct.Bstruct != nil {
		stage.StageBranch(astruct.Bstruct)
	}
	if astruct.Bstruct2 != nil {
		stage.StageBranch(astruct.Bstruct2)
	}
	if astruct.Dstruct != nil {
		stage.StageBranch(astruct.Dstruct)
	}
	if astruct.Dstruct2 != nil {
		stage.StageBranch(astruct.Dstruct2)
	}
	if astruct.Dstruct3 != nil {
		stage.StageBranch(astruct.Dstruct3)
	}
	if astruct.Dstruct4 != nil {
		stage.StageBranch(astruct.Dstruct4)
	}
	if astruct.AnAstruct != nil {
		stage.StageBranch(astruct.AnAstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range astruct.Anarrayofb {
		stage.StageBranch(_bstruct)
	}
	for _, _dstruct := range astruct.Dstruct4s {
		stage.StageBranch(_dstruct)
	}
	for _, _astruct := range astruct.Anarrayofa {
		stage.StageBranch(_astruct)
	}
	for _, _bstruct := range astruct.Anotherarrayofb {
		stage.StageBranch(_bstruct)
	}
	for _, _astructbstructuse := range astruct.AnarrayofbUse {
		stage.StageBranch(_astructbstructuse)
	}
	for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
		stage.StageBranch(_astructbstruct2use)
	}

}

func (astructbstruct2use *AstructBstruct2Use) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(astructbstruct2use) {
		return
	}

	astructbstruct2use.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astructbstruct2use.Bstrcut2 != nil {
		stage.StageBranch(astructbstruct2use.Bstrcut2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (astructbstructuse *AstructBstructUse) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(astructbstructuse) {
		return
	}

	astructbstructuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astructbstructuse.Bstruct2 != nil {
		stage.StageBranch(astructbstructuse.Bstruct2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bstruct *Bstruct) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(bstruct) {
		return
	}

	bstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dstruct *Dstruct) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(dstruct) {
		return
	}

	dstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dstruct.Gstruct != nil {
		stage.StageBranch(dstruct.Gstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstruct.Anarrayofb {
		stage.StageBranch(_bstruct)
	}
	for _, _gstruct := range dstruct.Gstructs {
		stage.StageBranch(_gstruct)
	}

}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(f0123456789012345678901234567890) {
		return
	}

	f0123456789012345678901234567890.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gstruct *Gstruct) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gstruct) {
		return
	}

	gstruct.Stage(stage)

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
	case *Astruct:
		toT := GongCopyBranchAstruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AstructBstruct2Use:
		toT := GongCopyBranchAstructBstruct2Use(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AstructBstructUse:
		toT := GongCopyBranchAstructBstructUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bstruct:
		toT := GongCopyBranchBstruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Dstruct:
		toT := GongCopyBranchDstruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *F0123456789012345678901234567890:
		toT := GongCopyBranchF0123456789012345678901234567890(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Gstruct:
		toT := GongCopyBranchGstruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAstruct(mapOrigCopy map[any]any, astructFrom *Astruct) (astructTo *Astruct) {
	var alreadyCopied bool
	astructTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, astructFrom)
	if alreadyCopied {
		return
	}
	astructFrom.GongCopyBasicFields(astructTo)

	//insertion point for the staging of instances referenced by pointers
	if astructFrom.Associationtob != nil {
		astructTo.Associationtob = GongCopyBranchBstruct(mapOrigCopy, astructFrom.Associationtob)
	}
	if astructFrom.Anotherassociationtob_2 != nil {
		astructTo.Anotherassociationtob_2 = astructFrom.Anotherassociationtob_2
	}
	if astructFrom.Bstruct != nil {
		astructTo.Bstruct = GongCopyBranchBstruct(mapOrigCopy, astructFrom.Bstruct)
	}
	if astructFrom.Bstruct2 != nil {
		astructTo.Bstruct2 = GongCopyBranchBstruct(mapOrigCopy, astructFrom.Bstruct2)
	}
	if astructFrom.Dstruct != nil {
		astructTo.Dstruct = GongCopyBranchDstruct(mapOrigCopy, astructFrom.Dstruct)
	}
	if astructFrom.Dstruct2 != nil {
		astructTo.Dstruct2 = GongCopyBranchDstruct(mapOrigCopy, astructFrom.Dstruct2)
	}
	if astructFrom.Dstruct3 != nil {
		astructTo.Dstruct3 = GongCopyBranchDstruct(mapOrigCopy, astructFrom.Dstruct3)
	}
	if astructFrom.Dstruct4 != nil {
		astructTo.Dstruct4 = GongCopyBranchDstruct(mapOrigCopy, astructFrom.Dstruct4)
	}
	if astructFrom.AnAstruct != nil {
		astructTo.AnAstruct = GongCopyBranchAstruct(mapOrigCopy, astructFrom.AnAstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range astructFrom.Anarrayofb {
		astructTo.Anarrayofb = append(astructTo.Anarrayofb, GongCopyBranchBstruct(mapOrigCopy, _bstruct))
	}
	for _, _dstruct := range astructFrom.Dstruct4s {
		astructTo.Dstruct4s = append(astructTo.Dstruct4s, GongCopyBranchDstruct(mapOrigCopy, _dstruct))
	}
	for _, _astruct := range astructFrom.Anarrayofa {
		astructTo.Anarrayofa = append(astructTo.Anarrayofa, GongCopyBranchAstruct(mapOrigCopy, _astruct))
	}
	for _, _bstruct := range astructFrom.Anotherarrayofb {
		astructTo.Anotherarrayofb = append(astructTo.Anotherarrayofb, GongCopyBranchBstruct(mapOrigCopy, _bstruct))
	}
	for _, _astructbstructuse := range astructFrom.AnarrayofbUse {
		astructTo.AnarrayofbUse = append(astructTo.AnarrayofbUse, GongCopyBranchAstructBstructUse(mapOrigCopy, _astructbstructuse))
	}
	for _, _astructbstruct2use := range astructFrom.Anarrayofb2Use {
		astructTo.Anarrayofb2Use = append(astructTo.Anarrayofb2Use, GongCopyBranchAstructBstruct2Use(mapOrigCopy, _astructbstruct2use))
	}

	return
}

func GongCopyBranchAstructBstruct2Use(mapOrigCopy map[any]any, astructbstruct2useFrom *AstructBstruct2Use) (astructbstruct2useTo *AstructBstruct2Use) {
	var alreadyCopied bool
	astructbstruct2useTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, astructbstruct2useFrom)
	if alreadyCopied {
		return
	}
	astructbstruct2useFrom.GongCopyBasicFields(astructbstruct2useTo)

	//insertion point for the staging of instances referenced by pointers
	if astructbstruct2useFrom.Bstrcut2 != nil {
		astructbstruct2useTo.Bstrcut2 = GongCopyBranchBstruct(mapOrigCopy, astructbstruct2useFrom.Bstrcut2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAstructBstructUse(mapOrigCopy map[any]any, astructbstructuseFrom *AstructBstructUse) (astructbstructuseTo *AstructBstructUse) {
	var alreadyCopied bool
	astructbstructuseTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, astructbstructuseFrom)
	if alreadyCopied {
		return
	}
	astructbstructuseFrom.GongCopyBasicFields(astructbstructuseTo)

	//insertion point for the staging of instances referenced by pointers
	if astructbstructuseFrom.Bstruct2 != nil {
		astructbstructuseTo.Bstruct2 = GongCopyBranchBstruct(mapOrigCopy, astructbstructuseFrom.Bstruct2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBstruct(mapOrigCopy map[any]any, bstructFrom *Bstruct) (bstructTo *Bstruct) {
	var alreadyCopied bool
	bstructTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, bstructFrom)
	if alreadyCopied {
		return
	}
	bstructFrom.GongCopyBasicFields(bstructTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDstruct(mapOrigCopy map[any]any, dstructFrom *Dstruct) (dstructTo *Dstruct) {
	var alreadyCopied bool
	dstructTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dstructFrom)
	if alreadyCopied {
		return
	}
	dstructFrom.GongCopyBasicFields(dstructTo)

	//insertion point for the staging of instances referenced by pointers
	if dstructFrom.Gstruct != nil {
		dstructTo.Gstruct = GongCopyBranchGstruct(mapOrigCopy, dstructFrom.Gstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstructFrom.Anarrayofb {
		dstructTo.Anarrayofb = append(dstructTo.Anarrayofb, GongCopyBranchBstruct(mapOrigCopy, _bstruct))
	}
	for _, _gstruct := range dstructFrom.Gstructs {
		dstructTo.Gstructs = append(dstructTo.Gstructs, GongCopyBranchGstruct(mapOrigCopy, _gstruct))
	}

	return
}

func GongCopyBranchF0123456789012345678901234567890(mapOrigCopy map[any]any, f0123456789012345678901234567890From *F0123456789012345678901234567890) (f0123456789012345678901234567890To *F0123456789012345678901234567890) {
	var alreadyCopied bool
	f0123456789012345678901234567890To, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, f0123456789012345678901234567890From)
	if alreadyCopied {
		return
	}
	f0123456789012345678901234567890From.GongCopyBasicFields(f0123456789012345678901234567890To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGstruct(mapOrigCopy map[any]any, gstructFrom *Gstruct) (gstructTo *Gstruct) {
	var alreadyCopied bool
	gstructTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gstructFrom)
	if alreadyCopied {
		return
	}
	gstructFrom.GongCopyBasicFields(gstructTo)

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
func (astruct *Astruct) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(astruct) {
		return
	}

	astruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astruct.Associationtob != nil {
		stage.UnstageBranch(astruct.Associationtob)
	}
	if astruct.Anotherassociationtob_2 != nil {
		stage.UnstageBranch(astruct.Anotherassociationtob_2)
	}
	if astruct.Bstruct != nil {
		stage.UnstageBranch(astruct.Bstruct)
	}
	if astruct.Bstruct2 != nil {
		stage.UnstageBranch(astruct.Bstruct2)
	}
	if astruct.Dstruct != nil {
		stage.UnstageBranch(astruct.Dstruct)
	}
	if astruct.Dstruct2 != nil {
		stage.UnstageBranch(astruct.Dstruct2)
	}
	if astruct.Dstruct3 != nil {
		stage.UnstageBranch(astruct.Dstruct3)
	}
	if astruct.Dstruct4 != nil {
		stage.UnstageBranch(astruct.Dstruct4)
	}
	if astruct.AnAstruct != nil {
		stage.UnstageBranch(astruct.AnAstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range astruct.Anarrayofb {
		stage.UnstageBranch(_bstruct)
	}
	for _, _dstruct := range astruct.Dstruct4s {
		stage.UnstageBranch(_dstruct)
	}
	for _, _astruct := range astruct.Anarrayofa {
		stage.UnstageBranch(_astruct)
	}
	for _, _bstruct := range astruct.Anotherarrayofb {
		stage.UnstageBranch(_bstruct)
	}
	for _, _astructbstructuse := range astruct.AnarrayofbUse {
		stage.UnstageBranch(_astructbstructuse)
	}
	for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
		stage.UnstageBranch(_astructbstruct2use)
	}

}

func (astructbstruct2use *AstructBstruct2Use) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(astructbstruct2use) {
		return
	}

	astructbstruct2use.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astructbstruct2use.Bstrcut2 != nil {
		stage.UnstageBranch(astructbstruct2use.Bstrcut2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (astructbstructuse *AstructBstructUse) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(astructbstructuse) {
		return
	}

	astructbstructuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astructbstructuse.Bstruct2 != nil {
		stage.UnstageBranch(astructbstructuse.Bstruct2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (bstruct *Bstruct) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(bstruct) {
		return
	}

	bstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dstruct *Dstruct) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(dstruct) {
		return
	}

	dstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dstruct.Gstruct != nil {
		stage.UnstageBranch(dstruct.Gstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstruct.Anarrayofb {
		stage.UnstageBranch(_bstruct)
	}
	for _, _gstruct := range dstruct.Gstructs {
		stage.UnstageBranch(_gstruct)
	}

}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(f0123456789012345678901234567890) {
		return
	}

	f0123456789012345678901234567890.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gstruct *Gstruct) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gstruct) {
		return
	}

	gstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Astruct) GongReconstructPointersFromReferences(stage *Stage, instance *Astruct) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Associationtob, stage.Bstructs_reference, instance.Associationtob)
	__gong__reconstructPointer(&reference.Anotherassociationtob_2, stage.Bstructs_reference, instance.Anotherassociationtob_2)
	__gong__reconstructPointer(&reference.Bstruct, stage.Bstructs_reference, instance.Bstruct)
	__gong__reconstructPointer(&reference.Bstruct2, stage.Bstructs_reference, instance.Bstruct2)
	__gong__reconstructPointer(&reference.Dstruct, stage.Dstructs_reference, instance.Dstruct)
	__gong__reconstructPointer(&reference.Dstruct2, stage.Dstructs_reference, instance.Dstruct2)
	__gong__reconstructPointer(&reference.Dstruct3, stage.Dstructs_reference, instance.Dstruct3)
	__gong__reconstructPointer(&reference.Dstruct4, stage.Dstructs_reference, instance.Dstruct4)
	__gong__reconstructPointer(&reference.AnAstruct, stage.Astructs_reference, instance.AnAstruct)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Anarrayofb, stage.Bstructs_reference, instance.Anarrayofb)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Dstruct4s, stage.Dstructs_reference, instance.Dstruct4s)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Anarrayofa, stage.Astructs_reference, instance.Anarrayofa)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Anotherarrayofb, stage.Bstructs_reference, instance.Anotherarrayofb)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AnarrayofbUse, stage.AstructBstructUses_reference, instance.AnarrayofbUse)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Anarrayofb2Use, stage.AstructBstruct2Uses_reference, instance.Anarrayofb2Use)
}

func (reference *AstructBstruct2Use) GongReconstructPointersFromReferences(stage *Stage, instance *AstructBstruct2Use) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Bstrcut2, stage.Bstructs_reference, instance.Bstrcut2)
	// insertion point for slice of pointers field
}

func (reference *AstructBstructUse) GongReconstructPointersFromReferences(stage *Stage, instance *AstructBstructUse) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Bstruct2, stage.Bstructs_reference, instance.Bstruct2)
	// insertion point for slice of pointers field
}

func (reference *Bstruct) GongReconstructPointersFromReferences(stage *Stage, instance *Bstruct) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Dstruct) GongReconstructPointersFromReferences(stage *Stage, instance *Dstruct) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Gstruct, stage.Gstructs_reference, instance.Gstruct)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Anarrayofb, stage.Bstructs_reference, instance.Anarrayofb)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Gstructs, stage.Gstructs_reference, instance.Gstructs)
}

func (reference *F0123456789012345678901234567890) GongReconstructPointersFromReferences(stage *Stage, instance *F0123456789012345678901234567890) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Gstruct) GongReconstructPointersFromReferences(stage *Stage, instance *Gstruct) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Astruct) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Associationtob, stage.Bstructs_instance)
	__gong__reconstructPointerFromInstance(&reference.Anotherassociationtob_2, stage.Bstructs_instance)
	__gong__reconstructPointerFromInstance(&reference.Bstruct, stage.Bstructs_instance)
	__gong__reconstructPointerFromInstance(&reference.Bstruct2, stage.Bstructs_instance)
	__gong__reconstructPointerFromInstance(&reference.Dstruct, stage.Dstructs_instance)
	__gong__reconstructPointerFromInstance(&reference.Dstruct2, stage.Dstructs_instance)
	__gong__reconstructPointerFromInstance(&reference.Dstruct3, stage.Dstructs_instance)
	__gong__reconstructPointerFromInstance(&reference.Dstruct4, stage.Dstructs_instance)
	__gong__reconstructPointerFromInstance(&reference.AnAstruct, stage.Astructs_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Anarrayofb, stage.Bstructs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Dstruct4s, stage.Dstructs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Anarrayofa, stage.Astructs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Anotherarrayofb, stage.Bstructs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AnarrayofbUse, stage.AstructBstructUses_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Anarrayofb2Use, stage.AstructBstruct2Uses_instance)
}

func (reference *AstructBstruct2Use) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Bstrcut2, stage.Bstructs_instance)
	// insertion point for slice of pointers fields
}

func (reference *AstructBstructUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Bstruct2, stage.Bstructs_instance)
	// insertion point for slice of pointers fields
}

func (reference *Bstruct) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Dstruct) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Gstruct, stage.Gstructs_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Anarrayofb, stage.Bstructs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Gstructs, stage.Gstructs_instance)
}

func (reference *F0123456789012345678901234567890) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Gstruct) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (astruct *Astruct) GongDiff(stage *Stage, astructOther *Astruct) (diffs []string) {
	// insertion point for field diffs
	if astruct.Name != astructOther.Name {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Name"))
	}
	if astruct.Field != astructOther.Field {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Field"))
	}
	if astruct.Associationtob != astructOther.Associationtob {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Associationtob"))
	}
	if ops := __gong__diffSliceOfPointers(stage, astruct, "Anarrayofb", astructOther.Anarrayofb, astruct.Anarrayofb); ops != "" {
		diffs = append(diffs, ops)
	}
	if astruct.Anotherassociationtob_2 != astructOther.Anotherassociationtob_2 {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Anotherassociationtob_2"))
	}
	if astruct.Date != astructOther.Date {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Date"))
	}
	if astruct.Date2 != astructOther.Date2 {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Date2"))
	}
	if astruct.Booleanfield != astructOther.Booleanfield {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Booleanfield"))
	}
	if astruct.Aenum != astructOther.Aenum {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Aenum"))
	}
	if astruct.Aenum_2 != astructOther.Aenum_2 {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Aenum_2"))
	}
	if astruct.Benum != astructOther.Benum {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Benum"))
	}
	if astruct.CEnum != astructOther.CEnum {
		diffs = append(diffs, astruct.GongMarshallField(stage, "CEnum"))
	}
	if astruct.CName != astructOther.CName {
		diffs = append(diffs, astruct.GongMarshallField(stage, "CName"))
	}
	if astruct.CFloatfield != astructOther.CFloatfield {
		diffs = append(diffs, astruct.GongMarshallField(stage, "CFloatfield"))
	}
	if astruct.Bstruct != astructOther.Bstruct {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Bstruct"))
	}
	if astruct.Bstruct2 != astructOther.Bstruct2 {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Bstruct2"))
	}
	if astruct.Dstruct != astructOther.Dstruct {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct"))
	}
	if astruct.Dstruct2 != astructOther.Dstruct2 {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct2"))
	}
	if astruct.Dstruct3 != astructOther.Dstruct3 {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct3"))
	}
	if astruct.Dstruct4 != astructOther.Dstruct4 {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct4"))
	}
	if ops := __gong__diffSliceOfPointers(stage, astruct, "Dstruct4s", astructOther.Dstruct4s, astruct.Dstruct4s); ops != "" {
		diffs = append(diffs, ops)
	}
	if astruct.Floatfield != astructOther.Floatfield {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Floatfield"))
	}
	if astruct.Intfield != astructOther.Intfield {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Intfield"))
	}
	if astruct.Anotherbooleanfield != astructOther.Anotherbooleanfield {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Anotherbooleanfield"))
	}
	if astruct.Duration1 != astructOther.Duration1 {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Duration1"))
	}
	if ops := __gong__diffSliceOfPointers(stage, astruct, "Anarrayofa", astructOther.Anarrayofa, astruct.Anarrayofa); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, astruct, "Anotherarrayofb", astructOther.Anotherarrayofb, astruct.Anotherarrayofb); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, astruct, "AnarrayofbUse", astructOther.AnarrayofbUse, astruct.AnarrayofbUse); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, astruct, "Anarrayofb2Use", astructOther.Anarrayofb2Use, astruct.Anarrayofb2Use); ops != "" {
		diffs = append(diffs, ops)
	}
	if astruct.AnAstruct != astructOther.AnAstruct {
		diffs = append(diffs, astruct.GongMarshallField(stage, "AnAstruct"))
	}
	if astruct.TextFieldBespokeSize != astructOther.TextFieldBespokeSize {
		diffs = append(diffs, astruct.GongMarshallField(stage, "TextFieldBespokeSize"))
	}
	if astruct.TextArea != astructOther.TextArea {
		diffs = append(diffs, astruct.GongMarshallField(stage, "TextArea"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (astructbstruct2use *AstructBstruct2Use) GongDiff(stage *Stage, astructbstruct2useOther *AstructBstruct2Use) (diffs []string) {
	// insertion point for field diffs
	if astructbstruct2use.Name != astructbstruct2useOther.Name {
		diffs = append(diffs, astructbstruct2use.GongMarshallField(stage, "Name"))
	}
	if astructbstruct2use.Bstrcut2 != astructbstruct2useOther.Bstrcut2 {
		diffs = append(diffs, astructbstruct2use.GongMarshallField(stage, "Bstrcut2"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (astructbstructuse *AstructBstructUse) GongDiff(stage *Stage, astructbstructuseOther *AstructBstructUse) (diffs []string) {
	// insertion point for field diffs
	if astructbstructuse.Name != astructbstructuseOther.Name {
		diffs = append(diffs, astructbstructuse.GongMarshallField(stage, "Name"))
	}
	if astructbstructuse.Bstruct2 != astructbstructuseOther.Bstruct2 {
		diffs = append(diffs, astructbstructuse.GongMarshallField(stage, "Bstruct2"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (bstruct *Bstruct) GongDiff(stage *Stage, bstructOther *Bstruct) (diffs []string) {
	// insertion point for field diffs
	if bstruct.Name != bstructOther.Name {
		diffs = append(diffs, bstruct.GongMarshallField(stage, "Name"))
	}
	if bstruct.Floatfield != bstructOther.Floatfield {
		diffs = append(diffs, bstruct.GongMarshallField(stage, "Floatfield"))
	}
	if bstruct.Floatfield2 != bstructOther.Floatfield2 {
		diffs = append(diffs, bstruct.GongMarshallField(stage, "Floatfield2"))
	}
	if bstruct.Intfield != bstructOther.Intfield {
		diffs = append(diffs, bstruct.GongMarshallField(stage, "Intfield"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dstruct *Dstruct) GongDiff(stage *Stage, dstructOther *Dstruct) (diffs []string) {
	// insertion point for field diffs
	if dstruct.Name != dstructOther.Name {
		diffs = append(diffs, dstruct.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, dstruct, "Anarrayofb", dstructOther.Anarrayofb, dstruct.Anarrayofb); ops != "" {
		diffs = append(diffs, ops)
	}
	if dstruct.Gstruct != dstructOther.Gstruct {
		diffs = append(diffs, dstruct.GongMarshallField(stage, "Gstruct"))
	}
	if ops := __gong__diffSliceOfPointers(stage, dstruct, "Gstructs", dstructOther.Gstructs, dstruct.Gstructs); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongDiff(stage *Stage, f0123456789012345678901234567890Other *F0123456789012345678901234567890) (diffs []string) {
	// insertion point for field diffs
	if f0123456789012345678901234567890.Name != f0123456789012345678901234567890Other.Name {
		diffs = append(diffs, f0123456789012345678901234567890.GongMarshallField(stage, "Name"))
	}
	if f0123456789012345678901234567890.Date != f0123456789012345678901234567890Other.Date {
		diffs = append(diffs, f0123456789012345678901234567890.GongMarshallField(stage, "Date"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gstruct *Gstruct) GongDiff(stage *Stage, gstructOther *Gstruct) (diffs []string) {
	// insertion point for field diffs
	if gstruct.Name != gstructOther.Name {
		diffs = append(diffs, gstruct.GongMarshallField(stage, "Name"))
	}
	if gstruct.Floatfield != gstructOther.Floatfield {
		diffs = append(diffs, gstruct.GongMarshallField(stage, "Floatfield"))
	}
	if gstruct.Floatfield2 != gstructOther.Floatfield2 {
		diffs = append(diffs, gstruct.GongMarshallField(stage, "Floatfield2"))
	}
	if gstruct.Intfield != gstructOther.Intfield {
		diffs = append(diffs, gstruct.GongMarshallField(stage, "Intfield"))
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
