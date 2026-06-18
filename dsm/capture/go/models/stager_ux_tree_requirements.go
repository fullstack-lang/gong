package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeRequirementBSinDiagram(diagram *Diagram, requirement *Requirement, parentNode *tree.Node) {
	reqNode := &tree.Node{
		Name:            requirement.Name,
		IsNodeClickable: true,
		IsInEditMode:    requirement.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, reqNode)

	reqNode.OnNameChange = func(newName string) {
		requirement.SetName(newName)
		requirement.SetIsInRenameMode(false)
		stager.stage.Commit()
	}

	addRenameButton(requirement, reqNode, stager)
}
