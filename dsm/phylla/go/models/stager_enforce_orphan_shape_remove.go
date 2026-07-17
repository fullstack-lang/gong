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

	refBaseVectorShape := make(map[*BaseVectorShape]bool)
	refStackOfGrowthCurveV2 := make(map[*StackOfGrowthCurve]bool)
	refStackGrowthCurveStartArcShapeV2 := make(map[*StackGrowthCurveStartArcShape]bool)
	refStackGrowthCurveEndArcShapeV2 := make(map[*StackGrowthCurveEndArcShape]bool)

	refTopStackOfGrowthCurveV2 := make(map[*TopStackOfGrowthCurve]bool)
	refTopStackGrowthCurveStartArcShapeV2 := make(map[*TopStackGrowthCurveStartArcShape]bool)
	refTopStackGrowthCurveEndArcShapeV2 := make(map[*TopStackGrowthCurveEndArcShape]bool)
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
		if plant.ShiftedLeftStackOfNormalVector != nil {
			refShiftedLeftStackOfNormalVector[plant.ShiftedLeftStackOfNormalVector] = true
			for _, shape := range plant.ShiftedLeftStackOfNormalVector.ShiftedLeftStackNormalVectors {
				if shape != nil {
					refShiftedLeftStackNormalVector[shape] = true
				}
			}
		}
		if plant.ShiftedLeftStackOfGrowthCurve != nil {
			refShiftedLeftStackOfGrowthCurve[plant.ShiftedLeftStackOfGrowthCurve] = true
			for _, shape := range plant.ShiftedLeftStackOfGrowthCurve.ShiftedLeftStackGrowthCurveStartArcShapes {
				if shape != nil {
					refShiftedLeftStackGrowthCurveStartArcShape[shape] = true
				}
			}
			for _, shape := range plant.ShiftedLeftStackOfGrowthCurve.ShiftedLeftStackGrowthCurveEndArcShapes {
				if shape != nil {
					refShiftedLeftStackGrowthCurveEndArcShape[shape] = true
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
	return
}
