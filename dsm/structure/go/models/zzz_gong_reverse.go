// generated code - do not edit
package models

// insertion point
func (inst *AllocatedResourceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "AllocatedResourceShapes":
			if _diagramstructure, ok := stage.DiagramStructure_AllocatedResourceShapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *AllocatedSystemShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "AllocatedSystemShapes":
			if _diagramstructure, ok := stage.DiagramStructure_AllocatedSystemShapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *ControlFlow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "ControlFlowsWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_ControlFlowsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "Part":
		switch reverseField.Fieldname {
		case "ControlFlows":
			if _part, ok := stage.Part_ControlFlows_reverseMap[inst]; ok {
				res = _part.Name
			}
		}
	}
	return
}

func (inst *ControlFlowShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "ControlFlow_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_ControlFlow_Shapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *Data) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DataFlow":
		switch reverseField.Fieldname {
		case "Datas":
			if _dataflow, ok := stage.DataFlow_Datas_reverseMap[inst]; ok {
				res = _dataflow.Name
			}
		}
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "DatasWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_DatasWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootDatas":
			if _library, ok := stage.Library_RootDatas_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "DatasWhoseNodeIsExpanded":
			if _library, ok := stage.Library_DatasWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *DataFlow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "DataFlowsWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_DataFlowsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		case "DataFlowsWhoseDataNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_DataFlowsWhoseDataNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootDataFlows":
			if _library, ok := stage.Library_RootDataFlows_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "DataFlowsWhoseNodeIsExpanded":
			if _library, ok := stage.Library_DataFlowsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "DataFlows":
			if _system, ok := stage.System_DataFlows_reverseMap[inst]; ok {
				res = _system.Name
			}
		}
	}
	return
}

func (inst *DataFlowShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "DataFlow_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_DataFlow_Shapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *DataShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Data_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_Data_Shapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *DiagramStructure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "System":
		switch reverseField.Fieldname {
		case "DiagramStructures":
			if _system, ok := stage.System_DiagramStructures_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "DiagramStructureWhoseNodeIsExpanded":
			if _system, ok := stage.System_DiagramStructureWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _system.Name
			}
		}
	}
	return
}

func (inst *ExternalPartShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "ExternalPart_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_ExternalPart_Shapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			if _library, ok := stage.Library_SubLibraries_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "SubLibrariesWhoseNodeIsExpanded":
			if _library, ok := stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_NotesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootNotes":
			if _library, ok := stage.Library_RootNotes_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "NotesWhoseNodeIsExpanded":
			if _library, ok := stage.Library_NotesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *NotePartShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "NotePartShapes":
			if _diagramstructure, ok := stage.DiagramStructure_NotePartShapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *NotePortShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "NotePortShapes":
			if _diagramstructure, ok := stage.DiagramStructure_NotePortShapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_Note_Shapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *Part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "PartWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_PartWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		case "ExternalPartWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_ExternalPartWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_ExternalPartsWhoseOutDataFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_ExternalPartsWhoseInDataFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "PartsWhoseNodeIsExpanded":
			if _library, ok := stage.Library_PartsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Parts":
			if _note, ok := stage.Note_Parts_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "Parts":
			if _system, ok := stage.System_Parts_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "PartWhoseNodeIsExpanded":
			if _system, ok := stage.System_PartWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "ExternalParts":
			if _system, ok := stage.System_ExternalParts_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "ExternalPartWhoseNodeIsExpanded":
			if _system, ok := stage.System_ExternalPartWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _system.Name
			}
		}
	}
	return
}

func (inst *PartAnchoredPath) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Part":
		switch reverseField.Fieldname {
		case "PartAnchoredPath":
			if _part, ok := stage.Part_PartAnchoredPath_reverseMap[inst]; ok {
				res = _part.Name
			}
		}
	}
	return
}

