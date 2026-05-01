// generated code - do not edit
package models

// insertion point
func (inst *ControlFlow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *ControlFlowShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ControlFlowShapes":
			if _diagramprocess, ok := stage.DiagramProcess_ControlFlowShapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *DataFlow) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "DataFlowsWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap[inst]; ok {
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

func (inst *DataFlowShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "DataFlowShapes":
			if _diagramprocess, ok := stage.DiagramProcess_DataFlowShapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

func (inst *DiagramProcess) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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
		}
	}
	return
}

func (inst *Participant) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ParticipantWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_ParticipantWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramprocess.Name
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
		}
	}
	return
}

func (inst *ParticipantShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Process) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ProcesssWhoseNodeIsExpanded":
			if _diagramprocess, ok := stage.DiagramProcess_ProcesssWhoseNodeIsExpanded_reverseMap[inst]; ok {
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

func (inst *ProcessShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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

func (inst *Task) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

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
		}
	}
	return
}

func (inst *TaskShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "TaskShapes":
			if _diagramprocess, ok := stage.DiagramProcess_TaskShapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
			}
		}
	}
	return
}

// insertion point
func (inst *ControlFlow) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ControlFlowsWhoseNodeIsExpanded":
			res = stage.DiagramProcess_ControlFlowsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Participant":
		switch reverseField.Fieldname {
		case "ControlFlows":
			res = stage.Participant_ControlFlows_reverseMap[inst]
		}
	}
	return res
}

func (inst *ControlFlowShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ControlFlowShapes":
			res = stage.DiagramProcess_ControlFlowShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *DataFlow) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "DataFlowsWhoseNodeIsExpanded":
			res = stage.DiagramProcess_DataFlowsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootDataFlows":
			res = stage.Library_RootDataFlows_reverseMap[inst]
		case "DataFlowsWhoseNodeIsExpanded":
			res = stage.Library_DataFlowsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Process":
		switch reverseField.Fieldname {
		case "DataFlows":
			res = stage.Process_DataFlows_reverseMap[inst]
		}
	}
	return res
}

func (inst *DataFlowShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "DataFlowShapes":
			res = stage.DiagramProcess_DataFlowShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *DiagramProcess) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Process":
		switch reverseField.Fieldname {
		case "DiagramProcesss":
			res = stage.Process_DiagramProcesss_reverseMap[inst]
		case "DiagramProcessWhoseNodeIsExpanded":
			res = stage.Process_DiagramProcessWhoseNodeIsExpanded_reverseMap[inst]
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
		}
	}
	return res
}

func (inst *Participant) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ParticipantWhoseNodeIsExpanded":
			res = stage.DiagramProcess_ParticipantWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Process":
		switch reverseField.Fieldname {
		case "Participants":
			res = stage.Process_Participants_reverseMap[inst]
		case "ParticipantWhoseNodeIsExpanded":
			res = stage.Process_ParticipantWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *ParticipantShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "Participant_Shapes":
			res = stage.DiagramProcess_Participant_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Process) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ProcesssWhoseNodeIsExpanded":
			res = stage.DiagramProcess_ProcesssWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootProcesses":
			res = stage.Library_RootProcesses_reverseMap[inst]
		case "ProcesssWhoseNodeIsExpanded":
			res = stage.Library_ProcesssWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Process":
		switch reverseField.Fieldname {
		case "SubProcesses":
			res = stage.Process_SubProcesses_reverseMap[inst]
		}
	}
	return res
}

func (inst *ProcessShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "Process_Shapes":
			res = stage.DiagramProcess_Process_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Task) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "TasksWhoseNodeIsExpanded":
			res = stage.DiagramProcess_TasksWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Participant":
		switch reverseField.Fieldname {
		case "Tasks":
			res = stage.Participant_Tasks_reverseMap[inst]
		case "TaskWhoseOutControlFlowsNodeIsExpanded":
			res = stage.Participant_TaskWhoseOutControlFlowsNodeIsExpanded_reverseMap[inst]
		case "TaskWhoseInControlFlowsNodeIsExpanded":
			res = stage.Participant_TaskWhoseInControlFlowsNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *TaskShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "TaskShapes":
			res = stage.DiagramProcess_TaskShapes_reverseMap[inst]
		}
	}
	return res
}
