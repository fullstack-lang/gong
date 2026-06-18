package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeConceptBSinDiagram(diagram *Diagram, concept *Concept, parentNode *tree.Node) {
	conceptNode := &tree.Node{
		Name:            concept.Name,
		IsNodeClickable: true,
		IsInEditMode:    concept.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, conceptNode)

	conceptNode.OnNameChange = func(newName string) {
		concept.SetName(newName)
		concept.SetIsInRenameMode(false)
		stager.stage.Commit()
	}

	addRenameButton(concept, conceptNode, stager)
}
