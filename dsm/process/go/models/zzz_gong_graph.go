// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (allocatedprocessshape *AllocatedProcessShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AllocatedProcessShapes[allocatedprocessshape]
	return ok
}

func (allocatedresourceshape *AllocatedResourceShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AllocatedResourceShapes[allocatedresourceshape]
	return ok
}

func (controlflow *ControlFlow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ControlFlows[controlflow]
	return ok
}

func (controlflowshape *ControlFlowShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ControlFlowShapes[controlflowshape]
	return ok
}

func (data *Data) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Datas[data]
	return ok
}

func (dataflow *DataFlow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DataFlows[dataflow]
	return ok
}

func (dataflowshape *DataFlowShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DataFlowShapes[dataflowshape]
	return ok
}

func (datashape *DataShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DataShapes[datashape]
	return ok
}

func (diagramprocess *DiagramProcess) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DiagramProcesss[diagramprocess]
	return ok
}

func (externalparticipantshape *ExternalParticipantShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ExternalParticipantShapes[externalparticipantshape]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteShapes[noteshape]
	return ok
}

func (notetaskshape *NoteTaskShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteTaskShapes[notetaskshape]
	return ok
}

func (participant *Participant) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Participants[participant]
	return ok
}

func (participantshape *ParticipantShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ParticipantShapes[participantshape]
	return ok
}

func (process *Process) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Processs[process]
	return ok
}

func (processshape *ProcessShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ProcessShapes[processshape]
	return ok
}

func (resource *Resource) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Resources[resource]
	return ok
}

func (task *Task) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tasks[task]
	return ok
}

