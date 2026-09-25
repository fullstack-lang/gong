// generated code - do not edit
package models

// insertion point
func (inst *AllocatedResourceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *AllocatedSystemShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *ControlFlow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *ControlFlowShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Data) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *DataFlow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *DataFlowShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *DataShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *DiagramLayerState) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DiagramStructure) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *ExternalPartShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *LayerDefinition) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *NotePartShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *NotePortShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Part) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
	case "SemanticTag":
		switch reverseField.Fieldname {
		case "Parts":
			if _semantictag, ok := stage.SemanticTag_Parts_reverseMap[inst]; ok {
				res = _semantictag.Name
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

func (inst *PartAnchoredPath) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *PartShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Port) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *PortShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Resource) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *SemanticTag) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "LayerDefinition":
		switch reverseField.Fieldname {
		case "Query":
			if _layerdefinition, ok := stage.LayerDefinition_Query_reverseMap[inst]; ok {
				res = _layerdefinition.Name
			}
		}
	}
	return
}

func (inst *System) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *SystemShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
