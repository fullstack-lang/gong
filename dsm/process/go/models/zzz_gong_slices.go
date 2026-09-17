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
	// Compute reverse map for named struct AllocatedProcessShape
	// insertion point per field

	// Compute reverse map for named struct AllocatedResourceShape
	// insertion point per field

	// Compute reverse map for named struct ControlFlow
	// insertion point per field

	// Compute reverse map for named struct ControlFlowShape
	// insertion point per field

	// Compute reverse map for named struct Data
	// insertion point per field

	// Compute reverse map for named struct DataFlow
	// insertion point per field
	stage.DataFlow_Datas_reverseMap = make(map[*Data]*DataFlow)
	for dataflow := range stage.DataFlows {
		_ = dataflow
		for _, _data := range dataflow.Datas {
			stage.DataFlow_Datas_reverseMap[_data] = dataflow
		}
	}

	// Compute reverse map for named struct DataFlowShape
	// insertion point per field

	// Compute reverse map for named struct DataShape
	// insertion point per field

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

	// Compute reverse map for named struct ExternalParticipantShape
	// insertion point per field

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

	// Compute reverse map for named struct NoteShape
	// insertion point per field

	// Compute reverse map for named struct NoteTaskShape
	// insertion point per field

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

	// Compute reverse map for named struct ParticipantShape
	// insertion point per field

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

	// Compute reverse map for named struct ProcessShape
	// insertion point per field

	// Compute reverse map for named struct Resource
	// insertion point per field

	// Compute reverse map for named struct Task
	// insertion point per field

	// Compute reverse map for named struct TaskShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.AllocatedProcessShapes {
		res = append(res, instance)
	}

	for instance := range stage.AllocatedResourceShapes {
		res = append(res, instance)
	}

	for instance := range stage.ControlFlows {
		res = append(res, instance)
	}

	for instance := range stage.ControlFlowShapes {
		res = append(res, instance)
	}

	for instance := range stage.Datas {
		res = append(res, instance)
	}

	for instance := range stage.DataFlows {
		res = append(res, instance)
	}

	for instance := range stage.DataFlowShapes {
		res = append(res, instance)
	}

	for instance := range stage.DataShapes {
		res = append(res, instance)
	}

	for instance := range stage.DiagramProcesss {
		res = append(res, instance)
	}

	for instance := range stage.ExternalParticipantShapes {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.NoteShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteTaskShapes {
		res = append(res, instance)
	}

	for instance := range stage.Participants {
		res = append(res, instance)
	}

	for instance := range stage.ParticipantShapes {
		res = append(res, instance)
	}

	for instance := range stage.Processs {
		res = append(res, instance)
	}

	for instance := range stage.ProcessShapes {
		res = append(res, instance)
	}

	for instance := range stage.Resources {
		res = append(res, instance)
	}

	for instance := range stage.Tasks {
		res = append(res, instance)
	}

	for instance := range stage.TaskShapes {
		res = append(res, instance)
	}

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
func (allocatedprocessshape *AllocatedProcessShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(allocatedprocessshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(allocatedprocessshape), uint64(stage.GetOrder(allocatedprocessshape)))
	return
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(allocatedresourceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(allocatedresourceshape), uint64(stage.GetOrder(allocatedresourceshape)))
	return
}

func (controlflow *ControlFlow) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlflow).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(controlflow), uint64(stage.GetOrder(controlflow)))
	return
}

func (controlflowshape *ControlFlowShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlflowshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(controlflowshape), uint64(stage.GetOrder(controlflowshape)))
	return
}

func (data *Data) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(data).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(data), uint64(stage.GetOrder(data)))
	return
}

func (dataflow *DataFlow) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(dataflow).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(dataflow), uint64(stage.GetOrder(dataflow)))
	return
}

func (dataflowshape *DataFlowShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(dataflowshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(dataflowshape), uint64(stage.GetOrder(dataflowshape)))
	return
}

func (datashape *DataShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datashape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(datashape), uint64(stage.GetOrder(datashape)))
	return
}

func (diagramprocess *DiagramProcess) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramprocess).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagramprocess), uint64(stage.GetOrder(diagramprocess)))
	return
}

