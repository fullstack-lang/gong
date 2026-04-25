// generated code - do not edit
package models

// insertion point
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

func (inst *ProcessCompositionShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ProcessComposition_Shapes":
			if _diagramprocess, ok := stage.DiagramProcess_ProcessComposition_Shapes_reverseMap[inst]; ok {
				res = _diagramprocess.Name
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

// insertion point
func (inst *DiagramProcess) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Process":
		switch reverseField.Fieldname {
		case "DiagramProcesss":
			res = stage.Process_DiagramProcesss_reverseMap[inst]
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

func (inst *ProcessCompositionShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramProcess":
		switch reverseField.Fieldname {
		case "ProcessComposition_Shapes":
			res = stage.DiagramProcess_ProcessComposition_Shapes_reverseMap[inst]
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
