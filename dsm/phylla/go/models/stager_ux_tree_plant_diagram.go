package models

import (
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePlantDiagram(
	plant *Plant,
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

	axesShape := plant.AxesShape
	if axesShape != nil {
		node := &tree.Node{
			Name:            axesShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(axesShape, GetPointerToGongstructName[*AxesShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenAxesShape = !plantDiagram.IsHiddenAxesShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenAxesShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)

		handleBtn := &tree.Button{
			Name:            "Hide Handle",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide handles",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				axesShape.IsWithHiddenHandle = !axesShape.IsWithHiddenHandle
				stager.stage.Commit()
			},
		}
		if axesShape.IsWithHiddenHandle {
			handleBtn.Icon = string(buttons.BUTTON_visibility)
			handleBtn.Name = "Show Handle"
			handleBtn.ToolTipText = "Show handles"
		}
		node.Buttons = append(node.Buttons, handleBtn)

		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	referenceRhombus := plant.ReferenceRhombus
	if referenceRhombus != nil {
		node := &tree.Node{
			Name:            referenceRhombus.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(referenceRhombus, GetPointerToGongstructName[*RhombusShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenReferenceRhombus = !plantDiagram.IsHiddenReferenceRhombus
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenReferenceRhombus {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	plantCircumferenceShape := plant.PlantCircumferenceShape
	if plantCircumferenceShape != nil {
		node := &tree.Node{
			Name:            plantCircumferenceShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(plantCircumferenceShape, GetPointerToGongstructName[*PlantCircumferenceShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenPlantCircumferenceShape = !plantDiagram.IsHiddenPlantCircumferenceShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenPlantCircumferenceShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	gridPathShape := plant.GridPathShape
	if gridPathShape != nil {
		node := &tree.Node{
			Name:            gridPathShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(gridPathShape, GetPointerToGongstructName[*GridPathShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenGridPathShape = !plantDiagram.IsHiddenGridPathShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenGridPathShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	rhombusGridShape := plant.RhombusGridShape
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
				plantDiagram.IsHiddenRhombusGridShape = !plantDiagram.IsHiddenRhombusGridShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenRhombusGridShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	explanationTextShape := plant.ExplanationTextShape
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
				plantDiagram.IsHiddenExplanationTextShape = !plantDiagram.IsHiddenExplanationTextShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenExplanationTextShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	rotatedReferenceRhombus := plant.RotatedReferenceRhombus
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
				plantDiagram.IsHiddenRotatedReferenceRhombus = !plantDiagram.IsHiddenRotatedReferenceRhombus
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenRotatedReferenceRhombus {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	rotatedPlantCircumferenceShape := plant.RotatedPlantCircumferenceShape
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
				plantDiagram.IsHiddenRotatedPlantCircumferenceShape = !plantDiagram.IsHiddenRotatedPlantCircumferenceShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenRotatedPlantCircumferenceShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	rotatedGridPathShape := plant.RotatedGridPathShape
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
				plantDiagram.IsHiddenRotatedGridPathShape = !plantDiagram.IsHiddenRotatedGridPathShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenRotatedGridPathShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	rotatedRhombusGridShape := plant.RotatedRhombusGridShape
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
				plantDiagram.IsHiddenRotatedRhombusGridShape = !plantDiagram.IsHiddenRotatedRhombusGridShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenRotatedRhombusGridShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	growthPathRhombusGridShape := plant.GrowthPathRhombusGridShape
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
				plantDiagram.IsHiddenGrowthPathRhombusGridShape = !plantDiagram.IsHiddenGrowthPathRhombusGridShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenGrowthPathRhombusGridShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}
}
