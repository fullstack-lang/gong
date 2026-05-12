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

	// Compute reverse map for named struct Participant
	// insertion point per field
	stage.Participant_Resources_reverseMap = make(map[*Resource]*Participant)
	for participant := range stage.Participants {
		_ = participant
		for _, _resource := range participant.Resources {
			stage.Participant_Resources_reverseMap[_resource] = participant
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
func (allocatedresourceshape *AllocatedResourceShape) GongCopy() GongstructIF {
	newInstance := new(AllocatedResourceShape)
	allocatedresourceshape.CopyBasicFields(newInstance)
	return newInstance
}

func (controlflow *ControlFlow) GongCopy() GongstructIF {
	newInstance := new(ControlFlow)
	controlflow.CopyBasicFields(newInstance)
	return newInstance
}

func (controlflowshape *ControlFlowShape) GongCopy() GongstructIF {
	newInstance := new(ControlFlowShape)
	controlflowshape.CopyBasicFields(newInstance)
	return newInstance
}

func (data *Data) GongCopy() GongstructIF {
	newInstance := new(Data)
	data.CopyBasicFields(newInstance)
	return newInstance
}

func (dataflow *DataFlow) GongCopy() GongstructIF {
	newInstance := new(DataFlow)
	dataflow.CopyBasicFields(newInstance)
	return newInstance
}

func (dataflowshape *DataFlowShape) GongCopy() GongstructIF {
	newInstance := new(DataFlowShape)
	dataflowshape.CopyBasicFields(newInstance)
	return newInstance
}

func (datashape *DataShape) GongCopy() GongstructIF {
	newInstance := new(DataShape)
	datashape.CopyBasicFields(newInstance)
	return newInstance
}

func (diagramprocess *DiagramProcess) GongCopy() GongstructIF {
	newInstance := new(DiagramProcess)
	diagramprocess.CopyBasicFields(newInstance)
	return newInstance
}

func (externalparticipantshape *ExternalParticipantShape) GongCopy() GongstructIF {
	newInstance := new(ExternalParticipantShape)
	externalparticipantshape.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.CopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.CopyBasicFields(newInstance)
	return newInstance
}

func (participant *Participant) GongCopy() GongstructIF {
	newInstance := new(Participant)
	participant.CopyBasicFields(newInstance)
	return newInstance
}

func (participantshape *ParticipantShape) GongCopy() GongstructIF {
	newInstance := new(ParticipantShape)
	participantshape.CopyBasicFields(newInstance)
	return newInstance
}

func (process *Process) GongCopy() GongstructIF {
	newInstance := new(Process)
	process.CopyBasicFields(newInstance)
	return newInstance
}

func (processshape *ProcessShape) GongCopy() GongstructIF {
	newInstance := new(ProcessShape)
	processshape.CopyBasicFields(newInstance)
	return newInstance
}

func (resource *Resource) GongCopy() GongstructIF {
	newInstance := new(Resource)
	resource.CopyBasicFields(newInstance)
	return newInstance
}

func (task *Task) GongCopy() GongstructIF {
	newInstance := new(Task)
	task.CopyBasicFields(newInstance)
	return newInstance
}

func (taskshape *TaskShape) GongCopy() GongstructIF {
	newInstance := new(TaskShape)
	taskshape.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (allocatedresourceshape *AllocatedResourceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(allocatedresourceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(allocatedresourceshape), uint64(GetOrderPointerGongstruct(stage, allocatedresourceshape)))
	return
}

func (controlflow *ControlFlow) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlflow).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(controlflow), uint64(GetOrderPointerGongstruct(stage, controlflow)))
	return
}

func (controlflowshape *ControlFlowShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(controlflowshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(controlflowshape), uint64(GetOrderPointerGongstruct(stage, controlflowshape)))
	return
}

func (data *Data) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(data).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(data), uint64(GetOrderPointerGongstruct(stage, data)))
	return
}

func (dataflow *DataFlow) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(dataflow).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(dataflow), uint64(GetOrderPointerGongstruct(stage, dataflow)))
	return
}

func (dataflowshape *DataFlowShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(dataflowshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(dataflowshape), uint64(GetOrderPointerGongstruct(stage, dataflowshape)))
	return
}

