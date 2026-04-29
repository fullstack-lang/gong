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

	case *ParticipantShape:
		ok = stage.IsStagedParticipantShape(target)

	case *Process:
		ok = stage.IsStagedProcess(target)

	case *ProcessShape:
		ok = stage.IsStagedProcessShape(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	case *TaskShape:
		ok = stage.IsStagedTaskShape(target)

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

	case *ParticipantShape:
		ok = stage.IsStagedParticipantShape(target)

	case *Process:
		ok = stage.IsStagedProcess(target)

	case *ProcessShape:
		ok = stage.IsStagedProcessShape(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	case *TaskShape:
		ok = stage.IsStagedTaskShape(target)

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

func (stage *Stage) IsStagedParticipantShape(participantshape *ParticipantShape) (ok bool) {

	_, ok = stage.ParticipantShapes[participantshape]

	return
}

func (stage *Stage) IsStagedProcess(process *Process) (ok bool) {

	_, ok = stage.Processs[process]

	return
}

func (stage *Stage) IsStagedProcessShape(processshape *ProcessShape) (ok bool) {

	_, ok = stage.ProcessShapes[processshape]

	return
}

func (stage *Stage) IsStagedTask(task *Task) (ok bool) {

	_, ok = stage.Tasks[task]

	return
}

func (stage *Stage) IsStagedTaskShape(taskshape *TaskShape) (ok bool) {

	_, ok = stage.TaskShapes[taskshape]

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

	case *ParticipantShape:
		stage.StageBranchParticipantShape(target)

	case *Process:
		stage.StageBranchProcess(target)

	case *ProcessShape:
		stage.StageBranchProcessShape(target)

	case *Task:
		stage.StageBranchTask(target)

	case *TaskShape:
		stage.StageBranchTaskShape(target)

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
	for _, _participantshape := range diagramprocess.Participant_Shapes {
		StageBranch(stage, _participantshape)
	}
	for _, _participant := range diagramprocess.ParticipantWhoseNodeIsExpanded {
		StageBranch(stage, _participant)
	}
	for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _taskshape := range diagramprocess.TaskShapes {
		StageBranch(stage, _taskshape)
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
	for _, _task := range participant.Tasks {
		StageBranch(stage, _task)
	}

}

func (stage *Stage) StageBranchParticipantShape(participantshape *ParticipantShape) {

	// check if instance is already staged
	if IsStaged(stage, participantshape) {
		return
	}

	participantshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if participantshape.Participant != nil {
		StageBranch(stage, participantshape.Participant)
	}

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
	for _, _diagramprocess := range process.DiagramProcesss {
		StageBranch(stage, _diagramprocess)
	}
	for _, _diagramprocess := range process.DiagramProcessWhoseNodeIsExpanded {
		StageBranch(stage, _diagramprocess)
	}
	for _, _process := range process.SubProcesses {
		StageBranch(stage, _process)
	}
	for _, _participant := range process.Participants {
		StageBranch(stage, _participant)
	}
	for _, _participant := range process.ParticipantWhoseNodeIsExpanded {
		StageBranch(stage, _participant)
	}

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

func (stage *Stage) StageBranchTask(task *Task) {

	// check if instance is already staged
	if IsStaged(stage, task) {
		return
	}

	task.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if IsStaged(stage, taskshape) {
		return
	}

	taskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		StageBranch(stage, taskshape.Task)
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

	case *ParticipantShape:
		toT := CopyBranchParticipantShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Process:
		toT := CopyBranchProcess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProcessShape:
		toT := CopyBranchProcessShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Task:
		toT := CopyBranchTask(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskShape:
		toT := CopyBranchTaskShape(mapOrigCopy, fromT)
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
	for _, _participantshape := range diagramprocessFrom.Participant_Shapes {
		diagramprocessTo.Participant_Shapes = append(diagramprocessTo.Participant_Shapes, CopyBranchParticipantShape(mapOrigCopy, _participantshape))
	}
	for _, _participant := range diagramprocessFrom.ParticipantWhoseNodeIsExpanded {
		diagramprocessTo.ParticipantWhoseNodeIsExpanded = append(diagramprocessTo.ParticipantWhoseNodeIsExpanded, CopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _task := range diagramprocessFrom.TasksWhoseNodeIsExpanded {
		diagramprocessTo.TasksWhoseNodeIsExpanded = append(diagramprocessTo.TasksWhoseNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskshape := range diagramprocessFrom.TaskShapes {
		diagramprocessTo.TaskShapes = append(diagramprocessTo.TaskShapes, CopyBranchTaskShape(mapOrigCopy, _taskshape))
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
	for _, _task := range participantFrom.Tasks {
		participantTo.Tasks = append(participantTo.Tasks, CopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func CopyBranchParticipantShape(mapOrigCopy map[any]any, participantshapeFrom *ParticipantShape) (participantshapeTo *ParticipantShape) {

	// participantshapeFrom has already been copied
	if _participantshapeTo, ok := mapOrigCopy[participantshapeFrom]; ok {
		participantshapeTo = _participantshapeTo.(*ParticipantShape)
		return
	}

	participantshapeTo = new(ParticipantShape)
	mapOrigCopy[participantshapeFrom] = participantshapeTo
	participantshapeFrom.CopyBasicFields(participantshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if participantshapeFrom.Participant != nil {
		participantshapeTo.Participant = CopyBranchParticipant(mapOrigCopy, participantshapeFrom.Participant)
	}

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
	for _, _diagramprocess := range processFrom.DiagramProcesss {
		processTo.DiagramProcesss = append(processTo.DiagramProcesss, CopyBranchDiagramProcess(mapOrigCopy, _diagramprocess))
	}
	for _, _diagramprocess := range processFrom.DiagramProcessWhoseNodeIsExpanded {
		processTo.DiagramProcessWhoseNodeIsExpanded = append(processTo.DiagramProcessWhoseNodeIsExpanded, CopyBranchDiagramProcess(mapOrigCopy, _diagramprocess))
	}
	for _, _process := range processFrom.SubProcesses {
		processTo.SubProcesses = append(processTo.SubProcesses, CopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _participant := range processFrom.Participants {
		processTo.Participants = append(processTo.Participants, CopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _participant := range processFrom.ParticipantWhoseNodeIsExpanded {
		processTo.ParticipantWhoseNodeIsExpanded = append(processTo.ParticipantWhoseNodeIsExpanded, CopyBranchParticipant(mapOrigCopy, _participant))
	}

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

func CopyBranchTask(mapOrigCopy map[any]any, taskFrom *Task) (taskTo *Task) {

	// taskFrom has already been copied
	if _taskTo, ok := mapOrigCopy[taskFrom]; ok {
		taskTo = _taskTo.(*Task)
		return
	}

	taskTo = new(Task)
	mapOrigCopy[taskFrom] = taskTo
	taskFrom.CopyBasicFields(taskTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTaskShape(mapOrigCopy map[any]any, taskshapeFrom *TaskShape) (taskshapeTo *TaskShape) {

	// taskshapeFrom has already been copied
	if _taskshapeTo, ok := mapOrigCopy[taskshapeFrom]; ok {
		taskshapeTo = _taskshapeTo.(*TaskShape)
		return
	}

	taskshapeTo = new(TaskShape)
	mapOrigCopy[taskshapeFrom] = taskshapeTo
	taskshapeFrom.CopyBasicFields(taskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskshapeFrom.Task != nil {
		taskshapeTo.Task = CopyBranchTask(mapOrigCopy, taskshapeFrom.Task)
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

	case *ParticipantShape:
		stage.UnstageBranchParticipantShape(target)

	case *Process:
		stage.UnstageBranchProcess(target)

	case *ProcessShape:
		stage.UnstageBranchProcessShape(target)

	case *Task:
		stage.UnstageBranchTask(target)

	case *TaskShape:
		stage.UnstageBranchTaskShape(target)

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
	for _, _participantshape := range diagramprocess.Participant_Shapes {
		UnstageBranch(stage, _participantshape)
	}
	for _, _participant := range diagramprocess.ParticipantWhoseNodeIsExpanded {
		UnstageBranch(stage, _participant)
	}
	for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _taskshape := range diagramprocess.TaskShapes {
		UnstageBranch(stage, _taskshape)
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
	for _, _task := range participant.Tasks {
		UnstageBranch(stage, _task)
	}

}

func (stage *Stage) UnstageBranchParticipantShape(participantshape *ParticipantShape) {

	// check if instance is already staged
	if !IsStaged(stage, participantshape) {
		return
	}

	participantshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if participantshape.Participant != nil {
		UnstageBranch(stage, participantshape.Participant)
	}

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
	for _, _diagramprocess := range process.DiagramProcesss {
		UnstageBranch(stage, _diagramprocess)
	}
	for _, _diagramprocess := range process.DiagramProcessWhoseNodeIsExpanded {
		UnstageBranch(stage, _diagramprocess)
	}
	for _, _process := range process.SubProcesses {
		UnstageBranch(stage, _process)
	}
	for _, _participant := range process.Participants {
		UnstageBranch(stage, _participant)
	}
	for _, _participant := range process.ParticipantWhoseNodeIsExpanded {
		UnstageBranch(stage, _participant)
	}

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

func (stage *Stage) UnstageBranchTask(task *Task) {

	// check if instance is already staged
	if !IsStaged(stage, task) {
		return
	}

	task.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskshape) {
		return
	}

	taskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		UnstageBranch(stage, taskshape.Task)
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
	reference.Participant_Shapes = reference.Participant_Shapes[:0]
	for _, _b := range instance.Participant_Shapes {
		reference.Participant_Shapes = append(reference.Participant_Shapes, stage.ParticipantShapes_reference[_b])
	}
	reference.ParticipantWhoseNodeIsExpanded = reference.ParticipantWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ParticipantWhoseNodeIsExpanded {
		reference.ParticipantWhoseNodeIsExpanded = append(reference.ParticipantWhoseNodeIsExpanded, stage.Participants_reference[_b])
	}
	reference.TasksWhoseNodeIsExpanded = reference.TasksWhoseNodeIsExpanded[:0]
	for _, _b := range instance.TasksWhoseNodeIsExpanded {
		reference.TasksWhoseNodeIsExpanded = append(reference.TasksWhoseNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.TaskShapes = reference.TaskShapes[:0]
	for _, _b := range instance.TaskShapes {
		reference.TaskShapes = append(reference.TaskShapes, stage.TaskShapes_reference[_b])
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
	reference.Tasks = reference.Tasks[:0]
	for _, _b := range instance.Tasks {
		reference.Tasks = append(reference.Tasks, stage.Tasks_reference[_b])
	}

	return
}

func (reference *ParticipantShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParticipantShape) () {
	// insertion point for pointers field
	if instance.Participant != nil {
		reference.Participant = stage.Participants_reference[instance.Participant]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Process) GongReconstructPointersFromReferences(stage *Stage, instance *Process) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.DiagramProcesss = reference.DiagramProcesss[:0]
	for _, _b := range instance.DiagramProcesss {
		reference.DiagramProcesss = append(reference.DiagramProcesss, stage.DiagramProcesss_reference[_b])
	}
	reference.DiagramProcessWhoseNodeIsExpanded = reference.DiagramProcessWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DiagramProcessWhoseNodeIsExpanded {
		reference.DiagramProcessWhoseNodeIsExpanded = append(reference.DiagramProcessWhoseNodeIsExpanded, stage.DiagramProcesss_reference[_b])
	}
	reference.SubProcesses = reference.SubProcesses[:0]
	for _, _b := range instance.SubProcesses {
		reference.SubProcesses = append(reference.SubProcesses, stage.Processs_reference[_b])
	}
	reference.Participants = reference.Participants[:0]
	for _, _b := range instance.Participants {
		reference.Participants = append(reference.Participants, stage.Participants_reference[_b])
	}
	reference.ParticipantWhoseNodeIsExpanded = reference.ParticipantWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ParticipantWhoseNodeIsExpanded {
		reference.ParticipantWhoseNodeIsExpanded = append(reference.ParticipantWhoseNodeIsExpanded, stage.Participants_reference[_b])
	}

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

func (reference *Task) GongReconstructPointersFromReferences(stage *Stage, instance *Task) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *TaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskShape) () {
	// insertion point for pointers field
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
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
	var _Participant_Shapes []*ParticipantShape
	for _, _reference := range reference.Participant_Shapes {
		if _instance, ok := stage.ParticipantShapes_instance[_reference]; ok {
			_Participant_Shapes = append(_Participant_Shapes, _instance)
		}
	}
	reference.Participant_Shapes = _Participant_Shapes
	var _ParticipantWhoseNodeIsExpanded []*Participant
	for _, _reference := range reference.ParticipantWhoseNodeIsExpanded {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_ParticipantWhoseNodeIsExpanded = append(_ParticipantWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ParticipantWhoseNodeIsExpanded = _ParticipantWhoseNodeIsExpanded
	var _TasksWhoseNodeIsExpanded []*Task
	for _, _reference := range reference.TasksWhoseNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TasksWhoseNodeIsExpanded = append(_TasksWhoseNodeIsExpanded, _instance)
		}
	}
	reference.TasksWhoseNodeIsExpanded = _TasksWhoseNodeIsExpanded
	var _TaskShapes []*TaskShape
	for _, _reference := range reference.TaskShapes {
		if _instance, ok := stage.TaskShapes_instance[_reference]; ok {
			_TaskShapes = append(_TaskShapes, _instance)
		}
	}
	reference.TaskShapes = _TaskShapes

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
	var _Tasks []*Task
	for _, _reference := range reference.Tasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_Tasks = append(_Tasks, _instance)
		}
	}
	reference.Tasks = _Tasks

	return
}

func (reference *ParticipantShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Participant; _reference != nil {
		reference.Participant = nil
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			reference.Participant = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *Process) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _DiagramProcesss []*DiagramProcess
	for _, _reference := range reference.DiagramProcesss {
		if _instance, ok := stage.DiagramProcesss_instance[_reference]; ok {
			_DiagramProcesss = append(_DiagramProcesss, _instance)
		}
	}
	reference.DiagramProcesss = _DiagramProcesss
	var _DiagramProcessWhoseNodeIsExpanded []*DiagramProcess
	for _, _reference := range reference.DiagramProcessWhoseNodeIsExpanded {
		if _instance, ok := stage.DiagramProcesss_instance[_reference]; ok {
			_DiagramProcessWhoseNodeIsExpanded = append(_DiagramProcessWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DiagramProcessWhoseNodeIsExpanded = _DiagramProcessWhoseNodeIsExpanded
	var _SubProcesses []*Process
	for _, _reference := range reference.SubProcesses {
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			_SubProcesses = append(_SubProcesses, _instance)
		}
	}
	reference.SubProcesses = _SubProcesses
	var _Participants []*Participant
	for _, _reference := range reference.Participants {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_Participants = append(_Participants, _instance)
		}
	}
	reference.Participants = _Participants
	var _ParticipantWhoseNodeIsExpanded []*Participant
	for _, _reference := range reference.ParticipantWhoseNodeIsExpanded {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_ParticipantWhoseNodeIsExpanded = append(_ParticipantWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ParticipantWhoseNodeIsExpanded = _ParticipantWhoseNodeIsExpanded

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

func (reference *Task) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *TaskShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
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
	if diagramprocess.IsProcesssNodeExpanded != diagramprocessOther.IsProcesssNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsProcesssNodeExpanded"))
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
	Participant_ShapesDifferent := false
	if len(diagramprocess.Participant_Shapes) != len(diagramprocessOther.Participant_Shapes) {
		Participant_ShapesDifferent = true
	} else {
		for i := range diagramprocess.Participant_Shapes {
			if (diagramprocess.Participant_Shapes[i] == nil) != (diagramprocessOther.Participant_Shapes[i] == nil) {
				Participant_ShapesDifferent = true
				break
			} else if diagramprocess.Participant_Shapes[i] != nil && diagramprocessOther.Participant_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.Participant_Shapes[i] != diagramprocessOther.Participant_Shapes[i] {
					Participant_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Participant_ShapesDifferent {
		ops := Diff(stage, diagramprocess, diagramprocessOther, "Participant_Shapes", diagramprocessOther.Participant_Shapes, diagramprocess.Participant_Shapes)
		diffs = append(diffs, ops)
	}
	if diagramprocess.IsParticipantsNodeExpanded != diagramprocessOther.IsParticipantsNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsParticipantsNodeExpanded"))
	}
	ParticipantWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.ParticipantWhoseNodeIsExpanded) != len(diagramprocessOther.ParticipantWhoseNodeIsExpanded) {
		ParticipantWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.ParticipantWhoseNodeIsExpanded {
			if (diagramprocess.ParticipantWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.ParticipantWhoseNodeIsExpanded[i] == nil) {
				ParticipantWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.ParticipantWhoseNodeIsExpanded[i] != nil && diagramprocessOther.ParticipantWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ParticipantWhoseNodeIsExpanded[i] != diagramprocessOther.ParticipantWhoseNodeIsExpanded[i] {
					ParticipantWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ParticipantWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ParticipantWhoseNodeIsExpanded", diagramprocessOther.ParticipantWhoseNodeIsExpanded, diagramprocess.ParticipantWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	TasksWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.TasksWhoseNodeIsExpanded) != len(diagramprocessOther.TasksWhoseNodeIsExpanded) {
		TasksWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.TasksWhoseNodeIsExpanded {
			if (diagramprocess.TasksWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.TasksWhoseNodeIsExpanded[i] == nil) {
				TasksWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.TasksWhoseNodeIsExpanded[i] != nil && diagramprocessOther.TasksWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.TasksWhoseNodeIsExpanded[i] != diagramprocessOther.TasksWhoseNodeIsExpanded[i] {
					TasksWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramprocess, diagramprocessOther, "TasksWhoseNodeIsExpanded", diagramprocessOther.TasksWhoseNodeIsExpanded, diagramprocess.TasksWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	TaskShapesDifferent := false
	if len(diagramprocess.TaskShapes) != len(diagramprocessOther.TaskShapes) {
		TaskShapesDifferent = true
	} else {
		for i := range diagramprocess.TaskShapes {
			if (diagramprocess.TaskShapes[i] == nil) != (diagramprocessOther.TaskShapes[i] == nil) {
				TaskShapesDifferent = true
				break
			} else if diagramprocess.TaskShapes[i] != nil && diagramprocessOther.TaskShapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.TaskShapes[i] != diagramprocessOther.TaskShapes[i] {
					TaskShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskShapesDifferent {
		ops := Diff(stage, diagramprocess, diagramprocessOther, "TaskShapes", diagramprocessOther.TaskShapes, diagramprocess.TaskShapes)
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
	TasksDifferent := false
	if len(participant.Tasks) != len(participantOther.Tasks) {
		TasksDifferent = true
	} else {
		for i := range participant.Tasks {
			if (participant.Tasks[i] == nil) != (participantOther.Tasks[i] == nil) {
				TasksDifferent = true
				break
			} else if participant.Tasks[i] != nil && participantOther.Tasks[i] != nil {
				// this is a pointer comparaison
				if participant.Tasks[i] != participantOther.Tasks[i] {
					TasksDifferent = true
					break
				}
			}
		}
	}
	if TasksDifferent {
		ops := Diff(stage, participant, participantOther, "Tasks", participantOther.Tasks, participant.Tasks)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (participantshape *ParticipantShape) GongDiff(stage *Stage, participantshapeOther *ParticipantShape) (diffs []string) {
	// insertion point for field diffs
	if participantshape.Name != participantshapeOther.Name {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "Name"))
	}
	if (participantshape.Participant == nil) != (participantshapeOther.Participant == nil) {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "Participant"))
	} else if participantshape.Participant != nil && participantshapeOther.Participant != nil {
		if participantshape.Participant != participantshapeOther.Participant {
			diffs = append(diffs, participantshape.GongMarshallField(stage, "Participant"))
		}
	}
	if participantshape.IsExpanded != participantshapeOther.IsExpanded {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "IsExpanded"))
	}
	if participantshape.X != participantshapeOther.X {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "X"))
	}
	if participantshape.Y != participantshapeOther.Y {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "Y"))
	}
	if participantshape.Width != participantshapeOther.Width {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "Width"))
	}
	if participantshape.Height != participantshapeOther.Height {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "Height"))
	}
	if participantshape.IsHidden != participantshapeOther.IsHidden {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "IsHidden"))
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
	DiagramProcessWhoseNodeIsExpandedDifferent := false
	if len(process.DiagramProcessWhoseNodeIsExpanded) != len(processOther.DiagramProcessWhoseNodeIsExpanded) {
		DiagramProcessWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range process.DiagramProcessWhoseNodeIsExpanded {
			if (process.DiagramProcessWhoseNodeIsExpanded[i] == nil) != (processOther.DiagramProcessWhoseNodeIsExpanded[i] == nil) {
				DiagramProcessWhoseNodeIsExpandedDifferent = true
				break
			} else if process.DiagramProcessWhoseNodeIsExpanded[i] != nil && processOther.DiagramProcessWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if process.DiagramProcessWhoseNodeIsExpanded[i] != processOther.DiagramProcessWhoseNodeIsExpanded[i] {
					DiagramProcessWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DiagramProcessWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, process, processOther, "DiagramProcessWhoseNodeIsExpanded", processOther.DiagramProcessWhoseNodeIsExpanded, process.DiagramProcessWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if process.IsSubProcessNodeExpanded != processOther.IsSubProcessNodeExpanded {
		diffs = append(diffs, process.GongMarshallField(stage, "IsSubProcessNodeExpanded"))
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
	ParticipantsDifferent := false
	if len(process.Participants) != len(processOther.Participants) {
		ParticipantsDifferent = true
	} else {
		for i := range process.Participants {
			if (process.Participants[i] == nil) != (processOther.Participants[i] == nil) {
				ParticipantsDifferent = true
				break
			} else if process.Participants[i] != nil && processOther.Participants[i] != nil {
				// this is a pointer comparaison
				if process.Participants[i] != processOther.Participants[i] {
					ParticipantsDifferent = true
					break
				}
			}
		}
	}
	if ParticipantsDifferent {
		ops := Diff(stage, process, processOther, "Participants", processOther.Participants, process.Participants)
		diffs = append(diffs, ops)
	}
	ParticipantWhoseNodeIsExpandedDifferent := false
	if len(process.ParticipantWhoseNodeIsExpanded) != len(processOther.ParticipantWhoseNodeIsExpanded) {
		ParticipantWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range process.ParticipantWhoseNodeIsExpanded {
			if (process.ParticipantWhoseNodeIsExpanded[i] == nil) != (processOther.ParticipantWhoseNodeIsExpanded[i] == nil) {
				ParticipantWhoseNodeIsExpandedDifferent = true
				break
			} else if process.ParticipantWhoseNodeIsExpanded[i] != nil && processOther.ParticipantWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if process.ParticipantWhoseNodeIsExpanded[i] != processOther.ParticipantWhoseNodeIsExpanded[i] {
					ParticipantWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ParticipantWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, process, processOther, "ParticipantWhoseNodeIsExpanded", processOther.ParticipantWhoseNodeIsExpanded, process.ParticipantWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
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

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (task *Task) GongDiff(stage *Stage, taskOther *Task) (diffs []string) {
	// insertion point for field diffs
	if task.Name != taskOther.Name {
		diffs = append(diffs, task.GongMarshallField(stage, "Name"))
	}
	if task.ComputedPrefix != taskOther.ComputedPrefix {
		diffs = append(diffs, task.GongMarshallField(stage, "ComputedPrefix"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskshape *TaskShape) GongDiff(stage *Stage, taskshapeOther *TaskShape) (diffs []string) {
	// insertion point for field diffs
	if taskshape.Name != taskshapeOther.Name {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Name"))
	}
	if (taskshape.Task == nil) != (taskshapeOther.Task == nil) {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Task"))
	} else if taskshape.Task != nil && taskshapeOther.Task != nil {
		if taskshape.Task != taskshapeOther.Task {
			diffs = append(diffs, taskshape.GongMarshallField(stage, "Task"))
		}
	}
	if taskshape.IsExpanded != taskshapeOther.IsExpanded {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "IsExpanded"))
	}
	if taskshape.X != taskshapeOther.X {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "X"))
	}
	if taskshape.Y != taskshapeOther.Y {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Y"))
	}
	if taskshape.Width != taskshapeOther.Width {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Width"))
	}
	if taskshape.Height != taskshapeOther.Height {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Height"))
	}
	if taskshape.IsHidden != taskshapeOther.IsHidden {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "IsHidden"))
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
