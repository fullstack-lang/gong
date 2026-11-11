package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

const defaultPortWidth = 40.0
const defaultPortHeight = 40.0

func (stager *Stager) treePort(
	diagramStructure *DiagramStructure,
	part *Part,
	partShape *PartShape,
	port *Port,
	parentNode *tree.Node,
	portWhoseNodeIsExpanded *[]*Port,
) {
	isPartShapeMissing := partShape == nil
	var portShape *PortShape
	var ok bool
	if !isPartShapeMissing && diagramStructure.map_PortShapeKey_PortShape != nil {
		portShape, ok = diagramStructure.map_PortShapeKey_PortShape[portShapeKey{
			port:            port,
			parentPartShape: partShape,
		}]
	}

	node := &tree.Node{
		Name:                    port.GetName(),
		IsExpanded:              slices.Contains(*portWhoseNodeIsExpanded, port),
		IsNodeClickable:         true,
		IsInEditMode:            port.isInRenameMode,
		HasCheckboxButton:       true,
		IsChecked:               ok,
		IsCheckboxDisabled:      isPartShapeMissing,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if isPartShapeMissing {
				return "Enclosing part shape must be present in the diagram"
			}
			if ok {
				return "Click to remove the port shape"
			}
			return "Click to create a port shape for this port within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, portWhoseNodeIsExpanded)
	node.OnNameChange = stager.onNameChange(port)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(port, "Port")
		if ok && portShape != nil {
			fillUpShapeAppearanceForm(stager, portShape)
		} else {
			clearShapeAppearanceForm(stager)
		}
		stager.stage.CommitWithSuspendedCallbacks()
	}

	addRenameButton(port, node, stager)
	deleteButton := &tree.Button{
		Name:            "Delete",
		Icon:            string(buttons.BUTTON_delete),
		ToolTipText:     "Delete the port",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			part.Ports = slices.DeleteFunc(part.Ports, func(p *Port) bool { return p == port })
			port.Unstage(stager.stage)
			if portShape != nil {
				portShape.UnstageVoid(stager.stage)
				diagramStructure.Port_Shapes = slices.DeleteFunc(diagramStructure.Port_Shapes, func(ps *PortShape) bool { return ps == portShape })
			}
			stager.stage.Commit()
		},
	}
	if node.Menu == nil {
		node.Menu = &tree.Menu{Name: "Menu"}
	}
	node.Menu.Buttons = append(node.Menu.Buttons, deleteButton)

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isPartShapeMissing {
			return
		}
		if isChecked {
			newPortShape := newShapeToDiagram(port, diagramStructure, &diagramStructure.Port_Shapes, stager, node.ClientOnY)
			newPortShape.ParentPartShape = partShape

			// default position on the left of the part shape
			newPortShape.X = partShape.X - defaultPortWidth/2
			newPortShape.Y = partShape.Y + partShape.Height/2 - defaultPortHeight/2
			newPortShape.Width = defaultPortWidth
			newPortShape.Height = defaultPortHeight

			newPortShape.Port = port
			newPortShape.Stage(stager.stage)
		} else if portShape != nil {
			portShape.UnstageVoid(stager.stage)
			idx := slices.Index(diagramStructure.Port_Shapes, portShape)
			if idx >= 0 {
				diagramStructure.Port_Shapes = slices.Delete(diagramStructure.Port_Shapes, idx, idx+1)
			}
		}
		stager.stage.Commit()
	}

	if ok && portShape != nil {
		visibilityButton := &tree.Button{
			Name:            diagramStructure.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				portShape.SetIsHidden(!portShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if portShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	nodeOutControlFlows := &tree.Node{
		Name:            "out control flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseOutControlFlowsNodeIsExpanded, port),
	}
	node.Children = append(node.Children, nodeOutControlFlows)
	nodeOutControlFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseOutControlFlowsNodeIsExpanded)
	for _, controlFlow := range port.outControlFlows {
		stager.treeControlFlowsWithinPortScoped(diagramStructure, partShape, controlFlow, nodeOutControlFlows)
	}

	nodeInControlFlows := &tree.Node{
		Name:            "in control flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseInControlFlowsNodeIsExpanded, port),
	}
	node.Children = append(node.Children, nodeInControlFlows)
	nodeInControlFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseInControlFlowsNodeIsExpanded)
	for _, controlFlow := range port.inControlFlows {
		stager.treeControlFlowsWithinPortScoped(diagramStructure, partShape, controlFlow, nodeInControlFlows)
	}

	var enclosingPartShape *PartShape
	isNestedFlow := isPartShapeMissing
	if partShape != nil {
		enclosingPartShape = partShape.ParentPartShape
		isNestedFlow = enclosingPartShape != nil
	}

	nodeOutDataFlows := &tree.Node{
		Name:            "out data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseOutDataFlowsNodeIsExpanded, port),
	}
	node.Children = append(node.Children, nodeOutDataFlows)
	nodeOutDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseOutDataFlowsNodeIsExpanded)
	for _, dataFlow := range port.outDataFlows {
		stager.treeDataFlowsWithinDiagramStructureWithinPortScoped(diagramStructure, enclosingPartShape, isNestedFlow, dataFlow, nodeOutDataFlows)
	}

	nodeInDataFlows := &tree.Node{
		Name:            "in data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseInDataFlowsNodeIsExpanded, port),
	}
	node.Children = append(node.Children, nodeInDataFlows)
	nodeInDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseInDataFlowsNodeIsExpanded)
	for _, dataFlow := range port.inDataFlows {
		stager.treeDataFlowsWithinDiagramStructureWithinPortScoped(diagramStructure, enclosingPartShape, isNestedFlow, dataFlow, nodeInDataFlows)
	}
}
