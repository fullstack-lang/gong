package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataFlowsWithinDiagramProcessWithinTask(
	diagramProcess *DiagramProcess,
	dataFlow *DataFlow,
	parentNode *tree.Node,
) {
	stage := stager.stage

	// find the dataFlowShape (if any)
	dataFlowShape, isShapePresent := diagramProcess.map_DataFlow_DataFlowShape[dataFlow]

	isStartShapePresent := false
	isEndShapePresent := false
	if dataFlow.Start != nil {
		_, isStartShapePresent = diagramProcess.map_Task_TaskShape[dataFlow.Start]
	}
	if dataFlow.End != nil {
		_, isEndShapePresent = diagramProcess.map_Task_TaskShape[dataFlow.End]
	}
	isCheckboxDisabled := !(isStartShapePresent && isEndShapePresent)

	dataFlowNode := &tree.Node{
		Name:                    dataFlow.GetName(),
		IsExpanded:              slices.Contains(diagramProcess.DataFlowsWhoseDataNodeIsExpanded, dataFlow),
		IsNodeClickable:         true,
		IsInEditMode:            dataFlow.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isShapePresent,
		IsCheckboxDisabled:      isCheckboxDisabled,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if isCheckboxDisabled {
				return "A data shape cannot be created if the start or end task shape is not present from the diagram"
			}
			if isShapePresent {
				return "Click to remove the data flow shape"
			}
			return "Click to create a data flow shape for this data flow within this diagram"
		}(),
	}

	if isCheckboxDisabled {
		dataFlowNode.CheckboxHasToolTip = true
		dataFlowNode.CheckboxToolTipText = "Start or end task shape is missing from the diagram"
	}
	parentNode.Children = append(parentNode.Children, dataFlowNode)

	addRenameButton(dataFlow, dataFlowNode, stager)

	dataFlowShape, shapePresent := diagramProcess.map_DataFlow_DataFlowShape[dataFlow]
	if shapePresent {
		dataFlowNode.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramProcess.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				dataFlowShape.SetIsHidden(!dataFlowShape.GetIsHidden())
				stage.Commit()
			},
		}
		if dataFlowShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		dataFlowNode.Buttons = append(dataFlowNode.Buttons, visibilityButton)
	}

	dataFlowNode.OnClick = func(frontNode *tree.Node) {
		if frontNode.Name != dataFlow.GetName() {
			dataFlow.SetName(frontNode.Name)
			dataFlow.SetIsInRenameMode(false)
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded != dataFlowNode.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(diagramProcess.DataFlowsWhoseDataNodeIsExpanded, dataFlow) == -1 {
					diagramProcess.DataFlowsWhoseDataNodeIsExpanded = append(diagramProcess.DataFlowsWhoseDataNodeIsExpanded, dataFlow)
				}
			} else {
				if idx := slices.Index(diagramProcess.DataFlowsWhoseDataNodeIsExpanded, dataFlow); idx != -1 {
					diagramProcess.DataFlowsWhoseDataNodeIsExpanded = slices.Delete(diagramProcess.DataFlowsWhoseDataNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
			return
		}
		if frontNode.IsChecked && !isShapePresent {
			isShapePresent = frontNode.IsChecked
			if dataFlowShape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			// shape = newShapeToDiagram(dataflow, diagramProcess, &diagramProcess.DataFlowShapes, stage)
			// addAssociationShapeToDiagram(stager, dataflow.Start, dataflow.End, &diagramProcess.DataFlowShapes)
			dataFlowShape := (&DataFlowShape{
				DataFlow: dataFlow,
			}).Stage(stage)
			dataFlowShape.SetName(dataFlow.Start.GetName() + " to " + dataFlow.End.GetName())
			dataFlowShape.SetAbstractStartElement(dataFlow.Start)
			dataFlowShape.SetAbstractEndElement(dataFlow.End)
			dataFlowShape.SetStartOrientation(ORIENTATION_VERTICAL)
			dataFlowShape.SetEndOrientation(ORIENTATION_VERTICAL)

			dataFlowShape.SetCornerOffsetRatio(1.1)
			dataFlowShape.SetStartRatio(0.5)
			dataFlowShape.SetEndRatio(0.5)
			diagramProcess.DataFlow_Shapes = append(diagramProcess.DataFlow_Shapes, dataFlowShape)

			stage.Commit()
			return
		}
		if !frontNode.IsChecked && isShapePresent {
			isShapePresent = frontNode.IsChecked
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
		stager.probeForm.FillUpFormFromGongstruct(dataFlow, GetPointerToGongstructName[*DataFlow]())
		stager.stage.Commit()
	}

	for _, data := range dataFlow.Datas {
		stager.treeDataWithinDiagramProcessWithinDataFlow(dataFlowNode, diagramProcess, dataFlow, dataFlowShape, shapePresent, data)
	}
}
