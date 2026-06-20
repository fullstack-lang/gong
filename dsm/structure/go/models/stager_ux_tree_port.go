package models

import (
	"slices"

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
	portNode := &tree.Node{
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
	parentNode.Children = append(parentNode.Children, portNode)
	portNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, portWhoseNodeIsExpanded)
	portNode.OnNameChange = stager.onNameChange(port)
	portNode.OnClick = onNodeClicked(stager, port)
	portNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			portShape := newShapeToDiagram(port, diagramStructure, &diagramStructure.Port_Shapes, stager.stage)

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

	nodeOutControlFlows := &tree.Node{
		Name:            "out control flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseOutControlFlowsNodeIsExpanded, port),
	}
	portNode.Children = append(portNode.Children, nodeOutControlFlows)
	nodeOutControlFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseOutControlFlowsNodeIsExpanded)
	for _, controlFlow := range port.outControlFlows {
		stager.treeControlFlowsWithinPort(diagramStructure, controlFlow, nodeOutControlFlows)
	}

	nodeInControlFlows := &tree.Node{
		Name:            "in control flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseInControlFlowsNodeIsExpanded, port),
	}
	portNode.Children = append(portNode.Children, nodeInControlFlows)
	nodeInControlFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseInControlFlowsNodeIsExpanded)
	for _, controlFlow := range port.inControlFlows {
		stager.treeControlFlowsWithinPort(diagramStructure, controlFlow, nodeInControlFlows)
	}

	nodeOutDataFlows := &tree.Node{
		Name:            "out data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseOutDataFlowsNodeIsExpanded, port),
	}
	portNode.Children = append(portNode.Children, nodeOutDataFlows)
	nodeOutDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseOutDataFlowsNodeIsExpanded)
	for _, dataFlow := range port.outDataFlows {
		stager.treeDataFlowsWithinDiagramStructureWithinPort(diagramStructure, dataFlow, nodeOutDataFlows)
	}

	nodeInDataFlows := &tree.Node{
		Name:            "in data flows",
		IsNodeClickable: true,
		IsExpanded:      slices.Contains(part.PortWhoseInDataFlowsNodeIsExpanded, port),
	}
	portNode.Children = append(portNode.Children, nodeInDataFlows)
	nodeInDataFlows.OnIsExpandedChange = onIsExpandedChangeSlice(stager, port, &part.PortWhoseInDataFlowsNodeIsExpanded)
	for _, dataFlow := range port.inDataFlows {
		stager.treeDataFlowsWithinDiagramStructureWithinPort(diagramStructure, dataFlow, nodeInDataFlows)
	}
}
