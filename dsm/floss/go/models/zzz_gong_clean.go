// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by CompareAnalysis
func (compareanalysis *CompareAnalysis) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &compareanalysis.DiagramFlossEquations) || modified
	modified = GongCleanSlice(stage, &compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &compareanalysis.FromSystem) || modified
	modified = GongCleanPointer(stage, &compareanalysis.ToSystem) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Complexity
func (complexity *Complexity) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramFlossEquation
func (diagramflossequation *DiagramFlossEquation) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagramflossequation.Note_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramflossequation.NoteComplexityShapes) || modified
	modified = GongCleanSlice(stage, &diagramflossequation.NotePerformanceShapes) || modified
	modified = GongCleanSlice(stage, &diagramflossequation.NoteEffortShapes) || modified
	modified = GongCleanSlice(stage, &diagramflossequation.NotesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramflossequation.ComplexitysWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramflossequation.PerformancesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramflossequation.EffortsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Effort
func (effort *Effort) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	modified = GongCleanSlice(stage, &library.RootSystems) || modified
	modified = GongCleanSlice(stage, &library.RootComplexitys) || modified
	modified = GongCleanSlice(stage, &library.RootPerformances) || modified
	modified = GongCleanSlice(stage, &library.RootEfforts) || modified
	modified = GongCleanSlice(stage, &library.RootCompareAnalysis) || modified
	modified = GongCleanSlice(stage, &library.RootNotes) || modified
	modified = GongCleanSlice(stage, &library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.SystemsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.ComplexitysWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.PerformancesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.EffortsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.CompareAnalysisWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.NotesWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &note.Complexities) || modified
	modified = GongCleanSlice(stage, &note.Performances) || modified
	modified = GongCleanSlice(stage, &note.Efforts) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteComplexityShape
func (notecomplexityshape *NoteComplexityShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &notecomplexityshape.Note) || modified
	modified = GongCleanPointer(stage, &notecomplexityshape.Complexity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteEffortShape
func (noteeffortshape *NoteEffortShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteeffortshape.Note) || modified
	modified = GongCleanPointer(stage, &noteeffortshape.Effort) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NotePerformanceShape
func (noteperformanceshape *NotePerformanceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteperformanceshape.Note) || modified
	modified = GongCleanPointer(stage, &noteperformanceshape.Performance) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Performance
func (performance *Performance) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by System
func (system *System) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &system.Complexities) || modified
	modified = GongCleanSlice(stage, &system.Performances) || modified
	modified = GongCleanSlice(stage, &system.Efforts) || modified
	modified = GongCleanSlice(stage, &system.DiagramFlossEquations) || modified
	modified = GongCleanSlice(stage, &system.DiagramFlossEquationsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.SubSystemes) || modified
	modified = GongCleanSlice(stage, &system.ComplexitysWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.PerformancesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.EffortsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
