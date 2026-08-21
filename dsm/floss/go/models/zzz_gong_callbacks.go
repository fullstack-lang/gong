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
	case *ComplexityShape:
		if stage.OnAfterComplexityShapeCreateCallback != nil {
			stage.OnAfterComplexityShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramFloss:
		if stage.OnAfterDiagramFlossCreateCallback != nil {
			stage.OnAfterDiagramFlossCreateCallback.OnAfterCreate(stage, target)
		}
	case *Effort:
		if stage.OnAfterEffortCreateCallback != nil {
			stage.OnAfterEffortCreateCallback.OnAfterCreate(stage, target)
		}
	case *EffortShape:
		if stage.OnAfterEffortShapeCreateCallback != nil {
			stage.OnAfterEffortShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Performance:
		if stage.OnAfterPerformanceCreateCallback != nil {
			stage.OnAfterPerformanceCreateCallback.OnAfterCreate(stage, target)
		}
	case *PerformanceShape:
		if stage.OnAfterPerformanceShapeCreateCallback != nil {
			stage.OnAfterPerformanceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *System:
		if stage.OnAfterSystemCreateCallback != nil {
			stage.OnAfterSystemCreateCallback.OnAfterCreate(stage, target)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeCreateCallback != nil {
			stage.OnAfterSystemShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *ComplexityShape:
		newTarget := any(new).(*ComplexityShape)
		if stage.OnAfterComplexityShapeUpdateCallback != nil {
			stage.OnAfterComplexityShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramFloss:
		newTarget := any(new).(*DiagramFloss)
		if stage.OnAfterDiagramFlossUpdateCallback != nil {
			stage.OnAfterDiagramFlossUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Effort:
		newTarget := any(new).(*Effort)
		if stage.OnAfterEffortUpdateCallback != nil {
			stage.OnAfterEffortUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EffortShape:
		newTarget := any(new).(*EffortShape)
		if stage.OnAfterEffortShapeUpdateCallback != nil {
			stage.OnAfterEffortShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Performance:
		newTarget := any(new).(*Performance)
		if stage.OnAfterPerformanceUpdateCallback != nil {
			stage.OnAfterPerformanceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PerformanceShape:
		newTarget := any(new).(*PerformanceShape)
		if stage.OnAfterPerformanceShapeUpdateCallback != nil {
			stage.OnAfterPerformanceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *System:
		newTarget := any(new).(*System)
		if stage.OnAfterSystemUpdateCallback != nil {
			stage.OnAfterSystemUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SystemShape:
		newTarget := any(new).(*SystemShape)
		if stage.OnAfterSystemShapeUpdateCallback != nil {
			stage.OnAfterSystemShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *ComplexityShape:
		if stage.OnAfterComplexityShapeDeleteCallback != nil {
			staged := any(staged).(*ComplexityShape)
			stage.OnAfterComplexityShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramFloss:
		if stage.OnAfterDiagramFlossDeleteCallback != nil {
			staged := any(staged).(*DiagramFloss)
			stage.OnAfterDiagramFlossDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Effort:
		if stage.OnAfterEffortDeleteCallback != nil {
			staged := any(staged).(*Effort)
			stage.OnAfterEffortDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EffortShape:
		if stage.OnAfterEffortShapeDeleteCallback != nil {
			staged := any(staged).(*EffortShape)
			stage.OnAfterEffortShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Performance:
		if stage.OnAfterPerformanceDeleteCallback != nil {
			staged := any(staged).(*Performance)
			stage.OnAfterPerformanceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PerformanceShape:
		if stage.OnAfterPerformanceShapeDeleteCallback != nil {
			staged := any(staged).(*PerformanceShape)
			stage.OnAfterPerformanceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *System:
		if stage.OnAfterSystemDeleteCallback != nil {
			staged := any(staged).(*System)
			stage.OnAfterSystemDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeDeleteCallback != nil {
			staged := any(staged).(*SystemShape)
			stage.OnAfterSystemShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *ComplexityShape:
		if stage.OnAfterComplexityShapeReadCallback != nil {
			stage.OnAfterComplexityShapeReadCallback.OnAfterRead(stage, target)
		}
	case *DiagramFloss:
		if stage.OnAfterDiagramFlossReadCallback != nil {
			stage.OnAfterDiagramFlossReadCallback.OnAfterRead(stage, target)
		}
	case *Effort:
		if stage.OnAfterEffortReadCallback != nil {
			stage.OnAfterEffortReadCallback.OnAfterRead(stage, target)
		}
	case *EffortShape:
		if stage.OnAfterEffortShapeReadCallback != nil {
			stage.OnAfterEffortShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Performance:
		if stage.OnAfterPerformanceReadCallback != nil {
			stage.OnAfterPerformanceReadCallback.OnAfterRead(stage, target)
		}
	case *PerformanceShape:
		if stage.OnAfterPerformanceShapeReadCallback != nil {
			stage.OnAfterPerformanceShapeReadCallback.OnAfterRead(stage, target)
		}
	case *System:
		if stage.OnAfterSystemReadCallback != nil {
			stage.OnAfterSystemReadCallback.OnAfterRead(stage, target)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeReadCallback != nil {
			stage.OnAfterSystemShapeReadCallback.OnAfterRead(stage, target)
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
	case *ComplexityShape:
		stage.OnAfterComplexityShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ComplexityShape])
	case *DiagramFloss:
		stage.OnAfterDiagramFlossUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramFloss])
	case *Effort:
		stage.OnAfterEffortUpdateCallback = any(callback).(OnAfterUpdateInterface[Effort])
	case *EffortShape:
		stage.OnAfterEffortShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[EffortShape])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Performance:
		stage.OnAfterPerformanceUpdateCallback = any(callback).(OnAfterUpdateInterface[Performance])
	case *PerformanceShape:
		stage.OnAfterPerformanceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PerformanceShape])
	case *System:
		stage.OnAfterSystemUpdateCallback = any(callback).(OnAfterUpdateInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[SystemShape])
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
	case *ComplexityShape:
		stage.OnAfterComplexityShapeCreateCallback = any(callback).(OnAfterCreateInterface[ComplexityShape])
	case *DiagramFloss:
		stage.OnAfterDiagramFlossCreateCallback = any(callback).(OnAfterCreateInterface[DiagramFloss])
	case *Effort:
		stage.OnAfterEffortCreateCallback = any(callback).(OnAfterCreateInterface[Effort])
	case *EffortShape:
		stage.OnAfterEffortShapeCreateCallback = any(callback).(OnAfterCreateInterface[EffortShape])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Performance:
		stage.OnAfterPerformanceCreateCallback = any(callback).(OnAfterCreateInterface[Performance])
	case *PerformanceShape:
		stage.OnAfterPerformanceShapeCreateCallback = any(callback).(OnAfterCreateInterface[PerformanceShape])
	case *System:
		stage.OnAfterSystemCreateCallback = any(callback).(OnAfterCreateInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeCreateCallback = any(callback).(OnAfterCreateInterface[SystemShape])
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
	case *ComplexityShape:
		stage.OnAfterComplexityShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ComplexityShape])
	case *DiagramFloss:
		stage.OnAfterDiagramFlossDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramFloss])
	case *Effort:
		stage.OnAfterEffortDeleteCallback = any(callback).(OnAfterDeleteInterface[Effort])
	case *EffortShape:
		stage.OnAfterEffortShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[EffortShape])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Performance:
		stage.OnAfterPerformanceDeleteCallback = any(callback).(OnAfterDeleteInterface[Performance])
	case *PerformanceShape:
		stage.OnAfterPerformanceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PerformanceShape])
	case *System:
		stage.OnAfterSystemDeleteCallback = any(callback).(OnAfterDeleteInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[SystemShape])
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
	case *ComplexityShape:
		stage.OnAfterComplexityShapeReadCallback = any(callback).(OnAfterReadInterface[ComplexityShape])
	case *DiagramFloss:
		stage.OnAfterDiagramFlossReadCallback = any(callback).(OnAfterReadInterface[DiagramFloss])
	case *Effort:
		stage.OnAfterEffortReadCallback = any(callback).(OnAfterReadInterface[Effort])
	case *EffortShape:
		stage.OnAfterEffortShapeReadCallback = any(callback).(OnAfterReadInterface[EffortShape])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Performance:
		stage.OnAfterPerformanceReadCallback = any(callback).(OnAfterReadInterface[Performance])
	case *PerformanceShape:
		stage.OnAfterPerformanceShapeReadCallback = any(callback).(OnAfterReadInterface[PerformanceShape])
	case *System:
		stage.OnAfterSystemReadCallback = any(callback).(OnAfterReadInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeReadCallback = any(callback).(OnAfterReadInterface[SystemShape])
	}
}