func (datashape *DataShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datashape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(datashape), uint64(GetOrderPointerGongstruct(stage, datashape)))
	return
}

func (diagramprocess *DiagramProcess) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramprocess).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagramprocess), uint64(GetOrderPointerGongstruct(stage, diagramprocess)))
	return
}

func (externalparticipantshape *ExternalParticipantShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(externalparticipantshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(externalparticipantshape), uint64(GetOrderPointerGongstruct(stage, externalparticipantshape)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
	return
}

func (note *Note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(note), uint64(GetOrderPointerGongstruct(stage, note)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(noteshape), uint64(GetOrderPointerGongstruct(stage, noteshape)))
	return
}

func (participant *Participant) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(participant).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(participant), uint64(GetOrderPointerGongstruct(stage, participant)))
	return
}

func (participantshape *ParticipantShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(participantshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(participantshape), uint64(GetOrderPointerGongstruct(stage, participantshape)))
	return
}

func (process *Process) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(process).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(process), uint64(GetOrderPointerGongstruct(stage, process)))
	return
}

func (processshape *ProcessShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(processshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(processshape), uint64(GetOrderPointerGongstruct(stage, processshape)))
	return
}

func (resource *Resource) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(resource).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(resource), uint64(GetOrderPointerGongstruct(stage, resource)))
	return
}

func (task *Task) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(task).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(task), uint64(GetOrderPointerGongstruct(stage, task)))
	return
}

