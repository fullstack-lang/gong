package models

import (
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

	// find the shape (if any)
	shape, ok := diagramStructure.map_Part_PartShape[part]
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

	if shape, ok := diagramStructure.map_Part_PartShape[part]; ok {
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

	node.OnNameChange = stager.onNameChange(part)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, part, &diagramStructure.PartWhoseNodeIsExpanded)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(part, diagramStructure, &diagramStructure.Part_Shapes, stage)

			stage.Commit()
			return
		} else {
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)
			idx := slices.Index(diagramStructure.Part_Shapes, shape)
			diagramStructure.Part_Shapes = slices.Delete(diagramStructure.Part_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, part)

	// Allocated Resources
	allocatedResourcesNode := &tree.Node{
		Name:            "Allocated Resources",
		FontStyle:       tree.ITALIC,
		IsExpanded:      part.IsResourcesNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, allocatedResourcesNode)
	allocatedResourcesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&part.IsResourcesNodeExpanded)

	for _, resource := range part.Resources {
		stager.treeAllocatedResourceWithinDiagramWithinPart(diagramStructure, resource, part, allocatedResourcesNode)
	}

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
		stager.treePort(diagramStructure, part, port, portsNode, &diagramStructure.PortsWhoseNodeIsExpanded)
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
	addCreateItemAndShapeButton(stager, conf)

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
