// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Angle0Shape:
		if stage.OnAfterAngle0ShapeCreateCallback != nil {
			stage.OnAfterAngle0ShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ArcNormalVectorShape:
		if stage.OnAfterArcNormalVectorShapeCreateCallback != nil {
			stage.OnAfterArcNormalVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ArcNormalVectorShapeGrid:
		if stage.OnAfterArcNormalVectorShapeGridCreateCallback != nil {
			stage.OnAfterArcNormalVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *AxesShape:
		if stage.OnAfterAxesShapeCreateCallback != nil {
			stage.OnAfterAxesShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *BaseVectorShape:
		if stage.OnAfterBaseVectorShapeCreateCallback != nil {
			stage.OnAfterBaseVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *BaseVectorShapeGrid:
		if stage.OnAfterBaseVectorShapeGridCreateCallback != nil {
			stage.OnAfterBaseVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *ChosenP1P2PairShape:
		if stage.OnAfterChosenP1P2PairShapeCreateCallback != nil {
			stage.OnAfterChosenP1P2PairShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeCreateCallback != nil {
			stage.OnAfterCircleGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Circumference3DShape:
		if stage.OnAfterCircumference3DShapeCreateCallback != nil {
			stage.OnAfterCircumference3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Clock2DDiagram:
		if stage.OnAfterClock2DDiagramCreateCallback != nil {
			stage.OnAfterClock2DDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Clock3DDiagram:
		if stage.OnAfterClock3DDiagramCreateCallback != nil {
			stage.OnAfterClock3DDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *ClockAbstract:
		if stage.OnAfterClockAbstractCreateCallback != nil {
			stage.OnAfterClockAbstractCreateCallback.OnAfterCreate(stage, target)
		}
	case *ClockTopCurveShape:
		if stage.OnAfterClockTopCurveShapeCreateCallback != nil {
			stage.OnAfterClockTopCurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *CutLine3DShape:
		if stage.OnAfterCutLine3DShapeCreateCallback != nil {
			stage.OnAfterCutLine3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EndArcShape:
		if stage.OnAfterEndArcShapeCreateCallback != nil {
			stage.OnAfterEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EndArcShapeGrid:
		if stage.OnAfterEndArcShapeGridCreateCallback != nil {
			stage.OnAfterEndArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *EndHalfwayArcShape:
		if stage.OnAfterEndHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EndHalfwayArcShapeGrid:
		if stage.OnAfterEndHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterEndHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExplanationTextShape:
		if stage.OnAfterExplanationTextShapeCreateCallback != nil {
			stage.OnAfterExplanationTextShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Eye3DShape:
		if stage.OnAfterEye3DShapeCreateCallback != nil {
			stage.OnAfterEye3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EyeCornersSampledPoints3DShape:
		if stage.OnAfterEyeCornersSampledPoints3DShapeCreateCallback != nil {
			stage.OnAfterEyeCornersSampledPoints3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EyeSampledPoints3DShape:
		if stage.OnAfterEyeSampledPoints3DShapeCreateCallback != nil {
			stage.OnAfterEyeSampledPoints3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EyeSeatBottomCurveShape:
		if stage.OnAfterEyeSeatBottomCurveShapeCreateCallback != nil {
			stage.OnAfterEyeSeatBottomCurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EyeStoolBottomCurveShape:
		if stage.OnAfterEyeStoolBottomCurveShapeCreateCallback != nil {
			stage.OnAfterEyeStoolBottomCurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EyeVolume3DShape:
		if stage.OnAfterEyeVolume3DShapeCreateCallback != nil {
			stage.OnAfterEyeVolume3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeCreateCallback != nil {
			stage.OnAfterGridPathShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurve2D:
		if stage.OnAfterGrowthCurve2DCreateCallback != nil {
			stage.OnAfterGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurve2DRibbon:
		if stage.OnAfterGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurve2DRibbonEndShape:
		if stage.OnAfterGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurve2DRibbonStartShape:
		if stage.OnAfterGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurveRhombusGridShape:
		if stage.OnAfterGrowthCurveRhombusGridShapeCreateCallback != nil {
			stage.OnAfterGrowthCurveRhombusGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurveRhombusShape:
		if stage.OnAfterGrowthCurveRhombusShapeCreateCallback != nil {
			stage.OnAfterGrowthCurveRhombusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthVectorShape:
		if stage.OnAfterGrowthVectorShapeCreateCallback != nil {
			stage.OnAfterGrowthVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *InitialRhombusGridShape:
		if stage.OnAfterInitialRhombusGridShapeCreateCallback != nil {
			stage.OnAfterInitialRhombusGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *InitialRhombusShape:
		if stage.OnAfterInitialRhombusShapeCreateCallback != nil {
			stage.OnAfterInitialRhombusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Key3DShape:
		if stage.OnAfterKey3DShapeCreateCallback != nil {
			stage.OnAfterKey3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *KeyHole3DShape:
		if stage.OnAfterKeyHole3DShapeCreateCallback != nil {
			stage.OnAfterKeyHole3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *KeyHoleShape:
		if stage.OnAfterKeyHoleShapeCreateCallback != nil {
			stage.OnAfterKeyHoleShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Leaves3DShape:
		if stage.OnAfterLeaves3DShapeCreateCallback != nil {
			stage.OnAfterLeaves3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *MidArcVectorShape:
		if stage.OnAfterMidArcVectorShapeCreateCallback != nil {
			stage.OnAfterMidArcVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *MidArcVectorShapeGrid:
		if stage.OnAfterMidArcVectorShapeGridCreateCallback != nil {
			stage.OnAfterMidArcVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *MusicAbstract:
		if stage.OnAfterMusicAbstractCreateCallback != nil {
			stage.OnAfterMusicAbstractCreateCallback.OnAfterCreate(stage, target)
		}
	case *OriginalPoints3DShape:
		if stage.OnAfterOriginalPoints3DShapeCreateCallback != nil {
			stage.OnAfterOriginalPoints3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParastichyMCurves3DShape:
		if stage.OnAfterParastichyMCurves3DShapeCreateCallback != nil {
			stage.OnAfterParastichyMCurves3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ParastichyNCurves3DShape:
		if stage.OnAfterParastichyNCurves3DShapeCreateCallback != nil {
			stage.OnAfterParastichyNCurves3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbon:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbonEndShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbonStartShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectory:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2CreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2CreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyRotatedSeatBottomCurveShape:
		if stage.OnAfterPartiallyRotatedSeatBottomCurveShapeCreateCallback != nil {
			stage.OnAfterPartiallyRotatedSeatBottomCurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyRotatedSeatTopCurveShape:
		if stage.OnAfterPartiallyRotatedSeatTopCurveShapeCreateCallback != nil {
			stage.OnAfterPartiallyRotatedSeatTopCurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyRotatedTorusShape:
		if stage.OnAfterPartiallyRotatedTorusShapeCreateCallback != nil {
			stage.OnAfterPartiallyRotatedTorusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PerpendicularVector:
		if stage.OnAfterPerpendicularVectorCreateCallback != nil {
			stage.OnAfterPerpendicularVectorCreateCallback.OnAfterCreate(stage, target)
		}
	case *PerpendicularVectorGrid:
		if stage.OnAfterPerpendicularVectorGridCreateCallback != nil {
			stage.OnAfterPerpendicularVectorGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *PerpendicularVectorGridHalfway:
		if stage.OnAfterPerpendicularVectorGridHalfwayCreateCallback != nil {
			stage.OnAfterPerpendicularVectorGridHalfwayCreateCallback.OnAfterCreate(stage, target)
		}
	case *PerpendicularVectorHalfway:
		if stage.OnAfterPerpendicularVectorHalfwayCreateCallback != nil {
			stage.OnAfterPerpendicularVectorHalfwayCreateCallback.OnAfterCreate(stage, target)
		}
	case *Plant2DDiagram:
		if stage.OnAfterPlant2DDiagramCreateCallback != nil {
			stage.OnAfterPlant2DDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Plant3DDiagram:
		if stage.OnAfterPlant3DDiagramCreateCallback != nil {
			stage.OnAfterPlant3DDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *PlantAbstract:
		if stage.OnAfterPlantAbstractCreateCallback != nil {
			stage.OnAfterPlantAbstractCreateCallback.OnAfterCreate(stage, target)
		}
	case *PlantCircumferenceShape:
		if stage.OnAfterPlantCircumferenceShapeCreateCallback != nil {
			stage.OnAfterPlantCircumferenceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PointsAndLines3DShape:
		if stage.OnAfterPointsAndLines3DShapeCreateCallback != nil {
			stage.OnAfterPointsAndLines3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PxShape:
		if stage.OnAfterPxShapeCreateCallback != nil {
			stage.OnAfterPxShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Rendered3DShape:
		if stage.OnAfterRendered3DShapeCreateCallback != nil {
			stage.OnAfterRendered3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *RhombusShape:
		if stage.OnAfterRhombusShapeCreateCallback != nil {
			stage.OnAfterRhombusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *RhombusStuff:
		if stage.OnAfterRhombusStuffCreateCallback != nil {
			stage.OnAfterRhombusStuffCreateCallback.OnAfterCreate(stage, target)
		}
	case *RotatedRhombusGridShape:
		if stage.OnAfterRotatedRhombusGridShapeCreateCallback != nil {
			stage.OnAfterRotatedRhombusGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *RotatedRhombusShape:
		if stage.OnAfterRotatedRhombusShapeCreateCallback != nil {
			stage.OnAfterRotatedRhombusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *RotatedSampledPoints3DShape:
		if stage.OnAfterRotatedSampledPoints3DShapeCreateCallback != nil {
			stage.OnAfterRotatedSampledPoints3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *RotatedSeatAndLegs3DShape:
		if stage.OnAfterRotatedSeatAndLegs3DShapeCreateCallback != nil {
			stage.OnAfterRotatedSeatAndLegs3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *SampledPoints3DShape:
		if stage.OnAfterSampledPoints3DShapeCreateCallback != nil {
			stage.OnAfterSampledPoints3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Seat3DShape:
		if stage.OnAfterSeat3DShapeCreateCallback != nil {
			stage.OnAfterSeat3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *SeatAndLegs3DShape:
		if stage.OnAfterSeatAndLegs3DShapeCreateCallback != nil {
			stage.OnAfterSeatAndLegs3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *SeatBottomCurveShape:
		if stage.OnAfterSeatBottomCurveShapeCreateCallback != nil {
			stage.OnAfterSeatBottomCurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *SeatTopCurveShape:
		if stage.OnAfterSeatTopCurveShapeCreateCallback != nil {
			stage.OnAfterSeatTopCurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedBottomTopStartArcShape:
		if stage.OnAfterShiftedBottomTopStartArcShapeCreateCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedBottomTopStartArcShapeGrid:
		if stage.OnAfterShiftedBottomTopStartArcShapeGridCreateCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftGrowthCurve2DRibbon:
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeCreateCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeCreateCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackNormalVector:
		if stage.OnAfterShiftedLeftStackNormalVectorCreateCallback != nil {
			stage.OnAfterShiftedLeftStackNormalVectorCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackOfGrowthCurve:
		if stage.OnAfterShiftedLeftStackOfGrowthCurveCreateCallback != nil {
			stage.OnAfterShiftedLeftStackOfGrowthCurveCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackOfNormalVector:
		if stage.OnAfterShiftedLeftStackOfNormalVectorCreateCallback != nil {
			stage.OnAfterShiftedLeftStackOfNormalVectorCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedRightGrowthCurve2DRibbon:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfGrowthCurve2D:
		if stage.OnAfterStackOfGrowthCurve2DCreateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfGrowthCurve2DByGrowthVector:
		if stage.OnAfterStackOfGrowthCurve2DByGrowthVectorCreateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DByGrowthVectorCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfGrowthCurve2DRibbon:
		if stage.OnAfterStackOfGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfPartiallyRotatedTorusShape:
		if stage.OnAfterStackOfPartiallyRotatedTorusShapeCreateCallback != nil {
			stage.OnAfterStackOfPartiallyRotatedTorusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfRotatedGrowthCurve2D:
		if stage.OnAfterStackOfRotatedGrowthCurve2DCreateCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfRotatedGrowthCurve2DRibbon:
		if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeCreateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeCreateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartArcShape:
		if stage.OnAfterStartArcShapeCreateCallback != nil {
			stage.OnAfterStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartArcShapeGrid:
		if stage.OnAfterStartArcShapeGridCreateCallback != nil {
			stage.OnAfterStartArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartHalfwayArcShape:
		if stage.OnAfterStartHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartHalfwayArcShapeGrid:
		if stage.OnAfterStartHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterStartHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *StemCylinder3DShape:
		if stage.OnAfterStemCylinder3DShapeCreateCallback != nil {
			stage.OnAfterStemCylinder3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Stool2DDiagram:
		if stage.OnAfterStool2DDiagramCreateCallback != nil {
			stage.OnAfterStool2DDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Stool3DDiagram:
		if stage.OnAfterStool3DDiagramCreateCallback != nil {
			stage.OnAfterStool3DDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *StoolAbstract:
		if stage.OnAfterStoolAbstractCreateCallback != nil {
			stage.OnAfterStoolAbstractCreateCallback.OnAfterCreate(stage, target)
		}
	case *TiledFloor3DShape:
		if stage.OnAfterTiledFloor3DShapeCreateCallback != nil {
			stage.OnAfterTiledFloor3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndArcShape:
		if stage.OnAfterTopEndArcShapeCreateCallback != nil {
			stage.OnAfterTopEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndArcShapeGrid:
		if stage.OnAfterTopEndArcShapeGridCreateCallback != nil {
			stage.OnAfterTopEndArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndHalfwayArcShape:
		if stage.OnAfterTopEndHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndHalfwayArcShapeGrid:
		if stage.OnAfterTopEndHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopGrowthCurve2D:
		if stage.OnAfterTopGrowthCurve2DCreateCallback != nil {
			stage.OnAfterTopGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopMidArcVectorShape:
		if stage.OnAfterTopMidArcVectorShapeCreateCallback != nil {
			stage.OnAfterTopMidArcVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopMidArcVectorShapeGrid:
		if stage.OnAfterTopMidArcVectorShapeGridCreateCallback != nil {
			stage.OnAfterTopMidArcVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfGrowthCurve2D:
		if stage.OnAfterTopStackOfGrowthCurve2DCreateCallback != nil {
			stage.OnAfterTopStackOfGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2D:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DCreateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeCreateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeCreateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartArcShape:
		if stage.OnAfterTopStartArcShapeCreateCallback != nil {
			stage.OnAfterTopStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartArcShapeGrid:
		if stage.OnAfterTopStartArcShapeGridCreateCallback != nil {
			stage.OnAfterTopStartArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartHalfwayArcShape:
		if stage.OnAfterTopStartHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartHalfwayArcShapeGrid:
		if stage.OnAfterTopStartHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *Torus3DShape:
		if stage.OnAfterTorus3DShapeCreateCallback != nil {
			stage.OnAfterTorus3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TorusEdge3DShape:
		if stage.OnAfterTorusEdge3DShapeCreateCallback != nil {
			stage.OnAfterTorusEdge3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TorusStackShape:
		if stage.OnAfterTorusStackShapeCreateCallback != nil {
			stage.OnAfterTorusStackShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Vase2DDiagram:
		if stage.OnAfterVase2DDiagramCreateCallback != nil {
			stage.OnAfterVase2DDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Vase3DDiagram:
		if stage.OnAfterVase3DDiagramCreateCallback != nil {
			stage.OnAfterVase3DDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *VaseAbstract:
		if stage.OnAfterVaseAbstractCreateCallback != nil {
			stage.OnAfterVaseAbstractCreateCallback.OnAfterCreate(stage, target)
		}
	case *VerticalTorusStackShape:
		if stage.OnAfterVerticalTorusStackShapeCreateCallback != nil {
			stage.OnAfterVerticalTorusStackShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *VolumeKey3DShape:
		if stage.OnAfterVolumeKey3DShapeCreateCallback != nil {
			stage.OnAfterVolumeKey3DShapeCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Angle0Shape:
		newTarget := any(new).(*Angle0Shape)
		if stage.OnAfterAngle0ShapeUpdateCallback != nil {
			stage.OnAfterAngle0ShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ArcNormalVectorShape:
		newTarget := any(new).(*ArcNormalVectorShape)
		if stage.OnAfterArcNormalVectorShapeUpdateCallback != nil {
			stage.OnAfterArcNormalVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ArcNormalVectorShapeGrid:
		newTarget := any(new).(*ArcNormalVectorShapeGrid)
		if stage.OnAfterArcNormalVectorShapeGridUpdateCallback != nil {
			stage.OnAfterArcNormalVectorShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AxesShape:
		newTarget := any(new).(*AxesShape)
		if stage.OnAfterAxesShapeUpdateCallback != nil {
			stage.OnAfterAxesShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BaseVectorShape:
		newTarget := any(new).(*BaseVectorShape)
		if stage.OnAfterBaseVectorShapeUpdateCallback != nil {
			stage.OnAfterBaseVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BaseVectorShapeGrid:
		newTarget := any(new).(*BaseVectorShapeGrid)
		if stage.OnAfterBaseVectorShapeGridUpdateCallback != nil {
			stage.OnAfterBaseVectorShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ChosenP1P2PairShape:
		newTarget := any(new).(*ChosenP1P2PairShape)
		if stage.OnAfterChosenP1P2PairShapeUpdateCallback != nil {
			stage.OnAfterChosenP1P2PairShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CircleGridShape:
		newTarget := any(new).(*CircleGridShape)
		if stage.OnAfterCircleGridShapeUpdateCallback != nil {
			stage.OnAfterCircleGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Circumference3DShape:
		newTarget := any(new).(*Circumference3DShape)
		if stage.OnAfterCircumference3DShapeUpdateCallback != nil {
			stage.OnAfterCircumference3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Clock2DDiagram:
		newTarget := any(new).(*Clock2DDiagram)
		if stage.OnAfterClock2DDiagramUpdateCallback != nil {
			stage.OnAfterClock2DDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Clock3DDiagram:
		newTarget := any(new).(*Clock3DDiagram)
		if stage.OnAfterClock3DDiagramUpdateCallback != nil {
			stage.OnAfterClock3DDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ClockAbstract:
		newTarget := any(new).(*ClockAbstract)
		if stage.OnAfterClockAbstractUpdateCallback != nil {
			stage.OnAfterClockAbstractUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ClockTopCurveShape:
		newTarget := any(new).(*ClockTopCurveShape)
		if stage.OnAfterClockTopCurveShapeUpdateCallback != nil {
			stage.OnAfterClockTopCurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CutLine3DShape:
		newTarget := any(new).(*CutLine3DShape)
		if stage.OnAfterCutLine3DShapeUpdateCallback != nil {
			stage.OnAfterCutLine3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EndArcShape:
		newTarget := any(new).(*EndArcShape)
		if stage.OnAfterEndArcShapeUpdateCallback != nil {
			stage.OnAfterEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EndArcShapeGrid:
		newTarget := any(new).(*EndArcShapeGrid)
		if stage.OnAfterEndArcShapeGridUpdateCallback != nil {
			stage.OnAfterEndArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EndHalfwayArcShape:
		newTarget := any(new).(*EndHalfwayArcShape)
		if stage.OnAfterEndHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EndHalfwayArcShapeGrid:
		newTarget := any(new).(*EndHalfwayArcShapeGrid)
		if stage.OnAfterEndHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterEndHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ExplanationTextShape:
		newTarget := any(new).(*ExplanationTextShape)
		if stage.OnAfterExplanationTextShapeUpdateCallback != nil {
			stage.OnAfterExplanationTextShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Eye3DShape:
		newTarget := any(new).(*Eye3DShape)
		if stage.OnAfterEye3DShapeUpdateCallback != nil {
			stage.OnAfterEye3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EyeCornersSampledPoints3DShape:
		newTarget := any(new).(*EyeCornersSampledPoints3DShape)
		if stage.OnAfterEyeCornersSampledPoints3DShapeUpdateCallback != nil {
			stage.OnAfterEyeCornersSampledPoints3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EyeSampledPoints3DShape:
		newTarget := any(new).(*EyeSampledPoints3DShape)
		if stage.OnAfterEyeSampledPoints3DShapeUpdateCallback != nil {
			stage.OnAfterEyeSampledPoints3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EyeSeatBottomCurveShape:
		newTarget := any(new).(*EyeSeatBottomCurveShape)
		if stage.OnAfterEyeSeatBottomCurveShapeUpdateCallback != nil {
			stage.OnAfterEyeSeatBottomCurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EyeStoolBottomCurveShape:
		newTarget := any(new).(*EyeStoolBottomCurveShape)
		if stage.OnAfterEyeStoolBottomCurveShapeUpdateCallback != nil {
			stage.OnAfterEyeStoolBottomCurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EyeVolume3DShape:
		newTarget := any(new).(*EyeVolume3DShape)
		if stage.OnAfterEyeVolume3DShapeUpdateCallback != nil {
			stage.OnAfterEyeVolume3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GridPathShape:
		newTarget := any(new).(*GridPathShape)
		if stage.OnAfterGridPathShapeUpdateCallback != nil {
			stage.OnAfterGridPathShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurve2D:
		newTarget := any(new).(*GrowthCurve2D)
		if stage.OnAfterGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurve2DRibbon:
		newTarget := any(new).(*GrowthCurve2DRibbon)
		if stage.OnAfterGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*GrowthCurve2DRibbonEndShape)
		if stage.OnAfterGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*GrowthCurve2DRibbonStartShape)
		if stage.OnAfterGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurveRhombusGridShape:
		newTarget := any(new).(*GrowthCurveRhombusGridShape)
		if stage.OnAfterGrowthCurveRhombusGridShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurveRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurveRhombusShape:
		newTarget := any(new).(*GrowthCurveRhombusShape)
		if stage.OnAfterGrowthCurveRhombusShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurveRhombusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthVectorShape:
		newTarget := any(new).(*GrowthVectorShape)
		if stage.OnAfterGrowthVectorShapeUpdateCallback != nil {
			stage.OnAfterGrowthVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *InitialRhombusGridShape:
		newTarget := any(new).(*InitialRhombusGridShape)
		if stage.OnAfterInitialRhombusGridShapeUpdateCallback != nil {
			stage.OnAfterInitialRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *InitialRhombusShape:
		newTarget := any(new).(*InitialRhombusShape)
		if stage.OnAfterInitialRhombusShapeUpdateCallback != nil {
			stage.OnAfterInitialRhombusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Key3DShape:
		newTarget := any(new).(*Key3DShape)
		if stage.OnAfterKey3DShapeUpdateCallback != nil {
			stage.OnAfterKey3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *KeyHole3DShape:
		newTarget := any(new).(*KeyHole3DShape)
		if stage.OnAfterKeyHole3DShapeUpdateCallback != nil {
			stage.OnAfterKeyHole3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *KeyHoleShape:
		newTarget := any(new).(*KeyHoleShape)
		if stage.OnAfterKeyHoleShapeUpdateCallback != nil {
			stage.OnAfterKeyHoleShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Leaves3DShape:
		newTarget := any(new).(*Leaves3DShape)
		if stage.OnAfterLeaves3DShapeUpdateCallback != nil {
			stage.OnAfterLeaves3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MidArcVectorShape:
		newTarget := any(new).(*MidArcVectorShape)
		if stage.OnAfterMidArcVectorShapeUpdateCallback != nil {
			stage.OnAfterMidArcVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MidArcVectorShapeGrid:
		newTarget := any(new).(*MidArcVectorShapeGrid)
		if stage.OnAfterMidArcVectorShapeGridUpdateCallback != nil {
			stage.OnAfterMidArcVectorShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MusicAbstract:
		newTarget := any(new).(*MusicAbstract)
		if stage.OnAfterMusicAbstractUpdateCallback != nil {
			stage.OnAfterMusicAbstractUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *OriginalPoints3DShape:
		newTarget := any(new).(*OriginalPoints3DShape)
		if stage.OnAfterOriginalPoints3DShapeUpdateCallback != nil {
			stage.OnAfterOriginalPoints3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParastichyMCurves3DShape:
		newTarget := any(new).(*ParastichyMCurves3DShape)
		if stage.OnAfterParastichyMCurves3DShapeUpdateCallback != nil {
			stage.OnAfterParastichyMCurves3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ParastichyNCurves3DShape:
		newTarget := any(new).(*ParastichyNCurves3DShape)
		if stage.OnAfterParastichyNCurves3DShapeUpdateCallback != nil {
			stage.OnAfterParastichyNCurves3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DRibbon:
		newTarget := any(new).(*PartiallyGrowthCurve2DRibbon)
		if stage.OnAfterPartiallyGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DRibbonEndShape)
		if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DRibbonStartShape)
		if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectory:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectory)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP1P2)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2UpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyRotatedSeatBottomCurveShape:
		newTarget := any(new).(*PartiallyRotatedSeatBottomCurveShape)
		if stage.OnAfterPartiallyRotatedSeatBottomCurveShapeUpdateCallback != nil {
			stage.OnAfterPartiallyRotatedSeatBottomCurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyRotatedSeatTopCurveShape:
		newTarget := any(new).(*PartiallyRotatedSeatTopCurveShape)
		if stage.OnAfterPartiallyRotatedSeatTopCurveShapeUpdateCallback != nil {
			stage.OnAfterPartiallyRotatedSeatTopCurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyRotatedTorusShape:
		newTarget := any(new).(*PartiallyRotatedTorusShape)
		if stage.OnAfterPartiallyRotatedTorusShapeUpdateCallback != nil {
			stage.OnAfterPartiallyRotatedTorusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PerpendicularVector:
		newTarget := any(new).(*PerpendicularVector)
		if stage.OnAfterPerpendicularVectorUpdateCallback != nil {
			stage.OnAfterPerpendicularVectorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PerpendicularVectorGrid:
		newTarget := any(new).(*PerpendicularVectorGrid)
		if stage.OnAfterPerpendicularVectorGridUpdateCallback != nil {
			stage.OnAfterPerpendicularVectorGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PerpendicularVectorGridHalfway:
		newTarget := any(new).(*PerpendicularVectorGridHalfway)
		if stage.OnAfterPerpendicularVectorGridHalfwayUpdateCallback != nil {
			stage.OnAfterPerpendicularVectorGridHalfwayUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PerpendicularVectorHalfway:
		newTarget := any(new).(*PerpendicularVectorHalfway)
		if stage.OnAfterPerpendicularVectorHalfwayUpdateCallback != nil {
			stage.OnAfterPerpendicularVectorHalfwayUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Plant2DDiagram:
		newTarget := any(new).(*Plant2DDiagram)
		if stage.OnAfterPlant2DDiagramUpdateCallback != nil {
			stage.OnAfterPlant2DDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Plant3DDiagram:
		newTarget := any(new).(*Plant3DDiagram)
		if stage.OnAfterPlant3DDiagramUpdateCallback != nil {
			stage.OnAfterPlant3DDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PlantAbstract:
		newTarget := any(new).(*PlantAbstract)
		if stage.OnAfterPlantAbstractUpdateCallback != nil {
			stage.OnAfterPlantAbstractUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PlantCircumferenceShape:
		newTarget := any(new).(*PlantCircumferenceShape)
		if stage.OnAfterPlantCircumferenceShapeUpdateCallback != nil {
			stage.OnAfterPlantCircumferenceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PointsAndLines3DShape:
		newTarget := any(new).(*PointsAndLines3DShape)
		if stage.OnAfterPointsAndLines3DShapeUpdateCallback != nil {
			stage.OnAfterPointsAndLines3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PxShape:
		newTarget := any(new).(*PxShape)
		if stage.OnAfterPxShapeUpdateCallback != nil {
			stage.OnAfterPxShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Rendered3DShape:
		newTarget := any(new).(*Rendered3DShape)
		if stage.OnAfterRendered3DShapeUpdateCallback != nil {
			stage.OnAfterRendered3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RhombusShape:
		newTarget := any(new).(*RhombusShape)
		if stage.OnAfterRhombusShapeUpdateCallback != nil {
			stage.OnAfterRhombusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RhombusStuff:
		newTarget := any(new).(*RhombusStuff)
		if stage.OnAfterRhombusStuffUpdateCallback != nil {
			stage.OnAfterRhombusStuffUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RotatedRhombusGridShape:
		newTarget := any(new).(*RotatedRhombusGridShape)
		if stage.OnAfterRotatedRhombusGridShapeUpdateCallback != nil {
			stage.OnAfterRotatedRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RotatedRhombusShape:
		newTarget := any(new).(*RotatedRhombusShape)
		if stage.OnAfterRotatedRhombusShapeUpdateCallback != nil {
			stage.OnAfterRotatedRhombusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RotatedSampledPoints3DShape:
		newTarget := any(new).(*RotatedSampledPoints3DShape)
		if stage.OnAfterRotatedSampledPoints3DShapeUpdateCallback != nil {
			stage.OnAfterRotatedSampledPoints3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RotatedSeatAndLegs3DShape:
		newTarget := any(new).(*RotatedSeatAndLegs3DShape)
		if stage.OnAfterRotatedSeatAndLegs3DShapeUpdateCallback != nil {
			stage.OnAfterRotatedSeatAndLegs3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SampledPoints3DShape:
		newTarget := any(new).(*SampledPoints3DShape)
		if stage.OnAfterSampledPoints3DShapeUpdateCallback != nil {
			stage.OnAfterSampledPoints3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Seat3DShape:
		newTarget := any(new).(*Seat3DShape)
		if stage.OnAfterSeat3DShapeUpdateCallback != nil {
			stage.OnAfterSeat3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SeatAndLegs3DShape:
		newTarget := any(new).(*SeatAndLegs3DShape)
		if stage.OnAfterSeatAndLegs3DShapeUpdateCallback != nil {
			stage.OnAfterSeatAndLegs3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SeatBottomCurveShape:
		newTarget := any(new).(*SeatBottomCurveShape)
		if stage.OnAfterSeatBottomCurveShapeUpdateCallback != nil {
			stage.OnAfterSeatBottomCurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SeatTopCurveShape:
		newTarget := any(new).(*SeatTopCurveShape)
		if stage.OnAfterSeatTopCurveShapeUpdateCallback != nil {
			stage.OnAfterSeatTopCurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedBottomTopStartArcShape:
		newTarget := any(new).(*ShiftedBottomTopStartArcShape)
		if stage.OnAfterShiftedBottomTopStartArcShapeUpdateCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedBottomTopStartArcShapeGrid:
		newTarget := any(new).(*ShiftedBottomTopStartArcShapeGrid)
		if stage.OnAfterShiftedBottomTopStartArcShapeGridUpdateCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftGrowthCurve2DRibbon:
		newTarget := any(new).(*ShiftedLeftGrowthCurve2DRibbon)
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*ShiftedLeftGrowthCurve2DRibbonEndShape)
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*ShiftedLeftGrowthCurve2DRibbonStartShape)
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		newTarget := any(new).(*ShiftedLeftPartiallyGrowthCurve2DRibbon)
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		newTarget := any(new).(*ShiftedLeftStackGrowthCurveEndArcShape)
		if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		newTarget := any(new).(*ShiftedLeftStackGrowthCurveStartArcShape)
		if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackNormalVector:
		newTarget := any(new).(*ShiftedLeftStackNormalVector)
		if stage.OnAfterShiftedLeftStackNormalVectorUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackNormalVectorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackOfGrowthCurve:
		newTarget := any(new).(*ShiftedLeftStackOfGrowthCurve)
		if stage.OnAfterShiftedLeftStackOfGrowthCurveUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackOfGrowthCurveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackOfNormalVector:
		newTarget := any(new).(*ShiftedLeftStackOfNormalVector)
		if stage.OnAfterShiftedLeftStackOfNormalVectorUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackOfNormalVectorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedRightGrowthCurve2DRibbon:
		newTarget := any(new).(*ShiftedRightGrowthCurve2DRibbon)
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*ShiftedRightGrowthCurve2DRibbonEndShape)
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*ShiftedRightGrowthCurve2DRibbonStartShape)
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurve2DEndHalfwayArcShape:
		newTarget := any(new).(*StackGrowthCurve2DEndHalfwayArcShape)
		if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*StackGrowthCurve2DRibbonEndShape)
		if stage.OnAfterStackGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*StackGrowthCurve2DRibbonStartShape)
		if stage.OnAfterStackGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurve2DStartHalfwayArcShape:
		newTarget := any(new).(*StackGrowthCurve2DStartHalfwayArcShape)
		if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfGrowthCurve2D:
		newTarget := any(new).(*StackOfGrowthCurve2D)
		if stage.OnAfterStackOfGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfGrowthCurve2DByGrowthVector:
		newTarget := any(new).(*StackOfGrowthCurve2DByGrowthVector)
		if stage.OnAfterStackOfGrowthCurve2DByGrowthVectorUpdateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DByGrowthVectorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfGrowthCurve2DRibbon:
		newTarget := any(new).(*StackOfGrowthCurve2DRibbon)
		if stage.OnAfterStackOfGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfPartiallyRotatedTorusShape:
		newTarget := any(new).(*StackOfPartiallyRotatedTorusShape)
		if stage.OnAfterStackOfPartiallyRotatedTorusShapeUpdateCallback != nil {
			stage.OnAfterStackOfPartiallyRotatedTorusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfRotatedGrowthCurve2D:
		newTarget := any(new).(*StackOfRotatedGrowthCurve2D)
		if stage.OnAfterStackOfRotatedGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfRotatedGrowthCurve2DRibbon:
		newTarget := any(new).(*StackOfRotatedGrowthCurve2DRibbon)
		if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackRotatedGrowthCurve2DEndArcShape:
		newTarget := any(new).(*StackRotatedGrowthCurve2DEndArcShape)
		if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeUpdateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*StackRotatedGrowthCurve2DRibbonEndShape)
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*StackRotatedGrowthCurve2DRibbonStartShape)
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackRotatedGrowthCurve2DStartArcShape:
		newTarget := any(new).(*StackRotatedGrowthCurve2DStartArcShape)
		if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeUpdateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StartArcShape:
		newTarget := any(new).(*StartArcShape)
		if stage.OnAfterStartArcShapeUpdateCallback != nil {
			stage.OnAfterStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StartArcShapeGrid:
		newTarget := any(new).(*StartArcShapeGrid)
		if stage.OnAfterStartArcShapeGridUpdateCallback != nil {
			stage.OnAfterStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StartHalfwayArcShape:
		newTarget := any(new).(*StartHalfwayArcShape)
		if stage.OnAfterStartHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StartHalfwayArcShapeGrid:
		newTarget := any(new).(*StartHalfwayArcShapeGrid)
		if stage.OnAfterStartHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterStartHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StemCylinder3DShape:
		newTarget := any(new).(*StemCylinder3DShape)
		if stage.OnAfterStemCylinder3DShapeUpdateCallback != nil {
			stage.OnAfterStemCylinder3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Stool2DDiagram:
		newTarget := any(new).(*Stool2DDiagram)
		if stage.OnAfterStool2DDiagramUpdateCallback != nil {
			stage.OnAfterStool2DDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Stool3DDiagram:
		newTarget := any(new).(*Stool3DDiagram)
		if stage.OnAfterStool3DDiagramUpdateCallback != nil {
			stage.OnAfterStool3DDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StoolAbstract:
		newTarget := any(new).(*StoolAbstract)
		if stage.OnAfterStoolAbstractUpdateCallback != nil {
			stage.OnAfterStoolAbstractUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TiledFloor3DShape:
		newTarget := any(new).(*TiledFloor3DShape)
		if stage.OnAfterTiledFloor3DShapeUpdateCallback != nil {
			stage.OnAfterTiledFloor3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopEndArcShape:
		newTarget := any(new).(*TopEndArcShape)
		if stage.OnAfterTopEndArcShapeUpdateCallback != nil {
			stage.OnAfterTopEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopEndArcShapeGrid:
		newTarget := any(new).(*TopEndArcShapeGrid)
		if stage.OnAfterTopEndArcShapeGridUpdateCallback != nil {
			stage.OnAfterTopEndArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopEndHalfwayArcShape:
		newTarget := any(new).(*TopEndHalfwayArcShape)
		if stage.OnAfterTopEndHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopEndHalfwayArcShapeGrid:
		newTarget := any(new).(*TopEndHalfwayArcShapeGrid)
		if stage.OnAfterTopEndHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopGrowthCurve2D:
		newTarget := any(new).(*TopGrowthCurve2D)
		if stage.OnAfterTopGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterTopGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopMidArcVectorShape:
		newTarget := any(new).(*TopMidArcVectorShape)
		if stage.OnAfterTopMidArcVectorShapeUpdateCallback != nil {
			stage.OnAfterTopMidArcVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopMidArcVectorShapeGrid:
		newTarget := any(new).(*TopMidArcVectorShapeGrid)
		if stage.OnAfterTopMidArcVectorShapeGridUpdateCallback != nil {
			stage.OnAfterTopMidArcVectorShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		newTarget := any(new).(*TopStackGrowthCurve2DEndHalfwayArcShape)
		if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		newTarget := any(new).(*TopStackGrowthCurve2DStartHalfwayArcShape)
		if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfGrowthCurve2D:
		newTarget := any(new).(*TopStackOfGrowthCurve2D)
		if stage.OnAfterTopStackOfGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterTopStackOfGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfRotatedGrowthCurve2D:
		newTarget := any(new).(*TopStackOfRotatedGrowthCurve2D)
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		newTarget := any(new).(*TopStackOfRotatedGrowthCurve2DEndArcShape)
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		newTarget := any(new).(*TopStackOfRotatedGrowthCurve2DStartArcShape)
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStartArcShape:
		newTarget := any(new).(*TopStartArcShape)
		if stage.OnAfterTopStartArcShapeUpdateCallback != nil {
			stage.OnAfterTopStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStartArcShapeGrid:
		newTarget := any(new).(*TopStartArcShapeGrid)
		if stage.OnAfterTopStartArcShapeGridUpdateCallback != nil {
			stage.OnAfterTopStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStartHalfwayArcShape:
		newTarget := any(new).(*TopStartHalfwayArcShape)
		if stage.OnAfterTopStartHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStartHalfwayArcShapeGrid:
		newTarget := any(new).(*TopStartHalfwayArcShapeGrid)
		if stage.OnAfterTopStartHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Torus3DShape:
		newTarget := any(new).(*Torus3DShape)
		if stage.OnAfterTorus3DShapeUpdateCallback != nil {
			stage.OnAfterTorus3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TorusEdge3DShape:
		newTarget := any(new).(*TorusEdge3DShape)
		if stage.OnAfterTorusEdge3DShapeUpdateCallback != nil {
			stage.OnAfterTorusEdge3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TorusStackShape:
		newTarget := any(new).(*TorusStackShape)
		if stage.OnAfterTorusStackShapeUpdateCallback != nil {
			stage.OnAfterTorusStackShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vase2DDiagram:
		newTarget := any(new).(*Vase2DDiagram)
		if stage.OnAfterVase2DDiagramUpdateCallback != nil {
			stage.OnAfterVase2DDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Vase3DDiagram:
		newTarget := any(new).(*Vase3DDiagram)
		if stage.OnAfterVase3DDiagramUpdateCallback != nil {
			stage.OnAfterVase3DDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *VaseAbstract:
		newTarget := any(new).(*VaseAbstract)
		if stage.OnAfterVaseAbstractUpdateCallback != nil {
			stage.OnAfterVaseAbstractUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *VerticalTorusStackShape:
		newTarget := any(new).(*VerticalTorusStackShape)
		if stage.OnAfterVerticalTorusStackShapeUpdateCallback != nil {
			stage.OnAfterVerticalTorusStackShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *VolumeKey3DShape:
		newTarget := any(new).(*VolumeKey3DShape)
		if stage.OnAfterVolumeKey3DShapeUpdateCallback != nil {
			stage.OnAfterVolumeKey3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Angle0Shape:
		if stage.OnAfterAngle0ShapeDeleteCallback != nil {
			staged := any(staged).(*Angle0Shape)
			stage.OnAfterAngle0ShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ArcNormalVectorShape:
		if stage.OnAfterArcNormalVectorShapeDeleteCallback != nil {
			staged := any(staged).(*ArcNormalVectorShape)
			stage.OnAfterArcNormalVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ArcNormalVectorShapeGrid:
		if stage.OnAfterArcNormalVectorShapeGridDeleteCallback != nil {
			staged := any(staged).(*ArcNormalVectorShapeGrid)
			stage.OnAfterArcNormalVectorShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AxesShape:
		if stage.OnAfterAxesShapeDeleteCallback != nil {
			staged := any(staged).(*AxesShape)
			stage.OnAfterAxesShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BaseVectorShape:
		if stage.OnAfterBaseVectorShapeDeleteCallback != nil {
			staged := any(staged).(*BaseVectorShape)
			stage.OnAfterBaseVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BaseVectorShapeGrid:
		if stage.OnAfterBaseVectorShapeGridDeleteCallback != nil {
			staged := any(staged).(*BaseVectorShapeGrid)
			stage.OnAfterBaseVectorShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ChosenP1P2PairShape:
		if stage.OnAfterChosenP1P2PairShapeDeleteCallback != nil {
			staged := any(staged).(*ChosenP1P2PairShape)
			stage.OnAfterChosenP1P2PairShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeDeleteCallback != nil {
			staged := any(staged).(*CircleGridShape)
			stage.OnAfterCircleGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Circumference3DShape:
		if stage.OnAfterCircumference3DShapeDeleteCallback != nil {
			staged := any(staged).(*Circumference3DShape)
			stage.OnAfterCircumference3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Clock2DDiagram:
		if stage.OnAfterClock2DDiagramDeleteCallback != nil {
			staged := any(staged).(*Clock2DDiagram)
			stage.OnAfterClock2DDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Clock3DDiagram:
		if stage.OnAfterClock3DDiagramDeleteCallback != nil {
			staged := any(staged).(*Clock3DDiagram)
			stage.OnAfterClock3DDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ClockAbstract:
		if stage.OnAfterClockAbstractDeleteCallback != nil {
			staged := any(staged).(*ClockAbstract)
			stage.OnAfterClockAbstractDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ClockTopCurveShape:
		if stage.OnAfterClockTopCurveShapeDeleteCallback != nil {
			staged := any(staged).(*ClockTopCurveShape)
			stage.OnAfterClockTopCurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CutLine3DShape:
		if stage.OnAfterCutLine3DShapeDeleteCallback != nil {
			staged := any(staged).(*CutLine3DShape)
			stage.OnAfterCutLine3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EndArcShape:
		if stage.OnAfterEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*EndArcShape)
			stage.OnAfterEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EndArcShapeGrid:
		if stage.OnAfterEndArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*EndArcShapeGrid)
			stage.OnAfterEndArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EndHalfwayArcShape:
		if stage.OnAfterEndHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*EndHalfwayArcShape)
			stage.OnAfterEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EndHalfwayArcShapeGrid:
		if stage.OnAfterEndHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*EndHalfwayArcShapeGrid)
			stage.OnAfterEndHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ExplanationTextShape:
		if stage.OnAfterExplanationTextShapeDeleteCallback != nil {
			staged := any(staged).(*ExplanationTextShape)
			stage.OnAfterExplanationTextShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Eye3DShape:
		if stage.OnAfterEye3DShapeDeleteCallback != nil {
			staged := any(staged).(*Eye3DShape)
			stage.OnAfterEye3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EyeCornersSampledPoints3DShape:
		if stage.OnAfterEyeCornersSampledPoints3DShapeDeleteCallback != nil {
			staged := any(staged).(*EyeCornersSampledPoints3DShape)
			stage.OnAfterEyeCornersSampledPoints3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EyeSampledPoints3DShape:
		if stage.OnAfterEyeSampledPoints3DShapeDeleteCallback != nil {
			staged := any(staged).(*EyeSampledPoints3DShape)
			stage.OnAfterEyeSampledPoints3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EyeSeatBottomCurveShape:
		if stage.OnAfterEyeSeatBottomCurveShapeDeleteCallback != nil {
			staged := any(staged).(*EyeSeatBottomCurveShape)
			stage.OnAfterEyeSeatBottomCurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EyeStoolBottomCurveShape:
		if stage.OnAfterEyeStoolBottomCurveShapeDeleteCallback != nil {
			staged := any(staged).(*EyeStoolBottomCurveShape)
			stage.OnAfterEyeStoolBottomCurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EyeVolume3DShape:
		if stage.OnAfterEyeVolume3DShapeDeleteCallback != nil {
			staged := any(staged).(*EyeVolume3DShape)
			stage.OnAfterEyeVolume3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeDeleteCallback != nil {
			staged := any(staged).(*GridPathShape)
			stage.OnAfterGridPathShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurve2D:
		if stage.OnAfterGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*GrowthCurve2D)
			stage.OnAfterGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurve2DRibbon:
		if stage.OnAfterGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*GrowthCurve2DRibbon)
			stage.OnAfterGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurve2DRibbonEndShape:
		if stage.OnAfterGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurve2DRibbonEndShape)
			stage.OnAfterGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurve2DRibbonStartShape:
		if stage.OnAfterGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurve2DRibbonStartShape)
			stage.OnAfterGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurveRhombusGridShape:
		if stage.OnAfterGrowthCurveRhombusGridShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurveRhombusGridShape)
			stage.OnAfterGrowthCurveRhombusGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurveRhombusShape:
		if stage.OnAfterGrowthCurveRhombusShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurveRhombusShape)
			stage.OnAfterGrowthCurveRhombusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthVectorShape:
		if stage.OnAfterGrowthVectorShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthVectorShape)
			stage.OnAfterGrowthVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *InitialRhombusGridShape:
		if stage.OnAfterInitialRhombusGridShapeDeleteCallback != nil {
			staged := any(staged).(*InitialRhombusGridShape)
			stage.OnAfterInitialRhombusGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *InitialRhombusShape:
		if stage.OnAfterInitialRhombusShapeDeleteCallback != nil {
			staged := any(staged).(*InitialRhombusShape)
			stage.OnAfterInitialRhombusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Key3DShape:
		if stage.OnAfterKey3DShapeDeleteCallback != nil {
			staged := any(staged).(*Key3DShape)
			stage.OnAfterKey3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *KeyHole3DShape:
		if stage.OnAfterKeyHole3DShapeDeleteCallback != nil {
			staged := any(staged).(*KeyHole3DShape)
			stage.OnAfterKeyHole3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *KeyHoleShape:
		if stage.OnAfterKeyHoleShapeDeleteCallback != nil {
			staged := any(staged).(*KeyHoleShape)
			stage.OnAfterKeyHoleShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Leaves3DShape:
		if stage.OnAfterLeaves3DShapeDeleteCallback != nil {
			staged := any(staged).(*Leaves3DShape)
			stage.OnAfterLeaves3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MidArcVectorShape:
		if stage.OnAfterMidArcVectorShapeDeleteCallback != nil {
			staged := any(staged).(*MidArcVectorShape)
			stage.OnAfterMidArcVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MidArcVectorShapeGrid:
		if stage.OnAfterMidArcVectorShapeGridDeleteCallback != nil {
			staged := any(staged).(*MidArcVectorShapeGrid)
			stage.OnAfterMidArcVectorShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MusicAbstract:
		if stage.OnAfterMusicAbstractDeleteCallback != nil {
			staged := any(staged).(*MusicAbstract)
			stage.OnAfterMusicAbstractDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *OriginalPoints3DShape:
		if stage.OnAfterOriginalPoints3DShapeDeleteCallback != nil {
			staged := any(staged).(*OriginalPoints3DShape)
			stage.OnAfterOriginalPoints3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParastichyMCurves3DShape:
		if stage.OnAfterParastichyMCurves3DShapeDeleteCallback != nil {
			staged := any(staged).(*ParastichyMCurves3DShape)
			stage.OnAfterParastichyMCurves3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ParastichyNCurves3DShape:
		if stage.OnAfterParastichyNCurves3DShapeDeleteCallback != nil {
			staged := any(staged).(*ParastichyNCurves3DShape)
			stage.OnAfterParastichyNCurves3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DRibbon:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DRibbon)
			stage.OnAfterPartiallyGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DRibbonEndShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DRibbonEndShape)
			stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DRibbonStartShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DRibbonStartShape)
			stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectory:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectory)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2DeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP1P2)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyRotatedSeatBottomCurveShape:
		if stage.OnAfterPartiallyRotatedSeatBottomCurveShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyRotatedSeatBottomCurveShape)
			stage.OnAfterPartiallyRotatedSeatBottomCurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyRotatedSeatTopCurveShape:
		if stage.OnAfterPartiallyRotatedSeatTopCurveShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyRotatedSeatTopCurveShape)
			stage.OnAfterPartiallyRotatedSeatTopCurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyRotatedTorusShape:
		if stage.OnAfterPartiallyRotatedTorusShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyRotatedTorusShape)
			stage.OnAfterPartiallyRotatedTorusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PerpendicularVector:
		if stage.OnAfterPerpendicularVectorDeleteCallback != nil {
			staged := any(staged).(*PerpendicularVector)
			stage.OnAfterPerpendicularVectorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PerpendicularVectorGrid:
		if stage.OnAfterPerpendicularVectorGridDeleteCallback != nil {
			staged := any(staged).(*PerpendicularVectorGrid)
			stage.OnAfterPerpendicularVectorGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PerpendicularVectorGridHalfway:
		if stage.OnAfterPerpendicularVectorGridHalfwayDeleteCallback != nil {
			staged := any(staged).(*PerpendicularVectorGridHalfway)
			stage.OnAfterPerpendicularVectorGridHalfwayDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PerpendicularVectorHalfway:
		if stage.OnAfterPerpendicularVectorHalfwayDeleteCallback != nil {
			staged := any(staged).(*PerpendicularVectorHalfway)
			stage.OnAfterPerpendicularVectorHalfwayDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Plant2DDiagram:
		if stage.OnAfterPlant2DDiagramDeleteCallback != nil {
			staged := any(staged).(*Plant2DDiagram)
			stage.OnAfterPlant2DDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Plant3DDiagram:
		if stage.OnAfterPlant3DDiagramDeleteCallback != nil {
			staged := any(staged).(*Plant3DDiagram)
			stage.OnAfterPlant3DDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PlantAbstract:
		if stage.OnAfterPlantAbstractDeleteCallback != nil {
			staged := any(staged).(*PlantAbstract)
			stage.OnAfterPlantAbstractDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PlantCircumferenceShape:
		if stage.OnAfterPlantCircumferenceShapeDeleteCallback != nil {
			staged := any(staged).(*PlantCircumferenceShape)
			stage.OnAfterPlantCircumferenceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PointsAndLines3DShape:
		if stage.OnAfterPointsAndLines3DShapeDeleteCallback != nil {
			staged := any(staged).(*PointsAndLines3DShape)
			stage.OnAfterPointsAndLines3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PxShape:
		if stage.OnAfterPxShapeDeleteCallback != nil {
			staged := any(staged).(*PxShape)
			stage.OnAfterPxShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Rendered3DShape:
		if stage.OnAfterRendered3DShapeDeleteCallback != nil {
			staged := any(staged).(*Rendered3DShape)
			stage.OnAfterRendered3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RhombusShape:
		if stage.OnAfterRhombusShapeDeleteCallback != nil {
			staged := any(staged).(*RhombusShape)
			stage.OnAfterRhombusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RhombusStuff:
		if stage.OnAfterRhombusStuffDeleteCallback != nil {
			staged := any(staged).(*RhombusStuff)
			stage.OnAfterRhombusStuffDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RotatedRhombusGridShape:
		if stage.OnAfterRotatedRhombusGridShapeDeleteCallback != nil {
			staged := any(staged).(*RotatedRhombusGridShape)
			stage.OnAfterRotatedRhombusGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RotatedRhombusShape:
		if stage.OnAfterRotatedRhombusShapeDeleteCallback != nil {
			staged := any(staged).(*RotatedRhombusShape)
			stage.OnAfterRotatedRhombusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RotatedSampledPoints3DShape:
		if stage.OnAfterRotatedSampledPoints3DShapeDeleteCallback != nil {
			staged := any(staged).(*RotatedSampledPoints3DShape)
			stage.OnAfterRotatedSampledPoints3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RotatedSeatAndLegs3DShape:
		if stage.OnAfterRotatedSeatAndLegs3DShapeDeleteCallback != nil {
			staged := any(staged).(*RotatedSeatAndLegs3DShape)
			stage.OnAfterRotatedSeatAndLegs3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SampledPoints3DShape:
		if stage.OnAfterSampledPoints3DShapeDeleteCallback != nil {
			staged := any(staged).(*SampledPoints3DShape)
			stage.OnAfterSampledPoints3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Seat3DShape:
		if stage.OnAfterSeat3DShapeDeleteCallback != nil {
			staged := any(staged).(*Seat3DShape)
			stage.OnAfterSeat3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SeatAndLegs3DShape:
		if stage.OnAfterSeatAndLegs3DShapeDeleteCallback != nil {
			staged := any(staged).(*SeatAndLegs3DShape)
			stage.OnAfterSeatAndLegs3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SeatBottomCurveShape:
		if stage.OnAfterSeatBottomCurveShapeDeleteCallback != nil {
			staged := any(staged).(*SeatBottomCurveShape)
			stage.OnAfterSeatBottomCurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SeatTopCurveShape:
		if stage.OnAfterSeatTopCurveShapeDeleteCallback != nil {
			staged := any(staged).(*SeatTopCurveShape)
			stage.OnAfterSeatTopCurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedBottomTopStartArcShape:
		if stage.OnAfterShiftedBottomTopStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedBottomTopStartArcShape)
			stage.OnAfterShiftedBottomTopStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedBottomTopStartArcShapeGrid:
		if stage.OnAfterShiftedBottomTopStartArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*ShiftedBottomTopStartArcShapeGrid)
			stage.OnAfterShiftedBottomTopStartArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftGrowthCurve2DRibbon:
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftGrowthCurve2DRibbon)
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftGrowthCurve2DRibbonEndShape)
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftGrowthCurve2DRibbonStartShape)
			stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbon:
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftPartiallyGrowthCurve2DRibbon)
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
			stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackGrowthCurveEndArcShape)
			stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackGrowthCurveStartArcShape)
			stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackNormalVector:
		if stage.OnAfterShiftedLeftStackNormalVectorDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackNormalVector)
			stage.OnAfterShiftedLeftStackNormalVectorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackOfGrowthCurve:
		if stage.OnAfterShiftedLeftStackOfGrowthCurveDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackOfGrowthCurve)
			stage.OnAfterShiftedLeftStackOfGrowthCurveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackOfNormalVector:
		if stage.OnAfterShiftedLeftStackOfNormalVectorDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackOfNormalVector)
			stage.OnAfterShiftedLeftStackOfNormalVectorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedRightGrowthCurve2DRibbon:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*ShiftedRightGrowthCurve2DRibbon)
			stage.OnAfterShiftedRightGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedRightGrowthCurve2DRibbonEndShape)
			stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedRightGrowthCurve2DRibbonStartShape)
			stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurve2DEndHalfwayArcShape)
			stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurve2DRibbonEndShape)
			stage.OnAfterStackGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurve2DRibbonStartShape)
			stage.OnAfterStackGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurve2DStartHalfwayArcShape)
			stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfGrowthCurve2D:
		if stage.OnAfterStackOfGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*StackOfGrowthCurve2D)
			stage.OnAfterStackOfGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfGrowthCurve2DByGrowthVector:
		if stage.OnAfterStackOfGrowthCurve2DByGrowthVectorDeleteCallback != nil {
			staged := any(staged).(*StackOfGrowthCurve2DByGrowthVector)
			stage.OnAfterStackOfGrowthCurve2DByGrowthVectorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfGrowthCurve2DRibbon:
		if stage.OnAfterStackOfGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*StackOfGrowthCurve2DRibbon)
			stage.OnAfterStackOfGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfPartiallyRotatedTorusShape:
		if stage.OnAfterStackOfPartiallyRotatedTorusShapeDeleteCallback != nil {
			staged := any(staged).(*StackOfPartiallyRotatedTorusShape)
			stage.OnAfterStackOfPartiallyRotatedTorusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfRotatedGrowthCurve2D:
		if stage.OnAfterStackOfRotatedGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*StackOfRotatedGrowthCurve2D)
			stage.OnAfterStackOfRotatedGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfRotatedGrowthCurve2DRibbon:
		if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*StackOfRotatedGrowthCurve2DRibbon)
			stage.OnAfterStackOfRotatedGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackRotatedGrowthCurve2DEndArcShape)
			stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*StackRotatedGrowthCurve2DRibbonEndShape)
			stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*StackRotatedGrowthCurve2DRibbonStartShape)
			stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackRotatedGrowthCurve2DStartArcShape)
			stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StartArcShape:
		if stage.OnAfterStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*StartArcShape)
			stage.OnAfterStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StartArcShapeGrid:
		if stage.OnAfterStartArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*StartArcShapeGrid)
			stage.OnAfterStartArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StartHalfwayArcShape:
		if stage.OnAfterStartHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*StartHalfwayArcShape)
			stage.OnAfterStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StartHalfwayArcShapeGrid:
		if stage.OnAfterStartHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*StartHalfwayArcShapeGrid)
			stage.OnAfterStartHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StemCylinder3DShape:
		if stage.OnAfterStemCylinder3DShapeDeleteCallback != nil {
			staged := any(staged).(*StemCylinder3DShape)
			stage.OnAfterStemCylinder3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Stool2DDiagram:
		if stage.OnAfterStool2DDiagramDeleteCallback != nil {
			staged := any(staged).(*Stool2DDiagram)
			stage.OnAfterStool2DDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Stool3DDiagram:
		if stage.OnAfterStool3DDiagramDeleteCallback != nil {
			staged := any(staged).(*Stool3DDiagram)
			stage.OnAfterStool3DDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StoolAbstract:
		if stage.OnAfterStoolAbstractDeleteCallback != nil {
			staged := any(staged).(*StoolAbstract)
			stage.OnAfterStoolAbstractDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TiledFloor3DShape:
		if stage.OnAfterTiledFloor3DShapeDeleteCallback != nil {
			staged := any(staged).(*TiledFloor3DShape)
			stage.OnAfterTiledFloor3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopEndArcShape:
		if stage.OnAfterTopEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopEndArcShape)
			stage.OnAfterTopEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopEndArcShapeGrid:
		if stage.OnAfterTopEndArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopEndArcShapeGrid)
			stage.OnAfterTopEndArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopEndHalfwayArcShape:
		if stage.OnAfterTopEndHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopEndHalfwayArcShape)
			stage.OnAfterTopEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopEndHalfwayArcShapeGrid:
		if stage.OnAfterTopEndHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopEndHalfwayArcShapeGrid)
			stage.OnAfterTopEndHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopGrowthCurve2D:
		if stage.OnAfterTopGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*TopGrowthCurve2D)
			stage.OnAfterTopGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopMidArcVectorShape:
		if stage.OnAfterTopMidArcVectorShapeDeleteCallback != nil {
			staged := any(staged).(*TopMidArcVectorShape)
			stage.OnAfterTopMidArcVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopMidArcVectorShapeGrid:
		if stage.OnAfterTopMidArcVectorShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopMidArcVectorShapeGrid)
			stage.OnAfterTopMidArcVectorShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackGrowthCurve2DEndHalfwayArcShape)
			stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackGrowthCurve2DStartHalfwayArcShape)
			stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfGrowthCurve2D:
		if stage.OnAfterTopStackOfGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*TopStackOfGrowthCurve2D)
			stage.OnAfterTopStackOfGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfRotatedGrowthCurve2D:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*TopStackOfRotatedGrowthCurve2D)
			stage.OnAfterTopStackOfRotatedGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackOfRotatedGrowthCurve2DEndArcShape)
			stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackOfRotatedGrowthCurve2DStartArcShape)
			stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStartArcShape:
		if stage.OnAfterTopStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStartArcShape)
			stage.OnAfterTopStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStartArcShapeGrid:
		if stage.OnAfterTopStartArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopStartArcShapeGrid)
			stage.OnAfterTopStartArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStartHalfwayArcShape:
		if stage.OnAfterTopStartHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStartHalfwayArcShape)
			stage.OnAfterTopStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStartHalfwayArcShapeGrid:
		if stage.OnAfterTopStartHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopStartHalfwayArcShapeGrid)
			stage.OnAfterTopStartHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Torus3DShape:
		if stage.OnAfterTorus3DShapeDeleteCallback != nil {
			staged := any(staged).(*Torus3DShape)
			stage.OnAfterTorus3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TorusEdge3DShape:
		if stage.OnAfterTorusEdge3DShapeDeleteCallback != nil {
			staged := any(staged).(*TorusEdge3DShape)
			stage.OnAfterTorusEdge3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TorusStackShape:
		if stage.OnAfterTorusStackShapeDeleteCallback != nil {
			staged := any(staged).(*TorusStackShape)
			stage.OnAfterTorusStackShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vase2DDiagram:
		if stage.OnAfterVase2DDiagramDeleteCallback != nil {
			staged := any(staged).(*Vase2DDiagram)
			stage.OnAfterVase2DDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Vase3DDiagram:
		if stage.OnAfterVase3DDiagramDeleteCallback != nil {
			staged := any(staged).(*Vase3DDiagram)
			stage.OnAfterVase3DDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *VaseAbstract:
		if stage.OnAfterVaseAbstractDeleteCallback != nil {
			staged := any(staged).(*VaseAbstract)
			stage.OnAfterVaseAbstractDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *VerticalTorusStackShape:
		if stage.OnAfterVerticalTorusStackShapeDeleteCallback != nil {
			staged := any(staged).(*VerticalTorusStackShape)
			stage.OnAfterVerticalTorusStackShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *VolumeKey3DShape:
		if stage.OnAfterVolumeKey3DShapeDeleteCallback != nil {
			staged := any(staged).(*VolumeKey3DShape)
			stage.OnAfterVolumeKey3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
