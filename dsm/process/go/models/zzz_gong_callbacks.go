// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
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

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
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

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
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

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		if stage.OnAfterAllocatedResourceShapeReadCallback != nil {
			stage.OnAfterAllocatedResourceShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ControlFlow:
		if stage.OnAfterControlFlowReadCallback != nil {
			stage.OnAfterControlFlowReadCallback.OnAfterRead(stage, target)
		}
	case *ControlFlowShape:
		if stage.OnAfterControlFlowShapeReadCallback != nil {
			stage.OnAfterControlFlowShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Data:
		if stage.OnAfterDataReadCallback != nil {
			stage.OnAfterDataReadCallback.OnAfterRead(stage, target)
		}
	case *DataFlow:
		if stage.OnAfterDataFlowReadCallback != nil {
			stage.OnAfterDataFlowReadCallback.OnAfterRead(stage, target)
		}
	case *DataFlowShape:
		if stage.OnAfterDataFlowShapeReadCallback != nil {
			stage.OnAfterDataFlowShapeReadCallback.OnAfterRead(stage, target)
		}
	case *DataShape:
		if stage.OnAfterDataShapeReadCallback != nil {
			stage.OnAfterDataShapeReadCallback.OnAfterRead(stage, target)
		}
	case *DiagramProcess:
		if stage.OnAfterDiagramProcessReadCallback != nil {
			stage.OnAfterDiagramProcessReadCallback.OnAfterRead(stage, target)
		}
	case *ExternalParticipantShape:
		if stage.OnAfterExternalParticipantShapeReadCallback != nil {
			stage.OnAfterExternalParticipantShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteReadCallback != nil {
			stage.OnAfterNoteReadCallback.OnAfterRead(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeReadCallback != nil {
			stage.OnAfterNoteShapeReadCallback.OnAfterRead(stage, target)
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
	case *Resource:
		if stage.OnAfterResourceReadCallback != nil {
			stage.OnAfterResourceReadCallback.OnAfterRead(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskReadCallback != nil {
			stage.OnAfterTaskReadCallback.OnAfterRead(stage, target)
		}
	case *TaskShape:
		if stage.OnAfterTaskShapeReadCallback != nil {
			stage.OnAfterTaskShapeReadCallback.OnAfterRead(stage, target)
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
	case *AllocatedResourceShape:
		stage.OnAfterAllocatedResourceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[AllocatedResourceShape])
	case *ControlFlow:
		stage.OnAfterControlFlowUpdateCallback = any(callback).(OnAfterUpdateInterface[ControlFlow])
	case *ControlFlowShape:
		stage.OnAfterControlFlowShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ControlFlowShape])
	case *Data:
		stage.OnAfterDataUpdateCallback = any(callback).(OnAfterUpdateInterface[Data])
	case *DataFlow:
		stage.OnAfterDataFlowUpdateCallback = any(callback).(OnAfterUpdateInterface[DataFlow])
	case *DataFlowShape:
		stage.OnAfterDataFlowShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[DataFlowShape])
	case *DataShape:
		stage.OnAfterDataShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[DataShape])
	case *DiagramProcess:
		stage.OnAfterDiagramProcessUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramProcess])
	case *ExternalParticipantShape:
		stage.OnAfterExternalParticipantShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ExternalParticipantShape])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Note:
		stage.OnAfterNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[Note])
	case *NoteShape:
		stage.OnAfterNoteShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteShape])
	case *Participant:
		stage.OnAfterParticipantUpdateCallback = any(callback).(OnAfterUpdateInterface[Participant])
	case *ParticipantShape:
		stage.OnAfterParticipantShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ParticipantShape])
	case *Process:
		stage.OnAfterProcessUpdateCallback = any(callback).(OnAfterUpdateInterface[Process])
	case *ProcessShape:
		stage.OnAfterProcessShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProcessShape])
	case *Resource:
		stage.OnAfterResourceUpdateCallback = any(callback).(OnAfterUpdateInterface[Resource])
	case *Task:
		stage.OnAfterTaskUpdateCallback = any(callback).(OnAfterUpdateInterface[Task])
	case *TaskShape:
		stage.OnAfterTaskShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TaskShape])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		stage.OnAfterAllocatedResourceShapeCreateCallback = any(callback).(OnAfterCreateInterface[AllocatedResourceShape])
	case *ControlFlow:
		stage.OnAfterControlFlowCreateCallback = any(callback).(OnAfterCreateInterface[ControlFlow])
	case *ControlFlowShape:
		stage.OnAfterControlFlowShapeCreateCallback = any(callback).(OnAfterCreateInterface[ControlFlowShape])
	case *Data:
		stage.OnAfterDataCreateCallback = any(callback).(OnAfterCreateInterface[Data])
	case *DataFlow:
		stage.OnAfterDataFlowCreateCallback = any(callback).(OnAfterCreateInterface[DataFlow])
	case *DataFlowShape:
		stage.OnAfterDataFlowShapeCreateCallback = any(callback).(OnAfterCreateInterface[DataFlowShape])
	case *DataShape:
		stage.OnAfterDataShapeCreateCallback = any(callback).(OnAfterCreateInterface[DataShape])
	case *DiagramProcess:
		stage.OnAfterDiagramProcessCreateCallback = any(callback).(OnAfterCreateInterface[DiagramProcess])
	case *ExternalParticipantShape:
		stage.OnAfterExternalParticipantShapeCreateCallback = any(callback).(OnAfterCreateInterface[ExternalParticipantShape])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Note:
		stage.OnAfterNoteCreateCallback = any(callback).(OnAfterCreateInterface[Note])
	case *NoteShape:
		stage.OnAfterNoteShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteShape])
	case *Participant:
		stage.OnAfterParticipantCreateCallback = any(callback).(OnAfterCreateInterface[Participant])
	case *ParticipantShape:
		stage.OnAfterParticipantShapeCreateCallback = any(callback).(OnAfterCreateInterface[ParticipantShape])
	case *Process:
		stage.OnAfterProcessCreateCallback = any(callback).(OnAfterCreateInterface[Process])
	case *ProcessShape:
		stage.OnAfterProcessShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProcessShape])
	case *Resource:
		stage.OnAfterResourceCreateCallback = any(callback).(OnAfterCreateInterface[Resource])
	case *Task:
		stage.OnAfterTaskCreateCallback = any(callback).(OnAfterCreateInterface[Task])
	case *TaskShape:
		stage.OnAfterTaskShapeCreateCallback = any(callback).(OnAfterCreateInterface[TaskShape])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		stage.OnAfterAllocatedResourceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[AllocatedResourceShape])
	case *ControlFlow:
		stage.OnAfterControlFlowDeleteCallback = any(callback).(OnAfterDeleteInterface[ControlFlow])
	case *ControlFlowShape:
		stage.OnAfterControlFlowShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ControlFlowShape])
	case *Data:
		stage.OnAfterDataDeleteCallback = any(callback).(OnAfterDeleteInterface[Data])
	case *DataFlow:
		stage.OnAfterDataFlowDeleteCallback = any(callback).(OnAfterDeleteInterface[DataFlow])
	case *DataFlowShape:
		stage.OnAfterDataFlowShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[DataFlowShape])
	case *DataShape:
		stage.OnAfterDataShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[DataShape])
	case *DiagramProcess:
		stage.OnAfterDiagramProcessDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramProcess])
	case *ExternalParticipantShape:
		stage.OnAfterExternalParticipantShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ExternalParticipantShape])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Note:
		stage.OnAfterNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[Note])
	case *NoteShape:
		stage.OnAfterNoteShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteShape])
	case *Participant:
		stage.OnAfterParticipantDeleteCallback = any(callback).(OnAfterDeleteInterface[Participant])
	case *ParticipantShape:
		stage.OnAfterParticipantShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ParticipantShape])
	case *Process:
		stage.OnAfterProcessDeleteCallback = any(callback).(OnAfterDeleteInterface[Process])
	case *ProcessShape:
		stage.OnAfterProcessShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProcessShape])
	case *Resource:
		stage.OnAfterResourceDeleteCallback = any(callback).(OnAfterDeleteInterface[Resource])
	case *Task:
		stage.OnAfterTaskDeleteCallback = any(callback).(OnAfterDeleteInterface[Task])
	case *TaskShape:
		stage.OnAfterTaskShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TaskShape])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AllocatedResourceShape:
		stage.OnAfterAllocatedResourceShapeReadCallback = any(callback).(OnAfterReadInterface[AllocatedResourceShape])
	case *ControlFlow:
		stage.OnAfterControlFlowReadCallback = any(callback).(OnAfterReadInterface[ControlFlow])
	case *ControlFlowShape:
		stage.OnAfterControlFlowShapeReadCallback = any(callback).(OnAfterReadInterface[ControlFlowShape])
	case *Data:
		stage.OnAfterDataReadCallback = any(callback).(OnAfterReadInterface[Data])
	case *DataFlow:
		stage.OnAfterDataFlowReadCallback = any(callback).(OnAfterReadInterface[DataFlow])
	case *DataFlowShape:
		stage.OnAfterDataFlowShapeReadCallback = any(callback).(OnAfterReadInterface[DataFlowShape])
	case *DataShape:
		stage.OnAfterDataShapeReadCallback = any(callback).(OnAfterReadInterface[DataShape])
	case *DiagramProcess:
		stage.OnAfterDiagramProcessReadCallback = any(callback).(OnAfterReadInterface[DiagramProcess])
	case *ExternalParticipantShape:
		stage.OnAfterExternalParticipantShapeReadCallback = any(callback).(OnAfterReadInterface[ExternalParticipantShape])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Note:
		stage.OnAfterNoteReadCallback = any(callback).(OnAfterReadInterface[Note])
	case *NoteShape:
		stage.OnAfterNoteShapeReadCallback = any(callback).(OnAfterReadInterface[NoteShape])
	case *Participant:
		stage.OnAfterParticipantReadCallback = any(callback).(OnAfterReadInterface[Participant])
	case *ParticipantShape:
		stage.OnAfterParticipantShapeReadCallback = any(callback).(OnAfterReadInterface[ParticipantShape])
	case *Process:
		stage.OnAfterProcessReadCallback = any(callback).(OnAfterReadInterface[Process])
	case *ProcessShape:
		stage.OnAfterProcessShapeReadCallback = any(callback).(OnAfterReadInterface[ProcessShape])
	case *Resource:
		stage.OnAfterResourceReadCallback = any(callback).(OnAfterReadInterface[Resource])
	case *Task:
		stage.OnAfterTaskReadCallback = any(callback).(OnAfterReadInterface[Task])
	case *TaskShape:
		stage.OnAfterTaskShapeReadCallback = any(callback).(OnAfterReadInterface[TaskShape])
	}
}
