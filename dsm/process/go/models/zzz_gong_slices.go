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
	// Compute reverse map for named struct ControlFlow
	// insertion point per field

	// Compute reverse map for named struct ControlFlowShape
	// insertion point per field

	// Compute reverse map for named struct DataFlow
	// insertion point per field

	// Compute reverse map for named struct DataFlowShape
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
	stage.DiagramProcess_TasksWhoseNodeIsExpanded_reverseMap = make(map[*Task]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _task := range diagramprocess.TasksWhoseNodeIsExpanded {
			stage.DiagramProcess_TasksWhoseNodeIsExpanded_reverseMap[_task] = diagramprocess
		}
	}
	stage.DiagramProcess_TaskShapes_reverseMap = make(map[*TaskShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _taskshape := range diagramprocess.TaskShapes {
			stage.DiagramProcess_TaskShapes_reverseMap[_taskshape] = diagramprocess
		}
	}
	stage.DiagramProcess_ControlFlowsWhoseNodeIsExpanded_reverseMap = make(map[*ControlFlow]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _controlflow := range diagramprocess.ControlFlowsWhoseNodeIsExpanded {
			stage.DiagramProcess_ControlFlowsWhoseNodeIsExpanded_reverseMap[_controlflow] = diagramprocess
		}
	}
	stage.DiagramProcess_ControlFlowShapes_reverseMap = make(map[*ControlFlowShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _controlflowshape := range diagramprocess.ControlFlowShapes {
			stage.DiagramProcess_ControlFlowShapes_reverseMap[_controlflowshape] = diagramprocess
		}
	}
	stage.DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap = make(map[*DataFlow]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _dataflow := range diagramprocess.DataFlowsWhoseNodeIsExpanded {
			stage.DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap[_dataflow] = diagramprocess
		}
	}
	stage.DiagramProcess_DataFlowShapes_reverseMap = make(map[*DataFlowShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _dataflowshape := range diagramprocess.DataFlowShapes {
			stage.DiagramProcess_DataFlowShapes_reverseMap[_dataflowshape] = diagramprocess
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

	// Compute reverse map for named struct Participant
	// insertion point per field
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

	// Compute reverse map for named struct ProcessShape
	// insertion point per field

	// Compute reverse map for named struct Task
	// insertion point per field

	// Compute reverse map for named struct TaskShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.ControlFlows {
		res = append(res, instance)
	}

	for instance := range stage.ControlFlowShapes {
		res = append(res, instance)
	}

	for instance := range stage.DataFlows {
		res = append(res, instance)
	}

	for instance := range stage.DataFlowShapes {
		res = append(res, instance)
	}

	for instance := range stage.DiagramProcesss {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
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

	for instance := range stage.Tasks {
		res = append(res, instance)
	}

	for instance := range stage.TaskShapes {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
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

func (diagramprocess *DiagramProcess) GongCopy() GongstructIF {
	newInstance := new(DiagramProcess)
	diagramprocess.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
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

func (diagramprocess *DiagramProcess) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramprocess).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagramprocess), uint64(GetOrderPointerGongstruct(stage, diagramprocess)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", controlflow.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", controlflowshape.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", dataflow.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", dataflowshape.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", diagramprocess.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", library.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", participant.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", participantshape.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", process.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", processshape.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", task.GetName())
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
				fieldsEdit += fmt.Sprintf("\n\t// %s", taskshape.GetName())
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

	stage.DiagramProcesss_reference = make(map[*DiagramProcess]*DiagramProcess)
	stage.DiagramProcesss_referenceOrder = make(map[*DiagramProcess]uint) // diff Unstage needs the reference order
	stage.DiagramProcesss_instance = make(map[*DiagramProcess]*DiagramProcess)
	for instance := range stage.DiagramProcesss {
		_copy := instance.GongCopy().(*DiagramProcess)
		stage.DiagramProcesss_reference[instance] = _copy
		stage.DiagramProcesss_instance[_copy] = instance
		stage.DiagramProcesss_referenceOrder[_copy] = instance.GongGetOrder(stage)
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
	for instance := range stage.ControlFlows {
		reference := stage.ControlFlows_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ControlFlowShapes {
		reference := stage.ControlFlowShapes_reference[instance]
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

	for instance := range stage.DiagramProcesss {
		reference := stage.DiagramProcesss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
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

func (diagramprocess *DiagramProcess) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramprocess.GongGetGongstructName(), diagramprocess.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramprocess *DiagramProcess) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramprocess.GongGetGongstructName(), diagramprocess.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
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

func (diagramprocess *DiagramProcess) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramprocess.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramProcess")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramprocess.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
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

func (diagramprocess *DiagramProcess) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramprocess.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
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
