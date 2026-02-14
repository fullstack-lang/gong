package models

func (stager *Stager) enforceDiagramMaps() {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		// computes all products presents in the diagram
		diagram.map_Product_ProductShape = make(map[*Product]*ProductShape)
		for _, shape := range diagram.Product_Shapes {
			if shape.Product != nil {
				diagram.map_Product_ProductShape[shape.Product] = shape
			}
		}

		diagram.map_Task_TaskShape = make(map[*Task]*TaskShape)
		for _, shape := range diagram.Task_Shapes {
			if shape.Task != nil {
				diagram.map_Task_TaskShape[shape.Task] = shape
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

		diagram.map_Resource_ResourceShape = make(map[*Resource]*ResourceShape)
		for _, shape := range diagram.Resource_Shapes {
			if shape.Resource != nil {
				diagram.map_Resource_ResourceShape[shape.Resource] = shape
			}
		}

		diagram.map_Resource_ResourceCompositionShape = make(map[*Resource]*ResourceCompositionShape)
		for _, shape := range diagram.ResourceComposition_Shapes {
			if shape.Resource != nil {
				diagram.map_Resource_ResourceCompositionShape[shape.Resource] = shape
			}
		}
	}
}
