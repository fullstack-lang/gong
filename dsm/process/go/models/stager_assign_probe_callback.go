package models

import (
	"log"
	"slices"
)

func (stager *Stager) assignProbeCallback() bool {
	if stager.probeForm == nil {
		return false
	}

	// We only assign it once. But since it's an assignment, doing it every time is harmless,
	// though we might want to ensure it's not nil before assigning.
	stager.probeForm.SetUpdateSliceOfPointersCallback(func(instance any, fieldName string, slicePtr any) {
		switch inst := instance.(type) {
		case *DataFlow:
			if slicePtr == &inst.Datas {
				stager.onDataFlowDatasUpdate(inst)
			}
		case *Note:
			switch slicePtr {
			case &inst.Tasks:
				stager.onNoteTasksUpdate(inst)
			case &inst.DataFlows:
				stager.onNoteDataFlowsUpdate(inst)
			case &inst.ControlFlows:
				stager.onNoteControlFlowsUpdate(inst)
			}
		case *Participant:
			if slicePtr == &inst.Resources {
				stager.onParticipantResourcesUpdate(inst)
			}
		case *ExternalParticipant:
			switch slicePtr {
			case &inst.Resources:
				stager.onExternalParticipantResourcesUpdate(inst)
			case &inst.Processes:
				stager.onExternalParticipantProcessesUpdate(inst)
			}
		}
	})

	return false
}

func (stager *Stager) onParticipantResourcesUpdate(participant *Participant) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		var valid []*AllocatedResourceShape
		for _, shape := range diagram.AllocatedResourceShapes {
			if shape.Participant == participant && (shape.Resource == nil || !slices.Contains(participant.Resources, shape.Resource)) {
				shape.Unstage(stager.stage)
				continue
			}
			valid = append(valid, shape)
		}
		diagram.AllocatedResourceShapes = valid
	}

	prevResources := stager.map_Participant_PrevResources[participant]
	if currentDiagram := stager.getCurrentDiagram(); currentDiagram != nil {
		var hasParticipantShape bool
		for _, pShape := range currentDiagram.Participant_Shapes {
			if pShape.Participant == participant {
				hasParticipantShape = true
				break
			}
		}
		if hasParticipantShape {
			for _, resource := range participant.Resources {
				if resource == nil || slices.Contains(prevResources, resource) {
					continue
				}
				var alreadyExists bool
				for _, existing := range currentDiagram.AllocatedResourceShapes {
					if existing.Participant == participant && existing.Resource == resource {
						alreadyExists = true
						break
					}
				}
				if !alreadyExists {
					allocatedResourceShape := (&AllocatedResourceShape{
						Name:        currentDiagram.GetName() + "-" + participant.GetName() + "-" + resource.GetName(),
						Participant: participant,
						Resource:    resource,
					}).Stage(stager.stage)
					currentDiagram.AllocatedResourceShapes = append(currentDiagram.AllocatedResourceShapes, allocatedResourceShape)
				}
			}
		}
	}

	if stager.map_Participant_PrevResources == nil {
		stager.map_Participant_PrevResources = make(map[*Participant][]*Resource)
	}
	stager.map_Participant_PrevResources[participant] = slices.Clone(participant.Resources)
}

func (stager *Stager) onExternalParticipantResourcesUpdate(externalParticipant *ExternalParticipant) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		var valid []*AllocatedResourceShape
		for _, shape := range diagram.AllocatedResourceShapes {
			if shape.ExternalParticipant == externalParticipant && (shape.Resource == nil || !slices.Contains(externalParticipant.Resources, shape.Resource)) {
				shape.Unstage(stager.stage)
				continue
			}
			valid = append(valid, shape)
		}
		diagram.AllocatedResourceShapes = valid
	}

	prevResources := stager.map_ExternalParticipant_PrevResources[externalParticipant]
	if currentDiagram := stager.getCurrentDiagram(); currentDiagram != nil {
		var hasExternalParticipantShape bool
		for _, epShape := range currentDiagram.ExternalParticipant_Shapes {
			if epShape.ExternalParticipant == externalParticipant {
				hasExternalParticipantShape = true
				break
			}
		}
		if hasExternalParticipantShape {
			for _, resource := range externalParticipant.Resources {
				if resource == nil || slices.Contains(prevResources, resource) {
					continue
				}
				var alreadyExists bool
				for _, existing := range currentDiagram.AllocatedResourceShapes {
					if existing.ExternalParticipant == externalParticipant && existing.Resource == resource {
						alreadyExists = true
						break
					}
				}
				if !alreadyExists {
					allocatedResourceShape := (&AllocatedResourceShape{
						Name:                currentDiagram.GetName() + "-" + externalParticipant.GetName() + "-" + resource.GetName(),
						ExternalParticipant: externalParticipant,
						Resource:            resource,
					}).Stage(stager.stage)
					currentDiagram.AllocatedResourceShapes = append(currentDiagram.AllocatedResourceShapes, allocatedResourceShape)
				}
			}
		}
	}

	if stager.map_ExternalParticipant_PrevResources == nil {
		stager.map_ExternalParticipant_PrevResources = make(map[*ExternalParticipant][]*Resource)
	}
	stager.map_ExternalParticipant_PrevResources[externalParticipant] = slices.Clone(externalParticipant.Resources)
}

