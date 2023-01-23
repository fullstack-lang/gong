package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	case *Astruct:
		ok = stage.IsStagedAstruct(target)
	case *Bstruct:
		ok = stage.IsStagedBstruct(target)
	}
	return
}

func (stage *StageStruct) IsStagedAstruct(astruct *Astruct) (ok bool) {

	_, ok = stage.Astructs[astruct]

	return
}

func (stage *StageStruct) IsStagedBstruct(bstruct *Bstruct) (ok bool) {

	_, ok = stage.Bstructs[bstruct]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	case *Astruct:
		stage.StageBranchAstruct(target)
	case *Bstruct:
		stage.StageBranchBstruct(target)
	}
}

func (stage *StageStruct) StageBranchAstruct(astruct *Astruct) {

	// check if instance is already staged
	if IsStaged(stage, astruct) {
		return
	}

	astruct.Stage()

	if astruct.Associationtob != nil {
		StageBranch(stage, astruct.Associationtob)
	}

	for _, _bstruct := range astruct.Anarrayofb {
		StageBranch(stage, _bstruct)
	}

}

// StageBranch stages bstruct and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers.
//
// the algorithm stops along the course of graph if a vertex is already staged
func (stage *StageStruct) StageBranchBstruct(bstruct *Bstruct) {

	// check if instance is already staged
	if stage.IsStagedBstruct(bstruct) {
		return
	}

	bstruct.Stage()

}
