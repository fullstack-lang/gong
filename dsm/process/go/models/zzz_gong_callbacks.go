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
func (allocatedprocessshape *AllocatedProcessShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAllocatedProcessShapeCreateCallback != nil {
		stage.OnAfterAllocatedProcessShapeCreateCallback.OnAfterCreate(stage, allocatedprocessshape)
	}
}

func (allocatedprocessshape *AllocatedProcessShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllocatedProcessShapeUpdateCallback != nil {
		var frontAllocatedProcessShape *AllocatedProcessShape
		if front != nil {
			frontAllocatedProcessShape, _ = front.(*AllocatedProcessShape)
		}
		stage.OnAfterAllocatedProcessShapeUpdateCallback.OnAfterUpdate(stage, allocatedprocessshape, frontAllocatedProcessShape)
	}
}

func (allocatedprocessshape *AllocatedProcessShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllocatedProcessShapeDeleteCallback != nil {
		var frontAllocatedProcessShape *AllocatedProcessShape
		if front != nil {
			frontAllocatedProcessShape, _ = front.(*AllocatedProcessShape)
		}
		stage.OnAfterAllocatedProcessShapeDeleteCallback.OnAfterDelete(stage, allocatedprocessshape, frontAllocatedProcessShape)
	}
}

func (allocatedresourceshape *AllocatedResourceShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAllocatedResourceShapeCreateCallback != nil {
		stage.OnAfterAllocatedResourceShapeCreateCallback.OnAfterCreate(stage, allocatedresourceshape)
	}
}

func (allocatedresourceshape *AllocatedResourceShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllocatedResourceShapeUpdateCallback != nil {
		var frontAllocatedResourceShape *AllocatedResourceShape
		if front != nil {
			frontAllocatedResourceShape, _ = front.(*AllocatedResourceShape)
		}
		stage.OnAfterAllocatedResourceShapeUpdateCallback.OnAfterUpdate(stage, allocatedresourceshape, frontAllocatedResourceShape)
	}
}

func (allocatedresourceshape *AllocatedResourceShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllocatedResourceShapeDeleteCallback != nil {
		var frontAllocatedResourceShape *AllocatedResourceShape
		if front != nil {
			frontAllocatedResourceShape, _ = front.(*AllocatedResourceShape)
		}
		stage.OnAfterAllocatedResourceShapeDeleteCallback.OnAfterDelete(stage, allocatedresourceshape, frontAllocatedResourceShape)
	}
}

func (controlflow *ControlFlow) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterControlFlowCreateCallback != nil {
		stage.OnAfterControlFlowCreateCallback.OnAfterCreate(stage, controlflow)
	}
}

func (controlflow *ControlFlow) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlFlowUpdateCallback != nil {
		var frontControlFlow *ControlFlow
		if front != nil {
			frontControlFlow, _ = front.(*ControlFlow)
		}
		stage.OnAfterControlFlowUpdateCallback.OnAfterUpdate(stage, controlflow, frontControlFlow)
	}
}

func (controlflow *ControlFlow) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlFlowDeleteCallback != nil {
		var frontControlFlow *ControlFlow
		if front != nil {
			frontControlFlow, _ = front.(*ControlFlow)
		}
		stage.OnAfterControlFlowDeleteCallback.OnAfterDelete(stage, controlflow, frontControlFlow)
	}
}

func (controlflowshape *ControlFlowShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterControlFlowShapeCreateCallback != nil {
		stage.OnAfterControlFlowShapeCreateCallback.OnAfterCreate(stage, controlflowshape)
	}
}

func (controlflowshape *ControlFlowShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlFlowShapeUpdateCallback != nil {
		var frontControlFlowShape *ControlFlowShape
		if front != nil {
			frontControlFlowShape, _ = front.(*ControlFlowShape)
		}
		stage.OnAfterControlFlowShapeUpdateCallback.OnAfterUpdate(stage, controlflowshape, frontControlFlowShape)
	}
}

