// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct DataFlow
	// insertion point per field
	stage.DataFlow_Datas_reverseMap = make(map[*Data]*DataFlow)
	for dataflow := range stage.DataFlows {
		_ = dataflow
		for _, _data := range dataflow.Datas {
			stage.DataFlow_Datas_reverseMap[_data] = dataflow
		}
	}

	// Compute reverse map for named struct DiagramProcess
	// insertion point per field
	stage.DiagramProcess_Process_Shapes_reverseMap = make(map[*ProcessShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _processshape := range diagramprocess.Process_Shapes {
			stage.DiagramProcess_Process_Shapes_reverseMap[_processshape] = diagramprocess
		}
	}
	stage.DiagramProcess_ProcesssWhoseNodeIsExpanded_reverseMap = make(map[*Process]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _process := range diagramprocess.ProcesssWhoseNodeIsExpanded {
			stage.DiagramProcess_ProcesssWhoseNodeIsExpanded_reverseMap[_process] = diagramprocess
		}
	}
	stage.DiagramProcess_Participant_Shapes_reverseMap = make(map[*ParticipantShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _participantshape := range diagramprocess.Participant_Shapes {
			stage.DiagramProcess_Participant_Shapes_reverseMap[_participantshape] = diagramprocess
		}
	}
	stage.DiagramProcess_ParticipantWhoseNodeIsExpanded_reverseMap = make(map[*Participant]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _participant := range diagramprocess.ParticipantWhoseNodeIsExpanded {
			stage.DiagramProcess_ParticipantWhoseNodeIsExpanded_reverseMap[_participant] = diagramprocess
		}
	}
	stage.DiagramProcess_ExternalParticipant_Shapes_reverseMap = make(map[*ExternalParticipantShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _externalparticipantshape := range diagramprocess.ExternalParticipant_Shapes {
			stage.DiagramProcess_ExternalParticipant_Shapes_reverseMap[_externalparticipantshape] = diagramprocess
		}
	}
	stage.DiagramProcess_ExternalParticipantWhoseNodeIsExpanded_reverseMap = make(map[*Participant]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _participant := range diagramprocess.ExternalParticipantWhoseNodeIsExpanded {
			stage.DiagramProcess_ExternalParticipantWhoseNodeIsExpanded_reverseMap[_participant] = diagramprocess
		}
	}
	stage.DiagramProcess_ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded_reverseMap = make(map[*Participant]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _participant := range diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded {
			stage.DiagramProcess_ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded_reverseMap[_participant] = diagramprocess
		}
	}
	stage.DiagramProcess_ExternalParticipantsWhoseInDataFlowsNodeIsExpanded_reverseMap = make(map[*Participant]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _participant := range diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded {
			stage.DiagramProcess_ExternalParticipantsWhoseInDataFlowsNodeIsExpanded_reverseMap[_participant] = diagramprocess
		}
	}
	stage.DiagramProcess_TasksWhoseNodeIsExpanded_reverseMap = make(map[*Task]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
			stage.DiagramProcess_TasksWhoseNodeIsExpanded_reverseMap[_task] = diagramprocess
		}
	}
	stage.DiagramProcess_Task_Shapes_reverseMap = make(map[*TaskShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _taskshape := range diagramprocess.Task_Shapes {
			stage.DiagramProcess_Task_Shapes_reverseMap[_taskshape] = diagramprocess
		}
	}
	stage.DiagramProcess_ControlFlowsWhoseNodeIsExpanded_reverseMap = make(map[*ControlFlow]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _controlflow := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
			stage.DiagramProcess_ControlFlowsWhoseNodeIsExpanded_reverseMap[_controlflow] = diagramprocess
		}
	}
	stage.DiagramProcess_ControlFlow_Shapes_reverseMap = make(map[*ControlFlowShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _controlflowshape := range diagramprocess.ControlFlow_Shapes {
			stage.DiagramProcess_ControlFlow_Shapes_reverseMap[_controlflowshape] = diagramprocess
		}
	}
	stage.DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap = make(map[*DataFlow]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _dataflow := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
			stage.DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap[_dataflow] = diagramprocess
		}
	}
	stage.DiagramProcess_DataFlow_Shapes_reverseMap = make(map[*DataFlowShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _dataflowshape := range diagramprocess.DataFlow_Shapes {
			stage.DiagramProcess_DataFlow_Shapes_reverseMap[_dataflowshape] = diagramprocess
		}
	}
	stage.DiagramProcess_DatasWhoseNodeIsExpanded_reverseMap = make(map[*Data]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _data := range diagramprocess.DatasWhoseNodeIsExpanded {
			stage.DiagramProcess_DatasWhoseNodeIsExpanded_reverseMap[_data] = diagramprocess
		}
	}
	stage.DiagramProcess_Data_Shapes_reverseMap = make(map[*DataShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _datashape := range diagramprocess.Data_Shapes {
			stage.DiagramProcess_Data_Shapes_reverseMap[_datashape] = diagramprocess
		}
	}
	stage.DiagramProcess_DataFlowsWhoseDataNodeIsExpanded_reverseMap = make(map[*DataFlow]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _dataflow := range diagramprocess.DataFlowsWhoseDataNodeIsExpanded {
			stage.DiagramProcess_DataFlowsWhoseDataNodeIsExpanded_reverseMap[_dataflow] = diagramprocess
		}
	}
	stage.DiagramProcess_AllocatedResourcesWhoseNodeIsExpanded_reverseMap = make(map[*Resource]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _resource := range diagramprocess.AllocatedResourcesWhoseNodeIsExpanded {
			stage.DiagramProcess_AllocatedResourcesWhoseNodeIsExpanded_reverseMap[_resource] = diagramprocess
		}
	}
	stage.DiagramProcess_AllocatedResourceShapes_reverseMap = make(map[*AllocatedResourceShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _allocatedresourceshape := range diagramprocess.AllocatedResourceShapes {
			stage.DiagramProcess_AllocatedResourceShapes_reverseMap[_allocatedresourceshape] = diagramprocess
		}
	}
	stage.DiagramProcess_AllocatedProcessesWhoseNodeIsExpanded_reverseMap = make(map[*Process]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _process := range diagramprocess.AllocatedProcessesWhoseNodeIsExpanded {
			stage.DiagramProcess_AllocatedProcessesWhoseNodeIsExpanded_reverseMap[_process] = diagramprocess
		}
	}
	stage.DiagramProcess_AllocatedProcessShapes_reverseMap = make(map[*AllocatedProcessShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _allocatedprocessshape := range diagramprocess.AllocatedProcessShapes {
			stage.DiagramProcess_AllocatedProcessShapes_reverseMap[_allocatedprocessshape] = diagramprocess
		}
	}
	stage.DiagramProcess_Note_Shapes_reverseMap = make(map[*NoteShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _noteshape := range diagramprocess.Note_Shapes {
			stage.DiagramProcess_Note_Shapes_reverseMap[_noteshape] = diagramprocess
		}
	}
	stage.DiagramProcess_NotesWhoseNodeIsExpanded_reverseMap = make(map[*Note]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _note := range diagramprocess.NotesWhoseNodeIsExpanded {
			stage.DiagramProcess_NotesWhoseNodeIsExpanded_reverseMap[_note] = diagramprocess
		}
	}
	stage.DiagramProcess_NoteTaskShapes_reverseMap = make(map[*NoteTaskShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _notetaskshape := range diagramprocess.NoteTaskShapes {
			stage.DiagramProcess_NoteTaskShapes_reverseMap[_notetaskshape] = diagramprocess
		}
	}

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}
	stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[_library] = library
		}
	}
	stage.Library_RootProcesses_reverseMap = make(map[*Process]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _process := range library.RootProcesses {
			stage.Library_RootProcesses_reverseMap[_process] = library
		}
	}
	stage.Library_ProcesssWhoseNodeIsExpanded_reverseMap = make(map[*Process]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _process := range library.ProcesssWhoseNodeIsExpanded {
			stage.Library_ProcesssWhoseNodeIsExpanded_reverseMap[_process] = library
		}
	}
	stage.Library_RootDataFlows_reverseMap = make(map[*DataFlow]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _dataflow := range library.RootDataFlows {
			stage.Library_RootDataFlows_reverseMap[_dataflow] = library
		}
	}
	stage.Library_DataFlowsWhoseNodeIsExpanded_reverseMap = make(map[*DataFlow]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
			stage.Library_DataFlowsWhoseNodeIsExpanded_reverseMap[_dataflow] = library
		}
	}
	stage.Library_RootDatas_reverseMap = make(map[*Data]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _data := range library.RootDatas {
			stage.Library_RootDatas_reverseMap[_data] = library
		}
	}
	stage.Library_DatasWhoseNodeIsExpanded_reverseMap = make(map[*Data]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _data := range library.DatasWhoseNodeIsExpanded {
			stage.Library_DatasWhoseNodeIsExpanded_reverseMap[_data] = library
		}
	}
	stage.Library_RootResources_reverseMap = make(map[*Resource]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _resource := range library.RootResources {
			stage.Library_RootResources_reverseMap[_resource] = library
		}
	}
	stage.Library_ResourcesWhoseNodeIsExpanded_reverseMap = make(map[*Resource]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
			stage.Library_ResourcesWhoseNodeIsExpanded_reverseMap[_resource] = library
		}
	}
	stage.Library_ParticipantsWhoseNodeIsExpanded_reverseMap = make(map[*Participant]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _participant := range library.ParticipantsWhoseNodeIsExpanded {
			stage.Library_ParticipantsWhoseNodeIsExpanded_reverseMap[_participant] = library
		}
	}
	stage.Library_RootNotes_reverseMap = make(map[*Note]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _note := range library.RootNotes {
			stage.Library_RootNotes_reverseMap[_note] = library
		}
	}
	stage.Library_NotesWhoseNodeIsExpanded_reverseMap = make(map[*Note]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _note := range library.NotesWhoseNodeIsExpanded {
			stage.Library_NotesWhoseNodeIsExpanded_reverseMap[_note] = library
		}
	}

	// Compute reverse map for named struct Note
	// insertion point per field
	stage.Note_Tasks_reverseMap = make(map[*Task]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _task := range note.Tasks {
			stage.Note_Tasks_reverseMap[_task] = note
		}
	}

	// Compute reverse map for named struct Participant
	// insertion point per field
	stage.Participant_Resources_reverseMap = make(map[*Resource]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _resource := range participant.Resources {
			stage.Participant_Resources_reverseMap[_resource] = participant
		}
	}
	stage.Participant_Processes_reverseMap = make(map[*Process]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _process := range participant.Processes {
			stage.Participant_Processes_reverseMap[_process] = participant
		}
	}
	stage.Participant_Tasks_reverseMap = make(map[*Task]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _task := range participant.Tasks {
			stage.Participant_Tasks_reverseMap[_task] = participant
		}
	}
	stage.Participant_ControlFlows_reverseMap = make(map[*ControlFlow]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _controlflow := range participant.ControlFlows {
			stage.Participant_ControlFlows_reverseMap[_controlflow] = participant
		}
	}
	stage.Participant_TaskWhoseOutControlFlowsNodeIsExpanded_reverseMap = make(map[*Task]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _task := range participant.TaskWhoseOutControlFlowsNodeIsExpanded {
			stage.Participant_TaskWhoseOutControlFlowsNodeIsExpanded_reverseMap[_task] = participant
		}
	}
	stage.Participant_TaskWhoseInControlFlowsNodeIsExpanded_reverseMap = make(map[*Task]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _task := range participant.TaskWhoseInControlFlowsNodeIsExpanded {
			stage.Participant_TaskWhoseInControlFlowsNodeIsExpanded_reverseMap[_task] = participant
		}
	}
	stage.Participant_TaskWhoseOutDataFlowsNodeIsExpanded_reverseMap = make(map[*Task]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _task := range participant.TaskWhoseOutDataFlowsNodeIsExpanded {
			stage.Participant_TaskWhoseOutDataFlowsNodeIsExpanded_reverseMap[_task] = participant
		}
	}
	stage.Participant_TaskWhoseInDataFlowsNodeIsExpanded_reverseMap = make(map[*Task]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _task := range participant.TaskWhoseInDataFlowsNodeIsExpanded {
			stage.Participant_TaskWhoseInDataFlowsNodeIsExpanded_reverseMap[_task] = participant
		}
	}

	// Compute reverse map for named struct Process
	// insertion point per field
	stage.Process_DiagramProcesss_reverseMap = make(map[*DiagramProcess]*Process)
	for process := range stage.Processs {
		_ = process
		for _, _diagramprocess := range process.DiagramProcesss {
			stage.Process_DiagramProcesss_reverseMap[_diagramprocess] = process
		}
	}
	stage.Process_DiagramProcessWhoseNodeIsExpanded_reverseMap = make(map[*DiagramProcess]*Process)
	for process := range stage.Processs {
		_ = process
		for _, _diagramprocess := range process.DiagramProcessWhoseNodeIsExpanded {
			stage.Process_DiagramProcessWhoseNodeIsExpanded_reverseMap[_diagramprocess] = process
		}
	}
	stage.Process_SubProcesses_reverseMap = make(map[*Process]*Process)
	for process := range stage.Processs {
		_ = process
		for _, _process := range process.SubProcesses {
			stage.Process_SubProcesses_reverseMap[_process] = process
		}
	}
	stage.Process_Participants_reverseMap = make(map[*Participant]*Process)
	for process := range stage.Processs {
		_ = process
		for _, _participant := range process.Participants {
			stage.Process_Participants_reverseMap[_participant] = process
		}
	}
	stage.Process_ParticipantWhoseNodeIsExpanded_reverseMap = make(map[*Participant]*Process)
	for process := range stage.Processs {
		_ = process
		for _, _participant := range process.ParticipantWhoseNodeIsExpanded {
			stage.Process_ParticipantWhoseNodeIsExpanded_reverseMap[_participant] = process
		}
	}
	stage.Process_DataFlows_reverseMap = make(map[*DataFlow]*Process)
	for process := range stage.Processs {
		_ = process
		for _, _dataflow := range process.DataFlows {
			stage.Process_DataFlows_reverseMap[_dataflow] = process
		}
	}
	stage.Process_ExternalParticipants_reverseMap = make(map[*Participant]*Process)
	for process := range stage.Processs {
		_ = process
		for _, _participant := range process.ExternalParticipants {
			stage.Process_ExternalParticipants_reverseMap[_participant] = process
		}
	}
	stage.Process_ExternalParticipantWhoseNodeIsExpanded_reverseMap = make(map[*Participant]*Process)
	for process := range stage.Processs {
		_ = process
		for _, _participant := range process.ExternalParticipantWhoseNodeIsExpanded {
			stage.Process_ExternalParticipantWhoseNodeIsExpanded_reverseMap[_participant] = process
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AllocatedProcessShapes)

	res = __gong__appendInstances(res, stage.AllocatedResourceShapes)

	res = __gong__appendInstances(res, stage.ControlFlows)

	res = __gong__appendInstances(res, stage.ControlFlowShapes)

	res = __gong__appendInstances(res, stage.Datas)

	res = __gong__appendInstances(res, stage.DataFlows)

	res = __gong__appendInstances(res, stage.DataFlowShapes)

	res = __gong__appendInstances(res, stage.DataShapes)

	res = __gong__appendInstances(res, stage.DiagramProcesss)

	res = __gong__appendInstances(res, stage.ExternalParticipantShapes)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.Notes)

	res = __gong__appendInstances(res, stage.NoteShapes)

	res = __gong__appendInstances(res, stage.NoteTaskShapes)

	res = __gong__appendInstances(res, stage.Participants)

	res = __gong__appendInstances(res, stage.ParticipantShapes)

	res = __gong__appendInstances(res, stage.Processs)

	res = __gong__appendInstances(res, stage.ProcessShapes)

	res = __gong__appendInstances(res, stage.Resources)

	res = __gong__appendInstances(res, stage.Tasks)

	res = __gong__appendInstances(res, stage.TaskShapes)

	return
}

// insertion point per named struct
func (allocatedprocessshape *AllocatedProcessShape) GongCopy() GongstructIF {
	newInstance := new(AllocatedProcessShape)
	allocatedprocessshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (allocatedresourceshape *AllocatedResourceShape) GongCopy() GongstructIF {
	newInstance := new(AllocatedResourceShape)
	allocatedresourceshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (controlflow *ControlFlow) GongCopy() GongstructIF {
	newInstance := new(ControlFlow)
	controlflow.GongCopyBasicFields(newInstance)
	return newInstance
}

func (controlflowshape *ControlFlowShape) GongCopy() GongstructIF {
	newInstance := new(ControlFlowShape)
	controlflowshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (data *Data) GongCopy() GongstructIF {
	newInstance := new(Data)
	data.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dataflow *DataFlow) GongCopy() GongstructIF {
	newInstance := new(DataFlow)
	dataflow.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dataflowshape *DataFlowShape) GongCopy() GongstructIF {
	newInstance := new(DataFlowShape)
	dataflowshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datashape *DataShape) GongCopy() GongstructIF {
	newInstance := new(DataShape)
	datashape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagramprocess *DiagramProcess) GongCopy() GongstructIF {
	newInstance := new(DiagramProcess)
	diagramprocess.GongCopyBasicFields(newInstance)
	return newInstance
}

func (externalparticipantshape *ExternalParticipantShape) GongCopy() GongstructIF {
	newInstance := new(ExternalParticipantShape)
	externalparticipantshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.GongCopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notetaskshape *NoteTaskShape) GongCopy() GongstructIF {
	newInstance := new(NoteTaskShape)
	notetaskshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (participant *Participant) GongCopy() GongstructIF {
	newInstance := new(Participant)
	participant.GongCopyBasicFields(newInstance)
	return newInstance
}

func (participantshape *ParticipantShape) GongCopy() GongstructIF {
	newInstance := new(ParticipantShape)
	participantshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (process *Process) GongCopy() GongstructIF {
	newInstance := new(Process)
	process.GongCopyBasicFields(newInstance)
	return newInstance
}

func (processshape *ProcessShape) GongCopy() GongstructIF {
	newInstance := new(ProcessShape)
	processshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (resource *Resource) GongCopy() GongstructIF {
	newInstance := new(Resource)
	resource.GongCopyBasicFields(newInstance)
	return newInstance
}

func (task *Task) GongCopy() GongstructIF {
	newInstance := new(Task)
	task.GongCopyBasicFields(newInstance)
	return newInstance
}

func (taskshape *TaskShape) GongCopy() GongstructIF {
	newInstance := new(TaskShape)
	taskshape.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (allocatedprocessshape *AllocatedProcessShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, allocatedprocessshape)
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, allocatedresourceshape)
}

func (controlflow *ControlFlow) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, controlflow)
}

func (controlflowshape *ControlFlowShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, controlflowshape)
}

func (data *Data) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, data)
}

func (dataflow *DataFlow) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dataflow)
}

