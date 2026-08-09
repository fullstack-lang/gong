package models

import (
	"fmt"

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
	plant *PlantAbstract,
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

	recordMovieBtn := &tree.Button{
		Name:            "Record Movie",
		Icon:            string(buttons.BUTTON_videocam),
		ToolTipText:     "Record movie frames from rot 0.0 to 1.0",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
		OnClick: func() {
			if stager.IsMovieRecording() {
				stager.StopMovieRecording()
			} else {
				stager.StartMovieRecording(plant, plantDiagram)
			}
		},
	}
	if stager.IsMovieRecording() {
		rotRatio := 0.0
		if plant.PlantType == Vase {
			rotRatio = plant.VaseAbstract.RotationRatio
		}
		recordMovieBtn.Name = "Stop Recording"
		recordMovieBtn.Icon = string(buttons.BUTTON_stop)
		recordMovieBtn.ToolTipText = fmt.Sprintf("Recording... frame %d (rot: %.3f)", stager.GetMovieRecordingFrameCount(), rotRatio)
	}
	plantDiagramNode.Buttons = append(plantDiagramNode.Buttons, recordMovieBtn)

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
		appendDiagramNode(stager, arcNodes, "Base Vector Grid", plant.BaseVectorShapeGrid, &plantDiagram.IsHiddenBaseVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "Arc Normal Vector Grid", plant.ArcNormalVectorShapeGrid, &plantDiagram.IsHiddenArcNormalVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "Start Arc Grid", plant.StartArcShapeGrid, &plantDiagram.IsHiddenStartArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Mid Arc Vector Shape Grid", plant.MidArcVectorShapeGrid, &plantDiagram.IsHiddenMidArcVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "End Arc Grid", plant.EndArcShapeGrid, &plantDiagram.IsHiddenEndArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Growth Curve 2D", plant.GrowthCurve2D, &plantDiagram.IsHiddenGrowthCurve2D)

		if plant.PlantType == Vase && plantDiagram.VaseDiagram != nil {
			vaseArcNodes := &tree.Node{
				Name:            "Arc Confs Vase",
				IsExpanded:      plantDiagram.IsVaseArcNodesExpanded,
				IsNodeClickable: true,
			}
			plantDiagramNode.Children = append(plantDiagramNode.Children, vaseArcNodes)
			vaseArcNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plantDiagram.IsVaseArcNodesExpanded)

			appendDiagramNode(stager, vaseArcNodes, "Top Growth Curve 2D", plant.TopGrowthCurve2D, &plantDiagram.VaseDiagram.IsHiddenTopGrowthCurve2D)
			appendDiagramNode(stager, vaseArcNodes, "Top End Arc Grid", plant.TopEndArcShapeGrid, &plantDiagram.VaseDiagram.IsHiddenTopEndArcShapeGrid)
			appendDiagramNode(stager, vaseArcNodes, "Top Mid Arc Vector Shape Grid", plant.TopMidArcVectorShapeGrid, &plantDiagram.VaseDiagram.IsHiddenTopMidArcVectorShapeGrid)
			appendDiagramNode(stager, vaseArcNodes, "Top Start Arc Grid", plant.TopStartArcShapeGrid, &plantDiagram.VaseDiagram.IsHiddenTopStartArcShapeGrid)
			appendDiagramNode(stager, vaseArcNodes, "Shifted Bottom Top Start Arc Grid", plant.ShiftedBottomTopStartArcShapeGrid, &plantDiagram.VaseDiagram.IsHiddenShiftedBottomTopStartArcShapeGrid)
			appendDiagramNode(stager, vaseArcNodes, "Perpendicular Vector Grid Halfway", plant.PerpendicularVectorGridHalfway, &plantDiagram.VaseDiagram.IsHiddenPerpendicularVectorGridHalfway)

			vaseClampingNodes := &tree.Node{
				Name:            "Clamping Confs Vase",
				IsExpanded:      plantDiagram.IsVaseClampingNodesExpanded,
				IsNodeClickable: true,
			}
			plantDiagramNode.Children = append(plantDiagramNode.Children, vaseClampingNodes)
			vaseClampingNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&plantDiagram.IsVaseClampingNodesExpanded)

			appendDiagramNode(stager, vaseClampingNodes, "Start Halfway Arc Shape Grid", plant.StartHalfwayArcShapeGrid, &plantDiagram.VaseDiagram.IsHiddenStartHalfwayArcShapeGrid)
			appendDiagramNode(stager, vaseClampingNodes, "Top Start Halfway Arc Shape Grid", plant.TopStartHalfwayArcShapeGrid, &plantDiagram.VaseDiagram.IsHiddenTopStartHalfwayArcShapeGrid)
			appendDiagramNode(stager, vaseClampingNodes, "End Halfway Arc Shape Grid", plant.EndHalfwayArcShapeGrid, &plantDiagram.VaseDiagram.IsHiddenEndHalfwayArcShapeGrid)
			appendDiagramNode(stager, vaseClampingNodes, "Top End Halfway Arc Shape Grid", plant.TopEndHalfwayArcShapeGrid, &plantDiagram.VaseDiagram.IsHiddenTopEndHalfwayArcShapeGrid)

			appendDiagramNode(stager, vaseClampingNodes, "Stack Of Rotated Growth Curve 2D", plant.StackOfRotatedGrowthCurve2D, &plantDiagram.VaseDiagram.IsHiddenStackOfGrowthCurve)
			appendDiagramNode(stager, vaseClampingNodes, "Top Stack Of Rotated Growth Curve 2D", plant.TopStackOfRotatedGrowthCurve2D, &plantDiagram.VaseDiagram.IsHiddenTopStackOfGrowthCurve)
			appendDiagramNode(stager, vaseClampingNodes, "Stack Of Rotated Growth Curve 2D Ribbon", plant.StackOfRotatedGrowthCurve2DRibbon, &plantDiagram.VaseDiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Growth Curve 2D Ribbon", plant.GrowthCurve2DRibbon, &plantDiagram.VaseDiagram.IsHiddenGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Shifted Right Growth Curve 2D Ribbon", plant.ShiftedRightGrowthCurve2DRibbon, &plantDiagram.VaseDiagram.IsHiddenShiftedRightGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Shifted Left Growth Curve 2D Ribbon", plant.ShiftedLeftGrowthCurve2DRibbon, &plantDiagram.VaseDiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Ribbon", plant.PartiallyGrowthCurve2DRibbon, &plantDiagram.VaseDiagram.IsHiddenPartiallyGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Shifted Left Partially Growth Curve 2D Ribbon", plant.ShiftedLeftPartiallyGrowthCurve2DRibbon, &plantDiagram.VaseDiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Trajectory", plant.PartiallyGrowthCurve2DTrajectory, &plantDiagram.VaseDiagram.IsHiddenPartiallyGrowthCurve2DTrajectory)
			appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Trajectory P1 P2", plant.PartiallyGrowthCurve2DTrajectoryP1P2, &plantDiagram.VaseDiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2)
			appendDiagramNode(stager, vaseClampingNodes, "Px Shape", plant.PxShape, &plantDiagram.VaseDiagram.IsHiddenPxShape)
			appendDiagramNode(stager, vaseClampingNodes, "Chosen P1 P2 Pair Shape", plant.ChosenP1P2PairShape, &plantDiagram.VaseDiagram.IsHiddenChosenP1P2PairShape)
			appendDiagramNode(stager, vaseClampingNodes, "Key Hole", plant.KeyHoleShape, &plantDiagram.VaseDiagram.IsHiddenKeyHoleShape)

			appendDiagramNode(stager, vaseClampingNodes, "Stack Of Growth Curve 2D", plant.StackOfGrowthCurve2D, &plantDiagram.VaseDiagram.IsHiddenStackOfGrowthCurve2D)
			appendDiagramNode(stager, vaseClampingNodes, "Top Stack Of Growth Curve 2D", plant.TopStackOfGrowthCurve2D, &plantDiagram.VaseDiagram.IsHiddenTopStackOfGrowthCurve2D)
			appendDiagramNode(stager, vaseClampingNodes, "Stack Of Growth Curve 2D Ribbon", plant.StackOfGrowthCurve2DRibbon, &plantDiagram.VaseDiagram.IsHiddenStackOfGrowthCurve2DRibbon)
		}
	}

	if is3DView && plantDiagram.VaseDiagram != nil {
		appendDiagramNode(stager, plantDiagramNode, "3D Torus Stack", plantDiagram.VaseDiagram.TorusStackShape, &plantDiagram.VaseDiagram.IsHiddenTorusStackShape)
		appendDiagramNode(stager, plantDiagramNode, "Vertical 3D Torus Stack", plantDiagram.VaseDiagram.VerticalTorusStackShape, &plantDiagram.VaseDiagram.IsHiddenVerticalTorusStackShape)
		appendDiagramNode(stager, plantDiagramNode, "Partially Rotated 3D Torus", plantDiagram.VaseDiagram.PartiallyRotatedTorusShape, &plantDiagram.VaseDiagram.IsHiddenPartiallyRotatedTorusShape)
		appendDiagramNode(stager, plantDiagramNode, "Stack Of Partially Rotated 3D Torus", plantDiagram.VaseDiagram.StackOfPartiallyRotatedTorusShape, &plantDiagram.VaseDiagram.IsHiddenStackOfPartiallyRotatedTorusShape)
		appendDiagramNode(stager, plantDiagramNode, "3D Key Hole", plantDiagram.VaseDiagram.KeyHole3DShape, &plantDiagram.VaseDiagram.IsHiddenKeyHole3DShape)
		appendDiagramNode(stager, plantDiagramNode, "3D Key", plantDiagram.VaseDiagram.Key3DShape, &plantDiagram.VaseDiagram.IsHiddenKey3DShape)
		appendDiagramNode(stager, plantDiagramNode, "3D Volume Key", plantDiagram.VaseDiagram.VolumeKey3DShape, &plantDiagram.VaseDiagram.IsHiddenVolumeKey3DShape)
		appendDiagramNode(stager, plantDiagramNode, "3D Torus Edge", plantDiagram.VaseDiagram.TorusEdge3DShape, &plantDiagram.VaseDiagram.IsHiddenTorusEdge3DShape)
		appendDiagramNode(stager, plantDiagramNode, "3D Points and lines between points", plantDiagram.VaseDiagram.PointsAndLines3DShape, &plantDiagram.VaseDiagram.IsHiddenPointsAndLines3DShape)
		appendDiagramNode(stager, plantDiagramNode, "3D Sampled Points", plantDiagram.VaseDiagram.SampledPoints3DShape, &plantDiagram.VaseDiagram.IsHiddenSampledPoints3DShape)
		appendDiagramNode(stager, plantDiagramNode, "3D Original Points", plantDiagram.VaseDiagram.OriginalPoints3DShape, &plantDiagram.VaseDiagram.IsHiddenOriginalPoints3DShape)
		appendDiagramNode(stager, plantDiagramNode, "3D Angle 0 Shape", plantDiagram.VaseDiagram.Angle0Shape, &plantDiagram.VaseDiagram.IsHiddenAngle0Shape)
	}
}
