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

	refPartiallyGrowthCurve2DRibbon := make(map[*PartiallyGrowthCurve2DRibbon]bool)
	refPartiallyGrowthCurve2DRibbonStartShape := make(map[*PartiallyGrowthCurve2DRibbonStartShape]bool)
	refPartiallyGrowthCurve2DRibbonEndShape := make(map[*PartiallyGrowthCurve2DRibbonEndShape]bool)

	refShiftedLeftPartiallyGrowthCurve2DRibbon := make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbon]bool)
	refShiftedLeftPartiallyGrowthCurve2DRibbonStartShape := make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape]bool)
	refShiftedLeftPartiallyGrowthCurve2DRibbonEndShape := make(map[*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape]bool)

	refPartiallyGrowthCurve2DTrajectory := make(map[*PartiallyGrowthCurve2DTrajectory]bool)
	refPartiallyGrowthCurve2DTrajectoryShape := make(map[*PartiallyGrowthCurve2DTrajectoryShape]bool)

	refPartiallyGrowthCurve2DTrajectoryP1P2 := make(map[*PartiallyGrowthCurve2DTrajectoryP1P2]bool)
	refPartiallyGrowthCurve2DTrajectoryP1PointShape := make(map[*PartiallyGrowthCurve2DTrajectoryP1PointShape]bool)
	refPartiallyGrowthCurve2DTrajectoryP2PointShape := make(map[*PartiallyGrowthCurve2DTrajectoryP2PointShape]bool)
	refPartiallyGrowthCurve2DTrajectoryP1CurveShape := make(map[*PartiallyGrowthCurve2DTrajectoryP1CurveShape]bool)
	refPartiallyGrowthCurve2DTrajectoryP2CurveShape := make(map[*PartiallyGrowthCurve2DTrajectoryP2CurveShape]bool)
	refPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape := make(map[*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape]bool)
	refPxShape := make(map[*PxShape]bool)
	refChosenP1P2PairShape := make(map[*ChosenP1P2PairShape]bool)




	refGrowthCurve2DRibbon := make(map[*GrowthCurve2DRibbon]bool)
	refGrowthCurve2DRibbonStartShape := make(map[*GrowthCurve2DRibbonStartShape]bool)
	refGrowthCurve2DRibbonEndShape := make(map[*GrowthCurve2DRibbonEndShape]bool)

	refShiftedRightGrowthCurve2DRibbon := make(map[*ShiftedRightGrowthCurve2DRibbon]bool)
	refShiftedRightGrowthCurve2DRibbonStartShape := make(map[*ShiftedRightGrowthCurve2DRibbonStartShape]bool)
	refShiftedRightGrowthCurve2DRibbonEndShape := make(map[*ShiftedRightGrowthCurve2DRibbonEndShape]bool)

	refShiftedLeftGrowthCurve2DRibbon := make(map[*ShiftedLeftGrowthCurve2DRibbon]bool)
	refShiftedLeftGrowthCurve2DRibbonStartShape := make(map[*ShiftedLeftGrowthCurve2DRibbonStartShape]bool)
	refShiftedLeftGrowthCurve2DRibbonEndShape := make(map[*ShiftedLeftGrowthCurve2DRibbonEndShape]bool)



	refTorusStackShape := make(map[*TorusStackShape]bool)
	refVerticalTorusStackShape := make(map[*VerticalTorusStackShape]bool)
	refPartiallyRotatedTorusShape := make(map[*PartiallyRotatedTorusShape]bool)
	refStackOfPartiallyRotatedTorusShape := make(map[*StackOfPartiallyRotatedTorusShape]bool)

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

		if plant.PartiallyGrowthCurve2DRibbon != nil {
			refPartiallyGrowthCurve2DRibbon[plant.PartiallyGrowthCurve2DRibbon] = true
			for _, start := range plant.PartiallyGrowthCurve2DRibbon.PartiallyGrowthCurve2DRibbonStartShapes {
				refPartiallyGrowthCurve2DRibbonStartShape[start] = true
			}
			for _, end := range plant.PartiallyGrowthCurve2DRibbon.PartiallyGrowthCurve2DRibbonEndShapes {
				refPartiallyGrowthCurve2DRibbonEndShape[end] = true
			}
		}

		if plant.ShiftedLeftPartiallyGrowthCurve2DRibbon != nil {
			refShiftedLeftPartiallyGrowthCurve2DRibbon[plant.ShiftedLeftPartiallyGrowthCurve2DRibbon] = true
			for _, start := range plant.ShiftedLeftPartiallyGrowthCurve2DRibbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
				refShiftedLeftPartiallyGrowthCurve2DRibbonStartShape[start] = true
			}
			for _, end := range plant.ShiftedLeftPartiallyGrowthCurve2DRibbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
				refShiftedLeftPartiallyGrowthCurve2DRibbonEndShape[end] = true
			}
		}

		if plant.PartiallyGrowthCurve2DTrajectory != nil {
			refPartiallyGrowthCurve2DTrajectory[plant.PartiallyGrowthCurve2DTrajectory] = true
			for _, shape := range plant.PartiallyGrowthCurve2DTrajectory.PartiallyGrowthCurve2DTrajectoryShapes {
				refPartiallyGrowthCurve2DTrajectoryShape[shape] = true
			}
		}

		if plant.PartiallyGrowthCurve2DTrajectoryP1P2 != nil {
			refPartiallyGrowthCurve2DTrajectoryP1P2[plant.PartiallyGrowthCurve2DTrajectoryP1P2] = true
			for _, shape := range plant.PartiallyGrowthCurve2DTrajectoryP1P2.P1PointShapes {
				refPartiallyGrowthCurve2DTrajectoryP1PointShape[shape] = true
			}
			for _, shape := range plant.PartiallyGrowthCurve2DTrajectoryP1P2.P2PointShapes {
				refPartiallyGrowthCurve2DTrajectoryP2PointShape[shape] = true
			}
			for _, shape := range plant.PartiallyGrowthCurve2DTrajectoryP1P2.P1CurveShapes {
				refPartiallyGrowthCurve2DTrajectoryP1CurveShape[shape] = true
			}
			for _, shape := range plant.PartiallyGrowthCurve2DTrajectoryP1P2.P2CurveShapes {
				refPartiallyGrowthCurve2DTrajectoryP2CurveShape[shape] = true
			}
			for _, shape := range plant.PartiallyGrowthCurve2DTrajectoryP1P2.P1P2PairLineShapes {
				refPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape[shape] = true
			}
		}

		if plant.PxShape != nil {
			refPxShape[plant.PxShape] = true
		}

		if plant.ChosenP1P2PairShape != nil {
			refChosenP1P2PairShape[plant.ChosenP1P2PairShape] = true
		}



		if plant.GrowthCurve2DRibbon != nil {
			refGrowthCurve2DRibbon[plant.GrowthCurve2DRibbon] = true
			for _, start := range plant.GrowthCurve2DRibbon.GrowthCurve2DRibbonStartShapes {
				refGrowthCurve2DRibbonStartShape[start] = true
			}
			for _, end := range plant.GrowthCurve2DRibbon.GrowthCurve2DRibbonEndShapes {
				refGrowthCurve2DRibbonEndShape[end] = true
			}
		}

		if plant.ShiftedRightGrowthCurve2DRibbon != nil {
			refShiftedRightGrowthCurve2DRibbon[plant.ShiftedRightGrowthCurve2DRibbon] = true
			for _, start := range plant.ShiftedRightGrowthCurve2DRibbon.ShiftedRightGrowthCurve2DRibbonStartShapes {
				refShiftedRightGrowthCurve2DRibbonStartShape[start] = true
			}
			for _, end := range plant.ShiftedRightGrowthCurve2DRibbon.ShiftedRightGrowthCurve2DRibbonEndShapes {
				refShiftedRightGrowthCurve2DRibbonEndShape[end] = true
			}
		}

		if plant.ShiftedLeftGrowthCurve2DRibbon != nil {
			refShiftedLeftGrowthCurve2DRibbon[plant.ShiftedLeftGrowthCurve2DRibbon] = true
			for _, start := range plant.ShiftedLeftGrowthCurve2DRibbon.ShiftedLeftGrowthCurve2DRibbonStartShapes {
				refShiftedLeftGrowthCurve2DRibbonStartShape[start] = true
			}
			for _, end := range plant.ShiftedLeftGrowthCurve2DRibbon.ShiftedLeftGrowthCurve2DRibbonEndShapes {
				refShiftedLeftGrowthCurve2DRibbonEndShape[end] = true
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
		if diagram.PartiallyRotatedTorusShape != nil {
			refPartiallyRotatedTorusShape[diagram.PartiallyRotatedTorusShape] = true
		}
		if diagram.StackOfPartiallyRotatedTorusShape != nil {
			refStackOfPartiallyRotatedTorusShape[diagram.StackOfPartiallyRotatedTorusShape] = true
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

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DRibbon](stage) {
		if !refPartiallyGrowthCurve2DRibbon[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DRibbonStartShape](stage) {
		if !refPartiallyGrowthCurve2DRibbonStartShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DRibbonEndShape](stage) {
		if !refPartiallyGrowthCurve2DRibbonEndShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftPartiallyGrowthCurve2DRibbon](stage) {
		if !refShiftedLeftPartiallyGrowthCurve2DRibbon[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape](stage) {
		if !refShiftedLeftPartiallyGrowthCurve2DRibbonStartShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape](stage) {
		if !refShiftedLeftPartiallyGrowthCurve2DRibbonEndShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DTrajectory](stage) {
		if !refPartiallyGrowthCurve2DTrajectory[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DTrajectoryShape](stage) {
		if !refPartiallyGrowthCurve2DTrajectoryShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DTrajectoryP1P2](stage) {
		if !refPartiallyGrowthCurve2DTrajectoryP1P2[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DTrajectoryP1PointShape](stage) {
		if !refPartiallyGrowthCurve2DTrajectoryP1PointShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DTrajectoryP2PointShape](stage) {
		if !refPartiallyGrowthCurve2DTrajectoryP2PointShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DTrajectoryP1CurveShape](stage) {
		if !refPartiallyGrowthCurve2DTrajectoryP1CurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DTrajectoryP2CurveShape](stage) {
		if !refPartiallyGrowthCurve2DTrajectoryP2CurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape](stage) {
		if !refPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PxShape](stage) {
		if !refPxShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ChosenP1P2PairShape](stage) {
		if !refChosenP1P2PairShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}




	for shape := range *GetGongstructInstancesSetFromPointerType[*GrowthCurve2DRibbon](stage) {
		if !refGrowthCurve2DRibbon[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*GrowthCurve2DRibbonStartShape](stage) {
		if !refGrowthCurve2DRibbonStartShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*GrowthCurve2DRibbonEndShape](stage) {
		if !refGrowthCurve2DRibbonEndShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedRightGrowthCurve2DRibbon](stage) {
		if !refShiftedRightGrowthCurve2DRibbon[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedRightGrowthCurve2DRibbonStartShape](stage) {
		if !refShiftedRightGrowthCurve2DRibbonStartShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedRightGrowthCurve2DRibbonEndShape](stage) {
		if !refShiftedRightGrowthCurve2DRibbonEndShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftGrowthCurve2DRibbon](stage) {
		if !refShiftedLeftGrowthCurve2DRibbon[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftGrowthCurve2DRibbonStartShape](stage) {
		if !refShiftedLeftGrowthCurve2DRibbonStartShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ShiftedLeftGrowthCurve2DRibbonEndShape](stage) {
		if !refShiftedLeftGrowthCurve2DRibbonEndShape[shape] {
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

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyRotatedTorusShape](stage) {
		if !refPartiallyRotatedTorusShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StackOfPartiallyRotatedTorusShape](stage) {
		if !refStackOfPartiallyRotatedTorusShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	return needCommit
}
