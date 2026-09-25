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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AllocatedResourceShapes)

	res = __gong__appendInstances(res, stage.AllocatedSystemShapes)

	res = __gong__appendInstances(res, stage.ControlFlows)

	res = __gong__appendInstances(res, stage.ControlFlowShapes)

	res = __gong__appendInstances(res, stage.Datas)

	res = __gong__appendInstances(res, stage.DataFlows)

	res = __gong__appendInstances(res, stage.DataFlowShapes)

	res = __gong__appendInstances(res, stage.DataShapes)

	res = __gong__appendInstances(res, stage.DiagramLayerStates)

	res = __gong__appendInstances(res, stage.DiagramStructures)

	res = __gong__appendInstances(res, stage.ExternalPartShapes)

	res = __gong__appendInstances(res, stage.LayerDefinitions)

	res = __gong__appendInstances(res, stage.Librarys)

	res = __gong__appendInstances(res, stage.Notes)

	res = __gong__appendInstances(res, stage.NotePartShapes)

	res = __gong__appendInstances(res, stage.NotePortShapes)

	res = __gong__appendInstances(res, stage.NoteShapes)

	res = __gong__appendInstances(res, stage.Parts)

	res = __gong__appendInstances(res, stage.PartAnchoredPaths)

	res = __gong__appendInstances(res, stage.PartShapes)

	res = __gong__appendInstances(res, stage.Ports)

	res = __gong__appendInstances(res, stage.PortShapes)

	res = __gong__appendInstances(res, stage.Resources)

	res = __gong__appendInstances(res, stage.SemanticTags)

	res = __gong__appendInstances(res, stage.Systems)

	res = __gong__appendInstances(res, stage.SystemShapes)

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
func (allocatedresourceshape *AllocatedResourceShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, allocatedresourceshape)
}

func (allocatedsystemshape *AllocatedSystemShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, allocatedsystemshape)
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

func (diagramlayerstate *DiagramLayerState) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagramlayerstate)
}

func (diagramstructure *DiagramStructure) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagramstructure)
}

func (externalpartshape *ExternalPartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, externalpartshape)
}

func (layerdefinition *LayerDefinition) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, layerdefinition)
}

func (library *Library) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, library)
}

func (note *Note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note)
}

func (notepartshape *NotePartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notepartshape)
}

func (noteportshape *NotePortShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteportshape)
}

func (noteshape *NoteShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, noteshape)
}

func (part *Part) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part)
}

func (partanchoredpath *PartAnchoredPath) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partanchoredpath)
}

func (partshape *PartShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, partshape)
}

func (port *Port) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, port)
}

func (portshape *PortShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, portshape)
}

func (resource *Resource) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, resource)
}

func (semantictag *SemanticTag) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, semantictag)
}

func (system *System) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, system)
}