func (inst *PartShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Part_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_Part_Shapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *Port) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "PortsWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_PortsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Ports":
			if _note, ok := stage.Note_Ports_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Part":
		switch reverseField.Fieldname {
		case "Ports":
			if _part, ok := stage.Part_Ports_reverseMap[inst]; ok {
				res = _part.Name
			}
		case "PortWhoseOutControlFlowsNodeIsExpanded":
			if _part, ok := stage.Part_PortWhoseOutControlFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _part.Name
			}
		case "PortWhoseInControlFlowsNodeIsExpanded":
			if _part, ok := stage.Part_PortWhoseInControlFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _part.Name
			}
		case "PortWhoseOutDataFlowsNodeIsExpanded":
			if _part, ok := stage.Part_PortWhoseOutDataFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _part.Name
			}
		case "PortWhoseInDataFlowsNodeIsExpanded":
			if _part, ok := stage.Part_PortWhoseInDataFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _part.Name
			}
		}
	}
	return
}

func (inst *PortShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Port_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_Port_Shapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

func (inst *Resource) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "AllocatedResourcesWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_AllocatedResourcesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootResources":
			if _library, ok := stage.Library_RootResources_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "ResourcesWhoseNodeIsExpanded":
			if _library, ok := stage.Library_ResourcesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *System) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "SystemsWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_SystemsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		case "AllocatedSystemesWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_AllocatedSystemesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootSystemes":
			if _library, ok := stage.Library_RootSystemes_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "SystemsWhoseNodeIsExpanded":
			if _library, ok := stage.Library_SystemsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "SubSystemes":
			if _system, ok := stage.System_SubSystemes_reverseMap[inst]; ok {
				res = _system.Name
			}
		}
	}
	return
}

func (inst *SystemShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "System_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_System_Shapes_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	}
	return
}

// insertion point
func (inst *AllocatedResourceShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "AllocatedResourceShapes":
			res = stage.DiagramStructure_AllocatedResourceShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *AllocatedSystemShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "AllocatedSystemShapes":
			res = stage.DiagramStructure_AllocatedSystemShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *ControlFlow) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "ControlFlowsWhoseNodeIsExpanded":
			res = stage.DiagramStructure_ControlFlowsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Part":
		switch reverseField.Fieldname {
		case "ControlFlows":
			res = stage.Part_ControlFlows_reverseMap[inst]
		}
	}
	return res
}

func (inst *ControlFlowShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "ControlFlow_Shapes":
			res = stage.DiagramStructure_ControlFlow_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Data) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DataFlow":
		switch reverseField.Fieldname {
		case "Datas":
			res = stage.DataFlow_Datas_reverseMap[inst]
		}
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "DatasWhoseNodeIsExpanded":
			res = stage.DiagramStructure_DatasWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootDatas":
			res = stage.Library_RootDatas_reverseMap[inst]
		case "DatasWhoseNodeIsExpanded":
			res = stage.Library_DatasWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *DataFlow) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "DataFlowsWhoseNodeIsExpanded":
			res = stage.DiagramStructure_DataFlowsWhoseNodeIsExpanded_reverseMap[inst]
		case "DataFlowsWhoseDataNodeIsExpanded":
			res = stage.DiagramStructure_DataFlowsWhoseDataNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootDataFlows":
			res = stage.Library_RootDataFlows_reverseMap[inst]
		case "DataFlowsWhoseNodeIsExpanded":
			res = stage.Library_DataFlowsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "DataFlows":
			res = stage.System_DataFlows_reverseMap[inst]
		}
	}
	return res
}