func (controlflowshape *ControlFlowShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlFlowShapeDeleteCallback != nil {
		var frontControlFlowShape *ControlFlowShape
		if front != nil {
			frontControlFlowShape, _ = front.(*ControlFlowShape)
		}
		stage.OnAfterControlFlowShapeDeleteCallback.OnAfterDelete(stage, controlflowshape, frontControlFlowShape)
	}
}

func (data *Data) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataCreateCallback != nil {
		stage.OnAfterDataCreateCallback.OnAfterCreate(stage, data)
	}
}

func (data *Data) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataUpdateCallback != nil {
		var frontData *Data
		if front != nil {
			frontData, _ = front.(*Data)
		}
		stage.OnAfterDataUpdateCallback.OnAfterUpdate(stage, data, frontData)
	}
}

func (data *Data) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataDeleteCallback != nil {
		var frontData *Data
		if front != nil {
			frontData, _ = front.(*Data)
		}
		stage.OnAfterDataDeleteCallback.OnAfterDelete(stage, data, frontData)
	}
}

func (dataflow *DataFlow) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataFlowCreateCallback != nil {
		stage.OnAfterDataFlowCreateCallback.OnAfterCreate(stage, dataflow)
	}
}

func (dataflow *DataFlow) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataFlowUpdateCallback != nil {
		var frontDataFlow *DataFlow
		if front != nil {
			frontDataFlow, _ = front.(*DataFlow)
		}
		stage.OnAfterDataFlowUpdateCallback.OnAfterUpdate(stage, dataflow, frontDataFlow)
	}
}

func (dataflow *DataFlow) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataFlowDeleteCallback != nil {
		var frontDataFlow *DataFlow
		if front != nil {
			frontDataFlow, _ = front.(*DataFlow)
		}
		stage.OnAfterDataFlowDeleteCallback.OnAfterDelete(stage, dataflow, frontDataFlow)
	}
}

func (dataflowshape *DataFlowShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataFlowShapeCreateCallback != nil {
		stage.OnAfterDataFlowShapeCreateCallback.OnAfterCreate(stage, dataflowshape)
	}
}

func (dataflowshape *DataFlowShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataFlowShapeUpdateCallback != nil {
		var frontDataFlowShape *DataFlowShape
		if front != nil {
			frontDataFlowShape, _ = front.(*DataFlowShape)
		}
		stage.OnAfterDataFlowShapeUpdateCallback.OnAfterUpdate(stage, dataflowshape, frontDataFlowShape)
	}
}

func (dataflowshape *DataFlowShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataFlowShapeDeleteCallback != nil {
		var frontDataFlowShape *DataFlowShape
		if front != nil {
			frontDataFlowShape, _ = front.(*DataFlowShape)
		}
		stage.OnAfterDataFlowShapeDeleteCallback.OnAfterDelete(stage, dataflowshape, frontDataFlowShape)
	}
}

func (datashape *DataShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataShapeCreateCallback != nil {
		stage.OnAfterDataShapeCreateCallback.OnAfterCreate(stage, datashape)
	}
}

func (datashape *DataShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataShapeUpdateCallback != nil {
		var frontDataShape *DataShape
		if front != nil {
			frontDataShape, _ = front.(*DataShape)
		}
		stage.OnAfterDataShapeUpdateCallback.OnAfterUpdate(stage, datashape, frontDataShape)
	}
}

func (datashape *DataShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataShapeDeleteCallback != nil {
		var frontDataShape *DataShape
		if front != nil {
			frontDataShape, _ = front.(*DataShape)
		}
		stage.OnAfterDataShapeDeleteCallback.OnAfterDelete(stage, datashape, frontDataShape)
	}
}

func (diagramprocess *DiagramProcess) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramProcessCreateCallback != nil {
		stage.OnAfterDiagramProcessCreateCallback.OnAfterCreate(stage, diagramprocess)
	}
}

func (diagramprocess *DiagramProcess) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramProcessUpdateCallback != nil {
		var frontDiagramProcess *DiagramProcess
		if front != nil {
			frontDiagramProcess, _ = front.(*DiagramProcess)
		}
		stage.OnAfterDiagramProcessUpdateCallback.OnAfterUpdate(stage, diagramprocess, frontDiagramProcess)
	}
}

