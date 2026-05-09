// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *ControlFlow:
		ok = stage.IsStagedControlFlow(target)

	case *ControlFlowShape:
		ok = stage.IsStagedControlFlowShape(target)

	case *Data:
		ok = stage.IsStagedData(target)

	case *DataFlow:
		ok = stage.IsStagedDataFlow(target)

	case *DataFlowShape:
		ok = stage.IsStagedDataFlowShape(target)

	case *DataShape:
		ok = stage.IsStagedDataShape(target)

	case *DiagramProcess:
		ok = stage.IsStagedDiagramProcess(target)

	case *ExternalParticipantShape:
		ok = stage.IsStagedExternalParticipantShape(target)

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

	case *Resource:
		ok = stage.IsStagedResource(target)

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
	case *ControlFlow:
		ok = stage.IsStagedControlFlow(target)

	case *ControlFlowShape:
		ok = stage.IsStagedControlFlowShape(target)

	case *Data:
		ok = stage.IsStagedData(target)

	case *DataFlow:
		ok = stage.IsStagedDataFlow(target)

	case *DataFlowShape:
		ok = stage.IsStagedDataFlowShape(target)

	case *DataShape:
		ok = stage.IsStagedDataShape(target)

	case *DiagramProcess:
		ok = stage.IsStagedDiagramProcess(target)

	case *ExternalParticipantShape:
		ok = stage.IsStagedExternalParticipantShape(target)

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

	case *Resource:
		ok = stage.IsStagedResource(target)

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
func (stage *Stage) IsStagedControlFlow(controlflow *ControlFlow) (ok bool) {

	_, ok = stage.ControlFlows[controlflow]

	return
}

func (stage *Stage) IsStagedControlFlowShape(controlflowshape *ControlFlowShape) (ok bool) {

	_, ok = stage.ControlFlowShapes[controlflowshape]

	return
}

func (stage *Stage) IsStagedData(data *Data) (ok bool) {

	_, ok = stage.Datas[data]

	return
}

func (stage *Stage) IsStagedDataFlow(dataflow *DataFlow) (ok bool) {

	_, ok = stage.DataFlows[dataflow]

	return
}

func (stage *Stage) IsStagedDataFlowShape(dataflowshape *DataFlowShape) (ok bool) {

	_, ok = stage.DataFlowShapes[dataflowshape]

	return
}

func (stage *Stage) IsStagedDataShape(datashape *DataShape) (ok bool) {

	_, ok = stage.DataShapes[datashape]

	return
}

func (stage *Stage) IsStagedDiagramProcess(diagramprocess *DiagramProcess) (ok bool) {

	_, ok = stage.DiagramProcesss[diagramprocess]

	return
}

