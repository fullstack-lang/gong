package models

import "slices"

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramProcess)

	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		_ = diagram

		diagram.map_Process_ProcessShape = make(map[*Process]*ProcessShape)
		updateMapElementDiagrams(stager, diagram, diagram.Process_Shapes, diagram.map_Process_ProcessShape)

		diagram.map_Participant_ParticipantShape = make(map[*Participant]*ParticipantShape)
		updateMapElementDiagrams(stager, diagram, diagram.Participant_Shapes, diagram.map_Participant_ParticipantShape)

		diagram.map_Task_TaskShape = make(map[*Task]*TaskShape)
		updateMapElementDiagrams(stager, diagram, diagram.Task_Shapes, diagram.map_Task_TaskShape)

		diagram.map_ControlFlow_ControlFlowShape = make(map[*ControlFlow]*ControlFlowShape)
		updateMapElementDiagrams(stager, diagram, diagram.ControlFlow_Shapes, diagram.map_ControlFlow_ControlFlowShape)

		diagram.map_DataFlow_DataFlowShape = make(map[*DataFlow]*DataFlowShape)
		updateMapElementDiagrams(stager, diagram, diagram.DataFlow_Shapes, diagram.map_DataFlow_DataFlowShape)
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
	diagramMap map[AT]CT,
) {
	for _, shape := range shapes {
		abstractElement, ok := shape.GetAbstractElement().(AT)
		if !ok {
			continue
		}

		// map the abstract element to its Shape within this diagram
		diagramMap[abstractElement] = shape

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
