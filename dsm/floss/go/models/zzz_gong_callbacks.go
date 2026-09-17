// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (compareanalysis *CompareAnalysis) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCompareAnalysisCreateCallback != nil {
		stage.OnAfterCompareAnalysisCreateCallback.OnAfterCreate(stage, compareanalysis)
	}
}

func (compareanalysis *CompareAnalysis) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompareAnalysisUpdateCallback != nil {
		var frontCompareAnalysis *CompareAnalysis
		if front != nil {
			frontCompareAnalysis, _ = front.(*CompareAnalysis)
		}
		stage.OnAfterCompareAnalysisUpdateCallback.OnAfterUpdate(stage, compareanalysis, frontCompareAnalysis)
	}
}

func (compareanalysis *CompareAnalysis) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompareAnalysisDeleteCallback != nil {
		var frontCompareAnalysis *CompareAnalysis
		if front != nil {
			frontCompareAnalysis, _ = front.(*CompareAnalysis)
		}
		stage.OnAfterCompareAnalysisDeleteCallback.OnAfterDelete(stage, compareanalysis, frontCompareAnalysis)
	}
}

func (complexity *Complexity) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterComplexityCreateCallback != nil {
		stage.OnAfterComplexityCreateCallback.OnAfterCreate(stage, complexity)
	}
}

func (complexity *Complexity) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterComplexityUpdateCallback != nil {
		var frontComplexity *Complexity
		if front != nil {
			frontComplexity, _ = front.(*Complexity)
		}
		stage.OnAfterComplexityUpdateCallback.OnAfterUpdate(stage, complexity, frontComplexity)
	}
}

func (complexity *Complexity) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterComplexityDeleteCallback != nil {
		var frontComplexity *Complexity
		if front != nil {
			frontComplexity, _ = front.(*Complexity)
		}
		stage.OnAfterComplexityDeleteCallback.OnAfterDelete(stage, complexity, frontComplexity)
	}
}

func (diagramflossequation *DiagramFlossEquation) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramFlossEquationCreateCallback != nil {
		stage.OnAfterDiagramFlossEquationCreateCallback.OnAfterCreate(stage, diagramflossequation)
	}
}

func (diagramflossequation *DiagramFlossEquation) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramFlossEquationUpdateCallback != nil {
		var frontDiagramFlossEquation *DiagramFlossEquation
		if front != nil {
			frontDiagramFlossEquation, _ = front.(*DiagramFlossEquation)
		}
		stage.OnAfterDiagramFlossEquationUpdateCallback.OnAfterUpdate(stage, diagramflossequation, frontDiagramFlossEquation)
	}
}

func (diagramflossequation *DiagramFlossEquation) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramFlossEquationDeleteCallback != nil {
		var frontDiagramFlossEquation *DiagramFlossEquation
		if front != nil {
			frontDiagramFlossEquation, _ = front.(*DiagramFlossEquation)
		}
		stage.OnAfterDiagramFlossEquationDeleteCallback.OnAfterDelete(stage, diagramflossequation, frontDiagramFlossEquation)
	}
}

func (effort *Effort) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEffortCreateCallback != nil {
		stage.OnAfterEffortCreateCallback.OnAfterCreate(stage, effort)
	}
}

func (effort *Effort) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEffortUpdateCallback != nil {
		var frontEffort *Effort
		if front != nil {
			frontEffort, _ = front.(*Effort)
		}
		stage.OnAfterEffortUpdateCallback.OnAfterUpdate(stage, effort, frontEffort)
	}
}

func (effort *Effort) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEffortDeleteCallback != nil {
		var frontEffort *Effort
		if front != nil {
			frontEffort, _ = front.(*Effort)
		}
		stage.OnAfterEffortDeleteCallback.OnAfterDelete(stage, effort, frontEffort)
	}
}

func (library *Library) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLibraryCreateCallback != nil {
		stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, library)
	}
}

func (library *Library) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryUpdateCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, library, frontLibrary)
	}
}

func (library *Library) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryDeleteCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, library, frontLibrary)
	}
}

func (note *Note) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteCreateCallback != nil {
		stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, note)
	}
}

func (note *Note) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteUpdateCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, note, frontNote)
	}
}

func (note *Note) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteDeleteCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, note, frontNote)
	}
}

func (notecomplexityshape *NoteComplexityShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteComplexityShapeCreateCallback != nil {
		stage.OnAfterNoteComplexityShapeCreateCallback.OnAfterCreate(stage, notecomplexityshape)
	}
}

func (notecomplexityshape *NoteComplexityShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteComplexityShapeUpdateCallback != nil {
		var frontNoteComplexityShape *NoteComplexityShape
		if front != nil {
			frontNoteComplexityShape, _ = front.(*NoteComplexityShape)
		}
		stage.OnAfterNoteComplexityShapeUpdateCallback.OnAfterUpdate(stage, notecomplexityshape, frontNoteComplexityShape)
	}
}

