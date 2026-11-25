package models

func (stager *Stager) computeDataFlowShapeFields() {
	stage := stager.stage

	// reset fields
	for _, dataFlowShape := range stage.GetInstancesSorted[*DataFlowShape]() {
		dataFlowShape.dataShapes = nil
	}
	for _, controlFlowShape := range stage.GetInstancesSorted[*ControlFlowShape]() {
		controlFlowShape.dataShapes = nil
	}

	// compute dataShapes by iterating over diagrams and their maps,
	// walking each DataFlow's and ControlFlow's Datas and their SubDatas in pre-order DFS.
	for _, diagram := range stage.GetInstancesSorted[*DiagramProcess]() {
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			if dataFlowShape.DataFlow != nil {
				df := dataFlowShape.DataFlow
				walkDatasPreOrder(df.Datas, func(data *Data) {
					key := dataShapeKey{
						dataFlow: df,
						data:     data,
					}
					if dataShape, ok := diagram.map_DataShapeKey_DataShape[key]; ok {
						dataFlowShape.dataShapes = append(dataFlowShape.dataShapes, dataShape)
					}
				})
			}
		}
		for _, controlFlowShape := range diagram.ControlFlow_Shapes {
			if controlFlowShape.ControlFlow != nil {
				cf := controlFlowShape.ControlFlow
				walkDatasPreOrder(cf.Datas, func(data *Data) {
					key := dataShapeKey{
						controlFlow: cf,
						data:        data,
					}
					if dataShape, ok := diagram.map_DataShapeKey_DataShape[key]; ok {
						controlFlowShape.dataShapes = append(controlFlowShape.dataShapes, dataShape)
					}
				})
			}
		}
	}
}
