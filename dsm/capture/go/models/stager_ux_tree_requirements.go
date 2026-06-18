package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeRequirementBSinDiagram(diagram *Diagram, requirement *Requirement, parentNode *tree.Node) {
	confNode := TreeNodeAndShapeConfigurationWithoutLink[
		*Requirement, Requirement,
		*Requirement, Requirement, // Parent is not used
		*RequirementShape, RequirementShape,
		*Diagram,
	]{
		diagram:                     diagram,
		parentNode:                  parentNode,
		element:                     requirement,
		parentElement:               nil,
		elementsWhoseNodeIsExpanded: &diagram.RequirementsWhoseNodeIsExpanded,
		shapes:                      &diagram.Requirement_Shapes,
		shapesMap:                   diagram.map_Requirement_RequirementShape,
	}
	_ = addNodeToTreeWithoutLink(stager, confNode)
}