func (diagramprocess *DiagramProcess) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramProcessDeleteCallback != nil {
		var frontDiagramProcess *DiagramProcess
		if front != nil {
			frontDiagramProcess, _ = front.(*DiagramProcess)
		}
		stage.OnAfterDiagramProcessDeleteCallback.OnAfterDelete(stage, diagramprocess, frontDiagramProcess)
	}
}

func (externalparticipantshape *ExternalParticipantShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterExternalParticipantShapeCreateCallback != nil {
		stage.OnAfterExternalParticipantShapeCreateCallback.OnAfterCreate(stage, externalparticipantshape)
	}
}

func (externalparticipantshape *ExternalParticipantShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExternalParticipantShapeUpdateCallback != nil {
		var frontExternalParticipantShape *ExternalParticipantShape
		if front != nil {
			frontExternalParticipantShape, _ = front.(*ExternalParticipantShape)
		}
		stage.OnAfterExternalParticipantShapeUpdateCallback.OnAfterUpdate(stage, externalparticipantshape, frontExternalParticipantShape)
	}
}

func (externalparticipantshape *ExternalParticipantShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExternalParticipantShapeDeleteCallback != nil {
		var frontExternalParticipantShape *ExternalParticipantShape
		if front != nil {
			frontExternalParticipantShape, _ = front.(*ExternalParticipantShape)
		}
		stage.OnAfterExternalParticipantShapeDeleteCallback.OnAfterDelete(stage, externalparticipantshape, frontExternalParticipantShape)
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

func (notetaskshape *NoteTaskShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteTaskShapeCreateCallback != nil {
		stage.OnAfterNoteTaskShapeCreateCallback.OnAfterCreate(stage, notetaskshape)
	}
}

func (notetaskshape *NoteTaskShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTaskShapeUpdateCallback != nil {
		var frontNoteTaskShape *NoteTaskShape
		if front != nil {
			frontNoteTaskShape, _ = front.(*NoteTaskShape)
		}
		stage.OnAfterNoteTaskShapeUpdateCallback.OnAfterUpdate(stage, notetaskshape, frontNoteTaskShape)
	}
}

func (notetaskshape *NoteTaskShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTaskShapeDeleteCallback != nil {
		var frontNoteTaskShape *NoteTaskShape
		if front != nil {
			frontNoteTaskShape, _ = front.(*NoteTaskShape)
		}
		stage.OnAfterNoteTaskShapeDeleteCallback.OnAfterDelete(stage, notetaskshape, frontNoteTaskShape)
	}
}

func (participant *Participant) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParticipantCreateCallback != nil {
		stage.OnAfterParticipantCreateCallback.OnAfterCreate(stage, participant)
	}
}

func (participant *Participant) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParticipantUpdateCallback != nil {
		var frontParticipant *Participant
		if front != nil {
			frontParticipant, _ = front.(*Participant)
		}
		stage.OnAfterParticipantUpdateCallback.OnAfterUpdate(stage, participant, frontParticipant)
	}
}

func (participant *Participant) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParticipantDeleteCallback != nil {
		var frontParticipant *Participant
		if front != nil {
			frontParticipant, _ = front.(*Participant)
		}
		stage.OnAfterParticipantDeleteCallback.OnAfterDelete(stage, participant, frontParticipant)
	}
}

func (participantshape *ParticipantShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParticipantShapeCreateCallback != nil {
		stage.OnAfterParticipantShapeCreateCallback.OnAfterCreate(stage, participantshape)
	}
}

func (participantshape *ParticipantShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParticipantShapeUpdateCallback != nil {
		var frontParticipantShape *ParticipantShape
		if front != nil {
			frontParticipantShape, _ = front.(*ParticipantShape)
		}
		stage.OnAfterParticipantShapeUpdateCallback.OnAfterUpdate(stage, participantshape, frontParticipantShape)
	}
}

