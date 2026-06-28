package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeExternalParts(
	diagramStructure *DiagramStructure,
	externalPart *Part,
	parentNode *tree.Node) {

	stage := stager.stage

	// find the shape (if any)
	shape, ok := diagramStructure.map_Part_ExternalPartShape[externalPart]
	node := &tree.Node{
		Name:                    externalPart.Name,
		IsExpanded:              slices.Contains(diagramStructure.ExternalPartWhoseNodeIsExpanded, externalPart),
		IsNodeClickable:         true,
		IsInEditMode:            externalPart.isInRenameMode,
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove the part shape"
			}
			return "Click to create a part shape for this part within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(externalPart, node, stager)

	if ok {
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

	node.OnNameChange = stager.onNameChange(externalPart)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, externalPart, &diagramStructure.ExternalPartWhoseNodeIsExpanded)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(externalPart, diagramStructure, &diagramStructure.ExternalPart_Shapes, stager, node.ClientOnY)

			stage.Commit()
			return
		} else {
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)
			idx := slices.Index(diagramStructure.ExternalPart_Shapes, shape)
			diagramStructure.ExternalPart_Shapes = slices.Delete(diagramStructure.ExternalPart_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, externalPart)

	// out data flows and in data flows
	nodeOutDataFlows := &tree.Node{
		Name:            "out data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(diagramStructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, externalPart),
	}
	node.Children = append(node.Children, nodeOutDataFlows)
	nodeOutDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, externalPart, &diagramStructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded)
	for _, dataFlow := range externalPart.outDataFlows {
		stager.treeDataFlowsWithinDiagramStructureWithinPort(diagramStructure, dataFlow, nodeOutDataFlows)
	}

	nodeInDataFlows := &tree.Node{
		Name:            "in data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(diagramStructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded, externalPart),
	}
	node.Children = append(node.Children, nodeInDataFlows)
	nodeInDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, externalPart, &diagramStructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded)
	for _, dataFlow := range externalPart.inDataFlows {
		stager.treeDataFlowsWithinDiagramStructureWithinPort(diagramStructure, dataFlow, nodeInDataFlows)
	}
}
