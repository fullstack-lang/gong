package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePort(
	diagramStructure *DiagramStructure,
	part *Part,
	port *Port,
	parentNode *tree.Node,
	portWhoseNodeIsExpanded *[]*Port,
) {
	portNodeConf := TreeNodeAndShapeConfigurationWithoutLink[
		*Port, Port,
		*Part, Part,
		*PortShape, PortShape,
		*DiagramStructure,
	]{
		diagram:                     diagramStructure,
		parentNode:                  parentNode,
		element:                     port,
		parentElement:               part,
		elementsWhoseNodeIsExpanded: portWhoseNodeIsExpanded,
		shapes:                      &diagramStructure.Port_Shapes,
		shapesMap:                   diagramStructure.map_Port_PortShape,
	}
	portNode := addNodeToTreeWithoutLink(stager, portNodeConf)

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
