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
		if plant != nil {
			if plant.CurrentView == VIEW_VASE_3D {
				stager.UpdateThreeJSStage()
			}
			if plant.CurrentView == VIEW_STOOL_3D {
				stager.UpdateStool3DStage()
			}
			if plant.CurrentView == VIEW_CLOCK_3D {
				stager.UpdateClock3DStage()
			}
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
func (stager *Stager) addHideAllButton(categoryNode *tree.Node, hiddenPtrs ...*bool) {
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

func (stager *Stager) handleDiagramCheck(diagramType interface{}, plant *PlantAbstract, view ViewType) {
	uncheckAllDiagrams(stager)
	// Actually we should set the passed diagram to checked, but it's done by the caller
	for p := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stager.stage) {
		p.IsSelected = (p == plant)
	}
	stager.selectedPlant = plant
	if plant.CurrentView != VIEW_ABOUT_SPIRAL_PLANTS {
		plant.CurrentView = view
	}
}

func (stager *Stager) treePlant2DDiagram(plant *PlantAbstract, diagram *Plant2DDiagram, parentNodes *[]*tree.Node, is3DView bool) {
	node := &tree.Node{
		Name: diagram.Name, IsExpanded: diagram.IsExpanded, IsNodeClickable: true,
		HasCheckboxButton: true, IsChecked: diagram.IsChecked, IsInEditMode: diagram.isInRenameMode,
		HasToolTip: true, ToolTipPosition: tree.Right, ToolTipText: "Check to select the diagram",
	}
	*parentNodes = append(*parentNodes, node)
	addRenameButton(diagram, node, stager)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			stager.handleDiagramCheck(diagram, plant, VIEW_PLANT_2D)
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
	node.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
	node.OnNameChange = stager.onNameChange(diagram)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(diagram, GetPointerToGongstructName[*Plant2DDiagram]())
		stager.handleDiagramCheck(diagram, plant, VIEW_PLANT_2D)
		diagram.IsChecked = true
		stager.stage.Commit()
	}

	if !is3DView {
		axesNode := appendDiagramNode(stager, node, "Axes", plant.AxesShape, &diagram.IsHiddenAxesShape)
		handleBtn := &tree.Button{Name: "Hide Handle", Icon: string(buttons.BUTTON_visibility_off), ToolTipText: "Hide handles", HasToolTip: true, ToolTipPosition: tree.Right}
		handleBtn.OnClick = stager.onToggleVisibility(&plant.AxesShape.IsWithHiddenHandle, handleBtn)
		if plant.AxesShape.IsWithHiddenHandle { handleBtn.Icon = string(buttons.BUTTON_visibility); handleBtn.Name = "Show Handle"; handleBtn.ToolTipText = "Show handles" }
		axesNode.Buttons = append(axesNode.Buttons, handleBtn)

		rhombusNodes := &tree.Node{Name: "Rhombus confs", IsExpanded: diagram.IsRhombusNodesExpanded, IsNodeClickable: true}
		node.Children = append(node.Children, rhombusNodes)
		rhombusNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsRhombusNodesExpanded)
		appendDiagramNode(stager, rhombusNodes, "Reference Rhombus", plant.RhombusStuff.ReferenceRhombus, &diagram.IsHiddenReferenceRhombus)
		appendDiagramNode(stager, rhombusNodes, "Plant Circumference", plant.RhombusStuff.PlantCircumferenceShape, &diagram.IsHiddenPlantCircumferenceShape)
		appendDiagramNode(stager, rhombusNodes, "Grid Path", plant.RhombusStuff.GridPathShape, &diagram.IsHiddenGridPathShape)
		appendDiagramNode(stager, rhombusNodes, "Initial Rhombus Grid", plant.RhombusStuff.InitialRhombusGridShape, &diagram.IsHiddenRhombusGridShape)
		appendDiagramNode(stager, rhombusNodes, "Explanation Text", plant.RhombusStuff.ExplanationTextShape, &diagram.IsHiddenExplanationTextShape)
		appendDiagramNode(stager, rhombusNodes, "Rotated Reference Rhombus", plant.RhombusStuff.RotatedReferenceRhombus, &diagram.IsHiddenRotatedReferenceRhombus)
		appendDiagramNode(stager, rhombusNodes, "Rotated Plant Circumference", plant.RhombusStuff.RotatedPlantCircumferenceShape, &diagram.IsHiddenRotatedPlantCircumferenceShape)
		appendDiagramNode(stager, rhombusNodes, "Rotated Grid Path", plant.RhombusStuff.RotatedGridPathShape, &diagram.IsHiddenRotatedGridPathShape)
		appendDiagramNode(stager, rhombusNodes, "Rotated Rhombus Grid 2", plant.RhombusStuff.RotatedRhombusGridShape2, &diagram.IsHiddenRotatedRhombusGridShape)
		appendDiagramNode(stager, rhombusNodes, "Growth Curve Rhombus Grid", plant.RhombusStuff.GrowthCurveRhombusGridShape, &diagram.IsHiddenGrowthPathRhombusGridShape)
		stager.addHideAllButton(rhombusNodes, &diagram.IsHiddenReferenceRhombus, &diagram.IsHiddenPlantCircumferenceShape, &diagram.IsHiddenGridPathShape, &diagram.IsHiddenRhombusGridShape, &diagram.IsHiddenExplanationTextShape, &diagram.IsHiddenRotatedReferenceRhombus, &diagram.IsHiddenRotatedPlantCircumferenceShape, &diagram.IsHiddenRotatedGridPathShape, &diagram.IsHiddenRotatedRhombusGridShape, &diagram.IsHiddenGrowthPathRhombusGridShape)

		arcNodes := &tree.Node{Name: "Arc confs", IsExpanded: diagram.IsArcNodesExpanded, IsNodeClickable: true}
		node.Children = append(node.Children, arcNodes)
		arcNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsArcNodesExpanded)
		appendDiagramNode(stager, arcNodes, "Growth Vector", plant.GrowthVectorShape, &diagram.IsHiddenGrowthVectorShape)
		appendDiagramNode(stager, arcNodes, "Perpendicular Vector Grid", plant.PerpendicularVectorGrid, &diagram.IsHiddenPerpendicularVectorGrid)
		appendDiagramNode(stager, arcNodes, "Base Vector Grid", plant.BaseVectorShapeGrid, &diagram.IsHiddenBaseVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "Arc Normal Vector Grid", plant.ArcNormalVectorShapeGrid, &diagram.IsHiddenArcNormalVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "Start Arc Grid", plant.StartArcShapeGrid, &diagram.IsHiddenStartArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Mid Arc Vector Shape Grid", plant.MidArcVectorShapeGrid, &diagram.IsHiddenMidArcVectorShapeGrid)
		appendDiagramNode(stager, arcNodes, "End Arc Grid", plant.EndArcShapeGrid, &diagram.IsHiddenEndArcShapeGrid)
		appendDiagramNode(stager, arcNodes, "Growth Curve 2D", plant.GrowthCurve2D, &diagram.IsHiddenGrowthCurve2D)
		appendDiagramNode(stager, arcNodes, "Growth Curve 2D (stack by growth vector)", plant.StackOfGrowthCurve2DByGrowthVector, &diagram.IsHiddenStackOfGrowthCurve2DByGrowthVector)
		stager.addHideAllButton(arcNodes, &diagram.IsHiddenGrowthVectorShape, &diagram.IsHiddenPerpendicularVectorGrid, &diagram.IsHiddenBaseVectorShapeGrid, &diagram.IsHiddenArcNormalVectorShapeGrid, &diagram.IsHiddenStartArcShapeGrid, &diagram.IsHiddenMidArcVectorShapeGrid, &diagram.IsHiddenEndArcShapeGrid, &diagram.IsHiddenGrowthCurve2D, &diagram.IsHiddenStackOfGrowthCurve2DByGrowthVector)
	}
}

