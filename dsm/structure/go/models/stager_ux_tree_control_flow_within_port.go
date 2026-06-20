package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeControlFlowsWithinPort(
	diagramStructure *DiagramStructure,
	controlFlow *ControlFlow,
	parentNode *tree.Node,
) {
	stage := stager.stage

	// find the shape (if any)
	shape, isShapePresent := diagramStructure.map_ControlFlow_ControlFlowShape[controlFlow]

	isStartShapePresent := false
	isEndShapePresent := false
	if controlFlow.Start != nil {
		_, isStartShapePresent = diagramStructure.map_Port_PortShape[controlFlow.Start]
	}
	if controlFlow.End != nil {
		_, isEndShapePresent = diagramStructure.map_Port_PortShape[controlFlow.End]
	}
	isCheckboxDisabled := !(isStartShapePresent && isEndShapePresent)

	node := &tree.Node{
		Name:                    controlFlow.GetName(),
		IsExpanded:              false,
		IsNodeClickable:         true,
		IsInEditMode:            controlFlow.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isShapePresent,
		IsCheckboxDisabled:      isCheckboxDisabled,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if isCheckboxDisabled {
				return "A control flow shape cannot be created if the start or end port shape is not present from the diagram"
			}
			if isShapePresent {
				return "Click to remove the control flow shape"
			}
			return "Click to create a control flow shape for this control flow within this diagram"
		}(),
	}

	if isCheckboxDisabled {
		node.CheckboxHasToolTip = true
		node.CheckboxToolTipText = "Start or end port shape is missing from the diagram"
	}
	parentNode.Children = append(parentNode.Children, node)

	addRenameButton(controlFlow, node, stager)

	if shape, ok := diagramStructure.map_ControlFlow_ControlFlowShape[controlFlow]; ok {
		node.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramStructure.GetName(),
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

	node.OnNameChange = stager.onNameChange(controlFlow)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isShapePresent {
			isShapePresent = isChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			// shape = newShapeToDiagram(controlflow, diagramStructure, &diagramStructure.ControlFlowShapes, stage)
			// addAssociationShapeToDiagram(stager, controlflow.Start, controlflow.End, &diagramStructure.ControlFlowShapes)
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
			diagramStructure.ControlFlow_Shapes = append(diagramStructure.ControlFlow_Shapes, controlFlowShape)

			stage.Commit()
			return
		}
		if !isChecked && isShapePresent {
			isShapePresent = isChecked
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)

			// not necessary since there is a semantic rule (gong clean) that remove the shape from the slice when it is unstaged
			idx := slices.Index(diagramStructure.ControlFlow_Shapes, shape)
			diagramStructure.ControlFlow_Shapes = slices.Delete(diagramStructure.ControlFlow_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, controlFlow)
}
