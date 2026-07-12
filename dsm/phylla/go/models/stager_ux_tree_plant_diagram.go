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

	rhombusGridShape := plant.InitialRhombusGridShape
	if rhombusGridShape != nil {
		node := &tree.Node{
			Name:            rhombusGridShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rhombusGridShape, GetPointerToGongstructName[*InitialRhombusGridShape]())
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

	rotatedRhombusGridShape := plant.RotatedRhombusGridShape2
	if rotatedRhombusGridShape != nil {
		node := &tree.Node{
			Name:            rotatedRhombusGridShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(rotatedRhombusGridShape, GetPointerToGongstructName[*RotatedRhombusGridShape]())
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

	growthPathRhombusGridShape := plant.GrowthCurveRhombusGridShape
	if growthPathRhombusGridShape != nil {
		node := &tree.Node{
			Name:            growthPathRhombusGridShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(growthPathRhombusGridShape, GetPointerToGongstructName[*GrowthCurveRhombusGridShape]())
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

	growthVectorShape := plant.GrowthVectorShape
	if growthVectorShape != nil {
		node := &tree.Node{
			Name:            growthVectorShape.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(growthVectorShape, GetPointerToGongstructName[*GrowthVectorShape]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenGrowthVectorShape = !plantDiagram.IsHiddenGrowthVectorShape
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenGrowthVectorShape {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	perpendicularVectorGrid := plant.PerpendicularVectorGrid
	if perpendicularVectorGrid != nil {
		node := &tree.Node{
			Name:            perpendicularVectorGrid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(perpendicularVectorGrid, GetPointerToGongstructName[*PerpendicularVectorGrid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenPerpendicularVectorGrid = !plantDiagram.IsHiddenPerpendicularVectorGrid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenPerpendicularVectorGrid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	perpendicularVectorGridHalfway := plant.PerpendicularVectorGridHalfway
	if perpendicularVectorGridHalfway != nil {
		node := &tree.Node{
			Name:            perpendicularVectorGridHalfway.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(perpendicularVectorGridHalfway, GetPointerToGongstructName[*PerpendicularVectorGridHalfway]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenPerpendicularVectorGridHalfway = !plantDiagram.IsHiddenPerpendicularVectorGridHalfway
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenPerpendicularVectorGridHalfway {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	baseVectorShapeGrid := plant.BaseVectorShapeGrid
	if baseVectorShapeGrid != nil {
		node := &tree.Node{
			Name:            baseVectorShapeGrid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(baseVectorShapeGrid, GetPointerToGongstructName[*BaseVectorShapeGrid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenBaseVectorShapeGrid = !plantDiagram.IsHiddenBaseVectorShapeGrid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenBaseVectorShapeGrid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	arcNormalVectorShapeGrid := plant.ArcNormalVectorShapeGrid
	if arcNormalVectorShapeGrid != nil {
		node := &tree.Node{
			Name:            arcNormalVectorShapeGrid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(arcNormalVectorShapeGrid, GetPointerToGongstructName[*ArcNormalVectorShapeGrid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenArcNormalVectorShapeGrid = !plantDiagram.IsHiddenArcNormalVectorShapeGrid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenArcNormalVectorShapeGrid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	startArcShapeGrid := plant.StartArcShapeGrid
	if startArcShapeGrid != nil {
		node := &tree.Node{
			Name:            startArcShapeGrid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(startArcShapeGrid, GetPointerToGongstructName[*StartArcShapeGrid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenStartArcShapeGrid = !plantDiagram.IsHiddenStartArcShapeGrid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenStartArcShapeGrid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	startArcShapeV2Grid := plant.StartArcShapeV2Grid
	if startArcShapeV2Grid != nil {
		node := &tree.Node{
			Name:            startArcShapeV2Grid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(startArcShapeV2Grid, GetPointerToGongstructName[*StartArcShapeV2Grid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenStartArcShapeV2Grid = !plantDiagram.IsHiddenStartArcShapeV2Grid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenStartArcShapeV2Grid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	topStartArcShapeV2Grid := plant.TopStartArcShapeV2Grid
	if topStartArcShapeV2Grid != nil {
		node := &tree.Node{
			Name:            topStartArcShapeV2Grid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(topStartArcShapeV2Grid, GetPointerToGongstructName[*TopStartArcShapeV2Grid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenTopStartArcShapeV2Grid = !plantDiagram.IsHiddenTopStartArcShapeV2Grid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenTopStartArcShapeV2Grid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	endArcShapeGrid := plant.EndArcShapeGrid
	if endArcShapeGrid != nil {
		node := &tree.Node{
			Name:            endArcShapeGrid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(endArcShapeGrid, GetPointerToGongstructName[*EndArcShapeGrid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenEndArcShapeGrid = !plantDiagram.IsHiddenEndArcShapeGrid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenEndArcShapeGrid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	endArcShapeV2Grid := plant.EndArcShapeV2Grid
	if endArcShapeV2Grid != nil {
		node := &tree.Node{
			Name:            endArcShapeV2Grid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(endArcShapeV2Grid, GetPointerToGongstructName[*EndArcShapeV2Grid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenEndArcShapeV2Grid = !plantDiagram.IsHiddenEndArcShapeV2Grid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenEndArcShapeV2Grid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	topEndArcShapeV2Grid := plant.TopEndArcShapeV2Grid
	if topEndArcShapeV2Grid != nil {
		node := &tree.Node{
			Name:            topEndArcShapeV2Grid.Name,
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(topEndArcShapeV2Grid, GetPointerToGongstructName[*TopEndArcShapeV2Grid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenTopEndArcShapeV2Grid = !plantDiagram.IsHiddenTopEndArcShapeV2Grid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenTopEndArcShapeV2Grid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	if growthCurveBezierShapeGrid := plant.GrowthCurveBezierShapeGrid; growthCurveBezierShapeGrid != nil {
		node := &tree.Node{
			Name: "GrowthCurveBezierShapeGrid",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(growthCurveBezierShapeGrid, GetPointerToGongstructName[*GrowthCurveBezierShapeGrid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenGrowthCurveBezierShapeGrid = !plantDiagram.IsHiddenGrowthCurveBezierShapeGrid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenGrowthCurveBezierShapeGrid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	if stackOfGrowthCurve := plant.StackOfGrowthCurve; stackOfGrowthCurve != nil {
		node := &tree.Node{
			Name: "StackOfGrowthCurve",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(stackOfGrowthCurve, GetPointerToGongstructName[*StackOfGrowthCurve]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenStackOfGrowthCurve = !plantDiagram.IsHiddenStackOfGrowthCurve
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenStackOfGrowthCurve {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}
}
