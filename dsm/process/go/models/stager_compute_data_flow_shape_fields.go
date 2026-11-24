package models

func (stager *Stager) computeDataFlowShapeFields() {
	stage := stager.stage

	// reset fields
	for _, dataFlowShape := range stage.GetInstancesSorted[*DataFlowShape]() {
		dataFlowShape.dataShapes = nil
	}

	// compute dataShapes by iterating over diagrams and their maps,
	// walking each DataFlow's Datas and their SubDatas in pre-order DFS.
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
	}
}
