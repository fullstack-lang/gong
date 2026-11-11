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
			case &inst.Parts:
				stager.onNotePartsUpdate(inst)
			case &inst.Ports:
				stager.onNotePortsUpdate(inst)
			case &inst.DataFlows:
				stager.onNoteDataFlowsUpdate(inst)
			}
		}
	})

	return false
}

func (stager *Stager) onDataFlowDatasUpdate(dataFlow *DataFlow) {
	// when a Data is added/removed to/from a DataFlow
	// update all diagrams accordingly
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {
		var matchingDFShapes []*DataFlowShape
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			if dataFlowShape.DataFlow == dataFlow {
				matchingDFShapes = append(matchingDFShapes, dataFlowShape)
			}
		}

		if len(matchingDFShapes) > 0 {
			type dfDataPair struct {
				dfShape *DataFlowShape
				data    *Data
			}
			existingDataShapes := make(map[dfDataPair]*DataShape)
			var validDataShapes []*DataShape

			for _, dataShape := range diagram.Data_Shapes {
				if dataShape.DataFlow == dataFlow {
					if dataShape.Data != nil {
						effDFS := effectiveDataFlowShape(diagram, dataFlow, dataShape.DataFlowShape)
						stillExists := slices.Contains(dataFlow.Datas, dataShape.Data) &&
							(effDFS != nil && slices.Contains(matchingDFShapes, effDFS))
						if stillExists {
							existingDataShapes[dfDataPair{dfShape: effDFS, data: dataShape.Data}] = dataShape
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

			// add missing DataShapes for each active DataFlowShape of this dataFlow
			for _, dfShape := range matchingDFShapes {
				for _, data := range dataFlow.Datas {
					pair := dfDataPair{dfShape: dfShape, data: data}
					if _, ok := existingDataShapes[pair]; !ok {
						newDataShape := new(DataShape).Stage(stager.stage)
						newDataShape.Name = data.Name
						newDataShape.Data = data
						newDataShape.DataFlow = dataFlow
						newDataShape.DataFlowShape = dfShape
						diagram.Data_Shapes = append(diagram.Data_Shapes, newDataShape)
					}
				}
			}
		}
	}

	log.Printf("Updated Data_Shapes for DataFlow %s", dataFlow.Name)
}

func (stager *Stager) onNotePartsUpdate(note *Note) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {
		var valid []*NotePartShape
		for _, shape := range diagram.NotePartShapes {
			if shape.Note == note && (shape.Part == nil || !slices.Contains(note.Parts, shape.Part)) {
				shape.Unstage(stager.stage)
				continue
			}
			valid = append(valid, shape)
		}
		diagram.NotePartShapes = valid
	}

	prevParts := stager.map_Note_PrevParts[note]
	if currentDiagram := stager.getCurrentDiagram(); currentDiagram != nil {
		for _, part := range note.Parts {
			if part == nil || slices.Contains(prevParts, part) {
				continue
			}
			var alreadyExists bool
			for _, existing := range currentDiagram.NotePartShapes {
				if existing.Note == note && existing.Part == part {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				addAssociationShapeToDiagram(stager, note, part, &currentDiagram.NotePartShapes)
			}
		}
	}

	if stager.map_Note_PrevParts == nil {
		stager.map_Note_PrevParts = make(map[*Note][]*Part)
	}
	stager.map_Note_PrevParts[note] = slices.Clone(note.Parts)
}

func (stager *Stager) onNotePortsUpdate(note *Note) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {
		var valid []*NotePortShape
		for _, shape := range diagram.NotePortShapes {
			if shape.Note == note && (shape.Port == nil || !slices.Contains(note.Ports, shape.Port)) {
				shape.Unstage(stager.stage)
				continue
			}
			valid = append(valid, shape)
		}
		diagram.NotePortShapes = valid
	}

	prevPorts := stager.map_Note_PrevPorts[note]
	if currentDiagram := stager.getCurrentDiagram(); currentDiagram != nil {
		for _, port := range note.Ports {
			if port == nil || slices.Contains(prevPorts, port) {
				continue
			}
			var alreadyExists bool
			for _, existing := range currentDiagram.NotePortShapes {
				if existing.Note == note && existing.Port == port {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				addAssociationShapeToDiagram(stager, note, port, &currentDiagram.NotePortShapes)
			}
		}
	}

	if stager.map_Note_PrevPorts == nil {
		stager.map_Note_PrevPorts = make(map[*Note][]*Port)
	}
	stager.map_Note_PrevPorts[note] = slices.Clone(note.Ports)
}

func (stager *Stager) onNoteDataFlowsUpdate(note *Note) {
	for _, diagram := range stager.stage.GetInstancesSorted[*DiagramStructure]() {
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
