package models

import (
	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func appendDiagramNode[T GongstructIF](
	stager *Stager,
	parentNode *tree.Node,
	name string,
	gongstructPointer T,
	isHidden *bool,
) *tree.Node {
	node := &tree.Node{
		Name:              name,
		IsNodeClickable:   true,
		HasCheckboxButton: false,
	}
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(gongstructPointer, GetPointerToGongstructName[T]())
		stager.stage.Commit()
	}
	btn := &tree.Button{
		Name:            "Hide",
		Icon:            string(buttons.BUTTON_visibility_off),
		ToolTipText:     "Hide from diagram",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
		OnClick: func() {
			*isHidden = !*isHidden
			stager.stage.Commit()
		},
	}
	if *isHidden {
		btn.Icon = string(buttons.BUTTON_visibility)
		btn.Name = "Show"
		btn.ToolTipText = "Show on diagram"
	}
	node.Buttons = append(node.Buttons, btn)
	parentNode.Children = append(parentNode.Children, node)
	return node
}

func (stager *Stager) treePlantDiagram(
	plant *Plant,
	plantDiagram *PlantDiagram,
	parentNodes *[]*tree.Node,
	is3DView bool,
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

	if !is3DView {
		axesNode := appendDiagramNode(stager, plantDiagramNode, "Axes", plant.AxesShape, &plantDiagram.IsHiddenAxesShape)

		handleBtn := &tree.Button{
			Name:            "Hide Handle",
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide handles",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				plant.AxesShape.IsWithHiddenHandle = !plant.AxesShape.IsWithHiddenHandle
				stager.stage.Commit()
			},
		}
		if plant.AxesShape.IsWithHiddenHandle {
			handleBtn.Icon = string(buttons.BUTTON_visibility)
			handleBtn.Name = "Show Handle"
			handleBtn.ToolTipText = "Show handles"
		}
		axesNode.Buttons = append(axesNode.Buttons, handleBtn)

		rhombusNodes := &tree.Node{
			Name:            "Rhombus confs",
			IsExpanded:      plantDiagram.IsRhombusNodesExpanded,
			IsNodeClickable: true,
		}
		plantDiagramNode.Children = append(plantDiagramNode.Children, rhombusNodes)
		rhombusNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plantDiagram.IsRhombusNodesExpanded)

		appendDiagramNode(stager, rhombusNodes, "Reference Rhombus", plant.RhombusStuff.ReferenceRhombus, &plantDiagram.IsHiddenReferenceRhombus)
		appendDiagramNode(stager, rhombusNodes, "Plant Circumference", plant.RhombusStuff.PlantCircumferenceShape, &plantDiagram.IsHiddenPlantCircumferenceShape)
		appendDiagramNode(stager, rhombusNodes, "Grid Path", plant.RhombusStuff.GridPathShape, &plantDiagram.IsHiddenGridPathShape)
		appendDiagramNode(stager, rhombusNodes, "Initial Rhombus Grid", plant.RhombusStuff.InitialRhombusGridShape, &plantDiagram.IsHiddenRhombusGridShape)
		appendDiagramNode(stager, rhombusNodes, "Explanation Text", plant.RhombusStuff.ExplanationTextShape, &plantDiagram.IsHiddenExplanationTextShape)
		appendDiagramNode(stager, rhombusNodes, "Rotated Reference Rhombus", plant.RhombusStuff.RotatedReferenceRhombus, &plantDiagram.IsHiddenRotatedReferenceRhombus)
		appendDiagramNode(stager, rhombusNodes, "Rotated Plant Circumference", plant.RhombusStuff.RotatedPlantCircumferenceShape, &plantDiagram.IsHiddenRotatedPlantCircumferenceShape)
		appendDiagramNode(stager, rhombusNodes, "Rotated Grid Path", plant.RhombusStuff.RotatedGridPathShape, &plantDiagram.IsHiddenRotatedGridPathShape)
		appendDiagramNode(stager, rhombusNodes, "Rotated Rhombus Grid 2", plant.RhombusStuff.RotatedRhombusGridShape2, &plantDiagram.IsHiddenRotatedRhombusGridShape)
		appendDiagramNode(stager, rhombusNodes, "Growth Curve Rhombus Grid", plant.RhombusStuff.GrowthCurveRhombusGridShape, &plantDiagram.IsHiddenGrowthPathRhombusGridShape)

		arcNodes := &tree.Node{
			Name:            "Arc confs",
			IsExpanded:      plantDiagram.IsArcNodesExpanded,
			IsNodeClickable: true,
		}
		plantDiagramNode.Children = append(plantDiagramNode.Children, arcNodes)
		arcNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plantDiagram.IsArcNodesExpanded)

		appendDiagramNode(stager, arcNodes, "Growth Vector", plant.GrowthVectorShape, &plantDiagram.IsHiddenGrowthVectorShape)
		appendDiagramNode(stager, arcNodes, "Perpendicular Vector Grid", plant.PerpendicularVectorGrid, &plantDiagram.IsHiddenPerpendicularVectorGrid)
		appendDiagramNode(stager, arcNodes, "Perpendicular Vector Grid Halfway", plant.PerpendicularVectorGridHalfway, &plantDiagram.IsHiddenPerpendicularVectorGridHalfway)
		appendDiagramNode(stager, arcNodes, "Base Vector Grid", plant.BaseVectorShapeGrid, &plantDiagram.IsHiddenBaseVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "Arc Normal Vector Grid", plant.ArcNormalVectorShapeGrid, &plantDiagram.IsHiddenArcNormalVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "Start Arc Grid", plant.StartArcShapeGrid, &plantDiagram.IsHiddenStartArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Top Start Arc Grid", plant.TopStartArcShapeGrid, &plantDiagram.IsHiddenTopStartArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Shifted Bottom Top Start Arc Grid", plant.ShiftedBottomTopStartArcShapeGrid, &plantDiagram.IsHiddenShiftedBottomTopStartArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Mid Arc Vector Shape Grid", plant.MidArcVectorShapeGrid, &plantDiagram.IsHiddenMidArcVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "Top Mid Arc Vector Shape Grid", plant.TopMidArcVectorShapeGrid, &plantDiagram.IsHiddenTopMidArcVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "End Arc Grid", plant.EndArcShapeGrid, &plantDiagram.IsHiddenEndArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Top End Arc Grid", plant.TopEndArcShapeGrid, &plantDiagram.IsHiddenTopEndArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Growth Curve 2D", plant.GrowthCurve2D, &plantDiagram.IsHiddenGrowthCurve2D)
		appendDiagramNode(stager, arcNodes, "Top Growth Curve 2D", plant.TopGrowthCurve2D, &plantDiagram.IsHiddenTopGrowthCurve2D)

		appendDiagramNode(stager, plantDiagramNode, "Start Halfway Arc Shape Grid", plant.StartHalfwayArcShapeGrid, &plantDiagram.IsHiddenStartHalfwayArcShapeGrid)
		appendDiagramNode(stager, plantDiagramNode, "Top Start Halfway Arc Shape Grid", plant.TopStartHalfwayArcShapeGrid, &plantDiagram.IsHiddenTopStartHalfwayArcShapeGrid)
		appendDiagramNode(stager, plantDiagramNode, "End Halfway Arc Shape Grid", plant.EndHalfwayArcShapeGrid, &plantDiagram.IsHiddenEndHalfwayArcShapeGrid)
		appendDiagramNode(stager, plantDiagramNode, "Top End Halfway Arc Shape Grid", plant.TopEndHalfwayArcShapeGrid, &plantDiagram.IsHiddenTopEndHalfwayArcShapeGrid)

		appendDiagramNode(stager, plantDiagramNode, "Stack Of Rotated Growth Curve 2D", plant.StackOfRotatedGrowthCurve2D, &plantDiagram.IsHiddenStackOfGrowthCurve)
		appendDiagramNode(stager, plantDiagramNode, "Top Stack Of Rotated Growth Curve 2D", plant.TopStackOfRotatedGrowthCurve2D, &plantDiagram.IsHiddenTopStackOfGrowthCurve)
		appendDiagramNode(stager, plantDiagramNode, "Stack Of Rotated Growth Curve 2D Ribbon", plant.StackOfRotatedGrowthCurve2DRibbon, &plantDiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon)

		appendDiagramNode(stager, plantDiagramNode, "Stack Of Growth Curve 2D", plant.StackOfGrowthCurve2D, &plantDiagram.IsHiddenStackOfGrowthCurve2D)
		appendDiagramNode(stager, plantDiagramNode, "Top Stack Of Growth Curve 2D", plant.TopStackOfGrowthCurve2D, &plantDiagram.IsHiddenTopStackOfGrowthCurve2D)
		appendDiagramNode(stager, plantDiagramNode, "Stack Of Growth Curve 2D Ribbon", plant.StackOfGrowthCurve2DRibbon, &plantDiagram.IsHiddenStackOfGrowthCurve2DRibbon)
		appendDiagramNode(stager, plantDiagramNode, "Stack Of Partially Rotated Growth Curve 2D Ribbon", plant.StackOfPartiallyRotatedGrowthCurve2DRibbon, &plantDiagram.IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon)
	}

	if is3DView {
		appendDiagramNode(stager, plantDiagramNode, "3D Torus Stack", plantDiagram.TorusStackShape, &plantDiagram.IsHiddenTorusStackShape)
		appendDiagramNode(stager, plantDiagramNode, "Vertical 3D Torus Stack", plantDiagram.VerticalTorusStackShape, &plantDiagram.IsHiddenVerticalTorusStackShape)
	}
}
