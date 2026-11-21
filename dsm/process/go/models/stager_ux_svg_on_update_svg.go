package models

import (
	"math"
	"slices"

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
		ASSOCIATION_TYPE_NOTE_TASK                 associationType = "NoteTask"
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
			targetAbstractElement = endExternalParticipantShape.ExternalParticipant
		}
	}
	if startExternalParticipantShape, ok := diagramProcess.map_SvgRect_ExternalParticipantShape[startRect]; ok {
		if endTaskShape, ok := diagramProcess.map_SvgRect_TaskShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_EXTERNAL_PARTICIPANT_TASK
			sourceAbstratctElement = startExternalParticipantShape.ExternalParticipant
			targetAbstractElement = endTaskShape.Task
		}
	}
	if startNoteShape, ok := diagramProcess.map_SvgRect_NoteShape[startRect]; ok {
		if endTaskShape, ok := diagramProcess.map_SvgRect_TaskShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_NOTE_TASK
			sourceAbstratctElement = startNoteShape.Note
			targetAbstractElement = endTaskShape.Task
		}
	}
	if startTaskShape, ok := diagramProcess.map_SvgRect_TaskShape[startRect]; ok {
		if endNoteShape, ok := diagramProcess.map_SvgRect_NoteShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_NOTE_TASK
			sourceAbstratctElement = endNoteShape.Note
			targetAbstractElement = startTaskShape.Task
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

			configureDefaultControlFlowShape(diagramProcess, controlFlow, controlFlowShape, 1.5)
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
			dataFlowShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
			dataFlowShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

			cornerOffset := 1.5
			if startShape := diagramProcess.map_Task_TaskShape[startTask]; startTask.IsForkJoin && startShape != nil && startShape.IsVertical {
				cornerOffset = 4.5
			}
			dataFlowShape.SetCornerOffsetRatio(cornerOffset)
			dataFlowShape.SetStartRatio(0.5)
			dataFlowShape.SetEndRatio(0.5)
			diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)
		}
	case ASSOCIATION_TYPE_EXTERNAL_PARTICIPANT_TASK:
		startExternalParticipant := sourceAbstratctElement.(*ExternalParticipant)
		endTask := targetAbstractElement.(*Task)

		dataFlow := (&DataFlow{
			Type:             DataFlow_ExternalParticipant2Task,
			StartParticipant: startExternalParticipant,
			EndTask:          endTask,
		})
		dataFlow.StartParticipant.owningProcess.DataFlows = append(dataFlow.StartParticipant.owningProcess.DataFlows, dataFlow)
		dataFlow.Name = dataFlow.StartParticipant.GetName() + " to " + dataFlow.EndTask.GetName()
		dataFlow.Stage(stage)

		dataFlowShape := (&DataFlowShape{
			DataFlow: dataFlow,
		}).Stage(stager.stage)

		dataFlowShape.SetName(dataFlow.StartParticipant.GetName() + " to " + dataFlow.EndTask.GetName())
		dataFlowShape.SetAbstractStartElement(dataFlow.StartParticipant)
		dataFlowShape.SetAbstractEndElement(dataFlow.EndTask)
		dataFlowShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
		dataFlowShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

		dataFlowShape.SetCornerOffsetRatio(1.5)
		dataFlowShape.SetStartRatio(0.5)
		dataFlowShape.SetEndRatio(0.5)
		diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)
	case ASSOCIATION_TYPE_TASK_EXTERNAL_PARTICIPANT:
		startTask := sourceAbstratctElement.(*Task)
		endExternalParticipant := targetAbstractElement.(*ExternalParticipant)

		dataFlow := (&DataFlow{
			Type:           DataFlow_Task2ExternalParticipant,
			StartTask:      startTask,
			EndParticipant: endExternalParticipant,
		})
		dataFlow.StartTask.owningParticipant.owningProcess.DataFlows = append(dataFlow.StartTask.owningParticipant.owningProcess.DataFlows, dataFlow)
		dataFlow.Name = dataFlow.StartTask.GetName() + " to " + dataFlow.EndParticipant.GetName()
		dataFlow.Stage(stage)

		dataFlowShape := (&DataFlowShape{
			DataFlow: dataFlow,
		}).Stage(stager.stage)

		dataFlowShape.SetName(dataFlow.StartTask.GetName() + " to " + dataFlow.EndParticipant.GetName())
		dataFlowShape.SetAbstractStartElement(dataFlow.StartTask)
		dataFlowShape.SetAbstractEndElement(dataFlow.EndParticipant)
		dataFlowShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
		dataFlowShape.SetEndOrientation(ORIENTATION_HORIZONTAL)

		dataFlowShape.SetCornerOffsetRatio(1.5)
		dataFlowShape.SetStartRatio(0.5)
		dataFlowShape.SetEndRatio(0.5)
		diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)
	case ASSOCIATION_TYPE_NOTE_TASK:
		note := sourceAbstratctElement.(*Note)
		task := targetAbstractElement.(*Task)

		if !slices.Contains(note.Tasks, task) {
			note.Tasks = append(note.Tasks, task)
		}

		key := noteTaskKey{Note: note, Task: task}
		if _, ok := diagramProcess.map_Note_NoteTaskShape[key]; !ok {
			addAssociationShapeToDiagram(stager, note, task, &diagramProcess.NoteTaskShapes)
		}
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}

