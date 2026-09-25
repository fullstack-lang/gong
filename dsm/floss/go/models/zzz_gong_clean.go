// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by CompareAnalysis
func (compareanalysis *CompareAnalysis) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&compareanalysis.DiagramFlossEquations) || modified
	modified = stage.CleanSlice(&compareanalysis.DiagramFlossEquationsWhoseNodeIsExpanded) || modified
	// insertion point per field
	modified = stage.CleanPointer(&compareanalysis.FromSystem) || modified
	modified = stage.CleanPointer(&compareanalysis.ToSystem) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramFlossEquation
func (diagramflossequation *DiagramFlossEquation) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&diagramflossequation.Note_Shapes) || modified
	modified = stage.CleanSlice(&diagramflossequation.NoteComplexityShapes) || modified
	modified = stage.CleanSlice(&diagramflossequation.NotePerformanceShapes) || modified
	modified = stage.CleanSlice(&diagramflossequation.NoteEffortShapes) || modified
	modified = stage.CleanSlice(&diagramflossequation.NotesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramflossequation.ComplexitysWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramflossequation.PerformancesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramflossequation.EffortsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	modified = stage.CleanSlice(&library.RootSystems) || modified
	modified = stage.CleanSlice(&library.RootComplexitys) || modified
	modified = stage.CleanSlice(&library.RootPerformances) || modified
	modified = stage.CleanSlice(&library.RootEfforts) || modified
	modified = stage.CleanSlice(&library.RootCompareAnalysis) || modified
	modified = stage.CleanSlice(&library.RootNotes) || modified
	modified = stage.CleanSlice(&library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.SystemsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.ComplexitysWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.PerformancesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.EffortsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.CompareAnalysisWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.NotesWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&note.Complexities) || modified
	modified = stage.CleanSlice(&note.Performances) || modified
	modified = stage.CleanSlice(&note.Efforts) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteComplexityShape
func (notecomplexityshape *NoteComplexityShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&notecomplexityshape.Note) || modified
	modified = stage.CleanPointer(&notecomplexityshape.Complexity) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteEffortShape
func (noteeffortshape *NoteEffortShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteeffortshape.Note) || modified
	modified = stage.CleanPointer(&noteeffortshape.Effort) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NotePerformanceShape
func (noteperformanceshape *NotePerformanceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteperformanceshape.Note) || modified
	modified = stage.CleanPointer(&noteperformanceshape.Performance) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by System
func (system *System) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&system.Complexities) || modified
	modified = stage.CleanSlice(&system.Performances) || modified
	modified = stage.CleanSlice(&system.Efforts) || modified
	modified = stage.CleanSlice(&system.SubSystems) || modified
	modified = stage.CleanSlice(&system.DiagramFlossEquations) || modified
	modified = stage.CleanSlice(&system.DiagramFlossEquationsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&system.ComplexitysWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&system.PerformancesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&system.EffortsWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
