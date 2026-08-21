// generated code - do not edit
package models

// insertion point
func (inst *CompareAnalysis) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootCompareAnalysis":
			if _library, ok := stage.Library_RootCompareAnalysis_reverseMap[inst]; ok {
				res = _library.Name
			}
		case "CompareAnalysisWhoseNodeIsExpanded":
			if _library, ok := stage.Library_CompareAnalysisWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *Complexity) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "ComplexitysWhoseNodeIsExpanded":
			if _diagramflossequation, ok := stage.DiagramFlossEquation_ComplexitysWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramflossequation.Name
			}
		}
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
	case "Note":
		switch reverseField.Fieldname {
		case "Complexities":
			if _note, ok := stage.Note_Complexities_reverseMap[inst]; ok {
				res = _note.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "Complexities":
			if _system, ok := stage.System_Complexities_reverseMap[inst]; ok {
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

func (inst *DiagramFlossEquation) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CompareAnalysis":
		switch reverseField.Fieldname {
		case "DiagramFlossEquations":
			if _compareanalysis, ok := stage.CompareAnalysis_DiagramFlossEquations_reverseMap[inst]; ok {
				res = _compareanalysis.Name
			}
		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			if _compareanalysis, ok := stage.CompareAnalysis_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _compareanalysis.Name
			}
		}
	case "System":
		switch reverseField.Fieldname {
		case "DiagramFlossEquations":
			if _system, ok := stage.System_DiagramFlossEquations_reverseMap[inst]; ok {
				res = _system.Name
			}
		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			if _system, ok := stage.System_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap[inst]; ok {
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
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "EffortsWhoseNodeIsExpanded":
			if _diagramflossequation, ok := stage.DiagramFlossEquation_EffortsWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramflossequation.Name
			}
		}
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
	case "Note":
		switch reverseField.Fieldname {
		case "Efforts":
			if _note, ok := stage.Note_Efforts_reverseMap[inst]; ok {
				res = _note.Name
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

func (inst *Note) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			if _diagramflossequation, ok := stage.DiagramFlossEquation_NotesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramflossequation.Name
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

func (inst *NoteComplexityShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "NoteComplexityShapes":
			if _diagramflossequation, ok := stage.DiagramFlossEquation_NoteComplexityShapes_reverseMap[inst]; ok {
				res = _diagramflossequation.Name
			}
		}
	}
	return
}

func (inst *NoteEffortShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "NoteEffortShapes":
			if _diagramflossequation, ok := stage.DiagramFlossEquation_NoteEffortShapes_reverseMap[inst]; ok {
				res = _diagramflossequation.Name
			}
		}
	}
	return
}

func (inst *NotePerformanceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "NotePerformanceShapes":
			if _diagramflossequation, ok := stage.DiagramFlossEquation_NotePerformanceShapes_reverseMap[inst]; ok {
				res = _diagramflossequation.Name
			}
		}
	}
	return
}

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			if _diagramflossequation, ok := stage.DiagramFlossEquation_Note_Shapes_reverseMap[inst]; ok {
				res = _diagramflossequation.Name
			}
		}
	}
	return
}

func (inst *Performance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "PerformancesWhoseNodeIsExpanded":
			if _diagramflossequation, ok := stage.DiagramFlossEquation_PerformancesWhoseNodeIsExpanded_reverseMap[inst]; ok {
				res = _diagramflossequation.Name
			}
		}
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
	case "Note":
		switch reverseField.Fieldname {
		case "Performances":
			if _note, ok := stage.Note_Performances_reverseMap[inst]; ok {
				res = _note.Name
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
		case "SubSystemes":
			if _system, ok := stage.System_SubSystemes_reverseMap[inst]; ok {
				res = _system.Name
			}
		}
	}
	return
}

// insertion point
func (inst *CompareAnalysis) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "RootCompareAnalysis":
			res = stage.Library_RootCompareAnalysis_reverseMap[inst]
		case "CompareAnalysisWhoseNodeIsExpanded":
			res = stage.Library_CompareAnalysisWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *Complexity) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "ComplexitysWhoseNodeIsExpanded":
			res = stage.DiagramFlossEquation_ComplexitysWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootComplexitys":
			res = stage.Library_RootComplexitys_reverseMap[inst]
		case "ComplexitysWhoseNodeIsExpanded":
			res = stage.Library_ComplexitysWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Complexities":
			res = stage.Note_Complexities_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "Complexities":
			res = stage.System_Complexities_reverseMap[inst]
		case "ComplexitysWhoseNodeIsExpanded":
			res = stage.System_ComplexitysWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *DiagramFlossEquation) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "CompareAnalysis":
		switch reverseField.Fieldname {
		case "DiagramFlossEquations":
			res = stage.CompareAnalysis_DiagramFlossEquations_reverseMap[inst]
		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			res = stage.CompareAnalysis_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "System":
		switch reverseField.Fieldname {
		case "DiagramFlossEquations":
			res = stage.System_DiagramFlossEquations_reverseMap[inst]
		case "DiagramFlossEquationsWhoseNodeIsExpanded":
			res = stage.System_DiagramFlossEquationsWhoseNodeIsExpanded_reverseMap[inst]
		}
	}
	return res
}

func (inst *Effort) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "EffortsWhoseNodeIsExpanded":
			res = stage.DiagramFlossEquation_EffortsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootEfforts":
			res = stage.Library_RootEfforts_reverseMap[inst]
		case "EffortsWhoseNodeIsExpanded":
			res = stage.Library_EffortsWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Efforts":
			res = stage.Note_Efforts_reverseMap[inst]
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

func (inst *Note) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "NotesWhoseNodeIsExpanded":
			res = stage.DiagramFlossEquation_NotesWhoseNodeIsExpanded_reverseMap[inst]
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

func (inst *NoteComplexityShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "NoteComplexityShapes":
			res = stage.DiagramFlossEquation_NoteComplexityShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteEffortShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "NoteEffortShapes":
			res = stage.DiagramFlossEquation_NoteEffortShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NotePerformanceShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "NotePerformanceShapes":
			res = stage.DiagramFlossEquation_NotePerformanceShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *NoteShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "Note_Shapes":
			res = stage.DiagramFlossEquation_Note_Shapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Performance) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "DiagramFlossEquation":
		switch reverseField.Fieldname {
		case "PerformancesWhoseNodeIsExpanded":
			res = stage.DiagramFlossEquation_PerformancesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Library":
		switch reverseField.Fieldname {
		case "RootPerformances":
			res = stage.Library_RootPerformances_reverseMap[inst]
		case "PerformancesWhoseNodeIsExpanded":
			res = stage.Library_PerformancesWhoseNodeIsExpanded_reverseMap[inst]
		}
	case "Note":
		switch reverseField.Fieldname {
		case "Performances":
			res = stage.Note_Performances_reverseMap[inst]
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
	case "Library":
		switch reverseField.Fieldname {
		case "RootSystems":
			res = stage.Library_RootSystems_reverseMap[inst]
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
