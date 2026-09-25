// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (compareanalysis *CompareAnalysis) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CompareAnalysiss[compareanalysis]
	return ok
}

func (complexity *Complexity) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Complexitys[complexity]
	return ok
}

func (diagramflossequation *DiagramFlossEquation) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DiagramFlossEquations[diagramflossequation]
	return ok
}

func (effort *Effort) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Efforts[effort]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (notecomplexityshape *NoteComplexityShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteComplexityShapes[notecomplexityshape]
	return ok
}

func (noteeffortshape *NoteEffortShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteEffortShapes[noteeffortshape]
	return ok
}

func (noteperformanceshape *NotePerformanceShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NotePerformanceShapes[noteperformanceshape]
	return ok
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteShapes[noteshape]
	return ok
}

func (performance *Performance) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Performances[performance]
	return ok
}

func (system *System) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Systems[system]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (compareanalysis *CompareAnalysis) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(compareanalysis) {
		return
	}

	compareanalysis.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if compareanalysis.FromSystem != nil {
		stage.StageBranch(compareanalysis.FromSystem)
	}
	if compareanalysis.ToSystem != nil {
		stage.StageBranch(compareanalysis.ToSystem)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramflossequation := range compareanalysis.DiagramFlossEquations {
		stage.StageBranch(_diagramflossequation)
	}
	for _, _diagramflossequation := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
		stage.StageBranch(_diagramflossequation)
	}

}

func (complexity *Complexity) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(complexity) {
		return
	}

	complexity.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagramflossequation *DiagramFlossEquation) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(diagramflossequation) {
		return
	}

	diagramflossequation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshape := range diagramflossequation.Note_Shapes {
		stage.StageBranch(_noteshape)
	}
	for _, _notecomplexityshape := range diagramflossequation.NoteComplexityShapes {
		stage.StageBranch(_notecomplexityshape)
	}
	for _, _noteperformanceshape := range diagramflossequation.NotePerformanceShapes {
		stage.StageBranch(_noteperformanceshape)
	}
	for _, _noteeffortshape := range diagramflossequation.NoteEffortShapes {
		stage.StageBranch(_noteeffortshape)
	}
	for _, _note := range diagramflossequation.NotesWhoseNodeIsExpanded {
		stage.StageBranch(_note)
	}
	for _, _complexity := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
		stage.StageBranch(_complexity)
	}
	for _, _performance := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
		stage.StageBranch(_performance)
	}
	for _, _effort := range diagramflossequation.EffortsWhoseNodeIsExpanded {
		stage.StageBranch(_effort)
	}

}

func (effort *Effort) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(effort) {
		return
	}

	effort.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.StageBranch(_library)
	}
	for _, _system := range library.RootSystems {
		stage.StageBranch(_system)
	}
	for _, _complexity := range library.RootComplexitys {
		stage.StageBranch(_complexity)
	}
	for _, _performance := range library.RootPerformances {
		stage.StageBranch(_performance)
	}
	for _, _effort := range library.RootEfforts {
		stage.StageBranch(_effort)
	}
	for _, _compareanalysis := range library.RootCompareAnalysis {
		stage.StageBranch(_compareanalysis)
	}
	for _, _note := range library.RootNotes {
		stage.StageBranch(_note)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.StageBranch(_library)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		stage.StageBranch(_system)
	}
	for _, _complexity := range library.ComplexitysWhoseNodeIsExpanded {
		stage.StageBranch(_complexity)
	}
	for _, _performance := range library.PerformancesWhoseNodeIsExpanded {
		stage.StageBranch(_performance)
	}
	for _, _effort := range library.EffortsWhoseNodeIsExpanded {
		stage.StageBranch(_effort)
	}
	for _, _compareanalysis := range library.CompareAnalysisWhoseNodeIsExpanded {
		stage.StageBranch(_compareanalysis)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		stage.StageBranch(_note)
	}

}

func (note *Note) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range note.Complexities {
		stage.StageBranch(_complexity)
	}
	for _, _performance := range note.Performances {
		stage.StageBranch(_performance)
	}
	for _, _effort := range note.Efforts {
		stage.StageBranch(_effort)
	}

}

