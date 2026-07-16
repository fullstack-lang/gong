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
		HasToolTip:        true,
		ToolTipPosition:   tree.Right,
		ToolTipText:       "Check to select the diagram",
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
	rhombusNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plantDiagram.IsRhombusNodesExpanded)

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

	startArcShapeV2Grid := plant.StartArcShapeGrid
	{
		node := &tree.Node{
			Name:            "Start Arc Grid",
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(startArcShapeV2Grid, GetPointerToGongstructName[*StartArcShapeGrid]())
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

	topStartArcShapeV2Grid := plant.TopStartArcShapeGrid
	{
		node := &tree.Node{
			Name:            "Top Start Arc Grid",
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(topStartArcShapeV2Grid, GetPointerToGongstructName[*TopStartArcShapeGrid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenTopStartArcShapeGrid = !plantDiagram.IsHiddenTopStartArcShapeGrid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenTopStartArcShapeGrid {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	endArcShapeV2Grid := plant.EndArcShapeGrid
	{
		node := &tree.Node{
			Name:            "End Arc Grid",
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(endArcShapeV2Grid, GetPointerToGongstructName[*EndArcShapeGrid]())
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

	topEndArcShapeV2Grid := plant.TopEndArcShapeGrid
	{
		node := &tree.Node{
			Name:            "Top End Arc Grid",
			IsNodeClickable: true,
		}
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(topEndArcShapeV2Grid, GetPointerToGongstructName[*TopEndArcShapeGrid]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenTopEndArcShapeGrid = !plantDiagram.IsHiddenTopEndArcShapeGrid
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenTopEndArcShapeGrid {
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
		stackOfGrowthCurveV2 := plant.StackOfGrowthCurve
		node := &tree.Node{
			Name: "Stack Of Growth Curve",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(stackOfGrowthCurveV2, GetPointerToGongstructName[*StackOfGrowthCurve]())
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

	{
		topStackOfGrowthCurveV2 := plant.TopStackOfGrowthCurve
		node := &tree.Node{
			Name: "Top Stack Of Growth Curve",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(topStackOfGrowthCurveV2, GetPointerToGongstructName[*TopStackOfGrowthCurve]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenTopStackOfGrowthCurve = !plantDiagram.IsHiddenTopStackOfGrowthCurve
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenTopStackOfGrowthCurve {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, btn)
		plantDiagramNode.Children = append(plantDiagramNode.Children, node)
	}

	{
		shiftedLeftStackOfGrowthCurve := plant.ShiftedLeftStackOfGrowthCurve
		node := &tree.Node{
			Name: "Shifted Left Stack Of Growth Curve",
		}
		node.IsExpanded = true
		node.HasCheckboxButton = false
		node.IsNodeClickable = true
		node.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(shiftedLeftStackOfGrowthCurve, GetPointerToGongstructName[*ShiftedLeftStackOfGrowthCurve]())
			stager.stage.Commit()
		}
		btn := &tree.Button{
			Name:            "Hide",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plantDiagram.IsHiddenShiftedLeftStackOfGrowthCurve = !plantDiagram.IsHiddenShiftedLeftStackOfGrowthCurve
				stager.stage.Commit()
			},
		}
		if plantDiagram.IsHiddenShiftedLeftStackOfGrowthCurve {
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
