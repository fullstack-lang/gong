// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Complexity:
		if stage.OnAfterComplexityCreateCallback != nil {
			stage.OnAfterComplexityCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramFloss:
		if stage.OnAfterDiagramFlossCreateCallback != nil {
			stage.OnAfterDiagramFlossCreateCallback.OnAfterCreate(stage, target)
		}
	case *Effort:
		if stage.OnAfterEffortCreateCallback != nil {
			stage.OnAfterEffortCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Performance:
		if stage.OnAfterPerformanceCreateCallback != nil {
			stage.OnAfterPerformanceCreateCallback.OnAfterCreate(stage, target)
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
	case *Complexity:
		newTarget := any(new).(*Complexity)
		if stage.OnAfterComplexityUpdateCallback != nil {
			stage.OnAfterComplexityUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Complexity:
		if stage.OnAfterComplexityDeleteCallback != nil {
			staged := any(staged).(*Complexity)
			stage.OnAfterComplexityDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Complexity:
		if stage.OnAfterComplexityReadCallback != nil {
			stage.OnAfterComplexityReadCallback.OnAfterRead(stage, target)
		}
	case *DiagramFloss:
		if stage.OnAfterDiagramFlossReadCallback != nil {
			stage.OnAfterDiagramFlossReadCallback.OnAfterRead(stage, target)
		}
	case *Effort:
		if stage.OnAfterEffortReadCallback != nil {
			stage.OnAfterEffortReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Performance:
		if stage.OnAfterPerformanceReadCallback != nil {
			stage.OnAfterPerformanceReadCallback.OnAfterRead(stage, target)
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
	case *Complexity:
		stage.OnAfterComplexityUpdateCallback = any(callback).(OnAfterUpdateInterface[Complexity])
	case *DiagramFloss:
		stage.OnAfterDiagramFlossUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramFloss])
	case *Effort:
		stage.OnAfterEffortUpdateCallback = any(callback).(OnAfterUpdateInterface[Effort])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Performance:
		stage.OnAfterPerformanceUpdateCallback = any(callback).(OnAfterUpdateInterface[Performance])
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
	case *Complexity:
		stage.OnAfterComplexityCreateCallback = any(callback).(OnAfterCreateInterface[Complexity])
	case *DiagramFloss:
		stage.OnAfterDiagramFlossCreateCallback = any(callback).(OnAfterCreateInterface[DiagramFloss])
	case *Effort:
		stage.OnAfterEffortCreateCallback = any(callback).(OnAfterCreateInterface[Effort])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Performance:
		stage.OnAfterPerformanceCreateCallback = any(callback).(OnAfterCreateInterface[Performance])
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
	case *Complexity:
		stage.OnAfterComplexityDeleteCallback = any(callback).(OnAfterDeleteInterface[Complexity])
	case *DiagramFloss:
		stage.OnAfterDiagramFlossDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramFloss])
	case *Effort:
		stage.OnAfterEffortDeleteCallback = any(callback).(OnAfterDeleteInterface[Effort])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Performance:
		stage.OnAfterPerformanceDeleteCallback = any(callback).(OnAfterDeleteInterface[Performance])
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
	case *Complexity:
		stage.OnAfterComplexityReadCallback = any(callback).(OnAfterReadInterface[Complexity])
	case *DiagramFloss:
		stage.OnAfterDiagramFlossReadCallback = any(callback).(OnAfterReadInterface[DiagramFloss])
	case *Effort:
		stage.OnAfterEffortReadCallback = any(callback).(OnAfterReadInterface[Effort])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Performance:
		stage.OnAfterPerformanceReadCallback = any(callback).(OnAfterReadInterface[Performance])
	case *System:
		stage.OnAfterSystemReadCallback = any(callback).(OnAfterReadInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeReadCallback = any(callback).(OnAfterReadInterface[SystemShape])
	}
}
