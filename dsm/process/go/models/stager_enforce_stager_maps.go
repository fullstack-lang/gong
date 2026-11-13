package models

import (
	"slices"
)

func (stager *Stager) enforceStagerMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramProcess)

	for _, diagramProcess := range stager.stage.GetInstancesSorted[*DiagramProcess]() {
		_ = diagramProcess

		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.Process_Shapes, &diagramProcess.map_Process_ProcessShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.Participant_Shapes, &diagramProcess.map_Participant_ParticipantShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.ExternalParticipant_Shapes, &diagramProcess.map_Participant_ExternalParticipantShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.Task_Shapes, &diagramProcess.map_Task_TaskShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.ControlFlow_Shapes, &diagramProcess.map_ControlFlow_ControlFlowShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.DataFlow_Shapes, &diagramProcess.map_DataFlow_DataFlowShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.Note_Shapes, &diagramProcess.map_Note_NoteShape)

		diagramProcess.map_DataShapeKey_DataShape = make(map[dataShapeKey]*DataShape)
		for _, dataShape := range diagramProcess.Data_Shapes {
			key := dataShapeKey{
				dataFlow: dataShape.DataFlow,
				data:     dataShape.Data,
			}
			diagramProcess.map_DataShapeKey_DataShape[key] = dataShape
		}

		diagramProcess.map_AllocatedResourceShapeKey_AllocatedResourceShape = make(map[allocatedResourceShapeKey]*AllocatedResourceShape)
		for _, allocatedResourceShape := range diagramProcess.AllocatedResourceShapes {
			key := allocatedResourceShapeKey{
				participant: allocatedResourceShape.Participant,
				resource:    allocatedResourceShape.Resource,
			}
			diagramProcess.map_AllocatedResourceShapeKey_AllocatedResourceShape[key] = allocatedResourceShape
		}

		diagramProcess.map_AllocatedProcessShapeKey_AllocatedProcessShape = make(map[allocatedProcessShapeKey]*AllocatedProcessShape)
		for _, allocatedProcessShape := range diagramProcess.AllocatedProcessShapes {
			key := allocatedProcessShapeKey{
				participant: allocatedProcessShape.Participant,
				process:     allocatedProcessShape.Process,
			}
			diagramProcess.map_AllocatedProcessShapeKey_AllocatedProcessShape[key] = allocatedProcessShape
		}

		diagramProcess.map_Note_NoteTaskShape = make(map[noteTaskKey]*NoteTaskShape)
		for _, noteTaskShape := range diagramProcess.NoteTaskShapes {
			key := noteTaskKey{
				Note: noteTaskShape.Note,
				Task: noteTaskShape.Task,
			}
			diagramProcess.map_Note_NoteTaskShape[key] = noteTaskShape
		}

		diagramProcess.map_Note_NoteDataFlowShape = make(map[noteDataFlowKey]*NoteDataFlowShape)
		for _, noteDataFlowShape := range diagramProcess.NoteDataFlowShapes {
			key := noteDataFlowKey{
				Note:     noteDataFlowShape.Note,
				DataFlow: noteDataFlowShape.DataFlow,
			}
			diagramProcess.map_Note_NoteDataFlowShape[key] = noteDataFlowShape
		}

		diagramProcess.map_Note_NoteControlFlowShape = make(map[noteControlFlowKey]*NoteControlFlowShape)
		for _, noteControlFlowShape := range diagramProcess.NoteControlFlowShapes {
			key := noteControlFlowKey{
				Note:        noteControlFlowShape.Note,
				ControlFlow: noteControlFlowShape.ControlFlow,
			}
			diagramProcess.map_Note_NoteControlFlowShape[key] = noteControlFlowShape
		}
	}

	stager.map_Note_PrevTasks = make(map[*Note][]*Task)
	stager.map_Note_PrevDataFlows = make(map[*Note][]*DataFlow)
	stager.map_Note_PrevControlFlows = make(map[*Note][]*ControlFlow)
	for _, note := range stager.stage.GetInstancesSorted[*Note]() {
		stager.map_Note_PrevTasks[note] = slices.Clone(note.Tasks)
		stager.map_Note_PrevDataFlows[note] = slices.Clone(note.DataFlows)
		stager.map_Note_PrevControlFlows[note] = slices.Clone(note.ControlFlows)
	}

	stager.rm_Data_DataFlows = stager.stage.GetSliceOfPointersReverseMap[DataFlow, Data](GongGetAssociationName[DataFlow]().Datas[0].Name)
	stager.rm_Resource_Participants = stager.stage.GetSliceOfPointersReverseMap[Participant, Resource](GongGetAssociationName[Participant]().Resources[0].Name)
}

// updateMapElementDiagrams is a helper function to update the map of abstract elements to their shapes for a given diagram
// and track the diagrams where each element is displayed.
func updateMapElementDiagrams[
	AT interface {
		AbstractType
		comparable
	},
	CT ConcreteType,
](
	stager *Stager,
	diagram *DiagramProcess,
	shapes []CT,
	diagramMapPtr *map[AT]CT, // Now a pointer to the map
) {
	// 1. Initialize the map right on the struct
	*diagramMapPtr = make(map[AT]CT)

	for _, shape := range shapes {
		abstractElement, ok := shape.GetAbstractElement().(AT)
		if !ok {
			continue
		}

		// 2. Dereference the pointer to map the abstract element to its Shape
		(*diagramMapPtr)[abstractElement] = shape

		// track all diagrams that display this element across the stage
		diagrams := stager.map_Element_Diagrams[abstractElement]
		if diagrams == nil {
			diagrams = []*DiagramProcess{diagram}
		}
		if !slices.Contains(diagrams, diagram) {
			diagrams = append(diagrams, diagram)
		}
		stager.map_Element_Diagrams[abstractElement] = diagrams
	}
}
