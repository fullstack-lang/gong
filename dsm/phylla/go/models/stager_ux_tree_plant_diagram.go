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

	rhombusNodes := &tree.Node{
		Name:            "Rhombus confs",
		IsExpanded:      plantDiagram.IsRhombusNodesExpanded,
		IsNodeClickable: true,
	}
	plantDiagramNode.Children = append(plantDiagramNode.Children, rhombusNodes)
	rhombusNodes.OnIsCheckedChanged = stager.onIsExpandedChangeBool(&plantDiagram.IsRhombusNodesExpanded)

	axesShape := plant.AxesShape
	{
		node := &tree.Node{
			Name:            "Axes",
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
	{
		node := &tree.Node{
			Name:            "Reference Rhombus",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	plantCircumferenceShape := plant.PlantCircumferenceShape
	{
		node := &tree.Node{
			Name:            "Plant Circumference",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	gridPathShape := plant.GridPathShape
	{
		node := &tree.Node{
			Name:            "Grid Path",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	rhombusGridShape := plant.InitialRhombusGridShape
	{
		node := &tree.Node{
			Name:            "Initial Rhombus Grid",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	explanationTextShape := plant.ExplanationTextShape
	{
		node := &tree.Node{
			Name:            "Explanation Text",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	rotatedReferenceRhombus := plant.RotatedReferenceRhombus
	{
		node := &tree.Node{
			Name:            "Rotated Reference Rhombus",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	rotatedPlantCircumferenceShape := plant.RotatedPlantCircumferenceShape
	{
		node := &tree.Node{
			Name:            "Rotated Plant Circumference",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	rotatedGridPathShape := plant.RotatedGridPathShape
	{
		node := &tree.Node{
			Name:            "Rotated Grid Path",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	rotatedRhombusGridShape := plant.RotatedRhombusGridShape2
	{
		node := &tree.Node{
			Name:            "Rotated Rhombus Grid 2",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	growthPathRhombusGridShape := plant.GrowthCurveRhombusGridShape
	{
		node := &tree.Node{
			Name:            "Growth Curve Rhombus Grid",
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
		rhombusNodes.Children = append(rhombusNodes.Children, node)
	}

	growthVectorShape := plant.GrowthVectorShape
	{
		node := &tree.Node{
			Name:            "Growth Vector",
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
	{
		node := &tree.Node{
			Name:            "Perpendicular Vector Grid",
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
	{
		node := &tree.Node{
			Name:            "Perpendicular Vector Grid Halfway",
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
	{
		node := &tree.Node{
			Name:            "Base Vector Grid",
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
	{
		node := &tree.Node{
			Name:            "Arc Normal Vector Grid",
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

	startArcShapeV2Grid := plant.StartArcShapeV2Grid
	{
		node := &tree.Node{
			Name:            "Start Arc V2 Grid",
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
	{
		node := &tree.Node{
			Name:            "Top Start Arc V2 Grid",
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

	endArcShapeV2Grid := plant.EndArcShapeV2Grid
	{
		node := &tree.Node{
			Name:            "End Arc V2 Grid",
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
	{
		node := &tree.Node{
			Name:            "Top End Arc V2 Grid",
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

	bottomStartArcShapeV2Grid := plant.BottomStartArcShapeV2Grid
	{
		node := &tree.Node{
			Name:            "Bottom Start Arc V2 Grid",
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(bottomStartArcShapeV2Grid, GetPointerToGongstructName[*BottomStartArcShapeV2Grid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenBottomStartArcShapeV2Grid = !plantDiagram.IsHiddenBottomStartArcShapeV2Grid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenBottomStartArcShapeV2Grid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	bottomEndArcShapeV2Grid := plant.BottomEndArcShapeV2Grid
	{
		node := &tree.Node{
			Name:            "Bottom End Arc V2 Grid",
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(bottomEndArcShapeV2Grid, GetPointerToGongstructName[*BottomEndArcShapeV2Grid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenBottomEndArcShapeV2Grid = !plantDiagram.IsHiddenBottomEndArcShapeV2Grid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenBottomEndArcShapeV2Grid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	{
		growthCurveBezierShapeGrid := plant.GrowthCurveBezierShapeGrid
		node := &tree.Node{
			Name: "Growth Curve Bezier Grid",
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

	{
		stackOfGrowthCurveV2 := plant.StackOfGrowthCurveV2
		node := &tree.Node{
			Name: "Stack Of Growth Curve V2",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(stackOfGrowthCurveV2, GetPointerToGongstructName[*StackOfGrowthCurveV2]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenStackOfGrowthCurveV2 = !plantDiagram.IsHiddenStackOfGrowthCurveV2
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenStackOfGrowthCurveV2 {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	{
		topStackOfGrowthCurveV2 := plant.TopStackOfGrowthCurveV2
		node := &tree.Node{
			Name: "Top Stack Of Growth Curve V2",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(topStackOfGrowthCurveV2, GetPointerToGongstructName[*TopStackOfGrowthCurveV2]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenTopStackOfGrowthCurveV2 = !plantDiagram.IsHiddenTopStackOfGrowthCurveV2
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenTopStackOfGrowthCurveV2 {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	{
		bottomStackOfGrowthCurveV2 := plant.BottomStackOfGrowthCurveV2
		node := &tree.Node{
			Name: "Bottom Stack Of Growth Curve V2",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(bottomStackOfGrowthCurveV2, GetPointerToGongstructName[*BottomStackOfGrowthCurveV2]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenBottomStackOfGrowthCurveV2 = !plantDiagram.IsHiddenBottomStackOfGrowthCurveV2
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenBottomStackOfGrowthCurveV2 {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}
	{
		growthCurve2D := plant.GrowthCurve2D
		node := &tree.Node{
			Name: "Growth Curve 2D",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(growthCurve2D, GetPointerToGongstructName[*GrowthCurve2D]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenGrowthCurve2D = !plantDiagram.IsHiddenGrowthCurve2D
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenGrowthCurve2D {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	{
		topGrowthCurve2D := plant.TopGrowthCurve2D
		node := &tree.Node{
			Name: "Top Growth Curve 2D",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(topGrowthCurve2D, GetPointerToGongstructName[*TopGrowthCurve2D]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenTopGrowthCurve2D = !plantDiagram.IsHiddenTopGrowthCurve2D
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenTopGrowthCurve2D {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}
}