func (stager *Stager) onExternalParticipantProcessesUpdate(externalParticipant *ExternalParticipant) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		var valid []*AllocatedProcessShape
		for _, shape := range diagram.AllocatedProcessShapes {
			if shape.ExternalParticipant == externalParticipant && (shape.Process == nil || !slices.Contains(externalParticipant.Processes, shape.Process)) {
				shape.Unstage(stager.stage)
				continue
			}
			valid = append(valid, shape)
		}
		diagram.AllocatedProcessShapes = valid
	}

	prevProcesses := stager.map_ExternalParticipant_PrevProcesses[externalParticipant]
	if currentDiagram := stager.getCurrentDiagram(); currentDiagram != nil {
		var hasExternalParticipantShape bool
		for _, epShape := range currentDiagram.ExternalParticipant_Shapes {
			if epShape.ExternalParticipant == externalParticipant {
				hasExternalParticipantShape = true
				break
			}
		}
		if hasExternalParticipantShape {
			for _, process := range externalParticipant.Processes {
				if process == nil || slices.Contains(prevProcesses, process) {
					continue
				}
				var alreadyExists bool
				for _, existing := range currentDiagram.AllocatedProcessShapes {
					if existing.ExternalParticipant == externalParticipant && existing.Process == process {
						alreadyExists = true
						break
					}
				}
				if !alreadyExists {
					allocatedProcessShape := (&AllocatedProcessShape{
						Name:                currentDiagram.GetName() + "-" + externalParticipant.GetName() + "-" + process.GetName(),
						ExternalParticipant: externalParticipant,
						Process:             process,
					}).Stage(stager.stage)
					currentDiagram.AllocatedProcessShapes = append(currentDiagram.AllocatedProcessShapes, allocatedProcessShape)
				}
			}
		}
	}

	if stager.map_ExternalParticipant_PrevProcesses == nil {
		stager.map_ExternalParticipant_PrevProcesses = make(map[*ExternalParticipant][]*Process)
	}
	stager.map_ExternalParticipant_PrevProcesses[externalParticipant] = slices.Clone(externalParticipant.Processes)
}

func (stager *Stager) onDataFlowDatasUpdate(dataFlow *DataFlow) {
	reachableDatas := collectReachableDatas(dataFlow.Datas)
	prevDatas, hasPrev := stager.map_DataFlow_PrevDatas[dataFlow]

	// when a Data is added/removed to/from a DataFlow
	// update all diagrams accordingly
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		// check if this diagram has a DataFlowShape for this dataFlow
		var hasDataFlowShape bool
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			if dataFlowShape.DataFlow == dataFlow {
				hasDataFlowShape = true
				break
			}
		}

		if hasDataFlowShape {
			existingDataShapes := make(map[*Data]*DataShape)
			var validDataShapes []*DataShape

			for _, dataShape := range diagram.Data_Shapes {
				if dataShape.DataFlow == dataFlow {
					if dataShape.Data != nil {
						if reachableDatas[dataShape.Data] {
							existingDataShapes[dataShape.Data] = dataShape
							validDataShapes = append(validDataShapes, dataShape)
						} else {
							dataShape.Unstage(stager.stage)
						}
					} else {
						dataShape.Unstage(stager.stage)
					}
				} else {
					validDataShapes = append(validDataShapes, dataShape)
				}
			}

			diagram.Data_Shapes = validDataShapes

			// add missing DataShapes for newly added top-level Data entries
			for _, data := range dataFlow.Datas {
				if data == nil {
					continue
				}
				if hasPrev && slices.Contains(prevDatas, data) {
					continue
				}
				if _, ok := existingDataShapes[data]; !ok {
					// create a new DataShape
					newDataShape := new(DataShape).Stage(stager.stage)
					newDataShape.Name = data.Name
					newDataShape.Data = data
					newDataShape.DataFlow = dataFlow
					diagram.Data_Shapes = append(diagram.Data_Shapes, newDataShape)
				}
			}
		}
	}

	if stager.map_DataFlow_PrevDatas == nil {
		stager.map_DataFlow_PrevDatas = make(map[*DataFlow][]*Data)
	}
	stager.map_DataFlow_PrevDatas[dataFlow] = slices.Clone(dataFlow.Datas)

	// The callback is invoked synchronously inside OnSave of the form,
	// and OnSave calls commit *after* the callback. But in case it doesn't,
	// we do not need to call commit explicitly if the form is saving.
	// However, if we modified diagram.Data_Shapes, we might want to log it.
	log.Printf("Updated Data_Shapes for DataFlow %s", dataFlow.Name)
}

