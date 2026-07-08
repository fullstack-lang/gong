// generated code - do not edit
package models

// insertion point
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

func (inst *Plant) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Plants":
			if _library, ok := stage.Library_Plants_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *PlantDiagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Plant":
		switch reverseField.Fieldname {
		case "PlantDiagramsWhoseNodeIsExpanded":
			if _plant, ok := stage.Plant_PlantDiagramsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _plant.Name
			}
		case "PlantDiagrams":
			if _plant, ok := stage.Plant_PlantDiagrams_reverseMap[inst]; ok {
				res = _plant.Name
			}
		}
	}
	return
}

// insertion point
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

func (inst *Plant) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Plants":
			res = stage.Library_Plants_reverseMap[inst]
		}
	}
	return res
}

func (inst *PlantDiagram) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Plant":
		switch reverseField.Fieldname {
		case "PlantDiagramsWhoseNodeIsExpanded":
			res = stage.Plant_PlantDiagramsWhoseNodeIsExpanded_reverseMap[inst]
		case "PlantDiagrams":
			res = stage.Plant_PlantDiagrams_reverseMap[inst]
		}
	}
	return res
}
