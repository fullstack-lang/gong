// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CompareAnalysis:
		if stage.OnAfterCompareAnalysisCreateCallback != nil {
			stage.OnAfterCompareAnalysisCreateCallback.OnAfterCreate(stage, target)
		}
	case *Complexity:
		if stage.OnAfterComplexityCreateCallback != nil {
			stage.OnAfterComplexityCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramFlossEquation:
		if stage.OnAfterDiagramFlossEquationCreateCallback != nil {
			stage.OnAfterDiagramFlossEquationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Effort:
		if stage.OnAfterEffortCreateCallback != nil {
			stage.OnAfterEffortCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteComplexityShape:
		if stage.OnAfterNoteComplexityShapeCreateCallback != nil {
			stage.OnAfterNoteComplexityShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteEffortShape:
		if stage.OnAfterNoteEffortShapeCreateCallback != nil {
			stage.OnAfterNoteEffortShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NotePerformanceShape:
		if stage.OnAfterNotePerformanceShapeCreateCallback != nil {
			stage.OnAfterNotePerformanceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeCreateCallback != nil {
			stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Performance:
		if stage.OnAfterPerformanceCreateCallback != nil {
			stage.OnAfterPerformanceCreateCallback.OnAfterCreate(stage, target)
		}
	case *System:
		if stage.OnAfterSystemCreateCallback != nil {
			stage.OnAfterSystemCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *CompareAnalysis:
		newTarget := any(new).(*CompareAnalysis)
		if stage.OnAfterCompareAnalysisUpdateCallback != nil {
			stage.OnAfterCompareAnalysisUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Complexity:
		newTarget := any(new).(*Complexity)
		if stage.OnAfterComplexityUpdateCallback != nil {
			stage.OnAfterComplexityUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramFlossEquation:
		newTarget := any(new).(*DiagramFlossEquation)
		if stage.OnAfterDiagramFlossEquationUpdateCallback != nil {
			stage.OnAfterDiagramFlossEquationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Effort:
		newTarget := any(new).(*Effort)
		if stage.OnAfterEffortUpdateCallback != nil {
			stage.OnAfterEffortUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note:
		newTarget := any(new).(*Note)
		if stage.OnAfterNoteUpdateCallback != nil {
			stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteComplexityShape:
		newTarget := any(new).(*NoteComplexityShape)
		if stage.OnAfterNoteComplexityShapeUpdateCallback != nil {
			stage.OnAfterNoteComplexityShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteEffortShape:
		newTarget := any(new).(*NoteEffortShape)
		if stage.OnAfterNoteEffortShapeUpdateCallback != nil {
			stage.OnAfterNoteEffortShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NotePerformanceShape:
		newTarget := any(new).(*NotePerformanceShape)
		if stage.OnAfterNotePerformanceShapeUpdateCallback != nil {
			stage.OnAfterNotePerformanceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteShape:
		newTarget := any(new).(*NoteShape)
		if stage.OnAfterNoteShapeUpdateCallback != nil {
			stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Performance:
		newTarget := any(new).(*Performance)
		if stage.OnAfterPerformanceUpdateCallback != nil {
			stage.OnAfterPerformanceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *System:
		newTarget := any(new).(*System)
		if stage.OnAfterSystemUpdateCallback != nil {
			stage.OnAfterSystemUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *CompareAnalysis:
		if stage.OnAfterCompareAnalysisDeleteCallback != nil {
			staged := any(staged).(*CompareAnalysis)
			stage.OnAfterCompareAnalysisDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Complexity:
		if stage.OnAfterComplexityDeleteCallback != nil {
			staged := any(staged).(*Complexity)
			stage.OnAfterComplexityDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramFlossEquation:
		if stage.OnAfterDiagramFlossEquationDeleteCallback != nil {
			staged := any(staged).(*DiagramFlossEquation)
			stage.OnAfterDiagramFlossEquationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Effort:
		if stage.OnAfterEffortDeleteCallback != nil {
			staged := any(staged).(*Effort)
			stage.OnAfterEffortDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note:
		if stage.OnAfterNoteDeleteCallback != nil {
			staged := any(staged).(*Note)
			stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteComplexityShape:
		if stage.OnAfterNoteComplexityShapeDeleteCallback != nil {
			staged := any(staged).(*NoteComplexityShape)
			stage.OnAfterNoteComplexityShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteEffortShape:
		if stage.OnAfterNoteEffortShapeDeleteCallback != nil {
			staged := any(staged).(*NoteEffortShape)
			stage.OnAfterNoteEffortShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NotePerformanceShape:
		if stage.OnAfterNotePerformanceShapeDeleteCallback != nil {
			staged := any(staged).(*NotePerformanceShape)
			stage.OnAfterNotePerformanceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeDeleteCallback != nil {
			staged := any(staged).(*NoteShape)
			stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Performance:
		if stage.OnAfterPerformanceDeleteCallback != nil {
			staged := any(staged).(*Performance)
			stage.OnAfterPerformanceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *System:
		if stage.OnAfterSystemDeleteCallback != nil {
			staged := any(staged).(*System)
			stage.OnAfterSystemDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CompareAnalysis:
		if stage.OnAfterCompareAnalysisReadCallback != nil {
			stage.OnAfterCompareAnalysisReadCallback.OnAfterRead(stage, target)
		}
	case *Complexity:
		if stage.OnAfterComplexityReadCallback != nil {
			stage.OnAfterComplexityReadCallback.OnAfterRead(stage, target)
		}
	case *DiagramFlossEquation:
		if stage.OnAfterDiagramFlossEquationReadCallback != nil {
			stage.OnAfterDiagramFlossEquationReadCallback.OnAfterRead(stage, target)
		}
	case *Effort:
		if stage.OnAfterEffortReadCallback != nil {
			stage.OnAfterEffortReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteReadCallback != nil {
			stage.OnAfterNoteReadCallback.OnAfterRead(stage, target)
		}
	case *NoteComplexityShape:
		if stage.OnAfterNoteComplexityShapeReadCallback != nil {
			stage.OnAfterNoteComplexityShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteEffortShape:
		if stage.OnAfterNoteEffortShapeReadCallback != nil {
			stage.OnAfterNoteEffortShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NotePerformanceShape:
		if stage.OnAfterNotePerformanceShapeReadCallback != nil {
			stage.OnAfterNotePerformanceShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeReadCallback != nil {
			stage.OnAfterNoteShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Performance:
		if stage.OnAfterPerformanceReadCallback != nil {
			stage.OnAfterPerformanceReadCallback.OnAfterRead(stage, target)
		}
	case *System:
		if stage.OnAfterSystemReadCallback != nil {
			stage.OnAfterSystemReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *CompareAnalysis:
		stage.OnAfterCompareAnalysisUpdateCallback = any(callback).(OnAfterUpdateInterface[CompareAnalysis])
	case *Complexity:
		stage.OnAfterComplexityUpdateCallback = any(callback).(OnAfterUpdateInterface[Complexity])
	case *DiagramFlossEquation:
		stage.OnAfterDiagramFlossEquationUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramFlossEquation])
	case *Effort:
		stage.OnAfterEffortUpdateCallback = any(callback).(OnAfterUpdateInterface[Effort])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Note:
		stage.OnAfterNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[Note])
	case *NoteComplexityShape:
		stage.OnAfterNoteComplexityShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteComplexityShape])
	case *NoteEffortShape:
		stage.OnAfterNoteEffortShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteEffortShape])
	case *NotePerformanceShape:
		stage.OnAfterNotePerformanceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NotePerformanceShape])
	case *NoteShape:
		stage.OnAfterNoteShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteShape])
	case *Performance:
		stage.OnAfterPerformanceUpdateCallback = any(callback).(OnAfterUpdateInterface[Performance])
	case *System:
		stage.OnAfterSystemUpdateCallback = any(callback).(OnAfterUpdateInterface[System])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *CompareAnalysis:
		stage.OnAfterCompareAnalysisCreateCallback = any(callback).(OnAfterCreateInterface[CompareAnalysis])
	case *Complexity:
		stage.OnAfterComplexityCreateCallback = any(callback).(OnAfterCreateInterface[Complexity])
	case *DiagramFlossEquation:
		stage.OnAfterDiagramFlossEquationCreateCallback = any(callback).(OnAfterCreateInterface[DiagramFlossEquation])
	case *Effort:
		stage.OnAfterEffortCreateCallback = any(callback).(OnAfterCreateInterface[Effort])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Note:
		stage.OnAfterNoteCreateCallback = any(callback).(OnAfterCreateInterface[Note])
	case *NoteComplexityShape:
		stage.OnAfterNoteComplexityShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteComplexityShape])
	case *NoteEffortShape:
		stage.OnAfterNoteEffortShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteEffortShape])
	case *NotePerformanceShape:
		stage.OnAfterNotePerformanceShapeCreateCallback = any(callback).(OnAfterCreateInterface[NotePerformanceShape])
	case *NoteShape:
		stage.OnAfterNoteShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteShape])
	case *Performance:
		stage.OnAfterPerformanceCreateCallback = any(callback).(OnAfterCreateInterface[Performance])
	case *System:
		stage.OnAfterSystemCreateCallback = any(callback).(OnAfterCreateInterface[System])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *CompareAnalysis:
		stage.OnAfterCompareAnalysisDeleteCallback = any(callback).(OnAfterDeleteInterface[CompareAnalysis])
	case *Complexity:
		stage.OnAfterComplexityDeleteCallback = any(callback).(OnAfterDeleteInterface[Complexity])
	case *DiagramFlossEquation:
		stage.OnAfterDiagramFlossEquationDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramFlossEquation])
	case *Effort:
		stage.OnAfterEffortDeleteCallback = any(callback).(OnAfterDeleteInterface[Effort])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Note:
		stage.OnAfterNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[Note])
	case *NoteComplexityShape:
		stage.OnAfterNoteComplexityShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteComplexityShape])
	case *NoteEffortShape:
		stage.OnAfterNoteEffortShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteEffortShape])
	case *NotePerformanceShape:
		stage.OnAfterNotePerformanceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NotePerformanceShape])
	case *NoteShape:
		stage.OnAfterNoteShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteShape])
	case *Performance:
		stage.OnAfterPerformanceDeleteCallback = any(callback).(OnAfterDeleteInterface[Performance])
	case *System:
		stage.OnAfterSystemDeleteCallback = any(callback).(OnAfterDeleteInterface[System])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *CompareAnalysis:
		stage.OnAfterCompareAnalysisReadCallback = any(callback).(OnAfterReadInterface[CompareAnalysis])
	case *Complexity:
		stage.OnAfterComplexityReadCallback = any(callback).(OnAfterReadInterface[Complexity])
	case *DiagramFlossEquation:
		stage.OnAfterDiagramFlossEquationReadCallback = any(callback).(OnAfterReadInterface[DiagramFlossEquation])
	case *Effort:
		stage.OnAfterEffortReadCallback = any(callback).(OnAfterReadInterface[Effort])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Note:
		stage.OnAfterNoteReadCallback = any(callback).(OnAfterReadInterface[Note])
	case *NoteComplexityShape:
		stage.OnAfterNoteComplexityShapeReadCallback = any(callback).(OnAfterReadInterface[NoteComplexityShape])
	case *NoteEffortShape:
		stage.OnAfterNoteEffortShapeReadCallback = any(callback).(OnAfterReadInterface[NoteEffortShape])
	case *NotePerformanceShape:
		stage.OnAfterNotePerformanceShapeReadCallback = any(callback).(OnAfterReadInterface[NotePerformanceShape])
	case *NoteShape:
		stage.OnAfterNoteShapeReadCallback = any(callback).(OnAfterReadInterface[NoteShape])
	case *Performance:
		stage.OnAfterPerformanceReadCallback = any(callback).(OnAfterReadInterface[Performance])
	case *System:
		stage.OnAfterSystemReadCallback = any(callback).(OnAfterReadInterface[System])
	}
}