func (stager *Stager) treeVase2DDiagram(plant *PlantAbstract, diagram *Vase2DDiagram, parentNodes *[]*tree.Node, is3DView bool) {
	node := &tree.Node{
		Name: diagram.Name, IsExpanded: diagram.IsExpanded, IsNodeClickable: true,
		HasCheckboxButton: true, IsChecked: diagram.IsChecked, IsInEditMode: diagram.isInRenameMode,
		HasToolTip: true, ToolTipPosition: tree.Right, ToolTipText: "Check to select the diagram",
	}
	*parentNodes = append(*parentNodes, node)
	addRenameButton(diagram, node, stager)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			stager.handleDiagramCheck(diagram, plant, VIEW_VASE_2D)
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
	node.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
	node.OnNameChange = stager.onNameChange(diagram)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(diagram, GetPointerToGongstructName[*Vase2DDiagram]())
		stager.handleDiagramCheck(diagram, plant, VIEW_VASE_2D)
		diagram.IsChecked = true
		stager.stage.Commit()
	}

	if !is3DView && plant.VaseAbstract != nil {
		vase := plant.VaseAbstract
		vaseArcNodes := &tree.Node{Name: "Arc Confs Vase", IsExpanded: diagram.IsVaseArcNodesExpanded, IsNodeClickable: true}
		node.Children = append(node.Children, vaseArcNodes)
		vaseArcNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsVaseArcNodesExpanded)
		appendDiagramNode(stager, vaseArcNodes, "Top Growth Curve 2D", vase.TopGrowthCurve2D, &diagram.IsHiddenTopGrowthCurve2D)
		appendDiagramNode(stager, vaseArcNodes, "Top End Arc Grid", vase.TopEndArcShapeGrid, &diagram.IsHiddenTopEndArcShapeGrid)
		appendDiagramNode(stager, vaseArcNodes, "Top Mid Arc Vector Shape Grid", vase.TopMidArcVectorShapeGrid, &diagram.IsHiddenTopMidArcVectorShapeGrid)
		appendDiagramNode(stager, vaseArcNodes, "Top Start Arc Grid", vase.TopStartArcShapeGrid, &diagram.IsHiddenTopStartArcShapeGrid)
		appendDiagramNode(stager, vaseArcNodes, "Shifted Bottom Top Start Arc Grid", vase.ShiftedBottomTopStartArcShapeGrid, &diagram.IsHiddenShiftedBottomTopStartArcShapeGrid)
		appendDiagramNode(stager, vaseArcNodes, "Perpendicular Vector Grid Halfway", vase.PerpendicularVectorGridHalfway, &diagram.IsHiddenPerpendicularVectorGridHalfway)
		stager.addHideAllButton(vaseArcNodes, &diagram.IsHiddenTopGrowthCurve2D, &diagram.IsHiddenTopEndArcShapeGrid, &diagram.IsHiddenTopMidArcVectorShapeGrid, &diagram.IsHiddenTopStartArcShapeGrid, &diagram.IsHiddenShiftedBottomTopStartArcShapeGrid, &diagram.IsHiddenPerpendicularVectorGridHalfway)

		vaseClampingNodes := &tree.Node{Name: "Clamping Confs Vase", IsExpanded: diagram.IsVaseClampingNodesExpanded, IsNodeClickable: true}
		node.Children = append(node.Children, vaseClampingNodes)
		vaseClampingNodes.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsVaseClampingNodesExpanded)
		appendDiagramNode(stager, vaseClampingNodes, "Start Halfway Arc Shape Grid", vase.StartHalfwayArcShapeGrid, &diagram.IsHiddenStartHalfwayArcShapeGrid)
		appendDiagramNode(stager, vaseClampingNodes, "Top Start Halfway Arc Shape Grid", vase.TopStartHalfwayArcShapeGrid, &diagram.IsHiddenTopStartHalfwayArcShapeGrid)
		appendDiagramNode(stager, vaseClampingNodes, "End Halfway Arc Shape Grid", vase.EndHalfwayArcShapeGrid, &diagram.IsHiddenEndHalfwayArcShapeGrid)
		appendDiagramNode(stager, vaseClampingNodes, "Top End Halfway Arc Shape Grid", vase.TopEndHalfwayArcShapeGrid, &diagram.IsHiddenTopEndHalfwayArcShapeGrid)
		appendDiagramNode(stager, vaseClampingNodes, "Stack Of Rotated Growth Curve 2D", vase.StackOfRotatedGrowthCurve2D, &diagram.IsHiddenStackOfGrowthCurve)
		appendDiagramNode(stager, vaseClampingNodes, "Top Stack Of Rotated Growth Curve 2D", vase.TopStackOfRotatedGrowthCurve2D, &diagram.IsHiddenTopStackOfGrowthCurve)
		appendDiagramNode(stager, vaseClampingNodes, "Stack Of Rotated Growth Curve 2D Ribbon", vase.StackOfRotatedGrowthCurve2DRibbon, &diagram.IsHiddenStackOfRotatedGrowthCurve2DRibbon)
		appendDiagramNode(stager, vaseClampingNodes, "Growth Curve 2D Ribbon", vase.GrowthCurve2DRibbon, &diagram.IsHiddenGrowthCurve2DRibbon)
		appendDiagramNode(stager, vaseClampingNodes, "ShiftedRight Growth Curve 2D Ribbon", vase.ShiftedRightGrowthCurve2DRibbon, &diagram.IsHiddenShiftedRightGrowthCurve2DRibbon)
		appendDiagramNode(stager, vaseClampingNodes, "ShiftedLeft Growth Curve 2D Ribbon", vase.ShiftedLeftGrowthCurve2DRibbon, &diagram.IsHiddenShiftedLeftGrowthCurve2DRibbon)
		appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Ribbon", vase.PartiallyGrowthCurve2DRibbon, &diagram.IsHiddenPartiallyGrowthCurve2DRibbon)
		appendDiagramNode(stager, vaseClampingNodes, "ShiftedLeft Partially Growth Curve 2D Ribbon", vase.ShiftedLeftPartiallyGrowthCurve2DRibbon, &diagram.IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon)
		appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Trajectory", vase.PartiallyGrowthCurve2DTrajectory, &diagram.IsHiddenPartiallyGrowthCurve2DTrajectory)
		appendDiagramNode(stager, vaseClampingNodes, "Partially Growth Curve 2D Trajectory P1 P2", vase.PartiallyGrowthCurve2DTrajectoryP1P2, &diagram.IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2)
		appendDiagramNode(stager, vaseClampingNodes, "Px Shape", vase.PxShape, &diagram.IsHiddenPxShape)
		appendDiagramNode(stager, vaseClampingNodes, "Chosen P1 P2 Pair Shape", vase.ChosenP1P2PairShape, &diagram.IsHiddenChosenP1P2PairShape)
		appendDiagramNode(stager, vaseClampingNodes, "Key Hole", vase.KeyHoleShape, &diagram.IsHiddenKeyHoleShape)
		appendDiagramNode(stager, vaseClampingNodes, "Stack Of Growth Curve 2D", vase.StackOfGrowthCurve2D, &diagram.IsHiddenStackOfGrowthCurve2D)
		appendDiagramNode(stager, vaseClampingNodes, "Top Stack Of Growth Curve 2D", vase.TopStackOfGrowthCurve2D, &diagram.IsHiddenTopStackOfGrowthCurve2D)
		appendDiagramNode(stager, vaseClampingNodes, "Stack Of Growth Curve 2D Ribbon", vase.StackOfGrowthCurve2DRibbon, &diagram.IsHiddenStackOfGrowthCurve2DRibbon)
	}
}

