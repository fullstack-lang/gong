package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// SVGUpdated implements SVGImplInterface.
func (stager *Stager) onUpdateSVG(frontSVG *svg.SVG) {
	stage := stager.stage
	diagramStructure := stager.diagramStructure
	svgObjectDiagramStructure := stager.svgObjectDiagramStructure

	if svgObjectDiagramStructure.DrawingState == frontSVG.DrawingState {
		// in any cases, have the form editor set up with the instance
		stager.probeForm.FillUpFormFromGongstruct(diagramStructure, "Diagram")
		return
	}

	// IMPORTANT : we are only interested when the updateSVG has finished drawing the connexion
	if frontSVG.DrawingState == svg.DRAWING_LINK {
		return
	}

	// DSM specific association
	type associationType string
	const (
		ASSOCIATION_TYPE_PORT_PORT                 associationType = "PortPort"
		ASSOCIATION_TYPE_EXTERNAL_PART_PORT associationType = "ExternalPartPort"
		ASSOCIATION_TYPE_PORT_EXTERNAL_PART associationType = "PortExternalPart"
	)

	var assocType associationType
	var sourceAbstratctElement, targetAbstractElement AbstractType
	startRect := frontSVG.StartRect
	endRect := frontSVG.EndRect

	if startPortShape, ok := diagramStructure.map_SvgRect_PortShape[startRect]; ok {
		if endPortShape, ok := diagramStructure.map_SvgRect_PortShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_PORT_PORT
			sourceAbstratctElement = startPortShape.Port
			targetAbstractElement = endPortShape.Port
		}
	}
	if startPortShape, ok := diagramStructure.map_SvgRect_PortShape[startRect]; ok {
		if endExternalPartShape, ok := diagramStructure.map_SvgRect_ExternalPartShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_PORT_EXTERNAL_PART
			sourceAbstratctElement = startPortShape.Port
			targetAbstractElement = endExternalPartShape.Part
		}
	}
	if startExternalPartShape, ok := diagramStructure.map_SvgRect_ExternalPartShape[startRect]; ok {
		if endPortShape, ok := diagramStructure.map_SvgRect_PortShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_EXTERNAL_PART_PORT
			sourceAbstratctElement = startExternalPartShape.Part
			targetAbstractElement = endPortShape.Port
		}
	}

	// create the association according to the association type
	switch assocType {
	case ASSOCIATION_TYPE_PORT_PORT:
		endPort := targetAbstractElement.(*Port)
		startPort := sourceAbstratctElement.(*Port)

		if startPort.owningPart == endPort.owningPart {
			controlFlow := (&ControlFlow{
				Start: startPort,
				End:   endPort,
			})
			startPort.owningPart.ControlFlows = append(startPort.owningPart.ControlFlows, controlFlow)

			controlFlow.Name = controlFlow.Start.GetName() + " to " + controlFlow.End.GetName()
			controlFlow.Stage(stage)

			controlFlowShape := (&ControlFlowShape{
				ControlFlow: controlFlow,
			}).Stage(stager.stage)

			controlFlowShape.SetName(controlFlow.Start.GetName() + " to " + controlFlow.End.GetName())
			controlFlowShape.SetAbstractStartElement(controlFlow.Start)
			controlFlowShape.SetAbstractEndElement(controlFlow.End)
			controlFlowShape.SetStartOrientation(ORIENTATION_VERTICAL)
			controlFlowShape.SetEndOrientation(ORIENTATION_VERTICAL)

			controlFlowShape.SetCornerOffsetRatio(1.5)
			controlFlowShape.SetStartRatio(0.5)
			controlFlowShape.SetEndRatio(0.5)
			diagramStructure.ControlFlow_Shapes = append(diagramStructure.ControlFlow_Shapes, controlFlowShape)
		} else {
			dataFlow := (&DataFlow{
				Type:      DataFlow_Port2Port,
				StartPort: startPort,
				EndPort:   endPort,
			})
			dataFlow.StartPort.owningPart.owningSystem.DataFlows = append(dataFlow.StartPort.owningPart.owningSystem.DataFlows, dataFlow)
			dataFlow.Name = dataFlow.StartPort.GetName() + " to " + dataFlow.EndPort.GetName()
			dataFlow.Stage(stage)

			dataFlowShape := (&DataFlowShape{
				DataFlow: dataFlow,
			}).Stage(stager.stage)

			dataFlowShape.SetName(dataFlow.StartPort.GetName() + " to " + dataFlow.EndPort.GetName())
			dataFlowShape.SetAbstractStartElement(dataFlow.StartPort)
			dataFlowShape.SetAbstractEndElement(dataFlow.EndPort)
			dataFlowShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
			dataFlowShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

			dataFlowShape.SetCornerOffsetRatio(1.5)
			dataFlowShape.SetStartRatio(0.5)
			dataFlowShape.SetEndRatio(0.5)
			diagramStructure.DataFlow_Shapes = append(diagramStructure.DataFlow_Shapes, dataFlowShape)
		}
	case ASSOCIATION_TYPE_EXTERNAL_PART_PORT:
		startPart := sourceAbstratctElement.(*Part)
		endPort := targetAbstractElement.(*Port)

		dataFlow := (&DataFlow{
			Type:                     DataFlow_ExternalPart2Port,
			StartExternalPart: startPart,
			EndPort:                  endPort,
		})
		dataFlow.StartExternalPart.owningSystem.DataFlows = append(dataFlow.StartExternalPart.owningSystem.DataFlows, dataFlow)
		dataFlow.Name = dataFlow.StartExternalPart.GetName() + " to " + dataFlow.EndPort.GetName()
		dataFlow.Stage(stage)

		dataFlowShape := (&DataFlowShape{
			DataFlow: dataFlow,
		}).Stage(stager.stage)

		dataFlowShape.SetName(dataFlow.StartExternalPart.GetName() + " to " + dataFlow.EndPort.GetName())
		dataFlowShape.SetAbstractStartElement(dataFlow.StartExternalPart)
		dataFlowShape.SetAbstractEndElement(dataFlow.EndPort)
		dataFlowShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
		dataFlowShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

		dataFlowShape.SetCornerOffsetRatio(1.5)
		dataFlowShape.SetStartRatio(0.5)
		dataFlowShape.SetEndRatio(0.5)
		diagramStructure.DataFlow_Shapes = append(diagramStructure.DataFlow_Shapes, dataFlowShape)
	case ASSOCIATION_TYPE_PORT_EXTERNAL_PART:
		startPort := sourceAbstratctElement.(*Port)
		endPart := targetAbstractElement.(*Part)

		dataFlow := (&DataFlow{
			Type:                   DataFlow_ExternalPart2Port,
			StartPort:              startPort,
			EndExternalPart: endPart,
		})
		dataFlow.StartPort.owningPart.owningSystem.DataFlows = append(dataFlow.StartPort.owningPart.owningSystem.DataFlows, dataFlow)
		dataFlow.Name = dataFlow.StartPort.GetName() + " to " + dataFlow.EndExternalPart.GetName()
		dataFlow.Stage(stage)

		dataFlowShape := (&DataFlowShape{
			DataFlow: dataFlow,
		}).Stage(stager.stage)

		dataFlowShape.SetName(dataFlow.StartPort.GetName() + " to " + dataFlow.EndExternalPart.GetName())
		dataFlowShape.SetAbstractStartElement(dataFlow.StartPort)
		dataFlowShape.SetAbstractEndElement(dataFlow.EndExternalPart)
		dataFlowShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
		dataFlowShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

		dataFlowShape.SetCornerOffsetRatio(1.5)
		dataFlowShape.SetStartRatio(0.5)
		dataFlowShape.SetEndRatio(0.5)
		diagramStructure.DataFlow_Shapes = append(diagramStructure.DataFlow_Shapes, dataFlowShape)
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
