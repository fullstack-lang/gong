package models

func (stager *Stager) computeDataFlowShapeFields() {
	stage := stager.stage

	// reset fields
	for _, dataFlowShape := range stage.GetInstancesSorted[*DataFlowShape]() {
		dataFlowShape.dataShapes = nil
		dataFlowShape.protocolShape = nil
		dataFlowShape.protocolLayerShapes = nil
	}

	// compute dataShapes, protocolShape, and protocolLayerShapes by iterating over diagrams and their maps
	for _, diagram := range stage.GetInstancesSorted[*DiagramStructure]() {
		for _, dataFlowShape := range diagram.DataFlow_Shapes {
			df := dataFlowShape.DataFlow
			if df == nil {
				continue
			}
			for _, data := range df.Datas {
				key := dataShapeKey{
					dataFlow: df,
					data:     data,
				}
				if dataShape, ok := diagram.map_DataShapeKey_DataShape[key]; ok {
					dataFlowShape.dataShapes = append(dataFlowShape.dataShapes, dataShape)
				}
			}
			if df.Protocol != nil {
				protoKey := protocolShapeKey{
					dataFlow: df,
					protocol: df.Protocol,
				}
				if protocolShape, ok := diagram.map_ProtocolShapeKey_ProtocolShape[protoKey]; ok {
					dataFlowShape.protocolShape = protocolShape
				}
				for _, layer := range df.Protocol.ProtocolLayers {
					layerKey := protocolLayerShapeKey{
						dataFlow:      df,
						protocolLayer: layer,
					}
					if layerShape, ok := diagram.map_ProtocolLayerShapeKey_ProtocolLayerShape[layerKey]; ok {
						dataFlowShape.protocolLayerShapes = append(dataFlowShape.protocolLayerShapes, layerShape)
					}
				}
			}
		}
	}
}
