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

	// Compute reverse map for named struct AllocatedSystemShape
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

	// Compute reverse map for named struct DiagramStructure
	// insertion point per field
	stage.DiagramStructure_System_Shapes_reverseMap = make(map[*SystemShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _systemshape := range diagramstructure.System_Shapes {
			stage.DiagramStructure_System_Shapes_reverseMap[_systemshape] = diagramstructure
		}
	}
	stage.DiagramStructure_SystemsWhoseNodeIsExpanded_reverseMap = make(map[*System]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _system := range diagramstructure.SystemsWhoseNodeIsExpanded {
			stage.DiagramStructure_SystemsWhoseNodeIsExpanded_reverseMap[_system] = diagramstructure
		}
	}
	stage.DiagramStructure_Part_Shapes_reverseMap = make(map[*PartShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _partshape := range diagramstructure.Part_Shapes {
			stage.DiagramStructure_Part_Shapes_reverseMap[_partshape] = diagramstructure
		}
	}
	stage.DiagramStructure_PartWhoseNodeIsExpanded_reverseMap = make(map[*Part]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _part := range diagramstructure.PartWhoseNodeIsExpanded {
			stage.DiagramStructure_PartWhoseNodeIsExpanded_reverseMap[_part] = diagramstructure
		}
	}
	stage.DiagramStructure_ExternalPart_Shapes_reverseMap = make(map[*ExternalPartShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _externalpartshape := range diagramstructure.ExternalPart_Shapes {
			stage.DiagramStructure_ExternalPart_Shapes_reverseMap[_externalpartshape] = diagramstructure
		}
	}
	stage.DiagramStructure_ExternalPartWhoseNodeIsExpanded_reverseMap = make(map[*Part]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _part := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
			stage.DiagramStructure_ExternalPartWhoseNodeIsExpanded_reverseMap[_part] = diagramstructure
		}
	}
	stage.DiagramStructure_ExternalPartsWhoseOutDataFlowsNodeIsExpanded_reverseMap = make(map[*Part]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _part := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
			stage.DiagramStructure_ExternalPartsWhoseOutDataFlowsNodeIsExpanded_reverseMap[_part] = diagramstructure
		}
	}
	stage.DiagramStructure_ExternalPartsWhoseInDataFlowsNodeIsExpanded_reverseMap = make(map[*Part]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _part := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
			stage.DiagramStructure_ExternalPartsWhoseInDataFlowsNodeIsExpanded_reverseMap[_part] = diagramstructure
		}
	}
	stage.DiagramStructure_PortsWhoseNodeIsExpanded_reverseMap = make(map[*Port]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _port := range diagramstructure.PortsWhoseNodeIsExpanded {
			stage.DiagramStructure_PortsWhoseNodeIsExpanded_reverseMap[_port] = diagramstructure
		}
	}
	stage.DiagramStructure_Port_Shapes_reverseMap = make(map[*PortShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _portshape := range diagramstructure.Port_Shapes {
			stage.DiagramStructure_Port_Shapes_reverseMap[_portshape] = diagramstructure
		}
	}
	stage.DiagramStructure_ControlFlowsWhoseNodeIsExpanded_reverseMap = make(map[*ControlFlow]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _controlflow := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
			stage.DiagramStructure_ControlFlowsWhoseNodeIsExpanded_reverseMap[_controlflow] = diagramstructure
		}
	}
	stage.DiagramStructure_ControlFlow_Shapes_reverseMap = make(map[*ControlFlowShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _controlflowshape := range diagramstructure.ControlFlow_Shapes {
			stage.DiagramStructure_ControlFlow_Shapes_reverseMap[_controlflowshape] = diagramstructure
		}
	}
	stage.DiagramStructure_DataFlowsWhoseNodeIsExpanded_reverseMap = make(map[*DataFlow]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _dataflow := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
			stage.DiagramStructure_DataFlowsWhoseNodeIsExpanded_reverseMap[_dataflow] = diagramstructure
		}
	}
	stage.DiagramStructure_DataFlow_Shapes_reverseMap = make(map[*DataFlowShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _dataflowshape := range diagramstructure.DataFlow_Shapes {
			stage.DiagramStructure_DataFlow_Shapes_reverseMap[_dataflowshape] = diagramstructure
		}
	}
	stage.DiagramStructure_DatasWhoseNodeIsExpanded_reverseMap = make(map[*Data]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _data := range diagramstructure.DatasWhoseNodeIsExpanded {
			stage.DiagramStructure_DatasWhoseNodeIsExpanded_reverseMap[_data] = diagramstructure
		}
	}
	stage.DiagramStructure_Data_Shapes_reverseMap = make(map[*DataShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _datashape := range diagramstructure.Data_Shapes {
			stage.DiagramStructure_Data_Shapes_reverseMap[_datashape] = diagramstructure
		}
	}
	stage.DiagramStructure_DataFlowsWhoseDataNodeIsExpanded_reverseMap = make(map[*DataFlow]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _dataflow := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
			stage.DiagramStructure_DataFlowsWhoseDataNodeIsExpanded_reverseMap[_dataflow] = diagramstructure
		}
	}
	stage.DiagramStructure_AllocatedResourcesWhoseNodeIsExpanded_reverseMap = make(map[*Resource]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _resource := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
			stage.DiagramStructure_AllocatedResourcesWhoseNodeIsExpanded_reverseMap[_resource] = diagramstructure
		}
	}
	stage.DiagramStructure_AllocatedResourceShapes_reverseMap = make(map[*AllocatedResourceShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _allocatedresourceshape := range diagramstructure.AllocatedResourceShapes {
			stage.DiagramStructure_AllocatedResourceShapes_reverseMap[_allocatedresourceshape] = diagramstructure
		}
	}
	stage.DiagramStructure_AllocatedSystemesWhoseNodeIsExpanded_reverseMap = make(map[*System]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _system := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
			stage.DiagramStructure_AllocatedSystemesWhoseNodeIsExpanded_reverseMap[_system] = diagramstructure
		}
	}
	stage.DiagramStructure_AllocatedSystemShapes_reverseMap = make(map[*AllocatedSystemShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _allocatedsystemshape := range diagramstructure.AllocatedSystemShapes {
			stage.DiagramStructure_AllocatedSystemShapes_reverseMap[_allocatedsystemshape] = diagramstructure
		}
	}
	stage.DiagramStructure_Note_Shapes_reverseMap = make(map[*NoteShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _noteshape := range diagramstructure.Note_Shapes {
			stage.DiagramStructure_Note_Shapes_reverseMap[_noteshape] = diagramstructure
		}
	}
	stage.DiagramStructure_NotesWhoseNodeIsExpanded_reverseMap = make(map[*Note]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _note := range diagramstructure.NotesWhoseNodeIsExpanded {
			stage.DiagramStructure_NotesWhoseNodeIsExpanded_reverseMap[_note] = diagramstructure
		}
	}
	stage.DiagramStructure_NotePortShapes_reverseMap = make(map[*NotePortShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _noteportshape := range diagramstructure.NotePortShapes {
			stage.DiagramStructure_NotePortShapes_reverseMap[_noteportshape] = diagramstructure
		}
	}

	// Compute reverse map for named struct ExternalPartShape
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
	stage.Library_RootSystemes_reverseMap = make(map[*System]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _system := range library.RootSystemes {
			stage.Library_RootSystemes_reverseMap[_system] = library
		}
	}
	stage.Library_SystemsWhoseNodeIsExpanded_reverseMap = make(map[*System]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _system := range library.SystemsWhoseNodeIsExpanded {
			stage.Library_SystemsWhoseNodeIsExpanded_reverseMap[_system] = library
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
	stage.Library_PartsWhoseNodeIsExpanded_reverseMap = make(map[*Part]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _part := range library.PartsWhoseNodeIsExpanded {
			stage.Library_PartsWhoseNodeIsExpanded_reverseMap[_part] = library
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
	stage.Note_Ports_reverseMap = make(map[*Port]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _port := range note.Ports {
			stage.Note_Ports_reverseMap[_port] = note
		}
	}

	// Compute reverse map for named struct NotePortShape
	// insertion point per field

	// Compute reverse map for named struct NoteShape
	// insertion point per field

	// Compute reverse map for named struct Part
	// insertion point per field
	stage.Part_Ports_reverseMap = make(map[*Port]*Part)
	for part := range stage.Parts {
		_ = part
		for _, _port := range part.Ports {
			stage.Part_Ports_reverseMap[_port] = part
		}
	}
	stage.Part_ControlFlows_reverseMap = make(map[*ControlFlow]*Part)
	for part := range stage.Parts {
		_ = part
		for _, _controlflow := range part.ControlFlows {
			stage.Part_ControlFlows_reverseMap[_controlflow] = part
		}
	}
	stage.Part_PortWhoseOutControlFlowsNodeIsExpanded_reverseMap = make(map[*Port]*Part)
	for part := range stage.Parts {
		_ = part
		for _, _port := range part.PortWhoseOutControlFlowsNodeIsExpanded {
			stage.Part_PortWhoseOutControlFlowsNodeIsExpanded_reverseMap[_port] = part
		}
	}
	stage.Part_PortWhoseInControlFlowsNodeIsExpanded_reverseMap = make(map[*Port]*Part)
	for part := range stage.Parts {
		_ = part
		for _, _port := range part.PortWhoseInControlFlowsNodeIsExpanded {
			stage.Part_PortWhoseInControlFlowsNodeIsExpanded_reverseMap[_port] = part
		}
	}
	stage.Part_PortWhoseOutDataFlowsNodeIsExpanded_reverseMap = make(map[*Port]*Part)
	for part := range stage.Parts {
		_ = part
		for _, _port := range part.PortWhoseOutDataFlowsNodeIsExpanded {
			stage.Part_PortWhoseOutDataFlowsNodeIsExpanded_reverseMap[_port] = part
		}
	}
	stage.Part_PortWhoseInDataFlowsNodeIsExpanded_reverseMap = make(map[*Port]*Part)
	for part := range stage.Parts {
		_ = part
		for _, _port := range part.PortWhoseInDataFlowsNodeIsExpanded {
			stage.Part_PortWhoseInDataFlowsNodeIsExpanded_reverseMap[_port] = part
		}
	}
	stage.Part_PartAnchoredPath_reverseMap = make(map[*PartAnchoredPath]*Part)
	for part := range stage.Parts {
		_ = part
		for _, _partanchoredpath := range part.PartAnchoredPath {
			stage.Part_PartAnchoredPath_reverseMap[_partanchoredpath] = part
		}
	}

	// Compute reverse map for named struct PartAnchoredPath
	// insertion point per field

	// Compute reverse map for named struct PartShape
	// insertion point per field

	// Compute reverse map for named struct Port
	// insertion point per field

	// Compute reverse map for named struct PortShape
	// insertion point per field

	// Compute reverse map for named struct Resource
	// insertion point per field

	// Compute reverse map for named struct System
	// insertion point per field
	stage.System_DiagramStructures_reverseMap = make(map[*DiagramStructure]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramstructure := range system.DiagramStructures {
			stage.System_DiagramStructures_reverseMap[_diagramstructure] = system
		}
	}
	stage.System_DiagramStructureWhoseNodeIsExpanded_reverseMap = make(map[*DiagramStructure]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramstructure := range system.DiagramStructureWhoseNodeIsExpanded {
			stage.System_DiagramStructureWhoseNodeIsExpanded_reverseMap[_diagramstructure] = system
		}
	}
	stage.System_SubSystemes_reverseMap = make(map[*System]*System)
	for system := range stage.Systems {
		_ = system
		for _, _system := range system.SubSystemes {
			stage.System_SubSystemes_reverseMap[_system] = system
		}
	}
	stage.System_Parts_reverseMap = make(map[*Part]*System)
	for system := range stage.Systems {
		_ = system
		for _, _part := range system.Parts {
			stage.System_Parts_reverseMap[_part] = system
		}
	}
	stage.System_PartWhoseNodeIsExpanded_reverseMap = make(map[*Part]*System)
	for system := range stage.Systems {
		_ = system
		for _, _part := range system.PartWhoseNodeIsExpanded {
			stage.System_PartWhoseNodeIsExpanded_reverseMap[_part] = system
		}
	}
	stage.System_DataFlows_reverseMap = make(map[*DataFlow]*System)
	for system := range stage.Systems {
		_ = system
		for _, _dataflow := range system.DataFlows {
			stage.System_DataFlows_reverseMap[_dataflow] = system
		}
	}
	stage.System_ExternalParts_reverseMap = make(map[*Part]*System)
	for system := range stage.Systems {
		_ = system
		for _, _part := range system.ExternalParts {
			stage.System_ExternalParts_reverseMap[_part] = system
		}
	}
	stage.System_ExternalPartWhoseNodeIsExpanded_reverseMap = make(map[*Part]*System)
	for system := range stage.Systems {
		_ = system
		for _, _part := range system.ExternalPartWhoseNodeIsExpanded {
			stage.System_ExternalPartWhoseNodeIsExpanded_reverseMap[_part] = system
		}
	}

	// Compute reverse map for named struct SystemShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.AllocatedResourceShapes {
		res = append(res, instance)
	}

	for instance := range stage.AllocatedSystemShapes {
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

	for instance := range stage.DiagramStructures {
		res = append(res, instance)
	}

	for instance := range stage.ExternalPartShapes {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.NotePortShapes {
		res = append(res, instance)
	}

	for instance := range stage.NoteShapes {
		res = append(res, instance)
	}

	for instance := range stage.Parts {
		res = append(res, instance)
	}

	for instance := range stage.PartAnchoredPaths {
		res = append(res, instance)
	}

	for instance := range stage.PartShapes {
		res = append(res, instance)
	}

	for instance := range stage.Ports {
		res = append(res, instance)
	}

	for instance := range stage.PortShapes {
		res = append(res, instance)
	}

	for instance := range stage.Resources {
		res = append(res, instance)
	}

	for instance := range stage.Systems {
		res = append(res, instance)
	}

	for instance := range stage.SystemShapes {
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

func (allocatedsystemshape *AllocatedSystemShape) GongCopy() GongstructIF {
	newInstance := new(AllocatedSystemShape)
	allocatedsystemshape.CopyBasicFields(newInstance)
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

func (diagramstructure *DiagramStructure) GongCopy() GongstructIF {
	newInstance := new(DiagramStructure)
	diagramstructure.CopyBasicFields(newInstance)
	return newInstance
}

func (externalpartshape *ExternalPartShape) GongCopy() GongstructIF {
	newInstance := new(ExternalPartShape)
	externalpartshape.CopyBasicFields(newInstance)
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

func (noteportshape *NotePortShape) GongCopy() GongstructIF {
	newInstance := new(NotePortShape)
	noteportshape.CopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.CopyBasicFields(newInstance)
	return newInstance
}

func (part *Part) GongCopy() GongstructIF {
	newInstance := new(Part)
	part.CopyBasicFields(newInstance)
	return newInstance
}

func (partanchoredpath *PartAnchoredPath) GongCopy() GongstructIF {
	newInstance := new(PartAnchoredPath)
	partanchoredpath.CopyBasicFields(newInstance)
	return newInstance
}

func (partshape *PartShape) GongCopy() GongstructIF {
	newInstance := new(PartShape)
	partshape.CopyBasicFields(newInstance)
	return newInstance
}

func (port *Port) GongCopy() GongstructIF {
	newInstance := new(Port)
	port.CopyBasicFields(newInstance)
	return newInstance
}

func (portshape *PortShape) GongCopy() GongstructIF {
	newInstance := new(PortShape)
	portshape.CopyBasicFields(newInstance)
	return newInstance
}

func (resource *Resource) GongCopy() GongstructIF {
	newInstance := new(Resource)
	resource.CopyBasicFields(newInstance)
	return newInstance
}

func (system *System) GongCopy() GongstructIF {
	newInstance := new(System)
	system.CopyBasicFields(newInstance)
	return newInstance
}

func (systemshape *SystemShape) GongCopy() GongstructIF {
	newInstance := new(SystemShape)
	systemshape.CopyBasicFields(newInstance)
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

func (allocatedsystemshape *AllocatedSystemShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(allocatedsystemshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(allocatedsystemshape), uint64(GetOrderPointerGongstruct(stage, allocatedsystemshape)))
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

func (diagramstructure *DiagramStructure) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramstructure).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagramstructure), uint64(GetOrderPointerGongstruct(stage, diagramstructure)))
	return
}

func (externalpartshape *ExternalPartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(externalpartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(externalpartshape), uint64(GetOrderPointerGongstruct(stage, externalpartshape)))
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

func (noteportshape *NotePortShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteportshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(noteportshape), uint64(GetOrderPointerGongstruct(stage, noteportshape)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(noteshape), uint64(GetOrderPointerGongstruct(stage, noteshape)))
	return
}

func (part *Part) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(part), uint64(GetOrderPointerGongstruct(stage, part)))
	return
}

func (partanchoredpath *PartAnchoredPath) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partanchoredpath).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(partanchoredpath), uint64(GetOrderPointerGongstruct(stage, partanchoredpath)))
	return
}

func (partshape *PartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(partshape), uint64(GetOrderPointerGongstruct(stage, partshape)))
	return
}

func (port *Port) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(port).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(port), uint64(GetOrderPointerGongstruct(stage, port)))
	return
}

func (portshape *PortShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(portshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(portshape), uint64(GetOrderPointerGongstruct(stage, portshape)))
	return
}

func (resource *Resource) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(resource).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(resource), uint64(GetOrderPointerGongstruct(stage, resource)))
	return
}

