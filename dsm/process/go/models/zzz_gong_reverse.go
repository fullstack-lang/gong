// generated code - do not edit
package models

// insertion point
func (inst *AllocatedProcessShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "AllocatedProcessShapes":
			if _diagramprocess, ok := stage.DiagramProcess_AllocatedProcessShapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *AllocatedResourceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "AllocatedResourceShapes":
			if _diagramprocess, ok := stage.DiagramProcess_AllocatedResourceShapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *ControlFlow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ControlFlowsWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_ControlFlowsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	case "Participant":
		switch reverseField.Fieldname {
		case "ControlFlows":
			if _participant, ok := stage.Participant_ControlFlows_reverseMap[inst]; ok {
				res = _participant.Name
			}
		}
	}
	return
}

func (inst *ControlFlowShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ControlFlow_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_ControlFlow_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
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
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "DatasWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_DatasWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
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
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "DataFlowsWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		case "DataFlowsWhoseDataNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_DataFlowsWhoseDataNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
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
	case "Process":
		switch reverseField.Fieldname {
		case "DataFlows":
			if _process, ok := stage.Process_DataFlows_reverseMap[inst]; ok {
				res = _process.Name
			}
		}
	}
	return
}

func (inst *DataFlowShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "DataFlow_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_DataFlow_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *DataShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "Data_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_Data_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *DiagramProcess) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Process":
		switch reverseField.Fieldname {
		case "DiagramProcesss":
			if _process, ok := stage.Process_DiagramProcesss_reverseMap[inst]; ok {
				res = _process.Name
			}
		case "DiagramProcessWhoseNodeIsExpanded":
			if _process, ok := stage.Process_DiagramProcessWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _process.Name
			}
		}
	}
	return
}

func (inst *ExternalParticipantShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ExternalParticipant_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_ExternalParticipant_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
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
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_NotesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
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

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_Note_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *NoteTaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "NoteTaskShapes":
			if _diagramprocess, ok := stage.DiagramProcess_NoteTaskShapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *Participant) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ParticipantWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_ParticipantWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		case "ExternalParticipantWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_ExternalParticipantWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		case "ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		case "ExternalParticipantsWhoseInDataFlowsNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_ExternalParticipantsWhoseInDataFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "ParticipantsWhoseNodeIsExpanded":
			if _library, ok := stage.Library_ParticipantsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Process":
		switch reverseField.Fieldname {
		case "Participants":
			if _process, ok := stage.Process_Participants_reverseMap[inst]; ok {
				res = _process.Name
			}
		case "ParticipantWhoseNodeIsExpanded":
			if _process, ok := stage.Process_ParticipantWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _process.Name
			}
		case "ExternalParticipants":
			if _process, ok := stage.Process_ExternalParticipants_reverseMap[inst]; ok {
				res = _process.Name
			}
		case "ExternalParticipantWhoseNodeIsExpanded":
			if _process, ok := stage.Process_ExternalParticipantWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _process.Name
			}
		}
	}
	return
}

func (inst *ParticipantShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "Participant_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_Participant_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *Process) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ProcesssWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_ProcesssWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		case "AllocatedProcessesWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_AllocatedProcessesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootProcesses":
			if _library, ok := stage.Library_RootProcesses_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "ProcesssWhoseNodeIsExpanded":
			if _library, ok := stage.Library_ProcesssWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "Participant":
		switch reverseField.Fieldname {
		case "Processes":
			if _participant, ok := stage.Participant_Processes_reverseMap[inst]; ok {
				res = _participant.Name
			}
		}
	case "Process":
		switch reverseField.Fieldname {
		case "SubProcesses":
			if _process, ok := stage.Process_SubProcesses_reverseMap[inst]; ok {
				res = _process.Name
			}
		}
	}
	return
}

func (inst *ProcessShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "Process_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_Process_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *Resource) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "AllocatedResourcesWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_AllocatedResourcesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
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
	case "Participant":
		switch reverseField.Fieldname {
		case "Resources":
			if _participant, ok := stage.Participant_Resources_reverseMap[inst]; ok {
				res = _participant.Name
			}
		}
	}
	return
}

func (inst *Task) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "TasksWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_TasksWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Tasks":
			if _note, ok := stage.Note_Tasks_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "Participant":
		switch reverseField.Fieldname {
		case "Tasks":
			if _participant, ok := stage.Participant_Tasks_reverseMap[inst]; ok {
				res = _participant.Name
			}
		case "TaskWhoseOutControlFlowsNodeIsExpanded":
			if _participant, ok := stage.Participant_TaskWhoseOutControlFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _participant.Name
			}
		case "TaskWhoseInControlFlowsNodeIsExpanded":
			if _participant, ok := stage.Participant_TaskWhoseInControlFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _participant.Name
			}
		case "TaskWhoseOutDataFlowsNodeIsExpanded":
			if _participant, ok := stage.Participant_TaskWhoseOutDataFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _participant.Name
			}
		case "TaskWhoseInDataFlowsNodeIsExpanded":
			if _participant, ok := stage.Participant_TaskWhoseInDataFlowsNodeIsExpanded_reverseMap[inst]; ok {
				res = _participant.Name
			}
		}
	}
	return
}

func (inst *TaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "Task_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_Task_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}
