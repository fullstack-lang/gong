package models

import "slices"

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramProcess)

	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		_ = diagram

		updateMapElementDiagrams(stager, diagram, diagram.Process_Shapes, &diagram.map_Process_ProcessShape)
		updateMapElementDiagrams(stager, diagram, diagram.Participant_Shapes, &diagram.map_Participant_ParticipantShape)
		updateMapElementDiagrams(stager, diagram, diagram.Task_Shapes, &diagram.map_Task_TaskShape)
		updateMapElementDiagrams(stager, diagram, diagram.ControlFlow_Shapes, &diagram.map_ControlFlow_ControlFlowShape)
		updateMapElementDiagrams(stager, diagram, diagram.DataFlow_Shapes, &diagram.map_DataFlow_DataFlowShape)
	}
}

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