func (externalparticipantshape *ExternalParticipantShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(externalparticipantshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(externalparticipantshape), uint64(stage.GetOrder(externalparticipantshape)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(library), uint64(stage.GetOrder(library)))
	return
}

func (note *Note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(note), uint64(stage.GetOrder(note)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(noteshape), uint64(stage.GetOrder(noteshape)))
	return
}

func (notetaskshape *NoteTaskShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notetaskshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notetaskshape), uint64(stage.GetOrder(notetaskshape)))
	return
}

func (participant *Participant) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(participant).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(participant), uint64(stage.GetOrder(participant)))
	return
}

func (participantshape *ParticipantShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(participantshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(participantshape), uint64(stage.GetOrder(participantshape)))
	return
}

func (process *Process) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(process).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(process), uint64(stage.GetOrder(process)))
	return
}

func (processshape *ProcessShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(processshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(processshape), uint64(stage.GetOrder(processshape)))
	return
}

func (resource *Resource) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(resource).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(resource), uint64(stage.GetOrder(resource)))
	return
}

func (task *Task) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(task).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(task), uint64(stage.GetOrder(task)))
	return
}

func (taskshape *TaskShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(taskshape), uint64(stage.GetOrder(taskshape)))
	return
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
	stage.AllocatedProcessShapes_reference = make(map[*AllocatedProcessShape]*AllocatedProcessShape)
	stage.AllocatedProcessShapes_referenceOrder = make(map[*AllocatedProcessShape]uint) // diff Unstage needs the reference order
	stage.AllocatedProcessShapes_instance = make(map[*AllocatedProcessShape]*AllocatedProcessShape)
	for instance := range stage.AllocatedProcessShapes {
		_copy := instance.GongCopy().(*AllocatedProcessShape)
		stage.AllocatedProcessShapes_reference[instance] = _copy
		stage.AllocatedProcessShapes_instance[_copy] = instance
		stage.AllocatedProcessShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.AllocatedResourceShapes_reference = make(map[*AllocatedResourceShape]*AllocatedResourceShape)
	stage.AllocatedResourceShapes_referenceOrder = make(map[*AllocatedResourceShape]uint) // diff Unstage needs the reference order
	stage.AllocatedResourceShapes_instance = make(map[*AllocatedResourceShape]*AllocatedResourceShape)
	for instance := range stage.AllocatedResourceShapes {
		_copy := instance.GongCopy().(*AllocatedResourceShape)
		stage.AllocatedResourceShapes_reference[instance] = _copy
		stage.AllocatedResourceShapes_instance[_copy] = instance
		stage.AllocatedResourceShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ControlFlows_reference = make(map[*ControlFlow]*ControlFlow)
	stage.ControlFlows_referenceOrder = make(map[*ControlFlow]uint) // diff Unstage needs the reference order
	stage.ControlFlows_instance = make(map[*ControlFlow]*ControlFlow)
	for instance := range stage.ControlFlows {
		_copy := instance.GongCopy().(*ControlFlow)
		stage.ControlFlows_reference[instance] = _copy
		stage.ControlFlows_instance[_copy] = instance
		stage.ControlFlows_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ControlFlowShapes_reference = make(map[*ControlFlowShape]*ControlFlowShape)
	stage.ControlFlowShapes_referenceOrder = make(map[*ControlFlowShape]uint) // diff Unstage needs the reference order
	stage.ControlFlowShapes_instance = make(map[*ControlFlowShape]*ControlFlowShape)
	for instance := range stage.ControlFlowShapes {
		_copy := instance.GongCopy().(*ControlFlowShape)
		stage.ControlFlowShapes_reference[instance] = _copy
		stage.ControlFlowShapes_instance[_copy] = instance
		stage.ControlFlowShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Datas_reference = make(map[*Data]*Data)
	stage.Datas_referenceOrder = make(map[*Data]uint) // diff Unstage needs the reference order
	stage.Datas_instance = make(map[*Data]*Data)
	for instance := range stage.Datas {
		_copy := instance.GongCopy().(*Data)
		stage.Datas_reference[instance] = _copy
		stage.Datas_instance[_copy] = instance
		stage.Datas_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DataFlows_reference = make(map[*DataFlow]*DataFlow)
	stage.DataFlows_referenceOrder = make(map[*DataFlow]uint) // diff Unstage needs the reference order
	stage.DataFlows_instance = make(map[*DataFlow]*DataFlow)
	for instance := range stage.DataFlows {
		_copy := instance.GongCopy().(*DataFlow)
		stage.DataFlows_reference[instance] = _copy
		stage.DataFlows_instance[_copy] = instance
		stage.DataFlows_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DataFlowShapes_reference = make(map[*DataFlowShape]*DataFlowShape)
	stage.DataFlowShapes_referenceOrder = make(map[*DataFlowShape]uint) // diff Unstage needs the reference order
	stage.DataFlowShapes_instance = make(map[*DataFlowShape]*DataFlowShape)
	for instance := range stage.DataFlowShapes {
		_copy := instance.GongCopy().(*DataFlowShape)
		stage.DataFlowShapes_reference[instance] = _copy
		stage.DataFlowShapes_instance[_copy] = instance
		stage.DataFlowShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DataShapes_reference = make(map[*DataShape]*DataShape)
	stage.DataShapes_referenceOrder = make(map[*DataShape]uint) // diff Unstage needs the reference order
	stage.DataShapes_instance = make(map[*DataShape]*DataShape)
	for instance := range stage.DataShapes {
		_copy := instance.GongCopy().(*DataShape)
		stage.DataShapes_reference[instance] = _copy
		stage.DataShapes_instance[_copy] = instance
		stage.DataShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DiagramProcesss_reference = make(map[*DiagramProcess]*DiagramProcess)
	stage.DiagramProcesss_referenceOrder = make(map[*DiagramProcess]uint) // diff Unstage needs the reference order
	stage.DiagramProcesss_instance = make(map[*DiagramProcess]*DiagramProcess)
	for instance := range stage.DiagramProcesss {
		_copy := instance.GongCopy().(*DiagramProcess)
		stage.DiagramProcesss_reference[instance] = _copy
		stage.DiagramProcesss_instance[_copy] = instance
		stage.DiagramProcesss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ExternalParticipantShapes_reference = make(map[*ExternalParticipantShape]*ExternalParticipantShape)
	stage.ExternalParticipantShapes_referenceOrder = make(map[*ExternalParticipantShape]uint) // diff Unstage needs the reference order
	stage.ExternalParticipantShapes_instance = make(map[*ExternalParticipantShape]*ExternalParticipantShape)
	for instance := range stage.ExternalParticipantShapes {
		_copy := instance.GongCopy().(*ExternalParticipantShape)
		stage.ExternalParticipantShapes_reference[instance] = _copy
		stage.ExternalParticipantShapes_instance[_copy] = instance
		stage.ExternalParticipantShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint) // diff Unstage needs the reference order
	stage.Librarys_instance = make(map[*Library]*Library)
	for instance := range stage.Librarys {
		_copy := instance.GongCopy().(*Library)
		stage.Librarys_reference[instance] = _copy
		stage.Librarys_instance[_copy] = instance
		stage.Librarys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint) // diff Unstage needs the reference order
	stage.Notes_instance = make(map[*Note]*Note)
	for instance := range stage.Notes {
		_copy := instance.GongCopy().(*Note)
		stage.Notes_reference[instance] = _copy
		stage.Notes_instance[_copy] = instance
		stage.Notes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteShapes_reference = make(map[*NoteShape]*NoteShape)
	stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint) // diff Unstage needs the reference order
	stage.NoteShapes_instance = make(map[*NoteShape]*NoteShape)
	for instance := range stage.NoteShapes {
		_copy := instance.GongCopy().(*NoteShape)
		stage.NoteShapes_reference[instance] = _copy
		stage.NoteShapes_instance[_copy] = instance
		stage.NoteShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.NoteTaskShapes_reference = make(map[*NoteTaskShape]*NoteTaskShape)
	stage.NoteTaskShapes_referenceOrder = make(map[*NoteTaskShape]uint) // diff Unstage needs the reference order
	stage.NoteTaskShapes_instance = make(map[*NoteTaskShape]*NoteTaskShape)
	for instance := range stage.NoteTaskShapes {
		_copy := instance.GongCopy().(*NoteTaskShape)
		stage.NoteTaskShapes_reference[instance] = _copy
		stage.NoteTaskShapes_instance[_copy] = instance
		stage.NoteTaskShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Participants_reference = make(map[*Participant]*Participant)
	stage.Participants_referenceOrder = make(map[*Participant]uint) // diff Unstage needs the reference order
	stage.Participants_instance = make(map[*Participant]*Participant)
	for instance := range stage.Participants {
		_copy := instance.GongCopy().(*Participant)
		stage.Participants_reference[instance] = _copy
		stage.Participants_instance[_copy] = instance
		stage.Participants_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ParticipantShapes_reference = make(map[*ParticipantShape]*ParticipantShape)
	stage.ParticipantShapes_referenceOrder = make(map[*ParticipantShape]uint) // diff Unstage needs the reference order
	stage.ParticipantShapes_instance = make(map[*ParticipantShape]*ParticipantShape)
	for instance := range stage.ParticipantShapes {
		_copy := instance.GongCopy().(*ParticipantShape)
		stage.ParticipantShapes_reference[instance] = _copy
		stage.ParticipantShapes_instance[_copy] = instance
		stage.ParticipantShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Processs_reference = make(map[*Process]*Process)
	stage.Processs_referenceOrder = make(map[*Process]uint) // diff Unstage needs the reference order
	stage.Processs_instance = make(map[*Process]*Process)
	for instance := range stage.Processs {
		_copy := instance.GongCopy().(*Process)
		stage.Processs_reference[instance] = _copy
		stage.Processs_instance[_copy] = instance
		stage.Processs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ProcessShapes_reference = make(map[*ProcessShape]*ProcessShape)
	stage.ProcessShapes_referenceOrder = make(map[*ProcessShape]uint) // diff Unstage needs the reference order
	stage.ProcessShapes_instance = make(map[*ProcessShape]*ProcessShape)
	for instance := range stage.ProcessShapes {
		_copy := instance.GongCopy().(*ProcessShape)
		stage.ProcessShapes_reference[instance] = _copy
		stage.ProcessShapes_instance[_copy] = instance
		stage.ProcessShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Resources_reference = make(map[*Resource]*Resource)
	stage.Resources_referenceOrder = make(map[*Resource]uint) // diff Unstage needs the reference order
	stage.Resources_instance = make(map[*Resource]*Resource)
	for instance := range stage.Resources {
		_copy := instance.GongCopy().(*Resource)
		stage.Resources_reference[instance] = _copy
		stage.Resources_instance[_copy] = instance
		stage.Resources_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tasks_reference = make(map[*Task]*Task)
	stage.Tasks_referenceOrder = make(map[*Task]uint) // diff Unstage needs the reference order
	stage.Tasks_instance = make(map[*Task]*Task)
	for instance := range stage.Tasks {
		_copy := instance.GongCopy().(*Task)
		stage.Tasks_reference[instance] = _copy
		stage.Tasks_instance[_copy] = instance
		stage.Tasks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.TaskShapes_reference = make(map[*TaskShape]*TaskShape)
	stage.TaskShapes_referenceOrder = make(map[*TaskShape]uint) // diff Unstage needs the reference order
	stage.TaskShapes_instance = make(map[*TaskShape]*TaskShape)
	for instance := range stage.TaskShapes {
		_copy := instance.GongCopy().(*TaskShape)
		stage.TaskShapes_reference[instance] = _copy
		stage.TaskShapes_instance[_copy] = instance
		stage.TaskShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.AllocatedProcessShapes {
		reference := stage.AllocatedProcessShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.AllocatedResourceShapes {
		reference := stage.AllocatedResourceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ControlFlows {
		reference := stage.ControlFlows_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ControlFlowShapes {
		reference := stage.ControlFlowShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Datas {
		reference := stage.Datas_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DataFlows {
		reference := stage.DataFlows_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DataFlowShapes {
		reference := stage.DataFlowShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DataShapes {
		reference := stage.DataShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DiagramProcesss {
		reference := stage.DiagramProcesss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ExternalParticipantShapes {
		reference := stage.ExternalParticipantShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Notes {
		reference := stage.Notes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteShapes {
		reference := stage.NoteShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteTaskShapes {
		reference := stage.NoteTaskShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Participants {
		reference := stage.Participants_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ParticipantShapes {
		reference := stage.ParticipantShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Processs {
		reference := stage.Processs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ProcessShapes {
		reference := stage.ProcessShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Resources {
		reference := stage.Resources_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tasks {
		reference := stage.Tasks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.TaskShapes {
		reference := stage.TaskShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

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
	if order, ok := stage.AllocatedProcessShape_stagedOrder[allocatedprocessshape]; ok {
		return order
	}
	if order, ok := stage.AllocatedProcessShapes_referenceOrder[allocatedprocessshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type AllocatedProcessShape was not staged and does not have a reference order", allocatedprocessshape)
		return 0
	}
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AllocatedResourceShape_stagedOrder[allocatedresourceshape]; ok {
		return order
	}
	if order, ok := stage.AllocatedResourceShapes_referenceOrder[allocatedresourceshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type AllocatedResourceShape was not staged and does not have a reference order", allocatedresourceshape)
		return 0
	}
}

func (controlflow *ControlFlow) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlFlow_stagedOrder[controlflow]; ok {
		return order
	}
	if order, ok := stage.ControlFlows_referenceOrder[controlflow]; ok {
		return order
	} else {
		log.Printf("instance %p of type ControlFlow was not staged and does not have a reference order", controlflow)
		return 0
	}
}

func (controlflowshape *ControlFlowShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlFlowShape_stagedOrder[controlflowshape]; ok {
		return order
	}
	if order, ok := stage.ControlFlowShapes_referenceOrder[controlflowshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ControlFlowShape was not staged and does not have a reference order", controlflowshape)
		return 0
	}
}

func (data *Data) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Data_stagedOrder[data]; ok {
		return order
	}
	if order, ok := stage.Datas_referenceOrder[data]; ok {
		return order
	} else {
		log.Printf("instance %p of type Data was not staged and does not have a reference order", data)
		return 0
	}
}

func (dataflow *DataFlow) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DataFlow_stagedOrder[dataflow]; ok {
		return order
	}
	if order, ok := stage.DataFlows_referenceOrder[dataflow]; ok {
		return order
	} else {
		log.Printf("instance %p of type DataFlow was not staged and does not have a reference order", dataflow)
		return 0
	}
}

func (dataflowshape *DataFlowShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DataFlowShape_stagedOrder[dataflowshape]; ok {
		return order
	}
	if order, ok := stage.DataFlowShapes_referenceOrder[dataflowshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type DataFlowShape was not staged and does not have a reference order", dataflowshape)
		return 0
	}
}

func (datashape *DataShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DataShape_stagedOrder[datashape]; ok {
		return order
	}
	if order, ok := stage.DataShapes_referenceOrder[datashape]; ok {
		return order
	} else {
		log.Printf("instance %p of type DataShape was not staged and does not have a reference order", datashape)
		return 0
	}
}

func (diagramprocess *DiagramProcess) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramProcess_stagedOrder[diagramprocess]; ok {
		return order
	}
	if order, ok := stage.DiagramProcesss_referenceOrder[diagramprocess]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramProcess was not staged and does not have a reference order", diagramprocess)
		return 0
	}
}

func (externalparticipantshape *ExternalParticipantShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ExternalParticipantShape_stagedOrder[externalparticipantshape]; ok {
		return order
	}
	if order, ok := stage.ExternalParticipantShapes_referenceOrder[externalparticipantshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ExternalParticipantShape was not staged and does not have a reference order", externalparticipantshape)
		return 0
	}
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Library_stagedOrder[library]; ok {
		return order
	}
	if order, ok := stage.Librarys_referenceOrder[library]; ok {
		return order
	} else {
		log.Printf("instance %p of type Library was not staged and does not have a reference order", library)
		return 0
	}
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Note_stagedOrder[note]; ok {
		return order
	}
	if order, ok := stage.Notes_referenceOrder[note]; ok {
		return order
	} else {
		log.Printf("instance %p of type Note was not staged and does not have a reference order", note)
		return 0
	}
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteShape_stagedOrder[noteshape]; ok {
		return order
	}
	if order, ok := stage.NoteShapes_referenceOrder[noteshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteShape was not staged and does not have a reference order", noteshape)
		return 0
	}
}

func (notetaskshape *NoteTaskShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NoteTaskShape_stagedOrder[notetaskshape]; ok {
		return order
	}
	if order, ok := stage.NoteTaskShapes_referenceOrder[notetaskshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NoteTaskShape was not staged and does not have a reference order", notetaskshape)
		return 0
	}
}

func (participant *Participant) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Participant_stagedOrder[participant]; ok {
		return order
	}
	if order, ok := stage.Participants_referenceOrder[participant]; ok {
		return order
	} else {
		log.Printf("instance %p of type Participant was not staged and does not have a reference order", participant)
		return 0
	}
}

func (participantshape *ParticipantShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ParticipantShape_stagedOrder[participantshape]; ok {
		return order
	}
	if order, ok := stage.ParticipantShapes_referenceOrder[participantshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ParticipantShape was not staged and does not have a reference order", participantshape)
		return 0
	}
}

func (process *Process) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Process_stagedOrder[process]; ok {
		return order
	}
	if order, ok := stage.Processs_referenceOrder[process]; ok {
		return order
	} else {
		log.Printf("instance %p of type Process was not staged and does not have a reference order", process)
		return 0
	}
}

func (processshape *ProcessShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ProcessShape_stagedOrder[processshape]; ok {
		return order
	}
	if order, ok := stage.ProcessShapes_referenceOrder[processshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ProcessShape was not staged and does not have a reference order", processshape)
		return 0
	}
}

func (resource *Resource) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Resource_stagedOrder[resource]; ok {
		return order
	}
	if order, ok := stage.Resources_referenceOrder[resource]; ok {
		return order
	} else {
		log.Printf("instance %p of type Resource was not staged and does not have a reference order", resource)
		return 0
	}
}

func (task *Task) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Task_stagedOrder[task]; ok {
		return order
	}
	if order, ok := stage.Tasks_referenceOrder[task]; ok {
		return order
	} else {
		log.Printf("instance %p of type Task was not staged and does not have a reference order", task)
		return 0
	}
}

func (taskshape *TaskShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TaskShape_stagedOrder[taskshape]; ok {
		return order
	}
	if order, ok := stage.TaskShapes_referenceOrder[taskshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type TaskShape was not staged and does not have a reference order", taskshape)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (allocatedprocessshape *AllocatedProcessShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", allocatedprocessshape.GongGetGongstructName(), allocatedprocessshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (allocatedprocessshape *AllocatedProcessShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", allocatedprocessshape.GongGetGongstructName(), allocatedprocessshape.GongGetOrder(stage))
}

func (allocatedresourceshape *AllocatedResourceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", allocatedresourceshape.GongGetGongstructName(), allocatedresourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (allocatedresourceshape *AllocatedResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", allocatedresourceshape.GongGetGongstructName(), allocatedresourceshape.GongGetOrder(stage))
}

func (controlflow *ControlFlow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlflow.GongGetGongstructName(), controlflow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlflow *ControlFlow) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlflow.GongGetGongstructName(), controlflow.GongGetOrder(stage))
}

func (controlflowshape *ControlFlowShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlflowshape.GongGetGongstructName(), controlflowshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlflowshape *ControlFlowShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlflowshape.GongGetGongstructName(), controlflowshape.GongGetOrder(stage))
}

func (data *Data) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", data.GongGetGongstructName(), data.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (data *Data) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", data.GongGetGongstructName(), data.GongGetOrder(stage))
}

func (dataflow *DataFlow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dataflow.GongGetGongstructName(), dataflow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dataflow *DataFlow) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dataflow.GongGetGongstructName(), dataflow.GongGetOrder(stage))
}

func (dataflowshape *DataFlowShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dataflowshape.GongGetGongstructName(), dataflowshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dataflowshape *DataFlowShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dataflowshape.GongGetGongstructName(), dataflowshape.GongGetOrder(stage))
}

func (datashape *DataShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datashape.GongGetGongstructName(), datashape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datashape *DataShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datashape.GongGetGongstructName(), datashape.GongGetOrder(stage))
}

func (diagramprocess *DiagramProcess) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramprocess.GongGetGongstructName(), diagramprocess.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramprocess *DiagramProcess) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramprocess.GongGetGongstructName(), diagramprocess.GongGetOrder(stage))
}

func (externalparticipantshape *ExternalParticipantShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", externalparticipantshape.GongGetGongstructName(), externalparticipantshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (externalparticipantshape *ExternalParticipantShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", externalparticipantshape.GongGetGongstructName(), externalparticipantshape.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

func (notetaskshape *NoteTaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notetaskshape *NoteTaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notetaskshape.GongGetGongstructName(), notetaskshape.GongGetOrder(stage))
}

func (participant *Participant) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", participant.GongGetGongstructName(), participant.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (participant *Participant) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", participant.GongGetGongstructName(), participant.GongGetOrder(stage))
}

func (participantshape *ParticipantShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", participantshape.GongGetGongstructName(), participantshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (participantshape *ParticipantShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", participantshape.GongGetGongstructName(), participantshape.GongGetOrder(stage))
}

func (process *Process) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", process.GongGetGongstructName(), process.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (process *Process) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", process.GongGetGongstructName(), process.GongGetOrder(stage))
}

func (processshape *ProcessShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", processshape.GongGetGongstructName(), processshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (processshape *ProcessShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", processshape.GongGetGongstructName(), processshape.GongGetOrder(stage))
}

func (resource *Resource) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resource.GongGetGongstructName(), resource.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resource *Resource) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resource.GongGetGongstructName(), resource.GongGetOrder(stage))
}

func (task *Task) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", task.GongGetGongstructName(), task.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (task *Task) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", task.GongGetGongstructName(), task.GongGetOrder(stage))
}

func (taskshape *TaskShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskshape.GongGetGongstructName(), taskshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (taskshape *TaskShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", taskshape.GongGetGongstructName(), taskshape.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (allocatedprocessshape *AllocatedProcessShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedprocessshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AllocatedProcessShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(allocatedprocessshape.Name))
	return
}

func (allocatedresourceshape *AllocatedResourceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedresourceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AllocatedResourceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(allocatedresourceshape.Name))
	return
}

func (controlflow *ControlFlow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlFlow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(controlflow.Name))
	return
}

func (controlflowshape *ControlFlowShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlFlowShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(controlflowshape.Name))
	return
}

func (data *Data) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", data.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Data")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(data.Name))
	return
}

func (dataflow *DataFlow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DataFlow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(dataflow.Name))
	return
}

