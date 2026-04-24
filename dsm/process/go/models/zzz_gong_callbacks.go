// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *DiagramProcess:
		if stage.OnAfterDiagramProcessCreateCallback != nil {
			stage.OnAfterDiagramProcessCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Process:
		if stage.OnAfterProcessCreateCallback != nil {
			stage.OnAfterProcessCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProcessCompositionShape:
		if stage.OnAfterProcessCompositionShapeCreateCallback != nil {
			stage.OnAfterProcessCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProcessShape:
		if stage.OnAfterProcessShapeCreateCallback != nil {
			stage.OnAfterProcessShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *DiagramProcess:
		newTarget := any(new).(*DiagramProcess)
		if stage.OnAfterDiagramProcessUpdateCallback != nil {
			stage.OnAfterDiagramProcessUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Process:
		newTarget := any(new).(*Process)
		if stage.OnAfterProcessUpdateCallback != nil {
			stage.OnAfterProcessUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ProcessCompositionShape:
		newTarget := any(new).(*ProcessCompositionShape)
		if stage.OnAfterProcessCompositionShapeUpdateCallback != nil {
			stage.OnAfterProcessCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ProcessShape:
		newTarget := any(new).(*ProcessShape)
		if stage.OnAfterProcessShapeUpdateCallback != nil {
			stage.OnAfterProcessShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *DiagramProcess:
		if stage.OnAfterDiagramProcessDeleteCallback != nil {
			staged := any(staged).(*DiagramProcess)
			stage.OnAfterDiagramProcessDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Process:
		if stage.OnAfterProcessDeleteCallback != nil {
			staged := any(staged).(*Process)
			stage.OnAfterProcessDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ProcessCompositionShape:
		if stage.OnAfterProcessCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*ProcessCompositionShape)
			stage.OnAfterProcessCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ProcessShape:
		if stage.OnAfterProcessShapeDeleteCallback != nil {
			staged := any(staged).(*ProcessShape)
			stage.OnAfterProcessShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *DiagramProcess:
		if stage.OnAfterDiagramProcessReadCallback != nil {
			stage.OnAfterDiagramProcessReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Process:
		if stage.OnAfterProcessReadCallback != nil {
			stage.OnAfterProcessReadCallback.OnAfterRead(stage, target)
		}
	case *ProcessCompositionShape:
		if stage.OnAfterProcessCompositionShapeReadCallback != nil {
			stage.OnAfterProcessCompositionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ProcessShape:
		if stage.OnAfterProcessShapeReadCallback != nil {
			stage.OnAfterProcessShapeReadCallback.OnAfterRead(stage, target)
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
	case *DiagramProcess:
		stage.OnAfterDiagramProcessUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramProcess])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Process:
		stage.OnAfterProcessUpdateCallback = any(callback).(OnAfterUpdateInterface[Process])
	case *ProcessCompositionShape:
		stage.OnAfterProcessCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProcessCompositionShape])
	case *ProcessShape:
		stage.OnAfterProcessShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProcessShape])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DiagramProcess:
		stage.OnAfterDiagramProcessCreateCallback = any(callback).(OnAfterCreateInterface[DiagramProcess])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Process:
		stage.OnAfterProcessCreateCallback = any(callback).(OnAfterCreateInterface[Process])
	case *ProcessCompositionShape:
		stage.OnAfterProcessCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProcessCompositionShape])
	case *ProcessShape:
		stage.OnAfterProcessShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProcessShape])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DiagramProcess:
		stage.OnAfterDiagramProcessDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramProcess])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Process:
		stage.OnAfterProcessDeleteCallback = any(callback).(OnAfterDeleteInterface[Process])
	case *ProcessCompositionShape:
		stage.OnAfterProcessCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProcessCompositionShape])
	case *ProcessShape:
		stage.OnAfterProcessShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProcessShape])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DiagramProcess:
		stage.OnAfterDiagramProcessReadCallback = any(callback).(OnAfterReadInterface[DiagramProcess])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Process:
		stage.OnAfterProcessReadCallback = any(callback).(OnAfterReadInterface[Process])
	case *ProcessCompositionShape:
		stage.OnAfterProcessCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[ProcessCompositionShape])
	case *ProcessShape:
		stage.OnAfterProcessShapeReadCallback = any(callback).(OnAfterReadInterface[ProcessShape])
	}
}
