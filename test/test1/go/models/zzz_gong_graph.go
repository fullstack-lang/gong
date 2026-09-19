// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (astruct *Astruct) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Astructs[astruct]

	return
}

func (stage *Stage) IsStagedAstruct(astruct *Astruct) (ok bool) {

	return astruct.GongIsStaged(stage)
}

func (astructbstruct2use *AstructBstruct2Use) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AstructBstruct2Uses[astructbstruct2use]

	return
}

func (stage *Stage) IsStagedAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) (ok bool) {

	return astructbstruct2use.GongIsStaged(stage)
}

func (astructbstructuse *AstructBstructUse) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AstructBstructUses[astructbstructuse]

	return
}

func (stage *Stage) IsStagedAstructBstructUse(astructbstructuse *AstructBstructUse) (ok bool) {

	return astructbstructuse.GongIsStaged(stage)
}

func (bstruct *Bstruct) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Bstructs[bstruct]

	return
}

func (stage *Stage) IsStagedBstruct(bstruct *Bstruct) (ok bool) {

	return bstruct.GongIsStaged(stage)
}

func (dstruct *Dstruct) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Dstructs[dstruct]

	return
}

func (stage *Stage) IsStagedDstruct(dstruct *Dstruct) (ok bool) {

	return dstruct.GongIsStaged(stage)
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.F0123456789012345678901234567890s[f0123456789012345678901234567890]

	return
}

func (stage *Stage) IsStagedF0123456789012345678901234567890(f0123456789012345678901234567890 *F0123456789012345678901234567890) (ok bool) {

	return f0123456789012345678901234567890.GongIsStaged(stage)
}

func (gstruct *Gstruct) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Gstructs[gstruct]

	return
}

func (stage *Stage) IsStagedGstruct(gstruct *Gstruct) (ok bool) {

	return gstruct.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (astruct *Astruct) GongStageBranch(stage *Stage) {
	stage.StageBranchAstruct(astruct)
}

func (stage *Stage) StageBranchAstruct(astruct *Astruct) {

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
	stage.StageBranchAstructBstruct2Use(astructbstruct2use)
}

func (stage *Stage) StageBranchAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) {

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
	stage.StageBranchAstructBstructUse(astructbstructuse)
}

func (stage *Stage) StageBranchAstructBstructUse(astructbstructuse *AstructBstructUse) {

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
	stage.StageBranchBstruct(bstruct)
}

func (stage *Stage) StageBranchBstruct(bstruct *Bstruct) {

	// check if instance is already staged
	if stage.IsStaged(bstruct) {
		return
	}

	bstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dstruct *Dstruct) GongStageBranch(stage *Stage) {
	stage.StageBranchDstruct(dstruct)
}

func (stage *Stage) StageBranchDstruct(dstruct *Dstruct) {

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
	stage.StageBranchF0123456789012345678901234567890(f0123456789012345678901234567890)
}

func (stage *Stage) StageBranchF0123456789012345678901234567890(f0123456789012345678901234567890 *F0123456789012345678901234567890) {

	// check if instance is already staged
	if stage.IsStaged(f0123456789012345678901234567890) {
		return
	}

	f0123456789012345678901234567890.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gstruct *Gstruct) GongStageBranch(stage *Stage) {
	stage.StageBranchGstruct(gstruct)
}

