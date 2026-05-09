package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataFlowsWithinProcessDiagram(
	diagramProcess *DiagramProcess,
	dataFlow *DataFlow,
	parentNode *tree.Node,
) {
	stage := stager.stage

	// find the dataFlowShape (if any)
	dataFlowShape, isShapePresent := diagramProcess.map_DataFlow_DataFlowShape[dataFlow]

	isStartShapePresent := false
	isEndShapePresent := false

	var startName, endName string
	var startElement, endElement AbstractType

	switch dataFlow.Type {
	case DataFlow_Task2Task:
		if dataFlow.StartTask != nil {
			_, isStartShapePresent = diagramProcess.map_Task_TaskShape[dataFlow.StartTask]
			startName = dataFlow.StartTask.GetName()
			startElement = dataFlow.StartTask
		}
		if dataFlow.EndTask != nil {
			_, isEndShapePresent = diagramProcess.map_Task_TaskShape[dataFlow.EndTask]
			endName = dataFlow.EndTask.GetName()
			endElement = dataFlow.EndTask
		}
	case DataFlow_ExternalParticipant2Task:
		if dataFlow.StartExternalParticipant != nil {
			_, isStartShapePresent = diagramProcess.map_Participant_ExternalParticipantShape[dataFlow.StartExternalParticipant]
			startName = dataFlow.StartExternalParticipant.GetName()
			startElement = dataFlow.StartExternalParticipant
		}
		if dataFlow.EndTask != nil {
			_, isEndShapePresent = diagramProcess.map_Task_TaskShape[dataFlow.EndTask]
			endName = dataFlow.EndTask.GetName()
			endElement = dataFlow.EndTask
		}
	case DataFlow_Task2ExternalParticipant:
		if dataFlow.StartTask != nil {
			_, isStartShapePresent = diagramProcess.map_Task_TaskShape[dataFlow.StartTask]
			startName = dataFlow.StartTask.GetName()
			startElement = dataFlow.StartTask
		}
		if dataFlow.EndExternalParticipant != nil {
			_, isEndShapePresent = diagramProcess.map_Participant_ExternalParticipantShape[dataFlow.EndExternalParticipant]
			endName = dataFlow.EndExternalParticipant.GetName()
			endElement = dataFlow.EndExternalParticipant
		}
	}
	isCheckboxDisabled := !(isStartShapePresent && isEndShapePresent)

	dataFlowNode := &tree.Node{
		Name:                    dataFlow.GetName(),
		IsExpanded:              slices.Contains(diagramProcess.DataFlowsWhoseNodeIsExpanded, dataFlow),
		IsNodeClickable:         true,
		IsInEditMode:            dataFlow.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isShapePresent,
		IsCheckboxDisabled:      isCheckboxDisabled,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if isCheckboxDisabled {
				return "A data flow cannot be created if the start or end shape is not present from the diagram"
			}
			if isShapePresent {
				return "Click to remove the data flow shape"
			}
			return "Click to create a data flow shape for this data flow within this diagram"
		}(),
	}

	if isCheckboxDisabled {
		dataFlowNode.CheckboxHasToolTip = true
		dataFlowNode.CheckboxToolTipText = "Start or end shape is missing from the diagram"
	}
	parentNode.Children = append(parentNode.Children, dataFlowNode)

	addRenameButton(dataFlow, dataFlowNode, stager)

	if shape, ok := diagramProcess.map_DataFlow_DataFlowShape[dataFlow]; ok {
		dataFlowNode.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramProcess.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				shape.SetIsHidden(!shape.GetIsHidden())
				stage.Commit()
			},
		}
		if shape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		dataFlowNode.Buttons = append(dataFlowNode.Buttons, visibilityButton)
	}

	dataFlowNode.OnNameChange = stager.onNameChange(dataFlow)
	dataFlowNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataFlow, &diagramProcess.DataFlowsWhoseNodeIsExpanded)
	dataFlowNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isShapePresent {
			isShapePresent = isChecked
			if dataFlowShape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			// shape = newShapeToDiagram(controlflow, diagramProcess, &diagramProcess.ControlFlowShapes, stage)
			// addAssociationShapeToDiagram(stager, controlflow.Start, controlflow.End, &diagramProcess.ControlFlowShapes)
			dataFlowShape := (&DataFlowShape{
				DataFlow: dataFlow,
			}).Stage(stage)
			dataFlowShape.SetName(startName + " to " + endName)
			dataFlowShape.SetAbstractStartElement(startElement)
			dataFlowShape.SetAbstractEndElement(endElement)
			dataFlowShape.SetStartOrientation(ORIENTATION_VERTICAL)
			dataFlowShape.SetEndOrientation(ORIENTATION_VERTICAL)

			dataFlowShape.SetCornerOffsetRatio(1.1)
			dataFlowShape.SetStartRatio(0.5)
			dataFlowShape.SetEndRatio(0.5)
			diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)

			stage.Commit()
			return
		}
		if !isChecked && isShapePresent {
			isShapePresent = isChecked
			if dataFlowShape == nil {
				log.Panic("remove a non existing shape to product")
			}
			dataFlowShape.UnstageVoid(stage)

			// not necessary since there is a semantic rule (gong clean) that remove the shape from the slice when it is unstaged
			idx := slices.Index(diagramProcess.DataFlow_Shapes, dataFlowShape)
			diagramProcess.DataFlow_Shapes = slices.Delete(diagramProcess.DataFlow_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	dataFlowNode.OnClick = onNodeClicked(stager, dataFlow)

	for _, data := range dataFlow.Datas {
		stager.treeDataWithinDiagramProcessWithinDataFlow(dataFlowNode, diagramProcess, dataFlow, dataFlowShape, isShapePresent, data)
	}
}
