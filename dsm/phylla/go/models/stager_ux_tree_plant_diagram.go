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

	referenceRhombus := plantDiagram.ReferenceRhombus
	if referenceRhombus != nil {
		referenceRhombusNode := &tree.Node{
			Name:            referenceRhombus.Name,
			IsNodeClickable: true,
		}
		referenceRhombusNode.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(referenceRhombus, GetPointerToGongstructName[*ReferenceRhombus]())
			stager.stage.Commit()
		}
		rrVisibilityButton := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				referenceRhombus.SetIsHidden(!referenceRhombus.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if referenceRhombus.GetIsHidden() {
			rrVisibilityButton.Icon = string(buttons.BUTTON_visibility)
			rrVisibilityButton.Name = "Show"
			rrVisibilityButton.ToolTipText = "Show on diagram"
		}
		referenceRhombusNode.Buttons = append(referenceRhombusNode.Buttons, rrVisibilityButton)

		plantDiagramNode.Children = append(plantDiagramNode.Children, referenceRhombusNode)
	}
	growthVectorShape := plantDiagram.GrowthVectorShape
	if growthVectorShape != nil {
		growthVectorShapeNode := &tree.Node{
			Name:            growthVectorShape.Name,
			IsNodeClickable: true,
		}
		growthVectorShapeNode.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(growthVectorShape, GetPointerToGongstructName[*GrowthVectorShape]())
			stager.stage.Commit()
		}
		gvVisibilityButton := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				growthVectorShape.SetIsHidden(!growthVectorShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if growthVectorShape.GetIsHidden() {
			gvVisibilityButton.Icon = string(buttons.BUTTON_visibility)
			gvVisibilityButton.Name = "Show"
			gvVisibilityButton.ToolTipText = "Show on diagram"
		}
		growthVectorShapeNode.Buttons = append(growthVectorShapeNode.Buttons, gvVisibilityButton)

		plantDiagramNode.Children = append(plantDiagramNode.Children, growthVectorShapeNode)
	}

	gridPathShape := plantDiagram.GridPathShape
	if gridPathShape != nil {
		gridPathShapeNode := &tree.Node{
			Name:            gridPathShape.Name,
			IsNodeClickable: true,
		}
		gridPathShapeNode.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(gridPathShape, GetPointerToGongstructName[*GridPathShape]())
			stager.stage.Commit()
		}
		gpVisibilityButton := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				gridPathShape.SetIsHidden(!gridPathShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if gridPathShape.GetIsHidden() {
			gpVisibilityButton.Icon = string(buttons.BUTTON_visibility)
			gpVisibilityButton.Name = "Show"
			gpVisibilityButton.ToolTipText = "Show on diagram"
		}
		gridPathShapeNode.Buttons = append(gridPathShapeNode.Buttons, gpVisibilityButton)

		plantDiagramNode.Children = append(plantDiagramNode.Children, gridPathShapeNode)
	}

	rotatedReferenceRhombus := plantDiagram.RotatedReferenceRhombus
	if rotatedReferenceRhombus != nil {
		node := &tree.Node{
			Name:            rotatedReferenceRhombus.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rotatedReferenceRhombus, GetPointerToGongstructName[*ReferenceRhombus]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				rotatedReferenceRhombus.SetIsHidden(!rotatedReferenceRhombus.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if rotatedReferenceRhombus.GetIsHidden() {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	rotatedGrowthVectorShape := plantDiagram.RotatedGrowthVectorShape
	if rotatedGrowthVectorShape != nil {
		node := &tree.Node{
			Name:            rotatedGrowthVectorShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rotatedGrowthVectorShape, GetPointerToGongstructName[*GrowthVectorShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				rotatedGrowthVectorShape.SetIsHidden(!rotatedGrowthVectorShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if rotatedGrowthVectorShape.GetIsHidden() {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	rotatedGridPathShape := plantDiagram.RotatedGridPathShape
	if rotatedGridPathShape != nil {
		node := &tree.Node{
			Name:            rotatedGridPathShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rotatedGridPathShape, GetPointerToGongstructName[*GridPathShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				rotatedGridPathShape.SetIsHidden(!rotatedGridPathShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if rotatedGridPathShape.GetIsHidden() {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}
}