func (system *System) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(system).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(system), uint64(GetOrderPointerGongstruct(stage, system)))
	return
}

func (systemshape *SystemShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(systemshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(systemshape), uint64(GetOrderPointerGongstruct(stage, systemshape)))
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
	var allocatedsystemshapes_newInstances []*AllocatedSystemShape
	var allocatedsystemshapes_deletedInstances []*AllocatedSystemShape

	// parse all staged instances and check if they have a reference
	for allocatedsystemshape := range stage.AllocatedSystemShapes {
		if ref, ok := stage.AllocatedSystemShapes_reference[allocatedsystemshape]; !ok {
			allocatedsystemshapes_newInstances = append(allocatedsystemshapes_newInstances, allocatedsystemshape)
			newInstancesSlice = append(newInstancesSlice, allocatedsystemshape.GongMarshallIdentifier(stage))
			if stage.AllocatedSystemShapes_referenceOrder == nil {
				stage.AllocatedSystemShapes_referenceOrder = make(map[*AllocatedSystemShape]uint)
			}
			stage.AllocatedSystemShapes_referenceOrder[allocatedsystemshape] = stage.AllocatedSystemShape_stagedOrder[allocatedsystemshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, allocatedsystemshape.GongMarshallUnstaging(stage))
			// delete(stage.AllocatedSystemShapes_referenceOrder, allocatedsystemshape)
			fieldInitializers, pointersInitializations := allocatedsystemshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AllocatedSystemShape_stagedOrder[ref] = stage.AllocatedSystemShape_stagedOrder[allocatedsystemshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := allocatedsystemshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, allocatedsystemshape)
			// delete(stage.AllocatedSystemShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if allocatedsystemshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", allocatedsystemshape.GetName())
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
	for _, ref := range stage.AllocatedSystemShapes_reference {
		instance := stage.AllocatedSystemShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.AllocatedSystemShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			allocatedsystemshapes_deletedInstances = append(allocatedsystemshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(allocatedsystemshapes_newInstances)
	lenDeletedInstances += len(allocatedsystemshapes_deletedInstances)
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
	var diagramstructures_newInstances []*DiagramStructure
	var diagramstructures_deletedInstances []*DiagramStructure

	// parse all staged instances and check if they have a reference
	for diagramstructure := range stage.DiagramStructures {
		if ref, ok := stage.DiagramStructures_reference[diagramstructure]; !ok {
			diagramstructures_newInstances = append(diagramstructures_newInstances, diagramstructure)
			newInstancesSlice = append(newInstancesSlice, diagramstructure.GongMarshallIdentifier(stage))
			if stage.DiagramStructures_referenceOrder == nil {
				stage.DiagramStructures_referenceOrder = make(map[*DiagramStructure]uint)
			}
			stage.DiagramStructures_referenceOrder[diagramstructure] = stage.DiagramStructure_stagedOrder[diagramstructure]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagramstructure.GongMarshallUnstaging(stage))
			// delete(stage.DiagramStructures_referenceOrder, diagramstructure)
			fieldInitializers, pointersInitializations := diagramstructure.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramStructure_stagedOrder[ref] = stage.DiagramStructure_stagedOrder[diagramstructure]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagramstructure.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagramstructure)
			// delete(stage.DiagramStructure_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if diagramstructure.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", diagramstructure.GetName())
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
	for _, ref := range stage.DiagramStructures_reference {
		instance := stage.DiagramStructures_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DiagramStructures[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagramstructures_deletedInstances = append(diagramstructures_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagramstructures_newInstances)
	lenDeletedInstances += len(diagramstructures_deletedInstances)
	var externalpartshapes_newInstances []*ExternalPartShape
	var externalpartshapes_deletedInstances []*ExternalPartShape

	// parse all staged instances and check if they have a reference
	for externalpartshape := range stage.ExternalPartShapes {
		if ref, ok := stage.ExternalPartShapes_reference[externalpartshape]; !ok {
			externalpartshapes_newInstances = append(externalpartshapes_newInstances, externalpartshape)
			newInstancesSlice = append(newInstancesSlice, externalpartshape.GongMarshallIdentifier(stage))
			if stage.ExternalPartShapes_referenceOrder == nil {
				stage.ExternalPartShapes_referenceOrder = make(map[*ExternalPartShape]uint)
			}
			stage.ExternalPartShapes_referenceOrder[externalpartshape] = stage.ExternalPartShape_stagedOrder[externalpartshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, externalpartshape.GongMarshallUnstaging(stage))
			// delete(stage.ExternalPartShapes_referenceOrder, externalpartshape)
			fieldInitializers, pointersInitializations := externalpartshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ExternalPartShape_stagedOrder[ref] = stage.ExternalPartShape_stagedOrder[externalpartshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := externalpartshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, externalpartshape)
			// delete(stage.ExternalPartShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if externalpartshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", externalpartshape.GetName())
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
	for _, ref := range stage.ExternalPartShapes_reference {
		instance := stage.ExternalPartShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ExternalPartShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			externalpartshapes_deletedInstances = append(externalpartshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(externalpartshapes_newInstances)
	lenDeletedInstances += len(externalpartshapes_deletedInstances)
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
	var noteportshapes_newInstances []*NotePortShape
	var noteportshapes_deletedInstances []*NotePortShape

	// parse all staged instances and check if they have a reference
	for noteportshape := range stage.NotePortShapes {
		if ref, ok := stage.NotePortShapes_reference[noteportshape]; !ok {
			noteportshapes_newInstances = append(noteportshapes_newInstances, noteportshape)
			newInstancesSlice = append(newInstancesSlice, noteportshape.GongMarshallIdentifier(stage))
			if stage.NotePortShapes_referenceOrder == nil {
				stage.NotePortShapes_referenceOrder = make(map[*NotePortShape]uint)
			}
			stage.NotePortShapes_referenceOrder[noteportshape] = stage.NotePortShape_stagedOrder[noteportshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, noteportshape.GongMarshallUnstaging(stage))
			// delete(stage.NotePortShapes_referenceOrder, noteportshape)
			fieldInitializers, pointersInitializations := noteportshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.NotePortShape_stagedOrder[ref] = stage.NotePortShape_stagedOrder[noteportshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := noteportshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, noteportshape)
			// delete(stage.NotePortShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if noteportshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", noteportshape.GetName())
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
	for _, ref := range stage.NotePortShapes_reference {
		instance := stage.NotePortShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.NotePortShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			noteportshapes_deletedInstances = append(noteportshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(noteportshapes_newInstances)
	lenDeletedInstances += len(noteportshapes_deletedInstances)
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
	var parts_newInstances []*Part
	var parts_deletedInstances []*Part

	// parse all staged instances and check if they have a reference
	for part := range stage.Parts {
		if ref, ok := stage.Parts_reference[part]; !ok {
			parts_newInstances = append(parts_newInstances, part)
			newInstancesSlice = append(newInstancesSlice, part.GongMarshallIdentifier(stage))
			if stage.Parts_referenceOrder == nil {
				stage.Parts_referenceOrder = make(map[*Part]uint)
			}
			stage.Parts_referenceOrder[part] = stage.Part_stagedOrder[part]
			newInstancesReverseSlice = append(newInstancesReverseSlice, part.GongMarshallUnstaging(stage))
			// delete(stage.Parts_referenceOrder, part)
			fieldInitializers, pointersInitializations := part.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Part_stagedOrder[ref] = stage.Part_stagedOrder[part]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := part.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, part)
			// delete(stage.Part_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if part.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", part.GetName())
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
	for _, ref := range stage.Parts_reference {
		instance := stage.Parts_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Parts[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parts_deletedInstances = append(parts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parts_newInstances)
	lenDeletedInstances += len(parts_deletedInstances)
	var partanchoredpaths_newInstances []*PartAnchoredPath
	var partanchoredpaths_deletedInstances []*PartAnchoredPath

	// parse all staged instances and check if they have a reference
	for partanchoredpath := range stage.PartAnchoredPaths {
		if ref, ok := stage.PartAnchoredPaths_reference[partanchoredpath]; !ok {
			partanchoredpaths_newInstances = append(partanchoredpaths_newInstances, partanchoredpath)
			newInstancesSlice = append(newInstancesSlice, partanchoredpath.GongMarshallIdentifier(stage))
			if stage.PartAnchoredPaths_referenceOrder == nil {
				stage.PartAnchoredPaths_referenceOrder = make(map[*PartAnchoredPath]uint)
			}
			stage.PartAnchoredPaths_referenceOrder[partanchoredpath] = stage.PartAnchoredPath_stagedOrder[partanchoredpath]
			newInstancesReverseSlice = append(newInstancesReverseSlice, partanchoredpath.GongMarshallUnstaging(stage))
			// delete(stage.PartAnchoredPaths_referenceOrder, partanchoredpath)
			fieldInitializers, pointersInitializations := partanchoredpath.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PartAnchoredPath_stagedOrder[ref] = stage.PartAnchoredPath_stagedOrder[partanchoredpath]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := partanchoredpath.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, partanchoredpath)
			// delete(stage.PartAnchoredPath_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if partanchoredpath.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", partanchoredpath.GetName())
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
	for _, ref := range stage.PartAnchoredPaths_reference {
		instance := stage.PartAnchoredPaths_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PartAnchoredPaths[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			partanchoredpaths_deletedInstances = append(partanchoredpaths_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(partanchoredpaths_newInstances)
	lenDeletedInstances += len(partanchoredpaths_deletedInstances)
	var partshapes_newInstances []*PartShape
	var partshapes_deletedInstances []*PartShape

	// parse all staged instances and check if they have a reference
	for partshape := range stage.PartShapes {
		if ref, ok := stage.PartShapes_reference[partshape]; !ok {
			partshapes_newInstances = append(partshapes_newInstances, partshape)
			newInstancesSlice = append(newInstancesSlice, partshape.GongMarshallIdentifier(stage))
			if stage.PartShapes_referenceOrder == nil {
				stage.PartShapes_referenceOrder = make(map[*PartShape]uint)
			}
			stage.PartShapes_referenceOrder[partshape] = stage.PartShape_stagedOrder[partshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, partshape.GongMarshallUnstaging(stage))
			// delete(stage.PartShapes_referenceOrder, partshape)
			fieldInitializers, pointersInitializations := partshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PartShape_stagedOrder[ref] = stage.PartShape_stagedOrder[partshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := partshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, partshape)
			// delete(stage.PartShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if partshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", partshape.GetName())
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
	for _, ref := range stage.PartShapes_reference {
		instance := stage.PartShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PartShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			partshapes_deletedInstances = append(partshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(partshapes_newInstances)
	lenDeletedInstances += len(partshapes_deletedInstances)
	var ports_newInstances []*Port
	var ports_deletedInstances []*Port

	// parse all staged instances and check if they have a reference
	for port := range stage.Ports {
		if ref, ok := stage.Ports_reference[port]; !ok {
			ports_newInstances = append(ports_newInstances, port)
			newInstancesSlice = append(newInstancesSlice, port.GongMarshallIdentifier(stage))
			if stage.Ports_referenceOrder == nil {
				stage.Ports_referenceOrder = make(map[*Port]uint)
			}
			stage.Ports_referenceOrder[port] = stage.Port_stagedOrder[port]
			newInstancesReverseSlice = append(newInstancesReverseSlice, port.GongMarshallUnstaging(stage))
			// delete(stage.Ports_referenceOrder, port)
			fieldInitializers, pointersInitializations := port.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Port_stagedOrder[ref] = stage.Port_stagedOrder[port]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := port.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, port)
			// delete(stage.Port_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if port.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", port.GetName())
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
	for _, ref := range stage.Ports_reference {
		instance := stage.Ports_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Ports[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			ports_deletedInstances = append(ports_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(ports_newInstances)
	lenDeletedInstances += len(ports_deletedInstances)
	var portshapes_newInstances []*PortShape
	var portshapes_deletedInstances []*PortShape

	// parse all staged instances and check if they have a reference
	for portshape := range stage.PortShapes {
		if ref, ok := stage.PortShapes_reference[portshape]; !ok {
			portshapes_newInstances = append(portshapes_newInstances, portshape)
			newInstancesSlice = append(newInstancesSlice, portshape.GongMarshallIdentifier(stage))
			if stage.PortShapes_referenceOrder == nil {
				stage.PortShapes_referenceOrder = make(map[*PortShape]uint)
			}
			stage.PortShapes_referenceOrder[portshape] = stage.PortShape_stagedOrder[portshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, portshape.GongMarshallUnstaging(stage))
			// delete(stage.PortShapes_referenceOrder, portshape)
			fieldInitializers, pointersInitializations := portshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PortShape_stagedOrder[ref] = stage.PortShape_stagedOrder[portshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := portshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, portshape)
			// delete(stage.PortShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if portshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", portshape.GetName())
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
	for _, ref := range stage.PortShapes_reference {
		instance := stage.PortShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PortShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			portshapes_deletedInstances = append(portshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(portshapes_newInstances)
	lenDeletedInstances += len(portshapes_deletedInstances)
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
	var systems_newInstances []*System
	var systems_deletedInstances []*System

	// parse all staged instances and check if they have a reference
	for system := range stage.Systems {
		if ref, ok := stage.Systems_reference[system]; !ok {
			systems_newInstances = append(systems_newInstances, system)
			newInstancesSlice = append(newInstancesSlice, system.GongMarshallIdentifier(stage))
			if stage.Systems_referenceOrder == nil {
				stage.Systems_referenceOrder = make(map[*System]uint)
			}
			stage.Systems_referenceOrder[system] = stage.System_stagedOrder[system]
			newInstancesReverseSlice = append(newInstancesReverseSlice, system.GongMarshallUnstaging(stage))
			// delete(stage.Systems_referenceOrder, system)
			fieldInitializers, pointersInitializations := system.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.System_stagedOrder[ref] = stage.System_stagedOrder[system]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := system.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, system)
			// delete(stage.System_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if system.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", system.GetName())
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
	for _, ref := range stage.Systems_reference {
		instance := stage.Systems_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Systems[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			systems_deletedInstances = append(systems_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(systems_newInstances)
	lenDeletedInstances += len(systems_deletedInstances)
	var systemshapes_newInstances []*SystemShape
	var systemshapes_deletedInstances []*SystemShape

	// parse all staged instances and check if they have a reference
	for systemshape := range stage.SystemShapes {
		if ref, ok := stage.SystemShapes_reference[systemshape]; !ok {
			systemshapes_newInstances = append(systemshapes_newInstances, systemshape)
			newInstancesSlice = append(newInstancesSlice, systemshape.GongMarshallIdentifier(stage))
			if stage.SystemShapes_referenceOrder == nil {
				stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint)
			}
			stage.SystemShapes_referenceOrder[systemshape] = stage.SystemShape_stagedOrder[systemshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, systemshape.GongMarshallUnstaging(stage))
			// delete(stage.SystemShapes_referenceOrder, systemshape)
			fieldInitializers, pointersInitializations := systemshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SystemShape_stagedOrder[ref] = stage.SystemShape_stagedOrder[systemshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := systemshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, systemshape)
			// delete(stage.SystemShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if systemshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", systemshape.GetName())
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
	for _, ref := range stage.SystemShapes_reference {
		instance := stage.SystemShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SystemShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			systemshapes_deletedInstances = append(systemshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(systemshapes_newInstances)
	lenDeletedInstances += len(systemshapes_deletedInstances)

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

	stage.AllocatedSystemShapes_reference = make(map[*AllocatedSystemShape]*AllocatedSystemShape)
	stage.AllocatedSystemShapes_referenceOrder = make(map[*AllocatedSystemShape]uint) // diff Unstage needs the reference order
	stage.AllocatedSystemShapes_instance = make(map[*AllocatedSystemShape]*AllocatedSystemShape)
	for instance := range stage.AllocatedSystemShapes {
		_copy := instance.GongCopy().(*AllocatedSystemShape)
		stage.AllocatedSystemShapes_reference[instance] = _copy
		stage.AllocatedSystemShapes_instance[_copy] = instance
		stage.AllocatedSystemShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.DiagramStructures_reference = make(map[*DiagramStructure]*DiagramStructure)
	stage.DiagramStructures_referenceOrder = make(map[*DiagramStructure]uint) // diff Unstage needs the reference order
	stage.DiagramStructures_instance = make(map[*DiagramStructure]*DiagramStructure)
	for instance := range stage.DiagramStructures {
		_copy := instance.GongCopy().(*DiagramStructure)
		stage.DiagramStructures_reference[instance] = _copy
		stage.DiagramStructures_instance[_copy] = instance
		stage.DiagramStructures_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ExternalPartShapes_reference = make(map[*ExternalPartShape]*ExternalPartShape)
	stage.ExternalPartShapes_referenceOrder = make(map[*ExternalPartShape]uint) // diff Unstage needs the reference order
	stage.ExternalPartShapes_instance = make(map[*ExternalPartShape]*ExternalPartShape)
	for instance := range stage.ExternalPartShapes {
		_copy := instance.GongCopy().(*ExternalPartShape)
		stage.ExternalPartShapes_reference[instance] = _copy
		stage.ExternalPartShapes_instance[_copy] = instance
		stage.ExternalPartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.NotePortShapes_reference = make(map[*NotePortShape]*NotePortShape)
	stage.NotePortShapes_referenceOrder = make(map[*NotePortShape]uint) // diff Unstage needs the reference order
	stage.NotePortShapes_instance = make(map[*NotePortShape]*NotePortShape)
	for instance := range stage.NotePortShapes {
		_copy := instance.GongCopy().(*NotePortShape)
		stage.NotePortShapes_reference[instance] = _copy
		stage.NotePortShapes_instance[_copy] = instance
		stage.NotePortShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Parts_reference = make(map[*Part]*Part)
	stage.Parts_referenceOrder = make(map[*Part]uint) // diff Unstage needs the reference order
	stage.Parts_instance = make(map[*Part]*Part)
	for instance := range stage.Parts {
		_copy := instance.GongCopy().(*Part)
		stage.Parts_reference[instance] = _copy
		stage.Parts_instance[_copy] = instance
		stage.Parts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartAnchoredPaths_reference = make(map[*PartAnchoredPath]*PartAnchoredPath)
	stage.PartAnchoredPaths_referenceOrder = make(map[*PartAnchoredPath]uint) // diff Unstage needs the reference order
	stage.PartAnchoredPaths_instance = make(map[*PartAnchoredPath]*PartAnchoredPath)
	for instance := range stage.PartAnchoredPaths {
		_copy := instance.GongCopy().(*PartAnchoredPath)
		stage.PartAnchoredPaths_reference[instance] = _copy
		stage.PartAnchoredPaths_instance[_copy] = instance
		stage.PartAnchoredPaths_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartShapes_reference = make(map[*PartShape]*PartShape)
	stage.PartShapes_referenceOrder = make(map[*PartShape]uint) // diff Unstage needs the reference order
	stage.PartShapes_instance = make(map[*PartShape]*PartShape)
	for instance := range stage.PartShapes {
		_copy := instance.GongCopy().(*PartShape)
		stage.PartShapes_reference[instance] = _copy
		stage.PartShapes_instance[_copy] = instance
		stage.PartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Ports_reference = make(map[*Port]*Port)
	stage.Ports_referenceOrder = make(map[*Port]uint) // diff Unstage needs the reference order
	stage.Ports_instance = make(map[*Port]*Port)
	for instance := range stage.Ports {
		_copy := instance.GongCopy().(*Port)
		stage.Ports_reference[instance] = _copy
		stage.Ports_instance[_copy] = instance
		stage.Ports_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PortShapes_reference = make(map[*PortShape]*PortShape)
	stage.PortShapes_referenceOrder = make(map[*PortShape]uint) // diff Unstage needs the reference order
	stage.PortShapes_instance = make(map[*PortShape]*PortShape)
	for instance := range stage.PortShapes {
		_copy := instance.GongCopy().(*PortShape)
		stage.PortShapes_reference[instance] = _copy
		stage.PortShapes_instance[_copy] = instance
		stage.PortShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Systems_reference = make(map[*System]*System)
	stage.Systems_referenceOrder = make(map[*System]uint) // diff Unstage needs the reference order
	stage.Systems_instance = make(map[*System]*System)
	for instance := range stage.Systems {
		_copy := instance.GongCopy().(*System)
		stage.Systems_reference[instance] = _copy
		stage.Systems_instance[_copy] = instance
		stage.Systems_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SystemShapes_reference = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint) // diff Unstage needs the reference order
	stage.SystemShapes_instance = make(map[*SystemShape]*SystemShape)
	for instance := range stage.SystemShapes {
		_copy := instance.GongCopy().(*SystemShape)
		stage.SystemShapes_reference[instance] = _copy
		stage.SystemShapes_instance[_copy] = instance
		stage.SystemShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.AllocatedResourceShapes {
		reference := stage.AllocatedResourceShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.AllocatedSystemShapes {
		reference := stage.AllocatedSystemShapes_reference[instance]
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

	for instance := range stage.DiagramStructures {
		reference := stage.DiagramStructures_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ExternalPartShapes {
		reference := stage.ExternalPartShapes_reference[instance]
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

	for instance := range stage.NotePortShapes {
		reference := stage.NotePortShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.NoteShapes {
		reference := stage.NoteShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Parts {
		reference := stage.Parts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartAnchoredPaths {
		reference := stage.PartAnchoredPaths_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartShapes {
		reference := stage.PartShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Ports {
		reference := stage.Ports_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PortShapes {
		reference := stage.PortShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Resources {
		reference := stage.Resources_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Systems {
		reference := stage.Systems_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SystemShapes {
		reference := stage.SystemShapes_reference[instance]
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

func (allocatedsystemshape *AllocatedSystemShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AllocatedSystemShape_stagedOrder[allocatedsystemshape]; ok {
		return order
	}
	if order, ok := stage.AllocatedSystemShapes_referenceOrder[allocatedsystemshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type AllocatedSystemShape was not staged and does not have a reference order", allocatedsystemshape)
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

func (diagramstructure *DiagramStructure) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramStructure_stagedOrder[diagramstructure]; ok {
		return order
	}
	if order, ok := stage.DiagramStructures_referenceOrder[diagramstructure]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramStructure was not staged and does not have a reference order", diagramstructure)
		return 0
	}
}

func (externalpartshape *ExternalPartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ExternalPartShape_stagedOrder[externalpartshape]; ok {
		return order
	}
	if order, ok := stage.ExternalPartShapes_referenceOrder[externalpartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type ExternalPartShape was not staged and does not have a reference order", externalpartshape)
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

func (noteportshape *NotePortShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NotePortShape_stagedOrder[noteportshape]; ok {
		return order
	}
	if order, ok := stage.NotePortShapes_referenceOrder[noteportshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NotePortShape was not staged and does not have a reference order", noteportshape)
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

func (part *Part) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_stagedOrder[part]; ok {
		return order
	}
	if order, ok := stage.Parts_referenceOrder[part]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part was not staged and does not have a reference order", part)
		return 0
	}
}

func (partanchoredpath *PartAnchoredPath) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartAnchoredPath_stagedOrder[partanchoredpath]; ok {
		return order
	}
	if order, ok := stage.PartAnchoredPaths_referenceOrder[partanchoredpath]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartAnchoredPath was not staged and does not have a reference order", partanchoredpath)
		return 0
	}
}

func (partshape *PartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartShape_stagedOrder[partshape]; ok {
		return order
	}
	if order, ok := stage.PartShapes_referenceOrder[partshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartShape was not staged and does not have a reference order", partshape)
		return 0
	}
}

func (port *Port) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Port_stagedOrder[port]; ok {
		return order
	}
	if order, ok := stage.Ports_referenceOrder[port]; ok {
		return order
	} else {
		log.Printf("instance %p of type Port was not staged and does not have a reference order", port)
		return 0
	}
}

func (portshape *PortShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PortShape_stagedOrder[portshape]; ok {
		return order
	}
	if order, ok := stage.PortShapes_referenceOrder[portshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PortShape was not staged and does not have a reference order", portshape)
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

func (system *System) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.System_stagedOrder[system]; ok {
		return order
	}
	if order, ok := stage.Systems_referenceOrder[system]; ok {
		return order
	} else {
		log.Printf("instance %p of type System was not staged and does not have a reference order", system)
		return 0
	}
}

func (systemshape *SystemShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SystemShape_stagedOrder[systemshape]; ok {
		return order
	}
	if order, ok := stage.SystemShapes_referenceOrder[systemshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type SystemShape was not staged and does not have a reference order", systemshape)
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

func (allocatedsystemshape *AllocatedSystemShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", allocatedsystemshape.GongGetGongstructName(), allocatedsystemshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (allocatedsystemshape *AllocatedSystemShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", allocatedsystemshape.GongGetGongstructName(), allocatedsystemshape.GongGetOrder(stage))
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

func (diagramstructure *DiagramStructure) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramstructure.GongGetGongstructName(), diagramstructure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramstructure *DiagramStructure) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramstructure.GongGetGongstructName(), diagramstructure.GongGetOrder(stage))
}

func (externalpartshape *ExternalPartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", externalpartshape.GongGetGongstructName(), externalpartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (externalpartshape *ExternalPartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", externalpartshape.GongGetGongstructName(), externalpartshape.GongGetOrder(stage))
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

func (noteportshape *NotePortShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteportshape.GongGetGongstructName(), noteportshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteportshape *NotePortShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteportshape.GongGetGongstructName(), noteportshape.GongGetOrder(stage))
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", noteshape.GongGetGongstructName(), noteshape.GongGetOrder(stage))
}

func (part *Part) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part.GongGetGongstructName(), part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part *Part) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part.GongGetGongstructName(), part.GongGetOrder(stage))
}

func (partanchoredpath *PartAnchoredPath) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partanchoredpath.GongGetGongstructName(), partanchoredpath.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partanchoredpath *PartAnchoredPath) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partanchoredpath.GongGetGongstructName(), partanchoredpath.GongGetOrder(stage))
}

func (partshape *PartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partshape.GongGetGongstructName(), partshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partshape *PartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partshape.GongGetGongstructName(), partshape.GongGetOrder(stage))
}

func (port *Port) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", port.GongGetGongstructName(), port.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (port *Port) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", port.GongGetGongstructName(), port.GongGetOrder(stage))
}

func (portshape *PortShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", portshape.GongGetGongstructName(), portshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (portshape *PortShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", portshape.GongGetGongstructName(), portshape.GongGetOrder(stage))
}

func (resource *Resource) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resource.GongGetGongstructName(), resource.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resource *Resource) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", resource.GongGetGongstructName(), resource.GongGetOrder(stage))
}

func (system *System) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system.GongGetGongstructName(), system.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system *System) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system.GongGetGongstructName(), system.GongGetOrder(stage))
}

func (systemshape *SystemShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", systemshape.GongGetGongstructName(), systemshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (systemshape *SystemShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", systemshape.GongGetGongstructName(), systemshape.GongGetOrder(stage))
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

func (allocatedsystemshape *AllocatedSystemShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedsystemshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AllocatedSystemShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(allocatedsystemshape.Name))
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

func (diagramstructure *DiagramStructure) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramStructure")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramstructure.Name))
	return
}

func (externalpartshape *ExternalPartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExternalPartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(externalpartshape.Name))
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

func (noteportshape *NotePortShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NotePortShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteportshape.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(noteshape.Name))
	return
}

func (part *Part) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(part.Name))
	return
}

func (partanchoredpath *PartAnchoredPath) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partanchoredpath.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartAnchoredPath")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(partanchoredpath.Name))
	return
}

func (partshape *PartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(partshape.Name))
	return
}

func (port *Port) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", port.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Port")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(port.Name))
	return
}

func (portshape *PortShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", portshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PortShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(portshape.Name))
	return
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Resource")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(resource.Name))
	return
}

func (system *System) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.Name))
	return
}

func (systemshape *SystemShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SystemShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(systemshape.Name))
	return
}

// insertion point for unstaging
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedresourceshape.GongGetReferenceIdentifier(stage))
	return
}

func (allocatedsystemshape *AllocatedSystemShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedsystemshape.GongGetReferenceIdentifier(stage))
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

func (diagramstructure *DiagramStructure) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramstructure.GongGetReferenceIdentifier(stage))
	return
}

func (externalpartshape *ExternalPartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", externalpartshape.GongGetReferenceIdentifier(stage))
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

func (noteportshape *NotePortShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteportshape.GongGetReferenceIdentifier(stage))
	return
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetReferenceIdentifier(stage))
	return
}

func (part *Part) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part.GongGetReferenceIdentifier(stage))
	return
}

func (partanchoredpath *PartAnchoredPath) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partanchoredpath.GongGetReferenceIdentifier(stage))
	return
}

func (partshape *PartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partshape.GongGetReferenceIdentifier(stage))
	return
}

func (port *Port) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", port.GongGetReferenceIdentifier(stage))
	return
}

func (portshape *PortShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", portshape.GongGetReferenceIdentifier(stage))
	return
}

func (resource *Resource) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetReferenceIdentifier(stage))
	return
}

func (system *System) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetReferenceIdentifier(stage))
	return
}

func (systemshape *SystemShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", systemshape.GongGetReferenceIdentifier(stage))
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
