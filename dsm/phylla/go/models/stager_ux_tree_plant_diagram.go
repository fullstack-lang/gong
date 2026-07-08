package models

import (
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePlantDiagram(
	plantDiagram *PlantDiagram,
	parentNodes *[]*tree.Node,
) {
	plantDiagramNode := &tree.Node{
		Name:              plantDiagram.Name,
		IsExpanded:        plantDiagram.IsExpanded,
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         plantDiagram.IsChecked,
		IsInEditMode:      plantDiagram.isInRenameMode,
	}
	*parentNodes = append(*parentNodes, plantDiagramNode)

	plantDiagramNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			for plantDiagram_ := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stager.stage) {
				plantDiagram_.IsChecked = false
			}
			plantDiagram.IsChecked = true
		} else {
			plantDiagram.IsChecked = false
			for plantDiagram_ := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stager.stage) {
				plantDiagram_.IsChecked = false
			}
		}
		stager.stage.Commit()
	}

	plantDiagramNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plantDiagram.IsExpanded)
	plantDiagramNode.OnNameChange = stager.onNameChange(plantDiagram)
	plantDiagramNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(plantDiagram, GetPointerToGongstructName[*PlantDiagram]())
		stager.stage.Commit()
	}

	axesShape := plantDiagram.AxesShape
	axesShapeNode := &tree.Node{
		Name:            axesShape.Name,
		IsNodeClickable: true,
	}
	axesShapeNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(axesShape, GetPointerToGongstructName[*AxesShape]())
		stager.stage.Commit()
	}
	visibilityButton := &tree.Button{
		Name:            "Hide",
		Icon:            string(buttons.BUTTON_visibility_off),
		ToolTipText:     "Hide from diagram",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
		OnClick: func() {
			axesShape.SetIsHidden(!axesShape.GetIsHidden())
			stager.stage.Commit()
		},
	}
	if axesShape.GetIsHidden() {
		visibilityButton.Icon = string(buttons.BUTTON_visibility)
		visibilityButton.Name = "Show"
		visibilityButton.ToolTipText = "Show on diagram"
	}
	axesShapeNode.Buttons = append(axesShapeNode.Buttons, visibilityButton)

	handleVisibilityButton := &tree.Button{
		Name:            "Hide Handle",
		Icon:            string(buttons.BUTTON_visibility_off),
		ToolTipText:     "Hide handles",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
		OnClick: func() {
			axesShape.SetIsWithHiddenHandle(!axesShape.GetIsWithHiddenHandle())
			stager.stage.Commit()
		},
	}
	if axesShape.GetIsWithHiddenHandle() {
		handleVisibilityButton.Icon = string(buttons.BUTTON_visibility)
		handleVisibilityButton.Name = "Show Handle"
		handleVisibilityButton.ToolTipText = "Show handles"
	}
	axesShapeNode.Buttons = append(axesShapeNode.Buttons, handleVisibilityButton)

	plantDiagramNode.Children = append(plantDiagramNode.Children, axesShapeNode)

}