func (taskshape *TaskShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TaskShapes[taskshape]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (allocatedprocessshape *AllocatedProcessShape) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(data) {
		return
	}

	data.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dataflow *DataFlow) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(resource) {
		return
	}

	resource.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (task *Task) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	allocatedprocessshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, allocatedprocessshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	allocatedresourceshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, allocatedresourceshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	controlflowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, controlflowFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	controlflowshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, controlflowshapeFrom)
	if alreadyCopied {
		return
	}
	controlflowshapeFrom.GongCopyBasicFields(controlflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshapeFrom.ControlFlow != nil {
		controlflowshapeTo.ControlFlow = GongCopyBranchControlFlow(mapOrigCopy, controlflowshapeFrom.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchData(mapOrigCopy map[any]any, dataFrom *Data) (dataTo *Data) {
	var alreadyCopied bool
	dataTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dataFrom)
	if alreadyCopied {
		return
	}
	dataFrom.GongCopyBasicFields(dataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDataFlow(mapOrigCopy map[any]any, dataflowFrom *DataFlow) (dataflowTo *DataFlow) {
	var alreadyCopied bool
	dataflowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dataflowFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	dataflowshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dataflowshapeFrom)
	if alreadyCopied {
		return
	}
	dataflowshapeFrom.GongCopyBasicFields(dataflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshapeFrom.DataFlow != nil {
		dataflowshapeTo.DataFlow = GongCopyBranchDataFlow(mapOrigCopy, dataflowshapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDataShape(mapOrigCopy map[any]any, datashapeFrom *DataShape) (datashapeTo *DataShape) {
	var alreadyCopied bool
	datashapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, datashapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	diagramprocessTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramprocessFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	externalparticipantshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, externalparticipantshapeFrom)
	if alreadyCopied {
		return
	}
	externalparticipantshapeFrom.GongCopyBasicFields(externalparticipantshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if externalparticipantshapeFrom.Participant != nil {
		externalparticipantshapeTo.Participant = GongCopyBranchParticipant(mapOrigCopy, externalparticipantshapeFrom.Participant)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {
	var alreadyCopied bool
	libraryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, libraryFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range noteFrom.Tasks {
		noteTo.Tasks = append(noteTo.Tasks, GongCopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func GongCopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {
	var alreadyCopied bool
	noteshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteshapeFrom)
	if alreadyCopied {
		return
	}
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteTaskShape(mapOrigCopy map[any]any, notetaskshapeFrom *NoteTaskShape) (notetaskshapeTo *NoteTaskShape) {
	var alreadyCopied bool
	notetaskshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notetaskshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	participantTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, participantFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	participantshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, participantshapeFrom)
	if alreadyCopied {
		return
	}
	participantshapeFrom.GongCopyBasicFields(participantshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if participantshapeFrom.Participant != nil {
		participantshapeTo.Participant = GongCopyBranchParticipant(mapOrigCopy, participantshapeFrom.Participant)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchProcess(mapOrigCopy map[any]any, processFrom *Process) (processTo *Process) {
	var alreadyCopied bool
	processTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, processFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	processshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, processshapeFrom)
	if alreadyCopied {
		return
	}
	processshapeFrom.GongCopyBasicFields(processshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if processshapeFrom.Process != nil {
		processshapeTo.Process = GongCopyBranchProcess(mapOrigCopy, processshapeFrom.Process)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {
	var alreadyCopied bool
	resourceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, resourceFrom)
	if alreadyCopied {
		return
	}
	resourceFrom.GongCopyBasicFields(resourceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTask(mapOrigCopy map[any]any, taskFrom *Task) (taskTo *Task) {
	var alreadyCopied bool
	taskTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskFrom)
	if alreadyCopied {
		return
	}
	taskFrom.GongCopyBasicFields(taskTo)

	//insertion point for the staging of instances referenced by pointers
	if taskFrom.Type != nil {
		taskTo.Type = GongCopyBranchProcess(mapOrigCopy, taskFrom.Type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskShape(mapOrigCopy map[any]any, taskshapeFrom *TaskShape) (taskshapeTo *TaskShape) {
	var alreadyCopied bool
	taskshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskshapeFrom)
	if alreadyCopied {
		return
	}
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

// insertion point for unstage branch per struct
func (allocatedprocessshape *AllocatedProcessShape) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(data) {
		return
	}

	data.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dataflow *DataFlow) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(resource) {
		return
	}

	resource.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (task *Task) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructPointer(&reference.Participant, stage.Participants_reference, instance.Participant)
	__gong__reconstructPointer(&reference.Process, stage.Processs_reference, instance.Process)
	// insertion point for slice of pointers field
}

func (reference *AllocatedResourceShape) GongReconstructPointersFromReferences(stage *Stage, instance *AllocatedResourceShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Participant, stage.Participants_reference, instance.Participant)
	__gong__reconstructPointer(&reference.Resource, stage.Resources_reference, instance.Resource)
	// insertion point for slice of pointers field
}

func (reference *ControlFlow) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlow) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Start, stage.Tasks_reference, instance.Start)
	__gong__reconstructPointer(&reference.End, stage.Tasks_reference, instance.End)
	// insertion point for slice of pointers field
}

func (reference *ControlFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlowShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ControlFlow, stage.ControlFlows_reference, instance.ControlFlow)
	// insertion point for slice of pointers field
}

func (reference *Data) GongReconstructPointersFromReferences(stage *Stage, instance *Data) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DataFlow) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlow) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.StartTask, stage.Tasks_reference, instance.StartTask)
	__gong__reconstructPointer(&reference.EndTask, stage.Tasks_reference, instance.EndTask)
	__gong__reconstructPointer(&reference.StartExternalParticipant, stage.Participants_reference, instance.StartExternalParticipant)
	__gong__reconstructPointer(&reference.EndExternalParticipant, stage.Participants_reference, instance.EndExternalParticipant)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Datas, stage.Datas_reference, instance.Datas)
}

func (reference *DataFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlowShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DataFlow, stage.DataFlows_reference, instance.DataFlow)
	// insertion point for slice of pointers field
}

func (reference *DataShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Data, stage.Datas_reference, instance.Data)
	__gong__reconstructPointer(&reference.DataFlow, stage.DataFlows_reference, instance.DataFlow)
	// insertion point for slice of pointers field
}

func (reference *DiagramProcess) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramProcess) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Process_Shapes, stage.ProcessShapes_reference, instance.Process_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ProcesssWhoseNodeIsExpanded, stage.Processs_reference, instance.ProcesssWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Participant_Shapes, stage.ParticipantShapes_reference, instance.Participant_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParticipantWhoseNodeIsExpanded, stage.Participants_reference, instance.ParticipantWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalParticipant_Shapes, stage.ExternalParticipantShapes_reference, instance.ExternalParticipant_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalParticipantWhoseNodeIsExpanded, stage.Participants_reference, instance.ExternalParticipantWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, stage.Participants_reference, instance.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, stage.Participants_reference, instance.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TasksWhoseNodeIsExpanded, stage.Tasks_reference, instance.TasksWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Task_Shapes, stage.TaskShapes_reference, instance.Task_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlFlowsWhoseNodeIsExpanded, stage.ControlFlows_reference, instance.ControlFlowsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlFlow_Shapes, stage.ControlFlowShapes_reference, instance.ControlFlow_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_reference, instance.DataFlowsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlow_Shapes, stage.DataFlowShapes_reference, instance.DataFlow_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DatasWhoseNodeIsExpanded, stage.Datas_reference, instance.DatasWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Data_Shapes, stage.DataShapes_reference, instance.Data_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlowsWhoseDataNodeIsExpanded, stage.DataFlows_reference, instance.DataFlowsWhoseDataNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AllocatedResourcesWhoseNodeIsExpanded, stage.Resources_reference, instance.AllocatedResourcesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AllocatedResourceShapes, stage.AllocatedResourceShapes_reference, instance.AllocatedResourceShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AllocatedProcessesWhoseNodeIsExpanded, stage.Processs_reference, instance.AllocatedProcessesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AllocatedProcessShapes, stage.AllocatedProcessShapes_reference, instance.AllocatedProcessShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note_Shapes, stage.NoteShapes_reference, instance.Note_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotesWhoseNodeIsExpanded, stage.Notes_reference, instance.NotesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteTaskShapes, stage.NoteTaskShapes_reference, instance.NoteTaskShapes)
}

func (reference *ExternalParticipantShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExternalParticipantShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Participant, stage.Participants_reference, instance.Participant)
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference, instance.SubLibrariesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootProcesses, stage.Processs_reference, instance.RootProcesses)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ProcesssWhoseNodeIsExpanded, stage.Processs_reference, instance.ProcesssWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootDataFlows, stage.DataFlows_reference, instance.RootDataFlows)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_reference, instance.DataFlowsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootDatas, stage.Datas_reference, instance.RootDatas)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DatasWhoseNodeIsExpanded, stage.Datas_reference, instance.DatasWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootResources, stage.Resources_reference, instance.RootResources)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ResourcesWhoseNodeIsExpanded, stage.Resources_reference, instance.ResourcesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParticipantsWhoseNodeIsExpanded, stage.Participants_reference, instance.ParticipantsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootNotes, stage.Notes_reference, instance.RootNotes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotesWhoseNodeIsExpanded, stage.Notes_reference, instance.NotesWhoseNodeIsExpanded)
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tasks, stage.Tasks_reference, instance.Tasks)
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	// insertion point for slice of pointers field
}

