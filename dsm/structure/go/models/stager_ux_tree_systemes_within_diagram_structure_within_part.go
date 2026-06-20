package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeSystemesWithinDiagramStructureWithinPart(
	diagramStructure *DiagramStructure,
	system *System,
	part *Part,
	parentNode *tree.Node) {

	stage := stager.stage
	key := allocatedSystemShapeKey{
		part: part,
		system:     system,
	}
	allocatedSystemShape, ok := diagramStructure.map_AllocatedSystemShapeKey_AllocatedSystemShape[key]

	node := &tree.Node{
		Name:              system.Name,
		IsNodeClickable:   true,
		IsInEditMode:      system.GetIsInRenameMode(),
		IsChecked:         ok,
		HasCheckboxButton: true,
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(system, node, stager)

	node.OnNameChange = stager.onNameChange(system)
	node.OnClick = onNodeClicked(stager, system)

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !ok {
			allocatedSystemShape = (&AllocatedSystemShape{
				Name:        diagramStructure.GetName() + "-" + part.GetName() + "-" + system.GetName(),
				Part: part,
				System:     system,
			}).Stage(stage)
			diagramStructure.AllocatedSystemShapes = append(diagramStructure.AllocatedSystemShapes, allocatedSystemShape)
			stage.Commit()
		} else if !isChecked && ok {
			allocatedSystemShape.Unstage(stage)
			stage.Commit()
		}
	}
}
