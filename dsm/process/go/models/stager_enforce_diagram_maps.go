package models

import "slices"

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramProcess)

	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		_ = diagram

		diagram.map_Process_ProcessShape = make(map[*Process]*ProcessShape)
		for _, shape := range diagram.Process_Shapes {
			if shape.Process != nil {
				diagram.map_Process_ProcessShape[shape.Process] = shape
				var diagrams = stager.map_Element_Diagrams[shape.Process]

				if diagrams == nil {
					diagrams = []*DiagramProcess{diagram}
				}

				if !slices.Contains(diagrams, diagram) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Process] = diagrams
			}
		}
	}
}
