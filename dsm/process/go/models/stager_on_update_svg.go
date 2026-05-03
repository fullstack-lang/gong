package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// SVGUpdated implements SVGImplInterface.
func (stager *Stager) onUpdateSVG(frontSVG *svg.SVG) {
	stage := stager.stage
	diagramProcess := stager.diagramProcess
	svgObjectDiagramProcess := stager.svgObjectDiagramProcess

	if svgObjectDiagramProcess.DrawingState == frontSVG.DrawingState {
		// in any cases, have the form editor set up with the instance
		stager.probeForm.FillUpFormFromGongstruct(diagramProcess, "Diagram")
		return
	}

	// IMPORTANT : we are only interested when the updateSVG has finished drawing the connexion
	if frontSVG.DrawingState == svg.DRAWING_LINK {
		return
	}

	// DSM specific association
	type associationType string
	const (
		ASSOCIATION_TYPE_TASK_TASK associationType = "TaskTask"
	)

	var assocType associationType
	var sourceAbstratctElement, targetAbstractElement AbstractType
	startRect := frontSVG.StartRect
	endRect := frontSVG.EndRect

	if startTaskShape, ok := diagramProcess.map_SvgRect_TaskShape[startRect]; ok {
		if endTaskShape, ok := diagramProcess.map_SvgRect_TaskShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_TASK
			sourceAbstratctElement = startTaskShape.Task
			targetAbstractElement = endTaskShape.Task
		}
	}

	// create the association according to the association type
	switch assocType {
	case ASSOCIATION_TYPE_TASK_TASK:
		endTask := targetAbstractElement.(*Task)
		startTask := sourceAbstratctElement.(*Task)

		if startTask.owningParticipant == endTask.owningParticipant {
			controlFlow := (&ControlFlow{
				Start: startTask,
				End:   endTask,
			})
			startTask.owningParticipant.ControlFlows = append(startTask.owningParticipant.ControlFlows, controlFlow)

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
			diagramProcess.ControlFlow_Shapes = append(diagramProcess.ControlFlow_Shapes, controlFlowShape)
		} else {
			dataFlow := (&DataFlow{
				Start: startTask,
				End:   endTask,
			})
			dataFlow.Start.owningParticipant.owningProcess.DataFlows = append(dataFlow.Start.owningParticipant.owningProcess.DataFlows, dataFlow)
			dataFlow.Name = dataFlow.Start.GetName() + " to " + dataFlow.End.GetName()
			dataFlow.Stage(stage)

			dataFlowShape := (&DataFlowShape{
				DataFlow: dataFlow,
			}).Stage(stager.stage)

			dataFlowShape.SetName(dataFlow.Start.GetName() + " to " + dataFlow.End.GetName())
			dataFlowShape.SetAbstractStartElement(dataFlow.Start)
			dataFlowShape.SetAbstractEndElement(dataFlow.End)
			dataFlowShape.SetStartOrientation(ORIENTATION_VERTICAL)
			dataFlowShape.SetEndOrientation(ORIENTATION_VERTICAL)

			dataFlowShape.SetCornerOffsetRatio(1.5)
			dataFlowShape.SetStartRatio(0.5)
			dataFlowShape.SetEndRatio(0.5)
			diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)
		}
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
