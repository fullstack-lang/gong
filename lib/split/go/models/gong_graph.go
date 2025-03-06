// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AsSplit:
		ok = stage.IsStagedAsSplit(target)

	case *AsSplitArea:
		ok = stage.IsStagedAsSplitArea(target)

	case *View:
		ok = stage.IsStagedView(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAsSplit(assplit *AsSplit) (ok bool) {

	_, ok = stage.AsSplits[assplit]

	return
}

func (stage *StageStruct) IsStagedAsSplitArea(assplitarea *AsSplitArea) (ok bool) {

	_, ok = stage.AsSplitAreas[assplitarea]

	return
}

func (stage *StageStruct) IsStagedView(view *View) (ok bool) {

	_, ok = stage.Views[view]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AsSplit:
		stage.StageBranchAsSplit(target)

	case *AsSplitArea:
		stage.StageBranchAsSplitArea(target)

	case *View:
		stage.StageBranchView(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAsSplit(assplit *AsSplit) {

	// check if instance is already staged
	if IsStaged(stage, assplit) {
		return
	}

	assplit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplit.AsSplitAreas {
		StageBranch(stage, _assplitarea)
	}

}

func (stage *StageStruct) StageBranchAsSplitArea(assplitarea *AsSplitArea) {

	// check if instance is already staged
	if IsStaged(stage, assplitarea) {
		return
	}

	assplitarea.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplit := range assplitarea.AsSplits {
		StageBranch(stage, _assplit)
	}

}

func (stage *StageStruct) StageBranchView(view *View) {

	// check if instance is already staged
	if IsStaged(stage, view) {
		return
	}

	view.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range view.RootAsSplitAreas {
		StageBranch(stage, _assplitarea)
	}

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
	case *AsSplit:
		toT := CopyBranchAsSplit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AsSplitArea:
		toT := CopyBranchAsSplitArea(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *View:
		toT := CopyBranchView(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAsSplit(mapOrigCopy map[any]any, assplitFrom *AsSplit) (assplitTo *AsSplit) {

	// assplitFrom has already been copied
	if _assplitTo, ok := mapOrigCopy[assplitFrom]; ok {
		assplitTo = _assplitTo.(*AsSplit)
		return
	}

	assplitTo = new(AsSplit)
	mapOrigCopy[assplitFrom] = assplitTo
	assplitFrom.CopyBasicFields(assplitTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplitFrom.AsSplitAreas {
		assplitTo.AsSplitAreas = append(assplitTo.AsSplitAreas, CopyBranchAsSplitArea(mapOrigCopy, _assplitarea))
	}

	return
}

func CopyBranchAsSplitArea(mapOrigCopy map[any]any, assplitareaFrom *AsSplitArea) (assplitareaTo *AsSplitArea) {

	// assplitareaFrom has already been copied
	if _assplitareaTo, ok := mapOrigCopy[assplitareaFrom]; ok {
		assplitareaTo = _assplitareaTo.(*AsSplitArea)
		return
	}

	assplitareaTo = new(AsSplitArea)
	mapOrigCopy[assplitareaFrom] = assplitareaTo
	assplitareaFrom.CopyBasicFields(assplitareaTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplit := range assplitareaFrom.AsSplits {
		assplitareaTo.AsSplits = append(assplitareaTo.AsSplits, CopyBranchAsSplit(mapOrigCopy, _assplit))
	}

	return
}

func CopyBranchView(mapOrigCopy map[any]any, viewFrom *View) (viewTo *View) {

	// viewFrom has already been copied
	if _viewTo, ok := mapOrigCopy[viewFrom]; ok {
		viewTo = _viewTo.(*View)
		return
	}

	viewTo = new(View)
	mapOrigCopy[viewFrom] = viewTo
	viewFrom.CopyBasicFields(viewTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range viewFrom.RootAsSplitAreas {
		viewTo.RootAsSplitAreas = append(viewTo.RootAsSplitAreas, CopyBranchAsSplitArea(mapOrigCopy, _assplitarea))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *AsSplit:
		stage.UnstageBranchAsSplit(target)

	case *AsSplitArea:
		stage.UnstageBranchAsSplitArea(target)

	case *View:
		stage.UnstageBranchView(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAsSplit(assplit *AsSplit) {

	// check if instance is already staged
	if !IsStaged(stage, assplit) {
		return
	}

	assplit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range assplit.AsSplitAreas {
		UnstageBranch(stage, _assplitarea)
	}

}

func (stage *StageStruct) UnstageBranchAsSplitArea(assplitarea *AsSplitArea) {

	// check if instance is already staged
	if !IsStaged(stage, assplitarea) {
		return
	}

	assplitarea.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplit := range assplitarea.AsSplits {
		UnstageBranch(stage, _assplit)
	}

}

func (stage *StageStruct) UnstageBranchView(view *View) {

	// check if instance is already staged
	if !IsStaged(stage, view) {
		return
	}

	view.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _assplitarea := range view.RootAsSplitAreas {
		UnstageBranch(stage, _assplitarea)
	}

}
