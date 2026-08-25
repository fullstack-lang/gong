// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *CompareAnalysis:
		ok = stage.IsStagedCompareAnalysis(target)

	case *Complexity:
		ok = stage.IsStagedComplexity(target)

	case *DiagramFlossEquation:
		ok = stage.IsStagedDiagramFlossEquation(target)

	case *Effort:
		ok = stage.IsStagedEffort(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteComplexityShape:
		ok = stage.IsStagedNoteComplexityShape(target)

	case *NoteEffortShape:
		ok = stage.IsStagedNoteEffortShape(target)

	case *NotePerformanceShape:
		ok = stage.IsStagedNotePerformanceShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *Performance:
		ok = stage.IsStagedPerformance(target)

	case *System:
		ok = stage.IsStagedSystem(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *CompareAnalysis:
		ok = stage.IsStagedCompareAnalysis(target)

	case *Complexity:
		ok = stage.IsStagedComplexity(target)

	case *DiagramFlossEquation:
		ok = stage.IsStagedDiagramFlossEquation(target)

	case *Effort:
		ok = stage.IsStagedEffort(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteComplexityShape:
		ok = stage.IsStagedNoteComplexityShape(target)

	case *NoteEffortShape:
		ok = stage.IsStagedNoteEffortShape(target)

	case *NotePerformanceShape:
		ok = stage.IsStagedNotePerformanceShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *Performance:
		ok = stage.IsStagedPerformance(target)

	case *System:
		ok = stage.IsStagedSystem(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedCompareAnalysis(compareanalysis *CompareAnalysis) (ok bool) {

	_, ok = stage.CompareAnalysiss[compareanalysis]

	return
}

func (stage *Stage) IsStagedComplexity(complexity *Complexity) (ok bool) {

	_, ok = stage.Complexitys[complexity]

	return
}

func (stage *Stage) IsStagedDiagramFlossEquation(diagramflossequation *DiagramFlossEquation) (ok bool) {

	_, ok = stage.DiagramFlossEquations[diagramflossequation]

	return
}

func (stage *Stage) IsStagedEffort(effort *Effort) (ok bool) {

	_, ok = stage.Efforts[effort]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNoteComplexityShape(notecomplexityshape *NoteComplexityShape) (ok bool) {

	_, ok = stage.NoteComplexityShapes[notecomplexityshape]

	return
}

func (stage *Stage) IsStagedNoteEffortShape(noteeffortshape *NoteEffortShape) (ok bool) {

	_, ok = stage.NoteEffortShapes[noteeffortshape]

	return
}

func (stage *Stage) IsStagedNotePerformanceShape(noteperformanceshape *NotePerformanceShape) (ok bool) {

	_, ok = stage.NotePerformanceShapes[noteperformanceshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedPerformance(performance *Performance) (ok bool) {

	_, ok = stage.Performances[performance]

	return
}

func (stage *Stage) IsStagedSystem(system *System) (ok bool) {

	_, ok = stage.Systems[system]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *CompareAnalysis:
		stage.StageBranchCompareAnalysis(target)

	case *Complexity:
		stage.StageBranchComplexity(target)

	case *DiagramFlossEquation:
		stage.StageBranchDiagramFlossEquation(target)

	case *Effort:
		stage.StageBranchEffort(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Note:
		stage.StageBranchNote(target)

	case *NoteComplexityShape:
		stage.StageBranchNoteComplexityShape(target)

	case *NoteEffortShape:
		stage.StageBranchNoteEffortShape(target)

	case *NotePerformanceShape:
		stage.StageBranchNotePerformanceShape(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *Performance:
		stage.StageBranchPerformance(target)

	case *System:
		stage.StageBranchSystem(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchCompareAnalysis(compareanalysis *CompareAnalysis) {

	// check if instance is already staged
	if IsStaged(stage, compareanalysis) {
		return
	}

	compareanalysis.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if compareanalysis.FromSystem != nil {
		StageBranch(stage, compareanalysis.FromSystem)
	}
	if compareanalysis.ToSystem != nil {
		StageBranch(stage, compareanalysis.ToSystem)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramflossequation := range compareanalysis.DiagramFlossEquations {
		StageBranch(stage, _diagramflossequation)
	}
	for _, _diagramflossequation := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
		StageBranch(stage, _diagramflossequation)
	}

}

func (stage *Stage) StageBranchComplexity(complexity *Complexity) {

	// check if instance is already staged
	if IsStaged(stage, complexity) {
		return
	}

	complexity.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDiagramFlossEquation(diagramflossequation *DiagramFlossEquation) {

	// check if instance is already staged
	if IsStaged(stage, diagramflossequation) {
		return
	}

	diagramflossequation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshape := range diagramflossequation.Note_Shapes {
		StageBranch(stage, _noteshape)
	}
	for _, _notecomplexityshape := range diagramflossequation.NoteComplexityShapes {
		StageBranch(stage, _notecomplexityshape)
	}
	for _, _noteperformanceshape := range diagramflossequation.NotePerformanceShapes {
		StageBranch(stage, _noteperformanceshape)
	}
	for _, _noteeffortshape := range diagramflossequation.NoteEffortShapes {
		StageBranch(stage, _noteeffortshape)
	}
	for _, _note := range diagramflossequation.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}
	for _, _complexity := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
		StageBranch(stage, _complexity)
	}
	for _, _performance := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
		StageBranch(stage, _performance)
	}
	for _, _effort := range diagramflossequation.EffortsWhoseNodeIsExpanded {
		StageBranch(stage, _effort)
	}

}

func (stage *Stage) StageBranchEffort(effort *Effort) {

	// check if instance is already staged
	if IsStaged(stage, effort) {
		return
	}

	effort.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if IsStaged(stage, library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		StageBranch(stage, _library)
	}
	for _, _system := range library.RootSystems {
		StageBranch(stage, _system)
	}
	for _, _complexity := range library.RootComplexitys {
		StageBranch(stage, _complexity)
	}
	for _, _performance := range library.RootPerformances {
		StageBranch(stage, _performance)
	}
	for _, _effort := range library.RootEfforts {
		StageBranch(stage, _effort)
	}
	for _, _compareanalysis := range library.RootCompareAnalysis {
		StageBranch(stage, _compareanalysis)
	}
	for _, _note := range library.RootNotes {
		StageBranch(stage, _note)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		StageBranch(stage, _library)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		StageBranch(stage, _system)
	}
	for _, _complexity := range library.ComplexitysWhoseNodeIsExpanded {
		StageBranch(stage, _complexity)
	}
	for _, _performance := range library.PerformancesWhoseNodeIsExpanded {
		StageBranch(stage, _performance)
	}
	for _, _effort := range library.EffortsWhoseNodeIsExpanded {
		StageBranch(stage, _effort)
	}
	for _, _compareanalysis := range library.CompareAnalysisWhoseNodeIsExpanded {
		StageBranch(stage, _compareanalysis)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}

}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if IsStaged(stage, note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range note.Complexities {
		StageBranch(stage, _complexity)
	}
	for _, _performance := range note.Performances {
		StageBranch(stage, _performance)
	}
	for _, _effort := range note.Efforts {
		StageBranch(stage, _effort)
	}

}

func (stage *Stage) StageBranchNoteComplexityShape(notecomplexityshape *NoteComplexityShape) {

	// check if instance is already staged
	if IsStaged(stage, notecomplexityshape) {
		return
	}

	notecomplexityshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notecomplexityshape.Note != nil {
		StageBranch(stage, notecomplexityshape.Note)
	}
	if notecomplexityshape.Complexity != nil {
		StageBranch(stage, notecomplexityshape.Complexity)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteEffortShape(noteeffortshape *NoteEffortShape) {

	// check if instance is already staged
	if IsStaged(stage, noteeffortshape) {
		return
	}

	noteeffortshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteeffortshape.Note != nil {
		StageBranch(stage, noteeffortshape.Note)
	}
	if noteeffortshape.Effort != nil {
		StageBranch(stage, noteeffortshape.Effort)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNotePerformanceShape(noteperformanceshape *NotePerformanceShape) {

	// check if instance is already staged
	if IsStaged(stage, noteperformanceshape) {
		return
	}

	noteperformanceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteperformanceshape.Note != nil {
		StageBranch(stage, noteperformanceshape.Note)
	}
	if noteperformanceshape.Performance != nil {
		StageBranch(stage, noteperformanceshape.Performance)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if IsStaged(stage, noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		StageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPerformance(performance *Performance) {

	// check if instance is already staged
	if IsStaged(stage, performance) {
		return
	}

	performance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSystem(system *System) {

	// check if instance is already staged
	if IsStaged(stage, system) {
		return
	}

	system.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range system.Complexities {
		StageBranch(stage, _complexity)
	}
	for _, _performance := range system.Performances {
		StageBranch(stage, _performance)
	}
	for _, _effort := range system.Efforts {
		StageBranch(stage, _effort)
	}
	for _, _system := range system.SubSystems {
		StageBranch(stage, _system)
	}
	for _, _diagramflossequation := range system.DiagramFlossEquations {
		StageBranch(stage, _diagramflossequation)
	}
	for _, _diagramflossequation := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
		StageBranch(stage, _diagramflossequation)
	}
	for _, _complexity := range system.ComplexitysWhoseNodeIsExpanded {
		StageBranch(stage, _complexity)
	}
	for _, _performance := range system.PerformancesWhoseNodeIsExpanded {
		StageBranch(stage, _performance)
	}
	for _, _effort := range system.EffortsWhoseNodeIsExpanded {
		StageBranch(stage, _effort)
	}

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *CompareAnalysis:
		toT := CopyBranchCompareAnalysis(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Complexity:
		toT := CopyBranchComplexity(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramFlossEquation:
		toT := CopyBranchDiagramFlossEquation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Effort:
		toT := CopyBranchEffort(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := CopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteComplexityShape:
		toT := CopyBranchNoteComplexityShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteEffortShape:
		toT := CopyBranchNoteEffortShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NotePerformanceShape:
		toT := CopyBranchNotePerformanceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := CopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Performance:
		toT := CopyBranchPerformance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System:
		toT := CopyBranchSystem(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchCompareAnalysis(mapOrigCopy map[any]any, compareanalysisFrom *CompareAnalysis) (compareanalysisTo *CompareAnalysis) {

	// compareanalysisFrom has already been copied
	if _compareanalysisTo, ok := mapOrigCopy[compareanalysisFrom]; ok {
		compareanalysisTo = _compareanalysisTo.(*CompareAnalysis)
		return
	}

	compareanalysisTo = new(CompareAnalysis)
	mapOrigCopy[compareanalysisFrom] = compareanalysisTo
	compareanalysisFrom.CopyBasicFields(compareanalysisTo)

	//insertion point for the staging of instances referenced by pointers
	if compareanalysisFrom.FromSystem != nil {
		compareanalysisTo.FromSystem = CopyBranchSystem(mapOrigCopy, compareanalysisFrom.FromSystem)
	}
	if compareanalysisFrom.ToSystem != nil {
		compareanalysisTo.ToSystem = CopyBranchSystem(mapOrigCopy, compareanalysisFrom.ToSystem)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramflossequation := range compareanalysisFrom.DiagramFlossEquations {
		compareanalysisTo.DiagramFlossEquations = append(compareanalysisTo.DiagramFlossEquations, CopyBranchDiagramFlossEquation(mapOrigCopy, _diagramflossequation))
	}
	for _, _diagramflossequation := range compareanalysisFrom.DiagramFlossEquationsWhoseNodeIsExpanded {
		compareanalysisTo.DiagramFlossEquationsWhoseNodeIsExpanded = append(compareanalysisTo.DiagramFlossEquationsWhoseNodeIsExpanded, CopyBranchDiagramFlossEquation(mapOrigCopy, _diagramflossequation))
	}

	return
}

func CopyBranchComplexity(mapOrigCopy map[any]any, complexityFrom *Complexity) (complexityTo *Complexity) {

	// complexityFrom has already been copied
	if _complexityTo, ok := mapOrigCopy[complexityFrom]; ok {
		complexityTo = _complexityTo.(*Complexity)
		return
	}

	complexityTo = new(Complexity)
	mapOrigCopy[complexityFrom] = complexityTo
	complexityFrom.CopyBasicFields(complexityTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDiagramFlossEquation(mapOrigCopy map[any]any, diagramflossequationFrom *DiagramFlossEquation) (diagramflossequationTo *DiagramFlossEquation) {

	// diagramflossequationFrom has already been copied
	if _diagramflossequationTo, ok := mapOrigCopy[diagramflossequationFrom]; ok {
		diagramflossequationTo = _diagramflossequationTo.(*DiagramFlossEquation)
		return
	}

	diagramflossequationTo = new(DiagramFlossEquation)
	mapOrigCopy[diagramflossequationFrom] = diagramflossequationTo
	diagramflossequationFrom.CopyBasicFields(diagramflossequationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshape := range diagramflossequationFrom.Note_Shapes {
		diagramflossequationTo.Note_Shapes = append(diagramflossequationTo.Note_Shapes, CopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _notecomplexityshape := range diagramflossequationFrom.NoteComplexityShapes {
		diagramflossequationTo.NoteComplexityShapes = append(diagramflossequationTo.NoteComplexityShapes, CopyBranchNoteComplexityShape(mapOrigCopy, _notecomplexityshape))
	}
	for _, _noteperformanceshape := range diagramflossequationFrom.NotePerformanceShapes {
		diagramflossequationTo.NotePerformanceShapes = append(diagramflossequationTo.NotePerformanceShapes, CopyBranchNotePerformanceShape(mapOrigCopy, _noteperformanceshape))
	}
	for _, _noteeffortshape := range diagramflossequationFrom.NoteEffortShapes {
		diagramflossequationTo.NoteEffortShapes = append(diagramflossequationTo.NoteEffortShapes, CopyBranchNoteEffortShape(mapOrigCopy, _noteeffortshape))
	}
	for _, _note := range diagramflossequationFrom.NotesWhoseNodeIsExpanded {
		diagramflossequationTo.NotesWhoseNodeIsExpanded = append(diagramflossequationTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _complexity := range diagramflossequationFrom.ComplexitysWhoseNodeIsExpanded {
		diagramflossequationTo.ComplexitysWhoseNodeIsExpanded = append(diagramflossequationTo.ComplexitysWhoseNodeIsExpanded, CopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range diagramflossequationFrom.PerformancesWhoseNodeIsExpanded {
		diagramflossequationTo.PerformancesWhoseNodeIsExpanded = append(diagramflossequationTo.PerformancesWhoseNodeIsExpanded, CopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range diagramflossequationFrom.EffortsWhoseNodeIsExpanded {
		diagramflossequationTo.EffortsWhoseNodeIsExpanded = append(diagramflossequationTo.EffortsWhoseNodeIsExpanded, CopyBranchEffort(mapOrigCopy, _effort))
	}

	return
}

func CopyBranchEffort(mapOrigCopy map[any]any, effortFrom *Effort) (effortTo *Effort) {

	// effortFrom has already been copied
	if _effortTo, ok := mapOrigCopy[effortFrom]; ok {
		effortTo = _effortTo.(*Effort)
		return
	}

	effortTo = new(Effort)
	mapOrigCopy[effortFrom] = effortTo
	effortFrom.CopyBasicFields(effortTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.CopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _system := range libraryFrom.RootSystems {
		libraryTo.RootSystems = append(libraryTo.RootSystems, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _complexity := range libraryFrom.RootComplexitys {
		libraryTo.RootComplexitys = append(libraryTo.RootComplexitys, CopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range libraryFrom.RootPerformances {
		libraryTo.RootPerformances = append(libraryTo.RootPerformances, CopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range libraryFrom.RootEfforts {
		libraryTo.RootEfforts = append(libraryTo.RootEfforts, CopyBranchEffort(mapOrigCopy, _effort))
	}
	for _, _compareanalysis := range libraryFrom.RootCompareAnalysis {
		libraryTo.RootCompareAnalysis = append(libraryTo.RootCompareAnalysis, CopyBranchCompareAnalysis(mapOrigCopy, _compareanalysis))
	}
	for _, _note := range libraryFrom.RootNotes {
		libraryTo.RootNotes = append(libraryTo.RootNotes, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _system := range libraryFrom.SystemsWhoseNodeIsExpanded {
		libraryTo.SystemsWhoseNodeIsExpanded = append(libraryTo.SystemsWhoseNodeIsExpanded, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _complexity := range libraryFrom.ComplexitysWhoseNodeIsExpanded {
		libraryTo.ComplexitysWhoseNodeIsExpanded = append(libraryTo.ComplexitysWhoseNodeIsExpanded, CopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range libraryFrom.PerformancesWhoseNodeIsExpanded {
		libraryTo.PerformancesWhoseNodeIsExpanded = append(libraryTo.PerformancesWhoseNodeIsExpanded, CopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range libraryFrom.EffortsWhoseNodeIsExpanded {
		libraryTo.EffortsWhoseNodeIsExpanded = append(libraryTo.EffortsWhoseNodeIsExpanded, CopyBranchEffort(mapOrigCopy, _effort))
	}
	for _, _compareanalysis := range libraryFrom.CompareAnalysisWhoseNodeIsExpanded {
		libraryTo.CompareAnalysisWhoseNodeIsExpanded = append(libraryTo.CompareAnalysisWhoseNodeIsExpanded, CopyBranchCompareAnalysis(mapOrigCopy, _compareanalysis))
	}
	for _, _note := range libraryFrom.NotesWhoseNodeIsExpanded {
		libraryTo.NotesWhoseNodeIsExpanded = append(libraryTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}

	return
}

func CopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.CopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range noteFrom.Complexities {
		noteTo.Complexities = append(noteTo.Complexities, CopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range noteFrom.Performances {
		noteTo.Performances = append(noteTo.Performances, CopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range noteFrom.Efforts {
		noteTo.Efforts = append(noteTo.Efforts, CopyBranchEffort(mapOrigCopy, _effort))
	}

	return
}

func CopyBranchNoteComplexityShape(mapOrigCopy map[any]any, notecomplexityshapeFrom *NoteComplexityShape) (notecomplexityshapeTo *NoteComplexityShape) {

	// notecomplexityshapeFrom has already been copied
	if _notecomplexityshapeTo, ok := mapOrigCopy[notecomplexityshapeFrom]; ok {
		notecomplexityshapeTo = _notecomplexityshapeTo.(*NoteComplexityShape)
		return
	}

	notecomplexityshapeTo = new(NoteComplexityShape)
	mapOrigCopy[notecomplexityshapeFrom] = notecomplexityshapeTo
	notecomplexityshapeFrom.CopyBasicFields(notecomplexityshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notecomplexityshapeFrom.Note != nil {
		notecomplexityshapeTo.Note = CopyBranchNote(mapOrigCopy, notecomplexityshapeFrom.Note)
	}
	if notecomplexityshapeFrom.Complexity != nil {
		notecomplexityshapeTo.Complexity = CopyBranchComplexity(mapOrigCopy, notecomplexityshapeFrom.Complexity)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteEffortShape(mapOrigCopy map[any]any, noteeffortshapeFrom *NoteEffortShape) (noteeffortshapeTo *NoteEffortShape) {

	// noteeffortshapeFrom has already been copied
	if _noteeffortshapeTo, ok := mapOrigCopy[noteeffortshapeFrom]; ok {
		noteeffortshapeTo = _noteeffortshapeTo.(*NoteEffortShape)
		return
	}

	noteeffortshapeTo = new(NoteEffortShape)
	mapOrigCopy[noteeffortshapeFrom] = noteeffortshapeTo
	noteeffortshapeFrom.CopyBasicFields(noteeffortshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteeffortshapeFrom.Note != nil {
		noteeffortshapeTo.Note = CopyBranchNote(mapOrigCopy, noteeffortshapeFrom.Note)
	}
	if noteeffortshapeFrom.Effort != nil {
		noteeffortshapeTo.Effort = CopyBranchEffort(mapOrigCopy, noteeffortshapeFrom.Effort)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNotePerformanceShape(mapOrigCopy map[any]any, noteperformanceshapeFrom *NotePerformanceShape) (noteperformanceshapeTo *NotePerformanceShape) {

	// noteperformanceshapeFrom has already been copied
	if _noteperformanceshapeTo, ok := mapOrigCopy[noteperformanceshapeFrom]; ok {
		noteperformanceshapeTo = _noteperformanceshapeTo.(*NotePerformanceShape)
		return
	}

	noteperformanceshapeTo = new(NotePerformanceShape)
	mapOrigCopy[noteperformanceshapeFrom] = noteperformanceshapeTo
	noteperformanceshapeFrom.CopyBasicFields(noteperformanceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteperformanceshapeFrom.Note != nil {
		noteperformanceshapeTo.Note = CopyBranchNote(mapOrigCopy, noteperformanceshapeFrom.Note)
	}
	if noteperformanceshapeFrom.Performance != nil {
		noteperformanceshapeTo.Performance = CopyBranchPerformance(mapOrigCopy, noteperformanceshapeFrom.Performance)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.CopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = CopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPerformance(mapOrigCopy map[any]any, performanceFrom *Performance) (performanceTo *Performance) {

	// performanceFrom has already been copied
	if _performanceTo, ok := mapOrigCopy[performanceFrom]; ok {
		performanceTo = _performanceTo.(*Performance)
		return
	}

	performanceTo = new(Performance)
	mapOrigCopy[performanceFrom] = performanceTo
	performanceFrom.CopyBasicFields(performanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSystem(mapOrigCopy map[any]any, systemFrom *System) (systemTo *System) {

	// systemFrom has already been copied
	if _systemTo, ok := mapOrigCopy[systemFrom]; ok {
		systemTo = _systemTo.(*System)
		return
	}

	systemTo = new(System)
	mapOrigCopy[systemFrom] = systemTo
	systemFrom.CopyBasicFields(systemTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range systemFrom.Complexities {
		systemTo.Complexities = append(systemTo.Complexities, CopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range systemFrom.Performances {
		systemTo.Performances = append(systemTo.Performances, CopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range systemFrom.Efforts {
		systemTo.Efforts = append(systemTo.Efforts, CopyBranchEffort(mapOrigCopy, _effort))
	}
	for _, _system := range systemFrom.SubSystems {
		systemTo.SubSystems = append(systemTo.SubSystems, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _diagramflossequation := range systemFrom.DiagramFlossEquations {
		systemTo.DiagramFlossEquations = append(systemTo.DiagramFlossEquations, CopyBranchDiagramFlossEquation(mapOrigCopy, _diagramflossequation))
	}
	for _, _diagramflossequation := range systemFrom.DiagramFlossEquationsWhoseNodeIsExpanded {
		systemTo.DiagramFlossEquationsWhoseNodeIsExpanded = append(systemTo.DiagramFlossEquationsWhoseNodeIsExpanded, CopyBranchDiagramFlossEquation(mapOrigCopy, _diagramflossequation))
	}
	for _, _complexity := range systemFrom.ComplexitysWhoseNodeIsExpanded {
		systemTo.ComplexitysWhoseNodeIsExpanded = append(systemTo.ComplexitysWhoseNodeIsExpanded, CopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range systemFrom.PerformancesWhoseNodeIsExpanded {
		systemTo.PerformancesWhoseNodeIsExpanded = append(systemTo.PerformancesWhoseNodeIsExpanded, CopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range systemFrom.EffortsWhoseNodeIsExpanded {
		systemTo.EffortsWhoseNodeIsExpanded = append(systemTo.EffortsWhoseNodeIsExpanded, CopyBranchEffort(mapOrigCopy, _effort))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *CompareAnalysis:
		stage.UnstageBranchCompareAnalysis(target)

	case *Complexity:
		stage.UnstageBranchComplexity(target)

	case *DiagramFlossEquation:
		stage.UnstageBranchDiagramFlossEquation(target)

	case *Effort:
		stage.UnstageBranchEffort(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Note:
		stage.UnstageBranchNote(target)

	case *NoteComplexityShape:
		stage.UnstageBranchNoteComplexityShape(target)

	case *NoteEffortShape:
		stage.UnstageBranchNoteEffortShape(target)

	case *NotePerformanceShape:
		stage.UnstageBranchNotePerformanceShape(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *Performance:
		stage.UnstageBranchPerformance(target)

	case *System:
		stage.UnstageBranchSystem(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchCompareAnalysis(compareanalysis *CompareAnalysis) {

	// check if instance is already staged
	if !IsStaged(stage, compareanalysis) {
		return
	}

	compareanalysis.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if compareanalysis.FromSystem != nil {
		UnstageBranch(stage, compareanalysis.FromSystem)
	}
	if compareanalysis.ToSystem != nil {
		UnstageBranch(stage, compareanalysis.ToSystem)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramflossequation := range compareanalysis.DiagramFlossEquations {
		UnstageBranch(stage, _diagramflossequation)
	}
	for _, _diagramflossequation := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
		UnstageBranch(stage, _diagramflossequation)
	}

}

func (stage *Stage) UnstageBranchComplexity(complexity *Complexity) {

	// check if instance is already staged
	if !IsStaged(stage, complexity) {
		return
	}

	complexity.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDiagramFlossEquation(diagramflossequation *DiagramFlossEquation) {

	// check if instance is already staged
	if !IsStaged(stage, diagramflossequation) {
		return
	}

	diagramflossequation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshape := range diagramflossequation.Note_Shapes {
		UnstageBranch(stage, _noteshape)
	}
	for _, _notecomplexityshape := range diagramflossequation.NoteComplexityShapes {
		UnstageBranch(stage, _notecomplexityshape)
	}
	for _, _noteperformanceshape := range diagramflossequation.NotePerformanceShapes {
		UnstageBranch(stage, _noteperformanceshape)
	}
	for _, _noteeffortshape := range diagramflossequation.NoteEffortShapes {
		UnstageBranch(stage, _noteeffortshape)
	}
	for _, _note := range diagramflossequation.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}
	for _, _complexity := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
		UnstageBranch(stage, _complexity)
	}
	for _, _performance := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
		UnstageBranch(stage, _performance)
	}
	for _, _effort := range diagramflossequation.EffortsWhoseNodeIsExpanded {
		UnstageBranch(stage, _effort)
	}

}

func (stage *Stage) UnstageBranchEffort(effort *Effort) {

	// check if instance is already staged
	if !IsStaged(stage, effort) {
		return
	}

	effort.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !IsStaged(stage, library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		UnstageBranch(stage, _library)
	}
	for _, _system := range library.RootSystems {
		UnstageBranch(stage, _system)
	}
	for _, _complexity := range library.RootComplexitys {
		UnstageBranch(stage, _complexity)
	}
	for _, _performance := range library.RootPerformances {
		UnstageBranch(stage, _performance)
	}
	for _, _effort := range library.RootEfforts {
		UnstageBranch(stage, _effort)
	}
	for _, _compareanalysis := range library.RootCompareAnalysis {
		UnstageBranch(stage, _compareanalysis)
	}
	for _, _note := range library.RootNotes {
		UnstageBranch(stage, _note)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		UnstageBranch(stage, _library)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		UnstageBranch(stage, _system)
	}
	for _, _complexity := range library.ComplexitysWhoseNodeIsExpanded {
		UnstageBranch(stage, _complexity)
	}
	for _, _performance := range library.PerformancesWhoseNodeIsExpanded {
		UnstageBranch(stage, _performance)
	}
	for _, _effort := range library.EffortsWhoseNodeIsExpanded {
		UnstageBranch(stage, _effort)
	}
	for _, _compareanalysis := range library.CompareAnalysisWhoseNodeIsExpanded {
		UnstageBranch(stage, _compareanalysis)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}

}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !IsStaged(stage, note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range note.Complexities {
		UnstageBranch(stage, _complexity)
	}
	for _, _performance := range note.Performances {
		UnstageBranch(stage, _performance)
	}
	for _, _effort := range note.Efforts {
		UnstageBranch(stage, _effort)
	}

}

func (stage *Stage) UnstageBranchNoteComplexityShape(notecomplexityshape *NoteComplexityShape) {

	// check if instance is already staged
	if !IsStaged(stage, notecomplexityshape) {
		return
	}

	notecomplexityshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notecomplexityshape.Note != nil {
		UnstageBranch(stage, notecomplexityshape.Note)
	}
	if notecomplexityshape.Complexity != nil {
		UnstageBranch(stage, notecomplexityshape.Complexity)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteEffortShape(noteeffortshape *NoteEffortShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteeffortshape) {
		return
	}

	noteeffortshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteeffortshape.Note != nil {
		UnstageBranch(stage, noteeffortshape.Note)
	}
	if noteeffortshape.Effort != nil {
		UnstageBranch(stage, noteeffortshape.Effort)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNotePerformanceShape(noteperformanceshape *NotePerformanceShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteperformanceshape) {
		return
	}

	noteperformanceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteperformanceshape.Note != nil {
		UnstageBranch(stage, noteperformanceshape.Note)
	}
	if noteperformanceshape.Performance != nil {
		UnstageBranch(stage, noteperformanceshape.Performance)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		UnstageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPerformance(performance *Performance) {

	// check if instance is already staged
	if !IsStaged(stage, performance) {
		return
	}

	performance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSystem(system *System) {

	// check if instance is already staged
	if !IsStaged(stage, system) {
		return
	}

	system.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range system.Complexities {
		UnstageBranch(stage, _complexity)
	}
	for _, _performance := range system.Performances {
		UnstageBranch(stage, _performance)
	}
	for _, _effort := range system.Efforts {
		UnstageBranch(stage, _effort)
	}
	for _, _system := range system.SubSystems {
		UnstageBranch(stage, _system)
	}
	for _, _diagramflossequation := range system.DiagramFlossEquations {
		UnstageBranch(stage, _diagramflossequation)
	}
	for _, _diagramflossequation := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
		UnstageBranch(stage, _diagramflossequation)
	}
	for _, _complexity := range system.ComplexitysWhoseNodeIsExpanded {
		UnstageBranch(stage, _complexity)
	}
	for _, _performance := range system.PerformancesWhoseNodeIsExpanded {
		UnstageBranch(stage, _performance)
	}
	for _, _effort := range system.EffortsWhoseNodeIsExpanded {
		UnstageBranch(stage, _effort)
	}

}

// insertion point for pointer reconstruction from references
func (reference *CompareAnalysis) GongReconstructPointersFromReferences(stage *Stage, instance *CompareAnalysis) {
	// insertion point for pointers field
	if instance.FromSystem != nil {
		reference.FromSystem = stage.Systems_reference[instance.FromSystem]
	}
	if instance.ToSystem != nil {
		reference.ToSystem = stage.Systems_reference[instance.ToSystem]
	}
	// insertion point for slice of pointers field
	reference.DiagramFlossEquations = reference.DiagramFlossEquations[:0]
	for _, _b := range instance.DiagramFlossEquations {
		reference.DiagramFlossEquations = append(reference.DiagramFlossEquations, stage.DiagramFlossEquations_reference[_b])
	}
	reference.DiagramFlossEquationsWhoseNodeIsExpanded = reference.DiagramFlossEquationsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DiagramFlossEquationsWhoseNodeIsExpanded {
		reference.DiagramFlossEquationsWhoseNodeIsExpanded = append(reference.DiagramFlossEquationsWhoseNodeIsExpanded, stage.DiagramFlossEquations_reference[_b])
	}
}

func (reference *Complexity) GongReconstructPointersFromReferences(stage *Stage, instance *Complexity) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DiagramFlossEquation) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramFlossEquation) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Note_Shapes = reference.Note_Shapes[:0]
	for _, _b := range instance.Note_Shapes {
		reference.Note_Shapes = append(reference.Note_Shapes, stage.NoteShapes_reference[_b])
	}
	reference.NoteComplexityShapes = reference.NoteComplexityShapes[:0]
	for _, _b := range instance.NoteComplexityShapes {
		reference.NoteComplexityShapes = append(reference.NoteComplexityShapes, stage.NoteComplexityShapes_reference[_b])
	}
	reference.NotePerformanceShapes = reference.NotePerformanceShapes[:0]
	for _, _b := range instance.NotePerformanceShapes {
		reference.NotePerformanceShapes = append(reference.NotePerformanceShapes, stage.NotePerformanceShapes_reference[_b])
	}
	reference.NoteEffortShapes = reference.NoteEffortShapes[:0]
	for _, _b := range instance.NoteEffortShapes {
		reference.NoteEffortShapes = append(reference.NoteEffortShapes, stage.NoteEffortShapes_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
	reference.ComplexitysWhoseNodeIsExpanded = reference.ComplexitysWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ComplexitysWhoseNodeIsExpanded {
		reference.ComplexitysWhoseNodeIsExpanded = append(reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_reference[_b])
	}
	reference.PerformancesWhoseNodeIsExpanded = reference.PerformancesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PerformancesWhoseNodeIsExpanded {
		reference.PerformancesWhoseNodeIsExpanded = append(reference.PerformancesWhoseNodeIsExpanded, stage.Performances_reference[_b])
	}
	reference.EffortsWhoseNodeIsExpanded = reference.EffortsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.EffortsWhoseNodeIsExpanded {
		reference.EffortsWhoseNodeIsExpanded = append(reference.EffortsWhoseNodeIsExpanded, stage.Efforts_reference[_b])
	}
}

func (reference *Effort) GongReconstructPointersFromReferences(stage *Stage, instance *Effort) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.RootSystems = reference.RootSystems[:0]
	for _, _b := range instance.RootSystems {
		reference.RootSystems = append(reference.RootSystems, stage.Systems_reference[_b])
	}
	reference.RootComplexitys = reference.RootComplexitys[:0]
	for _, _b := range instance.RootComplexitys {
		reference.RootComplexitys = append(reference.RootComplexitys, stage.Complexitys_reference[_b])
	}
	reference.RootPerformances = reference.RootPerformances[:0]
	for _, _b := range instance.RootPerformances {
		reference.RootPerformances = append(reference.RootPerformances, stage.Performances_reference[_b])
	}
	reference.RootEfforts = reference.RootEfforts[:0]
	for _, _b := range instance.RootEfforts {
		reference.RootEfforts = append(reference.RootEfforts, stage.Efforts_reference[_b])
	}
	reference.RootCompareAnalysis = reference.RootCompareAnalysis[:0]
	for _, _b := range instance.RootCompareAnalysis {
		reference.RootCompareAnalysis = append(reference.RootCompareAnalysis, stage.CompareAnalysiss_reference[_b])
	}
	reference.RootNotes = reference.RootNotes[:0]
	for _, _b := range instance.RootNotes {
		reference.RootNotes = append(reference.RootNotes, stage.Notes_reference[_b])
	}
	reference.SubLibrariesWhoseNodeIsExpanded = reference.SubLibrariesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SubLibrariesWhoseNodeIsExpanded {
		reference.SubLibrariesWhoseNodeIsExpanded = append(reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference[_b])
	}
	reference.SystemsWhoseNodeIsExpanded = reference.SystemsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SystemsWhoseNodeIsExpanded {
		reference.SystemsWhoseNodeIsExpanded = append(reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference[_b])
	}
	reference.ComplexitysWhoseNodeIsExpanded = reference.ComplexitysWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ComplexitysWhoseNodeIsExpanded {
		reference.ComplexitysWhoseNodeIsExpanded = append(reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_reference[_b])
	}
	reference.PerformancesWhoseNodeIsExpanded = reference.PerformancesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PerformancesWhoseNodeIsExpanded {
		reference.PerformancesWhoseNodeIsExpanded = append(reference.PerformancesWhoseNodeIsExpanded, stage.Performances_reference[_b])
	}
	reference.EffortsWhoseNodeIsExpanded = reference.EffortsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.EffortsWhoseNodeIsExpanded {
		reference.EffortsWhoseNodeIsExpanded = append(reference.EffortsWhoseNodeIsExpanded, stage.Efforts_reference[_b])
	}
	reference.CompareAnalysisWhoseNodeIsExpanded = reference.CompareAnalysisWhoseNodeIsExpanded[:0]
	for _, _b := range instance.CompareAnalysisWhoseNodeIsExpanded {
		reference.CompareAnalysisWhoseNodeIsExpanded = append(reference.CompareAnalysisWhoseNodeIsExpanded, stage.CompareAnalysiss_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Complexities = reference.Complexities[:0]
	for _, _b := range instance.Complexities {
		reference.Complexities = append(reference.Complexities, stage.Complexitys_reference[_b])
	}
	reference.Performances = reference.Performances[:0]
	for _, _b := range instance.Performances {
		reference.Performances = append(reference.Performances, stage.Performances_reference[_b])
	}
	reference.Efforts = reference.Efforts[:0]
	for _, _b := range instance.Efforts {
		reference.Efforts = append(reference.Efforts, stage.Efforts_reference[_b])
	}
}

func (reference *NoteComplexityShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteComplexityShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Complexity != nil {
		reference.Complexity = stage.Complexitys_reference[instance.Complexity]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteEffortShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteEffortShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Effort != nil {
		reference.Effort = stage.Efforts_reference[instance.Effort]
	}
	// insertion point for slice of pointers field
}

func (reference *NotePerformanceShape) GongReconstructPointersFromReferences(stage *Stage, instance *NotePerformanceShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Performance != nil {
		reference.Performance = stage.Performances_reference[instance.Performance]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	// insertion point for slice of pointers field
}

func (reference *Performance) GongReconstructPointersFromReferences(stage *Stage, instance *Performance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *System) GongReconstructPointersFromReferences(stage *Stage, instance *System) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Complexities = reference.Complexities[:0]
	for _, _b := range instance.Complexities {
		reference.Complexities = append(reference.Complexities, stage.Complexitys_reference[_b])
	}
	reference.Performances = reference.Performances[:0]
	for _, _b := range instance.Performances {
		reference.Performances = append(reference.Performances, stage.Performances_reference[_b])
	}
	reference.Efforts = reference.Efforts[:0]
	for _, _b := range instance.Efforts {
		reference.Efforts = append(reference.Efforts, stage.Efforts_reference[_b])
	}
	reference.SubSystems = reference.SubSystems[:0]
	for _, _b := range instance.SubSystems {
		reference.SubSystems = append(reference.SubSystems, stage.Systems_reference[_b])
	}
	reference.DiagramFlossEquations = reference.DiagramFlossEquations[:0]
	for _, _b := range instance.DiagramFlossEquations {
		reference.DiagramFlossEquations = append(reference.DiagramFlossEquations, stage.DiagramFlossEquations_reference[_b])
	}
	reference.DiagramFlossEquationsWhoseNodeIsExpanded = reference.DiagramFlossEquationsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DiagramFlossEquationsWhoseNodeIsExpanded {
		reference.DiagramFlossEquationsWhoseNodeIsExpanded = append(reference.DiagramFlossEquationsWhoseNodeIsExpanded, stage.DiagramFlossEquations_reference[_b])
	}
	reference.ComplexitysWhoseNodeIsExpanded = reference.ComplexitysWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ComplexitysWhoseNodeIsExpanded {
		reference.ComplexitysWhoseNodeIsExpanded = append(reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_reference[_b])
	}
	reference.PerformancesWhoseNodeIsExpanded = reference.PerformancesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PerformancesWhoseNodeIsExpanded {
		reference.PerformancesWhoseNodeIsExpanded = append(reference.PerformancesWhoseNodeIsExpanded, stage.Performances_reference[_b])
	}
	reference.EffortsWhoseNodeIsExpanded = reference.EffortsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.EffortsWhoseNodeIsExpanded {
		reference.EffortsWhoseNodeIsExpanded = append(reference.EffortsWhoseNodeIsExpanded, stage.Efforts_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *CompareAnalysis) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.FromSystem; _reference != nil {
		reference.FromSystem = nil
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			reference.FromSystem = _instance
		}
	}
	if _reference := reference.ToSystem; _reference != nil {
		reference.ToSystem = nil
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			reference.ToSystem = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _DiagramFlossEquations []*DiagramFlossEquation
	for _, _reference := range reference.DiagramFlossEquations {
		if _instance, ok := stage.DiagramFlossEquations_instance[_reference]; ok {
			_DiagramFlossEquations = append(_DiagramFlossEquations, _instance)
		}
	}
	reference.DiagramFlossEquations = _DiagramFlossEquations
	var _DiagramFlossEquationsWhoseNodeIsExpanded []*DiagramFlossEquation
	for _, _reference := range reference.DiagramFlossEquationsWhoseNodeIsExpanded {
		if _instance, ok := stage.DiagramFlossEquations_instance[_reference]; ok {
			_DiagramFlossEquationsWhoseNodeIsExpanded = append(_DiagramFlossEquationsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DiagramFlossEquationsWhoseNodeIsExpanded = _DiagramFlossEquationsWhoseNodeIsExpanded
}

func (reference *Complexity) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DiagramFlossEquation) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Note_Shapes []*NoteShape
	for _, _reference := range reference.Note_Shapes {
		if _instance, ok := stage.NoteShapes_instance[_reference]; ok {
			_Note_Shapes = append(_Note_Shapes, _instance)
		}
	}
	reference.Note_Shapes = _Note_Shapes
	var _NoteComplexityShapes []*NoteComplexityShape
	for _, _reference := range reference.NoteComplexityShapes {
		if _instance, ok := stage.NoteComplexityShapes_instance[_reference]; ok {
			_NoteComplexityShapes = append(_NoteComplexityShapes, _instance)
		}
	}
	reference.NoteComplexityShapes = _NoteComplexityShapes
	var _NotePerformanceShapes []*NotePerformanceShape
	for _, _reference := range reference.NotePerformanceShapes {
		if _instance, ok := stage.NotePerformanceShapes_instance[_reference]; ok {
			_NotePerformanceShapes = append(_NotePerformanceShapes, _instance)
		}
	}
	reference.NotePerformanceShapes = _NotePerformanceShapes
	var _NoteEffortShapes []*NoteEffortShape
	for _, _reference := range reference.NoteEffortShapes {
		if _instance, ok := stage.NoteEffortShapes_instance[_reference]; ok {
			_NoteEffortShapes = append(_NoteEffortShapes, _instance)
		}
	}
	reference.NoteEffortShapes = _NoteEffortShapes
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
	var _ComplexitysWhoseNodeIsExpanded []*Complexity
	for _, _reference := range reference.ComplexitysWhoseNodeIsExpanded {
		if _instance, ok := stage.Complexitys_instance[_reference]; ok {
			_ComplexitysWhoseNodeIsExpanded = append(_ComplexitysWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ComplexitysWhoseNodeIsExpanded = _ComplexitysWhoseNodeIsExpanded
	var _PerformancesWhoseNodeIsExpanded []*Performance
	for _, _reference := range reference.PerformancesWhoseNodeIsExpanded {
		if _instance, ok := stage.Performances_instance[_reference]; ok {
			_PerformancesWhoseNodeIsExpanded = append(_PerformancesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PerformancesWhoseNodeIsExpanded = _PerformancesWhoseNodeIsExpanded
	var _EffortsWhoseNodeIsExpanded []*Effort
	for _, _reference := range reference.EffortsWhoseNodeIsExpanded {
		if _instance, ok := stage.Efforts_instance[_reference]; ok {
			_EffortsWhoseNodeIsExpanded = append(_EffortsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.EffortsWhoseNodeIsExpanded = _EffortsWhoseNodeIsExpanded
}

func (reference *Effort) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
	var _RootSystems []*System
	for _, _reference := range reference.RootSystems {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_RootSystems = append(_RootSystems, _instance)
		}
	}
	reference.RootSystems = _RootSystems
	var _RootComplexitys []*Complexity
	for _, _reference := range reference.RootComplexitys {
		if _instance, ok := stage.Complexitys_instance[_reference]; ok {
			_RootComplexitys = append(_RootComplexitys, _instance)
		}
	}
	reference.RootComplexitys = _RootComplexitys
	var _RootPerformances []*Performance
	for _, _reference := range reference.RootPerformances {
		if _instance, ok := stage.Performances_instance[_reference]; ok {
			_RootPerformances = append(_RootPerformances, _instance)
		}
	}
	reference.RootPerformances = _RootPerformances
	var _RootEfforts []*Effort
	for _, _reference := range reference.RootEfforts {
		if _instance, ok := stage.Efforts_instance[_reference]; ok {
			_RootEfforts = append(_RootEfforts, _instance)
		}
	}
	reference.RootEfforts = _RootEfforts
	var _RootCompareAnalysis []*CompareAnalysis
	for _, _reference := range reference.RootCompareAnalysis {
		if _instance, ok := stage.CompareAnalysiss_instance[_reference]; ok {
			_RootCompareAnalysis = append(_RootCompareAnalysis, _instance)
		}
	}
	reference.RootCompareAnalysis = _RootCompareAnalysis
	var _RootNotes []*Note
	for _, _reference := range reference.RootNotes {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_RootNotes = append(_RootNotes, _instance)
		}
	}
	reference.RootNotes = _RootNotes
	var _SubLibrariesWhoseNodeIsExpanded []*Library
	for _, _reference := range reference.SubLibrariesWhoseNodeIsExpanded {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibrariesWhoseNodeIsExpanded = append(_SubLibrariesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SubLibrariesWhoseNodeIsExpanded = _SubLibrariesWhoseNodeIsExpanded
	var _SystemsWhoseNodeIsExpanded []*System
	for _, _reference := range reference.SystemsWhoseNodeIsExpanded {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SystemsWhoseNodeIsExpanded = append(_SystemsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SystemsWhoseNodeIsExpanded = _SystemsWhoseNodeIsExpanded
	var _ComplexitysWhoseNodeIsExpanded []*Complexity
	for _, _reference := range reference.ComplexitysWhoseNodeIsExpanded {
		if _instance, ok := stage.Complexitys_instance[_reference]; ok {
			_ComplexitysWhoseNodeIsExpanded = append(_ComplexitysWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ComplexitysWhoseNodeIsExpanded = _ComplexitysWhoseNodeIsExpanded
	var _PerformancesWhoseNodeIsExpanded []*Performance
	for _, _reference := range reference.PerformancesWhoseNodeIsExpanded {
		if _instance, ok := stage.Performances_instance[_reference]; ok {
			_PerformancesWhoseNodeIsExpanded = append(_PerformancesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PerformancesWhoseNodeIsExpanded = _PerformancesWhoseNodeIsExpanded
	var _EffortsWhoseNodeIsExpanded []*Effort
	for _, _reference := range reference.EffortsWhoseNodeIsExpanded {
		if _instance, ok := stage.Efforts_instance[_reference]; ok {
			_EffortsWhoseNodeIsExpanded = append(_EffortsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.EffortsWhoseNodeIsExpanded = _EffortsWhoseNodeIsExpanded
	var _CompareAnalysisWhoseNodeIsExpanded []*CompareAnalysis
	for _, _reference := range reference.CompareAnalysisWhoseNodeIsExpanded {
		if _instance, ok := stage.CompareAnalysiss_instance[_reference]; ok {
			_CompareAnalysisWhoseNodeIsExpanded = append(_CompareAnalysisWhoseNodeIsExpanded, _instance)
		}
	}
	reference.CompareAnalysisWhoseNodeIsExpanded = _CompareAnalysisWhoseNodeIsExpanded
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Complexities []*Complexity
	for _, _reference := range reference.Complexities {
		if _instance, ok := stage.Complexitys_instance[_reference]; ok {
			_Complexities = append(_Complexities, _instance)
		}
	}
	reference.Complexities = _Complexities
	var _Performances []*Performance
	for _, _reference := range reference.Performances {
		if _instance, ok := stage.Performances_instance[_reference]; ok {
			_Performances = append(_Performances, _instance)
		}
	}
	reference.Performances = _Performances
	var _Efforts []*Effort
	for _, _reference := range reference.Efforts {
		if _instance, ok := stage.Efforts_instance[_reference]; ok {
			_Efforts = append(_Efforts, _instance)
		}
	}
	reference.Efforts = _Efforts
}

func (reference *NoteComplexityShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Complexity; _reference != nil {
		reference.Complexity = nil
		if _instance, ok := stage.Complexitys_instance[_reference]; ok {
			reference.Complexity = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteEffortShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Effort; _reference != nil {
		reference.Effort = nil
		if _instance, ok := stage.Efforts_instance[_reference]; ok {
			reference.Effort = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NotePerformanceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Performance; _reference != nil {
		reference.Performance = nil
		if _instance, ok := stage.Performances_instance[_reference]; ok {
			reference.Performance = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Performance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *System) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Complexities []*Complexity
	for _, _reference := range reference.Complexities {
		if _instance, ok := stage.Complexitys_instance[_reference]; ok {
			_Complexities = append(_Complexities, _instance)
		}
	}
	reference.Complexities = _Complexities
	var _Performances []*Performance
	for _, _reference := range reference.Performances {
		if _instance, ok := stage.Performances_instance[_reference]; ok {
			_Performances = append(_Performances, _instance)
		}
	}
	reference.Performances = _Performances
	var _Efforts []*Effort
	for _, _reference := range reference.Efforts {
		if _instance, ok := stage.Efforts_instance[_reference]; ok {
			_Efforts = append(_Efforts, _instance)
		}
	}
	reference.Efforts = _Efforts
	var _SubSystems []*System
	for _, _reference := range reference.SubSystems {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SubSystems = append(_SubSystems, _instance)
		}
	}
	reference.SubSystems = _SubSystems
	var _DiagramFlossEquations []*DiagramFlossEquation
	for _, _reference := range reference.DiagramFlossEquations {
		if _instance, ok := stage.DiagramFlossEquations_instance[_reference]; ok {
			_DiagramFlossEquations = append(_DiagramFlossEquations, _instance)
		}
	}
	reference.DiagramFlossEquations = _DiagramFlossEquations
	var _DiagramFlossEquationsWhoseNodeIsExpanded []*DiagramFlossEquation
	for _, _reference := range reference.DiagramFlossEquationsWhoseNodeIsExpanded {
		if _instance, ok := stage.DiagramFlossEquations_instance[_reference]; ok {
			_DiagramFlossEquationsWhoseNodeIsExpanded = append(_DiagramFlossEquationsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DiagramFlossEquationsWhoseNodeIsExpanded = _DiagramFlossEquationsWhoseNodeIsExpanded
	var _ComplexitysWhoseNodeIsExpanded []*Complexity
	for _, _reference := range reference.ComplexitysWhoseNodeIsExpanded {
		if _instance, ok := stage.Complexitys_instance[_reference]; ok {
			_ComplexitysWhoseNodeIsExpanded = append(_ComplexitysWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ComplexitysWhoseNodeIsExpanded = _ComplexitysWhoseNodeIsExpanded
	var _PerformancesWhoseNodeIsExpanded []*Performance
	for _, _reference := range reference.PerformancesWhoseNodeIsExpanded {
		if _instance, ok := stage.Performances_instance[_reference]; ok {
			_PerformancesWhoseNodeIsExpanded = append(_PerformancesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PerformancesWhoseNodeIsExpanded = _PerformancesWhoseNodeIsExpanded
	var _EffortsWhoseNodeIsExpanded []*Effort
	for _, _reference := range reference.EffortsWhoseNodeIsExpanded {
		if _instance, ok := stage.Efforts_instance[_reference]; ok {
			_EffortsWhoseNodeIsExpanded = append(_EffortsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.EffortsWhoseNodeIsExpanded = _EffortsWhoseNodeIsExpanded
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (compareanalysis *CompareAnalysis) GongDiff(stage *Stage, compareanalysisOther *CompareAnalysis) (diffs []string) {
	// insertion point for field diffs
	if compareanalysis.Name != compareanalysisOther.Name {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "Name"))
	}
	if (compareanalysis.FromSystem == nil) != (compareanalysisOther.FromSystem == nil) {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "FromSystem"))
	} else if compareanalysis.FromSystem != nil && compareanalysisOther.FromSystem != nil {
		if compareanalysis.FromSystem != compareanalysisOther.FromSystem {
			diffs = append(diffs, compareanalysis.GongMarshallField(stage, "FromSystem"))
		}
	}
	if (compareanalysis.ToSystem == nil) != (compareanalysisOther.ToSystem == nil) {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "ToSystem"))
	} else if compareanalysis.ToSystem != nil && compareanalysisOther.ToSystem != nil {
		if compareanalysis.ToSystem != compareanalysisOther.ToSystem {
			diffs = append(diffs, compareanalysis.GongMarshallField(stage, "ToSystem"))
		}
	}
	if compareanalysis.Mu != compareanalysisOther.Mu {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "Mu"))
	}
	if compareanalysis.Epsilon != compareanalysisOther.Epsilon {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "Epsilon"))
	}
	DiagramFlossEquationsDifferent := false
	if len(compareanalysis.DiagramFlossEquations) != len(compareanalysisOther.DiagramFlossEquations) {
		DiagramFlossEquationsDifferent = true
	} else {
		for i := range compareanalysis.DiagramFlossEquations {
			if (compareanalysis.DiagramFlossEquations[i] == nil) != (compareanalysisOther.DiagramFlossEquations[i] == nil) {
				DiagramFlossEquationsDifferent = true
				break
			} else if compareanalysis.DiagramFlossEquations[i] != nil && compareanalysisOther.DiagramFlossEquations[i] != nil {
				// this is a pointer comparaison
				if compareanalysis.DiagramFlossEquations[i] != compareanalysisOther.DiagramFlossEquations[i] {
					DiagramFlossEquationsDifferent = true
					break
				}
			}
		}
	}
	if DiagramFlossEquationsDifferent {
		ops := Diff(stage, compareanalysis, compareanalysisOther, "DiagramFlossEquations", compareanalysisOther.DiagramFlossEquations, compareanalysis.DiagramFlossEquations)
		diffs = append(diffs, ops)
	}
	DiagramFlossEquationsWhoseNodeIsExpandedDifferent := false
	if len(compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded) != len(compareanalysisOther.DiagramFlossEquationsWhoseNodeIsExpanded) {
		DiagramFlossEquationsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
			if (compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded[i] == nil) != (compareanalysisOther.DiagramFlossEquationsWhoseNodeIsExpanded[i] == nil) {
				DiagramFlossEquationsWhoseNodeIsExpandedDifferent = true
				break
			} else if compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded[i] != nil && compareanalysisOther.DiagramFlossEquationsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded[i] != compareanalysisOther.DiagramFlossEquationsWhoseNodeIsExpanded[i] {
					DiagramFlossEquationsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DiagramFlossEquationsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, compareanalysis, compareanalysisOther, "DiagramFlossEquationsWhoseNodeIsExpanded", compareanalysisOther.DiagramFlossEquationsWhoseNodeIsExpanded, compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if compareanalysis.ComputedPrefix != compareanalysisOther.ComputedPrefix {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "ComputedPrefix"))
	}
	if compareanalysis.IsExpanded != compareanalysisOther.IsExpanded {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (complexity *Complexity) GongDiff(stage *Stage, complexityOther *Complexity) (diffs []string) {
	// insertion point for field diffs
	if complexity.Name != complexityOther.Name {
		diffs = append(diffs, complexity.GongMarshallField(stage, "Name"))
	}
	if complexity.Strength != complexityOther.Strength {
		diffs = append(diffs, complexity.GongMarshallField(stage, "Strength"))
	}
	if complexity.Description != complexityOther.Description {
		diffs = append(diffs, complexity.GongMarshallField(stage, "Description"))
	}
	if complexity.ComputedPrefix != complexityOther.ComputedPrefix {
		diffs = append(diffs, complexity.GongMarshallField(stage, "ComputedPrefix"))
	}
	if complexity.IsExpanded != complexityOther.IsExpanded {
		diffs = append(diffs, complexity.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagramflossequation *DiagramFlossEquation) GongDiff(stage *Stage, diagramflossequationOther *DiagramFlossEquation) (diffs []string) {
	// insertion point for field diffs
	if diagramflossequation.Name != diagramflossequationOther.Name {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "Name"))
	}
	if diagramflossequation.Description != diagramflossequationOther.Description {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "Description"))
	}
	if diagramflossequation.Scale != diagramflossequationOther.Scale {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "Scale"))
	}
	if diagramflossequation.FontSize != diagramflossequationOther.FontSize {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "FontSize"))
	}
	if diagramflossequation.ComputedPrefix != diagramflossequationOther.ComputedPrefix {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagramflossequation.IsExpanded != diagramflossequationOther.IsExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsExpanded"))
	}
	if diagramflossequation.IsChecked != diagramflossequationOther.IsChecked {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsChecked"))
	}
	if diagramflossequation.IsEditable_ != diagramflossequationOther.IsEditable_ {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsEditable_"))
	}
	if diagramflossequation.IsInDelta3ColumnsMode != diagramflossequationOther.IsInDelta3ColumnsMode {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsInDelta3ColumnsMode"))
	}
	if diagramflossequation.AreQuantitativeElementsVisible != diagramflossequationOther.AreQuantitativeElementsVisible {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "AreQuantitativeElementsVisible"))
	}
	if diagramflossequation.AreSubsystemsVisible != diagramflossequationOther.AreSubsystemsVisible {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "AreSubsystemsVisible"))
	}
	if diagramflossequation.Width != diagramflossequationOther.Width {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "Width"))
	}
	if diagramflossequation.Height != diagramflossequationOther.Height {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "Height"))
	}
	if diagramflossequation.DefaultBoxWidth != diagramflossequationOther.DefaultBoxWidth {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "DefaultBoxWidth"))
	}
	if diagramflossequation.DefaultBoxHeigth != diagramflossequationOther.DefaultBoxHeigth {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "DefaultBoxHeigth"))
	}
	Note_ShapesDifferent := false
	if len(diagramflossequation.Note_Shapes) != len(diagramflossequationOther.Note_Shapes) {
		Note_ShapesDifferent = true
	} else {
		for i := range diagramflossequation.Note_Shapes {
			if (diagramflossequation.Note_Shapes[i] == nil) != (diagramflossequationOther.Note_Shapes[i] == nil) {
				Note_ShapesDifferent = true
				break
			} else if diagramflossequation.Note_Shapes[i] != nil && diagramflossequationOther.Note_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramflossequation.Note_Shapes[i] != diagramflossequationOther.Note_Shapes[i] {
					Note_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Note_ShapesDifferent {
		ops := Diff(stage, diagramflossequation, diagramflossequationOther, "Note_Shapes", diagramflossequationOther.Note_Shapes, diagramflossequation.Note_Shapes)
		diffs = append(diffs, ops)
	}
	NoteComplexityShapesDifferent := false
	if len(diagramflossequation.NoteComplexityShapes) != len(diagramflossequationOther.NoteComplexityShapes) {
		NoteComplexityShapesDifferent = true
	} else {
		for i := range diagramflossequation.NoteComplexityShapes {
			if (diagramflossequation.NoteComplexityShapes[i] == nil) != (diagramflossequationOther.NoteComplexityShapes[i] == nil) {
				NoteComplexityShapesDifferent = true
				break
			} else if diagramflossequation.NoteComplexityShapes[i] != nil && diagramflossequationOther.NoteComplexityShapes[i] != nil {
				// this is a pointer comparaison
				if diagramflossequation.NoteComplexityShapes[i] != diagramflossequationOther.NoteComplexityShapes[i] {
					NoteComplexityShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteComplexityShapesDifferent {
		ops := Diff(stage, diagramflossequation, diagramflossequationOther, "NoteComplexityShapes", diagramflossequationOther.NoteComplexityShapes, diagramflossequation.NoteComplexityShapes)
		diffs = append(diffs, ops)
	}
	NotePerformanceShapesDifferent := false
	if len(diagramflossequation.NotePerformanceShapes) != len(diagramflossequationOther.NotePerformanceShapes) {
		NotePerformanceShapesDifferent = true
	} else {
		for i := range diagramflossequation.NotePerformanceShapes {
			if (diagramflossequation.NotePerformanceShapes[i] == nil) != (diagramflossequationOther.NotePerformanceShapes[i] == nil) {
				NotePerformanceShapesDifferent = true
				break
			} else if diagramflossequation.NotePerformanceShapes[i] != nil && diagramflossequationOther.NotePerformanceShapes[i] != nil {
				// this is a pointer comparaison
				if diagramflossequation.NotePerformanceShapes[i] != diagramflossequationOther.NotePerformanceShapes[i] {
					NotePerformanceShapesDifferent = true
					break
				}
			}
		}
	}
	if NotePerformanceShapesDifferent {
		ops := Diff(stage, diagramflossequation, diagramflossequationOther, "NotePerformanceShapes", diagramflossequationOther.NotePerformanceShapes, diagramflossequation.NotePerformanceShapes)
		diffs = append(diffs, ops)
	}
	NoteEffortShapesDifferent := false
	if len(diagramflossequation.NoteEffortShapes) != len(diagramflossequationOther.NoteEffortShapes) {
		NoteEffortShapesDifferent = true
	} else {
		for i := range diagramflossequation.NoteEffortShapes {
			if (diagramflossequation.NoteEffortShapes[i] == nil) != (diagramflossequationOther.NoteEffortShapes[i] == nil) {
				NoteEffortShapesDifferent = true
				break
			} else if diagramflossequation.NoteEffortShapes[i] != nil && diagramflossequationOther.NoteEffortShapes[i] != nil {
				// this is a pointer comparaison
				if diagramflossequation.NoteEffortShapes[i] != diagramflossequationOther.NoteEffortShapes[i] {
					NoteEffortShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteEffortShapesDifferent {
		ops := Diff(stage, diagramflossequation, diagramflossequationOther, "NoteEffortShapes", diagramflossequationOther.NoteEffortShapes, diagramflossequation.NoteEffortShapes)
		diffs = append(diffs, ops)
	}
	if diagramflossequation.IsNotesNodeExpanded != diagramflossequationOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(diagramflossequation.NotesWhoseNodeIsExpanded) != len(diagramflossequationOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramflossequation.NotesWhoseNodeIsExpanded {
			if (diagramflossequation.NotesWhoseNodeIsExpanded[i] == nil) != (diagramflossequationOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramflossequation.NotesWhoseNodeIsExpanded[i] != nil && diagramflossequationOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramflossequation.NotesWhoseNodeIsExpanded[i] != diagramflossequationOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramflossequation, diagramflossequationOther, "NotesWhoseNodeIsExpanded", diagramflossequationOther.NotesWhoseNodeIsExpanded, diagramflossequation.NotesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramflossequation.IsComplexitysNodeExpanded != diagramflossequationOther.IsComplexitysNodeExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
	}
	ComplexitysWhoseNodeIsExpandedDifferent := false
	if len(diagramflossequation.ComplexitysWhoseNodeIsExpanded) != len(diagramflossequationOther.ComplexitysWhoseNodeIsExpanded) {
		ComplexitysWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
			if (diagramflossequation.ComplexitysWhoseNodeIsExpanded[i] == nil) != (diagramflossequationOther.ComplexitysWhoseNodeIsExpanded[i] == nil) {
				ComplexitysWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramflossequation.ComplexitysWhoseNodeIsExpanded[i] != nil && diagramflossequationOther.ComplexitysWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramflossequation.ComplexitysWhoseNodeIsExpanded[i] != diagramflossequationOther.ComplexitysWhoseNodeIsExpanded[i] {
					ComplexitysWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ComplexitysWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramflossequation, diagramflossequationOther, "ComplexitysWhoseNodeIsExpanded", diagramflossequationOther.ComplexitysWhoseNodeIsExpanded, diagramflossequation.ComplexitysWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramflossequation.IsPerformancesNodeExpanded != diagramflossequationOther.IsPerformancesNodeExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
	}
	PerformancesWhoseNodeIsExpandedDifferent := false
	if len(diagramflossequation.PerformancesWhoseNodeIsExpanded) != len(diagramflossequationOther.PerformancesWhoseNodeIsExpanded) {
		PerformancesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
			if (diagramflossequation.PerformancesWhoseNodeIsExpanded[i] == nil) != (diagramflossequationOther.PerformancesWhoseNodeIsExpanded[i] == nil) {
				PerformancesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramflossequation.PerformancesWhoseNodeIsExpanded[i] != nil && diagramflossequationOther.PerformancesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramflossequation.PerformancesWhoseNodeIsExpanded[i] != diagramflossequationOther.PerformancesWhoseNodeIsExpanded[i] {
					PerformancesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PerformancesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramflossequation, diagramflossequationOther, "PerformancesWhoseNodeIsExpanded", diagramflossequationOther.PerformancesWhoseNodeIsExpanded, diagramflossequation.PerformancesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramflossequation.IsEffortsNodeExpanded != diagramflossequationOther.IsEffortsNodeExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsEffortsNodeExpanded"))
	}
	EffortsWhoseNodeIsExpandedDifferent := false
	if len(diagramflossequation.EffortsWhoseNodeIsExpanded) != len(diagramflossequationOther.EffortsWhoseNodeIsExpanded) {
		EffortsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramflossequation.EffortsWhoseNodeIsExpanded {
			if (diagramflossequation.EffortsWhoseNodeIsExpanded[i] == nil) != (diagramflossequationOther.EffortsWhoseNodeIsExpanded[i] == nil) {
				EffortsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramflossequation.EffortsWhoseNodeIsExpanded[i] != nil && diagramflossequationOther.EffortsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramflossequation.EffortsWhoseNodeIsExpanded[i] != diagramflossequationOther.EffortsWhoseNodeIsExpanded[i] {
					EffortsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if EffortsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramflossequation, diagramflossequationOther, "EffortsWhoseNodeIsExpanded", diagramflossequationOther.EffortsWhoseNodeIsExpanded, diagramflossequation.EffortsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (effort *Effort) GongDiff(stage *Stage, effortOther *Effort) (diffs []string) {
	// insertion point for field diffs
	if effort.Name != effortOther.Name {
		diffs = append(diffs, effort.GongMarshallField(stage, "Name"))
	}
	if effort.Strength != effortOther.Strength {
		diffs = append(diffs, effort.GongMarshallField(stage, "Strength"))
	}
	if effort.Description != effortOther.Description {
		diffs = append(diffs, effort.GongMarshallField(stage, "Description"))
	}
	if effort.ComputedPrefix != effortOther.ComputedPrefix {
		diffs = append(diffs, effort.GongMarshallField(stage, "ComputedPrefix"))
	}
	if effort.IsExpanded != effortOther.IsExpanded {
		diffs = append(diffs, effort.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (library *Library) GongDiff(stage *Stage, libraryOther *Library) (diffs []string) {
	// insertion point for field diffs
	if library.Name != libraryOther.Name {
		diffs = append(diffs, library.GongMarshallField(stage, "Name"))
	}
	if library.Description != libraryOther.Description {
		diffs = append(diffs, library.GongMarshallField(stage, "Description"))
	}
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
	}
	if library.IsExpanded != libraryOther.IsExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpanded"))
	}
	SubLibrariesDifferent := false
	if len(library.SubLibraries) != len(libraryOther.SubLibraries) {
		SubLibrariesDifferent = true
	} else {
		for i := range library.SubLibraries {
			if (library.SubLibraries[i] == nil) != (libraryOther.SubLibraries[i] == nil) {
				SubLibrariesDifferent = true
				break
			} else if library.SubLibraries[i] != nil && libraryOther.SubLibraries[i] != nil {
				// this is a pointer comparaison
				if library.SubLibraries[i] != libraryOther.SubLibraries[i] {
					SubLibrariesDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries)
		diffs = append(diffs, ops)
	}
	RootSystemsDifferent := false
	if len(library.RootSystems) != len(libraryOther.RootSystems) {
		RootSystemsDifferent = true
	} else {
		for i := range library.RootSystems {
			if (library.RootSystems[i] == nil) != (libraryOther.RootSystems[i] == nil) {
				RootSystemsDifferent = true
				break
			} else if library.RootSystems[i] != nil && libraryOther.RootSystems[i] != nil {
				// this is a pointer comparaison
				if library.RootSystems[i] != libraryOther.RootSystems[i] {
					RootSystemsDifferent = true
					break
				}
			}
		}
	}
	if RootSystemsDifferent {
		ops := Diff(stage, library, libraryOther, "RootSystems", libraryOther.RootSystems, library.RootSystems)
		diffs = append(diffs, ops)
	}
	RootComplexitysDifferent := false
	if len(library.RootComplexitys) != len(libraryOther.RootComplexitys) {
		RootComplexitysDifferent = true
	} else {
		for i := range library.RootComplexitys {
			if (library.RootComplexitys[i] == nil) != (libraryOther.RootComplexitys[i] == nil) {
				RootComplexitysDifferent = true
				break
			} else if library.RootComplexitys[i] != nil && libraryOther.RootComplexitys[i] != nil {
				// this is a pointer comparaison
				if library.RootComplexitys[i] != libraryOther.RootComplexitys[i] {
					RootComplexitysDifferent = true
					break
				}
			}
		}
	}
	if RootComplexitysDifferent {
		ops := Diff(stage, library, libraryOther, "RootComplexitys", libraryOther.RootComplexitys, library.RootComplexitys)
		diffs = append(diffs, ops)
	}
	RootPerformancesDifferent := false
	if len(library.RootPerformances) != len(libraryOther.RootPerformances) {
		RootPerformancesDifferent = true
	} else {
		for i := range library.RootPerformances {
			if (library.RootPerformances[i] == nil) != (libraryOther.RootPerformances[i] == nil) {
				RootPerformancesDifferent = true
				break
			} else if library.RootPerformances[i] != nil && libraryOther.RootPerformances[i] != nil {
				// this is a pointer comparaison
				if library.RootPerformances[i] != libraryOther.RootPerformances[i] {
					RootPerformancesDifferent = true
					break
				}
			}
		}
	}
	if RootPerformancesDifferent {
		ops := Diff(stage, library, libraryOther, "RootPerformances", libraryOther.RootPerformances, library.RootPerformances)
		diffs = append(diffs, ops)
	}
	RootEffortsDifferent := false
	if len(library.RootEfforts) != len(libraryOther.RootEfforts) {
		RootEffortsDifferent = true
	} else {
		for i := range library.RootEfforts {
			if (library.RootEfforts[i] == nil) != (libraryOther.RootEfforts[i] == nil) {
				RootEffortsDifferent = true
				break
			} else if library.RootEfforts[i] != nil && libraryOther.RootEfforts[i] != nil {
				// this is a pointer comparaison
				if library.RootEfforts[i] != libraryOther.RootEfforts[i] {
					RootEffortsDifferent = true
					break
				}
			}
		}
	}
	if RootEffortsDifferent {
		ops := Diff(stage, library, libraryOther, "RootEfforts", libraryOther.RootEfforts, library.RootEfforts)
		diffs = append(diffs, ops)
	}
	RootCompareAnalysisDifferent := false
	if len(library.RootCompareAnalysis) != len(libraryOther.RootCompareAnalysis) {
		RootCompareAnalysisDifferent = true
	} else {
		for i := range library.RootCompareAnalysis {
			if (library.RootCompareAnalysis[i] == nil) != (libraryOther.RootCompareAnalysis[i] == nil) {
				RootCompareAnalysisDifferent = true
				break
			} else if library.RootCompareAnalysis[i] != nil && libraryOther.RootCompareAnalysis[i] != nil {
				// this is a pointer comparaison
				if library.RootCompareAnalysis[i] != libraryOther.RootCompareAnalysis[i] {
					RootCompareAnalysisDifferent = true
					break
				}
			}
		}
	}
	if RootCompareAnalysisDifferent {
		ops := Diff(stage, library, libraryOther, "RootCompareAnalysis", libraryOther.RootCompareAnalysis, library.RootCompareAnalysis)
		diffs = append(diffs, ops)
	}
	RootNotesDifferent := false
	if len(library.RootNotes) != len(libraryOther.RootNotes) {
		RootNotesDifferent = true
	} else {
		for i := range library.RootNotes {
			if (library.RootNotes[i] == nil) != (libraryOther.RootNotes[i] == nil) {
				RootNotesDifferent = true
				break
			} else if library.RootNotes[i] != nil && libraryOther.RootNotes[i] != nil {
				// this is a pointer comparaison
				if library.RootNotes[i] != libraryOther.RootNotes[i] {
					RootNotesDifferent = true
					break
				}
			}
		}
	}
	if RootNotesDifferent {
		ops := Diff(stage, library, libraryOther, "RootNotes", libraryOther.RootNotes, library.RootNotes)
		diffs = append(diffs, ops)
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	SubLibrariesWhoseNodeIsExpandedDifferent := false
	if len(library.SubLibrariesWhoseNodeIsExpanded) != len(libraryOther.SubLibrariesWhoseNodeIsExpanded) {
		SubLibrariesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.SubLibrariesWhoseNodeIsExpanded {
			if (library.SubLibrariesWhoseNodeIsExpanded[i] == nil) != (libraryOther.SubLibrariesWhoseNodeIsExpanded[i] == nil) {
				SubLibrariesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.SubLibrariesWhoseNodeIsExpanded[i] != nil && libraryOther.SubLibrariesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.SubLibrariesWhoseNodeIsExpanded[i] != libraryOther.SubLibrariesWhoseNodeIsExpanded[i] {
					SubLibrariesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	if library.IsSystemsNodeExpanded != libraryOther.IsSystemsNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSystemsNodeExpanded"))
	}
	SystemsWhoseNodeIsExpandedDifferent := false
	if len(library.SystemsWhoseNodeIsExpanded) != len(libraryOther.SystemsWhoseNodeIsExpanded) {
		SystemsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.SystemsWhoseNodeIsExpanded {
			if (library.SystemsWhoseNodeIsExpanded[i] == nil) != (libraryOther.SystemsWhoseNodeIsExpanded[i] == nil) {
				SystemsWhoseNodeIsExpandedDifferent = true
				break
			} else if library.SystemsWhoseNodeIsExpanded[i] != nil && libraryOther.SystemsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.SystemsWhoseNodeIsExpanded[i] != libraryOther.SystemsWhoseNodeIsExpanded[i] {
					SystemsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SystemsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "SystemsWhoseNodeIsExpanded", libraryOther.SystemsWhoseNodeIsExpanded, library.SystemsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsComplexitysNodeExpanded != libraryOther.IsComplexitysNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
	}
	ComplexitysWhoseNodeIsExpandedDifferent := false
	if len(library.ComplexitysWhoseNodeIsExpanded) != len(libraryOther.ComplexitysWhoseNodeIsExpanded) {
		ComplexitysWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.ComplexitysWhoseNodeIsExpanded {
			if (library.ComplexitysWhoseNodeIsExpanded[i] == nil) != (libraryOther.ComplexitysWhoseNodeIsExpanded[i] == nil) {
				ComplexitysWhoseNodeIsExpandedDifferent = true
				break
			} else if library.ComplexitysWhoseNodeIsExpanded[i] != nil && libraryOther.ComplexitysWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.ComplexitysWhoseNodeIsExpanded[i] != libraryOther.ComplexitysWhoseNodeIsExpanded[i] {
					ComplexitysWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ComplexitysWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "ComplexitysWhoseNodeIsExpanded", libraryOther.ComplexitysWhoseNodeIsExpanded, library.ComplexitysWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsPerformancesNodeExpanded != libraryOther.IsPerformancesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
	}
	PerformancesWhoseNodeIsExpandedDifferent := false
	if len(library.PerformancesWhoseNodeIsExpanded) != len(libraryOther.PerformancesWhoseNodeIsExpanded) {
		PerformancesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.PerformancesWhoseNodeIsExpanded {
			if (library.PerformancesWhoseNodeIsExpanded[i] == nil) != (libraryOther.PerformancesWhoseNodeIsExpanded[i] == nil) {
				PerformancesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.PerformancesWhoseNodeIsExpanded[i] != nil && libraryOther.PerformancesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.PerformancesWhoseNodeIsExpanded[i] != libraryOther.PerformancesWhoseNodeIsExpanded[i] {
					PerformancesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PerformancesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "PerformancesWhoseNodeIsExpanded", libraryOther.PerformancesWhoseNodeIsExpanded, library.PerformancesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsEffortsNodeExpanded != libraryOther.IsEffortsNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsEffortsNodeExpanded"))
	}
	EffortsWhoseNodeIsExpandedDifferent := false
	if len(library.EffortsWhoseNodeIsExpanded) != len(libraryOther.EffortsWhoseNodeIsExpanded) {
		EffortsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.EffortsWhoseNodeIsExpanded {
			if (library.EffortsWhoseNodeIsExpanded[i] == nil) != (libraryOther.EffortsWhoseNodeIsExpanded[i] == nil) {
				EffortsWhoseNodeIsExpandedDifferent = true
				break
			} else if library.EffortsWhoseNodeIsExpanded[i] != nil && libraryOther.EffortsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.EffortsWhoseNodeIsExpanded[i] != libraryOther.EffortsWhoseNodeIsExpanded[i] {
					EffortsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if EffortsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "EffortsWhoseNodeIsExpanded", libraryOther.EffortsWhoseNodeIsExpanded, library.EffortsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsCompareAnalysisNodeExpanded != libraryOther.IsCompareAnalysisNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsCompareAnalysisNodeExpanded"))
	}
	CompareAnalysisWhoseNodeIsExpandedDifferent := false
	if len(library.CompareAnalysisWhoseNodeIsExpanded) != len(libraryOther.CompareAnalysisWhoseNodeIsExpanded) {
		CompareAnalysisWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.CompareAnalysisWhoseNodeIsExpanded {
			if (library.CompareAnalysisWhoseNodeIsExpanded[i] == nil) != (libraryOther.CompareAnalysisWhoseNodeIsExpanded[i] == nil) {
				CompareAnalysisWhoseNodeIsExpandedDifferent = true
				break
			} else if library.CompareAnalysisWhoseNodeIsExpanded[i] != nil && libraryOther.CompareAnalysisWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.CompareAnalysisWhoseNodeIsExpanded[i] != libraryOther.CompareAnalysisWhoseNodeIsExpanded[i] {
					CompareAnalysisWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if CompareAnalysisWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "CompareAnalysisWhoseNodeIsExpanded", libraryOther.CompareAnalysisWhoseNodeIsExpanded, library.CompareAnalysisWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsNotesNodeExpanded != libraryOther.IsNotesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(library.NotesWhoseNodeIsExpanded) != len(libraryOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.NotesWhoseNodeIsExpanded {
			if (library.NotesWhoseNodeIsExpanded[i] == nil) != (libraryOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.NotesWhoseNodeIsExpanded[i] != nil && libraryOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.NotesWhoseNodeIsExpanded[i] != libraryOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, library, libraryOther, "NotesWhoseNodeIsExpanded", libraryOther.NotesWhoseNodeIsExpanded, library.NotesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if note.Description != noteOther.Description {
		diffs = append(diffs, note.GongMarshallField(stage, "Description"))
	}
	ComplexitiesDifferent := false
	if len(note.Complexities) != len(noteOther.Complexities) {
		ComplexitiesDifferent = true
	} else {
		for i := range note.Complexities {
			if (note.Complexities[i] == nil) != (noteOther.Complexities[i] == nil) {
				ComplexitiesDifferent = true
				break
			} else if note.Complexities[i] != nil && noteOther.Complexities[i] != nil {
				// this is a pointer comparaison
				if note.Complexities[i] != noteOther.Complexities[i] {
					ComplexitiesDifferent = true
					break
				}
			}
		}
	}
	if ComplexitiesDifferent {
		ops := Diff(stage, note, noteOther, "Complexities", noteOther.Complexities, note.Complexities)
		diffs = append(diffs, ops)
	}
	PerformancesDifferent := false
	if len(note.Performances) != len(noteOther.Performances) {
		PerformancesDifferent = true
	} else {
		for i := range note.Performances {
			if (note.Performances[i] == nil) != (noteOther.Performances[i] == nil) {
				PerformancesDifferent = true
				break
			} else if note.Performances[i] != nil && noteOther.Performances[i] != nil {
				// this is a pointer comparaison
				if note.Performances[i] != noteOther.Performances[i] {
					PerformancesDifferent = true
					break
				}
			}
		}
	}
	if PerformancesDifferent {
		ops := Diff(stage, note, noteOther, "Performances", noteOther.Performances, note.Performances)
		diffs = append(diffs, ops)
	}
	EffortsDifferent := false
	if len(note.Efforts) != len(noteOther.Efforts) {
		EffortsDifferent = true
	} else {
		for i := range note.Efforts {
			if (note.Efforts[i] == nil) != (noteOther.Efforts[i] == nil) {
				EffortsDifferent = true
				break
			} else if note.Efforts[i] != nil && noteOther.Efforts[i] != nil {
				// this is a pointer comparaison
				if note.Efforts[i] != noteOther.Efforts[i] {
					EffortsDifferent = true
					break
				}
			}
		}
	}
	if EffortsDifferent {
		ops := Diff(stage, note, noteOther, "Efforts", noteOther.Efforts, note.Efforts)
		diffs = append(diffs, ops)
	}
	if note.ComputedPrefix != noteOther.ComputedPrefix {
		diffs = append(diffs, note.GongMarshallField(stage, "ComputedPrefix"))
	}
	if note.IsExpanded != noteOther.IsExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsExpanded"))
	}
	if note.IsComplexitysNodeExpanded != noteOther.IsComplexitysNodeExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
	}
	if note.IsPerformancesNodeExpanded != noteOther.IsPerformancesNodeExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
	}
	if note.IsEffortsNodeExpanded != noteOther.IsEffortsNodeExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsEffortsNodeExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notecomplexityshape *NoteComplexityShape) GongDiff(stage *Stage, notecomplexityshapeOther *NoteComplexityShape) (diffs []string) {
	// insertion point for field diffs
	if notecomplexityshape.Name != notecomplexityshapeOther.Name {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "Name"))
	}
	if (notecomplexityshape.Note == nil) != (notecomplexityshapeOther.Note == nil) {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "Note"))
	} else if notecomplexityshape.Note != nil && notecomplexityshapeOther.Note != nil {
		if notecomplexityshape.Note != notecomplexityshapeOther.Note {
			diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "Note"))
		}
	}
	if (notecomplexityshape.Complexity == nil) != (notecomplexityshapeOther.Complexity == nil) {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "Complexity"))
	} else if notecomplexityshape.Complexity != nil && notecomplexityshapeOther.Complexity != nil {
		if notecomplexityshape.Complexity != notecomplexityshapeOther.Complexity {
			diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "Complexity"))
		}
	}
	if notecomplexityshape.StartRatio != notecomplexityshapeOther.StartRatio {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "StartRatio"))
	}
	if notecomplexityshape.EndRatio != notecomplexityshapeOther.EndRatio {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "EndRatio"))
	}
	if notecomplexityshape.StartOrientation != notecomplexityshapeOther.StartOrientation {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "StartOrientation"))
	}
	if notecomplexityshape.EndOrientation != notecomplexityshapeOther.EndOrientation {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "EndOrientation"))
	}
	if notecomplexityshape.CornerOffsetRatio != notecomplexityshapeOther.CornerOffsetRatio {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notecomplexityshape.IsHidden != notecomplexityshapeOther.IsHidden {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteeffortshape *NoteEffortShape) GongDiff(stage *Stage, noteeffortshapeOther *NoteEffortShape) (diffs []string) {
	// insertion point for field diffs
	if noteeffortshape.Name != noteeffortshapeOther.Name {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "Name"))
	}
	if (noteeffortshape.Note == nil) != (noteeffortshapeOther.Note == nil) {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "Note"))
	} else if noteeffortshape.Note != nil && noteeffortshapeOther.Note != nil {
		if noteeffortshape.Note != noteeffortshapeOther.Note {
			diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "Note"))
		}
	}
	if (noteeffortshape.Effort == nil) != (noteeffortshapeOther.Effort == nil) {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "Effort"))
	} else if noteeffortshape.Effort != nil && noteeffortshapeOther.Effort != nil {
		if noteeffortshape.Effort != noteeffortshapeOther.Effort {
			diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "Effort"))
		}
	}
	if noteeffortshape.StartRatio != noteeffortshapeOther.StartRatio {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "StartRatio"))
	}
	if noteeffortshape.EndRatio != noteeffortshapeOther.EndRatio {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "EndRatio"))
	}
	if noteeffortshape.StartOrientation != noteeffortshapeOther.StartOrientation {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "StartOrientation"))
	}
	if noteeffortshape.EndOrientation != noteeffortshapeOther.EndOrientation {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "EndOrientation"))
	}
	if noteeffortshape.CornerOffsetRatio != noteeffortshapeOther.CornerOffsetRatio {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if noteeffortshape.IsHidden != noteeffortshapeOther.IsHidden {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteperformanceshape *NotePerformanceShape) GongDiff(stage *Stage, noteperformanceshapeOther *NotePerformanceShape) (diffs []string) {
	// insertion point for field diffs
	if noteperformanceshape.Name != noteperformanceshapeOther.Name {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "Name"))
	}
	if (noteperformanceshape.Note == nil) != (noteperformanceshapeOther.Note == nil) {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "Note"))
	} else if noteperformanceshape.Note != nil && noteperformanceshapeOther.Note != nil {
		if noteperformanceshape.Note != noteperformanceshapeOther.Note {
			diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "Note"))
		}
	}
	if (noteperformanceshape.Performance == nil) != (noteperformanceshapeOther.Performance == nil) {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "Performance"))
	} else if noteperformanceshape.Performance != nil && noteperformanceshapeOther.Performance != nil {
		if noteperformanceshape.Performance != noteperformanceshapeOther.Performance {
			diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "Performance"))
		}
	}
	if noteperformanceshape.StartRatio != noteperformanceshapeOther.StartRatio {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "StartRatio"))
	}
	if noteperformanceshape.EndRatio != noteperformanceshapeOther.EndRatio {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "EndRatio"))
	}
	if noteperformanceshape.StartOrientation != noteperformanceshapeOther.StartOrientation {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "StartOrientation"))
	}
	if noteperformanceshape.EndOrientation != noteperformanceshapeOther.EndOrientation {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "EndOrientation"))
	}
	if noteperformanceshape.CornerOffsetRatio != noteperformanceshapeOther.CornerOffsetRatio {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if noteperformanceshape.IsHidden != noteperformanceshapeOther.IsHidden {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteshape *NoteShape) GongDiff(stage *Stage, noteshapeOther *NoteShape) (diffs []string) {
	// insertion point for field diffs
	if noteshape.Name != noteshapeOther.Name {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Name"))
	}
	if (noteshape.Note == nil) != (noteshapeOther.Note == nil) {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
	} else if noteshape.Note != nil && noteshapeOther.Note != nil {
		if noteshape.Note != noteshapeOther.Note {
			diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
		}
	}
	if noteshape.X != noteshapeOther.X {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "X"))
	}
	if noteshape.Y != noteshapeOther.Y {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Y"))
	}
	if noteshape.Width != noteshapeOther.Width {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Width"))
	}
	if noteshape.Height != noteshapeOther.Height {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Height"))
	}
	if noteshape.IsHidden != noteshapeOther.IsHidden {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (performance *Performance) GongDiff(stage *Stage, performanceOther *Performance) (diffs []string) {
	// insertion point for field diffs
	if performance.Name != performanceOther.Name {
		diffs = append(diffs, performance.GongMarshallField(stage, "Name"))
	}
	if performance.Strength != performanceOther.Strength {
		diffs = append(diffs, performance.GongMarshallField(stage, "Strength"))
	}
	if performance.Description != performanceOther.Description {
		diffs = append(diffs, performance.GongMarshallField(stage, "Description"))
	}
	if performance.ComputedPrefix != performanceOther.ComputedPrefix {
		diffs = append(diffs, performance.GongMarshallField(stage, "ComputedPrefix"))
	}
	if performance.IsExpanded != performanceOther.IsExpanded {
		diffs = append(diffs, performance.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (system *System) GongDiff(stage *Stage, systemOther *System) (diffs []string) {
	// insertion point for field diffs
	if system.Name != systemOther.Name {
		diffs = append(diffs, system.GongMarshallField(stage, "Name"))
	}
	if system.Description != systemOther.Description {
		diffs = append(diffs, system.GongMarshallField(stage, "Description"))
	}
	ComplexitiesDifferent := false
	if len(system.Complexities) != len(systemOther.Complexities) {
		ComplexitiesDifferent = true
	} else {
		for i := range system.Complexities {
			if (system.Complexities[i] == nil) != (systemOther.Complexities[i] == nil) {
				ComplexitiesDifferent = true
				break
			} else if system.Complexities[i] != nil && systemOther.Complexities[i] != nil {
				// this is a pointer comparaison
				if system.Complexities[i] != systemOther.Complexities[i] {
					ComplexitiesDifferent = true
					break
				}
			}
		}
	}
	if ComplexitiesDifferent {
		ops := Diff(stage, system, systemOther, "Complexities", systemOther.Complexities, system.Complexities)
		diffs = append(diffs, ops)
	}
	PerformancesDifferent := false
	if len(system.Performances) != len(systemOther.Performances) {
		PerformancesDifferent = true
	} else {
		for i := range system.Performances {
			if (system.Performances[i] == nil) != (systemOther.Performances[i] == nil) {
				PerformancesDifferent = true
				break
			} else if system.Performances[i] != nil && systemOther.Performances[i] != nil {
				// this is a pointer comparaison
				if system.Performances[i] != systemOther.Performances[i] {
					PerformancesDifferent = true
					break
				}
			}
		}
	}
	if PerformancesDifferent {
		ops := Diff(stage, system, systemOther, "Performances", systemOther.Performances, system.Performances)
		diffs = append(diffs, ops)
	}
	EffortsDifferent := false
	if len(system.Efforts) != len(systemOther.Efforts) {
		EffortsDifferent = true
	} else {
		for i := range system.Efforts {
			if (system.Efforts[i] == nil) != (systemOther.Efforts[i] == nil) {
				EffortsDifferent = true
				break
			} else if system.Efforts[i] != nil && systemOther.Efforts[i] != nil {
				// this is a pointer comparaison
				if system.Efforts[i] != systemOther.Efforts[i] {
					EffortsDifferent = true
					break
				}
			}
		}
	}
	if EffortsDifferent {
		ops := Diff(stage, system, systemOther, "Efforts", systemOther.Efforts, system.Efforts)
		diffs = append(diffs, ops)
	}
	SubSystemsDifferent := false
	if len(system.SubSystems) != len(systemOther.SubSystems) {
		SubSystemsDifferent = true
	} else {
		for i := range system.SubSystems {
			if (system.SubSystems[i] == nil) != (systemOther.SubSystems[i] == nil) {
				SubSystemsDifferent = true
				break
			} else if system.SubSystems[i] != nil && systemOther.SubSystems[i] != nil {
				// this is a pointer comparaison
				if system.SubSystems[i] != systemOther.SubSystems[i] {
					SubSystemsDifferent = true
					break
				}
			}
		}
	}
	if SubSystemsDifferent {
		ops := Diff(stage, system, systemOther, "SubSystems", systemOther.SubSystems, system.SubSystems)
		diffs = append(diffs, ops)
	}
	if system.AreCPEsCompoundedFromSubSystems != systemOther.AreCPEsCompoundedFromSubSystems {
		diffs = append(diffs, system.GongMarshallField(stage, "AreCPEsCompoundedFromSubSystems"))
	}
	if system.ComputedPrefix != systemOther.ComputedPrefix {
		diffs = append(diffs, system.GongMarshallField(stage, "ComputedPrefix"))
	}
	if system.IsExpanded != systemOther.IsExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsExpanded"))
	}
	if system.SVG_Path != systemOther.SVG_Path {
		diffs = append(diffs, system.GongMarshallField(stage, "SVG_Path"))
	}
	if system.InverseAppliedScaling != systemOther.InverseAppliedScaling {
		diffs = append(diffs, system.GongMarshallField(stage, "InverseAppliedScaling"))
	}
	DiagramFlossEquationsDifferent := false
	if len(system.DiagramFlossEquations) != len(systemOther.DiagramFlossEquations) {
		DiagramFlossEquationsDifferent = true
	} else {
		for i := range system.DiagramFlossEquations {
			if (system.DiagramFlossEquations[i] == nil) != (systemOther.DiagramFlossEquations[i] == nil) {
				DiagramFlossEquationsDifferent = true
				break
			} else if system.DiagramFlossEquations[i] != nil && systemOther.DiagramFlossEquations[i] != nil {
				// this is a pointer comparaison
				if system.DiagramFlossEquations[i] != systemOther.DiagramFlossEquations[i] {
					DiagramFlossEquationsDifferent = true
					break
				}
			}
		}
	}
	if DiagramFlossEquationsDifferent {
		ops := Diff(stage, system, systemOther, "DiagramFlossEquations", systemOther.DiagramFlossEquations, system.DiagramFlossEquations)
		diffs = append(diffs, ops)
	}
	DiagramFlossEquationsWhoseNodeIsExpandedDifferent := false
	if len(system.DiagramFlossEquationsWhoseNodeIsExpanded) != len(systemOther.DiagramFlossEquationsWhoseNodeIsExpanded) {
		DiagramFlossEquationsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
			if (system.DiagramFlossEquationsWhoseNodeIsExpanded[i] == nil) != (systemOther.DiagramFlossEquationsWhoseNodeIsExpanded[i] == nil) {
				DiagramFlossEquationsWhoseNodeIsExpandedDifferent = true
				break
			} else if system.DiagramFlossEquationsWhoseNodeIsExpanded[i] != nil && systemOther.DiagramFlossEquationsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.DiagramFlossEquationsWhoseNodeIsExpanded[i] != systemOther.DiagramFlossEquationsWhoseNodeIsExpanded[i] {
					DiagramFlossEquationsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DiagramFlossEquationsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "DiagramFlossEquationsWhoseNodeIsExpanded", systemOther.DiagramFlossEquationsWhoseNodeIsExpanded, system.DiagramFlossEquationsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if system.IsSubSystemNodeExpanded != systemOther.IsSubSystemNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
	}
	if system.IsComplexitysNodeExpanded != systemOther.IsComplexitysNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
	}
	ComplexitysWhoseNodeIsExpandedDifferent := false
	if len(system.ComplexitysWhoseNodeIsExpanded) != len(systemOther.ComplexitysWhoseNodeIsExpanded) {
		ComplexitysWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.ComplexitysWhoseNodeIsExpanded {
			if (system.ComplexitysWhoseNodeIsExpanded[i] == nil) != (systemOther.ComplexitysWhoseNodeIsExpanded[i] == nil) {
				ComplexitysWhoseNodeIsExpandedDifferent = true
				break
			} else if system.ComplexitysWhoseNodeIsExpanded[i] != nil && systemOther.ComplexitysWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.ComplexitysWhoseNodeIsExpanded[i] != systemOther.ComplexitysWhoseNodeIsExpanded[i] {
					ComplexitysWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ComplexitysWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "ComplexitysWhoseNodeIsExpanded", systemOther.ComplexitysWhoseNodeIsExpanded, system.ComplexitysWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if system.IsPerformancesNodeExpanded != systemOther.IsPerformancesNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
	}
	PerformancesWhoseNodeIsExpandedDifferent := false
	if len(system.PerformancesWhoseNodeIsExpanded) != len(systemOther.PerformancesWhoseNodeIsExpanded) {
		PerformancesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.PerformancesWhoseNodeIsExpanded {
			if (system.PerformancesWhoseNodeIsExpanded[i] == nil) != (systemOther.PerformancesWhoseNodeIsExpanded[i] == nil) {
				PerformancesWhoseNodeIsExpandedDifferent = true
				break
			} else if system.PerformancesWhoseNodeIsExpanded[i] != nil && systemOther.PerformancesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.PerformancesWhoseNodeIsExpanded[i] != systemOther.PerformancesWhoseNodeIsExpanded[i] {
					PerformancesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PerformancesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "PerformancesWhoseNodeIsExpanded", systemOther.PerformancesWhoseNodeIsExpanded, system.PerformancesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if system.IsEffortsNodeExpanded != systemOther.IsEffortsNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsEffortsNodeExpanded"))
	}
	EffortsWhoseNodeIsExpandedDifferent := false
	if len(system.EffortsWhoseNodeIsExpanded) != len(systemOther.EffortsWhoseNodeIsExpanded) {
		EffortsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.EffortsWhoseNodeIsExpanded {
			if (system.EffortsWhoseNodeIsExpanded[i] == nil) != (systemOther.EffortsWhoseNodeIsExpanded[i] == nil) {
				EffortsWhoseNodeIsExpandedDifferent = true
				break
			} else if system.EffortsWhoseNodeIsExpanded[i] != nil && systemOther.EffortsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.EffortsWhoseNodeIsExpanded[i] != systemOther.EffortsWhoseNodeIsExpanded[i] {
					EffortsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if EffortsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "EffortsWhoseNodeIsExpanded", systemOther.EffortsWhoseNodeIsExpanded, system.EffortsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
