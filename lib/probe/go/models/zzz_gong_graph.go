// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Probe2:
		ok = stage.IsStagedProbe2(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedProbe2(probe2 *Probe2) (ok bool) {

	_, ok = stage.Probe2s[probe2]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Probe2:
		stage.StageBranchProbe2(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchProbe2(probe2 *Probe2) {

	// check if instance is already staged
	if IsStaged(stage, probe2) {
		return
	}

	probe2.Stage(stage)

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
	case *Probe2:
		toT := CopyBranchProbe2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchProbe2(mapOrigCopy map[any]any, probe2From *Probe2) (probe2To *Probe2) {

	// probe2From has already been copied
	if _probe2To, ok := mapOrigCopy[probe2From]; ok {
		probe2To = _probe2To.(*Probe2)
		return
	}

	probe2To = new(Probe2)
	mapOrigCopy[probe2From] = probe2To
	probe2From.CopyBasicFields(probe2To)

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
	case *Probe2:
		stage.UnstageBranchProbe2(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchProbe2(probe2 *Probe2) {

	// check if instance is already staged
	if !IsStaged(stage, probe2) {
		return
	}

	probe2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