func (stage *Stage) StageBranchGstruct(gstruct *Gstruct) {

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

	// astructFrom has already been copied
	if _astructTo, ok := mapOrigCopy[astructFrom]; ok {
		astructTo = _astructTo.(*Astruct)
		return
	}

	astructTo = new(Astruct)
	mapOrigCopy[astructFrom] = astructTo
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

	// astructbstruct2useFrom has already been copied
	if _astructbstruct2useTo, ok := mapOrigCopy[astructbstruct2useFrom]; ok {
		astructbstruct2useTo = _astructbstruct2useTo.(*AstructBstruct2Use)
		return
	}

	astructbstruct2useTo = new(AstructBstruct2Use)
	mapOrigCopy[astructbstruct2useFrom] = astructbstruct2useTo
	astructbstruct2useFrom.GongCopyBasicFields(astructbstruct2useTo)

	//insertion point for the staging of instances referenced by pointers
	if astructbstruct2useFrom.Bstrcut2 != nil {
		astructbstruct2useTo.Bstrcut2 = GongCopyBranchBstruct(mapOrigCopy, astructbstruct2useFrom.Bstrcut2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAstructBstructUse(mapOrigCopy map[any]any, astructbstructuseFrom *AstructBstructUse) (astructbstructuseTo *AstructBstructUse) {

	// astructbstructuseFrom has already been copied
	if _astructbstructuseTo, ok := mapOrigCopy[astructbstructuseFrom]; ok {
		astructbstructuseTo = _astructbstructuseTo.(*AstructBstructUse)
		return
	}

	astructbstructuseTo = new(AstructBstructUse)
	mapOrigCopy[astructbstructuseFrom] = astructbstructuseTo
	astructbstructuseFrom.GongCopyBasicFields(astructbstructuseTo)

	//insertion point for the staging of instances referenced by pointers
	if astructbstructuseFrom.Bstruct2 != nil {
		astructbstructuseTo.Bstruct2 = GongCopyBranchBstruct(mapOrigCopy, astructbstructuseFrom.Bstruct2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchBstruct(mapOrigCopy map[any]any, bstructFrom *Bstruct) (bstructTo *Bstruct) {

	// bstructFrom has already been copied
	if _bstructTo, ok := mapOrigCopy[bstructFrom]; ok {
		bstructTo = _bstructTo.(*Bstruct)
		return
	}

	bstructTo = new(Bstruct)
	mapOrigCopy[bstructFrom] = bstructTo
	bstructFrom.GongCopyBasicFields(bstructTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDstruct(mapOrigCopy map[any]any, dstructFrom *Dstruct) (dstructTo *Dstruct) {

	// dstructFrom has already been copied
	if _dstructTo, ok := mapOrigCopy[dstructFrom]; ok {
		dstructTo = _dstructTo.(*Dstruct)
		return
	}

	dstructTo = new(Dstruct)
	mapOrigCopy[dstructFrom] = dstructTo
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

	// f0123456789012345678901234567890From has already been copied
	if _f0123456789012345678901234567890To, ok := mapOrigCopy[f0123456789012345678901234567890From]; ok {
		f0123456789012345678901234567890To = _f0123456789012345678901234567890To.(*F0123456789012345678901234567890)
		return
	}

	f0123456789012345678901234567890To = new(F0123456789012345678901234567890)
	mapOrigCopy[f0123456789012345678901234567890From] = f0123456789012345678901234567890To
	f0123456789012345678901234567890From.GongCopyBasicFields(f0123456789012345678901234567890To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGstruct(mapOrigCopy map[any]any, gstructFrom *Gstruct) (gstructTo *Gstruct) {

	// gstructFrom has already been copied
	if _gstructTo, ok := mapOrigCopy[gstructFrom]; ok {
		gstructTo = _gstructTo.(*Gstruct)
		return
	}

	gstructTo = new(Gstruct)
	mapOrigCopy[gstructFrom] = gstructTo
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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (astruct *Astruct) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAstruct(astruct)
}

func (stage *Stage) UnstageBranchAstruct(astruct *Astruct) {

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
	stage.UnstageBranchAstructBstruct2Use(astructbstruct2use)
}

func (stage *Stage) UnstageBranchAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) {

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
	stage.UnstageBranchAstructBstructUse(astructbstructuse)
}

func (stage *Stage) UnstageBranchAstructBstructUse(astructbstructuse *AstructBstructUse) {

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
	stage.UnstageBranchBstruct(bstruct)
}

func (stage *Stage) UnstageBranchBstruct(bstruct *Bstruct) {

	// check if instance is already staged
	if !stage.IsStaged(bstruct) {
		return
	}

	bstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dstruct *Dstruct) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDstruct(dstruct)
}

func (stage *Stage) UnstageBranchDstruct(dstruct *Dstruct) {

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
	stage.UnstageBranchF0123456789012345678901234567890(f0123456789012345678901234567890)
}

func (stage *Stage) UnstageBranchF0123456789012345678901234567890(f0123456789012345678901234567890 *F0123456789012345678901234567890) {

	// check if instance is already staged
	if !stage.IsStaged(f0123456789012345678901234567890) {
		return
	}

	f0123456789012345678901234567890.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gstruct *Gstruct) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGstruct(gstruct)
}

func (stage *Stage) UnstageBranchGstruct(gstruct *Gstruct) {

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
	if instance.Associationtob != nil {
		reference.Associationtob = stage.Bstructs_reference[instance.Associationtob]
	}
	if instance.Anotherassociationtob_2 != nil {
		reference.Anotherassociationtob_2 = stage.Bstructs_reference[instance.Anotherassociationtob_2]
	}
	if instance.Bstruct != nil {
		reference.Bstruct = stage.Bstructs_reference[instance.Bstruct]
	}
	if instance.Bstruct2 != nil {
		reference.Bstruct2 = stage.Bstructs_reference[instance.Bstruct2]
	}
	if instance.Dstruct != nil {
		reference.Dstruct = stage.Dstructs_reference[instance.Dstruct]
	}
	if instance.Dstruct2 != nil {
		reference.Dstruct2 = stage.Dstructs_reference[instance.Dstruct2]
	}
	if instance.Dstruct3 != nil {
		reference.Dstruct3 = stage.Dstructs_reference[instance.Dstruct3]
	}
	if instance.Dstruct4 != nil {
		reference.Dstruct4 = stage.Dstructs_reference[instance.Dstruct4]
	}
	if instance.AnAstruct != nil {
		reference.AnAstruct = stage.Astructs_reference[instance.AnAstruct]
	}
	// insertion point for slice of pointers field
	reference.Anarrayofb = reference.Anarrayofb[:0]
	for _, _b := range instance.Anarrayofb {
		reference.Anarrayofb = append(reference.Anarrayofb, stage.Bstructs_reference[_b])
	}
	reference.Dstruct4s = reference.Dstruct4s[:0]
	for _, _b := range instance.Dstruct4s {
		reference.Dstruct4s = append(reference.Dstruct4s, stage.Dstructs_reference[_b])
	}
	reference.Anarrayofa = reference.Anarrayofa[:0]
	for _, _b := range instance.Anarrayofa {
		reference.Anarrayofa = append(reference.Anarrayofa, stage.Astructs_reference[_b])
	}
	reference.Anotherarrayofb = reference.Anotherarrayofb[:0]
	for _, _b := range instance.Anotherarrayofb {
		reference.Anotherarrayofb = append(reference.Anotherarrayofb, stage.Bstructs_reference[_b])
	}
	reference.AnarrayofbUse = reference.AnarrayofbUse[:0]
	for _, _b := range instance.AnarrayofbUse {
		reference.AnarrayofbUse = append(reference.AnarrayofbUse, stage.AstructBstructUses_reference[_b])
	}
	reference.Anarrayofb2Use = reference.Anarrayofb2Use[:0]
	for _, _b := range instance.Anarrayofb2Use {
		reference.Anarrayofb2Use = append(reference.Anarrayofb2Use, stage.AstructBstruct2Uses_reference[_b])
	}
}

func (reference *AstructBstruct2Use) GongReconstructPointersFromReferences(stage *Stage, instance *AstructBstruct2Use) {
	// insertion point for pointers field
	if instance.Bstrcut2 != nil {
		reference.Bstrcut2 = stage.Bstructs_reference[instance.Bstrcut2]
	}
	// insertion point for slice of pointers field
}

func (reference *AstructBstructUse) GongReconstructPointersFromReferences(stage *Stage, instance *AstructBstructUse) {
	// insertion point for pointers field
	if instance.Bstruct2 != nil {
		reference.Bstruct2 = stage.Bstructs_reference[instance.Bstruct2]
	}
	// insertion point for slice of pointers field
}

func (reference *Bstruct) GongReconstructPointersFromReferences(stage *Stage, instance *Bstruct) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Dstruct) GongReconstructPointersFromReferences(stage *Stage, instance *Dstruct) {
	// insertion point for pointers field
	if instance.Gstruct != nil {
		reference.Gstruct = stage.Gstructs_reference[instance.Gstruct]
	}
	// insertion point for slice of pointers field
	reference.Anarrayofb = reference.Anarrayofb[:0]
	for _, _b := range instance.Anarrayofb {
		reference.Anarrayofb = append(reference.Anarrayofb, stage.Bstructs_reference[_b])
	}
	reference.Gstructs = reference.Gstructs[:0]
	for _, _b := range instance.Gstructs {
		reference.Gstructs = append(reference.Gstructs, stage.Gstructs_reference[_b])
	}
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
	if _reference := reference.Associationtob; _reference != nil {
		reference.Associationtob = nil
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			reference.Associationtob = _instance
		}
	}
	if _reference := reference.Anotherassociationtob_2; _reference != nil {
		reference.Anotherassociationtob_2 = nil
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			reference.Anotherassociationtob_2 = _instance
		}
	}
	if _reference := reference.Bstruct; _reference != nil {
		reference.Bstruct = nil
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			reference.Bstruct = _instance
		}
	}
	if _reference := reference.Bstruct2; _reference != nil {
		reference.Bstruct2 = nil
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			reference.Bstruct2 = _instance
		}
	}
	if _reference := reference.Dstruct; _reference != nil {
		reference.Dstruct = nil
		if _instance, ok := stage.Dstructs_instance[_reference]; ok {
			reference.Dstruct = _instance
		}
	}
	if _reference := reference.Dstruct2; _reference != nil {
		reference.Dstruct2 = nil
		if _instance, ok := stage.Dstructs_instance[_reference]; ok {
			reference.Dstruct2 = _instance
		}
	}
	if _reference := reference.Dstruct3; _reference != nil {
		reference.Dstruct3 = nil
		if _instance, ok := stage.Dstructs_instance[_reference]; ok {
			reference.Dstruct3 = _instance
		}
	}
	if _reference := reference.Dstruct4; _reference != nil {
		reference.Dstruct4 = nil
		if _instance, ok := stage.Dstructs_instance[_reference]; ok {
			reference.Dstruct4 = _instance
		}
	}
	if _reference := reference.AnAstruct; _reference != nil {
		reference.AnAstruct = nil
		if _instance, ok := stage.Astructs_instance[_reference]; ok {
			reference.AnAstruct = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Anarrayofb []*Bstruct
	for _, _reference := range reference.Anarrayofb {
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			_Anarrayofb = append(_Anarrayofb, _instance)
		}
	}
	reference.Anarrayofb = _Anarrayofb
	var _Dstruct4s []*Dstruct
	for _, _reference := range reference.Dstruct4s {
		if _instance, ok := stage.Dstructs_instance[_reference]; ok {
			_Dstruct4s = append(_Dstruct4s, _instance)
		}
	}
	reference.Dstruct4s = _Dstruct4s
	var _Anarrayofa []*Astruct
	for _, _reference := range reference.Anarrayofa {
		if _instance, ok := stage.Astructs_instance[_reference]; ok {
			_Anarrayofa = append(_Anarrayofa, _instance)
		}
	}
	reference.Anarrayofa = _Anarrayofa
	var _Anotherarrayofb []*Bstruct
	for _, _reference := range reference.Anotherarrayofb {
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			_Anotherarrayofb = append(_Anotherarrayofb, _instance)
		}
	}
	reference.Anotherarrayofb = _Anotherarrayofb
	var _AnarrayofbUse []*AstructBstructUse
	for _, _reference := range reference.AnarrayofbUse {
		if _instance, ok := stage.AstructBstructUses_instance[_reference]; ok {
			_AnarrayofbUse = append(_AnarrayofbUse, _instance)
		}
	}
	reference.AnarrayofbUse = _AnarrayofbUse
	var _Anarrayofb2Use []*AstructBstruct2Use
	for _, _reference := range reference.Anarrayofb2Use {
		if _instance, ok := stage.AstructBstruct2Uses_instance[_reference]; ok {
			_Anarrayofb2Use = append(_Anarrayofb2Use, _instance)
		}
	}
	reference.Anarrayofb2Use = _Anarrayofb2Use
}

func (reference *AstructBstruct2Use) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Bstrcut2; _reference != nil {
		reference.Bstrcut2 = nil
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			reference.Bstrcut2 = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *AstructBstructUse) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Bstruct2; _reference != nil {
		reference.Bstruct2 = nil
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			reference.Bstruct2 = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Bstruct) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Dstruct) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Gstruct; _reference != nil {
		reference.Gstruct = nil
		if _instance, ok := stage.Gstructs_instance[_reference]; ok {
			reference.Gstruct = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Anarrayofb []*Bstruct
	for _, _reference := range reference.Anarrayofb {
		if _instance, ok := stage.Bstructs_instance[_reference]; ok {
			_Anarrayofb = append(_Anarrayofb, _instance)
		}
	}
	reference.Anarrayofb = _Anarrayofb
	var _Gstructs []*Gstruct
	for _, _reference := range reference.Gstructs {
		if _instance, ok := stage.Gstructs_instance[_reference]; ok {
			_Gstructs = append(_Gstructs, _instance)
		}
	}
	reference.Gstructs = _Gstructs
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
	if (astruct.Associationtob == nil) != (astructOther.Associationtob == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Associationtob"))
	} else if astruct.Associationtob != nil && astructOther.Associationtob != nil {
		if astruct.Associationtob != astructOther.Associationtob {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Associationtob"))
		}
	}
	AnarrayofbDifferent := false
	if len(astruct.Anarrayofb) != len(astructOther.Anarrayofb) {
		AnarrayofbDifferent = true
	} else {
		for i := range astruct.Anarrayofb {
			if (astruct.Anarrayofb[i] == nil) != (astructOther.Anarrayofb[i] == nil) {
				AnarrayofbDifferent = true
				break
			} else if astruct.Anarrayofb[i] != nil && astructOther.Anarrayofb[i] != nil {
				// this is a pointer comparaison
				if astruct.Anarrayofb[i] != astructOther.Anarrayofb[i] {
					AnarrayofbDifferent = true
					break
				}
			}
		}
	}
	if AnarrayofbDifferent {
		ops := stage.Diff(
			astruct,
			"Anarrayofb",
			len(astructOther.Anarrayofb),
			len(astruct.Anarrayofb),
			func(i, j int) bool {
				return astructOther.Anarrayofb[i] == astruct.Anarrayofb[j]
			},
			func(j int) string {
				return astruct.Anarrayofb[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if (astruct.Anotherassociationtob_2 == nil) != (astructOther.Anotherassociationtob_2 == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Anotherassociationtob_2"))
	} else if astruct.Anotherassociationtob_2 != nil && astructOther.Anotherassociationtob_2 != nil {
		if astruct.Anotherassociationtob_2 != astructOther.Anotherassociationtob_2 {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Anotherassociationtob_2"))
		}
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
	if (astruct.Bstruct == nil) != (astructOther.Bstruct == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Bstruct"))
	} else if astruct.Bstruct != nil && astructOther.Bstruct != nil {
		if astruct.Bstruct != astructOther.Bstruct {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Bstruct"))
		}
	}
	if (astruct.Bstruct2 == nil) != (astructOther.Bstruct2 == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Bstruct2"))
	} else if astruct.Bstruct2 != nil && astructOther.Bstruct2 != nil {
		if astruct.Bstruct2 != astructOther.Bstruct2 {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Bstruct2"))
		}
	}
	if (astruct.Dstruct == nil) != (astructOther.Dstruct == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct"))
	} else if astruct.Dstruct != nil && astructOther.Dstruct != nil {
		if astruct.Dstruct != astructOther.Dstruct {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct"))
		}
	}
	if (astruct.Dstruct2 == nil) != (astructOther.Dstruct2 == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct2"))
	} else if astruct.Dstruct2 != nil && astructOther.Dstruct2 != nil {
		if astruct.Dstruct2 != astructOther.Dstruct2 {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct2"))
		}
	}
	if (astruct.Dstruct3 == nil) != (astructOther.Dstruct3 == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct3"))
	} else if astruct.Dstruct3 != nil && astructOther.Dstruct3 != nil {
		if astruct.Dstruct3 != astructOther.Dstruct3 {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct3"))
		}
	}
	if (astruct.Dstruct4 == nil) != (astructOther.Dstruct4 == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct4"))
	} else if astruct.Dstruct4 != nil && astructOther.Dstruct4 != nil {
		if astruct.Dstruct4 != astructOther.Dstruct4 {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct4"))
		}
	}
	Dstruct4sDifferent := false
	if len(astruct.Dstruct4s) != len(astructOther.Dstruct4s) {
		Dstruct4sDifferent = true
	} else {
		for i := range astruct.Dstruct4s {
			if (astruct.Dstruct4s[i] == nil) != (astructOther.Dstruct4s[i] == nil) {
				Dstruct4sDifferent = true
				break
			} else if astruct.Dstruct4s[i] != nil && astructOther.Dstruct4s[i] != nil {
				// this is a pointer comparaison
				if astruct.Dstruct4s[i] != astructOther.Dstruct4s[i] {
					Dstruct4sDifferent = true
					break
				}
			}
		}
	}
	if Dstruct4sDifferent {
		ops := stage.Diff(
			astruct,
			"Dstruct4s",
			len(astructOther.Dstruct4s),
			len(astruct.Dstruct4s),
			func(i, j int) bool {
				return astructOther.Dstruct4s[i] == astruct.Dstruct4s[j]
			},
			func(j int) string {
				return astruct.Dstruct4s[j].GongGetIdentifier(stage)
			},
		)
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
	AnarrayofaDifferent := false
	if len(astruct.Anarrayofa) != len(astructOther.Anarrayofa) {
		AnarrayofaDifferent = true
	} else {
		for i := range astruct.Anarrayofa {
			if (astruct.Anarrayofa[i] == nil) != (astructOther.Anarrayofa[i] == nil) {
				AnarrayofaDifferent = true
				break
			} else if astruct.Anarrayofa[i] != nil && astructOther.Anarrayofa[i] != nil {
				// this is a pointer comparaison
				if astruct.Anarrayofa[i] != astructOther.Anarrayofa[i] {
					AnarrayofaDifferent = true
					break
				}
			}
		}
	}
	if AnarrayofaDifferent {
		ops := stage.Diff(
			astruct,
			"Anarrayofa",
			len(astructOther.Anarrayofa),
			len(astruct.Anarrayofa),
			func(i, j int) bool {
				return astructOther.Anarrayofa[i] == astruct.Anarrayofa[j]
			},
			func(j int) string {
				return astruct.Anarrayofa[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AnotherarrayofbDifferent := false
	if len(astruct.Anotherarrayofb) != len(astructOther.Anotherarrayofb) {
		AnotherarrayofbDifferent = true
	} else {
		for i := range astruct.Anotherarrayofb {
			if (astruct.Anotherarrayofb[i] == nil) != (astructOther.Anotherarrayofb[i] == nil) {
				AnotherarrayofbDifferent = true
				break
			} else if astruct.Anotherarrayofb[i] != nil && astructOther.Anotherarrayofb[i] != nil {
				// this is a pointer comparaison
				if astruct.Anotherarrayofb[i] != astructOther.Anotherarrayofb[i] {
					AnotherarrayofbDifferent = true
					break
				}
			}
		}
	}
	if AnotherarrayofbDifferent {
		ops := stage.Diff(
			astruct,
			"Anotherarrayofb",
			len(astructOther.Anotherarrayofb),
			len(astruct.Anotherarrayofb),
			func(i, j int) bool {
				return astructOther.Anotherarrayofb[i] == astruct.Anotherarrayofb[j]
			},
			func(j int) string {
				return astruct.Anotherarrayofb[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AnarrayofbUseDifferent := false
	if len(astruct.AnarrayofbUse) != len(astructOther.AnarrayofbUse) {
		AnarrayofbUseDifferent = true
	} else {
		for i := range astruct.AnarrayofbUse {
			if (astruct.AnarrayofbUse[i] == nil) != (astructOther.AnarrayofbUse[i] == nil) {
				AnarrayofbUseDifferent = true
				break
			} else if astruct.AnarrayofbUse[i] != nil && astructOther.AnarrayofbUse[i] != nil {
				// this is a pointer comparaison
				if astruct.AnarrayofbUse[i] != astructOther.AnarrayofbUse[i] {
					AnarrayofbUseDifferent = true
					break
				}
			}
		}
	}
	if AnarrayofbUseDifferent {
		ops := stage.Diff(
			astruct,
			"AnarrayofbUse",
			len(astructOther.AnarrayofbUse),
			len(astruct.AnarrayofbUse),
			func(i, j int) bool {
				return astructOther.AnarrayofbUse[i] == astruct.AnarrayofbUse[j]
			},
			func(j int) string {
				return astruct.AnarrayofbUse[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Anarrayofb2UseDifferent := false
	if len(astruct.Anarrayofb2Use) != len(astructOther.Anarrayofb2Use) {
		Anarrayofb2UseDifferent = true
	} else {
		for i := range astruct.Anarrayofb2Use {
			if (astruct.Anarrayofb2Use[i] == nil) != (astructOther.Anarrayofb2Use[i] == nil) {
				Anarrayofb2UseDifferent = true
				break
			} else if astruct.Anarrayofb2Use[i] != nil && astructOther.Anarrayofb2Use[i] != nil {
				// this is a pointer comparaison
				if astruct.Anarrayofb2Use[i] != astructOther.Anarrayofb2Use[i] {
					Anarrayofb2UseDifferent = true
					break
				}
			}
		}
	}
	if Anarrayofb2UseDifferent {
		ops := stage.Diff(
			astruct,
			"Anarrayofb2Use",
			len(astructOther.Anarrayofb2Use),
			len(astruct.Anarrayofb2Use),
			func(i, j int) bool {
				return astructOther.Anarrayofb2Use[i] == astruct.Anarrayofb2Use[j]
			},
			func(j int) string {
				return astruct.Anarrayofb2Use[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if (astruct.AnAstruct == nil) != (astructOther.AnAstruct == nil) {
		diffs = append(diffs, astruct.GongMarshallField(stage, "AnAstruct"))
	} else if astruct.AnAstruct != nil && astructOther.AnAstruct != nil {
		if astruct.AnAstruct != astructOther.AnAstruct {
			diffs = append(diffs, astruct.GongMarshallField(stage, "AnAstruct"))
		}
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
	if (astructbstruct2use.Bstrcut2 == nil) != (astructbstruct2useOther.Bstrcut2 == nil) {
		diffs = append(diffs, astructbstruct2use.GongMarshallField(stage, "Bstrcut2"))
	} else if astructbstruct2use.Bstrcut2 != nil && astructbstruct2useOther.Bstrcut2 != nil {
		if astructbstruct2use.Bstrcut2 != astructbstruct2useOther.Bstrcut2 {
			diffs = append(diffs, astructbstruct2use.GongMarshallField(stage, "Bstrcut2"))
		}
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
	if (astructbstructuse.Bstruct2 == nil) != (astructbstructuseOther.Bstruct2 == nil) {
		diffs = append(diffs, astructbstructuse.GongMarshallField(stage, "Bstruct2"))
	} else if astructbstructuse.Bstruct2 != nil && astructbstructuseOther.Bstruct2 != nil {
		if astructbstructuse.Bstruct2 != astructbstructuseOther.Bstruct2 {
			diffs = append(diffs, astructbstructuse.GongMarshallField(stage, "Bstruct2"))
		}
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
	AnarrayofbDifferent := false
	if len(dstruct.Anarrayofb) != len(dstructOther.Anarrayofb) {
		AnarrayofbDifferent = true
	} else {
		for i := range dstruct.Anarrayofb {
			if (dstruct.Anarrayofb[i] == nil) != (dstructOther.Anarrayofb[i] == nil) {
				AnarrayofbDifferent = true
				break
			} else if dstruct.Anarrayofb[i] != nil && dstructOther.Anarrayofb[i] != nil {
				// this is a pointer comparaison
				if dstruct.Anarrayofb[i] != dstructOther.Anarrayofb[i] {
					AnarrayofbDifferent = true
					break
				}
			}
		}
	}
	if AnarrayofbDifferent {
		ops := stage.Diff(
			dstruct,
			"Anarrayofb",
			len(dstructOther.Anarrayofb),
			len(dstruct.Anarrayofb),
			func(i, j int) bool {
				return dstructOther.Anarrayofb[i] == dstruct.Anarrayofb[j]
			},
			func(j int) string {
				return dstruct.Anarrayofb[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if (dstruct.Gstruct == nil) != (dstructOther.Gstruct == nil) {
		diffs = append(diffs, dstruct.GongMarshallField(stage, "Gstruct"))
	} else if dstruct.Gstruct != nil && dstructOther.Gstruct != nil {
		if dstruct.Gstruct != dstructOther.Gstruct {
			diffs = append(diffs, dstruct.GongMarshallField(stage, "Gstruct"))
		}
	}
	GstructsDifferent := false
	if len(dstruct.Gstructs) != len(dstructOther.Gstructs) {
		GstructsDifferent = true
	} else {
		for i := range dstruct.Gstructs {
			if (dstruct.Gstructs[i] == nil) != (dstructOther.Gstructs[i] == nil) {
				GstructsDifferent = true
				break
			} else if dstruct.Gstructs[i] != nil && dstructOther.Gstructs[i] != nil {
				// this is a pointer comparaison
				if dstruct.Gstructs[i] != dstructOther.Gstructs[i] {
					GstructsDifferent = true
					break
				}
			}
		}
	}
	if GstructsDifferent {
		ops := stage.Diff(
			dstruct,
			"Gstructs",
			len(dstructOther.Gstructs),
			len(dstruct.Gstructs),
			func(i, j int) bool {
				return dstructOther.Gstructs[i] == dstruct.Gstructs[j]
			},
			func(j int) string {
				return dstruct.Gstructs[j].GongGetIdentifier(stage)
			},
		)
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
