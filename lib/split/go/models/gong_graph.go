// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *SplitArea:
		ok = stage.IsStagedSplitArea(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedSplitArea(splitarea *SplitArea) (ok bool) {

	_, ok = stage.SplitAreas[splitarea]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *SplitArea:
		stage.StageBranchSplitArea(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchSplitArea(splitarea *SplitArea) {

	// check if instance is already staged
	if IsStaged(stage, splitarea) {
		return
	}

	splitarea.Stage(stage)

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
	case *SplitArea:
		toT := CopyBranchSplitArea(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchSplitArea(mapOrigCopy map[any]any, splitareaFrom *SplitArea) (splitareaTo *SplitArea) {

	// splitareaFrom has already been copied
	if _splitareaTo, ok := mapOrigCopy[splitareaFrom]; ok {
		splitareaTo = _splitareaTo.(*SplitArea)
		return
	}

	splitareaTo = new(SplitArea)
	mapOrigCopy[splitareaFrom] = splitareaTo
	splitareaFrom.CopyBasicFields(splitareaTo)

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
	case *SplitArea:
		stage.UnstageBranchSplitArea(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchSplitArea(splitarea *SplitArea) {

	// check if instance is already staged
	if !IsStaged(stage, splitarea) {
		return
	}

	splitarea.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
