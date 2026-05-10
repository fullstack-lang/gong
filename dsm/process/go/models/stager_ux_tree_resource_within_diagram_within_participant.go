package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeResourceWithinDiagramWithinParticipant(
	diagramProcess *DiagramProcess,
	resource *Resource,
	participant *Participant,
	parentNode *tree.Node) {

	stage := stager.stage
	_ = stage
	_ = participant
	_ = diagramProcess

	node := &tree.Node{
		Name:                    resource.Name,
		IsExpanded:              false,
		IsNodeClickable:         true,
		IsInEditMode:            resource.isInRenameMode,
		IsChecked:               false,
		IsCheckboxDisabled:      true,
		HasCheckboxButton:       false,
		HasSecondCheckboxButton: false,
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(resource, node, stager)

	node.OnNameChange = stager.onNameChange(resource)
	node.OnClick = onNodeClicked(stager, resource)

}