func (reference *NoteTaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteTaskShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	// insertion point for slice of pointers field
}

func (reference *Participant) GongReconstructPointersFromReferences(stage *Stage, instance *Participant) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Resources, stage.Resources_reference, instance.Resources)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Processes, stage.Processs_reference, instance.Processes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tasks, stage.Tasks_reference, instance.Tasks)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlFlows, stage.ControlFlows_reference, instance.ControlFlows)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskWhoseOutControlFlowsNodeIsExpanded, stage.Tasks_reference, instance.TaskWhoseOutControlFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskWhoseInControlFlowsNodeIsExpanded, stage.Tasks_reference, instance.TaskWhoseInControlFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskWhoseOutDataFlowsNodeIsExpanded, stage.Tasks_reference, instance.TaskWhoseOutDataFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskWhoseInDataFlowsNodeIsExpanded, stage.Tasks_reference, instance.TaskWhoseInDataFlowsNodeIsExpanded)
}

func (reference *ParticipantShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParticipantShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Participant, stage.Participants_reference, instance.Participant)
	// insertion point for slice of pointers field
}

func (reference *Process) GongReconstructPointersFromReferences(stage *Stage, instance *Process) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramProcesss, stage.DiagramProcesss_reference, instance.DiagramProcesss)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramProcessWhoseNodeIsExpanded, stage.DiagramProcesss_reference, instance.DiagramProcessWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubProcesses, stage.Processs_reference, instance.SubProcesses)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Participants, stage.Participants_reference, instance.Participants)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ParticipantWhoseNodeIsExpanded, stage.Participants_reference, instance.ParticipantWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlows, stage.DataFlows_reference, instance.DataFlows)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalParticipants, stage.Participants_reference, instance.ExternalParticipants)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalParticipantWhoseNodeIsExpanded, stage.Participants_reference, instance.ExternalParticipantWhoseNodeIsExpanded)
}

