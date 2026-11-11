package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataFlowsWithinDiagramStructureWithinPort(
	diagramStructure *DiagramStructure,
	dataFlow *DataFlow,
	parentNode *tree.Node,
) {
	stager.treeDataFlowsWithinDiagramStructureWithinPortScoped(
		diagramStructure,
		nil,
		false,
		dataFlow,
		parentNode,
	)
}

func (stager *Stager) treeDataFlowsWithinDiagramStructureWithinPortScoped(
	diagramStructure *DiagramStructure,
	enclosingPartShape *PartShape,
	isNested bool,
	dataFlow *DataFlow,
	parentNode *tree.Node,
) {
	stager.treeDataFlowsScoped(
		diagramStructure,
		dataFlow,
		enclosingPartShape,
		isNested,
		&diagramStructure.DataFlowsWhoseDataNodeIsExpanded,
		parentNode,
	)
}
