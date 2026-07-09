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
			stager.probeForm.FillUpFormFromGongstruct(referenceRhombus, GetPointerToGongstructName[*RhombusShape]())
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
	plantCircumferenceShape := plantDiagram.PlantCircumferenceShape
	if plantCircumferenceShape != nil {
		plantCircumferenceShapeNode := &tree.Node{
			Name:            plantCircumferenceShape.Name,
			IsNodeClickable: true,
		}
		plantCircumferenceShapeNode.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(plantCircumferenceShape, GetPointerToGongstructName[*PlantCircumferenceShape]())
			stager.stage.Commit()
		}
		gvVisibilityButton := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantCircumferenceShape.SetIsHidden(!plantCircumferenceShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if plantCircumferenceShape.GetIsHidden() {
			gvVisibilityButton.Icon = string(buttons.BUTTON_visibility)
			gvVisibilityButton.Name = "Show"
			gvVisibilityButton.ToolTipText = "Show on diagram"
		}
		plantCircumferenceShapeNode.Buttons = append(plantCircumferenceShapeNode.Buttons, gvVisibilityButton)

		plantDiagramNode.Children = append(plantDiagramNode.Children, plantCircumferenceShapeNode)
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

	rhombusGridShape := plantDiagram.RhombusGridShape
	if rhombusGridShape != nil {
		node := &tree.Node{
			Name:            rhombusGridShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rhombusGridShape, GetPointerToGongstructName[*RhombusGridShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				rhombusGridShape.SetIsHidden(!rhombusGridShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if rhombusGridShape.GetIsHidden() {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	explanationTextShape := plantDiagram.ExplanationTextShape
	if explanationTextShape != nil {
		node := &tree.Node{
			Name:            explanationTextShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(explanationTextShape, GetPointerToGongstructName[*ExplanationTextShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				explanationTextShape.SetIsHidden(!explanationTextShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if explanationTextShape.GetIsHidden() {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	rotatedReferenceRhombus := plantDiagram.RotatedReferenceRhombus
	if rotatedReferenceRhombus != nil {
		node := &tree.Node{
			Name:            rotatedReferenceRhombus.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rotatedReferenceRhombus, GetPointerToGongstructName[*RhombusShape]())
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

	rotatedPlantCircumferenceShape := plantDiagram.RotatedPlantCircumferenceShape
	if rotatedPlantCircumferenceShape != nil {
		node := &tree.Node{
			Name:            rotatedPlantCircumferenceShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rotatedPlantCircumferenceShape, GetPointerToGongstructName[*PlantCircumferenceShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				rotatedPlantCircumferenceShape.SetIsHidden(!rotatedPlantCircumferenceShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if rotatedPlantCircumferenceShape.GetIsHidden() {
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

	rotatedRhombusGridShape := plantDiagram.RotatedRhombusGridShape
	if rotatedRhombusGridShape != nil {
		node := &tree.Node{
			Name:            rotatedRhombusGridShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rotatedRhombusGridShape, GetPointerToGongstructName[*RhombusGridShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				rotatedRhombusGridShape.SetIsHidden(!rotatedRhombusGridShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if rotatedRhombusGridShape.GetIsHidden() {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	growthPathRhombusGridShape := plantDiagram.GrowthPathRhombusGridShape
	if growthPathRhombusGridShape != nil {
		node := &tree.Node{
			Name:            growthPathRhombusGridShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(growthPathRhombusGridShape, GetPointerToGongstructName[*RhombusGridShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				growthPathRhombusGridShape.SetIsHidden(!growthPathRhombusGridShape.GetIsHidden())
				stager.stage.Commit()
			},
		}
		if growthPathRhombusGridShape.GetIsHidden() {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

}
