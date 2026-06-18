package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeConceptBSinDiagram(diagram *Diagram, concept *Concept, parentNode *tree.Node) {
	confNode := TreeNodeAndShapeConfigurationWithoutLink[
		*Concept, Concept,
		*Concept, Concept, // Parent is not used
		*ConceptShape, ConceptShape,
		*Diagram,
	]{
		diagram:                     diagram,
		parentNode:                  parentNode,
		element:                     concept,
		parentElement:               nil,
		elementsWhoseNodeIsExpanded: &diagram.ConceptsWhoseNodeIsExpanded,
		shapes:                      &diagram.Concept_Shapes,
		shapesMap:                   diagram.map_Concept_ConceptShape,
	}
	_ = addNodeToTreeWithoutLink(stager, confNode)
}
