// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Complexity:
		ok = stage.IsStagedComplexity(target)

	case *DiagramFloss:
		ok = stage.IsStagedDiagramFloss(target)

	case *Effort:
		ok = stage.IsStagedEffort(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Performance:
		ok = stage.IsStagedPerformance(target)

	case *System:
		ok = stage.IsStagedSystem(target)

	case *SystemShape:
		ok = stage.IsStagedSystemShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Complexity:
		ok = stage.IsStagedComplexity(target)

	case *DiagramFloss:
		ok = stage.IsStagedDiagramFloss(target)

	case *Effort:
		ok = stage.IsStagedEffort(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Performance:
		ok = stage.IsStagedPerformance(target)

	case *System:
		ok = stage.IsStagedSystem(target)

	case *SystemShape:
		ok = stage.IsStagedSystemShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedComplexity(complexity *Complexity) (ok bool) {

	_, ok = stage.Complexitys[complexity]

	return
}

func (stage *Stage) IsStagedDiagramFloss(diagramfloss *DiagramFloss) (ok bool) {

	_, ok = stage.DiagramFlosss[diagramfloss]

	return
}

func (stage *Stage) IsStagedEffort(effort *Effort) (ok bool) {

	_, ok = stage.Efforts[effort]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedPerformance(performance *Performance) (ok bool) {

	_, ok = stage.Performances[performance]

	return
}

func (stage *Stage) IsStagedSystem(system *System) (ok bool) {

	_, ok = stage.Systems[system]

	return
}

func (stage *Stage) IsStagedSystemShape(systemshape *SystemShape) (ok bool) {

	_, ok = stage.SystemShapes[systemshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Complexity:
		stage.StageBranchComplexity(target)

	case *DiagramFloss:
		stage.StageBranchDiagramFloss(target)

	case *Effort:
		stage.StageBranchEffort(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Performance:
		stage.StageBranchPerformance(target)

	case *System:
		stage.StageBranchSystem(target)

	case *SystemShape:
		stage.StageBranchSystemShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchComplexity(complexity *Complexity) {

	// check if instance is already staged
	if IsStaged(stage, complexity) {
		return
	}

	complexity.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDiagramFloss(diagramfloss *DiagramFloss) {

	// check if instance is already staged
	if IsStaged(stage, diagramfloss) {
		return
	}

	diagramfloss.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramfloss.System_Shapes {
		StageBranch(stage, _systemshape)
	}
	for _, _system := range diagramfloss.SystemsWhoseNodeIsExpanded {
		StageBranch(stage, _system)
	}

}

func (stage *Stage) StageBranchEffort(effort *Effort) {

	// check if instance is already staged
	if IsStaged(stage, effort) {
		return
	}

	effort.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if IsStaged(stage, library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		StageBranch(stage, _library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		StageBranch(stage, _library)
	}
	for _, _system := range library.RootSystemes {
		StageBranch(stage, _system)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		StageBranch(stage, _system)
	}

}

func (stage *Stage) StageBranchPerformance(performance *Performance) {

	// check if instance is already staged
	if IsStaged(stage, performance) {
		return
	}

	performance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSystem(system *System) {

	// check if instance is already staged
	if IsStaged(stage, system) {
		return
	}

	system.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range system.Complexitys {
		StageBranch(stage, _complexity)
	}
	for _, _performance := range system.Performances {
		StageBranch(stage, _performance)
	}
	for _, _effort := range system.Efforts {
		StageBranch(stage, _effort)
	}
	for _, _diagramfloss := range system.DiagramFlosses {
		StageBranch(stage, _diagramfloss)
	}
	for _, _diagramfloss := range system.DiagramFlossWhoseNodeIsExpanded {
		StageBranch(stage, _diagramfloss)
	}
	for _, _system := range system.SubSystemes {
		StageBranch(stage, _system)
	}

}

func (stage *Stage) StageBranchSystemShape(systemshape *SystemShape) {

	// check if instance is already staged
	if IsStaged(stage, systemshape) {
		return
	}

	systemshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemshape.System != nil {
		StageBranch(stage, systemshape.System)
	}

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
	case *Complexity:
		toT := CopyBranchComplexity(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramFloss:
		toT := CopyBranchDiagramFloss(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Effort:
		toT := CopyBranchEffort(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Performance:
		toT := CopyBranchPerformance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System:
		toT := CopyBranchSystem(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SystemShape:
		toT := CopyBranchSystemShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchComplexity(mapOrigCopy map[any]any, complexityFrom *Complexity) (complexityTo *Complexity) {

	// complexityFrom has already been copied
	if _complexityTo, ok := mapOrigCopy[complexityFrom]; ok {
		complexityTo = _complexityTo.(*Complexity)
		return
	}

	complexityTo = new(Complexity)
	mapOrigCopy[complexityFrom] = complexityTo
	complexityFrom.CopyBasicFields(complexityTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDiagramFloss(mapOrigCopy map[any]any, diagramflossFrom *DiagramFloss) (diagramflossTo *DiagramFloss) {

	// diagramflossFrom has already been copied
	if _diagramflossTo, ok := mapOrigCopy[diagramflossFrom]; ok {
		diagramflossTo = _diagramflossTo.(*DiagramFloss)
		return
	}

	diagramflossTo = new(DiagramFloss)
	mapOrigCopy[diagramflossFrom] = diagramflossTo
	diagramflossFrom.CopyBasicFields(diagramflossTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramflossFrom.System_Shapes {
		diagramflossTo.System_Shapes = append(diagramflossTo.System_Shapes, CopyBranchSystemShape(mapOrigCopy, _systemshape))
	}
	for _, _system := range diagramflossFrom.SystemsWhoseNodeIsExpanded {
		diagramflossTo.SystemsWhoseNodeIsExpanded = append(diagramflossTo.SystemsWhoseNodeIsExpanded, CopyBranchSystem(mapOrigCopy, _system))
	}

	return
}

func CopyBranchEffort(mapOrigCopy map[any]any, effortFrom *Effort) (effortTo *Effort) {

	// effortFrom has already been copied
	if _effortTo, ok := mapOrigCopy[effortFrom]; ok {
		effortTo = _effortTo.(*Effort)
		return
	}

	effortTo = new(Effort)
	mapOrigCopy[effortFrom] = effortTo
	effortFrom.CopyBasicFields(effortTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.CopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _system := range libraryFrom.RootSystemes {
		libraryTo.RootSystemes = append(libraryTo.RootSystemes, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _system := range libraryFrom.SystemsWhoseNodeIsExpanded {
		libraryTo.SystemsWhoseNodeIsExpanded = append(libraryTo.SystemsWhoseNodeIsExpanded, CopyBranchSystem(mapOrigCopy, _system))
	}

	return
}

func CopyBranchPerformance(mapOrigCopy map[any]any, performanceFrom *Performance) (performanceTo *Performance) {

	// performanceFrom has already been copied
	if _performanceTo, ok := mapOrigCopy[performanceFrom]; ok {
		performanceTo = _performanceTo.(*Performance)
		return
	}

	performanceTo = new(Performance)
	mapOrigCopy[performanceFrom] = performanceTo
	performanceFrom.CopyBasicFields(performanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSystem(mapOrigCopy map[any]any, systemFrom *System) (systemTo *System) {

	// systemFrom has already been copied
	if _systemTo, ok := mapOrigCopy[systemFrom]; ok {
		systemTo = _systemTo.(*System)
		return
	}

	systemTo = new(System)
	mapOrigCopy[systemFrom] = systemTo
	systemFrom.CopyBasicFields(systemTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range systemFrom.Complexitys {
		systemTo.Complexitys = append(systemTo.Complexitys, CopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range systemFrom.Performances {
		systemTo.Performances = append(systemTo.Performances, CopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range systemFrom.Efforts {
		systemTo.Efforts = append(systemTo.Efforts, CopyBranchEffort(mapOrigCopy, _effort))
	}
	for _, _diagramfloss := range systemFrom.DiagramFlosses {
		systemTo.DiagramFlosses = append(systemTo.DiagramFlosses, CopyBranchDiagramFloss(mapOrigCopy, _diagramfloss))
	}
	for _, _diagramfloss := range systemFrom.DiagramFlossWhoseNodeIsExpanded {
		systemTo.DiagramFlossWhoseNodeIsExpanded = append(systemTo.DiagramFlossWhoseNodeIsExpanded, CopyBranchDiagramFloss(mapOrigCopy, _diagramfloss))
	}
	for _, _system := range systemFrom.SubSystemes {
		systemTo.SubSystemes = append(systemTo.SubSystemes, CopyBranchSystem(mapOrigCopy, _system))
	}

	return
}

func CopyBranchSystemShape(mapOrigCopy map[any]any, systemshapeFrom *SystemShape) (systemshapeTo *SystemShape) {

	// systemshapeFrom has already been copied
	if _systemshapeTo, ok := mapOrigCopy[systemshapeFrom]; ok {
		systemshapeTo = _systemshapeTo.(*SystemShape)
		return
	}

	systemshapeTo = new(SystemShape)
	mapOrigCopy[systemshapeFrom] = systemshapeTo
	systemshapeFrom.CopyBasicFields(systemshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if systemshapeFrom.System != nil {
		systemshapeTo.System = CopyBranchSystem(mapOrigCopy, systemshapeFrom.System)
	}

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
	case *Complexity:
		stage.UnstageBranchComplexity(target)

	case *DiagramFloss:
		stage.UnstageBranchDiagramFloss(target)

	case *Effort:
		stage.UnstageBranchEffort(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Performance:
		stage.UnstageBranchPerformance(target)

	case *System:
		stage.UnstageBranchSystem(target)

	case *SystemShape:
		stage.UnstageBranchSystemShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchComplexity(complexity *Complexity) {

	// check if instance is already staged
	if !IsStaged(stage, complexity) {
		return
	}

	complexity.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDiagramFloss(diagramfloss *DiagramFloss) {

	// check if instance is already staged
	if !IsStaged(stage, diagramfloss) {
		return
	}

	diagramfloss.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramfloss.System_Shapes {
		UnstageBranch(stage, _systemshape)
	}
	for _, _system := range diagramfloss.SystemsWhoseNodeIsExpanded {
		UnstageBranch(stage, _system)
	}

}

func (stage *Stage) UnstageBranchEffort(effort *Effort) {

	// check if instance is already staged
	if !IsStaged(stage, effort) {
		return
	}

	effort.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !IsStaged(stage, library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		UnstageBranch(stage, _library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		UnstageBranch(stage, _library)
	}
	for _, _system := range library.RootSystemes {
		UnstageBranch(stage, _system)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		UnstageBranch(stage, _system)
	}

}

func (stage *Stage) UnstageBranchPerformance(performance *Performance) {

	// check if instance is already staged
	if !IsStaged(stage, performance) {
		return
	}

	performance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSystem(system *System) {

	// check if instance is already staged
	if !IsStaged(stage, system) {
		return
	}

	system.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range system.Complexitys {
		UnstageBranch(stage, _complexity)
	}
	for _, _performance := range system.Performances {
		UnstageBranch(stage, _performance)
	}
	for _, _effort := range system.Efforts {
		UnstageBranch(stage, _effort)
	}
	for _, _diagramfloss := range system.DiagramFlosses {
		UnstageBranch(stage, _diagramfloss)
	}
	for _, _diagramfloss := range system.DiagramFlossWhoseNodeIsExpanded {
		UnstageBranch(stage, _diagramfloss)
	}
	for _, _system := range system.SubSystemes {
		UnstageBranch(stage, _system)
	}

}

func (stage *Stage) UnstageBranchSystemShape(systemshape *SystemShape) {

	// check if instance is already staged
	if !IsStaged(stage, systemshape) {
		return
	}

	systemshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemshape.System != nil {
		UnstageBranch(stage, systemshape.System)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Complexity) GongReconstructPointersFromReferences(stage *Stage, instance *Complexity) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DiagramFloss) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramFloss) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.System_Shapes = reference.System_Shapes[:0]
	for _, _b := range instance.System_Shapes {
		reference.System_Shapes = append(reference.System_Shapes, stage.SystemShapes_reference[_b])
	}
	reference.SystemsWhoseNodeIsExpanded = reference.SystemsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SystemsWhoseNodeIsExpanded {
		reference.SystemsWhoseNodeIsExpanded = append(reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference[_b])
	}
}

func (reference *Effort) GongReconstructPointersFromReferences(stage *Stage, instance *Effort) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.SubLibrariesWhoseNodeIsExpanded = reference.SubLibrariesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SubLibrariesWhoseNodeIsExpanded {
		reference.SubLibrariesWhoseNodeIsExpanded = append(reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference[_b])
	}
	reference.RootSystemes = reference.RootSystemes[:0]
	for _, _b := range instance.RootSystemes {
		reference.RootSystemes = append(reference.RootSystemes, stage.Systems_reference[_b])
	}
	reference.SystemsWhoseNodeIsExpanded = reference.SystemsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SystemsWhoseNodeIsExpanded {
		reference.SystemsWhoseNodeIsExpanded = append(reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference[_b])
	}
}

func (reference *Performance) GongReconstructPointersFromReferences(stage *Stage, instance *Performance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *System) GongReconstructPointersFromReferences(stage *Stage, instance *System) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Complexitys = reference.Complexitys[:0]
	for _, _b := range instance.Complexitys {
		reference.Complexitys = append(reference.Complexitys, stage.Complexitys_reference[_b])
	}
	reference.Performances = reference.Performances[:0]
	for _, _b := range instance.Performances {
		reference.Performances = append(reference.Performances, stage.Performances_reference[_b])
	}
	reference.Efforts = reference.Efforts[:0]
	for _, _b := range instance.Efforts {
		reference.Efforts = append(reference.Efforts, stage.Efforts_reference[_b])
	}
	reference.DiagramFlosses = reference.DiagramFlosses[:0]
	for _, _b := range instance.DiagramFlosses {
		reference.DiagramFlosses = append(reference.DiagramFlosses, stage.DiagramFlosss_reference[_b])
	}
	reference.DiagramFlossWhoseNodeIsExpanded = reference.DiagramFlossWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DiagramFlossWhoseNodeIsExpanded {
		reference.DiagramFlossWhoseNodeIsExpanded = append(reference.DiagramFlossWhoseNodeIsExpanded, stage.DiagramFlosss_reference[_b])
	}
	reference.SubSystemes = reference.SubSystemes[:0]
	for _, _b := range instance.SubSystemes {
		reference.SubSystemes = append(reference.SubSystemes, stage.Systems_reference[_b])
	}
}

func (reference *SystemShape) GongReconstructPointersFromReferences(stage *Stage, instance *SystemShape) {
	// insertion point for pointers field
	if instance.System != nil {
		reference.System = stage.Systems_reference[instance.System]
	}
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Complexity) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DiagramFloss) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _System_Shapes []*SystemShape
	for _, _reference := range reference.System_Shapes {
		if _instance, ok := stage.SystemShapes_instance[_reference]; ok {
			_System_Shapes = append(_System_Shapes, _instance)
		}
	}
	reference.System_Shapes = _System_Shapes
	var _SystemsWhoseNodeIsExpanded []*System
	for _, _reference := range reference.SystemsWhoseNodeIsExpanded {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SystemsWhoseNodeIsExpanded = append(_SystemsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SystemsWhoseNodeIsExpanded = _SystemsWhoseNodeIsExpanded
}

func (reference *Effort) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
	var _SubLibrariesWhoseNodeIsExpanded []*Library
	for _, _reference := range reference.SubLibrariesWhoseNodeIsExpanded {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibrariesWhoseNodeIsExpanded = append(_SubLibrariesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SubLibrariesWhoseNodeIsExpanded = _SubLibrariesWhoseNodeIsExpanded
	var _RootSystemes []*System
	for _, _reference := range reference.RootSystemes {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_RootSystemes = append(_RootSystemes, _instance)
		}
	}
	reference.RootSystemes = _RootSystemes
	var _SystemsWhoseNodeIsExpanded []*System
	for _, _reference := range reference.SystemsWhoseNodeIsExpanded {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SystemsWhoseNodeIsExpanded = append(_SystemsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SystemsWhoseNodeIsExpanded = _SystemsWhoseNodeIsExpanded
}

func (reference *Performance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *System) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Complexitys []*Complexity
	for _, _reference := range reference.Complexitys {
		if _instance, ok := stage.Complexitys_instance[_reference]; ok {
			_Complexitys = append(_Complexitys, _instance)
		}
	}
	reference.Complexitys = _Complexitys
	var _Performances []*Performance
	for _, _reference := range reference.Performances {
		if _instance, ok := stage.Performances_instance[_reference]; ok {
			_Performances = append(_Performances, _instance)
		}
	}
	reference.Performances = _Performances
	var _Efforts []*Effort
	for _, _reference := range reference.Efforts {
		if _instance, ok := stage.Efforts_instance[_reference]; ok {
			_Efforts = append(_Efforts, _instance)
		}
	}
	reference.Efforts = _Efforts
	var _DiagramFlosses []*DiagramFloss
	for _, _reference := range reference.DiagramFlosses {
		if _instance, ok := stage.DiagramFlosss_instance[_reference]; ok {
			_DiagramFlosses = append(_DiagramFlosses, _instance)
		}
	}
	reference.DiagramFlosses = _DiagramFlosses
	var _DiagramFlossWhoseNodeIsExpanded []*DiagramFloss
	for _, _reference := range reference.DiagramFlossWhoseNodeIsExpanded {
		if _instance, ok := stage.DiagramFlosss_instance[_reference]; ok {
			_DiagramFlossWhoseNodeIsExpanded = append(_DiagramFlossWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DiagramFlossWhoseNodeIsExpanded = _DiagramFlossWhoseNodeIsExpanded
	var _SubSystemes []*System
	for _, _reference := range reference.SubSystemes {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SubSystemes = append(_SubSystemes, _instance)
		}
	}
	reference.SubSystemes = _SubSystemes
}

func (reference *SystemShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.System; _reference != nil {
		reference.System = nil
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			reference.System = _instance
		}
	}
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (complexity *Complexity) GongDiff(stage *Stage, complexityOther *Complexity) (diffs []string) {
	// insertion point for field diffs
	if complexity.Name != complexityOther.Name {
		diffs = append(diffs, complexity.GongMarshallField(stage, "Name"))
	}
	if complexity.Strength != complexityOther.Strength {
		diffs = append(diffs, complexity.GongMarshallField(stage, "Strength"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagramfloss *DiagramFloss) GongDiff(stage *Stage, diagramflossOther *DiagramFloss) (diffs []string) {
	// insertion point for field diffs
	if diagramfloss.Name != diagramflossOther.Name {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "Name"))
	}
	if diagramfloss.Description != diagramflossOther.Description {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "Description"))
	}
	if diagramfloss.ComputedPrefix != diagramflossOther.ComputedPrefix {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagramfloss.IsExpanded != diagramflossOther.IsExpanded {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "IsExpanded"))
	}
	if diagramfloss.IsChecked != diagramflossOther.IsChecked {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "IsChecked"))
	}
	if diagramfloss.IsEditable_ != diagramflossOther.IsEditable_ {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "IsEditable_"))
	}
	if diagramfloss.IsShowPrefix != diagramflossOther.IsShowPrefix {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "IsShowPrefix"))
	}
	if diagramfloss.DefaultBoxWidth != diagramflossOther.DefaultBoxWidth {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "DefaultBoxWidth"))
	}
	if diagramfloss.DefaultBoxHeigth != diagramflossOther.DefaultBoxHeigth {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "DefaultBoxHeigth"))
	}
	if diagramfloss.Width != diagramflossOther.Width {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "Width"))
	}
	if diagramfloss.Height != diagramflossOther.Height {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "Height"))
	}
	System_ShapesDifferent := false
	if len(diagramfloss.System_Shapes) != len(diagramflossOther.System_Shapes) {
		System_ShapesDifferent = true
	} else {
		for i := range diagramfloss.System_Shapes {
			if (diagramfloss.System_Shapes[i] == nil) != (diagramflossOther.System_Shapes[i] == nil) {
				System_ShapesDifferent = true
				break
			} else if diagramfloss.System_Shapes[i] != nil && diagramflossOther.System_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramfloss.System_Shapes[i] != diagramflossOther.System_Shapes[i] {
					System_ShapesDifferent = true
					break
				}
			}
		}
	}
	if System_ShapesDifferent {
		ops := Diff(stage, diagramfloss, diagramflossOther, "System_Shapes", diagramflossOther.System_Shapes, diagramfloss.System_Shapes)
		diffs = append(diffs, ops)
	}
	if diagramfloss.IsSystemsNodeExpanded != diagramflossOther.IsSystemsNodeExpanded {
		diffs = append(diffs, diagramfloss.GongMarshallField(stage, "IsSystemsNodeExpanded"))
	}
	SystemsWhoseNodeIsExpandedDifferent := false
	if len(diagramfloss.SystemsWhoseNodeIsExpanded) != len(diagramflossOther.SystemsWhoseNodeIsExpanded) {
		SystemsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramfloss.SystemsWhoseNodeIsExpanded {
			if (diagramfloss.SystemsWhoseNodeIsExpanded[i] == nil) != (diagramflossOther.SystemsWhoseNodeIsExpanded[i] == nil) {
				SystemsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramfloss.SystemsWhoseNodeIsExpanded[i] != nil && diagramflossOther.SystemsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramfloss.SystemsWhoseNodeIsExpanded[i] != diagramflossOther.SystemsWhoseNodeIsExpanded[i] {
					SystemsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SystemsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramfloss, diagramflossOther, "SystemsWhoseNodeIsExpanded", diagramflossOther.SystemsWhoseNodeIsExpanded, diagramfloss.SystemsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (effort *Effort) GongDiff(stage *Stage, effortOther *Effort) (diffs []string) {
	// insertion point for field diffs
	if effort.Name != effortOther.Name {
		diffs = append(diffs, effort.GongMarshallField(stage, "Name"))
	}
	if effort.Strength != effortOther.Strength {
		diffs = append(diffs, effort.GongMarshallField(stage, "Strength"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (library *Library) GongDiff(stage *Stage, libraryOther *Library) (diffs []string) {
	// insertion point for field diffs
	if library.Name != libraryOther.Name {
		diffs = append(diffs, library.GongMarshallField(stage, "Name"))
	}
	if library.Description != libraryOther.Description {
		diffs = append(diffs, library.GongMarshallField(stage, "Description"))
	}
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
	}
	if library.IsExpanded != libraryOther.IsExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpanded"))
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	SubLibrariesDifferent := false
	if len(library.SubLibraries) != len(libraryOther.SubLibraries) {
		SubLibrariesDifferent = true
	} else {
		for i := range library.SubLibraries {
			if (library.SubLibraries[i] == nil) != (libraryOther.SubLibraries[i] == nil) {
				SubLibrariesDifferent = true
				break
			} else if library.SubLibraries[i] != nil && libraryOther.SubLibraries[i] != nil {
				// this is a pointer comparaison
				if library.SubLibraries[i] != libraryOther.SubLibraries[i] {
					SubLibrariesDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries)
		diffs = append(diffs, ops)
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	SubLibrariesWhoseNodeIsExpandedDifferent := false
	if len(library.SubLibrariesWhoseNodeIsExpanded) != len(libraryOther.SubLibrariesWhoseNodeIsExpanded) {
		SubLibrariesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.SubLibrariesWhoseNodeIsExpanded {
			if (library.SubLibrariesWhoseNodeIsExpanded[i] == nil) != (libraryOther.SubLibrariesWhoseNodeIsExpanded[i] == nil) {
				SubLibrariesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.SubLibrariesWhoseNodeIsExpanded[i] != nil && libraryOther.SubLibrariesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.SubLibrariesWhoseNodeIsExpanded[i] != libraryOther.SubLibrariesWhoseNodeIsExpanded[i] {
					SubLibrariesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	RootSystemesDifferent := false
	if len(library.RootSystemes) != len(libraryOther.RootSystemes) {
		RootSystemesDifferent = true
	} else {
		for i := range library.RootSystemes {
			if (library.RootSystemes[i] == nil) != (libraryOther.RootSystemes[i] == nil) {
				RootSystemesDifferent = true
				break
			} else if library.RootSystemes[i] != nil && libraryOther.RootSystemes[i] != nil {
				// this is a pointer comparaison
				if library.RootSystemes[i] != libraryOther.RootSystemes[i] {
					RootSystemesDifferent = true
					break
				}
			}
		}
	}
	if RootSystemesDifferent {
		ops := Diff(stage, library, libraryOther, "RootSystemes", libraryOther.RootSystemes, library.RootSystemes)
		diffs = append(diffs, ops)
	}
	if library.IsSystemesNodeExpanded != libraryOther.IsSystemesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSystemesNodeExpanded"))
	}
	SystemsWhoseNodeIsExpandedDifferent := false
	if len(library.SystemsWhoseNodeIsExpanded) != len(libraryOther.SystemsWhoseNodeIsExpanded) {
		SystemsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.SystemsWhoseNodeIsExpanded {
			if (library.SystemsWhoseNodeIsExpanded[i] == nil) != (libraryOther.SystemsWhoseNodeIsExpanded[i] == nil) {
				SystemsWhoseNodeIsExpandedDifferent = true
				break
			} else if library.SystemsWhoseNodeIsExpanded[i] != nil && libraryOther.SystemsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.SystemsWhoseNodeIsExpanded[i] != libraryOther.SystemsWhoseNodeIsExpanded[i] {
					SystemsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SystemsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "SystemsWhoseNodeIsExpanded", libraryOther.SystemsWhoseNodeIsExpanded, library.SystemsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (performance *Performance) GongDiff(stage *Stage, performanceOther *Performance) (diffs []string) {
	// insertion point for field diffs
	if performance.Name != performanceOther.Name {
		diffs = append(diffs, performance.GongMarshallField(stage, "Name"))
	}
	if performance.Strength != performanceOther.Strength {
		diffs = append(diffs, performance.GongMarshallField(stage, "Strength"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (system *System) GongDiff(stage *Stage, systemOther *System) (diffs []string) {
	// insertion point for field diffs
	if system.Name != systemOther.Name {
		diffs = append(diffs, system.GongMarshallField(stage, "Name"))
	}
	if system.Description != systemOther.Description {
		diffs = append(diffs, system.GongMarshallField(stage, "Description"))
	}
	ComplexitysDifferent := false
	if len(system.Complexitys) != len(systemOther.Complexitys) {
		ComplexitysDifferent = true
	} else {
		for i := range system.Complexitys {
			if (system.Complexitys[i] == nil) != (systemOther.Complexitys[i] == nil) {
				ComplexitysDifferent = true
				break
			} else if system.Complexitys[i] != nil && systemOther.Complexitys[i] != nil {
				// this is a pointer comparaison
				if system.Complexitys[i] != systemOther.Complexitys[i] {
					ComplexitysDifferent = true
					break
				}
			}
		}
	}
	if ComplexitysDifferent {
		ops := Diff(stage, system, systemOther, "Complexitys", systemOther.Complexitys, system.Complexitys)
		diffs = append(diffs, ops)
	}
	PerformancesDifferent := false
	if len(system.Performances) != len(systemOther.Performances) {
		PerformancesDifferent = true
	} else {
		for i := range system.Performances {
			if (system.Performances[i] == nil) != (systemOther.Performances[i] == nil) {
				PerformancesDifferent = true
				break
			} else if system.Performances[i] != nil && systemOther.Performances[i] != nil {
				// this is a pointer comparaison
				if system.Performances[i] != systemOther.Performances[i] {
					PerformancesDifferent = true
					break
				}
			}
		}
	}
	if PerformancesDifferent {
		ops := Diff(stage, system, systemOther, "Performances", systemOther.Performances, system.Performances)
		diffs = append(diffs, ops)
	}
	EffortsDifferent := false
	if len(system.Efforts) != len(systemOther.Efforts) {
		EffortsDifferent = true
	} else {
		for i := range system.Efforts {
			if (system.Efforts[i] == nil) != (systemOther.Efforts[i] == nil) {
				EffortsDifferent = true
				break
			} else if system.Efforts[i] != nil && systemOther.Efforts[i] != nil {
				// this is a pointer comparaison
				if system.Efforts[i] != systemOther.Efforts[i] {
					EffortsDifferent = true
					break
				}
			}
		}
	}
	if EffortsDifferent {
		ops := Diff(stage, system, systemOther, "Efforts", systemOther.Efforts, system.Efforts)
		diffs = append(diffs, ops)
	}
	if system.ComputedPrefix != systemOther.ComputedPrefix {
		diffs = append(diffs, system.GongMarshallField(stage, "ComputedPrefix"))
	}
	if system.IsExpanded != systemOther.IsExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsExpanded"))
	}
	if system.SVG_Path != systemOther.SVG_Path {
		diffs = append(diffs, system.GongMarshallField(stage, "SVG_Path"))
	}
	if system.InverseAppliedScaling != systemOther.InverseAppliedScaling {
		diffs = append(diffs, system.GongMarshallField(stage, "InverseAppliedScaling"))
	}
	DiagramFlossesDifferent := false
	if len(system.DiagramFlosses) != len(systemOther.DiagramFlosses) {
		DiagramFlossesDifferent = true
	} else {
		for i := range system.DiagramFlosses {
			if (system.DiagramFlosses[i] == nil) != (systemOther.DiagramFlosses[i] == nil) {
				DiagramFlossesDifferent = true
				break
			} else if system.DiagramFlosses[i] != nil && systemOther.DiagramFlosses[i] != nil {
				// this is a pointer comparaison
				if system.DiagramFlosses[i] != systemOther.DiagramFlosses[i] {
					DiagramFlossesDifferent = true
					break
				}
			}
		}
	}
	if DiagramFlossesDifferent {
		ops := Diff(stage, system, systemOther, "DiagramFlosses", systemOther.DiagramFlosses, system.DiagramFlosses)
		diffs = append(diffs, ops)
	}
	DiagramFlossWhoseNodeIsExpandedDifferent := false
	if len(system.DiagramFlossWhoseNodeIsExpanded) != len(systemOther.DiagramFlossWhoseNodeIsExpanded) {
		DiagramFlossWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.DiagramFlossWhoseNodeIsExpanded {
			if (system.DiagramFlossWhoseNodeIsExpanded[i] == nil) != (systemOther.DiagramFlossWhoseNodeIsExpanded[i] == nil) {
				DiagramFlossWhoseNodeIsExpandedDifferent = true
				break
			} else if system.DiagramFlossWhoseNodeIsExpanded[i] != nil && systemOther.DiagramFlossWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.DiagramFlossWhoseNodeIsExpanded[i] != systemOther.DiagramFlossWhoseNodeIsExpanded[i] {
					DiagramFlossWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DiagramFlossWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "DiagramFlossWhoseNodeIsExpanded", systemOther.DiagramFlossWhoseNodeIsExpanded, system.DiagramFlossWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if system.IsSubSystemNodeExpanded != systemOther.IsSubSystemNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
	}
	SubSystemesDifferent := false
	if len(system.SubSystemes) != len(systemOther.SubSystemes) {
		SubSystemesDifferent = true
	} else {
		for i := range system.SubSystemes {
			if (system.SubSystemes[i] == nil) != (systemOther.SubSystemes[i] == nil) {
				SubSystemesDifferent = true
				break
			} else if system.SubSystemes[i] != nil && systemOther.SubSystemes[i] != nil {
				// this is a pointer comparaison
				if system.SubSystemes[i] != systemOther.SubSystemes[i] {
					SubSystemesDifferent = true
					break
				}
			}
		}
	}
	if SubSystemesDifferent {
		ops := Diff(stage, system, systemOther, "SubSystemes", systemOther.SubSystemes, system.SubSystemes)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (systemshape *SystemShape) GongDiff(stage *Stage, systemshapeOther *SystemShape) (diffs []string) {
	// insertion point for field diffs
	if systemshape.Name != systemshapeOther.Name {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "Name"))
	}
	if (systemshape.System == nil) != (systemshapeOther.System == nil) {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "System"))
	} else if systemshape.System != nil && systemshapeOther.System != nil {
		if systemshape.System != systemshapeOther.System {
			diffs = append(diffs, systemshape.GongMarshallField(stage, "System"))
		}
	}
	if systemshape.IsExpanded != systemshapeOther.IsExpanded {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "IsExpanded"))
	}
	if systemshape.X != systemshapeOther.X {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "X"))
	}
	if systemshape.Y != systemshapeOther.Y {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "Y"))
	}
	if systemshape.Width != systemshapeOther.Width {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "Width"))
	}
	if systemshape.Height != systemshapeOther.Height {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "Height"))
	}
	if systemshape.IsHidden != systemshapeOther.IsHidden {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
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

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
