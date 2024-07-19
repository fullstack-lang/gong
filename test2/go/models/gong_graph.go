// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *A:
		ok = stage.IsStagedA(target)

	case *B:
		ok = stage.IsStagedB(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedA(a *A) (ok bool) {

	_, ok = stage.As[a]

	return
}

func (stage *StageStruct) IsStagedB(b *B) (ok bool) {

	_, ok = stage.Bs[b]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *A:
		stage.StageBranchA(target)

	case *B:
		stage.StageBranchB(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchA(a *A) {

	// check if instance is already staged
	if IsStaged(stage, a) {
		return
	}

	a.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a.B != nil {
		StageBranch(stage, a.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range a.Bs {
		StageBranch(stage, _b)
	}

}

func (stage *StageStruct) StageBranchB(b *B) {

	// check if instance is already staged
	if IsStaged(stage, b) {
		return
	}

	b.Stage(stage)

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
	case *A:
		toT := CopyBranchA(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *B:
		toT := CopyBranchB(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchA(mapOrigCopy map[any]any, aFrom *A) (aTo *A) {

	// aFrom has already been copied
	if _aTo, ok := mapOrigCopy[aFrom]; ok {
		aTo = _aTo.(*A)
		return
	}

	aTo = new(A)
	mapOrigCopy[aFrom] = aTo
	aFrom.CopyBasicFields(aTo)

	//insertion point for the staging of instances referenced by pointers
	if aFrom.B != nil {
		aTo.B = CopyBranchB(mapOrigCopy, aFrom.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range aFrom.Bs {
		aTo.Bs = append(aTo.Bs, CopyBranchB(mapOrigCopy, _b))
	}

	return
}

func CopyBranchB(mapOrigCopy map[any]any, bFrom *B) (bTo *B) {

	// bFrom has already been copied
	if _bTo, ok := mapOrigCopy[bFrom]; ok {
		bTo = _bTo.(*B)
		return
	}

	bTo = new(B)
	mapOrigCopy[bFrom] = bTo
	bFrom.CopyBasicFields(bTo)

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
	case *A:
		stage.UnstageBranchA(target)

	case *B:
		stage.UnstageBranchB(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchA(a *A) {

	// check if instance is already staged
	if !IsStaged(stage, a) {
		return
	}

	a.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a.B != nil {
		UnstageBranch(stage, a.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range a.Bs {
		UnstageBranch(stage, _b)
	}

}

func (stage *StageStruct) UnstageBranchB(b *B) {

	// check if instance is already staged
	if !IsStaged(stage, b) {
		return
	}

	b.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