func (dataflowshape *DataFlowShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dataflowshape)
}

func (datashape *DataShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datashape)
}

func (diagramprocess *DiagramProcess) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagramprocess)
}

func (externalparticipantshape *ExternalParticipantShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, externalparticipantshape)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (note *Note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note)
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteshape)
}

func (notetaskshape *NoteTaskShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notetaskshape)
}

func (participant *Participant) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, participant)
}

func (participantshape *ParticipantShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, participantshape)
}

func (process *Process) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, process)
}

func (processshape *ProcessShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, processshape)
}

func (resource *Resource) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, resource)
}

func (task *Task) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, task)
}

func (taskshape *TaskShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, taskshape)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.AllocatedProcessShapes,
		stage.AllocatedProcessShape_stagedOrder,
		stage.AllocatedProcessShapes_reference,
		&stage.AllocatedProcessShapes_referenceOrder,
		stage.AllocatedProcessShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AllocatedResourceShapes,
		stage.AllocatedResourceShape_stagedOrder,
		stage.AllocatedResourceShapes_reference,
		&stage.AllocatedResourceShapes_referenceOrder,
		stage.AllocatedResourceShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ControlFlows,
		stage.ControlFlow_stagedOrder,
		stage.ControlFlows_reference,
		&stage.ControlFlows_referenceOrder,
		stage.ControlFlows_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ControlFlowShapes,
		stage.ControlFlowShape_stagedOrder,
		stage.ControlFlowShapes_reference,
		&stage.ControlFlowShapes_referenceOrder,
		stage.ControlFlowShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Datas,
		stage.Data_stagedOrder,
		stage.Datas_reference,
		&stage.Datas_referenceOrder,
		stage.Datas_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DataFlows,
		stage.DataFlow_stagedOrder,
		stage.DataFlows_reference,
		&stage.DataFlows_referenceOrder,
		stage.DataFlows_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DataFlowShapes,
		stage.DataFlowShape_stagedOrder,
		stage.DataFlowShapes_reference,
		&stage.DataFlowShapes_referenceOrder,
		stage.DataFlowShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DataShapes,
		stage.DataShape_stagedOrder,
		stage.DataShapes_reference,
		&stage.DataShapes_referenceOrder,
		stage.DataShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DiagramProcesss,
		stage.DiagramProcess_stagedOrder,
		stage.DiagramProcesss_reference,
		&stage.DiagramProcesss_referenceOrder,
		stage.DiagramProcesss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ExternalParticipantShapes,
		stage.ExternalParticipantShape_stagedOrder,
		stage.ExternalParticipantShapes_reference,
		&stage.ExternalParticipantShapes_referenceOrder,
		stage.ExternalParticipantShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Librarys,
		stage.Library_stagedOrder,
		stage.Librarys_reference,
		&stage.Librarys_referenceOrder,
		stage.Librarys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Notes,
		stage.Note_stagedOrder,
		stage.Notes_reference,
		&stage.Notes_referenceOrder,
		stage.Notes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteShapes,
		stage.NoteShape_stagedOrder,
		stage.NoteShapes_reference,
		&stage.NoteShapes_referenceOrder,
		stage.NoteShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.NoteTaskShapes,
		stage.NoteTaskShape_stagedOrder,
		stage.NoteTaskShapes_reference,
		&stage.NoteTaskShapes_referenceOrder,
		stage.NoteTaskShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Participants,
		stage.Participant_stagedOrder,
		stage.Participants_reference,
		&stage.Participants_referenceOrder,
		stage.Participants_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ParticipantShapes,
		stage.ParticipantShape_stagedOrder,
		stage.ParticipantShapes_reference,
		&stage.ParticipantShapes_referenceOrder,
		stage.ParticipantShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Processs,
		stage.Process_stagedOrder,
		stage.Processs_reference,
		&stage.Processs_referenceOrder,
		stage.Processs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ProcessShapes,
		stage.ProcessShape_stagedOrder,
		stage.ProcessShapes_reference,
		&stage.ProcessShapes_referenceOrder,
		stage.ProcessShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Resources,
		stage.Resource_stagedOrder,
		stage.Resources_reference,
		&stage.Resources_referenceOrder,
		stage.Resources_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tasks,
		stage.Task_stagedOrder,
		stage.Tasks_reference,
		&stage.Tasks_referenceOrder,
		stage.Tasks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.TaskShapes,
		stage.TaskShape_stagedOrder,
		stage.TaskShapes_reference,
		&stage.TaskShapes_referenceOrder,
		stage.TaskShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.AllocatedProcessShapes, &stage.AllocatedProcessShapes_reference, &stage.AllocatedProcessShapes_referenceOrder, &stage.AllocatedProcessShapes_instance)

	__gong__computeReferencePass1(stage, stage.AllocatedResourceShapes, &stage.AllocatedResourceShapes_reference, &stage.AllocatedResourceShapes_referenceOrder, &stage.AllocatedResourceShapes_instance)

	__gong__computeReferencePass1(stage, stage.ControlFlows, &stage.ControlFlows_reference, &stage.ControlFlows_referenceOrder, &stage.ControlFlows_instance)

	__gong__computeReferencePass1(stage, stage.ControlFlowShapes, &stage.ControlFlowShapes_reference, &stage.ControlFlowShapes_referenceOrder, &stage.ControlFlowShapes_instance)

	__gong__computeReferencePass1(stage, stage.Datas, &stage.Datas_reference, &stage.Datas_referenceOrder, &stage.Datas_instance)

	__gong__computeReferencePass1(stage, stage.DataFlows, &stage.DataFlows_reference, &stage.DataFlows_referenceOrder, &stage.DataFlows_instance)

	__gong__computeReferencePass1(stage, stage.DataFlowShapes, &stage.DataFlowShapes_reference, &stage.DataFlowShapes_referenceOrder, &stage.DataFlowShapes_instance)

	__gong__computeReferencePass1(stage, stage.DataShapes, &stage.DataShapes_reference, &stage.DataShapes_referenceOrder, &stage.DataShapes_instance)

	__gong__computeReferencePass1(stage, stage.DiagramProcesss, &stage.DiagramProcesss_reference, &stage.DiagramProcesss_referenceOrder, &stage.DiagramProcesss_instance)

	__gong__computeReferencePass1(stage, stage.ExternalParticipantShapes, &stage.ExternalParticipantShapes_reference, &stage.ExternalParticipantShapes_referenceOrder, &stage.ExternalParticipantShapes_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.Notes, &stage.Notes_reference, &stage.Notes_referenceOrder, &stage.Notes_instance)

	__gong__computeReferencePass1(stage, stage.NoteShapes, &stage.NoteShapes_reference, &stage.NoteShapes_referenceOrder, &stage.NoteShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteTaskShapes, &stage.NoteTaskShapes_reference, &stage.NoteTaskShapes_referenceOrder, &stage.NoteTaskShapes_instance)

	__gong__computeReferencePass1(stage, stage.Participants, &stage.Participants_reference, &stage.Participants_referenceOrder, &stage.Participants_instance)

	__gong__computeReferencePass1(stage, stage.ParticipantShapes, &stage.ParticipantShapes_reference, &stage.ParticipantShapes_referenceOrder, &stage.ParticipantShapes_instance)

	__gong__computeReferencePass1(stage, stage.Processs, &stage.Processs_reference, &stage.Processs_referenceOrder, &stage.Processs_instance)

	__gong__computeReferencePass1(stage, stage.ProcessShapes, &stage.ProcessShapes_reference, &stage.ProcessShapes_referenceOrder, &stage.ProcessShapes_instance)

	__gong__computeReferencePass1(stage, stage.Resources, &stage.Resources_reference, &stage.Resources_referenceOrder, &stage.Resources_instance)

	__gong__computeReferencePass1(stage, stage.Tasks, &stage.Tasks_reference, &stage.Tasks_referenceOrder, &stage.Tasks_instance)

	__gong__computeReferencePass1(stage, stage.TaskShapes, &stage.TaskShapes_reference, &stage.TaskShapes_referenceOrder, &stage.TaskShapes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AllocatedProcessShapes, stage.AllocatedProcessShapes_reference, stage)

	__gong__computeReferencePass2(stage.AllocatedResourceShapes, stage.AllocatedResourceShapes_reference, stage)

	__gong__computeReferencePass2(stage.ControlFlows, stage.ControlFlows_reference, stage)

	__gong__computeReferencePass2(stage.ControlFlowShapes, stage.ControlFlowShapes_reference, stage)

	__gong__computeReferencePass2(stage.Datas, stage.Datas_reference, stage)

	__gong__computeReferencePass2(stage.DataFlows, stage.DataFlows_reference, stage)

	__gong__computeReferencePass2(stage.DataFlowShapes, stage.DataFlowShapes_reference, stage)

	__gong__computeReferencePass2(stage.DataShapes, stage.DataShapes_reference, stage)

	__gong__computeReferencePass2(stage.DiagramProcesss, stage.DiagramProcesss_reference, stage)

	__gong__computeReferencePass2(stage.ExternalParticipantShapes, stage.ExternalParticipantShapes_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.Notes, stage.Notes_reference, stage)

	__gong__computeReferencePass2(stage.NoteShapes, stage.NoteShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteTaskShapes, stage.NoteTaskShapes_reference, stage)

	__gong__computeReferencePass2(stage.Participants, stage.Participants_reference, stage)

	__gong__computeReferencePass2(stage.ParticipantShapes, stage.ParticipantShapes_reference, stage)

	__gong__computeReferencePass2(stage.Processs, stage.Processs_reference, stage)

	__gong__computeReferencePass2(stage.ProcessShapes, stage.ProcessShapes_reference, stage)

	__gong__computeReferencePass2(stage.Resources, stage.Resources_reference, stage)

	__gong__computeReferencePass2(stage.Tasks, stage.Tasks_reference, stage)

	__gong__computeReferencePass2(stage.TaskShapes, stage.TaskShapes_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (allocatedprocessshape *AllocatedProcessShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AllocatedProcessShape_stagedOrder, stage.AllocatedProcessShapes_referenceOrder, allocatedprocessshape, "AllocatedProcessShape")
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AllocatedResourceShape_stagedOrder, stage.AllocatedResourceShapes_referenceOrder, allocatedresourceshape, "AllocatedResourceShape")
}

