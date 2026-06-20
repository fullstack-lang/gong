package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeAllocatedResourceWithinDiagramWithinParticipant(
	diagramProcess *DiagramProcess,
	resource *Resource,
	participant *Participant,
	parentNode *tree.Node) {

	stage := stager.stage
	key := allocatedResourceShapeKey{
		participant: participant,
		resource:    resource,
	}
	allocatedResourceShape, ok := diagramProcess.map_AllocatedResourceShapeKey_AllocatedResourceShape[key]

	node := &tree.Node{
		Name:              resource.Name,
		IsNodeClickable:   true,
		IsInEditMode:      resource.isInRenameMode,
		IsChecked:         ok,
		HasCheckboxButton: true,
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(resource, node, stager)

	node.OnNameChange = stager.onNameChange(resource)
	node.OnClick = onNodeClicked(stager, resource)

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !ok {
			allocatedResourceShape = (&AllocatedResourceShape{
				Name:        diagramProcess.GetName() + "-" + participant.GetName() + "-" + resource.GetName(),
				Participant: participant,
				Resource:    resource,
			}).Stage(stage)
			diagramProcess.AllocatedResourceShapes = append(diagramProcess.AllocatedResourceShapes, allocatedResourceShape)
			stage.Commit()
		} else if !isChecked && ok {
			allocatedResourceShape.Unstage(stage)
			stage.Commit()
		}
	}
}
