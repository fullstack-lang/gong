package models

import (
	"fmt"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) onToggleVisibility(isHidden *bool, btn *tree.Button) func() {
	return func() {
		*isHidden = !*isHidden
		stager.stage.CommitWithSuspendedCallbacks()

		if *isHidden {
			btn.Icon = string(buttons.BUTTON_visibility)
			btn.Name = "Show"
			btn.ToolTipText = "Show on diagram"
		} else {
			btn.Icon = string(buttons.BUTTON_visibility_off)
			btn.Name = "Hide"
			btn.ToolTipText = "Hide from diagram"
		}

		stager.treeStage2D.Commit()
		stager.treeStage3D.Commit()
		stager.ux_svg_plant_diagram()

		// only regenerate the 3D stage when the user is actually looking at the 3D view
		plant := stager.GetCurrentPlant()
		if plant != nil && (plant.CurrentView == VIEW_VASE_3D || plant.CurrentView == VIEW_STOOL_3D) {
			stager.UpdateThreeJSStage()
		}
	}
}

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
		stager.stage.CommitWithSuspendedCallbacks()
	}
	btn := &tree.Button{
		Name:            "Hide",
		Icon:            string(buttons.BUTTON_visibility_off),
		ToolTipText:     "Hide from diagram",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
	}
	btn.OnClick = stager.onToggleVisibility(isHidden, btn)
	if *isHidden {
		btn.Icon = string(buttons.BUTTON_visibility)
		btn.Name = "Show"
		btn.ToolTipText = "Show on diagram"
	}
	node.Buttons = append(node.Buttons, btn)
	parentNode.Children = append(parentNode.Children, node)
	return node
}