func (reference *ProcessShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProcessShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Process, stage.Processs_reference, instance.Process)
	// insertion point for slice of pointers field
}

func (reference *Resource) GongReconstructPointersFromReferences(stage *Stage, instance *Resource) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Task) GongReconstructPointersFromReferences(stage *Stage, instance *Task) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Type, stage.Processs_reference, instance.Type)
	// insertion point for slice of pointers field
}

func (reference *TaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AllocatedProcessShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Participant, stage.Participants_instance)
	__gong__reconstructPointerFromInstance(&reference.Process, stage.Processs_instance)
	// insertion point for slice of pointers fields
}

func (reference *AllocatedResourceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Participant, stage.Participants_instance)
	__gong__reconstructPointerFromInstance(&reference.Resource, stage.Resources_instance)
	// insertion point for slice of pointers fields
}

func (reference *ControlFlow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Start, stage.Tasks_instance)
	__gong__reconstructPointerFromInstance(&reference.End, stage.Tasks_instance)
	// insertion point for slice of pointers fields
}

func (reference *ControlFlowShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ControlFlow, stage.ControlFlows_instance)
	// insertion point for slice of pointers fields
}

func (reference *Data) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DataFlow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.StartTask, stage.Tasks_instance)
	__gong__reconstructPointerFromInstance(&reference.EndTask, stage.Tasks_instance)
	__gong__reconstructPointerFromInstance(&reference.StartExternalParticipant, stage.Participants_instance)
	__gong__reconstructPointerFromInstance(&reference.EndExternalParticipant, stage.Participants_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Datas, stage.Datas_instance)
}

func (reference *DataFlowShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DataFlow, stage.DataFlows_instance)
	// insertion point for slice of pointers fields
}

func (reference *DataShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Data, stage.Datas_instance)
	__gong__reconstructPointerFromInstance(&reference.DataFlow, stage.DataFlows_instance)
	// insertion point for slice of pointers fields
}

func (reference *DiagramProcess) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Process_Shapes, stage.ProcessShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ProcesssWhoseNodeIsExpanded, stage.Processs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Participant_Shapes, stage.ParticipantShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParticipantWhoseNodeIsExpanded, stage.Participants_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalParticipant_Shapes, stage.ExternalParticipantShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalParticipantWhoseNodeIsExpanded, stage.Participants_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, stage.Participants_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, stage.Participants_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TasksWhoseNodeIsExpanded, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Task_Shapes, stage.TaskShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlFlowsWhoseNodeIsExpanded, stage.ControlFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlFlow_Shapes, stage.ControlFlowShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlow_Shapes, stage.DataFlowShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DatasWhoseNodeIsExpanded, stage.Datas_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Data_Shapes, stage.DataShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlowsWhoseDataNodeIsExpanded, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AllocatedResourcesWhoseNodeIsExpanded, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AllocatedResourceShapes, stage.AllocatedResourceShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AllocatedProcessesWhoseNodeIsExpanded, stage.Processs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AllocatedProcessShapes, stage.AllocatedProcessShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note_Shapes, stage.NoteShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotesWhoseNodeIsExpanded, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteTaskShapes, stage.NoteTaskShapes_instance)
}