func (dataflowshape *DataFlowShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DataFlowShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(dataflowshape.Name))
	return
}

func (datashape *DataShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datashape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DataShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(datashape.Name))
	return
}

func (diagramprocess *DiagramProcess) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramProcess")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagramprocess.Name))
	return
}

func (externalparticipantshape *ExternalParticipantShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", externalparticipantshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExternalParticipantShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(externalparticipantshape.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(library.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(note.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(noteshape.Name))
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteTaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notetaskshape.Name))
	return
}

func (participant *Participant) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", participant.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Participant")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(participant.Name))
	return
}

func (participantshape *ParticipantShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", participantshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParticipantShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(participantshape.Name))
	return
}

func (process *Process) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", process.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Process")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(process.Name))
	return
}

func (processshape *ProcessShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", processshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProcessShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(processshape.Name))
	return
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Resource")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(resource.Name))
	return
}

func (task *Task) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", task.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Task")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(task.Name))
	return
}

func (taskshape *TaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(taskshape.Name))
	return
}

// insertion point for unstaging
func (allocatedprocessshape *AllocatedProcessShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedprocessshape.GongGetReferenceIdentifier(stage))
	return
}

func (allocatedresourceshape *AllocatedResourceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedresourceshape.GongGetReferenceIdentifier(stage))
	return
}

