package models

func (stager *Stager) enforceOrphanShapeRemove() (needCommit bool) {
	stage := stager.stage

	// Sets of referenced shapes
	refAxes := make(map[*AxesShape]bool)
	refPlantCirc := make(map[*PlantCircumferenceShape]bool)
	refGridPath := make(map[*GridPathShape]bool)
	refCircleGrid := make(map[*CircleGridShape]bool)
	refExplanation := make(map[*ExplanationTextShape]bool)
	refGrowthVector := make(map[*GrowthVectorShape]bool)

	refInitialGrid := make(map[*InitialRhombusGridShape]bool)
	refRotatedGrid := make(map[*RotatedRhombusGridShape]bool)
	refGrowthCurveGrid := make(map[*GrowthCurveRhombusGridShape]bool)
	refVectorGrid := make(map[*PerpendicularVectorGrid]bool)

	refInitialShape := make(map[*InitialRhombusShape]bool)
	refRhombusShape := make(map[*RhombusShape]bool)
	refRotatedShape := make(map[*RotatedRhombusShape]bool)
	refGrowthCurveShape := make(map[*GrowthCurveRhombusShape]bool)
	refPerpendicularVector := make(map[*PerpendicularVector]bool)
	refPerpendicularVectorHalfway := make(map[*PerpendicularVectorHalfway]bool)

	refPerpendicularVectorGridHalfway := make(map[*PerpendicularVectorGridHalfway]bool)
	refBaseVectorShapeGrid := make(map[*BaseVectorShapeGrid]bool)
	refArcNormalVectorShapeGrid := make(map[*ArcNormalVectorShapeGrid]bool)
	refStartArcShapeV2Grid := make(map[*StartArcShapeGrid]bool)
	refTopStartArcShapeV2Grid := make(map[*TopStartArcShapeGrid]bool)
	refEndArcShapeV2Grid := make(map[*EndArcShapeGrid]bool)
	refTopEndArcShapeV2Grid := make(map[*TopEndArcShapeGrid]bool)

	refBaseVectorShape := make(map[*BaseVectorShape]bool)
	refStackOfGrowthCurveV2 := make(map[*StackOfRotatedGrowthCurve2D]bool)
	refStackGrowthCurveStartArcShapeV2 := make(map[*StackRotatedGrowthCurve2DStartArcShape]bool)
	refStackGrowthCurveEndArcShapeV2 := make(map[*StackRotatedGrowthCurve2DEndArcShape]bool)

	refTopStackOfGrowthCurveV2 := make(map[*TopStackOfRotatedGrowthCurve2D]bool)
	refTopStackGrowthCurveStartArcShapeV2 := make(map[*TopStackOfRotatedGrowthCurve2DStartArcShape]bool)
	refTopStackGrowthCurveEndArcShapeV2 := make(map[*TopStackOfRotatedGrowthCurve2DEndArcShape]bool)
	refShiftedBottomTopStartArcShapeGrid := make(map[*ShiftedBottomTopStartArcShapeGrid]bool)
	refShiftedBottomTopStartArcShape := make(map[*ShiftedBottomTopStartArcShape]bool)
	refMidArcVectorShapeGrid := make(map[*MidArcVectorShapeGrid]bool)
	refMidArcVectorShape := make(map[*MidArcVectorShape]bool)
	refTopMidArcVectorShapeGrid := make(map[*TopMidArcVectorShapeGrid]bool)
	refTopMidArcVectorShape := make(map[*TopMidArcVectorShape]bool)
	refShiftedLeftStackOfNormalVector := make(map[*ShiftedLeftStackOfNormalVector]bool)
	refShiftedLeftStackNormalVector := make(map[*ShiftedLeftStackNormalVector]bool)
	refShiftedLeftStackOfGrowthCurve := make(map[*ShiftedLeftStackOfGrowthCurve]bool)
	refShiftedLeftStackGrowthCurveStartArcShape := make(map[*ShiftedLeftStackGrowthCurveStartArcShape]bool)
	refShiftedLeftStackGrowthCurveEndArcShape := make(map[*ShiftedLeftStackGrowthCurveEndArcShape]bool)
	refRendered3DShape := make(map[*Rendered3DShape]bool)
	refArcNormalVectorShape := make(map[*ArcNormalVectorShape]bool)
	refStartArcShapeV2 := make(map[*StartArcShape]bool)
	refTopStartArcShapeV2 := make(map[*TopStartArcShape]bool)
	refEndArcShapeV2 := make(map[*EndArcShape]bool)
	refTopEndArcShapeV2 := make(map[*TopEndArcShape]bool)

	refStackOfGrowthCurve2D := make(map[*StackOfGrowthCurve2D]bool)
	refStackGrowthCurve2DStartHalfwayArcShape := make(map[*StackGrowthCurve2DStartHalfwayArcShape]bool)
	refStackGrowthCurve2DEndHalfwayArcShape := make(map[*StackGrowthCurve2DEndHalfwayArcShape]bool)

	refTopStackOfGrowthCurve2D := make(map[*TopStackOfGrowthCurve2D]bool)
	refTopStackGrowthCurve2DStartHalfwayArcShape := make(map[*TopStackGrowthCurve2DStartHalfwayArcShape]bool)
	refTopStackGrowthCurve2DEndHalfwayArcShape := make(map[*TopStackGrowthCurve2DEndHalfwayArcShape]bool)

	refStackOfGrowthCurve2DRibbon := make(map[*StackOfGrowthCurve2DRibbon]bool)
	refStackOfRotatedGrowthCurve2DRibbon := make(map[*StackOfRotatedGrowthCurve2DRibbon]bool)
	refStackGrowthCurve2DRibbonStartShape := make(map[*StackGrowthCurve2DRibbonStartShape]bool)
	refStackGrowthCurve2DRibbonEndShape := make(map[*StackGrowthCurve2DRibbonEndShape]bool)
	refStackRotatedGrowthCurve2DRibbonStartShape := make(map[*StackRotatedGrowthCurve2DRibbonStartShape]bool)
	refStackRotatedGrowthCurve2DRibbonEndShape := make(map[*StackRotatedGrowthCurve2DRibbonEndShape]bool)

	refTorusStackShape := make(map[*TorusStackShape]bool)
	refVerticalTorusStackShape := make(map[*VerticalTorusStackShape]bool)

	// Collect referenced shapes from all plants
	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		if plant.AxesShape != nil {
			refAxes[plant.AxesShape] = true
		}
		if plant.RhombusStuff != nil {
			if plant.RhombusStuff.ReferenceRhombus != nil {
				refRhombusShape[plant.RhombusStuff.ReferenceRhombus] = true
			}
			if plant.RhombusStuff.PlantCircumferenceShape != nil {
				refPlantCirc[plant.RhombusStuff.PlantCircumferenceShape] = true
			}
			if plant.RhombusStuff.GridPathShape != nil {
				refGridPath[plant.RhombusStuff.GridPathShape] = true
			}
			if plant.RhombusStuff.ExplanationTextShape != nil {
				refExplanation[plant.RhombusStuff.ExplanationTextShape] = true
			}
			if plant.RhombusStuff.RotatedReferenceRhombus != nil {
				refRhombusShape[plant.RhombusStuff.RotatedReferenceRhombus] = true
			}
			if plant.RhombusStuff.RotatedPlantCircumferenceShape != nil {
				refPlantCirc[plant.RhombusStuff.RotatedPlantCircumferenceShape] = true
			}
			if plant.RhombusStuff.RotatedGridPathShape != nil {
				refGridPath[plant.RhombusStuff.RotatedGridPathShape] = true
			}

			if plant.RhombusStuff.InitialRhombusGridShape != nil {
				refInitialGrid[plant.RhombusStuff.InitialRhombusGridShape] = true
				for _, shape := range plant.RhombusStuff.InitialRhombusGridShape.InitialRhombusShapes {
					if shape != nil {
						refInitialShape[shape] = true
					}
				}
			}

			if plant.RhombusStuff.RotatedRhombusGridShape2 != nil {
				refRotatedGrid[plant.RhombusStuff.RotatedRhombusGridShape2] = true
				for _, shape := range plant.RhombusStuff.RotatedRhombusGridShape2.RotatedRhombusShapes {
					if shape != nil {
						refRotatedShape[shape] = true
					}
				}
			}

			if plant.RhombusStuff.GrowthCurveRhombusGridShape != nil {
				refGrowthCurveGrid[plant.RhombusStuff.GrowthCurveRhombusGridShape] = true
				for _, shape := range plant.RhombusStuff.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes {
					if shape != nil {
						refGrowthCurveShape[shape] = true
					}
				}
			}
		}

		if plant.GrowthVectorShape != nil {
			refGrowthVector[plant.GrowthVectorShape] = true
		}

		if plant.PerpendicularVectorGrid != nil {
			refVectorGrid[plant.PerpendicularVectorGrid] = true
			for _, vec := range plant.PerpendicularVectorGrid.PerpendicularVectors {
				if vec != nil {
					refPerpendicularVector[vec] = true
				}
			}
		}

		if plant.PerpendicularVectorGridHalfway != nil {
			refPerpendicularVectorGridHalfway[plant.PerpendicularVectorGridHalfway] = true
			for _, vec := range plant.PerpendicularVectorGridHalfway.PerpendicularVectorHalfways {
				if vec != nil {
					refPerpendicularVectorHalfway[vec] = true
				}
			}
		}

		if plant.BaseVectorShapeGrid != nil {
			refBaseVectorShapeGrid[plant.BaseVectorShapeGrid] = true
			for _, shape := range plant.BaseVectorShapeGrid.BaseVectorShapes {
				if shape != nil {
					refBaseVectorShape[shape] = true
				}
			}
		}

		if plant.ArcNormalVectorShapeGrid != nil {
			refArcNormalVectorShapeGrid[plant.ArcNormalVectorShapeGrid] = true
			for _, shape := range plant.ArcNormalVectorShapeGrid.ArcNormalVectorShapes {
				if shape != nil {
					refArcNormalVectorShape[shape] = true
				}
			}
		}

		if plant.StartArcShapeGrid != nil {
			refStartArcShapeV2Grid[plant.StartArcShapeGrid] = true
			for _, shape := range plant.StartArcShapeGrid.StartArcShapes {
				if shape != nil {
					refStartArcShapeV2[shape] = true
				}
			}
		}

		if plant.TopStartArcShapeGrid != nil {
			refTopStartArcShapeV2Grid[plant.TopStartArcShapeGrid] = true
			for _, shape := range plant.TopStartArcShapeGrid.TopStartArcShapes {
				if shape != nil {
					refTopStartArcShapeV2[shape] = true
				}
			}
		}

		if plant.EndArcShapeGrid != nil {
			refEndArcShapeV2Grid[plant.EndArcShapeGrid] = true
			for _, shape := range plant.EndArcShapeGrid.EndArcShapes {
				if shape != nil {
					refEndArcShapeV2[shape] = true
				}
			}
		}

		if plant.TopEndArcShapeGrid != nil {
			refTopEndArcShapeV2Grid[plant.TopEndArcShapeGrid] = true
			for _, shape := range plant.TopEndArcShapeGrid.TopEndArcShapes {
				if shape != nil {
					refTopEndArcShapeV2[shape] = true
				}
			}
		}

		if plant.StackOfRotatedGrowthCurve2D != nil {
			refStackOfGrowthCurveV2[plant.StackOfRotatedGrowthCurve2D] = true
			for _, shape := range plant.StackOfRotatedGrowthCurve2D.StackRotatedGrowthCurve2DStartArcShapes {
				if shape != nil {
					refStackGrowthCurveStartArcShapeV2[shape] = true
				}
			}
			for _, shape := range plant.StackOfRotatedGrowthCurve2D.StackRotatedGrowthCurve2DEndArcShapes {
				if shape != nil {
					refStackGrowthCurveEndArcShapeV2[shape] = true
				}
			}
		}
		if plant.TopStackOfRotatedGrowthCurve2D != nil {
			refTopStackOfGrowthCurveV2[plant.TopStackOfRotatedGrowthCurve2D] = true
			for _, shape := range plant.TopStackOfRotatedGrowthCurve2D.TopStackOfRotatedGrowthCurve2DStartArcShapes {
				if shape != nil {
					refTopStackGrowthCurveStartArcShapeV2[shape] = true
				}
			}
			for _, shape := range plant.TopStackOfRotatedGrowthCurve2D.TopStackOfRotatedGrowthCurve2DEndArcShapes {
				if shape != nil {
					refTopStackGrowthCurveEndArcShapeV2[shape] = true
				}
			}
		}

		if plant.StackOfGrowthCurve2D != nil {
			refStackOfGrowthCurve2D[plant.StackOfGrowthCurve2D] = true
			for _, shape := range plant.StackOfGrowthCurve2D.StackGrowthCurve2DStartHalfwayArcShapes {
				if shape != nil {
					refStackGrowthCurve2DStartHalfwayArcShape[shape] = true
				}
			}
			for _, shape := range plant.StackOfGrowthCurve2D.StackGrowthCurve2DEndHalfwayArcShapes {
				if shape != nil {
					refStackGrowthCurve2DEndHalfwayArcShape[shape] = true
				}
			}
		}

		if plant.TopStackOfGrowthCurve2D != nil {
			refTopStackOfGrowthCurve2D[plant.TopStackOfGrowthCurve2D] = true
			for _, start := range plant.TopStackOfGrowthCurve2D.TopStackGrowthCurve2DStartHalfwayArcShapes {
				refTopStackGrowthCurve2DStartHalfwayArcShape[start] = true
			}
			for _, end := range plant.TopStackOfGrowthCurve2D.TopStackGrowthCurve2DEndHalfwayArcShapes {
				refTopStackGrowthCurve2DEndHalfwayArcShape[end] = true
			}
		}

		if plant.StackOfGrowthCurve2DRibbon != nil {
			refStackOfGrowthCurve2DRibbon[plant.StackOfGrowthCurve2DRibbon] = true
			for _, start := range plant.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonStartShapes {
				refStackGrowthCurve2DRibbonStartShape[start] = true
			}
			for _, end := range plant.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonEndShapes {
				refStackGrowthCurve2DRibbonEndShape[end] = true
			}
		}

		if plant.StackOfRotatedGrowthCurve2DRibbon != nil {
			refStackOfRotatedGrowthCurve2DRibbon[plant.StackOfRotatedGrowthCurve2DRibbon] = true
			for _, start := range plant.StackOfRotatedGrowthCurve2DRibbon.StackRotatedGrowthCurve2DRibbonStartShapes {
				refStackRotatedGrowthCurve2DRibbonStartShape[start] = true
			}
			for _, end := range plant.StackOfRotatedGrowthCurve2DRibbon.StackRotatedGrowthCurve2DRibbonEndShapes {
				refStackRotatedGrowthCurve2DRibbonEndShape[end] = true
			}
		}

		if plant.ShiftedBottomTopStartArcShapeGrid != nil {
			refShiftedBottomTopStartArcShapeGrid[plant.ShiftedBottomTopStartArcShapeGrid] = true
			for _, shape := range plant.ShiftedBottomTopStartArcShapeGrid.ShiftedBottomTopStartArcShapes {
				if shape != nil {
					refShiftedBottomTopStartArcShape[shape] = true
				}
			}
		}
		if plant.MidArcVectorShapeGrid != nil {
			refMidArcVectorShapeGrid[plant.MidArcVectorShapeGrid] = true
			for _, shape := range plant.MidArcVectorShapeGrid.MidArcVectorShapes {
				if shape != nil {
					refMidArcVectorShape[shape] = true
				}
			}
		}
		if plant.TopMidArcVectorShapeGrid != nil {
			refTopMidArcVectorShapeGrid[plant.TopMidArcVectorShapeGrid] = true
			for _, shape := range plant.TopMidArcVectorShapeGrid.TopMidArcVectorShapes {
				if shape != nil {
					refTopMidArcVectorShape[shape] = true
				}
			}
		}
	}

	for diagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if diagram.Rendered3DShape != nil {
			refRendered3DShape[diagram.Rendered3DShape] = true
		}
		if diagram.TorusStackShape != nil {
			refTorusStackShape[diagram.TorusStackShape] = true
		}
		if diagram.VerticalTorusStackShape != nil {
			refVerticalTorusStackShape[diagram.VerticalTorusStackShape] = true
		}
	}

	// Unstage unreferenced shapes
	for shape := range *GetGongstructInstancesSetFromPointerType[*AxesShape](stage) {
		if !refAxes[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*PlantCircumferenceShape](stage) {
		if !refPlantCirc[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*GridPathShape](stage) {
		if !refGridPath[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*CircleGridShape](stage) {
		if !refCircleGrid[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*ExplanationTextShape](stage) {
		if !refExplanation[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*GrowthVectorShape](stage) {
		if !refGrowthVector[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*InitialRhombusGridShape](stage) {
		if !refInitialGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}
	for grid := range *GetGongstructInstancesSetFromPointerType[*RotatedRhombusGridShape](stage) {
		if !refRotatedGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}
	for grid := range *GetGongstructInstancesSetFromPointerType[*GrowthCurveRhombusGridShape](stage) {
		if !refGrowthCurveGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}
	for grid := range *GetGongstructInstancesSetFromPointerType[*PerpendicularVectorGrid](stage) {
		if !refVectorGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*PerpendicularVectorGridHalfway](stage) {
		if !refPerpendicularVectorGridHalfway[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*BaseVectorShapeGrid](stage) {
		if !refBaseVectorShapeGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*ArcNormalVectorShapeGrid](stage) {
		if !refArcNormalVectorShapeGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*StartArcShapeGrid](stage) {
		if !refStartArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*TopStartArcShapeGrid](stage) {
		if !refTopStartArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*EndArcShapeGrid](stage) {
		if !refEndArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*TopEndArcShapeGrid](stage) {
		if !refTopEndArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for stack := range *GetGongstructInstancesSetFromPointerType[*StackOfRotatedGrowthCurve2D](stage) {
		if !refStackOfGrowthCurveV2[stack] {
			stack.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*RhombusShape](stage) {
		if !refRhombusShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*InitialRhombusShape](stage) {
		if !refInitialShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*RotatedRhombusShape](stage) {
		if !refRotatedShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*GrowthCurveRhombusShape](stage) {
		if !refGrowthCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for vec := range *GetGongstructInstancesSetFromPointerType[*PerpendicularVector](stage) {
		if !refPerpendicularVector[vec] {
			vec.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PerpendicularVectorHalfway](stage) {
		if !refPerpendicularVectorHalfway[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*BaseVectorShape](stage) {
		if !refBaseVectorShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ArcNormalVectorShape](stage) {
		if !refArcNormalVectorShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StartArcShape](stage) {
		if !refStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStartArcShape](stage) {
		if !refTopStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*EndArcShape](stage) {
		if !refEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*TopEndArcShape](stage) {
		if !refTopEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackRotatedGrowthCurve2DStartArcShape](stage) {
		if !refStackGrowthCurveStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*StackRotatedGrowthCurve2DEndArcShape](stage) {
		if !refStackGrowthCurveEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*Rendered3DShape](stage) {
		if !refRendered3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackOfRotatedGrowthCurve2D](stage) {
		if !refTopStackOfGrowthCurveV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackOfRotatedGrowthCurve2DStartArcShape](stage) {
		if !refTopStackGrowthCurveStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackOfRotatedGrowthCurve2DEndArcShape](stage) {
		if !refTopStackGrowthCurveEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedBottomTopStartArcShapeGrid](stage) {
		if !refShiftedBottomTopStartArcShapeGrid[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedBottomTopStartArcShape](stage) {
		if !refShiftedBottomTopStartArcShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*MidArcVectorShapeGrid](stage) {
		if !refMidArcVectorShapeGrid[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*MidArcVectorShape](stage) {
		if !refMidArcVectorShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopMidArcVectorShapeGrid](stage) {
		if !refTopMidArcVectorShapeGrid[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopMidArcVectorShape](stage) {
		if !refTopMidArcVectorShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftStackOfNormalVector](stage) {
		if !refShiftedLeftStackOfNormalVector[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftStackNormalVector](stage) {
		if !refShiftedLeftStackNormalVector[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftStackOfGrowthCurve](stage) {
		if !refShiftedLeftStackOfGrowthCurve[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftStackGrowthCurveStartArcShape](stage) {
		if !refShiftedLeftStackGrowthCurveStartArcShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftStackGrowthCurveEndArcShape](stage) {
		if !refShiftedLeftStackGrowthCurveEndArcShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for stack := range *GetGongstructInstancesSetFromPointerType[*StackOfGrowthCurve2D](stage) {
		if !refStackOfGrowthCurve2D[stack] {
			stack.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurve2DStartHalfwayArcShape](stage) {
		if !refStackGrowthCurve2DStartHalfwayArcShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurve2DEndHalfwayArcShape](stage) {
		if !refStackGrowthCurve2DEndHalfwayArcShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for stack := range *GetGongstructInstancesSetFromPointerType[*TopStackOfGrowthCurve2D](stage) {
		if !refTopStackOfGrowthCurve2D[stack] {
			stack.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackGrowthCurve2DStartHalfwayArcShape](stage) {
		if !refTopStackGrowthCurve2DStartHalfwayArcShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackGrowthCurve2DEndHalfwayArcShape](stage) {
		if !refTopStackGrowthCurve2DEndHalfwayArcShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackOfGrowthCurve2DRibbon](stage) {
		if !refStackOfGrowthCurve2DRibbon[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurve2DRibbonStartShape](stage) {
		if !refStackGrowthCurve2DRibbonStartShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurve2DRibbonEndShape](stage) {
		if !refStackGrowthCurve2DRibbonEndShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackOfRotatedGrowthCurve2DRibbon](stage) {
		if !refStackOfRotatedGrowthCurve2DRibbon[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackRotatedGrowthCurve2DRibbonStartShape](stage) {
		if !refStackRotatedGrowthCurve2DRibbonStartShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackRotatedGrowthCurve2DRibbonEndShape](stage) {
		if !refStackRotatedGrowthCurve2DRibbonEndShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*TorusStackShape](stage) {
		if !refTorusStackShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*VerticalTorusStackShape](stage) {
		if !refVerticalTorusStackShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	return needCommit
}
