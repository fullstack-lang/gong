package models

import (
	"slices"
)

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramProcess)

	for _, diagramProcess := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		_ = diagramProcess

		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.Process_Shapes, &diagramProcess.map_Process_ProcessShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.Participant_Shapes, &diagramProcess.map_Participant_ParticipantShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.ExternalParticipant_Shapes, &diagramProcess.map_Participant_ExternalParticipantShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.Task_Shapes, &diagramProcess.map_Task_TaskShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.ControlFlow_Shapes, &diagramProcess.map_ControlFlow_ControlFlowShape)
		updateMapElementDiagrams(stager, diagramProcess, diagramProcess.DataFlow_Shapes, &diagramProcess.map_DataFlow_DataFlowShape)

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
	}
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
