// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (allocatedprocessshape *AllocatedProcessShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AllocatedProcessShapes[allocatedprocessshape]

	return
}

func (stage *Stage) IsStagedAllocatedProcessShape(allocatedprocessshape *AllocatedProcessShape) (ok bool) {

	return allocatedprocessshape.GongIsStaged(stage)
}

func (allocatedresourceshape *AllocatedResourceShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AllocatedResourceShapes[allocatedresourceshape]

	return
}

func (stage *Stage) IsStagedAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) (ok bool) {

	return allocatedresourceshape.GongIsStaged(stage)
}

func (controlflow *ControlFlow) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ControlFlows[controlflow]

	return
}

func (stage *Stage) IsStagedControlFlow(controlflow *ControlFlow) (ok bool) {

	return controlflow.GongIsStaged(stage)
}

func (controlflowshape *ControlFlowShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ControlFlowShapes[controlflowshape]

	return
}

func (stage *Stage) IsStagedControlFlowShape(controlflowshape *ControlFlowShape) (ok bool) {

	return controlflowshape.GongIsStaged(stage)
}

func (data *Data) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Datas[data]

	return
}

func (stage *Stage) IsStagedData(data *Data) (ok bool) {

	return data.GongIsStaged(stage)
}

func (dataflow *DataFlow) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DataFlows[dataflow]

	return
}

func (stage *Stage) IsStagedDataFlow(dataflow *DataFlow) (ok bool) {

	return dataflow.GongIsStaged(stage)
}

func (dataflowshape *DataFlowShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DataFlowShapes[dataflowshape]

	return
}

func (stage *Stage) IsStagedDataFlowShape(dataflowshape *DataFlowShape) (ok bool) {

	return dataflowshape.GongIsStaged(stage)
}

func (datashape *DataShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DataShapes[datashape]

	return
}

func (stage *Stage) IsStagedDataShape(datashape *DataShape) (ok bool) {

	return datashape.GongIsStaged(stage)
}

func (diagramprocess *DiagramProcess) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DiagramProcesss[diagramprocess]

	return
}

func (stage *Stage) IsStagedDiagramProcess(diagramprocess *DiagramProcess) (ok bool) {

	return diagramprocess.GongIsStaged(stage)
}

func (externalparticipantshape *ExternalParticipantShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ExternalParticipantShapes[externalparticipantshape]

	return
}

func (stage *Stage) IsStagedExternalParticipantShape(externalparticipantshape *ExternalParticipantShape) (ok bool) {

	return externalparticipantshape.GongIsStaged(stage)
}

func (library *Library) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	return library.GongIsStaged(stage)
}

func (note *Note) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	return note.GongIsStaged(stage)
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	return noteshape.GongIsStaged(stage)
}

func (notetaskshape *NoteTaskShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteTaskShapes[notetaskshape]

	return
}

func (stage *Stage) IsStagedNoteTaskShape(notetaskshape *NoteTaskShape) (ok bool) {

	return notetaskshape.GongIsStaged(stage)
}

func (participant *Participant) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Participants[participant]

	return
}

func (stage *Stage) IsStagedParticipant(participant *Participant) (ok bool) {

	return participant.GongIsStaged(stage)
}

func (participantshape *ParticipantShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ParticipantShapes[participantshape]

	return
}

func (stage *Stage) IsStagedParticipantShape(participantshape *ParticipantShape) (ok bool) {

	return participantshape.GongIsStaged(stage)
}

func (process *Process) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Processs[process]

	return
}

func (stage *Stage) IsStagedProcess(process *Process) (ok bool) {

	return process.GongIsStaged(stage)
}

func (processshape *ProcessShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ProcessShapes[processshape]

	return
}

func (stage *Stage) IsStagedProcessShape(processshape *ProcessShape) (ok bool) {

	return processshape.GongIsStaged(stage)
}

func (resource *Resource) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Resources[resource]

	return
}

func (stage *Stage) IsStagedResource(resource *Resource) (ok bool) {

	return resource.GongIsStaged(stage)
}

func (task *Task) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Tasks[task]

	return
}

func (stage *Stage) IsStagedTask(task *Task) (ok bool) {

	return task.GongIsStaged(stage)
}

func (taskshape *TaskShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TaskShapes[taskshape]

	return
}

