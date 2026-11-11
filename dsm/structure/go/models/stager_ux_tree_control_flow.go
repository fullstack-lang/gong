package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeControlFlows(
	diagramStructure *DiagramStructure,
	controlFlow *ControlFlow,
	parentNode *tree.Node,
	controlflowWhoseNodeIsExpanded *[]*ControlFlow,
) {
	var partShape *PartShape
	if controlFlow.Start != nil && controlFlow.Start.owningPart != nil && diagramStructure.map_PartShapeKey_PartShape != nil {
		partShape = diagramStructure.map_PartShapeKey_PartShape[partShapeKey{
			part:            controlFlow.Start.owningPart,
			parentPartShape: nil,
		}]
	}
	stager.treeControlFlowsScoped(diagramStructure, partShape, controlFlow, parentNode, controlflowWhoseNodeIsExpanded)
}

func (stager *Stager) treeControlFlowsScoped(
	diagramStructure *DiagramStructure,
	partShape *PartShape,
	controlFlow *ControlFlow,
	parentNode *tree.Node,
	controlflowWhoseNodeIsExpanded *[]*ControlFlow,
) {
	stage := stager.stage

	var shape *ControlFlowShape
	var isShapePresent bool
	if partShape != nil && diagramStructure.map_ControlFlowShapeKey_ControlFlowShape != nil {
		shape, isShapePresent = diagramStructure.map_ControlFlowShapeKey_ControlFlowShape[controlFlowShapeKey{
			controlFlow:     controlFlow,
			parentPartShape: partShape,
		}]
	}

	isStartShapePresent := false
	isEndShapePresent := false
	if partShape != nil && diagramStructure.map_PortShapeKey_PortShape != nil {
		if controlFlow.Start != nil {
			_, isStartShapePresent = diagramStructure.map_PortShapeKey_PortShape[portShapeKey{
				port:            controlFlow.Start,
				parentPartShape: partShape,
			}]
		}
		if controlFlow.End != nil {
			_, isEndShapePresent = diagramStructure.map_PortShapeKey_PortShape[portShapeKey{
				port:            controlFlow.End,
				parentPartShape: partShape,
			}]
		}
	}
	isCheckboxDisabled := !(isStartShapePresent && isEndShapePresent)

	node := &tree.Node{
		Name:               controlFlow.GetName(),
		IsExpanded:         slices.Contains(*controlflowWhoseNodeIsExpanded, controlFlow),
		IsNodeClickable:    true,
		IsInEditMode:       controlFlow.GetIsInRenameMode(),
		HasCheckboxButton:  true,
		IsChecked:          isShapePresent,
		IsCheckboxDisabled: isCheckboxDisabled,
	}

	if isCheckboxDisabled {
		node.CheckboxHasToolTip = true
		node.CheckboxToolTipText = "Start or end port shape is missing from the diagram"
	}
	parentNode.Children = append(parentNode.Children, node)

	addRenameButton(controlFlow, node, stager)

	if isShapePresent && shape != nil {
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
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, controlFlow, controlflowWhoseNodeIsExpanded)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isShapePresent {
			isShapePresent = isChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			controlFlowShape := (&ControlFlowShape{
				ControlFlow:     controlFlow,
				ParentPartShape: partShape,
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
			idx := slices.Index(diagramStructure.ControlFlow_Shapes, shape)
			if idx >= 0 {
				diagramStructure.ControlFlow_Shapes = slices.Delete(diagramStructure.ControlFlow_Shapes, idx, idx+1)
			}
			stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, controlFlow)
}