func (reference *ExternalParticipantShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Participant, stage.Participants_instance)
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootProcesses, stage.Processs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ProcesssWhoseNodeIsExpanded, stage.Processs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootDataFlows, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootDatas, stage.Datas_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DatasWhoseNodeIsExpanded, stage.Datas_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootResources, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ResourcesWhoseNodeIsExpanded, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParticipantsWhoseNodeIsExpanded, stage.Participants_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootNotes, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotesWhoseNodeIsExpanded, stage.Notes_instance)
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tasks, stage.Tasks_instance)
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteTaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
	// insertion point for slice of pointers fields
}

func (reference *Participant) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Resources, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Processes, stage.Processs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tasks, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlFlows, stage.ControlFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskWhoseOutControlFlowsNodeIsExpanded, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskWhoseInControlFlowsNodeIsExpanded, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskWhoseOutDataFlowsNodeIsExpanded, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskWhoseInDataFlowsNodeIsExpanded, stage.Tasks_instance)
}

func (reference *ParticipantShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Participant, stage.Participants_instance)
	// insertion point for slice of pointers fields
}

func (reference *Process) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramProcesss, stage.DiagramProcesss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramProcessWhoseNodeIsExpanded, stage.DiagramProcesss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubProcesses, stage.Processs_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Participants, stage.Participants_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ParticipantWhoseNodeIsExpanded, stage.Participants_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlows, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalParticipants, stage.Participants_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalParticipantWhoseNodeIsExpanded, stage.Participants_instance)
}

func (reference *ProcessShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Process, stage.Processs_instance)
	// insertion point for slice of pointers fields
}

func (reference *Resource) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Task) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Type, stage.Processs_instance)
	// insertion point for slice of pointers fields
}

