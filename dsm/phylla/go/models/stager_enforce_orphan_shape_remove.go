package models

func (stager *Stager) enforceOrphanShapeRemove() (needCommit bool) {
	stage := stager.stage

	// Sets of referenced shapes
	refAxes := make(map[*AxesShape]bool)
	refPlantCirc := make(map[*PlantCircumferenceShape]bool)
	refGridPath := make(map[*GridPathShape]bool)
	refCircleGrid := make(map[*CircleGridShape]bool)
	refNextCircle := make(map[*NextCircleShape]bool)
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
	refStartArcShapeV2Grid := make(map[*StartArcShapeV2Grid]bool)
	refTopStartArcShapeV2Grid := make(map[*TopStartArcShapeV2Grid]bool)
	refEndArcShapeV2Grid := make(map[*EndArcShapeV2Grid]bool)
	refTopEndArcShapeV2Grid := make(map[*TopEndArcShapeV2Grid]bool)
	refBottomStartArcShapeV2Grid := make(map[*BottomStartArcShapeV2Grid]bool)
	refBottomEndArcShapeV2Grid := make(map[*BottomEndArcShapeV2Grid]bool)
	refGrowthCurveBezierShapeGrid := make(map[*GrowthCurveBezierShapeGrid]bool)
	refStackOfGrowthCurve := make(map[*StackOfGrowthCurve]bool)
	refGrowthCurveBezierShape := make(map[*GrowthCurveBezierShape]bool)
	refBaseVectorShape := make(map[*BaseVectorShape]bool)
	refStackOfGrowthCurveV2 := make(map[*StackOfGrowthCurveV2]bool)
	refStackGrowthCurveStartArcShapeV2 := make(map[*StackGrowthCurveStartArcShapeV2]bool)
	refStackGrowthCurveEndArcShapeV2 := make(map[*StackGrowthCurveEndArcShapeV2]bool)
	refStackGrowthCurveBezierShape := make(map[*StackGrowthCurveBezierShape]bool)
	refTopStackOfGrowthCurveV2 := make(map[*TopStackOfGrowthCurveV2]bool)
	refTopStackGrowthCurveStartArcShapeV2 := make(map[*TopStackGrowthCurveStartArcShapeV2]bool)
	refTopStackGrowthCurveEndArcShapeV2 := make(map[*TopStackGrowthCurveEndArcShapeV2]bool)
	refBottomStackOfGrowthCurveV2 := make(map[*BottomStackOfGrowthCurveV2]bool)
	refBottomStackGrowthCurveStartArcShapeV2 := make(map[*BottomStackGrowthCurveStartArcShapeV2]bool)
	refBottomStackGrowthCurveEndArcShapeV2 := make(map[*BottomStackGrowthCurveEndArcShapeV2]bool)
	refRendered3DShape := make(map[*Rendered3DShape]bool)
	refArcNormalVectorShape := make(map[*ArcNormalVectorShape]bool)
	refStartArcShapeV2 := make(map[*StartArcShapeV2]bool)
	refTopStartArcShapeV2 := make(map[*TopStartArcShapeV2]bool)
	refEndArcShapeV2 := make(map[*EndArcShapeV2]bool)
	refTopEndArcShapeV2 := make(map[*TopEndArcShapeV2]bool)
	refBottomStartArcShapeV2 := make(map[*BottomStartArcShapeV2]bool)
	refBottomEndArcShapeV2 := make(map[*BottomEndArcShapeV2]bool)

	// Collect referenced shapes from all plants
	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		if plant.AxesShape != nil {
			refAxes[plant.AxesShape] = true
		}
		if plant.ReferenceRhombus != nil {
			refRhombusShape[plant.ReferenceRhombus] = true
		}
		if plant.PlantCircumferenceShape != nil {
			refPlantCirc[plant.PlantCircumferenceShape] = true
		}
		if plant.GridPathShape != nil {
			refGridPath[plant.GridPathShape] = true
		}
		if plant.ExplanationTextShape != nil {
			refExplanation[plant.ExplanationTextShape] = true
		}
		if plant.RotatedReferenceRhombus != nil {
			refRhombusShape[plant.RotatedReferenceRhombus] = true
		}
		if plant.RotatedPlantCircumferenceShape != nil {
			refPlantCirc[plant.RotatedPlantCircumferenceShape] = true
		}
		if plant.RotatedGridPathShape != nil {
			refGridPath[plant.RotatedGridPathShape] = true
		}
		if plant.GrowthVectorShape != nil {
			refGrowthVector[plant.GrowthVectorShape] = true
		}

		if plant.InitialRhombusGridShape != nil {
			refInitialGrid[plant.InitialRhombusGridShape] = true
			for _, shape := range plant.InitialRhombusGridShape.InitialRhombusShapes {
				if shape != nil {
					refInitialShape[shape] = true
				}
			}
		}

		if plant.RotatedRhombusGridShape2 != nil {
			refRotatedGrid[plant.RotatedRhombusGridShape2] = true
			for _, shape := range plant.RotatedRhombusGridShape2.RotatedRhombusShapes {
				if shape != nil {
					refRotatedShape[shape] = true
				}
			}
		}

		if plant.GrowthCurveRhombusGridShape != nil {
			refGrowthCurveGrid[plant.GrowthCurveRhombusGridShape] = true
			for _, shape := range plant.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes {
				if shape != nil {
					refGrowthCurveShape[shape] = true
				}
			}
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

		if plant.StartArcShapeV2Grid != nil {
			refStartArcShapeV2Grid[plant.StartArcShapeV2Grid] = true
			for _, shape := range plant.StartArcShapeV2Grid.StartArcShapesV2 {
				if shape != nil {
					refStartArcShapeV2[shape] = true
				}
			}
		}

		if plant.TopStartArcShapeV2Grid != nil {
			refTopStartArcShapeV2Grid[plant.TopStartArcShapeV2Grid] = true
			for _, shape := range plant.TopStartArcShapeV2Grid.TopStartArcShapesV2 {
				if shape != nil {
					refTopStartArcShapeV2[shape] = true
				}
			}
		}

		if plant.EndArcShapeV2Grid != nil {
			refEndArcShapeV2Grid[plant.EndArcShapeV2Grid] = true
			for _, shape := range plant.EndArcShapeV2Grid.EndArcShapesV2 {
				if shape != nil {
					refEndArcShapeV2[shape] = true
				}
			}
		}

		if plant.TopEndArcShapeV2Grid != nil {
			refTopEndArcShapeV2Grid[plant.TopEndArcShapeV2Grid] = true
			for _, shape := range plant.TopEndArcShapeV2Grid.TopEndArcShapesV2 {
				if shape != nil {
					refTopEndArcShapeV2[shape] = true
				}
			}
		}

		if plant.BottomStartArcShapeV2Grid != nil {
			refBottomStartArcShapeV2Grid[plant.BottomStartArcShapeV2Grid] = true
			for _, shape := range plant.BottomStartArcShapeV2Grid.BottomStartArcShapesV2 {
				if shape != nil {
					refBottomStartArcShapeV2[shape] = true
				}
			}
		}

		if plant.BottomEndArcShapeV2Grid != nil {
			refBottomEndArcShapeV2Grid[plant.BottomEndArcShapeV2Grid] = true
			for _, shape := range plant.BottomEndArcShapeV2Grid.BottomEndArcShapesV2 {
				if shape != nil {
					refBottomEndArcShapeV2[shape] = true
				}
			}
		}

		if plant.GrowthCurveBezierShapeGrid != nil {
			refGrowthCurveBezierShapeGrid[plant.GrowthCurveBezierShapeGrid] = true
			for _, shape := range plant.GrowthCurveBezierShapeGrid.GrowthCurveBezierShapes {
				if shape != nil {
					refGrowthCurveBezierShape[shape] = true
				}
			}
		}

		if plant.StackOfGrowthCurve != nil {
			refStackOfGrowthCurve[plant.StackOfGrowthCurve] = true
			for _, shape := range plant.StackOfGrowthCurve.StackGrowthCurveBezierShapes {
				if shape != nil {
					refStackGrowthCurveBezierShape[shape] = true
				}
			}
		}
		if plant.StackOfGrowthCurveV2 != nil {
			refStackOfGrowthCurveV2[plant.StackOfGrowthCurveV2] = true
			for _, shape := range plant.StackOfGrowthCurveV2.StackGrowthCurveStartArcShapeV2s {
				if shape != nil {
					refStackGrowthCurveStartArcShapeV2[shape] = true
				}
			}
			for _, shape := range plant.StackOfGrowthCurveV2.StackGrowthCurveEndArcShapeV2s {
				if shape != nil {
					refStackGrowthCurveEndArcShapeV2[shape] = true
				}
			}
		}
		if plant.TopStackOfGrowthCurveV2 != nil {
			refTopStackOfGrowthCurveV2[plant.TopStackOfGrowthCurveV2] = true
			for _, shape := range plant.TopStackOfGrowthCurveV2.TopStackGrowthCurveStartArcShapeV2s {
				if shape != nil {
					refTopStackGrowthCurveStartArcShapeV2[shape] = true
				}
			}
			for _, shape := range plant.TopStackOfGrowthCurveV2.TopStackGrowthCurveEndArcShapeV2s {
				if shape != nil {
					refTopStackGrowthCurveEndArcShapeV2[shape] = true
				}
			}
		}
		if plant.BottomStackOfGrowthCurveV2 != nil {
			refBottomStackOfGrowthCurveV2[plant.BottomStackOfGrowthCurveV2] = true
			for _, shape := range plant.BottomStackOfGrowthCurveV2.BottomStackGrowthCurveStartArcShapeV2s {
				if shape != nil {
					refBottomStackGrowthCurveStartArcShapeV2[shape] = true
				}
			}
			for _, shape := range plant.BottomStackOfGrowthCurveV2.BottomStackGrowthCurveEndArcShapeV2s {
				if shape != nil {
					refBottomStackGrowthCurveEndArcShapeV2[shape] = true
				}
			}
		}
	}

	for diagram := range *GetGongstructInstancesSetFromPointerType[*PlantDiagram](stage) {
		if diagram.Rendered3DShape != nil {
			refRendered3DShape[diagram.Rendered3DShape] = true
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
	for shape := range *GetGongstructInstancesSetFromPointerType[*NextCircleShape](stage) {
		if !refNextCircle[shape] {
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

	for grid := range *GetGongstructInstancesSetFromPointerType[*StartArcShapeV2Grid](stage) {
		if !refStartArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*TopStartArcShapeV2Grid](stage) {
		if !refTopStartArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*EndArcShapeV2Grid](stage) {
		if !refEndArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*TopEndArcShapeV2Grid](stage) {
		if !refTopEndArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*BottomStartArcShapeV2Grid](stage) {
		if !refBottomStartArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*BottomEndArcShapeV2Grid](stage) {
		if !refBottomEndArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*GrowthCurveBezierShapeGrid](stage) {
		if !refGrowthCurveBezierShapeGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}
	for stack := range *GetGongstructInstancesSetFromPointerType[*StackOfGrowthCurve](stage) {
		if !refStackOfGrowthCurve[stack] {
			stack.Unstage(stage)
			needCommit = true
		}
	}
	for stack := range *GetGongstructInstancesSetFromPointerType[*StackOfGrowthCurveV2](stage) {
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

	for shape := range *GetGongstructInstancesSetFromPointerType[*StartArcShapeV2](stage) {
		if !refStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStartArcShapeV2](stage) {
		if !refTopStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*EndArcShapeV2](stage) {
		if !refEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*TopEndArcShapeV2](stage) {
		if !refTopEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomStartArcShapeV2](stage) {
		if !refBottomStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomEndArcShapeV2](stage) {
		if !refBottomEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*GrowthCurveBezierShape](stage) {
		if !refGrowthCurveBezierShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurveBezierShape](stage) {
		if !refStackGrowthCurveBezierShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurveStartArcShapeV2](stage) {
		if !refStackGrowthCurveStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurveEndArcShapeV2](stage) {
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
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackOfGrowthCurveV2](stage) {
		if !refTopStackOfGrowthCurveV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackGrowthCurveStartArcShapeV2](stage) {
		if !refTopStackGrowthCurveStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackGrowthCurveEndArcShapeV2](stage) {
		if !refTopStackGrowthCurveEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomStackOfGrowthCurveV2](stage) {
		if !refBottomStackOfGrowthCurveV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomStackGrowthCurveStartArcShapeV2](stage) {
		if !refBottomStackGrowthCurveStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomStackGrowthCurveEndArcShapeV2](stage) {
		if !refBottomStackGrowthCurveEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	return
}
