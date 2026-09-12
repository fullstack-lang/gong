// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AllocatedProcessShape:
		if stage.OnAfterAllocatedProcessShapeCreateCallback != nil {
			stage.OnAfterAllocatedProcessShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *AllocatedResourceShape:
		if stage.OnAfterAllocatedResourceShapeCreateCallback != nil {
			stage.OnAfterAllocatedResourceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ControlFlow:
		if stage.OnAfterControlFlowCreateCallback != nil {
			stage.OnAfterControlFlowCreateCallback.OnAfterCreate(stage, target)
		}
	case *ControlFlowShape:
		if stage.OnAfterControlFlowShapeCreateCallback != nil {
			stage.OnAfterControlFlowShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Data:
		if stage.OnAfterDataCreateCallback != nil {
			stage.OnAfterDataCreateCallback.OnAfterCreate(stage, target)
		}
	case *DataFlow:
		if stage.OnAfterDataFlowCreateCallback != nil {
			stage.OnAfterDataFlowCreateCallback.OnAfterCreate(stage, target)
		}
	case *DataFlowShape:
		if stage.OnAfterDataFlowShapeCreateCallback != nil {
			stage.OnAfterDataFlowShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *DataShape:
		if stage.OnAfterDataShapeCreateCallback != nil {
			stage.OnAfterDataShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *DiagramProcess:
		if stage.OnAfterDiagramProcessCreateCallback != nil {
			stage.OnAfterDiagramProcessCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExternalParticipantShape:
		if stage.OnAfterExternalParticipantShapeCreateCallback != nil {
			stage.OnAfterExternalParticipantShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeCreateCallback != nil {
			stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeCreateCallback != nil {
			stage.OnAfterNoteTaskShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *Resource:
		if stage.OnAfterResourceCreateCallback != nil {
			stage.OnAfterResourceCreateCallback.OnAfterCreate(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskCreateCallback != nil {
			stage.OnAfterTaskCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskShape:
		if stage.OnAfterTaskShapeCreateCallback != nil {
			stage.OnAfterTaskShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *AllocatedProcessShape:
		newTarget := any(new).(*AllocatedProcessShape)
		if stage.OnAfterAllocatedProcessShapeUpdateCallback != nil {
			stage.OnAfterAllocatedProcessShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AllocatedResourceShape:
		newTarget := any(new).(*AllocatedResourceShape)
		if stage.OnAfterAllocatedResourceShapeUpdateCallback != nil {
			stage.OnAfterAllocatedResourceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ControlFlow:
		newTarget := any(new).(*ControlFlow)
		if stage.OnAfterControlFlowUpdateCallback != nil {
			stage.OnAfterControlFlowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ControlFlowShape:
		newTarget := any(new).(*ControlFlowShape)
		if stage.OnAfterControlFlowShapeUpdateCallback != nil {
			stage.OnAfterControlFlowShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Data:
		newTarget := any(new).(*Data)
		if stage.OnAfterDataUpdateCallback != nil {
			stage.OnAfterDataUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DataFlow:
		newTarget := any(new).(*DataFlow)
		if stage.OnAfterDataFlowUpdateCallback != nil {
			stage.OnAfterDataFlowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DataFlowShape:
		newTarget := any(new).(*DataFlowShape)
		if stage.OnAfterDataFlowShapeUpdateCallback != nil {
			stage.OnAfterDataFlowShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DataShape:
		newTarget := any(new).(*DataShape)
		if stage.OnAfterDataShapeUpdateCallback != nil {
			stage.OnAfterDataShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DiagramProcess:
		newTarget := any(new).(*DiagramProcess)
		if stage.OnAfterDiagramProcessUpdateCallback != nil {
			stage.OnAfterDiagramProcessUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ExternalParticipantShape:
		newTarget := any(new).(*ExternalParticipantShape)
		if stage.OnAfterExternalParticipantShapeUpdateCallback != nil {
			stage.OnAfterExternalParticipantShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *NoteShape:
		newTarget := any(new).(*NoteShape)
		if stage.OnAfterNoteShapeUpdateCallback != nil {
			stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteTaskShape:
		newTarget := any(new).(*NoteTaskShape)
		if stage.OnAfterNoteTaskShapeUpdateCallback != nil {
			stage.OnAfterNoteTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Resource:
		newTarget := any(new).(*Resource)
		if stage.OnAfterResourceUpdateCallback != nil {
			stage.OnAfterResourceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Task:
		newTarget := any(new).(*Task)
		if stage.OnAfterTaskUpdateCallback != nil {
			stage.OnAfterTaskUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskShape:
		newTarget := any(new).(*TaskShape)
		if stage.OnAfterTaskShapeUpdateCallback != nil {
			stage.OnAfterTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *AllocatedProcessShape:
		if stage.OnAfterAllocatedProcessShapeDeleteCallback != nil {
			staged := any(staged).(*AllocatedProcessShape)
			stage.OnAfterAllocatedProcessShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AllocatedResourceShape:
		if stage.OnAfterAllocatedResourceShapeDeleteCallback != nil {
			staged := any(staged).(*AllocatedResourceShape)
			stage.OnAfterAllocatedResourceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ControlFlow:
		if stage.OnAfterControlFlowDeleteCallback != nil {
			staged := any(staged).(*ControlFlow)
			stage.OnAfterControlFlowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ControlFlowShape:
		if stage.OnAfterControlFlowShapeDeleteCallback != nil {
			staged := any(staged).(*ControlFlowShape)
			stage.OnAfterControlFlowShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Data:
		if stage.OnAfterDataDeleteCallback != nil {
			staged := any(staged).(*Data)
			stage.OnAfterDataDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DataFlow:
		if stage.OnAfterDataFlowDeleteCallback != nil {
			staged := any(staged).(*DataFlow)
			stage.OnAfterDataFlowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DataFlowShape:
		if stage.OnAfterDataFlowShapeDeleteCallback != nil {
			staged := any(staged).(*DataFlowShape)
			stage.OnAfterDataFlowShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DataShape:
		if stage.OnAfterDataShapeDeleteCallback != nil {
			staged := any(staged).(*DataShape)
			stage.OnAfterDataShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DiagramProcess:
		if stage.OnAfterDiagramProcessDeleteCallback != nil {
			staged := any(staged).(*DiagramProcess)
			stage.OnAfterDiagramProcessDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ExternalParticipantShape:
		if stage.OnAfterExternalParticipantShapeDeleteCallback != nil {
			staged := any(staged).(*ExternalParticipantShape)
			stage.OnAfterExternalParticipantShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *NoteShape:
		if stage.OnAfterNoteShapeDeleteCallback != nil {
			staged := any(staged).(*NoteShape)
			stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeDeleteCallback != nil {
			staged := any(staged).(*NoteTaskShape)
			stage.OnAfterNoteTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Resource:
		if stage.OnAfterResourceDeleteCallback != nil {
			staged := any(staged).(*Resource)
			stage.OnAfterResourceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Task:
		if stage.OnAfterTaskDeleteCallback != nil {
			staged := any(staged).(*Task)
			stage.OnAfterTaskDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskShape:
		if stage.OnAfterTaskShapeDeleteCallback != nil {
			staged := any(staged).(*TaskShape)
			stage.OnAfterTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
