// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *DiagramProcess:
		ok = stage.IsStagedDiagramProcess(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Participant:
		ok = stage.IsStagedParticipant(target)

	case *Process:
		ok = stage.IsStagedProcess(target)

	case *ProcessCompositionShape:
		ok = stage.IsStagedProcessCompositionShape(target)

	case *ProcessShape:
		ok = stage.IsStagedProcessShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *DiagramProcess:
		ok = stage.IsStagedDiagramProcess(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Participant:
		ok = stage.IsStagedParticipant(target)

	case *Process:
		ok = stage.IsStagedProcess(target)

	case *ProcessCompositionShape:
		ok = stage.IsStagedProcessCompositionShape(target)

	case *ProcessShape:
		ok = stage.IsStagedProcessShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedDiagramProcess(diagramprocess *DiagramProcess) (ok bool) {

	_, ok = stage.DiagramProcesss[diagramprocess]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedParticipant(participant *Participant) (ok bool) {

	_, ok = stage.Participants[participant]

	return
}

func (stage *Stage) IsStagedProcess(process *Process) (ok bool) {

	_, ok = stage.Processs[process]

	return
}

func (stage *Stage) IsStagedProcessCompositionShape(processcompositionshape *ProcessCompositionShape) (ok bool) {

	_, ok = stage.ProcessCompositionShapes[processcompositionshape]

	return
}

func (stage *Stage) IsStagedProcessShape(processshape *ProcessShape) (ok bool) {

	_, ok = stage.ProcessShapes[processshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *DiagramProcess:
		stage.StageBranchDiagramProcess(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Participant:
		stage.StageBranchParticipant(target)

	case *Process:
		stage.StageBranchProcess(target)

	case *ProcessCompositionShape:
		stage.StageBranchProcessCompositionShape(target)

	case *ProcessShape:
		stage.StageBranchProcessShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchDiagramProcess(diagramprocess *DiagramProcess) {

	// check if instance is already staged
	if IsStaged(stage, diagramprocess) {
		return
	}

	diagramprocess.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _processshape := range diagramprocess.Process_Shapes {
		StageBranch(stage, _processshape)
	}
	for _, _process := range diagramprocess.ProcesssWhoseNodeIsExpanded {
		StageBranch(stage, _process)
	}
	for _, _processcompositionshape := range diagramprocess.ProcessComposition_Shapes {
		StageBranch(stage, _processcompositionshape)
	}

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
	for _, _process := range library.RootProcesses {
		StageBranch(stage, _process)
	}
	for _, _process := range library.ProcesssWhoseNodeIsExpanded {
		StageBranch(stage, _process)
	}

}

func (stage *Stage) StageBranchParticipant(participant *Participant) {

	// check if instance is already staged
	if IsStaged(stage, participant) {
		return
	}

	participant.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchProcess(process *Process) {

	// check if instance is already staged
	if IsStaged(stage, process) {
		return
	}

	process.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _process := range process.SubProcesses {
		StageBranch(stage, _process)
	}
	for _, _diagramprocess := range process.DiagramProcesss {
		StageBranch(stage, _diagramprocess)
	}

}

func (stage *Stage) StageBranchProcessCompositionShape(processcompositionshape *ProcessCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, processcompositionshape) {
		return
	}

	processcompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if processcompositionshape.Process != nil {
		StageBranch(stage, processcompositionshape.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchProcessShape(processshape *ProcessShape) {

	// check if instance is already staged
	if IsStaged(stage, processshape) {
		return
	}

	processshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if processshape.Process != nil {
		StageBranch(stage, processshape.Process)
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
	case *DiagramProcess:
		toT := CopyBranchDiagramProcess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Participant:
		toT := CopyBranchParticipant(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Process:
		toT := CopyBranchProcess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProcessCompositionShape:
		toT := CopyBranchProcessCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProcessShape:
		toT := CopyBranchProcessShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchDiagramProcess(mapOrigCopy map[any]any, diagramprocessFrom *DiagramProcess) (diagramprocessTo *DiagramProcess) {

	// diagramprocessFrom has already been copied
	if _diagramprocessTo, ok := mapOrigCopy[diagramprocessFrom]; ok {
		diagramprocessTo = _diagramprocessTo.(*DiagramProcess)
		return
	}

	diagramprocessTo = new(DiagramProcess)
	mapOrigCopy[diagramprocessFrom] = diagramprocessTo
	diagramprocessFrom.CopyBasicFields(diagramprocessTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _processshape := range diagramprocessFrom.Process_Shapes {
		diagramprocessTo.Process_Shapes = append(diagramprocessTo.Process_Shapes, CopyBranchProcessShape(mapOrigCopy, _processshape))
	}
	for _, _process := range diagramprocessFrom.ProcesssWhoseNodeIsExpanded {
		diagramprocessTo.ProcesssWhoseNodeIsExpanded = append(diagramprocessTo.ProcesssWhoseNodeIsExpanded, CopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _processcompositionshape := range diagramprocessFrom.ProcessComposition_Shapes {
		diagramprocessTo.ProcessComposition_Shapes = append(diagramprocessTo.ProcessComposition_Shapes, CopyBranchProcessCompositionShape(mapOrigCopy, _processcompositionshape))
	}

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
	for _, _process := range libraryFrom.RootProcesses {
		libraryTo.RootProcesses = append(libraryTo.RootProcesses, CopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _process := range libraryFrom.ProcesssWhoseNodeIsExpanded {
		libraryTo.ProcesssWhoseNodeIsExpanded = append(libraryTo.ProcesssWhoseNodeIsExpanded, CopyBranchProcess(mapOrigCopy, _process))
	}

	return
}

func CopyBranchParticipant(mapOrigCopy map[any]any, participantFrom *Participant) (participantTo *Participant) {

	// participantFrom has already been copied
	if _participantTo, ok := mapOrigCopy[participantFrom]; ok {
		participantTo = _participantTo.(*Participant)
		return
	}

	participantTo = new(Participant)
	mapOrigCopy[participantFrom] = participantTo
	participantFrom.CopyBasicFields(participantTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchProcess(mapOrigCopy map[any]any, processFrom *Process) (processTo *Process) {

	// processFrom has already been copied
	if _processTo, ok := mapOrigCopy[processFrom]; ok {
		processTo = _processTo.(*Process)
		return
	}

	processTo = new(Process)
	mapOrigCopy[processFrom] = processTo
	processFrom.CopyBasicFields(processTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _process := range processFrom.SubProcesses {
		processTo.SubProcesses = append(processTo.SubProcesses, CopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _diagramprocess := range processFrom.DiagramProcesss {
		processTo.DiagramProcesss = append(processTo.DiagramProcesss, CopyBranchDiagramProcess(mapOrigCopy, _diagramprocess))
	}

	return
}

func CopyBranchProcessCompositionShape(mapOrigCopy map[any]any, processcompositionshapeFrom *ProcessCompositionShape) (processcompositionshapeTo *ProcessCompositionShape) {

	// processcompositionshapeFrom has already been copied
	if _processcompositionshapeTo, ok := mapOrigCopy[processcompositionshapeFrom]; ok {
		processcompositionshapeTo = _processcompositionshapeTo.(*ProcessCompositionShape)
		return
	}

	processcompositionshapeTo = new(ProcessCompositionShape)
	mapOrigCopy[processcompositionshapeFrom] = processcompositionshapeTo
	processcompositionshapeFrom.CopyBasicFields(processcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if processcompositionshapeFrom.Process != nil {
		processcompositionshapeTo.Process = CopyBranchProcess(mapOrigCopy, processcompositionshapeFrom.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchProcessShape(mapOrigCopy map[any]any, processshapeFrom *ProcessShape) (processshapeTo *ProcessShape) {

	// processshapeFrom has already been copied
	if _processshapeTo, ok := mapOrigCopy[processshapeFrom]; ok {
		processshapeTo = _processshapeTo.(*ProcessShape)
		return
	}

	processshapeTo = new(ProcessShape)
	mapOrigCopy[processshapeFrom] = processshapeTo
	processshapeFrom.CopyBasicFields(processshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if processshapeFrom.Process != nil {
		processshapeTo.Process = CopyBranchProcess(mapOrigCopy, processshapeFrom.Process)
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
	case *DiagramProcess:
		stage.UnstageBranchDiagramProcess(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Participant:
		stage.UnstageBranchParticipant(target)

	case *Process:
		stage.UnstageBranchProcess(target)

	case *ProcessCompositionShape:
		stage.UnstageBranchProcessCompositionShape(target)

	case *ProcessShape:
		stage.UnstageBranchProcessShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchDiagramProcess(diagramprocess *DiagramProcess) {

	// check if instance is already staged
	if !IsStaged(stage, diagramprocess) {
		return
	}

	diagramprocess.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _processshape := range diagramprocess.Process_Shapes {
		UnstageBranch(stage, _processshape)
	}
	for _, _process := range diagramprocess.ProcesssWhoseNodeIsExpanded {
		UnstageBranch(stage, _process)
	}
	for _, _processcompositionshape := range diagramprocess.ProcessComposition_Shapes {
		UnstageBranch(stage, _processcompositionshape)
	}

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
	for _, _process := range library.RootProcesses {
		UnstageBranch(stage, _process)
	}
	for _, _process := range library.ProcesssWhoseNodeIsExpanded {
		UnstageBranch(stage, _process)
	}

}

func (stage *Stage) UnstageBranchParticipant(participant *Participant) {

	// check if instance is already staged
	if !IsStaged(stage, participant) {
		return
	}

	participant.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchProcess(process *Process) {

	// check if instance is already staged
	if !IsStaged(stage, process) {
		return
	}

	process.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _process := range process.SubProcesses {
		UnstageBranch(stage, _process)
	}
	for _, _diagramprocess := range process.DiagramProcesss {
		UnstageBranch(stage, _diagramprocess)
	}

}

func (stage *Stage) UnstageBranchProcessCompositionShape(processcompositionshape *ProcessCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, processcompositionshape) {
		return
	}

	processcompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if processcompositionshape.Process != nil {
		UnstageBranch(stage, processcompositionshape.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchProcessShape(processshape *ProcessShape) {

	// check if instance is already staged
	if !IsStaged(stage, processshape) {
		return
	}

	processshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if processshape.Process != nil {
		UnstageBranch(stage, processshape.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *DiagramProcess) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramProcess) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Process_Shapes = reference.Process_Shapes[:0]
	for _, _b := range instance.Process_Shapes {
		reference.Process_Shapes = append(reference.Process_Shapes, stage.ProcessShapes_reference[_b])
	}
	reference.ProcesssWhoseNodeIsExpanded = reference.ProcesssWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ProcesssWhoseNodeIsExpanded {
		reference.ProcesssWhoseNodeIsExpanded = append(reference.ProcesssWhoseNodeIsExpanded, stage.Processs_reference[_b])
	}
	reference.ProcessComposition_Shapes = reference.ProcessComposition_Shapes[:0]
	for _, _b := range instance.ProcessComposition_Shapes {
		reference.ProcessComposition_Shapes = append(reference.ProcessComposition_Shapes, stage.ProcessCompositionShapes_reference[_b])
	}

	return
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.RootProcesses = reference.RootProcesses[:0]
	for _, _b := range instance.RootProcesses {
		reference.RootProcesses = append(reference.RootProcesses, stage.Processs_reference[_b])
	}
	reference.ProcesssWhoseNodeIsExpanded = reference.ProcesssWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ProcesssWhoseNodeIsExpanded {
		reference.ProcesssWhoseNodeIsExpanded = append(reference.ProcesssWhoseNodeIsExpanded, stage.Processs_reference[_b])
	}

	return
}

func (reference *Participant) GongReconstructPointersFromReferences(stage *Stage, instance *Participant) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *Process) GongReconstructPointersFromReferences(stage *Stage, instance *Process) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubProcesses = reference.SubProcesses[:0]
	for _, _b := range instance.SubProcesses {
		reference.SubProcesses = append(reference.SubProcesses, stage.Processs_reference[_b])
	}
	reference.DiagramProcesss = reference.DiagramProcesss[:0]
	for _, _b := range instance.DiagramProcesss {
		reference.DiagramProcesss = append(reference.DiagramProcesss, stage.DiagramProcesss_reference[_b])
	}

	return
}

func (reference *ProcessCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProcessCompositionShape) () {
	// insertion point for pointers field
	if instance.Process != nil {
		reference.Process = stage.Processs_reference[instance.Process]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *ProcessShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProcessShape) () {
	// insertion point for pointers field
	if instance.Process != nil {
		reference.Process = stage.Processs_reference[instance.Process]
	}
	// insertion point for slice of pointers field

	return
}

// insertion point for pointer reconstruction from instances
func (reference *DiagramProcess) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Process_Shapes []*ProcessShape
	for _, _reference := range reference.Process_Shapes {
		if _instance, ok := stage.ProcessShapes_instance[_reference]; ok {
			_Process_Shapes = append(_Process_Shapes, _instance)
		}
	}
	reference.Process_Shapes = _Process_Shapes
	var _ProcesssWhoseNodeIsExpanded []*Process
	for _, _reference := range reference.ProcesssWhoseNodeIsExpanded {
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			_ProcesssWhoseNodeIsExpanded = append(_ProcesssWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ProcesssWhoseNodeIsExpanded = _ProcesssWhoseNodeIsExpanded
	var _ProcessComposition_Shapes []*ProcessCompositionShape
	for _, _reference := range reference.ProcessComposition_Shapes {
		if _instance, ok := stage.ProcessCompositionShapes_instance[_reference]; ok {
			_ProcessComposition_Shapes = append(_ProcessComposition_Shapes, _instance)
		}
	}
	reference.ProcessComposition_Shapes = _ProcessComposition_Shapes

	return
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
	var _RootProcesses []*Process
	for _, _reference := range reference.RootProcesses {
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			_RootProcesses = append(_RootProcesses, _instance)
		}
	}
	reference.RootProcesses = _RootProcesses
	var _ProcesssWhoseNodeIsExpanded []*Process
	for _, _reference := range reference.ProcesssWhoseNodeIsExpanded {
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			_ProcesssWhoseNodeIsExpanded = append(_ProcesssWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ProcesssWhoseNodeIsExpanded = _ProcesssWhoseNodeIsExpanded

	return
}

func (reference *Participant) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *Process) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubProcesses []*Process
	for _, _reference := range reference.SubProcesses {
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			_SubProcesses = append(_SubProcesses, _instance)
		}
	}
	reference.SubProcesses = _SubProcesses
	var _DiagramProcesss []*DiagramProcess
	for _, _reference := range reference.DiagramProcesss {
		if _instance, ok := stage.DiagramProcesss_instance[_reference]; ok {
			_DiagramProcesss = append(_DiagramProcesss, _instance)
		}
	}
	reference.DiagramProcesss = _DiagramProcesss

	return
}

func (reference *ProcessCompositionShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Process; _reference != nil {
		reference.Process = nil
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			reference.Process = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *ProcessShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Process; _reference != nil {
		reference.Process = nil
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			reference.Process = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagramprocess *DiagramProcess) GongDiff(stage *Stage, diagramprocessOther *DiagramProcess) (diffs []string) {
	// insertion point for field diffs
	if diagramprocess.Name != diagramprocessOther.Name {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "Name"))
	}
	if diagramprocess.ComputedPrefix != diagramprocessOther.ComputedPrefix {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagramprocess.IsChecked != diagramprocessOther.IsChecked {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsChecked"))
	}
	if diagramprocess.IsEditable_ != diagramprocessOther.IsEditable_ {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsEditable_"))
	}
	if diagramprocess.IsShowPrefix != diagramprocessOther.IsShowPrefix {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsShowPrefix"))
	}
	if diagramprocess.DefaultBoxWidth != diagramprocessOther.DefaultBoxWidth {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "DefaultBoxWidth"))
	}
	if diagramprocess.DefaultBoxHeigth != diagramprocessOther.DefaultBoxHeigth {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "DefaultBoxHeigth"))
	}
	if diagramprocess.Width != diagramprocessOther.Width {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "Width"))
	}
	if diagramprocess.Height != diagramprocessOther.Height {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "Height"))
	}
	Process_ShapesDifferent := false
	if len(diagramprocess.Process_Shapes) != len(diagramprocessOther.Process_Shapes) {
		Process_ShapesDifferent = true
	} else {
		for i := range diagramprocess.Process_Shapes {
			if (diagramprocess.Process_Shapes[i] == nil) != (diagramprocessOther.Process_Shapes[i] == nil) {
				Process_ShapesDifferent = true
				break
			} else if diagramprocess.Process_Shapes[i] != nil && diagramprocessOther.Process_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.Process_Shapes[i] != diagramprocessOther.Process_Shapes[i] {
					Process_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Process_ShapesDifferent {
		ops := Diff(stage, diagramprocess, diagramprocessOther, "Process_Shapes", diagramprocessOther.Process_Shapes, diagramprocess.Process_Shapes)
		diffs = append(diffs, ops)
	}
	ProcesssWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.ProcesssWhoseNodeIsExpanded) != len(diagramprocessOther.ProcesssWhoseNodeIsExpanded) {
		ProcesssWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.ProcesssWhoseNodeIsExpanded {
			if (diagramprocess.ProcesssWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.ProcesssWhoseNodeIsExpanded[i] == nil) {
				ProcesssWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.ProcesssWhoseNodeIsExpanded[i] != nil && diagramprocessOther.ProcesssWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ProcesssWhoseNodeIsExpanded[i] != diagramprocessOther.ProcesssWhoseNodeIsExpanded[i] {
					ProcesssWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ProcesssWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ProcesssWhoseNodeIsExpanded", diagramprocessOther.ProcesssWhoseNodeIsExpanded, diagramprocess.ProcesssWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramprocess.IsProcesssNodeExpanded != diagramprocessOther.IsProcesssNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsProcesssNodeExpanded"))
	}
	ProcessComposition_ShapesDifferent := false
	if len(diagramprocess.ProcessComposition_Shapes) != len(diagramprocessOther.ProcessComposition_Shapes) {
		ProcessComposition_ShapesDifferent = true
	} else {
		for i := range diagramprocess.ProcessComposition_Shapes {
			if (diagramprocess.ProcessComposition_Shapes[i] == nil) != (diagramprocessOther.ProcessComposition_Shapes[i] == nil) {
				ProcessComposition_ShapesDifferent = true
				break
			} else if diagramprocess.ProcessComposition_Shapes[i] != nil && diagramprocessOther.ProcessComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ProcessComposition_Shapes[i] != diagramprocessOther.ProcessComposition_Shapes[i] {
					ProcessComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ProcessComposition_ShapesDifferent {
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ProcessComposition_Shapes", diagramprocessOther.ProcessComposition_Shapes, diagramprocess.ProcessComposition_Shapes)
		diffs = append(diffs, ops)
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
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
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
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	RootProcessesDifferent := false
	if len(library.RootProcesses) != len(libraryOther.RootProcesses) {
		RootProcessesDifferent = true
	} else {
		for i := range library.RootProcesses {
			if (library.RootProcesses[i] == nil) != (libraryOther.RootProcesses[i] == nil) {
				RootProcessesDifferent = true
				break
			} else if library.RootProcesses[i] != nil && libraryOther.RootProcesses[i] != nil {
				// this is a pointer comparaison
				if library.RootProcesses[i] != libraryOther.RootProcesses[i] {
					RootProcessesDifferent = true
					break
				}
			}
		}
	}
	if RootProcessesDifferent {
		ops := Diff(stage, library, libraryOther, "RootProcesses", libraryOther.RootProcesses, library.RootProcesses)
		diffs = append(diffs, ops)
	}
	ProcesssWhoseNodeIsExpandedDifferent := false
	if len(library.ProcesssWhoseNodeIsExpanded) != len(libraryOther.ProcesssWhoseNodeIsExpanded) {
		ProcesssWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.ProcesssWhoseNodeIsExpanded {
			if (library.ProcesssWhoseNodeIsExpanded[i] == nil) != (libraryOther.ProcesssWhoseNodeIsExpanded[i] == nil) {
				ProcesssWhoseNodeIsExpandedDifferent = true
				break
			} else if library.ProcesssWhoseNodeIsExpanded[i] != nil && libraryOther.ProcesssWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.ProcesssWhoseNodeIsExpanded[i] != libraryOther.ProcesssWhoseNodeIsExpanded[i] {
					ProcesssWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ProcesssWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "ProcesssWhoseNodeIsExpanded", libraryOther.ProcesssWhoseNodeIsExpanded, library.ProcesssWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (participant *Participant) GongDiff(stage *Stage, participantOther *Participant) (diffs []string) {
	// insertion point for field diffs
	if participant.Name != participantOther.Name {
		diffs = append(diffs, participant.GongMarshallField(stage, "Name"))
	}
	if participant.ComputedPrefix != participantOther.ComputedPrefix {
		diffs = append(diffs, participant.GongMarshallField(stage, "ComputedPrefix"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (process *Process) GongDiff(stage *Stage, processOther *Process) (diffs []string) {
	// insertion point for field diffs
	if process.Name != processOther.Name {
		diffs = append(diffs, process.GongMarshallField(stage, "Name"))
	}
	if process.ComputedPrefix != processOther.ComputedPrefix {
		diffs = append(diffs, process.GongMarshallField(stage, "ComputedPrefix"))
	}
	SubProcessesDifferent := false
	if len(process.SubProcesses) != len(processOther.SubProcesses) {
		SubProcessesDifferent = true
	} else {
		for i := range process.SubProcesses {
			if (process.SubProcesses[i] == nil) != (processOther.SubProcesses[i] == nil) {
				SubProcessesDifferent = true
				break
			} else if process.SubProcesses[i] != nil && processOther.SubProcesses[i] != nil {
				// this is a pointer comparaison
				if process.SubProcesses[i] != processOther.SubProcesses[i] {
					SubProcessesDifferent = true
					break
				}
			}
		}
	}
	if SubProcessesDifferent {
		ops := Diff(stage, process, processOther, "SubProcesses", processOther.SubProcesses, process.SubProcesses)
		diffs = append(diffs, ops)
	}
	DiagramProcesssDifferent := false
	if len(process.DiagramProcesss) != len(processOther.DiagramProcesss) {
		DiagramProcesssDifferent = true
	} else {
		for i := range process.DiagramProcesss {
			if (process.DiagramProcesss[i] == nil) != (processOther.DiagramProcesss[i] == nil) {
				DiagramProcesssDifferent = true
				break
			} else if process.DiagramProcesss[i] != nil && processOther.DiagramProcesss[i] != nil {
				// this is a pointer comparaison
				if process.DiagramProcesss[i] != processOther.DiagramProcesss[i] {
					DiagramProcesssDifferent = true
					break
				}
			}
		}
	}
	if DiagramProcesssDifferent {
		ops := Diff(stage, process, processOther, "DiagramProcesss", processOther.DiagramProcesss, process.DiagramProcesss)
		diffs = append(diffs, ops)
	}
	if process.IsSubProcessNodeExpanded != processOther.IsSubProcessNodeExpanded {
		diffs = append(diffs, process.GongMarshallField(stage, "IsSubProcessNodeExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (processcompositionshape *ProcessCompositionShape) GongDiff(stage *Stage, processcompositionshapeOther *ProcessCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if processcompositionshape.Name != processcompositionshapeOther.Name {
		diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "Name"))
	}
	if (processcompositionshape.Process == nil) != (processcompositionshapeOther.Process == nil) {
		diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "Process"))
	} else if processcompositionshape.Process != nil && processcompositionshapeOther.Process != nil {
		if processcompositionshape.Process != processcompositionshapeOther.Process {
			diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "Process"))
		}
	}
	if processcompositionshape.StartRatio != processcompositionshapeOther.StartRatio {
		diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "StartRatio"))
	}
	if processcompositionshape.EndRatio != processcompositionshapeOther.EndRatio {
		diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "EndRatio"))
	}
	if processcompositionshape.StartOrientation != processcompositionshapeOther.StartOrientation {
		diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if processcompositionshape.EndOrientation != processcompositionshapeOther.EndOrientation {
		diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if processcompositionshape.CornerOffsetRatio != processcompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if processcompositionshape.IsHidden != processcompositionshapeOther.IsHidden {
		diffs = append(diffs, processcompositionshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (processshape *ProcessShape) GongDiff(stage *Stage, processshapeOther *ProcessShape) (diffs []string) {
	// insertion point for field diffs
	if processshape.Name != processshapeOther.Name {
		diffs = append(diffs, processshape.GongMarshallField(stage, "Name"))
	}
	if (processshape.Process == nil) != (processshapeOther.Process == nil) {
		diffs = append(diffs, processshape.GongMarshallField(stage, "Process"))
	} else if processshape.Process != nil && processshapeOther.Process != nil {
		if processshape.Process != processshapeOther.Process {
			diffs = append(diffs, processshape.GongMarshallField(stage, "Process"))
		}
	}
	if processshape.IsExpanded != processshapeOther.IsExpanded {
		diffs = append(diffs, processshape.GongMarshallField(stage, "IsExpanded"))
	}
	if processshape.X != processshapeOther.X {
		diffs = append(diffs, processshape.GongMarshallField(stage, "X"))
	}
	if processshape.Y != processshapeOther.Y {
		diffs = append(diffs, processshape.GongMarshallField(stage, "Y"))
	}
	if processshape.Width != processshapeOther.Width {
		diffs = append(diffs, processshape.GongMarshallField(stage, "Width"))
	}
	if processshape.Height != processshapeOther.Height {
		diffs = append(diffs, processshape.GongMarshallField(stage, "Height"))
	}
	if processshape.IsHidden != processshapeOther.IsHidden {
		diffs = append(diffs, processshape.GongMarshallField(stage, "IsHidden"))
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
