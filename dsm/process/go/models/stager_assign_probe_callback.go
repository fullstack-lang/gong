package models

import "slices"

import "log"

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
		}
	})

	return false
}

func (stager *Stager) onDataFlowDatasUpdate(dataFlow *DataFlow) {
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
						existingDataShapes[dataShape.Data] = dataShape

						stillExists := slices.Contains(dataFlow.Datas, dataShape.Data)
						if stillExists {
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

			// add missing DataShapes
			for _, data := range dataFlow.Datas {
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