func (stager *Stager) treeVase3DDiagram(plant *PlantAbstract, diagram *Vase3DDiagram, parentNodes *[]*tree.Node, is3DView bool) {
	node := &tree.Node{
		Name: diagram.Name, IsExpanded: diagram.IsExpanded, IsNodeClickable: true,
		HasCheckboxButton: true, IsChecked: diagram.IsChecked, IsInEditMode: diagram.isInRenameMode,
		HasToolTip: true, ToolTipPosition: tree.Right, ToolTipText: "Check to select the diagram",
	}
	*parentNodes = append(*parentNodes, node)
	addRenameButton(diagram, node, stager)

	// Record Movie Button
	if is3DView && plant.CurrentView == VIEW_VASE_3D {
		recordMovieBtn := &tree.Button{
			Name: "Record Movie", Icon: string(buttons.BUTTON_videocam), ToolTipText: "Record movie frames from rot 0.0 to 1.0", HasToolTip: true, ToolTipPosition: tree.Right,
			OnClick: func() {
				if stager.IsMovieRecording() {
					stager.StopMovieRecording()
				} else {
					stager.StartMovieRecordingVase3D(plant, diagram)
				}
			},
		}
		if stager.IsMovieRecording() {
			rotRatio := 0.0
			if plant.PlantType == Vase && plant.VaseAbstract != nil {
				rotRatio = plant.VaseAbstract.RotationRatio
			}
			recordMovieBtn.Name = "Stop Recording"
			recordMovieBtn.Icon = string(buttons.BUTTON_stop)
			recordMovieBtn.ToolTipText = fmt.Sprintf("Recording... frame %d (rot: %.3f)", stager.GetMovieRecordingFrameCount(), rotRatio)
		}
		node.Buttons = append(node.Buttons, recordMovieBtn)
	}

	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			stager.handleDiagramCheck(diagram, plant, VIEW_VASE_3D)
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
	node.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
	node.OnNameChange = stager.onNameChange(diagram)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(diagram, GetPointerToGongstructName[*Vase3DDiagram]())
		stager.handleDiagramCheck(diagram, plant, VIEW_VASE_3D)
		diagram.IsChecked = true
		stager.stage.Commit()
	}

	if is3DView {
		appendDiagramNode(stager, node, "3D Torus Stack", diagram.TorusStackShape, &diagram.IsHiddenTorusStackShape)
		appendDiagramNode(stager, node, "Vertical 3D Torus Stack", diagram.VerticalTorusStackShape, &diagram.IsHiddenVerticalTorusStackShape)
		appendDiagramNode(stager, node, "Partially Rotated 3D Torus", diagram.PartiallyRotatedTorusShape, &diagram.IsHiddenPartiallyRotatedTorusShape)
		appendDiagramNode(stager, node, "Stack Of Partially Rotated 3D Torus", diagram.StackOfPartiallyRotatedTorusShape, &diagram.IsHiddenStackOfPartiallyRotatedTorusShape)
		appendDiagramNode(stager, node, "3D Key Hole", diagram.KeyHole3DShape, &diagram.IsHiddenKeyHole3DShape)
		appendDiagramNode(stager, node, "3D Key", diagram.Key3DShape, &diagram.IsHiddenKey3DShape)
		appendDiagramNode(stager, node, "3D Volume Key", diagram.VolumeKey3DShape, &diagram.IsHiddenVolumeKey3DShape)
		appendDiagramNode(stager, node, "3D Torus Edge", diagram.TorusEdge3DShape, &diagram.IsHiddenTorusEdge3DShape)
		appendDiagramNode(stager, node, "3D Points and lines between points", diagram.PointsAndLines3DShape, &diagram.IsHiddenPointsAndLines3DShape)
		appendDiagramNode(stager, node, "3D Sampled Points", diagram.SampledPoints3DShape, &diagram.IsHiddenSampledPoints3DShape)
		appendDiagramNode(stager, node, "3D Original Points", diagram.OriginalPoints3DShape, &diagram.IsHiddenOriginalPoints3DShape)
		appendDiagramNode(stager, node, "3D Angle 0 Shape", diagram.Angle0Shape, &diagram.IsHiddenAngle0Shape)
		appendDiagramNode(stager, node, "3D Tiled Floor", diagram.TiledFloor3DShape, &diagram.IsHiddenTiledFloor3DShape)
	}
}

