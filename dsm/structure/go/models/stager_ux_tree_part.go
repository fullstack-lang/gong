package models

import (
	"fmt"
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePartWithinLibrary(
	library *Library,
	part *Part,
	parentNode *tree.Node,
) {
	partNode := &tree.Node{
		Name:            part.GetName(),
		IsExpanded:      slices.Contains(library.PartsWhoseNodeIsExpanded, part),
		IsNodeClickable: true,
		IsInEditMode:    part.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, partNode)

	addRenameButton(part, partNode, stager)

	partNode.OnNameChange = stager.onNameChange(part)
	partNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, part, &library.PartsWhoseNodeIsExpanded)
	partNode.OnClick = onNodeClicked(stager, part)
}

func (stager *Stager) treeParts(
	diagramStructure *DiagramStructure,
	part *Part,
	parentNode *tree.Node) {

	stage := stager.stage

	// find the partShape (if any)
	partShape, ok := diagramStructure.map_Part_PartShape[part]
	node := &tree.Node{
		Name:                    part.Name,
		IsExpanded:              slices.Contains(diagramStructure.PartWhoseNodeIsExpanded, part),
		IsNodeClickable:         true,
		IsInEditMode:            part.isInRenameMode,
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
	addRenameButton(part, node, stager)

	if partShape, ok = diagramStructure.map_Part_PartShape[part]; ok {
		node.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramStructure.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				partShape.SetIsHidden(!partShape.GetIsHidden())
				stage.Commit()
			},
		}
		if partShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	node.OnNameChange = stager.onNameChange(part)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, part, &diagramStructure.PartWhoseNodeIsExpanded)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			if partShape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			partShape = newShapeToDiagram(part, diagramStructure, &diagramStructure.Part_Shapes, stager, node.ClientOnY)

			stage.Commit()
			return
		} else {
			if partShape == nil {
				log.Panic("remove a non existing shape to product")
			}
			partShape.UnstageVoid(stage)
			idx := slices.Index(diagramStructure.Part_Shapes, partShape)
			diagramStructure.Part_Shapes = slices.Delete(diagramStructure.Part_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, part)

	// Ports
	portsNode := &tree.Node{
		Name:            "Ports",
		FontStyle:       tree.ITALIC,
		IsExpanded:      part.IsPortsNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, portsNode)
	portsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&part.IsPortsNodeExpanded)

	for _, port := range part.Ports {
		stager.treePort(diagramStructure, part, partShape, port, portsNode, &diagramStructure.PortsWhoseNodeIsExpanded)
	}

	// loook forward to https://github.com/golang/go/issues/61731
	// proposal: spec: support type inference on generic structs
	conf := ItemAndShapeButtonConfiguration[
		Port, *Port, // AT, PAT (Added Element)
		Part, *Part, // ParentAT, PParentAT (Parent Element)
		PortShape, *PortShape, // CT, PCT (Concrete Shape)

	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Port, *Port, // AT, PAT (Added Element)
			Part, *Part, // ParentAT, PParentAT (Parent Element)
		]{
			parentNode:                         portsNode,
			sliceForNewAddedItem:               &part.Ports,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &part.IsPortsNodeExpanded,
			parentElement:                      part,
		},
		receivingDiagram:      diagramStructure,
		sliceForNewAddedShape: &diagramStructure.Port_Shapes,
	}
	callback := addCreateItemAndShapeButton(stager, conf)

	callback.OnBeforeCommit = func() {
		// the added shape must (?) be the last item
		newShape := diagramStructure.Port_Shapes[len(diagramStructure.Port_Shapes)-1]
		newShape.SetWidth(defaultPortWidth)
		newShape.SetHeight(defaultPortHeight)

		port := callback.createdItem
		port.SetName(fmt.Sprintf("P%d", len(part.Ports)))
		port.SetIsInRenameMode(false)
	}

	controlflowsNode := &tree.Node{
		Name:            "ControlFlows",
		FontStyle:       tree.ITALIC,
		IsExpanded:      part.IsControlFlowsNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, controlflowsNode)
	controlflowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&part.IsControlFlowsNodeExpanded)

	for _, controlflow := range part.ControlFlows {
		stager.treeControlFlows(diagramStructure, controlflow, controlflowsNode, &diagramStructure.ControlFlowsWhoseNodeIsExpanded)
	}
}
