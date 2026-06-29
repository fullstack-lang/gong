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
	_, ok := diagramStructure.map_Port_PortShape[port]
	node := &tree.Node{
		Name:                    port.GetName(),
		IsExpanded:              slices.Contains(*portWhoseNodeIsExpanded, port),
		IsNodeClickable:         true,
		IsInEditMode:            port.isInRenameMode,
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove the port shape"
			}
			return "Click to create a port shape for this port within this diagram"
		}(),
	}
	parentNode.Children = append(parentNode.Children, node)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, portWhoseNodeIsExpanded)
	node.OnNameChange = stager.onNameChange(port)
	node.OnClick = onNodeClicked(stager, port)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			portShape := newShapeToDiagram(port, diagramStructure, &diagramStructure.Port_Shapes, stager, node.ClientOnY)

			// default position on the left of the part shape
			portShape.X = partShape.X - defaultPortWidth/2
			portShape.Y = partShape.Y + partShape.Height/2
			portShape.Width = defaultPortWidth
			portShape.Height = defaultPortHeight

			portShape.Port = port
			portShape.Stage(stager.stage)
		} else {
			portShape := diagramStructure.map_Port_PortShape[port]
			portShape.UnstageVoid(stager.stage)
			idx := slices.Index(diagramStructure.Port_Shapes, portShape)
			diagramStructure.Port_Shapes = slices.Delete(diagramStructure.Port_Shapes, idx, idx+1)
		}
		stager.stage.Commit()
	}

	if ok {
		portShape := diagramStructure.map_Port_PortShape[port]
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
		stager.treeControlFlowsWithinPort(diagramStructure, controlFlow, nodeOutControlFlows)
	}

	nodeInControlFlows := &tree.Node{
		Name:            "in control flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseInControlFlowsNodeIsExpanded, port),
	}
	node.Children = append(node.Children, nodeInControlFlows)
	nodeInControlFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseInControlFlowsNodeIsExpanded)
	for _, controlFlow := range port.inControlFlows {
		stager.treeControlFlowsWithinPort(diagramStructure, controlFlow, nodeInControlFlows)
	}

	nodeOutDataFlows := &tree.Node{
		Name:            "out data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseOutDataFlowsNodeIsExpanded, port),
	}
	node.Children = append(node.Children, nodeOutDataFlows)
	nodeOutDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseOutDataFlowsNodeIsExpanded)
	for _, dataFlow := range port.outDataFlows {
		stager.treeDataFlowsWithinDiagramStructureWithinPort(diagramStructure, dataFlow, nodeOutDataFlows)
	}

	nodeInDataFlows := &tree.Node{
		Name:            "in data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseInDataFlowsNodeIsExpanded, port),
	}
	node.Children = append(node.Children, nodeInDataFlows)
	nodeInDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseInDataFlowsNodeIsExpanded)
	for _, dataFlow := range port.inDataFlows {
		stager.treeDataFlowsWithinDiagramStructureWithinPort(diagramStructure, dataFlow, nodeInDataFlows)
	}
}
