package models

import (
	"slices"

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
		ASSOCIATION_TYPE_PORT_PORT          associationType = "PortPort"
		ASSOCIATION_TYPE_EXTERNAL_PART_PORT associationType = "ExternalPartPort"
		ASSOCIATION_TYPE_PORT_EXTERNAL_PART associationType = "PortExternalPart"
		ASSOCIATION_TYPE_NOTE_PART          associationType = "NotePart"
		ASSOCIATION_TYPE_NOTE_PORT          associationType = "NotePort"
	)

	var assocType associationType
	var sourceAbstratctElement, targetAbstractElement AbstractType
	var matchedStartPortShape, matchedEndPortShape *PortShape
	startRect := frontSVG.StartRect
	endRect := frontSVG.EndRect

	if startPortShape, ok := diagramStructure.map_SvgRect_PortShape[startRect]; ok {
		if endPortShape, ok := diagramStructure.map_SvgRect_PortShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_PORT_PORT
			sourceAbstratctElement = startPortShape.Port
			targetAbstractElement = endPortShape.Port
			matchedStartPortShape = startPortShape
			matchedEndPortShape = endPortShape
		}
	}
	if startPortShape, ok := diagramStructure.map_SvgRect_PortShape[startRect]; ok {
		if endExternalPartShape, ok := diagramStructure.map_SvgRect_ExternalPartShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_PORT_EXTERNAL_PART
			sourceAbstratctElement = startPortShape.Port
			targetAbstractElement = endExternalPartShape.Part
			matchedStartPortShape = startPortShape
		}
	}
	if startExternalPartShape, ok := diagramStructure.map_SvgRect_ExternalPartShape[startRect]; ok {
		if endPortShape, ok := diagramStructure.map_SvgRect_PortShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_EXTERNAL_PART_PORT
			sourceAbstratctElement = startExternalPartShape.Part
			targetAbstractElement = endPortShape.Port
			matchedEndPortShape = endPortShape
		}
	}
	if startNoteShape, ok := diagramStructure.map_SvgRect_NoteShape[startRect]; ok {
		if endPart, ok := diagramStructure.map_SvgRect_Part[endRect]; ok {
			assocType = ASSOCIATION_TYPE_NOTE_PART
			sourceAbstratctElement = startNoteShape.Note
			targetAbstractElement = endPart
		}
	}
	if startNoteShape, ok := diagramStructure.map_SvgRect_NoteShape[startRect]; ok {
		if endPortShape, ok := diagramStructure.map_SvgRect_PortShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_NOTE_PORT
			sourceAbstratctElement = startNoteShape.Note
			targetAbstractElement = endPortShape.Port
		}
	}

	// create the association according to the association type
	switch assocType {
	case ASSOCIATION_TYPE_PORT_PORT:
		endPort := targetAbstractElement.(*Port)
		startPort := sourceAbstratctElement.(*Port)

		startParentPS := effectivePortParentPartShape(diagramStructure, matchedStartPortShape)
		endParentPS := effectivePortParentPartShape(diagramStructure, matchedEndPortShape)

		if startPort.owningPart == endPort.owningPart && startParentPS != nil && startParentPS == endParentPS {
			controlFlow := (&ControlFlow{
				Start: startPort,
				End:   endPort,
			})
			startPort.owningPart.ControlFlows = append(startPort.owningPart.ControlFlows, controlFlow)

			controlFlow.Name = controlFlow.Start.GetName() + " to " + controlFlow.End.GetName()
			controlFlow.Stage(stage)

			controlFlowShape := (&ControlFlowShape{
				ControlFlow:     controlFlow,
				ParentPartShape: startParentPS,
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
			owningSystem := diagramStructure.owningSystem
			var enclosingPartShape *PartShape
			if startParentPS != nil && endParentPS != nil && startParentPS.ParentPartShape == endParentPS.ParentPartShape {
				if startParentPS.ParentPartShape != nil &&
					startParentPS.ParentPartShape.Part != nil &&
					startParentPS.ParentPartShape.Part.Type != nil {
					owningSystem = startParentPS.ParentPartShape.Part.Type
					enclosingPartShape = startParentPS.ParentPartShape
				} else if startPort.owningPart != nil && startPort.owningPart.owningSystem != nil {
					owningSystem = startPort.owningPart.owningSystem
				}
			}
			if owningSystem != nil {
				owningSystem.DataFlows = append(owningSystem.DataFlows, dataFlow)
			}
			lib := stager.findSystemLibrary(owningSystem)
			dataFlow.SetOwningLibrary(lib)
			if lib != nil && !slices.Contains(lib.RootDataFlows, dataFlow) {
				lib.RootDataFlows = append(lib.RootDataFlows, dataFlow)
			}
			dataFlow.Name = dataFlow.StartPort.GetName() + " to " + dataFlow.EndPort.GetName()
			dataFlow.Stage(stage)

			dataFlowShape := (&DataFlowShape{
				DataFlow:        dataFlow,
				ParentPartShape: enclosingPartShape,
				StartPortShape:  matchedStartPortShape,
				EndPortShape:    matchedEndPortShape,
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
			Type:              DataFlow_ExternalPart2Port,
			StartExternalPart: startPart,
			EndPort:           endPort,
		})
		owningSystem := dataFlow.StartExternalPart.owningSystem
		owningSystem.DataFlows = append(owningSystem.DataFlows, dataFlow)
		lib := stager.findSystemLibrary(owningSystem)
		dataFlow.SetOwningLibrary(lib)
		if lib != nil && !slices.Contains(lib.RootDataFlows, dataFlow) {
			lib.RootDataFlows = append(lib.RootDataFlows, dataFlow)
		}
		dataFlow.Name = dataFlow.StartExternalPart.GetName() + " to " + dataFlow.EndPort.GetName()
		dataFlow.Stage(stage)

		dataFlowShape := (&DataFlowShape{
			DataFlow:     dataFlow,
			EndPortShape: matchedEndPortShape,
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
			Type:            DataFlow_Port2ExternalPart,
			StartPort:       startPort,
			EndExternalPart: endPart,
		})
		owningSystem := diagramStructure.owningSystem
		if owningSystem == nil && dataFlow.EndExternalPart != nil {
			owningSystem = dataFlow.EndExternalPart.owningSystem
		}
		if owningSystem != nil {
			owningSystem.DataFlows = append(owningSystem.DataFlows, dataFlow)
		}
		lib := stager.findSystemLibrary(owningSystem)
		dataFlow.SetOwningLibrary(lib)
		if lib != nil && !slices.Contains(lib.RootDataFlows, dataFlow) {
			lib.RootDataFlows = append(lib.RootDataFlows, dataFlow)
		}
		dataFlow.Name = dataFlow.StartPort.GetName() + " to " + dataFlow.EndExternalPart.GetName()
		dataFlow.Stage(stage)

		dataFlowShape := (&DataFlowShape{
			DataFlow:       dataFlow,
			StartPortShape: matchedStartPortShape,
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
	case ASSOCIATION_TYPE_NOTE_PART:
		note := sourceAbstratctElement.(*Note)
		part := targetAbstractElement.(*Part)

		if !slices.Contains(note.Parts, part) {
			note.Parts = append(note.Parts, part)
		}
		if _, exists := diagramStructure.map_Note_NotePartShape[notePartKey{Note: note, Part: part}]; !exists {
			addAssociationShapeToDiagram(stager, note, part, &diagramStructure.NotePartShapes)
		}
	case ASSOCIATION_TYPE_NOTE_PORT:
		note := sourceAbstratctElement.(*Note)
		port := targetAbstractElement.(*Port)

		if !slices.Contains(note.Ports, port) {
			note.Ports = append(note.Ports, port)
		}
		if _, exists := diagramStructure.map_Note_NotePortShape[notePortKey{Note: note, Port: port}]; !exists {
			addAssociationShapeToDiagram(stager, note, port, &diagramStructure.NotePortShapes)
		}
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
