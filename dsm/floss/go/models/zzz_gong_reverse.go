// generated code - do not edit
package models

// insertion point
func (inst *Complexity) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootComplexitys":
			if _library, ok := stage.Library_RootComplexitys_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "ComplexitysWhoseNodeIsExpanded":
			if _library, ok := stage.Library_ComplexitysWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "Complexitys":
			if _system, ok := stage.System_Complexitys_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "ComplexitysWhoseNodeIsExpanded":
			if _system, ok := stage.System_ComplexitysWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _system.Name
			}
		}
	}
	return
}

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

func (inst *Effort) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootEfforts":
			if _library, ok := stage.Library_RootEfforts_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "EffortsWhoseNodeIsExpanded":
			if _library, ok := stage.Library_EffortsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "Efforts":
			if _system, ok := stage.System_Efforts_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "EffortsWhoseNodeIsExpanded":
			if _system, ok := stage.System_EffortsWhoseNodeIsExpanded_reverseMap[inst]; ok {
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

func (inst *Performance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootPerformances":
			if _library, ok := stage.Library_RootPerformances_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "PerformancesWhoseNodeIsExpanded":
			if _library, ok := stage.Library_PerformancesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "Performances":
			if _system, ok := stage.System_Performances_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "PerformancesWhoseNodeIsExpanded":
			if _system, ok := stage.System_PerformancesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _system.Name
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
func (inst *Complexity) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootComplexitys":
			res = stage.Library_RootComplexitys_reverseMap[inst]
		case "ComplexitysWhoseNodeIsExpanded":
			res = stage.Library_ComplexitysWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "Complexitys":
			res = stage.System_Complexitys_reverseMap[inst]
		case "ComplexitysWhoseNodeIsExpanded":
			res = stage.System_ComplexitysWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

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

func (inst *Effort) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootEfforts":
			res = stage.Library_RootEfforts_reverseMap[inst]
		case "EffortsWhoseNodeIsExpanded":
			res = stage.Library_EffortsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "Efforts":
			res = stage.System_Efforts_reverseMap[inst]
		case "EffortsWhoseNodeIsExpanded":
			res = stage.System_EffortsWhoseNodeIsExpanded_reverseMap[inst]
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

func (inst *Performance) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootPerformances":
			res = stage.Library_RootPerformances_reverseMap[inst]
		case "PerformancesWhoseNodeIsExpanded":
			res = stage.Library_PerformancesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "Performances":
			res = stage.System_Performances_reverseMap[inst]
		case "PerformancesWhoseNodeIsExpanded":
			res = stage.System_PerformancesWhoseNodeIsExpanded_reverseMap[inst]
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
