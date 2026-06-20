// generated code - do not edit
package models

// insertion point
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
		case "DiagramStructuresWhoseNodeIsExpanded":
			if _system, ok := stage.System_DiagramStructuresWhoseNodeIsExpanded_reverseMap[inst]; ok {
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

func (inst *Link) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "LinksWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_LinksWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "Links":
			if _system, ok := stage.System_Links_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "LinksWhoseNodeIsExpanded":
			if _system, ok := stage.System_LinksWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _system.Name
			}
		}
	}
	return
}

func (inst *LinkAssociationShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Link_Shapes":
			if _diagramstructure, ok := stage.DiagramStructure_Link_Shapes_reverseMap[inst]; ok {
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
		case "PartsWhoseNodeIsExpanded":
			if _diagramstructure, ok := stage.DiagramStructure_PartsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramstructure.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "Parts":
			if _system, ok := stage.System_Parts_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "PartsWhoseNodeIsExpanded":
			if _system, ok := stage.System_PartsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _system.Name
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

func (inst *System) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootSystems":
			if _library, ok := stage.Library_RootSystems_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "SystemsWhoseNodeIsExpanded":
			if _library, ok := stage.Library_SystemsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "SubSystems":
			if _system, ok := stage.System_SubSystems_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "SubSystemsWhoseNodeIsExpanded":
			if _system, ok := stage.System_SubSystemsWhoseNodeIsExpanded_reverseMap[inst]; ok {
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
func (inst *DiagramStructure) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "System":
		switch reverseField.Fieldname {
		case "DiagramStructures":
			res = stage.System_DiagramStructures_reverseMap[inst]
		case "DiagramStructuresWhoseNodeIsExpanded":
			res = stage.System_DiagramStructuresWhoseNodeIsExpanded_reverseMap[inst]
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

func (inst *Link) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "LinksWhoseNodeIsExpanded":
			res = stage.DiagramStructure_LinksWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "Links":
			res = stage.System_Links_reverseMap[inst]
		case "LinksWhoseNodeIsExpanded":
			res = stage.System_LinksWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *LinkAssociationShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramStructure":
		switch reverseField.Fieldname {
		case "Link_Shapes":
			res = stage.DiagramStructure_Link_Shapes_reverseMap[inst]
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
		case "PartsWhoseNodeIsExpanded":
			res = stage.DiagramStructure_PartsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "Parts":
			res = stage.System_Parts_reverseMap[inst]
		case "PartsWhoseNodeIsExpanded":
			res = stage.System_PartsWhoseNodeIsExpanded_reverseMap[inst]
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

func (inst *System) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootSystems":
			res = stage.Library_RootSystems_reverseMap[inst]
		case "SystemsWhoseNodeIsExpanded":
			res = stage.Library_SystemsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "SubSystems":
			res = stage.System_SubSystems_reverseMap[inst]
		case "SubSystemsWhoseNodeIsExpanded":
			res = stage.System_SubSystemsWhoseNodeIsExpanded_reverseMap[inst]
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