func (participantshape *ParticipantShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParticipantShapeDeleteCallback != nil {
		var frontParticipantShape *ParticipantShape
		if front != nil {
			frontParticipantShape, _ = front.(*ParticipantShape)
		}
		stage.OnAfterParticipantShapeDeleteCallback.OnAfterDelete(stage, participantshape, frontParticipantShape)
	}
}

func (process *Process) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcessCreateCallback != nil {
		stage.OnAfterProcessCreateCallback.OnAfterCreate(stage, process)
	}
}

func (process *Process) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcessUpdateCallback != nil {
		var frontProcess *Process
		if front != nil {
			frontProcess, _ = front.(*Process)
		}
		stage.OnAfterProcessUpdateCallback.OnAfterUpdate(stage, process, frontProcess)
	}
}

func (process *Process) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcessDeleteCallback != nil {
		var frontProcess *Process
		if front != nil {
			frontProcess, _ = front.(*Process)
		}
		stage.OnAfterProcessDeleteCallback.OnAfterDelete(stage, process, frontProcess)
	}
}

func (processshape *ProcessShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcessShapeCreateCallback != nil {
		stage.OnAfterProcessShapeCreateCallback.OnAfterCreate(stage, processshape)
	}
}

func (processshape *ProcessShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcessShapeUpdateCallback != nil {
		var frontProcessShape *ProcessShape
		if front != nil {
			frontProcessShape, _ = front.(*ProcessShape)
		}
		stage.OnAfterProcessShapeUpdateCallback.OnAfterUpdate(stage, processshape, frontProcessShape)
	}
}

func (processshape *ProcessShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcessShapeDeleteCallback != nil {
		var frontProcessShape *ProcessShape
		if front != nil {
			frontProcessShape, _ = front.(*ProcessShape)
		}
		stage.OnAfterProcessShapeDeleteCallback.OnAfterDelete(stage, processshape, frontProcessShape)
	}
}

func (resource *Resource) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterResourceCreateCallback != nil {
		stage.OnAfterResourceCreateCallback.OnAfterCreate(stage, resource)
	}
}

func (resource *Resource) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceUpdateCallback != nil {
		var frontResource *Resource
		if front != nil {
			frontResource, _ = front.(*Resource)
		}
		stage.OnAfterResourceUpdateCallback.OnAfterUpdate(stage, resource, frontResource)
	}
}

func (resource *Resource) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceDeleteCallback != nil {
		var frontResource *Resource
		if front != nil {
			frontResource, _ = front.(*Resource)
		}
		stage.OnAfterResourceDeleteCallback.OnAfterDelete(stage, resource, frontResource)
	}
}

func (task *Task) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskCreateCallback != nil {
		stage.OnAfterTaskCreateCallback.OnAfterCreate(stage, task)
	}
}

func (task *Task) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskUpdateCallback != nil {
		var frontTask *Task
		if front != nil {
			frontTask, _ = front.(*Task)
		}
		stage.OnAfterTaskUpdateCallback.OnAfterUpdate(stage, task, frontTask)
	}
}

func (task *Task) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskDeleteCallback != nil {
		var frontTask *Task
		if front != nil {
			frontTask, _ = front.(*Task)
		}
		stage.OnAfterTaskDeleteCallback.OnAfterDelete(stage, task, frontTask)
	}
}

func (taskshape *TaskShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskShapeCreateCallback != nil {
		stage.OnAfterTaskShapeCreateCallback.OnAfterCreate(stage, taskshape)
	}
}

func (taskshape *TaskShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskShapeUpdateCallback != nil {
		var frontTaskShape *TaskShape
		if front != nil {
			frontTaskShape, _ = front.(*TaskShape)
		}
		stage.OnAfterTaskShapeUpdateCallback.OnAfterUpdate(stage, taskshape, frontTaskShape)
	}
}

func (taskshape *TaskShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskShapeDeleteCallback != nil {
		var frontTaskShape *TaskShape
		if front != nil {
			frontTaskShape, _ = front.(*TaskShape)
		}
		stage.OnAfterTaskShapeDeleteCallback.OnAfterDelete(stage, taskshape, frontTaskShape)
	}
}

