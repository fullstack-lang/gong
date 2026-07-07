package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePlant(plant *Plant, parentNodes *[]*tree.Node) {
	plantNode := &tree.Node{
		Name:            plant.Name,
		IsExpanded:      plant.IsExpanded,
		IsNodeClickable: true,
		IsInEditMode:    plant.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, plantNode)

	plantNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plant.IsExpanded)
	plantNode.OnNameChange = stager.onNameChange(plant)
	plantNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(plant, GetPointerToGongstructName[*Plant]())
		stager.selectedPlant = plant
		stager.stage.Commit()
	}
}