func (systemshape *SystemShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, systemshape)
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
	__gong__computeReferencePass1(stage, stage.AllocatedResourceShapes, &stage.AllocatedResourceShapes_reference, &stage.AllocatedResourceShapes_referenceOrder, &stage.AllocatedResourceShapes_instance)

	__gong__computeReferencePass1(stage, stage.AllocatedSystemShapes, &stage.AllocatedSystemShapes_reference, &stage.AllocatedSystemShapes_referenceOrder, &stage.AllocatedSystemShapes_instance)

	__gong__computeReferencePass1(stage, stage.ControlFlows, &stage.ControlFlows_reference, &stage.ControlFlows_referenceOrder, &stage.ControlFlows_instance)

	__gong__computeReferencePass1(stage, stage.ControlFlowShapes, &stage.ControlFlowShapes_reference, &stage.ControlFlowShapes_referenceOrder, &stage.ControlFlowShapes_instance)

	__gong__computeReferencePass1(stage, stage.Datas, &stage.Datas_reference, &stage.Datas_referenceOrder, &stage.Datas_instance)

	__gong__computeReferencePass1(stage, stage.DataFlows, &stage.DataFlows_reference, &stage.DataFlows_referenceOrder, &stage.DataFlows_instance)

	__gong__computeReferencePass1(stage, stage.DataFlowShapes, &stage.DataFlowShapes_reference, &stage.DataFlowShapes_referenceOrder, &stage.DataFlowShapes_instance)

	__gong__computeReferencePass1(stage, stage.DataShapes, &stage.DataShapes_reference, &stage.DataShapes_referenceOrder, &stage.DataShapes_instance)

	__gong__computeReferencePass1(stage, stage.DiagramLayerStates, &stage.DiagramLayerStates_reference, &stage.DiagramLayerStates_referenceOrder, &stage.DiagramLayerStates_instance)

	__gong__computeReferencePass1(stage, stage.DiagramStructures, &stage.DiagramStructures_reference, &stage.DiagramStructures_referenceOrder, &stage.DiagramStructures_instance)

	__gong__computeReferencePass1(stage, stage.ExternalPartShapes, &stage.ExternalPartShapes_reference, &stage.ExternalPartShapes_referenceOrder, &stage.ExternalPartShapes_instance)

	__gong__computeReferencePass1(stage, stage.LayerDefinitions, &stage.LayerDefinitions_reference, &stage.LayerDefinitions_referenceOrder, &stage.LayerDefinitions_instance)

	__gong__computeReferencePass1(stage, stage.Librarys, &stage.Librarys_reference, &stage.Librarys_referenceOrder, &stage.Librarys_instance)

	__gong__computeReferencePass1(stage, stage.Notes, &stage.Notes_reference, &stage.Notes_referenceOrder, &stage.Notes_instance)

	__gong__computeReferencePass1(stage, stage.NotePartShapes, &stage.NotePartShapes_reference, &stage.NotePartShapes_referenceOrder, &stage.NotePartShapes_instance)

	__gong__computeReferencePass1(stage, stage.NotePortShapes, &stage.NotePortShapes_reference, &stage.NotePortShapes_referenceOrder, &stage.NotePortShapes_instance)

	__gong__computeReferencePass1(stage, stage.NoteShapes, &stage.NoteShapes_reference, &stage.NoteShapes_referenceOrder, &stage.NoteShapes_instance)

	__gong__computeReferencePass1(stage, stage.Parts, &stage.Parts_reference, &stage.Parts_referenceOrder, &stage.Parts_instance)

	__gong__computeReferencePass1(stage, stage.PartAnchoredPaths, &stage.PartAnchoredPaths_reference, &stage.PartAnchoredPaths_referenceOrder, &stage.PartAnchoredPaths_instance)

	__gong__computeReferencePass1(stage, stage.PartShapes, &stage.PartShapes_reference, &stage.PartShapes_referenceOrder, &stage.PartShapes_instance)

	__gong__computeReferencePass1(stage, stage.Ports, &stage.Ports_reference, &stage.Ports_referenceOrder, &stage.Ports_instance)

	__gong__computeReferencePass1(stage, stage.PortShapes, &stage.PortShapes_reference, &stage.PortShapes_referenceOrder, &stage.PortShapes_instance)

	__gong__computeReferencePass1(stage, stage.Resources, &stage.Resources_reference, &stage.Resources_referenceOrder, &stage.Resources_instance)

	__gong__computeReferencePass1(stage, stage.SemanticTags, &stage.SemanticTags_reference, &stage.SemanticTags_referenceOrder, &stage.SemanticTags_instance)

	__gong__computeReferencePass1(stage, stage.Systems, &stage.Systems_reference, &stage.Systems_referenceOrder, &stage.Systems_instance)

	__gong__computeReferencePass1(stage, stage.SystemShapes, &stage.SystemShapes_reference, &stage.SystemShapes_referenceOrder, &stage.SystemShapes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AllocatedResourceShapes, stage.AllocatedResourceShapes_reference, stage)

	__gong__computeReferencePass2(stage.AllocatedSystemShapes, stage.AllocatedSystemShapes_reference, stage)

	__gong__computeReferencePass2(stage.ControlFlows, stage.ControlFlows_reference, stage)

	__gong__computeReferencePass2(stage.ControlFlowShapes, stage.ControlFlowShapes_reference, stage)

	__gong__computeReferencePass2(stage.Datas, stage.Datas_reference, stage)

	__gong__computeReferencePass2(stage.DataFlows, stage.DataFlows_reference, stage)

	__gong__computeReferencePass2(stage.DataFlowShapes, stage.DataFlowShapes_reference, stage)

	__gong__computeReferencePass2(stage.DataShapes, stage.DataShapes_reference, stage)

	__gong__computeReferencePass2(stage.DiagramLayerStates, stage.DiagramLayerStates_reference, stage)

	__gong__computeReferencePass2(stage.DiagramStructures, stage.DiagramStructures_reference, stage)

	__gong__computeReferencePass2(stage.ExternalPartShapes, stage.ExternalPartShapes_reference, stage)

	__gong__computeReferencePass2(stage.LayerDefinitions, stage.LayerDefinitions_reference, stage)

	__gong__computeReferencePass2(stage.Librarys, stage.Librarys_reference, stage)

	__gong__computeReferencePass2(stage.Notes, stage.Notes_reference, stage)

	__gong__computeReferencePass2(stage.NotePartShapes, stage.NotePartShapes_reference, stage)

	__gong__computeReferencePass2(stage.NotePortShapes, stage.NotePortShapes_reference, stage)

	__gong__computeReferencePass2(stage.NoteShapes, stage.NoteShapes_reference, stage)

	__gong__computeReferencePass2(stage.Parts, stage.Parts_reference, stage)

	__gong__computeReferencePass2(stage.PartAnchoredPaths, stage.PartAnchoredPaths_reference, stage)

	__gong__computeReferencePass2(stage.PartShapes, stage.PartShapes_reference, stage)

	__gong__computeReferencePass2(stage.Ports, stage.Ports_reference, stage)

	__gong__computeReferencePass2(stage.PortShapes, stage.PortShapes_reference, stage)

	__gong__computeReferencePass2(stage.Resources, stage.Resources_reference, stage)

	__gong__computeReferencePass2(stage.SemanticTags, stage.SemanticTags_reference, stage)

	__gong__computeReferencePass2(stage.Systems, stage.Systems_reference, stage)

	__gong__computeReferencePass2(stage.SystemShapes, stage.SystemShapes_reference, stage)

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
	return __gong__getOrder(stage.AllocatedResourceShape_stagedOrder, stage.AllocatedResourceShapes_referenceOrder, allocatedresourceshape, "AllocatedResourceShape")
}

