package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataFlowsWithinProcess(
	diagramProcess *DiagramProcess,
	dataFlow *DataFlow,
	parentNode *tree.Node,
) {
	stage := stager.stage

	// find the shape (if any)
	shape, isShapePresent := diagramProcess.map_DataFlow_DataFlowShape[dataFlow]

	node := &tree.Node{
		Name:              dataFlow.GetName(),
		IsExpanded:        false,
		IsNodeClickable:   true,
		IsInEditMode:      dataFlow.GetIsInRenameMode(),
		HasCheckboxButton: true,
		IsChecked:         isShapePresent,
	}
	parentNode.Children = append(parentNode.Children, node)

	addRenameButton(dataFlow, node, stager)

	if shape, ok := diagramProcess.map_DataFlow_DataFlowShape[dataFlow]; ok {
		node.IsChecked = true
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
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	node.OnClick = func(frontNode *tree.Node) {
		if frontNode.Name != dataFlow.GetName() {
			dataFlow.SetName(frontNode.Name)
			dataFlow.SetIsInRenameMode(false)
			stager.stage.Commit()
			return
		}
		if frontNode.IsChecked && !isShapePresent {
			isShapePresent = frontNode.IsChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			// shape = newShapeToDiagram(controlflow, diagramProcess, &diagramProcess.ControlFlowShapes, stage)
			// addAssociationShapeToDiagram(stager, controlflow.Start, controlflow.End, &diagramProcess.ControlFlowShapes)
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
			diagramProcess.DataFlowShapes = append(diagramProcess.DataFlowShapes, dataFlowShape)

			stage.Commit()
			return
		}
		if !frontNode.IsChecked && isShapePresent {
			isShapePresent = frontNode.IsChecked
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)

			// not necessary since there is a semantic rule (gong clean) that remove the shape from the slice when it is unstaged
			idx := slices.Index(diagramProcess.DataFlowShapes, shape)
			diagramProcess.DataFlowShapes = slices.Delete(diagramProcess.DataFlowShapes, idx, idx+1)
			stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(dataFlow, GetPointerToGongstructName[*DataFlow]())
		stager.stage.Commit()
	}
}
