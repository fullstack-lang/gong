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
		ASSOCIATION_TYPE_TASK_TASK                 associationType = "TaskTask"
		ASSOCIATION_TYPE_EXTERNAL_PARTICIPANT_TASK associationType = "ExternalParticipantTask"
		ASSOCIATION_TYPE_TASK_EXTERNAL_PARTICIPANT associationType = "TaskExternalParticipant"
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
	if startTaskShape, ok := diagramProcess.map_SvgRect_TaskShape[startRect]; ok {
		if endExternalParticipantShape, ok := diagramProcess.map_SvgRect_ExternalParticipantShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_EXTERNAL_PARTICIPANT
			sourceAbstratctElement = startTaskShape.Task
			targetAbstractElement = endExternalParticipantShape.Participant
		}
	}
	if startExternalParticipantShape, ok := diagramProcess.map_SvgRect_ExternalParticipantShape[startRect]; ok {
		if endTaskShape, ok := diagramProcess.map_SvgRect_TaskShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_EXTERNAL_PARTICIPANT_TASK
			sourceAbstratctElement = startExternalParticipantShape.Participant
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
				Type:      DataFlow_Task2Task,
				StartTask: startTask,
				EndTask:   endTask,
			})
			dataFlow.StartTask.owningParticipant.owningProcess.DataFlows = append(dataFlow.StartTask.owningParticipant.owningProcess.DataFlows, dataFlow)
			dataFlow.Name = dataFlow.StartTask.GetName() + " to " + dataFlow.EndTask.GetName()
			dataFlow.Stage(stage)

			dataFlowShape := (&DataFlowShape{
				DataFlow: dataFlow,
			}).Stage(stager.stage)

			dataFlowShape.SetName(dataFlow.StartTask.GetName() + " to " + dataFlow.EndTask.GetName())
			dataFlowShape.SetAbstractStartElement(dataFlow.StartTask)
			dataFlowShape.SetAbstractEndElement(dataFlow.EndTask)
			dataFlowShape.SetStartOrientation(ORIENTATION_VERTICAL)
			dataFlowShape.SetEndOrientation(ORIENTATION_VERTICAL)

			dataFlowShape.SetCornerOffsetRatio(1.5)
			dataFlowShape.SetStartRatio(0.5)
			dataFlowShape.SetEndRatio(0.5)
			diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)
		}
	case ASSOCIATION_TYPE_EXTERNAL_PARTICIPANT_TASK:
		startParticipant := sourceAbstratctElement.(*Participant)
		endTask := targetAbstractElement.(*Task)

		dataFlow := (&DataFlow{
			Type:                     DataFlow_ExternalParticipant2Task,
			StartExternalParticipant: startParticipant,
			EndTask:                  endTask,
		})
		dataFlow.StartExternalParticipant.owningProcess.DataFlows = append(dataFlow.StartExternalParticipant.owningProcess.DataFlows, dataFlow)
		dataFlow.Name = dataFlow.StartExternalParticipant.GetName() + " to " + dataFlow.EndTask.GetName()
		dataFlow.Stage(stage)

		dataFlowShape := (&DataFlowShape{
			DataFlow: dataFlow,
		}).Stage(stager.stage)

		dataFlowShape.SetName(dataFlow.StartExternalParticipant.GetName() + " to " + dataFlow.EndTask.GetName())
		dataFlowShape.SetAbstractStartElement(dataFlow.StartExternalParticipant)
		dataFlowShape.SetAbstractEndElement(dataFlow.EndTask)
		dataFlowShape.SetStartOrientation(ORIENTATION_VERTICAL)
		dataFlowShape.SetEndOrientation(ORIENTATION_VERTICAL)

		dataFlowShape.SetCornerOffsetRatio(1.5)
		dataFlowShape.SetStartRatio(0.5)
		dataFlowShape.SetEndRatio(0.5)
		diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)
	case ASSOCIATION_TYPE_TASK_EXTERNAL_PARTICIPANT:
		startTask := sourceAbstratctElement.(*Task)
		endParticipant := targetAbstractElement.(*Participant)

		dataFlow := (&DataFlow{
			Type:                   DataFlow_ExternalParticipant2Task,
			StartTask:              startTask,
			EndExternalParticipant: endParticipant,
		})
		dataFlow.StartTask.owningParticipant.owningProcess.DataFlows = append(dataFlow.StartTask.owningParticipant.owningProcess.DataFlows, dataFlow)
		dataFlow.Name = dataFlow.StartTask.GetName() + " to " + dataFlow.EndExternalParticipant.GetName()
		dataFlow.Stage(stage)

		dataFlowShape := (&DataFlowShape{
			DataFlow: dataFlow,
		}).Stage(stager.stage)

		dataFlowShape.SetName(dataFlow.StartTask.GetName() + " to " + dataFlow.EndExternalParticipant.GetName())
		dataFlowShape.SetAbstractStartElement(dataFlow.StartTask)
		dataFlowShape.SetAbstractEndElement(dataFlow.EndExternalParticipant)
		dataFlowShape.SetStartOrientation(ORIENTATION_VERTICAL)
		dataFlowShape.SetEndOrientation(ORIENTATION_VERTICAL)

		dataFlowShape.SetCornerOffsetRatio(1.5)
		dataFlowShape.SetStartRatio(0.5)
		dataFlowShape.SetEndRatio(0.5)
		diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