func (stager *Stager) treeStool2DDiagram(plant *PlantAbstract, diagram *Stool2DDiagram, parentNodes *[]*tree.Node, is3DView bool) {
	node := &tree.Node{
		Name: diagram.Name, IsExpanded: diagram.IsExpanded, IsNodeClickable: true,
		HasCheckboxButton: true, IsChecked: diagram.IsChecked, IsInEditMode: diagram.isInRenameMode,
		HasToolTip: true, ToolTipPosition: tree.Right, ToolTipText: "Check to select the diagram",
	}
	*parentNodes = append(*parentNodes, node)
	addRenameButton(diagram, node, stager)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			stager.handleDiagramCheck(diagram, plant, VIEW_PLANT_2D)
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
	node.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
	node.OnNameChange = stager.onNameChange(diagram)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(diagram, GetPointerToGongstructName[*Stool2DDiagram]())
		stager.handleDiagramCheck(diagram, plant, VIEW_PLANT_2D)
		diagram.IsChecked = true
		stager.stage.Commit()
	}
}

func (stager *Stager) treeStool3DDiagram(plant *PlantAbstract, diagram *Stool3DDiagram, parentNodes *[]*tree.Node, is3DView bool) {
	node := &tree.Node{
		Name: diagram.Name, IsExpanded: diagram.IsExpanded, IsNodeClickable: true,
		HasCheckboxButton: true, IsChecked: diagram.IsChecked, IsInEditMode: diagram.isInRenameMode,
		HasToolTip: true, ToolTipPosition: tree.Right, ToolTipText: "Check to select the diagram",
	}
	*parentNodes = append(*parentNodes, node)
	addRenameButton(diagram, node, stager)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			stager.handleDiagramCheck(diagram, plant, VIEW_STOOL_3D)
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
	node.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
	node.OnNameChange = stager.onNameChange(diagram)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(diagram, GetPointerToGongstructName[*Stool3DDiagram]())
		stager.handleDiagramCheck(diagram, plant, VIEW_STOOL_3D)
		diagram.IsChecked = true
		stager.stage.Commit()
	}

	if is3DView {
		appendDiagramNode(stager, node, "3D Torus", diagram.Torus3DShape, &diagram.IsHiddenTorus3DShape)
		appendDiagramNode(stager, node, "Rotated 3D Torus", diagram.RotatedTorusShape, &diagram.IsHiddenRotatedTorusShape)
		appendDiagramNode(stager, node, "3D Sampled Points", diagram.SampledPoints3DShape, &diagram.IsHiddenSampledPoints3DShape)
		appendDiagramNode(stager, node, "Rotated 3D Sampled Points", diagram.RotatedSampledPoints3DShape, &diagram.IsHiddenRotatedSampledPoints3DShape)
		appendDiagramNode(stager, node, "3D Eye Sampled Points", diagram.EyeSampledPoints3DShape, &diagram.IsHiddenEyeSampledPoints3DShape)
		appendDiagramNode(stager, node, "3D Eye Corners Sampled Points", diagram.EyeCornersSampledPoints3DShape, &diagram.IsHiddenEyeCornersSampledPoints3DShape)
		appendDiagramNode(stager, node, "3D Eye", diagram.Eye3DShape, &diagram.IsHiddenEye3DShape)
		appendDiagramNode(stager, node, "Seat Top Curve", diagram.SeatTopCurveShape, &diagram.IsHiddenSeatTopCurveShape)
		appendDiagramNode(stager, node, "Rotated Seat Top Curve", diagram.RotatedSeatTopCurveShape, &diagram.IsHiddenRotatedSeatTopCurveShape)
		appendDiagramNode(stager, node, "Seat Bottom Curve", diagram.SeatBottomCurveShape, &diagram.IsHiddenSeatBottomCurveShape)
		appendDiagramNode(stager, node, "Rotated Seat Bottom Curve", diagram.RotatedSeatBottomCurveShape, &diagram.IsHiddenRotatedSeatBottomCurveShape)
		appendDiagramNode(stager, node, "Seat Bottom Eye Curve", diagram.EyeSeatBottomCurveShape, &diagram.IsHiddenEyeSeatBottomCurveShape)
		appendDiagramNode(stager, node, "Stool Bottom Eye Curve", diagram.EyeStoolBottomCurveShape, &diagram.IsHiddenEyeStoolBottomCurveShape)
		appendDiagramNode(stager, node, "Seat", diagram.Seat3DShape, &diagram.IsHiddenSeat3DShape)
		appendDiagramNode(stager, node, "Eye Volume", diagram.EyeVolume3DShape, &diagram.IsHiddenEyeVolume3DShape)
		appendDiagramNode(stager, node, "Seat and Legs", diagram.SeatAndLegs3DShape, &diagram.IsHiddenSeatAndLegs3DShape)
		appendDiagramNode(stager, node, "Rotated Seat and Legs", diagram.RotatedSeatAndLegs3DShape, &diagram.IsHiddenRotatedSeatAndLegs3DShape)
		appendDiagramNode(stager, node, "3D Tiled Floor", diagram.TiledFloor3DShape, &diagram.IsHiddenTiledFloor3DShape)
	}
}