func (allocatedsystemshape *AllocatedSystemShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AllocatedSystemShape_stagedOrder, stage.AllocatedSystemShapes_referenceOrder, allocatedsystemshape, "AllocatedSystemShape")
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

func (diagramlayerstate *DiagramLayerState) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DiagramLayerState_stagedOrder, stage.DiagramLayerStates_referenceOrder, diagramlayerstate, "DiagramLayerState")
}

func (diagramstructure *DiagramStructure) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DiagramStructure_stagedOrder, stage.DiagramStructures_referenceOrder, diagramstructure, "DiagramStructure")
}

func (externalpartshape *ExternalPartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ExternalPartShape_stagedOrder, stage.ExternalPartShapes_referenceOrder, externalpartshape, "ExternalPartShape")
}

func (layerdefinition *LayerDefinition) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LayerDefinition_stagedOrder, stage.LayerDefinitions_referenceOrder, layerdefinition, "LayerDefinition")
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Library_stagedOrder, stage.Librarys_referenceOrder, library, "Library")
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_stagedOrder, stage.Notes_referenceOrder, note, "Note")
}

func (notepartshape *NotePartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePartShape_stagedOrder, stage.NotePartShapes_referenceOrder, notepartshape, "NotePartShape")
}

func (noteportshape *NotePortShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NotePortShape_stagedOrder, stage.NotePortShapes_referenceOrder, noteportshape, "NotePortShape")
}

func (noteshape *NoteShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.NoteShape_stagedOrder, stage.NoteShapes_referenceOrder, noteshape, "NoteShape")
}

func (part *Part) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_stagedOrder, stage.Parts_referenceOrder, part, "Part")
}

func (partanchoredpath *PartAnchoredPath) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartAnchoredPath_stagedOrder, stage.PartAnchoredPaths_referenceOrder, partanchoredpath, "PartAnchoredPath")
}

func (partshape *PartShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PartShape_stagedOrder, stage.PartShapes_referenceOrder, partshape, "PartShape")
}

func (port *Port) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Port_stagedOrder, stage.Ports_referenceOrder, port, "Port")
}

func (portshape *PortShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.PortShape_stagedOrder, stage.PortShapes_referenceOrder, portshape, "PortShape")
}