func (controlflow *ControlFlow) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlflow.GongGetReferenceIdentifier(stage))
	return
}

func (controlflowshape *ControlFlowShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlflowshape.GongGetReferenceIdentifier(stage))
	return
}

func (data *Data) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", data.GongGetReferenceIdentifier(stage))
	return
}

func (dataflow *DataFlow) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dataflow.GongGetReferenceIdentifier(stage))
	return
}

func (dataflowshape *DataFlowShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dataflowshape.GongGetReferenceIdentifier(stage))
	return
}

func (datashape *DataShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datashape.GongGetReferenceIdentifier(stage))
	return
}

func (diagramprocess *DiagramProcess) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramprocess.GongGetReferenceIdentifier(stage))
	return
}

func (externalparticipantshape *ExternalParticipantShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", externalparticipantshape.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (note *Note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetReferenceIdentifier(stage))
	return
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetReferenceIdentifier(stage))
	return
}

func (notetaskshape *NoteTaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notetaskshape.GongGetReferenceIdentifier(stage))
	return
}

func (participant *Participant) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", participant.GongGetReferenceIdentifier(stage))
	return
}

func (participantshape *ParticipantShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", participantshape.GongGetReferenceIdentifier(stage))
	return
}

func (process *Process) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", process.GongGetReferenceIdentifier(stage))
	return
}

func (processshape *ProcessShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", processshape.GongGetReferenceIdentifier(stage))
	return
}

func (resource *Resource) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetReferenceIdentifier(stage))
	return
}

func (task *Task) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", task.GongGetReferenceIdentifier(stage))
	return
}

func (taskshape *TaskShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetReferenceIdentifier(stage))
	return
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

// end of template
