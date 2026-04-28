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
	case *Participant:
		if stage.OnAfterParticipantCreateCallback != nil {
			stage.OnAfterParticipantCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParticipantShape:
		if stage.OnAfterParticipantShapeCreateCallback != nil {
			stage.OnAfterParticipantShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Process:
		if stage.OnAfterProcessCreateCallback != nil {
			stage.OnAfterProcessCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProcessShape:
		if stage.OnAfterProcessShapeCreateCallback != nil {
			stage.OnAfterProcessShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskCreateCallback != nil {
			stage.OnAfterTaskCreateCallback.OnAfterCreate(stage, target)
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
	case *Participant:
		newTarget := any(new).(*Participant)
		if stage.OnAfterParticipantUpdateCallback != nil {
			stage.OnAfterParticipantUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParticipantShape:
		newTarget := any(new).(*ParticipantShape)
		if stage.OnAfterParticipantShapeUpdateCallback != nil {
			stage.OnAfterParticipantShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Process:
		newTarget := any(new).(*Process)
		if stage.OnAfterProcessUpdateCallback != nil {
			stage.OnAfterProcessUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ProcessShape:
		newTarget := any(new).(*ProcessShape)
		if stage.OnAfterProcessShapeUpdateCallback != nil {
			stage.OnAfterProcessShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Task:
		newTarget := any(new).(*Task)
		if stage.OnAfterTaskUpdateCallback != nil {
			stage.OnAfterTaskUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Participant:
		if stage.OnAfterParticipantDeleteCallback != nil {
			staged := any(staged).(*Participant)
			stage.OnAfterParticipantDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParticipantShape:
		if stage.OnAfterParticipantShapeDeleteCallback != nil {
			staged := any(staged).(*ParticipantShape)
			stage.OnAfterParticipantShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Process:
		if stage.OnAfterProcessDeleteCallback != nil {
			staged := any(staged).(*Process)
			stage.OnAfterProcessDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ProcessShape:
		if stage.OnAfterProcessShapeDeleteCallback != nil {
			staged := any(staged).(*ProcessShape)
			stage.OnAfterProcessShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Task:
		if stage.OnAfterTaskDeleteCallback != nil {
			staged := any(staged).(*Task)
			stage.OnAfterTaskDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Participant:
		if stage.OnAfterParticipantReadCallback != nil {
			stage.OnAfterParticipantReadCallback.OnAfterRead(stage, target)
		}
	case *ParticipantShape:
		if stage.OnAfterParticipantShapeReadCallback != nil {
			stage.OnAfterParticipantShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Process:
		if stage.OnAfterProcessReadCallback != nil {
			stage.OnAfterProcessReadCallback.OnAfterRead(stage, target)
		}
	case *ProcessShape:
		if stage.OnAfterProcessShapeReadCallback != nil {
			stage.OnAfterProcessShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskReadCallback != nil {
			stage.OnAfterTaskReadCallback.OnAfterRead(stage, target)
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
	case *Participant:
		stage.OnAfterParticipantUpdateCallback = any(callback).(OnAfterUpdateInterface[Participant])
	case *ParticipantShape:
		stage.OnAfterParticipantShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ParticipantShape])
	case *Process:
		stage.OnAfterProcessUpdateCallback = any(callback).(OnAfterUpdateInterface[Process])
	case *ProcessShape:
		stage.OnAfterProcessShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProcessShape])
	case *Task:
		stage.OnAfterTaskUpdateCallback = any(callback).(OnAfterUpdateInterface[Task])
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
	case *Participant:
		stage.OnAfterParticipantCreateCallback = any(callback).(OnAfterCreateInterface[Participant])
	case *ParticipantShape:
		stage.OnAfterParticipantShapeCreateCallback = any(callback).(OnAfterCreateInterface[ParticipantShape])
	case *Process:
		stage.OnAfterProcessCreateCallback = any(callback).(OnAfterCreateInterface[Process])
	case *ProcessShape:
		stage.OnAfterProcessShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProcessShape])
	case *Task:
		stage.OnAfterTaskCreateCallback = any(callback).(OnAfterCreateInterface[Task])
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
	case *Participant:
		stage.OnAfterParticipantDeleteCallback = any(callback).(OnAfterDeleteInterface[Participant])
	case *ParticipantShape:
		stage.OnAfterParticipantShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ParticipantShape])
	case *Process:
		stage.OnAfterProcessDeleteCallback = any(callback).(OnAfterDeleteInterface[Process])
	case *ProcessShape:
		stage.OnAfterProcessShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProcessShape])
	case *Task:
		stage.OnAfterTaskDeleteCallback = any(callback).(OnAfterDeleteInterface[Task])
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
	case *Participant:
		stage.OnAfterParticipantReadCallback = any(callback).(OnAfterReadInterface[Participant])
	case *ParticipantShape:
		stage.OnAfterParticipantShapeReadCallback = any(callback).(OnAfterReadInterface[ParticipantShape])
	case *Process:
		stage.OnAfterProcessReadCallback = any(callback).(OnAfterReadInterface[Process])
	case *ProcessShape:
		stage.OnAfterProcessShapeReadCallback = any(callback).(OnAfterReadInterface[ProcessShape])
	case *Task:
		stage.OnAfterTaskReadCallback = any(callback).(OnAfterReadInterface[Task])
	}
}