func (stager *Stager) onNoteTasksUpdate(note *Note) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		var valid []*NoteTaskShape
		for _, shape := range diagram.NoteTaskShapes {
			if shape.Note == note && (shape.Task == nil || !slices.Contains(note.Tasks, shape.Task)) {
				shape.Unstage(stager.stage)
				continue
			}
			valid = append(valid, shape)
		}
		diagram.NoteTaskShapes = valid
	}

	prevTasks := stager.map_Note_PrevTasks[note]
	if currentDiagram := stager.getCurrentDiagram(); currentDiagram != nil {
		for _, task := range note.Tasks {
			if task == nil || slices.Contains(prevTasks, task) {
				continue
			}
			var alreadyExists bool
			for _, existing := range currentDiagram.NoteTaskShapes {
				if existing.Note == note && existing.Task == task {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				addAssociationShapeToDiagram(stager, note, task, &currentDiagram.NoteTaskShapes)
			}
		}
	}

	if stager.map_Note_PrevTasks == nil {
		stager.map_Note_PrevTasks = make(map[*Note][]*Task)
	}
	stager.map_Note_PrevTasks[note] = slices.Clone(note.Tasks)
}

func (stager *Stager) onNoteDataFlowsUpdate(note *Note) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		var valid []*NoteDataFlowShape
		for _, shape := range diagram.NoteDataFlowShapes {
			if shape.Note == note && (shape.DataFlow == nil || !slices.Contains(note.DataFlows, shape.DataFlow)) {
				shape.Unstage(stager.stage)
				continue
			}
			valid = append(valid, shape)
		}
		diagram.NoteDataFlowShapes = valid
	}

	prevDataFlows := stager.map_Note_PrevDataFlows[note]
	if currentDiagram := stager.getCurrentDiagram(); currentDiagram != nil {
		for _, dataFlow := range note.DataFlows {
			if dataFlow == nil || slices.Contains(prevDataFlows, dataFlow) {
				continue
			}
			var alreadyExists bool
			for _, existing := range currentDiagram.NoteDataFlowShapes {
				if existing.Note == note && existing.DataFlow == dataFlow {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				addAssociationShapeToDiagram(stager, note, dataFlow, &currentDiagram.NoteDataFlowShapes)
			}
		}
	}

	if stager.map_Note_PrevDataFlows == nil {
		stager.map_Note_PrevDataFlows = make(map[*Note][]*DataFlow)
	}
	stager.map_Note_PrevDataFlows[note] = slices.Clone(note.DataFlows)
}

func (stager *Stager) onNoteControlFlowsUpdate(note *Note) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		var valid []*NoteControlFlowShape
		for _, shape := range diagram.NoteControlFlowShapes {
			if shape.Note == note && (shape.ControlFlow == nil || !slices.Contains(note.ControlFlows, shape.ControlFlow)) {
				shape.Unstage(stager.stage)
				continue
			}
			valid = append(valid, shape)
		}
		diagram.NoteControlFlowShapes = valid
	}

	prevControlFlows := stager.map_Note_PrevControlFlows[note]
	if currentDiagram := stager.getCurrentDiagram(); currentDiagram != nil {
		for _, controlFlow := range note.ControlFlows {
			if controlFlow == nil || slices.Contains(prevControlFlows, controlFlow) {
				continue
			}
			var alreadyExists bool
			for _, existing := range currentDiagram.NoteControlFlowShapes {
				if existing.Note == note && existing.ControlFlow == controlFlow {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				addAssociationShapeToDiagram(stager, note, controlFlow, &currentDiagram.NoteControlFlowShapes)
			}
		}
	}

	if stager.map_Note_PrevControlFlows == nil {
		stager.map_Note_PrevControlFlows = make(map[*Note][]*ControlFlow)
	}
	stager.map_Note_PrevControlFlows[note] = slices.Clone(note.ControlFlows)
}