func (controlflow *ControlFlow) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ControlFlow_stagedOrder, stage.ControlFlows_referenceOrder, controlflow, "ControlFlow")
}

func (controlflowshape *ControlFlowShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ControlFlowShape_stagedOrder, stage.ControlFlowShapes_referenceOrder, controlflowshape, "ControlFlowShape")
}

func (data *Data) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Data_stagedOrder, stage.Datas_referenceOrder, data, "Data")
}

func (dataflow *DataFlow) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DataFlow_stagedOrder, stage.DataFlows_referenceOrder, dataflow, "DataFlow")
}

func (dataflowshape *DataFlowShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DataFlowShape_stagedOrder, stage.DataFlowShapes_referenceOrder, dataflowshape, "DataFlowShape")
}

func (datashape *DataShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DataShape_stagedOrder, stage.DataShapes_referenceOrder, datashape, "DataShape")
}

func (diagramprocess *DiagramProcess) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DiagramProcess_stagedOrder, stage.DiagramProcesss_referenceOrder, diagramprocess, "DiagramProcess")
}

func (externalparticipantshape *ExternalParticipantShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ExternalParticipantShape_stagedOrder, stage.ExternalParticipantShapes_referenceOrder, externalparticipantshape, "ExternalParticipantShape")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_stagedOrder, stage.Notes_referenceOrder, note, "Note")
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteShape_stagedOrder, stage.NoteShapes_referenceOrder, noteshape, "NoteShape")
}