func (notecomplexityshape *NoteComplexityShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(notecomplexityshape) {
		return
	}

	notecomplexityshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notecomplexityshape.Note != nil {
		stage.StageBranch(notecomplexityshape.Note)
	}
	if notecomplexityshape.Complexity != nil {
		stage.StageBranch(notecomplexityshape.Complexity)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteeffortshape *NoteEffortShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(noteeffortshape) {
		return
	}

	noteeffortshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteeffortshape.Note != nil {
		stage.StageBranch(noteeffortshape.Note)
	}
	if noteeffortshape.Effort != nil {
		stage.StageBranch(noteeffortshape.Effort)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteperformanceshape *NotePerformanceShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(noteperformanceshape) {
		return
	}

	noteperformanceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteperformanceshape.Note != nil {
		stage.StageBranch(noteperformanceshape.Note)
	}
	if noteperformanceshape.Performance != nil {
		stage.StageBranch(noteperformanceshape.Performance)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteshape *NoteShape) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.StageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (performance *Performance) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(performance) {
		return
	}

	performance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (system *System) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(system) {
		return
	}

	system.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range system.Complexities {
		stage.StageBranch(_complexity)
	}
	for _, _performance := range system.Performances {
		stage.StageBranch(_performance)
	}
	for _, _effort := range system.Efforts {
		stage.StageBranch(_effort)
	}
	for _, _system := range system.SubSystems {
		stage.StageBranch(_system)
	}
	for _, _diagramflossequation := range system.DiagramFlossEquations {
		stage.StageBranch(_diagramflossequation)
	}
	for _, _diagramflossequation := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
		stage.StageBranch(_diagramflossequation)
	}
	for _, _complexity := range system.ComplexitysWhoseNodeIsExpanded {
		stage.StageBranch(_complexity)
	}
	for _, _performance := range system.PerformancesWhoseNodeIsExpanded {
		stage.StageBranch(_performance)
	}
	for _, _effort := range system.EffortsWhoseNodeIsExpanded {
		stage.StageBranch(_effort)
	}

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *CompareAnalysis:
		toT := GongCopyBranchCompareAnalysis(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Complexity:
		toT := GongCopyBranchComplexity(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramFlossEquation:
		toT := GongCopyBranchDiagramFlossEquation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Effort:
		toT := GongCopyBranchEffort(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteComplexityShape:
		toT := GongCopyBranchNoteComplexityShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteEffortShape:
		toT := GongCopyBranchNoteEffortShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NotePerformanceShape:
		toT := GongCopyBranchNotePerformanceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := GongCopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Performance:
		toT := GongCopyBranchPerformance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System:
		toT := GongCopyBranchSystem(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCompareAnalysis(mapOrigCopy map[any]any, compareanalysisFrom *CompareAnalysis) (compareanalysisTo *CompareAnalysis) {
	var alreadyCopied bool
	compareanalysisTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, compareanalysisFrom)
	if alreadyCopied {
		return
	}
	compareanalysisFrom.GongCopyBasicFields(compareanalysisTo)

	//insertion point for the staging of instances referenced by pointers
	if compareanalysisFrom.FromSystem != nil {
		compareanalysisTo.FromSystem = GongCopyBranchSystem(mapOrigCopy, compareanalysisFrom.FromSystem)
	}
	if compareanalysisFrom.ToSystem != nil {
		compareanalysisTo.ToSystem = GongCopyBranchSystem(mapOrigCopy, compareanalysisFrom.ToSystem)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramflossequation := range compareanalysisFrom.DiagramFlossEquations {
		compareanalysisTo.DiagramFlossEquations = append(compareanalysisTo.DiagramFlossEquations, GongCopyBranchDiagramFlossEquation(mapOrigCopy, _diagramflossequation))
	}
	for _, _diagramflossequation := range compareanalysisFrom.DiagramFlossEquationsWhoseNodeIsExpanded {
		compareanalysisTo.DiagramFlossEquationsWhoseNodeIsExpanded = append(compareanalysisTo.DiagramFlossEquationsWhoseNodeIsExpanded, GongCopyBranchDiagramFlossEquation(mapOrigCopy, _diagramflossequation))
	}

	return
}

func GongCopyBranchComplexity(mapOrigCopy map[any]any, complexityFrom *Complexity) (complexityTo *Complexity) {
	var alreadyCopied bool
	complexityTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, complexityFrom)
	if alreadyCopied {
		return
	}
	complexityFrom.GongCopyBasicFields(complexityTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagramFlossEquation(mapOrigCopy map[any]any, diagramflossequationFrom *DiagramFlossEquation) (diagramflossequationTo *DiagramFlossEquation) {
	var alreadyCopied bool
	diagramflossequationTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramflossequationFrom)
	if alreadyCopied {
		return
	}
	diagramflossequationFrom.GongCopyBasicFields(diagramflossequationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshape := range diagramflossequationFrom.Note_Shapes {
		diagramflossequationTo.Note_Shapes = append(diagramflossequationTo.Note_Shapes, GongCopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _notecomplexityshape := range diagramflossequationFrom.NoteComplexityShapes {
		diagramflossequationTo.NoteComplexityShapes = append(diagramflossequationTo.NoteComplexityShapes, GongCopyBranchNoteComplexityShape(mapOrigCopy, _notecomplexityshape))
	}
	for _, _noteperformanceshape := range diagramflossequationFrom.NotePerformanceShapes {
		diagramflossequationTo.NotePerformanceShapes = append(diagramflossequationTo.NotePerformanceShapes, GongCopyBranchNotePerformanceShape(mapOrigCopy, _noteperformanceshape))
	}
	for _, _noteeffortshape := range diagramflossequationFrom.NoteEffortShapes {
		diagramflossequationTo.NoteEffortShapes = append(diagramflossequationTo.NoteEffortShapes, GongCopyBranchNoteEffortShape(mapOrigCopy, _noteeffortshape))
	}
	for _, _note := range diagramflossequationFrom.NotesWhoseNodeIsExpanded {
		diagramflossequationTo.NotesWhoseNodeIsExpanded = append(diagramflossequationTo.NotesWhoseNodeIsExpanded, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _complexity := range diagramflossequationFrom.ComplexitysWhoseNodeIsExpanded {
		diagramflossequationTo.ComplexitysWhoseNodeIsExpanded = append(diagramflossequationTo.ComplexitysWhoseNodeIsExpanded, GongCopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range diagramflossequationFrom.PerformancesWhoseNodeIsExpanded {
		diagramflossequationTo.PerformancesWhoseNodeIsExpanded = append(diagramflossequationTo.PerformancesWhoseNodeIsExpanded, GongCopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range diagramflossequationFrom.EffortsWhoseNodeIsExpanded {
		diagramflossequationTo.EffortsWhoseNodeIsExpanded = append(diagramflossequationTo.EffortsWhoseNodeIsExpanded, GongCopyBranchEffort(mapOrigCopy, _effort))
	}

	return
}

func GongCopyBranchEffort(mapOrigCopy map[any]any, effortFrom *Effort) (effortTo *Effort) {
	var alreadyCopied bool
	effortTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, effortFrom)
	if alreadyCopied {
		return
	}
	effortFrom.GongCopyBasicFields(effortTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {
	var alreadyCopied bool
	libraryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, libraryFrom)
	if alreadyCopied {
		return
	}
	libraryFrom.GongCopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _system := range libraryFrom.RootSystems {
		libraryTo.RootSystems = append(libraryTo.RootSystems, GongCopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _complexity := range libraryFrom.RootComplexitys {
		libraryTo.RootComplexitys = append(libraryTo.RootComplexitys, GongCopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range libraryFrom.RootPerformances {
		libraryTo.RootPerformances = append(libraryTo.RootPerformances, GongCopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range libraryFrom.RootEfforts {
		libraryTo.RootEfforts = append(libraryTo.RootEfforts, GongCopyBranchEffort(mapOrigCopy, _effort))
	}
	for _, _compareanalysis := range libraryFrom.RootCompareAnalysis {
		libraryTo.RootCompareAnalysis = append(libraryTo.RootCompareAnalysis, GongCopyBranchCompareAnalysis(mapOrigCopy, _compareanalysis))
	}
	for _, _note := range libraryFrom.RootNotes {
		libraryTo.RootNotes = append(libraryTo.RootNotes, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _system := range libraryFrom.SystemsWhoseNodeIsExpanded {
		libraryTo.SystemsWhoseNodeIsExpanded = append(libraryTo.SystemsWhoseNodeIsExpanded, GongCopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _complexity := range libraryFrom.ComplexitysWhoseNodeIsExpanded {
		libraryTo.ComplexitysWhoseNodeIsExpanded = append(libraryTo.ComplexitysWhoseNodeIsExpanded, GongCopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range libraryFrom.PerformancesWhoseNodeIsExpanded {
		libraryTo.PerformancesWhoseNodeIsExpanded = append(libraryTo.PerformancesWhoseNodeIsExpanded, GongCopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range libraryFrom.EffortsWhoseNodeIsExpanded {
		libraryTo.EffortsWhoseNodeIsExpanded = append(libraryTo.EffortsWhoseNodeIsExpanded, GongCopyBranchEffort(mapOrigCopy, _effort))
	}
	for _, _compareanalysis := range libraryFrom.CompareAnalysisWhoseNodeIsExpanded {
		libraryTo.CompareAnalysisWhoseNodeIsExpanded = append(libraryTo.CompareAnalysisWhoseNodeIsExpanded, GongCopyBranchCompareAnalysis(mapOrigCopy, _compareanalysis))
	}
	for _, _note := range libraryFrom.NotesWhoseNodeIsExpanded {
		libraryTo.NotesWhoseNodeIsExpanded = append(libraryTo.NotesWhoseNodeIsExpanded, GongCopyBranchNote(mapOrigCopy, _note))
	}

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range noteFrom.Complexities {
		noteTo.Complexities = append(noteTo.Complexities, GongCopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range noteFrom.Performances {
		noteTo.Performances = append(noteTo.Performances, GongCopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range noteFrom.Efforts {
		noteTo.Efforts = append(noteTo.Efforts, GongCopyBranchEffort(mapOrigCopy, _effort))
	}

	return
}

func GongCopyBranchNoteComplexityShape(mapOrigCopy map[any]any, notecomplexityshapeFrom *NoteComplexityShape) (notecomplexityshapeTo *NoteComplexityShape) {
	var alreadyCopied bool
	notecomplexityshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notecomplexityshapeFrom)
	if alreadyCopied {
		return
	}
	notecomplexityshapeFrom.GongCopyBasicFields(notecomplexityshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notecomplexityshapeFrom.Note != nil {
		notecomplexityshapeTo.Note = GongCopyBranchNote(mapOrigCopy, notecomplexityshapeFrom.Note)
	}
	if notecomplexityshapeFrom.Complexity != nil {
		notecomplexityshapeTo.Complexity = GongCopyBranchComplexity(mapOrigCopy, notecomplexityshapeFrom.Complexity)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteEffortShape(mapOrigCopy map[any]any, noteeffortshapeFrom *NoteEffortShape) (noteeffortshapeTo *NoteEffortShape) {
	var alreadyCopied bool
	noteeffortshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteeffortshapeFrom)
	if alreadyCopied {
		return
	}
	noteeffortshapeFrom.GongCopyBasicFields(noteeffortshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteeffortshapeFrom.Note != nil {
		noteeffortshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteeffortshapeFrom.Note)
	}
	if noteeffortshapeFrom.Effort != nil {
		noteeffortshapeTo.Effort = GongCopyBranchEffort(mapOrigCopy, noteeffortshapeFrom.Effort)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNotePerformanceShape(mapOrigCopy map[any]any, noteperformanceshapeFrom *NotePerformanceShape) (noteperformanceshapeTo *NotePerformanceShape) {
	var alreadyCopied bool
	noteperformanceshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteperformanceshapeFrom)
	if alreadyCopied {
		return
	}
	noteperformanceshapeFrom.GongCopyBasicFields(noteperformanceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteperformanceshapeFrom.Note != nil {
		noteperformanceshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteperformanceshapeFrom.Note)
	}
	if noteperformanceshapeFrom.Performance != nil {
		noteperformanceshapeTo.Performance = GongCopyBranchPerformance(mapOrigCopy, noteperformanceshapeFrom.Performance)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {
	var alreadyCopied bool
	noteshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteshapeFrom)
	if alreadyCopied {
		return
	}
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPerformance(mapOrigCopy map[any]any, performanceFrom *Performance) (performanceTo *Performance) {
	var alreadyCopied bool
	performanceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, performanceFrom)
	if alreadyCopied {
		return
	}
	performanceFrom.GongCopyBasicFields(performanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSystem(mapOrigCopy map[any]any, systemFrom *System) (systemTo *System) {
	var alreadyCopied bool
	systemTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, systemFrom)
	if alreadyCopied {
		return
	}
	systemFrom.GongCopyBasicFields(systemTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range systemFrom.Complexities {
		systemTo.Complexities = append(systemTo.Complexities, GongCopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range systemFrom.Performances {
		systemTo.Performances = append(systemTo.Performances, GongCopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range systemFrom.Efforts {
		systemTo.Efforts = append(systemTo.Efforts, GongCopyBranchEffort(mapOrigCopy, _effort))
	}
	for _, _system := range systemFrom.SubSystems {
		systemTo.SubSystems = append(systemTo.SubSystems, GongCopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _diagramflossequation := range systemFrom.DiagramFlossEquations {
		systemTo.DiagramFlossEquations = append(systemTo.DiagramFlossEquations, GongCopyBranchDiagramFlossEquation(mapOrigCopy, _diagramflossequation))
	}
	for _, _diagramflossequation := range systemFrom.DiagramFlossEquationsWhoseNodeIsExpanded {
		systemTo.DiagramFlossEquationsWhoseNodeIsExpanded = append(systemTo.DiagramFlossEquationsWhoseNodeIsExpanded, GongCopyBranchDiagramFlossEquation(mapOrigCopy, _diagramflossequation))
	}
	for _, _complexity := range systemFrom.ComplexitysWhoseNodeIsExpanded {
		systemTo.ComplexitysWhoseNodeIsExpanded = append(systemTo.ComplexitysWhoseNodeIsExpanded, GongCopyBranchComplexity(mapOrigCopy, _complexity))
	}
	for _, _performance := range systemFrom.PerformancesWhoseNodeIsExpanded {
		systemTo.PerformancesWhoseNodeIsExpanded = append(systemTo.PerformancesWhoseNodeIsExpanded, GongCopyBranchPerformance(mapOrigCopy, _performance))
	}
	for _, _effort := range systemFrom.EffortsWhoseNodeIsExpanded {
		systemTo.EffortsWhoseNodeIsExpanded = append(systemTo.EffortsWhoseNodeIsExpanded, GongCopyBranchEffort(mapOrigCopy, _effort))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (compareanalysis *CompareAnalysis) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(compareanalysis) {
		return
	}

	compareanalysis.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if compareanalysis.FromSystem != nil {
		stage.UnstageBranch(compareanalysis.FromSystem)
	}
	if compareanalysis.ToSystem != nil {
		stage.UnstageBranch(compareanalysis.ToSystem)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramflossequation := range compareanalysis.DiagramFlossEquations {
		stage.UnstageBranch(_diagramflossequation)
	}
	for _, _diagramflossequation := range compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded {
		stage.UnstageBranch(_diagramflossequation)
	}

}

func (complexity *Complexity) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(complexity) {
		return
	}

	complexity.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagramflossequation *DiagramFlossEquation) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(diagramflossequation) {
		return
	}

	diagramflossequation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _noteshape := range diagramflossequation.Note_Shapes {
		stage.UnstageBranch(_noteshape)
	}
	for _, _notecomplexityshape := range diagramflossequation.NoteComplexityShapes {
		stage.UnstageBranch(_notecomplexityshape)
	}
	for _, _noteperformanceshape := range diagramflossequation.NotePerformanceShapes {
		stage.UnstageBranch(_noteperformanceshape)
	}
	for _, _noteeffortshape := range diagramflossequation.NoteEffortShapes {
		stage.UnstageBranch(_noteeffortshape)
	}
	for _, _note := range diagramflossequation.NotesWhoseNodeIsExpanded {
		stage.UnstageBranch(_note)
	}
	for _, _complexity := range diagramflossequation.ComplexitysWhoseNodeIsExpanded {
		stage.UnstageBranch(_complexity)
	}
	for _, _performance := range diagramflossequation.PerformancesWhoseNodeIsExpanded {
		stage.UnstageBranch(_performance)
	}
	for _, _effort := range diagramflossequation.EffortsWhoseNodeIsExpanded {
		stage.UnstageBranch(_effort)
	}

}

func (effort *Effort) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(effort) {
		return
	}

	effort.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (library *Library) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.UnstageBranch(_library)
	}
	for _, _system := range library.RootSystems {
		stage.UnstageBranch(_system)
	}
	for _, _complexity := range library.RootComplexitys {
		stage.UnstageBranch(_complexity)
	}
	for _, _performance := range library.RootPerformances {
		stage.UnstageBranch(_performance)
	}
	for _, _effort := range library.RootEfforts {
		stage.UnstageBranch(_effort)
	}
	for _, _compareanalysis := range library.RootCompareAnalysis {
		stage.UnstageBranch(_compareanalysis)
	}
	for _, _note := range library.RootNotes {
		stage.UnstageBranch(_note)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.UnstageBranch(_library)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		stage.UnstageBranch(_system)
	}
	for _, _complexity := range library.ComplexitysWhoseNodeIsExpanded {
		stage.UnstageBranch(_complexity)
	}
	for _, _performance := range library.PerformancesWhoseNodeIsExpanded {
		stage.UnstageBranch(_performance)
	}
	for _, _effort := range library.EffortsWhoseNodeIsExpanded {
		stage.UnstageBranch(_effort)
	}
	for _, _compareanalysis := range library.CompareAnalysisWhoseNodeIsExpanded {
		stage.UnstageBranch(_compareanalysis)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		stage.UnstageBranch(_note)
	}

}

func (note *Note) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range note.Complexities {
		stage.UnstageBranch(_complexity)
	}
	for _, _performance := range note.Performances {
		stage.UnstageBranch(_performance)
	}
	for _, _effort := range note.Efforts {
		stage.UnstageBranch(_effort)
	}

}

func (notecomplexityshape *NoteComplexityShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(notecomplexityshape) {
		return
	}

	notecomplexityshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notecomplexityshape.Note != nil {
		stage.UnstageBranch(notecomplexityshape.Note)
	}
	if notecomplexityshape.Complexity != nil {
		stage.UnstageBranch(notecomplexityshape.Complexity)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteeffortshape *NoteEffortShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(noteeffortshape) {
		return
	}

	noteeffortshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteeffortshape.Note != nil {
		stage.UnstageBranch(noteeffortshape.Note)
	}
	if noteeffortshape.Effort != nil {
		stage.UnstageBranch(noteeffortshape.Effort)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteperformanceshape *NotePerformanceShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(noteperformanceshape) {
		return
	}

	noteperformanceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteperformanceshape.Note != nil {
		stage.UnstageBranch(noteperformanceshape.Note)
	}
	if noteperformanceshape.Performance != nil {
		stage.UnstageBranch(noteperformanceshape.Performance)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteshape *NoteShape) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.UnstageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (performance *Performance) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(performance) {
		return
	}

	performance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (system *System) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(system) {
		return
	}

	system.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _complexity := range system.Complexities {
		stage.UnstageBranch(_complexity)
	}
	for _, _performance := range system.Performances {
		stage.UnstageBranch(_performance)
	}
	for _, _effort := range system.Efforts {
		stage.UnstageBranch(_effort)
	}
	for _, _system := range system.SubSystems {
		stage.UnstageBranch(_system)
	}
	for _, _diagramflossequation := range system.DiagramFlossEquations {
		stage.UnstageBranch(_diagramflossequation)
	}
	for _, _diagramflossequation := range system.DiagramFlossEquationsWhoseNodeIsExpanded {
		stage.UnstageBranch(_diagramflossequation)
	}
	for _, _complexity := range system.ComplexitysWhoseNodeIsExpanded {
		stage.UnstageBranch(_complexity)
	}
	for _, _performance := range system.PerformancesWhoseNodeIsExpanded {
		stage.UnstageBranch(_performance)
	}
	for _, _effort := range system.EffortsWhoseNodeIsExpanded {
		stage.UnstageBranch(_effort)
	}

}

// insertion point for pointer reconstruction from references
func (reference *CompareAnalysis) GongReconstructPointersFromReferences(stage *Stage, instance *CompareAnalysis) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.FromSystem, stage.Systems_reference, instance.FromSystem)
	__gong__reconstructPointer(&reference.ToSystem, stage.Systems_reference, instance.ToSystem)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramFlossEquations, stage.DiagramFlossEquations_reference, instance.DiagramFlossEquations)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramFlossEquationsWhoseNodeIsExpanded, stage.DiagramFlossEquations_reference, instance.DiagramFlossEquationsWhoseNodeIsExpanded)
}

func (reference *Complexity) GongReconstructPointersFromReferences(stage *Stage, instance *Complexity) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DiagramFlossEquation) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramFlossEquation) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note_Shapes, stage.NoteShapes_reference, instance.Note_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteComplexityShapes, stage.NoteComplexityShapes_reference, instance.NoteComplexityShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotePerformanceShapes, stage.NotePerformanceShapes_reference, instance.NotePerformanceShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteEffortShapes, stage.NoteEffortShapes_reference, instance.NoteEffortShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotesWhoseNodeIsExpanded, stage.Notes_reference, instance.NotesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_reference, instance.ComplexitysWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PerformancesWhoseNodeIsExpanded, stage.Performances_reference, instance.PerformancesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.EffortsWhoseNodeIsExpanded, stage.Efforts_reference, instance.EffortsWhoseNodeIsExpanded)
}

func (reference *Effort) GongReconstructPointersFromReferences(stage *Stage, instance *Effort) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootSystems, stage.Systems_reference, instance.RootSystems)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootComplexitys, stage.Complexitys_reference, instance.RootComplexitys)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootPerformances, stage.Performances_reference, instance.RootPerformances)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootEfforts, stage.Efforts_reference, instance.RootEfforts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootCompareAnalysis, stage.CompareAnalysiss_reference, instance.RootCompareAnalysis)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootNotes, stage.Notes_reference, instance.RootNotes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference, instance.SubLibrariesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference, instance.SystemsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_reference, instance.ComplexitysWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PerformancesWhoseNodeIsExpanded, stage.Performances_reference, instance.PerformancesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.EffortsWhoseNodeIsExpanded, stage.Efforts_reference, instance.EffortsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.CompareAnalysisWhoseNodeIsExpanded, stage.CompareAnalysiss_reference, instance.CompareAnalysisWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotesWhoseNodeIsExpanded, stage.Notes_reference, instance.NotesWhoseNodeIsExpanded)
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Complexities, stage.Complexitys_reference, instance.Complexities)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Performances, stage.Performances_reference, instance.Performances)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Efforts, stage.Efforts_reference, instance.Efforts)
}

func (reference *NoteComplexityShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteComplexityShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Complexity, stage.Complexitys_reference, instance.Complexity)
	// insertion point for slice of pointers field
}

func (reference *NoteEffortShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteEffortShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Effort, stage.Efforts_reference, instance.Effort)
	// insertion point for slice of pointers field
}

func (reference *NotePerformanceShape) GongReconstructPointersFromReferences(stage *Stage, instance *NotePerformanceShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Performance, stage.Performances_reference, instance.Performance)
	// insertion point for slice of pointers field
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	// insertion point for slice of pointers field
}

func (reference *Performance) GongReconstructPointersFromReferences(stage *Stage, instance *Performance) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *System) GongReconstructPointersFromReferences(stage *Stage, instance *System) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Complexities, stage.Complexitys_reference, instance.Complexities)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Performances, stage.Performances_reference, instance.Performances)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Efforts, stage.Efforts_reference, instance.Efforts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubSystems, stage.Systems_reference, instance.SubSystems)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramFlossEquations, stage.DiagramFlossEquations_reference, instance.DiagramFlossEquations)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramFlossEquationsWhoseNodeIsExpanded, stage.DiagramFlossEquations_reference, instance.DiagramFlossEquationsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_reference, instance.ComplexitysWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PerformancesWhoseNodeIsExpanded, stage.Performances_reference, instance.PerformancesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.EffortsWhoseNodeIsExpanded, stage.Efforts_reference, instance.EffortsWhoseNodeIsExpanded)
}

// insertion point for pointer reconstruction from instances
func (reference *CompareAnalysis) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.FromSystem, stage.Systems_instance)
	__gong__reconstructPointerFromInstance(&reference.ToSystem, stage.Systems_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramFlossEquations, stage.DiagramFlossEquations_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramFlossEquationsWhoseNodeIsExpanded, stage.DiagramFlossEquations_instance)
}

func (reference *Complexity) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DiagramFlossEquation) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note_Shapes, stage.NoteShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteComplexityShapes, stage.NoteComplexityShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotePerformanceShapes, stage.NotePerformanceShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteEffortShapes, stage.NoteEffortShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotesWhoseNodeIsExpanded, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PerformancesWhoseNodeIsExpanded, stage.Performances_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.EffortsWhoseNodeIsExpanded, stage.Efforts_instance)
}

func (reference *Effort) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootSystems, stage.Systems_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootComplexitys, stage.Complexitys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootPerformances, stage.Performances_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootEfforts, stage.Efforts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootCompareAnalysis, stage.CompareAnalysiss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootNotes, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SystemsWhoseNodeIsExpanded, stage.Systems_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PerformancesWhoseNodeIsExpanded, stage.Performances_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.EffortsWhoseNodeIsExpanded, stage.Efforts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.CompareAnalysisWhoseNodeIsExpanded, stage.CompareAnalysiss_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotesWhoseNodeIsExpanded, stage.Notes_instance)
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Complexities, stage.Complexitys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Performances, stage.Performances_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Efforts, stage.Efforts_instance)
}

func (reference *NoteComplexityShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Complexity, stage.Complexitys_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteEffortShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Effort, stage.Efforts_instance)
	// insertion point for slice of pointers fields
}

func (reference *NotePerformanceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Performance, stage.Performances_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Performance) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *System) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Complexities, stage.Complexitys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Performances, stage.Performances_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Efforts, stage.Efforts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubSystems, stage.Systems_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramFlossEquations, stage.DiagramFlossEquations_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramFlossEquationsWhoseNodeIsExpanded, stage.DiagramFlossEquations_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ComplexitysWhoseNodeIsExpanded, stage.Complexitys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PerformancesWhoseNodeIsExpanded, stage.Performances_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.EffortsWhoseNodeIsExpanded, stage.Efforts_instance)
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (compareanalysis *CompareAnalysis) GongDiff(stage *Stage, compareanalysisOther *CompareAnalysis) (diffs []string) {
	// insertion point for field diffs
	if compareanalysis.Name != compareanalysisOther.Name {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "Name"))
	}
	if compareanalysis.FromSystem != compareanalysisOther.FromSystem {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "FromSystem"))
	}
	if compareanalysis.ToSystem != compareanalysisOther.ToSystem {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "ToSystem"))
	}
	if compareanalysis.Mu != compareanalysisOther.Mu {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "Mu"))
	}
	if compareanalysis.Epsilon != compareanalysisOther.Epsilon {
		diffs = append(diffs, compareanalysis.GongMarshallField(stage, "Epsilon"))
	}
	if ops := __gong__diffSliceOfPointers(stage, compareanalysis, "DiagramFlossEquations", compareanalysisOther.DiagramFlossEquations, compareanalysis.DiagramFlossEquations); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, compareanalysis, "DiagramFlossEquationsWhoseNodeIsExpanded", compareanalysisOther.DiagramFlossEquationsWhoseNodeIsExpanded, compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded); ops != "" {
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
	if diagramflossequation.AreCommonElementsHidden != diagramflossequationOther.AreCommonElementsHidden {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "AreCommonElementsHidden"))
	}
	if diagramflossequation.AreCPEArrowsVisible != diagramflossequationOther.AreCPEArrowsVisible {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "AreCPEArrowsVisible"))
	}
	if diagramflossequation.AreColumnTitlesVisible != diagramflossequationOther.AreColumnTitlesVisible {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "AreColumnTitlesVisible"))
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
	if ops := __gong__diffSliceOfPointers(stage, diagramflossequation, "Note_Shapes", diagramflossequationOther.Note_Shapes, diagramflossequation.Note_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramflossequation, "NoteComplexityShapes", diagramflossequationOther.NoteComplexityShapes, diagramflossequation.NoteComplexityShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramflossequation, "NotePerformanceShapes", diagramflossequationOther.NotePerformanceShapes, diagramflossequation.NotePerformanceShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramflossequation, "NoteEffortShapes", diagramflossequationOther.NoteEffortShapes, diagramflossequation.NoteEffortShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramflossequation.IsNotesNodeExpanded != diagramflossequationOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramflossequation, "NotesWhoseNodeIsExpanded", diagramflossequationOther.NotesWhoseNodeIsExpanded, diagramflossequation.NotesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramflossequation.IsComplexitysNodeExpanded != diagramflossequationOther.IsComplexitysNodeExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramflossequation, "ComplexitysWhoseNodeIsExpanded", diagramflossequationOther.ComplexitysWhoseNodeIsExpanded, diagramflossequation.ComplexitysWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramflossequation.IsPerformancesNodeExpanded != diagramflossequationOther.IsPerformancesNodeExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramflossequation, "PerformancesWhoseNodeIsExpanded", diagramflossequationOther.PerformancesWhoseNodeIsExpanded, diagramflossequation.PerformancesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramflossequation.IsEffortsNodeExpanded != diagramflossequationOther.IsEffortsNodeExpanded {
		diffs = append(diffs, diagramflossequation.GongMarshallField(stage, "IsEffortsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramflossequation, "EffortsWhoseNodeIsExpanded", diagramflossequationOther.EffortsWhoseNodeIsExpanded, diagramflossequation.EffortsWhoseNodeIsExpanded); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootSystems", libraryOther.RootSystems, library.RootSystems); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootComplexitys", libraryOther.RootComplexitys, library.RootComplexitys); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootPerformances", libraryOther.RootPerformances, library.RootPerformances); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootEfforts", libraryOther.RootEfforts, library.RootEfforts); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootCompareAnalysis", libraryOther.RootCompareAnalysis, library.RootCompareAnalysis); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootNotes", libraryOther.RootNotes, library.RootNotes); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, library, "SystemsWhoseNodeIsExpanded", libraryOther.SystemsWhoseNodeIsExpanded, library.SystemsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsComplexitysNodeExpanded != libraryOther.IsComplexitysNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "ComplexitysWhoseNodeIsExpanded", libraryOther.ComplexitysWhoseNodeIsExpanded, library.ComplexitysWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsPerformancesNodeExpanded != libraryOther.IsPerformancesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "PerformancesWhoseNodeIsExpanded", libraryOther.PerformancesWhoseNodeIsExpanded, library.PerformancesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsEffortsNodeExpanded != libraryOther.IsEffortsNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsEffortsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "EffortsWhoseNodeIsExpanded", libraryOther.EffortsWhoseNodeIsExpanded, library.EffortsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsCompareAnalysisNodeExpanded != libraryOther.IsCompareAnalysisNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsCompareAnalysisNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "CompareAnalysisWhoseNodeIsExpanded", libraryOther.CompareAnalysisWhoseNodeIsExpanded, library.CompareAnalysisWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsNotesNodeExpanded != libraryOther.IsNotesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "NotesWhoseNodeIsExpanded", libraryOther.NotesWhoseNodeIsExpanded, library.NotesWhoseNodeIsExpanded); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, note, "Complexities", noteOther.Complexities, note.Complexities); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Performances", noteOther.Performances, note.Performances); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Efforts", noteOther.Efforts, note.Efforts); ops != "" {
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
	if notecomplexityshape.Note != notecomplexityshapeOther.Note {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "Note"))
	}
	if notecomplexityshape.Complexity != notecomplexityshapeOther.Complexity {
		diffs = append(diffs, notecomplexityshape.GongMarshallField(stage, "Complexity"))
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
	if noteeffortshape.Note != noteeffortshapeOther.Note {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "Note"))
	}
	if noteeffortshape.Effort != noteeffortshapeOther.Effort {
		diffs = append(diffs, noteeffortshape.GongMarshallField(stage, "Effort"))
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
	if noteperformanceshape.Note != noteperformanceshapeOther.Note {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "Note"))
	}
	if noteperformanceshape.Performance != noteperformanceshapeOther.Performance {
		diffs = append(diffs, noteperformanceshape.GongMarshallField(stage, "Performance"))
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
	if noteshape.Note != noteshapeOther.Note {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
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
	if ops := __gong__diffSliceOfPointers(stage, system, "Complexities", systemOther.Complexities, system.Complexities); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "Performances", systemOther.Performances, system.Performances); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "Efforts", systemOther.Efforts, system.Efforts); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "SubSystems", systemOther.SubSystems, system.SubSystems); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, system, "DiagramFlossEquations", systemOther.DiagramFlossEquations, system.DiagramFlossEquations); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "DiagramFlossEquationsWhoseNodeIsExpanded", systemOther.DiagramFlossEquationsWhoseNodeIsExpanded, system.DiagramFlossEquationsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if system.IsSubSystemNodeExpanded != systemOther.IsSubSystemNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
	}
	if system.IsComplexitysNodeExpanded != systemOther.IsComplexitysNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsComplexitysNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "ComplexitysWhoseNodeIsExpanded", systemOther.ComplexitysWhoseNodeIsExpanded, system.ComplexitysWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if system.IsPerformancesNodeExpanded != systemOther.IsPerformancesNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsPerformancesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "PerformancesWhoseNodeIsExpanded", systemOther.PerformancesWhoseNodeIsExpanded, system.PerformancesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if system.IsEffortsNodeExpanded != systemOther.IsEffortsNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsEffortsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "EffortsWhoseNodeIsExpanded", systemOther.EffortsWhoseNodeIsExpanded, system.EffortsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
