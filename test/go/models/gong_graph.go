// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

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

	case *Fstruct:
		ok = stage.IsStagedFstruct(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAstruct(astruct *Astruct) (ok bool) {

	_, ok = stage.Astructs[astruct]

	return
}

func (stage *StageStruct) IsStagedAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) (ok bool) {

	_, ok = stage.AstructBstruct2Uses[astructbstruct2use]

	return
}

func (stage *StageStruct) IsStagedAstructBstructUse(astructbstructuse *AstructBstructUse) (ok bool) {

	_, ok = stage.AstructBstructUses[astructbstructuse]

	return
}

func (stage *StageStruct) IsStagedBstruct(bstruct *Bstruct) (ok bool) {

	_, ok = stage.Bstructs[bstruct]

	return
}

func (stage *StageStruct) IsStagedDstruct(dstruct *Dstruct) (ok bool) {

	_, ok = stage.Dstructs[dstruct]

	return
}

func (stage *StageStruct) IsStagedFstruct(fstruct *Fstruct) (ok bool) {

	_, ok = stage.Fstructs[fstruct]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

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

	case *Fstruct:
		stage.StageBranchFstruct(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAstruct(astruct *Astruct) {

	// check if instance is already staged
	if IsStaged(stage, astruct) {
		return
	}

	astruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astruct.AnonymousStructField1.Associationtob4 != nil {
		StageBranch(stage, astruct.AnonymousStructField1.Associationtob4)
	}
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

func (stage *StageStruct) StageBranchAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) {

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

func (stage *StageStruct) StageBranchAstructBstructUse(astructbstructuse *AstructBstructUse) {

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

func (stage *StageStruct) StageBranchBstruct(bstruct *Bstruct) {

	// check if instance is already staged
	if IsStaged(stage, bstruct) {
		return
	}

	bstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDstruct(dstruct *Dstruct) {

	// check if instance is already staged
	if IsStaged(stage, dstruct) {
		return
	}

	dstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstruct.Anarrayofb {
		StageBranch(stage, _bstruct)
	}

}

func (stage *StageStruct) StageBranchFstruct(fstruct *Fstruct) {

	// check if instance is already staged
	if IsStaged(stage, fstruct) {
		return
	}

	fstruct.Stage(stage)

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

	case *Fstruct:
		toT := CopyBranchFstruct(mapOrigCopy, fromT)
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
	if astructFrom.AnonymousStructField1.Associationtob4 != nil {
		astructTo.AnonymousStructField1.Associationtob4 = CopyBranchBstruct(mapOrigCopy, astructFrom.AnonymousStructField1.Associationtob4)
	}
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

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstructFrom.Anarrayofb {
		dstructTo.Anarrayofb = append(dstructTo.Anarrayofb, CopyBranchBstruct(mapOrigCopy, _bstruct))
	}

	return
}

func CopyBranchFstruct(mapOrigCopy map[any]any, fstructFrom *Fstruct) (fstructTo *Fstruct) {

	// fstructFrom has already been copied
	if _fstructTo, ok := mapOrigCopy[fstructFrom]; ok {
		fstructTo = _fstructTo.(*Fstruct)
		return
	}

	fstructTo = new(Fstruct)
	mapOrigCopy[fstructFrom] = fstructTo
	fstructFrom.CopyBasicFields(fstructTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

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

	case *Fstruct:
		stage.UnstageBranchFstruct(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAstruct(astruct *Astruct) {

	// check if instance is already staged
	if !IsStaged(stage, astruct) {
		return
	}

	astruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if astruct.AnonymousStructField1.Associationtob4 != nil {
		UnstageBranch(stage, astruct.AnonymousStructField1.Associationtob4)
	}
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

func (stage *StageStruct) UnstageBranchAstructBstruct2Use(astructbstruct2use *AstructBstruct2Use) {

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

func (stage *StageStruct) UnstageBranchAstructBstructUse(astructbstructuse *AstructBstructUse) {

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

func (stage *StageStruct) UnstageBranchBstruct(bstruct *Bstruct) {

	// check if instance is already staged
	if !IsStaged(stage, bstruct) {
		return
	}

	bstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDstruct(dstruct *Dstruct) {

	// check if instance is already staged
	if !IsStaged(stage, dstruct) {
		return
	}

	dstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _bstruct := range dstruct.Anarrayofb {
		UnstageBranch(stage, _bstruct)
	}

}

func (stage *StageStruct) UnstageBranchFstruct(fstruct *Fstruct) {

	// check if instance is already staged
	if !IsStaged(stage, fstruct) {
		return
	}

	fstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
