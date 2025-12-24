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
	}
}
