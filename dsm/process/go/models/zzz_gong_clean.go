// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

type GongCleaner interface {
	GongClean(stage *Stage) (modified bool)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by AllocatedProcessShape
func (allocatedprocessshape *AllocatedProcessShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&allocatedprocessshape.Participant) || modified
	modified = stage.CleanPointer(&allocatedprocessshape.Process) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AllocatedResourceShape
func (allocatedresourceshape *AllocatedResourceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&allocatedresourceshape.Participant) || modified
	modified = stage.CleanPointer(&allocatedresourceshape.Resource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlFlow
func (controlflow *ControlFlow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&controlflow.Start) || modified
	modified = stage.CleanPointer(&controlflow.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlFlowShape
func (controlflowshape *ControlFlowShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&controlflowshape.ControlFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataFlow
func (dataflow *DataFlow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&dataflow.Datas) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dataflow.StartTask) || modified
	modified = stage.CleanPointer(&dataflow.EndTask) || modified
	modified = stage.CleanPointer(&dataflow.StartExternalParticipant) || modified
	modified = stage.CleanPointer(&dataflow.EndExternalParticipant) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataFlowShape
func (dataflowshape *DataFlowShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dataflowshape.DataFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataShape
func (datashape *DataShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&datashape.Data) || modified
	modified = stage.CleanPointer(&datashape.DataFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramProcess
func (diagramprocess *DiagramProcess) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&diagramprocess.Process_Shapes) || modified
	modified = stage.CleanSlice(&diagramprocess.ProcesssWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.Participant_Shapes) || modified
	modified = stage.CleanSlice(&diagramprocess.ParticipantWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.ExternalParticipant_Shapes) || modified
	modified = stage.CleanSlice(&diagramprocess.ExternalParticipantWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.ExternalParticipantsWhoseInDataFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.TasksWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.Task_Shapes) || modified
	modified = stage.CleanSlice(&diagramprocess.ControlFlowsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.ControlFlow_Shapes) || modified
	modified = stage.CleanSlice(&diagramprocess.DataFlowsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.DataFlow_Shapes) || modified
	modified = stage.CleanSlice(&diagramprocess.DatasWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.Data_Shapes) || modified
	modified = stage.CleanSlice(&diagramprocess.DataFlowsWhoseDataNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.AllocatedResourcesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.AllocatedResourceShapes) || modified
	modified = stage.CleanSlice(&diagramprocess.AllocatedProcessesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.AllocatedProcessShapes) || modified
	modified = stage.CleanSlice(&diagramprocess.Note_Shapes) || modified
	modified = stage.CleanSlice(&diagramprocess.NotesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramprocess.NoteTaskShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ExternalParticipantShape
func (externalparticipantshape *ExternalParticipantShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&externalparticipantshape.Participant) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	modified = stage.CleanSlice(&library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootProcesses) || modified
	modified = stage.CleanSlice(&library.ProcesssWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootDataFlows) || modified
	modified = stage.CleanSlice(&library.DataFlowsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootDatas) || modified
	modified = stage.CleanSlice(&library.DatasWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootResources) || modified
	modified = stage.CleanSlice(&library.ResourcesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.ParticipantsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootNotes) || modified
	modified = stage.CleanSlice(&library.NotesWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&note.Tasks) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteTaskShape
func (notetaskshape *NoteTaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&notetaskshape.Note) || modified
	modified = stage.CleanPointer(&notetaskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Participant
func (participant *Participant) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&participant.Resources) || modified
	modified = stage.CleanSlice(&participant.Processes) || modified
	modified = stage.CleanSlice(&participant.Tasks) || modified
	modified = stage.CleanSlice(&participant.ControlFlows) || modified
	modified = stage.CleanSlice(&participant.TaskWhoseOutControlFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&participant.TaskWhoseInControlFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&participant.TaskWhoseOutDataFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&participant.TaskWhoseInDataFlowsNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ParticipantShape
func (participantshape *ParticipantShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&participantshape.Participant) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Process
func (process *Process) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&process.DiagramProcesss) || modified
	modified = stage.CleanSlice(&process.DiagramProcessWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&process.SubProcesses) || modified
	modified = stage.CleanSlice(&process.Participants) || modified
	modified = stage.CleanSlice(&process.ParticipantWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&process.DataFlows) || modified
	modified = stage.CleanSlice(&process.ExternalParticipants) || modified
	modified = stage.CleanSlice(&process.ExternalParticipantWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ProcessShape
func (processshape *ProcessShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&processshape.Process) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Task
func (task *Task) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&task.Type) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskShape
func (taskshape *TaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		if cleaner, ok := any(instance).(GongCleaner); ok {
			modified = cleaner.GongClean(stage) || modified
		}
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