func (notetaskshape *NoteTaskShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteTaskShape_stagedOrder, stage.NoteTaskShapes_referenceOrder, notetaskshape, "NoteTaskShape")
}

func (participant *Participant) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Participant_stagedOrder, stage.Participants_referenceOrder, participant, "Participant")
}

func (participantshape *ParticipantShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ParticipantShape_stagedOrder, stage.ParticipantShapes_referenceOrder, participantshape, "ParticipantShape")
}

func (process *Process) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Process_stagedOrder, stage.Processs_referenceOrder, process, "Process")
}

func (processshape *ProcessShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ProcessShape_stagedOrder, stage.ProcessShapes_referenceOrder, processshape, "ProcessShape")
}

func (resource *Resource) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Resource_stagedOrder, stage.Resources_referenceOrder, resource, "Resource")
}

func (task *Task) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Task_stagedOrder, stage.Tasks_referenceOrder, task, "Task")
}

func (taskshape *TaskShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.TaskShape_stagedOrder, stage.TaskShapes_referenceOrder, taskshape, "TaskShape")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (allocatedprocessshape *AllocatedProcessShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(allocatedprocessshape, allocatedprocessshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (allocatedprocessshape *AllocatedProcessShape) GongGetReferenceIdentifier(stage *Stage) string {
	return allocatedprocessshape.GongGetIdentifier(stage)
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(allocatedresourceshape, allocatedresourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (allocatedresourceshape *AllocatedResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return allocatedresourceshape.GongGetIdentifier(stage)
}

func (controlflow *ControlFlow) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(controlflow, controlflow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlflow *ControlFlow) GongGetReferenceIdentifier(stage *Stage) string {
	return controlflow.GongGetIdentifier(stage)
}

func (controlflowshape *ControlFlowShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(controlflowshape, controlflowshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlflowshape *ControlFlowShape) GongGetReferenceIdentifier(stage *Stage) string {
	return controlflowshape.GongGetIdentifier(stage)
}

func (data *Data) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(data, data.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (data *Data) GongGetReferenceIdentifier(stage *Stage) string {
	return data.GongGetIdentifier(stage)
}

func (dataflow *DataFlow) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dataflow, dataflow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dataflow *DataFlow) GongGetReferenceIdentifier(stage *Stage) string {
	return dataflow.GongGetIdentifier(stage)
}

func (dataflowshape *DataFlowShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dataflowshape, dataflowshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dataflowshape *DataFlowShape) GongGetReferenceIdentifier(stage *Stage) string {
	return dataflowshape.GongGetIdentifier(stage)
}

func (datashape *DataShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datashape, datashape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datashape *DataShape) GongGetReferenceIdentifier(stage *Stage) string {
	return datashape.GongGetIdentifier(stage)
}

func (diagramprocess *DiagramProcess) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagramprocess, diagramprocess.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramprocess *DiagramProcess) GongGetReferenceIdentifier(stage *Stage) string {
	return diagramprocess.GongGetIdentifier(stage)
}

func (externalparticipantshape *ExternalParticipantShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(externalparticipantshape, externalparticipantshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (externalparticipantshape *ExternalParticipantShape) GongGetReferenceIdentifier(stage *Stage) string {
	return externalparticipantshape.GongGetIdentifier(stage)
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(library, library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return library.GongGetIdentifier(stage)
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(note, note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return note.GongGetIdentifier(stage)
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteshape, noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteshape.GongGetIdentifier(stage)
}

func (notetaskshape *NoteTaskShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notetaskshape, notetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetaskshape *NoteTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return notetaskshape.GongGetIdentifier(stage)
}

func (participant *Participant) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(participant, participant.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (participant *Participant) GongGetReferenceIdentifier(stage *Stage) string {
	return participant.GongGetIdentifier(stage)
}

func (participantshape *ParticipantShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(participantshape, participantshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (participantshape *ParticipantShape) GongGetReferenceIdentifier(stage *Stage) string {
	return participantshape.GongGetIdentifier(stage)
}

func (process *Process) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(process, process.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (process *Process) GongGetReferenceIdentifier(stage *Stage) string {
	return process.GongGetIdentifier(stage)
}

func (processshape *ProcessShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(processshape, processshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (processshape *ProcessShape) GongGetReferenceIdentifier(stage *Stage) string {
	return processshape.GongGetIdentifier(stage)
}

func (resource *Resource) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(resource, resource.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resource *Resource) GongGetReferenceIdentifier(stage *Stage) string {
	return resource.GongGetIdentifier(stage)
}

func (task *Task) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(task, task.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (task *Task) GongGetReferenceIdentifier(stage *Stage) string {
	return task.GongGetIdentifier(stage)
}

func (taskshape *TaskShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(taskshape, taskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskshape *TaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return taskshape.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (allocatedprocessshape *AllocatedProcessShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(allocatedprocessshape.GongGetIdentifier(stage), "AllocatedProcessShape", allocatedprocessshape.Name)
}

func (allocatedresourceshape *AllocatedResourceShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(allocatedresourceshape.GongGetIdentifier(stage), "AllocatedResourceShape", allocatedresourceshape.Name)
}

func (controlflow *ControlFlow) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(controlflow.GongGetIdentifier(stage), "ControlFlow", controlflow.Name)
}

func (controlflowshape *ControlFlowShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(controlflowshape.GongGetIdentifier(stage), "ControlFlowShape", controlflowshape.Name)
}

func (data *Data) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(data.GongGetIdentifier(stage), "Data", data.Name)
}

func (dataflow *DataFlow) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dataflow.GongGetIdentifier(stage), "DataFlow", dataflow.Name)
}

func (dataflowshape *DataFlowShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dataflowshape.GongGetIdentifier(stage), "DataFlowShape", dataflowshape.Name)
}

func (datashape *DataShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datashape.GongGetIdentifier(stage), "DataShape", datashape.Name)
}

func (diagramprocess *DiagramProcess) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagramprocess.GongGetIdentifier(stage), "DiagramProcess", diagramprocess.Name)
}

func (externalparticipantshape *ExternalParticipantShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(externalparticipantshape.GongGetIdentifier(stage), "ExternalParticipantShape", externalparticipantshape.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (note *Note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note.GongGetIdentifier(stage), "Note", note.Name)
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteshape.GongGetIdentifier(stage), "NoteShape", noteshape.Name)
}

func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notetaskshape.GongGetIdentifier(stage), "NoteTaskShape", notetaskshape.Name)
}

func (participant *Participant) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(participant.GongGetIdentifier(stage), "Participant", participant.Name)
}

func (participantshape *ParticipantShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(participantshape.GongGetIdentifier(stage), "ParticipantShape", participantshape.Name)
}

func (process *Process) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(process.GongGetIdentifier(stage), "Process", process.Name)
}

func (processshape *ProcessShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(processshape.GongGetIdentifier(stage), "ProcessShape", processshape.Name)
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(resource.GongGetIdentifier(stage), "Resource", resource.Name)
}

func (task *Task) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(task.GongGetIdentifier(stage), "Task", task.Name)
}

func (taskshape *TaskShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(taskshape.GongGetIdentifier(stage), "TaskShape", taskshape.Name)
}

// insertion point for unstaging
func (allocatedprocessshape *AllocatedProcessShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(allocatedprocessshape.GongGetReferenceIdentifier(stage))
}

func (allocatedresourceshape *AllocatedResourceShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(allocatedresourceshape.GongGetReferenceIdentifier(stage))
}

func (controlflow *ControlFlow) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(controlflow.GongGetReferenceIdentifier(stage))
}

func (controlflowshape *ControlFlowShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(controlflowshape.GongGetReferenceIdentifier(stage))
}

func (data *Data) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(data.GongGetReferenceIdentifier(stage))
}

func (dataflow *DataFlow) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dataflow.GongGetReferenceIdentifier(stage))
}

func (dataflowshape *DataFlowShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dataflowshape.GongGetReferenceIdentifier(stage))
}

func (datashape *DataShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datashape.GongGetReferenceIdentifier(stage))
}

func (diagramprocess *DiagramProcess) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagramprocess.GongGetReferenceIdentifier(stage))
}

func (externalparticipantshape *ExternalParticipantShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(externalparticipantshape.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (note *Note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note.GongGetReferenceIdentifier(stage))
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteshape.GongGetReferenceIdentifier(stage))
}

func (notetaskshape *NoteTaskShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notetaskshape.GongGetReferenceIdentifier(stage))
}

func (participant *Participant) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(participant.GongGetReferenceIdentifier(stage))
}

func (participantshape *ParticipantShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(participantshape.GongGetReferenceIdentifier(stage))
}

func (process *Process) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(process.GongGetReferenceIdentifier(stage))
}

func (processshape *ProcessShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(processshape.GongGetReferenceIdentifier(stage))
}

func (resource *Resource) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(resource.GongGetReferenceIdentifier(stage))
}

func (task *Task) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(task.GongGetReferenceIdentifier(stage))
}

func (taskshape *TaskShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(taskshape.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
