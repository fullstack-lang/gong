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
	stage.DiagramProcess_ProcessComposition_Shapes_reverseMap = make(map[*ProcessCompositionShape]*DiagramProcess)
	for diagramprocess := range stage.DiagramProcesss {
		_ = diagramprocess
		for _, _processcompositionshape := range diagramprocess.ProcessComposition_Shapes {
			stage.DiagramProcess_ProcessComposition_Shapes_reverseMap[_processcompositionshape] = diagramprocess
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

	// Compute reverse map for named struct Participant
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

	// Compute reverse map for named struct ProcessCompositionShape
	// insertion point per field

	// Compute reverse map for named struct ProcessShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.DiagramProcesss {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Participants {
		res = append(res, instance)
	}

	for instance := range stage.Processs {
		res = append(res, instance)
	}

	for instance := range stage.ProcessCompositionShapes {
		res = append(res, instance)
	}

	for instance := range stage.ProcessShapes {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
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

func (process *Process) GongCopy() GongstructIF {
	newInstance := new(Process)
	process.CopyBasicFields(newInstance)
	return newInstance
}

func (processcompositionshape *ProcessCompositionShape) GongCopy() GongstructIF {
	newInstance := new(ProcessCompositionShape)
	processcompositionshape.CopyBasicFields(newInstance)
	return newInstance
}

func (processshape *ProcessShape) GongCopy() GongstructIF {
	newInstance := new(ProcessShape)
	processshape.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
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

func (process *Process) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(process).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(process), uint64(GetOrderPointerGongstruct(stage, process)))
	return
}

func (processcompositionshape *ProcessCompositionShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(processcompositionshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(processcompositionshape), uint64(GetOrderPointerGongstruct(stage, processcompositionshape)))
	return
}

func (processshape *ProcessShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(processshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(processshape), uint64(GetOrderPointerGongstruct(stage, processshape)))
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
	var processcompositionshapes_newInstances []*ProcessCompositionShape
	var processcompositionshapes_deletedInstances []*ProcessCompositionShape

	// parse all staged instances and check if they have a reference
	for processcompositionshape := range stage.ProcessCompositionShapes {
		if ref, ok := stage.ProcessCompositionShapes_reference[processcompositionshape]; !ok {
			processcompositionshapes_newInstances = append(processcompositionshapes_newInstances, processcompositionshape)
			newInstancesSlice = append(newInstancesSlice, processcompositionshape.GongMarshallIdentifier(stage))
			if stage.ProcessCompositionShapes_referenceOrder == nil {
				stage.ProcessCompositionShapes_referenceOrder = make(map[*ProcessCompositionShape]uint)
			}
			stage.ProcessCompositionShapes_referenceOrder[processcompositionshape] = stage.ProcessCompositionShape_stagedOrder[processcompositionshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, processcompositionshape.GongMarshallUnstaging(stage))
			// delete(stage.ProcessCompositionShapes_referenceOrder, processcompositionshape)
			fieldInitializers, pointersInitializations := processcompositionshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ProcessCompositionShape_stagedOrder[ref] = stage.ProcessCompositionShape_stagedOrder[processcompositionshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := processcompositionshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, processcompositionshape)
			// delete(stage.ProcessCompositionShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", processcompositionshape.GetName())
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
	for _, ref := range stage.ProcessCompositionShapes_reference {
		instance := stage.ProcessCompositionShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ProcessCompositionShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			processcompositionshapes_deletedInstances = append(processcompositionshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(processcompositionshapes_newInstances)
	lenDeletedInstances += len(processcompositionshapes_deletedInstances)
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

	stage.Processs_reference = make(map[*Process]*Process)
	stage.Processs_referenceOrder = make(map[*Process]uint) // diff Unstage needs the reference order
	stage.Processs_instance = make(map[*Process]*Process)
	for instance := range stage.Processs {
		_copy := instance.GongCopy().(*Process)
		stage.Processs_reference[instance] = _copy
		stage.Processs_instance[_copy] = instance
		stage.Processs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ProcessCompositionShapes_reference = make(map[*ProcessCompositionShape]*ProcessCompositionShape)
	stage.ProcessCompositionShapes_referenceOrder = make(map[*ProcessCompositionShape]uint) // diff Unstage needs the reference order
	stage.ProcessCompositionShapes_instance = make(map[*ProcessCompositionShape]*ProcessCompositionShape)
	for instance := range stage.ProcessCompositionShapes {
		_copy := instance.GongCopy().(*ProcessCompositionShape)
		stage.ProcessCompositionShapes_reference[instance] = _copy
		stage.ProcessCompositionShapes_instance[_copy] = instance
		stage.ProcessCompositionShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	// insertion point per named struct
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

	for instance := range stage.Processs {
		reference := stage.Processs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ProcessCompositionShapes {
		reference := stage.ProcessCompositionShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ProcessShapes {
		reference := stage.ProcessShapes_reference[instance]
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

func (processcompositionshape *ProcessCompositionShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ProcessCompositionShape_stagedOrder[processcompositionshape]; ok {
		return order
	}
	if order, ok := stage.ProcessCompositionShapes_referenceOrder[processcompositionshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ProcessCompositionShape was not staged and does not have a reference order", processcompositionshape)
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

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
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

func (process *Process) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", process.GongGetGongstructName(), process.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (process *Process) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", process.GongGetGongstructName(), process.GongGetOrder(stage))
}

func (processcompositionshape *ProcessCompositionShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", processcompositionshape.GongGetGongstructName(), processcompositionshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (processcompositionshape *ProcessCompositionShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", processcompositionshape.GongGetGongstructName(), processcompositionshape.GongGetOrder(stage))
}

func (processshape *ProcessShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", processshape.GongGetGongstructName(), processshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (processshape *ProcessShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", processshape.GongGetGongstructName(), processshape.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
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

func (process *Process) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", process.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Process")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(process.Name))
	return
}

func (processcompositionshape *ProcessCompositionShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", processcompositionshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProcessCompositionShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(processcompositionshape.Name))
	return
}

func (processshape *ProcessShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", processshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ProcessShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(processshape.Name))
	return
}

// insertion point for unstaging
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

func (process *Process) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", process.GongGetReferenceIdentifier(stage))
	return
}

func (processcompositionshape *ProcessCompositionShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", processcompositionshape.GongGetReferenceIdentifier(stage))
	return
}

func (processshape *ProcessShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", processshape.GongGetReferenceIdentifier(stage))
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
