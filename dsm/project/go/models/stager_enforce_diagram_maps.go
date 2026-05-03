package models

import "slices"

func (stager *Stager) enforceDiagramMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]DiagramIF)

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		// computes all products presents in the diagram
		diagram.map_Product_ProductShape = make(map[*Product]*ProductShape)
		for _, shape := range diagram.Product_Shapes {
			if shape.Product != nil {
				diagram.map_Product_ProductShape[shape.Product] = shape
				diagrams := stager.map_Element_Diagrams[shape.Product]

				if diagrams == nil {
					diagrams = []DiagramIF{diagram}
				}

				if !slices.Contains(diagrams, DiagramIF(diagram)) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Product] = diagrams
			}
		}

		diagram.map_Task_TaskShape = make(map[*Task]*TaskShape)
		for _, shape := range diagram.Task_Shapes {
			if shape.Task != nil {
				diagram.map_Task_TaskShape[shape.Task] = shape
				diagrams := stager.map_Element_Diagrams[shape.Task]

				if diagrams == nil {
					diagrams = []DiagramIF{diagram}
				}

				if !slices.Contains(diagrams, DiagramIF(diagram)) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Task] = diagrams
			}
		}

		diagram.map_Task_TaskInputShape = make(map[taskProductKey]*TaskInputShape)
		for _, shape := range diagram.TaskInputShapes {
			if shape.Task != nil && shape.Product != nil {
				diagram.map_Task_TaskInputShape[taskProductKey{Task: shape.Task, Product: shape.Product}] = shape
			}
		}

		diagram.map_Task_TaskOutputShape = make(map[taskProductKey]*TaskOutputShape)
		for _, shape := range diagram.TaskOutputShapes {
			if shape.Task != nil && shape.Product != nil {
				diagram.map_Task_TaskOutputShape[taskProductKey{Task: shape.Task, Product: shape.Product}] = shape
			}
		}

		diagram.map_Product_ProductCompositionShape = make(map[*Product]*ProductCompositionShape)
		for _, shape := range diagram.ProductComposition_Shapes {
			if shape.Product != nil {
				diagram.map_Product_ProductCompositionShape[shape.Product] = shape
			}
		}

		diagram.map_Task_TaskCompositionShape = make(map[*Task]*TaskCompositionShape)
		for _, shape := range diagram.TaskComposition_Shapes {
			if shape.Task != nil {
				diagram.map_Task_TaskCompositionShape[shape.Task] = shape
			}
		}

		diagram.map_Note_NoteShape = make(map[*Note]*NoteShape)
		for _, shape := range diagram.Note_Shapes {
			if shape.Note != nil {
				diagram.map_Note_NoteShape[shape.Note] = shape
			}
		}

		diagram.map_Note_NoteProductShape = make(map[noteProductKey]*NoteProductShape)
		for _, shape := range diagram.NoteProductShapes {
			if shape.Note != nil {
				diagram.map_Note_NoteProductShape[noteProductKey{Note: shape.Note, Product: shape.Product}] = shape
			}
		}

		diagram.map_Note_NoteTaskShape = make(map[noteTaskKey]*NoteTaskShape)
		for _, shape := range diagram.NoteTaskShapes {
			if shape.Note != nil && shape.Task != nil {
				diagram.map_Note_NoteTaskShape[noteTaskKey{Note: shape.Note, Task: shape.Task}] = shape
			}
		}

		diagram.map_Note_NoteResourceShape = make(map[noteResourceKey]*NoteResourceShape)
		for _, shape := range diagram.NoteResourceShapes {
			if shape.Note != nil && shape.Resource != nil {
				diagram.map_Note_NoteResourceShape[noteResourceKey{Note: shape.Note, Resource: shape.Resource}] = shape
			}
		}

		diagram.map_Resource_ResourceShape = make(map[*Resource]*ResourceShape)
		for _, shape := range diagram.Resource_Shapes {
			if shape.Resource != nil {
				diagram.map_Resource_ResourceShape[shape.Resource] = shape
				diagrams := stager.map_Element_Diagrams[shape.Resource]

				if diagrams == nil {
					diagrams = []DiagramIF{diagram}
				}

				if !slices.Contains(diagrams, DiagramIF(diagram)) {
					diagrams = append(diagrams, diagram)
				}
				stager.map_Element_Diagrams[shape.Resource] = diagrams
			}
		}

		diagram.map_Resource_ResourceCompositionShape = make(map[*Resource]*ResourceCompositionShape)
		for _, shape := range diagram.ResourceComposition_Shapes {
			if shape.Resource != nil {
				diagram.map_Resource_ResourceCompositionShape[shape.Resource] = shape
			}
		}

		diagram.map_Resource_ResourceTaskShape = make(map[resourceTaskKey]*ResourceTaskShape)
		for _, shape := range diagram.ResourceTaskShapes {
			if shape.Resource != nil && shape.Task != nil {
				diagram.map_Resource_ResourceTaskShape[resourceTaskKey{Resource: shape.Resource, Task: shape.Task}] = shape
			}
		}
	}
}
