package models

import (
	"slices"
)

func (stager *Stager) enforceStagerMaps() {
	stager.map_Element_Diagrams = make(map[AbstractType][]*DiagramStructure)

	for _, diagramStructure := range stager.stage.GetInstancesSorted[*DiagramStructure]() {
		_ = diagramStructure

		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.System_Shapes, &diagramStructure.map_System_SystemShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.Part_Shapes, &diagramStructure.map_Part_PartShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.ExternalPart_Shapes, &diagramStructure.map_Part_ExternalPartShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.Port_Shapes, &diagramStructure.map_Port_PortShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.ControlFlow_Shapes, &diagramStructure.map_ControlFlow_ControlFlowShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.DataFlow_Shapes, &diagramStructure.map_DataFlow_DataFlowShape)
		updateMapElementDiagrams(stager, diagramStructure, diagramStructure.Note_Shapes, &diagramStructure.map_Note_NoteShape)

		// Scoped PartShape lookup (prioritizing top-level shapes on map_Part_PartShape)
		diagramStructure.map_PartShapeKey_PartShape = make(map[partShapeKey]*PartShape)
		for _, partShape := range diagramStructure.Part_Shapes {
			if partShape == nil || partShape.Part == nil {
				continue
			}
			key := partShapeKey{
				part:            partShape.Part,
				parentPartShape: partShape.ParentPartShape,
			}
			diagramStructure.map_PartShapeKey_PartShape[key] = partShape
			if partShape.ParentPartShape == nil {
				diagramStructure.map_Part_PartShape[partShape.Part] = partShape
			}
		}

		// Scoped PortShape lookup
		diagramStructure.map_PortShapeKey_PortShape = make(map[portShapeKey]*PortShape)
		for _, portShape := range diagramStructure.Port_Shapes {
			if portShape == nil || portShape.Port == nil {
				continue
			}
			parentPS := effectivePortParentPartShape(diagramStructure, portShape)
			key := portShapeKey{
				port:            portShape.Port,
				parentPartShape: parentPS,
			}
			diagramStructure.map_PortShapeKey_PortShape[key] = portShape
			if parentPS != nil && parentPS.ParentPartShape == nil {
				diagramStructure.map_Port_PortShape[portShape.Port] = portShape
			}
		}

		// Scoped ControlFlowShape lookup
		diagramStructure.map_ControlFlowShapeKey_ControlFlowShape = make(map[controlFlowShapeKey]*ControlFlowShape)
		for _, cfShape := range diagramStructure.ControlFlow_Shapes {
			if cfShape == nil || cfShape.ControlFlow == nil {
				continue
			}
			parentPS := effectiveControlFlowParentPartShape(diagramStructure, cfShape)
			key := controlFlowShapeKey{
				controlFlow:     cfShape.ControlFlow,
				parentPartShape: parentPS,
			}
			diagramStructure.map_ControlFlowShapeKey_ControlFlowShape[key] = cfShape
			if parentPS != nil && parentPS.ParentPartShape == nil {
				diagramStructure.map_ControlFlow_ControlFlowShape[cfShape.ControlFlow] = cfShape
			}
		}

		// Scoped DataFlowShape lookup
		diagramStructure.map_DataFlowShapeKey_DataFlowShape = make(map[dataFlowShapeKey]*DataFlowShape)
		for _, dfShape := range diagramStructure.DataFlow_Shapes {
			if dfShape == nil || dfShape.DataFlow == nil {
				continue
			}
			key := dataFlowShapeKey{
				dataFlow:        dfShape.DataFlow,
				parentPartShape: dfShape.ParentPartShape,
			}
			diagramStructure.map_DataFlowShapeKey_DataFlowShape[key] = dfShape
			if dfShape.ParentPartShape == nil {
				diagramStructure.map_DataFlow_DataFlowShape[dfShape.DataFlow] = dfShape
			}
		}

		diagramStructure.map_DataShapeKey_DataShape = make(map[dataShapeKey]*DataShape)
		for _, dataShape := range diagramStructure.Data_Shapes {
			key := dataShapeKey{
				dataFlow:      dataShape.DataFlow,
				dataFlowShape: effectiveDataFlowShape(diagramStructure, dataShape.DataFlow, dataShape.DataFlowShape),
				data:          dataShape.Data,
			}
			diagramStructure.map_DataShapeKey_DataShape[key] = dataShape
		}

		diagramStructure.map_ProtocolShapeKey_ProtocolShape = make(map[protocolShapeKey]*ProtocolShape)
		for _, protocolShape := range diagramStructure.Protocol_Shapes {
			key := protocolShapeKey{
				dataFlow:      protocolShape.DataFlow,
				dataFlowShape: effectiveDataFlowShape(diagramStructure, protocolShape.DataFlow, protocolShape.DataFlowShape),
				protocol:      protocolShape.Protocol,
			}
			diagramStructure.map_ProtocolShapeKey_ProtocolShape[key] = protocolShape
		}

		diagramStructure.map_ProtocolLayerShapeKey_ProtocolLayerShape = make(map[protocolLayerShapeKey]*ProtocolLayerShape)
		for _, protocolLayerShape := range diagramStructure.ProtocolLayer_Shapes {
			key := protocolLayerShapeKey{
				dataFlow:      protocolLayerShape.DataFlow,
				dataFlowShape: effectiveDataFlowShape(diagramStructure, protocolLayerShape.DataFlow, protocolLayerShape.DataFlowShape),
				protocolLayer: protocolLayerShape.ProtocolLayer,
			}
			diagramStructure.map_ProtocolLayerShapeKey_ProtocolLayerShape[key] = protocolLayerShape
		}

		diagramStructure.map_AllocatedResourceShapeKey_AllocatedResourceShape = make(map[allocatedResourceShapeKey]*AllocatedResourceShape)
		for _, allocatedResourceShape := range diagramStructure.AllocatedResourceShapes {
			key := allocatedResourceShapeKey{
				part:     allocatedResourceShape.Part,
				resource: allocatedResourceShape.Resource,
			}
			diagramStructure.map_AllocatedResourceShapeKey_AllocatedResourceShape[key] = allocatedResourceShape
		}

		diagramStructure.map_AllocatedSystemShapeKey_AllocatedSystemShape = make(map[allocatedSystemShapeKey]*AllocatedSystemShape)
		for _, allocatedSystemShape := range diagramStructure.AllocatedSystemShapes {
			key := allocatedSystemShapeKey{
				part:   allocatedSystemShape.Part,
				system: allocatedSystemShape.System,
			}
			diagramStructure.map_AllocatedSystemShapeKey_AllocatedSystemShape[key] = allocatedSystemShape
		}

		diagramStructure.map_Note_NotePortShape = make(map[notePortKey]*NotePortShape)
		for _, notePortShape := range diagramStructure.NotePortShapes {
			key := notePortKey{
				Note: notePortShape.Note,
				Port: notePortShape.Port,
			}
			diagramStructure.map_Note_NotePortShape[key] = notePortShape
		}

		diagramStructure.map_Note_NotePartShape = make(map[notePartKey]*NotePartShape)
		for _, notePartShape := range diagramStructure.NotePartShapes {
			key := notePartKey{
				Note: notePartShape.Note,
				Part: notePartShape.Part,
			}
			diagramStructure.map_Note_NotePartShape[key] = notePartShape
		}

		diagramStructure.map_Note_NoteDataFlowShape = make(map[noteDataFlowKey]*NoteDataFlowShape)
		for _, noteDataFlowShape := range diagramStructure.NoteDataFlowShapes {
			key := noteDataFlowKey{
				Note:     noteDataFlowShape.Note,
				DataFlow: noteDataFlowShape.DataFlow,
			}
			diagramStructure.map_Note_NoteDataFlowShape[key] = noteDataFlowShape
		}
	}

	stager.map_Note_PrevParts = make(map[*Note][]*Part)
	stager.map_Note_PrevPorts = make(map[*Note][]*Port)
	stager.map_Note_PrevDataFlows = make(map[*Note][]*DataFlow)
	for _, note := range stager.stage.GetInstancesSorted[*Note]() {
		stager.map_Note_PrevParts[note] = slices.Clone(note.Parts)
		stager.map_Note_PrevPorts[note] = slices.Clone(note.Ports)
		stager.map_Note_PrevDataFlows[note] = slices.Clone(note.DataFlows)
	}

	stager.rm_Data_DataFlows = stager.stage.GetSliceOfPointersReverseMap[DataFlow, Data](GongGetAssociationName[DataFlow]().Datas[0].Name)
	stager.rm_ProtocolLayer_Protocols = stager.stage.GetSliceOfPointersReverseMap[Protocol, ProtocolLayer](GongGetAssociationName[Protocol]().ProtocolLayers[0].Name)

	stager.rm_Protocol_DataFlows = make(map[*Protocol][]*DataFlow)
	for _, df := range stager.stage.GetInstancesSorted[*DataFlow]() {
		if df.Protocol != nil {
			stager.rm_Protocol_DataFlows[df.Protocol] = append(stager.rm_Protocol_DataFlows[df.Protocol], df)
		}
	}
	stager.rm_Protocol_Ports = make(map[*Protocol][]*Port)
	for _, port := range stager.stage.GetInstancesSorted[*Port]() {
		if port.Protocol != nil {
			stager.rm_Protocol_Ports[port.Protocol] = append(stager.rm_Protocol_Ports[port.Protocol], port)
		}
	}
}

// updateMapElementDiagrams is a helper function to update the map of abstract elements to their shapes for a given diagram
// and track the diagrams where each element is displayed.
func updateMapElementDiagrams[
	AT interface {
		AbstractType
		comparable
	},
	CT ConcreteType,
](
	stager *Stager,
	diagram *DiagramStructure,
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
			diagrams = []*DiagramStructure{diagram}
		}
		if !slices.Contains(diagrams, diagram) {
			diagrams = append(diagrams, diagram)
		}
		stager.map_Element_Diagrams[abstractElement] = diagrams
	}
}