func (stager *Stager) treeClock2DDiagram(plant *PlantAbstract, diagram *Clock2DDiagram, parentNodes *[]*tree.Node, is3DView bool) {
	node := &tree.Node{
		Name: diagram.Name, IsExpanded: diagram.IsExpanded, IsNodeClickable: true,
		HasCheckboxButton: true, IsChecked: diagram.IsChecked, IsInEditMode: diagram.isInRenameMode,
		HasToolTip: true, ToolTipPosition: tree.Right, ToolTipText: "Check to select the diagram",
	}
	*parentNodes = append(*parentNodes, node)
	addRenameButton(diagram, node, stager)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			stager.handleDiagramCheck(diagram, plant, VIEW_PLANT_2D)
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
	node.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
	node.OnNameChange = stager.onNameChange(diagram)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(diagram, GetPointerToGongstructName[*Clock2DDiagram]())
		stager.handleDiagramCheck(diagram, plant, VIEW_PLANT_2D)
		diagram.IsChecked = true
		stager.stage.Commit()
	}
}

func (stager *Stager) treeClock3DDiagram(plant *PlantAbstract, diagram *Clock3DDiagram, parentNodes *[]*tree.Node, is3DView bool) {
	node := &tree.Node{
		Name: diagram.Name, IsExpanded: diagram.IsExpanded, IsNodeClickable: true,
		HasCheckboxButton: true, IsChecked: diagram.IsChecked, IsInEditMode: diagram.isInRenameMode,
		HasToolTip: true, ToolTipPosition: tree.Right, ToolTipText: "Check to select the diagram",
	}
	*parentNodes = append(*parentNodes, node)
	addRenameButton(diagram, node, stager)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			stager.handleDiagramCheck(diagram, plant, VIEW_CLOCK_3D)
			diagram.IsChecked = true
			stager.stage.Commit()
		} else {
			diagram.IsChecked = false
			stager.stage.Commit()
		}
	}
	node.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsExpanded)
	node.OnNameChange = stager.onNameChange(diagram)
	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(diagram, GetPointerToGongstructName[*Clock3DDiagram]())
		stager.handleDiagramCheck(diagram, plant, VIEW_CLOCK_3D)
		diagram.IsChecked = true
		stager.stage.Commit()
	}

	if is3DView {
		appendDiagramNode(stager, node, "3D Torus", diagram.Torus3DShape, &diagram.IsHiddenTorus3DShape)
		appendDiagramNode(stager, node, "3D Sampled Points", diagram.SampledPoints3DShape, &diagram.IsHiddenSampledPoints3DShape)
		appendDiagramNode(stager, node, "Clock Top Curve", diagram.ClockTopCurveShape, &diagram.IsHiddenClockTopCurveShape)
		appendDiagramNode(stager, node, "3D Tiled Floor", diagram.TiledFloor3DShape, &diagram.IsHiddenTiledFloor3DShape)
	}
}