func (inst *DataFlowShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "DataFlow_Shapes":
			res = stage.DiagramStructure_DataFlow_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *DataShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Data_Shapes":
			res = stage.DiagramStructure_Data_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *DiagramStructure) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "System":
		switch reverseField.Fieldname {
		case "DiagramStructures":
			res = stage.System_DiagramStructures_reverseMap[inst]
		case "DiagramStructureWhoseNodeIsExpanded":
			res = stage.System_DiagramStructureWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *ExternalPartShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "ExternalPart_Shapes":
			res = stage.DiagramStructure_ExternalPart_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Library) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			res = stage.Library_SubLibraries_reverseMap[inst]
		case "SubLibrariesWhoseNodeIsExpanded":
			res = stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *Note) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			res = stage.DiagramStructure_NotesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootNotes":
			res = stage.Library_RootNotes_reverseMap[inst]
		case "NotesWhoseNodeIsExpanded":
			res = stage.Library_NotesWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *NotePartShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "NotePartShapes":
			res = stage.DiagramStructure_NotePartShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NotePortShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "NotePortShapes":
			res = stage.DiagramStructure_NotePortShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			res = stage.DiagramStructure_Note_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Part) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "PartWhoseNodeIsExpanded":
			res = stage.DiagramStructure_PartWhoseNodeIsExpanded_reverseMap[inst]
		case "ExternalPartWhoseNodeIsExpanded":
			res = stage.DiagramStructure_ExternalPartWhoseNodeIsExpanded_reverseMap[inst]
		case "ExternalPartsWhoseOutDataFlowsNodeIsExpanded":
			res = stage.DiagramStructure_ExternalPartsWhoseOutDataFlowsNodeIsExpanded_reverseMap[inst]
		case "ExternalPartsWhoseInDataFlowsNodeIsExpanded":
			res = stage.DiagramStructure_ExternalPartsWhoseInDataFlowsNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "PartsWhoseNodeIsExpanded":
			res = stage.Library_PartsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Parts":
			res = stage.Note_Parts_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "Parts":
			res = stage.System_Parts_reverseMap[inst]
		case "PartWhoseNodeIsExpanded":
			res = stage.System_PartWhoseNodeIsExpanded_reverseMap[inst]
		case "ExternalParts":
			res = stage.System_ExternalParts_reverseMap[inst]
		case "ExternalPartWhoseNodeIsExpanded":
			res = stage.System_ExternalPartWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *PartAnchoredPath) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Part":
		switch reverseField.Fieldname {
		case "PartAnchoredPath":
			res = stage.Part_PartAnchoredPath_reverseMap[inst]
		}
	}
	return res
}

func (inst *PartShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Part_Shapes":
			res = stage.DiagramStructure_Part_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Port) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "PortsWhoseNodeIsExpanded":
			res = stage.DiagramStructure_PortsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Ports":
			res = stage.Note_Ports_reverseMap[inst]
		}
	case "Part":
		switch reverseField.Fieldname {
		case "Ports":
			res = stage.Part_Ports_reverseMap[inst]
		case "PortWhoseOutControlFlowsNodeIsExpanded":
			res = stage.Part_PortWhoseOutControlFlowsNodeIsExpanded_reverseMap[inst]
		case "PortWhoseInControlFlowsNodeIsExpanded":
			res = stage.Part_PortWhoseInControlFlowsNodeIsExpanded_reverseMap[inst]
		case "PortWhoseOutDataFlowsNodeIsExpanded":
			res = stage.Part_PortWhoseOutDataFlowsNodeIsExpanded_reverseMap[inst]
		case "PortWhoseInDataFlowsNodeIsExpanded":
			res = stage.Part_PortWhoseInDataFlowsNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *PortShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Port_Shapes":
			res = stage.DiagramStructure_Port_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Resource) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "AllocatedResourcesWhoseNodeIsExpanded":
			res = stage.DiagramStructure_AllocatedResourcesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootResources":
			res = stage.Library_RootResources_reverseMap[inst]
		case "ResourcesWhoseNodeIsExpanded":
			res = stage.Library_ResourcesWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *System) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "SystemsWhoseNodeIsExpanded":
			res = stage.DiagramStructure_SystemsWhoseNodeIsExpanded_reverseMap[inst]
		case "AllocatedSystemesWhoseNodeIsExpanded":
			res = stage.DiagramStructure_AllocatedSystemesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootSystemes":
			res = stage.Library_RootSystemes_reverseMap[inst]
		case "SystemsWhoseNodeIsExpanded":
			res = stage.Library_SystemsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "SubSystemes":
			res = stage.System_SubSystemes_reverseMap[inst]
		}
	}
	return res
}

func (inst *SystemShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "System_Shapes":
			res = stage.DiagramStructure_System_Shapes_reverseMap[inst]
		}
	}
	return res
}
