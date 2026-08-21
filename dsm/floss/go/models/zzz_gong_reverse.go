// generated code - do not edit
package models

// insertion point
func (inst *DiagramFloss) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "System":
		switch reverseField.Fieldname {
		case "DiagramFlosses":
			if _system, ok := stage.System_DiagramFlosses_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "DiagramFlossWhoseNodeIsExpanded":
			if _system, ok := stage.System_DiagramFlossWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _system.Name
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

func (inst *System) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFloss":
		switch reverseField.Fieldname {
		case "SystemsWhoseNodeIsExpanded":
			if _diagramfloss, ok := stage.DiagramFloss_SystemsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramfloss.Name
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
	case "DiagramFloss":
		switch reverseField.Fieldname {
		case "System_Shapes":
			if _diagramfloss, ok := stage.DiagramFloss_System_Shapes_reverseMap[inst]; ok {
				res = _diagramfloss.Name
			}
		}
	}
	return
}

// insertion point
func (inst *DiagramFloss) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "System":
		switch reverseField.Fieldname {
		case "DiagramFlosses":
			res = stage.System_DiagramFlosses_reverseMap[inst]
		case "DiagramFlossWhoseNodeIsExpanded":
			res = stage.System_DiagramFlossWhoseNodeIsExpanded_reverseMap[inst]
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

func (inst *System) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFloss":
		switch reverseField.Fieldname {
		case "SystemsWhoseNodeIsExpanded":
			res = stage.DiagramFloss_SystemsWhoseNodeIsExpanded_reverseMap[inst]
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
	case "DiagramFloss":
		switch reverseField.Fieldname {
		case "System_Shapes":
			res = stage.DiagramFloss_System_Shapes_reverseMap[inst]
		}
	}
	return res
}
