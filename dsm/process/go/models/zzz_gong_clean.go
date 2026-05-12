// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by AllocatedResourceShape
func (allocatedresourceshape *AllocatedResourceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &allocatedresourceshape.Participant) || modified
	modified = GongCleanPointer(stage, &allocatedresourceshape.Resource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlFlow
func (controlflow *ControlFlow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &controlflow.Start) || modified
	modified = GongCleanPointer(stage, &controlflow.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlFlowShape
func (controlflowshape *ControlFlowShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &controlflowshape.ControlFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Data
func (data *Data) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DataFlow
func (dataflow *DataFlow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &dataflow.Datas) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &dataflow.StartTask) || modified
	modified = GongCleanPointer(stage, &dataflow.EndTask) || modified
	modified = GongCleanPointer(stage, &dataflow.StartExternalParticipant) || modified
	modified = GongCleanPointer(stage, &dataflow.EndExternalParticipant) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataFlowShape
func (dataflowshape *DataFlowShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &dataflowshape.DataFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataShape
func (datashape *DataShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datashape.Data) || modified
	modified = GongCleanPointer(stage, &datashape.DataFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramProcess
func (diagramprocess *DiagramProcess) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagramprocess.Process_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.ProcesssWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.Participant_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.ParticipantWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.ExternalParticipant_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.ExternalParticipantWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.TasksWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.Task_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.ControlFlowsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.ControlFlow_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.DataFlowsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.DataFlow_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.DatasWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.Data_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.DataFlowsWhoseDataNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.AllocatedResourcesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.AllocatedResourceShapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.Note_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramprocess.NotesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramprocess.NoteTaskShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ExternalParticipantShape
func (externalparticipantshape *ExternalParticipantShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &externalparticipantshape.Participant) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	modified = GongCleanSlice(stage, &library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootProcesses) || modified
	modified = GongCleanSlice(stage, &library.ProcesssWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootDataFlows) || modified
	modified = GongCleanSlice(stage, &library.DataFlowsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootDatas) || modified
	modified = GongCleanSlice(stage, &library.DatasWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootResources) || modified
	modified = GongCleanSlice(stage, &library.ResourcesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootNotes) || modified
	modified = GongCleanSlice(stage, &library.NotesWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &note.Tasks) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteTaskShape
func (notetaskshape *NoteTaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &notetaskshape.Note) || modified
	modified = GongCleanPointer(stage, &notetaskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Participant
func (participant *Participant) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &participant.Resources) || modified
	modified = GongCleanSlice(stage, &participant.Tasks) || modified
	modified = GongCleanSlice(stage, &participant.ControlFlows) || modified
	modified = GongCleanSlice(stage, &participant.TaskWhoseOutControlFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &participant.TaskWhoseInControlFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &participant.TaskWhoseOutDataFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &participant.TaskWhoseInDataFlowsNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParticipantShape
func (participantshape *ParticipantShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &participantshape.Participant) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Process
func (process *Process) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &process.DiagramProcesss) || modified
	modified = GongCleanSlice(stage, &process.DiagramProcessWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &process.SubProcesses) || modified
	modified = GongCleanSlice(stage, &process.Participants) || modified
	modified = GongCleanSlice(stage, &process.ParticipantWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &process.DataFlows) || modified
	modified = GongCleanSlice(stage, &process.ExternalParticipants) || modified
	modified = GongCleanSlice(stage, &process.ExternalParticipantWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcessShape
func (processshape *ProcessShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &processshape.Process) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Resource
func (resource *Resource) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Task
func (task *Task) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &task.Type) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskShape
func (taskshape *TaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &taskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
