package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeControlFlows(
	diagramProcess *DiagramProcess,
	controlFlow *ControlFlow,
	parentNode *tree.Node,
	controlflowWhoseNodeIsExpanded *[]*ControlFlow,
) {
	stage := stager.stage

	// find the shape (if any)
	shape, isShapePresent := diagramProcess.map_ControlFlow_ControlFlowShape[controlFlow]

	node := &tree.Node{
		Name:              controlFlow.GetName(),
		IsExpanded:        slices.Index(*controlflowWhoseNodeIsExpanded, controlFlow) != -1,
		IsNodeClickable:   true,
		IsInEditMode:      controlFlow.GetIsInRenameMode(),
		HasCheckboxButton: true,
		IsChecked:         isShapePresent,
	}
	parentNode.Children = append(parentNode.Children, node)

	addRenameButton(controlFlow, node, stager)

	if shape, ok := diagramProcess.map_ControlFlow_ControlFlowShape[controlFlow]; ok {
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
		if frontNode.Name != controlFlow.GetName() {
			controlFlow.SetName(frontNode.Name)
			controlFlow.SetIsInRenameMode(false)
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
			controlFlowShape := (&ControlFlowShape{
				ControlFlow: controlFlow,
			}).Stage(stage)
			controlFlowShape.SetName(controlFlow.Start.GetName() + " to " + controlFlow.End.GetName())
			controlFlowShape.SetAbstractStartElement(controlFlow.Start)
			controlFlowShape.SetAbstractEndElement(controlFlow.End)
			controlFlowShape.SetStartOrientation(ORIENTATION_VERTICAL)
			controlFlowShape.SetEndOrientation(ORIENTATION_VERTICAL)

			controlFlowShape.SetCornerOffsetRatio(1.1)
			controlFlowShape.SetStartRatio(0.5)
			controlFlowShape.SetEndRatio(0.5)
			diagramProcess.ControlFlow_Shapes = append(diagramProcess.ControlFlow_Shapes, controlFlowShape)

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
			idx := slices.Index(diagramProcess.ControlFlow_Shapes, shape)
			diagramProcess.ControlFlow_Shapes = slices.Delete(diagramProcess.ControlFlow_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
		if frontNode.IsExpanded != node.IsExpanded {
			if frontNode.IsExpanded {
				if slices.Index(*controlflowWhoseNodeIsExpanded, controlFlow) == -1 {
					*controlflowWhoseNodeIsExpanded = append(*controlflowWhoseNodeIsExpanded, controlFlow)
				}
			} else {
				if idx := slices.Index(*controlflowWhoseNodeIsExpanded, controlFlow); idx != -1 {
					*controlflowWhoseNodeIsExpanded = slices.Delete(*controlflowWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(controlFlow, GetPointerToGongstructName[*ControlFlow]())
		stager.stage.Commit()
	}
}