func (taskshape *TaskShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(taskshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(taskshape), uint64(GetOrderPointerGongstruct(stage, taskshape)))
	return
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
	var allocatedresourceshapes_newInstances []*AllocatedResourceShape
	var allocatedresourceshapes_deletedInstances []*AllocatedResourceShape

	// parse all staged instances and check if they have a reference
	for allocatedresourceshape := range stage.AllocatedResourceShapes {
		if ref, ok := stage.AllocatedResourceShapes_reference[allocatedresourceshape]; !ok {
			allocatedresourceshapes_newInstances = append(allocatedresourceshapes_newInstances, allocatedresourceshape)
			newInstancesSlice = append(newInstancesSlice, allocatedresourceshape.GongMarshallIdentifier(stage))
			if stage.AllocatedResourceShapes_referenceOrder == nil {
				stage.AllocatedResourceShapes_referenceOrder = make(map[*AllocatedResourceShape]uint)
			}
			stage.AllocatedResourceShapes_referenceOrder[allocatedresourceshape] = stage.AllocatedResourceShape_stagedOrder[allocatedresourceshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, allocatedresourceshape.GongMarshallUnstaging(stage))
			// delete(stage.AllocatedResourceShapes_referenceOrder, allocatedresourceshape)
			fieldInitializers, pointersInitializations := allocatedresourceshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AllocatedResourceShape_stagedOrder[ref] = stage.AllocatedResourceShape_stagedOrder[allocatedresourceshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := allocatedresourceshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, allocatedresourceshape)
			// delete(stage.AllocatedResourceShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if allocatedresourceshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", allocatedresourceshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.AllocatedResourceShapes_reference {
		instance := stage.AllocatedResourceShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.AllocatedResourceShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			allocatedresourceshapes_deletedInstances = append(allocatedresourceshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(allocatedresourceshapes_newInstances)
	lenDeletedInstances += len(allocatedresourceshapes_deletedInstances)
	var controlflows_newInstances []*ControlFlow
	var controlflows_deletedInstances []*ControlFlow

	// parse all staged instances and check if they have a reference
	for controlflow := range stage.ControlFlows {
		if ref, ok := stage.ControlFlows_reference[controlflow]; !ok {
			controlflows_newInstances = append(controlflows_newInstances, controlflow)
			newInstancesSlice = append(newInstancesSlice, controlflow.GongMarshallIdentifier(stage))
			if stage.ControlFlows_referenceOrder == nil {
				stage.ControlFlows_referenceOrder = make(map[*ControlFlow]uint)
			}
			stage.ControlFlows_referenceOrder[controlflow] = stage.ControlFlow_stagedOrder[controlflow]
			newInstancesReverseSlice = append(newInstancesReverseSlice, controlflow.GongMarshallUnstaging(stage))
			// delete(stage.ControlFlows_referenceOrder, controlflow)
			fieldInitializers, pointersInitializations := controlflow.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ControlFlow_stagedOrder[ref] = stage.ControlFlow_stagedOrder[controlflow]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := controlflow.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, controlflow)
			// delete(stage.ControlFlow_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if controlflow.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", controlflow.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ControlFlows_reference {
		instance := stage.ControlFlows_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ControlFlows[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			controlflows_deletedInstances = append(controlflows_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(controlflows_newInstances)
	lenDeletedInstances += len(controlflows_deletedInstances)
	var controlflowshapes_newInstances []*ControlFlowShape
	var controlflowshapes_deletedInstances []*ControlFlowShape

	// parse all staged instances and check if they have a reference
	for controlflowshape := range stage.ControlFlowShapes {
		if ref, ok := stage.ControlFlowShapes_reference[controlflowshape]; !ok {
			controlflowshapes_newInstances = append(controlflowshapes_newInstances, controlflowshape)
			newInstancesSlice = append(newInstancesSlice, controlflowshape.GongMarshallIdentifier(stage))
			if stage.ControlFlowShapes_referenceOrder == nil {
				stage.ControlFlowShapes_referenceOrder = make(map[*ControlFlowShape]uint)
			}
			stage.ControlFlowShapes_referenceOrder[controlflowshape] = stage.ControlFlowShape_stagedOrder[controlflowshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, controlflowshape.GongMarshallUnstaging(stage))
			// delete(stage.ControlFlowShapes_referenceOrder, controlflowshape)
			fieldInitializers, pointersInitializations := controlflowshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ControlFlowShape_stagedOrder[ref] = stage.ControlFlowShape_stagedOrder[controlflowshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := controlflowshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, controlflowshape)
			// delete(stage.ControlFlowShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if controlflowshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", controlflowshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ControlFlowShapes_reference {
		instance := stage.ControlFlowShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ControlFlowShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			controlflowshapes_deletedInstances = append(controlflowshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(controlflowshapes_newInstances)
	lenDeletedInstances += len(controlflowshapes_deletedInstances)
	var datas_newInstances []*Data
	var datas_deletedInstances []*Data

	// parse all staged instances and check if they have a reference
	for data := range stage.Datas {
		if ref, ok := stage.Datas_reference[data]; !ok {
			datas_newInstances = append(datas_newInstances, data)
			newInstancesSlice = append(newInstancesSlice, data.GongMarshallIdentifier(stage))
			if stage.Datas_referenceOrder == nil {
				stage.Datas_referenceOrder = make(map[*Data]uint)
			}
			stage.Datas_referenceOrder[data] = stage.Data_stagedOrder[data]
			newInstancesReverseSlice = append(newInstancesReverseSlice, data.GongMarshallUnstaging(stage))
			// delete(stage.Datas_referenceOrder, data)
			fieldInitializers, pointersInitializations := data.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Data_stagedOrder[ref] = stage.Data_stagedOrder[data]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := data.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, data)
			// delete(stage.Data_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if data.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", data.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Datas_reference {
		instance := stage.Datas_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Datas[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			datas_deletedInstances = append(datas_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datas_newInstances)
	lenDeletedInstances += len(datas_deletedInstances)
	var dataflows_newInstances []*DataFlow
	var dataflows_deletedInstances []*DataFlow

	// parse all staged instances and check if they have a reference
	for dataflow := range stage.DataFlows {
		if ref, ok := stage.DataFlows_reference[dataflow]; !ok {
			dataflows_newInstances = append(dataflows_newInstances, dataflow)
			newInstancesSlice = append(newInstancesSlice, dataflow.GongMarshallIdentifier(stage))
			if stage.DataFlows_referenceOrder == nil {
				stage.DataFlows_referenceOrder = make(map[*DataFlow]uint)
			}
			stage.DataFlows_referenceOrder[dataflow] = stage.DataFlow_stagedOrder[dataflow]
			newInstancesReverseSlice = append(newInstancesReverseSlice, dataflow.GongMarshallUnstaging(stage))
			// delete(stage.DataFlows_referenceOrder, dataflow)
			fieldInitializers, pointersInitializations := dataflow.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DataFlow_stagedOrder[ref] = stage.DataFlow_stagedOrder[dataflow]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := dataflow.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, dataflow)
			// delete(stage.DataFlow_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if dataflow.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", dataflow.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DataFlows_reference {
		instance := stage.DataFlows_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DataFlows[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			dataflows_deletedInstances = append(dataflows_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(dataflows_newInstances)
	lenDeletedInstances += len(dataflows_deletedInstances)
	var dataflowshapes_newInstances []*DataFlowShape
	var dataflowshapes_deletedInstances []*DataFlowShape

	// parse all staged instances and check if they have a reference
	for dataflowshape := range stage.DataFlowShapes {
		if ref, ok := stage.DataFlowShapes_reference[dataflowshape]; !ok {
			dataflowshapes_newInstances = append(dataflowshapes_newInstances, dataflowshape)
			newInstancesSlice = append(newInstancesSlice, dataflowshape.GongMarshallIdentifier(stage))
			if stage.DataFlowShapes_referenceOrder == nil {
				stage.DataFlowShapes_referenceOrder = make(map[*DataFlowShape]uint)
			}
			stage.DataFlowShapes_referenceOrder[dataflowshape] = stage.DataFlowShape_stagedOrder[dataflowshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, dataflowshape.GongMarshallUnstaging(stage))
			// delete(stage.DataFlowShapes_referenceOrder, dataflowshape)
			fieldInitializers, pointersInitializations := dataflowshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DataFlowShape_stagedOrder[ref] = stage.DataFlowShape_stagedOrder[dataflowshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := dataflowshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, dataflowshape)
			// delete(stage.DataFlowShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if dataflowshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", dataflowshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DataFlowShapes_reference {
		instance := stage.DataFlowShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DataFlowShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			dataflowshapes_deletedInstances = append(dataflowshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(dataflowshapes_newInstances)
	lenDeletedInstances += len(dataflowshapes_deletedInstances)
	var datashapes_newInstances []*DataShape
	var datashapes_deletedInstances []*DataShape

	// parse all staged instances and check if they have a reference
	for datashape := range stage.DataShapes {
		if ref, ok := stage.DataShapes_reference[datashape]; !ok {
			datashapes_newInstances = append(datashapes_newInstances, datashape)
			newInstancesSlice = append(newInstancesSlice, datashape.GongMarshallIdentifier(stage))
			if stage.DataShapes_referenceOrder == nil {
				stage.DataShapes_referenceOrder = make(map[*DataShape]uint)
			}
			stage.DataShapes_referenceOrder[datashape] = stage.DataShape_stagedOrder[datashape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, datashape.GongMarshallUnstaging(stage))
			// delete(stage.DataShapes_referenceOrder, datashape)
			fieldInitializers, pointersInitializations := datashape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DataShape_stagedOrder[ref] = stage.DataShape_stagedOrder[datashape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := datashape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, datashape)
			// delete(stage.DataShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if datashape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", datashape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DataShapes_reference {
		instance := stage.DataShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DataShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			datashapes_deletedInstances = append(datashapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datashapes_newInstances)
	lenDeletedInstances += len(datashapes_deletedInstances)
	var diagramprocesss_newInstances []*DiagramProcess
	var diagramprocesss_deletedInstances []*DiagramProcess

	// parse all staged instances and check if they have a reference
	for diagramprocess := range stage.DiagramProcesss {
		if ref, ok := stage.DiagramProcesss_reference[diagramprocess]; !ok {
			diagramprocesss_newInstances = append(diagramprocesss_newInstances, diagramprocess)
			newInstancesSlice = append(newInstancesSlice, diagramprocess.GongMarshallIdentifier(stage))
			if stage.DiagramProcesss_referenceOrder == nil {
				stage.DiagramProcesss_referenceOrder = make(map[*DiagramProcess]uint)
			}
			stage.DiagramProcesss_referenceOrder[diagramprocess] = stage.DiagramProcess_stagedOrder[diagramprocess]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagramprocess.GongMarshallUnstaging(stage))
			// delete(stage.DiagramProcesss_referenceOrder, diagramprocess)
			fieldInitializers, pointersInitializations := diagramprocess.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramProcess_stagedOrder[ref] = stage.DiagramProcess_stagedOrder[diagramprocess]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagramprocess.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagramprocess)
			// delete(stage.DiagramProcess_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if diagramprocess.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", diagramprocess.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DiagramProcesss_reference {
		instance := stage.DiagramProcesss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DiagramProcesss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagramprocesss_deletedInstances = append(diagramprocesss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagramprocesss_newInstances)
	lenDeletedInstances += len(diagramprocesss_deletedInstances)
	var externalparticipantshapes_newInstances []*ExternalParticipantShape
	var externalparticipantshapes_deletedInstances []*ExternalParticipantShape

	// parse all staged instances and check if they have a reference
	for externalparticipantshape := range stage.ExternalParticipantShapes {
		if ref, ok := stage.ExternalParticipantShapes_reference[externalparticipantshape]; !ok {
			externalparticipantshapes_newInstances = append(externalparticipantshapes_newInstances, externalparticipantshape)
			newInstancesSlice = append(newInstancesSlice, externalparticipantshape.GongMarshallIdentifier(stage))
			if stage.ExternalParticipantShapes_referenceOrder == nil {
				stage.ExternalParticipantShapes_referenceOrder = make(map[*ExternalParticipantShape]uint)
			}
			stage.ExternalParticipantShapes_referenceOrder[externalparticipantshape] = stage.ExternalParticipantShape_stagedOrder[externalparticipantshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, externalparticipantshape.GongMarshallUnstaging(stage))
			// delete(stage.ExternalParticipantShapes_referenceOrder, externalparticipantshape)
			fieldInitializers, pointersInitializations := externalparticipantshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ExternalParticipantShape_stagedOrder[ref] = stage.ExternalParticipantShape_stagedOrder[externalparticipantshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := externalparticipantshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, externalparticipantshape)
			// delete(stage.ExternalParticipantShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if externalparticipantshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", externalparticipantshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ExternalParticipantShapes_reference {
		instance := stage.ExternalParticipantShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ExternalParticipantShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			externalparticipantshapes_deletedInstances = append(externalparticipantshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(externalparticipantshapes_newInstances)
	lenDeletedInstances += len(externalparticipantshapes_deletedInstances)
	var librarys_newInstances []*Library
	var librarys_deletedInstances []*Library

	// parse all staged instances and check if they have a reference
	for library := range stage.Librarys {
		if ref, ok := stage.Librarys_reference[library]; !ok {
			librarys_newInstances = append(librarys_newInstances, library)
			newInstancesSlice = append(newInstancesSlice, library.GongMarshallIdentifier(stage))
			if stage.Librarys_referenceOrder == nil {
				stage.Librarys_referenceOrder = make(map[*Library]uint)
			}
			stage.Librarys_referenceOrder[library] = stage.Library_stagedOrder[library]
			newInstancesReverseSlice = append(newInstancesReverseSlice, library.GongMarshallUnstaging(stage))
			// delete(stage.Librarys_referenceOrder, library)
			fieldInitializers, pointersInitializations := library.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Library_stagedOrder[ref] = stage.Library_stagedOrder[library]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := library.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, library)
			// delete(stage.Library_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if library.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", library.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Librarys_reference {
		instance := stage.Librarys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Librarys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			librarys_deletedInstances = append(librarys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(librarys_newInstances)
	lenDeletedInstances += len(librarys_deletedInstances)
	var notes_newInstances []*Note
	var notes_deletedInstances []*Note

	// parse all staged instances and check if they have a reference
	for note := range stage.Notes {
		if ref, ok := stage.Notes_reference[note]; !ok {
			notes_newInstances = append(notes_newInstances, note)
			newInstancesSlice = append(newInstancesSlice, note.GongMarshallIdentifier(stage))
			if stage.Notes_referenceOrder == nil {
				stage.Notes_referenceOrder = make(map[*Note]uint)
			}
			stage.Notes_referenceOrder[note] = stage.Note_stagedOrder[note]
			newInstancesReverseSlice = append(newInstancesReverseSlice, note.GongMarshallUnstaging(stage))
			// delete(stage.Notes_referenceOrder, note)
			fieldInitializers, pointersInitializations := note.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Note_stagedOrder[ref] = stage.Note_stagedOrder[note]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := note.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, note)
			// delete(stage.Note_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if note.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", note.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Notes_reference {
		instance := stage.Notes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Notes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			notes_deletedInstances = append(notes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(notes_newInstances)
	lenDeletedInstances += len(notes_deletedInstances)
	var noteshapes_newInstances []*NoteShape
	var noteshapes_deletedInstances []*NoteShape

	// parse all staged instances and check if they have a reference
	for noteshape := range stage.NoteShapes {
		if ref, ok := stage.NoteShapes_reference[noteshape]; !ok {
			noteshapes_newInstances = append(noteshapes_newInstances, noteshape)
			newInstancesSlice = append(newInstancesSlice, noteshape.GongMarshallIdentifier(stage))
			if stage.NoteShapes_referenceOrder == nil {
				stage.NoteShapes_referenceOrder = make(map[*NoteShape]uint)
			}
			stage.NoteShapes_referenceOrder[noteshape] = stage.NoteShape_stagedOrder[noteshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteshape.GongMarshallUnstaging(stage))
			// delete(stage.NoteShapes_referenceOrder, noteshape)
			fieldInitializers, pointersInitializations := noteshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NoteShape_stagedOrder[ref] = stage.NoteShape_stagedOrder[noteshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := noteshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteshape)
			// delete(stage.NoteShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if noteshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", noteshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.NoteShapes_reference {
		instance := stage.NoteShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NoteShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			noteshapes_deletedInstances = append(noteshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteshapes_newInstances)
	lenDeletedInstances += len(noteshapes_deletedInstances)
	var participants_newInstances []*Participant
	var participants_deletedInstances []*Participant

	// parse all staged instances and check if they have a reference
	for participant := range stage.Participants {
		if ref, ok := stage.Participants_reference[participant]; !ok {
			participants_newInstances = append(participants_newInstances, participant)
			newInstancesSlice = append(newInstancesSlice, participant.GongMarshallIdentifier(stage))
			if stage.Participants_referenceOrder == nil {
				stage.Participants_referenceOrder = make(map[*Participant]uint)
			}
			stage.Participants_referenceOrder[participant] = stage.Participant_stagedOrder[participant]
			newInstancesReverseSlice = append(newInstancesReverseSlice, participant.GongMarshallUnstaging(stage))
			// delete(stage.Participants_referenceOrder, participant)
			fieldInitializers, pointersInitializations := participant.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Participant_stagedOrder[ref] = stage.Participant_stagedOrder[participant]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := participant.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, participant)
			// delete(stage.Participant_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if participant.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", participant.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Participants_reference {
		instance := stage.Participants_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Participants[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			participants_deletedInstances = append(participants_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(participants_newInstances)
	lenDeletedInstances += len(participants_deletedInstances)
	var participantshapes_newInstances []*ParticipantShape
	var participantshapes_deletedInstances []*ParticipantShape

	// parse all staged instances and check if they have a reference
	for participantshape := range stage.ParticipantShapes {
		if ref, ok := stage.ParticipantShapes_reference[participantshape]; !ok {
			participantshapes_newInstances = append(participantshapes_newInstances, participantshape)
			newInstancesSlice = append(newInstancesSlice, participantshape.GongMarshallIdentifier(stage))
			if stage.ParticipantShapes_referenceOrder == nil {
				stage.ParticipantShapes_referenceOrder = make(map[*ParticipantShape]uint)
			}
			stage.ParticipantShapes_referenceOrder[participantshape] = stage.ParticipantShape_stagedOrder[participantshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, participantshape.GongMarshallUnstaging(stage))
			// delete(stage.ParticipantShapes_referenceOrder, participantshape)
			fieldInitializers, pointersInitializations := participantshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ParticipantShape_stagedOrder[ref] = stage.ParticipantShape_stagedOrder[participantshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := participantshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, participantshape)
			// delete(stage.ParticipantShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if participantshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", participantshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ParticipantShapes_reference {
		instance := stage.ParticipantShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ParticipantShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			participantshapes_deletedInstances = append(participantshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(participantshapes_newInstances)
	lenDeletedInstances += len(participantshapes_deletedInstances)
	var processs_newInstances []*Process
	var processs_deletedInstances []*Process

	// parse all staged instances and check if they have a reference
	for process := range stage.Processs {
		if ref, ok := stage.Processs_reference[process]; !ok {
			processs_newInstances = append(processs_newInstances, process)
			newInstancesSlice = append(newInstancesSlice, process.GongMarshallIdentifier(stage))
			if stage.Processs_referenceOrder == nil {
				stage.Processs_referenceOrder = make(map[*Process]uint)
			}
			stage.Processs_referenceOrder[process] = stage.Process_stagedOrder[process]
			newInstancesReverseSlice = append(newInstancesReverseSlice, process.GongMarshallUnstaging(stage))
			// delete(stage.Processs_referenceOrder, process)
			fieldInitializers, pointersInitializations := process.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Process_stagedOrder[ref] = stage.Process_stagedOrder[process]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := process.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, process)
			// delete(stage.Process_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if process.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", process.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Processs_reference {
		instance := stage.Processs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Processs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			processs_deletedInstances = append(processs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(processs_newInstances)
	lenDeletedInstances += len(processs_deletedInstances)
	var processshapes_newInstances []*ProcessShape
	var processshapes_deletedInstances []*ProcessShape

	// parse all staged instances and check if they have a reference
	for processshape := range stage.ProcessShapes {
		if ref, ok := stage.ProcessShapes_reference[processshape]; !ok {
			processshapes_newInstances = append(processshapes_newInstances, processshape)
			newInstancesSlice = append(newInstancesSlice, processshape.GongMarshallIdentifier(stage))
			if stage.ProcessShapes_referenceOrder == nil {
				stage.ProcessShapes_referenceOrder = make(map[*ProcessShape]uint)
			}
			stage.ProcessShapes_referenceOrder[processshape] = stage.ProcessShape_stagedOrder[processshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, processshape.GongMarshallUnstaging(stage))
			// delete(stage.ProcessShapes_referenceOrder, processshape)
			fieldInitializers, pointersInitializations := processshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ProcessShape_stagedOrder[ref] = stage.ProcessShape_stagedOrder[processshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := processshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, processshape)
			// delete(stage.ProcessShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if processshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", processshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.ProcessShapes_reference {
		instance := stage.ProcessShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ProcessShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			processshapes_deletedInstances = append(processshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(processshapes_newInstances)
	lenDeletedInstances += len(processshapes_deletedInstances)
	var resources_newInstances []*Resource
	var resources_deletedInstances []*Resource

	// parse all staged instances and check if they have a reference
	for resource := range stage.Resources {
		if ref, ok := stage.Resources_reference[resource]; !ok {
			resources_newInstances = append(resources_newInstances, resource)
			newInstancesSlice = append(newInstancesSlice, resource.GongMarshallIdentifier(stage))
			if stage.Resources_referenceOrder == nil {
				stage.Resources_referenceOrder = make(map[*Resource]uint)
			}
			stage.Resources_referenceOrder[resource] = stage.Resource_stagedOrder[resource]
			newInstancesReverseSlice = append(newInstancesReverseSlice, resource.GongMarshallUnstaging(stage))
			// delete(stage.Resources_referenceOrder, resource)
			fieldInitializers, pointersInitializations := resource.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Resource_stagedOrder[ref] = stage.Resource_stagedOrder[resource]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := resource.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, resource)
			// delete(stage.Resource_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if resource.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", resource.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Resources_reference {
		instance := stage.Resources_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Resources[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			resources_deletedInstances = append(resources_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(resources_newInstances)
	lenDeletedInstances += len(resources_deletedInstances)
	var tasks_newInstances []*Task
	var tasks_deletedInstances []*Task

	// parse all staged instances and check if they have a reference
	for task := range stage.Tasks {
		if ref, ok := stage.Tasks_reference[task]; !ok {
			tasks_newInstances = append(tasks_newInstances, task)
			newInstancesSlice = append(newInstancesSlice, task.GongMarshallIdentifier(stage))
			if stage.Tasks_referenceOrder == nil {
				stage.Tasks_referenceOrder = make(map[*Task]uint)
			}
			stage.Tasks_referenceOrder[task] = stage.Task_stagedOrder[task]
			newInstancesReverseSlice = append(newInstancesReverseSlice, task.GongMarshallUnstaging(stage))
			// delete(stage.Tasks_referenceOrder, task)
			fieldInitializers, pointersInitializations := task.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Task_stagedOrder[ref] = stage.Task_stagedOrder[task]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := task.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, task)
			// delete(stage.Task_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if task.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", task.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Tasks_reference {
		instance := stage.Tasks_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Tasks[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			tasks_deletedInstances = append(tasks_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(tasks_newInstances)
	lenDeletedInstances += len(tasks_deletedInstances)
	var taskshapes_newInstances []*TaskShape
	var taskshapes_deletedInstances []*TaskShape

	// parse all staged instances and check if they have a reference
	for taskshape := range stage.TaskShapes {
		if ref, ok := stage.TaskShapes_reference[taskshape]; !ok {
			taskshapes_newInstances = append(taskshapes_newInstances, taskshape)
			newInstancesSlice = append(newInstancesSlice, taskshape.GongMarshallIdentifier(stage))
			if stage.TaskShapes_referenceOrder == nil {
				stage.TaskShapes_referenceOrder = make(map[*TaskShape]uint)
			}
			stage.TaskShapes_referenceOrder[taskshape] = stage.TaskShape_stagedOrder[taskshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, taskshape.GongMarshallUnstaging(stage))
			// delete(stage.TaskShapes_referenceOrder, taskshape)
			fieldInitializers, pointersInitializations := taskshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TaskShape_stagedOrder[ref] = stage.TaskShape_stagedOrder[taskshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := taskshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, taskshape)
			// delete(stage.TaskShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if taskshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", taskshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.TaskShapes_reference {
		instance := stage.TaskShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.TaskShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			taskshapes_deletedInstances = append(taskshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(taskshapes_newInstances)
	lenDeletedInstances += len(taskshapes_deletedInstances)

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
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedresourceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AllocatedResourceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(allocatedresourceshape.Name))
	return
}

func (controlflow *ControlFlow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlflow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlFlow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlflow.Name))
	return
}

func (controlflowshape *ControlFlowShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlflowshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlFlowShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(controlflowshape.Name))
	return
}

func (data *Data) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", data.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Data")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(data.Name))
	return
}

func (dataflow *DataFlow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dataflow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DataFlow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(dataflow.Name))
	return
}

func (dataflowshape *DataFlowShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dataflowshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DataFlowShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(dataflowshape.Name))
	return
}

func (datashape *DataShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datashape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DataShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(datashape.Name))
	return
}

func (diagramprocess *DiagramProcess) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramProcess")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramprocess.Name))
	return
}

func (externalparticipantshape *ExternalParticipantShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", externalparticipantshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExternalParticipantShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(externalparticipantshape.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(note.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteshape.Name))
	return
}

func (participant *Participant) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", participant.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Participant")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(participant.Name))
	return
}

func (participantshape *ParticipantShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", participantshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ParticipantShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(participantshape.Name))
	return
}

func (process *Process) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", process.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Process")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(process.Name))
	return
}

func (processshape *ProcessShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", processshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProcessShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(processshape.Name))
	return
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Resource")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resource.Name))
	return
}

func (task *Task) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", task.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Task")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(task.Name))
	return
}

func (taskshape *TaskShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", taskshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TaskShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(taskshape.Name))
	return
}

// insertion point for unstaging
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

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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