// addHideAllButton adds a "hide all" button to a category node.
// The button only appears when at least one item in the category is currently visible.
// Because the tree is rebuilt on every commit, the button disappears automatically
// once all items are already hidden.
func (stager *Stager) addHideAllButton(categoryNode *tree.Node, hiddenPtrs ...*bool) {
	// anyVisible := false
	// for _, h := range hiddenPtrs {
	// 	if !*h {
	// 		anyVisible = true
	// 		break
	// 	}
	// }
	// if !anyVisible {
	// 	return
	// }

	captured := hiddenPtrs // capture slice for closure
	btn := &tree.Button{
		Name:            "Hide All",
		Icon:            string(buttons.BUTTON_visibility_off),
		ToolTipText:     "Hide all items in this category",
		HasToolTip:      true,
		ToolTipPosition: tree.Right,
		OnClick: func() {
			for _, h := range captured {
				*h = true
			}
			stager.stage.Commit()
		},
	}
	categoryNode.Buttons = append(categoryNode.Buttons, btn)
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
			for p := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
				p.IsSelected = (p == plant)
			}
			stager.selectedPlant = plant
			if plant.PlantType == Stool {
				if plant.CurrentView != VIEW_PLANT_2D && plant.CurrentView != VIEW_STOOL_3D {
					plant.CurrentView = VIEW_PLANT_2D
				}
			} else if plant.PlantType != Vase {
				plant.CurrentView = VIEW_PLANT_2D
			}
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
		for p := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
			p.IsSelected = (p == plant)
		}
		stager.selectedPlant = plant
		if plant.PlantType != Vase {
			plant.CurrentView = VIEW_PLANT_2D
		}
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
		}
		handleBtn.OnClick = stager.onToggleVisibility(&plant.AxesShape.IsWithHiddenHandle, handleBtn)
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
		stager.addHideAllButton(rhombusNodes,
			&plantDiagram.IsHiddenReferenceRhombus,
			&plantDiagram.IsHiddenPlantCircumferenceShape,
			&plantDiagram.IsHiddenGridPathShape,
			&plantDiagram.IsHiddenRhombusGridShape,
			&plantDiagram.IsHiddenExplanationTextShape,
			&plantDiagram.IsHiddenRotatedReferenceRhombus,
			&plantDiagram.IsHiddenRotatedPlantCircumferenceShape,
			&plantDiagram.IsHiddenRotatedGridPathShape,
			&plantDiagram.IsHiddenRotatedRhombusGridShape,
			&plantDiagram.IsHiddenGrowthPathRhombusGridShape,
		)

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
		appendDiagramNode(stager, arcNodes, "Growth Curve 2D (stack by growth vector)", plant.StackOfGrowthCurve2DByGrowthVector, &plantDiagram.IsHiddenStackOfGrowthCurve2DByGrowthVector)
		stager.addHideAllButton(arcNodes,
			&plantDiagram.IsHiddenGrowthVectorShape,
			&plantDiagram.IsHiddenPerpendicularVectorGrid,
			&plantDiagram.IsHiddenBaseVectorShapeGrid,
			&plantDiagram.IsHiddenArcNormalVectorShapeGrid,
			&plantDiagram.IsHiddenStartArcShapeGrid,
			&plantDiagram.IsHiddenMidArcVectorShapeGrid,
			&plantDiagram.IsHiddenEndArcShapeGrid,
			&plantDiagram.IsHiddenGrowthCurve2D,
			&plantDiagram.IsHiddenStackOfGrowthCurve2DByGrowthVector,
		)

		if plant.PlantType == Vase && plantDiagram.VaseDiagram != nil && plant.VaseAbstract != nil {
			vase := plant.VaseAbstract
			vaseDiagram := plantDiagram.VaseDiagram

			vaseArcNodes := &tree.Node{
				Name:            "Arc Confs Vase",
				IsExpanded:      vaseDiagram.IsVaseArcNodesExpanded,
				IsNodeClickable: true,
			}
			plantDiagramNode.Children = append(plantDiagramNode.Children, vaseArcNodes)
			vaseArcNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&vaseDiagram.IsVaseArcNodesExpanded)

			appendDiagramNode(stager, vaseArcNodes, "Top Growth Curve 2D", vase.TopGrowthCurve2D, &vaseDiagram.IsHiddenTopGrowthCurve2D)
			appendDiagramNode(stager, vaseArcNodes, "Top End Arc Grid", vase.TopEndArcShapeGrid, &vaseDiagram.IsHiddenTopEndArcShapeGrid)
			appendDiagramNode(stager, vaseArcNodes, "Top Mid Arc Vector Shape Grid", vase.TopMidArcVectorShapeGrid, &vaseDiagram.IsHiddenTopMidArcVectorShapeGrid)
			appendDiagramNode(stager, vaseArcNodes, "Top Start Arc Grid", vase.TopStartArcShapeGrid, &vaseDiagram.IsHiddenTopStartArcShapeGrid)
			appendDiagramNode(stager, vaseArcNodes, "Shifted Bottom Top Start Arc Grid", vase.ShiftedBottomTopStartArcShapeGrid, &vaseDiagram.IsHiddenShiftedBottomTopStartArcShapeGrid)
			appendDiagramNode(stager, vaseArcNodes, "Perpendicular Vector Grid Halfway", vase.PerpendicularVectorGridHalfway, &vaseDiagram.IsHiddenPerpendicularVectorGridHalfway)
			stager.addHideAllButton(vaseArcNodes,
				&vaseDiagram.IsHiddenTopGrowthCurve2D,
				&vaseDiagram.IsHiddenTopEndArcShapeGrid,
				&vaseDiagram.IsHiddenTopMidArcVectorShapeGrid,
				&vaseDiagram.IsHiddenTopStartArcShapeGrid,
				&vaseDiagram.IsHiddenShiftedBottomTopStartArcShapeGrid,
				&vaseDiagram.IsHiddenPerpendicularVectorGridHalfway,
			)

			vaseClampingNodes := &tree.Node{
				Name:            "Clamping Confs Vase",
				IsExpanded:      vaseDiagram.IsVaseClampingNodesExpanded,
				IsNodeClickable: true,
			}
			plantDiagramNode.Children = append(plantDiagramNode.Children, vaseClampingNodes)
			vaseClampingNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&vaseDiagram.IsVaseClampingNodesExpanded)

			appendDiagramNode(stager, vaseClampingNodes, "Start Halfway Arc Shape Grid", vase.StartHalfwayArcShapeGrid, &vaseDiagram.IsHiddenStartHalfwayArcShapeGrid)
			appendDiagramNode(stager, vaseClampingNodes, "Top Start Halfway Arc Shape Grid", vase.TopStartHalfwayArcShapeGrid, &vaseDiagram.IsHiddenTopStartHalfwayArcShapeGrid)
			appendDiagramNode(stager, vaseClampingNodes, "End Halfway Arc Shape Grid", vase.EndHalfwayArcShapeGrid, &vaseDiagram.IsHiddenEndHalfwayArcShapeGrid)
			appendDiagramNode(stager, vaseClampingNodes, "Top End Halfway Arc Shape Grid", vase.TopEndHalfwayArcShapeGrid, &vaseDiagram.IsHiddenTopEndHalfwayArcShapeGrid)

			appendDiagramNode(stager, vaseClampingNodes, "Stack Of Rotated Growth Curve 2D", vase.StackOfRotatedGrowthCurve2D, &vaseDiagram.IsHiddenStackOfGrowthCurve)
			appendDiagramNode(stager, vaseClampingNodes, "Top Stack Of Rotated Growth Curve 2D", vase.TopStackOfRotatedGrowthCurve2D, &vaseDiagram.IsHiddenTopStackOfGrowthCurve)
			appendDiagramNode(stager, vaseClampingNodes, "Stack Of Rotated Growth Curve 2D Ribbon", vase.StackOfRotatedGrowthCurve2DRibbon, &vaseDiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Growth Curve 2D Ribbon", vase.GrowthCurve2DRibbon, &vaseDiagram.IsHiddenGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Shifted Right Growth Curve 2D Ribbon", vase.ShiftedRightGrowthCurve2DRibbon, &vaseDiagram.IsHiddenShiftedRightGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Shifted Left Growth Curve 2D Ribbon", vase.ShiftedLeftGrowthCurve2DRibbon, &vaseDiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Ribbon", vase.PartiallyGrowthCurve2DRibbon, &vaseDiagram.IsHiddenPartiallyGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Shifted Left Partially Growth Curve 2D Ribbon", vase.ShiftedLeftPartiallyGrowthCurve2DRibbon, &vaseDiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon)
			appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Trajectory", vase.PartiallyGrowthCurve2DTrajectory, &vaseDiagram.IsHiddenPartiallyGrowthCurve2DTrajectory)
			appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Trajectory P1 P2", vase.PartiallyGrowthCurve2DTrajectoryP1P2, &vaseDiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2)
			appendDiagramNode(stager, vaseClampingNodes, "Px Shape", vase.PxShape, &vaseDiagram.IsHiddenPxShape)
			appendDiagramNode(stager, vaseClampingNodes, "Chosen P1 P2 Pair Shape", vase.ChosenP1P2PairShape, &vaseDiagram.IsHiddenChosenP1P2PairShape)
			appendDiagramNode(stager, vaseClampingNodes, "Key Hole", vase.KeyHoleShape, &vaseDiagram.IsHiddenKeyHoleShape)

			appendDiagramNode(stager, vaseClampingNodes, "Stack Of Growth Curve 2D", vase.StackOfGrowthCurve2D, &vaseDiagram.IsHiddenStackOfGrowthCurve2D)
			appendDiagramNode(stager, vaseClampingNodes, "Top Stack Of Growth Curve 2D", vase.TopStackOfGrowthCurve2D, &vaseDiagram.IsHiddenTopStackOfGrowthCurve2D)
			appendDiagramNode(stager, vaseClampingNodes, "Stack Of Growth Curve 2D Ribbon", vase.StackOfGrowthCurve2DRibbon, &vaseDiagram.IsHiddenStackOfGrowthCurve2DRibbon)
			stager.addHideAllButton(vaseClampingNodes,
				&vaseDiagram.IsHiddenStartHalfwayArcShapeGrid,
				&vaseDiagram.IsHiddenTopStartHalfwayArcShapeGrid,
				&vaseDiagram.IsHiddenEndHalfwayArcShapeGrid,
				&vaseDiagram.IsHiddenTopEndHalfwayArcShapeGrid,
				&vaseDiagram.IsHiddenStackOfGrowthCurve,
				&vaseDiagram.IsHiddenTopStackOfGrowthCurve,
				&vaseDiagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon,
				&vaseDiagram.IsHiddenGrowthCurve2DRibbon,
				&vaseDiagram.IsHiddenShiftedRightGrowthCurve2DRibbon,
				&vaseDiagram.IsHiddenShiftedLeftGrowthCurve2DRibbon,
				&vaseDiagram.IsHiddenPartiallyGrowthCurve2DRibbon,
				&vaseDiagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon,
				&vaseDiagram.IsHiddenPartiallyGrowthCurve2DTrajectory,
				&vaseDiagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2,
				&vaseDiagram.IsHiddenPxShape,
				&vaseDiagram.IsHiddenChosenP1P2PairShape,
				&vaseDiagram.IsHiddenKeyHoleShape,
				&vaseDiagram.IsHiddenStackOfGrowthCurve2D,
				&vaseDiagram.IsHiddenTopStackOfGrowthCurve2D,
				&vaseDiagram.IsHiddenStackOfGrowthCurve2DRibbon,
			)
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

	if is3DView && plantDiagram.StoolDiagram != nil {
		appendDiagramNode(stager, plantDiagramNode, "3D Sampled Points", plantDiagram.StoolDiagram.SampledPoints3DShape, &plantDiagram.StoolDiagram.IsHiddenSampledPoints3DShape)
	}
}