func (notecomplexityshape *NoteComplexityShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteComplexityShapeDeleteCallback != nil {
		var frontNoteComplexityShape *NoteComplexityShape
		if front != nil {
			frontNoteComplexityShape, _ = front.(*NoteComplexityShape)
		}
		stage.OnAfterNoteComplexityShapeDeleteCallback.OnAfterDelete(stage, notecomplexityshape, frontNoteComplexityShape)
	}
}

func (noteeffortshape *NoteEffortShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteEffortShapeCreateCallback != nil {
		stage.OnAfterNoteEffortShapeCreateCallback.OnAfterCreate(stage, noteeffortshape)
	}
}

func (noteeffortshape *NoteEffortShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteEffortShapeUpdateCallback != nil {
		var frontNoteEffortShape *NoteEffortShape
		if front != nil {
			frontNoteEffortShape, _ = front.(*NoteEffortShape)
		}
		stage.OnAfterNoteEffortShapeUpdateCallback.OnAfterUpdate(stage, noteeffortshape, frontNoteEffortShape)
	}
}

func (noteeffortshape *NoteEffortShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteEffortShapeDeleteCallback != nil {
		var frontNoteEffortShape *NoteEffortShape
		if front != nil {
			frontNoteEffortShape, _ = front.(*NoteEffortShape)
		}
		stage.OnAfterNoteEffortShapeDeleteCallback.OnAfterDelete(stage, noteeffortshape, frontNoteEffortShape)
	}
}

func (noteperformanceshape *NotePerformanceShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNotePerformanceShapeCreateCallback != nil {
		stage.OnAfterNotePerformanceShapeCreateCallback.OnAfterCreate(stage, noteperformanceshape)
	}
}

func (noteperformanceshape *NotePerformanceShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePerformanceShapeUpdateCallback != nil {
		var frontNotePerformanceShape *NotePerformanceShape
		if front != nil {
			frontNotePerformanceShape, _ = front.(*NotePerformanceShape)
		}
		stage.OnAfterNotePerformanceShapeUpdateCallback.OnAfterUpdate(stage, noteperformanceshape, frontNotePerformanceShape)
	}
}

func (noteperformanceshape *NotePerformanceShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePerformanceShapeDeleteCallback != nil {
		var frontNotePerformanceShape *NotePerformanceShape
		if front != nil {
			frontNotePerformanceShape, _ = front.(*NotePerformanceShape)
		}
		stage.OnAfterNotePerformanceShapeDeleteCallback.OnAfterDelete(stage, noteperformanceshape, frontNotePerformanceShape)
	}
}

func (noteshape *NoteShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteShapeCreateCallback != nil {
		stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, noteshape)
	}
}

func (noteshape *NoteShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteShapeUpdateCallback != nil {
		var frontNoteShape *NoteShape
		if front != nil {
			frontNoteShape, _ = front.(*NoteShape)
		}
		stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, noteshape, frontNoteShape)
	}
}

func (noteshape *NoteShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteShapeDeleteCallback != nil {
		var frontNoteShape *NoteShape
		if front != nil {
			frontNoteShape, _ = front.(*NoteShape)
		}
		stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, noteshape, frontNoteShape)
	}
}

func (performance *Performance) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPerformanceCreateCallback != nil {
		stage.OnAfterPerformanceCreateCallback.OnAfterCreate(stage, performance)
	}
}

func (performance *Performance) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerformanceUpdateCallback != nil {
		var frontPerformance *Performance
		if front != nil {
			frontPerformance, _ = front.(*Performance)
		}
		stage.OnAfterPerformanceUpdateCallback.OnAfterUpdate(stage, performance, frontPerformance)
	}
}

func (performance *Performance) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerformanceDeleteCallback != nil {
		var frontPerformance *Performance
		if front != nil {
			frontPerformance, _ = front.(*Performance)
		}
		stage.OnAfterPerformanceDeleteCallback.OnAfterDelete(stage, performance, frontPerformance)
	}
}

func (system *System) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystemCreateCallback != nil {
		stage.OnAfterSystemCreateCallback.OnAfterCreate(stage, system)
	}
}

func (system *System) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemUpdateCallback != nil {
		var frontSystem *System
		if front != nil {
			frontSystem, _ = front.(*System)
		}
		stage.OnAfterSystemUpdateCallback.OnAfterUpdate(stage, system, frontSystem)
	}
}

func (system *System) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemDeleteCallback != nil {
		var frontSystem *System
		if front != nil {
			frontSystem, _ = front.(*System)
		}
		stage.OnAfterSystemDeleteCallback.OnAfterDelete(stage, system, frontSystem)
	}
}