func (resource *Resource) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Resource_stagedOrder, stage.Resources_referenceOrder, resource, "Resource")
}

func (semantictag *SemanticTag) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SemanticTag_stagedOrder, stage.SemanticTags_referenceOrder, semantictag, "SemanticTag")
}

func (system *System) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.System_stagedOrder, stage.Systems_referenceOrder, system, "System")
}

func (systemshape *SystemShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SystemShape_stagedOrder, stage.SystemShapes_referenceOrder, systemshape, "SystemShape")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (allocatedresourceshape *AllocatedResourceShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(allocatedresourceshape, allocatedresourceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (allocatedresourceshape *AllocatedResourceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return allocatedresourceshape.GongGetIdentifier(stage)
}

func (allocatedsystemshape *AllocatedSystemShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(allocatedsystemshape, allocatedsystemshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (allocatedsystemshape *AllocatedSystemShape) GongGetReferenceIdentifier(stage *Stage) string {
	return allocatedsystemshape.GongGetIdentifier(stage)
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

func (diagramlayerstate *DiagramLayerState) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagramlayerstate, diagramlayerstate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramlayerstate *DiagramLayerState) GongGetReferenceIdentifier(stage *Stage) string {
	return diagramlayerstate.GongGetIdentifier(stage)
}

func (diagramstructure *DiagramStructure) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagramstructure, diagramstructure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramstructure *DiagramStructure) GongGetReferenceIdentifier(stage *Stage) string {
	return diagramstructure.GongGetIdentifier(stage)
}

func (externalpartshape *ExternalPartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(externalpartshape, externalpartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (externalpartshape *ExternalPartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return externalpartshape.GongGetIdentifier(stage)
}

func (layerdefinition *LayerDefinition) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(layerdefinition, layerdefinition.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layerdefinition *LayerDefinition) GongGetReferenceIdentifier(stage *Stage) string {
	return layerdefinition.GongGetIdentifier(stage)
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

func (notepartshape *NotePartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notepartshape, notepartshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notepartshape *NotePartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return notepartshape.GongGetIdentifier(stage)
}

func (noteportshape *NotePortShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteportshape, noteportshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteportshape *NotePortShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteportshape.GongGetIdentifier(stage)
}

func (noteshape *NoteShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(noteshape, noteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (noteshape *NoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return noteshape.GongGetIdentifier(stage)
}

func (part *Part) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part, part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part *Part) GongGetReferenceIdentifier(stage *Stage) string {
	return part.GongGetIdentifier(stage)
}

func (partanchoredpath *PartAnchoredPath) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partanchoredpath, partanchoredpath.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partanchoredpath *PartAnchoredPath) GongGetReferenceIdentifier(stage *Stage) string {
	return partanchoredpath.GongGetIdentifier(stage)
}

func (partshape *PartShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(partshape, partshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partshape *PartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return partshape.GongGetIdentifier(stage)
}

func (port *Port) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(port, port.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (port *Port) GongGetReferenceIdentifier(stage *Stage) string {
	return port.GongGetIdentifier(stage)
}

func (portshape *PortShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(portshape, portshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (portshape *PortShape) GongGetReferenceIdentifier(stage *Stage) string {
	return portshape.GongGetIdentifier(stage)
}

func (resource *Resource) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(resource, resource.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (resource *Resource) GongGetReferenceIdentifier(stage *Stage) string {
	return resource.GongGetIdentifier(stage)
}

func (semantictag *SemanticTag) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(semantictag, semantictag.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (semantictag *SemanticTag) GongGetReferenceIdentifier(stage *Stage) string {
	return semantictag.GongGetIdentifier(stage)
}

func (system *System) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(system, system.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system *System) GongGetReferenceIdentifier(stage *Stage) string {
	return system.GongGetIdentifier(stage)
}

func (systemshape *SystemShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(systemshape, systemshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (systemshape *SystemShape) GongGetReferenceIdentifier(stage *Stage) string {
	return systemshape.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(allocatedresourceshape.GongGetIdentifier(stage), "AllocatedResourceShape", allocatedresourceshape.Name)
}

func (allocatedsystemshape *AllocatedSystemShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(allocatedsystemshape.GongGetIdentifier(stage), "AllocatedSystemShape", allocatedsystemshape.Name)
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

func (diagramlayerstate *DiagramLayerState) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagramlayerstate.GongGetIdentifier(stage), "DiagramLayerState", diagramlayerstate.Name)
}

func (diagramstructure *DiagramStructure) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagramstructure.GongGetIdentifier(stage), "DiagramStructure", diagramstructure.Name)
}

func (externalpartshape *ExternalPartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(externalpartshape.GongGetIdentifier(stage), "ExternalPartShape", externalpartshape.Name)
}

func (layerdefinition *LayerDefinition) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(layerdefinition.GongGetIdentifier(stage), "LayerDefinition", layerdefinition.Name)
}

func (library *Library) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(library.GongGetIdentifier(stage), "Library", library.Name)
}

func (note *Note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note.GongGetIdentifier(stage), "Note", note.Name)
}

func (notepartshape *NotePartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notepartshape.GongGetIdentifier(stage), "NotePartShape", notepartshape.Name)
}

func (noteportshape *NotePortShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteportshape.GongGetIdentifier(stage), "NotePortShape", noteportshape.Name)
}

func (noteshape *NoteShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(noteshape.GongGetIdentifier(stage), "NoteShape", noteshape.Name)
}

func (part *Part) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part.GongGetIdentifier(stage), "Part", part.Name)
}

func (partanchoredpath *PartAnchoredPath) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partanchoredpath.GongGetIdentifier(stage), "PartAnchoredPath", partanchoredpath.Name)
}

func (partshape *PartShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(partshape.GongGetIdentifier(stage), "PartShape", partshape.Name)
}

func (port *Port) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(port.GongGetIdentifier(stage), "Port", port.Name)
}

func (portshape *PortShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(portshape.GongGetIdentifier(stage), "PortShape", portshape.Name)
}

func (resource *Resource) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(resource.GongGetIdentifier(stage), "Resource", resource.Name)
}

