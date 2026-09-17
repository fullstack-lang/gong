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

	// Compute reverse map for named struct DiagramLayerState
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
	stage.DiagramStructure_NotePartShapes_reverseMap = make(map[*NotePartShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _notepartshape := range diagramstructure.NotePartShapes {
			stage.DiagramStructure_NotePartShapes_reverseMap[_notepartshape] = diagramstructure
		}
	}

	// Compute reverse map for named struct ExternalPartShape
	// insertion point per field

	// Compute reverse map for named struct LayerDefinition
	// insertion point per field
	stage.LayerDefinition_Query_reverseMap = make(map[*SemanticTag]*LayerDefinition)
	for layerdefinition := range stage.LayerDefinitions {
		_ = layerdefinition
		for _, _semantictag := range layerdefinition.Query {
			stage.LayerDefinition_Query_reverseMap[_semantictag] = layerdefinition
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
	stage.Note_Parts_reverseMap = make(map[*Part]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _part := range note.Parts {
			stage.Note_Parts_reverseMap[_part] = note
		}
	}
	stage.Note_Ports_reverseMap = make(map[*Port]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _port := range note.Ports {
			stage.Note_Ports_reverseMap[_port] = note
		}
	}

	// Compute reverse map for named struct NotePartShape
	// insertion point per field

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

	// Compute reverse map for named struct SemanticTag
	// insertion point per field
	stage.SemanticTag_Parts_reverseMap = make(map[*Part]*SemanticTag)
	for semantictag := range stage.SemanticTags {
		_ = semantictag
		for _, _part := range semantictag.Parts {
			stage.SemanticTag_Parts_reverseMap[_part] = semantictag
		}
	}

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

	for instance := range stage.DiagramLayerStates {
		res = append(res, instance)
	}

	for instance := range stage.DiagramStructures {
		res = append(res, instance)
	}

	for instance := range stage.ExternalPartShapes {
		res = append(res, instance)
	}

	for instance := range stage.LayerDefinitions {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.NotePartShapes {
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

	for instance := range stage.SemanticTags {
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
	allocatedresourceshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (allocatedsystemshape *AllocatedSystemShape) GongCopy() GongstructIF {
	newInstance := new(AllocatedSystemShape)
	allocatedsystemshape.GongCopyBasicFields(newInstance)
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

func (diagramlayerstate *DiagramLayerState) GongCopy() GongstructIF {
	newInstance := new(DiagramLayerState)
	diagramlayerstate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagramstructure *DiagramStructure) GongCopy() GongstructIF {
	newInstance := new(DiagramStructure)
	diagramstructure.GongCopyBasicFields(newInstance)
	return newInstance
}

func (externalpartshape *ExternalPartShape) GongCopy() GongstructIF {
	newInstance := new(ExternalPartShape)
	externalpartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (layerdefinition *LayerDefinition) GongCopy() GongstructIF {
	newInstance := new(LayerDefinition)
	layerdefinition.GongCopyBasicFields(newInstance)
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

func (notepartshape *NotePartShape) GongCopy() GongstructIF {
	newInstance := new(NotePartShape)
	notepartshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteportshape *NotePortShape) GongCopy() GongstructIF {
	newInstance := new(NotePortShape)
	noteportshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (noteshape *NoteShape) GongCopy() GongstructIF {
	newInstance := new(NoteShape)
	noteshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part *Part) GongCopy() GongstructIF {
	newInstance := new(Part)
	part.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partanchoredpath *PartAnchoredPath) GongCopy() GongstructIF {
	newInstance := new(PartAnchoredPath)
	partanchoredpath.GongCopyBasicFields(newInstance)
	return newInstance
}

func (partshape *PartShape) GongCopy() GongstructIF {
	newInstance := new(PartShape)
	partshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (port *Port) GongCopy() GongstructIF {
	newInstance := new(Port)
	port.GongCopyBasicFields(newInstance)
	return newInstance
}

func (portshape *PortShape) GongCopy() GongstructIF {
	newInstance := new(PortShape)
	portshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (resource *Resource) GongCopy() GongstructIF {
	newInstance := new(Resource)
	resource.GongCopyBasicFields(newInstance)
	return newInstance
}

func (semantictag *SemanticTag) GongCopy() GongstructIF {
	newInstance := new(SemanticTag)
	semantictag.GongCopyBasicFields(newInstance)
	return newInstance
}

func (system *System) GongCopy() GongstructIF {
	newInstance := new(System)
	system.GongCopyBasicFields(newInstance)
	return newInstance
}

func (systemshape *SystemShape) GongCopy() GongstructIF {
	newInstance := new(SystemShape)
	systemshape.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (allocatedresourceshape *AllocatedResourceShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(allocatedresourceshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(allocatedresourceshape), uint64(stage.GetOrder(allocatedresourceshape)))
	return
}

func (allocatedsystemshape *AllocatedSystemShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(allocatedsystemshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(allocatedsystemshape), uint64(stage.GetOrder(allocatedsystemshape)))
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

func (diagramlayerstate *DiagramLayerState) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramlayerstate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagramlayerstate), uint64(stage.GetOrder(diagramlayerstate)))
	return
}

func (diagramstructure *DiagramStructure) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramstructure).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagramstructure), uint64(stage.GetOrder(diagramstructure)))
	return
}

func (externalpartshape *ExternalPartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(externalpartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(externalpartshape), uint64(stage.GetOrder(externalpartshape)))
	return
}

func (layerdefinition *LayerDefinition) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(layerdefinition).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(layerdefinition), uint64(stage.GetOrder(layerdefinition)))
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

func (notepartshape *NotePartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notepartshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notepartshape), uint64(stage.GetOrder(notepartshape)))
	return
}

func (noteportshape *NotePortShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteportshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(noteportshape), uint64(stage.GetOrder(noteportshape)))
	return
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(noteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(noteshape), uint64(stage.GetOrder(noteshape)))
	return
}

func (part *Part) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(part), uint64(stage.GetOrder(part)))
	return
}

func (partanchoredpath *PartAnchoredPath) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partanchoredpath).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partanchoredpath), uint64(stage.GetOrder(partanchoredpath)))
	return
}

