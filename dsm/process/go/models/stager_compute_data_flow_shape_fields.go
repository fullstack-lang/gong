package models

func (stager *Stager) computeDataFlowShapeFields() {
	stage := stager.stage

	// reset fields
	for _, dataFlowShape := range GetGongstrucsSorted[*DataFlowShape](stage) {
		dataFlowShape.dataShapes = nil
	}

	// compute dataShapes by iterating over diagrams and their maps
	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stage) {
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			if dataFlowShape.DataFlow != nil {
				for _, data := range dataFlowShape.DataFlow.Datas {
					if dataShape, ok := diagram.map_Data_DataShape[data]; ok {
						dataFlowShape.dataShapes = append(dataFlowShape.dataShapes, dataShape)
					}
				}
			}
		}
	}
}