func (semantictag *SemanticTag) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(semantictag.GongGetIdentifier(stage), "SemanticTag", semantictag.Name)
}

func (system *System) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(system.GongGetIdentifier(stage), "System", system.Name)
}

func (systemshape *SystemShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(systemshape.GongGetIdentifier(stage), "SystemShape", systemshape.Name)
}

// insertion point for unstaging
func (allocatedresourceshape *AllocatedResourceShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(allocatedresourceshape.GongGetReferenceIdentifier(stage))
}

func (allocatedsystemshape *AllocatedSystemShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(allocatedsystemshape.GongGetReferenceIdentifier(stage))
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

func (diagramlayerstate *DiagramLayerState) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagramlayerstate.GongGetReferenceIdentifier(stage))
}

func (diagramstructure *DiagramStructure) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagramstructure.GongGetReferenceIdentifier(stage))
}

func (externalpartshape *ExternalPartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(externalpartshape.GongGetReferenceIdentifier(stage))
}

func (layerdefinition *LayerDefinition) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(layerdefinition.GongGetReferenceIdentifier(stage))
}

func (library *Library) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(library.GongGetReferenceIdentifier(stage))
}

func (note *Note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note.GongGetReferenceIdentifier(stage))
}

func (notepartshape *NotePartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notepartshape.GongGetReferenceIdentifier(stage))
}

func (noteportshape *NotePortShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteportshape.GongGetReferenceIdentifier(stage))
}

func (noteshape *NoteShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(noteshape.GongGetReferenceIdentifier(stage))
}

func (part *Part) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part.GongGetReferenceIdentifier(stage))
}

func (partanchoredpath *PartAnchoredPath) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partanchoredpath.GongGetReferenceIdentifier(stage))
}

func (partshape *PartShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(partshape.GongGetReferenceIdentifier(stage))
}

func (port *Port) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(port.GongGetReferenceIdentifier(stage))
}

func (portshape *PortShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(portshape.GongGetReferenceIdentifier(stage))
}

func (resource *Resource) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(resource.GongGetReferenceIdentifier(stage))
}

func (semantictag *SemanticTag) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(semantictag.GongGetReferenceIdentifier(stage))
}

func (system *System) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(system.GongGetReferenceIdentifier(stage))
}

func (systemshape *SystemShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(systemshape.GongGetReferenceIdentifier(stage))
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