func (partshape *PartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(partshape), uint64(stage.GetOrder(partshape)))
	return
}

func (port *Port) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(port).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(port), uint64(stage.GetOrder(port)))
	return
}

func (portshape *PortShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(portshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(portshape), uint64(stage.GetOrder(portshape)))
	return
}

func (resource *Resource) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(resource).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(resource), uint64(stage.GetOrder(resource)))
	return
}

func (semantictag *SemanticTag) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(semantictag).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(semantictag), uint64(stage.GetOrder(semantictag)))
	return
}

func (system *System) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(system).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(system), uint64(stage.GetOrder(system)))
	return
}

func (systemshape *SystemShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(systemshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(systemshape), uint64(stage.GetOrder(systemshape)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
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
		stage.AllocatedSystemShapes,
		stage.AllocatedSystemShape_stagedOrder,
		stage.AllocatedSystemShapes_reference,
		&stage.AllocatedSystemShapes_referenceOrder,
		stage.AllocatedSystemShapes_instance,
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
		stage.DiagramLayerStates,
		stage.DiagramLayerState_stagedOrder,
		stage.DiagramLayerStates_reference,
		&stage.DiagramLayerStates_referenceOrder,
		stage.DiagramLayerStates_instance,
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
		stage.DiagramStructures,
		stage.DiagramStructure_stagedOrder,
		stage.DiagramStructures_reference,
		&stage.DiagramStructures_referenceOrder,
		stage.DiagramStructures_instance,
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
		stage.ExternalPartShapes,
		stage.ExternalPartShape_stagedOrder,
		stage.ExternalPartShapes_reference,
		&stage.ExternalPartShapes_referenceOrder,
		stage.ExternalPartShapes_instance,
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
		stage.LayerDefinitions,
		stage.LayerDefinition_stagedOrder,
		stage.LayerDefinitions_reference,
		&stage.LayerDefinitions_referenceOrder,
		stage.LayerDefinitions_instance,
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
		stage.NotePartShapes,
		stage.NotePartShape_stagedOrder,
		stage.NotePartShapes_reference,
		&stage.NotePartShapes_referenceOrder,
		stage.NotePartShapes_instance,
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
		stage.NotePortShapes,
		stage.NotePortShape_stagedOrder,
		stage.NotePortShapes_reference,
		&stage.NotePortShapes_referenceOrder,
		stage.NotePortShapes_instance,
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
		stage.Parts,
		stage.Part_stagedOrder,
		stage.Parts_reference,
		&stage.Parts_referenceOrder,
		stage.Parts_instance,
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
		stage.PartAnchoredPaths,
		stage.PartAnchoredPath_stagedOrder,
		stage.PartAnchoredPaths_reference,
		&stage.PartAnchoredPaths_referenceOrder,
		stage.PartAnchoredPaths_instance,
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
		stage.PartShapes,
		stage.PartShape_stagedOrder,
		stage.PartShapes_reference,
		&stage.PartShapes_referenceOrder,
		stage.PartShapes_instance,
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
		stage.Ports,
		stage.Port_stagedOrder,
		stage.Ports_reference,
		&stage.Ports_referenceOrder,
		stage.Ports_instance,
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
		stage.PortShapes,
		stage.PortShape_stagedOrder,
		stage.PortShapes_reference,
		&stage.PortShapes_referenceOrder,
		stage.PortShapes_instance,
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
		stage.SemanticTags,
		stage.SemanticTag_stagedOrder,
		stage.SemanticTags_reference,
		&stage.SemanticTags_referenceOrder,
		stage.SemanticTags_instance,
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
		stage.Systems,
		stage.System_stagedOrder,
		stage.Systems_reference,
		&stage.Systems_referenceOrder,
		stage.Systems_instance,
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
		stage.SystemShapes,
		stage.SystemShape_stagedOrder,
		stage.SystemShapes_reference,
		&stage.SystemShapes_referenceOrder,
		stage.SystemShapes_instance,
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

	stage.DiagramLayerStates_reference = make(map[*DiagramLayerState]*DiagramLayerState)
	stage.DiagramLayerStates_referenceOrder = make(map[*DiagramLayerState]uint) // diff Unstage needs the reference order
	stage.DiagramLayerStates_instance = make(map[*DiagramLayerState]*DiagramLayerState)
	for instance := range stage.DiagramLayerStates {
		_copy := instance.GongCopy().(*DiagramLayerState)
		stage.DiagramLayerStates_reference[instance] = _copy
		stage.DiagramLayerStates_instance[_copy] = instance
		stage.DiagramLayerStates_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.LayerDefinitions_reference = make(map[*LayerDefinition]*LayerDefinition)
	stage.LayerDefinitions_referenceOrder = make(map[*LayerDefinition]uint) // diff Unstage needs the reference order
	stage.LayerDefinitions_instance = make(map[*LayerDefinition]*LayerDefinition)
	for instance := range stage.LayerDefinitions {
		_copy := instance.GongCopy().(*LayerDefinition)
		stage.LayerDefinitions_reference[instance] = _copy
		stage.LayerDefinitions_instance[_copy] = instance
		stage.LayerDefinitions_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.NotePartShapes_reference = make(map[*NotePartShape]*NotePartShape)
	stage.NotePartShapes_referenceOrder = make(map[*NotePartShape]uint) // diff Unstage needs the reference order
	stage.NotePartShapes_instance = make(map[*NotePartShape]*NotePartShape)
	for instance := range stage.NotePartShapes {
		_copy := instance.GongCopy().(*NotePartShape)
		stage.NotePartShapes_reference[instance] = _copy
		stage.NotePartShapes_instance[_copy] = instance
		stage.NotePartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.SemanticTags_reference = make(map[*SemanticTag]*SemanticTag)
	stage.SemanticTags_referenceOrder = make(map[*SemanticTag]uint) // diff Unstage needs the reference order
	stage.SemanticTags_instance = make(map[*SemanticTag]*SemanticTag)
	for instance := range stage.SemanticTags {
		_copy := instance.GongCopy().(*SemanticTag)
		stage.SemanticTags_reference[instance] = _copy
		stage.SemanticTags_instance[_copy] = instance
		stage.SemanticTags_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	for instance := range stage.DiagramLayerStates {
		reference := stage.DiagramLayerStates_reference[instance]
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

	for instance := range stage.LayerDefinitions {
		reference := stage.LayerDefinitions_reference[instance]
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

	for instance := range stage.NotePartShapes {
		reference := stage.NotePartShapes_reference[instance]
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

	for instance := range stage.SemanticTags {
		reference := stage.SemanticTags_reference[instance]
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

func (diagramlayerstate *DiagramLayerState) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramLayerState_stagedOrder[diagramlayerstate]; ok {
		return order
	}
	if order, ok := stage.DiagramLayerStates_referenceOrder[diagramlayerstate]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramLayerState was not staged and does not have a reference order", diagramlayerstate)
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

func (layerdefinition *LayerDefinition) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LayerDefinition_stagedOrder[layerdefinition]; ok {
		return order
	}
	if order, ok := stage.LayerDefinitions_referenceOrder[layerdefinition]; ok {
		return order
	} else {
		log.Printf("instance %p of type LayerDefinition was not staged and does not have a reference order", layerdefinition)
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

func (notepartshape *NotePartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.NotePartShape_stagedOrder[notepartshape]; ok {
		return order
	}
	if order, ok := stage.NotePartShapes_referenceOrder[notepartshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type NotePartShape was not staged and does not have a reference order", notepartshape)
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

func (semantictag *SemanticTag) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SemanticTag_stagedOrder[semantictag]; ok {
		return order
	}
	if order, ok := stage.SemanticTags_referenceOrder[semantictag]; ok {
		return order
	} else {
		log.Printf("instance %p of type SemanticTag was not staged and does not have a reference order", semantictag)
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

func (diagramlayerstate *DiagramLayerState) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramlayerstate.GongGetGongstructName(), diagramlayerstate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramlayerstate *DiagramLayerState) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramlayerstate.GongGetGongstructName(), diagramlayerstate.GongGetOrder(stage))
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

func (layerdefinition *LayerDefinition) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layerdefinition.GongGetGongstructName(), layerdefinition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layerdefinition *LayerDefinition) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layerdefinition.GongGetGongstructName(), layerdefinition.GongGetOrder(stage))
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

func (notepartshape *NotePartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notepartshape.GongGetGongstructName(), notepartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notepartshape *NotePartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notepartshape.GongGetGongstructName(), notepartshape.GongGetOrder(stage))
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

func (semantictag *SemanticTag) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", semantictag.GongGetGongstructName(), semantictag.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (semantictag *SemanticTag) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", semantictag.GongGetGongstructName(), semantictag.GongGetOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(allocatedresourceshape.Name))
	return
}

func (allocatedsystemshape *AllocatedSystemShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", allocatedsystemshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AllocatedSystemShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(allocatedsystemshape.Name))
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

func (diagramlayerstate *DiagramLayerState) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramlayerstate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramLayerState")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagramlayerstate.Name))
	return
}

func (diagramstructure *DiagramStructure) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramStructure")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagramstructure.Name))
	return
}

func (externalpartshape *ExternalPartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", externalpartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ExternalPartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(externalpartshape.Name))
	return
}

func (layerdefinition *LayerDefinition) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layerdefinition.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LayerDefinition")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(layerdefinition.Name))
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

func (notepartshape *NotePartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notepartshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NotePartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notepartshape.Name))
	return
}

func (noteportshape *NotePortShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteportshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NotePortShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(noteportshape.Name))
	return
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", noteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "NoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(noteshape.Name))
	return
}

func (part *Part) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(part.Name))
	return
}

func (partanchoredpath *PartAnchoredPath) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partanchoredpath.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartAnchoredPath")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partanchoredpath.Name))
	return
}

func (partshape *PartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(partshape.Name))
	return
}

func (port *Port) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", port.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Port")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(port.Name))
	return
}

func (portshape *PortShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", portshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PortShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(portshape.Name))
	return
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", resource.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Resource")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(resource.Name))
	return
}

func (semantictag *SemanticTag) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", semantictag.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SemanticTag")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(semantictag.Name))
	return
}

func (system *System) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(system.Name))
	return
}

func (systemshape *SystemShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SystemShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(systemshape.Name))
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

func (diagramlayerstate *DiagramLayerState) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramlayerstate.GongGetReferenceIdentifier(stage))
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

func (layerdefinition *LayerDefinition) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layerdefinition.GongGetReferenceIdentifier(stage))
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

func (notepartshape *NotePartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notepartshape.GongGetReferenceIdentifier(stage))
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

func (semantictag *SemanticTag) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", semantictag.GongGetReferenceIdentifier(stage))
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
