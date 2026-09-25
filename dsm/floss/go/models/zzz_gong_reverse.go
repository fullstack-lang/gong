// generated code - do not edit
package models

// insertion point
func (inst *CompareAnalysis) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Complexity) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *DiagramFlossEquation) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Effort) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *NoteComplexityShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *NoteEffortShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *NotePerformanceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *NoteShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *Performance) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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

func (inst *System) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

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
		}
	}
	return
}