func clampForkJoinRatio(ratio float64) float64 {
	if ratio < 0.15 {
		return 0.15
	}
	if ratio > 0.85 {
		return 0.85
	}
	return ratio
}

func configureDefaultControlFlowShape(
	diagramProcess *DiagramProcess,
	controlFlow *ControlFlow,
	controlFlowShape *ControlFlowShape,
	defaultCornerOffset float64,
) {
	startTask := controlFlow.Start
	endTask := controlFlow.End

	controlFlowShape.SetName(startTask.GetName() + " to " + endTask.GetName())
	controlFlowShape.SetAbstractStartElement(startTask)
	controlFlowShape.SetAbstractEndElement(endTask)

	startShape := diagramProcess.map_Task_TaskShape[startTask]
	endShape := diagramProcess.map_Task_TaskShape[endTask]

	startOrientation := ORIENTATION_VERTICAL
	if startTask.IsForkJoin && startShape != nil && startShape.IsVertical {
		startOrientation = ORIENTATION_HORIZONTAL
	} else if startTask.IsDecisionNode && startShape != nil && endShape != nil {
		startCenterX := startShape.X + startShape.Width/2.0
		endCenterX := endShape.X + endShape.Width/2.0
		if math.Abs(endCenterX-startCenterX) > 60.0 {
			startOrientation = ORIENTATION_HORIZONTAL
		}
	}

	endOrientation := ORIENTATION_VERTICAL
	if endTask.IsForkJoin && endShape != nil && endShape.IsVertical {
		endOrientation = ORIENTATION_HORIZONTAL
	}

	startRatio := 0.5
	endRatio := 0.5

	if startTask.IsForkJoin && startShape != nil && endShape != nil {
		if startShape.IsVertical && startShape.Height > 0 {
			endCenterY := endShape.Y + endShape.Height/2.0
			startRatio = clampForkJoinRatio((endCenterY - startShape.Y) / startShape.Height)
		} else if !startShape.IsVertical && startShape.Width > 0 {
			endCenterX := endShape.X + endShape.Width/2.0
			startRatio = clampForkJoinRatio((endCenterX - startShape.X) / startShape.Width)
		}
	}

	if endTask.IsForkJoin && endShape != nil && startShape != nil {
		if endShape.IsVertical && endShape.Height > 0 {
			startCenterY := startShape.Y + startShape.Height/2.0
			endRatio = clampForkJoinRatio((startCenterY - endShape.Y) / endShape.Height)
		} else if !endShape.IsVertical && endShape.Width > 0 {
			startCenterX := startShape.X + startShape.Width/2.0
			endRatio = clampForkJoinRatio((startCenterX - endShape.X) / endShape.Width)
		}
	}

	cornerOffsetRatio := defaultCornerOffset
	if startTask.IsForkJoin {
		if startOrientation == ORIENTATION_VERTICAL && endOrientation == ORIENTATION_VERTICAL {
			cornerOffsetRatio = 4.5
		} else if startOrientation == ORIENTATION_HORIZONTAL && endOrientation == ORIENTATION_HORIZONTAL {
			cornerOffsetRatio = 4.5
		}
	}

	controlFlowShape.SetStartOrientation(startOrientation)
	controlFlowShape.SetEndOrientation(endOrientation)
	controlFlowShape.SetCornerOffsetRatio(cornerOffsetRatio)
	controlFlowShape.SetStartRatio(startRatio)
	controlFlowShape.SetEndRatio(endRatio)
}
