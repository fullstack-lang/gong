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
	refTiledFloor3DShape := make(map[*TiledFloor3DShape]bool)
	refStemCylinder3DShape := make(map[*StemCylinder3DShape]bool)
	refParastichyNCurves3DShape := make(map[*ParastichyNCurves3DShape]bool)
	refParastichyMCurves3DShape := make(map[*ParastichyMCurves3DShape]bool)
	refCutLine3DShape := make(map[*CutLine3DShape]bool)
	refCircumference3DShape := make(map[*Circumference3DShape]bool)
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
	refKeyHoleShape := make(map[*KeyHoleShape]bool)

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
	refPointsAndLines3DShape := make(map[*PointsAndLines3DShape]bool)
	refKeyHole3DShape := make(map[*KeyHole3DShape]bool)
	refKey3DShape := make(map[*Key3DShape]bool)
	refVolumeKey3DShape := make(map[*VolumeKey3DShape]bool)
	refTorusEdge3DShape := make(map[*TorusEdge3DShape]bool)
	refSampledPoints3DShape := make(map[*SampledPoints3DShape]bool)
	refRotatedSampledPoints3DShape := make(map[*RotatedSampledPoints3DShape]bool)
	refEyeSampledPoints3DShape := make(map[*EyeSampledPoints3DShape]bool)
	refEyeCornersSampledPoints3DShape := make(map[*EyeCornersSampledPoints3DShape]bool)
	refOriginalPoints3DShape := make(map[*OriginalPoints3DShape]bool)
	refAngle0Shape := make(map[*Angle0Shape]bool)
	refEye3DShape := make(map[*Eye3DShape]bool)
	refEyeSeatBottomCurveShape := make(map[*EyeSeatBottomCurveShape]bool)
	refEyeStoolBottomCurveShape := make(map[*EyeStoolBottomCurveShape]bool)
	refSeat3DShape := make(map[*Seat3DShape]bool)
	refEyeVolume3DShape := make(map[*EyeVolume3DShape]bool)
	refSeatAndLegs3DShape := make(map[*SeatAndLegs3DShape]bool)
	refRotatedSeatAndLegs3DShape := make(map[*RotatedSeatAndLegs3DShape]bool)
	refClockTopCurveShape := make(map[*ClockTopCurveShape]bool)
	refSeatTopCurveShape := make(map[*SeatTopCurveShape]bool)
	refPartiallyRotatedSeatTopCurveShape := make(map[*PartiallyRotatedSeatTopCurveShape]bool)
	refSeatBottomCurveShape := make(map[*SeatBottomCurveShape]bool)
	refPartiallyRotatedSeatBottomCurveShape := make(map[*PartiallyRotatedSeatBottomCurveShape]bool)
	refTorus3DShape := make(map[*Torus3DShape]bool)

	// Collect referenced shapes from all plants
	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stage) {
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

		if plant.EndArcShapeGrid != nil {
			refEndArcShapeV2Grid[plant.EndArcShapeGrid] = true
			for _, shape := range plant.EndArcShapeGrid.EndArcShapes {
				if shape != nil {
					refEndArcShapeV2[shape] = true
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

		if vase := plant.VaseAbstract; vase != nil {
			if vase.PerpendicularVectorGridHalfway != nil {
				refPerpendicularVectorGridHalfway[vase.PerpendicularVectorGridHalfway] = true
				for _, vec := range vase.PerpendicularVectorGridHalfway.PerpendicularVectorHalfways {
					if vec != nil {
						refPerpendicularVectorHalfway[vec] = true
					}
				}
			}

			if vase.TopStartArcShapeGrid != nil {
				refTopStartArcShapeV2Grid[vase.TopStartArcShapeGrid] = true
				for _, shape := range vase.TopStartArcShapeGrid.TopStartArcShapes {
					if shape != nil {
						refTopStartArcShapeV2[shape] = true
					}
				}
			}

			if vase.TopEndArcShapeGrid != nil {
				refTopEndArcShapeV2Grid[vase.TopEndArcShapeGrid] = true
				for _, shape := range vase.TopEndArcShapeGrid.TopEndArcShapes {
					if shape != nil {
						refTopEndArcShapeV2[shape] = true
					}
				}
			}

			if vase.StackOfRotatedGrowthCurve2D != nil {
				refStackOfGrowthCurveV2[vase.StackOfRotatedGrowthCurve2D] = true
				for _, shape := range vase.StackOfRotatedGrowthCurve2D.StackRotatedGrowthCurve2DStartArcShapes {
					if shape != nil {
						refStackGrowthCurveStartArcShapeV2[shape] = true
					}
				}
				for _, shape := range vase.StackOfRotatedGrowthCurve2D.StackRotatedGrowthCurve2DEndArcShapes {
					if shape != nil {
						refStackGrowthCurveEndArcShapeV2[shape] = true
					}
				}
			}
			if vase.TopStackOfRotatedGrowthCurve2D != nil {
				refTopStackOfGrowthCurveV2[vase.TopStackOfRotatedGrowthCurve2D] = true
				for _, shape := range vase.TopStackOfRotatedGrowthCurve2D.TopStackOfRotatedGrowthCurve2DStartArcShapes {
					if shape != nil {
						refTopStackGrowthCurveStartArcShapeV2[shape] = true
					}
				}
				for _, shape := range vase.TopStackOfRotatedGrowthCurve2D.TopStackOfRotatedGrowthCurve2DEndArcShapes {
					if shape != nil {
						refTopStackGrowthCurveEndArcShapeV2[shape] = true
					}
				}
			}

			if vase.StackOfGrowthCurve2D != nil {
				refStackOfGrowthCurve2D[vase.StackOfGrowthCurve2D] = true
				for _, shape := range vase.StackOfGrowthCurve2D.StackGrowthCurve2DStartHalfwayArcShapes {
					if shape != nil {
						refStackGrowthCurve2DStartHalfwayArcShape[shape] = true
					}
				}
				for _, shape := range vase.StackOfGrowthCurve2D.StackGrowthCurve2DEndHalfwayArcShapes {
					if shape != nil {
						refStackGrowthCurve2DEndHalfwayArcShape[shape] = true
					}
				}
			}

			if vase.TopStackOfGrowthCurve2D != nil {
				refTopStackOfGrowthCurve2D[vase.TopStackOfGrowthCurve2D] = true
				for _, start := range vase.TopStackOfGrowthCurve2D.TopStackGrowthCurve2DStartHalfwayArcShapes {
					refTopStackGrowthCurve2DStartHalfwayArcShape[start] = true
				}
				for _, end := range vase.TopStackOfGrowthCurve2D.TopStackGrowthCurve2DEndHalfwayArcShapes {
					refTopStackGrowthCurve2DEndHalfwayArcShape[end] = true
				}
			}

			if vase.StackOfGrowthCurve2DRibbon != nil {
				refStackOfGrowthCurve2DRibbon[vase.StackOfGrowthCurve2DRibbon] = true
				for _, start := range vase.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonStartShapes {
					refStackGrowthCurve2DRibbonStartShape[start] = true
				}
				for _, end := range vase.StackOfGrowthCurve2DRibbon.StackGrowthCurve2DRibbonEndShapes {
					refStackGrowthCurve2DRibbonEndShape[end] = true
				}
			}

			if vase.StackOfRotatedGrowthCurve2DRibbon != nil {
				refStackOfRotatedGrowthCurve2DRibbon[vase.StackOfRotatedGrowthCurve2DRibbon] = true
				for _, start := range vase.StackOfRotatedGrowthCurve2DRibbon.StackRotatedGrowthCurve2DRibbonStartShapes {
					refStackRotatedGrowthCurve2DRibbonStartShape[start] = true
				}
				for _, end := range vase.StackOfRotatedGrowthCurve2DRibbon.StackRotatedGrowthCurve2DRibbonEndShapes {
					refStackRotatedGrowthCurve2DRibbonEndShape[end] = true
				}
			}

			if vase.PartiallyGrowthCurve2DRibbon != nil {
				refPartiallyGrowthCurve2DRibbon[vase.PartiallyGrowthCurve2DRibbon] = true
				for _, start := range vase.PartiallyGrowthCurve2DRibbon.PartiallyGrowthCurve2DRibbonStartShapes {
					refPartiallyGrowthCurve2DRibbonStartShape[start] = true
				}
				for _, end := range vase.PartiallyGrowthCurve2DRibbon.PartiallyGrowthCurve2DRibbonEndShapes {
					refPartiallyGrowthCurve2DRibbonEndShape[end] = true
				}
			}

			if vase.ShiftedLeftPartiallyGrowthCurve2DRibbon != nil {
				refShiftedLeftPartiallyGrowthCurve2DRibbon[vase.ShiftedLeftPartiallyGrowthCurve2DRibbon] = true
				for _, start := range vase.ShiftedLeftPartiallyGrowthCurve2DRibbon.ShiftedLeftPartiallyGrowthCurve2DRibbonStartShapes {
					refShiftedLeftPartiallyGrowthCurve2DRibbonStartShape[start] = true
				}
				for _, end := range vase.ShiftedLeftPartiallyGrowthCurve2DRibbon.ShiftedLeftPartiallyGrowthCurve2DRibbonEndShapes {
					refShiftedLeftPartiallyGrowthCurve2DRibbonEndShape[end] = true
				}
			}

			if vase.PartiallyGrowthCurve2DTrajectory != nil {
				refPartiallyGrowthCurve2DTrajectory[vase.PartiallyGrowthCurve2DTrajectory] = true
				for _, shape := range vase.PartiallyGrowthCurve2DTrajectory.PartiallyGrowthCurve2DTrajectoryShapes {
					refPartiallyGrowthCurve2DTrajectoryShape[shape] = true
				}
			}

			if vase.PartiallyGrowthCurve2DTrajectoryP1P2 != nil {
				refPartiallyGrowthCurve2DTrajectoryP1P2[vase.PartiallyGrowthCurve2DTrajectoryP1P2] = true
				for _, shape := range vase.PartiallyGrowthCurve2DTrajectoryP1P2.P1PointShapes {
					refPartiallyGrowthCurve2DTrajectoryP1PointShape[shape] = true
				}
				for _, shape := range vase.PartiallyGrowthCurve2DTrajectoryP1P2.P2PointShapes {
					refPartiallyGrowthCurve2DTrajectoryP2PointShape[shape] = true
				}
				for _, shape := range vase.PartiallyGrowthCurve2DTrajectoryP1P2.P1CurveShapes {
					refPartiallyGrowthCurve2DTrajectoryP1CurveShape[shape] = true
				}
				for _, shape := range vase.PartiallyGrowthCurve2DTrajectoryP1P2.P2CurveShapes {
					refPartiallyGrowthCurve2DTrajectoryP2CurveShape[shape] = true
				}
				for _, shape := range vase.PartiallyGrowthCurve2DTrajectoryP1P2.P1P2PairLineShapes {
					refPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape[shape] = true
				}
			}

			if vase.PxShape != nil {
				refPxShape[vase.PxShape] = true
			}

			if vase.ChosenP1P2PairShape != nil {
				refChosenP1P2PairShape[vase.ChosenP1P2PairShape] = true
			}

			if vase.KeyHoleShape != nil {
				refKeyHoleShape[vase.KeyHoleShape] = true
			}

			if vase.GrowthCurve2DRibbon != nil {
				refGrowthCurve2DRibbon[vase.GrowthCurve2DRibbon] = true
				for _, start := range vase.GrowthCurve2DRibbon.GrowthCurve2DRibbonStartShapes {
					refGrowthCurve2DRibbonStartShape[start] = true
				}
				for _, end := range vase.GrowthCurve2DRibbon.GrowthCurve2DRibbonEndShapes {
					refGrowthCurve2DRibbonEndShape[end] = true
				}
			}

			if vase.ShiftedRightGrowthCurve2DRibbon != nil {
				refShiftedRightGrowthCurve2DRibbon[vase.ShiftedRightGrowthCurve2DRibbon] = true
				for _, start := range vase.ShiftedRightGrowthCurve2DRibbon.ShiftedRightGrowthCurve2DRibbonStartShapes {
					refShiftedRightGrowthCurve2DRibbonStartShape[start] = true
				}
				for _, end := range vase.ShiftedRightGrowthCurve2DRibbon.ShiftedRightGrowthCurve2DRibbonEndShapes {
					refShiftedRightGrowthCurve2DRibbonEndShape[end] = true
				}
			}

			if vase.ShiftedLeftGrowthCurve2DRibbon != nil {
				refShiftedLeftGrowthCurve2DRibbon[vase.ShiftedLeftGrowthCurve2DRibbon] = true
				for _, start := range vase.ShiftedLeftGrowthCurve2DRibbon.ShiftedLeftGrowthCurve2DRibbonStartShapes {
					refShiftedLeftGrowthCurve2DRibbonStartShape[start] = true
				}
				for _, end := range vase.ShiftedLeftGrowthCurve2DRibbon.ShiftedLeftGrowthCurve2DRibbonEndShapes {
					refShiftedLeftGrowthCurve2DRibbonEndShape[end] = true
				}
			}

			if vase.ShiftedBottomTopStartArcShapeGrid != nil {
				refShiftedBottomTopStartArcShapeGrid[vase.ShiftedBottomTopStartArcShapeGrid] = true
				for _, shape := range vase.ShiftedBottomTopStartArcShapeGrid.ShiftedBottomTopStartArcShapes {
					if shape != nil {
						refShiftedBottomTopStartArcShape[shape] = true
					}
				}
			}
			if vase.TopMidArcVectorShapeGrid != nil {
				refTopMidArcVectorShapeGrid[vase.TopMidArcVectorShapeGrid] = true
				for _, shape := range vase.TopMidArcVectorShapeGrid.TopMidArcVectorShapes {
					if shape != nil {
						refTopMidArcVectorShape[shape] = true
					}
				}
			}
		}
	}

	for diagram := range *GetGongstructInstancesSetFromPointerType[*Vase3DDiagram](stage) {
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
		if diagram.PointsAndLines3DShape != nil {
			refPointsAndLines3DShape[diagram.PointsAndLines3DShape] = true
		}
		if diagram.SampledPoints3DShape != nil {
			refSampledPoints3DShape[diagram.SampledPoints3DShape] = true
		}
		if diagram.OriginalPoints3DShape != nil {
			refOriginalPoints3DShape[diagram.OriginalPoints3DShape] = true
		}
		if diagram.Angle0Shape != nil {
			refAngle0Shape[diagram.Angle0Shape] = true
		}
		if diagram.KeyHole3DShape != nil {
			refKeyHole3DShape[diagram.KeyHole3DShape] = true
		}
		if diagram.Key3DShape != nil {
			refKey3DShape[diagram.Key3DShape] = true
		}
		if diagram.VolumeKey3DShape != nil {
			refVolumeKey3DShape[diagram.VolumeKey3DShape] = true
		}
		if diagram.TorusEdge3DShape != nil {
			refTorusEdge3DShape[diagram.TorusEdge3DShape] = true
		}
		if diagram.TiledFloor3DShape != nil {
			refTiledFloor3DShape[diagram.TiledFloor3DShape] = true
		}
	}
	for diagram := range *GetGongstructInstancesSetFromPointerType[*Stool3DDiagram](stage) {
		if diagram.Rendered3DShape != nil {
			refRendered3DShape[diagram.Rendered3DShape] = true
		}
		if diagram.Torus3DShape != nil {
			refTorus3DShape[diagram.Torus3DShape] = true
		}
		if diagram.RotatedTorusShape != nil {
			refPartiallyRotatedTorusShape[diagram.RotatedTorusShape] = true
		}
		if diagram.SampledPoints3DShape != nil {
			refSampledPoints3DShape[diagram.SampledPoints3DShape] = true
		}
		if diagram.RotatedSampledPoints3DShape != nil {
			refRotatedSampledPoints3DShape[diagram.RotatedSampledPoints3DShape] = true
		}
		if diagram.EyeSampledPoints3DShape != nil {
			refEyeSampledPoints3DShape[diagram.EyeSampledPoints3DShape] = true
		}
		if diagram.EyeCornersSampledPoints3DShape != nil {
			refEyeCornersSampledPoints3DShape[diagram.EyeCornersSampledPoints3DShape] = true
		}
		if diagram.Eye3DShape != nil {
			refEye3DShape[diagram.Eye3DShape] = true
		}
		if diagram.SeatTopCurveShape != nil {
			refSeatTopCurveShape[diagram.SeatTopCurveShape] = true
		}
		if diagram.RotatedSeatTopCurveShape != nil {
			refPartiallyRotatedSeatTopCurveShape[diagram.RotatedSeatTopCurveShape] = true
		}
		if diagram.SeatBottomCurveShape != nil {
			refSeatBottomCurveShape[diagram.SeatBottomCurveShape] = true
		}
		if diagram.RotatedSeatBottomCurveShape != nil {
			refPartiallyRotatedSeatBottomCurveShape[diagram.RotatedSeatBottomCurveShape] = true
		}
		if diagram.EyeSeatBottomCurveShape != nil {
			refEyeSeatBottomCurveShape[diagram.EyeSeatBottomCurveShape] = true
		}
		if diagram.EyeStoolBottomCurveShape != nil {
			refEyeStoolBottomCurveShape[diagram.EyeStoolBottomCurveShape] = true
		}
		if diagram.Seat3DShape != nil {
			refSeat3DShape[diagram.Seat3DShape] = true
		}
		if diagram.EyeVolume3DShape != nil {
			refEyeVolume3DShape[diagram.EyeVolume3DShape] = true
		}
		if diagram.SeatAndLegs3DShape != nil {
			refSeatAndLegs3DShape[diagram.SeatAndLegs3DShape] = true
		}
		if diagram.RotatedSeatAndLegs3DShape != nil {
			refRotatedSeatAndLegs3DShape[diagram.RotatedSeatAndLegs3DShape] = true
		}
		if diagram.TiledFloor3DShape != nil {
			refTiledFloor3DShape[diagram.TiledFloor3DShape] = true
		}
	}
	for diagram := range *GetGongstructInstancesSetFromPointerType[*Clock3DDiagram](stage) {
		if diagram.Rendered3DShape != nil {
			refRendered3DShape[diagram.Rendered3DShape] = true
		}
		if diagram.Torus3DShape != nil {
			refTorus3DShape[diagram.Torus3DShape] = true
		}
		if diagram.SampledPoints3DShape != nil {
			refSampledPoints3DShape[diagram.SampledPoints3DShape] = true
		}
		if diagram.ClockTopCurveShape != nil {
			refClockTopCurveShape[diagram.ClockTopCurveShape] = true
		}
		if diagram.TiledFloor3DShape != nil {
			refTiledFloor3DShape[diagram.TiledFloor3DShape] = true
		}
	}
	for diagram := range *GetGongstructInstancesSetFromPointerType[*Plant3DDiagram](stage) {
		if diagram.Rendered3DShape != nil {
			refRendered3DShape[diagram.Rendered3DShape] = true
		}
		if diagram.StemCylinder3DShape != nil {
			refStemCylinder3DShape[diagram.StemCylinder3DShape] = true
		}
		if diagram.ParastichyNCurves3DShape != nil {
			refParastichyNCurves3DShape[diagram.ParastichyNCurves3DShape] = true
		}
		if diagram.ParastichyMCurves3DShape != nil {
			refParastichyMCurves3DShape[diagram.ParastichyMCurves3DShape] = true
		}
		if diagram.CutLine3DShape != nil {
			refCutLine3DShape[diagram.CutLine3DShape] = true
		}
		if diagram.Circumference3DShape != nil {
			refCircumference3DShape[diagram.Circumference3DShape] = true
		}
		if diagram.TiledFloor3DShape != nil {
			refTiledFloor3DShape[diagram.TiledFloor3DShape] = true
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
	for shape := range *GetGongstructInstancesSetFromPointerType[*TiledFloor3DShape](stage) {
		if !refTiledFloor3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*OriginalPoints3DShape](stage) {
		if !refOriginalPoints3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*Angle0Shape](stage) {
		if !refAngle0Shape[shape] {
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

	for shape := range *GetGongstructInstancesSetFromPointerType[*KeyHoleShape](stage) {
		if !refKeyHoleShape[shape] {
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

	for shape := range *GetGongstructInstancesSetFromPointerType[*KeyHole3DShape](stage) {
		if !refKeyHole3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*Key3DShape](stage) {
		if !refKey3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*VolumeKey3DShape](stage) {
		if !refVolumeKey3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*TorusEdge3DShape](stage) {
		if !refTorusEdge3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*SampledPoints3DShape](stage) {
		if !refSampledPoints3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ClockTopCurveShape](stage) {
		if !refClockTopCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*SeatTopCurveShape](stage) {
		if !refSeatTopCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyRotatedSeatTopCurveShape](stage) {
		if !refPartiallyRotatedSeatTopCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*SeatBottomCurveShape](stage) {
		if !refSeatBottomCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*PartiallyRotatedSeatBottomCurveShape](stage) {
		if !refPartiallyRotatedSeatBottomCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*Torus3DShape](stage) {
		if !refTorus3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*RotatedSampledPoints3DShape](stage) {
		if !refRotatedSampledPoints3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*EyeSampledPoints3DShape](stage) {
		if !refEyeSampledPoints3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*EyeCornersSampledPoints3DShape](stage) {
		if !refEyeCornersSampledPoints3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*Eye3DShape](stage) {
		if !refEye3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*EyeSeatBottomCurveShape](stage) {
		if !refEyeSeatBottomCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*EyeStoolBottomCurveShape](stage) {
		if !refEyeStoolBottomCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*Seat3DShape](stage) {
		if !refSeat3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*EyeVolume3DShape](stage) {
		if !refEyeVolume3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*SeatAndLegs3DShape](stage) {
		if !refSeatAndLegs3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*RotatedSeatAndLegs3DShape](stage) {
		if !refRotatedSeatAndLegs3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*StemCylinder3DShape](stage) {
		if !refStemCylinder3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ParastichyNCurves3DShape](stage) {
		if !refParastichyNCurves3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*ParastichyMCurves3DShape](stage) {
		if !refParastichyMCurves3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*CutLine3DShape](stage) {
		if !refCutLine3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*Circumference3DShape](stage) {
		if !refCircumference3DShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	return needCommit
}
