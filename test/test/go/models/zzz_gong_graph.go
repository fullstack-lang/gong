// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Astruct:
		ok = stage.IsStagedAstruct(target)

	case *AstructBstruct2Use:
		ok = stage.IsStagedAstructBstruct2Use(target)

	case *AstructBstructUse:
		ok = stage.IsStagedAstructBstructUse(target)

	case *Bstruct:
		ok = stage.IsStagedBstruct(target)

	case *Dstruct:
		ok = stage.IsStagedDstruct(target)

	case *F0123456789012345678901234567890:
		ok = stage.IsStagedF0123456789012345678901234567890(target)

	case *Gstruct:
		ok = stage.IsStagedGstruct(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Astruct:
		ok = stage.IsStagedAstruct(target)

	case *AstructBstruct2Use:
		ok = stage.IsStagedAstructBstruct2Use(target)

	case *AstructBstructUse:
		ok = stage.IsStagedAstructBstructUse(target)

	case *Bstruct:
		ok = stage.IsStagedBstruct(target)

	case *Dstruct:
		ok = stage.IsStagedDstruct(target)

	case *F0123456789012345678901234567890:
		ok = stage.IsStagedF0123456789012345678901234567890(target)

	case *Gstruct:
		ok = stage.IsStagedGstruct(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAstruct(astruct *Astruct) (ok bool) {

	_, ok = stage.Astructs[astruct]

	return
}

func (stage *Stage) IsStagedAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) (ok bool) {

	_, ok = stage.AstructBstruct2Uses[astructbstruct2use]

	return
}

func (stage *Stage) IsStagedAstructBstructUse(astructbstructuse *AstructBstructUse) (ok bool) {

	_, ok = stage.AstructBstructUses[astructbstructuse]

	return
}

func (stage *Stage) IsStagedBstruct(bstruct *Bstruct) (ok bool) {

	_, ok = stage.Bstructs[bstruct]

	return
}

func (stage *Stage) IsStagedDstruct(dstruct *Dstruct) (ok bool) {

	_, ok = stage.Dstructs[dstruct]

	return
}

func (stage *Stage) IsStagedF0123456789012345678901234567890(f0123456789012345678901234567890 *F0123456789012345678901234567890) (ok bool) {

	_, ok = stage.F0123456789012345678901234567890s[f0123456789012345678901234567890]

	return
}

func (stage *Stage) IsStagedGstruct(gstruct *Gstruct) (ok bool) {

	_, ok = stage.Gstructs[gstruct]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Astruct:
		stage.StageBranchAstruct(target)

	case *AstructBstruct2Use:
		stage.StageBranchAstructBstruct2Use(target)

	case *AstructBstructUse:
		stage.StageBranchAstructBstructUse(target)

	case *Bstruct:
		stage.StageBranchBstruct(target)

	case *Dstruct:
		stage.StageBranchDstruct(target)

	case *F0123456789012345678901234567890:
		stage.StageBranchF0123456789012345678901234567890(target)

	case *Gstruct:
		stage.StageBranchGstruct(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAstruct(astruct *Astruct) {

	// check if instance is already staged
	if IsStaged(stage, astruct) {
		return
	}

	astruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astruct.Associationtob != nil {
		StageBranch(stage, astruct.Associationtob)
	}
	if astruct.Anotherassociationtob_2 != nil {
		StageBranch(stage, astruct.Anotherassociationtob_2)
	}
	if astruct.Bstruct != nil {
		StageBranch(stage, astruct.Bstruct)
	}
	if astruct.Bstruct2 != nil {
		StageBranch(stage, astruct.Bstruct2)
	}
	if astruct.Dstruct != nil {
		StageBranch(stage, astruct.Dstruct)
	}
	if astruct.Dstruct2 != nil {
		StageBranch(stage, astruct.Dstruct2)
	}
	if astruct.Dstruct3 != nil {
		StageBranch(stage, astruct.Dstruct3)
	}
	if astruct.Dstruct4 != nil {
		StageBranch(stage, astruct.Dstruct4)
	}
	if astruct.AnAstruct != nil {
		StageBranch(stage, astruct.AnAstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range astruct.Anarrayofb {
		StageBranch(stage, _bstruct)
	}
	for _, _dstruct := range astruct.Dstruct4s {
		StageBranch(stage, _dstruct)
	}
	for _, _astruct := range astruct.Anarrayofa {
		StageBranch(stage, _astruct)
	}
	for _, _bstruct := range astruct.Anotherarrayofb {
		StageBranch(stage, _bstruct)
	}
	for _, _astructbstructuse := range astruct.AnarrayofbUse {
		StageBranch(stage, _astructbstructuse)
	}
	for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
		StageBranch(stage, _astructbstruct2use)
	}

}

func (stage *Stage) StageBranchAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) {

	// check if instance is already staged
	if IsStaged(stage, astructbstruct2use) {
		return
	}

	astructbstruct2use.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astructbstruct2use.Bstrcut2 != nil {
		StageBranch(stage, astructbstruct2use.Bstrcut2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchAstructBstructUse(astructbstructuse *AstructBstructUse) {

	// check if instance is already staged
	if IsStaged(stage, astructbstructuse) {
		return
	}

	astructbstructuse.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astructbstructuse.Bstruct2 != nil {
		StageBranch(stage, astructbstructuse.Bstruct2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchBstruct(bstruct *Bstruct) {

	// check if instance is already staged
	if IsStaged(stage, bstruct) {
		return
	}

	bstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDstruct(dstruct *Dstruct) {

	// check if instance is already staged
	if IsStaged(stage, dstruct) {
		return
	}

	dstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dstruct.Gstruct != nil {
		StageBranch(stage, dstruct.Gstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstruct.Anarrayofb {
		StageBranch(stage, _bstruct)
	}
	for _, _gstruct := range dstruct.Gstructs {
		StageBranch(stage, _gstruct)
	}

}

func (stage *Stage) StageBranchF0123456789012345678901234567890(f0123456789012345678901234567890 *F0123456789012345678901234567890) {

	// check if instance is already staged
	if IsStaged(stage, f0123456789012345678901234567890) {
		return
	}

	f0123456789012345678901234567890.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGstruct(gstruct *Gstruct) {

	// check if instance is already staged
	if IsStaged(stage, gstruct) {
		return
	}

	gstruct.Stage(stage)

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
	case *Astruct:
		toT := CopyBranchAstruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AstructBstruct2Use:
		toT := CopyBranchAstructBstruct2Use(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AstructBstructUse:
		toT := CopyBranchAstructBstructUse(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bstruct:
		toT := CopyBranchBstruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Dstruct:
		toT := CopyBranchDstruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *F0123456789012345678901234567890:
		toT := CopyBranchF0123456789012345678901234567890(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Gstruct:
		toT := CopyBranchGstruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAstruct(mapOrigCopy map[any]any, astructFrom *Astruct) (astructTo *Astruct) {

	// astructFrom has already been copied
	if _astructTo, ok := mapOrigCopy[astructFrom]; ok {
		astructTo = _astructTo.(*Astruct)
		return
	}

	astructTo = new(Astruct)
	mapOrigCopy[astructFrom] = astructTo
	astructFrom.CopyBasicFields(astructTo)

	//insertion point for the staging of instances referenced by pointers
	if astructFrom.Associationtob != nil {
		astructTo.Associationtob = CopyBranchBstruct(mapOrigCopy, astructFrom.Associationtob)
	}
	if astructFrom.Anotherassociationtob_2 != nil {
		astructTo.Anotherassociationtob_2 = astructFrom.Anotherassociationtob_2
	}
	if astructFrom.Bstruct != nil {
		astructTo.Bstruct = CopyBranchBstruct(mapOrigCopy, astructFrom.Bstruct)
	}
	if astructFrom.Bstruct2 != nil {
		astructTo.Bstruct2 = CopyBranchBstruct(mapOrigCopy, astructFrom.Bstruct2)
	}
	if astructFrom.Dstruct != nil {
		astructTo.Dstruct = CopyBranchDstruct(mapOrigCopy, astructFrom.Dstruct)
	}
	if astructFrom.Dstruct2 != nil {
		astructTo.Dstruct2 = CopyBranchDstruct(mapOrigCopy, astructFrom.Dstruct2)
	}
	if astructFrom.Dstruct3 != nil {
		astructTo.Dstruct3 = CopyBranchDstruct(mapOrigCopy, astructFrom.Dstruct3)
	}
	if astructFrom.Dstruct4 != nil {
		astructTo.Dstruct4 = CopyBranchDstruct(mapOrigCopy, astructFrom.Dstruct4)
	}
	if astructFrom.AnAstruct != nil {
		astructTo.AnAstruct = CopyBranchAstruct(mapOrigCopy, astructFrom.AnAstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range astructFrom.Anarrayofb {
		astructTo.Anarrayofb = append(astructTo.Anarrayofb, CopyBranchBstruct(mapOrigCopy, _bstruct))
	}
	for _, _dstruct := range astructFrom.Dstruct4s {
		astructTo.Dstruct4s = append(astructTo.Dstruct4s, CopyBranchDstruct(mapOrigCopy, _dstruct))
	}
	for _, _astruct := range astructFrom.Anarrayofa {
		astructTo.Anarrayofa = append(astructTo.Anarrayofa, CopyBranchAstruct(mapOrigCopy, _astruct))
	}
	for _, _bstruct := range astructFrom.Anotherarrayofb {
		astructTo.Anotherarrayofb = append(astructTo.Anotherarrayofb, CopyBranchBstruct(mapOrigCopy, _bstruct))
	}
	for _, _astructbstructuse := range astructFrom.AnarrayofbUse {
		astructTo.AnarrayofbUse = append(astructTo.AnarrayofbUse, CopyBranchAstructBstructUse(mapOrigCopy, _astructbstructuse))
	}
	for _, _astructbstruct2use := range astructFrom.Anarrayofb2Use {
		astructTo.Anarrayofb2Use = append(astructTo.Anarrayofb2Use, CopyBranchAstructBstruct2Use(mapOrigCopy, _astructbstruct2use))
	}

	return
}

func CopyBranchAstructBstruct2Use(mapOrigCopy map[any]any, astructbstruct2useFrom *AstructBstruct2Use) (astructbstruct2useTo *AstructBstruct2Use) {

	// astructbstruct2useFrom has already been copied
	if _astructbstruct2useTo, ok := mapOrigCopy[astructbstruct2useFrom]; ok {
		astructbstruct2useTo = _astructbstruct2useTo.(*AstructBstruct2Use)
		return
	}

	astructbstruct2useTo = new(AstructBstruct2Use)
	mapOrigCopy[astructbstruct2useFrom] = astructbstruct2useTo
	astructbstruct2useFrom.CopyBasicFields(astructbstruct2useTo)

	//insertion point for the staging of instances referenced by pointers
	if astructbstruct2useFrom.Bstrcut2 != nil {
		astructbstruct2useTo.Bstrcut2 = CopyBranchBstruct(mapOrigCopy, astructbstruct2useFrom.Bstrcut2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAstructBstructUse(mapOrigCopy map[any]any, astructbstructuseFrom *AstructBstructUse) (astructbstructuseTo *AstructBstructUse) {

	// astructbstructuseFrom has already been copied
	if _astructbstructuseTo, ok := mapOrigCopy[astructbstructuseFrom]; ok {
		astructbstructuseTo = _astructbstructuseTo.(*AstructBstructUse)
		return
	}

	astructbstructuseTo = new(AstructBstructUse)
	mapOrigCopy[astructbstructuseFrom] = astructbstructuseTo
	astructbstructuseFrom.CopyBasicFields(astructbstructuseTo)

	//insertion point for the staging of instances referenced by pointers
	if astructbstructuseFrom.Bstruct2 != nil {
		astructbstructuseTo.Bstruct2 = CopyBranchBstruct(mapOrigCopy, astructbstructuseFrom.Bstruct2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBstruct(mapOrigCopy map[any]any, bstructFrom *Bstruct) (bstructTo *Bstruct) {

	// bstructFrom has already been copied
	if _bstructTo, ok := mapOrigCopy[bstructFrom]; ok {
		bstructTo = _bstructTo.(*Bstruct)
		return
	}

	bstructTo = new(Bstruct)
	mapOrigCopy[bstructFrom] = bstructTo
	bstructFrom.CopyBasicFields(bstructTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDstruct(mapOrigCopy map[any]any, dstructFrom *Dstruct) (dstructTo *Dstruct) {

	// dstructFrom has already been copied
	if _dstructTo, ok := mapOrigCopy[dstructFrom]; ok {
		dstructTo = _dstructTo.(*Dstruct)
		return
	}

	dstructTo = new(Dstruct)
	mapOrigCopy[dstructFrom] = dstructTo
	dstructFrom.CopyBasicFields(dstructTo)

	//insertion point for the staging of instances referenced by pointers
	if dstructFrom.Gstruct != nil {
		dstructTo.Gstruct = CopyBranchGstruct(mapOrigCopy, dstructFrom.Gstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstructFrom.Anarrayofb {
		dstructTo.Anarrayofb = append(dstructTo.Anarrayofb, CopyBranchBstruct(mapOrigCopy, _bstruct))
	}
	for _, _gstruct := range dstructFrom.Gstructs {
		dstructTo.Gstructs = append(dstructTo.Gstructs, CopyBranchGstruct(mapOrigCopy, _gstruct))
	}

	return
}

func CopyBranchF0123456789012345678901234567890(mapOrigCopy map[any]any, f0123456789012345678901234567890From *F0123456789012345678901234567890) (f0123456789012345678901234567890To *F0123456789012345678901234567890) {

	// f0123456789012345678901234567890From has already been copied
	if _f0123456789012345678901234567890To, ok := mapOrigCopy[f0123456789012345678901234567890From]; ok {
		f0123456789012345678901234567890To = _f0123456789012345678901234567890To.(*F0123456789012345678901234567890)
		return
	}

	f0123456789012345678901234567890To = new(F0123456789012345678901234567890)
	mapOrigCopy[f0123456789012345678901234567890From] = f0123456789012345678901234567890To
	f0123456789012345678901234567890From.CopyBasicFields(f0123456789012345678901234567890To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGstruct(mapOrigCopy map[any]any, gstructFrom *Gstruct) (gstructTo *Gstruct) {

	// gstructFrom has already been copied
	if _gstructTo, ok := mapOrigCopy[gstructFrom]; ok {
		gstructTo = _gstructTo.(*Gstruct)
		return
	}

	gstructTo = new(Gstruct)
	mapOrigCopy[gstructFrom] = gstructTo
	gstructFrom.CopyBasicFields(gstructTo)

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
	case *Astruct:
		stage.UnstageBranchAstruct(target)

	case *AstructBstruct2Use:
		stage.UnstageBranchAstructBstruct2Use(target)

	case *AstructBstructUse:
		stage.UnstageBranchAstructBstructUse(target)

	case *Bstruct:
		stage.UnstageBranchBstruct(target)

	case *Dstruct:
		stage.UnstageBranchDstruct(target)

	case *F0123456789012345678901234567890:
		stage.UnstageBranchF0123456789012345678901234567890(target)

	case *Gstruct:
		stage.UnstageBranchGstruct(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAstruct(astruct *Astruct) {

	// check if instance is already staged
	if !IsStaged(stage, astruct) {
		return
	}

	astruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astruct.Associationtob != nil {
		UnstageBranch(stage, astruct.Associationtob)
	}
	if astruct.Anotherassociationtob_2 != nil {
		UnstageBranch(stage, astruct.Anotherassociationtob_2)
	}
	if astruct.Bstruct != nil {
		UnstageBranch(stage, astruct.Bstruct)
	}
	if astruct.Bstruct2 != nil {
		UnstageBranch(stage, astruct.Bstruct2)
	}
	if astruct.Dstruct != nil {
		UnstageBranch(stage, astruct.Dstruct)
	}
	if astruct.Dstruct2 != nil {
		UnstageBranch(stage, astruct.Dstruct2)
	}
	if astruct.Dstruct3 != nil {
		UnstageBranch(stage, astruct.Dstruct3)
	}
	if astruct.Dstruct4 != nil {
		UnstageBranch(stage, astruct.Dstruct4)
	}
	if astruct.AnAstruct != nil {
		UnstageBranch(stage, astruct.AnAstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range astruct.Anarrayofb {
		UnstageBranch(stage, _bstruct)
	}
	for _, _dstruct := range astruct.Dstruct4s {
		UnstageBranch(stage, _dstruct)
	}
	for _, _astruct := range astruct.Anarrayofa {
		UnstageBranch(stage, _astruct)
	}
	for _, _bstruct := range astruct.Anotherarrayofb {
		UnstageBranch(stage, _bstruct)
	}
	for _, _astructbstructuse := range astruct.AnarrayofbUse {
		UnstageBranch(stage, _astructbstructuse)
	}
	for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
		UnstageBranch(stage, _astructbstruct2use)
	}

}

func (stage *Stage) UnstageBranchAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) {

	// check if instance is already staged
	if !IsStaged(stage, astructbstruct2use) {
		return
	}

	astructbstruct2use.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astructbstruct2use.Bstrcut2 != nil {
		UnstageBranch(stage, astructbstruct2use.Bstrcut2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchAstructBstructUse(astructbstructuse *AstructBstructUse) {

	// check if instance is already staged
	if !IsStaged(stage, astructbstructuse) {
		return
	}

	astructbstructuse.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astructbstructuse.Bstruct2 != nil {
		UnstageBranch(stage, astructbstructuse.Bstruct2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchBstruct(bstruct *Bstruct) {

	// check if instance is already staged
	if !IsStaged(stage, bstruct) {
		return
	}

	bstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDstruct(dstruct *Dstruct) {

	// check if instance is already staged
	if !IsStaged(stage, dstruct) {
		return
	}

	dstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dstruct.Gstruct != nil {
		UnstageBranch(stage, dstruct.Gstruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstruct.Anarrayofb {
		UnstageBranch(stage, _bstruct)
	}
	for _, _gstruct := range dstruct.Gstructs {
		UnstageBranch(stage, _gstruct)
	}

}

func (stage *Stage) UnstageBranchF0123456789012345678901234567890(f0123456789012345678901234567890 *F0123456789012345678901234567890) {

	// check if instance is already staged
	if !IsStaged(stage, f0123456789012345678901234567890) {
		return
	}

	f0123456789012345678901234567890.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGstruct(gstruct *Gstruct) {

	// check if instance is already staged
	if !IsStaged(stage, gstruct) {
		return
	}

	gstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
		diffs = append(diffs, "Associationtob")
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
		diffs = append(diffs, astruct.GongMarshallField(stage, "Anarrayofb"))
	}
	if (astruct.Anotherassociationtob_2 == nil) != (astructOther.Anotherassociationtob_2 == nil) {
		diffs = append(diffs, "Anotherassociationtob_2")
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
		diffs = append(diffs, "Bstruct")
	} else if astruct.Bstruct != nil && astructOther.Bstruct != nil {
		if astruct.Bstruct != astructOther.Bstruct {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Bstruct"))
		}
	}
	if (astruct.Bstruct2 == nil) != (astructOther.Bstruct2 == nil) {
		diffs = append(diffs, "Bstruct2")
	} else if astruct.Bstruct2 != nil && astructOther.Bstruct2 != nil {
		if astruct.Bstruct2 != astructOther.Bstruct2 {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Bstruct2"))
		}
	}
	if (astruct.Dstruct == nil) != (astructOther.Dstruct == nil) {
		diffs = append(diffs, "Dstruct")
	} else if astruct.Dstruct != nil && astructOther.Dstruct != nil {
		if astruct.Dstruct != astructOther.Dstruct {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct"))
		}
	}
	if (astruct.Dstruct2 == nil) != (astructOther.Dstruct2 == nil) {
		diffs = append(diffs, "Dstruct2")
	} else if astruct.Dstruct2 != nil && astructOther.Dstruct2 != nil {
		if astruct.Dstruct2 != astructOther.Dstruct2 {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct2"))
		}
	}
	if (astruct.Dstruct3 == nil) != (astructOther.Dstruct3 == nil) {
		diffs = append(diffs, "Dstruct3")
	} else if astruct.Dstruct3 != nil && astructOther.Dstruct3 != nil {
		if astruct.Dstruct3 != astructOther.Dstruct3 {
			diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct3"))
		}
	}
	if (astruct.Dstruct4 == nil) != (astructOther.Dstruct4 == nil) {
		diffs = append(diffs, "Dstruct4")
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
		diffs = append(diffs, astruct.GongMarshallField(stage, "Dstruct4s"))
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
		diffs = append(diffs, astruct.GongMarshallField(stage, "Anarrayofa"))
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
		diffs = append(diffs, astruct.GongMarshallField(stage, "Anotherarrayofb"))
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
		diffs = append(diffs, astruct.GongMarshallField(stage, "AnarrayofbUse"))
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
		diffs = append(diffs, astruct.GongMarshallField(stage, "Anarrayofb2Use"))
	}
	if (astruct.AnAstruct == nil) != (astructOther.AnAstruct == nil) {
		diffs = append(diffs, "AnAstruct")
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
		diffs = append(diffs, "Bstrcut2")
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
		diffs = append(diffs, "Bstruct2")
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
		diffs = append(diffs, dstruct.GongMarshallField(stage, "Anarrayofb"))
	}
	if (dstruct.Gstruct == nil) != (dstructOther.Gstruct == nil) {
		diffs = append(diffs, "Gstruct")
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
		diffs = append(diffs, dstruct.GongMarshallField(stage, "Gstructs"))
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
