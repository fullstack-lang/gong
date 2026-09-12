// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
