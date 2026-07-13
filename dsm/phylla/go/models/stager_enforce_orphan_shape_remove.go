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
	refStartArcShapeV2Grid := make(map[*StartArcShapeGrid]bool)
	refTopStartArcShapeV2Grid := make(map[*TopStartArcShapeGrid]bool)
	refEndArcShapeV2Grid := make(map[*EndArcShapeGrid]bool)
	refTopEndArcShapeV2Grid := make(map[*TopEndArcShapeGrid]bool)
	refBottomStartArcShapeV2Grid := make(map[*BottomStartArcShapeGrid]bool)
	refBottomEndArcShapeV2Grid := make(map[*BottomEndArcShapeGrid]bool)
	refGrowthCurveBezierShapeGrid := make(map[*GrowthCurveBezierShapeGrid]bool)

	refGrowthCurveBezierShape := make(map[*GrowthCurveBezierShape]bool)
	refBaseVectorShape := make(map[*BaseVectorShape]bool)
	refStackOfGrowthCurveV2 := make(map[*StackOfGrowthCurve]bool)
	refStackGrowthCurveStartArcShapeV2 := make(map[*StackGrowthCurveStartArcShape]bool)
	refStackGrowthCurveEndArcShapeV2 := make(map[*StackGrowthCurveEndArcShape]bool)

	refTopStackOfGrowthCurveV2 := make(map[*TopStackOfGrowthCurve]bool)
	refTopStackGrowthCurveStartArcShapeV2 := make(map[*TopStackGrowthCurveStartArcShape]bool)
	refTopStackGrowthCurveEndArcShapeV2 := make(map[*TopStackGrowthCurveEndArcShape]bool)
	refBottomStackOfGrowthCurveV2 := make(map[*BottomStackOfGrowthCurve]bool)
	refBottomStackGrowthCurveStartArcShapeV2 := make(map[*BottomStackGrowthCurveStartArcShape]bool)
	refBottomStackGrowthCurveEndArcShapeV2 := make(map[*BottomStackGrowthCurveEndArcShape]bool)
	refRendered3DShape := make(map[*Rendered3DShape]bool)
	refArcNormalVectorShape := make(map[*ArcNormalVectorShape]bool)
	refStartArcShapeV2 := make(map[*StartArcShape]bool)
	refTopStartArcShapeV2 := make(map[*TopStartArcShape]bool)
	refEndArcShapeV2 := make(map[*EndArcShape]bool)
	refTopEndArcShapeV2 := make(map[*TopEndArcShape]bool)
	refBottomStartArcShapeV2 := make(map[*BottomStartArcShape]bool)
	refBottomEndArcShapeV2 := make(map[*BottomEndArcShape]bool)

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

		if plant.BottomStartArcShapeGrid != nil {
			refBottomStartArcShapeV2Grid[plant.BottomStartArcShapeGrid] = true
			for _, shape := range plant.BottomStartArcShapeGrid.BottomStartArcShapes {
				if shape != nil {
					refBottomStartArcShapeV2[shape] = true
				}
			}
		}

		if plant.BottomEndArcShapeGrid != nil {
			refBottomEndArcShapeV2Grid[plant.BottomEndArcShapeGrid] = true
			for _, shape := range plant.BottomEndArcShapeGrid.BottomEndArcShapes {
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
			refStackOfGrowthCurveV2[plant.StackOfGrowthCurve] = true
			for _, shape := range plant.StackOfGrowthCurve.StackGrowthCurveStartArcShapes {
				if shape != nil {
					refStackGrowthCurveStartArcShapeV2[shape] = true
				}
			}
			for _, shape := range plant.StackOfGrowthCurve.StackGrowthCurveEndArcShapes {
				if shape != nil {
					refStackGrowthCurveEndArcShapeV2[shape] = true
				}
			}
		}
		if plant.TopStackOfGrowthCurve != nil {
			refTopStackOfGrowthCurveV2[plant.TopStackOfGrowthCurve] = true
			for _, shape := range plant.TopStackOfGrowthCurve.TopStackGrowthCurveStartArcShapes {
				if shape != nil {
					refTopStackGrowthCurveStartArcShapeV2[shape] = true
				}
			}
			for _, shape := range plant.TopStackOfGrowthCurve.TopStackGrowthCurveEndArcShapes {
				if shape != nil {
					refTopStackGrowthCurveEndArcShapeV2[shape] = true
				}
			}
		}
		if plant.BottomStackOfGrowthCurve != nil {
			refBottomStackOfGrowthCurveV2[plant.BottomStackOfGrowthCurve] = true
			for _, shape := range plant.BottomStackOfGrowthCurve.BottomStackGrowthCurveStartArcShapes {
				if shape != nil {
					refBottomStackGrowthCurveStartArcShapeV2[shape] = true
				}
			}
			for _, shape := range plant.BottomStackOfGrowthCurve.BottomStackGrowthCurveEndArcShapes {
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

	for grid := range *GetGongstructInstancesSetFromPointerType[*BottomStartArcShapeGrid](stage) {
		if !refBottomStartArcShapeV2Grid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*BottomEndArcShapeGrid](stage) {
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

	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomStartArcShape](stage) {
		if !refBottomStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomEndArcShape](stage) {
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

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurveStartArcShape](stage) {
		if !refStackGrowthCurveStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*StackGrowthCurveEndArcShape](stage) {
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
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackOfGrowthCurve](stage) {
		if !refTopStackOfGrowthCurveV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackGrowthCurveStartArcShape](stage) {
		if !refTopStackGrowthCurveStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*TopStackGrowthCurveEndArcShape](stage) {
		if !refTopStackGrowthCurveEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomStackOfGrowthCurve](stage) {
		if !refBottomStackOfGrowthCurveV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomStackGrowthCurveStartArcShape](stage) {
		if !refBottomStackGrowthCurveStartArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*BottomStackGrowthCurveEndArcShape](stage) {
		if !refBottomStackGrowthCurveEndArcShapeV2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	return
}