func (stage *Stage) IsStagedTaskShape(taskshape *TaskShape) (ok bool) {

	return taskshape.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (allocatedprocessshape *AllocatedProcessShape) GongStageBranch(stage *Stage) {
	stage.StageBranchAllocatedProcessShape(allocatedprocessshape)
}

func (stage *Stage) StageBranchAllocatedProcessShape(allocatedprocessshape *AllocatedProcessShape) {

	// check if instance is already staged
	if stage.IsStaged(allocatedprocessshape) {
		return
	}

	allocatedprocessshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedprocessshape.Participant != nil {
		stage.StageBranch(allocatedprocessshape.Participant)
	}
	if allocatedprocessshape.Process != nil {
		stage.StageBranch(allocatedprocessshape.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (allocatedresourceshape *AllocatedResourceShape) GongStageBranch(stage *Stage) {
	stage.StageBranchAllocatedResourceShape(allocatedresourceshape)
}

func (stage *Stage) StageBranchAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) {

	// check if instance is already staged
	if stage.IsStaged(allocatedresourceshape) {
		return
	}

	allocatedresourceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshape.Participant != nil {
		stage.StageBranch(allocatedresourceshape.Participant)
	}
	if allocatedresourceshape.Resource != nil {
		stage.StageBranch(allocatedresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlflow *ControlFlow) GongStageBranch(stage *Stage) {
	stage.StageBranchControlFlow(controlflow)
}

func (stage *Stage) StageBranchControlFlow(controlflow *ControlFlow) {

	// check if instance is already staged
	if stage.IsStaged(controlflow) {
		return
	}

	controlflow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflow.Start != nil {
		stage.StageBranch(controlflow.Start)
	}
	if controlflow.End != nil {
		stage.StageBranch(controlflow.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlflowshape *ControlFlowShape) GongStageBranch(stage *Stage) {
	stage.StageBranchControlFlowShape(controlflowshape)
}

func (stage *Stage) StageBranchControlFlowShape(controlflowshape *ControlFlowShape) {

	// check if instance is already staged
	if stage.IsStaged(controlflowshape) {
		return
	}

	controlflowshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshape.ControlFlow != nil {
		stage.StageBranch(controlflowshape.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (data *Data) GongStageBranch(stage *Stage) {
	stage.StageBranchData(data)
}

func (stage *Stage) StageBranchData(data *Data) {

	// check if instance is already staged
	if stage.IsStaged(data) {
		return
	}

	data.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dataflow *DataFlow) GongStageBranch(stage *Stage) {
	stage.StageBranchDataFlow(dataflow)
}

func (stage *Stage) StageBranchDataFlow(dataflow *DataFlow) {

	// check if instance is already staged
	if stage.IsStaged(dataflow) {
		return
	}

	dataflow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflow.StartTask != nil {
		stage.StageBranch(dataflow.StartTask)
	}
	if dataflow.EndTask != nil {
		stage.StageBranch(dataflow.EndTask)
	}
	if dataflow.StartExternalParticipant != nil {
		stage.StageBranch(dataflow.StartExternalParticipant)
	}
	if dataflow.EndExternalParticipant != nil {
		stage.StageBranch(dataflow.EndExternalParticipant)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflow.Datas {
		stage.StageBranch(_data)
	}

}

func (dataflowshape *DataFlowShape) GongStageBranch(stage *Stage) {
	stage.StageBranchDataFlowShape(dataflowshape)
}

func (stage *Stage) StageBranchDataFlowShape(dataflowshape *DataFlowShape) {

	// check if instance is already staged
	if stage.IsStaged(dataflowshape) {
		return
	}

	dataflowshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshape.DataFlow != nil {
		stage.StageBranch(dataflowshape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datashape *DataShape) GongStageBranch(stage *Stage) {
	stage.StageBranchDataShape(datashape)
}

func (stage *Stage) StageBranchDataShape(datashape *DataShape) {

	// check if instance is already staged
	if stage.IsStaged(datashape) {
		return
	}

	datashape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datashape.Data != nil {
		stage.StageBranch(datashape.Data)
	}
	if datashape.DataFlow != nil {
		stage.StageBranch(datashape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagramprocess *DiagramProcess) GongStageBranch(stage *Stage) {
	stage.StageBranchDiagramProcess(diagramprocess)
}

func (stage *Stage) StageBranchDiagramProcess(diagramprocess *DiagramProcess) {

	// check if instance is already staged
	if stage.IsStaged(diagramprocess) {
		return
	}

	diagramprocess.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _processshape := range diagramprocess.Process_Shapes {
		stage.StageBranch(_processshape)
	}
	for _, _process := range diagramprocess.ProcesssWhoseNodeIsExpanded {
		stage.StageBranch(_process)
	}
	for _, _participantshape := range diagramprocess.Participant_Shapes {
		stage.StageBranch(_participantshape)
	}
	for _, _participant := range diagramprocess.ParticipantWhoseNodeIsExpanded {
		stage.StageBranch(_participant)
	}
	for _, _externalparticipantshape := range diagramprocess.ExternalParticipant_Shapes {
		stage.StageBranch(_externalparticipantshape)
	}
	for _, _participant := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
		stage.StageBranch(_participant)
	}
	for _, _participant := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
		stage.StageBranch(_participant)
	}
	for _, _participant := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
		stage.StageBranch(_participant)
	}
	for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
		stage.StageBranch(_task)
	}
	for _, _taskshape := range diagramprocess.Task_Shapes {
		stage.StageBranch(_taskshape)
	}
	for _, _controlflow := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
		stage.StageBranch(_controlflow)
	}
	for _, _controlflowshape := range diagramprocess.ControlFlow_Shapes {
		stage.StageBranch(_controlflowshape)
	}
	for _, _dataflow := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
		stage.StageBranch(_dataflow)
	}
	for _, _dataflowshape := range diagramprocess.DataFlow_Shapes {
		stage.StageBranch(_dataflowshape)
	}
	for _, _data := range diagramprocess.DatasWhoseNodeIsExpanded {
		stage.StageBranch(_data)
	}
	for _, _datashape := range diagramprocess.Data_Shapes {
		stage.StageBranch(_datashape)
	}
	for _, _dataflow := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
		stage.StageBranch(_dataflow)
	}
	for _, _resource := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
		stage.StageBranch(_resource)
	}
	for _, _allocatedresourceshape := range diagramprocess.AllocatedResourceShapes {
		stage.StageBranch(_allocatedresourceshape)
	}
	for _, _process := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
		stage.StageBranch(_process)
	}
	for _, _allocatedprocessshape := range diagramprocess.AllocatedProcessShapes {
		stage.StageBranch(_allocatedprocessshape)
	}
	for _, _noteshape := range diagramprocess.Note_Shapes {
		stage.StageBranch(_noteshape)
	}
	for _, _note := range diagramprocess.NotesWhoseNodeIsExpanded {
		stage.StageBranch(_note)
	}
	for _, _notetaskshape := range diagramprocess.NoteTaskShapes {
		stage.StageBranch(_notetaskshape)
	}

}

func (externalparticipantshape *ExternalParticipantShape) GongStageBranch(stage *Stage) {
	stage.StageBranchExternalParticipantShape(externalparticipantshape)
}

func (stage *Stage) StageBranchExternalParticipantShape(externalparticipantshape *ExternalParticipantShape) {

	// check if instance is already staged
	if stage.IsStaged(externalparticipantshape) {
		return
	}

	externalparticipantshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if externalparticipantshape.Participant != nil {
		stage.StageBranch(externalparticipantshape.Participant)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongStageBranch(stage *Stage) {
	stage.StageBranchLibrary(library)
}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if stage.IsStaged(library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.StageBranch(_library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.StageBranch(_library)
	}
	for _, _process := range library.RootProcesses {
		stage.StageBranch(_process)
	}
	for _, _process := range library.ProcesssWhoseNodeIsExpanded {
		stage.StageBranch(_process)
	}
	for _, _dataflow := range library.RootDataFlows {
		stage.StageBranch(_dataflow)
	}
	for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
		stage.StageBranch(_dataflow)
	}
	for _, _data := range library.RootDatas {
		stage.StageBranch(_data)
	}
	for _, _data := range library.DatasWhoseNodeIsExpanded {
		stage.StageBranch(_data)
	}
	for _, _resource := range library.RootResources {
		stage.StageBranch(_resource)
	}
	for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
		stage.StageBranch(_resource)
	}
	for _, _participant := range library.ParticipantsWhoseNodeIsExpanded {
		stage.StageBranch(_participant)
	}
	for _, _note := range library.RootNotes {
		stage.StageBranch(_note)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		stage.StageBranch(_note)
	}

}

func (note *Note) GongStageBranch(stage *Stage) {
	stage.StageBranchNote(note)
}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range note.Tasks {
		stage.StageBranch(_task)
	}

}

func (noteshape *NoteShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteShape(noteshape)
}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if stage.IsStaged(noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.StageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notetaskshape *NoteTaskShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteTaskShape(notetaskshape)
}

func (stage *Stage) StageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if stage.IsStaged(notetaskshape) {
		return
	}

	notetaskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		stage.StageBranch(notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		stage.StageBranch(notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (participant *Participant) GongStageBranch(stage *Stage) {
	stage.StageBranchParticipant(participant)
}

func (stage *Stage) StageBranchParticipant(participant *Participant) {

	// check if instance is already staged
	if stage.IsStaged(participant) {
		return
	}

	participant.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _resource := range participant.Resources {
		stage.StageBranch(_resource)
	}
	for _, _process := range participant.Processes {
		stage.StageBranch(_process)
	}
	for _, _task := range participant.Tasks {
		stage.StageBranch(_task)
	}
	for _, _controlflow := range participant.ControlFlows {
		stage.StageBranch(_controlflow)
	}
	for _, _task := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
		stage.StageBranch(_task)
	}
	for _, _task := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
		stage.StageBranch(_task)
	}
	for _, _task := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
		stage.StageBranch(_task)
	}
	for _, _task := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
		stage.StageBranch(_task)
	}

}

func (participantshape *ParticipantShape) GongStageBranch(stage *Stage) {
	stage.StageBranchParticipantShape(participantshape)
}

func (stage *Stage) StageBranchParticipantShape(participantshape *ParticipantShape) {

	// check if instance is already staged
	if stage.IsStaged(participantshape) {
		return
	}

	participantshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if participantshape.Participant != nil {
		stage.StageBranch(participantshape.Participant)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (process *Process) GongStageBranch(stage *Stage) {
	stage.StageBranchProcess(process)
}

func (stage *Stage) StageBranchProcess(process *Process) {

	// check if instance is already staged
	if stage.IsStaged(process) {
		return
	}

	process.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramprocess := range process.DiagramProcesss {
		stage.StageBranch(_diagramprocess)
	}
	for _, _diagramprocess := range process.DiagramProcessWhoseNodeIsExpanded {
		stage.StageBranch(_diagramprocess)
	}
	for _, _process := range process.SubProcesses {
		stage.StageBranch(_process)
	}
	for _, _participant := range process.Participants {
		stage.StageBranch(_participant)
	}
	for _, _participant := range process.ParticipantWhoseNodeIsExpanded {
		stage.StageBranch(_participant)
	}
	for _, _dataflow := range process.DataFlows {
		stage.StageBranch(_dataflow)
	}
	for _, _participant := range process.ExternalParticipants {
		stage.StageBranch(_participant)
	}
	for _, _participant := range process.ExternalParticipantWhoseNodeIsExpanded {
		stage.StageBranch(_participant)
	}

}

func (processshape *ProcessShape) GongStageBranch(stage *Stage) {
	stage.StageBranchProcessShape(processshape)
}

func (stage *Stage) StageBranchProcessShape(processshape *ProcessShape) {

	// check if instance is already staged
	if stage.IsStaged(processshape) {
		return
	}

	processshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if processshape.Process != nil {
		stage.StageBranch(processshape.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resource *Resource) GongStageBranch(stage *Stage) {
	stage.StageBranchResource(resource)
}

func (stage *Stage) StageBranchResource(resource *Resource) {

	// check if instance is already staged
	if stage.IsStaged(resource) {
		return
	}

	resource.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (task *Task) GongStageBranch(stage *Stage) {
	stage.StageBranchTask(task)
}

func (stage *Stage) StageBranchTask(task *Task) {

	// check if instance is already staged
	if stage.IsStaged(task) {
		return
	}

	task.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if task.Type != nil {
		stage.StageBranch(task.Type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskshape *TaskShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTaskShape(taskshape)
}

func (stage *Stage) StageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if stage.IsStaged(taskshape) {
		return
	}

	taskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		stage.StageBranch(taskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AllocatedProcessShape:
		toT := GongCopyBranchAllocatedProcessShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AllocatedResourceShape:
		toT := GongCopyBranchAllocatedResourceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlFlow:
		toT := GongCopyBranchControlFlow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlFlowShape:
		toT := GongCopyBranchControlFlowShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Data:
		toT := GongCopyBranchData(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataFlow:
		toT := GongCopyBranchDataFlow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataFlowShape:
		toT := GongCopyBranchDataFlowShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataShape:
		toT := GongCopyBranchDataShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramProcess:
		toT := GongCopyBranchDiagramProcess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExternalParticipantShape:
		toT := GongCopyBranchExternalParticipantShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := GongCopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteTaskShape:
		toT := GongCopyBranchNoteTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Participant:
		toT := GongCopyBranchParticipant(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ParticipantShape:
		toT := GongCopyBranchParticipantShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Process:
		toT := GongCopyBranchProcess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProcessShape:
		toT := GongCopyBranchProcessShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Resource:
		toT := GongCopyBranchResource(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Task:
		toT := GongCopyBranchTask(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskShape:
		toT := GongCopyBranchTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAllocatedProcessShape(mapOrigCopy map[any]any, allocatedprocessshapeFrom *AllocatedProcessShape) (allocatedprocessshapeTo *AllocatedProcessShape) {

	// allocatedprocessshapeFrom has already been copied
	if _allocatedprocessshapeTo, ok := mapOrigCopy[allocatedprocessshapeFrom]; ok {
		allocatedprocessshapeTo = _allocatedprocessshapeTo.(*AllocatedProcessShape)
		return
	}

	allocatedprocessshapeTo = new(AllocatedProcessShape)
	mapOrigCopy[allocatedprocessshapeFrom] = allocatedprocessshapeTo
	allocatedprocessshapeFrom.GongCopyBasicFields(allocatedprocessshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if allocatedprocessshapeFrom.Participant != nil {
		allocatedprocessshapeTo.Participant = GongCopyBranchParticipant(mapOrigCopy, allocatedprocessshapeFrom.Participant)
	}
	if allocatedprocessshapeFrom.Process != nil {
		allocatedprocessshapeTo.Process = GongCopyBranchProcess(mapOrigCopy, allocatedprocessshapeFrom.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAllocatedResourceShape(mapOrigCopy map[any]any, allocatedresourceshapeFrom *AllocatedResourceShape) (allocatedresourceshapeTo *AllocatedResourceShape) {

	// allocatedresourceshapeFrom has already been copied
	if _allocatedresourceshapeTo, ok := mapOrigCopy[allocatedresourceshapeFrom]; ok {
		allocatedresourceshapeTo = _allocatedresourceshapeTo.(*AllocatedResourceShape)
		return
	}

	allocatedresourceshapeTo = new(AllocatedResourceShape)
	mapOrigCopy[allocatedresourceshapeFrom] = allocatedresourceshapeTo
	allocatedresourceshapeFrom.GongCopyBasicFields(allocatedresourceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshapeFrom.Participant != nil {
		allocatedresourceshapeTo.Participant = GongCopyBranchParticipant(mapOrigCopy, allocatedresourceshapeFrom.Participant)
	}
	if allocatedresourceshapeFrom.Resource != nil {
		allocatedresourceshapeTo.Resource = GongCopyBranchResource(mapOrigCopy, allocatedresourceshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlFlow(mapOrigCopy map[any]any, controlflowFrom *ControlFlow) (controlflowTo *ControlFlow) {

	// controlflowFrom has already been copied
	if _controlflowTo, ok := mapOrigCopy[controlflowFrom]; ok {
		controlflowTo = _controlflowTo.(*ControlFlow)
		return
	}

	controlflowTo = new(ControlFlow)
	mapOrigCopy[controlflowFrom] = controlflowTo
	controlflowFrom.GongCopyBasicFields(controlflowTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowFrom.Start != nil {
		controlflowTo.Start = GongCopyBranchTask(mapOrigCopy, controlflowFrom.Start)
	}
	if controlflowFrom.End != nil {
		controlflowTo.End = GongCopyBranchTask(mapOrigCopy, controlflowFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlFlowShape(mapOrigCopy map[any]any, controlflowshapeFrom *ControlFlowShape) (controlflowshapeTo *ControlFlowShape) {

	// controlflowshapeFrom has already been copied
	if _controlflowshapeTo, ok := mapOrigCopy[controlflowshapeFrom]; ok {
		controlflowshapeTo = _controlflowshapeTo.(*ControlFlowShape)
		return
	}

	controlflowshapeTo = new(ControlFlowShape)
	mapOrigCopy[controlflowshapeFrom] = controlflowshapeTo
	controlflowshapeFrom.GongCopyBasicFields(controlflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshapeFrom.ControlFlow != nil {
		controlflowshapeTo.ControlFlow = GongCopyBranchControlFlow(mapOrigCopy, controlflowshapeFrom.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchData(mapOrigCopy map[any]any, dataFrom *Data) (dataTo *Data) {

	// dataFrom has already been copied
	if _dataTo, ok := mapOrigCopy[dataFrom]; ok {
		dataTo = _dataTo.(*Data)
		return
	}

	dataTo = new(Data)
	mapOrigCopy[dataFrom] = dataTo
	dataFrom.GongCopyBasicFields(dataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDataFlow(mapOrigCopy map[any]any, dataflowFrom *DataFlow) (dataflowTo *DataFlow) {

	// dataflowFrom has already been copied
	if _dataflowTo, ok := mapOrigCopy[dataflowFrom]; ok {
		dataflowTo = _dataflowTo.(*DataFlow)
		return
	}

	dataflowTo = new(DataFlow)
	mapOrigCopy[dataflowFrom] = dataflowTo
	dataflowFrom.GongCopyBasicFields(dataflowTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowFrom.StartTask != nil {
		dataflowTo.StartTask = GongCopyBranchTask(mapOrigCopy, dataflowFrom.StartTask)
	}
	if dataflowFrom.EndTask != nil {
		dataflowTo.EndTask = GongCopyBranchTask(mapOrigCopy, dataflowFrom.EndTask)
	}
	if dataflowFrom.StartExternalParticipant != nil {
		dataflowTo.StartExternalParticipant = GongCopyBranchParticipant(mapOrigCopy, dataflowFrom.StartExternalParticipant)
	}
	if dataflowFrom.EndExternalParticipant != nil {
		dataflowTo.EndExternalParticipant = GongCopyBranchParticipant(mapOrigCopy, dataflowFrom.EndExternalParticipant)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflowFrom.Datas {
		dataflowTo.Datas = append(dataflowTo.Datas, GongCopyBranchData(mapOrigCopy, _data))
	}

	return
}

func GongCopyBranchDataFlowShape(mapOrigCopy map[any]any, dataflowshapeFrom *DataFlowShape) (dataflowshapeTo *DataFlowShape) {

	// dataflowshapeFrom has already been copied
	if _dataflowshapeTo, ok := mapOrigCopy[dataflowshapeFrom]; ok {
		dataflowshapeTo = _dataflowshapeTo.(*DataFlowShape)
		return
	}

	dataflowshapeTo = new(DataFlowShape)
	mapOrigCopy[dataflowshapeFrom] = dataflowshapeTo
	dataflowshapeFrom.GongCopyBasicFields(dataflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshapeFrom.DataFlow != nil {
		dataflowshapeTo.DataFlow = GongCopyBranchDataFlow(mapOrigCopy, dataflowshapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDataShape(mapOrigCopy map[any]any, datashapeFrom *DataShape) (datashapeTo *DataShape) {

	// datashapeFrom has already been copied
	if _datashapeTo, ok := mapOrigCopy[datashapeFrom]; ok {
		datashapeTo = _datashapeTo.(*DataShape)
		return
	}

	datashapeTo = new(DataShape)
	mapOrigCopy[datashapeFrom] = datashapeTo
	datashapeFrom.GongCopyBasicFields(datashapeTo)

	//insertion point for the staging of instances referenced by pointers
	if datashapeFrom.Data != nil {
		datashapeTo.Data = GongCopyBranchData(mapOrigCopy, datashapeFrom.Data)
	}
	if datashapeFrom.DataFlow != nil {
		datashapeTo.DataFlow = GongCopyBranchDataFlow(mapOrigCopy, datashapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagramProcess(mapOrigCopy map[any]any, diagramprocessFrom *DiagramProcess) (diagramprocessTo *DiagramProcess) {

	// diagramprocessFrom has already been copied
	if _diagramprocessTo, ok := mapOrigCopy[diagramprocessFrom]; ok {
		diagramprocessTo = _diagramprocessTo.(*DiagramProcess)
		return
	}

	diagramprocessTo = new(DiagramProcess)
	mapOrigCopy[diagramprocessFrom] = diagramprocessTo
	diagramprocessFrom.GongCopyBasicFields(diagramprocessTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _processshape := range diagramprocessFrom.Process_Shapes {
		diagramprocessTo.Process_Shapes = append(diagramprocessTo.Process_Shapes, GongCopyBranchProcessShape(mapOrigCopy, _processshape))
	}
	for _, _process := range diagramprocessFrom.ProcesssWhoseNodeIsExpanded {
		diagramprocessTo.ProcesssWhoseNodeIsExpanded = append(diagramprocessTo.ProcesssWhoseNodeIsExpanded, GongCopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _participantshape := range diagramprocessFrom.Participant_Shapes {
		diagramprocessTo.Participant_Shapes = append(diagramprocessTo.Participant_Shapes, GongCopyBranchParticipantShape(mapOrigCopy, _participantshape))
	}
	for _, _participant := range diagramprocessFrom.ParticipantWhoseNodeIsExpanded {
		diagramprocessTo.ParticipantWhoseNodeIsExpanded = append(diagramprocessTo.ParticipantWhoseNodeIsExpanded, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _externalparticipantshape := range diagramprocessFrom.ExternalParticipant_Shapes {
		diagramprocessTo.ExternalParticipant_Shapes = append(diagramprocessTo.ExternalParticipant_Shapes, GongCopyBranchExternalParticipantShape(mapOrigCopy, _externalparticipantshape))
	}
	for _, _participant := range diagramprocessFrom.ExternalParticipantWhoseNodeIsExpanded {
		diagramprocessTo.ExternalParticipantWhoseNodeIsExpanded = append(diagramprocessTo.ExternalParticipantWhoseNodeIsExpanded, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _participant := range diagramprocessFrom.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
		diagramprocessTo.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(diagramprocessTo.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _participant := range diagramprocessFrom.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
		diagramprocessTo.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(diagramprocessTo.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _task := range diagramprocessFrom.TasksWhoseNodeIsExpanded {
		diagramprocessTo.TasksWhoseNodeIsExpanded = append(diagramprocessTo.TasksWhoseNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskshape := range diagramprocessFrom.Task_Shapes {
		diagramprocessTo.Task_Shapes = append(diagramprocessTo.Task_Shapes, GongCopyBranchTaskShape(mapOrigCopy, _taskshape))
	}
	for _, _controlflow := range diagramprocessFrom.ControlFlowsWhoseNodeIsExpanded {
		diagramprocessTo.ControlFlowsWhoseNodeIsExpanded = append(diagramprocessTo.ControlFlowsWhoseNodeIsExpanded, GongCopyBranchControlFlow(mapOrigCopy, _controlflow))
	}
	for _, _controlflowshape := range diagramprocessFrom.ControlFlow_Shapes {
		diagramprocessTo.ControlFlow_Shapes = append(diagramprocessTo.ControlFlow_Shapes, GongCopyBranchControlFlowShape(mapOrigCopy, _controlflowshape))
	}
	for _, _dataflow := range diagramprocessFrom.DataFlowsWhoseNodeIsExpanded {
		diagramprocessTo.DataFlowsWhoseNodeIsExpanded = append(diagramprocessTo.DataFlowsWhoseNodeIsExpanded, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _dataflowshape := range diagramprocessFrom.DataFlow_Shapes {
		diagramprocessTo.DataFlow_Shapes = append(diagramprocessTo.DataFlow_Shapes, GongCopyBranchDataFlowShape(mapOrigCopy, _dataflowshape))
	}
	for _, _data := range diagramprocessFrom.DatasWhoseNodeIsExpanded {
		diagramprocessTo.DatasWhoseNodeIsExpanded = append(diagramprocessTo.DatasWhoseNodeIsExpanded, GongCopyBranchData(mapOrigCopy, _data))
	}
	for _, _datashape := range diagramprocessFrom.Data_Shapes {
		diagramprocessTo.Data_Shapes = append(diagramprocessTo.Data_Shapes, GongCopyBranchDataShape(mapOrigCopy, _datashape))
	}
	for _, _dataflow := range diagramprocessFrom.DataFlowsWhoseDataNodeIsExpanded {
		diagramprocessTo.DataFlowsWhoseDataNodeIsExpanded = append(diagramprocessTo.DataFlowsWhoseDataNodeIsExpanded, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _resource := range diagramprocessFrom.AllocatedResourcesWhoseNodeIsExpanded {
		diagramprocessTo.AllocatedResourcesWhoseNodeIsExpanded = append(diagramprocessTo.AllocatedResourcesWhoseNodeIsExpanded, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _allocatedresourceshape := range diagramprocessFrom.AllocatedResourceShapes {
		diagramprocessTo.AllocatedResourceShapes = append(diagramprocessTo.AllocatedResourceShapes, GongCopyBranchAllocatedResourceShape(mapOrigCopy, _allocatedresourceshape))
	}
	for _, _process := range diagramprocessFrom.AllocatedProcessesWhoseNodeIsExpanded {
		diagramprocessTo.AllocatedProcessesWhoseNodeIsExpanded = append(diagramprocessTo.AllocatedProcessesWhoseNodeIsExpanded, GongCopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _allocatedprocessshape := range diagramprocessFrom.AllocatedProcessShapes {
		diagramprocessTo.AllocatedProcessShapes = append(diagramprocessTo.AllocatedProcessShapes, GongCopyBranchAllocatedProcessShape(mapOrigCopy, _allocatedprocessshape))
	}
	for _, _noteshape := range diagramprocessFrom.Note_Shapes {
		diagramprocessTo.Note_Shapes = append(diagramprocessTo.Note_Shapes, GongCopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramprocessFrom.NotesWhoseNodeIsExpanded {
		diagramprocessTo.NotesWhoseNodeIsExpanded = append(diagramprocessTo.NotesWhoseNodeIsExpanded, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _notetaskshape := range diagramprocessFrom.NoteTaskShapes {
		diagramprocessTo.NoteTaskShapes = append(diagramprocessTo.NoteTaskShapes, GongCopyBranchNoteTaskShape(mapOrigCopy, _notetaskshape))
	}

	return
}

func GongCopyBranchExternalParticipantShape(mapOrigCopy map[any]any, externalparticipantshapeFrom *ExternalParticipantShape) (externalparticipantshapeTo *ExternalParticipantShape) {

	// externalparticipantshapeFrom has already been copied
	if _externalparticipantshapeTo, ok := mapOrigCopy[externalparticipantshapeFrom]; ok {
		externalparticipantshapeTo = _externalparticipantshapeTo.(*ExternalParticipantShape)
		return
	}

	externalparticipantshapeTo = new(ExternalParticipantShape)
	mapOrigCopy[externalparticipantshapeFrom] = externalparticipantshapeTo
	externalparticipantshapeFrom.GongCopyBasicFields(externalparticipantshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if externalparticipantshapeFrom.Participant != nil {
		externalparticipantshapeTo.Participant = GongCopyBranchParticipant(mapOrigCopy, externalparticipantshapeFrom.Participant)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.GongCopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _process := range libraryFrom.RootProcesses {
		libraryTo.RootProcesses = append(libraryTo.RootProcesses, GongCopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _process := range libraryFrom.ProcesssWhoseNodeIsExpanded {
		libraryTo.ProcesssWhoseNodeIsExpanded = append(libraryTo.ProcesssWhoseNodeIsExpanded, GongCopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _dataflow := range libraryFrom.RootDataFlows {
		libraryTo.RootDataFlows = append(libraryTo.RootDataFlows, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _dataflow := range libraryFrom.DataFlowsWhoseNodeIsExpanded {
		libraryTo.DataFlowsWhoseNodeIsExpanded = append(libraryTo.DataFlowsWhoseNodeIsExpanded, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _data := range libraryFrom.RootDatas {
		libraryTo.RootDatas = append(libraryTo.RootDatas, GongCopyBranchData(mapOrigCopy, _data))
	}
	for _, _data := range libraryFrom.DatasWhoseNodeIsExpanded {
		libraryTo.DatasWhoseNodeIsExpanded = append(libraryTo.DatasWhoseNodeIsExpanded, GongCopyBranchData(mapOrigCopy, _data))
	}
	for _, _resource := range libraryFrom.RootResources {
		libraryTo.RootResources = append(libraryTo.RootResources, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _resource := range libraryFrom.ResourcesWhoseNodeIsExpanded {
		libraryTo.ResourcesWhoseNodeIsExpanded = append(libraryTo.ResourcesWhoseNodeIsExpanded, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _participant := range libraryFrom.ParticipantsWhoseNodeIsExpanded {
		libraryTo.ParticipantsWhoseNodeIsExpanded = append(libraryTo.ParticipantsWhoseNodeIsExpanded, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _note := range libraryFrom.RootNotes {
		libraryTo.RootNotes = append(libraryTo.RootNotes, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _note := range libraryFrom.NotesWhoseNodeIsExpanded {
		libraryTo.NotesWhoseNodeIsExpanded = append(libraryTo.NotesWhoseNodeIsExpanded, GongCopyBranchNote(mapOrigCopy, _note))
	}

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range noteFrom.Tasks {
		noteTo.Tasks = append(noteTo.Tasks, GongCopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func GongCopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteTaskShape(mapOrigCopy map[any]any, notetaskshapeFrom *NoteTaskShape) (notetaskshapeTo *NoteTaskShape) {

	// notetaskshapeFrom has already been copied
	if _notetaskshapeTo, ok := mapOrigCopy[notetaskshapeFrom]; ok {
		notetaskshapeTo = _notetaskshapeTo.(*NoteTaskShape)
		return
	}

	notetaskshapeTo = new(NoteTaskShape)
	mapOrigCopy[notetaskshapeFrom] = notetaskshapeTo
	notetaskshapeFrom.GongCopyBasicFields(notetaskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshapeFrom.Note != nil {
		notetaskshapeTo.Note = GongCopyBranchNote(mapOrigCopy, notetaskshapeFrom.Note)
	}
	if notetaskshapeFrom.Task != nil {
		notetaskshapeTo.Task = GongCopyBranchTask(mapOrigCopy, notetaskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchParticipant(mapOrigCopy map[any]any, participantFrom *Participant) (participantTo *Participant) {

	// participantFrom has already been copied
	if _participantTo, ok := mapOrigCopy[participantFrom]; ok {
		participantTo = _participantTo.(*Participant)
		return
	}

	participantTo = new(Participant)
	mapOrigCopy[participantFrom] = participantTo
	participantFrom.GongCopyBasicFields(participantTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _resource := range participantFrom.Resources {
		participantTo.Resources = append(participantTo.Resources, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _process := range participantFrom.Processes {
		participantTo.Processes = append(participantTo.Processes, GongCopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _task := range participantFrom.Tasks {
		participantTo.Tasks = append(participantTo.Tasks, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _controlflow := range participantFrom.ControlFlows {
		participantTo.ControlFlows = append(participantTo.ControlFlows, GongCopyBranchControlFlow(mapOrigCopy, _controlflow))
	}
	for _, _task := range participantFrom.TaskWhoseOutControlFlowsNodeIsExpanded {
		participantTo.TaskWhoseOutControlFlowsNodeIsExpanded = append(participantTo.TaskWhoseOutControlFlowsNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range participantFrom.TaskWhoseInControlFlowsNodeIsExpanded {
		participantTo.TaskWhoseInControlFlowsNodeIsExpanded = append(participantTo.TaskWhoseInControlFlowsNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range participantFrom.TaskWhoseOutDataFlowsNodeIsExpanded {
		participantTo.TaskWhoseOutDataFlowsNodeIsExpanded = append(participantTo.TaskWhoseOutDataFlowsNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range participantFrom.TaskWhoseInDataFlowsNodeIsExpanded {
		participantTo.TaskWhoseInDataFlowsNodeIsExpanded = append(participantTo.TaskWhoseInDataFlowsNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func GongCopyBranchParticipantShape(mapOrigCopy map[any]any, participantshapeFrom *ParticipantShape) (participantshapeTo *ParticipantShape) {

	// participantshapeFrom has already been copied
	if _participantshapeTo, ok := mapOrigCopy[participantshapeFrom]; ok {
		participantshapeTo = _participantshapeTo.(*ParticipantShape)
		return
	}

	participantshapeTo = new(ParticipantShape)
	mapOrigCopy[participantshapeFrom] = participantshapeTo
	participantshapeFrom.GongCopyBasicFields(participantshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if participantshapeFrom.Participant != nil {
		participantshapeTo.Participant = GongCopyBranchParticipant(mapOrigCopy, participantshapeFrom.Participant)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchProcess(mapOrigCopy map[any]any, processFrom *Process) (processTo *Process) {

	// processFrom has already been copied
	if _processTo, ok := mapOrigCopy[processFrom]; ok {
		processTo = _processTo.(*Process)
		return
	}

	processTo = new(Process)
	mapOrigCopy[processFrom] = processTo
	processFrom.GongCopyBasicFields(processTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramprocess := range processFrom.DiagramProcesss {
		processTo.DiagramProcesss = append(processTo.DiagramProcesss, GongCopyBranchDiagramProcess(mapOrigCopy, _diagramprocess))
	}
	for _, _diagramprocess := range processFrom.DiagramProcessWhoseNodeIsExpanded {
		processTo.DiagramProcessWhoseNodeIsExpanded = append(processTo.DiagramProcessWhoseNodeIsExpanded, GongCopyBranchDiagramProcess(mapOrigCopy, _diagramprocess))
	}
	for _, _process := range processFrom.SubProcesses {
		processTo.SubProcesses = append(processTo.SubProcesses, GongCopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _participant := range processFrom.Participants {
		processTo.Participants = append(processTo.Participants, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _participant := range processFrom.ParticipantWhoseNodeIsExpanded {
		processTo.ParticipantWhoseNodeIsExpanded = append(processTo.ParticipantWhoseNodeIsExpanded, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _dataflow := range processFrom.DataFlows {
		processTo.DataFlows = append(processTo.DataFlows, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _participant := range processFrom.ExternalParticipants {
		processTo.ExternalParticipants = append(processTo.ExternalParticipants, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _participant := range processFrom.ExternalParticipantWhoseNodeIsExpanded {
		processTo.ExternalParticipantWhoseNodeIsExpanded = append(processTo.ExternalParticipantWhoseNodeIsExpanded, GongCopyBranchParticipant(mapOrigCopy, _participant))
	}

	return
}

func GongCopyBranchProcessShape(mapOrigCopy map[any]any, processshapeFrom *ProcessShape) (processshapeTo *ProcessShape) {

	// processshapeFrom has already been copied
	if _processshapeTo, ok := mapOrigCopy[processshapeFrom]; ok {
		processshapeTo = _processshapeTo.(*ProcessShape)
		return
	}

	processshapeTo = new(ProcessShape)
	mapOrigCopy[processshapeFrom] = processshapeTo
	processshapeFrom.GongCopyBasicFields(processshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if processshapeFrom.Process != nil {
		processshapeTo.Process = GongCopyBranchProcess(mapOrigCopy, processshapeFrom.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {

	// resourceFrom has already been copied
	if _resourceTo, ok := mapOrigCopy[resourceFrom]; ok {
		resourceTo = _resourceTo.(*Resource)
		return
	}

	resourceTo = new(Resource)
	mapOrigCopy[resourceFrom] = resourceTo
	resourceFrom.GongCopyBasicFields(resourceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTask(mapOrigCopy map[any]any, taskFrom *Task) (taskTo *Task) {

	// taskFrom has already been copied
	if _taskTo, ok := mapOrigCopy[taskFrom]; ok {
		taskTo = _taskTo.(*Task)
		return
	}

	taskTo = new(Task)
	mapOrigCopy[taskFrom] = taskTo
	taskFrom.GongCopyBasicFields(taskTo)

	//insertion point for the staging of instances referenced by pointers
	if taskFrom.Type != nil {
		taskTo.Type = GongCopyBranchProcess(mapOrigCopy, taskFrom.Type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskShape(mapOrigCopy map[any]any, taskshapeFrom *TaskShape) (taskshapeTo *TaskShape) {

	// taskshapeFrom has already been copied
	if _taskshapeTo, ok := mapOrigCopy[taskshapeFrom]; ok {
		taskshapeTo = _taskshapeTo.(*TaskShape)
		return
	}

	taskshapeTo = new(TaskShape)
	mapOrigCopy[taskshapeFrom] = taskshapeTo
	taskshapeFrom.GongCopyBasicFields(taskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskshapeFrom.Task != nil {
		taskshapeTo.Task = GongCopyBranchTask(mapOrigCopy, taskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (allocatedprocessshape *AllocatedProcessShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAllocatedProcessShape(allocatedprocessshape)
}

func (stage *Stage) UnstageBranchAllocatedProcessShape(allocatedprocessshape *AllocatedProcessShape) {

	// check if instance is already staged
	if !stage.IsStaged(allocatedprocessshape) {
		return
	}

	allocatedprocessshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedprocessshape.Participant != nil {
		stage.UnstageBranch(allocatedprocessshape.Participant)
	}
	if allocatedprocessshape.Process != nil {
		stage.UnstageBranch(allocatedprocessshape.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (allocatedresourceshape *AllocatedResourceShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAllocatedResourceShape(allocatedresourceshape)
}

func (stage *Stage) UnstageBranchAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) {

	// check if instance is already staged
	if !stage.IsStaged(allocatedresourceshape) {
		return
	}

	allocatedresourceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshape.Participant != nil {
		stage.UnstageBranch(allocatedresourceshape.Participant)
	}
	if allocatedresourceshape.Resource != nil {
		stage.UnstageBranch(allocatedresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlflow *ControlFlow) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchControlFlow(controlflow)
}

func (stage *Stage) UnstageBranchControlFlow(controlflow *ControlFlow) {

	// check if instance is already staged
	if !stage.IsStaged(controlflow) {
		return
	}

	controlflow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflow.Start != nil {
		stage.UnstageBranch(controlflow.Start)
	}
	if controlflow.End != nil {
		stage.UnstageBranch(controlflow.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlflowshape *ControlFlowShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchControlFlowShape(controlflowshape)
}

func (stage *Stage) UnstageBranchControlFlowShape(controlflowshape *ControlFlowShape) {

	// check if instance is already staged
	if !stage.IsStaged(controlflowshape) {
		return
	}

	controlflowshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshape.ControlFlow != nil {
		stage.UnstageBranch(controlflowshape.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (data *Data) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchData(data)
}

func (stage *Stage) UnstageBranchData(data *Data) {

	// check if instance is already staged
	if !stage.IsStaged(data) {
		return
	}

	data.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dataflow *DataFlow) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDataFlow(dataflow)
}

func (stage *Stage) UnstageBranchDataFlow(dataflow *DataFlow) {

	// check if instance is already staged
	if !stage.IsStaged(dataflow) {
		return
	}

	dataflow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflow.StartTask != nil {
		stage.UnstageBranch(dataflow.StartTask)
	}
	if dataflow.EndTask != nil {
		stage.UnstageBranch(dataflow.EndTask)
	}
	if dataflow.StartExternalParticipant != nil {
		stage.UnstageBranch(dataflow.StartExternalParticipant)
	}
	if dataflow.EndExternalParticipant != nil {
		stage.UnstageBranch(dataflow.EndExternalParticipant)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflow.Datas {
		stage.UnstageBranch(_data)
	}

}

func (dataflowshape *DataFlowShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDataFlowShape(dataflowshape)
}

func (stage *Stage) UnstageBranchDataFlowShape(dataflowshape *DataFlowShape) {

	// check if instance is already staged
	if !stage.IsStaged(dataflowshape) {
		return
	}

	dataflowshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshape.DataFlow != nil {
		stage.UnstageBranch(dataflowshape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datashape *DataShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDataShape(datashape)
}

func (stage *Stage) UnstageBranchDataShape(datashape *DataShape) {

	// check if instance is already staged
	if !stage.IsStaged(datashape) {
		return
	}

	datashape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datashape.Data != nil {
		stage.UnstageBranch(datashape.Data)
	}
	if datashape.DataFlow != nil {
		stage.UnstageBranch(datashape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagramprocess *DiagramProcess) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDiagramProcess(diagramprocess)
}

func (stage *Stage) UnstageBranchDiagramProcess(diagramprocess *DiagramProcess) {

	// check if instance is already staged
	if !stage.IsStaged(diagramprocess) {
		return
	}

	diagramprocess.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _processshape := range diagramprocess.Process_Shapes {
		stage.UnstageBranch(_processshape)
	}
	for _, _process := range diagramprocess.ProcesssWhoseNodeIsExpanded {
		stage.UnstageBranch(_process)
	}
	for _, _participantshape := range diagramprocess.Participant_Shapes {
		stage.UnstageBranch(_participantshape)
	}
	for _, _participant := range diagramprocess.ParticipantWhoseNodeIsExpanded {
		stage.UnstageBranch(_participant)
	}
	for _, _externalparticipantshape := range diagramprocess.ExternalParticipant_Shapes {
		stage.UnstageBranch(_externalparticipantshape)
	}
	for _, _participant := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
		stage.UnstageBranch(_participant)
	}
	for _, _participant := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
		stage.UnstageBranch(_participant)
	}
	for _, _participant := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
		stage.UnstageBranch(_participant)
	}
	for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
		stage.UnstageBranch(_task)
	}
	for _, _taskshape := range diagramprocess.Task_Shapes {
		stage.UnstageBranch(_taskshape)
	}
	for _, _controlflow := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
		stage.UnstageBranch(_controlflow)
	}
	for _, _controlflowshape := range diagramprocess.ControlFlow_Shapes {
		stage.UnstageBranch(_controlflowshape)
	}
	for _, _dataflow := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
		stage.UnstageBranch(_dataflow)
	}
	for _, _dataflowshape := range diagramprocess.DataFlow_Shapes {
		stage.UnstageBranch(_dataflowshape)
	}
	for _, _data := range diagramprocess.DatasWhoseNodeIsExpanded {
		stage.UnstageBranch(_data)
	}
	for _, _datashape := range diagramprocess.Data_Shapes {
		stage.UnstageBranch(_datashape)
	}
	for _, _dataflow := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
		stage.UnstageBranch(_dataflow)
	}
	for _, _resource := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
		stage.UnstageBranch(_resource)
	}
	for _, _allocatedresourceshape := range diagramprocess.AllocatedResourceShapes {
		stage.UnstageBranch(_allocatedresourceshape)
	}
	for _, _process := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
		stage.UnstageBranch(_process)
	}
	for _, _allocatedprocessshape := range diagramprocess.AllocatedProcessShapes {
		stage.UnstageBranch(_allocatedprocessshape)
	}
	for _, _noteshape := range diagramprocess.Note_Shapes {
		stage.UnstageBranch(_noteshape)
	}
	for _, _note := range diagramprocess.NotesWhoseNodeIsExpanded {
		stage.UnstageBranch(_note)
	}
	for _, _notetaskshape := range diagramprocess.NoteTaskShapes {
		stage.UnstageBranch(_notetaskshape)
	}

}

func (externalparticipantshape *ExternalParticipantShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchExternalParticipantShape(externalparticipantshape)
}

func (stage *Stage) UnstageBranchExternalParticipantShape(externalparticipantshape *ExternalParticipantShape) {

	// check if instance is already staged
	if !stage.IsStaged(externalparticipantshape) {
		return
	}

	externalparticipantshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if externalparticipantshape.Participant != nil {
		stage.UnstageBranch(externalparticipantshape.Participant)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLibrary(library)
}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !stage.IsStaged(library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.UnstageBranch(_library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.UnstageBranch(_library)
	}
	for _, _process := range library.RootProcesses {
		stage.UnstageBranch(_process)
	}
	for _, _process := range library.ProcesssWhoseNodeIsExpanded {
		stage.UnstageBranch(_process)
	}
	for _, _dataflow := range library.RootDataFlows {
		stage.UnstageBranch(_dataflow)
	}
	for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
		stage.UnstageBranch(_dataflow)
	}
	for _, _data := range library.RootDatas {
		stage.UnstageBranch(_data)
	}
	for _, _data := range library.DatasWhoseNodeIsExpanded {
		stage.UnstageBranch(_data)
	}
	for _, _resource := range library.RootResources {
		stage.UnstageBranch(_resource)
	}
	for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
		stage.UnstageBranch(_resource)
	}
	for _, _participant := range library.ParticipantsWhoseNodeIsExpanded {
		stage.UnstageBranch(_participant)
	}
	for _, _note := range library.RootNotes {
		stage.UnstageBranch(_note)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		stage.UnstageBranch(_note)
	}

}

func (note *Note) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNote(note)
}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range note.Tasks {
		stage.UnstageBranch(_task)
	}

}

func (noteshape *NoteShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteShape(noteshape)
}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !stage.IsStaged(noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.UnstageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notetaskshape *NoteTaskShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteTaskShape(notetaskshape)
}

func (stage *Stage) UnstageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if !stage.IsStaged(notetaskshape) {
		return
	}

	notetaskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		stage.UnstageBranch(notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		stage.UnstageBranch(notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (participant *Participant) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchParticipant(participant)
}

func (stage *Stage) UnstageBranchParticipant(participant *Participant) {

	// check if instance is already staged
	if !stage.IsStaged(participant) {
		return
	}

	participant.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _resource := range participant.Resources {
		stage.UnstageBranch(_resource)
	}
	for _, _process := range participant.Processes {
		stage.UnstageBranch(_process)
	}
	for _, _task := range participant.Tasks {
		stage.UnstageBranch(_task)
	}
	for _, _controlflow := range participant.ControlFlows {
		stage.UnstageBranch(_controlflow)
	}
	for _, _task := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
		stage.UnstageBranch(_task)
	}
	for _, _task := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
		stage.UnstageBranch(_task)
	}
	for _, _task := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
		stage.UnstageBranch(_task)
	}
	for _, _task := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
		stage.UnstageBranch(_task)
	}

}

func (participantshape *ParticipantShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchParticipantShape(participantshape)
}

func (stage *Stage) UnstageBranchParticipantShape(participantshape *ParticipantShape) {

	// check if instance is already staged
	if !stage.IsStaged(participantshape) {
		return
	}

	participantshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if participantshape.Participant != nil {
		stage.UnstageBranch(participantshape.Participant)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (process *Process) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchProcess(process)
}

func (stage *Stage) UnstageBranchProcess(process *Process) {

	// check if instance is already staged
	if !stage.IsStaged(process) {
		return
	}

	process.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramprocess := range process.DiagramProcesss {
		stage.UnstageBranch(_diagramprocess)
	}
	for _, _diagramprocess := range process.DiagramProcessWhoseNodeIsExpanded {
		stage.UnstageBranch(_diagramprocess)
	}
	for _, _process := range process.SubProcesses {
		stage.UnstageBranch(_process)
	}
	for _, _participant := range process.Participants {
		stage.UnstageBranch(_participant)
	}
	for _, _participant := range process.ParticipantWhoseNodeIsExpanded {
		stage.UnstageBranch(_participant)
	}
	for _, _dataflow := range process.DataFlows {
		stage.UnstageBranch(_dataflow)
	}
	for _, _participant := range process.ExternalParticipants {
		stage.UnstageBranch(_participant)
	}
	for _, _participant := range process.ExternalParticipantWhoseNodeIsExpanded {
		stage.UnstageBranch(_participant)
	}

}

func (processshape *ProcessShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchProcessShape(processshape)
}

func (stage *Stage) UnstageBranchProcessShape(processshape *ProcessShape) {

	// check if instance is already staged
	if !stage.IsStaged(processshape) {
		return
	}

	processshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if processshape.Process != nil {
		stage.UnstageBranch(processshape.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resource *Resource) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchResource(resource)
}

func (stage *Stage) UnstageBranchResource(resource *Resource) {

	// check if instance is already staged
	if !stage.IsStaged(resource) {
		return
	}

	resource.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (task *Task) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTask(task)
}

func (stage *Stage) UnstageBranchTask(task *Task) {

	// check if instance is already staged
	if !stage.IsStaged(task) {
		return
	}

	task.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if task.Type != nil {
		stage.UnstageBranch(task.Type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskshape *TaskShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTaskShape(taskshape)
}

func (stage *Stage) UnstageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if !stage.IsStaged(taskshape) {
		return
	}

	taskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		stage.UnstageBranch(taskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AllocatedProcessShape) GongReconstructPointersFromReferences(stage *Stage, instance *AllocatedProcessShape) {
	// insertion point for pointers field
	if instance.Participant != nil {
		reference.Participant = stage.Participants_reference[instance.Participant]
	}
	if instance.Process != nil {
		reference.Process = stage.Processs_reference[instance.Process]
	}
	// insertion point for slice of pointers field
}

func (reference *AllocatedResourceShape) GongReconstructPointersFromReferences(stage *Stage, instance *AllocatedResourceShape) {
	// insertion point for pointers field
	if instance.Participant != nil {
		reference.Participant = stage.Participants_reference[instance.Participant]
	}
	if instance.Resource != nil {
		reference.Resource = stage.Resources_reference[instance.Resource]
	}
	// insertion point for slice of pointers field
}

func (reference *ControlFlow) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlow) {
	// insertion point for pointers field
	if instance.Start != nil {
		reference.Start = stage.Tasks_reference[instance.Start]
	}
	if instance.End != nil {
		reference.End = stage.Tasks_reference[instance.End]
	}
	// insertion point for slice of pointers field
}

func (reference *ControlFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlowShape) {
	// insertion point for pointers field
	if instance.ControlFlow != nil {
		reference.ControlFlow = stage.ControlFlows_reference[instance.ControlFlow]
	}
	// insertion point for slice of pointers field
}

func (reference *Data) GongReconstructPointersFromReferences(stage *Stage, instance *Data) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DataFlow) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlow) {
	// insertion point for pointers field
	if instance.StartTask != nil {
		reference.StartTask = stage.Tasks_reference[instance.StartTask]
	}
	if instance.EndTask != nil {
		reference.EndTask = stage.Tasks_reference[instance.EndTask]
	}
	if instance.StartExternalParticipant != nil {
		reference.StartExternalParticipant = stage.Participants_reference[instance.StartExternalParticipant]
	}
	if instance.EndExternalParticipant != nil {
		reference.EndExternalParticipant = stage.Participants_reference[instance.EndExternalParticipant]
	}
	// insertion point for slice of pointers field
	reference.Datas = reference.Datas[:0]
	for _, _b := range instance.Datas {
		reference.Datas = append(reference.Datas, stage.Datas_reference[_b])
	}
}

func (reference *DataFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlowShape) {
	// insertion point for pointers field
	if instance.DataFlow != nil {
		reference.DataFlow = stage.DataFlows_reference[instance.DataFlow]
	}
	// insertion point for slice of pointers field
}

func (reference *DataShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataShape) {
	// insertion point for pointers field
	if instance.Data != nil {
		reference.Data = stage.Datas_reference[instance.Data]
	}
	if instance.DataFlow != nil {
		reference.DataFlow = stage.DataFlows_reference[instance.DataFlow]
	}
	// insertion point for slice of pointers field
}

func (reference *DiagramProcess) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramProcess) {
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
	reference.ExternalParticipant_Shapes = reference.ExternalParticipant_Shapes[:0]
	for _, _b := range instance.ExternalParticipant_Shapes {
		reference.ExternalParticipant_Shapes = append(reference.ExternalParticipant_Shapes, stage.ExternalParticipantShapes_reference[_b])
	}
	reference.ExternalParticipantWhoseNodeIsExpanded = reference.ExternalParticipantWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ExternalParticipantWhoseNodeIsExpanded {
		reference.ExternalParticipantWhoseNodeIsExpanded = append(reference.ExternalParticipantWhoseNodeIsExpanded, stage.Participants_reference[_b])
	}
	reference.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = reference.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[:0]
	for _, _b := range instance.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
		reference.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(reference.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, stage.Participants_reference[_b])
	}
	reference.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = reference.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[:0]
	for _, _b := range instance.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
		reference.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(reference.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, stage.Participants_reference[_b])
	}
	reference.TasksWhoseNodeIsExpanded = reference.TasksWhoseNodeIsExpanded[:0]
	for _, _b := range instance.TasksWhoseNodeIsExpanded {
		reference.TasksWhoseNodeIsExpanded = append(reference.TasksWhoseNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.Task_Shapes = reference.Task_Shapes[:0]
	for _, _b := range instance.Task_Shapes {
		reference.Task_Shapes = append(reference.Task_Shapes, stage.TaskShapes_reference[_b])
	}
	reference.ControlFlowsWhoseNodeIsExpanded = reference.ControlFlowsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ControlFlowsWhoseNodeIsExpanded {
		reference.ControlFlowsWhoseNodeIsExpanded = append(reference.ControlFlowsWhoseNodeIsExpanded, stage.ControlFlows_reference[_b])
	}
	reference.ControlFlow_Shapes = reference.ControlFlow_Shapes[:0]
	for _, _b := range instance.ControlFlow_Shapes {
		reference.ControlFlow_Shapes = append(reference.ControlFlow_Shapes, stage.ControlFlowShapes_reference[_b])
	}
	reference.DataFlowsWhoseNodeIsExpanded = reference.DataFlowsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DataFlowsWhoseNodeIsExpanded {
		reference.DataFlowsWhoseNodeIsExpanded = append(reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_reference[_b])
	}
	reference.DataFlow_Shapes = reference.DataFlow_Shapes[:0]
	for _, _b := range instance.DataFlow_Shapes {
		reference.DataFlow_Shapes = append(reference.DataFlow_Shapes, stage.DataFlowShapes_reference[_b])
	}
	reference.DatasWhoseNodeIsExpanded = reference.DatasWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DatasWhoseNodeIsExpanded {
		reference.DatasWhoseNodeIsExpanded = append(reference.DatasWhoseNodeIsExpanded, stage.Datas_reference[_b])
	}
	reference.Data_Shapes = reference.Data_Shapes[:0]
	for _, _b := range instance.Data_Shapes {
		reference.Data_Shapes = append(reference.Data_Shapes, stage.DataShapes_reference[_b])
	}
	reference.DataFlowsWhoseDataNodeIsExpanded = reference.DataFlowsWhoseDataNodeIsExpanded[:0]
	for _, _b := range instance.DataFlowsWhoseDataNodeIsExpanded {
		reference.DataFlowsWhoseDataNodeIsExpanded = append(reference.DataFlowsWhoseDataNodeIsExpanded, stage.DataFlows_reference[_b])
	}
	reference.AllocatedResourcesWhoseNodeIsExpanded = reference.AllocatedResourcesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.AllocatedResourcesWhoseNodeIsExpanded {
		reference.AllocatedResourcesWhoseNodeIsExpanded = append(reference.AllocatedResourcesWhoseNodeIsExpanded, stage.Resources_reference[_b])
	}
	reference.AllocatedResourceShapes = reference.AllocatedResourceShapes[:0]
	for _, _b := range instance.AllocatedResourceShapes {
		reference.AllocatedResourceShapes = append(reference.AllocatedResourceShapes, stage.AllocatedResourceShapes_reference[_b])
	}
	reference.AllocatedProcessesWhoseNodeIsExpanded = reference.AllocatedProcessesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.AllocatedProcessesWhoseNodeIsExpanded {
		reference.AllocatedProcessesWhoseNodeIsExpanded = append(reference.AllocatedProcessesWhoseNodeIsExpanded, stage.Processs_reference[_b])
	}
	reference.AllocatedProcessShapes = reference.AllocatedProcessShapes[:0]
	for _, _b := range instance.AllocatedProcessShapes {
		reference.AllocatedProcessShapes = append(reference.AllocatedProcessShapes, stage.AllocatedProcessShapes_reference[_b])
	}
	reference.Note_Shapes = reference.Note_Shapes[:0]
	for _, _b := range instance.Note_Shapes {
		reference.Note_Shapes = append(reference.Note_Shapes, stage.NoteShapes_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
	reference.NoteTaskShapes = reference.NoteTaskShapes[:0]
	for _, _b := range instance.NoteTaskShapes {
		reference.NoteTaskShapes = append(reference.NoteTaskShapes, stage.NoteTaskShapes_reference[_b])
	}
}

func (reference *ExternalParticipantShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExternalParticipantShape) {
	// insertion point for pointers field
	if instance.Participant != nil {
		reference.Participant = stage.Participants_reference[instance.Participant]
	}
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
	reference.RootProcesses = reference.RootProcesses[:0]
	for _, _b := range instance.RootProcesses {
		reference.RootProcesses = append(reference.RootProcesses, stage.Processs_reference[_b])
	}
	reference.ProcesssWhoseNodeIsExpanded = reference.ProcesssWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ProcesssWhoseNodeIsExpanded {
		reference.ProcesssWhoseNodeIsExpanded = append(reference.ProcesssWhoseNodeIsExpanded, stage.Processs_reference[_b])
	}
	reference.RootDataFlows = reference.RootDataFlows[:0]
	for _, _b := range instance.RootDataFlows {
		reference.RootDataFlows = append(reference.RootDataFlows, stage.DataFlows_reference[_b])
	}
	reference.DataFlowsWhoseNodeIsExpanded = reference.DataFlowsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DataFlowsWhoseNodeIsExpanded {
		reference.DataFlowsWhoseNodeIsExpanded = append(reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_reference[_b])
	}
	reference.RootDatas = reference.RootDatas[:0]
	for _, _b := range instance.RootDatas {
		reference.RootDatas = append(reference.RootDatas, stage.Datas_reference[_b])
	}
	reference.DatasWhoseNodeIsExpanded = reference.DatasWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DatasWhoseNodeIsExpanded {
		reference.DatasWhoseNodeIsExpanded = append(reference.DatasWhoseNodeIsExpanded, stage.Datas_reference[_b])
	}
	reference.RootResources = reference.RootResources[:0]
	for _, _b := range instance.RootResources {
		reference.RootResources = append(reference.RootResources, stage.Resources_reference[_b])
	}
	reference.ResourcesWhoseNodeIsExpanded = reference.ResourcesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ResourcesWhoseNodeIsExpanded {
		reference.ResourcesWhoseNodeIsExpanded = append(reference.ResourcesWhoseNodeIsExpanded, stage.Resources_reference[_b])
	}
	reference.ParticipantsWhoseNodeIsExpanded = reference.ParticipantsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ParticipantsWhoseNodeIsExpanded {
		reference.ParticipantsWhoseNodeIsExpanded = append(reference.ParticipantsWhoseNodeIsExpanded, stage.Participants_reference[_b])
	}
	reference.RootNotes = reference.RootNotes[:0]
	for _, _b := range instance.RootNotes {
		reference.RootNotes = append(reference.RootNotes, stage.Notes_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Tasks = reference.Tasks[:0]
	for _, _b := range instance.Tasks {
		reference.Tasks = append(reference.Tasks, stage.Tasks_reference[_b])
	}
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteTaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteTaskShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

func (reference *Participant) GongReconstructPointersFromReferences(stage *Stage, instance *Participant) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Resources = reference.Resources[:0]
	for _, _b := range instance.Resources {
		reference.Resources = append(reference.Resources, stage.Resources_reference[_b])
	}
	reference.Processes = reference.Processes[:0]
	for _, _b := range instance.Processes {
		reference.Processes = append(reference.Processes, stage.Processs_reference[_b])
	}
	reference.Tasks = reference.Tasks[:0]
	for _, _b := range instance.Tasks {
		reference.Tasks = append(reference.Tasks, stage.Tasks_reference[_b])
	}
	reference.ControlFlows = reference.ControlFlows[:0]
	for _, _b := range instance.ControlFlows {
		reference.ControlFlows = append(reference.ControlFlows, stage.ControlFlows_reference[_b])
	}
	reference.TaskWhoseOutControlFlowsNodeIsExpanded = reference.TaskWhoseOutControlFlowsNodeIsExpanded[:0]
	for _, _b := range instance.TaskWhoseOutControlFlowsNodeIsExpanded {
		reference.TaskWhoseOutControlFlowsNodeIsExpanded = append(reference.TaskWhoseOutControlFlowsNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.TaskWhoseInControlFlowsNodeIsExpanded = reference.TaskWhoseInControlFlowsNodeIsExpanded[:0]
	for _, _b := range instance.TaskWhoseInControlFlowsNodeIsExpanded {
		reference.TaskWhoseInControlFlowsNodeIsExpanded = append(reference.TaskWhoseInControlFlowsNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.TaskWhoseOutDataFlowsNodeIsExpanded = reference.TaskWhoseOutDataFlowsNodeIsExpanded[:0]
	for _, _b := range instance.TaskWhoseOutDataFlowsNodeIsExpanded {
		reference.TaskWhoseOutDataFlowsNodeIsExpanded = append(reference.TaskWhoseOutDataFlowsNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.TaskWhoseInDataFlowsNodeIsExpanded = reference.TaskWhoseInDataFlowsNodeIsExpanded[:0]
	for _, _b := range instance.TaskWhoseInDataFlowsNodeIsExpanded {
		reference.TaskWhoseInDataFlowsNodeIsExpanded = append(reference.TaskWhoseInDataFlowsNodeIsExpanded, stage.Tasks_reference[_b])
	}
}

func (reference *ParticipantShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParticipantShape) {
	// insertion point for pointers field
	if instance.Participant != nil {
		reference.Participant = stage.Participants_reference[instance.Participant]
	}
	// insertion point for slice of pointers field
}

func (reference *Process) GongReconstructPointersFromReferences(stage *Stage, instance *Process) {
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
	reference.DataFlows = reference.DataFlows[:0]
	for _, _b := range instance.DataFlows {
		reference.DataFlows = append(reference.DataFlows, stage.DataFlows_reference[_b])
	}
	reference.ExternalParticipants = reference.ExternalParticipants[:0]
	for _, _b := range instance.ExternalParticipants {
		reference.ExternalParticipants = append(reference.ExternalParticipants, stage.Participants_reference[_b])
	}
	reference.ExternalParticipantWhoseNodeIsExpanded = reference.ExternalParticipantWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ExternalParticipantWhoseNodeIsExpanded {
		reference.ExternalParticipantWhoseNodeIsExpanded = append(reference.ExternalParticipantWhoseNodeIsExpanded, stage.Participants_reference[_b])
	}
}

func (reference *ProcessShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProcessShape) {
	// insertion point for pointers field
	if instance.Process != nil {
		reference.Process = stage.Processs_reference[instance.Process]
	}
	// insertion point for slice of pointers field
}

func (reference *Resource) GongReconstructPointersFromReferences(stage *Stage, instance *Resource) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Task) GongReconstructPointersFromReferences(stage *Stage, instance *Task) {
	// insertion point for pointers field
	if instance.Type != nil {
		reference.Type = stage.Processs_reference[instance.Type]
	}
	// insertion point for slice of pointers field
}

func (reference *TaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskShape) {
	// insertion point for pointers field
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AllocatedProcessShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Participant; _reference != nil {
		reference.Participant = nil
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			reference.Participant = _instance
		}
	}
	if _reference := reference.Process; _reference != nil {
		reference.Process = nil
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			reference.Process = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *AllocatedResourceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Participant; _reference != nil {
		reference.Participant = nil
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			reference.Participant = _instance
		}
	}
	if _reference := reference.Resource; _reference != nil {
		reference.Resource = nil
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			reference.Resource = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ControlFlow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Start; _reference != nil {
		reference.Start = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Start = _instance
		}
	}
	if _reference := reference.End; _reference != nil {
		reference.End = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.End = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ControlFlowShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ControlFlow; _reference != nil {
		reference.ControlFlow = nil
		if _instance, ok := stage.ControlFlows_instance[_reference]; ok {
			reference.ControlFlow = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Data) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DataFlow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.StartTask; _reference != nil {
		reference.StartTask = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.StartTask = _instance
		}
	}
	if _reference := reference.EndTask; _reference != nil {
		reference.EndTask = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.EndTask = _instance
		}
	}
	if _reference := reference.StartExternalParticipant; _reference != nil {
		reference.StartExternalParticipant = nil
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			reference.StartExternalParticipant = _instance
		}
	}
	if _reference := reference.EndExternalParticipant; _reference != nil {
		reference.EndExternalParticipant = nil
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			reference.EndExternalParticipant = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Datas []*Data
	for _, _reference := range reference.Datas {
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			_Datas = append(_Datas, _instance)
		}
	}
	reference.Datas = _Datas
}

func (reference *DataFlowShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DataFlow; _reference != nil {
		reference.DataFlow = nil
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			reference.DataFlow = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DataShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Data; _reference != nil {
		reference.Data = nil
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			reference.Data = _instance
		}
	}
	if _reference := reference.DataFlow; _reference != nil {
		reference.DataFlow = nil
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			reference.DataFlow = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DiagramProcess) GongReconstructPointersFromInstances(stage *Stage) {
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
	var _ExternalParticipant_Shapes []*ExternalParticipantShape
	for _, _reference := range reference.ExternalParticipant_Shapes {
		if _instance, ok := stage.ExternalParticipantShapes_instance[_reference]; ok {
			_ExternalParticipant_Shapes = append(_ExternalParticipant_Shapes, _instance)
		}
	}
	reference.ExternalParticipant_Shapes = _ExternalParticipant_Shapes
	var _ExternalParticipantWhoseNodeIsExpanded []*Participant
	for _, _reference := range reference.ExternalParticipantWhoseNodeIsExpanded {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_ExternalParticipantWhoseNodeIsExpanded = append(_ExternalParticipantWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ExternalParticipantWhoseNodeIsExpanded = _ExternalParticipantWhoseNodeIsExpanded
	var _ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded []*Participant
	for _, _reference := range reference.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(_ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, _instance)
		}
	}
	reference.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = _ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded
	var _ExternalParticipantsWhoseInDataFlowsNodeIsExpanded []*Participant
	for _, _reference := range reference.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(_ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, _instance)
		}
	}
	reference.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = _ExternalParticipantsWhoseInDataFlowsNodeIsExpanded
	var _TasksWhoseNodeIsExpanded []*Task
	for _, _reference := range reference.TasksWhoseNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TasksWhoseNodeIsExpanded = append(_TasksWhoseNodeIsExpanded, _instance)
		}
	}
	reference.TasksWhoseNodeIsExpanded = _TasksWhoseNodeIsExpanded
	var _Task_Shapes []*TaskShape
	for _, _reference := range reference.Task_Shapes {
		if _instance, ok := stage.TaskShapes_instance[_reference]; ok {
			_Task_Shapes = append(_Task_Shapes, _instance)
		}
	}
	reference.Task_Shapes = _Task_Shapes
	var _ControlFlowsWhoseNodeIsExpanded []*ControlFlow
	for _, _reference := range reference.ControlFlowsWhoseNodeIsExpanded {
		if _instance, ok := stage.ControlFlows_instance[_reference]; ok {
			_ControlFlowsWhoseNodeIsExpanded = append(_ControlFlowsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ControlFlowsWhoseNodeIsExpanded = _ControlFlowsWhoseNodeIsExpanded
	var _ControlFlow_Shapes []*ControlFlowShape
	for _, _reference := range reference.ControlFlow_Shapes {
		if _instance, ok := stage.ControlFlowShapes_instance[_reference]; ok {
			_ControlFlow_Shapes = append(_ControlFlow_Shapes, _instance)
		}
	}
	reference.ControlFlow_Shapes = _ControlFlow_Shapes
	var _DataFlowsWhoseNodeIsExpanded []*DataFlow
	for _, _reference := range reference.DataFlowsWhoseNodeIsExpanded {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_DataFlowsWhoseNodeIsExpanded = append(_DataFlowsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DataFlowsWhoseNodeIsExpanded = _DataFlowsWhoseNodeIsExpanded
	var _DataFlow_Shapes []*DataFlowShape
	for _, _reference := range reference.DataFlow_Shapes {
		if _instance, ok := stage.DataFlowShapes_instance[_reference]; ok {
			_DataFlow_Shapes = append(_DataFlow_Shapes, _instance)
		}
	}
	reference.DataFlow_Shapes = _DataFlow_Shapes
	var _DatasWhoseNodeIsExpanded []*Data
	for _, _reference := range reference.DatasWhoseNodeIsExpanded {
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			_DatasWhoseNodeIsExpanded = append(_DatasWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DatasWhoseNodeIsExpanded = _DatasWhoseNodeIsExpanded
	var _Data_Shapes []*DataShape
	for _, _reference := range reference.Data_Shapes {
		if _instance, ok := stage.DataShapes_instance[_reference]; ok {
			_Data_Shapes = append(_Data_Shapes, _instance)
		}
	}
	reference.Data_Shapes = _Data_Shapes
	var _DataFlowsWhoseDataNodeIsExpanded []*DataFlow
	for _, _reference := range reference.DataFlowsWhoseDataNodeIsExpanded {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_DataFlowsWhoseDataNodeIsExpanded = append(_DataFlowsWhoseDataNodeIsExpanded, _instance)
		}
	}
	reference.DataFlowsWhoseDataNodeIsExpanded = _DataFlowsWhoseDataNodeIsExpanded
	var _AllocatedResourcesWhoseNodeIsExpanded []*Resource
	for _, _reference := range reference.AllocatedResourcesWhoseNodeIsExpanded {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_AllocatedResourcesWhoseNodeIsExpanded = append(_AllocatedResourcesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.AllocatedResourcesWhoseNodeIsExpanded = _AllocatedResourcesWhoseNodeIsExpanded
	var _AllocatedResourceShapes []*AllocatedResourceShape
	for _, _reference := range reference.AllocatedResourceShapes {
		if _instance, ok := stage.AllocatedResourceShapes_instance[_reference]; ok {
			_AllocatedResourceShapes = append(_AllocatedResourceShapes, _instance)
		}
	}
	reference.AllocatedResourceShapes = _AllocatedResourceShapes
	var _AllocatedProcessesWhoseNodeIsExpanded []*Process
	for _, _reference := range reference.AllocatedProcessesWhoseNodeIsExpanded {
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			_AllocatedProcessesWhoseNodeIsExpanded = append(_AllocatedProcessesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.AllocatedProcessesWhoseNodeIsExpanded = _AllocatedProcessesWhoseNodeIsExpanded
	var _AllocatedProcessShapes []*AllocatedProcessShape
	for _, _reference := range reference.AllocatedProcessShapes {
		if _instance, ok := stage.AllocatedProcessShapes_instance[_reference]; ok {
			_AllocatedProcessShapes = append(_AllocatedProcessShapes, _instance)
		}
	}
	reference.AllocatedProcessShapes = _AllocatedProcessShapes
	var _Note_Shapes []*NoteShape
	for _, _reference := range reference.Note_Shapes {
		if _instance, ok := stage.NoteShapes_instance[_reference]; ok {
			_Note_Shapes = append(_Note_Shapes, _instance)
		}
	}
	reference.Note_Shapes = _Note_Shapes
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
	var _NoteTaskShapes []*NoteTaskShape
	for _, _reference := range reference.NoteTaskShapes {
		if _instance, ok := stage.NoteTaskShapes_instance[_reference]; ok {
			_NoteTaskShapes = append(_NoteTaskShapes, _instance)
		}
	}
	reference.NoteTaskShapes = _NoteTaskShapes
}

func (reference *ExternalParticipantShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Participant; _reference != nil {
		reference.Participant = nil
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			reference.Participant = _instance
		}
	}
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
	var _RootDataFlows []*DataFlow
	for _, _reference := range reference.RootDataFlows {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_RootDataFlows = append(_RootDataFlows, _instance)
		}
	}
	reference.RootDataFlows = _RootDataFlows
	var _DataFlowsWhoseNodeIsExpanded []*DataFlow
	for _, _reference := range reference.DataFlowsWhoseNodeIsExpanded {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_DataFlowsWhoseNodeIsExpanded = append(_DataFlowsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DataFlowsWhoseNodeIsExpanded = _DataFlowsWhoseNodeIsExpanded
	var _RootDatas []*Data
	for _, _reference := range reference.RootDatas {
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			_RootDatas = append(_RootDatas, _instance)
		}
	}
	reference.RootDatas = _RootDatas
	var _DatasWhoseNodeIsExpanded []*Data
	for _, _reference := range reference.DatasWhoseNodeIsExpanded {
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			_DatasWhoseNodeIsExpanded = append(_DatasWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DatasWhoseNodeIsExpanded = _DatasWhoseNodeIsExpanded
	var _RootResources []*Resource
	for _, _reference := range reference.RootResources {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_RootResources = append(_RootResources, _instance)
		}
	}
	reference.RootResources = _RootResources
	var _ResourcesWhoseNodeIsExpanded []*Resource
	for _, _reference := range reference.ResourcesWhoseNodeIsExpanded {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_ResourcesWhoseNodeIsExpanded = append(_ResourcesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ResourcesWhoseNodeIsExpanded = _ResourcesWhoseNodeIsExpanded
	var _ParticipantsWhoseNodeIsExpanded []*Participant
	for _, _reference := range reference.ParticipantsWhoseNodeIsExpanded {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_ParticipantsWhoseNodeIsExpanded = append(_ParticipantsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ParticipantsWhoseNodeIsExpanded = _ParticipantsWhoseNodeIsExpanded
	var _RootNotes []*Note
	for _, _reference := range reference.RootNotes {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_RootNotes = append(_RootNotes, _instance)
		}
	}
	reference.RootNotes = _RootNotes
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Tasks []*Task
	for _, _reference := range reference.Tasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_Tasks = append(_Tasks, _instance)
		}
	}
	reference.Tasks = _Tasks
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteTaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Participant) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Resources []*Resource
	for _, _reference := range reference.Resources {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_Resources = append(_Resources, _instance)
		}
	}
	reference.Resources = _Resources
	var _Processes []*Process
	for _, _reference := range reference.Processes {
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			_Processes = append(_Processes, _instance)
		}
	}
	reference.Processes = _Processes
	var _Tasks []*Task
	for _, _reference := range reference.Tasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_Tasks = append(_Tasks, _instance)
		}
	}
	reference.Tasks = _Tasks
	var _ControlFlows []*ControlFlow
	for _, _reference := range reference.ControlFlows {
		if _instance, ok := stage.ControlFlows_instance[_reference]; ok {
			_ControlFlows = append(_ControlFlows, _instance)
		}
	}
	reference.ControlFlows = _ControlFlows
	var _TaskWhoseOutControlFlowsNodeIsExpanded []*Task
	for _, _reference := range reference.TaskWhoseOutControlFlowsNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TaskWhoseOutControlFlowsNodeIsExpanded = append(_TaskWhoseOutControlFlowsNodeIsExpanded, _instance)
		}
	}
	reference.TaskWhoseOutControlFlowsNodeIsExpanded = _TaskWhoseOutControlFlowsNodeIsExpanded
	var _TaskWhoseInControlFlowsNodeIsExpanded []*Task
	for _, _reference := range reference.TaskWhoseInControlFlowsNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TaskWhoseInControlFlowsNodeIsExpanded = append(_TaskWhoseInControlFlowsNodeIsExpanded, _instance)
		}
	}
	reference.TaskWhoseInControlFlowsNodeIsExpanded = _TaskWhoseInControlFlowsNodeIsExpanded
	var _TaskWhoseOutDataFlowsNodeIsExpanded []*Task
	for _, _reference := range reference.TaskWhoseOutDataFlowsNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TaskWhoseOutDataFlowsNodeIsExpanded = append(_TaskWhoseOutDataFlowsNodeIsExpanded, _instance)
		}
	}
	reference.TaskWhoseOutDataFlowsNodeIsExpanded = _TaskWhoseOutDataFlowsNodeIsExpanded
	var _TaskWhoseInDataFlowsNodeIsExpanded []*Task
	for _, _reference := range reference.TaskWhoseInDataFlowsNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TaskWhoseInDataFlowsNodeIsExpanded = append(_TaskWhoseInDataFlowsNodeIsExpanded, _instance)
		}
	}
	reference.TaskWhoseInDataFlowsNodeIsExpanded = _TaskWhoseInDataFlowsNodeIsExpanded
}

func (reference *ParticipantShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Participant; _reference != nil {
		reference.Participant = nil
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			reference.Participant = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Process) GongReconstructPointersFromInstances(stage *Stage) {
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
	var _DataFlows []*DataFlow
	for _, _reference := range reference.DataFlows {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_DataFlows = append(_DataFlows, _instance)
		}
	}
	reference.DataFlows = _DataFlows
	var _ExternalParticipants []*Participant
	for _, _reference := range reference.ExternalParticipants {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_ExternalParticipants = append(_ExternalParticipants, _instance)
		}
	}
	reference.ExternalParticipants = _ExternalParticipants
	var _ExternalParticipantWhoseNodeIsExpanded []*Participant
	for _, _reference := range reference.ExternalParticipantWhoseNodeIsExpanded {
		if _instance, ok := stage.Participants_instance[_reference]; ok {
			_ExternalParticipantWhoseNodeIsExpanded = append(_ExternalParticipantWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ExternalParticipantWhoseNodeIsExpanded = _ExternalParticipantWhoseNodeIsExpanded
}

func (reference *ProcessShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Process; _reference != nil {
		reference.Process = nil
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			reference.Process = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Resource) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Task) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Type; _reference != nil {
		reference.Type = nil
		if _instance, ok := stage.Processs_instance[_reference]; ok {
			reference.Type = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (allocatedprocessshape *AllocatedProcessShape) GongDiff(stage *Stage, allocatedprocessshapeOther *AllocatedProcessShape) (diffs []string) {
	// insertion point for field diffs
	if allocatedprocessshape.Name != allocatedprocessshapeOther.Name {
		diffs = append(diffs, allocatedprocessshape.GongMarshallField(stage, "Name"))
	}
	if (allocatedprocessshape.Participant == nil) != (allocatedprocessshapeOther.Participant == nil) {
		diffs = append(diffs, allocatedprocessshape.GongMarshallField(stage, "Participant"))
	} else if allocatedprocessshape.Participant != nil && allocatedprocessshapeOther.Participant != nil {
		if allocatedprocessshape.Participant != allocatedprocessshapeOther.Participant {
			diffs = append(diffs, allocatedprocessshape.GongMarshallField(stage, "Participant"))
		}
	}
	if (allocatedprocessshape.Process == nil) != (allocatedprocessshapeOther.Process == nil) {
		diffs = append(diffs, allocatedprocessshape.GongMarshallField(stage, "Process"))
	} else if allocatedprocessshape.Process != nil && allocatedprocessshapeOther.Process != nil {
		if allocatedprocessshape.Process != allocatedprocessshapeOther.Process {
			diffs = append(diffs, allocatedprocessshape.GongMarshallField(stage, "Process"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (allocatedresourceshape *AllocatedResourceShape) GongDiff(stage *Stage, allocatedresourceshapeOther *AllocatedResourceShape) (diffs []string) {
	// insertion point for field diffs
	if allocatedresourceshape.Name != allocatedresourceshapeOther.Name {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Name"))
	}
	if (allocatedresourceshape.Participant == nil) != (allocatedresourceshapeOther.Participant == nil) {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Participant"))
	} else if allocatedresourceshape.Participant != nil && allocatedresourceshapeOther.Participant != nil {
		if allocatedresourceshape.Participant != allocatedresourceshapeOther.Participant {
			diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Participant"))
		}
	}
	if (allocatedresourceshape.Resource == nil) != (allocatedresourceshapeOther.Resource == nil) {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Resource"))
	} else if allocatedresourceshape.Resource != nil && allocatedresourceshapeOther.Resource != nil {
		if allocatedresourceshape.Resource != allocatedresourceshapeOther.Resource {
			diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Resource"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlflow *ControlFlow) GongDiff(stage *Stage, controlflowOther *ControlFlow) (diffs []string) {
	// insertion point for field diffs
	if controlflow.Name != controlflowOther.Name {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Name"))
	}
	if controlflow.Description != controlflowOther.Description {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Description"))
	}
	if controlflow.ComputedPrefix != controlflowOther.ComputedPrefix {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "ComputedPrefix"))
	}
	if controlflow.IsExpanded != controlflowOther.IsExpanded {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "IsExpanded"))
	}
	if (controlflow.Start == nil) != (controlflowOther.Start == nil) {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Start"))
	} else if controlflow.Start != nil && controlflowOther.Start != nil {
		if controlflow.Start != controlflowOther.Start {
			diffs = append(diffs, controlflow.GongMarshallField(stage, "Start"))
		}
	}
	if (controlflow.End == nil) != (controlflowOther.End == nil) {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "End"))
	} else if controlflow.End != nil && controlflowOther.End != nil {
		if controlflow.End != controlflowOther.End {
			diffs = append(diffs, controlflow.GongMarshallField(stage, "End"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlflowshape *ControlFlowShape) GongDiff(stage *Stage, controlflowshapeOther *ControlFlowShape) (diffs []string) {
	// insertion point for field diffs
	if controlflowshape.Name != controlflowshapeOther.Name {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "Name"))
	}
	if (controlflowshape.ControlFlow == nil) != (controlflowshapeOther.ControlFlow == nil) {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "ControlFlow"))
	} else if controlflowshape.ControlFlow != nil && controlflowshapeOther.ControlFlow != nil {
		if controlflowshape.ControlFlow != controlflowshapeOther.ControlFlow {
			diffs = append(diffs, controlflowshape.GongMarshallField(stage, "ControlFlow"))
		}
	}
	if controlflowshape.StartRatio != controlflowshapeOther.StartRatio {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "StartRatio"))
	}
	if controlflowshape.EndRatio != controlflowshapeOther.EndRatio {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "EndRatio"))
	}
	if controlflowshape.StartOrientation != controlflowshapeOther.StartOrientation {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "StartOrientation"))
	}
	if controlflowshape.EndOrientation != controlflowshapeOther.EndOrientation {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "EndOrientation"))
	}
	if controlflowshape.CornerOffsetRatio != controlflowshapeOther.CornerOffsetRatio {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if controlflowshape.IsHidden != controlflowshapeOther.IsHidden {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (data *Data) GongDiff(stage *Stage, dataOther *Data) (diffs []string) {
	// insertion point for field diffs
	if data.Name != dataOther.Name {
		diffs = append(diffs, data.GongMarshallField(stage, "Name"))
	}
	if data.Acronym != dataOther.Acronym {
		diffs = append(diffs, data.GongMarshallField(stage, "Acronym"))
	}
	if data.Description != dataOther.Description {
		diffs = append(diffs, data.GongMarshallField(stage, "Description"))
	}
	if data.ComputedPrefix != dataOther.ComputedPrefix {
		diffs = append(diffs, data.GongMarshallField(stage, "ComputedPrefix"))
	}
	if data.IsExpanded != dataOther.IsExpanded {
		diffs = append(diffs, data.GongMarshallField(stage, "IsExpanded"))
	}
	if data.SVG_Path != dataOther.SVG_Path {
		diffs = append(diffs, data.GongMarshallField(stage, "SVG_Path"))
	}
	if data.InverseAppliedScaling != dataOther.InverseAppliedScaling {
		diffs = append(diffs, data.GongMarshallField(stage, "InverseAppliedScaling"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dataflow *DataFlow) GongDiff(stage *Stage, dataflowOther *DataFlow) (diffs []string) {
	// insertion point for field diffs
	if dataflow.Name != dataflowOther.Name {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Name"))
	}
	DatasDifferent := false
	if len(dataflow.Datas) != len(dataflowOther.Datas) {
		DatasDifferent = true
	} else {
		for i := range dataflow.Datas {
			if (dataflow.Datas[i] == nil) != (dataflowOther.Datas[i] == nil) {
				DatasDifferent = true
				break
			} else if dataflow.Datas[i] != nil && dataflowOther.Datas[i] != nil {
				// this is a pointer comparaison
				if dataflow.Datas[i] != dataflowOther.Datas[i] {
					DatasDifferent = true
					break
				}
			}
		}
	}
	if DatasDifferent {
		ops := stage.Diff(
			dataflow,
			"Datas",
			len(dataflowOther.Datas),
			len(dataflow.Datas),
			func(i, j int) bool {
				return dataflowOther.Datas[i] == dataflow.Datas[j]
			},
			func(j int) string {
				return dataflow.Datas[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if dataflow.Description != dataflowOther.Description {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Description"))
	}
	if dataflow.ComputedPrefix != dataflowOther.ComputedPrefix {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "ComputedPrefix"))
	}
	if dataflow.IsExpanded != dataflowOther.IsExpanded {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "IsExpanded"))
	}
	if dataflow.Type != dataflowOther.Type {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Type"))
	}
	if (dataflow.StartTask == nil) != (dataflowOther.StartTask == nil) {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "StartTask"))
	} else if dataflow.StartTask != nil && dataflowOther.StartTask != nil {
		if dataflow.StartTask != dataflowOther.StartTask {
			diffs = append(diffs, dataflow.GongMarshallField(stage, "StartTask"))
		}
	}
	if (dataflow.EndTask == nil) != (dataflowOther.EndTask == nil) {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "EndTask"))
	} else if dataflow.EndTask != nil && dataflowOther.EndTask != nil {
		if dataflow.EndTask != dataflowOther.EndTask {
			diffs = append(diffs, dataflow.GongMarshallField(stage, "EndTask"))
		}
	}
	if (dataflow.StartExternalParticipant == nil) != (dataflowOther.StartExternalParticipant == nil) {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "StartExternalParticipant"))
	} else if dataflow.StartExternalParticipant != nil && dataflowOther.StartExternalParticipant != nil {
		if dataflow.StartExternalParticipant != dataflowOther.StartExternalParticipant {
			diffs = append(diffs, dataflow.GongMarshallField(stage, "StartExternalParticipant"))
		}
	}
	if (dataflow.EndExternalParticipant == nil) != (dataflowOther.EndExternalParticipant == nil) {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "EndExternalParticipant"))
	} else if dataflow.EndExternalParticipant != nil && dataflowOther.EndExternalParticipant != nil {
		if dataflow.EndExternalParticipant != dataflowOther.EndExternalParticipant {
			diffs = append(diffs, dataflow.GongMarshallField(stage, "EndExternalParticipant"))
		}
	}
	if dataflow.IsDatasNodeExpanded != dataflowOther.IsDatasNodeExpanded {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "IsDatasNodeExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dataflowshape *DataFlowShape) GongDiff(stage *Stage, dataflowshapeOther *DataFlowShape) (diffs []string) {
	// insertion point for field diffs
	if dataflowshape.Name != dataflowshapeOther.Name {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "Name"))
	}
	if (dataflowshape.DataFlow == nil) != (dataflowshapeOther.DataFlow == nil) {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "DataFlow"))
	} else if dataflowshape.DataFlow != nil && dataflowshapeOther.DataFlow != nil {
		if dataflowshape.DataFlow != dataflowshapeOther.DataFlow {
			diffs = append(diffs, dataflowshape.GongMarshallField(stage, "DataFlow"))
		}
	}
	if dataflowshape.StartRatio != dataflowshapeOther.StartRatio {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "StartRatio"))
	}
	if dataflowshape.EndRatio != dataflowshapeOther.EndRatio {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "EndRatio"))
	}
	if dataflowshape.StartOrientation != dataflowshapeOther.StartOrientation {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "StartOrientation"))
	}
	if dataflowshape.EndOrientation != dataflowshapeOther.EndOrientation {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "EndOrientation"))
	}
	if dataflowshape.CornerOffsetRatio != dataflowshapeOther.CornerOffsetRatio {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if dataflowshape.IsHidden != dataflowshapeOther.IsHidden {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datashape *DataShape) GongDiff(stage *Stage, datashapeOther *DataShape) (diffs []string) {
	// insertion point for field diffs
	if datashape.Name != datashapeOther.Name {
		diffs = append(diffs, datashape.GongMarshallField(stage, "Name"))
	}
	if (datashape.Data == nil) != (datashapeOther.Data == nil) {
		diffs = append(diffs, datashape.GongMarshallField(stage, "Data"))
	} else if datashape.Data != nil && datashapeOther.Data != nil {
		if datashape.Data != datashapeOther.Data {
			diffs = append(diffs, datashape.GongMarshallField(stage, "Data"))
		}
	}
	if (datashape.DataFlow == nil) != (datashapeOther.DataFlow == nil) {
		diffs = append(diffs, datashape.GongMarshallField(stage, "DataFlow"))
	} else if datashape.DataFlow != nil && datashapeOther.DataFlow != nil {
		if datashape.DataFlow != datashapeOther.DataFlow {
			diffs = append(diffs, datashape.GongMarshallField(stage, "DataFlow"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagramprocess *DiagramProcess) GongDiff(stage *Stage, diagramprocessOther *DiagramProcess) (diffs []string) {
	// insertion point for field diffs
	if diagramprocess.Name != diagramprocessOther.Name {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "Name"))
	}
	if diagramprocess.Description != diagramprocessOther.Description {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "Description"))
	}
	if diagramprocess.ComputedPrefix != diagramprocessOther.ComputedPrefix {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagramprocess.IsExpanded != diagramprocessOther.IsExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsExpanded"))
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
		ops := stage.Diff(
			diagramprocess,
			"Process_Shapes",
			len(diagramprocessOther.Process_Shapes),
			len(diagramprocess.Process_Shapes),
			func(i, j int) bool {
				return diagramprocessOther.Process_Shapes[i] == diagramprocess.Process_Shapes[j]
			},
			func(j int) string {
				return diagramprocess.Process_Shapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagramprocess,
			"ProcesssWhoseNodeIsExpanded",
			len(diagramprocessOther.ProcesssWhoseNodeIsExpanded),
			len(diagramprocess.ProcesssWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.ProcesssWhoseNodeIsExpanded[i] == diagramprocess.ProcesssWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.ProcesssWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagramprocess,
			"Participant_Shapes",
			len(diagramprocessOther.Participant_Shapes),
			len(diagramprocess.Participant_Shapes),
			func(i, j int) bool {
				return diagramprocessOther.Participant_Shapes[i] == diagramprocess.Participant_Shapes[j]
			},
			func(j int) string {
				return diagramprocess.Participant_Shapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagramprocess,
			"ParticipantWhoseNodeIsExpanded",
			len(diagramprocessOther.ParticipantWhoseNodeIsExpanded),
			len(diagramprocess.ParticipantWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.ParticipantWhoseNodeIsExpanded[i] == diagramprocess.ParticipantWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.ParticipantWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ExternalParticipant_ShapesDifferent := false
	if len(diagramprocess.ExternalParticipant_Shapes) != len(diagramprocessOther.ExternalParticipant_Shapes) {
		ExternalParticipant_ShapesDifferent = true
	} else {
		for i := range diagramprocess.ExternalParticipant_Shapes {
			if (diagramprocess.ExternalParticipant_Shapes[i] == nil) != (diagramprocessOther.ExternalParticipant_Shapes[i] == nil) {
				ExternalParticipant_ShapesDifferent = true
				break
			} else if diagramprocess.ExternalParticipant_Shapes[i] != nil && diagramprocessOther.ExternalParticipant_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ExternalParticipant_Shapes[i] != diagramprocessOther.ExternalParticipant_Shapes[i] {
					ExternalParticipant_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ExternalParticipant_ShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"ExternalParticipant_Shapes",
			len(diagramprocessOther.ExternalParticipant_Shapes),
			len(diagramprocess.ExternalParticipant_Shapes),
			func(i, j int) bool {
				return diagramprocessOther.ExternalParticipant_Shapes[i] == diagramprocess.ExternalParticipant_Shapes[j]
			},
			func(j int) string {
				return diagramprocess.ExternalParticipant_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagramprocess.IsExternalParticipantsNodeExpanded != diagramprocessOther.IsExternalParticipantsNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsExternalParticipantsNodeExpanded"))
	}
	ExternalParticipantWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.ExternalParticipantWhoseNodeIsExpanded) != len(diagramprocessOther.ExternalParticipantWhoseNodeIsExpanded) {
		ExternalParticipantWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
			if (diagramprocess.ExternalParticipantWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.ExternalParticipantWhoseNodeIsExpanded[i] == nil) {
				ExternalParticipantWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.ExternalParticipantWhoseNodeIsExpanded[i] != nil && diagramprocessOther.ExternalParticipantWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ExternalParticipantWhoseNodeIsExpanded[i] != diagramprocessOther.ExternalParticipantWhoseNodeIsExpanded[i] {
					ExternalParticipantWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ExternalParticipantWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"ExternalParticipantWhoseNodeIsExpanded",
			len(diagramprocessOther.ExternalParticipantWhoseNodeIsExpanded),
			len(diagramprocess.ExternalParticipantWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.ExternalParticipantWhoseNodeIsExpanded[i] == diagramprocess.ExternalParticipantWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.ExternalParticipantWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ExternalParticipantsWhoseOutDataFlowsNodeIsExpandedDifferent := false
	if len(diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded) != len(diagramprocessOther.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded) {
		ExternalParticipantsWhoseOutDataFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
			if (diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[i] == nil) != (diagramprocessOther.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[i] == nil) {
				ExternalParticipantsWhoseOutDataFlowsNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[i] != nil && diagramprocessOther.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[i] != diagramprocessOther.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[i] {
					ExternalParticipantsWhoseOutDataFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ExternalParticipantsWhoseOutDataFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded",
			len(diagramprocessOther.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded),
			len(diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[i] == diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ExternalParticipantsWhoseInDataFlowsNodeIsExpandedDifferent := false
	if len(diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded) != len(diagramprocessOther.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded) {
		ExternalParticipantsWhoseInDataFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
			if (diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[i] == nil) != (diagramprocessOther.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[i] == nil) {
				ExternalParticipantsWhoseInDataFlowsNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[i] != nil && diagramprocessOther.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[i] != diagramprocessOther.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[i] {
					ExternalParticipantsWhoseInDataFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ExternalParticipantsWhoseInDataFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"ExternalParticipantsWhoseInDataFlowsNodeIsExpanded",
			len(diagramprocessOther.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded),
			len(diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[i] == diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagramprocess,
			"TasksWhoseNodeIsExpanded",
			len(diagramprocessOther.TasksWhoseNodeIsExpanded),
			len(diagramprocess.TasksWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.TasksWhoseNodeIsExpanded[i] == diagramprocess.TasksWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.TasksWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Task_ShapesDifferent := false
	if len(diagramprocess.Task_Shapes) != len(diagramprocessOther.Task_Shapes) {
		Task_ShapesDifferent = true
	} else {
		for i := range diagramprocess.Task_Shapes {
			if (diagramprocess.Task_Shapes[i] == nil) != (diagramprocessOther.Task_Shapes[i] == nil) {
				Task_ShapesDifferent = true
				break
			} else if diagramprocess.Task_Shapes[i] != nil && diagramprocessOther.Task_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.Task_Shapes[i] != diagramprocessOther.Task_Shapes[i] {
					Task_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Task_ShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"Task_Shapes",
			len(diagramprocessOther.Task_Shapes),
			len(diagramprocess.Task_Shapes),
			func(i, j int) bool {
				return diagramprocessOther.Task_Shapes[i] == diagramprocess.Task_Shapes[j]
			},
			func(j int) string {
				return diagramprocess.Task_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ControlFlowsWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.ControlFlowsWhoseNodeIsExpanded) != len(diagramprocessOther.ControlFlowsWhoseNodeIsExpanded) {
		ControlFlowsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
			if (diagramprocess.ControlFlowsWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.ControlFlowsWhoseNodeIsExpanded[i] == nil) {
				ControlFlowsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.ControlFlowsWhoseNodeIsExpanded[i] != nil && diagramprocessOther.ControlFlowsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ControlFlowsWhoseNodeIsExpanded[i] != diagramprocessOther.ControlFlowsWhoseNodeIsExpanded[i] {
					ControlFlowsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ControlFlowsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"ControlFlowsWhoseNodeIsExpanded",
			len(diagramprocessOther.ControlFlowsWhoseNodeIsExpanded),
			len(diagramprocess.ControlFlowsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.ControlFlowsWhoseNodeIsExpanded[i] == diagramprocess.ControlFlowsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.ControlFlowsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ControlFlow_ShapesDifferent := false
	if len(diagramprocess.ControlFlow_Shapes) != len(diagramprocessOther.ControlFlow_Shapes) {
		ControlFlow_ShapesDifferent = true
	} else {
		for i := range diagramprocess.ControlFlow_Shapes {
			if (diagramprocess.ControlFlow_Shapes[i] == nil) != (diagramprocessOther.ControlFlow_Shapes[i] == nil) {
				ControlFlow_ShapesDifferent = true
				break
			} else if diagramprocess.ControlFlow_Shapes[i] != nil && diagramprocessOther.ControlFlow_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.ControlFlow_Shapes[i] != diagramprocessOther.ControlFlow_Shapes[i] {
					ControlFlow_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ControlFlow_ShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"ControlFlow_Shapes",
			len(diagramprocessOther.ControlFlow_Shapes),
			len(diagramprocess.ControlFlow_Shapes),
			func(i, j int) bool {
				return diagramprocessOther.ControlFlow_Shapes[i] == diagramprocess.ControlFlow_Shapes[j]
			},
			func(j int) string {
				return diagramprocess.ControlFlow_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DataFlowsWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.DataFlowsWhoseNodeIsExpanded) != len(diagramprocessOther.DataFlowsWhoseNodeIsExpanded) {
		DataFlowsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
			if (diagramprocess.DataFlowsWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.DataFlowsWhoseNodeIsExpanded[i] == nil) {
				DataFlowsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.DataFlowsWhoseNodeIsExpanded[i] != nil && diagramprocessOther.DataFlowsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.DataFlowsWhoseNodeIsExpanded[i] != diagramprocessOther.DataFlowsWhoseNodeIsExpanded[i] {
					DataFlowsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DataFlowsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"DataFlowsWhoseNodeIsExpanded",
			len(diagramprocessOther.DataFlowsWhoseNodeIsExpanded),
			len(diagramprocess.DataFlowsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.DataFlowsWhoseNodeIsExpanded[i] == diagramprocess.DataFlowsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.DataFlowsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DataFlow_ShapesDifferent := false
	if len(diagramprocess.DataFlow_Shapes) != len(diagramprocessOther.DataFlow_Shapes) {
		DataFlow_ShapesDifferent = true
	} else {
		for i := range diagramprocess.DataFlow_Shapes {
			if (diagramprocess.DataFlow_Shapes[i] == nil) != (diagramprocessOther.DataFlow_Shapes[i] == nil) {
				DataFlow_ShapesDifferent = true
				break
			} else if diagramprocess.DataFlow_Shapes[i] != nil && diagramprocessOther.DataFlow_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.DataFlow_Shapes[i] != diagramprocessOther.DataFlow_Shapes[i] {
					DataFlow_ShapesDifferent = true
					break
				}
			}
		}
	}
	if DataFlow_ShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"DataFlow_Shapes",
			len(diagramprocessOther.DataFlow_Shapes),
			len(diagramprocess.DataFlow_Shapes),
			func(i, j int) bool {
				return diagramprocessOther.DataFlow_Shapes[i] == diagramprocess.DataFlow_Shapes[j]
			},
			func(j int) string {
				return diagramprocess.DataFlow_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DatasWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.DatasWhoseNodeIsExpanded) != len(diagramprocessOther.DatasWhoseNodeIsExpanded) {
		DatasWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.DatasWhoseNodeIsExpanded {
			if (diagramprocess.DatasWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.DatasWhoseNodeIsExpanded[i] == nil) {
				DatasWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.DatasWhoseNodeIsExpanded[i] != nil && diagramprocessOther.DatasWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.DatasWhoseNodeIsExpanded[i] != diagramprocessOther.DatasWhoseNodeIsExpanded[i] {
					DatasWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DatasWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"DatasWhoseNodeIsExpanded",
			len(diagramprocessOther.DatasWhoseNodeIsExpanded),
			len(diagramprocess.DatasWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.DatasWhoseNodeIsExpanded[i] == diagramprocess.DatasWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.DatasWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Data_ShapesDifferent := false
	if len(diagramprocess.Data_Shapes) != len(diagramprocessOther.Data_Shapes) {
		Data_ShapesDifferent = true
	} else {
		for i := range diagramprocess.Data_Shapes {
			if (diagramprocess.Data_Shapes[i] == nil) != (diagramprocessOther.Data_Shapes[i] == nil) {
				Data_ShapesDifferent = true
				break
			} else if diagramprocess.Data_Shapes[i] != nil && diagramprocessOther.Data_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.Data_Shapes[i] != diagramprocessOther.Data_Shapes[i] {
					Data_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Data_ShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"Data_Shapes",
			len(diagramprocessOther.Data_Shapes),
			len(diagramprocess.Data_Shapes),
			func(i, j int) bool {
				return diagramprocessOther.Data_Shapes[i] == diagramprocess.Data_Shapes[j]
			},
			func(j int) string {
				return diagramprocess.Data_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DataFlowsWhoseDataNodeIsExpandedDifferent := false
	if len(diagramprocess.DataFlowsWhoseDataNodeIsExpanded) != len(diagramprocessOther.DataFlowsWhoseDataNodeIsExpanded) {
		DataFlowsWhoseDataNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
			if (diagramprocess.DataFlowsWhoseDataNodeIsExpanded[i] == nil) != (diagramprocessOther.DataFlowsWhoseDataNodeIsExpanded[i] == nil) {
				DataFlowsWhoseDataNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.DataFlowsWhoseDataNodeIsExpanded[i] != nil && diagramprocessOther.DataFlowsWhoseDataNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.DataFlowsWhoseDataNodeIsExpanded[i] != diagramprocessOther.DataFlowsWhoseDataNodeIsExpanded[i] {
					DataFlowsWhoseDataNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DataFlowsWhoseDataNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"DataFlowsWhoseDataNodeIsExpanded",
			len(diagramprocessOther.DataFlowsWhoseDataNodeIsExpanded),
			len(diagramprocess.DataFlowsWhoseDataNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.DataFlowsWhoseDataNodeIsExpanded[i] == diagramprocess.DataFlowsWhoseDataNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.DataFlowsWhoseDataNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AllocatedResourcesWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.AllocatedResourcesWhoseNodeIsExpanded) != len(diagramprocessOther.AllocatedResourcesWhoseNodeIsExpanded) {
		AllocatedResourcesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
			if (diagramprocess.AllocatedResourcesWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.AllocatedResourcesWhoseNodeIsExpanded[i] == nil) {
				AllocatedResourcesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.AllocatedResourcesWhoseNodeIsExpanded[i] != nil && diagramprocessOther.AllocatedResourcesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.AllocatedResourcesWhoseNodeIsExpanded[i] != diagramprocessOther.AllocatedResourcesWhoseNodeIsExpanded[i] {
					AllocatedResourcesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if AllocatedResourcesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"AllocatedResourcesWhoseNodeIsExpanded",
			len(diagramprocessOther.AllocatedResourcesWhoseNodeIsExpanded),
			len(diagramprocess.AllocatedResourcesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.AllocatedResourcesWhoseNodeIsExpanded[i] == diagramprocess.AllocatedResourcesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.AllocatedResourcesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AllocatedResourceShapesDifferent := false
	if len(diagramprocess.AllocatedResourceShapes) != len(diagramprocessOther.AllocatedResourceShapes) {
		AllocatedResourceShapesDifferent = true
	} else {
		for i := range diagramprocess.AllocatedResourceShapes {
			if (diagramprocess.AllocatedResourceShapes[i] == nil) != (diagramprocessOther.AllocatedResourceShapes[i] == nil) {
				AllocatedResourceShapesDifferent = true
				break
			} else if diagramprocess.AllocatedResourceShapes[i] != nil && diagramprocessOther.AllocatedResourceShapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.AllocatedResourceShapes[i] != diagramprocessOther.AllocatedResourceShapes[i] {
					AllocatedResourceShapesDifferent = true
					break
				}
			}
		}
	}
	if AllocatedResourceShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"AllocatedResourceShapes",
			len(diagramprocessOther.AllocatedResourceShapes),
			len(diagramprocess.AllocatedResourceShapes),
			func(i, j int) bool {
				return diagramprocessOther.AllocatedResourceShapes[i] == diagramprocess.AllocatedResourceShapes[j]
			},
			func(j int) string {
				return diagramprocess.AllocatedResourceShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AllocatedProcessesWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.AllocatedProcessesWhoseNodeIsExpanded) != len(diagramprocessOther.AllocatedProcessesWhoseNodeIsExpanded) {
		AllocatedProcessesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
			if (diagramprocess.AllocatedProcessesWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.AllocatedProcessesWhoseNodeIsExpanded[i] == nil) {
				AllocatedProcessesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.AllocatedProcessesWhoseNodeIsExpanded[i] != nil && diagramprocessOther.AllocatedProcessesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.AllocatedProcessesWhoseNodeIsExpanded[i] != diagramprocessOther.AllocatedProcessesWhoseNodeIsExpanded[i] {
					AllocatedProcessesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if AllocatedProcessesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"AllocatedProcessesWhoseNodeIsExpanded",
			len(diagramprocessOther.AllocatedProcessesWhoseNodeIsExpanded),
			len(diagramprocess.AllocatedProcessesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.AllocatedProcessesWhoseNodeIsExpanded[i] == diagramprocess.AllocatedProcessesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.AllocatedProcessesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AllocatedProcessShapesDifferent := false
	if len(diagramprocess.AllocatedProcessShapes) != len(diagramprocessOther.AllocatedProcessShapes) {
		AllocatedProcessShapesDifferent = true
	} else {
		for i := range diagramprocess.AllocatedProcessShapes {
			if (diagramprocess.AllocatedProcessShapes[i] == nil) != (diagramprocessOther.AllocatedProcessShapes[i] == nil) {
				AllocatedProcessShapesDifferent = true
				break
			} else if diagramprocess.AllocatedProcessShapes[i] != nil && diagramprocessOther.AllocatedProcessShapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.AllocatedProcessShapes[i] != diagramprocessOther.AllocatedProcessShapes[i] {
					AllocatedProcessShapesDifferent = true
					break
				}
			}
		}
	}
	if AllocatedProcessShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"AllocatedProcessShapes",
			len(diagramprocessOther.AllocatedProcessShapes),
			len(diagramprocess.AllocatedProcessShapes),
			func(i, j int) bool {
				return diagramprocessOther.AllocatedProcessShapes[i] == diagramprocess.AllocatedProcessShapes[j]
			},
			func(j int) string {
				return diagramprocess.AllocatedProcessShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Note_ShapesDifferent := false
	if len(diagramprocess.Note_Shapes) != len(diagramprocessOther.Note_Shapes) {
		Note_ShapesDifferent = true
	} else {
		for i := range diagramprocess.Note_Shapes {
			if (diagramprocess.Note_Shapes[i] == nil) != (diagramprocessOther.Note_Shapes[i] == nil) {
				Note_ShapesDifferent = true
				break
			} else if diagramprocess.Note_Shapes[i] != nil && diagramprocessOther.Note_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.Note_Shapes[i] != diagramprocessOther.Note_Shapes[i] {
					Note_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Note_ShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"Note_Shapes",
			len(diagramprocessOther.Note_Shapes),
			len(diagramprocess.Note_Shapes),
			func(i, j int) bool {
				return diagramprocessOther.Note_Shapes[i] == diagramprocess.Note_Shapes[j]
			},
			func(j int) string {
				return diagramprocess.Note_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(diagramprocess.NotesWhoseNodeIsExpanded) != len(diagramprocessOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramprocess.NotesWhoseNodeIsExpanded {
			if (diagramprocess.NotesWhoseNodeIsExpanded[i] == nil) != (diagramprocessOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramprocess.NotesWhoseNodeIsExpanded[i] != nil && diagramprocessOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.NotesWhoseNodeIsExpanded[i] != diagramprocessOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramprocess,
			"NotesWhoseNodeIsExpanded",
			len(diagramprocessOther.NotesWhoseNodeIsExpanded),
			len(diagramprocess.NotesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramprocessOther.NotesWhoseNodeIsExpanded[i] == diagramprocess.NotesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramprocess.NotesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagramprocess.IsNotesNodeExpanded != diagramprocessOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NoteTaskShapesDifferent := false
	if len(diagramprocess.NoteTaskShapes) != len(diagramprocessOther.NoteTaskShapes) {
		NoteTaskShapesDifferent = true
	} else {
		for i := range diagramprocess.NoteTaskShapes {
			if (diagramprocess.NoteTaskShapes[i] == nil) != (diagramprocessOther.NoteTaskShapes[i] == nil) {
				NoteTaskShapesDifferent = true
				break
			} else if diagramprocess.NoteTaskShapes[i] != nil && diagramprocessOther.NoteTaskShapes[i] != nil {
				// this is a pointer comparaison
				if diagramprocess.NoteTaskShapes[i] != diagramprocessOther.NoteTaskShapes[i] {
					NoteTaskShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteTaskShapesDifferent {
		ops := stage.Diff(
			diagramprocess,
			"NoteTaskShapes",
			len(diagramprocessOther.NoteTaskShapes),
			len(diagramprocess.NoteTaskShapes),
			func(i, j int) bool {
				return diagramprocessOther.NoteTaskShapes[i] == diagramprocess.NoteTaskShapes[j]
			},
			func(j int) string {
				return diagramprocess.NoteTaskShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (externalparticipantshape *ExternalParticipantShape) GongDiff(stage *Stage, externalparticipantshapeOther *ExternalParticipantShape) (diffs []string) {
	// insertion point for field diffs
	if externalparticipantshape.Name != externalparticipantshapeOther.Name {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "Name"))
	}
	if (externalparticipantshape.Participant == nil) != (externalparticipantshapeOther.Participant == nil) {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "Participant"))
	} else if externalparticipantshape.Participant != nil && externalparticipantshapeOther.Participant != nil {
		if externalparticipantshape.Participant != externalparticipantshapeOther.Participant {
			diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "Participant"))
		}
	}
	if externalparticipantshape.IsExpanded != externalparticipantshapeOther.IsExpanded {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "IsExpanded"))
	}
	if externalparticipantshape.X != externalparticipantshapeOther.X {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "X"))
	}
	if externalparticipantshape.Y != externalparticipantshapeOther.Y {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "Y"))
	}
	if externalparticipantshape.Width != externalparticipantshapeOther.Width {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "Width"))
	}
	if externalparticipantshape.Height != externalparticipantshapeOther.Height {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "Height"))
	}
	if externalparticipantshape.IsHidden != externalparticipantshapeOther.IsHidden {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "IsHidden"))
	}
	if externalparticipantshape.TailHeigth != externalparticipantshapeOther.TailHeigth {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "TailHeigth"))
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
		ops := stage.Diff(
			library,
			"SubLibraries",
			len(libraryOther.SubLibraries),
			len(library.SubLibraries),
			func(i, j int) bool {
				return libraryOther.SubLibraries[i] == library.SubLibraries[j]
			},
			func(j int) string {
				return library.SubLibraries[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			library,
			"SubLibrariesWhoseNodeIsExpanded",
			len(libraryOther.SubLibrariesWhoseNodeIsExpanded),
			len(library.SubLibrariesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.SubLibrariesWhoseNodeIsExpanded[i] == library.SubLibrariesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.SubLibrariesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			library,
			"RootProcesses",
			len(libraryOther.RootProcesses),
			len(library.RootProcesses),
			func(i, j int) bool {
				return libraryOther.RootProcesses[i] == library.RootProcesses[j]
			},
			func(j int) string {
				return library.RootProcesses[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsProcessesNodeExpanded != libraryOther.IsProcessesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsProcessesNodeExpanded"))
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
		ops := stage.Diff(
			library,
			"ProcesssWhoseNodeIsExpanded",
			len(libraryOther.ProcesssWhoseNodeIsExpanded),
			len(library.ProcesssWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.ProcesssWhoseNodeIsExpanded[i] == library.ProcesssWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.ProcesssWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootDataFlowsDifferent := false
	if len(library.RootDataFlows) != len(libraryOther.RootDataFlows) {
		RootDataFlowsDifferent = true
	} else {
		for i := range library.RootDataFlows {
			if (library.RootDataFlows[i] == nil) != (libraryOther.RootDataFlows[i] == nil) {
				RootDataFlowsDifferent = true
				break
			} else if library.RootDataFlows[i] != nil && libraryOther.RootDataFlows[i] != nil {
				// this is a pointer comparaison
				if library.RootDataFlows[i] != libraryOther.RootDataFlows[i] {
					RootDataFlowsDifferent = true
					break
				}
			}
		}
	}
	if RootDataFlowsDifferent {
		ops := stage.Diff(
			library,
			"RootDataFlows",
			len(libraryOther.RootDataFlows),
			len(library.RootDataFlows),
			func(i, j int) bool {
				return libraryOther.RootDataFlows[i] == library.RootDataFlows[j]
			},
			func(j int) string {
				return library.RootDataFlows[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsDataFlowsNodeExpanded != libraryOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	DataFlowsWhoseNodeIsExpandedDifferent := false
	if len(library.DataFlowsWhoseNodeIsExpanded) != len(libraryOther.DataFlowsWhoseNodeIsExpanded) {
		DataFlowsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.DataFlowsWhoseNodeIsExpanded {
			if (library.DataFlowsWhoseNodeIsExpanded[i] == nil) != (libraryOther.DataFlowsWhoseNodeIsExpanded[i] == nil) {
				DataFlowsWhoseNodeIsExpandedDifferent = true
				break
			} else if library.DataFlowsWhoseNodeIsExpanded[i] != nil && libraryOther.DataFlowsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.DataFlowsWhoseNodeIsExpanded[i] != libraryOther.DataFlowsWhoseNodeIsExpanded[i] {
					DataFlowsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DataFlowsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"DataFlowsWhoseNodeIsExpanded",
			len(libraryOther.DataFlowsWhoseNodeIsExpanded),
			len(library.DataFlowsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.DataFlowsWhoseNodeIsExpanded[i] == library.DataFlowsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.DataFlowsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootDatasDifferent := false
	if len(library.RootDatas) != len(libraryOther.RootDatas) {
		RootDatasDifferent = true
	} else {
		for i := range library.RootDatas {
			if (library.RootDatas[i] == nil) != (libraryOther.RootDatas[i] == nil) {
				RootDatasDifferent = true
				break
			} else if library.RootDatas[i] != nil && libraryOther.RootDatas[i] != nil {
				// this is a pointer comparaison
				if library.RootDatas[i] != libraryOther.RootDatas[i] {
					RootDatasDifferent = true
					break
				}
			}
		}
	}
	if RootDatasDifferent {
		ops := stage.Diff(
			library,
			"RootDatas",
			len(libraryOther.RootDatas),
			len(library.RootDatas),
			func(i, j int) bool {
				return libraryOther.RootDatas[i] == library.RootDatas[j]
			},
			func(j int) string {
				return library.RootDatas[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsDatasNodeExpanded != libraryOther.IsDatasNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsDatasNodeExpanded"))
	}
	DatasWhoseNodeIsExpandedDifferent := false
	if len(library.DatasWhoseNodeIsExpanded) != len(libraryOther.DatasWhoseNodeIsExpanded) {
		DatasWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.DatasWhoseNodeIsExpanded {
			if (library.DatasWhoseNodeIsExpanded[i] == nil) != (libraryOther.DatasWhoseNodeIsExpanded[i] == nil) {
				DatasWhoseNodeIsExpandedDifferent = true
				break
			} else if library.DatasWhoseNodeIsExpanded[i] != nil && libraryOther.DatasWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.DatasWhoseNodeIsExpanded[i] != libraryOther.DatasWhoseNodeIsExpanded[i] {
					DatasWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DatasWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"DatasWhoseNodeIsExpanded",
			len(libraryOther.DatasWhoseNodeIsExpanded),
			len(library.DatasWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.DatasWhoseNodeIsExpanded[i] == library.DatasWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.DatasWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootResourcesDifferent := false
	if len(library.RootResources) != len(libraryOther.RootResources) {
		RootResourcesDifferent = true
	} else {
		for i := range library.RootResources {
			if (library.RootResources[i] == nil) != (libraryOther.RootResources[i] == nil) {
				RootResourcesDifferent = true
				break
			} else if library.RootResources[i] != nil && libraryOther.RootResources[i] != nil {
				// this is a pointer comparaison
				if library.RootResources[i] != libraryOther.RootResources[i] {
					RootResourcesDifferent = true
					break
				}
			}
		}
	}
	if RootResourcesDifferent {
		ops := stage.Diff(
			library,
			"RootResources",
			len(libraryOther.RootResources),
			len(library.RootResources),
			func(i, j int) bool {
				return libraryOther.RootResources[i] == library.RootResources[j]
			},
			func(j int) string {
				return library.RootResources[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsResourcesNodeExpanded != libraryOther.IsResourcesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	ResourcesWhoseNodeIsExpandedDifferent := false
	if len(library.ResourcesWhoseNodeIsExpanded) != len(libraryOther.ResourcesWhoseNodeIsExpanded) {
		ResourcesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.ResourcesWhoseNodeIsExpanded {
			if (library.ResourcesWhoseNodeIsExpanded[i] == nil) != (libraryOther.ResourcesWhoseNodeIsExpanded[i] == nil) {
				ResourcesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.ResourcesWhoseNodeIsExpanded[i] != nil && libraryOther.ResourcesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.ResourcesWhoseNodeIsExpanded[i] != libraryOther.ResourcesWhoseNodeIsExpanded[i] {
					ResourcesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ResourcesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"ResourcesWhoseNodeIsExpanded",
			len(libraryOther.ResourcesWhoseNodeIsExpanded),
			len(library.ResourcesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.ResourcesWhoseNodeIsExpanded[i] == library.ResourcesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.ResourcesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ParticipantsWhoseNodeIsExpandedDifferent := false
	if len(library.ParticipantsWhoseNodeIsExpanded) != len(libraryOther.ParticipantsWhoseNodeIsExpanded) {
		ParticipantsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.ParticipantsWhoseNodeIsExpanded {
			if (library.ParticipantsWhoseNodeIsExpanded[i] == nil) != (libraryOther.ParticipantsWhoseNodeIsExpanded[i] == nil) {
				ParticipantsWhoseNodeIsExpandedDifferent = true
				break
			} else if library.ParticipantsWhoseNodeIsExpanded[i] != nil && libraryOther.ParticipantsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.ParticipantsWhoseNodeIsExpanded[i] != libraryOther.ParticipantsWhoseNodeIsExpanded[i] {
					ParticipantsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ParticipantsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"ParticipantsWhoseNodeIsExpanded",
			len(libraryOther.ParticipantsWhoseNodeIsExpanded),
			len(library.ParticipantsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.ParticipantsWhoseNodeIsExpanded[i] == library.ParticipantsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.ParticipantsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootNotesDifferent := false
	if len(library.RootNotes) != len(libraryOther.RootNotes) {
		RootNotesDifferent = true
	} else {
		for i := range library.RootNotes {
			if (library.RootNotes[i] == nil) != (libraryOther.RootNotes[i] == nil) {
				RootNotesDifferent = true
				break
			} else if library.RootNotes[i] != nil && libraryOther.RootNotes[i] != nil {
				// this is a pointer comparaison
				if library.RootNotes[i] != libraryOther.RootNotes[i] {
					RootNotesDifferent = true
					break
				}
			}
		}
	}
	if RootNotesDifferent {
		ops := stage.Diff(
			library,
			"RootNotes",
			len(libraryOther.RootNotes),
			len(library.RootNotes),
			func(i, j int) bool {
				return libraryOther.RootNotes[i] == library.RootNotes[j]
			},
			func(j int) string {
				return library.RootNotes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsNotesNodeExpanded != libraryOther.IsNotesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(library.NotesWhoseNodeIsExpanded) != len(libraryOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.NotesWhoseNodeIsExpanded {
			if (library.NotesWhoseNodeIsExpanded[i] == nil) != (libraryOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.NotesWhoseNodeIsExpanded[i] != nil && libraryOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.NotesWhoseNodeIsExpanded[i] != libraryOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"NotesWhoseNodeIsExpanded",
			len(libraryOther.NotesWhoseNodeIsExpanded),
			len(library.NotesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.NotesWhoseNodeIsExpanded[i] == library.NotesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.NotesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if note.Description != noteOther.Description {
		diffs = append(diffs, note.GongMarshallField(stage, "Description"))
	}
	if note.ComputedPrefix != noteOther.ComputedPrefix {
		diffs = append(diffs, note.GongMarshallField(stage, "ComputedPrefix"))
	}
	if note.IsExpanded != noteOther.IsExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsExpanded"))
	}
	if note.IsTasksNodeExpanded != noteOther.IsTasksNodeExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsTasksNodeExpanded"))
	}
	TasksDifferent := false
	if len(note.Tasks) != len(noteOther.Tasks) {
		TasksDifferent = true
	} else {
		for i := range note.Tasks {
			if (note.Tasks[i] == nil) != (noteOther.Tasks[i] == nil) {
				TasksDifferent = true
				break
			} else if note.Tasks[i] != nil && noteOther.Tasks[i] != nil {
				// this is a pointer comparaison
				if note.Tasks[i] != noteOther.Tasks[i] {
					TasksDifferent = true
					break
				}
			}
		}
	}
	if TasksDifferent {
		ops := stage.Diff(
			note,
			"Tasks",
			len(noteOther.Tasks),
			len(note.Tasks),
			func(i, j int) bool {
				return noteOther.Tasks[i] == note.Tasks[j]
			},
			func(j int) string {
				return note.Tasks[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteshape *NoteShape) GongDiff(stage *Stage, noteshapeOther *NoteShape) (diffs []string) {
	// insertion point for field diffs
	if noteshape.Name != noteshapeOther.Name {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Name"))
	}
	if (noteshape.Note == nil) != (noteshapeOther.Note == nil) {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
	} else if noteshape.Note != nil && noteshapeOther.Note != nil {
		if noteshape.Note != noteshapeOther.Note {
			diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
		}
	}
	if noteshape.X != noteshapeOther.X {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "X"))
	}
	if noteshape.Y != noteshapeOther.Y {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Y"))
	}
	if noteshape.Width != noteshapeOther.Width {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Width"))
	}
	if noteshape.Height != noteshapeOther.Height {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Height"))
	}
	if noteshape.IsHidden != noteshapeOther.IsHidden {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notetaskshape *NoteTaskShape) GongDiff(stage *Stage, notetaskshapeOther *NoteTaskShape) (diffs []string) {
	// insertion point for field diffs
	if notetaskshape.Name != notetaskshapeOther.Name {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Name"))
	}
	if (notetaskshape.Note == nil) != (notetaskshapeOther.Note == nil) {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
	} else if notetaskshape.Note != nil && notetaskshapeOther.Note != nil {
		if notetaskshape.Note != notetaskshapeOther.Note {
			diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
		}
	}
	if (notetaskshape.Task == nil) != (notetaskshapeOther.Task == nil) {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
	} else if notetaskshape.Task != nil && notetaskshapeOther.Task != nil {
		if notetaskshape.Task != notetaskshapeOther.Task {
			diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
		}
	}
	if notetaskshape.StartRatio != notetaskshapeOther.StartRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "StartRatio"))
	}
	if notetaskshape.EndRatio != notetaskshapeOther.EndRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "EndRatio"))
	}
	if notetaskshape.StartOrientation != notetaskshapeOther.StartOrientation {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "StartOrientation"))
	}
	if notetaskshape.EndOrientation != notetaskshapeOther.EndOrientation {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "EndOrientation"))
	}
	if notetaskshape.CornerOffsetRatio != notetaskshapeOther.CornerOffsetRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notetaskshape.IsHidden != notetaskshapeOther.IsHidden {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "IsHidden"))
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
	if participant.IsProcessResource != participantOther.IsProcessResource {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsProcessResource"))
	}
	if participant.Description != participantOther.Description {
		diffs = append(diffs, participant.GongMarshallField(stage, "Description"))
	}
	ResourcesDifferent := false
	if len(participant.Resources) != len(participantOther.Resources) {
		ResourcesDifferent = true
	} else {
		for i := range participant.Resources {
			if (participant.Resources[i] == nil) != (participantOther.Resources[i] == nil) {
				ResourcesDifferent = true
				break
			} else if participant.Resources[i] != nil && participantOther.Resources[i] != nil {
				// this is a pointer comparaison
				if participant.Resources[i] != participantOther.Resources[i] {
					ResourcesDifferent = true
					break
				}
			}
		}
	}
	if ResourcesDifferent {
		ops := stage.Diff(
			participant,
			"Resources",
			len(participantOther.Resources),
			len(participant.Resources),
			func(i, j int) bool {
				return participantOther.Resources[i] == participant.Resources[j]
			},
			func(j int) string {
				return participant.Resources[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if participant.IsResourcesNodeExpanded != participantOther.IsResourcesNodeExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	ProcessesDifferent := false
	if len(participant.Processes) != len(participantOther.Processes) {
		ProcessesDifferent = true
	} else {
		for i := range participant.Processes {
			if (participant.Processes[i] == nil) != (participantOther.Processes[i] == nil) {
				ProcessesDifferent = true
				break
			} else if participant.Processes[i] != nil && participantOther.Processes[i] != nil {
				// this is a pointer comparaison
				if participant.Processes[i] != participantOther.Processes[i] {
					ProcessesDifferent = true
					break
				}
			}
		}
	}
	if ProcessesDifferent {
		ops := stage.Diff(
			participant,
			"Processes",
			len(participantOther.Processes),
			len(participant.Processes),
			func(i, j int) bool {
				return participantOther.Processes[i] == participant.Processes[j]
			},
			func(j int) string {
				return participant.Processes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if participant.IsProcessesNodeExpanded != participantOther.IsProcessesNodeExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsProcessesNodeExpanded"))
	}
	if participant.ComputedPrefix != participantOther.ComputedPrefix {
		diffs = append(diffs, participant.GongMarshallField(stage, "ComputedPrefix"))
	}
	if participant.IsExpanded != participantOther.IsExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsExpanded"))
	}
	if participant.IsTasksNodeExpanded != participantOther.IsTasksNodeExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsTasksNodeExpanded"))
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
		ops := stage.Diff(
			participant,
			"Tasks",
			len(participantOther.Tasks),
			len(participant.Tasks),
			func(i, j int) bool {
				return participantOther.Tasks[i] == participant.Tasks[j]
			},
			func(j int) string {
				return participant.Tasks[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if participant.IsControlFlowsNodeExpanded != participantOther.IsControlFlowsNodeExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsControlFlowsNodeExpanded"))
	}
	ControlFlowsDifferent := false
	if len(participant.ControlFlows) != len(participantOther.ControlFlows) {
		ControlFlowsDifferent = true
	} else {
		for i := range participant.ControlFlows {
			if (participant.ControlFlows[i] == nil) != (participantOther.ControlFlows[i] == nil) {
				ControlFlowsDifferent = true
				break
			} else if participant.ControlFlows[i] != nil && participantOther.ControlFlows[i] != nil {
				// this is a pointer comparaison
				if participant.ControlFlows[i] != participantOther.ControlFlows[i] {
					ControlFlowsDifferent = true
					break
				}
			}
		}
	}
	if ControlFlowsDifferent {
		ops := stage.Diff(
			participant,
			"ControlFlows",
			len(participantOther.ControlFlows),
			len(participant.ControlFlows),
			func(i, j int) bool {
				return participantOther.ControlFlows[i] == participant.ControlFlows[j]
			},
			func(j int) string {
				return participant.ControlFlows[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TaskWhoseOutControlFlowsNodeIsExpandedDifferent := false
	if len(participant.TaskWhoseOutControlFlowsNodeIsExpanded) != len(participantOther.TaskWhoseOutControlFlowsNodeIsExpanded) {
		TaskWhoseOutControlFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
			if (participant.TaskWhoseOutControlFlowsNodeIsExpanded[i] == nil) != (participantOther.TaskWhoseOutControlFlowsNodeIsExpanded[i] == nil) {
				TaskWhoseOutControlFlowsNodeIsExpandedDifferent = true
				break
			} else if participant.TaskWhoseOutControlFlowsNodeIsExpanded[i] != nil && participantOther.TaskWhoseOutControlFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if participant.TaskWhoseOutControlFlowsNodeIsExpanded[i] != participantOther.TaskWhoseOutControlFlowsNodeIsExpanded[i] {
					TaskWhoseOutControlFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TaskWhoseOutControlFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			participant,
			"TaskWhoseOutControlFlowsNodeIsExpanded",
			len(participantOther.TaskWhoseOutControlFlowsNodeIsExpanded),
			len(participant.TaskWhoseOutControlFlowsNodeIsExpanded),
			func(i, j int) bool {
				return participantOther.TaskWhoseOutControlFlowsNodeIsExpanded[i] == participant.TaskWhoseOutControlFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return participant.TaskWhoseOutControlFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TaskWhoseInControlFlowsNodeIsExpandedDifferent := false
	if len(participant.TaskWhoseInControlFlowsNodeIsExpanded) != len(participantOther.TaskWhoseInControlFlowsNodeIsExpanded) {
		TaskWhoseInControlFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
			if (participant.TaskWhoseInControlFlowsNodeIsExpanded[i] == nil) != (participantOther.TaskWhoseInControlFlowsNodeIsExpanded[i] == nil) {
				TaskWhoseInControlFlowsNodeIsExpandedDifferent = true
				break
			} else if participant.TaskWhoseInControlFlowsNodeIsExpanded[i] != nil && participantOther.TaskWhoseInControlFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if participant.TaskWhoseInControlFlowsNodeIsExpanded[i] != participantOther.TaskWhoseInControlFlowsNodeIsExpanded[i] {
					TaskWhoseInControlFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TaskWhoseInControlFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			participant,
			"TaskWhoseInControlFlowsNodeIsExpanded",
			len(participantOther.TaskWhoseInControlFlowsNodeIsExpanded),
			len(participant.TaskWhoseInControlFlowsNodeIsExpanded),
			func(i, j int) bool {
				return participantOther.TaskWhoseInControlFlowsNodeIsExpanded[i] == participant.TaskWhoseInControlFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return participant.TaskWhoseInControlFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if participant.IsDataFlowsNodeExpanded != participantOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	TaskWhoseOutDataFlowsNodeIsExpandedDifferent := false
	if len(participant.TaskWhoseOutDataFlowsNodeIsExpanded) != len(participantOther.TaskWhoseOutDataFlowsNodeIsExpanded) {
		TaskWhoseOutDataFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
			if (participant.TaskWhoseOutDataFlowsNodeIsExpanded[i] == nil) != (participantOther.TaskWhoseOutDataFlowsNodeIsExpanded[i] == nil) {
				TaskWhoseOutDataFlowsNodeIsExpandedDifferent = true
				break
			} else if participant.TaskWhoseOutDataFlowsNodeIsExpanded[i] != nil && participantOther.TaskWhoseOutDataFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if participant.TaskWhoseOutDataFlowsNodeIsExpanded[i] != participantOther.TaskWhoseOutDataFlowsNodeIsExpanded[i] {
					TaskWhoseOutDataFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TaskWhoseOutDataFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			participant,
			"TaskWhoseOutDataFlowsNodeIsExpanded",
			len(participantOther.TaskWhoseOutDataFlowsNodeIsExpanded),
			len(participant.TaskWhoseOutDataFlowsNodeIsExpanded),
			func(i, j int) bool {
				return participantOther.TaskWhoseOutDataFlowsNodeIsExpanded[i] == participant.TaskWhoseOutDataFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return participant.TaskWhoseOutDataFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TaskWhoseInDataFlowsNodeIsExpandedDifferent := false
	if len(participant.TaskWhoseInDataFlowsNodeIsExpanded) != len(participantOther.TaskWhoseInDataFlowsNodeIsExpanded) {
		TaskWhoseInDataFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
			if (participant.TaskWhoseInDataFlowsNodeIsExpanded[i] == nil) != (participantOther.TaskWhoseInDataFlowsNodeIsExpanded[i] == nil) {
				TaskWhoseInDataFlowsNodeIsExpandedDifferent = true
				break
			} else if participant.TaskWhoseInDataFlowsNodeIsExpanded[i] != nil && participantOther.TaskWhoseInDataFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if participant.TaskWhoseInDataFlowsNodeIsExpanded[i] != participantOther.TaskWhoseInDataFlowsNodeIsExpanded[i] {
					TaskWhoseInDataFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TaskWhoseInDataFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			participant,
			"TaskWhoseInDataFlowsNodeIsExpanded",
			len(participantOther.TaskWhoseInDataFlowsNodeIsExpanded),
			len(participant.TaskWhoseInDataFlowsNodeIsExpanded),
			func(i, j int) bool {
				return participantOther.TaskWhoseInDataFlowsNodeIsExpanded[i] == participant.TaskWhoseInDataFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return participant.TaskWhoseInDataFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
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
	if participantshape.WidthWeight != participantshapeOther.WidthWeight {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "WidthWeight"))
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
	if process.Description != processOther.Description {
		diffs = append(diffs, process.GongMarshallField(stage, "Description"))
	}
	if process.ComputedPrefix != processOther.ComputedPrefix {
		diffs = append(diffs, process.GongMarshallField(stage, "ComputedPrefix"))
	}
	if process.IsExpanded != processOther.IsExpanded {
		diffs = append(diffs, process.GongMarshallField(stage, "IsExpanded"))
	}
	if process.SVG_Path != processOther.SVG_Path {
		diffs = append(diffs, process.GongMarshallField(stage, "SVG_Path"))
	}
	if process.InverseAppliedScaling != processOther.InverseAppliedScaling {
		diffs = append(diffs, process.GongMarshallField(stage, "InverseAppliedScaling"))
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
		ops := stage.Diff(
			process,
			"DiagramProcesss",
			len(processOther.DiagramProcesss),
			len(process.DiagramProcesss),
			func(i, j int) bool {
				return processOther.DiagramProcesss[i] == process.DiagramProcesss[j]
			},
			func(j int) string {
				return process.DiagramProcesss[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			process,
			"DiagramProcessWhoseNodeIsExpanded",
			len(processOther.DiagramProcessWhoseNodeIsExpanded),
			len(process.DiagramProcessWhoseNodeIsExpanded),
			func(i, j int) bool {
				return processOther.DiagramProcessWhoseNodeIsExpanded[i] == process.DiagramProcessWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return process.DiagramProcessWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			process,
			"SubProcesses",
			len(processOther.SubProcesses),
			len(process.SubProcesses),
			func(i, j int) bool {
				return processOther.SubProcesses[i] == process.SubProcesses[j]
			},
			func(j int) string {
				return process.SubProcesses[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			process,
			"Participants",
			len(processOther.Participants),
			len(process.Participants),
			func(i, j int) bool {
				return processOther.Participants[i] == process.Participants[j]
			},
			func(j int) string {
				return process.Participants[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			process,
			"ParticipantWhoseNodeIsExpanded",
			len(processOther.ParticipantWhoseNodeIsExpanded),
			len(process.ParticipantWhoseNodeIsExpanded),
			func(i, j int) bool {
				return processOther.ParticipantWhoseNodeIsExpanded[i] == process.ParticipantWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return process.ParticipantWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DataFlowsDifferent := false
	if len(process.DataFlows) != len(processOther.DataFlows) {
		DataFlowsDifferent = true
	} else {
		for i := range process.DataFlows {
			if (process.DataFlows[i] == nil) != (processOther.DataFlows[i] == nil) {
				DataFlowsDifferent = true
				break
			} else if process.DataFlows[i] != nil && processOther.DataFlows[i] != nil {
				// this is a pointer comparaison
				if process.DataFlows[i] != processOther.DataFlows[i] {
					DataFlowsDifferent = true
					break
				}
			}
		}
	}
	if DataFlowsDifferent {
		ops := stage.Diff(
			process,
			"DataFlows",
			len(processOther.DataFlows),
			len(process.DataFlows),
			func(i, j int) bool {
				return processOther.DataFlows[i] == process.DataFlows[j]
			},
			func(j int) string {
				return process.DataFlows[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if process.IsDataFlowsNodeExpanded != processOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, process.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	ExternalParticipantsDifferent := false
	if len(process.ExternalParticipants) != len(processOther.ExternalParticipants) {
		ExternalParticipantsDifferent = true
	} else {
		for i := range process.ExternalParticipants {
			if (process.ExternalParticipants[i] == nil) != (processOther.ExternalParticipants[i] == nil) {
				ExternalParticipantsDifferent = true
				break
			} else if process.ExternalParticipants[i] != nil && processOther.ExternalParticipants[i] != nil {
				// this is a pointer comparaison
				if process.ExternalParticipants[i] != processOther.ExternalParticipants[i] {
					ExternalParticipantsDifferent = true
					break
				}
			}
		}
	}
	if ExternalParticipantsDifferent {
		ops := stage.Diff(
			process,
			"ExternalParticipants",
			len(processOther.ExternalParticipants),
			len(process.ExternalParticipants),
			func(i, j int) bool {
				return processOther.ExternalParticipants[i] == process.ExternalParticipants[j]
			},
			func(j int) string {
				return process.ExternalParticipants[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ExternalParticipantWhoseNodeIsExpandedDifferent := false
	if len(process.ExternalParticipantWhoseNodeIsExpanded) != len(processOther.ExternalParticipantWhoseNodeIsExpanded) {
		ExternalParticipantWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range process.ExternalParticipantWhoseNodeIsExpanded {
			if (process.ExternalParticipantWhoseNodeIsExpanded[i] == nil) != (processOther.ExternalParticipantWhoseNodeIsExpanded[i] == nil) {
				ExternalParticipantWhoseNodeIsExpandedDifferent = true
				break
			} else if process.ExternalParticipantWhoseNodeIsExpanded[i] != nil && processOther.ExternalParticipantWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if process.ExternalParticipantWhoseNodeIsExpanded[i] != processOther.ExternalParticipantWhoseNodeIsExpanded[i] {
					ExternalParticipantWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ExternalParticipantWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			process,
			"ExternalParticipantWhoseNodeIsExpanded",
			len(processOther.ExternalParticipantWhoseNodeIsExpanded),
			len(process.ExternalParticipantWhoseNodeIsExpanded),
			func(i, j int) bool {
				return processOther.ExternalParticipantWhoseNodeIsExpanded[i] == process.ExternalParticipantWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return process.ExternalParticipantWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
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
func (resource *Resource) GongDiff(stage *Stage, resourceOther *Resource) (diffs []string) {
	// insertion point for field diffs
	if resource.Name != resourceOther.Name {
		diffs = append(diffs, resource.GongMarshallField(stage, "Name"))
	}
	if resource.Acronym != resourceOther.Acronym {
		diffs = append(diffs, resource.GongMarshallField(stage, "Acronym"))
	}
	if resource.Description != resourceOther.Description {
		diffs = append(diffs, resource.GongMarshallField(stage, "Description"))
	}
	if resource.ComputedPrefix != resourceOther.ComputedPrefix {
		diffs = append(diffs, resource.GongMarshallField(stage, "ComputedPrefix"))
	}
	if resource.IsExpanded != resourceOther.IsExpanded {
		diffs = append(diffs, resource.GongMarshallField(stage, "IsExpanded"))
	}
	if resource.SVG_Path != resourceOther.SVG_Path {
		diffs = append(diffs, resource.GongMarshallField(stage, "SVG_Path"))
	}
	if resource.InverseAppliedScaling != resourceOther.InverseAppliedScaling {
		diffs = append(diffs, resource.GongMarshallField(stage, "InverseAppliedScaling"))
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
	if task.Description != taskOther.Description {
		diffs = append(diffs, task.GongMarshallField(stage, "Description"))
	}
	if task.ComputedPrefix != taskOther.ComputedPrefix {
		diffs = append(diffs, task.GongMarshallField(stage, "ComputedPrefix"))
	}
	if task.IsExpanded != taskOther.IsExpanded {
		diffs = append(diffs, task.GongMarshallField(stage, "IsExpanded"))
	}
	if task.IsStartTask != taskOther.IsStartTask {
		diffs = append(diffs, task.GongMarshallField(stage, "IsStartTask"))
	}
	if task.IsEndTask != taskOther.IsEndTask {
		diffs = append(diffs, task.GongMarshallField(stage, "IsEndTask"))
	}
	if (task.Type == nil) != (taskOther.Type == nil) {
		diffs = append(diffs, task.GongMarshallField(stage, "Type"))
	} else if task.Type != nil && taskOther.Type != nil {
		if task.Type != taskOther.Type {
			diffs = append(diffs, task.GongMarshallField(stage, "Type"))
		}
	}
	if task.IsTaskNameNotProcessName != taskOther.IsTaskNameNotProcessName {
		diffs = append(diffs, task.GongMarshallField(stage, "IsTaskNameNotProcessName"))
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

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if equal(i, j) {
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
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := 0; k < n; k++ {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
