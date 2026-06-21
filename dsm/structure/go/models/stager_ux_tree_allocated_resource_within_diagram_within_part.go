package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeAllocatedResourceWithinDiagramWithinPart(
	diagramStructure *DiagramStructure,
	resource *Resource,
	part *Part,
	parentNode *tree.Node) {

	stage := stager.stage
	key := allocatedResourceShapeKey{
		part:     part,
		resource: resource,
	}
	allocatedResourceShape, ok := diagramStructure.map_AllocatedResourceShapeKey_AllocatedResourceShape[key]

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
				Name:     diagramStructure.GetName() + "-" + part.GetName() + "-" + resource.GetName(),
				Part:     part,
				Resource: resource,
			}).Stage(stage)
			diagramStructure.AllocatedResourceShapes = append(diagramStructure.AllocatedResourceShapes, allocatedResourceShape)
			stage.Commit()
		} else if !isChecked && ok {
			allocatedResourceShape.Unstage(stage)
			stage.Commit()
		}
	}
}