func (stage *Stage) IsStagedExternalParticipantShape(externalparticipantshape *ExternalParticipantShape) (ok bool) {

	_, ok = stage.ExternalParticipantShapes[externalparticipantshape]

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

func (stage *Stage) IsStagedResource(resource *Resource) (ok bool) {

	_, ok = stage.Resources[resource]

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
	case *ControlFlow:
		stage.StageBranchControlFlow(target)

	case *ControlFlowShape:
		stage.StageBranchControlFlowShape(target)

	case *Data:
		stage.StageBranchData(target)

	case *DataFlow:
		stage.StageBranchDataFlow(target)

	case *DataFlowShape:
		stage.StageBranchDataFlowShape(target)

	case *DataShape:
		stage.StageBranchDataShape(target)

	case *DiagramProcess:
		stage.StageBranchDiagramProcess(target)

	case *ExternalParticipantShape:
		stage.StageBranchExternalParticipantShape(target)

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

	case *Resource:
		stage.StageBranchResource(target)

	case *Task:
		stage.StageBranchTask(target)

	case *TaskShape:
		stage.StageBranchTaskShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchControlFlow(controlflow *ControlFlow) {

	// check if instance is already staged
	if IsStaged(stage, controlflow) {
		return
	}

	controlflow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflow.Start != nil {
		StageBranch(stage, controlflow.Start)
	}
	if controlflow.End != nil {
		StageBranch(stage, controlflow.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchControlFlowShape(controlflowshape *ControlFlowShape) {

	// check if instance is already staged
	if IsStaged(stage, controlflowshape) {
		return
	}

	controlflowshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshape.ControlFlow != nil {
		StageBranch(stage, controlflowshape.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchData(data *Data) {

	// check if instance is already staged
	if IsStaged(stage, data) {
		return
	}

	data.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDataFlow(dataflow *DataFlow) {

	// check if instance is already staged
	if IsStaged(stage, dataflow) {
		return
	}

	dataflow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflow.StartTask != nil {
		StageBranch(stage, dataflow.StartTask)
	}
	if dataflow.EndTask != nil {
		StageBranch(stage, dataflow.EndTask)
	}
	if dataflow.StartExternalParticipant != nil {
		StageBranch(stage, dataflow.StartExternalParticipant)
	}
	if dataflow.EndExternalParticipant != nil {
		StageBranch(stage, dataflow.EndExternalParticipant)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflow.Datas {
		StageBranch(stage, _data)
	}

}

func (stage *Stage) StageBranchDataFlowShape(dataflowshape *DataFlowShape) {

	// check if instance is already staged
	if IsStaged(stage, dataflowshape) {
		return
	}

	dataflowshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshape.DataFlow != nil {
		StageBranch(stage, dataflowshape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDataShape(datashape *DataShape) {

	// check if instance is already staged
	if IsStaged(stage, datashape) {
		return
	}

	datashape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datashape.Data != nil {
		StageBranch(stage, datashape.Data)
	}
	if datashape.DataFlow != nil {
		StageBranch(stage, datashape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

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
	for _, _externalparticipantshape := range diagramprocess.ExternalParticipant_Shapes {
		StageBranch(stage, _externalparticipantshape)
	}
	for _, _participant := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
		StageBranch(stage, _participant)
	}
	for _, _participant := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
		StageBranch(stage, _participant)
	}
	for _, _participant := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
		StageBranch(stage, _participant)
	}
	for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _taskshape := range diagramprocess.Task_Shapes {
		StageBranch(stage, _taskshape)
	}
	for _, _controlflow := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
		StageBranch(stage, _controlflow)
	}
	for _, _controlflowshape := range diagramprocess.ControlFlow_Shapes {
		StageBranch(stage, _controlflowshape)
	}
	for _, _dataflow := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
		StageBranch(stage, _dataflow)
	}
	for _, _dataflowshape := range diagramprocess.DataFlow_Shapes {
		StageBranch(stage, _dataflowshape)
	}
	for _, _data := range diagramprocess.DatasWhoseNodeIsExpanded {
		StageBranch(stage, _data)
	}
	for _, _datashape := range diagramprocess.Data_Shapes {
		StageBranch(stage, _datashape)
	}
	for _, _dataflow := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
		StageBranch(stage, _dataflow)
	}

}

func (stage *Stage) StageBranchExternalParticipantShape(externalparticipantshape *ExternalParticipantShape) {

	// check if instance is already staged
	if IsStaged(stage, externalparticipantshape) {
		return
	}

	externalparticipantshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if externalparticipantshape.Participant != nil {
		StageBranch(stage, externalparticipantshape.Participant)
	}

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
	for _, _process := range library.RootProcesses {
		StageBranch(stage, _process)
	}
	for _, _process := range library.ProcesssWhoseNodeIsExpanded {
		StageBranch(stage, _process)
	}
	for _, _dataflow := range library.RootDataFlows {
		StageBranch(stage, _dataflow)
	}
	for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
		StageBranch(stage, _dataflow)
	}
	for _, _data := range library.RootDatas {
		StageBranch(stage, _data)
	}
	for _, _data := range library.DatasWhoseNodeIsExpanded {
		StageBranch(stage, _data)
	}
	for _, _resource := range library.RootResources {
		StageBranch(stage, _resource)
	}
	for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
		StageBranch(stage, _resource)
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
	for _, _controlflow := range participant.ControlFlows {
		StageBranch(stage, _controlflow)
	}
	for _, _task := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
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
	for _, _dataflow := range process.DataFlows {
		StageBranch(stage, _dataflow)
	}
	for _, _participant := range process.ExternalParticipants {
		StageBranch(stage, _participant)
	}
	for _, _participant := range process.ExternalParticipantWhoseNodeIsExpanded {
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

func (stage *Stage) StageBranchResource(resource *Resource) {

	// check if instance is already staged
	if IsStaged(stage, resource) {
		return
	}

	resource.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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
	case *ControlFlow:
		toT := CopyBranchControlFlow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlFlowShape:
		toT := CopyBranchControlFlowShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Data:
		toT := CopyBranchData(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataFlow:
		toT := CopyBranchDataFlow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataFlowShape:
		toT := CopyBranchDataFlowShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataShape:
		toT := CopyBranchDataShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramProcess:
		toT := CopyBranchDiagramProcess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExternalParticipantShape:
		toT := CopyBranchExternalParticipantShape(mapOrigCopy, fromT)
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

	case *Resource:
		toT := CopyBranchResource(mapOrigCopy, fromT)
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
func CopyBranchControlFlow(mapOrigCopy map[any]any, controlflowFrom *ControlFlow) (controlflowTo *ControlFlow) {

	// controlflowFrom has already been copied
	if _controlflowTo, ok := mapOrigCopy[controlflowFrom]; ok {
		controlflowTo = _controlflowTo.(*ControlFlow)
		return
	}

	controlflowTo = new(ControlFlow)
	mapOrigCopy[controlflowFrom] = controlflowTo
	controlflowFrom.CopyBasicFields(controlflowTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowFrom.Start != nil {
		controlflowTo.Start = CopyBranchTask(mapOrigCopy, controlflowFrom.Start)
	}
	if controlflowFrom.End != nil {
		controlflowTo.End = CopyBranchTask(mapOrigCopy, controlflowFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchControlFlowShape(mapOrigCopy map[any]any, controlflowshapeFrom *ControlFlowShape) (controlflowshapeTo *ControlFlowShape) {

	// controlflowshapeFrom has already been copied
	if _controlflowshapeTo, ok := mapOrigCopy[controlflowshapeFrom]; ok {
		controlflowshapeTo = _controlflowshapeTo.(*ControlFlowShape)
		return
	}

	controlflowshapeTo = new(ControlFlowShape)
	mapOrigCopy[controlflowshapeFrom] = controlflowshapeTo
	controlflowshapeFrom.CopyBasicFields(controlflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshapeFrom.ControlFlow != nil {
		controlflowshapeTo.ControlFlow = CopyBranchControlFlow(mapOrigCopy, controlflowshapeFrom.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchData(mapOrigCopy map[any]any, dataFrom *Data) (dataTo *Data) {

	// dataFrom has already been copied
	if _dataTo, ok := mapOrigCopy[dataFrom]; ok {
		dataTo = _dataTo.(*Data)
		return
	}

	dataTo = new(Data)
	mapOrigCopy[dataFrom] = dataTo
	dataFrom.CopyBasicFields(dataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDataFlow(mapOrigCopy map[any]any, dataflowFrom *DataFlow) (dataflowTo *DataFlow) {

	// dataflowFrom has already been copied
	if _dataflowTo, ok := mapOrigCopy[dataflowFrom]; ok {
		dataflowTo = _dataflowTo.(*DataFlow)
		return
	}

	dataflowTo = new(DataFlow)
	mapOrigCopy[dataflowFrom] = dataflowTo
	dataflowFrom.CopyBasicFields(dataflowTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowFrom.StartTask != nil {
		dataflowTo.StartTask = CopyBranchTask(mapOrigCopy, dataflowFrom.StartTask)
	}
	if dataflowFrom.EndTask != nil {
		dataflowTo.EndTask = CopyBranchTask(mapOrigCopy, dataflowFrom.EndTask)
	}
	if dataflowFrom.StartExternalParticipant != nil {
		dataflowTo.StartExternalParticipant = CopyBranchParticipant(mapOrigCopy, dataflowFrom.StartExternalParticipant)
	}
	if dataflowFrom.EndExternalParticipant != nil {
		dataflowTo.EndExternalParticipant = CopyBranchParticipant(mapOrigCopy, dataflowFrom.EndExternalParticipant)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflowFrom.Datas {
		dataflowTo.Datas = append(dataflowTo.Datas, CopyBranchData(mapOrigCopy, _data))
	}

	return
}

func CopyBranchDataFlowShape(mapOrigCopy map[any]any, dataflowshapeFrom *DataFlowShape) (dataflowshapeTo *DataFlowShape) {

	// dataflowshapeFrom has already been copied
	if _dataflowshapeTo, ok := mapOrigCopy[dataflowshapeFrom]; ok {
		dataflowshapeTo = _dataflowshapeTo.(*DataFlowShape)
		return
	}

	dataflowshapeTo = new(DataFlowShape)
	mapOrigCopy[dataflowshapeFrom] = dataflowshapeTo
	dataflowshapeFrom.CopyBasicFields(dataflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshapeFrom.DataFlow != nil {
		dataflowshapeTo.DataFlow = CopyBranchDataFlow(mapOrigCopy, dataflowshapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDataShape(mapOrigCopy map[any]any, datashapeFrom *DataShape) (datashapeTo *DataShape) {

	// datashapeFrom has already been copied
	if _datashapeTo, ok := mapOrigCopy[datashapeFrom]; ok {
		datashapeTo = _datashapeTo.(*DataShape)
		return
	}

	datashapeTo = new(DataShape)
	mapOrigCopy[datashapeFrom] = datashapeTo
	datashapeFrom.CopyBasicFields(datashapeTo)

	//insertion point for the staging of instances referenced by pointers
	if datashapeFrom.Data != nil {
		datashapeTo.Data = CopyBranchData(mapOrigCopy, datashapeFrom.Data)
	}
	if datashapeFrom.DataFlow != nil {
		datashapeTo.DataFlow = CopyBranchDataFlow(mapOrigCopy, datashapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

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
	for _, _externalparticipantshape := range diagramprocessFrom.ExternalParticipant_Shapes {
		diagramprocessTo.ExternalParticipant_Shapes = append(diagramprocessTo.ExternalParticipant_Shapes, CopyBranchExternalParticipantShape(mapOrigCopy, _externalparticipantshape))
	}
	for _, _participant := range diagramprocessFrom.ExternalParticipantWhoseNodeIsExpanded {
		diagramprocessTo.ExternalParticipantWhoseNodeIsExpanded = append(diagramprocessTo.ExternalParticipantWhoseNodeIsExpanded, CopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _participant := range diagramprocessFrom.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
		diagramprocessTo.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded = append(diagramprocessTo.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, CopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _participant := range diagramprocessFrom.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
		diagramprocessTo.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded = append(diagramprocessTo.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, CopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _task := range diagramprocessFrom.TasksWhoseNodeIsExpanded {
		diagramprocessTo.TasksWhoseNodeIsExpanded = append(diagramprocessTo.TasksWhoseNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskshape := range diagramprocessFrom.Task_Shapes {
		diagramprocessTo.Task_Shapes = append(diagramprocessTo.Task_Shapes, CopyBranchTaskShape(mapOrigCopy, _taskshape))
	}
	for _, _controlflow := range diagramprocessFrom.ControlFlowsWhoseNodeIsExpanded {
		diagramprocessTo.ControlFlowsWhoseNodeIsExpanded = append(diagramprocessTo.ControlFlowsWhoseNodeIsExpanded, CopyBranchControlFlow(mapOrigCopy, _controlflow))
	}
	for _, _controlflowshape := range diagramprocessFrom.ControlFlow_Shapes {
		diagramprocessTo.ControlFlow_Shapes = append(diagramprocessTo.ControlFlow_Shapes, CopyBranchControlFlowShape(mapOrigCopy, _controlflowshape))
	}
	for _, _dataflow := range diagramprocessFrom.DataFlowsWhoseNodeIsExpanded {
		diagramprocessTo.DataFlowsWhoseNodeIsExpanded = append(diagramprocessTo.DataFlowsWhoseNodeIsExpanded, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _dataflowshape := range diagramprocessFrom.DataFlow_Shapes {
		diagramprocessTo.DataFlow_Shapes = append(diagramprocessTo.DataFlow_Shapes, CopyBranchDataFlowShape(mapOrigCopy, _dataflowshape))
	}
	for _, _data := range diagramprocessFrom.DatasWhoseNodeIsExpanded {
		diagramprocessTo.DatasWhoseNodeIsExpanded = append(diagramprocessTo.DatasWhoseNodeIsExpanded, CopyBranchData(mapOrigCopy, _data))
	}
	for _, _datashape := range diagramprocessFrom.Data_Shapes {
		diagramprocessTo.Data_Shapes = append(diagramprocessTo.Data_Shapes, CopyBranchDataShape(mapOrigCopy, _datashape))
	}
	for _, _dataflow := range diagramprocessFrom.DataFlowsWhoseDataNodeIsExpanded {
		diagramprocessTo.DataFlowsWhoseDataNodeIsExpanded = append(diagramprocessTo.DataFlowsWhoseDataNodeIsExpanded, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}

	return
}

func CopyBranchExternalParticipantShape(mapOrigCopy map[any]any, externalparticipantshapeFrom *ExternalParticipantShape) (externalparticipantshapeTo *ExternalParticipantShape) {

	// externalparticipantshapeFrom has already been copied
	if _externalparticipantshapeTo, ok := mapOrigCopy[externalparticipantshapeFrom]; ok {
		externalparticipantshapeTo = _externalparticipantshapeTo.(*ExternalParticipantShape)
		return
	}

	externalparticipantshapeTo = new(ExternalParticipantShape)
	mapOrigCopy[externalparticipantshapeFrom] = externalparticipantshapeTo
	externalparticipantshapeFrom.CopyBasicFields(externalparticipantshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if externalparticipantshapeFrom.Participant != nil {
		externalparticipantshapeTo.Participant = CopyBranchParticipant(mapOrigCopy, externalparticipantshapeFrom.Participant)
	}

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
	for _, _process := range libraryFrom.RootProcesses {
		libraryTo.RootProcesses = append(libraryTo.RootProcesses, CopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _process := range libraryFrom.ProcesssWhoseNodeIsExpanded {
		libraryTo.ProcesssWhoseNodeIsExpanded = append(libraryTo.ProcesssWhoseNodeIsExpanded, CopyBranchProcess(mapOrigCopy, _process))
	}
	for _, _dataflow := range libraryFrom.RootDataFlows {
		libraryTo.RootDataFlows = append(libraryTo.RootDataFlows, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _dataflow := range libraryFrom.DataFlowsWhoseNodeIsExpanded {
		libraryTo.DataFlowsWhoseNodeIsExpanded = append(libraryTo.DataFlowsWhoseNodeIsExpanded, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _data := range libraryFrom.RootDatas {
		libraryTo.RootDatas = append(libraryTo.RootDatas, CopyBranchData(mapOrigCopy, _data))
	}
	for _, _data := range libraryFrom.DatasWhoseNodeIsExpanded {
		libraryTo.DatasWhoseNodeIsExpanded = append(libraryTo.DatasWhoseNodeIsExpanded, CopyBranchData(mapOrigCopy, _data))
	}
	for _, _resource := range libraryFrom.RootResources {
		libraryTo.RootResources = append(libraryTo.RootResources, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _resource := range libraryFrom.ResourcesWhoseNodeIsExpanded {
		libraryTo.ResourcesWhoseNodeIsExpanded = append(libraryTo.ResourcesWhoseNodeIsExpanded, CopyBranchResource(mapOrigCopy, _resource))
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
	for _, _controlflow := range participantFrom.ControlFlows {
		participantTo.ControlFlows = append(participantTo.ControlFlows, CopyBranchControlFlow(mapOrigCopy, _controlflow))
	}
	for _, _task := range participantFrom.TaskWhoseOutControlFlowsNodeIsExpanded {
		participantTo.TaskWhoseOutControlFlowsNodeIsExpanded = append(participantTo.TaskWhoseOutControlFlowsNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range participantFrom.TaskWhoseInControlFlowsNodeIsExpanded {
		participantTo.TaskWhoseInControlFlowsNodeIsExpanded = append(participantTo.TaskWhoseInControlFlowsNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range participantFrom.TaskWhoseOutDataFlowsNodeIsExpanded {
		participantTo.TaskWhoseOutDataFlowsNodeIsExpanded = append(participantTo.TaskWhoseOutDataFlowsNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range participantFrom.TaskWhoseInDataFlowsNodeIsExpanded {
		participantTo.TaskWhoseInDataFlowsNodeIsExpanded = append(participantTo.TaskWhoseInDataFlowsNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
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
	for _, _dataflow := range processFrom.DataFlows {
		processTo.DataFlows = append(processTo.DataFlows, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _participant := range processFrom.ExternalParticipants {
		processTo.ExternalParticipants = append(processTo.ExternalParticipants, CopyBranchParticipant(mapOrigCopy, _participant))
	}
	for _, _participant := range processFrom.ExternalParticipantWhoseNodeIsExpanded {
		processTo.ExternalParticipantWhoseNodeIsExpanded = append(processTo.ExternalParticipantWhoseNodeIsExpanded, CopyBranchParticipant(mapOrigCopy, _participant))
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

func CopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {

	// resourceFrom has already been copied
	if _resourceTo, ok := mapOrigCopy[resourceFrom]; ok {
		resourceTo = _resourceTo.(*Resource)
		return
	}

	resourceTo = new(Resource)
	mapOrigCopy[resourceFrom] = resourceTo
	resourceFrom.CopyBasicFields(resourceTo)

	//insertion point for the staging of instances referenced by pointers

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
	case *ControlFlow:
		stage.UnstageBranchControlFlow(target)

	case *ControlFlowShape:
		stage.UnstageBranchControlFlowShape(target)

	case *Data:
		stage.UnstageBranchData(target)

	case *DataFlow:
		stage.UnstageBranchDataFlow(target)

	case *DataFlowShape:
		stage.UnstageBranchDataFlowShape(target)

	case *DataShape:
		stage.UnstageBranchDataShape(target)

	case *DiagramProcess:
		stage.UnstageBranchDiagramProcess(target)

	case *ExternalParticipantShape:
		stage.UnstageBranchExternalParticipantShape(target)

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

	case *Resource:
		stage.UnstageBranchResource(target)

	case *Task:
		stage.UnstageBranchTask(target)

	case *TaskShape:
		stage.UnstageBranchTaskShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchControlFlow(controlflow *ControlFlow) {

	// check if instance is already staged
	if !IsStaged(stage, controlflow) {
		return
	}

	controlflow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflow.Start != nil {
		UnstageBranch(stage, controlflow.Start)
	}
	if controlflow.End != nil {
		UnstageBranch(stage, controlflow.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchControlFlowShape(controlflowshape *ControlFlowShape) {

	// check if instance is already staged
	if !IsStaged(stage, controlflowshape) {
		return
	}

	controlflowshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshape.ControlFlow != nil {
		UnstageBranch(stage, controlflowshape.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchData(data *Data) {

	// check if instance is already staged
	if !IsStaged(stage, data) {
		return
	}

	data.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDataFlow(dataflow *DataFlow) {

	// check if instance is already staged
	if !IsStaged(stage, dataflow) {
		return
	}

	dataflow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflow.StartTask != nil {
		UnstageBranch(stage, dataflow.StartTask)
	}
	if dataflow.EndTask != nil {
		UnstageBranch(stage, dataflow.EndTask)
	}
	if dataflow.StartExternalParticipant != nil {
		UnstageBranch(stage, dataflow.StartExternalParticipant)
	}
	if dataflow.EndExternalParticipant != nil {
		UnstageBranch(stage, dataflow.EndExternalParticipant)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflow.Datas {
		UnstageBranch(stage, _data)
	}

}

func (stage *Stage) UnstageBranchDataFlowShape(dataflowshape *DataFlowShape) {

	// check if instance is already staged
	if !IsStaged(stage, dataflowshape) {
		return
	}

	dataflowshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshape.DataFlow != nil {
		UnstageBranch(stage, dataflowshape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDataShape(datashape *DataShape) {

	// check if instance is already staged
	if !IsStaged(stage, datashape) {
		return
	}

	datashape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datashape.Data != nil {
		UnstageBranch(stage, datashape.Data)
	}
	if datashape.DataFlow != nil {
		UnstageBranch(stage, datashape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

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
	for _, _externalparticipantshape := range diagramprocess.ExternalParticipant_Shapes {
		UnstageBranch(stage, _externalparticipantshape)
	}
	for _, _participant := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
		UnstageBranch(stage, _participant)
	}
	for _, _participant := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
		UnstageBranch(stage, _participant)
	}
	for _, _participant := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
		UnstageBranch(stage, _participant)
	}
	for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _taskshape := range diagramprocess.Task_Shapes {
		UnstageBranch(stage, _taskshape)
	}
	for _, _controlflow := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
		UnstageBranch(stage, _controlflow)
	}
	for _, _controlflowshape := range diagramprocess.ControlFlow_Shapes {
		UnstageBranch(stage, _controlflowshape)
	}
	for _, _dataflow := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
		UnstageBranch(stage, _dataflow)
	}
	for _, _dataflowshape := range diagramprocess.DataFlow_Shapes {
		UnstageBranch(stage, _dataflowshape)
	}
	for _, _data := range diagramprocess.DatasWhoseNodeIsExpanded {
		UnstageBranch(stage, _data)
	}
	for _, _datashape := range diagramprocess.Data_Shapes {
		UnstageBranch(stage, _datashape)
	}
	for _, _dataflow := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
		UnstageBranch(stage, _dataflow)
	}

}

func (stage *Stage) UnstageBranchExternalParticipantShape(externalparticipantshape *ExternalParticipantShape) {

	// check if instance is already staged
	if !IsStaged(stage, externalparticipantshape) {
		return
	}

	externalparticipantshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if externalparticipantshape.Participant != nil {
		UnstageBranch(stage, externalparticipantshape.Participant)
	}

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
	for _, _process := range library.RootProcesses {
		UnstageBranch(stage, _process)
	}
	for _, _process := range library.ProcesssWhoseNodeIsExpanded {
		UnstageBranch(stage, _process)
	}
	for _, _dataflow := range library.RootDataFlows {
		UnstageBranch(stage, _dataflow)
	}
	for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
		UnstageBranch(stage, _dataflow)
	}
	for _, _data := range library.RootDatas {
		UnstageBranch(stage, _data)
	}
	for _, _data := range library.DatasWhoseNodeIsExpanded {
		UnstageBranch(stage, _data)
	}
	for _, _resource := range library.RootResources {
		UnstageBranch(stage, _resource)
	}
	for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
		UnstageBranch(stage, _resource)
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
	for _, _controlflow := range participant.ControlFlows {
		UnstageBranch(stage, _controlflow)
	}
	for _, _task := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
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
	for _, _dataflow := range process.DataFlows {
		UnstageBranch(stage, _dataflow)
	}
	for _, _participant := range process.ExternalParticipants {
		UnstageBranch(stage, _participant)
	}
	for _, _participant := range process.ExternalParticipantWhoseNodeIsExpanded {
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

func (stage *Stage) UnstageBranchResource(resource *Resource) {

	// check if instance is already staged
	if !IsStaged(stage, resource) {
		return
	}

	resource.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

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
func (reference *ControlFlow) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlow) {
	// insertion point for pointers field
	if instance.Start != nil {
		reference.Start = stage.Tasks_reference[instance.Start]
	}
	if instance.End != nil {
		reference.End = stage.Tasks_reference[instance.End]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *ControlFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlowShape) {
	// insertion point for pointers field
	if instance.ControlFlow != nil {
		reference.ControlFlow = stage.ControlFlows_reference[instance.ControlFlow]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Data) GongReconstructPointersFromReferences(stage *Stage, instance *Data) {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
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

	return
}

func (reference *DataFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlowShape) {
	// insertion point for pointers field
	if instance.DataFlow != nil {
		reference.DataFlow = stage.DataFlows_reference[instance.DataFlow]
	}
	// insertion point for slice of pointers field

	return
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

	return
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

	return
}

func (reference *ExternalParticipantShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExternalParticipantShape) {
	// insertion point for pointers field
	if instance.Participant != nil {
		reference.Participant = stage.Participants_reference[instance.Participant]
	}
	// insertion point for slice of pointers field

	return
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

	return
}

func (reference *Participant) GongReconstructPointersFromReferences(stage *Stage, instance *Participant) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
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

	return
}

func (reference *ParticipantShape) GongReconstructPointersFromReferences(stage *Stage, instance *ParticipantShape) {
	// insertion point for pointers field
	if instance.Participant != nil {
		reference.Participant = stage.Participants_reference[instance.Participant]
	}
	// insertion point for slice of pointers field

	return
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

	return
}

func (reference *ProcessShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProcessShape) {
	// insertion point for pointers field
	if instance.Process != nil {
		reference.Process = stage.Processs_reference[instance.Process]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Resource) GongReconstructPointersFromReferences(stage *Stage, instance *Resource) {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *Task) GongReconstructPointersFromReferences(stage *Stage, instance *Task) {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *TaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskShape) {
	// insertion point for pointers field
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field

	return
}

// insertion point for pointer reconstruction from instances
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

	return
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

	return
}

func (reference *Data) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
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

	return
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

	return
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

	return
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

	return
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

	return
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

	return
}

func (reference *Participant) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
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

	return
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

	return
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

	return
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

	return
}

func (reference *Resource) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *Task) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
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

	return
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlflow *ControlFlow) GongDiff(stage *Stage, controlflowOther *ControlFlow) (diffs []string) {
	// insertion point for field diffs
	if controlflow.Name != controlflowOther.Name {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Name"))
	}
	if controlflow.ComputedPrefix != controlflowOther.ComputedPrefix {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "ComputedPrefix"))
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
	if data.ComputedPrefix != dataOther.ComputedPrefix {
		diffs = append(diffs, data.GongMarshallField(stage, "ComputedPrefix"))
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
	if dataflow.ComputedPrefix != dataflowOther.ComputedPrefix {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "ComputedPrefix"))
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
		ops := Diff(stage, dataflow, dataflowOther, "Datas", dataflowOther.Datas, dataflow.Datas)
		diffs = append(diffs, ops)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ExternalParticipant_Shapes", diagramprocessOther.ExternalParticipant_Shapes, diagramprocess.ExternalParticipant_Shapes)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ExternalParticipantWhoseNodeIsExpanded", diagramprocessOther.ExternalParticipantWhoseNodeIsExpanded, diagramprocess.ExternalParticipantWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded", diagramprocessOther.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded, diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded", diagramprocessOther.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded, diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "Task_Shapes", diagramprocessOther.Task_Shapes, diagramprocess.Task_Shapes)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ControlFlowsWhoseNodeIsExpanded", diagramprocessOther.ControlFlowsWhoseNodeIsExpanded, diagramprocess.ControlFlowsWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "ControlFlow_Shapes", diagramprocessOther.ControlFlow_Shapes, diagramprocess.ControlFlow_Shapes)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "DataFlowsWhoseNodeIsExpanded", diagramprocessOther.DataFlowsWhoseNodeIsExpanded, diagramprocess.DataFlowsWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "DataFlow_Shapes", diagramprocessOther.DataFlow_Shapes, diagramprocess.DataFlow_Shapes)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "DatasWhoseNodeIsExpanded", diagramprocessOther.DatasWhoseNodeIsExpanded, diagramprocess.DatasWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "Data_Shapes", diagramprocessOther.Data_Shapes, diagramprocess.Data_Shapes)
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
		ops := Diff(stage, diagramprocess, diagramprocessOther, "DataFlowsWhoseDataNodeIsExpanded", diagramprocessOther.DataFlowsWhoseDataNodeIsExpanded, diagramprocess.DataFlowsWhoseDataNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "ProcesssWhoseNodeIsExpanded", libraryOther.ProcesssWhoseNodeIsExpanded, library.ProcesssWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "RootDataFlows", libraryOther.RootDataFlows, library.RootDataFlows)
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
		ops := Diff(stage, library, libraryOther, "DataFlowsWhoseNodeIsExpanded", libraryOther.DataFlowsWhoseNodeIsExpanded, library.DataFlowsWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "RootDatas", libraryOther.RootDatas, library.RootDatas)
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
		ops := Diff(stage, library, libraryOther, "DatasWhoseNodeIsExpanded", libraryOther.DatasWhoseNodeIsExpanded, library.DatasWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "RootResources", libraryOther.RootResources, library.RootResources)
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
		ops := Diff(stage, library, libraryOther, "ResourcesWhoseNodeIsExpanded", libraryOther.ResourcesWhoseNodeIsExpanded, library.ResourcesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
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
		ops := Diff(stage, participant, participantOther, "Tasks", participantOther.Tasks, participant.Tasks)
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
		ops := Diff(stage, participant, participantOther, "ControlFlows", participantOther.ControlFlows, participant.ControlFlows)
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
		ops := Diff(stage, participant, participantOther, "TaskWhoseOutControlFlowsNodeIsExpanded", participantOther.TaskWhoseOutControlFlowsNodeIsExpanded, participant.TaskWhoseOutControlFlowsNodeIsExpanded)
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
		ops := Diff(stage, participant, participantOther, "TaskWhoseInControlFlowsNodeIsExpanded", participantOther.TaskWhoseInControlFlowsNodeIsExpanded, participant.TaskWhoseInControlFlowsNodeIsExpanded)
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
		ops := Diff(stage, participant, participantOther, "TaskWhoseOutDataFlowsNodeIsExpanded", participantOther.TaskWhoseOutDataFlowsNodeIsExpanded, participant.TaskWhoseOutDataFlowsNodeIsExpanded)
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
		ops := Diff(stage, participant, participantOther, "TaskWhoseInDataFlowsNodeIsExpanded", participantOther.TaskWhoseInDataFlowsNodeIsExpanded, participant.TaskWhoseInDataFlowsNodeIsExpanded)
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
		ops := Diff(stage, process, processOther, "DataFlows", processOther.DataFlows, process.DataFlows)
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
		ops := Diff(stage, process, processOther, "ExternalParticipants", processOther.ExternalParticipants, process.ExternalParticipants)
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
		ops := Diff(stage, process, processOther, "ExternalParticipantWhoseNodeIsExpanded", processOther.ExternalParticipantWhoseNodeIsExpanded, process.ExternalParticipantWhoseNodeIsExpanded)
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
	if resource.ComputedPrefix != resourceOther.ComputedPrefix {
		diffs = append(diffs, resource.GongMarshallField(stage, "ComputedPrefix"))
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
	if task.IsStartTask != taskOther.IsStartTask {
		diffs = append(diffs, task.GongMarshallField(stage, "IsStartTask"))
	}
	if task.IsEndTask != taskOther.IsEndTask {
		diffs = append(diffs, task.GongMarshallField(stage, "IsEndTask"))
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