func (reference *TaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
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
	if allocatedprocessshape.Participant != allocatedprocessshapeOther.Participant {
		diffs = append(diffs, allocatedprocessshape.GongMarshallField(stage, "Participant"))
	}
	if allocatedprocessshape.Process != allocatedprocessshapeOther.Process {
		diffs = append(diffs, allocatedprocessshape.GongMarshallField(stage, "Process"))
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
	if allocatedresourceshape.Participant != allocatedresourceshapeOther.Participant {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Participant"))
	}
	if allocatedresourceshape.Resource != allocatedresourceshapeOther.Resource {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Resource"))
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
	if controlflow.Start != controlflowOther.Start {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Start"))
	}
	if controlflow.End != controlflowOther.End {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "End"))
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
	if controlflowshape.ControlFlow != controlflowshapeOther.ControlFlow {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "ControlFlow"))
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
	if ops := __gong__diffSliceOfPointers(stage, dataflow, "Datas", dataflowOther.Datas, dataflow.Datas); ops != "" {
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
	if dataflow.StartTask != dataflowOther.StartTask {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "StartTask"))
	}
	if dataflow.EndTask != dataflowOther.EndTask {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "EndTask"))
	}
	if dataflow.StartExternalParticipant != dataflowOther.StartExternalParticipant {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "StartExternalParticipant"))
	}
	if dataflow.EndExternalParticipant != dataflowOther.EndExternalParticipant {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "EndExternalParticipant"))
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
	if dataflowshape.DataFlow != dataflowshapeOther.DataFlow {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "DataFlow"))
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
	if datashape.Data != datashapeOther.Data {
		diffs = append(diffs, datashape.GongMarshallField(stage, "Data"))
	}
	if datashape.DataFlow != datashapeOther.DataFlow {
		diffs = append(diffs, datashape.GongMarshallField(stage, "DataFlow"))
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
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "Process_Shapes", diagramprocessOther.Process_Shapes, diagramprocess.Process_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramprocess.IsProcesssNodeExpanded != diagramprocessOther.IsProcesssNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsProcesssNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "ProcesssWhoseNodeIsExpanded", diagramprocessOther.ProcesssWhoseNodeIsExpanded, diagramprocess.ProcesssWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "Participant_Shapes", diagramprocessOther.Participant_Shapes, diagramprocess.Participant_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramprocess.IsParticipantsNodeExpanded != diagramprocessOther.IsParticipantsNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsParticipantsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "ParticipantWhoseNodeIsExpanded", diagramprocessOther.ParticipantWhoseNodeIsExpanded, diagramprocess.ParticipantWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "ExternalParticipant_Shapes", diagramprocessOther.ExternalParticipant_Shapes, diagramprocess.ExternalParticipant_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramprocess.IsExternalParticipantsNodeExpanded != diagramprocessOther.IsExternalParticipantsNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsExternalParticipantsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "ExternalParticipantWhoseNodeIsExpanded", diagramprocessOther.ExternalParticipantWhoseNodeIsExpanded, diagramprocess.ExternalParticipantWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", diagramprocessOther.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", diagramprocessOther.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "TasksWhoseNodeIsExpanded", diagramprocessOther.TasksWhoseNodeIsExpanded, diagramprocess.TasksWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "Task_Shapes", diagramprocessOther.Task_Shapes, diagramprocess.Task_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "ControlFlowsWhoseNodeIsExpanded", diagramprocessOther.ControlFlowsWhoseNodeIsExpanded, diagramprocess.ControlFlowsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "ControlFlow_Shapes", diagramprocessOther.ControlFlow_Shapes, diagramprocess.ControlFlow_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "DataFlowsWhoseNodeIsExpanded", diagramprocessOther.DataFlowsWhoseNodeIsExpanded, diagramprocess.DataFlowsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "DataFlow_Shapes", diagramprocessOther.DataFlow_Shapes, diagramprocess.DataFlow_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "DatasWhoseNodeIsExpanded", diagramprocessOther.DatasWhoseNodeIsExpanded, diagramprocess.DatasWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "Data_Shapes", diagramprocessOther.Data_Shapes, diagramprocess.Data_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "DataFlowsWhoseDataNodeIsExpanded", diagramprocessOther.DataFlowsWhoseDataNodeIsExpanded, diagramprocess.DataFlowsWhoseDataNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "AllocatedResourcesWhoseNodeIsExpanded", diagramprocessOther.AllocatedResourcesWhoseNodeIsExpanded, diagramprocess.AllocatedResourcesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "AllocatedResourceShapes", diagramprocessOther.AllocatedResourceShapes, diagramprocess.AllocatedResourceShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "AllocatedProcessesWhoseNodeIsExpanded", diagramprocessOther.AllocatedProcessesWhoseNodeIsExpanded, diagramprocess.AllocatedProcessesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "AllocatedProcessShapes", diagramprocessOther.AllocatedProcessShapes, diagramprocess.AllocatedProcessShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "Note_Shapes", diagramprocessOther.Note_Shapes, diagramprocess.Note_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "NotesWhoseNodeIsExpanded", diagramprocessOther.NotesWhoseNodeIsExpanded, diagramprocess.NotesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramprocess.IsNotesNodeExpanded != diagramprocessOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagramprocess.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramprocess, "NoteTaskShapes", diagramprocessOther.NoteTaskShapes, diagramprocess.NoteTaskShapes); ops != "" {
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
	if externalparticipantshape.Participant != externalparticipantshapeOther.Participant {
		diffs = append(diffs, externalparticipantshape.GongMarshallField(stage, "Participant"))
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
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootProcesses", libraryOther.RootProcesses, library.RootProcesses); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsProcessesNodeExpanded != libraryOther.IsProcessesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsProcessesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "ProcesssWhoseNodeIsExpanded", libraryOther.ProcesssWhoseNodeIsExpanded, library.ProcesssWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootDataFlows", libraryOther.RootDataFlows, library.RootDataFlows); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsDataFlowsNodeExpanded != libraryOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "DataFlowsWhoseNodeIsExpanded", libraryOther.DataFlowsWhoseNodeIsExpanded, library.DataFlowsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootDatas", libraryOther.RootDatas, library.RootDatas); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsDatasNodeExpanded != libraryOther.IsDatasNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsDatasNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "DatasWhoseNodeIsExpanded", libraryOther.DatasWhoseNodeIsExpanded, library.DatasWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootResources", libraryOther.RootResources, library.RootResources); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsResourcesNodeExpanded != libraryOther.IsResourcesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "ResourcesWhoseNodeIsExpanded", libraryOther.ResourcesWhoseNodeIsExpanded, library.ResourcesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "ParticipantsWhoseNodeIsExpanded", libraryOther.ParticipantsWhoseNodeIsExpanded, library.ParticipantsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootNotes", libraryOther.RootNotes, library.RootNotes); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsNotesNodeExpanded != libraryOther.IsNotesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "NotesWhoseNodeIsExpanded", libraryOther.NotesWhoseNodeIsExpanded, library.NotesWhoseNodeIsExpanded); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, note, "Tasks", noteOther.Tasks, note.Tasks); ops != "" {
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
	if noteshape.Note != noteshapeOther.Note {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
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
	if notetaskshape.Note != notetaskshapeOther.Note {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
	}
	if notetaskshape.Task != notetaskshapeOther.Task {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
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
	if ops := __gong__diffSliceOfPointers(stage, participant, "Resources", participantOther.Resources, participant.Resources); ops != "" {
		diffs = append(diffs, ops)
	}
	if participant.IsResourcesNodeExpanded != participantOther.IsResourcesNodeExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, participant, "Processes", participantOther.Processes, participant.Processes); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, participant, "Tasks", participantOther.Tasks, participant.Tasks); ops != "" {
		diffs = append(diffs, ops)
	}
	if participant.IsControlFlowsNodeExpanded != participantOther.IsControlFlowsNodeExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsControlFlowsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, participant, "ControlFlows", participantOther.ControlFlows, participant.ControlFlows); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, participant, "TaskWhoseOutControlFlowsNodeIsExpanded", participantOther.TaskWhoseOutControlFlowsNodeIsExpanded, participant.TaskWhoseOutControlFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, participant, "TaskWhoseInControlFlowsNodeIsExpanded", participantOther.TaskWhoseInControlFlowsNodeIsExpanded, participant.TaskWhoseInControlFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if participant.IsDataFlowsNodeExpanded != participantOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, participant.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, participant, "TaskWhoseOutDataFlowsNodeIsExpanded", participantOther.TaskWhoseOutDataFlowsNodeIsExpanded, participant.TaskWhoseOutDataFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, participant, "TaskWhoseInDataFlowsNodeIsExpanded", participantOther.TaskWhoseInDataFlowsNodeIsExpanded, participant.TaskWhoseInDataFlowsNodeIsExpanded); ops != "" {
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
	if participantshape.Participant != participantshapeOther.Participant {
		diffs = append(diffs, participantshape.GongMarshallField(stage, "Participant"))
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
	if ops := __gong__diffSliceOfPointers(stage, process, "DiagramProcesss", processOther.DiagramProcesss, process.DiagramProcesss); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, process, "DiagramProcessWhoseNodeIsExpanded", processOther.DiagramProcessWhoseNodeIsExpanded, process.DiagramProcessWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if process.IsSubProcessNodeExpanded != processOther.IsSubProcessNodeExpanded {
		diffs = append(diffs, process.GongMarshallField(stage, "IsSubProcessNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, process, "SubProcesses", processOther.SubProcesses, process.SubProcesses); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, process, "Participants", processOther.Participants, process.Participants); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, process, "ParticipantWhoseNodeIsExpanded", processOther.ParticipantWhoseNodeIsExpanded, process.ParticipantWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, process, "DataFlows", processOther.DataFlows, process.DataFlows); ops != "" {
		diffs = append(diffs, ops)
	}
	if process.IsDataFlowsNodeExpanded != processOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, process.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, process, "ExternalParticipants", processOther.ExternalParticipants, process.ExternalParticipants); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, process, "ExternalParticipantWhoseNodeIsExpanded", processOther.ExternalParticipantWhoseNodeIsExpanded, process.ExternalParticipantWhoseNodeIsExpanded); ops != "" {
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
	if processshape.Process != processshapeOther.Process {
		diffs = append(diffs, processshape.GongMarshallField(stage, "Process"))
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
	if task.Type != taskOther.Type {
		diffs = append(diffs, task.GongMarshallField(stage, "Type"))
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
	if taskshape.Task != taskshapeOther.Task {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Task"))
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

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
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
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
